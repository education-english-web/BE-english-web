package logkeeper

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var (
	singleton LogKeeper
	once      sync.Once
)

const (
	timestampFormatForLogStreamName = "20060102" // yyyymmdd
)

type cloudWatch struct {
	clwClient         *cloudwatchlogs.CloudWatchLogs
	logGroupName      *string
	logStreamName     *string
	nextSequenceToken *string
	sync.Mutex
}

// Push pushs log to the log store
func (clw *cloudWatch) Push(message string, timestamp time.Time) error {
	// PutLogEvents returns a token that will be used in the next call
	// So, if there are more than one requests running concurrently, the later one will be omitted due to invalid token
	clw.Mutex.Lock()
	defer clw.Mutex.Unlock()

	logStreamName := timestamp.Format(timestampFormatForLogStreamName)
	if clw.logStreamName == nil || *clw.logStreamName != logStreamName {
		if err := clw.validateLogStream(logStreamName); err != nil {
			return fmt.Errorf("error while validating log stream: %w", err)
		}
	}

	// put log event to cloud watch
	if err := clw.putEvent(message, timestamp); err != nil {
		return fmt.Errorf("error while putting event: %w", err)
	}

	return nil
}

func (clw *cloudWatch) validateLogStream(name string) error {
	clw.logStreamName = aws.String(name)

	logStreams, err := clw.clwClient.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        clw.logGroupName,
		LogStreamNamePrefix: aws.String(name),
		Limit:               aws.Int64(1),
	})
	if err != nil {
		return fmt.Errorf("error while describing log stream: %w", err)
	}

	if len(logStreams.LogStreams) > 0 {
		clw.nextSequenceToken = logStreams.LogStreams[0].UploadSequenceToken

		return nil
	}

	if _, err := clw.clwClient.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  clw.logGroupName,
		LogStreamName: aws.String(name),
	}); err == nil {
		return nil
	}

	aerr, ok := err.(awserr.Error)
	if !ok {
		return fmt.Errorf("error while asserting aws error: %w", err)
	}

	switch aerr.Code() {
	case cloudwatchlogs.ErrCodeResourceAlreadyExistsException:
		return clw.validateLogStream(name)
	default:
		return fmt.Errorf("unhandled error while creating log stream: %w", aerr)
	}
}

func (clw *cloudWatch) putEvent(message string, timestamp time.Time) error {
	resp, err := clw.clwClient.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message: aws.String(message),
				//nolint:gomnd
				Timestamp: aws.Int64(timestamp.UnixNano() / 1000000), // timestamp is required in millisecond format
			},
		},
		LogGroupName:  clw.logGroupName,
		LogStreamName: clw.logStreamName,
		SequenceToken: clw.nextSequenceToken,
	})
	if err == nil {
		clw.nextSequenceToken = resp.NextSequenceToken

		return nil
	}

	aerr, ok := err.(awserr.Error)
	if !ok {
		return fmt.Errorf("error while asserting aws error: %w", err)
	}

	switch aerr.Code() {
	case cloudwatchlogs.ErrCodeInvalidSequenceTokenException:
		if invalidTokenErr, ok := aerr.(*cloudwatchlogs.InvalidSequenceTokenException); ok {
			clw.nextSequenceToken = invalidTokenErr.ExpectedSequenceToken
		}

		return clw.putEvent(message, timestamp)
	default:
		return fmt.Errorf("unhandled error while putting log event: %w", aerr)
	}
}

// Init initializes a log keeper
func Init(region, logGroupName, endpointURL string) error {
	var outErr error

	config := &aws.Config{
		Region: aws.String(region),
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
			Timeout: 5 * 60 * time.Second,
		},
	}

	if endpointURL != "" {
		config.Endpoint = aws.String(endpointURL)
	}

	once.Do(
		func() {
			// create an aws session
			sess := session.Must(
				session.NewSession(config),
			)

			// create new cloud watch log client from session
			clwClient := cloudwatchlogs.New(sess)

			singleton = &cloudWatch{
				clwClient:    clwClient,
				logGroupName: aws.String(logGroupName),
			}
		},
	)

	return outErr
}

// GetCloudWatch return the singleton LogKeeper implementation
func GetCloudWatch() LogKeeper {
	return singleton
}

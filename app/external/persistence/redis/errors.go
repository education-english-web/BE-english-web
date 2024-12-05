package redis

import (
	"context"
	"errors"

	appErrors "github.com/education-english-web/BE-english-web/app/errors"
)

/*
handleError

To handle some common errors from MySQL and return our custom application errors
*/
func handleError(err error) error {
	if err != nil && errors.Is(err, context.Canceled) {
		return appErrors.ErrContextCancelled
	}

	return err
}

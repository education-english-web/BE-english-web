package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/education-english-web/BE-english-web/app/domain/repository"
	appErrs "github.com/education-english-web/BE-english-web/app/errors"
)

// csrfRepo interact with jwt stored in redis
type csrfRepo struct {
	redisClient *redis.Client
}

// NewCSRFRepository constructor of CSRFRepository
func NewCSRFRepository(redisClient *redis.Client) repository.CSRFRepository {
	return &csrfRepo{
		redisClient: redisClient,
	}
}

// Get token of given user
func (r *csrfRepo) Get(ctx context.Context, userID, uid string) (string, error) {
	key := r.formatKey(userID, uid)

	token, err := r.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", appErrs.ErrNotFound
	}

	return token, handleError(err)
}

// Save persistent a token with given expired time
func (r *csrfRepo) Save(ctx context.Context, userID, uid, value string, expiresAt time.Duration) error {
	key := r.formatKey(userID, uid)

	return handleError(r.redisClient.Set(ctx, key, value, expiresAt).Err())
}

// Delete tokens
func (r *csrfRepo) Delete(ctx context.Context, userID, uid string) error {
	key := r.formatKey(userID, uid)

	stsCmd := r.redisClient.Del(ctx, key)

	return handleError(stsCmd.Err())
}

// DeleteAll all tokens or user
func (r *csrfRepo) DeleteAll(ctx context.Context, userID string) error {
	key := r.formatKey(userID, "*")

	iter := r.redisClient.Scan(ctx, 0, key, 0).Iterator()

	for iter.Next(ctx) {
		err := r.redisClient.Del(ctx, iter.Val()).Err()
		if err != nil {
			return handleError(err)
		}
	}

	if err := iter.Err(); err != nil {
		return handleError(err)
	}

	return nil
}

func (r *csrfRepo) formatKey(userID, uid string) string {
	return fmt.Sprintf("education:user_id:%d:uid:%s", userID, uid)
}

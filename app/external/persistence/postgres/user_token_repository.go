package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
	"github.com/education-english-web/BE-english-web/app/errors"
)

// userTokenRepository interacts with postgres for users
type userTokenRepository struct {
	db *gorm.DB
}

// NewUserTokenRepository provides an instance that implements interface repository.InternalUserRepository
func NewUserTokenRepository(db *gorm.DB) repository.UserTokenRepository {
	return &userTokenRepository{
		db: db,
	}
}

// Create creates an user token
func (r *userTokenRepository) Create(ctx context.Context, userToken *entity.UserToken) error {
	return handleError(r.db.WithContext(ctx).Create(userToken).Scan(userToken).Error)
}

// Get an user token
func (r *userTokenRepository) Get(ctx context.Context, spec specifications.I) (entity.UserToken, error) {
	var userToken entity.UserToken

	tx := spec.GormQuery(r.db.WithContext(ctx))

	if err := tx.Take(&userToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userToken, errors.ErrNotFound
		}

		return userToken, handleError(err)
	}

	return userToken, nil
}

// Update updates an user token
func (r *userTokenRepository) Update(ctx context.Context, userToken *entity.UserToken) error {
	return handleError(r.db.WithContext(ctx).Save(userToken).Error)
}

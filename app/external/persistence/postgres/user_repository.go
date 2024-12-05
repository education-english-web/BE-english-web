package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
	"github.com/education-english-web/BE-english-web/app/errors"
)

// userRepository interacts with postgres for users
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository provides an instance that implements interface repository.InternalUserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create creates an internal user
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return handleError(r.db.WithContext(ctx).Create(user).Error)
}

// Get gets an internal user
func (r *userRepository) Get(ctx context.Context, spec specifications.I) (entity.User, error) {
	var user entity.User

	tx := spec.GormQuery(r.db.WithContext(ctx))

	if err := tx.Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.ErrNotFound
		}

		return user, handleError(err)
	}

	return user, nil
}

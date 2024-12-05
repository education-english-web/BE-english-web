package registry

import (
	"github.com/google/wire"

	"github.com/education-english-web/BE-english-web/app/usecases"
)

var (
	UserAddUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecases.NewUserAddUsecase,
	)

	UserMeUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecases.NewUserMeUsecase,
	)

	UserLoginUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecases.NewUserLoginUsecase,
	)

	UserRefreshTokenUsecaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		usecases.NewUserRefreshTokenUsecase,
	)
)

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package registry

import (
	"github.com/google/wire"

	"github.com/education-english-web/BE-english-web/app/usecases"
)

// InjectedInternalUserAddUsecase provides DI for the use case
func InjectedUserAddUsecase() usecases.UserAddUsecase {
	wire.Build(UserAddUsecaseSet)

	return nil
}

// InjectedInternalUserMeUsecase provides DI for the use case
func InjectedUserMeUsecase() usecases.UserMeUsecase {
	wire.Build(UserMeUsecaseSet)

	return nil
}

// InjectedInternalUserLoginUsecase provides DI for the use case
func InjectedUserLoginUsecase() usecases.UserLoginUsecase {
	wire.Build(UserLoginUsecaseSet)

	return nil
}

// InjectedInternalUserRefreshTokenUsecase provides DI for the use case
func InjectedUserRefreshTokenUsecase() usecases.RefreshTokenUsecase {
	wire.Build(UserRefreshTokenUsecaseSet)

	return nil
}

//func InjectedCSRFRepo() repository.CSRFRepository {
//	wire.Build(singletonSet, repositorySet)
//
//	return nil
//}

//func InjectedProxyLoginEventRepo() repository.ProxyLoginEventRepository {
//	wire.Build(singletonSet, repositorySet)
//
//	return nil
//}

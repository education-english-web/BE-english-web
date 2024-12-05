package registry

import (
	"github.com/google/wire"

	"github.com/education-english-web/BE-english-web/app/external/persistence/postgres"
	"github.com/education-english-web/BE-english-web/app/external/persistence/redis"
)

// Dependency Injection: All repository set for wire generate
var (
	repositorySet = wire.NewSet(
		// mysql
		postgres.NewUserRepository,
		postgres.NewUserTokenRepository,

		// redis
		redis.NewCSRFRepository,
	)
)

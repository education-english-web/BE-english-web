package framework

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/education-english-web/BE-english-web/app/config"
	_ "github.com/education-english-web/BE-english-web/app/interface/api/docs"
	"github.com/education-english-web/BE-english-web/app/interface/api/handler"
	apiMiddleware "github.com/education-english-web/BE-english-web/app/interface/api/middleware"
	"github.com/education-english-web/BE-english-web/app/services"
	"github.com/education-english-web/BE-english-web/pkg/middleware"
)

// Handler define mapping routes
// @title Edu backend
// @version 1.0
// @description This is the project of stampless team
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func Handler(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: middleware.LogFormatterJSON,
			Output:    gin.DefaultWriter,
			SkipPaths: []string{"/", "/api/healthz"},
		}),
		tracerMiddleware(cfg),
		middleware.Recovery(),
		middleware.Secure(),
		middleware.Headers,
	)

	if cfg.Env == config.ENVProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	if cfg.Env != config.ENVStaging && cfg.Env != config.ENVProduction {
		router.Use(middleware.CorsMiddleware(cfg.CORS.AllowHosts))
	}

	router.GET("/", root)
	router.GET("/api/healthz", middleware.Health)

	if cfg.Env != config.ENVHeroku {
		router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// handler
	userHandler := handler.NewUserHandler()

	//// repo for validators
	//csrfRepository := registry.InjectedCSRFRepo()
	//proxyLoginEventRepo := registry.InjectedProxyLoginEventRepo()
	// validators
	//userValidator := apiMiddleware.NewUserValidator()
	//csrfValidator := apiMiddleware.NewCSRFValidator(csrfRepository)
	//roleValidator := apiMiddleware.NewRolesValidator()
	//termsOfUseValidator := apiMiddleware.NewTermsOfUseValidator()
	//proxyLoginEventValidator := apiMiddleware.NewProxyLoginEventValidator(proxyLoginEventRepo)

	//// middlewares
	jwt := services.NewJWT(cfg.JWTSecret)
	userAuthMiddleware := apiMiddleware.NewUserAuthenticator(jwt)
	//UserOnlyMiddleware := apiMiddleware.NewCookieAuthenticator(jwt, proxyLoginEventValidator)
	//apiAuthMiddleware := apiMiddleware.NewCookieAuthenticator(
	//	jwt,
	//	userValidator,
	//	proxyLoginEventValidator,
	//	csrfValidator,
	//	roleValidator,
	//	termsOfUseValidator,
	//)

	v1NoAuth := router.Group("/api/v1")
	v1NoAuth.POST("/users/sign-up", userHandler.Add)
	v1NoAuth.POST("/users/login", userHandler.Login)
	//v1NoAuth.POST("/users/logout", userHandler.Logout)
	v1NoAuth.POST("users/refresh", userHandler.RefreshToken)

	v1Auth := router.Group("/api/v1")
	v1Auth.Use(userAuthMiddleware.Authenticate)
	v1Auth.GET("/users/me", userHandler.Me)
	return router
}

func tracerMiddleware(cfg *config.Config) gin.HandlerFunc {
	tracerEngine := cfg.TracerEngine
	switch tracerEngine {
	case "":
		return middleware.Nop()
	case middleware.TracerEngineNop:
		return middleware.Nop()
	case middleware.TracerEngineDatadog:
		return middleware.DatadogTracer(cfg.Datadog.ServiceName)
	default:
		panic(fmt.Errorf("unsupported tracer engine: %s", tracerEngine))
	}
}

func root(ctx *gin.Context) {
	type svcInfo struct {
		JSONAPI struct {
			Version string `json:"version,omitempty"`
			Name    string `json:"name,omitempty"`
		} `json:"jsonapi"`
	}

	info := svcInfo{}
	info.JSONAPI.Version = "v1"
	info.JSONAPI.Name = "EDU API"

	ctx.JSON(http.StatusOK, info)
}

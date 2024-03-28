package v1

import (
	"exam3/api-gateway_exam3/api/handlers/tokens"
	"exam3/api-gateway_exam3/config"
	"exam3/api-gateway_exam3/pkg/logger"
	"exam3/api-gateway_exam3/services"
	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	jwtHandler     tokens.JwtHandler
	enforcer       *casbin.Enforcer
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	jwtHandler     tokens.JwtHandler
	Enforcer       *casbin.Enforcer
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		jwtHandler:     c.jwtHandler,
		enforcer:       c.Enforcer,
		cfg:            c.Cfg,
	}
}

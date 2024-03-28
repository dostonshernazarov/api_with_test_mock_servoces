package main

import (
	"exam3/api-gateway_exam3/api"
	"exam3/api-gateway_exam3/config"
	"exam3/api-gateway_exam3/pkg/logger"
	"exam3/api-gateway_exam3/services"
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway_exam3")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// casbin with CSV -------------------------------------------
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatal("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy", logger.Error(err))
		return
	}
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Enforcer:       casbinEnforcer,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}

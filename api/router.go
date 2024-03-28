package api

import (
	_ "exam3/api-gateway_exam3/api/docs" // swag
	v1 "exam3/api-gateway_exam3/api/handlers/v1"
	"exam3/api-gateway_exam3/config"
	"exam3/api-gateway_exam3/pkg/logger"
	"exam3/api-gateway_exam3/services"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Enforcer       *casbin.Enforcer
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Enforcer:       option.Enforcer,
	})

	api := router.Group("/v1")
	//api.Use(casbinC.CheckCasbinPermission(option.Enforcer, option.Conf))

	// users
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.ListUsers)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)

	//order
	api.POST("/order", handlerV1.CreateUserProduct)
	api.GET("/order", handlerV1.GetAllProductUserByUserId)
	api.PUT("/order/:id", handlerV1.UpdateProductUserByID)
	api.DELETE("/order/:id", handlerV1.DeleteProductUserByID)

	//registr
	api.POST("/users/signup", handlerV1.Registr)
	api.GET("/users/verify", handlerV1.Verification)
	api.POST("/users/login", handlerV1.LogIn)
	api.GET("/users/retoken", handlerV1.RefreshAccessToken)

	//product
	api.POST("/product", handlerV1.CreateProduct)
	api.GET("/product/:id", handlerV1.GetProductByID)
	api.GET("/product", handlerV1.ListProducts)
	api.PUT("/product/:id", handlerV1.UpdateProduct)
	api.DELETE("/product/:id", handlerV1.DeleteProduct)

	// rbac
	//api.GET("/rbac/policy", handlerV1.ListAllPolicies)
	//api.GET("/rbac/roles", handlerV1.ListAllRoles)
	//api.POST("/rbac/create", handlerV1.CreateNewRole)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

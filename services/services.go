package services

import (
	"exam3/api-gateway_exam3/config"
	pbp "exam3/api-gateway_exam3/genproto/product_proto"
	pbu "exam3/api-gateway_exam3/genproto/user_proto"
	"exam3/api-gateway_exam3/mock_data"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	ProductService() pbp.ProductServiceClient
	MockService() mock_data.MockServiceClient
}

type serviceManager struct {
	userService    pbu.UserServiceClient
	productService pbp.ProductServiceClient
	mockService    mock_data.MockServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func (s *serviceManager) MockService() mock_data.MockServiceClient {
	return s.mockService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connProd, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProductServiceHost, conf.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pbu.NewUserServiceClient(connUser),
		productService: pbp.NewProductServiceClient(connProd),
		mockService:    mock_data.NewMockServiceClient(),
	}

	return serviceManager, nil
}

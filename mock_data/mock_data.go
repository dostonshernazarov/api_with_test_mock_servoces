package mock_data

import (
	"context"
	pbp "exam3/api-gateway_exam3/genproto/product_proto"
	pbu "exam3/api-gateway_exam3/genproto/user_proto"
)

type MockServiceClient interface {
	Create(ctx context.Context, user *pbu.User) (*pbu.User, error)
	CheckUniqueEmail(ctx context.Context, req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error)
	GetUserByRfshToken(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error)
	GetUserByEmail(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error)
	UpdateUser(ctx context.Context, req *pbu.User) (*pbu.User, error)
	GetUserByID(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error)
	GetAllUsers(ctx context.Context, req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersRespons, error)
	DeleteUserByID(ctx context.Context, req *pbu.IdRequest) (*pbu.DeleteUserByIDRespons, error)
	//Product
	CreateProduct(ctx context.Context, req *pbp.Product) (*pbp.Product, error)
	UpdateProduct(ctx context.Context, product *pbp.Product) (*pbp.Product, error)
	GetProductByID(ctx context.Context, request *pbp.IdRequest) (*pbp.Product, error)
	GetAllProducts(ctx context.Context, request *pbp.GetAllProductRequest) (*pbp.GetAllProdRes, error)
	DeleteProductByID(ctx context.Context, request *pbp.IdRequest) (*pbp.DeleteProductByIDRespons, error)
	CreateProductUser(ctx context.Context, product *pbp.ProductUser) (*pbp.ProductUser, error)
	GetAllProductUserByUserId(ctx context.Context, req *pbp.GetAllProductUserByUserIdReq) (*pbp.GetAllProductsRespons, error)
	UpdateProductUserByID(ctx context.Context, user *pbp.ProductUser) (*pbp.ProductUser, error)
	DeleteProductUserByID(ctx context.Context, request *pbp.IdRequest) (*pbp.DeleteProductByIDRespons, error)
}

type mockServiceClient struct {
}

func NewMockServiceClient() MockServiceClient {
	return &mockServiceClient{}
}

func (c *mockServiceClient) Create(ctx context.Context, user *pbu.User) (*pbu.User, error) {
	return user, nil
}

func (c *mockServiceClient) CheckUniqueEmail(ctx context.Context, req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error) {
	return &pbu.CheckUniqueRespons{
		IsExist: true,
	}, nil
}

func (c *mockServiceClient) GetUserByRfshToken(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error) {
	return &pbu.User{
		Id:           "b124c9a5-3ae3-4597-9540-c9c2b06ed050",
		FullName:     "Mock full name",
		City:         "Mock city",
		Email:        "Mock Email",
		Password:     "Mock password",
		Phone:        "Mock phone",
		Role:         "Mock user",
		RefreshToken: "Mock token",
	}, nil
}

func (c *mockServiceClient) GetUserByEmail(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error) {
	return &pbu.User{
		Id:           "b124c9a5-3ae3-4597-9540-c9c2b06ed050",
		FullName:     "Mock full name",
		City:         "Mock city",
		Email:        "Mock Email",
		Password:     "Mock password",
		Phone:        "Mock phone",
		Role:         "Mock user",
		RefreshToken: "Mock token",
	}, nil
}

func (c *mockServiceClient) UpdateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	return req, nil
}

func (c *mockServiceClient) GetUserByID(ctx context.Context, req *pbu.IdRequest) (*pbu.User, error) {
	return &pbu.User{
		Id:           req.Id,
		FullName:     "Mock full name",
		City:         "Mock city",
		Email:        "Mock Email",
		Password:     "Mock password",
		Phone:        "Mock phone",
		Role:         "Mock user",
		RefreshToken: "Mock token",
	}, nil
}

func (c *mockServiceClient) GetAllUsers(ctx context.Context, req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersRespons, error) {
	users := []*pbu.User{{
		Id:           "b124c9a5-3ae3-4597-9540-c9c2b06ed050",
		FullName:     "Mock full name",
		City:         "Mock city",
		Email:        "Mock Email",
		Password:     "Mock password",
		Phone:        "Mock phone",
		Role:         "Mock user",
		RefreshToken: "Mock token",
	},
		{
			Id:           "5b9fc9c7-24ac-48c3-8c55-a3cd033df8ea",
			FullName:     "Mock full name 2",
			City:         "Mock city,2",
			Email:        "Mock Email 2",
			Password:     "Mock password 2",
			Phone:        "Mock phone 2",
			Role:         "Mock user 2",
			RefreshToken: "Mock token 2",
		},
	}
	return &pbu.GetAllUsersRespons{
		User: users,
	}, nil

}

func (c *mockServiceClient) DeleteUserByID(ctx context.Context, req *pbu.IdRequest) (*pbu.DeleteUserByIDRespons, error) {
	return &pbu.DeleteUserByIDRespons{
		Result: "User has been deleted",
	}, nil
}

//Methods for products

func (c *mockServiceClient) CreateProduct(ctx context.Context, req *pbp.Product) (*pbp.Product, error) {
	return req, nil
}

func (c *mockServiceClient) UpdateProduct(ctx context.Context, product *pbp.Product) (*pbp.Product, error) {
	return product, nil
}

func (c *mockServiceClient) GetProductByID(ctx context.Context, request *pbp.IdRequest) (*pbp.Product, error) {
	return &pbp.Product{
		Id:          request.Id,
		Name:        "Mock product name",
		Description: "Mock description",
		Category:    "Mock category",
		Price:       "Mock price",
		ContactInfo: "Mock contact",
	}, nil
}

func (c *mockServiceClient) GetAllProducts(ctx context.Context, request *pbp.GetAllProductRequest) (*pbp.GetAllProdRes, error) {
	products := []*pbp.Product{
		{
			Id:          "95212dcb-3176-4eff-bb12-e35a5079ff7d",
			Name:        "Mock product name",
			Description: "Mock description",
			Category:    "Mock category",
			Price:       "Mock price",
			ContactInfo: "Mock contact",
		},
		{
			Id:          "efcd8e34-6249-4be6-8d02-e1a1667ea3b0",
			Name:        "Mock product name 2",
			Description: "Mock description 2",
			Category:    "Mock category 2",
			Price:       "Mock price 2",
			ContactInfo: "Mock contact 2",
		},
	}
	return &pbp.GetAllProdRes{
		Products: products,
	}, nil
}

func (c *mockServiceClient) DeleteProductByID(ctx context.Context, request *pbp.IdRequest) (*pbp.DeleteProductByIDRespons, error) {
	return &pbp.DeleteProductByIDRespons{
		Result: "Prodect has been deleted",
	}, nil
}

func (c *mockServiceClient) CreateProductUser(ctx context.Context, product *pbp.ProductUser) (*pbp.ProductUser, error) {
	return product, nil
}

func (c *mockServiceClient) GetAllProductUserByUserId(ctx context.Context, req *pbp.GetAllProductUserByUserIdReq) (*pbp.GetAllProductsRespons, error) {

	products := []*pbp.Product{
		{
			Id:          "95212dcb-3176-4eff-bb12-e35a5079ff7d",
			Name:        "Mock product name",
			Description: "Mock description",
			Category:    "Mock category",
			Price:       "Mock price",
			ContactInfo: "Mock contact",
		},
		{
			Id:          "efcd8e34-6249-4be6-8d02-e1a1667ea3b0",
			Name:        "Mock product name 2",
			Description: "Mock description 2",
			Category:    "Mock category 2",
			Price:       "Mock price 2",
			ContactInfo: "Mock contact 2",
		},
	}
	return &pbp.GetAllProductsRespons{
		User: &pbp.User{
			Id:           req.UserId,
			FullName:     "Mock full name",
			City:         "Mock city",
			Email:        "Mock Email",
			Password:     "Mock password",
			Phone:        "Mock phone",
			Role:         "Mock user",
			RefreshToken: "Mock token",
		},
		Products: products,
	}, nil
}

func (c *mockServiceClient) UpdateProductUserByID(ctx context.Context, user *pbp.ProductUser) (*pbp.ProductUser, error) {
	return user, nil
}

func (c *mockServiceClient) DeleteProductUserByID(ctx context.Context, request *pbp.IdRequest) (*pbp.DeleteProductByIDRespons, error) {
	return &pbp.DeleteProductByIDRespons{
		Result: "Product and user have been deleted",
	}, nil
}

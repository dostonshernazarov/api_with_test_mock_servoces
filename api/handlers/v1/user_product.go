package v1

import (
	"context"
	"exam3/api-gateway_exam3/api/handlers/models"
	"exam3/api-gateway_exam3/api/handlers/tokens"
	pbp "exam3/api-gateway_exam3/genproto/product_proto"
	pbu "exam3/api-gateway_exam3/genproto/user_proto"
	"exam3/api-gateway_exam3/pkg/etc"
	l "exam3/api-gateway_exam3/pkg/logger"
	"exam3/api-gateway_exam3/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateUser ...
// @Summary CreateUser
// @Description Api for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param Product body models.User true "CreateUser"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/ [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        models.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	if body.Id == "" {
		id := uuid.New()
		body.Id = id.String()
	}

	h.jwtHandler = tokens.JwtHandler{
		Sub:  body.Id,
		Iss:  "client",
		Role: "user",
		Log:  h.log,
	}

	access, refresh, err := h.jwtHandler.GenerateJwt()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "error while generating jwt",
		})
		h.log.Error("error generate new jwt tokens", l.Error(err))
		return
	}
	body.Password, err = etc.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Oops. Something went wrong with password",
		})
		h.log.Error("error in hash password", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	res, err := h.serviceManager.MockService().Create(ctx, &pbu.User{
		Id:           body.Id,
		FullName:     body.FullName,
		City:         "city",
		Email:        body.Email,
		Password:     body.Password,
		Phone:        body.Phone,
		Role:         "user",
		RefreshToken: refresh,
	})

	c.JSON(http.StatusOK, &models.RegisterResponseModel{
		UserID:      body.Id,
		AccessToken: access,
	})

	c.JSON(http.StatusOK, res)
}

// GetUser gets user by id
// @Summary GetUser
// @Description Api for getting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//Mock ______________
	response, err := h.serviceManager.MockService().GetUserByID(ctx, &pbu.IdRequest{Id: id})
	//Mock______________end

	//response, err := h.serviceManager.UserService().GetUserByID(ctx, &pbu.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUsers returns list of users
// @Summary ListUser
// @Description Api returns list of users
// @Tags user
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.Users
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/ [get]
func (h *handlerV1) ListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_________
	response, err := h.serviceManager.MockService().GetAllUsers(
		ctx, &pbu.GetAllUsersRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	//mock________end

	//response, err := h.serviceManager.UserService().GetAllUsers(
	//	ctx, &pbu.GetAllUsersRequest{
	//		Limit: params.Limit,
	//		Page:  params.Page,
	//	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary UpdateUser
// @Description Api returns updates user
// @Tags user
// @Accept json
// @Produce json
// @Succes 200 {Object} models.User
// @Param id path string true "ID"
// @Param User body models.User true "UpdateUser"
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pbu.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_______________
	response, err := h.serviceManager.MockService().UpdateUser(ctx, &body)
	//mock_____________end

	//response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary DeleteUser
// @Description Api deletes user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Succes 200 {Object} models.Delete
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock__________________
	response, err := h.serviceManager.MockService().DeleteUserByID(
		ctx, &pbu.IdRequest{
			Id: guid,
		})
	//mock________________end

	//response, err := h.serviceManager.UserService().DeleteUserByID(
	//	ctx, &pbu.IdRequest{
	//		Id: guid,
	//	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateUserProduct ...
// @Summary CreateUserProduct
// @Description Api for creating a new user order product
// @Tags order
// @Accept json
// @Produce json
// @Param User body models.UserOrder true "createUserOrder"
// @Success 200 {object} string
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/order/ [post]
func (h *handlerV1) CreateUserProduct(c *gin.Context) {
	var (
		body        models.UserOrder
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	if body.Id == "" {
		id := uuid.New()
		body.Id = id.String()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_____________
	_, err = h.serviceManager.MockService().CreateProductUser(ctx, &pbp.ProductUser{
		Id:        body.Id,
		ProductId: body.ProductId,
		UserId:    body.UserId,
	})
	//mock__________end

	//_, err = h.serviceManager.ProductService().CreateProductUser(ctx, &pbp.ProductUser{
	//	Id:        body.Id,
	//	ProductId: body.ProductId,
	//	UserId:    body.UserId,
	//})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, "User order product")
}

// GetAllProductUserByUserId returns list of products with user
// @Summary GetAllProductUserByUserId
// @Description Api returns list of users
// @Tags order
// @Accept json
// @Produce json
// @Param userID query string true "userID"
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.UserProducts
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/order [get]
func (h *handlerV1) GetAllProductUserByUserId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	user_id := c.Query("userID")
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock___________
	response, err := h.serviceManager.MockService().GetAllProductUserByUserId(ctx, &pbp.GetAllProductUserByUserIdReq{
		UserId: user_id,
		Page:   params.Page,
		Limit:  params.Limit,
	})
	//mock___________end

	//response, err := h.serviceManager.ProductService().GetAllProductUserByUserId(ctx, &pbp.GetAllProductUserByUserIdReq{
	//	UserId: user_id,
	//	Page:   params.Page,
	//	Limit:  params.Limit,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all product user", l.Error(err))
		return
	}

	userRes, err := h.serviceManager.MockService().GetUserByID(ctx, &pbu.IdRequest{Id: user_id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, userRes)
	c.JSON(http.StatusOK, response)
}

// UpdateProductUserByID updates user order product by id
// @Summary UpdateProductUserByID
// @Description Api returns updates user
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param UserOrder body models.UserOrder true "UserOrder"
// @Succes 200 {Object} models.UserOrder
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/order/{id} [put]
func (h *handlerV1) UpdateProductUserByID(c *gin.Context) {
	var (
		body        models.UserOrder
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock__________
	response, err := h.serviceManager.MockService().UpdateProductUserByID(ctx, &pbp.ProductUser{
		Id:        body.Id,
		ProductId: body.ProductId,
		UserId:    body.UserId,
	})
	//mock________end

	//response, err := h.serviceManager.ProductService().UpdateProductUserByID(ctx, &pbp.ProductUser{
	//	Id:        body.Id,
	//	ProductId: body.ProductId,
	//	UserId:    body.UserId,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProductUserByID deletes user order product by id
// @Summary DeleteProductUserByID
// @Description Api deletes user
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Succes 200 {Object} models.Delete
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/order/{id} [delete]
func (h *handlerV1) DeleteProductUserByID(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	fmt.Println(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock___________
	response, err := h.serviceManager.MockService().DeleteProductUserByID(ctx, &pbp.IdRequest{Id: id})
	//mock__________end

	//response, err := h.serviceManager.ProductService().DeleteProductUserByID(ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

//Product

// CreateProduct ...
// @Summary CreateProduct
// @Description Api for creating a new user
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "CreateProduct"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/ [post]
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		body        models.Product
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	if body.Id == "" {
		id := uuid.New()
		body.Id = id.String()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock____________
	response, err := h.serviceManager.MockService().CreateProduct(ctx, &pbp.Product{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		Category:    body.Category,
		Price:       body.Price,
		ContactInfo: body.ContactInfo,
	})
	//mock_________end

	//response, err := h.serviceManager.ProductService().CreateProduct(ctx, &pbp.Product{
	//	Id:          body.Id,
	//	Name:        body.Name,
	//	Description: body.Description,
	//	Category:    body.Category,
	//	Price:       body.Price,
	//	ContactInfo: body.ContactInfo,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetProductByID gets product by id
// @Summary GetProductByID
// @Description Api for getting product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/{id} [get]
func (h *handlerV1) GetProductByID(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock__________
	response, err := h.serviceManager.MockService().GetProductByID(ctx, &pbp.IdRequest{
		Id: id,
	})
	//mockj__________end

	//response, err := h.serviceManager.ProductService().GetProductByID(ctx, &pbp.IdRequest{
	//	Id: id,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProducts returns list of products
// @Summary ListProducts
// @Description Api returns list of products
// @Tags product
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.AllProducts
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/ [get]
func (h *handlerV1) ListProducts(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock_________
	response, err := h.serviceManager.MockService().GetAllProducts(ctx, &pbp.GetAllProductRequest{
		Page:  params.Page,
		Limit: params.Limit,
	})
	//mock_________end

	//response, err := h.serviceManager.ProductService().GetAllProducts(ctx, &pbp.GetAllProductRequest{
	//	Page:  params.Page,
	//	Limit: params.Limit,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct updates product by id
// @Summary UpdateProduct
// @Description Api returns updates product
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "UpdateProduct"
// @Param id path string true "ID"
// @Succes 200 {Object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/{id} [put]
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		body        models.Product
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock___________
	response, err := h.serviceManager.MockService().UpdateProduct(ctx, &pbp.Product{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		Category:    body.Category,
		Price:       body.Price,
		ContactInfo: body.ContactInfo,
	})
	//mock___________end

	//response, err := h.serviceManager.ProductService().UpdateProduct(ctx, &pbp.Product{
	//	Id:          body.Id,
	//	Name:        body.Name,
	//	Description: body.Description,
	//	Category:    body.Category,
	//	Price:       body.Price,
	//	ContactInfo: body.ContactInfo,
	//})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct deletes product by id
// @Summary DeleteProduct
// @Description Api deletes product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Succes 200 {Object} models.Delete
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/{id} [delete]
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	//mock___________
	response, err := h.serviceManager.MockService().DeleteProductByID(ctx, &pbp.IdRequest{Id: id})
	//mock________end

	//response, err := h.serviceManager.ProductService().DeleteProductByID(ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

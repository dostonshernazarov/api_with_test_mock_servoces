package tests

import (
	"encoding/json"
	"exam3/api-gateway_exam3/test_api/handler"
	"exam3/api-gateway_exam3/test_api/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// USER
	require.NoError(t, SetupMinimumInstance(""))
	file, err := OpenFile("user.json")

	require.NoError(t, err)
	req := NewRequest(http.MethodPost, "/users", file)
	res := httptest.NewRecorder()
	r := gin.Default()

	r.POST("/users", handler.CreateUser)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var user *storage.User

	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	require.Equal(t, user.FullName, "Doston Shernazarov")
	require.Equal(t, user.Phone, "916266005")
	require.Equal(t, user.Email, "dostonshernazarov2001@gmail.com")
	require.Equal(t, user.Code, "12345")
	require.NotNil(t, user.Id)

	getReq := NewRequest(http.MethodGet, "/users/get", nil)
	args := getReq.URL.Query()
	args.Add("id", user.Id)
	getReq.URL.RawQuery = args.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", handler.GetUser)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)

	var getUser *storage.User

	bdByte, err := io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getUser))
	assert.Equal(t, user.Id, getUser.Id)
	assert.Equal(t, user.Password, getUser.Password)
	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.FullName, getUser.FullName)

	ListReq := NewRequest(http.MethodGet, "/users/all", file)
	listRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/all", handler.GetAllUsers)
	r.ServeHTTP(listRes, ListReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bbLists, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bbLists)

	delReq := NewRequest(http.MethodDelete, "/users/del?id="+user.Id, file)
	delRes := httptest.NewRecorder()

	r.DELETE("/users/del", handler.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resUserB storage.ResponseMessage
	bbDel, err := io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bbDel, &resUserB))

	// REGISTR
	fileReg, err := OpenFile("registr.json")
	regReq := NewRequest(http.MethodPost, "/users/register", fileReg)
	regRes := httptest.NewRecorder()
	r.POST("/users/register", handler.RegisterUser)
	r.ServeHTTP(regRes, regReq)
	assert.Equal(t, http.StatusOK, regRes.Code)
	var resp storage.ResponseMessage
	bodyBytes, err := io.ReadAll(regRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resp))
	require.Equal(t, "We send verification password you email", resp.Content)
	require.NotNil(t, resp.Content)

	verURLCorrect := "/users/verify/12345"
	verReqCorrect := NewRequest(http.MethodGet, verURLCorrect, file)
	verResCorrect := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/verify/:code", handler.Verification)
	r.ServeHTTP(verResCorrect, verReqCorrect)

	// PRODUCT
	fileProd, err := OpenFile("product.json")
	req = NewRequest(http.MethodPost, "/products/create", fileProd)
	res = httptest.NewRecorder()
	r = gin.Default()
	r.POST("/products/create", handler.CreateProduct)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var product storage.Product
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &product))
	require.Equal(t, product.Name, "Iphone 15")
	require.Equal(t, product.Description, "Smartphone by Apple")
	require.Equal(t, product.Price, "12.3")
	require.Equal(t, product.Category, "phone")
	require.Equal(t, product.ContactInfo, "916266005")

	//get product
	getReqP := NewRequest(http.MethodGet, "/products/get", fileProd)
	q := getReqP.URL.Query()
	q.Add("id", string(product.Id))
	getReqP.URL.RawQuery = q.Encode()
	getResP := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products/get", handler.GetProductByID)
	r.ServeHTTP(getResP, getReqP)
	assert.Equal(t, http.StatusOK, getResP.Code)
	var getProduct storage.Product
	bodyGetP, err := io.ReadAll(getResP.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyGetP, &getProduct))
	require.Equal(t, product.Id, getProduct.Id)
	require.Equal(t, product.Name, getProduct.Name)
	require.Equal(t, product.Description, getProduct.Description)
	require.Equal(t, product.Category, getProduct.Category)
	require.Equal(t, product.Price, getProduct.Price)
	require.Equal(t, product.ContactInfo, getProduct.ContactInfo)

	listReqq := NewRequest(http.MethodGet, "/products/all", fileProd)
	listRess := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products/all", handler.ListProducts)
	r.ServeHTTP(listRess, listReqq)
	// assert.Equal(t, http.StatusOK, listRess.Code)
	bodyBytes, err = io.ReadAll(listRess.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	delReq = NewRequest(http.MethodDelete, "/product/delete?id="+product.Id, fileProd)
	delRes = httptest.NewRecorder()
	r.DELETE("/product/delete", handler.DeleteProduct)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var content storage.ResponseMessage
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &content))
	require.Equal(t, "product has been deleted", content.Content)

	// USER PRODUCT
	fileUserProd, err := OpenFile("user_product.json")
	req = NewRequest(http.MethodPost, "/user/prod/create", fileUserProd)
	res = httptest.NewRecorder()
	r = gin.Default()
	r.POST("/user/prod/create", handler.CreateUserProduct)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var productUser storage.UserOrder
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &productUser))
	require.NotNil(t, productUser.UserId)
	require.NotNil(t, productUser.Id)
	require.NotNil(t, productUser.ProductId)

	delReqUP := NewRequest(http.MethodDelete, "/user/prod/delete?id="+product.Id, fileUserProd)
	delResUP := httptest.NewRecorder()
	r.DELETE("/user/prod/delete", handler.DeleteProductUserByID)
	r.ServeHTTP(delRes, delReqUP)
	assert.Equal(t, http.StatusOK, delResUP.Code)
	var contentR storage.ResponseMessage
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &contentR))
	require.Equal(t, "user product has been deleted", contentR.Content)
}

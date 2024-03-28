package handler

import (
	"encoding/json"
	"exam3/api-gateway_exam3/test_api/storage"
	"exam3/api-gateway_exam3/test_api/storage/kv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

//Registr

func RegisterUser(c *gin.Context) {
	var newUser storage.RegisterModel
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	if err := kv.Set(id.String(), string(userJson)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "We send verification password you email",
	})
}

func Verification(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

// User
func GetUser(c *gin.Context) {
	id := c.Query("id")

	userGet, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.User
	if err := json.Unmarshal([]byte(userGet), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func CreateUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJ, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJ)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)

}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been deleted",
	})
}

func GetAllUsers(c *gin.Context) {
	userList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*storage.User
	for _, l := range userList {
		var user storage.User

		if err := json.Unmarshal([]byte(l), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}

// Product
func CreateProduct(c *gin.Context) {
	var body storage.Product

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = uuid.NewString()

	prodJ, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(body.Id, string(prodJ)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func GetProductByID(c *gin.Context) {
	id := c.Query("id")

	prodGet, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.Product
	if err := json.Unmarshal([]byte(prodGet), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListProducts(c *gin.Context) {
	prodList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var prods []*storage.User
	for _, l := range prodList {
		var prod storage.User

		if err := json.Unmarshal([]byte(l), &prod); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		prods = append(prods, &prod)
	}

	c.JSON(http.StatusOK, prods)
}

func UpdateProduct(c *gin.Context) {
	var prod storage.Product
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	prodJ, err := json.Marshal(prod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(prod.Id, string(prodJ)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated product",
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "product has been deleted",
	})
}

// User Product

func CreateUserProduct(c *gin.Context) {
	var body storage.UserOrder

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if body.Id == "" {
		id := uuid.New()
		body.Id = id.String()
	}

	userJ, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(body.Id, string(userJ)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, body)
}

func DeleteProductUserByID(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "user product has been deleted",
	})
}

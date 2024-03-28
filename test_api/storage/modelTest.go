package storage

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

import "regexp"

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       string `json:"price"`
	ContactInfo string `json:"contact_info"`
}

type AllProducts struct {
	Products []*Product `json:"products"`
}

type User struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	RefreshToken string `json:"refresh_token"`
	Code         string `json:"code"`
}

type UserOrder struct {
	Id        string `json:"id"`
	ProductId string `json:"product_Id"`
	UserId    string `json:"user_id"`
}

type UserLoginRequest struct {
	UserNameOrEmail string `json:"user_name_or_email"`
	Password        string `json:"password"`
}

type RegisterModel struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Code     string `json:"code"`
}

type ResponseMessage struct {
	Content string `json:"content"`
}

type RegisterResponseModel struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Users struct {
	Users []*User `json:"users"`
}

type UserProducts struct {
	User     User
	Products []*Product
}
type Delete struct {
	Result string
}

func (u *RegisterModel) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.FullName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
	)
}

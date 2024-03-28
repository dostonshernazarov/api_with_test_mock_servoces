package models

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

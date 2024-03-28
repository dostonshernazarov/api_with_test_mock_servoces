package models

type Policy struct {
	Role     string `json:"role"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

type ListPolicyResponse struct {
	Policies []*Policy `json:"policies"`
}

type CreateUserRoleRequest struct {
	RoleName string `json:"role_name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

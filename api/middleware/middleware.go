package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/spf13/cast"
	"net/http"
	"strings"

	token "exam3/api-gateway_exam3/api/tokens"
	"exam3/api-gateway_exam3/config"
	"github.com/gin-gonic/gin"
)

type CasbinHandler struct {
	enforcer *casbin.Enforcer
	cnf      config.Config
}

func CheckCasbinPermission(enforce *casbin.Enforcer, cfg config.Config) gin.HandlerFunc {
	casbH := &CasbinHandler{
		enforcer: enforce,
		cnf:      cfg,
	}

	return func(ctx *gin.Context) {
		permission, err := casbH.CheckPermission(ctx.Request)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)

		}
		if !permission {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}
	}
}

// GetRole gets role from Authorization header if there is a token then it is
// parsed and in role got from role claim. If there is no token then role is
// unauthorized
func (a *CasbinHandler) GetRole(r *http.Request) (string, error) {
	var (
		t   string
		err error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", nil
	} else if strings.Contains(jwtToken, "Bearer") {
		t = strings.TrimPrefix(jwtToken, "Bearer ")
	} else {
		t = jwtToken
	}

	claims, err := token.ExtractClaim(t, []byte(a.cnf.SignInKey))
	if err != nil {
		return "unauthorized", err
	}

	return cast.ToString(claims["role"]), nil
}

// CheckPermission checks whether user is allowed to use certain endpoint
func (a *CasbinHandler) CheckPermission(r *http.Request) (bool, error) {
	user, err := a.GetRole(r)
	if err != nil {
		return false, err
	}
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

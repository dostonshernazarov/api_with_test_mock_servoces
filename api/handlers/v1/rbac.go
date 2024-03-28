package v1

import (
	"encoding/json"
	"exam3/api-gateway_exam3/api/handlers/models"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security      ApiKeyAuth
// @Summary       Get list of policeis
// @Description   This API gets list of policies
// @Tags          rbac
// @Accept        json
// @Produce       json
// @Param         role query string true "Role"
// @Succes        200 {object} models.ListPolePolicyResponse
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/rbac/policy [GET]
func (h *handlerV1) ListAllPolicies(ctx *gin.Context) {
	role := ctx.Query("role")

	var resp models.ListPolicyResponse

	for _, p := range h.enforcer.GetFilteredPolicy(0, role) {
		resp.Policies = append(resp.Policies, &models.Policy{
			Role:     p[0],
			Endpoint: p[1],
			Method:   p[2],
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Security      ApiKeyAuth
// @Summary       Get list of roles
// @Description   This API get list of roles
// @Tags          rbac
// @Accept        json
// @Produce       json
// @Param         limit query int false "limit"
// @Param         offset query int false "offset"
// @Succes        200 {object} []string
// @Failure       404 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/rbac/roles [GET]
func (h *handlerV1) ListAllRoles(ctx *gin.Context) {
	resp := h.enforcer.GetAllRoles()
	ctx.JSON(http.StatusOK, resp)
}

// @Security      ApiKeyAuth
// @Summary       Create new user-role
// @Description   Create new user-role
// @Tags          rbac
// @Accept        json
// @Produce       json
// @Param         body body models.CreateUserRoleRequest true "body"
// @Success       200 {object} models.CreateUserRoleRequest
// @Failure     404 {object} models.Error
// @Failure     500 {object} models.Error
// @Router        /v1/rbac/add-user-role [POST]
func (h *handlerV1) CreateNewRole(ctx *gin.Context) {
	var reqBody models.CreateUserRoleRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody); err != nil {
		h.log.Error("rbacHandler/CreateUserRole", zap.Error(err))
		return
	}

	if _, err := h.enforcer.AddPolicy(reqBody.RoleName, reqBody.Path, reqBody.Method); err != nil {
		h.log.Error("Error on grant access", zap.Error(err))
		return
	}
	h.enforcer.SavePolicy()
	ctx.JSON(http.StatusOK, reqBody)
}

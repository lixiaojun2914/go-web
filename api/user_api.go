package api

import (
	"code/service"
	"code/service/dto"
	"code/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tags 用户管理
// @Summary 用户登陆
// @Description 用户登陆详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError()
	if err != nil {
		return
	}

	iUser, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

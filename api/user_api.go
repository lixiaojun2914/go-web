package api

import (
	"code/service"
	"code/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrCodeAddUser     = 100011
	ErrCodeGetUserById = 100012
	ErrCodeGetUserList = 100013
	ErrCodeUpdateUser  = 100014
	ErrCodeDeleteUser  = 100015
	ErrCodeLogin       = 100016
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
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, token, err := m.Service.Login(iUserLoginDTO)
	if err == nil {
		err = service.SetLoginUserTokenToRedis(iUser.ID, token)
	}

	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ErrCodeLogin,
			Msg:    err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	//file, _ := c.FormFile("file")
	//stFilePath := fmt.Sprintf("./upload/%s", file.Filename)
	//_ = c.SaveUploadedFile(file, stFilePath)
	//iUserAddDTO.Avatar = stFilePath

	err := m.Service.AddUser(&iUserAddDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ErrCodeAddUser,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: iUserAddDTO,
	})
}

func (m UserApi) GetUserByID(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.GetUserByID(iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ErrCodeGetUserById,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: iUser,
	})
}

func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDTO}).GetError(); err != nil {
		return
	}

	giUserList, nTotal, err := m.Service.GetUserList(iUserListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ErrCodeGetUserList,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  giUserList,
		Total: nTotal,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTP dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTP, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateUser(iUserUpdateDTP)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ErrCodeUpdateUser,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{})
}

func (m UserApi) DeleteUserByID(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteUserByID(iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ErrCodeDeleteUser,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{})
}

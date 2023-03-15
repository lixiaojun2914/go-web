package api

import (
	"code/service/dto"
	"code/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tags 用户管理
// @Summary 用户登陆
// @Description 用户登陆详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	err := ctx.ShouldBind(&iUserLoginDTO)
	if err != nil {
		Fail(ctx, ResponseJson{
			Msg: parseValidateErrors(err.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}
	OK(ctx, ResponseJson{
		Data: iUserLoginDTO,
	})
}

// 自定义校验器
func parseValidateErrors(errs validator.ValidationErrors, target any) error {
	var errResult error

	// 通过反射获取指针指向元素的类型对象
	fieleds := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		fieled, _ := fieleds.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := fieled.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessageTag = fieled.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

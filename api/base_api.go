package api

import (
	"code/global"
	"code/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx     *gin.Context
	DTO     any
	BindUri bool
	BindAll bool
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	// 绑定请求上下文
	m.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindAll || option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBindUri(option.DTO))
		}

		if option.BindAll || !option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBind(option.DTO))
		}

		if errResult != nil {
			errResult = m.ParseValidateErrors(errResult, option.DTO)
			m.AddError(errResult)
			Fail(m.Ctx, ResponseJson{
				Msg: m.GetError().Error(),
			})
		}
	}
	return m
}

// 错误处理
func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}

// 自定义校验器
func (m *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}
	// 通过反射获取指针指向元素的类型对象
	fieleds := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
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

func (m *BaseApi) OK(resp ResponseJson) {
	OK(m.Ctx, resp)
}

func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(m.Ctx, resp)
}

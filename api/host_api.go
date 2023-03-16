package api

import (
	"code/service"
	"code/service/dto"
	"github.com/gin-gonic/gin"
)

type HostApi struct {
	BaseApi
	Service *service.HostService
}

func NewHostApi() HostApi {
	return HostApi{
		Service: service.NewHostService(),
	}
}

func (m HostApi) Shutdown(c *gin.Context) {
	var iShutdownHostDTO dto.ShutdownHostDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iShutdownHostDTO}).GetError()
	if err != nil {
		return
	}

	err = m.Service.Shutdown(iShutdownHostDTO)

	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Msg: "Shutdown Success",
	})
}

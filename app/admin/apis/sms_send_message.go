package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/apis"
)

type SmsSendMessage struct {
	apis.Api
}

func (e *SmsSendMessage) SendMessage(c *gin.Context) {
	req := dto.SmsSendMessageReq{}
	s := service.SmsSendMessage{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	message, _ := s.SendMessage(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(200, message)
}

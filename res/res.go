package res

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//goland:noinspection ALL
func GetStatusMessage(status int, msg string) *Message {
	return &Message{
		Status:  status,
		Message: msg,
	}
}

//goland:noinspection ALL
func GetStatusTplMessage(status int, msg string, err error) (*Message, error) {
	message := &Message{
		Status:  status,
		Message: msg,
	}
	if err != nil {
		return message, err
	}

	return message, nil
}

func GetMessage(status, code int, msg string) *Message {
	return &Message{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}

//goland:noinspection ALL
func GetTplMessage(status, code int, msg string, err error) (*Message, error) {
	message := &Message{
		Status:  status,
		Code:    code,
		Message: msg,
	}
	if err != nil {
		return message, err
	}

	return message, nil
}

type baseData struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (m *Message) SendJson(
	ctx *gin.Context,
) {
	m.setStatusCode(ctx)
	ctx.JSON(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
	})
}

func (m *Message) SendJsonData(
	ctx *gin.Context,
	data any,
) {
	m.setStatusCode(ctx)
	ctx.JSON(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
		Data: data,
	})
}

func (m *Message) SendRawJsonData(
	ctx *gin.Context,
	data any,
) {
	m.setStatusCode(ctx)
	ctx.JSON(m.Status, data)
}

func (m *Message) SendXml(
	ctx *gin.Context,
) {
	m.setStatusCode(ctx)
	ctx.XML(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
	})
}

func (m *Message) SendXmlData(
	ctx *gin.Context,
	data any,
) {
	m.setStatusCode(ctx)
	ctx.XML(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
		Data: data,
	})
}

func (m *Message) SendYaml(
	ctx *gin.Context,
) {
	m.setStatusCode(ctx)
	ctx.YAML(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
	})
}

func (m *Message) SendYamlData(
	ctx *gin.Context,
	data any,
) {
	m.setStatusCode(ctx)
	ctx.YAML(m.Status, &baseData{
		Code: m.Code,
		Msg:  m.Message,
		Data: data,
	})
}

func (m *Message) SendFile(
	ctx *gin.Context,
	filePath string,
) {
	m.setStatusCode(ctx)
	ctx.Status(m.Status)
	ctx.File(filePath)
}

func (m *Message) SendFileAttachment(
	ctx *gin.Context,
	filePath,
	fileName string,
) {
	m.setStatusCode(ctx)
	ctx.Status(m.Status)
	ctx.FileAttachment(filePath, fileName)
}

func (m *Message) SendProtoBuf(
	ctx *gin.Context,
	protobuf any,
) {
	m.setStatusCode(ctx)
	ctx.ProtoBuf(m.Status, protobuf)
}

func (m *Message) SendRedirect(
	ctx *gin.Context,
	location string,
) {
	m.setStatusCode(ctx)
	ctx.Redirect(http.StatusMovedPermanently, location)
}

func (m *Message) SendData(
	ctx *gin.Context,
	contentType string,
	data []byte,
) {
	m.setStatusCode(ctx)
	ctx.Data(m.Status, contentType, data)
}

func (m *Message) SendString(
	ctx *gin.Context,
	format string,
	values ...any,
) {
	m.setStatusCode(ctx)
	ctx.String(m.Status, format, values...)
}

func (m *Message) SendHTML(
	ctx *gin.Context,
	name string,
	obj any,
) {
	m.setStatusCode(ctx)
	ctx.HTML(m.Status, name, obj)
}

func (m *Message) setStatusCode(ctx *gin.Context) {
	ctx.Header("black_gin_status", fmt.Sprintf("%d", m.Status))
	ctx.Header("black_gin-code", fmt.Sprintf("%d", m.Code))
}

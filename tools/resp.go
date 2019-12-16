package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespFail(c *gin.Context, msg interface{}) {
	Resp(c, 0, nil, fmt.Sprint(msg))
}

func RespOk(c *gin.Context, data interface{}, msg interface{}) {
	Resp(c, 1, data, fmt.Sprint(msg))
}

func Resp(c *gin.Context, status int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"detail": data,
	})
}

package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/wagaru/task/internal/errcode"
)

func (d *delivery) ToResponse(c *gin.Context, status int, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(status, data)
}

func (d *delivery) ToErrorResponse(c *gin.Context, err error) {
	e, ok := err.(*errcode.Error)
	if !ok {
		e = errcode.UnknownError
	}

	data := gin.H{
		"code":    e.Code(),
		"message": e.Message(),
	}
	c.JSON(e.StatusCode(), data)
}

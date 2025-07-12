package helpers

import (
	"github.com/faisallbhr/gin-boilerplate/structs"
	"github.com/gin-gonic/gin"
)

// ResponseSuccess mengirimkan response sukses universal
func ResponseSuccess(c *gin.Context, data any, message string, status int, meta *structs.Meta) {
	c.JSON(status, structs.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// ResponseError mengirimkan response gagal universal
func ResponseError(c *gin.Context, message string, status int, errors map[string]string) {
	c.JSON(status, structs.ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

package response

import (
	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, message string, data any) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
}

func Paginated(
	c *gin.Context,
	status int,
	message string,
	data any,
	pagination dto.Pagination,
) {
	c.JSON(status, gin.H{
		"success":    true,
		"message":    message,
		"data":       data,
		"pagination": pagination,
	})
}

package web

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int
	Content    interface{}
}

func CheckResponse(ctx *gin.Context, response Response) {
	switch response.StatusCode {
	case 200 | 201 | 204:
		createSuccessResponse(ctx, response.StatusCode, response.Content)
	case 400 | 404 | 500:
		createErrorResponse(ctx, response.StatusCode, response.Content.(string))
	}
}

func createSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func createErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.String(statusCode, message)
}

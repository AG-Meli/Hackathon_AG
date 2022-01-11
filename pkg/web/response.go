package web

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int
	Content    interface{}
}

func CheckResponse(ctx *gin.Context, response Response) {
	switch {
	case response.StatusCode == 200 || response.StatusCode == 201 || response.StatusCode == 204:
		createSuccessResponse(ctx, response.StatusCode, response.Content)
	case response.StatusCode == 400 || response.StatusCode == 404 || response.StatusCode == 500:
		createErrorResponse(ctx, response.StatusCode, response.Content.(string))
	}
}

func createSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func createErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.String(statusCode, message)
}

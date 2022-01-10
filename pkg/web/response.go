package web

import "github.com/gin-gonic/gin"

var (
	NOT_FOUND             error
	BAD_REQUEST           error
	INTERNAL_SERVER_ERROR error
	SUCCESS_CONTENT       error
	SUCCESS_NO_CONTENT    error
	SUCCESS_CREATED       error
)

func CheckResponse(ctx *gin.Context, reference error, content interface{}) {
	switch reference {
	case NOT_FOUND:
		createErrorResponse(ctx, 404, content.(error))
	case BAD_REQUEST:
		createErrorResponse(ctx, 400, content.(error))
	case INTERNAL_SERVER_ERROR:
		createErrorResponse(ctx, 500, content.(error))
	case SUCCESS_CONTENT:
		createSuccessResponse(ctx, 200, content)
	case SUCCESS_NO_CONTENT:
		createSuccessResponse(ctx, 204, content)
	case SUCCESS_CREATED:
		createSuccessResponse(ctx, 201, content)
	}
}

func createSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func createErrorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.String(statusCode, err.Error())
}

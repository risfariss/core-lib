package response

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

type IResponse interface {
	SetPagination(queryParam url.Values) (out PaginationObject)
	SetResponseFailed(ctx *gin.Context, httpCode int, message string, err string)
	SetResponseSuccess(ctx *gin.Context, httpCode int, message string, data any, pagination any)
}

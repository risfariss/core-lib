package request

import "github.com/gin-gonic/gin"

type IRequest interface {
	CheckToken(ctx *gin.Context) (userId string, err error)
}

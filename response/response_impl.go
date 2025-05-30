package response

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

type Response struct {
	/*todo fill this block if need connection with another interface*/
}

func InitializeResponse() IResponse {
	return &Response{}
}

type PaginationObject struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Order     string `json:"order"`
	Offset    int    `json:"offset"`
	TotalData int    `json:"totalData"`
	TotalPage int    `json:"totalPage"`
}

type ResponseObject struct {
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

func (r *Response) SetPagination(queryParam url.Values) (out PaginationObject) {
	if queryParam.Get("page") != "" {
		page, err := strconv.Atoi(queryParam.Get("page"))
		if err != nil {
			page = 0
		}
		out.Page = page
	}

	if queryParam.Get("limit") != "" {
		limit, err := strconv.Atoi(queryParam.Get("limit"))
		if err != nil {
			limit = 0
		}
		out.Limit = limit
	}

	if queryParam.Get("order") != "" {
		out.Order = "ID " + queryParam.Get("order")
	}

	if out.Page == 0 {
		out.Page = 1
	}

	if out.Limit == 0 {
		out.Limit = 10
	}

	if out.Order == "" {
		out.Order = "ID DESC"
	}

	out.Offset = (out.Page - 1) * out.Limit

	return out
}

func (r *Response) SetResponseFailed(ctx *gin.Context, httpCode int, message string, err string) {
	res := ResponseObject{
		Status:  false,
		Message: message,
		Error:   err,
	}
	ctx.AbortWithStatusJSON(httpCode, res)
}

func (r *Response) SetResponseSuccess(ctx *gin.Context, httpCode int, message string, data any, pagination any) {
	res := ResponseObject{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
	ctx.JSON(httpCode, res)
}

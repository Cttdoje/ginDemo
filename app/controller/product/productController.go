package product

import (
	"fmt"
	"ginDemo/app/controller/param_bind"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(context *gin.Context) {
	value, _ := context.GetPostForm("name1")
	fmt.Println(value)

	if err := context.ShouldBind(&param_bind.ProductAdd{}); err != nil {
		context.String(http.StatusInternalServerError, "error")
	} else {
		context.String(http.StatusOK, "success")
	}
}

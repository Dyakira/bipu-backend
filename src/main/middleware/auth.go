package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"net/http"
	"fmt"
)


func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Handle the exchange code to initiate a transport.
		session := sessions.Default(ctx)

		uid := session.Get("uid")
		if uid != nil {
			log.Println(strconv.Itoa(uid.(int)) + "is login success")
			ctx.Next()
			return
		}

		ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("need login"))
		return
	}
}

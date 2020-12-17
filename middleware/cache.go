package middleware

import (
	"log"
	"net/http"

	"github.com/dgkg/keypass/cache"
	"github.com/gin-gonic/gin"
)

func NewCacheMiddleware(client cache.CacheDB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := client.Get(ctx, ctx.Request.URL.String())
		if err != nil || len(data) == 0 {
			log.Print("try to get some cache but no cache found")
			ctx.Next()
		} else {
			log.Print("cache found:", string(data))
			ctx.Header("Content-Type", "application/json")
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(data)
			ctx.Abort()
		}
	}
}

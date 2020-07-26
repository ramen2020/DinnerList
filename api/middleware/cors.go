package middleware

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

var reg = regexp.MustCompile("")

func CORS(c *gin.Context) {

	if c.Request.Method == "OPTIONS" {
		if true {
			headers := c.Request.Header.Get("Access-Control-Request-Headers")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
			c.Data(200, "text/plain", []byte{})
			c.Abort()
		} else {
			c.Data(403, "text/plain", []byte{})
			c.Abort()
		}
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}

	return
}

package middleware

import (
	"fmt"
	"memo_gin/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

//HandlerFunc 将 gin 中间件使用的处理程序定义为返回值。
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization") //请求头
		fmt.Printf("token: %v\n", token)
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token) //验证token
			fmt.Printf("claim: %v\n", claim)
			if err != nil {
				code = 403 //token无权限，是假的
				fmt.Printf("err: %v\n", err)
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //token无效
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort() //中止
			return
		}
		c.Next()
	}
}

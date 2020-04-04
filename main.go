package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"weixin-demo/util"
)

const Token = "coleliedev"

func main() {
	router := gin.Default()

	router.GET("/", WXCheckSignature)

	log.Fatalln(router.Run(":80"))
}

// WXCheckSignature 微信接入校验
func WXCheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := util.CheckSignature(signature, timestamp, nonce, Token)
	if !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)
}

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
	router.POST("/", WXMsgReceive)

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
		log.Println("[微信接入] - 微信公众号接入校验失败!")
		return
	}

	log.Println("[微信接入] - 微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)
}

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
}

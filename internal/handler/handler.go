package handler

import "github.com/gin-gonic/gin"

type ChannelHandler interface {
	Get() gin.HandlerFunc
	Post() gin.HandlerFunc
}

type LineBotHandler interface {
	Post() gin.HandlerFunc
}

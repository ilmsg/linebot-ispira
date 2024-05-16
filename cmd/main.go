package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	inMemDb "github.com/ilmsg/linebot-ispira/internal/db/in_memory"
	chHandler "github.com/ilmsg/linebot-ispira/internal/handler/channel_handler"
	linebotHandler "github.com/ilmsg/linebot-ispira/internal/handler/linebot_handler"
	inMemSrv "github.com/ilmsg/linebot-ispira/internal/services/in_memory"
)

func main() {
	fmt.Println("hi.")

	inMemoryDB := inMemDb.NewInMemoryDB()
	channelSrv := inMemSrv.NewChannelInMemoryService(inMemoryDB)

	app := gin.Default()
	groupChannel := app.Group("/channels")
	groupWebhook := app.Group("/webhook")

	gh := chHandler.NewChannelHandler(channelSrv)
	groupChannel.GET("/", gh.Get())
	groupChannel.POST("/", gh.Post())

	lh := linebotHandler.NewLinebotHandler(channelSrv)
	groupWebhook.POST("/:channelId", lh.Post()) // webhook

	app.Run(":6660")
}

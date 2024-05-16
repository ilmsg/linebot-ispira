package linebothandler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/linebot-ispira/internal/handler"
	"github.com/ilmsg/linebot-ispira/internal/services"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type linebotHandler struct {
	chSrv services.ChannelService
}

func (l *linebotHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {

		ch, err := l.chSrv.Get(c.Param("channelId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		bot, err := messaging_api.NewMessagingApiAPI(ch.Token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		cb, err := webhook.ParseRequest(ch.Secret, c.Request)
		if err != nil {
			log.Printf("Cannot parse request: %+v\n", err)
			if errors.Is(err, webhook.ErrInvalidSignature) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, event := range cb.Events {
			log.Printf("/webhook called%+v\n", event)

			switch e := event.(type) {
			case webhook.MessageEvent:
				switch message := e.Message.(type) {
				case webhook.TextMessageContent:
					if _, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								messaging_api.TextMessage{
									Text: message.Text,
								},
							},
						},
					); err != nil {
						log.Print(err)
					} else {
						log.Println("Sent text reply.")
					}
				}
			}
		}
	}
}

func NewLinebotHandler(chSrv services.ChannelService) handler.LineBotHandler {
	return &linebotHandler{chSrv}
}

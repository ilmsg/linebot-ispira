package channelhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/linebot-ispira/internal/handler"
	"github.com/ilmsg/linebot-ispira/internal/models"
	"github.com/ilmsg/linebot-ispira/internal/services"
)

type channelHandler struct {
	chSrv services.ChannelService
}

func (g *channelHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		channels, err := g.chSrv.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, channels)
	}
}

func (g *channelHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var channel *models.Channel
		if err := c.ShouldBindJSON(&channel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := g.chSrv.Create(channel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "created."})
	}
}

func NewChannelHandler(chSrv services.ChannelService) handler.ChannelHandler {
	return &channelHandler{chSrv}
}

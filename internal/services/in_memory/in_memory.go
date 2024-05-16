package inmemory

import (
	"github.com/ilmsg/linebot-ispira/internal/db"
	"github.com/ilmsg/linebot-ispira/internal/models"
	"github.com/ilmsg/linebot-ispira/internal/services"
)

type channelInMemoryService struct {
	db db.Channel
}

// Create implements services.ChannelService.
func (c *channelInMemoryService) Create(channel *models.Channel) error {
	return c.db.Create(channel)
}

// Get implements services.ChannelService.
func (c *channelInMemoryService) Get(lineId string) (*models.Channel, error) {
	return c.db.Get(lineId)
}

// List implements services.ChannelService.
func (c *channelInMemoryService) List() ([]*models.Channel, error) {
	return c.db.List()
}

func NewChannelInMemoryService(db db.Channel) services.ChannelService {
	return &channelInMemoryService{db}
}

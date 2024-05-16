package inmemory

import (
	"errors"

	"github.com/ilmsg/linebot-ispira/internal/db"
	"github.com/ilmsg/linebot-ispira/internal/models"
)

type inMemoryDB struct {
	channels map[string]*models.Channel
}

// Create implements db.Channel.
func (i *inMemoryDB) Create(channel *models.Channel) error {
	i.channels[channel.ID] = channel
	return nil
}

// Get implements db.Channel.
func (i *inMemoryDB) Get(lineId string) (*models.Channel, error) {
	channel, ok := i.channels[lineId]
	if !ok {
		return nil, errors.New("channel not found")
	}
	return channel, nil
}

// List implements db.Channel.
func (i *inMemoryDB) List() ([]*models.Channel, error) {
	var channels []*models.Channel = []*models.Channel{}
	for _, channel := range i.channels {
		channels = append(channels, channel)
	}
	return channels, nil
}

func NewInMemoryDB() db.Channel {
	return &inMemoryDB{channels: make(map[string]*models.Channel)}
}

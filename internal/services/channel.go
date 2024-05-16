package services

import "github.com/ilmsg/linebot-ispira/internal/models"

type ChannelService interface {
	List() ([]*models.Channel, error)
	Get(lineId string) (*models.Channel, error)
	Create(*models.Channel) error
}

type LinebotService interface {
}

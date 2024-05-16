package db

import "github.com/ilmsg/linebot-ispira/internal/models"

type Channel interface {
	List() ([]*models.Channel, error)
	Get(lineId string) (*models.Channel, error)
	Create(*models.Channel) error
}

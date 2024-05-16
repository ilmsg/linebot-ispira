package models

type Channel struct {
	ID     string `json:"linebot_id"`
	Secret string `json:"linebot_secret"`
	Token  string `json:"linebot_token"`
}

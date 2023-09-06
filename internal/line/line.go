package line

import (
	"log"

	"github.com/a23667788/m800-assignment/internal/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewLineBot(channel config.Channel) (*linebot.Client, error) {
	bot, err := linebot.New(channel.ChannelSecret, channel.ChannelAccessToken)
	if err != nil {
		log.Println("linebot.New error", err)
		return nil, err
	}

	return bot, nil
}

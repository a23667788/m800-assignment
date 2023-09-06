package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/a23667788/m800-assignment/internal/config"
	"github.com/a23667788/m800-assignment/internal/database"
	"github.com/a23667788/m800-assignment/internal/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type m800Service struct {
	Port int

	Client *linebot.Client
	Mongo  *database.Mongo
}

func NewM800Service(setting *config.Setting, port int) *m800Service {
	if setting == nil {
		log.Fatalln("setting is nil")
	}

	// connect to line
	bot, err := line.NewLineBot(setting.Channel)
	if err != nil {
		log.Fatalln(err)
	}

	// connect to mongo db
	db, err := database.NewMongoDb(setting.MongoDb)
	if err != nil {
		log.Fatalln(err)
	}

	return &m800Service{Port: port, Client: bot, Mongo: db}
}

func (s *m800Service) StartWebServer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hi")
	})

	router.POST("/callback", func(c *gin.Context) {
		events, err := s.Client.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Println(err)
			}
			return
		}

		for _, event := range events {
			// Save all messages in different collections based on event type
			err = s.Mongo.Insert(string(event.Type), event)
			if err != nil {
				log.Fatalln("Insert Document error", err)
			}

			// Reply same message to users.
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// Handle text message
					handleTextMessage(s.Client, event.ReplyToken, message.Text)
				}
			}
		}

	})

	portStr := strconv.Itoa(s.Port)
	router.Run(":" + portStr)

}

func handleTextMessage(bot *linebot.Client, replyToken string, messageText string) {
	// Echo the received message back to the user
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(messageText)).Do(); err != nil {
		log.Println(err)
	}
}
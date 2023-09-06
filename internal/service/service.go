package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a23667788/m800-assignment/internal/config"
	"github.com/a23667788/m800-assignment/internal/database"
	"github.com/a23667788/m800-assignment/internal/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
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

	router.POST("/message", func(c *gin.Context) {
		var lineMessage line.LineMessage

		if err := c.ShouldBindJSON(&lineMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if lineMessage.UserID == "" {
			c.JSON(http.StatusNotFound, "User Id is empty")
		} else if lineMessage.Message == "" {
			c.JSON(http.StatusBadRequest, "Message is empty")
		} else {
			lineTextMessage := linebot.NewTextMessage(lineMessage.Message)
			if _, err := s.Client.PushMessage(lineMessage.UserID, lineTextMessage).Do(); err != nil {
				c.JSON(http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusOK, "Message sent")
			}
		}
	})

	router.GET("/message", func(c *gin.Context) {
		userid := c.Query("userid")

		res, err := queryLineEvent(s.Mongo, "message", "source.userid", userid)
		if err != nil {
			log.Fatalln(err)
		}

		if len(*res) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			c.JSON(http.StatusOK, res)
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

func queryLineEvent(db *database.Mongo, collectionName string, field string, filterString string) (*[]database.RawEvent, error) {
	if db == nil {
		return nil, fmt.Errorf("nil db")
	}
	if collectionName == "" {
		return nil, fmt.Errorf("empty collectionName")
	}

	var res []database.RawEvent
	var err error
	if filterString != "" {
		filter := bson.M{field: bson.M{"$eq": filterString}}
		res, err = db.QueryEvent("message", filter)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = db.QueryAllEvent("message")
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

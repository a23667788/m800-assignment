package database

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

type RawEvent struct {
	ReplyToken        string                     `json:"replyToken,omitempty"`
	Type              linebot.EventType          `json:"type"`
	Mode              linebot.EventMode          `json:"mode"`
	Timestamp         time.Time                  `json:"timestamp"`
	Source            *linebot.EventSource       `json:"source"`
	Message           *rawEventMessage           `json:"message,omitempty"`
	Postback          *linebot.Postback          `json:"postback,omitempty"`
	Beacon            *rawBeaconEvent            `json:"beacon,omitempty"`
	AccountLink       *rawAccountLinkEvent       `json:"link,omitempty"`
	Joined            *rawMemberEvent            `json:"joined,omitempty"`
	Left              *rawMemberEvent            `json:"left,omitempty"`
	Things            *rawThingsEvent            `json:"things,omitempty"`
	Unsend            *linebot.Unsend            `json:"unsend,omitempty"`
	VideoPlayComplete *linebot.VideoPlayComplete `json:"videoPlayComplete,omitempty"`
}

type rawMemberEvent struct {
	Members []*linebot.EventSource `json:"members"`
}

type rawEventMessage struct {
	ID                  string                      `json:"id"`
	Type                linebot.MessageType         `json:"type"`
	Text                string                      `json:"text,omitempty"`
	Duration            int                         `json:"duration,omitempty"`
	Title               string                      `json:"title,omitempty"`
	Address             string                      `json:"address,omitempty"`
	FileName            string                      `json:"fileName,omitempty"`
	FileSize            int                         `json:"fileSize,omitempty"`
	Latitude            float64                     `json:"latitude,omitempty"`
	Longitude           float64                     `json:"longitude,omitempty"`
	PackageID           string                      `json:"packageId,omitempty"`
	StickerID           string                      `json:"stickerId,omitempty"`
	StickerResourceType linebot.StickerResourceType `json:"stickerResourceType,omitempty"`
	Keywords            []string                    `json:"keywords,omitempty"`
	Emojis              []*linebot.Emoji            `json:"emojis,omitempty"`
	Mention             *linebot.Mention            `json:"mention,omitempty"`
}

type rawBeaconEvent struct {
	Hwid string                  `json:"hwid"`
	Type linebot.BeaconEventType `json:"type"`
	DM   string                  `json:"dm,omitempty"`
}
type rawAccountLinkEvent struct {
	Result linebot.AccountLinkResult `json:"result"`
	Nonce  string                    `json:"nonce"`
}

type rawThingsResult struct {
	ScenarioID             string                   `json:"scenarioId"`
	Revision               int                      `json:"revision"`
	StartTime              int64                    `json:"startTime"`
	EndTime                int64                    `json:"endTime"`
	ResultCode             linebot.ThingsResultCode `json:"resultCode"`
	ActionResults          []*rawThingsActionResult `json:"actionResults"`
	BLENotificationPayload string                   `json:"bleNotificationPayload,omitempty"`
	ErrorReason            string                   `json:"errorReason,omitempty"`
}

type rawThingsActionResult struct {
	Type linebot.ThingsActionResultType `json:"type,omitempty"`
	Data string                         `json:"data,omitempty"`
}

type rawThingsEvent struct {
	DeviceID string           `json:"deviceId"`
	Type     string           `json:"type"`
	Result   *rawThingsResult `json:"result,omitempty"`
}

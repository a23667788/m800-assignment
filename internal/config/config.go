package config

type Setting struct {
	Line    Line
	Channel Channel
	MongoDb MongoDb
}

type Channel struct {
	ChannelId          string
	ChannelSecret      string
	ChannelAccessToken string
}

type Line struct {
	Name    string
	Id      string
	Address string
	UserId  string
}

type MongoDb struct {
	Name             string
	ConnectionString string
	CollectionName   string
}

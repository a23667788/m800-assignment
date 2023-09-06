package main

import (
	"github.com/a23667788/m800-assignment/internal/config"
	"github.com/spf13/viper"
)

func InitializeSetting(filePath string) *config.Setting {
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var setting config.Setting
	initializeLineSetting(&setting.Line)
	initializeChannelSetting(&setting.Channel)
	initializeMongoDbSetting(&setting.MongoDb)

	return &setting
}

func initializeLineSetting(line *config.Line) {
	line.Name = viper.GetString("name")
	line.Id = viper.GetString("id")
	line.Address = viper.GetString("adress")
	line.UserId = viper.GetString("user_id")
}

func initializeChannelSetting(channel *config.Channel) {
	channel.ChannelId = viper.GetString("channel.channel_id")
	channel.ChannelSecret = viper.GetString("channel.channel_secret")
	channel.ChannelAccessToken = viper.GetString("channel.channel_access_token")
}

func initializeMongoDbSetting(mongo *config.MongoDb) {
	mongo.Name = viper.GetString("mongodb.name")
	mongo.ConnectionString = viper.GetString("mongodb.connectionString")
}

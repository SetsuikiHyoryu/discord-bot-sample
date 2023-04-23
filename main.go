package main

import (
	"discord-weather-bot/bot"
	"log"
	"os"
)

func main() {
	// 读取环境变量
	botToken, succeed := os.LookupEnv("BOT_TOKEN")

	if !succeed {
		log.Fatal("请创建 Discord bot token，并赋值到名为 BOT_TOKEN 的环境变量中。")
	}

	openWeatherToken, succeed := os.LookupEnv("OPENWEATHER_TOKEN")

	if !succeed {
		log.Fatal("请创建 Open Weather token，并赋值到名为 OPENWEATHER_TOKEN 的环境变量中。")
	}

	// 启动 bot
	bot.BotToken = botToken
	bot.OpenWeatherToken = openWeatherToken
	bot.Run()
}

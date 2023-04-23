package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// 保存 bot API tokens
var (
	OpenWeatherToken string
	BotToken         string
)

func Run() {
	// 创建一个 discord session
	discord, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		log.Fatal(err)
	}

	// 注入一个处理信息的 event handler
	discord.AddHandler(messageHandler)

	// 打开 session
	discord.Open()
	// Run 函数执行完毕时会执行 defer 的函数
	defer discord.Close()

	// 一直运行直到手动终止
	fmt.Println("Bot is running...")
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel
}

// 捕捉并处理会话中的所有消息
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// 忽略 bot 自己的消息（如果消息作者和 session 的用户一致，则表示是 bot 的消息）
	if message.Author.ID == session.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "weather"):
		session.ChannelMessageSend(message.ChannelID, "我可以帮助到您, 请用 '!zip <zip code>'")
	case strings.Contains(message.Content, "bot"):
		session.ChannelMessageSend(message.ChannelID, "我在")
	case strings.Contains(message.Content, "!zip"):
		currentWeather := getCurrentWeather(message.Content)
		session.ChannelMessageSendComplex(message.ChannelID, currentWeather)
	}
}

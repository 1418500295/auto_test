package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var bot *tgbotapi.BotAPI

// 向telegram机器人发送消息
const (
	chatID   = "" //要傳送訊息給指定用戶
	youToken = ""
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(youToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false

	link := fmt.Sprintf(`<a href="%s">[google]</a>`, "https://www.google.com.tw/")
	sendMsg(link)
}

func sendMsg(msg string) {
	msg = "html"
	NewMsg := tgbotapi.NewDocumentUpload(chatID, msg)
	//NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息
	_, err := bot.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}

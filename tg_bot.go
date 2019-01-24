package main

import (
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

var (
	telegramBotToken string
)

func init() {
	flag.StringVar(&telegramBotToken,
		"telegrambottoken",
		"775526701:AAEAuRjoMT8BoSmymZMmmu6oKvrr6Q1Uou0",
		"Telegram Bot Token")
	flag.Parse()

	if telegramBotToken == "" {
		log.Print("-telegramBotToken is required")
		os.Exit(1)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		reply := "Не знаю что сказать вам " + update.Message.From.UserName
		if update.Message == nil {
			continue
		}

		log.Printf("[ChatId - %o] \n [%s] Text: %s \nChat.ID %d \nCommandArguments() %s \nPinnedMessage %v \nCommand() %q \n",
			update.Message.Chat.ID,
			update.Message.From.UserName,
			update.Message.Text,
			update.Message.Chat.ID,
			update.Message.CommandArguments(),
			update.Message.PinnedMessage,
			update.Message.Command())
		//)
		var text = update.Message.Command()

		//switch  strings.ToLower(text){
		//case "start":
		//	reply = "Привет " + update.Message.From.UserName + ". Я телеграмм-бот"
		//case "hello":
		//	reply = "world"
		//}

		log.Println("Command: ", text)

		if !strings.Contains(text, "/") {
			reply = text
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		//bot.Send(msg)
		bot.Send(msg)

	}

}

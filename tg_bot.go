package main

import (
	"flag"
	"fmt"
	"github.com/dafanasev/go-yandex-translate"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

var (
	telegramBotToken                            string
	tr                                          *translate.Translator
	sourceLanguage, targetLanguage, text, token string
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
		//reply := "Не знаю что сказать вам " + update.Message.From.UserName
		if update.Message == nil {
			continue
		}

		log.Printf("ChatId - %o \nFrom.UserName: %s \nText: %s \nChat.ID %d \nCommandArguments() %s \nPinnedMessage %v \nCommand() %q \n",
			update.Message.Chat.ID,
			update.Message.From.UserName,
			update.Message.Text,
			update.Message.Chat.ID,
			update.Message.CommandArguments(),
			update.Message.PinnedMessage,
			update.Message.Command())

		switch strings.ToLower(update.Message.Command()) {
		case "start":
			update.Message.Text = "Привет " + update.Message.From.UserName + ". Я телеграмм-бот"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if strings.Compare(update.Message.Command(), "") == 0 {
			//update.Message.Text = "Не знаю что сказать на " + update.Message.Text + " вам " + update.Message.From.UserName
			msg.ReplyToMessageID = update.Message.MessageID

			Initial()
			update.Message.Text = Translation(update.Message.Text)

		}

		msg.Text = update.Message.Text

		bot.Send(msg)
	}
}

func InstalledLanguage() {
	response, err := tr.GetLangs(sourceLanguage)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println("Lang's: 	", response.Langs)
		//fmt.Println("Dirs:	", response.Dirs)
		//fmt.Println("Message:	", response.Message)
		fmt.Println("Code:	", response.Code)
		fmt.Println()
	}
}

func SetLanguage(s, t string) {
	sourceLanguage, targetLanguage = s, t
}

func Initial() {
	sourceLanguage, targetLanguage, text, token = "ru", "en", "Привет мир!", "trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad"
	tr = translate.New(token)
	SetLanguage(sourceLanguage, targetLanguage)
	InstalledLanguage()
}

func Translation(t string) string {
	translation, err := tr.Translate(targetLanguage, t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", translation.Result())
	}
	return translation.Result()
}

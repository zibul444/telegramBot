package main

import (
	"flag"
	"github.com/dafanasev/go-yandex-translate"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

var (
	telegramBotToken                      string
	tr                                    *translate.Translator
	sourceLanguage, targetLanguage, token string
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
		if update.Message == nil {
			continue
		}

		log.Printf(
			"ChatId - %o \nFrom.UserName: %s \nText: %s \nChat.ID %d \nCommandArguments() %s \nPinnedMessage %v \nCommand() %q \n _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_",
			update.Message.Chat.ID,
			update.Message.From.UserName,
			update.Message.Text,
			update.Message.Chat.ID,
			update.Message.CommandArguments(),
			update.Message.PinnedMessage,
			update.Message.Command())

		switch strings.ToLower(update.Message.Command()) {
		case "start":
			update.Message.Text = "Привет " + update.Message.From.UserName +
				". Я телеграмм-бот\n /replace - для смены языка\n введите текс для перевода: "
		case "replace":
			replaceLanguag()
			update.Message.Text = "Язык изменен, теперь " +
				"\nsourceLanguage: " + sourceLanguage + "\ntargetLanguage: " + targetLanguage
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if strings.Compare(update.Message.Command(), "") == 0 {
			msg.ReplyToMessageID = update.Message.MessageID

			InitialTranslation(
				"ru",
				"en",
				"trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad")
			update.Message.Text = Translation(update.Message.Text)
		}

		msg.Text = update.Message.Text
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func InstalledLanguage() {
	response, err := tr.GetLangs(sourceLanguage)
	if err != nil {
		log.Println(err)
	} else {
		//log.Println("Lang's: 	", response.Langs)
		//log.Println("Dirs:	", response.Dirs)
		//log.Println("Message:	", response.Message)
		log.Println("Code:	", response.Code)
		log.Println()
	}
}

func SetLanguage(s, t string) {
	sourceLanguage, targetLanguage = s, t
}

func replaceLanguag() {
	sourceLanguage, targetLanguage = targetLanguage, sourceLanguage
}

func InitialTranslation(sl, tl, t string) {
	//sourceLanguage, targetLanguage, token = "ru", "en", "trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad"
	sourceLanguage, targetLanguage, token = sl, tl, t
	tr = translate.New(token)
	SetLanguage(sourceLanguage, targetLanguage)
	InstalledLanguage()
}

func Translation(t string) string {
	translation, err := tr.Translate(targetLanguage, t)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Result: ", translation.Result())
	}
	return translation.Result()
}

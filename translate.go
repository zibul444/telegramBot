package main

import (
	"flag"
	"fmt"
	"github.com/dafanasev/go-yandex-translate"
)

var tr *translate.Translator
var sourceLanguage, targetLanguage, text, token string

func init() {
	flag.StringVar(&token,
		"k",
		"775526701:AAEAuRjoMT8BoSmymZMmmu6oKvrr6Q1Uou0",
		"Идентификатор API")
	flag.StringVar(&sourceLanguage, "s", "ru", "Исходный язык текста")
	flag.StringVar(&targetLanguage, "t", "en", "Целевой язык перевода")
	flag.StringVar(&text, "l", "Привет мир!", "Текст для перевода")

	flag.Parse()
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
	tr = translate.New(token)
	SetLanguage(sourceLanguage, targetLanguage)
	InstalledLanguage()
}

func Translation(language, text string) {
	translation, err := tr.Translate(language, text)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", translation.Result())
	}
}

func main() {
	Initial()
	Translation(targetLanguage, text)
}

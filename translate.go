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
		"trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad",
		"Идентификатор API")
	flag.StringVar(&sourceLanguage, "s", "en", "Исходный язык текста")
	flag.StringVar(&targetLanguage, "t", "ru", "Целевой язык перевода")
	flag.StringVar(&text, "l", "Hello world", "Текст для перевода")

	flag.Parse()
}

func installedLanguage() {
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

func setLanguage(s, t string) {
	sourceLanguage, targetLanguage = s, t
}

func initial() {
	tr = translate.New(token)
	setLanguage(sourceLanguage, targetLanguage)
	installedLanguage()
}

func funcName(language, text string) {
	translation, err := tr.Translate(language, text)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", translation.Result())
	}
}

func main() {
	initial()
	funcName(targetLanguage, text)
}

package main

//
//import (
//	"flag"
//	"fmt"
//	"github.com/dafanasev/go-yandex-translate"
//)
//
//var tr *translate.Translator
//var sourceLanguage, targetLanguage, text, token string
//
//func init() {
//	flag.StringVar(&token,
//		"k",
//		"trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad",
//		"Идентификатор API")
//	//flag.StringVar(&sourceLanguage, "s", "ru", "Исходный язык текста")
//	//flag.StringVar(&targetLanguage, "t", "en", "Целевой язык перевода")
//
//	flag.StringVar(&sourceLanguage, "s", "en", "Исходный язык текста")
//	flag.StringVar(&targetLanguage, "t", "ru", "Целевой язык перевода")
//
//	flag.StringVar(&text, "l", "undefined", "Текст для перевода")
//
//	flag.Parse()
//}
//
//func InstalledLanguage() {
//	response, err := tr.GetLangs(sourceLanguage)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		//fmt.Println("Lang's: 	", response.Langs)
//		//fmt.Println("Dirs:	", response.Dirs)
//		//fmt.Println("Message:	", response.Message)
//		fmt.Println("Code:	", response.Code)
//		fmt.Println()
//	}
//}
//
//func SetLanguage(s, t string) {
//	sourceLanguage, targetLanguage = s, t
//}
//
//func Initial() {
//	tr = translate.New(token)
//	SetLanguage(sourceLanguage, targetLanguage)
//	InstalledLanguage()
//}
//
//func Translation(text string) string {
//	translation, err := tr.Translate(targetLanguage, text)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Result: ", translation.Result())
//	}
//	return translation.Result()
//}
//
//
//func main() {
//	Initial()
//	Translation(text)
//}

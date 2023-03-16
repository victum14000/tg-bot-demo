package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var stop = false

var sas = tgbotapi.ReplyKeyboardRemove{
	RemoveKeyboard: true,
}

var menuButtons = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("👤 Search with name", "👤 Search with name")),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("💼 View History", "💼 View History")),
	// tgbotapi.NewInlineKeyboardButtonData("🛏️ Конец РД", "🛏️ Конец РД")),
)
var stopIt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⛔ Cancel")),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("💼 История", "💼 История")),
	// tgbotapi.NewInlineKeyboardButtonData("🛏️ Конец РД", "🛏️ Конец РД")),
)

func tgMain() {
	var bot, err = tgbotapi.NewBotAPI("your bot token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	go updateTg(updates, bot)
}

func tgInfo(name string) {
	var bot, err = tgbotapi.NewBotAPI("your bot token")
	if err != nil {
		log.Panic(err)
	}
	text := fmt.Sprintf("Employee %v has started working day!", name)
	msg := tgbotapi.NewMessage(your chat id, text)
	msg.ReplyMarkup = menuButtons
	bot.Send(msg)
	fmt.Println("\n ТГ отправлено")
}

func myLoop(name string) {
	stop = false
	var bot, err = tgbotapi.NewBotAPI("your bot token")
	if err != nil {
		log.Panic(err)
	}
	counter := 0
	for i := len(sqliteImport()) - 1; i >= 0; i-- {
		p := sqliteImport()[i]
		text := fmt.Sprintf(" Name: %v \n Working time: %v \n Fees: %v \n Was late for: %v \n Date: %v", p.userName, p.workTime, p.pens, p.lateTime, p.timeNow)
		if name != p.userName && name != "" {
			continue
		}
		msg := tgbotapi.NewMessage(your chat id, text)
		bot.Send(msg)
		time.Sleep(time.Millisecond * 200)
		counter++
		if counter >= 30 || stop {
			break
		}
	}
	if counter == 0 {
		text := "Ничего не найдено, попробуйте ввести имя и фамилию, например: \"Никита Пузанов\" "
		msg := tgbotapi.NewMessage(your chat id, text)
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	text := "Done!"
	msg := tgbotapi.NewMessage(your chat id, text)
	msg.ReplyMarkup = menuButtons
	bot.Send(msg)
}

func updateTg(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	var input = false
	for {
		for update := range updates {
			if update.CallbackQuery != nil {
				fmt.Println("\n Новая рекация")
				if update.CallbackQuery.Message.Chat.ID == your chat id {
					fmt.Println("\n айди верный")

					switch update.CallbackQuery.Data {
					case "💼 View History":
						fmt.Println("\n нажатие было История")
						text := "Loading... \n30 recent records: "
						msg := tgbotapi.NewMessage(your chat id, text)
						msg.ReplyMarkup = stopIt
						bot.Send(msg)
						time.Sleep(time.Second * 3)
						go myLoop("")
						fmt.Println("\n история отправлена")
					case "👤 Search with name":
						fmt.Println("\n нажатие было Поиск")
						text := "Enter the name to search for:"
						msg := tgbotapi.NewMessage(your chat id, text)
						bot.Send(msg)
						input = true
					}
				}
			}
			if update.Message != nil {
				fmt.Println("\n Новое сообщение")
				if update.Message.Chat.ID == your chat id {
					fmt.Println("\n айди верный")
					if input {
						input = false
						text := "loading... \n30 recent records: "
						msg := tgbotapi.NewMessage(your chat id, text)
						msg.ReplyMarkup = stopIt
						bot.Send(msg)
						time.Sleep(time.Second * 3)
						go myLoop(update.Message.Text)
						break
					}

					switch update.Message.Text {
					case "⛔ Cancel":
						stop = true
					default:
						fmt.Println("\n switch сработал")
						text := "Hello!"
						msg := tgbotapi.NewMessage(your chat id, text)
						msg.ReplyMarkup = sas
						bot.Send(msg)
						time.Sleep(time.Second * 1)
						text = "Please use the buttons below to get started"
						msg = tgbotapi.NewMessage(your chat id, text)
						msg.ReplyMarkup = menuButtons
						bot.Send(msg)
						fmt.Println("\n инфо отправлено")
					}
				}
			}
		}
	}
}

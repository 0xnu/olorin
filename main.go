package main

import (
	"fmt"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	emojiCounts := make(map[string]int)
	userChoices := make(map[int]string)
	var mu sync.Mutex
	var lastKeyboardMsgID int

	for update := range updates {
		if update.CallbackQuery != nil {
			userID := update.CallbackQuery.From.ID
			newEmoji := update.CallbackQuery.Data
			chatID := update.CallbackQuery.Message.Chat.ID

			mu.Lock()

			if lastEmoji, exists := userChoices[userID]; exists {
				emojiCounts[lastEmoji]--
			}

			userChoices[userID] = newEmoji
			emojiCounts[newEmoji]++

			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👍 %d", emojiCounts["thumbs_up"]), "thumbs_up"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👎 %d", emojiCounts["thumbs_down"]), "thumbs_down"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🔥 %d", emojiCounts["fire"]), "fire"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🥰 %d", emojiCounts["love"]), "love"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🚀 %d", emojiCounts["launch"]), "launch"),
				),
			)

			editMsg := tgbotapi.EditMessageTextConfig{
				BaseEdit: tgbotapi.BaseEdit{
					ChatID:      chatID,
					MessageID:   lastKeyboardMsgID,
					ReplyMarkup: &inlineKeyboard,
				},
				Text:                  "Follow Us: <a href=\"https://twitter.com/LatestJamz\">@LatestJamz</a>",
				ParseMode:             "HTML",
				DisableWebPagePreview: true,
			}

			bot.Send(editMsg)

			mu.Unlock()

			callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.AnswerCallbackQuery(callbackConfig)
		}

		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/sharemusic" {
			chatID := update.Message.Chat.ID
			messageID := update.Message.MessageID

			delConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: messageID,
			}
			bot.DeleteMessage(delConfig)

			audioConfig := tgbotapi.NewAudioUpload(chatID, "./audio/my_g_kizz_daniel.mp3")
			audioConfig.Title = "My G"
			audioConfig.Performer = "Kizz Daniel"
			audioConfig.MimeType = "audio/mp3"
			audioConfig.Caption = "Genre: Afrobeats\nReleased: Fri, 28th July 2023"

			_, err := bot.Send(audioConfig)
			if err != nil {
				log.Panic(err)
			}

			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👍 %d", emojiCounts["thumbs_up"]), "thumbs_up"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👎 %d", emojiCounts["thumbs_down"]), "thumbs_down"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🔥 %d", emojiCounts["fire"]), "fire"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🥰 %d", emojiCounts["love"]), "love"),
					tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🚀 %d", emojiCounts["launch"]), "launch"),
				),
			)

			keyboardMsg := tgbotapi.NewMessage(chatID, "Follow Us: <a href=\"https://twitter.com/LatestJamz\">@LatestJamz</a>")
			keyboardMsg.ParseMode = "HTML"
			keyboardMsg.ReplyMarkup = inlineKeyboard
			keyboardMsg.DisableWebPagePreview = true

			newMessage, err := bot.Send(keyboardMsg)
			if err != nil {
				log.Panic(err)
			}

			lastKeyboardMsgID = newMessage.MessageID
		}
	}
}
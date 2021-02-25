package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL = errors.New("url is invalid")
	errUnautorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Произошла неизвестная ошибка.")
	switch err {
	case errInvalidURL:
		msg.Text = "Это невалидная ссылка!"
		b.bot.Send(msg)
	case errUnautorized:
		msg.Text = "Ты не авторизирован! Используй команду /start ."
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Не удалось сохранить ссылку. Попробуй еще раз позже."
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}

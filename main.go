
package main

import (
 "fmt"
 "log"
 

 tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)



var token = "7717459803:AAEsKjO39AIeoVLtxZs9Fv7x8yrKlvpOn5A"


var registeredUsers = map[int64]bool{}

func main() {
 bot, err := tgBotAPI.NewBotAPI(token)
 if err != nil {
  log.Fatal(err)
 }

 log.Printf("Authorized on account %s", bot.Self.UserName)

 u := tgBotAPI.NewUpdate(0)
 u.Timeout = 60

 updates, err := bot.GetUpdatesChan(u)

 for update := range updates {
  if update.Message == nil && update.CallbackQuery == nil {
   continue
  }

  if update.Message != nil {
   switch update.Message.Text {
   case "/start":
    handleStart(update.Message.Chat.ID, bot)
   }
  }

  if update.CallbackQuery != nil {
   handleCallback(update.CallbackQuery, bot)
  }
 }
}

func handleStart(chatID int64, bot *tgBotAPI.BotAPI) {
 if !registeredUsers[chatID] {
  msg := tgBotAPI.NewMessage(chatID, "Вы не зарегистрированы. Пожалуйста, зарегистрируйтесь или войдите в аккаунт.")

  registrationKeyboard := tgBotAPI.NewInlineKeyboardMarkup(
   tgBotAPI.NewInlineKeyboardRow(
    tgBotAPI.NewInlineKeyboardButtonData("Зарегистрироваться", "register"),
    tgBotAPI.NewInlineKeyboardButtonData("Войти", "login"),
   ),
  )

  msg.ReplyMarkup = registrationKeyboard

  if _, err := bot.Send(msg); err != nil {
   fmt.Println(err)
  }
 } else {
  msg := tgBotAPI.NewMessage(chatID, "Добро пожаловать обратно! Выберите способ регистрации:")
  registrationKeyboard := tgBotAPI.NewInlineKeyboardMarkup(
   tgBotAPI.NewInlineKeyboardRow(
    tgBotAPI.NewInlineKeyboardButtonData("Регистрация через GitHub", "github_reg"),
    tgBotAPI.NewInlineKeyboardButtonData("Регистрация через Яндекс", "yandex_reg"),
   ),
  )

  msg.ReplyMarkup = registrationKeyboard

  if _, err := bot.Send(msg); err != nil {
   fmt.Println(err)
  }
 }
}


func handleCallback(callbackQuery *tgBotAPI.CallbackQuery, bot *tgBotAPI.BotAPI) { 
    chatID := callbackQuery.Message.Chat.ID 
   
    switch callbackQuery.Data { 
    case "register": 
     registeredUsers[chatID] = true // Регистрация пользователя 
     msg := tgBotAPI.NewMessage(chatID, "Зарегестрируйтесь через:") 
     msg.ReplyMarkup = tgBotAPI.NewInlineKeyboardMarkup( 
      tgBotAPI.NewInlineKeyboardRow( 
       tgBotAPI.NewInlineKeyboardButtonData("Регистрация через GitHub", "github_reg"), 
       tgBotAPI.NewInlineKeyboardButtonData("Регистрация через Яндекс", "yandex_reg"), 
      ), 
     ) 
     msg.ReplyToMessageID = callbackQuery.Message.MessageID 
   
    
     if _, err := bot.Send(msg); err != nil { 
      fmt.Println(err) 
     }
   
     
     if _, err := bot.AnswerCallbackQuery(tgBotAPI.NewCallback(callbackQuery.ID, "")); err != nil { 
      fmt.Println(err) 
     }
   
    case "login": 
     msg := tgBotAPI.NewMessage(chatID, "Введите ваши данные для входа:") 
     msg.ReplyToMessageID = callbackQuery.Message.MessageID
     if _, err := bot.Send(msg); err != nil {
      fmt.Println(err) 
     }
   
    
     if _, err := bot.AnswerCallbackQuery(tgBotAPI.NewCallback(callbackQuery.ID, "")); err != nil { 
      fmt.Println(err) 
     }
    }
   }
   
   


package main

import (
	"fmt"
	"time"
)

func main() {
	go tgMain()

	var user Worker = &User{}

	seleniumBot := NewSeleniumBot(user)
	for seleniumBot.driver == nil {
		time.Sleep(time.Second * 5)
	}
	for {
		seleniumBot.wbPoint()
		for user.name() == "" {
			time.Sleep(time.Second * 5)
		}
		user.wTime()
		fmt.Println("\nmain user", user)
		sqlite(user)
		tgInfo(user.name())
		user.addName("")
		fmt.Println("\nОбнулили имя")
		fmt.Println("\nЖдем 15 часов main")
		time.Sleep(time.Hour * 15)
		fmt.Println("Подождали 15 часов main")
	}

}

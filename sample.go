package main

import (
	"fmt"

	"github.com/tukejonny/g0tiu5a-bot/ctftimes"
)

func main() {
	var events []ctftimes.Event
	events = ctftimes.GetAPIData()
	for _, event := range events {
		fmt.Println("**********")
		fmt.Println(event.Title)
		fmt.Println(event.Format)
		fmt.Println(event.Weight)
		fmt.Println(event.Start)
		fmt.Println(event.Finish)
		fmt.Println(event.Url)
	}
}

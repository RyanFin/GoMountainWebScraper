package main

import (
	"RyanFin/GoMountainWebScraper/src/service"
	"fmt"
)

func main() {
	webScraper := service.WebScraper{}

	err := webScraper.BuildMountain()
	if err != nil {
		fmt.Println(err)
	}
}

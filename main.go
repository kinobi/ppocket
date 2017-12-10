package main

import (
	"log"
	"os"

	"github.com/kinobi/ppocket/pocket"
)

const urlCallback = "https://ppocket.herokuapp.com/"

func main() {
	ppocketUserAccessToken := os.Getenv("PPOCKET_USER_ACCESS_TOKEN")
	ppocketUsername := os.Getenv("PPOCKET_USERNAME")
	ppocketConsumerKey := os.Getenv("PPOCKET_API_CONSUMER_KEY")
	if ppocketConsumerKey == "" {
		log.Fatalln("Consumer key is missing: $PPOCKET_API_CONSUMER_KEY is empty")
	}

	if ppocketUserAccessToken == "" {
		ppocketUserAccessToken, ppocketUsername = pocket.OAuthProcess(ppocketConsumerKey, urlCallback)
	}

	log.Printf("Welcome to PPocket %s", ppocketUsername)
}

package main

import (
	"fmt"
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

	fmt.Println("Welcome to PPocket", ppocketUsername)

	res, err := pocket.Retrieve(ppocketConsumerKey, ppocketUserAccessToken)
	if err != nil {
		log.Fatalf("Failed to retrieve Pocket list: %s", err)
	}

	for _, item := range res.List {
		fmt.Printf("* %s [%s]\n", item.Title, item.URL)
	}
}

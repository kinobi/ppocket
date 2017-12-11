package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kinobi/ppocket/pocket"
)

const urlCallback = "https://ppocket.herokuapp.com/"

func main() {
	ppocketConsumerKey := flag.String("k", os.Getenv("PPOCKET_API_CONSUMER_KEY"), "Consumer key")
	if *ppocketConsumerKey == "" {
		log.Fatalln("Consumer key is missing")
	}
	ppocketUserAccessToken := flag.String("a", os.Getenv("PPOCKET_USER_ACCESS_TOKEN"), "User access token")
	flag.Parse()

	ppocketUsername := os.Getenv("PPOCKET_USERNAME")

	if *ppocketUserAccessToken == "" {
		*ppocketUserAccessToken, ppocketUsername = pocket.OAuthProcess(*ppocketConsumerKey, urlCallback)
	}

	fmt.Println("Welcome to PPocket", ppocketUsername)

	res, err := pocket.Retrieve(*ppocketConsumerKey, *ppocketUserAccessToken)
	if err != nil {
		log.Fatalf("Failed to retrieve Pocket list: %s", err)
	}

	for _, item := range res.List {
		fmt.Printf("* %s => %s \n[%s words | status: %v]\n", item.GivenTitle, item.GivenURL, item.WordCount, item.Status)
	}
}

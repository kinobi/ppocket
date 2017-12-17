package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

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

	since := time.Date(2017, 12, 16, 0, 0, 0, 0, time.UTC)

	query := pocket.NewGetQuery(
		pocket.WithState(pocket.QueryStateAll),
		pocket.WithFavorite(pocket.QueryFavoriteOrNot),
		pocket.WithTag("php"),
		pocket.WithSort(pocket.QuerySortNewest),
		pocket.WithSince(&since),
		pocket.WithPagination(4, 4),
	)

	res, err := pocket.Get(*ppocketConsumerKey, *ppocketUserAccessToken, query)
	if err != nil {
		log.Fatalf("Failed to retrieve Pocket list: %s", err)
	}
	for _, item := range res.List {
		fmt.Printf("* %s => %s \n", item.ResolvedTitle, item.GivenURL)
		for tag := range item.Tags {
			fmt.Println("\t- " + tag)
		}
		fmt.Printf("\t[%s words | status: %#v | favorite: %v]\n\n", item.WordCount, item.Status, item.Favorite)
	}
}

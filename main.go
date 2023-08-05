package main

import (
	"fmt"
	"log"
	"regexp"
	"context"

	"github.com/mattn/go-mastodon"
)

const regex = `<.*?>`

func main() {

	Init()

	c := mastodon.NewClient(&mastodon.Config{
		Server:       Server,
		ClientID:     C_key,
		ClientSecret: C_secret,
	})
	err := c.Authenticate(context.Background(), User, Pass)
	if err != nil {
		log.Fatal(err)
	}

	/* Get Timeline
	timeline, err := c.GetTimelineHome(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}
	*/


	notes, err := c.GetNotifications(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range notes {

		fmt.Printf("Account: %s\nStatus: %s\n", msg.Account.URL, stripHtmlRegex(msg.Status.Content))

		if stripHtmlRegex(msg.Status.Content) == BotName + " stop" && msg.Account.URL == Admin {
			fmt.Println("Stopping")
			//os.Exit(0)
		}

	}

	/* Clear Notifications
	err = c.ClearNotifications(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	*/

ReadFile()

}
/*
	// Post Status
	_, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status:      msg,
		//Visibility:	"direct",
	})
	if err != nil {
		log.Fatal(err)
	}
*/

func stripHtmlRegex(s string) string {
    r := regexp.MustCompile(regex)
    return r.ReplaceAllString(s, "")
}

func Toot(msg string) {

	c := mastodon.NewClient(&mastodon.Config{
		Server:       Server,
		ClientID:     C_key,
		ClientSecret: C_secret,
	})
	err := c.Authenticate(context.Background(), User, Pass)
	if err != nil {
		log.Fatal(err)
	}

    // Post Status
    _, errr := c.PostStatus(context.Background(), &mastodon.Toot{
        Status:      msg,
        //Visibility:   "direct",
    })
    if errr != nil {
        log.Fatal(err)
    }

}

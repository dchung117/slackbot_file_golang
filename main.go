package main

import (
	"fmt"
	"os"

	"github.com/dchung117/slackbot_file_golang/secret"

	"github.com/slack-go/slack"
)

func main() {

	// set bot token and channel id environment variable
	os.Setenv("SLACK_BOT_TOKEN", secret.SLACK_BOT_TOKEN)
	os.Setenv("CHANNEL_ID", secret.CHANNEL_ID)

	// create api client connection to slack
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// set channel array
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	// set file array (place path to file here)
	fileArr := []string{"/Users/davidchung/Downloads/wine-reviews.csv", "/Users/davidchung/Downloads/quicksort.py"}

	// loop through files and upload
	for i := 0; i < len(fileArr); i++ {
		// set upload parameters
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		// upload file to channel
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// print out file details
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}

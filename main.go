package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SultanKs4/gi-daily/setting"
	"github.com/SultanKs4/gi-daily/util"
)

func init() {
	setting.Setup()
}

func telegram(text string, photo string) {
	resp, err := util.SendMessageTelegram(text, photo)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("send telegram message: %v", resp["ok"])
	}
}

func main() {
	jar, err := util.JarCookies("https://hk4e-api-os.mihoyo.com/event/sol/")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Jar: jar,
	}

	awards, err := util.GetReward(client)
	if err != nil {
		log.Fatal(err)
	}

	status, day, err := util.GetStatus(client)
	if err != nil {
		log.Fatal(err)
	}

	claim, err := util.ClaimReward(client)
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		if claim["message"] == "OK" {
			award := awards[int(day)].(map[string]interface{})
			text := fmt.Sprintf("Day %v: %v x%v claimed successfully", day+1, award["name"], award["cnt"])
			fmt.Println(text)
			telegram(text, award["icon"].(string))
		} else {
			text := fmt.Sprintf("claim reward error: %v", claim["message"])
			fmt.Println(text)
			telegram(text, "")
		}
	} else {
		text := fmt.Sprintf("Day %v: Already claimed!", day)
		fmt.Println(text)
		telegram(text, "")
	}
}

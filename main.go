package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
  "log"
	"os"
	"io/ioutil"
)

func parse(url string, chFailedUrls chan string, chIsFinished chan bool) {
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  body := response.Body
  defer body.Close()

  bytes, err := ioutil.ReadAll(body)
  data := string(bytes)

	scanner := bufio.NewScanner(strings.NewReader(data))
	itemDef := new(ItemDef)
	for scanner.Scan() {

		switch line := scanner.Text(); {

		case strings.Contains(line, "|name ="):
			itemDef.Name = parseStringOf(line, "name")
			break

		case strings.Contains(line, "|members ="):
			itemDef.Members = parseBooleanOf(line, "members")
			break

		case strings.Contains(line, "|quest ="):
			itemDef.QuestItem = parseBooleanOf(line, "quest")
			break

		case strings.Contains(line, "|tradeable ="):
			itemDef.Tradeable = parseBooleanOf(line, "tradeable")
			break

		case strings.Contains(line, "|equipable ="):
			itemDef.Equipable = parseBooleanOf(line, "equipable")
			break;

		case strings.Contains(line, "|stackable ="):
			itemDef.Stackable = parseBooleanOf(line, "stackable")
			break;

		case strings.Contains(line, "|noteable ="):
			itemDef.Noteable = parseBooleanOf(line, "noteable")
			break;

		case strings.Contains(line, "|examine ="):
			itemDef.Examine = parseStringOf(line, "examine")
			break;

		}

	}

	file, err := os.OpenFile("item_defs.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s", itemDef.ToJson())

	chIsFinished <- true
}

func main() {
	start := time.Now()

	urls := [10]string{
		"https://oldschool.runescape.wiki/w/Abyssal_whip?action=raw",
		"https://oldschool.runescape.wiki/w/Adamant_longsword?action=raw",
		"https://oldschool.runescape.wiki/w/Abyssal_dagger?action=raw",
		"https://oldschool.runescape.wiki/w/Ancient_staff?action=raw",
		"https://oldschool.runescape.wiki/w/Abyssal_bludgeon?action=raw",
		"https://oldschool.runescape.wiki/w/3rd_age_wand?action=raw",
		"https://oldschool.runescape.wiki/w/3rd_age_plateskirt?action=raw",
		"https://oldschool.runescape.wiki/w/Air_rune?action=raw",
		"https://oldschool.runescape.wiki/w/Ahrim%27s_staff?action=raw",
		"https://oldschool.runescape.wiki/w/Adamant_scimitar?action=raw",
	}

	chFailedUrls := make(chan string)
	chIsFinished := make(chan bool)

	/* Concurrently parse each url */
	for _, url := range urls {
		go parse(url, chFailedUrls, chIsFinished)
	}

	failedUrls := make([]string, 0)
	for i := 0; i < len(urls); {
		select {
		case url := <-chFailedUrls:
			failedUrls = append(failedUrls, url)
		case <-chIsFinished:
			i++
		}
	}
	fmt.Println("Could not fetch these urls: ", failedUrls)
	fmt.Println("Duration: ", time.Since(start))

}

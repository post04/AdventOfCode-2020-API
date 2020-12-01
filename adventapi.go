package adventapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetBody() string {
	var link = "https://adventofcode.com/2020/stats"
	resp, err := http.Get(link)
	if err != nil {
		return "ERROR"
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	thing := string(bodyBytes)
	return thing
}

func GetNumToMultiply(body string) int {
	var numToMul int
	arr := strings.Split(body, "\n")
	for _, thing := range arr {
		if strings.HasPrefix(thing, "<p>Here are the current") {
			keks := strings.Split(thing, "*</span> star represents up to ")
			var kekw = keks[1]
			finalarr := strings.Split(kekw, " users.</p>")
			numToMul1, err := strconv.Atoi(finalarr[0])
			if err != nil {
				fmt.Println(err)
			}
			numToMul = numToMul1
		}
	}
	return numToMul
}

func GetCurrentDay(body string) int {
	arr := strings.Split(body, "\n")

	for _, thing := range arr {
		if strings.HasPrefix(thing, "<a href=\"/2020/day/") {
			thingSplit := strings.Split(thing, "\"> ")
			thing2Split := strings.Split(thingSplit[1], "  <span ")
			finalint, err := strconv.Atoi(thing2Split[0])
			if err != nil {
				return 999
			}
			return finalint
		}
	}

	return 999
}

func GetAllStats(body string) []string {
	arr := strings.Split(body, "\n")
	var finalkek []string
	for _, thing := range arr {
		if strings.HasPrefix(thing, "<a href=\"/2020/day/") {
			var finalstring = ""
			kek2 := strings.Split(thing, "</span>")
			for _, thing2 := range kek2 {
				if strings.HasPrefix(thing2, "<a href=\"/2020/day/") {
					lol := strings.Split(thing2, "\"> ")
					day := strings.Split(strings.Split(lol[1], "\">")[0], "  <span")[0]
					bothStats := strings.Split(lol[1], "\">")[1]
					finalstring += day + "-" + bothStats + "-"
				}
				if strings.HasPrefix(thing2, "  <span class=\"stats-firstonly\">") {
					finalstring += strings.Split(strings.Split(thing2, "<span class=\"stats-firstonly\">")[1], " ")[0]
				}

			}
			finalkek = append(finalkek, finalstring)
		}
	}
	return finalkek
}

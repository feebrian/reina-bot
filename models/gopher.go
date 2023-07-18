package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Gopher struct {
	Name string `json:"name"`
}

var KuteGoAPIURL = os.Getenv("KUTEGO_APIURL")

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!reina" {
		resp, err := http.Get(KuteGoAPIURL + "/gopher/" + "dr-who")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			_, err := s.ChannelFileSend(m.ChannelID, "dr-who.png", resp.Body)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Error: Can't get dr-who Gopher!")
		}
	}

	if m.Content == "!random" {
		resp, err := http.Get(KuteGoAPIURL + "/gopher/random")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			_, err := s.ChannelFileSend(m.ChannelID, "random-gopher.png", resp.Body)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Error: Can't get random Gopher!")
		}
	}

	if m.Content == "!gophers" {
		resp, err := http.Get(KuteGoAPIURL + "/gophers/")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var data []Gopher
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Fatal(err)
			}

			var gophers strings.Builder
			for _, gopher := range data {
				gophers.WriteString(gopher.Name + "\n")
			}

			_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Error: can't get list of Gophers!")
		}
	}
}

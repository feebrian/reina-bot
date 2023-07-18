package models

import "github.com/bwmarrin/discordgo"

type Gopher struct {
	Name string `json:"name"`
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}

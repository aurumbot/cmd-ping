package main

import (
	f "github.com/aurumbot/dat/foundation"
	dsg "github.com/bwmarrin/discordgo"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["hi"] = &f.Command{
		Name:    "Ping The Bot",
		Help:    "Pings the bot to see if its online.",
		Perms:   -1,
		Version: "v1.1",
		Action:  ping,
	}
}

func ping(session *dsg.Session, message *dsg.Message) {
	session.ChannelMessageSend(message.ChannelID, "Hello, World!")
}

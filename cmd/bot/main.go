package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	config := struct {
		conf.Version
		Bot struct {
			Token string
		}
	}{
		Version: conf.Version{
			SVN:  "dev",
			Desc: "",
		},
	}

	// parse configurations
	if err := conf.Parse(os.Args[1:], "MyBot", &config); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, _ := conf.Usage("MyBot", &config)

			log.Println(usage)
			return
		case conf.ErrVersionWanted:
			version, _ := conf.VersionString("MyBot", &config)

			log.Println(version)
			return
		}
	}

	discord, err := discordgo.New("Bot " + config.Bot.Token)
	fmt.Println(discord, err)

	discord.AddHandler(messageCreate)

	if err = discord.Open(); err != nil {
		fmt.Println(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}

package router

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type Router struct {
	log          *log.Logger
	commands     []*Command
	commandMaker string
}

type Command struct {
	Name        string
	Description string
	Action      func(s *discordgo.Session, m *discordgo.MessageCreate)
}

func New(log *log.Logger, commandMarker string) *Router {
	return &Router{log: log, commandMaker: commandMarker}
}

func (r *Router) parseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	textMessage := m.Content
	fields := strings.Fields(textMessage)

	if !r.isCommand(m.Content) {
		return
	}

	command := fields[0][1:]
	values := fields[1:]
	r.log.Println(command, values)

	for _, cmd := range r.commands {
		if cmd.Name == command {
			cmd.Action(s, m)
			return
		}
	}

	_, err := s.ChannelMessageSendReply(
		m.Message.ChannelID,
		fmt.Sprintf("O comando %v não é valido.", command),
		m.Reference())

	if err != nil {
		r.log.Println(err)
	}
}

func (r *Router) isCommand(message string) bool {
	fields := strings.Fields(message)
	if len(fields) < 1 {
		return false
	}

	firstField := fields[0]

	return strings.HasPrefix(firstField, r.commandMaker) && len(firstField) > len(r.commandMaker)
}

func (r *Router) OnCreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !r.isCommand(m.Content) {
		return
	}

	r.parseCommand(s, m)

}
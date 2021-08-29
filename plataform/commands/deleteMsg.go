package commands

import (
	"github.com/bwmarrin/discordgo"
	"mybot/plataform/router"
)

var DeleteMsgCmd = router.Command{
	Name:        "deleteMsg",
	Description: "Delete uma quantidade definida de mensagens onde é invocado.",
	Action:      deleteMsg,
}

func deleteMsg(s *discordgo.Session, m *discordgo.MessageCreate) error {
	return nil
}

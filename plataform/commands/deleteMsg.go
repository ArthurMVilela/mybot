package commands

import (
	"github.com/bwmarrin/discordgo"
	"mybot/plataform/router"
	"strconv"
	"strings"
)

var DeleteMsgCmd = router.Command{
	Name:        "deleteMsg",
	Description: "Delete uma quantidade definida de mensagens onde é invocado.",
	Action:      deleteMsg,
}

func deleteMsg(s *discordgo.Session, m *discordgo.MessageCreate) error {
	args := strings.Fields(m.Content)

	if len(args) < 2 {
		_, err := s.ChannelMessageSendReply(
			m.Message.ChannelID,
			"Argumentos inválido. Use %deleteMsg help para ajuda.",
			m.Reference())
		return err
	}

	if args[1] == "help" {
		helpEmbed := discordgo.MessageEmbed{
			Title:       "Delete Mensagens",
			Description: "deleta mensagens em um canal.\nvantepaco [quantidade] [autor]\n",
			Color:       0x009933,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "quantidade",
					Value:  "Quantidade de mensagens a serem deletada, mínimo 1 e máximo 100",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "autor",
					Value:  "Autor das mensagens a serem deletadas. Opcional",
					Inline: false,
				},
			},
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &helpEmbed)
		return err
	}

	amount, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil || (amount < 1 || amount > 100) {
		_, err := s.ChannelMessageSendReply(
			m.Message.ChannelID,
			"Argumentos inválido. Use %deleteMsg help para ajuda.",
			m.Reference())
		return err
	}

	messages, _ := s.ChannelMessages(m.ChannelID, int(amount), "", "", "")
	var messagesId []string
	for _, m := range messages {
		messagesId = append(messagesId, m.ID)
	}

	err = s.ChannelMessagesBulkDelete(m.ChannelID, messagesId)

	return err
}

package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"mybot/plataform/router"
	"strconv"
	"strings"
	"time"
)

var VampetacoCmd = router.Command{
	Name:        "vampetaco",
	Description: "Invoca um vampetaço no canal chamado",
	Action:      vampetaco,
}

var vampsPicsSFW = []string{
	"https://pbs.twimg.com/media/EUfH4EVXYAIeVGh.jpg",
	"https://ecosinternos.files.wordpress.com/2020/10/vampeta-540x338-1.jpg?w=540",
	"https://pbs.twimg.com/media/Edzf4ypWsAAtG8Y.jpg:large",
	"https://uploads.metropoles.com/wp-content/uploads/2021/03/13154752/Design-sem-nome-2021-03-13T154729.596-600x400.jpg",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSXr_AGxe25hOG1_mJukMfSwwLPIOXxlqZpLHu01Y9S6Zl2vgTXlHVQMJi4Yng6CTUguAo&usqp=CAU",
	"https://f.i.uol.com.br/fotografia/2013/10/31/331934-970x600-1.jpeg",
	"https://scontent.fcgh17-1.fna.fbcdn.net/v/t1.18169-9/22308819_1895786947404296_892622575886006285_n.jpg?_nc_cat=100&ccb=1-5&_nc_sid=730e14&_nc_ohc=uuooey9kLLoAX_mk0kP&_nc_ht=scontent.fcgh17-1.fna&oh=1be48d3788bbf1b754608492031e1985&oe=614FBA02",
	"https://frutilau.files.wordpress.com/2011/01/vampeta1.jpg?w=584",
	"https://pbs.twimg.com/media/EckbF9OXoAEh6zs.jpg",
}

var vampsPicsNSFW = []string{
	"https://alvarocantor.files.wordpress.com/2010/07/vampeta.jpg",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSMBFNhcvzpd4bV82ee0LqEFyPBxpu5eWzxQNBlV32obPfHmg9EdR6EhChV4vI_8zyzjr4&usqp=CAU",
	"https://pbs.twimg.com/media/CrNcxpZWIAE3QP8.jpg",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS6zS8WIoVlFS78Sfx_c0WQIhX8brGnn2dWuLP7llvR47jNjtTwvDLRK9yfa-DsRGChljA&usqp=CAU",
	"https://pbs.twimg.com/media/EV6Sm8AXgAE8WLo.jpg",
	"https://pbs.twimg.com/media/EctO1AVWoAcPpwt.jpg",
	"https://pbs.twimg.com/media/EWk33IhXkAInuEk.jpg",
}

func vampetaco(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Fields(m.Content)
	mode := strings.ToLower(args[1])
	if mode != "sfw" && mode != "nsfw" {
		mode = "sfw"
	}

	amount, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		return
	}

	var pics []string

	if mode == "sfw" {
		pics = vampsPicsSFW
	} else {
		pics = vampsPicsNSFW
	}

	for i := 0; i < int(amount); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(pics) - 1)
		picUrl := pics[index]

		embeded := discordgo.MessageEmbed{
			Title: "Vampeta",
			Color: 0x009933,
			Image: &discordgo.MessageEmbedImage{URL: picUrl},
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embeded)
		log.Println(err)

	}
}

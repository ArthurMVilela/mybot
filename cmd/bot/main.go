package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"github.com/bwmarrin/discordgo"
	"log"
	"mybot/plataform/commands"
	"mybot/plataform/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := log.New(os.Stdout, "[MyBot] ", log.Lshortfile|log.Lmicroseconds)

	if err := run(log); err != nil {
		log.Fatal(err)
	}
}

func run(log *log.Logger) error {
	log.Println("Iniciando Bot.")

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

	log.Println("Lendo configurações.")
	// parse configurations
	if err := conf.Parse(os.Args[1:], "MyBot", &config); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, _ := conf.Usage("MyBot", &config)

			_, err := fmt.Fprint(log.Writer(), usage)
			return err
		case conf.ErrVersionWanted:
			version, _ := conf.VersionString("MyBot", &config)

			_, err := fmt.Fprint(log.Writer(), version)
			return err
		}
		return err
	}

	log.Println("Autentificando bot.")
	discord, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		return err
	}

	r := router.New(log, "%")
	r.AddCommand(&commands.VampetacoCmd)
	r.AddCommand(&commands.DeleteMsgCmd)

	log.Println("Adicionando Handler.")

	discord.AddHandler(ready)
	discord.AddHandler(r.OnCreateMessage)

	log.Println("Abrindo websocket.")
	if err = discord.Open(); err != nil {
		return err
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Fechando websocket.")
	return discord.Close()
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("ready")
}

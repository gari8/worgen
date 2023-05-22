package main

import (
	"github.com/gari8/worgen/config"
	"github.com/gari8/worgen/gen"
	"log"
)

func main() {
	c := config.NewConfig()
	if err := c.Load(); err != nil {
		log.Fatal(err)
	}
	switch c.Mode {
	case config.ModeGen:
		err := c.Conversation()
		if err != nil {
			log.Fatal(err)
		}
		archive, err := gen.NewTemplate(c).ReadTemplates(c.ImportPath)
		if err != nil {
			log.Fatal(err)
		}
		ar := gen.NewArchive(archive)
		for _, f := range ar.Files {
			if err := f.CreateFile(); err != nil {
				log.Fatal(err)
			}
		}
		log.Printf(genOk, c.AppName)
	case config.ModeHelp:
		log.Println(defaultResp)
	default:
		log.Println(defaultResp)
	}
}

const (
	genOk = `
if you want to run service, please run command below:
cd %s
go mod tidy
docker-compose up`
	defaultResp = `
if you want to create new project, please run command below:
worgen new`
)

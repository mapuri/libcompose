package app

import (
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/contiv/deploy/hooks"
	"github.com/docker/libcompose/project"
)

func plugin(p *project.Project, context *cli.Context) error {
	logrus.Infof("VJ =====> Project %#v  Context %#v", p, context)
	cliLabels := ""
	event := project.NoEvent
	switch context.Command.Name {
	case "up":
		event = project.EventProjectUpStart
	case "start":
		event = project.EventProjectUpStart
	case "down":
		event = project.EventProjectDownStart
	case "delete":
		event = project.EventProjectDownStart
	case "kill":
		event = project.EventProjectDownStart
	case "create":
	case "build":
	case "ps":
	case "port":
	case "pull":
	case "log":
	case "restart":
	}

	if event == project.NoEvent {
		return nil
	}

	if event == project.EventProjectUpStart && cliLabels != "" {
		if err := hooks.PopulateEnvLabels(p, cliLabels); err != nil {
			log.Fatalf("Unable to insert environment labels. Error %v", err)
		}
	}

	if err := hooks.NetHooks(p, event); err != nil {
		log.Fatalf("Unable to generate network labels. Error %v", err)
	}

	return nil
}

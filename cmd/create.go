package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/cjimti/gdeck/part"
	"github.com/desertbit/grumble"
)

func init() {
	createCmd := &grumble.Command{
		Name: "create",
		Help: "create components such as projects, decks, suits and cards",
	}

	App.AddCommand(createCmd)

	createCmd.AddCommand(&grumble.Command{
		Name: "project",
		Help: "create a project",
		Run: func(c *grumble.Context) error {
			createProject()
			return nil
		},
	})

	createCmd.AddCommand(&grumble.Command{
		Name: "deck",
		Help: "create a deck",
		Run: func(c *grumble.Context) error {
			createProject()
			return nil
		},
	})

	createCmd.AddCommand(&grumble.Command{
		Name: "suit",
		Help: "create a suit (classification)",
		Run: func(c *grumble.Context) error {
			createProject()
			return nil
		},
	})

	createCmd.AddCommand(&grumble.Command{
		Name: "card",
		Help: "create a card",
		Run: func(c *grumble.Context) error {
			createProject()
			return nil
		},
	})

}

func createProject() {

	name := ""
	prompt := &survey.Input{
		Message: "Project Name:",
	}
	survey.AskOne(prompt, &name, nil)

	machineName := machineName(name)

	description := ""
	prompt = &survey.Input{
		Message: "Project Description:",
	}
	survey.AskOne(prompt, &description, nil)

	component := part.Component{
		Kind:        "Project",
		MachineName: machineName,
		Name:        name,
		Description: description,
	}

	filename := machineName + "-gdeck.yml"
	project := part.Project{
		Component: component,
	}

	saved := confirmAndSave(filename, project)
	if saved {
		fmt.Println()
		fmt.Printf("NOTICE: Project %s was saved as %s.\n", name, filename)
	}
}

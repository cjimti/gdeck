package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/desertbit/grumble"
	"github.com/olekukonko/tablewriter"
)

func init() {
	listCmd := &grumble.Command{
		Name: "list",
		Help: "list components such as projects, suits, decks and cards",
	}

	App.AddCommand(listCmd)

	listCmd.AddCommand(&grumble.Command{
		Name: "projects",
		Help: "list projects",
		Run: func(c *grumble.Context) error {
			listProjects()
			return nil
		},
	})

}

func listProjects() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Machine Name", "Project", "File Name", "Description"})

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filename := f.Name()

		// if file ends with -gdeck.yml it's a project file
		// load it and get the name
		match, _ := regexp.MatchString("-gdeck\\.yml$", filename)

		if match {
			project := loadProject(filename)
			table.Append([]string{project.Component.MachineName, project.Component.Name, filename, project.Component.Description})
		}
	}

	table.Render()

}

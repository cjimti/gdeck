package cmd

import (
	"fmt"
	"os"

	"github.com/desertbit/grumble"
	"github.com/olekukonko/tablewriter"
)

func init() {
	openCmd := &grumble.Command{
		Name:      "open",
		Help:      "a project: open [project]",
		Usage:     "use 'open [project]' to open an existing project",
		AllowArgs: true,
		Run: func(c *grumble.Context) error {
			if len(c.Args) != 1 {
				return fmt.Errorf("invalid number of arguments")
			}

			openProject(c.Args[0])
			return nil
		},
	}

	App.AddCommand(openCmd)

}

func openProject(machineName string) {
	filename := machineName + "-gdeck.yml"

	fmt.Printf("Opening project '%s' from file '%s'.\n", machineName, filename)

	project := loadProject(filename)
	table := tablewriter.NewWriter(os.Stdout)

	table.Append([]string{"NAME", project.Component.Name})
	table.Append([]string{"MACHINE NAME", project.Component.MachineName})
	table.Append([]string{"DESCRIPTION", project.Component.Description})
	table.Append([]string{"DECKS", string(len(project.Decks))})
	table.Render()

	global.Project = project
}

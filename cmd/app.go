package cmd

import (
	"fmt"

	"github.com/cjimti/gdeck/part"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var global struct {
	Project part.Project
	Env     map[string]string
}

var App = grumble.New(&grumble.Config{
	Name:                  "Deck",
	Description:           "Card deck creation system.",
	HistoryFile:           "/tmp/deck.hist",
	Prompt:                "g-deck » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,

	Flags: func(f *grumble.Flags) {
		f.String("d", "directory", "DEFAULT", "set an alternative root directory path")
		f.Bool("v", "verbose", false, "enable verbose mode")
	},
})

func init() {
	version := "v0.0.0"
	global.Env = map[string]string{}

	App.SetPrintASCIILogo(func(a *grumble.App) {
		fmt.Println(`   _______   _______   ______  __  ___ `)
		fmt.Println(`  |       \ |   ____| /      ||  |/  / `)
		fmt.Println(`  |  .--.  ||  |__   |  ,----'|  '  /  `)
		fmt.Println(`  |  |  |  ||   __|  |  |     |    <   `)
		fmt.Println("  |  '--'  ||  |____ |  `----.|  .  \\  ")
		fmt.Println(`  |_______/ |_______| \______||__|\__\ `)
		fmt.Println("                               ", version)
		fmt.Println()
		fmt.Println(`  - Deck Engineering Construction Kit - `)
		fmt.Println()
	})

}

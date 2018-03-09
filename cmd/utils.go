package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey"
	"github.com/cjimti/gdeck/part"
	"github.com/go-yaml/yaml"
)

func loadProject(filename string) part.Project {
	ymlData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	project := part.Project{}

	err = yaml.Unmarshal([]byte(ymlData), &project)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return project
}

func machineName(name string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	machineName := strings.ToLower(reg.ReplaceAllString(name, ""))

	prompt := &survey.Input{
		Message: "Machine Name:",
		Help: "\n The Machine Name is used for file names and referencing components." +
			"\n This should not container spaces or special characters other than - and _" +
			"\n The Default should be acceptable.",
		Default: machineName,
	}

	survey.AskOne(prompt, &machineName, nil)

	return machineName
}

func confirmAndSave(filename string, component interface{}) bool {

	save := false
	saveMessage := fmt.Sprintf("Save project file %s?", filename)
	savePrompt := &survey.Confirm{
		Message: saveMessage,
	}
	survey.AskOne(savePrompt, &save, nil)

	if save == false {
		fmt.Println()
		fmt.Printf("NOTICE: %s was not saved.\n", filename)
		return false
	}

	if exists := fileExists(filename); exists != false {
		overMessage := fmt.Sprintf("WARNING: Project File %s exists. Overwrite?", filename)
		overPrompt := &survey.Confirm{
			Message: overMessage,
		}

		survey.AskOne(overPrompt, &save, nil)
	}

	if save == false {
		fmt.Println()
		fmt.Printf("NOTICE: %s was not saved.\n", filename)
		return false
	}

	// Marshal to YML and Save
	d, err := yaml.Marshal(component)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = ioutil.WriteFile("./"+filename, d, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}

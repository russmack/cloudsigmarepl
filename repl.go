// Package cloudsigmarepl is a repl for the CloudSigma REST API.
package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"os"
)

var (
	session *Session
)

// Session struct holds properties that we don't want to have to specify repeatedly.
type Session struct {
	Location string
	Username string
	Password string
}

func main() {
	initSession()

	// Create the repl, add command state machines, and start the repl.
	repl := replizer.NewRepl()
	repl.Name = "CloudSigma IaaS"

	// Create a statemachine per command available in the repl.

	repl.AddCommand(&replizer.Command{
		Instruction: "config location",
		StartFn:     ListConfigLocation().Start,
		Help:        "Show the current location for the session.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "set config location",
		StartFn:     EditConfigLocation().Start,
		Help:        "Set the location for the session.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "cloud status",
		StartFn:     NewCloudStatus().Start,
		Help:        "Get the status of the cloud.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "locations",
		StartFn:     NewLocations().Start,
		Help:        "Request available locations.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "usage",
		StartFn:     NewUsage().Start,
		// TODO: usage not in docs.
		Help: "????",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "balance",
		StartFn:     NewBalance().Start,
		Help:        "Request account balance.",
	})

	//repl.Add("create server", NewCreateServer().Start)

	repl.AddCommand(&replizer.Command{
		Instruction: "notification contacts",
		StartFn:     NewListNotifyContacts().Start,
		Help:        "Request notification contacts.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "create notification contact",
		StartFn:     NewCreateNotifyContacts().Start,
		Help:        "Request notification contacts.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "notification preferences",
		StartFn:     NewListNotifyPrefs().Start,
		Help:        "Request notification preferences for a specified contact.",
	})

	repl.AddCommand(&replizer.Command{
		Instruction: "edit notification preferences",
		StartFn:     NewEditNotifyPrefs().Start,
		Help:        "Edit notification preferences for a specified contact.",
	})

	repl.Start()
}

// initSession sets the initial state of the session, based on the config file.
func initSession() {
	config := cloudsigma.NewConfig()
	_, err := config.Load()
	if err != nil {
		fmt.Println("Unable to load config.", err)
		os.Exit(1)
	}
	session = &Session{}
	session.Username = config.Login().Username
	session.Password = config.Login().Password
}

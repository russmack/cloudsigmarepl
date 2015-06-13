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
	repl.FormatResponse = replizer.PrettyJson

	// Create a statemachine per command available in the repl.
	addCommands(repl)
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

func addNewCommand(repl *replizer.Repl, instr string, startFn replizer.CommandStartFn, help string) {
	repl.AddCommand(&replizer.Command{
		Instruction: instr,
		StartFn:     startFn,
		Help:        help,
	})
}

func addCommands(repl *replizer.Repl) {
	// Config location
	addNewCommand(repl, "config location", ListConfigLocation().Start,
		"Show the current location for the session.")

	addNewCommand(repl, "set config location", EditConfigLocation().Start,
		"Set the location for the session.")

	// Cloud
	addNewCommand(repl, "cloud status", NewCloudStatus().Start,
		"Get the status of the cloud.")

	addNewCommand(repl, "locations", NewLocations().Start,
		"Request available locations.")

	// Billing
	// TODO: usage not in docs.
	addNewCommand(repl, "usage", NewUsage().Start,
		"Not in docs.")

	addNewCommand(repl, "balance", NewBalance().Start,
		"Request account balance.")
	addNewCommand(repl, "current usage", NewCurrentUsage().Start,
		"Request current usage.")
	addNewCommand(repl, "burst usage", NewBurstUsage().Start,
		"Request burst usage.")

	// Servers
	addNewCommand(repl, "list servers", NewListServers().Start,
		"List all servers.")
	addNewCommand(repl, "start server", NewStartServer().Start,
		"Start a server.")
	addNewCommand(repl, "stop server", NewStopServer().Start,
		"Stop a server.")
	addNewCommand(repl, "create server", NewCreateServer().Start,
		"Create a new server.")

	// Drives
	addNewCommand(repl, "list drives", NewListDrives().Start,
		"List all drives.")
	addNewCommand(repl, "create drive", NewCreateDrive().Start,
		"Create a new drive.")

	// Notification Contacts
	addNewCommand(repl, "notification contacts", NewListNotifyContacts().Start,
		"Request notification contacts.")

	addNewCommand(repl, "create notification contact", NewCreateNotifyContacts().Start,
		"Create a notification contact.")

	addNewCommand(repl, "edit notification contact", NewEditNotifyContacts().Start,
		"Edit a notification contact.")

	addNewCommand(repl, "delete notification contact", NewDeleteNotifyContacts().Start,
		"Delete a notification contact.")

	// Notification Preferences
	addNewCommand(repl, "notification preferences", NewListNotifyPrefs().Start,
		"Request notification preferences for a specified contact.")

	addNewCommand(repl, "edit notification preferences", NewEditNotifyPrefs().Start,
		"Edit notification preferences for a specified contact.")
}

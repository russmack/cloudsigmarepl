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
	repl.Add("config location", NewGetConfigLocation().Start)
	repl.Add("set config location", NewSetConfigLocation().Start)
	repl.Add("cloud status", NewCloudStatus().Start)
	repl.Add("locations", NewLocations().Start)
	//repl.Add("usage", NewUsage().Start)
	//repl.Add("balance", NewBalance().Start)
	//repl.Add("create server", NewCreateServer().Start)
	//repl.Add("notification contacts", NewGetNotifyContacts().Start)
	//repl.Add("notification preferences", NewGetNotifyPrefs().Start)
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

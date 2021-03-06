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

func sendRequest(channels *replizer.Channels, args *cloudsigma.Args) error {
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		channels.ResponseChan <- fmt.Sprintf("Error calling client. %s", err)
		return err
	}
	channels.ResponseChan <- string(resp)
	return nil
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
	addNewCommand(repl, "capabilities", NewCapabilities().Start,
		"Request location's capabilities.")

	// Billing
	// TODO: usage not in docs.
	addNewCommand(repl, "usage", NewUsage().Start,
		"Not in docs.")
	addNewCommand(repl, "profile", NewProfile().Start,
		"Request profile.")
	addNewCommand(repl, "balance", NewBalance().Start,
		"Request account balance.")
	addNewCommand(repl, "list subscriptions", NewListSubscriptions().Start,
		"List subscriptions.")
	addNewCommand(repl, "list transactions", NewListTransactions().Start,
		"List transactions.")
	addNewCommand(repl, "pricing", NewListPricing().Start,
		"Request pricing.")
	addNewCommand(repl, "discounts", NewListDiscounts().Start,
		"List discounts.")
	addNewCommand(repl, "current usage", NewCurrentUsage().Start,
		"Request current usage.")
	addNewCommand(repl, "burst usage", NewBurstUsage().Start,
		"Request burst usage.")
	addNewCommand(repl, "daily burst usage", NewDailyBurstUsage().Start,
		"Request daily burst usage.")
	addNewCommand(repl, "licenses", NewListLicenses().Start,
		"List licenses.")

	// Servers
	addNewCommand(repl, "list servers", NewListServers().Start,
		"List all servers.")
	addNewCommand(repl, "start server", NewStartServer().Start,
		"Start a server.")
	addNewCommand(repl, "stop server", NewStopServer().Start,
		"Stop a server.")
	addNewCommand(repl, "shutdown server", NewShutdownServer().Start,
		"Shutdown a server.")
	addNewCommand(repl, "create server", NewCreateServer().Start,
		"Create a new server.")
	addNewCommand(repl, "delete server", NewDeleteServer().Start,
		"Delete a server.")
	//addNewCommand(repl, "attach drive", NewAttachDrive().Start,
	//	"Attach a drive to a server.")

	// Drives
	addNewCommand(repl, "list drives", NewListDrives().Start,
		"List all drives.")
	addNewCommand(repl, "list drives detailed", NewListDrivesDetailed().Start,
		"List all drives detailed.")
	addNewCommand(repl, "get drive", NewGetDrive().Start,
		"Get a drive.")
	addNewCommand(repl, "create drive", NewCreateDrive().Start,
		"Create a new drive.")
	addNewCommand(repl, "delete drive", NewDeleteDrive().Start,
		"Delete a drive.")

	// Snapshots
	addNewCommand(repl, "list snapshots", NewListSnapshots().Start,
		"List all snapshots.")
	addNewCommand(repl, "list snapshots detailed", NewListSnapshotsDetailed().Start,
		"List all snapshots detailed.")
	addNewCommand(repl, "get snapshot", NewGetSnapshot().Start,
		"Get a snapshot.")
	addNewCommand(repl, "create snapshot", NewCreateSnapshot().Start,
		"Create a snapshot.")
	addNewCommand(repl, "delete snapshot", NewDeleteSnapshot().Start,
		"Delete a snapshot.")

	// Network
	addNewCommand(repl, "list vlans", NewListVlans().Start,
		"List VLANs.")
	addNewCommand(repl, "list vlans detailed", NewListVlansDetailed().Start,
		"List VLANs detailed.")
	addNewCommand(repl, "get vlan", NewGetVlan().Start,
		"Get a single VLAN.")
	addNewCommand(repl, "list ips", NewListIps().Start,
		"List IP addresses.")
	addNewCommand(repl, "list ips detailed", NewListIpsDetailed().Start,
		"List IPs detailed.")
	addNewCommand(repl, "get ip", NewGetIp().Start,
		"Get a single IP.")

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

	// Access Control Lists
	addNewCommand(repl, "list acls", NewListAcls().Start,
		"List access control lists.")
	addNewCommand(repl, "get acl", NewGetAcl().Start,
		"Get an ACL.")

	// Tags
	addNewCommand(repl, "list tags", NewListTags().Start,
		"List tags.")
	addNewCommand(repl, "get tag", NewGetTag().Start,
		"Get a tag.")

	// Logs
	addNewCommand(repl, "list logs", NewListLogs().Start,
		"List audit logs.")

	// Jobs
	addNewCommand(repl, "list jobs", NewListJobs().Start,
		"List jobs.")

	// Firewall Policies
	addNewCommand(repl, "list fwpolicies", NewListFwPolicies().Start,
		"List firewall policies.")

	// Keypairs
	addNewCommand(repl, "list keypairs", NewListKeypairs().Start,
		"List SSH keypairs.")

}

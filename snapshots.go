package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	//"strconv"
)

type CommandListSnapshots struct {
	channels *replizer.Channels
}

// type CommandStartServer struct {
// 	channels *replizer.Channels
// }
// type CommandStopServer struct {
// 	channels *replizer.Channels
// }
type CommandCreateSnapshot struct {
	channels *replizer.Channels
}

type SnapshotCargo struct {
	Uuid string
	Body cloudsigma.SnapshotRequest
}

func NewListSnapshots() *CommandListSnapshots {
	return &CommandListSnapshots{}
}

// func NewStartServer() *CommandStartServer {
// 	return &CommandStartServer{}
// }
// func NewStopServer() *CommandStopServer {
// 	return &CommandStopServer{}
// }
func NewCreateSnapshot() *CommandCreateSnapshot {
	return &CommandCreateSnapshot{}
}

// Start is the start state of the CommandListSnapshots state machine.
func (m *CommandListSnapshots) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listSnapshots
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (g *CommandListSnapshots) listSnapshots(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	args := o.NewList()
	g.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(g.channels, args)
	return nil
}

/*
// Start is the start state of the CommandStartServer state machine.
func (m *CommandStartServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.startServerUuid
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStartServer) startServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.startServerUuid(c)
	}
	return m.startServerSendRequest(c)
}

func (m *CommandStartServer) startServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return nil
	}

	args := o.NewStart(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		m.channels.ResponseChan <- fmt.Sprintf("Error calling client. %s", err)
		return nil
	}
	m.channels.ResponseChan <- string(resp)
	return nil
}

// Start is the start state of the CommandStopServer state machine.
func (m *CommandStopServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.stopServerUuid
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStopServer) stopServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.stopServerUuid(c)
	}
	return m.stopServerSendRequest(c)
}

func (m *CommandStopServer) stopServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return nil
	}

	args := o.NewStop(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
*/

/*
// Start is the start state of the CommandCreateSnapshot state machine.
func (m *CommandCreateSnapshot) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createSnapshotName
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateSnapshot) createSnapshotName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Name:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Body.Name = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.createSnapshotName(c)
	}
	return m.createSnapshotMedia(c)
}

func (m *CommandCreateSnapshot) createSnapshotMedia(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Media (disk | cdrom):"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Body.Media = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.createSnapshotMedia(c)
	}
	return m.createSnapshotSize(c)
}

func (m *CommandCreateSnapshot) createSnapshotSize(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Size:"
	s := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			m.channels.ResponseChan <- "Invalid size."
			return m.createSnapshotSize(c)
		} else {
			c.Body.Size = n
		}
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.createSnapshotSize(c)
	}
	return m.createSnapshotSendRequest(c)
}

func (m *CommandCreateSnapshot) createSnapshotSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return nil
	}
	newSnapshots := []cloudsigma.SnapshotRequest{
		cloudsigma.SnapshotRequest{c.Body.Media, c.Body.Name, c.Body.Size},
	}
	args := o.NewCreate(newSnapshots)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
*/

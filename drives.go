package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandListDrives struct {
	channels *replizer.Channels
}

// type CommandStartServer struct {
// 	channels *replizer.Channels
// }
// type CommandStopServer struct {
// 	channels *replizer.Channels
// }
type CommandCreateDrive struct {
	channels *replizer.Channels
}

type DriveCargo struct {
	Uuid string
	Body cloudsigma.DriveRequest
}

func NewListDrives() *CommandListDrives {
	return &CommandListDrives{}
}

// func NewStartServer() *CommandStartServer {
// 	return &CommandStartServer{}
// }
// func NewStopServer() *CommandStopServer {
// 	return &CommandStopServer{}
// }
func NewCreateDrive() *CommandCreateDrive {
	return &CommandCreateDrive{}
}

// Start is the start state of the CommandListDrives state machine.
func (m *CommandListDrives) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listDrives
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (g *CommandListDrives) listDrives(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDrives()
	args := o.NewList()
	g.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		g.channels.ResponseChan <- fmt.Sprintf("Error calling client. %s", err)
		return nil
	}
	g.channels.ResponseChan <- string(resp)
	return nil
}

/*
// Start is the start state of the CommandStartServer state machine.
func (m *CommandStartServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.startServerUuid
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStartServer) startServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.startServerUuid(c)
	}
	return m.startServerSendRequest(c)
}

func (m *CommandStartServer) startServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(DriveCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Drive."
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
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStopServer) stopServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.stopServerUuid(c)
	}
	return m.stopServerSendRequest(c)
}

func (m *CommandStopServer) stopServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(DriveCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Drive."
		return nil
	}

	args := o.NewStop(c.Uuid)
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
*/

// Start is the start state of the CommandCreateDrive state machine.
func (m *CommandCreateDrive) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createDriveName
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateDrive) createDriveName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Name:"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Body.Name = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.createDriveName(c)
	}
	return m.createDriveMedia(c)
}

func (m *CommandCreateDrive) createDriveMedia(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Media (disk | cdrom):"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Body.Media = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.createDriveMedia(c)
	}
	return m.createDriveSize(c)
}

func (m *CommandCreateDrive) createDriveSize(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Size:"
	s := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			m.channels.ResponseChan <- "Invalid size."
			return m.createDriveSize(c)
		} else {
			c.Body.Size = n
		}
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.createDriveSize(c)
	}
	return m.createDriveSendRequest(c)
}

func (m *CommandCreateDrive) createDriveSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDrives()
	c, ok := cargo.(DriveCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Drive."
		return nil
	}
	newDrives := []cloudsigma.DriveRequest{
		cloudsigma.DriveRequest{c.Body.Media, c.Body.Name, c.Body.Size},
	}
	args := o.NewCreate(newDrives)
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

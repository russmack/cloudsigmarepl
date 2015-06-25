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
type CommandListDrivesDetailed struct {
	channels *replizer.Channels
}
type CommandGetDrive struct {
	channels *replizer.Channels
}
type CommandCreateDrive struct {
	channels *replizer.Channels
}
type CommandDeleteDrive struct {
	channels *replizer.Channels
}

type DriveCargo struct {
	Uuid string
	Body cloudsigma.DriveRequest
}

func NewListDrives() *CommandListDrives {
	return &CommandListDrives{}
}
func NewListDrivesDetailed() *CommandListDrivesDetailed {
	return &CommandListDrivesDetailed{}
}
func NewGetDrive() *CommandGetDrive {
	return &CommandGetDrive{}
}
func NewDeleteDrive() *CommandDeleteDrive {
	return &CommandDeleteDrive{}
}
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
	_ = sendRequest(g.channels, args)
	return nil
}

func (m *CommandListDrivesDetailed) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listDrivesDetailed
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandListDrivesDetailed) listDrivesDetailed(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDrives()
	args := o.NewListDetailed()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetDrive) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getDriveUuid
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetDrive) getDriveUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive in getDriveUuid."
		return m.getDriveUuid(c)
	}
	return m.getDriveSendRequest(c)
}

func (m *CommandGetDrive) getDriveSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDrives()
	c, ok := cargo.(DriveCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Drive in getDriveSendRequest."
		return nil
	}
	args := o.NewGet(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

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
	_ = sendRequest(m.channels, args)
	return nil
}

// Start is the start state of the CommandDeleteDrive state machine.
func (m *CommandDeleteDrive) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.deleteDriveUuid
	cargo := DriveCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandDeleteDrive) deleteDriveUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(DriveCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Drive."
		return m.deleteDriveUuid(c)
	}
	return m.deleteDriveSendRequest(c)
}

func (m *CommandDeleteDrive) deleteDriveSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDrives()
	c, ok := cargo.(DriveCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Drive."
		return nil
	}

	args := o.NewDelete(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

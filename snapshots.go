package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListSnapshots struct {
	channels *replizer.Channels
}
type CommandListSnapshotsDetailed struct {
	channels *replizer.Channels
}
type CommandGetSnapshot struct {
	channels *replizer.Channels
}
type CommandDeleteSnapshot struct {
	channels *replizer.Channels
}
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
func NewListSnapshotsDetailed() *CommandListSnapshotsDetailed {
	return &CommandListSnapshotsDetailed{}
}
func NewGetSnapshot() *CommandGetSnapshot {
	return &CommandGetSnapshot{}
}
func NewDeleteSnapshot() *CommandDeleteSnapshot {
	return &CommandDeleteSnapshot{}
}

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

func (m *CommandListSnapshotsDetailed) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listSnapshotsDetailed
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandListSnapshotsDetailed) listSnapshotsDetailed(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	args := o.NewListDetailed()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetSnapshot) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getSnapshotUuid
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetSnapshot) getSnapshotUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot in getSnapshotUuid."
		return m.getSnapshotUuid(c)
	}
	return m.getSnapshotSendRequest(c)
}

func (m *CommandGetSnapshot) getSnapshotSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot in getSnapshotSendRequest."
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

// Start is the start state of the CommandCreateSnapshot state machine.
func (m *CommandCreateSnapshot) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createSnapshotDrive
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateSnapshot) createSnapshotDrive(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Drive uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Body.Drive = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.createSnapshotDrive(c)
	}
	return m.createSnapshotName(c)
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
	return m.createSnapshotSendRequest(c)
}

func (m *CommandCreateSnapshot) createSnapshotSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return nil
	}
	newSnapshots := cloudsigma.SnapshotRequest{c.Body.Drive, c.Body.Name}
	args := o.NewCreate(newSnapshots)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

// Start is the start state of the CommandDeleteSnapshot state machine.
func (m *CommandDeleteSnapshot) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.deleteSnapshotUuid
	cargo := SnapshotCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandDeleteSnapshot) deleteSnapshotUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(SnapshotCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Snapshot."
		return m.deleteSnapshotUuid(c)
	}
	return m.deleteSnapshotSendRequest(c)
}

func (m *CommandDeleteSnapshot) deleteSnapshotSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSnapshots()
	c, ok := cargo.(SnapshotCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Snapshot."
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

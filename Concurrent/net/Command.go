package net

import "strconv"

type ICommand interface {
	execute() string
}

type QueryCommand struct {
	command []string
}
type ReportCommand struct {
	command []string
}
type StopCommand struct {
	command []string
}
type ErrorCommand struct {
	command []string
}

func NewQueryCommand(command []string) *QueryCommand {
	c := new(QueryCommand)
	c.command = command
	return c
}
func NewReportCommand(command []string) *ReportCommand {
	c := new(ReportCommand)
	c.command = command
	return c
}
func NewStopCommand(command []string) *StopCommand {
	c := new(StopCommand)
	c.command = command
	return c
}
func NewErrorCommand(command []string) *StopCommand {
	c := new(StopCommand)
	c.command = command
	return c
}
func (q *QueryCommand) execute() string {
	dao := GetDAO()
	if len(q.command) == 3 {
		return dao.query2(q.command[1], q.command[2])
	} else if len(q.command) == 4 {
		v, _ := strconv.Atoi(q.command[3])
		return dao.query3(q.command[1], q.command[2], v)

	} else {
		return "ERROR;Bad Command"
	}
}

func (r *ReportCommand) execute() string {
	dao := GetDAO()
	return dao.report(r.command[1])
}

func (s *StopCommand) execute() string {
	return "Server stopped"
}

func (e *ErrorCommand) execute() string {
	return "Unknown command: " + e.command[0]
}

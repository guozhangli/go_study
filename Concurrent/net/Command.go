package net

import "strconv"

type ICommand interface {
	execute() string
}

type QueryCommand struct{}
type ReportCommand struct{}
type StopCommand struct{}
type ErrorCommand struct{}

var command []string

func NewQueryCommand(command []string) *QueryCommand {
	command = command
	return new(QueryCommand)
}
func NewReportCommand(command []string) *ReportCommand {
	command = command
	return new(ReportCommand)
}
func NewStopCommand(command []string) *StopCommand {
	command = command
	return new(StopCommand)
}
func NewErrorCommand(command []string) *ErrorCommand {
	command = command
	return new(ErrorCommand)
}
func (q *QueryCommand) execute() string {
	dao := GetDAO()
	if len(command) == 3 {
		return dao.query2(command[1], command[2])
	} else if len(command) == 4 {
		v, _ := strconv.Atoi(command[3])
		return dao.query3(command[1], command[2], v)

	} else {
		return "ERROR;Bad Command"
	}
}

func (r *ReportCommand) execute() string {
	dao := GetDAO()
	return dao.report(command[1])
}

func (s *StopCommand) execute() string {
	return "Server stopped"
}

func (e *ErrorCommand) execute() string {
	return "Unknown command: " + command[0]
}

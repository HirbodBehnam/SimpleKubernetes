package master

import "github.com/go-faster/errors"

var InsufficientResourceErr = errors.New("insufficient resource in this slave")
var TasksFullErr = errors.New("tasks full in this node")

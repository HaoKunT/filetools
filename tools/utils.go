package tools

import "errors"

type status int

const (
	Equal = iota
	ascending
	descending
)

const INT_MAX = int(^uint(0) >> 1)

var breakWalkError = errors.New("break walk")

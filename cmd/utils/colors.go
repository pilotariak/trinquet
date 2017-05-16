package utils

import (
	"github.com/fatih/color"
)

var (
	GreenOut  = color.New(color.FgGreen).SprintFunc()
	YellowOut = color.New(color.FgYellow).SprintFunc()
	RedOut    = color.New(color.FgRed).SprintFunc()
)

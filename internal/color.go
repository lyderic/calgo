package internal

import "github.com/fatih/color"

var (
	Green  = color.New(color.FgGreen).PrintfFunc()
	Red    = color.New(color.FgRed).PrintfFunc()
	Yellow = color.New(color.FgYellow).PrintfFunc()
	Cyan   = color.New(color.FgCyan).PrintfFunc()
	Blue   = color.New(color.FgBlue).PrintfFunc()
)

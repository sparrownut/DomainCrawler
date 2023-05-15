package utils

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var using = false

func Printsuc(text string, args ...any) string {
while:
	if using {
		goto while
	}
	using = true
	c := color.New(color.FgHiGreen, color.Bold)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[+]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	using = false
	return "[+]" + fmt.Sprintf(text+"\n", args...)
}
func Printerr(text string, args ...any) string {
while:
	if using {
		goto while
	}
	using = true
	c := color.New(color.FgHiRed)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[-]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	using = false
	return "[-]" + fmt.Sprintf(text+"\n", args...)
}
func Printminfo(text string, args ...any) string {
while:
	if using {
		goto while
	}
	using = true
	c := color.New(color.FgYellow)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Printf("[>]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	using = false
	return "[>]" + fmt.Sprintf(text+"\n", args...)
}
func Printhinfo(text string, args ...any) string {
while:
	if using {
		goto while
	}
	using = true
	c := color.New(color.FgHiYellow, color.Bold)

	if !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[!]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	using = false
	return "[!]" + fmt.Sprintf(text+"\n", args...)
}
func Printcritical(text string, args ...any) {
while:
	if using {
		goto while
	}
	using = true
	c := color.New(color.FgHiBlue, color.BgHiRed, color.Bold)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[■]")

		if err != nil {
			return
		}

	}
	_, err := c.Printf(text+"\n", args...)

	if err != nil {
		return
	}
	fmt.Print()
	using = false
}

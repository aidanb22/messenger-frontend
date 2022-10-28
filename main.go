package main

import (
	_ "github.com/42wim/go.rice"
	"github.com/ablancas22/messenger-frontend/cmd"
)

//go:generate rice embed-go

func main() {
	a := cmd.App{}
	a.Initialize("production")
	a.Run()
}

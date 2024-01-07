package main

import (
	"os/exec"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var checkSwitch = goyek.Define(goyek.Task{
	Name:  "checkSwitch",
	Usage: "Performs an exhaustive check on the switch statements",
	Action: func(a *goyek.A) {

		if _, err := exec.LookPath("exhaustive"); err != nil {
			cmd.Exec(a, "go install github.com/nishanths/exhaustive/cmd/exhaustive@latest")
		}

		cmd.Exec(a, "exhaustive ./...", cmd.Dir("../"))
	},
})

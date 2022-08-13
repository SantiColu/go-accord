package client

import (
	"fmt"
	"os"

	"github.com/SantiColu/go-accord/pkg/tui"
)

func Run() {

	t := tui.NewTUI()
	err := t.Start()
	if err != nil {
		fmt.Printf("failed to start: %v\n", err)
		os.Exit(2)
	}

}

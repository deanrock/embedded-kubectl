package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/kubernetes/pkg/kubectl/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewDefaultKubectlCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

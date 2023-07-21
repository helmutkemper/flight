package main

import (
	"github.com/helmutkemper/chaos/factory"
	"testing"
	"time"
)

func TestLocalDevOps(t *testing.T) {

	primordial := factory.NewPrimordial().
		NetworkCreate("test_network", "10.0.0.0/16", "10.0.0.1").
		Test(t, "./end")

	// Container with test project archived in a local folder, "./mongodbClient"
	factory.NewContainerFromFolder(
		"server:latest",
		"./server",
		//"./cmd/server",
	).
		// Passing the connection through environment var makes the code more organized
		EnvironmentVar(
			[]string{
				"SERVER_PORT=8081",
				"SERVER_PORT=8082",
				"SERVER_PORT=8083",
			},
		).
		// Mount the dockerfile automatically
		MakeDockerfile().
		// Wait for the container to run
		WaitForFlagTimeout("Server started at", 10*time.Second).
		FailFlag("./bug", "panic:").
		Create("server", 3).
		Start()

	if !primordial.Monitor(10 * time.Minute) {
		t.Fail()
	}
}

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
	factory.NewContainerFromGit(
		"server:latest",
		"https://github.com/helmutkemper/flight.git",
		//"./cmd/server",
	).
		// Passing the connection through environment var makes the code more organized
		Ports("tcp", 8080, 8081, 8082, 8083).
		ReplaceBeforeBuild("./Dockerfile", "./cmd/localDevOps/Dockerfile-server").
		//DockerfilePath("./cmd/localDevOps/Dockerfile").
		// Wait for the container to run
		WaitForFlagTimeout("Server started at", 10*time.Second).
		FailFlag("./bug", "panic:").
		Create("server", 3).
		Start()

	factory.NewContainerFromGit(
		"proxy:latest",
		"https://github.com/helmutkemper/flight.git",
		//"./cmd/server",
	).
		// Passing the connection through environment var makes the code more organized
		Ports("tcp", 9999, 9999).
		ReplaceBeforeBuild("./Dockerfile", "./cmd/localDevOps/Dockerfile-proxy").
		//DockerfilePath("./cmd/localDevOps/Dockerfile").
		// Wait for the container to run
		WaitForFlagTimeout("Starting proxy", 10*time.Second).
		FailFlag("./bug", "panic:").
		Create("proxy", 1).
		Start()

	if !primordial.Monitor(10 * time.Minute) {
		t.Fail()
	}
}

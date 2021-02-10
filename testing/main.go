package main

import (
	"log"
	"os/exec"
	"regexp"
)

func main() {
	dockerContainers := runCommand("docker container ls")
	// fmt.Printf("%s\n", dockerContainers)

	//Check we actually have the system running by looking for certain containers
	match, _ := regexp.MatchString("distributed-storage_master", string(dockerContainers))
	if !match {
		log.Fatalf("Can't run tests, the system doesn't appear to be running...\n")
	}

	//TEST 1
	//Check if a container is able to be restart and join the network successfully

	//TEST 2
	//Check that the neighbours for the containers are correct

	//TEST 3
	//Check that new containers get added to the tail

	//TEST 4
	//Check that you can't add a container to the chain that already exists in the chain
}

func runCommand(cmd string) string {
	output, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("Error running command '%v' - %v", cmd, err.Error())
	}
	return string(output)
}

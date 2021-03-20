package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"
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
	//Check that the neighbours for the containers are correct
	masterState := runCommand("docker-compose logs --no-log-prefix master | tail -5")
	checkNeighbours(masterState)

	//TEST 2
	//Check if a container is able to be restart and join the network successfully. Will do node-0
	_ = runCommand("docker-compose kill node-0")
	time.Sleep(2 * time.Second)
	//Make sure the node isnt alive
	dockerContainers = runCommand("docker container ls")
	match, _ = regexp.MatchString("distributed-storage_node-0", string(dockerContainers))
	if match {
		log.Fatalf("Node is still alive which shouldnt be\n")
	}
	//Restart the container
	_ = runCommand("docker-compose start node-0")
	time.Sleep(1 * time.Second)
	dockerContainers = runCommand("docker container ls")
	match, _ = regexp.MatchString("distributed-storage_node-0", string(dockerContainers))
	if !match {
		log.Fatalf("Node isn't alive that should be\n")
	}

	//TEST 3
	//Check that new containers get added to the tail, check that node-0 is the tail
	masterState = runCommand("docker-compose logs --no-log-prefix master | tail -5")
	tail := getTailNode(masterState)
	if tail != "node-0:7000" {
		log.Fatalf("Node 0 is not at the tail")
	}

	//TEST 4
	//Add a key value pair to the system
	addKV := runCommand("make request OP=write KEY=test VALUE=testing")
	checkKeyValueWrite(addKV, "test", "testing")

	readKV := runCommand("make request OP=read KEY=test")
	checkKeyValueRead(readKV, "testing")

	addKV2 := runCommand("make request OP=write KEY=test VALUE=hello")
	checkKeyValueWrite(addKV2, "test", "hello")

	addKV3 := runCommand("make request OP=write KEY=chain VALUE=replication")
	checkKeyValueWrite(addKV3, "chain", "replication")

	//Kill a node and see if the data is retained still
	_ = runCommand("docker-compose kill node-0")
	time.Sleep(2 * time.Second)
	readKV2 := runCommand("make request OP=read KEY=chain")
	checkKeyValueRead(readKV2, "replication")

	fmt.Println("All tests successful!")
}

func runCommand(cmd string) string {
	output, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("Error running command '%v' - %v", cmd, err.Error())
	}
	return string(output)
}

func checkNeighbours(input string) {
	splitLines := strings.Split(input, "\n")[1:4]

	lineOne := strings.Split(splitLines[0], "|")
	lineTwo := strings.Split(splitLines[1], "|")
	lineThree := strings.Split(splitLines[2], "|")

	if (lineOne[1] != lineTwo[0]) && (lineTwo[1] != lineThree[0]) && (lineOne[0] != "") && (lineThree[2] != "") {
		log.Fatalln("Neighbours dont match")
	}
}

func getTailNode(input string) string {
	splitLines := strings.Split(input, "\n")[1:4]
	lineThree := strings.Split(splitLines[2], "|")

	return lineThree[1]
}

func checkKeyValueWrite(input, expectedKey, expectedValue string) {
	lines := strings.Split(input, "\n")
	split := strings.Split(lines[1], ": ")[1]
	keyValue := strings.Split(split, "->")
	if expectedKey != keyValue[0] || expectedValue != keyValue[1] {
		log.Fatalf("Key value doesnt match expected values")
	}
}

func checkKeyValueRead(input, expectedValue string) {
	lines := strings.Split(input, "\n")
	split := strings.Split(lines[2], ": ")[1]
	if split != expectedValue {
		log.Fatalf("Value not matching the expected value")
	}
}

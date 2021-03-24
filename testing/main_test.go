package main

import (
	"log"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSystemAlive(t *testing.T) {
	//Make sure all of the expected containers are alive and running so we cant test properly
	dockerContainers := string(runCommand("docker container ls"))
	assert.Contains(t, dockerContainers, "master", "Master service is not alive  - restart the system and try again")
	assert.Contains(t, dockerContainers, "node-0", "Node 0 is not alive  - restart the system and try again")
	assert.Contains(t, dockerContainers, "node-1", "Node 1 is not alive  - restart the system and try again")
	assert.Contains(t, dockerContainers, "node-2", "Node 2 is not alive  - restart the system and try again")
}

func TestNeighbours(t *testing.T) {
	//Get the neighbours for the master
	masterState := runCommand("docker-compose logs --no-log-prefix master | tail -5")
	splitLines := strings.Split(masterState, "\n")[1:4]

	//Take each line (one for each node in the chain) and split them on the "\" character to determine their neighbours
	lineOne := strings.Split(splitLines[0], "|")
	lineTwo := strings.Split(splitLines[1], "|")
	lineThree := strings.Split(splitLines[2], "|")

	assert.Equal(t, lineOne[1], lineTwo[0], "Node 2's predecessor is not node 1") //these are symmetrical so don't need to test every case
	assert.Equal(t, lineOne[1], lineTwo[0], "Node 2's successor is not node 3")   //these are symmetrical so don't need to test every case
	assert.Equal(t, lineOne[0], "", "The first chain node (head) should have no predecessor")
	assert.Equal(t, lineThree[2], "", "The last chain node (tail) should have no successor")
}

func TestNodeKill(t *testing.T) {
	_ = runCommand("docker-compose kill node-0")

	//Make sure the node isnt alive (docker wise)
	dockerContainers := string(runCommand("docker container ls"))
	assert.NotContains(t, dockerContainers, "node-0", "Node 0 is still alive after being killed")

	//Make sure the node is not alive for the chain
	time.Sleep(4 * time.Second)
	masterState := runCommand("docker-compose logs --no-log-prefix master | tail -4")
	splitLines := strings.Split(masterState, "\n")
	assert.Contains(t, splitLines[0], "Neighbour Info", "Incorrect neighbour info used")
	assert.NotContains(t, splitLines[1], "node-0", "Node 0 still in master chain state after being killed")
	assert.NotContains(t, splitLines[2], "node-0", "Node 0 still in master chain state after being killed")
}

func TestNodeStart(t *testing.T) {
	_ = runCommand("docker-compose start node-0")
	time.Sleep(2 * time.Second)
	dockerContainers := string(runCommand("docker container ls"))
	assert.Contains(t, dockerContainers, "node-0", "Node 0 is not alive, even after being restarted")

	time.Sleep(3 * time.Second)
	masterState := runCommand("docker-compose logs --no-log-prefix master | tail -5")
	splitLines := strings.Split(masterState, "\n")
	assert.Contains(t, splitLines[0], "Neighbour Info", "Incorrect neighbour info used")
	assert.Contains(t, splitLines[3], "node-0", "Node 0 is not in the master chain's state after being restarted")
}

func TestNewNodeIsTail(t *testing.T) {
	masterState := runCommand("docker-compose logs --no-log-prefix master | tail -5")
	splitLines := strings.Split(masterState, "\n")

	//Check it's the tail
	lineFour := strings.Split(splitLines[3], "|")
	assert.Contains(t, lineFour[1], "node-0", "The newly started node (node-0) is not the tail.")

	//Check the tail has no successor
	assert.Equal(t, lineFour[2], "", "The new tail node has a neighbour")

	//Check that the new tails predecessor has the correct neighbour information
	lineThree := strings.Split(splitLines[2], "|")
	assert.Equal(t, lineThree[2], lineFour[1], "Incorrect neighbours for the new tails predecessor")
}

func TestReadEmpty(t *testing.T) {
	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")

	assert.Contains(t, readKV, "Requesting read: test", "Error reading the key provided")
	assert.Contains(t, readKV, "Recieved empty value", "Incorrect value received from the read operations")
}

func TestWriteRequest(t *testing.T) {
	addKV := runCommand("docker-compose run --rm -e OP=write -e KEY=test -e VALUE=testing client")

	assert.Contains(t, addKV, "Requesting write: test->testing", "Error creating write request for key=test, value=testing")
	assert.Contains(t, addKV, "Write persisted", "Error writing test->testing to the chain")
}

func TestReadRequest(t *testing.T) {
	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")

	assert.Contains(t, readKV, "Requesting read: test", "Error reading the value from key=test")
	assert.Contains(t, readKV, "Recieved value: testing", "Incorrect value read from the chain for key=test")
}

func TestSecondWriteRequestOverwrite(t *testing.T) {
	addKV := runCommand("docker-compose run --rm -e OP=write -e KEY=test -e VALUE=hello client")

	assert.Contains(t, addKV, "Requesting write: test->hello", "Error creating write request for key=test, value=hello")
	assert.Contains(t, addKV, "Write persisted", "Error overwriting the key=test with value=hello")
}

func TestReadRequestOverwrite(t *testing.T) {
	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")

	assert.Contains(t, readKV, "Requesting read: test", "Error reading the value from key=test")
	assert.Contains(t, readKV, "Recieved value: hello", "Incorrect value read from the chain for key=test, should have been overwritten to value=hello from value=testing")
}

func TestReadAfterTailKill(t *testing.T) {
	//Kill the tail and then check that we can still read (from the middle chain node now)
	_ = runCommand("docker-compose kill node-0")
	time.Sleep(3 * time.Second)

	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")

	assert.Contains(t, readKV, "Requesting read: test", "Error reading the value from key=test from new tail")
	assert.Contains(t, readKV, "Recieved value: hello", "Incorrect value read from the chain for key=test. (after tail kill)")
}

func TestBatchDataTransfer(t *testing.T) {
	//Kill the tail and then check that we can still read (from the middle chain node now)
	_ = runCommand("docker-compose start node-0")
	time.Sleep(3 * time.Second)

	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")

	assert.Contains(t, readKV, "Requesting read: test", "Error reading the value from key=test from new tail")
	assert.Contains(t, readKV, "Recieved value: hello", "Incorrect value read from the chain for key=test. (after new tail batch transfer)")
}

func TestOneNodeSystem(t *testing.T) {
	//Testing whether the system works with only 1 node
	_ = runCommand("docker-compose kill node-1")
	_ = runCommand("docker-compose kill node-2")
	time.Sleep(6 * time.Second)

	readKV := runCommand("docker-compose run --rm -e OP=read -e KEY=test -e VALUE= client")
	assert.Contains(t, readKV, "Requesting read: test", "Error reading the value from key=test from one node system")
	assert.Contains(t, readKV, "Recieved value: hello", "Incorrect value read from the chain for key=test. (one node system)")

	writeKV := runCommand("docker-compose run --rm -e OP=write -e KEY=one -e VALUE=two client")
	assert.Contains(t, writeKV, "Requesting write: one->two", "Error creating write request for key=test, value=hello")
	assert.Contains(t, writeKV, "Write persisted", "Error overwriting the key=test with value=hello")

	readKV2 := runCommand("docker-compose run --rm -e OP=read -e KEY=one -e VALUE= client")
	assert.Contains(t, readKV2, "Requesting read: one", "Error reading the value from key=one from one node system")
	assert.Contains(t, readKV2, "Recieved value: two", "Incorrect value read from the chain for key=one. (one node system)")
}

func runCommand(cmd string) string {
	output, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("Error running command '%v' - %v", cmd, err.Error())
	}
	return string(output)
}

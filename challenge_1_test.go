package main

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	// log.Println("\nbefore udpate: ", msg)
	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	// log.Println("\nafter udpate: ", msg)
	if msg != "Hello, universe!" {
		t.Errorf("Expected updated value to be: Hello, universe!, but found: %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello, world!"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	log.Println("\nLog output: ", output)
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find: Hello, world!, but got %s", output)
	}
}

func Test_challenge(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	challenge()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	log.Println("\nLog output: ", output)
	if !strings.Contains(output, "universe") {
		t.Errorf("Expected to find: universe, but got %s", output)
	}
	if !strings.Contains(output, "cosmos") {
		t.Errorf("Expected to find: cosmos, but got %s", output)
	}
	if !strings.Contains(output, "world") {
		t.Errorf("Expected to find: world, but got %s", output)
	}
}

package main

import (
	"bufio"
	"bytes"
	"io"
	"os"

	//"strconv"
	"strings"
	"testing"
)

func Test_tableTest_isPrime(t *testing.T){
	
	primeTests := []struct{
		name string
		testNum int
		expected bool
		msg string
	}{
		{"prime", 7, true,"7 is a prime number!"},
		{"zero", 0, false,"0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"not prime", 8, false,"8 is not a prime number because it is divisible by 2"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
		
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result{
			t.Errorf("%s expected true but got false!", e.name)
		}

		if !e.expected && result{
			t.Errorf("%s expected false but got true!", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_group_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_group_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumber(t *testing.T){

	checkNumberTests := []struct{
		name string
		input string
		expected string
	}{
		{"prime", "7", "7 is a prime number"},
		{"empty", "", "Please Enter a number!"},
		{"typed", "three", "Please Enter a number!"},
		{"decimal", "1.11", "Please Enter a number!"},
		{"quit", "q", ""},
		{"not prime", "8", "8 is not a prime number because it is divisible by 2"},
		{"zero", "0","0 is not prime, by definition!"},
		{"one", "1", "1 is not prime, by definition!"},
		{"two", "2", "2 is a prime number!"},
		{"negative number", "-11", "Negative numbers are not prime, by definition!"},
	}

	for _, e := range checkNumberTests {
		inputNumber := strings.NewReader(e.input)
		scanner := bufio.NewScanner(inputNumber)
	
		res, _ := checkNumber(scanner)
	
		if !strings.Contains(string(res), e.expected) {
			t.Errorf("%s: %s expected, got %s", e.name, e.expected, string(res))
		}
	
		// _, err := strconv.Atoi(res)
		// if err != nil {
		// 	t.Errorf("Expected a number; got %s", err)
		// }
	}
	

}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
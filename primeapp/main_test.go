package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	input := "17\nq\nabc\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	expected := []struct {
		msg string
		res bool
	}{
		{"17 is prime", false},
		{"", true},
		{"Please enter a whole number!", false},
	}

	for i, e := range expected {
		msg, quit := checkNumbers(scanner)

		if msg != e.msg {
			t.Errorf("Test %d: expected message %q but got %q", i, e.msg, msg)
		}

		if quit != e.res {
			t.Errorf("Test %d: expected quit %t but got %t", i, e.res, quit)
		}
	}
}

func Test_prompt(t *testing.T) {
	rescue := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescue

	if string(out) != "this is value: test" {
		t.Errorf("Expected %s, got %s", "this is value: test", out)
	}
}

func Test_intro(t *testing.T) {
	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	w.Close()
	os.Stdout = temp

	var buf bytes.Buffer
	io.Copy(&buf, r)

	expected := "Is it Prime?\n" +
		"------------\n" +
		"Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n" +
		"-> "

	if buf.String() != expected {
		t.Errorf("Unexpected output from intro(). Got: %s, Expected: %s", buf.String(), expected)
	}
}

func Test_readUserInput(t *testing.T) {
	input := "5\nq\n"
	reader := strings.NewReader(input)

	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	doneChan := make(chan bool)

	go readUserInput(reader, doneChan)

	<-doneChan

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = temp

	expected := "Please enter a whole number!\n-> 7 is a prime number!\n-> Goodbye.\n"

	if buf.String() != expected {
		t.Errorf("readUserInput failed, expected %s but got %s", expected, buf.String())
	}
}

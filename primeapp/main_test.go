package main

import (
	"bufio"
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
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "this is value: test" {
		t.Errorf("Expected %s, got %s", "this is value: test", out)
	}
}

package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	subproc := exec.Command("./sedplus")
	input := "hello world\n"
	subproc.Stdin = strings.NewReader(input)
	output, err := subproc.Output()
	if err != nil {
		t.Error(err)
	}
	if input != string(output) {
		t.Errorf("Wanted: %v, Got: %v", input, string(output))
	}
	subproc.Wait()
}

func Test_numericOnly(t *testing.T) {
	check(t, numericOnly("abc123$Dsb"), "123")
}
func Test_numericOnly_onlynumbers(t *testing.T) {
	check(t, numericOnly("987654321"), "987654321")
}

func check(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Error("Expected to get back", expected, "but got", actual)
	}
}

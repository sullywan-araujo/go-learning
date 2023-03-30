package main

import "testing"

func TestOla(t *testing.T) {
	resultName := Hello("Bruna")
	waitName := WaitHello("Bruna")

	if resultName != waitName {
		t.Errorf("resultado '%s', esperado '%s'", resultName, waitName)
	}
}

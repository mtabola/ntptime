package ntptime

import (
	"os"
	"os/exec"
	"testing"
)

func TestPrintPresiceTime(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		PrintPresiceTime()
		return
	}
	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestPrintPresiceTime")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process run with err %v, want exit status 1", err)
}

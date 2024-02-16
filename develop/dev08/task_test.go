package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestOS(t *testing.T) {
	testCases := []struct {
		command  string
		expected string
		checker  func() string
	}{
		{
			command: "exec firefox",
			checker: func() string {
				check := exec.Command("pgrep", "firefox")
				if check.Run() != nil {
					return "false"
				}
				return "true"
			},
			expected: "true",
		},
		{
			command: "cd ..",
			checker: func() string {
				dir, _ := os.Getwd()
				current := strings.Split(dir, "/")
				return current[len(current)-1]
			},
			expected: "develop",
		},
		{
			command: "cd ../develop/dev08",
			checker: func() string {
				dir, _ := os.Getwd()
				log.Println(dir)
				current := strings.Split(dir, "/")
				return current[len(current)-1]
			},
			expected: "dev08",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		args := strings.Split(tc.command, " ")
		err := CommandExecute(args)
		if err != nil {
			t.Fatal(err)
		}
		if res := tc.checker(); res != tc.expected {
			t.Errorf("something went wrong, %v not equal %v", res, tc.expected)
		}
	}

}

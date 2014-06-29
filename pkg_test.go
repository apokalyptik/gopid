package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestPid(t *testing.T) {
	pid := os.Getpid()
	fp, err := Do("./test.pid")
	if err != nil {
		t.Errorf("Expecting no error on first lock attempt, got %s", err.Error())
	}
	if fp == nil {
		t.Errorf("Expecting fd on first lock attempt, got nil")
	} else {
		fp.Seek(0, 0)
		contents, err := ioutil.ReadAll(fp)
		if err != nil {
			t.Errorf("Expected no error reading pidfile, got %s", err.Error())
		}
		if string(contents) != fmt.Sprintf("%d", pid) {
			t.Errorf("Expected %d as file contents, got %s", pid, string(contents))
		}
	}
	fp2, err := Do("./test.pid")
	if err == nil {
		t.Errorf("Expecting error no second lock attempt, got nil")
	}
	if fp2 != nil {
		t.Errorf("Expecting no fp on second lock attempt, got %#v", fp2)
	}
	os.Remove("./test.pid")
}

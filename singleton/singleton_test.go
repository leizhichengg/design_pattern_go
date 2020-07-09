package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance1()
	instance2 := GetInstance1()
	if instance1 != instance2 {
		t.Fatal("Instance not equal.")
	}
}

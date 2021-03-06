package note

import (
	"github.com/SUSE/saptune/system"
	"testing"
)

func TestVmwareGuestIOElevator(t *testing.T) {
	if !system.IsUserRoot() {
		t.Skip("the test requires root access")
	}
	ioel, err := VmwareGuestIOElevator{}.Initialise()
	if ioel.Name() == "" {
		t.Fatal(ioel.Name())
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice) == 0 {
		t.Skip("the test case will not continue because inspection result turns out empty")
	}
	for name, elevator := range ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice {
		if name == "" || elevator == "" {
			t.Fatal(ioel)
		}
	}
	optimised, err := ioel.Optimise()
	if err != nil {
		t.Fatal(err)
	}
	if len(ioel.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice) == 0 {
		t.Fatal(ioel)
	}
	// All elevators now must be set to noop
	for name, elevator := range optimised.(VmwareGuestIOElevator).BlockDeviceSchedulers.SchedulerChoice {
		if name == "" || elevator != "noop" {
			t.Fatal(optimised)
		}
	}
}

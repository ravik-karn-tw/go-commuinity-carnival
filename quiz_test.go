package greeting

import (
	"reflect"
	"testing"
	"time"
)

func TestWelcome(t *testing.T) {
	expected := "Welcome to Community Carnival"
	actual := Welcome()
	if actual != expected {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

func TestCelebrationBeforeIPODate(t *testing.T) {
	expected := "Part of big celebration"
	actual := Celebration(time.Date(2018, time.September, 16, 0, 0, 0, 0, time.UTC))
	if actual != expected {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

func TestCelebrationAfterIPODate(t *testing.T) {
	expected := "Missed the big celebration"
	actual := Celebration(time.Date(2022, time.January, 16, 0, 0, 0, 0, time.UTC))
	if actual != expected {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

func TestNumberOfVisitors(t *testing.T) {
	expected := 18
	actual := NumberOfVisitors(5)
	if actual != expected {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

func TestNumberOfThoughtworker(t *testing.T) {
	expected := 3400
	actual := NumberOfThoughtworker(time.Date(2019, time.January, 16, 0, 0, 0, 0, time.UTC))
	if actual != expected {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

func TestLeaveBalance(t *testing.T) {
	expected := map[string]int{
		"id":          1,
		"casualLeave": 11,
		"annualLeave": 10,
	}
	employee := Employee{1, 12, 12}
	employee.TakeCasualLeave()
	employee.TakeAnnualLeave()
	employee.TakeAnnualLeave()
	actual := employee.LeaveBalance()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Failed %v: Expected: %v Actual: %v", t.Name(), expected, actual)
	}
}

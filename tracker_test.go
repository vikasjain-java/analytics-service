package main

import (
	"testing"
)

func TestGetAllVisitors(t *testing.T) {
	tracker := NewVisitorTracker()

	tracker.AddVisit("http://foo.com", "Alice")
	tracker.AddVisit("http://foo.com", "Bob")
	tracker.AddVisit("http://bar.com", "Alice")
	tracker.AddVisit("http://bar.com", "Bob")
	tracker.AddVisit("http://bar.com", "Charlie")

	allVisitors := tracker.GetAllVisitors()

	if count, exists := allVisitors["http://foo.com"]; !exists || count != 2 {
		t.Errorf("Expected 2 unique visitors for http://foo.com, got %d", count)
	}

	if count, exists := allVisitors["http://bar.com"]; !exists || count != 3 {
		t.Errorf("Expected 3 unique visitors for http://bar.com, got %d", count)
	}
}

package perf

import (
	"testing"
)

func TestAvailableEvents(t *testing.T) {
	events, err := AvailableEvents()
	if err != nil {
		t.Fatal(err)
	}
	if len(events) == 0 {
		t.Fatalf("Expected available events, got: %v", events)
	}
}

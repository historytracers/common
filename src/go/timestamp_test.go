package common

import (
	"strconv"
	"testing"
)

func TestHTUpdateTimestamp(t *testing.T) {
	ts := HTUpdateTimestamp()
	if ts == "" {
		t.Error("expected non-empty timestamp")
	}
	_, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		t.Errorf("expected numeric timestamp, got %q: %v", ts, err)
	}
}

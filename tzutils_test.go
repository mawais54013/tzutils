package tzutils

import (
    "testing"
)

func TestCurrentTimeIn(t *testing.T) {
    _, err := CurrentTimeIn("America/New_York")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
}

func TestConvertTime(t *testing.T) {
    nyTime, _ := CurrentTimeIn("America/New_York")
    _, err := ConvertTime(nyTime, "America/New_York", "Europe/London")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
}

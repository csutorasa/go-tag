package gotagio_test

import (
	"testing"
	"time"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestTimeWriter(t *testing.T) {
	var d time.Time
	writeValue(t, gotagio.NewTimeWriter(time.DateTime), &d, "2024-09-16 21:43:25")
	if d != time.Date(2024, 9, 16, 21, 43, 25, 0, time.UTC) {
		t.Fatal("failed to set time")
	}
}

func TestTimeReader(t *testing.T) {
	d := time.Date(2024, 9, 16, 21, 43, 25, 0, time.UTC)
	s := readValue(t, gotagio.NewTimeReader(time.DateTime), d)
	if s != "2024-09-16 21:43:25" {
		t.Fatal("failed to read time")
	}
}

func TestWriteDuration(t *testing.T) {
	var d time.Duration
	writeValue(t, gotagio.WriteDuration, &d, "42m")
	if d != 42*time.Minute {
		t.Fatal("failed to set duration")
	}
}

func TestReadDuration(t *testing.T) {
	d := 42 * time.Minute
	s := readValue(t, gotagio.ReadDuration, d)
	if s != "42m0s" {
		t.Fatal("failed to read duration")
	}
}

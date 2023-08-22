package must

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	// t.Parallel() - cannot test in parallel because of 'fatal' global.
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseDuration("333s"), time.Duration(333*time.Second); got != want {
		t.Errorf("ParseDuration = %v, want %v", got, want)
	}
	if _ = ParseDuration("ABC"); fakeFatalErr == nil {
		t.Error("ParseDuration: fakeFatalGot = nil, want err")
	}
}

func TestParseTime(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseTime(time.RFC3339, "2021-12-06T00:00:00Z"), time.Date(2021, 12, 6, 0, 0, 0, 0, time.UTC); got != want {
		t.Errorf("ParseTime = %v, want %v", got, want)
	}
	if _ = ParseTime(time.RFC3339, "ABC"); fakeFatalErr == nil {
		t.Error("ParseTime: fakeFatalGot = nil, want err")
	}
}

func TestParseTimeInLocation(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseTimeInLocation(time.RFC3339, "2021-12-06T00:00:00Z", time.UTC), time.Date(2021, 12, 6, 0, 0, 0, 0, time.UTC); got != want {
		t.Errorf("ParseTimeInLocation = %v, want %v", got, want)
	}
	if _ = ParseTimeInLocation(time.RFC3339, "ABC", time.UTC); fakeFatalErr == nil {
		t.Error("ParseTimeInLocation: fakeFatalGot = nil, want err")
	}
}

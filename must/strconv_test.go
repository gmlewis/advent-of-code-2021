package must

import "testing"

func TestAtoi(t *testing.T) {
	// t.Parallel() - cannot test in parallel because of 'fatal' global.
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := Atoi("333"), 333; got != want {
		t.Errorf("Atoi = %v, want %v", got, want)
	}
	if _ = Atoi("ABC"); fakeFatalErr == nil {
		t.Error("Atoi: fakeFatalGot = nil, want err")
	}
}

func TestParseBool(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseBool("true"), true; got != want {
		t.Errorf("ParseBool = %v, want %v", got, want)
	}
	if _ = ParseBool("ABC"); fakeFatalErr == nil {
		t.Error("ParseBool: fakeFatalGot = nil, want err")
	}
}

func TestParseComplex(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseComplex("1+2i", 64), complex(1, 2); got != want {
		t.Errorf("ParseComplex = %v, want %v", got, want)
	}
	if _ = ParseComplex("ABC", 64); fakeFatalErr == nil {
		t.Error("ParseComplex: fakeFatalGot = nil, want err")
	}
}

func TestParseFloat(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseFloat("1", 64), 1.0; got != want {
		t.Errorf("ParseFloat = %v, want %v", got, want)
	}
	if _ = ParseFloat("ABC", 64); fakeFatalErr == nil {
		t.Error("ParseFloat: fakeFatalGot = nil, want err")
	}
}

func TestParseInt(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseInt("1", 10, 64), 1; got != want {
		t.Errorf("ParseInt = %v, want %v", got, want)
	}
	if _ = ParseInt("ABC", 10, 64); fakeFatalErr == nil {
		t.Error("ParseInt: fakeFatalGot = nil, want err")
	}
}

func TestParseUint(t *testing.T) {
	// t.Parallel()
	fatal = fakeFatal
	fakeFatalErr = nil

	if got, want := ParseUint("1", 10, 64), uint64(1); got != want {
		t.Errorf("ParseUint = %v, want %v", got, want)
	}
	if _ = ParseUint("ABC", 10, 64); fakeFatalErr == nil {
		t.Error("ParseUint: fakeFatalGot = nil, want err")
	}
}

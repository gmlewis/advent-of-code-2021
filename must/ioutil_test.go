package must

import (
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	fatal = fakeFatal
	fakeFatalErr = nil

	buf := ReadFile("ioutil_test.go")
	if want := "package must"; !strings.HasPrefix(buf, want) {
		t.Errorf("ReadFile = %v, want %v", buf, want)
	}

	buf = ReadFile("bogus-file-name")
	if fakeFatalErr == nil {
		t.Error("ReadFile: fakeFatalGot = nil, want err")
	}
}

func TestReadFileLines(t *testing.T) {
	fatal = fakeFatal
	fakeFatalErr = nil

	lines := ReadFileLines("ioutil_test.go")
	if want := "package must"; len(lines) == 0 || !strings.HasPrefix(lines[0], want) {
		t.Errorf("ReadFileLines = %v, want %v", lines, want)
	}

	lines = ReadFileLines("bogus-file-name")
	if fakeFatalErr == nil {
		t.Error("ReadFileLines: fakeFatalGot = nil, want err")
	}
}

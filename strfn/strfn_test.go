package strfn

import (
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

func TestCompare(t *testing.T) {
	f := Compare("b")
	if got, want := f("a"), -1; got != want {
		t.Errorf("Compare = %v, want %v", got, want)
	}
	if got, want := f("b"), 0; got != want {
		t.Errorf("Compare = %v, want %v", got, want)
	}
	if got, want := f("c"), 1; got != want {
		t.Errorf("Compare = %v, want %v", got, want)
	}
}

func TestContains(t *testing.T) {
	f := Contains("b")
	if got, want := f("acd"), false; got != want {
		t.Errorf("Contains = %v, want %v", got, want)
	}
	if got, want := f("abcd"), true; got != want {
		t.Errorf("Contains = %v, want %v", got, want)
	}
}

func TestContainsAny(t *testing.T) {
	f := ContainsAny("abc")
	if got, want := f("def"), false; got != want {
		t.Errorf("ContainsAny = %v, want %v", got, want)
	}
	if got, want := f("cde"), true; got != want {
		t.Errorf("ContainsAny = %v, want %v", got, want)
	}
}

func TestContainsRune(t *testing.T) {
	f := ContainsRune('b')
	if got, want := f("acd"), false; got != want {
		t.Errorf("ContainsRune = %v, want %v", got, want)
	}
	if got, want := f("abcd"), true; got != want {
		t.Errorf("ContainsRune = %v, want %v", got, want)
	}
}

func TestCount(t *testing.T) {
	f := Count("ab")
	if got, want := f("acd"), 0; got != want {
		t.Errorf("Count = %v, want %v", got, want)
	}
	if got, want := f("abcd"), 1; got != want {
		t.Errorf("Count = %v, want %v", got, want)
	}
	if got, want := f("abba"), 1; got != want {
		t.Errorf("Count = %v, want %v", got, want)
	}
}

func TestEqualFold(t *testing.T) {
	f := EqualFold("yo")
	if got, want := f("acd"), false; got != want {
		t.Errorf("EqualFold = %v, want %v", got, want)
	}
	if got, want := f("Yo"), true; got != want {
		t.Errorf("EqualFold = %v, want %v", got, want)
	}
	if got, want := f("yo"), true; got != want {
		t.Errorf("EqualFold = %v, want %v", got, want)
	}
}

func TestFieldsFunc(t *testing.T) {
	f := FieldsFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("acd"), []string{"a", "d"}; !cmp.Equal(got, want) {
		t.Errorf("FieldsFunc = %v, want %v", got, want)
	}
	if got, want := f("Yo"), []string{"Yo"}; !cmp.Equal(got, want) {
		t.Errorf("FieldsFunc = %v, want %v", got, want)
	}
}

func TestHasPrefix(t *testing.T) {
	f := HasPrefix("yo")
	if got, want := f("abcd yo ho"), false; got != want {
		t.Errorf("HasPrefix = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), true; got != want {
		t.Errorf("HasPrefix = %v, want %v", got, want)
	}
}

func TestHasSuffix(t *testing.T) {
	f := HasSuffix("yo")
	if got, want := f("yo ho"), false; got != want {
		t.Errorf("HasSuffix = %v, want %v", got, want)
	}
	if got, want := f("ho yo"), true; got != want {
		t.Errorf("HasSuffix = %v, want %v", got, want)
	}
}

func TestIndex(t *testing.T) {
	f := Index("yo")
	if got, want := f("abcd"), -1; got != want {
		t.Errorf("Index = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), 0; got != want {
		t.Errorf("Index = %v, want %v", got, want)
	}
	if got, want := f("ho yo"), 3; got != want {
		t.Errorf("Index = %v, want %v", got, want)
	}
}

func TestIndexAny(t *testing.T) {
	f := IndexAny("abc")
	if got, want := f("abcd"), 0; got != want {
		t.Errorf("IndexAny = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), -1; got != want {
		t.Errorf("IndexAny = %v, want %v", got, want)
	}
	if got, want := f("yo ho chow"), 6; got != want {
		t.Errorf("IndexAny = %v, want %v", got, want)
	}
}

func TestIndexByte(t *testing.T) {
	f := IndexByte(48)
	if got, want := f("abcd"), -1; got != want {
		t.Errorf("IndexByte = %v, want %v", got, want)
	}
	if got, want := f("yo ho 0"), 6; got != want {
		t.Errorf("IndexByte = %v, want %v", got, want)
	}
}

func TestIndexFunc(t *testing.T) {
	f := IndexFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("abcd"), 2; got != want {
		t.Errorf("IndexFunc = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), -1; got != want {
		t.Errorf("IndexFunc = %v, want %v", got, want)
	}
}

func TestIndexRune(t *testing.T) {
	f := IndexRune('c')
	if got, want := f("abcd"), 2; got != want {
		t.Errorf("IndexRune = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), -1; got != want {
		t.Errorf("IndexRune = %v, want %v", got, want)
	}
}

func TestJoin(t *testing.T) {
	f := Join("")
	if got, want := f([]string{"yo", "ho"}), "yoho"; got != want {
		t.Errorf("IndexRune = %v, want %v", got, want)
	}
	f = Join(" ")
	if got, want := f([]string{"yo", "ho"}), "yo ho"; got != want {
		t.Errorf("IndexRune = %v, want %v", got, want)
	}
}

func TestLastIndex(t *testing.T) {
	f := LastIndex("yo")
	if got, want := f("abcd"), -1; got != want {
		t.Errorf("LastIndex = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), 0; got != want {
		t.Errorf("LastIndex = %v, want %v", got, want)
	}
	if got, want := f("ho yo"), 3; got != want {
		t.Errorf("LastIndex = %v, want %v", got, want)
	}
}

func TestLastIndexAny(t *testing.T) {
	f := LastIndexAny("abc")
	if got, want := f("abcd"), 2; got != want {
		t.Errorf("LastIndexAny = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), -1; got != want {
		t.Errorf("LastIndexAny = %v, want %v", got, want)
	}
	if got, want := f("yo ho chow"), 6; got != want {
		t.Errorf("LastIndexAny = %v, want %v", got, want)
	}
}

func TestLastIndexByte(t *testing.T) {
	f := LastIndexByte(48)
	if got, want := f("abcd"), -1; got != want {
		t.Errorf("LastIndexByte = %v, want %v", got, want)
	}
	if got, want := f("yo ho 0"), 6; got != want {
		t.Errorf("LastIndexByte = %v, want %v", got, want)
	}
}

func TestLastIndexFunc(t *testing.T) {
	f := LastIndexFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("abcd"), 2; got != want {
		t.Errorf("LastIndexFunc = %v, want %v", got, want)
	}
	if got, want := f("yo ho"), -1; got != want {
		t.Errorf("LastIndexFunc = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	f := Map(func(r rune) rune { return r + 2 })
	if got, want := f("abcd"), "cdef"; got != want {
		t.Errorf("Map = %v, want %v", got, want)
	}
	if got, want := f("0123"), "2345"; got != want {
		t.Errorf("Map = %v, want %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	f := Repeat(3)
	if got, want := f("abc"), "abcabcabc"; got != want {
		t.Errorf("Repeat = %v, want %v", got, want)
	}
}

func TestReplace(t *testing.T) {
	f := Replace("cd", "yo", 2)
	if got, want := f("acdcdcdcd"), "ayoyocdcd"; got != want {
		t.Errorf("Replace = %v, want %v", got, want)
	}
	if got, want := f("yo"), "yo"; got != want {
		t.Errorf("Replace = %v, want %v", got, want)
	}
}

func TestReplaceAll(t *testing.T) {
	f := ReplaceAll("cd", "yo")
	if got, want := f("acdcdcdcd"), "ayoyoyoyo"; got != want {
		t.Errorf("ReplaceAll = %v, want %v", got, want)
	}
	if got, want := f("yo"), "yo"; got != want {
		t.Errorf("ReplaceAll = %v, want %v", got, want)
	}
}

func TestSplit(t *testing.T) {
	f := Split(",")
	if got, want := f("acd,cd,cd,"), []string{"acd", "cd", "cd", ""}; !cmp.Equal(got, want) {
		t.Errorf("Split = %v, want %v", got, want)
	}
	if got, want := f("yo"), []string{"yo"}; !cmp.Equal(got, want) {
		t.Errorf("Split = %v, want %v", got, want)
	}
}

func TestSplitAfter(t *testing.T) {
	f := SplitAfter(",")
	if got, want := f("acd,cd,cd,"), []string{"acd,", "cd,", "cd,", ""}; !cmp.Equal(got, want) {
		t.Errorf("SplitAfter = %v, want %v", got, want)
	}
	if got, want := f("yo"), []string{"yo"}; !cmp.Equal(got, want) {
		t.Errorf("SplitAfter = %v, want %v", got, want)
	}
}

func TestSplitAfterN(t *testing.T) {
	f := SplitAfterN(",", 2)
	if got, want := f("acd,cd,cd,"), []string{"acd,", "cd,cd,"}; !cmp.Equal(got, want) {
		t.Errorf("SplitAfterN = %v, want %v", got, want)
	}
	if got, want := f("yo"), []string{"yo"}; !cmp.Equal(got, want) {
		t.Errorf("SplitAfterN = %v, want %v", got, want)
	}
}

func TestSplitN(t *testing.T) {
	f := SplitN(",", 2)
	if got, want := f("acd,cd,cd,"), []string{"acd", "cd,cd,"}; !cmp.Equal(got, want) {
		t.Errorf("SplitN = %v, want %v", got, want)
	}
	if got, want := f("yo"), []string{"yo"}; !cmp.Equal(got, want) {
		t.Errorf("SplitN = %v, want %v", got, want)
	}
}

func TestToLowerSpecial(t *testing.T) {
	f := ToLowerSpecial(unicode.TurkishCase)
	if got, want := f("İ"), "i"; got != want {
		t.Errorf("ToLowerSpecial = %v, want %v", got, want)
	}
}

func TestToTitleSpecial(t *testing.T) {
	f := ToTitleSpecial(unicode.TurkishCase)
	if got, want := f("İ"), "İ"; got != want {
		t.Errorf("ToTitleSpecial = %v, want %v", got, want)
	}
}

func TestToUpperSpecial(t *testing.T) {
	f := ToUpperSpecial(unicode.TurkishCase)
	if got, want := f("İ"), "İ"; got != want {
		t.Errorf("ToUpperSpecial = %v, want %v", got, want)
	}
}

func TestToValidUTF8(t *testing.T) {
	f := ToValidUTF8("yo")
	if got, want := f("\367\122"), "yoR"; got != want {
		t.Errorf("ToValidUTF8 = %v, want %v", got, want)
	}
}

func TestTrim(t *testing.T) {
	f := Trim("abc")
	if got, want := f("acd,cd,cd,"), "d,cd,cd,"; got != want {
		t.Errorf("Trim = %v, want %v", got, want)
	}
	if got, want := f("yo"), "yo"; got != want {
		t.Errorf("Trim = %v, want %v", got, want)
	}
}

func TestTrimFunc(t *testing.T) {
	f := TrimFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimFunc = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "yo"; got != want {
		t.Errorf("TrimFunc = %v, want %v", got, want)
	}
}

func TestTrimLeft(t *testing.T) {
	f := TrimLeft("c,")
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimLeft = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "yoc"; got != want {
		t.Errorf("TrimLeft = %v, want %v", got, want)
	}
}

func TestTrimLeftFunc(t *testing.T) {
	f := TrimLeftFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimLeftFunc = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "yoc"; got != want {
		t.Errorf("TrimLeftFunc = %v, want %v", got, want)
	}
}

func TestTrimPrefix(t *testing.T) {
	f := TrimPrefix("c")
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimPrefix = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "yoc"; got != want {
		t.Errorf("TrimPrefix = %v, want %v", got, want)
	}
}

func TestTrimRight(t *testing.T) {
	f := TrimRight("c")
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimRight = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "cyo"; got != want {
		t.Errorf("TrimRight = %v, want %v", got, want)
	}
}

func TestTrimRightFunc(t *testing.T) {
	f := TrimRightFunc(func(r rune) bool { return r == 'c' })
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimRightFunc = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "cyo"; got != want {
		t.Errorf("TrimRightFunc = %v, want %v", got, want)
	}
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TestTrimSuffix(t *testing.T) {
	f := TrimSuffix("c")
	if got, want := f("acd,cd,cd,"), "acd,cd,cd,"; got != want {
		t.Errorf("TrimSuffix = %v, want %v", got, want)
	}
	if got, want := f("cyoc"), "cyo"; got != want {
		t.Errorf("TrimSuffix = %v, want %v", got, want)
	}
}

// Package strfn provides curried functional versions of the functions
// provided by the "strings" package so they can be more easily
// used within "enum" and "maps" functions.
package strfn

import (
	"strings"
	"unicode"
)

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare(b string) func(a string) int {
	return func(a string) int { return strings.Compare(a, b) }
}

// Contains reports whether substr is within s.
func Contains(substr string) func(s string) bool {
	return func(s string) bool { return strings.Contains(s, substr) }
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(chars string) func(s string) bool {
	return func(s string) bool { return strings.ContainsAny(s, chars) }
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(r rune) func(s string) bool {
	return func(s string) bool { return strings.ContainsRune(s, r) }
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(substr string) func(s string) int {
	return func(s string) int { return strings.Count(s, substr) }
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold(t string) func(s string) bool {
	return func(s string) bool { return strings.EqualFold(s, t) }
}

// FieldsFunc splits the string s at each run of Unicode code
// points c satisfying f(c) and returns an array of slices of s.
// If all code points in s satisfy f(c) or the string is empty,
// an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc(f func(rune) bool) func(s string) []string {
	return func(s string) []string { return strings.FieldsFunc(s, f) }
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(prefix string) func(s string) bool {
	return func(s string) bool { return strings.HasPrefix(s, prefix) }
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(suffix string) func(s string) bool {
	return func(s string) bool { return strings.HasSuffix(s, suffix) }
}

// Index returns the index of the first instance of substr in s,
// or -1 if substr is not present in s.
func Index(substr string) func(s string) int {
	return func(s string) int { return strings.Index(s, substr) }
}

// IndexAny returns the index of the first instance of any
// Unicode code point from chars in s, or -1 if no Unicode
// code point from chars is present in s.
func IndexAny(chars string) func(s string) int {
	return func(s string) int { return strings.IndexAny(s, chars) }
}

// IndexByte returns the index of the first instance of c in s,
// or -1 if c is not present in s.
func IndexByte(c byte) func(s string) int {
	return func(s string) int { return strings.IndexByte(s, c) }
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(f func(rune) bool) func(s string) int {
	return func(s string) int { return strings.IndexFunc(s, f) }
}

// IndexRune returns the index of the first instance of the
// Unicode code point r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of
// any invalid UTF-8 byte sequence.
func IndexRune(r rune) func(s string) int {
	return func(s string) int { return strings.IndexRune(s, r) }
}

// Join concatenates the elements of its first argument to
// create a single string. The separator string sep is placed
// between elements in the resulting string.
func Join(sep string) func(elems []string) string {
	return func(elems []string) string { return strings.Join(elems, sep) }
}

// LastIndex returns the index of the last instance of substr
// in s, or -1 if substr is not present in s.
func LastIndex(substr string) func(s string) int {
	return func(s string) int { return strings.LastIndex(s, substr) }
}

// LastIndexAny returns the index of the last instance of any
// Unicode code point from chars in s, or -1 if no Unicode
// code point from chars is present in s.
func LastIndexAny(chars string) func(s string) int {
	return func(s string) int { return strings.LastIndexAny(s, chars) }
}

// LastIndexByte returns the index of the last instance
// of c in s, or -1 if c is not present in s.
func LastIndexByte(c byte) func(s string) int {
	return func(s string) int { return strings.LastIndexByte(s, c) }
}

// LastIndexFunc returns the index into s of the last Unicode
// code point satisfying f(c), or -1 if none do.
func LastIndexFunc(f func(rune) bool) func(s string) int {
	return func(s string) int { return strings.LastIndexFunc(s, f) }
}

// Map returns a copy of the string s with all its characters
// modified according to the mapping function. If mapping returns
// a negative value, the character is dropped from the string
// with no replacement.
func Map(mapping func(rune) rune) func(s string) string {
	return func(s string) string { return strings.Map(mapping, s) }
}

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count) overflows.
func Repeat(count int) func(s string) string {
	return func(s string) string { return strings.Repeat(s, count) }
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
//
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
//
// If n < 0, there is no limit on the number of replacements.
func Replace(old, new string, n int) func(s string) string {
	return func(s string) string { return strings.Replace(s, old, new, n) }
}

// ReplaceAll returns a copy of the string s with all non-overlapping
// instances of old replaced by new.
//
// If old is empty, it matches at the beginning of the string and after
// each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
func ReplaceAll(old, new string) func(s string) string {
	return func(s string) string { return strings.ReplaceAll(s, old, new) }
}

// Split slices s into all substrings separated by sep and returns
// a slice of the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence.
// If both s and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
func Split(sep string) func(s string) []string {
	return func(s string) []string { return strings.Split(s, sep) }
}

// SplitAfter slices s into all substrings after each instance
// of sep and returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter
// returns a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
// If both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(sep string) func(s string) []string {
	return func(s string) []string { return strings.SplitAfter(s, sep) }
}

// SplitAfterN slices s into substrings after each instance
// of sep and returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//
// n > 0: at most n substrings; the last substring will be the unsplit remainder.
// n == 0: the result is nil (zero substrings)
// n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are
// handled as described in the documentation for SplitAfter.
func SplitAfterN(sep string, n int) func(s string) []string {
	return func(s string) []string { return strings.SplitAfterN(s, sep, n) }
}

// SplitN slices s into substrings separated by sep and
// returns a slice of the substrings between those separators.
//
// The count determines the number of substrings to return:
//
// n > 0: at most n substrings; the last substring will be the unsplit remainder.
// n == 0: the result is nil (zero substrings)
// n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are
// handled as described in the documentation for Split.
func SplitN(sep string, n int) func(s string) []string {
	return func(s string) []string { return strings.SplitN(s, sep, n) }
}

// ToLowerSpecial returns a copy of the string s with all
// Unicode letters mapped to their lower case using the
// case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase) func(s string) string {
	return func(s string) string { return strings.ToLowerSpecial(c, s) }
}

// ToTitleSpecial returns a copy of the string s with all
// Unicode letters mapped to their Unicode title case,
// giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase) func(s string) string {
	return func(s string) string { return strings.ToTitleSpecial(c, s) }
}

// ToUpperSpecial returns a copy of the string s with all
// Unicode letters mapped to their upper case using the
// case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase) func(s string) string {
	return func(s string) string { return strings.ToUpperSpecial(c, s) }
}

// ToValidUTF8 returns a copy of the string s with each run
// of invalid UTF-8 byte sequences replaced by the replacement
// string, which may be empty.
func ToValidUTF8(replacement string) func(s string) string {
	return func(s string) string { return strings.ToValidUTF8(s, replacement) }
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(cutset string) func(s string) string {
	return func(s string) string { return strings.Trim(s, cutset) }
}

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(f func(rune) bool) func(s string) string {
	return func(s string) string { return strings.TrimFunc(s, f) }
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use TrimPrefix instead.
func TrimLeft(cutset string) func(s string) string {
	return func(s string) string { return strings.TrimLeft(s, cutset) }
}

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(f func(rune) bool) func(s string) string {
	return func(s string) string { return strings.TrimLeftFunc(s, f) }
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(prefix string) func(s string) string {
	return func(s string) string { return strings.TrimPrefix(s, prefix) }
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimRight(cutset string) func(s string) string {
	return func(s string) string { return strings.TrimRight(s, cutset) }
}

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(f func(rune) bool) func(s string) string {
	return func(s string) string { return strings.TrimRightFunc(s, f) }
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(suffix string) func(s string) string {
	return func(s string) string { return strings.TrimSuffix(s, suffix) }
}

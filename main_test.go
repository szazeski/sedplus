package main

import (
	"testing"
)

func Test_numericOnly(t *testing.T) {
	check(t, numericOnly("abc123$Dsb"), "123")
}
func Test_alphaOnly(t *testing.T) {
	check(t, alphaOnly("abc123$Dsb"), "abcDsb")
}
func Test_alphanumericOnly(t *testing.T) {
	check(t, alphanumericOnly("abc123$Dsb"), "abc123Dsb")
}
func Test_replaceTimestamps(t *testing.T) {
	check(t, replaceTimestamps("12:03:32 Log Message"), "Log Message")

	check(t, replaceTimestamps("1:2 Log Message"), "Log Message")
}

// future feature
func Test_replaceTimestamps_Apache(t *testing.T) {
	input := "1.2.3.4 - - [24/Dec/2023:22:13:43 +0000] \"GET /wp-content/uploads/2016/08/insufficientStorageSpace3000-scaled.jpg HTTP/1.1\" 200 141235 \"\n\n"
	t.SkipNow()
	check(t, replaceTimestamps(input), "\"1.2.3.4 - - [24/Dec/2023:+0000] \"GET /wp-content/uploads/2016/08/insufficientStorageSpace3000-scaled.jpg HTTP/1.1\" 200 141235 \"\n\n")
}

func check(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Error("Expected to get back", expected, "but got", actual)
	}
}

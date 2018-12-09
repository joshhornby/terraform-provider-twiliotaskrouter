package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestJsonBytesEqualQuotedAndUnquoted(t *testing.T) {
	unquoted := `{"test": "test"}`
	quoted := "{\"test\": \"test\"}"

	if !jsonBytesEqual([]byte(unquoted), []byte(quoted)) {
		t.Errorf("Expected jsonBytesEqual to return true for %s == %s", unquoted, quoted)
	}

	unquotedDiff := `{"test": "test"}`
	quotedDiff := "{\"test\": \"tested\"}"

	if jsonBytesEqual([]byte(unquotedDiff), []byte(quotedDiff)) {
		t.Errorf("Expected jsonBytesEqual to return false for %s == %s", unquotedDiff, quotedDiff)
	}
}

func TestJsonBytesEqualWhitespaceAndNoWhitespace(t *testing.T) {
	noWhitespace := `{"test":"test"}`
	whitespace := `
{
  "test": "test"
}`

	if !jsonBytesEqual([]byte(noWhitespace), []byte(whitespace)) {
		t.Errorf("Expected jsonBytesEqual to return true for %s == %s", noWhitespace, whitespace)
	}

	noWhitespaceDiff := `{"test":"test"}`
	whitespaceDiff := `
{
  "test": "tested"
}`

	if jsonBytesEqual([]byte(noWhitespaceDiff), []byte(whitespaceDiff)) {
		t.Errorf("Expected jsonBytesEqual to return false for %s == %s", noWhitespaceDiff, whitespaceDiff)
	}
}

func TestSuppressEquivalentJsonDiffsWhitespaceAndNoWhitespace(t *testing.T) {
	d := new(schema.ResourceData)
	noWhitespace := `{"test":"test"}`
	whitespace := `
{
  "test": "test"
}`
	if !suppressEquivalentJsonDiffs("", noWhitespace, whitespace, d) {
		t.Errorf("Expected suppressEquivalentJsonDiffs to return true for %s == %s", noWhitespace, whitespace)
	}
	noWhitespaceDiff := `{"test":"test"}`
	whitespaceDiff := `
{
  "test": "tested"
}`
	if suppressEquivalentJsonDiffs("", noWhitespaceDiff, whitespaceDiff, d) {
		t.Errorf("Expected suppressEquivalentJsonDiffs to return false for %s == %s", noWhitespaceDiff, whitespaceDiff)
	}
}

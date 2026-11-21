package interpreter

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	result, err := Run("2 * (3 + 4)")
	if err != nil {
		t.Fatalf("Run returned an error: %v", err)
	}
	if result != "14" {
		t.Fatalf("Run returned %q, want %q", result, "14")
	}
}

func TestRunReportsParserErrors(t *testing.T) {
	_, err := Run("1 +")
	var parseError *ParseError
	if !errors.As(err, &parseError) {
		t.Fatalf("Run error = %v, want *ParseError", err)
	}
	if len(parseError.Messages) == 0 {
		t.Fatal("ParseError has no messages")
	}
}

func TestRunReportsUnsupportedStatements(t *testing.T) {
	_, err := Run("let answer = 42;")
	if err == nil {
		t.Fatal("Run did not report an unsupported statement")
	}
}

func TestRunReportsUnsupportedExpressions(t *testing.T) {
	_, err := Run("answer + 1")
	if err == nil {
		t.Fatal("Run did not report an unsupported expression")
	}
}

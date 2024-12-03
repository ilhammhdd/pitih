package relationalDB_test

import (
	"testing"

	relationalDB "ilhammhdd.com/pitih/internal/pkg/relational-db"
)

func TestNewPlaceHolder(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{1, "(?)"},
		{2, "(?,?)"},
		{3, "(?,?,?)"},
		{10, "(?,?,?,?,?,?,?,?,?,?)"},
	}

	var result string
	for i := 0; i < len(testCases); i++ {
		result = string(relationalDB.NewPlaceHolder(testCases[i].n))
		if result != testCases[i].expected {
			t.Fatalf("n: %d expected: \"%s\" got: \"%s\"", testCases[i].n, testCases[i].expected, result)
		}
	}
}

func TestNewNPlaceHolder(t *testing.T) {
	testCases := []struct {
		nPlaceHolders, nParams int
		expected               string
	}{
		{0, 0, ""},
		{1, 1, "(?)"},
		{1, 2, "(?,?)"},
		{2, 2, "(?,?),(?,?)"},
		{2, 3, "(?,?,?),(?,?,?)"},
		{10, 2, "(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?)"}, //59
	}

	var result string
	for i := 0; i < len(testCases); i++ {
		result = string(relationalDB.NewNPlaceHolder(testCases[i].nPlaceHolders, testCases[i].nParams))
		if result != testCases[i].expected {
			t.Fatalf("nPlaceholders: %d nParams: %d expected: \"%s\" got: \"%s\"", testCases[i].nPlaceHolders, testCases[i].nParams, testCases[i].expected, result)
		}
	}
}

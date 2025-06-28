package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	}{
		{
			name: "Example from problem description (adjusted for valid date)",
			input: "0012-10-2013-10-2013", // Adjusted year to be valid
			expected: "12-10-2013",
		}, 
		{
			name: "Single valid date",
			input: "01-01-2013",
			expected: "01-01-2013",
		},
		{
			name: "Multiple valid dates, one most frequent",
			input: "01-01-2013-02-02-2014-01-01-2013",
			expected: "01-01-2013",
		},
		{
			name: "Dates with different days/months/years",
			input: "01-01-2013-01-02-2013-01-01-2013-01-03-2014",
			expected: "01-01-2013",
		},
		{
			name: "Edge case year 2015",
			input: "31-12-2015",
			expected: "31-12-2015",
		},
		{
			name: "Date at the beginning of string",
			input: "01-01-2013abc",
			expected: "01-01-2013",
		},
		{
			name: "Date at the end of string",
			input: "abc01-01-2013",
			expected: "01-01-2013",
		},
		{
			name: "Date in middle of string",
			input: "abc01-01-2013def",
			expected: "01-01-2013",
		},
		{
			name: "Multiple occurrences of the same date, interleaved with invalid dates",
			input: "01-01-2013-99-99-9999-01-01-2013-10-10-2010-01-01-2013",
			expected: "01-01-2013",
		},
		{
			name: "February 29th in a non-leap year (should be invalid)",
			input: "29-02-2013",
			expected: "", // Should be invalid, problem states 2013-2015 are not leap years, so no 29th Feb
		},
		{
			name: "February 28th in a non-leap year (should be valid)",
			input: "28-02-2013",
			expected: "28-02-2013",
		},
		{
			name: "Complex string with multiple valid and invalid dates",
			input: "01-01-2013abc02-02-2014def01-01-2013ghi03-03-2015jkl01-01-2013",
			expected: "01-01-2013",
		},
		{
			name: "Another complex string",
			input: "10-10-2013-10-10-2014-10-10-2013-11-11-2015-10-10-2013",
			expected: "10-10-2013",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := solve(tt.input)
			if actual != tt.expected {
				t.Errorf("For input %q, expected %q but got %q", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		name string
		dateStr string
		expected bool
	}{
		{"Valid date 01-01-2013", "01-01-2013", true},
		{"Valid date 31-12-2015", "31-12-2015", true},
		{"Invalid year 2012", "01-01-2012", false},
		{"Invalid year 2016", "01-01-2016", false},
		{"Invalid month 00", "01-00-2013", false},
		{"Invalid month 13", "01-13-2013", false},
		{"Invalid day 00", "00-01-2013", false},
		{"Invalid day 32 for Jan", "32-01-2013", false},
		{"Invalid day 31 for Apr", "31-04-2013", false},
		{"Invalid format missing dash", "0101-2013", false},
		{"Invalid format wrong length", "01-01-20133", false},
		{"Invalid format non-digit day", "ab-01-2013", false},
		{"Invalid format non-digit month", "01-cd-2013", false},
		{"Invalid format non-digit year", "01-01-efgh", false},
		{"February 29th 2013 (not leap year)", "29-02-2013", false},
		{"February 28th 2013 (valid)", "28-02-2013", true},
		{"February 29th 2014 (not leap year)", "29-02-2014", false},
		{"February 29th 2015 (not leap year)", "29-02-2015", false},
		{"April 30th (valid)", "30-04-2013", true},
		{"April 31st (invalid)", "31-04-2013", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isValidDate(tt.dateStr)
			if actual != tt.expected {
				t.Errorf("For date string %q, expected %v but got %v", tt.dateStr, tt.expected, actual)
			}
		})
	}
}

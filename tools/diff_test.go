package tools

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

var SimpleStringTestCases = []struct {
	first    []string
	second   []string
	expected []string
}{
	{strings.Split("ABCDEF", ""), strings.Split("ABCDEF", ""), strings.Split("ABCDEF", "")},
	{strings.Split("ABC", ""), strings.Split("XYZ", ""), strings.Split("", "")},
	{strings.Split("AABCXY", ""), strings.Split("XYZ", ""), strings.Split("XY", "")},
	{strings.Split("", ""), strings.Split("", ""), strings.Split("", "")},
	{strings.Split("ABCD", ""), strings.Split("AC", ""), strings.Split("AC", "")},
}

var SimpleArrayTestCases = []struct {
	first    []string
	second   []string
	expected []string
}{
	{[]string{"This is a test which contains:", "this is the lcs"}, []string{"this is the lcs", "we're testing"}, []string{"this is the lcs"}},
	{[]string{
		"Coding Challenges helps you become a better software engineer through that build real applications.",
		"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
		"I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
		"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities."},
		[]string{
			"Helping you become a better software engineer through coding challenges that build real applications.",
			"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
			"These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
			"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
		},
		[]string{
			"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
			"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
		}},
}

func TestFindCommonSimpleString(t *testing.T) {
	for idx, testCase := range SimpleStringTestCases {
		testCaseName := fmt.Sprintf("Running %d ", idx)
		t.Run(testCaseName, func(t *testing.T) {
			expected := testCase.expected
			calculated := FindCommon(testCase.first, testCase.second)

			if !reflect.DeepEqual(expected, calculated) {
				t.Errorf("Expected - %v, Got - %v", expected, calculated)
				t.Fail()
			}
		})
	}
}

func TestFindCommonArrayOfString(t *testing.T) {
	for idx, testCase := range SimpleArrayTestCases {
		testCaseName := fmt.Sprintf("Running %d ", idx)
		t.Run(testCaseName, func(t *testing.T) {
			expected := testCase.expected
			calculated := FindCommon(testCase.first, testCase.second)

			if !reflect.DeepEqual(expected, calculated) {
				t.Errorf("Expected - %v, Got - %v", expected, calculated)
				t.Fail()
			}
		})
	}
}

func TestDiffFile(t *testing.T) {
	originalFileData, err := os.ReadFile("../data/original.txt")

	if err != nil {
		t.Fatal(err.Error())
	}

	newFileData, err := os.ReadFile("../data/new.txt")

	if err != nil {
		t.Fatal(err.Error())
	}

	originalFileContent := strings.Split(string(originalFileData), "\n")
	newFileContent := strings.Split(string(newFileData), "\n")

	expected := []string{
		"> This is an important",
		"> notice! It should",
		"> therefore be located at",
		"> the beginning of this",
		"> document!",
		"> ",
		"< This paragraph contains",
		"< text that is outdated.",
		"< It will be deleted in the",
		"< near future.",
		"< ",
		"< check this dokument. On",
		"> check this document. On",
		"> ",
		"> This paragraph contains",
		"> important new additions",
		"> to this document.",
	}

	diff := FindDiff(originalFileContent, newFileContent)

	if !reflect.DeepEqual(diff, expected) {
		t.Errorf("Expected - %v, Got - %v", expected, diff)
		t.Fail()
	}
}

func TestDiffCCFile(t *testing.T) {
	originalFileData, err := os.ReadFile("../data/originalcc.txt")

	if err != nil {
		t.Fatal(err.Error())
	}

	newFileData, err := os.ReadFile("../data/newcc.txt")

	if err != nil {
		t.Fatal(err.Error())
	}

	originalFileContent := strings.Split(string(originalFileData), "\n")
	newFileContent := strings.Split(string(newFileData), "\n")

	expected := []string{
		"< Coding Challenges helps you become a better software engineer through that build real applications.",
		"> Helping you become a better software engineer through coding challenges that build real applications.",
		"< I’ve used or am using these coding challenges as exercises to learn a new programming language or technology.",
		"> These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
	}

	diff := FindDiff(originalFileContent, newFileContent)

	if !reflect.DeepEqual(diff, expected) {
		t.Errorf("Expected - %v, Got - %v", expected, diff)
		t.Fail()
	}
}

package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFormatEnt(t *testing.T) {
	fmt.Printf("\n\t::: Test Target formatEnt() :::\n")

	type craque struct {
		A int    // Code (0), Data (1), Platform (2), Tool (3)
		B int    // Adopt (0), Trial (1), Assess (2), Hold (3)
		C string // Tech Subject
		D bool   // Not sure what this is yet
		E string // External URL
		F int    // Ring Delta
	}

	// a full test will go through all options
	testCSV := []string{"craque", "Code", "Hold", "mattic"}
	testLoc := craque{A: 0, B: 3, C: "craque", D: true, E: ".", F: 0}
	testRow := formatEnt(testCSV)
	if cmp.Equal(testLoc.A, testRow.Quadrant) {
		// fmt.Println("equal")
	} else {
		t.Error()
	}

	// fmt.Printf("%v \n", testLoc)
	// fmt.Printf("%v \n", testRow)
}

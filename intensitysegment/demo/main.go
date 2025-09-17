package main

import (
	"fmt"

	is "github.com/lesiona-z/jerryaiinterview/intensitysegment"
)

func main() {
	// see segment_test.go for tests

	var segments = is.IntensitySegment{}
	// Should be "[]
	fmt.Println("init segments: ", segments.ToString())

	err := segments.Add(10, 30, 1)
	if err != nil {
		fmt.Println("Error adding segment:", err)
	} else {
		// Should be: "[[10,1],[30,0]]
		fmt.Println("After adding [10,30] with intensity 1: ", segments.ToString())
	}

	err = segments.Add(20, 40, 1)
	if err != nil {
		fmt.Println("Error adding segment:", err)
	} else {
		// Should be: "[[10,1],[20,2],[30,1],[40,0]]
		fmt.Println("After adding [20,40] with intensity 1: ", segments.ToString())
	}

	err = segments.Add(10, 40, -1)
	if err != nil {
		fmt.Println("Error adding segment:", err)
	} else {
		// Should be "[[20,1],[30,0]]"
		fmt.Println("After adding [10,40] with intensity -1: ", segments.ToString())
	}

	err = segments.Add(10, 40, -1)
	if err != nil {
		fmt.Println("Error adding segment:", err)
	} else {
		// Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
		fmt.Println("After adding [10,40] with intensity -1: ", segments.ToString())
	}
}

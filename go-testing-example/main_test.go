package main

import "testing"
//go test -v -cover

/* Visual of paths covered from tests*/
//first create: go test -coverprofile=coverage.out
//then generate html: go tool cover -html=coverage.out
func TestAddTwoNums(t *testing.T) {
	if AddTwoNums(1, 2) != 3 {
		t.Error("Expected 3")
	}
}


//table driven testing
func TestTableAddTwoNums(t *testing.T){
	var tests = []struct{
		input1 int
		input2 int
		expected int
	}{
		{1,2,3},
		{2,4,6},
		{3,6,9},
	}

	for _,vals := range tests{
		if output := AddTwoNums(vals.input1, vals.input2); output != vals.expected {
			t.Errorf("With inputs: %v, %v ---- We expected: %v ---- Instead got: %v", vals.input1, vals.input2, vals.expected, output)
		}
	}
}


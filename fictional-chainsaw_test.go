package main

import "testing"

func TestSquare(t *testing.T) {

    expectedOutput := 4

    if square(2) != expectedOutput {
        t.Errorf("Test failed, expected: '%d', got: '%d'", expectedOutput, square(2))
    }

}

package game

import (
	"fmt"
	"testing"
)

func TestGridToWorldHex(t *testing.T) {
	// Size 10.0 to make numbers easy
	// Width = sqrt(3)*10 = 17.32
	// Height = 2*10 = 20
	// VertStep = 1.5*10 = 15

	// Case 0,0
	res := GridToWorldHex(0, 0, 10.0)
	fmt.Printf("TEST_OUTPUT: (0,0) -> X=%f, Y=%f\n", res.X, res.Y)

	// Case 1,0 (One col right) -> x should be 17.32, y should be 0
	res = GridToWorldHex(1, 0, 10.0)
	fmt.Printf("TEST_OUTPUT: (1,0) -> X=%f, Y=%f\n", res.X, res.Y)

	// Case 0,1 (One row down) -> x should be width/2=8.66, y should be 15
	res = GridToWorldHex(0, 1, 10.0)
	fmt.Printf("TEST_OUTPUT: (0,1) -> X=%f, Y=%f\n", res.X, res.Y)
}

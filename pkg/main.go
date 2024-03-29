package main

import (
	"fmt"
	"github.com/sanleon/unit_test_sample/pkg/ops"
)

func main() {
	op := ops.GetKeyOperator()

	key := op.Generate(2, 3)

	a, b, _ := op.Degenerate(key)

	fmt.Printf("key=%v, a=%v, b=%v", key, a, b)
}

package loop_test

import (
	"testing"
	"math"
)

func TestWhileLoop(t *testing.T){
	n:=0
	for n < 5{
		t.Log(n)
		n++
	}
	t.Log(math.MaxInt64)
	t.Log(math.MinInt64)
	t.Log(math.MaxFloat64)
	// t.Log(math.MinFloat64)
	t.Log(math.MaxUint32)
	// t.Log(math.MinUint32)
}
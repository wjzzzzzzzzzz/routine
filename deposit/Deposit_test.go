package concurrent

import (
	"image"
	"testing"
)

func Test_DepositA(t *testing.T) {
	DepositA()
}

func Test_DepositB(t *testing.T) {
	DepositB()
}
func Test_Demo(t *testing.T) {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()
	x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
}
func Test_Currenetmap(t *testing.T) {

}

package num

import (
	"testing"
	"log"
)
func init(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func Test_reverse(t *testing.T) {
	i := reverse(1234)
	log.Println(i)

}
func Test_twoSum(t *testing.T) {
	sum := twoSum([]int{3, 2, 0, 4, 5, 3}, 6)
	log.Println(sum)
}

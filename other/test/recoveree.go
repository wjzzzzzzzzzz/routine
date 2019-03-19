package main

func main() {
	var attr = struct {
		perm int
		owner int
	}{2, 0755}
	var null struct{}

	set := make(map[string]struct{})
	set["a"] = null


}

type Node struct {
	_  map[string]interface{}
	id int
}
type File struct {
	name string
	size int
	attr struct {
		perm int
		owner int
	}
}

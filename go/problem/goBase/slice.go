package goBase

func sliceTest() {
	s := make([]int, 3, 3)

	s = append(s, 1)
	s[0] = 2

	fmt.Println(s)
}
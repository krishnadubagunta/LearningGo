package Slices

func AppendToSlice(slice []int, elements ...int) []int {
	return append(slice, elements...)
}

func ConcatTwoSlices(slice []int, slice2 []int) []int {
	return append(slice, slice2...)
}

package arrays

func Reverse(a []int32) []int32 {
    var aux int32
	
	for x, y := 0, len(a)-1; x < len(a)/2; x, y = x+1, y-1 {
		aux = a[x]
		a[x] = a[y]
		a[y] = aux 
	}
	
	return a
}
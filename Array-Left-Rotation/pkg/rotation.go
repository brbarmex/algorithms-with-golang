package rotation

func RotateLeft(d int32, arr []int32) []int32 {
	var aux int32
    for x := 0; x < int(d); x++ {
		aux = arr[0]
		for y:= 0; y < len(arr)-1; y++ {
			arr[y] = arr[y+1]
		}
		arr[len(arr)-1]=aux
	}

	return arr
}
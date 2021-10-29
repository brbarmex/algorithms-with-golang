package crypto

func Encrypt(v string) string {

	bytes := []byte(v)

	for x, ancii := range bytes {
		if ancii > 64 && ancii < 91 ||
			ancii > 96 && ancii < 122 {
			bytes[x] = ancii + 3
		}
	}

	var aux byte

	for x, y := 0, len(v)-1; x < len(v)/2; x, y = x+1, y-1 {
		aux = bytes[x]
		bytes[x] = bytes[y]
		bytes[y] = aux 
	}

	for x := len(v)/2; x < len(v); x++{
		bytes[x] = bytes[x]-1
	} 

	return string(bytes)
}

package decodexoredarray

func decode(encoded []int, first int) []int {
	decoded := make([]int, len(encoded)+1)

	decoded[0] = first
	for i := 0; i < len(encoded); i++ {
		decoded[i+1] = decoded[i] ^ encoded[i]
	}

	return decoded
}

package somador

// SomarNumeros recebe um slice de inteiros e retorna a soma deles.
func SomarNumeros(nums []int) int {
	soma := 0
	for _, num := range nums {
		soma += num
	}
	return soma
}


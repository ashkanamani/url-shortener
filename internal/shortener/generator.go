package shortener

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeToBase62(n uint64) string {
	if n == 0 {
		return string(alphabet[0])
	}
	var result []byte

	for n > 0 {
		r := n % 62
		result = append([]byte{alphabet[r]}, result...)
		n /= 62
	}
	return string(result)
}

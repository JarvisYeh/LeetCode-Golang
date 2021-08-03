package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	dict := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := []byte{}
	isLeadZero := true

	for i := 7; i >= 0; i-- {
		digit := (num >> (4 * i)) & 0xf
		if !isLeadZero || digit != 0 {
			isLeadZero = false
			res = append(res, dict[(num>>(4*i))&0xf])
		}
	}
	return string(res)

}

func main() {
	toHex(26)
}

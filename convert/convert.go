package convert

import "fmt"

// String Convert
func ToString(value interface{}) string {
	point := []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"}
	if isValueInList(fmt.Sprintf("%T", value), point) {
		return fmt.Sprintf("%d", value)
	}

	if fmt.Sprintf("%T", value) == "float32" || fmt.Sprintf("%T", value) == "float64" {
		return fmt.Sprintf("%.2f", value)
	}
	return ""
}

// Check Slice Element
func isValueInList(value string, s []string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func Num2alpha(number int) (char string) {
	var (
		k    int
		temp []int
	)
	//用来匹配的字符A-Z
	charSlice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if number <= 26 {
		return charSlice[number]
	}

	for {
		//从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
		k = number % 26
		if k == 0 {
			temp = append(temp, 26)
			k = 26
		} else {
			temp = append(temp, k)
		}
		//减去Num最后一位数的值，因为已经记录在temp中
		number = (number - k) / 26
		//小于等于26直接进行匹配，不需要进行数据拆分
		if number <= 26 {
			temp = append(temp, number)
			break
		}
	}

	for _, v := range temp {
		char = charSlice[v] + char
	}

	return
}

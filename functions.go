package util

//计算字符串长度，1个汉字相当2个英文字符/UTF-8
func Strlen(str string) int{
	n := len(str)
	c := 0
	for i:=0; i<n; i++ {
		if str[i] < 128 {
			c ++
		} else if str[i] > 252 { //6 bytes
			c += 2
			i += 5
		} else if str[i] > 248 { //5 bytes
			c += 2
			i += 4
		} else if str[i] > 240 { //4 bytes
			c += 2
			i += 3
		} else if str[i] > 224 { //3 bytes
			c += 2
			i += 2
		} else if str[i] > 192 { //2 bytes
			c += 2
			i += 1
		}
	}

	return c
}

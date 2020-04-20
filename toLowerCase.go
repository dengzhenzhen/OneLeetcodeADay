package main

func toLowerCase(str string) string {
	str_bytes := []byte(str)
	for i, v := range str_bytes {
		if byte('A') <= v && v <= byte('Z') {
			str_bytes[i] += 32
		}
	}
	return string(str_bytes)
}

package go_func

import (
	"strconv"
	"strings"
)

// implode 函数
func Implode(fenge string, str_array []string) string {
	var data_string = ""
	for k, element := range str_array {
		if k != 1 {
			data_string += fenge
		}
		data_string += element
	}
	return data_string
}

// in_array函数
func In_array(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

//explode 函数
func Explode(target string, src string) []string {
	arr := strings.Split(target, src)
	return arr
}

//strtofloat 函数
func Strtofloat(src string) float64 {
	f, _ := strconv.ParseFloat(src, 64) //string转float64
	src_data := float64(f)
	return src_data
}
func Floattostr(src float64) string {
	src_data := strconv.FormatFloat(src, 'G', -1, 64) // float64s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	return src_data
}

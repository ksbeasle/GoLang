package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

func main() {
	/* ----------------------------3.10----------------------------*/
	fmt.Println("--------------")
	str1 := "12345"
	str2 := "123456789"
	strWithCommas1 := comma(str1)
	strWithCommas2 := comma(str2)
	fmt.Println(strWithCommas1)
	fmt.Println(strWithCommas2)
	/* ----------------------------3.10----------------------------*/

	/* ----------------------------3.12----------------------------*/
	fmt.Println("--------------")
	str1 = "margana"
	str2 = "anagram"
	str3 := "aanagramz"
	fmt.Println(isAnagram(str1, str2))
	fmt.Println(isAnagram(str2, str3))
	/* ----------------------------3.12----------------------------*/

}

//Exercise 3.10 non-recursive version of comma() and using bytes.Buffer
func comma(s string) string {
	var buf bytes.Buffer
	for i, count := len(s)-1, 0; i >= 0; i-- {
		if count == 3 {
			buf.WriteString(",")
			count = 0
		}

		buf.WriteByte(s[i])
		count++

	}
	return reverse(buf.String())

}

//Helper function to reverse a string using bytes.Buffer
func reverse(s string) string {
	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}
	return buf.String()
}

//Exercise 3.12 Anagram
func isAnagram(s1, s2 string) bool {
	// str1 := strings.Split(s1, "")
	// sort.Strings(str1)
	// str2 := strings.Split(s2, "")
	// sort.Strings(str2)
	// strr := strings.Join(str1, "")
	// strr1 := strings.Join(str2, "")
	// return strr1 == strr
	str1 := []byte(s1)
	sort.Slice(str1, func(i int, j int) bool { return str1[i] < str1[j] })

	str2 := []byte(s2)
	sort.Slice(str2, func(i int, j int) bool { return str2[i] < str2[j] })

	return reflect.DeepEqual(str1, str2)
}

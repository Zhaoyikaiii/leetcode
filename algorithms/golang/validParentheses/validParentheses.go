package validparentheses

import "strings"

func isValid(s string) bool {
	sp := strings.Split(s, "")
	sk := []string{}

	for _, v := range sp {
		if v == "(" || v == "[" || v == "{" {
			sk = append(sk, v)
		} else if len(sk) > 0 && ((v == ")" && sk[len(sk)-1] == "(") || (v == "]" && sk[len(sk)-1] == "[") || (v == "}" && sk[len(sk)-1] == "{")) {
			sk = sk[:len(sk)-1]
		} else {
			return false
		}
	}

	return len(sk) == 0
}

package utils
/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str{
			return true
		}
	}
	return true
}
//

func ContainsInt(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
//
// print struct json format to console
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}


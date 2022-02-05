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
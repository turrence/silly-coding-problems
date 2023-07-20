package leetcode

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.

*/

func romanStringToNum(s string) int {
	romanTable := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	val, sLen := 0, len(s)
	i := 0

	// this checks characters in of the string in pairs
	// if 2 characters don't exist in the map, then it is a standalone
	// character and then we just add that single value
	for i = 0; i < sLen-1; {
		// second value is exclusive index
		if num, ok := romanTable[s[i:i+2]]; ok {
			val += num
			i += 2
		} else {
			val += romanTable[s[i:i+1]]
			i += 1
		}
	}

	// there is an edgecase in which we don't get to add the last
	// character in the string due to our bounds
	if i == sLen-1 {
		val += romanTable[s[i:i+1]]
	}

	return val
}

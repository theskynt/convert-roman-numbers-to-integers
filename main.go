package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
	}

	n := len(s)
	total := romanMap[s[n-1]]

	for i := n - 2 ; i >= 0 ; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		}else {
			total += romanMap[s[i]]
		}
	}

	return total

}
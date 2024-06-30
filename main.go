package main

func romanToInt(s string) int {
<<<<<<< HEAD
	romanMap := map[byte]int{
=======
	romanMap := map[byte] int {
>>>>>>> 297b62f ([3] should return 1 when input is I)
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
	}

	n := len(s)
	total := romanMap[s[n-1]]

<<<<<<< HEAD
	for i := n-2 ; i >= 0 ; i-- {
=======
	for i := n - 2 ; i >= 0 ; i-- {
>>>>>>> 297b62f ([3] should return 1 when input is I)
		if romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			total += romanMap[s[i]]
		}
	}

	return total
}
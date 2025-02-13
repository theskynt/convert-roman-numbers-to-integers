package main

func romanValue(b byte)int{
	if b == 'I'{
		return 1
	}
	if b == 'V' {
		return 5
	}
	if b == 'X' {
		return 10
	}
	if b == 'L' {
		return 50
	}
	if b == 'C' {
		return 100
	}
	return 0
}

func romanToInt(s string)int{
	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		value := romanValue(s[i])

		if i < n-1 && value < romanValue(s[i+1]){
			total -= value
		}else{
			total += value
		}
	}
	return total
}
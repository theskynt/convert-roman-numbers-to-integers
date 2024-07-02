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
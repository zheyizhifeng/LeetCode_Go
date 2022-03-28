// @algorithm @lc id=12 lang=golang 
// @title integer-to-roman


package main
import ("strings")
// @test(3)="III"
// @test(58)="LVIII"
// @test(1994)="MCMXCIV"
func getRomanChar(n int) string {
	switch (n) {
		case 1: return "I";
		case 4: return "IV";
		case 5: return "V";
		case 9: return "IX";
		case 10: return "X";
		case 40: return "XL";
		case 50: return "L";
		case 90: return "XC";
		case 100: return "C";
		case 400: return "CD";
		case 500: return "D";
		case 900: return "CM";
		case 1000: return "M";
		default: return "";
	}
}
func intToRoman(num int) string {
	ans := ""
	tenTimes := 1
	for num > 0 {
		bit := num % 10
		if bit == 4 || bit == 9 {
			ans = getRomanChar(bit * tenTimes) + ans
		} else {
			fives, ones := bit / 5, bit % 5
			ans = getRomanChar(fives * tenTimes * 5) + strings.Repeat(getRomanChar(tenTimes), ones) + ans
		}
		num /= 10
		tenTimes *= 10;
	}
	return ans
}
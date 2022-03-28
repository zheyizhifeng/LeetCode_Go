// @algorithm @lc id=13 lang=golang 
// @title roman-to-integer

package main
// @test("III")=3
// @test("LVIII")=58
// @test("MCMXCIV")=1994
func getRomanInt(c rune) int {
	switch c {
		case 'I': return 1;
		case 'V': return 5;
		case 'X': return 10;
		case 'L': return 50;
		case 'C': return 100;
		case 'D': return 500;
		case 'M': return 1000;
		default: return 0;
	}
}

func romanToInt(s string) int {
	ans := 0;
	len := len(s);
	prevInt := 0;
	for _, cur := range s {
		curInt := getRomanInt(cur)
		if prevInt < curInt {
			ans -= prevInt;
		} else {
			ans += prevInt;
		}
		prevInt = curInt;
	}
	return ans + getRomanInt(rune(s[len - 1]));
}
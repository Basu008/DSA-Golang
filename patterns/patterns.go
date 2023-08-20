package patterns

import (
	"fmt"
	"math"
)

func Pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func Pattern2(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func Pattern3(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func Pattern4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}

func Pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func Pattern6(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func Pattern7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < (2*n)-1; j++ {
			if j < n {
				if i+j >= n-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			} else {
				if j-i <= n-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
}

func Pattern8(n int) {
	limit := (2 * n) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < limit; j++ {
			if j < n {
				if j-i < 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			} else {
				if i+j >= limit {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
		}
		fmt.Println("")
	}
}

func Pattern9(n int) {
	Pattern7(n)
	Pattern8(n)
}

func Pattern10(n int) {
	limit := (2 * n) - 1
	for i := 0; i < limit; i++ {
		for j := 0; j < n; j++ {
			if i >= n {
				if i+j < limit {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			} else {
				if i-j >= 0 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
}

func Pattern11(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println("")
	}
}

func Pattern12(n int) {
	for i := 1; i <= n; i++ {
		end := 0
		spaces := (n - i) * 2
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			end = j
		}
		for k := 1; k <= spaces; k++ {
			fmt.Print(" ")
		}
		for j := end; j > 0; j-- {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func Pattern13(n int) {
	output := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(output)
			fmt.Print(" ")
			output++
		}
		fmt.Println("")
	}
}

func Pattern14(n int) {
	for i := 1; i <= n; i++ {
		for j := 65; j-65 <= i; j++ {
			fmt.Print(string(rune(j)))
		}
		fmt.Println("")
	}
}

func Pattern15(n int) {
	for i := 0; i <= n; i++ {
		letters := n - i
		alphabet := 'A'
		for j := 1; j <= letters; j++ {
			fmt.Print(string(alphabet))
			alphabet++
		}
		fmt.Println("")
	}
}

func Pattern16(n int) {
	alphabet := 'A'
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string(alphabet))
		}
		alphabet++
		fmt.Println("")
	}
}

func Pattern17(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		ch := 'A'
		letters := (2 * i) - 1
		midPoint := (letters + 1) / 2
		for k := 1; k <= letters; k++ {
			fmt.Print(string(ch))
			if k < midPoint {
				ch++
			} else {
				ch--
			}
		}
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func Pattern18(n int) {
	target := 'A' + n
	for i := 1; i <= n; i++ {
		for j := target - i; j < target; j++ {
			fmt.Print(string(rune(j)))
		}
		fmt.Print("\n")
	}
}

func Pattern19(n int) {
	for i := 0; i < n*2; i++ {
		var midPoint = ((n * 2) / 2) - 1
		var spaces int
		var stars int
		if i <= midPoint {
			spaces = i * 2
		} else {
			spaces = (n * 2) - ((i - midPoint) * 2)
		}
		stars = n*2 - spaces
		for j := 1; j <= stars/2; j++ {
			fmt.Print("*")
		}
		for k := 1; k <= spaces; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= stars/2; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func Pattern20(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		spaces := ((n * 2) - (i * 2))
		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	spaces := 2
	for i := 1; i < n; i++ {
		stars := n*2 - spaces
		for j := 1; j <= stars/2; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= stars/2; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
		spaces += 2
	}
}

func Pattern21(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == n || j == 1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func Pattern22(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			top := j
			left := i
			right := 2*n - 2 - j
			bottom := 2*n - 2 - i
			value := float64(n) - math.Min(math.Min(float64(left), float64(right)), math.Min(float64(top), float64(bottom)))
			fmt.Print(value)
		}
		fmt.Println("")
	}
}

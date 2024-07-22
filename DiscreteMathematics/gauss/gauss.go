package main

import (
	"fmt"
	"os"
)

type Fraction struct {
	numerator   int
	denominator int
}

func abs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func gcd(num1, num2 int) int {
	num1, num2 = abs(num1), abs(num2)
	for num2 != 0 {
		t := num2
		num2 = num1 % num2
		num1 = t
	}

	return num1
}

func have_sollution(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		all_zero := true

		for j := 0; j < len(matrix[0])-1; j++ {
			if matrix[i][i] != 0 {
				all_zero = false
			}
		}

		if all_zero && matrix[i][len(matrix[0])-1] != 0 {
			return false
		}
	}

	return true
}

func multiply_number_with_frac(num int, frac Fraction) Fraction {
	return Fraction{frac.numerator * num, frac.denominator}
}

func add_fractoins(frac1, frac2 Fraction) Fraction {
	new_numerator := (frac1.numerator * frac2.denominator) + (frac1.denominator * frac2.numerator)
	new_denominator := frac1.denominator * frac2.denominator
	return Fraction{new_numerator, new_denominator}
}

func substract_number_with_fractoins(num int, frac Fraction) Fraction {
	return Fraction{(num * frac.denominator) - frac.numerator, frac.denominator}
}

func division_number_with_fractions(num int, frac Fraction) Fraction {
	return Fraction{frac.numerator, frac.denominator * num}
}

func multiply_lines(i_bas, j_bas int, matrix *[][]int) {
	i := i_bas + 1
	elem_base := (*matrix)[i_bas][j_bas]

	for i < len(*matrix) {
		j := 0
		for j < len((*matrix)[0]) {
			(*matrix)[i][j] *= elem_base
			j++
		}
		i++
	}
}

func subtract_lines(i_bas, j_bas int, matrix *[][]int) {
	i := i_bas + 1

	for i < len(*matrix) {
		j := j_bas
		counter := (*matrix)[i][j] / (*matrix)[i_bas][j]

		for j < len((*matrix)[0]) {
			(*matrix)[i][j] -= (*matrix)[i_bas][j] * counter
			j++
		}
		i++
	}
}

func solve_matrix(n int, matrix [][]int, ans *[]Fraction) {
	i := n - 1
	for i >= 0 {
		sum := Fraction{0, 1}
		j := 0
		for j < n {

			if i != j {
				new_frac := multiply_number_with_frac(matrix[i][j], (*ans)[j])
				sum = add_fractoins(sum, new_frac)
			}
			j++
		}

		ans_frac := substract_number_with_fractoins(matrix[i][n], sum)
		ans_frac = division_number_with_fractions(matrix[i][i], ans_frac)

		del := gcd(ans_frac.numerator, ans_frac.denominator)
		if ans_frac.numerator < 0 && ans_frac.denominator < 0 {
			ans_frac = Fraction{abs(ans_frac.numerator / del), abs(ans_frac.denominator / del)}

		} else if ans_frac.numerator < 0 || ans_frac.denominator < 0 {
			ans_frac = Fraction{-abs(ans_frac.numerator / del), abs(ans_frac.denominator / del)}
		} else {
			ans_frac = Fraction{ans_frac.numerator / del, ans_frac.denominator / del}
		}

		(*ans)[i] = ans_frac
		i--
	}
}

func find_no_zero_line(i_base int, matrix *[][]int) bool {
	if (*matrix)[i_base][i_base] == 0 {

		for i := i_base + 1; i < len(*matrix); i++ {
			if (*matrix)[i][i_base] != 0 {
				(*matrix)[i_base], (*matrix)[i] = (*matrix)[i], (*matrix)[i_base]
				return true
			}
		}
		return false
	}
	return true
}

func main() {
	var n, num int

	fmt.Scan(&n)
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < (n + 1); j++ {
			fmt.Scan(&num)
			matrix[i] = append(matrix[i], num)
		}
	}

	for i := 0; i < n; i++ {

		if find_no_zero_line(i, &matrix) {
			multiply_lines(i, i, &matrix)
			subtract_lines(i, i, &matrix)
		} else {
			fmt.Print("No solution")
			os.Exit(0)
		}
	}

	var ans []Fraction
	for i := 0; i < n; i++ {
		ans = append(ans, Fraction{0, 1})
	}

	if have_sollution(matrix) {
		solve_matrix(n, matrix, &ans)

		for _, elem := range ans {
			fmt.Printf("%d/%d\n", elem.numerator, elem.denominator)
		}
	} else {
		fmt.Print("No solution")
	}
}
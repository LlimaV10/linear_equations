package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"os"
)

func print_matrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%-10.4f", m[i][j])
		}
		fmt.Printf("\n")
	}
}

func get_matrix_minor(m [][]float64, i int, j int) [][]float64 {
	m1 := make([]([]float64), len(m) - 1)
	for z := 0; z < len(m) - 1; z++ {
		m1[z] = make([]float64, len(m) - 1)
	}
	for y := 0; y < i; y++ {
		for x := 0; x < j; x++ {
			m1[y][x] = m[y][x]
		}
		for x := j + 1; x < len(m); x++ {
			m1[y][x - 1] = m[y][x]
		}
	}
	for y := i + 1; y < len(m); y++ {
		for x := 0; x < j; x++ {
			m1[y - 1][x] = m[y][x]
		}
		for x := j + 1; x < len(m); x++ {
			m1[y - 1][x - 1] = m[y][x]
		}
	}
	return m1
}

func get_matrix_determinant(m [][]float64) float64 {
	if len(m) == 1 {
		return m[0][0]
	}
	if len(m) == 2 {
		return m[0][0] * m[1][1] - m[0][1] * m[1][0]
	}
	sum := 0.0
	for i := 0; i < len(m[0]); i++ {
		m2 := get_matrix_minor(m, 0, i)
		sum += math.Pow(float64(-1), float64(i)) * m[0][i] * get_matrix_determinant(m2)
	}
	return sum
}

func matrix_transpond(m [][]float64) [][]float64 {
	m1 := make([]([]float64), len(m))
	for z := 0; z < len(m); z++ {
		m1[z] = make([]float64, len(m))
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			m1[i][j] = m[j][i]
		}
	}
	return m1
}

func get_reverse_matrix(m [][]float64) [][]float64 {
	m1 := make([]([]float64), len(m))
	for z := 0; z < len(m); z++ {
		m1[z] = make([]float64, len(m))
	}
	det := get_matrix_determinant(m)
	fmt.Printf("Детерминант |A| = %.4f\n", det)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			mm := get_matrix_minor(m, i, j);
			m1[i][j] = math.Pow(float64(-1), float64(i + j)) * get_matrix_determinant(mm) / det
		}
	}
	fmt.Println("Матрица алгебраических дополнений разделённая на |A|:")
	print_matrix(m1)
	m1 = matrix_transpond(m1)
	fmt.Println("Транспонированная матрица алгебраических дополнений разделённая на |A|")
	fmt.Println("то бишь обратная матрица A^-1:")
	print_matrix(m1)
	return m1
}

func print_equation(m [][]float64, r []float64) {
	for i := 0; i < len(m); i++ {
		fmt.Printf("%.2f * x%d", m[i][0], 1)
		for j := 1; j < len(m[i]); j++ {
			if (m[i][j] > 0) {
				fmt.Printf(" + %.2f * x%d", m[i][j], j + 1)
			} else {
				fmt.Printf(" - %.2f * x%d", -m[i][j], j + 1)
			}
		}
		fmt.Printf(" = %.2f\n", r[i])
	}
}

func get_mult(m [][]float64, r []float64) []float64 {
	x := make([]float64, len(r))
	for i := 0; i < len(m); i++ {
		x[i] = 0
		for j := 0; j < len(m[i]); j++ {
			x[i] += m[i][j] * r[j]
		}
	}
	return x
}

func main() {
	file, err := os.Open("matrix.txt")
	if err != nil{
		fmt.Println(err) 
		os.Exit(1) 
	}
	defer file.Close()
	data := make([]byte, 512)
	n, err := file.Read(data)
	spl := strings.Split(string(data[:n]), "\r\n")
	mlen := len(strings.Split(spl[0], " "))
	m := make([][]float64, mlen)
	for i := 0; i < mlen; i++ {
		sm := strings.Split(spl[i], " ")
		m[i] = make([]float64, mlen)
		for j := 0; j < len(sm); j++ {
			m[i][j], err = strconv.ParseFloat(sm[j], 64)
		}
	}
	r := make([]float64, mlen)
	for j := 0; j < mlen; j++ {
		r[j], err = strconv.ParseFloat(spl[mlen + j], 64)
	}
	fmt.Println("Уравнение:")
	print_equation(m, r)
	fmt.Println("Матрица 'A':")
	print_matrix(m)
	reverse := get_reverse_matrix(m)
	res := get_mult(reverse, r)
	fmt.Println("Ответ:")
	fmt.Printf(" X = %10.4f\n", res[0])
	for j := 1; j < len(res); j++ {
		fmt.Printf("%15.4f\n", res[j])
	}
}

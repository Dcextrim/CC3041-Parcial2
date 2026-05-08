package main

import "fmt"

// Teclado:
//   [ 1 ][ 2 ][ 3 ]
//   [ 4 ][ 5 ][ 6 ]
//   [ 7 ][ 8 ][ 9 ]
//   [ * ][ 0 ][ # ]

var adj = map[int][]int{
	0: {0, 8},
	1: {1, 2, 4},
	2: {1, 2, 3, 5},
	3: {2, 3, 6},
	4: {1, 4, 5, 7},
	5: {2, 4, 5, 6, 8},
	6: {3, 5, 6, 9},
	7: {4, 7, 8},
	8: {5, 7, 8, 9, 0},
	9: {6, 8, 9},
}

func contarCombinaciones(n int) int {
	if n <= 0 {
		return 0
	}

	// Caso base: longitud 1
	dp := make([]int, 10)
	for d := 0; d <= 9; d++ {
		dp[d] = 1
	}

	for paso := 2; paso <= n; paso++ {
		nuevoDp := make([]int, 10)
		for d := 0; d <= 9; d++ {
			for _, vecino := range adj[d] {
				nuevoDp[vecino] += dp[d]
			}
		}
		dp = nuevoDp
	}

	total := 0
	for d := 0; d <= 9; d++ {
		total += dp[d]
	}
	return total
}

// contarDetallado muestra la tabla DP paso a paso
func contarDetallado(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, 10)
	for d := 0; d <= 9; d++ {
		dp[d] = 1
	}
	fmt.Printf("  Paso 1 (base): dig[0..9]=%v  total=%d\n", dp, sumar(dp))

	for paso := 2; paso <= n; paso++ {
		nuevoDp := make([]int, 10)
		for d := 0; d <= 9; d++ {
			for _, vecino := range adj[d] {
				nuevoDp[vecino] += dp[d]
			}
		}
		dp = nuevoDp
		fmt.Printf("  Paso %d:        dig[0..9]=%v  total=%d\n", paso, dp, sumar(dp))
	}

	total := sumar(dp)
	return total
}

func sumar(dp []int) int {
	s := 0
	for _, v := range dp {
		s += v
	}
	return s
}

func main() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║  Nokia 3230 - Combinaciones DP (con self)   ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Teclado:")
	fmt.Println("  [ 1 ][ 2 ][ 3 ]")
	fmt.Println("  [ 4 ][ 5 ][ 6 ]")
	fmt.Println("  [ 7 ][ 8 ][ 9 ]")
	fmt.Println("  [ * ][ 0 ][ # ]")
	fmt.Println()
	fmt.Println("Reglas:")
	fmt.Println("  - Puede iniciar en cualquier dígito 0-9")
	fmt.Println("  - Siguiente tecla: arriba/abajo/izq/der o MISMA tecla")
	fmt.Println("  - No puede pisar * ni #")
	fmt.Println()

	casos := []int{1, 2, 3, 4, 5}

	for _, n := range casos {
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("Caso n = %d\n", n)
		resultado := contarDetallado(n)
		fmt.Printf("  ➜  Total combinaciones válidas: %d\n\n", resultado)
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("Verificación del ejemplo del examen:")
	r := contarCombinaciones(2)
	check := "NO"
	if r == 36 {
		check = "SI"
	}
	fmt.Printf("  n=2 esperado=36, obtenido=%d  %s\n", r, check)
}

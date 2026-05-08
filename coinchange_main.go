package main

import "fmt"

// denominaciones disponibles ordenadas de mayor a menor (greedy-choice)
var denominaciones = []int{25, 10, 5, 1}

// Result representa cuántas monedas se usaron de cada denominación
type Result struct {
	denominacion int
	cantidad     int
}

// hacerSencillo resuelve el problema con greedy.
func hacerSencillo(m int) (int, []Result) {
	restante := m
	total := 0
	var results []Result

	for _, c := range denominaciones {
		if restante == 0 {
			break
		}

		cantidad := restante / c // mayor cantidad posible de esta denominación
		restante -= cantidad * c
		total += cantidad

		if cantidad > 0 {
			results = append(results, Result{c, cantidad})
		}
	}

	return total, results
}

// formatDenom convierte centavos a string legible: 25→"Q0.25", 1→"Q0.01"
func formatDenom(c int) string {
	return fmt.Sprintf("Q0.%02d", c)
}

func printCase(label string, m int) {
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Caso: %s\n", label)
	fmt.Printf("Monto = %d centavos (Q%.2f)\n", m, float64(m)/100)
	fmt.Println()

	total, results := hacerSencillo(m)

	fmt.Printf("  %-14s %s\n", "Denominación", "Cantidad")
	for _, r := range results {
		fmt.Printf("  %-14s %d\n", formatDenom(r.denominacion), r.cantidad)
	}

	fmt.Printf("\n  ➜  Total de monedas: %d\n\n", total)
}

func main() {
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║     Hacer Sencillo — Algoritmo Greedy           ║")
	fmt.Println("║     Denominaciones: {1, 5, 10, 25} centavos     ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Println()

	// ── Caso 1 ───────────────────────────
	// Q2.93 = 293 centavos → esperado: 11×Q0.25 + 1×Q0.10 + 1×Q0.05 + 3×Q0.01
	printCase("Ejemplo del examen (Q2.93)", 293)

	// ── Caso 2: monto redondo ────────────────────────────────
	// Q1.00 = 100 centavos → 4×Q0.25
	printCase("Monto redondo (Q1.00)", 100)

	// ── Caso 3: monto que fuerza todas las denominaciones ────
	// Q0.41 = 41 centavos → 1×Q0.25 + 1×Q0.10 + 1×Q0.05 + 1×Q0.01
	printCase("Todas las denominaciones (Q0.41)", 41)

	// ── Caso 4: monto grande ─────────────────────────────────
	// Q9.99 = 999 centavos
	printCase("Monto grande (Q9.99)", 999)

	// ── Caso 5: monto mínimo ─────────────────────────────────
	// Q0.01 = 1 centavo → 1×Q0.01
	printCase("Monto mínimo (Q0.01)", 1)
}

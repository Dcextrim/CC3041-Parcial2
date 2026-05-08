package main

import (
	"fmt"
	"sort"
)

// Item representa un artículo disponible para robar
type Item struct {
	id    int
	price float64 // precio total del artículo
	units float64 // unidades disponibles
	ratio float64 // valor por unidad (price / units)
}

// Result representa cuánto se tomó de cada artículo
type Result struct {
	item  Item
	taken float64
	value float64
}

// fractionalKnapsack resuelve el problema con greedy.
func fractionalKnapsack(items []Item, W float64) (float64, []Result) {
	// Calcular ratio de cada artículo
	for i := range items {
		items[i].ratio = items[i].price / items[i].units
	}

	// Ordenar por ratio descendente (greedy-choice)
	sort.Slice(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	totalValue := 0.0
	remaining := W
	var results []Result

	for _, item := range items {
		if remaining == 0 {
			break
		}

		// Tomar la mayor cantidad posible de este artículo
		taken := min(item.units, remaining)
		value := taken * item.ratio

		totalValue += value
		remaining -= taken

		results = append(results, Result{item, taken, value})
	}

	return totalValue, results
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func printCase(label string, items []Item, W float64) {
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Caso: %s  |  Capacidad W = %.0f\n", label, W)
	fmt.Printf("%-10s %-10s %-10s %-10s\n", "Artículo", "Precio", "Unidades", "Ratio p/w")
	for _, it := range items {
		fmt.Printf("%-10d %-10.2f %-10.2f %-10.2f\n", it.id, it.price, it.units, it.price/it.units)
	}
	fmt.Println()

	// Hacemos copia para no mutar el slice original
	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)

	totalValue, results := fractionalKnapsack(itemsCopy, W)

	fmt.Printf("%-10s %-12s %-12s %-12s\n", "Artículo", "Ratio p/w", "Tomado", "Valor")
	for _, r := range results {
		fmt.Printf("%-10d %-12.2f %-12.2f %-12.2f\n",
			r.item.id, r.item.ratio, r.taken, r.value)
	}
	fmt.Printf("\n  ➜  Valor total máximo: $%.2f\n\n", totalValue)
}

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Println("║       Fractional Knapsack — Algoritmo Greedy        ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
	fmt.Println()

	// ── Caso 1: ──────────────────────────
	// Item 1: $60 / 10u = $6/u
	// Item 2: $100 / 20u = $5/u
	// Item 3: $120 / 30u = $4/u
	// W = 50 → esperado $240
	printCase("Ejemplo del examen", []Item{
		{id: 1, price: 60, units: 10},
		{id: 2, price: 100, units: 20},
		{id: 3, price: 120, units: 30},
	}, 50)

	// ── Caso 2: capacidad exacta ─────────────────────────────
	// La bolsa alcanza exactamente para todos los artículos
	printCase("Capacidad exacta (toma todo)", []Item{
		{id: 1, price: 500, units: 5},
		{id: 2, price: 200, units: 10},
		{id: 3, price: 300, units: 15},
	}, 30)

	// ── Caso 3: capacidad muy pequeña ───────────────────────
	// La bolsa solo alcanza para fracciones de los artículos
	printCase("Capacidad pequeña (solo fracciones)", []Item{
		{id: 1, price: 100, units: 50},
		{id: 2, price: 280, units: 40},
		{id: 3, price: 120, units: 30},
	}, 10)

	// ── Caso 4: un solo artículo pero muy valioso ────────────
	printCase("Un artículo muy valioso, W menor a disponibilidad", []Item{
		{id: 1, price: 1000, units: 100},
		{id: 2, price: 50, units: 5},
	}, 20)
}

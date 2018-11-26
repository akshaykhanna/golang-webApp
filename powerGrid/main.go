package main

import "fmt"

func main() {
	fmt.Println("Power Grid console application")

	powerPlantsCapacity := []int{20, 70, 86, 50}
	utilization := 0.75

	totalUtil := 0.
	for _, powerPlantCapacity := range powerPlantsCapacity {
		totalUtil += utilization * float64(powerPlantCapacity)
	}

	fmt.Println("powerPlantsCapacity :", powerPlantsCapacity)
	fmt.Println("utilization :", utilization)
	fmt.Println("totalUtil :", totalUtil)
}

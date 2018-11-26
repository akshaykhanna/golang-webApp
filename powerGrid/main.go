package main

import "fmt"

func main() {
	fmt.Println("Power Grid console application")

	powerPlantsCapacity := []int{20, 70, 86, 50}
	activePowerPlants := []int{0, 2}
	gridLoad := 75.

	var opiton int
	fmt.Println("Enter your choice !")
	fmt.Println("1 Power plant report")
	fmt.Println("2 Power grid report")

	fmt.Scan(&opiton)
	fmt.Println(opiton)

	switch {
	case opiton == 1:
		for index, value := range powerPlantsCapacity {
			fmt.Printf("Plant %d capacity is %d. \n", index, value)
		}
		break
	case opiton == 2:
		effectivePowerOutput := 0.
		for _, activePowerPlantIndex := range activePowerPlants {
			effectivePowerOutput += float64(powerPlantsCapacity[activePowerPlantIndex])
		}
		fmt.Printf("effectivePowerOutput: %f \n", effectivePowerOutput)
		fmt.Println("Utilization : ", (gridLoad/effectivePowerOutput)*100)
	}

}

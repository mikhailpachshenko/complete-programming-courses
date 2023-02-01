package main

import "fmt"

func main() {
	var (
		laps          = 4
		speed         = 358
		engine        = 254
		wheels        = 93
		steeringWheel = 49
		wind          = 21
		rain          = 17
	)

	fmt.Println("——————————")
	fmt.Println("Супер гонки. Круг", laps)
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("——————————")
	fmt.Println("Водитель: Шумахер")
	fmt.Println("Скорость:", speed)
	fmt.Println("——————————")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("——————————")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}

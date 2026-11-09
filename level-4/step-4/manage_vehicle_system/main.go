package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	result := make(map[string]float64)

	for _, vehicle := range vehicles {
		vehicleType := reflect.TypeOf(vehicle).String()
		travelTime := vehicle.CalculateTravelTime(distance)

		result[vehicleType] = travelTime
	}

	return result
}

func main() {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)

	fmt.Printf("Ожидается время для автомобиля %.2f часа\n", travelTimes["main.Car"])
	fmt.Printf("Ожидается время для мотоцикла %.2f часа", travelTimes["main.Motorcycle"])
}

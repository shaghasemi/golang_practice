package main

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	return &Car{speed: speed, battery: battery}
}

func GetSpeed(car *Car) int {
	return car.speed
}

func GetBattery(car *Car) int {
	return car.battery
}

func ChargeCar(car *Car, minutes int) {
	if minutes <= 0 {
		return
	}

	if car.battery >= 100 {
		return
	}

	var chargeAdded int
	chargeAdded = minutes / 2
	car.battery = car.battery + chargeAdded

	if car.battery > 100 {
		car.battery = 100
	}
}

func TryFinish(car *Car, distance int) string {
	if distance <= 0 {
		return "0.00"
	}

	chargeDrained := distance / 2

	if chargeDrained > car.battery {
		car.battery = 0
		return ""
	}
	car.battery -= chargeDrained

	timeElapsed := float64(distance) / float64(car.speed)

	return fmt.Sprintf("%.2f", timeElapsed)
}

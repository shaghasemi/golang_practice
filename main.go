package main

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
	if car.battery >= 100 {
		return
	}
	var chargeAdded int
	chargeAdded = minutes % 2
	car.battery = car.battery + chargeAdded
	if car.battery > 100 {
		car.battery = 100
	}
}

func TryFinish(car *Car, distance int) string {
	if distance <= 0 {
		return "0.00"
	}

	var chargeDrained int
	chargeDrained = distance % 2

	car.battery = car.battery - chargeDrained

	if car.battery < 0 {
		car.battery = 0
		return ""
	}

	var timeElapsed float32
	timeElapsed = distance / car.speed

	return float32(timeElapsed)
}

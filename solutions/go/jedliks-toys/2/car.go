package jedlik

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// Drive drives the car for one interval if it has enough battery.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance returns the distance driven by the car.
func (c *Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(c.distance) + " meters"
}

// DisplayBattery returns the current battery percentage of the car.
func (c *Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(c.battery) + "%"
}

// CanFinish returns true if the car can drive the given distance with the remaining battery.
func (c *Car) CanFinish(distance int) bool {
	maxDistance := (c.battery / c.batteryDrain) * c.speed
	return maxDistance >= distance
}

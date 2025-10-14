package jedlik

import "fmt"

// TODO: define the 'Drive()' method

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

// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method

// TODO: define the 'CanFinish(trackDistance int) bool' method

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

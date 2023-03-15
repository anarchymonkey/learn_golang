package main

import (
	"errors"
	"fmt"
	"math"
)

// as we all know that go does not have any logic of classes embedded in it
// but we can mimic the whole operation by applying the method to a type

type CarParts struct {
	doors         int
	gears         int
	currentGear   int
	color         string
	engineVersion string
}

func (c CarParts) getColorOfCar() string {
	return c.color
}

func (c CarParts) getNoOfDoorsInTheCar() int {
	return c.doors
}

func (c CarParts) getNoOfGearsInTheCar() int {
	return c.gears
}

func (c CarParts) getCurrentGear() int {
	return c.currentGear
}

func (c CarParts) changeGear(gearNo int) (int, error) {

	if gearNo > c.gears {
		return 0, errors.New("the gear to change cannot be more than the no of gears in the car, thus stopping the car")
	}
	return gearNo, nil
}

// a reciever as a pointer can directly change the values of the underlying structure
func (c *CarParts) addEngineVersion(engineVersion string) {
	c.engineVersion = engineVersion
}

type CustomFloat float64

// custom types can also have methods
func (f CustomFloat) getPowerOfFloats(exponent int) float64 {
	return float64(math.Pow(float64(f), float64(exponent)))
}

func main() {
	fmt.Println("Lets learn about methods")

	car := CarParts{
		doors:       4,
		gears:       6,
		color:       "Cherry Red",
		currentGear: 0,
	}

	fmt.Printf("The color of the car is %s \n", car.getColorOfCar())
	fmt.Printf("The no of doors in the car is %d \n", car.getNoOfDoorsInTheCar())
	fmt.Printf("The no of gears in the car is %d \n", car.getNoOfGearsInTheCar())

	res, error := car.changeGear(3)

	if error != nil {
		defer func() {
			fmt.Println("The deffered errors are", error)
		}()
	}

	car.currentGear = res

	fmt.Println("The current gear after change is", car.getCurrentGear())

	car.addEngineVersion("IvtechEngine")

	car.addEngineVersion("IVTECH M2")

	fmt.Println("The engine version is", car.engineVersion)

	var f CustomFloat = CustomFloat(1.234)

	fmt.Printf("The floating point value after powering with 2 is %v \n", f.getPowerOfFloats(2))
}

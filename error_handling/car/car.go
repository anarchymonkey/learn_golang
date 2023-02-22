package car

import (
	"errors"
	"fmt"
)

func StartCar(carName string) (string, error) {
	if carName == "" {
		return "", errors.New("there has been some error! Cannot send empty car name")
	}
	message := fmt.Sprintf("Starting car, name of the car is %v", carName)

	return message, nil
}

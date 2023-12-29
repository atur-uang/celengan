package services

type Vehicle struct {
	wheel int
}

func (vehicle Vehicle) GetWheel() int {
	return vehicle.wheel
}

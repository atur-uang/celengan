package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVehicle_GetWheel(t *testing.T) {
	vehicle := Vehicle{
		wheel: 89,
	}
	assert.Equal(t, 89, vehicle.GetWheel())

}

func TestVehicle_CustomWheel(t *testing.T) {
	vehicle := Vehicle{
		wheel: 2,
	}
	assert.Equal(t, 2, vehicle.GetWheel())

}

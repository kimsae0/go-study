package main

import (
	"github.com/kimsae0/go-study/interface/school"
	// "github.com/kimsae0/go-study/interface/work"
)

type Rider interface {
	Ride(vehicle string)
}

// Sender 인터페이스를 입력으로 받는다
func RideVehicle(name string, rider Rider) {
	rider.Ride(name)
}

func main() {

	schoolRider := &school.SchoolRider{}
	RideVehicle("버스", schoolRider)
	RideVehicle("지하철", schoolRider)

	// workRider := &work.WorkRider{}
	// RideVehicle("버스", workRider)
	// RideVehicle("지하철", workRider)
}
package home

import "fmt"

type HomeRider struct {
}

func (s *HomeRider) Ride(vehicle string) {
	fmt.Printf("집에 %v를 타고 갑니다.\n", vehicle)
}

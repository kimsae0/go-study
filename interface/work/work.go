package work

import "fmt"

type WorkRider struct {
}

func (w *WorkRider) Ride(vehicle string) {
	fmt.Printf("직장에 %v를 타고 갑니다.\n", vehicle)
}


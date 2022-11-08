package school

import "fmt"

type SchoolRider struct {
}

func (s *SchoolRider) Send(vehicle string) {
	fmt.Printf("학교에 %v를 타고 갑니다.\n", vehicle)
}

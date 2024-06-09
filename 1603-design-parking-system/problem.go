package designparkingsystem

type ParkingSystem struct {
	lots [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		lots: [3]int{big, medium, small},
	}
}

func (s *ParkingSystem) AddCar(carType int) bool {
	if s.lots[carType-1] > 0 {
		s.lots[carType-1]--
		return true
	}
	return false
}

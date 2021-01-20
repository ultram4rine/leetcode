package main

func main() {}

type ParkingSystem struct {
	big    int
	medium int
	small  int

	bigCars    int
	mediumCars int
	smallCars  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		{
			if this.bigCars == this.big {
				return false
			}
			this.bigCars++
			return true
		}
	case 2:
		{
			if this.mediumCars == this.medium {
				return false
			}
			this.mediumCars++
			return true
		}
	case 3:
		{
			if this.smallCars == this.small {
				return false
			}
			this.smallCars++
			return true
		}
	default:
		return false
	}
}

 package speed

// TODO: define the 'Car' type struct

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	battery := 100
	distance := 0
    return Car {
        battery: battery,
        batteryDrain: batteryDrain,
        distance: distance,
        speed: speed,
    }
}

// TODO: define the 'Track' type struct

type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	speed := car.speed
    batteryDrain := car.batteryDrain
    var battery int
    var distance int
    if car.battery >= car.batteryDrain {
        battery = car.battery - car.batteryDrain
        distance = car.distance + speed
    } else {
    	battery = car.battery
        distance = car.distance
    }
    return Car {
        speed: speed,
        distance: distance,
        batteryDrain: batteryDrain,
        battery: battery,
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distance := track.distance - car.distance
    a_required_ratio := distance/car.speed
    battery_consumed := a_required_ratio * car.batteryDrain
    battery_left := car.battery - battery_consumed

    if battery_left > 0 {
        return true
    } else {
    	return false
    }
}

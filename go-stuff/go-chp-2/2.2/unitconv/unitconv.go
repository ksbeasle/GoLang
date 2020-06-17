package unitconv

type Ft float64
type Pounds float64
type Meters float64
type Kilograms float64

func FeetToMeters(f Ft) Meters {
	return Meters(f / 3.2808)
}

func PoundsToKilograms(p Pounds) Kilograms {
	return Kilograms(p * 0.45359237)
}

package zoo

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLong
	SetLocation(LatLong)
	GetName() string
	SetName(string)
	CanFly() bool
	Speak() string
}

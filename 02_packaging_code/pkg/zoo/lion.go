package zoo

type Lion struct {
	name       string
	manelength int
	location   LatLong
}

func (l *Lion) GetLocation() LatLong {
	return l.location
}

func (l *Lion) SetLocation(loc LatLong) {
	l.location = loc
}

func (l *Lion) CanFly() bool {
	return false
}

func (l *Lion) Speak() string {
	return "roar"
}

func (l *Lion) GetManeLength() int {
	return l.manelength
}

func (l *Lion) SetManeLength(lenght int) {
	l.manelength = lenght
}

func (l *Lion) GetName() string {
	return l.name
}

func (l *Lion) SetName(name string) {
	l.name = name
}

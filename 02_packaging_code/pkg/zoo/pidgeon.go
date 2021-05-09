package zoo

type Pigeon struct {
	name     string
	location LatLong
}

func (p *Pigeon) GetLocation() LatLong {
	return p.location
}

func (p *Pigeon) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pigeon) CanFly() bool {
	return true
}

func (p *Pigeon) Speak() string {
	return "hoot"
}

func (p *Pigeon) GetName() string {
	return p.name
}

func (p *Pigeon) SetName(name string) {
	p.name = name
}

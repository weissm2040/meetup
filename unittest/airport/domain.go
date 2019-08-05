package airport

// Airplane represents an aircraft in a particular location.
type Airplane struct {
	ID       string
	Name     string
	Location string
}

// Airport represents a physical airport and the airplanes currently at that airport.
type Airport struct {
	ID        string
	Name      string
	Airplanes []Airplane
}

package airport

//go:generate moq -out mock_test.go -pkg airport_test . AirplaneRepo Repo

import (
	"fmt"
)

// Repo represents methods for persisting and fetching airports.
type Repo interface {
	Save(Airport) (Airport, error)
	Fetch(string) (Airport, error)
}

// AirplaneRepo represents methods for persisting and fetching airplanes.
type AirplaneRepo interface {
	Save(Airplane) (Airplane, error)
	Fetch(string) (Airplane, error)
}

// Service can have an airplan arrive or depart from an airport.
type Service struct {
	Airports  Repo
	Airplanes AirplaneRepo
}

// Arrive will confirm the arrival of an airplane at an airport.
func (a *Service) Arrive(airplaneID string, airportID string) error {
	airplane, err := a.Airplanes.Fetch(airplaneID)
	if err != nil {
		return err
	}
	if airplane.Location != "" {
		return fmt.Errorf("airplane %v has not left previous location", airplaneID)
	}

	airport, err := a.Airports.Fetch(airportID)
	if err != nil {
		return err
	}

	airplane.Location = airport.ID
	airport.Airplanes = append(airport.Airplanes, airplane)

	_, err = a.Airplanes.Save(airplane)
	if err != nil {
		return err
	}

	_, err = a.Airports.Save(airport)
	if err != nil {
		return err
	}

	return nil
}

// Depart will confirm departure of an airplane from an airport.
func (a *Service) Depart(airplaneID string, airportID string) error {
	airplane, err := a.Airplanes.Fetch(airplaneID)
	if err != nil {
		return err
	}
	if airplane.Location == "" {
		return fmt.Errorf("airplane %v is in the air", airplaneID)
	}

	airport, err := a.Airports.Fetch(airportID)
	if err != nil {
		return err
	}
	if airplane.Location != airport.ID {
		return fmt.Errorf("airplane %v is not at airport %v", airplaneID, airportID)
	}

	airplane.Location = ""
	// remove airplane from airport
	newAirplanes := make([]Airplane, 0, len(airport.Airplanes)-1)
	for _, a := range airport.Airplanes {
		if a.ID != airplaneID {
			newAirplanes = append(newAirplanes, a)
		}
	}
	airport.Airplanes = newAirplanes

	_, err = a.Airplanes.Save(airplane)
	if err != nil {
		return err
	}

	_, err = a.Airports.Save(airport)
	if err != nil {
		return err
	}

	return nil
}

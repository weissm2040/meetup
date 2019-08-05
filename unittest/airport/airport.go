package airport

//go:generate moq -out mock_test.go -pkg airport_test . Transaction AirplaneRepo AirportRepo

import (
	"fmt"
)

// Transaction represents methods to start and commit a transaction.
type Transaction interface {
	Begin() error
	Commit() error
}

// AirportRepo represents methods for persisting and fetching airports.
type AirportRepo interface {
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
	tx        Transaction
	airports  AirportRepo
	airplanes AirplaneRepo
}

// NewAirportService creates a new airport service instance.
func NewAirportService(transaction Transaction, airports AirportRepo, airplanes AirplaneRepo) *Service {
	return &Service{
		tx:        transaction,
		airports:  airports,
		airplanes: airplanes,
	}
}

// Arrive will confirm the arrival of an airplane at an airport.
func (a *Service) Arrive(airplaneID string, airportID string) error {
	err := a.tx.Begin()
	if err != nil {
		return err
	}

	airplane, err := a.airplanes.Fetch(airplaneID)
	if err != nil {
		return err
	}
	if airplane.Location != "" {
		return fmt.Errorf("airplane %v has not left previous location", airplaneID)
	}

	airport, err := a.airports.Fetch(airportID)
	if err != nil {
		return err
	}

	airplane.Location = airport.ID
	airport.Airplanes = append(airport.Airplanes, airplane)

	_, err = a.airplanes.Save(airplane)
	if err != nil {
		return err
	}

	_, err = a.airports.Save(airport)
	if err != nil {
		return err
	}

	return nil
}

// Depart will confirm departure of an airplane from an airport.
func (a *Service) Depart(airplaneID string, airportID string) error {
	airplane, err := a.airplanes.Fetch(airplaneID)
	if err != nil {
		return err
	}
	if airplane.Location == "" {
		return fmt.Errorf("airplane %v is in the air", airplaneID)
	}

	airport, err := a.airports.Fetch(airportID)
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

	_, err = a.airplanes.Save(airplane)
	if err != nil {
		return err
	}

	_, err = a.airports.Save(airport)
	if err != nil {
		return err
	}

	return nil
}

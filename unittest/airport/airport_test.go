package airport_test

import (
	"fmt"
	"testing"

	"github.com/weissm2040/meetup/unittest/airport"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAirportServiceArrive(t *testing.T) {
	Convey("Given an airport service", t, func() {
		tx := &TransactionMock{
			BeginFunc: func() error {
				return nil
			},
			CommitFunc: func() error {
				return nil
			},
		}
		airports := &AirportRepoMock{
			FetchFunc: func(id string) (airport.Airport, error) {
				return airport.Airport{
					ID:   id,
					Name: "Adelaide Airport",
				}, nil
			},
			SaveFunc: func(in airport.Airport) (airport.Airport, error) {
				return in, nil
			},
		}
		airplanes := &AirplaneRepoMock{
			FetchFunc: func(id string) (airport.Airplane, error) {
				return airport.Airplane{
					ID:   id,
					Name: "Boeing 737 1000A",
				}, nil
			},
			SaveFunc: func(in airport.Airplane) (airport.Airplane, error) {
				return in, nil
			},
		}
		service := airport.NewAirportService(tx, airports, airplanes)
		Convey("When all repo calls succeed", func() {
			Convey("Then arrive should succeed", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeNil)
				So(airplanes.SaveCalls(), ShouldHaveLength, 1)
				So(airplanes.SaveCalls()[0].In1, ShouldResemble, airport.Airplane{
					ID:       "airplane1",
					Name:     "Boeing 737 1000A",
					Location: "airport1",
				})
				So(airports.SaveCalls(), ShouldHaveLength, 1)
				So(airports.SaveCalls()[0].In1, ShouldResemble, airport.Airport{
					ID:   "airport1",
					Name: "Adelaide Airport",
					Airplanes: []airport.Airplane{
						airport.Airplane{
							ID:       "airplane1",
							Name:     "Boeing 737 1000A",
							Location: "airport1",
						},
					},
				})
			})
		})
		Convey("When fetching the airplane errors", func() {
			airplanes.FetchFunc = func(_ string) (airport.Airplane, error) {
				return airport.Airplane{}, fmt.Errorf("fetch airplane error")
			}
			Convey("Then arrive should error", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeError, "fetch airplane error")
			})
		})
		Convey("When the airplane is at an airport", func() {
			airplanes.FetchFunc = func(id string) (airport.Airplane, error) {
				return airport.Airplane{
					ID:       id,
					Name:     "Boeing 737 1000A",
					Location: "airport2",
				}, nil
			}
			Convey("Then arrive should error", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeError, "airplane airplane1 has not left previous location")
			})
		})
		Convey("When fetching the airport errors", func() {
			airports.FetchFunc = func(_ string) (airport.Airport, error) {
				return airport.Airport{}, fmt.Errorf("fetch airport error")
			}
			Convey("Then arrive should error", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeError, "fetch airport error")
			})
		})
		Convey("When saving the airplane errors", func() {
			airplanes.SaveFunc = func(_ airport.Airplane) (airport.Airplane, error) {
				return airport.Airplane{}, fmt.Errorf("save airplane error")
			}
			Convey("Then arrive should error", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeError, "save airplane error")
			})
		})
		Convey("When saving the airport errors", func() {
			airports.SaveFunc = func(_ airport.Airport) (airport.Airport, error) {
				return airport.Airport{}, fmt.Errorf(("save airport error"))
			}
			Convey("Then arrive should error", func() {
				err := service.Arrive("airplane1", "airport1")
				So(err, ShouldBeError, "save airport error")
			})
		})
	})
}

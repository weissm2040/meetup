// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package airport_test

import (
	"github.com/weissm2040/meetup/unittest/airport"
	"sync"
)

var (
	lockAirplaneRepoMockFetch sync.RWMutex
	lockAirplaneRepoMockSave  sync.RWMutex
)

// Ensure, that AirplaneRepoMock does implement AirplaneRepo.
// If this is not the case, regenerate this file with moq.
var _ airport.AirplaneRepo = &AirplaneRepoMock{}

// AirplaneRepoMock is a mock implementation of AirplaneRepo.
//
//     func TestSomethingThatUsesAirplaneRepo(t *testing.T) {
//
//         // make and configure a mocked AirplaneRepo
//         mockedAirplaneRepo := &AirplaneRepoMock{
//             FetchFunc: func(in1 string) (airport.Airplane, error) {
// 	               panic("mock out the Fetch method")
//             },
//             SaveFunc: func(in1 airport.Airplane) (airport.Airplane, error) {
// 	               panic("mock out the Save method")
//             },
//         }
//
//         // use mockedAirplaneRepo in code that requires AirplaneRepo
//         // and then make assertions.
//
//     }
type AirplaneRepoMock struct {
	// FetchFunc mocks the Fetch method.
	FetchFunc func(in1 string) (airport.Airplane, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(in1 airport.Airplane) (airport.Airplane, error)

	// calls tracks calls to the methods.
	calls struct {
		// Fetch holds details about calls to the Fetch method.
		Fetch []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// In1 is the in1 argument value.
			In1 airport.Airplane
		}
	}
}

// Fetch calls FetchFunc.
func (mock *AirplaneRepoMock) Fetch(in1 string) (airport.Airplane, error) {
	if mock.FetchFunc == nil {
		panic("AirplaneRepoMock.FetchFunc: method is nil but AirplaneRepo.Fetch was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockAirplaneRepoMockFetch.Lock()
	mock.calls.Fetch = append(mock.calls.Fetch, callInfo)
	lockAirplaneRepoMockFetch.Unlock()
	return mock.FetchFunc(in1)
}

// FetchCalls gets all the calls that were made to Fetch.
// Check the length with:
//     len(mockedAirplaneRepo.FetchCalls())
func (mock *AirplaneRepoMock) FetchCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockAirplaneRepoMockFetch.RLock()
	calls = mock.calls.Fetch
	lockAirplaneRepoMockFetch.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *AirplaneRepoMock) Save(in1 airport.Airplane) (airport.Airplane, error) {
	if mock.SaveFunc == nil {
		panic("AirplaneRepoMock.SaveFunc: method is nil but AirplaneRepo.Save was just called")
	}
	callInfo := struct {
		In1 airport.Airplane
	}{
		In1: in1,
	}
	lockAirplaneRepoMockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	lockAirplaneRepoMockSave.Unlock()
	return mock.SaveFunc(in1)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedAirplaneRepo.SaveCalls())
func (mock *AirplaneRepoMock) SaveCalls() []struct {
	In1 airport.Airplane
} {
	var calls []struct {
		In1 airport.Airplane
	}
	lockAirplaneRepoMockSave.RLock()
	calls = mock.calls.Save
	lockAirplaneRepoMockSave.RUnlock()
	return calls
}

var (
	lockRepoMockFetch sync.RWMutex
	lockRepoMockSave  sync.RWMutex
)

// Ensure, that RepoMock does implement Repo.
// If this is not the case, regenerate this file with moq.
var _ airport.Repo = &RepoMock{}

// RepoMock is a mock implementation of Repo.
//
//     func TestSomethingThatUsesRepo(t *testing.T) {
//
//         // make and configure a mocked Repo
//         mockedRepo := &RepoMock{
//             FetchFunc: func(in1 string) (airport.Airport, error) {
// 	               panic("mock out the Fetch method")
//             },
//             SaveFunc: func(in1 airport.Airport) (airport.Airport, error) {
// 	               panic("mock out the Save method")
//             },
//         }
//
//         // use mockedRepo in code that requires Repo
//         // and then make assertions.
//
//     }
type RepoMock struct {
	// FetchFunc mocks the Fetch method.
	FetchFunc func(in1 string) (airport.Airport, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(in1 airport.Airport) (airport.Airport, error)

	// calls tracks calls to the methods.
	calls struct {
		// Fetch holds details about calls to the Fetch method.
		Fetch []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// In1 is the in1 argument value.
			In1 airport.Airport
		}
	}
}

// Fetch calls FetchFunc.
func (mock *RepoMock) Fetch(in1 string) (airport.Airport, error) {
	if mock.FetchFunc == nil {
		panic("RepoMock.FetchFunc: method is nil but Repo.Fetch was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockRepoMockFetch.Lock()
	mock.calls.Fetch = append(mock.calls.Fetch, callInfo)
	lockRepoMockFetch.Unlock()
	return mock.FetchFunc(in1)
}

// FetchCalls gets all the calls that were made to Fetch.
// Check the length with:
//     len(mockedRepo.FetchCalls())
func (mock *RepoMock) FetchCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockRepoMockFetch.RLock()
	calls = mock.calls.Fetch
	lockRepoMockFetch.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *RepoMock) Save(in1 airport.Airport) (airport.Airport, error) {
	if mock.SaveFunc == nil {
		panic("RepoMock.SaveFunc: method is nil but Repo.Save was just called")
	}
	callInfo := struct {
		In1 airport.Airport
	}{
		In1: in1,
	}
	lockRepoMockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	lockRepoMockSave.Unlock()
	return mock.SaveFunc(in1)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedRepo.SaveCalls())
func (mock *RepoMock) SaveCalls() []struct {
	In1 airport.Airport
} {
	var calls []struct {
		In1 airport.Airport
	}
	lockRepoMockSave.RLock()
	calls = mock.calls.Save
	lockRepoMockSave.RUnlock()
	return calls
}

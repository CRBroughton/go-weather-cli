package weather

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type WeatherTest struct {
	name          string
	server        *httptest.Server
	response      WeatherResponse
	expectedError error
}

func TestGetCurrentWeather(t *testing.T) {
	fakeResponse := WeatherResponse{
		Lat:            37.7749,
		Lon:            -122.4194,
		Timezone:       "America/Los_Angeles",
		TimezoneOffset: -25200,
		Current: CurrentWeather{
			Dt:         1679167420,
			Sunrise:    1679153840,
			Sunset:     1679199540,
			Temp:       72.5,
			FeelsLike:  74.2,
			Pressure:   1013,
			Humidity:   56,
			DewPoint:   57.3,
			Uvi:        3.2,
			Clouds:     0,
			Visibility: 10000,
			WindSpeed:  8.5,
			WindDeg:    45,
			WindGust:   11.4,
			Weather: []WeatherDescription{
				{
					ID:          800,
					Main:        "Clear",
					Description: "clear sky",
					Icon:        "01d",
				},
			},
		},
	}

	byteResponse, err := json.Marshal(fakeResponse)
	if err != nil {
		t.Fatal("Failed to convert mock response to byte[]")
	}

	test := WeatherTest{
		name: "Weather test",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(byteResponse)
		})),
		response:      fakeResponse,
		expectedError: nil,
	}

	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		res, err := getCurrentWeather(test.server.URL)

		if !reflect.DeepEqual(res, test.response) {
			t.Errorf("FAILED TEST: expected %v, got %v\n", test.response, res)
		}

		if !errors.Is(err, test.expectedError) {
			t.Errorf("Expected error FAILED: expected %v, got %v\n", test.expectedError, err)
		}
	})
}

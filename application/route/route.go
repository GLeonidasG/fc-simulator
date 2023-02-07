package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

type PartialRoutePosition struct {
	ID        string    `json:"routeId"`
	ClientID  string    `json:"clientId"`
	Positions []float64 `json:"positions"`
	Finished  bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID must be informed")
	}

	file, err := os.Open("destinations/"+r.ID+".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})

	}

	return nil

}

func (r *Route) ExportJsonPosition() ([]string, error) {
  var route PartialRoutePosition
  var results []string
  var totals = len(r.Positions) - 1
  for k, v := range r.Positions {
    route.ID = r.ID
    route.ClientID = r.ClientID
    route.Positions = []float64{v.Lat, v.Long}
    route.Finished = false
    if totals == k {
      route.Finished = true
    }
    jsonRoute, err := json.Marshal(route)
    if err != nil {
      return nil, err
    }

    results = append(results, string(jsonRoute))
  }
  return results, nil
}

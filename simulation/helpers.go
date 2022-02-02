package simulation

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

const (
	MAP_CITY_SEPARATOR      = " "
	MAP_DIRECTION_SEPARATOR = "="
)

// loadMap parses the map input file and constructs the map and city objects.
func loadMap(mapFilePath string) (map[string]*City, map[string][]Road, error) {
	cities := make(map[string]*City)
	simMap := make(map[string][]Road)

	file, err := os.Open(mapFilePath)
	if err != nil {
		return cities, simMap, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), MAP_CITY_SEPARATOR)
		fromCity := splitLine[0]
		cities[fromCity] = &City{Name: fromCity}
		simMap[fromCity] = make([]Road, 0)
		for i := 1; i < len(splitLine); i++ {
			direction := strings.Split(splitLine[i], MAP_DIRECTION_SEPARATOR)[0]
			toCity := strings.Split(splitLine[i], MAP_DIRECTION_SEPARATOR)[1]
			if _, ok := cities[toCity]; !ok {
				cities[toCity] = &City{Name: toCity}
			}

			simMap[fromCity] = append(simMap[fromCity], Road{Direction: direction, Destination: toCity})
		}
	}

	return cities, simMap, nil
}

// initAliens creates a list of aliens with unique ids and randomly generated names.
func initAliens(numAliens int) map[int]*Alien {
	aliens := make(map[int]*Alien)

	for i := 0; i < numAliens; i++ {
		alien := Alien{
			Name: fmt.Sprintf("%v %v", randomdata.SillyName(), randomdata.LastName()),
		}
		aliens[i] = &alien
	}

	return aliens
}

// prepSimulation constructs a simulation object with the initial simulation state.
func prepSimulation(mapFilePath string, numAliens int) (Simulation, error) {
	cities, simMap, err := loadMap(mapFilePath)
	if err != nil {
		return Simulation{}, err
	}

	if len(cities) == 0 {
		return Simulation{}, errors.New("Simulation must have at least one city")
	}

	aliens := initAliens(numAliens)

	// Construct list of city names used for random selection.
	cityNames := make([]string, len(cities))

	i := 0
	for city := range cities {
		cityNames[i] = city
		i++
	}

	// Randomly assign aliens to starting cities.
	for _, alien := range aliens {
		startingCity := cityNames[rand.Intn(len(cityNames))]
		alien.CurrentCity = cities[startingCity]
	}

	// Create simulation object used for simulation.
	simulation := Simulation{
		Cities: cities,
		SimMap: simMap,
		Aliens: aliens,
	}

	return simulation, nil
}

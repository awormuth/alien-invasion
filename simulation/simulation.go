package simulation

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CITY_DESTRUCTION_THRESH = 2
	SIMULATION_ITERATIONS   = 10000
)

func (simulation *Simulation) updateAliens() {
	for alienId, alien := range simulation.Aliens {
		// Choose next destination randomly.
		roads := simulation.SimMap[alien.CurrentCity.Name]

		destinationCity := ""
		for i := 0; i < len(roads); i++ {
			road := roads[rand.Intn(len(roads))]

			// Only choose cities that aren't destroyed.
			if !simulation.Cities[road.Destination].IsDestroyed {
				destinationCity = road.Destination
				break
			}
		}

		// Alien is trapped.
		if destinationCity == "" {
			continue
		}

		simulation.Aliens[alienId].CurrentCity = simulation.Cities[destinationCity]
	}
}

// destroyCities checks each city and destroys it if there are too many aliens.
func (simulation *Simulation) destroyCities() {
	aliensInCities := make(map[string][]int)

	for alienId, alien := range simulation.Aliens {
		aliensInCities[alien.CurrentCity.Name] = append(aliensInCities[alien.CurrentCity.Name], alienId)
	}

	for city, alienIds := range aliensInCities {
		// Destroy the city of there are too many aliens.
		if len(aliensInCities[city]) >= CITY_DESTRUCTION_THRESH {
			simulation.Cities[city].IsDestroyed = true

			fmt.Printf("%v has been destroyed by aliens %v\n", city, alienIds)

			// Also destroy aliens that are currently in the city.
			for _, alienId := range alienIds {
				delete(simulation.Aliens, alienId)
			}
		}
	}
}

// printCities prints cities for debugging purposes.
func (simulation *Simulation) printCities() {
	for _, city := range simulation.Cities {
		fmt.Printf("{%v %v} ", city.Name, city.IsDestroyed)
	}

	fmt.Println()
}

// printAlienState prints alien states for debugging purposes.
func (simulation *Simulation) printAlienState() {
	for alienId, alien := range simulation.Aliens {
		fmt.Printf("{%v %v %v} ", alienId, alien.Name, alien.CurrentCity)
	}

	fmt.Println()
}

// printSimMap prints the current state of the simulation map.
func (simulation *Simulation) printSimMap() {
	fmt.Println()
	for city, roads := range simulation.SimMap {
		if simulation.Cities[city].IsDestroyed {
			continue
		}

		fmt.Print(city)

		for _, road := range roads {
			if simulation.Cities[road.Destination].IsDestroyed {
				continue
			}

			fmt.Printf(" %v=%v", road.Direction, road.Destination)
		}

		fmt.Println()
	}
	fmt.Println()
}

// printStartMessage prints a message to information the user simulation is starting.
func (simulation *Simulation) printStartMessage() {
	fmt.Println()
	fmt.Println("👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐")
	fmt.Println("👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐")
	fmt.Println("👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐👾👽🪐")
	fmt.Println()
	fmt.Printf("Running simulation with %v aliens using map:", len(simulation.Aliens))
	simulation.printSimMap()
}

// start updates the alien state SIMULATION_ITERATIONS times, or until all of the aliens are destroyed.
func (simulation *Simulation) start() {
	simulation.printStartMessage()

	for i := 0; i < SIMULATION_ITERATIONS; i++ {
		simulation.updateAliens()
		simulation.destroyCities()
		if len(simulation.Aliens) == 0 {
			break
		}
	}

	fmt.Print("\nRemaining cities:")
	simulation.printSimMap()
}

func Run(mapFilePath string, numAliens int) (Simulation, error) {
	// Seed rand for non-deterministic simulations.
	rand.Seed(time.Now().UnixNano())

	simulation, err := prepSimulation(mapFilePath, numAliens)
	if err != nil {
		return Simulation{}, err
	}

	simulation.start()

	return simulation, nil
}

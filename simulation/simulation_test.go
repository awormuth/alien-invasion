package simulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPrepSimulation makes sure cities, roads, and aliens are loaded properly into the simulation object.
func TestPrepSimulation(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/single_path", 100)
	if assert.Nil(t, err) {
		assert.Equal(t, 100, len(simulation.Aliens), "Simulation should have 100 aliens.")

		for _, city := range simulation.Cities {
			assert.Equal(t, false, city.IsDestroyed, "City should not be destroyed.")
		}
	}
}

// TestCityDestruction makes city cities are destroyed properly under normal simulation conditions.
func TestCityDestruction(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/no_paths", 2)
	if assert.Nil(t, err) {
		simulation.start()

		assert.Less(t, len(simulation.Aliens), 1000, "Simulation should have less than 100 aliens.")
		assert.Equal(t, true, simulation.Cities["Foo"].IsDestroyed, "Foo should be destroyed.")
	}
}

// TestAlienMovement tests the behavior where an alien can move from one city to another.
func TestAlienMovement(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/single_path", 1)
	if assert.Nil(t, err) {
		currentCity := simulation.Aliens[0].CurrentCity
		simulation.updateAliens()
		assert.NotEqual(t, currentCity, simulation.Aliens[0].CurrentCity, "Alien should not be in the same city.")
	}
}

// TestAlienTrapped tests the "trapped" behavior where an alien cannot move.
func TestAlienTrapped(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/no_paths", 1)
	if assert.Nil(t, err) {
		currentCity := simulation.Aliens[0].CurrentCity
		simulation.updateAliens()
		assert.Equal(t, currentCity, simulation.Aliens[0].CurrentCity, "Alien be in the same city.")
	}
}

// TestNoCities ensures an error message is fired if the map has no cities.
func TestNoCities(t *testing.T) {
	_, err := prepSimulation("../city_maps/no_cities", 1)
	assert.Error(t, err, "Simulation must have at least one city")
}

// TestMediumSparseCity ensures the medium_sparse map can be loaded and simulated.
func TestMediumSparseCity(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/medium_sparse", 10000)
	if assert.Nil(t, err) {
		simulation.start()
	}
}

// TestLargeDenseCity ensures the large_dense map can be loaded and simulated.
func TestLargeDenseCity(t *testing.T) {
	simulation, err := prepSimulation("../city_maps/large_dense", 10000)
	if assert.Nil(t, err) {
		simulation.start()
	}
}

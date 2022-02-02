package main

import (
	"alien_invasion/simulation"
	"flag"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	mapFilePath string
	numAliens   int
	useWeb      bool
)

// RunSimulation is a REST API endpoint used for running simulations from the web.
func RunSimulation(c *gin.Context) {
	numAliens, err := strconv.Atoi(c.PostForm("num_aliens"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error parsing num_aliens parameter. Please make sure you're using an integer."})
		fmt.Println(err)
		return
	}

	sim, err := simulation.Run(c.PostForm("map_file_path"), numAliens)
	if err != nil {
		c.JSON(500, gin.H{"error": "An unexpected error occured while running the simulation. Please make sure you're using a valid map file."})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"Result": sim})
}

func main() {
	flag.StringVar(&mapFilePath, "map_file_path", "", "Location of map file used for simulation")
	flag.IntVar(&numAliens, "num_aliens", 0, "Number of aliens used for simulation")
	flag.BoolVar(&useWeb, "use_web", false, "Flag to use web interface instead of command line")
	flag.Parse()

	// If the user does not specify the use_web flag, it will use other command line flags to run the simulation.
	if !useWeb {
		_, err := simulation.Run(mapFilePath, numAliens)
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	// Start a REST API for managing similation requests from a web client.
	router := gin.Default()
	router.POST("/runSimulation", RunSimulation)
	router.Run(":8080")
}

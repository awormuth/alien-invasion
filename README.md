<img src="https://i.ibb.co/P6MGfpv/aliens.png" width="350">

## What is Alien Invasion?

Mad aliens are about to invade the earth and we are tasked with simulating the invasion. Full project spec can be seen [here](https://drive.google.com/file/d/1xxTCVMiyzhWTWmsKL4e5DRvUcEbMLy6g/view?usp=sharing).

- [x] Phase 1: Brainwash humans to build simulation of our invasion
- [x] Phase 2: Analyze simulation results to determine path of maximum destruction
- [ ] Phase 3: Initiate world annihilation 

## Code overview

The `simulation` package contains all of the logic for initializing and running a simulation. Structs and helper functions are split out into separate files for readability, while the core logic remains in `simulation/simulation.go`.

A simulation has a few key steps:

1. Initialization

    1. During initiatization, the map file is parsed and the simulation object is populated with cities, roads, and the initial alien state.

2. Simulation

    1. The simulation runs up to 10,000 iterations and stops early if all of the aliens are destroyed. 
    2. The simulation object is kept up to date with the state of the simulation as it's running. 
    3. Hashmaps are used for quick lookup and destruction of aliens.

3. Results

    1. The final state of the simulation is printed to the console for review.

<img src="https://i.ibb.co/z4WNdQr/alien.gif" width="600">

Tests can be found in `simulation/simulation_test.go`, which cover a variety of cases:

1. Unit tests for critical functions
    1. Map initialization
    2. City destruction
    3. Alien movement
2. Tests with various map inputs
    1. No cities
    2. Single city with no paths
    3. Two cities with single path
    4. Small map
    5. Medium sparse map
    6. Large dense map

## How to run?

You can run `main.go` with the following flags:

`map_file_path` - Map file used for the simulation

`num_aliens` - Number of aliens spawned for the simulation

__Example:__

```
go run main.go --map_file_path=city_maps/small_map --num_aliens=10
```

If you are missing project dependencies, you can run `go mod tidy` to pull them.

You can run the tests using the following command:

```
go test simulation/*
```

__Bonus:__

You can also run an alien invasion using a web interface! 

Step 1: Run the same `main.go` file using the `--use_web` flag.

```
go run main.go --use_web
```

Step 2: Open `index.html` in a web browser (tested with Chrome)

Step 3: Input `map_file_path` and `num_aliens` and click Submit

Step 4: See simulation results!

## Assumptions

1. The map and number of aliens will fit in memory. The current code will not handle extremely large maps or overwhelming numbers of aliens.
2. Every alien makes a move before determining whether a city can be destroyed. If needed, we can modify the simulation to "tick" on each individual alien move. 
3. The code assumes that map input is provided in a valid "structure" and format, meaning that roads lead to valid cities, and do point directly back to itself. We do not do extensive input validation, so a valid map file is important.
4. City names are unique. Since we use hashmaps for performance improvements, we require city names to be unique.
5. The map is treated more as a graph instead of a grid. Roads are treated as edges in a graph, rather than physical space or distance in a particular direction.

## Future Directions

1. Build a more sophisticated frontend to visualize the simulation and simulation results
2. Improve performance and scalability to handle larger maps and complex alien interactions
3. Get feedback from users on what to build next!
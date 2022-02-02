package simulation

type Alien struct {
	Name        string
	CurrentCity *City
}

type City struct {
	Name        string
	IsDestroyed bool
}

type Road struct {
	Direction   string
	Destination string
}

type Simulation struct {
	Cities map[string]*City
	SimMap map[string][]Road
	Aliens map[int]*Alien
}

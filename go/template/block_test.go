package template

func ExampleBlocks() {
	Blocks()
	// Output:
	// Names:
	// - Gamora
	// - Groot
	// - Nebula
	// - Rocket
	// - Star-Lord
	//
	// - Gamora
	// - Groot
	// - Nebula
	// - Rocket
	// - Star-Lord
	// Names:Gamora, Groot, Nebula, Rocket, Star-Lord
	// Names:Gamora, Groot, Nebula, Rocket, Star-Lord
	// Gamora, Groot, Nebula, Rocket, Star-Lord
}

package main

import (
	"fmt"

	"github.com/RudysAcosta/design_patterns/factory"
)

func main() {
	testFactory()
}

// test factory
func testFactory() {
	pokemonFactory := &factory.PokemonFactory{}

	// Create different types of Pokemon using the factory.
	charmander := pokemonFactory.CreatePokemon("Charmander")
	squirtle := pokemonFactory.CreatePokemon("Squirtle")
	bulbasaur := pokemonFactory.CreatePokemon("Bulbasaur")

	// Perform attacks of the created Pokemon.
	fmt.Println(charmander.Attack()) // Output: Charmander uses Flamethrower!
	fmt.Println(squirtle.Attack())   // Output: Squirtle uses Water Gun!
	fmt.Println(bulbasaur.Attack())  // Output: Bulbasaur uses Vine Whip!

}

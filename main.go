package main

import (
	"fmt"

	"github.com/RudysAcosta/design_patterns/abstract_factory/notification"
	"github.com/RudysAcosta/design_patterns/factory"
)

func main() {
	// testFactory()
	TestAbstractFactory()
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

func TestAbstractFactory() {
	smsFactory, _ := notification.GetNotificationFactory("SMS")
	emailFactory, _ := notification.GetNotificationFactory("Email")

	notification.SendNotification(smsFactory)
	notification.SendNotification(emailFactory)

	notification.GetMethod(smsFactory)
	notification.GetMethod(emailFactory)

}

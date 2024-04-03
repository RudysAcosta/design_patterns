package factory

type IPokemon interface {
	Attack() string
}

// Charmander is a kind pokemon
type Charmander struct{}

// Attack implements the Attack method of the Pokemon interface for Charmander.
func (c *Charmander) Attack() string {
	return "Charmander uses Flamethrower!"
}

// Squirtle is a type of Pokemon.
type Squirtle struct{}

// Attack implements the Attack method of the Pokemon interface for Charmander.
func (s *Squirtle) Attack() string {
	return "Squirtle uses Water Gun!"
}

// Bulbasaur is a type of Pokemon.
type Bulbasaur struct{}

// Attack implements the Attack method of the Pokemon interface for Charmander.
func (s *Bulbasaur) Attack() string {
	return "Bulbasaur uses Water Gun!"
}

type PokemonFactory struct{}

func (pf *PokemonFactory) CreatePokemon(pokemonType string) IPokemon {
	switch pokemonType {
	case "Charmander":
		return &Charmander{}
	case "Squirtle":
		return &Squirtle{}
	case "Bulbasaur":
		return &Bulbasaur{}
	default:
		return nil
	}
}

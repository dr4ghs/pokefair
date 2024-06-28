package types

import "fmt"

// PokemonType numeric representation of a Pokémon type as an unsinged 32 bit 
// integer. This data is used as a binary flag, in which every bit represents a 
// specific type.
type PokemonType uint32

const (
	NONE     PokemonType = 0
	NORMAL               = 1
	FIRE                 = 1 << 1
	WATER                = 1 << 2
	GRASS                = 1 << 3
	ELECTRIC             = 1 << 4
	ICE                  = 1 << 5
	FIGHTING             = 1 << 6
	POISON               = 1 << 7
	GROUND               = 1 << 8
	FLYING               = 1 << 9
	PSYCHIC              = 1 << 10
	BUG                  = 1 << 11
	ROCK                 = 1 << 12
	GHOST                = 1 << 13
	DRAGON               = 1 << 14
	DARK                 = 1 << 15
	STEEL                = 1 << 16
	FAIRY                = 1 << 17
)

// toUInt8 converts a PokemonType value to a unsigned 8 bit integer.
func toUInt8(t PokemonType) uint8 {
	var c uint8 = 0

	for t > 0 {
		t >>= 1
		c++
	}

	return c
}

// Type effectiveness numeric representation
// i (immune)   = 0 
// r (resisted) = 1
// n (neutral)  = 2
// w (weak)     = 3
const (
	i uint8 = iota
	r
	n
	w
)

// PokemonTypeChart Pokemon type effectiveness chart data type. 
// Each row represents a type, ranging from NORMAL (1) to FAIRY (18). Columns 
// represents compressed type modifier data. A single cell (unsigned 8 bit 
// integer) contains up to 4 different type's effectiveness modifier. 
// The smaller type is in the lowest bits of the byte.
type PokemonTypeChart [18][5]uint8

var typeChart = PokemonTypeChart{
	// NORMAL
	{ 
    n | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		r | i<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY
	},

	// FIRE
	{
		n | r<<2 | r<<4 | w<<6, // NORMAL, FIRE, WATER, GRASS
		n | w<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | w<<6, // GROUND, FLYING, PSYCHIC, BUG
		r | n<<2 | r<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		w | n<<2,               // STEEL, FAIRY 
	},

	// WATER
	{
		n | w<<2 | r<<4 | r<<6, // NORMAL, FIRE, WATER, GRASS
    n | n<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		w | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		w | n<<2 | r<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		n | n<<2,               // STEEL, FAIRY 
	},

	// GRASS
	{
		n | r<<2 | w<<4 | r<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | n<<4 | r<<6, // ELECTRIC, ICE, FIGHTING, POISON
		w | r<<2 | n<<4 | r<<6, // GROUND, FLYING, PSYCHIC, BUG
		w | n<<2 | r<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY  
	},

	// ELECTRIC
	{
		n | n<<2 | w<<4 | r<<6, // NORMAL, FIRE, WATER, GRASS
		r | n<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		i | w<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | r<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		n | n<<2,               // STEEL, FAIRY
	},

	// ICE
	{
		n | r<<2 | r<<4 | w<<6, // NORMAL, FIRE, WATER, GRASS
		n | r<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		w | w<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | w<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY
	},

	// FIGHTING
	{
		w | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | w<<2 | n<<4 | r<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | r<<2 | r<<4 | r<<6, // GROUND, FLYING, PSYCHIC, BUG
		w | i<<2 | n<<4 | w<<6, // ROCK, GHOST, DRAGON, DARK
		w | r<<2,               // STEEL, FAIRY
	},

	// POISON
	{
		n | n<<2 | n<<4 | w<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | n<<4 | r<<6, // ELECTRIC, ICE, FIGHTING, POISON
		r | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		r | r<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		i | w<<2,               // STEEL, FAIRY
	},

	// GROUND
	{
		n | w<<2 | n<<4 | r<<6, // NORMAL, FIRE, WATER, GRASS
		w | n<<2 | n<<4 | w<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | i<<2 | n<<4 | r<<6, // GROUND, FLYING, PSYCHIC, BUG
		w | n<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		w | n<<2,               // STEEL, FAIRY
	},

	// FLYING
	{
		n | n<<2 | n<<4 | w<<6, // NORMAL, FIRE, WATER, GRASS
		r | n<<2 | w<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | w<<6, // GROUND, FLYING, PSYCHIC, BUG
		r | n<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY
	},

	// PSYCHIC
	{
		n | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | w<<4 | w<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | r<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | n<<4 | i<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY
	},

	// BUG
	{
		n | r<<2 | n<<4 | w<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | r<<4 | r<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | r<<2 | w<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | r<<2 | n<<4 | w<<6, // ROCK, GHOST, DRAGON, DARK
		r | r<<2,               // STEEL, FAIRY
	},

	// ROCK
	{
		n | w<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | w<<2 | r<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		r | w<<2 | n<<4 | w<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY
	},

	// GHOST
	{
		i | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | w<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | w<<2 | n<<4 | r<<6, // ROCK, GHOST, DRAGON, DARK
		n | n<<2,               // STEEL, FAIRY
	},

	// DRAGON
	{
		n | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | w<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | i<<2,               // STEEL, FAIRY
	},

	// DARK
	{
		n | n<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | r<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | w<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | w<<2 | n<<4 | r<<6, // ROCK, GHOST, DRAGON, DARK
		n | r<<2,               // STEEL, FAIRY
	},

	// STEEL
	{
		n | r<<2 | r<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		r | w<<2 | n<<4 | n<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		w | n<<2 | n<<4 | n<<6, // ROCK, GHOST, DRAGON, DARK
		r | w<<2,               // STEEL, FAIRY },
  },

	// FAIRY
	{
		n | r<<2 | n<<4 | n<<6, // NORMAL, FIRE, WATER, GRASS
		n | n<<2 | w<<4 | r<<6, // ELECTRIC, ICE, FIGHTING, POISON
		n | n<<2 | n<<4 | n<<6, // GROUND, FLYING, PSYCHIC, BUG
		n | n<<2 | w<<4 | w<<6, // ROCK, GHOST, DRAGON, DARK
		r | n<<2,               // STEEL, FAIRY  
	},
}

// GetTypeMultiplier calculates the final damage multiplier of the attacker type
// calculated upon the defender type. If the defender type is combined, the
// function returns the result of the multiplication of both types multipliers.
func GetTypeMultiplier(
	attacker PokemonType,
	defender PokemonType,
) (float32, error) {
	// Checking if the attacker type is not invalid nor combined.
	if _, err := ValidateType(attacker, 1); err != nil {
		return 0, fmt.Errorf("Invalid attacker type, either invalid or composed")
	}

	// Checking if the defender type is not invalid nor composed by more than 2
	// types.
	if _, err := ValidateType(defender, 2); err != nil {
		return 0, fmt.Errorf("Invalid defender type, either invalid or composed by more than 2 types")
	}

	// Starting with a total multiplier of value 1, we iterate over all defender
	// types and multiply it by the single type multiplier.
	var multiplier float32 = 1
	var mask uint8 = 1

	for defender > 0 {
		if defender&1 > 0 {
			multiplier *= calculateMultiplier(toUInt8(attacker), mask, typeChart)
		}

		defender >>= 1
		mask++
	}

	return multiplier, nil
}

// ValidateType validates a Pokemon type. To be considered valid, a type cannot
// be NONE (0) or grater than 18. If composed, it cannot be composed by more
// than 2 types.
// The expected params indicates how many combined types are expected the type
// to be. If expected 2 but the passed type is not combined, it is still
// considered a valid type. This due to validation of Pokémon type.
func ValidateType(
	typ PokemonType,
	expected uint8,
) (PokemonType, error) {
	// Checking the the expected value is either 1 or 2.
	if expected < 1 || expected > 2 {
		return NONE, fmt.Errorf("Wrong typre quantity expected: %d", expected)
	}

	// Checking if the passed type values is in between 1 and 18.
	if toUInt8(typ) < 1 || toUInt8(typ) > 18 {
		return NONE, fmt.Errorf("Unknown type: value %d", typ)
	}

	// Type validation is achieved by checking the first less significant bit of
	// the type: if the bit is 1 a type is count and a counter is increased.
	// In order to check the next type, a right shift of 1 bit is performed on a
	// temporary type copy named test. This operation is repeated as long as the
	// temporary type is greater than 0. This allows to stop checking for other
	// types if we already reached the higher type value.
	test := typ
	var count uint8 = 0

	for test > 0 {
		count += uint8(test & 1)
		test >>= 1
	}

	// If the counted types are more than the expected, returns from the function
	// with an error.
	if count > expected {
		return NONE, fmt.Errorf("Too many types: got %d, expected %d", count, expected)
	}

	return typ, nil
}

// calculateMultiplier calculates the damage multiplayer of attacker type on
// defender type using the given type chart. Calculations are made with bitwise
// operations. It can return 2 for super effective damage, 1 for neutral damage,
// 0.5 for resisted damage or 0 for immune damage.
//
// Both attacker and defender type must not be combined types.
// Type santitization should be performed before using this function.
func calculateMultiplier(
	attacker uint8,
	defender uint8,
	typeChart PokemonTypeChart,
) float32 {
	// Extract the attacker type damaging chart
	chart := typeChart[attacker-1]

	// Since multiplier values are represented by half nibbles and four of these
	// values are compressed in one byte, it's necessary to perform some
	// calculations in order to extract the correct value.
	//
	// First we need to find in which byte the searched value is store using the
	// defender uint8 value minus one (since values type enum starts at 1 and we
	// need an array index) and dividing it by the max index of the chart array,
	// which is 4.
	selection := (defender - 1) / 4

	// Then we calculate the amount of left shifts we have to perform in order to
	// get the mask in the right position.
	// This is achieved by retrieving the reminder of the division of the type
	// index and the max chart index and multipling it by 2 (which is the mask's
	// bits amount).
	shifts := 2 * ((defender - 1) % 4)

	// The multiplier value is then extracted by left shifting the mask in the
	// right position, AND it with the byte cotaining the value and removing all
	// less significant zeros by right shifting the result with the same mask.
	value := (chart[selection] & (3 << shifts)) >> shifts

	// If the extracted values is 0, it means that we have an immunity. So we can
	// return from the function.
	if value == 0 {
		return 0
	}

	// Else, the real multiplier is calculated by raising 2 to the power of the
	// extracted value minus one and divided by 2.
	return float32(uint8(1<<(value-1))) / 2
}

package core

type PokemonNature uint8

const (
  HARDY   PokemonNature = PokemonNature((ATTACK << 4) | ATTACK)
  LONELY  PokemonNature = PokemonNature((ATTACK << 4) | DEFENSE)
  BRAVE   PokemonNature = PokemonNature((ATTACK << 4) | SPEED)
  ADAMANT PokemonNature = PokemonNature((ATTACK << 4) | SPECIAL_ATTACK)
  NAUGHTY PokemonNature = PokemonNature((ATTACK << 4) | SPECIAL_DEFENSE)
  BOLD    PokemonNature = PokemonNature((DEFENSE << 4) | ATTACK)
  DOCILE  PokemonNature = PokemonNature((DEFENSE << 4) | DEFENSE)
  RELAXED PokemonNature = PokemonNature((DEFENSE << 4) | SPEED)
  IMPISH  PokemonNature = PokemonNature((DEFENSE << 4) | SPECIAL_ATTACK)
  LAX     PokemonNature = PokemonNature((DEFENSE << 4) | SPECIAL_DEFENSE)
  TIMID   PokemonNature = PokemonNature((SPEED << 4) | ATTACK)
  HASTY   PokemonNature = PokemonNature((SPEED << 4) | DEFENSE)
  SERIOUS PokemonNature = PokemonNature((SPEED << 4) | SPEED)
  JOLLY   PokemonNature = PokemonNature((SPEED << 4) | SPECIAL_ATTACK)
  NAIVE   PokemonNature = PokemonNature((SPEED << 4) | SPECIAL_DEFENSE)
  MODEST  PokemonNature = PokemonNature((SPECIAL_ATTACK << 4) | ATTACK)
  MILD    PokemonNature = PokemonNature((SPECIAL_ATTACK << 4) | DEFENSE)
  QUIET   PokemonNature = PokemonNature((SPECIAL_ATTACK << 4) | SPEED)
  BASHFUL PokemonNature = PokemonNature((SPECIAL_ATTACK << 4) | SPECIAL_ATTACK)
  RASH    PokemonNature = PokemonNature((SPECIAL_ATTACK << 4) | SPECIAL_DEFENSE)
  CALM    PokemonNature = PokemonNature((SPECIAL_DEFENSE << 4) | ATTACK)
  GENTLE  PokemonNature = PokemonNature((SPECIAL_DEFENSE << 4) | DEFENSE)
  SASSY   PokemonNature = PokemonNature((SPECIAL_DEFENSE << 4) | SPEED)
  CAREFUL PokemonNature = PokemonNature((SPECIAL_DEFENSE << 4) | SPECIAL_ATTACK)
  QUIRKY  PokemonNature = PokemonNature((SPECIAL_DEFENSE << 4) | SPECIAL_DEFENSE)
)

func GetNatureModifier(
  nature PokemonNature, 
  stat PokemonStatsName,
) float32 {
  var modifier float32 = 1

  if PokemonStatsName(nature >> 4) == stat {
    modifier += 0.1
  }

  if PokemonStatsName(nature & 0b1111) == stat {
    modifier -= 0.1
  }

  return modifier
}

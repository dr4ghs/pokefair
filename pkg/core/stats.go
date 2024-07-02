package core

type PokemonStatsName uint8

const (
  STATS_NONE      PokemonStat = iota
  HP
  ATTACK
  DEFENSE
  SPECIAL_ATTACK
  SPECIAL_DEFENSE
  SPEED
)

type PokemonBaseStats [6]uint8

type PokemonStats [6]uint16

func CalculateStats(
  level uint8, 
  base PokemonBaseStats, 
  nature PokemonNature,
) PokemonStats {
  lvl := uint16(level)

  s[0] = ((2 * uint16(base[0]) * lvl) / 100) + lvl + 10

  for i := 1; i < len(base); i++ {
    stat := uint16(base[i])

    s[i] = ()
  }
}

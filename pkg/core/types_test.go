package core

import (
	"testing"
)

// ValidateType ================================================================
func TestValidateType(t *testing.T) {
	t.Run("GuardClausesCheck", validateType_guardClausesCheck)
	t.Run("ExpectedQuantityCheck", validateType_expectedQuantityCheck)
}

func validateType_guardClausesCheck(t *testing.T) {
	t.Log("Testing expected values")
	for i := 0; i < 4; i++ {
		_, err := ValidateType(NORMAL, uint8(i))

		if err == nil && (i < 1 || i > 2) {
			t.Errorf("Expected %d, should fail", i)
		} else if err != nil && (i > 0 && i < 3) {
			t.Errorf("Expected %d, hould pass", i)
		}
	}

	t.Log("Testing type numeric value")
	if _, err := ValidateType(NORMAL, 1); err != nil {
		t.Error("Type NORMAL is valid, should work")
	}

	if _, err := ValidateType(TYPES_NONE, 1); err == nil {
		t.Error("NONE type is not valid, should fail")
	}

	if _, err := ValidateType(20, 1); err == nil {
		t.Error("Unknown type of value 20 is not valid, should fail")
	}
}

func validateType_expectedQuantityCheck(t *testing.T) {
	t.Log("Testing NORMAL, must be 1")
	if _, err := ValidateType(NORMAL, 1); err != nil {
		t.Fail()
	}

	t.Log("Testing NORMAL with expected 2")
	if _, err := ValidateType(NORMAL, 2); err != nil {
		t.Error("Given 1, expected 2. Should work.")
	}

	t.Log("Testing NORMAL/FIRE, must be 2")
	if _, err := ValidateType(NORMAL|FIRE, 2); err != nil {
		t.Fail()
	}

	t.Log("Testing NORMAL/FIRE with expected 1")
	if _, err := ValidateType(NORMAL|FIRE, 1); err == nil {
		t.Error("Given 2, expected 1. Should fail.")
	}

	t.Log("Testing NORMAL/FIRE/WATER. Illegal, should fail ")
	if _, err := ValidateType(NORMAL|FIRE|WATER, 2); err == nil {
		t.Fail()
	}
}

// CalculateMultiplier =========================================================
func TestCalculateMultiplier(t *testing.T) {
	t.Log("FIRE on GRASS, expected 2")
	weakness := calculateMultiplier(toUInt8(FIRE), toUInt8(GRASS), typeChart)
	if weakness != 2 {
		t.Errorf("Multiplier should be 2, got %f\n", weakness)
	}

	t.Log("FIRE on NORMAL, expected 2")
	neutral := calculateMultiplier(toUInt8(FIRE), toUInt8(NORMAL), typeChart)
	if neutral != 1 {
		t.Errorf("Multiplier should be 1, got %f\n", neutral)
	}

	t.Log("Fire on ROCK, expected 0.5")
	resistence := calculateMultiplier(toUInt8(FIRE), toUInt8(ROCK), typeChart)
	if resistence != 0.5 {
		t.Errorf("Multiplier should be 0.5, got %f\n", resistence)
	}

	t.Log("GHOST on NORMAL, expected 0")
	immunity := calculateMultiplier(toUInt8(GHOST), toUInt8(NORMAL), typeChart)
	if immunity != 0 {
		t.Errorf("Multiplier should be 0, got %f\n", immunity)
	}
}

// GetTypeMultiplier ===========================================================
func TestGetTypeMultiplier(t *testing.T) {
	t.Run("GuardClauses", getTypeMultiplier_GuardClauses)
	t.Run("Immune", getTypeMultiplier_Immune)
	t.Run("Resisted", getTypeMultiplier_Resisted)
	t.Run("Neutral", getTypeMultiplier_Neutral)
	t.Run("SuperEffective", getTypeMultiplier_SuperEffective)
	t.Run("Mixed", getTypeMultiplier_Mixed)
}

func getTypeMultiplier_GuardClauses(t *testing.T) {
	t.Log("Testing attacker")
	if _, err := GetTypeMultiplier(FIRE, FIRE); err != nil {
		t.Error("Attacker type FIRE is not combined, should work")
	}

	if _, err := GetTypeMultiplier(FIRE|WATER, FIRE); err == nil {
		t.Error("Attacker type FIRE/WATER is combined, should not work")
	}

	if _, err := GetTypeMultiplier(TYPES_NONE, FIRE); err == nil {
		t.Error("Invalid attacker type NONE, should not work")
	}

	if _, err := GetTypeMultiplier(20, FIRE); err == nil {
		t.Error("Unknown attacker type of value 20 is not valid, should not work")
	}

  t.Log("Testing defender")
  if _, err := GetTypeMultiplier(FIRE, FIRE); err != nil {
    t.Error("Defender type FIRE is not composed, should workd")
  }

  if _, err := GetTypeMultiplier(FIRE, FIRE|NORMAL); err != nil {
    t.Error("Defender type FIRE/NORMAL is combined by 2, should work")
  }

  if _, err := GetTypeMultiplier(FIRE, TYPES_NONE); err == nil {
    t.Error("Invalid defender type NONE, should not work")
  }

  if _, err := GetTypeMultiplier(FIRE, FIRE|WATER|GRASS); err == nil {
    t.Error("Defender type FIRE/WATER/GRASS is combined by more than 2, should not work")
  }
}

func getTypeMultiplier_Immune(t *testing.T) {
	singleType, _ := GetTypeMultiplier(GROUND, FLYING)
	if singleType != 0 {
		t.Errorf("GROUND on FLYING: expected 0, got %f", singleType)
	}

	dualType, _ := GetTypeMultiplier(GROUND, FLYING|ELECTRIC)
	if dualType != 0 {
		t.Errorf("GROUND on FLYING/ELECTRIC: expected 0, got %f", dualType)
	}
}

func getTypeMultiplier_Resisted(t *testing.T) {
  singleType, _ := GetTypeMultiplier(FIRE, WATER)
  if singleType != 0.5 {
    t.Errorf("FIRE on WATER: expected 0.5, got %f", singleType)
  }

  dualType, _ := GetTypeMultiplier(FIRE, WATER|FIRE)
  if dualType != 0.25 {
    t.Errorf("FIRE on WATER/FIRE: expected 0.25, got %f", dualType)
  }
}

func getTypeMultiplier_Neutral(t *testing.T) {
  singleType, err := GetTypeMultiplier(NORMAL, FIGHTING)
  if singleType != 1 {
    t.Error(err)
    t.Errorf("FIRE on FIGHTING: expected 1, got %f", singleType)
  }

  dualType, err := GetTypeMultiplier(FIRE, FIGHTING|POISON)
  if dualType != 1 {
    t.Error(err)
    t.Errorf("FIRE on FIGHTING/POISON: expected 1, got %f", dualType)
  }
}

func getTypeMultiplier_SuperEffective(t *testing.T) {
  singleType, _ := GetTypeMultiplier(FAIRY, DRAGON)
  if singleType != 2 {
    t.Errorf("FAIRY on DRAGON: expected 2, got %f", singleType)
  }

  dualType, _ := GetTypeMultiplier(FAIRY, DRAGON|FIGHTING)
  if dualType != 4 {
    t.Errorf("FIRE on BUG/GRASS: expected 4, got %f", dualType)
  }
}

func getTypeMultiplier_Mixed(t *testing.T) {
  if value, _ := GetTypeMultiplier(GROUND, FLYING|GRASS); value != 0 {
    t.Errorf("GROUND on FLYING/GRASS: expected 0, got %f", value)
  }

  if value, _ := GetTypeMultiplier(GROUND, FLYING|NORMAL); value != 0 {
    t.Errorf("GROUND on FLYING/NORMAL: expected 0, got %f", value)
  }
  
  if value, _ := GetTypeMultiplier(GROUND, FLYING/ROCK); value != 0 {
    t.Errorf("GROUND on FLYING/ROCK: expected 0, got %f", value)
  }

  if value, _ := GetTypeMultiplier(GROUND, GRASS|NORMAL); value != 0.5 {
    t.Errorf("GROUND on GRASS/NORMAL: expected 0.5, got %f", value)
  }

  if value, _ := GetTypeMultiplier(GROUND, GRASS|STEEL); value != 1 {
    t.Errorf("GROUND on GRASS/STEEL: expected 1, got %f", value)
  }

  if value, _ := GetTypeMultiplier(GROUND, NORMAL|STEEL); value != 2 {
    t.Errorf("GROUND on NORMAL/STEEL: expected 2, got %f", value)
  }
}

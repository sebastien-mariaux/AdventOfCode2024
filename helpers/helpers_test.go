package helpers

import (
	"testing"
)


func TestRemoveAtIndex3(t *testing.T) {
  values := []int{90, 88, 86, 86, 84, 81}
  expected :=  []int{90, 88, 86, 84, 81}
  result := DeleteAtIndex(values, 3)
  if !equalSlices(result, expected) {
    t.Errorf("Failed - Expected %v, Got %v", expected, result)
  }
}

func TestRemoveAtIndex0(t *testing.T) {
  values := []int{90, 88, 86, 86, 84, 81}
  expected :=  []int{88, 86, 86, 84, 81}
  result := DeleteAtIndex(values, 0)
  if !equalSlices(result, expected) {
    t.Errorf("Failed - Expected %v, Got %v", expected, result)
  }
}

func equalSlices(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}
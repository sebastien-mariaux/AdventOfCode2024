package helpers

import (
	"os"
)

func OpenInput(sample bool) string {
  var file string
  if sample {
    file = "sample.txt"
  } else {
    file = "data.txt"
  }
  content, _ := os.ReadFile(file)
  return string(content)
}



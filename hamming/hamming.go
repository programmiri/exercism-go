//Package hamming adds a method to calculate
//the Hamming difference between two DNA strands.
package hamming

import "errors"

// Distance returns the Hamming difference between two given DNA strands.
func Distance(a, b string) (int, error) {
  if len(a) != len(b) {
    return 0, errors.New("Input values have to be the same length")
  }
  var count int
  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      count++
    }
  }
  return count, nil
}

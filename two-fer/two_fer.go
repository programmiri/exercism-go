// Package twofer creates a sentence
package twofer

//ShareWith creates a sentence that uses a given name
//or, if that is not given, "you"
func ShareWith(name string) string {
  if name == "" {
    name = "you"
  }
  return "One for " + name + ", one for me."
}

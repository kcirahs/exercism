// Package twofer takes a given name and returns a response using that name.
package twofer

// ShareWith takes the given name as a string and returns the phrase:
// "One for X, one for me." If the string is empty, it returns a general phrase.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."

}

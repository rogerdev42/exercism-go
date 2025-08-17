// Package twofer implements the classic "Two-fer" (two for one) exercise.
//
// The package provides a function ShareWith, which returns a string in the form:
//   "One for <name>, one for me."
//
// If no name is provided, the function defaults to "you":
//   "One for you, one for me."
package twofer

import "fmt"

// ShareWith returns a phrase sharing one cookie with the given name.
//
// If the name is an empty string, "you" is used instead.
// Examples:
//   ShareWith("Alice") -> "One for Alice, one for me."
//   ShareWith("")      -> "One for you, one for me."
func ShareWith(name string) string {
     if name == "" {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}

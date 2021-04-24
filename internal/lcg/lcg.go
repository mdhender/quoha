/*
 * quoha - a game server backed by mariadb
 *
 * Copyright (C) 2021  Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Package lcg implements a psuedo-random number generator
package lcg

type LCG interface {
	GetCurrent() int
	IntN(n int) int
	Seed(n int)
}

// New returns an initialized generator
func New(seed int) *Generator {
	return &Generator{x: seed}
}

// Generator implements an LCG for our PRNG.
type Generator struct {
	x int
}

// GetCurrent returns the current value.
func (ng *Generator) GetCurrent() int {
	return ng.x
}

// GetNext returns the LCG's current value and generates the next value.
// We return an int, but it's actually just 32 bits.
func (ng *Generator) GetNext() int {
	// constants from Knuth's MMIX
	ng.x = 6364136223846793005*ng.x + 1442695040888963407
	// returns only the most significant 32 bits
	return ng.GetCurrent()
}

// IntN returns, as an int, a non-negative pseudo-random number in [0,n).
func (ng *Generator) IntN(n int) int {
	ng.GetNext()
	t := (n * ng.x) >> 32
	if t < 0 {
		t *= -1
	}
	return t % n
}

// Seed updates the seed.
func (ng *Generator) Seed(seed int) {
	ng.x = seed
}

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

package state

import "testing"

// Specification: State API

func TestDefault(t *testing.T) {
	// Given a new state
	s := State{}
	// When we query the ID of the game
	id := s.GetID()
	// Then we should get an empty string
	if expected := ""; id != expected {
		t.Errorf("id: expected %q, got %q\n", expected, id)
	}
	// When we query the name of the game
	name := s.GetName()
	// Then we should get an empty string
	if expected := ""; name != expected {
		t.Errorf("name: expected %q, got %q\n", expected, name)
	}
	// When we execute the list of players
	players := s.GetListOfPlayers()
	// Then we should get an empty list of players
	if expected := 0; len(players) != expected {
		t.Errorf("players: expected %d, got %d\n", expected, len(players))
	}
}

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

type State struct {
	id      string // uuid
	name    string
	players Players
}

func (s State) GetID() string {
	return s.id
}

func (s State) GetName() string {
	return s.name
}

type Players []Player
type Player struct {
	id   string // uuid
	name string
}

func (s State) GetListOfPlayers() Players {
	return s.players
}

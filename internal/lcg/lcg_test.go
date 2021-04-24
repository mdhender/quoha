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

package lcg

import "testing"

func TestLCG(t *testing.T) {
	l := New(10)
	for i := 0; i < 1024; i++ {
		if actual := l.IntN(10); !(0 <= actual && actual < 10) {
			t.Errorf("expected 0 <= actual < 10, got %d\n", actual)
		}
	}
}

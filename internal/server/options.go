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

package server

import (
	"fmt"
	"github.com/mdhender/quoha/internal/way"
	"net"
	"time"
)

// Address sets the host and port that the server will listen on.
// It does not set the flag for http vs https.
func Address(host string, port int) func(*Server) error {
	return func(s *Server) error {
		s.Addr = net.JoinHostPort(host, fmt.Sprintf("%d", port))
		return nil
	}
}

// IdleTimeout sets the duration to wait before dropping a client connection.
func IdleTimeout(t time.Duration) func(*Server) error {
	return func(s *Server) error {
		s.IdleTimeout = t
		return nil
	}
}

// ReadTimeout sets the maximum time to wait for a read from a client ot complete.
func ReadTimeout(t time.Duration) func(*Server) error {
	return func(s *Server) error {
		s.ReadTimeout = t
		return nil
	}
}

// Router sets the mux used for the server.
func Router(r *way.Router) func(*Server) error {
	return func(s *Server) error {
		if r == nil {
			return fmt.Errorf("assert(router != nil)")
		}
		s.Handler = r
		return nil
	}
}

// WriteTimeout sets the maximum time to wait for a write to a client to complete.
func WriteTimeout(t time.Duration) func(*Server) error {
	return func(s *Server) error {
		s.WriteTimeout = t
		return nil
	}
}

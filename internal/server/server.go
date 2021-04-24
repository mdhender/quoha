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
	"database/sql"
	"github.com/mdhender/quoha/internal/way"
	"net/http"
	"time"
)

type Server struct {
	http.Server
	db     *sql.DB
	router *way.Router
}

// Default returns a server with updated timeouts per:
//   https://blog.cloudflare.com/exposing-go-on-the-internet/
//   https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
//   https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
func Default() *Server {
	s := Server{}
	s.Addr = ":3000"
	s.IdleTimeout = 10 * time.Second
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return &s
}

// ServeHTTP implements the http handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

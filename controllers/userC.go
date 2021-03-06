/*
 Domesticated Apricot - An open source member and event management platform
    Copyright (C) 2017  Sam Schurter

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controllers

import (
	"github.com/makeict/Domesticated-Apricot/views"
	"net/http"
)

//separate files for each individual controller
//a controller is a struct, with methods defined on the struct for each action

type UserController struct {
	Controller
}

var User UserController

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	user := []string{"first name", "last name"}
	if err := views.User.Index.Render(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

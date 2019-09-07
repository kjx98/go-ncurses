// Licensed under the GPL-v3
// Copyright: Sebastian Tilders <info@informatikonline.net> (c) 2019

package ncurses

import (
	// #include <curses.h>
	"C"
)

// Draws a border around w *Window.
func (w *Window) Box() {
    com := Command{
    	Name: DRAWBOX,
    	Scope: LOCAL,
    	Window: w,
    }
    w.sendCommand(com,true)
}
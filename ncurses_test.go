// Licensed under the GPL-v3
// Copyright: Sebastian Tilders <info@informatikonline.net> (c) 2019

package ncurses

import (
	"fmt"
	"os"
)

func Example_goTMDialog() {
	w, _ := Initscr()

	// Ensure, that ncurses will be properly exited
	defer Endwin()
	defer func() {
		if r := recover(); r != nil {
			Endwin()
			fmt.Printf("panic:\n%s\n", r)
			os.Exit(-1)
		}
	}()

	// Enable color mode
	StartColor()

	// Define color pairs
	AddColorPair("bw", ColorGreen, ColorBlack)
	AddColorPair("wb", ColorWhite, ColorBlue)

	// Set cursor visiblity to hidden
	SetCursor(CURSOR_HIDDEN)

	// Automatically refresh after each command
	w.AutoRefresh = true

	// Set color for stdscr-window to system defaults.
	w.Wbkgd("std")

	// Draw a border around main window (stdscr)
	w.Box()

	// Create a new window for greeting-text at cell (x=20,y=5) with a size of 25 x 5 cells.
	w2, err := NewWindow("dialog", Position{20, 5}, Size{25, 5})

	// This can fail if the terminal is too small.
	if err != nil {
		panic(err)
	}
	w2.AutoRefresh = true

	// Use color pair wb (2)
	w2.Wbkgd("wb")

	// Draw a border around our "Greeting Window".
	w2.Box()

	// Move cursor relative to the window borders of w2
	w2.Move(2, 3)

	// Output our greeting text
	fmt.Fprintf(w2, "Hello from Go\u2122-Lang!")

	// Move cursor relative to the beginning of our main window
	w.Move(17, 19)

	// Output exit instruction for the user
	fmt.Fprintf(w, " => Press a key to exit <=")

	// Wait for user input (e.g. keypress)
	w.Getch()
}

func ExampleInitscr_doThis() {
	fmt.Println("Before: This is okay!")
	w, _ := Initscr()

	// Normally, you could use defer Endwin()
	fmt.Fprintf(w, "While: This is also okay!")

	Endwin()
	fmt.Printf("After: This is also okay!")
	// do some things
}

func ExampleInitscr_dontDoThis() {
	Initscr()
	// Make sure, to call Endwin() before exiting
	defer Endwin()
	fmt.Println("Don't do this after you've called ncurses.InitScr()")
	// do some things
}

func ExampleAddColorPair() {
	w, _ := Initscr()
	defer Endwin()
	AddColorPair("mypair", ColorWhite, ColorBlue)
	SetColor("mypair")
	fmt.Fprintf(w, "Hello in White and Blue")
	// Output: Hello in White and Blue
}

func ExampleNewWindow() {
	Initscr()
	defer Endwin()
	w2, err := NewWindow("MyWindow", Position{4, 4}, Size{20, 4})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w2, "Hello from MyWindow")
}

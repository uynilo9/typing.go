package typing

import (
	"fmt"
	"os"
	"time"
)

type Typer struct {
	Content  string
	position uint
}

type Option struct {
	Count uint
	Delay time.Duration
}

func (t Typer) Type(o Option) Typer {
	if t.position+o.Count > uint(len(t.Content)) {
		fmt.Println("\n\u001B[31mERROR\u001B[0m The number of characters to type is out of range of the content")
		os.Exit(1)
	}
	if o.Count == 0 {
		o.Count = uint(len(t.Content))
	}
	for i := range o.Count {
		fmt.Print(string(t.Content[t.position+i]))
		time.Sleep(time.Millisecond * o.Delay)
	}

	t.position += o.Count
	return t
}

func (t Typer) Delete(o Option) Typer {
	if t.position-o.Count < uint(len(t.Content)) {
		fmt.Println("\n\u001B[31mERROR\u001B[0m The number of characters to delete is out of range of the content")
		os.Exit(1)
	}
	for range o.Count {
		fmt.Print("\b \b")
		time.Sleep(time.Millisecond * o.Delay)
	}

	t.position -= o.Count
	return t
}

func (t Typer) Paste(count uint) Typer {
	if t.position+count > uint(len(t.Content)) {
		fmt.Println("\n\u001B[31mERROR\u001B[0m The number of characters to paste is out of range of the content")
		os.Exit(1)
	}
	fmt.Print(string(t.Content[t.position : t.position+count]))

	t.position += count
	return t
}

func (t Typer) Cut(count uint) Typer {
	if t.position-count < uint(len(t.Content)) {
		fmt.Println("\n\u001B[31mERROR\u001B[0m The number of characters to cut is out of range of the content")
		os.Exit(1)
	}
	for range count {
		fmt.Print("\b \b")
	}

	t.position -= count
	return t
}

func (t Typer) Cursor(show bool) Typer {
	cursor := "\u001B[?25h"
	if !show {
		cursor = "\u001B[?25l"
	}
	fmt.Print(cursor)
	return t
}

func (t Typer) Sleep(d time.Duration) Typer {
	time.Sleep(time.Millisecond * d)
	return t
}

func (t Typer) Newline() {
	fmt.Println()
}

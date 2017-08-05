package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type brainfuck struct {
	source string
	cells  []byte
	cell   int
}

const memsize = 30000

func main() {
	if len(os.Args) < 1 {
		fmt.Println("No file provided")
		os.Exit(2)
	}

	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	bf := newBrainfuck(string(source))
	err = bf.run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func newBrainfuck(source string) (bf brainfuck) {
	bf.cells = make([]byte, memsize)
	bf.source = source
	return
}

func (bf brainfuck) run() (err error) {
	depth := 0
	var char rune
	for index := 0; index < len(bf.source); index++ {
		switch bf.source[index] {
		case '+':
			bf.cells[bf.cell]++
		case '-':
			bf.cells[bf.cell]--
		case '>':
			bf.cell++
		case '<':
			bf.cell--
		case '.':
			fmt.Printf("%c", bf.cells[bf.cell])
		case ',':
			_, err = fmt.Scanf("%c", &char)
			if err != nil {
				return
			}
			bf.cells[bf.cell] = byte(char)
		case '[':
			if bf.cells[bf.cell] == 0 {
				index++
				for depth > 0 || bf.source[index] != ']' {
					if bf.source[index] == '[' {
						depth++
					} else if bf.source[index] == ']' {
						depth--
					}
					index++
				}
			}
		case ']':
			//When only at one bracket of depth and cell value is 0, can just move to next index
			//rather than having to itterate back over to [ first
			//TODO could calcuate at other depths whether this is possible
			if depth > 0 && bf.cells[bf.cell] == 0 {
				depth--
				continue
			}
			index--
			for depth > 0 || bf.source[index] != '[' {
				if bf.source[index] == ']' {
					depth++
				} else if bf.source[index] == '[' {
					depth--
				}
				index--
			}
			//Loop moves pointer back up one so move behind one (without evalutating) to ensure correct location
			index--
		}
	}
	return
}

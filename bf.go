package main

import (
	"fmt"
	"os"
)

type BF struct {
	ptr   int
	tape  []int
	code  string
	pc    int
	table map[int]int
}

func newBF(code string, tape_len int) *BF {
	var bf BF
	bf.tape = make([]int, tape_len, tape_len)
	bf.code = code
	bf.table = make_table(code)
	return &bf
}

func make_table(code string) map[int]int {
	m := map[int]int{}
	stack := make([]int, 100)
	sp := 0
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case 91: // "["
			sp += 1
			stack[sp] = i
		case 93: // "]"
			key := stack[sp]
			m[key] = i
			sp -= 1
		default:
			// do nothing
		}
	}
	return m
}

func (bf *BF) eval() {
	for ; bf.pc < len(bf.code); bf.pc++ {
		switch bf.code[bf.pc] {
		case 43: // +
			bf.inc()
		case 45: // -
			bf.dec()
		case 60: // <
			bf.left()
		case 62: // >
			bf.right()
		case 91: // [
			if bf.tape[bf.ptr] == 0 {
				bf.pc = bf.table[bf.pc]
			}
		case 93: // ]
			for k, v := range bf.table {
				if v == bf.pc {
					bf.pc = k - 1
					break
				}
			}
		case 46: // .
			bf._print()
		case 44: // ,
			b := []byte{1}
			_, err := os.Stdin.Read(b)
			bf.tape[bf.ptr] = 0
			if err == nil {
				bf.tape[bf.ptr] = int(b[0])
			}
		default:
			// do nothing
		}
	}
}
func (bf *BF) inc() {
	bf.tape[bf.ptr]++
}
func (bf *BF) dec() {
	bf.tape[bf.ptr]--
}
func (bf *BF) right() {
	bf.ptr++
}
func (bf *BF) left() {
	bf.ptr--
}
func (bf *BF) _print() {
	fmt.Printf("%c", bf.tape[bf.ptr])
}
func main() {
	var code string
	code = ">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.[-]>++++++++[<++" +
		"++>-]<.>+++++++++++[<+++++>-]<.>++++++++[<+++>-]<.+++.------.--------.[-]>" +
		"++++++++[<++++>-]<+.[-]++++++++++."
	code = "+[,.]"
	bf := newBF(code, 16)
	bf.eval()
	//	fmt.Printf("\n%v\n", bf)
}

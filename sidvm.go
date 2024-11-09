// In this project I'll be build a virtual maching using go language

/* sidvm.go */

/*

	* 16 bit virtual machine

	* i.e. 65536 memory addresses or 65KB of memory

	* we will have 4 general purpose registers
	AX, BX, CX and DX
	(that we'll we use for addition and stuff like that)

	* we will have a stack pointer register
	SP

	* we will have a instruction pointer register(next instruction in the program to execute)
	IP

	*let's see what we'll add later


*/

package main

const stackSize = 65536

type Reg uint16

type Registers struct {
	ax, bx, cx, dx, sp, ip Reg
}

type CPU struct {
	r Registers
}

type IM struct {
	o    Opcode
	size int8
}

type Opcode int16

const (
	mov Opcode = 0x01
	nop Opcode = 0x02
)

type Instruction struct {
	o Opcode
	a []int16
}

// type Stack struct {
// }
type Stack [stackSize]int8

type Program []Instruction

type VM struct {
	c CPU
	s *Stack
	p *Program
}

var instrMAP = []IM{
	{o: mov, size: 0x03},
	{o: nop, size: 0x01},
}

// func virtualmachine(pr *Program, progsz int16) (*VM, error){
// 	if pr == nil || progsz == 0 {
// 		return nil, fmt.Errorf("invalid program")
// 	}

// 	vm := &VM{
// 		c: CPU{
// 			r: Registers{}
// 		},
// 		s:

// 	}
// }

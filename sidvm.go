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

type Reg uint16

type Registers struct {
	ax, bx, cx, dx, sp, ip Reg
}

type CPU struct {
	r Registers
}

type Stack struct {
}

type VM struct {
	c CPU
	s *Stack
	p *Program
}

Drones
======

At some point this will turn into a video game.  The concept is robot
combat, but instead of controlling the game directly players provide
code for a VM that controls their robot.  Currently, only the VM is
working.

VM Documentation
----------------

The Drones VM is a strictly 16-bit machine: all of its registers are
16 bits, and its memory is addressed in 16-bit words (not bytes!).

### Registers

* ip - Instruction pointer.  Tells the machine where to fetch the next
  opcode from.
* bp - Base pointer.  Points to the return address for the current
  activation record.
* sp - Stack pointer.  Points one element below the top of the stack
  (the stack grown downwards).
* a - Accumulator.  Opcodes that compute a result will generally store
  their output here.
* b - A second register, used as the source for some which operate on
  multiple values.
* p - Pointer register, used for dynamic memory loading operations.
* i - Index register, used for indexed addressing.

### Operation

The initial state of the machine puts sp and bp at the highest address
(the stack grown downwards from the top of memory) and every other
register at 0.  In a clock cycle, the machine loads the instruction at
address ip and an argument from address ip + 1, increments ip by 2,
and then executes the opcode with the argument.  Note that _all_
opcodes take up 2 16-bit words in memory, for those which don't use
arguments the content of the second word is irrelevant.  A clever
programmer or compiler may re-use those spaces in memory.  Execution
continues in this way for as long as the machine runs.  In the event
that something goes horrible wrong (divide by zero or attempt to
access out-of-bounds memory, for instance), the instruction point is
reset back to 0 and execution continues.

### Stack Frame

There are two special opcodes for dealing with function calls, call
and ret.  When call is executed, the VM jumps to its argument address.
At the same time, it pushes the address of the instruction that _would
have_ been executed next and the current base pointer to the stack,
and sets bp to point to the return address that it just pushed.  When
ret is executed, it restores the previous base pointer and jumps to
the return address.  This has the effect of dropping anything from the
stack that was pushed during the function's execution (so local
variables can be safely pushed to the stack).  Everything pushed
before beginning function execution will be just above the base
pointer, so you can access arguments pushed onto the stack by the
calling function using indexed addressing.

### Instruction Set

The instruction set of the drones VM is extremely minimal.  Most every
operation operates on a and possibly b and stores its result in a.
Loading from constants or memory can only be done into a, and from
there a set of opcodes is available to swap values into other
registers of interest.  The one exception to this rule is the lbp
instruction, which loads the base pointer into the p register.

#### Jumps
* jmp - Jumps to the argument address.
* jz - Jumps to the argument address only if a == 0.
* jnz - Jumps to the argument address only if a != 0.

#### Swaps
* sab - Swaps the a and b registers.
* sap - Swaps the a and p registers.
* sai - Swaps the a and i registers.

#### Loads
* lbp - Loads the b register into p.
* ldc - Loads a constant argument into a.
* ldm - Loads from memory at the argument address into a.
* ldp - Loads from memory at the address in p into a.
* ldi - Loads from memory at the address in (p + i) into a.

#### Stack Manipulation
* push - Pushes the content of a onto the stack.
* pop - Pops the top of the stack into a.

#### Buses
* rb - Reads from the bus number in the argument into a.
* wb - Writes from a into the bus number in the argument.

#### Function Calls
* call - Jumps to the address in the argument preparing a new stack
  frame for a function call.  Does not automatically preserve
  registers, push them manually if you want to save them.
* ret - Returns from a function called with call.

#### Arithmetic
* add - a <- a + b
* sub - a <- a - b
* mul - a <- a * b
* div - a <- a / b, b <- a % b
* sdiv - a <- a / b, b <- a % b (signed)

#### Bitwise Operations
* and - a <- a & b
* or - a <- a | b
* xor - a <- a ^ b
* shl - a <- a << b
* shr - a <- a >> b
* not - a <- ~a

#### Comparisons
* lt - a <- a < b ? 0xffff : 0
* lts - a <- a < b ? 0xffff : 0 (signed)
* le - a <- a <= b ? 0xffff : 0
* les - a <- a <= b ? 0xffff : 0 (signed)
* gt - a <- a > b ? 0xffff : 0
* gts - a <- a > b ? 0xffff : 0 (signed)
* ge - a <- a >= b ? 0xffff : 0
* ges - a <- a >= b ? 0xffff : 0 (signed)
* eq - a <- a == b ? 0xffff : 0
* neq - a <- a != b ? 0xfff : 0

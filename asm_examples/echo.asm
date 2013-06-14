start:
		rb 3
		jnz write
		jmp start
write:
		rb 2
		wb 0
		sab
		ldc 1
		wb 1
		ldc 0
		wb 3
		jmp start

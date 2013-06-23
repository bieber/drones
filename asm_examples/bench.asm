start:	ldc 12
		push
		ldc 15
		push
		call mult
		pop
		pop
		jmp start

mult:	lbp
		ldc 1
		sai
		ldi
		sab
		ldc 2
		sai
		ldi
		mul
		ret

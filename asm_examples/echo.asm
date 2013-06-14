start:	rb 3		# Check bus 3 to see if a new character is ready.
		jnz write	# Jump to write if there is one.
		jmp start	# Otherwise loop back to start.

write:	rb 2		# Read the character out of bus 2...
		wb 0		# and write it right back out to bus 0.
		ldc 1		# Load constant 1 into a...
		wb 1		# ...and write it out to bus 1 to write the character.
		ldc 0		# Load constant 0 into a...
		wb 3		# ...and write it out to bus 3 to ask for another character.
		jmp start	# Then jump back to start and do it all over again.


./main:	file format mach-o arm64


Disassembly of section __TEXT,__text:

0000000100003f84 <_main>:
100003f84: ff 43 00 d1 	sub	sp, sp, #16
100003f88: ff 0f 00 b9 	str	wzr, [sp, #12]
100003f8c: ff 0b 00 b9 	str	wzr, [sp, #8]
100003f90: e8 0b 40 b9 	ldr	w8, [sp, #8]
100003f94: 08 91 01 71 	subs	w8, w8, #100
100003f98: aa 00 00 54 	b.ge	0x100003fac <_main+0x28>
100003f9c: e8 0b 40 b9 	ldr	w8, [sp, #8]
100003fa0: 08 05 00 11 	add	w8, w8, #1
100003fa4: e8 0b 00 b9 	str	w8, [sp, #8]
100003fa8: fa ff ff 17 	b	0x100003f90 <_main+0xc>
100003fac: 00 00 80 52 	mov	w0, #0
100003fb0: ff 43 00 91 	add	sp, sp, #16
100003fb4: c0 03 5f d6 	ret

#include"textflag.h"

GLOBL ·Num(SB),NOPTR,$8
DATA ·Num+0(SB)/8, $21

GLOBL ·strdata(SB), $16
DATA ·strdata+0(SB)/8, $"hello123"

GLOBL ·Str(SB), NOPTR, $16
DATA ·Str+0(SB)/8, $·Str+16(SB)
DATA ·Str+8(SB)/8, $16
DATA ·Str+16(SB)/8, $"hello123"
DATA ·Str+24(SB)/8, $"321hello"

TEXT ·Swap(SB), $0-32
    MOVQ a+0(FP), AX 
    MOVQ b+8(FP), BX 
    MOVQ BX, ret0+16(FP)
    MOVQ AX, ret2+24(FP)
    RET

#include "textflag.h"
GLOBL ·Id(SB),$8

DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

//GLOBL ·NameDate(SB),NOPTR,$8
GLOBL ·NameData(SB),$8
DATA ·NameData(SB)/8,$"gopher"

GLOBL ·Name(SB),$16
DATA ·Name+0(SB)/8,$·NameData(SB)
DATA ·Name+8(SB)/8,$6

GLOBL ·Name1(SB),$24
DATA ·Name1+0(SB)/8,$·Name1+16(SB)
DATA ·Name1+8(SB)/8,$6
DATA ·Name1+16(SB)/8,$"gopher"

// func Swap(a, b int) (int, int)
TEXT ·Swap(SB), $0-32
    MOVQ a+0(FP),AX
    MOVQ b+8(FP),BX
    MOVQ BX,ret0+16(FP)
    MOVQ AX,ret1+24(FP)
    RET

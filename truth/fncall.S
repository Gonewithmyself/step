"".Add STEXT nosplit size=65 args=0x18 locals=0x10 funcid=0x0
	0x0000 00000 (fncall.go:5)	TEXT	"".Add(SB), NOSPLIT|ABIInternal, $16-24
	0x0000 00000 (fncall.go:5)	SUBQ	$16, SP
	0x0004 00004 (fncall.go:5)	MOVQ	BP, 8(SP)
	0x0009 00009 (fncall.go:5)	LEAQ	8(SP), BP
	0x000e 00014 (fncall.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (fncall.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (fncall.go:5)	MOVQ	$0, "".~r2+40(SP)
	0x0017 00023 (fncall.go:6)	MOVQ	"".y+32(SP), AX
	0x001c 00028 (fncall.go:6)	LEAQ	3(AX), CX
	0x0020 00032 (fncall.go:6)	MOVQ	CX, "".z(SP)
	0x0024 00036 (fncall.go:7)	MOVQ	"".x+24(SP), CX
	0x0029 00041 (fncall.go:7)	LEAQ	(AX)(CX*1), AX
	0x002d 00045 (fncall.go:7)	LEAQ	3(AX), AX
	0x0031 00049 (fncall.go:7)	MOVQ	AX, "".~r2+40(SP)
	0x0036 00054 (fncall.go:7)	MOVQ	8(SP), BP
	0x003b 00059 (fncall.go:7)	ADDQ	$16, SP
	0x003f 00063 (fncall.go:7)	NOP
	0x0040 00064 (fncall.go:7)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 c7  H...H.l$.H.l$.H.
	0x0010 44 24 28 00 00 00 00 48 8b 44 24 20 48 8d 48 03  D$(....H.D$ H.H.
	0x0020 48 89 0c 24 48 8b 4c 24 18 48 8d 04 08 48 8d 40  H..$H.L$.H...H.@
	0x0030 03 48 89 44 24 28 48 8b 6c 24 08 48 83 c4 10 90  .H.D$(H.l$.H....
	0x0040 c3                                               .
"".Sub STEXT nosplit size=25 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (fncall.go:10)	TEXT	"".Sub(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (fncall.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (fncall.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (fncall.go:10)	MOVQ	$0, "".~r2+24(SP)
	0x0009 00009 (fncall.go:11)	MOVQ	"".x+8(SP), AX
	0x000e 00014 (fncall.go:11)	SUBQ	"".y+16(SP), AX
	0x0013 00019 (fncall.go:11)	MOVQ	AX, "".~r2+24(SP)
	0x0018 00024 (fncall.go:11)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 8b 44 24 08 48 2b  H.D$.....H.D$.H+
	0x0010 44 24 10 48 89 44 24 18 c3                       D$.H.D$..
"".A STEXT size=138 args=0x18 locals=0x30 funcid=0x0
	0x0000 00000 (fncall.go:14)	TEXT	"".A(SB), ABIInternal, $48-24
	0x0000 00000 (fncall.go:14)	MOVQ	TLS, CX
	0x0009 00009 (fncall.go:14)	PCDATA	$0, $-2
	0x0009 00009 (fncall.go:14)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (fncall.go:14)	PCDATA	$0, $-1
	0x0010 00016 (fncall.go:14)	CMPQ	SP, 16(CX)
	0x0014 00020 (fncall.go:14)	PCDATA	$0, $-2
	0x0014 00020 (fncall.go:14)	JLS	128
	0x0016 00022 (fncall.go:14)	PCDATA	$0, $-1
	0x0016 00022 (fncall.go:14)	SUBQ	$48, SP
	0x001a 00026 (fncall.go:14)	MOVQ	BP, 40(SP)
	0x001f 00031 (fncall.go:14)	LEAQ	40(SP), BP
	0x0024 00036 (fncall.go:14)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (fncall.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (fncall.go:14)	MOVQ	$0, "".~r2+72(SP)
	0x002d 00045 (fncall.go:15)	MOVQ	"".y+64(SP), AX
	0x0032 00050 (fncall.go:15)	MOVQ	"".x+56(SP), CX
	0x0037 00055 (fncall.go:15)	MOVQ	CX, (SP)
	0x003b 00059 (fncall.go:15)	MOVQ	AX, 8(SP)
	0x0040 00064 (fncall.go:15)	PCDATA	$1, $0
	0x0040 00064 (fncall.go:15)	CALL	"".Add(SB)
	0x0045 00069 (fncall.go:15)	MOVQ	16(SP), AX
	0x004a 00074 (fncall.go:15)	MOVQ	AX, "".z+24(SP)
	0x004f 00079 (fncall.go:16)	MOVQ	"".y+64(SP), CX
	0x0054 00084 (fncall.go:16)	MOVQ	AX, (SP)
	0x0058 00088 (fncall.go:16)	MOVQ	CX, 8(SP)
	0x005d 00093 (fncall.go:16)	NOP
	0x0060 00096 (fncall.go:16)	CALL	"".Sub(SB)
	0x0065 00101 (fncall.go:16)	MOVQ	16(SP), AX
	0x006a 00106 (fncall.go:16)	MOVQ	AX, ""..autotmp_4+32(SP)
	0x006f 00111 (fncall.go:16)	MOVQ	AX, "".~r2+72(SP)
	0x0074 00116 (fncall.go:16)	MOVQ	40(SP), BP
	0x0079 00121 (fncall.go:16)	ADDQ	$48, SP
	0x007d 00125 (fncall.go:16)	RET
	0x007e 00126 (fncall.go:16)	NOP
	0x007e 00126 (fncall.go:14)	PCDATA	$1, $-1
	0x007e 00126 (fncall.go:14)	PCDATA	$0, $-2
	0x007e 00126 (fncall.go:14)	NOP
	0x0080 00128 (fncall.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x0085 00133 (fncall.go:14)	PCDATA	$0, $-1
	0x0085 00133 (fncall.go:14)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 6a 48 83 ec 30 48 89 6c 24 28 48  H;a.vjH..0H.l$(H
	0x0020 8d 6c 24 28 48 c7 44 24 48 00 00 00 00 48 8b 44  .l$(H.D$H....H.D
	0x0030 24 40 48 8b 4c 24 38 48 89 0c 24 48 89 44 24 08  $@H.L$8H..$H.D$.
	0x0040 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24 18 48  .....H.D$.H.D$.H
	0x0050 8b 4c 24 40 48 89 04 24 48 89 4c 24 08 0f 1f 00  .L$@H..$H.L$....
	0x0060 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24 20 48  .....H.D$.H.D$ H
	0x0070 89 44 24 48 48 8b 6c 24 28 48 83 c4 30 c3 66 90  .D$HH.l$(H..0.f.
	0x0080 e8 00 00 00 00 e9 76 ff ff ff                    ......v...
	rel 12+4 t=17 TLS+0
	rel 65+4 t=8 "".Add+0
	rel 97+4 t=8 "".Sub+0
	rel 129+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 74 72 75 74 68                                   truth
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........

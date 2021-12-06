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
"".A STEXT size=175 args=0x20 locals=0x28 funcid=0x0
	0x0000 00000 (fncall.go:14)	TEXT	"".A(SB), ABIInternal, $40-32
	0x0000 00000 (fncall.go:14)	MOVQ	TLS, CX
	0x0009 00009 (fncall.go:14)	PCDATA	$0, $-2
	0x0009 00009 (fncall.go:14)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (fncall.go:14)	PCDATA	$0, $-1
	0x0010 00016 (fncall.go:14)	CMPQ	SP, 16(CX)
	0x0014 00020 (fncall.go:14)	PCDATA	$0, $-2
	0x0014 00020 (fncall.go:14)	JLS	165
	0x001a 00026 (fncall.go:14)	PCDATA	$0, $-1
	0x001a 00026 (fncall.go:14)	SUBQ	$40, SP
	0x001e 00030 (fncall.go:14)	MOVQ	BP, 32(SP)
	0x0023 00035 (fncall.go:14)	LEAQ	32(SP), BP
	0x0028 00040 (fncall.go:14)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (fncall.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (fncall.go:14)	MOVQ	$0, "".w+64(SP)
	0x0031 00049 (fncall.go:14)	MOVQ	$0, "".z+72(SP)
	0x003a 00058 (fncall.go:15)	MOVQ	"".y+56(SP), AX
	0x003f 00063 (fncall.go:15)	MOVQ	"".x+48(SP), CX
	0x0044 00068 (fncall.go:15)	MOVQ	CX, (SP)
	0x0048 00072 (fncall.go:15)	MOVQ	AX, 8(SP)
	0x004d 00077 (fncall.go:15)	PCDATA	$1, $0
	0x004d 00077 (fncall.go:15)	CALL	"".Add(SB)
	0x0052 00082 (fncall.go:15)	MOVQ	16(SP), AX
	0x0057 00087 (fncall.go:15)	MOVQ	AX, "".c+24(SP)
	0x005c 00092 (fncall.go:16)	MOVQ	"".x+48(SP), CX
	0x0061 00097 (fncall.go:16)	MOVQ	CX, (SP)
	0x0065 00101 (fncall.go:16)	MOVQ	AX, 8(SP)
	0x006a 00106 (fncall.go:16)	CALL	"".Add(SB)
	0x006f 00111 (fncall.go:16)	MOVQ	16(SP), AX
	0x0074 00116 (fncall.go:16)	MOVQ	AX, "".z+72(SP)
	0x0079 00121 (fncall.go:17)	MOVQ	"".y+56(SP), AX
	0x007e 00126 (fncall.go:17)	MOVQ	"".x+48(SP), CX
	0x0083 00131 (fncall.go:17)	MOVQ	CX, (SP)
	0x0087 00135 (fncall.go:17)	MOVQ	AX, 8(SP)
	0x008c 00140 (fncall.go:17)	CALL	"".Sub(SB)
	0x0091 00145 (fncall.go:17)	MOVQ	16(SP), AX
	0x0096 00150 (fncall.go:17)	MOVQ	AX, "".w+64(SP)
	0x009b 00155 (fncall.go:18)	MOVQ	32(SP), BP
	0x00a0 00160 (fncall.go:18)	ADDQ	$40, SP
	0x00a4 00164 (fncall.go:18)	RET
	0x00a5 00165 (fncall.go:18)	NOP
	0x00a5 00165 (fncall.go:14)	PCDATA	$1, $-1
	0x00a5 00165 (fncall.go:14)	PCDATA	$0, $-2
	0x00a5 00165 (fncall.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x00aa 00170 (fncall.go:14)	PCDATA	$0, $-1
	0x00aa 00170 (fncall.go:14)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 8b 00 00 00 48 83 ec 28 48 89  H;a.......H..(H.
	0x0020 6c 24 20 48 8d 6c 24 20 48 c7 44 24 40 00 00 00  l$ H.l$ H.D$@...
	0x0030 00 48 c7 44 24 48 00 00 00 00 48 8b 44 24 38 48  .H.D$H....H.D$8H
	0x0040 8b 4c 24 30 48 89 0c 24 48 89 44 24 08 e8 00 00  .L$0H..$H.D$....
	0x0050 00 00 48 8b 44 24 10 48 89 44 24 18 48 8b 4c 24  ..H.D$.H.D$.H.L$
	0x0060 30 48 89 0c 24 48 89 44 24 08 e8 00 00 00 00 48  0H..$H.D$......H
	0x0070 8b 44 24 10 48 89 44 24 48 48 8b 44 24 38 48 8b  .D$.H.D$HH.D$8H.
	0x0080 4c 24 30 48 89 0c 24 48 89 44 24 08 e8 00 00 00  L$0H..$H.D$.....
	0x0090 00 48 8b 44 24 10 48 89 44 24 40 48 8b 6c 24 20  .H.D$.H.D$@H.l$ 
	0x00a0 48 83 c4 28 c3 e8 00 00 00 00 e9 51 ff ff ff     H..(.......Q...
	rel 12+4 t=17 TLS+0
	rel 78+4 t=8 "".Add+0
	rel 107+4 t=8 "".Add+0
	rel 141+4 t=8 "".Sub+0
	rel 166+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 74 72 75 74 68                                   truth
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........

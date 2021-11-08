"".Add STEXT nosplit size=38 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (add.go:3)	TEXT	"".Add(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (add.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (add.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (add.go:4)	MOVQ	"".a+8(SP), AX
	0x0005 00005 (add.go:4)	MOVQ	"".b+16(SP), CX
	0x000a 00010 (add.go:4)	MOVQ	AX, DX
	0x000d 00013 (add.go:4)	SUBQ	CX, AX
	0x0010 00016 (add.go:6)	MOVQ	CX, BX
	0x0013 00019 (add.go:6)	SUBQ	DX, CX
	0x0016 00022 (add.go:5)	CMPQ	DX, BX
	0x0019 00025 (add.go:8)	CMOVQLT	CX, AX
	0x001d 00029 (add.go:8)	ADDQ	DX, AX
	0x0020 00032 (add.go:8)	MOVQ	AX, "".~r2+24(SP)
	0x0025 00037 (add.go:5)	RET
	0x0000 48 8b 44 24 08 48 8b 4c 24 10 48 89 c2 48 29 c8  H.D$.H.L$.H..H).
	0x0010 48 89 cb 48 29 d1 48 39 da 48 0f 4c c1 48 01 d0  H..H).H9.H.L.H..
	0x0020 48 89 44 24 18 c3                                H.D$..
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 70 6b 67                                         pkg
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........

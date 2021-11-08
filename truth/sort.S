"".sort1 STEXT size=229 args=0x0 locals=0x28 funcid=0x0
	0x0000 00000 (sort.go:3)	TEXT	"".sort1(SB), ABIInternal, $40-0
	0x0000 00000 (sort.go:3)	MOVQ	TLS, CX
	0x0009 00009 (sort.go:3)	PCDATA	$0, $-2
	0x0009 00009 (sort.go:3)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (sort.go:3)	PCDATA	$0, $-1
	0x0010 00016 (sort.go:3)	CMPQ	SP, 16(CX)
	0x0014 00020 (sort.go:3)	PCDATA	$0, $-2
	0x0014 00020 (sort.go:3)	JLS	214
	0x001a 00026 (sort.go:3)	PCDATA	$0, $-1
	0x001a 00026 (sort.go:3)	SUBQ	$40, SP
	0x001e 00030 (sort.go:3)	MOVQ	BP, 32(SP)
	0x0023 00035 (sort.go:3)	LEAQ	32(SP), BP
	0x0028 00040 (sort.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (sort.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (sort.go:5)	XORL	AX, AX
	0x002a 00042 (sort.go:5)	JMP	192
	0x002f 00047 (sort.go:5)	MOVQ	AX, "".i+24(SP)
	0x0034 00052 (sort.go:6)	LEAQ	type.[9]int(SB), AX
	0x003b 00059 (sort.go:6)	MOVQ	AX, (SP)
	0x003f 00063 (sort.go:6)	PCDATA	$1, $0
	0x003f 00063 (sort.go:6)	NOP
	0x0040 00064 (sort.go:6)	CALL	runtime.newobject(SB)
	0x0045 00069 (sort.go:6)	MOVQ	8(SP), AX
	0x004a 00074 (sort.go:6)	MOVQ	$9, (AX)
	0x0051 00081 (sort.go:6)	MOVQ	$8, 8(AX)
	0x0059 00089 (sort.go:6)	MOVQ	$7, 16(AX)
	0x0061 00097 (sort.go:6)	MOVQ	$6, 24(AX)
	0x0069 00105 (sort.go:6)	MOVQ	$5, 32(AX)
	0x0071 00113 (sort.go:6)	MOVQ	$4, 40(AX)
	0x0079 00121 (sort.go:6)	MOVQ	$3, 48(AX)
	0x0081 00129 (sort.go:6)	MOVQ	$2, 56(AX)
	0x0089 00137 (sort.go:6)	MOVQ	$1, 64(AX)
	0x0091 00145 (sort.go:7)	MOVQ	AX, (SP)
	0x0095 00149 (sort.go:7)	MOVQ	$9, 8(SP)
	0x009e 00158 (sort.go:7)	MOVQ	$9, 16(SP)
	0x00a7 00167 (sort.go:7)	CALL	"".bubbleSort(SB)
	0x00ac 00172 (sort.go:5)	MOVQ	"".i+24(SP), AX
	0x00b1 00177 (sort.go:5)	INCQ	AX
	0x00b4 00180 (sort.go:5)	NOP
	0x00c0 00192 (sort.go:5)	CMPQ	AX, $10000
	0x00c6 00198 (sort.go:5)	JLT	47
	0x00cc 00204 (sort.go:5)	PCDATA	$1, $-1
	0x00cc 00204 (sort.go:5)	MOVQ	32(SP), BP
	0x00d1 00209 (sort.go:5)	ADDQ	$40, SP
	0x00d5 00213 (sort.go:5)	RET
	0x00d6 00214 (sort.go:5)	NOP
	0x00d6 00214 (sort.go:3)	PCDATA	$1, $-1
	0x00d6 00214 (sort.go:3)	PCDATA	$0, $-2
	0x00d6 00214 (sort.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x00db 00219 (sort.go:3)	PCDATA	$0, $-1
	0x00db 00219 (sort.go:3)	NOP
	0x00e0 00224 (sort.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 bc 00 00 00 48 83 ec 28 48 89  H;a.......H..(H.
	0x0020 6c 24 20 48 8d 6c 24 20 31 c0 e9 91 00 00 00 48  l$ H.l$ 1......H
	0x0030 89 44 24 18 48 8d 05 00 00 00 00 48 89 04 24 90  .D$.H......H..$.
	0x0040 e8 00 00 00 00 48 8b 44 24 08 48 c7 00 09 00 00  .....H.D$.H.....
	0x0050 00 48 c7 40 08 08 00 00 00 48 c7 40 10 07 00 00  .H.@.....H.@....
	0x0060 00 48 c7 40 18 06 00 00 00 48 c7 40 20 05 00 00  .H.@.....H.@ ...
	0x0070 00 48 c7 40 28 04 00 00 00 48 c7 40 30 03 00 00  .H.@(....H.@0...
	0x0080 00 48 c7 40 38 02 00 00 00 48 c7 40 40 01 00 00  .H.@8....H.@@...
	0x0090 00 48 89 04 24 48 c7 44 24 08 09 00 00 00 48 c7  .H..$H.D$.....H.
	0x00a0 44 24 10 09 00 00 00 e8 00 00 00 00 48 8b 44 24  D$..........H.D$
	0x00b0 18 48 ff c0 66 0f 1f 84 00 00 00 00 00 0f 1f 00  .H..f...........
	0x00c0 48 3d 10 27 00 00 0f 8c 63 ff ff ff 48 8b 6c 24  H=.'....c...H.l$
	0x00d0 20 48 83 c4 28 c3 e8 00 00 00 00 0f 1f 44 00 00   H..(........D..
	0x00e0 e9 1b ff ff ff                                   .....
	rel 12+4 t=17 TLS+0
	rel 55+4 t=16 type.[9]int+0
	rel 65+4 t=8 runtime.newobject+0
	rel 168+4 t=8 "".bubbleSort+0
	rel 215+4 t=8 runtime.morestack_noctxt+0
"".sort2 STEXT size=237 args=0x0 locals=0x70 funcid=0x0
	0x0000 00000 (sort.go:11)	TEXT	"".sort2(SB), ABIInternal, $112-0
	0x0000 00000 (sort.go:11)	MOVQ	TLS, CX
	0x0009 00009 (sort.go:11)	PCDATA	$0, $-2
	0x0009 00009 (sort.go:11)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (sort.go:11)	PCDATA	$0, $-1
	0x0010 00016 (sort.go:11)	CMPQ	SP, 16(CX)
	0x0014 00020 (sort.go:11)	PCDATA	$0, $-2
	0x0014 00020 (sort.go:11)	JLS	227
	0x001a 00026 (sort.go:11)	PCDATA	$0, $-1
	0x001a 00026 (sort.go:11)	SUBQ	$112, SP
	0x001e 00030 (sort.go:11)	MOVQ	BP, 104(SP)
	0x0023 00035 (sort.go:11)	LEAQ	104(SP), BP
	0x0028 00040 (sort.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (sort.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (sort.go:11)	XORL	AX, AX
	0x002a 00042 (sort.go:12)	JMP	205
	0x002f 00047 (sort.go:12)	MOVQ	AX, "".i+24(SP)
	0x0034 00052 (sort.go:13)	MOVQ	$0, ""..autotmp_4+32(SP)
	0x003d 00061 (sort.go:13)	XORPS	X0, X0
	0x0040 00064 (sort.go:13)	MOVUPS	X0, ""..autotmp_4+40(SP)
	0x0045 00069 (sort.go:13)	MOVUPS	X0, ""..autotmp_4+56(SP)
	0x004a 00074 (sort.go:13)	MOVUPS	X0, ""..autotmp_4+72(SP)
	0x004f 00079 (sort.go:13)	MOVUPS	X0, ""..autotmp_4+88(SP)
	0x0054 00084 (sort.go:13)	MOVQ	$9, ""..autotmp_4+32(SP)
	0x005d 00093 (sort.go:13)	MOVQ	$8, ""..autotmp_4+40(SP)
	0x0066 00102 (sort.go:13)	MOVQ	$7, ""..autotmp_4+48(SP)
	0x006f 00111 (sort.go:13)	MOVQ	$6, ""..autotmp_4+56(SP)
	0x0078 00120 (sort.go:13)	MOVQ	$5, ""..autotmp_4+64(SP)
	0x0081 00129 (sort.go:13)	MOVQ	$4, ""..autotmp_4+72(SP)
	0x008a 00138 (sort.go:13)	MOVQ	$3, ""..autotmp_4+80(SP)
	0x0093 00147 (sort.go:13)	MOVQ	$2, ""..autotmp_4+88(SP)
	0x009c 00156 (sort.go:13)	MOVQ	$1, ""..autotmp_4+96(SP)
	0x00a5 00165 (sort.go:14)	LEAQ	""..autotmp_4+32(SP), AX
	0x00aa 00170 (sort.go:14)	MOVQ	AX, (SP)
	0x00ae 00174 (sort.go:14)	MOVQ	$9, 8(SP)
	0x00b7 00183 (sort.go:14)	MOVQ	$9, 16(SP)
	0x00c0 00192 (sort.go:14)	PCDATA	$1, $0
	0x00c0 00192 (sort.go:14)	CALL	"".bubbleSort(SB)
	0x00c5 00197 (sort.go:12)	MOVQ	"".i+24(SP), AX
	0x00ca 00202 (sort.go:12)	INCQ	AX
	0x00cd 00205 (sort.go:12)	CMPQ	AX, $10000
	0x00d3 00211 (sort.go:12)	JLT	47
	0x00d9 00217 (sort.go:12)	PCDATA	$1, $-1
	0x00d9 00217 (sort.go:12)	MOVQ	104(SP), BP
	0x00de 00222 (sort.go:12)	ADDQ	$112, SP
	0x00e2 00226 (sort.go:12)	RET
	0x00e3 00227 (sort.go:12)	NOP
	0x00e3 00227 (sort.go:11)	PCDATA	$1, $-1
	0x00e3 00227 (sort.go:11)	PCDATA	$0, $-2
	0x00e3 00227 (sort.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x00e8 00232 (sort.go:11)	PCDATA	$0, $-1
	0x00e8 00232 (sort.go:11)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 c9 00 00 00 48 83 ec 70 48 89  H;a.......H..pH.
	0x0020 6c 24 68 48 8d 6c 24 68 31 c0 e9 9e 00 00 00 48  l$hH.l$h1......H
	0x0030 89 44 24 18 48 c7 44 24 20 00 00 00 00 0f 57 c0  .D$.H.D$ .....W.
	0x0040 0f 11 44 24 28 0f 11 44 24 38 0f 11 44 24 48 0f  ..D$(..D$8..D$H.
	0x0050 11 44 24 58 48 c7 44 24 20 09 00 00 00 48 c7 44  .D$XH.D$ ....H.D
	0x0060 24 28 08 00 00 00 48 c7 44 24 30 07 00 00 00 48  $(....H.D$0....H
	0x0070 c7 44 24 38 06 00 00 00 48 c7 44 24 40 05 00 00  .D$8....H.D$@...
	0x0080 00 48 c7 44 24 48 04 00 00 00 48 c7 44 24 50 03  .H.D$H....H.D$P.
	0x0090 00 00 00 48 c7 44 24 58 02 00 00 00 48 c7 44 24  ...H.D$X....H.D$
	0x00a0 60 01 00 00 00 48 8d 44 24 20 48 89 04 24 48 c7  `....H.D$ H..$H.
	0x00b0 44 24 08 09 00 00 00 48 c7 44 24 10 09 00 00 00  D$.....H.D$.....
	0x00c0 e8 00 00 00 00 48 8b 44 24 18 48 ff c0 48 3d 10  .....H.D$.H..H=.
	0x00d0 27 00 00 0f 8c 56 ff ff ff 48 8b 6c 24 68 48 83  '....V...H.l$hH.
	0x00e0 c4 70 c3 e8 00 00 00 00 e9 13 ff ff ff           .p...........
	rel 12+4 t=17 TLS+0
	rel 193+4 t=8 "".bubbleSort+0
	rel 228+4 t=8 runtime.morestack_noctxt+0
"".bubbleSort STEXT nosplit size=145 args=0x18 locals=0x18 funcid=0x0
	0x0000 00000 (sort.go:18)	TEXT	"".bubbleSort(SB), NOSPLIT|ABIInternal, $24-24
	0x0000 00000 (sort.go:18)	SUBQ	$24, SP
	0x0004 00004 (sort.go:18)	MOVQ	BP, 16(SP)
	0x0009 00009 (sort.go:18)	LEAQ	16(SP), BP
	0x000e 00014 (sort.go:18)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x000e 00014 (sort.go:18)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x000e 00014 (sort.go:19)	MOVQ	"".arr+32(SP), DX
	0x0013 00019 (sort.go:20)	MOVQ	"".arr+40(SP), BX
	0x0018 00024 (sort.go:20)	XORL	AX, AX
	0x001a 00026 (sort.go:20)	JMP	95
	0x001c 00028 (sort.go:21)	MOVQ	BX, CX
	0x001f 00031 (sort.go:21)	MOVQ	SI, BX
	0x0022 00034 (sort.go:21)	MOVQ	BX, SI
	0x0025 00037 (sort.go:21)	SUBQ	AX, BX
	0x0028 00040 (sort.go:21)	DECQ	BX
	0x002b 00043 (sort.go:21)	CMPQ	CX, BX
	0x002e 00046 (sort.go:21)	JGE	89
	0x0030 00048 (sort.go:22)	LEAQ	1(CX), BX
	0x0034 00052 (sort.go:22)	CMPQ	SI, BX
	0x0037 00055 (sort.go:22)	JLS	133
	0x0039 00057 (sort.go:22)	MOVQ	8(DX)(CX*8), DI
	0x003e 00062 (sort.go:22)	NOP
	0x0040 00064 (sort.go:22)	CMPQ	SI, CX
	0x0043 00067 (sort.go:22)	JLS	118
	0x0045 00069 (sort.go:22)	MOVQ	(DX)(CX*8), R8
	0x0049 00073 (sort.go:22)	CMPQ	DI, R8
	0x004c 00076 (sort.go:22)	JLE	28
	0x004e 00078 (sort.go:23)	MOVQ	DI, (DX)(CX*8)
	0x0052 00082 (sort.go:23)	MOVQ	R8, 8(DX)(CX*8)
	0x0057 00087 (sort.go:23)	JMP	28
	0x0059 00089 (sort.go:20)	INCQ	AX
	0x005c 00092 (sort.go:20)	MOVQ	SI, BX
	0x005f 00095 (sort.go:20)	LEAQ	-1(BX), SI
	0x0063 00099 (sort.go:20)	CMPQ	AX, SI
	0x0066 00102 (sort.go:20)	JGE	108
	0x0068 00104 (sort.go:20)	XORL	CX, CX
	0x006a 00106 (sort.go:21)	JMP	34
	0x006c 00108 (sort.go:20)	MOVQ	16(SP), BP
	0x0071 00113 (sort.go:20)	ADDQ	$24, SP
	0x0075 00117 (sort.go:20)	RET
	0x0076 00118 (sort.go:22)	MOVQ	CX, AX
	0x0079 00121 (sort.go:22)	MOVQ	SI, CX
	0x007c 00124 (sort.go:22)	PCDATA	$1, $1
	0x007c 00124 (sort.go:22)	NOP
	0x0080 00128 (sort.go:22)	CALL	runtime.panicIndex(SB)
	0x0085 00133 (sort.go:22)	MOVQ	BX, AX
	0x0088 00136 (sort.go:22)	MOVQ	SI, CX
	0x008b 00139 (sort.go:22)	CALL	runtime.panicIndex(SB)
	0x0090 00144 (sort.go:22)	XCHGL	AX, AX
	0x0000 48 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 8b  H...H.l$.H.l$.H.
	0x0010 54 24 20 48 8b 5c 24 28 31 c0 eb 43 48 89 d9 48  T$ H.\$(1..CH..H
	0x0020 89 f3 48 89 de 48 29 c3 48 ff cb 48 39 d9 7d 29  ..H..H).H..H9.})
	0x0030 48 8d 59 01 48 39 de 76 4c 48 8b 7c ca 08 66 90  H.Y.H9.vLH.|..f.
	0x0040 48 39 ce 76 31 4c 8b 04 ca 4c 39 c7 7e ce 48 89  H9.v1L...L9.~.H.
	0x0050 3c ca 4c 89 44 ca 08 eb c3 48 ff c0 48 89 f3 48  <.L.D....H..H..H
	0x0060 8d 73 ff 48 39 f0 7d 04 31 c9 eb b6 48 8b 6c 24  .s.H9.}.1...H.l$
	0x0070 10 48 83 c4 18 c3 48 89 c8 48 89 f1 0f 1f 40 00  .H....H..H....@.
	0x0080 e8 00 00 00 00 48 89 d8 48 89 f1 e8 00 00 00 00  .....H..H.......
	0x0090 90                                               .
	rel 129+4 t=8 runtime.panicIndex+0
	rel 140+4 t=8 runtime.panicIndex+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 74 72 75 74 68                                   truth
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]int- SRODATA dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..eqfunc72 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 48 00 00 00 00 00 00 00  ........H.......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[9]int- SRODATA dupok size=10
	0x0000 00 00 07 2a 5b 39 5d 69 6e 74                    ...*[9]int
type.*[9]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ce 17 bf 1b 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[9]int-+0
	rel 48+8 t=1 type.[9]int+0
runtime.gcbits. SRODATA dupok size=0
type.[9]int SRODATA dupok size=72
	0x0000 48 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  H...............
	0x0010 2f 88 ea 48 0a 08 08 11 00 00 00 00 00 00 00 00  /..H............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 09 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc72+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[9]int-+0
	rel 44+4 t=6 type.*[9]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........

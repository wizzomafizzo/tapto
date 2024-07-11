//go:build linux

package mister

import "github.com/bendahl/uinput"

var KeyboardMap = map[string]int{
	"esc":       1,
	"1":         2,
	"2":         3,
	"3":         4,
	"4":         5,
	"5":         6,
	"6":         7,
	"7":         8,
	"8":         9,
	"9":         10,
	"0":         11,
	"-":         12,
	"=":         13,
	"backspace": 14,
	"tab":       15,
	"q":         16,
	"w":         17,
	"e":         18,
	"r":         19,
	"t":         20,
	"y":         21,
	"u":         22,
	"i":         23,
	"o":         24,
	"p":         25,
	"[":         26,
	"]":         27,
	"enter":     28,
	"lctrl":     29,
	"a":         30,
	"s":         31,
	"d":         32,
	"f":         33,
	"g":         34,
	"h":         35,
	"j":         36,
	"k":         37,
	"l":         38,
	";":         39,
	"'":         40,
	"`":         41,
	"lshift":    42,
	"\\":        43,
	"backslash": 43,
	"z":         44,
	"x":         45,
	"c":         46,
	"v":         47,
	"b":         48,
	"n":         49,
	"m":         50,
	",":         51,
	".":         52,
	"/":         53,
	"rshift":    54,
	"lalt":      56,
	" ":         57,
	"space":     57,
	"caps":      58,
	"f1":        59,
	"f2":        60,
	"f3":        61,
	"f4":        62,
	"f5":        63,
	"f6":        64,
	"f7":        65,
	"f8":        66,
	"f9":        67,
	"f10":       68,
	"num":       69,
	"scroll":    70,
	"f11":       87,
	"f12":       88,
	"home":      102,
	"up":        103,
	"pgup":      104,
	"left":      105,
	"right":     106,
	"end":       107,
	"down":      108,
	"pgdn":      109,
	"ins":       110,
	"del":       111,
	"volup":     115,
	"voldn":     114,
	// shift
	"~":  -41,
	"!":  -2,
	"@":  -3,
	"#":  -4,
	"$":  -5,
	"%":  -6,
	"^":  -7,
	"&":  -8,
	"*":  -9,
	"(":  -10,
	")":  -11,
	"_":  -12,
	"+":  -13,
	"{":  -26,
	"}":  -27,
	"|":  -43,
	":":  -39,
	"\"": -40,
	"<":  -51,
	">":  -52,
	"?":  -53,
	"Q":  -16,
	"W":  -17,
	"E":  -18,
	"R":  -19,
	"T":  -20,
	"Y":  -21,
	"U":  -22,
	"I":  -23,
	"O":  -24,
	"P":  -25,
	"A":  -30,
	"S":  -31,
	"D":  -32,
	"F":  -33,
	"G":  -34,
	"H":  -35,
	"J":  -36,
	"K":  -37,
	"L":  -38,
	"Z":  -44,
	"X":  -45,
	"C":  -46,
	"V":  -47,
	"B":  -48,
	"N":  -49,
	"M":  -50,
}

var GamepadMap = map[string]int{
	"^":        uinput.ButtonDpadUp,
	"{up}":     uinput.ButtonDpadUp,
	"V":        uinput.ButtonDpadDown,
	"{down}":   uinput.ButtonDpadDown,
	"<":        uinput.ButtonDpadLeft,
	"{left}":   uinput.ButtonDpadLeft,
	">":        uinput.ButtonDpadRight,
	"{right}":  uinput.ButtonDpadRight,
	"A":        uinput.ButtonEast,
	"a":        uinput.ButtonEast,
	"B":        uinput.ButtonSouth,
	"b":        uinput.ButtonSouth,
	"X":        uinput.ButtonNorth,
	"x":        uinput.ButtonNorth,
	"Y":        uinput.ButtonWest,
	"y":        uinput.ButtonWest,
	"{start}":  uinput.ButtonStart,
	"{select}": uinput.ButtonSelect,
	"{menu}":   uinput.ButtonMode,
	"L":        uinput.ButtonBumperLeft,
	"l":        uinput.ButtonBumperLeft,
	"{l1}":     uinput.ButtonBumperLeft,
	"R":        uinput.ButtonBumperRight,
	"r":        uinput.ButtonBumperRight,
	"{r1}":     uinput.ButtonBumperRight,
	"{l2}":     uinput.ButtonTriggerLeft,
	"{r2}":     uinput.ButtonTriggerRight,
}

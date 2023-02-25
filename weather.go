package main

type Weather int8

const (
	WM00 Weather = iota
	WM01
	WM02
	WM03
	WM04
	WM05
	WM06
	WM07
	WM08
	WM09
	WM10
	WM11
	WM12
	WM13
	WM14
	WM15
	WM16
	WM17
	WM18
	WM19
	WM20
	WM21
	WM22
	WM23
	WM24
	WM25
	WM26
	WM27
	WM28
	WM29
	WM30
	WM31
	WM32
	WM33
	WM34
	WM35
	WM36
	WM37
	WM38
	WM39
	WM40
	WM41
	WM42
	WM43
	WM44
	WM45
	WM46
	WM47
	WM48
	WM49
	WM50
	WM51
	WM52
	WM53
	WM54
	WM55
	WM56
	WM57
	WM58
	WM59
	WM60
	WM61
	WM62
	WM63
	WM64
	WM65
	WM66
	WM67
	WM68
	WM69
	WM70
	WM71
	WM72
	WM73
	WM74
	WM75
	WM76
	WM77
	WM78
	WM79
	WM80
	WM81
	WM82
	WM83
	WM84
	WM85
	WM86
	WM87
	WM88
	WM89
	WM90
	WM91
	WM92
	WM93
	WM94
	WM95
	WM96
	WM97
	WM98
	WM99
)

func (w Weather) GetString() string {
	switch {
	// case WM00, WM01, WM02, WM03:
	case WM00 <= w && w <= WM03:
		return "晴れ"
	case WM04 <= w && w <= WM19:
		return "曇り"
	case WM20 <= w && w <= WM29:
		return "一時雨"
	case WM30 <= w && w <= WM39:
		return "雪"
	case WM40 <= w && w <= WM49:
		return "霧"
	case WM50 <= w && w <= WM59:
		return "小雨"
	case WM60 <= w && w <= WM69:
		return "雨"
	case WM70 <= w && w <= WM79:
		return "大雨"
	case WM80 <= w && w <= WM99:
		return "雷雨"
	default:
		return "不明な天気コード"
	}
}

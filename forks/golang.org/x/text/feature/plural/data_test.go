// Code generated by running "go generate" in github.com/ie310mu/ie310go/forks/golang.org/x/text. DO NOT EDIT.

package plural

type pluralTest struct {
	locales string
	form    int
	integer []string
	decimal []string
}

var ordinalTests = []pluralTest{ // 66 elements
	0:  {locales: "af am ar bg bs ce cs da de dsb el es et eu fa fi fy gl gsw he hr hsb id in is iw ja km kn ko ky lt lv ml mn my nb nl pa pl prg ps pt root ru sd sh si sk sl sr sw ta te th tr ur uz yue zh zu", form: 0, integer: []string{"0~15", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	1:  {locales: "sv", form: 2, integer: []string{"1", "2", "21", "22", "31", "32", "41", "42", "51", "52", "61", "62", "71", "72", "81", "82", "101", "1001"}, decimal: []string(nil)},
	2:  {locales: "sv", form: 0, integer: []string{"0", "3~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	3:  {locales: "fil fr ga hy lo mo ms ro tl vi", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	4:  {locales: "fil fr ga hy lo mo ms ro tl vi", form: 0, integer: []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	5:  {locales: "hu", form: 2, integer: []string{"1", "5"}, decimal: []string(nil)},
	6:  {locales: "hu", form: 0, integer: []string{"0", "2~4", "6~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	7:  {locales: "ne", form: 2, integer: []string{"1~4"}, decimal: []string(nil)},
	8:  {locales: "ne", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	9:  {locales: "be", form: 4, integer: []string{"2", "3", "22", "23", "32", "33", "42", "43", "52", "53", "62", "63", "72", "73", "82", "83", "102", "1002"}, decimal: []string(nil)},
	10: {locales: "be", form: 0, integer: []string{"0", "1", "4~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	11: {locales: "uk", form: 4, integer: []string{"3", "23", "33", "43", "53", "63", "73", "83", "103", "1003"}, decimal: []string(nil)},
	12: {locales: "uk", form: 0, integer: []string{"0~2", "4~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	13: {locales: "tk", form: 4, integer: []string{"6", "9", "10", "16", "19", "26", "29", "36", "39", "106", "1006"}, decimal: []string(nil)},
	14: {locales: "tk", form: 0, integer: []string{"0~5", "7", "8", "11~15", "17", "18", "20", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	15: {locales: "kk", form: 5, integer: []string{"6", "9", "10", "16", "19", "20", "26", "29", "30", "36", "39", "40", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	16: {locales: "kk", form: 0, integer: []string{"0~5", "7", "8", "11~15", "17", "18", "21", "101", "1001"}, decimal: []string(nil)},
	17: {locales: "it", form: 5, integer: []string{"8", "11", "80", "800"}, decimal: []string(nil)},
	18: {locales: "it", form: 0, integer: []string{"0~7", "9", "10", "12~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	19: {locales: "ka", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	20: {locales: "ka", form: 5, integer: []string{"0", "2~16", "102", "1002"}, decimal: []string(nil)},
	21: {locales: "ka", form: 0, integer: []string{"21~36", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	22: {locales: "sq", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	23: {locales: "sq", form: 5, integer: []string{"4", "24", "34", "44", "54", "64", "74", "84", "104", "1004"}, decimal: []string(nil)},
	24: {locales: "sq", form: 0, integer: []string{"0", "2", "3", "5~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	25: {locales: "en", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string(nil)},
	26: {locales: "en", form: 3, integer: []string{"2", "22", "32", "42", "52", "62", "72", "82", "102", "1002"}, decimal: []string(nil)},
	27: {locales: "en", form: 4, integer: []string{"3", "23", "33", "43", "53", "63", "73", "83", "103", "1003"}, decimal: []string(nil)},
	28: {locales: "en", form: 0, integer: []string{"0", "4~18", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	29: {locales: "mr", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	30: {locales: "mr", form: 3, integer: []string{"2", "3"}, decimal: []string(nil)},
	31: {locales: "mr", form: 4, integer: []string{"4"}, decimal: []string(nil)},
	32: {locales: "mr", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	33: {locales: "ca", form: 2, integer: []string{"1", "3"}, decimal: []string(nil)},
	34: {locales: "ca", form: 3, integer: []string{"2"}, decimal: []string(nil)},
	35: {locales: "ca", form: 4, integer: []string{"4"}, decimal: []string(nil)},
	36: {locales: "ca", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	37: {locales: "mk", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string(nil)},
	38: {locales: "mk", form: 3, integer: []string{"2", "22", "32", "42", "52", "62", "72", "82", "102", "1002"}, decimal: []string(nil)},
	39: {locales: "mk", form: 5, integer: []string{"7", "8", "27", "28", "37", "38", "47", "48", "57", "58", "67", "68", "77", "78", "87", "88", "107", "1007"}, decimal: []string(nil)},
	40: {locales: "mk", form: 0, integer: []string{"0", "3~6", "9~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	41: {locales: "az", form: 2, integer: []string{"1", "2", "5", "7", "8", "11", "12", "15", "17", "18", "20~22", "25", "101", "1001"}, decimal: []string(nil)},
	42: {locales: "az", form: 4, integer: []string{"3", "4", "13", "14", "23", "24", "33", "34", "43", "44", "53", "54", "63", "64", "73", "74", "100", "1003"}, decimal: []string(nil)},
	43: {locales: "az", form: 5, integer: []string{"0", "6", "16", "26", "36", "40", "46", "56", "106", "1006"}, decimal: []string(nil)},
	44: {locales: "az", form: 0, integer: []string{"9", "10", "19", "29", "30", "39", "49", "59", "69", "79", "109", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	45: {locales: "gu hi", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	46: {locales: "gu hi", form: 3, integer: []string{"2", "3"}, decimal: []string(nil)},
	47: {locales: "gu hi", form: 4, integer: []string{"4"}, decimal: []string(nil)},
	48: {locales: "gu hi", form: 5, integer: []string{"6"}, decimal: []string(nil)},
	49: {locales: "gu hi", form: 0, integer: []string{"0", "5", "7~20", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	50: {locales: "as bn", form: 2, integer: []string{"1", "5", "7~10"}, decimal: []string(nil)},
	51: {locales: "as bn", form: 3, integer: []string{"2", "3"}, decimal: []string(nil)},
	52: {locales: "as bn", form: 4, integer: []string{"4"}, decimal: []string(nil)},
	53: {locales: "as bn", form: 5, integer: []string{"6"}, decimal: []string(nil)},
	54: {locales: "as bn", form: 0, integer: []string{"0", "11~25", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	55: {locales: "or", form: 2, integer: []string{"1", "5", "7~9"}, decimal: []string(nil)},
	56: {locales: "or", form: 3, integer: []string{"2", "3"}, decimal: []string(nil)},
	57: {locales: "or", form: 4, integer: []string{"4"}, decimal: []string(nil)},
	58: {locales: "or", form: 5, integer: []string{"6"}, decimal: []string(nil)},
	59: {locales: "or", form: 0, integer: []string{"0", "10~24", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	60: {locales: "cy", form: 1, integer: []string{"0", "7~9"}, decimal: []string(nil)},
	61: {locales: "cy", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	62: {locales: "cy", form: 3, integer: []string{"2"}, decimal: []string(nil)},
	63: {locales: "cy", form: 4, integer: []string{"3", "4"}, decimal: []string(nil)},
	64: {locales: "cy", form: 5, integer: []string{"5", "6"}, decimal: []string(nil)},
	65: {locales: "cy", form: 0, integer: []string{"10~25", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
} // Size: 4776 bytes

var cardinalTests = []pluralTest{ // 113 elements
	0:   {locales: "bm bo dz id ig ii in ja jbo jv jw kde kea km ko lkt lo ms my nqo root sah ses sg th to vi wo yo yue zh", form: 0, integer: []string{"0~15", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	1:   {locales: "am as bn fa gu hi kn mr zu", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0~1.0", "0.00~0.04"}},
	2:   {locales: "am as bn fa gu hi kn mr zu", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"1.1~2.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	3:   {locales: "ff fr hy kab", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0~1.5"}},
	4:   {locales: "ff fr hy kab", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	5:   {locales: "pt", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0~1.5"}},
	6:   {locales: "pt", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	7:   {locales: "ast ca de en et fi fy gl io it ji nl pt_PT sv sw ur yi", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	8:   {locales: "ast ca de en et fi fy gl io it ji nl pt_PT sv sw ur yi", form: 0, integer: []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	9:   {locales: "si", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0", "0.1", "1.0", "0.00", "0.01", "1.00", "0.000", "0.001", "1.000", "0.0000", "0.0001", "1.0000"}},
	10:  {locales: "si", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.2~0.9", "1.1~1.8", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	11:  {locales: "ak bh guw ln mg nso pa ti wa", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0", "1.0", "0.00", "1.00", "0.000", "1.000", "0.0000", "1.0000"}},
	12:  {locales: "ak bh guw ln mg nso pa ti wa", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	13:  {locales: "tzm", form: 2, integer: []string{"0", "1", "11~24"}, decimal: []string{"0.0", "1.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "19.0", "20.0", "21.0", "22.0", "23.0", "24.0"}},
	14:  {locales: "tzm", form: 0, integer: []string{"2~10", "100~106", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	15:  {locales: "af asa az bem bez bg brx ce cgg chr ckb dv ee el eo es eu fo fur gsw ha haw hu jgo jmc ka kaj kcg kk kkj kl ks ksb ku ky lb lg mas mgo ml mn nah nb nd ne nn nnh no nr ny nyn om or os pap ps rm rof rwk saq sd sdh seh sn so sq ss ssy st syr ta te teo tig tk tn tr ts ug uz ve vo vun wae xh xog", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	16:  {locales: "af asa az bem bez bg brx ce cgg chr ckb dv ee el eo es eu fo fur gsw ha haw hu jgo jmc ka kaj kcg kk kkj kl ks ksb ku ky lb lg mas mgo ml mn nah nb nd ne nn nnh no nr ny nyn om or os pap ps rm rof rwk saq sd sdh seh sn so sq ss ssy st syr ta te teo tig tk tn tr ts ug uz ve vo vun wae xh xog", form: 0, integer: []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	17:  {locales: "da", form: 2, integer: []string{"1"}, decimal: []string{"0.1~1.6"}},
	18:  {locales: "da", form: 0, integer: []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "2.0~3.4", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	19:  {locales: "is", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string{"0.1~1.6", "10.1", "100.1", "1000.1"}},
	20:  {locales: "is", form: 0, integer: []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	21:  {locales: "mk", form: 2, integer: []string{"1", "11", "21", "31", "41", "51", "61", "71", "101", "1001"}, decimal: []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"}},
	22:  {locales: "mk", form: 0, integer: []string{"0", "2~10", "12~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "0.2~1.0", "1.2~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	23:  {locales: "fil tl", form: 2, integer: []string{"0~3", "5", "7", "8", "10~13", "15", "17", "18", "20", "21", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~0.3", "0.5", "0.7", "0.8", "1.0~1.3", "1.5", "1.7", "1.8", "2.0", "2.1", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	24:  {locales: "fil tl", form: 0, integer: []string{"4", "6", "9", "14", "16", "19", "24", "26", "104", "1004"}, decimal: []string{"0.4", "0.6", "0.9", "1.4", "1.6", "1.9", "2.4", "2.6", "10.4", "100.4", "1000.4"}},
	25:  {locales: "lv prg", form: 1, integer: []string{"0", "10~20", "30", "40", "50", "60", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "10.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	26:  {locales: "lv prg", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string{"0.1", "1.0", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"}},
	27:  {locales: "lv prg", form: 0, integer: []string{"2~9", "22~29", "102", "1002"}, decimal: []string{"0.2~0.9", "1.2~1.9", "10.2", "100.2", "1000.2"}},
	28:  {locales: "lag", form: 1, integer: []string{"0"}, decimal: []string{"0.0", "0.00", "0.000", "0.0000"}},
	29:  {locales: "lag", form: 2, integer: []string{"1"}, decimal: []string{"0.1~1.6"}},
	30:  {locales: "lag", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	31:  {locales: "ksh", form: 1, integer: []string{"0"}, decimal: []string{"0.0", "0.00", "0.000", "0.0000"}},
	32:  {locales: "ksh", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	33:  {locales: "ksh", form: 0, integer: []string{"2~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	34:  {locales: "iu kw naq se sma smi smj smn sms", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	35:  {locales: "iu kw naq se sma smi smj smn sms", form: 3, integer: []string{"2"}, decimal: []string{"2.0", "2.00", "2.000", "2.0000"}},
	36:  {locales: "iu kw naq se sma smi smj smn sms", form: 0, integer: []string{"0", "3~17", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	37:  {locales: "shi", form: 2, integer: []string{"0", "1"}, decimal: []string{"0.0~1.0", "0.00~0.04"}},
	38:  {locales: "shi", form: 4, integer: []string{"2~10"}, decimal: []string{"2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "2.00", "3.00", "4.00", "5.00", "6.00", "7.00", "8.00"}},
	39:  {locales: "shi", form: 0, integer: []string{"11~26", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"1.1~1.9", "2.1~2.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	40:  {locales: "mo ro", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	41:  {locales: "mo ro", form: 4, integer: []string{"0", "2~16", "101", "1001"}, decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	42:  {locales: "mo ro", form: 0, integer: []string{"20~35", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	43:  {locales: "bs hr sh sr", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"}},
	44:  {locales: "bs hr sh sr", form: 4, integer: []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"}, decimal: []string{"0.2~0.4", "1.2~1.4", "2.2~2.4", "3.2~3.4", "4.2~4.4", "5.2", "10.2", "100.2", "1000.2"}},
	45:  {locales: "bs hr sh sr", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "0.5~1.0", "1.5~2.0", "2.5~2.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	46:  {locales: "gd", form: 2, integer: []string{"1", "11"}, decimal: []string{"1.0", "11.0", "1.00", "11.00", "1.000", "11.000", "1.0000"}},
	47:  {locales: "gd", form: 3, integer: []string{"2", "12"}, decimal: []string{"2.0", "12.0", "2.00", "12.00", "2.000", "12.000", "2.0000"}},
	48:  {locales: "gd", form: 4, integer: []string{"3~10", "13~19"}, decimal: []string{"3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "19.0", "3.00"}},
	49:  {locales: "gd", form: 0, integer: []string{"0", "20~34", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~0.9", "1.1~1.6", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	50:  {locales: "sl", form: 2, integer: []string{"1", "101", "201", "301", "401", "501", "601", "701", "1001"}, decimal: []string(nil)},
	51:  {locales: "sl", form: 3, integer: []string{"2", "102", "202", "302", "402", "502", "602", "702", "1002"}, decimal: []string(nil)},
	52:  {locales: "sl", form: 4, integer: []string{"3", "4", "103", "104", "203", "204", "303", "304", "403", "404", "503", "504", "603", "604", "703", "704", "1003"}, decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	53:  {locales: "sl", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	54:  {locales: "dsb hsb", form: 2, integer: []string{"1", "101", "201", "301", "401", "501", "601", "701", "1001"}, decimal: []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"}},
	55:  {locales: "dsb hsb", form: 3, integer: []string{"2", "102", "202", "302", "402", "502", "602", "702", "1002"}, decimal: []string{"0.2", "1.2", "2.2", "3.2", "4.2", "5.2", "6.2", "7.2", "10.2", "100.2", "1000.2"}},
	56:  {locales: "dsb hsb", form: 4, integer: []string{"3", "4", "103", "104", "203", "204", "303", "304", "403", "404", "503", "504", "603", "604", "703", "704", "1003"}, decimal: []string{"0.3", "0.4", "1.3", "1.4", "2.3", "2.4", "3.3", "3.4", "4.3", "4.4", "5.3", "5.4", "6.3", "6.4", "7.3", "7.4", "10.3", "100.3", "1000.3"}},
	57:  {locales: "dsb hsb", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "0.5~1.0", "1.5~2.0", "2.5~2.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	58:  {locales: "he iw", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	59:  {locales: "he iw", form: 3, integer: []string{"2"}, decimal: []string(nil)},
	60:  {locales: "he iw", form: 5, integer: []string{"20", "30", "40", "50", "60", "70", "80", "90", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	61:  {locales: "he iw", form: 0, integer: []string{"0", "3~17", "101", "1001"}, decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	62:  {locales: "cs sk", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	63:  {locales: "cs sk", form: 4, integer: []string{"2~4"}, decimal: []string(nil)},
	64:  {locales: "cs sk", form: 5, integer: []string(nil), decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	65:  {locales: "cs sk", form: 0, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	66:  {locales: "pl", form: 2, integer: []string{"1"}, decimal: []string(nil)},
	67:  {locales: "pl", form: 4, integer: []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"}, decimal: []string(nil)},
	68:  {locales: "pl", form: 5, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	69:  {locales: "pl", form: 0, integer: []string(nil), decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	70:  {locales: "be", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "71.0", "81.0", "101.0", "1001.0"}},
	71:  {locales: "be", form: 4, integer: []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"}, decimal: []string{"2.0", "3.0", "4.0", "22.0", "23.0", "24.0", "32.0", "33.0", "102.0", "1002.0"}},
	72:  {locales: "be", form: 5, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "11.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	73:  {locales: "be", form: 0, integer: []string(nil), decimal: []string{"0.1~0.9", "1.1~1.7", "10.1", "100.1", "1000.1"}},
	74:  {locales: "lt", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "71.0", "81.0", "101.0", "1001.0"}},
	75:  {locales: "lt", form: 4, integer: []string{"2~9", "22~29", "102", "1002"}, decimal: []string{"2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "22.0", "102.0", "1002.0"}},
	76:  {locales: "lt", form: 5, integer: []string(nil), decimal: []string{"0.1~0.9", "1.1~1.7", "10.1", "100.1", "1000.1"}},
	77:  {locales: "lt", form: 0, integer: []string{"0", "10~20", "30", "40", "50", "60", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0", "10.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	78:  {locales: "mt", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	79:  {locales: "mt", form: 4, integer: []string{"0", "2~10", "102~107", "1002"}, decimal: []string{"0.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "10.0", "102.0", "1002.0"}},
	80:  {locales: "mt", form: 5, integer: []string{"11~19", "111~117", "1011"}, decimal: []string{"11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "111.0", "1011.0"}},
	81:  {locales: "mt", form: 0, integer: []string{"20~35", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	82:  {locales: "ru uk", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"}, decimal: []string(nil)},
	83:  {locales: "ru uk", form: 4, integer: []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"}, decimal: []string(nil)},
	84:  {locales: "ru uk", form: 5, integer: []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	85:  {locales: "ru uk", form: 0, integer: []string(nil), decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	86:  {locales: "br", form: 2, integer: []string{"1", "21", "31", "41", "51", "61", "81", "101", "1001"}, decimal: []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "81.0", "101.0", "1001.0"}},
	87:  {locales: "br", form: 3, integer: []string{"2", "22", "32", "42", "52", "62", "82", "102", "1002"}, decimal: []string{"2.0", "22.0", "32.0", "42.0", "52.0", "62.0", "82.0", "102.0", "1002.0"}},
	88:  {locales: "br", form: 4, integer: []string{"3", "4", "9", "23", "24", "29", "33", "34", "39", "43", "44", "49", "103", "1003"}, decimal: []string{"3.0", "4.0", "9.0", "23.0", "24.0", "29.0", "33.0", "34.0", "103.0", "1003.0"}},
	89:  {locales: "br", form: 5, integer: []string{"1000000"}, decimal: []string{"1000000.0", "1000000.00", "1000000.000"}},
	90:  {locales: "br", form: 0, integer: []string{"0", "5~8", "10~20", "100", "1000", "10000", "100000"}, decimal: []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0"}},
	91:  {locales: "ga", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	92:  {locales: "ga", form: 3, integer: []string{"2"}, decimal: []string{"2.0", "2.00", "2.000", "2.0000"}},
	93:  {locales: "ga", form: 4, integer: []string{"3~6"}, decimal: []string{"3.0", "4.0", "5.0", "6.0", "3.00", "4.00", "5.00", "6.00", "3.000", "4.000", "5.000", "6.000", "3.0000", "4.0000", "5.0000", "6.0000"}},
	94:  {locales: "ga", form: 5, integer: []string{"7~10"}, decimal: []string{"7.0", "8.0", "9.0", "10.0", "7.00", "8.00", "9.00", "10.00", "7.000", "8.000", "9.000", "10.000", "7.0000", "8.0000", "9.0000", "10.0000"}},
	95:  {locales: "ga", form: 0, integer: []string{"0", "11~25", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.0~0.9", "1.1~1.6", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	96:  {locales: "gv", form: 2, integer: []string{"1", "11", "21", "31", "41", "51", "61", "71", "101", "1001"}, decimal: []string(nil)},
	97:  {locales: "gv", form: 3, integer: []string{"2", "12", "22", "32", "42", "52", "62", "72", "102", "1002"}, decimal: []string(nil)},
	98:  {locales: "gv", form: 4, integer: []string{"0", "20", "40", "60", "80", "100", "120", "140", "1000", "10000", "100000", "1000000"}, decimal: []string(nil)},
	99:  {locales: "gv", form: 5, integer: []string(nil), decimal: []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	100: {locales: "gv", form: 0, integer: []string{"3~10", "13~19", "23", "103", "1003"}, decimal: []string(nil)},
	101: {locales: "ar ars", form: 1, integer: []string{"0"}, decimal: []string{"0.0", "0.00", "0.000", "0.0000"}},
	102: {locales: "ar ars", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	103: {locales: "ar ars", form: 3, integer: []string{"2"}, decimal: []string{"2.0", "2.00", "2.000", "2.0000"}},
	104: {locales: "ar ars", form: 4, integer: []string{"3~10", "103~110", "1003"}, decimal: []string{"3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "103.0", "1003.0"}},
	105: {locales: "ar ars", form: 5, integer: []string{"11~26", "111", "1011"}, decimal: []string{"11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "111.0", "1011.0"}},
	106: {locales: "ar ars", form: 0, integer: []string{"100~102", "200~202", "300~302", "400~402", "500~502", "600", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
	107: {locales: "cy", form: 1, integer: []string{"0"}, decimal: []string{"0.0", "0.00", "0.000", "0.0000"}},
	108: {locales: "cy", form: 2, integer: []string{"1"}, decimal: []string{"1.0", "1.00", "1.000", "1.0000"}},
	109: {locales: "cy", form: 3, integer: []string{"2"}, decimal: []string{"2.0", "2.00", "2.000", "2.0000"}},
	110: {locales: "cy", form: 4, integer: []string{"3"}, decimal: []string{"3.0", "3.00", "3.000", "3.0000"}},
	111: {locales: "cy", form: 5, integer: []string{"6"}, decimal: []string{"6.0", "6.00", "6.000", "6.0000"}},
	112: {locales: "cy", form: 0, integer: []string{"4", "5", "7~20", "100", "1000", "10000", "100000", "1000000"}, decimal: []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"}},
} // Size: 8160 bytes

// Total table size 12936 bytes (12KiB); checksum: 8456DC5D

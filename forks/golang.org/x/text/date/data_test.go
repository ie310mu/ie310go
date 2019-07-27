// Code generated by running "go generate" in github.com/ie310mu/ie310go/forks/golang.org/x/text. DO NOT EDIT.

package date

var enumMap = map[string]uint16{
	"":                           0,
	"calendars":                  0,
	"fields":                     1,
	"timeZoneNames":              2,
	"buddhist":                   0,
	"chinese":                    1,
	"coptic":                     2,
	"dangi":                      3,
	"ethiopic":                   4,
	"ethiopic-amete-alem":        5,
	"generic":                    6,
	"gregorian":                  7,
	"hebrew":                     8,
	"indian":                     9,
	"islamic":                    10,
	"islamic-civil":              11,
	"islamic-rgsa":               12,
	"islamic-tbla":               13,
	"islamic-umalqura":           14,
	"japanese":                   15,
	"persian":                    16,
	"roc":                        17,
	"months":                     0,
	"days":                       1,
	"quarters":                   2,
	"dayPeriods":                 3,
	"eras":                       4,
	"dateFormats":                5,
	"timeFormats":                6,
	"dateTimeFormats":            7,
	"monthPatterns":              8,
	"cyclicNameSets":             9,
	"format":                     0,
	"stand-alone":                1,
	"numeric":                    2,
	"widthAbbreviated":           0,
	"widthNarrow":                1,
	"widthWide":                  2,
	"widthAll":                   3,
	"widthShort":                 4,
	"leap7":                      0,
	"sun":                        0,
	"mon":                        1,
	"tue":                        2,
	"wed":                        3,
	"thu":                        4,
	"fri":                        5,
	"sat":                        6,
	"am":                         0,
	"pm":                         1,
	"midnight":                   2,
	"morning1":                   3,
	"afternoon1":                 4,
	"evening1":                   5,
	"night1":                     6,
	"noon":                       7,
	"morning2":                   8,
	"afternoon2":                 9,
	"night2":                     10,
	"evening2":                   11,
	"variant":                    1,
	"short":                      0,
	"long":                       1,
	"full":                       2,
	"medium":                     3,
	"dayPartsCycleType":          0,
	"daysCycleType":              1,
	"monthsCycleType":            2,
	"solarTermsCycleType":        3,
	"yearsCycleType":             4,
	"zodiacsCycleType":           5,
	"eraField":                   0,
	"era-shortField":             1,
	"era-narrowField":            2,
	"yearField":                  3,
	"year-shortField":            4,
	"year-narrowField":           5,
	"quarterField":               6,
	"quarter-shortField":         7,
	"quarter-narrowField":        8,
	"monthField":                 9,
	"month-shortField":           10,
	"month-narrowField":          11,
	"weekField":                  12,
	"week-shortField":            13,
	"week-narrowField":           14,
	"weekOfMonthField":           15,
	"weekOfMonth-shortField":     16,
	"weekOfMonth-narrowField":    17,
	"dayField":                   18,
	"day-shortField":             19,
	"day-narrowField":            20,
	"dayOfYearField":             21,
	"dayOfYear-shortField":       22,
	"dayOfYear-narrowField":      23,
	"weekdayField":               24,
	"weekday-shortField":         25,
	"weekday-narrowField":        26,
	"weekdayOfMonthField":        27,
	"weekdayOfMonth-shortField":  28,
	"weekdayOfMonth-narrowField": 29,
	"sunField":                   30,
	"sun-shortField":             31,
	"sun-narrowField":            32,
	"monField":                   33,
	"mon-shortField":             34,
	"mon-narrowField":            35,
	"tueField":                   36,
	"tue-shortField":             37,
	"tue-narrowField":            38,
	"wedField":                   39,
	"wed-shortField":             40,
	"wed-narrowField":            41,
	"thuField":                   42,
	"thu-shortField":             43,
	"thu-narrowField":            44,
	"friField":                   45,
	"fri-shortField":             46,
	"fri-narrowField":            47,
	"satField":                   48,
	"sat-shortField":             49,
	"sat-narrowField":            50,
	"dayperiod-shortField":       51,
	"dayperiodField":             52,
	"dayperiod-narrowField":      53,
	"hourField":                  54,
	"hour-shortField":            55,
	"hour-narrowField":           56,
	"minuteField":                57,
	"minute-shortField":          58,
	"minute-narrowField":         59,
	"secondField":                60,
	"second-shortField":          61,
	"second-narrowField":         62,
	"zoneField":                  63,
	"zone-shortField":            64,
	"zone-narrowField":           65,
	"displayName":                0,
	"relative":                   1,
	"relativeTime":               2,
	"relativePeriod":             3,
	"before1":                    0,
	"current":                    1,
	"after1":                     2,
	"before2":                    3,
	"after2":                     4,
	"after3":                     5,
	"future":                     0,
	"past":                       1,
	"other":                      0,
	"one":                        1,
	"zero":                       2,
	"two":                        3,
	"few":                        4,
	"many":                       5,
	"zoneFormat":                 0,
	"regionFormat":               1,
	"zone":                       2,
	"metaZone":                   3,
	"hourFormat":                 0,
	"gmtFormat":                  1,
	"gmtZeroFormat":              2,
	"genericTime":                0,
	"daylightTime":               1,
	"standardTime":               2,
	"Etc/UTC":                    0,
	"Europe/London":              1,
	"Europe/Dublin":              2,
	"Pacific/Honolulu":           3,
	"Afghanistan":                0,
	"Africa_Central":             1,
	"Africa_Eastern":             2,
	"Africa_Southern":            3,
	"Africa_Western":             4,
	"Alaska":                     5,
	"Amazon":                     6,
	"America_Central":            7,
	"America_Eastern":            8,
	"America_Mountain":           9,
	"America_Pacific":            10,
	"Anadyr":                     11,
	"Apia":                       12,
	"Arabian":                    13,
	"Argentina":                  14,
	"Argentina_Western":          15,
	"Armenia":                    16,
	"Atlantic":                   17,
	"Australia_Central":          18,
	"Australia_CentralWestern":   19,
	"Australia_Eastern":          20,
	"Australia_Western":          21,
	"Azerbaijan":                 22,
	"Azores":                     23,
	"Bangladesh":                 24,
	"Bhutan":                     25,
	"Bolivia":                    26,
	"Brasilia":                   27,
	"Brunei":                     28,
	"Cape_Verde":                 29,
	"Chamorro":                   30,
	"Chatham":                    31,
	"Chile":                      32,
	"China":                      33,
	"Choibalsan":                 34,
	"Christmas":                  35,
	"Cocos":                      36,
	"Colombia":                   37,
	"Cook":                       38,
	"Cuba":                       39,
	"Davis":                      40,
	"DumontDUrville":             41,
	"East_Timor":                 42,
	"Easter":                     43,
	"Ecuador":                    44,
	"Europe_Central":             45,
	"Europe_Eastern":             46,
	"Europe_Further_Eastern":     47,
	"Europe_Western":             48,
	"Falkland":                   49,
	"Fiji":                       50,
	"French_Guiana":              51,
	"French_Southern":            52,
	"Galapagos":                  53,
	"Gambier":                    54,
	"Georgia":                    55,
	"Gilbert_Islands":            56,
	"GMT":                        57,
	"Greenland_Eastern":          58,
	"Greenland_Western":          59,
	"Gulf":                       60,
	"Guyana":                     61,
	"Hawaii_Aleutian":            62,
	"Hong_Kong":                  63,
	"Hovd":                       64,
	"India":                      65,
	"Indian_Ocean":               66,
	"Indochina":                  67,
	"Indonesia_Central":          68,
	"Indonesia_Eastern":          69,
	"Indonesia_Western":          70,
	"Iran":                       71,
	"Irkutsk":                    72,
	"Israel":                     73,
	"Japan":                      74,
	"Kamchatka":                  75,
	"Kazakhstan_Eastern":         76,
	"Kazakhstan_Western":         77,
	"Korea":                      78,
	"Kosrae":                     79,
	"Krasnoyarsk":                80,
	"Kyrgystan":                  81,
	"Line_Islands":               82,
	"Lord_Howe":                  83,
	"Macquarie":                  84,
	"Magadan":                    85,
	"Malaysia":                   86,
	"Maldives":                   87,
	"Marquesas":                  88,
	"Marshall_Islands":           89,
	"Mauritius":                  90,
	"Mawson":                     91,
	"Mexico_Northwest":           92,
	"Mexico_Pacific":             93,
	"Mongolia":                   94,
	"Moscow":                     95,
	"Myanmar":                    96,
	"Nauru":                      97,
	"Nepal":                      98,
	"New_Caledonia":              99,
	"New_Zealand":                100,
	"Newfoundland":               101,
	"Niue":                       102,
	"Norfolk":                    103,
	"Noronha":                    104,
	"Novosibirsk":                105,
	"Omsk":                       106,
	"Pakistan":                   107,
	"Palau":                      108,
	"Papua_New_Guinea":           109,
	"Paraguay":                   110,
	"Peru":                       111,
	"Philippines":                112,
	"Phoenix_Islands":            113,
	"Pierre_Miquelon":            114,
	"Pitcairn":                   115,
	"Ponape":                     116,
	"Pyongyang":                  117,
	"Reunion":                    118,
	"Rothera":                    119,
	"Sakhalin":                   120,
	"Samara":                     121,
	"Samoa":                      122,
	"Seychelles":                 123,
	"Singapore":                  124,
	"Solomon":                    125,
	"South_Georgia":              126,
	"Suriname":                   127,
	"Syowa":                      128,
	"Tahiti":                     129,
	"Taipei":                     130,
	"Tajikistan":                 131,
	"Tokelau":                    132,
	"Tonga":                      133,
	"Truk":                       134,
	"Turkmenistan":               135,
	"Tuvalu":                     136,
	"Uruguay":                    137,
	"Uzbekistan":                 138,
	"Vanuatu":                    139,
	"Venezuela":                  140,
	"Vladivostok":                141,
	"Volgograd":                  142,
	"Vostok":                     143,
	"Wake":                       144,
	"Wallis":                     145,
	"Yakutsk":                    146,
	"Yekaterinburg":              147,
	"Guam":                       148,
	"North_Mariana":              149,
	"Acre":                       150,
	"Almaty":                     151,
	"Aqtau":                      152,
	"Aqtobe":                     153,
	"Casey":                      154,
	"Lanka":                      155,
	"Macau":                      156,
	"Qyzylorda":                  157,
}

// Total table size 0 bytes (0KiB); checksum: 811C9DC5

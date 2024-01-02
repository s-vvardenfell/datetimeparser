package datetimeparser

import (
	"time"
)

var (
	endYearShort   = (time.Now().Year() + 1) % 100 // 24 for 2023
	startYearLong  = 1999
	endYearLong    = time.Now().Year() + 1 // 2024 for 2023
	chineeseFormat = "2006年01月02日15時04分"

	tzOffsetLayouts = [2]string{
		"-0700",
		"-07:00",
	}

	numbersInWords = map[string]int{
		"one":    1,
		"two":    2,
		"three":  3,
		"four":   4,
		"five":   5,
		"six":    6,
		"seven":  7,
		"eight":  8,
		"nine":   9,
		"ten":    10,
		"eleven": 11,
		"twelve": 12,
	}

	continuousFormats = []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04",
		"2006-01-02T15:04:05",
		"2006-1-2T15:4:5-07:00",
		"2006-01-02T15:04-0700",
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02MST15:04:05-07:00",
		"2006-01-02T15:04:05.9999999-07:00",
		"2006-01-02Z15:04:05.9999999-07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
	}

	namedTimeZones = map[string]int{
		"ACDT":  +10*3600 + 1800, // Australian Central Daylight Savings Time
		"ACST":  +9*3600 + 1800,  // Australian Central Standard Time
		"ACT":   -5 * 3600,       // Acre Time
		"ACWST": +8*3600 + 2700,  // Australian Central Western Standard Time (unofficial)
		"ADT":   -3 * 3600,       // Atlantic Daylight Time
		"AEDT":  +11 * 3600,      // Australian Eastern Daylight Savings Time
		"AEST":  +10 * 3600,      // Australian Eastern Standard Time
		"AFT":   +4*3600 + 1800,  // Afghanistan Time
		"AKDT":  -8 * 3600,       // Alaska Daylight Time
		"AKST":  -9 * 3600,       // Alaska Standard Time
		"AMST":  -3 * 3600,       // Amazon Summer Time (Brazil)[1]
		"AMT":   +4 * 3600,       // Armenia Time
		"ART":   -3 * 3600,       // Argentina Time
		"AST":   -4 * 3600,       // Atlantic Standard Time
		"AWST":  +8 * 3600,       // Australian Western Standard Time
		"AZOST": 0 * 3600,        // Azores Summer Time
		"AZOT":  -1 * 3600,       // Azores Standard Time
		"AZT":   +4 * 3600,       // Azerbaijan Time
		"BDT":   +8 * 3600,       // Brunei Time
		"BIOT":  +6 * 3600,       // British Indian Ocean Time
		"BIT":   -12 * 3600,      // Baker Island Time
		"BOT":   -4 * 3600,       // Bolivia Time
		"BRST":  -2 * 3600,       // Bras�lia Summer Time
		"BRT":   -3 * 3600,       // Brasilia Time
		"BST":   +6 * 3600,       // Bangladesh Standard Time
		"BTT":   +6 * 3600,       // Bhutan Time
		"CAT":   +2 * 3600,       // Central Africa Time
		"CCT":   +6*3600 + 1800,  // Cocos Islands Time
		"CDT":   -5 * 3600,       // Central Daylight Time (North America)
		"CEST":  +2 * 3600,       // Central European Summer Time (Cf. HAEC)
		"CET":   +1 * 3600,       // Central European Time
		"CHADT": +13*3600 + 2700, // Chatham Daylight Time
		"CHAST": +12*3600 + 2700, // Chatham Standard Time
		"CHOT":  +8 * 3600,       // Choibalsan Standard Time
		"CHOST": +9 * 3600,       // Choibalsan Summer Time
		"CHST":  +10 * 3600,      // Chamorro Standard Time
		"CHUT":  +10 * 3600,      // Chuuk Time
		"CIST":  -8 * 3600,       // Clipperton Island Standard Time
		"CIT":   +8 * 3600,       // Central Indonesia Time
		"CKT":   -10 * 3600,      // Cook Island Time
		"CLST":  -3 * 3600,       // Chile Summer Time
		"CLT":   -4 * 3600,       // Chile Standard Time
		"COST":  -4 * 3600,       // Colombia Summer Time
		"COT":   -5 * 3600,       // Colombia Time
		"CST":   +8 * 3600,       // China Standard Time
		"CT":    +8 * 3600,       // China Time
		"CVT":   -1 * 3600,       // Cape Verde Time
		"CWST":  +8*3600 + 2700,  // Central Western Standard Time (Australia) unofficial
		"CXT":   +7 * 3600,       // Christmas Island Time
		"DAVT":  +7 * 3600,       // Davis Time
		"DDUT":  +10 * 3600,      // Dumont d'Urville Time
		"DFT":   +1 * 3600,       // AIX-specific equivalent of Central European Time[NB 1]
		"EASST": -5 * 3600,       // Easter Island Summer Time
		"EAST":  -6 * 3600,       // Easter Island Standard Time
		"EAT":   +3 * 3600,       // East Africa Time
		"ECT":   -4 * 3600,       // Eastern Caribbean Time (does not recognise DST)
		"EDT":   -4 * 3600,       // Eastern Daylight Time (North America)
		"EEST":  +3 * 3600,       // Eastern European Summer Time
		"EET":   +2 * 3600,       // Eastern European Time
		"EGST":  0 * 3600,        // Eastern Greenland Summer Time
		"EGT":   -1 * 3600,       // Eastern Greenland Time
		"EIT":   +9 * 3600,       // Eastern Indonesian Time
		"EST":   -5 * 3600,       // Eastern Standard Time (North America)
		"FET":   +3 * 3600,       // Further-eastern European Time
		"FJT":   +12 * 3600,      // Fiji Time
		"FKST":  -3 * 3600,       // Falkland Islands Summer Time
		"FKT":   -4 * 3600,       // Falkland Islands Time
		"FNT":   -2 * 3600,       // Fernando de Noronha Time
		"GALT":  -6 * 3600,       // Gal�pagos Time
		"GAMT":  -9 * 3600,       // Gambier Islands Time
		"GET":   +4 * 3600,       // Georgia Standard Time
		"GFT":   -3 * 3600,       // French Guiana Time
		"GILT":  +12 * 3600,      // Gilbert Island Time
		"GIT":   -9 * 3600,       // Gambier Island Time
		"GMT":   0 * 3600,        // Greenwich Mean Time
		"GST":   -2 * 3600,       // South Georgia and the South Sandwich Islands Time
		"GYT":   -4 * 3600,       // Guyana Time
		"HDT":   -9 * 3600,       // Hawaii�Aleutian Daylight Time
		"HAEC":  +2 * 3600,       // Heure Avanc�e d'Europe Centrale French-language name for CEST
		"HST":   -10 * 3600,      // Hawaii�Aleutian Standard Time
		"HKT":   +8 * 3600,       // Hong Kong Time
		"HMT":   +5 * 3600,       // Heard and McDonald Islands Time
		"HOVST": +8 * 3600,       // Khovd Summer Time
		"HOVT":  +7 * 3600,       // Khovd Standard Time
		"ICT":   +7 * 3600,       // Indochina Time
		"IDLW":  -12 * 3600,      // International Day Line West time zone
		"IDT":   +3 * 3600,       // Israel Daylight Time
		"IOT":   +3 * 3600,       // Indian Ocean Time
		"IRDT":  +4*3600 + 1800,  // Iran Daylight Time
		"IRKT":  +8 * 3600,       // Irkutsk Time
		"IRST":  +3*3600 + 1800,  // Iran Standard Time
		"IST":   +5*3600 + 1800,  // Indian Standard Time
		"JST":   +9 * 3600,       // Japan Standard Time
		"KGT":   +6 * 3600,       // Kyrgyzstan Time
		"KOST":  +11 * 3600,      // Kosrae Time
		"KRAT":  +7 * 3600,       // Krasnoyarsk Time
		"KST":   +9 * 3600,       // Korea Standard Time
		"LHST":  +10*3600 + 1800, // Lord Howe Standard Time
		"LINT":  +14 * 3600,      // Line Islands Time
		"MAGT":  +12 * 3600,      // Magadan Time
		"MART":  -9*3600 - 1800,  // Marquesas Islands Time
		"MAWT":  +5 * 3600,       // Mawson Station Time
		"MDT":   -6 * 3600,       // Mountain Daylight Time (North America)
		"MET":   +1 * 3600,       // Middle European Time Same zone as CET
		"MEST":  +2 * 3600,       // Middle European Summer Time Same zone as CEST
		"MHT":   +12 * 3600,      // Marshall Islands Time
		"MIST":  +11 * 3600,      // Macquarie Island Station Time
		"MIT":   -9*3600 - 1800,  // Marquesas Islands Time
		"MMT":   +6*3600 + 1800,  // Myanmar Standard Time
		"MSK":   +3 * 3600,       // Moscow Time
		"MST":   -7 * 3600,       // Mountain Standard Time (North America)
		"MUT":   +4 * 3600,       // Mauritius Time
		"MVT":   +5 * 3600,       // Maldives Time
		"MYT":   +8 * 3600,       // Malaysia Time
		"NCT":   +11 * 3600,      // New Caledonia Time
		"NDT":   -2*3600 - 1800,  // Newfoundland Daylight Time
		"NFT":   +11 * 3600,      // Norfolk Island Time
		"NPT":   +5*3600 + 2700,  // Nepal Time
		"NST":   -3*3600 - 1800,  // Newfoundland Standard Time
		"NT":    -3*3600 - 1800,  // Newfoundland Time
		"NUT":   -11 * 3600,      // Niue Time
		"NZDT":  +13 * 3600,      // New Zealand Daylight Time
		"NZST":  +12 * 3600,      // New Zealand Standard Time
		"OMST":  +6 * 3600,       // Omsk Time
		"ORAT":  +5 * 3600,       // Oral Time
		"PDT":   -7 * 3600,       // Pacific Daylight Time (North America)
		"PET":   -5 * 3600,       // Peru Time
		"PETT":  +12 * 3600,      // Kamchatka Time
		"PGT":   +10 * 3600,      // Papua New Guinea Time
		"PHOT":  +13 * 3600,      // Phoenix Island Time
		"PHT":   +8 * 3600,       // Philippine Time
		"PKT":   +5 * 3600,       // Pakistan Standard Time
		"PMDT":  -2 * 3600,       // Saint Pierre and Miquelon Daylight Time
		"PMST":  -3 * 3600,       // Saint Pierre and Miquelon Standard Time
		"PONT":  +11 * 3600,      // Pohnpei Standard Time
		"PST":   -8 * 3600,       // Pacific Standard Time (North America)
		"PYST":  -3 * 3600,       // Paraguay Summer Time[6]
		"PYT":   -4 * 3600,       // Paraguay Time[7]
		"RET":   +4 * 3600,       // R�union Time
		"ROTT":  -3 * 3600,       // Rothera Research Station Time
		"SAKT":  +11 * 3600,      // Sakhalin Island Time
		"SAMT":  +4 * 3600,       // Samara Time
		"SAST":  +2 * 3600,       // South African Standard Time
		"SBT":   +11 * 3600,      // Solomon Islands Time
		"SCT":   +4 * 3600,       // Seychelles Time
		"SDT":   -10 * 3600,      // Samoa Daylight Time
		"SGT":   +8 * 3600,       // Singapore Time
		"SLST":  +5*3600 + 1800,  // Sri Lanka Standard Time
		"SRET":  +11 * 3600,      // Srednekolymsk Time
		"SRT":   -3 * 3600,       // Suriname Time
		"SST":   +8 * 3600,       // Singapore Standard Time
		"SYOT":  +3 * 3600,       // Showa Station Time
		"TAHT":  -10 * 3600,      // Tahiti Time
		"THA":   +7 * 3600,       // Thailand Standard Time
		"TFT":   +5 * 3600,       // Indian/Kerguelen
		"TJT":   +5 * 3600,       // Tajikistan Time
		"TKT":   +13 * 3600,      // Tokelau Time
		"TLT":   +9 * 3600,       // Timor Leste Time
		"TMT":   +5 * 3600,       // Turkmenistan Time
		"TRT":   +3 * 3600,       // Turkey Time
		"TOT":   +13 * 3600,      // Tonga Time
		"TVT":   +12 * 3600,      // Tuvalu Time
		"ULAST": +9 * 3600,       // Ulaanbaatar Summer Time
		"ULAT":  +8 * 3600,       // Ulaanbaatar Standard Time
		"USZ1":  +2 * 3600,       // Kaliningrad Time
		"UTC":   0 * 3600,        // Coordinated Universal Time
		"UT":    0 * 3600,        // Coordinated Universal Time
		"UYST":  -2 * 3600,       // Uruguay Summer Time
		"UYT":   -3 * 3600,       // Uruguay Standard Time
		"UZT":   +5 * 3600,       // Uzbekistan Time
		"VET":   -4 * 3600,       // Venezuelan Standard Time
		"VLAT":  +10 * 3600,      // Vladivostok Time
		"VOLT":  +4 * 3600,       // Volgograd Time
		"VOST":  +6 * 3600,       // Vostok Station Time
		"VUT":   +11 * 3600,      // Vanuatu Time
		"WAKT":  +12 * 3600,      // Wake Island Time
		"WAST":  +2 * 3600,       // West Africa Summer Time
		"WAT":   +1 * 3600,       // West Africa Time
		"WEST":  +1 * 3600,       // Western European Summer Time
		"WET":   0 * 3600,        // Western European Time
		"WIT":   +7 * 3600,       // Western Indonesian Time
		"WST":   +8 * 3600,       // Western Standard Time
		"YAKT":  +9 * 3600,       // Yakutsk Time
		"YEKT":  +5 * 3600,       // Yekaterinburg Time
		"Z":     0 * 3600,
	}
)

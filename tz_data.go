package tz

// GENERATED FILE DO NOT MODIFY DIRECTLY. GENERATED AT: 2021-03-18 17:33:12.486096739 +0300 MSK

var (
	countries = []Country{
		{
			Code: "AF",
			Name: "Afghanistan",
			Zones: []*Zone{
				{
					Id:          3,
					CountryCode: "AF",
					Name:        "Asia/Kabul",
					TimeZones: []*TimeZone{
						{
							ZoneId: 3,
							Abbr:   "LMT",
							Offset: 16608,
						},
						{
							ZoneId: 3,
							Abbr:   "AFT",
							Offset: 14400,
						},
						{
							ZoneId: 3,
							Abbr:   "AFT",
							Offset: 16200,
						},
					},
				},
			},
		},
		{
			Code: "AL",
			Name: "Albania",
			Zones: []*Zone{
				{
					Id:          6,
					CountryCode: "AL",
					Name:        "Europe/Tirane",
					TimeZones: []*TimeZone{
						{
							ZoneId: 6,
							Abbr:   "LMT",
							Offset: 4760,
						},
						{
							ZoneId: 6,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 6,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "DZ",
			Name: "Algeria",
			Zones: []*Zone{
				{
					Id:          140,
					CountryCode: "DZ",
					Name:        "Africa/Algiers",
					TimeZones: []*TimeZone{
						{
							ZoneId: 140,
							Abbr:   "LMT",
							Offset: 732,
						},
						{
							ZoneId: 140,
							Abbr:   "PMT",
							Offset: 561,
						},
						{
							ZoneId: 140,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 140,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 140,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 140,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "AS",
			Name: "American Samoa",
			Zones: []*Zone{
				{
					Id:          31,
					CountryCode: "AS",
					Name:        "Pacific/Pago_Pago",
					TimeZones: []*TimeZone{
						{
							ZoneId: 31,
							Abbr:   "LMT",
							Offset: 45432,
						},
						{
							ZoneId: 31,
							Abbr:   "LMT",
							Offset: -40968,
						},
						{
							ZoneId: 31,
							Abbr:   "SST",
							Offset: -39600,
						},
					},
				},
			},
		},
		{
			Code: "AD",
			Name: "Andorra",
			Zones: []*Zone{
				{
					Id:          1,
					CountryCode: "AD",
					Name:        "Europe/Andorra",
					TimeZones: []*TimeZone{
						{
							ZoneId: 1,
							Abbr:   "LMT",
							Offset: 364,
						},
						{
							ZoneId: 1,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 1,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 1,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "AO",
			Name: "Angola",
			Zones: []*Zone{
				{
					Id:          8,
					CountryCode: "AO",
					Name:        "Africa/Luanda",
					TimeZones: []*TimeZone{
						{
							ZoneId: 8,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 8,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 8,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 8,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "AI",
			Name: "Anguilla",
			Zones: []*Zone{
				{
					Id:          5,
					CountryCode: "AI",
					Name:        "America/Anguilla",
					TimeZones: []*TimeZone{
						{
							ZoneId: 5,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 5,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "AQ",
			Name: "Antarctica",
			Zones: []*Zone{
				{
					Id:          10,
					CountryCode: "AQ",
					Name:        "Antarctica/Casey",
					TimeZones: []*TimeZone{
						{
							ZoneId: 10,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 10,
							Abbr:   "CAST",
							Offset: 28800,
						},
						{
							ZoneId: 10,
							Abbr:   "CAST",
							Offset: 39600,
						},
					},
				},
				{
					Id:          11,
					CountryCode: "AQ",
					Name:        "Antarctica/Davis",
					TimeZones: []*TimeZone{
						{
							ZoneId: 11,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 11,
							Abbr:   "DAVT",
							Offset: 25200,
						},
						{
							ZoneId: 11,
							Abbr:   "DAVT",
							Offset: 18000,
						},
					},
				},
				{
					Id:          12,
					CountryCode: "AQ",
					Name:        "Antarctica/DumontDUrville",
					TimeZones: []*TimeZone{
						{
							ZoneId: 12,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 12,
							Abbr:   "DDUT",
							Offset: 36000,
						},
					},
				},
				{
					Id:          13,
					CountryCode: "AQ",
					Name:        "Antarctica/Mawson",
					TimeZones: []*TimeZone{
						{
							ZoneId: 13,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 13,
							Abbr:   "MAWT",
							Offset: 21600,
						},
						{
							ZoneId: 13,
							Abbr:   "MAWT",
							Offset: 18000,
						},
					},
				},
				{
					Id:          9,
					CountryCode: "AQ",
					Name:        "Antarctica/McMurdo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 9,
							Abbr:   "LMT",
							Offset: 41944,
						},
						{
							ZoneId: 9,
							Abbr:   "NZMT",
							Offset: 41400,
						},
						{
							ZoneId: 9,
							Abbr:   "NZST",
							Offset: 45000,
						},
						{
							ZoneId: 9,
							Abbr:   "NZST",
							Offset: 43200,
						},
						{
							ZoneId: 9,
							Abbr:   "NZDT",
							Offset: 46800,
						},
					},
				},
				{
					Id:          14,
					CountryCode: "AQ",
					Name:        "Antarctica/Palmer",
					TimeZones: []*TimeZone{
						{
							ZoneId: 14,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 14,
							Abbr:   "CLST",
							Offset: -10800,
						},
						{
							ZoneId: 14,
							Abbr:   "CLT",
							Offset: -14400,
						},
						{
							ZoneId: 14,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          15,
					CountryCode: "AQ",
					Name:        "Antarctica/Rothera",
					TimeZones: []*TimeZone{
						{
							ZoneId: 15,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 15,
							Abbr:   "ART",
							Offset: -10800,
						},
					},
				},
				{
					Id:          16,
					CountryCode: "AQ",
					Name:        "Antarctica/Syowa",
					TimeZones: []*TimeZone{
						{
							ZoneId: 16,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 16,
							Abbr:   "SYOT",
							Offset: 10800,
						},
					},
				},
				{
					Id:          17,
					CountryCode: "AQ",
					Name:        "Antarctica/Troll",
					TimeZones: []*TimeZone{
						{
							ZoneId: 17,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 17,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
				{
					Id:          18,
					CountryCode: "AQ",
					Name:        "Antarctica/Vostok",
					TimeZones: []*TimeZone{
						{
							ZoneId: 18,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 18,
							Abbr:   "VOST",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "AG",
			Name: "Antigua and Barbuda",
			Zones: []*Zone{
				{
					Id:          4,
					CountryCode: "AG",
					Name:        "America/Antigua",
					TimeZones: []*TimeZone{
						{
							ZoneId: 4,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 4,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "AR",
			Name: "Argentina",
			Zones: []*Zone{
				{
					Id:          19,
					CountryCode: "AR",
					Name:        "America/Argentina/Buenos_Aires",
					TimeZones: []*TimeZone{
						{
							ZoneId: 19,
							Abbr:   "LMT",
							Offset: -14028,
						},
						{
							ZoneId: 19,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 19,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 19,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 19,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 19,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          24,
					CountryCode: "AR",
					Name:        "America/Argentina/Catamarca",
					TimeZones: []*TimeZone{
						{
							ZoneId: 24,
							Abbr:   "LMT",
							Offset: -15788,
						},
						{
							ZoneId: 24,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 24,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 24,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 24,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 24,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          20,
					CountryCode: "AR",
					Name:        "America/Argentina/Cordoba",
					TimeZones: []*TimeZone{
						{
							ZoneId: 20,
							Abbr:   "LMT",
							Offset: -15408,
						},
						{
							ZoneId: 20,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 20,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 20,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 20,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 20,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          22,
					CountryCode: "AR",
					Name:        "America/Argentina/Jujuy",
					TimeZones: []*TimeZone{
						{
							ZoneId: 22,
							Abbr:   "LMT",
							Offset: -15672,
						},
						{
							ZoneId: 22,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 22,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 22,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 22,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 22,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          25,
					CountryCode: "AR",
					Name:        "America/Argentina/La_Rioja",
					TimeZones: []*TimeZone{
						{
							ZoneId: 25,
							Abbr:   "LMT",
							Offset: -16044,
						},
						{
							ZoneId: 25,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 25,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 25,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 25,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 25,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          27,
					CountryCode: "AR",
					Name:        "America/Argentina/Mendoza",
					TimeZones: []*TimeZone{
						{
							ZoneId: 27,
							Abbr:   "LMT",
							Offset: -16516,
						},
						{
							ZoneId: 27,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 27,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 27,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 27,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 27,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          29,
					CountryCode: "AR",
					Name:        "America/Argentina/Rio_Gallegos",
					TimeZones: []*TimeZone{
						{
							ZoneId: 29,
							Abbr:   "LMT",
							Offset: -16612,
						},
						{
							ZoneId: 29,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 29,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 29,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 29,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 29,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          21,
					CountryCode: "AR",
					Name:        "America/Argentina/Salta",
					TimeZones: []*TimeZone{
						{
							ZoneId: 21,
							Abbr:   "LMT",
							Offset: -15700,
						},
						{
							ZoneId: 21,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 21,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 21,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 21,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 21,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          26,
					CountryCode: "AR",
					Name:        "America/Argentina/San_Juan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 26,
							Abbr:   "LMT",
							Offset: -16444,
						},
						{
							ZoneId: 26,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 26,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 26,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 26,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 26,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          28,
					CountryCode: "AR",
					Name:        "America/Argentina/San_Luis",
					TimeZones: []*TimeZone{
						{
							ZoneId: 28,
							Abbr:   "LMT",
							Offset: -15924,
						},
						{
							ZoneId: 28,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 28,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 28,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 28,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 28,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          23,
					CountryCode: "AR",
					Name:        "America/Argentina/Tucuman",
					TimeZones: []*TimeZone{
						{
							ZoneId: 23,
							Abbr:   "LMT",
							Offset: -15652,
						},
						{
							ZoneId: 23,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 23,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 23,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 23,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 23,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          30,
					CountryCode: "AR",
					Name:        "America/Argentina/Ushuaia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 30,
							Abbr:   "LMT",
							Offset: -16392,
						},
						{
							ZoneId: 30,
							Abbr:   "CMT",
							Offset: -15408,
						},
						{
							ZoneId: 30,
							Abbr:   "ART",
							Offset: -14400,
						},
						{
							ZoneId: 30,
							Abbr:   "ARST",
							Offset: -10800,
						},
						{
							ZoneId: 30,
							Abbr:   "ART",
							Offset: -10800,
						},
						{
							ZoneId: 30,
							Abbr:   "ARST",
							Offset: -7200,
						},
					},
				},
			},
		},
		{
			Code: "AM",
			Name: "Armenia",
			Zones: []*Zone{
				{
					Id:          7,
					CountryCode: "AM",
					Name:        "Asia/Yerevan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 7,
							Abbr:   "LMT",
							Offset: 10680,
						},
						{
							ZoneId: 7,
							Abbr:   "AMT",
							Offset: 10800,
						},
						{
							ZoneId: 7,
							Abbr:   "AMT",
							Offset: 14400,
						},
						{
							ZoneId: 7,
							Abbr:   "AMST",
							Offset: 18000,
						},
						{
							ZoneId: 7,
							Abbr:   "AMST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "AW",
			Name: "Aruba",
			Zones: []*Zone{
				{
					Id:          45,
					CountryCode: "AW",
					Name:        "America/Aruba",
					TimeZones: []*TimeZone{
						{
							ZoneId: 45,
							Abbr:   "LMT",
							Offset: -16547,
						},
						{
							ZoneId: 45,
							Abbr:   "AST",
							Offset: -16200,
						},
						{
							ZoneId: 45,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "AU",
			Name: "Australia",
			Zones: []*Zone{
				{
					Id:          34,
					CountryCode: "AU",
					Name:        "Antarctica/Macquarie",
					TimeZones: []*TimeZone{
						{
							ZoneId: 34,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 34,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 34,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          41,
					CountryCode: "AU",
					Name:        "Australia/Adelaide",
					TimeZones: []*TimeZone{
						{
							ZoneId: 41,
							Abbr:   "LMT",
							Offset: 33260,
						},
						{
							ZoneId: 41,
							Abbr:   "ACST",
							Offset: 32400,
						},
						{
							ZoneId: 41,
							Abbr:   "ACST",
							Offset: 34200,
						},
						{
							ZoneId: 41,
							Abbr:   "ACDT",
							Offset: 37800,
						},
					},
				},
				{
					Id:          39,
					CountryCode: "AU",
					Name:        "Australia/Brisbane",
					TimeZones: []*TimeZone{
						{
							ZoneId: 39,
							Abbr:   "LMT",
							Offset: 36728,
						},
						{
							ZoneId: 39,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 39,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          38,
					CountryCode: "AU",
					Name:        "Australia/Broken_Hill",
					TimeZones: []*TimeZone{
						{
							ZoneId: 38,
							Abbr:   "LMT",
							Offset: 33948,
						},
						{
							ZoneId: 38,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 38,
							Abbr:   "ACST",
							Offset: 32400,
						},
						{
							ZoneId: 38,
							Abbr:   "ACST",
							Offset: 34200,
						},
						{
							ZoneId: 38,
							Abbr:   "ACDT",
							Offset: 37800,
						},
					},
				},
				{
					Id:          42,
					CountryCode: "AU",
					Name:        "Australia/Darwin",
					TimeZones: []*TimeZone{
						{
							ZoneId: 42,
							Abbr:   "LMT",
							Offset: 31400,
						},
						{
							ZoneId: 42,
							Abbr:   "ACST",
							Offset: 32400,
						},
						{
							ZoneId: 42,
							Abbr:   "ACST",
							Offset: 34200,
						},
						{
							ZoneId: 42,
							Abbr:   "ACDT",
							Offset: 37800,
						},
					},
				},
				{
					Id:          44,
					CountryCode: "AU",
					Name:        "Australia/Eucla",
					TimeZones: []*TimeZone{
						{
							ZoneId: 44,
							Abbr:   "LMT",
							Offset: 30928,
						},
						{
							ZoneId: 44,
							Abbr:   "ACWST",
							Offset: 31500,
						},
						{
							ZoneId: 44,
							Abbr:   "ACWDT",
							Offset: 35100,
						},
					},
				},
				{
					Id:          35,
					CountryCode: "AU",
					Name:        "Australia/Hobart",
					TimeZones: []*TimeZone{
						{
							ZoneId: 35,
							Abbr:   "LMT",
							Offset: 35356,
						},
						{
							ZoneId: 35,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 35,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          40,
					CountryCode: "AU",
					Name:        "Australia/Lindeman",
					TimeZones: []*TimeZone{
						{
							ZoneId: 40,
							Abbr:   "LMT",
							Offset: 35756,
						},
						{
							ZoneId: 40,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 40,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          33,
					CountryCode: "AU",
					Name:        "Australia/Lord_Howe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 33,
							Abbr:   "LMT",
							Offset: 38180,
						},
						{
							ZoneId: 33,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 33,
							Abbr:   "LHDT",
							Offset: 37800,
						},
						{
							ZoneId: 33,
							Abbr:   "LHST",
							Offset: 41400,
						},
						{
							ZoneId: 33,
							Abbr:   "LHST",
							Offset: 39600,
						},
					},
				},
				{
					Id:          36,
					CountryCode: "AU",
					Name:        "Australia/Melbourne",
					TimeZones: []*TimeZone{
						{
							ZoneId: 36,
							Abbr:   "LMT",
							Offset: 34792,
						},
						{
							ZoneId: 36,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 36,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          43,
					CountryCode: "AU",
					Name:        "Australia/Perth",
					TimeZones: []*TimeZone{
						{
							ZoneId: 43,
							Abbr:   "LMT",
							Offset: 27804,
						},
						{
							ZoneId: 43,
							Abbr:   "AWST",
							Offset: 28800,
						},
						{
							ZoneId: 43,
							Abbr:   "AWDT",
							Offset: 32400,
						},
					},
				},
				{
					Id:          37,
					CountryCode: "AU",
					Name:        "Australia/Sydney",
					TimeZones: []*TimeZone{
						{
							ZoneId: 37,
							Abbr:   "LMT",
							Offset: 36292,
						},
						{
							ZoneId: 37,
							Abbr:   "AEST",
							Offset: 36000,
						},
						{
							ZoneId: 37,
							Abbr:   "AEDT",
							Offset: 39600,
						},
					},
				},
			},
		},
		{
			Code: "AT",
			Name: "Austria",
			Zones: []*Zone{
				{
					Id:          32,
					CountryCode: "AT",
					Name:        "Europe/Vienna",
					TimeZones: []*TimeZone{
						{
							ZoneId: 32,
							Abbr:   "LMT",
							Offset: 3921,
						},
						{
							ZoneId: 32,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 32,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "AZ",
			Name: "Azerbaijan",
			Zones: []*Zone{
				{
					Id:          47,
					CountryCode: "AZ",
					Name:        "Asia/Baku",
					TimeZones: []*TimeZone{
						{
							ZoneId: 47,
							Abbr:   "LMT",
							Offset: 11964,
						},
						{
							ZoneId: 47,
							Abbr:   "AZT",
							Offset: 10800,
						},
						{
							ZoneId: 47,
							Abbr:   "AZT",
							Offset: 14400,
						},
						{
							ZoneId: 47,
							Abbr:   "AZST",
							Offset: 18000,
						},
						{
							ZoneId: 47,
							Abbr:   "AZST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "BS",
			Name: "Bahamas",
			Zones: []*Zone{
				{
					Id:          78,
					CountryCode: "BS",
					Name:        "America/Nassau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 78,
							Abbr:   "LMT",
							Offset: -18570,
						},
						{
							ZoneId: 78,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 78,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 78,
							Abbr:   "EPT",
							Offset: -14400,
						},
						{
							ZoneId: 78,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "BH",
			Name: "Bahrain",
			Zones: []*Zone{
				{
					Id:          54,
					CountryCode: "BH",
					Name:        "Asia/Bahrain",
					TimeZones: []*TimeZone{
						{
							ZoneId: 54,
							Abbr:   "LMT",
							Offset: 12368,
						},
						{
							ZoneId: 54,
							Abbr:   "GST",
							Offset: 14400,
						},
						{
							ZoneId: 54,
							Abbr:   "AST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "BD",
			Name: "Bangladesh",
			Zones: []*Zone{
				{
					Id:          50,
					CountryCode: "BD",
					Name:        "Asia/Dhaka",
					TimeZones: []*TimeZone{
						{
							ZoneId: 50,
							Abbr:   "LMT",
							Offset: 21700,
						},
						{
							ZoneId: 50,
							Abbr:   "HMT",
							Offset: 21200,
						},
						{
							ZoneId: 50,
							Abbr:   "BURT",
							Offset: 23400,
						},
						{
							ZoneId: 50,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 50,
							Abbr:   "BST",
							Offset: 21600,
						},
						{
							ZoneId: 50,
							Abbr:   "BDT",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "BB",
			Name: "Barbados",
			Zones: []*Zone{
				{
					Id:          49,
					CountryCode: "BB",
					Name:        "America/Barbados",
					TimeZones: []*TimeZone{
						{
							ZoneId: 49,
							Abbr:   "LMT",
							Offset: -14309,
						},
						{
							ZoneId: 49,
							Abbr:   "BMT",
							Offset: -14309,
						},
						{
							ZoneId: 49,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 49,
							Abbr:   "ADT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "BY",
			Name: "Belarus",
			Zones: []*Zone{
				{
					Id:          81,
					CountryCode: "BY",
					Name:        "Europe/Minsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 81,
							Abbr:   "LMT",
							Offset: 6616,
						},
						{
							ZoneId: 81,
							Abbr:   "MMT",
							Offset: 6600,
						},
						{
							ZoneId: 81,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 81,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 81,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 81,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 81,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 81,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "BE",
			Name: "Belgium",
			Zones: []*Zone{
				{
					Id:          51,
					CountryCode: "BE",
					Name:        "Europe/Brussels",
					TimeZones: []*TimeZone{
						{
							ZoneId: 51,
							Abbr:   "LMT",
							Offset: 1050,
						},
						{
							ZoneId: 51,
							Abbr:   "BMT",
							Offset: 1050,
						},
						{
							ZoneId: 51,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 51,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 51,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 51,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "BZ",
			Name: "Belize",
			Zones: []*Zone{
				{
					Id:          82,
					CountryCode: "BZ",
					Name:        "America/Belize",
					TimeZones: []*TimeZone{
						{
							ZoneId: 82,
							Abbr:   "LMT",
							Offset: -21168,
						},
						{
							ZoneId: 82,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 82,
							Abbr:   "CDT",
							Offset: -19800,
						},
						{
							ZoneId: 82,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 82,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 82,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "BJ",
			Name: "Benin",
			Zones: []*Zone{
				{
					Id:          56,
					CountryCode: "BJ",
					Name:        "Africa/Porto-Novo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 56,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 56,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 56,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 56,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "BM",
			Name: "Bermuda",
			Zones: []*Zone{
				{
					Id:          58,
					CountryCode: "BM",
					Name:        "Atlantic/Bermuda",
					TimeZones: []*TimeZone{
						{
							ZoneId: 58,
							Abbr:   "LMT",
							Offset: -15558,
						},
						{
							ZoneId: 58,
							Abbr:   "BMT",
							Offset: -15558,
						},
						{
							ZoneId: 58,
							Abbr:   "BST",
							Offset: -11958,
						},
						{
							ZoneId: 58,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 58,
							Abbr:   "ADT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "BT",
			Name: "Bhutan",
			Zones: []*Zone{
				{
					Id:          79,
					CountryCode: "BT",
					Name:        "Asia/Thimphu",
					TimeZones: []*TimeZone{
						{
							ZoneId: 79,
							Abbr:   "LMT",
							Offset: 21516,
						},
						{
							ZoneId: 79,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 79,
							Abbr:   "BTT",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "BO",
			Name: "Bolivia, Plurinational State of",
			Zones: []*Zone{
				{
					Id:          60,
					CountryCode: "BO",
					Name:        "America/La_Paz",
					TimeZones: []*TimeZone{
						{
							ZoneId: 60,
							Abbr:   "LMT",
							Offset: -16356,
						},
						{
							ZoneId: 60,
							Abbr:   "CMT",
							Offset: -16356,
						},
						{
							ZoneId: 60,
							Abbr:   "BST",
							Offset: -12756,
						},
						{
							ZoneId: 60,
							Abbr:   "BOT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "BQ",
			Name: "Bonaire, Sint Eustatius and Saba",
			Zones: []*Zone{
				{
					Id:          61,
					CountryCode: "BQ",
					Name:        "America/Kralendijk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 61,
							Abbr:   "LMT",
							Offset: -16547,
						},
						{
							ZoneId: 61,
							Abbr:   "ANT",
							Offset: -16200,
						},
						{
							ZoneId: 61,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "BA",
			Name: "Bosnia and Herzegovina",
			Zones: []*Zone{
				{
					Id:          48,
					CountryCode: "BA",
					Name:        "Europe/Sarajevo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 48,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 48,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 48,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "BW",
			Name: "Botswana",
			Zones: []*Zone{
				{
					Id:          80,
					CountryCode: "BW",
					Name:        "Africa/Gaborone",
					TimeZones: []*TimeZone{
						{
							ZoneId: 80,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 80,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code:  "BV",
			Name:  "Bouvet Island",
			Zones: []*Zone{},
		},
		{
			Code: "BR",
			Name: "Brazil",
			Zones: []*Zone{
				{
					Id:          66,
					CountryCode: "BR",
					Name:        "America/Araguaina",
					TimeZones: []*TimeZone{
						{
							ZoneId: 66,
							Abbr:   "LMT",
							Offset: -11568,
						},
						{
							ZoneId: 66,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 66,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          68,
					CountryCode: "BR",
					Name:        "America/Bahia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 68,
							Abbr:   "LMT",
							Offset: -9244,
						},
						{
							ZoneId: 68,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 68,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          63,
					CountryCode: "BR",
					Name:        "America/Belem",
					TimeZones: []*TimeZone{
						{
							ZoneId: 63,
							Abbr:   "LMT",
							Offset: -11636,
						},
						{
							ZoneId: 63,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 63,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          74,
					CountryCode: "BR",
					Name:        "America/Boa_Vista",
					TimeZones: []*TimeZone{
						{
							ZoneId: 74,
							Abbr:   "LMT",
							Offset: -14560,
						},
						{
							ZoneId: 74,
							Abbr:   "AMT",
							Offset: -14400,
						},
						{
							ZoneId: 74,
							Abbr:   "AMST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          70,
					CountryCode: "BR",
					Name:        "America/Campo_Grande",
					TimeZones: []*TimeZone{
						{
							ZoneId: 70,
							Abbr:   "LMT",
							Offset: -13108,
						},
						{
							ZoneId: 70,
							Abbr:   "AMT",
							Offset: -14400,
						},
						{
							ZoneId: 70,
							Abbr:   "AMST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          71,
					CountryCode: "BR",
					Name:        "America/Cuiaba",
					TimeZones: []*TimeZone{
						{
							ZoneId: 71,
							Abbr:   "LMT",
							Offset: -13460,
						},
						{
							ZoneId: 71,
							Abbr:   "BRT",
							Offset: -14400,
						},
						{
							ZoneId: 71,
							Abbr:   "BRST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          76,
					CountryCode: "BR",
					Name:        "America/Eirunepe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 76,
							Abbr:   "LMT",
							Offset: -16768,
						},
						{
							ZoneId: 76,
							Abbr:   "ACT",
							Offset: -18000,
						},
						{
							ZoneId: 76,
							Abbr:   "AMT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          64,
					CountryCode: "BR",
					Name:        "America/Fortaleza",
					TimeZones: []*TimeZone{
						{
							ZoneId: 64,
							Abbr:   "LMT",
							Offset: -9240,
						},
						{
							ZoneId: 64,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 64,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          67,
					CountryCode: "BR",
					Name:        "America/Maceio",
					TimeZones: []*TimeZone{
						{
							ZoneId: 67,
							Abbr:   "LMT",
							Offset: -8572,
						},
						{
							ZoneId: 67,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 67,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          75,
					CountryCode: "BR",
					Name:        "America/Manaus",
					TimeZones: []*TimeZone{
						{
							ZoneId: 75,
							Abbr:   "LMT",
							Offset: -14404,
						},
						{
							ZoneId: 75,
							Abbr:   "AMT",
							Offset: -14400,
						},
						{
							ZoneId: 75,
							Abbr:   "AMST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          62,
					CountryCode: "BR",
					Name:        "America/Noronha",
					TimeZones: []*TimeZone{
						{
							ZoneId: 62,
							Abbr:   "LMT",
							Offset: -7780,
						},
						{
							ZoneId: 62,
							Abbr:   "FNT",
							Offset: -7200,
						},
						{
							ZoneId: 62,
							Abbr:   "FNST",
							Offset: -3600,
						},
					},
				},
				{
					Id:          73,
					CountryCode: "BR",
					Name:        "America/Porto_Velho",
					TimeZones: []*TimeZone{
						{
							ZoneId: 73,
							Abbr:   "LMT",
							Offset: -15336,
						},
						{
							ZoneId: 73,
							Abbr:   "AMT",
							Offset: -14400,
						},
						{
							ZoneId: 73,
							Abbr:   "AMST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          65,
					CountryCode: "BR",
					Name:        "America/Recife",
					TimeZones: []*TimeZone{
						{
							ZoneId: 65,
							Abbr:   "LMT",
							Offset: -8376,
						},
						{
							ZoneId: 65,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 65,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          77,
					CountryCode: "BR",
					Name:        "America/Rio_Branco",
					TimeZones: []*TimeZone{
						{
							ZoneId: 77,
							Abbr:   "LMT",
							Offset: -16272,
						},
						{
							ZoneId: 77,
							Abbr:   "ACT",
							Offset: -18000,
						},
						{
							ZoneId: 77,
							Abbr:   "AMT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          72,
					CountryCode: "BR",
					Name:        "America/Santarem",
					TimeZones: []*TimeZone{
						{
							ZoneId: 72,
							Abbr:   "LMT",
							Offset: -13128,
						},
						{
							ZoneId: 72,
							Abbr:   "BRT",
							Offset: -14400,
						},
						{
							ZoneId: 72,
							Abbr:   "BRST",
							Offset: -10800,
						},
						{
							ZoneId: 72,
							Abbr:   "BRT",
							Offset: -10800,
						},
					},
				},
				{
					Id:          69,
					CountryCode: "BR",
					Name:        "America/Sao_Paulo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 69,
							Abbr:   "LMT",
							Offset: -11188,
						},
						{
							ZoneId: 69,
							Abbr:   "BRT",
							Offset: -10800,
						},
						{
							ZoneId: 69,
							Abbr:   "BRST",
							Offset: -7200,
						},
					},
				},
			},
		},
		{
			Code: "IO",
			Name: "British Indian Ocean Territory",
			Zones: []*Zone{
				{
					Id:          194,
					CountryCode: "IO",
					Name:        "Indian/Chagos",
					TimeZones: []*TimeZone{
						{
							ZoneId: 194,
							Abbr:   "LMT",
							Offset: 17380,
						},
						{
							ZoneId: 194,
							Abbr:   "IOT",
							Offset: 18000,
						},
						{
							ZoneId: 194,
							Abbr:   "IOT",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "BN",
			Name: "Brunei Darussalam",
			Zones: []*Zone{
				{
					Id:          59,
					CountryCode: "BN",
					Name:        "Asia/Brunei",
					TimeZones: []*TimeZone{
						{
							ZoneId: 59,
							Abbr:   "LMT",
							Offset: 27580,
						},
						{
							ZoneId: 59,
							Abbr:   "BNT",
							Offset: 27000,
						},
						{
							ZoneId: 59,
							Abbr:   "BNT",
							Offset: 28800,
						},
					},
				},
			},
		},
		{
			Code: "BG",
			Name: "Bulgaria",
			Zones: []*Zone{
				{
					Id:          53,
					CountryCode: "BG",
					Name:        "Europe/Sofia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 53,
							Abbr:   "LMT",
							Offset: 5596,
						},
						{
							ZoneId: 53,
							Abbr:   "IMT",
							Offset: 7016,
						},
						{
							ZoneId: 53,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 53,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 53,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 53,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "BF",
			Name: "Burkina Faso",
			Zones: []*Zone{
				{
					Id:          52,
					CountryCode: "BF",
					Name:        "Africa/Ouagadougou",
					TimeZones: []*TimeZone{
						{
							ZoneId: 52,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 52,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "BI",
			Name: "Burundi",
			Zones: []*Zone{
				{
					Id:          55,
					CountryCode: "BI",
					Name:        "Africa/Bujumbura",
					TimeZones: []*TimeZone{
						{
							ZoneId: 55,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 55,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "KH",
			Name: "Cambodia",
			Zones: []*Zone{
				{
					Id:          205,
					CountryCode: "KH",
					Name:        "Asia/Phnom_Penh",
					TimeZones: []*TimeZone{
						{
							ZoneId: 205,
							Abbr:   "LMT",
							Offset: 24124,
						},
						{
							ZoneId: 205,
							Abbr:   "BMT",
							Offset: 24124,
						},
						{
							ZoneId: 205,
							Abbr:   "ICT",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "CM",
			Name: "Cameroon",
			Zones: []*Zone{
				{
					Id:          122,
					CountryCode: "CM",
					Name:        "Africa/Douala",
					TimeZones: []*TimeZone{
						{
							ZoneId: 122,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 122,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 122,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 122,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "CA",
			Name: "Canada",
			Zones: []*Zone{
				{
					Id:          94,
					CountryCode: "CA",
					Name:        "America/Atikokan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 94,
							Abbr:   "LMT",
							Offset: -21988,
						},
						{
							ZoneId: 94,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 94,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 94,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 94,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 94,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          88,
					CountryCode: "CA",
					Name:        "America/Blanc-Sablon",
					TimeZones: []*TimeZone{
						{
							ZoneId: 88,
							Abbr:   "LMT",
							Offset: -13708,
						},
						{
							ZoneId: 88,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 88,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 88,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 88,
							Abbr:   "APT",
							Offset: -10800,
						},
					},
				},
				{
					Id:          102,
					CountryCode: "CA",
					Name:        "America/Cambridge_Bay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 102,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 102,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 102,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 102,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 102,
							Abbr:   "MDDT",
							Offset: -18000,
						},
						{
							ZoneId: 102,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 102,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 102,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 102,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          105,
					CountryCode: "CA",
					Name:        "America/Creston",
					TimeZones: []*TimeZone{
						{
							ZoneId: 105,
							Abbr:   "LMT",
							Offset: -27964,
						},
						{
							ZoneId: 105,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 105,
							Abbr:   "PST",
							Offset: -28800,
						},
					},
				},
				{
					Id:          109,
					CountryCode: "CA",
					Name:        "America/Dawson",
					TimeZones: []*TimeZone{
						{
							ZoneId: 109,
							Abbr:   "LMT",
							Offset: -33460,
						},
						{
							ZoneId: 109,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 109,
							Abbr:   "YDT",
							Offset: -28800,
						},
						{
							ZoneId: 109,
							Abbr:   "YWT",
							Offset: -28800,
						},
						{
							ZoneId: 109,
							Abbr:   "YPT",
							Offset: -28800,
						},
						{
							ZoneId: 109,
							Abbr:   "YDDT",
							Offset: -25200,
						},
						{
							ZoneId: 109,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 109,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 109,
							Abbr:   "MST",
							Offset: -25200,
						},
					},
				},
				{
					Id:          106,
					CountryCode: "CA",
					Name:        "America/Dawson_Creek",
					TimeZones: []*TimeZone{
						{
							ZoneId: 106,
							Abbr:   "LMT",
							Offset: -28856,
						},
						{
							ZoneId: 106,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 106,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 106,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 106,
							Abbr:   "PPT",
							Offset: -25200,
						},
						{
							ZoneId: 106,
							Abbr:   "MST",
							Offset: -25200,
						},
					},
				},
				{
					Id:          101,
					CountryCode: "CA",
					Name:        "America/Edmonton",
					TimeZones: []*TimeZone{
						{
							ZoneId: 101,
							Abbr:   "LMT",
							Offset: -27232,
						},
						{
							ZoneId: 101,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 101,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 101,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 101,
							Abbr:   "MPT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          107,
					CountryCode: "CA",
					Name:        "America/Fort_Nelson",
					TimeZones: []*TimeZone{
						{
							ZoneId: 107,
							Abbr:   "LMT",
							Offset: -29447,
						},
						{
							ZoneId: 107,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 107,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 107,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 107,
							Abbr:   "PPT",
							Offset: -25200,
						},
						{
							ZoneId: 107,
							Abbr:   "MST",
							Offset: -25200,
						},
					},
				},
				{
					Id:          85,
					CountryCode: "CA",
					Name:        "America/Glace_Bay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 85,
							Abbr:   "LMT",
							Offset: -14388,
						},
						{
							ZoneId: 85,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 85,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 85,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 85,
							Abbr:   "APT",
							Offset: -10800,
						},
					},
				},
				{
					Id:          87,
					CountryCode: "CA",
					Name:        "America/Goose_Bay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 87,
							Abbr:   "LMT",
							Offset: -14500,
						},
						{
							ZoneId: 87,
							Abbr:   "NST",
							Offset: -12652,
						},
						{
							ZoneId: 87,
							Abbr:   "NDT",
							Offset: -9052,
						},
						{
							ZoneId: 87,
							Abbr:   "NST",
							Offset: -12600,
						},
						{
							ZoneId: 87,
							Abbr:   "NDT",
							Offset: -9000,
						},
						{
							ZoneId: 87,
							Abbr:   "NWT",
							Offset: -9000,
						},
						{
							ZoneId: 87,
							Abbr:   "NPT",
							Offset: -9000,
						},
						{
							ZoneId: 87,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 87,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 87,
							Abbr:   "ADDT",
							Offset: -7200,
						},
					},
				},
				{
					Id:          84,
					CountryCode: "CA",
					Name:        "America/Halifax",
					TimeZones: []*TimeZone{
						{
							ZoneId: 84,
							Abbr:   "LMT",
							Offset: -15264,
						},
						{
							ZoneId: 84,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 84,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 84,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 84,
							Abbr:   "APT",
							Offset: -10800,
						},
					},
				},
				{
					Id:          104,
					CountryCode: "CA",
					Name:        "America/Inuvik",
					TimeZones: []*TimeZone{
						{
							ZoneId: 104,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 104,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 104,
							Abbr:   "PDDT",
							Offset: -21600,
						},
						{
							ZoneId: 104,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 104,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          92,
					CountryCode: "CA",
					Name:        "America/Iqaluit",
					TimeZones: []*TimeZone{
						{
							ZoneId: 92,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 92,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 92,
							Abbr:   "EPT",
							Offset: -14400,
						},
						{
							ZoneId: 92,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 92,
							Abbr:   "EDDT",
							Offset: -10800,
						},
						{
							ZoneId: 92,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 92,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 92,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          86,
					CountryCode: "CA",
					Name:        "America/Moncton",
					TimeZones: []*TimeZone{
						{
							ZoneId: 86,
							Abbr:   "LMT",
							Offset: -15548,
						},
						{
							ZoneId: 86,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 86,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 86,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 86,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 86,
							Abbr:   "APT",
							Offset: -10800,
						},
					},
				},
				{
					Id:          90,
					CountryCode: "CA",
					Name:        "America/Nipigon",
					TimeZones: []*TimeZone{
						{
							ZoneId: 90,
							Abbr:   "LMT",
							Offset: -21184,
						},
						{
							ZoneId: 90,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 90,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 90,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 90,
							Abbr:   "EPT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          93,
					CountryCode: "CA",
					Name:        "America/Pangnirtung",
					TimeZones: []*TimeZone{
						{
							ZoneId: 93,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 93,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 93,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 93,
							Abbr:   "APT",
							Offset: -10800,
						},
						{
							ZoneId: 93,
							Abbr:   "ADDT",
							Offset: -7200,
						},
						{
							ZoneId: 93,
							Abbr:   "ADT",
							Offset: -10800,
						},
						{
							ZoneId: 93,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 93,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 93,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 93,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          96,
					CountryCode: "CA",
					Name:        "America/Rainy_River",
					TimeZones: []*TimeZone{
						{
							ZoneId: 96,
							Abbr:   "LMT",
							Offset: -22696,
						},
						{
							ZoneId: 96,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 96,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 96,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 96,
							Abbr:   "CPT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          98,
					CountryCode: "CA",
					Name:        "America/Rankin_Inlet",
					TimeZones: []*TimeZone{
						{
							ZoneId: 98,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 98,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 98,
							Abbr:   "CDDT",
							Offset: -14400,
						},
						{
							ZoneId: 98,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 98,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          99,
					CountryCode: "CA",
					Name:        "America/Regina",
					TimeZones: []*TimeZone{
						{
							ZoneId: 99,
							Abbr:   "LMT",
							Offset: -25116,
						},
						{
							ZoneId: 99,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 99,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 99,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 99,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 99,
							Abbr:   "CST",
							Offset: -21600,
						},
					},
				},
				{
					Id:          97,
					CountryCode: "CA",
					Name:        "America/Resolute",
					TimeZones: []*TimeZone{
						{
							ZoneId: 97,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 97,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 97,
							Abbr:   "CDDT",
							Offset: -14400,
						},
						{
							ZoneId: 97,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 97,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          83,
					CountryCode: "CA",
					Name:        "America/St_Johns",
					TimeZones: []*TimeZone{
						{
							ZoneId: 83,
							Abbr:   "LMT",
							Offset: -12652,
						},
						{
							ZoneId: 83,
							Abbr:   "NST",
							Offset: -12652,
						},
						{
							ZoneId: 83,
							Abbr:   "NDT",
							Offset: -9052,
						},
						{
							ZoneId: 83,
							Abbr:   "NST",
							Offset: -12600,
						},
						{
							ZoneId: 83,
							Abbr:   "NDT",
							Offset: -9000,
						},
						{
							ZoneId: 83,
							Abbr:   "NWT",
							Offset: -9000,
						},
						{
							ZoneId: 83,
							Abbr:   "NPT",
							Offset: -9000,
						},
						{
							ZoneId: 83,
							Abbr:   "NDDT",
							Offset: -5400,
						},
					},
				},
				{
					Id:          100,
					CountryCode: "CA",
					Name:        "America/Swift_Current",
					TimeZones: []*TimeZone{
						{
							ZoneId: 100,
							Abbr:   "LMT",
							Offset: -25880,
						},
						{
							ZoneId: 100,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 100,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 100,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 100,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 100,
							Abbr:   "CST",
							Offset: -21600,
						},
					},
				},
				{
					Id:          91,
					CountryCode: "CA",
					Name:        "America/Thunder_Bay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 91,
							Abbr:   "LMT",
							Offset: -21420,
						},
						{
							ZoneId: 91,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 91,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 91,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 91,
							Abbr:   "EPT",
							Offset: -14400,
						},
						{
							ZoneId: 91,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          89,
					CountryCode: "CA",
					Name:        "America/Toronto",
					TimeZones: []*TimeZone{
						{
							ZoneId: 89,
							Abbr:   "LMT",
							Offset: -19052,
						},
						{
							ZoneId: 89,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 89,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 89,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 89,
							Abbr:   "EPT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          110,
					CountryCode: "CA",
					Name:        "America/Vancouver",
					TimeZones: []*TimeZone{
						{
							ZoneId: 110,
							Abbr:   "LMT",
							Offset: -29548,
						},
						{
							ZoneId: 110,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 110,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 110,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 110,
							Abbr:   "PPT",
							Offset: -25200,
						},
					},
				},
				{
					Id:          108,
					CountryCode: "CA",
					Name:        "America/Whitehorse",
					TimeZones: []*TimeZone{
						{
							ZoneId: 108,
							Abbr:   "LMT",
							Offset: -32412,
						},
						{
							ZoneId: 108,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 108,
							Abbr:   "YDT",
							Offset: -28800,
						},
						{
							ZoneId: 108,
							Abbr:   "YWT",
							Offset: -28800,
						},
						{
							ZoneId: 108,
							Abbr:   "YPT",
							Offset: -28800,
						},
						{
							ZoneId: 108,
							Abbr:   "YDDT",
							Offset: -25200,
						},
						{
							ZoneId: 108,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 108,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 108,
							Abbr:   "MST",
							Offset: -25200,
						},
					},
				},
				{
					Id:          95,
					CountryCode: "CA",
					Name:        "America/Winnipeg",
					TimeZones: []*TimeZone{
						{
							ZoneId: 95,
							Abbr:   "LMT",
							Offset: -23316,
						},
						{
							ZoneId: 95,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 95,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 95,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 95,
							Abbr:   "CPT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          103,
					CountryCode: "CA",
					Name:        "America/Yellowknife",
					TimeZones: []*TimeZone{
						{
							ZoneId: 103,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 103,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 103,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 103,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 103,
							Abbr:   "MDDT",
							Offset: -18000,
						},
						{
							ZoneId: 103,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
			},
		},
		{
			Code: "CV",
			Name: "Cape Verde",
			Zones: []*Zone{
				{
					Id:          128,
					CountryCode: "CV",
					Name:        "Atlantic/Cape_Verde",
					TimeZones: []*TimeZone{
						{
							ZoneId: 128,
							Abbr:   "LMT",
							Offset: -5644,
						},
						{
							ZoneId: 128,
							Abbr:   "CVT",
							Offset: -7200,
						},
						{
							ZoneId: 128,
							Abbr:   "CVST",
							Offset: -3600,
						},
						{
							ZoneId: 128,
							Abbr:   "CVT",
							Offset: -3600,
						},
					},
				},
			},
		},
		{
			Code: "KY",
			Name: "Cayman Islands",
			Zones: []*Zone{
				{
					Id:          214,
					CountryCode: "KY",
					Name:        "America/Cayman",
					TimeZones: []*TimeZone{
						{
							ZoneId: 214,
							Abbr:   "LMT",
							Offset: -19088,
						},
						{
							ZoneId: 214,
							Abbr:   "CMT",
							Offset: -19176,
						},
						{
							ZoneId: 214,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "CF",
			Name: "Central African Republic",
			Zones: []*Zone{
				{
					Id:          114,
					CountryCode: "CF",
					Name:        "Africa/Bangui",
					TimeZones: []*TimeZone{
						{
							ZoneId: 114,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 114,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 114,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 114,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "TD",
			Name: "Chad",
			Zones: []*Zone{
				{
					Id:          358,
					CountryCode: "TD",
					Name:        "Africa/Ndjamena",
					TimeZones: []*TimeZone{
						{
							ZoneId: 358,
							Abbr:   "LMT",
							Offset: 3612,
						},
						{
							ZoneId: 358,
							Abbr:   "WAT",
							Offset: 3600,
						},
						{
							ZoneId: 358,
							Abbr:   "WAST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "CL",
			Name: "Chile",
			Zones: []*Zone{
				{
					Id:          120,
					CountryCode: "CL",
					Name:        "America/Punta_Arenas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 120,
							Abbr:   "LMT",
							Offset: -17020,
						},
						{
							ZoneId: 120,
							Abbr:   "SMT",
							Offset: -16966,
						},
						{
							ZoneId: 120,
							Abbr:   "CLT",
							Offset: -18000,
						},
						{
							ZoneId: 120,
							Abbr:   "CLT",
							Offset: -14400,
						},
						{
							ZoneId: 120,
							Abbr:   "CLST",
							Offset: -14400,
						},
						{
							ZoneId: 120,
							Abbr:   "CLST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          119,
					CountryCode: "CL",
					Name:        "America/Santiago",
					TimeZones: []*TimeZone{
						{
							ZoneId: 119,
							Abbr:   "LMT",
							Offset: -16966,
						},
						{
							ZoneId: 119,
							Abbr:   "SMT",
							Offset: -16966,
						},
						{
							ZoneId: 119,
							Abbr:   "CLT",
							Offset: -18000,
						},
						{
							ZoneId: 119,
							Abbr:   "CLT",
							Offset: -14400,
						},
						{
							ZoneId: 119,
							Abbr:   "CLST",
							Offset: -14400,
						},
						{
							ZoneId: 119,
							Abbr:   "CLST",
							Offset: -10800,
						},
					},
				},
				{
					Id:          121,
					CountryCode: "CL",
					Name:        "Pacific/Easter",
					TimeZones: []*TimeZone{
						{
							ZoneId: 121,
							Abbr:   "LMT",
							Offset: -26248,
						},
						{
							ZoneId: 121,
							Abbr:   "EMT",
							Offset: -26248,
						},
						{
							ZoneId: 121,
							Abbr:   "EAST",
							Offset: -25200,
						},
						{
							ZoneId: 121,
							Abbr:   "EASST",
							Offset: -21600,
						},
						{
							ZoneId: 121,
							Abbr:   "EASST",
							Offset: -18000,
						},
						{
							ZoneId: 121,
							Abbr:   "EAST",
							Offset: -21600,
						},
					},
				},
			},
		},
		{
			Code: "CN",
			Name: "China",
			Zones: []*Zone{
				{
					Id:          123,
					CountryCode: "CN",
					Name:        "Asia/Shanghai",
					TimeZones: []*TimeZone{
						{
							ZoneId: 123,
							Abbr:   "LMT",
							Offset: 29143,
						},
						{
							ZoneId: 123,
							Abbr:   "CST",
							Offset: 28800,
						},
						{
							ZoneId: 123,
							Abbr:   "CDT",
							Offset: 32400,
						},
					},
				},
				{
					Id:          124,
					CountryCode: "CN",
					Name:        "Asia/Urumqi",
					TimeZones: []*TimeZone{
						{
							ZoneId: 124,
							Abbr:   "LMT",
							Offset: 21020,
						},
						{
							ZoneId: 124,
							Abbr:   "URUT",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "CX",
			Name: "Christmas Island",
			Zones: []*Zone{
				{
					Id:          130,
					CountryCode: "CX",
					Name:        "Indian/Christmas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 130,
							Abbr:   "LMT",
							Offset: 25372,
						},
						{
							ZoneId: 130,
							Abbr:   "CXT",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "CC",
			Name: "Cocos (Keeling) Islands",
			Zones: []*Zone{
				{
					Id:          111,
					CountryCode: "CC",
					Name:        "Indian/Cocos",
					TimeZones: []*TimeZone{
						{
							ZoneId: 111,
							Abbr:   "LMT",
							Offset: 23260,
						},
						{
							ZoneId: 111,
							Abbr:   "CCT",
							Offset: 23400,
						},
					},
				},
			},
		},
		{
			Code: "CO",
			Name: "Colombia",
			Zones: []*Zone{
				{
					Id:          125,
					CountryCode: "CO",
					Name:        "America/Bogota",
					TimeZones: []*TimeZone{
						{
							ZoneId: 125,
							Abbr:   "LMT",
							Offset: -17776,
						},
						{
							ZoneId: 125,
							Abbr:   "BMT",
							Offset: -17776,
						},
						{
							ZoneId: 125,
							Abbr:   "COT",
							Offset: -18000,
						},
						{
							ZoneId: 125,
							Abbr:   "COST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "KM",
			Name: "Comoros",
			Zones: []*Zone{
				{
					Id:          209,
					CountryCode: "KM",
					Name:        "Indian/Comoro",
					TimeZones: []*TimeZone{
						{
							ZoneId: 209,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 209,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 209,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 209,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "CG",
			Name: "Congo",
			Zones: []*Zone{
				{
					Id:          115,
					CountryCode: "CG",
					Name:        "Africa/Brazzaville",
					TimeZones: []*TimeZone{
						{
							ZoneId: 115,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 115,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 115,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 115,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "CD",
			Name: "Congo, the Democratic Republic of the",
			Zones: []*Zone{
				{
					Id:          112,
					CountryCode: "CD",
					Name:        "Africa/Kinshasa",
					TimeZones: []*TimeZone{
						{
							ZoneId: 112,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 112,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 112,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 112,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
				{
					Id:          113,
					CountryCode: "CD",
					Name:        "Africa/Lubumbashi",
					TimeZones: []*TimeZone{
						{
							ZoneId: 113,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 113,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "CK",
			Name: "Cook Islands",
			Zones: []*Zone{
				{
					Id:          118,
					CountryCode: "CK",
					Name:        "Pacific/Rarotonga",
					TimeZones: []*TimeZone{
						{
							ZoneId: 118,
							Abbr:   "LMT",
							Offset: -38344,
						},
						{
							ZoneId: 118,
							Abbr:   "CKT",
							Offset: -37800,
						},
						{
							ZoneId: 118,
							Abbr:   "CKHST",
							Offset: -34200,
						},
						{
							ZoneId: 118,
							Abbr:   "CKT",
							Offset: -36000,
						},
					},
				},
			},
		},
		{
			Code: "CR",
			Name: "Costa Rica",
			Zones: []*Zone{
				{
					Id:          126,
					CountryCode: "CR",
					Name:        "America/Costa_Rica",
					TimeZones: []*TimeZone{
						{
							ZoneId: 126,
							Abbr:   "LMT",
							Offset: -20173,
						},
						{
							ZoneId: 126,
							Abbr:   "SJMT",
							Offset: -20173,
						},
						{
							ZoneId: 126,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 126,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "HR",
			Name: "Croatia",
			Zones: []*Zone{
				{
					Id:          183,
					CountryCode: "HR",
					Name:        "Europe/Zagreb",
					TimeZones: []*TimeZone{
						{
							ZoneId: 183,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 183,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 183,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "CU",
			Name: "Cuba",
			Zones: []*Zone{
				{
					Id:          127,
					CountryCode: "CU",
					Name:        "America/Havana",
					TimeZones: []*TimeZone{
						{
							ZoneId: 127,
							Abbr:   "LMT",
							Offset: -19768,
						},
						{
							ZoneId: 127,
							Abbr:   "HMT",
							Offset: -19776,
						},
						{
							ZoneId: 127,
							Abbr:   "CST",
							Offset: -18000,
						},
						{
							ZoneId: 127,
							Abbr:   "CDT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "CW",
			Name: "Curaçao",
			Zones: []*Zone{
				{
					Id:          129,
					CountryCode: "CW",
					Name:        "America/Curacao",
					TimeZones: []*TimeZone{
						{
							ZoneId: 129,
							Abbr:   "LMT",
							Offset: -16547,
						},
						{
							ZoneId: 129,
							Abbr:   "AST",
							Offset: -16200,
						},
						{
							ZoneId: 129,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "CY",
			Name: "Cyprus",
			Zones: []*Zone{
				{
					Id:          132,
					CountryCode: "CY",
					Name:        "Asia/Famagusta",
					TimeZones: []*TimeZone{
						{
							ZoneId: 132,
							Abbr:   "LMT",
							Offset: 8148,
						},
						{
							ZoneId: 132,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 132,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 132,
							Abbr:   "EET",
							Offset: 10800,
						},
					},
				},
				{
					Id:          131,
					CountryCode: "CY",
					Name:        "Asia/Nicosia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 131,
							Abbr:   "LMT",
							Offset: 8008,
						},
						{
							ZoneId: 131,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 131,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "CZ",
			Name: "Czech Republic",
			Zones: []*Zone{
				{
					Id:          133,
					CountryCode: "CZ",
					Name:        "Europe/Prague",
					TimeZones: []*TimeZone{
						{
							ZoneId: 133,
							Abbr:   "LMT",
							Offset: 3464,
						},
						{
							ZoneId: 133,
							Abbr:   "PMT",
							Offset: 3464,
						},
						{
							ZoneId: 133,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 133,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 133,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "CI",
			Name: "Côte d'Ivoire",
			Zones: []*Zone{
				{
					Id:          117,
					CountryCode: "CI",
					Name:        "Africa/Abidjan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 117,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 117,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "DK",
			Name: "Denmark",
			Zones: []*Zone{
				{
					Id:          137,
					CountryCode: "DK",
					Name:        "Europe/Copenhagen",
					TimeZones: []*TimeZone{
						{
							ZoneId: 137,
							Abbr:   "LMT",
							Offset: 3020,
						},
						{
							ZoneId: 137,
							Abbr:   "CMT",
							Offset: 3020,
						},
						{
							ZoneId: 137,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 137,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "DJ",
			Name: "Djibouti",
			Zones: []*Zone{
				{
					Id:          136,
					CountryCode: "DJ",
					Name:        "Africa/Djibouti",
					TimeZones: []*TimeZone{
						{
							ZoneId: 136,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 136,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 136,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 136,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "DM",
			Name: "Dominica",
			Zones: []*Zone{
				{
					Id:          138,
					CountryCode: "DM",
					Name:        "America/Dominica",
					TimeZones: []*TimeZone{
						{
							ZoneId: 138,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 138,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "DO",
			Name: "Dominican Republic",
			Zones: []*Zone{
				{
					Id:          139,
					CountryCode: "DO",
					Name:        "America/Santo_Domingo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 139,
							Abbr:   "LMT",
							Offset: -16776,
						},
						{
							ZoneId: 139,
							Abbr:   "SDMT",
							Offset: -16800,
						},
						{
							ZoneId: 139,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 139,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 139,
							Abbr:   "EHDT",
							Offset: -16200,
						},
						{
							ZoneId: 139,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "EC",
			Name: "Ecuador",
			Zones: []*Zone{
				{
					Id:          141,
					CountryCode: "EC",
					Name:        "America/Guayaquil",
					TimeZones: []*TimeZone{
						{
							ZoneId: 141,
							Abbr:   "LMT",
							Offset: -19160,
						},
						{
							ZoneId: 141,
							Abbr:   "QMT",
							Offset: -18840,
						},
						{
							ZoneId: 141,
							Abbr:   "ECT",
							Offset: -18000,
						},
						{
							ZoneId: 141,
							Abbr:   "ECT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          142,
					CountryCode: "EC",
					Name:        "Pacific/Galapagos",
					TimeZones: []*TimeZone{
						{
							ZoneId: 142,
							Abbr:   "LMT",
							Offset: -21504,
						},
						{
							ZoneId: 142,
							Abbr:   "GALT",
							Offset: -18000,
						},
						{
							ZoneId: 142,
							Abbr:   "GALT",
							Offset: -21600,
						},
					},
				},
			},
		},
		{
			Code: "EG",
			Name: "Egypt",
			Zones: []*Zone{
				{
					Id:          144,
					CountryCode: "EG",
					Name:        "Africa/Cairo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 144,
							Abbr:   "LMT",
							Offset: 7509,
						},
						{
							ZoneId: 144,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 144,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "SV",
			Name: "El Salvador",
			Zones: []*Zone{
				{
					Id:          353,
					CountryCode: "SV",
					Name:        "America/El_Salvador",
					TimeZones: []*TimeZone{
						{
							ZoneId: 353,
							Abbr:   "LMT",
							Offset: -21408,
						},
						{
							ZoneId: 353,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 353,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "GQ",
			Name: "Equatorial Guinea",
			Zones: []*Zone{
				{
					Id:          174,
					CountryCode: "GQ",
					Name:        "Africa/Malabo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 174,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 174,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 174,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 174,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "ER",
			Name: "Eritrea",
			Zones: []*Zone{
				{
					Id:          146,
					CountryCode: "ER",
					Name:        "Africa/Asmara",
					TimeZones: []*TimeZone{
						{
							ZoneId: 146,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 146,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 146,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 146,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "EE",
			Name: "Estonia",
			Zones: []*Zone{
				{
					Id:          143,
					CountryCode: "EE",
					Name:        "Europe/Tallinn",
					TimeZones: []*TimeZone{
						{
							ZoneId: 143,
							Abbr:   "LMT",
							Offset: 5940,
						},
						{
							ZoneId: 143,
							Abbr:   "TMT",
							Offset: 5940,
						},
						{
							ZoneId: 143,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 143,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 143,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 143,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 143,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 143,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "ET",
			Name: "Ethiopia",
			Zones: []*Zone{
				{
					Id:          150,
					CountryCode: "ET",
					Name:        "Africa/Addis_Ababa",
					TimeZones: []*TimeZone{
						{
							ZoneId: 150,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 150,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 150,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 150,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "FK",
			Name: "Falkland Islands (Malvinas)",
			Zones: []*Zone{
				{
					Id:          153,
					CountryCode: "FK",
					Name:        "Atlantic/Stanley",
					TimeZones: []*TimeZone{
						{
							ZoneId: 153,
							Abbr:   "LMT",
							Offset: -13884,
						},
						{
							ZoneId: 153,
							Abbr:   "SMT",
							Offset: -13884,
						},
						{
							ZoneId: 153,
							Abbr:   "FKT",
							Offset: -14400,
						},
						{
							ZoneId: 153,
							Abbr:   "FKST",
							Offset: -10800,
						},
						{
							ZoneId: 153,
							Abbr:   "FKT",
							Offset: -10800,
						},
						{
							ZoneId: 153,
							Abbr:   "FKST",
							Offset: -7200,
						},
					},
				},
			},
		},
		{
			Code: "FO",
			Name: "Faroe Islands",
			Zones: []*Zone{
				{
					Id:          157,
					CountryCode: "FO",
					Name:        "Atlantic/Faroe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 157,
							Abbr:   "LMT",
							Offset: -1624,
						},
						{
							ZoneId: 157,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 157,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "FJ",
			Name: "Fiji",
			Zones: []*Zone{
				{
					Id:          152,
					CountryCode: "FJ",
					Name:        "Pacific/Fiji",
					TimeZones: []*TimeZone{
						{
							ZoneId: 152,
							Abbr:   "LMT",
							Offset: 42944,
						},
						{
							ZoneId: 152,
							Abbr:   "FJT",
							Offset: 43200,
						},
						{
							ZoneId: 152,
							Abbr:   "FJT",
							Offset: 46800,
						},
					},
				},
			},
		},
		{
			Code: "FI",
			Name: "Finland",
			Zones: []*Zone{
				{
					Id:          151,
					CountryCode: "FI",
					Name:        "Europe/Helsinki",
					TimeZones: []*TimeZone{
						{
							ZoneId: 151,
							Abbr:   "LMT",
							Offset: 5989,
						},
						{
							ZoneId: 151,
							Abbr:   "HMT",
							Offset: 5989,
						},
						{
							ZoneId: 151,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 151,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "FR",
			Name: "France",
			Zones: []*Zone{
				{
					Id:          158,
					CountryCode: "FR",
					Name:        "Europe/Paris",
					TimeZones: []*TimeZone{
						{
							ZoneId: 158,
							Abbr:   "LMT",
							Offset: 561,
						},
						{
							ZoneId: 158,
							Abbr:   "PMT",
							Offset: 561,
						},
						{
							ZoneId: 158,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 158,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 158,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 158,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 158,
							Abbr:   "WEMT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "GF",
			Name: "French Guiana",
			Zones: []*Zone{
				{
					Id:          163,
					CountryCode: "GF",
					Name:        "America/Cayenne",
					TimeZones: []*TimeZone{
						{
							ZoneId: 163,
							Abbr:   "LMT",
							Offset: -12560,
						},
						{
							ZoneId: 163,
							Abbr:   "GFT",
							Offset: -14400,
						},
						{
							ZoneId: 163,
							Abbr:   "GFT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "PF",
			Name: "French Polynesia",
			Zones: []*Zone{
				{
					Id:          288,
					CountryCode: "PF",
					Name:        "Pacific/Gambier",
					TimeZones: []*TimeZone{
						{
							ZoneId: 288,
							Abbr:   "LMT",
							Offset: -32388,
						},
						{
							ZoneId: 288,
							Abbr:   "GAMT",
							Offset: -32400,
						},
					},
				},
				{
					Id:          287,
					CountryCode: "PF",
					Name:        "Pacific/Marquesas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 287,
							Abbr:   "LMT",
							Offset: -33480,
						},
						{
							ZoneId: 287,
							Abbr:   "MART",
							Offset: -34200,
						},
					},
				},
				{
					Id:          286,
					CountryCode: "PF",
					Name:        "Pacific/Tahiti",
					TimeZones: []*TimeZone{
						{
							ZoneId: 286,
							Abbr:   "LMT",
							Offset: -35896,
						},
						{
							ZoneId: 286,
							Abbr:   "TAHT",
							Offset: -36000,
						},
					},
				},
			},
		},
		{
			Code: "TF",
			Name: "French Southern Territories",
			Zones: []*Zone{
				{
					Id:          359,
					CountryCode: "TF",
					Name:        "Indian/Kerguelen",
					TimeZones: []*TimeZone{
						{
							ZoneId: 359,
							Abbr:   "TFT",
							Offset: 0,
						},
						{
							ZoneId: 359,
							Abbr:   "TFT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "GA",
			Name: "Gabon",
			Zones: []*Zone{
				{
					Id:          159,
					CountryCode: "GA",
					Name:        "Africa/Libreville",
					TimeZones: []*TimeZone{
						{
							ZoneId: 159,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 159,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 159,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 159,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "GM",
			Name: "Gambia",
			Zones: []*Zone{
				{
					Id:          171,
					CountryCode: "GM",
					Name:        "Africa/Banjul",
					TimeZones: []*TimeZone{
						{
							ZoneId: 171,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 171,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "GE",
			Name: "Georgia",
			Zones: []*Zone{
				{
					Id:          162,
					CountryCode: "GE",
					Name:        "Asia/Tbilisi",
					TimeZones: []*TimeZone{
						{
							ZoneId: 162,
							Abbr:   "LMT",
							Offset: 10751,
						},
						{
							ZoneId: 162,
							Abbr:   "TBMT",
							Offset: 10751,
						},
						{
							ZoneId: 162,
							Abbr:   "GET",
							Offset: 10800,
						},
						{
							ZoneId: 162,
							Abbr:   "GET",
							Offset: 14400,
						},
						{
							ZoneId: 162,
							Abbr:   "GEST",
							Offset: 18000,
						},
						{
							ZoneId: 162,
							Abbr:   "GEST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "DE",
			Name: "Germany",
			Zones: []*Zone{
				{
					Id:          134,
					CountryCode: "DE",
					Name:        "Europe/Berlin",
					TimeZones: []*TimeZone{
						{
							ZoneId: 134,
							Abbr:   "LMT",
							Offset: 3208,
						},
						{
							ZoneId: 134,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 134,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 134,
							Abbr:   "CEMT",
							Offset: 10800,
						},
					},
				},
				{
					Id:          135,
					CountryCode: "DE",
					Name:        "Europe/Busingen",
					TimeZones: []*TimeZone{
						{
							ZoneId: 135,
							Abbr:   "LMT",
							Offset: 2048,
						},
						{
							ZoneId: 135,
							Abbr:   "BMT",
							Offset: 1786,
						},
						{
							ZoneId: 135,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 135,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "GH",
			Name: "Ghana",
			Zones: []*Zone{
				{
					Id:          165,
					CountryCode: "GH",
					Name:        "Africa/Accra",
					TimeZones: []*TimeZone{
						{
							ZoneId: 165,
							Abbr:   "LMT",
							Offset: -52,
						},
						{
							ZoneId: 165,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 165,
							Abbr:   "GHST",
							Offset: 1200,
						},
						{
							ZoneId: 165,
							Abbr:   "GMT",
							Offset: 1800,
						},
						{
							ZoneId: 165,
							Abbr:   "GHST",
							Offset: 1800,
						},
					},
				},
			},
		},
		{
			Code: "GI",
			Name: "Gibraltar",
			Zones: []*Zone{
				{
					Id:          166,
					CountryCode: "GI",
					Name:        "Europe/Gibraltar",
					TimeZones: []*TimeZone{
						{
							ZoneId: 166,
							Abbr:   "LMT",
							Offset: -1284,
						},
						{
							ZoneId: 166,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 166,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 166,
							Abbr:   "BDST",
							Offset: 7200,
						},
						{
							ZoneId: 166,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 166,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "GR",
			Name: "Greece",
			Zones: []*Zone{
				{
					Id:          175,
					CountryCode: "GR",
					Name:        "Europe/Athens",
					TimeZones: []*TimeZone{
						{
							ZoneId: 175,
							Abbr:   "LMT",
							Offset: 5692,
						},
						{
							ZoneId: 175,
							Abbr:   "AMT",
							Offset: 5692,
						},
						{
							ZoneId: 175,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 175,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 175,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 175,
							Abbr:   "CET",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "GL",
			Name: "Greenland",
			Zones: []*Zone{
				{
					Id:          168,
					CountryCode: "GL",
					Name:        "America/Danmarkshavn",
					TimeZones: []*TimeZone{
						{
							ZoneId: 168,
							Abbr:   "LMT",
							Offset: -4480,
						},
						{
							ZoneId: 168,
							Abbr:   "WGT",
							Offset: -10800,
						},
						{
							ZoneId: 168,
							Abbr:   "WGST",
							Offset: -7200,
						},
						{
							ZoneId: 168,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
				{
					Id:          167,
					CountryCode: "GL",
					Name:        "America/Nuuk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 167,
							Abbr:   "LMT",
							Offset: -12416,
						},
						{
							ZoneId: 167,
							Abbr:   "WGT",
							Offset: -10800,
						},
						{
							ZoneId: 167,
							Abbr:   "WGST",
							Offset: -7200,
						},
					},
				},
				{
					Id:          169,
					CountryCode: "GL",
					Name:        "America/Scoresbysund",
					TimeZones: []*TimeZone{
						{
							ZoneId: 169,
							Abbr:   "LMT",
							Offset: -5272,
						},
						{
							ZoneId: 169,
							Abbr:   "CGT",
							Offset: -7200,
						},
						{
							ZoneId: 169,
							Abbr:   "EGT",
							Offset: -3600,
						},
						{
							ZoneId: 169,
							Abbr:   "EGST",
							Offset: 0,
						},
					},
				},
				{
					Id:          170,
					CountryCode: "GL",
					Name:        "America/Thule",
					TimeZones: []*TimeZone{
						{
							ZoneId: 170,
							Abbr:   "LMT",
							Offset: -16508,
						},
						{
							ZoneId: 170,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 170,
							Abbr:   "ADT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "GD",
			Name: "Grenada",
			Zones: []*Zone{
				{
					Id:          161,
					CountryCode: "GD",
					Name:        "America/Grenada",
					TimeZones: []*TimeZone{
						{
							ZoneId: 161,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 161,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "GP",
			Name: "Guadeloupe",
			Zones: []*Zone{
				{
					Id:          173,
					CountryCode: "GP",
					Name:        "America/Guadeloupe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 173,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 173,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "GU",
			Name: "Guam",
			Zones: []*Zone{
				{
					Id:          178,
					CountryCode: "GU",
					Name:        "Pacific/Guam",
					TimeZones: []*TimeZone{
						{
							ZoneId: 178,
							Abbr:   "LMT",
							Offset: -51660,
						},
						{
							ZoneId: 178,
							Abbr:   "LMT",
							Offset: 34740,
						},
						{
							ZoneId: 178,
							Abbr:   "GST",
							Offset: 36000,
						},
						{
							ZoneId: 178,
							Abbr:   "GST",
							Offset: 32400,
						},
						{
							ZoneId: 178,
							Abbr:   "GDT",
							Offset: 39600,
						},
						{
							ZoneId: 178,
							Abbr:   "ChST",
							Offset: 36000,
						},
					},
				},
			},
		},
		{
			Code: "GT",
			Name: "Guatemala",
			Zones: []*Zone{
				{
					Id:          177,
					CountryCode: "GT",
					Name:        "America/Guatemala",
					TimeZones: []*TimeZone{
						{
							ZoneId: 177,
							Abbr:   "LMT",
							Offset: -21724,
						},
						{
							ZoneId: 177,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 177,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "GG",
			Name: "Guernsey",
			Zones: []*Zone{
				{
					Id:          164,
					CountryCode: "GG",
					Name:        "Europe/Guernsey",
					TimeZones: []*TimeZone{
						{
							ZoneId: 164,
							Abbr:   "LMT",
							Offset: -75,
						},
						{
							ZoneId: 164,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 164,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 164,
							Abbr:   "BDST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "GN",
			Name: "Guinea",
			Zones: []*Zone{
				{
					Id:          172,
					CountryCode: "GN",
					Name:        "Africa/Conakry",
					TimeZones: []*TimeZone{
						{
							ZoneId: 172,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 172,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "GW",
			Name: "Guinea-Bissau",
			Zones: []*Zone{
				{
					Id:          179,
					CountryCode: "GW",
					Name:        "Africa/Bissau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 179,
							Abbr:   "LMT",
							Offset: -3740,
						},
						{
							ZoneId: 179,
							Abbr:   "WAT",
							Offset: -3600,
						},
						{
							ZoneId: 179,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "GY",
			Name: "Guyana",
			Zones: []*Zone{
				{
					Id:          180,
					CountryCode: "GY",
					Name:        "America/Guyana",
					TimeZones: []*TimeZone{
						{
							ZoneId: 180,
							Abbr:   "LMT",
							Offset: -13960,
						},
						{
							ZoneId: 180,
							Abbr:   "GYT",
							Offset: -13500,
						},
						{
							ZoneId: 180,
							Abbr:   "GYT",
							Offset: -10800,
						},
						{
							ZoneId: 180,
							Abbr:   "GYT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "HT",
			Name: "Haiti",
			Zones: []*Zone{
				{
					Id:          184,
					CountryCode: "HT",
					Name:        "America/Port-au-Prince",
					TimeZones: []*TimeZone{
						{
							ZoneId: 184,
							Abbr:   "LMT",
							Offset: -17360,
						},
						{
							ZoneId: 184,
							Abbr:   "PPMT",
							Offset: -17340,
						},
						{
							ZoneId: 184,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 184,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code:  "HM",
			Name:  "Heard Island and McDonald Islands",
			Zones: []*Zone{},
		},
		{
			Code: "VA",
			Name: "Holy See (Vatican City State)",
			Zones: []*Zone{
				{
					Id:          411,
					CountryCode: "VA",
					Name:        "Europe/Vatican",
					TimeZones: []*TimeZone{
						{
							ZoneId: 411,
							Abbr:   "LMT",
							Offset: 2996,
						},
						{
							ZoneId: 411,
							Abbr:   "RMT",
							Offset: 2996,
						},
						{
							ZoneId: 411,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 411,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "HN",
			Name: "Honduras",
			Zones: []*Zone{
				{
					Id:          182,
					CountryCode: "HN",
					Name:        "America/Tegucigalpa",
					TimeZones: []*TimeZone{
						{
							ZoneId: 182,
							Abbr:   "LMT",
							Offset: -20932,
						},
						{
							ZoneId: 182,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 182,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "HK",
			Name: "Hong Kong",
			Zones: []*Zone{
				{
					Id:          181,
					CountryCode: "HK",
					Name:        "Asia/Hong_Kong",
					TimeZones: []*TimeZone{
						{
							ZoneId: 181,
							Abbr:   "LMT",
							Offset: 27402,
						},
						{
							ZoneId: 181,
							Abbr:   "HKT",
							Offset: 28800,
						},
						{
							ZoneId: 181,
							Abbr:   "HKST",
							Offset: 32400,
						},
						{
							ZoneId: 181,
							Abbr:   "HKWT",
							Offset: 30600,
						},
						{
							ZoneId: 181,
							Abbr:   "JST",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "HU",
			Name: "Hungary",
			Zones: []*Zone{
				{
					Id:          185,
					CountryCode: "HU",
					Name:        "Europe/Budapest",
					TimeZones: []*TimeZone{
						{
							ZoneId: 185,
							Abbr:   "LMT",
							Offset: 4580,
						},
						{
							ZoneId: 185,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 185,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "IS",
			Name: "Iceland",
			Zones: []*Zone{
				{
					Id:          197,
					CountryCode: "IS",
					Name:        "Atlantic/Reykjavik",
					TimeZones: []*TimeZone{
						{
							ZoneId: 197,
							Abbr:   "LMT",
							Offset: -5280,
						},
						{
							ZoneId: 197,
							Abbr:   "IST",
							Offset: -3600,
						},
						{
							ZoneId: 197,
							Abbr:   "ISST",
							Offset: 0,
						},
						{
							ZoneId: 197,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "IN",
			Name: "India",
			Zones: []*Zone{
				{
					Id:          193,
					CountryCode: "IN",
					Name:        "Asia/Kolkata",
					TimeZones: []*TimeZone{
						{
							ZoneId: 193,
							Abbr:   "LMT",
							Offset: 21208,
						},
						{
							ZoneId: 193,
							Abbr:   "HMT",
							Offset: 21200,
						},
						{
							ZoneId: 193,
							Abbr:   "MMT",
							Offset: 19270,
						},
						{
							ZoneId: 193,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 193,
							Abbr:   "IST",
							Offset: 23400,
						},
					},
				},
			},
		},
		{
			Code: "ID",
			Name: "Indonesia",
			Zones: []*Zone{
				{
					Id:          186,
					CountryCode: "ID",
					Name:        "Asia/Jakarta",
					TimeZones: []*TimeZone{
						{
							ZoneId: 186,
							Abbr:   "LMT",
							Offset: 25632,
						},
						{
							ZoneId: 186,
							Abbr:   "BMT",
							Offset: 25632,
						},
						{
							ZoneId: 186,
							Abbr:   "JAVT",
							Offset: 26400,
						},
						{
							ZoneId: 186,
							Abbr:   "WIB",
							Offset: 27000,
						},
						{
							ZoneId: 186,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 186,
							Abbr:   "WIB",
							Offset: 28800,
						},
						{
							ZoneId: 186,
							Abbr:   "WIB",
							Offset: 25200,
						},
					},
				},
				{
					Id:          189,
					CountryCode: "ID",
					Name:        "Asia/Jayapura",
					TimeZones: []*TimeZone{
						{
							ZoneId: 189,
							Abbr:   "LMT",
							Offset: 33768,
						},
						{
							ZoneId: 189,
							Abbr:   "WIT",
							Offset: 32400,
						},
						{
							ZoneId: 189,
							Abbr:   "ACST",
							Offset: 34200,
						},
					},
				},
				{
					Id:          188,
					CountryCode: "ID",
					Name:        "Asia/Makassar",
					TimeZones: []*TimeZone{
						{
							ZoneId: 188,
							Abbr:   "LMT",
							Offset: 28656,
						},
						{
							ZoneId: 188,
							Abbr:   "MMT",
							Offset: 28656,
						},
						{
							ZoneId: 188,
							Abbr:   "WITA",
							Offset: 28800,
						},
						{
							ZoneId: 188,
							Abbr:   "JST",
							Offset: 32400,
						},
					},
				},
				{
					Id:          187,
					CountryCode: "ID",
					Name:        "Asia/Pontianak",
					TimeZones: []*TimeZone{
						{
							ZoneId: 187,
							Abbr:   "LMT",
							Offset: 26240,
						},
						{
							ZoneId: 187,
							Abbr:   "PMT",
							Offset: 26240,
						},
						{
							ZoneId: 187,
							Abbr:   "WIB",
							Offset: 27000,
						},
						{
							ZoneId: 187,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 187,
							Abbr:   "WITA",
							Offset: 28800,
						},
						{
							ZoneId: 187,
							Abbr:   "WIB",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "IR",
			Name: "Iran, Islamic Republic of",
			Zones: []*Zone{
				{
					Id:          196,
					CountryCode: "IR",
					Name:        "Asia/Tehran",
					TimeZones: []*TimeZone{
						{
							ZoneId: 196,
							Abbr:   "LMT",
							Offset: 12344,
						},
						{
							ZoneId: 196,
							Abbr:   "TMT",
							Offset: 12344,
						},
						{
							ZoneId: 196,
							Abbr:   "IRDT",
							Offset: 12600,
						},
						{
							ZoneId: 196,
							Abbr:   "IRDT",
							Offset: 14400,
						},
						{
							ZoneId: 196,
							Abbr:   "IRST",
							Offset: 18000,
						},
						{
							ZoneId: 196,
							Abbr:   "IRST",
							Offset: 16200,
						},
					},
				},
			},
		},
		{
			Code: "IQ",
			Name: "Iraq",
			Zones: []*Zone{
				{
					Id:          195,
					CountryCode: "IQ",
					Name:        "Asia/Baghdad",
					TimeZones: []*TimeZone{
						{
							ZoneId: 195,
							Abbr:   "LMT",
							Offset: 10660,
						},
						{
							ZoneId: 195,
							Abbr:   "BMT",
							Offset: 10656,
						},
						{
							ZoneId: 195,
							Abbr:   "AST",
							Offset: 10800,
						},
						{
							ZoneId: 195,
							Abbr:   "ADT",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "IE",
			Name: "Ireland",
			Zones: []*Zone{
				{
					Id:          190,
					CountryCode: "IE",
					Name:        "Europe/Dublin",
					TimeZones: []*TimeZone{
						{
							ZoneId: 190,
							Abbr:   "LMT",
							Offset: -1500,
						},
						{
							ZoneId: 190,
							Abbr:   "DMT",
							Offset: -1521,
						},
						{
							ZoneId: 190,
							Abbr:   "IST",
							Offset: 2079,
						},
						{
							ZoneId: 190,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 190,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 190,
							Abbr:   "IST",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "IM",
			Name: "Isle of Man",
			Zones: []*Zone{
				{
					Id:          192,
					CountryCode: "IM",
					Name:        "Europe/Isle_of_Man",
					TimeZones: []*TimeZone{
						{
							ZoneId: 192,
							Abbr:   "LMT",
							Offset: -75,
						},
						{
							ZoneId: 192,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 192,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 192,
							Abbr:   "BDST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "IL",
			Name: "Israel",
			Zones: []*Zone{
				{
					Id:          191,
					CountryCode: "IL",
					Name:        "Asia/Jerusalem",
					TimeZones: []*TimeZone{
						{
							ZoneId: 191,
							Abbr:   "LMT",
							Offset: 8454,
						},
						{
							ZoneId: 191,
							Abbr:   "JMT",
							Offset: 8440,
						},
						{
							ZoneId: 191,
							Abbr:   "IST",
							Offset: 7200,
						},
						{
							ZoneId: 191,
							Abbr:   "IDT",
							Offset: 10800,
						},
						{
							ZoneId: 191,
							Abbr:   "IDDT",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "IT",
			Name: "Italy",
			Zones: []*Zone{
				{
					Id:          198,
					CountryCode: "IT",
					Name:        "Europe/Rome",
					TimeZones: []*TimeZone{
						{
							ZoneId: 198,
							Abbr:   "LMT",
							Offset: 2996,
						},
						{
							ZoneId: 198,
							Abbr:   "RMT",
							Offset: 2996,
						},
						{
							ZoneId: 198,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 198,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "JM",
			Name: "Jamaica",
			Zones: []*Zone{
				{
					Id:          200,
					CountryCode: "JM",
					Name:        "America/Jamaica",
					TimeZones: []*TimeZone{
						{
							ZoneId: 200,
							Abbr:   "LMT",
							Offset: -18430,
						},
						{
							ZoneId: 200,
							Abbr:   "KMT",
							Offset: -18430,
						},
						{
							ZoneId: 200,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 200,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "JP",
			Name: "Japan",
			Zones: []*Zone{
				{
					Id:          202,
					CountryCode: "JP",
					Name:        "Asia/Tokyo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 202,
							Abbr:   "LMT",
							Offset: 33539,
						},
						{
							ZoneId: 202,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 202,
							Abbr:   "JDT",
							Offset: 36000,
						},
					},
				},
			},
		},
		{
			Code: "JE",
			Name: "Jersey",
			Zones: []*Zone{
				{
					Id:          199,
					CountryCode: "JE",
					Name:        "Europe/Jersey",
					TimeZones: []*TimeZone{
						{
							ZoneId: 199,
							Abbr:   "LMT",
							Offset: -75,
						},
						{
							ZoneId: 199,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 199,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 199,
							Abbr:   "BDST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "JO",
			Name: "Jordan",
			Zones: []*Zone{
				{
					Id:          201,
					CountryCode: "JO",
					Name:        "Asia/Amman",
					TimeZones: []*TimeZone{
						{
							ZoneId: 201,
							Abbr:   "LMT",
							Offset: 8624,
						},
						{
							ZoneId: 201,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 201,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "KZ",
			Name: "Kazakhstan",
			Zones: []*Zone{
				{
					Id:          215,
					CountryCode: "KZ",
					Name:        "Asia/Almaty",
					TimeZones: []*TimeZone{
						{
							ZoneId: 215,
							Abbr:   "LMT",
							Offset: 18468,
						},
						{
							ZoneId: 215,
							Abbr:   "ALMT",
							Offset: 18000,
						},
						{
							ZoneId: 215,
							Abbr:   "ALMT",
							Offset: 21600,
						},
						{
							ZoneId: 215,
							Abbr:   "ALMST",
							Offset: 25200,
						},
					},
				},
				{
					Id:          219,
					CountryCode: "KZ",
					Name:        "Asia/Aqtau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 219,
							Abbr:   "LMT",
							Offset: 12064,
						},
						{
							ZoneId: 219,
							Abbr:   "FORT",
							Offset: 14400,
						},
						{
							ZoneId: 219,
							Abbr:   "SHET",
							Offset: 18000,
						},
						{
							ZoneId: 219,
							Abbr:   "SHET",
							Offset: 21600,
						},
						{
							ZoneId: 219,
							Abbr:   "SHEST",
							Offset: 21600,
						},
						{
							ZoneId: 219,
							Abbr:   "MSK+1",
							Offset: 14400,
						},
						{
							ZoneId: 219,
							Abbr:   "MSK+1",
							Offset: 18000,
						},
						{
							ZoneId: 219,
							Abbr:   "MSD+1",
							Offset: 21600,
						},
						{
							ZoneId: 219,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
					},
				},
				{
					Id:          218,
					CountryCode: "KZ",
					Name:        "Asia/Aqtobe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 218,
							Abbr:   "LMT",
							Offset: 13720,
						},
						{
							ZoneId: 218,
							Abbr:   "AKTT",
							Offset: 14400,
						},
						{
							ZoneId: 218,
							Abbr:   "AKTT",
							Offset: 18000,
						},
						{
							ZoneId: 218,
							Abbr:   "AKTST",
							Offset: 21600,
						},
						{
							ZoneId: 218,
							Abbr:   "AQTT",
							Offset: 18000,
						},
					},
				},
				{
					Id:          220,
					CountryCode: "KZ",
					Name:        "Asia/Atyrau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 220,
							Abbr:   "LMT",
							Offset: 12464,
						},
						{
							ZoneId: 220,
							Abbr:   "FORT",
							Offset: 10800,
						},
						{
							ZoneId: 220,
							Abbr:   "SHET",
							Offset: 18000,
						},
						{
							ZoneId: 220,
							Abbr:   "SHET",
							Offset: 21600,
						},
						{
							ZoneId: 220,
							Abbr:   "SHEST",
							Offset: 21600,
						},
						{
							ZoneId: 220,
							Abbr:   "MSK+1",
							Offset: 14400,
						},
						{
							ZoneId: 220,
							Abbr:   "MSK+1",
							Offset: 18000,
						},
						{
							ZoneId: 220,
							Abbr:   "MSD+1",
							Offset: 21600,
						},
						{
							ZoneId: 220,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
					},
				},
				{
					Id:          221,
					CountryCode: "KZ",
					Name:        "Asia/Oral",
					TimeZones: []*TimeZone{
						{
							ZoneId: 221,
							Abbr:   "LMT",
							Offset: 12324,
						},
						{
							ZoneId: 221,
							Abbr:   "URAT",
							Offset: 10800,
						},
						{
							ZoneId: 221,
							Abbr:   "URAT",
							Offset: 18000,
						},
						{
							ZoneId: 221,
							Abbr:   "URAST",
							Offset: 21600,
						},
						{
							ZoneId: 221,
							Abbr:   "URAT",
							Offset: 14400,
						},
						{
							ZoneId: 221,
							Abbr:   "URAST",
							Offset: 18000,
						},
						{
							ZoneId: 221,
							Abbr:   "ORAT",
							Offset: 14400,
						},
						{
							ZoneId: 221,
							Abbr:   "ORAST",
							Offset: 18000,
						},
					},
				},
				{
					Id:          217,
					CountryCode: "KZ",
					Name:        "Asia/Qostanay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 217,
							Abbr:   "LMT",
							Offset: 15268,
						},
						{
							ZoneId: 217,
							Abbr:   "QYZT",
							Offset: 14400,
						},
						{
							ZoneId: 217,
							Abbr:   "QYZT",
							Offset: 18000,
						},
						{
							ZoneId: 217,
							Abbr:   "QYZST",
							Offset: 21600,
						},
					},
				},
				{
					Id:          216,
					CountryCode: "KZ",
					Name:        "Asia/Qyzylorda",
					TimeZones: []*TimeZone{
						{
							ZoneId: 216,
							Abbr:   "LMT",
							Offset: 15712,
						},
						{
							ZoneId: 216,
							Abbr:   "KIZT",
							Offset: 14400,
						},
						{
							ZoneId: 216,
							Abbr:   "MSK+2",
							Offset: 18000,
						},
						{
							ZoneId: 216,
							Abbr:   "QYZST",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "KE",
			Name: "Kenya",
			Zones: []*Zone{
				{
					Id:          203,
					CountryCode: "KE",
					Name:        "Africa/Nairobi",
					TimeZones: []*TimeZone{
						{
							ZoneId: 203,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 203,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 203,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 203,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "KI",
			Name: "Kiribati",
			Zones: []*Zone{
				{
					Id:          207,
					CountryCode: "KI",
					Name:        "Pacific/Enderbury",
					TimeZones: []*TimeZone{
						{
							ZoneId: 207,
							Abbr:   "LMT",
							Offset: -41060,
						},
						{
							ZoneId: 207,
							Abbr:   "PHOT",
							Offset: -43200,
						},
						{
							ZoneId: 207,
							Abbr:   "PHOT",
							Offset: -39600,
						},
						{
							ZoneId: 207,
							Abbr:   "PHOT",
							Offset: 46800,
						},
					},
				},
				{
					Id:          208,
					CountryCode: "KI",
					Name:        "Pacific/Kiritimati",
					TimeZones: []*TimeZone{
						{
							ZoneId: 208,
							Abbr:   "LMT",
							Offset: -37760,
						},
						{
							ZoneId: 208,
							Abbr:   "LINT",
							Offset: -38400,
						},
						{
							ZoneId: 208,
							Abbr:   "LINT",
							Offset: -36000,
						},
						{
							ZoneId: 208,
							Abbr:   "LINT",
							Offset: 50400,
						},
					},
				},
				{
					Id:          206,
					CountryCode: "KI",
					Name:        "Pacific/Tarawa",
					TimeZones: []*TimeZone{
						{
							ZoneId: 206,
							Abbr:   "LMT",
							Offset: 41524,
						},
						{
							ZoneId: 206,
							Abbr:   "GILT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "KP",
			Name: "Korea, Democratic People's Republic of",
			Zones: []*Zone{
				{
					Id:          211,
					CountryCode: "KP",
					Name:        "Asia/Pyongyang",
					TimeZones: []*TimeZone{
						{
							ZoneId: 211,
							Abbr:   "LMT",
							Offset: 30180,
						},
						{
							ZoneId: 211,
							Abbr:   "KST",
							Offset: 30600,
						},
						{
							ZoneId: 211,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 211,
							Abbr:   "KST",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "KR",
			Name: "Korea, Republic of",
			Zones: []*Zone{
				{
					Id:          212,
					CountryCode: "KR",
					Name:        "Asia/Seoul",
					TimeZones: []*TimeZone{
						{
							ZoneId: 212,
							Abbr:   "LMT",
							Offset: 30472,
						},
						{
							ZoneId: 212,
							Abbr:   "KST",
							Offset: 30600,
						},
						{
							ZoneId: 212,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 212,
							Abbr:   "KST",
							Offset: 32400,
						},
						{
							ZoneId: 212,
							Abbr:   "KDT",
							Offset: 36000,
						},
						{
							ZoneId: 212,
							Abbr:   "KDT",
							Offset: 34200,
						},
					},
				},
			},
		},
		{
			Code: "KW",
			Name: "Kuwait",
			Zones: []*Zone{
				{
					Id:          213,
					CountryCode: "KW",
					Name:        "Asia/Kuwait",
					TimeZones: []*TimeZone{
						{
							ZoneId: 213,
							Abbr:   "LMT",
							Offset: 11212,
						},
						{
							ZoneId: 213,
							Abbr:   "AST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "KG",
			Name: "Kyrgyzstan",
			Zones: []*Zone{
				{
					Id:          204,
					CountryCode: "KG",
					Name:        "Asia/Bishkek",
					TimeZones: []*TimeZone{
						{
							ZoneId: 204,
							Abbr:   "LMT",
							Offset: 17904,
						},
						{
							ZoneId: 204,
							Abbr:   "FRUT",
							Offset: 18000,
						},
						{
							ZoneId: 204,
							Abbr:   "FRUT",
							Offset: 21600,
						},
						{
							ZoneId: 204,
							Abbr:   "FRUT",
							Offset: 25200,
						},
						{
							ZoneId: 204,
							Abbr:   "FRUST",
							Offset: 25200,
						},
						{
							ZoneId: 204,
							Abbr:   "KGST",
							Offset: 21600,
						},
						{
							ZoneId: 204,
							Abbr:   "KGT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "LA",
			Name: "Lao People's Democratic Republic",
			Zones: []*Zone{
				{
					Id:          222,
					CountryCode: "LA",
					Name:        "Asia/Vientiane",
					TimeZones: []*TimeZone{
						{
							ZoneId: 222,
							Abbr:   "LMT",
							Offset: 24124,
						},
						{
							ZoneId: 222,
							Abbr:   "BMT",
							Offset: 24124,
						},
						{
							ZoneId: 222,
							Abbr:   "ICT",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "LV",
			Name: "Latvia",
			Zones: []*Zone{
				{
					Id:          231,
					CountryCode: "LV",
					Name:        "Europe/Riga",
					TimeZones: []*TimeZone{
						{
							ZoneId: 231,
							Abbr:   "LMT",
							Offset: 5794,
						},
						{
							ZoneId: 231,
							Abbr:   "RMT",
							Offset: 5794,
						},
						{
							ZoneId: 231,
							Abbr:   "LST",
							Offset: 9394,
						},
						{
							ZoneId: 231,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 231,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 231,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 231,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 231,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 231,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "LB",
			Name: "Lebanon",
			Zones: []*Zone{
				{
					Id:          223,
					CountryCode: "LB",
					Name:        "Asia/Beirut",
					TimeZones: []*TimeZone{
						{
							ZoneId: 223,
							Abbr:   "LMT",
							Offset: 8520,
						},
						{
							ZoneId: 223,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 223,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "LS",
			Name: "Lesotho",
			Zones: []*Zone{
				{
					Id:          228,
					CountryCode: "LS",
					Name:        "Africa/Maseru",
					TimeZones: []*TimeZone{
						{
							ZoneId: 228,
							Abbr:   "LMT",
							Offset: 6720,
						},
						{
							ZoneId: 228,
							Abbr:   "SAST",
							Offset: 5400,
						},
						{
							ZoneId: 228,
							Abbr:   "SAST",
							Offset: 7200,
						},
						{
							ZoneId: 228,
							Abbr:   "SAST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "LR",
			Name: "Liberia",
			Zones: []*Zone{
				{
					Id:          227,
					CountryCode: "LR",
					Name:        "Africa/Monrovia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 227,
							Abbr:   "LMT",
							Offset: -2588,
						},
						{
							ZoneId: 227,
							Abbr:   "MMT",
							Offset: -2588,
						},
						{
							ZoneId: 227,
							Abbr:   "MMT",
							Offset: -2670,
						},
						{
							ZoneId: 227,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "LY",
			Name: "Libya",
			Zones: []*Zone{
				{
					Id:          232,
					CountryCode: "LY",
					Name:        "Africa/Tripoli",
					TimeZones: []*TimeZone{
						{
							ZoneId: 232,
							Abbr:   "LMT",
							Offset: 3164,
						},
						{
							ZoneId: 232,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 232,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 232,
							Abbr:   "EET",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "LI",
			Name: "Liechtenstein",
			Zones: []*Zone{
				{
					Id:          225,
					CountryCode: "LI",
					Name:        "Europe/Vaduz",
					TimeZones: []*TimeZone{
						{
							ZoneId: 225,
							Abbr:   "LMT",
							Offset: 2048,
						},
						{
							ZoneId: 225,
							Abbr:   "BMT",
							Offset: 1786,
						},
						{
							ZoneId: 225,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 225,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "LT",
			Name: "Lithuania",
			Zones: []*Zone{
				{
					Id:          229,
					CountryCode: "LT",
					Name:        "Europe/Vilnius",
					TimeZones: []*TimeZone{
						{
							ZoneId: 229,
							Abbr:   "LMT",
							Offset: 6076,
						},
						{
							ZoneId: 229,
							Abbr:   "WMT",
							Offset: 5040,
						},
						{
							ZoneId: 229,
							Abbr:   "KMT",
							Offset: 5736,
						},
						{
							ZoneId: 229,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 229,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 229,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 229,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 229,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 229,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "LU",
			Name: "Luxembourg",
			Zones: []*Zone{
				{
					Id:          230,
					CountryCode: "LU",
					Name:        "Europe/Luxembourg",
					TimeZones: []*TimeZone{
						{
							ZoneId: 230,
							Abbr:   "LMT",
							Offset: 1476,
						},
						{
							ZoneId: 230,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 230,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 230,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 230,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 230,
							Abbr:   "WEST",
							Offset: 7200,
						},
						{
							ZoneId: 230,
							Abbr:   "WET",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "MO",
			Name: "Macao",
			Zones: []*Zone{
				{
					Id:          247,
					CountryCode: "MO",
					Name:        "Asia/Macau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 247,
							Abbr:   "LMT",
							Offset: 27250,
						},
						{
							ZoneId: 247,
							Abbr:   "CST",
							Offset: 28800,
						},
						{
							ZoneId: 247,
							Abbr:   "CST",
							Offset: 32400,
						},
						{
							ZoneId: 247,
							Abbr:   "CDT",
							Offset: 36000,
						},
						{
							ZoneId: 247,
							Abbr:   "CDT",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "MK",
			Name: "Macedonia, the Former Yugoslav Republic of",
			Zones: []*Zone{
				{
					Id:          241,
					CountryCode: "MK",
					Name:        "Europe/Skopje",
					TimeZones: []*TimeZone{
						{
							ZoneId: 241,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 241,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 241,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MG",
			Name: "Madagascar",
			Zones: []*Zone{
				{
					Id:          238,
					CountryCode: "MG",
					Name:        "Indian/Antananarivo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 238,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 238,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 238,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 238,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "MW",
			Name: "Malawi",
			Zones: []*Zone{
				{
					Id:          255,
					CountryCode: "MW",
					Name:        "Africa/Blantyre",
					TimeZones: []*TimeZone{
						{
							ZoneId: 255,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 255,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MY",
			Name: "Malaysia",
			Zones: []*Zone{
				{
					Id:          267,
					CountryCode: "MY",
					Name:        "Asia/Kuala_Lumpur",
					TimeZones: []*TimeZone{
						{
							ZoneId: 267,
							Abbr:   "LMT",
							Offset: 24406,
						},
						{
							ZoneId: 267,
							Abbr:   "SMT",
							Offset: 24925,
						},
						{
							ZoneId: 267,
							Abbr:   "MALT",
							Offset: 25200,
						},
						{
							ZoneId: 267,
							Abbr:   "MALST",
							Offset: 26400,
						},
						{
							ZoneId: 267,
							Abbr:   "MALT",
							Offset: 27000,
						},
						{
							ZoneId: 267,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 267,
							Abbr:   "MYT",
							Offset: 28800,
						},
					},
				},
				{
					Id:          268,
					CountryCode: "MY",
					Name:        "Asia/Kuching",
					TimeZones: []*TimeZone{
						{
							ZoneId: 268,
							Abbr:   "LMT",
							Offset: 26480,
						},
						{
							ZoneId: 268,
							Abbr:   "BORT",
							Offset: 27000,
						},
						{
							ZoneId: 268,
							Abbr:   "BORT",
							Offset: 28800,
						},
						{
							ZoneId: 268,
							Abbr:   "BORTST",
							Offset: 30000,
						},
						{
							ZoneId: 268,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 268,
							Abbr:   "MYT",
							Offset: 28800,
						},
					},
				},
			},
		},
		{
			Code: "MV",
			Name: "Maldives",
			Zones: []*Zone{
				{
					Id:          254,
					CountryCode: "MV",
					Name:        "Indian/Maldives",
					TimeZones: []*TimeZone{
						{
							ZoneId: 254,
							Abbr:   "LMT",
							Offset: 17640,
						},
						{
							ZoneId: 254,
							Abbr:   "MMT",
							Offset: 17640,
						},
						{
							ZoneId: 254,
							Abbr:   "MVT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "ML",
			Name: "Mali",
			Zones: []*Zone{
				{
					Id:          242,
					CountryCode: "ML",
					Name:        "Africa/Bamako",
					TimeZones: []*TimeZone{
						{
							ZoneId: 242,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 242,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "MT",
			Name: "Malta",
			Zones: []*Zone{
				{
					Id:          252,
					CountryCode: "MT",
					Name:        "Europe/Malta",
					TimeZones: []*TimeZone{
						{
							ZoneId: 252,
							Abbr:   "LMT",
							Offset: 3484,
						},
						{
							ZoneId: 252,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 252,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MH",
			Name: "Marshall Islands",
			Zones: []*Zone{
				{
					Id:          240,
					CountryCode: "MH",
					Name:        "Pacific/Kwajalein",
					TimeZones: []*TimeZone{
						{
							ZoneId: 240,
							Abbr:   "LMT",
							Offset: 40160,
						},
						{
							ZoneId: 240,
							Abbr:   "+11",
							Offset: 39600,
						},
						{
							ZoneId: 240,
							Abbr:   "+10",
							Offset: 36000,
						},
						{
							ZoneId: 240,
							Abbr:   "+09",
							Offset: 32400,
						},
						{
							ZoneId: 240,
							Abbr:   "--12",
							Offset: -43200,
						},
						{
							ZoneId: 240,
							Abbr:   "+12",
							Offset: 43200,
						},
					},
				},
				{
					Id:          239,
					CountryCode: "MH",
					Name:        "Pacific/Majuro",
					TimeZones: []*TimeZone{
						{
							ZoneId: 239,
							Abbr:   "LMT",
							Offset: 41088,
						},
						{
							ZoneId: 239,
							Abbr:   "+11",
							Offset: 39600,
						},
						{
							ZoneId: 239,
							Abbr:   "+09",
							Offset: 32400,
						},
						{
							ZoneId: 239,
							Abbr:   "+10",
							Offset: 36000,
						},
						{
							ZoneId: 239,
							Abbr:   "+12",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "MQ",
			Name: "Martinique",
			Zones: []*Zone{
				{
					Id:          249,
					CountryCode: "MQ",
					Name:        "America/Martinique",
					TimeZones: []*TimeZone{
						{
							ZoneId: 249,
							Abbr:   "LMT",
							Offset: -14660,
						},
						{
							ZoneId: 249,
							Abbr:   "FFMT",
							Offset: -14660,
						},
						{
							ZoneId: 249,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 249,
							Abbr:   "ADT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "MR",
			Name: "Mauritania",
			Zones: []*Zone{
				{
					Id:          250,
					CountryCode: "MR",
					Name:        "Africa/Nouakchott",
					TimeZones: []*TimeZone{
						{
							ZoneId: 250,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 250,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "MU",
			Name: "Mauritius",
			Zones: []*Zone{
				{
					Id:          253,
					CountryCode: "MU",
					Name:        "Indian/Mauritius",
					TimeZones: []*TimeZone{
						{
							ZoneId: 253,
							Abbr:   "LMT",
							Offset: 13800,
						},
						{
							ZoneId: 253,
							Abbr:   "MUT",
							Offset: 14400,
						},
						{
							ZoneId: 253,
							Abbr:   "MUST",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "YT",
			Name: "Mayotte",
			Zones: []*Zone{
				{
					Id:          421,
					CountryCode: "YT",
					Name:        "Indian/Mayotte",
					TimeZones: []*TimeZone{
						{
							ZoneId: 421,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 421,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 421,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 421,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "MX",
			Name: "Mexico",
			Zones: []*Zone{
				{
					Id:          266,
					CountryCode: "MX",
					Name:        "America/Bahia_Banderas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 266,
							Abbr:   "LMT",
							Offset: -25260,
						},
						{
							ZoneId: 266,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 266,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 266,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 266,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 266,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          257,
					CountryCode: "MX",
					Name:        "America/Cancun",
					TimeZones: []*TimeZone{
						{
							ZoneId: 257,
							Abbr:   "LMT",
							Offset: -20824,
						},
						{
							ZoneId: 257,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 257,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 257,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 257,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          262,
					CountryCode: "MX",
					Name:        "America/Chihuahua",
					TimeZones: []*TimeZone{
						{
							ZoneId: 262,
							Abbr:   "LMT",
							Offset: -25460,
						},
						{
							ZoneId: 262,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 262,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 262,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 262,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          264,
					CountryCode: "MX",
					Name:        "America/Hermosillo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 264,
							Abbr:   "LMT",
							Offset: -26632,
						},
						{
							ZoneId: 264,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 264,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 264,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 264,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          260,
					CountryCode: "MX",
					Name:        "America/Matamoros",
					TimeZones: []*TimeZone{
						{
							ZoneId: 260,
							Abbr:   "LMT",
							Offset: -24000,
						},
						{
							ZoneId: 260,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 260,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          261,
					CountryCode: "MX",
					Name:        "America/Mazatlan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 261,
							Abbr:   "LMT",
							Offset: -25540,
						},
						{
							ZoneId: 261,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 261,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 261,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 261,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          258,
					CountryCode: "MX",
					Name:        "America/Merida",
					TimeZones: []*TimeZone{
						{
							ZoneId: 258,
							Abbr:   "LMT",
							Offset: -21508,
						},
						{
							ZoneId: 258,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 258,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 258,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          256,
					CountryCode: "MX",
					Name:        "America/Mexico_City",
					TimeZones: []*TimeZone{
						{
							ZoneId: 256,
							Abbr:   "LMT",
							Offset: -23796,
						},
						{
							ZoneId: 256,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 256,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 256,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 256,
							Abbr:   "CWT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          259,
					CountryCode: "MX",
					Name:        "America/Monterrey",
					TimeZones: []*TimeZone{
						{
							ZoneId: 259,
							Abbr:   "LMT",
							Offset: -24076,
						},
						{
							ZoneId: 259,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 259,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          263,
					CountryCode: "MX",
					Name:        "America/Ojinaga",
					TimeZones: []*TimeZone{
						{
							ZoneId: 263,
							Abbr:   "LMT",
							Offset: -25060,
						},
						{
							ZoneId: 263,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 263,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 263,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 263,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          265,
					CountryCode: "MX",
					Name:        "America/Tijuana",
					TimeZones: []*TimeZone{
						{
							ZoneId: 265,
							Abbr:   "LMT",
							Offset: -28084,
						},
						{
							ZoneId: 265,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 265,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 265,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 265,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 265,
							Abbr:   "PPT",
							Offset: -25200,
						},
					},
				},
			},
		},
		{
			Code: "FM",
			Name: "Micronesia, Federated States of",
			Zones: []*Zone{
				{
					Id:          154,
					CountryCode: "FM",
					Name:        "Pacific/Chuuk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 154,
							Abbr:   "LMT",
							Offset: -49972,
						},
						{
							ZoneId: 154,
							Abbr:   "LMT",
							Offset: 36428,
						},
						{
							ZoneId: 154,
							Abbr:   "CHUT",
							Offset: 36000,
						},
						{
							ZoneId: 154,
							Abbr:   "CHUT",
							Offset: 32400,
						},
					},
				},
				{
					Id:          156,
					CountryCode: "FM",
					Name:        "Pacific/Kosrae",
					TimeZones: []*TimeZone{
						{
							ZoneId: 156,
							Abbr:   "LMT",
							Offset: -47284,
						},
						{
							ZoneId: 156,
							Abbr:   "LMT",
							Offset: 39116,
						},
						{
							ZoneId: 156,
							Abbr:   "KOST",
							Offset: 39600,
						},
						{
							ZoneId: 156,
							Abbr:   "KOST",
							Offset: 32400,
						},
						{
							ZoneId: 156,
							Abbr:   "KOST",
							Offset: 36000,
						},
						{
							ZoneId: 156,
							Abbr:   "KOST",
							Offset: 43200,
						},
					},
				},
				{
					Id:          155,
					CountryCode: "FM",
					Name:        "Pacific/Pohnpei",
					TimeZones: []*TimeZone{
						{
							ZoneId: 155,
							Abbr:   "LMT",
							Offset: -48428,
						},
						{
							ZoneId: 155,
							Abbr:   "LMT",
							Offset: 37972,
						},
						{
							ZoneId: 155,
							Abbr:   "PONT",
							Offset: 39600,
						},
						{
							ZoneId: 155,
							Abbr:   "PONT",
							Offset: 32400,
						},
						{
							ZoneId: 155,
							Abbr:   "PONT",
							Offset: 36000,
						},
					},
				},
			},
		},
		{
			Code: "MD",
			Name: "Moldova, Republic of",
			Zones: []*Zone{
				{
					Id:          235,
					CountryCode: "MD",
					Name:        "Europe/Chisinau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 235,
							Abbr:   "LMT",
							Offset: 6920,
						},
						{
							ZoneId: 235,
							Abbr:   "CMT",
							Offset: 6900,
						},
						{
							ZoneId: 235,
							Abbr:   "BMT",
							Offset: 6264,
						},
						{
							ZoneId: 235,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 235,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 235,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 235,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 235,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 235,
							Abbr:   "MSD",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "MC",
			Name: "Monaco",
			Zones: []*Zone{
				{
					Id:          234,
					CountryCode: "MC",
					Name:        "Europe/Monaco",
					TimeZones: []*TimeZone{
						{
							ZoneId: 234,
							Abbr:   "LMT",
							Offset: 1772,
						},
						{
							ZoneId: 234,
							Abbr:   "PMT",
							Offset: 561,
						},
						{
							ZoneId: 234,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 234,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 234,
							Abbr:   "WEMT",
							Offset: 7200,
						},
						{
							ZoneId: 234,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 234,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MN",
			Name: "Mongolia",
			Zones: []*Zone{
				{
					Id:          246,
					CountryCode: "MN",
					Name:        "Asia/Choibalsan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 246,
							Abbr:   "LMT",
							Offset: 27480,
						},
						{
							ZoneId: 246,
							Abbr:   "ULAT",
							Offset: 25200,
						},
						{
							ZoneId: 246,
							Abbr:   "ULAT",
							Offset: 28800,
						},
						{
							ZoneId: 246,
							Abbr:   "CHOST",
							Offset: 36000,
						},
						{
							ZoneId: 246,
							Abbr:   "CHOT",
							Offset: 32400,
						},
						{
							ZoneId: 246,
							Abbr:   "CHOT",
							Offset: 28800,
						},
						{
							ZoneId: 246,
							Abbr:   "CHOST",
							Offset: 32400,
						},
					},
				},
				{
					Id:          245,
					CountryCode: "MN",
					Name:        "Asia/Hovd",
					TimeZones: []*TimeZone{
						{
							ZoneId: 245,
							Abbr:   "LMT",
							Offset: 21996,
						},
						{
							ZoneId: 245,
							Abbr:   "HOVT",
							Offset: 21600,
						},
						{
							ZoneId: 245,
							Abbr:   "HOVT",
							Offset: 25200,
						},
						{
							ZoneId: 245,
							Abbr:   "HOVST",
							Offset: 28800,
						},
					},
				},
				{
					Id:          244,
					CountryCode: "MN",
					Name:        "Asia/Ulaanbaatar",
					TimeZones: []*TimeZone{
						{
							ZoneId: 244,
							Abbr:   "LMT",
							Offset: 25652,
						},
						{
							ZoneId: 244,
							Abbr:   "ULAT",
							Offset: 25200,
						},
						{
							ZoneId: 244,
							Abbr:   "ULAT",
							Offset: 28800,
						},
						{
							ZoneId: 244,
							Abbr:   "ULAST",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "ME",
			Name: "Montenegro",
			Zones: []*Zone{
				{
					Id:          236,
					CountryCode: "ME",
					Name:        "Europe/Podgorica",
					TimeZones: []*TimeZone{
						{
							ZoneId: 236,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 236,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 236,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MS",
			Name: "Montserrat",
			Zones: []*Zone{
				{
					Id:          251,
					CountryCode: "MS",
					Name:        "America/Montserrat",
					TimeZones: []*TimeZone{
						{
							ZoneId: 251,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 251,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "MA",
			Name: "Morocco",
			Zones: []*Zone{
				{
					Id:          233,
					CountryCode: "MA",
					Name:        "Africa/Casablanca",
					TimeZones: []*TimeZone{
						{
							ZoneId: 233,
							Abbr:   "LMT",
							Offset: -1820,
						},
						{
							ZoneId: 233,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 233,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "MZ",
			Name: "Mozambique",
			Zones: []*Zone{
				{
					Id:          269,
					CountryCode: "MZ",
					Name:        "Africa/Maputo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 269,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 269,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "MM",
			Name: "Myanmar",
			Zones: []*Zone{
				{
					Id:          243,
					CountryCode: "MM",
					Name:        "Asia/Yangon",
					TimeZones: []*TimeZone{
						{
							ZoneId: 243,
							Abbr:   "LMT",
							Offset: 23087,
						},
						{
							ZoneId: 243,
							Abbr:   "RMT",
							Offset: 23087,
						},
						{
							ZoneId: 243,
							Abbr:   "BURT",
							Offset: 23400,
						},
						{
							ZoneId: 243,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 243,
							Abbr:   "MMT",
							Offset: 23400,
						},
					},
				},
			},
		},
		{
			Code: "NA",
			Name: "Namibia",
			Zones: []*Zone{
				{
					Id:          270,
					CountryCode: "NA",
					Name:        "Africa/Windhoek",
					TimeZones: []*TimeZone{
						{
							ZoneId: 270,
							Abbr:   "LMT",
							Offset: 4104,
						},
						{
							ZoneId: 270,
							Abbr:   "SAST",
							Offset: 5400,
						},
						{
							ZoneId: 270,
							Abbr:   "SAST",
							Offset: 7200,
						},
						{
							ZoneId: 270,
							Abbr:   "SAST",
							Offset: 10800,
						},
						{
							ZoneId: 270,
							Abbr:   "CAT",
							Offset: 7200,
						},
						{
							ZoneId: 270,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "NR",
			Name: "Nauru",
			Zones: []*Zone{
				{
					Id:          279,
					CountryCode: "NR",
					Name:        "Pacific/Nauru",
					TimeZones: []*TimeZone{
						{
							ZoneId: 279,
							Abbr:   "LMT",
							Offset: 40060,
						},
						{
							ZoneId: 279,
							Abbr:   "NRT",
							Offset: 41400,
						},
						{
							ZoneId: 279,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 279,
							Abbr:   "NRT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "NP",
			Name: "Nepal",
			Zones: []*Zone{
				{
					Id:          278,
					CountryCode: "NP",
					Name:        "Asia/Kathmandu",
					TimeZones: []*TimeZone{
						{
							ZoneId: 278,
							Abbr:   "LMT",
							Offset: 20476,
						},
						{
							ZoneId: 278,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 278,
							Abbr:   "NPT",
							Offset: 20700,
						},
					},
				},
			},
		},
		{
			Code: "NL",
			Name: "Netherlands",
			Zones: []*Zone{
				{
					Id:          276,
					CountryCode: "NL",
					Name:        "Europe/Amsterdam",
					TimeZones: []*TimeZone{
						{
							ZoneId: 276,
							Abbr:   "LMT",
							Offset: 1172,
						},
						{
							ZoneId: 276,
							Abbr:   "AMT",
							Offset: 1172,
						},
						{
							ZoneId: 276,
							Abbr:   "NST",
							Offset: 4772,
						},
						{
							ZoneId: 276,
							Abbr:   "NST",
							Offset: 4800,
						},
						{
							ZoneId: 276,
							Abbr:   "AMT",
							Offset: 1200,
						},
						{
							ZoneId: 276,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 276,
							Abbr:   "CET",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "NC",
			Name: "New Caledonia",
			Zones: []*Zone{
				{
					Id:          271,
					CountryCode: "NC",
					Name:        "Pacific/Noumea",
					TimeZones: []*TimeZone{
						{
							ZoneId: 271,
							Abbr:   "LMT",
							Offset: 39948,
						},
						{
							ZoneId: 271,
							Abbr:   "NCT",
							Offset: 39600,
						},
						{
							ZoneId: 271,
							Abbr:   "NCST",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "NZ",
			Name: "New Zealand",
			Zones: []*Zone{
				{
					Id:          281,
					CountryCode: "NZ",
					Name:        "Pacific/Auckland",
					TimeZones: []*TimeZone{
						{
							ZoneId: 281,
							Abbr:   "LMT",
							Offset: 41944,
						},
						{
							ZoneId: 281,
							Abbr:   "NZMT",
							Offset: 41400,
						},
						{
							ZoneId: 281,
							Abbr:   "NZST",
							Offset: 45000,
						},
						{
							ZoneId: 281,
							Abbr:   "NZST",
							Offset: 43200,
						},
						{
							ZoneId: 281,
							Abbr:   "NZDT",
							Offset: 46800,
						},
					},
				},
				{
					Id:          282,
					CountryCode: "NZ",
					Name:        "Pacific/Chatham",
					TimeZones: []*TimeZone{
						{
							ZoneId: 282,
							Abbr:   "LMT",
							Offset: 44028,
						},
						{
							ZoneId: 282,
							Abbr:   "CHADT",
							Offset: 44100,
						},
						{
							ZoneId: 282,
							Abbr:   "CHADT",
							Offset: 45900,
						},
						{
							ZoneId: 282,
							Abbr:   "CHAST",
							Offset: 49500,
						},
					},
				},
			},
		},
		{
			Code: "NI",
			Name: "Nicaragua",
			Zones: []*Zone{
				{
					Id:          275,
					CountryCode: "NI",
					Name:        "America/Managua",
					TimeZones: []*TimeZone{
						{
							ZoneId: 275,
							Abbr:   "LMT",
							Offset: -20708,
						},
						{
							ZoneId: 275,
							Abbr:   "MMT",
							Offset: -20712,
						},
						{
							ZoneId: 275,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 275,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 275,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "NE",
			Name: "Niger",
			Zones: []*Zone{
				{
					Id:          272,
					CountryCode: "NE",
					Name:        "Africa/Niamey",
					TimeZones: []*TimeZone{
						{
							ZoneId: 272,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 272,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 272,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 272,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "NG",
			Name: "Nigeria",
			Zones: []*Zone{
				{
					Id:          274,
					CountryCode: "NG",
					Name:        "Africa/Lagos",
					TimeZones: []*TimeZone{
						{
							ZoneId: 274,
							Abbr:   "LMT",
							Offset: 815,
						},
						{
							ZoneId: 274,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 274,
							Abbr:   "WAT",
							Offset: 1800,
						},
						{
							ZoneId: 274,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "NU",
			Name: "Niue",
			Zones: []*Zone{
				{
					Id:          280,
					CountryCode: "NU",
					Name:        "Pacific/Niue",
					TimeZones: []*TimeZone{
						{
							ZoneId: 280,
							Abbr:   "LMT",
							Offset: -40780,
						},
						{
							ZoneId: 280,
							Abbr:   "NUT",
							Offset: -40800,
						},
						{
							ZoneId: 280,
							Abbr:   "NUT",
							Offset: -41400,
						},
						{
							ZoneId: 280,
							Abbr:   "NUT",
							Offset: -39600,
						},
					},
				},
			},
		},
		{
			Code: "NF",
			Name: "Norfolk Island",
			Zones: []*Zone{
				{
					Id:          273,
					CountryCode: "NF",
					Name:        "Pacific/Norfolk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 273,
							Abbr:   "LMT",
							Offset: 40312,
						},
						{
							ZoneId: 273,
							Abbr:   "NMT",
							Offset: 40320,
						},
						{
							ZoneId: 273,
							Abbr:   "NFT",
							Offset: 41400,
						},
						{
							ZoneId: 273,
							Abbr:   "NFST",
							Offset: 45000,
						},
						{
							ZoneId: 273,
							Abbr:   "NFT",
							Offset: 39600,
						},
						{
							ZoneId: 273,
							Abbr:   "NFDT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "MP",
			Name: "Northern Mariana Islands",
			Zones: []*Zone{
				{
					Id:          248,
					CountryCode: "MP",
					Name:        "Pacific/Saipan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 248,
							Abbr:   "LMT",
							Offset: -51660,
						},
						{
							ZoneId: 248,
							Abbr:   "LMT",
							Offset: 34740,
						},
						{
							ZoneId: 248,
							Abbr:   "GST",
							Offset: 36000,
						},
						{
							ZoneId: 248,
							Abbr:   "GST",
							Offset: 32400,
						},
						{
							ZoneId: 248,
							Abbr:   "GDT",
							Offset: 39600,
						},
						{
							ZoneId: 248,
							Abbr:   "ChST",
							Offset: 36000,
						},
					},
				},
			},
		},
		{
			Code: "NO",
			Name: "Norway",
			Zones: []*Zone{
				{
					Id:          277,
					CountryCode: "NO",
					Name:        "Europe/Oslo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 277,
							Abbr:   "LMT",
							Offset: 2580,
						},
						{
							ZoneId: 277,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 277,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "OM",
			Name: "Oman",
			Zones: []*Zone{
				{
					Id:          283,
					CountryCode: "OM",
					Name:        "Asia/Muscat",
					TimeZones: []*TimeZone{
						{
							ZoneId: 283,
							Abbr:   "LMT",
							Offset: 13272,
						},
						{
							ZoneId: 283,
							Abbr:   "GST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "PK",
			Name: "Pakistan",
			Zones: []*Zone{
				{
					Id:          292,
					CountryCode: "PK",
					Name:        "Asia/Karachi",
					TimeZones: []*TimeZone{
						{
							ZoneId: 292,
							Abbr:   "LMT",
							Offset: 16092,
						},
						{
							ZoneId: 292,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 292,
							Abbr:   "IST",
							Offset: 23400,
						},
						{
							ZoneId: 292,
							Abbr:   "KART",
							Offset: 18000,
						},
						{
							ZoneId: 292,
							Abbr:   "PKT",
							Offset: 18000,
						},
						{
							ZoneId: 292,
							Abbr:   "PKST",
							Offset: 21600,
						},
					},
				},
			},
		},
		{
			Code: "PW",
			Name: "Palau",
			Zones: []*Zone{
				{
					Id:          302,
					CountryCode: "PW",
					Name:        "Pacific/Palau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 302,
							Abbr:   "LMT",
							Offset: -54124,
						},
						{
							ZoneId: 302,
							Abbr:   "LMT",
							Offset: 32276,
						},
						{
							ZoneId: 302,
							Abbr:   "+09",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "PS",
			Name: "Palestine, State of",
			Zones: []*Zone{
				{
					Id:          297,
					CountryCode: "PS",
					Name:        "Asia/Gaza",
					TimeZones: []*TimeZone{
						{
							ZoneId: 297,
							Abbr:   "LMT",
							Offset: 8272,
						},
						{
							ZoneId: 297,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 297,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 297,
							Abbr:   "IST",
							Offset: 7200,
						},
						{
							ZoneId: 297,
							Abbr:   "IDT",
							Offset: 10800,
						},
					},
				},
				{
					Id:          298,
					CountryCode: "PS",
					Name:        "Asia/Hebron",
					TimeZones: []*TimeZone{
						{
							ZoneId: 298,
							Abbr:   "LMT",
							Offset: 8423,
						},
						{
							ZoneId: 298,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 298,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 298,
							Abbr:   "IST",
							Offset: 7200,
						},
						{
							ZoneId: 298,
							Abbr:   "IDT",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "PA",
			Name: "Panama",
			Zones: []*Zone{
				{
					Id:          284,
					CountryCode: "PA",
					Name:        "America/Panama",
					TimeZones: []*TimeZone{
						{
							ZoneId: 284,
							Abbr:   "LMT",
							Offset: -19088,
						},
						{
							ZoneId: 284,
							Abbr:   "CMT",
							Offset: -19176,
						},
						{
							ZoneId: 284,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
			},
		},
		{
			Code: "PG",
			Name: "Papua New Guinea",
			Zones: []*Zone{
				{
					Id:          290,
					CountryCode: "PG",
					Name:        "Pacific/Bougainville",
					TimeZones: []*TimeZone{
						{
							ZoneId: 290,
							Abbr:   "LMT",
							Offset: 37336,
						},
						{
							ZoneId: 290,
							Abbr:   "PMMT",
							Offset: 35312,
						},
						{
							ZoneId: 290,
							Abbr:   "PGT",
							Offset: 36000,
						},
						{
							ZoneId: 290,
							Abbr:   "PGT",
							Offset: 32400,
						},
						{
							ZoneId: 290,
							Abbr:   "BST",
							Offset: 39600,
						},
					},
				},
				{
					Id:          289,
					CountryCode: "PG",
					Name:        "Pacific/Port_Moresby",
					TimeZones: []*TimeZone{
						{
							ZoneId: 289,
							Abbr:   "LMT",
							Offset: 35320,
						},
						{
							ZoneId: 289,
							Abbr:   "PMMT",
							Offset: 35312,
						},
						{
							ZoneId: 289,
							Abbr:   "PGT",
							Offset: 36000,
						},
					},
				},
			},
		},
		{
			Code: "PY",
			Name: "Paraguay",
			Zones: []*Zone{
				{
					Id:          303,
					CountryCode: "PY",
					Name:        "America/Asuncion",
					TimeZones: []*TimeZone{
						{
							ZoneId: 303,
							Abbr:   "LMT",
							Offset: -13840,
						},
						{
							ZoneId: 303,
							Abbr:   "AMT",
							Offset: -13840,
						},
						{
							ZoneId: 303,
							Abbr:   "PYT",
							Offset: -14400,
						},
						{
							ZoneId: 303,
							Abbr:   "PYT",
							Offset: -10800,
						},
						{
							ZoneId: 303,
							Abbr:   "PYST",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "PE",
			Name: "Peru",
			Zones: []*Zone{
				{
					Id:          285,
					CountryCode: "PE",
					Name:        "America/Lima",
					TimeZones: []*TimeZone{
						{
							ZoneId: 285,
							Abbr:   "LMT",
							Offset: -18492,
						},
						{
							ZoneId: 285,
							Abbr:   "LMT",
							Offset: -18516,
						},
						{
							ZoneId: 285,
							Abbr:   "PET",
							Offset: -18000,
						},
						{
							ZoneId: 285,
							Abbr:   "PEST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "PH",
			Name: "Philippines",
			Zones: []*Zone{
				{
					Id:          291,
					CountryCode: "PH",
					Name:        "Asia/Manila",
					TimeZones: []*TimeZone{
						{
							ZoneId: 291,
							Abbr:   "LMT",
							Offset: -57360,
						},
						{
							ZoneId: 291,
							Abbr:   "LMT",
							Offset: 29040,
						},
						{
							ZoneId: 291,
							Abbr:   "PST",
							Offset: 28800,
						},
						{
							ZoneId: 291,
							Abbr:   "PDT",
							Offset: 32400,
						},
						{
							ZoneId: 291,
							Abbr:   "JST",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "PN",
			Name: "Pitcairn",
			Zones: []*Zone{
				{
					Id:          295,
					CountryCode: "PN",
					Name:        "Pacific/Pitcairn",
					TimeZones: []*TimeZone{
						{
							ZoneId: 295,
							Abbr:   "LMT",
							Offset: -31220,
						},
						{
							ZoneId: 295,
							Abbr:   "PNT",
							Offset: -30600,
						},
						{
							ZoneId: 295,
							Abbr:   "PST",
							Offset: -28800,
						},
					},
				},
			},
		},
		{
			Code: "PL",
			Name: "Poland",
			Zones: []*Zone{
				{
					Id:          293,
					CountryCode: "PL",
					Name:        "Europe/Warsaw",
					TimeZones: []*TimeZone{
						{
							ZoneId: 293,
							Abbr:   "LMT",
							Offset: 5040,
						},
						{
							ZoneId: 293,
							Abbr:   "WMT",
							Offset: 5040,
						},
						{
							ZoneId: 293,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 293,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 293,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 293,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "PT",
			Name: "Portugal",
			Zones: []*Zone{
				{
					Id:          301,
					CountryCode: "PT",
					Name:        "Atlantic/Azores",
					TimeZones: []*TimeZone{
						{
							ZoneId: 301,
							Abbr:   "LMT",
							Offset: -6160,
						},
						{
							ZoneId: 301,
							Abbr:   "HMT",
							Offset: -6872,
						},
						{
							ZoneId: 301,
							Abbr:   "AZOT",
							Offset: -7200,
						},
						{
							ZoneId: 301,
							Abbr:   "AZOST",
							Offset: -3600,
						},
						{
							ZoneId: 301,
							Abbr:   "AZOST",
							Offset: 0,
						},
						{
							ZoneId: 301,
							Abbr:   "AZOT",
							Offset: -3600,
						},
						{
							ZoneId: 301,
							Abbr:   "WET",
							Offset: 0,
						},
					},
				},
				{
					Id:          300,
					CountryCode: "PT",
					Name:        "Atlantic/Madeira",
					TimeZones: []*TimeZone{
						{
							ZoneId: 300,
							Abbr:   "LMT",
							Offset: -4056,
						},
						{
							ZoneId: 300,
							Abbr:   "FMT",
							Offset: -4056,
						},
						{
							ZoneId: 300,
							Abbr:   "WAT",
							Offset: -3600,
						},
						{
							ZoneId: 300,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 300,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
				{
					Id:          299,
					CountryCode: "PT",
					Name:        "Europe/Lisbon",
					TimeZones: []*TimeZone{
						{
							ZoneId: 299,
							Abbr:   "LMT",
							Offset: -2205,
						},
						{
							ZoneId: 299,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 299,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 299,
							Abbr:   "WEMT",
							Offset: 7200,
						},
						{
							ZoneId: 299,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 299,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "PR",
			Name: "Puerto Rico",
			Zones: []*Zone{
				{
					Id:          296,
					CountryCode: "PR",
					Name:        "America/Puerto_Rico",
					TimeZones: []*TimeZone{
						{
							ZoneId: 296,
							Abbr:   "LMT",
							Offset: -15865,
						},
						{
							ZoneId: 296,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 296,
							Abbr:   "AWT",
							Offset: -10800,
						},
						{
							ZoneId: 296,
							Abbr:   "APT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "QA",
			Name: "Qatar",
			Zones: []*Zone{
				{
					Id:          304,
					CountryCode: "QA",
					Name:        "Asia/Qatar",
					TimeZones: []*TimeZone{
						{
							ZoneId: 304,
							Abbr:   "LMT",
							Offset: 12368,
						},
						{
							ZoneId: 304,
							Abbr:   "GST",
							Offset: 14400,
						},
						{
							ZoneId: 304,
							Abbr:   "AST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "RO",
			Name: "Romania",
			Zones: []*Zone{
				{
					Id:          306,
					CountryCode: "RO",
					Name:        "Europe/Bucharest",
					TimeZones: []*TimeZone{
						{
							ZoneId: 306,
							Abbr:   "LMT",
							Offset: 6264,
						},
						{
							ZoneId: 306,
							Abbr:   "BMT",
							Offset: 6264,
						},
						{
							ZoneId: 306,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 306,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "RU",
			Name: "Russian Federation",
			Zones: []*Zone{
				{
					Id:          334,
					CountryCode: "RU",
					Name:        "Asia/Anadyr",
					TimeZones: []*TimeZone{
						{
							ZoneId: 334,
							Abbr:   "LMT",
							Offset: 42596,
						},
						{
							ZoneId: 334,
							Abbr:   "ANAT",
							Offset: 43200,
						},
						{
							ZoneId: 334,
							Abbr:   "ANAT",
							Offset: 46800,
						},
						{
							ZoneId: 334,
							Abbr:   "ANAST",
							Offset: 50400,
						},
						{
							ZoneId: 334,
							Abbr:   "ANAST",
							Offset: 46800,
						},
						{
							ZoneId: 334,
							Abbr:   "ANAT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          320,
					CountryCode: "RU",
					Name:        "Asia/Barnaul",
					TimeZones: []*TimeZone{
						{
							ZoneId: 320,
							Abbr:   "LMT",
							Offset: 20100,
						},
						{
							ZoneId: 320,
							Abbr:   "MSK+3",
							Offset: 21600,
						},
						{
							ZoneId: 320,
							Abbr:   "MSK+4",
							Offset: 25200,
						},
						{
							ZoneId: 320,
							Abbr:   "MSD+4",
							Offset: 28800,
						},
						{
							ZoneId: 320,
							Abbr:   "MSK+4",
							Offset: 21600,
						},
						{
							ZoneId: 320,
							Abbr:   "MSD+3",
							Offset: 25200,
						},
						{
							ZoneId: 320,
							Abbr:   "MSK+3",
							Offset: 25200,
						},
					},
				},
				{
					Id:          325,
					CountryCode: "RU",
					Name:        "Asia/Chita",
					TimeZones: []*TimeZone{
						{
							ZoneId: 325,
							Abbr:   "LMT",
							Offset: 27232,
						},
						{
							ZoneId: 325,
							Abbr:   "YAKT",
							Offset: 28800,
						},
						{
							ZoneId: 325,
							Abbr:   "YAKT",
							Offset: 32400,
						},
						{
							ZoneId: 325,
							Abbr:   "YAKST",
							Offset: 36000,
						},
						{
							ZoneId: 325,
							Abbr:   "YAKT",
							Offset: 36000,
						},
						{
							ZoneId: 325,
							Abbr:   "IRKT",
							Offset: 28800,
						},
					},
				},
				{
					Id:          324,
					CountryCode: "RU",
					Name:        "Asia/Irkutsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 324,
							Abbr:   "LMT",
							Offset: 25025,
						},
						{
							ZoneId: 324,
							Abbr:   "IMT",
							Offset: 25025,
						},
						{
							ZoneId: 324,
							Abbr:   "IRKT",
							Offset: 25200,
						},
						{
							ZoneId: 324,
							Abbr:   "IRKT",
							Offset: 28800,
						},
						{
							ZoneId: 324,
							Abbr:   "IRKST",
							Offset: 32400,
						},
						{
							ZoneId: 324,
							Abbr:   "IRKT",
							Offset: 32400,
						},
					},
				},
				{
					Id:          333,
					CountryCode: "RU",
					Name:        "Asia/Kamchatka",
					TimeZones: []*TimeZone{
						{
							ZoneId: 333,
							Abbr:   "LMT",
							Offset: 38076,
						},
						{
							ZoneId: 333,
							Abbr:   "PETT",
							Offset: 39600,
						},
						{
							ZoneId: 333,
							Abbr:   "PETT",
							Offset: 43200,
						},
						{
							ZoneId: 333,
							Abbr:   "PETST",
							Offset: 46800,
						},
					},
				},
				{
					Id:          327,
					CountryCode: "RU",
					Name:        "Asia/Khandyga",
					TimeZones: []*TimeZone{
						{
							ZoneId: 327,
							Abbr:   "LMT",
							Offset: 32533,
						},
						{
							ZoneId: 327,
							Abbr:   "YAKT",
							Offset: 28800,
						},
						{
							ZoneId: 327,
							Abbr:   "YAKT",
							Offset: 32400,
						},
						{
							ZoneId: 327,
							Abbr:   "YAKST",
							Offset: 36000,
						},
						{
							ZoneId: 327,
							Abbr:   "VLAT",
							Offset: 36000,
						},
						{
							ZoneId: 327,
							Abbr:   "VLAST",
							Offset: 39600,
						},
						{
							ZoneId: 327,
							Abbr:   "VLAT",
							Offset: 39600,
						},
						{
							ZoneId: 327,
							Abbr:   "YAKT",
							Offset: 36000,
						},
					},
				},
				{
					Id:          323,
					CountryCode: "RU",
					Name:        "Asia/Krasnoyarsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 323,
							Abbr:   "LMT",
							Offset: 22286,
						},
						{
							ZoneId: 323,
							Abbr:   "KRAT",
							Offset: 21600,
						},
						{
							ZoneId: 323,
							Abbr:   "KRAT",
							Offset: 25200,
						},
						{
							ZoneId: 323,
							Abbr:   "KRAST",
							Offset: 28800,
						},
						{
							ZoneId: 323,
							Abbr:   "KRAT",
							Offset: 28800,
						},
					},
				},
				{
					Id:          330,
					CountryCode: "RU",
					Name:        "Asia/Magadan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 330,
							Abbr:   "LMT",
							Offset: 36192,
						},
						{
							ZoneId: 330,
							Abbr:   "MAGT",
							Offset: 36000,
						},
						{
							ZoneId: 330,
							Abbr:   "MAGT",
							Offset: 39600,
						},
						{
							ZoneId: 330,
							Abbr:   "MAGST",
							Offset: 43200,
						},
						{
							ZoneId: 330,
							Abbr:   "MAGT",
							Offset: 43200,
						},
					},
				},
				{
					Id:          322,
					CountryCode: "RU",
					Name:        "Asia/Novokuznetsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 322,
							Abbr:   "LMT",
							Offset: 20928,
						},
						{
							ZoneId: 322,
							Abbr:   "KRAT",
							Offset: 21600,
						},
						{
							ZoneId: 322,
							Abbr:   "KRAT",
							Offset: 25200,
						},
						{
							ZoneId: 322,
							Abbr:   "KRAST",
							Offset: 28800,
						},
					},
				},
				{
					Id:          319,
					CountryCode: "RU",
					Name:        "Asia/Novosibirsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 319,
							Abbr:   "LMT",
							Offset: 19900,
						},
						{
							ZoneId: 319,
							Abbr:   "NOVT",
							Offset: 21600,
						},
						{
							ZoneId: 319,
							Abbr:   "NOVT",
							Offset: 25200,
						},
						{
							ZoneId: 319,
							Abbr:   "NOVST",
							Offset: 28800,
						},
						{
							ZoneId: 319,
							Abbr:   "NOVST",
							Offset: 25200,
						},
					},
				},
				{
					Id:          318,
					CountryCode: "RU",
					Name:        "Asia/Omsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 318,
							Abbr:   "LMT",
							Offset: 17610,
						},
						{
							ZoneId: 318,
							Abbr:   "OMST",
							Offset: 18000,
						},
						{
							ZoneId: 318,
							Abbr:   "OMST",
							Offset: 21600,
						},
						{
							ZoneId: 318,
							Abbr:   "OMSST",
							Offset: 25200,
						},
						{
							ZoneId: 318,
							Abbr:   "OMST",
							Offset: 25200,
						},
					},
				},
				{
					Id:          331,
					CountryCode: "RU",
					Name:        "Asia/Sakhalin",
					TimeZones: []*TimeZone{
						{
							ZoneId: 331,
							Abbr:   "LMT",
							Offset: 34248,
						},
						{
							ZoneId: 331,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 331,
							Abbr:   "SAKT",
							Offset: 39600,
						},
						{
							ZoneId: 331,
							Abbr:   "SAKST",
							Offset: 43200,
						},
						{
							ZoneId: 331,
							Abbr:   "SAKT",
							Offset: 36000,
						},
						{
							ZoneId: 331,
							Abbr:   "SAKST",
							Offset: 39600,
						},
					},
				},
				{
					Id:          332,
					CountryCode: "RU",
					Name:        "Asia/Srednekolymsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 332,
							Abbr:   "LMT",
							Offset: 36892,
						},
						{
							ZoneId: 332,
							Abbr:   "MAGT",
							Offset: 36000,
						},
						{
							ZoneId: 332,
							Abbr:   "MAGT",
							Offset: 39600,
						},
						{
							ZoneId: 332,
							Abbr:   "MAGST",
							Offset: 43200,
						},
						{
							ZoneId: 332,
							Abbr:   "MAGT",
							Offset: 43200,
						},
						{
							ZoneId: 332,
							Abbr:   "SRET",
							Offset: 39600,
						},
					},
				},
				{
					Id:          321,
					CountryCode: "RU",
					Name:        "Asia/Tomsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 321,
							Abbr:   "LMT",
							Offset: 20391,
						},
						{
							ZoneId: 321,
							Abbr:   "MSK+3",
							Offset: 21600,
						},
						{
							ZoneId: 321,
							Abbr:   "MSD+3",
							Offset: 25200,
						},
						{
							ZoneId: 321,
							Abbr:   "MSD+4",
							Offset: 28800,
						},
					},
				},
				{
					Id:          329,
					CountryCode: "RU",
					Name:        "Asia/Ust-Nera",
					TimeZones: []*TimeZone{
						{
							ZoneId: 329,
							Abbr:   "LMT",
							Offset: 34374,
						},
						{
							ZoneId: 329,
							Abbr:   "YAKT",
							Offset: 28800,
						},
						{
							ZoneId: 329,
							Abbr:   "YAKT",
							Offset: 32400,
						},
						{
							ZoneId: 329,
							Abbr:   "MAGST",
							Offset: 43200,
						},
						{
							ZoneId: 329,
							Abbr:   "MAGT",
							Offset: 39600,
						},
						{
							ZoneId: 329,
							Abbr:   "MAGT",
							Offset: 36000,
						},
						{
							ZoneId: 329,
							Abbr:   "MAGT",
							Offset: 43200,
						},
						{
							ZoneId: 329,
							Abbr:   "VLAT",
							Offset: 39600,
						},
						{
							ZoneId: 329,
							Abbr:   "VLAT",
							Offset: 36000,
						},
					},
				},
				{
					Id:          328,
					CountryCode: "RU",
					Name:        "Asia/Vladivostok",
					TimeZones: []*TimeZone{
						{
							ZoneId: 328,
							Abbr:   "LMT",
							Offset: 31651,
						},
						{
							ZoneId: 328,
							Abbr:   "VLAT",
							Offset: 32400,
						},
						{
							ZoneId: 328,
							Abbr:   "VLAT",
							Offset: 36000,
						},
						{
							ZoneId: 328,
							Abbr:   "VLAST",
							Offset: 39600,
						},
						{
							ZoneId: 328,
							Abbr:   "VLAT",
							Offset: 39600,
						},
					},
				},
				{
					Id:          326,
					CountryCode: "RU",
					Name:        "Asia/Yakutsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 326,
							Abbr:   "LMT",
							Offset: 31138,
						},
						{
							ZoneId: 326,
							Abbr:   "YAKT",
							Offset: 28800,
						},
						{
							ZoneId: 326,
							Abbr:   "YAKT",
							Offset: 32400,
						},
						{
							ZoneId: 326,
							Abbr:   "YAKST",
							Offset: 36000,
						},
						{
							ZoneId: 326,
							Abbr:   "YAKT",
							Offset: 36000,
						},
					},
				},
				{
					Id:          317,
					CountryCode: "RU",
					Name:        "Asia/Yekaterinburg",
					TimeZones: []*TimeZone{
						{
							ZoneId: 317,
							Abbr:   "LMT",
							Offset: 14553,
						},
						{
							ZoneId: 317,
							Abbr:   "PMT",
							Offset: 13505,
						},
						{
							ZoneId: 317,
							Abbr:   "SVET",
							Offset: 14400,
						},
						{
							ZoneId: 317,
							Abbr:   "SVET",
							Offset: 18000,
						},
						{
							ZoneId: 317,
							Abbr:   "SVEST",
							Offset: 21600,
						},
						{
							ZoneId: 317,
							Abbr:   "YEKT",
							Offset: 18000,
						},
						{
							ZoneId: 317,
							Abbr:   "YEKST",
							Offset: 21600,
						},
						{
							ZoneId: 317,
							Abbr:   "YEKT",
							Offset: 21600,
						},
					},
				},
				{
					Id:          313,
					CountryCode: "RU",
					Name:        "Europe/Astrakhan",
					TimeZones: []*TimeZone{
						{
							ZoneId: 313,
							Abbr:   "LMT",
							Offset: 11532,
						},
						{
							ZoneId: 313,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 313,
							Abbr:   "MSK+1",
							Offset: 14400,
						},
						{
							ZoneId: 313,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
						{
							ZoneId: 313,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 313,
							Abbr:   "MSK",
							Offset: 14400,
						},
					},
				},
				{
					Id:          308,
					CountryCode: "RU",
					Name:        "Europe/Kaliningrad",
					TimeZones: []*TimeZone{
						{
							ZoneId: 308,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 308,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 308,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 308,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 308,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 308,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 308,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 308,
							Abbr:   "EET",
							Offset: 10800,
						},
					},
				},
				{
					Id:          311,
					CountryCode: "RU",
					Name:        "Europe/Kirov",
					TimeZones: []*TimeZone{
						{
							ZoneId: 311,
							Abbr:   "LMT",
							Offset: 11928,
						},
						{
							ZoneId: 311,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 311,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 311,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
					},
				},
				{
					Id:          309,
					CountryCode: "RU",
					Name:        "Europe/Moscow",
					TimeZones: []*TimeZone{
						{
							ZoneId: 309,
							Abbr:   "LMT",
							Offset: 9017,
						},
						{
							ZoneId: 309,
							Abbr:   "MMT",
							Offset: 9017,
						},
						{
							ZoneId: 309,
							Abbr:   "MMT",
							Offset: 9079,
						},
						{
							ZoneId: 309,
							Abbr:   "MST",
							Offset: 12679,
						},
						{
							ZoneId: 309,
							Abbr:   "MDST",
							Offset: 16279,
						},
						{
							ZoneId: 309,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 309,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 309,
							Abbr:   "MSM",
							Offset: 18000,
						},
						{
							ZoneId: 309,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 309,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 309,
							Abbr:   "MSK",
							Offset: 14400,
						},
					},
				},
				{
					Id:          316,
					CountryCode: "RU",
					Name:        "Europe/Samara",
					TimeZones: []*TimeZone{
						{
							ZoneId: 316,
							Abbr:   "LMT",
							Offset: 12020,
						},
						{
							ZoneId: 316,
							Abbr:   "KUYT",
							Offset: 10800,
						},
						{
							ZoneId: 316,
							Abbr:   "KUYT",
							Offset: 14400,
						},
						{
							ZoneId: 316,
							Abbr:   "KUYST",
							Offset: 18000,
						},
						{
							ZoneId: 316,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 316,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 316,
							Abbr:   "SAMT",
							Offset: 14400,
						},
						{
							ZoneId: 316,
							Abbr:   "SAMST",
							Offset: 18000,
						},
						{
							ZoneId: 316,
							Abbr:   "SAMT",
							Offset: 10800,
						},
					},
				},
				{
					Id:          314,
					CountryCode: "RU",
					Name:        "Europe/Saratov",
					TimeZones: []*TimeZone{
						{
							ZoneId: 314,
							Abbr:   "LMT",
							Offset: 11058,
						},
						{
							ZoneId: 314,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 314,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 314,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
					},
				},
				{
					Id:          315,
					CountryCode: "RU",
					Name:        "Europe/Ulyanovsk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 315,
							Abbr:   "LMT",
							Offset: 11616,
						},
						{
							ZoneId: 315,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 315,
							Abbr:   "MSK+1",
							Offset: 14400,
						},
						{
							ZoneId: 315,
							Abbr:   "MSD+1",
							Offset: 18000,
						},
						{
							ZoneId: 315,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 315,
							Abbr:   "MSK",
							Offset: 7200,
						},
						{
							ZoneId: 315,
							Abbr:   "MSK",
							Offset: 14400,
						},
					},
				},
				{
					Id:          312,
					CountryCode: "RU",
					Name:        "Europe/Volgograd",
					TimeZones: []*TimeZone{
						{
							ZoneId: 312,
							Abbr:   "LMT",
							Offset: 10660,
						},
						{
							ZoneId: 312,
							Abbr:   "TSAT",
							Offset: 10800,
						},
						{
							ZoneId: 312,
							Abbr:   "STAT",
							Offset: 14400,
						},
						{
							ZoneId: 312,
							Abbr:   "VOLST",
							Offset: 18000,
						},
						{
							ZoneId: 312,
							Abbr:   "VOLT",
							Offset: 14400,
						},
						{
							ZoneId: 312,
							Abbr:   "VOLT",
							Offset: 10800,
						},
						{
							ZoneId: 312,
							Abbr:   "VOLST",
							Offset: 14400,
						},
						{
							ZoneId: 312,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 312,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 312,
							Abbr:   "MSK",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "RW",
			Name: "Rwanda",
			Zones: []*Zone{
				{
					Id:          335,
					CountryCode: "RW",
					Name:        "Africa/Kigali",
					TimeZones: []*TimeZone{
						{
							ZoneId: 335,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 335,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "RE",
			Name: "Réunion",
			Zones: []*Zone{
				{
					Id:          305,
					CountryCode: "RE",
					Name:        "Indian/Reunion",
					TimeZones: []*TimeZone{
						{
							ZoneId: 305,
							Abbr:   "LMT",
							Offset: 13312,
						},
						{
							ZoneId: 305,
							Abbr:   "RET",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "BL",
			Name: "Saint Barthélemy",
			Zones: []*Zone{
				{
					Id:          57,
					CountryCode: "BL",
					Name:        "America/St_Barthelemy",
					TimeZones: []*TimeZone{
						{
							ZoneId: 57,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 57,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "SH",
			Name: "Saint Helena, Ascension and Tristan da Cunha",
			Zones: []*Zone{
				{
					Id:          342,
					CountryCode: "SH",
					Name:        "Atlantic/St_Helena",
					TimeZones: []*TimeZone{
						{
							ZoneId: 342,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 342,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "KN",
			Name: "Saint Kitts and Nevis",
			Zones: []*Zone{
				{
					Id:          210,
					CountryCode: "KN",
					Name:        "America/St_Kitts",
					TimeZones: []*TimeZone{
						{
							ZoneId: 210,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 210,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "LC",
			Name: "Saint Lucia",
			Zones: []*Zone{
				{
					Id:          224,
					CountryCode: "LC",
					Name:        "America/St_Lucia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 224,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 224,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "MF",
			Name: "Saint Martin (French part)",
			Zones: []*Zone{
				{
					Id:          237,
					CountryCode: "MF",
					Name:        "America/Marigot",
					TimeZones: []*TimeZone{
						{
							ZoneId: 237,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 237,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "PM",
			Name: "Saint Pierre and Miquelon",
			Zones: []*Zone{
				{
					Id:          294,
					CountryCode: "PM",
					Name:        "America/Miquelon",
					TimeZones: []*TimeZone{
						{
							ZoneId: 294,
							Abbr:   "LMT",
							Offset: -13480,
						},
						{
							ZoneId: 294,
							Abbr:   "AST",
							Offset: -14400,
						},
						{
							ZoneId: 294,
							Abbr:   "PMDT",
							Offset: -10800,
						},
						{
							ZoneId: 294,
							Abbr:   "PMST",
							Offset: -7200,
						},
					},
				},
			},
		},
		{
			Code: "VC",
			Name: "Saint Vincent and the Grenadines",
			Zones: []*Zone{
				{
					Id:          412,
					CountryCode: "VC",
					Name:        "America/St_Vincent",
					TimeZones: []*TimeZone{
						{
							ZoneId: 412,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 412,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "WS",
			Name: "Samoa",
			Zones: []*Zone{
				{
					Id:          419,
					CountryCode: "WS",
					Name:        "Pacific/Apia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 419,
							Abbr:   "LMT",
							Offset: 45184,
						},
						{
							ZoneId: 419,
							Abbr:   "LMT",
							Offset: -41216,
						},
						{
							ZoneId: 419,
							Abbr:   "SAMT",
							Offset: -41400,
						},
						{
							ZoneId: 419,
							Abbr:   "WST",
							Offset: -39600,
						},
						{
							ZoneId: 419,
							Abbr:   "WST",
							Offset: -36000,
						},
						{
							ZoneId: 419,
							Abbr:   "WST",
							Offset: 50400,
						},
						{
							ZoneId: 419,
							Abbr:   "WST",
							Offset: 46800,
						},
					},
				},
			},
		},
		{
			Code: "SM",
			Name: "San Marino",
			Zones: []*Zone{
				{
					Id:          347,
					CountryCode: "SM",
					Name:        "Europe/San_Marino",
					TimeZones: []*TimeZone{
						{
							ZoneId: 347,
							Abbr:   "LMT",
							Offset: 2996,
						},
						{
							ZoneId: 347,
							Abbr:   "RMT",
							Offset: 2996,
						},
						{
							ZoneId: 347,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 347,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "ST",
			Name: "Sao Tome and Principe",
			Zones: []*Zone{
				{
					Id:          352,
					CountryCode: "ST",
					Name:        "Africa/Sao_Tome",
					TimeZones: []*TimeZone{
						{
							ZoneId: 352,
							Abbr:   "LMT",
							Offset: 1616,
						},
						{
							ZoneId: 352,
							Abbr:   "LMT",
							Offset: -2205,
						},
						{
							ZoneId: 352,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 352,
							Abbr:   "WAT",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "SA",
			Name: "Saudi Arabia",
			Zones: []*Zone{
				{
					Id:          336,
					CountryCode: "SA",
					Name:        "Asia/Riyadh",
					TimeZones: []*TimeZone{
						{
							ZoneId: 336,
							Abbr:   "LMT",
							Offset: 11212,
						},
						{
							ZoneId: 336,
							Abbr:   "AST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "SN",
			Name: "Senegal",
			Zones: []*Zone{
				{
					Id:          348,
					CountryCode: "SN",
					Name:        "Africa/Dakar",
					TimeZones: []*TimeZone{
						{
							ZoneId: 348,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 348,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "RS",
			Name: "Serbia",
			Zones: []*Zone{
				{
					Id:          307,
					CountryCode: "RS",
					Name:        "Europe/Belgrade",
					TimeZones: []*TimeZone{
						{
							ZoneId: 307,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 307,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 307,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "SC",
			Name: "Seychelles",
			Zones: []*Zone{
				{
					Id:          338,
					CountryCode: "SC",
					Name:        "Indian/Mahe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 338,
							Abbr:   "LMT",
							Offset: 13308,
						},
						{
							ZoneId: 338,
							Abbr:   "SCT",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "SL",
			Name: "Sierra Leone",
			Zones: []*Zone{
				{
					Id:          346,
					CountryCode: "SL",
					Name:        "Africa/Freetown",
					TimeZones: []*TimeZone{
						{
							ZoneId: 346,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 346,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "SG",
			Name: "Singapore",
			Zones: []*Zone{
				{
					Id:          341,
					CountryCode: "SG",
					Name:        "Asia/Singapore",
					TimeZones: []*TimeZone{
						{
							ZoneId: 341,
							Abbr:   "LMT",
							Offset: 24925,
						},
						{
							ZoneId: 341,
							Abbr:   "SMT",
							Offset: 24925,
						},
						{
							ZoneId: 341,
							Abbr:   "SMT",
							Offset: 25200,
						},
						{
							ZoneId: 341,
							Abbr:   "SMT",
							Offset: 26400,
						},
						{
							ZoneId: 341,
							Abbr:   "SMT",
							Offset: 27000,
						},
						{
							ZoneId: 341,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 341,
							Abbr:   "SST",
							Offset: 27000,
						},
						{
							ZoneId: 341,
							Abbr:   "SST",
							Offset: 28800,
						},
					},
				},
			},
		},
		{
			Code: "SX",
			Name: "Sint Maarten (Dutch part)",
			Zones: []*Zone{
				{
					Id:          354,
					CountryCode: "SX",
					Name:        "America/Lower_Princes",
					TimeZones: []*TimeZone{
						{
							ZoneId: 354,
							Abbr:   "LMT",
							Offset: -16547,
						},
						{
							ZoneId: 354,
							Abbr:   "AST",
							Offset: -16200,
						},
						{
							ZoneId: 354,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "SK",
			Name: "Slovakia",
			Zones: []*Zone{
				{
					Id:          345,
					CountryCode: "SK",
					Name:        "Europe/Bratislava",
					TimeZones: []*TimeZone{
						{
							ZoneId: 345,
							Abbr:   "LMT",
							Offset: 3464,
						},
						{
							ZoneId: 345,
							Abbr:   "PMT",
							Offset: 3464,
						},
						{
							ZoneId: 345,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 345,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 345,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "SI",
			Name: "Slovenia",
			Zones: []*Zone{
				{
					Id:          343,
					CountryCode: "SI",
					Name:        "Europe/Ljubljana",
					TimeZones: []*TimeZone{
						{
							ZoneId: 343,
							Abbr:   "LMT",
							Offset: 4920,
						},
						{
							ZoneId: 343,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 343,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "SB",
			Name: "Solomon Islands",
			Zones: []*Zone{
				{
					Id:          337,
					CountryCode: "SB",
					Name:        "Pacific/Guadalcanal",
					TimeZones: []*TimeZone{
						{
							ZoneId: 337,
							Abbr:   "LMT",
							Offset: 38388,
						},
						{
							ZoneId: 337,
							Abbr:   "SBT",
							Offset: 39600,
						},
					},
				},
			},
		},
		{
			Code: "SO",
			Name: "Somalia",
			Zones: []*Zone{
				{
					Id:          349,
					CountryCode: "SO",
					Name:        "Africa/Mogadishu",
					TimeZones: []*TimeZone{
						{
							ZoneId: 349,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 349,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 349,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 349,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "ZA",
			Name: "South Africa",
			Zones: []*Zone{
				{
					Id:          422,
					CountryCode: "ZA",
					Name:        "Africa/Johannesburg",
					TimeZones: []*TimeZone{
						{
							ZoneId: 422,
							Abbr:   "LMT",
							Offset: 6720,
						},
						{
							ZoneId: 422,
							Abbr:   "SAST",
							Offset: 5400,
						},
						{
							ZoneId: 422,
							Abbr:   "SAST",
							Offset: 7200,
						},
						{
							ZoneId: 422,
							Abbr:   "SAST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "GS",
			Name: "South Georgia and the South Sandwich Islands",
			Zones: []*Zone{
				{
					Id:          176,
					CountryCode: "GS",
					Name:        "Atlantic/South_Georgia",
					TimeZones: []*TimeZone{
						{
							ZoneId: 176,
							Abbr:   "LMT",
							Offset: -8768,
						},
						{
							ZoneId: 176,
							Abbr:   "GST",
							Offset: -7200,
						},
					},
				},
			},
		},
		{
			Code: "SS",
			Name: "South Sudan",
			Zones: []*Zone{
				{
					Id:          351,
					CountryCode: "SS",
					Name:        "Africa/Juba",
					TimeZones: []*TimeZone{
						{
							ZoneId: 351,
							Abbr:   "LMT",
							Offset: 7588,
						},
						{
							ZoneId: 351,
							Abbr:   "CAT",
							Offset: 7200,
						},
						{
							ZoneId: 351,
							Abbr:   "CAST",
							Offset: 10800,
						},
						{
							ZoneId: 351,
							Abbr:   "EAT",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "ES",
			Name: "Spain",
			Zones: []*Zone{
				{
					Id:          148,
					CountryCode: "ES",
					Name:        "Africa/Ceuta",
					TimeZones: []*TimeZone{
						{
							ZoneId: 148,
							Abbr:   "LMT",
							Offset: -1276,
						},
						{
							ZoneId: 148,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 148,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 148,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 148,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
				{
					Id:          149,
					CountryCode: "ES",
					Name:        "Atlantic/Canary",
					TimeZones: []*TimeZone{
						{
							ZoneId: 149,
							Abbr:   "LMT",
							Offset: -3696,
						},
						{
							ZoneId: 149,
							Abbr:   "WAT",
							Offset: -3600,
						},
						{
							ZoneId: 149,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 149,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
				{
					Id:          147,
					CountryCode: "ES",
					Name:        "Europe/Madrid",
					TimeZones: []*TimeZone{
						{
							ZoneId: 147,
							Abbr:   "LMT",
							Offset: -884,
						},
						{
							ZoneId: 147,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 147,
							Abbr:   "WEST",
							Offset: 3600,
						},
						{
							ZoneId: 147,
							Abbr:   "WEMT",
							Offset: 7200,
						},
						{
							ZoneId: 147,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 147,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "LK",
			Name: "Sri Lanka",
			Zones: []*Zone{
				{
					Id:          226,
					CountryCode: "LK",
					Name:        "Asia/Colombo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 226,
							Abbr:   "LMT",
							Offset: 19164,
						},
						{
							ZoneId: 226,
							Abbr:   "MMT",
							Offset: 19172,
						},
						{
							ZoneId: 226,
							Abbr:   "IST",
							Offset: 19800,
						},
						{
							ZoneId: 226,
							Abbr:   "IST",
							Offset: 21600,
						},
						{
							ZoneId: 226,
							Abbr:   "IST",
							Offset: 23400,
						},
					},
				},
			},
		},
		{
			Code: "SD",
			Name: "Sudan",
			Zones: []*Zone{
				{
					Id:          339,
					CountryCode: "SD",
					Name:        "Africa/Khartoum",
					TimeZones: []*TimeZone{
						{
							ZoneId: 339,
							Abbr:   "LMT",
							Offset: 7808,
						},
						{
							ZoneId: 339,
							Abbr:   "CAT",
							Offset: 7200,
						},
						{
							ZoneId: 339,
							Abbr:   "CAST",
							Offset: 10800,
						},
						{
							ZoneId: 339,
							Abbr:   "EAT",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "SR",
			Name: "Suriname",
			Zones: []*Zone{
				{
					Id:          350,
					CountryCode: "SR",
					Name:        "America/Paramaribo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 350,
							Abbr:   "LMT",
							Offset: -13240,
						},
						{
							ZoneId: 350,
							Abbr:   "PMT",
							Offset: -13252,
						},
						{
							ZoneId: 350,
							Abbr:   "PMT",
							Offset: -13236,
						},
						{
							ZoneId: 350,
							Abbr:   "SRT",
							Offset: -12600,
						},
						{
							ZoneId: 350,
							Abbr:   "SRT",
							Offset: -10800,
						},
					},
				},
			},
		},
		{
			Code: "SJ",
			Name: "Svalbard and Jan Mayen",
			Zones: []*Zone{
				{
					Id:          344,
					CountryCode: "SJ",
					Name:        "Arctic/Longyearbyen",
					TimeZones: []*TimeZone{
						{
							ZoneId: 344,
							Abbr:   "LMT",
							Offset: 2580,
						},
						{
							ZoneId: 344,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 344,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "SZ",
			Name: "Swaziland",
			Zones: []*Zone{
				{
					Id:          356,
					CountryCode: "SZ",
					Name:        "Africa/Mbabane",
					TimeZones: []*TimeZone{
						{
							ZoneId: 356,
							Abbr:   "LMT",
							Offset: 6720,
						},
						{
							ZoneId: 356,
							Abbr:   "SAST",
							Offset: 5400,
						},
						{
							ZoneId: 356,
							Abbr:   "SAST",
							Offset: 7200,
						},
						{
							ZoneId: 356,
							Abbr:   "SAST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "SE",
			Name: "Sweden",
			Zones: []*Zone{
				{
					Id:          340,
					CountryCode: "SE",
					Name:        "Europe/Stockholm",
					TimeZones: []*TimeZone{
						{
							ZoneId: 340,
							Abbr:   "LMT",
							Offset: 4332,
						},
						{
							ZoneId: 340,
							Abbr:   "SET",
							Offset: 3614,
						},
						{
							ZoneId: 340,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 340,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "CH",
			Name: "Switzerland",
			Zones: []*Zone{
				{
					Id:          116,
					CountryCode: "CH",
					Name:        "Europe/Zurich",
					TimeZones: []*TimeZone{
						{
							ZoneId: 116,
							Abbr:   "LMT",
							Offset: 2048,
						},
						{
							ZoneId: 116,
							Abbr:   "BMT",
							Offset: 1786,
						},
						{
							ZoneId: 116,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 116,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "SY",
			Name: "Syrian Arab Republic",
			Zones: []*Zone{
				{
					Id:          355,
					CountryCode: "SY",
					Name:        "Asia/Damascus",
					TimeZones: []*TimeZone{
						{
							ZoneId: 355,
							Abbr:   "LMT",
							Offset: 8712,
						},
						{
							ZoneId: 355,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 355,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "TW",
			Name: "Taiwan, Province of China",
			Zones: []*Zone{
				{
					Id:          371,
					CountryCode: "TW",
					Name:        "Asia/Taipei",
					TimeZones: []*TimeZone{
						{
							ZoneId: 371,
							Abbr:   "LMT",
							Offset: 29160,
						},
						{
							ZoneId: 371,
							Abbr:   "CST",
							Offset: 28800,
						},
						{
							ZoneId: 371,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 371,
							Abbr:   "CDT",
							Offset: 32400,
						},
					},
				},
			},
		},
		{
			Code: "TJ",
			Name: "Tajikistan",
			Zones: []*Zone{
				{
					Id:          362,
					CountryCode: "TJ",
					Name:        "Asia/Dushanbe",
					TimeZones: []*TimeZone{
						{
							ZoneId: 362,
							Abbr:   "LMT",
							Offset: 16512,
						},
						{
							ZoneId: 362,
							Abbr:   "DUST",
							Offset: 18000,
						},
						{
							ZoneId: 362,
							Abbr:   "DUST",
							Offset: 21600,
						},
						{
							ZoneId: 362,
							Abbr:   "DUSST",
							Offset: 25200,
						},
						{
							ZoneId: 362,
							Abbr:   "TJT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "TZ",
			Name: "Tanzania, United Republic of",
			Zones: []*Zone{
				{
					Id:          372,
					CountryCode: "TZ",
					Name:        "Africa/Dar_es_Salaam",
					TimeZones: []*TimeZone{
						{
							ZoneId: 372,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 372,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 372,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 372,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "TH",
			Name: "Thailand",
			Zones: []*Zone{
				{
					Id:          361,
					CountryCode: "TH",
					Name:        "Asia/Bangkok",
					TimeZones: []*TimeZone{
						{
							ZoneId: 361,
							Abbr:   "LMT",
							Offset: 24124,
						},
						{
							ZoneId: 361,
							Abbr:   "BMT",
							Offset: 24124,
						},
						{
							ZoneId: 361,
							Abbr:   "ICT",
							Offset: 25200,
						},
					},
				},
			},
		},
		{
			Code: "TL",
			Name: "Timor-Leste",
			Zones: []*Zone{
				{
					Id:          364,
					CountryCode: "TL",
					Name:        "Asia/Dili",
					TimeZones: []*TimeZone{
						{
							ZoneId: 364,
							Abbr:   "LMT",
							Offset: 30140,
						},
						{
							ZoneId: 364,
							Abbr:   "TLT",
							Offset: 28800,
						},
						{
							ZoneId: 364,
							Abbr:   "TLT",
							Offset: 32400,
						},
						{
							ZoneId: 364,
							Abbr:   "WITA",
							Offset: 28800,
						},
					},
				},
			},
		},
		{
			Code: "TG",
			Name: "Togo",
			Zones: []*Zone{
				{
					Id:          360,
					CountryCode: "TG",
					Name:        "Africa/Lome",
					TimeZones: []*TimeZone{
						{
							ZoneId: 360,
							Abbr:   "LMT",
							Offset: -968,
						},
						{
							ZoneId: 360,
							Abbr:   "GMT",
							Offset: 0,
						},
					},
				},
			},
		},
		{
			Code: "TK",
			Name: "Tokelau",
			Zones: []*Zone{
				{
					Id:          363,
					CountryCode: "TK",
					Name:        "Pacific/Fakaofo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 363,
							Abbr:   "LMT",
							Offset: -41096,
						},
						{
							ZoneId: 363,
							Abbr:   "TKT",
							Offset: -39600,
						},
						{
							ZoneId: 363,
							Abbr:   "TKT",
							Offset: 46800,
						},
					},
				},
			},
		},
		{
			Code: "TO",
			Name: "Tonga",
			Zones: []*Zone{
				{
					Id:          367,
					CountryCode: "TO",
					Name:        "Pacific/Tongatapu",
					TimeZones: []*TimeZone{
						{
							ZoneId: 367,
							Abbr:   "LMT",
							Offset: 44360,
						},
						{
							ZoneId: 367,
							Abbr:   "TOT",
							Offset: 44400,
						},
						{
							ZoneId: 367,
							Abbr:   "TOT",
							Offset: 46800,
						},
						{
							ZoneId: 367,
							Abbr:   "TOST",
							Offset: 50400,
						},
					},
				},
			},
		},
		{
			Code: "TT",
			Name: "Trinidad and Tobago",
			Zones: []*Zone{
				{
					Id:          369,
					CountryCode: "TT",
					Name:        "America/Port_of_Spain",
					TimeZones: []*TimeZone{
						{
							ZoneId: 369,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 369,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "TN",
			Name: "Tunisia",
			Zones: []*Zone{
				{
					Id:          366,
					CountryCode: "TN",
					Name:        "Africa/Tunis",
					TimeZones: []*TimeZone{
						{
							ZoneId: 366,
							Abbr:   "LMT",
							Offset: 2444,
						},
						{
							ZoneId: 366,
							Abbr:   "PMT",
							Offset: 561,
						},
						{
							ZoneId: 366,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 366,
							Abbr:   "CEST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "TR",
			Name: "Turkey",
			Zones: []*Zone{
				{
					Id:          368,
					CountryCode: "TR",
					Name:        "Europe/Istanbul",
					TimeZones: []*TimeZone{
						{
							ZoneId: 368,
							Abbr:   "LMT",
							Offset: 6952,
						},
						{
							ZoneId: 368,
							Abbr:   "IMT",
							Offset: 7016,
						},
						{
							ZoneId: 368,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 368,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 368,
							Abbr:   "TRT",
							Offset: 10800,
						},
						{
							ZoneId: 368,
							Abbr:   "TRST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "TM",
			Name: "Turkmenistan",
			Zones: []*Zone{
				{
					Id:          365,
					CountryCode: "TM",
					Name:        "Asia/Ashgabat",
					TimeZones: []*TimeZone{
						{
							ZoneId: 365,
							Abbr:   "LMT",
							Offset: 14012,
						},
						{
							ZoneId: 365,
							Abbr:   "ASHT",
							Offset: 14400,
						},
						{
							ZoneId: 365,
							Abbr:   "ASHT",
							Offset: 18000,
						},
						{
							ZoneId: 365,
							Abbr:   "ASHST",
							Offset: 21600,
						},
						{
							ZoneId: 365,
							Abbr:   "TMT",
							Offset: 14400,
						},
						{
							ZoneId: 365,
							Abbr:   "TMT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "TC",
			Name: "Turks and Caicos Islands",
			Zones: []*Zone{
				{
					Id:          357,
					CountryCode: "TC",
					Name:        "America/Grand_Turk",
					TimeZones: []*TimeZone{
						{
							ZoneId: 357,
							Abbr:   "LMT",
							Offset: -17072,
						},
						{
							ZoneId: 357,
							Abbr:   "KMT",
							Offset: -18430,
						},
						{
							ZoneId: 357,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 357,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 357,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "TV",
			Name: "Tuvalu",
			Zones: []*Zone{
				{
					Id:          370,
					CountryCode: "TV",
					Name:        "Pacific/Funafuti",
					TimeZones: []*TimeZone{
						{
							ZoneId: 370,
							Abbr:   "LMT",
							Offset: 43012,
						},
						{
							ZoneId: 370,
							Abbr:   "TVT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "UG",
			Name: "Uganda",
			Zones: []*Zone{
				{
					Id:          376,
					CountryCode: "UG",
					Name:        "Africa/Kampala",
					TimeZones: []*TimeZone{
						{
							ZoneId: 376,
							Abbr:   "LMT",
							Offset: 8836,
						},
						{
							ZoneId: 376,
							Abbr:   "EAT",
							Offset: 9000,
						},
						{
							ZoneId: 376,
							Abbr:   "EAT",
							Offset: 10800,
						},
						{
							ZoneId: 376,
							Abbr:   "EAT",
							Offset: 9900,
						},
					},
				},
			},
		},
		{
			Code: "UA",
			Name: "Ukraine",
			Zones: []*Zone{
				{
					Id:          373,
					CountryCode: "UA",
					Name:        "Europe/Kiev",
					TimeZones: []*TimeZone{
						{
							ZoneId: 373,
							Abbr:   "LMT",
							Offset: 7324,
						},
						{
							ZoneId: 373,
							Abbr:   "KMT",
							Offset: 7324,
						},
						{
							ZoneId: 373,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 373,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 373,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 373,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 373,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 373,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
				{
					Id:          310,
					CountryCode: "UA",
					Name:        "Europe/Simferopol",
					TimeZones: []*TimeZone{
						{
							ZoneId: 310,
							Abbr:   "LMT",
							Offset: 8184,
						},
						{
							ZoneId: 310,
							Abbr:   "SMT",
							Offset: 8160,
						},
						{
							ZoneId: 310,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 310,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 310,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 310,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 310,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 310,
							Abbr:   "EEST",
							Offset: 10800,
						},
						{
							ZoneId: 310,
							Abbr:   "MSK",
							Offset: 14400,
						},
					},
				},
				{
					Id:          374,
					CountryCode: "UA",
					Name:        "Europe/Uzhgorod",
					TimeZones: []*TimeZone{
						{
							ZoneId: 374,
							Abbr:   "LMT",
							Offset: 5352,
						},
						{
							ZoneId: 374,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 374,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 374,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 374,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 374,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 374,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
				{
					Id:          375,
					CountryCode: "UA",
					Name:        "Europe/Zaporozhye",
					TimeZones: []*TimeZone{
						{
							ZoneId: 375,
							Abbr:   "LMT",
							Offset: 8440,
						},
						{
							ZoneId: 375,
							Abbr:   "EET",
							Offset: 8400,
						},
						{
							ZoneId: 375,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 375,
							Abbr:   "MSK",
							Offset: 10800,
						},
						{
							ZoneId: 375,
							Abbr:   "CEST",
							Offset: 7200,
						},
						{
							ZoneId: 375,
							Abbr:   "CET",
							Offset: 3600,
						},
						{
							ZoneId: 375,
							Abbr:   "MSD",
							Offset: 14400,
						},
						{
							ZoneId: 375,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "AE",
			Name: "United Arab Emirates",
			Zones: []*Zone{
				{
					Id:          2,
					CountryCode: "AE",
					Name:        "Asia/Dubai",
					TimeZones: []*TimeZone{
						{
							ZoneId: 2,
							Abbr:   "LMT",
							Offset: 13272,
						},
						{
							ZoneId: 2,
							Abbr:   "GST",
							Offset: 14400,
						},
					},
				},
			},
		},
		{
			Code: "GB",
			Name: "United Kingdom",
			Zones: []*Zone{
				{
					Id:          160,
					CountryCode: "GB",
					Name:        "Europe/London",
					TimeZones: []*TimeZone{
						{
							ZoneId: 160,
							Abbr:   "LMT",
							Offset: -75,
						},
						{
							ZoneId: 160,
							Abbr:   "GMT",
							Offset: 0,
						},
						{
							ZoneId: 160,
							Abbr:   "BST",
							Offset: 3600,
						},
						{
							ZoneId: 160,
							Abbr:   "BDST",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "US",
			Name: "United States",
			Zones: []*Zone{
				{
					Id:          406,
					CountryCode: "US",
					Name:        "America/Adak",
					TimeZones: []*TimeZone{
						{
							ZoneId: 406,
							Abbr:   "LMT",
							Offset: 44002,
						},
						{
							ZoneId: 406,
							Abbr:   "LMT",
							Offset: -42398,
						},
						{
							ZoneId: 406,
							Abbr:   "NST",
							Offset: -39600,
						},
						{
							ZoneId: 406,
							Abbr:   "NWT",
							Offset: -36000,
						},
						{
							ZoneId: 406,
							Abbr:   "NPT",
							Offset: -36000,
						},
						{
							ZoneId: 406,
							Abbr:   "BST",
							Offset: -39600,
						},
						{
							ZoneId: 406,
							Abbr:   "BDT",
							Offset: -36000,
						},
						{
							ZoneId: 406,
							Abbr:   "AHST",
							Offset: -36000,
						},
						{
							ZoneId: 406,
							Abbr:   "HST",
							Offset: -36000,
						},
						{
							ZoneId: 406,
							Abbr:   "HDT",
							Offset: -32400,
						},
					},
				},
				{
					Id:          400,
					CountryCode: "US",
					Name:        "America/Anchorage",
					TimeZones: []*TimeZone{
						{
							ZoneId: 400,
							Abbr:   "LMT",
							Offset: 50424,
						},
						{
							ZoneId: 400,
							Abbr:   "LMT",
							Offset: -35976,
						},
						{
							ZoneId: 400,
							Abbr:   "AST",
							Offset: -36000,
						},
						{
							ZoneId: 400,
							Abbr:   "AWT",
							Offset: -32400,
						},
						{
							ZoneId: 400,
							Abbr:   "APT",
							Offset: -32400,
						},
						{
							ZoneId: 400,
							Abbr:   "AHST",
							Offset: -36000,
						},
						{
							ZoneId: 400,
							Abbr:   "AHDT",
							Offset: -32400,
						},
						{
							ZoneId: 400,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 400,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 400,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          397,
					CountryCode: "US",
					Name:        "America/Boise",
					TimeZones: []*TimeZone{
						{
							ZoneId: 397,
							Abbr:   "LMT",
							Offset: -27889,
						},
						{
							ZoneId: 397,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 397,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 397,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 397,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 397,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 397,
							Abbr:   "MDT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          389,
					CountryCode: "US",
					Name:        "America/Chicago",
					TimeZones: []*TimeZone{
						{
							ZoneId: 389,
							Abbr:   "LMT",
							Offset: -21036,
						},
						{
							ZoneId: 389,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 389,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 389,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 389,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 389,
							Abbr:   "CPT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          396,
					CountryCode: "US",
					Name:        "America/Denver",
					TimeZones: []*TimeZone{
						{
							ZoneId: 396,
							Abbr:   "LMT",
							Offset: -25196,
						},
						{
							ZoneId: 396,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 396,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 396,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 396,
							Abbr:   "MPT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          380,
					CountryCode: "US",
					Name:        "America/Detroit",
					TimeZones: []*TimeZone{
						{
							ZoneId: 380,
							Abbr:   "LMT",
							Offset: -19931,
						},
						{
							ZoneId: 380,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 380,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 380,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 380,
							Abbr:   "EPT",
							Offset: -14400,
						},
						{
							ZoneId: 380,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          383,
					CountryCode: "US",
					Name:        "America/Indiana/Indianapolis",
					TimeZones: []*TimeZone{
						{
							ZoneId: 383,
							Abbr:   "LMT",
							Offset: -20678,
						},
						{
							ZoneId: 383,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 383,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 383,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 383,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 383,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 383,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          391,
					CountryCode: "US",
					Name:        "America/Indiana/Knox",
					TimeZones: []*TimeZone{
						{
							ZoneId: 391,
							Abbr:   "LMT",
							Offset: -20790,
						},
						{
							ZoneId: 391,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 391,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 391,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 391,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 391,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          386,
					CountryCode: "US",
					Name:        "America/Indiana/Marengo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 386,
							Abbr:   "LMT",
							Offset: -20723,
						},
						{
							ZoneId: 386,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 386,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 386,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 386,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 386,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 386,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          387,
					CountryCode: "US",
					Name:        "America/Indiana/Petersburg",
					TimeZones: []*TimeZone{
						{
							ZoneId: 387,
							Abbr:   "LMT",
							Offset: -20947,
						},
						{
							ZoneId: 387,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 387,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 387,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 387,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 387,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 387,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          390,
					CountryCode: "US",
					Name:        "America/Indiana/Tell_City",
					TimeZones: []*TimeZone{
						{
							ZoneId: 390,
							Abbr:   "LMT",
							Offset: -20823,
						},
						{
							ZoneId: 390,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 390,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 390,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 390,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 390,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 390,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          388,
					CountryCode: "US",
					Name:        "America/Indiana/Vevay",
					TimeZones: []*TimeZone{
						{
							ZoneId: 388,
							Abbr:   "LMT",
							Offset: -20416,
						},
						{
							ZoneId: 388,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 388,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 388,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 388,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 388,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 388,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          384,
					CountryCode: "US",
					Name:        "America/Indiana/Vincennes",
					TimeZones: []*TimeZone{
						{
							ZoneId: 384,
							Abbr:   "LMT",
							Offset: -21007,
						},
						{
							ZoneId: 384,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 384,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 384,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 384,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 384,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 384,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          385,
					CountryCode: "US",
					Name:        "America/Indiana/Winamac",
					TimeZones: []*TimeZone{
						{
							ZoneId: 385,
							Abbr:   "LMT",
							Offset: -20785,
						},
						{
							ZoneId: 385,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 385,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 385,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 385,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 385,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 385,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          401,
					CountryCode: "US",
					Name:        "America/Juneau",
					TimeZones: []*TimeZone{
						{
							ZoneId: 401,
							Abbr:   "LMT",
							Offset: 54139,
						},
						{
							ZoneId: 401,
							Abbr:   "LMT",
							Offset: -32261,
						},
						{
							ZoneId: 401,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 401,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 401,
							Abbr:   "PPT",
							Offset: -25200,
						},
						{
							ZoneId: 401,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 401,
							Abbr:   "YDT",
							Offset: -28800,
						},
						{
							ZoneId: 401,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 401,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 401,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          381,
					CountryCode: "US",
					Name:        "America/Kentucky/Louisville",
					TimeZones: []*TimeZone{
						{
							ZoneId: 381,
							Abbr:   "LMT",
							Offset: -20582,
						},
						{
							ZoneId: 381,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 381,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 381,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 381,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 381,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 381,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          382,
					CountryCode: "US",
					Name:        "America/Kentucky/Monticello",
					TimeZones: []*TimeZone{
						{
							ZoneId: 382,
							Abbr:   "LMT",
							Offset: -20364,
						},
						{
							ZoneId: 382,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 382,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 382,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 382,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 382,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 382,
							Abbr:   "EDT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          399,
					CountryCode: "US",
					Name:        "America/Los_Angeles",
					TimeZones: []*TimeZone{
						{
							ZoneId: 399,
							Abbr:   "LMT",
							Offset: -28378,
						},
						{
							ZoneId: 399,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 399,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 399,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 399,
							Abbr:   "PPT",
							Offset: -25200,
						},
					},
				},
				{
					Id:          392,
					CountryCode: "US",
					Name:        "America/Menominee",
					TimeZones: []*TimeZone{
						{
							ZoneId: 392,
							Abbr:   "LMT",
							Offset: -21027,
						},
						{
							ZoneId: 392,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 392,
							Abbr:   "CDT",
							Offset: -18000,
						},
						{
							ZoneId: 392,
							Abbr:   "CWT",
							Offset: -18000,
						},
						{
							ZoneId: 392,
							Abbr:   "CPT",
							Offset: -18000,
						},
						{
							ZoneId: 392,
							Abbr:   "EST",
							Offset: -18000,
						},
					},
				},
				{
					Id:          403,
					CountryCode: "US",
					Name:        "America/Metlakatla",
					TimeZones: []*TimeZone{
						{
							ZoneId: 403,
							Abbr:   "LMT",
							Offset: 54822,
						},
						{
							ZoneId: 403,
							Abbr:   "LMT",
							Offset: -31578,
						},
						{
							ZoneId: 403,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 403,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 403,
							Abbr:   "PPT",
							Offset: -25200,
						},
						{
							ZoneId: 403,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 403,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 403,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          379,
					CountryCode: "US",
					Name:        "America/New_York",
					TimeZones: []*TimeZone{
						{
							ZoneId: 379,
							Abbr:   "LMT",
							Offset: -17762,
						},
						{
							ZoneId: 379,
							Abbr:   "EST",
							Offset: -18000,
						},
						{
							ZoneId: 379,
							Abbr:   "EDT",
							Offset: -14400,
						},
						{
							ZoneId: 379,
							Abbr:   "EWT",
							Offset: -14400,
						},
						{
							ZoneId: 379,
							Abbr:   "EPT",
							Offset: -14400,
						},
					},
				},
				{
					Id:          405,
					CountryCode: "US",
					Name:        "America/Nome",
					TimeZones: []*TimeZone{
						{
							ZoneId: 405,
							Abbr:   "LMT",
							Offset: 46702,
						},
						{
							ZoneId: 405,
							Abbr:   "LMT",
							Offset: -39698,
						},
						{
							ZoneId: 405,
							Abbr:   "NST",
							Offset: -39600,
						},
						{
							ZoneId: 405,
							Abbr:   "NWT",
							Offset: -36000,
						},
						{
							ZoneId: 405,
							Abbr:   "NPT",
							Offset: -36000,
						},
						{
							ZoneId: 405,
							Abbr:   "BST",
							Offset: -39600,
						},
						{
							ZoneId: 405,
							Abbr:   "BDT",
							Offset: -36000,
						},
						{
							ZoneId: 405,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 405,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 405,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          395,
					CountryCode: "US",
					Name:        "America/North_Dakota/Beulah",
					TimeZones: []*TimeZone{
						{
							ZoneId: 395,
							Abbr:   "LMT",
							Offset: -24427,
						},
						{
							ZoneId: 395,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 395,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 395,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 395,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 395,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 395,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          393,
					CountryCode: "US",
					Name:        "America/North_Dakota/Center",
					TimeZones: []*TimeZone{
						{
							ZoneId: 393,
							Abbr:   "LMT",
							Offset: -24312,
						},
						{
							ZoneId: 393,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 393,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 393,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 393,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 393,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 393,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          394,
					CountryCode: "US",
					Name:        "America/North_Dakota/New_Salem",
					TimeZones: []*TimeZone{
						{
							ZoneId: 394,
							Abbr:   "LMT",
							Offset: -24339,
						},
						{
							ZoneId: 394,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 394,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 394,
							Abbr:   "MWT",
							Offset: -21600,
						},
						{
							ZoneId: 394,
							Abbr:   "MPT",
							Offset: -21600,
						},
						{
							ZoneId: 394,
							Abbr:   "CST",
							Offset: -21600,
						},
						{
							ZoneId: 394,
							Abbr:   "CDT",
							Offset: -18000,
						},
					},
				},
				{
					Id:          398,
					CountryCode: "US",
					Name:        "America/Phoenix",
					TimeZones: []*TimeZone{
						{
							ZoneId: 398,
							Abbr:   "LMT",
							Offset: -26898,
						},
						{
							ZoneId: 398,
							Abbr:   "MST",
							Offset: -25200,
						},
						{
							ZoneId: 398,
							Abbr:   "MDT",
							Offset: -21600,
						},
						{
							ZoneId: 398,
							Abbr:   "MWT",
							Offset: -21600,
						},
					},
				},
				{
					Id:          402,
					CountryCode: "US",
					Name:        "America/Sitka",
					TimeZones: []*TimeZone{
						{
							ZoneId: 402,
							Abbr:   "LMT",
							Offset: 53927,
						},
						{
							ZoneId: 402,
							Abbr:   "LMT",
							Offset: -32473,
						},
						{
							ZoneId: 402,
							Abbr:   "PST",
							Offset: -28800,
						},
						{
							ZoneId: 402,
							Abbr:   "PWT",
							Offset: -25200,
						},
						{
							ZoneId: 402,
							Abbr:   "PPT",
							Offset: -25200,
						},
						{
							ZoneId: 402,
							Abbr:   "PDT",
							Offset: -25200,
						},
						{
							ZoneId: 402,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 402,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 402,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          404,
					CountryCode: "US",
					Name:        "America/Yakutat",
					TimeZones: []*TimeZone{
						{
							ZoneId: 404,
							Abbr:   "LMT",
							Offset: 52865,
						},
						{
							ZoneId: 404,
							Abbr:   "LMT",
							Offset: -33535,
						},
						{
							ZoneId: 404,
							Abbr:   "YST",
							Offset: -32400,
						},
						{
							ZoneId: 404,
							Abbr:   "YWT",
							Offset: -28800,
						},
						{
							ZoneId: 404,
							Abbr:   "YPT",
							Offset: -28800,
						},
						{
							ZoneId: 404,
							Abbr:   "YDT",
							Offset: -28800,
						},
						{
							ZoneId: 404,
							Abbr:   "AKST",
							Offset: -32400,
						},
						{
							ZoneId: 404,
							Abbr:   "AKDT",
							Offset: -28800,
						},
					},
				},
				{
					Id:          407,
					CountryCode: "US",
					Name:        "Pacific/Honolulu",
					TimeZones: []*TimeZone{
						{
							ZoneId: 407,
							Abbr:   "LMT",
							Offset: -37886,
						},
						{
							ZoneId: 407,
							Abbr:   "HST",
							Offset: -37800,
						},
						{
							ZoneId: 407,
							Abbr:   "HDT",
							Offset: -34200,
						},
						{
							ZoneId: 407,
							Abbr:   "HWT",
							Offset: -34200,
						},
						{
							ZoneId: 407,
							Abbr:   "HPT",
							Offset: -34200,
						},
						{
							ZoneId: 407,
							Abbr:   "HST",
							Offset: -36000,
						},
					},
				},
			},
		},
		{
			Code: "UM",
			Name: "United States Minor Outlying Islands",
			Zones: []*Zone{
				{
					Id:          377,
					CountryCode: "UM",
					Name:        "Pacific/Midway",
					TimeZones: []*TimeZone{
						{
							ZoneId: 377,
							Abbr:   "LMT",
							Offset: 45432,
						},
						{
							ZoneId: 377,
							Abbr:   "LMT",
							Offset: -40968,
						},
						{
							ZoneId: 377,
							Abbr:   "SST",
							Offset: -39600,
						},
					},
				},
				{
					Id:          378,
					CountryCode: "UM",
					Name:        "Pacific/Wake",
					TimeZones: []*TimeZone{
						{
							ZoneId: 378,
							Abbr:   "LMT",
							Offset: 39988,
						},
						{
							ZoneId: 378,
							Abbr:   "WAKT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "UY",
			Name: "Uruguay",
			Zones: []*Zone{
				{
					Id:          408,
					CountryCode: "UY",
					Name:        "America/Montevideo",
					TimeZones: []*TimeZone{
						{
							ZoneId: 408,
							Abbr:   "LMT",
							Offset: -13491,
						},
						{
							ZoneId: 408,
							Abbr:   "MMT",
							Offset: -13491,
						},
						{
							ZoneId: 408,
							Abbr:   "UYT",
							Offset: -14400,
						},
						{
							ZoneId: 408,
							Abbr:   "UYST",
							Offset: -10800,
						},
						{
							ZoneId: 408,
							Abbr:   "UYT",
							Offset: -12600,
						},
						{
							ZoneId: 408,
							Abbr:   "UYST",
							Offset: -9000,
						},
						{
							ZoneId: 408,
							Abbr:   "UYT",
							Offset: -10800,
						},
						{
							ZoneId: 408,
							Abbr:   "UYST",
							Offset: -7200,
						},
						{
							ZoneId: 408,
							Abbr:   "UYST",
							Offset: -5400,
						},
					},
				},
			},
		},
		{
			Code: "UZ",
			Name: "Uzbekistan",
			Zones: []*Zone{
				{
					Id:          409,
					CountryCode: "UZ",
					Name:        "Asia/Samarkand",
					TimeZones: []*TimeZone{
						{
							ZoneId: 409,
							Abbr:   "LMT",
							Offset: 16073,
						},
						{
							ZoneId: 409,
							Abbr:   "SAMT",
							Offset: 14400,
						},
						{
							ZoneId: 409,
							Abbr:   "SAMT",
							Offset: 18000,
						},
						{
							ZoneId: 409,
							Abbr:   "SAMST",
							Offset: 21600,
						},
						{
							ZoneId: 409,
							Abbr:   "UZT",
							Offset: 18000,
						},
					},
				},
				{
					Id:          410,
					CountryCode: "UZ",
					Name:        "Asia/Tashkent",
					TimeZones: []*TimeZone{
						{
							ZoneId: 410,
							Abbr:   "LMT",
							Offset: 16631,
						},
						{
							ZoneId: 410,
							Abbr:   "TAST",
							Offset: 18000,
						},
						{
							ZoneId: 410,
							Abbr:   "TAST",
							Offset: 21600,
						},
						{
							ZoneId: 410,
							Abbr:   "TASST",
							Offset: 25200,
						},
						{
							ZoneId: 410,
							Abbr:   "UZT",
							Offset: 18000,
						},
					},
				},
			},
		},
		{
			Code: "VU",
			Name: "Vanuatu",
			Zones: []*Zone{
				{
					Id:          417,
					CountryCode: "VU",
					Name:        "Pacific/Efate",
					TimeZones: []*TimeZone{
						{
							ZoneId: 417,
							Abbr:   "LMT",
							Offset: 40396,
						},
						{
							ZoneId: 417,
							Abbr:   "VUT",
							Offset: 39600,
						},
						{
							ZoneId: 417,
							Abbr:   "VUST",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "VE",
			Name: "Venezuela, Bolivarian Republic of",
			Zones: []*Zone{
				{
					Id:          413,
					CountryCode: "VE",
					Name:        "America/Caracas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 413,
							Abbr:   "LMT",
							Offset: -16064,
						},
						{
							ZoneId: 413,
							Abbr:   "CMT",
							Offset: -16060,
						},
						{
							ZoneId: 413,
							Abbr:   "VET",
							Offset: -16200,
						},
						{
							ZoneId: 413,
							Abbr:   "VET",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "VN",
			Name: "Viet Nam",
			Zones: []*Zone{
				{
					Id:          416,
					CountryCode: "VN",
					Name:        "Asia/Ho_Chi_Minh",
					TimeZones: []*TimeZone{
						{
							ZoneId: 416,
							Abbr:   "LMT",
							Offset: 25600,
						},
						{
							ZoneId: 416,
							Abbr:   "PLMT",
							Offset: 25590,
						},
						{
							ZoneId: 416,
							Abbr:   "ICT",
							Offset: 25200,
						},
						{
							ZoneId: 416,
							Abbr:   "IDT",
							Offset: 28800,
						},
						{
							ZoneId: 416,
							Abbr:   "JST",
							Offset: 32400,
						},
						{
							ZoneId: 416,
							Abbr:   "IDT",
							Offset: 25200,
						},
						{
							ZoneId: 416,
							Abbr:   "ICT",
							Offset: 28800,
						},
					},
				},
			},
		},
		{
			Code: "VG",
			Name: "Virgin Islands, British",
			Zones: []*Zone{
				{
					Id:          414,
					CountryCode: "VG",
					Name:        "America/Tortola",
					TimeZones: []*TimeZone{
						{
							ZoneId: 414,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 414,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "VI",
			Name: "Virgin Islands, U.S.",
			Zones: []*Zone{
				{
					Id:          415,
					CountryCode: "VI",
					Name:        "America/St_Thomas",
					TimeZones: []*TimeZone{
						{
							ZoneId: 415,
							Abbr:   "LMT",
							Offset: -14764,
						},
						{
							ZoneId: 415,
							Abbr:   "AST",
							Offset: -14400,
						},
					},
				},
			},
		},
		{
			Code: "WF",
			Name: "Wallis and Futuna",
			Zones: []*Zone{
				{
					Id:          418,
					CountryCode: "WF",
					Name:        "Pacific/Wallis",
					TimeZones: []*TimeZone{
						{
							ZoneId: 418,
							Abbr:   "LMT",
							Offset: 44120,
						},
						{
							ZoneId: 418,
							Abbr:   "WFT",
							Offset: 43200,
						},
					},
				},
			},
		},
		{
			Code: "EH",
			Name: "Western Sahara",
			Zones: []*Zone{
				{
					Id:          145,
					CountryCode: "EH",
					Name:        "Africa/El_Aaiun",
					TimeZones: []*TimeZone{
						{
							ZoneId: 145,
							Abbr:   "LMT",
							Offset: -3168,
						},
						{
							ZoneId: 145,
							Abbr:   "WAT",
							Offset: -3600,
						},
						{
							ZoneId: 145,
							Abbr:   "WET",
							Offset: 0,
						},
						{
							ZoneId: 145,
							Abbr:   "WEST",
							Offset: 3600,
						},
					},
				},
			},
		},
		{
			Code: "YE",
			Name: "Yemen",
			Zones: []*Zone{
				{
					Id:          420,
					CountryCode: "YE",
					Name:        "Asia/Aden",
					TimeZones: []*TimeZone{
						{
							ZoneId: 420,
							Abbr:   "LMT",
							Offset: 11212,
						},
						{
							ZoneId: 420,
							Abbr:   "AST",
							Offset: 10800,
						},
					},
				},
			},
		},
		{
			Code: "ZM",
			Name: "Zambia",
			Zones: []*Zone{
				{
					Id:          423,
					CountryCode: "ZM",
					Name:        "Africa/Lusaka",
					TimeZones: []*TimeZone{
						{
							ZoneId: 423,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 423,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "ZW",
			Name: "Zimbabwe",
			Zones: []*Zone{
				{
					Id:          424,
					CountryCode: "ZW",
					Name:        "Africa/Harare",
					TimeZones: []*TimeZone{
						{
							ZoneId: 424,
							Abbr:   "LMT",
							Offset: 7820,
						},
						{
							ZoneId: 424,
							Abbr:   "CAT",
							Offset: 7200,
						},
					},
				},
			},
		},
		{
			Code: "AX",
			Name: "Åland Islands",
			Zones: []*Zone{
				{
					Id:          46,
					CountryCode: "AX",
					Name:        "Europe/Mariehamn",
					TimeZones: []*TimeZone{
						{
							ZoneId: 46,
							Abbr:   "LMT",
							Offset: 5989,
						},
						{
							ZoneId: 46,
							Abbr:   "HMT",
							Offset: 5989,
						},
						{
							ZoneId: 46,
							Abbr:   "EET",
							Offset: 7200,
						},
						{
							ZoneId: 46,
							Abbr:   "EEST",
							Offset: 10800,
						},
					},
				},
			},
		},
	}
)

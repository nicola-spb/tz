package tz

import (
	"sync"
	"strconv"
)

var (
	once      	sync.Once
	countryMap    	map[string]*Country
	zoneMap		map[int]*Zone
	// map of timezone by country, abbr and offset
	timezoneCAOMap		map[string][]*TimeZone
	// map of timezone by abbr and offset
	timezoneAOMap		map[string][]*TimeZone
	// map of timezone by country
	timezoneCMap		map[string][]*TimeZone

)

// Country contains a country's information
type Country struct {
	Code  string
	Name  string
	Zones []*Zone
}

// Zone contains a zone`s information
type Zone struct {
	Id		int
	CountryCode 	string
	Name        	string
	TimeZones	[]*TimeZone
}

// TimeZone contains timezone`s information
type TimeZone struct {
	ZoneId		int
	Abbr 		string
	Offset        	int
}

func init() {
	// initialize maps
	once.Do(func() {
		countryMap = make(map[string]*Country)
		zoneMap = make(map[int]*Zone)
		timezoneCAOMap = make(map[string][]*TimeZone)
		timezoneAOMap = make(map[string][]*TimeZone)
		timezoneCMap = make(map[string][]*TimeZone)

		countriesLen := len(countries)
		for i := 0; i < countriesLen; i++ {
			country := &countries[i]
			countryMap[countries[i].Code] = country
			for _, z := range country.Zones {
				zoneMap[z.Id] = z
				for _, t := range z.TimeZones {
					caoKey := getTimeZoneCAOKey(country.Code, t.Abbr, t.Offset)
					timezoneCAOMap[caoKey] = append(timezoneCAOMap[caoKey], t)

					aoKey := getTimeZoneAOKey(t.Abbr, t.Offset)
					timezoneAOMap[aoKey] = append(timezoneAOMap[aoKey], t)

					timezoneAOMap[country.Code] = append(timezoneCMap[country.Code], t)
				}
			}

		}
	})
}

// GetCountries returns an array of all countries.
func GetCountries() []Country {
	return countries
}

// GetCountry get country by country code
func GetCountry(countryCode string) (c *Country, found bool) {
	c, found = countryMap[countryCode]
	return
}

// GetZonesByCountry get all zones in country
func GetZonesByCountry(countryCode string) (c []*Zone, found bool) {
	if c, ok := countryMap[countryCode]; ok {
		return c.Zones, ok
	}
	return []*Zone{}, false
}

// GetTimeZones get timezones by country, abbreviation and offset
func GetTimeZones(country string, abbr string, offset int) (timezones []*TimeZone, found bool) {
	timezones, found = timezoneCAOMap[getTimeZoneCAOKey(country, abbr, offset)]
	return
}

// GetTimeZonesByAbbrAndOffset get timezones by zone abbreviation and offset
func GetTimeZonesByAbbrAndOffset(abbr string, offset int) (timezones []*TimeZone, found bool) {
	timezones, found = timezoneAOMap[getTimeZoneAOKey(abbr, offset)]
	return
}

// GetTimeZonesByCountry get timezones by country
func GetTimeZonesByCountry(countryCode string) (timezones []*TimeZone, found bool) {
	timezones, found = timezoneCMap[countryCode]
	return
}

// GetZone get zone by zone id and whether it was found
func GetZone(zoneId int) (z *Zone, found bool) {
	z, found = zoneMap[zoneId]
	return
}

// getTimeZoneCAOKey get key of country, zone abbreviation and offset
func getTimeZoneCAOKey(country string, abbr string, offset int) string {
	return country + "-" + abbr + "-" + strconv.Itoa(offset)
}

// getTimeZoneAOKey get key of zone abbreviation and offset
func getTimeZoneAOKey(abbr string, offset int) string {
	return abbr + "-" + strconv.Itoa(offset)
}




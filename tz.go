package tz

import (
	"strconv"
	"sync"
)

var (
	once sync.Once
	// map of countries by country code
	countryMap map[string]*Country
	// map of zones by id
	zoneMap map[int]*Zone
	// map of zones by abbr and offset
	zoneAOMap map[string][]*Zone
	// map of zones by country, abbr and offset
	zoneCAOMap map[string][]*Zone
)

// Country contains a country's information
type Country struct {
	Code  string
	Name  string
	Zones []*Zone
}

// Zone contains a zone`s information
type Zone struct {
	Id          int
	CountryCode string
	Name        string
	TimeZones   []*TimeZone
}

// TimeZone contains timezone`s information.
type TimeZone struct {
	ZoneId int
	Abbr   string
	Offset int
}

func init() {
	once.Do(func() {
		countryMap = make(map[string]*Country)
		zoneMap = make(map[int]*Zone)
		zoneAOMap = make(map[string][]*Zone)
		zoneCAOMap = make(map[string][]*Zone)
		countriesLen := len(countries)
		for i := 0; i < countriesLen; i++ {
			country := &countries[i]
			countryMap[countries[i].Code] = country
			for _, z := range country.Zones {
				zoneMap[z.Id] = z
				for _, t := range z.TimeZones {
					aoKey := getZoneAOKey(t.Abbr, t.Offset)
					zoneAOMap[aoKey] = append(zoneAOMap[aoKey], z)

					caoKey := getZoneCAOKey(country.Code, t.Abbr, t.Offset)
					zoneCAOMap[caoKey] = append(zoneCAOMap[caoKey], z)
				}
			}

		}
	})
}

// GetCountries returns an array of all countries.
func GetCountries() []Country {
	return countries
}

// GetCountry get country by country code.
func GetCountry(countryCode string) (c *Country, found bool) {
	c, found = countryMap[countryCode]
	return
}

// GetZonesByCountry get all zones in country.
func GetZonesByCountry(countryCode string) ([]*Zone, bool) {
	if c, ok := countryMap[countryCode]; ok {
		return c.Zones, ok
	}
	return []*Zone{}, false
}

// GetZones get timezones by country, abbreviation and offset.
func GetZones(country string, abbr string, offset int) (zones []*Zone, found bool) {
	zones, found = zoneCAOMap[getZoneCAOKey(country, abbr, offset)]
	return
}

// GetZonesByAbbrAndOffset get timezones by zone abbreviation and offset.
func GetZonesByAbbrAndOffset(abbr string, offset int) (zones []*Zone, found bool) {
	zones, found = zoneAOMap[getZoneAOKey(abbr, offset)]
	return
}

// GetZoneById get zone by zone id and whether it was found.
func GetZoneById(zoneId int) (z *Zone, found bool) {
	z, found = zoneMap[zoneId]
	return
}

// getZoneCAOKey get key of country, zone abbreviation and offset.
func getZoneCAOKey(country string, abbr string, offset int) string {
	return country + "-" + abbr + "-" + strconv.Itoa(offset)
}

// getZoneAOKey get key of zone abbreviation and offset.
func getZoneAOKey(abbr string, offset int) string {
	return abbr + "-" + strconv.Itoa(offset)
}

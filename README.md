Package tz
==========

![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/tz/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/tz)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/tz)](https://goreportcard.com/report/github.com/go-playground/tz)
[![GoDoc](https://godoc.org/github.com/go-playground/tz?status.svg)](https://godoc.org/github.com/go-playground/tz)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package tz contains Timezone Country and Zone data generated from timezonedb.com

This library is nothing special, it contains alphabetically sorted Country, Zone and Timezone info that can be used within your project.

#### Motivation
I got timezone abbreviation and offset in browser and I wanted to find timezone by this data. This package can do it.

I found https://github.com/go-playground/tz and forked it. I hope it is

Thank you @go-playground!

##### Get timezones by country, abbreviation and offset
```go
tz.GetTimeZones(country string, abbr string, offset int)
```

##### Get timezones by abbreviation and offset
```go
tz.GetTimeZonesByAbbrAndOffset(abbr string, offset int)
```

##### Get timezones by country
```go
tz.GetTimeZonesByCountry(countryCode string)
```

##### Get all countries
```go
tz.GetCountries()
```

##### Get country by code
```go
tz.GetCountry(countryCode string)
```

##### Get all zones in country
```go
tz.GetZonesByCountry(countryCode string)
```

##### Example using with 
```go

timezones := tz.GetTimeZones("AF", "LMT", 16608)
if len(timezones) > 0 {
    loc, _ := time.LoadLocation(timezones[0])
    // now that you have location can use with Go's time package which handles timezone offsets & Daylight savings times.
    time.ParseInLocation(...)
    time.Now().In(loc)
}

```


#### NOTES
This is intended to be used along side Go's own time logic eg. `time.LoadLocation`

# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

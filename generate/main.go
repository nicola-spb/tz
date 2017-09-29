package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"text/template"
	"time"

	"tz"
	"strconv"
)

const (
	workDir     	 	= "tzdb"
	timezonedbFileName 	= "timezonedb.csv.zip"
	timezonedbURL 		= "https://timezonedb.com/files/" + timezonedbFileName
	timeZoneFileName 	= "timezone.csv"
	countryFileName  	= "country.csv"
	zoneFileName 	 	= "zone.csv"
	outputFile  	 	= "../tz_data.go"
)

type countryColumn int

// Country columns
const (
	countryCode countryColumn = iota
	countryName
)

type zoneColumn int

// Zone columns
const (
	zoneId zoneColumn = iota
	zoneCode
	zoneName
)

type timezoneColumn int

// Timezone columns
const (
	timeZoneId timezoneColumn = iota
	timeZoneAbbr
	timeZoneStart
	timeZoneOffset
)

type byCountryName []*tz.Country

func (a byCountryName) Len() int {
	return len(a)
}

func (a byCountryName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byCountryName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type byZoneName []*tz.Zone

func (a byZoneName) Len() int {
	return len(a)
}

func (a byZoneName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byZoneName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}


type byTimeZoneAbbr []*tz.TimeZone

func (a byTimeZoneAbbr) Len() int {
	return len(a)
}

func (a byTimeZoneAbbr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byTimeZoneAbbr) Less(i, j int) bool {
	return a[i].Abbr < a[j].Abbr
}

func main() {
	fmt.Print("- Parsing template...")
	tmpl, err := template.New("gen").Parse(output)
	if err != nil {
		log.Fatal("ERROR parsing template:", err)
	}
	fmt.Println(" success")

	fmt.Print("- Get current directory...")
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("ERROR determining current working DIR:", err)
	}
	fmt.Println(" success")

	fmt.Print("- Creating temp directory...")
	cmd := exec.Command("mkdir", "-p", workDir)
	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR creating working DIR:", err)
	}	

	err = os.Chdir(workDir)
	if err != nil {
		log.Fatal("ERROR switching to working DIR:", err)
	}
	fmt.Println(" success")

	fmt.Print("- Downloading timezonedb: ", timezonedbURL, "...")
	cmd = exec.Command("curl", "-O", timezonedbURL)
	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR downloading file:", err)
	}
	fmt.Println(" success")

	fmt.Print("- Unzipping archive...")
	cmd = exec.Command("unzip", timezonedbFileName)
	err = cmd.Run()
	if err != nil {
		log.Fatal("ERROR Unzipping file:", err)
	}
	fmt.Println(" success")

	fmt.Println("- Parsing files...")
	countries, err := process()
	if err != nil {
		log.Fatal("ERROR processing files:", err)
	}
	fmt.Println("  success")

	err = os.Chdir(cwd)
	if err != nil {
		log.Fatal("ERROR switching to original working DIR:", err)
	}

	fmt.Print("- Generating code to: ", outputFile, "...")
	err = os.Remove(outputFile)
	if err != nil {
		log.Fatal("ERROR failed delete file:", err)
	}

	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("ERROR writing/creating tz data file:", err)
	}

	err = tmpl.Execute(f, countries)
	if err != nil {
		log.Fatal("ERROR executing template:", err)
	}
	// close before formatting
	f.Close()
	fmt.Println(" success")

	fmt.Print("- Removing temp directory...")
	err = os.RemoveAll(workDir)
	if err != nil {
		log.Fatal("ERROR removing working DIR:", err)
	}
	fmt.Println(" success")

	fmt.Print("- Formatting code...")
	// after file written run gofmt on file
	cmd = exec.Command("gofmt", "-s", "-w", outputFile)
	if err = cmd.Run(); err != nil {
		log.Fatal("ERROR running gofmt:", err)
	}
	fmt.Println(" success")

}

func process() ([]*tz.Country, error) {
	var err error

	var countries []*tz.Country
	countryMap := make(map[string]*tz.Country)
	zoneMap := make(map[int]*tz.Zone)

	// parse countries
	if countries, err = parseCountries(countryMap, countries); err != nil {
		return nil, err
	}

	// parse zones
	if zoneMap, err = parseZones(countryMap); err != nil {
		return nil, err
	}

	// parse timezones
	if err = parseTimeZones(zoneMap); err != nil {
		return nil, err
	}

	// sort countries
	sort.Sort(byCountryName(countries))
	// sort zones
	for _, c := range countries {
		sort.Sort(byZoneName(c.Zones))
	}

	return countries, nil
}

// parseCountries parse countries file
func parseCountries(countryMap map[string]*tz.Country, countries []*tz.Country) ([]*tz.Country, error) {
	countryFile, err := os.Open(countryFileName)
	if err != nil {
		return countries, err
	}
	defer countryFile.Close()

	r := csv.NewReader(countryFile)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		country := &tz.Country{
			Code: row[countryCode],
			Name: row[countryName],
		}
		countryMap[country.Code] = country
		countries = append(countries, country)
	}

	return countries, nil
}

// parseZones parse zones file
func parseZones(countryMap map[string]*tz.Country) (map[int]*tz.Zone, error) {
	zoneFile, err := os.Open(zoneFileName)
	zoneMap := map[int]*tz.Zone{}
	if err != nil {
		return zoneMap, err
	}
	defer zoneFile.Close()

	r := csv.NewReader(zoneFile)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		zone := &tz.Zone{
			CountryCode: 	row[zoneCode],
			Name:        	row[zoneName],
		}
		zone.Id, _ = strconv.Atoi(row[zoneId])

		// test zone is working in Go
		if _, err = time.LoadLocation(zone.Name); err != nil {
			fmt.Println("  error: failed load timezone ", zone.Id, ":", zone.Name, ";", err)
			continue
		}

		if c, ok := countryMap[zone.CountryCode]; ok {
			zoneMap[zone.Id] = zone
			c.Zones = append(c.Zones, zone)
		} else {
			fmt.Println("  error: failed find country \"", zone.CountryCode, "\"")
			continue
		}
	}

	return zoneMap, nil
}

// parseTimeZones parse timezones file
func parseTimeZones(zoneMap map[int]*tz.Zone) error {
	timeZoneFile, err := os.Open(timeZoneFileName)
	timeZoneMap := map[string]bool{}
	if err != nil {
		return err
	}
	defer timeZoneFile.Close()

	r := csv.NewReader(timeZoneFile)
	failedTimeZones := make(map[int]bool)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		timeZone := &tz.TimeZone{
			Abbr: row[timeZoneAbbr],
		}
		timeZone.ZoneId, _ = strconv.Atoi(row[timeZoneId])
		timeZone.Offset, _ = strconv.Atoi(row[timeZoneOffset])

		uniqueTimeZoneKey := fmt.Sprint(timeZone.ZoneId, timeZone.Abbr, timeZone.Offset)
		if _, exists := timeZoneMap[uniqueTimeZoneKey]; !exists {
			if z, ok := zoneMap[timeZone.ZoneId]; ok {
				timeZoneMap[uniqueTimeZoneKey] = true
				z.TimeZones = append(z.TimeZones, timeZone)
			} else {
				failedTimeZones[timeZone.ZoneId] = true
				continue
			}
		}
	}
	for failedZone, _ := range failedTimeZones {
		fmt.Println("  error: timezones for \"", failedZone, "\" zone wasn`t loaded")
	}

	return nil
}

var output = `package tz

// GENERATED FILE DO NOT MODIFY DIRECTLY. GENERATED AT: ` + time.Now().String() + `

var (
	countries = []Country{
			{{ range $c := .}}{
				Code: "{{ $c.Code }}",
				Name: "{{ $c.Name }}",
				Zones: []*Zone{
					{{ range $z := $c.Zones }}{
						Id: {{ $z.Id }},
						CountryCode: "{{ $z.CountryCode }}",
						Name: "{{ $z.Name }}",
						TimeZones: []*TimeZone{
							{{ range $tz := $z.TimeZones }}{
								ZoneId: {{ $tz.ZoneId }},
								Abbr: "{{ $tz.Abbr }}",
								Offset: {{ $tz.Offset }},
							},
							{{ end }}
						},
					},
					{{ end }}
				},
			},
			{{ end }}
	}
)


`

// func main() {

// 	time.Local = time.UTC

// 	loc, err := time.LoadLocation("America/Toronto")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	utc := time.Now()

// 	fmt.Println("   NOW UTC:", utc)

// 	local := utc.In(loc)
// 	fmt.Println("LOCAL TIME:", local)

// 	edt, err := time.Parse("2006-01-02", "2016-04-01")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	est, err := time.Parse("2006-01-02", "2016-12-01")
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	fmt.Println("EDT UTC:", edt)
// 	fmt.Println("EST UTC:", est)

// 	fmt.Println("EDT LOCAL:", edt.In(loc))
// 	fmt.Println("EST LOCAL:", est.In(loc))
// }

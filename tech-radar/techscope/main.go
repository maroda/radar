package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TechList ::: Tech Radar Incoming JSON Fields
// What may be the incoming data post?
type TechList struct {
	Subject  string
	Category string
	Ring     string
	State    string
}

// RadRing ::: Tech Radar Radian/Ring Schema
// This is the default schema found in the original Tech Radar index.html
// These are rows, they are currently kept in a map.
type RadRing struct {
	Quadrant int    // Code (0), Data (1), Platform (2), Tool (3)
	Ring     int    // Adopt (0), Trial (1), Assess (2), Hold (3)
	Label    string // Tech Subject
	Active   bool   // Not sure what this is yet
	Link     string // External URL
	Moved    int    // Ring Delta
}

// readFile ::: Generic File Reader, returns a byte array of the file.
func readFile(f *string) string {
	if len(*f) == 0 {
		log.Error().Msg("no path given")
		return "ENOENT"
	}

	var fileBuf []byte
	fileBuf, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Error()
	}

	return string(fileBuf)
}

// formatEnt :::
func formatEnt(e []string) RadRing {
	// 	fmt.Printf("%v \n", e)
	var newRadarRow RadRing
	var radian, ring int

	switch e[1] {
	case "Code":
		radian = 0
	case "Data":
		radian = 1
	case "Platform":
		radian = 2
	case "Tool":
		radian = 3
	}

	switch e[2] {
	case "Adopt":
		ring = 0
	case "Trial":
		ring = 1
	case "Assess":
		ring = 2
	case "Hold":
		ring = 3
	}

	newRadarRow.Quadrant = radian
	newRadarRow.Ring = ring
	newRadarRow.Label = e[0]
	newRadarRow.Active = true
	newRadarRow.Link = "."
	newRadarRow.Moved = 0

	return newRadarRow
}

// readCSV ::: Transforms a CSV into JSON blob
func readCSV(f string) map[int][]byte {
	entryMap := make(map[int][]byte)

	readFile, err := os.Open(f)
	if err != nil {
		log.Error()
		return entryMap
	}
	defer readFile.Close()

	// new csv reader
	rcsv := csv.NewReader(readFile)

	// knock off the first line of headers
	_, err = rcsv.Read()
	if err != nil {
		log.Error()
		return entryMap
	}

	// now read the remainder of the csv file
	// and transform each row into a json blob
	for i := 0; ; i++ {
		row, err := rcsv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Error()
			return entryMap
		}

		newRow := formatEnt(row)
		newEnt, err := json.MarshalIndent(newRow, "", "    ")
		if err != nil {
			log.Error()
		}

		// add the new blob to the bag
		// there is a potential problem here that labels in the json need to be lowercase
		entryMap[i] = newEnt
	}
	return entryMap
}

// readJSON ::: Future usage for in-line automation.
func readJSON() {
	var tRad TechList
	listFile := "TRad.json"
	source, err := ioutil.ReadFile(listFile)
	if err != nil {
		log.Error()
	}

	if err := json.Unmarshal(source, &tRad); err != nil {
		log.Fatal().Err(err).Msg("failed to decode JSON")
	}

	fmt.Println(tRad.Subject)
	fmt.Println(tRad.Category)
	fmt.Println(tRad.Ring)
	fmt.Println(tRad.State)
}

func main() {
	// Zerolog
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Runtime Flags
	var list string
	flag.StringVar(&list, "list", "", "List Filename")

	flag.Parse()

	// Get the entries from the CSV file given with the -list flag,
	// (returned as a map of JSON blobs) then print for importability.
	entryMap := readCSV(list)
	for _, entry := range entryMap {
		fmt.Printf("%s,\n", entry)
	}
}

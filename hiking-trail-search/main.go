package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var tpl *template.Template
var trails []HikingTrail

type HikingTrail struct {
	AccessID    string
	AccessName  string
	Restrooms   bool
	Picnics     bool
	Fishing     bool
	Address     string
	Fee         bool
	BikeRacks   bool
	BikeTrails  bool
	Grills      bool
	TrashCans   string
	Difficulty  string
	RecycleBins bool
	DogCompost  bool
}

func init() {
	tpl = template.Must(template.ParseFiles("search.html"))

	hikingTrails, err := getTrailValues("BoulderTrailHeads.csv")
	if err != nil {
		log.Fatal("Error loading trails data")
		return
	}
	trails = hikingTrails
}

func getTrailValues(path string) ([]HikingTrail, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while reading the file "+path, err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
		return nil, err
	}

	var hikingTrails []HikingTrail

	//create a map with fields
	fields := make(map[string]int)
	for i, name := range records[0] {
		fields[name] = i
	}

	//check for any missing column
	for _, name := range []string{"AccessID", "AccessName", "RESTROOMS", "PICNIC", "FISHING", "Address",
		"Fee", "BikeRack", "BikeTrail", "Grills", "TrashCans", "ADAtrail", "RecycleBin", "DogCompost"} {
		if _, valid := fields[name]; !valid {
			log.Fatal("field missing : ", name)
		}
	}

	//create an array of all the hiking trails
	for _, record := range records[1:] {
		hikingTrails = append(hikingTrails, HikingTrail{
			AccessID:    record[fields["AccessID"]],
			AccessName:  record[fields["AccessName"]],
			Restrooms:   record[fields["RESTROOMS"]] == "Yes",
			Picnics:     record[fields["PICNIC"]] == "Yes",
			Fishing:     record[fields["FISHING"]] == "Yes",
			Address:     record[fields["Address"]],
			Fee:         record[fields["Fee"]] == "Yes",
			BikeRacks:   record[fields["BikeRack"]] == "Yes",
			BikeTrails:  record[fields["BikeTrail"]] == "Yes",
			Grills:      record[fields["Grills"]] == "Yes",
			TrashCans:   record[fields["TrashCans"]],
			Difficulty:  record[fields["ADAtrail"]],
			RecycleBins: record[fields["RecycleBin"]] == "Yes",
			DogCompost:  record[fields["DogCompost"]] == "Yes",
		})
	}

	return hikingTrails, nil

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
		fmt.Println(fmt.Sprint(err) + ": ")
		return
	}

	// get user input from the form
	restrooms := r.FormValue("restrooms") == "on"
	picnics := r.FormValue("picnics") == "on"
	fishing := r.FormValue("fishing") == "on"
	fee := r.FormValue("fee") == "on"
	bikeRacks := r.FormValue("bike_racks") == "on"
	bikeTrails := r.FormValue("bike_trails") == "on"
	grills := r.FormValue("grills") == "on"
	trashCans := r.FormValue("trash_cans")
	difficulty := r.FormValue("difficulty")
	recycleBins := r.FormValue("recycle_bins") == "on"
	dogCompost := r.FormValue("dog_compost") == "on"

	filteredTrails := filterTrails(trails, restrooms, picnics, fishing, fee, bikeRacks, bikeTrails, grills, recycleBins, dogCompost, trashCans, difficulty)

	tpl, err := template.ParseFiles("search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, filteredTrails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

}

func filterTrails(trails []HikingTrail, restrooms, picnics, fishing, fee, bikeRacks, bikeTrails, grills, recycleBins, dogCompost bool, trashCans, difficulty string) []HikingTrail {
	var filtered []HikingTrail
	for _, trail := range trails {

		// filter with checkboxes
		if (!restrooms || trail.Restrooms) &&
			(!picnics || trail.Picnics) &&
			(!fishing || trail.Fishing) &&
			(!fee || trail.Fee) &&
			(!bikeRacks || trail.BikeRacks) &&
			(!bikeTrails || trail.BikeTrails) &&
			(!grills || trail.Grills) &&
			(!recycleBins || trail.RecycleBins) &&
			(!dogCompost || trail.DogCompost) &&
			(trashCans == "0" || filterTrashCans(trashCans, trail.TrashCans)) &&
			(difficulty == "default" || filterDifficulty(difficulty, trail.Difficulty)) {
			filtered = append(filtered, trail)
		}
	}
	return filtered
}

func filterTrashCans(filterValue string, trailValue string) bool {

	if filterValue != "" {
		value := string(filterValue[0])

		// string to int
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		trail, err := strconv.Atoi(trailValue)

		if trail >= i {
			return true
		}
	}
	return false
}

func filterDifficulty(filterValue string, trailValue string) bool {

	if filterValue != "" {
		if strings.Contains(trailValue, filterValue) {
			return true
		}
	}
	return false
}

package main

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/deckarep/gosx-notifier"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Not used
type StoreStock struct {
	CG iPhone6p `json:"R245"`
	RS iPhone6p `json:"R092"`
	SC iPhone6p `json:"R410"`
	WC iPhone6p `json:"R226"`
}

// iPhone 6 plus
type iPhone6p struct {
	Stock bool `json:"MGAH2B/A"`
}

//a slice of string sites that you are interested in watching
var StoreCodes map[string]string = map[string]string{
	"R245": "Covent Garden",
	"R092": "Regent Street",
	"R410": "Stratford City",
	"R226": "White City",
}

func main() {
	ch := make(chan string)
	go pinger(ch, "https://reserve.cdn-apple.com/GB/en_GB/reserve/iPhone/availability.json")

	for {
		select {
		case result := <-ch:
			if strings.HasPrefix(result, "-") {
				s := strings.Trim(result, "-")
				showNotification(s)
			} else {
				showNotification("Store has stock!!! - " + result)
			}
		}
	}
}

func showNotification(message string) {

	note := gosxnotifier.NewNotification(message)
	note.Title = "iPhone6Plus Stock"
	note.Sound = gosxnotifier.Default

	note.Push()
}

//Prefixing a site with a + means it's up, while - means it's down
func pinger(ch chan string, site string) {
	for {
		resp, err := http.Get(site)
		defer resp.Body.Close()
		if err != nil {
			ch <- "- Cant get website"
		}

		if resp.StatusCode != 200 {
			ch <- "- 200 Fail"
		} else {

			// Loop through look for stock
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ch <- "- Cannot readall body"
			}

			ss := StoreStock{}
			err = json.Unmarshal(body, &ss)
			if err != nil {
				log.Errorf("Error: %v", err)
				ch <- "- Cannot unmarshall"
			}

			log.Debugf("Stock for CG is %v", ss.CG.Stock)
			if ss.CG.Stock == true {
				ch <- "Covent Garden"
			}

			log.Debugf("Stock for RS is %v", ss.RS.Stock)
			if ss.CG.Stock == true {
				ch <- "Regrent Street"
			}

			log.Debugf("Stock for SC is %v", ss.SC.Stock)
			if ss.SC.Stock == true {
				ch <- "Straford City"
			}

			log.Debugf("Stock for WC is %v", ss.WC.Stock)
			if ss.WC.Stock == true {
				ch <- "White City"
			}

			// Loop and return
			// for storeCode, stockStatus := range ss {
			// 	log.Debugf("Stock: %v", stockStatus)
			// 	if stockStatus.Stock == true {
			//		ch <- "storeCode"
			// 	}
			// }
		}

		time.Sleep(30 * time.Second)
	}
}

// array(
// 	"R227" => "Bentall Centre",
// 	"R113" => "Bluewater",
// 	"R340" => "Braehead",
// 	"R163" => "Brent Cross",
// 	"R496" => "Bromley",
// 	"R135" => "Buchanan Street",
// 	"R118" => "Bullring",
// 	"R252" => "Cabot Circus",
// 	"R391" => "Chapelfield",
// 	"R244" => "Churchill Square",
// 	"R245" => "Covent Garden",
// 	"R393" => "Cribbs Causeway",
// 	"R545" => "Drake Circus",
// 	"R341" => "Eldon Square",
// 	"R482" => "Festival Place",
// 	"R270" => "Grand Arcade",
// 	"R308" => "Highcross",
// 	"R242" => "Lakeside",
// 	"R239" => "Liverpool ONE",
// 	"R215" => "Manchester Arndale",
// 	"R153" => "Meadowhall",
// 	"R423" => "Metrocentre",
// 	"R269" => "Milton Keynes",
// 	"R279" => "Princesshay",
// 	"R092" => "Regent Street",
// 	"R335" => "SouthGate",
// 	"R334" => "St David's 2",
// 	"R410" => "Stratford City",
// 	"R176" => "The Oracle",
// 	"R255" => "Touchwood Centre",
// 	"R136" => "Trafford Centre",
// 	"R372" => "Trinity Leeds",
// 	"R363" => "Union Square",
// 	"R313" => "Victoria Square",
// 	"R527" => "Watford",
// 	"R174" => "WestQuay",
// 	"R226" => "White City"
// );

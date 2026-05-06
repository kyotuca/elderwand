/*
Package api manages data retrieval functions
*/
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

const url = "https://tennis-api-atp-wta-itf.p.rapidapi.com"

type TennisData struct {
	nextMatch time.Time
	apiKey    string
}

func NewTennisData(currentChampionID int) *TennisData {
	return &TennisData{time.Now(), os.Getenv("TENNISAPI_KEY")}
}

func (p *TennisData) prepareRequest(endpoint string) *http.Request {
	req, err := http.NewRequest("GET", url+endpoint, nil)
	if err != nil {
		log.Fatal("Error preparing request:\n", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-rapidapi-host", "tennis-api-atp-wta-itf.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", p.apiKey)
	return req
}

func (p *TennisData) logRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal("Log dump failed:\n", err)
	}
	log.Printf("\n --- REQUEST --- \n %s \n---------------------\n", string(dump))
}

type Player struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CountryAcr string `json:"countryAcr"`
}

type Match struct {
	ID       int        `json:"id"`
	Date     *time.Time `json:"date"`
	Live     *string    `json:"live"`
	Complete *int       `json:"complete"`
	Draw     int        `json:"Draw"`
	Player1  Player     `json:"player1"`
	Player2  Player     `json:"player2"`
}

type Tournament struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *TennisData) retrieveNextMatch(playerID int) *Match {
	client := &http.Client{}
	endPoint := fmt.Sprintf("/tennis/v2/atp/fixtures/player/%d", playerID)

	req := p.prepareRequest(endPoint)
	p.logRequest(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Request failed", err)
	}
	var data []Match
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Decoding error", err)
	}

	if len(data) == 0 {
		return nil
	}

	return &data[0]
}

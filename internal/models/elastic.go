package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Elastic struct {
	Domain string
}

type QueryResponse struct {
	Hits `json:"hits"`
}

type Hits struct {
	Total    uint    `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Source RoommateOffer `json:"_source"`
}

func NewElasticClient() *Elastic {
	return &Elastic{
		Domain: "https://search-capstone18es-527sqz4yrcalm26l6vlodv33nm.us-east-1.es.amazonaws.com",
	}
}

func (e *Elastic) Put(offer *RoommateOffer) error {
	body, err := json.Marshal(offer)
	if err != nil {
		return err
	}
	url := e.Domain + "/offers/offer/" + strconv.Itoa(int(offer.ID))
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	r, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func getJson(url string, target interface{}) error {
	r, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func (e *Elastic) Get(offer *MatchRequest, results *QueryResponse) error {
	query := "/offers/offer/_search?sort=CreatedAt:desc&q="
	if offer.GenderRequirement != "" {
		query += "genderRequirement:" + offer.GenderRequirement
	}
	if offer.Bedrooms != 0 {
		query += " AND bedrooms:" + strconv.Itoa(int(offer.Bedrooms))
	}
	if offer.Bathrooms != 0 {
		query += " AND bathrooms:" + strconv.Itoa(int(offer.Bathrooms))
	}
	if offer.State != "" {
		query += " AND state:" + offer.State
	}
	if offer.City != "" {
		query += " AND city:" + offer.City
	}
	if offer.Zip != 0 {
		query += " AND zip:" + strconv.Itoa(int(offer.Zip))
	}
	if offer.PropertyType != "" {
		query += " AND propertyType:" + offer.PropertyType
	}
	if offer.TargetResidentCount != 0 {
		query += " AND targetResidentCount:" + strconv.Itoa(int(offer.TargetResidentCount))
	}
	query += fmt.Sprintf(" AND preChosenProperty:%v AND petsAllowed:%v AND smokingAllowed:%v",
		offer.PreChosenProperty, offer.PetsAllowed, offer.SmokingAllowed)
	query = strings.Replace(query, " ", "%20", -1)
	url := e.Domain + query
	fmt.Println(url)
	return getJson(url, results)
}

func (e *Elastic) Delete(offerID string) error {
	url := e.Domain + "/offers/offer/" + offerID
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	r, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Elastic struct {
	Domain string
}

type QueryResponse struct {
	Hits `json:"hits"`
}

type Hits struct {
	Total    uint            `json:"total"`
	MaxScore uint            `json:"max_score"`
	Hits     []RoommateOffer `json:"hits"`
}

func NewElasticClient() *Elastic {
	return &Elastic{
		Domain: "https://search-capstone18es-527sqz4yrcalm26l6vlodv33nm.us-east-1.es.amazonaws.com",
	}
}

func (e *Elastic) Put(offer *RoommateOffer) (*http.Response, error) {
	body, err := json.Marshal(offer)
	if err != nil {
		return nil, err
	}
	url := e.Domain + "/offers/offer/" + strconv.Itoa(int(offer.ID))
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return (&http.Client{Timeout: 10 * time.Second}).Do(req)
}

func (e *Elastic) Get(offer *RoommateOffer) (*QueryResponse, error) {
	query := fmt.Sprintf(`/_search?q="genderRequirement:%v AND preChosenProperty:%v
			AND propertyType:%v AND zip:%v AND petsAllowed:%v AND bathrooms:%v AND bedrooms:%v
			AND smokingAllowed:%v AND targetResidentCount:%v"`, offer.GenderRequirement,
		offer.PreChosenProperty, offer.PropertyType, offer.Zip, offer.PetsAllowed,
		offer.Bathrooms, offer.Bedrooms, offer.SmokingAllowed, offer.TargetResidentCount)
	url := e.Domain + query
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	queryRes := QueryResponse{}
	if err := json.NewDecoder(res.Body).Decode(&queryRes); err != nil {
		return nil, err
	}
	return &queryRes, nil
}

func (e *Elastic) Delete(offerID string) (*http.Response, error) {
	url := e.Domain + "/offers/offer/" + offerID
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return (&http.Client{Timeout: 10 * time.Second}).Do(req)
}

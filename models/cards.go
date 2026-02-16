package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Card struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Supertype string   `json:"supertype"`
	Subtypes  []string `json:"subtypes"`
	HP        string   `json:"hp"`
	Types     []string `json:"types"`

	Attacks []struct {
		Name   string `json:"name"`
		Damage string `json:"damage"`
		Text   string `json:"text"`
	} `json:"attacks"`

	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`

	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances"`

	Set struct {
		Name   string `json:"name"`
		Series string `json:"series"`
		Images struct {
			Logo   string `json:"logo"`
			Symbol string `json:"symbol"`
		} `json:"images"`
	} `json:"set"`

	Rarity string `json:"rarity"`

	Images struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
}

type APIResponse struct {
	Data []Card `json:"data"`
}

func GetCards(page int, pageSize int) ([]Card, error) {
	url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?page=%d&pageSize=%d", page, pageSize)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	return apiResp.Data, nil
}

func GetCardByID(id string) (Card, error) {
	url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return Card{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Card{}, err
	}

	var apiResp struct {
		Data Card `json:"data"`
	}

	json.Unmarshal(body, &apiResp)

	return apiResp.Data, nil
}

func GetFilteredCards(page int, size int, typeFilter, hpFilter, supertypeFilter string) ([]Card, error) {

	query := ""

	if typeFilter != "" {
		query += fmt.Sprintf("types:%s ", typeFilter)
	}

	if hpFilter != "" {
		query += fmt.Sprintf("hp:[%s TO 500] ", hpFilter)
	}

	if supertypeFilter != "" {
		query += fmt.Sprintf("supertype:%s ", supertypeFilter)
	}

	url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?page=%d&pageSize=%d&q=%s",
		page, size, url.QueryEscape(query))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	return apiResp.Data, nil
}
func SearchCards(query string) ([]Card, error) {
	url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?q=name:%s", url.QueryEscape(query))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	return apiResp.Data, nil
}

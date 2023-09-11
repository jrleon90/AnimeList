package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type animeResponse struct {
	Data struct {
		Title         string `json:"title"`
		TitleEnglish  string `json:"title_english"`
		TitleJapanese string `json:"title_japanese"`
	}
}

func GetAnimeFact(animeId string) (animeResponse, error) {
	url := fmt.Sprintf("https://api.jikan.moe/v4/anime/%s/full", animeId)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return animeResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return animeResponse{}, err
	}
	object := animeResponse{}
	json.Unmarshal(body, &object)
	return object, nil
}

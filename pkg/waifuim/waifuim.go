package waifuim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/browser"
)

type Artist struct {
	ArtistID   int     `json:"artist_id"`
	Name       string  `json:"name"`
	Patreon    *string `json:"patreon"`
	Pixiv      string  `json:"pixiv"`
	Twitter    string  `json:"twitter"`
	DeviantArt *string `json:"deviant_art"`
}

type Tag struct {
	TagID       int    `json:"tag_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsNSFW      bool   `json:"is_nsfw"`
}

type Image struct {
	Signature     string  `json:"signature"`
	Extension     string  `json:"extension"`
	ImageID       int     `json:"image_id"`
	Favorites     int     `json:"favorites"`
	DominantColor string  `json:"dominant_color"`
	Source        string  `json:"source"`
	Artist        Artist  `json:"artist"`
	UploadedAt    string  `json:"uploaded_at"`
	LikedAt       *string `json:"liked_at"`
	IsNSFW        bool    `json:"is_nsfw"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	ByteSize      int     `json:"byte_size"`
	URL           string  `json:"url"`
	PreviewURL    string  `json:"preview_url"`
	Tags          []Tag   `json:"tags"`
}

type ApiResponse struct {
	Images []Image `json:"images"`
}

func GetWaifu() {
	apiUrl := "https://api.waifu.im/search"

	// Creating URL parameters
	params := url.Values{}
	params.Add("included_tags", "waifu")
	queryParams := params.Encode()

	// Complete URL with query parameters
	requestUrl := fmt.Sprintf("%s?%s", apiUrl, queryParams)

	// Performing the GET request
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var apiResponse ApiResponse
		err = json.Unmarshal(bodyBytes, &apiResponse)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		fmt.Println("Finding a random waifu to cheer you up", apiResponse.Images[0].URL)

		// Open the image URL in the default web browser
		err = browser.OpenURL(apiResponse.Images[0].URL)
		if err != nil {
			fmt.Println("Error opening browser:", err)
		}
	} else {
		fmt.Println("Request failed with status code:", resp.StatusCode)
	}
}

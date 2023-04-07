package send_request

import (
	"aquilon/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func SendRequest() []models.ApiResponse {
	var response []models.ApiResponse

	params := url.Values{}
	params.Add("dt_start", os.Getenv("DT_START"))
	params.Add("dt_end", os.Getenv("DT_END"))
	url := fmt.Sprintf("%s?%s", os.Getenv("API_ENDPOINT"), params.Encode())

	header := http.Header{}
	token := fmt.Sprintf("Bearer %s", os.Getenv("API_AUTH_TOKEN"))
	header.Add("Authorization", token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = header

	if !checkUseRest() {
		f, _ := os.ReadFile("files/testRestData.json")

		jsonString := strings.ReplaceAll(string(f), "\n", "")
		jsonString = strings.ReplaceAll(jsonString, "\r", "")
		jsonString = strings.ReplaceAll(jsonString, "\t", "")

		err := json.Unmarshal([]byte(jsonString), &response)
		if err != nil {
			panic(err)
		}

		return response
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	return response
}

func checkUseRest() bool {
	useRest, err := strconv.ParseBool(os.Getenv("USE_REST"))
	if err != nil {
		return false
	}
	return useRest
}

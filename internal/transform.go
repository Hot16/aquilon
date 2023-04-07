package internal

import (
	"aquilon/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TransformClients(apiClients []models.ApiResponse) []models.Clients {
	var clients []models.Clients
	for _, apiClient := range apiClients {
		client := transform(apiClient)
		clients = append(clients, client)
	}
	return clients
}

func transform(apiResponse models.ApiResponse) models.Clients {
	dt, err := time.Parse("2006-01-02 15:04:05", apiResponse.Dt)
	if err != nil {
		panic(err)
	}

	clientId, err := strconv.ParseUint(apiResponse.ClientId, 10, 64)
	if err != nil {
		panic(err)
	}

	return models.Clients{
		Id:           replaceF(apiResponse.Id),
		Dt:           uint64(dt.Unix()),
		ClientId:     clientId,
		Type:         replaceF(apiResponse.Type),
		SubmitId:     apiResponse.SubmitId,
		Referer:      replaceF(apiResponse.Referer),
		Os:           replaceF(apiResponse.Os),
		LeadSource:   replaceF(apiResponse.LeadSource),
		CreativeName: replaceF(apiResponse.CreativeName),
		Country:      replaceF(apiResponse.Country),
	}
}

func replaceF(s string) string {
	s = strings.ToLower(strings.ReplaceAll(s, "/", ""))
	s = strings.ReplaceAll(s, "\\", "")
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	return s
}

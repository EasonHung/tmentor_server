package http_utils

import (
	"fmt"
	"io"
	"mentor_app/chatroom/middleware/log"
	"net/http"
)

func GetString(url string, queryMap map[string]string) (error, string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}

	query := req.URL.Query()
	for key, value := range queryMap {
		query.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Logger.Error("[DB] error occured when insert message", err)
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		bodyString = bodyString[1 : len(bodyString)-1]
		return nil, bodyString
	}

	return nil, ""
}

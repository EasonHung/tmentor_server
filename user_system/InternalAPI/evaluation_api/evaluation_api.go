package evaluation_api

import (
	"bytes"
	"encoding/json"
	"mentor_app/user_system/initialize"
	"net/http"

	"github.com/pkg/errors"
)

func InitEvaluation(traceId string, userId string) error {
	client := &http.Client{}
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)
	if err != nil {
		err = errors.Wrap(err, "error call init evaluation api")
		return err
	}

	url := initialize.GLOBAL_CONFIG.Servers.EvaluationSystem + "/user/new"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		err = errors.Wrap(err, "error call init evaluation api")
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", traceId)
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error call init evaluation api")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.Errorf("error call init evaluation api")
	}

	return nil
}

// func GetScoreAndCount(traceId string, userId string) (error, GetScoreAndCountRes) {
// 	client := &http.Client{}
// 	resData := GetScoreAndCountRes{}
// 	url := initialize.GLOBAL_CONFIG.Servers.EvaluationSystem + "/user/scoreAndCount?userId=" + userId
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		err = errors.Wrap(err, "error call get evaluation score and count")
// 		return err, GetScoreAndCountRes{}
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("request-id", traceId)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		err = errors.Wrap(err, "error call init evaluation api")
// 		return err, GetScoreAndCountRes{}
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return errors.Errorf("error call init evaluation api"), GetScoreAndCountRes{}
// 	}

// 	err = json.NewDecoder(resp.Body).Decode(&resData)
// 	if err != nil {
// 		err = errors.Wrap(err, "error call verify token api")
// 		return err, GetScoreAndCountRes{}
// 	}

// 	return nil, resData
// }

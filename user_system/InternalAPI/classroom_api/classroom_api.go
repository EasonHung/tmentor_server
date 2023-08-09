package classroom_api

import (
	"bytes"
	"encoding/json"
	"mentor_app/user_system/initialize"
	"net/http"

	"github.com/pkg/errors"
)

func InitClassroom(traceId string, userId string) error {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	url := initialize.GLOBAL_CONFIG.Servers.Classroom + "/info/init"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		err = errors.Wrap(err, "error call init classroom api")
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", traceId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error call init evaluation api")
		return err
	}

	if resp.StatusCode != 200 {
		return errors.Errorf("error call init classroom api")
	}
	defer resp.Body.Close()

	return nil
}

func GetStudentCount(traceId string, userId string) (error, GetStudentCountRes) {
	client := &http.Client{}
	resData := GetStudentCountRes{}
	url := initialize.GLOBAL_CONFIG.Servers.Classroom + "/info/studentCount?userId=" + userId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = errors.Wrap(err, "error call get evaluation score and count")
		return err, GetStudentCountRes{}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", traceId)
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "error call init evaluation api")
		return err, GetStudentCountRes{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.Errorf("error call init evaluation api"), GetStudentCountRes{}
	}

	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		err = errors.Wrap(err, "error call verify token api")
		return err, GetStudentCountRes{}
	}

	return nil, resData
}

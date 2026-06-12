package api

import (
	"bytes"
	"client/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type FilApi struct {
	baseURL string
}

func InitFilApi(baseURL string) *FilApi {
	return &FilApi{baseURL: baseURL}
}

func (api *FilApi) executeRequest(req *http.Request, result interface{}) (int, error) {
	_client := http.Client{
		Timeout: time.Second * 5,
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := _client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return resp.StatusCode, readErr
	}

	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("Erreur response - %s", string(bytes.TrimSpace(body)))
	}

	if result != nil {
		if err = json.Unmarshal(body, result); err != nil {
			return resp.StatusCode, fmt.Errorf("Erreur decode données - %s", err.Error())
		}
	}

	return resp.StatusCode, nil
}

func (api *FilApi) ReadAll() ([]dto.FilDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils", nil)
	if err != nil {
		return nil, err
	}

	var list []dto.FilDto
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (api *FilApi) ReadByIdMessages(id int) ([]dto.MessageDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fil/"+strconv.Itoa(id)+"/messages", nil)
	if err != nil {
		return nil, err
	}

	var messages []dto.MessageDto

	status, err := api.executeRequest(req, &messages)
	if err != nil {
		if status == http.StatusNotFound {
			return []dto.MessageDto{}, nil
		}
		return nil, err
	}
	fmt.Println(messages)
	return messages, nil
}

func (api *FilApi) FilsPetanque() ([]dto.FilDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils/petanque", nil)
	if err != nil {
		return nil, err
	}

	var list []dto.FilDto
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}
	fmt.Println(list)
	return list, nil
}

func (api *FilApi) FilsCuisine() ([]dto.FilDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils/cuisine", nil)
	if err != nil {
		return nil, err
	}

	var list []dto.FilDto
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}
	fmt.Println(list)
	return list, nil
}

func (api *FilApi) FilsNature() ([]dto.FilDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fils/nature", nil)
	if err != nil {
		return nil, err
	}

	var list []dto.FilDto
	_, err = api.executeRequest(req, &list)
	if err != nil {
		return nil, err
	}
	fmt.Println(list)
	return list, nil
}

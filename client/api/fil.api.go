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

func (api *FilApi) Create(fil dto.FilDto, token string) (int, dto.FilDto, error) {
	payload, err := json.Marshal(fil)
	if err != nil {
		return 0, dto.FilDto{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/fils", bytes.NewReader(payload))
	if err != nil {
		return 0, dto.FilDto{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	var created dto.FilDto
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, dto.FilDto{}, err
	}

	return created.Id, created, nil
}

func (api *FilApi) CreateMessage(message dto.MessageDto, filId int, token string) (int, dto.MessageDto, error) {
	payload, err := json.Marshal(message)
	if err != nil {
		return 0, dto.MessageDto{}, err
	}

	url := fmt.Sprintf("%s/fil/%d/messages", api.baseURL, filId)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return 0, dto.MessageDto{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	var created dto.MessageDto
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, dto.MessageDto{}, err
	}

	return created.Id, created, nil
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

func (api *FilApi) ReadById(id int) (dto.FilDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fil/"+strconv.Itoa(id), nil)
	if err != nil {
		return dto.FilDto{}, err
	}

	var fil dto.FilDto
	status, err := api.executeRequest(req, &fil)
	if err != nil {
		if status == http.StatusNotFound {
			return dto.FilDto{}, nil
		}
		return dto.FilDto{}, err
	}

	return fil, nil
}

func (api *FilApi) ReadByIdMessages(id int, tri string) ([]dto.MessageDto, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseURL+"/fil/"+strconv.Itoa(id)+"/messages?tri="+tri, nil)
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

func (api *FilApi) AjoutReaction(reaction dto.Reaction, token string) (int, dto.Reaction, error) {

	payload, err := json.Marshal(reaction)
	if err != nil {
		return -1, dto.Reaction{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/message/reaction", bytes.NewReader(payload))
	if err != nil {
		return 0, dto.Reaction{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	var created dto.Reaction
	_, err = api.executeRequest(req, &created)
	if err != nil {
		return 0, dto.Reaction{}, err
	}

	return 1, created, nil
}

func (api *FilApi) UpdateFilById(fil dto.FilDto, token string) error {
	payload, err := json.Marshal(fil)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, api.baseURL+"/fil/"+strconv.Itoa(fil.Id), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("Fil introuvable")
		}
		return err
	}

	return nil
}

func (api *FilApi) DeleteFilById(id int, token string) error {
	req, err := http.NewRequest(http.MethodDelete, api.baseURL+"/fil/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusForbidden {
			return fmt.Errorf("Vous n'avez pas les droits pour supprimer ce fil")
		}
		return err
	}

	return nil
}

func (api *FilApi) UpdateMessageById(message dto.MessageDto, token string) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, api.baseURL+"/message/update", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	status, err := api.executeRequest(req, nil)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("message introuvable")
		}
		return err
	}

	return nil
}

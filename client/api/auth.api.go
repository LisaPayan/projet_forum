package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"client/dto"
)

// AuthApi contient l'URL de base utilisee pour les appels d'authentification.
type AuthApi struct {
	baseURL string
}

// InitAuthApi initialise le client API auth avec l'URL de base fournie.
func InitAuthApi(baseURL string) *AuthApi {
	return &AuthApi{baseURL: strings.TrimRight(baseURL, "/")}
}

// Login envoie les identifiants a la route POST /login de l'API.
func (api *AuthApi) Login(data dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/login", bytes.NewReader(payload))
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return dto.LoginResponseDto{}, fmt.Errorf("Connexion refusee - code %v", resp.StatusCode)
	}

	var result dto.LoginResponseDto
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("Erreur decode token - %s", err.Error())
	}

	return result, nil
}

// Me appelle la route protegee GET /me avec le token JWT de l'utilisateur.
// La reponse JSON est decodee directement dans MeResponseDto pour eviter de manipuler une string brute.
func (api *AuthApi) Me(token string) (dto.MeResponseDto, error) {
	client := http.Client{Timeout: time.Second * 5}

	// Creation de la requete HTTP vers l'API distante.
	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/me", api.baseURL), nil)
	if requestErr != nil {
		log.Printf("Erreur request (auth.me) : %v", requestErr)
		return dto.MeResponseDto{}, fmt.Errorf("Erreur request (auth.me) : %v", requestErr)
	}
	// Le token est envoye dans le header Authorization au format Bearer.
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Printf("Erreur envoie (auth.me) : %v", responseErr)
		return dto.MeResponseDto{}, fmt.Errorf("Erreur envois (auth.me) : %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Erreur result (auth.me) - Code :%v", response.StatusCode)
		return dto.MeResponseDto{}, fmt.Errorf("Erreur result (auth.me) - Code : %v", response.StatusCode)
	}

	// Decode le body JSON dans le DTO :
	// {"code":200,"message":"Hello user 1 with role admin"}
	var data dto.MeResponseDto
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return dto.MeResponseDto{}, fmt.Errorf("Erreur decode me - %s", err.Error())
	}

	return data, nil
}

func (api *AuthApi) Register(data dto.RegisterRequestDto) (dto.RegisterResponseDto, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+"/register", bytes.NewReader(payload))
	if err != nil {
		return dto.RegisterResponseDto{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return dto.RegisterResponseDto{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errorResult map[string]string
		json.NewDecoder(resp.Body).Decode(&errorResult)
		if msg, ok := errorResult["error"]; ok {
			return dto.RegisterResponseDto{}, fmt.Errorf("%s", msg)
		}
		return dto.RegisterResponseDto{}, fmt.Errorf("erreur d'inscription - code %v", resp.StatusCode)
	}

	var result dto.RegisterResponseDto
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return dto.RegisterResponseDto{}, fmt.Errorf("erreur décodage réponse inscription - %s", err.Error())
	}

	return result, nil
}

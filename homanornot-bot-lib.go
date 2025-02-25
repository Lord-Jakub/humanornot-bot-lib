package humanornotbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Payloads for requests
type ChatCreationPayload struct {
	UserID string `json:"user_id"`
	Origin string `json:"origin"`
}

type SendMessagePayload struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

type WaitMessagePayload struct {
	UserID string `json:"user_id"`
}

type GuessPayload struct {
	UserID      string `json:"user_id"`
	PartnerType string `json:"partner_type"`
}

// Structure for response
// Important are mostly ChatID and array of messages
type ChatResponse struct {
	ChatID              string    `json:"chat_id"`
	ChatTime            int       `json:"chat_time"`
	UserID              string    `json:"user_id"`
	CreatedAt           int64     `json:"created_at"`
	NumParticipants     int       `json:"num_participants"`
	PartnerID           string    `json:"partner_id"`
	PartnerGroup        string    `json:"partner_group"`
	IsMyTurn            bool      `json:"is_my_turn"`
	TurnTime            int       `json:"turn_time"`
	IsActive            bool      `json:"is_active"`
	ChatCounter         int       `json:"chat_counter,omitempty"`
	FinishedChatCounter int       `json:"finished_chat_counter,omitempty"`
	SpotOnGuessCounter  int       `json:"spot_on_guess_counter,omitempty"`
	PartnerType         string    `json:"partner_type,omitempty"`
	Messages            []Message `json:"messages"`
}

type Message struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	CreatedAt int64  `json:"created_at"`
	Text      string `json:"text"`
}

// Function for creating chat (POST)
func CreateChat() (*ChatResponse, error) {
	url := "https://api.humanornot.ai/human-or-not/chat/"

	payload := ChatCreationPayload{
		UserID: "58278235-eb96-401e-bce9-7d63fc159635",
		Origin: "honLandPage",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error during payload serialization: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error during request creation: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("hon-client-version", "0.0.5")
	req.Header.Set("origin", "https://app.humanornot.ai")
	req.Header.Set("referer", "https://app.humanornot.ai/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error during sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error during reading of response: %v", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf("Error during parsing json response %v", err)
	}
	return &chatResp, nil
}

// Function for sending message (PUT)
func SendMessage(chatID, messageText string) (*ChatResponse, error) {
	url := fmt.Sprintf("https://api.humanornot.ai/human-or-not/chat/%s/send-message", chatID)

	payload := SendMessagePayload{
		UserID: "58278235-eb96-401e-bce9-7d63fc159635",
		Text:   messageText,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error during serialization of payload %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("hon-client-version", "0.0.5")
	req.Header.Set("origin", "https://app.humanornot.ai")
	req.Header.Set("referer", "https://app.humanornot.ai/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error during sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error during reading of response: %v", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		fmt.Printf("Error during parsing of json response: %v\n", err)
		//return nil, fmt.Errorf("Error during parsing of json response: %v", err)
	}
	return &chatResp, nil
}

// Function for waiting for message (PUT)
func WaitMessage(chatID string) (*ChatResponse, error) {
	url := fmt.Sprintf("https://api.humanornot.ai/human-or-not/chat/%s/wait-message", chatID)

	payload := WaitMessagePayload{
		UserID: "58278235-eb96-401e-bce9-7d63fc159635",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error during serialization payload: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("hon-client-version", "0.0.5")
	req.Header.Set("origin", "https://app.humanornot.ai")
	req.Header.Set("referer", "https://app.humanornot.ai/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error during sending of message: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error during reading of response: %v", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf("Error during parasing of json response: %v", err)
	}
	return &chatResp, nil
}

// Sending guess
func GuessChat(chatID, partner_type string) (*ChatResponse, error) {
	url := fmt.Sprintf("https://api.humanornot.ai/human-or-not/chat/%s/guess", chatID)

	payload := GuessPayload{
		UserID:      "58278235-eb96-401e-bce9-7d63fc159635",
		PartnerType: partner_type,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error during payload serialization: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("hon-client-version", "0.0.5")
	req.Header.Set("origin", "https://app.humanornot.ai")
	req.Header.Set("referer", "https://app.humanornot.ai/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error during reading response: %v", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf("Error during parrasing of json response %v", err)
	}
	return &chatResp, nil
}

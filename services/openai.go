package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OpenAIResponse represents a generic OpenAI API response
// You can expand this struct as needed

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// CallOpenAI sends the prompt and data to OpenAI and returns the response
func CallOpenAI(apiKey, prompt string) (string, error) {
	url := "https://api.openai.com/v1/completions"
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model":      "gpt-3.5-turbo-instruct",
		"prompt":     prompt,
		"max_tokens": 512,
	})
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API error: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var openaiResp OpenAIResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", err
	}
	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}
	return openaiResp.Choices[0].Text, nil
}

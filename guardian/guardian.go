package guardian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	apiUrl = "https://api.peeace.net/v1/guardians/text"
)

type GuardianInput struct {
	Text           string
	APIKey         string
	ScoreThreshold float64
}

func RequestGuardian(in GuardianInput) (map[string]interface{}, error) {
	if in.ScoreThreshold < 0.0 || 1.0 < in.ScoreThreshold {
		return nil, fmt.Errorf("ScoreThreshold is out of range")
	}
	if in.APIKey == "" {
		return nil, fmt.Errorf("Error API key is required.")
	}

	// リクエストを作成
	body, err := json.Marshal(map[string]string{
		"text": string(in.Text),
	})
	if err != nil {
		return nil, fmt.Errorf("Error marshalling request body: %q", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %q", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", in.APIKey))

	// リクエスト
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %q", err)
	}
	defer resp.Body.Close()

	// レスポンスをチェック
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status: %q", resp.Status)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("Error decoding response body: %q", err)
	}

	return responseBody, nil
}

package guardian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiUrl                = "https://api.peeace.net/v1/guardians/text"
	defaultScoreThreshold = 0.5
)

func validateInput(in GuardianInput) error {
	if in.ScoreThreshold < 0.0 || 1.0 < in.ScoreThreshold {
		return ErrScoreThresholdOutOfRange
	}

	if in.APIKey == "" {
		return ErrAPIKeyRequired
	}

	return nil
}

func RequestGuardian(in GuardianInput) (out GuardianOutput, err error) {
	if err := validateInput(in); err != nil {
		return out, fmt.Errorf("Validation Error: %q ", err.Error())
	}
	// user not set
	if in.ScoreThreshold == 0 {
		in.ScoreThreshold = defaultScoreThreshold
	}

	// リクエストを作成
	body, err := json.Marshal(map[string]interface{}{
		"text":            string(in.Text),
		"score_threshold": in.ScoreThreshold,
	})
	if err != nil {
		return out, fmt.Errorf("Error marshalling request body: %q", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return out, fmt.Errorf("Error creating request: %q", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", in.APIKey))

	// リクエスト
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return out, fmt.Errorf("Error sending request: %q", err)
	}
	defer resp.Body.Close()

	// レスポンスをチェック
	if resp.StatusCode != http.StatusOK {
		return out, fmt.Errorf("Request failed with status: %q", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return out, fmt.Errorf("Error decoding response body: %q", err)
	}

	return out, nil
}

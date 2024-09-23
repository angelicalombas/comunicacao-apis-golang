package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID uint `json:"id"`
}

func CheckUserExists(userID uint) (bool, error) {
	url := fmt.Sprintf("http://user-service:8081/users/%d", userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return false, err
	}

	return user.ID == userID, nil
}

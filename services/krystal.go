package services

import (
	"encoding/json"
	"fmt"
	"io"
	"lp-assistant/models"
	"net/http"
)

// FetchStrategies fetches LP strategies for a given wallet
func FetchStrategies(wallet string, page, perPage int, status string) ([]models.StrategyProfile, error) {
	url := fmt.Sprintf("https://api.krystal.app/all/v2/strategies/profile?wallet=%s&page=%d&status=%s&perPage=%d", wallet, page, status, perPage)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result struct {
		Data []models.StrategyProfile `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// FetchPoolDetail fetches pool details for a given pool
func FetchPoolDetail(chainId int, protocol, poolAddress string) (*models.PoolDetail, error) {
	url := fmt.Sprintf("https://api.krystal.app/all/v1/lp_explorer/pool_detail?chainId=%d&protocol=%s&poolAddress=%s&includeTicks=true", chainId, protocol, poolAddress)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result struct {
		Data models.PoolDetail `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

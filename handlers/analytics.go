package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"lp-assistant/config"
	"lp-assistant/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyticsHandler handles the /analytics GET endpoint
func AnalyticsHandler(c *gin.Context) {
	wallet := c.Query("wallet")
	if wallet == "" {
		log.Println("[WARN] wallet is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet is required"})
		return
	}
	strategies, err := services.FetchStrategies(wallet, 0, 500, "open")
	if err != nil {
		log.Printf("[ERROR] failed to fetch strategies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch strategies", "details": err.Error()})
		return
	}
	if len(strategies) == 0 {
		log.Printf("[INFO] No strategies found for wallet %s", wallet)
		c.JSON(http.StatusOK, gin.H{"message": "No strategies found", "strategies": []string{}})
		return
	}
	var analyticsData []map[string]interface{}
	for _, s := range strategies {
		pool, err := services.FetchPoolDetail(s.ChainID, s.Protocol.Key, s.PoolAddress)
		if err != nil {
			log.Printf("[WARN] failed to fetch pool detail for %s: %v", s.PoolAddress, err)
			continue // skip failed pools
		}
		analyticsData = append(analyticsData, map[string]interface{}{
			"strategy": s,
			"pool":     pool,
		})
	}
	if len(analyticsData) == 0 {
		log.Printf("[INFO] No pool data could be fetched for wallet %s", wallet)
		c.JSON(http.StatusOK, gin.H{"message": "No pool data could be fetched", "strategies": strategies})
		return
	}
	// Prepare prompt for OpenAI
	prompt := fmt.Sprintf(`Analyze the following LP strategies and pools: %v. 
Return your answer as a JSON object with these fields:
{
  "total": <number of strategies>,
  "tvl_ranges": {
    "0-100k": <count>,
    "100k-1m": <count>,
    "1m-10m": <count>,
    ">10m": <count>
  },
  "proportion": <proportion info>,
  "summary": "<summary text>",
  "score": <score 0-100>,
  "recommendation": "<recommendation text>"
}
Only return valid JSON, no explanation.`, analyticsData)
	openaiKey := config.GetOpenAIKey()
	if openaiKey == "" {
		log.Println("[ERROR] OpenAI API key not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API key not set"})
		return
	}
	aiResult, err := services.CallOpenAI(openaiKey, prompt)
	if err != nil {
		log.Printf("[ERROR] OpenAI analytics failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI analytics failed", "details": err.Error()})
		return
	}
	var aiJSON map[string]interface{}
	if err := json.Unmarshal([]byte(aiResult), &aiJSON); err != nil {
		log.Printf("[ERROR] Failed to parse OpenAI JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse OpenAI JSON", "details": err.Error(), "raw": aiResult})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"analytics":  aiJSON,
		"strategies": strategies,
	})
}

// AnalyticsHandlerWithDeps allows dependency injection for easier testing
func AnalyticsHandlerWithDeps(
	fetchStrategies func(string, int, int, string) ([]interface{}, error),
	fetchPoolDetail func(int, string, string) (interface{}, error),
	callOpenAI func(string, string) (string, error),
	getOpenAIKey func() string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		wallet := c.Query("wallet")
		if wallet == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wallet is required"})
			return
		}
		strategies, err := fetchStrategies(wallet, 0, 500, "open")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch strategies", "details": err.Error()})
			return
		}
		if len(strategies) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No strategies found", "strategies": []string{}})
			return
		}
		var analyticsData []map[string]interface{}
		for _, s := range strategies {
			strat, ok := s.(interface {
				GetChainID() int
				GetProtocolKey() string
				GetPoolAddress() string
			})
			if !ok {
				continue
			}
			pool, err := fetchPoolDetail(strat.GetChainID(), strat.GetProtocolKey(), strat.GetPoolAddress())
			if err != nil {
				continue
			}
			analyticsData = append(analyticsData, map[string]interface{}{
				"strategy": strat,
				"pool":     pool,
			})
		}
		if len(analyticsData) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No pool data could be fetched", "strategies": strategies})
			return
		}
		prompt := "test"
		openaiKey := getOpenAIKey()
		if openaiKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API key not set"})
			return
		}
		aiResult, err := callOpenAI(openaiKey, prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI analytics failed", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"analytics":  aiResult,
			"strategies": strategies,
		})
	}
}

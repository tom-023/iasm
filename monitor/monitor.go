package monitor

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tom-023/iasm/config"
	"github.com/tom-023/iasm/logger"
	"go.uber.org/zap"
)

func isRespond(url string, timeout time.Duration) bool {
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		logger.Logger.Info("Failed to reach URL", zap.String("url", url), zap.Error(err))
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		logger.Logger.Warn("URL responded successfully", zap.String("url", url))
		return true
	}

	logger.Logger.Info("Unexpected response status", zap.String("url", url), zap.Int("status", resp.StatusCode))
	return false
}

// Monitor checks a list of URLs with the specified timeout
func Monitor() {
	c := config.GetConfig()
	urlsStr := c.GetString("urls")
	timeoutStr := c.GetString("timeout")

	if urlsStr == "" {
		logger.Logger.Fatal("No URLs configured for monitoring")
		return
	}

	// urlsをカンマで分割して[]stringに変換
	urls := strings.Split(urlsStr, ",")

	// timeoutが設定されていない場合はデフォルト値（1分）を使用
	var timeout time.Duration
	if timeoutStr == "" {
		timeout = 1 * time.Minute
		logger.Logger.Info("No timeout configured, using default of 1 minute")
	} else {
		// timeoutStrをtime.Durationに変換
		var err error
		timeout, err = time.ParseDuration(timeoutStr)
		if err != nil {
			logger.Logger.Fatal("Invalid timeout format", zap.String("timeout", timeoutStr), zap.Error(err))
			return
		}
	}

	for _, url := range urls {
		logger.Logger.Info(fmt.Sprintf("Check URL: %s", url))
		if isRespond(url, timeout) {
			// 通知処理を後ほど追加予定
			fmt.Printf("URL %s failed to respond as expected\n", url)
		}
	}
}

package monitor

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tom-023/iasm/config"
	"github.com/tom-023/iasm/logger"
	"github.com/tom-023/iasm/notify"
	"go.uber.org/zap"
)

func isRespond(url string, timeout time.Duration) bool {
	//client := http.Client{
	//	Timeout: timeout,
	//}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // リダイレクトを無効化して、リダイレクトレスポンスを返す
		},
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

	logger.Logger.Info("URL response status", zap.String("url", url), zap.Int("status", resp.StatusCode))
	return false
}

func Monitor() {
	c := config.GetConfig()
	urlsStr := c.GetString("urls")
	timeoutStr := c.GetString("timeout")
	slackToken := c.GetString("slack_token")
	slackChannel := c.GetString("slack_channel")

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

	var alertUrls []string

	for _, url := range urls {
		if isRespond(url, timeout) {
			alertUrls = append(alertUrls, fmt.Sprintf("- %s", url))
		}
	}

	if len(alertUrls) > 0 {
		message := fmt.Sprintf("*[Alert] The following URL is now available for connection.*\n%s", strings.Join(alertUrls, "\n"))
		err := notify.NotifySlack(slackToken, slackChannel, message)
		if err != nil {
			logger.Logger.Error("Failed to send Slack notification", zap.Error(err))
		}
	}
}

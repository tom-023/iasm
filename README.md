# iasm
InternalAppShieldMonitor (IASM) is a tool that monitors whether internal applications are incommunicable from the outside.

## Features
- Periodic checks of specified URLs for external accessibility. 
- Sends alerts to Slack when monitored URLs are accessible.

## Usage

### Required Environment Variables
Set the following environment variables to configure the application:

- URLS
Comma-separated list of URLs to monitor.
Example: https://example.com,https://test.example.com
Required

- MONITOR_INTERVAL
Interval between monitoring checks. Acceptable formats: 1m, 5m, 10s.
Default: 5m

- TIMEOUT
Timeout duration for each request.
Default: 1m

- SLACK_TOKEN
Slack Bot token for sending messages.
Required

- SLACK_CHANNEL
Slack channel for sending alerts.
Example: #alerts
Required

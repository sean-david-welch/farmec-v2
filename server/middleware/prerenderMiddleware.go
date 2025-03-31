package middleware

import (
	"io"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type PrerenderMiddleware struct {
	Token string
}

func NewPrerenderMiddleware(token string) *PrerenderMiddleware {
	return &PrerenderMiddleware{
		Token: token,
	}
}

func (middleware *PrerenderMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip if this is not a GET request
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		userAgent := c.Request.Header.Get("User-Agent")
		// Regex to detect search engine bots
		botRegex := regexp.MustCompile(`(?i)(googlebot|bingbot|yandex|baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora|pinterest|slackbot|vkShare|W3C_Validator)`)

		// Check if request is from a bot
		if botRegex.MatchString(userAgent) {
			// Get the current URL
			host := c.Request.Host
			// Format the URL for Prerender
			prerenderURL := "https://service.prerender.io/https://" + host + c.Request.URL.Path

			// Make request to Prerender service
			prerenderReq, _ := http.NewRequest("GET", prerenderURL, nil)

			// Add Prerender token
			prerenderReq.Header.Add("X-Prerender-Token", middleware.Token)

			client := &http.Client{}
			resp, err := client.Do(prerenderReq)

			if err == nil && resp.StatusCode == 200 {
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						return
					}
				}(resp.Body)

				// Copy status code
				c.Status(resp.StatusCode)

				// Copy headers
				for k, v := range resp.Header {
					if k != "Content-Length" {
						for _, val := range v {
							c.Header(k, val)
						}
					}
				}

				// Copy the prerendered content
				io.Copy(c.Writer, resp.Body)

				// Abort the chain since we've already written the response
				c.Abort()
				return
			}
		}

		// If not a bot or prerender failed, continue the chain
		c.Next()
	}
}

package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"io"
	"loan-ext/app/entity/vo"
	"sort"
	"strconv"
	"strings"
	"time"
)

// generateSignature creates an HMAC signature for the given JSON body using the provided secret key.
func generateSignature(secretKey string, requestBodyMap map[string]interface{}, requestTimestamp string) string {
	// Sort JSON keys alphabetically (including "timestamp")
	keys := make([]string, 0, len(requestBodyMap)+1)
	for key := range requestBodyMap {
		keys = append(keys, key)
	}
	keys = append(keys, "timestamp")
	sort.Strings(keys)

	// Concatenate parameters as "key=value" pairs
	var sortedParams []string
	for _, key := range keys {
		var value string
		if key == "timestamp" {
			value = requestTimestamp
		} else {
			value = fmt.Sprintf("%v", requestBodyMap[key])
		}
		sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", key, value))
	}
	data := strings.Join(sortedParams, "&")

	// Generate HMAC-SHA256 signature
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// validateTimestamp checks whether the timestamp is valid and within an acceptable range.
func validateTimestamp(requestTimestamp string, allowedSkewMinutes float64) error {
	if requestTimestamp == "" {
		return errors.New(vo.INVALID_TIMESTAMP)
	}
	timestampSeconds, err := strconv.ParseInt(requestTimestamp, 10, 64)
	if err != nil {
		return errors.New(vo.TIMESTAMP_ERROR)
	}
	parsedTime := time.Unix(timestampSeconds, 0)

	// Check time difference only once
	timeDifference := time.Since(parsedTime).Minutes()

	if timeDifference > allowedSkewMinutes || timeDifference < -allowedSkewMinutes {
		return errors.New(vo.INVALID_TIMESTAMP)
	}

	return nil
}

// Signature is a Gin middleware function for verifying HMAC signatures.
func Signature(secretKeyParam ...string) gin.HandlerFunc {
	secretKey := config.GetString("system.secret")
	if len(secretKeyParam) > 0 {
		secretKey = secretKeyParam[0]
	}

	return func(c *gin.Context) {
		// Retrieve the signature and timestamp from request headers.
		providedSignature := c.GetHeader("signature")
		if providedSignature == "" {
			vo.NewBase().Fail(c, vo.SIGNATURE_MISSING)
			c.Abort()
			return
		}

		requestTimestamp := c.GetHeader("timestamp") // Changed timestamp to X-Timestamp for consistency
		if err := validateTimestamp(requestTimestamp, 10); err != nil {
			vo.NewBase().Fail(c, err.Error())
			c.Abort()
			return
		}

		// Read and restore the request body using io.TeeReader.
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			vo.NewBase().Fail(c, vo.ERR_READ_BODY, err)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(strings.NewReader(string(body)))

		// Parse the request body as JSON.
		var requestBodyMap map[string]interface{}
		json.Unmarshal(body, &requestBodyMap)

		// Generate the expected signature.
		expectedSignature := generateSignature(secretKey, requestBodyMap, requestTimestamp)

		// Compare the provided signature with the expected signature.
		if providedSignature != expectedSignature {
			vo.NewBase().Fail(c, vo.INVALID_SIGNATURE)
			c.Abort()
			return
		}

		// Proceed with the request if validation passes.
		c.Next()
	}
}

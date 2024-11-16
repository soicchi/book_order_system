package middlewares

import (
	"encoding/json"
	"log/slog"

	"github.com/soicchi/book_order_system/internal/logging"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var sensitiveFields = []string{"password", "email"}

type SensitiveField int

const (
	Password SensitiveField = iota
	Email
)

func (sf SensitiveField) String() string {
	switch sf {
	case Password:
		return "password"
	case Email:
		return "email"
	default:
		return ""
	}
}

func CustomBodyDump(logger logging.Logger) middleware.BodyDumpHandler {
	return func(ctx echo.Context, reqBody, resBody []byte) {
		var req, res map[string]interface{}

		// convert request bytes to map
		json.Unmarshal(reqBody, &req)
		json.Unmarshal(resBody, &res)

		// mask sensitive fields
		maskedRequest := maskFields(req)
		maskedResponse := maskFields(res)

		// output logs
		logger.LogAttrs(ctx.Request().Context(), slog.LevelInfo, "request dump", slog.Any("request", maskedRequest))
		logger.LogAttrs(ctx.Request().Context(), slog.LevelInfo, "response dump", slog.Any("response", maskedResponse))
	}
}

func shouldMaskField(field string) bool {
	switch field {
	case Password.String(), Email.String():
		return true
	}

	return false
}

func maskFields(body map[string]interface{}) map[string]interface{} {
	// body is empty
	if len(body) == 0 {
		return map[string]interface{}{}
	}

	for key, value := range body {
		if shouldMaskField(key) {
			body[key] = "*****"
		} else if nestedBody, ok := value.(map[string]interface{}); ok {
			maskFields(nestedBody)
		}
	}

	return body
}

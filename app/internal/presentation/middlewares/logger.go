package middlewares

import (
	"encoding/json"
	"log/slog"
	"slices"

	"github.com/soicchi/book_order_system/internal/logging"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var sensitiveFields = []string{"password", "email"}

func CustomBodyDump(logger logging.Logger) middleware.BodyDumpHandler {
	return func(ctx echo.Context, reqBody, resBody []byte) {
		var req, res map[string]interface{}

		// convert request bytes to map
		json.Unmarshal(reqBody, &req)
		json.Unmarshal(resBody, &res)

		// mask sensitive fields
		maskedRequest := maskField(req)
		maskedResponse := maskField(res)

		// output logs
		logger.LogAttrs(ctx.Request().Context(), slog.LevelInfo, "request dump", slog.Any("request", maskedRequest))
		logger.LogAttrs(ctx.Request().Context(), slog.LevelInfo, "response dump", slog.Any("response", maskedResponse))
	}
}

func shouldMaskField(field string) bool {
	return slices.Contains(sensitiveFields, field)
}

func maskField(body map[string]interface{}) *map[string]interface{} {
	// body is empty
	if body == nil {
		return nil
	}

	for key, value := range body {
		if shouldMaskField(key) {
			body[key] = "*****"
		} else if nestedBody, ok := value.(map[string]interface{}); ok {
			maskField(nestedBody)
		}
	}

	return &body
}

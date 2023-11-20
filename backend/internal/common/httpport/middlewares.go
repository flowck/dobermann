package httpport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/flowck/doberman/internal/common/logs"
)

func loggerMiddleware(logger *logs.Logger) echo.MiddlewareFunc {
	if logger.Level.String() == logs.DebugLevel.String() {
		return middleware.Logger()
	}

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.WithFields(logs.Fields{
				"method": v.Method,
				"URI":    v.URI,
				"status": v.Status,
				"error":  v.Error,
			}).Info("")
			return nil
		},
	})
}

func errorHandler(logger *logs.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		logger.WithError(err).Debug("Handler failed with error")

		code := http.StatusInternalServerError
		errorSlug := "internal-server-error"
		errorMessage := err.Error()

		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			errorMessage = e.Error()
		case *HandlerError:
			code = e.Code
			errorSlug = e.Slug()
			errorMessage = e.Error()
		}

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, map[string]string{
				"error":   errorSlug,
				"message": errorMessage,
			})
		}
		if err != nil {
			logger.WithError(err).Error("Failed to send error response")
		}
	}
}

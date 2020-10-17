package healthcontrollers

import (
	json "encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

// Health is a method to Health Check
func Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := map[string]interface{}{
			"message": "Deploy OK",
		}
		jsonResponse, _ := json.Marshal(response)
		return c.JSON(http.StatusOK, string(jsonResponse))
	}
}

package netflix

import "github.com/labstack/echo"

func (h *Handler) GetMovies(c echo.Context) error {
	return c.JSON(200, "movies")
}

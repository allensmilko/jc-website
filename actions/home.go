package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("currentBanner", "assets/images/home-banner.png")
	c.Set("profileImage", "assets/images/profile.png")
	c.Set("portFoliotitle", "Portfolio")

	fmt.Print(c)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

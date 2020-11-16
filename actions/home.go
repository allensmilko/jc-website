package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jc_website/websitestructs"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	buf := `[
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-1.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-2.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-3.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-4.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-5.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-6.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-7.png"
		},
		{
			"name": "Creation of all the DNA of a new born brand.",
			"image": "assets/images/pg-8.png"
		}
	]`

	var projects []websitestructs.Project
	fmt.Println("TEST---------AAA-----------")

	fmt.Printf("%#v", projects)

	json.Unmarshal([]byte(buf), &projects)
	fmt.Println("TEST--------------------")

	fmt.Printf("Birds : %+v", projects)

	c.Set("currentBanner", "assets/images/home-banner.png")
	c.Set("profileImage", "assets/images/profile.png")
	c.Set("portFoliotitle", "Portfolio")
	c.Set("projects", projects)

	return c.Render(http.StatusOK, r.HTML("index.html"))
}

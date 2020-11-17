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
	json.Unmarshal([]byte(buf), &projects)
	c.Set("projects", projects)

	bufHab := `[
		{
			"name": "UX - UI Design",
			"image": "assets/images/ux.png",
			"description": "Design of differents types of digital platforms."
		},
		{
			"name": "Brand Identity",
			"image": "assets/images/brand.png",
			"description": "Creation of all the DNA of a newborn brand."
		},
		{
			"name": "Editorial Design",
			"image": "assets/images/editorial.png",
			"description": "Books, magazines and also all kind of print stuff."
		},
		{
			"name": "Video & Animation",
			"image": "assets/images/video-icon.png",
			"description": "Motion graphics and video edition for your brand."
		}
	]`

	var habilities []websitestructs.Hability
	json.Unmarshal([]byte(bufHab), &habilities)
	c.Set("habilities", habilities)

	fmt.Println("TEST--------------------")

	fmt.Printf("Birds : %+v", habilities)

	c.Set("currentBanner", "assets/images/home-banner.png")
	c.Set("profileImage", "assets/images/profile.png")
	c.Set("portFoliotitle", "Portfolio")
	c.Set("habilitiesTitle", "Habilidades")

	return c.Render(http.StatusOK, r.HTML("index.html"))
}

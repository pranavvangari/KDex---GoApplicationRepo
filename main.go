package main

import ( //imports necessary packages for this application

	"net/http" //imports the standard Go package for making HTTP requests

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2" //imports the Fiber web framework
	//needed for printing information
	//needed for reading information from http response
)

func main() {

	app := fiber.New() //creates a new instance of the Fiber app

	app.Get("/api/jupiter", func(c *fiber.Ctx) error {
		/*for the code above
		1. defines a route handler for the root path ('/') using the Get method
		2. func(c *fiber.Ctx) error is an anonymous function that will be executed whenever a GET request is made to the root path
		3. the fiber.Ctx parameter c represents the request context, which contains information about the incoming HTTP request and provides methods for handling the response
		*/

		resp, err := http.Get("https://en.wikipedia.org/wiki/Jupiter") //this line makes an HTTP GET request to "https://en.wikipedia.org/wiki/Jupiter" using the http.Get function from net/http package. the response -> resp, errors -> err

		if err != nil {
			return err
		}

		defer resp.Body.Close() //defers the closing of the response body until the surrounding function returns (func(c *fiber.Ctx) error)

		// body, err := io.ReadAll(resp.Body)

		// if err != nil {
		// 	return err
		// }

		// bodyString := string(body)

		// fmt.Println("response body of first 500 characters: ", bodyString[:500])

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return err
		}

		contentSection := doc.Find("#mw-content-text > div.mw-parser-output") //find the main content section of the Wikipedia page
		var content []string
		contentSection.Find("p").Each(func(i int, s *goquery.Selection) {
			content = append(content, s.Text())
		})

		//contentHTML := strings.Join(content, "<br>") // try to align the content -- not needed when using Postman

		//return c.SendString(contentHTML)

		return c.JSON(content)
	})

	app.Listen(":3000")

}
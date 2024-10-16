// Ortelius v11 ApplicationVersion Microservice that handles creating and retrieving ApplicationVersion
package main

import (
	"context"
	"encoding/json"

	_ "github.com/ortelius/scec-appver/docs"

	driver "github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ortelius/scec-commons/database"
	"github.com/ortelius/scec-commons/model"
)

var logger = database.InitLogger()
var dbconn = database.InitializeDatabase()

// GetApplicationVersions godoc
// @Summary Get a List of ApplicationVersion
// @Description Get a list of ApplicationVersion for the user.
// @Tags ApplicationVersion
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/appver [get]
func GetApplicationVersions(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	// query all the ApplicationVersion in the collection
	aql := `FOR appver in evidence
			FILTER (appver.objtype == 'ApplicationVersion')
			RETURN appver`

	// execute the query with no parameters
	if cursor, err = dbconn.Database.Query(ctx, aql, nil); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err) // log error
	}

	defer cursor.Close() // close the cursor when returning from this function

	applications := model.NewApplications() // define a list of ApplicationVersions to be returned

	for cursor.HasMore() { // loop thru all of the documents

		appver := model.NewApplicationVersion() // fetched ApplicationVersion
		var meta driver.DocumentMeta            // data about the fetch

		// fetch a document from the cursor
		if meta, err = cursor.ReadDocument(ctx, appver); err != nil {
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
		applications.Applications = append(applications.Applications, appver) // add the Application Version to the list
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key)  // log the key
	}

	return c.JSON(applications) // return the list of ApplicationVersion in JSON format
}

// GetApplicationVersion godoc
// @Summary Get a ApplicationVersion
// @Description Get a ApplicationVersionbased on the _key or name.
// @Tags ApplicationVersion
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/appver/:key [get]
func GetApplicationVersion(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	key := c.Params("key")                // key from URL
	parameters := map[string]interface{}{ // parameters
		"key": key,
	}

	// query the ApplicationVersion that match the key or name
	aql := `FOR appver in evidence
			FILTER (appver.name == @key or appver._key == @key)
			RETURN appver`

	// run the query with patameters
	if cursor, err = dbconn.Database.Query(ctx, aql, &driver.QueryOptions{BindVars: parameters}); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err)
	}

	defer cursor.Close() // close the cursor when returning from this function

	appver := model.NewApplicationVersion() // define a appver to be returned

	if cursor.HasMore() { // appver found
		var meta driver.DocumentMeta // data about the fetch

		if meta, err = cursor.ReadDocument(ctx, appver); err != nil { // fetch the document into the object
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key)

	} else { // not found so get from NFT Storage
		if jsonStr, exists := database.MakeJSON(key); exists {
			if err := json.Unmarshal([]byte(jsonStr), appver); err != nil { // convert the JSON string from LTF into the object
				logger.Sugar().Errorf("Failed to unmarshal from LTS: %v", err)
			}
		}
	}

	return c.JSON(appver) // return the appver in JSON format
}

// NewApplicationVersion godoc
// @Summary Create a ApplicationVersion
// @Description Create a new ApplicationVersion and persist it
// @Tags ApplicationVersion
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/appver [post]
func NewApplicationVersion(c *fiber.Ctx) error {

	var err error                           // for error handling
	var meta driver.DocumentMeta            // data about the document
	var ctx = context.Background()          // use default database context
	appver := model.NewApplicationVersion() // define an appver to be returned

	if err = c.BodyParser(appver); err != nil { // parse the JSON into the appver object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, dbStr := database.MakeNFT(appver) // normalize the object into NFTs and JSON string for db persistence

	logger.Sugar().Infof("%s=%s\n", cid, dbStr) // log the new nft

	var resp driver.CollectionDocumentCreateResponse
	// add the appver to the database.  Ignore if it already exists since it will be identical
	if resp, err = dbconn.Collections["applications"].CreateDocument(ctx, appver); err != nil && !shared.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	meta = resp.DocumentMeta
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collections["applications"].Name(), dbconn.Database.Name(), meta.Key)

	return c.JSON(appver) // return the appver object in JSON format.  This includes the new _key
}

// setupRoutes defines maps the routes to the functions
func setupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault)        // handle displaying the swagger
	app.Get("/msapi/appver", GetApplicationVersions)     // list of ApplicationVersion
	app.Get("/msapi/appver/:key", GetApplicationVersion) // single ApplicationVersion based on name or key
	app.Post("/msapi/appver", NewApplicationVersion)     // save a single ApplicationVersion
}

// @title Ortelius v11 ApplicationVersion Microservice
// @version 11.0.0
// @description RestAPI for the ApplicationVersion Object
// @description ![Release](https://img.shields.io/github/v/release/ortelius/scec-appver?sort=semver)
// @description ![license](https://img.shields.io/github/license/ortelius/.github)
// @description
// @description ![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-appver/build-push-chart.yml)
// @description [![MegaLinter](https://github.com/ortelius/scec-appver/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-appver/actions?query=workflow%3AMegaLinter+branch%3Amain)
// @description ![CodeQL](https://github.com/ortelius/scec-appver/workflows/CodeQL/badge.svg)
// @description [![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-appver/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-appver)
// @description
// @description ![Discord](https://img.shields.io/discord/722468819091849316)

// @termsOfService http://swagger.io/terms/
// @contact.name Ortelius Google Group
// @contact.email ortelius-dev@googlegroups.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /msapi/appver
func main() {
	port := ":" + database.GetEnvDefault("MS_PORT", "8080") // database port
	app := fiber.New()                                      // create a new fiber application
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "*",
	}))

	setupRoutes(app) // define the routes for this microservice

	if err := app.Listen(port); err != nil { // start listening for incoming connections
		logger.Sugar().Fatalf("Failed get the microservice running: %v", err)
	}
}

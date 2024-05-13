// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Ortelius Google Group",
            "email": "ortelius-dev@googlegroups.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/msapi/appver": {
            "get": {
                "description": "Get a list of ApplicationVersion for the user.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApplicationVersion"
                ],
                "summary": "Get a List of ApplicationVersion",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create a new ApplicationVersion and persist it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApplicationVersion"
                ],
                "summary": "Create a ApplicationVersion",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/msapi/appver/:key": {
            "get": {
                "description": "Get a ApplicationVersionbased on the _key or name.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApplicationVersion"
                ],
                "summary": "Get a ApplicationVersion",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "11.0.0",
	Host:             "localhost:3000",
	BasePath:         "/msapi/appver",
	Schemes:          []string{},
	Title:            "Ortelius v11 ApplicationVersion Microservice",
	Description:      "RestAPI for the ApplicationVersion Object\n![Release](https://img.shields.io/github/v/release/ortelius/scec-appver?sort=semver)\n![license](https://img.shields.io/github/license/ortelius/.github)\n\n![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-appver/build-push-chart.yml)\n[![MegaLinter](https://github.com/ortelius/scec-appver/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-appver/actions?query=workflow%3AMegaLinter+branch%3Amain)\n![CodeQL](https://github.com/ortelius/scec-appver/workflows/CodeQL/badge.svg)\n[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-appver/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-appver)\n\n![Discord](https://img.shields.io/discord/722468819091849316)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

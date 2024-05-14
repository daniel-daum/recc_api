// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Daniel Daum"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Returns JSON response \"healthy\"",
                "produces": [
                    "application/json"
                ],
                "summary": "Health",
                "responses": {
                    "200": {
                        "description": "HealthResponse JSON",
                        "schema": {
                            "$ref": "#/definitions/main.HealthResponse"
                        }
                    }
                }
            }
        },
        "/reference": {
            "get": {
                "description": "Returns Recc API scalar documentation",
                "produces": [
                    "text/html"
                ],
                "summary": "Reference",
                "responses": {}
            }
        }
    },
    "definitions": {
        "main.HealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "healthy"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Recc API",
	Description:      "The Recc API documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

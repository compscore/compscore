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
            "name": "1nv8rzim",
            "url": "https://github.com/compsore/compscore/issues"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login": {
            "post": {
                "description": "Authenticate into another team and return a JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Authenticate into another team",
                "parameters": [
                    {
                        "description": "Team name",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdminLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cookie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/admin/password": {
            "post": {
                "description": "Reset password of another team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Reset password of another team",
                "parameters": [
                    {
                        "description": "Team name and new password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdminPasswordReset"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/check/{check}": {
            "get": {
                "description": "Get a check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "check"
                ],
                "summary": "Get a check",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Check ID",
                        "name": "check",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Check"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/checks": {
            "get": {
                "description": "Get all checks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "check"
                ],
                "summary": "Get all checks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Check"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/credential/{check}": {
            "get": {
                "security": [
                    {
                        "ServiceAuth": []
                    }
                ],
                "description": "Get a credential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credential"
                ],
                "summary": "Get a credential",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Check name",
                        "name": "check",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Credential"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ServiceAuth": []
                    }
                ],
                "description": "Update a credential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credential"
                ],
                "summary": "Update a credential",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Check name",
                        "name": "check",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CredentialEdit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/credentials": {
            "get": {
                "security": [
                    {
                        "ServiceAuth": []
                    }
                ],
                "description": "Get all credentials for a team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credential"
                ],
                "summary": "Get all credentials for a team",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Credential"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/engine/start": {
            "post": {
                "security": [
                    {
                        "ServiceAuth": []
                    }
                ],
                "description": "Start the engine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "engine"
                ],
                "summary": "Start the engine",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Status"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/engine/status": {
            "get": {
                "description": "Status of the engine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "engine"
                ],
                "summary": "Status of the engine",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/engine/stop": {
            "post": {
                "security": [
                    {
                        "ServiceAuth": []
                    }
                ],
                "description": "Stop the engine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "engine"
                ],
                "summary": "Stop the engine",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Status"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate a user and return a JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate a user",
                "parameters": [
                    {
                        "description": "Username and password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cookie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/password": {
            "post": {
                "description": "Change a user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Change a user's password",
                "parameters": [
                    {
                        "description": "Old and new password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password changed"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/round/latest": {
            "get": {
                "description": "Get the latest round",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "round"
                ],
                "summary": "Get the latest round",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Round"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/round/{round}": {
            "get": {
                "description": "Get a round",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "round"
                ],
                "summary": "Get a round",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Round number",
                        "name": "round",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Round"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/rounds": {
            "get": {
                "description": "Get all rounds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "round"
                ],
                "summary": "Get all rounds",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Round"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/team/{team}": {
            "get": {
                "description": "Get a team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "team"
                ],
                "summary": "Get a team",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Team ID",
                        "name": "team",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Team"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AdminLogin": {
            "description": "body of admin login request",
            "type": "object",
            "properties": {
                "team": {
                    "type": "string"
                }
            }
        },
        "models.AdminPasswordReset": {
            "description": "body of admin password reset request",
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                }
            }
        },
        "models.ChangePassword": {
            "description": "body of change password request",
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "models.Check": {
            "description": "score check",
            "type": "object",
            "properties": {
                "edges": {
                    "type": "object",
                    "properties": {
                        "credential": {
                            "type": "array",
                            "items": {}
                        },
                        "status": {
                            "type": "array",
                            "items": {}
                        }
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Cookie": {
            "description": "response of login request",
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string"
                },
                "expiration": {
                    "type": "integer"
                },
                "httponly": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "secure": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.Credential": {
            "description": "credential of a check",
            "type": "object",
            "properties": {
                "edges": {
                    "type": "object",
                    "properties": {
                        "check": {},
                        "team": {}
                    }
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.CredentialEdit": {
            "description": "body of credential edit request",
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "description": "response of login request",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "description": "body of login request",
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Round": {
            "description": "scoring round",
            "type": "object",
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "edges": {
                    "type": "object",
                    "properties": {
                        "status": {
                            "type": "array",
                            "items": {}
                        }
                    }
                },
                "number": {
                    "type": "integer"
                }
            }
        },
        "models.Status": {
            "description": "status of a check",
            "type": "object",
            "properties": {
                "edges": {
                    "type": "object",
                    "properties": {
                        "check": {},
                        "round": {},
                        "team": {}
                    }
                },
                "error": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.Team": {
            "description": "team",
            "type": "object",
            "properties": {
                "edges": {
                    "type": "object",
                    "properties": {
                        "credential": {
                            "type": "array",
                            "items": {}
                        },
                        "status": {
                            "type": "array",
                            "items": {}
                        }
                    }
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "roles": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ServiceAuth": {
            "description": "JWT for authentication",
            "type": "apiKey",
            "name": "auth",
            "in": "cookie"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Compscore API",
	Description:      "This is the API for the Compscore application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

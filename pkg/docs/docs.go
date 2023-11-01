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
        "/api/admin/login": {
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
        "/api/admin/password": {
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
        "/api/checks": {
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
        "/api/login": {
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
        "/api/password": {
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

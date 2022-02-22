// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/csv/films": {
            "get": {
                "description": "Export all films data to csv file.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Export all films data to csv file.",
                "responses": {
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Import films data from csv file.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Import films data from csv file.",
                "responses": {
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/favourites": {
            "post": {
                "description": "Allow to add a film to favourites list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Favourites"
                ],
                "summary": "Add film to favourites",
                "parameters": [
                    {
                        "description": "Add favourites",
                        "name": "film_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.addFavouriteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns object with favourite pk, film_id and user_id",
                        "schema": {
                            "$ref": "#/definitions/http.addFavouriteResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/favourites/{film_id}": {
            "delete": {
                "description": "Allow to remove a film from favourites list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Favourites"
                ],
                "summary": "Remove a film from favourites",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Remove favourites",
                        "name": "film_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.emptyResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/films": {
            "get": {
                "description": "Allow to get all films by some filters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Get all films.",
                "responses": {
                    "200": {
                        "description": "Returns films",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/film.Film"
                            }
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Allow to register a new film .",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Create a new film.",
                "parameters": [
                    {
                        "description": "create film",
                        "name": "film",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createFilmRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.createFilmResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/films/{film_id}": {
            "get": {
                "description": "Given a film_id returns this film.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Get a specific film.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Film id",
                        "name": "film_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns film",
                        "schema": {
                            "$ref": "#/definitions/film.Film"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Allow to remove a film from records.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Destroy a film from records.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Destroy film",
                        "name": "film_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.emptyResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Allow to update one or many fields of films.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Allow to update one or many fields of films.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "film id",
                        "name": "film_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data to update",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateFilmRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.emptyResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Returns a valid token if credentials are right.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get valid credentials.",
                "parameters": [
                    {
                        "description": "username \u0026 password",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns a token",
                        "schema": {
                            "$ref": "#/definitions/http.loginResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Send current token to blacklist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Invalidate actual token.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.emptyResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register user inside API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Allow users to register to consume this API.",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.registerResponse"
                        }
                    },
                    "400": {
                        "description": "error 400",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "error 500",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "film.Film": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "director": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "release_year": {
                    "type": "integer"
                },
                "starring": {
                    "type": "string"
                },
                "sypnosis": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.addFavouriteRequest": {
            "type": "object",
            "properties": {
                "film_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http.addFavouriteResponse": {
            "type": "object",
            "properties": {
                "film_id": {
                    "type": "integer",
                    "example": 2
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 7
                }
            }
        },
        "http.createFilmRequest": {
            "type": "object",
            "required": [
                "director",
                "genre",
                "release_year",
                "starring",
                "title"
            ],
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "director": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Quentin Tarantino"
                },
                "genre": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "action, comedy, war"
                },
                "release_year": {
                    "type": "integer",
                    "example": 2009
                },
                "starring": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Brad Pitt, Christoph Waltz, Michael Fassbender"
                },
                "sypnosis": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Inglourious Basterds"
                }
            }
        },
        "http.createFilmResponse": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "director": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "release_year": {
                    "type": "integer"
                },
                "starring": {
                    "type": "string"
                },
                "sypnosis": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.emptyResponse": {
            "type": "object"
        },
        "http.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "http.loginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 5,
                    "example": "secret"
                },
                "username": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 5,
                    "example": "user21"
                }
            }
        },
        "http.loginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
                }
            }
        },
        "http.registerResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.registerUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 5,
                    "example": "secret"
                },
                "username": {
                    "type": "string",
                    "maxLength": 75,
                    "minLength": 5,
                    "example": "user21"
                }
            }
        },
        "http.updateFilmRequest": {
            "type": "object",
            "properties": {
                "director": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Quentin Tarantino"
                },
                "genre": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "action, comedy, war"
                },
                "release_year": {
                    "type": "integer",
                    "example": 2009
                },
                "starring": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Brad Pitt, Christoph Waltz, Michael Fassbender"
                },
                "sypnosis": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Inglourious Basterds"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Kentech-Films",
	Description:      "This API provides endpoints to manage films and register users.\n[Read me](https://github.com/migalpha/kentech-films)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}

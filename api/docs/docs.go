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
            "name": "Nathan Faucett",
            "email": "nathanfaucett@gmail.com"
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
        "/cameras": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "camera"
                ],
                "summary": "Get all cameras",
                "operationId": "cameras",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Camera"
                            }
                        }
                    }
                }
            }
        },
        "/cameras/{hardwareId}": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "camera"
                ],
                "summary": "Get camera by hardware id",
                "operationId": "camera-by-hardware-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hardware Id",
                        "name": "hardwareId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Camera"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "camera"
                ],
                "summary": "update camera by hardware id",
                "operationId": "update-camera-by-hardware-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hardware Id",
                        "name": "hardwareId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Camera",
                        "name": "updates",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpsertCamera"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Camera"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "Get Health Check",
                "operationId": "health-check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Health"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Health"
                        }
                    }
                }
            }
        },
        "/ice-servers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "Get ICE servers",
                "operationId": "ice-servers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ICEServer"
                            }
                        }
                    }
                }
            }
        },
        "/p2p-access": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "Get p2p access info",
                "operationId": "p2p-access",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/P2PAccess"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token"
                ],
                "summary": "create a token by authenticating a user",
                "operationId": "token",
                "parameters": [
                    {
                        "description": "user credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Credentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "current-user"
                ],
                "summary": "Get current user",
                "operationId": "current-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    }
                }
            }
        },
        "/user/reset-password": {
            "patch": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "current-user"
                ],
                "summary": "Resets a user's password",
                "operationId": "reset-password",
                "parameters": [
                    {
                        "description": "reset user's password",
                        "name": "resetPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ResetPassword"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Errors"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "app"
                ],
                "summary": "Get Version",
                "operationId": "version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Version"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Camera": {
            "type": "object",
            "required": [
                "discovered",
                "hardwareId",
                "mediaUris",
                "name",
                "record",
                "saved",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string",
                    "format": "date-time"
                },
                "discovered": {
                    "type": "boolean"
                },
                "hardwareId": {
                    "type": "string"
                },
                "mediaUris": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "record": {
                    "type": "boolean"
                },
                "recordWindow": {
                    "type": "integer"
                },
                "saved": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "Credentials": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "ErrorMessage": {
            "type": "object",
            "required": [
                "error",
                "parameters"
            ],
            "properties": {
                "error": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "Errors": {
            "type": "object",
            "required": [
                "errors"
            ],
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/ErrorMessage"
                        }
                    }
                }
            }
        },
        "Health": {
            "type": "object",
            "required": [
                "date"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "ICEServer": {
            "type": "object",
            "properties": {
                "credential": {},
                "credentialType": {
                    "type": "string"
                },
                "urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "P2PAccess": {
            "type": "object",
            "required": [
                "host",
                "id",
                "password",
                "ssl"
            ],
            "properties": {
                "host": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "ssl": {
                    "type": "boolean"
                }
            }
        },
        "ResetPassword": {
            "type": "object",
            "required": [
                "password",
                "passwordConfirmation"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "passwordConfirmation": {
                    "type": "string"
                }
            }
        },
        "Token": {
            "type": "object",
            "required": [
                "accessToken",
                "expiresIn",
                "issuedTokenType",
                "refreshToken",
                "refreshTokenExpiresIn",
                "tokenType"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "type": "integer"
                },
                "issuedTokenType": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "refreshTokenExpiresIn": {
                    "type": "integer"
                },
                "tokenType": {
                    "type": "string"
                }
            }
        },
        "UpsertCamera": {
            "type": "object"
        },
        "User": {
            "type": "object",
            "required": [
                "createdAt",
                "updatedAt",
                "username"
            ],
            "properties": {
                "createdAt": {
                    "type": "string",
                    "format": "date-time"
                },
                "updatedAt": {
                    "type": "string",
                    "format": "date-time"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Version": {
            "type": "object",
            "required": [
                "build",
                "version"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "Locale": {
            "type": "apiKey",
            "name": "X-Locale",
            "in": "header"
        },
        "Timezone": {
            "type": "apiKey",
            "name": "X-Timezone",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "IPCameras API",
	Description:      "IPCameras API API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

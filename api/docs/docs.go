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
                "operationId": "cameraByHardwareId",
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
                "operationId": "updateCameraByHardwareId",
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
                "operationId": "healthCheck",
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
        "/p2p-access": {
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
                "summary": "Get p2p access info",
                "operationId": "p2pAccess",
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
        "UpsertCamera": {
            "type": "object"
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
        "TenentId": {
            "type": "apiKey",
            "name": "Tenent-Id",
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

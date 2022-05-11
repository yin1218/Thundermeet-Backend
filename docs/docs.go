// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Wu, Chien Yin and Yeh, Hsiao Li"
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
        "/v1/events/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The body to create an event",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/EventFormat"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Update"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/login/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/resetPassword": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ForgotInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "EventFormat": {
            "type": "object",
            "required": [
                "dateOrDays",
                "endTime",
                "eventName",
                "isPriorityEnabled",
                "startTime"
            ],
            "properties": {
                "dateOrDays": {
                    "description": "required",
                    "type": "boolean",
                    "example": true
                },
                "endDate": {
                    "description": "optional",
                    "type": "string",
                    "example": "2021-01-10T11:00:00.000Z"
                },
                "endDay": {
                    "description": "optional",
                    "type": "string",
                    "example": "7"
                },
                "endTime": {
                    "description": "required",
                    "type": "string",
                    "example": "1975-08-19T23:00:00.000Z"
                },
                "eventName": {
                    "description": "required",
                    "type": "string",
                    "example": "OR first meet"
                },
                "isPriorityEnabled": {
                    "description": "required",
                    "type": "boolean",
                    "example": true
                },
                "startDate": {
                    "description": "optional",
                    "type": "string",
                    "example": "2021-01-01T11:00:00.000Z"
                },
                "startDay": {
                    "description": "optional",
                    "type": "string",
                    "example": "1"
                },
                "startTime": {
                    "description": "required",
                    "type": "string",
                    "example": "1975-08-19T11:00:00.000Z"
                }
            }
        },
        "Login": {
            "type": "object",
            "required": [
                "password",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                }
            }
        },
        "Register": {
            "type": "object",
            "required": [
                "password",
                "passwordAnswer",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                },
                "userName": {
                    "type": "string",
                    "example": "Christine Wang"
                }
            }
        },
        "Update": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userName": {
                    "type": "string",
                    "example": "Christine Wang"
                }
            }
        },
        "controller.ForgotInfo": {
            "type": "object",
            "required": [
                "password",
                "passwordAnswer",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ThunderMeet APIs",
	Description:      "This is the backend server for the Thundermeet App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

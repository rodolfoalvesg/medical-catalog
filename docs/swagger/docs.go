// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/v1/medical-catalog/categories": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create a new category\nOtherwise it will return HTTP status code 409.",
                "tags": [
                    "Category"
                ],
                "summary": "Create a new category",
                "parameters": [
                    {
                        "description": "Category input",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category created successfully",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "409": {
                        "description": "Category already exists",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    }
                }
            }
        },
        "/admin/v1/medical-catalog/login/signin": {
            "post": {
                "description": "Create a new login to access the system.",
                "tags": [
                    "Login"
                ],
                "summary": "Create a new login",
                "parameters": [
                    {
                        "description": "Login input",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successfully",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid email or password",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    }
                }
            }
        },
        "/admin/v1/medical-catalog/specialties": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create a new specialty\nOtherwise it will return HTTP status code 409.",
                "tags": [
                    "Specialty"
                ],
                "summary": "Create a new specialty",
                "parameters": [
                    {
                        "description": "Specialty input",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateSpecialtyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Specialty created successfully",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "409": {
                        "description": "Specialty already exists",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    }
                }
            }
        },
        "/service/v1/medical-catalog/users": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create a new user to access the system.\nOtherwise it will return HTTP status code 409.",
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User input",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "RequestID": {
                                "type": "string",
                                "description": "x-request-id"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CreateCategoryInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.CreateSpecialtyInput": {
            "type": "object",
            "properties": {
                "specialtyName": {
                    "type": "string"
                }
            }
        },
        "schema.CreateUserInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.SignInInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "description": "This is a medical catalog API.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Medical Catalog API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

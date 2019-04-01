// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "FSME server for UI",
    "title": "FSME",
    "version": "0.1.0"
  },
  "host": "127.0.0.1:8080",
  "paths": {
    "/": {
      "get": {
        "description": "Get list fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "getFSMList",
        "parameters": [
          {
            "type": "boolean",
            "format": "bool",
            "name": "enable",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int",
            "default": 10,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "get list the fsm",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/fsm"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Create a new fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "addFSM",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{_id}": {
      "get": {
        "description": "Get a fsm by ID",
        "tags": [
          "fsm"
        ],
        "operationId": "getFSM",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "destroyFSM",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "description": "Update a fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "updateFSM",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "string",
          "name": "_id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{_id}/event": {
      "post": {
        "description": "Send event to fsm by ID",
        "tags": [
          "event"
        ],
        "operationId": "sendEventFSM",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/event"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "string",
          "name": "_id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "code": {
          "type": "integer",
          "format": "int"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "event": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "state": {
          "type": "string"
        }
      }
    },
    "fsm": {
      "type": "object",
      "required": [
        "rules"
      ],
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "callbacks": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "events": {
          "type": "string"
        },
        "rules": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "FSME server for UI",
    "title": "FSME",
    "version": "0.1.0"
  },
  "host": "127.0.0.1:8080",
  "paths": {
    "/": {
      "get": {
        "description": "Get list fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "getFSMList",
        "parameters": [
          {
            "type": "boolean",
            "format": "bool",
            "name": "enable",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int",
            "default": 10,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "get list the fsm",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/fsm"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Create a new fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "addFSM",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{_id}": {
      "get": {
        "description": "Get a fsm by ID",
        "tags": [
          "fsm"
        ],
        "operationId": "getFSM",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "destroyFSM",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "description": "Update a fsm",
        "tags": [
          "fsm"
        ],
        "operationId": "updateFSM",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/fsm"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "string",
          "name": "_id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{_id}/event": {
      "post": {
        "description": "Send event to fsm by ID",
        "tags": [
          "event"
        ],
        "operationId": "sendEventFSM",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/event"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "string",
          "name": "_id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "code": {
          "type": "integer",
          "format": "int"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "event": {
      "type": "object",
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "state": {
          "type": "string"
        }
      }
    },
    "fsm": {
      "type": "object",
      "required": [
        "rules"
      ],
      "properties": {
        "_id": {
          "type": "string",
          "readOnly": true
        },
        "callbacks": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "events": {
          "type": "string"
        },
        "rules": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
}

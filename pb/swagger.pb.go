package pb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/leagues": {
      "get": {
        "summary": "List returns all available League",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbGetLeaguesResponse"
            }
          }
        },
        "tags": [
          "LeagueService"
        ]
      },
      "post": {
        "summary": "Create creates a new league",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateLeagueResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateLeagueRequest"
            }
          }
        ],
        "tags": [
          "LeagueService"
        ]
      }
    },
    "/v1/leagues/{name}": {
      "get": {
        "summary": "Get return a league",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbGetLeagueResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "LeagueService"
        ]
      }
    }
  },
  "definitions": {
    "pbCreateLeagueRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "website": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbCreateLeagueResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "league": {
          "$ref": "#/definitions/pbLeague"
        }
      }
    },
    "pbDetails": {
      "type": "object",
      "properties": {
        "website": {
          "type": "string",
          "format": "string"
        },
        "address": {
          "type": "string",
          "format": "string"
        },
        "email": {
          "type": "string",
          "format": "string"
        },
        "phonenumber": {
          "type": "string",
          "format": "string"
        },
        "fax": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDiscipline": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "title": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbGetLeagueRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbGetLeagueResponse": {
      "type": "object",
      "properties": {
        "league": {
          "$ref": "#/definitions/pbLeague"
        }
      }
    },
    "pbGetLeaguesRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbGetLeaguesResponse": {
      "type": "object",
      "properties": {
        "leagues": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbLeague"
          }
        }
      }
    },
    "pbLeague": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "details": {
          "$ref": "#/definitions/pbDetails"
        },
        "levels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbLevel"
          }
        },
        "disciplines": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbDiscipline"
          }
        }
      },
      "title": "League define a pelota league"
    },
    "pbLevel": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "title": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
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
    },
    "/v1/tournaments": {
      "get": {
        "summary": "List returns all available Tournament",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbGetTournamentsResponse"
            }
          }
        },
        "tags": [
          "TournamentService"
        ]
      },
      "post": {
        "summary": "Create creates a new tournament",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateTournamentResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateTournamentRequest"
            }
          }
        ],
        "tags": [
          "TournamentService"
        ]
      }
    },
    "/v1/tournaments/{name}": {
      "get": {
        "summary": "Get return a tournament",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbGetTournamentResponse"
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
          "TournamentService"
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
    "pbCreateTournamentRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbGroup"
          }
        }
      }
    },
    "pbCreateTournamentResponse": {
      "type": "object",
      "properties": {
        "tournament": {
          "$ref": "#/definitions/pbTournament"
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
    "pbGetTournamentRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbGetTournamentResponse": {
      "type": "object",
      "properties": {
        "tournament": {
          "$ref": "#/definitions/pbTournament"
        }
      }
    },
    "pbGetTournamentsRequest": {
      "type": "object"
    },
    "pbGetTournamentsResponse": {
      "type": "object",
      "properties": {
        "tournaments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbTournament"
          }
        }
      }
    },
    "pbGroup": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "team": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbTeam"
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
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
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
    },
    "pbRound": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "first": {
          "$ref": "#/definitions/pbTeam"
        },
        "second": {
          "$ref": "#/definitions/pbTeam"
        },
        "score": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbTeam": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbTournament": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "groups": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/pbGroup"
          }
        },
        "rounds": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbRound"
          }
        }
      },
      "title": "Tournament define a pelota league"
    }
  }
}
`
)

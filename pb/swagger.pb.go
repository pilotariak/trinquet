package pb 

const (
api = `{"swagger":"2.0","basePath":"","info":{"title":"Trinquet REST API","version":"1.0.0","description":"\nFor more information about the usage of the Trinquet REST API, see\n[https://github.com/pilotariak/trinquet](https://github.com/pilotariak/trinquet).\n"},"schemes":null,"consumes":["application/json"],"produces":["application/json"],"paths":{"/v1beta/leagues":{"get":{"operationId":"List","responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaGetLeaguesResponse"}}},"summary":"List returns all available League","tags":["LeagueService"]},"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1betaCreateLeagueRequest"}}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaCreateLeagueResponse"}}},"summary":"Create creates a new league","tags":["LeagueService"]}},"/v1beta/tournaments":{"get":{"operationId":"List","responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaGetTournamentsResponse"}}},"summary":"List returns all available Tournament","tags":["TournamentService"]},"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1betaCreateTournamentRequest"}}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaCreateTournamentResponse"}}},"summary":"Create creates a new tournament","tags":["TournamentService"]}},"/v1beta/tournaments/{name}":{"get":{"operationId":"Get","parameters":[{"format":"string","in":"path","name":"name","required":true,"type":"string"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaGetTournamentResponse"}}},"summary":"Get return a tournament","tags":["TournamentService"]}},"/v1ibeta/leagues/{name}":{"get":{"operationId":"Get","parameters":[{"format":"string","in":"path","name":"name","required":true,"type":"string"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v1betaGetLeagueResponse"}}},"summary":"Get return a league","tags":["LeagueService"]}}},"definitions":{"v1betaCreateLeagueRequest":{"properties":{"name":{"format":"string","type":"string"},"website":{"format":"string","type":"string"}},"type":"object"},"v1betaCreateLeagueResponse":{"properties":{"code":{"format":"int32","type":"integer"},"league":{"$ref":"#/definitions/v1betaLeague"}},"type":"object"},"v1betaCreateTournamentRequest":{"properties":{"groups":{"items":{"$ref":"#/definitions/v1betaGroup"},"type":"array"},"name":{"format":"string","type":"string"}},"type":"object"},"v1betaCreateTournamentResponse":{"properties":{"tournament":{"$ref":"#/definitions/v1betaTournament"}},"type":"object"},"v1betaDiscipline":{"properties":{"description":{"format":"string","type":"string"},"id":{"format":"string","type":"string"}},"type":"object"},"v1betaGetLeagueRequest":{"properties":{"name":{"format":"string","type":"string"}},"type":"object"},"v1betaGetLeagueResponse":{"properties":{"league":{"$ref":"#/definitions/v1betaLeague"}},"type":"object"},"v1betaGetLeaguesRequest":{"properties":{"name":{"format":"string","type":"string"}},"type":"object"},"v1betaGetLeaguesResponse":{"properties":{"leagues":{"items":{"$ref":"#/definitions/v1betaLeague"},"type":"array"}},"type":"object"},"v1betaGetTournamentRequest":{"properties":{"name":{"format":"string","type":"string"}},"type":"object"},"v1betaGetTournamentResponse":{"properties":{"tournament":{"$ref":"#/definitions/v1betaTournament"}},"type":"object"},"v1betaGetTournamentsRequest":{"type":"object"},"v1betaGetTournamentsResponse":{"properties":{"tournaments":{"items":{"$ref":"#/definitions/v1betaTournament"},"type":"array"}},"type":"object"},"v1betaGroup":{"properties":{"name":{"format":"string","type":"string"},"team":{"items":{"$ref":"#/definitions/v1betaTeam"},"type":"array"}},"type":"object"},"v1betaLeague":{"description":"League define a pelota league","properties":{"details":{"additionalProperties":{"format":"string","type":"string"},"type":"object"},"disciplines":{"items":{"$ref":"#/definitions/v1betaDiscipline"},"type":"array"},"levels":{"items":{"$ref":"#/definitions/v1betaLevel"},"type":"array"},"name":{"format":"string","type":"string"}},"type":"object"},"v1betaLevel":{"properties":{"description":{"format":"string","type":"string"},"id":{"format":"string","type":"string"}},"type":"object"},"v1betaRound":{"properties":{"first":{"$ref":"#/definitions/v1betaTeam"},"name":{"format":"string","type":"string"},"score":{"format":"string","type":"string"},"second":{"$ref":"#/definitions/v1betaTeam"}},"type":"object"},"v1betaTeam":{"properties":{"name":{"format":"string","type":"string"}},"type":"object"},"v1betaTournament":{"description":"Tournament define a pelota league","properties":{"groups":{"additionalProperties":{"$ref":"#/definitions/v1betaGroup"},"type":"object"},"name":{"format":"string","type":"string"},"rounds":{"items":{"$ref":"#/definitions/v1betaRound"},"type":"array"}},"type":"object"}}}
`
league = `{
  "swagger": "2.0",
  "info": {
    "title": "league.proto",
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
    "/v1beta/leagues": {
      "get": {
        "summary": "List returns all available League",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1betaGetLeaguesResponse"
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
              "$ref": "#/definitions/v1betaCreateLeagueResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1betaCreateLeagueRequest"
            }
          }
        ],
        "tags": [
          "LeagueService"
        ]
      }
    },
    "/v1beta/tournaments": {
      "get": {
        "summary": "List returns all available Tournament",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1betaGetTournamentsResponse"
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
              "$ref": "#/definitions/v1betaCreateTournamentResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1betaCreateTournamentRequest"
            }
          }
        ],
        "tags": [
          "TournamentService"
        ]
      }
    },
    "/v1beta/tournaments/{name}": {
      "get": {
        "summary": "Get return a tournament",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1betaGetTournamentResponse"
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
    },
    "/v1ibeta/leagues/{name}": {
      "get": {
        "summary": "Get return a league",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1betaGetLeagueResponse"
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
    "v1betaCreateLeagueRequest": {
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
    "v1betaCreateLeagueResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "league": {
          "$ref": "#/definitions/v1betaLeague"
        }
      }
    },
    "v1betaCreateTournamentRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaGroup"
          }
        }
      }
    },
    "v1betaCreateTournamentResponse": {
      "type": "object",
      "properties": {
        "tournament": {
          "$ref": "#/definitions/v1betaTournament"
        }
      }
    },
    "v1betaDiscipline": {
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
    "v1betaGetLeagueRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v1betaGetLeagueResponse": {
      "type": "object",
      "properties": {
        "league": {
          "$ref": "#/definitions/v1betaLeague"
        }
      }
    },
    "v1betaGetLeaguesRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v1betaGetLeaguesResponse": {
      "type": "object",
      "properties": {
        "leagues": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaLeague"
          }
        }
      }
    },
    "v1betaGetTournamentRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v1betaGetTournamentResponse": {
      "type": "object",
      "properties": {
        "tournament": {
          "$ref": "#/definitions/v1betaTournament"
        }
      }
    },
    "v1betaGetTournamentsRequest": {
      "type": "object"
    },
    "v1betaGetTournamentsResponse": {
      "type": "object",
      "properties": {
        "tournaments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaTournament"
          }
        }
      }
    },
    "v1betaGroup": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "team": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaTeam"
          }
        }
      }
    },
    "v1betaLeague": {
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
            "$ref": "#/definitions/v1betaLevel"
          }
        },
        "disciplines": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaDiscipline"
          }
        }
      },
      "title": "League define a pelota league"
    },
    "v1betaLevel": {
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
    "v1betaRound": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "first": {
          "$ref": "#/definitions/v1betaTeam"
        },
        "second": {
          "$ref": "#/definitions/v1betaTeam"
        },
        "score": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v1betaTeam": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v1betaTournament": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "groups": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/v1betaGroup"
          }
        },
        "rounds": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaRound"
          }
        }
      },
      "title": "Tournament define a pelota league"
    }
  }
}
`
)

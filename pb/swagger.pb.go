package pb 

const (
api = `{"swagger":"2.0","basePath":"","info":{"title":"Trinquet REST API","version":"1.0.0","description":"\nFor more information about the usage of the Trinquet REST API, see\n[https://github.com/pilotariak/trinquet](https://github.com/pilotariak/trinquet).\n"},"schemes":null,"consumes":["application/json"],"produces":["application/json"],"paths":{"/v1/auth":{"post":{"operationId":"Login","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1LoginRequest"}}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1LoginResponse"}}},"summary":"Perform authentication","tags":["AuthService"]}},"/v1beta/leagues":{"get":{"operationId":"List","parameters":[{"in":"query","name":"name","required":false,"type":"string"}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetLeaguesResponse"}}},"summary":"List returns all available League","tags":["LeagueService"]},"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1betaCreateLeagueRequest"}}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaCreateLeagueResponse"}}},"summary":"Create creates a new league","tags":["LeagueService"]}},"/v1beta/pilotaris":{"get":{"operationId":"List","responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetPilotarisResponse"}}},"summary":"List returns all available Pilotaris","tags":["PilotariService"]}},"/v1beta/pilotris":{"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1betaCreatePilotariRequest"}}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaCreatePilotariResponse"}}},"summary":"Create creates a new Pilotari","tags":["PilotariService"]}},"/v1beta/tournaments":{"get":{"operationId":"List","responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetTournamentsResponse"}}},"summary":"List returns all available Tournament","tags":["TournamentService"]},"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v1betaCreateTournamentRequest"}}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaCreateTournamentResponse"}}},"summary":"Create creates a new tournament","tags":["TournamentService"]}},"/v1beta/tournaments/{name}":{"get":{"operationId":"Get","parameters":[{"in":"path","name":"name","required":true,"type":"string"}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetTournamentResponse"}}},"summary":"Get return a tournament","tags":["TournamentService"]}},"/v1ibeta/leagues/{name}":{"get":{"operationId":"Get","parameters":[{"in":"path","name":"name","required":true,"type":"string"}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetLeagueResponse"}}},"summary":"Get return a league","tags":["LeagueService"]}},"/v1ibeta/pilotaris/{email}":{"get":{"operationId":"Get","parameters":[{"in":"path","name":"email","required":true,"type":"string"}],"responses":{"200":{"description":"A successful response.","schema":{"$ref":"#/definitions/v1betaGetPilotariResponse"}}},"summary":"Get return a Pilotari","tags":["PilotariService"]}}},"definitions":{"v1LoginRequest":{"description":"The request message containing the username and password.","properties":{"password":{"type":"string"},"username":{"type":"string"}},"type":"object"},"v1LoginResponse":{"description":"The response message containing the token.","properties":{"token":{"type":"string"}},"type":"object"},"v1betaCreateLeagueRequest":{"properties":{"name":{"type":"string"},"website":{"type":"string"}},"type":"object"},"v1betaCreateLeagueResponse":{"properties":{"code":{"format":"int32","type":"integer"},"league":{"$ref":"#/definitions/v1betaLeague"}},"type":"object"},"v1betaCreatePilotariRequest":{"properties":{"email":{"type":"string"},"firstname":{"type":"string"},"lastname":{"type":"string"},"phone_number":{"type":"string"}},"type":"object"},"v1betaCreatePilotariResponse":{"properties":{"Pilotari":{"$ref":"#/definitions/v1betaPilotari"}},"type":"object"},"v1betaCreateTournamentRequest":{"properties":{"groups":{"items":{"$ref":"#/definitions/v1betaGroup"},"type":"array"},"name":{"type":"string"}},"type":"object"},"v1betaCreateTournamentResponse":{"properties":{"tournament":{"$ref":"#/definitions/v1betaTournament"}},"type":"object"},"v1betaDiscipline":{"properties":{"description":{"type":"string"},"id":{"type":"string"}},"type":"object"},"v1betaGetLeagueResponse":{"properties":{"league":{"$ref":"#/definitions/v1betaLeague"}},"type":"object"},"v1betaGetLeaguesResponse":{"properties":{"leagues":{"items":{"$ref":"#/definitions/v1betaLeague"},"type":"array"}},"type":"object"},"v1betaGetPilotariResponse":{"properties":{"Pilotari":{"$ref":"#/definitions/v1betaPilotari"}},"type":"object"},"v1betaGetPilotarisResponse":{"properties":{"pilotaris":{"items":{"$ref":"#/definitions/v1betaPilotari"},"type":"array"}},"type":"object"},"v1betaGetTournamentResponse":{"properties":{"tournament":{"$ref":"#/definitions/v1betaTournament"}},"type":"object"},"v1betaGetTournamentsResponse":{"properties":{"tournaments":{"items":{"$ref":"#/definitions/v1betaTournament"},"type":"array"}},"type":"object"},"v1betaGroup":{"properties":{"name":{"type":"string"},"team":{"items":{"$ref":"#/definitions/v1betaTeam"},"type":"array"}},"type":"object"},"v1betaLeague":{"description":"League define a pelota league","properties":{"details":{"additionalProperties":{"type":"string"},"type":"object"},"disciplines":{"items":{"$ref":"#/definitions/v1betaDiscipline"},"type":"array"},"levels":{"items":{"$ref":"#/definitions/v1betaLevel"},"type":"array"},"name":{"type":"string"}},"type":"object"},"v1betaLevel":{"properties":{"description":{"type":"string"},"id":{"type":"string"}},"type":"object"},"v1betaPilotari":{"properties":{"email":{"type":"string"},"firstname":{"type":"string"},"lastname":{"type":"string"},"phone_number":{"type":"string"}},"type":"object"},"v1betaRound":{"properties":{"first":{"$ref":"#/definitions/v1betaTeam"},"name":{"type":"string"},"score":{"type":"string"},"second":{"$ref":"#/definitions/v1betaTeam"}},"type":"object"},"v1betaTeam":{"properties":{"name":{"type":"string"}},"type":"object"},"v1betaTournament":{"description":"Tournament define a pelota league","properties":{"groups":{"additionalProperties":{"$ref":"#/definitions/v1betaGroup"},"type":"object"},"name":{"type":"string"},"rounds":{"items":{"$ref":"#/definitions/v1betaRound"},"type":"array"}},"type":"object"}}}
`
auth = `{
  "swagger": "2.0",
  "info": {
    "title": "auth.proto",
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
    "/v1/auth": {
      "post": {
        "summary": "Perform authentication",
        "operationId": "Login",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1LoginResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1LoginRequest"
            }
          }
        ],
        "tags": [
          "AuthService"
        ]
      }
    }
  },
  "definitions": {
    "v1LoginRequest": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "description": "The request message containing the username and password."
    },
    "v1LoginResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      },
      "description": "The response message containing the token."
    }
  }
}
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
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1betaGetLeaguesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "LeagueService"
        ]
      },
      "post": {
        "summary": "Create creates a new league",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "A successful response.",
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
    "/v1ibeta/leagues/{name}": {
      "get": {
        "summary": "Get return a league",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "A successful response.",
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
            "type": "string"
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
          "type": "string"
        },
        "website": {
          "type": "string"
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
    "v1betaDiscipline": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "title": {
          "type": "string"
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
    "v1betaLeague": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "details": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
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
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}
`
pilotari = `{
  "swagger": "2.0",
  "info": {
    "title": "pilotari.proto",
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
    "/v1beta/pilotaris": {
      "get": {
        "summary": "List returns all available Pilotaris",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1betaGetPilotarisResponse"
            }
          }
        },
        "tags": [
          "PilotariService"
        ]
      }
    },
    "/v1beta/pilotris": {
      "post": {
        "summary": "Create creates a new Pilotari",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1betaCreatePilotariResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1betaCreatePilotariRequest"
            }
          }
        ],
        "tags": [
          "PilotariService"
        ]
      }
    },
    "/v1ibeta/pilotaris/{email}": {
      "get": {
        "summary": "Get return a Pilotari",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1betaGetPilotariResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "email",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "PilotariService"
        ]
      }
    }
  },
  "definitions": {
    "v1betaCreatePilotariRequest": {
      "type": "object",
      "properties": {
        "firstname": {
          "type": "string"
        },
        "lastname": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "phone_number": {
          "type": "string"
        }
      }
    },
    "v1betaCreatePilotariResponse": {
      "type": "object",
      "properties": {
        "Pilotari": {
          "$ref": "#/definitions/v1betaPilotari"
        }
      }
    },
    "v1betaGetPilotariResponse": {
      "type": "object",
      "properties": {
        "Pilotari": {
          "$ref": "#/definitions/v1betaPilotari"
        }
      }
    },
    "v1betaGetPilotarisResponse": {
      "type": "object",
      "properties": {
        "pilotaris": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaPilotari"
          }
        }
      }
    },
    "v1betaPilotari": {
      "type": "object",
      "properties": {
        "firstname": {
          "type": "string"
        },
        "lastname": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "phone_number": {
          "type": "string"
        }
      }
    }
  }
}
`
tournament = `{
  "swagger": "2.0",
  "info": {
    "title": "tournament.proto",
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
    "/v1beta/tournaments": {
      "get": {
        "summary": "List returns all available Tournament",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "A successful response.",
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
            "description": "A successful response.",
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
            "description": "A successful response.",
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
            "type": "string"
          }
        ],
        "tags": [
          "TournamentService"
        ]
      }
    }
  },
  "definitions": {
    "v1betaCreateTournamentRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
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
    "v1betaGetTournamentResponse": {
      "type": "object",
      "properties": {
        "tournament": {
          "$ref": "#/definitions/v1betaTournament"
        }
      }
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
          "type": "string"
        },
        "team": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1betaTeam"
          }
        }
      }
    },
    "v1betaRound": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "first": {
          "$ref": "#/definitions/v1betaTeam"
        },
        "second": {
          "$ref": "#/definitions/v1betaTeam"
        },
        "score": {
          "type": "string"
        }
      }
    },
    "v1betaTeam": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "v1betaTournament": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
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

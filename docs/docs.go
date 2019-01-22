// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-22 11:04:38.7440225 +0800 CST m=+0.039893701

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "WebAPI v0.0.1",
        "title": "WebAPI",
        "contact": {},
        "license": {},
        "version": "0.0.1"
    },
    "host": "{{.Host}}",
    "basePath": "/",
    "paths": {
        "/Data/GetAccountInfo": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户"
                ],
                "summary": "获取账户信息",
                "responses": {
                    "200": {
                        "description": "{\"Code\":0,\"Msg\":\"\",\"Body\":\"msg\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/Login": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账号"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userName",
                        "name": "userName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userPwd",
                        "name": "userPwd",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"Code\":0,\"Msg\":\"\",\"Body\":\"abc\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

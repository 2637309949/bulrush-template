define({ "api": [
  {
    "type": "get",
    "url": "/chache/xxx",
    "title": "测试缓存路由",
    "group": "Cache",
    "description": "<p>xxxbbb</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "profiles",
            "description": "<p>List of user profiles.</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "profiles.age",
            "description": "<p>Users age.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "profiles.image",
            "description": "<p>Avatar-Image.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n   \"message\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "accessToken",
            "description": "<p>令牌</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ids",
            "description": "<p>顶级评分项ID, 如果多个就用用&quot;,&quot;分割</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "label",
            "description": "<p>顶级评分项label, 如果多个就用用&quot;,&quot;分割</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./routes/hello.go",
    "groupTitle": "Cache",
    "name": "GetChacheXxx",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/chache/xxx"
      }
    ]
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_home_double_Work_K11_repo_bulrush_bulrush_template_doc_main_js",
    "groupTitle": "_home_double_Work_K11_repo_bulrush_bulrush_template_doc_main_js",
    "name": ""
  }
] });

define({ "api": [
  {
    "type": "get",
    "url": "/gorm/mock/init",
    "title": "GORM测试数据",
    "group": "Test",
    "description": "<p>队列路由</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "Mess",
            "description": "<p>实体类</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Mess.message",
            "description": "<p>消息内容</p>"
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
    "version": "0.0.0",
    "filename": "./routes/test_mock.go",
    "groupTitle": "Test",
    "name": "GetGormMockInit",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/gorm/mock/init"
      }
    ]
  },
  {
    "type": "get",
    "url": "/gorm/mock/login",
    "title": "模拟登录",
    "group": "Test",
    "description": "<p>队列路由</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "Mess",
            "description": "<p>实体类</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Mess.message",
            "description": "<p>消息内容</p>"
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
    "version": "0.0.0",
    "filename": "./routes/test_mock.go",
    "groupTitle": "Test",
    "name": "GetGormMockLogin",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/gorm/mock/login"
      }
    ]
  },
  {
    "type": "get",
    "url": "/test/chache",
    "title": "缓存路由",
    "group": "Test",
    "description": "<p>缓存路由</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "Mess",
            "description": "<p>实体类</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Mess.message",
            "description": "<p>消息内容</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n   \"message\": \"hello\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/test_cache.go",
    "groupTitle": "Test",
    "name": "GetTestChache",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/test/chache"
      }
    ]
  },
  {
    "type": "get",
    "url": "/test/mgo/adduser",
    "title": "添加用户",
    "group": "Test",
    "description": "<p>添加用户</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "Users",
            "description": "<p>实体类数组</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Users.Name",
            "description": "<p>名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Users.Password",
            "description": "<p>密码</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n    \"ID\": \"5d064cc213756cdb0f662607\",\n    \"Created\": null,\n    \"Modified\": null,\n    \"Deleted\": null,\n    \"CreatorID\": \"5d064cc2e8cd7d3029885465\",\n    \"Creator\": null,\n    \"ModifierID\": \"5d064cc2e8cd7d3029885466\",\n    \"Modifier\": null,\n    \"DeleterID\": \"\",\n    \"Deleter\": null,\n    \"Name\": \"double\",\n    \"Password\": \"111111\",\n    \"Age\": 24,\n    \"RoleIds\": null,\n    \"Roles\": [],\n    \"Test1\": null,\n    \"Test2\": null\n}",
          "type": "json"
        }
      ]
    },
    "parameter": {
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"cond\": { \"age\": { $gte: 1 } },\n  \"range\": \"page\",\n  \"page\": 1,\n  \"size\": 10,\n  \"project\": \"_id,name\",\n  \"preload\": \"Creator\",\n  \"order\": \"-_created\",\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/test_mgo.go",
    "groupTitle": "Test",
    "name": "GetTestMgoAdduser",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/test/mgo/adduser"
      }
    ]
  },
  {
    "type": "get",
    "url": "/test/mq/hello",
    "title": "队列路由",
    "group": "Test",
    "description": "<p>队列路由</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "Mess",
            "description": "<p>实体类</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Mess.message",
            "description": "<p>消息内容</p>"
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
    "version": "0.0.0",
    "filename": "./routes/test_mq.go",
    "groupTitle": "Test",
    "name": "GetTestMqHello",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/test/mq/hello"
      }
    ]
  },
  {
    "type": "get",
    "url": "/test/seq/users",
    "title": "查询用户",
    "group": "Test",
    "description": "<p>查询用户</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "Users",
            "description": "<p>实体类数组</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Users.Name",
            "description": "<p>名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Users.Password",
            "description": "<p>密码</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n    \"ID\": 1,\n    \"CreatedAt\": \"2019-07-12T20:59:57+08:00\",\n    \"UpdatedAt\": \"2019-07-12T20:59:57+08:00\",\n    \"DeletedAt\": null,\n    \"Creator\": null,\n    \"CreatorID\": 0,\n    \"Modifier\": null,\n    \"ModifierID\": 0,\n    \"Deleter\": null,\n    \"DeleterID\": 0,\n    \"Name\": \"L1212\",\n    \"Age\": 23\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/test_gorm.go",
    "groupTitle": "Test",
    "name": "GetTestSeqUsers",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/test/seq/users"
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
    "group": "_home_double_work_K11_repo_bulrush_bulrush_template_doc_main_js",
    "groupTitle": "_home_double_work_K11_repo_bulrush_bulrush_template_doc_main_js",
    "name": ""
  }
] });

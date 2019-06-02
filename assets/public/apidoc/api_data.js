define({ "api": [
  {
    "type": "get",
    "url": "/chache",
    "title": "缓存测试",
    "group": "Test",
    "description": "<p>测试系统可用性</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "accessToken",
            "description": "<p>认证令牌</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n       \"message\":    \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "routes/hello.go",
    "groupTitle": "Test",
    "name": "GetChache"
  },
  {
    "type": "get",
    "url": "/ping",
    "title": "数据库测试",
    "group": "Test",
    "description": "<p>测试系统可用性</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "accessToken",
            "description": "<p>认证令牌</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n       \"message\":    \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "routes/hello.go",
    "groupTitle": "Test",
    "name": "GetPing"
  }
] });

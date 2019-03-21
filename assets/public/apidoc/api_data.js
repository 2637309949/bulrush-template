define({ "api": [
  {
    "type": "get",
    "url": "/hello",
    "title": "检验Token有效性",
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
    "name": "GetHello"
  }
] });

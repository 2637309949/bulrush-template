define({ "api": [
  {
    "type": "get",
    "url": "/chache",
    "title": "测试缓存路由",
    "group": "Cache",
    "description": "<p>DEMO: /chache</p>",
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
    "success": {
      "examples": [
        {
          "title": "正常返回",
          "content": "HTTP/1.1 200 OK\n{\n   \"message\": \"ok\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "routes/hello.go",
    "groupTitle": "Cache",
    "name": "GetChache",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8080/api/v1/chache"
      }
    ]
  }
] });

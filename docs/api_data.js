define({
  "api": [
    {
      "type": "post",
      "url": "/ts/daily/gz/scoring/results",
      "title": "保存GZ评分结果",
      "group": "TenantScore",
      "description": "<p>list里的结构跟/ts/daily/gz/scoring/:ids基本一致,多添加的字段是store_id,tenant_type,answer</p>",
      "parameter": {
        "fields": {
          "Parameter": [
            {
              "group": "Parameter",
              "type": "String",
              "optional": false,
              "field": "site",
              "description": "<p>所在项目</p>"
            },
            {
              "group": "Parameter",
              "type": "String",
              "optional": false,
              "field": "stat_time",
              "description": "<p>提交的评分所在月份, 跨月的需要多次提交</p>"
            },
            {
              "group": "Parameter",
              "type": "Array",
              "optional": false,
              "field": "list",
              "description": "<p>评分结果</p>"
            },
            {
              "group": "Parameter",
              "type": "String",
              "optional": false,
              "field": "score_editor",
              "description": "<p>提交人ID</p>"
            }
          ]
        },
        "examples": [
          {
            "title": "Request-Example:",
            "content": "{\n    \"site\": \"2\",\n    \"stat_time\": 1530374400,\n    \"score_editor\": \"super\",\n    \"list\": [\n        {\n            \"_id\": \"1ce4fcafac2011e886c2c12334fc5470\",             // 父级评分项ID\n            \"item_label\": \"配合商场运营（陈列、现场营业管理）\",       // 父级评分项标签\n            \"tenant_type\": \"new\",                                  // 租户类型，必选，从查询可评分店铺接口获取, new: \"新租户\", old: \"老租户\"\n            \"store_id\": \"cff001e4f44a11e7bf7814dae9eb7407\",        // 店铺ID\n            \"children\": [{\n                    \"item_id\": \"1ce4fcb0ac2011e886c2c12334fc5470\",                                         // 子级评分项ID\n                    \"item_label\": \"是否遵循K11的Retail Therapy-陈列、Rank标识、盐的故事、一店一物等\",          // 子级评分项标签\n                    \"score_time\": 1530373400,                                                              // 具体评分时间（必须在stat_time月份之内）\n                    \"comment\": \"xxx\",                                                                      // 备注\n                    \"answer\": 4,                                                                           // 答案, 答案为Y/N/或分值\n                    \"type\": \"加分\",                                                                        // 选项类型\n                    \"selected\": \"是的\",                                                                    // 选项\n                    \"attachment\": [{                                                                       // 附件\n                        \"status\": \"done\",\n                        \"name\": \"测试\",\n                        \"uid\": \"xxxx\",\n                        \"url\": \"/www/www/xx/x.jpg\"\n                    }]\n            }]\n        }\n    ]\n}",
            "type": "json"
          }
        ]
      },
      "version": "0.0.0",
      "filename": "src/routes/gz_score.js",
      "groupTitle": "TenantScore",
      "name": "PostTsDailyGzScoringResults",
      "sampleRequest": [
        {
          "url": "https://k11-central.hofo.co/api/v1/ts/daily/gz/scoring/results"
        }
      ]
    },
  ]
});

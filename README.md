# text2lang

简单的文本语言识别API,使用GIN

## 支持的语言

日语、中文、英语、韩语、阿拉伯语、俄语、泰语、印地语、希伯来语、高棉语、老挝语
## 使用


clone存储库安装go并cd进目录后启动服务器:
```bash
go mod tidy
go run main.go
```

使用Docker:
```bash
docker build -t text2lang .
docker run -p 8080:8080 text2lang
```

部署到koyeb:  

[![Deploy to Koyeb](https://www.koyeb.com/static/images/deploy/button.svg)](https://app.koyeb.com/deploy?name=text2lang&type=git&repository=IcyBlue17%2Ftext2lang&branch=main&builder=dockerfile&instance_type=free&regions=fra&instances_min=0&autoscaling_sleep_idle_delay=3900&ports=7891%3Bhttp%3B%2F&hc_protocol%5B7891%5D=tcp&hc_grace_period%5B7891%5D=5&hc_interval%5B7891%5D=30&hc_restart_limit%5B7891%5D=3&hc_timeout%5B7891%5D=5&hc_path%5B7891%5D=%2F&hc_method%5B7891%5D=get)

## API使用

post到`http://你的地址:7891/`

请求:
```json
{"text": "こんにちは"}
```

返回:
```json
{"success": true, "code": 200, "lang": "ja"}
```

## 原理

统计文本中各语言字符数量，返回占比最高的语言的代码

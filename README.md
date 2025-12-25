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

## API使用

post到`http://你的地址:8080/`

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

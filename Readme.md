# sub2sing-box

## Cli Command

### convert

```
Convert common proxy to sing-box proxy

Usage:
   convert [flags]

Flags:
  -d, --delete string           delete proxy with regex
  -h, --help                    help for convert
  -o, --output string           output file path
  -p, --proxy strings           common proxies
  -r, --rename stringToString   rename proxy with regex (default [])
  -s, --subscription strings    subscription urls
  -t, --template string         template file path
```

### server

```
Run the server

Usage:
   server [flags]

Flags:
  -h, --help          help for server
  -p, --port uint16   server port (default 8080)
```

#### api

##### GET /convert

- `data`: Base64 编码的 JSON 字符串，包含以下字段：
  - `subscription`: []string
  - `proxy`: []string
  - `delete`: string 可选
  - `rename`: string 可选
  - `template`: map[string]string 可选

示例
```
{
  "subscription": ["url1", "url2"],
  "proxy": ["p1", "p2"],
  "delete": "reg",
  "template": "t",
  "rename": {
    "text": "replaceTo"
  }
}
```

## Template

Template 中使用 `<all-proxy-tags>` 指明节点插入位置，例如

```
{
  "type": "selector",
  "tag": "节点选择",
  "outbounds": ["<all-proxy-tags>", "direct"],
  "interrupt_exist_connections": true
},
```

## Docker

`docker run -p 8080:8080 nite07/sub2sing-box`
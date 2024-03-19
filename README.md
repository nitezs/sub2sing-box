# sub2sing-box

## cli

- convert: 转换
- server: 启动 Web UI
- version: 版本信息

`sub2sing-box <command> -h` 查看帮助

## api

##### GET /convert

- `data`: Base64 编码(url safe)的 JSON 字符串，包含以下字段：
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

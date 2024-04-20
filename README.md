> [!important]
> [moved to forgejo](https://git.nite07.com/nite07/sub2sing-box)# sub2sing-box
## cli

- convert: 转换
- server: 启动 Web UI
- version: 版本信息

`sub2sing-box <command> -h` 查看帮助

## api

### GET /convert?data=xxx

data 为 base64 URL 编码的请求体，示例

```
{
  "subscriptions": ["订阅地址1", "订阅地址2"],
  "proxies": ["代理1", "代理2"],
  "template": "模板路径",
  "delete": "",
  "rename": {"原文本": "新文本"},
  "group": false,
  "group-type": "selector",
  "sort": "name",
  "sort-type": "asc"
}
```

## Template 占位符

- `<all-proxy-tags>`: 插入所有节点标签
  ```
  {
    "type": "selector",
    "tag": "节点选择",
    "outbounds": ["<all-proxy-tags>",   "direct"],
    "interrupt_exist_connections": true
  }
  ```
- `<all-country-tags>`: 插入所有国家标签
  ```
  {
    "type": "selector",
    "tag": "节点选择",
    "outbounds": ["<all-country-tags>", "direct"],
    "interrupt_exist_connections": true
  }
  ```
- `<国家(地区)二字码>`: 插入国家(地区)所有节点标签，例如 `<tw>`
  ```
  {
    "type": "selector",
    "tag": "巴哈姆特",
    "outbounds": ["<tw>", "direct"],
    "interrupt_exist_connections": true
  }
  ```

## Docker

`docker run -p 8080:8080 nite07/sub2sing-box`

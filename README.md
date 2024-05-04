# sub2sing-box

## 控制台命令

- convert: 转换
- server: 启动 Web UI
- version: 版本信息

`sub2sing-box <command> -h` 查看帮助

## Api

### GET /convert?data=xxx

data 为 base64 URL 编码(将编码字符串中`/`替换为`_`，`+`替换为`-`)的 JSON 结构体，示例

```
{
  "subscriptions": ["订阅地址1", "订阅地址2"],
  "proxies": ["代理1", "代理2"],
  "template": "模板路径/网络地址",
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

# sub2sing-box

将订阅/节点连接转换为 sing-box 配置的工具。

## 控制台命令

使用 `sub2sing-box <command> -h` 查看各命令的帮助信息。

## 配置

示例:

```json
{
  "subscriptions": ["订阅地址1", "订阅地址2"],
  "proxies": ["代理1", "代理2"],
  "template": "模板路径或网络地址",
  "delete": "",
  "rename": { "原文本": "新文本" },
  "group-type": "selector",
  "sort": "name",
  "sort-type": "asc"
}
```

将上述 JSON 内容保存为 `config.json`，然后执行:

```
sub2sing-box convert -c ./config.json
```

即可生成 sing-box 配置，无需每次重复设置参数。

## 模板

### 默认模板

默认模板位于 `templates` 目录，使用 `tun+fakeip` 配置，可以根据需求自行修改模板内容。

### 占位符

模板中可使用以下占位符:

- `<all-proxy-tags>`: 插入所有节点标签
- `<all-country-tags>`: 插入所有国家标签
- `<国家(地区)二字码>`: 插入指定国家(地区)所有节点标签，如 `<tw>`

占位符使用示例:

```json
{
  "type": "selector",
  "tag": "节点选择",
  "outbounds": ["<all-proxy-tags>", "direct"],
  "interrupt_exist_connections": true
}

{
  "type": "selector",
  "tag": "节点选择",
  "outbounds": ["<all-country-tags>", "direct"],
  "interrupt_exist_connections": true
}

{
  "type": "selector",
  "tag": "巴哈姆特",
  "outbounds": ["<tw>", "direct"],
  "interrupt_exist_connections": true
}
```

## Docker 使用

```
docker run -p 8080:8080 nite07/sub2sing-box
```

可以挂载目录添加自定义模板

## Server 模式 API

### GET /convert

| query | 描述                                                                                                                    |
| ----- | ----------------------------------------------------------------------------------------------------------------------- |
| data  | 同上方配置，但需要使用 [base64 URL safe 编码](<https://gchq.github.io/CyberChef/#recipe=To_Base64('A-Za-z0-9%2B/%3D')>) |

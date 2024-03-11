# sub2sing-box

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


打开 `Preferences->Docker Engine`

![](./assets/docker%20-%20change%20proxies%20under%20Mac.png)

```json
{
  "debug": true,
  "experimental": false,
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirrors.c.163.com"
  ]
}
```

点击 Apply/Restart
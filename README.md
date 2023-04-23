# README

这是一个学习 Discord Bot 的应用。

## 教程

[[How To] Building a Simple Discord Bot using DiscordGo](https://www.youtube.com/watch?v=G7A3nnMvfCk)

## 环境变量

main.go 中所调用的环境变量来自于环境变量文件 `.env`，因为涉及隐私信息，就没有上传到 github 中，需要自己创建。  
不同的 OS 读取和书写 `.env` 文件的方式不一样。

※ 环境变量中所用到的 token 的获取方式跟着视频做一下就知道了。

### linux / mac

```env
# .env
import VARIABLE_NAME=variable
```

```shell
source .env
```

### windows

```env
# .env
VARIABLE_NAME=variable
```

```powershell
./source-env.ps1
```

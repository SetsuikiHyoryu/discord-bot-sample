# 在 PowerShell 中获取 `.env` 文件中的环境变量
# `.env` 文件中的环境变量需要是 `key=value` 的键值对
get-content .env | foreach {
  $name, $value = $_.split('=')
  set-content env:\$name $value
}


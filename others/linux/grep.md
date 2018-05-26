
```bash
# grep 用于在文本中执行关键词搜索，并显示匹配的结果
grep [OPTION] file_path
# -b 将可执行文件当作文本文件搜索
# -c 仅显示找到的行数
# -i 忽略大小写
# -n 显示行号
# -v 反向选择 —— 仅列出没有“关键词”的行

# 譬如，查找当前系统种不允许登录的用户信息
grep /sbin/nologin /etc/passwd

```


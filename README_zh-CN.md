[English](https://github.com/HaoKunT/filetools/blob/master/README.md)|[中文](https://github.com/HaoKunT/filetools/blob/master/README_zh-CN.md)
# filetools
一些关于文件的小工具集合

## 为什么写这个工具
出于学习golang和替换一些文件工具的目的。`filetools`工具将一些零散的文件命令（例如`cp`, `mv`）集成为一个命令，并以子命令的形式对文件命令进行区分。

## Build
`build-release.sh`是一个shell脚本，脚本将生成不同架构和平台的二进制文件
``` bash
./build-release.sh
```
二进制文件将放在`release`目录下

## 使用方法
``` bash
./filetools -h
```

## 现在已有的命令
- `compare`: 比较两个目录的不同，也可用于比较文件
- `copy`: 将一个文件或目录拷贝至其他地方，这个命令可以有进度条
- `info`: 输出某文件或目录的一些信息
- `rename`: 重命名文件，这个命令可以保持原有文件
- `list`: 列出目录下的所有文件，可以排序，限制输出的数量，这个命令可以有效的查找目录下的大文件，如果你硬盘空间不够了，可以找一下哪些文件较大，自行决定是否删除

## 国际化帮助
filetools工具现已使用`github.com/chai2010/gettext-go`支持gettext，翻译文件在`local.zip`，使用`github.com/go-bindata/go-bindata`将翻译文件打包进二进制文件，`local_pack.go`文件是由`go-bindata`自动生成的，请不要编辑。

如果你想制作你自己的本地化文件，请按以下步骤操作
1. 编辑本地文件，例如`local/zh_CN/LC_MESSAGES/filetools.po`
2. 使用`msg -o local/zh_CN/LC_MESSAGES/filetools.mo local/zh_CN/LC_MESSAGES/filetools.po`命令构建MO文件
3. 压缩`local`文件夹到`local.zip`文件夹
4. 使用`go-bindata -o=local_pack.go -pkg=main local.zip`命令生成go文件。（你需要先安装go-bindata）
5. 构建二进制文件

## todo
未来将会增加更多的子命令和参数项
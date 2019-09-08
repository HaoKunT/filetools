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

## todo
未来将会增加更多的子命令和参数项
[English](https://github.com/HaoKunT/filetools/blob/master/README.md)|[中文](https://github.com/HaoKunT/filetools/blob/master/README_zh-CN.md)
# filetools 
--------------------------------
[![Build Status](https://travis-ci.org/HaoKunT/filetools.svg?branch=master)](https://travis-ci.org/HaoKunT/filetools)

some tools of file

## Why
For practice and replace so many different tools, the `filetools` will combine more features about file into one command.

## Build
`build-release.sh` is a shell script to build the release, it can be all different architectures and platforms binary files
``` bash
./build-release.sh
```
The release file will be build in `release` directory

## Usage
``` bash
./filetools -h
```

## Now command
- `compare`: Compare the different between two directories, it also can comapre the files
- `copy`: Copy the files or directories to other path, it has progressBar
- `info`: Output the information about a file or directory
- `rename`: Rename a file, it can keep the original file
- `list`: List files in a directory, it can sort files, limit the number of files, so it's useful to find the large files in a directory. if you find your disk space is not enough, you can findout which files are large and determine wether it can be deleted.
- `delete`: Delete a file or directory, this command is soft delete by default, it means we will move the file to the ~/.ft-trash directory, if you accidentally delete a file, you can recover it from the trash. We don't move the file into the cycle on Windows because it need use cgo, but I don't want it. You also can delete it directly, pass `-hd` to hard delete it. In future, we will use a flag to clean the trash.
- `move`: Move a file or directory, it may work like `mv`, but we are not satisfied with it, you should use it be careful. Of course, the normal condition is ok, I means there are some extremely case not been tested, if you find any bugs on it, please tell me.

## i18n
filetools support gettext by use `github.com/chai2010/gettext-go`, the local file is `local.zip`, use the new feature in golang1.16 to embed the local file into binary file, so you need update your golang version to 1.16 or higher

if you want to make your local file, please follow the steps below
1. Edit the local file, like `local/zh_CN/LC_MESSAGES/filetools.po`
2. Use command `msgfmt -o local/zh_CN/LC_MESSAGES/filetools.mo local/zh_CN/LC_MESSAGES/filetools.po` to build MO file
3. zip the `local` directory into `local.zip`
4. Build the binary file

The step `2, 3` can run `build-local.sh` directly(need `msgfmt`, `zip` command and bash environment)

## todo
More command and flag will be add in the future

- [ ] add file compression and unzip.
- [ ] add file download.
- [ ] add file found.
- [ ] add regex match.
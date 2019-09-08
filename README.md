[English](https://github.com/HaoKunT/filetools/blob/master/README.md)|[中文](https://github.com/HaoKunT/filetools/blob/master/README_zh-CN.md)
# filetools
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

## todo
More command and flag will be add in the future

# VNDR

[![Build Status](https://travis-ci.org/LK4D4/vndr.svg?branch=master)](https://travis-ci.org/LK4D4/vndr)

Vndr is simple vendoring tool, which is inspired by Docker vendor script.
Vndr has next command line arguments:

* `-verbose` adds additional output, helpful for debugging issues.
* `-whitelist` allows you to specify several regular expressions for paths
  which will *not* be cleaned in the final stage of vendoring -- this is useful
  for running tests in a vendored project or otherwise ensuring that some
  important files are retained after `vndr` is done cleaning unused files from
  your `vendor/` directory.
* `-strict` exits with non-zero status on non-trivial warning

* `-copyFromLocal` 从本地`GOPATH`下复制库到项目`vendor`下

## 痛点以及改进点
* 在开发k8s 项目时，版本依赖复杂，需要忽略一些 repo 单独管理
* 各个包管理器对于ignore 支持很差，glide 在update 后会删除ignore 的文件
* vndr 默认从外网下载包，添加了可以从本地GOPATH 拷贝依赖包的功能
* 增加两种忽略策略，既可以忽略本地项目的某一个文件夹，也可以忽略指定的repo
* **使用前如果使用ignore 功能，请先创建vndr.ignore文件**

## Installation

Execute

    go get github.com/xianlubird/vndr
    
## vndr.ignore
`vndr.ignore` 是用来描述需要忽略的包和路径，书写规范如下
```
ignoreFolders:
  - helm
ignorePaths:
  - github.com/kubernetes/kompose
```
`ignoreFolders`是指忽略本项目某一个文件夹，该文件夹下的所有import 都会被忽略


`ignorePaths` 是指忽略某个依赖，该`repo`以及子依赖都会被忽略

## vendor.conf

`vendor.conf` is the configuration file of `vndr`. It must have multiple lines,
which have format:
```
Import path               | revision                               | Repository(optional)
github.com/example/example 03a4d9dcf2f92eae8e90ed42aa2656f63fdd0b14 https://github.com/LK4D4/example.git

```
You can use `Repository` field for vendoring forks instead of original repos.
This config format is also accepted by [trash](https://github.com/rancher/trash).

## Initialization

You can initiate your project with vendor directory and `vendor.conf` using command
`vndr init`. This will populate your vendor directory with latest versions of
all dependencies and also write `vendor.conf` config which you can use for changing
versions later.

## Updating or using existing vendor.conf

If you already have `vendor.conf` you can just change versions there as you like,
set `$GOPATH` and run `vndr` in your repository with that file:
```
vndr
```
(Note: Your repository must be in proper place in `$GOPATH`, i.e. `$GOPATH/src/github.com/LK4D4/vndr`).

Also it's possible to vendor or update only one dependency:
```
vndr github.com/example/example 03a4d9dcf2f92eae8e90ed42aa2656f63fdd0b14 https://github.com/LK4D4/example.git
```
or
```
vndr github.com/example/example
```
to take revision and repo from `vendor.conf`.

If you experience any problems, please try to run `vndr -verbose`.

Sometimes `vndr` might suggest you to change your `vendor.conf`:
* in case of duplicated or non-top packages it will write suggested file to
`vendor.conf.tmp`, you should diff your file with it and make changes accordingly.
* in case of unused packages it will just print warning

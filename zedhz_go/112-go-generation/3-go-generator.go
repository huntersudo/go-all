package main

import "fmt"

/**
要玩 Go 的代码生成，你需要三个东西：
一个函数模板，在里面设置好相应的占位符；
一个脚本，用于按规则来替换文本并生成新的代码；
一行注释代码。

todo 可以看到，函数模板中我们有如下的占位符：
PACKAGE_NAME：包名
GENERIC_NAME  ：名字
GENERIC_TYPE  ：实际的类型
 */

/**
然后，我们有一个叫gen.sh的生成脚本，如下所示：

#!/bin/bash

set -e

SRC_FILE=${1}
PACKAGE=${2}
TYPE=${3}
DES=${4}
#uppcase the first char
PREFIX="$(tr '[:lower:]' '[:upper:]' <<< ${TYPE:0:1})${TYPE:1}"

DES_FILE=$(echo ${TYPE}| tr '[:upper:]' '[:lower:]')_${DES}.go

sed 's/PACKAGE_NAME/'"${PACKAGE}"'/g' ${SRC_FILE} | \
    sed 's/GENERIC_TYPE/'"${TYPE}"'/g' | \
    sed 's/GENERIC_NAME/'"${PREFIX}"'/g' > ${DES_FILE}
 */

/**
这里需要 4 个参数：
模板源文件；
包名；
实际需要具体化的类型；
用于构造目标文件名的后缀。

然后，我们用 sed 命令去替换刚刚的函数模板，并生成到目标文件中（关于 sed 命令，我给你推荐一篇文章：《sed 简明教程》）。生成代码
 */


//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample() {
	var s string = "Hello"
	c := NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateExample: %s (%T)\n", v, v)
}

// todo 结果：

[![logo](./logo.png)](https://github.com/phachon/fastgo)

[![build](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/phachon/fastgo)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/phachon/fastgo)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phachon/fastgo/master/LICENSE)
[![go_Report](https://goreportcard.com/badge/github.com/phachon/fastgo)](https://goreportcard.com/report/github.com/phachon/fastgo)
[![release](https://img.shields.io/github/release/phachon/fastgo.svg?style=flat)](https://github.com/phachon/fastgo/releases) 
[![powered_by](https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat)]()
[![platforms](https://img.shields.io/badge/platform-All-yellow.svg?style=flat)]()

fastgo 是一个基于 [fasthttp](https://github.com/valyala/fasthttp), [fasthttprouter](https://github.com/buaazp/fasthttprouter), [fasthttpsession](https://github.com/phachon/fasthttpsession) 搭建的快速的 Go Web 开发框架!


# 安装

要求是 Go 至少是 v1.7。

```shell
$ go get -u github.com/phachon/fastgo
$ go get ./...
```

# 使用
```go
package main

import (
    "github.com/phachon/fastgo"
)


type AuthorController struct{
	fastgo.Controller
}

func (this *AuthorController) Index()  {
    
	this.ReturnJson(map[string]string{
		"data": "ok",
	})
}

func main() {
    author := &AuthorController{}
    fastgo.Route.GET("/author/index", author, "Index")

    fastgo.Run()
}

```

# 文档

文档地址: [http://godoc.org/github.com/phachon/fastgo](http://godoc.org/github.com/phachon/fastgo)


## 示例

[完整的 MVC 登录示例](./_example/mvc)

# 组件
- [fasthttp](https://github.com/valyala/fasthttp)
- [fasthttprouter](https://github.com/buaazp/fasthttprouter)
- [fasthttpsession](https://github.com/phachon/fasthttpsession)
- [go-logger](https://github.com/phachon/go-logger)

## 反馈

- 如果您喜欢该项目，请 [Start](https://github.com/phachon/fastgo/stargazers).
- 如果在使用过程中有任何问题， 请提交 [Issue](https://github.com/phachon/fastgo/issues).
- 如果您发现并解决了bug，请提交 [Pull Request](https://github.com/phachon/fastgo/pulls).
- 如果您想扩展 session 存储，欢迎 [Fork](https://github.com/phachon/fastgo/network/members) and merge this rep.
- 如果你想交个朋友，欢迎发邮件给 [phachon@163.com](mailto:phachon@163.com).

## License

MIT

Thanks
---------
Create By phachon@163.com

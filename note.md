# 搭建博客过程的一些记录
## lecture 1
介绍了一些要用到的框架，以及要养成看官方文档的习惯，如何快速的去了解、去查找自己想要的东西？

这正和我反思的自己看东西的问题非常相关，进入一个网站（例如官方文档），看到庞大的内容量瞬间就不想看了，或者就看的特别随意。

但实际上看东西是有技巧的，先看看界面上有些什么（以https://beego.me/举例），
对于新使用者而言，最先看应该是快速入门，看看这个东西怎么用起来，也就是程序的hello world，
然后对于稍微深一点的东西，参考开发文档按照分类来查询。

对于某一块具体内容的设计，可以看看**视频教程**怎么做的，也可以在开发者社区看看


## lecture 2
对于基本的beego程序，介绍了是怎么run起来的，分析了实现的源码分别用三个例子重新实现了hello world

__钻研思路如下：__

temp实现了一个最简单的hello world
然后钻研这个程序是如何实现的，发现一共分为两步，
第一是注册了路由，绑定了对应函数，第二是运行

对于第一步注册路由，使用http.handlerFunc
本质上是用一个开一个mux，然后调用mux.handle()

所以在temp2中我们自己实现了handlerFunc
并且发现可以对mux里面注册多个函数，绑定多个路径

对于第二步运行，使用http.ListenandServe()
本质上是用http里面的Server结构体，定义了地址、handler等一系列属性，然后调用server.ListenAndServe()

所以在temp3中我们自己实现了一个server结构体，然后调用运行
思考：在temp3中如何实现temp2中绑定多个函数呢？
方式：针对Handler的ServeHTTP函数，进行路由分发

## lecture 3
1、介绍了前端语法里是怎么展示数据的，在前端模板中写了一些例子
    1. 循环输出数据 {{range .nums}} {{end}}
    2. 输出struct中的多个属性 {{with .User}} {{end}}
    3. 临时变量 {{$tplVar := .TplVar}} {{$tplVar}}
    4. html文本读取展示 {{str2html .Html}} "<div>hello</div>"展示结果为"hello"
    5. 嵌套模板使用，名为"test" {{template "test"}} 

2、介绍了sqlite的用法以及orm（对象关系映射），如何通过代码方便的初始化以及调用数据库的数据
    1. 定义数据结构，如何与orm联系，`orm:"index"`表示为该项建索引，`orm:"size(5000)"`表示更改设定值大小为5000字节
    2. 数据库注册，见models/models.go

3、用Bootstrap框架，重写前端页面，当然是比较简单的版本，又强调了一次如何去使用这些成熟的框架（看文档，检查元素等trick）
    1. 一些语法的应用，见views/home.html


```
1. 重定向的问题。  
在login的post方法中，现在写了一句`c.Redirect("/", 302)`，但是最初由于不清楚301（永久重定向）和302（临时重定向）的区别，写的是`c.Redirect("/", 301)`

这造成了一个现象，在登录之后点击注销的时候(本应发送'/login?exit=true'请求)，实际直接发送的'/'请求，打在logincontroller中get方法里的log也没打印 。

从现象来看，应该是永久的将传到'/login'的请求重定向到'/'了。即使我把代码中的301改成了302，重新运行也没解决。（手敲链接啥的都没用）

然后在检测各个方法时，将`c.Redirect("/", 302)`这句话先写成了`c.Redirect("/login?exit=true", 302)`，并且运行了一遍，发现这次并没有被重定向，我打在代码里的log也成功显示了。

再改回`c.Redirect("/", 302)`之后，一切正常！！！  

但是这个问题为什么会出现 以及 为什么这样又解决了还不是很清楚

有个简单的猜测（查询无果，待询问），在本地一旦运行一次301重定向，可能就出现了某个配置，记录了这个重定向，导致之后怎么访问'/login'都被跳转到'/'

即使我后面改成了302，由于重定向后的地址相同，所以配置仍然存在

在我改成重定向到"/login?exit=true"之后，由于地址不同，那个配置取消了，然后一切正常
```

# GoMiniblink

#### 介绍
Miniblink的免费版封装，官网：https://miniblink.net/

1.  不使用CGO
2.  面向对象
3.  跨平台设计，但目前只有一个windows实现
4.  组件和窗体两种模式
5.  JS互操作
6.  监控与拦截请求
7.  透明窗体
8.  支持本地目录加载模式

Go封装的功能比较少，其实就是 https://gitee.com/aochulai/NetMiniblink 的简化版，因为我出Go封装的目标是VIP，所以免费版就懒得像NetMiniblink一样写得那么完善啦，不过VIP版会向NetMiniblink完整度看齐。

### 简单的例子
    package main
    
    import (
    	gm "gitee.com/aochulai/GoMiniblink"
    	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
    	ws "gitee.com/aochulai/GoMiniblink/forms/windows"
    )
    
    func main() {
        //windows版本初始化
        cs.App = new(ws.Provider).Init()
        
        //创建一个窗体并设置基本属性
        frm := new(cs.Form).Init()
        frm.SetTitle("普通窗口")
        frm.SetSize(800, 500)
    	
        //创建浏览器控件并设置基本属性
        mb := new(gm.MiniblinkBrowser).Init()
        mb.SetSize(700, 400)
        
        //添加浏览器控件到窗体
        frm.AddChild(mb)
        //注册回调, EvLoad回调在窗体首次显示前触发
        frm.EvLoad["回调名称"] = func(s cs.GUI) {
            //加载网址
            mb.LoadUri("https://www.baidu.com")
        }
        //将frm作为主窗口打开
        cs.Run(frm)
    }
    
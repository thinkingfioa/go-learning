# go-learning
```
@author 鲁伟林
随着Oracle宣布JDK收费，Go语言席卷全球，云计算等各种新技术。Go作为后起之秀吸引了我的注意力。
鄙人一直Java语言开发已有5年，写过多个Java优秀项目。希望维护管理此仓库，记录学习Go语言部分心得，与其他同行参考。

本博客中涉及的完整代码：
GitHub地址: https://github.com/thinkingfioa/go-learning
本人博客地址: https://blog.csdn.net/thinking_fioa
文中如若有任何错误，欢迎指出。个人邮箱: thinking_fioa@163.com
```

## TODO 
- 学习cache2go项目代码
- nio方式Go语言Socket数据传输

## 项目结构介绍
|序号|名称|说明|
|:---:|:---:|:---:|
|1|goinaction|阅读《Go语言实战》笔记|
|2|go-practice|多个练手项目|

## 一、goinaction介绍
记录《Go语言实战》中各章节学习过程，写下一些自己的思考和总结。
分为4大部分，可直接查看[GitHub地址](https://github.com/thinkingfioa/go-learning/tree/master/goinaction)，或者通过专栏博客分章节产看

专栏博客地址

|序号|名称|
|:---:|:---|
|1|[Go专栏 (一) ——— 初识Go语言](https://blog.csdn.net/thinking_fioa/article/details/89289675)|
|1|[Go专栏 (二) ——— 初识Go语言](https://blog.csdn.net/thinking_fioa/article/details/89605937)|
|2|[Go专栏 (三) ——— 数组、切片和映射](https://blog.csdn.net/thinking_fioa/article/details/89289737)|
|3|[Go专栏 (四) ——— Go语言的类型系统](https://blog.csdn.net/thinking_fioa/article/details/89289876)|
|4|[Go专栏 (五) ——— 并发和并发模式](https://blog.csdn.net/thinking_fioa/article/details/89289980)|

## 二、go-practice介绍
go-practice目标是提供多个用Go语言实现的短小精悍的程序，方便初学者学习Go语言和了解Go语言的诸多特性，同时也多有一定开发经验的同学也有较大的益处。go-practice项目是官方依赖管理工具dep构建而成。

子项目介绍如下，[详细解释请阅读go-practice项目介绍](https://github.com/thinkingfioa/go-learning/tree/master/go-practice)

|序号|名称|介绍|
|:---:|:---:|:---:|
|1|concurrentcache|并发安全的Go语言cache|
|2|go-socket-nio|nio方式实现Socket数据传输|
|3|leetcode-in-action|leetcode上刷题记录|


## 参考

1. 《Go语言实战》
2. [muesli/cache2go](https://github.com/muesli/cache2go)
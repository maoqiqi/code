# NoSQL入门和概述

> GitHub地址：[NoSQL入门和概述](https://github.com/maoqiqi/One/blob/master/md/Java8.md)<br/>
> 邮箱：fengqi.mao.march@gmail.com<br/>
> 博客：http://blog.csdn.net/u011810138<br/>

NoSQL（NoSQL = Not Only SQL），意即“不仅仅是SQL”，泛指非关系型的数据库。

随着互联网web2.0网站的兴起，传统的关系数据库在应付web2.0网站，特别是超大规模和高并发的SNS类型的web2.0纯动态网站已经显得力不从心，
暴露了很多难以克服的问题，而非关系型的数据库则由于其本身的特点得到了非常迅速的发展。

NoSQL数据库的产生就是为了解决大规模数据集合多重数据种类带来的挑战，尤其是大数据应用难题，包括超大规模数据的存储。
例如谷歌或Facebook每天为他们的用户收集万亿比特的数据。这些类型的数据存储不需要固定的模式，无需多余操作就可以横向扩展。

## 互联网时代背景下大机遇，为什么用NoSQL

##### 1. 单机MySQL的美好年代

在90年代，一个网站的访问量一般都不大，用单个数据库完全可以轻松应付。
在那个时候，更多的都是静态网页，动态交互类型的网站不多。

上述架构下，我们来看看数据存储的瓶颈是什么？
- 数据量的总大小，一个机器放不下时
- 数据的索引（B + Tree），一个机器的内存放不下时
- 访问量(读写混合)，一个实例不能承受

##### 2. Memcached(缓存)+MySQL+垂直拆分
##### 3. Mysql主从读写分离
##### 4. 分表分库+水平拆分+MySQL集群
##### 5. MySQL的扩展性瓶颈
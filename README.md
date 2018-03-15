# NoSQL入门和概述

> GitHub地址：[NoSQL入门和概述]（https://github.com/maoqiqi/One/blob/master/md/Java8.md）<br/>
> 邮箱：fengqi.mao.march@gmail.com<br/>
> 博客：http://blog.csdn.net/u011810138<br/>

**NoSQL（NoSQL = Not Only SQL）**，意即“不仅仅是SQL”，泛指非关系型的数据库。

随着互联网web2.0网站的兴起，传统的关系数据库在应付web2.0网站，特别是超大规模和高并发的SNS类型的web2.0纯动态网站已经显得力不从心，
暴露了很多难以克服的问题，而非关系型的数据库则由于其本身的特点得到了非常迅速的发展。

NoSQL数据库的产生就是为了解决大规模数据集合多重数据种类带来的挑战，尤其是大数据应用难题，包括超大规模数据的存储。
例如谷歌或Facebook每天为他们的用户收集万亿比特的数据。这些类型的数据存储不需要固定的模式，无需多余操作就可以横向扩展。

## 互联网时代背景下大机遇，为什么用NoSQL

#### 1. 单机MySQL的美好年代

在90年代，一个网站的访问量一般都不大，用单个数据库完全可以轻松应付。
在那个时候，更多的都是静态网页，动态交互类型的网站不多。

上述架构下，我们来看看数据存储的瓶颈是什么？
- 数据量的总大小，一个机器放不下时
- 数据的索引（B + Tree），一个机器的内存放不下时
- 访问量（读写混合），一个实例不能承受

#### 2. Memcached（缓存）+MySQL+垂直拆分

后来，随着访问量的上升，几乎大部分使用MySQL架构的网站在数据库上都开始出现了性能问题，web程序不再仅仅专注在功能上，同时也在追求性能。
程序员们开始大量的使用缓存技术来缓解数据库的压力，优化数据库的结构和索引。
开始比较流行的是通过文件缓存来缓解数据库压力，但是当访问量继续增大的时候，多台web机器通过文件缓存不能共享，大量的小文件缓存也带了了比较高的IO压力。
在这个时候，Memcached就自然的成为一个非常时尚的技术产品。

Memcached作为一个独立的分布式的缓存服务器，为多个web服务器提供了一个共享的高性能缓存服务，
在Memcached服务器上，又发展了根据hash算法来进行多台Memcached缓存服务的扩展，
然后又出现了一致性hash来解决增加或减少缓存服务器导致重新hash带来的大量缓存失效的弊端。

#### 3. MySQL主从读写分离

由于数据库的写入压力增加，Memcached只能缓解数据库的读取压力。
读写集中在一个数据库上让数据库不堪重负，大部分网站开始使用主从复制技术来达到读写分离，以提高读写性能和读库的可扩展性。
MySQL的master-slave模式成为这个时候的网站标配了。

#### 4. 分表分库+水平拆分+MySQL集群

在Memcached的高速缓存，MySQL的主从复制，读写分离的基础之上，这时MySQL主库的写压力开始出现瓶颈，
而数据量的持续猛增，由于MyISAM使用表锁，在高并发下会出现严重的锁问题，大量的高并发MySQL应用开始使用InnoDB引擎代替MyISAM。

同时，开始流行使用分表分库来缓解写压力和数据增长的扩展问题。
这个时候，分表分库成了一个热门技术，是面试的热门问题也是业界讨论的热门技术问题。
也就在这个时候，MySQL推出了还不太稳定的表分区，这也给技术实力一般的公司带来了希望。
虽然MySQL推出了MySQL Cluster集群，但性能也不能很好满足互联网的要求，只是在高可靠性上提供了非常大的保证。

#### 5. MySQL的扩展性瓶颈

MySQL数据库也经常存储一些大文本字段，导致数据库表非常的大，在做数据库恢复的时候就导致非常的慢，不容易快速恢复数据库。
比如1000万4KB大小的文本就接近40GB的大小，如果能把这些数据从MySQL省去，MySQL将变得非常的小。
关系数据库很强大，但是它并不能很好的应付所有的应用场景。
MySQL的扩展性差（需要复杂的技术来实现），大数据下IO压力大，表结构更改困难，正是当前使用MySQL的开发人员面临的问题。

## 能干嘛？

#### 1.易扩展

NoSQL数据库种类繁多，但是一个共同的特点都是去掉关系数据库的关系型特性。
数据之间无关系，这样就非常容易扩展。
也无形之间，在架构的层面上带来了可扩展的能力。

#### 2.大数据量高性能

NoSQL数据库都具有非常高的读写性能，尤其在大数据量下，同样表现优秀。
这得益于它的无关系性，数据库的结构简单。

一般MySQL使用Query Cache，每次表的更新Cache就失效，是一种大粒度的Cache，在针对web2.0的交互频繁的应用，Cache性能不高。
而NoSQL的Cache是记录级的，是一种细粒度的Cache，所以NoSQL在这个层面上来说就要性能高很多了。

#### 3.多样灵活的数据模型

NoSQL无需事先为要存储的数据建立字段，随时可以存储自定义的数据格式。
而在关系数据库里，增删字段是一件非常麻烦的事情。如果是非常大数据量的表，增加字段简直就是一个噩梦。

#### 4.传统RDBMS VS NoSQL

RDBMS
- 高度组织化结构化数据
- 结构化查询语言（SQL）
- 数据和关系都存储在单独的表中。
- 数据操纵语言，数据定义语言
- 严格的一致性
- 基础事务

NoSQL
- 代表着不仅仅是SQL
- 没有声明性查询语言
- 没有预定义的模式
- 键 - 值对存储，列存储，文档存储，图形数据库
- 最终一致性，而非ACID属性
- 非结构化和不可预知的数据
- CAP定理
- 高性能，高可用性和可伸缩性

## 3V+3高

**大数据时代的3V：海量Volume，多样Variety，实时Velocity**

**互联网需求的3高：高并发，高可扩，高性能**

## 当下的NoSQL经典应用

## NoSQL数据模型简介

## NoSQL数据库的四大分类

## 在分布式数据库中CAP原理CAP+BASE

## Redis

### 一. Redis入门介绍

#### redis入门概述

#### redis的安装

#### redis启动后杂项基础知识讲解

### 二. Redis数据类型

#### 1. Redis的五大数据类型

> * String（字符串）

string 是redis最基本的类型，你可以理解成与Memcached一模一样的类型，一个key对应一个value。

string 类型是二进制安全的。意思是redis的string可以包含任何数据。比如jpg图片或者序列化的对象 。

string 类型是Redis最基本的数据类型，一个redis中字符串value最多可以是512M。

> * Hash（哈希）

hash 是一个键值对集合。

hash 是一个string类型的field和value的映射表，hash特别适合用于存储对象。

类似Java里面的Map<String,Object>。

> * List（列表）

List 是简单的字符串列表，按照插入顺序排序。你可以添加一个元素导列表的头部（左边）或者尾部（右边）。它的底层实际是个链表。

> * Set（集合）

Set 是string类型的无序集合。它是通过HashTable实现实现的。

> * zset（sorted set：有序集合）

zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

zset的成员是唯一的,但分数（score）却可以重复。

**注意：**[哪里去获得redis常见数据类型操作命令]（http://redisdoc.com/）

#### 2. String（字符串）

> SET key value [EX seconds] [PX milliseconds] [NX|XX]

将字符串值 value 关联到 key 。

如果 key 已经持有其他值， SET 就覆写旧值，无视类型。

对于某个原本带有生存时间（TTL）的键来说， 当 SET 命令成功在这个键上执行时， 这个键原有的 TTL 将被清除。

> GET key

返回 key 所关联的字符串值。

如果 key 不存在那么返回特殊值 nil 。

假如 key 储存的值不是字符串类型，返回一个错误，因为 GET 只能用于处理字符串值。

> STRLEN key

返回 key 所储存的字符串值的长度。

当 key 储存的不是字符串值时，返回一个错误。

> APPEND key value

如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。

如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。

> SETEX key seconds value

将值 value 关联到 key ，并将 key 的生存时间设为 seconds （以秒为单位）。

如果 key 已经存在， SETEX 命令将覆写旧值。

这个命令类似于以下两个命令：

```
SET key value
EXPIRE key seconds  # 设置生存时间
```

不同之处是， SETEX 是一个原子性（atomic）操作，关联值和设置生存时间两个动作会在同一时间内完成，该命令在 Redis 用作缓存时，非常实用。

> SETNX key value

将 key 的值设为 value ，当且仅当 key 不存在。

若给定的 key 已经存在，则 SETNX 不做任何动作。

SETNX 是『SET if Not eXists』（如果不存在，则 SET）的简写。

> GETRANGE key start end

返回 key 中字符串值的子字符串，字符串的截取范围由 start 和 end 两个偏移量决定（包括 start 和 end 在内）。

负数偏移量表示从字符串最后开始计数， -1 表示最后一个字符， -2 表示倒数第二个，以此类推。

GETRANGE 通过保证子字符串的值域（range）不超过实际字符串的值域来处理超出范围的值域请求。

> SETRANGE key offset value

用 value 参数覆写（overwrite）给定 key 所储存的字符串值，从偏移量 offset 开始。

不存在的 key 当作空白字符串处理。

SETRANGE 命令会确保字符串足够长以便将 value 设置在指定的偏移量上，如果给定 key 原来储存的字符串长度比偏移量小（比如字符串只有 5 个字符长，但你设置的 offset 是 10 ），那么原字符和偏移量之间的空白将用零字节（zerobytes, "\x00" ）来填充。

注意你能使用的最大偏移量是 2^29-1（536870911） ，因为 Redis 字符串的大小被限制在 512 兆（megabytes）以内。如果你需要使用比这更大的空间，你可以使用多个 key 。

> INCR key

将 key 中储存的数字值增一。

如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。

如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。

本操作的值限制在 64 位（bit）有符号数字表示之内。

> DECR key

将 key 中储存的数字值减一。

如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECR 操作。

如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。

本操作的值限制在 64 位（bit）有符号数字表示之内。

> INCRBY key increment

将 key 所储存的值加上增量 increment 。

如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCRBY 命令。

如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。

本操作的值限制在 64 位（bit）有符号数字表示之内。

> DECRBY key decrement

将 key 所储存的值减去减量 decrement 。

如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECRBY 操作。

如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。

本操作的值限制在 64 位（bit）有符号数字表示之内。

> MSET key value [key value ...]

同时设置一个或多个 key-value 对。

如果某个给定 key 已经存在，那么 MSET 会用新值覆盖原来的旧值，如果这不是你所希望的效果，请考虑使用 MSETNX 命令：它只会在所有给定 key 都不存在的情况下进行设置操作。

MSET 是一个原子性（atomic）操作，所有给定 key 都会在同一时间内被设置，某些给定 key 被更新而另一些给定 key 没有改变的情况，不可能发生。

> MGET key [key ...]

返回所有（一个或多个）给定 key 的值。

如果给定的 key 里面，有某个 key 不存在，那么这个 key 返回特殊值 nil 。因此，该命令永不失败。

> MSETNX key value [key value ...]

MSETNX key value [key value ...]

同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。

即使只有一个给定 key 已存在， MSETNX 也会拒绝执行所有给定 key 的设置操作。

MSETNX 是原子性的，因此它可以用作设置多个不同 key 表示不同字段（field）的唯一性逻辑对象（unique logic object），所有字段要么全被设置，要么全不被设置。

> GETSET key value

将给定 key 的值设为 value ，并返回 key 的旧值（old value）。

当 key 存在但不是字符串类型时，返回一个错误。

#### 3. Hash（哈希）

> HSET key field value

HSET key field value

将哈希表 key 中的域 field 的值设为 value 。

如果 key 不存在，一个新的哈希表被创建并进行 HSET 操作。

如果域 field 已经存在于哈希表中，旧值将被覆盖。

> HGET key field

返回哈希表 key 中给定域 field 的值。

> HMSET key field value [field value ...]

同时将多个 field-value （域-值）对设置到哈希表 key 中。

此命令会覆盖哈希表中已存在的域。

如果 key 不存在，一个空哈希表被创建并执行 HMSET 操作。

> HMGET key field [field ...]

返回哈希表 key 中，一个或多个给定域的值。

如果给定的域不存在于哈希表，那么返回一个 nil 值。

因为不存在的 key 被当作一个空哈希表来处理，所以对一个不存在的 key 进行 HMGET 操作将返回一个只带有 nil 值的表。

> HGETALL key

返回哈希表 key 中，所有的域和值。

在返回值里，紧跟每个域名（field name）之后是域的值（value），所以返回值的长度是哈希表大小的两倍。

> HDEL key field [field ...]

删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。

> HLEN key

返回哈希表 key 中域的数量。

> HEXISTS key field

查看哈希表 key 中，给定域 field 是否存在。

> HKEYS key

返回哈希表 key 中的所有域。

> HVALS key

返回哈希表 key 中所有域的值。

> HINCRBY key field increment

为哈希表 key 中的域 field 的值加上增量 increment 。

增量也可以为负数，相当于对给定域进行减法操作。

如果 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令。

如果域 field 不存在，那么在执行命令前，域的值被初始化为 0 。

对一个储存字符串值的域 field 执行 HINCRBY 命令将造成一个错误。

本操作的值被限制在 64 位（bit）有符号数字表示之内。

> HINCRBYFLOAT key field increment

HINCRBYFLOAT key field increment

为哈希表 key 中的域 field 加上浮点数增量 increment 。

如果哈希表中没有域 field ，那么 HINCRBYFLOAT 会先将域 field 的值设为 0 ，然后再执行加法操作。

如果键 key 不存在，那么 HINCRBYFLOAT 会先创建一个哈希表，再创建域 field ，最后再执行加法操作。

当以下任意一个条件发生时，返回一个错误：

* 域 field 的值不是字符串类型（因为 redis 中的数字和浮点数都以字符串的形式保存，所以它们都属于字符串类型）
* 域 field 当前的值或给定的增量 increment 不能解释（parse）为双精度浮点数（double precision floating point number）

HINCRBYFLOAT 命令的详细功能和 INCRBYFLOAT 命令类似。

> HSETNX key field value

将哈希表 key 中的域 field 的值设置为 value ，当且仅当域 field 不存在。

若域 field 已经存在，该操作无效。

如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。

#### 4. List（列表）

> LPUSH key value [value ...]

LPUSH key value [value ...]

将一个或多个值 value 插入到列表 key 的表头

如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表头： 比如说，对空列表 mylist 执行命令 LPUSH mylist a b c ，列表的值将是 c b a ，这等同于原子性地执行 LPUSH mylist a 、 LPUSH mylist b 和 LPUSH mylist c 三个命令。

如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。

当 key 存在但不是列表类型时，返回一个错误。

> RPUSH key value [value ...]

将一个或多个值 value 插入到列表 key 的表尾(最右边)。

如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表尾：比如对一个空列表 mylist 执行 RPUSH mylist a b c ，得出的结果列表为 a b c ，等同于执行命令 RPUSH mylist a 、 RPUSH mylist b 、 RPUSH mylist c 。

如果 key 不存在，一个空列表会被创建并执行 RPUSH 操作。

当 key 存在但不是列表类型时，返回一个错误。

> LRANGE key start stop

返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

**注意LRANGE命令和编程语言区间函数的区别**

假如你有一个包含一百个元素的列表，对该列表执行 LRANGE list 0 10 ，结果是一个包含11个元素的列表，这表明 stop 下标也在 LRANGE 命令的取值范围之内(闭区间)，这和某些语言的区间函数可能不一致，比如Ruby的 Range.new 、 Array#slice 和Python的 range() 函数。

**超出范围的下标**

超出范围的下标值不会引起错误。

如果 start 下标比列表的最大下标 end ( LLEN list 减去 1 )还要大，那么 LRANGE 返回一个空列表。

如果 stop 下标比 end 下标还要大，Redis将 stop 的值设置为 end 。

> LPOP key

移除并返回列表 key 的头元素。

> RPOP key

移除并返回列表 key 的尾元素。

> LINDEX key index

返回列表 key 中，下标为 index 的元素。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

如果 key 不是列表类型，返回一个错误。

> LLEN key

返回列表 key 的长度。

如果 key 不存在，则 key 被解释为一个空列表，返回 0 .

如果 key 不是列表类型，返回一个错误。

> LREM key count value

根据参数 count 的值，移除列表中与参数 value 相等的元素。

count 的值可以是以下几种：

* count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
* count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
* count = 0 : 移除表中所有与 value 相等的值。

> LTRIM key start stop

对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。

举个例子，执行命令 LTRIM list 0 2 ，表示只保留列表 list 的前三个元素，其余元素全部删除。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

当 key 不是列表类型时，返回一个错误。

LTRIM 命令通常和 LPUSH 命令或 RPUSH 命令配合使用，举个例子：

```
LPUSH log newest_log
LTRIM log 0 99
```

这个例子模拟了一个日志程序，每次将最新日志 newest_log 放到 log 列表中，并且只保留最新的 100 项。注意当这样使用 LTRIM 命令时，时间复杂度是O(1)，因为平均情况下，每次只有一个元素被移除。

**注意LTRIM命令和编程语言区间函数的区别**

假如你有一个包含一百个元素的列表 list ，对该列表执行 LTRIM list 0 10 ，结果是一个包含11个元素的列表，这表明 stop 下标也在 LTRIM 命令的取值范围之内(闭区间)，这和某些语言的区间函数可能不一致，比如Ruby的 Range.new 、 Array#slice 和Python的 range() 函数。

**超出范围的下标**

超出范围的下标值不会引起错误。

如果 start 下标比列表的最大下标 end ( LLEN list 减去 1 )还要大，或者 start > stop ， LTRIM 返回一个空列表(因为 LTRIM 已经将整个列表清空)。

如果 stop 下标比 end 下标还要大，Redis将 stop 的值设置为 end 。

> RPOPLPUSH source destination

命令 RPOPLPUSH 在一个原子时间内，执行以下两个动作：

* 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。
* 将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。

举个例子，你有两个列表 source 和 destination ， source 列表有元素 a, b, c ， destination 列表有元素 x, y, z ，执行 RPOPLPUSH source destination 之后， source 列表包含元素 a, b ， destination 列表包含元素 c, x, y, z ，并且元素 c 会被返回给客户端。

如果 source 不存在，值 nil 被返回，并且不执行其他动作。

如果 source 和 destination 相同，则列表中的表尾元素被移动到表头，并返回该元素，可以把这种特殊情况视作列表的旋转(rotation)操作。

> LSET key index value

将列表 key 下标为 index 的元素的值设置为 value 。

当 index 参数超出范围，或对一个空列表( key 不存在)进行 LSET 时，返回一个错误。

> LINSERT key BEFORE|AFTER pivot value

当 pivot 不存在于列表 key 时，不执行任何操作。

当 key 不存在时， key 被视为空列表，不执行任何操作。

如果 key 不是列表类型，返回一个错误。

**性能总结**

它是一个字符串链表，left、right都可以插入添加；

如果键不存在，创建新的链表；

如果键已存在，新增内容；

如果值全移除，对应的键也就消失了。

链表的操作无论是头和尾效率都极高，但假如是对中间元素进行操作，效率就很惨淡了。

#### 5. Set（集合）

> SADD key member [member ...]

将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。

假如 key 不存在，则创建一个只包含 member 元素作成员的集合。

当 key 不是集合类型时，返回一个错误。

> SMEMBERS key

返回集合 key 中的所有成员。

不存在的 key 被视为空集合。

> SISMEMBER key member

判断 member 元素是否集合 key 的成员。

> SCARD key

返回集合 key 的基数(集合中元素的数量)。

> SREM key member [member ...]

移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略。

当 key 不是集合类型，返回一个错误。

> SRANDMEMBER key [count]

如果命令执行时，只提供了 key 参数，那么返回集合中的一个随机元素。

从 Redis 2.6 版本开始， SRANDMEMBER 命令接受可选的 count 参数：

* 如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合。
* 如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值。

该操作和 SPOP 相似，但 SPOP 将随机元素从集合中移除并返回，而 SRANDMEMBER 则仅仅返回随机元素，而不对集合进行任何改动。

> SPOP key

移除并返回集合中的一个随机元素。

如果只想获取一个随机元素，但不想该元素从集合中被移除的话，可以使用 SRANDMEMBER 命令。

> SMOVE source destination member

将 member 元素从 source 集合移动到 destination 集合。

SMOVE 是原子性操作。

如果 source 集合不存在或不包含指定的 member 元素，则 SMOVE 命令不执行任何操作，仅返回 0 。否则， member 元素从 source 集合中被移除，并添加到 destination 集合中去。

当 destination 集合已经包含 member 元素时， SMOVE 命令只是简单地将 source 集合中的 member 元素删除。

当 source 或 destination 不是集合类型时，返回一个错误。

> SDIFF key [key ...]

返回一个集合的全部成员，该集合是所有给定集合之间的差集。

不存在的 key 被视为空集。

> SINTER key [key ...]

返回一个集合的全部成员，该集合是所有给定集合的交集。

不存在的 key 被视为空集。

当给定集合当中有一个空集时，结果也为空集(根据集合运算定律)。

> SUNION key [key ...]

返回一个集合的全部成员，该集合是所有给定集合的并集。

不存在的 key 被视为空集。

#### 6. zset（sorted set：有序集合）

> ZADD key score member [[score member] [score member] ...]

将一个或多个 member 元素及其 score 值加入到有序集 key 当中。

如果某个 member 已经是有序集的成员，那么更新这个 member 的 score 值，并通过重新插入这个 member 元素，来保证该 member 在正确的位置上。

score 值可以是整数值或双精度浮点数。

如果 key 不存在，则创建一个空的有序集并执行 ZADD 操作。

当 key 存在但不是有序集类型时，返回一个错误。

对有序集的更多介绍请参见 sorted set 。

> ZRANGE key start stop [WITHSCORES]

返回有序集 key 中，指定区间内的成员。

其中成员的位置按 score 值递增(从小到大)来排序。

具有相同 score 值的成员按字典序(lexicographical order )来排列。

如果你需要成员按 score 值递减(从大到小)来排列，请使用 ZREVRANGE 命令。

下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。

你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。

超出范围的下标并不会引起错误。

比如说，当 start 的值比有序集的最大下标还要大，或是 start > stop 时， ZRANGE 命令只是简单地返回一个空列表。

另一方面，假如 stop 参数的值比有序集的最大下标还要大，那么 Redis 将 stop 当作最大下标来处理。

可以通过使用 WITHSCORES 选项，来让成员和它的 score 值一并返回，返回列表以 value1,score1, ..., valueN,scoreN 的格式表示。

客户端库可能会返回一些更复杂的数据类型，比如数组、元组等。

#### 7. Key（键）

> KEYS pattern

查找所有符合给定模式 pattern 的 key 。

KEYS * 匹配数据库中所有 key 。

KEYS h?llo 匹配 hello ， hallo 和 hxllo 等。

KEYS h*llo 匹配 hllo 和 heeeeello 等。

KEYS h[ae]llo 匹配 hello 和 hallo ，但不匹配 hillo 。

特殊符号用 \ 隔开

> EXISTS key

检查给定 key 是否存在。

> MOVE key db

将当前数据库的 key 移动到给定的数据库 db 当中。

如果当前数据库(源数据库)和给定数据库(目标数据库)有相同名字的给定 key ，或者 key 不存在于当前数据库，那么 MOVE 没有任何效果。

因此，也可以利用这一特性，将 MOVE 当作锁(locking)原语(primitive)。

> EXPIRE key seconds

为给定 key 设置生存时间，当 key 过期时(生存时间为 0 )，它会被自动删除。

在 Redis 中，带有生存时间的 key 被称为『易失的』(volatile)。

生存时间可以通过使用 DEL 命令来删除整个 key 来移除，或者被 SET 和 GETSET 命令覆写(overwrite)，这意味着，如果一个命令只是修改(alter)一个带生存时间的 key 的值而不是用一个新的 key 值来代替(replace)它的话，那么生存时间不会被改变。

比如说，对一个 key 执行 INCR 命令，对一个列表进行 LPUSH 命令，或者对一个哈希表执行 HSET 命令，这类操作都不会修改 key 本身的生存时间。

另一方面，如果使用 RENAME 对一个 key 进行改名，那么改名后的 key 的生存时间和改名前一样。

RENAME 命令的另一种可能是，尝试将一个带生存时间的 key 改名成另一个带生存时间的 another_key ，这时旧的 another_key (以及它的生存时间)会被删除，然后旧的 key 会改名为 another_key ，因此，新的 another_key 的生存时间也和原本的 key 一样。

使用 PERSIST 命令可以在不删除 key 的情况下，移除 key 的生存时间，让 key 重新成为一个『持久的』(persistent) key 。

**更新生存时间**

可以对一个已经带有生存时间的 key 执行 EXPIRE 命令，新指定的生存时间会取代旧的生存时间。

**过期时间的精确度**

在 Redis 2.4 版本中，过期时间的延迟在 1 秒钟之内 —— 也即是，就算 key 已经过期，但它还是可能在过期之后一秒钟之内被访问到，而在新的 Redis 2.6 版本中，延迟被降低到 1 毫秒之内。

**Redis 2.1.3 之前的不同之处**

在 Redis 2.1.3 之前的版本中，修改一个带有生存时间的 key 会导致整个 key 被删除，这一行为是受当时复制(replication)层的限制而作出的，现在这一限制已经被修复。

> TTL key

以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。

> TYPE key

返回 key 所储存的值的类型。

### 三. 解析配置文件redis.conf

### 四. Redis的持久化

### 五. Redis的事务

### 六. Redis的发布订阅

### 七. Redis的复制（Master/Slave）

### 八. Redis的Java客户端Jedis

## MemCached

## MongoDB
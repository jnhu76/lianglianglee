<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=11&#32;&#32;Redis事件驱动框架（下）：Redis有哪些事件？>
        <link rel="icon" href="/static/favicon.png">
        <title>11  Redis事件驱动框架（下）：Redis有哪些事件？ </title>
        
        <link rel="stylesheet" href="/static/index.css">
        <link rel="stylesheet" href="/static/highlight.min.css">
        <script src="/static/highlight.min.js"></script>
        
        <meta name="generator" content="Hexo 4.2.0">
        <script defer src="https://umami.lianglianglee.com/script.js"
         data-website-id="83e5d5db-9d06-40e3-b780-cbae722fdf8c"></script>
    </head>

<body>
    <div class="book-container">
        <div class="book-sidebar">
            <div class="book-brand">
                <a href="/">
                    <img src="/static/favicon.png">
                    <span>技术文章摘抄</span>
                </a>
            </div>
            <div class="book-menu uncollapsible">
                <ul class="uncollapsible">
                    <li><a href="/" class="current-tab">首页</a></li>
                    <li><a href="../">上一级</a></li>
                </ul>
                <ul class="uncollapsible">
                    
                    <li>
                        <a class="menu-item" id="00  开篇词  阅读Redis源码能给你带来什么？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e9%98%85%e8%af%bbRedis%e6%ba%90%e7%a0%81%e8%83%bd%e7%bb%99%e4%bd%a0%e5%b8%a6%e6%9d%a5%e4%bb%80%e4%b9%88%ef%bc%9f.md">00  开篇词  阅读Redis源码能给你带来什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  带你快速攻略Redis源码的整体架构.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%20%e5%b8%a6%e4%bd%a0%e5%bf%ab%e9%80%9f%e6%94%bb%e7%95%a5Redis%e6%ba%90%e7%a0%81%e7%9a%84%e6%95%b4%e4%bd%93%e6%9e%b6%e6%9e%84.md">01  带你快速攻略Redis源码的整体架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  键值对中字符串的实现，用char还是结构体？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/02%20%20%e9%94%ae%e5%80%bc%e5%af%b9%e4%b8%ad%e5%ad%97%e7%ac%a6%e4%b8%b2%e7%9a%84%e5%ae%9e%e7%8e%b0%ef%bc%8c%e7%94%a8char%e8%bf%98%e6%98%af%e7%bb%93%e6%9e%84%e4%bd%93%ef%bc%9f.md">02  键值对中字符串的实现，用char还是结构体？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  如何实现一个性能优异的Hash表？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/03%20%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e6%80%a7%e8%83%bd%e4%bc%98%e5%bc%82%e7%9a%84Hash%e8%a1%a8%ef%bc%9f.md">03  如何实现一个性能优异的Hash表？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  内存友好的数据结构该如何细化设计？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/04%20%20%e5%86%85%e5%ad%98%e5%8f%8b%e5%a5%bd%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e8%af%a5%e5%a6%82%e4%bd%95%e7%bb%86%e5%8c%96%e8%ae%be%e8%ae%a1%ef%bc%9f.md">04  内存友好的数据结构该如何细化设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  有序集合为何能同时支持点查询和范围查询？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%20%e6%9c%89%e5%ba%8f%e9%9b%86%e5%90%88%e4%b8%ba%e4%bd%95%e8%83%bd%e5%90%8c%e6%97%b6%e6%94%af%e6%8c%81%e7%82%b9%e6%9f%a5%e8%af%a2%e5%92%8c%e8%8c%83%e5%9b%b4%e6%9f%a5%e8%af%a2%ef%bc%9f.md">05  有序集合为何能同时支持点查询和范围查询？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  从ziplist到quicklist，再到listpack的启发.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/06%20%20%e4%bb%8eziplist%e5%88%b0quicklist%ef%bc%8c%e5%86%8d%e5%88%b0listpack%e7%9a%84%e5%90%af%e5%8f%91.md">06  从ziplist到quicklist，再到listpack的启发.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  为什么Stream使用了Radix Tree？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%20%e4%b8%ba%e4%bb%80%e4%b9%88Stream%e4%bd%bf%e7%94%a8%e4%ba%86Radix%20Tree%ef%bc%9f.md">07  为什么Stream使用了Radix Tree？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  Redis server启动后会做哪些操作？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%20Redis%20server%e5%90%af%e5%8a%a8%e5%90%8e%e4%bc%9a%e5%81%9a%e5%93%aa%e4%ba%9b%e6%93%8d%e4%bd%9c%ef%bc%9f.md">08  Redis server启动后会做哪些操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  Redis事件驱动框架（上）：何时使用select、poll、epoll？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%20Redis%e4%ba%8b%e4%bb%b6%e9%a9%b1%e5%8a%a8%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%bd%95%e6%97%b6%e4%bd%bf%e7%94%a8select%e3%80%81poll%e3%80%81epoll%ef%bc%9f.md">09  Redis事件驱动框架（上）：何时使用select、poll、epoll？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  Redis事件驱动框架（中）：Redis实现了Reactor模型吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%20Redis%e4%ba%8b%e4%bb%b6%e9%a9%b1%e5%8a%a8%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9aRedis%e5%ae%9e%e7%8e%b0%e4%ba%86Reactor%e6%a8%a1%e5%9e%8b%e5%90%97%ef%bc%9f.md">10  Redis事件驱动框架（中）：Redis实现了Reactor模型吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  Redis事件驱动框架（下）：Redis有哪些事件？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%20Redis%e4%ba%8b%e4%bb%b6%e9%a9%b1%e5%8a%a8%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aRedis%e6%9c%89%e5%93%aa%e4%ba%9b%e4%ba%8b%e4%bb%b6%ef%bc%9f.md">11  Redis事件驱动框架（下）：Redis有哪些事件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  Redis真的是单线程吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%20Redis%e7%9c%9f%e7%9a%84%e6%98%af%e5%8d%95%e7%ba%bf%e7%a8%8b%e5%90%97%ef%bc%9f.md">12  Redis真的是单线程吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  Redis 6.0多IO线程的效率提高了吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%20Redis%206.0%e5%a4%9aIO%e7%ba%bf%e7%a8%8b%e7%9a%84%e6%95%88%e7%8e%87%e6%8f%90%e9%ab%98%e4%ba%86%e5%90%97%ef%bc%9f.md">13  Redis 6.0多IO线程的效率提高了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  从代码实现看分布式锁的原子性保证.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%20%e4%bb%8e%e4%bb%a3%e7%a0%81%e5%ae%9e%e7%8e%b0%e7%9c%8b%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e7%9a%84%e5%8e%9f%e5%ad%90%e6%80%a7%e4%bf%9d%e8%af%81.md">14  从代码实现看分布式锁的原子性保证.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  为什么LRU算法原理和代码实现不一样？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/15%20%20%e4%b8%ba%e4%bb%80%e4%b9%88LRU%e7%ae%97%e6%b3%95%e5%8e%9f%e7%90%86%e5%92%8c%e4%bb%a3%e7%a0%81%e5%ae%9e%e7%8e%b0%e4%b8%8d%e4%b8%80%e6%a0%b7%ef%bc%9f.md">15  为什么LRU算法原理和代码实现不一样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  LFU算法和其他算法相比有优势吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%20LFU%e7%ae%97%e6%b3%95%e5%92%8c%e5%85%b6%e4%bb%96%e7%ae%97%e6%b3%95%e7%9b%b8%e6%af%94%e6%9c%89%e4%bc%98%e5%8a%bf%e5%90%97%ef%bc%9f.md">16  LFU算法和其他算法相比有优势吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  Lazy Free会影响缓存替换吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%20Lazy%20Free%e4%bc%9a%e5%bd%b1%e5%93%8d%e7%bc%93%e5%ad%98%e6%9b%bf%e6%8d%a2%e5%90%97%ef%bc%9f.md">17  Lazy Free会影响缓存替换吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  如何生成和解读RDB文件？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/18%20%20%e5%a6%82%e4%bd%95%e7%94%9f%e6%88%90%e5%92%8c%e8%a7%a3%e8%af%bbRDB%e6%96%87%e4%bb%b6%ef%bc%9f.md">18  如何生成和解读RDB文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  AOF重写（上）：触发时机与重写的影响.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/19%20%20AOF%e9%87%8d%e5%86%99%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e8%a7%a6%e5%8f%91%e6%97%b6%e6%9c%ba%e4%b8%8e%e9%87%8d%e5%86%99%e7%9a%84%e5%bd%b1%e5%93%8d.md">19  AOF重写（上）：触发时机与重写的影响.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  AOF重写（下）：重写时的新写操作记录在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%20AOF%e9%87%8d%e5%86%99%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e9%87%8d%e5%86%99%e6%97%b6%e7%9a%84%e6%96%b0%e5%86%99%e6%93%8d%e4%bd%9c%e8%ae%b0%e5%bd%95%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">20  AOF重写（下）：重写时的新写操作记录在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  主从复制：基于状态机的设计与实现.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/21%20%20%e4%b8%bb%e4%bb%8e%e5%a4%8d%e5%88%b6%ef%bc%9a%e5%9f%ba%e4%ba%8e%e7%8a%b6%e6%80%81%e6%9c%ba%e7%9a%84%e8%ae%be%e8%ae%a1%e4%b8%8e%e5%ae%9e%e7%8e%b0.md">21  主从复制：基于状态机的设计与实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  哨兵也和Redis实例一样初始化吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%20%e5%93%a8%e5%85%b5%e4%b9%9f%e5%92%8cRedis%e5%ae%9e%e4%be%8b%e4%b8%80%e6%a0%b7%e5%88%9d%e5%a7%8b%e5%8c%96%e5%90%97%ef%bc%9f.md">22  哨兵也和Redis实例一样初始化吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  从哨兵Leader选举学习Raft协议实现（上）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%20%e4%bb%8e%e5%93%a8%e5%85%b5Leader%e9%80%89%e4%b8%be%e5%ad%a6%e4%b9%a0Raft%e5%8d%8f%e8%ae%ae%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8a%ef%bc%89.md">23  从哨兵Leader选举学习Raft协议实现（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  从哨兵Leader选举学习Raft协议实现（下）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%20%e4%bb%8e%e5%93%a8%e5%85%b5Leader%e9%80%89%e4%b8%be%e5%ad%a6%e4%b9%a0Raft%e5%8d%8f%e8%ae%ae%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8b%ef%bc%89.md">24  从哨兵Leader选举学习Raft协议实现（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  PubSub在主从故障切换时是如何发挥作用的？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%20PubSub%e5%9c%a8%e4%b8%bb%e4%bb%8e%e6%95%85%e9%9a%9c%e5%88%87%e6%8d%a2%e6%97%b6%e6%98%af%e5%a6%82%e4%bd%95%e5%8f%91%e6%8c%a5%e4%bd%9c%e7%94%a8%e7%9a%84%ef%bc%9f.md">25  PubSub在主从故障切换时是如何发挥作用的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  从Ping-Pong消息学习Gossip协议的实现.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%20%e4%bb%8ePing-Pong%e6%b6%88%e6%81%af%e5%ad%a6%e4%b9%a0Gossip%e5%8d%8f%e8%ae%ae%e7%9a%84%e5%ae%9e%e7%8e%b0.md">26  从Ping-Pong消息学习Gossip协议的实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  从MOVED、ASK看集群节点如何处理命令？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%20%e4%bb%8eMOVED%e3%80%81ASK%e7%9c%8b%e9%9b%86%e7%be%a4%e8%8a%82%e7%82%b9%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e5%91%bd%e4%bb%a4%ef%bc%9f.md">27  从MOVED、ASK看集群节点如何处理命令？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  Redis Cluster数据迁移会阻塞吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%20Redis%20Cluster%e6%95%b0%e6%8d%ae%e8%bf%81%e7%a7%bb%e4%bc%9a%e9%98%bb%e5%a1%9e%e5%90%97%ef%bc%9f.md">28  Redis Cluster数据迁移会阻塞吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  如何正确实现循环缓冲区？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e5%ae%9e%e7%8e%b0%e5%be%aa%e7%8e%af%e7%bc%93%e5%86%b2%e5%8c%ba%ef%bc%9f.md">29  如何正确实现循环缓冲区？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  如何在系统中实现延迟监控？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%20%e5%a6%82%e4%bd%95%e5%9c%a8%e7%b3%bb%e7%bb%9f%e4%b8%ad%e5%ae%9e%e7%8e%b0%e5%bb%b6%e8%bf%9f%e7%9b%91%e6%8e%a7%ef%bc%9f.md">30  如何在系统中实现延迟监控？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  从Module的实现学习动态扩展功能.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/31%20%20%e4%bb%8eModule%e7%9a%84%e5%ae%9e%e7%8e%b0%e5%ad%a6%e4%b9%a0%e5%8a%a8%e6%80%81%e6%89%a9%e5%b1%95%e5%8a%9f%e8%83%bd.md">31  从Module的实现学习动态扩展功能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  如何在一个系统中实现单元测试？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%20%e5%a6%82%e4%bd%95%e5%9c%a8%e4%b8%80%e4%b8%aa%e7%b3%bb%e7%bb%9f%e4%b8%ad%e5%ae%9e%e7%8e%b0%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%ef%bc%9f.md">32  如何在一个系统中实现单元测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  Redis源码阅读，让我们从新开始.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%ba%90%e7%a0%81%e5%89%96%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20Redis%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb%ef%bc%8c%e8%ae%a9%e6%88%91%e4%bb%ac%e4%bb%8e%e6%96%b0%e5%bc%80%e5%a7%8b.md">结束语  Redis源码阅读，让我们从新开始.md</a>
                    </li>
                    
                    <li><a href="https://lianglianglee.com/assets/%E6%8D%90%E8%B5%A0.md">捐赠</a></li>
                </ul>

            </div>
        </div>

        <div class="sidebar-toggle" onclick="sidebar_toggle()" onmouseover="add_inner()" onmouseleave="remove_inner()">
            <div class="sidebar-toggle-inner"></div>
        </div>
        <div class="off-canvas-content">
            <div class="columns">
                <div class="column col-12 col-lg-12">
                    <div class="book-navbar">
                        
                        <header class="navbar">
                            <section class="navbar-section">
                                <a onclick="open_sidebar()">
                                    <i class="icon icon-menu"></i>
                                </a>
                            </section>
                        </header>
                    </div>
                    <div class="book-content" style="max-width: 960px; margin: 0 auto;
    overflow-x: auto;
    overflow-y: hidden;">
                        <div class="book-post">
                            
                            
                            
                            <p id="tip" align="center"></p>
                            <h1 id="title" data-id="11  Redis事件驱动框架（下）：Redis有哪些事件？" class="title">11  Redis事件驱动框架（下）：Redis有哪些事件？</h1>
                            <div><p>在【第 9 讲】中，我给你介绍了 Linux 提供的三种 IO 多路复用机制，分别是 select、poll 和 epoll，这是 Redis 实现事件驱动框架的操作系统层支撑技术。</p>

<p>紧接着在上节课，我带你学习了 Redis 事件驱动框架的基本工作机制，其中介绍了事件驱动框架基于的 Reactor 模型，并以 IO 事件中的客户端连接事件为例，给你介绍了<strong>框架运行的基本流程</strong>：从 server 初始化时调用 aeCreateFileEvent 函数注册监听事件，到 server 初始化完成后调用 aeMain 函数，而 aeMain 函数循环执行 aeProceeEvent 函数，来捕获和处理客户端请求触发的事件。</p>

<p>但是在上节课当中，我们主要关注的是框架基本流程，所以到这里，你或许仍然存有一些疑问，比如说：</p>

<ul>
<li>Redis 事件驱动框架监听的 IO 事件，除了上节课介绍的客户端连接以外，还有没有其他事件？而除了 IO 事件以外，框架还会监听其他事件么？</li>
<li>这些事件的创建和处理又分别对应了 Redis 源码中的哪些具体操作？</li>
</ul>

<p>今天这节课，我就来给你介绍下 Redis 事件驱动框架中的两大类事件类型：<strong>IO 事件和时间事件，以及它们相应的处理机制。</strong></p>

<p>事实上，了解和学习这部分内容，一方面可以帮助我们更加全面地掌握，Redis 事件驱动框架是如何以事件形式，处理 server 运行过程中面临的请求操作和多种任务的。比如，正常的客户端读写请求是以什么事件、由哪个函数进行处理，以及后台快照任务又是如何及时启动的。</p>

<p>因为事件驱动框架是 Redis server 运行后的核心循环流程，了解它何时用什么函数处理哪种事件，对我们排查 server 运行过程中遇到的问题，是很有帮助的。</p>

<p>另一方面，我们还可以学习到如何在一个框架中，同时处理 IO 事件和时间事件。我们平时开发服务器端程序，经常需要处理周期性任务，而 Redis 关于两类事件的处理实现，就给了我们一个不错的参考。</p>

<p>好，为了对这两类事件有个相对全面的了解，接下来，我们先从事件驱动框架循环流程的数据结构及其初始化开始学起，因为这里面就包含了针对这两类事件的数据结构定义和初始化操作。</p>

<h2 id="aeeventloop-结构体与初始化">aeEventLoop 结构体与初始化</h2>

<p>首先，我们来看下 Redis 事件驱动框架循环流程对应的数据结构 aeEventLoop。这个结构体是在事件驱动框架代码<a href="https://github.com/redis/redis/tree/5.0/src/ae.h" target="_blank">ae.h</a>中定义的，记录了框架循环运行过程中的信息，其中，就包含了记录两类事件的变量，分别是：</p>

<ul>
<li><strong>aeFileEvent 类型的指针 *events，表示 IO 事件</strong>。之所以类型名称为 aeFileEvent，是因为所有的 IO 事件都会用文件描述符进行标识；</li>
<li><strong>aeTimeEvent 类型的指针 *timeEventHead，表示时间事件</strong>，即按一定时间周期触发的事件。</li>
</ul>

<p>此外，aeEventLoop 结构体中还有一个 <strong>aeFiredEvent 类型的指针 *fired</strong>，这个并不是一类专门的事件类型，它只是用来记录已触发事件对应的文件描述符信息。</p>

<p>下面的代码显示了 Redis 中事件循环的结构体定义，你可以看下。</p>

<pre><code class="language-c">typedef struct aeEventLoop {
    …
    aeFileEvent *events; //IO事件数组
    aeFiredEvent *fired; //已触发事件数组
    aeTimeEvent *timeEventHead; //记录时间事件的链表头
    …
    void *apidata; //和API调用接口相关的数据
    aeBeforeSleepProc *beforesleep; //进入事件循环流程前执行的函数
    aeBeforeSleepProc *aftersleep;  //退出事件循环流程后执行的函数
} aeEventLoop;
</code></pre>

<p>了解了 aeEventLoop 结构体后，我们再来看下，这个结构体是如何初始化的，这其中就包括了 IO 事件数组和时间事件链表的初始化。</p>

<h3 id="aecreateeventloop-函数的初始化操作">aeCreateEventLoop 函数的初始化操作</h3>

<p>因为 Redis server 在完成初始化后，就要开始运行事件驱动框架的循环流程，所以，aeEventLoop 结构体在<a href="https://github.com/redis/redis/tree/5.0/src/server.c" target="_blank">server.c</a>的 initServer 函数中，就通过调用 <strong>aeCreateEventLoop 函数</strong>进行初始化了。这个函数的参数只有一个，是 setsize。</p>

<p>下面的代码展示了 initServer 函数中对 aeCreateEventLoop 函数的调用。</p>

<pre><code class="language-c">initServer() {
…
//调用aeCreateEventLoop函数创建aeEventLoop结构体，并赋值给server结构的el变量
server.el = aeCreateEventLoop(server.maxclients+CONFIG_FDSET_INCR);
…
}
</code></pre>

<p>从这里我们可以看到<strong>参数 setsize</strong> 的大小，其实是由 server 结构的 maxclients 变量和宏定义 CONFIG_FDSET_INCR 共同决定的。其中，maxclients 变量的值大小，可以在 Redis 的配置文件 redis.conf 中进行定义，默认值是 1000。而宏定义 CONFIG_FDSET_INCR 的大小，等于宏定义 CONFIG_MIN_RESERVED_FDS 的值再加上 96，如下所示，这里的两个宏定义都是在<a href="https://github.com/redis/redis/blob/5.0/src/server.h" target="_blank">server.h</a>文件中定义的。</p>

<pre><code class="language-c">#define CONFIG_MIN_RESERVED_FDS 32
#define CONFIG_FDSET_INCR (CONFIG_MIN_RESERVED_FDS+96)
</code></pre>

<p>好了，到这里，你可能有疑问了：aeCreateEventLoop 函数的参数 setsize，设置为最大客户端数量加上一个宏定义值，可是<strong>这个参数有什么用呢</strong>？这就和 aeCreateEventLoop 函数具体执行的初始化操作有关了。</p>

<p>接下来，我们就来看下 aeCreateEventLoop 函数执行的操作，大致可以分成以下三个步骤。</p>

<p><strong>第一步，aeCreateEventLoop 函数会创建一个 aeEventLoop 结构体类型的变量 eventLoop</strong>。然后，该函数会给 eventLoop 的成员变量分配内存空间，比如，按照传入的参数 setsize，给 IO 事件数组和已触发事件数组分配相应的内存空间。此外，该函数还会给 eventLoop 的成员变量赋初始值。</p>

<p><strong>第二步，aeCreateEventLoop 函数会调用 aeApiCreate 函数</strong>。aeApiCreate 函数封装了操作系统提供的 IO 多路复用函数，假设 Redis 运行在 Linux 操作系统上，并且 IO 多路复用机制是 epoll，那么此时，aeApiCreate 函数就会调用 epoll_create 创建 epoll 实例，同时会创建 epoll_event 结构的数组，数组大小等于参数 setsize。</p>

<p>这里你需要注意，aeApiCreate 函数是把创建的 epoll 实例描述符和 epoll_event 数组，保存在了 aeApiState 结构体类型的变量 state，如下所示：</p>

<pre><code class="language-c">typedef struct aeApiState {  //aeApiState结构体定义
    int epfd;   //epoll实例的描述符
    struct epoll_event *events;   //epoll_event结构体数组，记录监听事件
} aeApiState;

static int aeApiCreate(aeEventLoop *eventLoop) {
    aeApiState *state = zmalloc(sizeof(aeApiState));
    ...
    //将epoll_event数组保存在aeApiState结构体变量state中
    state-&gt;events = zmalloc(sizeof(struct epoll_event)*eventLoop-&gt;setsize);
    ...
    //将epoll实例描述符保存在aeApiState结构体变量state中
    state-&gt;epfd = epoll_create(1024); 
</code></pre>

<p>紧接着，aeApiCreate 函数把 state 变量赋值给 eventLoop 中的 apidata。这样一来，eventLoop 结构体中就有了 epoll 实例和 epoll_event 数组的信息，这样就可以用来基于 epoll 创建和处理事件了。我一会儿还会给你具体介绍。</p>

<pre><code class="language-c">eventLoop-&gt;apidata = state;
</code></pre>

<p><strong>第三步，aeCreateEventLoop 函数会把所有网络 IO 事件对应文件描述符的掩码，初始化为 AE_NONE，表示暂时不对任何事件进行监听。</strong></p>

<p>我把 aeCreateEventLoop 函数的主要部分代码放在这里，你可以看下。</p>

<pre><code class="language-c">aeEventLoop *aeCreateEventLoop(int setsize) {
    aeEventLoop *eventLoop;
    int i;
   
    //给eventLoop变量分配内存空间
  if ((eventLoop = zmalloc(sizeof(*eventLoop))) == NULL) goto err;
  //给IO事件、已触发事件分配内存空间
    eventLoop-&gt;events = zmalloc(sizeof(aeFileEvent)*setsize);
    eventLoop-&gt;fired = zmalloc(sizeof(aeFiredEvent)*setsize);
    …
    eventLoop-&gt;setsize = setsize;
    eventLoop-&gt;lastTime = time(NULL);
    //设置时间事件的链表头为NULL
    eventLoop-&gt;timeEventHead = NULL;
  …
  //调用aeApiCreate函数，去实际调用操作系统提供的IO多路复用函数
  if (aeApiCreate(eventLoop) == -1) goto err;
   
    //将所有网络IO事件对应文件描述符的掩码设置为AE_NONE
    for (i = 0; i &lt; setsize; i++)
        eventLoop-&gt;events[i].mask = AE_NONE;
    return eventLoop;
 
    //初始化失败后的处理逻辑，
    err:
    …
}
</code></pre>

<p>好，那么从 aeCreateEventLoop 函数的执行流程中，我们其实可以看到以下<strong>两个关键点</strong>：</p>

<ul>
<li>事件驱动框架监听的 IO 事件数组大小就等于参数 setsize，这样决定了和 Redis server 连接的客户端数量。所以，当你遇到客户端连接 Redis 时报错“max number of clients reached”，你就可以去 redis.conf 文件修改 maxclients 配置项，以扩充框架能监听的客户端数量。</li>
<li>当使用 Linux 系统的 epoll 机制时，框架循环流程初始化操作，会通过 aeApiCreate 函数创建 epoll_event 结构数组，并调用 epoll_create 函数创建 epoll 实例，这都是使用 epoll 机制的准备工作要求，你也可以再回顾下第 9 讲中对 epoll 使用的介绍。</li>
</ul>

<p>到这里，框架就可以创建和处理具体的 IO 事件和时间事件了。所以接下来，我们就先来了解下 IO 事件及其处理机制。</p>

<h2 id="io-事件处理">IO 事件处理</h2>

<p>事实上，Redis 的 IO 事件主要包括三类，分别是可读事件、可写事件和屏障事件。</p>

<p>其中，可读事件和可写事件其实比较好理解，也就是对应于 Redis 实例，我们可以<strong>从客户端读取数据或是向客户端写入数据</strong>。而屏障事件的主要作用是用来<strong>反转事件的处理顺序</strong>。比如在默认情况下，Redis 会先给客户端返回结果，但是如果面临需要把数据尽快写入磁盘的情况，Redis 就会用到屏障事件，把写数据和回复客户端的顺序做下调整，先把数据落盘，再给客户端回复。</p>

<p>我在上节课给你介绍过，在 Redis 源码中，IO 事件的数据结构是 aeFileEvent 结构体，IO 事件的创建是通过 aeCreateFileEvent 函数来完成的。下面的代码展示了 aeFileEvent 结构体的定义，你可以再回顾下：</p>

<pre><code class="language-c">typedef struct aeFileEvent {
    int mask; //掩码标记，包括可读事件、可写事件和屏障事件
    aeFileProc *rfileProc;   //处理可读事件的回调函数
    aeFileProc *wfileProc;   //处理可写事件的回调函数
    void *clientData;  //私有数据
} aeFileEvent;
</code></pre>

<p>而对于 aeCreateFileEvent 函数来说，在上节课我们已经了解了它是通过 aeApiAddEvent 函数来完成事件注册的。那么接下来，我们再从代码级别看下它是如何执行的，这可以帮助我们更加透彻地理解，事件驱动框架对 IO 事件监听是如何基于 epoll 机制对应封装的。</p>

<h3 id="io-事件创建">IO 事件创建</h3>

<p>首先，我们来看 aeCreateFileEvent 函数的原型定义，如下所示：</p>

<pre><code class="language-c">int aeCreateFileEvent(aeEventLoop *eventLoop, int fd, int mask, aeFileProc *proc, void *clientData)
</code></pre>

<p>这个函数的参数有 5 个，分别是循环流程结构体 *eventLoop、IO 事件对应的文件描述符 fd、事件类型掩码 mask、事件处理回调函数*proc，以及事件私有数据*clientData。</p>

<p>因为循环流程结构体*eventLoop中有 IO 事件数组，这个数组的元素是 aeFileEvent 类型，所以，每个数组元素都对应记录了一个文件描述符（比如一个套接字）相关联的监听事件类型和回调函数。</p>

<p>aeCreateFileEvent 函数会先根据传入的文件描述符 fd，在 eventLoop 的 IO 事件数组中，获取该描述符关联的 IO 事件指针变量*fe，如下所示：</p>

<pre><code class="language-c">aeFileEvent *fe = &amp;eventLoop-&gt;events[fd];
</code></pre>

<p>紧接着，aeCreateFileEvent 函数会调用 aeApiAddEvent 函数，添加要监听的事件：</p>

<pre><code class="language-c">if (aeApiAddEvent(eventLoop, fd, mask) == -1)
   return AE_ERR;
</code></pre>

<p>aeApiAddEvent 函数实际上会调用操作系统提供的 IO 多路复用函数，来完成事件的添加。我们还是假设 Redis 实例运行在使用 epoll 机制的 Linux 上，那么 aeApiAddEvent 函数就会调用 epoll_ctl 函数，添加要监听的事件。我在第 9 讲中其实已经给你介绍过 epoll_ctl 函数，这个函数会接收 4 个参数，分别是：</p>

<ul>
<li>epoll 实例；</li>
<li>要执行的操作类型（是添加还是修改）；</li>
<li>要监听的文件描述符；</li>
<li>epoll_event 类型变量。</li>
</ul>

<p>那么，<strong>这个调用过程是如何准备 epoll_ctl 函数需要的参数，从而完成执行的呢？</strong></p>

<p>首先，epoll 实例是我刚才给你介绍的 aeCreateEventLoop 函数，它是通过调用 aeApiCreate 函数来创建的，保存在了 eventLoop 结构体的 apidata 变量中，类型是 aeApiState。所以，aeApiAddEvent 函数会先获取该变量，如下所示：</p>

<pre><code class="language-c">static int aeApiAddEvent(aeEventLoop *eventLoop, int fd, int mask) {
    //从eventLoop结构体中获取aeApiState变量，里面保存了epoll实例
  aeApiState *state = eventLoop-&gt;apidata;
    ...
 }
</code></pre>

<p>其次，对于要执行的操作类型的设置，aeApiAddEvent 函数会根据传入的文件描述符 fd，在 eventLoop 结构体中 IO 事件数组中查找该 fd。因为 IO 事件数组的每个元素，都对应了一个文件描述符，而该数组初始化时，每个元素的值都设置为了 AE_NONE。</p>

<p>所以，如果要监听的文件描述符 fd 在数组中的类型不是 AE_NONE，则表明该描述符已做过设置，那么操作类型就是修改操作，对应 epoll 机制中的宏定义 EPOLL_CTL_MOD。否则，操作类型就是添加操作，对应 epoll 机制中的宏定义 EPOLL_CTL_ADD。这部分代码如下所示：</p>

<pre><code class="language-c">//如果文件描述符fd对应的IO事件已存在，则操作类型为修改，否则为添加
 int op = eventLoop-&gt;events[fd].mask == AE_NONE ?
            EPOLL_CTL_ADD : EPOLL_CTL_MOD;
</code></pre>

<p>第三，epoll_ctl 函数需要的监听文件描述符，就是 aeApiAddEvent 函数接收到的参数 fd。</p>

<p>最后，epoll_ctl 函数还需要一个 epoll_event 类型变量，因此 aeApiAddEvent 函数在调用 epoll_ctl 函数前，会新创建 epoll_event 类型<strong>变量 ee。</strong>然后，aeApiAddEvent 函数会设置变量 ee 中的监听事件类型和监听文件描述符。</p>

<p>aeApiAddEvent 函数的参数 mask，表示的是要监听的事件类型掩码。所以，aeApiAddEvent 函数会根据掩码值是可读（AE_READABLE）或可写（AE_WRITABLE）事件，来设置 ee 监听的事件类型是 EPOLLIN 还是 EPOLLOUT。这样一来，Redis 事件驱动框架中的读写事件就能够和 epoll 机制中的读写事件对应上来。下面的代码展示了这部分逻辑，你可以看下。</p>

<pre><code class="language-c">…
struct epoll_event ee = {0}; //创建epoll_event类型变量
…
//将可读或可写IO事件类型转换为epoll监听的类型EPOLLIN或EPOLLOUT
if (mask &amp; AE_READABLE) ee.events |= EPOLLIN;
if (mask &amp; AE_WRITABLE) ee.events |= EPOLLOUT;
ee.data.fd = fd;  //将要监听的文件描述符赋值给ee
…  
</code></pre>

<p>好了，到这里，aeApiAddEvent 函数就准备好了 epoll 实例、操作类型、监听文件描述符以及 epoll_event 类型变量，然后，它就会调用 epoll_ctl 开始实际创建监听事件了，如下所示：</p>

<pre><code class="language-c">static int aeApiAddEvent(aeEventLoop *eventLoop, int fd, int mask) {
...
//调用epoll_ctl实际创建监听事件
if (epoll_ctl(state-&gt;epfd,op,fd,&amp;ee) == -1) return -1;
return 0;
}
</code></pre>

<p>了解了这些代码后，我们可以学习到事件驱动框架是如何基于 epoll，封装实现了 IO 事件的创建。那么，在 Redis server 启动运行后，最开始监听的 IO 事件是可读事件，对应于客户端的连接请求。具体是 initServer 函数调用了 aeCreateFileEvent 函数，创建可读事件，并设置回调函数为 acceptTcpHandler，用来处理客户端连接。这部分内容，你也可以再回顾下第 10 讲。</p>

<p>接下来，我们再来看下一旦有了客户端连接请求后，IO 事件具体是如何处理的呢？</p>

<h3 id="读事件处理">读事件处理</h3>

<p>当 Redis server 接收到客户端的连接请求时，就会使用注册好的 <strong>acceptTcpHandler 函数</strong>进行处理。</p>

<p>acceptTcpHandler 函数是在<a href="https://github.com/redis/redis/tree/5.0/src/networking.c" target="_blank">networking.c</a>文件中，它会接受客户端连接，并创建已连接套接字 cfd。然后，acceptCommonHandler 函数（在 networking.c 文件中）会被调用，同时，刚刚创建的已连接套接字 cfd 会作为参数，传递给 acceptCommonHandler 函数。</p>

<p>acceptCommonHandler 函数会调用 createClient 函数（在 networking.c 文件中）创建客户端。而在 createClient 函数中，我们就会看到，aeCreateFileEvent 函数被再次调用了。</p>

<p>此时，aeCreateFileEvent 函数会针对已连接套接字上，创建监听事件，类型为 AE_READABLE，回调函数是 readQueryFromClient（在 networking.c 文件中）。</p>

<p>好了，到这里，事件驱动框架就增加了对一个客户端已连接套接字的监听。一旦客户端有请求发送到 server，框架就会回调 readQueryFromClient 函数处理请求。这样一来，客户端请求就能通过事件驱动框架进行处理了。</p>

<p>下面代码展示了 createClient 函数调用 aeCreateFileEvent 的过程，你可以看下。</p>

<pre><code class="language-c">client *createClient(int fd) {
…
if (fd != -1) {
        …
        //调用aeCreateFileEvent，监听读事件，对应客户端读写请求，使用readQueryFromclient回调函数处理
        if (aeCreateFileEvent(server.el,fd,AE_READABLE,
            readQueryFromClient, c) == AE_ERR)
        {
            close(fd);
            zfree(c);
            return NULL;
        } }
…
}
</code></pre>

<p>为了便于你掌握从监听客户端连接请求到监听客户端常规读写请求的事件创建过程，我画了下面这张图，你可以看下。</p>

<p><img src="assets/5212d6d61a9e5d8aaefc88497fd5aee9-20221013235205-sj0vdny.jpg" alt="" /></p>

<p>了解了事件驱动框架中的读事件处理之后，我们再来看下写事件的处理。</p>

<h3 id="写事件处理">写事件处理</h3>

<p>Redis 实例在收到客户端请求后，会在处理客户端命令后，将要返回的数据写入客户端输出缓冲区。下图就展示了这个过程的函数调用逻辑：</p>

<p><img src="assets/3107452361f9366189e4d4080f3153df-20221013235205-w4uhglw.jpg" alt="" /></p>

<p>而在 Redis 事件驱动框架每次循环进入事件处理函数前，也就是在框架主函数 aeMain 中调用 aeProcessEvents，来处理监听到的已触发事件或是到时的时间事件之前，都会调用 server.c 文件中的 <strong>beforeSleep 函数</strong>，进行一些任务处理，这其中就包括了调用 handleClientsWithPendingWrites 函数，它会将 Redis sever 客户端缓冲区中的数据写回客户端。</p>

<p>下面给出的代码是事件驱动框架的主函数 aeMain。在该函数每次调用 aeProcessEvents 函数前，就会调用 beforeSleep 函数，你可以看下。</p>

<pre><code class="language-c">void aeMain(aeEventLoop *eventLoop) {
    eventLoop-&gt;stop = 0;
  while (!eventLoop-&gt;stop) {
      //如果beforeSleep函数不为空，则调用beforeSleep函数
        if (eventLoop-&gt;beforesleep != NULL)
            eventLoop-&gt;beforesleep(eventLoop);
        //调用完beforeSleep函数，再处理事件
        aeProcessEvents(eventLoop, AE_ALL_EVENTS|AE_CALL_AFTER_SLEEP);
    }
}
</code></pre>

<p>这里你要知道，beforeSleep 函数调用的 handleClientsWithPendingWrites 函数，会遍历每一个待写回数据的客户端，然后调用 writeToClient 函数，将客户端输出缓冲区中的数据写回。下面这张图展示了这个流程，你可以看下。</p>

<p><img src="assets/18c8bfaab29a890fb15edcca3971bb05-20221013235205-dd3oiep.jpg" alt="" /></p>

<p>不过，如果输出缓冲区的数据还没有写完，此时，handleClientsWithPendingWrites 函数就会<strong>调用 aeCreateFileEvent 函数，创建可写事件</strong>，并设置回调函数 sendReplyToClient。sendReplyToClient 函数里面会调用 writeToClient 函数写回数据。</p>

<p>下面的代码展示了 handleClientsWithPendingWrite 函数的基本流程，你可以看下。</p>

<pre><code class="language-c">int handleClientsWithPendingWrites(void) {
    listIter li;
  listNode *ln;
  …
    //获取待写回的客户端列表
  listRewind(server.clients_pending_write,&amp;li);
  //遍历每一个待写回的客户端
  while((ln = listNext(&amp;li))) {
     client *c = listNodeValue(ln);
     …
     //调用writeToClient将当前客户端的输出缓冲区数据写回
     if (writeToClient(c-&gt;fd,c,0) == C_ERR) continue;
     //如果还有待写回数据
     if (clientHasPendingReplies(c)) {
              int ae_flags = AE_WRITABLE;
              //创建可写事件的监听，以及设置回调函数
               if (aeCreateFileEvent(server.el, c-&gt;fd, ae_flags,
                  sendReplyToClient, c) == AE_ERR)
              {
                     …
              }
    } }
}
</code></pre>

<p>好了，我们刚才了解的是读写事件对应的回调处理函数。实际上，为了能及时处理这些事件，Redis 事件驱动框架的 aeMain 函数还会循环<strong>调用 aeProcessEvents 函数，来检测已触发的事件，并调用相应的回调函数进行处理。</strong></p>

<p>从 aeProcessEvents 函数的代码中，我们可以看到该函数会调用 aeApiPoll 函数，查询监听的文件描述符中，有哪些已经就绪。一旦有描述符就绪，aeProcessEvents 函数就会根据事件的可读或可写类型，调用相应的回调函数进行处理。aeProcessEvents 函数调用的基本流程如下所示：</p>

<pre><code class="language-c">int aeProcessEvents(aeEventLoop *eventLoop, int flags){
…
//调用aeApiPoll获取就绪的描述符
numevents = aeApiPoll(eventLoop, tvp);
…
for (j = 0; j &lt; numevents; j++) {
  aeFileEvent *fe = &amp;eventLoop-&gt;events[eventLoop-&gt;fired[j].fd];
  …
    //如果触发的是可读事件，调用事件注册时设置的读事件回调处理函数
  if (!invert &amp;&amp; fe-&gt;mask &amp; mask &amp; AE_READABLE) {
        fe-&gt;rfileProc(eventLoop,fd,fe-&gt;clientData,mask);
                  fired++;
  }
    //如果触发的是可写事件，调用事件注册时设置的写事件回调处理函数
  if (fe-&gt;mask &amp; mask &amp; AE_WRITABLE) {
                  if (!fired || fe-&gt;wfileProc != fe-&gt;rfileProc) {
                      fe-&gt;wfileProc(eventLoop,fd,fe-&gt;clientData,mask);
                      fired++;
                  }
              }
  …
  } }
  …
}
</code></pre>

<p>到这里，我们就了解了 IO 事件的创建函数 aeCreateFileEvent，以及在处理客户端请求时对应的读写事件和它们的处理函数。那么接下来，我们再来看看事件驱动框架中的时间事件是怎么创建和处理的。</p>

<h2 id="时间事件处理">时间事件处理</h2>

<p>其实，相比于 IO 事件有可读、可写、屏障类型，以及不同类型 IO 事件有不同回调函数来说，时间事件的处理就比较简单了。下面，我们就来分别学习下它的定义、创建、回调函数和触发处理。</p>

<h3 id="时间事件定义">时间事件定义</h3>

<p>首先，我们来看下时间事件的结构体定义，代码如下所示：</p>

<pre><code class="language-c">typedef struct aeTimeEvent {
    long long id; //时间事件ID
    long when_sec; //事件到达的秒级时间戳
    long when_ms; //事件到达的毫秒级时间戳
    aeTimeProc *timeProc; //时间事件触发后的处理函数
    aeEventFinalizerProc *finalizerProc;  //事件结束后的处理函数
    void *clientData; //事件相关的私有数据
    struct aeTimeEvent *prev;  //时间事件链表的前向指针
    struct aeTimeEvent *next;  //时间事件链表的后向指针
} aeTimeEvent;
</code></pre>

<p>时间事件结构体中主要的变量，包括以秒记录和以毫秒记录的时间事件触发时的时间戳 when_sec 和 when_ms，以及时间事件触发后的处理函数*timeProc。另外，在时间事件的结构体中，还包含了前向和后向指针*prev和*next，这表明<strong>时间事件是以链表的形式组织起来的</strong>。</p>

<p>在了解了时间事件结构体的定义以后，我们接着来看下，时间事件是如何创建的。</p>

<h3 id="时间事件创建">时间事件创建</h3>

<p>与 IO 事件创建使用 aeCreateFileEvent 函数类似，<strong>时间事件的创建函数是 aeCreateTimeEvent 函数</strong>。这个函数的原型定义如下所示：</p>

<pre><code class="language-c">long long aeCreateTimeEvent(aeEventLoop *eventLoop, long long milliseconds, aeTimeProc *proc, void *clientData, aeEventFinalizerProc *finalizerProc)
</code></pre>

<p>在它的参数中，有两个需要我们重点了解下，以便于我们理解时间事件的处理。一个是 <strong>milliseconds</strong>，这是所创建时间事件的触发时间距离当前时间的时长，是用毫秒表示的。另一个是 *<strong>proc</strong>，这是所创建时间事件触发后的回调函数。</p>

<p>aeCreateTimeEvent 函数的执行逻辑不复杂，主要就是创建一个时间事件的<strong>变量 te</strong>，对它进行初始化，并把它插入到框架循环流程结构体 eventLoop 中的时间事件链表中。在这个过程中，aeCreateTimeEvent 函数会<strong>调用 aeAddMillisecondsToNow 函数</strong>，根据传入的 milliseconds 参数，计算所创建时间事件具体的触发时间戳，并赋值给 te。</p>

<p>实际上，Redis server 在初始化时，除了创建监听的 IO 事件外，也会调用 aeCreateTimeEvent 函数创建时间事件。下面代码显示了 initServer 函数对 aeCreateTimeEvent 函数的调用：</p>

<pre><code class="language-c">initServer() {
…
//创建时间事件
if (aeCreateTimeEvent(server.el, 1, serverCron, NULL, NULL) == AE_ERR){
… //报错信息
}
}
</code></pre>

<p>从代码中，我们可以看到，<strong>时间事件触发后的回调函数是 serverCron</strong>。所以接下来，我们就来了解下 serverCron 函数。</p>

<h3 id="时间事件回调函数">时间事件回调函数</h3>

<p>serverCron 函数是在 server.c 文件中实现的。<strong>一方面</strong>，它会顺序调用一些函数，来实现时间事件被触发后，执行一些后台任务。比如，serverCron 函数会检查是否有进程结束信号，若有就执行 server 关闭操作。serverCron 会调用 databaseCron 函数，处理过期 key 或进行 rehash 等。你可以参考下面给出的代码：	</p>

<pre><code class="language-c">...
//如果收到进程结束信号，则执行server关闭操作
 if (server.shutdown_asap) {
        if (prepareForShutdown(SHUTDOWN_NOFLAGS) == C_OK) exit(0);
        ...
 }
...
clientCron();  //执行客户端的异步操作
databaseCron(); //执行数据库的后台操作
...
</code></pre>

<p><strong>另一方面</strong>，serverCron 函数还会以不同的频率周期性执行一些任务，这是通过执行宏 run_with_period 来实现的。</p>

<p>run_with_period 宏定义如下，该宏定义会根据 Redis 实例配置文件 redis.conf 中定义的 hz 值，来判断参数 <em>ms</em> 表示的时间戳是否到达。一旦到达，serverCron 就可以执行相应的任务了。</p>

<pre><code class="language-c">#define run_with_period(_ms_) if ((_ms_ &lt;= 1000/server.hz) || !(server.cronloops%((_ms_)/(1000/server.hz))))
</code></pre>

<p>比如，serverCron 函数中会以 1 秒 1 次的频率，检查 AOF 文件是否有写错误。如果有的话，serverCron 就会调用 flushAppendOnlyFile 函数，再次刷回 AOF 文件的缓存数据。下面的代码展示了这一周期性任务：</p>

<pre><code class="language-c">serverCron() {
   …
   //每1秒执行1次，检查AOF是否有写错误
   run_with_period(1000) {
        if (server.aof_last_write_status == C_ERR)
            flushAppendOnlyFile(0);
    }
   …
}
</code></pre>

<p>如果你想了解更多的周期性任务，可以再详细阅读下 serverCron 函数中，以 run_with_period 宏定义包含的代码块。</p>

<p>好了，了解了时间事件触发后的回调函数 serverCron，我们最后来看下，时间事件是如何触发处理的。</p>

<h3 id="时间事件的触发处理">时间事件的触发处理</h3>

<p>其实，时间事件的检测触发比较简单，事件驱动框架的 aeMain 函数会循环调用 aeProcessEvents 函数，来处理各种事件。而 aeProcessEvents 函数在执行流程的最后，会<strong>调用 processTimeEvents 函数处理相应到时的任务</strong>。</p>

<pre><code class="language-c">aeProcessEvents(){
…
//检测时间事件是否触发
if (flags &amp; AE_TIME_EVENTS)
        processed += processTimeEvents(eventLoop);
…
}
</code></pre>

<p>那么，具体到 proecessTimeEvent 函数来说，它的基本流程就是从时间事件链表上逐一取出每一个事件，然后根据当前时间判断该事件的触发时间戳是否已满足。如果已满足，那么就调用该事件对应的回调函数进行处理。这样一来，周期性任务就能在不断循环执行的 aeProcessEvents 函数中，得到执行了。</p>

<p>下面的代码显示了 processTimeEvents 函数的基本流程，你可以再看下。</p>

<pre><code class="language-c">static int processTimeEvents(aeEventLoop *eventLoop) {
...
te = eventLoop-&gt;timeEventHead;  //从时间事件链表中取出事件
while(te) {
   ...
  aeGetTime(&amp;now_sec, &amp;now_ms);  //获取当前时间
  if (now_sec &gt; te-&gt;when_sec || (now_sec == te-&gt;when_sec &amp;&amp; now_ms &gt;= te-&gt;when_ms))   //如果当前时间已经满足当前事件的触发时间戳
  {
     ...
    retval = te-&gt;timeProc(eventLoop, id, te-&gt;clientData); //调用注册的回调函数处理
    ...
  }
  te = te-&gt;next;   //获取下一个时间事件
  ...
}
</code></pre>

<h2 id="小结">小结</h2>

<p>这节课，我给你介绍了 Redis 事件驱动框架中的两类事件：IO 事件和时间事件。</p>

<p>对于 IO 事件来说，它可以进一步分成可读、可写和屏障事件。因为可读、可写事件在 Redis 和客户端通信处理请求过程中使用广泛，所以今天我们重点学习了这两种 IO 事件。当 Redis server 创建 Socket 后，就会注册可读事件，并使用 acceptTCPHandler 回调函数处理客户端的连接请求。</p>

<p>当 server 和客户端完成连接建立后，server 会在已连接套接字上监听可读事件，并使用 readQueryFromClient 函数处理客户端读写请求。这里，你需要再注意下，<strong>无论客户端发送的请求是读或写操作，对于 server 来说，都是要读取客户端的请求并解析处理</strong>。所以，server 在客户端的已连接套接字上注册的是可读事件。</p>

<p>而当实例需要向客户端写回数据时，实例会在事件驱动框架中注册可写事件，并使用 sendReplyToClient 作为回调函数，将缓冲区中数据写回客户端。我总结了一张表格，以便你再回顾下 IO 事件和相应套接字、回调函数的对应关系。</p>

<p><img src="assets/4b354d1c4fd6eed66b58c14689018a0c-20221013235205-ncd78c2.jpg" alt="" /></p>

<p>然后，对于时间事件来说，它主要是用于在事件驱动框架中注册一些周期性执行的任务，以便 Redis server 进行后台处理。时间事件的回调函数是 serverCron 函数，你可以做进一步阅读了解其中的具体任务。</p>

<p>好了，从第 9 讲开始，我用了 3 节课，向你介绍 Redis 事件驱动框架的运行机制，本质上来说，事件驱动框架是基于操作系统提供的 IO 多路复用机制进行了封装，并加上了时间事件的处理。这是一个非常经典的事件框架实现，我希望你可以学习并掌握好它，然后用在你自己的系统开发中。</p>

<h2 id="每课一问">每课一问</h2>

<p>已知，Redis 事件驱动框架的 aeApiCreate、aeApiAddEvent 等等这些函数，是对操作系统提供的 IO 多路复用函数进行了封装，具体的 IO 多路复用函数分别是在<a href="https://github.com/redis/redis/tree/5.0/src/ae_epoll.c" target="_blank">ae_epoll.c</a>，<a href="https://github.com/redis/redis/tree/5.0/src/ae_evport.c" target="_blank">ae_evport.c</a>，<a href="https://github.com/redis/redis/tree/5.0/src/ae_kqueue.c" target="_blank">ae_kqueue.c</a>，<a href="https://github.com/redis/redis/blob/5.0/src/ae_select.c" target="_blank">ae_select.c</a>四个代码文件中定义的。那么你知道，Redis 在调用 aeApiCreate，aeApiAddEvent 这些函数时，是根据什么条件来决定，具体调用哪个文件中的 IO 多路复用函数的吗？</p>

<p>‍</p>
</div>
                        </div>
                        <div>
                            <div id="prePage" style="float: left">

                            </div>
                            <div id="nextPage" style="float: right">

                            </div>
                        </div>

                    </div>
                </div>
            </div>
            <div class="copyright">
                <hr />
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#84e8e8e8bdb0b5b5b4b3c4e3e9e5ede8aae7ebe9" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f123a5698b1ede4',t:'MTczNDA1MzgzNS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;&#32;Redis事件驱动框架（上）：何时使用select、poll、epoll？>
        <link rel="icon" href="/static/favicon.png">
        <title>09  Redis事件驱动框架（上）：何时使用select、poll、epoll？ </title>
        
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
                            <h1 id="title" data-id="09  Redis事件驱动框架（上）：何时使用select、poll、epoll？" class="title">09  Redis事件驱动框架（上）：何时使用select、poll、epoll？</h1>
                            <div><p>Redis 作为一个 Client-Server 架构的数据库，其源码中少不了用来实现网络通信的部分。而你应该也清楚，通常系统实现网络通信的基本方法是<strong>使用 Socket 编程模型</strong>，包括创建 Socket、监听端口、处理连接请求和读写请求。但是，由于基本的 Socket 编程模型一次只能处理一个客户端连接上的请求，所以当要处理高并发请求时，一种方案就是使用多线程，让每个线程负责处理一个客户端的请求。</p>

<p>而 Redis 负责客户端请求解析和处理的线程只有一个，那么如果直接采用基本 Socket 模型，就会影响 Redis 支持高并发的客户端访问。</p>

<p>因此，为了实现高并发的网络通信，我们常用的 Linux 操作系统，就提供了 select、poll 和 epoll 三种编程模型，而在 Linux 上运行的 Redis，通常就会采用其中的 <strong>epoll 模型</strong>来进行网络通信。</p>

<p>这里你可能就要问了：<strong>为啥 Redis 通常会选择 epoll 模型呢？这三种编程模型之间有什么区别？</strong>如果我们自己要开发高并发的服务器处理程序时，应该如何选择使用呢？</p>

<p>今天这节课，我就来和你聊聊，Redis 在高并发网络通信编程模型上的选择和设计思想。通过这节课的学习，你可以掌握 select、poll 和 epoll 三种模型的工作机制和使用方法。了解这些内容，一方面可以帮助你理解 Redis 整体网络通信框架的工作基础，另一方面，也可以让你学会如何进行高并发网络通信的开发。</p>

<p>那么，要想理解 select、poll 和 epoll 的优势，我们需要有个对比基础，也就是基本的 Socket 编程模型。所以接下来，我们就先来了解下基本的 Socket 编程模型，以及它的不足之处。</p>

<h2 id="为什么-redis-不使用基本的-socket-编程模型">为什么 Redis 不使用基本的 Socket 编程模型？</h2>

<p>刚刚我们说过，使用 Socket 模型实现网络通信时，需要经过创建 Socket、监听端口、处理连接和读写请求等多个步骤，现在我们就来具体了解下这些步骤中的关键操作，以此帮助我们分析 Socket 模型中的不足。</p>

<p>首先，当我们需要让服务器端和客户端进行通信时，可以在服务器端通过以下三步，来创建监听客户端连接的监听套接字（Listening Socket）：</p>

<ol>
<li>调用 socket 函数，创建一个套接字。我们通常把这个套接字称为主动套接字（Active Socket）；</li>
<li>调用 bind 函数，将主动套接字和当前服务器的 IP 和监听端口进行绑定；</li>
<li>调用 listen 函数，将主动套接字转换为监听套接字，开始监听客户端的连接。</li>
</ol>

<p><img src="assets/eaf5b29b824994a6e9e3bc5bfdeb1a05-20221013235113-d684zhc.jpg" alt="" /></p>

<p>在完成上述三步之后，服务器端就可以接收客户端的连接请求了。为了能及时地收到客户端的连接请求，我们可以运行一个循环流程，在该流程中调用 accept 函数，用于接收客户端连接请求。</p>

<p>这里你需要注意的是，accept 函数是阻塞函数，也就是说，如果此时一直没有客户端连接请求，那么，服务器端的执行流程会一直阻塞在 accept 函数。一旦有客户端连接请求到达，accept 将不再阻塞，而是处理连接请求，和客户端建立连接，并返回已连接套接字（Connected Socket）。</p>

<p>最后，服务器端可以通过调用 recv 或 send 函数，在刚才返回的已连接套接字上，接收并处理读写请求，或是将数据发送给客户端。</p>

<p>下面的代码展示了这一过程，你可以看下。</p>

<pre><code class="language-c">listenSocket = socket(); //调用socket系统调用创建一个主动套接字
bind(listenSocket);  //绑定地址和端口
listen(listenSocket); //将默认的主动套接字转换为服务器使用的被动套接字，也就是监听套接字
while (1) { //循环监听是否有客户端连接请求到来
   connSocket = accept(listenSocket); //接受客户端连接
   recv(connsocket); //从客户端读取数据，只能同时处理一个客户端
   send(connsocket); //给客户端返回数据，只能同时处理一个客户端
}
</code></pre>

<p>不过，从上述代码中，你可能会发现，虽然它能够实现服务器端和客户端之间的通信，但是程序每调用一次 accept 函数，只能处理一个客户端连接。因此，如果想要处理多个并发客户端的请求，我们就需要使用<strong>多线程</strong>的方法，来处理通过 accept 函数建立的多个客户端连接上的请求。</p>

<p>使用这种方法后，我们需要在 accept 函数返回已连接套接字后，创建一个线程，并将已连接套接字传递给创建的线程，由该线程负责这个连接套接字上后续的数据读写。同时，服务器端的执行流程会再次调用 accept 函数，等待下一个客户端连接。</p>

<p>以下给出的示例代码，就展示了使用多线程来提升服务器端的并发客户端处理能力：</p>

<pre><code class="language-c">listenSocket = socket(); //调用socket系统调用创建一个主动套接字
bind(listenSocket);  //绑定地址和端口
listen(listenSocket); //将默认的主动套接字转换为服务器使用的被动套接字，即监听套接字
while (1) { //循环监听是否有客户端连接到来
   connSocket = accept(listenSocket); //接受客户端连接，返回已连接套接字
   pthread_create(processData, connSocket); //创建新线程对已连接套接字进行处理
   
}

//处理已连接套接字上的读写请求
processData(connSocket){
   recv(connsocket); //从客户端读取数据，只能同时处理一个客户端
   send(connsocket); //给客户端返回数据，只能同时处理一个客户端
}
</code></pre>

<p>不过，虽然这种方法能提升服务器端的并发处理能力，遗憾的是，<strong>Redis 的主执行流程是由一个线程在执行，无法使用多线程的方式来提升并发处理能力。</strong>所以，该方法对 Redis 并不起作用。</p>

<p>那么，还有没有什么其他方法，能帮助 Redis 提升并发客户端的处理能力呢？</p>

<p>这就要用到操作系统提供的 <strong>IO 多路复用功能</strong>了。在基本的 Socket 编程模型中，accept 函数只能在一个监听套接字上监听客户端的连接，recv 函数也只能在一个已连接套接字上，等待客户端发送的请求。</p>

<p>而 IO 多路复用机制，可以让程序通过调用多路复用函数，同时监听多个套接字上的请求。这里既可以包括监听套接字上的连接请求，也可以包括已连接套接字上的读写请求。这样当有一个或多个套接字上有请求时，多路复用函数就会返回。此时，程序就可以处理这些就绪套接字上的请求，比如读取就绪的已连接套接字上的请求内容。</p>

<p>因为 Linux 操作系统在实际应用中比较广泛，所以这节课，我们主要来学习 Linux 上的 IO 多路复用机制。Linux 提供的 IO 多路复用机制主要有三种，分别是 select、poll 和 epoll。下面，我们就分别来学习下这三种机制的实现思路和使用方法。然后，我们再来看看，为什么 Redis 通常是选择使用 epoll 这种机制来实现网络通信。</p>

<h2 id="使用-select-和-poll-机制实现-io-多路复用">使用 select 和 poll 机制实现 IO 多路复用</h2>

<p>首先，我们来了解下 select 机制的编程模型。</p>

<p>不过在具体学习之前，我们需要知道，对于一种 IO 多路复用机制来说，我们需要掌握哪些要点，这样可以帮助我们快速抓住不同机制的联系与区别。其实，当我们学习 IO 多路复用机制时，我们需要能回答以下问题：</p>

<ul>
<li>第一，多路复用机制会监听套接字上的哪些事件？</li>
<li>第二，多路复用机制可以监听多少个套接字？</li>
<li>第三，当有套接字就绪时，多路复用机制要如何找到就绪的套接字？</li>
</ul>

<h3 id="select-机制与使用">select 机制与使用</h3>

<p>select 机制中的一个重要函数就是 select 函数。对于 select 函数来说，它的参数包括监听的文件描述符数量<strong>nfds、被监听描述符的三个集合*</strong>readfds、<em>__writefds和</em><strong>exceptfds，以及监听时阻塞等待的超时时长*</strong>timeout。下面的代码显示了 select 函数的原型，你可以看下。</p>

<pre><code class="language-c">int select (int __nfds, fd_set *__readfds, fd_set *__writefds, fd_set *__exceptfds, struct timeval *__timeout)
</code></pre>

<p>这里你需要注意的是，Linux 针对每一个套接字都会有一个文件描述符，也就是一个非负整数，用来唯一标识该套接字。所以，在多路复用机制的函数中，Linux 通常会用文件描述符作为参数。有了文件描述符，函数也就能找到对应的套接字，进而进行监听、读写等操作。</p>

<p>所以，select 函数的参数<strong>readfds、</strong>writefds和__exceptfds表示的是，被监听描述符的集合，其实就是被监听套接字的集合。那么，为什么会有三个集合呢？</p>

<p>这就和我刚才提出的第一个问题相关，也就是<strong>多路复用机制会监听哪些事件</strong>。select 函数使用三个集合，表示监听的三类事件，分别是读数据事件（对应<strong>readfds集合）、写数据事件（对应</strong>writefds集合）和异常事件（对应__exceptfds集合）。</p>

<p>我们进一步可以看到，参数 <strong>readfds、</strong>writefds 和 <strong>exceptfds 的类型是 fd_set 结构体，它主要定义部分如下所示。其中，</strong>fd_mask类型是 long int 类型的别名，__FD_SETSIZE 和 __NFDBITS 这两个宏定义的大小默认为 1024 和 32。</p>

<pre><code class="language-c">typedef struct {
   …
   __fd_mask  __fds_bits[__FD_SETSIZE / __NFDBITS];
   …
} fd_set
</code></pre>

<p>所以，fd_set 结构体的定义，其实就是一个 long int 类型的数组，该数组中一共有 32 个元素（1024/32=32），每个元素是 32 位（long int 类型的大小），而每一位可以用来表示一个文件描述符的状态。</p>

<p>好了，了解了 fd_set 结构体的定义，我们就可以回答刚才提出的第二个问题了。select 函数对每一个描述符集合，都可以监听 1024 个描述符。</p>

<p>接下来，我们再来了解下<strong>如何使用 select 机制来实现网络通信</strong>。</p>

<p>首先，我们在调用 select 函数前，可以先创建好传递给 select 函数的描述符集合，然后再创建监听套接字。而为了让创建的监听套接字能被 select 函数监控，我们需要把这个套接字的描述符加入到创建好的描述符集合中。</p>

<p>然后，我们就可以调用 select 函数，并把创建好的描述符集合作为参数传递给 select 函数。程序在调用 select 函数后，会发生阻塞。而当 select 函数检测到有描述符就绪后，就会结束阻塞，并返回就绪的文件描述符个数。</p>

<p>那么此时，我们就可以在描述符集合中查找哪些描述符就绪了。然后，我们对已就绪描述符对应的套接字进行处理。比如，如果是 __readfds 集合中有描述符就绪，这就表明这些就绪描述符对应的套接字上，有读事件发生，此时，我们就在该套接字上读取数据。</p>

<p>而因为 select 函数一次可以监听 1024 个文件描述符的状态，所以 select 函数在返回时，也可能会一次返回多个就绪的文件描述符。这样一来，我们就可以使用一个循环流程，依次对就绪描述符对应的套接字进行读写或异常处理操作。</p>

<p>我也画了张图，展示了使用 select 函数进行网络通信的基本流程，你可以看下。</p>

<p><img src="assets/49b513c6b9f9a440e8883ff93b33b49f-20221013235113-i8hpv49.jpg" alt="" /></p>

<p>下面的代码展示的是使用 select 函数，进行并发客户端处理的关键步骤和主要函数调用：</p>

<pre><code class="language-c">
int sock_fd,conn_fd; //监听套接字和已连接套接字的变量
sock_fd = socket() //创建套接字
bind(sock_fd)   //绑定套接字
listen(sock_fd) //在套接字上进行监听，将套接字转为监听套接字

fd_set rset;  //被监听的描述符集合，关注描述符上的读事件
 
int max_fd = sock_fd

//初始化rset数组，使用FD_ZERO宏设置每个元素为0 
FD_ZERO(&amp;rset);
//使用FD_SET宏设置rset数组中位置为sock_fd的文件描述符为1，表示需要监听该文件描述符
FD_SET(sock_fd,&amp;rset);

//设置超时时间 
struct timeval timeout;
timeout.tv_sec = 3;
timeout.tv_usec = 0;
 
while(1) {
   //调用select函数，检测rset数组保存的文件描述符是否已有读事件就绪，返回就绪的文件描述符个数
   n = select(max_fd+1, &amp;rset, NULL, NULL, &amp;timeout);
 
   //调用FD_ISSET宏，在rset数组中检测sock_fd对应的文件描述符是否就绪
   if (FD_ISSET(sock_fd, &amp;rset)) {
       //如果sock_fd已经就绪，表明已有客户端连接；调用accept函数建立连接
       conn_fd = accept();
       //设置rset数组中位置为conn_fd的文件描述符为1，表示需要监听该文件描述符
       FD_SET(conn_fd, &amp;rset);
   }

   //依次检查已连接套接字的文件描述符
   for (i = 0; i &lt; maxfd; i++) {
        //调用FD_ISSET宏，在rset数组中检测文件描述符是否就绪
       if (FD_ISSET(i, &amp;rset)) {
         //有数据可读，进行读数据处理
       }
   }
}
</code></pre>

<p>不过从刚才的介绍中，你或许会发现 select 函数存在<strong>两个设计上的不足</strong>：</p>

<ul>
<li>首先，select 函数对单个进程能监听的文件描述符数量是有限制的，它能监听的文件描述符个数由 __FD_SETSIZE 决定，默认值是 1024。</li>
<li>其次，当 select 函数返回后，我们需要遍历描述符集合，才能找到具体是哪些描述符就绪了。这个遍历过程会产生一定开销，从而降低程序的性能。</li>
</ul>

<p>所以，为了解决 select 函数受限于 1024 个文件描述符的不足，poll 函数对此做了改进。</p>

<h3 id="poll-机制与使用">poll 机制与使用</h3>

<p>poll 机制的主要函数是 poll 函数，我们先来看下它的原型定义，如下所示：</p>

<pre><code class="language-c">int poll (struct pollfd *__fds, nfds_t __nfds, int __timeout);
</code></pre>

<p>其中，参数 *__fds 是 pollfd 结构体数组，参数 <strong>nfds 表示的是 *</strong>fds 数组的元素个数，而 __timeout 表示 poll 函数阻塞的超时时间。</p>

<p>pollfd 结构体里包含了要监听的描述符，以及该描述符上要监听的事件类型。这个我们可以从 pollfd 结构体的定义中看出来，如下所示。pollfd 结构体中包含了三个成员变量 fd、events 和 revents，分别表示要监听的文件描述符、要监听的事件类型和实际发生的事件类型。</p>

<pre><code class="language-c">struct pollfd {
    int fd;         //进行监听的文件描述符
    short int events;       //要监听的事件类型
    short int revents;      //实际发生的事件类型
};
</code></pre>

<p>pollfd 结构体中要监听和实际发生的事件类型，是通过以下三个宏定义来表示的，分别是 POLLRDNORM、POLLWRNORM 和 POLLERR，它们分别表示可读、可写和错误事件。</p>

<pre><code class="language-c">#define POLLRDNORM  0x040       //可读事件
#define POLLWRNORM  0x100       //可写事件
#define POLLERR     0x008       //错误事件
</code></pre>

<p>好了，了解了 poll 函数的参数后，我们来看下如何使用 poll 函数完成网络通信。这个流程主要可以分成三步：</p>

<ul>
<li>第一步，创建 pollfd 数组和监听套接字，并进行绑定；</li>
<li>第二步，将监听套接字加入 pollfd 数组，并设置其监听读事件，也就是客户端的连接请求；</li>
<li>第三步，循环调用 poll 函数，检测 pollfd 数组中是否有就绪的文件描述符。</li>
</ul>

<p>而在第三步的循环过程中，其处理逻辑又分成了两种情况：</p>

<ul>
<li>如果是连接套接字就绪，这表明是有客户端连接，我们可以调用 accept 接受连接，并创建已连接套接字，并将其加入 pollfd 数组，并监听读事件；</li>
<li>如果是已连接套接字就绪，这表明客户端有读写请求，我们可以调用 recv/send 函数处理读写请求。</li>
</ul>

<p>我画了下面这张图，展示了使用 poll 函数的流程，你可以学习掌握下。</p>

<p><img src="assets/b1dab536cc9509f476db2c527fdea619-20221013235113-jm8fzak.jpg" alt="" /></p>

<p>另外，为了便于你掌握在代码中使用 poll 函数，我也写了一份示例代码，如下所示：</p>

<pre><code class="language-c">int sock_fd,conn_fd; //监听套接字和已连接套接字的变量
sock_fd = socket() //创建套接字
bind(sock_fd)   //绑定套接字
listen(sock_fd) //在套接字上进行监听，将套接字转为监听套接字

//poll函数可以监听的文件描述符数量，可以大于1024
#define MAX_OPEN = 2048

//pollfd结构体数组，对应文件描述符
struct pollfd client[MAX_OPEN];

//将创建的监听套接字加入pollfd数组，并监听其可读事件
client[0].fd = sock_fd;
client[0].events = POLLRDNORM; 
maxfd = 0;

//初始化client数组其他元素为-1
for (i = 1; i &lt; MAX_OPEN; i++)
    client[i].fd = -1; 

while(1) {
   //调用poll函数，检测client数组里的文件描述符是否有就绪的，返回就绪的文件描述符个数
   n = poll(client, maxfd+1, &amp;timeout);
   //如果监听套件字的文件描述符有可读事件，则进行处理
   if (client[0].revents &amp; POLLRDNORM) {
       //有客户端连接；调用accept函数建立连接
       conn_fd = accept();

       //保存已建立连接套接字
       for (i = 1; i &lt; MAX_OPEN; i++){
         if (client[i].fd &lt; 0) {
           client[i].fd = conn_fd; //将已建立连接的文件描述符保存到client数组
           client[i].events = POLLRDNORM; //设置该文件描述符监听可读事件
           break;
          }
       }
       maxfd = i; 
   }
   
   //依次检查已连接套接字的文件描述符
   for (i = 1; i &lt; MAX_OPEN; i++) {
       if (client[i].revents &amp; (POLLRDNORM | POLLERR)) {
         //有数据可读或发生错误，进行读数据处理或错误处理
       }
   }
}
</code></pre>

<p>其实，和 select 函数相比，poll 函数的改进之处主要就在于，它允许一次监听超过 1024 个文件描述符。但是当调用了 poll 函数后，我们仍然需要遍历每个文件描述符，检测该描述符是否就绪，然后再进行处理。</p>

<p><strong>那么，有没有办法可以避免遍历每个描述符呢？</strong>这就是我接下来向你介绍的 epoll 机制。</p>

<h2 id="使用-epoll-机制实现-io-多路复用">使用 epoll 机制实现 IO 多路复用</h2>

<p>首先，epoll 机制是使用 epoll_event 结构体，来记录待监听的文件描述符及其监听的事件类型的，这和 poll 机制中使用 pollfd 结构体比较类似。</p>

<p>那么，对于 epoll_event 结构体来说，其中包含了 epoll_data_t 联合体变量，以及整数类型的 events 变量。epoll_data_t 联合体中有记录文件描述符的成员变量 fd，而 events 变量会取值使用不同的宏定义值，来表示 epoll_data_t 变量中的文件描述符所关注的事件类型，比如一些常见的事件类型包括以下这几种。</p>

<ul>
<li>EPOLLIN：读事件，表示文件描述符对应套接字有数据可读。</li>
<li>EPOLLOUT：写事件，表示文件描述符对应套接字有数据要写。</li>
<li>EPOLLERR：错误事件，表示文件描述符对于套接字出错。</li>
</ul>

<p>下面的代码展示了 epoll_event 结构体以及 epoll_data 联合体的定义，你可以看下。</p>

<pre><code class="language-c">typedef union epoll_data
{
  ...
  int fd;  //记录文件描述符
  ...
} epoll_data_t;


struct epoll_event
{
  uint32_t events;  //epoll监听的事件类型
  epoll_data_t data; //应用程序数据
};
</code></pre>

<p>好了，现在我们知道，在使用 select 或 poll 函数的时候，创建好文件描述符集合或 pollfd 数组后，就可以往数组中添加我们需要监听的文件描述符。</p>

<p>但是对于 epoll 机制来说，我们则需要先调用 epoll_create 函数，创建一个 epoll 实例。这个 epoll 实例内部维护了两个结构，分别是<strong>记录要监听的文件描述符</strong>和<strong>已经就绪的文件描述符</strong>，而对于已经就绪的文件描述符来说，它们会被返回给用户程序进行处理。</p>

<p>所以，我们在使用 epoll 机制时，就不用像使用 select 和 poll 一样，遍历查询哪些文件描述符已经就绪了。这样一来， epoll 的效率就比 select 和 poll 有了更高的提升。</p>

<p>在创建了 epoll 实例后，我们需要再使用 epoll_ctl 函数，给被监听的文件描述符添加监听事件类型，以及使用 epoll_wait 函数获取就绪的文件描述符。</p>

<p>我画了一张图，展示了使用 epoll 进行网络通信的流程，你可以看下。</p>

<p><img src="assets/1ee730305558d9d83ff8e52eb4d966fb-20221013235113-11n924w.jpg" alt="" /></p>

<p>下面的代码展示了使用 epoll 函数的流程，你也可以看下。</p>

<pre><code class="language-c">int sock_fd,conn_fd; //监听套接字和已连接套接字的变量
sock_fd = socket() //创建套接字
bind(sock_fd)   //绑定套接字
listen(sock_fd) //在套接字上进行监听，将套接字转为监听套接字
  
epfd = epoll_create(EPOLL_SIZE); //创建epoll实例，
//创建epoll_event结构体数组，保存套接字对应文件描述符和监听事件类型  
ep_events = (epoll_event*)malloc(sizeof(epoll_event) * EPOLL_SIZE);

//创建epoll_event变量
struct epoll_event ee
//监听读事件
ee.events = EPOLLIN;
//监听的文件描述符是刚创建的监听套接字
ee.data.fd = sock_fd;

//将监听套接字加入到监听列表中  
epoll_ctl(epfd, EPOLL_CTL_ADD, sock_fd, &amp;ee); 
  
while (1) {
   //等待返回已经就绪的描述符 
   n = epoll_wait(epfd, ep_events, EPOLL_SIZE, -1); 
   //遍历所有就绪的描述符   
   for (int i = 0; i &lt; n; i++) {
       //如果是监听套接字描述符就绪，表明有一个新客户端连接到来 
       if (ep_events[i].data.fd == sock_fd) { 
          conn_fd = accept(sock_fd); //调用accept()建立连接
          ee.events = EPOLLIN;  
          ee.data.fd = conn_fd;
          //添加对新创建的已连接套接字描述符的监听，监听后续在已连接套接字上的读事件    
          epoll_ctl(epfd, EPOLL_CTL_ADD, conn_fd, &amp;ee); 
            
       } else { //如果是已连接套接字描述符就绪，则可以读数据
           ...//读取数据并处理
       }
   }
}
</code></pre>

<p>好了，到这里，你就了解了 epoll 函数的使用方法了。实际上，也正是因为 epoll 能自定义监听的描述符数量，以及可以直接返回就绪的描述符，Redis 在设计和实现网络通信框架时，就基于 epoll 机制中的 epoll_create、epoll_ctl 和 epoll_wait 等函数和读写事件，进行了封装开发，实现了用于网络通信的事件驱动框架，从而使得 Redis 虽然是单线程运行，但是仍然能高效应对高并发的客户端访问。</p>

<h2 id="小结">小结</h2>

<p>今天这节课，我给你介绍了 Redis 网络通信依赖的操作系统底层机制，也就是 IO 多路复用机制。</p>

<p>由于 Redis 是单线程程序，如果使用基本的 Socket 编程模型的话，只能对一个监听套接字或一个已连接套接字进行监听。而当 Redis 实例面临很多并发的客户端时，这种处理方式的效率就会很低。</p>

<p>所以，和基本的 Socket 通信相比，使用 IO 多路复用机制，就可以一次性获得就绪的多个套接字，从而避免了逐个检测套接字的开销。</p>

<p>这节课，我是以最常用的 Linux 操作系统为例，给你具体介绍了 Linux 系统提供的三种 IO 多路复用机制，分别是 select、poll 和 epoll。这三种机制在能监听的描述符数量和查找就绪描述符的方法上是不一样的，你可以重点参考下图，来掌握它们的不同之处。这些差异，其实也决定了 epoll 相比于 select 和 poll 来说，效率更高，也应用更广泛。</p>

<p><img src="assets/c04feac38f984a0c407985ec793ca95c-20221013235113-uoid3yq.jpg" alt="" /></p>

<p>最后我想说的是，虽然这节课我没有给你介绍 Redis 的源码，但是学习 IO 多路复用的机制和使用流程，其实就是掌握 Redis 事件驱动框架的基础。Redis 的<a href="https://github.com/redis/redis/tree/5.0/src/ae_select.c" target="_blank">ae_select.c</a>和<a href="https://github.com/redis/redis/tree/5.0/src/ae_epoll.c" target="_blank">ae_epoll.c</a>文件，就分别使用了 select 和 epoll 这两种机制，实现 IO 多路复用。而在接下来的第 10、11 两节课上，我还会给分别你介绍，Redis 事件驱动框架是如何基于 epoll 进行封装开发和运行的，以及 Redis 事件驱动框架的事件类型和处理方法。这样一来，你就能对 Redis 事件驱动框架的底层支撑、框架运行和事件类型与处理，有个全面的掌握了。</p>

<h2 id="每课一问">每课一问</h2>

<p>在 Redis 事件驱动框架代码中，分别使用了 Linux 系统上的 select 和 epoll 两种机制，你知道为什么 Redis 没有使用 poll 这一机制吗？</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#117d7d7d28252020212651767c70787d3f727e7c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1239e1fc24ede4',t:'MTczNDA1MzgxNi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
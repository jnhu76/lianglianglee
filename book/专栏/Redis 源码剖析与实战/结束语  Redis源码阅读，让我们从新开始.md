<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=结束语&#32;&#32;Redis源码阅读，让我们从新开始>
        <link rel="icon" href="/static/favicon.png">
        <title>结束语  Redis源码阅读，让我们从新开始 </title>
        
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
                            <h1 id="title" data-id="结束语  Redis源码阅读，让我们从新开始" class="title">结束语  Redis源码阅读，让我们从新开始</h1>
                            <div><p>不知不觉中，我和你又一起走过了 3 个多月的时光。在这 3 个多月的时间里，我和你一起并肩作战，去学习和了解了 Redis 的源码。跟第一季的课程内容相比，这一季学习的内容的确更有难度，也更加需要你能静下心来钻研。</p>

<p>这里先感谢你的一路陪伴，我们一起走到了现在。做这样一门专栏，我自己也是收获了很多、成长了很多。那么最后一节课，我就把我在做这门课程里三点最重要的认知分享给你，我们一起持续精进。</p>

<h2 id="用源码重新认知你的知识体系">用源码重新认知你的知识体系</h2>

<p>其实，要是用一句话来总结我的感受，那就是，阅读源码让我感到“<strong>从新开始</strong>”。</p>

<p>我在学习 Redis 源码之前，已经对 Redis 的一些基本原理、一些常见的后端系统设计都有了了解和掌握，本身也有一些 C 语言的开发经历。我相信，我当时的状态和此刻正在阅读这篇结束语的你可能很相似。</p>

<p>而在学习了 Redis 源码后，我发现自己在 C 语言编程技巧、计算机系统关键机制，还有系统设计原则等等很多方面，都有了新的认识。这些新认知，是源于对 Redis 源码设计与实现的学习，而源码学习本身又给我提供了高于 Redis 的通用知识的掌握，这让我受益匪浅。</p>

<p>就比如说，我以前在学习操作系统时，了解了进程间通信的方法有消息队列、命名管道、无名管道、共享内存等等，但是一直没能建立直观的认知。而在阅读 Redis 源码时，我发现 Redis 广泛地使用了无名管道，来支持父子进程间的通信。这一下子就在我的知识体系中，增加了对管道实际开发使用的新认知，这也让我有了一种实践正好结合理论的体会。</p>

<p>而另一方面，我以前在实现一些数据结构和算法时，都会按照它们在书本上的定义去实现。但是，在阅读 Redis 源码过程中，我发现其实实践和理论又是有差异的。就像 Redis 中的字符串根据不同长度，使用了不同的数据结构实现；有序集合使用了两种数据结构的组合来实现；以及 LRU 算法采用了近似方法来实现等等。</p>

<p>这些实际代码让我的知识体系，对实践结合理论又有了新的认识：其实在实际系统开发中，我们通常要考虑性能、空间、复杂度等约束条件，会在理论基础上进行优化开发。这一新认知对我后来的开发工作有了很大帮助，我会有意识地识别所开发系统面临的约束，进而优化自己的实现方法。</p>

<p>其实从 Redis 的源码中，我们可以掌握很多计算机系统知识，这些新知识，或许我们在目前的工作里还用不到，甚至在日后不断学习的过程中，还会被更新迭代掉。但是我们要清楚一点，就是<strong>我们在某一阶段所掌握的知识，往往会是下一阶段知识的基础。</strong></p>

<p>源码阅读本身就是一个结合之前学习的理论和开发知识，进一步学习实践开发知识的过程，这是一种从知识再到知识的过程，也是让我们重新认知自己知识体系的过程。</p>

<h2 id="用源码重新磨炼你的意志力">用源码重新磨炼你的意志力</h2>

<p>阅读源码是一件很辛苦的事情，尤其是当我们面对一个庞大的代码结构时，往往就会感到无从下手了。而等到我们好不容易摸清了代码结构，知道了要从哪些关键函数开始看起时，我们又会面临代码中复杂的调用关系、高级的语法实现，同时，还要尝试去理解代码开发者的思路。这些都是我们在阅读代码过程中的拦路虎，很容易就让我们打退堂鼓了。</p>

<p>我自己在阅读源码时，这些问题也都碰到了。不过，<strong>我这个过程看成是对自己意志力的一个磨炼，越是遇到困难，越要迎难而上，而不能轻言放弃。</strong>***<strong><em>把</em></strong>***</p>

<p>虽然我们也能通过坚持做某些事来磨炼自己的意志力，但是阅读代码的挑战性更大。这是因为代码是细节，而掌握细节需要我们有足够的静心、耐心和细心。这和学习原理不一样，学习原理的时候，我们的头脑往往转得很快，有些机制我们会想当然地认同了。</p>

<p>而阅读代码就不能这样了，一段代码不理解就是不理解，我们是无法想当然认同的。我们只有在不断尝试理解代码的过程中，<strong>正视自己想要放弃的心理和消极情绪，并能找到原因记录下来，然后逐渐减少阻力，以及慢慢提高自己想要放弃它的心理阈值</strong>。这正是阅读源码给我的意志力带来的新磨炼。</p>

<p>当然除了有意志力的支持，我们也需要有合理的方法。我之前看过一本书叫做《干劲的开关》，其中有句话是这样说的：“影响结果的不是斗志，而是科学”。所以我在读源码的时候，我就把阅读代码的目标拆分得更加细粒度化，每天、每周完成一些小目标，日积月累，等到我把 Redis 源码主要部分阅读完后，我收获了很大的成就感，因为我做到了。</p>

<p>而且在那之后，我也发现自己再做其他一些具有挑战性的工作时，阅读源码时得到磨炼的意志力就会发挥积极作用，让我自己不再畏惧困难，而是会积极应对。那么相对应地，我希望你在阅读源码的时候，也能够不要被代码的复杂结构或是错综调用关系所吓倒，而是规划好切实可达的目标，一步一个脚印地去完成代码的学习。</p>

<h2 id="用源码重新塑造你的做事原则">用源码重新塑造你的做事原则</h2>

<p>我之前在做事时，通常都是直线思维，定了一个目标就希望一次性完成这个目标。但有时受限于自己的知识背景和能力，对如何一次性完成目标会感到很困惑。</p>

<p>而在阅读 Redis 源码时，我遇到了相同的困惑：我一直奔着一定要把主要代码和关键技术掌握好这个目标而学习。但是在源码阅读的过程中，我有时在阅读了部分代码后，又会忘了之前学习的一些细节。而且对于在学习时已经厘清的概念和方法，等过一段时间之后，我发现又会变得模糊了。</p>

<p>后来，我自己在开发一个系统时，经常会去再回顾 Redis 源码。等这个系统开发完成后，我发现，原先变得模糊的 Redis 代码细节，已经变成深刻的记忆沉淀下来了。</p>

<p>在那个时候，我想明白了，<strong>源码阅读从来都不是一个一次性的学习过程。相反，源码阅读过程就像是 DNA 的双螺旋结构一样，是一个循环向上的过程。</strong>从源码阅读中学习开发知识，了解系统实现，然后再用学到的知识反哺自己的系统开发。而在开发过程中，又会再次阅读源码，进行学习，将自己的认知重新提升一个层次。这个过程周而复始，循环向上。</p>

<p>其实，我们日常的学习和做事跟源码阅读也是很相似的，它是一个循环向上的过程。很多事情并不是一蹴而就的，我们需要经历“认知、实践、再认知、再实践”这样一个过程。在这个过程中，我们会遇到困难，也会有收获，但是这些困难或收获都是为了下一次的认知和实践打基础。所以，<strong>我们不要因为一时的挫折而气馁，也不要因为一时的成就而停滞，就像生命之源的 DNA 结构一样，我们螺旋上升。</strong></p>

<h2 id="写在最后"><strong>写在最后</strong></h2>

<p>今天正好是周六，是个承前启后的时刻，这一季课程也要在今天正式画上一个句号了。不过，这个句号既是一个阶段的结束，更是一个新阶段的开始。</p>

<p>其实很多时候，你的成功并不取决于你知道了多少，而是你知道在无知的时候该怎么做。我希望上面讲的三点关于阅读源码的认知，能够给你的工作和生活提供一些指导方向。当我们在学习一个不会的知识点时，当我们在学习一门新的语言时，当我们面对生活中各种各样的情况必须孤立地做出反应时，我希望你能联想到学习源码的底层逻辑，从而更好地作出自己的选择。</p>

<p>最后，到了我们说再见的时候了，再次感谢你的一路相伴，我相信现在我们都成长了很多。而从今天起，我们在学习的道路上又将是一个新征程，让我们从新开始，学以致用。</p>

<p>最后的最后，我还给你准备了一份毕业问卷，希望你能花两三分钟填写一下，我非常期待能听到你对这门课的反馈。</p>

<p>好了，天长地久有时尽，学习之路绵绵无绝期，我们下一次再会！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#177b7b7b2e232626272057707a767e7b3974787a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f123f698b84ede4',t:'MTczNDA1NDA0My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
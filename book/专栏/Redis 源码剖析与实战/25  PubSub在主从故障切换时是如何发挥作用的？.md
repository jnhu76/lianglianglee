<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=25&#32;&#32;PubSub在主从故障切换时是如何发挥作用的？>
        <link rel="icon" href="/static/favicon.png">
        <title>25  PubSub在主从故障切换时是如何发挥作用的？ </title>
        
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
                            <h1 id="title" data-id="25  PubSub在主从故障切换时是如何发挥作用的？" class="title">25  PubSub在主从故障切换时是如何发挥作用的？</h1>
                            <div><p>在前面两节课，我们学习了哨兵工作的基本过程：哨兵会使用 sentinelRedisInstance 结构体来记录主节点的信息，在这个结构体中又记录了监听同一主节点的其他哨兵的信息。<strong>那么，一个哨兵是如何获得其他哨兵的信息的呢？</strong></p>

<p>这其实就和哨兵在运行过程中，使用的<strong>发布订阅（Pub/Sub）</strong>通信方法有关了。Pub/Sub 通信方法可以让哨兵订阅一个或多个频道，当频道中有消息时，哨兵可以收到相应消息；同时，哨兵也可以向频道中发布自己生成的消息，以便订阅该频道的其他客户端能收到消息。</p>

<p>今天这节课，我就带你来了解发布订阅通信方法的实现，以及它在哨兵工作过程中的应用。同时，你还可以了解哨兵之间是如何发现彼此的，以及客户端是如何知道故障切换完成的。Pub/Sub 通信方法在分布式系统中可以用作多对多的信息交互，在学完这节课之后，当你要实现分布式节点间通信时，就可以把它应用起来。</p>

<p>好了，接下来，我们先来看下发布订阅通信方法的实现。</p>

<h2 id="发布订阅通信方法的实现">发布订阅通信方法的实现</h2>

<p>发布订阅通信方法的基本模型是包含<strong>发布者、频道和订阅者</strong>，发布者把消息发布到频道上，而订阅者会订阅频道，一旦频道上有消息，频道就会把消息发送给订阅者。一个频道可以有多个订阅者，而对于一个订阅者来说，它也可以订阅多个频道，从而获得多个发布者发布的消息。</p>

<p>下图展示的就是发布者 - 频道 - 订阅者的基本模型，你可以看下。</p>

<p><img src="assets/b41f7bc078c210ef35d17bf3dc9be09d-20221014000254-298sfqw.jpg" alt="" /></p>

<h3 id="频道的实现">频道的实现</h3>

<p>了解了发布订阅方法的基本模型后，我们就来看下频道是如何实现的，因为在发布订阅通信方法中，频道很重要，它是发布者和订阅者之间通信的基础。</p>

<p>其实，Redis 的全局变量 server 使用了一个成员变量 <strong>pubsub_channels</strong> 来保存频道，pubsub_channels 的初始化是在 initServer 函数（在<a href="https://github.com/redis/redis/tree/5.0/src/server.c" target="_blank">server.c</a>文件中）中完成的。initServer 函数会调用 dictCreate 创建一个 <strong>keylistDictType 类型的哈希表</strong>，然后用这个哈希表来保存频道的信息，如下所示：</p>

<pre><code class="language-c">void initServer(void) {
…
server.pubsub_channels = dictCreate(&amp;keylistDictType,NULL);
…
}
</code></pre>

<p>注意，当哈希表是 keylistDictType 类型时，它保存的哈希项的 value 就是一个列表。而之所以采用这种类型来保存频道信息，是因为 Redis 把频道的名称作为哈希项的 key，而把订阅频道的订阅者作为哈希项的 value。就像刚才我们介绍的，一个频道可以有多个订阅者，所以 Redis 在实现时，就会用列表把订阅同一个频道的订阅者保存起来。</p>

<p>pubsub_channels 哈希表保存频道和订阅者的示意图如下所示：</p>

<p><img src="assets/21775d173b9db1650d3285a18d7d7a7f-20221014000254-xs8fmxc.jpg" alt="" /></p>

<p>了解了频道是如何实现的之后，下面我们再分别看下发布命令和订阅命令的实现。</p>

<h3 id="发布命令的实现">发布命令的实现</h3>

<p>发布命令在 Redis 的实现中对应的是 <strong>publish</strong>。我在【第 14 讲】中给你介绍过，Redis server 在初始化时，会初始化一个命令表 redisCommandTable，表中就记录了 Redis 支持的各种命令，以及对应的实现函数。</p>

<p>这张命令表是在 server.c 文件中定义的，当你需要了解 Redis 某个命令的具体实现函数时，一个快捷的方式就是在这张表中查找对应命令，然后就能定位到该命令的实现函数了。我们同样可以用这个方法来定位 publish 命令，这样就可以看到它<strong>对应的实现函数是 publishCommand</strong>（在<a href="https://github.com/redis/redis/tree/5.0/src/pubsub.c" target="_blank">pubsub.c</a>文件中），如下所示：</p>

<pre><code class="language-c">struct redisCommand redisCommandTable[] = {
…
{&quot;publish&quot;,publishCommand,3,&quot;pltF&quot;,0,NULL,0,0,0,0,0},
…
}
</code></pre>

<p>我们来看下 publishCommand 函数，它是调用 <strong>pubsubPublishMessage 函数</strong>（在 pubsub.c 文件中）来完成消息的实际发送，然后，再返回接收消息的订阅者数量的，如下所示：</p>

<pre><code class="language-c">void publishCommand(client *c) {
    //调用pubsubPublishMessage发布消息
    int receivers = pubsubPublishMessage(c-&gt;argv[1],c-&gt;argv[2]);
    … //如果Redis启用了cluster，那么在集群中发送publish命令
    addReplyLongLong(c,receivers); //返回接收消息的订阅者数量
}
</code></pre>

<p>而对于 pubsubPublishMessage 函数来说，它的原型如下。你可以看到，它的两个参数分别是要<strong>发布消息的频道</strong>，以及<strong>要发布的具体消息</strong>。</p>

<pre><code class="language-c">int pubsubPublishMessage(robj *channel, robj *message)
</code></pre>

<p>pubsubPublishMessage 函数会在 server.pubsub_channels 哈希表中，查找要发布的频道。如果找见了，它就会遍历这个 channel 对应的订阅者列表，然后依次向每个订阅者发送要发布的消息。这样一来，只要订阅者订阅了这个频道，那么发布者发布消息时，它就能收到了。</p>

<pre><code class="language-c">//查找频道是否存在
de = dictFind(server.pubsub_channels,channel);
    if (de) { //频道存在
        …
        //遍历频道对应的订阅者，向订阅者发送要发布的消息
        while ((ln = listNext(&amp;li)) != NULL) {
            client *c = ln-&gt;value;
            …
            addReplyBulk(c,channel);
            addReplyBulk(c,message);
            receivers++;
        }
    }
</code></pre>

<p>好了，了解了发布命令后，我们再来看下订阅命令的实现。</p>

<h3 id="订阅命令的实现">订阅命令的实现</h3>

<p>和查找发布命令的方法一样，我们可以在 redisCommandTable 表中，找到订阅命令 <strong>subscribe</strong> 对应的实现函数是 <strong>subscribeCommand</strong>（在 pubsub.c 文件中）。</p>

<p>subscribeCommand 函数的逻辑比较简单，它会直接调用 pubsubSubscribeChannel 函数（在 pubsub.c 文件中）来完成订阅操作，如下所示：</p>

<pre><code class="language-c">void subscribeCommand(client *c) {
    int j;
    for (j = 1; j &lt; c-&gt;argc; j++)
        pubsubSubscribeChannel(c,c-&gt;argv[j]);
    c-&gt;flags |= CLIENT_PUBSUB;
}
</code></pre>

<p>从代码中，你可以看到，subscribeCommand 函数的参数是 client 类型的变量，而它会根据 client 的 <strong>argc</strong> 成员变量执行一个循环，并把 client 的每个 <strong>argv</strong> 成员变量传给 pubsubSubscribeChannel 函数执行。</p>

<p>对于 client 的 argc 和 argv 来说，它们分别代表了要执行命令的参数个数和具体参数值，那么，<strong>这里的参数值是指什么呢?</strong></p>

<p>其实，我们来看下 pubsubSubscribeChannel 函数的原型就能知道了，如下所示：</p>

<pre><code class="language-c">int pubsubSubscribeChannel(client *c, robj *channel)
</code></pre>

<p>pubsubSubscribeChannel 函数的参数除了 client 变量外，还会<strong>接收频道的信息</strong>，这也就是说，subscribeCommand 会按照 subscribe 执行时附带的频道名称，来逐个订阅频道。我也在下面展示了 subscribe 命令执行的一个示例，你可以看下。当这个 subscribe 命令执行时，它会订阅三个频道，分别是 channel1、channel2 和 channel3：</p>

<pre><code class="language-c">subscribe channel1 channel2 channel3
</code></pre>

<p>下面我们来具体看下 pubsubSubscribeChannel 函数的实现。这个函数的逻辑也比较清晰，主要可以分成三步。</p>

<p><strong>首先</strong>，它把要订阅的频道加入到 server 记录的 pubsub_channels 中。如果这个频道是新创建的，那么它会在 pubsub_channels 哈希表中新建一个哈希项，代表新创建的这个频道，并且会创建一个列表，用来保存这个频道对应的订阅者。</p>

<p>如果频道已经在 pubsub_channels 哈希表中存在了，那么 pubsubSubscribeChannel 函数就直接获取该频道对应的订阅者列表。</p>

<p><strong>然后</strong>，pubsubSubscribeChannel 函数把执行 subscribe 命令的订阅者，加入到订阅者列表中。</p>

<p><strong>最后</strong>，pubsubSubscribeChannel 函数会把成功订阅的频道个数返回给订阅者。</p>

<p>下面的代码展示了这部分的逻辑，你可以看下。</p>

<pre><code class="language-c">if (dictAdd(c-&gt;pubsub_channels,channel,NULL) == DICT_OK) {
   …
   de = dictFind(server.pubsub_channels,channel); //在pubsub_channels哈希表中查找频道
   if (de == NULL) { //如果频道不存在
      clients = listCreate();  //创建订阅者对应的列表
      dictAdd(server.pubsub_channels,channel,clients); //新插入频道对应的哈希项
      …
    } else {
      clients = dictGetVal(de); //频道已存在，获取订阅者列表
    }
    listAddNodeTail(clients,c); //将订阅者加入到订阅者列表
}
 
…
addReplyLongLong(c,clientSubscriptionsCount(c)); //给订阅者返回成功订阅的频道数量
</code></pre>

<p>现在，你就了解了 Redis 中发布订阅方法的实现。接下来，我们来看下哨兵在工作过程中，又是如何使用发布订阅功能的。</p>

<h2 id="发布订阅方法在哨兵中的应用">发布订阅方法在哨兵中的应用</h2>

<p>首先，我们来看下哨兵用来发布消息的函数 sentinelEvent。</p>

<h3 id="sentinelevent-函数与消息生成">sentinelEvent 函数与消息生成</h3>

<p>哨兵在使用发布订阅方法时，封装了 <strong>sentinelEvent 函数</strong>（在<a href="https://github.com/redis/redis/tree/5.0/src/sentinel.c" target="_blank">sentinel.c</a>文件中），用来发布消息。所以，你在阅读 sentinel.c 文件中关于哨兵的源码时，如果看到 sentinelEvent，这就表明哨兵正在用它来发布消息。</p>

<p>我在【第 22 讲】中给你介绍过 sentinelEvent 函数，你可以再回顾下。这个函数的原型如下所示：</p>

<pre><code class="language-c">void sentinelEvent(int level, char *type, sentinelRedisInstance *ri, const char *fmt, ...) 
</code></pre>

<p>实际上，这个函数最终是通过调用刚才我提到的 pubsubPublishMessage 函数，来实现向某一个频道发布消息的。那么，当我们要发布一条消息时，需要确定两个方面的内容：<strong>一个是要发布的频道，另一个是要发布的消息</strong>。</p>

<p>sentinelEvent 函数的第二个参数 type，表示的就是要发布的频道，而要发布的消息，就是由这个函数第四个参数 fmt 后面的省略号来表示的。</p>

<p>看到这里，你可以会有一个疑问，<strong>为什么 sentinelEvent 函数参数中会有省略号？</strong></p>

<p>其实，这里的省略号表示的是<strong>可变参数</strong>，当我们无法列出传递给函数的所有实参类型和数目时，我们可以用省略号来表示可变参数，这就是说，我们可以给 sentinelEvent 函数传递 4 个、5 个、6 个甚至更多的参数。</p>

<p>我在这里就以 sentinelEvent 函数的实现为例，给你介绍下可变参数的使用，这样一来，当你在开发分布式通信程序时，需要生成内容不定的消息时，就可以把哨兵源码中实现的方法用起来。</p>

<p>在 sentinelEvent 函数中，为了使用了可变参数，它主要包含了四个步骤：</p>

<ul>
<li>首先，我们需要定义一个 va_list 类型的变量，假设是 ap。这个变量是指向可变参数的指针。</li>
<li>然后，当我们要在函数中使用可变参数了，就需要通过 <strong>va_start 宏</strong>来获取可变参数中的第一个参数。va_start 宏有两个参数，一个是刚才定义的 va_list 类型变量 ap，另一个是可变参数的前一个参数，也就是 sentinelEvent 函数参数中，省略号前的参数 fmt。</li>
<li>紧接着，我们可以使用 vsnprintf 函数，来按照 fmt 定义的格式，打印可变参数中的内容。vsnprintf 函数会逐个获取可变参数中的每一个参数，并进行打印。</li>
<li>最后，我们在获取完所有参数后，需要调用 va_end 宏将刚才创建的 ap 指针关闭。</li>
</ul>

<p>下面的代码展示了刚才介绍的这个过程，你可以再看下。</p>

<pre><code class="language-c">void sentinelEvent(int level, char *type, sentinelRedisInstance *ri, const char *fmt, ...) {
    va_list ap;
    ... 
    if (fmt[0] != '\0') {
        va_start(ap, fmt);
        vsnprintf(msg+strlen(msg), sizeof(msg)-strlen(msg), fmt, ap);
        va_end(ap);
    }
    ...
}
</code></pre>

<p>为了让你有个更加直观的了解，我在下面列了三个 sentinelEvent 函数的调用示例，你可以再学习掌握下。</p>

<p>第一个对应了哨兵调用 sentinelCheckSubjectivelyDown 函数<strong>检测出主节点主观下线后</strong>，sentinelCheckSubjectivelyDown 函数调用 sentinelEvent 函数，向“+sdown”频道发布消息。此时，传递给 sentinelEvent 的参数就是 4 个，并没有可变参数，如下所示：</p>

<pre><code class="language-c">sentinelEvent(LL_WARNING,&quot;+sdown&quot;,ri,&quot;%@&quot;);
</code></pre>

<p>第二个对应了<strong>哨兵在初始化时</strong>，在 sentinelGenerateInitialMonitorEvents 函数中，调用 sentinelEvent 函数向“+monitor”频道发布消息，此时，传递给 sentinelEvent 的参数有 5 个，包含了 1 个可变参数，表示的是哨兵的 quorum 阈值，如下所示：</p>

<pre><code class="language-c">sentinelEvent(LL_WARNING,&quot;+monitor&quot;,ri,&quot;%@ quorum %d&quot;,ri-&gt;quorum);
</code></pre>

<p>最后一个对应了<strong>哨兵在完成主节点切换后</strong>，在 sentinelFailoverSwitchToPromotedSlave 函数中，调用 sentinelEvent 函数向“+switch-master”频道发布消息。此时，传递给 sentinelEvent 的可变参数一共有 5 个，对应了故障切换前的主节点名称、IP 和端口号，以及切换后升级为主节点的从节点 IP 和端口号，如下所示：</p>

<pre><code class="language-c">sentinelEvent(LL_WARNING,&quot;+switch-master&quot;,master,&quot;%s %s %d %s %d&quot;,
        master-&gt;name, master-&gt;addr-&gt;ip, master-&gt;addr-&gt;port,
        ref-&gt;addr-&gt;ip, ref-&gt;addr-&gt;port);
</code></pre>

<p>这样一来，你也就了解了，哨兵在工作过程中是通过 sentinelEvent 函数和 pubsubPublishMessage 函数，来实现消息的发布的。在哨兵的整个工作过程中，它会在一些关键节点上，<strong>使用 sentinelEvent 函数往不同的频道上发布消息</strong>。除了刚才给你举例的三个频道 +monitor、+sdown、+switch-master 以外，我还把哨兵在工作过程中会用到的消息发布频道列在了下表中，你可以了解下。</p>

<p><img src="assets/53c920eec89a2108351d816a80cbe272-20221014000254-xggnu7x.jpg" alt="" /></p>

<p>其实，在哨兵的工作过程中，如果有客户端想要了解故障切换的整体情况或进度，比如主节点是否被判断为主观下线、主节点是否被判断为客观下线、Leader 是否完成选举、新主节点是否切换完成，等等，就可以通过 subscribe 命令，订阅上面这张表中的相应频道。这样一来，客户端就可以了解故障切换的过程了。</p>

<p>好，下面我们再来看下，哨兵在工作过程中对消息的订阅是如何实现的。</p>

<h3 id="哨兵订阅与-hello-频道">哨兵订阅与 hello 频道</h3>

<p>首先你要知道，每个哨兵会订阅它所监听的主节点的&rdquo;<strong>sentinel</strong>:hello&rdquo;频道。在【第 23 讲】中，我给你介绍过，哨兵会周期性调用 sentinelTimer 函数来完成周期性的任务，这其中，就有哨兵订阅主节点 hello 频道的操作。</p>

<p>具体来说，哨兵在周期性执行 sentinelTimer 函数时，会调用 sentinelHandleRedisInstance 函数，进而调用 sentinelReconnectInstance 函数。而在 sentinelReconnectInstance 函数中，哨兵会调用 redisAsyncCommand 函数，向主节点发送 subscribe 命令，订阅的频道由宏定义 SENTINEL_HELLO_CHANNEL（在 sentinel.c 文件中）指定，也就是&rdquo;<strong>sentinel</strong>:hello&rdquo;频道。这部分的代码如下所示：</p>

<pre><code class="language-c">retval = redisAsyncCommand(link-&gt;pc,
                sentinelReceiveHelloMessages, ri, &quot;%s %s&quot;,
                sentinelInstanceMapCommand(ri,&quot;SUBSCRIBE&quot;),
                SENTINEL_HELLO_CHANNEL);
</code></pre>

<p>从代码中，我们也可以看到，当在&rdquo;<strong>sentinel</strong>:hello&rdquo;频道上收到 hello 消息后，哨兵会回调 sentinelReceiveHelloMessages 函数来进行处理。而 sentinelReceiveHelloMessages 函数，实际是通过调用 <strong>sentinelProcessHelloMessage 函数</strong>，来完成 hello 消息的处理的。</p>

<p>对于 sentinelProcessHelloMessage 函数来说，它主要是从 hello 消息中获得发布 hello 消息的哨兵实例的基本信息，比如 IP、端口号、quorum 阈值等。如果当前哨兵并没有记录发布 hello 消息的哨兵实例的信息，那么，sentinelProcessHelloMessage 函数就会调用 <strong>createSentinelRedisInstance 函数</strong>，来创建发布 hello 消息的哨兵实例的信息记录，这样一来，当前哨兵就拥有了其他哨兵实例的信息了。</p>

<p>好了，了解了哨兵对&rdquo;<strong>sentinel</strong>:hello&rdquo;频道的订阅和处理后，我们还需要搞清楚一个问题，即<strong>哨兵是在什么时候发布 hello 消息的呢？</strong></p>

<p>这其实是哨兵在 sentinelTimer 函数中，调用 sentinelSendPeriodicCommands 函数时，由 sentinelSendPeriodicCommands 函数调用 sentinelSendHello 函数来完成的。</p>

<p><strong>sentinelSendHello 函数</strong>会调用 redisAsyncCommand 函数，向主节点的&rdquo;<strong>sentinel</strong>:hello&rdquo;频道发布 hello 消息。在它发送的 hello 消息中，包含了发布 hello 消息的哨兵实例的 IP、端口号、ID 和当前的纪元，以及该哨兵监听的主节点的名称、IP、端口号和纪元信息。</p>

<p>下面的代码就展示了 hello 消息的生成和发布，你可以看下。</p>

<pre><code class="language-c">//hello消息包含的内容
snprintf(payload,sizeof(payload),
        &quot;%s,%d,%s,%llu,&quot; //当前哨兵实例的信息，包括ip、端口号、ID和当前纪元
        &quot;%s,%s,%d,%llu&quot;, //当前主节点的信息，包括名称、IP、端口号和纪元
        announce_ip, announce_port, sentinel.myid,
        (unsigned long long) sentinel.current_epoch,
        master-&gt;name,master_addr-&gt;ip,master_addr-&gt;port,
        (unsigned long long) master-&gt;config_epoch);
//向主节点的hello频道发布hello消息
retval = redisAsyncCommand(ri-&gt;link-&gt;cc,
        sentinelPublishReplyCallback, ri, &quot;%s %s %s&quot;,
        sentinelInstanceMapCommand(ri,&quot;PUBLISH&quot;),
        SENTINEL_HELLO_CHANNEL,payload);
</code></pre>

<p>这样，当哨兵通过 sentinelSendHello，向自己监听的主节点的&rdquo;<strong>sentinel</strong>:hello&rdquo;频道发布 hello 消息时，和该哨兵监听同一个主节点的其他哨兵，也会订阅主节点的&rdquo;<strong>sentinel</strong>:hello&rdquo;频道，从而就可以获得该频道上的 hello 消息了。</p>

<p>通过这样的通信方式，监听同一主节点的哨兵就能相互知道彼此的访问信息了。如此一来，哨兵就可以基于这些访问信息，执行主节点状态共同判断，以及进行 Leader 选举等操作了。</p>

<h2 id="小结">小结</h2>

<p>今天这节课，我们了解了 Redis 实现的发布订阅通信方法。这个方法是提供了频道的方式，让要通信的双方按照频道来完成消息交互。而<strong>不同频道的不同名称，就代表了哨兵工作过程中的不同状态</strong>。当客户端需要了解哨兵的工作进度或是主节点的状态判断时，就可以通过订阅哨兵发布消息的频道来完成。</p>

<p>当然，对于一个哨兵来说，它一定会订阅的频道是它所监听的主节点的&rdquo;<strong>sentinel</strong>:hello&rdquo;频道。通过这个频道，监听同一主节点的不同哨兵就能通过频道上的 hello 消息，来交互彼此的访问信息了，比如哨兵的 IP、端口号等。</p>

<p>此外，在这节课，我还给你介绍了一个 <strong>C 语言函数可变参数的使用小技巧</strong>，当你开发发布订阅功能时，都需要生成发布的消息，而可变参数就可以用来生成长度不定的消息。希望你能把这个小技巧应用起来。</p>

<h2 id="每课一问">每课一问</h2>

<p>如果我们在哨兵实例上执行 publish 命令，那么，这条命令是不是就是由 pubsub.c 文件中的 publishCommand 函数来处理的呢?</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#90fcfcfca9a4a1a1a0a7d0f7fdf1f9fcbef3fffd" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f123da878eaede4',t:'MTczNDA1Mzk3MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
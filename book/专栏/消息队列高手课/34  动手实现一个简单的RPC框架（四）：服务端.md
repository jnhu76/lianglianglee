<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=34&#32;&#32;动手实现一个简单的RPC框架（四）：服务端>
        <link rel="icon" href="/static/favicon.png">
        <title>34  动手实现一个简单的RPC框架（四）：服务端 </title>
        
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
                        <a class="menu-item" id="00 开篇词  优秀的程序员，你的技术栈中不能只有“增删改查”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%bc%98%e7%a7%80%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%ef%bc%8c%e4%bd%a0%e7%9a%84%e6%8a%80%e6%9c%af%e6%a0%88%e4%b8%ad%e4%b8%8d%e8%83%bd%e5%8f%aa%e6%9c%89%e2%80%9c%e5%a2%9e%e5%88%a0%e6%94%b9%e6%9f%a5%e2%80%9d.md">00 开篇词  优秀的程序员，你的技术栈中不能只有“增删改查”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 预习  怎样更好地学习这门课？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/00%20%e9%a2%84%e4%b9%a0%20%20%e6%80%8e%e6%a0%b7%e6%9b%b4%e5%a5%bd%e5%9c%b0%e5%ad%a6%e4%b9%a0%e8%bf%99%e9%97%a8%e8%af%be%ef%bc%9f.md">00 预习  怎样更好地学习这门课？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  为什么需要消息队列？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/01%20%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%ef%bc%9f.md">01  为什么需要消息队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  该如何选择消息队列？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/02%20%20%e8%af%a5%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%ef%bc%9f.md">02  该如何选择消息队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  消息模型：主题和队列有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/03%20%20%e6%b6%88%e6%81%af%e6%a8%a1%e5%9e%8b%ef%bc%9a%e4%b8%bb%e9%a2%98%e5%92%8c%e9%98%9f%e5%88%97%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">03  消息模型：主题和队列有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  如何利用事务消息实现分布式事务？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/04%20%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e4%ba%8b%e5%8a%a1%e6%b6%88%e6%81%af%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9f.md">04  如何利用事务消息实现分布式事务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  如何确保消息不会丢失.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/05%20%20%e5%a6%82%e4%bd%95%e7%a1%ae%e4%bf%9d%e6%b6%88%e6%81%af%e4%b8%8d%e4%bc%9a%e4%b8%a2%e5%a4%b1.md">05  如何确保消息不会丢失.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  如何处理消费过程中的重复消息？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/06%20%20%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e6%b6%88%e8%b4%b9%e8%bf%87%e7%a8%8b%e4%b8%ad%e7%9a%84%e9%87%8d%e5%a4%8d%e6%b6%88%e6%81%af%ef%bc%9f.md">06  如何处理消费过程中的重复消息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  消息积压了该如何处理？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/07%20%20%e6%b6%88%e6%81%af%e7%a7%af%e5%8e%8b%e4%ba%86%e8%af%a5%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%ef%bc%9f.md">07  消息积压了该如何处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  答疑解惑（一）  网关如何接收服务端的秒杀结果？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/08%20%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%ef%bc%88%e4%b8%80%ef%bc%89%20%20%e7%bd%91%e5%85%b3%e5%a6%82%e4%bd%95%e6%8e%a5%e6%94%b6%e6%9c%8d%e5%8a%a1%e7%ab%af%e7%9a%84%e7%a7%92%e6%9d%80%e7%bb%93%e6%9e%9c%ef%bc%9f.md">08  答疑解惑（一）  网关如何接收服务端的秒杀结果？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  学习开源代码该如何入手？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/09%20%20%e5%ad%a6%e4%b9%a0%e5%bc%80%e6%ba%90%e4%bb%a3%e7%a0%81%e8%af%a5%e5%a6%82%e4%bd%95%e5%85%a5%e6%89%8b%ef%bc%9f.md">09  学习开源代码该如何入手？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  如何使用异步设计提升系统性能？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/10%20%20%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%bc%82%e6%ad%a5%e8%ae%be%e8%ae%a1%e6%8f%90%e5%8d%87%e7%b3%bb%e7%bb%9f%e6%80%a7%e8%83%bd%ef%bc%9f.md">10  如何使用异步设计提升系统性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  如何实现高性能的异步网络传输？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/11%20%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84%e5%bc%82%e6%ad%a5%e7%bd%91%e7%bb%9c%e4%bc%a0%e8%be%93%ef%bc%9f.md">11  如何实现高性能的异步网络传输？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  序列化与反序列化：如何通过网络传输结构化的数据？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/12%20%20%e5%ba%8f%e5%88%97%e5%8c%96%e4%b8%8e%e5%8f%8d%e5%ba%8f%e5%88%97%e5%8c%96%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e7%bd%91%e7%bb%9c%e4%bc%a0%e8%be%93%e7%bb%93%e6%9e%84%e5%8c%96%e7%9a%84%e6%95%b0%e6%8d%ae%ef%bc%9f.md">12  序列化与反序列化：如何通过网络传输结构化的数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  传输协议：应用程序之间对话的语言.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/13%20%20%e4%bc%a0%e8%be%93%e5%8d%8f%e8%ae%ae%ef%bc%9a%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e4%b9%8b%e9%97%b4%e5%af%b9%e8%af%9d%e7%9a%84%e8%af%ad%e8%a8%80.md">13  传输协议：应用程序之间对话的语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  内存管理：如何避免内存溢出和频繁的垃圾回收？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/14%20%20%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e5%86%85%e5%ad%98%e6%ba%a2%e5%87%ba%e5%92%8c%e9%a2%91%e7%b9%81%e7%9a%84%e5%9e%83%e5%9c%be%e5%9b%9e%e6%94%b6%ef%bc%9f.md">14  内存管理：如何避免内存溢出和频繁的垃圾回收？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  Kafka如何实现高性能IO？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/15%20%20Kafka%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%80%a7%e8%83%bdIO%ef%bc%9f.md">15  Kafka如何实现高性能IO？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  缓存策略：如何使用缓存来减少磁盘IO？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/16%20%20%e7%bc%93%e5%ad%98%e7%ad%96%e7%95%a5%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e7%bc%93%e5%ad%98%e6%9d%a5%e5%87%8f%e5%b0%91%e7%a3%81%e7%9b%98IO%ef%bc%9f.md">16  缓存策略：如何使用缓存来减少磁盘IO？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  如何正确使用锁保护共享数据，协调异步线程？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/17%20%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e4%bd%bf%e7%94%a8%e9%94%81%e4%bf%9d%e6%8a%a4%e5%85%b1%e4%ba%ab%e6%95%b0%e6%8d%ae%ef%bc%8c%e5%8d%8f%e8%b0%83%e5%bc%82%e6%ad%a5%e7%ba%bf%e7%a8%8b%ef%bc%9f.md">17  如何正确使用锁保护共享数据，协调异步线程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  如何用硬件同步原语（CAS）替代锁？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/18%20%20%e5%a6%82%e4%bd%95%e7%94%a8%e7%a1%ac%e4%bb%b6%e5%90%8c%e6%ad%a5%e5%8e%9f%e8%af%ad%ef%bc%88CAS%ef%bc%89%e6%9b%bf%e4%bb%a3%e9%94%81%ef%bc%9f.md">18  如何用硬件同步原语（CAS）替代锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  数据压缩：时间换空间的游戏.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/19%20%20%e6%95%b0%e6%8d%ae%e5%8e%8b%e7%bc%a9%ef%bc%9a%e6%97%b6%e9%97%b4%e6%8d%a2%e7%a9%ba%e9%97%b4%e7%9a%84%e6%b8%b8%e6%88%8f.md">19  数据压缩：时间换空间的游戏.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  RocketMQ Producer源码分析：消息生产的实现过程.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/20%20%20RocketMQ%20Producer%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%ef%bc%9a%e6%b6%88%e6%81%af%e7%94%9f%e4%ba%a7%e7%9a%84%e5%ae%9e%e7%8e%b0%e8%bf%87%e7%a8%8b.md">20  RocketMQ Producer源码分析：消息生产的实现过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  Kafka Consumer源码分析：消息消费的实现过程.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/21%20%20Kafka%20Consumer%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%ef%bc%9a%e6%b6%88%e6%81%af%e6%b6%88%e8%b4%b9%e7%9a%84%e5%ae%9e%e7%8e%b0%e8%bf%87%e7%a8%8b.md">21  Kafka Consumer源码分析：消息消费的实现过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  Kafka和RocketMQ的消息复制实现的差异点在哪？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/22%20%20Kafka%e5%92%8cRocketMQ%e7%9a%84%e6%b6%88%e6%81%af%e5%a4%8d%e5%88%b6%e5%ae%9e%e7%8e%b0%e7%9a%84%e5%b7%ae%e5%bc%82%e7%82%b9%e5%9c%a8%e5%93%aa%ef%bc%9f.md">22  Kafka和RocketMQ的消息复制实现的差异点在哪？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  RocketMQ客户端如何在集群中找到正确的节点？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/23%20%20RocketMQ%e5%ae%a2%e6%88%b7%e7%ab%af%e5%a6%82%e4%bd%95%e5%9c%a8%e9%9b%86%e7%be%a4%e4%b8%ad%e6%89%be%e5%88%b0%e6%ad%a3%e7%a1%ae%e7%9a%84%e8%8a%82%e7%82%b9%ef%bc%9f.md">23  RocketMQ客户端如何在集群中找到正确的节点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  Kafka的协调服务ZooKeeper：实现分布式系统的“瑞士军刀”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/24%20%20Kafka%e7%9a%84%e5%8d%8f%e8%b0%83%e6%9c%8d%e5%8a%a1ZooKeeper%ef%bc%9a%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e7%b3%bb%e7%bb%9f%e7%9a%84%e2%80%9c%e7%91%9e%e5%a3%ab%e5%86%9b%e5%88%80%e2%80%9d.md">24  Kafka的协调服务ZooKeeper：实现分布式系统的“瑞士军刀”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  RocketMQ与Kafka中如何实现事务？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/25%20%20RocketMQ%e4%b8%8eKafka%e4%b8%ad%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%ba%8b%e5%8a%a1%ef%bc%9f.md">25  RocketMQ与Kafka中如何实现事务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  MQTT协议：如何支持海量的在线IoT设备.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/26%20%20MQTT%e5%8d%8f%e8%ae%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e6%94%af%e6%8c%81%e6%b5%b7%e9%87%8f%e7%9a%84%e5%9c%a8%e7%ba%bfIoT%e8%ae%be%e5%a4%87.md">26  MQTT协议：如何支持海量的在线IoT设备.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  Pulsar的存储计算分离设计：全新的消息队列设计思路.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/27%20%20Pulsar%e7%9a%84%e5%ad%98%e5%82%a8%e8%ae%a1%e7%ae%97%e5%88%86%e7%a6%bb%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%85%a8%e6%96%b0%e7%9a%84%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e8%ae%be%e8%ae%a1%e6%80%9d%e8%b7%af.md">27  Pulsar的存储计算分离设计：全新的消息队列设计思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  答疑解惑（二）：我的100元哪儿去了？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/28%20%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%88%91%e7%9a%84100%e5%85%83%e5%93%aa%e5%84%bf%e5%8e%bb%e4%ba%86%ef%bc%9f.md">28  答疑解惑（二）：我的100元哪儿去了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  流计算与消息（一）：通过Flink理解流计算的原理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/29%20%20%e6%b5%81%e8%ae%a1%e7%ae%97%e4%b8%8e%e6%b6%88%e6%81%af%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e9%80%9a%e8%bf%87Flink%e7%90%86%e8%a7%a3%e6%b5%81%e8%ae%a1%e7%ae%97%e7%9a%84%e5%8e%9f%e7%90%86.md">29  流计算与消息（一）：通过Flink理解流计算的原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  流计算与消息（二）：在流计算中使用Kafka链接计算任务.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/30%20%20%e6%b5%81%e8%ae%a1%e7%ae%97%e4%b8%8e%e6%b6%88%e6%81%af%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%9c%a8%e6%b5%81%e8%ae%a1%e7%ae%97%e4%b8%ad%e4%bd%bf%e7%94%a8Kafka%e9%93%be%e6%8e%a5%e8%ae%a1%e7%ae%97%e4%bb%bb%e5%8a%a1.md">30  流计算与消息（二）：在流计算中使用Kafka链接计算任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  动手实现一个简单的RPC框架（一）：原理和程序的结构.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/31%20%20%e5%8a%a8%e6%89%8b%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84RPC%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%8e%9f%e7%90%86%e5%92%8c%e7%a8%8b%e5%ba%8f%e7%9a%84%e7%bb%93%e6%9e%84.md">31  动手实现一个简单的RPC框架（一）：原理和程序的结构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  动手实现一个简单的RPC框架（二）：通信与序列化.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/32%20%20%e5%8a%a8%e6%89%8b%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84RPC%e6%a1%86%e6%9e%b6%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e9%80%9a%e4%bf%a1%e4%b8%8e%e5%ba%8f%e5%88%97%e5%8c%96.md">32  动手实现一个简单的RPC框架（二）：通信与序列化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  动手实现一个简单的RPC框架（三）：客户端.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/33%20%20%e5%8a%a8%e6%89%8b%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84RPC%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%ae%a2%e6%88%b7%e7%ab%af.md">33  动手实现一个简单的RPC框架（三）：客户端.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  动手实现一个简单的RPC框架（四）：服务端.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/34%20%20%e5%8a%a8%e6%89%8b%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84RPC%e6%a1%86%e6%9e%b6%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e6%9c%8d%e5%8a%a1%e7%ab%af.md">34  动手实现一个简单的RPC框架（四）：服务端.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  答疑解惑（三）：主流消息队列都是如何存储消息的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/35%20%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e4%b8%bb%e6%b5%81%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%83%bd%e6%98%af%e5%a6%82%e4%bd%95%e5%ad%98%e5%82%a8%e6%b6%88%e6%81%af%e7%9a%84%ef%bc%9f.md">35  答疑解惑（三）：主流消息队列都是如何存储消息的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐  JMQ的Broker是如何异步处理消息的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/%e5%8a%a0%e9%a4%90%20%20JMQ%e7%9a%84Broker%e6%98%af%e5%a6%82%e4%bd%95%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86%e6%b6%88%e6%81%af%e7%9a%84%ef%bc%9f.md">加餐  JMQ的Broker是如何异步处理消息的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  程序员如何构建知识体系？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%ab%98%e6%89%8b%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e7%a8%8b%e5%ba%8f%e5%91%98%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%ef%bc%9f.md">结束语  程序员如何构建知识体系？.md</a>
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
                            <h1 id="title" data-id="34  动手实现一个简单的RPC框架（四）：服务端" class="title">34  动手实现一个简单的RPC框架（四）：服务端</h1>
                            <div><p>你好，我是李玥。</p>

<p>上节课我们一起学习了如何来构建这个 RPC 框架中最关键的部分，也就是：在客户端，如何根据用户注册的服务接口来动态生成桩的方法。在这里，除了和语言特性相关的一些动态编译小技巧之外，你更应该掌握的是其中动态代理这种设计思想，它的使用场景以及实现方法。</p>

<p>这节课我们一起来实现这个框架的最后一部分：服务端。对于我们这个 RPC 框架来说，服务端可以分为两个部分：注册中心和 RPC 服务。其中，注册中心的作用是帮助客户端来寻址，找到对应 RPC 服务的物理地址，RPC 服务用于接收客户端桩的请求，调用业务服务的方法，并返回结果。</p>

<h2 id="注册中心是如何实现的">注册中心是如何实现的？</h2>

<p>我们先来看看注册中心是如何实现的。一般来说，一个完整的注册中心也是分为客户端和服务端两部分的，客户端给调用方提供 API，并实现与服务端的通信；服务端提供真正的业务功能，记录每个 RPC 服务发来的注册信息，并保存到它的元数据中。当有客户端来查询服务地址的时候，它会从元数据中获取服务地址，返回给客户端。</p>

<p>由于注册中心并不是这个 RPC 框架的重点内容，所以在这里，我们只实现了一个单机版的注册中心，它只有客户端没有服务端，所有的客户端依靠读写同一个元数据文件来实现元数据共享。所以，我们这个注册中心只能支持单机运行，并不支持跨服务器调用。</p>

<p>但是，我们在这里，同样采用的是“面向接口编程”的设计模式，这样，你可以在不改动一行代码的情况下，就可以通过增加一个 SPI 插件的方式，提供一个可以跨服务器调用的真正的注册中心实现，比如说，一个基于 HTTP 协议实现的注册中心。我们再来复习一下，这种面向接口编程的设计是如何在注册中心中来应用的。</p>

<p>首先，我们在 RPC 服务的接入点，接口 RpcAccessPoint 中增加一个获取注册中心实例的方法：</p>

<pre><code>public interface RpcAccessPoint extends Closeable{
    /**
     * 获取注册中心的引用
     * @param nameServiceUri 注册中心 URI
     * @return 注册中心引用
     */
    NameService getNameService(URI nameServiceUri);
 
    // ...
}
</code></pre>

<p>这个方法的参数就是注册中心的 URI，也就是它的地址，返回值就是访问这个注册中心的实例。然后我们再给 NameService 接口增加两个方法：</p>

<pre><code>public interface NameService {
 
    /**
     * 所有支持的协议
     */
    Collection&lt;String&gt; supportedSchemes();
 
    /**
     * 连接注册中心
     * @param nameServiceUri 注册中心地址
     */
    void connect(URI nameServiceUri);
 
    // ...
}
</code></pre>

<p>其中 supportedSchemes 方法，返回可以支持的所有协议，比如我们在这个例子中的实现，它的协议是“file”。connect 方法就是给定注册中心服务端的 URI，去建立与注册中心服务端的连接。</p>

<p>下面我们来看获取注册中心的方法 getNameService 的实现，它的实现也很简单，就是通过 SPI 机制加载所有的 NameService 的实现类，然后根据给定的 URI 中的协议，去匹配支持这个协议的实现类，然后返回这个实现的引用就可以了。由于这部分实现是通用并且不会改变的，我们直接把实现代码放在 RpcAccessPoint 这个接口中。</p>

<p>这样我们就实现了一个可扩展的注册中心接口，系统可以根据 URI 中的协议，动态地来选择不同的注册中心实现。增加一种注册中心的实现，也不需要修改任何代码，只要按照 SPI 的规范，把协议的实现加入到运行时 CLASSPATH 中就可以了。（这里设置 CLASSPATH 的目的，在于告诉 Java 执行环境，在哪些目录下可以找到你所要执行的 Java 程序所需要的类或者包。）</p>

<p>我们这个例子中注册中心的实现类是 LocalFileNameService，它的实现比较简单，就是去读写一个本地文件，实现注册服务 registerService 方法时，把服务提供者保存到本地文件中；实现查找服务 lookupService 时，就是去本地文件中读出所有的服务提供者，找到对应的服务提供者，然后返回。</p>

<p>这里面有一点需要注意的是，由于这个本地文件它是一个共享资源，它会被 RPC 框架所有的客户端和服务端并发读写。所以，这时你要怎么做呢？对，<strong>必须要加锁！</strong></p>

<p>由于我们这个文件可能被多个进程读写，所以这里不能使用我们之前讲过的，编程语言提供的那些锁，原因是这些锁只能在进程内起作用，它锁不住其他进程。我们这里面必须使用由操作系统提供的文件锁。这个锁的使用和其他的锁并没有什么区别，同样是在访问共享文件之前先获取锁，访问共享资源结束后必须释放锁。具体的代码你可以去查看 LocalFileNameService 这个实现类。</p>

<h2 id="rpc-服务是怎么实现的">RPC 服务是怎么实现的？</h2>

<p>接下来，我们再来看看 RPC 服务是怎么实现的。RPC 服务也就是 RPC 框架的服务端。我们在之前讲解这个 RPC 框架的实现原理时讲到过，RPC 框架的服务端主要需要实现下面这两个功能：</p>

<ol>
<li>服务端的业务代码把服务的实现类注册到 RPC 框架中 ;</li>
<li>接收客户端桩发出的请求，调用服务的实现类并返回结果。</li>
</ol>

<p>把服务的实现类注册到 RPC 框架中，这个逻辑的实现很简单，我们只要使用一个合适的数据结构，记录下所有注册的实例就可以了，后面在处理客户端请求的时候，会用到这个数据结构来查找服务实例。</p>

<p>然后我们来看，RPC 框架的服务端如何来处理客户端发送的 RPC 请求。首先来看服务端中，使用 Netty 接收所有请求数据的处理类 RequestInvocation 的 channelRead0 方法。</p>

<pre><code>@Override
protected void channelRead0(ChannelHandlerContext channelHandlerContext, Command request) throws Exception {
    RequestHandler handler = requestHandlerRegistry.get(request.getHeader().getType());
    if(null != handler) {
        Command response = handler.handle(request);
        if(null != response) {
            channelHandlerContext.writeAndFlush(response).addListener((ChannelFutureListener) channelFuture -&gt; {
                if (!channelFuture.isSuccess()) {
                    logger.warn(&quot;Write response failed!&quot;, channelFuture.cause());
                    channelHandlerContext.channel().close();
                }
            });
        } else {
            logger.warn(&quot;Response is null!&quot;);
        }
    } else {
        throw new Exception(String.format(&quot;No handler for request with type: %d!&quot;, request.getHeader().getType()));
    }
}
</code></pre>

<p>这段代码的处理逻辑就是，根据请求命令的 Hdader 中的请求类型 type，去 requestHandlerRegistry 中查找对应的请求处理器 RequestHandler，然后调用请求处理器去处理请求，最后把结果发送给客户端。</p>

<p>这种通过“请求中的类型”，把请求分发到对应的处理类或者处理方法的设计，我们在 RocketMQ 和 Kafka 的源代码中都见到过，在服务端处理请求的场景中，这是一个很常用的方法。我们这里使用的也是同样的设计，不同的是，我们使用了一个命令注册机制，让这个路由分发的过程省略了大量的 if-else 或者是 switch 代码。这样做的好处是，可以很方便地扩展命令处理器，而不用修改路由分发的方法，并且代码看起来更加优雅。这个命令注册机制的实现类是 RequestHandlerRegistry，你可以自行去查看。</p>

<p>因为我们这个 RPC 框架中只需要处理一种类型的请求：RPC 请求，所以我们只实现了一个命令处理器：RpcRequestHandler。这部分代码是这个 RPC 框架服务端最核心的部分，你需要重点掌握。另外，为了便于你理解，在这里我只保留了核心业务逻辑，你在充分理解这部分核心业务逻辑之后，可以再去查看项目中完整的源代码，补全错误处理部分。</p>

<p>我们先来看它处理客户端请求，也就是这个 handle 方法的实现。</p>

<pre><code>@Override
public Command handle(Command requestCommand) {
    Header header = requestCommand.getHeader();
    // 从 payload 中反序列化 RpcRequest
    RpcRequest rpcRequest = SerializeSupport.parse(requestCommand.getPayload());
    // 查找所有已注册的服务提供方，寻找 rpcRequest 中需要的服务
    Object serviceProvider = serviceProviders.get(rpcRequest.getInterfaceName());
    // 找到服务提供者，利用 Java 反射机制调用服务的对应方法
    String arg = SerializeSupport.parse(rpcRequest.getSerializedArguments());
    Method method = serviceProvider.getClass().getMethod(rpcRequest.getMethodName(), String.class);
    String result = (String ) method.invoke(serviceProvider, arg);
    // 把结果封装成响应命令并返回
    return new Command(new ResponseHeader(type(), header.getVersion(), header.getRequestId()), SerializeSupport.serialize(result));
    // ...
}
</code></pre>

<ol>
<li>把 requestCommand 的 payload 属性反序列化成为 RpcRequest；</li>
<li>根据 rpcRequest 中的服务名，去成员变量 serviceProviders 中查找已注册服务实现类的实例；</li>
<li>找到服务提供者之后，利用 Java 反射机制调用服务的对应方法；</li>
<li>把结果封装成响应命令并返回，在 RequestInvocation 中，它会把这个响应命令发送给客户端。</li>
</ol>

<p>再来看成员变量 serviceProviders，它的定义是：Map<String/*service name*/, Object/*service provider*/> serviceProviders。它实际上就是一个 Map，Key 就是服务名，Value 就是服务提供方，也就是服务实现类的实例。这个 Map 的数据从哪儿来的呢？我们来看一下 RpcRequestHandler 这个类的定义：</p>

<pre><code>@Singleton
public class RpcRequestHandler implements RequestHandler, ServiceProviderRegistry {
    @Override
    public synchronized &lt;T&gt; void addServiceProvider(Class&lt;? extends T&gt; serviceClass, T serviceProvider) {
        serviceProviders.put(serviceClass.getCanonicalName(), serviceProvider);
        logger.info(&quot;Add service: {}, provider: {}.&quot;,
                serviceClass.getCanonicalName(),
                serviceProvider.getClass().getCanonicalName());
    }
    // ...
}
</code></pre>

<p>可以看到，这个类不仅实现了处理客户端请求的 RequestHandler 接口，同时还实现了注册 RPC 服务 ServiceProviderRegistry 接口，也就是说，RPC 框架服务端需要实现的两个功能——注册 RPC 服务和处理客户端 RPC 请求，都是在这一个类 RpcRequestHandler 中实现的，所以说，这个类是这个 RPC 框架服务端最核心的部分。成员变量 serviceProviders 这个 Map 中的数据，也就是在 addServiceProvider 这个方法的实现中添加进去的。</p>

<p>还有一点需要注意的是，我们 RpcRequestHandler 上增加了一个注解 @Singleton，限定这个类它是一个单例模式，这样确保在进程中任何一个地方，无论通过 ServiceSupport 获取 RequestHandler 或者 ServiceProviderRegistry 这两个接口的实现类，拿到的都是 RpcRequestHandler 这个类的唯一的一个实例。这个 @Singleton 的注解和获取单例的实现在 ServiceSupport 中，你可以自行查看代码。顺便说一句，在 Spring 中，也提供了单例 Bean 的支持，它的实现原理也是类似的。</p>

<h2 id="小结">小结</h2>

<p>以上就是实现这个 RPC 框架服务端的全部核心内容，照例我们来做一个总结。</p>

<p>首先我们一起来实现了一个注册中心，注册中心的接口设计采用了依赖倒置的设计原则（也就是“面向接口编程”的设计），并且还提供了一个“根据 URI 协议，自动加载对应实现类”的机制，使得我们可以通过扩展不同的协议，增加不同的注册中心实现。</p>

<p>这种“通过请求参数中的类型，来动态加载对应实现”的设计，在我们这个 RPC 框架中不止这一处用到，在“处理客户端命令并路由到对应的处理类”这部分代码中，使用的也是这样一种设计。</p>

<p>在 RPC 框架的服务端处理客户端请求的业务逻辑中，我们分两层做了两次请求分发：</p>

<ol>
<li>在 RequestInvocation 类中，根据请求命令中的请求类型 (command.getHeader().getType())，分发到对应的请求处理器 RequestHandler 中；</li>
<li>RpcRequestHandler 类中，根据 RPC 请求中的服务名，把 RPC 请求分发到对应的服务实现类的实例中去。</li>
</ol>

<p>这两次分发采用的设计是差不多的，但你需要注意的是，这并不是一种过度设计。原因是，我们这两次分发分别是在不同的业务抽象分层中，第一次分发是在服务端的网络传输层抽象中，它是网络传输的一部分，而第二次分发是 RPC 框架服务端的业务层，是 RPC 框架服务端的一部分。良好的分层设计，目的也是让系统各部分更加的“松耦合，高内聚”。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a0cccccc999491919097e0c7cdc1c9cc8ec3cfcd" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1587e67bc37771',t:'MTczNDA4ODQ2OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
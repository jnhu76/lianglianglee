<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=特别放送（一）经典的Kafka学习资料有哪些？>
        <link rel="icon" href="/static/favicon.png">
        <title>特别放送（一）经典的Kafka学习资料有哪些？ </title>
        
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
                        <a class="menu-item" id="00 导读 构建Kafka工程和源码阅读环境、Scala语言热身.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/00%20%e5%af%bc%e8%af%bb%20%e6%9e%84%e5%bb%baKafka%e5%b7%a5%e7%a8%8b%e5%92%8c%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb%e7%8e%af%e5%a2%83%e3%80%81Scala%e8%af%ad%e8%a8%80%e7%83%ad%e8%ba%ab.md">00 导读 构建Kafka工程和源码阅读环境、Scala语言热身.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词  阅读源码，逐渐成了职业进阶道路上的“必选项”.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e9%98%85%e8%af%bb%e6%ba%90%e7%a0%81%ef%bc%8c%e9%80%90%e6%b8%90%e6%88%90%e4%ba%86%e8%81%8c%e4%b8%9a%e8%bf%9b%e9%98%b6%e9%81%93%e8%b7%af%e4%b8%8a%e7%9a%84%e2%80%9c%e5%bf%85%e9%80%89%e9%a1%b9%e2%80%9d.md">00 开篇词  阅读源码，逐渐成了职业进阶道路上的“必选项”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 重磅加餐 带你快速入门Scala语言.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/00%20%e9%87%8d%e7%a3%85%e5%8a%a0%e9%a4%90%20%e5%b8%a6%e4%bd%a0%e5%bf%ab%e9%80%9f%e5%85%a5%e9%97%a8Scala%e8%af%ad%e8%a8%80.md">00 重磅加餐 带你快速入门Scala语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 日志段：保存消息文件的对象是怎么实现的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/01%20%e6%97%a5%e5%bf%97%e6%ae%b5%ef%bc%9a%e4%bf%9d%e5%ad%98%e6%b6%88%e6%81%af%e6%96%87%e4%bb%b6%e7%9a%84%e5%af%b9%e8%b1%a1%e6%98%af%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e7%9a%84%ef%bc%9f.md">01 日志段：保存消息文件的对象是怎么实现的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 日志（上）：日志究竟是如何加载日志段的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/02%20%e6%97%a5%e5%bf%97%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%97%a5%e5%bf%97%e7%a9%b6%e7%ab%9f%e6%98%af%e5%a6%82%e4%bd%95%e5%8a%a0%e8%bd%bd%e6%97%a5%e5%bf%97%e6%ae%b5%e7%9a%84%ef%bc%9f.md">02 日志（上）：日志究竟是如何加载日志段的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 日志（下）：彻底搞懂Log对象的常见操作.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/03%20%e6%97%a5%e5%bf%97%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%bd%bb%e5%ba%95%e6%90%9e%e6%87%82Log%e5%af%b9%e8%b1%a1%e7%9a%84%e5%b8%b8%e8%a7%81%e6%93%8d%e4%bd%9c.md">03 日志（下）：彻底搞懂Log对象的常见操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 索引（上）：改进的二分查找算法在Kafka索引的应用.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/04%20%e7%b4%a2%e5%bc%95%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%94%b9%e8%bf%9b%e7%9a%84%e4%ba%8c%e5%88%86%e6%9f%a5%e6%89%be%e7%ae%97%e6%b3%95%e5%9c%a8Kafka%e7%b4%a2%e5%bc%95%e7%9a%84%e5%ba%94%e7%94%a8.md">04 索引（上）：改进的二分查找算法在Kafka索引的应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 索引（下）：位移索引和时间戳索引的区别是什么？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/05%20%e7%b4%a2%e5%bc%95%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%bd%8d%e7%a7%bb%e7%b4%a2%e5%bc%95%e5%92%8c%e6%97%b6%e9%97%b4%e6%88%b3%e7%b4%a2%e5%bc%95%e7%9a%84%e5%8c%ba%e5%88%ab%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">05 索引（下）：位移索引和时间戳索引的区别是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 请求通道：如何实现Kafka请求队列？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/06%20%e8%af%b7%e6%b1%82%e9%80%9a%e9%81%93%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0Kafka%e8%af%b7%e6%b1%82%e9%98%9f%e5%88%97%ef%bc%9f.md">06 请求通道：如何实现Kafka请求队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 SocketServer（上）：Kafka到底是怎么应用NIO实现网络通信的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/07%20SocketServer%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aKafka%e5%88%b0%e5%ba%95%e6%98%af%e6%80%8e%e4%b9%88%e5%ba%94%e7%94%a8NIO%e5%ae%9e%e7%8e%b0%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1%e7%9a%84%ef%bc%9f.md">07 SocketServer（上）：Kafka到底是怎么应用NIO实现网络通信的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 SocketServer（中）：请求还要区分优先级？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/08%20SocketServer%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e8%af%b7%e6%b1%82%e8%bf%98%e8%a6%81%e5%8c%ba%e5%88%86%e4%bc%98%e5%85%88%e7%ba%a7%ef%bc%9f.md">08 SocketServer（中）：请求还要区分优先级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 SocketServer（下）：请求处理全流程源码分析.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/09%20SocketServer%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%af%b7%e6%b1%82%e5%a4%84%e7%90%86%e5%85%a8%e6%b5%81%e7%a8%8b%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90.md">09 SocketServer（下）：请求处理全流程源码分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 KafkaApis：Kafka最重要的源码入口，没有之一.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/10%20KafkaApis%ef%bc%9aKafka%e6%9c%80%e9%87%8d%e8%a6%81%e7%9a%84%e6%ba%90%e7%a0%81%e5%85%a5%e5%8f%a3%ef%bc%8c%e6%b2%a1%e6%9c%89%e4%b9%8b%e4%b8%80.md">10 KafkaApis：Kafka最重要的源码入口，没有之一.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Controller元数据：Controller都保存有哪些东西？有几种状态？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/11%20Controller%e5%85%83%e6%95%b0%e6%8d%ae%ef%bc%9aController%e9%83%bd%e4%bf%9d%e5%ad%98%e6%9c%89%e5%93%aa%e4%ba%9b%e4%b8%9c%e8%a5%bf%ef%bc%9f%e6%9c%89%e5%87%a0%e7%a7%8d%e7%8a%b6%e6%80%81%ef%bc%9f.md">11 Controller元数据：Controller都保存有哪些东西？有几种状态？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 ControllerChannelManager：Controller如何管理请求发送？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/12%20ControllerChannelManager%ef%bc%9aController%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e8%af%b7%e6%b1%82%e5%8f%91%e9%80%81%ef%bc%9f.md">12 ControllerChannelManager：Controller如何管理请求发送？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 ControllerEventManager：变身单线程后的Controller如何处理事件？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/13%20ControllerEventManager%ef%bc%9a%e5%8f%98%e8%ba%ab%e5%8d%95%e7%ba%bf%e7%a8%8b%e5%90%8e%e7%9a%84Controller%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e4%ba%8b%e4%bb%b6%ef%bc%9f.md">13 ControllerEventManager：变身单线程后的Controller如何处理事件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Controller选举是怎么实现的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/14%20Controller%e9%80%89%e4%b8%be%e6%98%af%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e7%9a%84%ef%bc%9f.md">14 Controller选举是怎么实现的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 如何理解Controller在Kafka集群中的作用？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/15%20%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3Controller%e5%9c%a8Kafka%e9%9b%86%e7%be%a4%e4%b8%ad%e7%9a%84%e4%bd%9c%e7%94%a8%ef%bc%9f.md">15 如何理解Controller在Kafka集群中的作用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 TopicDeletionManager： Topic是怎么被删除的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/16%20TopicDeletionManager%ef%bc%9a%20Topic%e6%98%af%e6%80%8e%e4%b9%88%e8%a2%ab%e5%88%a0%e9%99%a4%e7%9a%84%ef%bc%9f.md">16 TopicDeletionManager： Topic是怎么被删除的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 ReplicaStateMachine：揭秘副本状态机实现原理.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/17%20ReplicaStateMachine%ef%bc%9a%e6%8f%ad%e7%a7%98%e5%89%af%e6%9c%ac%e7%8a%b6%e6%80%81%e6%9c%ba%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">17 ReplicaStateMachine：揭秘副本状态机实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 PartitionStateMachine：分区状态转换如何实现？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/18%20PartitionStateMachine%ef%bc%9a%e5%88%86%e5%8c%ba%e7%8a%b6%e6%80%81%e8%bd%ac%e6%8d%a2%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">18 PartitionStateMachine：分区状态转换如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 TimingWheel：探究Kafka定时器背后的高效时间轮算法.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/19%20TimingWheel%ef%bc%9a%e6%8e%a2%e7%a9%b6Kafka%e5%ae%9a%e6%97%b6%e5%99%a8%e8%83%8c%e5%90%8e%e7%9a%84%e9%ab%98%e6%95%88%e6%97%b6%e9%97%b4%e8%bd%ae%e7%ae%97%e6%b3%95.md">19 TimingWheel：探究Kafka定时器背后的高效时间轮算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 DelayedOperation：Broker是怎么延时处理请求的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/20%20DelayedOperation%ef%bc%9aBroker%e6%98%af%e6%80%8e%e4%b9%88%e5%bb%b6%e6%97%b6%e5%a4%84%e7%90%86%e8%af%b7%e6%b1%82%e7%9a%84%ef%bc%9f.md">20 DelayedOperation：Broker是怎么延时处理请求的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 AbstractFetcherThread：拉取消息分几步？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/21%20AbstractFetcherThread%ef%bc%9a%e6%8b%89%e5%8f%96%e6%b6%88%e6%81%af%e5%88%86%e5%87%a0%e6%ad%a5%ef%bc%9f.md">21 AbstractFetcherThread：拉取消息分几步？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 ReplicaFetcherThread：Follower如何拉取Leader消息？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/22%20ReplicaFetcherThread%ef%bc%9aFollower%e5%a6%82%e4%bd%95%e6%8b%89%e5%8f%96Leader%e6%b6%88%e6%81%af%ef%bc%9f.md">22 ReplicaFetcherThread：Follower如何拉取Leader消息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 ReplicaManager（上）：必须要掌握的副本管理类定义和核心字段.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/23%20ReplicaManager%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%bf%85%e9%a1%bb%e8%a6%81%e6%8e%8c%e6%8f%a1%e7%9a%84%e5%89%af%e6%9c%ac%e7%ae%a1%e7%90%86%e7%b1%bb%e5%ae%9a%e4%b9%89%e5%92%8c%e6%a0%b8%e5%bf%83%e5%ad%97%e6%ae%b5.md">23 ReplicaManager（上）：必须要掌握的副本管理类定义和核心字段.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 ReplicaManager（中）：副本管理器是如何读写副本的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/24%20ReplicaManager%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%89%af%e6%9c%ac%e7%ae%a1%e7%90%86%e5%99%a8%e6%98%af%e5%a6%82%e4%bd%95%e8%af%bb%e5%86%99%e5%89%af%e6%9c%ac%e7%9a%84%ef%bc%9f.md">24 ReplicaManager（中）：副本管理器是如何读写副本的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 ReplicaManager（下）：副本管理器是如何管理副本的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/25%20ReplicaManager%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%89%af%e6%9c%ac%e7%ae%a1%e7%90%86%e5%99%a8%e6%98%af%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e5%89%af%e6%9c%ac%e7%9a%84%ef%bc%9f.md">25 ReplicaManager（下）：副本管理器是如何管理副本的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 MetadataCache：Broker是怎么异步更新元数据缓存的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/26%20MetadataCache%ef%bc%9aBroker%e6%98%af%e6%80%8e%e4%b9%88%e5%bc%82%e6%ad%a5%e6%9b%b4%e6%96%b0%e5%85%83%e6%95%b0%e6%8d%ae%e7%bc%93%e5%ad%98%e7%9a%84%ef%bc%9f.md">26 MetadataCache：Broker是怎么异步更新元数据缓存的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 消费者组元数据（上）：消费者组都有哪些元数据？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/27%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e5%85%83%e6%95%b0%e6%8d%ae%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e5%85%83%e6%95%b0%e6%8d%ae%ef%bc%9f.md">27 消费者组元数据（上）：消费者组都有哪些元数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 消费者组元数据（下）：Kafka如何管理这些元数据？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/28%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e5%85%83%e6%95%b0%e6%8d%ae%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aKafka%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e8%bf%99%e4%ba%9b%e5%85%83%e6%95%b0%e6%8d%ae%ef%bc%9f.md">28 消费者组元数据（下）：Kafka如何管理这些元数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 GroupMetadataManager：组元数据管理器是个什么东西？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/29%20GroupMetadataManager%ef%bc%9a%e7%bb%84%e5%85%83%e6%95%b0%e6%8d%ae%e7%ae%a1%e7%90%86%e5%99%a8%e6%98%af%e4%b8%aa%e4%bb%80%e4%b9%88%e4%b8%9c%e8%a5%bf%ef%bc%9f.md">29 GroupMetadataManager：组元数据管理器是个什么东西？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 GroupMetadataManager：位移主题保存的只是位移吗？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/30%20GroupMetadataManager%ef%bc%9a%e4%bd%8d%e7%a7%bb%e4%b8%bb%e9%a2%98%e4%bf%9d%e5%ad%98%e7%9a%84%e5%8f%aa%e6%98%af%e4%bd%8d%e7%a7%bb%e5%90%97%ef%bc%9f.md">30 GroupMetadataManager：位移主题保存的只是位移吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 GroupMetadataManager：查询位移时，不用读取位移主题？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/31%20GroupMetadataManager%ef%bc%9a%e6%9f%a5%e8%af%a2%e4%bd%8d%e7%a7%bb%e6%97%b6%ef%bc%8c%e4%b8%8d%e7%94%a8%e8%af%bb%e5%8f%96%e4%bd%8d%e7%a7%bb%e4%b8%bb%e9%a2%98%ef%bc%9f.md">31 GroupMetadataManager：查询位移时，不用读取位移主题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 GroupCoordinator：在Rebalance中，Coordinator如何处理成员入组？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/32%20GroupCoordinator%ef%bc%9a%e5%9c%a8Rebalance%e4%b8%ad%ef%bc%8cCoordinator%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e6%88%90%e5%91%98%e5%85%a5%e7%bb%84%ef%bc%9f.md">32 GroupCoordinator：在Rebalance中，Coordinator如何处理成员入组？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 GroupCoordinator：在Rebalance中，如何进行组同步？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/33%20GroupCoordinator%ef%bc%9a%e5%9c%a8Rebalance%e4%b8%ad%ef%bc%8c%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e7%bb%84%e5%90%8c%e6%ad%a5%ef%bc%9f.md">33 GroupCoordinator：在Rebalance中，如何进行组同步？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送（一）经典的Kafka学习资料有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%88%e4%b8%80%ef%bc%89%e7%bb%8f%e5%85%b8%e7%9a%84Kafka%e5%ad%a6%e4%b9%a0%e8%b5%84%e6%96%99%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">特别放送（一）经典的Kafka学习资料有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送（三）我是怎么度过日常一天的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%88%e4%b8%89%ef%bc%89%e6%88%91%e6%98%af%e6%80%8e%e4%b9%88%e5%ba%a6%e8%bf%87%e6%97%a5%e5%b8%b8%e4%b8%80%e5%a4%a9%e7%9a%84%ef%bc%9f.md">特别放送（三）我是怎么度过日常一天的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送（二）一篇文章带你了解参与开源社区的全部流程.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%88%e4%ba%8c%ef%bc%89%e4%b8%80%e7%af%87%e6%96%87%e7%ab%a0%e5%b8%a6%e4%bd%a0%e4%ba%86%e8%a7%a3%e5%8f%82%e4%b8%8e%e5%bc%80%e6%ba%90%e7%a4%be%e5%8c%ba%e7%9a%84%e5%85%a8%e9%83%a8%e6%b5%81%e7%a8%8b.md">特别放送（二）一篇文章带你了解参与开源社区的全部流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送（五） Kafka 社区的重磅功能：移除 ZooKeeper 依赖.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%88%e4%ba%94%ef%bc%89%20Kafka%20%e7%a4%be%e5%8c%ba%e7%9a%84%e9%87%8d%e7%a3%85%e5%8a%9f%e8%83%bd%ef%bc%9a%e7%a7%bb%e9%99%a4%20ZooKeeper%20%e4%be%9d%e8%b5%96.md">特别放送（五） Kafka 社区的重磅功能：移除 ZooKeeper 依赖.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送（四）20道经典的Kafka面试题详解.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%ef%bc%88%e5%9b%9b%ef%bc%8920%e9%81%93%e7%bb%8f%e5%85%b8%e7%9a%84Kafka%e9%9d%a2%e8%af%95%e9%a2%98%e8%af%a6%e8%a7%a3.md">特别放送（四）20道经典的Kafka面试题详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 源码学习，我们才刚上路呢.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%ba%90%e7%a0%81%e5%ad%a6%e4%b9%a0%ef%bc%8c%e6%88%91%e4%bb%ac%e6%89%8d%e5%88%9a%e4%b8%8a%e8%b7%af%e5%91%a2.md">结束语 源码学习，我们才刚上路呢.md</a>
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
                            <h1 id="title" data-id="特别放送（一）经典的Kafka学习资料有哪些？" class="title">特别放送（一）经典的Kafka学习资料有哪些？</h1>
                            <div><p>你好，我是胡夕。我们的课程已经更新一段时间了，你每节课都按时学了吗？如果你一直跟着学习的话，相信你一定会有很多收获的。</p>

<p>当然了，我也知道，在学习源码的过程中，除了有进步的快乐，可能还会有一些痛苦，毕竟，源码并不是那么容易就可以掌握的。</p>

<p>如果你在学习的过程中遇到任何问题，都可以给我留言，我会尽快回复你，帮助你解决问题。如果你发现自己被一些不好的情绪包围了，除了要努力坚持以外，我建议你学着从这种情绪中暂时跳脱出来，让自己转换到一些轻松的话题上。今天，我要讲的特别放送的内容就非常让人放松，因为我会给你分享一些Kafka的学习资料。</p>

<p>实际上，在更新的这段时间里，经常有同学问我：“老师，我想更多地了解下Kafka，你能给我推荐一些相关的学习资料吗？”今天，我就借着这个特别放送的环节，专门为你搜罗了各种Kafka学习资料，包括书籍、视频、博客等一切影音像资料，我还把它们做成了清单，一起分享给你。</p>

<p>这些资料的深浅程度不一样，有的偏重于基础理论，有的深入底层架构，有的侧重于实际案例，有的是分享最佳实践。</p>

<p>如果你期望得到实际使用上的指导，那么可以重点关注下我提到的社区维护的文档以及各类Kafka实战书籍。如果你对Kafka源码的学习兴趣更加浓厚，那么，这节课里的各类大神级的博客以及视频资料是你需要重点关注的。因为他们经常会直接给出源码级的分析，学习这类资料既能开拓我们的思路与视野，也能加深我们对源码知识的理解，可以说是具有多重好处。</p>

<p>总之，我建议你基于自身的学习目标与兴趣，有针对性地利用这些资料。</p>

<p>我把这份清单大体分为英文资料和中文资料两大部分，我先给出收集到的所有英文资料清单。</p>

<h2 id="英文资料">英文资料</h2>

<ol>
<li><a href="https://kafka.apache.org/documentation/" target="_blank">Apache Kafka官方网站</a></li>
</ol>

<p>我不知道你有没有认真地读过官网上面的文字，这里面的所有内容都是出自Kafka Committer之手，文字言简意赅，而且内容翔实丰富。我推荐你重点研读一下其中的<strong>Configuration篇</strong>、<strong>Operations篇</strong>以及<strong>Security篇</strong>，特别是Configuration中的参数部分。熟练掌握这些关键的参数配置，是进阶学习Kafka的必要条件。</p>

<ol>
<li>Confluent公司自己维护的<a href="http://docs.confluent.io/current/" target="_blank">官方文档</a></li>
</ol>

<p>Confluent公司是Kafka初创团队创建的商业化公司，主要是为了提供基于Kafka的商业化解决方案。我们经常把他们提供的产品称为Confluent Kafka。就我个人的感觉而言，这个公司的官网质量要比社区版（即Apache Kafka官网）上乘，特别是关于Security和Kafka Streams两部分的介绍，明显高出社区版一筹。因此，我推荐你重点学习Confluent官网上关于<a href="https://docs.confluent.io/current/security/index.html" target="_blank">Security配置</a>和<a href="https://docs.confluent.io/current/streams/index.html" target="_blank">Kafka Streams组件</a>的文档。</p>

<ol>
<li>Kafka的<a href="https://issues.apache.org/jira/issues/?filter=-4&amp;jql=project%20%3D%20KAFKA%20ORDER%20BY%20created%20DESC" target="_blank">Jira列表</a>，也就是我们俗称的Bug列表</li>
</ol>

<p>你可以在这个页面上搜索自己在实际环境中碰到的Kafka异常名称，然后结合自己的Kafka版本来看，这样的话，你通常能判断出该异常是不是由已知Bug造成的，从而避免浪费更多的时间去定位问题原因。另外，你还可以通过认领Jira的方式来为社区贡献代码。后面我会单独再用一节课的时间，给你具体介绍一下为社区贡献代码的完整流程。</p>

<ol>
<li>Kafka的<a href="https://cwiki.apache.org/confluence/display/KAFKA/Kafka+Improvement+Proposals" target="_blank">KIP网站</a></li>
</ol>

<p>KIP的完整名称是Kafka Improvement Proposals，即Kafka的新功能提案。在这里你能够了解到Kafka社区关于新功能的所有提案与相关讨论。有意思的是，有些KIP的讨论要比KIP本身精彩得多。针对各种新功能，全球开发者在这里审思明辨，彼此讨论，有时协同互利，有时针锋相对，实在是酣畅淋漓。KIP的另一大魅力则在于它非常民主——<strong>任何人都能申请新功能提案，将自己的想法付诸实践</strong>。</p>

<ol>
<li>StackOverflow的<a href="https://stackoverflow.com/questions/tagged/apache-kafka?sort=newest&amp;pageSize=15" target="_blank">Kafka专区</a></li>
</ol>

<p>大名鼎鼎的StackOverflow网站我就不过多介绍了。这上面的Kafka问题五花八门，而且难度参差不齐，不过你通常都能找到你想要的答案。同时，如果你对Kafka非常了解，不妨尝试回答一些问题，不仅能赚取一些积分，还能和全球的使用者一起交流，实在是一举两得。</p>

<ol>
<li><a href="https://www.confluent.io/blog/" target="_blank">Confluent博客</a></li>
</ol>

<p>这里面的文章含金量特别高，你一定要仔细研读下。举个简单的例子，我已经被太多人问过这样的问题了：“一个Kafka集群到底能支撑多少分区？”其实我们都知道这种问题是没有标准答案的，你需要了解的是原理！碰巧，Confluent博客上就有一篇这样的原理性解释<a href="https://www.confluent.jp/blog/apache-kafka-supports-200k-partitions-per-cluster/" target="_blank">文章</a>，是Kafka创始人饶军写的，你不妨看一看。</p>

<ol>
<li>Kafka社区非公开的各种资料，这包括社区维护的<a href="https://cwiki.apache.org/confluence/display/KAFKA/Index" target="_blank">Confluence文档</a>和Google Docs</li>
</ol>

<p>你几乎能在Confluence Space中找到所有的Kafka设计文档，其中关于Controller和新版Clients设计的文章非常值得一读；而Google Docs主要是两篇：一篇是Kafka事务的<a href="https://docs.google.com/document/d/11Jqy_GjUGtdXJK94XGsEIK7CP1SnQGdp2eF0wSw9ra8/edit" target="_blank">详细设计文档</a>，另一篇是<a href="https://docs.google.com/document/d/1rLDmzDOGQQeSiMANP0rC2RYp_L7nUGHzFD9MQISgXYM/edit" target="_blank">Controller重设计文档</a>。这两篇是我目前所见过的最详细的Kafka设计文档。国内的很多Kafka书籍和博客多是援引这两篇文章，甚至是直接翻译的，足见它们的价值非凡。</p>

<ol>
<li>Kafka社区的<a href="https://twitter.com/apachekafka" target="_blank">Twitter首页</a>和Confluent的<a href="https://twitter.com/confluentinc" target="_blank">Twitter首页</a></li>
</ol>

<p>你可能会说，Twitter算哪门子学习资料啊？但实际上，很多时候，你就是能够在这上面搜罗到有价值的Kafka文章，特别是Confluent的Twitter，它会定期推送关于Kafka最佳实践方面的内容。每次看到这类文章， 我都有一种意外淘到宝藏的感觉。我给你举个例子，Kafka Twitter在2019年10月12日转载了一篇名为<a href="https://medium.com/swlh/exploit-apache-kafkas-message-format-to-save-storage-and-bandwidth-7e0c533edf26" target="_blank"><em>Exploit Apache Kafka’s Message Format to Save Storage and Bandwidth</em></a> 的文章，这里面的内容水准很高，读起来非常过瘾，我建议你好好读一读。</p>

<ol>
<li>YouTube上的<a href="https://www.youtube.com/results?search_query=apache+kafka&amp;sp=EgIIAw%253D%253D" target="_blank">Kafka视频</a></li>
</ol>

<p>这些视频内容主要包括Kafka原理的讲解、业界牛人分享等。有的时候，你会发现，我们对文字类资料的掌握程度远不如看视频来得深入。如果你的英语基础还不错的话，我推荐你重点关注下YouTube上的这些分享视频。</p>

<p>好了，上面这九个资料就是我总结的Kafka英文学习资料。总体上看，这些资料都不要求太高的英文基础。即使是像YouTube上的英文视频，也是支持实时翻译的，你不用担心出现无法理解内容的情况。</p>

<p>接下来，我来给出中文资料清单。</p>

<h2 id="中文资料">中文资料</h2>

<p>首先，我给出我认为比较好的五本Apache Kafka书籍。</p>

<ol>
<li><a href="https://book.douban.com/subject/27665114/" target="_blank">《Kafka权威指南》</a></li>
</ol>

<p>这本书是“Kafka Definitive Guide”的中译本。实际上，这本书的中英两个版本我都看过，应该说中文版翻译得很棒，你直接看中译本就行了。这本书虽然很薄，但它包含的内容几乎囊括了Kafka的方方面面，而且这本书由Committer执笔，质量上无可挑剔。</p>

<ol>
<li><a href="https://book.douban.com/subject/27179953/" target="_blank">《Kafka技术内幕》</a></li>
</ol>

<p>这本书出版后一跃成为市面上能见到的Kafka最好书籍。这本书当得起“技术内幕”这四个字，里面很多Kafka原理的讲解清晰而深入，我自己读起来觉得收获很大。</p>

<ol>
<li><a href="https://book.douban.com/subject/30437872/" target="_blank">《深入理解Kafka：核心设计与实践原理》</a></li>
</ol>

<p>我与这本书的作者相识，他同时精通Kafka和RabbitMQ，可以说是消息中间件领域内的大家。这本书成书于2019年，是目前市面上最新的一本Kafka书籍。我推荐你买来读一下。</p>

<ol>
<li><a href="https://book.douban.com/subject/33425155/" target="_blank">《Kafka Streams实战》</a></li>
</ol>

<p>这本书是“Kafka Streams in action”的中译本，由Kafka Committer撰写而成。该书是我见到的<strong>最深入讲解Kafka Streams的书籍</strong>。如果你想学习基于Kafka Streams的实时流处理，那么这本书是不能不看的。</p>

<ol>
<li><a href="https://book.douban.com/subject/30221096/" target="_blank">《Apache Kafka实战》</a></li>
</ol>

<p>我这本书是2018年出版的，和之前那些面向Kafka设计原理的国内佳作不同的是，该书以讨论Kafka实际应用为主。我在里面分享了我这几年参与Kafka社区以及在使用Kafka的过程中积累的各种“江湖杂技”。如果你以使用为主，那么我推荐你看下这本书。</p>

<p>书籍的推荐告一段落，下面，我再介绍三个网站给你。</p>

<p>第一个是OrcHome。据我所知，OrcHome是国人自己维护的一个Kafka教程网站。这个网站最具特色的是它有个<a href="https://www.orchome.com/kafka/issues" target="_blank">Kafka问答区</a>，你可以在这上面提问，会有人专门负责解答你提出的问题。</p>

<p>第二个是我本人的<a href="https://www.cnblogs.com/huxi2b/" target="_blank">博客</a>。这个博客里大多是关于Kafka或者是其他大数据技术的原创文章。另外，我也会定期分享转载一些国内外的优秀博客。</p>

<p>第三个是知乎的<a href="https://www.zhihu.com/topic/20012159/newest" target="_blank">Kafka专区</a>。和StackOverflow不同的是，这个专区上的问题多以理论探讨为主。通常大家对于这类问题的解答还是很踊跃的，我也经常在这里回复问题。如果你对Kafka的某些原理想要做深入的了解，不妨在知乎的这个专区上提出你的问题，我相信很快就会有人回复的。</p>

<h2 id="总结">总结</h2>

<p>好了，上面的这些内容就是我总结的Kafka学习资料清单，希望它们对你是有帮助的。我把它们整理在了一张表格里，你可以重点看下。</p>

<p><img src="assets/4d773e45c4a3f86c5d9e86bb4a7ac7b9.jpg" alt="" /></p>

<p>另外，我强烈建议你把社区官网和Confluent官网文档仔仔细细地读几遍，我保证你一定会有很大的收获，毕竟，相比于清单上的其他项，官网文档是最最权威的第一手资料。</p>

<h2 id="课后讨论">课后讨论</h2>

<p>最后，我也请你分享一下你自己的Kafka学习书单、网站、影音像资料以及好的学习方法。</p>

<p>欢迎你在留言区畅所欲言，跟我交流讨论，也欢迎你把今天的内容分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#93ffffffaaa7a2a2a3a4d3f4fef2faffbdf0fcfe" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f104b51195563c6',t:'MTczNDAzMzU1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=12&#32;ControllerChannelManager：Controller如何管理请求发送？>
        <link rel="icon" href="/static/favicon.png">
        <title>12 ControllerChannelManager：Controller如何管理请求发送？ </title>
        
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
                            <h1 id="title" data-id="12 ControllerChannelManager：Controller如何管理请求发送？" class="title">12 ControllerChannelManager：Controller如何管理请求发送？</h1>
                            <div><p>你好，我是胡夕。上节课，我们深入研究了ControllerContext.scala源码文件，掌握了Kafka集群定义的重要元数据。今天，我们来学习下Controller是如何给其他Broker发送请求的。</p>

<p>掌握了这部分实现原理，你就能更好地了解Controller究竟是如何与集群Broker进行交互，从而实现管理集群元数据的功能的。而且，阅读这部分源码，还能帮你定位和解决线上问题。我先跟你分享一个真实的案例。</p>

<p>当时还是在Kafka 0.10.0.1时代，我们突然发现，在线上环境中，很多元数据变更无法在集群的所有Broker上同步了。具体表现为，创建了主题后，有些Broker依然无法感知到。</p>

<p>我的第一感觉是Controller出现了问题，但又苦于无从排查和验证。后来，我想到，会不会是Controller端请求队列中积压的请求太多造成的呢？因为当时Controller所在的Broker本身承载着非常重的业务，这是非常有可能的原因。</p>

<p>在看了相关代码后，我们就在相应的源码中新加了一个监控指标，用于实时监控Controller的请求队列长度。当更新到生产环境后，我们很轻松地定位了问题。果然，由于Controller所在的Broker自身负载过大，导致Controller端的请求积压，从而造成了元数据更新的滞后。精准定位了问题之后，解决起来就很容易了。后来，社区于0.11版本正式引入了相关的监控指标。</p>

<p>你看，阅读源码，除了可以学习优秀开发人员编写的代码之外，我们还能根据自身的实际情况做定制化方案，实现一些非开箱即用的功能。</p>

<h2 id="controller发送请求类型">Controller发送请求类型</h2>

<p>下面，我们就正式进入到Controller请求发送管理部分的学习。你可能会问：“Controller也会给Broker发送请求吗？”当然！<strong>Controller会给集群中的所有Broker（包括它自己所在的Broker）机器发送网络请求</strong>。发送请求的目的，是让Broker执行相应的指令。我用一张图，来展示下Controller都会发送哪些请求，如下所示：</p>

<p><img src="assets/3e8b0a34f003db5d67d5adafe8781ef7.jpg" alt="" /></p>

<p>当前，Controller只会向Broker发送三类请求，分别是LeaderAndIsrRequest、StopReplicaRequest和UpdateMetadataRequest。注意，这里我使用的是“当前”！我只是说，目前仅有这三类，不代表以后不会有变化。事实上，我几乎可以肯定，以后能发送的RPC协议种类一定会变化的。因此，你需要掌握请求发送的原理。毕竟，所有请求发送都是通过相同的机制完成的。</p>

<p>还记得我在[第8节课]提到的控制类请求吗？没错，这三类请求就是典型的控制类请求。我来解释下它们的作用。</p>

<ul>
<li>LeaderAndIsrRequest：最主要的功能是，告诉Broker相关主题各个分区的Leader副本位于哪台Broker上、ISR中的副本都在哪些Broker上。在我看来，<strong>它应该被赋予最高的优先级，毕竟，它有令数据类请求直接失效的本领</strong>。试想一下，如果这个请求中的Leader副本变更了，之前发往老的Leader的PRODUCE请求是不是全部失效了？因此，我认为它是非常重要的控制类请求。</li>
<li>StopReplicaRequest：告知指定Broker停止它上面的副本对象，该请求甚至还能删除副本底层的日志数据。这个请求主要的使用场景，是<strong>分区副本迁移</strong>和<strong>删除主题</strong>。在这两个场景下，都要涉及停掉Broker上的副本操作。</li>
<li>UpdateMetadataRequest：顾名思义，该请求会更新Broker上的元数据缓存。集群上的所有元数据变更，都首先发生在Controller端，然后再经由这个请求广播给集群上的所有Broker。在我刚刚分享的案例中，正是因为这个请求被处理得不及时，才导致集群Broker无法获取到最新的元数据信息。</li>
</ul>

<p>现在，社区越来越倾向于<strong>将重要的数据结构源代码从服务器端的core工程移动到客户端的clients工程中</strong>。这三类请求Java类的定义就封装在clients中，它们的抽象基类是AbstractControlRequest类，这个类定义了这三类请求的公共字段。</p>

<p>我用代码展示下这三类请求及其抽象父类的定义，以便让你对Controller发送的请求类型有个基本的认识。这些类位于clients工程下的src/main/java/org/apache/kafka/common/requests路径下。</p>

<p>先来看AbstractControlRequest类的主要代码：</p>

<pre><code>public abstract class AbstractControlRequest extends AbstractRequest {
    public static final long UNKNOWN_BROKER_EPOCH = -1L;
    public static abstract class Builder&lt;T extends AbstractRequest&gt; extends AbstractRequest.Builder&lt;T&gt; {
        protected final int controllerId;
        protected final int controllerEpoch;
        protected final long brokerEpoch;
        ......
}
</code></pre>

<p>区别于其他的数据类请求，抽象类请求必然包含3个字段。</p>

<ul>
<li><strong>controllerId</strong>：Controller所在的Broker ID。</li>
<li><strong>controllerEpoch</strong>：Controller的版本信息。</li>
<li><strong>brokerEpoch</strong>：目标Broker的Epoch。</li>
</ul>

<p>后面这两个Epoch字段用于隔离Zombie Controller和Zombie Broker，以保证集群的一致性。</p>

<p>在同一源码路径下，你能找到LeaderAndIsrRequest、StopReplicaRequest和UpdateMetadataRequest的定义，如下所示：</p>

<pre><code>public class LeaderAndIsrRequest extends AbstractControlRequest { ...... }
public class StopReplicaRequest extends AbstractControlRequest { ...... }
public class UpdateMetadataRequest extends AbstractControlRequest { ...... }
</code></pre>

<h2 id="requestsendthread">RequestSendThread</h2>

<p>说完了Controller发送什么请求，接下来我们说说怎么发。</p>

<p>Kafka源码非常喜欢生产者-消费者模式。该模式的好处在于，<strong>解耦生产者和消费者逻辑，分离两者的集中性交互</strong>。学完了“请求处理”模块，现在，你一定很赞同这个说法吧。还记得Broker端的SocketServer组件吗？它就在内部定义了一个线程共享的请求队列：它下面的Processor线程扮演Producer，而KafkaRequestHandler线程扮演Consumer。</p>

<p>对于Controller而言，源码同样使用了这个模式：它依然是一个线程安全的阻塞队列，Controller事件处理线程（第13节课会详细说它）负责向这个队列写入待发送的请求，而一个名为RequestSendThread的线程负责执行真正的请求发送。如下图所示：</p>

<p><img src="assets/825d084eb1517daace5532d1c93b0321.jpg" alt="" /></p>

<p>Controller会为集群中的每个Broker都创建一个对应的RequestSendThread线程。Broker上的这个线程，持续地从阻塞队列中获取待发送的请求。</p>

<p>那么，Controller往阻塞队列上放什么数据呢？这其实是由源码中的QueueItem类定义的。代码如下：</p>

<pre><code>case class QueueItem(apiKey: ApiKeys, request: AbstractControlRequest.Builder[_ &lt;: AbstractControlRequest], callback: AbstractResponse =&gt; Unit, enqueueTimeMs: Long)
</code></pre>

<p>每个QueueItem的核心字段都是<strong>AbstractControlRequest.Builder对象</strong>。你基本上可以认为，它就是阻塞队列上AbstractControlRequest类型。</p>

<p>需要注意的是这里的“&lt;:”符号，它在Scala中表示<strong>上边界</strong>的意思，即字段request必须是AbstractControlRequest的子类，也就是上面说到的那三类请求。</p>

<p>这也就是说，每个QueueItem实际保存的都是那三类请求中的其中一类。如果使用一个BlockingQueue对象来保存这些QueueItem，那么，代码就实现了一个请求阻塞队列。这就是RequestSendThread类做的事情。</p>

<p>接下来，我们就来学习下RequestSendThread类的定义。我给一些主要的字段添加了注释。</p>

<pre><code>class RequestSendThread(val controllerId: Int, // Controller所在Broker的Id
    val controllerContext: ControllerContext, // Controller元数据信息
    val queue: BlockingQueue[QueueItem], // 请求阻塞队列
    val networkClient: NetworkClient, // 用于执行发送的网络I/O类
    val brokerNode: Node, // 目标Broker节点
    val config: KafkaConfig, // Kafka配置信息
    val time: Time, 
    val requestRateAndQueueTimeMetrics: Timer,
    val stateChangeLogger: StateChangeLogger,
    name: String) extends ShutdownableThread(name = name) {
    ......
}
</code></pre>

<p>其实，RequestSendThread最重要的是它的<strong>doWork方法</strong>，也就是执行线程逻辑的方法：</p>

<pre><code>override def doWork(): Unit = {
    def backoff(): Unit = pause(100, TimeUnit.MILLISECONDS)
    val QueueItem(apiKey, requestBuilder, callback, enqueueTimeMs) = queue.take() // 以阻塞的方式从阻塞队列中取出请求
    requestRateAndQueueTimeMetrics.update(time.milliseconds() - enqueueTimeMs, TimeUnit.MILLISECONDS) // 更新统计信息
    var clientResponse: ClientResponse = null
    try {
      var isSendSuccessful = false
      while (isRunning &amp;&amp; !isSendSuccessful) {
        try {
          // 如果没有创建与目标Broker的TCP连接，或连接暂时不可用
          if (!brokerReady()) {
            isSendSuccessful = false
            backoff() // 等待重试
          }
          else {
            val clientRequest = networkClient.newClientRequest(brokerNode.idString, requestBuilder,
              time.milliseconds(), true)
            // 发送请求，等待接收Response
            clientResponse = NetworkClientUtils.sendAndReceive(networkClient, clientRequest, time)
            isSendSuccessful = true
          }
        } catch {
          case e: Throwable =&gt;
            warn(s&quot;Controller $controllerId epoch ${controllerContext.epoch} fails to send request $requestBuilder &quot; +
              s&quot;to broker $brokerNode. Reconnecting to broker.&quot;, e)
            // 如果出现异常，关闭与对应Broker的连接
            networkClient.close(brokerNode.idString)
            isSendSuccessful = false
            backoff()
        }
      }
      // 如果接收到了Response
      if (clientResponse != null) {
        val requestHeader = clientResponse.requestHeader
        val api = requestHeader.apiKey
        // 此Response的请求类型必须是LeaderAndIsrRequest、StopReplicaRequest或UpdateMetadataRequest中的一种
        if (api != ApiKeys.LEADER_AND_ISR &amp;&amp; api != ApiKeys.STOP_REPLICA &amp;&amp; api != ApiKeys.UPDATE_METADATA)
          throw new KafkaException(s&quot;Unexpected apiKey received: $apiKey&quot;)
        val response = clientResponse.responseBody
        stateChangeLogger.withControllerEpoch(controllerContext.epoch)
          .trace(s&quot;Received response &quot; +
          s&quot;${response.toString(requestHeader.apiVersion)} for request $api with correlation id &quot; +
          s&quot;${requestHeader.correlationId} sent to broker $brokerNode&quot;)

        if (callback != null) {
          callback(response) // 处理回调
        }
      }
    } catch {
      case e: Throwable =&gt;
        error(s&quot;Controller $controllerId fails to send a request to broker $brokerNode&quot;, e)
        networkClient.close(brokerNode.idString)
    }
  }
</code></pre>

<p>我用一张图来说明doWork的执行逻辑：</p>

<p><img src="assets/869727e22f882509a149d1065a8a1719.jpg" alt="" /></p>

<p>总体上来看，doWork的逻辑很直观。它的主要作用是从阻塞队列中取出待发送的请求，然后把它发送出去，之后等待Response的返回。在等待Response的过程中，线程将一直处于阻塞状态。当接收到Response之后，调用callback执行请求处理完成后的回调逻辑。</p>

<p>需要注意的是，RequestSendThread线程对请求发送的处理方式与Broker处理请求不太一样。它调用的sendAndReceive方法在发送完请求之后，会原地进入阻塞状态，等待Response返回。只有接收到Response，并执行完回调逻辑之后，该线程才能从阻塞队列中取出下一个待发送请求进行处理。</p>

<h2 id="controllerchannelmanager">ControllerChannelManager</h2>

<p>了解了RequestSendThread线程的源码之后，我们进入到ControllerChannelManager类的学习。</p>

<p>这个类和RequestSendThread是合作共赢的关系。在我看来，它有两大类任务。</p>

<ul>
<li>管理Controller与集群Broker之间的连接，并为每个Broker创建RequestSendThread线程实例；</li>
<li>将要发送的请求放入到指定Broker的阻塞队列中，等待该Broker专属的RequestSendThread线程进行处理。</li>
</ul>

<p>由此可见，它们是紧密相连的。</p>

<p>ControllerChannelManager类最重要的数据结构是brokerStateInfo，它是在下面这行代码中定义的：</p>

<pre><code>protected val brokerStateInfo = new HashMap[Int, ControllerBrokerStateInfo]
</code></pre>

<p>这是一个HashMap类型，Key是Integer类型，其实就是集群中Broker的ID信息，而Value是一个ControllerBrokerStateInfo。</p>

<p>你可能不太清楚ControllerBrokerStateInfo类是什么，我先解释一下。它本质上是一个POJO类，仅仅是承载若干数据结构的容器，如下所示：</p>

<pre><code>case class ControllerBrokerStateInfo(networkClient: NetworkClient,
    brokerNode: Node,
    messageQueue: BlockingQueue[QueueItem],
    requestSendThread: RequestSendThread,
    queueSizeGauge: Gauge[Int],
    requestRateAndTimeMetrics: Timer,
reconfigurableChannelBuilder: Option[Reconfigurable])
</code></pre>

<p>它有三个非常关键的字段。</p>

<ul>
<li><strong>brokerNode</strong>：目标Broker节点对象，里面封装了目标Broker的连接信息，比如主机名、端口号等。</li>
<li><strong>messageQueue</strong>：请求消息阻塞队列。你可以发现，Controller为每个目标Broker都创建了一个消息队列。</li>
<li><strong>requestSendThread</strong>：Controller使用这个线程给目标Broker发送请求。</li>
</ul>

<p>你一定要记住这三个字段，因为它们是实现Controller发送请求的关键因素。</p>

<p>为什么呢？我们思考一下，如果Controller要给Broker发送请求，肯定需要解决三个问题：发给谁？发什么？怎么发？“发给谁”就是由brokerNode决定的；messageQueue里面保存了要发送的请求，因而解决了“发什么”的问题；最后的“怎么发”就是依赖requestSendThread变量实现的。</p>

<p>好了，我们现在回到ControllerChannelManager。它定义了5个public方法，我来一一介绍下。</p>

<ul>
<li><strong>startup方法</strong>：Controller组件在启动时，会调用ControllerChannelManager的startup方法。该方法会从元数据信息中找到集群的Broker列表，然后依次为它们调用addBroker方法，把它们加到brokerStateInfo变量中，最后再依次启动brokerStateInfo中的RequestSendThread线程。</li>
<li><strong>shutdown方法</strong>：关闭所有RequestSendThread线程，并清空必要的资源。</li>
<li><strong>sendRequest方法</strong>：从名字看，就是发送请求，实际上就是把请求对象提交到请求队列。</li>
<li><strong>addBroker方法</strong>：添加目标Broker到brokerStateInfo数据结构中，并创建必要的配套资源，如请求队列、RequestSendThread线程对象等。最后，RequestSendThread启动线程。</li>
<li><strong>removeBroker方法</strong>：从brokerStateInfo移除目标Broker的相关数据。</li>
</ul>

<p>这里面大部分的方法逻辑都很简单，从方法名字就可以看得出来。我重点说一下<strong>addBroker</strong>，以及<strong>底层相关的私有方法addNewBroker和startRequestSendThread方法</strong>。</p>

<p>毕竟，addBroker是最重要的逻辑。每当集群中扩容了新的Broker时，Controller就会调用这个方法为新Broker增加新的RequestSendThread线程。</p>

<p>我们先来看addBroker：</p>

<pre><code>def addBroker(broker: Broker): Unit = {
    brokerLock synchronized {
      // 如果该Broker是新Broker的话
      if (!brokerStateInfo.contains(broker.id)) {
        // 将新Broker加入到Controller管理，并创建对应的RequestSendThread线程
        addNewBroker(broker) 
        // 启动RequestSendThread线程
        startRequestSendThread(broker.id)
      }
    }
  }
</code></pre>

<p>整个代码段被brokerLock保护起来了。还记得brokerStateInfo的定义吗？它仅仅是一个HashMap对象，因为不是线程安全的，所以任何访问该变量的地方，都需要锁的保护。</p>

<p>这段代码的逻辑是，判断目标Broker的序号，是否已经保存在brokerStateInfo中。如果是，就说明这个Broker之前已经添加过了，就没必要再次添加了；否则，addBroker方法会对目前的Broker执行两个操作：</p>

<ol>
<li>把该Broker节点添加到brokerStateInfo中；</li>
<li>启动与该Broker对应的RequestSendThread线程。</li>
</ol>

<p>这两步分别是由addNewBroker和startRequestSendThread方法实现的。</p>

<p>addNewBroker方法的逻辑比较复杂，我用注释的方式给出主要步骤：</p>

<pre><code>private def addNewBroker(broker: Broker): Unit = {
  // 为该Broker构造请求阻塞队列
  val messageQueue = new LinkedBlockingQueue[QueueItem]
  debug(s&quot;Controller ${config.brokerId} trying to connect to broker ${broker.id}&quot;)
  val controllerToBrokerListenerName = config.controlPlaneListenerName.getOrElse(config.interBrokerListenerName)
  val controllerToBrokerSecurityProtocol = config.controlPlaneSecurityProtocol.getOrElse(config.interBrokerSecurityProtocol)
  // 获取待连接Broker节点对象信息
  val brokerNode = broker.node(controllerToBrokerListenerName)
  val logContext = new LogContext(s&quot;[Controller id=${config.brokerId}, targetBrokerId=${brokerNode.idString}] &quot;)
  val (networkClient, reconfigurableChannelBuilder) = {
    val channelBuilder = ChannelBuilders.clientChannelBuilder(
      controllerToBrokerSecurityProtocol,
      JaasContext.Type.SERVER,
      config,
      controllerToBrokerListenerName,
      config.saslMechanismInterBrokerProtocol,
      time,
      config.saslInterBrokerHandshakeRequestEnable,
      logContext
    )
    val reconfigurableChannelBuilder = channelBuilder match {
      case reconfigurable: Reconfigurable =&gt;
        config.addReconfigurable(reconfigurable)
        Some(reconfigurable)
      case _ =&gt; None
    }
    // 创建NIO Selector实例用于网络数据传输
    val selector = new Selector(
      NetworkReceive.UNLIMITED,
      Selector.NO_IDLE_TIMEOUT_MS,
      metrics,
      time,
      &quot;controller-channel&quot;,
      Map(&quot;broker-id&quot; -&gt; brokerNode.idString).asJava,
      false,
      channelBuilder,
      logContext
    )
    // 创建NetworkClient实例
    // NetworkClient类是Kafka clients工程封装的顶层网络客户端API
    // 提供了丰富的方法实现网络层IO数据传输
    val networkClient = new NetworkClient(
      selector,
      new ManualMetadataUpdater(Seq(brokerNode).asJava),
      config.brokerId.toString,
      1,
      0,
      0,
      Selectable.USE_DEFAULT_BUFFER_SIZE,
      Selectable.USE_DEFAULT_BUFFER_SIZE,
      config.requestTimeoutMs,
      ClientDnsLookup.DEFAULT,
      time,
      false,
      new ApiVersions,
      logContext
    )
    (networkClient, reconfigurableChannelBuilder)
  }
  // 为这个RequestSendThread线程设置线程名称
  val threadName = threadNamePrefix match {
    case None =&gt; s&quot;Controller-${config.brokerId}-to-broker-${broker.id}-send-thread&quot;
    case Some(name) =&gt; s&quot;$name:Controller-${config.brokerId}-to-broker-${broker.id}-send-thread&quot;
  }
  // 构造请求处理速率监控指标
  val requestRateAndQueueTimeMetrics = newTimer(
    RequestRateAndQueueTimeMetricName, TimeUnit.MILLISECONDS, TimeUnit.SECONDS, brokerMetricTags(broker.id)
  )
  // 创建RequestSendThread实例
  val requestThread = new RequestSendThread(config.brokerId, controllerContext, messageQueue, networkClient,
    brokerNode, config, time, requestRateAndQueueTimeMetrics, stateChangeLogger, threadName)
  requestThread.setDaemon(false)

  val queueSizeGauge = newGauge(QueueSizeMetricName, () =&gt; messageQueue.size, brokerMetricTags(broker.id))
  // 创建该Broker专属的ControllerBrokerStateInfo实例
  // 并将其加入到brokerStateInfo统一管理
  brokerStateInfo.put(broker.id, ControllerBrokerStateInfo(networkClient, brokerNode, messageQueue,
    requestThread, queueSizeGauge, requestRateAndQueueTimeMetrics, reconfigurableChannelBuilder))
}
</code></pre>

<p>为了方便你理解，我还画了一张流程图形象说明它的执行流程：</p>

<p><img src="assets/4f34c319f9480c16ac12dee78d5ba322.jpg" alt="" /></p>

<p>addNewBroker的关键在于，<strong>要为目标Broker创建一系列的配套资源</strong>，比如，NetworkClient用于网络I/O操作、messageQueue用于阻塞队列、requestThread用于发送请求，等等。</p>

<p>至于startRequestSendThread方法，就简单得多了，只有几行代码而已。</p>

<pre><code>protected def startRequestSendThread(brokerId: Int): Unit = {
  // 获取指定Broker的专属RequestSendThread实例
  val requestThread = brokerStateInfo(brokerId).requestSendThread
  if (requestThread.getState == Thread.State.NEW)
    // 启动线程
    requestThread.start()
}
</code></pre>

<p>它首先根据给定的Broker序号信息，从brokerStateInfo中找出对应的ControllerBrokerStateInfo对象。有了这个对象，也就有了为该目标Broker服务的所有配套资源。下一步就是从ControllerBrokerStateInfo中拿出RequestSendThread对象，再启动它就好了。</p>

<h2 id="总结">总结</h2>

<p>今天，我结合ControllerChannelManager.scala文件，重点分析了Controller向Broker发送请求机制的实现原理。</p>

<p>Controller主要通过ControllerChannelManager类来实现与其他Broker之间的请求发送。其中，ControllerChannelManager类中定义的RequestSendThread是主要的线程实现类，用于实际发送请求给集群Broker。除了RequestSendThread之外，ControllerChannelManager还定义了相应的管理方法，如添加Broker、移除Broker等。通过这些管理方法，Controller在集群扩缩容时能够快速地响应到这些变化，完成对应Broker连接的创建与销毁。</p>

<p>我们来回顾下这节课的重点。</p>

<ul>
<li>Controller端请求：Controller发送三类请求给Broker，分别是LeaderAndIsrRequest、StopReplicaRequest和UpdateMetadataRequest。</li>
<li>RequestSendThread：该线程负责将请求发送给集群中的相关或所有Broker。</li>
<li>请求阻塞队列+RequestSendThread：Controller会为集群上所有Broker创建对应的请求阻塞队列和RequestSendThread线程。</li>
</ul>

<p>其实，今天讲的所有东西都只是这节课的第二张图中“消费者”的部分，我们并没有详细了解请求是怎么被放到请求队列中的。接下来，我们就会针对这个问题，深入地去探讨Controller单线程的事件处理器是如何实现的。</p>

<p><img src="assets/00fce28d26a94389f2bb5e957b650bb9.jpg" alt="" /></p>

<h2 id="课后讨论">课后讨论</h2>

<p>你觉得，为每个Broker都创建一个RequestSendThread的方案有什么优缺点？</p>

<p>欢迎你在留言区写下你的思考和答案，跟我交流讨论，也欢迎你把今天的内容分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#18747474212c2929282f587f75797174367b7775" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fb3cd6e307791',t:'MTczNDAyNzM1My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=00&#32;导读&#32;构建Kafka工程和源码阅读环境、Scala语言热身>
        <link rel="icon" href="/static/favicon.png">
        <title>00 导读 构建Kafka工程和源码阅读环境、Scala语言热身 </title>
        
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
                            <h1 id="title" data-id="00 导读 构建Kafka工程和源码阅读环境、Scala语言热身" class="title">00 导读 构建Kafka工程和源码阅读环境、Scala语言热身</h1>
                            <div><p>你好，我是胡夕。</p>

<p>从今天开始，我们就要正式走入Kafka源码的世界了。既然咱们这个课程是教你阅读Kafka源码的，那么，你首先就得掌握如何在自己的电脑上搭建Kafka的源码环境，甚至是知道怎么对它们进行调试。在这节课，我展示了很多实操步骤，建议你都跟着操作一遍，否则很难会有特别深刻的认识。</p>

<p>话不多说，现在，我们就先来搭建源码环境吧。</p>

<h2 id="环境准备">环境准备</h2>

<p>在阅读Kafka源码之前，我们要先做一些必要的准备工作。这涉及到一些工具软件的安装，比如Java、Gradle、Scala、IDE、Git，等等。</p>

<p>如果你是在Linux或Mac系统下搭建环境，你需要安装Java、IDE和Git；如果你使用的是Windows，那么你需要全部安装它们。</p>

<p>咱们这个课程统一使用下面的版本进行源码讲解。</p>

<ul>
<li>Oracle Java 8：我们使用的是Oracle的JDK及Hotspot JVM。如果你青睐于其他厂商或开源的Java版本（比如OpenJDK），你可以选择安装不同厂商的JVM版本。</li>
<li>Gradle 6.3：我在这门课里带你阅读的Kafka源码是社区的Trunk分支。Trunk分支目前演进到了2.5版本，已经支持Gradle 6.x版本。你最好安装Gradle 6.3或更高版本。</li>
<li>Scala 2.13：社区Trunk分支编译当前支持两个Scala版本，分别是2.12和2.13。默认使用2.13进行编译，因此我推荐你安装Scala 2.13版本。</li>
<li>IDEA + Scala插件：这门课使用IDEA作为IDE来阅读和配置源码。我对Eclipse充满敬意，只是我个人比较习惯使用IDEA。另外，你需要为IDEA安装Scala插件，这样可以方便你阅读Scala源码。</li>
<li>Git：安装Git主要是为了管理Kafka源码版本。如果你要成为一名社区代码贡献者，Git管理工具是必不可少的。</li>
</ul>

<h2 id="构建kafka工程">构建Kafka工程</h2>

<p>等你准备好以上这些之后，我们就可以来构建Kafka工程了。</p>

<p>首先，我们下载Kafka源代码。方法很简单，找一个干净的源码路径，然后执行下列命令去下载社区的Trunk代码即可：</p>

<pre><code>$ git clone https://github.com/apache/kafka.git
</code></pre>

<p>在漫长的等待之后，你的路径上会新增一个名为kafka的子目录，它就是Kafka项目的根目录。如果在你的环境中，上面这条命令无法执行的话，你可以在浏览器中输入<a href="https://codeload.github.com/apache/kafka/zip/trunk" target="_blank">https://codeload.github.com/apache/kafka/zip/trunk</a>下载源码ZIP包并解压，只是这样你就失去了Git管理，你要手动链接到远程仓库，具体方法可以参考这篇<a href="https://help.github.com/articles/fork-a-repo/" target="_blank">Git文档</a>。</p>

<p>下载完成后，你要进入工程所在的路径下，也就是进入到名为kafka的路径下，然后执行相应的命令来构建Kafka工程。</p>

<p>如果你是在Mac或Linux平台上搭建环境，那么直接运行下列命令构建即可：</p>

<pre><code>$ ./gradlew build
</code></pre>

<p>该命令首先下载Gradle Wrapper所需的jar文件，然后对Kafka工程进行构建。需要注意的是，在执行这条命令时，你很可能会遇到下面的这个异常：</p>

<pre><code>Failed to connect to raw.githubusercontent.com port 443: Connection refused
</code></pre>

<p>如果碰到了这个异常，你也不用惊慌，你可以去这个<a href="https://raw.githubusercontent.com/gradle/gradle/v6.3.0/gradle/wrapper/gradle-wrapper.jar" target="_blank">官网链接</a>或者是我提供的<a href="https://pan.baidu.com/s/1tuVHunoTwHfbtoqMvoTNoQ" target="_blank">链接</a>（提取码：ntvd）直接下载Wrapper所需的Jar包，手动把这个Jar文件拷贝到kafka路径下的gradle/wrapper子目录下，然后重新执行gradlew build命令去构建工程。</p>

<p>我想提醒你的是，官网链接包含的版本号是v6.3.0，但是该版本后续可能会变化，因此，你最好先打开gradlew文件，去看一下社区使用的是哪个版本的Gradle。<strong>一旦你发现版本不再是v6.3.0了，那就不要再使用我提供的链接了。这个时候，你需要直接去官网下载对应版本的Jar包</strong>。</p>

<p>举个例子，假设gradlew文件中使用的Gradle版本变更为v6.4.0，那么你需要把官网链接URL中的版本号修改为v6.4.0，然后去下载这个版本的Wrapper Jar包。</p>

<p>如果你是在Windows平台上构建，那你就不能使用Gradle Wrapper了，因为Kafka没有提供Windows平台上可运行的Wrapper Bat文件。这个时候，你只能使用你自己的环境中自行安装的Gradle。具体命令是：</p>

<pre><code>kafka&gt; gradle.bat build
</code></pre>

<p>无论是gradle.bat build命令，还是gradlew build命令，首次运行时都要花费相当长的时间去下载必要的Jar包，你要耐心地等待。</p>

<p>下面，我用一张图给你展示下Kafka工程的各个目录以及文件：</p>

<p><img src="assets/a2ef664cd8d5494f55919643df1305f7.png" alt="" /></p>

<p>这里我再简单介绍一些主要的组件路径。</p>

<ul>
<li><strong>bin目录</strong>：保存Kafka工具行脚本，我们熟知的kafka-server-start和kafka-console-producer等脚本都存放在这里。</li>
<li><strong>clients目录</strong>：保存Kafka客户端代码，比如生产者和消费者的代码都在该目录下。</li>
<li><strong>config目录</strong>：保存Kafka的配置文件，其中比较重要的配置文件是server.properties。</li>
<li><strong>connect目录</strong>：保存Connect组件的源代码。我在开篇词里提到过，Kafka Connect组件是用来实现Kafka与外部系统之间的实时数据传输的。</li>
<li><strong>core目录</strong>：保存Broker端代码。Kafka服务器端代码全部保存在该目录下。</li>
<li><strong>streams目录</strong>：保存Streams组件的源代码。Kafka Streams是实现Kafka实时流处理的组件。</li>
</ul>

<p>其他的目录要么不太重要，要么和配置相关，这里我就不展开讲了。</p>

<p>除了上面的gradlew build命令之外，我再介绍一些常用的构建命令，帮助你调试Kafka工程。</p>

<p>我们先看一下测试相关的命令。Kafka源代码分为4大部分：Broker端代码、Clients端代码、Connect端代码和Streams端代码。如果你想要测试这4个部分的代码，可以分别运行以下4条命令：</p>

<pre><code>$ ./gradlew core:test
$ ./gradlew clients:test
$ ./gradlew connect:[submodule]:test
$ ./gradlew streams:test
</code></pre>

<p>你可能注意到了，在这4条命令中，Connect组件的测试方法不太一样。这是因为Connect工程下细分了多个子模块，比如api、runtime等，所以，你需要显式地指定要测试的子模块名，才能进行测试。</p>

<p>如果你要单独对某一个具体的测试用例进行测试，比如单独测试Broker端core包的LogTest类，可以用下面的命令：</p>

<pre><code>$ ./gradlew core:test --tests kafka.log.LogTest
</code></pre>

<p>另外，如果你要构建整个Kafka工程并打包出一个可运行的二进制环境，就需要运行下面的命令：</p>

<pre><code>$ ./gradlew clean releaseTarGz
</code></pre>

<p>成功运行后，core、clients和streams目录下就会分别生成对应的二进制发布包，它们分别是：</p>

<ul>
<li><strong>kafka-2.12-2.5.0-SNAPSHOT.tgz</strong>。它是Kafka的Broker端发布包，把该文件解压之后就是标准的Kafka运行环境。该文件位于core路径的/build/distributions目录。</li>
<li><strong>kafka-clients-2.5.0-SNAPSHOT.jar</strong>。该Jar包是Clients端代码编译打包之后的二进制发布包。该文件位于clients目录下的/build/libs目录。</li>
<li><strong>kafka-streams-2.5.0-SNAPSHOT.jar</strong>。该Jar包是Streams端代码编译打包之后的二进制发布包。该文件位于streams目录下的/build/libs目录。</li>
</ul>

<h2 id="搭建源码阅读环境">搭建源码阅读环境</h2>

<p>刚刚我介绍了如何使用Gradle工具来构建Kafka项目工程，现在我来带你看一下如何利用IDEA搭建Kafka源码阅读环境。实际上，整个过程非常简单。我们打开IDEA，点击“文件”，随后点击“打开”，选择上一步中的Kafka文件路径即可。</p>

<p>项目工程被导入之后，IDEA会对项目进行自动构建，等构建完成之后，你可以找到core目录源码下的Kafka.scala文件。打开它，然后右键点击Kafka，你应该就能看到这样的输出结果了：</p>

<p><img src="assets/ce0a63e7627c641da471b48a62860ad2.png" alt="" /></p>

<p>这就是无参执行Kafka主文件的运行结果。通过这段输出，我们能够学会启动Broker所必需的参数，即指定server.properties文件的地址。这也是启动Kafka Broker的标准命令。</p>

<p>在开篇词中我也说了，这个课程会聚焦于讲解Kafka Broker端源代码。因此，在正式学习这部分源码之前，我先给你简单介绍一下Broker端源码的组织架构。下图展示了Kafka core包的代码架构：</p>

<p><img src="assets/dfdd73cc95ecc5390ebeb73c324437b2.png" alt="" /></p>

<p>我来给你解释几个比较关键的代码包。</p>

<ul>
<li>controller包：保存了Kafka控制器（Controller）代码，而控制器组件是Kafka的核心组件，后面我们会针对这个包的代码进行详细分析。</li>
<li>coordinator包：保存了<strong>消费者端的GroupCoordinator代码</strong>和<strong>用于事务的TransactionCoordinator代码</strong>。对coordinator包进行分析，特别是对消费者端的GroupCoordinator代码进行分析，是我们弄明白Broker端协调者组件设计原理的关键。</li>
<li>log包：保存了Kafka最核心的日志结构代码，包括日志、日志段、索引文件等，后面会有详细介绍。另外，该包下还封装了Log Compaction的实现机制，是非常重要的源码包。</li>
<li>network包：封装了Kafka服务器端网络层的代码，特别是SocketServer.scala这个文件，是Kafka实现Reactor模式的具体操作类，非常值得一读。</li>
<li>server包：顾名思义，它是Kafka的服务器端主代码，里面的类非常多，很多关键的Kafka组件都存放在这里，比如后面要讲到的状态机、Purgatory延时机制等。</li>
</ul>

<p>在后续的课程中，我会挑选Kafka最主要的代码类进行详细分析，帮助你深入了解Kafka Broker端重要组件的实现原理。</p>

<p>另外，虽然这门课不会涵盖测试用例的代码分析，但在我看来，<strong>弄懂测试用例是帮助你快速了解Kafka组件的最有效的捷径之一</strong>。如果时间允许的话，我建议你多读一读Kafka各个组件下的测试用例，它们通常都位于代码包的src/test目录下。拿Kafka日志源码Log来说，它对应的LogTest.scala测试文件就写得非常完备，里面多达几十个测试用例，涵盖了Log的方方面面，你一定要读一下。</p>

<h2 id="scala-语言热身">Scala 语言热身</h2>

<p>因为Broker端的源码完全是基于Scala的，所以在开始阅读这部分源码之前，我还想花一点时间快速介绍一下 Scala 语言的语法特点。我先拿几个真实的 Kafka 源码片段来帮你热热身。</p>

<p>先来看第一个：</p>

<pre><code>def sizeInBytes(segments: Iterable[LogSegment]): Long =
    segments.map(_.size.toLong).sum
</code></pre>

<p>这是一个典型的 Scala 方法，方法名是 sizeInBytes。它接收一组 LogSegment 对象，返回一个长整型。LogSegment 对象就是我们后面要谈到的日志段。你在 Kafka 分区目录下看到的每一个.log 文件本质上就是一个 LogSegment。从名字上来看，这个方法计算的是这组 LogSegment 的总字节数。具体方法是遍历每个输入 LogSegment，调用其 size 方法并将其累加求和之后返回。</p>

<p>再来看一个：</p>

<pre><code>val firstOffset: Option[Long] = ......

def numMessages: Long = {
    firstOffset match {
      case Some(firstOffsetVal) if (firstOffsetVal &gt;= 0 &amp;&amp; lastOffset &gt;= 0) =&gt; (lastOffset - firstOffsetVal + 1)
      case _ =&gt; 0
    }
  }
</code></pre>

<p>该方法是 LogAppendInfo 对象的一个方法，统计的是 Broker 端一次性批量写入的消息数。这里你需要重点关注 <strong>match</strong> 和 <strong>case</strong> 这两个关键字，你可以近似地认为它们等同于 Java 中的 switch，但它们的功能要强大得多。该方法统计写入消息数的逻辑是：如果 firstOffsetVal 和 lastOffset 值都大于 0，则写入消息数等于两者的差值+1；如果不存在 firstOffsetVal，则无法统计写入消息数，简单返回 0 即可。</p>

<p>倘若对你而言，弄懂上面这两段代码有些吃力，我建议你去快速地学习一下Scala语言。重点学什么呢？我建议你重点学习下Scala中对于<strong>集合的遍历语法</strong>，以及<strong>基于match的模式匹配用法</strong>。</p>

<p>另外，由于Scala具有的函数式编程风格，你至少<strong>要理解Java中Lambda表达式的含义</strong>，这会在很大程度上帮你扫清阅读障碍。</p>

<p>相反地，如果上面的代码对你来讲很容易理解，那么，读懂Broker端80%的源码应该没有什么问题。你可能还会关心，剩下的那晦涩难懂的20%源码怎么办呢？其实没关系，你可以等慢慢精通了Scala语言之后再进行阅读，它们不会对你熟练掌握核心源码造成影响的。另外，后面涉及到比较难的Scala语法特性时，我还会再具体给你解释的，所以，还是那句话，你完全不用担心语言的问题！</p>

<h2 id="总结">总结</h2>

<p>今天是我们开启Kafka源码分析的“热身课”，我给出了构建Kafka工程以及搭建Kafka源码阅读环境的具体方法。我建议你对照上面的内容完整地走一遍流程，亲身体会一下Kafka工程的构建与源码工程的导入。毕竟，这些都是后面阅读具体Kafka代码的前提条件。</p>

<p>最后我想再强调一下，阅读任何一个大型项目的源码都不是一件容易的事情，我希望你在任何时候都不要轻言放弃。很多时候，碰到读不懂的代码你就多读几遍，也许稍后就会有醍醐灌顶的感觉。</p>

<h2 id="课后讨论">课后讨论</h2>

<p>熟悉Kafka的话，你一定听说过kafka-console-producer.sh脚本。我前面提到过，该脚本位于工程的bin目录下，你能找到它对应的Java类是哪个文件吗？这个搜索过程能够给你一些寻找Kafka所需类文件的好思路，你不妨去试试看。</p>

<p>欢迎你在留言区畅所欲言，跟我交流讨论，也欢迎你把文章分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#761a1a1a4f424747464136111b171f1a5815191b" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fb086fc2a7791',t:'MTczNDAyNzIxOS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
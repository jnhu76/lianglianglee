<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=11&#32;Controller元数据：Controller都保存有哪些东西？有几种状态？>
        <link rel="icon" href="/static/favicon.png">
        <title>11 Controller元数据：Controller都保存有哪些东西？有几种状态？ </title>
        
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
                            <h1 id="title" data-id="11 Controller元数据：Controller都保存有哪些东西？有几种状态？" class="title">11 Controller元数据：Controller都保存有哪些东西？有几种状态？</h1>
                            <div><p>你好，我是胡夕。从今天开始，我们正式进入到第三大模块的学习：控制器（Controller）模块 。</p>

<p>提起Kafka中的Controller组件，我相信你一定不陌生。从某种意义上说，它是Kafka最核心的组件。一方面，它要为集群中的所有主题分区选举领导者副本；另一方面，它还承载着集群的全部元数据信息，并负责将这些元数据信息同步到其他Broker上。既然我们是Kafka源码解读课，那就绝对不能错过这么重量级的组件。</p>

<p>我画了一张图片，希望借助它帮你建立起对这个模块的整体认知。今天，我们先学习下Controller元数据。</p>

<p><img src="assets/13c0d8b3f52c295c70c71a154dae185f.jpg" alt="" /></p>

<h2 id="案例分享">案例分享</h2>

<p>在正式学习源码之前，我想向你分享一个真实的案例。</p>

<p>在我们公司的Kafka集群环境上，曾经出现了一个比较“诡异”的问题：某些核心业务的主题分区一直处于“不可用”状态。</p>

<p>通过使用“kafka-topics”命令查询，我们发现，这些分区的Leader显示是-1。之前，这些Leader所在的Broker机器因为负载高宕机了，当Broker重启回来后，Controller竟然无法成功地为这些分区选举Leader，因此，它们一直处于“不可用”状态。</p>

<p>由于是生产环境，我们的当务之急是马上恢复受损分区，然后才能调研问题的原因。有人提出，重启这些分区旧Leader所在的所有Broker机器——这很容易想到，毕竟“重启大法”一直很好用。但是，这一次竟然没有任何作用。</p>

<p>之后，有人建议升级重启大法，即重启集群的所有Broker——这在当时是不能接受的。且不说有很多业务依然在运行着，单是重启Kafka集群本身，就是一件非常缺乏计划性的事情。毕竟，生产环境怎么能随意重启呢？！</p>

<p>后来，我突然想到了Controller组件中重新选举Controller的代码。一旦Controller被选举出来，它就会向所有Broker更新集群元数据，也就是说，会“重刷”这些分区的状态。</p>

<p>那么问题来了，我们如何在避免重启集群的情况下，干掉已有Controller并执行新的Controller选举呢？答案就在源码中的<strong>ControllerZNode.path</strong>上，也就是ZooKeeper的/controller节点。倘若我们手动删除了/controller节点，Kafka集群就会触发Controller选举。于是，我们马上实施了这个方案，效果出奇得好：之前的受损分区全部恢复正常，业务数据得以正常生产和消费。</p>

<p>当然，给你分享这个案例的目的，并不是让你记住可以随意干掉/controller节点——这个操作其实是有一点危险的。事实上，我只是想通过这个真实的例子，向你说明，很多打开“精通Kafka之门”的钥匙是隐藏在源码中的。那么，接下来，我们就开始找“钥匙”吧。</p>

<h2 id="集群元数据">集群元数据</h2>

<p>想要完整地了解Controller的工作原理，我们首先就要学习它管理了哪些数据。毕竟，Controller的很多代码仅仅是做数据的管理操作而已。今天，我们就来重点学习Kafka集群元数据都有哪些。</p>

<p>如果说ZooKeeper是整个Kafka集群元数据的“真理之源（Source of Truth）”，那么Controller可以说是集群元数据的“真理之源副本（Backup Source of Truth）”。好吧，后面这个词是我自己发明的。你只需要理解，Controller承载了ZooKeeper上的所有元数据即可。</p>

<p>事实上，集群Broker是不会与ZooKeeper直接交互去获取元数据的。相反地，它们总是与Controller进行通信，获取和更新最新的集群数据。而且社区已经打算把ZooKeeper“干掉”了（我会在之后的“特别放送”里具体给你解释社区干掉ZooKeeper的操作），以后Controller将成为新的“真理之源”。</p>

<p>我们总说元数据，那么，到底什么是集群的元数据，或者说，Kafka集群的元数据都定义了哪些内容呢？我用一张图给你完整地展示一下，当前Kafka定义的所有集群元数据信息。</p>

<p><img src="assets/f146aceb78a5da31d887618303b5ff54.jpg" alt="" /></p>

<p>可以看到，目前，Controller定义的元数据有17项之多。不过，并非所有的元数据都同等重要，你也不用完整地记住它们，我们只需要重点关注那些最重要的元数据，并结合源代码来了解下这些元数据都是用来做什么的。</p>

<p>在了解具体的元数据之前，我要先介绍下ControllerContext类。刚刚我们提到的这些元数据信息全部封装在这个类里。应该这么说，<strong>这个类是Controller组件的数据容器类</strong>。</p>

<h2 id="controllercontext">ControllerContext</h2>

<p>Controller组件的源代码位于core包的src/main/scala/kafka/controller路径下，这里面有很多Scala源文件，<strong>ControllerContext类就位于这个路径下的ControllerContext.scala文件中。</strong></p>

<p>该文件只有几百行代码，其中，最重要的数据结构就是ControllerContext类。前面说过，<strong>它定义了前面提到的所有元数据信息，以及许多实用的工具方法</strong>。比如，获取集群上所有主题分区对象的allPartitions方法、获取某主题分区副本列表的partitionReplicaAssignment方法，等等。</p>

<p>首先，我们来看下ControllerContext类的定义，如下所示：</p>

<pre><code>class ControllerContext {
  val stats = new ControllerStats // Controller统计信息类 
  var offlinePartitionCount = 0   // 离线分区计数器
  val shuttingDownBrokerIds = mutable.Set.empty[Int]  // 关闭中Broker的Id列表
  private val liveBrokers = mutable.Set.empty[Broker] // 当前运行中Broker对象列表
  private val liveBrokerEpochs = mutable.Map.empty[Int, Long] 	// 运行中Broker Epoch列表
  var epoch: Int = KafkaController.InitialControllerEpoch   // Controller当前Epoch值
  var epochZkVersion: Int = KafkaController.InitialControllerEpochZkVersion	// Controller对应ZooKeeper节点的Epoch值
  val allTopics = mutable.Set.empty[String]	// 集群主题列表
  val partitionAssignments = mutable.Map.empty[String, mutable.Map[Int, ReplicaAssignment]]	// 主题分区的副本列表
  val partitionLeadershipInfo = mutable.Map.empty[TopicPartition, LeaderIsrAndControllerEpoch]	// 主题分区的Leader/ISR副本信息
  val partitionsBeingReassigned = mutable.Set.empty[TopicPartition]	// 正处于副本重分配过程的主题分区列表
  val partitionStates = mutable.Map.empty[TopicPartition, PartitionState] // 主题分区状态列表 
  val replicaStates = mutable.Map.empty[PartitionAndReplica, ReplicaState]	// 主题分区的副本状态列表
  val replicasOnOfflineDirs = mutable.Map.empty[Int, Set[TopicPartition]]	// 不可用磁盘路径上的副本列表
  val topicsToBeDeleted = mutable.Set.empty[String]	// 待删除主题列表
  val topicsWithDeletionStarted = mutable.Set.empty[String]	// 已开启删除的主题列表
  val topicsIneligibleForDeletion = mutable.Set.empty[String]	// 暂时无法执行删除的主题列表
  ......
}
</code></pre>

<p>不多不少，这段代码中定义的字段正好17个，它们一一对应着上图中的那些元数据信息。下面，我选取一些重要的元数据，来详细解释下它们的含义。</p>

<p>这些元数据理解起来还是比较简单的，掌握了它们之后，你在理解MetadataCache，也就是元数据缓存的时候，就容易得多了。比如，接下来我要讲到的liveBrokers信息，就是Controller通过UpdateMetadataRequest请求同步给其他Broker的MetadataCache的。</p>

<h3 id="controllerstats">ControllerStats</h3>

<p>第一个是ControllerStats类的变量。它的完整代码如下：</p>

<pre><code>private[controller] class ControllerStats extends KafkaMetricsGroup {
  // 统计每秒发生的Unclean Leader选举次数
  val uncleanLeaderElectionRate = newMeter(&quot;UncleanLeaderElectionsPerSec&quot;, &quot;elections&quot;, TimeUnit.SECONDS)
  // Controller事件通用的统计速率指标的方法
  val rateAndTimeMetrics: Map[ControllerState, KafkaTimer] = ControllerState.values.flatMap { state =&gt;
    state.rateAndTimeMetricName.map { metricName =&gt;
      state -&gt; new KafkaTimer(newTimer(metricName, TimeUnit.MILLISECONDS, TimeUnit.SECONDS))
    }
  }.toMap
}
</code></pre>

<p>顾名思义，它表征的是Controller的一些统计信息。目前，源码中定义了两大类统计指标：<strong>UncleanLeaderElectionsPerSec和所有Controller事件状态的执行速率与时间</strong>。</p>

<p>其中，<strong>前者是计算Controller每秒执行的Unclean Leader选举数量，通常情况下，执行Unclean Leader选举可能造成数据丢失，一般不建议开启它</strong>。一旦开启，你就需要时刻关注这个监控指标的值，确保Unclean Leader选举的速率维持在一个很低的水平，否则会出现很多数据丢失的情况。</p>

<p><strong>后者是统计所有Controller状态的速率和时间信息</strong>，单位是毫秒。当前，Controller定义了很多事件，比如，TopicDeletion是执行主题删除的Controller事件、ControllerChange是执行Controller重选举的事件。ControllerStats的这个指标通过在每个事件名后拼接字符串RateAndTimeMs的方式，为每类Controller事件都创建了对应的速率监控指标。</p>

<p>由于Controller事件有很多种，对应的速率监控指标也有很多，有一些Controller事件是需要你额外关注的。</p>

<p>举个例子，IsrChangeNotification事件是标志ISR列表变更的事件，如果这个事件经常出现，说明副本的ISR列表经常发生变化，而这通常被认为是非正常情况，因此，你最好关注下这个事件的速率监控指标。</p>

<h3 id="offlinepartitioncount">offlinePartitionCount</h3>

<p><strong>该字段统计集群中所有离线或处于不可用状态的主题分区数量</strong>。所谓的不可用状态，就是我最开始举的例子中“Leader=-1”的情况。</p>

<p>ControllerContext中的updatePartitionStateMetrics方法根据<strong>给定主题分区的当前状态和目标状态</strong>，来判断该分区是否是离线状态的分区。如果是，则累加offlinePartitionCount字段的值，否则递减该值。方法代码如下：</p>

<pre><code>// 更新offlinePartitionCount元数据
private def updatePartitionStateMetrics(
  partition: TopicPartition, 
  currentState: PartitionState,
  targetState: PartitionState): Unit = {
  // 如果该主题当前并未处于删除中状态
  if (!isTopicDeletionInProgress(partition.topic)) {
    // targetState表示该分区要变更到的状态
    // 如果当前状态不是OfflinePartition，即离线状态并且目标状态是离线状态
    // 这个if语句判断是否要将该主题分区状态转换到离线状态
    if (currentState != OfflinePartition &amp;&amp; targetState == OfflinePartition) {
      offlinePartitionCount = offlinePartitionCount + 1
    // 如果当前状态已经是离线状态，但targetState不是
    // 这个else if语句判断是否要将该主题分区状态转换到非离线状态
    } else if (currentState == OfflinePartition &amp;&amp; targetState != OfflinePartition) {
      offlinePartitionCount = offlinePartitionCount - 1
    }
  }
}
</code></pre>

<p>该方法首先要判断，此分区所属的主题当前是否处于删除操作的过程中。如果是的话，Kafka就不能修改这个分区的状态，那么代码什么都不做，直接返回。否则，代码会判断该分区是否要转换到离线状态。如果targetState是OfflinePartition，那么就将offlinePartitionCount值加1，毕竟多了一个离线状态的分区。相反地，如果currentState是offlinePartition，而targetState反而不是，那么就将offlinePartitionCount值减1。</p>

<h3 id="shuttingdownbrokerids">shuttingDownBrokerIds</h3>

<p>顾名思义，<strong>该字段保存所有正在关闭中的Broker ID列表</strong>。当Controller在管理集群Broker时，它要依靠这个字段来甄别Broker当前是否已关闭，因为处于关闭状态的Broker是不适合执行某些操作的，如分区重分配（Reassignment）以及主题删除等。</p>

<p>另外，Kafka必须要为这些关闭中的Broker执行很多清扫工作，Controller定义了一个onBrokerFailure方法，它就是用来做这个的。代码如下：</p>

<pre><code>private def onBrokerFailure(deadBrokers: Seq[Int]): Unit = {
  info(s&quot;Broker failure callback for ${deadBrokers.mkString(&quot;,&quot;)}&quot;)
  // deadBrokers：给定的一组已终止运行的Broker Id列表
  // 更新Controller元数据信息，将给定Broker从元数据的replicasOnOfflineDirs中移除
  deadBrokers.foreach(controllerContext.replicasOnOfflineDirs.remove)
  // 找出这些Broker上的所有副本对象
  val deadBrokersThatWereShuttingDown =
    deadBrokers.filter(id =&gt; controllerContext.shuttingDownBrokerIds.remove(id))
  if (deadBrokersThatWereShuttingDown.nonEmpty)
    info(s&quot;Removed ${deadBrokersThatWereShuttingDown.mkString(&quot;,&quot;)} from list of shutting down brokers.&quot;)
  // 执行副本清扫工作
  val allReplicasOnDeadBrokers = controllerContext.replicasOnBrokers(deadBrokers.toSet)
  onReplicasBecomeOffline(allReplicasOnDeadBrokers)
  // 取消这些Broker上注册的ZooKeeper监听器
  unregisterBrokerModificationsHandler(deadBrokers)
}
</code></pre>

<p>该方法接收一组已终止运行的Broker ID列表，首先是更新Controller元数据信息，将给定Broker从元数据的replicasOnOfflineDirs和shuttingDownBrokerIds中移除，然后为这组Broker执行必要的副本清扫工作，也就是onReplicasBecomeOffline方法做的事情。</p>

<p>该方法主要依赖于分区状态机和副本状态机来完成对应的工作。在后面的课程中，我们会专门讨论副本状态机和分区状态机，这里你只要简单了解下它要做的事情就行了。后面等我们学完了这两个状态机之后，你可以再看下这个方法的具体实现原理。</p>

<p>这个方法的主要目的是把给定的副本标记成Offline状态，即不可用状态。具体分为以下这几个步骤：</p>

<ol>
<li>利用分区状态机将给定副本所在的分区标记为Offline状态；</li>
<li>将集群上所有新分区和Offline分区状态变更为Online状态；</li>
<li>将相应的副本对象状态变更为Offline。</li>
</ol>

<h3 id="livebrokers">liveBrokers</h3>

<p><strong>该字段保存当前所有运行中的Broker对象</strong>。每个Broker对象就是一个<Id，EndPoint，机架信息>的三元组。ControllerContext中定义了很多方法来管理该字段，如addLiveBrokersAndEpochs、removeLiveBrokers和updateBrokerMetadata等。我拿updateBrokerMetadata方法进行说明，以下是源码：</p>

<pre><code>def updateBrokerMetadata(oldMetadata: Broker, newMetadata: Broker): Unit = {
    liveBrokers -= oldMetadata
    liveBrokers += newMetadata
  }
</code></pre>

<p>每当新增或移除已有Broker时，ZooKeeper就会更新其保存的Broker数据，从而引发Controller修改元数据，也就是会调用updateBrokerMetadata方法来增减Broker列表中的对象。怎么样，超简单吧？！</p>

<h3 id="livebrokerepochs">liveBrokerEpochs</h3>

<p><strong>该字段保存所有运行中Broker的Epoch信息</strong>。Kafka使用Epoch数据防止Zombie Broker，即一个非常老的Broker被选举成为Controller。</p>

<p>另外，源码大多使用这个字段来获取所有运行中Broker的ID序号，如下面这个方法定义的那样：</p>

<pre><code>def liveBrokerIds: Set[Int] = liveBrokerEpochs.keySet -- shuttingDownBrokerIds
</code></pre>

<p>liveBrokerEpochs的keySet方法返回Broker序号列表，然后从中移除关闭中的Broker序号，剩下的自然就是处于运行中的Broker序号列表了。</p>

<h3 id="epoch-epochzkversion">epoch &amp; epochZkVersion</h3>

<p>这两个字段一起说，因为它们都有“epoch”字眼，放在一起说，可以帮助你更好地理解两者的区别。epoch实际上就是ZooKeeper中/controller_epoch节点的值，你可以认为它就是Controller在整个Kafka集群的版本号，而epochZkVersion实际上是/controller_epoch节点的dataVersion值。</p>

<p>Kafka使用epochZkVersion来判断和防止Zombie Controller。这也就是说，原先在老Controller任期内的Controller操作在新Controller不能成功执行，因为新Controller的epochZkVersion要比老Controller的大。</p>

<p>另外，你可能会问：“这里的两个Epoch和上面的liveBrokerEpochs有啥区别呢？”实际上，这里的两个Epoch值都是属于Controller侧的数据，而liveBrokerEpochs是每个Broker自己的Epoch值。</p>

<h3 id="alltopics">allTopics</h3>

<p><strong>该字段保存集群上所有的主题名称</strong>。每当有主题的增减，Controller就要更新该字段的值。</p>

<p>比如Controller有个processTopicChange方法，从名字上来看，它就是处理主题变更的。我们来看下它的代码实现，我把主要逻辑以注释的方式标注了出来：</p>

<pre><code>private def processTopicChange(): Unit = {
    if (!isActive) return // 如果Contorller已经关闭，直接返回
    val topics = zkClient.getAllTopicsInCluster(true) // 从ZooKeeper中获取当前所有主题列表
    val newTopics = topics -- controllerContext.allTopics // 找出当前元数据中不存在、ZooKeeper中存在的主题，视为新增主题
    val deletedTopics = controllerContext.allTopics -- topics // 找出当前元数据中存在、ZooKeeper中不存在的主题，视为已删除主题
    controllerContext.allTopics = topics // 更新Controller元数据
    // 为新增主题和已删除主题执行后续处理操作
    registerPartitionModificationsHandlers(newTopics.toSeq)
    val addedPartitionReplicaAssignment = zkClient.getFullReplicaAssignmentForTopics(newTopics)
    deletedTopics.foreach(controllerContext.removeTopic)
    addedPartitionReplicaAssignment.foreach {
      case (topicAndPartition, newReplicaAssignment) =&gt; controllerContext.updatePartitionFullReplicaAssignment(topicAndPartition, newReplicaAssignment)
    }
    info(s&quot;New topics: [$newTopics], deleted topics: [$deletedTopics], new partition replica assignment &quot; +
      s&quot;[$addedPartitionReplicaAssignment]&quot;)
    if (addedPartitionReplicaAssignment.nonEmpty)
      onNewPartitionCreation(addedPartitionReplicaAssignment.keySet)
  }
</code></pre>

<h3 id="partitionassignments">partitionAssignments</h3>

<p><strong>该字段保存所有主题分区的副本分配情况</strong>。在我看来，<strong>这是Controller最重要的元数据了</strong>。事实上，你可以从这个字段衍生、定义很多实用的方法，来帮助Kafka从各种维度获取数据。</p>

<p>比如，如果Kafka要获取某个Broker上的所有分区，那么，它可以这样定义：</p>

<pre><code>partitionAssignments.flatMap {
      case (topic, topicReplicaAssignment) =&gt; topicReplicaAssignment.filter {
        case (_, partitionAssignment) =&gt; partitionAssignment.replicas.contains(brokerId)
      }.map {
        case (partition, _) =&gt; new TopicPartition(topic, partition)
      }
    }.toSet
</code></pre>

<p>再比如，如果Kafka要获取某个主题的所有分区对象，代码可以这样写：</p>

<pre><code>partitionAssignments.getOrElse(topic, mutable.Map.empty).map {
      case (partition, _) =&gt; new TopicPartition(topic, partition)
    }.toSet
</code></pre>

<p>实际上，这两段代码分别是ControllerContext.scala中partitionsOnBroker方法和partitionsForTopic两个方法的主体实现代码。</p>

<p>讲到这里，9个重要的元数据字段我就介绍完了。前面说过，ControllerContext中一共定义了17个元数据字段，你可以结合这9个字段，把其余8个的定义也过一遍，做到心中有数。<strong>你对Controller元数据掌握得越好，就越能清晰地理解Controller在集群中发挥的作用</strong>。</p>

<p>值得注意的是，在学习每个元数据字段时，除了它的定义之外，我建议你去搜索一下，与之相关的工具方法都是如何实现的。如果后面你想要新增获取或更新元数据的方法，你要对操作它们的代码有很强的把控力才行。</p>

<h2 id="总结">总结</h2>

<p>今天，我们揭开了Kafka重要组件Controller的学习大幕。我给出了Controller模块的学习路线，还介绍了Controller的重要元数据。</p>

<ul>
<li>Controller元数据：Controller当前定义了17种元数据，涵盖Kafka集群数据的方方面面。</li>
<li>ControllerContext：定义元数据以及操作它们的类。</li>
<li>关键元数据字段：最重要的元数据包括offlinePartitionCount、liveBrokers、partitionAssignments等。</li>
<li>ControllerContext工具方法：ControllerContext 类定义了很多实用方法来管理这些元数据信息。</li>
</ul>

<p>下节课，我们将学习Controller是如何给Broker发送请求的。Controller与Broker进行交互与通信，是Controller奠定王者地位的重要一环，我会向你详细解释它是如何做到这一点的。</p>

<h2 id="课后讨论">课后讨论</h2>

<p>我今天并未给出所有的元数据说明，请你自行结合代码分析一下，partitionLeadershipInfo里面保存的是什么数据？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#39555555000d0808090e795e54585055175a5654" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fb39fbc547791',t:'MTczNDAyNzM0NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
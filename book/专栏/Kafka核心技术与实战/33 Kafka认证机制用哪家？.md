<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=33&#32;Kafka认证机制用哪家？>
        <link rel="icon" href="/static/favicon.png">
        <title>33 Kafka认证机制用哪家？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么要学习Kafka？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6%e4%b9%a0Kafka%ef%bc%9f.md">00 开篇词 为什么要学习Kafka？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  消息引擎系统ABC.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%20%e6%b6%88%e6%81%af%e5%bc%95%e6%93%8e%e7%b3%bb%e7%bb%9fABC.md">01  消息引擎系统ABC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 一篇文章带你快速搞定Kafka术语.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/02%20%e4%b8%80%e7%af%87%e6%96%87%e7%ab%a0%e5%b8%a6%e4%bd%a0%e5%bf%ab%e9%80%9f%e6%90%9e%e5%ae%9aKafka%e6%9c%af%e8%af%ad.md">02 一篇文章带你快速搞定Kafka术语.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 Kafka只是消息引擎系统吗？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/03%20Kafka%e5%8f%aa%e6%98%af%e6%b6%88%e6%81%af%e5%bc%95%e6%93%8e%e7%b3%bb%e7%bb%9f%e5%90%97%ef%bc%9f.md">03 Kafka只是消息引擎系统吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 我应该选择哪种Kafka？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/04%20%e6%88%91%e5%ba%94%e8%af%a5%e9%80%89%e6%8b%a9%e5%93%aa%e7%a7%8dKafka%ef%bc%9f.md">04 我应该选择哪种Kafka？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 聊聊Kafka的版本号.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%e8%81%8a%e8%81%8aKafka%e7%9a%84%e7%89%88%e6%9c%ac%e5%8f%b7.md">05 聊聊Kafka的版本号.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Kafka线上集群部署方案怎么做？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/06%20Kafka%e7%ba%bf%e4%b8%8a%e9%9b%86%e7%be%a4%e9%83%a8%e7%bd%b2%e6%96%b9%e6%a1%88%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f.md">06 Kafka线上集群部署方案怎么做？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 最最最重要的集群参数配置（上）.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%e6%9c%80%e6%9c%80%e6%9c%80%e9%87%8d%e8%a6%81%e7%9a%84%e9%9b%86%e7%be%a4%e5%8f%82%e6%95%b0%e9%85%8d%e7%bd%ae%ef%bc%88%e4%b8%8a%ef%bc%89.md">07 最最最重要的集群参数配置（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 最最最重要的集群参数配置（下）.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%e6%9c%80%e6%9c%80%e6%9c%80%e9%87%8d%e8%a6%81%e7%9a%84%e9%9b%86%e7%be%a4%e5%8f%82%e6%95%b0%e9%85%8d%e7%bd%ae%ef%bc%88%e4%b8%8b%ef%bc%89.md">08 最最最重要的集群参数配置（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  生产者消息分区机制原理剖析.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%20%e7%94%9f%e4%ba%a7%e8%80%85%e6%b6%88%e6%81%af%e5%88%86%e5%8c%ba%e6%9c%ba%e5%88%b6%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90.md">09  生产者消息分区机制原理剖析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 生产者压缩算法面面观.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%e7%94%9f%e4%ba%a7%e8%80%85%e5%8e%8b%e7%bc%a9%e7%ae%97%e6%b3%95%e9%9d%a2%e9%9d%a2%e8%a7%82.md">10 生产者压缩算法面面观.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 无消息丢失配置怎么实现？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%e6%97%a0%e6%b6%88%e6%81%af%e4%b8%a2%e5%a4%b1%e9%85%8d%e7%bd%ae%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">11 无消息丢失配置怎么实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 客户端都有哪些不常见但是很高级的功能？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%e5%ae%a2%e6%88%b7%e7%ab%af%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e4%b8%8d%e5%b8%b8%e8%a7%81%e4%bd%86%e6%98%af%e5%be%88%e9%ab%98%e7%ba%a7%e7%9a%84%e5%8a%9f%e8%83%bd%ef%bc%9f.md">12 客户端都有哪些不常见但是很高级的功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  Java生产者是如何管理TCP连接的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%20Java%e7%94%9f%e4%ba%a7%e8%80%85%e6%98%af%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86TCP%e8%bf%9e%e6%8e%a5%e7%9a%84%ef%bc%9f.md">13  Java生产者是如何管理TCP连接的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 幂等生产者和事务生产者是一回事吗？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%e5%b9%82%e7%ad%89%e7%94%9f%e4%ba%a7%e8%80%85%e5%92%8c%e4%ba%8b%e5%8a%a1%e7%94%9f%e4%ba%a7%e8%80%85%e6%98%af%e4%b8%80%e5%9b%9e%e4%ba%8b%e5%90%97%ef%bc%9f.md">14 幂等生产者和事务生产者是一回事吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 消费者组到底是什么？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/15%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">15 消费者组到底是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 揭开神秘的“位移主题”面纱.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%e6%8f%ad%e5%bc%80%e7%a5%9e%e7%a7%98%e7%9a%84%e2%80%9c%e4%bd%8d%e7%a7%bb%e4%b8%bb%e9%a2%98%e2%80%9d%e9%9d%a2%e7%ba%b1.md">16 揭开神秘的“位移主题”面纱.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 消费者组重平衡能避免吗？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e9%87%8d%e5%b9%b3%e8%a1%a1%e8%83%bd%e9%81%bf%e5%85%8d%e5%90%97%ef%bc%9f.md">17 消费者组重平衡能避免吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Kafka中位移提交那些事儿.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/18%20Kafka%e4%b8%ad%e4%bd%8d%e7%a7%bb%e6%8f%90%e4%ba%a4%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">18 Kafka中位移提交那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 CommitFailedException异常怎么处理？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/19%20CommitFailedException%e5%bc%82%e5%b8%b8%e6%80%8e%e4%b9%88%e5%a4%84%e7%90%86%ef%bc%9f.md">19 CommitFailedException异常怎么处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 多线程开发消费者实例.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%e5%a4%9a%e7%ba%bf%e7%a8%8b%e5%bc%80%e5%8f%91%e6%b6%88%e8%b4%b9%e8%80%85%e5%ae%9e%e4%be%8b.md">20 多线程开发消费者实例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Java 消费者是如何管理TCP连接的_.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/21%20Java%20%e6%b6%88%e8%b4%b9%e8%80%85%e6%98%af%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86TCP%e8%bf%9e%e6%8e%a5%e7%9a%84_.md">21 Java 消费者是如何管理TCP连接的_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 消费者组消费进度监控都怎么实现？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e6%b6%88%e8%b4%b9%e8%bf%9b%e5%ba%a6%e7%9b%91%e6%8e%a7%e9%83%bd%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">22 消费者组消费进度监控都怎么实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Kafka副本机制详解.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/23%20Kafka%e5%89%af%e6%9c%ac%e6%9c%ba%e5%88%b6%e8%af%a6%e8%a7%a3.md">23 Kafka副本机制详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 请求是怎么被处理的？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%e8%af%b7%e6%b1%82%e6%98%af%e6%80%8e%e4%b9%88%e8%a2%ab%e5%a4%84%e7%90%86%e7%9a%84%ef%bc%9f.md">24 请求是怎么被处理的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 消费者组重平衡全流程解析.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e9%87%8d%e5%b9%b3%e8%a1%a1%e5%85%a8%e6%b5%81%e7%a8%8b%e8%a7%a3%e6%9e%90.md">25 消费者组重平衡全流程解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 你一定不能错过的Kafka控制器.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%e4%bd%a0%e4%b8%80%e5%ae%9a%e4%b8%8d%e8%83%bd%e9%94%99%e8%bf%87%e7%9a%84Kafka%e6%8e%a7%e5%88%b6%e5%99%a8.md">26 你一定不能错过的Kafka控制器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 关于高水位和Leader Epoch的讨论.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%e5%85%b3%e4%ba%8e%e9%ab%98%e6%b0%b4%e4%bd%8d%e5%92%8cLeader%20Epoch%e7%9a%84%e8%ae%a8%e8%ae%ba.md">27 关于高水位和Leader Epoch的讨论.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 主题管理知多少_.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%e4%b8%bb%e9%a2%98%e7%ae%a1%e7%90%86%e7%9f%a5%e5%a4%9a%e5%b0%91_.md">28 主题管理知多少_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 Kafka动态配置了解下？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/29%20Kafka%e5%8a%a8%e6%80%81%e9%85%8d%e7%bd%ae%e4%ba%86%e8%a7%a3%e4%b8%8b%ef%bc%9f.md">29 Kafka动态配置了解下？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 怎么重设消费者组位移？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%e6%80%8e%e4%b9%88%e9%87%8d%e8%ae%be%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84%e4%bd%8d%e7%a7%bb%ef%bc%9f.md">30 怎么重设消费者组位移？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 常见工具脚本大汇总.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/31%20%e5%b8%b8%e8%a7%81%e5%b7%a5%e5%85%b7%e8%84%9a%e6%9c%ac%e5%a4%a7%e6%b1%87%e6%80%bb.md">31 常见工具脚本大汇总.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 KafkaAdminClient：Kafka的运维利器.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/32%20KafkaAdminClient%ef%bc%9aKafka%e7%9a%84%e8%bf%90%e7%bb%b4%e5%88%a9%e5%99%a8.md">32 KafkaAdminClient：Kafka的运维利器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 Kafka认证机制用哪家？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/33%20Kafka%e8%ae%a4%e8%af%81%e6%9c%ba%e5%88%b6%e7%94%a8%e5%93%aa%e5%ae%b6%ef%bc%9f.md">33 Kafka认证机制用哪家？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 云环境下的授权该怎么做？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/34%20%e4%ba%91%e7%8e%af%e5%a2%83%e4%b8%8b%e7%9a%84%e6%8e%88%e6%9d%83%e8%af%a5%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f.md">34 云环境下的授权该怎么做？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 跨集群备份解决方案MirrorMaker.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/35%20%e8%b7%a8%e9%9b%86%e7%be%a4%e5%a4%87%e4%bb%bd%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88MirrorMaker.md">35 跨集群备份解决方案MirrorMaker.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 你应该怎么监控Kafka？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/36%20%e4%bd%a0%e5%ba%94%e8%af%a5%e6%80%8e%e4%b9%88%e7%9b%91%e6%8e%a7Kafka%ef%bc%9f.md">36 你应该怎么监控Kafka？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 主流的Kafka监控框架.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/37%20%e4%b8%bb%e6%b5%81%e7%9a%84Kafka%e7%9b%91%e6%8e%a7%e6%a1%86%e6%9e%b6.md">37 主流的Kafka监控框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 调优Kafka，你做到了吗？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/38%20%e8%b0%83%e4%bc%98Kafka%ef%bc%8c%e4%bd%a0%e5%81%9a%e5%88%b0%e4%ba%86%e5%90%97%ef%bc%9f.md">38 调优Kafka，你做到了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 从0搭建基于Kafka的企业级实时日志流处理平台.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/39%20%e4%bb%8e0%e6%90%ad%e5%bb%ba%e5%9f%ba%e4%ba%8eKafka%e7%9a%84%e4%bc%81%e4%b8%9a%e7%ba%a7%e5%ae%9e%e6%97%b6%e6%97%a5%e5%bf%97%e6%b5%81%e5%a4%84%e7%90%86%e5%b9%b3%e5%8f%b0.md">39 从0搭建基于Kafka的企业级实时日志流处理平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 Kafka Streams与其他流处理平台的差异在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/40%20Kafka%20Streams%e4%b8%8e%e5%85%b6%e4%bb%96%e6%b5%81%e5%a4%84%e7%90%86%e5%b9%b3%e5%8f%b0%e7%9a%84%e5%b7%ae%e5%bc%82%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">40 Kafka Streams与其他流处理平台的差异在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 Kafka Streams DSL开发实例.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/41%20Kafka%20Streams%20DSL%e5%bc%80%e5%8f%91%e5%ae%9e%e4%be%8b.md">41 Kafka Streams DSL开发实例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 Kafka Streams在金融领域的应用.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/42%20Kafka%20Streams%e5%9c%a8%e9%87%91%e8%9e%8d%e9%a2%86%e5%9f%9f%e7%9a%84%e5%ba%94%e7%94%a8.md">42 Kafka Streams在金融领域的应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 搭建开发环境、阅读源码方法、经典学习资料大揭秘.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%20%e6%90%ad%e5%bb%ba%e5%bc%80%e5%8f%91%e7%8e%af%e5%a2%83%e3%80%81%e9%98%85%e8%af%bb%e6%ba%90%e7%a0%81%e6%96%b9%e6%b3%95%e3%80%81%e7%bb%8f%e5%85%b8%e5%ad%a6%e4%b9%a0%e8%b5%84%e6%96%99%e5%a4%a7%e6%8f%ad%e7%a7%98.md">加餐 搭建开发环境、阅读源码方法、经典学习资料大揭秘.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 黄云：行百里者半九十.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e9%bb%84%e4%ba%91%ef%bc%9a%e8%a1%8c%e7%99%be%e9%87%8c%e8%80%85%e5%8d%8a%e4%b9%9d%e5%8d%81.md">用户故事 黄云：行百里者半九十.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 以梦为马，莫负韶华！.md" href="/%e4%b8%93%e6%a0%8f/Kafka%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e4%bb%a5%e6%a2%a6%e4%b8%ba%e9%a9%ac%ef%bc%8c%e8%8e%ab%e8%b4%9f%e9%9f%b6%e5%8d%8e%ef%bc%81.md">结束语 以梦为马，莫负韶华！.md</a>
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
                            <h1 id="title" data-id="33 Kafka认证机制用哪家？" class="title">33 Kafka认证机制用哪家？</h1>
                            <div><p>你好，我是胡夕。今天我要和你分享的主题是：Kafka的认证机制。</p>

<h2 id="什么是认证机制">什么是认证机制？</h2>

<p>所谓认证，又称“验证”“鉴权”，英文是authentication，是指通过一定的手段，完成对用户身份的确认。认证的主要目的是确认当前声称为某种身份的用户确实是所声称的用户。</p>

<p>在计算机领域，经常和认证搞混的一个术语就是授权，英文是authorization。授权一般是指对信息安全或计算机安全相关的资源定义与授予相应的访问权限。</p>

<p>举个简单的例子来区分下两者：认证要解决的是你要证明你是谁的问题，授权要解决的则是你能做什么的问题。</p>

<p>在Kafka中，认证和授权是两套独立的安全配置。我们今天主要讨论Kafka的认证机制，在专栏的下一讲内容中，我们将讨论授权机制。</p>

<h2 id="kafka认证机制">Kafka认证机制</h2>

<p>自0.9.0.0版本开始，Kafka正式引入了认证机制，用于实现基础的安全用户认证，这是将Kafka上云或进行多租户管理的必要步骤。截止到当前最新的2.3版本，Kafka支持基于SSL和基于SASL的安全认证机制。</p>

<p><strong>基于SSL的认证主要是指Broker和客户端的双路认证</strong>（2-way authentication）。通常来说，SSL加密（Encryption）已经启用了单向认证，即客户端认证Broker的证书（Certificate）。如果要做SSL认证，那么我们要启用双路认证，也就是说Broker也要认证客户端的证书。</p>

<p>对了，你可能会说，SSL不是已经过时了吗？现在都叫TLS（Transport Layer Security）了吧？但是，Kafka的源码中依然是使用SSL而不是TLS来表示这类东西的。不过，今天出现的所有SSL字眼，你都可以认为它们是和TLS等价的。</p>

<p>Kafka还支持通过SASL做客户端认证。<strong>SASL是提供认证和数据安全服务的框架</strong>。Kafka支持的SASL机制有5种，它们分别是在不同版本中被引入的，你需要根据你自己使用的Kafka版本，来选择该版本所支持的认证机制。</p>

<ol>
<li>GSSAPI：也就是Kerberos使用的安全接口，是在0.9版本中被引入的。</li>
<li>PLAIN：是使用简单的用户名/密码认证的机制，在0.10版本中被引入。</li>
<li>SCRAM：主要用于解决PLAIN机制安全问题的新机制，是在0.10.2版本中被引入的。</li>
<li>OAUTHBEARER：是基于OAuth 2认证框架的新机制，在2.0版本中被引进。</li>
<li>Delegation Token：补充现有SASL机制的轻量级认证机制，是在1.1.0版本被引入的。</li>
</ol>

<h2 id="认证机制的比较">认证机制的比较</h2>

<p>Kafka为我们提供了这么多种认证机制，在实际使用过程中，我们应该如何选择合适的认证框架呢？下面我们就来比较一下。</p>

<p>目前来看，使用SSL做信道加密的情况更多一些，但使用SSL实现认证不如使用SASL。毕竟，SASL能够支持你选择不同的实现机制，如GSSAPI、SCRAM、PLAIN等。因此，我的建议是<strong>你可以使用SSL来做通信加密，使用SASL来做Kafka的认证实现</strong>。</p>

<p>SASL下又细分了很多种认证机制，我们应该如何选择呢？</p>

<p>SASL/GSSAPI主要是给Kerberos使用的。如果你的公司已经做了Kerberos认证（比如使用Active Directory），那么使用GSSAPI是最方便的了。因为你不需要额外地搭建Kerberos，只要让你们的Kerberos管理员给每个Broker和要访问Kafka集群的操作系统用户申请principal就好了。总之，<strong>GSSAPI适用于本身已经做了Kerberos认证的场景，这样的话，SASL/GSSAPI可以实现无缝集成</strong>。</p>

<p>而SASL/PLAIN，就像前面说到的，它是一个简单的用户名/密码认证机制，通常与SSL加密搭配使用。注意，这里的PLAIN和PLAINTEXT是两回事。<strong>PLAIN在这里是一种认证机制，而PLAINTEXT说的是未使用SSL时的明文传输</strong>。对于一些小公司而言，搭建公司级的Kerberos可能并没有什么必要，他们的用户系统也不复杂，特别是访问Kafka集群的用户可能不是很多。对于SASL/PLAIN而言，这就是一个非常合适的应用场景。<strong>总体来说，SASL/PLAIN的配置和运维成本相对较小，适合于小型公司中的Kafka集群</strong>。</p>

<p>但是，SASL/PLAIN有这样一个弊端：它不能动态地增减认证用户，你必须重启Kafka集群才能令变更生效。为什么呢？这是因为所有认证用户信息全部保存在静态文件中，所以只能重启Broker，才能重新加载变更后的静态文件。</p>

<p>我们知道，重启集群在很多场景下都是令人不爽的，即使是轮替式升级（Rolling Upgrade）。SASL/SCRAM就解决了这样的问题。它通过将认证用户信息保存在ZooKeeper的方式，避免了动态修改需要重启Broker的弊端。在实际使用过程中，你可以使用Kafka提供的命令动态地创建和删除用户，无需重启整个集群。因此，<strong>如果你打算使用SASL/PLAIN，不妨改用SASL/SCRAM试试。不过要注意的是，后者是0.10.2版本引入的。你至少要升级到这个版本后才能使用</strong>。</p>

<p>SASL/OAUTHBEARER是2.0版本引入的新认证机制，主要是为了实现与OAuth 2框架的集成。OAuth是一个开发标准，允许用户授权第三方应用访问该用户在某网站上的资源，而无需将用户名和密码提供给第三方应用。Kafka不提倡单纯使用OAUTHBEARER，因为它生成的不安全的JSON Web Token，必须配以SSL加密才能用在生产环境中。当然，鉴于它是2.0版本才推出来的，而且目前没有太多的实际使用案例，我们可以先观望一段时间，再酌情将其应用于生产环境中。</p>

<p>Delegation Token是在1.1版本引入的，它是一种轻量级的认证机制，主要目的是补充现有的SASL或SSL认证。如果要使用Delegation Token，你需要先配置好SASL认证，然后再利用Kafka提供的API去获取对应的Delegation Token。这样，Broker和客户端在做认证的时候，可以直接使用这个token，不用每次都去KDC获取对应的ticket（Kerberos认证）或传输Keystore文件（SSL认证）。</p>

<p>为了方便你更好地理解和记忆，我把这些认证机制汇总在下面的表格里了。你可以对照着表格，进行一下区分。</p>

<p><img src="assets/4a52c2eb1ae631697b5ec3d298f7333d.jpg" alt="" /></p>

<h2 id="sasl-scram-sha-256配置实例">SASL/SCRAM-SHA-256配置实例</h2>

<p>接下来，我给出SASL/SCRAM的一个配置实例，来说明一下如何在Kafka集群中开启认证。其他认证机制的设置方法也是类似的，比如它们都涉及认证用户的创建、Broker端以及Client端特定参数的配置等。</p>

<p>我的测试环境是本地Mac上的两个Broker组成的Kafka集群，连接端口分别是9092和9093。</p>

<h3 id="第1步-创建用户">第1步：创建用户</h3>

<p>配置SASL/SCRAM的第一步，是创建能否连接Kafka集群的用户。在本次测试中，我会创建3个用户，分别是admin用户、writer用户和reader用户。admin用户用于实现Broker间通信，writer用户用于生产消息，reader用户用于消费消息。</p>

<p>我们使用下面这3条命令，分别来创建它们。</p>

<pre><code>$ cd kafka_2.12-2.3.0/
$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'SCRAM-SHA-256=[password=admin],SCRAM-SHA-512=[password=admin]' --entity-type users --entity-name admin
Completed Updating config for entity: user-principal 'admin'.


$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'SCRAM-SHA-256=[password=writer],SCRAM-SHA-512=[password=writer]' --entity-type users --entity-name writer
Completed Updating config for entity: user-principal 'writer'.


$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'SCRAM-SHA-256=[password=reader],SCRAM-SHA-512=[password=reader]' --entity-type users --entity-name reader
Completed Updating config for entity: user-principal 'reader'.
</code></pre>

<p>在专栏前面，我们提到过，kafka-configs脚本是用来设置主题级别参数的。其实，它的功能还有很多。比如在这个例子中，我们使用它来创建SASL/SCRAM认证中的用户信息。我们可以使用下列命令来查看刚才创建的用户数据。</p>

<pre><code>$ bin/kafka-configs.sh --zookeeper localhost:2181 --describe --entity-type users  --entity-name writer
Configs for user-principal 'writer' are SCRAM-SHA-512=salt=MWt6OGplZHF6YnF5bmEyam9jamRwdWlqZWQ=,stored_key=hR7+vgeCEz61OmnMezsqKQkJwMCAoTTxw2jftYiXCHxDfaaQU7+9/dYBq8bFuTio832mTHk89B4Yh9frj/ampw==,server_key=C0k6J+9/InYRohogXb3HOlG7s84EXAs/iw0jGOnnQAt4jxQODRzeGxNm+18HZFyPn7qF9JmAqgtcU7hgA74zfA==,iterations=4096,SCRAM-SHA-256=salt=MWV0cDFtbXY5Nm5icWloajdnbjljZ3JqeGs=,stored_key=sKjmeZe4sXTAnUTL1CQC7DkMtC+mqKtRY0heEHvRyPk=,server_key=kW7CC3PBj+JRGtCOtIbAMefL8aiL8ZrUgF5tfomsWVA=,iterations=4096
</code></pre>

<p>这段命令包含了writer用户加密算法SCRAM-SHA-256以及SCRAM-SHA-512对应的盐值(Salt)、ServerKey和StoreKey。这些都是SCRAM机制的术语，我们不需要了解它们的含义，因为它们并不影响我们接下来的配置。</p>

<h3 id="第2步-创建jaas文件">第2步：创建JAAS文件</h3>

<p>配置了用户之后，我们需要为每个Broker创建一个对应的JAAS文件。因为本例中的两个Broker实例是在一台机器上，所以我只创建了一份JAAS文件。但是你要切记，在实际场景中，你需要为每台单独的物理Broker机器都创建一份JAAS文件。</p>

<p>JAAS的文件内容如下：</p>

<pre><code>KafkaServer {
org.apache.kafka.common.security.scram.ScramLoginModule required
username=&quot;admin&quot;
password=&quot;admin&quot;;
};
</code></pre>

<p>关于这个文件内容，你需要注意以下两点：</p>

<ul>
<li>不要忘记最后一行和倒数第二行结尾处的分号；</li>
<li>JAAS文件中不需要任何空格键。</li>
</ul>

<p>这里，我们使用admin用户实现Broker之间的通信。接下来，我们来配置Broker的server.properties文件，下面这些内容，是需要单独配置的：</p>

<pre><code>sasl.enabled.mechanisms=SCRAM-SHA-256


sasl.mechanism.inter.broker.protocol=SCRAM-SHA-256


security.inter.broker.protocol=SASL_PLAINTEXT


listeners=SASL_PLAINTEXT://localhost:9092
</code></pre>

<p>第1项内容表明开启SCRAM认证机制，并启用SHA-256算法；第2项的意思是为Broker间通信也开启SCRAM认证，同样使用SHA-256算法；第3项表示Broker间通信不配置SSL，本例中我们不演示SSL的配置；最后1项是设置listeners使用SASL_PLAINTEXT，依然是不使用SSL。</p>

<p>另一台Broker的配置基本和它类似，只是要使用不同的端口，在这个例子中，端口是9093。</p>

<h3 id="第3步-启动broker">第3步：启动Broker</h3>

<p>现在我们分别启动这两个Broker。在启动时，你需要指定JAAS文件的位置，如下所示：</p>

<pre><code>$KAFKA_OPTS=-Djava.security.auth.login.config=&lt;your_path&gt;/kafka-broker.jaas bin/kafka-server-start.sh config/server1.properties
......
[2019-07-02 13:30:34,822] INFO Kafka commitId: fc1aaa116b661c8a (org.apache.kafka.common.utils.AppInfoParser)
[2019-07-02 13:30:34,822] INFO Kafka startTimeMs: 1562045434820 (org.apache.kafka.common.utils.AppInfoParser)
[2019-07-02 13:30:34,823] INFO [KafkaServer id=0] started (kafka.server.KafkaServer)


$KAFKA_OPTS=-Djava.security.auth.login.config=&lt;your_path&gt;/kafka-broker.jaas bin/kafka-server-start.sh config/server2.properties
......
[2019-07-02 13:32:31,976] INFO Kafka commitId: fc1aaa116b661c8a (org.apache.kafka.common.utils.AppInfoParser)
[2019-07-02 13:32:31,976] INFO Kafka startTimeMs: 1562045551973 (org.apache.kafka.common.utils.AppInfoParser)
[2019-07-02 13:32:31,978] INFO [KafkaServer id=1] started (kafka.server.KafkaServer)
</code></pre>

<p>此时，两台Broker都已经成功启动了。</p>

<h3 id="第4步-发送消息">第4步：发送消息</h3>

<p>在创建好测试主题之后，我们使用kafka-console-producer脚本来尝试发送消息。由于启用了认证，客户端需要做一些相应的配置。我们创建一个名为producer.conf的配置文件，内容如下：</p>

<pre><code>security.protocol=SASL_PLAINTEXT
sasl.mechanism=SCRAM-SHA-256
sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required username=&quot;writer&quot; password=&quot;writer&quot;;
</code></pre>

<p>之后运行Console Producer程序：</p>

<pre><code>$ bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093 --topic test  --producer.config &lt;your_path&gt;/producer.conf
&gt;hello, world
&gt;   
</code></pre>

<p>可以看到，Console Producer程序发送消息成功。</p>

<h3 id="第5步-消费消息">第5步：消费消息</h3>

<p>接下来，我们使用Console Consumer程序来消费一下刚刚生产的消息。同样地，我们需要为kafka-console-consumer脚本创建一个名为consumer.conf的脚本，内容如下：</p>

<pre><code>security.protocol=SASL_PLAINTEXT
sasl.mechanism=SCRAM-SHA-256
sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required username=&quot;reader&quot; password=&quot;reader&quot;;
</code></pre>

<p>之后运行Console Consumer程序：</p>

<pre><code>$ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092,localhost:9093 --topic test --from-beginning --consumer.config &lt;your_path&gt;/consumer.conf 
hello, world
</code></pre>

<p>很显然，我们是可以正常消费的。</p>

<h3 id="第6步-动态增减用户">第6步：动态增减用户</h3>

<p>最后，我们来演示SASL/SCRAM动态增减用户的场景。假设我删除了writer用户，同时又添加了一个新用户：new_writer，那么，我们需要执行的命令如下：</p>

<pre><code>$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --delete-config 'SCRAM-SHA-256' --entity-type users --entity-name writer
Completed Updating config for entity: user-principal 'writer'.

$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --delete-config 'SCRAM-SHA-512' --entity-type users --entity-name writer
Completed Updating config for entity: user-principal 'writer'.

$ bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'SCRAM-SHA-256=[iterations=8192,password=new_writer]' --entity-type users --entity-name new_writer
Completed Updating config for entity: user-principal 'new_writer'.
</code></pre>

<p>现在，我们依然使用刚才的producer.conf来验证，以确认Console Producer程序不能发送消息。</p>

<pre><code>$ bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093 --topic test  --producer.config /Users/huxi/testenv/producer.conf
&gt;[2019-07-02 13:54:29,695] ERROR [Producer clientId=console-producer] Connection to node -1 (localhost/127.0.0.1:9092) failed authentication due to: Authentication failed during authentication due to invalid credentials with SASL mechanism SCRAM-SHA-256 (org.apache.kafka.clients.NetworkClient)
......
</code></pre>

<p>很显然，此时Console Producer已经不能发送消息了。因为它使用的producer.conf文件指定的是已经被删除的writer用户。如果我们修改producer.conf的内容，改为指定新创建的new_writer用户，结果如下：</p>

<pre><code>$ bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093 --topic test  --producer.config &lt;your_path&gt;/producer.conf
&gt;Good!  
</code></pre>

<p>现在，Console Producer可以正常发送消息了。</p>

<p>这个过程完整地展示了SASL/SCRAM是如何在不重启Broker的情况下增减用户的。</p>

<p>至此，SASL/SCRAM配置就完成了。在专栏下一讲中，我会详细介绍一下如何赋予writer和reader用户不同的权限。</p>

<h2 id="小结">小结</h2>

<p>好了，我们来小结一下。今天，我们讨论了Kafka目前提供的几种认证机制，我给出了它们各自的优劣势以及推荐使用建议。其实，在真实的使用场景中，认证和授权往往是结合在一起使用的。在专栏下一讲中，我会详细向你介绍Kafka的授权机制，即ACL机制，敬请期待。</p>

<p><img src="assets/af705cb98a2f46acd45b184ec201005c.jpg" alt="" /></p>

<h2 id="开放讨论">开放讨论</h2>

<p>请谈一谈你的Kafka集群上的用户认证机制，并分享一个你遇到过的“坑”。</p>

<p>欢迎写下你的思考和答案，我们一起讨论。如果你觉得有所收获，也欢迎把文章分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ddb1b1b1e4e9ececedea9dbab0bcb4b1f3beb2b0" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fad7dabdb7791',t:'MTczNDAyNzA5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
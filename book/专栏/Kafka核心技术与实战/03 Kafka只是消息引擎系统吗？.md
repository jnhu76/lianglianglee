<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=03&#32;Kafka只是消息引擎系统吗？>
        <link rel="icon" href="/static/favicon.png">
        <title>03 Kafka只是消息引擎系统吗？ </title>
        
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
                            <h1 id="title" data-id="03 Kafka只是消息引擎系统吗？" class="title">03 Kafka只是消息引擎系统吗？</h1>
                            <div><p>你好，我是胡夕。今天我们来聊一个老生常谈的话题：Kafka只是消息引擎系统吗？</p>

<p>要搞清楚这个问题，我们不可避免地要了解一下Apache Kafka的发展历程。有的时候我们会觉得说了解一个系统或框架的前世今生似乎没什么必要，直接开始学具体的技术不是更快更好吗？其实，不论是学习哪种技术，直接扎到具体的细节中，亦或是从一个很小的点开始学习，你很快就会感到厌烦。为什么呢？因为你虽然快速地搞定了某个技术细节，但无法建立全局的认知观，这会导致你只是在单个的点上有所进展，却没法将其串联成一条线进而扩展成一个面，从而实现系统地学习。</p>

<p>我这么说是有依据的，因为这就是我当初学习Kafka的方式。你可能不会相信，我阅读Kafka源码就是从utils包开始的。显然，我们不用看源码也知道这玩意是干什么用的，对吧？就是个工具类包嘛，而且这种阅读源码的方式是极其低效的。就像我说的，我是在一个点一个点地学习，但全部学完之后压根没有任何感觉，依然不了解Kafka，因为不知道这些包中的代码组合在一起能达成什么效果。所以我说它是很低效的学习方法。</p>

<p>后来我修改了学习的方法，转而从自上而下的角度去理解Kafka，竟然发现了很多之前学习过程中忽略掉的东西。更特别的是，我发现这种学习方法能够帮助我维持较长时间的学习兴趣，不会阶段性地产生厌烦情绪。特别是在了解Apache Kafka整个发展历史的过程中我愉快地学到了很多运营大型开源软件社区的知识和经验，可谓是技术之外的一大收获。</p>

<p>纵观Kafka的发展脉络，它的确是从消息引擎起家的，但正如文章标题所问，<strong>Apache Kafka真的只是消息引擎吗</strong>？通常，在回答这个问题之前很多文章可能就要这样展开了：那我们先来讨论下什么是消息引擎以及消息引擎能做什么事情。算了，我还是直给吧，就不从“唐尧虞舜”说起了。这个问题的答案是，<strong>Apache Kafka是消息引擎系统，也是一个分布式流处理平台</strong>（Distributed Streaming Platform）。如果你通读全篇文字但只能记住一句话，我希望你记住的就是这句。再强调一遍，Kafka是消息引擎系统，也是分布式流处理平台。</p>

<p>众所周知，Kafka是LinkedIn公司内部孵化的项目。根据我和Kafka创始团队成员的交流以及查阅到的公开信息显示，LinkedIn最开始有强烈的数据强实时处理方面的需求，其内部的诸多子系统要执行多种类型的数据处理与分析，主要包括业务系统和应用程序性能监控，以及用户行为数据处理等。</p>

<p>当时他们碰到的主要问题包括：</p>

<ul>
<li>数据正确性不足。因为数据的收集主要采用轮询（Polling）的方式，如何确定轮询的间隔时间就变成了一个高度经验化的事情。虽然可以采用一些类似于启发式算法（Heuristic）来帮助评估间隔时间值，但一旦指定不当，必然会造成较大的数据偏差。</li>
<li>系统高度定制化，维护成本高。各个业务子系统都需要对接数据收集模块，引入了大量的定制开销和人工成本。</li>
</ul>

<p>为了解决这些问题，LinkedIn工程师尝试过使用ActiveMQ来解决这些问题，但效果并不理想。显然需要有一个“大一统”的系统来取代现有的工作方式，而这个系统就是Kafka。</p>

<p>Kafka自诞生伊始是以消息引擎系统的面目出现在大众视野中的。如果翻看0.10.0.0之前的官网说明，你会发现Kafka社区将其清晰地定位为一个分布式、分区化且带备份功能的提交日志（Commit Log）服务。</p>

<p>这里引出一个题外话，你可能好奇Kafka这个名字的由来，实际上Kafka作者之一Jay Kreps曾经谈及过命名的原因。</p>

<blockquote>
<p>因为Kafka系统的写性能很强，所以找了个作家的名字来命名似乎是一个好主意。大学期间我上了很多文学课，非常喜欢Franz Kafka这个作家，另外为开源软件起这个名字听上去很酷。</p>
</blockquote>

<p>言归正传，Kafka在设计之初就旨在提供三个方面的特性：</p>

<ul>
<li>提供一套API实现生产者和消费者；</li>
<li>降低网络传输和磁盘存储开销；</li>
<li>实现高伸缩性架构。</li>
</ul>

<p>在专栏后面的课程中，我们将陆续探讨Kafka是如何做到以上三点的。总之随着Kafka的不断完善，Jay等大神们终于意识到将其开源惠及更多的人是一个非常棒的主意，因此在2011年Kafka正式进入到Apache基金会孵化并于次年10月顺利毕业成为Apache顶级项目。</p>

<p>开源之后的Kafka被越来越多的公司应用到它们企业内部的数据管道中，特别是在大数据工程领域，Kafka在承接上下游、串联数据流管道方面发挥了重要的作用：所有的数据几乎都要从一个系统流入Kafka然后再流向下游的另一个系统中。这样的使用方式屡见不鲜以至于引发了Kafka社区的思考：与其我把数据从一个系统传递到下一个系统中做处理，我为何不自己实现一套流处理框架呢？基于这个考量，Kafka社区于0.10.0.0版本正式推出了流处理组件Kafka Streams，也正是从这个版本开始，Kafka正式“变身”为分布式的流处理平台，而不仅仅是消息引擎系统了。今天Apache Kafka是和Apache Storm、Apache Spark和Apache Flink同等级的实时流处理平台。</p>

<p>诚然，目前国内对Kafka是流处理平台的认知还尚不普及，其核心的流处理组件Kafka Streams更是少有大厂在使用。但我们也欣喜地看到，随着在Kafka峰会上各路大神们的鼎力宣传，如今利用Kafka构建流处理平台的案例层出不穷，而了解并有意愿使用Kafka Streams的厂商也是越来越多，因此我个人对于Kafka流处理平台的前景也是非常乐观的。</p>

<p>你可能会有这样的疑问：作为流处理平台，Kafka与其他主流大数据流式计算框架相比，优势在哪里呢？我能想到的有两点。</p>

<p><strong>第一点是更容易实现端到端的正确性（Correctness）</strong>。Google大神Tyler曾经说过，流处理要最终替代它的“兄弟”批处理需要具备两点核心优势：<strong>要实现正确性和提供能够推导时间的工具。实现正确性是流处理能够匹敌批处理的基石</strong>。正确性一直是批处理的强项，而实现正确性的基石则是要求框架能提供精确一次处理语义，即处理一条消息有且只有一次机会能够影响系统状态。目前主流的大数据流处理框架都宣称实现了精确一次处理语义，但这是有限定条件的，即它们只能实现框架内的精确一次处理语义，无法实现端到端的。</p>

<p>这是为什么呢？因为当这些框架与外部消息引擎系统结合使用时，它们无法影响到外部系统的处理语义，所以如果你搭建了一套环境使得Spark或Flink从Kafka读取消息之后进行有状态的数据计算，最后再写回Kafka，那么你只能保证在Spark或Flink内部，这条消息对于状态的影响只有一次。但是计算结果有可能多次写入到Kafka，因为它们不能控制Kafka的语义处理。相反地，Kafka则不是这样，因为所有的数据流转和计算都在Kafka内部完成，故Kafka可以实现端到端的精确一次处理语义。</p>

<p><strong>可能助力Kafka胜出的第二点是它自己对于流式计算的定位</strong>。官网上明确标识Kafka Streams是一个用于搭建实时流处理的客户端库而非是一个完整的功能系统。这就是说，你不能期望着Kafka提供类似于集群调度、弹性部署等开箱即用的运维特性，你需要自己选择适合的工具或系统来帮助Kafka流处理应用实现这些功能。</p>

<p>读到这你可能会说这怎么算是优点呢？坦率来说，这的确是一个“双刃剑”的设计，也是Kafka社区“剑走偏锋”不正面PK其他流计算框架的特意考量。大型公司的流处理平台一定是大规模部署的，因此具备集群调度功能以及灵活的部署方案是不可或缺的要素。但毕竟这世界上还存在着很多中小企业，它们的流处理数据量并不巨大，逻辑也并不复杂，部署几台或十几台机器足以应付。在这样的需求之下，搭建重量级的完整性平台实在是“杀鸡焉用牛刀”，而这正是Kafka流处理组件的用武之地。因此从这个角度来说，未来在流处理框架中，Kafka应该是有一席之地的。</p>

<p>除了消息引擎和流处理平台，Kafka还有别的用途吗？当然有！你能想象吗，Kafka能够被用作分布式存储系统。Kafka作者之一Jay Kreps曾经专门写过一篇文章阐述为什么能把<a href="https://www.confluent.io/blog/okay-store-data-apache-kafka/" target="_blank">Kafka用作分布式存储</a>。不过我觉得你姑且了解下就好了，我从没有见过在实际生产环境中，有人把Kafka当作持久化存储来用 。</p>

<p>说了这么多，我只想阐述这样的一个观点：Apache Kafka从一个优秀的消息引擎系统起家，逐渐演变成现在分布式的流处理平台。你不仅要熟练掌握它作为消息引擎系统的非凡特性及使用技巧，最好还要多了解下其流处理组件的设计与案例应用。</p>

<p><img src="assets/eafc2d9315a8ba858649c1ecd4201459.jpg" alt="" /></p>

<h2 id="开放讨论">开放讨论</h2>

<p>你觉得Kafka未来的演进路线是怎么样的？如果你是Kafka社区的“掌舵人”，你准备带领整个社区奔向什么方向呢？（提示下，你可以把自己想象成Linus再去思考）</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c6aaaaaafff2f7f7f6f186a1aba7afaae8a5a9ab" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fa6c7f8a47791',t:'MTczNDAyNjgyMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=19&#32;CommitFailedException异常怎么处理？>
        <link rel="icon" href="/static/favicon.png">
        <title>19 CommitFailedException异常怎么处理？ </title>
        
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
                            <h1 id="title" data-id="19 CommitFailedException异常怎么处理？" class="title">19 CommitFailedException异常怎么处理？</h1>
                            <div><p>你好，我是胡夕。今天我来跟你聊聊CommitFailedException异常的处理。</p>

<p>说起这个异常，我相信用过Kafka Java Consumer客户端API的你一定不会感到陌生。<strong>所谓CommitFailedException，顾名思义就是Consumer客户端在提交位移时出现了错误或异常，而且还是那种不可恢复的严重异常</strong>。如果异常是可恢复的瞬时错误，提交位移的API自己就能规避它们了，因为很多提交位移的API方法是支持自动错误重试的，比如我们在上一期中提到的<strong>commitSync方法</strong>。</p>

<p>每次和CommitFailedException一起出现的，还有一段非常著名的注释。为什么说它很“著名”呢？第一，我想不出在近50万行的Kafka源代码中，还有哪个异常类能有这种待遇，可以享有这么大段的注释，来阐述其异常的含义；第二，纵然有这么长的文字解释，却依然有很多人对该异常想表达的含义感到困惑。</p>

<p>现在，我们一起领略下这段文字的风采，看看社区对这个异常的最新解释：</p>

<blockquote>
<p>Commit cannot be completed since the group has already rebalanced and assigned the partitions to another member. This means that the time between subsequent calls to poll() was longer than the configured max.poll.interval.ms, which typically implies that the poll loop is spending too much time message processing. You can address this either by increasing max.poll.interval.ms or by reducing the maximum size of batches returned in poll() with max.poll.records.</p>
</blockquote>

<p>这段话前半部分的意思是，本次提交位移失败了，原因是消费者组已经开启了Rebalance过程，并且将要提交位移的分区分配给了另一个消费者实例。出现这个情况的原因是，你的消费者实例连续两次调用poll方法的时间间隔超过了期望的max.poll.interval.ms参数值。这通常表明，你的消费者实例花费了太长的时间进行消息处理，耽误了调用poll方法。</p>

<p>在后半部分，社区给出了两个相应的解决办法（即橙色字部分）：</p>

<ol>
<li>增加期望的时间间隔max.poll.interval.ms参数值。</li>
<li>减少poll方法一次性返回的消息数量，即减少max.poll.records参数值。</li>
</ol>

<p>在详细讨论这段文字之前，我还想提一句，实际上这段文字总共有3个版本，除了上面的这个最新版本，还有2个版本，它们分别是：</p>

<blockquote>
<p>Commit cannot be completed since the group has already rebalanced and assigned the partitions to another member. This means that the time between subsequent calls to poll() was longer than the configured session.timeout.ms, which typically implies that the poll loop is spending too much time message processing. You can address this either by increasing the session timeout or by reducing the maximum size of batches returned in poll() with max.poll.records.</p>

<p>Commit cannot be completed since the group has already rebalanced and assigned the partitions to another member. This means that the time between subsequent calls to poll() was longer than the configured max.poll.interval.ms, which typically implies that the poll loop is spending too much time message processing. You can address this either by increasing the session timeout or by reducing the maximum size of batches returned in poll() with max.poll.records.</p>
</blockquote>

<p>这两个较早的版本和最新版相差不大，我就不详细解释了，具体的差异我用橙色标注了。我之所以列出这些版本，就是想让你在日后看到它们时能做到心中有数，知道它们说的是一个事情。</p>

<p>其实不论是哪段文字，它们都表征位移提交出现了异常。下面我们就来讨论下该异常是什么时候被抛出的。从源代码方面来说，CommitFailedException异常通常发生在手动提交位移时，即用户显式调用KafkaConsumer.commitSync()方法时。从使用场景来说，有两种典型的场景可能遭遇该异常。</p>

<p><strong>场景一</strong></p>

<p>我们先说说最常见的场景。当消息处理的总时间超过预设的max.poll.interval.ms参数值时，Kafka Consumer端会抛出CommitFailedException异常。这是该异常最“正宗”的登场方式。你只需要写一个Consumer程序，使用KafkaConsumer.subscribe方法随意订阅一个主题，之后设置Consumer端参数max.poll.interval.ms=5秒，最后在循环调用KafkaConsumer.poll方法之间，插入Thread.sleep(6000)和手动提交位移，就可以成功复现这个异常了。在这里，我展示一下主要的代码逻辑。</p>

<pre><code>…
Properties props = new Properties();
…
props.put(&quot;max.poll.interval.ms&quot;, 5000);
consumer.subscribe(Arrays.asList(&quot;test-topic&quot;));

while (true) {
    ConsumerRecords&lt;String, String&gt; records = 
		consumer.poll(Duration.ofSeconds(1));
    // 使用Thread.sleep模拟真实的消息处理逻辑
    Thread.sleep(6000L);
    consumer.commitSync();
}
</code></pre>

<p>如果要防止这种场景下抛出异常，你需要简化你的消息处理逻辑。具体来说有4种方法。</p>

<ol>
<li><p><strong>缩短单条消息处理的时间</strong>。比如，之前下游系统消费一条消息的时间是100毫秒，优化之后成功地下降到50毫秒，那么此时Consumer端的TPS就提升了一倍。</p></li>

<li><p><strong>增加Consumer端允许下游系统消费一批消息的最大时长</strong>。这取决于Consumer端参数max.poll.interval.ms的值。在最新版的Kafka中，该参数的默认值是5分钟。如果你的消费逻辑不能简化，那么提高该参数值是一个不错的办法。值得一提的是，Kafka 0.10.1.0之前的版本是没有这个参数的，因此如果你依然在使用0.10.1.0之前的客户端API，那么你需要增加session.timeout.ms参数的值。不幸的是，session.timeout.ms参数还有其他的含义，因此增加该参数的值可能会有其他方面的“不良影响”，这也是社区在0.10.1.0版本引入max.poll.interval.ms参数，将这部分含义从session.timeout.ms中剥离出来的原因之一。</p></li>

<li><p><strong>减少下游系统一次性消费的消息总数</strong>。这取决于Consumer端参数max.poll.records的值。当前该参数的默认值是500条，表明调用一次KafkaConsumer.poll方法，最多返回500条消息。可以说，该参数规定了单次poll方法能够返回的消息总数的上限。如果前两种方法对你都不适用的话，降低此参数值是避免CommitFailedException异常最简单的手段。</p></li>

<li><p><strong>下游系统使用多线程来加速消费</strong>。这应该算是“最高级”同时也是最难实现的解决办法了。具体的思路就是，让下游系统手动创建多个消费线程处理poll方法返回的一批消息。之前你使用Kafka Consumer消费数据更多是单线程的，所以当消费速度无法匹及Kafka Consumer消息返回的速度时，它就会抛出CommitFailedException异常。如果是多线程，你就可以灵活地控制线程数量，随时调整消费承载能力，再配以目前多核的硬件条件，该方法可谓是防止CommitFailedException最高档的解决之道。事实上，很多主流的大数据流处理框架使用的都是这个方法，比如Apache Flink在集成Kafka时，就是创建了多个KafkaConsumerThread线程，自行处理多线程间的数据消费。不过，凡事有利就有弊，这个方法实现起来并不容易，特别是在多个线程间如何处理位移提交这个问题上，更是极容易出错。在专栏后面的内容中，我将着重和你讨论一下多线程消费的实现方案。</p></li>
</ol>

<p>综合以上这4个处理方法，我个人推荐你首先尝试采用方法1来预防此异常的发生。优化下游系统的消费逻辑是百利而无一害的法子，不像方法2、3那样涉及到Kafka Consumer端TPS与消费延时（Latency）的权衡。如果方法1实现起来有难度，那么你可以按照下面的法则来实践方法2、3。</p>

<p>首先，你需要弄清楚你的下游系统消费每条消息的平均延时是多少。比如你的消费逻辑是从Kafka获取到消息后写入到下游的MongoDB中，假设访问MongoDB的平均延时不超过2秒，那么你可以认为消息处理需要花费2秒的时间。如果按照max.poll.records等于500来计算，一批消息的总消费时长大约是1000秒，因此你的Consumer端的max.poll.interval.ms参数值就不能低于1000秒。如果你使用默认配置，那默认值5分钟显然是不够的，你将有很大概率遭遇CommitFailedException异常。将max.poll.interval.ms增加到1000秒以上的做法就属于上面的第2种方法。</p>

<p>除了调整max.poll.interval.ms之外，你还可以选择调整max.poll.records值，减少每次poll方法返回的消息数。还拿刚才的例子来说，你可以设置max.poll.records值为150，甚至更少，这样每批消息的总消费时长不会超过300秒（150*2=300），即max.poll.interval.ms的默认值5分钟。这种减少max.poll.records值的做法就属于上面提到的方法3。</p>

<p><strong>场景二</strong></p>

<p>Okay，现在我们已经说完了关于CommitFailedException异常的经典发生场景以及应对办法。从理论上讲，关于该异常你了解到这个程度，已经足以帮助你应对应用开发过程中由该异常带来的“坑”了 。但其实，该异常还有一个不太为人所知的出现场景。了解这个冷门场景，可以帮助你拓宽Kafka Consumer的知识面，也能提前预防一些古怪的问题。下面我们就来说说这个场景。</p>

<p>之前我们花了很多时间学习Kafka的消费者，不过大都集中在消费者组上，即所谓的Consumer Group。其实，Kafka Java Consumer端还提供了一个名为Standalone Consumer的独立消费者。它没有消费者组的概念，每个消费者实例都是独立工作的，彼此之间毫无联系。不过，你需要注意的是，独立消费者的位移提交机制和消费者组是一样的，因此独立消费者的位移提交也必须遵守之前说的那些规定，比如独立消费者也要指定group.id参数才能提交位移。你可能会觉得奇怪，既然是独立消费者，为什么还要指定group.id呢？没办法，谁让社区就是这么设计的呢？总之，消费者组和独立消费者在使用之前都要指定group.id。</p>

<p>现在问题来了，如果你的应用中同时出现了设置相同group.id值的消费者组程序和独立消费者程序，那么当独立消费者程序手动提交位移时，Kafka就会立即抛出CommitFailedException异常，因为Kafka无法识别这个具有相同group.id的消费者实例，于是就向它返回一个错误，表明它不是消费者组内合法的成员。</p>

<p>虽然说这个场景很冷门，但也并非完全不会遇到。在一个大型公司中，特别是那些将Kafka作为全公司级消息引擎系统的公司中，每个部门或团队都可能有自己的消费者应用，谁能保证各自的Consumer程序配置的group.id没有重复呢？一旦出现不凑巧的重复，发生了上面提到的这种场景，你使用之前提到的哪种方法都不能规避该异常。令人沮丧的是，无论是刚才哪个版本的异常说明，都完全没有提及这个场景，因此，如果是这个原因引发的CommitFailedException异常，前面的4种方法全部都是无效的。</p>

<p>更为尴尬的是，无论是社区官网，还是网上的文章，都没有提到过这种使用场景。我个人认为，这应该算是Kafka的一个bug。比起返回CommitFailedException异常只是表明提交位移失败，更好的做法应该是，在Consumer端应用程序的某个地方，能够以日志或其他方式友善地提示你错误的原因，这样你才能正确处理甚至是预防该异常。</p>

<h2 id="小结">小结</h2>

<p>总结一下，今天我们详细讨论了Kafka Consumer端经常碰到的CommitFailedException异常。我们从它的含义说起，再到它出现的时机和场景，以及每种场景下的应对之道。当然，我也留了个悬念，在专栏后面的内容中，我会详细说说多线程消费的实现方式。希望通过今天的分享，你能清晰地掌握CommitFailedException异常发生的方方面面，从而能在今后更有效地应对此异常。</p>

<p><img src="assets/df3691cee68c7878efd21e79719bec88.jpg" alt="" /></p>

<h2 id="开放讨论">开放讨论</h2>

<p>请比较一下今天我们提到的预防该异常的4种方法，并说说你对它们的理解。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#0e626262373a3f3f3e394e69636f6762206d6163" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0fa9f4bc5e7791',t:'MTczNDAyNjk1MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
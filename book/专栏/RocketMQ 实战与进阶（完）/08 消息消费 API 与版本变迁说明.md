<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=08&#32;消息消费&#32;API&#32;与版本变迁说明>
        <link rel="icon" href="/static/favicon.png">
        <title>08 消息消费 API 与版本变迁说明 </title>
        
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
                        <a class="menu-item" id="01 搭建学习环境准备篇.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/01%20%e6%90%ad%e5%bb%ba%e5%ad%a6%e4%b9%a0%e7%8e%af%e5%a2%83%e5%87%86%e5%a4%87%e7%af%87.md">01 搭建学习环境准备篇.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 RocketMQ 核心概念扫盲篇.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/02%20RocketMQ%20%e6%a0%b8%e5%bf%83%e6%a6%82%e5%bf%b5%e6%89%ab%e7%9b%b2%e7%af%87.md">02 RocketMQ 核心概念扫盲篇.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 消息发送 API 详解与版本变迁说明.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/03%20%e6%b6%88%e6%81%af%e5%8f%91%e9%80%81%20API%20%e8%af%a6%e8%a7%a3%e4%b8%8e%e7%89%88%e6%9c%ac%e5%8f%98%e8%bf%81%e8%af%b4%e6%98%8e.md">03 消息发送 API 详解与版本变迁说明.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 结合实际应用场景谈消息发送.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/04%20%e7%bb%93%e5%90%88%e5%ae%9e%e9%99%85%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%e8%b0%88%e6%b6%88%e6%81%af%e5%8f%91%e9%80%81.md">04 结合实际应用场景谈消息发送.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 消息发送核心参数与工作原理详解.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/05%20%e6%b6%88%e6%81%af%e5%8f%91%e9%80%81%e6%a0%b8%e5%bf%83%e5%8f%82%e6%95%b0%e4%b8%8e%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86%e8%af%a6%e8%a7%a3.md">05 消息发送核心参数与工作原理详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 消息发送常见错误与解决方案.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/06%20%e6%b6%88%e6%81%af%e5%8f%91%e9%80%81%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%e4%b8%8e%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88.md">06 消息发送常见错误与解决方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 事务消息使用及方案选型思考.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/07%20%e4%ba%8b%e5%8a%a1%e6%b6%88%e6%81%af%e4%bd%bf%e7%94%a8%e5%8f%8a%e6%96%b9%e6%a1%88%e9%80%89%e5%9e%8b%e6%80%9d%e8%80%83.md">07 事务消息使用及方案选型思考.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 消息消费 API 与版本变迁说明.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/08%20%e6%b6%88%e6%81%af%e6%b6%88%e8%b4%b9%20API%20%e4%b8%8e%e7%89%88%e6%9c%ac%e5%8f%98%e8%bf%81%e8%af%b4%e6%98%8e.md">08 消息消费 API 与版本变迁说明.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 DefaultMQPushConsumer 核心参数与工作原理.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/09%20DefaultMQPushConsumer%20%e6%a0%b8%e5%bf%83%e5%8f%82%e6%95%b0%e4%b8%8e%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86.md">09 DefaultMQPushConsumer 核心参数与工作原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 DefaultMQPushConsumer 使用示例与注意事项.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/10%20DefaultMQPushConsumer%20%e4%bd%bf%e7%94%a8%e7%a4%ba%e4%be%8b%e4%b8%8e%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9.md">10 DefaultMQPushConsumer 使用示例与注意事项.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 DefaultLitePullConsumer 核心参数与实战.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/11%20DefaultLitePullConsumer%20%e6%a0%b8%e5%bf%83%e5%8f%82%e6%95%b0%e4%b8%8e%e5%ae%9e%e6%88%98.md">11 DefaultLitePullConsumer 核心参数与实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 结合实际场景再聊 DefaultLitePullConsumer 的使用.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/12%20%e7%bb%93%e5%90%88%e5%ae%9e%e9%99%85%e5%9c%ba%e6%99%af%e5%86%8d%e8%81%8a%20DefaultLitePullConsumer%20%e7%9a%84%e4%bd%bf%e7%94%a8.md">12 结合实际场景再聊 DefaultLitePullConsumer 的使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 结合实际场景顺序消费、消息过滤实战.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/13%20%e7%bb%93%e5%90%88%e5%ae%9e%e9%99%85%e5%9c%ba%e6%99%af%e9%a1%ba%e5%ba%8f%e6%b6%88%e8%b4%b9%e3%80%81%e6%b6%88%e6%81%af%e8%bf%87%e6%bb%a4%e5%ae%9e%e6%88%98.md">13 结合实际场景顺序消费、消息过滤实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 消息消费积压问题排查实战.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/14%20%e6%b6%88%e6%81%af%e6%b6%88%e8%b4%b9%e7%a7%af%e5%8e%8b%e9%97%ae%e9%a2%98%e6%8e%92%e6%9f%a5%e5%ae%9e%e6%88%98.md">14 消息消费积压问题排查实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 RocketMQ 常用命令实战.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/15%20RocketMQ%20%e5%b8%b8%e7%94%a8%e5%91%bd%e4%bb%a4%e5%ae%9e%e6%88%98.md">15 RocketMQ 常用命令实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 RocketMQ 集群性能摸高.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/16%20RocketMQ%20%e9%9b%86%e7%be%a4%e6%80%a7%e8%83%bd%e6%91%b8%e9%ab%98.md">16 RocketMQ 集群性能摸高.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 RocketMQ 集群性能调优.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/17%20RocketMQ%20%e9%9b%86%e7%be%a4%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98.md">17 RocketMQ 集群性能调优.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 RocketMQ 集群平滑运维.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/18%20RocketMQ%20%e9%9b%86%e7%be%a4%e5%b9%b3%e6%bb%91%e8%bf%90%e7%bb%b4.md">18 RocketMQ 集群平滑运维.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 RocketMQ 集群监控（一）.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/19%20RocketMQ%20%e9%9b%86%e7%be%a4%e7%9b%91%e6%8e%a7%ef%bc%88%e4%b8%80%ef%bc%89.md">19 RocketMQ 集群监控（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 RocketMQ 集群监控（二）.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/20%20RocketMQ%20%e9%9b%86%e7%be%a4%e7%9b%91%e6%8e%a7%ef%bc%88%e4%ba%8c%ef%bc%89.md">20 RocketMQ 集群监控（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 RocketMQ 集群告警.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/21%20RocketMQ%20%e9%9b%86%e7%be%a4%e5%91%8a%e8%ad%a6.md">21 RocketMQ 集群告警.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 RocketMQ 集群踩坑记.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/22%20RocketMQ%20%e9%9b%86%e7%be%a4%e8%b8%a9%e5%9d%91%e8%ae%b0.md">22 RocketMQ 集群踩坑记.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 消息轨迹、ACL 与多副本搭建.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/23%20%e6%b6%88%e6%81%af%e8%bd%a8%e8%bf%b9%e3%80%81ACL%20%e4%b8%8e%e5%a4%9a%e5%89%af%e6%9c%ac%e6%90%ad%e5%bb%ba.md">23 消息轨迹、ACL 与多副本搭建.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 RocketMQ-Console 常用页面指标获取逻辑.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/24%20RocketMQ-Console%20%e5%b8%b8%e7%94%a8%e9%a1%b5%e9%9d%a2%e6%8c%87%e6%a0%87%e8%8e%b7%e5%8f%96%e9%80%bb%e8%be%91.md">24 RocketMQ-Console 常用页面指标获取逻辑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 RocketMQ Nameserver 背后的设计理念.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/25%20RocketMQ%20Nameserver%20%e8%83%8c%e5%90%8e%e7%9a%84%e8%ae%be%e8%ae%a1%e7%90%86%e5%bf%b5.md">25 RocketMQ Nameserver 背后的设计理念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 Java 并发编程实战.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/26%20Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98.md">26 Java 并发编程实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 从 RocketMQ 学基于文件的编程模式（一）.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/27%20%e4%bb%8e%20RocketMQ%20%e5%ad%a6%e5%9f%ba%e4%ba%8e%e6%96%87%e4%bb%b6%e7%9a%84%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%b8%80%ef%bc%89.md">27 从 RocketMQ 学基于文件的编程模式（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 从 RocketMQ 学基于文件的编程模式（二）.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/28%20%e4%bb%8e%20RocketMQ%20%e5%ad%a6%e5%9f%ba%e4%ba%8e%e6%96%87%e4%bb%b6%e7%9a%84%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%ba%8c%ef%bc%89.md">28 从 RocketMQ 学基于文件的编程模式（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 从 RocketMQ 学 Netty 网络编程技巧.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/29%20%e4%bb%8e%20RocketMQ%20%e5%ad%a6%20Netty%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e6%8a%80%e5%b7%a7.md">29 从 RocketMQ 学 Netty 网络编程技巧.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 RocketMQ 学习方法之我见.md" href="/%e4%b8%93%e6%a0%8f/RocketMQ%20%e5%ae%9e%e6%88%98%e4%b8%8e%e8%bf%9b%e9%98%b6%ef%bc%88%e5%ae%8c%ef%bc%89/30%20RocketMQ%20%e5%ad%a6%e4%b9%a0%e6%96%b9%e6%b3%95%e4%b9%8b%e6%88%91%e8%a7%81.md">30 RocketMQ 学习方法之我见.md</a>
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
                            <h1 id="title" data-id="08 消息消费 API 与版本变迁说明" class="title">08 消息消费 API 与版本变迁说明</h1>
                            <div><p>从本篇开始我们将详细介绍 RockeMQ 的消息消费端的 API。</p>

<h3 id="消息消费类图">消息消费类图</h3>

<p>RocketMQ 消费端的 API 如下图所示：</p>

<p><img src="assets/20200811225030941.png" alt="1" /></p>

<p>其核心类图如下所示。</p>

<p><strong>MQAdmin</strong></p>

<p>MQ 一些基本的管理功能，例如创建 Topic，这里稍微有点奇怪，消费端应该不需要继承该接口。该类在消息发送 API 章节已详细介绍，再次不再重复说明。</p>

<p><strong>MQConsumer</strong></p>

<p>MQ 消费者，这个接口定义得过于简单，如果该接口需要，可以将其子接口一些共同的方法提取到该接口中。</p>

<pre><code>Set&lt;MessageQueue&gt; fetchSubscribeMessageQueues(final String topic)

</code></pre>

<p>获取分配该 Topic 所有的读队列。</p>

<p><strong>MQPushConsumer</strong></p>

<p>RocketMQ 支持推、拉两种模式，该接口是<strong>拉模式</strong>的接口定义。</p>

<pre><code>void start()

</code></pre>

<p>启动消费者。</p>

<pre><code>void shutdown()

</code></pre>

<p>关闭消费者。</p>

<pre><code>void registerMessageQueueListener(String topic, MessageQueueListener listener)

</code></pre>

<p>注册消息队列变更回调时间，即消费端分配到的队列发生变化时触发的回调函数，其声明如下：</p>

<p><img src="assets/2020081122504049.png" alt="2" /></p>

<p>其参数说明如下：</p>

<ul>
<li><code>String topic</code>：主题。</li>
<li><code>Set&lt;MessageQueue&gt; mqAll</code>：该 Topic 所有的队列集合。</li>
<li><code>Set&lt;MessageQueue&gt; mqDivided</code>：分配给当前消费者的消费队列。</li>
</ul>

<pre><code>PullResult pull(MessageQueue mq, String subExpression, long offset,int maxNums, long timeout)

</code></pre>

<p>消息拉取，应用程序可以通过调用该方法从 RocketMQ 服务器拉取一篇消息，其参数含义说明如下：</p>

<ul>
<li><code>MessageQueue mq</code>：消息消费队列。</li>
<li><code>String subExpression</code>：消息过滤表达式，基于 Tag、SQL92 的过滤表达式。</li>
<li><code>long offset</code>：消息偏移量，消息在 ConsumeQueue 中的偏移量。</li>
<li><code>int maxNums</code>：一次消息拉取返回的最大消息条数。</li>
<li><code>long timeout</code>：本次拉取的超时时间。</li>
</ul>

<pre><code>PullResult pull(MessageQueue mq, MessageSelector selector, long offset,int maxNums, long timeout)

</code></pre>

<p>pull 重载方法，通过 MessageSelector 构建消息过滤对象，可以通过 MessageSelector 的 buildSql、buildTag 两个方法构建过滤表达式。</p>

<pre><code>void pull(MessageQueue mq, String subExpression, long offset, int maxNums,PullCallback pullCallback)

</code></pre>

<p>异步拉取，调用其异步回调函数 PullCallback。</p>

<pre><code>PullResult pullBlockIfNotFound(MessageQueue mq, String subExpression,long offset, int maxNums)

</code></pre>

<p>拉取消息，如果服务端没有新消息待拉取，一直阻塞等待，直到有消息返回，同样该方法有一个重载放假支持异步拉取。</p>

<pre><code>void updateConsumeOffset(MessageQueue mq, long offset)

</code></pre>

<p>更新消息消费处理进度。</p>

<pre><code>long fetchConsumeOffset(MessageQueue mq, boolean fromStore)

</code></pre>

<p>获取指定消息消费队列的消费进度。其中参数 fromStore 如果为 true，表示从消息消费进度存储文件中获取消费进度。</p>

<pre><code>Set&lt;MessageQueue&gt; fetchMessageQueuesInBalance(String topic)

</code></pre>

<p>获取当前正在处理的消息消费队列（通过消息队列负载机制分配的队列）。</p>

<pre><code>void sendMessageBack(MessageExt msg, int delayLevel, String brokerName, String consumerGroup)

</code></pre>

<p>消息消费失败后发送的 ACK。</p>

<p><strong>MQPushConsumer</strong></p>

<p>RocketMQ <strong>推模式</strong>消费者接口。</p>

<pre><code>void start()

</code></pre>

<p>启动消费者。</p>

<pre><code>void shutdown()

</code></pre>

<p>关闭消费者。</p>

<pre><code>void registerMessageListener(MessageListenerConcurrently messageListener)

</code></pre>

<p>注册并发消费模式监听器。</p>

<pre><code>void registerMessageListener(MessageListenerOrderly messageListener)

</code></pre>

<p>注册顺序消费模式监听器。</p>

<pre><code>void subscribe(String topic, String subExpression)

</code></pre>

<p>订阅主题。其参数说明如下：</p>

<ul>
<li><code>String topic</code>： 订阅的主题，RocketMQ 支持一个消费者订阅多个主题，操作方式是多次调用该方法。</li>
<li><code>String subExpression</code>：消息过滤表达式，例如传入订阅的 tag，SQL92 表达式。</li>
</ul>

<pre><code>void subscribe(String topic, MessageSelector selector)

</code></pre>

<p>订阅主题，重载方法，MessageSelector 提供了 buildSQL、buildTag 的订阅方式。</p>

<pre><code>void unsubscribe(String topic)

</code></pre>

<p>取消订阅。</p>

<pre><code>void suspend()

</code></pre>

<p>挂起消费。</p>

<pre><code>void resume()

</code></pre>

<p>恢复继续消费。</p>

<p><strong>DefaultMQPushConsumer</strong></p>

<p>RocketMQ 消息推模式默认实现类。</p>

<p><strong>DefaultMQPullConsumer</strong></p>

<p>RocketMQ 消息拉取模式默认实现类。</p>

<p>在 RocketMQ 的内部实现原理中，其实现机制为 PULL 模式，而 PUSH 模式是一种伪推送，是对 PULL 模式的封装，PUSH 模式的实现原理如下图所示：</p>

<p><img src="assets/2020081122504049.png" alt="3" /></p>

<p>即 PUSH 模式就是对 PULL 模式的封装，每拉去一批消息后，提交到消费端的线程池（异步），然后马上向 Broker 拉取消息，即实现类似“推”的效果。</p>

<p>从 PULL 模式来看，消息的消费主要包含如下几个方面：</p>

<ul>
<li>消息拉取，消息拉取模式通过 PULL 相关的 API 从 Broker 指定消息消费队列中拉取一批消息到消费消费客户端，多个消费者需要手动完成队列的分配。</li>
<li>消息消费端处理完消费，需要向 Broker 端报告消息处理队列，然后继续拉取下一批消息。</li>
<li>如果遇到消息消费失败，需要告知 Broker，该条消息消费失败，后续需要重试，通过手动调用 sendMessageBack 方法实现。</li>
</ul>

<p>而 PUSH 模式就上述这些处理操作无需使用者考虑，只需告诉 RocketMQ 消费者在拉取消息后需要调用的事件监听器即可，消息消费进度的存储、消息消费的重试统一由 RocketMQ Client 来实现。</p>

<h3 id="消息消费-api-简单使用示例">消息消费 API 简单使用示例</h3>

<p>从上文基本可以得知，推模式 API 与拉模式 API 在使用层面的差别，可以简单理解为汽车领域的自动挡与手动挡。在实际业务类场景中，通常使用的是推送风格的 API，适合实时监控；但在大数据领域，通常是跑批处理即定时类任务，故大数据领域通常使用拉模式更多。</p>

<p>接下来编写几个示例代码对拉取、推送相关 API 进行一个使用方面的演示。</p>

<h4 id="rocketmq-拉模式核心-api-使用示例"><strong>RocketMQ 拉模式核心 API 使用示例</strong></h4>

<p>使用场景：例如公司大数据团队需要对订单进行分析，为了提高计算效能，采取每 2 个小时调度一次，每批任务处理任务启动之前的所有消息。</p>

<p>首先我先给出一个基于 RocketMQ PULL 的 API 的编程示例代码，本示例接近生产实践，示例代码如下：</p>

<pre><code>import org.apache.rocketmq.client.consumer.DefaultMQPullConsumer;
import org.apache.rocketmq.client.consumer.PullResult;
import org.apache.rocketmq.common.message.MessageExt;
import org.apache.rocketmq.common.message.MessageQueue;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
public class PullConsumerTest {
    public static void main(String[] args) throws Exception {
        Semaphore semaphore = new Semaphore();
        Thread t = new Thread(new Task(semaphore));
        t.start();
        CountDownLatch cdh = new CountDownLatch(1);
        try {
            //程序运行 120s　后介绍
            cdh.await(120 * 1000, TimeUnit.MILLISECONDS);
        } finally {
            semaphore.running = false;
        }
    }
    /**
     * 消息拉取核心实现逻辑
     */
    static class Task implements Runnable {
        Semaphore s = new Semaphore();
        public Task(Semaphore s ) {
            this.s = s;
        }
        public void run() {
            try {
                DefaultMQPullConsumer consumer = new 
                    DefaultMQPullConsumer(&quot;dw_pull_consumer&quot;);
                consumer.setNamesrvAddr(&quot;127.0.01:9876&quot;);
                consumer.start();
                Map&lt;MessageQueue, Long&gt; offsetTable = new HashMap&lt;MessageQueue, Long&gt;();
                Set&lt;MessageQueue&gt; msgQueueList = consumer.
                    fetchSubscribeMessageQueues(&quot;TOPIC_TEST&quot;); // 获取该 Topic 的所有队列
                if(msgQueueList != null &amp;&amp; !msgQueueList.isEmpty()) {
                    boolean noFoundFlag = false;
                    while(this.s.running) {
                        if(noFoundFlag) { //　没有找到消息，暂停一下消费
                            Thread.sleep(1000);
                        }
                        for( MessageQueue q : msgQueueList ) {
                            PullResult pullResult = consumer.pull(q, &quot;*&quot;,                                          decivedPulloffset(offsetTable
                             , q, consumer) , 3000);
                            System.out.println(&quot;pullStatus:&quot; + 
                                               pullResult.getPullStatus());
                            switch (pullResult.getPullStatus()) {
                                case FOUND:
                                    doSomething(pullResult.getMsgFoundList());
                                    break;
                                case NO_MATCHED_MSG:
                                    break;
                                case NO_NEW_MSG:
                                case OFFSET_ILLEGAL:
                                    noFoundFlag = true;
                                    break;
                                default:
                                    continue ;
                            }
                            //提交位点
                            consumer.updateConsumeOffset(q, 
                                 pullResult.getNextBeginOffset());
                        }
                        System.out.println(&quot;balacne queue is empty: &quot; + consumer.
                              fetchMessageQueuesInBalance(&quot;TOPIC_TEST&quot;).isEmpty());
                    }
                } else {
                    System.out.println(&quot;end,because queue is enmpty&quot;);
                }
                consumer.shutdown();
                System.out.println(&quot;consumer shutdown&quot;);
            } catch (Throwable e) {
                e.printStackTrace();
            }
        }
    }
    /** 拉取到消息后具体的处理逻辑　*/
    private static void doSomething(List&lt;MessageExt&gt; msgs) {
        System.out.println(&quot;本次拉取到的消息条数:&quot; + msgs.size());
    }
    public static long decivedPulloffset(Map&lt;MessageQueue, Long&gt; offsetTable, 
             MessageQueue queue, DefaultMQPullConsumer consumer) throws Exception {
        long offset = consumer.fetchConsumeOffset(queue, false);
        if(offset &lt; 0 ) {
            offset = 0;
        }
        System.out.println(&quot;offset:&quot; + offset);
        return offset;
    }
    static class Semaphore {
        public volatile boolean running = true;
    }
}

</code></pre>

<p>关于上述代码，提供了优雅的线程拉取的方法，消息的拉取实现主要在任务 Task 的 run 方法中，主要的实现技巧如下：</p>

<ul>
<li>首先根据 MQConsumer 的 fetchSubscribeMessageQueues 的方法获取 Topic 的所有队列信息。</li>
<li>然后遍历所有队列，依次通过 MQConsuemr 的 PULL 方法从 Broker 端拉取消息。</li>
<li>对拉取的消息进行消费处理。</li>
<li>通过调用 MQConsumer 的 updateConsumeOffset 方法更新位点，但需要注意的是这个方法并不是实时向 Broker 提交，而是客户端会启用以线程，默认每隔 5s 向 Broker 集中上报一次。</li>
</ul>

<p>上面的示例演示的是一个消费组只有一个消费者，如果有多个消费组呢？这里就涉及到队列的重新分配，而每一个消费者是否只负责拉取分配的队列，是不是觉得这个直接使用 PULL 模式，是不是觉得有点复杂了。笔者也觉得是，接下来我们来看一下 PUSH 模式。</p>

<h4 id="rocketmq-推模式使用示例"><strong>RocketMQ 推模式使用示例</strong></h4>

<p>在 RocketMQ 中绝大数场景中，通常会选择使用 PUSH 模式，因为 PUSH 模式是对 PULL 模式的封装，将消息的拉取、消息队列的自动负载、消息进度（位点）自动提交、消息消费重试都进行了封装，无需使用者关心，其示例代码如下：</p>

<pre><code>public static void main(String[] args) throws InterruptedException, MQClientException {
        DefaultMQPushConsumer consumer = new 
            DefaultMQPushConsumer(&quot;dw_test_consumer_6&quot;);
        consumer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        consumer.setConsumeFromWhere(ConsumeFromWhere.CONSUME_FROM_FIRST_OFFSET);
        consumer.subscribe(&quot;TOPIC_TEST&quot;, &quot;*&quot;);
        consumer.setAllocateMessageQueueStrategy(new 
               AllocateMessageQueueAveragelyByCircle());
        consumer.registerMessageListener(new MessageListenerConcurrently() {
            @Override
            public ConsumeConcurrentlyStatus consumeMessage(List&lt;MessageExt&gt; msgs,
                ConsumeConcurrentlyContext context) {
                try {
                    System.out.printf(&quot;%s Receive New Messages: %s %n&quot;, 
                          Thread.currentThread().getName(), msgs);
                    return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
                } catch (Throwable e) {
                    e.printStackTrace();
                    return ConsumeConcurrentlyStatus.RECONSUME_LATER;
                }
            }
        });
        consumer.start();
        System.out.printf(&quot;Consumer Started.%n&quot;);
    }

</code></pre>

<p>上面的代码是不是非常简单。在后续的文章我们会重点对 PUSH 模式消费者的核心属性即工作原理做详细的介绍。使用 DefaultMQPushConsumer 开发一个消费者，其代码基本覆盖如下几个方面：</p>

<ul>
<li>首先 new DefaultMQPushConsumer 对象，并指定一个消费组名。</li>
<li>然后设置相关参数，例如 nameSrvAdd、消费失败重试次数、线程数等（可以设置哪些参数将在下篇文章中详细介绍）。</li>
<li>通过调用 setConsumeFromWhere 方法指定初次启动时从什么地方消费，默认是最新的消息开始消费。</li>
<li>通过调用 setAllocateMessageQueueStrategy 指定队列负载机制，默认平均分配。</li>
<li>通过调用 registerMessageListener 设置消息监听器，即消息处理逻辑，最终返回 CONSUME_SUCCESS（成功消费）或 RECONSUME_LATER（需要重试）。</li>
</ul>

<blockquote>
<p>温馨提示，关于消息消费的详细使用，将在后面的文章中进行详细介绍。</p>
</blockquote>

<h3 id="消息消费-api-版本演变说明">消息消费 API 版本演变说明</h3>

<p>RocketMQ 在消费端的 API 相对来说还是比较稳定的，只是在 RocketMQ 4.6.0 版本引入了 DefaultLitePullConsumer，如果上述 PULL 示例代码引用的 Client 包版本为 4.6.0，细心的读者朋友们肯定会发现 DefaultMQPullConsumer 已过期，取代它的正是 DefaultLitePullConsumer。那这是什么原因呢？</p>

<p>我想只要大家使用过 DefaultMQPullConsumer 编写代码后，就会发现这个类的 API 太底层，使用者需要考虑的问题太多，例如队列负载、消费进度存储等等方面，可以毫不夸张地说，要用好 DefaultMQPullConsumer 还是不那么容易的。</p>

<p>RocketMQ 设计者也意识到了这样的问题，故引入了 DefaultLitePullConsumer，按照官方文档上的介绍，该类具备如下特性：</p>

<ul>
<li>支持以订阅方式进行消息消费，支持消费队列自动再平衡。</li>
<li>支持手动分配队列的方式进行消息消费，此模式不支持队列自动再平衡。</li>
<li>提供 seek/commit 方法来重置、提交消费位点。</li>
</ul>

<p>温馨提示：由于 DefaultLitePullConsumer 的内容比较多，我们将在后续单独一篇文章对其参数、方法、使用示例进行详细介绍，故本篇只是告知其引入的目的。</p>

<h3 id="小结">小结</h3>

<p>本篇详细介绍了 RocketMQ PUSH、PULl 两种消息消费模式的核心 API，并对相关 API 进行了演示，后续会详细结合实际场景，并梳理核心参数，给出使用建议，后面引出了 RocketMQ 消费者一个重大的版本变更。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#620e0e0e5b565353525522050f030b0e4c010d0f" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12423aaccbede4',t:'MTczNDA1NDE1OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
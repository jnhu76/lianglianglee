<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=23&#32;消息轨迹、ACL&#32;与多副本搭建>
        <link rel="icon" href="/static/favicon.png">
        <title>23 消息轨迹、ACL 与多副本搭建 </title>
        
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
                            <h1 id="title" data-id="23 消息轨迹、ACL 与多副本搭建" class="title">23 消息轨迹、ACL 与多副本搭建</h1>
                            <div><h3 id="消息轨迹">消息轨迹</h3>

<h4 id="消息轨迹含义"><strong>消息轨迹含义</strong></h4>

<p>一条消息什么时候由哪台机器产生的、发送的耗时、消息大小、发送状态、存储在哪个 Broker 上、什么时候存储的以及存储在哪台 Broker 上、什么时候消费的、消费状态等信息，这些信息即消息轨迹，用于追踪消息从诞生到被消费的整个生命周期。</p>

<p>这些信息对于业务同学排查定位有着重要的意义，发送和消费往往在不同的业务部门。有了消息轨迹后一条消息有没有发送，发送成功了没，有没有消费一目了然，降低了彼此的沟通成本。</p>

<h4 id="消息轨迹使用"><strong>消息轨迹使用</strong></h4>

<p><strong>1. Broker 设置</strong></p>

<p>RocketMQ 的默认消息轨迹功能默认是关闭的，可以将 Broker 属性 traceTopicEnable 设置为 true 开启。消息轨迹默认存储在 RMQ_SYS_TRACE_TOPIC 的主题中，可以通过 msgTraceTopicName 修改。</p>

<table>
<thead>
<tr>
<th align="left">属性</th>
<th align="left">默认参数</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">traceTopicEnable</td>
<td align="left">false</td>
</tr>

<tr>
<td align="left">msgTraceTopicName</td>
<td align="left">RMQ_SYS_TRACE_TOPIC</td>
</tr>
</tbody>
</table>
<p><strong>2. 发送端使用</strong></p>

<p><strong>发送轨迹 API</strong></p>

<pre><code>public DefaultMQProducer(final String producerGroup, boolean enableMsgTrace, final String customizedTraceTopic){
   this(null, producerGroup, null, enableMsgTrace, customizedTraceTopic);
}

</code></pre>

<p>说明：enableMsgTrace 是否开启发送轨迹，默认 false；customizedTraceTopic 设置收集消息轨迹的自定义主题，默认为 RMQ_SYS_TRACE_TOPIC。</p>

<p><strong>发送代码示例</strong></p>

<pre><code>public static void main(String[] args) throws MQClientException, InterruptedException {
        DefaultMQProducer producer = new DefaultMQProducer(&quot;ProducerGroupName&quot;,true);
        producer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        producer.start();

        for (int i = 0; i &lt; 1; i++)
            try {
                Message msg = new Message(&quot;TopicTest&quot;,
                    &quot;TagA&quot;,
                    &quot;OrderID111&quot;,
                    &quot;Hello world&quot;.getBytes(RemotingHelper.DEFAULT_CHARSET));
                SendResult sendResult = producer.send(msg);
                System.out.printf(&quot;%s%n&quot;, sendResult);

            } catch (Exception e) {
                e.printStackTrace();
            }

        producer.shutdown();
    }

</code></pre>

<p>说明：创建 DefaultMQProducer 时将 enableMsgTrace 设置为 true 开启发送消息轨迹。</p>

<p><strong>3. 消费端使用</strong></p>

<p><strong>消费轨迹 API</strong></p>

<pre><code>public DefaultMQPushConsumer(final String consumerGroup, boolean enableMsgTrace, final String customizedTraceTopic) {
        this(null, consumerGroup, null, new AllocateMessageQueueAveragely(), enableMsgTrace, customizedTraceTopic);
}

</code></pre>

<p>说明：enableMsgTrace 是否开启消费轨迹，默认 false；customizedTraceTopic 设置收集消息轨迹的自定义主题，默认为 RMQ_SYS_TRACE_TOPIC。</p>

<p><strong>消费代码示例</strong></p>

<pre><code>public static void main(String[] args) throws InterruptedException, MQClientException {
        DefaultMQPushConsumer consumer = new DefaultMQPushConsumer(&quot;CID_JODIE_1&quot;,true);
        consumer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        consumer.subscribe(&quot;TopicTest&quot;, &quot;*&quot;);
        consumer.setConsumeFromWhere(ConsumeFromWhere.CONSUME_FROM_FIRST_OFFSET);
        consumer.setConsumeTimestamp(&quot;20181109221800&quot;);
        consumer.registerMessageListener(new MessageListenerConcurrently() {

            @Override
            public ConsumeConcurrentlyStatus consumeMessage(List&lt;MessageExt&gt; msgs, ConsumeConcurrentlyContext context) {
                System.out.printf(&quot;%s Receive New Messages: %s %n&quot;, Thread.currentThread().getName(), msgs);
                return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
            }
        });
        consumer.start();
        System.out.printf(&quot;Consumer Started.%n&quot;);
}

</code></pre>

<p>说明：创建 DefaultMQPushConsumer 将 enableMsgTrace 设置为 true 开启消费消息轨迹。</p>

<p><strong>4. 消息轨迹效果</strong></p>

<p>通过发送和消费一条消息，在 RocketMQ-Console 中看下消息轨迹的效果截图。</p>

<p><strong>发送消息内容</strong></p>

<pre><code>SendResult [sendStatus=SEND_OK, msgId=A9FE1075810A18B4AAC24A40738B0000, offsetMsgId=A9FE107500002A9F0000000000002147, messageQueue=MessageQueue [topic=TopicTest, brokerName=liangyong, queueId=1], queueOffset=2]

</code></pre>

<p><strong>消费消息内容</strong></p>

<pre><code>Receive New Messages: [MessageExt [brokerName=liangyong, queueId=1, storeSize=189, queueOffset=2, sysFlag=0, bornTimestamp=1600135337872, bornHost=/169.254.16.117:65532, storeTimestamp=1600135337883, storeHost=/169.254.16.117:10911, msgId=A9FE107500002A9F0000000000002147, commitLogOffset=8519, bodyCRC=198614610, reconsumeTimes=0, preparedTransactionOffset=0, toString()=Message{topic='TopicTest', flag=0, properties={MIN_OFFSET=0, MAX_OFFSET=3, KEYS=OrderID111, CONSUME_START_TIME=1600135337915, UNIQ_KEY=A9FE1075810A18B4AAC24A40738B0000, WAIT=true, TAGS=TagA}, body=[72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100], transactionId='null'}]] 

</code></pre>

<p><strong>消息轨迹展现</strong></p>

<p>在 RocketMQ 控制台，可以通过 message key 或者 message id 检索消息内容，如下图：</p>

<p><img src="assets/20200915100845698.jpg" alt="img" /></p>

<p>点击 MESSAGE TRACE DETAIL 可以查看消息轨迹，如下图：</p>

<p><img src="assets/20200915101143460.jpg" alt="img" /></p>

<h4 id="消息轨迹原理"><strong>消息轨迹原理</strong></h4>

<p><strong>发送轨迹原理</strong>：在消息发送前与发送后收集指标信息，并将指标信息异步发送到轨迹主题。</p>

<p><img src="assets/20200917223954432.png" alt="img" /></p>

<p><strong>消费轨迹原理</strong>：消费的消息轨迹有两部分，一部分在拉取消息后处理消息前收集指标异步发送都轨迹主题；另一部分处理消息后收集消息指标异步发送到轨迹主题。</p>

<p><img src="assets/20200918101244746.png" alt="img" /></p>

<h4 id="轨迹格式说明"><strong>轨迹格式说明</strong></h4>

<p>消息轨迹类型有三种，Pub 指发送轨迹，SubBefore 指消费前轨迹，SubAfter 指消费后轨迹。</p>

<p><strong>发送轨迹 Pub</strong></p>

<table>
<thead>
<tr>
<th align="left">名称</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">TraceType</td>
<td align="left">Pub 表示发送轨迹</td>
</tr>

<tr>
<td align="left">timeStamp</td>
<td align="left">存储时间</td>
</tr>

<tr>
<td align="left">regionId</td>
<td align="left">机房可用区，默认为 DefaultRegion（目前没有使用）</td>
</tr>

<tr>
<td align="left">groupName</td>
<td align="left">生产者组 producerGroup</td>
</tr>

<tr>
<td align="left">topic</td>
<td align="left">主题名称</td>
</tr>

<tr>
<td align="left">msgId</td>
<td align="left">消息 ID，由客户端生成</td>
</tr>

<tr>
<td align="left">tags</td>
<td align="left">消息 tag</td>
</tr>

<tr>
<td align="left">keys</td>
<td align="left">消息 key</td>
</tr>

<tr>
<td align="left">storeHost</td>
<td align="left">消息存储 Broker IP</td>
</tr>

<tr>
<td align="left">bodyLength</td>
<td align="left">消息体大小</td>
</tr>

<tr>
<td align="left">costTime</td>
<td align="left">发送消息耗时</td>
</tr>

<tr>
<td align="left">msgType</td>
<td align="left">消息类型：普通消息（Normal_Msg）、事物半消息（Trans_Msg_Half）、 事物提交消息（Trans_msg_Commit）、延迟消息（Delay_Msg）</td>
</tr>

<tr>
<td align="left">offsetMsgId</td>
<td align="left">消息 Id，由 Broker 生成</td>
</tr>

<tr>
<td align="left">isSuccess</td>
<td align="left">发送是否成功，true 表示成功、false 表示失败</td>
</tr>
</tbody>
</table>
<p><strong>消费前轨迹 SubBefore</strong></p>

<table>
<thead>
<tr>
<th align="left">名称</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">traceType</td>
<td align="left">SubBefore 表示消费前轨迹</td>
</tr>

<tr>
<td align="left">timeStamp</td>
<td align="left">消息存储时间</td>
</tr>

<tr>
<td align="left">regionId</td>
<td align="left">机房可用区（目前未使用）</td>
</tr>

<tr>
<td align="left">groupName</td>
<td align="left">消费组名称</td>
</tr>

<tr>
<td align="left">requestId</td>
<td align="left">请求标识</td>
</tr>

<tr>
<td align="left">msgId</td>
<td align="left">消息 Id</td>
</tr>

<tr>
<td align="left">retryTimes</td>
<td align="left">重试次数</td>
</tr>

<tr>
<td align="left">keys</td>
<td align="left">消息 key</td>
</tr>
</tbody>
</table>
<p><strong>消费前轨迹 SubAfter</strong></p>

<table>
<thead>
<tr>
<th align="left">名称</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">traceType</td>
<td align="left">SubAfter 表示消费后轨迹</td>
</tr>

<tr>
<td align="left">requestId</td>
<td align="left">请求标识</td>
</tr>

<tr>
<td align="left">msgId</td>
<td align="left">消息 Id</td>
</tr>

<tr>
<td align="left">costTime</td>
<td align="left">消费耗时</td>
</tr>

<tr>
<td align="left">isSuccess</td>
<td align="left">消费结果，true 消费成功、false 消费失败</td>
</tr>

<tr>
<td align="left">keys</td>
<td align="left">消息 key</td>
</tr>

<tr>
<td align="left">contextCode</td>
<td align="left">Broker 返回的消费状态，0:SUCCESS，1:TIME_OUT，2:EXCEPTION，3:RETURNNULL，4:FAILED</td>
</tr>
</tbody>
</table>

<h4 id="消息轨迹结语"><strong>消息轨迹结语</strong></h4>

<ul>
<li>在生产环境中使用消息轨迹，可以将 MQ 集群的一台 Broker 用于收集消息轨迹，避免轨迹消息对集群性能造成影响。</li>
<li>开源版本的消息轨迹中缺少消费的 IP 信息，即：我们不能查询到消息被哪个机器消费掉了。</li>
<li>开源版本中的消息轨迹组织格式用 char 字符拼接，解析使用数组，扩展性和兼容性不太友好。</li>
<li>基于此两位作者负责的 RocketMQ 集群中均未开启轨迹功能。</li>
</ul>

<h3 id="acl">ACL</h3>

<h4 id="acl-含义"><strong>ACL 含义</strong></h4>

<p>访问控制表（Access Control List，ACL）描述用户或角色对资源的访问控制权限，RocketMQ 中的 ACL 见下表说明。</p>

<p><strong>RocketMQ 中的 ACL 含义说明：</strong></p>

<table>
<thead>
<tr>
<th align="left">含义</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">用户</td>
<td align="left">在 plain_acl.yml 配置文件用 accessKey 表示</td>
</tr>

<tr>
<td align="left">角色</td>
<td align="left">admin 和其他角色</td>
</tr>

<tr>
<td align="left">资源</td>
<td align="left">包括主题和消费组</td>
</tr>

<tr>
<td align="left">权限</td>
<td align="left">DENY 表示无权限 ANY 表示拥有 PUB 或者 SUB 权限 PUB 表示拥有主题发送权限 SUB 表示拥有消费组订阅权限</td>
</tr>
</tbody>
</table>

<h4 id="acl-使用示例"><strong>ACL 使用示例</strong></h4>

<p>将 <code>aclEnable = true</code> 添加到 Broker 配置文件中，另外添加 <code>${ROCKETMQ_HOME}/conf/plain_acl.yml</code> 文件，用于 ACL 控制。</p>

<p><strong>1. Broker 配置</strong></p>

<pre><code>brokerClusterName = DefaultCluster
brokerName = broker-a
brokerId = 0
deleteWhen = 04
fileReservedTime = 48
brokerRole = ASYNC_MASTER
flushDiskType = ASYNC_FLUSH
traceTopicEnable = true
aclEnable = true

</code></pre>

<p><strong>说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">含义</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">aclEnable</td>
<td align="left">默认 false，开启 ACL 需要设置为 true</td>
</tr>

<tr>
<td align="left">ROCKETMQ_HOME</td>
<td align="left">可以通过 -Drocketmq.home.dir 指定 MQ 根目录</td>
</tr>

<tr>
<td align="left">ACL fileName</td>
<td align="left">默认 /conf/plain_acl.yml，可以通过 -Drocketmq.acl.plain.file 指定 ACL 文件名称</td>
</tr>
</tbody>
</table>
<p><strong>2. plain_acl.yml 配置</strong></p>

<p>ACL 配置文件由全局白名单配置（globalWhiteRemoteAddresses）和账户配置（accounts）两部分构成。</p>

<pre><code>globalWhiteRemoteAddresses:

accounts:
- accessKey: RocketMQ
  secretKey: 12345678
  whiteRemoteAddress:
  admin: false
  defaultTopicPerm: DENY
  defaultGroupPerm: SUB
  topicPerms:
  - TopicTes1=DENY
  - TopicTest2=PUB|SUB
  groupPerms:
  - consumerTest=DENY

- accessKey: rocketmq2
  secretKey: 12345678
  whiteRemoteAddress: 192.168.1.*
  admin: true

</code></pre>

<p><strong>说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">含义</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">globalWhiteRemoteAddresses</td>
<td align="left">全局白名单配置，策略如下： 空：忽略白名单，继续执行下面校验 全匹配模式：全部放行不会执行后面校验 例如：* 或 <em>.</em>.<em>.</em> 或 <em>:</em>:<em>:</em>:<em>:</em>:<em>:</em> 多 IP 模式：表示白名单 IP 在设置区间段的放行 例如：192.168.0.{1,2} 或 192.168.1.1,192.168.1.2 或 192.168.*. 或 <em>192.168.1-10.5-50</em></td>
</tr>

<tr>
<td align="left">accessKey</td>
<td align="left">唯一用户标识</td>
</tr>

<tr>
<td align="left">secretKey</td>
<td align="left">访问密码</td>
</tr>

<tr>
<td align="left">whiteRemoteAddress</td>
<td align="left">用户级白名单，格式同 globalWhiteRemoteAddresses</td>
</tr>

<tr>
<td align="left">admin</td>
<td align="left">是否为管理员，管理员拥有所有资源访问权限</td>
</tr>

<tr>
<td align="left">defaultTopicPerm</td>
<td align="left">默认主题权限，默认值 DENY</td>
</tr>

<tr>
<td align="left">defaultGroupPerm</td>
<td align="left">默认消费组权限，默认值 DENY</td>
</tr>

<tr>
<td align="left">topicPerms</td>
<td align="left">详细的主题权限</td>
</tr>

<tr>
<td align="left">groupPerms</td>
<td align="left">详细的消费组权限</td>
</tr>
</tbody>
</table>
<p><strong>3. ACL 发送示例</strong></p>

<p>在上面的配置文件中，将 TopicTes1 设置了 DENY 权限，即禁止发送和消费；将 TopicTest2 设置成了 PUB|SUB 权限，即允许发送和订阅权限。下面例子尝试向主题 TopicTes1 发送消息，观察其是否可以成功。</p>

<p><strong>禁止发送示例</strong></p>

<pre><code>public class AclSendTest {
    private static final String ACL_ACCESS_KEY = &quot;RocketMQ&quot;;
    private static final String ACL_SECRET_KEY = &quot;12345678&quot;;

    public static void main(String[] args) throws MQClientException, InterruptedException {
        producer();
    }
    public static void producer() throws MQClientException {
        DefaultMQProducer producer = new DefaultMQProducer(&quot;ProducerGroupName&quot;, getAclRPCHook());
        producer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        producer.start();
        for (int i = 0; i &lt; 1; i++)
            try {
                {
                    Message msg = new Message(&quot;TopicTest1&quot;,
                            &quot;TagA&quot;,
                            &quot;OrderID188&quot;,
                            &quot;Hello world&quot;.getBytes(RemotingHelper.DEFAULT_CHARSET));
                    SendResult sendResult = producer.send(msg);
                    System.out.printf(&quot;%s%n&quot;, sendResult);
                }

            } catch (Exception e) {
                e.printStackTrace();
            }

        producer.shutdown();
    }

    static RPCHook getAclRPCHook() {
        return new AclClientRPCHook(new SessionCredentials(ACL_ACCESS_KEY,ACL_SECRET_KEY));
    }
}

</code></pre>

<p><strong>禁止发送截图</strong></p>

<p><img src="assets/20200920202144594.jpg" alt="img" /></p>

<p><strong>禁止发送说明</strong></p>

<p>用户 RocketMQ 向主题 TopicTes1 发送消息时抛出 AclException，拒绝访问，如果将代码中主题换成 TopicTes2，则可以发送成功，接着看下文。</p>

<p><strong>4. ACL 消费示例</strong></p>

<p>在上面的配置文件中，将 consumerTest 设置了 DENY 权限，即禁止消费消息。由于 TopicTes2 设置为允许发送，我们下面尝试向 TopicTes2 发送一条消息，consumerTest 订阅了 TopicTes2 观察其是否可以消费。</p>

<p><strong>允许发送示例</strong></p>

<pre><code>public class AclSendTest {
    private static final String ACL_ACCESS_KEY = &quot;RocketMQ&quot;;
    private static final String ACL_SECRET_KEY = &quot;12345678&quot;;

    public static void main(String[] args) throws MQClientException, InterruptedException {
        producer();
    }
    public static void producer() throws MQClientException {
        DefaultMQProducer producer = new DefaultMQProducer(&quot;ProducerGroupName&quot;, getAclRPCHook());
        producer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        producer.start();
        for (int i = 0; i &lt; 1; i++)
            try {
                {
                    Message msg = new Message(&quot;TopicTest2&quot;,
                            &quot;TagA&quot;,
                            &quot;OrderID188&quot;,
                            &quot;Hello world&quot;.getBytes(RemotingHelper.DEFAULT_CHARSET));
                    SendResult sendResult = producer.send(msg);
                    System.out.printf(&quot;%s%n&quot;, sendResult);
                }

            } catch (Exception e) {
                e.printStackTrace();
            }

        producer.shutdown();
    }

    static RPCHook getAclRPCHook() {
        return new AclClientRPCHook(new SessionCredentials(ACL_ACCESS_KEY,ACL_SECRET_KEY));
    }
}

</code></pre>

<p><strong>允许发送结果</strong></p>

<pre><code>SendResult [sendStatus=SEND_OK, msgId=C0A800667FB218B4AAC2663AB66F0000, offsetMsgId=C0A8006600002A9F00000000000085EA, messageQueue=MessageQueue [topic=TopicTest2, brokerName=broker-a, queueId=0], queueOffset=2]

</code></pre>

<p><strong>禁止消费示例</strong></p>

<pre><code>public class AclConsumeTest {
    private static final String ACL_ACCESS_KEY = &quot;RocketMQ&quot;;
    private static final String ACL_SECRET_KEY = &quot;12345678&quot;;

    public static void main(String[] args) throws MQClientException, InterruptedException {
        pushConsumer();
    }

    public static void pushConsumer() throws MQClientException {

        DefaultMQPushConsumer consumer = new DefaultMQPushConsumer(&quot;consumerTest&quot;, getAclRPCHook(), new AllocateMessageQueueAveragely());
        consumer.setNamesrvAddr(&quot;127.0.0.1:9876&quot;);
        consumer.subscribe(&quot;TopicTest2&quot;, &quot;*&quot;);
        consumer.setConsumeFromWhere(ConsumeFromWhere.CONSUME_FROM_FIRST_OFFSET);
        consumer.setConsumeTimestamp(&quot;20180422221800&quot;);
        consumer.registerMessageListener(new MessageListenerConcurrently() {

            @Override
            public ConsumeConcurrentlyStatus consumeMessage(List&lt;MessageExt&gt; msgs, ConsumeConcurrentlyContext context) {
                System.out.printf(&quot;%s Receive New Messages: %s %n&quot;, Thread.currentThread().getName(), msgs);
                printBody(msgs);
                return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
            }
        });
        consumer.start();
        System.out.printf(&quot;Consumer Started.%n&quot;);
    }

    static RPCHook getAclRPCHook() {
        return new AclClientRPCHook(new SessionCredentials(ACL_ACCESS_KEY,ACL_SECRET_KEY));
    }

    private static void printBody(List&lt;MessageExt&gt; msg) {
        if (msg == null || msg.size() == 0)
            return;
        for (MessageExt m : msg) {
            if (m != null) {
                System.out.printf(&quot;msgId : %s  body : %s  \n\r&quot;, m.getMsgId(), new String(m.getBody()));
            }
        }
    }
}

</code></pre>

<p><strong>禁止消费截图</strong></p>

<p><img src="assets/20200920203904115.jpg" alt="img" /></p>

<p><strong>禁止消费说明</strong></p>

<p>我们向 TopicTest2 成功发送了一条消息，但由于消费组 consumerTest 被设置成禁止消费，所有未能收到该消息。</p>

<h4 id="acl-命令汇总"><strong>ACL 命令汇总</strong></h4>

<p>RocketMQ 提供了一系列命令动态更新 Acl 配置文件，使设置的权限及时生效。</p>

<p><strong>1. 获取 ACL 配置版本</strong></p>

<p>使用 clusterAclConfigVersion 命令查看版本信息。</p>

<p><strong>参数说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">-b</td>
<td align="left">Broker 地址，更新特定的 Broker</td>
</tr>

<tr>
<td align="left">-c</td>
<td align="left">集群名称，更新集群中的所有 Broker</td>
</tr>

<tr>
<td align="left">-n</td>
<td align="left">namesrv 地址</td>
</tr>
</tbody>
</table>
<p><strong>命令示例</strong></p>

<pre><code>$ bin/mqadmin clusterAclConfigVersion -n x.x.x.x:9876 -c DefaultCluster
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
#Cluster Name     #Broker Name            #Broker Addr            #AclConfigVersionNum  #AclLastUpdateTime
DefaultCluster    broker-a                x.x.x.x:10911      0                     2020-09-20 22:42:59
get cluster's plain access config version success.

</code></pre>

<p><strong>2. 获取 Acl 权限配置</strong></p>

<p>使用 getAccessConfigSubCommand 获取 ACL 的配置信息。</p>

<p><strong>参数说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">-b</td>
<td align="left">Broker 地址，更新特定的 Broker</td>
</tr>

<tr>
<td align="left">-c</td>
<td align="left">集群名称，更新集群中的所有 Broker</td>
</tr>

<tr>
<td align="left">-n</td>
<td align="left">namesrv 地址</td>
</tr>
</tbody>
</table>
<p><strong>命令示例</strong></p>

<pre><code>$ bin/mqadmin getAccessConfigSubCommand -n x.x.x.x:9876 -c DefaultCluster
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.

globalWhiteRemoteAddresses: [10.10.103.*, 192.168.0.*]

accounts:
  accessKey         : RocketMQ
  secretKey         : 12345678
  whiteRemoteAddress:
  admin             : false
  defaultTopicPerm  : DENY
  defaultGroupPerm  : SUB
  topicPerms        : [topicA=DENY, topicB=PUB|SUB, topicC=SUB]
  groupPerms        : [groupA=DENY, groupB=PUB|SUB, groupC=SUB]

  accessKey         : rocketmq2
  secretKey         : 12345678
  whiteRemoteAddress: 192.168.1.*
  admin             : true
  defaultTopicPerm  :
  defaultGroupPerm  :
  topicPerms        :
  groupPerms        :

</code></pre>

<p><strong>3. 更新全局白名单</strong></p>

<p>使用 updateGlobalWhiteAddr 对 ACL 的全局白名单 globalWhiteRemoteAddresses 进行变更。</p>

<p><strong>参数说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">-b</td>
<td align="left">Broker 地址，更新特定的 Broker</td>
</tr>

<tr>
<td align="left">-c</td>
<td align="left">集群名称，更新集群中的所有 Broker</td>
</tr>

<tr>
<td align="left">-n</td>
<td align="left">namesrv 地址</td>
</tr>

<tr>
<td align="left">-g</td>
<td align="left">全局白名单值，例如：10.10.103.<em>,192.168.0.</em></td>
</tr>
</tbody>
</table>
<p><strong>命令示例</strong></p>

<pre><code>$ bin/mqadmin updateGlobalWhiteAddr -n x.x.x.x:9876 -c DefaultCluster -g 10.10.113.*,192.168.20.*
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
update global white remote addresses to x.x.x.x:10911 success.

</code></pre>

<p><strong>查看生效</strong></p>

<pre><code>$ bin/mqadmin getAccessConfigSubCommand -n x.x.x.x:9876 -c DefaultCluster
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.

globalWhiteRemoteAddresses: [10.10.113.*, 192.168.20.*]

accounts:
  accessKey         : RocketMQ
  secretKey         : 12345678
  whiteRemoteAddress:
  admin             : false
  defaultTopicPerm  : DENY
  defaultGroupPerm  : SUB
  topicPerms        : [topicA=DENY, topicB=PUB|SUB, topicC=SUB]
  groupPerms        : [groupA=DENY, groupB=PUB|SUB, groupC=SUB]

  accessKey         : rocketmq2
  secretKey         : 12345678
  whiteRemoteAddress: 192.168.1.*
  admin             : true
  defaultTopicPerm  :
  defaultGroupPerm  :
  topicPerms        :
  groupPerms        :

</code></pre>

<p>说明：全局白名单已经被更新。</p>

<p><strong>4. 更新用户配置</strong></p>

<p>对于用户账户的配置的变更通过 updateAclConfig 来实现。</p>

<p><strong>参数说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">-a</td>
<td align="left">指定 accessKey，变更哪个用户的配置</td>
</tr>

<tr>
<td align="left">-b</td>
<td align="left">Broker 地址，更新特定的 Broker</td>
</tr>

<tr>
<td align="left">-c</td>
<td align="left">集群名称，更新集群中的所有 Broker</td>
</tr>

<tr>
<td align="left">-n</td>
<td align="left">namesrv 地址</td>
</tr>

<tr>
<td align="left">-g</td>
<td align="left">设置 groupPerms 消费组权限，格式为：<code>groupD=DENY,groupD=SUB</code></td>
</tr>

<tr>
<td align="left">-i</td>
<td align="left">设置 Acl 文件中 defaultTopicPerm 权限</td>
</tr>

<tr>
<td align="left">-m</td>
<td align="left">设置 Acl 文件中 admin 权限</td>
</tr>

<tr>
<td align="left">-s</td>
<td align="left">设置 Acl 文件中 secretKey 密钥值</td>
</tr>

<tr>
<td align="left">-t</td>
<td align="left">设置 topicPerms 主题权限，格式为：topicA=DENY,topicD=SUB</td>
</tr>

<tr>
<td align="left">-u</td>
<td align="left">设置 Acl 文件中的默认消费组 defaultGroupPerm 权限</td>
</tr>

<tr>
<td align="left">-w</td>
<td align="left">设置 Acl 文件中该用户下的白名单权限 whiteRemoteAddress</td>
</tr>
</tbody>
</table>
<p><strong>命令示例</strong></p>

<pre><code>$ bin/mqadmin updateAclConfig -n x.x.x.x:9876 -c DefaultCluster -a RocketMQ -s 87654321 -t testTopicA=DENY,testTopicb=SUB
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
create or update plain access config to x.x.x.x:10911 success.

</code></pre>

<p><strong>查看生效</strong></p>

<pre><code>$ bin/mqadmin getAccessConfigSubCommand -n uat-mq2.ttbike.com.cn:9876 -c DefaultCluster
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.

globalWhiteRemoteAddresses: [10.10.113.*, 192.168.20.*]

accounts:
  accessKey         : rocketmq2
  secretKey         : 12345678
  whiteRemoteAddress: 192.168.1.*
  admin             : true
  defaultTopicPerm  :
  defaultGroupPerm  :
  topicPerms        :
  groupPerms        :

  accessKey         : RocketMQ
  secretKey         : 87654321
  whiteRemoteAddress:
  admin             : false
  defaultTopicPerm  : DENY
  defaultGroupPerm  : SUB
  topicPerms        : [testTopicA=DENY, testTopicb=SUB]
  groupPerms        : [groupA=DENY, groupB=PUB|SUB, groupC=SUB]

</code></pre>

<p>说明：用户 RocketMQ 的密钥 secretKey 和主题权限 topicPerms 已变更生效。</p>

<p><strong>5. 删除用户配置</strong></p>

<p>通过 deleteAccessConfig 删除指定用户的 ACL 配置信息。</p>

<p><strong>参数说明</strong></p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">-b</td>
<td align="left">Broker 地址，更新特定的 Broker</td>
</tr>

<tr>
<td align="left">-c</td>
<td align="left">集群名称，更新集群中的所有 Broker</td>
</tr>

<tr>
<td align="left">-n</td>
<td align="left">namesrv 地址</td>
</tr>

<tr>
<td align="left">-a</td>
<td align="left">指定特定用户 accessKey</td>
</tr>
</tbody>
</table>
<p><strong>命令示例</strong></p>

<pre><code>$ bin/mqadmin deleteAccessConfig -n x.x.x.x:9876 -c DefaultCluster -a RocketMQ
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
delete plain access config account to x.x.x.x:10911 success.

</code></pre>

<p><strong>查看生效</strong></p>

<pre><code>$ bin/mqadmin getAccessConfigSubCommand -n x.x.x.x:9876 -c DefaultCluster
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.

globalWhiteRemoteAddresses: [10.10.113.*, 192.168.20.*]

accounts:
  accessKey         : rocketmq2
  secretKey         : 12345678
  whiteRemoteAddress: 192.168.1.*
  admin             : true
  defaultTopicPerm  :
  defaultGroupPerm  :
  topicPerms        :
  groupPerms        :

</code></pre>

<p>说明：用户 RocketMQ 的相关 Acl 配置已被全部删除。</p>

<h4 id="acl-原理简述"><strong>ACL 原理简述</strong></h4>

<p><strong>1. 规则加载</strong></p>

<p>规则配置在 plain_acl.yml 文件中，需要加载到 Broker 缓存中使其生效。启动时会加载 acl 文件的内容，当其变更了也需要动态加载规则内容，详见如下流程。</p>

<p><img src="assets/20200921223149842.png" alt="img" /></p>

<p><strong>2. 权限校验</strong></p>

<p>注册的钩子程序通过 NettyServerHandler 实现，在 Broker 当前的 channel 接收到客户端消息时执行校验逻辑，入口为 NettyServerHandler#channelRead0() 以及 processRequestCommand#doBeforeRpcHooks。规则校验的入口位于 PlainAccessValidator#validate 方法，下面是校验简图。</p>

<p><img src="assets/20200921231728157.png" alt="img" /></p>

<h4 id="acl-结语"><strong>ACL 结语</strong></h4>

<ul>
<li>评估是否需要 ACL 通常大部分场景是不需要的，如果对数据敏感可以通过对内容加密实现。</li>
<li>如果几千个资源（主题和消费组）都配置了 ACL 那么配置文件是庞大的。</li>
<li>如果要用建议局部使用，比如：资金往来的主题、消费组等。</li>
<li>基于此两位作者负责的 RocketMQ 集群均未开启 ACL 功能。</li>
</ul>

<h3 id="多副本搭建">多副本搭建</h3>

<h4 id="多副本意义"><strong>多副本意义</strong></h4>

<p>RocketMQ 开源版本在 4.5.0 版本开始支持多副本（DLedger），在以前的版本中只支持主从模式。</p>

<p>主从模式存在的问题：</p>

<ul>
<li>如果主节点挂掉了不能动态切换到从节点，这一组 Broker 节点不能提供写入服务；</li>
<li>设置主从异步复制模式时，如果主节点意外挂掉，数据可能没有全部复制到从节点，存在数据丢失风险。</li>
</ul>

<p>多副本使用 Raft 协议在节点意外掉线后能够完成自动选主，提高集群的高可用和保证数据的一致性。</p>

<h4 id="多副本搭建-1"><strong>多副本搭建</strong></h4>

<p>由于 DLedger 基于 Raft 协议开发的功能，需要过半数选举，最少 3 个节点组成一个 Raft 组。</p>

<p><strong>broker-n0.conf</strong></p>

<pre><code>brokerClusterName = RaftCluster
brokerName=RaftNode00
listenPort=30911
namesrvAddr=127.0.0.1:9876
storePathRootDir=/tmp/rmqstore/node00
storePathCommitLog=/tmp/rmqstore/node00/commitlog
enableDLegerCommitLog=true
dLegerGroup=RaftNode00
dLegerPeers=n0-127.0.0.1:40911;n1-127.0.0.1:40912;n2-127.0.0.1:40913
### must be unique
dLegerSelfId=n0
sendMessageThreadPoolNums=16

</code></pre>

<p><strong>broker-n1.conf</strong></p>

<pre><code>brokerClusterName = RaftCluster
brokerName=RaftNode00
listenPort=30921
namesrvAddr=127.0.0.1:9876
storePathRootDir=/tmp/rmqstore/node01
storePathCommitLog=/tmp/rmqstore/node01/commitlog
enableDLegerCommitLog=true
dLegerGroup=RaftNode00
dLegerPeers=n0-127.0.0.1:40911;n1-127.0.0.1:40912;n2-127.0.0.1:40913
### must be unique
dLegerSelfId=n1
sendMessageThreadPoolNums=16

</code></pre>

<p><strong>broker-n2.conf</strong></p>

<pre><code>brokerClusterName = RaftCluster
brokerName=RaftNode00
listenPort=30931
namesrvAddr=127.0.0.1:9876
storePathRootDir=/tmp/rmqstore/node02
storePathCommitLog=/tmp/rmqstore/node02/commitlog
enableDLegerCommitLog=true
dLegerGroup=RaftNode00
dLegerPeers=n0-127.0.0.1:40911;n1-127.0.0.1:40912;n2-127.0.0.1:40913
### must be unique
dLegerSelfId=n2
sendMessageThreadPoolNums=16

</code></pre>

<p><strong>启动三个节点：</strong></p>

<pre><code>nohup bin/mqbroker -c conf/dledger/broker-n0.conf &amp;
nohup bin/mqbroker -c conf/dledger/broker-n1.conf &amp;
nohup bin/mqbroker -c conf/dledger/broker-n2.conf &amp;

</code></pre>

<p><strong>查看是否启动成功：</strong></p>

<pre><code>$ bin/mqadmin clusterList -n localhost:9876
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
#Cluster Name     #Broker Name            #BID  #Addr                  #Version                #InTPS(LOAD)       #OutTPS(LOAD) #PCWait(ms) #Hour #SPACE
RaftCluster       RaftNode00              0     x.x.x.x:30921     V4_7_0                   0.00(0,0ms)         0.00(0,0ms)          0 444663.49 -1.0000
RaftCluster       RaftNode00              1     x.x.x.x:30911     V4_7_0                   0.00(0,0ms)         0.00(0,0ms)          0 444663.49 -1.0000
RaftCluster       RaftNode00              3     x.x.x.x:30931     V4_7_0                   0.00(0,0ms)         0.00(0,0ms)          0 444663.49 -1.0000

</code></pre>

<p>说明：BID 为 0 表示表示 Master，其他两个均为 Follower。</p>

<p><strong>控制台截图：</strong></p>

<p><img src="assets/20200922235326160.jpg" alt="img" /></p>

<p><strong>查看发送消息：</strong></p>

<p><img src="assets/20200922235557734.jpg" alt="img" /></p>

<p>说明：通过以上步骤，我们完成多副本的搭建过程。</p>

<h4 id="重新选主"><strong>重新选主</strong></h4>

<p>我们通过 kill 掉 Master 的方式来验证 DLedger 选主情况，上面的 clusterList 截图中我们看到 Master 为 x.x.x.x:30921，将该进程 kill 掉后观察一下。</p>

<pre><code>$ bin/mqadmin clusterList -n localhost:9876
RocketMQLog:WARN No appenders could be found for logger (io.netty.util.internal.PlatformDependent0).
RocketMQLog:WARN Please initialize the logger system properly.
#Cluster Name     #Broker Name            #BID  #Addr                  #Version                #InTPS(LOAD)       #OutTPS(LOAD) #PCWait(ms) #Hour #SPACE
RaftCluster       RaftNode00              0     x.x.x.x:30931     V4_7_0                   0.00(0,0ms)         0.00(0,0ms)          0 444664.03 -1.0000
RaftCluster       RaftNode00              1     x.x.x.x:30911     V4_7_0                   0.00(0,0ms)         0.00(0,0ms)          0 444664.03 -1.0000

</code></pre>

<p>说明：kill 掉原 Master 后，完成自动选主，新的 Master 为 x.x.x.x:30931。</p>

<h4 id="参数说明"><strong>参数说明</strong></h4>

<p>配置文件中多副本的参数说明见下面表格。</p>

<table>
<thead>
<tr>
<th align="left">参数</th>
<th align="left">说明</th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">enableDLegerCommitLog</td>
<td align="left">是否启用 DLedger，默认 false</td>
</tr>

<tr>
<td align="left">dLegerGroup</td>
<td align="left">节点所属的 Raft 组，建议与 broker 一致</td>
</tr>

<tr>
<td align="left">dLegerPeers</td>
<td align="left">集群节点信息，示例：n0-127.0.0.1:40911;n1-127.0.0.1:40912;n2-127.0.0.1:40913</td>
</tr>

<tr>
<td align="left">dLegerSelfId</td>
<td align="left">当前节点 id。取自 legerPeers 中条目的开头，即上述示例中的 n0，并且特别需要强调，只能第一个字符为英文，其他字符需要配置成数字</td>
</tr>
</tbody>
</table>

<h4 id="参考资料"><strong>参考资料</strong></h4>

<p>Raft 的学习资料见下面链接，供学习使用。DLedger 的源码解读，见《RocketMQ 技术内幕》第二版。</p>

<ul>
<li><a href="https://github.com/brandonwang001/raft_translation/blob/master/raft_translation.pdf" target="_blank">Raft 论文译文</a></li>
<li><a href="https://ramcloud.atlassian.net/wiki/download/attachments/6586375/raft.pdf" target="_blank">Raft 论文原文</a></li>
<li><a href="http://thesecretlivesofdata.com/raft/" target="_blank">Raft 动画</a></li>
<li><a href="https://raft.github.io/" target="_blank">Raft 可视化操作</a></li>
</ul>

<h4 id="多副本结语"><strong>多副本结语</strong></h4>

<p>使用多副本时，请做好压测，压测的 TPS 是否满足业务的需求，作者曾做过多副本压测与主从异步的 TPS 有相当大的差距。</p>

<p>在 TPS 满足的情况下，建议使用多副本架构，尤其是支付类可以优先使用。</p>

<p>如果线上已经存在了主从默认的架构如何升级到 DLedger 模式呢？</p>

<ol>
<li>可以参考前面平滑扩所容的方式，将多副本组成的 Raft 组加入到原集群中</li>
<li>关闭原主从架构节点的写入权限</li>
<li>在过了日志存储时间后，将主从架构节点下线</li>
</ol>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c1adadadf8f5f0f0f1f681a6aca0a8adefa2aeac" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1246b43b21ede4',t:'MTczNDA1NDM0Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
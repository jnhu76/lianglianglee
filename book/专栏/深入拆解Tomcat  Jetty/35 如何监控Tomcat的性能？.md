<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=35&#32;如何监控Tomcat的性能？>
        <link rel="icon" href="/static/favicon.png">
        <title>35 如何监控Tomcat的性能？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 Java程序员如何快速成长？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/00%20%e5%bc%80%e7%af%87%e8%af%8d%20Java%e7%a8%8b%e5%ba%8f%e5%91%98%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e6%88%90%e9%95%bf%ef%bc%9f.md">00 开篇词 Java程序员如何快速成长？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 Web容器学习路径.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/01%20Web%e5%ae%b9%e5%99%a8%e5%ad%a6%e4%b9%a0%e8%b7%af%e5%be%84.md">01 Web容器学习路径.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 HTTP协议必知必会.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/02%20HTTP%e5%8d%8f%e8%ae%ae%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a.md">02 HTTP协议必知必会.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 你应该知道的Servlet规范和Servlet容器.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/03%20%e4%bd%a0%e5%ba%94%e8%af%a5%e7%9f%a5%e9%81%93%e7%9a%84Servlet%e8%a7%84%e8%8c%83%e5%92%8cServlet%e5%ae%b9%e5%99%a8.md">03 你应该知道的Servlet规范和Servlet容器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 实战：纯手工打造和运行一个Servlet.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/04%20%e5%ae%9e%e6%88%98%ef%bc%9a%e7%ba%af%e6%89%8b%e5%b7%a5%e6%89%93%e9%80%a0%e5%92%8c%e8%bf%90%e8%a1%8c%e4%b8%80%e4%b8%aaServlet.md">04 实战：纯手工打造和运行一个Servlet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 Tomcat系统架构（上）： 连接器是如何设计的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/05%20Tomcat%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%c2%a0%e8%bf%9e%e6%8e%a5%e5%99%a8%e6%98%af%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e7%9a%84%ef%bc%9f.md">05 Tomcat系统架构（上）： 连接器是如何设计的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Tomcat系统架构（下）：聊聊多层容器的设计.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/06%20Tomcat%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%81%8a%e8%81%8a%e5%a4%9a%e5%b1%82%e5%ae%b9%e5%99%a8%e7%9a%84%e8%ae%be%e8%ae%a1.md">06 Tomcat系统架构（下）：聊聊多层容器的设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Tomcat如何实现一键式启停？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/07%20Tomcat%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e9%94%ae%e5%bc%8f%e5%90%af%e5%81%9c%ef%bc%9f.md">07 Tomcat如何实现一键式启停？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 Tomcat的“高层们”都负责做什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/08%20Tomcat%e7%9a%84%e2%80%9c%e9%ab%98%e5%b1%82%e4%bb%ac%e2%80%9d%e9%83%bd%e8%b4%9f%e8%b4%a3%e5%81%9a%e4%bb%80%e4%b9%88%ef%bc%9f.md">08 Tomcat的“高层们”都负责做什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 比较：Jetty架构特点之Connector组件.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/09%20%e6%af%94%e8%be%83%ef%bc%9aJetty%e6%9e%b6%e6%9e%84%e7%89%b9%e7%82%b9%e4%b9%8bConnector%e7%bb%84%e4%bb%b6.md">09 比较：Jetty架构特点之Connector组件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 比较：Jetty架构特点之Handler组件.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/10%20%e6%af%94%e8%be%83%ef%bc%9aJetty%e6%9e%b6%e6%9e%84%e7%89%b9%e7%82%b9%e4%b9%8bHandler%e7%bb%84%e4%bb%b6.md">10 比较：Jetty架构特点之Handler组件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 总结：从Tomcat和Jetty中提炼组件化设计规范.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/11%20%e6%80%bb%e7%bb%93%ef%bc%9a%e4%bb%8eTomcat%e5%92%8cJetty%e4%b8%ad%e6%8f%90%e7%82%bc%e7%bb%84%e4%bb%b6%e5%8c%96%e8%ae%be%e8%ae%a1%e8%a7%84%e8%8c%83.md">11 总结：从Tomcat和Jetty中提炼组件化设计规范.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 实战：优化并提高Tomcat启动速度.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/12%20%e5%ae%9e%e6%88%98%ef%bc%9a%e4%bc%98%e5%8c%96%e5%b9%b6%e6%8f%90%e9%ab%98Tomcat%e5%90%af%e5%8a%a8%e9%80%9f%e5%ba%a6.md">12 实战：优化并提高Tomcat启动速度.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 热点问题答疑（1）：如何学习源码？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/13%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%881%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ad%a6%e4%b9%a0%e6%ba%90%e7%a0%81%ef%bc%9f.md">13 热点问题答疑（1）：如何学习源码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 NioEndpoint组件：Tomcat如何实现非阻塞I_O？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/14%20NioEndpoint%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%9d%9e%e9%98%bb%e5%a1%9eI_O%ef%bc%9f.md">14 NioEndpoint组件：Tomcat如何实现非阻塞I_O？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Nio2Endpoint组件：Tomcat如何实现异步I_O？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/15%20Nio2Endpoint%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%bc%82%e6%ad%a5I_O%ef%bc%9f.md">15 Nio2Endpoint组件：Tomcat如何实现异步I_O？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 AprEndpoint组件：Tomcat APR提高I_O性能的秘密.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/16%20AprEndpoint%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%20APR%e6%8f%90%e9%ab%98I_O%e6%80%a7%e8%83%bd%e7%9a%84%e7%a7%98%e5%af%86.md">16 AprEndpoint组件：Tomcat APR提高I_O性能的秘密.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 Executor组件：Tomcat如何扩展Java线程池？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/17%20Executor%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e6%89%a9%e5%b1%95Java%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f.md">17 Executor组件：Tomcat如何扩展Java线程池？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 新特性：Tomcat如何支持WebSocket？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/18%20%e6%96%b0%e7%89%b9%e6%80%a7%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e6%94%af%e6%8c%81WebSocket%ef%bc%9f.md">18 新特性：Tomcat如何支持WebSocket？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 比较：Jetty的线程策略EatWhatYouKill.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/19%20%e6%af%94%e8%be%83%ef%bc%9aJetty%e7%9a%84%e7%ba%bf%e7%a8%8b%e7%ad%96%e7%95%a5EatWhatYouKill.md">19 比较：Jetty的线程策略EatWhatYouKill.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 总结：Tomcat和Jetty中的对象池技术.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/20%20%e6%80%bb%e7%bb%93%ef%bc%9aTomcat%e5%92%8cJetty%e4%b8%ad%e7%9a%84%e5%af%b9%e8%b1%a1%e6%b1%a0%e6%8a%80%e6%9c%af.md">20 总结：Tomcat和Jetty中的对象池技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 总结：Tomcat和Jetty的高性能、高并发之道.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/21%20%e6%80%bb%e7%bb%93%ef%bc%9aTomcat%e5%92%8cJetty%e7%9a%84%e9%ab%98%e6%80%a7%e8%83%bd%e3%80%81%e9%ab%98%e5%b9%b6%e5%8f%91%e4%b9%8b%e9%81%93.md">21 总结：Tomcat和Jetty的高性能、高并发之道.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 热点问题答疑（2）：内核如何阻塞与唤醒进程？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/22%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%882%ef%bc%89%ef%bc%9a%e5%86%85%e6%a0%b8%e5%a6%82%e4%bd%95%e9%98%bb%e5%a1%9e%e4%b8%8e%e5%94%a4%e9%86%92%e8%bf%9b%e7%a8%8b%ef%bc%9f.md">22 热点问题答疑（2）：内核如何阻塞与唤醒进程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Host容器：Tomcat如何实现热部署和热加载？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/23%20Host%e5%ae%b9%e5%99%a8%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%83%ad%e9%83%a8%e7%bd%b2%e5%92%8c%e7%83%ad%e5%8a%a0%e8%bd%bd%ef%bc%9f.md">23 Host容器：Tomcat如何实现热部署和热加载？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 Context容器（上）：Tomcat如何打破双亲委托机制？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/24%20Context%e5%ae%b9%e5%99%a8%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e6%89%93%e7%a0%b4%e5%8f%8c%e4%ba%b2%e5%a7%94%e6%89%98%e6%9c%ba%e5%88%b6%ef%bc%9f.md">24 Context容器（上）：Tomcat如何打破双亲委托机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 Context容器（中）：Tomcat如何隔离Web应用？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/25%20Context%e5%ae%b9%e5%99%a8%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e9%9a%94%e7%a6%bbWeb%e5%ba%94%e7%94%a8%ef%bc%9f.md">25 Context容器（中）：Tomcat如何隔离Web应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 Context容器（下）：Tomcat如何实现Servlet规范？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/26%20Context%e5%ae%b9%e5%99%a8%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0Servlet%e8%a7%84%e8%8c%83%ef%bc%9f.md">26 Context容器（下）：Tomcat如何实现Servlet规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 新特性：Tomcat如何支持异步Servlet？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/27%20%e6%96%b0%e7%89%b9%e6%80%a7%ef%bc%9aTomcat%e5%a6%82%e4%bd%95%e6%94%af%e6%8c%81%e5%bc%82%e6%ad%a5Servlet%ef%bc%9f.md">27 新特性：Tomcat如何支持异步Servlet？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 新特性：Spring Boot如何使用内嵌式的Tomcat和Jetty？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/28%20%e6%96%b0%e7%89%b9%e6%80%a7%ef%bc%9aSpring%20Boot%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%86%85%e5%b5%8c%e5%bc%8f%e7%9a%84Tomcat%e5%92%8cJetty%ef%bc%9f.md">28 新特性：Spring Boot如何使用内嵌式的Tomcat和Jetty？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 比较：Jetty如何实现具有上下文信息的责任链？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/29%20%e6%af%94%e8%be%83%ef%bc%9aJetty%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%85%b7%e6%9c%89%e4%b8%8a%e4%b8%8b%e6%96%87%e4%bf%a1%e6%81%af%e7%9a%84%e8%b4%a3%e4%bb%bb%e9%93%be%ef%bc%9f.md">29 比较：Jetty如何实现具有上下文信息的责任链？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 热点问题答疑（3）：Spring框架中的设计模式.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/30%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%883%ef%bc%89%ef%bc%9aSpring%e6%a1%86%e6%9e%b6%e4%b8%ad%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f.md">30 热点问题答疑（3）：Spring框架中的设计模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 Logger组件：Tomcat的日志框架及实战.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/31%20Logger%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e7%9a%84%e6%97%a5%e5%bf%97%e6%a1%86%e6%9e%b6%e5%8f%8a%e5%ae%9e%e6%88%98.md">31 Logger组件：Tomcat的日志框架及实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 Manager组件：Tomcat的Session管理机制解析.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/32%20Manager%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e7%9a%84Session%e7%ae%a1%e7%90%86%e6%9c%ba%e5%88%b6%e8%a7%a3%e6%9e%90.md">32 Manager组件：Tomcat的Session管理机制解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 Cluster组件：Tomcat的集群通信原理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/33%20Cluster%e7%bb%84%e4%bb%b6%ef%bc%9aTomcat%e7%9a%84%e9%9b%86%e7%be%a4%e9%80%9a%e4%bf%a1%e5%8e%9f%e7%90%86.md">33 Cluster组件：Tomcat的集群通信原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 JVM GC原理及调优的基本思路.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/34%20JVM%20GC%e5%8e%9f%e7%90%86%e5%8f%8a%e8%b0%83%e4%bc%98%e7%9a%84%e5%9f%ba%e6%9c%ac%e6%80%9d%e8%b7%af.md">34 JVM GC原理及调优的基本思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 如何监控Tomcat的性能？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/35%20%e5%a6%82%e4%bd%95%e7%9b%91%e6%8e%a7Tomcat%e7%9a%84%e6%80%a7%e8%83%bd%ef%bc%9f.md">35 如何监控Tomcat的性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Tomcat I_O和线程池的并发调优.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/36%20Tomcat%20I_O%e5%92%8c%e7%ba%bf%e7%a8%8b%e6%b1%a0%e7%9a%84%e5%b9%b6%e5%8f%91%e8%b0%83%e4%bc%98.md">36 Tomcat I_O和线程池的并发调优.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 Tomcat内存溢出的原因分析及调优.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/37%20Tomcat%e5%86%85%e5%ad%98%e6%ba%a2%e5%87%ba%e7%9a%84%e5%8e%9f%e5%9b%a0%e5%88%86%e6%9e%90%e5%8f%8a%e8%b0%83%e4%bc%98.md">37 Tomcat内存溢出的原因分析及调优.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 Tomcat拒绝连接原因分析及网络优化.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/38%20Tomcat%e6%8b%92%e7%bb%9d%e8%bf%9e%e6%8e%a5%e5%8e%9f%e5%9b%a0%e5%88%86%e6%9e%90%e5%8f%8a%e7%bd%91%e7%bb%9c%e4%bc%98%e5%8c%96.md">38 Tomcat拒绝连接原因分析及网络优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 Tomcat进程占用CPU过高怎么办？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/39%20Tomcat%e8%bf%9b%e7%a8%8b%e5%8d%a0%e7%94%a8CPU%e8%bf%87%e9%ab%98%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">39 Tomcat进程占用CPU过高怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 谈谈Jetty性能调优的思路.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/40%20%e8%b0%88%e8%b0%88Jetty%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e7%9a%84%e6%80%9d%e8%b7%af.md">40 谈谈Jetty性能调优的思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 热点问题答疑（4）： Tomcat和Jetty有哪些不同？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/41%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%884%ef%bc%89%ef%bc%9a%20Tomcat%e5%92%8cJetty%e6%9c%89%e5%93%aa%e4%ba%9b%e4%b8%8d%e5%90%8c%ef%bc%9f.md">41 热点问题答疑（4）： Tomcat和Jetty有哪些不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 如何持续保持对学习的兴趣？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e5%a6%82%e4%bd%95%e6%8c%81%e7%bb%ad%e4%bf%9d%e6%8c%81%e5%af%b9%e5%ad%a6%e4%b9%a0%e7%9a%84%e5%85%b4%e8%b6%a3%ef%bc%9f.md">特别放送 如何持续保持对学习的兴趣？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 静下心来，品味经典.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%8b%86%e8%a7%a3Tomcat%20%20Jetty/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e9%9d%99%e4%b8%8b%e5%bf%83%e6%9d%a5%ef%bc%8c%e5%93%81%e5%91%b3%e7%bb%8f%e5%85%b8.md">结束语 静下心来，品味经典.md</a>
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
                            <h1 id="title" data-id="35 如何监控Tomcat的性能？" class="title">35 如何监控Tomcat的性能？</h1>
                            <div><p>专栏上一期我们分析了JVM GC的基本原理以及监控和分析工具，今天我们接着来聊如何监控Tomcat的各种指标，因为只有我们掌握了这些指标和信息，才能对Tomcat内部发生的事情一目了然，让我们明白系统的瓶颈在哪里，进而做出调优的决策。</p>

<p>在今天的文章里，我们首先来看看到底都需要监控Tomcat哪些关键指标，接着来具体学习如何通过JConsole来监控它们。如果系统没有暴露JMX接口，我们还可以通过命令行来查看Tomcat的性能指标。</p>

<p>Web应用的响应时间是我们关注的一个重点，最后我们通过一个实战案例，来看看Web应用的下游服务响应时间比较长的情况下，Tomcat的各项指标是什么样子的。</p>

<h2 id="tomcat的关键指标">Tomcat的关键指标</h2>

<p>Tomcat的关键指标有<strong>吞吐量、响应时间、错误数、线程池、CPU以及JVM内存</strong>。</p>

<p>我来简单介绍一下这些指标背后的意义。其中前三个指标是我们最关心的业务指标，Tomcat作为服务器，就是要能够又快有好地处理请求，因此吞吐量要大、响应时间要短，并且错误数要少。</p>

<p>而后面三个指标是跟系统资源有关的，当某个资源出现瓶颈就会影响前面的业务指标，比如线程池中的线程数量不足会影响吞吐量和响应时间；但是线程数太多会耗费大量CPU，也会影响吞吐量；当内存不足时会触发频繁地GC，耗费CPU，最后也会反映到业务指标上来。</p>

<p>那如何监控这些指标呢？Tomcat可以通过JMX将上述指标暴露出来的。JMX（Java Management Extensions，即Java管理扩展）是一个为应用程序、设备、系统等植入监控管理功能的框架。JMX使用管理MBean来监控业务资源，这些MBean在JMX MBean服务器上注册，代表JVM中运行的应用程序或服务。每个MBean都有一个属性列表。JMX客户端可以连接到MBean Server来读写MBean的属性值。你可以通过下面这张图来理解一下JMX的工作原理：</p>

<p><img src="assets/ba04d8418bf14995beb63148d209dd95.jpg" alt="" /></p>

<p>Tomcat定义了一系列MBean来对外暴露系统状态，接下来我们来看看如何通过JConsole来监控这些指标。</p>

<h2 id="通过jconsole监控tomcat">通过JConsole监控Tomcat</h2>

<p>首先我们需要开启JMX的远程监听端口，具体来说就是设置若干JVM参数。我们可以在Tomcat的bin目录下新建一个名为<code>setenv.sh</code>的文件（或者<code>setenv.bat</code>，根据你的操作系统类型），然后输入下面的内容：</p>

<pre><code>export JAVA_OPTS=&quot;${JAVA_OPTS} -Dcom.sun.management.jmxremote&quot;
export JAVA_OPTS=&quot;${JAVA_OPTS} -Dcom.sun.management.jmxremote.port=9001&quot;
export JAVA_OPTS=&quot;${JAVA_OPTS} -Djava.rmi.server.hostname=x.x.x.x&quot;
export JAVA_OPTS=&quot;${JAVA_OPTS} -Dcom.sun.management.jmxremote.ssl=false&quot;
export JAVA_OPTS=&quot;${JAVA_OPTS} -Dcom.sun.management.jmxremote.authenticate=false&quot;
</code></pre>

<p>重启Tomcat，这样JMX的监听端口9001就开启了，接下来通过JConsole来连接这个端口。</p>

<pre><code>jconsole x.x.x.x:9001
</code></pre>

<p>我们可以看到JConsole的主界面：</p>

<p><img src="assets/32f4f1bca7534022bdafbc5a63e5262d.jpg" alt="" /></p>

<p>前面我提到的需要监控的关键指标有<strong>吞吐量、响应时间、错误数、线程池、CPU以及JVM内存</strong>，接下来我们就来看看怎么在JConsole上找到这些指标。</p>

<p><strong>吞吐量、响应时间、错误数</strong></p>

<p>在MBeans标签页下选择GlobalRequestProcessor，这里有Tomcat请求处理的统计信息。你会看到Tomcat中的各种连接器，展开“http-nio-8080”，你会看到这个连接器上的统计信息，其中maxTime表示最长的响应时间，processingTime表示平均响应时间，requestCount表示吞吐量，errorCount就是错误数。</p>

<p><img src="assets/7cff80831cc9486884e086ac42d3ff7e.jpg" alt="" /></p>

<p><strong>线程池</strong></p>

<p>选择“线程”标签页，可以看到当前Tomcat进程中有多少线程，如下图所示：</p>

<p><img src="assets/7f49c6e34e864382834cc97d3a60f55b.jpg" alt="" /></p>

<p>图的左下方是线程列表，右边是线程的运行栈，这些都是非常有用的信息。如果大量线程阻塞，通过观察线程栈，能看到线程阻塞在哪个函数，有可能是I/O等待，或者是死锁。</p>

<p><strong>CPU</strong></p>

<p>在主界面可以找到CPU使用率指标，请注意这里的CPU使用率指的是Tomcat进程占用的CPU，不是主机总的CPU使用率。</p>

<p><img src="assets/9ce5607c19e94bb686836a15321c90fb.jpg" alt="" /></p>

<p><strong>JVM内存</strong></p>

<p>选择“内存”标签页，你能看到Tomcat进程的JVM内存使用情况。</p>

<p><img src="assets/9cc6db06db5144748356134276c27af3.jpg" alt="" /></p>

<p>你还可以查看JVM各内存区域的使用情况，大的层面分堆区和非堆区。堆区里有分为Eden、Survivor和Old。选择“VM Summary”标签，可以看到虚拟机内的详细信息。</p>

<p><img src="assets/63399aa296064ffbb4f2f40456a46026.jpg" alt="" /></p>

<h2 id="命令行查看tomcat指标">命令行查看Tomcat指标</h2>

<p>极端情况下如果Web应用占用过多CPU或者内存，又或者程序中发生了死锁，导致Web应用对外没有响应，监控系统上看不到数据，这个时候需要我们登陆到目标机器，通过命令行来查看各种指标。</p>

<p>1.首先我们通过ps命令找到Tomcat进程，拿到进程ID。</p>

<p><img src="assets/efb5812b1e5449bc854de1298c3662c2.jpg" alt="" /></p>

<p>2.接着查看进程状态的大致信息，通过<code>cat/proc/&lt;pid&gt;/status</code>命令：</p>

<p><img src="assets/4d72ae5a23494a5b883a50c4c8d619d7.jpg" alt="" /></p>

<p>3.监控进程的CPU和内存资源使用情况：</p>

<p><img src="assets/8cd9274e317e4040a16bb97a8ed74420.jpg" alt="" /></p>

<p>4.查看Tomcat的网络连接，比如Tomcat在8080端口上监听连接请求，通过下面的命令查看连接列表：</p>

<p><img src="assets/df139429ce3c424cb4e70de4240f4824.jpg" alt="" /></p>

<p>你还可以分别统计处在“已连接”状态和“TIME_WAIT”状态的连接数：</p>

<p><img src="assets/254a8bdac83244edbad7293adc732562.jpg" alt="" /></p>

<p>5.通过ifstat来查看网络流量，大致可以看出Tomcat当前的请求数和负载状况。</p>

<p><img src="assets/eff2640e671948d189b226a3513719d9.jpg" alt="" /></p>

<h2 id="实战案例">实战案例</h2>

<p>在这个实战案例中，我们会创建一个Web应用，根据传入的参数latency来休眠相应的秒数，目的是模拟当前的Web应用在访问下游服务时遇到的延迟。然后用JMeter来压测这个服务，通过JConsole来观察Tomcat的各项指标，分析和定位问题。</p>

<p>主要的步骤有：</p>

<p>1.创建一个Spring Boot程序，加入下面代码所示的一个RestController：</p>

<pre><code>@RestController
public class DownStreamLatency {

    @RequestMapping(&quot;/greeting/latency/{seconds}&quot;)
    public Greeting greeting(@PathVariable long seconds) {

        try {
            Thread.sleep(seconds * 1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        Greeting greeting = new Greeting(&quot;Hello World!&quot;);

        return greeting;
    }
}
</code></pre>

<p>从上面的代码我们看到，程序会读取URL传过来的seconds参数，先休眠相应的秒数，再返回请求。这样做的目的是，客户端压测工具能够控制服务端的延迟。</p>

<p>为了方便观察Tomcat的线程数跟延迟之间的关系，还需要加大Tomcat的最大线程数，我们可以在<code>application.properties</code>文件中加入这样一行：</p>

<pre><code>server.tomcat.max-threads=1000server.tomcat.max-threads=1000
</code></pre>

<p>2.启动JMeter开始压测，这里我们将压测的线程数设置为100：</p>

<p><img src="assets/15cd5593291b4afdb85ccc334f4234ea.jpg" alt="" /></p>

<p>请你注意的是，我们还需要将客户端的Timeout设置为1000毫秒，这是因为JMeter的测试线程在收到响应之前，不会发出下一次请求，这就意味我们没法按照固定的吞吐量向服务端加压。而加了Timeout以后，JMeter会有固定的吞吐量向Tomcat发送请求。</p>

<p><img src="assets/2b254711e13c4be999a580545c2a845a.jpg" alt="" /></p>

<p>3.开启测试，这里分三个阶段，第一个阶段将服务端休眠时间设为2秒，然后暂停一段时间。第二和第三阶段分别将休眠时间设置成4秒和6秒。</p>

<p><img src="assets/8e15d9cfa64149939d534d0365f30a8a.jpg" alt="" /></p>

<p>4.最后我们通过JConsole来观察结果：</p>

<p><img src="assets/7382146bba144c0bbee59a9428e6b286.jpg" alt="" /></p>

<p>下面我们从线程数、内存和CPU这三个指标来分析Tomcat的性能问题。</p>

<ul>
<li>首先看线程数，在第一阶段时间之前，线程数大概是40，第一阶段压测开始后，线程数增长到250。为什么是250呢？这是因为JMeter每秒会发出100个请求，每一个请求休眠2秒，因此Tomcat需要200个工作线程来干活；此外Tomcat还有一些其他线程用来处理网络通信和后台任务，所以总数是250左右。第一阶段压测暂停后，线程数又下降到40，这是因为线程池会回收空闲线程。第二阶段测试开始后，线程数涨到了420，这是因为每个请求休眠了4秒；同理，我们看到第三阶段测试的线程数是620。</li>
<li>我们再来看CPU，在三个阶段的测试中，CPU的峰值始终比较稳定，这是因为JMeter控制了总体的吞吐量，因为服务端用来处理这些请求所需要消耗的CPU基本也是一样的。</li>
<li>各测试阶段的内存使用量略有增加，这是因为线程数增加了，创建线程也需要消耗内存。</li>
</ul>

<p>从上面的测试结果我们可以得出一个结论：对于一个Web应用来说，下游服务的延迟越大，Tomcat所需要的线程数越多，但是CPU保持稳定。所以如果你在实际工作碰到线程数飙升但是CPU没有增加的情况，这个时候你需要怀疑，你的Web应用所依赖的下游服务是不是出了问题，响应时间是否变长了。</p>

<h2 id="本期精华">本期精华</h2>

<p>今天我们学习了Tomcat中的关键的性能指标以及如何监控这些指标：主要有<strong>吞吐量、响应时间、错误数、线程池、CPU以及JVM内存。</strong></p>

<p>在实际工作中，我们需要通过观察这些指标来诊断系统遇到的性能问题，找到性能瓶颈。如果我们监控到CPU上升，这时我们可以看看吞吐量是不是也上升了，如果是那说明正常；如果不是的话，可以看看GC的活动，如果GC活动频繁，并且内存居高不下，基本可以断定是内存泄漏。</p>

<h2 id="课后思考">课后思考</h2>

<p>请问工作中你如何监控Web应用的健康状态？遇到性能问题的时候是如何做问题定位的呢？</p>

<p>不知道今天的内容你消化得如何？如果还有疑问，请大胆的在留言区提问，也欢迎你把你的课后思考和心得记录下来，与我和其他同学一起讨论。如果你觉得今天有所收获，欢迎你把它分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#036f6f6f3a373232333443646e626a6f2d606c6e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f15b3b5dd647771',t:'MTczNDA5MDI2NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
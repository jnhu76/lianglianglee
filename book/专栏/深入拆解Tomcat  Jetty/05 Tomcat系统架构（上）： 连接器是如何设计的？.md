<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=05&#32;Tomcat系统架构（上）： 连接器是如何设计的？>
        <link rel="icon" href="/static/favicon.png">
        <title>05 Tomcat系统架构（上）： 连接器是如何设计的？ </title>
        
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
                            <h1 id="title" data-id="05 Tomcat系统架构（上）： 连接器是如何设计的？" class="title">05 Tomcat系统架构（上）： 连接器是如何设计的？</h1>
                            <div><p>05 Tomcat系统架构（上）： 连接器是如何设计的？</p>

<p>在面试时我们可能经常被问到：你做的XX项目的架构是如何设计的，请讲一下实现的思路。对于面试官来说，可以通过你对复杂系统设计的理解，了解你的技术水平以及处理复杂问题的思路。</p>

<p>今天咱们就来一步一步分析Tomcat的设计思路，看看Tomcat的设计者们当时是怎么回答这个问题的。一方面我们可以学到Tomcat的总体架构，学会从宏观上怎么去设计一个复杂系统，怎么设计顶层模块，以及模块之间的关系；另一方面也为我们深入学习Tomcat的工作原理打下基础。</p>

<h2 id="tomcat总体架构">Tomcat总体架构</h2>

<p>我们知道如果要设计一个系统，首先是要了解需求。通过专栏前面的文章，我们已经了解了Tomcat要实现2个核心功能：</p>

<ul>
<li><p>处理Socket连接，负责网络字节流与Request和Response对象的转化。</p></li>

<li><p>加载和管理Servlet，以及具体处理Request请求。</p></li>
</ul>

<p><strong>因此Tomcat设计了两个核心组件连接器（Connector）和容器（Container）来分别做这两件事情。连接器负责对外交流，容器负责内部处理。</strong></p>

<p>所以连接器和容器可以说是Tomcat架构里最重要的两部分，需要你花些精力理解清楚。这两部分内容我会分成两期，今天我来分析连接器是如何设计的，下一期我会介绍容器的设计。</p>

<p>在开始讲连接器前，我先铺垫一下Tomcat支持的多种I/O模型和应用层协议。</p>

<p>Tomcat支持的I/O模型有：</p>

<ul>
<li><p>NIO：非阻塞I/O，采用Java NIO类库实现。</p></li>

<li><p>NIO.2：异步I/O，采用JDK 7最新的NIO.2类库实现。</p></li>

<li><p>APR：采用Apache可移植运行库实现，是C/C++编写的本地库。</p></li>
</ul>

<p>Tomcat支持的应用层协议有：</p>

<ul>
<li><p>HTTP/1.1：这是大部分Web应用采用的访问协议。</p></li>

<li><p>AJP：用于和Web服务器集成（如Apache）。</p></li>

<li><p>HTTP/2：HTTP 2.0大幅度的提升了Web性能。</p></li>
</ul>

<p>Tomcat为了实现支持多种I/O模型和应用层协议，一个容器可能对接多个连接器，就好比一个房间有多个门。但是单独的连接器或者容器都不能对外提供服务，需要把它们组装起来才能工作，组装后这个整体叫作Service组件。这里请你注意，Service本身没有做什么重要的事情，只是在连接器和容器外面多包了一层，把它们组装在一起。Tomcat内可能有多个Service，这样的设计也是出于灵活性的考虑。通过在Tomcat中配置多个Service，可以实现通过不同的端口号来访问同一台机器上部署的不同应用。</p>

<p>到此我们得到这样一张关系图：</p>

<p><img src="assets/7b0377be2dde4696be981e7691b982f1.jpg" alt="" /></p>

<p>从图上你可以看到，最顶层是Server，这里的Server指的就是一个Tomcat实例。一个Server中有一个或者多个Service，一个Service中有多个连接器和一个容器。连接器与容器之间通过标准的ServletRequest和ServletResponse通信。</p>

<h2 id="连接器">连接器</h2>

<p>连接器对Servlet容器屏蔽了协议及I/O模型等的区别，无论是HTTP还是AJP，在容器中获取到的都是一个标准的ServletRequest对象。</p>

<p>我们可以把连接器的功能需求进一步细化，比如：</p>

<ul>
<li><p>监听网络端口。</p></li>

<li><p>接受网络连接请求。</p></li>

<li><p>读取网络请求字节流。</p></li>

<li><p>根据具体应用层协议（HTTP/AJP）解析字节流，生成统一的Tomcat Request对象。</p></li>

<li><p>将Tomcat Request对象转成标准的ServletRequest。</p></li>

<li><p>调用Servlet容器，得到ServletResponse。</p></li>

<li><p>将ServletResponse转成Tomcat Response对象。</p></li>

<li><p>将Tomcat Response转成网络字节流。</p></li>

<li><p>将响应字节流写回给浏览器。</p></li>
</ul>

<p>需求列清楚后，我们要考虑的下一个问题是，连接器应该有哪些子模块？优秀的模块化设计应该考虑<strong>高内聚、低耦合</strong>。</p>

<ul>
<li><p><strong>高内聚</strong>是指相关度比较高的功能要尽可能集中，不要分散。</p></li>

<li><p><strong>低耦合</strong>是指两个相关的模块要尽可能减少依赖的部分和降低依赖的程度，不要让两个模块产生强依赖。</p></li>
</ul>

<p>通过分析连接器的详细功能列表，我们发现连接器需要完成3个<strong>高内聚</strong>的功能：</p>

<ul>
<li><p>网络通信。</p></li>

<li><p>应用层协议解析。</p></li>

<li><p>Tomcat Request/Response与ServletRequest/ServletResponse的转化。</p></li>
</ul>

<p>因此Tomcat的设计者设计了3个组件来实现这3个功能，分别是Endpoint、Processor和Adapter。</p>

<p>组件之间通过抽象接口交互。这样做还有一个好处是<strong>封装变化。</strong>这是面向对象设计的精髓，将系统中经常变化的部分和稳定的部分隔离，有助于增加复用性，并降低系统耦合度。</p>

<p>网络通信的I/O模型是变化的，可能是非阻塞I/O、异步I/O或者APR。应用层协议也是变化的，可能是HTTP、HTTPS、AJP。浏览器端发送的请求信息也是变化的。</p>

<p>但是整体的处理逻辑是不变的，Endpoint负责提供字节流给Processor，Processor负责提供Tomcat Request对象给Adapter，Adapter负责提供ServletRequest对象给容器。</p>

<p>如果要支持新的I/O方案、新的应用层协议，只需要实现相关的具体子类，上层通用的处理逻辑是不变的。</p>

<p>由于I/O模型和应用层协议可以自由组合，比如NIO + HTTP或者NIO.2 + AJP。Tomcat的设计者将网络通信和应用层协议解析放在一起考虑，设计了一个叫ProtocolHandler的接口来封装这两种变化点。各种协议和通信模型的组合有相应的具体实现类。比如：Http11NioProtocol和AjpNioProtocol。</p>

<p>除了这些变化点，系统也存在一些相对稳定的部分，因此Tomcat设计了一系列抽象基类来<strong>封装这些稳定的部分</strong>，抽象基类AbstractProtocol实现了ProtocolHandler接口。每一种应用层协议有自己的抽象基类，比如AbstractAjpProtocol和AbstractHttp11Protocol，具体协议的实现类扩展了协议层抽象基类。下面我整理一下它们的继承关系。</p>

<p><img src="assets/01abbc34b3944300ac81a847b6e15b10.jpg" alt="" /></p>

<p>通过上面的图，你可以清晰地看到它们的继承和层次关系，这样设计的目的是尽量将稳定的部分放到抽象基类，同时每一种I/O模型和协议的组合都有相应的具体实现类，我们在使用时可以自由选择。</p>

<p>小结一下，连接器模块用三个核心组件：Endpoint、Processor和Adapter来分别做三件事情，其中Endpoint和Processor放在一起抽象成了ProtocolHandler组件，它们的关系如下图所示。</p>

<p><img src="assets/887abb917136488aab259e8919a0bcb4.jpg" alt="" /></p>

<p>下面我来详细介绍这两个顶层组件ProtocolHandler和Adapter。</p>

<p><strong>ProtocolHandler组件</strong></p>

<p>由上文我们知道，连接器用ProtocolHandler来处理网络连接和应用层协议，包含了2个重要部件：Endpoint和Processor，下面我来详细介绍它们的工作原理。</p>

<ul>
<li>Endpoint</li>
</ul>

<p>Endpoint是通信端点，即通信监听的接口，是具体的Socket接收和发送处理器，是对传输层的抽象，因此Endpoint是用来实现TCP/IP协议的。</p>

<p>Endpoint是一个接口，对应的抽象实现类是AbstractEndpoint，而AbstractEndpoint的具体子类，比如在NioEndpoint和Nio2Endpoint中，有两个重要的子组件：Acceptor和SocketProcessor。</p>

<p>其中Acceptor用于监听Socket连接请求。SocketProcessor用于处理接收到的Socket请求，它实现Runnable接口，在run方法里调用协议处理组件Processor进行处理。为了提高处理能力，SocketProcessor被提交到线程池来执行。而这个线程池叫作执行器（Executor)，我在后面的专栏会详细介绍Tomcat如何扩展原生的Java线程池。</p>

<ul>
<li>Processor</li>
</ul>

<p>如果说Endpoint是用来实现TCP/IP协议的，那么Processor用来实现HTTP协议，Processor接收来自Endpoint的Socket，读取字节流解析成Tomcat Request和Response对象，并通过Adapter将其提交到容器处理，Processor是对应用层协议的抽象。</p>

<p>Processor是一个接口，定义了请求的处理等方法。它的抽象实现类AbstractProcessor对一些协议共有的属性进行封装，没有对方法进行实现。具体的实现有AjpProcessor、Http11Processor等，这些具体实现类实现了特定协议的解析方法和请求处理方式。</p>

<p>我们再来看看连接器的组件图：</p>

<p><img src="assets/fa6d757632aa419fb3486afa166323d3.jpg" alt="" /></p>

<p>从图中我们看到，Endpoint接收到Socket连接后，生成一个SocketProcessor任务提交到线程池去处理，SocketProcessor的run方法会调用Processor组件去解析应用层协议，Processor通过解析生成Request对象后，会调用Adapter的Service方法。</p>

<p>到这里我们学习了ProtocolHandler的总体架构和工作原理，关于Endpoint的详细设计，后面我还会专门介绍Endpoint是如何最大限度地利用Java NIO的非阻塞以及NIO.2的异步特性，来实现高并发。</p>

<p><strong>Adapter组件</strong></p>

<p>我在前面说过，由于协议不同，客户端发过来的请求信息也不尽相同，Tomcat定义了自己的Request类来“存放”这些请求信息。ProtocolHandler接口负责解析请求并生成Tomcat Request类。但是这个Request对象不是标准的ServletRequest，也就意味着，不能用Tomcat Request作为参数来调用容器。Tomcat设计者的解决方案是引入CoyoteAdapter，这是适配器模式的经典运用，连接器调用CoyoteAdapter的sevice方法，传入的是Tomcat Request对象，CoyoteAdapter负责将Tomcat Request转成ServletRequest，再调用容器的service方法。</p>

<h2 id="本期精华">本期精华</h2>

<p>Tomcat的整体架构包含了两个核心组件连接器和容器。连接器负责对外交流，容器负责内部处理。连接器用ProtocolHandler接口来封装通信协议和I/O模型的差异，ProtocolHandler内部又分为Endpoint和Processor模块，Endpoint负责底层Socket通信，Processor负责应用层协议解析。连接器通过适配器Adapter调用容器。</p>

<p>通过对Tomcat整体架构的学习，我们可以得到一些设计复杂系统的基本思路。首先要分析需求，根据高内聚低耦合的原则确定子模块，然后找出子模块中的变化点和不变点，用接口和抽象基类去封装不变点，在抽象基类中定义模板方法，让子类自行实现抽象方法，也就是具体子类去实现变化点。</p>

<h2 id="课后思考">课后思考</h2>

<p>回忆一下你在工作中曾经独立设计过的系统，或者你碰到过的设计类面试题，结合今天专栏的内容，你有没有一些新的思路？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#82eeeeeebbb6b3b3b2b5c2e5efe3ebeeace1edef" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f15aad2ba517771',t:'MTczNDA4OTg5OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=19&#32;比较：Jetty的线程策略EatWhatYouKill>
        <link rel="icon" href="/static/favicon.png">
        <title>19 比较：Jetty的线程策略EatWhatYouKill </title>
        
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
                            <h1 id="title" data-id="19 比较：Jetty的线程策略EatWhatYouKill" class="title">19 比较：Jetty的线程策略EatWhatYouKill</h1>
                            <div><p>我在前面的专栏里介绍了Jetty的总体架构设计，简单回顾一下，Jetty总体上是由一系列Connector、一系列Handler和一个ThreadPool组成，它们的关系如下图所示：</p>

<p><img src="assets/adde0c3870d847b6adfe5e2c239fe3ef.jpg" alt="" /></p>

<p>相比较Tomcat的连接器，Jetty的Connector在设计上有自己的特点。Jetty的Connector支持NIO通信模型，我们知道<strong>NIO模型中的主角就是Selector</strong>，Jetty在Java原生Selector的基础上封装了自己的Selector，叫作ManagedSelector。ManagedSelector在线程策略方面做了大胆尝试，将I/O事件的侦测和处理放到同一个线程来处理，充分利用了CPU缓存并减少了线程上下文切换的开销。</p>

<p>具体的数字是，根据Jetty的官方测试，这种名为“EatWhatYouKill”的线程策略将吞吐量提高了8倍。你一定很好奇它是如何实现的吧，今天我们就来看一看这背后的原理是什么。</p>

<h2 id="selector编程的一般思路">Selector编程的一般思路</h2>

<p>常规的NIO编程思路是，将I/O事件的侦测和请求的处理分别用不同的线程处理。具体过程是：</p>

<p>启动一个线程，在一个死循环里不断地调用select方法，检测Channel的I/O状态，一旦I/O事件达到，比如数据就绪，就把该I/O事件以及一些数据包装成一个Runnable，将Runnable放到新线程中去处理。</p>

<p>在这个过程中按照职责划分，有两个线程在干活，一个是I/O事件检测线程，另一个是I/O事件处理线程。我们仔细思考一下这两者的关系，其实它们是生产者和消费者的关系。I/O事件侦测线程作为生产者，负责“生产”I/O事件，也就是负责接活儿的老板；I/O处理线程是消费者，它“消费”并处理I/O事件，就是干苦力的员工。把这两个工作用不同的线程来处理，好处是它们互不干扰和阻塞对方。</p>

<h2 id="jetty中的selector编程">Jetty中的Selector编程</h2>

<p>然而世事无绝对，将I/O事件检测和业务处理这两种工作分开的思路也有缺点。当Selector检测读就绪事件时，数据已经被拷贝到内核中的缓存了，同时CPU的缓存中也有这些数据了，我们知道CPU本身的缓存比内存快多了，这时当应用程序去读取这些数据时，如果用另一个线程去读，很有可能这个读线程使用另一个CPU核，而不是之前那个检测数据就绪的CPU核，这样CPU缓存中的数据就用不上了，并且线程切换也需要开销。</p>

<p>因此Jetty的Connector做了一个大胆尝试，那就是用<strong>把I/O事件的生产和消费放到同一个线程来处理</strong>，如果这两个任务由同一个线程来执行，如果执行过程中线程不阻塞，操作系统会用同一个CPU核来执行这两个任务，这样就能利用CPU缓存了。那具体是如何做的呢，我们还是来详细分析一下Connector中的ManagedSelector组件。</p>

<p><strong>ManagedSelector</strong></p>

<p>ManagedSelector的本质就是一个Selector，负责I/O事件的检测和分发。为了方便使用，Jetty在Java原生的Selector上做了一些扩展，就变成了ManagedSelector，我们先来看看它有哪些成员变量：</p>

<pre><code>public class ManagedSelector extends ContainerLifeCycle implements Dumpable
{
    //原子变量，表明当前的ManagedSelector是否已经启动
    private final AtomicBoolean _started = new AtomicBoolean(false);
    
    //表明是否阻塞在select调用上
    private boolean _selecting = false;
    
    //管理器的引用，SelectorManager管理若干ManagedSelector的生命周期
    private final SelectorManager _selectorManager;
    
    //ManagedSelector不止一个，为它们每人分配一个id
    private final int _id;
    
    //关键的执行策略，生产者和消费者是否在同一个线程处理由它决定
    private final ExecutionStrategy _strategy;
    
    //Java原生的Selector
    private Selector _selector;
    
    //&quot;Selector更新任务&quot;队列
    private Deque&lt;SelectorUpdate&gt; _updates = new ArrayDeque&lt;&gt;();
    private Deque&lt;SelectorUpdate&gt; _updateable = new ArrayDeque&lt;&gt;();
    
    ...
}
</code></pre>

<p>这些成员变量中其他的都好理解，就是“Selector更新任务”队列<code>_updates</code>和执行策略<code>_strategy</code>可能不是很直观。</p>

<p><strong>SelectorUpdate接口</strong></p>

<p>为什么需要一个“Selector更新任务”队列呢，对于Selector的用户来说，我们对Selector的操作无非是将Channel注册到Selector或者告诉Selector我对什么I/O事件感兴趣，那么这些操作其实就是对Selector状态的更新，Jetty把这些操作抽象成SelectorUpdate接口。</p>

<pre><code>/**
 * A selector update to be done when the selector has been woken.
 */
public interface SelectorUpdate
{
    void update(Selector selector);
}
</code></pre>

<p>这意味着如果你不能直接操作ManageSelector中的Selector，而是需要向ManagedSelector提交一个任务类，这个类需要实现SelectorUpdate接口update方法，在update方法里定义你想要对ManagedSelector做的操作。</p>

<p>比如Connector中Endpoint组件对读就绪事件感兴趣，它就向ManagedSelector提交了一个内部任务类ManagedSelector.SelectorUpdate：</p>

<pre><code>_selector.submit(_updateKeyAction);
</code></pre>

<p>这个<code>_updateKeyAction</code>就是一个SelectorUpdate实例，它的update方法实现如下：</p>

<pre><code>private final ManagedSelector.SelectorUpdate _updateKeyAction = new ManagedSelector.SelectorUpdate()
{
    @Override
    public void update(Selector selector)
    {
        //这里的updateKey其实就是调用了SelectionKey.interestOps(OP_READ);
        updateKey();
    }
};
</code></pre>

<p>我们看到在update方法里，调用了SelectionKey类的interestOps方法，传入的参数是<code>OP_READ</code>，意思是现在我对这个Channel上的读就绪事件感兴趣了。</p>

<p>那谁来负责执行这些update方法呢，答案是ManagedSelector自己，它在一个死循环里拉取这些SelectorUpdate任务类逐个执行。</p>

<p><strong>Selectable接口</strong></p>

<p>那I/O事件到达时，ManagedSelector怎么知道应该调哪个函数来处理呢？其实也是通过一个任务类接口，这个接口就是Selectable，它返回一个Runnable，这个Runnable其实就是I/O事件就绪时相应的处理逻辑。</p>

<pre><code>public interface Selectable
{
    //当某一个Channel的I/O事件就绪后，ManagedSelector会调用的回调函数
    Runnable onSelected();

    //当所有事件处理完了之后ManagedSelector会调的回调函数，我们先忽略。
    void updateKey();
}
</code></pre>

<p>ManagedSelector在检测到某个Channel上的I/O事件就绪时，也就是说这个Channel被选中了，ManagedSelector调用这个Channel所绑定的附件类的onSelected方法来拿到一个Runnable。</p>

<p>这句话有点绕，其实就是ManagedSelector的使用者，比如Endpoint组件在向ManagedSelector注册读就绪事件时，同时也要告诉ManagedSelector在事件就绪时执行什么任务，具体来说就是传入一个附件类，这个附件类需要实现Selectable接口。ManagedSelector通过调用这个onSelected拿到一个Runnable，然后把Runnable扔给线程池去执行。</p>

<p>那Endpoint的onSelected是如何实现的呢？</p>

<pre><code>@Override
public Runnable onSelected()
{
    int readyOps = _key.readyOps();

    boolean fillable = (readyOps &amp; SelectionKey.OP_READ) != 0;
    boolean flushable = (readyOps &amp; SelectionKey.OP_WRITE) != 0;

    // return task to complete the job
    Runnable task= fillable 
            ? (flushable 
                    ? _runCompleteWriteFillable 
                    : _runFillable)
            : (flushable 
                    ? _runCompleteWrite 
                    : null);

    return task;
}
</code></pre>

<p>上面的代码逻辑很简单，就是读事件到了就读，写事件到了就写。</p>

<p><strong>ExecutionStrategy</strong></p>

<p>铺垫了这么多，终于要上主菜了。前面我主要介绍了ManagedSelector的使用者如何跟ManagedSelector交互，也就是如何注册Channel以及I/O事件，提供什么样的处理类来处理I/O事件，接下来我们来看看ManagedSelector是如何统一管理和维护用户注册的Channel集合。再回到今天开始的讨论，ManagedSelector将I/O事件的生产和消费看作是生产者消费者模式，为了充分利用CPU缓存，生产和消费尽量放到同一个线程处理，那这是如何实现的呢？Jetty定义了ExecutionStrategy接口：</p>

<pre><code>public interface ExecutionStrategy
{
    //只在HTTP2中用到，简单起见，我们先忽略这个方法。
    public void dispatch();

    //实现具体执行策略，任务生产出来后可能由当前线程执行，也可能由新线程来执行
    public void produce();
    
    //任务的生产委托给Producer内部接口，
    public interface Producer
    {
        //生产一个Runnable(任务)
        Runnable produce();
    }
}
</code></pre>

<p>我们看到ExecutionStrategy接口比较简单，它将具体任务的生产委托内部接口Producer，而在自己的produce方法里来实现具体执行逻辑，<strong>也就是生产出来的任务要么由当前线程执行，要么放到新线程中执行</strong>。Jetty提供了一些具体策略实现类：ProduceConsume、ProduceExecuteConsume、ExecuteProduceConsume和EatWhatYouKill。它们的区别是：</p>

<ul>
<li>ProduceConsume：任务生产者自己依次生产和执行任务，对应到NIO通信模型就是用一个线程来侦测和处理一个ManagedSelector上所有的I/O事件，后面的I/O事件要等待前面的I/O事件处理完，效率明显不高。通过图来理解，图中绿色表示生产一个任务，蓝色表示执行这个任务。</li>
</ul>

<p><img src="assets/8b07041f80444fa685f477b532d1925c.jpg" alt="" /></p>

<ul>
<li>ProduceExecuteConsume：任务生产者开启新线程来运行任务，这是典型的I/O事件侦测和处理用不同的线程来处理，缺点是不能利用CPU缓存，并且线程切换成本高。同样我们通过一张图来理解，图中的棕色表示线程切换。</li>
</ul>

<p><img src="assets/585c608b909e4a7fb5b340a67d4e4b73.jpg" alt="" /></p>

<ul>
<li>ExecuteProduceConsume：任务生产者自己运行任务，但是该策略可能会新建一个新线程以继续生产和执行任务。这种策略也被称为“吃掉你杀的猎物”，它来自狩猎伦理，认为一个人不应该杀死他不吃掉的东西，对应线程来说，不应该生成自己不打算运行的任务。它的优点是能利用CPU缓存，但是潜在的问题是如果处理I/O事件的业务代码执行时间过长，会导致线程大量阻塞和线程饥饿。</li>
</ul>

<p><img src="assets/c76559efbb77422997e3829b932468fc.jpg" alt="" /></p>

<ul>
<li>EatWhatYouKill：这是Jetty对ExecuteProduceConsume策略的改良，在线程池线程充足的情况下等同于ExecuteProduceConsume；当系统比较忙线程不够时，切换成ProduceExecuteConsume策略。为什么要这么做呢，原因是ExecuteProduceConsume是在同一线程执行I/O事件的生产和消费，它使用的线程来自Jetty全局的线程池，这些线程有可能被业务代码阻塞，如果阻塞得多了，全局线程池中的线程自然就不够用了，最坏的情况是连I/O事件的侦测都没有线程可用了，会导致Connector拒绝浏览器请求。于是Jetty做了一个优化，在低线程情况下，就执行ProduceExecuteConsume策略，I/O侦测用专门的线程处理，I/O事件的处理扔给线程池处理，其实就是放到线程池的队列里慢慢处理。</li>
</ul>

<p>分析了这几种线程策略，我们再来看看Jetty是如何实现ExecutionStrategy接口的。答案其实就是实现Produce接口生产任务，一旦任务生产出来，ExecutionStrategy会负责执行这个任务。</p>

<pre><code>private class SelectorProducer implements ExecutionStrategy.Producer
{
    private Set&lt;SelectionKey&gt; _keys = Collections.emptySet();
    private Iterator&lt;SelectionKey&gt; _cursor = Collections.emptyIterator();

    @Override
    public Runnable produce()
    {
        while (true)
        {
            //如何Channel集合中有I/O事件就绪，调用前面提到的Selectable接口获取Runnable,直接返回给ExecutionStrategy去处理
            Runnable task = processSelected();
            if (task != null)
                return task;
            
           //如果没有I/O事件就绪，就干点杂活，看看有没有客户提交了更新Selector的任务，就是上面提到的SelectorUpdate任务类。
            processUpdates();
            updateKeys();

           //继续执行select方法，侦测I/O就绪事件
            if (!select())
                return null;
        }
    }
 }
</code></pre>

<p>SelectorProducer是ManagedSelector的内部类，SelectorProducer实现了ExecutionStrategy中的Producer接口中的produce方法，需要向ExecutionStrategy返回一个Runnable。在这个方法里SelectorProducer主要干了三件事情</p>

<ol>
<li>如果Channel集合中有I/O事件就绪，调用前面提到的Selectable接口获取Runnable，直接返回给ExecutionStrategy去处理。</li>
<li>如果没有I/O事件就绪，就干点杂活，看看有没有客户提交了更新Selector上事件注册的任务，也就是上面提到的SelectorUpdate任务类。</li>
<li>干完杂活继续执行select方法，侦测I/O就绪事件。</li>
</ol>

<h2 id="本期精华">本期精华</h2>

<p>多线程虽然是提高并发的法宝，但并不是说线程越多越好，CPU缓存以及线程上下文切换的开销也是需要考虑的。Jetty巧妙设计了EatWhatYouKill的线程策略，尽量用同一个线程侦测I/O事件和处理I/O事件，充分利用了CPU缓存，并减少了线程切换的开销。</p>

<h2 id="课后思考">课后思考</h2>

<p>文章提到ManagedSelector的使用者不能直接向它注册I/O事件，而是需要向ManagedSelector提交一个SelectorUpdate事件，ManagedSelector将这些事件Queue起来由自己来统一处理，这样做有什么好处呢？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e18d8d8dd8d5d0d0d1d6a1868c80888dcf828e8c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f15afd7f8e27771',t:'MTczNDA5MDEwNS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;&#32;编程思想：Netty&#32;中应用了哪些设计模式？>
        <link rel="icon" href="/static/favicon.png">
        <title>29  编程思想：Netty 中应用了哪些设计模式？ </title>
        
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
                        <a class="menu-item" id="00 学好 Netty，是你修炼 Java 内功的必经之路.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/00%20%e5%ad%a6%e5%a5%bd%20Netty%ef%bc%8c%e6%98%af%e4%bd%a0%e4%bf%ae%e7%82%bc%20Java%20%e5%86%85%e5%8a%9f%e7%9a%84%e5%bf%85%e7%bb%8f%e4%b9%8b%e8%b7%af.md">00 学好 Netty，是你修炼 Java 内功的必经之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  初识 Netty：为什么 Netty 这么流行？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/01%20%20%e5%88%9d%e8%af%86%20Netty%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%20Netty%20%e8%bf%99%e4%b9%88%e6%b5%81%e8%a1%8c%ef%bc%9f.md">01  初识 Netty：为什么 Netty 这么流行？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  纵览全局：把握 Netty 整体架构脉络.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/02%20%20%e7%ba%b5%e8%a7%88%e5%85%a8%e5%b1%80%ef%bc%9a%e6%8a%8a%e6%8f%a1%20Netty%20%e6%95%b4%e4%bd%93%e6%9e%b6%e6%9e%84%e8%84%89%e7%bb%9c.md">02  纵览全局：把握 Netty 整体架构脉络.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  引导器作用：客户端和服务端启动都要做些什么？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/03%20%20%e5%bc%95%e5%af%bc%e5%99%a8%e4%bd%9c%e7%94%a8%ef%bc%9a%e5%ae%a2%e6%88%b7%e7%ab%af%e5%92%8c%e6%9c%8d%e5%8a%a1%e7%ab%af%e5%90%af%e5%8a%a8%e9%83%bd%e8%a6%81%e5%81%9a%e4%ba%9b%e4%bb%80%e4%b9%88%ef%bc%9f.md">03  引导器作用：客户端和服务端启动都要做些什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 事件调度层：为什么 EventLoop 是 Netty 的精髓？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/04%20%e4%ba%8b%e4%bb%b6%e8%b0%83%e5%ba%a6%e5%b1%82%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%20EventLoop%20%e6%98%af%20Netty%20%e7%9a%84%e7%b2%be%e9%ab%93%ef%bc%9f.md">04 事件调度层：为什么 EventLoop 是 Netty 的精髓？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  服务编排层：Pipeline 如何协调各类 Handler ？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/05%20%20%e6%9c%8d%e5%8a%a1%e7%bc%96%e6%8e%92%e5%b1%82%ef%bc%9aPipeline%20%e5%a6%82%e4%bd%95%e5%8d%8f%e8%b0%83%e5%90%84%e7%b1%bb%20Handler%20%ef%bc%9f.md">05  服务编排层：Pipeline 如何协调各类 Handler ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  粘包拆包问题：如何获取一个完整的网络包？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/06%20%20%e7%b2%98%e5%8c%85%e6%8b%86%e5%8c%85%e9%97%ae%e9%a2%98%ef%bc%9a%e5%a6%82%e4%bd%95%e8%8e%b7%e5%8f%96%e4%b8%80%e4%b8%aa%e5%ae%8c%e6%95%b4%e7%9a%84%e7%bd%91%e7%bb%9c%e5%8c%85%ef%bc%9f.md">06  粘包拆包问题：如何获取一个完整的网络包？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  接头暗语：如何利用 Netty 实现自定义协议通信？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/07%20%20%e6%8e%a5%e5%a4%b4%e6%9a%97%e8%af%ad%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20Netty%20%e5%ae%9e%e7%8e%b0%e8%87%aa%e5%ae%9a%e4%b9%89%e5%8d%8f%e8%ae%ae%e9%80%9a%e4%bf%a1%ef%bc%9f.md">07  接头暗语：如何利用 Netty 实现自定义协议通信？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  开箱即用：Netty 支持哪些常用的解码器？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/08%20%20%e5%bc%80%e7%ae%b1%e5%8d%b3%e7%94%a8%ef%bc%9aNetty%20%e6%94%af%e6%8c%81%e5%93%aa%e4%ba%9b%e5%b8%b8%e7%94%a8%e7%9a%84%e8%a7%a3%e7%a0%81%e5%99%a8%ef%bc%9f.md">08  开箱即用：Netty 支持哪些常用的解码器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  数据传输：writeAndFlush 处理流程剖析.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/09%20%20%e6%95%b0%e6%8d%ae%e4%bc%a0%e8%be%93%ef%bc%9awriteAndFlush%20%e5%a4%84%e7%90%86%e6%b5%81%e7%a8%8b%e5%89%96%e6%9e%90.md">09  数据传输：writeAndFlush 处理流程剖析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  双刃剑：合理管理 Netty 堆外内存.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/10%20%20%e5%8f%8c%e5%88%83%e5%89%91%ef%bc%9a%e5%90%88%e7%90%86%e7%ae%a1%e7%90%86%20Netty%20%e5%a0%86%e5%a4%96%e5%86%85%e5%ad%98.md">10  双刃剑：合理管理 Netty 堆外内存.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  另起炉灶：Netty 数据传输载体 ByteBuf 详解.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/11%20%20%e5%8f%a6%e8%b5%b7%e7%82%89%e7%81%b6%ef%bc%9aNetty%20%e6%95%b0%e6%8d%ae%e4%bc%a0%e8%be%93%e8%bd%bd%e4%bd%93%20ByteBuf%20%e8%af%a6%e8%a7%a3.md">11  另起炉灶：Netty 数据传输载体 ByteBuf 详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  他山之石：高性能内存分配器 jemalloc 基本原理.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/12%20%20%e4%bb%96%e5%b1%b1%e4%b9%8b%e7%9f%b3%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e5%86%85%e5%ad%98%e5%88%86%e9%85%8d%e5%99%a8%20jemalloc%20%e5%9f%ba%e6%9c%ac%e5%8e%9f%e7%90%86.md">12  他山之石：高性能内存分配器 jemalloc 基本原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  举一反三：Netty 高性能内存管理设计（上）.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/13%20%20%e4%b8%be%e4%b8%80%e5%8f%8d%e4%b8%89%ef%bc%9aNetty%20%e9%ab%98%e6%80%a7%e8%83%bd%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89.md">13  举一反三：Netty 高性能内存管理设计（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  举一反三：Netty 高性能内存管理设计（下）.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/14%20%20%e4%b8%be%e4%b8%80%e5%8f%8d%e4%b8%89%ef%bc%9aNetty%20%e9%ab%98%e6%80%a7%e8%83%bd%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89.md">14  举一反三：Netty 高性能内存管理设计（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  轻量级对象回收站：Recycler 对象池技术解析.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/15%20%20%e8%bd%bb%e9%87%8f%e7%ba%a7%e5%af%b9%e8%b1%a1%e5%9b%9e%e6%94%b6%e7%ab%99%ef%bc%9aRecycler%20%e5%af%b9%e8%b1%a1%e6%b1%a0%e6%8a%80%e6%9c%af%e8%a7%a3%e6%9e%90.md">15  轻量级对象回收站：Recycler 对象池技术解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  IO 加速：与众不同的 Netty 零拷贝技术.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/16%20%20IO%20%e5%8a%a0%e9%80%9f%ef%bc%9a%e4%b8%8e%e4%bc%97%e4%b8%8d%e5%90%8c%e7%9a%84%20Netty%20%e9%9b%b6%e6%8b%b7%e8%b4%9d%e6%8a%80%e6%9c%af.md">16  IO 加速：与众不同的 Netty 零拷贝技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  源码篇：从 Linux 出发深入剖析服务端启动流程.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/17%20%20%e6%ba%90%e7%a0%81%e7%af%87%ef%bc%9a%e4%bb%8e%20Linux%20%e5%87%ba%e5%8f%91%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%e6%9c%8d%e5%8a%a1%e7%ab%af%e5%90%af%e5%8a%a8%e6%b5%81%e7%a8%8b.md">17  源码篇：从 Linux 出发深入剖析服务端启动流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  源码篇：解密 Netty Reactor 线程模型.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/18%20%20%e6%ba%90%e7%a0%81%e7%af%87%ef%bc%9a%e8%a7%a3%e5%af%86%20Netty%20Reactor%20%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b.md">18  源码篇：解密 Netty Reactor 线程模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  源码篇：一个网络请求在 Netty 中的旅程.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/19%20%20%e6%ba%90%e7%a0%81%e7%af%87%ef%bc%9a%e4%b8%80%e4%b8%aa%e7%bd%91%e7%bb%9c%e8%af%b7%e6%b1%82%e5%9c%a8%20Netty%20%e4%b8%ad%e7%9a%84%e6%97%85%e7%a8%8b.md">19  源码篇：一个网络请求在 Netty 中的旅程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  技巧篇：Netty 的 FastThreadLocal 究竟比 ThreadLocal 快在哪儿？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/20%20%20%e6%8a%80%e5%b7%a7%e7%af%87%ef%bc%9aNetty%20%e7%9a%84%20FastThreadLocal%20%e7%a9%b6%e7%ab%9f%e6%af%94%20ThreadLocal%20%e5%bf%ab%e5%9c%a8%e5%93%aa%e5%84%bf%ef%bc%9f.md">20  技巧篇：Netty 的 FastThreadLocal 究竟比 ThreadLocal 快在哪儿？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  技巧篇：延迟任务处理神器之时间轮 HashedWheelTimer.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/21%20%20%e6%8a%80%e5%b7%a7%e7%af%87%ef%bc%9a%e5%bb%b6%e8%bf%9f%e4%bb%bb%e5%8a%a1%e5%a4%84%e7%90%86%e7%a5%9e%e5%99%a8%e4%b9%8b%e6%97%b6%e9%97%b4%e8%bd%ae%20HashedWheelTimer.md">21  技巧篇：延迟任务处理神器之时间轮 HashedWheelTimer.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  技巧篇：高性能无锁队列 Mpsc Queue.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/22%20%20%e6%8a%80%e5%b7%a7%e7%af%87%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e6%97%a0%e9%94%81%e9%98%9f%e5%88%97%20Mpsc%20Queue.md">22  技巧篇：高性能无锁队列 Mpsc Queue.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  架构设计：如何实现一个高性能分布式 RPC 框架.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/23%20%20%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e9%ab%98%e6%80%a7%e8%83%bd%e5%88%86%e5%b8%83%e5%bc%8f%20RPC%20%e6%a1%86%e6%9e%b6.md">23  架构设计：如何实现一个高性能分布式 RPC 框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  服务发布与订阅：搭建生产者和消费者的基础框架.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/24%20%20%e6%9c%8d%e5%8a%a1%e5%8f%91%e5%b8%83%e4%b8%8e%e8%ae%a2%e9%98%85%ef%bc%9a%e6%90%ad%e5%bb%ba%e7%94%9f%e4%ba%a7%e8%80%85%e5%92%8c%e6%b6%88%e8%b4%b9%e8%80%85%e7%9a%84%e5%9f%ba%e7%a1%80%e6%a1%86%e6%9e%b6.md">24  服务发布与订阅：搭建生产者和消费者的基础框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  远程通信：通信协议设计以及编解码的实现.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/25%20%20%e8%bf%9c%e7%a8%8b%e9%80%9a%e4%bf%a1%ef%bc%9a%e9%80%9a%e4%bf%a1%e5%8d%8f%e8%ae%ae%e8%ae%be%e8%ae%a1%e4%bb%a5%e5%8f%8a%e7%bc%96%e8%a7%a3%e7%a0%81%e7%9a%84%e5%ae%9e%e7%8e%b0.md">25  远程通信：通信协议设计以及编解码的实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  服务治理：服务发现与负载均衡机制的实现.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/26%20%20%e6%9c%8d%e5%8a%a1%e6%b2%bb%e7%90%86%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%8f%91%e7%8e%b0%e4%b8%8e%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e6%9c%ba%e5%88%b6%e7%9a%84%e5%ae%9e%e7%8e%b0.md">26  服务治理：服务发现与负载均衡机制的实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  动态代理：为用户屏蔽 RPC 调用的底层细节.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/27%20%20%e5%8a%a8%e6%80%81%e4%bb%a3%e7%90%86%ef%bc%9a%e4%b8%ba%e7%94%a8%e6%88%b7%e5%b1%8f%e8%94%bd%20RPC%20%e8%b0%83%e7%94%a8%e7%9a%84%e5%ba%95%e5%b1%82%e7%bb%86%e8%8a%82.md">27  动态代理：为用户屏蔽 RPC 调用的底层细节.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  实战总结：RPC 实战总结与进阶延伸.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/28%20%20%e5%ae%9e%e6%88%98%e6%80%bb%e7%bb%93%ef%bc%9aRPC%20%e5%ae%9e%e6%88%98%e6%80%bb%e7%bb%93%e4%b8%8e%e8%bf%9b%e9%98%b6%e5%bb%b6%e4%bc%b8.md">28  实战总结：RPC 实战总结与进阶延伸.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  编程思想：Netty 中应用了哪些设计模式？.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/29%20%20%e7%bc%96%e7%a8%8b%e6%80%9d%e6%83%b3%ef%bc%9aNetty%20%e4%b8%ad%e5%ba%94%e7%94%a8%e4%ba%86%e5%93%aa%e4%ba%9b%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%9f.md">29  编程思想：Netty 中应用了哪些设计模式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  实践总结：Netty 在项目开发中的一些最佳实践.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/30%20%20%e5%ae%9e%e8%b7%b5%e6%80%bb%e7%bb%93%ef%bc%9aNetty%20%e5%9c%a8%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e4%b8%ad%e7%9a%84%e4%b8%80%e4%ba%9b%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5.md">30  实践总结：Netty 在项目开发中的一些最佳实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 结束语  技术成长之路：如何打造自己的技术体系.md" href="/%e4%b8%93%e6%a0%8f/Netty%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%e4%b8%8e%20RPC%20%e5%ae%9e%e8%b7%b5-%e5%ae%8c/31%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e6%8a%80%e6%9c%af%e6%88%90%e9%95%bf%e4%b9%8b%e8%b7%af%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%93%e9%80%a0%e8%87%aa%e5%b7%b1%e7%9a%84%e6%8a%80%e6%9c%af%e4%bd%93%e7%b3%bb.md">31 结束语  技术成长之路：如何打造自己的技术体系.md</a>
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
                            <h1 id="title" data-id="29  编程思想：Netty 中应用了哪些设计模式？" class="title">29  编程思想：Netty 中应用了哪些设计模式？</h1>
                            <div><p>设计模式的运用是面试过程中常考的，学习设计模式切勿死记硬背，结合优秀项目的源码去理解设计模式的使用会事半功倍。Netty 源码中运用了大量的设计模式，常见的设计模式在 Netty 源码中都有所体现。本节课我们便一起梳理 Netty 源码中所包含的设计模式，希望能帮助你更深入地了解 Netty 的设计精髓，并可以结合 Netty 源码向面试官讲述你对设计模式的理解。</p>

<h3 id="单例模式">单例模式</h3>

<p>单例模式是最常见的设计模式，它可以保证全局只有一个实例，避免线程安全问题。单例模式有很多种实现方法，其中我比较推荐三种最佳实践：<strong>双重检验锁</strong>、<strong>静态内部类方式</strong>、<strong>饿汉方式</strong>和<strong>枚举方式</strong>，其中双重检验锁和静态内部类方式属于懒汉式单例，饿汉方式和枚举方式属于饿汉式单例。</p>

<h4 id="双重检验锁">双重检验锁</h4>

<p>在多线程环境下，为了提高实例初始化的性能，不是每次获取实例时在方法上加锁，而是当实例未创建时才会加锁，如下所示：</p>

<pre><code>public class SingletonTest {

    private SingletonTest instance;

    public static SingletonTest getInstance() {

        if (instance == null) {

            synchronized (this) {

                if (instance == null) {

                    instance = new SingletonTest();

                }

            }

        }

        return instance;

    }

}
</code></pre>

<h4 id="静态内部类方式">静态内部类方式</h4>

<p>静态内部类方式实现单例巧妙地利用了 Java 类加载机制，保证其在多线程环境下的线程安全性。当一个类被加载时，其静态内部类是不会被同时加载的，只有第一次被调用时才会初始化，而且我们不能通过反射的方式获取内部的属性。由此可见，静态内部类方式实现单例更加安全，可以防止被反射入侵。具体实现方式如下：</p>

<pre><code>public class SingletonTest {

    private SingletonTest() {

    }

    public static Singleton getInstance() {

        return SingletonInstance.instance;

    }

    private static class SingletonInstance {

        private static final Singleton instance = new Singleton();

    }

}
</code></pre>

<h4 id="饿汉方式">饿汉方式</h4>

<p>饿汉式实现单例非常简单，类加载的时候就创建出实例。饿汉方式使用私有构造函数实现全局单个实例的初始化，并使用 public static final 加以修饰，实现延迟加载和保证线程安全性。实现方式如下所示：</p>

<pre><code>public class SingletonTest {

    private static Singleton instance = new Singleton();

    private Singleton() {

    }

    public static Singleton getInstance() {

        return instance;

    }

}
</code></pre>

<h4 id="枚举方式">枚举方式</h4>

<p>枚举方式是一种天然的单例实现，在项目开发中枚举方式是非常推荐使用的。它能够保证序列化和反序列化过程中实例的唯一性，而且不用担心线程安全问题。枚举方式实现单例如下所示：</p>

<pre><code>public enum SingletonTest {

    SERVICE_A {

        @Override

        protected void hello() {

            System.out.println(&quot;hello, service A&quot;);

        }

    },

    SERVICE_B {

        @Override

        protected void hello() {

            System.out.println(&quot;hello, service B&quot;);

        }

    };

    protected abstract void hello();

}
</code></pre>

<p>在《源码篇：解密 Netty Reactor 的线程模型》课程中，我们介绍了 NioEventLoop 的核心原理。NioEventLoop 通过核心方法 select() 不断轮询注册的 I/O 事件，Netty 提供了选择策略 SelectStrategy 对象，它用于控制 select 循环行为，包含 CONTINUE、SELECT、BUSY_WAIT 三种策略。SelectStrategy 对象的默认实现就是使用的饿汉式单例，源码如下：</p>

<pre><code>final class DefaultSelectStrategy implements SelectStrategy {

    static final SelectStrategy INSTANCE = new DefaultSelectStrategy();

    private DefaultSelectStrategy() { }

    @Override

    public int calculateStrategy(IntSupplier selectSupplier, boolean hasTasks) throws Exception {

        return hasTasks ? selectSupplier.get() : SelectStrategy.SELECT;

    }

}
</code></pre>

<p>此外 Netty 中还有不少饿汉方式实现单例的实践，例如 MqttEncoder、ReadTimeoutException 等。</p>

<h3 id="工厂方法模式">工厂方法模式</h3>

<p>工厂模式封装了对象创建的过程，使用者不需要关心对象创建的细节。在需要生成复杂对象的场景下，都可以使用工厂模式实现。工厂模式分为三种：<strong>简单工厂模式</strong>、<strong>工厂方法模式</strong>和<strong>抽象工厂模式</strong>。</p>

<ul>
<li><strong>简单工厂模式</strong>。定义一个工厂类，根据参数类型返回不同类型的实例。适用于对象实例类型不多的场景，如果对象实例类型太多，每增加一种类型就要在工厂类中增加相应的创建逻辑，这是违背开放封闭原则的。</li>
<li><strong>工厂方法模式</strong>。简单工厂模式的升级版，不再是提供一个统一的工厂类来创建所有对象的实例，而是每种类型的对象实例都对应不同的工厂类，每个具体的工厂类只能创建一个类型的对象实例。</li>
<li><strong>抽象工厂模式</strong>。较少使用，适用于创建多个产品的场景。如果按照工厂方法模式的实现思路，需要在具体工厂类中实现多个工厂方法，是非常不友好的。抽象工厂模式就是把这些工厂方法单独剥离到抽象工厂类中，然后创建工厂对象并通过组合的方式来获取工厂方法。</li>
</ul>

<p>Netty 中使用的就是工厂方法模式，这也是项目开发中最常用的一种工厂模式。工厂方法模式如何使用呢？我们先来看个简单的例子：</p>

<pre><code>public class TSLAFactory implements CarFactory {

    @Override

    public Car createCar() {

        return new TSLA();

    }

}

public class BMWFactory implements CarFactory {

    @Override

    public Car createCar() {

        return new BMW();

    }

}
</code></pre>

<p>Netty 在创建 Channel 的时候使用的就是工厂方法模式，因为服务端和客户端的 Channel 是不一样的。Netty 将反射和工厂方法模式结合在一起，只使用一个工厂类，然后根据传入的 Class 参数来构建出对应的 Channel，不需要再为每一种 Channel 类型创建一个工厂类。具体源码实现如下：</p>

<pre><code>public class ReflectiveChannelFactory&lt;T extends Channel&gt; implements ChannelFactory&lt;T&gt; {

    private final Constructor&lt;? extends T&gt; constructor;

    public ReflectiveChannelFactory(Class&lt;? extends T&gt; clazz) {

        ObjectUtil.checkNotNull(clazz, &quot;clazz&quot;);

        try {

            this.constructor = clazz.getConstructor();

        } catch (NoSuchMethodException e) {

            throw new IllegalArgumentException(&quot;Class &quot; + StringUtil.simpleClassName(clazz) +

                    &quot; does not have a public non-arg constructor&quot;, e);

        }

    }

    @Override

    public T newChannel() {

        try {

            return constructor.newInstance();

        } catch (Throwable t) {

            throw new ChannelException(&quot;Unable to create Channel from class &quot; + constructor.getDeclaringClass(), t);

        }

    }

    @Override

    public String toString() {

        return StringUtil.simpleClassName(ReflectiveChannelFactory.class) +

                '(' + StringUtil.simpleClassName(constructor.getDeclaringClass()) + &quot;.class)&quot;;

    }

}
</code></pre>

<p>虽然通过反射技术可以有效地减少工厂类的数据量，但是反射相比直接创建工厂类有性能损失，所以对于性能敏感的场景，应当谨慎使用反射。</p>

<h3 id="责任链模式">责任链模式</h3>

<p>想必学完本专栏的前面课程后，责任链模式大家应该再熟悉不过了，自然而然联想到 ChannlPipeline 和 ChannelHandler。ChannlPipeline 内部是由一组 ChannelHandler 实例组成的，内部通过双向链表将不同的 ChannelHandler 链接在一起，如下图所示。</p>

<p><img src="assets/Cip5yGAKfLmANJ_0AAZUvQP4FxQ293.png" alt="Drawing 0.png" /></p>

<p>对于 Netty 中责任链模式的实现，也遵循了责任链模式的四个基本要素：</p>

<h4 id="责任处理器接口">责任处理器接口</h4>

<p>ChannelHandler 对应的就是责任处理器接口，ChannelHandler 有两个重要的子接口：ChannelInboundHandler和ChannelOutboundHandler，分别拦截入站和出站的各种 I/O 事件。</p>

<h4 id="动态创建责任链-添加-删除责任处理器">动态创建责任链，添加、删除责任处理器</h4>

<p>ChannelPipeline 负责创建责任链，其内部采用双向链表实现，ChannelPipeline 的内部结构定义如下所示：</p>

<pre><code>public class DefaultChannelPipeline implements ChannelPipeline {

    static final InternalLogger logger = InternalLoggerFactory.getInstance(DefaultChannelPipeline.class);

    private static final String HEAD_NAME = generateName0(HeadContext.class);

    private static final String TAIL_NAME = generateName0(TailContext.class);

    // 省略其他代码

    final AbstractChannelHandlerContext head; // 头结点

    final AbstractChannelHandlerContext tail; // 尾节点

    private final Channel channel;

    private final ChannelFuture succeededFuture;

    private final VoidChannelPromise voidPromise;

    private final boolean touch = ResourceLeakDetector.isEnabled();

    

    // 省略其他代码

}
</code></pre>

<p>ChannelPipeline 提供了一系列 add 和 remove 相关接口用于动态添加和删除 ChannelHandler 处理器，如下所示：</p>

<p><img src="assets/CgpVE2AKfMSAehuyAA0_w5_fB-c556.png" alt="Drawing 1.png" /></p>

<h4 id="上下文">上下文</h4>

<p>从 ChannelPipeline 内部结构定义可以看出，ChannelHandlerContext 负责保存责任链节点上下文信息。ChannelHandlerContext 是对 ChannelHandler 的封装，每个 ChannelHandler 都对应一个 ChannelHandlerContext，实际上 ChannelPipeline 维护的是与 ChannelHandlerContext 的关系。</p>

<h4 id="责任传播和终止机制">责任传播和终止机制</h4>

<p>ChannelHandlerContext 提供了 fire 系列的方法用于事件传播，如下所示：</p>

<p><img src="assets/Cip5yGAKfM6AHTnGAA43hnlCx54417.png" alt="Drawing 2.png" /></p>

<p>以 ChannelInboundHandlerAdapter 的 channelRead 方法为例，ChannelHandlerContext 会默认调用 fireChannelRead 方法将事件默认传递到下一个处理器。如果我们重写了 ChannelInboundHandlerAdapter 的 channelRead 方法，并且没有调用 fireChannelRead 进行事件传播，那么表示此次事件传播已终止。</p>

<h3 id="观察者模式">观察者模式</h3>

<p>观察者模式有两个角色：观察者和被观察。被观察者发布消息，观察者订阅消息，没有订阅的观察者是收不到消息的。首先我们通过一个简单的例子看下观察者模式的是如何实现的。</p>

<pre><code>// 被观察者

public interface Observable {

    void registerObserver(Observer observer);

    void removeObserver(Observer observer);

    void notifyObservers(String message);

}

// 观察者

public interface Observer {

    void notify(String message);

}

// 默认被观察者实现

public class DefaultObservable implements Observable {

    private final List&lt;Observer&gt; observers = new ArrayList&lt;&gt;();

    @Override

    public void registerObserver(Observer observer) {

        observers.add(observer);

    }

    @Override

    public void removeObserver(Observer observer) {

        observers.remove(observer);

    }

    @Override

    public void notifyObservers(String message) {

        for (Observer observer : observers) {

            observer.notify(message);

        }

    }

}
</code></pre>

<p>Netty 中观察者模式的运用非常多，但是并没有以上示例代码这么直观，我们平时经常使用的ChannelFuture#addListener 接口就是观察者模式的实现。我们先来看下 ChannelFuture 使用的示例：</p>

<pre><code>ChannelFuture channelFuture = channel.writeAndFlush(object);

channelFuture.addListener(future -&gt; {

    if (future.isSuccess()) {

        // do something

    } else {

        // do something

    }

});
</code></pre>

<p>addListener 方法会将添加监听器添加到 ChannelFuture 当中，并在 ChannelFuture 执行完毕的时候立刻通知已经注册的监听器。所以 ChannelFuture 是被观察者，addListener 方法用于添加观察者。</p>

<h3 id="建造者模式">建造者模式</h3>

<p>建造者模式非常简单，通过链式调用来设置对象的属性，在对象属性繁多的场景下非常有用。建造者模式的优势就是可以像搭积木一样自由选择需要的属性，并不是强绑定的。对于使用者来说，必须清楚需要设置哪些属性，在不同场景下可能需要的属性也是不一样的。</p>

<p>Netty 中 ServerBootStrap 和 Bootstrap 引导器是最经典的建造者模式实现，在构建过程中需要设置非常多的参数，例如配置线程池 EventLoopGroup、设置 Channel 类型、注册 ChannelHandler、设置 Channel 参数、端口绑定等。ServerBootStrap 引导器的具体使用可以参考《引导器作用：客户端和服务端启动都要做些什么？》课程，在此我就不多作赘述了。</p>

<h3 id="策略模式">策略模式</h3>

<p>策略模式针对同一个问题提供多种策略的处理方式，这些策略之间可以相互替换，在一定程度上提高了系统的灵活性。策略模式非常符合开闭原则，使用者在不修改现有系统的情况下选择不同的策略，而且便于扩展增加新的策略。</p>

<p>Netty 在多处地方使用了策略模式，例如 EventExecutorChooser 提供了不同的策略选择 NioEventLoop，newChooser() 方法会根据线程池的大小是否是 2 的幂次，以此来动态的选择取模运算的方式，从而提高性能。EventExecutorChooser 源码实现如下所示：</p>

<pre><code>public final class DefaultEventExecutorChooserFactory implements EventExecutorChooserFactory {

    public static final DefaultEventExecutorChooserFactory INSTANCE = new DefaultEventExecutorChooserFactory();

    private DefaultEventExecutorChooserFactory() { }

    @SuppressWarnings(&quot;unchecked&quot;)

    @Override

    public EventExecutorChooser newChooser(EventExecutor[] executors) {

        if (isPowerOfTwo(executors.length)) {

            return new PowerOfTwoEventExecutorChooser(executors);

        } else {

            return new GenericEventExecutorChooser(executors);

        }

    }

    

    // 省略其他代码

}
</code></pre>

<h3 id="装饰者模式">装饰者模式</h3>

<p>装饰器模式是对被装饰类的功能增强，在不修改被装饰类的前提下，能够为被装饰类添加新的功能特性。当我们需要为一个类扩展功能时会使用装饰器模式，但是该模式的缺点是需要增加额外的代码。我们先通过一个简单的例子学习下装饰器模式应当如何使用，如下所示：</p>

<pre><code>public interface Shape {

    void draw();

}

class Circle implements Shape {

    @Override

    public void draw() {

        System.out.print(&quot;draw a circle.&quot;);

    }

}

abstract class ShapeDecorator implements Shape {

    protected Shape shapeDecorated;

    public ShapeDecorator(Shape shapeDecorated) {

        this.shapeDecorated = shapeDecorated;

    }

    public void draw() {

        shapeDecorated.draw();

    }

}

class FillReadColorShapeDecorator extends ShapeDecorator {

    public FillReadColorShapeDecorator(Shape shapeDecorated) {

        super(shapeDecorated);

    }

    @Override

    public void draw() {

        shapeDecorated.draw();

        fillColor();

    }

    private void fillColor() {

        System.out.println(&quot;Fill Read Color.&quot;);

    }

}
</code></pre>

<p>我们创建了一个 Shape 接口的抽象装饰类 ShapeDecorator，并维护 Shape 原始对象，FillReadColorShapeDecorator 是用于装饰 ShapeDecorator 的实体类，它不对 draw() 方法做任何修改，而是直接调用 Shape 对象原有的 draw() 方法，然后再调用 fillColor() 方法进行颜色填充。</p>

<p>下面我们再来看一下 Netty 中 WrappedByteBuf 是如何装饰 ByteBuf 的，源码如下所示：</p>

<pre><code>class WrappedByteBuf extends ByteBuf {

    protected final ByteBuf buf;

    protected WrappedByteBuf(ByteBuf buf) {

        if (buf == null) {

            throw new NullPointerException(&quot;buf&quot;);

        }

        this.buf = buf;

    }

    @Override

    public final boolean hasMemoryAddress() {

        return buf.hasMemoryAddress();

    }

    @Override

    public final long memoryAddress() {

        return buf.memoryAddress();

    }

    

    // 省略其他代码

}
</code></pre>

<p>WrappedByteBuf 是所有 ByteBuf 装饰器的基类，它并没有什么特别的，也是在构造函数里传入了原始的 ByteBuf 实例作为被装饰者。WrappedByteBuf 有两个子类 UnreleasableByteBuf 和 SimpleLeakAwareByteBuf，它们是真正实现对 ByteBuf 的功能增强，例如 UnreleasableByteBuf 类的 release() 方法是直接返回 false 表示不可被释放，源码实现如下所示。</p>

<pre><code>final class UnreleasableByteBuf extends WrappedByteBuf {

    private SwappedByteBuf swappedBuf;

    UnreleasableByteBuf(ByteBuf buf) {

        super(buf instanceof UnreleasableByteBuf ? buf.unwrap() : buf);

    }

    @Override

    public boolean release() {

        return false;

    }

    // 省略其他代码

}
</code></pre>

<p>不知道你会不会有一个疑问，装饰器模式和代理模式都是实现目标类增强，他们有什么区别吗？装饰器模式和代理模式的实现确实是非常相似的，都需要维护原始的目标对象，装饰器模式侧重于为目标类增加新的功能，代理模式更侧重于在现有功能的基础上进行扩展。</p>

<h3 id="总结">总结</h3>

<p>学习设计模式切勿死记硬背，不仅要吸收设计模式的思想，还要理解为什么使用该设计模式。锻炼代码设计能力比较好的办法就是读优秀框架的源码，Netty 就是一个非常丰富的学习资源。我们需要了解源码中设计模式的使用场景，不断吸收消化，并能够做到在项目开发中学以致用。</p>

<p>本节课所介绍的设计模式在 Netty 中并不能涵盖所有，还有很多等待我们去挖掘。留一个课后任务，学习以下设计模式在 Netty 中的运用。</p>

<ul>
<li>模板方法模式：ServerBootStrap 和 Bootstrap 的 init 过程实现；</li>
<li>迭代器模式：CompositeByteBuf；</li>
<li>适配器模式：ScheduledFutureTask。</li>
</ul>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#4b272727727f7a7a7b7c0b2c262a222765282426" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1185a1395294f1',t:'MTczNDA0NjQzMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
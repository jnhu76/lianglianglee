<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=07&#32;&#32;接头暗语：如何利用&#32;Netty&#32;实现自定义协议通信？>
        <link rel="icon" href="/static/favicon.png">
        <title>07  接头暗语：如何利用 Netty 实现自定义协议通信？ </title>
        
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
                            <h1 id="title" data-id="07  接头暗语：如何利用 Netty 实现自定义协议通信？" class="title">07  接头暗语：如何利用 Netty 实现自定义协议通信？</h1>
                            <div><p>既然是网络编程，自然离不开通信协议，应用层之间通信需要实现各种各样的网络协议。在项目开发的过程中，我们就需要去构建满足自己业务场景的应用层协议。在上节课中我们介绍了如何使用网络协议解决 TCP 拆包/粘包的底层问题，本节课我们将在此基础上继续讨论如何设计一个高效、可扩展、易维护的自定义通信协议，以及如何使用 Netty 实现自定义通信协议。</p>

<h3 id="通信协议设计">通信协议设计</h3>

<p>所谓协议，就是通信双方事先商量好的接口暗语，在 TCP 网络编程中，发送方和接收方的数据包格式都是二进制，发送方将对象转化成二进制流发送给接收方，接收方获得二进制数据后需要知道如何解析成对象，所以协议是<strong>双方能够正常通信的基础</strong>。</p>

<p>目前市面上已经有不少通用的协议，例如 HTTP、HTTPS、JSON-RPC、FTP、IMAP、Protobuf 等。<strong>通用协议</strong>兼容性好，易于维护，各种异构系统之间可以实现无缝对接。如果在满足业务场景以及性能需求的前提下，推荐采用通用协议的方案。相比通用协议，自定义协议主要有以下优点。</p>

<ul>
<li><strong>极致性能</strong>：通用的通信协议考虑了很多兼容性的因素，必然在性能方面有所损失。</li>
<li><strong>扩展性</strong>：自定义的协议相比通用协议更好扩展，可以更好地满足自己的业务需求。</li>
<li><strong>安全性</strong>：通用协议是公开的，很多漏洞已经很多被黑客攻破。自定义协议更加<strong>安全</strong>，因为黑客需要先破解你的协议内容。</li>
</ul>

<p>那么如何设计自定义的通信协议呢？这个答案见仁见智，但是设计通信协议有经验方法可循。结合实战经验我们一起看下一个完备的网络协议需要具备哪些基本要素。</p>

<h4 id="1-魔数">1. 魔数</h4>

<p>魔数是通信双方协商的一个暗号，通常采用固定的几个字节表示。魔数的作用是<strong>防止任何人随便向服务器的端口上发送数据</strong>。服务端在接收到数据时会解析出前几个固定字节的魔数，然后做正确性比对。如果和约定的魔数不匹配，则认为是非法数据，可以直接关闭连接或者采取其他措施以增强系统的安全防护。魔数的思想在压缩算法、Java Class 文件等场景中都有所体现，例如 Class 文件开头就存储了魔数 0xCAFEBABE，在加载 Class 文件时首先会验证魔数的正确性。</p>

<h4 id="2-协议版本号">2. 协议版本号</h4>

<p>随着业务需求的变化，协议可能需要对结构或字段进行改动，不同版本的协议对应的解析方法也是不同的。所以在生产级项目中强烈建议预留<strong>协议版本号</strong>这个字段。</p>

<h4 id="3-序列化算法">3. 序列化算法</h4>

<p>序列化算法字段表示数据发送方应该采用何种方法将请求的对象转化为二进制，以及如何再将二进制转化为对象，如 JSON、Hessian、Java 自带序列化等。</p>

<h4 id="4-报文类型">4. 报文类型</h4>

<p>在不同的业务场景中，报文可能存在不同的类型。例如在 RPC 框架中有请求、响应、心跳等类型的报文，在 IM 即时通信的场景中有登陆、创建群聊、发送消息、接收消息、退出群聊等类型的报文。</p>

<h4 id="5-长度域字段">5. 长度域字段</h4>

<p>长度域字段代表<strong>请求数据</strong>的长度，接收方根据长度域字段获取一个完整的报文。</p>

<h4 id="6-请求数据">6. 请求数据</h4>

<p>请求数据通常为序列化之后得到的<strong>二进制流</strong>，每种请求数据的内容是不一样的。</p>

<h4 id="7-状态">7. 状态</h4>

<p>状态字段用于标识<strong>请求是否正常</strong>。一般由被调用方设置。例如一次 RPC 调用失败，状态字段可被服务提供方设置为异常状态。</p>

<h4 id="8-保留字段">8. 保留字段</h4>

<p>保留字段是可选项，为了应对协议升级的可能性，可以预留若干字节的保留字段，以备不时之需。</p>

<p>通过以上协议基本要素的学习，我们可以得到一个较为通用的协议示例：</p>

<pre><code>+---------------------------------------------------------------+

| 魔数 2byte | 协议版本号 1byte | 序列化算法 1byte | 报文类型 1byte  |

+---------------------------------------------------------------+

| 状态 1byte |        保留字段 4byte     |      数据长度 4byte     | 

+---------------------------------------------------------------+

|                   数据内容 （长度不定）                          |

+---------------------------------------------------------------+
</code></pre>

<h3 id="netty-如何实现自定义通信协议">Netty 如何实现自定义通信协议</h3>

<p>在学习完如何设计协议之后，我们又该如何在 Netty 中实现自定义的通信协议呢？其实 Netty 作为一个非常优秀的网络通信框架，已经为我们提供了非常丰富的编解码抽象基类，帮助我们更方便地基于这些抽象基类扩展实现自定义协议。</p>

<p>首先我们看下 Netty 中编解码器是如何分类的。</p>

<p><strong>Netty 常用编码器类型：</strong></p>

<ul>
<li>MessageToByteEncoder 对象编码成字节流；</li>
<li>MessageToMessageEncoder 一种消息类型编码成另外一种消息类型。</li>
</ul>

<p><strong>Netty 常用解码器类型：</strong></p>

<ul>
<li>ByteToMessageDecoder/ReplayingDecoder 将字节流解码为消息对象；</li>
<li>MessageToMessageDecoder 将一种消息类型解码为另外一种消息类型。</li>
</ul>

<p>编解码器可以分为<strong>一次解码器</strong>和<strong>二次解码器</strong>，一次解码器用于解决 TCP 拆包/粘包问题，按协议解析后得到的字节数据。如果你需要对解析后的字节数据做对象模型的转换，这时候便需要用到二次解码器，同理编码器的过程是反过来的。</p>

<ul>
<li>一次编解码器：MessageToByteEncoder/ByteToMessageDecoder。</li>
<li>二次编解码器：MessageToMessageEncoder/MessageToMessageDecoder。</li>
</ul>

<p>下面我们对 Netty 中常用的抽象编解码类进行详细的介绍。</p>

<h4 id="抽象编码类">抽象编码类</h4>

<p><img src="assets/CgqCHl-lDlWAUGw2AAMI8fuvukI452.png" alt="Drawing 0.png" /></p>

<p>通过抽象编码类的继承图可以看出，编码类是 ChanneOutboundHandler 的抽象类实现，具体操作的是 Outbound 出站数据。</p>

<ul>
<li><strong>MessageToByteEncoder</strong></li>
</ul>

<p>MessageToByteEncoder 用于将对象编码成<strong>字节流</strong>，MessageToByteEncoder 提供了唯一的 encode 抽象方法，我们只需要实现<strong>encode 方法</strong>即可完成自定义编码。那么encode() 方法是在什么时候被调用的呢？我们一起看下MessageToByteEncoder 的核心源码片段，如下所示。</p>

<pre><code>@Override

public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) throws Exception {

    ByteBuf buf = null;

    try {

        if (acceptOutboundMessage(msg)) { // 1. 消息类型是否匹配

            @SuppressWarnings(&quot;unchecked&quot;)

            I cast = (I) msg;

            buf = allocateBuffer(ctx, cast, preferDirect); // 2. 分配 ByteBuf 资源

            try {

                encode(ctx, cast, buf); // 3. 执行 encode 方法完成数据编码

            } finally {

                ReferenceCountUtil.release(cast);

            }

            if (buf.isReadable()) {

                ctx.write(buf, promise); // 4. 向后传递写事件

            } else {

                buf.release();

                ctx.write(Unpooled.EMPTY_BUFFER, promise);

            }

            buf = null;

        } else {

            ctx.write(msg, promise);

        }

    } catch (EncoderException e) {

        throw e;

    } catch (Throwable e) {

        throw new EncoderException(e);

    } finally {

        if (buf != null) {

            buf.release();

        }

    }

}
</code></pre>

<p>MessageToByteEncoder 重写了 ChanneOutboundHandler 的 write() 方法，其主要逻辑分为以下几个步骤：</p>

<ol>
<li>acceptOutboundMessage 判断是否有匹配的消息类型，如果匹配需要执行编码流程，如果不匹配直接继续传递给下一个 ChannelOutboundHandler；</li>
<li>分配 ByteBuf 资源，默认使用<strong>堆外内存；</strong></li>
<li>调用子类实现的 encode 方法完成数据编码，一旦消息被成功编码，会通过调用 ReferenceCountUtil.release(cast) 自动释放；</li>
<li>如果 ByteBuf 可读，说明已经成功编码得到数据，然后写入 ChannelHandlerContext 交到下一个节点；如果 ByteBuf 不可读，则释放 ByteBuf 资源，向下传递空的 ByteBuf 对象。</li>
</ol>

<p>编码器实现非常简单，不需要关注拆包/粘包问题。如下例子，展示了如何将字符串类型的数据写入到 ByteBuf 实例，ByteBuf 实例将传递给 ChannelPipeline 链表中的下一个 ChannelOutboundHandler。</p>

<pre><code>public class StringToByteEncoder extends MessageToByteEncoder&lt;String&gt; {

        @Override

        protected void encode(ChannelHandlerContext channelHandlerContext, String data, ByteBuf byteBuf) throws Exception {

            byteBuf.writeBytes(data.getBytes());

        }

}
</code></pre>

<ul>
<li><strong>MessageToMessageEncoder</strong></li>
</ul>

<p>MessageToMessageEncoder 与 MessageToByteEncoder 类似，同样只需要实现 encode 方法。与 MessageToByteEncoder 不同的是，MessageToMessageEncoder 是将一种格式的消息转换为另外一种格式的消息。其中第二个 Message 所指的可以是任意一个对象，如果该对象是 ByteBuf 类型，那么基本上和 MessageToByteEncoder 的实现原理是一致的。此外 MessageToByteEncoder 的输出结果是对象列表，编码后的结果属于<strong>中间对象</strong>，最终仍然会转化成 ByteBuf 进行传输。</p>

<p>MessageToMessageEncoder 常用的<strong>实现子类</strong>有 StringEncoder、LineEncoder、Base64Encoder 等。以 StringEncoder 为例看下 MessageToMessageEncoder 的用法。源码示例如下：将 CharSequence 类型（String、StringBuilder、StringBuffer 等）转换成 ByteBuf 类型，结合 StringDecoder 可以直接实现 String 类型数据的编解码。</p>

<pre><code>@Override

protected void encode(ChannelHandlerContext ctx, CharSequence msg, List&lt;Object&gt; out) throws Exception {

    if (msg.length() == 0) {

        return;

    }

    out.add(ByteBufUtil.encodeString(ctx.alloc(), CharBuffer.wrap(msg), charset));

}
</code></pre>

<h4 id="抽象解码类">抽象解码类</h4>

<p>同样，我们先看下抽象解码类的继承关系图。解码类是 ChanneInboundHandler 的抽象类实现，操作的是 Inbound 入站数据。解码器实现的难度要远大于编码器，因为解码器需要考虑拆包/粘包问题。由于接收方有可能没有接收到完整的消息，所以解码框架需要对入站的数据做缓冲操作，直至获取到完整的消息。</p>

<p><img src="assets/CgqCHl-lDmiAU8l2AAPWlIx6BfA268.png" alt="Drawing 1.png" /></p>

<ul>
<li><strong>抽象解码类 ByteToMessageDecoder。</strong></li>
</ul>

<p>首先，我们看下 ByteToMessageDecoder 定义的抽象方法：</p>

<pre><code>public abstract class ByteToMessageDecoder extends ChannelInboundHandlerAdapter {

    protected abstract void decode(ChannelHandlerContext ctx, ByteBuf in, List&lt;Object&gt; out) throws Exception;

    protected void decodeLast(ChannelHandlerContext ctx, ByteBuf in, List&lt;Object&gt; out) throws Exception {

        if (in.isReadable()) {

            decodeRemovalReentryProtection(ctx, in, out);

        }

    }

}
</code></pre>

<p>decode() 是用户必须实现的抽象方法，在该方法在调用时需要传入接收的数据 ByteBuf，及用来添加编码后消息的 List。由于 TCP 粘包问题，ByteBuf 中可能包含多个有效的报文，或者不够一个完整的报文。Netty 会重复回调 decode() 方法，直到没有解码出新的完整报文可以添加到 List 当中，或者 ByteBuf 没有更多可读取的数据为止。如果此时 List 的内容不为空，那么会传递给  ChannelPipeline 中的下一个ChannelInboundHandler。</p>

<p>此外 ByteToMessageDecoder 还定义了 decodeLast() 方法。为什么抽象解码器要比编码器多一个 decodeLast() 方法呢？因为 decodeLast 在 Channel 关闭后会被调用一次，主要用于处理 ByteBuf 最后剩余的字节数据。Netty 中 decodeLast 的默认实现只是简单调用了 decode() 方法。如果有特殊的业务需求，则可以通过重写 decodeLast() 方法扩展自定义逻辑。</p>

<p>ByteToMessageDecoder 还有一个抽象子类是 ReplayingDecoder。它封装了缓冲区的管理，在读取缓冲区数据时，你无须再对字节长度进行检查。因为如果没有足够长度的字节数据，ReplayingDecoder 将终止解码操作。ReplayingDecoder 的性能相比直接使用 ByteToMessageDecoder 要慢，大部分情况下并不推荐使用 ReplayingDecoder。</p>

<ul>
<li><strong>抽象解码类 MessageToMessageDecoder。</strong></li>
</ul>

<p>MessageToMessageDecoder 与 ByteToMessageDecoder 作用类似，都是将一种消息类型的编码成另外一种消息类型。与 ByteToMessageDecoder 不同的是 MessageToMessageDecoder 并不会对数据报文进行缓存，它主要用作转换消息模型。比较推荐的做法是使用 ByteToMessageDecoder 解析 TCP 协议，解决拆包/粘包问题。解析得到有效的 ByteBuf 数据，然后传递给后续的 MessageToMessageDecoder 做数据对象的转换，具体流程如下图所示。
<img src="assets/CgqCHl-op8iAMMV_AAH0Z4-w0bM373.png" alt="Lark20201109-102121.png" /></p>

<h3 id="通信协议实战">通信协议实战</h3>

<p>在上述通信协议设计的小节内容中，我们提到了协议的基本要素并给出了一个较为通用的协议示例。下面我们通过 Netty 的编辑码框架实现该协议的解码器，加深我们对 Netty 编解码框架的理解。</p>

<p>在实现协议编码器之前，我们首先需要清楚一个问题：如何判断 ByteBuf 是否存在完整的报文？最常用的做法就是通过读取消息长度 dataLength 进行判断。如果 ByteBuf 的可读数据长度小于 dataLength，说明 ByteBuf 还不够获取一个完整的报文。在该协议前面的消息头部分包含了魔数、协议版本号、数据长度等固定字段，共 14 个字节。固定字段长度和数据长度可以作为我们判断消息完整性的依据，具体编码器实现逻辑示例如下：</p>

<pre><code>/*

+---------------------------------------------------------------+

| 魔数 2byte | 协议版本号 1byte | 序列化算法 1byte | 报文类型 1byte  |

+---------------------------------------------------------------+

| 状态 1byte |        保留字段 4byte     |      数据长度 4byte     | 

+---------------------------------------------------------------+

|                   数据内容 （长度不定）                          |

+---------------------------------------------------------------+

 */

@Override

public final void decode(ChannelHandlerContext ctx, ByteBuf in, List&lt;Object&gt; out) {

    // 判断 ByteBuf 可读取字节

    if (in.readableBytes() &lt; 14) { 

        return;

    }

    in.markReaderIndex(); // 标记 ByteBuf 读指针位置

    in.skipBytes(2); // 跳过魔数

    in.skipBytes(1); // 跳过协议版本号

    byte serializeType = in.readByte();

    in.skipBytes(1); // 跳过报文类型

    in.skipBytes(1); // 跳过状态字段

    in.skipBytes(4); // 跳过保留字段

    int dataLength = in.readInt();

    if (in.readableBytes() &lt; dataLength) {

        in.resetReaderIndex(); // 重置 ByteBuf 读指针位置

        return;

    }

    byte[] data = new byte[dataLength];

    in.readBytes(data);

    SerializeService serializeService = getSerializeServiceByType(serializeType);

    Object obj = serializeService.deserialize(data);

    if (obj != null) {

        out.add(obj);

    }

}
</code></pre>

<p>上述实现中所涉及的 ByteBuf API，在本章节就不做详细阐述了。在本专栏第三章我们会深入学习 ByteBuf。</p>

<h3 id="总结">总结</h3>

<p>本节课我们学习了协议设计的基本要素，以及如何使用 Netty 实现自定义协议。Netty 提供了一组 ChannelHandler 实现的抽象类，在项目开发中基于这些抽象类实现自定义的编解码器具备较好的可扩展性，最后通过具体示例协议的实战加深对编解码器的理解。你学会了吗？</p>

<p>当然 Netty 在编解码方面所做的工作远不止于此。它还提供了丰富的开箱即用的编解码器，下节课我们便一起探索实用的编解码技巧。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#9bf7f7f7a2afaaaaabacdbfcf6faf2f7b5f8f4f6" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f113963a84363d2',t:'MTczNDA0MzMxMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=05&#32;&#32;服务编排层：Pipeline&#32;如何协调各类&#32;Handler&#32;？>
        <link rel="icon" href="/static/favicon.png">
        <title>05  服务编排层：Pipeline 如何协调各类 Handler ？ </title>
        
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
                            <h1 id="title" data-id="05  服务编排层：Pipeline 如何协调各类 Handler ？" class="title">05  服务编排层：Pipeline 如何协调各类 Handler ？</h1>
                            <div><p>通过上节课的学习，我们知道 EventLoop 可以说是 Netty 的调度中心，负责监听多种事件类型：I/O 事件、信号事件、定时事件等，然而实际的业务处理逻辑则是由 ChannelPipeline 中所定义的 ChannelHandler 完成的，ChannelPipeline 和 ChannelHandler 也是我们在平时应用开发的过程中打交道最多的组件。Netty 服务编排层的核心组件 ChannelPipeline 和 ChannelHandler 为用户提供了 I/O 事件的全部控制权。今天这节课我们便一起深入学习 Netty 是如何利用这两个组件，将数据玩转起来。</p>

<p>在学习这节课之前，我先抛出几个问题。</p>

<ul>
<li>ChannelPipeline 与 ChannelHandler 的关系是什么？它们之间是如何协同工作的？</li>
<li>ChannelHandler 的类型有哪些？有什么区别？</li>
<li>Netty 中 I/O 事件是如何传播的？</li>
</ul>

<p>希望你在学习完本课时后，可以找到问题的答案。</p>

<h3 id="channelpipeline-概述">ChannelPipeline 概述</h3>

<p>Pipeline 的字面意思是管道、流水线。它在 Netty 中起到的作用，和一个工厂的流水线类似。原始的网络字节流经过 Pipeline ，被一步步加工包装，最后得到加工后的成品。经过前面课程核心组件的初步学习，我们已经对 ChannelPipeline 有了初步的印象：它是 Netty 的核心处理链，用以实现网络事件的动态编排和有序传播。</p>

<p>今天我们将从以下几个方面一起探讨 ChannelPipeline 的实现原理：</p>

<ul>
<li>ChannelPipeline 内部结构；</li>
<li>ChannelHandler 接口设计；</li>
<li>ChannelPipeline 事件传播机制；</li>
<li>ChannelPipeline 异常传播机制。</li>
</ul>

<h3 id="channelpipeline-内部结构">ChannelPipeline 内部结构</h3>

<p>首先我们要理清楚 ChannelPipeline 的内部结构是什么样子，这样才能理解 ChannelPipeline 的处理流程。ChannelPipeline 作为 Netty 的核心编排组件，负责调度各种类型的 ChannelHandler，实际数据的加工处理操作则是由 ChannelHandler 完成的。</p>

<p>ChannelPipeline 可以看作是 ChannelHandler 的容器载体，它是由一组 ChannelHandler 实例组成的，内部通过双向链表将不同的 ChannelHandler 链接在一起，如下图所示。当有 I/O 读写事件触发时，ChannelPipeline 会依次调用 ChannelHandler 列表对 Channel 的数据进行拦截和处理。</p>

<p><img src="assets/CgqCHl-dLiiAcORMAAYJnrq5ceE455.png" alt="image.png" /></p>

<p>由上图可知，每个 Channel 会绑定一个 ChannelPipeline，每一个 ChannelPipeline 都包含多个 ChannelHandlerContext，所有 ChannelHandlerContext 之间组成了双向链表。又因为每个 ChannelHandler 都对应一个 ChannelHandlerContext，所以实际上 ChannelPipeline 维护的是它与 ChannelHandlerContext 的关系。那么你可能会有疑问，为什么这里会多一层 ChannelHandlerContext 的封装呢？</p>

<p>其实这是一种比较常用的编程思想。ChannelHandlerContext 用于保存 ChannelHandler 上下文；ChannelHandlerContext 则包含了 ChannelHandler 生命周期的所有事件，如 connect、bind、read、flush、write、close 等。可以试想一下，如果没有 ChannelHandlerContext 的这层封装，那么我们在做 ChannelHandler 之间传递的时候，前置后置的通用逻辑就要在每个 ChannelHandler 里都实现一份。这样虽然能解决问题，但是代码结构的耦合，会非常不优雅。</p>

<p>根据网络数据的流向，ChannelPipeline 分为入站 ChannelInboundHandler 和出站 ChannelOutboundHandler 两种处理器。在客户端与服务端通信的过程中，数据从客户端发向服务端的过程叫出站，反之称为入站。数据先由一系列 InboundHandler 处理后入站，然后再由相反方向的 OutboundHandler 处理完成后出站，如下图所示。我们经常使用的解码器 Decoder 就是入站操作，编码器 Encoder 就是出站操作。服务端接收到客户端数据需要先经过 Decoder 入站处理后，再通过 Encoder 出站通知客户端。</p>

<p><img src="assets/Ciqc1F-dLm2APCjcAAPRZBy9s5c466.png" alt="image.png" /></p>

<p>接下来我们详细分析下 ChannelPipeline 双向链表的构造，ChannelPipeline 的双向链表分别维护了 HeadContext 和 TailContext 的头尾节点。我们自定义的 ChannelHandler 会插入到 Head 和 Tail 之间，这两个节点在 Netty 中已经默认实现了，它们在 ChannelPipeline 中起到了至关重要的作用。首先我们看下 HeadContext 和 TailContext 的继承关系，如下图所示。</p>

<p><img src="assets/Ciqc1F-aW9qADWwSAAndrBdsXyc104.png" alt="image.png" /></p>

<p>HeadContext 既是 Inbound 处理器，也是 Outbound 处理器。它分别实现了 ChannelInboundHandler 和 ChannelOutboundHandler。网络数据写入操作的入口就是由 HeadContext 节点完成的。HeadContext 作为 Pipeline 的头结点负责读取数据并开始传递 InBound 事件，当数据处理完成后，数据会反方向经过 Outbound 处理器，最终传递到 HeadContext，所以 HeadContext 又是处理 Outbound 事件的最后一站。此外 HeadContext 在传递事件之前，还会执行一些前置操作。</p>

<p>TailContext 只实现了 ChannelInboundHandler 接口。它会在 ChannelInboundHandler 调用链路的最后一步执行，主要用于终止 Inbound 事件传播，例如释放 Message 数据资源等。TailContext 节点作为 OutBound 事件传播的第一站，仅仅是将 OutBound 事件传递给上一个节点。</p>

<p>从整个 ChannelPipeline 调用链路来看，如果由 Channel 直接触发事件传播，那么调用链路将贯穿整个 ChannelPipeline。然而也可以在其中某一个 ChannelHandlerContext 触发同样的方法，这样只会从当前的 ChannelHandler 开始执行事件传播，该过程不会从头贯穿到尾，在一定场景下，可以提高程序性能。</p>

<h3 id="channelhandler-接口设计">ChannelHandler 接口设计</h3>

<p>在学习 ChannelPipeline 事件传播机制之前，我们需要了解 I/O 事件的生命周期。整个 ChannelHandler 是围绕 I/O 事件的生命周期所设计的，例如建立连接、读数据、写数据、连接销毁等。ChannelHandler 有两个重要的<strong>子接口</strong>：<strong>ChannelInboundHandler</strong>和<strong>ChannelOutboundHandler</strong>，分别拦截<strong>入站和出站的各种 I/O 事件</strong>。</p>

<p><strong>1. ChannelInboundHandler 的事件回调方法与触发时机。</strong></p>

<table>
<thead>
<tr>
<th>事件回调方法</th>
<th>触发时机</th>
</tr>
</thead>

<tbody>
<tr>
<td>channelRegistered</td>
<td>Channel 被注册到 EventLoop</td>
</tr>

<tr>
<td>channelUnregistered</td>
<td>Channel 从 EventLoop 中取消注册</td>
</tr>

<tr>
<td>channelActive</td>
<td>Channel 处于就绪状态，可以被读写</td>
</tr>

<tr>
<td>channelInactive</td>
<td>Channel 处于非就绪状态Channel 可以从远端读取到数据</td>
</tr>

<tr>
<td>channelRead</td>
<td>Channel 可以从远端读取到数据</td>
</tr>

<tr>
<td>channelReadComplete</td>
<td>Channel 读取数据完成</td>
</tr>

<tr>
<td>userEventTriggered</td>
<td>用户事件触发时</td>
</tr>

<tr>
<td>channelWritabilityChanged</td>
<td>Channel 的写状态发生变化</td>
</tr>
</tbody>
</table>
<p><strong>2. ChannelOutboundHandler 的事件回调方法与触发时机。</strong></p>

<p>ChannelOutboundHandler 的事件回调方法非常清晰，直接通过 ChannelOutboundHandler 的接口列表可以看到每种操作所对应的回调方法，如下图所示。这里每个回调方法都是在相应操作执行之前触发，在此就不多做赘述了。此外 ChannelOutboundHandler 中绝大部分接口都包含ChannelPromise 参数，以便于在操作完成时能够及时获得通知。</p>

<p><img src="assets/CgqCHl-aW-2AJmXxAAVxQEbkD5w806.png" alt="image" /></p>

<h3 id="事件传播机制">事件传播机制</h3>

<p>在上文中我们介绍了 ChannelPipeline 可分为入站 ChannelInboundHandler 和出站 ChannelOutboundHandler 两种处理器，与此对应传输的事件类型可以分为<strong>Inbound 事件</strong>和<strong>Outbound 事件</strong>。</p>

<p>我们通过一个代码示例，一起体验下 ChannelPipeline 的事件传播机制。</p>

<pre><code>serverBootstrap.childHandler(new ChannelInitializer&lt;SocketChannel&gt;() {

    @Override

    public void initChannel(SocketChannel ch) {

        ch.pipeline()

                .addLast(new SampleInBoundHandler(&quot;SampleInBoundHandlerA&quot;, false))

                .addLast(new SampleInBoundHandler(&quot;SampleInBoundHandlerB&quot;, false))

                .addLast(new SampleInBoundHandler(&quot;SampleInBoundHandlerC&quot;, true));

        ch.pipeline()

                .addLast(new SampleOutBoundHandler(&quot;SampleOutBoundHandlerA&quot;))

                .addLast(new SampleOutBoundHandler(&quot;SampleOutBoundHandlerB&quot;))

                .addLast(new SampleOutBoundHandler(&quot;SampleOutBoundHandlerC&quot;));
    }

}

public class SampleInBoundHandler extends ChannelInboundHandlerAdapter {

    private final String name;

    private final boolean flush;

    public SampleInBoundHandler(String name, boolean flush) {

        this.name = name;

        this.flush = flush;

    }

    @Override

    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {

        System.out.println(&quot;InBoundHandler: &quot; + name);

        if (flush) {

            ctx.channel().writeAndFlush(msg);

        } else {

            super.channelRead(ctx, msg);

        }

    }

}
public class SampleOutBoundHandler extends ChannelOutboundHandlerAdapter {

    private final String name;

    public SampleOutBoundHandler(String name) {

        this.name = name;

    }

    @Override

    public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) throws Exception {

        System.out.println(&quot;OutBoundHandler: &quot; + name);

        super.write(ctx, msg, promise);

    }

}
</code></pre>

<p>通过 Pipeline 的 addLast 方法分别添加了三个 InboundHandler 和 OutboundHandler，添加顺序都是 A -&gt; B -&gt; C，下图可以表示初始化后 ChannelPipeline 的内部结构。</p>

<p><img src="assets/CgqCHl-dLuOAPXJFAAJ3Qmmho38501.png" alt="image.png" /></p>

<p>当客户端向服务端发送请求时，会触发 SampleInBoundHandler 调用链的 channelRead 事件。经过 SampleInBoundHandler 调用链处理完成后，在 SampleInBoundHandlerC 中会调用 writeAndFlush 方法向客户端写回数据，此时会触发 SampleOutBoundHandler 调用链的 write 事件。最后我们看下代码示例的控制台输出：</p>

<p><img src="assets/CgqCHl-aW_yAKkKnAAWUaqNNpiI795.png" alt="image" /></p>

<p>由此可见，Inbound 事件和 Outbound 事件的传播方向是不一样的。Inbound 事件的传播方向为 Head -&gt; Tail，而 Outbound 事件传播方向是 Tail -&gt; Head，两者恰恰相反。在 Netty 应用编程中一定要理清楚事件传播的顺序。推荐你在系统设计时模拟客户端和服务端的场景画出 ChannelPipeline 的内部结构图，以避免搞混调用关系。</p>

<h3 id="异常传播机制">异常传播机制</h3>

<p>ChannelPipeline 事件传播的实现采用了经典的责任链模式，调用链路环环相扣。那么如果有一个节点处理逻辑异常会出现什么现象呢？我们通过修改 SampleInBoundHandler 的实现来模拟业务逻辑异常：</p>

<pre><code>public class SampleInBoundHandler extends ChannelInboundHandlerAdapter {

    private final String name;

    private final boolean flush;

    public SampleInBoundHandler(String name, boolean flush) {

        this.name = name;

        this.flush = flush;

    }

    @Override

    public void channelRead(ChannelHandlerContext ctx, Object msg) {

        System.out.println(&quot;InBoundHandler: &quot; + name);

        if (flush) {

            ctx.channel().writeAndFlush(msg);

        } else {

            throw new RuntimeException(&quot;InBoundHandler: &quot; + name);

        }

    }

    @Override

    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {

        System.out.println(&quot;InBoundHandlerException: &quot; + name);

        ctx.fireExceptionCaught(cause);

    }

}
</code></pre>

<p>在 channelRead 事件处理中，第一个 A 节点就会抛出 RuntimeException。同时我们重写了 ChannelInboundHandlerAdapter 中的 exceptionCaught 方法，只是在开头加上了控制台输出，方便观察异常传播的行为。下面看一下代码运行的控制台输出结果：</p>

<p><img src="assets/Ciqc1F-aXAiAV52JABzDltoTrWE345.png" alt="image" /></p>

<p>由输出结果可以看出 ctx.fireExceptionCaugh 会将异常按顺序从 Head 节点传播到 Tail 节点。如果用户没有对异常进行拦截处理，最后将由 Tail 节点统一处理，在 TailContext 源码中可以找到具体实现：</p>

<pre><code>protected void onUnhandledInboundException(Throwable cause) {

    try {

        logger.warn(

                &quot;An exceptionCaught() event was fired, and it reached at the tail of the pipeline. &quot; +

                        &quot;It usually means the last handler in the pipeline did not handle the exception.&quot;,

                cause);

    } finally {

        ReferenceCountUtil.release(cause);

    }

}
</code></pre>

<p>虽然 Netty 中 TailContext 提供了兜底的异常处理逻辑，但是在很多场景下，并不能满足我们的需求。假如你需要拦截指定的异常类型，并做出相应的异常处理，应该如何实现呢？我们接着往下看。</p>

<h3 id="异常处理的最佳实践">异常处理的最佳实践</h3>

<p>在 Netty 应用开发的过程中，良好的异常处理机制会让排查问题的过程事半功倍。所以推荐用户对异常进行统一拦截，然后根据实际业务场景实现更加完善的异常处理机制。通过异常传播机制的学习，我们应该可以想到最好的方法是在 ChannelPipeline 自定义处理器的末端添加统一的异常处理器，此时 ChannelPipeline 的内部结构如下图所示。</p>

<p><img src="assets/Ciqc1F-dLz2AMj8yAALx2oNWK94344.png" alt="image.png" /></p>

<p>用户自定义的异常处理器代码示例如下：</p>

<pre><code>public class ExceptionHandler extends ChannelDuplexHandler {

    @Override

    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {

        if (cause instanceof RuntimeException) {

            System.out.println(&quot;Handle Business Exception Success.&quot;);

        }

    }

}
</code></pre>

<p>加入统一的异常处理器后，可以看到异常已经被优雅地拦截并处理掉了。这也是 Netty 推荐的最佳异常处理实践。</p>

<p><img src="assets/CgqCHl-aXBCAS8QAAAWXhTFjQOE519.png" alt="image" /></p>

<h3 id="总结">总结</h3>

<p>本节课我们深入分析了 Pipeline 的设计原理与事件传播机制。那么课程最初我提出的几个问题你是否已经都找到答案了？我来做个简单的总结：</p>

<ul>
<li>ChannelPipeline 是双向链表结构，包含 ChannelInboundHandler 和 ChannelOutboundHandler 两种处理器。</li>
<li>ChannelHandlerContext 是对 ChannelHandler 的封装，每个 ChannelHandler 都对应一个 ChannelHandlerContext，实际上 ChannelPipeline 维护的是与 ChannelHandlerContext 的关系。</li>
<li>Inbound 事件和 Outbound 事件的传播方向相反，Inbound 事件的传播方向为 Head -&gt; Tail，而 Outbound 事件传播方向是 Tail -&gt; Head。</li>
<li>异常事件的处理顺序与 ChannelHandler 的添加顺序相同，会依次向后传播，与 Inbound 事件和 Outbound 事件无关。</li>
</ul>

<p>ChannelPipeline 精妙的设计思想值得我们学以致用，建议有兴趣的同学可以深入学习下这个组件的核心源码。在未来源码篇的课程中我们将会继续深入了解 ChannelPipeline 这个组件。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#066a6a6a3f323737363146616b676f6a2865696b" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1138c2b8b063d2',t:'MTczNDA0MzI4NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
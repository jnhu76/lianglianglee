<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;&#32;加餐：HTTP&#32;协议&#32;&#43;&#32;JSON-RPC，Dubbo&#32;跨语言就是如此简单>
        <link rel="icon" href="/static/favicon.png">
        <title>29  加餐：HTTP 协议 &#43; JSON-RPC，Dubbo 跨语言就是如此简单 </title>
        
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
                        <a class="menu-item" id="00 开篇词  深入掌握 Dubbo 原理与实现，提升你的职场竞争力.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e6%b7%b1%e5%85%a5%e6%8e%8c%e6%8f%a1%20Dubbo%20%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e7%8e%b0%ef%bc%8c%e6%8f%90%e5%8d%87%e4%bd%a0%e7%9a%84%e8%81%8c%e5%9c%ba%e7%ab%9e%e4%ba%89%e5%8a%9b.md">00 开篇词  深入掌握 Dubbo 原理与实现，提升你的职场竞争力.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  Dubbo 源码环境搭建：千里之行，始于足下.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/01%20%20Dubbo%20%e6%ba%90%e7%a0%81%e7%8e%af%e5%a2%83%e6%90%ad%e5%bb%ba%ef%bc%9a%e5%8d%83%e9%87%8c%e4%b9%8b%e8%a1%8c%ef%bc%8c%e5%a7%8b%e4%ba%8e%e8%b6%b3%e4%b8%8b.md">01  Dubbo 源码环境搭建：千里之行，始于足下.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Dubbo 的配置总线：抓住 URL，就理解了半个 Dubbo.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/02%20Dubbo%20%e7%9a%84%e9%85%8d%e7%bd%ae%e6%80%bb%e7%ba%bf%ef%bc%9a%e6%8a%93%e4%bd%8f%20URL%ef%bc%8c%e5%b0%b1%e7%90%86%e8%a7%a3%e4%ba%86%e5%8d%8a%e4%b8%aa%20Dubbo.md">02 Dubbo 的配置总线：抓住 URL，就理解了半个 Dubbo.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  Dubbo SPI 精析，接口实现两极反转（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/03%20%20Dubbo%20SPI%20%e7%b2%be%e6%9e%90%ef%bc%8c%e6%8e%a5%e5%8f%a3%e5%ae%9e%e7%8e%b0%e4%b8%a4%e6%9e%81%e5%8f%8d%e8%bd%ac%ef%bc%88%e4%b8%8a%ef%bc%89.md">03  Dubbo SPI 精析，接口实现两极反转（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  Dubbo SPI 精析，接口实现两极反转（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/04%20%20Dubbo%20SPI%20%e7%b2%be%e6%9e%90%ef%bc%8c%e6%8e%a5%e5%8f%a3%e5%ae%9e%e7%8e%b0%e4%b8%a4%e6%9e%81%e5%8f%8d%e8%bd%ac%ef%bc%88%e4%b8%8b%ef%bc%89.md">04  Dubbo SPI 精析，接口实现两极反转（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  海量定时任务，一个时间轮搞定.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/05%20%20%e6%b5%b7%e9%87%8f%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1%ef%bc%8c%e4%b8%80%e4%b8%aa%e6%97%b6%e9%97%b4%e8%bd%ae%e6%90%9e%e5%ae%9a.md">05  海量定时任务，一个时间轮搞定.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  ZooKeeper 与 Curator，求你别用 ZkClient 了（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/06%20%20ZooKeeper%20%e4%b8%8e%20Curator%ef%bc%8c%e6%b1%82%e4%bd%a0%e5%88%ab%e7%94%a8%20ZkClient%20%e4%ba%86%ef%bc%88%e4%b8%8a%ef%bc%89.md">06  ZooKeeper 与 Curator，求你别用 ZkClient 了（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  ZooKeeper 与 Curator，求你别用 ZkClient 了（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/07%20%20ZooKeeper%20%e4%b8%8e%20Curator%ef%bc%8c%e6%b1%82%e4%bd%a0%e5%88%ab%e7%94%a8%20ZkClient%20%e4%ba%86%ef%bc%88%e4%b8%8b%ef%bc%89.md">07  ZooKeeper 与 Curator，求你别用 ZkClient 了（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  代理模式与常见实现.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/08%20%20%e4%bb%a3%e7%90%86%e6%a8%a1%e5%bc%8f%e4%b8%8e%e5%b8%b8%e8%a7%81%e5%ae%9e%e7%8e%b0.md">08  代理模式与常见实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  Netty 入门，用它做网络编程都说好（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/09%20%20Netty%20%e5%85%a5%e9%97%a8%ef%bc%8c%e7%94%a8%e5%ae%83%e5%81%9a%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e9%83%bd%e8%af%b4%e5%a5%bd%ef%bc%88%e4%b8%8a%ef%bc%89.md">09  Netty 入门，用它做网络编程都说好（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  Netty 入门，用它做网络编程都说好（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/10%20%20Netty%20%e5%85%a5%e9%97%a8%ef%bc%8c%e7%94%a8%e5%ae%83%e5%81%9a%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e9%83%bd%e8%af%b4%e5%a5%bd%ef%bc%88%e4%b8%8b%ef%bc%89.md">10  Netty 入门，用它做网络编程都说好（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  简易版 RPC 框架实现（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/11%20%20%e7%ae%80%e6%98%93%e7%89%88%20RPC%20%e6%a1%86%e6%9e%b6%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8a%ef%bc%89.md">11  简易版 RPC 框架实现（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  简易版 RPC 框架实现（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/12%20%20%e7%ae%80%e6%98%93%e7%89%88%20RPC%20%e6%a1%86%e6%9e%b6%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8b%ef%bc%89.md">12  简易版 RPC 框架实现（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  本地缓存：降低 ZooKeeper 压力的一个常用手段.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/13%20%20%e6%9c%ac%e5%9c%b0%e7%bc%93%e5%ad%98%ef%bc%9a%e9%99%8d%e4%bd%8e%20ZooKeeper%20%e5%8e%8b%e5%8a%9b%e7%9a%84%e4%b8%80%e4%b8%aa%e5%b8%b8%e7%94%a8%e6%89%8b%e6%ae%b5.md">13  本地缓存：降低 ZooKeeper 压力的一个常用手段.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  重试机制是网络操作的基本保证.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/14%20%20%e9%87%8d%e8%af%95%e6%9c%ba%e5%88%b6%e6%98%af%e7%bd%91%e7%bb%9c%e6%93%8d%e4%bd%9c%e7%9a%84%e5%9f%ba%e6%9c%ac%e4%bf%9d%e8%af%81.md">14  重试机制是网络操作的基本保证.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  ZooKeeper 注册中心实现，官方推荐注册中心实践.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/15%20%20ZooKeeper%20%e6%b3%a8%e5%86%8c%e4%b8%ad%e5%bf%83%e5%ae%9e%e7%8e%b0%ef%bc%8c%e5%ae%98%e6%96%b9%e6%8e%a8%e8%8d%90%e6%b3%a8%e5%86%8c%e4%b8%ad%e5%bf%83%e5%ae%9e%e8%b7%b5.md">15  ZooKeeper 注册中心实现，官方推荐注册中心实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  Dubbo Serialize 层：多种序列化算法，总有一款适合你.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/16%20%20Dubbo%20Serialize%20%e5%b1%82%ef%bc%9a%e5%a4%9a%e7%a7%8d%e5%ba%8f%e5%88%97%e5%8c%96%e7%ae%97%e6%b3%95%ef%bc%8c%e6%80%bb%e6%9c%89%e4%b8%80%e6%ac%be%e9%80%82%e5%90%88%e4%bd%a0.md">16  Dubbo Serialize 层：多种序列化算法，总有一款适合你.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  Dubbo Remoting 层核心接口分析：这居然是一套兼容所有 NIO 框架的设计？.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/17%20%20Dubbo%20Remoting%20%e5%b1%82%e6%a0%b8%e5%bf%83%e6%8e%a5%e5%8f%a3%e5%88%86%e6%9e%90%ef%bc%9a%e8%bf%99%e5%b1%85%e7%84%b6%e6%98%af%e4%b8%80%e5%a5%97%e5%85%bc%e5%ae%b9%e6%89%80%e6%9c%89%20NIO%20%e6%a1%86%e6%9e%b6%e7%9a%84%e8%ae%be%e8%ae%a1%ef%bc%9f.md">17  Dubbo Remoting 层核心接口分析：这居然是一套兼容所有 NIO 框架的设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  Buffer 缓冲区：我们不生产数据，我们只是数据的搬运工.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/18%20%20Buffer%20%e7%bc%93%e5%86%b2%e5%8c%ba%ef%bc%9a%e6%88%91%e4%bb%ac%e4%b8%8d%e7%94%9f%e4%ba%a7%e6%95%b0%e6%8d%ae%ef%bc%8c%e6%88%91%e4%bb%ac%e5%8f%aa%e6%98%af%e6%95%b0%e6%8d%ae%e7%9a%84%e6%90%ac%e8%bf%90%e5%b7%a5.md">18  Buffer 缓冲区：我们不生产数据，我们只是数据的搬运工.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  Transporter 层核心实现：编解码与线程模型一文打尽（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/19%20%20Transporter%20%e5%b1%82%e6%a0%b8%e5%bf%83%e5%ae%9e%e7%8e%b0%ef%bc%9a%e7%bc%96%e8%a7%a3%e7%a0%81%e4%b8%8e%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b%e4%b8%80%e6%96%87%e6%89%93%e5%b0%bd%ef%bc%88%e4%b8%8a%ef%bc%89.md">19  Transporter 层核心实现：编解码与线程模型一文打尽（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  Transporter 层核心实现：编解码与线程模型一文打尽（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/20%20%20Transporter%20%e5%b1%82%e6%a0%b8%e5%bf%83%e5%ae%9e%e7%8e%b0%ef%bc%9a%e7%bc%96%e8%a7%a3%e7%a0%81%e4%b8%8e%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b%e4%b8%80%e6%96%87%e6%89%93%e5%b0%bd%ef%bc%88%e4%b8%8b%ef%bc%89.md">20  Transporter 层核心实现：编解码与线程模型一文打尽（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  Exchange 层剖析：彻底搞懂 Request-Response 模型（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/21%20%20Exchange%20%e5%b1%82%e5%89%96%e6%9e%90%ef%bc%9a%e5%bd%bb%e5%ba%95%e6%90%9e%e6%87%82%20Request-Response%20%e6%a8%a1%e5%9e%8b%ef%bc%88%e4%b8%8a%ef%bc%89.md">21  Exchange 层剖析：彻底搞懂 Request-Response 模型（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  Exchange 层剖析：彻底搞懂 Request-Response 模型（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/22%20%20Exchange%20%e5%b1%82%e5%89%96%e6%9e%90%ef%bc%9a%e5%bd%bb%e5%ba%95%e6%90%9e%e6%87%82%20Request-Response%20%e6%a8%a1%e5%9e%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">22  Exchange 层剖析：彻底搞懂 Request-Response 模型（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  核心接口介绍，RPC 层骨架梳理.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/23%20%20%e6%a0%b8%e5%bf%83%e6%8e%a5%e5%8f%a3%e4%bb%8b%e7%bb%8d%ef%bc%8cRPC%20%e5%b1%82%e9%aa%a8%e6%9e%b6%e6%a2%b3%e7%90%86.md">23  核心接口介绍，RPC 层骨架梳理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  从 Protocol 起手，看服务暴露和服务引用的全流程（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/24%20%20%e4%bb%8e%20Protocol%20%e8%b5%b7%e6%89%8b%ef%bc%8c%e7%9c%8b%e6%9c%8d%e5%8a%a1%e6%9a%b4%e9%9c%b2%e5%92%8c%e6%9c%8d%e5%8a%a1%e5%bc%95%e7%94%a8%e7%9a%84%e5%85%a8%e6%b5%81%e7%a8%8b%ef%bc%88%e4%b8%8a%ef%bc%89.md">24  从 Protocol 起手，看服务暴露和服务引用的全流程（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  从 Protocol 起手，看服务暴露和服务引用的全流程（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/25%20%20%e4%bb%8e%20Protocol%20%e8%b5%b7%e6%89%8b%ef%bc%8c%e7%9c%8b%e6%9c%8d%e5%8a%a1%e6%9a%b4%e9%9c%b2%e5%92%8c%e6%9c%8d%e5%8a%a1%e5%bc%95%e7%94%a8%e7%9a%84%e5%85%a8%e6%b5%81%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">25  从 Protocol 起手，看服务暴露和服务引用的全流程（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  加餐：直击 Dubbo “心脏”，带你一起探秘 Invoker（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/26%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e7%9b%b4%e5%87%bb%20Dubbo%20%e2%80%9c%e5%bf%83%e8%84%8f%e2%80%9d%ef%bc%8c%e5%b8%a6%e4%bd%a0%e4%b8%80%e8%b5%b7%e6%8e%a2%e7%a7%98%20Invoker%ef%bc%88%e4%b8%8a%ef%bc%89.md">26  加餐：直击 Dubbo “心脏”，带你一起探秘 Invoker（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  加餐：直击 Dubbo “心脏”，带你一起探秘 Invoker（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/27%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e7%9b%b4%e5%87%bb%20Dubbo%20%e2%80%9c%e5%bf%83%e8%84%8f%e2%80%9d%ef%bc%8c%e5%b8%a6%e4%bd%a0%e4%b8%80%e8%b5%b7%e6%8e%a2%e7%a7%98%20Invoker%ef%bc%88%e4%b8%8b%ef%bc%89.md">27  加餐：直击 Dubbo “心脏”，带你一起探秘 Invoker（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  复杂问题简单化，代理帮你隐藏了多少底层细节？.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/28%20%20%e5%a4%8d%e6%9d%82%e9%97%ae%e9%a2%98%e7%ae%80%e5%8d%95%e5%8c%96%ef%bc%8c%e4%bb%a3%e7%90%86%e5%b8%ae%e4%bd%a0%e9%9a%90%e8%97%8f%e4%ba%86%e5%a4%9a%e5%b0%91%e5%ba%95%e5%b1%82%e7%bb%86%e8%8a%82%ef%bc%9f.md">28  复杂问题简单化，代理帮你隐藏了多少底层细节？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  加餐：HTTP 协议 &#43; JSON-RPC，Dubbo 跨语言就是如此简单.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/29%20%20%e5%8a%a0%e9%a4%90%ef%bc%9aHTTP%20%e5%8d%8f%e8%ae%ae%20&#43;%20JSON-RPC%ef%bc%8cDubbo%20%e8%b7%a8%e8%af%ad%e8%a8%80%e5%b0%b1%e6%98%af%e5%a6%82%e6%ad%a4%e7%ae%80%e5%8d%95.md">29  加餐：HTTP 协议 &#43; JSON-RPC，Dubbo 跨语言就是如此简单.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  Filter 接口，扩展 Dubbo 框架的常用手段指北.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/30%20%20Filter%20%e6%8e%a5%e5%8f%a3%ef%bc%8c%e6%89%a9%e5%b1%95%20Dubbo%20%e6%a1%86%e6%9e%b6%e7%9a%84%e5%b8%b8%e7%94%a8%e6%89%8b%e6%ae%b5%e6%8c%87%e5%8c%97.md">30  Filter 接口，扩展 Dubbo 框架的常用手段指北.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  加餐：深潜 Directory 实现，探秘服务目录玄机.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/31%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e6%b7%b1%e6%bd%9c%20Directory%20%e5%ae%9e%e7%8e%b0%ef%bc%8c%e6%8e%a2%e7%a7%98%e6%9c%8d%e5%8a%a1%e7%9b%ae%e5%bd%95%e7%8e%84%e6%9c%ba.md">31  加餐：深潜 Directory 实现，探秘服务目录玄机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  路由机制：请求到底怎么走，它说了算（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/32%20%20%e8%b7%af%e7%94%b1%e6%9c%ba%e5%88%b6%ef%bc%9a%e8%af%b7%e6%b1%82%e5%88%b0%e5%ba%95%e6%80%8e%e4%b9%88%e8%b5%b0%ef%bc%8c%e5%ae%83%e8%af%b4%e4%ba%86%e7%ae%97%ef%bc%88%e4%b8%8a%ef%bc%89.md">32  路由机制：请求到底怎么走，它说了算（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  路由机制：请求到底怎么走，它说了算（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/33%20%20%e8%b7%af%e7%94%b1%e6%9c%ba%e5%88%b6%ef%bc%9a%e8%af%b7%e6%b1%82%e5%88%b0%e5%ba%95%e6%80%8e%e4%b9%88%e8%b5%b0%ef%bc%8c%e5%ae%83%e8%af%b4%e4%ba%86%e7%ae%97%ef%bc%88%e4%b8%8b%ef%bc%89.md">33  路由机制：请求到底怎么走，它说了算（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  加餐：初探 Dubbo 动态配置的那些事儿.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/34%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e5%88%9d%e6%8e%a2%20Dubbo%20%e5%8a%a8%e6%80%81%e9%85%8d%e7%bd%ae%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">34  加餐：初探 Dubbo 动态配置的那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  负载均衡：公平公正物尽其用的负载均衡策略，这里都有（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/35%20%20%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%ef%bc%9a%e5%85%ac%e5%b9%b3%e5%85%ac%e6%ad%a3%e7%89%a9%e5%b0%bd%e5%85%b6%e7%94%a8%e7%9a%84%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e7%ad%96%e7%95%a5%ef%bc%8c%e8%bf%99%e9%87%8c%e9%83%bd%e6%9c%89%ef%bc%88%e4%b8%8a%ef%bc%89.md">35  负载均衡：公平公正物尽其用的负载均衡策略，这里都有（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  负载均衡：公平公正物尽其用的负载均衡策略，这里都有（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/36%20%20%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%ef%bc%9a%e5%85%ac%e5%b9%b3%e5%85%ac%e6%ad%a3%e7%89%a9%e5%b0%bd%e5%85%b6%e7%94%a8%e7%9a%84%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e7%ad%96%e7%95%a5%ef%bc%8c%e8%bf%99%e9%87%8c%e9%83%bd%e6%9c%89%ef%bc%88%e4%b8%8b%ef%bc%89.md">36  负载均衡：公平公正物尽其用的负载均衡策略，这里都有（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  集群容错：一个好汉三个帮（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/37%20%20%e9%9b%86%e7%be%a4%e5%ae%b9%e9%94%99%ef%bc%9a%e4%b8%80%e4%b8%aa%e5%a5%bd%e6%b1%89%e4%b8%89%e4%b8%aa%e5%b8%ae%ef%bc%88%e4%b8%8a%ef%bc%89.md">37  集群容错：一个好汉三个帮（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  集群容错：一个好汉三个帮（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/38%20%20%e9%9b%86%e7%be%a4%e5%ae%b9%e9%94%99%ef%bc%9a%e4%b8%80%e4%b8%aa%e5%a5%bd%e6%b1%89%e4%b8%89%e4%b8%aa%e5%b8%ae%ef%bc%88%e4%b8%8b%ef%bc%89.md">38  集群容错：一个好汉三个帮（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  加餐：多个返回值不用怕，Merger 合并器来帮忙.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/39%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e5%a4%9a%e4%b8%aa%e8%bf%94%e5%9b%9e%e5%80%bc%e4%b8%8d%e7%94%a8%e6%80%95%ef%bc%8cMerger%20%e5%90%88%e5%b9%b6%e5%99%a8%e6%9d%a5%e5%b8%ae%e5%bf%99.md">39  加餐：多个返回值不用怕，Merger 合并器来帮忙.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  加餐：模拟远程调用，Mock 机制帮你搞定.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/40%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e6%a8%a1%e6%8b%9f%e8%bf%9c%e7%a8%8b%e8%b0%83%e7%94%a8%ef%bc%8cMock%20%e6%9c%ba%e5%88%b6%e5%b8%ae%e4%bd%a0%e6%90%9e%e5%ae%9a.md">40  加餐：模拟远程调用，Mock 机制帮你搞定.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41  加餐：一键通关服务发布全流程.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/41%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e4%b8%80%e9%94%ae%e9%80%9a%e5%85%b3%e6%9c%8d%e5%8a%a1%e5%8f%91%e5%b8%83%e5%85%a8%e6%b5%81%e7%a8%8b.md">41  加餐：一键通关服务发布全流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42  加餐：服务引用流程全解析.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/42%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%bc%95%e7%94%a8%e6%b5%81%e7%a8%8b%e5%85%a8%e8%a7%a3%e6%9e%90.md">42  加餐：服务引用流程全解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43  服务自省设计方案：新版本新方案.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/43%20%20%e6%9c%8d%e5%8a%a1%e8%87%aa%e7%9c%81%e8%ae%be%e8%ae%a1%e6%96%b9%e6%a1%88%ef%bc%9a%e6%96%b0%e7%89%88%e6%9c%ac%e6%96%b0%e6%96%b9%e6%a1%88.md">43  服务自省设计方案：新版本新方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44  元数据方案深度剖析，如何避免注册中心数据量膨胀？.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/44%20%20%e5%85%83%e6%95%b0%e6%8d%ae%e6%96%b9%e6%a1%88%e6%b7%b1%e5%ba%a6%e5%89%96%e6%9e%90%ef%bc%8c%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e6%b3%a8%e5%86%8c%e4%b8%ad%e5%bf%83%e6%95%b0%e6%8d%ae%e9%87%8f%e8%86%a8%e8%83%80%ef%bc%9f.md">44  元数据方案深度剖析，如何避免注册中心数据量膨胀？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45  加餐：深入服务自省方案中的服务发布订阅（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/45%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e6%b7%b1%e5%85%a5%e6%9c%8d%e5%8a%a1%e8%87%aa%e7%9c%81%e6%96%b9%e6%a1%88%e4%b8%ad%e7%9a%84%e6%9c%8d%e5%8a%a1%e5%8f%91%e5%b8%83%e8%ae%a2%e9%98%85%ef%bc%88%e4%b8%8a%ef%bc%89.md">45  加餐：深入服务自省方案中的服务发布订阅（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46  加餐：深入服务自省方案中的服务发布订阅（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/46%20%20%e5%8a%a0%e9%a4%90%ef%bc%9a%e6%b7%b1%e5%85%a5%e6%9c%8d%e5%8a%a1%e8%87%aa%e7%9c%81%e6%96%b9%e6%a1%88%e4%b8%ad%e7%9a%84%e6%9c%8d%e5%8a%a1%e5%8f%91%e5%b8%83%e8%ae%a2%e9%98%85%ef%bc%88%e4%b8%8b%ef%bc%89.md">46  加餐：深入服务自省方案中的服务发布订阅（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47  配置中心设计与实现：集中化配置 and 本地化配置，我都要（上）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/47%20%20%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83%e8%ae%be%e8%ae%a1%e4%b8%8e%e5%ae%9e%e7%8e%b0%ef%bc%9a%e9%9b%86%e4%b8%ad%e5%8c%96%e9%85%8d%e7%bd%ae%20and%20%e6%9c%ac%e5%9c%b0%e5%8c%96%e9%85%8d%e7%bd%ae%ef%bc%8c%e6%88%91%e9%83%bd%e8%a6%81%ef%bc%88%e4%b8%8a%ef%bc%89.md">47  配置中心设计与实现：集中化配置 and 本地化配置，我都要（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48  配置中心设计与实现：集中化配置 and 本地化配置，我都要（下）.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/48%20%20%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83%e8%ae%be%e8%ae%a1%e4%b8%8e%e5%ae%9e%e7%8e%b0%ef%bc%9a%e9%9b%86%e4%b8%ad%e5%8c%96%e9%85%8d%e7%bd%ae%20and%20%e6%9c%ac%e5%9c%b0%e5%8c%96%e9%85%8d%e7%bd%ae%ef%bc%8c%e6%88%91%e9%83%bd%e8%a6%81%ef%bc%88%e4%b8%8b%ef%bc%89.md">48  配置中心设计与实现：集中化配置 and 本地化配置，我都要（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 结束语  认真学习，缩小差距.md" href="/%e4%b8%93%e6%a0%8f/Dubbo%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/49%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e8%ae%a4%e7%9c%9f%e5%ad%a6%e4%b9%a0%ef%bc%8c%e7%bc%a9%e5%b0%8f%e5%b7%ae%e8%b7%9d.md">49 结束语  认真学习，缩小差距.md</a>
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
                            <h1 id="title" data-id="29  加餐：HTTP 协议 &#43; JSON-RPC，Dubbo 跨语言就是如此简单" class="title">29  加餐：HTTP 协议 &#43; JSON-RPC，Dubbo 跨语言就是如此简单</h1>
                            <div><p>在前面课时介绍 Protocol 和 Invoker 实现时，我们重点介绍了 AbstractProtocol 以及 DubboInvoker 实现。其实，Protocol 还有一个实现分支是 AbstractProxyProtocol，如下图所示：</p>

<p><img src="assets/CgqCHl-hFBOAU2UWAAFM9gWuXAk914.png" alt="Lark20201103-162545.png" /></p>

<p>AbstractProxyProtocol 继承关系图</p>

<p>从图中我们可以看到：gRPC、HTTP、WebService、Hessian、Thrift 等协议对应的 Protocol 实现，都是继承自 AbstractProxyProtocol 抽象类。</p>

<p>目前互联网的技术栈百花齐放，很多公司会使用 Node.js、Python、Rails、Go 等语言来开发 一些 Web 端应用，同时又有很多服务会使用 Java 技术栈实现，这就出现了大量的跨语言调用的需求。Dubbo 作为一个 RPC 框架，自然也希望能实现这种跨语言的调用，目前 Dubbo 中使用“HTTP 协议 + JSON-RPC”的方式来达到这一目的，其中 HTTP 协议和 JSON 都是天然跨语言的标准，在各种语言中都有成熟的类库。</p>

<p>下面我们就重点来分析 Dubbo 对 HTTP 协议的支持。首先，我会介绍 JSON-RPC 的基础，并通过一个示例，帮助你快速入门，然后介绍 Dubbo 中 HttpProtocol 的具体实现，也就是如何将 HTTP 协议与 JSON-RPC 结合使用，实现跨语言调用的效果。</p>

<h3 id="json-rpc">JSON-RPC</h3>

<p>Dubbo 中支持的 HTTP 协议实际上使用的是 JSON-RPC 协议。</p>

<p><strong>JSON-RPC 是基于 JSON 的跨语言远程调用协议</strong>。Dubbo 中的 dubbo-rpc-xml、dubbo-rpc-webservice 等模块支持的 XML-RPC、WebService 等协议与 JSON-RPC 一样，都是基于文本的协议，只不过 JSON 的格式比 XML、WebService 等格式更加简洁、紧凑。与 Dubbo 协议、Hessian 协议等二进制协议相比，JSON-RPC 更便于调试和实现，可见 JSON-RPC 协议还是一款非常优秀的远程调用协议。</p>

<p>在 Java 体系中，有很多成熟的 JSON-RPC 框架，例如 jsonrpc4j、jpoxy 等，其中，jsonrpc4j 本身体积小巧，使用方便，既可以独立使用，也可以与 Spring 无缝集合，非常适合基于 Spring 的项目。</p>

<p>下面我们先来看看 JSON-RPC 协议中请求的基本格式：</p>

<pre><code>{

    &quot;id&quot;:1，

    &quot;method&quot;:&quot;sayHello&quot;,

    &quot;params&quot;:[

        &quot;Dubbo json-rpc&quot;

    ]

}
</code></pre>

<p>JSON-RPC<strong>请求</strong>中各个字段的含义如下：</p>

<ul>
<li>id 字段，用于唯一标识一次远程调用。</li>
<li>method 字段，指定了调用的方法名。</li>
<li>params 数组，表示方法传入的参数，如果方法无参数传入，则传入空数组。</li>
</ul>

<p>在 JSON-RPC 的服务端收到调用请求之后，会查找到相应的方法并进行调用，然后将方法的返回值整理成如下格式，返回给客户端：</p>

<pre><code>{

    &quot;id&quot;:1，

    &quot;result&quot;:&quot;Hello Dubbo json-rpc&quot;,

    &quot;error&quot;:null

}
</code></pre>

<p>JSON-RPC<strong>响应</strong>中各个字段的含义如下：</p>

<ul>
<li>id 字段，用于唯一标识一次远程调用，该值与请求中的 id 字段值保持一致。</li>
<li>result 字段，记录了方法的返回值，若无返回值，则返回空；若调用错误，返回 null。</li>
<li>error 字段，表示调用发生异常时的异常信息，方法执行无异常时该字段为 null。</li>
</ul>

<h3 id="jsonrpc4j-基础使用">jsonrpc4j 基础使用</h3>

<p>Dubbo 使用 jsonrpc4j 库来实现 JSON-RPC 协议，下面我们使用 jsonrpc4j 编写一个简单的 JSON-RPC 服务端示例程序和客户端示例程序，并通过这两个示例程序说明 jsonrpc4j 最基本的使用方式。</p>

<p>首先，我们需要创建服务端和客户端都需要的 domain 类以及服务接口。我们先来创建一个 User 类，作为最基础的数据对象：</p>

<pre><code>public class User implements Serializable {

    private int userId;

    private String name;

    private int age;

    // 省略上述字段的getter/setter方法以及toString()方法

}
</code></pre>

<p>接下来创建一个 UserService 接口作为服务接口，其中定义了 5 个方法，分别用来创建 User、查询 User 以及相关信息、删除 User：</p>

<pre><code>public interface UserService {

    User createUser(int userId, String name, int age); 

    User getUser(int userId);

    String getUserName(int userId);

    int getUserId(String name);

    void deleteAll();

}
</code></pre>

<p>UserServiceImpl 是 UserService 接口的实现类，其中使用一个 ArrayList 集合管理 User 对象，具体实现如下：</p>

<pre><code>public class UserServiceImpl implements UserService {

    // 管理所有User对象

    private List&lt;User&gt; users = new ArrayList&lt;&gt;(); 

    @Override

    public User createUser(int userId, String name, int age) {

        System.out.println(&quot;createUser method&quot;);

        User user = new User();

        user.setUserId(userId);

        user.setName(name);

        user.setAge(age);

        users.add(user); // 创建User对象并添加到users集合中

        return user;

    }

    @Override

    public User getUser(int userId) {

        System.out.println(&quot;getUser method&quot;);

        // 根据userId从users集合中查询对应的User对象

        return users.stream().filter(u -&gt; u.getUserId() == userId).findAny().get();

    }

    @Override

    public String getUserName(int userId) {

        System.out.println(&quot;getUserName method&quot;);

        // 根据userId从users集合中查询对应的User对象之后，获取该User的name

        return getUser(userId).getName();

    }

    @Override

    public int getUserId(String name) {

        System.out.println(&quot;getUserId method&quot;);

        // 根据name从users集合中查询对应的User对象，然后获取该User的id

        return users.stream().filter(u -&gt; u.getName().equals(name)).findAny().get().getUserId();

    }

    @Override

    public void deleteAll() {

        System.out.println(&quot;deleteAll&quot;);

        users.clear(); // 清空users集合

    }

}
</code></pre>

<p>整个用户管理业务的核心大致如此。下面我们来看服务端如何将 UserService 与 JSON-RPC 关联起来。</p>

<p>首先，我们创建 RpcServlet 类，它是 HttpServlet 的子类，并覆盖了 HttpServlet 的 service() 方法。我们知道，HttpServlet 在收到 GET 和 POST 请求的时候，最终会调用其 service() 方法进行处理；HttpServlet 还会将 HTTP 请求和响应封装成 HttpServletRequest 和 HttpServletResponse 传入 service() 方法之中。这里的 RpcServlet 实现之中会创建一个 JsonRpcServer，并在 service() 方法中将 HTTP 请求委托给 JsonRpcServer 进行处理：</p>

<pre><code>public class RpcServlet extends HttpServlet {

    private JsonRpcServer rpcServer = null;

    public RpcServlet() {

        super();

        // JsonRpcServer会按照json-rpc请求，调用UserServiceImpl中的方法

        rpcServer = new JsonRpcServer(new UserServiceImpl(), UserService.class);

    }

    @Override

    protected void service(HttpServletRequest request,

                           HttpServletResponse response) throws ServletException, IOException {

        rpcServer.handle(request, response);

    }

}
</code></pre>

<p>最后，我们创建一个 JsonRpcServer 作为服务端的入口类，在其 main() 方法中会启动 Jetty 作为 Web 容器，具体实现如下：</p>

<pre><code>public class JsonRpcServer {

    public static void main(String[] args) throws Throwable {

        // 服务器的监听端口

        Server server = new Server(9999);

        // 关联一个已经存在的上下文

        WebAppContext context = new WebAppContext();

        // 设置描述符位置

        context.setDescriptor(&quot;/dubbo-demo/json-rpc-demo/src/main/webapp/WEB-INF/web.xml&quot;);

        // 设置Web内容上下文路径

        context.setResourceBase(&quot;/dubbo-demo/json-rpc-demo/src/main/webapp&quot;);

        // 设置上下文路径

        context.setContextPath(&quot;/&quot;);

        context.setParentLoaderPriority(true);

        server.setHandler(context);

        server.start();

        server.join();

    }

}
</code></pre>

<p>这里使用到的 web.xml 配置文件如下：</p>

<pre><code>&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot;?&gt;

&lt;web-app

        xmlns:xsi=&quot;http://www.w3.org/2001/XMLSchema-instance&quot;

        xmlns=&quot;http://xmlns.jcp.org/xml/ns/javaee&quot;

        xsi:schemaLocation=&quot;http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_3_1.xsd&quot;

        version=&quot;3.1&quot;&gt;

    &lt;servlet&gt;

        &lt;servlet-name&gt;RpcServlet&lt;/servlet-name&gt;

        &lt;servlet-class&gt;com.demo.RpcServlet&lt;/servlet-class&gt;

    &lt;/servlet&gt;

    &lt;servlet-mapping&gt;

        &lt;servlet-name&gt;RpcServlet&lt;/servlet-name&gt;

        &lt;url-pattern&gt;/rpc&lt;/url-pattern&gt;

    &lt;/servlet-mapping&gt;

&lt;/web-app&gt;
</code></pre>

<p><strong>完成服务端的编写之后，下面我们再继续编写 JSON-RPC 的客户端</strong>。在 JsonRpcClient 中会创建 JsonRpcHttpClient，并通过 JsonRpcHttpClient 请求服务端：</p>

<pre><code>public class JsonRpcClient {

    private static JsonRpcHttpClient rpcHttpClient;

    public static void main(String[] args) throws Throwable {

        // 创建JsonRpcHttpClient

        rpcHttpClient = new JsonRpcHttpClient(new URL(&quot;http://127.0.0.1:9999/rpc&quot;));

        JsonRpcClient jsonRpcClient = new JsonRpcClient();

        jsonRpcClient.deleteAll(); // 调用deleteAll()方法删除全部User

        // 调用createUser()方法创建User

        System.out.println(jsonRpcClient.createUser(1, &quot;testName&quot;, 30));

        // 调用getUser()、getUserName()、getUserId()方法进行查询

        System.out.println(jsonRpcClient.getUser(1));

        System.out.println(jsonRpcClient.getUserName(1));

        System.out.println(jsonRpcClient.getUserId(&quot;testName&quot;));

    }

    public void deleteAll() throws Throwable {

        // 调用服务端的deleteAll()方法

        rpcHttpClient.invoke(&quot;deleteAll&quot;, null); 

    }

    public User createUser(int userId, String name, int age) throws Throwable {

        Object[] params = new Object[]{userId, name, age};

        // 调用服务端的createUser()方法

        return rpcHttpClient.invoke(&quot;createUser&quot;, params, User.class);

    }

    public User getUser(int userId) throws Throwable {

        Integer[] params = new Integer[]{userId};

        // 调用服务端的getUser()方法

        return rpcHttpClient.invoke(&quot;getUser&quot;, params, User.class);

    }

    public String getUserName(int userId) throws Throwable {

        Integer[] params = new Integer[]{userId};

        // 调用服务端的getUserName()方法

        return rpcHttpClient.invoke(&quot;getUserName&quot;, params, String.class);

    }

    public int getUserId(String name) throws Throwable {

        String[] params = new String[]{name};

        // 调用服务端的getUserId()方法

        return rpcHttpClient.invoke(&quot;getUserId&quot;, params, Integer.class);

    }

}

// 输出：

// User{userId=1, name='testName', age=30}

// User{userId=1, name='testName', age=30}

// testName

// 1
</code></pre>

<h3 id="abstractproxyprotocol">AbstractProxyProtocol</h3>

<p>在 AbstractProxyProtocol 的 export() 方法中，首先会根据 URL 检查 exporterMap 缓存，如果查询失败，则会调用 ProxyFactory.getProxy() 方法将 Invoker 封装成业务接口的代理类，然后通过子类实现的 doExport() 方法启动底层的 ProxyProtocolServer，并初始化 serverMap 集合。具体实现如下：</p>

<pre><code>public &lt;T&gt; Exporter&lt;T&gt; export(final Invoker&lt;T&gt; invoker) throws RpcException {

    // 首先查询exporterMap集合

    final String uri = serviceKey(invoker.getUrl());

    Exporter&lt;T&gt; exporter = (Exporter&lt;T&gt;) exporterMap.get(uri);

    if (exporter != null) {

        if (Objects.equals(exporter.getInvoker().getUrl(), invoker.getUrl())) {

            return exporter;

        }

    }

    // 通过ProxyFactory创建代理类，将Invoker封装成业务接口的代理类

    final Runnable runnable = doExport(proxyFactory.getProxy(invoker, true), invoker.getInterface(), invoker.getUrl());

    // doExport()方法返回的Runnable是一个回调，其中会销毁底层的Server，将会在unexport()方法中调用该Runnable

    exporter = new AbstractExporter&lt;T&gt;(invoker) {

        public void unexport() {

            super.unexport();

            exporterMap.remove(uri);

            if (runnable != null) {

                runnable.run();

            }

        }

    };

    exporterMap.put(uri, exporter);

    return exporter;

}
</code></pre>

<p>在 HttpProtocol 的 doExport() 方法中，与前面介绍的 DubboProtocol 的实现类似，也要启动一个 RemotingServer。为了适配各种 HTTP 服务器，例如，Tomcat、Jetty 等，Dubbo 在 Transporter 层抽象出了一个 HttpServer 的接口。</p>

<p><img src="assets/Ciqc1F-hFD-ATNiiAABhOjw9PMM486.png" alt="Drawing 1.png" /></p>

<p>dubbo-remoting-http 模块位置</p>

<p>dubbo-remoting-http 模块的入口是 HttpBinder 接口，它被 @SPI 注解修饰，是一个扩展接口，有三个扩展实现，默认使用的是 JettyHttpBinder 实现，如下图所示：</p>

<p><img src="assets/Ciqc1F-hFEaAUtnPAABBFL3GfzE890.png" alt="Drawing 2.png" /></p>

<p>JettyHttpBinder 继承关系图</p>

<p>HttpBinder 接口中的 bind() 方法被 @Adaptive 注解修饰，会根据 URL 的 server 参数选择相应的 HttpBinder 扩展实现，不同 HttpBinder 实现返回相应的 HttpServer 实现。HttpServer 的继承关系如下图所示：</p>

<p><img src="assets/CgqCHl-hFFSAApv5AABUwFas2rw795.png" alt="Drawing 3.png" /></p>

<p>HttpServer 继承关系图</p>

<p>这里我们以 JettyHttpServer 为例简单介绍 HttpServer 的实现，在 JettyHttpServer 中会初始化 Jetty Server，其中会配置 Jetty Server 使用到的线程池以及处理请求 Handler：</p>

<pre><code>public JettyHttpServer(URL url, final HttpHandler handler) {

    // 初始化AbstractHttpServer中的url字段和handler字段

    super(url, handler); 

    this.url = url;

    DispatcherServlet.addHttpHandler( // 添加HttpHandler

        url.getParameter(Constants.BIND_PORT_KEY, 

        url.getPort()), handler);

    // 创建线程池

    int threads = url.getParameter(THREADS_KEY, DEFAULT_THREADS);

    QueuedThreadPool threadPool = new QueuedThreadPool();

    threadPool.setDaemon(true);

    threadPool.setMaxThreads(threads);

    threadPool.setMinThreads(threads);

    // 创建Jetty Server

    server = new Server(threadPool);

    // 创建ServerConnector，并指定绑定的ip和port

    ServerConnector connector = new ServerConnector(server);

    String bindIp = url.getParameter(Constants.BIND_IP_KEY, url.getHost());

    if (!url.isAnyHost() &amp;&amp; NetUtils.isValidLocalHost(bindIp)) {

        connector.setHost(bindIp);

    }

    connector.setPort(url.getParameter(Constants.BIND_PORT_KEY, url.getPort()));

    server.addConnector(connector);

    // 创建ServletHandler并与Jetty Server关联，由DispatcherServlet处理全部的请求

    ServletHandler servletHandler = new ServletHandler();

    ServletHolder servletHolder = servletHandler.addServletWithMapping(DispatcherServlet.class, &quot;/*&quot;);

    servletHolder.setInitOrder(2);

    // 创建ServletContextHandler并与Jetty Server关联

    ServletContextHandler context = new ServletContextHandler(server, &quot;/&quot;, ServletContextHandler.SESSIONS);

    context.setServletHandler(servletHandler);

   ServletManager.getInstance().addServletContext(url.getParameter(Constants.BIND_PORT_KEY, url.getPort()), context.getServletContext());

    server.start();

}
</code></pre>

<p>我们可以看到 JettyHttpServer 收到的全部请求将委托给 DispatcherServlet 这个 HttpServlet 实现，而 DispatcherServlet 的 service() 方法会把请求委托给对应接端口的 HttpHandler 处理：</p>

<pre><code>protected void service(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    // 从HANDLERS集合中查询端口对应的HttpHandler对象

    HttpHandler handler = HANDLERS.get(request.getLocalPort());

    if (handler == null) { // 端口没有对应的HttpHandler实现

        response.sendError(HttpServletResponse.SC_NOT_FOUND, &quot;Service not found.&quot;);

    } else { // 将请求委托给HttpHandler对象处理

        handler.handle(request, response);

    }

}
</code></pre>

<p>了解了 Dubbo 对 HttpServer 的抽象以及 JettyHttpServer 的核心之后，我们回到 HttpProtocol 中的 doExport() 方法继续分析。</p>

<p>在 HttpProtocol.doExport() 方法中会通过 HttpBinder 创建前面介绍的 HttpServer 对象，并记录到 serverMap 中用来接收 HTTP 请求。这里初始化 HttpServer 以及处理请求用到的 HttpHandler 是 HttpProtocol 中的内部类，在其他使用 HTTP 协议作为基础的 RPC 协议实现中也有类似的 HttpHandler 实现类，如下图所示：</p>

<p><img src="assets/CgqCHl-hFGCARUTkAABNZnY-dJg331.png" alt="Drawing 4.png" /></p>

<p>HttpHandler 继承关系图</p>

<p>在 HttpProtocol.InternalHandler 中的 handle() 实现中，会将请求委托给 skeletonMap 集合中记录的 JsonRpcServer 对象进行处理：</p>

<pre><code>public void handle(HttpServletRequest request, HttpServletResponse response) throws ServletException {

    String uri = request.getRequestURI();

    JsonRpcServer skeleton = skeletonMap.get(uri);

    if (cors) { ... // 处理跨域问题 }

    if (request.getMethod().equalsIgnoreCase(&quot;OPTIONS&quot;)) {

        response.setStatus(200); // 处理OPTIONS请求

    } else if (request.getMethod().equalsIgnoreCase(&quot;POST&quot;)) {

        // 只处理POST请求

        RpcContext.getContext().setRemoteAddress(

            request.getRemoteAddr(), request.getRemotePort());

        skeleton.handle(request.getInputStream(),   response.getOutputStream());

    } else {// 其他Method类型的请求，例如，GET请求，直接返回500

        response.setStatus(500);

    }

}
</code></pre>

<p>skeletonMap 集合中的 JsonRpcServer 是与 HttpServer 对象一同在 doExport() 方法中初始化的。最后，我们来看 HttpProtocol.doExport() 方法的实现：</p>

<pre><code>protected &lt;T&gt; Runnable doExport(final T impl, Class&lt;T&gt; type, URL url) throws RpcException {

    String addr = getAddr(url);

    // 先查询serverMap缓存

    ProtocolServer protocolServer = serverMap.get(addr);

    if (protocolServer == null) { // 查询缓存失败

        // 创建HttpServer,注意，传入的HttpHandler实现是InternalHandler

        RemotingServer remotingServer = httpBinder.bind(url, new InternalHandler(url.getParameter(&quot;cors&quot;, false)));

        serverMap.put(addr, new ProxyProtocolServer(remotingServer));

    }

    // 创建JsonRpcServer对象，并将URL与JsonRpcServer的映射关系记录到skeletonMap集合中

    final String path = url.getAbsolutePath();

    final String genericPath = path + &quot;/&quot; + GENERIC_KEY;

    JsonRpcServer skeleton = new JsonRpcServer(impl, type);

    JsonRpcServer genericServer = new JsonRpcServer(impl, GenericService.class);

    skeletonMap.put(path, skeleton);

    skeletonMap.put(genericPath, genericServer);

    return () -&gt; {// 返回Runnable回调，在Exporter中的unexport()方法中执行

        skeletonMap.remove(path);

        skeletonMap.remove(genericPath);

    };

}
</code></pre>

<p>介绍完 HttpProtocol 暴露服务的相关实现之后，下面我们再来看 HttpProtocol 中引用服务相关的方法实现，即 protocolBindinRefer() 方法实现。该方法首先通过 doRefer() 方法创建业务接口的代理，这里会使用到 jsonrpc4j 库中的 JsonProxyFactoryBean 与 Spring 进行集成，在其 afterPropertiesSet() 方法中会创建 JsonRpcHttpClient 对象：</p>

<pre><code>public void afterPropertiesSet() {

    ... ... // 省略ObjectMapper等对象

    try {

        // 创建JsonRpcHttpClient，用于后续发送json-rpc请求

        jsonRpcHttpClient = new JsonRpcHttpClient(objectMapper, new URL(getServiceUrl()), extraHttpHeaders);

        jsonRpcHttpClient.setRequestListener(requestListener);

        jsonRpcHttpClient.setSslContext(sslContext);

        jsonRpcHttpClient.setHostNameVerifier(hostNameVerifier);

    } catch (MalformedURLException mue) {

        throw new RuntimeException(mue);

    }

}
</code></pre>

<p>下面来看 doRefer() 方法的具体实现：</p>

<pre><code>protected &lt;T&gt; T doRefer(final Class&lt;T&gt; serviceType, URL url) throws RpcException {

    final String generic = url.getParameter(GENERIC_KEY);

    final boolean isGeneric = ProtocolUtils.isGeneric(generic) || serviceType.equals(GenericService.class);

    JsonProxyFactoryBean jsonProxyFactoryBean = new JsonProxyFactoryBean();

    ... // 省略其他初始化逻辑

    jsonProxyFactoryBean.afterPropertiesSet();

    // 返回的是serviceType类型的代理对象

    return (T) jsonProxyFactoryBean.getObject(); 

}
</code></pre>

<p>在 AbstractProxyProtocol.protocolBindingRefer() 方法中，会通过 ProxyFactory.getInvoker() 方法将 doRefer() 方法返回的代理对象转换成 Invoker 对象，并记录到 Invokers 集合中，具体实现如下：</p>

<pre><code>protected &lt;T&gt; Invoker&lt;T&gt; protocolBindingRefer(final Class&lt;T&gt; type, final URL url) throws RpcException {

    final Invoker&lt;T&gt; target = proxyFactory.getInvoker(doRefer(type, url), type, url);

    Invoker&lt;T&gt; invoker = new AbstractInvoker&lt;T&gt;(type, url) {

        @Override

        protected Result doInvoke(Invocation invocation) throws Throwable {

              Result result = target.invoke(invocation);

              // 省略处理异常的逻辑

              return result;

        }

    };

    invokers.add(invoker); // 将Invoker添加到invokers集合中

    return invoker;

}
</code></pre>

<h3 id="总结">总结</h3>

<p>本课时重点介绍了在 Dubbo 中如何通过“HTTP 协议 + JSON-RPC”的方案实现跨语言调用。首先我们介绍了 JSON-RPC 中请求和响应的基本格式，以及其实现库 jsonrpc4j 的基本使用；接下来我们还详细介绍了 Dubbo 中 AbstractProxyProtocol、HttpProtocol 等核心类，剖析了 Dubbo 中“HTTP 协议 + JSON-RPC”方案的落地实现。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#610d0d0d58555050515621060c00080d4f020e0c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0d94689d94e8fe',t:'MTczNDAwNTA5NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=04&#32;ZooKeeper&#32;如何进行序列化？>
        <link rel="icon" href="/static/favicon.png">
        <title>04 ZooKeeper 如何进行序列化？ </title>
        
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
                        <a class="menu-item" id="00 开篇词：选择 ZooKeeper，一步到位掌握分布式开发.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%ef%bc%9a%e9%80%89%e6%8b%a9%20ZooKeeper%ef%bc%8c%e4%b8%80%e6%ad%a5%e5%88%b0%e4%bd%8d%e6%8e%8c%e6%8f%a1%e5%88%86%e5%b8%83%e5%bc%8f%e5%bc%80%e5%8f%91.md">00 开篇词：选择 ZooKeeper，一步到位掌握分布式开发.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 ZooKeeper 数据模型：节点的特性与应用.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/01%20ZooKeeper%20%e6%95%b0%e6%8d%ae%e6%a8%a1%e5%9e%8b%ef%bc%9a%e8%8a%82%e7%82%b9%e7%9a%84%e7%89%b9%e6%80%a7%e4%b8%8e%e5%ba%94%e7%94%a8.md">01 ZooKeeper 数据模型：节点的特性与应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 发布订阅模式：如何使用 Watch 机制实现分布式通知.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/02%20%e5%8f%91%e5%b8%83%e8%ae%a2%e9%98%85%e6%a8%a1%e5%bc%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Watch%20%e6%9c%ba%e5%88%b6%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e9%80%9a%e7%9f%a5.md">02 发布订阅模式：如何使用 Watch 机制实现分布式通知.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 ACL 权限控制：如何避免未经授权的访问？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/03%20ACL%20%e6%9d%83%e9%99%90%e6%8e%a7%e5%88%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e6%9c%aa%e7%bb%8f%e6%8e%88%e6%9d%83%e7%9a%84%e8%ae%bf%e9%97%ae%ef%bc%9f.md">03 ACL 权限控制：如何避免未经授权的访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 ZooKeeper 如何进行序列化？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/04%20ZooKeeper%20%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e5%ba%8f%e5%88%97%e5%8c%96%ef%bc%9f.md">04 ZooKeeper 如何进行序列化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 深入分析 Jute 的底层实现原理.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/05%20%e6%b7%b1%e5%85%a5%e5%88%86%e6%9e%90%20Jute%20%e7%9a%84%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">05 深入分析 Jute 的底层实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 ZooKeeper 的网络通信协议详解.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/06%20ZooKeeper%20%e7%9a%84%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1%e5%8d%8f%e8%ae%ae%e8%af%a6%e8%a7%a3.md">06 ZooKeeper 的网络通信协议详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 单机模式：服务器如何从初始化到对外提供服务？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/07%20%e5%8d%95%e6%9c%ba%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%99%a8%e5%a6%82%e4%bd%95%e4%bb%8e%e5%88%9d%e5%a7%8b%e5%8c%96%e5%88%b0%e5%af%b9%e5%a4%96%e6%8f%90%e4%be%9b%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">07 单机模式：服务器如何从初始化到对外提供服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 集群模式：服务器如何从初始化到对外提供服务？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/08%20%e9%9b%86%e7%be%a4%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%99%a8%e5%a6%82%e4%bd%95%e4%bb%8e%e5%88%9d%e5%a7%8b%e5%8c%96%e5%88%b0%e5%af%b9%e5%a4%96%e6%8f%90%e4%be%9b%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">08 集群模式：服务器如何从初始化到对外提供服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 创建会话：避开日常开发的那些“坑”.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/09%20%e5%88%9b%e5%bb%ba%e4%bc%9a%e8%af%9d%ef%bc%9a%e9%81%bf%e5%bc%80%e6%97%a5%e5%b8%b8%e5%bc%80%e5%8f%91%e7%9a%84%e9%82%a3%e4%ba%9b%e2%80%9c%e5%9d%91%e2%80%9d.md">09 创建会话：避开日常开发的那些“坑”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 ClientCnxn：客户端核心工作类工作原理解析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/10%20ClientCnxn%ef%bc%9a%e5%ae%a2%e6%88%b7%e7%ab%af%e6%a0%b8%e5%bf%83%e5%b7%a5%e4%bd%9c%e7%b1%bb%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86%e8%a7%a3%e6%9e%90.md">10 ClientCnxn：客户端核心工作类工作原理解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 分桶策略：如何实现高效的会话管理？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/11%20%e5%88%86%e6%a1%b6%e7%ad%96%e7%95%a5%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e7%9a%84%e4%bc%9a%e8%af%9d%e7%ae%a1%e7%90%86%ef%bc%9f.md">11 分桶策略：如何实现高效的会话管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 服务端是如何处理一次会话请求的？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/12%20%e6%9c%8d%e5%8a%a1%e7%ab%af%e6%98%af%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e4%b8%80%e6%ac%a1%e4%bc%9a%e8%af%9d%e8%af%b7%e6%b1%82%e7%9a%84%ef%bc%9f.md">12 服务端是如何处理一次会话请求的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Curator：如何降低 ZooKeeper 使用的复杂性？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/13%20Curator%ef%bc%9a%e5%a6%82%e4%bd%95%e9%99%8d%e4%bd%8e%20ZooKeeper%20%e4%bd%bf%e7%94%a8%e7%9a%84%e5%a4%8d%e6%9d%82%e6%80%a7%ef%bc%9f.md">13 Curator：如何降低 ZooKeeper 使用的复杂性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Leader 选举：如何保证分布式数据的一致性？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/14%20Leader%20%e9%80%89%e4%b8%be%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e5%88%86%e5%b8%83%e5%bc%8f%e6%95%b0%e6%8d%ae%e7%9a%84%e4%b8%80%e8%87%b4%e6%80%a7%ef%bc%9f.md">14 Leader 选举：如何保证分布式数据的一致性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 ZooKeeper 究竟是怎么选中 Leader 的？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/15%20ZooKeeper%20%e7%a9%b6%e7%ab%9f%e6%98%af%e6%80%8e%e4%b9%88%e9%80%89%e4%b8%ad%20Leader%20%e7%9a%84%ef%bc%9f.md">15 ZooKeeper 究竟是怎么选中 Leader 的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 ZooKeeper 集群中 Leader 与 Follower 的数据同步策略.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/16%20ZooKeeper%20%e9%9b%86%e7%be%a4%e4%b8%ad%20Leader%20%e4%b8%8e%20Follower%20%e7%9a%84%e6%95%b0%e6%8d%ae%e5%90%8c%e6%ad%a5%e7%ad%96%e7%95%a5.md">16 ZooKeeper 集群中 Leader 与 Follower 的数据同步策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 集群中 Leader 的作用：事务的请求处理与调度分析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/17%20%e9%9b%86%e7%be%a4%e4%b8%ad%20Leader%20%e7%9a%84%e4%bd%9c%e7%94%a8%ef%bc%9a%e4%ba%8b%e5%8a%a1%e7%9a%84%e8%af%b7%e6%b1%82%e5%a4%84%e7%90%86%e4%b8%8e%e8%b0%83%e5%ba%a6%e5%88%86%e6%9e%90.md">17 集群中 Leader 的作用：事务的请求处理与调度分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 集群中 Follow 的作用：非事务请求的处理与 Leader 的选举分析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/18%20%e9%9b%86%e7%be%a4%e4%b8%ad%20Follow%20%e7%9a%84%e4%bd%9c%e7%94%a8%ef%bc%9a%e9%9d%9e%e4%ba%8b%e5%8a%a1%e8%af%b7%e6%b1%82%e7%9a%84%e5%a4%84%e7%90%86%e4%b8%8e%20Leader%20%e7%9a%84%e9%80%89%e4%b8%be%e5%88%86%e6%9e%90.md">18 集群中 Follow 的作用：非事务请求的处理与 Leader 的选举分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Observer 的作用与 Follow 有哪些不同？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/19%20Observer%20%e7%9a%84%e4%bd%9c%e7%94%a8%e4%b8%8e%20Follow%20%e6%9c%89%e5%93%aa%e4%ba%9b%e4%b8%8d%e5%90%8c%ef%bc%9f.md">19 Observer 的作用与 Follow 有哪些不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 一个运行中的 ZooKeeper 服务会产生哪些数据和文件？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/20%20%e4%b8%80%e4%b8%aa%e8%bf%90%e8%a1%8c%e4%b8%ad%e7%9a%84%20ZooKeeper%20%e6%9c%8d%e5%8a%a1%e4%bc%9a%e4%ba%a7%e7%94%9f%e5%93%aa%e4%ba%9b%e6%95%b0%e6%8d%ae%e5%92%8c%e6%96%87%e4%bb%b6%ef%bc%9f.md">20 一个运行中的 ZooKeeper 服务会产生哪些数据和文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 ZooKeeper 分布式锁：实现和原理解析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/21%20ZooKeeper%20%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%ef%bc%9a%e5%ae%9e%e7%8e%b0%e5%92%8c%e5%8e%9f%e7%90%86%e8%a7%a3%e6%9e%90.md">21 ZooKeeper 分布式锁：实现和原理解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 基于 ZooKeeper 命名服务的应用：分布式 ID 生成器.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/22%20%e5%9f%ba%e4%ba%8e%20ZooKeeper%20%e5%91%bd%e5%90%8d%e6%9c%8d%e5%8a%a1%e7%9a%84%e5%ba%94%e7%94%a8%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%20ID%20%e7%94%9f%e6%88%90%e5%99%a8.md">22 基于 ZooKeeper 命名服务的应用：分布式 ID 生成器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 使用 ZooKeeper 实现负载均衡服务器功能.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/23%20%e4%bd%bf%e7%94%a8%20ZooKeeper%20%e5%ae%9e%e7%8e%b0%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e6%9c%8d%e5%8a%a1%e5%99%a8%e5%8a%9f%e8%83%bd.md">23 使用 ZooKeeper 实现负载均衡服务器功能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 ZooKeeper 在 Kafka 和 Dubbo 中的工业级实现案例分析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/24%20ZooKeeper%20%e5%9c%a8%20Kafka%20%e5%92%8c%20Dubbo%20%e4%b8%ad%e7%9a%84%e5%b7%a5%e4%b8%9a%e7%ba%a7%e5%ae%9e%e7%8e%b0%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90.md">24 ZooKeeper 在 Kafka 和 Dubbo 中的工业级实现案例分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 如何搭建一个高可用的 ZooKeeper 生产环境？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/25%20%e5%a6%82%e4%bd%95%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%e9%ab%98%e5%8f%af%e7%94%a8%e7%9a%84%20ZooKeeper%20%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%ef%bc%9f.md">25 如何搭建一个高可用的 ZooKeeper 生产环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 JConsole 与四字母命令：如何监控服务器上 ZooKeeper 的运行状态？.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/26%20JConsole%20%e4%b8%8e%e5%9b%9b%e5%ad%97%e6%af%8d%e5%91%bd%e4%bb%a4%ef%bc%9a%e5%a6%82%e4%bd%95%e7%9b%91%e6%8e%a7%e6%9c%8d%e5%8a%a1%e5%99%a8%e4%b8%8a%20ZooKeeper%20%e7%9a%84%e8%bf%90%e8%a1%8c%e7%8a%b6%e6%80%81%ef%bc%9f.md">26 JConsole 与四字母命令：如何监控服务器上 ZooKeeper 的运行状态？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 crontab 与 PurgeTxnLog：线上系统日志清理的最佳时间和方式.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/27%20crontab%20%e4%b8%8e%20PurgeTxnLog%ef%bc%9a%e7%ba%bf%e4%b8%8a%e7%b3%bb%e7%bb%9f%e6%97%a5%e5%bf%97%e6%b8%85%e7%90%86%e7%9a%84%e6%9c%80%e4%bd%b3%e6%97%b6%e9%97%b4%e5%92%8c%e6%96%b9%e5%bc%8f.md">27 crontab 与 PurgeTxnLog：线上系统日志清理的最佳时间和方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 彻底掌握二阶段提交三阶段提交算法原理.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/28%20%e5%bd%bb%e5%ba%95%e6%8e%8c%e6%8f%a1%e4%ba%8c%e9%98%b6%e6%ae%b5%e6%8f%90%e4%ba%a4%e4%b8%89%e9%98%b6%e6%ae%b5%e6%8f%90%e4%ba%a4%e7%ae%97%e6%b3%95%e5%8e%9f%e7%90%86.md">28 彻底掌握二阶段提交三阶段提交算法原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 ZAB 协议算法：崩溃恢复和消息广播.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/29%20ZAB%20%e5%8d%8f%e8%ae%ae%e7%ae%97%e6%b3%95%ef%bc%9a%e5%b4%a9%e6%ba%83%e6%81%a2%e5%a4%8d%e5%92%8c%e6%b6%88%e6%81%af%e5%b9%bf%e6%92%ad.md">29 ZAB 协议算法：崩溃恢复和消息广播.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 ZAB 与 Paxos 算法的联系与区别.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/30%20ZAB%20%e4%b8%8e%20Paxos%20%e7%ae%97%e6%b3%95%e7%9a%84%e8%81%94%e7%b3%bb%e4%b8%8e%e5%8c%ba%e5%88%ab.md">30 ZAB 与 Paxos 算法的联系与区别.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 ZooKeeper 中二阶段提交算法的实现分析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/31%20ZooKeeper%20%e4%b8%ad%e4%ba%8c%e9%98%b6%e6%ae%b5%e6%8f%90%e4%ba%a4%e7%ae%97%e6%b3%95%e7%9a%84%e5%ae%9e%e7%8e%b0%e5%88%86%e6%9e%90.md">31 ZooKeeper 中二阶段提交算法的实现分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 ZooKeeper 数据存储底层实现解析.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/32%20ZooKeeper%20%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0%e8%a7%a3%e6%9e%90.md">32 ZooKeeper 数据存储底层实现解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 结束语  分布技术发展与 ZooKeeper 应用前景.md" href="/%e4%b8%93%e6%a0%8f/ZooKeeper%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90%e4%b8%8e%e5%ae%9e%e6%88%98-%e5%ae%8c/33%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e5%88%86%e5%b8%83%e6%8a%80%e6%9c%af%e5%8f%91%e5%b1%95%e4%b8%8e%20ZooKeeper%20%e5%ba%94%e7%94%a8%e5%89%8d%e6%99%af.md">33 结束语  分布技术发展与 ZooKeeper 应用前景.md</a>
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
                            <h1 id="title" data-id="04 ZooKeeper 如何进行序列化？" class="title">04 ZooKeeper 如何进行序列化？</h1>
                            <div><p>通过前几课时的学习，我们大概清楚了使用 ZooKeeper 实现一些功能的主要方式，也就是通过客户端与服务端之间的相互通信。那么首先要解决的问题就是通过网络传输数据，而要想通过网络传输我们定义好的 Java 对象数据，必须要先对其进行序列化。例如，我们通过 ZooKeeper 客户端发送 ACL 权限控制请求时，需要把请求信息封装成 packet 类型，经过序列化后才能通过网络将 ACL 信息发送给 ZooKeeper 服务端进行处理。</p>

<h4 id="什么是序列化-为什么要进行序列化操作">什么是序列化，为什么要进行序列化操作</h4>

<p>序列化是指将我们定义好的 Java 类型转化成数据流的形式。之所以这么做是因为在网络传输过程中，TCP 协议采用“流通信”的方式，提供了可以读写的字节流。而这种设计的好处在于避免了在网络传输过程中经常出现的问题：比如消息丢失、消息重复和排序等问题。那么什么时候需要序列化呢？如果我们需要通过网络传递对象或将对象信息进行持久化的时候，就需要将该对象进行序列化。</p>

<p>我们较为熟悉的序列化操作是在 Java中，当我们要序列化一个对象的时候，首先要实现一个 Serializable 接口。</p>

<pre><code class="language-java">public class User implements Serializable{

  private static final long serialVersionUID = 1L;

  private Long ids;

  private String name;

  ...

}
</code></pre>

<p>实现了 Serializable 接口后其实没有做什么实际的工作，它是一个没有任何内容的空接口，起到的作用就是标识该类是需要进行序列化的，这个就与我们后边要重点讲解的 ZooKeeper 序列化实现方法有很大的不同，这里请你先记住当前的写法，后边我们会展开讲解。</p>

<pre><code class="language-java">public interface Serializable {

}
</code></pre>

<p>定义好序列化接口后，我们再看一下如何进行序列化和反序列化的操作。Java 中进行序列化和反序列化的过程中，主要用到了 ObjectInputStream 和 ObjectOutputStream 两个 IO 类。</p>

<p>ObjectOutputStream 负责将对象进行序列化并存储到本地。而 ObjectInputStream 从本地存储中读取对象信息反序列化对象。</p>

<pre><code class="language-java">//序列化

ObjectOutputStream oo = new ObjectOutputStream()

oo.writeObject(user);

//反序列化

ObjectInputStream ois = new ObjectInputStream();

User user = (User) ois.readObject();
</code></pre>

<p>到目前为止我们了解了什么是序列化，以及为什么要进行序列化，并通过我们熟悉的 Java 编程语言中的序列化实现，进一步对序列化操作有更加具体的了解。我们知道，当我们要把对象进行本地存储或网络传输时是需要进行序列化操作的，而在 ZooKeeper 中需要频繁的网络传输工作，那么在 ZooKeeper 中是如何进行序列化的呢，我们带着这个问题继续下面的学习。</p>

<h4 id="zookeeper-中的序列化方案">ZooKeeper 中的序列化方案</h4>

<p>在 ZooKeeper 中并没有采用和 Java 一样的序列化方式，而是采用了一个 Jute 的序列解决方案作为 ZooKeeper 框架自身的序列化方式，说到 Jute 框架，它最早作为 Hadoop 中的序列化组件。之后 Jute 从 Hadoop 中独立出来，成为一个独立的序列化解决方案。ZooKeeper 从最开始就采用 Jute 作为其序列化解决方案，直到其最新的版本依然没有更改。</p>

<p>虽然 ZooKeeper 一直将 Jute 框架作为序列化解决方案，但这并不意味着 Jute 相对其他框架性能更好，反倒是 Apache Avro、Thrift 等框架在性能上优于前者。之所以 ZooKeeper 一直采用 Jute 作为序列化解决方案，主要是新老版本的兼容等问题，这里也请你注意，也许在之后的版本中，ZooKeeper 会选择更加高效的序列化解决方案。</p>

<h4 id="使用-jute-实现序列化">使用 Jute 实现序列化</h4>

<p>简单介绍了 Jute 框架的发展过程，下面我们来看一下如何使用 Jute 在 ZooKeeper 中实现序列化。如果我们要想将某个定义的类进行序列化，首先需要该类实现 Record 接口的 serilize 和 deserialize 方法，这两个方法分别是序列化和反序列化方法。下边这段代码给出了我们一般在 ZooKeeper 中进行序列化的具体实现：首先，我们定义了一个 test_jute 类，为了能够对它进行序列化，需要该 test_jute 类实现 Record 接口，并在对应的 serialize 序列化方法和 deserialize 反序列化方法中编辑具体的实现逻辑。</p>

<pre><code class="language-java">class test_jute implements Record{

  private long ids；

  private String name;

  ...

  public void serialize(OutpurArchive a_,String tag){

    ...

  }

  public void deserialize(INputArchive a_,String tag){

    ...

  }

}
</code></pre>

<p>而在序列化方法 serialize 中，我们要实现的逻辑是，首先通过字符类型参数 tag 传递标记序列化标识符，之后使用 writeLong 和 writeString 等方法分别将对象属性字段进行序列化。</p>

<pre><code class="language-java">public void serialize(OutpurArchive a_,String tag) throws ...{

  a_.startRecord(this.tag);

  a_.writeLong(ids,&quot;ids&quot;);

  a_.writeString(type,&quot;name&quot;);

  a_.endRecord(this,tag);

}
</code></pre>

<p>而调用 derseralize 在实现反序列化的过程则与我们上边说的序列化过程正好相反。</p>

<pre><code class="language-java">public void deserialize(INputArchive a_,String tag) throws {

  a_.startRecord(tag);

  ids = a_.readLong(&quot;ids&quot;);

  name = a_.readString(&quot;name&quot;);

  a_.endRecord(tag);

}
</code></pre>

<p>到这里我们就介绍完了如何在 ZooKeeper 中使用 Jute 实现序列化，需要注意的是，<strong>在实现了Record 接口后，具体的序列化和反序列化逻辑要我们自己在 serialize 和 deserialize 函数中完成</strong>。</p>

<p>序列化和反序列化的实现逻辑编码方式相对固定，首先通过 startRecord 开启一段序列化操作，之后通过 writeLong、writeString 或 readLong、 readString 等方法执行序列化或反序列化。本例中只是实现了长整型和字符型的序列化和反序列化操作，除此之外 ZooKeeper 中的 Jute 框架还支持 整数类型（Int）、布尔类型（Bool）、双精度类型（Double）以及 Byte/Buffer 类型。</p>

<h4 id="jute-在-zookeeper-中的底层实现">Jute 在 ZooKeeper 中的底层实现</h4>

<p>正因为 ZooKeeper 的设计目的是将复杂的底层操作封装成简单易用的接口，从而方便用户调用，也使得我们在使用 ZooKeeper 实现序列化的时候能够更加容易。</p>

<p>学会了利用 Jute 实现序列化和反序列化后，我们深入底层，看一下 ZooKeeper 框架具体是如何实现序列化操作的。正如上边我们提到的，通过简单的实现 Record 接口就可以实现序列化，那么我们接下来就以这个接口作为入口，详细分析其底层原理。</p>

<p>Record 接口可以理解为 ZooKeeper 中专门用来进行网络传输或本地存储时使用的数据类型。因此所有我们实现的类要想传输或者存储到本地都要实现该 Record 接口。</p>

<pre><code class="language-java">public interface Record{

  public void serialize(OutputArchive archive, String tag)

    throws IOException;

public void deserialize(InputArchive archive, String tag)

    throws IOException;

}
</code></pre>

<p>Record 接口的内部实现逻辑非常简单，只是定义了一个 序列化方法 serialize 和一个反序列化方法 deserialize 。而在 Record 起到关键作用的则是两个重要的类：OutputArchive 和 InputArchive ，其实这两个类才是真正的序列化和反序列化工具类。</p>

<p>在 OutputArchive 中定义了可进行序列化的参数类型，根据不同的序列化方式调用不同的实现类进行序列化操作。如下图所示，Jute 可以通过 Binary 、 Csv 、Xml 等方式进行序列化操作。</p>

<p><img src="assets/Ciqc1F6-a8CAUVfhAABcAV_NNXw402.png" alt="image" /></p>

<p>而对应于序列化操作，在反序列化时也会相应调用不同的实现类，来进行反序列化操作。
如下图所示：</p>

<p><img src="assets/CgqCHl6-a8mAOP1YAABW8fO1GAM913.png" alt="image" /></p>

<p>注意：无论是序列化还是反序列化，都可以对多个对象进行操作，所以当我们在定义序列化和反序列化方法时，需要字符类型参数 tag 表示要序列化或反序列化哪个对象。</p>

<h4 id="总结">总结</h4>

<p>本课时介绍了什么是序列化以及为什么要进行序列化，简单来说，就是将对象编译成字节码的形式，方便将对象信息存储到本地或通过网络传输。</p>

<p>之后我们复习了 Java 中的序列化方式，与 ZooKeeper 中的序列化使用不同的是，在 Java 中，Serializable 接口是一个空接口，只是起到标识作用，实现了该接口的对象是需要进行序列化的。而在 ZooKeeper 的实现中，序列化对象需要实现的 Record 接口里边存在两个重要的方法，分别是 serialize 和 deserialize 方法。需要序列化的对象在继承 Record 接口的同时，还需要实现这两个方法。而具体的实现过程就是我们要序列化和反序列化的实现逻辑。</p>

<p>到现在我们已经对 ZooKeeper 中的序列化知识有了一个全面的掌握，请你思考一个问题：如果说序列化就是将对象转化成字节流的格式，那么为什么 ZooKeeper 的 Jute 序列化框架还提供了对 Byte/Buffer 这两种类型的序列化操作呢？</p>

<p>其实这其中最关键的作用就是在不同操作系统中存在大端和小端的问题，为了避免不同操作系统环境的差异，在传输字节类型时也要进行序列化和反序列化。这里你需要在日常使用中多注意。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d6bababaefe2e7e7e6e196b1bbb7bfbaf8b5b9bb" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12c8fa6b4363f7',t:'MTczNDA1OTY3OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
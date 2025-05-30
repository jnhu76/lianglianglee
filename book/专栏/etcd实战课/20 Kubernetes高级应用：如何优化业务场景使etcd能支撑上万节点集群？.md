<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=20&#32;Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？>
        <link rel="icon" href="/static/favicon.png">
        <title>20 Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么你要学习etcd_.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e8%a6%81%e5%ad%a6%e4%b9%a0etcd_.md">00 开篇词 为什么你要学习etcd_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 etcd的前世今生：为什么Kubernetes使用etcd？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/01%20etcd%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Kubernetes%e4%bd%bf%e7%94%a8etcd%ef%bc%9f.md">01 etcd的前世今生：为什么Kubernetes使用etcd？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 基础架构：etcd一个读请求是如何执行的？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/02%20%e5%9f%ba%e7%a1%80%e6%9e%b6%e6%9e%84%ef%bc%9aetcd%e4%b8%80%e4%b8%aa%e8%af%bb%e8%af%b7%e6%b1%82%e6%98%af%e5%a6%82%e4%bd%95%e6%89%a7%e8%a1%8c%e7%9a%84%ef%bc%9f.md">02 基础架构：etcd一个读请求是如何执行的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 基础架构：etcd一个写请求是如何执行的？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/03%20%e5%9f%ba%e7%a1%80%e6%9e%b6%e6%9e%84%ef%bc%9aetcd%e4%b8%80%e4%b8%aa%e5%86%99%e8%af%b7%e6%b1%82%e6%98%af%e5%a6%82%e4%bd%95%e6%89%a7%e8%a1%8c%e7%9a%84%ef%bc%9f.md">03 基础架构：etcd一个写请求是如何执行的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 Raft协议：etcd如何实现高可用、数据强一致的？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/04%20Raft%e5%8d%8f%e8%ae%ae%ef%bc%9aetcd%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e5%8f%af%e7%94%a8%e3%80%81%e6%95%b0%e6%8d%ae%e5%bc%ba%e4%b8%80%e8%87%b4%e7%9a%84%ef%bc%9f.md">04 Raft协议：etcd如何实现高可用、数据强一致的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 鉴权：如何保护你的数据安全？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/05%20%e9%89%b4%e6%9d%83%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bf%9d%e6%8a%a4%e4%bd%a0%e7%9a%84%e6%95%b0%e6%8d%ae%e5%ae%89%e5%85%a8%ef%bc%9f.md">05 鉴权：如何保护你的数据安全？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 租约：如何检测你的客户端存活？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/06%20%e7%a7%9f%e7%ba%a6%ef%bc%9a%e5%a6%82%e4%bd%95%e6%a3%80%e6%b5%8b%e4%bd%a0%e7%9a%84%e5%ae%a2%e6%88%b7%e7%ab%af%e5%ad%98%e6%b4%bb%ef%bc%9f.md">06 租约：如何检测你的客户端存活？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 MVCC：如何实现多版本并发控制？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/07%20MVCC%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%a4%9a%e7%89%88%e6%9c%ac%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%ef%bc%9f.md">07 MVCC：如何实现多版本并发控制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 Watch：如何高效获取数据变化通知？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/08%20Watch%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e8%8e%b7%e5%8f%96%e6%95%b0%e6%8d%ae%e5%8f%98%e5%8c%96%e9%80%9a%e7%9f%a5%ef%bc%9f.md">08 Watch：如何高效获取数据变化通知？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 事务：如何安全地实现多key操作？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/09%20%e4%ba%8b%e5%8a%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%89%e5%85%a8%e5%9c%b0%e5%ae%9e%e7%8e%b0%e5%a4%9akey%e6%93%8d%e4%bd%9c%ef%bc%9f.md">09 事务：如何安全地实现多key操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 boltdb：如何持久化存储你的key-value数据？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/10%20boltdb%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8c%81%e4%b9%85%e5%8c%96%e5%ad%98%e5%82%a8%e4%bd%a0%e7%9a%84key-value%e6%95%b0%e6%8d%ae%ef%bc%9f.md">10 boltdb：如何持久化存储你的key-value数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 压缩：如何回收旧版本数据？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/11%20%e5%8e%8b%e7%bc%a9%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9b%9e%e6%94%b6%e6%97%a7%e7%89%88%e6%9c%ac%e6%95%b0%e6%8d%ae%ef%bc%9f.md">11 压缩：如何回收旧版本数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 一致性：为什么基于Raft实现的etcd还会出现数据不一致？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/12%20%e4%b8%80%e8%87%b4%e6%80%a7%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%9f%ba%e4%ba%8eRaft%e5%ae%9e%e7%8e%b0%e7%9a%84etcd%e8%bf%98%e4%bc%9a%e5%87%ba%e7%8e%b0%e6%95%b0%e6%8d%ae%e4%b8%8d%e4%b8%80%e8%87%b4%ef%bc%9f.md">12 一致性：为什么基于Raft实现的etcd还会出现数据不一致？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 db大小：为什么etcd社区建议db大小不超过8G？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/13%20db%e5%a4%a7%e5%b0%8f%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88etcd%e7%a4%be%e5%8c%ba%e5%bb%ba%e8%ae%aedb%e5%a4%a7%e5%b0%8f%e4%b8%8d%e8%b6%85%e8%bf%878G%ef%bc%9f.md">13 db大小：为什么etcd社区建议db大小不超过8G？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 延时：为什么你的etcd请求会出现超时？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/14%20%e5%bb%b6%e6%97%b6%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e7%9a%84etcd%e8%af%b7%e6%b1%82%e4%bc%9a%e5%87%ba%e7%8e%b0%e8%b6%85%e6%97%b6%ef%bc%9f.md">14 延时：为什么你的etcd请求会出现超时？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 内存：为什么你的etcd内存占用那么高？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/15%20%e5%86%85%e5%ad%98%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e7%9a%84etcd%e5%86%85%e5%ad%98%e5%8d%a0%e7%94%a8%e9%82%a3%e4%b9%88%e9%ab%98%ef%bc%9f.md">15 内存：为什么你的etcd内存占用那么高？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 性能及稳定性（上）：如何优化及扩展etcd性能？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/16%20%e6%80%a7%e8%83%bd%e5%8f%8a%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%e5%8f%8a%e6%89%a9%e5%b1%95etcd%e6%80%a7%e8%83%bd%ef%bc%9f.md">16 性能及稳定性（上）：如何优化及扩展etcd性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 性能及稳定性（下）：如何优化及扩展etcd性能_.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/17%20%e6%80%a7%e8%83%bd%e5%8f%8a%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%e5%8f%8a%e6%89%a9%e5%b1%95etcd%e6%80%a7%e8%83%bd_.md">17 性能及稳定性（下）：如何优化及扩展etcd性能_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 实战：如何基于Raft从0到1构建一个支持多存储引擎分布式KV服务？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/18%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8eRaft%e4%bb%8e0%e5%88%b01%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e6%94%af%e6%8c%81%e5%a4%9a%e5%ad%98%e5%82%a8%e5%bc%95%e6%93%8e%e5%88%86%e5%b8%83%e5%bc%8fKV%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">18 实战：如何基于Raft从0到1构建一个支持多存储引擎分布式KV服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Kubernetes基础应用：创建一个Pod背后etcd发生了什么？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/19%20Kubernetes%e5%9f%ba%e7%a1%80%e5%ba%94%e7%94%a8%ef%bc%9a%e5%88%9b%e5%bb%ba%e4%b8%80%e4%b8%aaPod%e8%83%8c%e5%90%8eetcd%e5%8f%91%e7%94%9f%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">19 Kubernetes基础应用：创建一个Pod背后etcd发生了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/20%20Kubernetes%e9%ab%98%e7%ba%a7%e5%ba%94%e7%94%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%e4%b8%9a%e5%8a%a1%e5%9c%ba%e6%99%af%e4%bd%bfetcd%e8%83%bd%e6%94%af%e6%92%91%e4%b8%8a%e4%b8%87%e8%8a%82%e7%82%b9%e9%9b%86%e7%be%a4%ef%bc%9f.md">20 Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 分布式锁：为什么基于etcd实现分布式锁比Redis锁更安全？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/21%20%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%9f%ba%e4%ba%8eetcd%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e6%af%94Redis%e9%94%81%e6%9b%b4%e5%ae%89%e5%85%a8%ef%bc%9f.md">21 分布式锁：为什么基于etcd实现分布式锁比Redis锁更安全？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 配置及服务发现：解析etcd在API Gateway开源项目中应用.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/22%20%e9%85%8d%e7%bd%ae%e5%8f%8a%e6%9c%8d%e5%8a%a1%e5%8f%91%e7%8e%b0%ef%bc%9a%e8%a7%a3%e6%9e%90etcd%e5%9c%a8API%20Gateway%e5%bc%80%e6%ba%90%e9%a1%b9%e7%9b%ae%e4%b8%ad%e5%ba%94%e7%94%a8.md">22 配置及服务发现：解析etcd在API Gateway开源项目中应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 选型：etcd_ZooKeeper_Consul等我们该如何选择？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/23%20%e9%80%89%e5%9e%8b%ef%bc%9aetcd_ZooKeeper_Consul%e7%ad%89%e6%88%91%e4%bb%ac%e8%af%a5%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%ef%bc%9f.md">23 选型：etcd_ZooKeeper_Consul等我们该如何选择？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 运维：如何构建高可靠的etcd集群运维体系？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/24%20%e8%bf%90%e7%bb%b4%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e9%ab%98%e5%8f%af%e9%9d%a0%e7%9a%84etcd%e9%9b%86%e7%be%a4%e8%bf%90%e7%bb%b4%e4%bd%93%e7%b3%bb%ef%bc%9f.md">24 运维：如何构建高可靠的etcd集群运维体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 成员变更：为什么集群看起来正常，移除节点却会失败呢？.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e6%88%90%e5%91%98%e5%8f%98%e6%9b%b4%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9b%86%e7%be%a4%e7%9c%8b%e8%b5%b7%e6%9d%a5%e6%ad%a3%e5%b8%b8%ef%bc%8c%e7%a7%bb%e9%99%a4%e8%8a%82%e7%82%b9%e5%8d%b4%e4%bc%9a%e5%a4%b1%e8%b4%a5%e5%91%a2%ef%bc%9f.md">特别放送 成员变更：为什么集群看起来正常，移除节点却会失败呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 搞懂etcd，掌握通往分布式存储系统之门的钥匙.md" href="/%e4%b8%93%e6%a0%8f/etcd%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%90%9e%e6%87%82etcd%ef%bc%8c%e6%8e%8c%e6%8f%a1%e9%80%9a%e5%be%80%e5%88%86%e5%b8%83%e5%bc%8f%e5%ad%98%e5%82%a8%e7%b3%bb%e7%bb%9f%e4%b9%8b%e9%97%a8%e7%9a%84%e9%92%a5%e5%8c%99.md">结束语 搞懂etcd，掌握通往分布式存储系统之门的钥匙.md</a>
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
                            <h1 id="title" data-id="20 Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？" class="title">20 Kubernetes高级应用：如何优化业务场景使etcd能支撑上万节点集群？</h1>
                            <div><p>你好，我是唐聪。</p>

<p>你知道吗？ 虽然Kubernetes社区官网文档目前声称支持最大集群节点数为5000，但是云厂商已经号称支持15000节点的Kubernetes集群了，那么为什么一个小小的etcd能支撑15000节点Kubernetes集群呢？</p>

<p>今天我就和你聊聊为了支撑15000节点，Kubernetes和etcd的做的一系列优化。我将重点和你分析Kubernetes针对etcd的瓶颈是如何从应用层采取一系列优化措施，去解决大规模集群场景中各个痛点。</p>

<p>当你遇到etcd性能瓶颈时，希望这节课介绍的大规模Kubernetes集群的最佳实践经验和优化技术，能让你获得启发，帮助你解决类似问题。</p>

<h2 id="大集群核心问题分析">大集群核心问题分析</h2>

<p>在大规模Kubernetes集群中会遇到哪些问题呢？</p>

<p>大规模Kubernetes集群的外在表现是节点数成千上万，资源对象数量高达几十万。本质是更频繁地查询、写入更大的资源对象。</p>

<p>首先是查询相关问题。在大集群中最重要的就是如何最大程度地减少expensive request。因为对几十万级别的对象数量来说，按标签、namespace查询Pod，获取所有Node等场景时，很容易造成etcd和kube-apiserver OOM和丢包，乃至雪崩等问题发生。</p>

<p>其次是写入相关问题。Kubernetes为了维持上万节点的心跳，会产生大量写请求。而按照我们基础篇介绍的etcd MVCC、boltdb、线性读等原理，etcd适用场景是读多写少，大量写请求可能会导致db size持续增长、写性能达到瓶颈被限速、影响读性能。</p>

<p>最后是大资源对象相关问题。etcd适合存储较小的key-value数据，etcd本身也做了一系列硬限制，比如key的value大小默认不能超过1.5MB。</p>

<p>本讲我就和你重点分析下Kubernetes是如何优化以上问题，以实现支撑上万节点的。以及我会简单和你讲下etcd针对Kubernetes场景做了哪些优化。</p>

<h2 id="如何减少expensive-request">如何减少expensive request</h2>

<p>首先是第一个问题，Kubernetes如何减少expensive request？</p>

<p>在这个问题中，我将Kubernetes解决此问题的方案拆分成几个核心点和你分析。</p>

<h3 id="分页">分页</h3>

<p>首先List资源操作是个基本功能点。各个组件在启动的时候，都不可避免会产生List操作，从etcd获取集群资源数据，构建初始状态。因此优化的第一步就是要避免一次性读取数十万的资源操作。</p>

<p>解决方案是Kubernetes List接口支持分页特性。分页特性依赖底层存储支持，早期的etcd v2并未支持分页被饱受诟病，非常容易出现kube-apiserver大流量、高负载等问题。在etcd v3中，实现了指定返回Limit数量的范围查询，因此也赋能kube-apiserver 对外提供了分页能力。</p>

<p>如下所示，在List接口的ListOption结构体中，Limit和Continue参数就是为了实现分页特性而增加的。</p>

<p>Limit表示一次List请求最多查询的对象数量，一般为500。如果实际对象数量大于Limit，kube-apiserver则会更新ListMeta的Continue字段，client发起的下一个List请求带上这个字段就可获取下一批对象数量。直到kube-apiserver返回空的Continue值，就获取完成了整个对象结果集。</p>

<pre><code>// ListOptions is the query options to a standard REST 
list call.
type ListOptions struct {
   ...
   Limit int64 `json:&quot;limit,omitempty&quot; 
protobuf:&quot;varint,7,opt,name=limit&quot;`
   Continue string `json:&quot;continue,omitempty&quot; 
protobuf:&quot;bytes,8,opt,name=continue&quot;`
}
</code></pre>

<p>了解完kube-apiserver的分页特性后，我们接着往下看Continue字段具体含义，以及它是如何影响etcd查询结果的。</p>

<p>我们知道etcd分页是通过范围查询和Limit实现，ListOption中的Limit对应etcd查询接口中的Limit参数。你可以大胆猜测下，Continue字段是不是跟查询的范围起始key相关呢？</p>

<p>Continue字段的确包含查询范围的起始key，它本质上是个结构体，还包含APIVersion和ResourceVersion。你之所以看到的是一个奇怪字符串，那是因为kube-apiserver使用base64库对其进行了URL编码，下面是它的原始结构体。</p>

<pre><code>type continueToken struct {
   APIVersion      string `json:&quot;v&quot;`
   ResourceVersion int64  `json:&quot;rv&quot;`
   StartKey        string `json:&quot;start&quot;`
}
</code></pre>

<p>当kube-apiserver收到带Continue的分页查询时，解析Continue，获取StartKey、ResourceVersion，etcd查询Range接口指定startKey，增加clienv3.WithRange、clientv3.WithLimit、clientv3.WithRev即可。</p>

<p>当你通过分页多次查询Kubernetes资源对象，得到的最终结果集合与不带Limit查询结果是一致的吗？kube-apiserver是如何保证分页查询的一致性呢？ 这个问题我把它作为了思考题，我们一起讨论。</p>

<h3 id="资源按namespace拆分">资源按namespace拆分</h3>

<p>通过分页特性提供机制避免一次拉取大量资源对象后，接下来就是业务最佳实践上要避免同namespace存储大量资源，尽量将资源对象拆分到不同namespace下。</p>

<p>为什么拆分到不同namespace下有助于提升性能呢?</p>

<p>正如我在<a href="https://time.geekbang.org/column/article/347992" target="_blank">19</a>中所介绍的，Kubernetes资源对象存储在etcd中的key前缀包含namespace，因此它相当于是个高效的索引字段。etcd treeIndex模块从B-tree中匹配前缀时，可快速过滤出符合条件的key-value数据。</p>

<p>Kubernetes社区承诺<a href="https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md" target="_blank">SLO</a>达标的前提是，你在使用Kubernetes集群过程中必须合理配置集群和使用扩展特性，并遵循<a href="https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md" target="_blank">一系列条件限制</a>（比如同namespace下的Service数量不超过5000个）。</p>

<h3 id="informer机制">Informer机制</h3>

<p>各组件启动发起一轮List操作加载完初始状态数据后，就进入了控制器的一致性协调逻辑。在一致性协调逻辑中，在19讲Kubernetes 基础篇中，我和你介绍了Kubernetes使用的是Watch特性来获取数据变化通知，而不是List定时轮询，这也是减少List操作一大核心策略。</p>

<p>Kubernetes社区在client-go项目中提供了一个通用的Informer组件来负责client与kube-apiserver进行资源和事件同步，显著降低了开发者使用Kubernetes API、开发高性能Kubernetes扩展组件的复杂度。</p>

<p>Informer机制的Reflector封装了Watch、List操作，结合本地Cache、Indexer，实现了控制器加载完初始状态数据后，接下来的其他操作都只需要从本地缓存读取，极大降低了kube-apiserver和etcd的压力。</p>

<p>下面是Kubernetes社区给出的一个控制器使用Informer机制的架构图。黄色部分是控制器相关基础组件，蓝色部分是client-go的Informer机制的组件，它由Reflector、Queue、Informer、Indexer、Thread safe store(Local Cache)组成。</p>

<p><img src="assets/45205409d60a4715926c869c993fa5de.jpg" alt="" /></p>

<p>Informer机制的基本工作流程如下：</p>

<ul>
<li>client启动或与kube-apiserver出现连接中断再次Watch时，报&rdquo;too old resource version&rdquo;等错误后，通过Reflector组件的List操作，从kube-apiserver获取初始状态数据，随后通过Watch机制实时监听数据变化。</li>
<li>收到事件后添加到Delta FIFO队列，由Informer组件进行处理。</li>
<li>Informer将delta FIFO队列中的事件转发给Indexer组件，Indexer组件将事件持久化存储在本地的缓存中。</li>
<li>控制器开发者可通过Informer组件注册Add、Update、Delete事件的回调函数。Informer组件收到事件后会回调业务函数，比如典型的控制器使用场景，一般是将各个事件添加到WorkQueue中，控制器的各个协调goroutine从队列取出消息，解析key，通过key从Informer机制维护的本地Cache中读取数据。</li>
</ul>

<p>通过以上流程分析，你可以发现除了启动、连接中断等场景才会触发List操作，其他时候都是从本地Cache读取。</p>

<p>那连接中断等场景为什么触发client List操作呢？</p>

<h3 id="watch-bookmark机制">Watch bookmark机制</h3>

<p>要搞懂这个问题，你得了解kube-apiserver Watch特性的原理。</p>

<p>接下来我就和你介绍下Kubernetes的Watch特性。我们知道Kubernetes通过全局递增的Resource Version来实现增量数据同步逻辑，尽量避免连接中断等异常场景下client发起全量List同步操作。</p>

<p>那么在什么场景下会触发全量List同步操作呢？这就取决于client请求的Resource Version以及kube-apiserver中是否还保存了相关的历史版本数据。</p>

<p>在<a href="https://time.geekbang.org/column/article/341060" target="_blank">08</a>Watch特性中，我和你提到实现历史版本数据存储两大核心机制，滑动窗口和MVCC。与etcd v3使用MVCC机制不一样的是，Kubernetes采用的是滑动窗口机制。</p>

<p>kube-apiserver的滑动窗口机制是如何实现的呢?</p>

<p>它通过为每个类型资源（Pod,Node等）维护一个cyclic buffer，来存储最近的一系列变更事件实现。</p>

<p>下面Kubernetes核心的watchCache结构体中的cache数组、startIndex、endIndex就是用来实现cyclic buffer的。滑动窗口中的第一个元素就是cache[startIndex%capacity]，最后一个元素则是cache[endIndex%capacity]。</p>

<pre><code>// watchCache is a &quot;sliding window&quot; (with a limited capacity) of objects
// observed from a watch.
type watchCache struct {
   sync.RWMutex

   // Condition on which lists are waiting for the fresh enough
   // resource version.
   cond *sync.Cond

   // Maximum size of history window.
   capacity int

   // upper bound of capacity since event cache has a dynamic size.
   upperBoundCapacity int

   // lower bound of capacity since event cache has a dynamic size.
   lowerBoundCapacity int

   // cache is used a cyclic buffer - its first element (with the smallest
   // resourceVersion) is defined by startIndex, its last element is defined
   // by endIndex (if cache is full it will be startIndex + capacity).
   // Both startIndex and endIndex can be greater than buffer capacity -
   // you should always apply modulo capacity to get an index in cache array.
   cache      []*watchCacheEvent
   startIndex int
   endIndex   int

   // store will effectively support LIST operation from the &quot;end of cache
   // history&quot; i.e. from the moment just after the newest cached watched event.
   // It is necessary to effectively allow clients to start watching at now.
   // NOTE: We assume that &lt;store&gt; is thread-safe.
   store cache.Indexer

   // ResourceVersion up to which the watchCache is propagated.
   resourceVersion uint64
}
</code></pre>

<p>下面我以Pod资源的历史事件滑动窗口为例，和你聊聊它在什么场景可能会触发client全量List同步操作。</p>

<p>如下图所示，kube-apiserver启动后，通过List机制，加载初始Pod状态数据，随后通过Watch机制监听最新Pod数据变化。当你不断对Pod资源进行增加、删除、修改后，携带新Resource Version（简称RV）的Pod事件就会不断被加入到cyclic buffer。假设cyclic buffer容量为100，RV1是最小的一个Watch事件的Resource Version，RV 100是最大的一个Watch事件的Resource Version。</p>

<p>当版本号为RV101的Pod事件到达时，RV1就会被淘汰，kube-apiserver维护的Pod最小版本号就变成了RV2。然而在Kubernetes集群中，不少组件都只关心cyclic buffer中与自己相关的事件。</p>

<p><img src="assets/7a404110de164eeb972eacaa766dbcfe.jpg" alt="" /></p>

<p>比如图中的kubelet只关注运行在自己节点上的Pod，假设只有RV1是它关心的Pod事件版本号，在未实现Watch bookmark特性之前，其他RV2到RV101的事件是不会推送给它的，因此它内存中维护的Resource Version依然是RV1。</p>

<p>若此kubelet随后与kube-apiserver连接出现异常，它将使用版本号RV1发起Watch重连操作。但是kube-apsierver cyclic buffer中的Pod最小版本号已是RV2，因此会返回&rdquo;too old resource version&rdquo;错误给client，client只能发起List操作，在获取到最新版本号后，才能重新进入监听逻辑。</p>

<p>那么我们能否定时将最新的版本号推送给各个client来解决以上问题呢?</p>

<p>是的，这就是Kubernetes的Watch bookmark机制核心思想。即使队列中无client关注的更新事件，Informer机制的Reflector组件中Resource Version也需要更新。</p>

<p>Watch bookmark机制通过新增一个bookmark类型的事件来实现的。kube-apiserver会通过定时器将各类型资源最新的Resource Version推送给kubelet等client，在client与kube-apiserver网络异常重连等场景，大大降低了client重建Watch的开销，减少了relist expensive request。</p>

<h3 id="更高效的watch恢复机制">更高效的Watch恢复机制</h3>

<p>虽然Kubernetes社区通过Watch bookmark机制缓解了client与kube-apiserver重连等场景下可能导致的relist expensive request操作，然而在kube-apiserver重启、滚动更新时，它依然还是有可能导致大量的relist操作，这是为什么呢？ 如何进一步减少kube-apiserver重启场景下的List操作呢？</p>

<p>如下图所示，在kube-apiserver重启后，kubelet等client会立刻带上Resource Version发起重建Watch的请求。问题就在kube-apiserver重启后，watchCache中的cyclic buffer是空的，此时watchCache中的最小Resource Version(listResourceVersion)是etcd的最新全局版本号，也就是图中的RV200。</p>

<p><img src="assets/3248cc1770434fe39bcbbd050f412b1b.jpg" alt="" /></p>

<p>在不少场景下，client请求重建Watch的Resource Version是可能小于listResourceVersion的。</p>

<p>比如在上面的这个案例图中，集群内Pod稳定运行未发生变化，kubelet假设收到了最新的RV100事件。然而这个集群其他资源如ConfigMap，被管理员不断的修改，它就会导致导致etcd版本号新增，ConfigMap滑动窗口也会不断存储变更事件，从图中可以看到，它记录最大版本号为RV200。</p>

<p>因此kube-apiserver重启后，client请求重建Pod Watch的Resource Version是RV100，而Pod watchCache中的滑动窗口最小Resource Version是RV200。很显然，RV100不在Pod watchCache所维护的滑动窗口中，kube-apiserver就会返回&rdquo;too old resource version&rdquo;错误给client，client只能发起relist expensive request操作同步最新数据。</p>

<p>为了进一步降低kube-apiserver重启对client Watch中断的影响，Kubernetes在1.20版本中又进一步实现了<a href="https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption" target="_blank">更高效的Watch恢复机制</a>。它通过etcd Watch机制的Notify特性，实现了将etcd最新的版本号定时推送给kube-apiserver。kube-apiserver在将其转换成ResourceVersion后，再通过bookmark机制推送给client，避免了kube-apiserver重启后client可能发起的List操作。</p>

<h2 id="如何控制db-size">如何控制db size</h2>

<p>分析完Kubernetes如何减少expensive request，我们再看看Kubernetes是如何控制db size的。</p>

<p>首先，我们知道Kubernetes的kubelet组件会每隔10秒上报一次心跳给kube-apiserver。</p>

<p>其次，Node资源对象因为包含若干个镜像、数据卷等信息，导致Node资源对象会较大，一次心跳消息可能高达15KB以上。</p>

<p>最后，etcd是基于COW(Copy-on-write)机制实现的MVCC数据库，每次修改都会产生新的key-value，若大量写入会导致db size持续增长。</p>

<p>早期Kubernetes集群由于以上原因，当节点数成千上万时，kubelet产生的大量写请求就较容易造成db大小达到配额，无法写入。</p>

<p>那么如何解决呢？</p>

<p>本质上还是Node资源对象大的问题。实际上我们需要更新的仅仅是Node资源对象的心跳状态，而在etcd中我们存储的是整个Node资源对象，并未将心跳状态拆分出来。</p>

<p>因此Kuberentes的解决方案就是将Node资源进行拆分，把心跳状态信息从Node对象中剥离出来，通过下面的Lease对象来描述它。</p>

<pre><code>// Lease defines a lease concept.
type Lease struct {
   metav1.TypeMeta `json:&quot;,inline&quot;`
   metav1.ObjectMeta `json:&quot;metadata,omitempty&quot; protobuf:&quot;bytes,1,opt,name=metadata&quot;`
   Spec LeaseSpec `json:&quot;spec,omitempty&quot; protobuf:&quot;bytes,2,opt,name=spec&quot;`
}

// LeaseSpec is a specification of a Lease.
type LeaseSpec struct {
   HolderIdentity *string `json:&quot;holderIdentity,omitempty&quot; protobuf:&quot;bytes,1,opt,name=holderIdentity&quot;`
   LeaseDurationSeconds *int32 `json:&quot;leaseDurationSeconds,omitempty&quot; protobuf:&quot;varint,2,opt,name=leaseDurationSeconds&quot;`
   AcquireTime *metav1.MicroTime `json:&quot;acquireTime,omitempty&quot; protobuf:&quot;bytes,3,opt,name=acquireTime&quot;`
   RenewTime *metav1.MicroTime `json:&quot;renewTime,omitempty&quot; protobuf:&quot;bytes,4,opt,name=renewTime&quot;`
   LeaseTransitions *int32 `json:&quot;leaseTransitions,omitempty&quot; protobuf:&quot;varint,5,opt,name=leaseTransitions&quot;`
}
</code></pre>

<p>因为Lease对象非常小，更新的代价远小于Node对象，所以这样显著降低了kube-apiserver的CPU开销、etcd db size，Kubernetes 1.14版本后已经默认启用Node心跳切换到Lease API。</p>

<h2 id="如何优化key-value大小">如何优化key-value大小</h2>

<p>最后，我们再看看Kubernetes是如何解决etcd key-value大小限制的。</p>

<p>在成千上万个节点的集群中，一个服务可能背后有上万个Pod。而服务对应的Endpoints资源含有大量的独立的endpoints信息，这会导致Endpoints资源大小达到etcd的value大小限制，etcd拒绝更新。</p>

<p>另外，kube-proxy等组件会实时监听Endpoints资源，一个endpoint变化就会产生较大的流量，导致kube-apiserver等组件流量超大、出现一系列性能瓶颈。</p>

<p>如何解决以上Endpoints资源过大的问题呢？</p>

<p>答案依然是拆分、化大为小。Kubernetes社区设计了EndpointSlice概念，每个EndpointSlice最大支持保存100个endpoints，成功解决了key-value过大、变更同步导致流量超大等一系列瓶颈。</p>

<h2 id="etcd优化">etcd优化</h2>

<p>Kubernetes社区在解决大集群的挑战的同时，etcd社区也在不断优化、新增特性，提升etcd在Kubernetes场景下的稳定性和性能。这里我简单列举两个，一个是etcd并发读特性，一个是Watch特性的Notify机制。</p>

<h3 id="并发读特性">并发读特性</h3>

<p>通过以上介绍的各种机制、策略，虽然Kubernetes能大大缓解expensive read request问题，但是它并不是从本质上来解决问题的。</p>

<p>为什么etcd无法支持大量的read expensive request呢？</p>

<p>除了我们一直强调的容易导致OOM、大流量导致丢包外，etcd根本性瓶颈是在etcd 3.4版本之前，expensive read request会长时间持有MVCC模块的buffer读锁RLock。而写请求执行完后，需升级锁至Lock，expensive request导致写事务阻塞在升级锁过程中，最终导致写请求超时。</p>

<p>为了解决此问题，etcd 3.4版本实现了并发读特性。核心解决方案是去掉了读写锁，每个读事务拥有一个buffer。在收到读请求创建读事务对象时，全量拷贝写事务维护的buffer到读事务buffer中。</p>

<p>通过并发读特性，显著降低了List Pod和CRD等expensive read request对写性能的影响，延时不再突增、抖动。</p>

<h3 id="改善watch-notify机制">改善Watch Notify机制</h3>

<p>为了配合Kubernetes社区实现更高效的Watch恢复机制，etcd改善了Watch Notify机制，早期Notify消息发送间隔是固定的10分钟。</p>

<p>在etcd 3.4.11版本中，新增了&ndash;experimental-watch-progress-notify-interval参数使Notify间隔时间可配置，最小支持为100ms，满足了Kubernetes业务场景的诉求。</p>

<p>最后，你要注意的是，默认通过clientv3 Watch API创建的watcher是不会开启此特性的。你需要创建Watcher的时候，设置clientv3.WithProgressNotify选项，这样etcd server就会定时发送提醒消息给client，消息中就会携带etcd当前最新的全局版本号。</p>

<h2 id="小结">小结</h2>

<p>最后我们来小结下今天的内容。</p>

<p>首先我和你剖析了大集群核心问题，即expensive request、db size、key-value大小。</p>

<p>针对expensive request，我分别为你阐述了Kubernetes的分页机制、资源按namespace拆分部署策略、核心的Informer机制、优化client与kube-apiserver连接异常后的Watch恢复效率的bookmark机制、以及进一步优化kube-apiserver重建场景下Watch恢复效率的Notify机制。从这个问题优化思路中我们可以看到，优化无止境。从大方向到边界问题，Kubernetes社区一步步将expensive request降低到极致。</p>

<p>针对db size和key-value大小，Kubernetes社区的解决方案核心思想是拆分，通过Lease和EndpointSlice资源对象成功解决了大规模集群过程遇到db size和key-value瓶颈。</p>

<p>最后etcd社区也在努力提升、优化相关特性，etcd 3.4版本中的并发读特性和可配置化的Watch Notify间隔时间就是最典型的案例。自从etcd被redhat捐赠给CNCF后，etcd核心就围绕着Kubernetes社区展开工作，努力打造更快、更稳的etcd。</p>

<h2 id="思考题">思考题</h2>

<p>最后我给你留了两个思考题。</p>

<p>首先，在Kubernetes集群中，当你通过分页API分批多次查询得到全量Node资源的时候，它能保证Node全量数据的完整性、一致性（所有节点时间点一致）吗？如果能，是如何保证的呢?</p>

<p>其次，你在使用Kubernetes集群中是否有遇到一些稳定性、性能以及令你困惑的问题呢？欢迎留言和我一起讨论。</p>

<p>感谢你的阅读，如果你认为这节课的内容有收获，也欢迎把它分享给你的朋友，谢谢。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#91fdfdfda8a5a0a0a1a6d1f6fcf0f8fdbff2fefc" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12e163ef7e63f7',t:'MTczNDA2MDY3Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=01&#32;如何证明分布式系统的&#32;CAP&#32;理论？>
        <link rel="icon" href="/static/favicon.png">
        <title>01 如何证明分布式系统的 CAP 理论？ </title>
        
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
                        <a class="menu-item" id="00 开篇词：搭建分布式知识体系，挑战高薪 Offer.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%ef%bc%9a%e6%90%ad%e5%bb%ba%e5%88%86%e5%b8%83%e5%bc%8f%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%ef%bc%8c%e6%8c%91%e6%88%98%e9%ab%98%e8%96%aa%20Offer.md">00 开篇词：搭建分布式知识体系，挑战高薪 Offer.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 如何证明分布式系统的 CAP 理论？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/01%20%e5%a6%82%e4%bd%95%e8%af%81%e6%98%8e%e5%88%86%e5%b8%83%e5%bc%8f%e7%b3%bb%e7%bb%9f%e7%9a%84%20CAP%20%e7%90%86%e8%ae%ba%ef%bc%9f.md">01 如何证明分布式系统的 CAP 理论？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 不同数据一致性模型有哪些应用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/02%20%e4%b8%8d%e5%90%8c%e6%95%b0%e6%8d%ae%e4%b8%80%e8%87%b4%e6%80%a7%e6%a8%a1%e5%9e%8b%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%ef%bc%9f.md">02 不同数据一致性模型有哪些应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 如何透彻理解 Paxos 算法？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/03%20%e5%a6%82%e4%bd%95%e9%80%8f%e5%bd%bb%e7%90%86%e8%a7%a3%20Paxos%20%e7%ae%97%e6%b3%95%ef%bc%9f.md">03 如何透彻理解 Paxos 算法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 ZooKeeper 如何保证数据一致性？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/04%20ZooKeeper%20%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e6%95%b0%e6%8d%ae%e4%b8%80%e8%87%b4%e6%80%a7%ef%bc%9f.md">04 ZooKeeper 如何保证数据一致性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 共识问题：区块链如何确认记账权？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/05%20%e5%85%b1%e8%af%86%e9%97%ae%e9%a2%98%ef%bc%9a%e5%8c%ba%e5%9d%97%e9%93%be%e5%a6%82%e4%bd%95%e7%a1%ae%e8%ae%a4%e8%ae%b0%e8%b4%a6%e6%9d%83%ef%bc%9f.md">05 共识问题：区块链如何确认记账权？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 如何准备一线互联网公司面试？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/06%20%e5%a6%82%e4%bd%95%e5%87%86%e5%a4%87%e4%b8%80%e7%ba%bf%e4%ba%92%e8%81%94%e7%bd%91%e5%85%ac%e5%8f%b8%e9%9d%a2%e8%af%95%ef%bc%9f.md">06 如何准备一线互联网公司面试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 分布式事务有哪些解决方案？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/07%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%e6%9c%89%e5%93%aa%e4%ba%9b%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88%ef%bc%9f.md">07 分布式事务有哪些解决方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 对比两阶段提交，三阶段协议有哪些改进？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/08%20%e5%af%b9%e6%af%94%e4%b8%a4%e9%98%b6%e6%ae%b5%e6%8f%90%e4%ba%a4%ef%bc%8c%e4%b8%89%e9%98%b6%e6%ae%b5%e5%8d%8f%e8%ae%ae%e6%9c%89%e5%93%aa%e4%ba%9b%e6%94%b9%e8%bf%9b%ef%bc%9f.md">08 对比两阶段提交，三阶段协议有哪些改进？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 MySQL 数据库如何实现 XA 规范？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/09%20MySQL%20%e6%95%b0%e6%8d%ae%e5%ba%93%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%20XA%20%e8%a7%84%e8%8c%83%ef%bc%9f.md">09 MySQL 数据库如何实现 XA 规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 如何在业务中体现 TCC 事务模型？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/10%20%e5%a6%82%e4%bd%95%e5%9c%a8%e4%b8%9a%e5%8a%a1%e4%b8%ad%e4%bd%93%e7%8e%b0%20TCC%20%e4%ba%8b%e5%8a%a1%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">10 如何在业务中体现 TCC 事务模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 分布式锁有哪些应用场景和实现？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/11%20%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%e5%92%8c%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">11 分布式锁有哪些应用场景和实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 如何使用 Redis 快速实现分布式锁？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/12%20%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Redis%20%e5%bf%ab%e9%80%9f%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%ef%bc%9f.md">12 如何使用 Redis 快速实现分布式锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 分布式事务考点梳理 &#43; 高频面试题.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/13%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%e8%80%83%e7%82%b9%e6%a2%b3%e7%90%86%20&#43;%20%e9%ab%98%e9%a2%91%e9%9d%a2%e8%af%95%e9%a2%98.md">13 分布式事务考点梳理 &#43; 高频面试题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 如何理解 RPC 远程服务调用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/14%20%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%20RPC%20%e8%bf%9c%e7%a8%8b%e6%9c%8d%e5%8a%a1%e8%b0%83%e7%94%a8%ef%bc%9f.md">14 如何理解 RPC 远程服务调用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 为什么微服务需要 API 网关？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/15%20%e4%b8%ba%e4%bb%80%e4%b9%88%e5%be%ae%e6%9c%8d%e5%8a%a1%e9%9c%80%e8%a6%81%20API%20%e7%bd%91%e5%85%b3%ef%bc%9f.md">15 为什么微服务需要 API 网关？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 如何实现服务注册与发现？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/16%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%9c%8d%e5%8a%a1%e6%b3%a8%e5%86%8c%e4%b8%8e%e5%8f%91%e7%8e%b0%ef%bc%9f.md">16 如何实现服务注册与发现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 如何实现分布式调用跟踪？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/17%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e8%b0%83%e7%94%a8%e8%b7%9f%e8%b8%aa%ef%bc%9f.md">17 如何实现分布式调用跟踪？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 分布式下如何实现配置管理？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/18%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%8b%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%85%8d%e7%bd%ae%e7%ae%a1%e7%90%86%ef%bc%9f.md">18 分布式下如何实现配置管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 容器化升级对服务有哪些影响？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/19%20%e5%ae%b9%e5%99%a8%e5%8c%96%e5%8d%87%e7%ba%a7%e5%af%b9%e6%9c%8d%e5%8a%a1%e6%9c%89%e5%93%aa%e4%ba%9b%e5%bd%b1%e5%93%8d%ef%bc%9f.md">19 容器化升级对服务有哪些影响？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 ServiceMesh：服务网格有哪些应用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/20%20ServiceMesh%ef%bc%9a%e6%9c%8d%e5%8a%a1%e7%bd%91%e6%a0%bc%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%ef%bc%9f.md">20 ServiceMesh：服务网格有哪些应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Dubbo vs Spring Cloud：两大技术栈如何选型？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/21%20Dubbo%20vs%20Spring%20Cloud%ef%bc%9a%e4%b8%a4%e5%a4%a7%e6%8a%80%e6%9c%af%e6%a0%88%e5%a6%82%e4%bd%95%e9%80%89%e5%9e%8b%ef%bc%9f.md">21 Dubbo vs Spring Cloud：两大技术栈如何选型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 分布式服务考点梳理 &#43; 高频面试题.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/22%20%e5%88%86%e5%b8%83%e5%bc%8f%e6%9c%8d%e5%8a%a1%e8%80%83%e7%82%b9%e6%a2%b3%e7%90%86%20&#43;%20%e9%ab%98%e9%a2%91%e9%9d%a2%e8%af%95%e9%a2%98.md">22 分布式服务考点梳理 &#43; 高频面试题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 读写分离如何在业务中落地？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/23%20%e8%af%bb%e5%86%99%e5%88%86%e7%a6%bb%e5%a6%82%e4%bd%95%e5%9c%a8%e4%b8%9a%e5%8a%a1%e4%b8%ad%e8%90%bd%e5%9c%b0%ef%bc%9f.md">23 读写分离如何在业务中落地？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 为什么需要分库分表，如何实现？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/24%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8%ef%bc%8c%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">24 为什么需要分库分表，如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 存储拆分后，如何解决唯一主键问题？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/25%20%e5%ad%98%e5%82%a8%e6%8b%86%e5%88%86%e5%90%8e%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e5%94%af%e4%b8%80%e4%b8%bb%e9%94%ae%e9%97%ae%e9%a2%98%ef%bc%9f.md">25 存储拆分后，如何解决唯一主键问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 分库分表以后，如何实现扩容？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/26%20%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8%e4%bb%a5%e5%90%8e%ef%bc%8c%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%89%a9%e5%ae%b9%ef%bc%9f.md">26 分库分表以后，如何实现扩容？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 NoSQL 数据库有哪些典型应用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/27%20NoSQL%20%e6%95%b0%e6%8d%ae%e5%ba%93%e6%9c%89%e5%93%aa%e4%ba%9b%e5%85%b8%e5%9e%8b%e5%ba%94%e7%94%a8%ef%bc%9f.md">27 NoSQL 数据库有哪些典型应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 ElasticSearch 是如何建立索引的？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/28%20ElasticSearch%20%e6%98%af%e5%a6%82%e4%bd%95%e5%bb%ba%e7%ab%8b%e7%b4%a2%e5%bc%95%e7%9a%84%ef%bc%9f.md">28 ElasticSearch 是如何建立索引的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 分布式存储考点梳理 &#43; 高频面试题.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/29%20%e5%88%86%e5%b8%83%e5%bc%8f%e5%ad%98%e5%82%a8%e8%80%83%e7%82%b9%e6%a2%b3%e7%90%86%20&#43;%20%e9%ab%98%e9%a2%91%e9%9d%a2%e8%af%95%e9%a2%98.md">29 分布式存储考点梳理 &#43; 高频面试题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 消息队列有哪些应用场景？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/30%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%ef%bc%9f.md">30 消息队列有哪些应用场景？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 集群消费和广播消费有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/31%20%e9%9b%86%e7%be%a4%e6%b6%88%e8%b4%b9%e5%92%8c%e5%b9%bf%e6%92%ad%e6%b6%88%e8%b4%b9%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">31 集群消费和广播消费有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 业务上需要顺序消费，怎么保证时序性？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/32%20%e4%b8%9a%e5%8a%a1%e4%b8%8a%e9%9c%80%e8%a6%81%e9%a1%ba%e5%ba%8f%e6%b6%88%e8%b4%b9%ef%bc%8c%e6%80%8e%e4%b9%88%e4%bf%9d%e8%af%81%e6%97%b6%e5%ba%8f%e6%80%a7%ef%bc%9f.md">32 业务上需要顺序消费，怎么保证时序性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 消息幂等：如何保证消息不被重复消费？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/33%20%e6%b6%88%e6%81%af%e5%b9%82%e7%ad%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e6%b6%88%e6%81%af%e4%b8%8d%e8%a2%ab%e9%87%8d%e5%a4%8d%e6%b6%88%e8%b4%b9%ef%bc%9f.md">33 消息幂等：如何保证消息不被重复消费？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 高可用：如何实现消息队列的 HA？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/34%20%e9%ab%98%e5%8f%af%e7%94%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e7%9a%84%20HA%ef%bc%9f.md">34 高可用：如何实现消息队列的 HA？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 消息队列选型：Kafka 如何实现高性能？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/35%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%80%89%e5%9e%8b%ef%bc%9aKafka%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%80%a7%e8%83%bd%ef%bc%9f.md">35 消息队列选型：Kafka 如何实现高性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 消息队列选型：RocketMQ 适用哪些场景？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/36%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e9%80%89%e5%9e%8b%ef%bc%9aRocketMQ%20%e9%80%82%e7%94%a8%e5%93%aa%e4%ba%9b%e5%9c%ba%e6%99%af%ef%bc%9f.md">36 消息队列选型：RocketMQ 适用哪些场景？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 消息队列考点梳理 &#43; 高频面试题.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/37%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e8%80%83%e7%82%b9%e6%a2%b3%e7%90%86%20&#43;%20%e9%ab%98%e9%a2%91%e9%9d%a2%e8%af%95%e9%a2%98.md">37 消息队列考点梳理 &#43; 高频面试题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 不止业务缓存，分布式系统中还有哪些缓存？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/38%20%e4%b8%8d%e6%ad%a2%e4%b8%9a%e5%8a%a1%e7%bc%93%e5%ad%98%ef%bc%8c%e5%88%86%e5%b8%83%e5%bc%8f%e7%b3%bb%e7%bb%9f%e4%b8%ad%e8%bf%98%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bc%93%e5%ad%98%ef%bc%9f.md">38 不止业务缓存，分布式系统中还有哪些缓存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 如何避免缓存穿透、缓存击穿、缓存雪崩？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/39%20%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e7%bc%93%e5%ad%98%e7%a9%bf%e9%80%8f%e3%80%81%e7%bc%93%e5%ad%98%e5%87%bb%e7%a9%bf%e3%80%81%e7%bc%93%e5%ad%98%e9%9b%aa%e5%b4%a9%ef%bc%9f.md">39 如何避免缓存穿透、缓存击穿、缓存雪崩？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 经典问题：先更新数据库，还是先更新缓存？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/40%20%e7%bb%8f%e5%85%b8%e9%97%ae%e9%a2%98%ef%bc%9a%e5%85%88%e6%9b%b4%e6%96%b0%e6%95%b0%e6%8d%ae%e5%ba%93%ef%bc%8c%e8%bf%98%e6%98%af%e5%85%88%e6%9b%b4%e6%96%b0%e7%bc%93%e5%ad%98%ef%bc%9f.md">40 经典问题：先更新数据库，还是先更新缓存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 失效策略：缓存过期都有哪些策略？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/41%20%e5%a4%b1%e6%95%88%e7%ad%96%e7%95%a5%ef%bc%9a%e7%bc%93%e5%ad%98%e8%bf%87%e6%9c%9f%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e7%ad%96%e7%95%a5%ef%bc%9f.md">41 失效策略：缓存过期都有哪些策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 负载均衡：一致性哈希解决了哪些问题？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/42%20%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%ef%bc%9a%e4%b8%80%e8%87%b4%e6%80%a7%e5%93%88%e5%b8%8c%e8%a7%a3%e5%86%b3%e4%ba%86%e5%93%aa%e4%ba%9b%e9%97%ae%e9%a2%98%ef%bc%9f.md">42 负载均衡：一致性哈希解决了哪些问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 缓存高可用：缓存如何保证高可用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/43%20%e7%bc%93%e5%ad%98%e9%ab%98%e5%8f%af%e7%94%a8%ef%bc%9a%e7%bc%93%e5%ad%98%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e9%ab%98%e5%8f%af%e7%94%a8%ef%bc%9f.md">43 缓存高可用：缓存如何保证高可用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 分布式缓存考点梳理 &#43; 高频面试题.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/44%20%e5%88%86%e5%b8%83%e5%bc%8f%e7%bc%93%e5%ad%98%e8%80%83%e7%82%b9%e6%a2%b3%e7%90%86%20&#43;%20%e9%ab%98%e9%a2%91%e9%9d%a2%e8%af%95%e9%a2%98.md">44 分布式缓存考点梳理 &#43; 高频面试题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 从双十一看高可用的保障方式.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/45%20%e4%bb%8e%e5%8f%8c%e5%8d%81%e4%b8%80%e7%9c%8b%e9%ab%98%e5%8f%af%e7%94%a8%e7%9a%84%e4%bf%9d%e9%9a%9c%e6%96%b9%e5%bc%8f.md">45 从双十一看高可用的保障方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 高并发场景下如何实现系统限流？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/46%20%e9%ab%98%e5%b9%b6%e5%8f%91%e5%9c%ba%e6%99%af%e4%b8%8b%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%b3%bb%e7%bb%9f%e9%99%90%e6%b5%81%ef%bc%9f.md">46 高并发场景下如何实现系统限流？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 降级和熔断：如何增强服务稳定性？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/47%20%e9%99%8d%e7%ba%a7%e5%92%8c%e7%86%94%e6%96%ad%ef%bc%9a%e5%a6%82%e4%bd%95%e5%a2%9e%e5%bc%ba%e6%9c%8d%e5%8a%a1%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%9f.md">47 降级和熔断：如何增强服务稳定性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 如何选择适合业务的负载均衡策略？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/48%20%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e9%80%82%e5%90%88%e4%b8%9a%e5%8a%a1%e7%9a%84%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e7%ad%96%e7%95%a5%ef%bc%9f.md">48 如何选择适合业务的负载均衡策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 线上服务有哪些稳定性指标？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/49%20%e7%ba%bf%e4%b8%8a%e6%9c%8d%e5%8a%a1%e6%9c%89%e5%93%aa%e4%ba%9b%e7%a8%b3%e5%ae%9a%e6%80%a7%e6%8c%87%e6%a0%87%ef%bc%9f.md">49 线上服务有哪些稳定性指标？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 分布式下有哪些好用的监控组件？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/50%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%8b%e6%9c%89%e5%93%aa%e4%ba%9b%e5%a5%bd%e7%94%a8%e7%9a%84%e7%9b%91%e6%8e%a7%e7%bb%84%e4%bb%b6%ef%bc%9f.md">50 分布式下有哪些好用的监控组件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 分布式下如何实现统一日志系统？.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/51%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%8b%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%bb%9f%e4%b8%80%e6%97%a5%e5%bf%97%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">51 分布式下如何实现统一日志系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52 分布式路漫漫，厚积薄发才是王道.md" href="/%e4%b8%93%e6%a0%8f/%e5%88%86%e5%b8%83%e5%bc%8f%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%9845%e8%ae%b2-%e5%ae%8c/52%20%e5%88%86%e5%b8%83%e5%bc%8f%e8%b7%af%e6%bc%ab%e6%bc%ab%ef%bc%8c%e5%8e%9a%e7%a7%af%e8%96%84%e5%8f%91%e6%89%8d%e6%98%af%e7%8e%8b%e9%81%93.md">52 分布式路漫漫，厚积薄发才是王道.md</a>
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
                            <h1 id="title" data-id="01 如何证明分布式系统的 CAP 理论？" class="title">01 如何证明分布式系统的 CAP 理论？</h1>
                            <div><p>本课时我们主要介绍分布式系统中最基础的 CAP 理论及其应用。
对于开发或设计分布式系统的架构师、工程师来说，CAP 是必须要掌握的基础理论，CAP 理论可以帮助架构师对系统设计中目标进行取舍，合理地规划系统拆分的维度。下面我们先讲讲分布式系统的特点。</p>

<h2 id="分布式系统的特点">分布式系统的特点</h2>

<p>随着移动互联网的快速发展，互联网的用户数量越来越多，产生的数据规模也越来越大，对应用系统提出了更高的要求，我们的系统必须支持高并发访问和海量数据处理。
分布式系统技术就是用来解决集中式架构的性能瓶颈问题，来适应快速发展的业务规模，一般来说，分布式系统是建立在网络之上的硬件或者软件系统，彼此之间通过消息等方式进行通信和协调。
分布式系统的核心是<strong>可扩展性</strong>，通过对服务、存储的扩展，来提高系统的处理能力，通过对多台服务器协同工作，来完成单台服务器无法处理的任务，尤其是高并发或者大数据量的任务。
除了对可扩展性的需求，分布式系统<strong>还有不出现单点故障、服务或者存储无状态等特点</strong>。</p>

<ul>
<li>单点故障（Single Point Failure）是指在系统中某个组件一旦失效，这会让整个系统无法工作，而不出现单点故障，单点不影响整体，就是分布式系统的设计目标之一；</li>
<li>无状态，是因为无状态的服务才能满足部分机器宕机不影响全部，可以随时进行扩展的需求。
由于分布式系统的特点，在分布式环境中更容易出现问题，比如节点之间通信失败、网络分区故障、多个副本的数据不一致等，为了更好地在分布式系统下进行开发，学者们提出了一系列的理论，其中具有代表性的就是 CAP 理论。</li>
</ul>

<h2 id="cap-代表什么含义">CAP 代表什么含义</h2>

<p>CAP 理论可以表述为，一个分布式系统最多只能同时满足一致性（Consistency）、可用性（Availability）和分区容忍性（Partition Tolerance）这三项中的两项。
<img src="assets/Ciqah16ER_SAGmCqAADG3jNX34o901.png" alt="img" /></p>

<p><strong>一致性</strong>是指“所有节点同时看到相同的数据”，即更新操作成功并返回客户端完成后，所有节点在同一时间的数据完全一致，等同于所有节点拥有数据的最新版本。
<strong>可用性</strong>是指“任何时候，读写都是成功的”，即服务一直可用，而且是正常响应时间。我们平时会看到一些 IT 公司的对外宣传，比如系统稳定性已经做到 3 个 9、4 个 9，即 99.9%、99.99%，这里的 N 个 9 就是对可用性的一个描述，叫做 SLA，即服务水平协议。比如我们说月度 99.95% 的 SLA，则意味着每个月服务出现故障的时间只能占总时间的 0.05%，如果这个月是 30 天，那么就是 21.6 分钟。
<strong>分区容忍性</strong>具体是指“当部分节点出现消息丢失或者分区故障的时候，分布式系统仍然能够继续运行”，即系统容忍网络出现分区，并且在遇到某节点或网络分区之间网络不可达的情况下，仍然能够对外提供满足一致性和可用性的服务。
在分布式系统中，由于系统的各层拆分，P 是确定的，CAP 的应用模型就是 CP 架构和 AP 架构。分布式系统所关注的，就是在 Partition Tolerance 的前提下，如何实现更好的 A 和更稳定的 C。</p>

<h3 id="cap-理论的证明">CAP 理论的证明</h3>

<p>CAP 理论的证明有多种方式，通过<strong>反证</strong>的方式是最直观的。反证法来证明 CAP 定理，最早是由 Lynch 提出的，通过一个实际场景，如果 CAP 三者可同时满足，由于允许 P 的存在，则一定存在 Server 之间的丢包，如此则不能保证 C。
<img src="assets/Cgq2xl6ER_SAIiA0AACyIE8xkbY529.png" alt="img" />
首先构造一个单机系统，如上图，Client A 可以发送指令到 Server 并且设置更新 X 的值，Client 1 从 Server 读取该值，在单点情况下，即没有网络分区的情况下，通过简单的事务机制，可以保证 Client 1 读到的始终是最新值，不存在一致性的问题。
<img src="assets/Ciqah16ER_SAbt2BAAGrjnzOmj0352.png" alt="img" />
我们在系统中增加一组节点，因为允许分区容错，Write 操作可能在 Server 1 上成功，在 Server 2 上失败，这时候对于 Client 1 和 Client 2，就会读取到不一致的值，出现不一致的情况。如果要保持 X 值的一致性，Write 操作必须同时失败， 也就是降低系统的可用性。
可以看到，在分布式系统中，无法同时满足 CAP 定律中的“一致性”“可用性”和“分区容错性”三者。
在该证明中，对 CAP 的定义进行了更明确的声明：</p>

<ul>
<li>Consistency，一致性被称为原子对象，任何的读写都应该看起来是“原子”的，或串行的，写后面的读一定能读到前面写的内容，所有的读写请求都好像被全局排序；</li>
<li>Availability，对任何非失败节点都应该在有限时间内给出请求的回应（请求的可终止性）；</li>
<li>Partition Tolerance，允许节点之间丢失任意多的消息，当网络分区发生时，节点之间的消息可能会完全丢失。</li>
</ul>

<h3 id="cap-理论的应用">CAP 理论的应用</h3>

<p>CAP 理论提醒我们，在架构设计中，不要把精力浪费在如何设计能满足三者的完美分布式系统上，而要合理进行取舍，CAP 理论类似数学上的不可能三角，只能三者选其二，不能全部获得。
不同业务对于一致性的要求是不同的。举个例来讲，在微博上发表评论和点赞，用户对不一致是不敏感的，可以容忍相对较长时间的不一致，只要做好本地的交互，并不会影响用户体验；而我们在电商购物时，产品价格数据则是要求强一致性的，如果商家更改价格不能实时生效，则会对交易成功率有非常大的影响。
需要注意的是，CAP 理论中是忽略网络延迟的，也就是当事务提交时，节点间的数据复制一定是需要花费时间的。即使是同一个机房，从节点 A 复制到节点 B，由于现实中网络不是实时的，所以总会有一定的时间不一致。</p>

<h3 id="cp-和-ap-架构的取舍">CP 和 AP 架构的取舍</h3>

<p>在通常的分布式系统中，为了保证数据的高可用，通常会将数据保留多个<strong>副本</strong>（Replica），网络分区是既成的现实，于是只能在可用性和一致性两者间做出选择。CAP 理论关注的是在绝对情况下，在工程上，可用性和一致性并不是完全对立的，我们关注的往往是如何在保持相对一致性的前提下，提高系统的可用性。
业务上对一致性的要求会直接反映在系统设计中，典型的就是 CP 和 AP 结构。</p>

<ul>
<li>CP 架构：对于 CP 来说，放弃可用性，追求一致性和分区容错性。</li>
</ul>

<p>我们熟悉的 ZooKeeper，就是采用了 CP 一致性，ZooKeeper 是一个分布式的服务框架，主要用来解决分布式集群中应用系统的协调和一致性问题。其核心算法是 Zab，所有设计都是为了一致性。在 CAP 模型中，ZooKeeper 是 CP，这意味着面对网络分区时，为了保持一致性，它是不可用的。关于 Zab 协议，将会在后面的 ZooKeeper 课时中介绍。</p>

<ul>
<li>AP 架构：对于 AP 来说，放弃强一致性，追求分区容错性和可用性，这是很多分布式系统设计时的选择，后面的 Base 也是根据 AP 来扩展的。</li>
</ul>

<p>和 ZooKeeper 相对的是 Eureka，Eureka 是 Spring Cloud 微服务技术栈中的服务发现组件，Eureka 的各个节点都是平等的，几个节点挂掉不影响正常节点的工作，剩余的节点依然可以提供注册和查询服务，只要有一台 Eureka 还在，就能保证注册服务可用，只不过查到的信息可能不是最新的版本，不保证一致性。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#345858580d0005050403745359555d581a575b59" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f13c442dc739508',t:'MTczNDA2OTk3MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=13&#32;&#32;微内核架构：ShardingSphere&#32;如何实现系统的扩展性？>
        <link rel="icon" href="/static/favicon.png">
        <title>13  微内核架构：ShardingSphere 如何实现系统的扩展性？ </title>
        
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
                        <a class="menu-item" id="00 如何正确学习一款分库分表开源框架？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/00%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e5%ad%a6%e4%b9%a0%e4%b8%80%e6%ac%be%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8%e5%bc%80%e6%ba%90%e6%a1%86%e6%9e%b6%ef%bc%9f.md">00 如何正确学习一款分库分表开源框架？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  从理论到实践：如何让分库分表真正落地？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/01%20%20%e4%bb%8e%e7%90%86%e8%ae%ba%e5%88%b0%e5%ae%9e%e8%b7%b5%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8%e7%9c%9f%e6%ad%a3%e8%90%bd%e5%9c%b0%ef%bc%9f.md">01  从理论到实践：如何让分库分表真正落地？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  顶级项目：ShardingSphere 是一款什么样的 Apache 开源软件？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/02%20%20%e9%a1%b6%e7%ba%a7%e9%a1%b9%e7%9b%ae%ef%bc%9aShardingSphere%20%e6%98%af%e4%b8%80%e6%ac%be%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%20Apache%20%e5%bc%80%e6%ba%90%e8%bd%af%e4%bb%b6%ef%bc%9f.md">02  顶级项目：ShardingSphere 是一款什么样的 Apache 开源软件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  规范兼容：JDBC 规范与 ShardingSphere 是什么关系？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/03%20%20%e8%a7%84%e8%8c%83%e5%85%bc%e5%ae%b9%ef%bc%9aJDBC%20%e8%a7%84%e8%8c%83%e4%b8%8e%20ShardingSphere%20%e6%98%af%e4%bb%80%e4%b9%88%e5%85%b3%e7%b3%bb%ef%bc%9f.md">03  规范兼容：JDBC 规范与 ShardingSphere 是什么关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  应用集成：在业务系统中使用 ShardingSphere 的方式有哪些？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/04%20%20%e5%ba%94%e7%94%a8%e9%9b%86%e6%88%90%ef%bc%9a%e5%9c%a8%e4%b8%9a%e5%8a%a1%e7%b3%bb%e7%bb%9f%e4%b8%ad%e4%bd%bf%e7%94%a8%20ShardingSphere%20%e7%9a%84%e6%96%b9%e5%bc%8f%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">04  应用集成：在业务系统中使用 ShardingSphere 的方式有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  配置驱动：ShardingSphere 中的配置体系是如何设计的？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/05%20%20%e9%85%8d%e7%bd%ae%e9%a9%b1%e5%8a%a8%ef%bc%9aShardingSphere%20%e4%b8%ad%e7%9a%84%e9%85%8d%e7%bd%ae%e4%bd%93%e7%b3%bb%e6%98%af%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e7%9a%84%ef%bc%9f.md">05  配置驱动：ShardingSphere 中的配置体系是如何设计的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  数据分片：如何实现分库、分表、分库&#43;分表以及强制路由？（上）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/06%20%20%e6%95%b0%e6%8d%ae%e5%88%86%e7%89%87%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%88%86%e5%ba%93%e3%80%81%e5%88%86%e8%a1%a8%e3%80%81%e5%88%86%e5%ba%93&#43;%e5%88%86%e8%a1%a8%e4%bb%a5%e5%8f%8a%e5%bc%ba%e5%88%b6%e8%b7%af%e7%94%b1%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">06  数据分片：如何实现分库、分表、分库&#43;分表以及强制路由？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  数据分片：如何实现分库、分表、分库&#43;分表以及强制路由？（下）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/07%20%20%e6%95%b0%e6%8d%ae%e5%88%86%e7%89%87%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%88%86%e5%ba%93%e3%80%81%e5%88%86%e8%a1%a8%e3%80%81%e5%88%86%e5%ba%93&#43;%e5%88%86%e8%a1%a8%e4%bb%a5%e5%8f%8a%e5%bc%ba%e5%88%b6%e8%b7%af%e7%94%b1%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">07  数据分片：如何实现分库、分表、分库&#43;分表以及强制路由？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  读写分离：如何集成分库分表&#43;数据库主从架构？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/08%20%20%e8%af%bb%e5%86%99%e5%88%86%e7%a6%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e9%9b%86%e6%88%90%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8&#43;%e6%95%b0%e6%8d%ae%e5%ba%93%e4%b8%bb%e4%bb%8e%e6%9e%b6%e6%9e%84%ef%bc%9f.md">08  读写分离：如何集成分库分表&#43;数据库主从架构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  分布式事务：如何使用强一致性事务与柔性事务？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/09%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%bc%ba%e4%b8%80%e8%87%b4%e6%80%a7%e4%ba%8b%e5%8a%a1%e4%b8%8e%e6%9f%94%e6%80%a7%e4%ba%8b%e5%8a%a1%ef%bc%9f.md">09  分布式事务：如何使用强一致性事务与柔性事务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  数据脱敏：如何确保敏感数据的安全访问？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/10%20%20%e6%95%b0%e6%8d%ae%e8%84%b1%e6%95%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e7%a1%ae%e4%bf%9d%e6%95%8f%e6%84%9f%e6%95%b0%e6%8d%ae%e7%9a%84%e5%ae%89%e5%85%a8%e8%ae%bf%e9%97%ae%ef%bc%9f.md">10  数据脱敏：如何确保敏感数据的安全访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  编排治理：如何实现分布式环境下的动态配置管理？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/11%20%20%e7%bc%96%e6%8e%92%e6%b2%bb%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e7%8e%af%e5%a2%83%e4%b8%8b%e7%9a%84%e5%8a%a8%e6%80%81%e9%85%8d%e7%bd%ae%e7%ae%a1%e7%90%86%ef%bc%9f.md">11  编排治理：如何实现分布式环境下的动态配置管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  从应用到原理：如何高效阅读 ShardingSphere 源码？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/12%20%20%e4%bb%8e%e5%ba%94%e7%94%a8%e5%88%b0%e5%8e%9f%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e9%98%85%e8%af%bb%20ShardingSphere%20%e6%ba%90%e7%a0%81%ef%bc%9f.md">12  从应用到原理：如何高效阅读 ShardingSphere 源码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  微内核架构：ShardingSphere 如何实现系统的扩展性？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/13%20%20%e5%be%ae%e5%86%85%e6%a0%b8%e6%9e%b6%e6%9e%84%ef%bc%9aShardingSphere%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%89%a9%e5%b1%95%e6%80%a7%ef%bc%9f.md">13  微内核架构：ShardingSphere 如何实现系统的扩展性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  分布式主键：ShardingSphere 中有哪些分布式主键实现方式？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/14%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%bb%e9%94%ae%ef%bc%9aShardingSphere%20%e4%b8%ad%e6%9c%89%e5%93%aa%e4%ba%9b%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%bb%e9%94%ae%e5%ae%9e%e7%8e%b0%e6%96%b9%e5%bc%8f%ef%bc%9f.md">14  分布式主键：ShardingSphere 中有哪些分布式主键实现方式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  解析引擎：SQL 解析流程应该包括哪些核心阶段？（上）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/15%20%20%e8%a7%a3%e6%9e%90%e5%bc%95%e6%93%8e%ef%bc%9aSQL%20%e8%a7%a3%e6%9e%90%e6%b5%81%e7%a8%8b%e5%ba%94%e8%af%a5%e5%8c%85%e6%8b%ac%e5%93%aa%e4%ba%9b%e6%a0%b8%e5%bf%83%e9%98%b6%e6%ae%b5%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">15  解析引擎：SQL 解析流程应该包括哪些核心阶段？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  解析引擎：SQL 解析流程应该包括哪些核心阶段？（下）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/16%20%20%e8%a7%a3%e6%9e%90%e5%bc%95%e6%93%8e%ef%bc%9aSQL%20%e8%a7%a3%e6%9e%90%e6%b5%81%e7%a8%8b%e5%ba%94%e8%af%a5%e5%8c%85%e6%8b%ac%e5%93%aa%e4%ba%9b%e6%a0%b8%e5%bf%83%e9%98%b6%e6%ae%b5%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">16  解析引擎：SQL 解析流程应该包括哪些核心阶段？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  路由引擎：如何理解分片路由核心类 ShardingRouter 的运作机制？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/17%20%20%e8%b7%af%e7%94%b1%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e5%88%86%e7%89%87%e8%b7%af%e7%94%b1%e6%a0%b8%e5%bf%83%e7%b1%bb%20ShardingRouter%20%e7%9a%84%e8%bf%90%e4%bd%9c%e6%9c%ba%e5%88%b6%ef%bc%9f.md">17  路由引擎：如何理解分片路由核心类 ShardingRouter 的运作机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  路由引擎：如何实现数据访问的分片路由和广播路由？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/18%20%20%e8%b7%af%e7%94%b1%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%e7%9a%84%e5%88%86%e7%89%87%e8%b7%af%e7%94%b1%e5%92%8c%e5%b9%bf%e6%92%ad%e8%b7%af%e7%94%b1%ef%bc%9f.md">18  路由引擎：如何实现数据访问的分片路由和广播路由？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  路由引擎：如何在路由过程中集成多种路由策略和路由算法？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/19%20%20%e8%b7%af%e7%94%b1%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e8%b7%af%e7%94%b1%e8%bf%87%e7%a8%8b%e4%b8%ad%e9%9b%86%e6%88%90%e5%a4%9a%e7%a7%8d%e8%b7%af%e7%94%b1%e7%ad%96%e7%95%a5%e5%92%8c%e8%b7%af%e7%94%b1%e7%ae%97%e6%b3%95%ef%bc%9f.md">19  路由引擎：如何在路由过程中集成多种路由策略和路由算法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  改写引擎：如何理解装饰器模式下的 SQL 改写实现机制？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/20%20%20%e6%94%b9%e5%86%99%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e8%a3%85%e9%a5%b0%e5%99%a8%e6%a8%a1%e5%bc%8f%e4%b8%8b%e7%9a%84%20SQL%20%e6%94%b9%e5%86%99%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6%ef%bc%9f.md">20  改写引擎：如何理解装饰器模式下的 SQL 改写实现机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  执行引擎：分片环境下 SQL 执行的整体流程应该如何进行抽象？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/21%20%20%e6%89%a7%e8%a1%8c%e5%bc%95%e6%93%8e%ef%bc%9a%e5%88%86%e7%89%87%e7%8e%af%e5%a2%83%e4%b8%8b%20SQL%20%e6%89%a7%e8%a1%8c%e7%9a%84%e6%95%b4%e4%bd%93%e6%b5%81%e7%a8%8b%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e6%8a%bd%e8%b1%a1%ef%bc%9f.md">21  执行引擎：分片环境下 SQL 执行的整体流程应该如何进行抽象？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  执行引擎：如何把握 ShardingSphere 中的 Executor 执行模型？（上）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/22%20%20%e6%89%a7%e8%a1%8c%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8a%8a%e6%8f%a1%20ShardingSphere%20%e4%b8%ad%e7%9a%84%20Executor%20%e6%89%a7%e8%a1%8c%e6%a8%a1%e5%9e%8b%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">22  执行引擎：如何把握 ShardingSphere 中的 Executor 执行模型？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  执行引擎：如何把握 ShardingSphere 中的 Executor 执行模型？（下）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/23%20%20%e6%89%a7%e8%a1%8c%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8a%8a%e6%8f%a1%20ShardingSphere%20%e4%b8%ad%e7%9a%84%20Executor%20%e6%89%a7%e8%a1%8c%e6%a8%a1%e5%9e%8b%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">23  执行引擎：如何把握 ShardingSphere 中的 Executor 执行模型？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  归并引擎：如何理解数据归并的类型以及简单归并策略的实现过程？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/24%20%20%e5%bd%92%e5%b9%b6%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e6%95%b0%e6%8d%ae%e5%bd%92%e5%b9%b6%e7%9a%84%e7%b1%bb%e5%9e%8b%e4%bb%a5%e5%8f%8a%e7%ae%80%e5%8d%95%e5%bd%92%e5%b9%b6%e7%ad%96%e7%95%a5%e7%9a%84%e5%ae%9e%e7%8e%b0%e8%bf%87%e7%a8%8b%ef%bc%9f.md">24  归并引擎：如何理解数据归并的类型以及简单归并策略的实现过程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  归并引擎：如何理解流式归并和内存归并在复杂归并场景下的应用方式？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/25%20%20%e5%bd%92%e5%b9%b6%e5%bc%95%e6%93%8e%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e6%b5%81%e5%bc%8f%e5%bd%92%e5%b9%b6%e5%92%8c%e5%86%85%e5%ad%98%e5%bd%92%e5%b9%b6%e5%9c%a8%e5%a4%8d%e6%9d%82%e5%bd%92%e5%b9%b6%e5%9c%ba%e6%99%af%e4%b8%8b%e7%9a%84%e5%ba%94%e7%94%a8%e6%96%b9%e5%bc%8f%ef%bc%9f.md">25  归并引擎：如何理解流式归并和内存归并在复杂归并场景下的应用方式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  读写分离：普通主从架构和分片主从架构分别是如何实现的？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/26%20%20%e8%af%bb%e5%86%99%e5%88%86%e7%a6%bb%ef%bc%9a%e6%99%ae%e9%80%9a%e4%b8%bb%e4%bb%8e%e6%9e%b6%e6%9e%84%e5%92%8c%e5%88%86%e7%89%87%e4%b8%bb%e4%bb%8e%e6%9e%b6%e6%9e%84%e5%88%86%e5%88%ab%e6%98%af%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%9a%84%ef%bc%9f.md">26  读写分离：普通主从架构和分片主从架构分别是如何实现的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  分布式事务：如何理解 ShardingSphere 中对分布式事务的抽象过程？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/27%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%20ShardingSphere%20%e4%b8%ad%e5%af%b9%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%e7%9a%84%e6%8a%bd%e8%b1%a1%e8%bf%87%e7%a8%8b%ef%bc%9f.md">27  分布式事务：如何理解 ShardingSphere 中对分布式事务的抽象过程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  分布式事务：ShardingSphere 中如何集成强一致性事务和柔性事务支持？（上）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/28%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9aShardingSphere%20%e4%b8%ad%e5%a6%82%e4%bd%95%e9%9b%86%e6%88%90%e5%bc%ba%e4%b8%80%e8%87%b4%e6%80%a7%e4%ba%8b%e5%8a%a1%e5%92%8c%e6%9f%94%e6%80%a7%e4%ba%8b%e5%8a%a1%e6%94%af%e6%8c%81%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">28  分布式事务：ShardingSphere 中如何集成强一致性事务和柔性事务支持？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  分布式事务：ShardingSphere 中如何集成强一致性事务和柔性事务支持？（下）.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/29%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9aShardingSphere%20%e4%b8%ad%e5%a6%82%e4%bd%95%e9%9b%86%e6%88%90%e5%bc%ba%e4%b8%80%e8%87%b4%e6%80%a7%e4%ba%8b%e5%8a%a1%e5%92%8c%e6%9f%94%e6%80%a7%e4%ba%8b%e5%8a%a1%e6%94%af%e6%8c%81%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">29  分布式事务：ShardingSphere 中如何集成强一致性事务和柔性事务支持？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  数据脱敏：如何基于改写引擎实现低侵入性数据脱敏方案？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/30%20%20%e6%95%b0%e6%8d%ae%e8%84%b1%e6%95%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%e6%94%b9%e5%86%99%e5%bc%95%e6%93%8e%e5%ae%9e%e7%8e%b0%e4%bd%8e%e4%be%b5%e5%85%a5%e6%80%a7%e6%95%b0%e6%8d%ae%e8%84%b1%e6%95%8f%e6%96%b9%e6%a1%88%ef%bc%9f.md">30  数据脱敏：如何基于改写引擎实现低侵入性数据脱敏方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 配置中心：如何基于配置中心实现配置信息的动态化管理？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/31%20%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83%e5%ae%9e%e7%8e%b0%e9%85%8d%e7%bd%ae%e4%bf%a1%e6%81%af%e7%9a%84%e5%8a%a8%e6%80%81%e5%8c%96%e7%ae%a1%e7%90%86%ef%bc%9f.md">31 配置中心：如何基于配置中心实现配置信息的动态化管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 注册中心：如何基于注册中心实现数据库访问熔断机制？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/32%20%e6%b3%a8%e5%86%8c%e4%b8%ad%e5%bf%83%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%e6%b3%a8%e5%86%8c%e4%b8%ad%e5%bf%83%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e5%ba%93%e8%ae%bf%e9%97%ae%e7%86%94%e6%96%ad%e6%9c%ba%e5%88%b6%ef%bc%9f.md">32 注册中心：如何基于注册中心实现数据库访问熔断机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 链路跟踪：如何基于 Hook 机制以及 OpenTracing 协议实现数据访问链路跟踪？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/33%20%e9%93%be%e8%b7%af%e8%b7%9f%e8%b8%aa%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20Hook%20%e6%9c%ba%e5%88%b6%e4%bb%a5%e5%8f%8a%20OpenTracing%20%e5%8d%8f%e8%ae%ae%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%e9%93%be%e8%b7%af%e8%b7%9f%e8%b8%aa%ef%bc%9f.md">33 链路跟踪：如何基于 Hook 机制以及 OpenTracing 协议实现数据访问链路跟踪？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 系统集成：如何完成 ShardingSphere 内核与 Spring&#43;SpringBoot 的无缝整合？.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/34%20%e7%b3%bb%e7%bb%9f%e9%9b%86%e6%88%90%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%8c%e6%88%90%20ShardingSphere%20%e5%86%85%e6%a0%b8%e4%b8%8e%20Spring&#43;SpringBoot%20%e7%9a%84%e6%97%a0%e7%bc%9d%e6%95%b4%e5%90%88%ef%bc%9f.md">34 系统集成：如何完成 ShardingSphere 内核与 Spring&#43;SpringBoot 的无缝整合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 结语：ShardingSphere 总结及展望.md" href="/%e4%b8%93%e6%a0%8f/ShardingSphere%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e7%b2%be%e8%ae%b2-%e5%ae%8c/35%20%e7%bb%93%e8%af%ad%ef%bc%9aShardingSphere%20%e6%80%bb%e7%bb%93%e5%8f%8a%e5%b1%95%e6%9c%9b.md">35 结语：ShardingSphere 总结及展望.md</a>
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
                            <h1 id="title" data-id="13  微内核架构：ShardingSphere 如何实现系统的扩展性？" class="title">13  微内核架构：ShardingSphere 如何实现系统的扩展性？</h1>
                            <div><p>我们已经在课程中多次提到 ShardingSphere 使用了微内核架构来实现框架的扩展性。随着课程的演进，我们会发现，用于实现配置中心的 ConfigCenter、用于数据脱敏的 ShardingEncryptor 以及用于数据库治理的注册中心接口 RegistryCenter 等大量组件的实现也都使用了微内核架构。那么，究竟什么是微内核架构呢？今天我们就来讨论这个架构模式的基本原理以及在 ShardingSphere 中的应用。</p>

<h3 id="什么是微内核架构">什么是微内核架构？</h3>

<p><strong>微内核是一种典型的架构模式</strong> ，区别于普通的设计模式，架构模式是一种高层模式，用于描述系统级的结构组成、相互关系及相关约束。微内核架构在开源框架中的应用也比较广泛，除了 ShardingSphere 之外，在主流的 PRC 框架 Dubbo 中也实现了自己的微内核架构。那么，在介绍什么是微内核架构之前，我们有必要先阐述这些开源框架会使用微内核架构的原因。</p>

<h4 id="为什么要使用微内核架构">为什么要使用微内核架构？</h4>

<p><strong>微内核架构本质上是为了提高系统的扩展性</strong> 。所谓扩展性，是指系统在经历不可避免的变更时所具有的灵活性，以及针对提供这样的灵活性所需要付出的成本间的平衡能力。也就是说，当在往系统中添加新业务时，不需要改变原有的各个组件，只需把新业务封闭在一个新的组件中就能完成整体业务的升级，我们认为这样的系统具有较好的可扩展性。</p>

<p>就架构设计而言，扩展性是软件设计的永恒话题。而要实现系统扩展性，一种思路是提供可插拔式的机制来应对所发生的变化。当系统中现有的某个组件不满足要求时，我们可以实现一个新的组件来替换它，而整个过程对于系统的运行而言应该是无感知的，我们也可以根据需要随时完成这种新旧组件的替换。</p>

<p>比如在下个课时中我们将要介绍的 ShardingSphere 中提供的分布式主键功能，分布式主键的实现可能有很多种，而扩展性在这个点上的体现就是， <strong>我们可以使用任意一种新的分布式主键实现来替换原有的实现，而不需要依赖分布式主键的业务代码做任何的改变</strong> 。</p>

<p><img src="assets/CgqCHl8esVaAVlUFAACJmGjQZDA482.png" alt="image.png" /></p>

<p>微内核架构模式为这种实现扩展性的思路提供了架构设计上的支持，ShardingSphere 基于微内核架构实现了高度的扩展性。在介绍如何实现微内核架构之前，我们先对微内核架构的具体组成结构和基本原理做简要的阐述。</p>

<h4 id="什么是微内核架构-1">什么是微内核架构？</h4>

<p>从组成结构上讲， <strong>微内核架构包含两部分组件：内核系统和插件</strong> 。这里的内核系统通常提供系统运行所需的最小功能集，而插件是独立的组件，包含自定义的各种业务代码，用来向内核系统增强或扩展额外的业务能力。在 ShardingSphere 中，前面提到的分布式主键就是插件，而 ShardingSphere 的运行时环境构成了内核系统。</p>

<p><img src="assets/CgqCHl8esWOAJ-5cAACfxz06p_E616.png" alt="image" /></p>

<p>那么这里的插件具体指的是什么呢？这就需要我们明确两个概念，一个概念就是经常在说的 <strong>API</strong> ，这是系统对外暴露的接口。而另一个概念就是 <strong>SPI</strong>（Service Provider Interface，服务提供接口），这是插件自身所具备的扩展点。就两者的关系而言，API 面向业务开发人员，而 SPI 面向框架开发人员，两者共同构成了 ShardingSphere 本身。</p>

<p><img src="assets/Ciqc1F8esXOADonEAACE9HEUTJc298.png" alt="image" /></p>

<p>可插拔式的实现机制说起来简单，做起来却不容易，我们需要考虑两方面内容。一方面，我们需要梳理系统的变化并把它们抽象成多个 SPI 扩展点。另一方面， <strong>当我们实现了这些 SPI 扩展点之后，就需要构建一个能够支持这种可插拔机制的具体实现，从而提供一种 SPI 运行时环境</strong> 。</p>

<p>那么，ShardingSphere 是如何实现微内核架构的呢？让我们来一起看一下。</p>

<h3 id="如何实现微内核架构">如何实现微内核架构？</h3>

<p>事实上，JDK 已经为我们提供了一种微内核架构的实现方式，这种实现方式针对如何设计和实现 SPI 提出了一些开发和配置上的规范，ShardingSphere 使用的就是这种规范。首先，我们需要设计一个服务接口，并根据需要提供不同的实现类。接下来，我们将模拟实现分布式主键的应用场景。</p>

<p>基于 SPI 的约定，创建一个单独的工程来存放服务接口，并给出接口定义。请注意 <strong>这个服务接口的完整类路径为 com.tianyilan.KeyGenerator</strong> ，接口中只包含一个获取目标主键的简单示例方法。</p>

<pre><code>package com.tianyilan; 

public interface KeyGenerator{ 

    String getKey(); 
}
</code></pre>

<p>针对该接口，提供两个简单的实现类，分别是基于 UUID 的 UUIDKeyGenerator 和基于雪花算法的 SnowflakeKeyGenerator。为了让演示过程更简单，这里我们直接返回一个模拟的结果，真实的实现过程我们会在下一课时中详细介绍。</p>

<pre><code>public class UUIDKeyGenerator implements KeyGenerator { 

    @Override 
    public String getKey() { 

       return &quot;UUIDKey&quot;; 
    } 
} 
	
public class SnowflakeKeyGenerator implements KeyGenerator { 

    @Override 
    public String getKey() { 

       return &quot;SnowflakeKey&quot;; 
    } 
}
</code></pre>

<p>接下来的这个步骤很关键， <strong>在这个代码工程的 META-INF/services/ 目录下，需要创建一个以服务接口完整类路径 com.tianyilan.KeyGenerator 命名的文件</strong> ，文件的内容是指向该接口所对应的两个实现类的完整类路径 com.tianyilan.UUIDKeyGenerator 和 com.tianyilan. SnowflakeKeyGenerator。</p>

<p>我们把这个代码工程打成一个 jar 包，然后新建另一个代码工程，该代码工程需要这个 jar 包，并完成如下所示的 Main 函数。</p>

<pre><code>import java.util.ServiceLoader; 
import com.tianyilan. KeyGenerator; 

public class Main { 
    public static void main(String[] args) { 

       ServiceLoader&lt;KeyGenerator&gt; generators = ServiceLoader.load(KeyGenerator.class); 

       for (KeyGenerator generator : generators) { 
           System.out.println(generator.getClass()); 
           String key = generator.getKey(); 
           System.out.println(key); 
       } 
    } 
}
</code></pre>

<p>现在，该工程的角色是 SPI 服务的使用者，这里使用了 JDK 提供的 ServiceLoader 工具类来获取所有 KeyGenerator 的实现类。现在在 jar 包的 META-INF/services/com.tianyilan.KeyGenerator 文件中有两个 KeyGenerator 实现类的定义。执行这段 Main 函数，我们将得到的输出结果如下：</p>

<pre><code>	class com.tianyilan.UUIDKeyGenerator 
	UUIDKey 
	class com.tianyilan.SnowflakeKeyGenerator 
	SnowflakeKey
</code></pre>

<p>如果我们调整 META-INF/services/com.tianyilan.KeyGenerator 文件中的内容，去掉 com.tianyilan.UUIDKeyGenerator 的定义，并重新打成 jar 包供 SPI 服务的使用者进行引用。再次执行 Main 函数，则只会得到基于 SnowflakeKeyGenerator 的输出结果。</p>

<p>至此， 完整 的 SPI 提供者和使用者的实现过程演示完毕。我们通过一张图，总结基于 JDK 的 SPI 机制实现微内核架构的开发流程：</p>

<p><img src="assets/Ciqc1F8esYqAdXABAADVVh6mYnA926.png" alt="image" /></p>

<p>这个示例非常简单，但却是 ShardingSphere 中实现微内核架构的基础。接下来，就让我们把话题转到 ShardingSphere，看看 ShardingSphere 中应用 SPI 机制的具体方法。</p>

<h3 id="shardingsphere-如何基于微内核架构实现扩展性">ShardingSphere 如何基于微内核架构实现扩展性？</h3>

<p>ShardingSphere 中微内核架构的实现过程并不复杂，基本就是对 JDK 中 SPI 机制的封装。让我们一起来看一下。</p>

<h4 id="shardingsphere-中的微内核架构基础实现机制">ShardingSphere 中的微内核架构基础实现机制</h4>

<p>我们发现，在 ShardingSphere 源码的根目录下，存在一个独立的工程 shardingsphere-spi。显然，从命名上看，这个工程中应该包含了 ShardingSphere 实现 SPI 的相关代码。我们快速浏览该工程，发现里面只有一个接口定义和两个工具类。我们先来看这个接口定义 TypeBasedSPI：</p>

<pre><code>public interface TypeBasedSPI { 

    //获取SPI对应的类型 
    String getType(); 

    //获取属性 
    Properties getProperties(); 

    //设置属性 
    void setProperties(Properties properties); 
}
</code></pre>

<p>从定位上看，这个接口在 ShardingSphere 中应该是一个顶层接口，我们已经在上一课时给出了这一接口的实现类类层结构。接下来再看一下 NewInstanceServiceLoader 类，从命名上看，不难想象该类的作用类似于一种 ServiceLoader，用于加载新的目标对象实例：</p>

<pre><code>public final class NewInstanceServiceLoader { 

    private static final Map&lt;Class, Collection&lt;Class&lt;?&gt;&gt;&gt; SERVICE_MAP = new HashMap&lt;&gt;(); 

    //通过ServiceLoader获取新的SPI服务实例并注册到SERVICE_MAP中
    public static &lt;T&gt; void register(final Class&lt;T&gt; service) { 
        for (T each : ServiceLoader.load(service)) { 
            registerServiceClass(service, each); 
        } 
    } 

    @SuppressWarnings(&quot;unchecked&quot;) 
    private static &lt;T&gt; void registerServiceClass(final Class&lt;T&gt; service, final T instance) { 
        Collection&lt;Class&lt;?&gt;&gt; serviceClasses = SERVICE_MAP.get(service); 
        if (null == serviceClasses) { 
            serviceClasses = new LinkedHashSet&lt;&gt;(); 
        } 
        serviceClasses.add(instance.getClass()); 
        SERVICE_MAP.put(service, serviceClasses); 
    } 

    @SneakyThrows 
    @SuppressWarnings(&quot;unchecked&quot;) 
    public static &lt;T&gt; Collection&lt;T&gt; newServiceInstances(final Class&lt;T&gt; service) { 
        Collection&lt;T&gt; result = new LinkedList&lt;&gt;(); 
        if (null == SERVICE_MAP.get(service)) { 
            return result; 
        } 
        for (Class&lt;?&gt; each : SERVICE_MAP.get(service)) { 
            result.add((T) each.newInstance()); 
        } 
        return result; 
    } 
}
</code></pre>

<p>在上面这段代码中， 首先看到了熟悉的 ServiceLoader.load(service) 方法，这是 JDK 中 ServiceLoader 工具类的具体应用。同时，注意到 ShardingSphere 使用了一个 HashMap 来保存类的定义以及类的实例之 间 的一对多关系，可以认为，这是一种用于提高访问效率的缓存机制。</p>

<p>最后，我们来看一下 TypeBasedSPIServiceLoader 的实现，该类依赖于前面介绍的 NewInstanceServiceLoader 类。 下面这段代码演示了 基于 NewInstanceServiceLoader 获取实例类列表，并根据所传入的类型做过滤：</p>

<pre><code>    //使用NewInstanceServiceLoader获取实例类列表，并根据类型做过滤 
    private Collection&lt;T&gt; loadTypeBasedServices(final String type) { 
        return Collections2.filter(NewInstanceServiceLoader.newServiceInstances(classType), new Predicate&lt;T&gt;() { 

            @Override 
            public boolean apply(final T input) { 
                return type.equalsIgnoreCase(input.getType()); 
            } 
        }); 
    }
</code></pre>

<p>TypeBasedSPIServiceLoader 对外暴露了服务的接口，对通过 loadTypeBasedServices 方法获取的服务实例设置对应的属性然后返回：</p>

<pre><code>	//基于类型通过SPI创建实例 
    public final T newService(final String type, final Properties props) { 
        Collection&lt;T&gt; typeBasedServices = loadTypeBasedServices(type); 
        if (typeBasedServices.isEmpty()) { 
            throw new RuntimeException(String.format(&quot;Invalid `%s` SPI type `%s`.&quot;, classType.getName(), type)); 
        } 
        T result = typeBasedServices.iterator().next(); 
        result.setProperties(props); 
        return result; 
	}
</code></pre>

<p>同时，TypeBasedSPIServiceLoader 也对外暴露了不需要传入类型的 newService 方法，该方法使用了 loadFirstTypeBasedService 工具方法来获取第一个服务实例：</p>

<pre><code>	//基于默认类型通过SPI创建实例 
    public final T newService() { 
        T result = loadFirstTypeBasedService(); 
        result.setProperties(new Properties()); 
        return result; 
	} 
	
    private T loadFirstTypeBasedService() { 
        Collection&lt;T&gt; instances = NewInstanceServiceLoader.newServiceInstances(classType); 
        if (instances.isEmpty()) { 
            throw new RuntimeException(String.format(&quot;Invalid `%s` SPI, no implementation class load from SPI.&quot;, classType.getName())); 
        } 
        return instances.iterator().next(); 
	}
</code></pre>

<p>这样，shardingsphere-spi 代码工程中的内容就介绍完毕。 <strong>这部分内容相当于是 ShardingSphere 中所提供的插件运行时环境</strong> 。下面我们基于 ShardingSphere 中提供的几个典型应用场景来讨论这个运行时环境的具体使用方法。</p>

<h4 id="微内核架构在-shardingsphere-中的应用">微内核架构在 ShardingSphere 中的应用</h4>

<ul>
<li>SQL 解析器 SQLParser</li>
</ul>

<p>我们将在 15 课时中介绍 SQLParser 类，该类负责将具体某一条 SQL 解析成一个抽象语法树的整个过程。而这个 SQLParser 的生成由 SQLParserFactory 负责：</p>

<pre><code>	public final class SQLParserFactory { 
	
	    public static SQLParser newInstance(final String databaseTypeName, final String sql) { 
	     //通过SPI机制加载所有扩展 
	     for (SQLParserEntry each : NewInstanceServiceLoader.newServiceInstances(SQLParserEntry.class)) { 
	        … 
	    } 
	}
</code></pre>

<p>可以看到，这里并没有使用前面介绍的 TypeBasedSPIServiceLoader 来加载实例，而是直接使用更为底层的 NewInstanceServiceLoader。</p>

<p>这里引入的 SQLParserEntry 接口就位于 shardingsphere-sql-parser-spi 工程的 org.apache.shardingsphere.sql.parser.spi 包中。显然，从包的命名上看，该接口是一个 SPI 接口。在 SQLParserEntry 类层结构接口中包含一批实现类，分别对应各个具体的数据库：</p>

<p><img src="assets/Ciqc1F8ed26ANXCOAAArBJH3uDs890.png" alt="Drawing 4.png" /></p>

<p>SQLParserEntry 实现类图</p>

<p>我们先来看针对 MySQL 的代码工程 shardingsphere-sql-parser-mysql，在 META-INF/services 目录下，我们找到了一个 org.apache.shardingsphere.sql.parser.spi.SQLParserEntry 文件:</p>

<p><img src="assets/CgqCHl8ed3aABqWdAABTnSG89Jg177.png" alt="Drawing 5.png" /></p>

<p>MySQL 代码工程中的 SPI 配置</p>

<p>可以看到这里指向了 org.apache.shardingsphere.sql.parser.MySQLParserEntry 类。再来到 Oracle 的代码工程 shardingsphere-sql-parser-oracle，在 META-INF/services 目录下，同样找到了一个 org.apache.shardingsphere.sql.parser.spi.SQLParserEntry 文件：</p>

<p><img src="assets/Ciqc1F8ed4GABKlZAABTzlYzJvc755.png" alt="Drawing 6.png" /></p>

<p>Oracle 代码工程中的 SPI 配置</p>

<p>显然，这里应该指向 org.apache.shardingsphere.sql.parser.OracleParserEntry 类，通过这种方式，系统在运行时就会根据类路径动态加载 SPI。</p>

<p>可以注意到，在 SQLParserEntry 接口的类层结构中，实际并没有使用到 TypeBasedSPI 接口 ，而是完全采用了 JDK 原生的 SPI 机制。</p>

<ul>
<li>配置中心 ConfigCenter</li>
</ul>

<p>接下来，我们来找一个使用 TypeBasedSPI 的示例，比如代表配置中心的 ConfigCenter：</p>

<pre><code>public interface ConfigCenter extends TypeBasedSPI
</code></pre>

<p>显然，ConfigCenter 接口继承了 TypeBasedSPI 接口，而在 ShardingSphere 中也存在两个 ConfigCenter 接口的实现类，一个是 ApolloConfigCenter，一个是 CuratorZookeeperConfigCenter。</p>

<p>在 sharding-orchestration-core 工程的 org.apache.shardingsphere.orchestration.internal.configcenter 中，我们找到了 ConfigCenterServiceLoader 类，该类扩展了前面提到的 TypeBasedSPIServiceLoader 类：</p>

<pre><code>public final class ConfigCenterServiceLoader extends TypeBasedSPIServiceLoader&lt;ConfigCenter&gt; { 

    static { 
        NewInstanceServiceLoader.register(ConfigCenter.class); 
    } 

    public ConfigCenterServiceLoader() { 
        super(ConfigCenter.class); 
    } 

    //基于SPI加载ConfigCenter 
    public ConfigCenter load(final ConfigCenterConfiguration configCenterConfig) { 
        Preconditions.checkNotNull(configCenterConfig, &quot;Config center configuration cannot be null.&quot;); 
        ConfigCenter result = newService(configCenterConfig.getType(), configCenterConfig.getProperties()); 
        result.init(configCenterConfig); 
        return result; 
    } 
}
</code></pre>

<p>那么它是如何实现的呢？ 首先，ConfigCenterServiceLoader 类通过 NewInstanceServiceLoader.register(ConfigCenter.class) 语句将所有 ConfigCenter 注册到系统中，这一步会通过 JDK 的 ServiceLoader 工具类加载类路径中的所有 ConfigCenter 实例。</p>

<p>我们可以看到在上面的 load 方法中，通过父类 TypeBasedSPIServiceLoader 的 newService 方法，基于类型创建了 SPI 实例。</p>

<p>以 ApolloConfigCenter 为例，我们来看它的使用方法。在 sharding-orchestration-config-apollo 工程的 META-INF/services 目录下，应该存在一个名为 org.apache.shardingsphere.orchestration.config.api.ConfigCenter 的配置文件，指向 ApolloConfigCenter 类：</p>

<p><img src="assets/CgqCHl8eekGAa88DAABIbz4-Q20783.png" alt="Drawing 7.png" /></p>

<p>Apollo 代码工程中的 SPI 配置</p>

<p>其他的 ConfigCenter 实现也是一样，你可以自行查阅 sharding-orchestration-config-zookeeper-curator 等工程中的 SPI 配置文件。</p>

<p>至此，我们全面了解了 ShardingSphere 中的微内核架构，也就可以基于 ShardingSphere 所提供的各种 SPI 扩展点提供满足自身需求的具体实现。</p>

<h3 id="从源码解析到日常开发">从源码解析到日常开发</h3>

<p>在日常开发过程中，我们一般可以直接使用 JDK 的 ServiceLoader 类来实现 SPI 机制。当然，我们也可以采用像 ShardingSphere 的方式对 ServiceLoader 类进行一层简单的封装，并添加属性设置等自定义功能。</p>

<p>同时，我们也应该注意到，ServiceLoader 这种实现方案也有一定缺点：</p>

<ul>
<li>一方面，META/services 这个配置文件的加载地址是写死在代码中，缺乏灵活性。</li>
<li>另一方面，ServiceLoader 内部采用了基于迭代器的加载方法，会把配置文件中的所有 SPI 实现类都加载到内存中，效率不高。</li>
</ul>

<p>所以如果需要提供更高的灵活性和性能，我们也可以基于 ServiceLoader 的实现方法自己开发适合自身需求的 SPI 加载 机制。</p>

<h3 id="总结">总结</h3>

<p>微内核架构是 ShardingSphere 中最核心的基础架构，为这个框架提供了高度的灵活度，以及可插拔的扩展性。微内核架构也是一种同样的架构模式，本课时我们对这个架构模式的特点和组成结构做了介绍，并基于 JDK 中提供的 SPI 机制给出了实现这一架构模式的具体方案。</p>

<p>ShardingSphere 中大量使用了微内核架构来解耦系统内核和各个组件之间的关联关系，我们基于解析引擎和配置中心给出了具体的实现案例。在学习这些案例时 <strong>，重点在于掌握 ShardingSphere 中对 JDK 中 SPI的封装机制。</strong></p>

<p>这里给你留一道思考题：ShardingSphere 中使用微内核架构时对 JDK 中的 SPI 机制做了哪些封装？</p>

<p>本课时的内容就到这里，下一课时，我们将继续探索 ShardingSphere 中的基础设施，并给出分布式主键的设计原理和多种实现方案，记得按时来听课。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d5b9b9b9ece1e4e4e5e295b2b8b4bcb9fbb6bab8" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f125f8f6abfede4',t:'MTczNDA1NTM1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
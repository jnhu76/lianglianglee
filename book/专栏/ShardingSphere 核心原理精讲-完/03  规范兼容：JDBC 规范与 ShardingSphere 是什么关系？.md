<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=03&#32;&#32;规范兼容：JDBC&#32;规范与&#32;ShardingSphere&#32;是什么关系？>
        <link rel="icon" href="/static/favicon.png">
        <title>03  规范兼容：JDBC 规范与 ShardingSphere 是什么关系？ </title>
        
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
                            <h1 id="title" data-id="03  规范兼容：JDBC 规范与 ShardingSphere 是什么关系？" class="title">03  规范兼容：JDBC 规范与 ShardingSphere 是什么关系？</h1>
                            <div><p>我们知道 ShardingSphere 是一种典型的客户端分片解决方案，而客户端分片的实现方式之一就是重写 JDBC 规范。在上一课时中，我们也介绍了，ShardingSphere 在设计上从一开始就完全兼容 JDBC 规范，ShardingSphere 对外暴露的一套分片操作接口与 JDBC 规范中所提供的接口完全一致。</p>

<p>讲到这里，你可能会觉得有点神奇，<strong>ShardingSphere 究竟是通过什么方式，实现了与 JDBC 规范完全兼容的 API 来提供分片功能呢？</strong></p>

<p>这个问题非常重要，值得我们专门花一个课时的内容来进行分析和讲解。可以说，<strong>理解 JDBC 规范以及 ShardingSphere 对 JDBC 规范的重写方式，是正确使用 ShardingSphere 实现数据分片的前提</strong>。今天，我们就深入讨论 JDBC 规范与 ShardingSphere 的这层关系，帮你从底层设计上解开其中的神奇之处。</p>

<h3 id="jdbc-规范简介">JDBC 规范简介</h3>

<p>ShardingSphere 提供了与 JDBC 规范完全兼容的实现过程，在对这一过程进行详细展开之前，先来回顾一下 JDBC 规范。<strong>JDBC（Java Database Connectivity）的设计初衷是提供一套用于各种数据库的统一标准</strong>，而不同的数据库厂家共同遵守这套标准，并提供各自的实现方案供应用程序调用。作为统一标准，JDBC 规范具有完整的架构体系，如下图所示：</p>

<p><img src="assets/CgqCHl7xtaiASay6AAB0vuO1kAA457.png" alt="Drawing 0.png" /></p>

<p>JDBC 架构中的 Driver Manager 负责加载各种不同的驱动程序（Driver），并根据不同的请求，向调用者返回相应的数据库连接（Connection）。而应用程序通过调用 JDBC API 来实现对数据库的操作。<strong>对于开发人员而言，JDBC API 是我们访问数据库的主要途径，也是 ShardingSphere 重写 JDBC 规范并添加分片功能的入口</strong>。如果我们使用 JDBC 开发一个访问数据库的处理流程，常见的代码风格如下所示：</p>

<pre><code>// 创建池化的数据源
PooledDataSource dataSource = new PooledDataSource ();
// 设置MySQL Driver
dataSource.setDriver (&quot;com.mysql.jdbc.Driver&quot;);
// 设置数据库URL、用户名和密码
dataSource.setUrl (&quot;jdbc:mysql://localhost:3306/test&quot;);
dataSource.setUsername (&quot;root&quot;);
dataSource.setPassword (&quot;root&quot;);
// 获取连接
Connection connection = dataSource.getConnection();
// 执行查询
PreparedStatement statement = connection.prepareStatement (&quot;select * from user&quot;);
// 获取查询结果应该处理
ResultSet resultSet = statement.executeQuery();
while (resultSet.next()) {
	…
}
// 关闭资源
statement.close();
resultSet.close();
connection.close();
</code></pre>

<p>这段代码中包含了 JDBC API 中的核心接口，使用这些核心接口是我们基于 JDBC 进行数据访问的基本方式，这里有必要对这些接口的作用和使用方法做一些展开。事实上，随着课程内容的不断演进，你会发现在 ShardingSphere 中，完成日常开发所使用的也就是这些接口。</p>

<h4 id="datasource">DataSource</h4>

<p>DataSource 在 JDBC 规范中代表的是一种数据源，核心作用是获取数据库连接对象 Connection。在 JDBC 规范中，实际可以直接通过 DriverManager 获取 Connection。我们知道获取 Connection 的过程需要建立与数据库之间的连接，而这个过程会产生较大的系统开销。</p>

<p>为了提高性能，通常会建立一个中间层，该中间层将 DriverManager 生成的 Connection 存放到连接池中，然后从池中获取 Connection，可以认为，DataSource 就是这样一个中间层。在日常开发过程中，我们通常都会基于 DataSource 来获取 Connection。而在 ShardingSphere 中，暴露给业务开发人员的同样是一个经过增强的 DataSource 对象。DataSource 接口的定义是这样的：</p>

<pre><code>public interface DataSource  extends CommonDataSource, Wrapper {
  Connection getConnection() throws SQLException; 
  Connection getConnection(String username, String password)
    throws SQLException;
}
</code></pre>

<p>可以看到，DataSource 接口提供了两个获取 Connection 的重载方法，并继承了 CommonDataSource 接口，该接口是 JDBC 中关于数据源定义的根接口。除了 DataSource 接口之外，它还有两个子接口：</p>

<p><img src="assets/CgqCHl7xtbuALDZqAABj4c2IofU664.png" alt="Drawing 1.png" /></p>

<p>其中，DataSource 是官方定义的获取 Connection 的基础接口，ConnectionPoolDataSource 是从连接池 ConnectionPool 中获取的 Connection 接口。而 XADataSource 则用来实现在分布式事务环境下获取 Connection，我们在讨论 ShardingSphere 的分布式事务时会接触到这个接口。</p>

<p><strong>请注意，DataSource 接口同时还继承了一个 Wrapper 接口</strong>。从接口的命名上看，可以判断该接口应该起到一种包装器的作用，事实上，由于很多数据库供应商提供了超越标准 JDBC API 的扩展功能，所以，Wrapper 接口可以把一个由第三方供应商提供的、非 JDBC 标准的接口包装成标准接口。以 DataSource 接口为例，如果我们想要实现自己的数据源 MyDataSource，就可以提供一个实现了 Wrapper 接口的 MyDataSourceWrapper 类来完成包装和适配：</p>

<p><img src="assets/CgqCHl7xtdOAZEGNAABnV-ZtNrk288.png" alt="Drawing 2.png" /></p>

<p>在 JDBC 规范中，除了 DataSource 之外，Connection、Statement、ResultSet 等核心对象也都继承了这个接口。显然，ShardingSphere 提供的就是非 JDBC 标准的接口，所以也应该会用到这个 Wrapper 接口，并提供了类似的实现方案。</p>

<h4 id="connection">Connection</h4>

<p>DataSource 的目的是获取 Connection 对象，我们可以把 Connection 理解为一种<strong>会话（Session）机制</strong>。Connection 代表一个数据库连接，负责完成与数据库之间的通信。所有 SQL 的执行都是在某个特定 Connection 环境中进行的，同时它还提供了一组重载方法，分别用于创建 Statement 和 PreparedStatement。另一方面，Connection 也涉及事务相关的操作，为了实现分片操作，ShardingSphere 同样也实现了定制化的 Connection 类 ShardingConnection。</p>

<h4 id="statement">Statement</h4>

<p>JDBC 规范中的 Statement 存在两种类型，一种是<strong>普通的 Statement</strong>，一种是<strong>支持预编译的 PreparedStatement</strong>。所谓预编译，是指数据库的编译器会对 SQL 语句提前编译，然后将预编译的结果缓存到数据库中，这样下次执行时就可以替换参数并直接使用编译过的语句，从而提高 SQL 的执行效率。当然，这种预编译也需要成本，所以在日常开发中，对数据库只执行一次性读写操作时，用 Statement 对象进行处理比较合适；而当涉及 SQL 语句的多次执行时，可以使用 PreparedStatement。</p>

<p>如果需要查询数据库中的数据，只需要调用 Statement 或 PreparedStatement 对象的 executeQuery 方法即可，该方法以 SQL 语句作为参数，执行完后返回一个 JDBC 的 ResultSet 对象。当然，Statement 或 PreparedStatement 中提供了一大批执行 SQL 更新和查询的重载方法。在 ShardingSphere 中，同样也提供了 ShardingStatement 和 ShardingPreparedStatement 这两个支持分片操作的 Statement 对象。</p>

<h4 id="resultset">ResultSet</h4>

<p>一旦通过 Statement 或 PreparedStatement 执行了 SQL 语句并获得了 ResultSet 对象后，那么就可以通过调用 Resulset 对象中的 next() 方法遍历整个结果集。如果 next() 方法返回为 true，就意味结果集中存在数据，则可以调用 ResultSet 对象的一系列 getXXX() 方法来取得对应的结果值。对于分库分表操作而言，因为涉及从多个数据库或数据表中获取目标数据，势必需要对获取的结果进行归并。因此，ShardingSphere 中也提供了分片环境下的 ShardingResultSet 对象。</p>

<p>作为总结，我们梳理了基于 JDBC 规范进行数据库访问的开发流程图，如下图所示：</p>

<p><img src="assets/Ciqc1F7xteqAQsj5AAB1bj_eu10085.png" alt="Drawing 3.png" /></p>

<p>ShardingSphere 提供了与 JDBC 规范完全兼容的 API。也就是说，开发人员可以基于这个开发流程和 JDBC 中的核心接口完成分片引擎、数据脱敏等操作，我们来看一下。</p>

<h3 id="基于适配器模式的-jdbc-重写实现方案">基于适配器模式的 JDBC 重写实现方案</h3>

<p>在 ShardingSphere 中，实现与 JDBC 规范兼容性的基本策略就是采用了设计模式中的适配器模式（Adapter Pattern）。<strong>适配器模式通常被用作连接两个不兼容接口之间的桥梁，涉及为某一个接口加入独立的或不兼容的功能。</strong></p>

<p>作为一套适配 JDBC 规范的实现方案，ShardingSphere 需要对上面介绍的 JDBC API 中的 DataSource、Connection、Statement 及 ResultSet 等核心对象都完成重写。虽然这些对象承载着不同功能，但重写机制应该是共通的，否则就需要对不同对象都实现定制化开发，显然，这不符合我们的设计原则。为此，ShardingSphere 抽象并开发了一套基于适配器模式的实现方案，整体结构是这样的，如下图所示：</p>

<p><img src="assets/Ciqc1F7xtfeAIlV7AABhpWkSy7c199.png" alt="Drawing 4.png" /></p>

<p>首先，我们看到这里有一个 JdbcObject 接口，这个接口泛指 JDBC API 中的 DataSource、Connection、Statement 等核心接口。前面提到，这些接口都继承自包装器 Wrapper 接口。ShardingSphere 为这个 Wrapper 接口提供了一个实现类 WrapperAdapter，这点在图中得到了展示。在 ShardingSphere 代码工程 sharding-jdbc-core 的 org.apache.shardingsphere.shardingjdbc.jdbc.adapter 包中包含了所有与 Adapter 相关的实现类：</p>

<p><img src="assets/Ciqc1F7xtgWAb3PaAAAW8D9SY1w475.png" alt="Drawing 5.png" /></p>

<p>在 ShardingSphere 基于适配器模式的实现方案图的底部，有一个 ShardingJdbcObject 类的定义。这个类也是一种泛指，代表 ShardingSphere 中用于分片的 ShardingDataSource、ShardingConnection、ShardingStatement 等对象。</p>

<p>最后发现 ShardingJdbcObject 继承自一个 AbstractJdbcObjectAdapter，而 AbstractJdbcObjectAdapter 又继承自 AbstractUnsupportedOperationJdbcObject，这两个类都是抽象类，而且也都泛指一组类。两者的区别在于，AbstractJdbcObjectAdapter 只提供了针对 JdbcObject 接口的一部分实现方法，这些方法是我们完成分片操作所需要的。而对于那些我们不需要的方法实现，则全部交由 AbstractUnsupportedOperationJdbcObject 进行实现，这两个类的所有方法的合集，就是原有 JdbcObject 接口的所有方法定义。</p>

<p>这样，我们大致了解了 ShardingSphere 对 JDBC 规范中核心接口的重写机制。这个重写机制非常重要，在 ShardingSphere 中应用也很广泛，我们可以通过示例对这一机制做进一步理解。</p>

<h3 id="shardingsphere-重写-jdbc-规范示例-shardingconnection">ShardingSphere 重写 JDBC 规范示例：ShardingConnection</h3>

<p>通过前面的介绍，我们知道 ShardingSphere 的分片引擎中提供了一系列 ShardingJdbcObject 来支持分片操作，包括 ShardingDataSource、ShardingConnection、ShardingStatement、ShardingPreparedStament 等。这里以最具代表性的 ShardingConnection 为例，来讲解它的实现过程。请注意，今天我们关注的还是重写机制，不会对 ShardingConnection 中的具体功能以及与其他类之间的交互过程做过多展开讲解。</p>

<h4 id="shardingconnection-类层结构">ShardingConnection 类层结构</h4>

<p>ShardingConnection 是对 JDBC 中 Connection 的适配和包装，所以它需要提供 Connection 接口中定义的方法，包括 createConnection、getMetaData、各种重载的 prepareStatement 和 createStatement 以及针对事务的 setAutoCommit、commit 和 rollback 方法等。ShardingConnection 对这些方法都进行了重写，如下图所示：</p>

<p><img src="assets/CgqCHl7xthSAHp4DAACJDJsvmyk879.png" alt="Drawing 6.png" />
ShardingConnection 中的方法列表图</p>

<p>ShardingConnection 类的一条类层结构支线就是适配器模式的具体应用，这部分内容的类层结构与前面介绍的重写机制的类层结构是完全一致的，如下图所示：</p>

<p><img src="assets/CgqCHl9MudqAOb9wAAEeGv0YTn807.jpeg" alt="111.jpeg" /></p>

<h4 id="abstractconnectionadapter">AbstractConnectionAdapter</h4>

<p>我们首先来看看 AbstractConnectionAdapter 抽象类，ShardingConnection 直接继承了它。在 AbstractConnectionAdapter 中发现了一个 cachedConnections 属性，它是一个 Map 对象，该对象其实缓存了这个经过封装的 ShardingConnection 背后真实的 Connection 对象。如果我们对一个 AbstractConnectionAdapter 重复使用，那么这些 cachedConnections 也会一直被缓存，直到调用 close 方法。可以从 AbstractConnectionAdapter 的 getConnections 方法中理解具体的操作过程：</p>

<pre><code>public final List&lt;Connection&gt; getConnections(final ConnectionMode connectionMode, final String dataSourceName, final int connectionSize) throws SQLException {
        //获取DataSource
        DataSource dataSource = getDataSourceMap().get(dataSourceName);
        Preconditions.checkState(null != dataSource, &quot;Missing the data source name: '%s'&quot;, dataSourceName);
        Collection&lt;Connection&gt; connections;

        //根据数据源从cachedConnections中获取connections
        synchronized (cachedConnections) {
            connections = cachedConnections.get(dataSourceName);
        }

        //如果connections多于想要的connectionSize，则只获取所需部分
        List&lt;Connection&gt; result;
        if (connections.size() &gt;= connectionSize) {
            result = new ArrayList&lt;&gt;(connections).subList(0, connectionSize);
        } else if (!connections.isEmpty()) {//如果connections不够
            result = new ArrayList&lt;&gt;(connectionSize);
            result.addAll(connections);
            //创建新的connections
            List&lt;Connection&gt; newConnections = createConnections(dataSourceName, connectionMode, dataSource, connectionSize - connections.size());
            result.addAll(newConnections);
            synchronized (cachedConnections) {
                //将新创建的connections也放入缓存中进行管理
                cachedConnections.putAll(dataSourceName, newConnections);
            }
        } else {//如果缓存中没有对应dataSource的Connections，同样进行创建并放入缓存中
            result = new ArrayList&lt;&gt;(createConnections(dataSourceName, connectionMode, dataSource, connectionSize));
            synchronized (cachedConnections) {
                cachedConnections.putAll(dataSourceName, result);
            }
        }
        return result;
}
</code></pre>

<p>这段代码有三个判断，流程上比较简单，参考注释即可，需要关注的是其中的 createConnections 方法：</p>

<pre><code>private List&lt;Connection&gt; createConnections(final String dataSourceName, final ConnectionMode connectionMode, final DataSource dataSource, final int connectionSize) throws SQLException {
        if (1 == connectionSize) {
            Connection connection = createConnection(dataSourceName, dataSource);
            replayMethodsInvocation(connection);
            return Collections.singletonList(connection);
        }
        if (ConnectionMode.CONNECTION_STRICTLY == connectionMode) {
            return createConnections(dataSourceName, dataSource, connectionSize);
        }
        synchronized (dataSource) {
            return createConnections(dataSourceName, dataSource, connectionSize);
        }
}
</code></pre>

<p>这段代码涉及了 ConnectionMode（连接模式），这是 ShardingSphere 执行引擎中的重要概念，今天我们先不展开，将在第 21 课时“执行引擎：分片环境下SQL执行的整体流程应该如何进行抽象？”中详细讲解。这里，可以看到 createConnections 方法批量调用了一个 createConnection 抽象方法，该方法需要 AbstractConnectionAdapter 的子类进行实现：</p>

<pre><code>protected abstract Connection createConnection(String dataSourceName, DataSource dataSource) throws SQLException;
</code></pre>

<p>同时，我们看到对于创建的 Connection 对象，都需要执行这样一个语句：</p>

<pre><code>replayMethodsInvocation(connection);
</code></pre>

<p>这行代码比较难以理解，让我们来到定义它的地方，即 WrapperAdapter 类。</p>

<h4 id="wrapperadapter">WrapperAdapter</h4>

<p>从命名上看，WrapperAdapter 是一个包装器的适配类，实现了 JDBC 中的 Wrapper 接口。我们在该类中找到了这样一对方法定义：</p>

<pre><code>    //记录方法调用
    public final void recordMethodInvocation(final Class&lt;?&gt; targetClass, final String methodName, final Class&lt;?&gt;[] argumentTypes, final Object[] arguments) {
        jdbcMethodInvocations.add(new JdbcMethodInvocation(targetClass.getMethod(methodName, argumentTypes), arguments));
	}
	 
	//重放方法调用
    public final void replayMethodsInvocation(final Object target) {
        for (JdbcMethodInvocation each : jdbcMethodInvocations) {
            each.invoke(target);
        }
	}
</code></pre>

<p>这两个方法都用到了 JdbcMethodInvocation 类：</p>

<pre><code>public class JdbcMethodInvocation {

    @Getter
    private final Method method;

    @Getter
    private final Object[] arguments;

    public void invoke(final Object target) {
        method.invoke(target, arguments);
    }
}
</code></pre>

<p>显然，JdbcMethodInvocation 类中用到了反射技术根据传入的 method 和 arguments 对象执行对应方法。</p>

<p>了解了 JdbcMethodInvocation 类的原理后，我们就不难理解 recordMethodInvocation 和 replayMethodsInvocation 方法的作用。其中，recordMethodInvocation 用于记录需要执行的方法和参数，而 replayMethodsInvocation 则根据这些方法和参数通过反射技术进行执行。</p>

<p>对于执行 replayMethodsInvocation，我们必须先找到 recordMethodInvocation 的调用入口。通过代码的调用关系，可以看到在 AbstractConnectionAdapter 中对其进行了调用，具体来说就是 setAutoCommit、setReadOnly 和 setTransactionIsolation 这三处方法。这里以 setReadOnly 方法为例给出它的实现：</p>

<pre><code>    @Override
    public final void setReadOnly(final boolean readOnly) throws SQLException {
        this.readOnly = readOnly; 
        //调用recordMethodInvocation方法记录方法调用的元数据
        recordMethodInvocation(Connection.class, &quot;setReadOnly&quot;, new Class[]{boolean.class}, new Object[]{readOnly});

        //执行回调
        forceExecuteTemplate.execute(cachedConnections.values(), new ForceExecuteCallback&lt;Connection&gt;() {

            @Override
            public void execute(final Connection connection) throws SQLException {
                connection.setReadOnly(readOnly);
            }
        });
    }
</code></pre>

<h4 id="abstractunsupportedoperationconnection">AbstractUnsupportedOperationConnection</h4>

<p>另一方面，从类层关系上，可以看到 AbstractConnectionAdapter 直接继承的是 AbstractUnsupportedOperationConnection 而不是 WrapperAdapter，而在 AbstractUnsupportedOperationConnection 中都是一组直接抛出异常的方法。这里截取部分代码：</p>

<pre><code>public abstract class AbstractUnsupportedOperationConnection extends WrapperAdapter implements Connection {

    @Override
    public final CallableStatement prepareCall(final String sql) throws SQLException {
        throw new SQLFeatureNotSupportedException(&quot;prepareCall&quot;);
    }

    @Override
    public final CallableStatement prepareCall(final String sql, final int resultSetType, final int resultSetConcurrency) throws SQLException {
        throw new SQLFeatureNotSupportedException(&quot;prepareCall&quot;);
}
…
}
</code></pre>

<p>AbstractUnsupportedOperationConnection 这种处理方式的目的就是明确哪些操作是 AbstractConnectionAdapter 及其子类 ShardingConnection 所不能支持的，属于职责分离的一种具体实现方法。</p>

<h3 id="小结">小结</h3>

<p>JDBC 规范是理解和应用 ShardingSphere 的基础，ShardingSphere 对 JDBC 规范进行了重写，并提供了完全兼容的一套接口。在这一课时中，我们首先给出了 JDBC 规范中各个核心接口的介绍；正是在这些接口的基础上，ShardingSphere 基于适配器模式对 JDBC 规范进行了重写；我们对这一重写方案进行了抽象，并基于 ShardingConnectin 类的实现过程详细阐述了 ShardingSphere 重写 Connection 接口的源码分析。</p>

<p>这里给你留一道思考题：ShardingSphere如何基于适配器模式实现对JDBC中核心类的重写？</p>

<p>JDBC 规范与 ShardingSphere 的兼容性概念至关重要。在掌握了这个概念之后，下一课时将介绍应用集成方面的话题，即在业务系统中使用 ShardingSphere 的具体方式。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a4c8c8c89d9095959493e4c3c9c5cdc88ac7cbc9" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f125cf96a0eede4',t:'MTczNDA1NTI1NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
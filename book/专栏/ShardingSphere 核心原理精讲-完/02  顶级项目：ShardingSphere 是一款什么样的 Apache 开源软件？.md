<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=02&#32;&#32;顶级项目：ShardingSphere&#32;是一款什么样的&#32;Apache&#32;开源软件？>
        <link rel="icon" href="/static/favicon.png">
        <title>02  顶级项目：ShardingSphere 是一款什么样的 Apache 开源软件？ </title>
        
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
                            <h1 id="title" data-id="02  顶级项目：ShardingSphere 是一款什么样的 Apache 开源软件？" class="title">02  顶级项目：ShardingSphere 是一款什么样的 Apache 开源软件？</h1>
                            <div><p>本课时将为你讲解 ShardingSphere 是一款什么样的 Apache 开源软件。</p>

<p>在上一课时中，我详细分析了分库分表的表现形式以及分片架构的解决方案和代表性框架。可以看到，ShardingSphere 同时实现了客户端分片和代理服务器组件，并提供了分布式数据库的相关功能特性。作为一款优秀的开源软件，ShardingSphere 能够取得目前的成就也不是一蹴而就，下面我们先来回顾一下 ShardingSphere 的发展历程。</p>

<h3 id="shardingsphere-的发展历程-从-sharding-jdbc-到-apache-顶级项目">ShardingSphere 的发展历程：从 Sharding-JDBC 到 Apache 顶级项目</h3>

<p>说到 ShardingSphere 的起源，我们不得不提 Sharding-JDBC 框架，该框架是一款起源于当当网内部的应用框架，并于 2017 年初正式开源。从 Sharding-JDBC 到 Apache 顶级项目，ShardingSphere 的发展经历了不同的演进阶段。纵观整个 ShardingSphere 的发展历史，我们可以得到时间线与阶段性里程碑的演进过程图：</p>

<p><img src="assets/CgqCHl7-6OWAVF7WAACmgRfIB7Y558.png" alt="1.png" /></p>

<p>从版本发布角度，我们也可以进一步梳理 ShardingSphere 发展历程中主线版本与核心功能之间的演进关系图：</p>

<p><img src="assets/Ciqc1F7rSMqAGrqLAABmt5OcOuc557.png" alt="2.png" /></p>

<p>基于 GitHub 上星数的增长轨迹，也可以从另一个维度很好地反映出 ShardingSphere 的发展历程：</p>

<p><img src="assets/CgqCHl7rSNaAE3gNAADqRUDMk-w922.png" alt="3.png" /></p>

<h3 id="shardingsphere-的设计理念-不是颠覆-而是兼容">ShardingSphere 的设计理念：不是颠覆，而是兼容</h3>

<p>对于一款开源中间件来说，要得到长足的发展，一方面依赖于社区的贡献，另外在很大程度上还取决于自身的设计和发展理念。</p>

<p>ShardingSphere 的定位非常明确，就是一种关系型数据库中间件，而并非一个全新的关系型数据库。ShardingSphere 认为，在当下，关系型数据库依然占有巨大市场，但凡涉及数据的持久化，关系型数据库仍然是系统的标准配置，也是各个公司核心业务的基石，在可预见的未来中，这点很难撼动。所以，<strong>ShardingSphere 在当前阶段更加关注在原有基础上进行兼容和扩展，而非颠覆</strong>。那么 ShardingSphere 是如何做到这一点呢？</p>

<p>ShardingSphere 构建了一个生态圈，这个生态圈由一套开源的分布式数据库中间件解决方案所构成。按照目前的规划，ShardingSphere 由 Sharding-JDBC、Sharding-Proxy 和 Sharding-Sidecar 这三款相互独立的产品组成，其中前两款已经正式发布，而 Sharding-Sidecar 正在规划中。我们可以从这三款产品出发，分析 ShardingSphere 的设计理念。</p>

<h4 id="sharding-jdbc">Sharding-JDBC</h4>

<p>ShardingSphere 的前身是 Sharding-JDBC，所以这是整个框架中最为成熟的组件。Sharding-JDBC 的定位是一个轻量级 Java 框架，在 JDBC 层提供了扩展性服务。我们知道 JDBC 是一种开发规范，指定了 DataSource、Connection、Statement、PreparedStatement、ResultSet 等一系列接口。而各大数据库供应商通过实现这些接口提供了自身对 JDBC 规范的支持，使得 JDBC 规范成为 Java 领域中被广泛采用的数据库访问标准。</p>

<p>基于这一点，Sharding-JDBC 一开始的设计就完全兼容 JDBC 规范，Sharding-JDBC 对外暴露的一套分片操作接口与 JDBC 规范中所提供的接口完全一致。开发人员只需要了解 JDBC，就可以使用 Sharding-JDBC 来实现分库分表，Sharding-JDBC 内部屏蔽了所有的分片规则和处理逻辑的复杂性。显然，<strong>这种方案天生就是一种具有高度兼容性的方案，能够为开发人员提供最简单、最直接的开发支持</strong>。关于 Sharding-JDBC 与 JDBC 规范的兼容性话题，我们将会在下一课时中详细讨论。</p>

<p><img src="assets/CgqCHl7rSOuAXZt6AAC4cmjERnk488.png" alt="4.png" />
Sharding-JDBC 与 JDBC 规范的兼容性示意图</p>

<p>在实际开发过程中，Sharding-JDBC 以 JAR 包的形式提供服务。<strong>开发人员可以使用这个 JAR 包直连数据库，无需额外的部署和依赖管理</strong>。在应用 Sharding-JDBC 时，需要注意到 Sharding-JDBC 背后依赖的是一套完整而强大的分片引擎：</p>

<p><img src="assets/CgqCHl7rSPSAUJHuAADsN1Pqjqs981.png" alt="5.png" /></p>

<p>由于 Sharding-JDBC 提供了一套与 JDBC 规范完全一致的 API，所以它可以很方便地与遵循 JDBC 规范的各种组件和框架进行无缝集成。例如，用于提供数据库连接的 DBCP、C3P0 等数据库连接池组件，以及用于提供对象-关系映射的 Hibernate、MyBatis 等 ORM 框架。当然，作为一款支持多数据库的开源框架，Sharding-JDBC 支持 MySQL、Oracle、SQLServer 等主流关系型数据库。</p>

<h4 id="sharding-proxy">Sharding-Proxy</h4>

<p>ShardingSphere 中的 Sharding-Proxy 组件定位为一个透明化的数据库代理端，所以它是代理服务器分片方案的一种具体实现方式。在代理方案的设计和实现上，Sharding-Proxy 同样充分考虑了兼容性。</p>

<p>Sharding-Proxy 所提供的兼容性首先体现在对异构语言的支持上，为了完成对异构语言的支持，Sharding-Proxy 专门对数据库二进制协议进行了封装，并提供了一个代理服务端组件。其次，从客户端组件上讲，针对目前市面上流行的 Navicat、MySQL Command Client 等客户端工具，Sharding-Proxy 也能够兼容遵循 MySQL 和 PostgreSQL 协议的各类访问客户端。当然，和 Sharding-JDBC 一样，Sharding-Proxy 也支持 MySQL 和 PostgreSQL 等多种数据库。</p>

<p>接下来，我们看一下 Sharding-Proxy 的整体架构。对于应用程序而言，这种代理机制是完全透明的，可以直接把它当作 MySQL 或 PostgreSQL 进行使用：</p>

<p><img src="assets/CgqCHl7rSQKAC1u6AADoQEq6hys890.png" alt="6.png" /></p>

<p>总结一下，我们可以直接把 Sharding-Proxy 视为一个数据库，用来代理后面分库分表的多个数据库，它屏蔽了后端多个数据库的复杂性。同时，也看到 Sharding-Proxy 的运行同样需要依赖于完成分片操作的分片引擎以及用于管理数据库的治理组件。</p>

<p>虽然 Sharding-JDBC 和 Sharding-Proxy 具有不同的关注点，但事实上，我们完全可以将它们整合在一起进行使用，也就是说这两个组件之间也存在兼容性。</p>

<p>前面已经介绍过，我们使用 Sharding-JDBC 的方式是在应用程序中直接嵌入 JAR 包，这种方式适合于业务开发人员。而 Sharding-Proxy 提供静态入口以及异构语言的支持，适用于需要对分片数据库进行管理的中间件开发和运维人员。基于底层共通的分片引擎，以及数据库治理功能，可以混合使用 Sharding-JDBC 和 Sharding-Proxy，以便应对不同的应用场景和不同的开发人员：</p>

<p><img src="assets/Ciqc1F7rSRCAS66OAAECtLTiErU161.png" alt="7.png" /></p>

<h4 id="sharding-sidecar">Sharding-Sidecar</h4>

<p>Sidecar 设计模式受到了越来越多的关注和采用，这个模式的目标是把系统中各种异构的服务组件串联起来，并进行高效的服务治理。ShardingSphere 也基于该模式设计了 Sharding-Sidecar 组件。截止到目前，ShardingSphere 给出了 Sharding-Sidecar 的规划，但还没有提供具体的实现方案，这里不做具体展开。作为 Sidecar 模式的具体实现，我们可以想象 <strong>Sharding-Sidecar</strong>** 的作用就是以 Sidecar 的形式代理所有对数据库的访问**。这也是一种兼容性的设计思路，通过无中心、零侵入的方案将分布式的数据访问应用与数据库有机串联起来。</p>

<h3 id="shardingsphere-的核心功能-从数据分片到编排治理">ShardingSphere 的核心功能：从数据分片到编排治理</h3>

<p>介绍完 ShardingSphere 的设计理念之后，我们再来关注它的核心功能和实现机制。这里把 ShardingSphere 的整体功能拆分成四大部分，即<strong>基础设施</strong>、<strong>分片引擎</strong>、<strong>分布式事务</strong>和<strong>治理与集成</strong>，这四大部分也构成了本课程介绍 ShardingSphere 的整体行文结构，下面我们来分别进行介绍：</p>

<h4 id="基础设施">基础设施</h4>

<p>作为一款开源框架，ShardingSphere 在架构上也提供了很多基础设施类的组件，这些组件更多与它的内部实现机制有关，我们将会在后续的源码解析部分详细展开讲解。但对开发人员而言，可以认为微内核架构和分布式主键是框架提供的基础设施类的核心功能。</p>

<ul>
<li>微内核架构</li>
</ul>

<p>ShardingSphere 在设计上采用了<strong>微内核（MicroKernel）架构模式</strong>，来确保系统具有高度可扩展性。微内核架构包含两部分组件，即内核系统和插件。使用微内核架构对系统进行升级，要做的只是用新插件替换旧插件，而不需要改变整个系统架构：</p>

<p><img src="assets/CgqCHl7rSSCAcY4sAABRDnG4TnQ180.png" alt="8.png" /></p>

<p>在 ShardingSphere 中，抽象了一大批插件接口，包含用实现 SQL 解析的 SQLParserEntry、用于实现配置中心的 ConfigCenter、用于数据脱敏的 ShardingEncryptor，以及用于数据库治理的注册中心接口 RegistryCenter 等。开发人员完全可以根据自己的需要，基于这些插件定义来提供定制化实现，并动态加载到 ShardingSphere 运行时环境中。</p>

<ul>
<li>分布式主键</li>
</ul>

<p>在本地数据库场景下，我们可以使用数据库自带的自增序列来完成主键的生成。但在分片场景下，我们将面对从本地数据库转换到分布式数据库的应用场景，这就需要考虑主键在各个数据库中的全局唯一性。为此，我们需要引入分布式主键机制。ShardingSphere 同样提供了分布式主键的实现机制，默认采用的是 SnowFlake（雪花）算法。</p>

<h4 id="分片引擎">分片引擎</h4>

<p>对于分片引擎，ShardingSphere 同时支持数据分片和读写分离机制。</p>

<ul>
<li>数据分片</li>
</ul>

<p>数据分片是 ShardingSphere 的核心功能，常规的，基于垂直拆分和水平拆分的分库分表操作它都支持。同时，ShardingSphere 也预留了分片扩展点，开发人员也可以基于需要实现分片策略的定制化开发。</p>

<ul>
<li>读写分离</li>
</ul>

<p>在分库分表的基础上，ShardingSphere 也实现了基于数据库主从架构的读写分离机制。而且，这种读写分离机制可以和数据分片完美地进行整合。</p>

<h4 id="分布式事务">分布式事务</h4>

<p>分布式事务是分布式环境下确保数据一致性的基本功能，作为分布式数据库的一种生态圈，ShardingSphere 也提供了对分布式事务的全面支持。</p>

<ul>
<li>标准化事务处理接口</li>
</ul>

<p>ShardingSphere 支持本地事务、基于 XA 两阶段提交的强一致性事务以及基于 BASE 的柔性最终一致性事务。同时，ShardingSphere 抽象了一组标准化的事务处理接口，并通过分片事务管理器 ShardingTransactionManager 进行统一管理。我们也可以根据需要实现自己的 ShardingTransactionManager 从而对分布式事务进行扩展。</p>

<ul>
<li>强一致性事务与柔性事务</li>
</ul>

<p>ShardingSphere 内置了一组分布式事务的实现方案，其中强一致性事务内置集成了 Atomikos、Narayana 和 Bitronix 等技术来实现 XA 事务管理器；另一方面，ShardingSphere 内部也整合了 Seata 来提供柔性事务功能。</p>

<h4 id="治理与集成">治理与集成</h4>

<p>对于分布式数据库而言，治理的范畴可以很广，ShardingSphere 也提供了注册中心、配置中心等一系列功能来支持数据库治理。另一方面，ShardingSphere 作为一款支持快速开发的开源框架，也完成了与其他主流框架的无缝集成。</p>

<ul>
<li>数据脱敏</li>
</ul>

<p>数据脱敏是确保数据访问安全的常见需求，通常做法是对原始的 SQL 进行改写，从而实现对原文数据进行加密。当我们想要获取原始数据时，在实现上就需要通过对数据库中所存储的密文数据进行解密才能完成。我们可以根据需要实现一套类似的加解密机制，但 ShardingSphere 的强大之处在于，<strong>它将这套机制内嵌到了 SQL 的执行过程中，业务开发人员不需要关注具体的加解密实现细节，而只要通过简单的配置就能实现数据的自动脱敏。</strong></p>

<ul>
<li>配置中心</li>
</ul>

<p>关于配置信息的管理，我们可以基于 YAML 格式或 XML 格式的配置文件完成配置信息的维护，这在 ShardingSphere 中都得到了支持。更进一步，在 ShardingSphere 中，它还提供了配置信息动态化的管理机制，可以支持数据源、表与分片及读写分离策略的动态切换。</p>

<ul>
<li>注册中心</li>
</ul>

<p>相较配置中心，注册中心在 ShardingSphere 中的应用更为广泛。ShardingSphere 中的注册中心提供了基于 Nacos 和 ZooKeeper 的两种实现方式。而在应用场景上，我们可以基于注册中心完成数据库实例管理、数据库熔断禁用等治理功能。</p>

<ul>
<li>链路跟踪</li>
</ul>

<p>SQL 解析与 SQL 执行是数据分片的最核心步骤，ShardingSphere 在完成这两个步骤的同时，也会将运行时的数据通过标准协议提交到链路跟踪系统。ShardingSphere 使用 OpenTracing API 发送性能追踪数据。像 SkyWalking、Zipkin 和 Jaeger 等面向 OpenTracing 协议的具体产品都可以和 ShardingSphere 自动完成对接。</p>

<ul>
<li>系统集成</li>
</ul>

<p>这里所谓的系统集成，指的是 ShardingSphere 和 Spring 系列框架的集成。到目前为止，ShardingSphere 实现了两种系统的集成机制，一种是命名空间机制，即通过扩展 Spring Schema 来实现与 Spring 框架的集成；而另一种则是通过编写自定义的 starter 组件来完成与 Spring Boot 的集成。这样，无论开发人员采用哪一种 Spring 框架，对于使用 ShardingSphere 而言都是零学习成本。</p>

<h3 id="小结">小结</h3>

<p>从今天开始，我们正式引入了 ShardingSphere。本课时回顾了 ShardingSphere 的发展历程，并结合该框架所提供的三大核心产品 Sharding-JDBC、Sharding-Proxy 及 Sharding-Sidecar 阐述了其在设计上的思想和理念。同时，我们也给出了 ShardingSphere 作为分布式数据库所具有的分片引擎、分布式事务及治理与集成方面的核心功能，这些功能在课程的后续内容都会进行详细展开讲解。</p>

<p>这里给你留一道思考题：ShardingSphere是一款兼容性很强的开源框架，它的兼容性具体体现在哪些方面？</p>

<p>通过前面内容的介绍，我们知道了 ShardingSphere 实现分片引擎的方案是重写 JDBC 规范，从而为应用程序提供与 JDBC 完全兼容的使用方式。下一课时将对 ShardingSphere 与 JDBC 规范之间的这层关系展开讨论。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#503c3c3c69646161606710373d31393c7e333f3d" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f125ca399c2ede4',t:'MTczNDA1NTI0MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
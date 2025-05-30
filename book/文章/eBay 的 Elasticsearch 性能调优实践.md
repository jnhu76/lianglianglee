<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=eBay&#32;的&#32;Elasticsearch&#32;性能调优实践>
        <link rel="icon" href="/static/favicon.png">
        <title>eBay 的 Elasticsearch 性能调优实践 </title>
        
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
                        <a class="menu-item" id="AQS 万字图文全面解析.md" href="/%e6%96%87%e7%ab%a0/AQS%20%e4%b8%87%e5%ad%97%e5%9b%be%e6%96%87%e5%85%a8%e9%9d%a2%e8%a7%a3%e6%9e%90.md">AQS 万字图文全面解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Docker 镜像构建原理及源码分析.md" href="/%e6%96%87%e7%ab%a0/Docker%20%e9%95%9c%e5%83%8f%e6%9e%84%e5%bb%ba%e5%8e%9f%e7%90%86%e5%8f%8a%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90.md">Docker 镜像构建原理及源码分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="ElasticSearch 小白从入门到精通.md" href="/%e6%96%87%e7%ab%a0/ElasticSearch%20%e5%b0%8f%e7%99%bd%e4%bb%8e%e5%85%a5%e9%97%a8%e5%88%b0%e7%b2%be%e9%80%9a.md">ElasticSearch 小白从入门到精通.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="JVM CPU Profiler技术原理及源码深度解析.md" href="/%e6%96%87%e7%ab%a0/JVM%20CPU%20Profiler%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86%e5%8f%8a%e6%ba%90%e7%a0%81%e6%b7%b1%e5%ba%a6%e8%a7%a3%e6%9e%90.md">JVM CPU Profiler技术原理及源码深度解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="JVM 垃圾收集器.md" href="/%e6%96%87%e7%ab%a0/JVM%20%e5%9e%83%e5%9c%be%e6%94%b6%e9%9b%86%e5%99%a8.md">JVM 垃圾收集器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="JVM 面试的 30 个知识点.md" href="/%e6%96%87%e7%ab%a0/JVM%20%e9%9d%a2%e8%af%95%e7%9a%84%2030%20%e4%b8%aa%e7%9f%a5%e8%af%86%e7%82%b9.md">JVM 面试的 30 个知识点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java IO 体系、线程模型大总结.md" href="/%e6%96%87%e7%ab%a0/Java%20IO%20%e4%bd%93%e7%b3%bb%e3%80%81%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b%e5%a4%a7%e6%80%bb%e7%bb%93.md">Java IO 体系、线程模型大总结.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java NIO浅析.md" href="/%e6%96%87%e7%ab%a0/Java%20NIO%e6%b5%85%e6%9e%90.md">Java NIO浅析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java 面试题集锦（网络篇）.md" href="/%e6%96%87%e7%ab%a0/Java%20%e9%9d%a2%e8%af%95%e9%a2%98%e9%9b%86%e9%94%a6%ef%bc%88%e7%bd%91%e7%bb%9c%e7%af%87%ef%bc%89.md">Java 面试题集锦（网络篇）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java-直接内存 DirectMemory 详解.md" href="/%e6%96%87%e7%ab%a0/Java-%e7%9b%b4%e6%8e%a5%e5%86%85%e5%ad%98%20DirectMemory%20%e8%af%a6%e8%a7%a3.md">Java-直接内存 DirectMemory 详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java中9种常见的CMS GC问题分析与解决（上）.md" href="/%e6%96%87%e7%ab%a0/Java%e4%b8%ad9%e7%a7%8d%e5%b8%b8%e8%a7%81%e7%9a%84CMS%20GC%e9%97%ae%e9%a2%98%e5%88%86%e6%9e%90%e4%b8%8e%e8%a7%a3%e5%86%b3%ef%bc%88%e4%b8%8a%ef%bc%89.md">Java中9种常见的CMS GC问题分析与解决（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java中9种常见的CMS GC问题分析与解决（下）.md" href="/%e6%96%87%e7%ab%a0/Java%e4%b8%ad9%e7%a7%8d%e5%b8%b8%e8%a7%81%e7%9a%84CMS%20GC%e9%97%ae%e9%a2%98%e5%88%86%e6%9e%90%e4%b8%8e%e8%a7%a3%e5%86%b3%ef%bc%88%e4%b8%8b%ef%bc%89.md">Java中9种常见的CMS GC问题分析与解决（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java中的SPI.md" href="/%e6%96%87%e7%ab%a0/Java%e4%b8%ad%e7%9a%84SPI.md">Java中的SPI.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java中的ThreadLocal.md" href="/%e6%96%87%e7%ab%a0/Java%e4%b8%ad%e7%9a%84ThreadLocal.md">Java中的ThreadLocal.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java线程池实现原理及其在美团业务中的实践.md" href="/%e6%96%87%e7%ab%a0/Java%e7%ba%bf%e7%a8%8b%e6%b1%a0%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%e5%8f%8a%e5%85%b6%e5%9c%a8%e7%be%8e%e5%9b%a2%e4%b8%9a%e5%8a%a1%e4%b8%ad%e7%9a%84%e5%ae%9e%e8%b7%b5.md">Java线程池实现原理及其在美团业务中的实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Java魔法类：Unsafe应用解析.md" href="/%e6%96%87%e7%ab%a0/Java%e9%ad%94%e6%b3%95%e7%b1%bb%ef%bc%9aUnsafe%e5%ba%94%e7%94%a8%e8%a7%a3%e6%9e%90.md">Java魔法类：Unsafe应用解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Kafka 源码阅读笔记.md" href="/%e6%96%87%e7%ab%a0/Kafka%20%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb%e7%ac%94%e8%ae%b0.md">Kafka 源码阅读笔记.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Kafka、ActiveMQ、RabbitMQ、RocketMQ 区别以及高可用原理.md" href="/%e6%96%87%e7%ab%a0/Kafka%e3%80%81ActiveMQ%e3%80%81RabbitMQ%e3%80%81RocketMQ%20%e5%8c%ba%e5%88%ab%e4%bb%a5%e5%8f%8a%e9%ab%98%e5%8f%af%e7%94%a8%e5%8e%9f%e7%90%86.md">Kafka、ActiveMQ、RabbitMQ、RocketMQ 区别以及高可用原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB Buffer Pool.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%20Buffer%20Pool.md">MySQL · 引擎特性 · InnoDB Buffer Pool.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB IO子系统.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%20IO%e5%ad%90%e7%b3%bb%e7%bb%9f.md">MySQL · 引擎特性 · InnoDB IO子系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB 事务系统.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%20%e4%ba%8b%e5%8a%a1%e7%b3%bb%e7%bb%9f.md">MySQL · 引擎特性 · InnoDB 事务系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB 同步机制.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%20%e5%90%8c%e6%ad%a5%e6%9c%ba%e5%88%b6.md">MySQL · 引擎特性 · InnoDB 同步机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB 数据页解析.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%20%e6%95%b0%e6%8d%ae%e9%a1%b5%e8%a7%a3%e6%9e%90.md">MySQL · 引擎特性 · InnoDB 数据页解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · InnoDB崩溃恢复.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20InnoDB%e5%b4%a9%e6%ba%83%e6%81%a2%e5%a4%8d.md">MySQL · 引擎特性 · InnoDB崩溃恢复.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL · 引擎特性 · 临时表那些事儿.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%c2%b7%20%e5%bc%95%e6%93%8e%e7%89%b9%e6%80%a7%20%c2%b7%20%e4%b8%b4%e6%97%b6%e8%a1%a8%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">MySQL · 引擎特性 · 临时表那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 主从复制 半同步复制.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%b8%bb%e4%bb%8e%e5%a4%8d%e5%88%b6%20%e5%8d%8a%e5%90%8c%e6%ad%a5%e5%a4%8d%e5%88%b6.md">MySQL 主从复制 半同步复制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 主从复制 基于GTID复制.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%b8%bb%e4%bb%8e%e5%a4%8d%e5%88%b6%20%e5%9f%ba%e4%ba%8eGTID%e5%a4%8d%e5%88%b6.md">MySQL 主从复制 基于GTID复制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 主从复制.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%b8%bb%e4%bb%8e%e5%a4%8d%e5%88%b6.md">MySQL 主从复制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 事务日志(redo log和undo log).md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%ba%8b%e5%8a%a1%e6%97%a5%e5%bf%97%28redo%20log%e5%92%8cundo%20log%29.md">MySQL 事务日志(redo log和undo log).md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 亿级别数据迁移实战代码分享.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%ba%bf%e7%ba%a7%e5%88%ab%e6%95%b0%e6%8d%ae%e8%bf%81%e7%a7%bb%e5%ae%9e%e6%88%98%e4%bb%a3%e7%a0%81%e5%88%86%e4%ba%ab.md">MySQL 亿级别数据迁移实战代码分享.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 从一条数据说起-InnoDB行存储数据结构.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e4%bb%8e%e4%b8%80%e6%9d%a1%e6%95%b0%e6%8d%ae%e8%af%b4%e8%b5%b7-InnoDB%e8%a1%8c%e5%ad%98%e5%82%a8%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84.md">MySQL 从一条数据说起-InnoDB行存储数据结构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 地基基础：事务和锁的面纱.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e5%9c%b0%e5%9f%ba%e5%9f%ba%e7%a1%80%ef%bc%9a%e4%ba%8b%e5%8a%a1%e5%92%8c%e9%94%81%e7%9a%84%e9%9d%a2%e7%ba%b1.md">MySQL 地基基础：事务和锁的面纱.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 地基基础：数据字典.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e5%9c%b0%e5%9f%ba%e5%9f%ba%e7%a1%80%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%ad%97%e5%85%b8.md">MySQL 地基基础：数据字典.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 地基基础：数据库字符集.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e5%9c%b0%e5%9f%ba%e5%9f%ba%e7%a1%80%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%ba%93%e5%ad%97%e7%ac%a6%e9%9b%86.md">MySQL 地基基础：数据库字符集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 性能优化：碎片整理.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%9a%e7%a2%8e%e7%89%87%e6%95%b4%e7%90%86.md">MySQL 性能优化：碎片整理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 故障诊断：一个 ALTER TALBE 执行了很久，你慌不慌？.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e6%95%85%e9%9a%9c%e8%af%8a%e6%96%ad%ef%bc%9a%e4%b8%80%e4%b8%aa%20ALTER%20TALBE%20%e6%89%a7%e8%a1%8c%e4%ba%86%e5%be%88%e4%b9%85%ef%bc%8c%e4%bd%a0%e6%85%8c%e4%b8%8d%e6%85%8c%ef%bc%9f.md">MySQL 故障诊断：一个 ALTER TALBE 执行了很久，你慌不慌？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 故障诊断：如何在日志中轻松定位大事务.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e6%95%85%e9%9a%9c%e8%af%8a%e6%96%ad%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e6%97%a5%e5%bf%97%e4%b8%ad%e8%bd%bb%e6%9d%be%e5%ae%9a%e4%bd%8d%e5%a4%a7%e4%ba%8b%e5%8a%a1.md">MySQL 故障诊断：如何在日志中轻松定位大事务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 故障诊断：教你快速定位加锁的 SQL.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e6%95%85%e9%9a%9c%e8%af%8a%e6%96%ad%ef%bc%9a%e6%95%99%e4%bd%a0%e5%bf%ab%e9%80%9f%e5%ae%9a%e4%bd%8d%e5%8a%a0%e9%94%81%e7%9a%84%20SQL.md">MySQL 故障诊断：教你快速定位加锁的 SQL.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 日志详解.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e6%97%a5%e5%bf%97%e8%af%a6%e8%a7%a3.md">MySQL 日志详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL 的半同步是什么？.md" href="/%e6%96%87%e7%ab%a0/MySQL%20%e7%9a%84%e5%8d%8a%e5%90%8c%e6%ad%a5%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">MySQL 的半同步是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL中的事务和MVCC.md" href="/%e6%96%87%e7%ab%a0/MySQL%e4%b8%ad%e7%9a%84%e4%ba%8b%e5%8a%a1%e5%92%8cMVCC.md">MySQL中的事务和MVCC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL事务_事务隔离级别详解.md" href="/%e6%96%87%e7%ab%a0/MySQL%e4%ba%8b%e5%8a%a1_%e4%ba%8b%e5%8a%a1%e9%9a%94%e7%a6%bb%e7%ba%a7%e5%88%ab%e8%af%a6%e8%a7%a3.md">MySQL事务_事务隔离级别详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL优化：优化 select count().md" href="/%e6%96%87%e7%ab%a0/MySQL%e4%bc%98%e5%8c%96%ef%bc%9a%e4%bc%98%e5%8c%96%20select%20count%28%29.md">MySQL优化：优化 select count().md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL共享锁、排他锁、悲观锁、乐观锁.md" href="/%e6%96%87%e7%ab%a0/MySQL%e5%85%b1%e4%ba%ab%e9%94%81%e3%80%81%e6%8e%92%e4%bb%96%e9%94%81%e3%80%81%e6%82%b2%e8%a7%82%e9%94%81%e3%80%81%e4%b9%90%e8%a7%82%e9%94%81.md">MySQL共享锁、排他锁、悲观锁、乐观锁.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="MySQL的MVCC（多版本并发控制）.md" href="/%e6%96%87%e7%ab%a0/MySQL%e7%9a%84MVCC%ef%bc%88%e5%a4%9a%e7%89%88%e6%9c%ac%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%ef%bc%89.md">MySQL的MVCC（多版本并发控制）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="QingStor 对象存储架构设计及最佳实践.md" href="/%e6%96%87%e7%ab%a0/QingStor%20%e5%af%b9%e8%b1%a1%e5%ad%98%e5%82%a8%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%e5%8f%8a%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5.md">QingStor 对象存储架构设计及最佳实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="RocketMQ 面试题集锦.md" href="/%e6%96%87%e7%ab%a0/RocketMQ%20%e9%9d%a2%e8%af%95%e9%a2%98%e9%9b%86%e9%94%a6.md">RocketMQ 面试题集锦.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="SnowFlake 雪花算法生成分布式 ID.md" href="/%e6%96%87%e7%ab%a0/SnowFlake%20%e9%9b%aa%e8%8a%b1%e7%ae%97%e6%b3%95%e7%94%9f%e6%88%90%e5%88%86%e5%b8%83%e5%bc%8f%20ID.md">SnowFlake 雪花算法生成分布式 ID.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring Boot 2.x 结合 k8s 实现分布式微服务架构.md" href="/%e6%96%87%e7%ab%a0/Spring%20Boot%202.x%20%e7%bb%93%e5%90%88%20k8s%20%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e5%be%ae%e6%9c%8d%e5%8a%a1%e6%9e%b6%e6%9e%84.md">Spring Boot 2.x 结合 k8s 实现分布式微服务架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring Boot 教程：如何开发一个 starter.md" href="/%e6%96%87%e7%ab%a0/Spring%20Boot%20%e6%95%99%e7%a8%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bc%80%e5%8f%91%e4%b8%80%e4%b8%aa%20starter.md">Spring Boot 教程：如何开发一个 starter.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring MVC 原理.md" href="/%e6%96%87%e7%ab%a0/Spring%20MVC%20%e5%8e%9f%e7%90%86.md">Spring MVC 原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring MyBatis和Spring整合的奥秘.md" href="/%e6%96%87%e7%ab%a0/Spring%20MyBatis%e5%92%8cSpring%e6%95%b4%e5%90%88%e7%9a%84%e5%a5%a5%e7%a7%98.md">Spring MyBatis和Spring整合的奥秘.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring 帮助你更好的理解Spring循环依赖.md" href="/%e6%96%87%e7%ab%a0/Spring%20%e5%b8%ae%e5%8a%a9%e4%bd%a0%e6%9b%b4%e5%a5%bd%e7%9a%84%e7%90%86%e8%a7%a3Spring%e5%be%aa%e7%8e%af%e4%be%9d%e8%b5%96.md">Spring 帮助你更好的理解Spring循环依赖.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring 循环依赖及解决方式.md" href="/%e6%96%87%e7%ab%a0/Spring%20%e5%be%aa%e7%8e%af%e4%be%9d%e8%b5%96%e5%8f%8a%e8%a7%a3%e5%86%b3%e6%96%b9%e5%bc%8f.md">Spring 循环依赖及解决方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spring中眼花缭乱的BeanDefinition.md" href="/%e6%96%87%e7%ab%a0/Spring%e4%b8%ad%e7%9c%bc%e8%8a%b1%e7%bc%ad%e4%b9%b1%e7%9a%84BeanDefinition.md">Spring中眼花缭乱的BeanDefinition.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Vert.x 基础入门.md" href="/%e6%96%87%e7%ab%a0/Vert.x%20%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8.md">Vert.x 基础入门.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="eBay 的 Elasticsearch 性能调优实践.md" href="/%e6%96%87%e7%ab%a0/eBay%20%e7%9a%84%20Elasticsearch%20%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e8%b7%b5.md">eBay 的 Elasticsearch 性能调优实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不可不说的Java“锁”事.md" href="/%e6%96%87%e7%ab%a0/%e4%b8%8d%e5%8f%af%e4%b8%8d%e8%af%b4%e7%9a%84Java%e2%80%9c%e9%94%81%e2%80%9d%e4%ba%8b.md">不可不说的Java“锁”事.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="互联网并发限流实战.md" href="/%e6%96%87%e7%ab%a0/%e4%ba%92%e8%81%94%e7%bd%91%e5%b9%b6%e5%8f%91%e9%99%90%e6%b5%81%e5%ae%9e%e6%88%98.md">互联网并发限流实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="从ReentrantLock的实现看AQS的原理及应用.md" href="/%e6%96%87%e7%ab%a0/%e4%bb%8eReentrantLock%e7%9a%84%e5%ae%9e%e7%8e%b0%e7%9c%8bAQS%e7%9a%84%e5%8e%9f%e7%90%86%e5%8f%8a%e5%ba%94%e7%94%a8.md">从ReentrantLock的实现看AQS的原理及应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="从SpringCloud开始，聊微服务架构.md" href="/%e6%96%87%e7%ab%a0/%e4%bb%8eSpringCloud%e5%bc%80%e5%a7%8b%ef%bc%8c%e8%81%8a%e5%be%ae%e6%9c%8d%e5%8a%a1%e6%9e%b6%e6%9e%84.md">从SpringCloud开始，聊微服务架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="全面了解 JDK 线程池实现原理.md" href="/%e6%96%87%e7%ab%a0/%e5%85%a8%e9%9d%a2%e4%ba%86%e8%a7%a3%20JDK%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">全面了解 JDK 线程池实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="分布式一致性理论与算法.md" href="/%e6%96%87%e7%ab%a0/%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%80%e8%87%b4%e6%80%a7%e7%90%86%e8%ae%ba%e4%b8%8e%e7%ae%97%e6%b3%95.md">分布式一致性理论与算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="分布式一致性算法 Raft.md" href="/%e6%96%87%e7%ab%a0/%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%80%e8%87%b4%e6%80%a7%e7%ae%97%e6%b3%95%20Raft.md">分布式一致性算法 Raft.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="分布式唯一 ID 解析.md" href="/%e6%96%87%e7%ab%a0/%e5%88%86%e5%b8%83%e5%bc%8f%e5%94%af%e4%b8%80%20ID%20%e8%a7%a3%e6%9e%90.md">分布式唯一 ID 解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="分布式链路追踪：集群管理设计.md" href="/%e6%96%87%e7%ab%a0/%e5%88%86%e5%b8%83%e5%bc%8f%e9%93%be%e8%b7%af%e8%bf%bd%e8%b8%aa%ef%bc%9a%e9%9b%86%e7%be%a4%e7%ae%a1%e7%90%86%e8%ae%be%e8%ae%a1.md">分布式链路追踪：集群管理设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="动态代理种类及原理，你知道多少？.md" href="/%e6%96%87%e7%ab%a0/%e5%8a%a8%e6%80%81%e4%bb%a3%e7%90%86%e7%a7%8d%e7%b1%bb%e5%8f%8a%e5%8e%9f%e7%90%86%ef%bc%8c%e4%bd%a0%e7%9f%a5%e9%81%93%e5%a4%9a%e5%b0%91%ef%bc%9f.md">动态代理种类及原理，你知道多少？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="响应式架构与 RxJava 在有赞零售的实践.md" href="/%e6%96%87%e7%ab%a0/%e5%93%8d%e5%ba%94%e5%bc%8f%e6%9e%b6%e6%9e%84%e4%b8%8e%20RxJava%20%e5%9c%a8%e6%9c%89%e8%b5%9e%e9%9b%b6%e5%94%ae%e7%9a%84%e5%ae%9e%e8%b7%b5.md">响应式架构与 RxJava 在有赞零售的实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大数据算法——布隆过滤器.md" href="/%e6%96%87%e7%ab%a0/%e5%a4%a7%e6%95%b0%e6%8d%ae%e7%ae%97%e6%b3%95%e2%80%94%e2%80%94%e5%b8%83%e9%9a%86%e8%bf%87%e6%bb%a4%e5%99%a8.md">大数据算法——布隆过滤器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="如何优雅地记录操作日志？.md" href="/%e6%96%87%e7%ab%a0/%e5%a6%82%e4%bd%95%e4%bc%98%e9%9b%85%e5%9c%b0%e8%ae%b0%e5%bd%95%e6%93%8d%e4%bd%9c%e6%97%a5%e5%bf%97%ef%bc%9f.md">如何优雅地记录操作日志？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="如何设计一个亿级消息量的 IM 系统.md" href="/%e6%96%87%e7%ab%a0/%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e4%b8%80%e4%b8%aa%e4%ba%bf%e7%ba%a7%e6%b6%88%e6%81%af%e9%87%8f%e7%9a%84%20IM%20%e7%b3%bb%e7%bb%9f.md">如何设计一个亿级消息量的 IM 系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="异步网络模型.md" href="/%e6%96%87%e7%ab%a0/%e5%bc%82%e6%ad%a5%e7%bd%91%e7%bb%9c%e6%a8%a1%e5%9e%8b.md">异步网络模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="当我们在讨论CQRS时，我们在讨论些神马？.md" href="/%e6%96%87%e7%ab%a0/%e5%bd%93%e6%88%91%e4%bb%ac%e5%9c%a8%e8%ae%a8%e8%ae%baCQRS%e6%97%b6%ef%bc%8c%e6%88%91%e4%bb%ac%e5%9c%a8%e8%ae%a8%e8%ae%ba%e4%ba%9b%e7%a5%9e%e9%a9%ac%ef%bc%9f.md">当我们在讨论CQRS时，我们在讨论些神马？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="彻底理解 MySQL 的索引机制.md" href="/%e6%96%87%e7%ab%a0/%e5%bd%bb%e5%ba%95%e7%90%86%e8%a7%a3%20MySQL%20%e7%9a%84%e7%b4%a2%e5%bc%95%e6%9c%ba%e5%88%b6.md">彻底理解 MySQL 的索引机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="最全的 116 道 Redis 面试题解答.md" href="/%e6%96%87%e7%ab%a0/%e6%9c%80%e5%85%a8%e7%9a%84%20116%20%e9%81%93%20Redis%20%e9%9d%a2%e8%af%95%e9%a2%98%e8%a7%a3%e7%ad%94.md">最全的 116 道 Redis 面试题解答.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="有赞权限系统(SAM).md" href="/%e6%96%87%e7%ab%a0/%e6%9c%89%e8%b5%9e%e6%9d%83%e9%99%90%e7%b3%bb%e7%bb%9f%28SAM%29.md">有赞权限系统(SAM).md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="有赞零售中台建设方法的探索与实践.md" href="/%e6%96%87%e7%ab%a0/%e6%9c%89%e8%b5%9e%e9%9b%b6%e5%94%ae%e4%b8%ad%e5%8f%b0%e5%bb%ba%e8%ae%be%e6%96%b9%e6%b3%95%e7%9a%84%e6%8e%a2%e7%b4%a2%e4%b8%8e%e5%ae%9e%e8%b7%b5.md">有赞零售中台建设方法的探索与实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="服务注册与发现原理剖析（Eureka、Zookeeper、Nacos）.md" href="/%e6%96%87%e7%ab%a0/%e6%9c%8d%e5%8a%a1%e6%b3%a8%e5%86%8c%e4%b8%8e%e5%8f%91%e7%8e%b0%e5%8e%9f%e7%90%86%e5%89%96%e6%9e%90%ef%bc%88Eureka%e3%80%81Zookeeper%e3%80%81Nacos%ef%bc%89.md">服务注册与发现原理剖析（Eureka、Zookeeper、Nacos）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="深入浅出Cache.md" href="/%e6%96%87%e7%ab%a0/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%baCache.md">深入浅出Cache.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="深入理解 MySQL 底层实现.md" href="/%e6%96%87%e7%ab%a0/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20MySQL%20%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0.md">深入理解 MySQL 底层实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="漫画讲解 git rebase VS git merge.md" href="/%e6%96%87%e7%ab%a0/%e6%bc%ab%e7%94%bb%e8%ae%b2%e8%a7%a3%20git%20rebase%20VS%20git%20merge.md">漫画讲解 git rebase VS git merge.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="生成浏览器唯一稳定 ID 的探索.md" href="/%e6%96%87%e7%ab%a0/%e7%94%9f%e6%88%90%e6%b5%8f%e8%a7%88%e5%99%a8%e5%94%af%e4%b8%80%e7%a8%b3%e5%ae%9a%20ID%20%e7%9a%84%e6%8e%a2%e7%b4%a2.md">生成浏览器唯一稳定 ID 的探索.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="缓存 如何保证缓存与数据库的双写一致性？.md" href="/%e6%96%87%e7%ab%a0/%e7%bc%93%e5%ad%98%20%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e7%bc%93%e5%ad%98%e4%b8%8e%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e5%8f%8c%e5%86%99%e4%b8%80%e8%87%b4%e6%80%a7%ef%bc%9f.md">缓存 如何保证缓存与数据库的双写一致性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="网易严选怎么做全链路监控的？.md" href="/%e6%96%87%e7%ab%a0/%e7%bd%91%e6%98%93%e4%b8%a5%e9%80%89%e6%80%8e%e4%b9%88%e5%81%9a%e5%85%a8%e9%93%be%e8%b7%af%e7%9b%91%e6%8e%a7%e7%9a%84%ef%bc%9f.md">网易严选怎么做全链路监控的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="美团万亿级 KV 存储架构与实践.md" href="/%e6%96%87%e7%ab%a0/%e7%be%8e%e5%9b%a2%e4%b8%87%e4%ba%bf%e7%ba%a7%20KV%20%e5%ad%98%e5%82%a8%e6%9e%b6%e6%9e%84%e4%b8%8e%e5%ae%9e%e8%b7%b5.md">美团万亿级 KV 存储架构与实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="美团点评Kubernetes集群管理实践.md" href="/%e6%96%87%e7%ab%a0/%e7%be%8e%e5%9b%a2%e7%82%b9%e8%af%84Kubernetes%e9%9b%86%e7%be%a4%e7%ae%a1%e7%90%86%e5%ae%9e%e8%b7%b5.md">美团点评Kubernetes集群管理实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="美团百亿规模API网关服务Shepherd的设计与实现.md" href="/%e6%96%87%e7%ab%a0/%e7%be%8e%e5%9b%a2%e7%99%be%e4%ba%bf%e8%a7%84%e6%a8%a1API%e7%bd%91%e5%85%b3%e6%9c%8d%e5%8a%a1Shepherd%e7%9a%84%e8%ae%be%e8%ae%a1%e4%b8%8e%e5%ae%9e%e7%8e%b0.md">美团百亿规模API网关服务Shepherd的设计与实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="解读《阿里巴巴 Java 开发手册》背后的思考.md" href="/%e6%96%87%e7%ab%a0/%e8%a7%a3%e8%af%bb%e3%80%8a%e9%98%bf%e9%87%8c%e5%b7%b4%e5%b7%b4%20Java%20%e5%bc%80%e5%8f%91%e6%89%8b%e5%86%8c%e3%80%8b%e8%83%8c%e5%90%8e%e7%9a%84%e6%80%9d%e8%80%83.md">解读《阿里巴巴 Java 开发手册》背后的思考.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="认识 MySQL 和 Redis 的数据一致性问题.md" href="/%e6%96%87%e7%ab%a0/%e8%ae%a4%e8%af%86%20MySQL%20%e5%92%8c%20Redis%20%e7%9a%84%e6%95%b0%e6%8d%ae%e4%b8%80%e8%87%b4%e6%80%a7%e9%97%ae%e9%a2%98.md">认识 MySQL 和 Redis 的数据一致性问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="进阶：Dockerfile 高阶使用指南及镜像优化.md" href="/%e6%96%87%e7%ab%a0/%e8%bf%9b%e9%98%b6%ef%bc%9aDockerfile%20%e9%ab%98%e9%98%b6%e4%bd%bf%e7%94%a8%e6%8c%87%e5%8d%97%e5%8f%8a%e9%95%9c%e5%83%8f%e4%bc%98%e5%8c%96.md">进阶：Dockerfile 高阶使用指南及镜像优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="铁总在用的高性能分布式缓存计算框架 Geode.md" href="/%e6%96%87%e7%ab%a0/%e9%93%81%e6%80%bb%e5%9c%a8%e7%94%a8%e7%9a%84%e9%ab%98%e6%80%a7%e8%83%bd%e5%88%86%e5%b8%83%e5%bc%8f%e7%bc%93%e5%ad%98%e8%ae%a1%e7%ae%97%e6%a1%86%e6%9e%b6%20Geode.md">铁总在用的高性能分布式缓存计算框架 Geode.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="阿里云PolarDB及其共享存储PolarFS技术实现分析（上）.md" href="/%e6%96%87%e7%ab%a0/%e9%98%bf%e9%87%8c%e4%ba%91PolarDB%e5%8f%8a%e5%85%b6%e5%85%b1%e4%ba%ab%e5%ad%98%e5%82%a8PolarFS%e6%8a%80%e6%9c%af%e5%ae%9e%e7%8e%b0%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89.md">阿里云PolarDB及其共享存储PolarFS技术实现分析（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="阿里云PolarDB及其共享存储PolarFS技术实现分析（下）.md" href="/%e6%96%87%e7%ab%a0/%e9%98%bf%e9%87%8c%e4%ba%91PolarDB%e5%8f%8a%e5%85%b6%e5%85%b1%e4%ba%ab%e5%ad%98%e5%82%a8PolarFS%e6%8a%80%e6%9c%af%e5%ae%9e%e7%8e%b0%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89.md">阿里云PolarDB及其共享存储PolarFS技术实现分析（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="面试最常被问的 Java 后端题.md" href="/%e6%96%87%e7%ab%a0/%e9%9d%a2%e8%af%95%e6%9c%80%e5%b8%b8%e8%a2%ab%e9%97%ae%e7%9a%84%20Java%20%e5%90%8e%e7%ab%af%e9%a2%98.md">面试最常被问的 Java 后端题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="领域驱动设计在互联网业务开发中的实践.md" href="/%e6%96%87%e7%ab%a0/%e9%a2%86%e5%9f%9f%e9%a9%b1%e5%8a%a8%e8%ae%be%e8%ae%a1%e5%9c%a8%e4%ba%92%e8%81%94%e7%bd%91%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e4%b8%ad%e7%9a%84%e5%ae%9e%e8%b7%b5.md">领域驱动设计在互联网业务开发中的实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="领域驱动设计的菱形对称架构.md" href="/%e6%96%87%e7%ab%a0/%e9%a2%86%e5%9f%9f%e9%a9%b1%e5%8a%a8%e8%ae%be%e8%ae%a1%e7%9a%84%e8%8f%b1%e5%bd%a2%e5%af%b9%e7%a7%b0%e6%9e%b6%e6%9e%84.md">领域驱动设计的菱形对称架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="高效构建 Docker 镜像的最佳实践.md" href="/%e6%96%87%e7%ab%a0/%e9%ab%98%e6%95%88%e6%9e%84%e5%bb%ba%20Docker%20%e9%95%9c%e5%83%8f%e7%9a%84%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5.md">高效构建 Docker 镜像的最佳实践.md</a>
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
                            <h1 id="title" data-id="eBay 的 Elasticsearch 性能调优实践" class="title">eBay 的 Elasticsearch 性能调优实践</h1>
                            <div><p>Elasticsearch 是一个基于 Apache Lucene 的开源搜索和分析引擎，允许用户近实时地存储、搜索和分析数据。Pronto 是 eBay 托管 Elasticsearch 集群的平台，使 eBay 内部客户易于部署、运维和扩展 Elasticsearch 以进行全文搜索、实时分析和日志事件监控。今天 Pronto 管理着 60 多个 Elasticsearch 集群，达 2000 多个节点。日采集数据量达到 180 亿个文档，日均查询量达到 35 亿。该平台提供了创建、修复、安全、监控、告警和诊断的一整套功能。</p>

<p>虽然 Elasticsearch 专为快速查询而设计，但其性能在很大程度上取决于应用程序的场景、索引的数据量以及应用程序和用户查询数据的速度。本文总结了 Pronto 团队面临的挑战以及应对挑战所构建的流程和工具，还给出了对几种配置进行基准测试的一些结果。</p>

<h2 id="挑战">挑战</h2>

<p>迄今遇到的 Pronto/Elasticsearch 使用场景所面临的挑战包括：</p>

<ol>
<li>高吞吐量：一些集群每天采集的数据高达 5TB，一些集群每天的搜索请求超过 4 亿。如果 Elasticsearch 无法及时处理这些请求，上游的请求将发生积压。</li>
<li>低搜索延迟：对于性能比较关键的集群，尤其是面向线上的系统，低搜索延迟是必需的，否则用户体验将受到影响。</li>
<li>由于数据或查询是可变的，所以最佳设置也是变化的。不存在所有情况都是最佳的设置。例如，将索引拆分成更多的分片对于费时的查询是有好处的，但是这可能会影响其他查询性能。</li>
</ol>

<h2 id="解决方案">解决方案</h2>

<p>为了帮助我们的客户应对这些挑战，Pronto 团队为用户案例上线和整个集群生命周期，针对性能测试、调优和监控构建了一套策略方法。</p>

<ol>
<li>预估集群大小：在新的用户案例上线之前，收集客户提供的信息，如吞吐量、文档大小、文档数量和搜索类型，以估计 Elasticsearch 集群的初始大小。</li>
<li>优化索引设计：与客户一起评估索引设计。</li>
<li>索引性能调优：根据用户场景进行索引性能和搜索性能调优。</li>
<li>搜索性能调优：使用用户真实数据 / 查询运行性能测试，比较和分析不同 Elasticsearch 配置参数的测试结果。</li>
<li>运行性能测试：在用户案例上线后，集群将受到监控，并且每当数据发生变化，查询更改或流量增加时，用户都可以自由地重新运行性能测试。</li>
</ol>

<h3 id="预估集群大小">预估集群大小</h3>

<p>Pronto 团队为每种类型的机器和每个支持的 Elasticsearch 版本运行基准测试，收集性能数据，然后结合客户提供的信息，估算群集初始大小，包括：</p>

<ul>
<li>索引吞吐量</li>
<li>文档大小</li>
<li>搜索吞吐量</li>
<li>查询类型</li>
<li>热点索引文档数量</li>
<li>保留策略</li>
<li>响应时间需求</li>
<li>SLA 级别</li>
</ul>

<h3 id="优化索引设计">优化索引设计</h3>

<p>在开始索引数据和运行查询之前，我们先考虑一下。索引到底表示什么？Elastic 的官方答案是“具有某种相似特征的文档集合”。因此，下一个问题是“应该使用哪些特征来对数据进行分组？应该把所有文档放入一个索引还是多个索引？”，答案是，这取决于使用的查询。以下是关于如何根据最常用的查询组织索引的一些建议。</p>

<ul>
<li><strong>如果查询有一个过滤字段并且它的值是可枚举的，那么把数据分成多个索引。</strong>例如，你有大量的全球产品信息被采集到 Elasticsearch 中，大多数查询都有一个过滤条件“地区”，并且很少有机会运行跨地区查询。如下查询体可以被优化：</li>
</ul>

<pre><code>{
    &quot;query&quot;: {
        &quot;bool&quot;: {
            &quot;must&quot;: {
                &quot;match&quot;: {
                    &quot;title&quot;: &quot;${title}&quot;
                }
            },
            &quot;filter&quot;: {
                &quot;term&quot;: {
                    &quot;region&quot;: &quot;US&quot;
                }
            }
        }
    }
}
</code></pre>

<p>在这种情况下，如果索引按照美国、欧盟等地区分成几个较小的索引，可以从查询中删除过滤子句，查询性能会更好。如果我们需要运行一个跨地区查询，我们可以将多个索引或通配符传递给 Elasticsearch。</p>

<ul>
<li><strong>如果查询有一个过滤字段并且它的值是不可枚举的，建议使用路由。</strong>通过使用过滤字段值作为路由键，我们可以将具有相同过滤字段值的文档索引至同一个分片，并移除过滤子句。</li>
</ul>

<p>例如，Elasticsearch 集群中存有数以百万记的订单数据，大多数查询都包含有买方 ID 作为限定从句。为每个买家创建索引是不可能的，所以我们不能通过买方 ID 将数据拆分成多个索引。一个合适的解决方案是，使用路由将具有相同买方 ID 的所有订单放入相同分片中。然后几乎所有的查询都可以在匹配路由键的分片内完成。</p>

<ul>
<li><strong>如果查询具有日期范围过滤子句，则按日期建立数据。</strong>这适用于大多数日志记录和监控场景。我们可以按天、周或月组织索引，然后可以获得指定的日期范围内的索引列表，这样，Elasticsearch 只需要查询一个较小的数据集而不是整个数据集。另外，当数据过期时，删除旧的索引也很容易。</li>
<li><strong>明确设置映射。</strong>虽然 Elasticsearch 可以动态创建映射，但创建的映射可能并不适用于所有场景。例如，Elasticsearch 5.x 中的默认字符串字段映射是“keyword”和“text”类型。这在很多情况下是没有必要的。</li>
<li><strong>如果文档使用用户定义的 ID 或路由进行索引，要避免造成分片不平衡。</strong> Elasticsearch 使用随机 ID 生成器和散列算法来确保文档均匀地分配给分片。当使用用户定义的 ID 或路由时，ID 或路由键可能不够随机，造成一些分片明显比其他分片大。在这种情况下，这个分片上的读 / 写操作会比其他的慢得多。我们可以优化 ID/ 路由键或使用<a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-routing-field.html#routing-index-partition" target="_blank"> index.routing_partition_size </a>（5.3 和更高版本中可用）。</li>
<li><strong>确保分片均匀分布在节点上。</strong>一个节点如果比其他节点的分片多，则会比其他节点承担更多的负载，成为整个系统的瓶颈。</li>
</ul>

<h3 id="索引性能调优">索引性能调优</h3>

<p>对于日志记录和监控等重度索引场景，索引性能是关键指标。这里有一些建议：</p>

<ul>
<li><strong>使用批量请求。</strong></li>
<li><strong>使用多线程发送请求。</strong></li>
<li><strong>增加刷新时间间隔。</strong>每次刷新事件发生时，Elasticsearch 都会创建一个新的 Lucene 段，并在稍后进行合并。增加刷新时间间隔将降低创建和合并的开销。请注意，文档只有在刷新后才能搜索到。</li>
</ul>

<p><img src="assets/d296638b00d0b40908c4ab5712aa1952.png" alt="img" /></p>

<p>性能和刷新时间间隔之间的关系</p>

<p>从上图可以看出，随着刷新时间间隔的增加，吞吐量增加，响应时间减少。我们可以使用下面的请求来检查我们有多少段以及刷新和合并花了多少时间。</p>

<p><code>Index/_stats?filter_path= indices.**.refresh,indices.**.segments,indices.**.merges</code>- <strong>减少副本数量。</strong>对于每个索引请求，Elasticsearch 需要将文档写入主分片和所有副本分片。显然，副本过多会减慢索引速度，但另一方面，这将提高搜索性能。我们将在本文后面讨论这个问题。</p>

<p><img src="assets/47d20b17cdc09959f3e1eedb03a296de.png" alt="img" />
性能和副本数之间的关系</p>

<p>从上图可以看出，随着副本数增加，吞吐量下降，响应时间增加。</p>

<ul>
<li><strong>尽可能使用自动生成的 ID。</strong> Elasticsearch 自动生成的 ID 保证是唯一的，能避免版本查找。如果客户真的需要使用自定义的 ID，我们建议选择一个对 Lucene 友好的 ID，比如零填充顺序 ID、UUID-1 或者纳秒级时间。这些 ID 具有一致的顺序模式，能良好压缩。相比之下，像 UUID-4 这样的 ID 本质上是随机的，压缩率低，会降低 Lucene 的速度。</li>
</ul>

<h3 id="搜索性能调优">搜索性能调优</h3>

<p>使用 Elasticsearch 的主要原因是支持搜索数据。用户应该能够快速找到他们正在寻找的信息。搜索性能取决于很多因素。</p>

<ul>
<li><strong>尽可能使用过滤器上下文（Filter）替代查询上下文（Query）。</strong> 查询子句用于回答“这个文档与此子句相匹配的程度”，而过滤器子句用于回答“这个文档是否匹配这个子句”，Elasticsearch 只需要回答“是”或“否”，不需要为过滤器子句计算相关性分数，而且过滤器结果可以缓存。有关详细信息，请参阅<a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/query-filter-context.html" target="_blank">查询和过滤上下文</a>。</li>
</ul>

<p><img src="assets/2be7679a1b248987f612dc3a4fd1110b.png" alt="img" />
查询和过滤器性能比较</p>

<ul>
<li><strong>增加刷新时间间隔。</strong>正如我们在<a href="https://www.ebayinc.com/stories/blogs/tech/elasticsearch-performance-tuning-practice-at-ebay/#tune_indexing_performance" target="_blank">索引性能调优</a>中所提到的，Elasticsearch 每次刷新时都会创建一个新的段。增加刷新时间间隔将有助于减少段数并降低搜索的 IO 成本。并且，一旦发生刷新并且数据改变，缓存将会失效。增加刷新时间间隔可以使 Elasticsearch 更高效地利用缓存。</li>
<li><strong>增加副本数。</strong> Elasticsearch 可以在主分片或副本分片上执行搜索。副本越多，搜索可用的节点就越多。</li>
</ul>

<p><img src="assets/40d1a26c967a0f4c37f612f1f8e9ca9a.png" alt="img" />
搜索性能和副本数之间的关系</p>

<p>从上图可以看出，搜索吞吐量几乎与副本数量成线性关系。注意在这个测试中，测试集群有足够的数据节点来确保每个分片都有一个专有节点。如果这个条件不能满足，搜索吞吐量就不会这么好。</p>

<ul>
<li><strong>尝试不同的分片数。</strong>“应该为索引设置多少分片呢？”这可能是最常见的问题。遗憾的是，没有适合所有应用场景的分片数。这完全取决于你的情况。</li>
</ul>

<p>分片太少会使搜索无法扩展。例如，如果分片数设置为 1，则索引中的所有文档都将存储在一个分片中。对于每个搜索，只有一个节点能够参与计算。如果索引中的文件数量很多，查询会很耗时。从另一方面来说，创建的索引分片太多也会对性能造成不利影响，因为 Elasticsearch 需要在所有分片上运行查询（除非在请求中指定了路由键），然后提取并合并所有返回的结果。</p>

<p>根据我们的经验，如果索引小于 1G，可以将分片数设置为 1。对于大多数场景，我们可以将分片数保留为默认值 5，但是如果分片大小超过 30GB，我们应该增加分片 ，将索引分成更多的分片。创建索引后，分片数不能更改，但是我们可以创建新的索引并使用 reindex API 迁移数据。</p>

<p>我们测试了一个拥有 1 亿个文档，大约 150GB 的索引。我们使用了 100 个线程发送搜索请求。</p>

<p><img src="assets/4f035bac3fe936ed084ed3352592bf10.png" alt="img" />
搜索性能和分片数量之间的关系</p>

<p>从上图可以看出，最优的分片数量为 11 个。开始时搜索吞吐量增大（响应时间减少），但随着分片数量的增加，搜索吞吐量减小（响应时间增加）。</p>

<p>请注意，在这个测试中，就像在副本数测试中一样，每个分片都有一个独占节点。如果这个条件不能满足，搜索吞吐量将不会像这个图那样好。</p>

<p>在这种情况下，我们建议你尝试一个小于最优值的分片数，因为如果分片多，并且使每个分片都有一个独占数据节点，那么就需要很多节点。</p>

<ul>
<li><strong>节点查询缓存。</strong><a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/query-cache.html" target="_blank">节点查询缓存</a>只缓存过滤器上下文中使用的查询。与查询子句不同，过滤子句是“是”或“否”的问题。Elasticsearch 使用位集（bit set）机制来缓存过滤结果，以便后面使用相同的过滤器的查询进行加速。请注意，Elasticsearch 只对保存超过 10,000（或文档总数的 3％，以较大者为准）个文档的段启用查询缓存。有关更多详细信息，请参阅<a href="https://www.elastic.co/guide/en/elasticsearch/guide/current/filter-caching.html#_independent_query_caching" target="_blank">缓存</a>。</li>
</ul>

<p>我们可以使用下面的请求来检查一个节点查询缓存是否生效。</p>

<pre><code>GET index_name/_stats?filter_path=indices.**.query_cache
{
  &quot;indices&quot;: {
    &quot;index_name&quot;: {
      &quot;primaries&quot;: {
        &quot;query_cache&quot;: {
          &quot;memory_size_in_bytes&quot;: 46004616,
          &quot;total_count&quot;: 1588886,
          &quot;hit_count&quot;: 515001,
          &quot;miss_count&quot;: 1073885,
          &quot;cache_size&quot;: 630,
          &quot;cache_count&quot;: 630,
          &quot;evictions&quot;: 0
        }
      },
      &quot;total&quot;: {
        &quot;query_cache&quot;: {
          &quot;memory_size_in_bytes&quot;: 46004616,
          &quot;total_count&quot;: 1588886,
          &quot;hit_count&quot;: 515001,
          &quot;miss_count&quot;: 1073885,
          &quot;cache_size&quot;: 630,
          &quot;cache_count&quot;: 630,
          &quot;evictions&quot;: 0
        }
      }
    }
  }
}
</code></pre>

<ul>
<li>分片查询缓存。</li>
</ul>

<p>如果大多数查询是聚合查询，我们应该考虑</p>

<p>分片查询缓存</p>

<p>。分片查询缓存可以缓存聚合结果，以便 Elasticsearch 以低开销直接处理请求。有几件事情需要注意：</p>

<ul>
<li>设置“size”为 0。分片查询缓存只缓存聚合结果和建议。它不会缓存命中，因此如果将 size 设置为非零，则无法从缓存中获益。</li>
<li>查询请求的负载（Payload）必须完全相同。分片查询缓存使用请求负载作为缓存键，因此需要确保后续查询请求的负载必须和之前的完全一致。由于负载中 JSON 键的顺序变化会导致负载变化，故建议对负载的键进行排序来确保顺序一致。</li>
<li>处理好日期时间。不要直接在查询中使用像 Date.now 这样的变量。否则，每个请求的请求体都不同，从而导致缓存始终无效。我们建议将日期时间整理为小时或天，以便更有效地利用缓存。</li>
</ul>

<p>我们可以使用下面的请求来检查分片查询缓存是否有效。</p>

<pre><code>GET index_name/_stats?filter_path=indices.**.request_cache
{
  &quot;indices&quot;: {
    &quot;index_name&quot;: {
      &quot;primaries&quot;: {
        &quot;request_cache&quot;: {
          &quot;memory_size_in_bytes&quot;: 0,
          &quot;evictions&quot;: 0,
          &quot;hit_count&quot;: 541,
          &quot;miss_count&quot;: 514098
        }
      },
      &quot;total&quot;: {
        &quot;request_cache&quot;: {
          &quot;memory_size_in_bytes&quot;: 0,
          &quot;evictions&quot;: 0,
          &quot;hit_count&quot;: 982,
          &quot;miss_count&quot;: 947321
        }
      }
    }
  }
}
</code></pre>

<ul>
<li><strong>仅检索必要的字段。</strong>如果文档很大，而你只需要几个字段，请使用<a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-stored-fields.html" target="_blank"> stored_fields </a>检索需要的字段而不是所有字段。</li>
<li><strong>避免搜索停用词。</strong>诸如“a”和“the”等停用词可能导致查询命中结果数暴增。假设你有一百万个文档。搜索“fox”可能会返回几十个命中文档，但搜索“the fox”可能会返回索引中的所有文档，因为“the”几乎出现在所有文档中。Elasticsearch 需要对所有命中的结果进行评分和排序，以致像“the fox”这样的查询会减慢整个系统。你可以使用停用词过滤器来删除停用词，或使用“and”运算符将查询更改为“the AND fox”，获得更精确的结果。</li>
</ul>

<p>如果某些单词在索引中经常使用，但不在默认停用词列表中，则可以使用<a href="https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-cutoff" target="_blank">截止频率</a>来动态处理它们。</p>

<ul>
<li><strong>如果不关心文档返回的顺序，则按 _doc 排序。</strong> Elasticsearch 默认使用“_score”字段按评分排序。如果不关心顺序，可以使用&rdquo;sort&rdquo;:&ldquo;_doc&rdquo;让 Elasticsearch 按索引顺序返回命中文档，可以节省排序开销。</li>
<li><strong>避免使用脚本查询（script query）计算动态字段，建议在索引时计算并在文档中添加该字段。</strong>例如，我们有一个包含大量用户信息的索引，我们需要查询以&rdquo;1234&rdquo;开头的所有用户。你可能运行一个脚本查询，如&rdquo;source&rdquo;:“doc[‘num’].value.startsWith（‘1234’）”。这个查询非常耗费资源，并且减慢整个系统。索引时考虑添加一个名为“num_prefix”的字段。然后我们可以查询&rdquo;name_prefix&rdquo;:“1234”。</li>
<li><strong>避免使用通配符查询。</strong></li>
</ul>

<h2 id="运行性能测试">运行性能测试</h2>

<p>对于每一次变更，都需要运行性能测试来验证变更是否适用。因为 Elasticsearch 是一个 RESTful 服务，所以可以使用 Rally、Apache Jmeter 和 Gatling 等工具来运行性能测试。因为 Pronto 团队需要在每种类型的机器和 Elasticsearch 版本上运行大量的基准测试，而且需要在许多 Elasticsearch 集群上针对不同 Elasticsearch 配置参数运行性能测试，所以这些工具不能满足我们的要求。</p>

<p>Pronto 团队建立了基于<a href="https://gatling.io/" target="_blank"> Gatling </a>的在线性能分析服务，帮助客户和我们运行性能测试和回归测试。该服务提供的功能有：</p>

<ol>
<li>轻松添加和编辑测试。用户无需 Gatling 或 Scala 知识即可根据输入的查询或文档结构生成测试。</li>
<li>顺序运行多个测试，无需人工干预。该服务可以检查状态并在每次测试之前 / 之后更改 Elasticsearch 设置。</li>
<li>帮助用户比较和分析测试结果。测试期间的测试结果和集群统计信息将保留下来，并可以通过预定义的 Kibana 可视化进行分析。</li>
<li>从命令行或 Web UI 运行测试。该服务提供了与其他系统集成的 Rest API。</li>
</ol>

<p>架构如下：</p>

<p>性能测试服务架构</p>

<p>用户可以查看每个测试的 Gatling 报告，并查看 Kibana 预定义的可视化图像，以便进一步分析和比较，如下所示。</p>

<p><img src="assets/34796429d29e25201e8e30178fc675d7.png" alt="img" />
Gatling 报告</p>

<p><img src="assets/9b1f27a4e022ca694527f129f1e7a946.png" alt="img" />
Gatling 报告</p>

<h2 id="总结">总结</h2>

<p>本文总结了在设计满足高期望的采集和搜索性能的 Elasticsearch 集群时应该考虑的索引 / 分片 / 副本设计以及一些其他配置，还说明了 Pronto 如何在策略上帮助客户进行初始规模调整、索引设计和调优以及性能测试。截至今天，Pronto 团队已经帮助包括订单管理系统（OMS）和搜索引擎优化（SEO）在内的众多客户实现了苛刻的性能目标，从而为 eBay 的关键业务作出了贡献。</p>

<p>Elasticsearch 的性能取决于很多因素，包括文档结构、文档大小、索引设置 / 映射、请求率、数据集大小和查询命中次数等等。针对一种情况的建议不一定适用于另一种情况，因此，彻底进行性能测试、收集数据、根据负载调整配置以及优化集群以满足性能要求非常重要。</p>

<p><strong>查看英文原文：</strong><a href="https://www.ebayinc.com/stories/blogs/tech/elasticsearch-performance-tuning-practice-at-ebay/" target="_blank"> https://www.ebayinc.com/stories/blogs/tech/elasticsearch-performance-tuning-practice-at-ebay/</a></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#adc1c1c194999c9c9d9aedcac0ccc4c183cec2c0" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0c72856cefede4',t:'MTczMzk5MzIyMS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
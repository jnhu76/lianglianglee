<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=阿里云PolarDB及其共享存储PolarFS技术实现分析（上）>
        <link rel="icon" href="/static/favicon.png">
        <title>阿里云PolarDB及其共享存储PolarFS技术实现分析（上） </title>
        
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
                            <h1 id="title" data-id="阿里云PolarDB及其共享存储PolarFS技术实现分析（上）" class="title">阿里云PolarDB及其共享存储PolarFS技术实现分析（上）</h1>
                            <div><p>PolarDB是阿里云基于MySQL推出的云原生数据库（Cloud Native Database）产品，通过将数据库中计算和存储分离，多个计算节点访问同一份存储数据的方式来解决目前MySQL数据库存在的运维和扩展性问题；通过引入RDMA和SPDK等新硬件来改造传统的网络和IO协议栈来极大提升数据库性能。代表了未来数据库发展的一个方向。本系列共2篇文章，主要分析为什么会出现PolarDB以及其技术实现。</p>

<p>由于PolarDB并不开源，因此只能基于阿里云公开的技术资料进行解读。这些资料包括从去年下半年开始陆续在阿里云栖社区、云栖大会等场合发布的PolarDB相关资料，以及今年以来公开的PolarDB后端共享存储PolarFS相关文章。</p>

<p>PolarDB出现背景</p>

<p>MySQL云服务遇到的问题</p>

<p>首先来了解下为什么会出现PolarDB。阿里云数据库团队具备国内领先的技术能力，为MySQL等数据库在国内的推广起到了很大的作用。在阿里云上也维护了非常庞大的MySQL云服务（RDS）集群，但也遇到了很多棘手的问题。举例如下：</p>

<ul>
<li>实例数据量太大，单实例几个TB的数据，这样即使使用xtrabackup物理备份，也需要很长的备份时间，且备份期间写入量大的话可能导致redo日志被覆盖引起备份失败；</li>
<li>大实例故障恢复需要重建时，耗时太长，影响服务可用性（此时存活节点也挂了，那么完蛋了）。时间长有2个原因，一是备份需要很长时间，二是恢复的时候回放redo也需要较长时间；</li>
<li>大实例做只读扩展麻烦，因为只读实例的数据是单独一份的，所以也需要通过备份来重建；</li>
<li>RDS实例集群很大，包括成千上万个实例，可能同时有很多实例同时在备份，会占用云服务巨大的网络和IO带宽，导致云服务不稳定；</li>
<li>云服务一般使用云硬盘，导致数据库的性能没有物理机实例好，比如IO延时过高；</li>
<li>主库写入量大的时候，会导致主从复制延迟过大，semi-sync/半同步复制也没法彻底解决，这是由于mysql基于binlog复制，需要走完整的mysql事务处理流程。</li>
<li>对于需要读写分离，且要求部署多个只读节点的用户，最明显的感觉就是每增加一个只读实例，成本是线性增长的。</li>
</ul>

<p>其实不仅仅是阿里云RDS，网易云上的RDS服务也有数千个实例，同样遇到了类似的问题，我们是亲身经历而非感同身受。应该说就目前的MySQL技术实现方案，要解决上述任何一个问题都不是件容易的事情，甚至有几个问题是无法避免的。</p>

<p>现有解决方案及不足</p>

<p>那么，跳出MySQL，是否有解决方案呢，分析目前业界的数据库和存储领域技术，可以发现基于共享存储是个可选的方案，所谓数据库共享存储方案指的是RDS实例（一般指一主一从的高可用实例）和只读实例共享同一份数据，这样在实例故障或只读扩展时就无需拷贝数据了，只需简单得把故障节点重新拉起来，或者新建个只读计算节点即可，省时省力更省钱。共享存储可通过快照技术（snapshot/checkpoint）和写时拷贝（copy-on-write，COW）来解决数据备份和误操作恢复问题，将所需备份的数据量分摊到较长的一段时间内，而不需要瞬时完成，这样就不会导致多实例同时备份导致网络和IO数据风暴。下图就是一个典型的数据库共享存储方案，Primary节点即数据库主节点，对外提供读写服务，Read Only节点可以是Primary的灾备节点，也可以是对外提供只读服务的节点，他们共享一份底层数据。</p>

<p><img src="assets/20181012175033a45a1419-5df7-4ec2-9601-43ece82896b8.jpg" alt="img" /></p>

<p>理想很丰满，但现实却很骨感，目前可用的共享存储方案寥寥无几，比如在Hadoop生态圈占统治地位的HDFS，以及在通用存储领域风生水起的Ceph，只是如果将其作为在线数据处理（OLTP）服务的共享存储，最终对用户呈现的性能是不可接受的。除此之外，还存在大量与现有数据库实现整合适配等问题。</p>

<p>PolarDB实现方案</p>

<p>云原生数据库</p>

<p>说道云原生数据库，就不得不提Aurora。其在2014年下半年发布后，轰动了整个数据库领域。Aurora对MySQL存储层进行了大刀阔斧的改造，将其拆为独立的存储节点(主要做数据块存储，数据库快照的服务器)。上层的MySQL计算节点(主要做SQL解析以及存储引擎计算的服务器)共享同一个存储节点，可在同一个共享存储上快速部署新的计算节点，高效解决服务能力扩展和服务高可用问题。基于日志即数据的思想，大大减少了计算节点和存储节点间的网络IO，进一步提升了数据库的性能。再利用存储领域成熟的快照技术，解决数据库数据备份问题。被公认为关系型数据库的未来发展方向之一。截止2018年上半年，Aurora已经实现了多个计算节点同时提供写服务的能力，继续在云原生数据库上保持领先的地位。</p>

<p>不难推断，在Aurora发布3年后推出的PolarDB，肯定对Aurora进行了深入的研究，并借鉴了很多技术实现方法。关于Aurora的分析，国内外，包括公司内都已进行了深入分析，本文不再展开描述。下面着重介绍PolarDB实现。我们采用先存储后计算的方式，先讲清楚PolarFS共享存储的实现，再分析PolarDB计算层如何适配PolarFS。</p>

<p>PolarDB架构</p>

<p><img src="assets/201810121750331dd25433-6361-4349-9069-a7102623b6eb.png" alt="img" /></p>

<p>上图为PolarFS视角看到的PolarDB实现架构。一套PolarDB至少包括3个部分，分别为最底层的共享存储，与用户交互的MySQL节点，还有用户进行系统管理的PolarCtrl。而其中PolarFS又可进一步拆分为libpfs、PolarSwitch和ChunkServer。下面进行简单说明：</p>

<ul>
<li>MySQL节点，即图中的POLARDB，负责用户SQL解析、事务处理等数据库相关操作，扮演计算节点角色；</li>
<li>libpfs是一个用户空间文件系统库，提供POSIX兼容的文件操作API接口，嵌入到PolarDB负责数据库IO（File IO）接入；</li>
<li>PolarSwitch运行在计算节点主机（Host）上，每个Host部署一个PolarSwitch的守护进程，其将数据库文件IO变换为块设备IO，并发送到具体的后端节点（即ChunkServer）；</li>
<li>ChunkServer部署在存储节点上，用于处理块设备IO（Block IO）请求和节点内的存储资源分布；</li>
<li>PolarCtrl是系统的控制平面，PolarFS集群的控制核心，所有的计算和存储节点均部署有PolarCtrl的Agent。</li>
</ul>

<p>PolarFS的存储组织</p>

<p>与大多数存储系统一样，PolarFS对存储资源也进行了多层封装和管理，PolarFS的存储层次包括：Volume、Chunk和Block，分别对应存储领域中的数据卷，数据区和数据块，在有些系统中Chunk又被成为Extent，均表示一段连续的块组成的更大的区域，作为分配的基本单位。一张图可以大致表现各层的关系：</p>

<p><img src="assets/20181012175033932353c3-f60e-4824-82ab-fb726c962c6a.png" alt="img" /></p>

<p><strong>Volume</strong></p>

<p>当用户申请创建PolarDB数据库实例时，系统就会为该实例创建一个Volume（卷，本文后续将这两种表达混用），每个卷都有多个Chunk组成，其大小就是用户指定的数据库实例大小，PolarDB支持用户创建的实例大小范围是10GB至100TB，满足绝大部分云数据库实例的容量要求。</p>

<p>跟其他传统的块设备一样，卷上的读写IO以512B大小对齐，对卷上同个Chunk的修改操作是原子的。当然，卷还是块设备层面的概念，在提供给数据库实例使用前，需在卷上格式化一个PolarFS文件系统（PFS）实例，跟ext4、btrfs一样，PFS上也会在卷上存放文件系统元数据。这些元数据包括inode、directory entry和空闲块等对象。同时，PFS也是一个日志文件系统，为了实现文件系统的元数据一致性，元数据的更新会首先记录在卷上的Journal（日志）文件中，然后才更新指定的元数据。</p>

<p>跟传统文件系统不一样的是PolarFS是个共享文件系统即一个卷会被挂载到多个计算节点上，也就是说可能存在有多个客户端（挂载点）对文件系统进行读写和更新操作，所以PolarFS在卷上额外维护了一个Paxos文件。每个客户端在更新Journal文件前，都需要使用Paxos文件执行Disk Paxos算法实现对Journal文件的互斥访问。更详细的PolarFS元数据更新实现，后续单独作为一个小节。</p>

<p><strong>Chunk</strong></p>

<p>前面提到，每个卷内部会被划分为多个Chunk（区），区是数据分布的最小粒度，每个区都位于单块SSD盘上，其目的是利于数据高可靠和高可用的管理，详见后续章节。每个Chunk大小设置为10GB，远大于其他类似的存储系统，例如GFS为64MB，Linux LVM的物理区（PE）为4MB。这样做的目的是减少卷到区映射的元数据量大小（例如，100TB的卷只包含10K个映射项）。一方面，全局元数据的存放和管理会更容易；另一方面，元数据可以全都缓存在内存中，避免关键IO路径上的额外元数据访问开销。</p>

<p>当然，Chunk设置为10GB也有不足。当上层数据库应用出现区域级热点访问时，Chunk内热点无法进一步打散，但是由于每个存储节点提供的Chunk数量往往远大于节点数量（节点:Chunk在1:1000量级），PolarFS支持Chunk的在线迁移，其上服务着大量数据库实例，因此可以将热点Chunk分布到不同节点上以获得整体的负载均衡。</p>

<p>在PolarFS上，卷上的每个Chunk都有3个副本，分布在不同的ChunkServer上，3个副本基于ParallelRaft分布式一致性协议来保证数据高可靠和高可用。</p>

<p><strong>Block</strong></p>

<p>在ChunkServer内，Chunk会被进一步划分为163,840个Block（块），每个块大小为64KB。Chunk至Block的映射信息由ChunkServer自行管理和保存。每个Chunk除了用于存放数据库数据的Block外，还包含一些额外Block用来实现预写日志（Write Ahead Log，WAL）。</p>

<p>需要注意的是，虽然Chunk被进一步划分为块，但Chunk内的各个Block在SSD盘是物理连续的。PolarFS的VLDB文章里提到“Blocks are allocated and mapped to a chunk on demand to achieve thin provisioning”。thin provisioning就是精简配置，是存储上常用的技术，就是用户创建一个100GB大小的卷，但其实在卷创建时并没有实际分配100GB存储空间给它，仅仅是逻辑上为其创建10个Chunk，随着用户数据不断写入，PolarFS不断分配物理存储空间供其使用，这样能够实现存储系统按需扩容，大大节省存储成本。</p>

<p>那么为何PolarFS要引入Block这个概念呢，其中一个是跟卷上的具体文件相关，我们知道一个文件系统会有多个文件，比如InnoDB数据文件*.ibd。每个文件大小会动态增长，文件系统采用预分配（fallocate()）为文件提前分配更多的空间，这样在真正写数据的时无需进行文件系统元数据操作，进而优化了性能。显然，每次给文件分配一个Chunk，即10GB空间是不合理的，64KB或其倍数才是合适的值。上面提到了精简配置和预分配，看起来是冲突的方法，但其实是统一的，精简配置的粒度比预分配的粒度大，比如精简配置了10GB，预分配了64KB。这样对用户使用没有任何影响，同时还节省了存储成本。</p>

<p>PolarFS组件解析</p>

<p>首先展示一张能够更加清晰描述与数据流相关的各个组件作用的示意图，并逐一对其进行解释。</p>

<p><img src="assets/20181012175034d5fb0058-7b69-43ba-a25b-19a75f533ef0.png" alt="img" /></p>

<p><strong>libpfs</strong></p>

<p>libpfs是一个用户空间文件系统（即上图User Space File System）库，负责数据库IO（File IO）接入。更直观点，libpfs提供了供计算节点/PolarDB访问底层存储的API接口，进行文件读写和元数据更新等操作，如下图所示：</p>

<p><img src="assets/20181012175034e664ea5a-6091-4efd-a3c6-dbb93b49cfe3.png" alt="img" /></p>

<p>pfs_mount()用于将指定卷上文件系统挂载到对应的数据库计算节点上，该操作会获取卷上的文件系统元数据信息，将其缓存在计算节点上，这些元数据信息包括目录树（the directory tree），文件映射表（the file mapping table）和块映射表（the block mapping table）等，其中目录树描述了文件目录层级结构信息，每个文件名对应的inode节点信息（目录项）。inode节点信息就是文件系统中唯一标识一个文件的FileID。文件映射表描述了该文件都有哪些Block组成。通过上图我们还发现了pfs_mount_growfs()，该API可以让用户方便得进行数据库扩容，在对卷进行扩容后，通过调用该API将增加的空间映射到文件系统层。</p>

<p><img src="assets/20181012175035369fb3c9-464d-488f-8264-e62e38d2e26f.png" alt="img" /></p>

<p>上图右侧的表描述了目录树中的某个文件的前3个块分别对应的是卷的第348,1500和201这几个块。假如数据库操作需要回刷一个脏页，该页在该表所属文件的偏移位置128KB处，也就是说要写该文件偏移128KB开始的16KB数据，通过文件映射表知道该写操作其实写的是卷的第201个块。这就是lipfs发送给PolarSwitch的请求包含的内容：volumeid，offset和len。其中offset就是201*64KB，len就是16KB。</p>

<p><strong>PolarSwitch</strong></p>

<p>PolarSwitch是部署在计算节点的Daemon，即上图的Data Router&amp;Cache模块，它负责接收libpfs发送而来的文件IO请求，PolarSwitch将其划分为对应的一到多个Chunk，并将请求发往Chunk所属的ChunkServer完成访问。具体来说PolarSwitch根据自己缓存的volumeid到Chunk的映射表，知道该文件请求属于那个Chunk。请求如果跨Chunk的话，会将其进一步拆分为多个块IO请求。PolarSwitch还缓存了该Chunk的三个副本分别属于那几个ChunkServer以及哪个ChunkServer是当前的Leader节点。PolarSwitch只将请求发送给Leader节点。</p>

<p><strong>ChunkServer</strong></p>

<p>ChunkServer部署在存储节点上，即上图的Data Chunk Server，用于处理块IO（Block IO）请求和节点内的存储资源分布。一个存储节点可以有多个ChunkServer，每个ChunkServer绑定到一个CPU核，并管理一块独立的NVMe SSD盘，因此ChunkServer之间没有资源竞争。</p>

<p>ChunkServer负责存储Chunk和提供Chunk上的IO随机访问。每个Chunk都包括一个WAL，对Chunk的修改会先写Log再执行修改操作，保证数据的原子性和持久性。ChunkServer使用了3D XPoint SSD和普通NVMe SSD混合型WAL buffer，Log会优先存放到更快的3DXPoint SSD中。</p>

<p>前面提到Chunk有3副本，这三个副本基于ParallelRaft协议，作为该Chunk Leader的ChunkServer会将块IO请求发送给Follow节点其他ChunkServer）上，通过ParallelRaft一致性协议来保证已提交的Chunk数据不丢失。</p>

<p><strong>PolarCtrl</strong></p>

<p>PolarCtrl是系统的控制平面，相应地Agent代理被部署到所有的计算和存储节点上，PolarCtrl与各个节点的交互通过Agent进行。PolarCtrl是PolarFS集群的控制核心，后端使用一个关系数据库云服务来管理PolarDB的元数据。其主要职责包括：</p>

<ul>
<li>监控ChunkServer的健康状况，包括剔除出现故障的ChunkServer，维护Chunk多个副本的关系，迁移负载过高的ChunkServer上的部分Chunk等；</li>
<li>Volume创建及Chunk的布局管理，比如Volume上的Chunk应该分配到哪些ChunkServer上；</li>
<li>Volume至Chunk的元数据信息维护；</li>
<li>向PolarSwitch推送元信息缓存更新，比如因为计算节点执行DDL导致卷上文件系统元数据更新，这些更新可通过PolarCtrl推送给PolarSwitch；</li>
<li>监控Volume和Chunk的IO性能，根据一定的规则进行迁移操作；</li>
<li>周期性地发起副本内和副本间的CRC数据校验。</li>
</ul>

<p>本篇主要是介绍了PolarDB数据库及其后端共享存储PolarFS系统的基本架构和组成模块，是最基础的部分。下一篇重点分析PolarFS的数据IO流程，元数据更新流程，以及PolarDB数据库节点如何适配PolarFS这样的共享存储系统。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b9d5d5d5808d8888898ef9ded4d8d0d597dad6d4" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0c92b8089ebed2',t:'MTczMzk5NDU0MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
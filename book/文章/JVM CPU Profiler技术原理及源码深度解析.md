<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=JVM&#32;CPU&#32;Profiler技术原理及源码深度解析>
        <link rel="icon" href="/static/favicon.png">
        <title>JVM CPU Profiler技术原理及源码深度解析 </title>
        
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
                            <h1 id="title" data-id="JVM CPU Profiler技术原理及源码深度解析" class="title">JVM CPU Profiler技术原理及源码深度解析</h1>
                            <div><p>研发人员在遇到线上报警或需要优化系统性能时，常常需要分析程序运行行为和性能瓶颈。Profiling技术是一种在应用运行时收集程序相关信息的动态分析手段，常用的JVM Profiler可以从多个方面对程序进行动态分析，如CPU、Memory、Thread、Classes、GC等，其中CPU Profiling的应用最为广泛。CPU Profiling经常被用于分析代码的执行热点，如“哪个方法占用CPU的执行时间最长”、“每个方法占用CPU的比例是多少”等等，通过CPU Profiling得到上述相关信息后，研发人员就可以轻松针对热点瓶颈进行分析和性能优化，进而突破性能瓶颈，大幅提升系统的吞吐量。</p>

<p>本文介绍了JVM平台上CPU Profiler的实现原理，希望能帮助读者在使用类似工具的同时也能清楚其内部的技术实现。</p>

<h2 id="cpu-profiler简介">CPU Profiler简介</h2>

<p>社区实现的JVM Profiler很多，比如已经商用且功能强大的<a href="https://www.ej-technologies.com/products/jprofiler/overview.html" target="_blank">JProfiler</a>，也有免费开源的产品，如<a href="https://github.com/uber-common/jvm-profiler" target="_blank">JVM-Profiler</a>，功能各有所长。我们日常使用的Intellij IDEA最新版内部也集成了一个简单好用的Profiler，详细的介绍参见<a href="https://blog.jetbrains.com/idea/2018/09/intellij-idea-2018-3-eap-git-submodules-jvm-profiler-macos-and-linux-and-more/" target="_blank">官方Blog</a>。</p>

<p>在用IDEA打开需要诊断的Java项目后，在“Preferences -&gt; Build, Execution, Deployment -&gt; Java Profiler”界面添加一个“CPU Profiler”，然后回到项目，单击右上角的“Run with Profiler”启动项目并开始CPU Profiling过程。一定时间后（推荐5min），在Profiler界面点击“Stop Profiling and Show Results”，即可看到Profiling的结果，包含火焰图和调用树，如下图所示：</p>

<p><img src="assets/80cac68ffeaf0064ca261d5acf285353439115.png" alt="Intellij IDEA - 性能火焰图" /></p>

<p>Intellij IDEA - 性能火焰图</p>

<p><img src="assets/d212c393113d821841023d66c50cb8b8710861.png" alt="Intellij IDEA - 调用堆栈树" /></p>

<p>Intellij IDEA - 调用堆栈树</p>

<p>火焰图是根据调用栈的样本集生成的可视化性能分析图，《<a href="http://www.ruanyifeng.com/blog/2017/09/flame-graph.html" target="_blank">如何读懂火焰图？</a>》一文对火焰图进行了不错的讲解，大家可以参考一下。简而言之，看火焰图时我们需要关注“平顶”，因为那里就是我们程序的CPU热点。调用树是另一种可视化分析的手段，与火焰图一样，也是根据同一份样本集而生成，按需选择即可。</p>

<p>这里要说明一下，因为我们没有在项目中引入任何依赖，仅仅是“Run with Profiler”，Profiler就能获取我们程序运行时的信息。这个功能其实是通过JVM Agent实现的，为了更好地帮助大家系统性的了解它，我们在这里先对JVM Agent做个简单的介绍。</p>

<h2 id="jvm-agent简介">JVM Agent简介</h2>

<p>JVM Agent是一个按一定规则编写的特殊程序库，可以在启动阶段通过命令行参数传递给JVM，作为一个伴生库与目标JVM运行在同一个进程中。在Agent中可以通过固定的接口获取JVM进程内的相关信息。Agent既可以是用C/C++/Rust编写的JVMTI Agent，也可以是用Java编写的Java Agent。</p>

<p>执行Java命令，我们可以看到Agent相关的命令行参数：</p>

<pre><code>Plain Text
    -agentlib:&lt;库名&gt;[=&lt;选项&gt;]
                  加载本机代理库 &lt;库名&gt;, 例如 -agentlib:jdwp
                  另请参阅 -agentlib:jdwp=help
    -agentpath:&lt;路径名&gt;[=&lt;选项&gt;]
                  按完整路径名加载本机代理库
    -javaagent:&lt;jar 路径&gt;[=&lt;选项&gt;]
                  加载 Java 编程语言代理, 请参阅 java.lang.instrument
</code></pre>

<h3 id="jvmti-agent">JVMTI Agent</h3>

<p>JVMTI（JVM Tool Interface）是JVM提供的一套标准的C/C++编程接口，是实现Debugger、Profiler、Monitor、Thread Analyser等工具的统一基础，在主流Java虚拟机中都有实现。</p>

<p>当我们要基于JVMTI实现一个Agent时，需要实现如下入口函数：</p>

<pre><code>// $JAVA_HOME/include/jvmti.h

JNIEXPORT jint JNICALL Agent_OnLoad(JavaVM *vm, char *options, void *reserved);
</code></pre>

<p>使用C/C++实现该函数，并将代码编译为动态连接库（Linux上是.so），通过-agentpath参数将库的完整路径传递给Java进程，JVM就会在启动阶段的合适时机执行该函数。在函数内部，我们可以通过JavaVM指针参数拿到JNI和JVMTI的函数指针表，这样我们就拥有了与JVM进行各种复杂交互的能力。</p>

<p>更多JVMTI相关的细节可以参考<a href="https://docs.oracle.com/en/java/javase/12/docs/specs/jvmti.html" target="_blank">官方文档</a>。</p>

<h3 id="java-agent">Java Agent</h3>

<p>在很多场景下，我们没有必要必须使用C/C++来开发JVMTI Agent，因为成本高且不易维护。JVM自身基于JVMTI封装了一套Java的Instrument API接口，允许使用Java语言开发Java Agent（只是一个jar包），大大降低了Agent的开发成本。社区开源的产品如<a href="https://github.com/oldmanpushcart/greys-anatomy" target="_blank">Greys</a>、<a href="https://github.com/alibaba/arthas" target="_blank">Arthas</a>、<a href="https://github.com/alibaba/jvm-sandbox" target="_blank">JVM-Sandbox</a>、<a href="https://github.com/uber-common/jvm-profiler" target="_blank">JVM-Profiler</a>等都是纯Java编写的，也是以Java Agent形式来运行。</p>

<p>在Java Agent中，我们需要在jar包的MANIFEST.MF中将Premain-Class指定为一个入口类，并在该入口类中实现如下方法：</p>

<pre><code>public static void premain(String args, Instrumentation ins) {
    // implement
}
</code></pre>

<p>这样打包出来的jar就是一个Java Agent，可以通过-javaagent参数将jar传递给Java进程伴随启动，JVM同样会在启动阶段的合适时机执行该方法。</p>

<p>在该方法内部，参数Instrumentation接口提供了Retransform Classes的能力，我们利用该接口就可以对宿主进程的Class进行修改，实现方法耗时统计、故障注入、Trace等功能。Instrumentation接口提供的能力较为单一，仅与Class字节码操作相关，但由于我们现在已经处于宿主进程环境内，就可以利用JMX直接获取宿主进程的内存、线程、锁等信息。无论是Instrument API还是JMX，它们内部仍是统一基于JVMTI来实现。</p>

<p>更多Instrument API相关的细节可以参考<a href="https://docs.oracle.com/en/java/javase/12/docs/api/java.instrument/java/lang/instrument/package-summary.html" target="_blank">官方文档</a>。</p>

<h2 id="cpu-profiler原理解析">CPU Profiler原理解析</h2>

<p>在了解完Profiler如何以Agent的形式执行后，我们可以开始尝试构造一个简单的CPU Profiler。但在此之前，还有必要了解下CPU Profiling技术的两种实现方式及其区别。</p>

<h3 id="sampling-vs-instrumentation">Sampling vs Instrumentation</h3>

<p>使用过JProfiler的同学应该都知道，JProfiler的CPU Profiling功能提供了两种方式选项: Sampling和Instrumentation，它们也是实现CPU Profiler的两种手段。</p>

<p>Sampling方式顾名思义，基于对StackTrace的“采样”进行实现，核心原理如下：</p>

<ol>
<li>引入Profiler依赖，或直接利用Agent技术注入目标JVM进程并启动Profiler。</li>
<li>启动一个采样定时器，以固定的采样频率每隔一段时间（毫秒级）对所有线程的调用栈进行Dump。</li>
<li>汇总并统计每次调用栈的Dump结果，在一定时间内采到足够的样本后，导出统计结果，内容是每个方法被采样到的次数及方法的调用关系。</li>
</ol>

<p>Instrumentation则是利用Instrument API，对所有必要的Class进行字节码增强，在进入每个方法前进行埋点，方法执行结束后统计本次方法执行耗时，最终进行汇总。二者都能得到想要的结果，那么它们有什么区别呢？或者说，孰优孰劣？</p>

<p>Instrumentation方式对几乎所有方法添加了额外的AOP逻辑，这会导致对线上服务造成巨额的性能影响，但其优势是：绝对精准的方法调用次数、调用时间统计。</p>

<p>Sampling方式基于无侵入的额外线程对所有线程的调用栈快照进行固定频率抽样，相对前者来说它的性能开销很低。但由于它基于“采样”的模式，以及JVM固有的只能在安全点（Safe Point）进行采样的“缺陷”，会导致统计结果存在一定的偏差。譬如说：某些方法执行时间极短，但执行频率很高，真实占用了大量的CPU Time，但Sampling Profiler的采样周期不能无限调小，这会导致性能开销骤增，所以会导致大量的样本调用栈中并不存在刚才提到的”高频小方法“，进而导致最终结果无法反映真实的CPU热点。更多Sampling相关的问题可以参考《<a href="https://psy-lob-saw.blogspot.com/2016/02/why-most-sampling-java-profilers-are.html" target="_blank">Why (Most) Sampling Java Profilers Are Fucking Terrible</a>》。</p>

<p>具体到“孰优孰劣”的问题层面，这两种实现技术并没有非常明显的高下之判，只有在分场景讨论下才有意义。Sampling由于低开销的特性，更适合用在CPU密集型的应用中，以及不可接受大量性能开销的线上服务中。而Instrumentation则更适合用在I/O密集的应用中、对性能开销不敏感以及确实需要精确统计的场景中。社区的Profiler更多的是基于Sampling来实现，本文也是基于Sampling来进行讲解。</p>

<h3 id="基于java-agent-jmx实现">基于Java Agent + JMX实现</h3>

<p>一个最简单的Sampling CPU Profiler可以用Java Agent + JMX方式来实现。以Java Agent为入口，进入目标JVM进程后开启一个ScheduledExecutorService，定时利用JMX的threadMXBean.dumpAllThreads()来导出所有线程的StackTrace，最终汇总并导出即可。</p>

<p>Uber的<a href="https://github.com/uber-common/jvm-profiler" target="_blank">JVM-Profiler</a>实现原理也是如此，关键部分代码如下：</p>

<pre><code class="language-java">// com/uber/profiling/profilers/StacktraceCollectorProfiler.java

/*
 * StacktraceCollectorProfiler等同于文中所述CpuProfiler，仅命名偏好不同而已
 * jvm-profiler的CpuProfiler指代的是CpuLoad指标的Profiler
 */

// 实现了Profiler接口，外部由统一的ScheduledExecutorService对所有Profiler定时执行
@Override
public void profile() {
    ThreadInfo[] threadInfos = threadMXBean.dumpAllThreads(false, false);
    // ...
    for (ThreadInfo threadInfo : threadInfos) {
        String threadName = threadInfo.getThreadName();
        // ...
        StackTraceElement[] stackTraceElements = threadInfo.getStackTrace();
        // ...
        for (int i = stackTraceElements.length - 1; i &gt;= 0; i--) {
            StackTraceElement stackTraceElement = stackTraceElements[i];
            // ...
        }
        // ...
    }
}
</code></pre>

<p>Uber提供的定时器默认Interval是100ms，对于CPU Profiler来说，这略显粗糙。但由于dumpAllThreads()的执行开销不容小觑，Interval不宜设置的过小，所以该方法的CPU Profiling结果会存在不小的误差。</p>

<p>JVM-Profiler的优点在于支持多种指标的Profiling（StackTrace、CPUBusy、Memory、I/O、Method），且支持将Profiling结果通过Kafka上报回中心Server进行分析，也即支持集群诊断。</p>

<h3 id="基于jvmti-getstacktrace实现">基于JVMTI + GetStackTrace实现</h3>

<p>使用Java实现Profiler相对较简单，但也存在一些问题，譬如说Java Agent代码与业务代码共享AppClassLoader，被JVM直接加载的agent.jar如果引入了第三方依赖，可能会对业务Class造成污染。截止发稿时，JVM-Profiler都存在这个问题，它引入了Kafka-Client、http-Client、Jackson等组件，如果与业务代码中的组件版本发生冲突，可能会引发未知错误。<a href="https://github.com/oldmanpushcart/greys-anatomy" target="_blank">Greys</a>/<a href="https://github.com/alibaba/arthas" target="_blank">Arthas</a>/<a href="https://github.com/alibaba/jvm-sandbox" target="_blank">JVM-Sandbox</a>的解决方式是分离入口与核心代码，使用定制的ClassLoader加载核心代码，避免影响业务代码。</p>

<p>在更底层的C/C++层面，我们可以直接对接JVMTI接口，使用原生C API对JVM进行操作，功能更丰富更强大，但开发效率偏低。基于上节同样的原理开发CPU Profiler，使用JVMTI需要进行如下这些步骤：</p>

<p>1.编写Agent_OnLoad()，在入口通过JNI的JavaVM*指针的GetEnv()函数拿到JVMTI的jvmtiEnv指针：</p>

<pre><code class="language-c">// agent.c

JNIEXPORT jint JNICALL Agent_OnLoad(JavaVM *vm, char *options, void *reserved) {
    jvmtiEnv *jvmti;
    (*vm)-&gt;GetEnv((void **)&amp;jvmti, JVMTI_VERSION_1_0);
    // ...
    return JNI_OK;
}
</code></pre>

<p>2.开启一个线程定时循环，定时使用jvmtiEnv指针配合调用如下几个JVMTI函数：</p>

<pre><code class="language-java">// 获取所有线程的jthread
jvmtiError GetAllThreads(jvmtiEnv *env, jint *threads_count_ptr, jthread **threads_ptr);

// 根据jthread获取该线程信息（name、daemon、priority...）
jvmtiError GetThreadInfo(jvmtiEnv *env, jthread thread, jvmtiThreadInfo* info_ptr);

// 根据jthread获取该线程调用栈
jvmtiError GetStackTrace(jvmtiEnv *env,
                         jthread thread,
                         jint start_depth,
                         jint max_frame_count,
                         jvmtiFrameInfo *frame_buffer,
                         jint *count_ptr);
</code></pre>

<p>主逻辑大致是：首先调用GetAllThreads()获取所有线程的“句柄”jthread，然后遍历根据jthread调用GetThreadInfo()获取线程信息，按线程名过滤掉不需要的线程后，继续遍历根据jthread调用GetStackTrace()获取线程的调用栈。</p>

<p>3.在Buffer中保存每一次的采样结果，最终生成必要的统计数据即可。</p>

<p>按如上步骤即可实现基于JVMTI的CPU Profiler。但需要说明的是，即便是基于原生JVMTI接口使用GetStackTrace()的方式获取调用栈，也存在与JMX相同的问题——只能在安全点（Safe Point）进行采样。</p>

<h3 id="safepoint-bias问题">SafePoint Bias问题</h3>

<p>基于Sampling的CPU Profiler通过采集程序在不同时间点的调用栈样本来近似地推算出热点方法，因此，从理论上来讲Sampling CPU Profiler必须遵循以下两个原则：</p>

<ol>
<li>样本必须足够多。</li>
<li>程序中所有正在运行的代码点都必须以相同的概率被Profiler采样。</li>
</ol>

<p>如果只能在安全点采样，就违背了第二条原则。因为我们只能采集到位于安全点时刻的调用栈快照，意味着某些代码可能永远没有机会被采样，即使它真实耗费了大量的CPU执行时间，这种现象被称为“SafePoint Bias”。</p>

<p>上文我们提到，基于JMX与基于JVMTI的Profiler实现都存在SafePoint Bias，但一个值得了解的细节是：单独来说，JVMTI的GetStackTrace()函数并不需要在Caller的安全点执行，但当调用GetStackTrace()获取其他线程的调用栈时，必须等待，直到目标线程进入安全点；而且，GetStackTrace()仅能通过单独的线程同步定时调用，不能在UNIX信号处理器的Handler中被异步调用。综合来说，GetStackTrace()存在与JMX一样的SafePoint Bias。更多安全点相关的知识可以参考《Safepoints: Meaning, Side Effects and Overheads》。</p>

<p>那么，如何避免SafePoint Bias？社区提供了一种Hack思路——AsyncGetCallTrace。</p>

<h3 id="基于jvmti-asyncgetcalltrace实现">基于JVMTI + AsyncGetCallTrace实现</h3>

<p>如上节所述，假如我们拥有一个函数可以获取当前线程的调用栈且不受安全点干扰，另外它还支持在UNIX信号处理器中被异步调用，那么我们只需注册一个UNIX信号处理器，在Handler中调用该函数获取当前线程的调用栈即可。由于UNIX信号会被发送给进程的随机一线程进行处理，因此最终信号会均匀分布在所有线程上，也就均匀获取了所有线程的调用栈样本。</p>

<p>OracleJDK/OpenJDK内部提供了这么一个函数——AsyncGetCallTrace，它的原型如下：</p>

<pre><code>// 栈帧
typedef struct {
 jint lineno;
 jmethodID method_id;
} AGCT_CallFrame;

// 调用栈
typedef struct {
    JNIEnv *env;
    jint num_frames;
    AGCT_CallFrame *frames;
} AGCT_CallTrace;

// 根据ucontext将调用栈填充进trace指针
void AsyncGetCallTrace(AGCT_CallTrace *trace, jint depth, void *ucontext);
</code></pre>

<p>通过原型可以看到，该函数的使用方式非常简洁，直接通过ucontext就能获取到完整的Java调用栈。</p>

<p>顾名思义，AsyncGetCallTrace是“async”的，不受安全点影响，这样的话采样就可能发生在任何时间，包括Native代码执行期间、GC期间等，在这时我们是无法获取Java调用栈的，AGCT_CallTrace的num_frames字段正常情况下标识了获取到的调用栈深度，但在如前所述的异常情况下它就表示为负数，最常见的-2代表此刻正在GC。</p>

<p>由于AsyncGetCallTrace非标准JVMTI函数，因此我们无法在jvmti.h中找到该函数声明，且由于其目标文件也早已链接进JVM二进制文件中，所以无法通过简单的声明来获取该函数的地址，这需要通过一些Trick方式来解决。简单说，Agent最终是作为动态链接库加载到目标JVM进程的地址空间中，因此可以在Agent_OnLoad内通过glibc提供的dlsym()函数拿到当前地址空间（即目标JVM进程地址空间）名为“AsyncGetCallTrace”的符号地址。这样就拿到了该函数的指针，按照上述原型进行类型转换后，就可以正常调用了。</p>

<p>通过AsyncGetCallTrace实现CPU Profiler的大致流程：</p>

<p>1.编写Agent_OnLoad()，在入口拿到jvmtiEnv和AsyncGetCallTrace指针，获取AsyncGetCallTrace方式如下:</p>

<pre><code>typedef void (*AsyncGetCallTrace)(AGCT_CallTrace *traces, jint depth, void *ucontext);
// ...
AsyncGetCallTrace agct_ptr = (AsyncGetCallTrace)dlsym(RTLD_DEFAULT, &quot;AsyncGetCallTrace&quot;);
if (agct_ptr == NULL) {
    void *libjvm = dlopen(&quot;libjvm.so&quot;, RTLD_NOW);
    if (!libjvm) {
        // 处理dlerror()...
    }
    agct_ptr = (AsyncGetCallTrace)dlsym(libjvm, &quot;AsyncGetCallTrace&quot;);
}
</code></pre>

<p>2.在OnLoad阶段，我们还需要做一件事，即注册OnClassLoad和OnClassPrepare这两个Hook，原因是jmethodID是延迟分配的，使用AGCT获取Traces依赖预先分配好的数据。我们在OnClassPrepare的CallBack中尝试获取该Class的所有Methods，这样就使JVMTI提前分配了所有方法的jmethodID，如下所示：</p>

<pre><code>void JNICALL OnClassLoad(jvmtiEnv *jvmti, JNIEnv* jni, jthread thread, jclass klass) {}

void JNICALL OnClassPrepare(jvmtiEnv *jvmti, JNIEnv *jni, jthread thread, jclass klass) {
    jint method_count;
    jmethodID *methods;
    jvmti-&gt;GetClassMethods(klass, &amp;method_count, &amp;methods);
    delete [] methods;
}

// ...

jvmtiEventCallbacks callbacks = {0};
callbacks.ClassLoad = OnClassLoad;
callbacks.ClassPrepare = OnClassPrepare;
jvmti-&gt;SetEventCallbacks(&amp;callbacks, sizeof(callbacks));
jvmti-&gt;SetEventNotificationMode(JVMTI_ENABLE, JVMTI_EVENT_CLASS_LOAD, NULL);
jvmti-&gt;SetEventNotificationMode(JVMTI_ENABLE, JVMTI_EVENT_CLASS_PREPARE, NULL);
</code></pre>

<p>3.利用SIGPROF信号来进行定时采样：</p>

<pre><code>// 这里信号handler传进来的的ucontext即AsyncGetCallTrace需要的ucontext
void signal_handler(int signo, siginfo_t *siginfo, void *ucontext) {
    // 使用AsyncCallTrace进行采样，注意处理num_frames为负的异常情况
}

// ...

// 注册SIGPROF信号的handler
struct sigaction sa;
sigemptyset(&amp;sa.sa_mask);
sa.sa_sigaction = signal_handler;
sa.sa_flags = SA_RESTART | SA_SIGINFO;
sigaction(SIGPROF, &amp;sa, NULL);

// 定时产生SIGPROF信号
// interval是nanoseconds表示的采样间隔，AsyncGetCallTrace相对于同步采样来说可以适当高频一些
long sec = interval / 1000000000;
long usec = (interval % 1000000000) / 1000;
struct itimerval tv = {{sec, usec}, {sec, usec}};
setitimer(ITIMER_PROF, &amp;tv, NULL);
</code></pre>

<p>4.在Buffer中保存每一次的采样结果，最终生成必要的统计数据即可。</p>

<p>按如上步骤即可实现基于AsyncGetCallTrace的CPU Profiler，这是社区中目前性能开销最低、相对效率最高的CPU Profiler实现方式，在Linux环境下结合perf_events还能做到同时采样Java栈与Native栈，也就能同时分析Native代码中存在的性能热点。该方式的典型开源实现有<a href="https://github.com/jvm-profiling-tools/async-profiler" target="_blank">Async-Profiler</a>和<a href="https://github.com/jvm-profiling-tools/honest-profiler" target="_blank">Honest-Profiler</a>，Async-Profiler实现质量较高，感兴趣的话建议大家阅读参考文章。有趣的是，IntelliJ IDEA内置的Java Profiler，其实就是Async-Profiler的包装。更多关于AsyncGetCallTrace的内容，大家可以参考《<a href="https://psy-lob-saw.blogspot.com/2016/06/the-pros-and-cons-of-agct.html" target="_blank">The Pros and Cons of AsyncGetCallTrace Profilers</a>》。</p>

<h3 id="生成性能火焰图">生成性能火焰图</h3>

<p>现在我们拥有了采样调用栈的能力，但是调用栈样本集是以二维数组的数据结构形式存在于内存中的，如何将其转换为可视化的火焰图呢？</p>

<p>火焰图通常是一个svg文件，部分优秀项目可以根据文本文件自动生成火焰图文件，仅对文本文件的格式有一定要求。FlameGraph项目的核心只是一个Perl脚本，可以根据我们提供的调用栈文本生成相应的火焰图svg文件。调用栈的文本格式相当简单，如下所示：</p>

<pre><code>base_func;func1;func2;func3 10
base_func;funca;funcb 15
</code></pre>

<p>将我们采样到的调用栈样本集进行整合后，需输出如上所示的文本格式。每一行代表一“类“调用栈，空格左边是调用栈的方法名排列，以分号分割，左栈底右栈顶，空格右边是该样本出现的次数。</p>

<p>将样本文件交给flamegraph.pl脚本执行，就能输出相应的火焰图了：</p>

<pre><code>$ flamegraph.pl stacktraces.txt &gt; stacktraces.svg
</code></pre>

<p>效果如下图所示：</p>

<p><img src="assets/ae2b3dda630d2de82eb632a6e8d5bee9336049.png" alt="通过flamegraph.pl生成的火焰图" /></p>

<p>通过flamegraph.pl生成的火焰图</p>

<h2 id="hotspot的dynamic-attach机制解析">HotSpot的Dynamic Attach机制解析</h2>

<p>到目前为止，我们已经了解了CPU Profiler完整的工作原理，然而使用过JProfiler/Arthas的同学可能会有疑问，很多情况下可以直接对线上运行中的服务进行Profling，并不需要在Java进程的启动参数添加Agent参数，这是通过什么手段做到的？答案是Dynamic Attach。</p>

<p>JDK在1.6以后提供了Attach API，允许向运行中的JVM进程添加Agent，这项手段被广泛使用在各种Profiler和字节码增强工具中，其官方简介如下：</p>

<blockquote>
<p>This is a Sun extension that allows a tool to ‘attach’ to another process running Java code and launch a JVM TI agent or a java.lang.instrument agent in that process.</p>
</blockquote>

<p>总的来说，Dynamic Attach是HotSpot提供的一种特殊能力，它允许一个进程向另一个运行中的JVM进程发送一些命令并执行，命令并不限于加载Agent，还包括Dump内存、Dump线程等等。</p>

<h3 id="通过sun-tools进行attach">通过sun.tools进行Attach</h3>

<p>Attach虽然是HotSpot提供的能力，但JDK在Java层面也对其做了封装。</p>

<p>前文已经提到，对于Java Agent来说，PreMain方法在Agent作为启动参数运行的时候执行，其实我们还可以额外实现一个AgentMain方法，并在MANIFEST.MF中将Agent-Class指定为该Class：</p>

<pre><code>public static void agentmain(String args, Instrumentation ins) {
    // implement
}
</code></pre>

<p>这样打包出来的jar，既可以作为-javaagent参数启动，也可以被Attach到运行中的目标JVM进程。JDK已经封装了简单的API让我们直接Attach一个Java Agent，下面以Arthas中的代码进行演示：</p>

<pre><code>// com/taobao/arthas/core/Arthas.java

import com.sun.tools.attach.VirtualMachine;
import com.sun.tools.attach.VirtualMachineDescriptor;

// ...

private void attachAgent(Configure configure) throws Exception {
    VirtualMachineDescriptor virtualMachineDescriptor = null;
  
    // 拿到所有JVM进程，找出目标进程
    for (VirtualMachineDescriptor descriptor : VirtualMachine.list()) {
        String pid = descriptor.id();
        if (pid.equals(Integer.toString(configure.getJavaPid()))) {
            virtualMachineDescriptor = descriptor;
        }
    }
    VirtualMachine virtualMachine = null;
    try {
        // 针对某个JVM进程调用VirtualMachine.attach()方法，拿到VirtualMachine实例
        if (null == virtualMachineDescriptor) {
            virtualMachine = VirtualMachine.attach(&quot;&quot; + configure.getJavaPid());
        } else {
            virtualMachine = VirtualMachine.attach(virtualMachineDescriptor);
        }

        // ...

        // 调用VirtualMachine#loadAgent()，将arthasAgentPath指定的jar attach到目标JVM进程中
        // 第二个参数为attach参数，即agentmain的首个String参数args
        virtualMachine.loadAgent(arthasAgentPath, configure.getArthasCore() + &quot;;&quot; + configure.toString());
    } finally {
        if (null != virtualMachine) {
            // 调用VirtualMachine#detach()释放
            virtualMachine.detach();
        }
    }
}
</code></pre>

<h3 id="直接对hotspot进行attach">直接对HotSpot进行Attach</h3>

<p>sun.tools封装的API足够简单易用，但只能使用Java编写，也只能用在Java Agent上，因此有些时候我们必须手工对JVM进程直接进行Attach。对于JVMTI，除了Agent_OnLoad()之外，我们还需实现一个Agent_OnAttach()函数，当将JVMTI Agent Attach到目标进程时，从该函数开始执行：</p>

<pre><code>// $JAVA_HOME/include/jvmti.h

JNIEXPORT jint JNICALL Agent_OnAttach(JavaVM *vm, char *options, void *reserved);
</code></pre>

<p>下面我们以Async-Profiler中的jattach源码为线索，探究一下如何利用Attach机制给运行中的JVM进程发送命令。jattach是Async-Profiler提供的一个Driver，使用方式比较直观：</p>

<pre><code>Usage:
    jattach &lt;pid&gt; &lt;cmd&gt; [args ...]
Args:
    &lt;pid&gt;  目标JVM进程的进程ID
    &lt;cmd&gt;  要执行的命令
    &lt;args&gt; 命令参数
</code></pre>

<p>使用方式如：</p>

<pre><code>$ jattach 1234 load /absolute/path/to/agent/libagent.so true
</code></pre>

<p>执行上述命令，libagent.so就被加载到ID为1234的JVM进程中并开始执行Agent_OnAttach函数了。有一点需要注意，执行Attach的进程euid及egid，与被Attach的目标JVM进程必须相同。接下来开始分析jattach源码。</p>

<p>如下所示的Main函数描述了一次Attach的整体流程：</p>

<pre><code>// async-profiler/src/jattach/jattach.c

int main(int argc, char** argv) {
    // 解析命令行参数
    // 检查euid与egid
    // ...

    if (!check_socket(nspid) &amp;&amp; !start_attach_mechanism(pid, nspid)) {
        perror(&quot;Could not start attach mechanism&quot;);
        return 1;
    }

    int fd = connect_socket(nspid);
    if (fd == -1) {
        perror(&quot;Could not connect to socket&quot;);
        return 1;
    }

    printf(&quot;Connected to remote JVM\n&quot;);
    if (!write_command(fd, argc - 2, argv + 2)) {
        perror(&quot;Error writing to socket&quot;);
        close(fd);
        return 1;
    }
    printf(&quot;Response code = &quot;);
    fflush(stdout);

    int result = read_response(fd);
    close(fd);
    return result;
}
</code></pre>

<p>忽略掉命令行参数解析与检查euid和egid的过程。jattach首先调用了check_socket函数进行了“socket检查？”，check_socket源码如下：</p>

<pre><code>// async-profiler/src/jattach/jattach.c

// Check if remote JVM has already opened socket for Dynamic Attach
static int check_socket(int pid) {
    char path[MAX_PATH];
    snprintf(path, MAX_PATH, &quot;%s/.java_pid%d&quot;, get_temp_directory(), pid); // get_temp_directory()在Linux下固定返回&quot;/tmp&quot;
    struct stat stats;
    return stat(path, &amp;stats) == 0 &amp;&amp; S_ISSOCK(stats.st_mode);
}
</code></pre>

<p>我们知道，UNIX操作系统提供了一种基于文件的Socket接口，称为“UNIX Socket”（一种常用的进程间通信方式）。在该函数中使用S_ISSOCK宏来判断该文件是否被绑定到了UNIX Socket，如此看来，“/tmp/.java_pid”文件很有可能就是外部进程与JVM进程间通信的桥梁。</p>

<p>查阅官方文档，得到如下描述：</p>

<blockquote>
<p>The attach listener thread then communicates with the source JVM in an OS dependent manner: - On Solaris, the Doors IPC mechanism is used. The door is attached to a file in the file system so that clients can access it. - On Linux, a Unix domain socket is used. This socket is bound to a file in the filesystem so that clients can access it. - On Windows, the created thread is given the name of a pipe which is served by the client. The result of the operations are written to this pipe by the target JVM.</p>
</blockquote>

<p>证明了我们的猜想是正确的。目前为止check_socket函数的作用很容易理解了：判断外部进程与目标JVM进程之间是否已经建立了UNIX Socket连接。</p>

<p>回到Main函数，在使用check_socket确定连接尚未建立后，紧接着调用start_attach_mechanism函数，函数名很直观地描述了它的作用，源码如下：</p>

<pre><code>// async-profiler/src/jattach/jattach.c

// Force remote JVM to start Attach listener.
// HotSpot will start Attach listener in response to SIGQUIT if it sees .attach_pid file
static int start_attach_mechanism(int pid, int nspid) {
    char path[MAX_PATH];
    snprintf(path, MAX_PATH, &quot;/proc/%d/cwd/.attach_pid%d&quot;, nspid, nspid);

    int fd = creat(path, 0660);
    if (fd == -1 || (close(fd) == 0 &amp;&amp; !check_file_owner(path))) {
        // Failed to create attach trigger in current directory. Retry in /tmp
        snprintf(path, MAX_PATH, &quot;%s/.attach_pid%d&quot;, get_temp_directory(), nspid);
        fd = creat(path, 0660);
        if (fd == -1) {
            return 0;
        }
        close(fd);
    }

    // We have to still use the host namespace pid here for the kill() call
    kill(pid, SIGQUIT);

    // Start with 20 ms sleep and increment delay each iteration
    struct timespec ts = {0, 20000000};
    int result;
    do {
        nanosleep(&amp;ts, NULL);
        result = check_socket(nspid);
    } while (!result &amp;&amp; (ts.tv_nsec += 20000000) &lt; 300000000);

    unlink(path);
    return result;
}
</code></pre>

<p>start_attach_mechanism函数首先创建了一个名为“/tmp/.attach_pid”的空文件，然后向目标JVM进程发送了一个SIGQUIT信号，这个信号似乎触发了JVM的某种机制？紧接着，start_attach_mechanism函数开始陷入了一种等待，每20ms调用一次check_socket函数检查连接是否被建立，如果等了300ms还没有成功就放弃。函数的最后调用Unlink删掉.attach_pid文件并返回。</p>

<p>如此看来，HotSpot似乎提供了一种特殊的机制，只要给它发送一个SIGQUIT信号，并预先准备好.attach_pid文件，HotSpot会主动创建一个地址为“/tmp/.java_pid”的UNIX Socket，接下来主动Connect这个地址即可建立连接执行命令。</p>

<p>查阅文档，得到如下描述：</p>

<blockquote>
<p>Dynamic attach has an attach listener thread in the target JVM. This is a thread that is started when the first attach request occurs. On Linux and Solaris, the client creates a file named .attach_pid(pid) and sends a SIGQUIT to the target JVM process. The existence of this file causes the SIGQUIT handler in HotSpot to start the attach listener thread. On Windows, the client uses the Win32 CreateRemoteThread function to create a new thread in the target process.</p>
</blockquote>

<p>这样一来就很明确了，在Linux上我们只需创建一个“/tmp/.attach_pid”文件，并向目标JVM进程发送一个SIGQUIT信号，HotSpot就会开始监听“/tmp/.java_pid”地址上的UNIX Socket，接收并执行相关Attach的命令。至于为什么一定要创建.attach_pid文件才可以触发Attach Listener的创建，经查阅资料，我们得到了两种说法：一是JVM不止接收从外部Attach进程发送的SIGQUIT信号，必须配合外部进程创建的外部文件才能确定这是一次Attach请求；二是为了安全。</p>

<p>继续看jattach的源码，果不其然，它调用了connect_socket函数对“/tmp/.java_pid”进行连接，connect_socket源码如下：</p>

<pre><code>// async-profiler/src/jattach/jattach.c

// Connect to UNIX domain socket created by JVM for Dynamic Attach
static int connect_socket(int pid) {
    int fd = socket(PF_UNIX, SOCK_STREAM, 0);
    if (fd == -1) {
        return -1;
    }

    struct sockaddr_un addr;
    addr.sun_family = AF_UNIX;
    snprintf(addr.sun_path, sizeof(addr.sun_path), &quot;%s/.java_pid%d&quot;, get_temp_directory(), pid);

    if (connect(fd, (struct sockaddr*)&amp;addr, sizeof(addr)) == -1) {
        close(fd);
        return -1;
    }
    return fd;
}
</code></pre>

<p>一个很普通的Socket创建函数，返回Socket文件描述符。</p>

<p>回到Main函数，主流程紧接着调用write_command函数向该Socket写入了从命令行传进来的参数，并且调用read_response函数接收从目标JVM进程返回的数据。两个很常见的Socket读写函数，源码如下：</p>

<pre><code>// async-profiler/src/jattach/jattach.c

// Send command with arguments to socket
static int write_command(int fd, int argc, char** argv) {
    // Protocol version
    if (write(fd, &quot;1&quot;, 2) &lt;= 0) {
        return 0;
    }

    int i;
    for (i = 0; i &lt; 4; i++) {
        const char* arg = i &lt; argc ? argv[i] : &quot;&quot;;
        if (write(fd, arg, strlen(arg) + 1) &lt;= 0) {
            return 0;
        }
    }
    return 1;
}

// Mirror response from remote JVM to stdout
static int read_response(int fd) {
    char buf[8192];
    ssize_t bytes = read(fd, buf, sizeof(buf) - 1);
    if (bytes &lt;= 0) {
        perror(&quot;Error reading response&quot;);
        return 1;
    }

    // First line of response is the command result code
    buf[bytes] = 0;
    int result = atoi(buf);

    do {
        fwrite(buf, 1, bytes, stdout);
        bytes = read(fd, buf, sizeof(buf));
    } while (bytes &gt; 0);
    return result;
}
</code></pre>

<p>浏览write_command函数就可知外部进程与目标JVM进程之间发送的数据格式相当简单，基本如下所示：</p>

<pre><code>&lt;PROTOCOL VERSION&gt;\0&lt;COMMAND&gt;\0&lt;ARG1&gt;\0&lt;ARG2&gt;\0&lt;ARG3&gt;\0
</code></pre>

<p>以先前我们使用的Load命令为例，发送给HotSpot时格式如下：</p>

<pre><code>1\0load\0/absolute/path/to/agent/libagent.so\0true\0\0
</code></pre>

<p>至此，我们已经了解了如何手工对JVM进程直接进行Attach。</p>

<h3 id="attach补充介绍">Attach补充介绍</h3>

<p>Load命令仅仅是HotSpot所支持的诸多命令中的一种，用于动态加载基于JVMTI的Agent，完整的命令表如下所示：</p>

<pre><code>static AttachOperationFunctionInfo funcs[] = {
  { &quot;agentProperties&quot;,  get_agent_properties },
  { &quot;datadump&quot;,         data_dump },
  { &quot;dumpheap&quot;,         dump_heap },
  { &quot;load&quot;,             JvmtiExport::load_agent_library },
  { &quot;properties&quot;,       get_system_properties },
  { &quot;threaddump&quot;,       thread_dump },
  { &quot;inspectheap&quot;,      heap_inspection },
  { &quot;setflag&quot;,          set_flag },
  { &quot;printflag&quot;,        print_flag },
  { &quot;jcmd&quot;,             jcmd },
  { NULL,               NULL }
};
</code></pre>

<p>读者可以尝试下threaddump命令，然后对相同的进程进行jstack，对比观察输出，其实是完全相同的，其它命令大家可以自行进行探索。</p>

<h2 id="总结">总结</h2>

<p>总的来说，善用各类Profiler是提升性能优化效率的一把利器，了解Profiler本身的实现原理更能帮助我们避免对工具的各种误用。CPU Profiler所依赖的Attach、JVMTI、Instrumentation、JMX等皆是JVM平台比较通用的技术，在此基础上，我们去实现Memory Profiler、Thread Profiler、GC Analyzer等工具也没有想象中那么神秘和复杂了。</p>

<h2 id="参考资料">参考资料</h2>

<ul>
<li><a href="https://docs.oracle.com/en/java/javase/12/docs/specs/jvmti.html" target="_blank">JVM Tool Interface</a></li>
<li><a href="https://psy-lob-saw.blogspot.com/2016/06/the-pros-and-cons-of-agct.html" target="_blank">The Pros and Cons of AsyncGetCallTrace Profilers</a></li>
<li><a href="https://psy-lob-saw.blogspot.com/2016/02/why-most-sampling-java-profilers-are.html" target="_blank">Why (Most) Sampling Java Profilers Are Fucking Terrible</a></li>
<li><a href="https://psy-lob-saw.blogspot.com/2015/12/safepoints.html" target="_blank">Safepoints: Meaning, Side Effects and Overheads</a></li>
<li><a href="http://openjdk.java.net/groups/hotspot/docs/Serviceability.html" target="_blank">Serviceability in HotSpot</a></li>
<li><a href="http://www.ruanyifeng.com/blog/2017/09/flame-graph.html" target="_blank">如何读懂火焰图？</a></li>
<li><a href="https://blog.jetbrains.com/idea/2018/09/intellij-idea-2018-3-eap-git-submodules-jvm-profiler-macos-and-linux-and-more/" target="_blank">IntelliJ IDEA 2018.3 EAP: Git Submodules, JVM Profiler (macOS and Linux) and more</a></li>
</ul>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#1f737373262b2e2e2f285f78727e7673317c7072" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0c6236f97cede4',t:'MTczMzk5MjU1NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
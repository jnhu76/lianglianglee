<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=06&#32;&#32;日志框架千千万，MyBatis&#32;都能兼容的秘密是什么？>
        <link rel="icon" href="/static/favicon.png">
        <title>06  日志框架千千万，MyBatis 都能兼容的秘密是什么？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  领略 MyBatis 设计思维，突破持久化技术瓶颈.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e9%a2%86%e7%95%a5%20MyBatis%20%e8%ae%be%e8%ae%a1%e6%80%9d%e7%bb%b4%ef%bc%8c%e7%aa%81%e7%a0%b4%e6%8c%81%e4%b9%85%e5%8c%96%e6%8a%80%e6%9c%af%e7%93%b6%e9%a2%88.md">00 开篇词  领略 MyBatis 设计思维，突破持久化技术瓶颈.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  常见持久层框架赏析，到底是什么让你选择 MyBatis？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/01%20%20%e5%b8%b8%e8%a7%81%e6%8c%81%e4%b9%85%e5%b1%82%e6%a1%86%e6%9e%b6%e8%b5%8f%e6%9e%90%ef%bc%8c%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%e8%ae%a9%e4%bd%a0%e9%80%89%e6%8b%a9%20MyBatis%ef%bc%9f.md">01  常见持久层框架赏析，到底是什么让你选择 MyBatis？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  订单系统持久层示例分析，20 分钟带你快速上手 MyBatis.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/02%20%20%e8%ae%a2%e5%8d%95%e7%b3%bb%e7%bb%9f%e6%8c%81%e4%b9%85%e5%b1%82%e7%a4%ba%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%8c20%20%e5%88%86%e9%92%9f%e5%b8%a6%e4%bd%a0%e5%bf%ab%e9%80%9f%e4%b8%8a%e6%89%8b%20MyBatis.md">02  订单系统持久层示例分析，20 分钟带你快速上手 MyBatis.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  MyBatis 源码环境搭建及整体架构解析.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/03%20%20MyBatis%20%e6%ba%90%e7%a0%81%e7%8e%af%e5%a2%83%e6%90%ad%e5%bb%ba%e5%8f%8a%e6%95%b4%e4%bd%93%e6%9e%b6%e6%9e%84%e8%a7%a3%e6%9e%90.md">03  MyBatis 源码环境搭建及整体架构解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  MyBatis 反射工具箱：带你领略不一样的反射设计思路.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/04%20%20MyBatis%20%e5%8f%8d%e5%b0%84%e5%b7%a5%e5%85%b7%e7%ae%b1%ef%bc%9a%e5%b8%a6%e4%bd%a0%e9%a2%86%e7%95%a5%e4%b8%8d%e4%b8%80%e6%a0%b7%e7%9a%84%e5%8f%8d%e5%b0%84%e8%ae%be%e8%ae%a1%e6%80%9d%e8%b7%af.md">04  MyBatis 反射工具箱：带你领略不一样的反射设计思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  数据库类型体系与 Java 类型体系之间的“爱恨情仇”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/05%20%20%e6%95%b0%e6%8d%ae%e5%ba%93%e7%b1%bb%e5%9e%8b%e4%bd%93%e7%b3%bb%e4%b8%8e%20Java%20%e7%b1%bb%e5%9e%8b%e4%bd%93%e7%b3%bb%e4%b9%8b%e9%97%b4%e7%9a%84%e2%80%9c%e7%88%b1%e6%81%a8%e6%83%85%e4%bb%87%e2%80%9d.md">05  数据库类型体系与 Java 类型体系之间的“爱恨情仇”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  日志框架千千万，MyBatis 都能兼容的秘密是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/06%20%20%e6%97%a5%e5%bf%97%e6%a1%86%e6%9e%b6%e5%8d%83%e5%8d%83%e4%b8%87%ef%bc%8cMyBatis%20%e9%83%bd%e8%83%bd%e5%85%bc%e5%ae%b9%e7%9a%84%e7%a7%98%e5%af%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">06  日志框架千千万，MyBatis 都能兼容的秘密是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  深入数据源和事务，把握持久化框架的两个关键命脉.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/07%20%20%e6%b7%b1%e5%85%a5%e6%95%b0%e6%8d%ae%e6%ba%90%e5%92%8c%e4%ba%8b%e5%8a%a1%ef%bc%8c%e6%8a%8a%e6%8f%a1%e6%8c%81%e4%b9%85%e5%8c%96%e6%a1%86%e6%9e%b6%e7%9a%84%e4%b8%a4%e4%b8%aa%e5%85%b3%e9%94%ae%e5%91%bd%e8%84%89.md">07  深入数据源和事务，把握持久化框架的两个关键命脉.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  Mapper 文件与 Java 接口的优雅映射之道.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/08%20%20Mapper%20%e6%96%87%e4%bb%b6%e4%b8%8e%20Java%20%e6%8e%a5%e5%8f%a3%e7%9a%84%e4%bc%98%e9%9b%85%e6%98%a0%e5%b0%84%e4%b9%8b%e9%81%93.md">08  Mapper 文件与 Java 接口的优雅映射之道.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  基于 MyBatis 缓存分析装饰器模式的最佳实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/09%20%20%e5%9f%ba%e4%ba%8e%20MyBatis%20%e7%bc%93%e5%ad%98%e5%88%86%e6%9e%90%e8%a3%85%e9%a5%b0%e5%99%a8%e6%a8%a1%e5%bc%8f%e7%9a%84%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5.md">09  基于 MyBatis 缓存分析装饰器模式的最佳实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  鸟瞰 MyBatis 初始化，把握 MyBatis 启动流程脉络（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/10%20%20%e9%b8%9f%e7%9e%b0%20MyBatis%20%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%8c%e6%8a%8a%e6%8f%a1%20MyBatis%20%e5%90%af%e5%8a%a8%e6%b5%81%e7%a8%8b%e8%84%89%e7%bb%9c%ef%bc%88%e4%b8%8a%ef%bc%89.md">10  鸟瞰 MyBatis 初始化，把握 MyBatis 启动流程脉络（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  鸟瞰 MyBatis 初始化，把握 MyBatis 启动流程脉络（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/11%20%20%e9%b8%9f%e7%9e%b0%20MyBatis%20%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%8c%e6%8a%8a%e6%8f%a1%20MyBatis%20%e5%90%af%e5%8a%a8%e6%b5%81%e7%a8%8b%e8%84%89%e7%bb%9c%ef%bc%88%e4%b8%8b%ef%bc%89.md">11  鸟瞰 MyBatis 初始化，把握 MyBatis 启动流程脉络（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  深入分析动态 SQL 语句解析全流程（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/12%20%20%e6%b7%b1%e5%85%a5%e5%88%86%e6%9e%90%e5%8a%a8%e6%80%81%20SQL%20%e8%af%ad%e5%8f%a5%e8%a7%a3%e6%9e%90%e5%85%a8%e6%b5%81%e7%a8%8b%ef%bc%88%e4%b8%8a%ef%bc%89.md">12  深入分析动态 SQL 语句解析全流程（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  深入分析动态 SQL 语句解析全流程（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/13%20%20%e6%b7%b1%e5%85%a5%e5%88%86%e6%9e%90%e5%8a%a8%e6%80%81%20SQL%20%e8%af%ad%e5%8f%a5%e8%a7%a3%e6%9e%90%e5%85%a8%e6%b5%81%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">13  深入分析动态 SQL 语句解析全流程（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  探究 MyBatis 结果集映射机制背后的秘密（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/14%20%20%e6%8e%a2%e7%a9%b6%20MyBatis%20%e7%bb%93%e6%9e%9c%e9%9b%86%e6%98%a0%e5%b0%84%e6%9c%ba%e5%88%b6%e8%83%8c%e5%90%8e%e7%9a%84%e7%a7%98%e5%af%86%ef%bc%88%e4%b8%8a%ef%bc%89.md">14  探究 MyBatis 结果集映射机制背后的秘密（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  探究 MyBatis 结果集映射机制背后的秘密（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/15%20%20%e6%8e%a2%e7%a9%b6%20MyBatis%20%e7%bb%93%e6%9e%9c%e9%9b%86%e6%98%a0%e5%b0%84%e6%9c%ba%e5%88%b6%e8%83%8c%e5%90%8e%e7%9a%84%e7%a7%98%e5%af%86%ef%bc%88%e4%b8%8b%ef%bc%89.md">15  探究 MyBatis 结果集映射机制背后的秘密（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  StatementHandler：参数绑定、SQL 执行和结果映射的奠基者.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/16%20%20StatementHandler%ef%bc%9a%e5%8f%82%e6%95%b0%e7%bb%91%e5%ae%9a%e3%80%81SQL%20%e6%89%a7%e8%a1%8c%e5%92%8c%e7%bb%93%e6%9e%9c%e6%98%a0%e5%b0%84%e7%9a%84%e5%a5%a0%e5%9f%ba%e8%80%85.md">16  StatementHandler：参数绑定、SQL 执行和结果映射的奠基者.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  Executor 才是执行 SQL 语句的幕后推手（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/17%20%20Executor%20%e6%89%8d%e6%98%af%e6%89%a7%e8%a1%8c%20SQL%20%e8%af%ad%e5%8f%a5%e7%9a%84%e5%b9%95%e5%90%8e%e6%8e%a8%e6%89%8b%ef%bc%88%e4%b8%8a%ef%bc%89.md">17  Executor 才是执行 SQL 语句的幕后推手（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  Executor 才是执行 SQL 语句的幕后推手（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/18%20%20Executor%20%e6%89%8d%e6%98%af%e6%89%a7%e8%a1%8c%20SQL%20%e8%af%ad%e5%8f%a5%e7%9a%84%e5%b9%95%e5%90%8e%e6%8e%a8%e6%89%8b%ef%bc%88%e4%b8%8b%ef%bc%89.md">18  Executor 才是执行 SQL 语句的幕后推手（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  深入 MyBatis 内核与业务逻辑的桥梁——接口层.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/19%20%20%e6%b7%b1%e5%85%a5%20MyBatis%20%e5%86%85%e6%a0%b8%e4%b8%8e%e4%b8%9a%e5%8a%a1%e9%80%bb%e8%be%91%e7%9a%84%e6%a1%a5%e6%a2%81%e2%80%94%e2%80%94%e6%8e%a5%e5%8f%a3%e5%b1%82.md">19  深入 MyBatis 内核与业务逻辑的桥梁——接口层.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  插件体系让 MyBatis 世界更加精彩.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/20%20%20%e6%8f%92%e4%bb%b6%e4%bd%93%e7%b3%bb%e8%ae%a9%20MyBatis%20%e4%b8%96%e7%95%8c%e6%9b%b4%e5%8a%a0%e7%b2%be%e5%bd%a9.md">20  插件体系让 MyBatis 世界更加精彩.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  深挖 MyBatis 与 Spring 集成底层原理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/21%20%20%e6%b7%b1%e6%8c%96%20MyBatis%20%e4%b8%8e%20Spring%20%e9%9b%86%e6%88%90%e5%ba%95%e5%b1%82%e5%8e%9f%e7%90%86.md">21  深挖 MyBatis 与 Spring 集成底层原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  基于 MyBatis 的衍生框架一览.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/22%20%20%e5%9f%ba%e4%ba%8e%20MyBatis%20%e7%9a%84%e8%a1%8d%e7%94%9f%e6%a1%86%e6%9e%b6%e4%b8%80%e8%a7%88.md">22  基于 MyBatis 的衍生框架一览.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结束语  会使用只能默默“搬砖”，懂原理才能快速晋升.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90%20MyBatis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86-%e5%ae%8c/23%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bc%9a%e4%bd%bf%e7%94%a8%e5%8f%aa%e8%83%bd%e9%bb%98%e9%bb%98%e2%80%9c%e6%90%ac%e7%a0%96%e2%80%9d%ef%bc%8c%e6%87%82%e5%8e%9f%e7%90%86%e6%89%8d%e8%83%bd%e5%bf%ab%e9%80%9f%e6%99%8b%e5%8d%87.md">23 结束语  会使用只能默默“搬砖”，懂原理才能快速晋升.md</a>
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
                            <h1 id="title" data-id="06  日志框架千千万，MyBatis 都能兼容的秘密是什么？" class="title">06  日志框架千千万，MyBatis 都能兼容的秘密是什么？</h1>
                            <div><p>Apache Commons Logging、Log4j、Log4j2、java.util.logging 等是 Java 开发中常用的几款日志框架，这些日志框架来源于不同的开源组织，给用户暴露的接口也有很多不同之处，所以很多开源框架会自己定义一套统一的日志接口，兼容上述第三方日志框架，供上层使用。</p>

<p>一般实现的方式是使用<strong>适配器模式，将各个第三方日志框架接口转换为框架内部自定义的日志接口</strong>。MyBatis 也提供了类似的实现。</p>

<h3 id="适配器模式">适配器模式</h3>

<p>适配器模式主要解决的是<strong>由于接口不能兼容而导致类无法使用的问题，这在处理遗留代码以及集成第三方框架的时候用得比较多</strong>。其核心原理是：<strong>通过组合的方式，将需要适配的类转换成使用者能够使用的接口</strong>。</p>

<p>适配器模式的类图如下所示：</p>

<p><img src="assets/Cgp9HWAfYoOAKO6lAAEyIgsMVKA161.png" alt="2.png" /></p>

<p>适配器模式类图</p>

<p>在该类图中，你可以看到适配器模式涉及的三个核心角色。</p>

<ul>
<li><strong>目标接口（Target）</strong>：使用者能够直接使用的接口。以处理遗留代码为例，Target 就是最新定义的业务接口。</li>
<li><strong>需要适配的类/要使用的实现类（Adaptee）</strong>：定义了真正要执行的业务逻辑，但是其接口不能被使用者直接使用。这里依然以处理遗留代码为例，Adaptee 就是遗留业务实现，由于编写 Adaptee 的时候还没有定义 Target 接口，所以 Adaptee 无法实现 Target 接口。</li>
<li><strong>适配器（Adapter）</strong>：在实现 Target 接口的同时，维护了一个指向 Adaptee 对象的引用。Adapter 底层会依赖 Adaptee 的逻辑来实现 Target 接口的功能，这样就能够复用 Adaptee 类中的遗留逻辑来完成业务。</li>
</ul>

<p>适配器模式带来的最大好处就是<strong>复用已有的逻辑</strong>，避免直接去修改 Adaptee 实现的接口，这符合开放-封闭原则（也就是程序要对扩展开放、对修改关闭）。</p>

<p>MyBatis 使用的日志接口是自己定义的 Log 接口，但是 Apache Commons Logging、Log4j、Log4j2 等日志框架提供给用户的都是自己的 Logger 接口。为了统一这些第三方日志框架，<strong>MyBatis 使用适配器模式添加了针对不同日志框架的 Adapter 实现</strong>，使得第三方日志框架的 Logger 接口转换成 MyBatis 中的 Log 接口，从而实现集成第三方日志框架打印日志的功能。</p>

<h3 id="日志模块">日志模块</h3>

<p><strong>MyBatis 自定义的 Log 接口位于 org.apache.ibatis.logging 包中，相关的适配器也位于该包中</strong>，下面我们就来看看该模块的具体实现。</p>

<p>首先是 LogFactory 工厂类，它负责创建 Log 对象。这些 Log 接口的实现类中，就包含了多种第三方日志框架的适配器，如下图所示：</p>

<p><img src="assets/CioPOWAfYo6AbKZWAAKqpRwXpuA169.png" alt="3.png" /></p>

<p>Log 接口继承关系图</p>

<p>在 LogFactory 类中有<a href="https://github.com/xxxlxy2008/mybatis/blob/master/src/main/java/org/apache/ibatis/logging/LogFactory.java#L31-L43" target="_blank">一段静态代码块</a>，其中会依次加载各个第三方日志框架的适配器。在静态代码块执行的 tryImplementation() 方法中，首先会检测 logConstructor 字段是否为空，如果不为空，则表示已经成功确定当前使用的日志框架，直接返回；如果为空，则在当前线程中执行传入的 Runnable.run() 方法，尝试确定当前使用的日志框架。</p>

<p>以 JDK Logging 的加载流程（useJdkLogging() 方法）为例，其具体代码实现和注释如下：</p>

<pre><code>public static synchronized void useJdkLogging() {

    setImplementation(org.apache.ibatis.logging.jdk14.Jdk14LoggingImpl.class);

}
private static void setImplementation(Class&lt;? extends Log&gt; implClass) {

  try {

    // 获取implClass这个适配器的构造方法

    Constructor&lt;? extends Log&gt; candidate = implClass.getConstructor(String.class);

    // 尝试加载implClass这个适配器，加载失败会抛出异常

    Log log = candidate.newInstance(LogFactory.class.getName());

    // 加载成功，则更新logConstructor字段，记录适配器的构造方法

    logConstructor = candidate;

  } catch (Throwable t) {

    throw new LogException(&quot;Error setting Log implementation.  Cause: &quot; + t, t);

  }

}
</code></pre>

<p>下面我们以 Jdk14LoggingImpl 为例介绍一下 MyBatis Log 接口的实现。</p>

<p>Jdk14LoggingImpl 作为 Java Logging 的适配器，在实现 MyBatis Log 接口的同时，在内部还封装了一个 java.util.logging.Logger 对象（这是 JDK 提供的日志框架），如下图所示：</p>

<p><img src="assets/CioPOWAeM0WAMQPTAABNRPFy7R0954.png" alt="Drawing 2.png" /></p>

<p>Jdk14LoggingImpl 继承关系图</p>

<p>Jdk14LoggingImpl 对 Log 接口的实现也比较简单，其中会将日志输出操作委托给底层封装的java.util.logging.Logger 对象的相应方法，这与前文介绍的典型适配器模式的实现完全一致。Jdk14LoggingImpl 中的核心实现以及注释如下：</p>

<pre><code>public class Jdk14LoggingImpl implements Log {

  // 指向一个java.util.logging.Logger对象

  private final Logger log;
  public Jdk14LoggingImpl(String clazz) {

    // 初始化log字段

    log = Logger.getLogger(clazz);

  }
  @Override

  public void error(String s, Throwable e) {

    // 全部调用依赖java.util.logging.Logger对象进行实现

    log.log(Level.SEVERE, s, e);

  }

  // 省略其他级别的日志输出方法

}
</code></pre>

<p>在 MyBatis 的 org.apache.ibatis.logging 包下面，除了集成三方日志框架的适配器实现之外，还有一个 jdbc 包，这个包的功能不是将日志写入数据库中，而是将数据库操作涉及的信息通过指定的 Log 打印到日志文件中。我们可以通过这个包，将执行的 SQL 语句、SQL 绑定的参数、SQL 执行之后影响的行数等信息，统统打印到日志中，这个功能主要是在测试环境进行调试的时候使用，很少在线上开启，因为这会产生非常多的日志，拖慢系统性能。</p>

<h3 id="代理模式">代理模式</h3>

<p>在后面即将介绍的 org.apache.ibatis.logging.jdbc 包中，使用到了 JDK 动态代理的相关知识，所以这里我们就先来介绍一下经典的静态代理模式，以及 JDK 提供的动态代理。</p>

<h4 id="1-静态代理模式">1. 静态代理模式</h4>

<p>经典的静态代理模式，其类图如下所示：</p>

<p><img src="assets/Cgp9HWAfYrOAWv7JAAD2hkpzuWw664.png" alt="4.png" /></p>

<p>代理模式类图</p>

<p>从该类图中，你可以看到与代理模式相关的三个核心角色。</p>

<ul>
<li><strong>Subject</strong>：程序中的业务接口，定义了相关的业务方法。</li>
<li><strong>RealSubject</strong>：实现了 Subject 接口的业务实现类，其实现中完成了真正的业务逻辑。</li>
<li><strong>Proxy</strong>：代理类，实现了 Subject 接口，其中会持有一个 Subject 类型的字段，指向一个 RealSubject 对象。</li>
</ul>

<p>在使用的时候，会将 RealSubject 对象封装到 Proxy 对象中，然后访问 Proxy 的相关方法，而不是直接访问 RealSubject 对象。在 Proxy 的方法实现中，不仅会调用 RealSubject 对象的相应方法完成业务逻辑，还会在 RealSubject 方法执行前后进行预处理和后置处理。</p>

<p>通过对代理模式的描述可知，<strong>Proxy 能够控制使用方对 RealSubject 对象的访问，或是在执行业务逻辑之前执行统一的预处理逻辑，在执行业务逻辑之后执行统一的后置处理逻辑</strong>。</p>

<p><strong>代理模式除了实现访问控制以外，还能用于实现延迟加载</strong>。例如，查询数据库涉及网络 I/O 和磁盘 I/O，会是一个比较耗时的操作，有些时候从数据库加载到内存的数据，也并非系统真正会使用到的数据，所以就有了延迟加载这种优化操作。</p>

<p>延迟加载可以有效地避免数据库资源的浪费，其主要原理是：用户在访问数据库时，会立刻拿到一个代理对象，此时并没有执行任何 SQL 到数据库中查询数据，代理对象中自然也不会包含任何真正的有效数据；当用户真正需要使用数据时，会访问代理对象，此时会由代理对象去执行 SQL，完成数据库的查询。MyBatis 也提供了延迟加载功能，原理大同小异，具体的实现方式也是通过代理实现的。</p>

<p>针对每个 RealSubject 类，都需要创建一个 Proxy 代理类，当 RealSubject 这种需要被代理的类变得很多的时候，相应地就需要定义大量的 Proxy 类，这也是经典代理模式面临的一个问题。JDK 动态代理可以有效地解决这个问题，所以接下来我们就来一起分析 JDK 动态代理的核心原理。</p>

<h4 id="2-jdk-动态代理">2. JDK 动态代理</h4>

<p><strong>JDK 动态代理的核心是 InvocationHandler 接口</strong>。这里我先给出了一个 InvocationHandler 的示例实现，如下所示：</p>

<pre><code>public class DemoInvokerHandler implements InvocationHandler {

    private Object target; // 真正的业务对象，也就是RealSubject对象

    // DemoInvokerHandler构造方法

    public DemoInvokerHandler(Object target) { 

        this.target = target;

    }

    public Object invoke(Object proxy, Method method, Object[] args)

             throws Throwable {

        ... // 在执行业务逻辑之前的预处理逻辑

        Object result = method.invoke(target, args);

        ... // 在执行业务逻辑之后的后置处理逻辑

        return result;

    }

    public Object getProxy() {

        // 创建代理对象

        return Proxy.newProxyInstance(Thread.currentThread()

            .getContextClassLoader(),

                target.getClass().getInterfaces(), this);

    }

}
</code></pre>

<p>接下来，我们可以创建一个 main() 方法来模拟使用方创建并使用 DemoInvokerHandler 动态生成代理对象，示例代码如下：</p>

<pre><code>public class Main {

    public static void main(String[] args) {

        Subject subject = new RealSubject();

        DemoInvokerHandler invokerHandler = 

            new DemoInvokerHandler(subject);

        // 获取代理对象

        Subject proxy = (Subject) invokerHandler.getProxy();

        // 调用代理对象的方法，它会调用DemoInvokerHandler.invoke()方法

        proxy.operation();

    }

}
</code></pre>

<p>现在假设有多个业务逻辑类，需要相同的预处理逻辑和后置处理逻辑，那么只需要提供一个 InvocationHandler 接口实现类即可。<strong>在程序运行过程中，JDK 动态代理会为每个业务类动态生成相应的代理类实现</strong>，并加载到 JVM 中，然后创建对应的代理实例对象。</p>

<p>下面我们就接着来深入分析一下 JDK 动态代理底层动态创建代理类的原理。不同 JDK 版本 Proxy 类的实现会有些许差异，但总体的核心思路基本一致，这里我们就以 JDK 1.8.0 版本为例进行说明。</p>

<p>首先，从前面的示例代码中可以看出，JDK 动态代理的入口方法是 Proxy.newProxyInstance()，这个静态方法有以下三个参数。</p>

<ul>
<li>loader（ClassLoader 类型）：加载动态生成的代理类的类加载器。</li>
<li>interfaces（Class[] 类型）：业务类实现的接口。</li>
<li>h（InvocationHandler 类型）：自定义的 InvocationHandler 对象。</li>
</ul>

<p>下面进入 Proxy.newProxyInstance() 方法，查看其具体实现如下：</p>

<pre><code>public static Object newProxyInstance(ClassLoader loader,

     Class[] interfaces, InvocationHandler h) 

         throws IllegalArgumentException {

    final Class&lt;?&gt;[] intfs = interfaces.clone();

    ... // 省略权限检查等代码

    Class&lt;?&gt; cl = getProxyClass0(loader, intfs);  // 获取代理类

    ... // 省略try/catch代码块和相关异常处理

    // 获取代理类的构造方法

    final Constructor&lt;?&gt; cons = cl.getConstructor(constructorParams);

    final InvocationHandler ih = h;

    return cons.newInstance(new Object[]{h});  // 创建代理对象

}
</code></pre>

<p>从 newProxyInstance() 方法的具体实现代码中我们可以看到，JDK 动态代理是在 getProxyClass0() 方法中完成代理类的生成和加载。getProxyClass0() 方法的具体实现如下：</p>

<pre><code>private static Class getProxyClass0 (ClassLoader loader, 

        Class... interfaces) {

    // 边界检查，限制接口数量（略）

    // 如果指定的类加载器中已经创建了实现指定接口的代理类，则查找缓存；

    // 否则通过ProxyClassFactory创建实现指定接口的代理类

    return proxyClassCache.get(loader, interfaces);

}
</code></pre>

<p>proxyClassCache 是定义在 Proxy 类中一个静态字段，它是 WeakCache 类型的集合，用于缓存已经创建过的代理类，具体定义如下：</p>

<pre><code>private static final WeakCache&lt;ClassLoader, Class&lt;?&gt;[], Class&lt;?&gt;&gt; proxyClassCache

     = new WeakCache&lt;&gt;(new KeyFactory(), 

           new ProxyClassFactory());
</code></pre>

<p>WeakCache.get() 方法会首先尝试从缓存中查找代理类，如果查找失败，则会创建相应的 Factory 对象并调用其 get() 方法获取代理类。Factory 是 WeakCache 中的内部类，在 Factory.get() 方法中会通过 ProxyClassFactory.apply() 方法创建并加载代理类。</p>

<p>在 ProxyClassFactory.apply() 方法中，首先会检测代理类需要实现的接口集合，然后确定代理类的名称，之后创建代理类并将其写入文件中，最后加载代理类，返回对应的 Class 对象用于后续的实例化代理类对象。该方法的具体实现如下：</p>

<pre><code>public Class apply(ClassLoader loader, Class[] interfaces) {

    // ... 对interfaces集合进行一系列检测（略）

    // ... 选择定义代理类的包名（略）

    // 代理类的名称是通过包名、代理类名称前缀以及编号这三项组成的

    long num = nextUniqueNumber.getAndIncrement();

    String proxyName = proxyPkg + proxyClassNamePrefix + num;

    // 生成代理类，并写入文件

    byte[] proxyClassFile = ProxyGenerator.generateProxyClass(

            proxyName, interfaces, accessFlags);

    // 加载代理类，并返回Class对象

    return defineClass0(loader, proxyName, proxyClassFile, 0, 

      proxyClassFile.length);

}
</code></pre>

<p>ProxyGenerator.generateProxyClass() 方法会按照指定的名称和接口集合生成代理类的字节码，并根据条件决定是否保存到磁盘上。该方法的具体代码如下：</p>

<pre><code>public static byte[] generateProxyClass(final String name,

       Class[] interfaces) {

    ProxyGenerator gen = new ProxyGenerator(name, interfaces);

    // 动态生成代理类的字节码，具体生成过程不再详细介绍

    final byte[] classFile = gen.generateClassFile();

    // 如果saveGeneratedFiles值为true，会将生成的代理类的字节码保存到文件中

    if (saveGeneratedFiles) { 

        java.security.AccessController.doPrivileged(

            new java.security.PrivilegedAction() {

                public Void run() {

                    // 省略try/catch代码块

                    FileOutputStream file = new FileOutputStream(

                        dotToSlash(name) + &quot;.class&quot;);

                    file.write(classFile);

                    file.close();

                    return null;

                }

            }

        );

    }

    return classFile; // 返回上面生成的代理类的字节码

}
</code></pre>

<p>最后，为了清晰地看到 JDK 动态生成的代理类的真正代码，我们需要将上述生成的代理类的字节码进行反编译。上述示例为 RealSubject 生成的代理类，反编译后得到的代码如下：</p>

<pre><code>public final class $Proxy143 

      extends Proxy implements Subject {  // 实现了Subject接口

    // 这里省略了从Object类继承下来的相关方法和属性

    private static Method m3;

    static {

        // 省略了try/catch代码块

        // 记录了operation()方法对应的Method对象

        m3 = Class.forName(&quot;design.proxy.Subject&quot;)

          .getMethod(&quot;operation&quot;, new Class[0]);

    }

    // 构造方法的参数就是我们在示例中使用的DemoInvokerHandler对象

    public $Proxy11(InvocationHandler var1) throws {

        super(var1); 

    }

    public final void operation() throws {

        // 省略了try/catch代码块

        // 调用DemoInvokerHandler对象的invoke()方法

        // 最终调用RealSubject对象的对应方法

        super.h.invoke(this, m3, (Object[]) null);

    }

}
</code></pre>

<p>到此为止，JDK 动态代理的基本使用以及核心原理就分析完了。这里我做一个简单的总结，JDK 动态代理的实现原理是：<strong>动态创建代理类，然后通过指定类加载器进行加载</strong>。在创建代理对象时，需要将 InvocationHandler 对象作为构造参数传入；当调用代理对象时，会调用 InvocationHandler.invoke() 方法，从而执行代理逻辑，最终调用真正业务对象的相应方法。</p>

<h3 id="jdbc-logger">JDBC Logger</h3>

<p>了解了代理模式以及 JDK 动态代理的基础知识之后，下面我们开始分析 org.apache.ibatis.logging.jdbc 包中的内容。</p>

<p>首先来看其中<strong>最基础的抽象类—— BaseJdbcLogger，它是 jdbc 包下其他 Logger 类的父类</strong>，继承关系如下图所示：</p>

<p><img src="assets/CioPOWAfYsuAc9WjAAFm7izVaMI477.png" alt="1.png" /></p>

<p>BaseJdbcLogger 继承关系图</p>

<p>在 BaseJdbcLogger 这个抽象类中，定义了 SET_METHODS 和 EXECUTE_METHODS 两个 Set 类型的集合。其中，SET_METHODS 用于记录绑定 SQL 参数涉及的全部 set*() 方法名称，例如 setString() 方法、setInt() 方法等。EXECUTE_METHODS 用于记录执行 SQL 语句涉及的所有方法名称，例如 execute() 方法、executeUpdate() 方法、executeQuery() 方法、addBatch() 方法等。这两个集合都是在 BaseJdbcLogger 的静态代码块中被填充的。</p>

<p>从上面的 BaseJdbcLogger 继承关系图中可以看到，BaseJdbcLogger 的子类同时会实现 InvocationHandler 接口。</p>

<p>我们先来看其中的 ConnectionLogger 实现，其底层维护了一个 Connection 对象的引用，在<a href="https://github.com/xxxlxy2008/mybatis/blob/master/src/main/java/org/apache/ibatis/logging/jdbc/ConnectionLogger.java#L84-L90" target="_blank">ConnectionLogger.newInstance() 方法</a>中会使用 JDK 动态代理的方式为这个 Connection 对象创建相应的代理对象。</p>

<p>invoke() 方法是代理对象的核心方法，在该方法中，ConnectionLogger 会为 prepareStatement()、prepareCall()、createStatement() 三个方法添加代理逻辑。下面来看 invoke() 方法的具体实现，具体代码以及注释如下：</p>

<pre><code>public Object invoke(Object proxy, Method method, Object[] params)

        throws Throwable {

    try {

        if (Object.class.equals(method.getDeclaringClass())) {

            // 如果调用的是从Object继承的方法，则直接调用，不做任何拦截

            return method.invoke(this, params);

        }

        // 调用prepareStatement()方法、prepareCall()方法的时候，

        // 会在创建PreparedStatement对象之后，用PreparedStatementLogger为其创建代理对象

        if (&quot;prepareStatement&quot;.equals(method.getName()) || &quot;prepareCall&quot;.equals(method.getName())) {

            if (isDebugEnabled()) {

                // 通过statementLog这个Log输出日志

                debug(&quot; Preparing: &quot; + removeExtraWhitespace((String) params[0]), true);

            }

            PreparedStatement stmt = (PreparedStatement) method.invoke(connection, params);

            stmt = PreparedStatementLogger.newInstance(stmt, statementLog, queryStack);

            return stmt;

        } else if (&quot;createStatement&quot;.equals(method.getName())) {

            // 调用createStatement()方法的时候，

            // 会在创建Statement对象之后，用StatementLogger为其创建代理对象

            Statement stmt = (Statement) method.invoke(connection, params);

            stmt = StatementLogger.newInstance(stmt, statementLog, queryStack);

            return stmt;

        } else {

            // 除了上述三个方法之外，其他方法的调用将直接传递给底层Connection对象的相应方法处理

            return method.invoke(connection, params);

        }

    } catch (Throwable t) {

        throw ExceptionUtil.unwrapThrowable(t);

    }

}
</code></pre>

<p>下面我们来看 PreparedStatementLogger，在其 invoke() 方法中调用了 SET_METHODS 集合中的方法、EXECUTE_METHODS 集合中的方法或 getResultSet() 方法时，会添加相应的代理逻辑。StatementLogger 中的 Invoke() 方法实现与之类似，这里就不再赘述。</p>

<p>最后我们再看下 ResultSetLogger 对 InvocationHandler 接口的实现，其中会针对 ResultSet.next() 方法进行后置处理，主要是打印结果集中每一行数据以及统计结果集总行数等信息，具体实现和注释如下：</p>

<pre><code>public Object invoke(Object proxy, Method method, Object[] params) throws Throwable {

  try {

      if (Object.class.equals(method.getDeclaringClass())) {

          // 如果调用Object的方法，则直接调用，不做任何其他处理

          return method.invoke(this, params);

      }

      Object o = method.invoke(rs, params);

      // 针对ResultSet.next()方法进行后置处理

      if (&quot;next&quot;.equals(method.getName())) { 

          if ((Boolean) o) { // 检测next()方法的返回值，确定是否还存在下一行数据

              rows++; // 记录ResultSet中的行数

              if (isTraceEnabled()) {

                  // 获取数据集的列元数据

                  ResultSetMetaData rsmd = rs.getMetaData();

                  // 获取数据集的列数

                  final int columnCount = rsmd.getColumnCount();

                  if (first) { // 如果是数据集的第一行数据，会输出表头信息

                      first = false;

                      // 这里除了输出表头，还会记录BLOB等超大类型的列名

                      printColumnHeaders(rsmd, columnCount);

                  }

                  // 输出当前遍历的这行记录，这里会过滤掉超大类型列的数据，不进行输出

                  printColumnValues(columnCount);

              }

          } else { // 完成结果集的遍历之后，这里会在日志中输出总行数

              debug(&quot;     Total: &quot; + rows, false);

          }

      }

      clearColumnInfo(); // 清空column*集合

      return o;

  } catch (Throwable t) {

      throw ExceptionUtil.unwrapThrowable(t);

  }

}
</code></pre>

<h3 id="总结">总结</h3>

<p>在这一讲中，我们主要介绍的是 MyBatis 基础模块中的日志模块。</p>

<ul>
<li>首先，介绍了适配器模式的核心知识点，这也是日志模块底层的设计思想。</li>
<li>然后，说明了日志模块是如何基于适配器模式集成多种三方日志框架的。</li>
<li>接下来，详细讲解了静态代理模式以及 JDK 动态代理的实现原理。</li>
<li>最后，深入分析了 JDBC Logger 是如何基于 JDK 动态代理实现日志功能的。</li>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#214d4d4d18151010111661464c40484d0f424e4c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f158abe7bb37771',t:'MTczNDA4ODU4Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
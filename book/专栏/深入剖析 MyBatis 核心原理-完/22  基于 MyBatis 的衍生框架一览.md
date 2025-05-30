<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=22&#32;&#32;基于&#32;MyBatis&#32;的衍生框架一览>
        <link rel="icon" href="/static/favicon.png">
        <title>22  基于 MyBatis 的衍生框架一览 </title>
        
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
                            <h1 id="title" data-id="22  基于 MyBatis 的衍生框架一览" class="title">22  基于 MyBatis 的衍生框架一览</h1>
                            <div><p>在前面的课时中，我们深入分析了 MyBatis 的内核，了解了 MyBatis 处理一条 SQL 的完整流程，剖析了 MyBatis 中动态 SQL、结果集映射、缓存等核心功能的实现原理。在日常工作中，除了单纯使用 MyBatis 之外，还可能会涉及 MyBatis 的衍生框架，这一讲我们就来介绍一下工作中常用的 MyBatis 衍生框架。</p>

<h3 id="mybatis-generator">MyBatis-Generator</h3>

<p>虽然使用 MyBatis 编写 DAO 层已经非常方便，但是我们还是要编写 Mapper 接口和相应的 Mapper.xml 配置文件。为了进一步节省编码时间，我们<strong>可以选择 MyBatis-Generator 工具自动生成 Mapper 接口和 Mapper.xml 配置文件</strong>。</p>

<p>这里我们通过一个简单示例介绍一下 MyBatis-Generator 工具的基本功能。</p>

<p>MyBatis-Generator 目前最新的版本是 1.4.0 版本，首先我们需要下载<a href="https://github.com/mybatis/generator/releases/download/mybatis-generator-1.4.0/mybatis-generator-core-1.4.0-bundle.zip?fileGuid=xxQTRXtVcqtHK6j8" target="_blank">这个</a>最新的 zip 包，并进行解压，得到 mybatis-generator-core-1.4.0.jar 这个 jar 包。</p>

<p>由于我们本地使用的是 MySQL 数据库，所以需要准备一个 mysql-connector-java 的 jar 包，我们可以从本地的 Maven 仓库中获得，具体的目录是：.m2/repository/mysql/mysql-connector-java/，在这个目录中选择一个最新版本的 jar 包拷贝到 mybatis-generator-core-1.4.0.jar 同目录下。</p>

<p>接下来，我们需要编写一个 generatorConfig.xml 配置文件，其中会告诉 MyBatis-Generator 去连接哪个数据库、连接数据库的用户名和密码分别是什么、需要根据哪些表生成哪些配置文件和类，以及这些生成文件的存放位置。下面是一个 generatorConfig.xml 配置文件的完整示例：</p>

<pre><code>&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot;?&gt;

&lt;!DOCTYPE generatorConfiguration

        PUBLIC &quot;-//mybatis.org//DTD MyBatis Generator Configuration 1.0//EN&quot;

        &quot;http://mybatis.org/dtd/mybatis-generator-config_1_0.dtd&quot;&gt;

&lt;generatorConfiguration&gt;

    &lt;!-- 使用的数据库驱动jar包 --&gt;

    &lt;classPathEntry location=&quot;mysql-connector-java-8.0.22.jar&quot;/&gt;

    &lt;!-- 指定数据库地址、数据库用户名和密码 --&gt;

    &lt;context id=&quot;DB2Tables&quot; targetRuntime=&quot;MyBatis3&quot;&gt;

        &lt;jdbcConnection driverClass=&quot;com.mysql.jdbc.Driver&quot;

                connectionURL=&quot;jdbc:mysql://localhost:3306/test&quot;

                userId=&quot;root&quot;  password=&quot;xxx&quot;&gt;

        &lt;/jdbcConnection&gt;

        &lt;javaTypeResolver&gt;

            &lt;property name=&quot;forceBigDecimals&quot; value=&quot;false&quot;/&gt;

        &lt;/javaTypeResolver&gt;

        &lt;!-- 生成的Model类存放位置 --&gt;

        &lt;javaModelGenerator targetPackage=&quot;org.example&quot; targetProject=&quot;src&quot;&gt;

            &lt;!-- 是否支持生成子package --&gt;

            &lt;property name=&quot;enableSubPackages&quot; value=&quot;true&quot;/&gt;

            &lt;!-- 对String进行操作时，会添加trim()方法进行处理 --&gt;

            &lt;property name=&quot;trimStrings&quot; value=&quot;true&quot;/&gt;

        &lt;/javaModelGenerator&gt;

        &lt;!-- 生成的Mapper.xml映射配置文件的存放位置--&gt;

        &lt;sqlMapGenerator targetPackage=&quot;org.example.mapper&quot; targetProject=&quot;src&quot;&gt;

            &lt;property name=&quot;enableSubPackages&quot; value=&quot;true&quot;/&gt;

        &lt;/sqlMapGenerator&gt;

        &lt;!-- 生成的Mapper接口的存放位置--&gt;

        &lt;javaClientGenerator type=&quot;XMLMAPPER&quot; targetPackage=&quot;org.example.mapper&quot; 

                     targetProject=&quot;src&quot;&gt;

            &lt;property name=&quot;enableSubPackages&quot; value=&quot;true&quot;/&gt;

        &lt;/javaClientGenerator&gt;

        &lt;!-- 数据库表与Model类之间的映射关系，根据t_customer表进行映射--&gt;

        &lt;table schema=&quot;test&quot; tableName=&quot;t_customer&quot; domainObjectName=&quot;Customer&quot;

               enableCountByExample=&quot;false&quot; enableUpdateByExample=&quot;false&quot; 

               enableDeleteByExample=&quot;false&quot;

               enableSelectByExample=&quot;false&quot; selectByExampleQueryId=&quot;false&quot;&gt;

        &lt;/table&gt;

    &lt;/context&gt;

&lt;/generatorConfiguration&gt;
</code></pre>

<p>然后，我们准备一下数据库中的表，在 MySQL 中建立一个 test 数据库，并创建 t_customer 表，使用到的建库建表语句如下：</p>

<pre><code>create databases test; # 创建数据库

use test;

DROP TABLE IF EXISTS `t_customer`; # 删除已有的t_customer表

CREATE TABLE `t_customer` ( # 创建t_customer表

  `id` int(255) NOT NULL,

  `name` varchar(255) DEFAULT NULL,

  `password` varchar(255) DEFAULT NULL,

  `account` bigint(255) DEFAULT NULL,

  PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
</code></pre>

<p>最后，我们在 mybatis-generator-core-1.4.0.jar 包同目录下新建一个 src 目录，存放生成的代码，然后执行如下命令，逆向生成需要的代码：</p>

<pre><code>java -jar mybatis-generator-core-1.4.0.jar -configfile generatorConfig.xml
</code></pre>

<p>命令正常执行完成之后，可以看到 src 目录下生成的文件如下图所示：</p>

<p><img src="assets/CioPOWBtTDqAYagkAABmeFv2Z84519.png" alt="Drawing 0.png" /></p>

<p>MyBatis-Generator 工具类生成结果图</p>

<p>生成的 Customer.java 类是一个 Model 类（或者说 Domain 类），包含了 id、name、password、account 属性；CustomerMapper.xml 是 Customer 对应的 Mapper.xml 配置文件，其中定义了按照 id 进行查询和删除的 select、delete 语句，以及全字段写入和更新的 insert、update 语句；CustomerMapper 接口中包含了与 CustomerMapper.xml 对应的方法。该示例中生成的代码并不复杂，在你生成代码之后，也希望你能够自己分析一下。</p>

<h3 id="mybatis-分页插件">MyBatis 分页插件</h3>

<p>MyBatis 本身提供了 RowBounds 参数，可以实现分页的效果，但是在前面[第 14 讲]中我们提到过，通过 RowBounds 方式实现分页的时候，本质是将整个结果集数据加载到内存中，然后在内存中过滤出需要的数据，这其实也是我们常说的“内存分页”。而真正的分页是为了解决数据量太大，无法直接加载到内存或无法直接传输的问题，显然“内存分页”并没有解决这个问题。</p>

<p>你如果用过 MySQL 的话，应该知道我们常用 limit 方式进行分页，例如下面这条 select 语句：</p>

<pre><code>select * from t_customer limit 5,10;
</code></pre>

<p>使用 Oracle 实现分页时，则需要用 rownum 实现，可见在不同数据库中实现物理分页的写法各不相同。</p>

<p>如果我们想屏蔽底层数据库的分页 SQL 语句的差异，同时使用 MyBatis 的 RowBounds 参数实现“物理分页”，可以考虑使用 MyBatis 的分页插件<a href="https://github.com/pagehelper/Mybatis-PageHelper?fileGuid=xxQTRXtVcqtHK6j8" target="_blank">PageHelper</a>。PageHelper 的使用比较简单，只需要在 pom.xml 中引入 PageHelper 依赖包，并在 mybatis-config.xml 配置文件中配置 PageInterceptor 插件即可，核心配置如下：</p>

<pre><code>&lt;plugins&gt;

    &lt;plugin interceptor=&quot;com.github.pagehelper.PageInterceptor&quot;&gt;

        &lt;property name=&quot;helperDialect&quot; value=&quot;mysql&quot;/&gt;

	&lt;/plugin&gt;

&lt;/plugins&gt;
</code></pre>

<p><strong>PageHelper 核心原理是使用 MyBatis 的插件机制，整个插件的入口是在 PageInterceptor</strong>。</p>

<p>在 PageInterceptor 初始化的时候，会根据配置的 helperDialect 属性以及 MyBatis 使用的 JDBC URL 信息确定底层连接的数据库类型，并创建一个 Dialect 对象。我们可以再来看 PageInterceptor 的注解信息，会发现 PageInterceptor 会拦截 Executor 中带有 RowBounds 参数的两个查询方法。拦截到目标方法之后，PageInterceptor.intercept() 方法会通过 Dialect 对象完成分页操作，核心代码如下：</p>

<pre><code>List resultList;

// 判断是否需要进行分页

if (!dialect.skip(ms, parameter, rowBounds)) {

    // 是否需要查询总记录数，这可以帮助我们显示总页数

    if (dialect.beforeCount(ms, parameter, rowBounds)) {

        // 查询总记录数

        Long count = count(executor, ms, parameter, rowBounds, null, boundSql);

        // 处理查询总记录数，返回true时继续分页查询，false时直接返回，会返回false的原因很多，可能是count为0，或是当前已经到最后一页等原因

        if (!dialect.afterCount(count, parameter, rowBounds)) {

            return dialect.afterPage(new ArrayList(), parameter, rowBounds);

        }

    }

    // 执行分页查询

    resultList = ExecutorUtil.pageQuery(dialect, executor,

            ms, parameter, rowBounds, resultHandler, boundSql, cacheKey);

} else {

    // 如果不需要，直接交给Executor执行查询，返回结果

    resultList = executor.query(ms, parameter, rowBounds, resultHandler, cacheKey, boundSql);

}

// 在afterPage()方法中会完成总页数的计算等后置操作

return dialect.afterPage(resultList, parameter, rowBounds);
</code></pre>

<p>通过对 PageInterceptor 的分析我们看到，<strong>核心的分页逻辑都是在 Dialect 中完成的</strong>，PageHelper 针对每个数据库都提供了一个 Dialect 接口实现。下图展示了 MySQL 数据库对应的 Dialect 接口实现：</p>

<p><img src="assets/Cgp9HWBtTFKAVlWCAACyAbYHCQg938.png" alt="Drawing 1.png" /></p>

<p>MySqlDialect 的继承关系图</p>

<p>在上图中，PageHelper 是一个通用的 Dialect 实现，会将上述分页操作委托给当前线程绑定的 Dialect 实现进行处理，这主要是靠其中的 autoDialect 字段（PageAutoDialect 类型）实现的。AbstractDialect 中只提供了一个生成“查询总记录数”SQL 语句（即 select count(*) 语句）的功能。</p>

<p>AbstractRowBoundsDialect 这条继承线是针对 RowBounds 进行分页的 Dialect 实现，其中会根据 RowBounds 实现 Dialect 接口，例如，在 MySqlRowBoundsDialect 中的 getPageSql() 方法实现中会改写 SQL 语句，添加 limit 子句，其中的 offset、limit 参数均来自传入的 RowBounds 参数。</p>

<p>如果没有用 RowBounds 参数进行分页，而是在传入的 SQL 语句绑定实参（即 Executor.query() 方法的第二个参数 parameter）中指定 pageNum、pageSize 等分页信息，则会走 AbstractHelperDialect 这条继承线。在 PageObjectUtil 这个工具类中，会从绑定实参中解析出分页信息并封装成 Page 对象，然后传递给 AbstractHelperDialect 完成分页操作。例如，在 MySqlDialect 实现中的 getPageSql() 方法和 processPageParameter() 方法，都会从 Page 参数中获取分页信息，这两个方法的具体实现就留给你自己分析了。</p>

<p>到此为止，PageHelper 分页插件中的分页功能就介绍完了，除了基本的分页功能，PageHelper 还提供了分页使用的缓存等相关能力，这里就不再展开详细分析了，你若感兴趣的话可以下载其源码进行深入分析。</p>

<h3 id="mybatis-plus">MyBatis-Plus</h3>

<p>MyBatis-Plus 是国人开发的一款 MyBatis 增强工具，通过其名字就能看出，<strong>它并没有改变 MyBatis 本身的功能，而是在 MyBatis 的基础上提供了很多增强功能，使我们的开发更加简洁高效</strong>。也正是由于其“只做增强不做改变”的特性，让我们可以在使用 MyBatis 的项目中无感知地引入 MyBatis-Plus。</p>

<p>MyBatis-Plus 对 MyBatis 的很多方面进行了增强，例如：</p>

<ul>
<li>内置了通用的 Mapper 和通用的 Service，只需要添加少量配置即可实现 DAO 层和 Service 层；</li>
<li>内置了一个分布式唯一 ID 生成器，可以提供分布式环境下的 ID 生成策略；</li>
<li>通过 Maven 插件可以集成生成代码能力，可以快速生成 Mapper、Service 以及 Controller 层的代码，同时支持模块引擎的生成；</li>
<li>内置了分页插件，可以实现和 PageHelper 类似的“物理分页”，而且分页插件支持多种数据库；</li>
<li>内置了一款性能分析插件，通过该插件我们可以获取一条 SQL 语句的执行时间，可以更快地帮助我们发现慢查询。</li>
</ul>

<p>既然 MyBatis-Plus 在 MyBatis 之上提供了这么多的扩展，那么我们就来快速上手体验一下 MyBatis-Plus。这里我们依旧选用 MySQL 数据库，复用上面介绍 MyBatis-Generator 示例时用到的 test 库和 t_customer 表。</p>

<p>首先，新建一个 Spring Boot 项目，这里我们可以使用 Spring 官网提供的<a href="https://start.spring.io/?fileGuid=xxQTRXtVcqtHK6j8" target="_blank">项目生成器</a>快速生成，导入 IDEA 之后会发现 Spring Boot 的配置和启动类都已经生成好了，如下图所示：</p>

<p><img src="assets/Cgp9HWBtTGCAB50qAADaNi9sMew051.png" alt="Drawing 2.png" /></p>

<p>Spring Boot 示例项目的结构图</p>

<p>接下来我们打开 pom.xml 文件，看到其中已经自动添加了 Spring Boot 的全部依赖，此时只需要添加 mysql-connector-java 依赖以及 MyBatis-Plus 依赖即可（目前 MyBatis-Plus 最新版本是 3.4.2）：</p>

<pre><code>&lt;dependency&gt;

    &lt;groupId&gt;com.baomidou&lt;/groupId&gt;

    &lt;artifactId&gt;mybatis-plus-boot-starter&lt;/artifactId&gt;

    &lt;version&gt;3.4.2&lt;/version&gt;

&lt;/dependency&gt;

&lt;dependency&gt;

    &lt;groupId&gt;mysql&lt;/groupId&gt;

    &lt;artifactId&gt;mysql-connector-java&lt;/artifactId&gt;

&lt;/dependency&gt;
</code></pre>

<p>再接下来，我们修改 application.properties 文件，添加数据库的相关配置：</p>

<pre><code>spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver

spring.datasource.url=jdbc:mysql://localhost:3306/test?useUnicode=true&amp;characterEncoding=UTF-8

spring.datasource.username=root

spring.datasource.password=xxx
</code></pre>

<p>然后，我们开始编写 Customer 类和 CustomerMapper 接口，这两个类非常简单，Customer 类中需要定义 t_customer 表中各列对应的属性，如下所示：</p>

<pre><code>@TableName(value = &quot;t_customer&quot;) // 通过@TableName注解，指定Customer与 t_customer表的关联关系

public class Customer {

    private Integer id;

    private String name;

    private String password;

    private Long account;

    // 省略上述字段的getter/setter方法，以及toString()方法

}
</code></pre>

<p>CustomerMapper 接口的定义更加简单，只需要继承 BaseMapper 即可，具体定义如下：</p>

<pre><code>public interface CustomerMapper extends BaseMapper&lt;Customer&gt; {

  // 无须提供任何方法定义，而是从BaseMapper继承

}
</code></pre>

<p>最后，我们修改一下这个 Spring Boot 项目的启动类 DemoApplication，在其中添加 @MapperScan 注解指定 Mapper 接口所在的包，该注解会自动进行扫描，DemoApplication 的具体实现如下：</p>

<pre><code>@SpringBootApplication

@MapperScan(&quot;com.example.demo.mapper&quot;)

public class DemoApplication {

    public static void main(String[] args) {

        SpringApplication.run(DemoApplication.class, args);

    }

}
</code></pre>

<p>完成上述示例的编写之后，我们可以添加一个测试用例来查询 t_customer 表中的数据，具体实现如下：</p>

<pre><code>@RunWith(SpringRunner.class)

@SpringBootTest

class DemoApplicationTests {

    @Autowired

    private CustomerMapper customerMapper;

    @Test

    public void testSelect() {

        Customer customer = new Customer();

        customer.setId(1);

        customer.setName(&quot;Bob&quot;);

        customer.setPassword(&quot;pwd&quot;);

        customer.setAccount(10097L);

        int insert = customerMapper.insert(customer);

        System.out.println(&quot;affect row num:&quot; + insert);

        List&lt;Customer&gt; userList = customerMapper.selectList(null);

        userList.forEach(System.out::println);

    }

}
</code></pre>

<p>执行该单元测试之后，得到如下输出：</p>

<pre><code>affect row num:1

Customer{id=1, name='Bob', password='pwd', account=10097}
</code></pre>

<p>MyBatis-Plus 的基础使用示例就介绍到这里了。另外，MyBatis-Plus<a href="https://baomidou.com/guide/#特性?fileGuid=xxQTRXtVcqtHK6j8" target="_blank">官方文档</a>中还提供了很多核心功能的说明和介绍，同时 MyBatis-Plus 还提供了<a href="https://github.com/baomidou/mybatis-plus-samples?fileGuid=xxQTRXtVcqtHK6j8" target="_blank">示例 GitHub 仓库</a>，其中包含了非常多的 MyBatis-Plus 示例代码和使用技巧，非常值得你参考。</p>

<h3 id="总结">总结</h3>

<p>在这一讲我们重点介绍了 MyBatis 相关的辅助工具以及在 MyBatis 之上衍生出来的扩展框架。</p>

<ul>
<li>首先，分析了 MyBatis-Generator 工具，它可以根据我们已有的数据表快速生成 MyBatis 中的 Domain 类、Mapper 接口以及 Mapper.xml 文件。</li>
<li>然后，介绍了 MyBatis 分页插件—— PageHelper，PageHelper 可以让我们直接使用 RowBounds API 实现“内存分页”，同时也可以帮助我们实现对不同数据库产品的分页功能。</li>
<li>最后，还讲解了 MyBatis-Plus 框架，MyBatis-Plus 内置了默认的 DAO 和 Service 实现以及分页功能，可以大幅度提高开发效率，你也可以结合我展示的示例来帮助你快速上手 MyBatis-Plus 框架。</li>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#cfa3a3a3f6fbfefefff88fa8a2aea6a3e1aca0a2" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f158e90ea087771',t:'MTczNDA4ODc0My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=28&#32;如何在原生应用中混编Flutter工程？>
        <link rel="icon" href="/static/favicon.png">
        <title>28 如何在原生应用中混编Flutter工程？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么每一位大前端从业者都应该学习Flutter？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%af%8f%e4%b8%80%e4%bd%8d%e5%a4%a7%e5%89%8d%e7%ab%af%e4%bb%8e%e4%b8%9a%e8%80%85%e9%83%bd%e5%ba%94%e8%af%a5%e5%ad%a6%e4%b9%a0Flutter%ef%bc%9f.md">00 开篇词 为什么每一位大前端从业者都应该学习Flutter？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 预习篇 · 从0开始搭建Flutter工程环境.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20%e4%bb%8e0%e5%bc%80%e5%a7%8b%e6%90%ad%e5%bb%baFlutter%e5%b7%a5%e7%a8%8b%e7%8e%af%e5%a2%83.md">01 预习篇 · 从0开始搭建Flutter工程环境.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 预习篇 · Dart语言概览.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/02%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20Dart%e8%af%ad%e8%a8%80%e6%a6%82%e8%a7%88.md">02 预习篇 · Dart语言概览.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 深入理解跨平台方案的历史发展逻辑.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/03%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e8%b7%a8%e5%b9%b3%e5%8f%b0%e6%96%b9%e6%a1%88%e7%9a%84%e5%8e%86%e5%8f%b2%e5%8f%91%e5%b1%95%e9%80%bb%e8%be%91.md">03 深入理解跨平台方案的历史发展逻辑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 Flutter区别于其他方案的关键技术是什么？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/04%20Flutter%e5%8c%ba%e5%88%ab%e4%ba%8e%e5%85%b6%e4%bb%96%e6%96%b9%e6%a1%88%e7%9a%84%e5%85%b3%e9%94%ae%e6%8a%80%e6%9c%af%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">04 Flutter区别于其他方案的关键技术是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 从标准模板入手，体会Flutter代码是如何运行在原生系统上的.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%e4%bb%8e%e6%a0%87%e5%87%86%e6%a8%a1%e6%9d%bf%e5%85%a5%e6%89%8b%ef%bc%8c%e4%bd%93%e4%bc%9aFlutter%e4%bb%a3%e7%a0%81%e6%98%af%e5%a6%82%e4%bd%95%e8%bf%90%e8%a1%8c%e5%9c%a8%e5%8e%9f%e7%94%9f%e7%b3%bb%e7%bb%9f%e4%b8%8a%e7%9a%84.md">05 从标准模板入手，体会Flutter代码是如何运行在原生系统上的.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 基础语法与类型变量：Dart是如何表示信息的？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/06%20%e5%9f%ba%e7%a1%80%e8%af%ad%e6%b3%95%e4%b8%8e%e7%b1%bb%e5%9e%8b%e5%8f%98%e9%87%8f%ef%bc%9aDart%e6%98%af%e5%a6%82%e4%bd%95%e8%a1%a8%e7%a4%ba%e4%bf%a1%e6%81%af%e7%9a%84%ef%bc%9f.md">06 基础语法与类型变量：Dart是如何表示信息的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 函数、类与运算符：Dart是如何处理信息的？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%e5%87%bd%e6%95%b0%e3%80%81%e7%b1%bb%e4%b8%8e%e8%bf%90%e7%ae%97%e7%ac%a6%ef%bc%9aDart%e6%98%af%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e4%bf%a1%e6%81%af%e7%9a%84%ef%bc%9f.md">07 函数、类与运算符：Dart是如何处理信息的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 综合案例：掌握Dart核心特性.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%e7%bb%bc%e5%90%88%e6%a1%88%e4%be%8b%ef%bc%9a%e6%8e%8c%e6%8f%a1Dart%e6%a0%b8%e5%bf%83%e7%89%b9%e6%80%a7.md">08 综合案例：掌握Dart核心特性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 Widget，构建Flutter界面的基石.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/09%20Widget%ef%bc%8c%e6%9e%84%e5%bb%baFlutter%e7%95%8c%e9%9d%a2%e7%9a%84%e5%9f%ba%e7%9f%b3.md">09 Widget，构建Flutter界面的基石.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Widget中的State到底是什么？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/10%20Widget%e4%b8%ad%e7%9a%84State%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">10 Widget中的State到底是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 提到生命周期，我们是在说什么？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%e6%8f%90%e5%88%b0%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%ef%bc%8c%e6%88%91%e4%bb%ac%e6%98%af%e5%9c%a8%e8%af%b4%e4%bb%80%e4%b9%88%ef%bc%9f.md">11 提到生命周期，我们是在说什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 经典控件（一）：文本、图片和按钮在Flutter中怎么用？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%e7%bb%8f%e5%85%b8%e6%8e%a7%e4%bb%b6%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%96%87%e6%9c%ac%e3%80%81%e5%9b%be%e7%89%87%e5%92%8c%e6%8c%89%e9%92%ae%e5%9c%a8Flutter%e4%b8%ad%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f.md">12 经典控件（一）：文本、图片和按钮在Flutter中怎么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 经典控件（二）：UITableView_ListView在Flutter中是什么？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%e7%bb%8f%e5%85%b8%e6%8e%a7%e4%bb%b6%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aUITableView_ListView%e5%9c%a8Flutter%e4%b8%ad%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">13 经典控件（二）：UITableView_ListView在Flutter中是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 经典布局：如何定义子控件在父容器中排版的位置？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%e7%bb%8f%e5%85%b8%e5%b8%83%e5%b1%80%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9a%e4%b9%89%e5%ad%90%e6%8e%a7%e4%bb%b6%e5%9c%a8%e7%88%b6%e5%ae%b9%e5%99%a8%e4%b8%ad%e6%8e%92%e7%89%88%e7%9a%84%e4%bd%8d%e7%bd%ae%ef%bc%9f.md">14 经典布局：如何定义子控件在父容器中排版的位置？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 组合与自绘，我该选用何种方式自定义Widget？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/15%20%e7%bb%84%e5%90%88%e4%b8%8e%e8%87%aa%e7%bb%98%ef%bc%8c%e6%88%91%e8%af%a5%e9%80%89%e7%94%a8%e4%bd%95%e7%a7%8d%e6%96%b9%e5%bc%8f%e8%87%aa%e5%ae%9a%e4%b9%89Widget%ef%bc%9f.md">15 组合与自绘，我该选用何种方式自定义Widget？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 从夜间模式说起，如何定制不同风格的App主题？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%e4%bb%8e%e5%a4%9c%e9%97%b4%e6%a8%a1%e5%bc%8f%e8%af%b4%e8%b5%b7%ef%bc%8c%e5%a6%82%e4%bd%95%e5%ae%9a%e5%88%b6%e4%b8%8d%e5%90%8c%e9%a3%8e%e6%a0%bc%e7%9a%84App%e4%b8%bb%e9%a2%98%ef%bc%9f.md">16 从夜间模式说起，如何定制不同风格的App主题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 依赖管理（一）：图片、配置和字体在Flutter中怎么用？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%e4%be%9d%e8%b5%96%e7%ae%a1%e7%90%86%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%9b%be%e7%89%87%e3%80%81%e9%85%8d%e7%bd%ae%e5%92%8c%e5%ad%97%e4%bd%93%e5%9c%a8Flutter%e4%b8%ad%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f.md">17 依赖管理（一）：图片、配置和字体在Flutter中怎么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 依赖管理（二）：第三方组件库在Flutter中要如何管理？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/18%20%e4%be%9d%e8%b5%96%e7%ae%a1%e7%90%86%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%ac%ac%e4%b8%89%e6%96%b9%e7%bb%84%e4%bb%b6%e5%ba%93%e5%9c%a8Flutter%e4%b8%ad%e8%a6%81%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%ef%bc%9f.md">18 依赖管理（二）：第三方组件库在Flutter中要如何管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 用户交互事件该如何响应？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/19%20%e7%94%a8%e6%88%b7%e4%ba%a4%e4%ba%92%e4%ba%8b%e4%bb%b6%e8%af%a5%e5%a6%82%e4%bd%95%e5%93%8d%e5%ba%94%ef%bc%9f.md">19 用户交互事件该如何响应？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 关于跨组件传递数据，你只需要记住这三招.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%e5%85%b3%e4%ba%8e%e8%b7%a8%e7%bb%84%e4%bb%b6%e4%bc%a0%e9%80%92%e6%95%b0%e6%8d%ae%ef%bc%8c%e4%bd%a0%e5%8f%aa%e9%9c%80%e8%a6%81%e8%ae%b0%e4%bd%8f%e8%bf%99%e4%b8%89%e6%8b%9b.md">20 关于跨组件传递数据，你只需要记住这三招.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 路由与导航，Flutter是这样实现页面切换的.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/21%20%e8%b7%af%e7%94%b1%e4%b8%8e%e5%af%bc%e8%88%aa%ef%bc%8cFlutter%e6%98%af%e8%bf%99%e6%a0%b7%e5%ae%9e%e7%8e%b0%e9%a1%b5%e9%9d%a2%e5%88%87%e6%8d%a2%e7%9a%84.md">21 路由与导航，Flutter是这样实现页面切换的.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 如何构造炫酷的动画效果？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%e5%a6%82%e4%bd%95%e6%9e%84%e9%80%a0%e7%82%ab%e9%85%b7%e7%9a%84%e5%8a%a8%e7%94%bb%e6%95%88%e6%9e%9c%ef%bc%9f.md">22 如何构造炫酷的动画效果？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 单线程模型怎么保证UI运行流畅？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%e5%8d%95%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b%e6%80%8e%e4%b9%88%e4%bf%9d%e8%af%81UI%e8%bf%90%e8%a1%8c%e6%b5%81%e7%95%85%ef%bc%9f.md">23 单线程模型怎么保证UI运行流畅？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 HTTP网络编程与JSON解析.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/24%20HTTP%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e4%b8%8eJSON%e8%a7%a3%e6%9e%90.md">24 HTTP网络编程与JSON解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 本地存储与数据库的使用和优化.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%e6%9c%ac%e5%9c%b0%e5%ad%98%e5%82%a8%e4%b8%8e%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e4%bd%bf%e7%94%a8%e5%92%8c%e4%bc%98%e5%8c%96.md">25 本地存储与数据库的使用和优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 如何在Dart层兼容Android_iOS平台特定实现？（一）.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%e5%a6%82%e4%bd%95%e5%9c%a8Dart%e5%b1%82%e5%85%bc%e5%ae%b9Android_iOS%e5%b9%b3%e5%8f%b0%e7%89%b9%e5%ae%9a%e5%ae%9e%e7%8e%b0%ef%bc%9f%ef%bc%88%e4%b8%80%ef%bc%89.md">26 如何在Dart层兼容Android_iOS平台特定实现？（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 如何在Dart层兼容Android_iOS平台特定实现？（二）.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%e5%a6%82%e4%bd%95%e5%9c%a8Dart%e5%b1%82%e5%85%bc%e5%ae%b9Android_iOS%e5%b9%b3%e5%8f%b0%e7%89%b9%e5%ae%9a%e5%ae%9e%e7%8e%b0%ef%bc%9f%ef%bc%88%e4%ba%8c%ef%bc%89.md">27 如何在Dart层兼容Android_iOS平台特定实现？（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 如何在原生应用中混编Flutter工程？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%e5%a6%82%e4%bd%95%e5%9c%a8%e5%8e%9f%e7%94%9f%e5%ba%94%e7%94%a8%e4%b8%ad%e6%b7%b7%e7%bc%96Flutter%e5%b7%a5%e7%a8%8b%ef%bc%9f.md">28 如何在原生应用中混编Flutter工程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 混合开发，该用何种方案管理导航栈？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%e6%b7%b7%e5%90%88%e5%bc%80%e5%8f%91%ef%bc%8c%e8%af%a5%e7%94%a8%e4%bd%95%e7%a7%8d%e6%96%b9%e6%a1%88%e7%ae%a1%e7%90%86%e5%af%bc%e8%88%aa%e6%a0%88%ef%bc%9f.md">29 混合开发，该用何种方案管理导航栈？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 为什么需要做状态管理，怎么做？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e5%81%9a%e7%8a%b6%e6%80%81%e7%ae%a1%e7%90%86%ef%bc%8c%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f.md">30 为什么需要做状态管理，怎么做？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 如何实现原生推送能力？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/31%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%94%9f%e6%8e%a8%e9%80%81%e8%83%bd%e5%8a%9b%ef%bc%9f.md">31 如何实现原生推送能力？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 适配国际化，除了多语言我们还需要注意什么_.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%e9%80%82%e9%85%8d%e5%9b%bd%e9%99%85%e5%8c%96%ef%bc%8c%e9%99%a4%e4%ba%86%e5%a4%9a%e8%af%ad%e8%a8%80%e6%88%91%e4%bb%ac%e8%bf%98%e9%9c%80%e8%a6%81%e6%b3%a8%e6%84%8f%e4%bb%80%e4%b9%88_.md">32 适配国际化，除了多语言我们还需要注意什么_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 如何适配不同分辨率的手机屏幕？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/33%20%e5%a6%82%e4%bd%95%e9%80%82%e9%85%8d%e4%b8%8d%e5%90%8c%e5%88%86%e8%be%a8%e7%8e%87%e7%9a%84%e6%89%8b%e6%9c%ba%e5%b1%8f%e5%b9%95%ef%bc%9f.md">33 如何适配不同分辨率的手机屏幕？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 如何理解Flutter的编译模式？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/34%20%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3Flutter%e7%9a%84%e7%bc%96%e8%af%91%e6%a8%a1%e5%bc%8f%ef%bc%9f.md">34 如何理解Flutter的编译模式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 Hot Reload是怎么做到的？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/35%20Hot%20Reload%e6%98%af%e6%80%8e%e4%b9%88%e5%81%9a%e5%88%b0%e7%9a%84%ef%bc%9f.md">35 Hot Reload是怎么做到的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 如何通过工具链优化开发调试效率？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/36%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e5%b7%a5%e5%85%b7%e9%93%be%e4%bc%98%e5%8c%96%e5%bc%80%e5%8f%91%e8%b0%83%e8%af%95%e6%95%88%e7%8e%87%ef%bc%9f.md">36 如何通过工具链优化开发调试效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 如何检测并优化Flutter App的整体性能表现？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/37%20%e5%a6%82%e4%bd%95%e6%a3%80%e6%b5%8b%e5%b9%b6%e4%bc%98%e5%8c%96Flutter%20App%e7%9a%84%e6%95%b4%e4%bd%93%e6%80%a7%e8%83%bd%e8%a1%a8%e7%8e%b0%ef%bc%9f.md">37 如何检测并优化Flutter App的整体性能表现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 如何通过自动化测试提高交付质量？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/38%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%e6%8f%90%e9%ab%98%e4%ba%a4%e4%bb%98%e8%b4%a8%e9%87%8f%ef%bc%9f.md">38 如何通过自动化测试提高交付质量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 线上出现问题，该如何做好异常捕获与信息采集？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/39%20%e7%ba%bf%e4%b8%8a%e5%87%ba%e7%8e%b0%e9%97%ae%e9%a2%98%ef%bc%8c%e8%af%a5%e5%a6%82%e4%bd%95%e5%81%9a%e5%a5%bd%e5%bc%82%e5%b8%b8%e6%8d%95%e8%8e%b7%e4%b8%8e%e4%bf%a1%e6%81%af%e9%87%87%e9%9b%86%ef%bc%9f.md">39 线上出现问题，该如何做好异常捕获与信息采集？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 衡量Flutter App线上质量，我们需要关注这三个指标.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/40%20%e8%a1%a1%e9%87%8fFlutter%20App%e7%ba%bf%e4%b8%8a%e8%b4%a8%e9%87%8f%ef%bc%8c%e6%88%91%e4%bb%ac%e9%9c%80%e8%a6%81%e5%85%b3%e6%b3%a8%e8%bf%99%e4%b8%89%e4%b8%aa%e6%8c%87%e6%a0%87.md">40 衡量Flutter App线上质量，我们需要关注这三个指标.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 组件化和平台化，该如何组织合理稳定的Flutter工程结构？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/41%20%e7%bb%84%e4%bb%b6%e5%8c%96%e5%92%8c%e5%b9%b3%e5%8f%b0%e5%8c%96%ef%bc%8c%e8%af%a5%e5%a6%82%e4%bd%95%e7%bb%84%e7%bb%87%e5%90%88%e7%90%86%e7%a8%b3%e5%ae%9a%e7%9a%84Flutter%e5%b7%a5%e7%a8%8b%e7%bb%93%e6%9e%84%ef%bc%9f.md">41 组件化和平台化，该如何组织合理稳定的Flutter工程结构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 如何构建高效的Flutter App打包发布环境？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/42%20%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e9%ab%98%e6%95%88%e7%9a%84Flutter%20App%e6%89%93%e5%8c%85%e5%8f%91%e5%b8%83%e7%8e%af%e5%a2%83%ef%bc%9f.md">42 如何构建高效的Flutter App打包发布环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 如何构建自己的Flutter混合开发框架（一）？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/43%20%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e8%87%aa%e5%b7%b1%e7%9a%84Flutter%e6%b7%b7%e5%90%88%e5%bc%80%e5%8f%91%e6%a1%86%e6%9e%b6%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9f.md">43 如何构建自己的Flutter混合开发框架（一）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 如何构建自己的Flutter混合开发框架（二）？.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/44%20%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e8%87%aa%e5%b7%b1%e7%9a%84Flutter%e6%b7%b7%e5%90%88%e5%bc%80%e5%8f%91%e6%a1%86%e6%9e%b6%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9f.md">44 如何构建自己的Flutter混合开发框架（二）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送   温故而知新，与你说说专栏的那些思考题.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%20%20%e6%b8%a9%e6%95%85%e8%80%8c%e7%9f%a5%e6%96%b0%ef%bc%8c%e4%b8%8e%e4%bd%a0%e8%af%b4%e8%af%b4%e4%b8%93%e6%a0%8f%e7%9a%84%e9%82%a3%e4%ba%9b%e6%80%9d%e8%80%83%e9%a2%98.md">特别放送   温故而知新，与你说说专栏的那些思考题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 勿畏难，勿轻略.md" href="/%e4%b8%93%e6%a0%8f/Flutter%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%8b%bf%e7%95%8f%e9%9a%be%ef%bc%8c%e5%8b%bf%e8%bd%bb%e7%95%a5.md">结束语 勿畏难，勿轻略.md</a>
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
                            <h1 id="title" data-id="28 如何在原生应用中混编Flutter工程？" class="title">28 如何在原生应用中混编Flutter工程？</h1>
                            <div><p>你好，我是陈航。今天，我来和你聊聊如何在原生应用中接入Flutter。</p>

<p>在前面两篇文章中，我与你分享了如何在Dart层引入Android/iOS平台特定的能力，来提升App的功能体验。</p>

<p>使用Flutter从头开始写一个App，是一件轻松惬意的事情。但，对于成熟产品来说，完全摒弃原有App的历史沉淀，而全面转向Flutter并不现实。用Flutter去统一iOS/Android技术栈，把它作为已有原生App的扩展能力，通过逐步试验有序推进从而提升终端开发效率，可能才是现阶段Flutter最具吸引力的地方。</p>

<p>那么，Flutter工程与原生工程该如何组织管理？不同平台的Flutter工程打包构建产物该如何抽取封装？封装后的产物该如何引入原生工程？原生工程又该如何使用封装后的Flutter能力？</p>

<p>这些问题使得在已有原生App中接入Flutter看似并不是一件容易的事情。那接下来，我就和你介绍下如何在原生App中以最自然的方式接入Flutter。</p>

<h2 id="准备工作">准备工作</h2>

<p>既然是要在原生应用中混编Flutter，相信你一定已经准备好原生应用工程来实施今天的改造了。如果你还没有准备好也没关系，我会以一个最小化的示例和你演示这个改造过程。</p>

<p>首先，我们分别用Xcode与Android Studio快速建立一个只有首页的基本工程，工程名分别为iOSDemo与AndroidDemo。</p>

<p>这时，Android工程就已经准备好了；而对于iOS工程来说，由于基本工程并不支持以组件化的方式管理项目，因此我们还需要多做一步，将其改造成使用CocoaPods管理的工程，也就是要在iOSDemo根目录下创建一个只有基本信息的Podfile文件：</p>

<pre><code>use_frameworks!
platform :ios, '8.0'
target 'iOSDemo' do
#todo
end
</code></pre>

<p>然后，在命令行输入pod install后，会自动生成一个iOSDemo.xcworkspace文件，这时我们就完成了iOS工程改造。</p>

<h2 id="flutter混编方案介绍">Flutter混编方案介绍</h2>

<p>如果你想要在已有的原生App里嵌入一些Flutter页面，有两个办法：</p>

<ul>
<li>将原生工程作为Flutter工程的子工程，由Flutter统一管理。这种模式，就是统一管理模式。</li>
<li>将Flutter工程作为原生工程共用的子模块，维持原有的原生工程管理方式不变。这种模式，就是三端分离模式。</li>
</ul>

<p><img src="assets/28ff7c1a0ab44377bff8f9688671a5eb.jpg" alt="" /></p>

<p>图1 Flutter混编工程管理方式</p>

<p>由于Flutter早期提供的混编方式能力及相关资料有限，国内较早使用Flutter混合开发的团队大多使用的是统一管理模式。但是，随着功能迭代的深入，这种方案的弊端也随之显露，不仅三端（Android、iOS、Flutter）代码耦合严重，相关工具链耗时也随之大幅增长，导致开发效率降低。</p>

<p>所以，后续使用Flutter混合开发的团队陆续按照三端代码分离的模式来进行依赖治理，实现了Flutter工程的轻量级接入。</p>

<p>除了可以轻量级接入，三端代码分离模式把Flutter模块作为原生工程的子模块，还可以快速实现Flutter功能的“热插拔”，降低原生工程的改造成本。而Flutter工程通过Android Studio进行管理，无需打开原生工程，可直接进行Dart代码和原生代码的开发调试。</p>

<p><strong>三端工程分离模式的关键是抽离Flutter工程，将不同平台的构建产物依照标准组件化的形式进行管理</strong>，即Android使用aar、iOS使用pod。换句话说，接下来介绍的混编方案会将Flutter模块打包成aar和pod，这样原生工程就可以像引用其他第三方原生组件库那样快速接入Flutter了。</p>

<p>听起来是不是很兴奋？接下来，我们就开始正式采用三端分离模式来接入Flutter模块吧。</p>

<h2 id="集成flutter">集成Flutter</h2>

<p>我曾在前面的文章中提到，Flutter的工程结构比较特殊，包括Flutter工程和原生工程的目录（即iOS和Android两个目录）。在这种情况下，原生工程就会依赖于Flutter相关的库和资源，从而无法脱离父目录进行独立构建和运行。</p>

<p>原生工程对Flutter的依赖主要分为两部分：</p>

<ul>
<li>Flutter库和引擎，也就是Flutter的Framework库和引擎库；</li>
<li>Flutter工程，也就是我们自己实现的Flutter模块功能，主要包括Flutter工程lib目录下的Dart代码实现的这部分功能。</li>
</ul>

<p>在已经有原生工程的情况下，我们需要在同级目录创建Flutter模块，构建iOS和Android各自的Flutter依赖库。这也很好实现，Flutter就为我们提供了这样的命令。我们只需要在原生项目的同级目录下，执行Flutter命令创建名为flutter_library的模块即可：</p>

<pre><code>Flutter create -t module flutter_library
</code></pre>

<p>这里的Flutter模块，也是Flutter工程，我们用Android Studio打开它，其目录如下图所示：</p>

<p><img src="assets/d7f154a2e5f244e384802f3cd746fc7f.jpg" alt="" /></p>

<p>图2 Flutter模块工程结构</p>

<p>可以看到，和传统的Flutter工程相比，Flutter模块工程也有内嵌的Android工程与iOS工程，因此我们可以像普通工程一样使用Android Studio进行开发调试。</p>

<p>仔细查看可以发现，<strong>Flutter模块有一个细微的变化</strong>：Android工程下多了一个Flutter目录，这个目录下的build.gradle配置就是我们构建aar的打包配置。这就是模块工程既能像Flutter传统工程一样使用Android Studio开发调试，又能打包构建aar与pod的秘密。</p>

<p>实际上，iOS工程的目录结构也有细微变化，但这个差异并不影响打包构建，因此我就不再展开了。</p>

<p>然后，我们打开main.dart文件，将其逻辑更新为以下代码逻辑，即一个写着“Hello from Flutter”的全屏红色的Flutter Widget：</p>

<pre><code>import 'package:flutter/material.dart';
import 'dart:ui';

void main() =&gt; runApp(_widgetForRoute(window.defaultRouteName));//独立运行传入默认路由

Widget _widgetForRoute(String route) {
  switch (route) {
    default:
      return MaterialApp(
        home: Scaffold(
          backgroundColor: const Color(0xFFD63031),//ARGB红色
          body: Center(
            child: Text(
              'Hello from Flutter', //显示的文字
              textDirection: TextDirection.ltr,
              style: TextStyle(
                fontSize: 20.0,
                color: Colors.blue,
              ),
            ),
          ),
        ),
      );
  }
}
</code></pre>

<p>注意：我们创建的Widget实际上是包在一个switch-case语句中的。这是因为封装的Flutter模块一般会有多个页面级Widget，原生App代码则会通过传入路由标识字符串，告诉Flutter究竟应该返回何种Widget。为了简化案例，在这里我们忽略标识字符串，统一返回一个MaterialApp。</p>

<p>接下来，我们要做的事情就是把这段代码编译打包，构建出对应的Android和iOS依赖库，实现原生工程的接入。</p>

<p>现在，我们首先来看看Android工程如何接入。</p>

<h3 id="android模块集成">Android模块集成</h3>

<p>之前我们提到原生工程对Flutter的依赖主要分为两部分，对应到Android平台，这两部分分别是：</p>

<ul>
<li>Flutter库和引擎，也就是icudtl.dat、libFlutter.so，还有一些class文件。这些文件都封装在Flutter.jar中。</li>
<li>Flutter工程产物，主要包括应用程序数据段isolate_snapshot_data、应用程序指令段isolate_snapshot_instr、虚拟机数据段vm_snapshot_data、虚拟机指令段vm_snapshot_instr、资源文件Flutter_assets。</li>
</ul>

<p>搞清楚Flutter工程的Android编译产物之后，我们对Android的Flutter依赖抽取步骤如下：</p>

<p>首先在Flutter_library的根目录下，执行aar打包构建命令：</p>

<pre><code>Flutter build apk --debug
</code></pre>

<p>这条命令的作用是编译工程产物，并将Flutter.jar和工程产物编译结果封装成一个aar。你很快就会想到，如果是构建release产物，只需要把debug换成release就可以了。</p>

<p><strong>其次</strong>，打包构建的flutter-debug.aar位于.android/Flutter/build/outputs/aar/目录下，我们把它拷贝到原生Android工程AndroidDemo的app/libs目录下，并在App的打包配置build.gradle中添加对它的依赖:</p>

<pre><code>...
repositories {
    flatDir {
        dirs 'libs'   // aar目录
    }
}
android {
    ...
    compileOptions {
        sourceCompatibility 1.8 //Java 1.8
        targetCompatibility 1.8 //Java 1.8
    }
    ...
}

dependencies {
    ...
    implementation(name: 'flutter-debug', ext: 'aar')//Flutter模块aar
    ...
}
</code></pre>

<p>Sync一下，Flutter模块就被添加到了Android项目中。</p>

<p>再次，我们试着改一下MainActivity.java的代码，把它的contentView改成Flutter的widget：</p>

<pre><code>protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    View FlutterView = Flutter.createView(this, getLifecycle(), &quot;defaultRoute&quot;); //传入路由标识符
    setContentView(FlutterView);//用FlutterView替代Activity的ContentView
}
</code></pre>

<p>最后点击运行，可以看到一个写着“Hello from Flutter”的全屏红色的Flutter Widget就展示出来了。至此，我们完成了Android工程的接入。</p>

<p><img src="assets/66f1a9f681b84e0c972db14fe5c7983b.jpg" alt="" /></p>

<p>图3 Android工程接入示例</p>

<h3 id="ios模块集成">iOS模块集成</h3>

<p>iOS工程接入的情况要稍微复杂一些。在iOS平台，原生工程对Flutter的依赖分别是：</p>

<ul>
<li>Flutter库和引擎，即Flutter.framework；</li>
<li>Flutter工程的产物，即App.framework。</li>
</ul>

<p>iOS平台的Flutter模块抽取，实际上就是通过打包命令生成这两个产物，并将它们封装成一个pod供原生工程引用。</p>

<p>类似地，首先我们在Flutter_library的根目录下，执行iOS打包构建命令：</p>

<pre><code>Flutter build ios --debug
</code></pre>

<p>这条命令的作用是编译Flutter工程生成两个产物：Flutter.framework和App.framework。同样，把debug换成release就可以构建release产物（当然，你还需要处理一下签名问题）。</p>

<p><strong>其次</strong>，在iOSDemo的根目录下创建一个名为FlutterEngine的目录，并把这两个framework文件拷贝进去。iOS的模块化产物工作要比Android多一个步骤，因为我们需要把这两个产物手动封装成pod。因此，我们还需要在该目录下创建FlutterEngine.podspec，即Flutter模块的组件定义：</p>

<pre><code>Pod::Spec.new do |s|
  s.name             = 'FlutterEngine'
  s.version          = '0.1.0'
  s.summary          = 'XXXXXXX'
  s.description      = &lt;&lt;-DESC
TODO: Add long description of the pod here.
                       DESC
  s.homepage         = 'https://github.com/xx/FlutterEngine'
  s.license          = { :type =&gt; 'MIT', :file =&gt; 'LICENSE' }
  s.author           = { 'chenhang' =&gt; '<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="97fff6f9f0fee4f9fef4f2d7f0faf6fefbb9f4f8fa">[email&#160;protected]</a>' }
  s.source       = { :git =&gt; &quot;&quot;, :tag =&gt; &quot;#{s.version}&quot; }
  s.ios.deployment_target = '8.0'
  s.ios.vendored_frameworks = 'App.framework', 'Flutter.framework'
end
</code></pre>

<p>pod lib lint一下，Flutter模块组件就已经做好了。趁热打铁，我们再修改Podfile文件把它集成到iOSDemo工程中：</p>

<pre><code>...
target 'iOSDemo' do
    pod 'FlutterEngine', :path =&gt; './'
end
</code></pre>

<p>pod install一下，Flutter模块就集成进iOS原生工程中了。</p>

<p>再次，我们试着修改一下AppDelegate.m的代码，把window的rootViewController改成FlutterViewController：</p>

<pre><code>- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions

{
    self.window = [[UIWindow alloc] initWithFrame:[UIScreen mainScreen].bounds];
    FlutterViewController *vc = [[FlutterViewController alloc]init];
    [vc setInitialRoute:@&quot;defaultRoute&quot;]; //路由标识符
    self.window.rootViewController = vc;
    [self.window makeKeyAndVisible];
    return YES;
}
</code></pre>

<p>最后点击运行，一个写着“Hello from Flutter”的全屏红色的Flutter Widget也展示出来了。至此，iOS工程的接入我们也顺利搞定了。</p>

<p><img src="assets/50ae2b3e7fa847308a4e427b11d0c132.jpg" alt="" /></p>

<p>图4 iOS工程接入示例</p>

<h2 id="总结">总结</h2>

<p>通过分离Android、iOS和Flutter三端工程，抽离Flutter库和引擎及工程代码为组件库，以Android和iOS平台最常见的aar和pod形式接入原生工程，我们就可以低成本地接入Flutter模块，愉快地使用Flutter扩展原生App的边界了。</p>

<p>但，我们还可以做得更好。</p>

<p>如果每次通过构建Flutter模块工程，都是手动搬运Flutter编译产物，那很容易就会因为工程管理混乱导致Flutter组件库被覆盖，从而引发难以排查的Bug。而要解决此类问题的话，我们可以引入CI自动构建框架，把Flutter编译产物构建自动化，原生工程通过接入不同版本的构建产物，实现更优雅的三端分离模式。</p>

<p>而关于自动化构建，我会在后面的文章中和你详细介绍，这里就不再赘述了。</p>

<p>接下来，我们简单回顾一下今天的内容。</p>

<p>原生工程混编Flutter的方式有两种。一种是，将Flutter工程内嵌Android和iOS工程，由Flutter统一管理的集中模式；另一种是，将Flutter工程作为原生工程共用的子模块，由原生工程各自管理的三端工程分离模式。目前，业界采用的基本都是第二种方式。</p>

<p>而对于三端工程分离模式最主要的则是抽离Flutter工程，将不同平台的构建产物依照标准组件化的形式进行管理，即：针对Android平台打包构建生成aar，通过build.gradle进行依赖管理；针对iOS平台打包构建生成framework，将其封装成独立的pod，并通过podfile进行依赖管理。</p>

<p>我把今天分享所涉及到的知识点打包到了GitHub（<a href="https://github.com/cyndibaby905/28_module_page" target="_blank">flutter_module_page</a>、<a href="https://github.com/cyndibaby905/28_iOSDemo" target="_blank">iOS_demo</a>、<a href="https://github.com/cyndibaby905/28_AndroidDemo" target="_blank">Android_Demo</a>）中，你可以下载下来，反复运行几次，加深理解与记忆。</p>

<h2 id="思考题">思考题</h2>

<p>最后，我给你下留一个思考题吧。</p>

<p>对于有资源依赖的Flutter模块工程而言，其打包构建的产物，以及抽离Flutter组件库的过程会有什么不同吗？</p>

<p>欢迎你在评论区给我留言分享你的观点，我会在下一篇文章中等待你！感谢你的收听，也欢迎你把这篇文章分享给更多的朋友一起阅读。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#2e424242171a1f1f1e196e49434f4742004d4143" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0de3269b08653b',t:'MTczNDAwODMyMS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
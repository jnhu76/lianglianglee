<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=40&#32;衡量Flutter&#32;App线上质量，我们需要关注这三个指标>
        <link rel="icon" href="/static/favicon.png">
        <title>40 衡量Flutter App线上质量，我们需要关注这三个指标 </title>
        
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
                            <h1 id="title" data-id="40 衡量Flutter App线上质量，我们需要关注这三个指标" class="title">40 衡量Flutter App线上质量，我们需要关注这三个指标</h1>
                            <div><p>你好，我是陈航。</p>

<p>在上一篇文章中，我与你分享了如何捕获Flutter应用的未处理异常。所谓异常，指的是Dart代码在运行时意外发生的错误事件。对于单一异常来说，我们可以使用try-catch，或是catchError去处理；而如果我们想对异常进行集中的拦截治理，则需要使用Zone，并结合FlutterError进行统一管理。异常一旦被抓取，我们就可以利用第三方数据上报服务（比如Bugly），上报其上下文信息了。</p>

<p>这些线上异常的监控数据，对于开发者尽早发现线上隐患，确定问题根因至关重要。如果我们想进一步评估应用整体的稳定性的话，就需要把异常信息与页面的渲染关联起来。比如，页面渲染过程是否出现了异常，而导致功能不可用？</p>

<p>而对于以“丝般顺滑”著称的Flutter应用而言，页面渲染的性能同样需要我们重点关注。比如，界面渲染是否出现会掉帧卡顿现象，或者页面加载是否会出现性能问题导致耗时过长？这些问题，虽不至于让应用完全不能使用，但也很容易引起用户对应用质量的质疑，甚至是反感。</p>

<p>通过上面的分析，可以看到，衡量线上Flutter应用整体质量的指标，可以分为以下3类：</p>

<ul>
<li>页面异常率；</li>
<li>页面帧率；</li>
<li>页面加载时长。</li>
</ul>

<p>其中，页面异常率反应了页面的健康程度，页面帧率反应了视觉效果的顺滑程度，而页面加载时长则反应了整个渲染过程中点对点的延时情况。</p>

<p>这三项数据指标，是度量Flutter应用是否优秀的重要质量指标。通过梳理这些指标的统计口径，建立起Flutter应用的质量监控能力，这样一来我们不仅可以及早发现线上隐患，还可以确定质量基线，从而持续提升用户体验。</p>

<p>所以在今天的分享中，我会与你详细讲述这3项指标是如何采集的。</p>

<h2 id="页面异常率">页面异常率</h2>

<p>页面异常率指的是，页面渲染过程中出现异常的概率。它度量的是页面维度下功能不可用的情况，其统计公式为：<strong>页面异常率=异常发生次数/整体页面PV数</strong>。</p>

<p>在了解了页面异常率的统计口径之后，接下来我们分别来看一下这个公式中的分子与分母应该如何统计吧。</p>

<p>我们先来看看<strong>异常发生次数的统计方法</strong>。通过上一篇文章，我们已经知道了在Flutter中，未处理异常需要通过Zone与FlutterError去捕获。所以，如果我们想统计异常发生次数的话，依旧是利用这两个方法，只不过要在异常拦截的方法中，通过一个计数器进行累加，统一记录。</p>

<p>下面的例子演示了异常发生次数的具体统计方法。我们使用全局变量exceptionCount，在异常捕获的回调方法_reportError中持续地累加捕获到的异常次数：</p>

<pre><code>int exceptionCount = 0; 
Future&lt;Null&gt; _reportError(dynamic error, dynamic stackTrace) async {
  exceptionCount++; //累加异常次数
  FlutterCrashPlugin.postException(error, stackTrace);
}

Future&lt;Null&gt; main() async {
  FlutterError.onError = (FlutterErrorDetails details) async {
    //将异常转发至Zone
    Zone.current.handleUncaughtError(details.exception, details.stack);
  };

  runZoned&lt;Future&lt;Null&gt;&gt;(() async {
    runApp(MyApp());
  }, onError: (error, stackTrace) async {
    //拦截异常
    await _reportError(error, stackTrace);
  });
}
</code></pre>

<p>接下来，我们再看看<strong>整体页面PV数如何统计</strong>吧。整体页面PV数，其实就是页面的打开次数。通过第21篇文章“<a href="https://time.geekbang.org/column/article/118421" target="_blank">路由与导航，Flutter是这样实现页面切换的</a>”，我们已经知道了Flutter页面的切换需要经过Navigator来实现，所以页面切换状态也需要通过Navigator才能感知到。</p>

<p>与注册页面路由类似的，在MaterialApp中，我们可以通过NavigatorObservers属性，去监听页面的打开与关闭。下面的例子演示了<strong>NavigatorObserver的具体用法</strong>。在下面的代码中，我们定义了一个继承自NavigatorObserver的观察者，并在其didPush方法中，去统计页面的打开行为：</p>

<pre><code>int totalPV = 0;
//导航监听器
class MyObserver extends NavigatorObserver{
  @override
  void didPush(Route route, Route previousRoute) {
    super.didPush(route, previousRoute);
    totalPV++;//累加PV
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return  MaterialApp(
    //设置路由监听
       navigatorObservers: [
         MyObserver(),
       ],
       home: HomePage(),
    ); 
  }   
}
</code></pre>

<p>现在，我们已经收集到了异常发生次数和整体页面PV数这两个参数，接下来我们就可以计算出页面异常率了：</p>

<pre><code>double pageException() {
  if(totalPV == 0) return 0;
  return exceptionCount/totalPV;
}
</code></pre>

<p>可以看到，页面异常率的计算还是相对比较简单的。</p>

<h2 id="页面帧率">页面帧率</h2>

<p>页面帧率，即FPS，是图像领域中的定义，指的是画面每秒传输帧数。由于人眼的视觉暂留特质，当所见到的画面传输帧数高于一定数量的时候，就会认为是连贯性的视觉效果。因此，对于动态页面而言，每秒钟展示的帧数越多，画面就越流畅。</p>

<p>由此我们可以得出，<strong>FPS的计算口径为单位时间内渲染的帧总数</strong>。在移动设备中，FPS的推荐数值通常是60Hz，即每秒刷新页面60次。</p>

<p>为什么是60Hz，而不是更高或更低的值呢？这是因为显示过程，是由VSync信号周期性驱动的，而VSync信号的周期就是每秒60次，这也是FPS的上限。</p>

<p>CPU与GPU在接收到VSync信号后，就会计算图形图像，准备渲染内容，并将其提交到帧缓冲区，等待下一次VSync信号到来时显示到屏幕上。如果在一个VSync时间内，CPU或者GPU没有完成内容提交，这一帧就会被丢弃，等待下一次机会再显示，而这时页面会保留之前的内容不变，造成界面卡顿。因此，FPS低于60Hz时就会出现掉帧现象，而如果低于45Hz则会有比较严重的卡顿现象。</p>

<p>为方便开发者统计FPS，Flutter在全局window对象上提供了帧回调机制。我们可以<strong>在window对象上注册onReportTimings方法</strong>，将最近绘制帧耗费的时间（即FrameTiming），以回调的形式告诉我们。有了每一帧的绘制时间后，我们就可以计算FPS了。</p>

<p>需要注意的是，onReportTimings方法只有在有帧被绘制时才有数据回调，如果用户没有和App发生交互，界面状态没有变化时，是不会产生新的帧的。考虑到单个帧的绘制时间差异较大，逐帧计算可能会产生数据跳跃，所以为了让FPS的计算更加平滑，我们需要保留最近25个FrameTiming用于求和计算。</p>

<p>而另一方面，对于FPS的计算，我们并不能孤立地只考虑帧绘制时间，而应该结合VSync信号的周期，即1/60秒（即16.67毫秒）来综合评估。</p>

<p>由于帧的渲染是依靠VSync信号驱动的，如果帧绘制的时间没有超过16.67毫秒，我们也需要把它当成16.67毫秒来算，因为绘制完成的帧必须要等到下一次VSync信号来了之后才能渲染。而如果帧绘制时间超过了16.67毫秒，则会占用后续的VSync信号周期，从而打乱后续的绘制次序，产生卡顿现象。这里有两种情况：</p>

<ul>
<li>如果帧绘制时间正好是16.67的整数倍，比如50，则代表它花费了3个VSync信号周期，即本来可以绘制3帧，但实际上只绘制了1帧；</li>
<li>如果帧绘制时间不是16.67的整数倍，比如51，那么它花费的VSync信号周期应该向上取整，即4个，这意味着本来可以绘制4帧，实际上只绘制了1帧。</li>
</ul>

<p>所以我们的FPS计算公式最终确定为：<strong>FPS=60*实际渲染的帧数/本来应该在这个时间内渲染完成的帧数</strong>。</p>

<p>下面的示例演示了如何通过onReportTimings回调函数实现FPS的计算。在下面的代码中，我们定义了一个容量为25的列表，用于存储最近的帧绘制耗时FrameTiming。在FPS的计算函数中，我们将列表中每帧绘制时间与VSync周期frameInterval进行比较，得出本来应该绘制的帧数，最后两者相除就得到了FPS指标。</p>

<p>需要注意的是，Android Studio提供的Flutter插件里展示的FPS信息，其实也来自于onReportTimings回调，所以我们在注册回调时需要保留原始回调引用，否则插件就读不到FPS信息了。</p>

<pre><code>import 'dart:ui';

var orginalCallback;

void main() {
  runApp(MyApp());
  //设置帧回调函数并保存原始帧回调函数
  orginalCallback = window.onReportTimings;
  window.onReportTimings = onReportTimings;
}

//仅缓存最近25帧绘制耗时
const maxframes = 25;
final lastFrames = List&lt;FrameTiming&gt;();
//基准VSync信号周期
const frameInterval = const Duration(microseconds: Duration.microsecondsPerSecond ~/ 60);

void onReportTimings(List&lt;FrameTiming&gt; timings) {
  lastFrames.addAll(timings);
  //仅保留25帧
  if(lastFrames.length &gt; maxframes) {
    lastFrames.removeRange(0, lastFrames.length - maxframes);
  }
  //如果有原始帧回调函数，则执行
  if (orginalCallback != null) {
    orginalCallback(timings);
  }
}

double get fps {
  int sum = 0;
  for (FrameTiming timing in lastFrames) {
    //计算渲染耗时
    int duration = timing.timestampInMicroseconds(FramePhase.rasterFinish) - timing.timestampInMicroseconds(FramePhase.buildStart);
    //判断耗时是否在Vsync信号周期内
    if(duration &lt; frameInterval.inMicroseconds) {
      sum += 1;
    } else {
      //有丢帧，向上取整
      int count = (duration/frameInterval.inMicroseconds).ceil();
      sum += count;
    }
  }
  return lastFrames.length/sum * 60;
}
</code></pre>

<p>运行这段代码，可以看到，我们统计的FPS指标和Flutter插件展示的FPS走势是一致的。</p>

<p><img src="assets/10ca7dd6b08849fb8924d55ef79c5664.jpg" alt="" /></p>

<p>图1 FPS指标走势</p>

<h2 id="页面加载时长">页面加载时长</h2>

<p>页面加载时长，指的是页面从创建到可见的时间。它反应的是代码中创建页面视图是否存在过度绘制，或者绘制不合理导致创建视图时间过长的情况。</p>

<p>从定义可以看出，<strong>页面加载时长的统计口径为页面可见的时间-页面创建的时间</strong>。获取页面创建的时间比较容易，我们只需要在页面的初始化函数里记录时间即可。那么，<strong>页面可见的时间应该如何统计</strong>呢？</p>

<p>在第11篇文章“<a href="https://time.geekbang.org/column/article/109490" target="_blank">提到生命周期，我们是在说什么？</a>”中，我在介绍Widget的生命周期时，曾向你介绍过Flutter的帧回调机制。WidgetsBinding提供了单次Frame回调addPostFrameCallback方法，它会在当前Frame绘制完成之后进行回调，并且只会回调一次。一旦监听到Frame绘制完成回调后，我们就可以确认页面已经被渲染出来了，因此我们可以借助这个方法去获取页面可见的时间。</p>

<p>下面的例子演示了如何通过帧回调机制获取页面加载时长。在下面的代码中，我们在页面MyPage的初始化方法中记录了页面的创建时间startTime，然后在页面状态的初始化方法中，通过addPostFrameCallback注册了单次帧绘制回调，并在回调函数中记录了页面的渲染完成时间endTime。将这两个时间做减法，我们就得到了MyPage的页面加载时长：</p>

<pre><code>class MyHomePage extends StatefulWidget {
  int startTime;
  int endTime;
  MyHomePage({Key key}) : super(key: key) {
    //页面初始化时记录启动时间
    startTime = DateTime.now().millisecondsSinceEpoch;
  }
  @override
  _MyHomePageState createState() =&gt; _MyHomePageState();
}

class _MyHomePageState extends State&lt;MyHomePage&gt; {
  @override
  void initState() {
    super.initState();
    //通过帧绘制回调获取渲染完成时间
    WidgetsBinding.instance.addPostFrameCallback((_) {
      widget.endTime = DateTime.now().millisecondsSinceEpoch;
      int timeSpend = widget.endTime - widget.startTime;
      print(&quot;Page render time:${timeSpend} ms&quot;);
    });
  }
  ...
}
</code></pre>

<p>试着运行一下代码，观察命令行输出：</p>

<pre><code>flutter: Page render time:548 ms
</code></pre>

<p>可以看到，通过单次帧绘制回调统计得出的页面加载时间为548毫秒。</p>

<p>至此，我们就已经得到了页面异常率、页面帧率和页面加载时长这3个指标了。</p>

<h2 id="总结">总结</h2>

<p>好了，今天的分享就到这里，我们来总结下主要内容吧。</p>

<p>今天我们一起学习了衡量Flutter应用线上质量的3个指标，即页面异常率、页面帧率和页面加载时长，以及分别对应的数据采集方式。</p>

<p>其中，页面异常率表示页面渲染过程中的稳定性，可以通过集中捕获未处理异常，结合NavigatorObservers观察页面PV，计算得出页面维度下功能不可用的概率。</p>

<p>页面帧率则表示了页面的流畅情况，可以利用Flutter提供的帧绘制耗时回调onReportTimings，以加权的形式计算出本应该绘制的帧数，得到更为准确的FPS。</p>

<p>而页面加载时长，反应的是渲染过程的延时情况。我们可以借助于单次帧回调机制，来获取页面渲染完成时间，从而得到整体页面的加载时长。</p>

<p>通过这3个数据指标统计方法，我们再去评估Flutter应用的性能时，就有一个具体的数字化标准了。而有了数据之后，我们不仅可以及早发现问题隐患，准确定位及修复问题，还可以根据它们去评估应用的健康程度和页面的渲染性能，从而确定后续的优化方向。</p>

<p>我把今天分享涉及的知识点打包到了<a href="https://github.com/cyndibaby905/40_peformance_demo" target="_blank">GitHub</a>中，你可以下载下来，反复运行几次，加深理解与记忆。</p>

<h2 id="思考题">思考题</h2>

<p>最后，我给你留一道思考题吧。</p>

<p>如果页面的渲染需要依赖单个或多个网络接口数据，这时的页面加载时长应该如何统计呢？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#701c1c1c49444141404730171d11191c5e131f1d" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0de8135be0653b',t:'MTczNDAwODUyMi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
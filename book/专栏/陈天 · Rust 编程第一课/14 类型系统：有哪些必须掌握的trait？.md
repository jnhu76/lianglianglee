<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=14&#32;类型系统：有哪些必须掌握的trait？>
        <link rel="icon" href="/static/favicon.png">
        <title>14 类型系统：有哪些必须掌握的trait？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 让Rust成为你的下一门主力语言.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e8%ae%a9Rust%e6%88%90%e4%b8%ba%e4%bd%a0%e7%9a%84%e4%b8%8b%e4%b8%80%e9%97%a8%e4%b8%bb%e5%8a%9b%e8%af%ad%e8%a8%80.md">00 开篇词 让Rust成为你的下一门主力语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 内存：值放堆上还是放栈上，这是一个问题.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/01%20%e5%86%85%e5%ad%98%ef%bc%9a%e5%80%bc%e6%94%be%e5%a0%86%e4%b8%8a%e8%bf%98%e6%98%af%e6%94%be%e6%a0%88%e4%b8%8a%ef%bc%8c%e8%bf%99%e6%98%af%e4%b8%80%e4%b8%aa%e9%97%ae%e9%a2%98.md">01 内存：值放堆上还是放栈上，这是一个问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 串讲：编程开发中，那些你需要掌握的基本概念.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/02%20%e4%b8%b2%e8%ae%b2%ef%bc%9a%e7%bc%96%e7%a8%8b%e5%bc%80%e5%8f%91%e4%b8%ad%ef%bc%8c%e9%82%a3%e4%ba%9b%e4%bd%a0%e9%9c%80%e8%a6%81%e6%8e%8c%e6%8f%a1%e7%9a%84%e5%9f%ba%e6%9c%ac%e6%a6%82%e5%bf%b5.md">02 串讲：编程开发中，那些你需要掌握的基本概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 初窥门径：从你的第一个Rust程序开始！.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/03%20%e5%88%9d%e7%aa%a5%e9%97%a8%e5%be%84%ef%bc%9a%e4%bb%8e%e4%bd%a0%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aaRust%e7%a8%8b%e5%ba%8f%e5%bc%80%e5%a7%8b%ef%bc%81.md">03 初窥门径：从你的第一个Rust程序开始！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 get hands dirty：来写个实用的CLI小工具.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/04%20get%20hands%20dirty%ef%bc%9a%e6%9d%a5%e5%86%99%e4%b8%aa%e5%ae%9e%e7%94%a8%e7%9a%84CLI%e5%b0%8f%e5%b7%a5%e5%85%b7.md">04 get hands dirty：来写个实用的CLI小工具.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 get hands dirty：做一个图片服务器有多难？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/05%20get%20hands%20dirty%ef%bc%9a%e5%81%9a%e4%b8%80%e4%b8%aa%e5%9b%be%e7%89%87%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%9c%89%e5%a4%9a%e9%9a%be%ef%bc%9f.md">05 get hands dirty：做一个图片服务器有多难？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 get hands dirty：SQL查询工具怎么一鱼多吃？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/06%20get%20hands%20dirty%ef%bc%9aSQL%e6%9f%a5%e8%af%a2%e5%b7%a5%e5%85%b7%e6%80%8e%e4%b9%88%e4%b8%80%e9%b1%bc%e5%a4%9a%e5%90%83%ef%bc%9f.md">06 get hands dirty：SQL查询工具怎么一鱼多吃？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 所有权：值的生杀大权到底在谁手上？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/07%20%e6%89%80%e6%9c%89%e6%9d%83%ef%bc%9a%e5%80%bc%e7%9a%84%e7%94%9f%e6%9d%80%e5%a4%a7%e6%9d%83%e5%88%b0%e5%ba%95%e5%9c%a8%e8%b0%81%e6%89%8b%e4%b8%8a%ef%bc%9f.md">07 所有权：值的生杀大权到底在谁手上？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 所有权：值的借用是如何工作的？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/08%20%e6%89%80%e6%9c%89%e6%9d%83%ef%bc%9a%e5%80%bc%e7%9a%84%e5%80%9f%e7%94%a8%e6%98%af%e5%a6%82%e4%bd%95%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%9f.md">08 所有权：值的借用是如何工作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 所有权：一个值可以有多个所有者么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/09%20%e6%89%80%e6%9c%89%e6%9d%83%ef%bc%9a%e4%b8%80%e4%b8%aa%e5%80%bc%e5%8f%af%e4%bb%a5%e6%9c%89%e5%a4%9a%e4%b8%aa%e6%89%80%e6%9c%89%e8%80%85%e4%b9%88%ef%bc%9f.md">09 所有权：一个值可以有多个所有者么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 生命周期：你创建的值究竟能活多久？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/10%20%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%ef%bc%9a%e4%bd%a0%e5%88%9b%e5%bb%ba%e7%9a%84%e5%80%bc%e7%a9%b6%e7%ab%9f%e8%83%bd%e6%b4%bb%e5%a4%9a%e4%b9%85%ef%bc%9f.md">10 生命周期：你创建的值究竟能活多久？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 内存管理：从创建到消亡，值都经历了什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/11%20%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%ef%bc%9a%e4%bb%8e%e5%88%9b%e5%bb%ba%e5%88%b0%e6%b6%88%e4%ba%a1%ef%bc%8c%e5%80%bc%e9%83%bd%e7%bb%8f%e5%8e%86%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">11 内存管理：从创建到消亡，值都经历了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 类型系统：Rust的类型系统有什么特点？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/12%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9aRust%e7%9a%84%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%e6%9c%89%e4%bb%80%e4%b9%88%e7%89%b9%e7%82%b9%ef%bc%9f.md">12 类型系统：Rust的类型系统有什么特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 类型系统：如何使用trait来定义接口？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/13%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8trait%e6%9d%a5%e5%ae%9a%e4%b9%89%e6%8e%a5%e5%8f%a3%ef%bc%9f.md">13 类型系统：如何使用trait来定义接口？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 类型系统：有哪些必须掌握的trait？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/14%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9a%e6%9c%89%e5%93%aa%e4%ba%9b%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84trait%ef%bc%9f.md">14 类型系统：有哪些必须掌握的trait？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 数据结构：这些浓眉大眼的结构竟然都是智能指针？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/15%20%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%ef%bc%9a%e8%bf%99%e4%ba%9b%e6%b5%93%e7%9c%89%e5%a4%a7%e7%9c%bc%e7%9a%84%e7%bb%93%e6%9e%84%e7%ab%9f%e7%84%b6%e9%83%bd%e6%98%af%e6%99%ba%e8%83%bd%e6%8c%87%e9%92%88%ef%bc%9f.md">15 数据结构：这些浓眉大眼的结构竟然都是智能指针？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 数据结构：Vec_T_、&amp;[T]、Box_[T]_ ，你真的了解集合容器么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/16%20%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%ef%bc%9aVec_T_%e3%80%81&amp;[T]%e3%80%81Box_[T]_%20%ef%bc%8c%e4%bd%a0%e7%9c%9f%e7%9a%84%e4%ba%86%e8%a7%a3%e9%9b%86%e5%90%88%e5%ae%b9%e5%99%a8%e4%b9%88%ef%bc%9f.md">16 数据结构：Vec_T_、&amp;[T]、Box_[T]_ ，你真的了解集合容器么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 数据结构：软件系统核心部件哈希表，内存如何布局？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/17%20%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%ef%bc%9a%e8%bd%af%e4%bb%b6%e7%b3%bb%e7%bb%9f%e6%a0%b8%e5%bf%83%e9%83%a8%e4%bb%b6%e5%93%88%e5%b8%8c%e8%a1%a8%ef%bc%8c%e5%86%85%e5%ad%98%e5%a6%82%e4%bd%95%e5%b8%83%e5%b1%80%ef%bc%9f.md">17 数据结构：软件系统核心部件哈希表，内存如何布局？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 错误处理：为什么Rust的错误处理与众不同？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/18%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Rust%e7%9a%84%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%e4%b8%8e%e4%bc%97%e4%b8%8d%e5%90%8c%ef%bc%9f.md">18 错误处理：为什么Rust的错误处理与众不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 闭包：FnOnce、FnMut和Fn，为什么有这么多类型？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/19%20%e9%97%ad%e5%8c%85%ef%bc%9aFnOnce%e3%80%81FnMut%e5%92%8cFn%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e6%9c%89%e8%bf%99%e4%b9%88%e5%a4%9a%e7%b1%bb%e5%9e%8b%ef%bc%9f.md">19 闭包：FnOnce、FnMut和Fn，为什么有这么多类型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 4 Steps ：如何更好地阅读Rust源码？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/20%204%20Steps%20%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9b%b4%e5%a5%bd%e5%9c%b0%e9%98%85%e8%af%bbRust%e6%ba%90%e7%a0%81%ef%bc%9f.md">20 4 Steps ：如何更好地阅读Rust源码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 阶段实操（1）：构建一个简单的KV server-基本流程.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/21%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%881%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e5%9f%ba%e6%9c%ac%e6%b5%81%e7%a8%8b.md">21 阶段实操（1）：构建一个简单的KV server-基本流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 阶段实操（2）：构建一个简单的KV server-基本流程.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/22%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%882%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e5%9f%ba%e6%9c%ac%e6%b5%81%e7%a8%8b.md">22 阶段实操（2）：构建一个简单的KV server-基本流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 类型系统：如何在实战中使用泛型编程？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/23%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%ae%9e%e6%88%98%e4%b8%ad%e4%bd%bf%e7%94%a8%e6%b3%9b%e5%9e%8b%e7%bc%96%e7%a8%8b%ef%bc%9f.md">23 类型系统：如何在实战中使用泛型编程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 类型系统：如何在实战中使用trait object？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/24%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%ae%9e%e6%88%98%e4%b8%ad%e4%bd%bf%e7%94%a8trait%20object%ef%bc%9f.md">24 类型系统：如何在实战中使用trait object？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 类型系统：如何围绕trait来设计和架构系统？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/25%20%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9b%b4%e7%bb%95trait%e6%9d%a5%e8%ae%be%e8%ae%a1%e5%92%8c%e6%9e%b6%e6%9e%84%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">25 类型系统：如何围绕trait来设计和架构系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 阶段实操（3）：构建一个简单的KV server-高级trait技巧.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/26%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%883%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e9%ab%98%e7%ba%a7trait%e6%8a%80%e5%b7%a7.md">26 阶段实操（3）：构建一个简单的KV server-高级trait技巧.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 生态系统：有哪些常有的Rust库可以为我所用？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/27%20%e7%94%9f%e6%80%81%e7%b3%bb%e7%bb%9f%ef%bc%9a%e6%9c%89%e5%93%aa%e4%ba%9b%e5%b8%b8%e6%9c%89%e7%9a%84Rust%e5%ba%93%e5%8f%af%e4%bb%a5%e4%b8%ba%e6%88%91%e6%89%80%e7%94%a8%ef%bc%9f.md">27 生态系统：有哪些常有的Rust库可以为我所用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 网络开发（上）：如何使用Rust处理网络请求？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/28%20%e7%bd%91%e7%bb%9c%e5%bc%80%e5%8f%91%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8Rust%e5%a4%84%e7%90%86%e7%bd%91%e7%bb%9c%e8%af%b7%e6%b1%82%ef%bc%9f.md">28 网络开发（上）：如何使用Rust处理网络请求？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 网络开发（下）：如何使用Rust处理网络请求？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/29%20%e7%bd%91%e7%bb%9c%e5%bc%80%e5%8f%91%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8Rust%e5%a4%84%e7%90%86%e7%bd%91%e7%bb%9c%e8%af%b7%e6%b1%82%ef%bc%9f.md">29 网络开发（下）：如何使用Rust处理网络请求？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 Unsafe Rust：如何用C&#43;&#43;的方式打开Rust？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/30%20Unsafe%20Rust%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8C&#43;&#43;%e7%9a%84%e6%96%b9%e5%bc%8f%e6%89%93%e5%bc%80Rust%ef%bc%9f.md">30 Unsafe Rust：如何用C&#43;&#43;的方式打开Rust？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 FFI：Rust如何和你的语言架起沟通桥梁？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/31%20FFI%ef%bc%9aRust%e5%a6%82%e4%bd%95%e5%92%8c%e4%bd%a0%e7%9a%84%e8%af%ad%e8%a8%80%e6%9e%b6%e8%b5%b7%e6%b2%9f%e9%80%9a%e6%a1%a5%e6%a2%81%ef%bc%9f.md">31 FFI：Rust如何和你的语言架起沟通桥梁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 实操项目：使用PyO3开发Python3模块.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/32%20%e5%ae%9e%e6%93%8d%e9%a1%b9%e7%9b%ae%ef%bc%9a%e4%bd%bf%e7%94%a8PyO3%e5%bc%80%e5%8f%91Python3%e6%a8%a1%e5%9d%97.md">32 实操项目：使用PyO3开发Python3模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 并发处理（上）：从atomics到Channel，Rust都提供了什么工具？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/33%20%e5%b9%b6%e5%8f%91%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%bb%8eatomics%e5%88%b0Channel%ef%bc%8cRust%e9%83%bd%e6%8f%90%e4%be%9b%e4%ba%86%e4%bb%80%e4%b9%88%e5%b7%a5%e5%85%b7%ef%bc%9f.md">33 并发处理（上）：从atomics到Channel，Rust都提供了什么工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 并发处理（下）：从atomics到Channel，Rust都提供了什么工具？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/34%20%e5%b9%b6%e5%8f%91%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%bb%8eatomics%e5%88%b0Channel%ef%bc%8cRust%e9%83%bd%e6%8f%90%e4%be%9b%e4%ba%86%e4%bb%80%e4%b9%88%e5%b7%a5%e5%85%b7%ef%bc%9f.md">34 并发处理（下）：从atomics到Channel，Rust都提供了什么工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 实操项目：如何实现一个基本的MPSC channel？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/35%20%e5%ae%9e%e6%93%8d%e9%a1%b9%e7%9b%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e5%9f%ba%e6%9c%ac%e7%9a%84MPSC%20channel%ef%bc%9f.md">35 实操项目：如何实现一个基本的MPSC channel？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 阶段实操（4）：构建一个简单的KV server-网络处理.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/36%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%884%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e7%bd%91%e7%bb%9c%e5%a4%84%e7%90%86.md">36 阶段实操（4）：构建一个简单的KV server-网络处理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 阶段实操（5）：构建一个简单的KV server-网络安全.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/37%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%885%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e7%bd%91%e7%bb%9c%e5%ae%89%e5%85%a8.md">37 阶段实操（5）：构建一个简单的KV server-网络安全.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 异步处理：Future是什么？它和async_await是什么关系？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/38%20%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86%ef%bc%9aFuture%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e5%ae%83%e5%92%8casync_await%e6%98%af%e4%bb%80%e4%b9%88%e5%85%b3%e7%b3%bb%ef%bc%9f.md">38 异步处理：Future是什么？它和async_await是什么关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 异步处理：async_await内部是怎么实现的？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/39%20%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86%ef%bc%9aasync_await%e5%86%85%e9%83%a8%e6%98%af%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e7%9a%84%ef%bc%9f.md">39 异步处理：async_await内部是怎么实现的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 异步处理：如何处理异步IO？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/40%20%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e5%bc%82%e6%ad%a5IO%ef%bc%9f.md">40 异步处理：如何处理异步IO？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 阶段实操（6）：构建一个简单的KV server-异步处理.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/41%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%886%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86.md">41 阶段实操（6）：构建一个简单的KV server-异步处理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 阶段实操（7）：构建一个简单的KV server-如何做大的重构？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/42%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%887%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e5%a6%82%e4%bd%95%e5%81%9a%e5%a4%a7%e7%9a%84%e9%87%8d%e6%9e%84%ef%bc%9f.md">42 阶段实操（7）：构建一个简单的KV server-如何做大的重构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 生产环境：真实世界下的一个Rust项目包含哪些要素？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/43%20%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%ef%bc%9a%e7%9c%9f%e5%ae%9e%e4%b8%96%e7%95%8c%e4%b8%8b%e7%9a%84%e4%b8%80%e4%b8%aaRust%e9%a1%b9%e7%9b%ae%e5%8c%85%e5%90%ab%e5%93%aa%e4%ba%9b%e8%a6%81%e7%b4%a0%ef%bc%9f.md">43 生产环境：真实世界下的一个Rust项目包含哪些要素？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 数据处理：应用程序和数据如何打交道？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/44%20%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%ef%bc%9a%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e5%92%8c%e6%95%b0%e6%8d%ae%e5%a6%82%e4%bd%95%e6%89%93%e4%ba%a4%e9%81%93%ef%bc%9f.md">44 数据处理：应用程序和数据如何打交道？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 阶段实操（8）：构建一个简单的KV server-配置_测试_监控_CI_CD.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/45%20%e9%98%b6%e6%ae%b5%e5%ae%9e%e6%93%8d%ef%bc%888%ef%bc%89%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84KV%20server-%e9%85%8d%e7%bd%ae_%e6%b5%8b%e8%af%95_%e7%9b%91%e6%8e%a7_CI_CD.md">45 阶段实操（8）：构建一个简单的KV server-配置_测试_监控_CI_CD.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 软件架构：如何用Rust架构复杂系统？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/46%20%e8%bd%af%e4%bb%b6%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8Rust%e6%9e%b6%e6%9e%84%e5%a4%8d%e6%9d%82%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">46 软件架构：如何用Rust架构复杂系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 Rust2021版次问世了！.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20Rust2021%e7%89%88%e6%ac%a1%e9%97%ae%e4%b8%96%e4%ba%86%ef%bc%81.md">加餐 Rust2021版次问世了！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 代码即数据：为什么我们需要宏编程能力？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e4%bb%a3%e7%a0%81%e5%8d%b3%e6%95%b0%e6%8d%ae%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e4%bb%ac%e9%9c%80%e8%a6%81%e5%ae%8f%e7%bc%96%e7%a8%8b%e8%83%bd%e5%8a%9b%ef%bc%9f.md">加餐 代码即数据：为什么我们需要宏编程能力？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 宏编程（上）：用最“笨”的方式撰写宏.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e5%ae%8f%e7%bc%96%e7%a8%8b%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e7%94%a8%e6%9c%80%e2%80%9c%e7%ac%a8%e2%80%9d%e7%9a%84%e6%96%b9%e5%bc%8f%e6%92%b0%e5%86%99%e5%ae%8f.md">加餐 宏编程（上）：用最“笨”的方式撰写宏.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 宏编程（下）：用 syn_quote 优雅地构建宏.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e5%ae%8f%e7%bc%96%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e7%94%a8%20syn_quote%20%e4%bc%98%e9%9b%85%e5%9c%b0%e6%9e%84%e5%bb%ba%e5%ae%8f.md">加餐 宏编程（下）：用 syn_quote 优雅地构建宏.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 愚昧之巅：你的Rust学习常见问题汇总.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e6%84%9a%e6%98%a7%e4%b9%8b%e5%b7%85%ef%bc%9a%e4%bd%a0%e7%9a%84Rust%e5%ad%a6%e4%b9%a0%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e6%b1%87%e6%80%bb.md">加餐 愚昧之巅：你的Rust学习常见问题汇总.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 期中测试：参考实现讲解.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e6%9c%9f%e4%b8%ad%e6%b5%8b%e8%af%95%ef%bc%9a%e5%8f%82%e8%80%83%e5%ae%9e%e7%8e%b0%e8%ae%b2%e8%a7%a3.md">加餐 期中测试：参考实现讲解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 期中测试：来写一个简单的grep命令行.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e6%9c%9f%e4%b8%ad%e6%b5%8b%e8%af%95%ef%bc%9a%e6%9d%a5%e5%86%99%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84grep%e5%91%bd%e4%bb%a4%e8%a1%8c.md">加餐 期中测试：来写一个简单的grep命令行.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 这个专栏你可以怎么学，以及Rust是否值得学？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e8%bf%99%e4%b8%aa%e4%b8%93%e6%a0%8f%e4%bd%a0%e5%8f%af%e4%bb%a5%e6%80%8e%e4%b9%88%e5%ad%a6%ef%bc%8c%e4%bb%a5%e5%8f%8aRust%e6%98%af%e5%90%a6%e5%80%bc%e5%be%97%e5%ad%a6%ef%bc%9f.md">加餐 这个专栏你可以怎么学，以及Rust是否值得学？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助场 开悟之坡（上）：Rust的现状、机遇与挑战.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e5%9c%ba%20%e5%bc%80%e6%82%9f%e4%b9%8b%e5%9d%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aRust%e7%9a%84%e7%8e%b0%e7%8a%b6%e3%80%81%e6%9c%ba%e9%81%87%e4%b8%8e%e6%8c%91%e6%88%98.md">大咖助场 开悟之坡（上）：Rust的现状、机遇与挑战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助场 开悟之坡（下）：Rust的现状、机遇与挑战.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e5%9c%ba%20%e5%bc%80%e6%82%9f%e4%b9%8b%e5%9d%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aRust%e7%9a%84%e7%8e%b0%e7%8a%b6%e3%80%81%e6%9c%ba%e9%81%87%e4%b8%8e%e6%8c%91%e6%88%98.md">大咖助场 开悟之坡（下）：Rust的现状、机遇与挑战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别策划 学习锦囊（一）：听听课代表们怎么说.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%89%b9%e5%88%ab%e7%ad%96%e5%88%92%20%e5%ad%a6%e4%b9%a0%e9%94%a6%e5%9b%8a%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%90%ac%e5%90%ac%e8%af%be%e4%bb%a3%e8%a1%a8%e4%bb%ac%e6%80%8e%e4%b9%88%e8%af%b4.md">特别策划 学习锦囊（一）：听听课代表们怎么说.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别策划 学习锦囊（三）：听听课代表们怎么说.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%89%b9%e5%88%ab%e7%ad%96%e5%88%92%20%e5%ad%a6%e4%b9%a0%e9%94%a6%e5%9b%8a%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%90%ac%e5%90%ac%e8%af%be%e4%bb%a3%e8%a1%a8%e4%bb%ac%e6%80%8e%e4%b9%88%e8%af%b4.md">特别策划 学习锦囊（三）：听听课代表们怎么说.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别策划 学习锦囊（二）：听听课代表们怎么说.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%89%b9%e5%88%ab%e7%ad%96%e5%88%92%20%e5%ad%a6%e4%b9%a0%e9%94%a6%e5%9b%8a%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%90%ac%e5%90%ac%e8%af%be%e4%bb%a3%e8%a1%a8%e4%bb%ac%e6%80%8e%e4%b9%88%e8%af%b4.md">特别策划 学习锦囊（二）：听听课代表们怎么说.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 绝望之谷：改变从学习开始.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e7%bb%9d%e6%9c%9b%e4%b9%8b%e8%b0%b7%ef%bc%9a%e6%94%b9%e5%8f%98%e4%bb%8e%e5%ad%a6%e4%b9%a0%e5%bc%80%e5%a7%8b.md">用户故事 绝望之谷：改变从学习开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 语言不仅是工具，还是思维方式.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e8%af%ad%e8%a8%80%e4%b8%8d%e4%bb%85%e6%98%af%e5%b7%a5%e5%85%b7%ef%bc%8c%e8%bf%98%e6%98%af%e6%80%9d%e7%bb%b4%e6%96%b9%e5%bc%8f.md">用户故事 语言不仅是工具，还是思维方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 永续之原：Rust学习，如何持续精进？.md" href="/%e4%b8%93%e6%a0%8f/%e9%99%88%e5%a4%a9%20%c2%b7%20Rust%20%e7%bc%96%e7%a8%8b%e7%ac%ac%e4%b8%80%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%b0%b8%e7%bb%ad%e4%b9%8b%e5%8e%9f%ef%bc%9aRust%e5%ad%a6%e4%b9%a0%ef%bc%8c%e5%a6%82%e4%bd%95%e6%8c%81%e7%bb%ad%e7%b2%be%e8%bf%9b%ef%bc%9f.md">结束语 永续之原：Rust学习，如何持续精进？.md</a>
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
                            <h1 id="title" data-id="14 类型系统：有哪些必须掌握的trait？" class="title">14 类型系统：有哪些必须掌握的trait？</h1>
                            <div><p>你好，我是陈天。</p>

<p>开发软件系统时，我们弄清楚需求，要对需求进行架构上的分析和设计。在这个过程中，合理地定义和使用 trait，会让代码结构具有很好的扩展性，让系统变得非常灵活。</p>

<p>之前在 get hands dirty 系列中就粗略见识到了 trait 的巨大威力，使用了 From<T>/TryFrom<T> trait 进行类型间的转换（[第 5 讲]），还使用了 Deref trait （[第 6 讲]）让类型在不暴露其内部结构代码的同时，让内部结构的方法可以对外使用。</p>

<p>经过上两讲的学习，相信你现在对trait 的理解就深入了。在实际解决问题的过程中，<strong>用好这些 trait，会让你的代码结构更加清晰，阅读和使用都更加符合 Rust 生态的习惯</strong>。比如数据结构实现了 Debug trait，那么当你想打印数据结构时，就可以用 {:?} 来打印；如果你的数据结构实现了 From<T>，那么，可以直接使用 into() 方法做数据转换。</p>

<h2 id="trait">trait</h2>

<p>Rust 语言的标准库定义了大量的标准 trait，来先来数已经学过的，看看攒了哪些：</p>

<ul>
<li>Clone/Copy trait，约定了数据被深拷贝和浅拷贝的行为；</li>
<li>Read/Write trait，约定了对 I/O 读写的行为；</li>
<li>Iterator，约定了迭代器的行为；</li>
<li>Debug，约定了数据如何被以 debug 的方式显示出来的行为；</li>
<li>Default，约定数据类型的缺省值如何产生的行为；</li>
<li>From<T>/TryFrom<T>，约定了数据间如何转换的行为。</li>
</ul>

<p>我们会再学习几类重要的 trait，包括和内存分配释放相关的 trait、用于区别不同类型协助编译器做类型安全检查的标记 trait、进行类型转换的 trait、操作符相关的 trait，以及 Debug/Display/Default。</p>

<p>在学习这些 trait的过程中，你也可以结合之前讲的内容，有意识地思考一下Rust为什么这么设计，在增进对语言理解的同时，也能写出更加优雅的 Rust 代码。</p>

<h2 id="内存相关-clone-copy-drop">内存相关：Clone/Copy/Drop</h2>

<p>首先来看内存相关的 Clone/Copy/Drop。这三个 trait 在介绍所有权的时候已经学习过，这里我们再深入研究一下它们的定义和使用场景。</p>

<h3 id="clone-trait">Clone trait</h3>

<p>首先看 Clone：</p>

<pre><code>pub trait Clone {
  fn clone(&amp;self) -&gt; Self;

  fn clone_from(&amp;mut self, source: &amp;Self) {
    *self = source.clone()
  }
}
</code></pre>

<p>Clone trait 有两个方法， <code>clone()</code> 和 <code>clone_from()</code> ，后者有缺省实现，所以平时我们只需要实现 <code>clone()</code> 方法即可。你也许会疑惑，这个 <code>clone_from()</code> 有什么作用呢？因为看起来 <code>a.clone_from(&amp;b)</code> ，和 <code>a = b.clone()</code> 是等价的。</p>

<p>其实不是，如果 a 已经存在，在 clone 过程中会分配内存，那么<strong>用 <code>a.clone_from(&amp;b)</code> 可以避免内存分配，提高效率</strong>。</p>

<p>Clone trait 可以通过派生宏直接实现，这样能简化不少代码。如果在你的数据结构里，每一个字段都已经实现了Clone trait，你可以用 <code>#[derive(Clone)]</code> ，看下面的<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=9726c98022668f3249b711719a11bf09" target="_blank">代码</a>，定义了 Developer 结构和 Language 枚举：</p>

<pre><code>#[derive(Clone, Debug)]
struct Developer {
  name: String,
  age: u8,
  lang: Language
}

#[allow(dead_code)]
#[derive(Clone, Debug)]
enum Language {
  Rust,
  TypeScript,
  Elixir,
  Haskell
}

fn main() {
    let dev = Developer {
        name: &quot;Tyr&quot;.to_string(),
        age: 18,
        lang: Language::Rust
    };
    let dev1 = dev.clone();
    println!(&quot;dev: {:?}, addr of dev name: {:p}&quot;, dev, dev.name.as_str());
    println!(&quot;dev1: {:?}, addr of dev1 name: {:p}&quot;, dev1, dev1.name.as_str())
}
</code></pre>

<p>如果没有为 Language 实现 Clone 的话，Developer 的派生宏 Clone 将会编译出错。运行这段代码可以看到，对于 name，也就是 String 类型的 Clone，其堆上的内存也被 Clone 了一份，所以 Clone 是深度拷贝，栈内存和堆内存一起拷贝。</p>

<p>值得注意的是，clone 方法的接口是 &amp;self，这在绝大多数场合下都是适用的，我们在 clone 一个数据时只需要有已有数据的只读引用。但对 Rc<T> 这样在 clone() 时维护引用计数的数据结构，clone() 过程中会改变自己，所以要用 Cell<T> 这样提供内部可变性的结构来进行改变，如果你也有类似的需求，可以参考。</p>

<h3 id="copy-trait">Copy trait</h3>

<p>和 Clone trait 不同的是，Copy trait 没有任何额外的方法，它只是一个标记 trait（marker trait）。它的 trait 定义：</p>

<pre><code>pub trait Copy: Clone {}
</code></pre>

<p>所以看这个定义，如果要实现 Copy trait 的话，必须实现 Clone trait，然后实现一个空的 Copy trait。你是不是有点疑惑：这样不包含任何行为的 trait 有什么用呢？</p>

<p>这样的 trait <strong>虽然没有任何行为，但它可以用作 trait bound 来进行类型安全检查</strong>，所以我们管它叫<strong>标记 trait</strong>。</p>

<p>和 Clone 一样，如果数据结构的所有字段都实现了 Copy，也可以用 <code>#[derive(Copy)]</code> 宏来为数据结构实现 Copy。试着为 Developer 和 Language 加上 Copy：</p>

<pre><code>#[derive(Clone, Copy, Debug)]
struct Developer {
  name: String,
  age: u8,
  lang: Language
}

#[derive(Clone, Copy, Debug)]
enum Language {
  Rust,
  TypeScript,
  Elixir,
  Haskell
}
</code></pre>

<p>这个代码会出错。因为 String 类型没有实现 Copy。 因此，Developer 数据结构只能 clone，无法 copy。我们知道，如果类型实现了 Copy，那么在赋值、函数调用的时候，值会被拷贝，否则所有权会被移动。</p>

<p>所以上面的代码 Developer 类型在做参数传递时，会执行 Move 语义，而 Language 会执行 Copy 语义。</p>

<p>在讲所有权可变/不可变引用的时候提到，不可变引用实现了 Copy，而可变引用 &amp;mut T 没有实现 Copy。为什么是这样？</p>

<p>因为如果可变引用实现了 Copy trait，那么生成一个可变引用然后把它赋值给另一个变量时，就会违背所有权规则：同一个作用域下只能有一个可变引用。可见，Rust 标准库在哪些结构可以 Copy、哪些不可以 Copy 上，有着仔细的考量。</p>

<h3 id="drop-trait">Drop trait</h3>

<p>在内存管理中已经详细探讨过 Drop trait。这里我们再看一下它的定义：</p>

<pre><code>pub trait Drop {
    fn drop(&amp;mut self);
}
</code></pre>

<p>大部分场景无需为数据结构提供 Drop trait，系统默认会依次对数据结构的每个域做 drop。但有两种情况你可能需要手工实现 Drop。</p>

<p>第一种是希望在数据结束生命周期的时候做一些事情，比如记日志。</p>

<p>第二种是需要对资源回收的场景。编译器并不知道你额外使用了哪些资源，也就无法帮助你 drop 它们。比如说锁资源的释放，在 MutexGuard<T> 中实现了 Drop 来释放锁资源：</p>

<pre><code>impl&lt;T: ?Sized&gt; Drop for MutexGuard&lt;'_, T&gt; {
    #[inline]
    fn drop(&amp;mut self) {
        unsafe {
            self.lock.poison.done(&amp;self.poison);
            self.lock.inner.raw_unlock();
        }
    }
}
</code></pre>

<p>需要注意的是，Copy trait 和 Drop trait 是互斥的，两者不能共存，当你尝试为同一种数据类型实现 Copy 时，也实现 Drop，编译器就会报错。这其实很好理解：<strong>Copy是按位做浅拷贝，那么它会默认拷贝的数据没有需要释放的资源；而Drop恰恰是为了释放额外的资源而生的</strong>。</p>

<p>我们还是写一段代码来辅助理解，在代码中，强行用 Box::into_raw 获得堆内存的指针，放入 RawBuffer 结构中，这样就接管了这块堆内存的释放。</p>

<p>虽然 RawBuffer 可以实现 Copy trait，但这样一来就无法实现 Drop trait。如果程序非要这么写，会导致内存泄漏，因为该释放的堆内存没有释放。</p>

<p>但是这个操作不会破坏 Rust 的正确性保证：即便你 Copy 了 N 份 RawBuffer，由于无法实现 Drop trait，RawBuffer 指向的那同一块堆内存不会释放，所以不会出现 use after free 的内存安全问题。（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=76de1040b516b50f671c3abfe71cfb37" target="_blank">代码</a>）</p>

<pre><code>use std::{fmt, slice};

// 注意这里，我们实现了 Copy，这是因为 *mut u8/usize 都支持 Copy
#[derive(Clone, Copy)]
struct RawBuffer {
    // 裸指针用 *const/*mut 来表述，这和引用的 &amp; 不同
    ptr: *mut u8,
    len: usize,
}

impl From&lt;Vec&lt;u8&gt;&gt; for RawBuffer {
    fn from(vec: Vec&lt;u8&gt;) -&gt; Self {
        let slice = vec.into_boxed_slice();
        Self {
            len: slice.len(),
            // into_raw 之后，Box 就不管这块内存的释放了，RawBuffer 需要处理释放
            ptr: Box::into_raw(slice) as *mut u8,
        }
    }
}

// 如果 RawBuffer 实现了 Drop trait，就可以在所有者退出时释放堆内存
// 然后，Drop trait 会跟 Copy trait 冲突，要么不实现 Copy，要么不实现 Drop
// 如果不实现 Drop，那么就会导致内存泄漏，但它不会对正确性有任何破坏
// 比如不会出现 use after free 这样的问题。
// 你可以试着把下面注释去掉，看看会出什么问题
// impl Drop for RawBuffer {
//     #[inline]
//     fn drop(&amp;mut self) {
//         let data = unsafe { Box::from_raw(slice::from_raw_parts_mut(self.ptr, self.len)) };
//         drop(data)
//     }
// }

impl fmt::Debug for RawBuffer {
    fn fmt(&amp;self, f: &amp;mut fmt::Formatter&lt;'_&gt;) -&gt; fmt::Result {
        let data = self.as_ref();
        write!(f, &quot;{:p}: {:?}&quot;, self.ptr, data)
    }
}

impl AsRef&lt;[u8]&gt; for RawBuffer {
    fn as_ref(&amp;self) -&gt; &amp;[u8] {
        unsafe { slice::from_raw_parts(self.ptr, self.len) }
    }
}

fn main() {
    let data = vec![1, 2, 3, 4];

    let buf: RawBuffer = data.into();

    // 因为 buf 允许 Copy，所以这里 Copy 了一份
    use_buffer(buf);

    // buf 还能用
    println!(&quot;buf: {:?}&quot;, buf);
}

fn use_buffer(buf: RawBuffer) {
    println!(&quot;buf to die: {:?}&quot;, buf);

    // 这里不用特意 drop，写出来只是为了说明 Copy 出来的 buf 被 Drop 了
    drop(buf)
}
</code></pre>

<p>对于代码安全来说，内存泄漏危害大？还是 use after free 危害大呢？肯定是后者。Rust 的底线是内存安全，所以两害相权取其轻。</p>

<p>实际上，任何编程语言都无法保证不发生人为的内存泄漏，比如程序在运行时，开发者疏忽了，对哈希表只添加不删除，就会造成内存泄漏。但 Rust 会保证即使开发者疏忽了，也不会出现内存安全问题。</p>

<p>建议你仔细阅读这段代码中的注释，试着把注释掉的 Drop trait 恢复，然后再把代码改得可以编译通过，认真思考一下 Rust 这样做的良苦用心。</p>

<h2 id="标记-trait-sized-send-sync-unpin">标记 trait：Sized/Send/Sync/Unpin</h2>

<p>好，讲完内存相关的主要 trait，来看标记 trait。</p>

<p>刚才我们已经看到了一个标记 trait：Copy。Rust 还支持其它几种标记 trait：<a href="https://doc.rust-lang.org/std/marker/trait.Sized.html" target="_blank">Sized</a>/<a href="https://doc.rust-lang.org/std/marker/trait.Send.html" target="_blank">Send</a>/<a href="https://doc.rust-lang.org/std/marker/trait.Sync.html" target="_blank">Sync</a>/<a href="https://doc.rust-lang.org/std/marker/trait.Unpin.html" target="_blank">Unpin</a>。</p>

<p>Sized trait 用于标记有具体大小的类型。在使用泛型参数时，Rust 编译器会自动为泛型参数加上 Sized 约束，比如下面的 Data<T> 和处理 Data<T> 的函数 process_data：</p>

<pre><code>struct Data&lt;T&gt; {
    inner: T,
}

fn process_data&lt;T&gt;(data: Data&lt;T&gt;) {
    todo!();
}
</code></pre>

<p>它等价于：</p>

<pre><code>struct Data&lt;T: Sized&gt; {
    inner: T,
}

fn process_data&lt;T: Sized&gt;(data: Data&lt;T&gt;) {
    todo!();
}
</code></pre>

<p>大部分时候，我们都希望能自动添加这样的约束，因为这样定义出的泛型结构，在编译期，大小是固定的，可以作为参数传递给函数。如果没有这个约束，T 是大小不固定的类型， process_data 函数会无法编译。</p>

<p>但是这个自动添加的约束有时候不太适用，<strong>在少数情况下，需要 T 是可变类型的，怎么办？Rust 提供了 ?Sized 来摆脱这个约束</strong>。</p>

<p>如果开发者显式定义了<code>T: ?Sized</code>，那么 T 就可以是任意大小。如果你对（[第12讲]）之前说的 Cow 还有印象，可能会记得 Cow 中泛型参数 B 的约束是 ?Sized：</p>

<pre><code>pub enum Cow&lt;'a, B: ?Sized + 'a&gt; where B: ToOwned,
{
    // 借用的数据
    Borrowed(&amp;'a B),
    // 拥有的数据
    Owned(&lt;B as ToOwned&gt;::Owned),
}
</code></pre>

<p>这样 B 就可以是 [T] 或者 str 类型，大小都是不固定的。要注意 Borrowed(&amp;&lsquo;a B) 大小是固定的，因为它内部是对 B 的一个引用，而引用的大小是固定的。</p>

<h3 id="send-sync">Send/Sync</h3>

<p>说完了 Sized，我们再来看 Send/Sync，定义是：</p>

<pre><code>pub unsafe auto trait Send {}
pub unsafe auto trait Sync {}
</code></pre>

<p>这两个 trait 都是 unsafe auto trait，auto 意味着编译器会在合适的场合，自动为数据结构添加它们的实现，而 unsafe 代表实现的这个 trait 可能会违背 Rust 的内存安全准则，如果开发者手工实现这两个 trait ，要自己为它们的安全性负责。</p>

<p>Send/Sync 是 Rust 并发安全的基础：</p>

<ul>
<li>如果一个类型 T 实现了 Send trait，意味着 T 可以安全地从一个线程移动到另一个线程，也就是说所有权可以在线程间移动。</li>
<li>如果一个类型 T 实现了 Sync trait，则意味着 &amp;T 可以安全地在多个线程中共享。一个类型 T 满足 Sync trait，当且仅当 &amp;T 满足 Send trait。</li>
</ul>

<p>对于 Send/Sync 在线程安全中的作用，可以这么看，<strong>如果一个类型T: Send，那么 T 在某个线程中的独占访问是线程安全的；如果一个类型 T: Sync，那么 T 在线程间的只读共享是安全的</strong>。</p>

<p>对于我们自己定义的数据结构，如果其内部的所有域都实现了 Send/Sync，那么这个数据结构会被自动添加 Send/Sync 。基本上原生数据结构都支持 Send/Sync，也就是说，绝大多数自定义的数据结构都是满足 Send/Sync 的。标准库中，不支持 Send/Sync 的数据结构主要有：</p>

<ul>
<li>裸指针 *const T/*mut T。它们是不安全的，所以既不是 Send 也不是 Sync。</li>
<li>UnsafeCell<T> 不支持 Sync。也就是说，任何使用了 Cell 或者 RefCell 的数据结构不支持 Sync。</li>
<li>引用计数 Rc 不支持 Send 也不支持 Sync。所以 Rc 无法跨线程。</li>
</ul>

<p>之前介绍过 Rc/RefCell（[第9讲]），我们来看看，如果尝试跨线程使用 Rc/RefCell，会发生什么。在 Rust 下，如果想创建一个新的线程，需要使用 <a href="https://doc.rust-lang.org/std/thread/fn.spawn.html" target="_blank">std::thread::spawn</a>：</p>

<pre><code>pub fn spawn&lt;F, T&gt;(f: F) -&gt; JoinHandle&lt;T&gt; 
where
    F: FnOnce() -&gt; T,
    F: Send + 'static,
    T: Send + 'static,
</code></pre>

<p>它的参数是一个闭包（后面会讲），这个闭包需要 Send + &lsquo;static：</p>

<ul>
<li>&lsquo;static 意思是闭包捕获的自由变量必须是一个拥有所有权的类型，或者是一个拥有静态生命周期的引用；</li>
<li>Send 意思是，这些被捕获自由变量的所有权可以从一个线程移动到另一个线程。</li>
</ul>

<p>从这个接口上，可以得出结论：如果在线程间传递 Rc，是无法编译通过的，因为 <a href="https://doc.rust-lang.org/std/rc/struct.Rc.html#impl-Send" target="_blank">Rc 的实现不支持 Send 和 Sync</a>。写段代码验证一下（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=9892c9cd4baa26dcfcca99e4e4869cc5" target="_blank">代码</a>）：</p>

<pre><code>// Rc 既不是 Send，也不是 Sync
fn rc_is_not_send_and_sync() {
    let a = Rc::new(1);
    let b = a.clone();
    let c = a.clone();
    thread::spawn(move || {
        println!(&quot;c= {:?}&quot;, c);
    });
}
</code></pre>

<p>果然，这段代码不通过。-
<img src="assets/131b21c05850e8f5070d952a777613e6.jpg" alt="" /></p>

<p>那么，RefCell<T> 可以在线程间转移所有权么？<a href="https://doc.rust-lang.org/std/cell/struct.RefCell.html#impl-Send" target="_blank">RefCell 实现了 Send，但没有实现 Sync</a>，所以，看起来是可以工作的（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=1a820a1bd4eca214956e85a1333e5df0" target="_blank">代码</a>）：</p>

<pre><code>fn refcell_is_send() {
    let a = RefCell::new(1);
    thread::spawn(move || {
        println!(&quot;a= {:?}&quot;, a);
    });
}
</code></pre>

<p>验证一下发现，这是 OK 的。</p>

<p>既然 Rc 不能 Send，我们无法跨线程使用 Rc<RefCell<T>&gt; 这样的数据，那么使用<a href="https://doc.rust-lang.org/std/sync/struct.Arc.html#impl-Send" target="_blank">支持 Send/Sync 的 Arc</a>呢，使用 Arc<RefCell<T>&gt; 来获得，一个可以在多线程间共享，且可以修改的类型，可以么（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=16b78ecc207cbae4a511a316681ad49e" target="_blank">代码</a>）？</p>

<pre><code>// RefCell 现在有多个 Arc 持有它，虽然 Arc 是 Send/Sync，但 RefCell 不是 Sync
fn refcell_is_not_sync() {
    let a = Arc::new(RefCell::new(1));
    let b = a.clone();
    let c = a.clone();
    thread::spawn(move || {
        println!(&quot;c= {:?}&quot;, c);
    });
}
</code></pre>

<p>不可以。</p>

<p>因为 Arc 内部的数据是共享的，需要支持 Sync 的数据结构，但是RefCell 不是 Sync，编译失败。所以在多线程情况下，我们只能使用支持 Send/Sync 的 Arc ，和 Mutex 一起，构造一个可以在多线程间共享且可以修改的类型（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=e20084ea53dbd030e3f75ce0b07b6421" target="_blank">代码</a>）：</p>

<pre><code>use std::{
    sync::{Arc, Mutex},
    thread,
};

// Arc&lt;Mutex&lt;T&gt;&gt; 可以多线程共享且修改数据
fn arc_mutext_is_send_sync() {
    let a = Arc::new(Mutex::new(1));
    let b = a.clone();
    let c = a.clone();
    let handle = thread::spawn(move || {
        let mut g = c.lock().unwrap();
        *g += 1;
    });

    {
        let mut g = b.lock().unwrap();
        *g += 1;
    }

    handle.join().unwrap();
    println!(&quot;a= {:?}&quot;, a);
}

fn main() {
    arc_mutext_is_send_sync();
}
</code></pre>

<p>这几段代码建议你都好好阅读和运行一下，对于编译出错的情况，仔细看看编译器给出的错误，会帮助你理解好 Send/Sync trait 以及它们如何保证并发安全。</p>

<p>最后一个标记 trait Unpin，是用于自引用类型的，在后面讲到 Future trait 时，再详细讲这个 trait。</p>

<h2 id="类型转换相关-from-t-into-t-asref-t-asmut-t">类型转换相关：From<T>/Into<T>/AsRef<T>/AsMut<T></h2>

<p>好，学完了标记 trait，来看看和类型转换相关的 trait。在软件开发的过程中，我们经常需要在某个上下文中，把一种数据结构转换成另一种数据结构。</p>

<p>不过转换有很多方式，看下面的代码，你觉得哪种方式更好呢？</p>

<pre><code>// 第一种方法，为每一种转换提供一个方法
// 把字符串 s 转换成 Path
let v = s.to_path();
// 把字符串 s 转换成 u64
let v = s.to_u64();

// 第二种方法，为 s 和要转换的类型之间实现一个 Into&lt;T&gt; trait
// v 的类型根据上下文得出
let v = s.into();
// 或者也可以显式地标注 v 的类型
let v: u64 = s.into();
</code></pre>

<p>第一种方式，在类型 T 的实现里，要为每一种可能的转换提供一个方法；第二种，我们为类型 T 和类型 U 之间的转换实现一个数据转换 trait，这样可以用同一个方法来实现不同的转换。</p>

<p>显然，第二种方法要更好，因为它符合软件开发的开闭原则（Open-Close Principle），“<strong>软件中的对象（类、模块、函数等等）对扩展是开放的，但是对修改是封闭的</strong>”。</p>

<p>在第一种方式下，未来每次要添加对新类型的转换，都要重新修改类型 T 的实现，而第二种方式，我们只需要添加一个对于数据转换 trait 的新实现即可。</p>

<p>基于这个思路，对值类型的转换和对引用类型的转换，Rust 提供了两套不同的 trait：</p>

<ul>
<li>值类型到值类型的转换：From<T>/Into<T>/TryFrom<T>/TryInto<T></li>
<li>引用类型到引用类型的转换：AsRef<T>/AsMut<T></li>
</ul>

<h3 id="from-t-into-t">From<T>/Into<T></h3>

<p>先看 <a href="https://doc.rust-lang.org/std/convert/trait.From.html" target="_blank">From<T></a> 和 <a href="https://doc.rust-lang.org/std/convert/trait.Into.html" target="_blank">Into<T></a>。这两个 trait 的定义如下：</p>

<pre><code>pub trait From&lt;T&gt; {
    fn from(T) -&gt; Self;
}

pub trait Into&lt;T&gt; {
    fn into(self) -&gt; T;
}
</code></pre>

<p>在实现 From<T> 的时候会自动实现 Into<T>。这是因为：</p>

<pre><code>// 实现 From 会自动实现 Into
impl&lt;T, U&gt; Into&lt;U&gt; for T where U: From&lt;T&gt; {
    fn into(self) -&gt; U {
        U::from(self)
    }
}
</code></pre>

<p>所以大部分情况下，只用实现 From<T>，然后这两种方式都能做数据转换，比如：</p>

<pre><code>let s = String::from(&quot;Hello world!&quot;);
let s: String = &quot;Hello world!&quot;.into();
</code></pre>

<p>这两种方式是等价的，怎么选呢？From<T> 可以根据上下文做类型推导，使用场景更多；而且因为实现了 From<T> 会自动实现 Into<T>，反之不会。<strong>所以需要的时候，不要去实现 Into<T>，只要实现 From<T> 就好了</strong>。</p>

<p>此外，From<T> 和 Into<T> 还是自反的：把类型 T 的值转换成类型 T，会直接返回。这是因为标准库有如下的实现：</p>

<pre><code>// From（以及 Into）是自反的
impl&lt;T&gt; From&lt;T&gt; for T {
    fn from(t: T) -&gt; T {
        t
    }
}
</code></pre>

<p>有了 From<T> 和 Into<T>，很多函数的接口就可以变得灵活，比如函数如果接受一个 IpAddr 为参数，我们可以使用 Into<IpAddr> 让更多的类型可以被这个函数使用，看下面的<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=f8be081138a8bb2c736e30badcc5ae41" target="_blank">代码</a>：</p>

<pre><code>use std::net::{IpAddr, Ipv4Addr, Ipv6Addr};

fn print(v: impl Into&lt;IpAddr&gt;) {
    println!(&quot;{:?}&quot;, v.into());
}

fn main() {
    let v4: Ipv4Addr = &quot;2.2.2.2&quot;.parse().unwrap();
    let v6: Ipv6Addr = &quot;::1&quot;.parse().unwrap();

    // IPAddr 实现了 From&lt;[u8; 4]，转换 IPv4 地址
    print([1, 1, 1, 1]);
    // IPAddr 实现了 From&lt;[u16; 8]，转换 IPv6 地址
    print([0xfe80, 0, 0, 0, 0xaede, 0x48ff, 0xfe00, 0x1122]);
    // IPAddr 实现了 From&lt;Ipv4Addr&gt;
    print(v4);
    // IPAddr 实现了 From&lt;Ipv6Addr&gt;
    print(v6);
}
</code></pre>

<p>所以，合理地使用 From<T>/Into<T>，可以让代码变得简洁，符合 Rust 可读性强的风格，更符合开闭原则。</p>

<p>注意，如果你的数据类型在转换过程中有可能出现错误，可以使用 <a href="https://doc.rust-lang.org/std/convert/trait.TryFrom.html" target="_blank">TryFrom<T></a> 和 <a href="https://doc.rust-lang.org/std/convert/trait.TryInto.html" target="_blank">TryInto<T></a>，它们的用法和 From<T>/Into<T> 一样，只是 trait 内多了一个关联类型 Error，且返回的结果是 Result<T, Self::Error>。</p>

<h3 id="asref-t-asmut-t">AsRef<T>/AsMut<T></h3>

<p>搞明白了 From<T>/Into<T> 后，AsRef<T> 和 AsMut<T> 就很好理解了，用于从引用到引用的转换。还是先看它们的定义：</p>

<pre><code>pub trait AsRef&lt;T&gt; where T: ?Sized {
    fn as_ref(&amp;self) -&gt; &amp;T;
}

pub trait AsMut&lt;T&gt; where T: ?Sized {
    fn as_mut(&amp;mut self) -&gt; &amp;mut T;
}
</code></pre>

<p>在 trait 的定义上，都允许 T 使用大小可变的类型，如 str、[u8] 等。AsMut<T> 除了使用可变引用生成可变引用外，其它都和 AsRef<T> 一样，所以我们重点看 AsRef<T>。</p>

<p>看标准库中打开文件的接口 <a href="https://doc.rust-lang.org/std/fs/struct.File.html#method.open" target="_blank">std::fs::File::open</a>：</p>

<pre><code>pub fn open&lt;P: AsRef&lt;Path&gt;&gt;(path: P) -&gt; Result&lt;File&gt;
</code></pre>

<p>它的参数 path 是符合 AsRef<Path> 的类型，所以，你可以为这个参数传入 String、&amp;str、PathBuf、Path 等类型。而且，当你使用 path.as_ref() 时，会得到一个 &amp;Path。</p>

<p>来写一段代码体验一下 AsRef<T> 的使用和实现（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=176092de75680a60821d6523e6340773" target="_blank">代码</a>）：</p>

<pre><code>#[allow(dead_code)]
enum Language {
    Rust,
    TypeScript,
    Elixir,
    Haskell,
}

impl AsRef&lt;str&gt; for Language {
    fn as_ref(&amp;self) -&gt; &amp;str {
        match self {
            Language::Rust =&gt; &quot;Rust&quot;,
            Language::TypeScript =&gt; &quot;TypeScript&quot;,
            Language::Elixir =&gt; &quot;Elixir&quot;,
            Language::Haskell =&gt; &quot;Haskell&quot;,
        }
    }
}

fn print_ref(v: impl AsRef&lt;str&gt;) {
    println!(&quot;{}&quot;, v.as_ref());
}

fn main() {
    let lang = Language::Rust;
    // &amp;str 实现了 AsRef&lt;str&gt;
    print_ref(&quot;Hello world!&quot;);
    // String 实现了 AsRef&lt;str&gt;
    print_ref(&quot;Hello world!&quot;.to_string());
    // 我们自己定义的 enum 也实现了 AsRef&lt;str&gt;
    print_ref(lang);
}
</code></pre>

<p>现在对在 Rust 下，如何使用 From/Into/AsRef/AsMut 进行类型间转换，有了深入了解，未来我们还会在实战中使用到这些 trait。</p>

<p>刚才的小例子中要额外说明一下的是，如果你的代码出现 v.as_ref().clone() 这样的语句，也就是说你要对 v 进行引用转换，然后又得到了拥有所有权的值，那么你应该实现 From<T>，然后做 v.into()。</p>

<h2 id="操作符相关-deref-derefmut">操作符相关：Deref/DerefMut</h2>

<p>操作符相关的 trait ，上一讲我们已经看到了 Add<Rhs> trait，它允许你重载加法运算符。Rust 为所有的运算符都提供了 trait，你可以为自己的类型重载某些操作符。这里用下图简单概括一下，更详细的信息你可以阅读<a href="https://doc.rust-lang.org/std/ops/index.html" target="_blank">官方文档</a>。</p>

<p><img src="assets/a28619aae702e186aa115af94300dc19.jpg" alt="" /></p>

<p>今天重点要介绍的操作符是 <a href="https://doc.rust-lang.org/std/ops/trait.Deref.html" target="_blank">Deref</a> 和 <a href="https://doc.rust-lang.org/std/ops/trait.DerefMut.html" target="_blank">DerefMut</a>。来看它们的定义：</p>

<pre><code>pub trait Deref {
    // 解引用出来的结果类型
    type Target: ?Sized;
    fn deref(&amp;self) -&gt; &amp;Self::Target;
}

pub trait DerefMut: Deref {
    fn deref_mut(&amp;mut self) -&gt; &amp;mut Self::Target;
}
</code></pre>

<p>可以看到，DerefMut “继承”了 Deref，只是它额外提供了一个 deref_mut 方法，用来获取可变的解引用。所以这里重点学习 Deref。</p>

<p>对于普通的引用，解引用很直观，因为它只有一个指向值的地址，从这个地址可以获取到所需要的值，比如下面的例子：</p>

<pre><code>let mut x = 42;
let y = &amp;mut x;
// 解引用，内部调用 DerefMut（其实现就是 *self）
*y += 1;
</code></pre>

<p>但对智能指针来说，拿什么域来解引用就不那么直观了，我们来看之前学过的 Rc 是怎么实现 Deref 的：</p>

<pre><code>impl&lt;T: ?Sized&gt; Deref for Rc&lt;T&gt; {
    type Target = T;

    fn deref(&amp;self) -&gt; &amp;T {
        &amp;self.inner().value
    }
}
</code></pre>

<p>可以看到，它最终指向了堆上的 RcBox 内部的 value 的地址，然后如果对其解引用的话，得到了 value 对应的值。以下图为例，最终打印出 v = 1。-
<img src="assets/5068f84af27d696f6a062c5a2f43f4d1.jpg" alt="" /></p>

<p>从图中还可以看到，Deref 和 DerefMut 是自动调用的，*b 会被展开为 *(b.deref())。</p>

<p>在 Rust 里，绝大多数智能指针都实现了 Deref，我们也可以为自己的数据结构实现 Deref。看一个例子（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=084d38d49c6b29d6074a2c4570551601" target="_blank">代码</a>）：</p>

<pre><code>use std::ops::{Deref, DerefMut};

#[derive(Debug)]
struct Buffer&lt;T&gt;(Vec&lt;T&gt;);

impl&lt;T&gt; Buffer&lt;T&gt; {
    pub fn new(v: impl Into&lt;Vec&lt;T&gt;&gt;) -&gt; Self {
        Self(v.into())
    }
}

impl&lt;T&gt; Deref for Buffer&lt;T&gt; {
    type Target = [T];

    fn deref(&amp;self) -&gt; &amp;Self::Target {
        &amp;self.0
    }
}

impl&lt;T&gt; DerefMut for Buffer&lt;T&gt; {
    fn deref_mut(&amp;mut self) -&gt; &amp;mut Self::Target {
        &amp;mut self.0
    }
}

fn main() {
    let mut buf = Buffer::new([1, 3, 2, 4]);
    // 因为实现了 Deref 和 DerefMut，这里 buf 可以直接访问 Vec&lt;T&gt; 的方法
    // 下面这句相当于：(&amp;mut buf).deref_mut().sort()，也就是 (&amp;mut buf.0).sort()
    buf.sort();
    println!(&quot;buf: {:?}&quot;, buf);
}
</code></pre>

<p>但是在这个例子里，数据结构 Buffer<T> 包裹住了 Vec<T>，但这样一来，原本 Vec<T> 实现了的很多方法，现在使用起来就很不方便，需要用 buf.0 来访问。怎么办？</p>

<p><strong>可以实现 Deref 和 DerefMut，这样在解引用的时候，直接访问到 buf.0</strong>，省去了代码的啰嗦和数据结构内部字段的隐藏。</p>

<p>在这段代码里，还有一个值得注意的地方：写 buf.sort() 的时候，并没有做解引用的操作，为什么会相当于访问了 buf.0.sort() 呢？这是因为 sort() 方法第一个参数是 &amp;mut self，此时 Rust 编译器会强制做 Deref/DerefMut 的解引用，所以这相当于 (*(&amp;mut buf)).sort()。</p>

<h2 id="其它-debug-display-default">其它：Debug/Display/Default</h2>

<p>现在我们对运算符相关的 trait 有了足够的了解，最后来看看其它一些常用的 trait：<a href="https://doc.rust-lang.org/std/fmt/trait.Debug.html" target="_blank">Debug</a>/<a href="https://doc.rust-lang.org/std/fmt/trait.Display.html" target="_blank">Display</a>/<a href="https://doc.rust-lang.org/std/default/trait.Default.html" target="_blank">Default</a>。</p>

<p>先看 Debug/Display，它们的定义如下：</p>

<pre><code>pub trait Debug {
    fn fmt(&amp;self, f: &amp;mut Formatter&lt;'_&gt;) -&gt; Result&lt;(), Error&gt;;
}

pub trait Display {
    fn fmt(&amp;self, f: &amp;mut Formatter&lt;'_&gt;) -&gt; Result&lt;(), Error&gt;;
}
</code></pre>

<p>可以看到，Debug 和 Display 两个 trait 的签名一样，都接受一个 &amp;self 和一个 &amp;mut Formatter。那为什么要有两个一样的 trait 呢？</p>

<p>这是因为 <strong>Debug 是为开发者调试打印数据结构所设计的，而 Display 是给用户显示数据结构所设计的</strong>。这也是为什么 Debug trait 的实现可以通过派生宏直接生成，而 Display 必须手工实现。在使用的时候，Debug 用 {:?} 来打印，Display 用 {} 打印。</p>

<p>最后看 Default trait。它的定义如下：</p>

<pre><code>pub trait Default {
    fn default() -&gt; Self;
}
</code></pre>

<p>Default trait 用于为类型提供缺省值。它也可以通过 derive 宏 #[derive(Default)] 来生成实现，前提是类型中的每个字段都实现了 Default trait。在初始化一个数据结构时，我们可以部分初始化，然后剩余的部分使用 Default::default()。</p>

<p>Debug/Display/Default 如何使用，统一看个例子（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=77bdb7c373ad7762bf0e3c2081c96719" target="_blank">代码</a>）：</p>

<pre><code>use std::fmt;
// struct 可以 derive Default，但我们需要所有字段都实现了 Default
#[derive(Clone, Debug, Default)]
struct Developer {
    name: String,
    age: u8,
    lang: Language,
}

// enum 不能 derive Default
#[allow(dead_code)]
#[derive(Clone, Debug)]
enum Language {
    Rust,
    TypeScript,
    Elixir,
    Haskell,
}

// 手工实现 Default
impl Default for Language {
    fn default() -&gt; Self {
        Language::Rust
    }
}

impl Developer {
    pub fn new(name: &amp;str) -&gt; Self {
        // 用 ..Default::default() 为剩余字段使用缺省值
        Self {
            name: name.to_owned(),
            ..Default::default()
        }
    }
}

impl fmt::Display for Developer {
    fn fmt(&amp;self, f: &amp;mut fmt::Formatter&lt;'_&gt;) -&gt; fmt::Result {
        write!(
            f,
            &quot;{}({} years old): {:?} developer&quot;,
            self.name, self.age, self.lang
        )
    }
}

fn main() {
    // 使用 T::default()
    let dev1 = Developer::default();
    // 使用 Default::default()，但此时类型无法通过上下文推断，需要提供类型
    let dev2: Developer = Default::default();
    // 使用 T::new
    let dev3 = Developer::new(&quot;Tyr&quot;);
    println!(&quot;dev1: {}\\ndev2: {}\\ndev3: {:?}&quot;, dev1, dev2, dev3);
}
</code></pre>

<p>它们实现起来非常简单，你可以看文中的代码。</p>

<h2 id="小结">小结</h2>

<p>今天介绍了内存管理、类型转换、操作符、数据显示等相关的基本 trait，还介绍了标记 trait，它是一种特殊的 trait，主要是用于协助编译器检查类型安全。-
<img src="assets/c40e3efef2bec9140c95054547958a5e.jpg" alt="" /></p>

<p>在我们使用 Rust 开发时，trait 占据了非常核心的地位。<strong>一个设计良好的 trait 可以大大提升整个系统的可用性和扩展性</strong>。</p>

<p>很多优秀的第三方库，都围绕着 trait 展开它们的能力，比如上一讲提到的 tower-service 中的 <a href="https://docs.rs/tower-service/0.3.1/tower_service/trait.Service.html" target="_blank">Service trait</a>，再比如你日后可能会经常使用到的 parser combinator 库 <a href="https://docs.rs/nom/6.2.1/nom/" target="_blank">nom</a> 的 <a href="https://docs.rs/nom/6.2.1/nom/trait.Parser.html" target="_blank">Parser trait</a>。</p>

<p>因为 trait 实现了延迟绑定。不知道你是否还记得，之前串讲编程基础概念的时候，就谈到了延迟绑定。在软件开发中，延迟绑定会带来极大的灵活性。</p>

<p>从数据的角度看，数据结构是具体数据的延迟绑定，泛型结构是具体数据结构的延迟绑定；从代码的角度看，函数是一组实现某个功能的表达式的延迟绑定，泛型函数是函数的延迟绑定。那么 trait 是什么的延迟绑定呢？</p>

<p><strong>trait 是行为的延迟绑定</strong>。我们可以在不知道具体要处理什么数据结构的前提下，先通过 trait 把系统的很多行为约定好。这也是为什么开头解释标准trait时，频繁用到了“约定……行为”。</p>

<p>相信通过今天的学习，你能对 trait 有更深刻的认识，在撰写自己的数据类型时，就能根据需要实现这些 trait。</p>

<h3 id="思考题">思考题</h3>

<ol>
<li><p>Vec<T> 可以实现 Copy trait 么？为什么？-</p></li>

<li><p>在使用 Arc<Mutex<T>&gt; 时，为什么下面这段代码可以直接使用 shared.lock()？</p>

<p>use std::sync::{Arc, Mutex};
let shared = Arc::new(Mutex::new(1));
let mut g = shared.lock().unwrap();
*g += 1;</p></li>
</ol>

<p>3.有余力的同学可以尝试一下，为下面的 List<T> 类型实现 Index，使得所有的测试都能通过。这段代码使用了 std::collections::LinkedList，你可以参考<a href="https://doc.rust-lang.org/std/collections/linked_list/struct.LinkedList.html" target="_blank">官方文档</a>阅读它支持的方法（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=c7ddd7b647ef42753cc86a2a86e5a753" target="_blank">代码</a>）：</p>

<pre><code>use std::{
    collections::LinkedList,
    ops::{Deref, DerefMut, Index},
};
struct List&lt;T&gt;(LinkedList&lt;T&gt;);

impl&lt;T&gt; Deref for List&lt;T&gt; {
    type Target = LinkedList&lt;T&gt;;

    fn deref(&amp;self) -&gt; &amp;Self::Target {
        &amp;self.0
    }
}

impl&lt;T&gt; DerefMut for List&lt;T&gt; {
    fn deref_mut(&amp;mut self) -&gt; &amp;mut Self::Target {
        &amp;mut self.0
    }
}

impl&lt;T&gt; Default for List&lt;T&gt; {
    fn default() -&gt; Self {
        Self(Default::default())
    }
}

impl&lt;T&gt; Index&lt;isize&gt; for List&lt;T&gt; {
    type Output = T;

    fn index(&amp;self, index: isize) -&gt; &amp;Self::Output {
        todo!();
    }
}

#[test]
fn it_works() {
    let mut list: List&lt;u32&gt; = List::default();
    for i in 0..16 {
        list.push_back(i);
    }

    assert_eq!(list[0], 0);
    assert_eq!(list[5], 5);
    assert_eq!(list[15], 15);
    assert_eq!(list[16], 0);
    assert_eq!(list[-1], 15);
    assert_eq!(list[128], 0);
    assert_eq!(list[-128], 0);
}
</code></pre>

<p>今天你已经完成了Rust学习的第14次打卡，坚持学习，如果你觉得有收获，也欢迎分享给身边的朋友，邀TA一起讨论。我们下节课见～</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#7c10101045484d4d4c4b3c1b111d1510521f1311" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a4ceb299f7755',t:'MTczNDEzODQ4Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
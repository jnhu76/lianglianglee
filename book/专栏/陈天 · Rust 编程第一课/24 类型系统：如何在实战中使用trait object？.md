<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=24&#32;类型系统：如何在实战中使用trait&#32;object？>
        <link rel="icon" href="/static/favicon.png">
        <title>24 类型系统：如何在实战中使用trait object？ </title>
        
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
                            <h1 id="title" data-id="24 类型系统：如何在实战中使用trait object？" class="title">24 类型系统：如何在实战中使用trait object？</h1>
                            <div><p>你好，我是陈天。</p>

<p>今天我们来看看 trait object 是如何在实战中使用的。</p>

<p>照例先来回顾一下 trait object。当我们在运行时想让某个具体类型，只表现出某个 trait 的行为，可以通过将其赋值给一个 dyn T，无论是 &amp;dyn T，还是 Box<dyn T>，还是 Arc<dyn T>，都可以，这里，T 是当前数据类型实现的某个 trait。此时，原有的类型被抹去，Rust 会创建一个 trait object，并为其分配满足该 trait 的 vtable。</p>

<p>你可以再阅读一下[第 13 讲]的这个图，来回顾 trait object 是怎么回事：-
<img src="assets/4900097edab0yye11233e14ef857be1d.jpg" alt="" /></p>

<p>在编译 dyn T 时，Rust 会为使用了 trait object 类型的 trait 实现，生成相应的 vtable，放在可执行文件中（一般在 TEXT 或 RODATA 段）：-
<img src="assets/9ddeafee9740e891f6bf9c1584e6905e.jpg" alt="" /></p>

<p>这样，当 trait object 调用 trait 的方法时，它会先从 vptr 中找到对应的 vtable，进而找到对应的方法来执行。</p>

<p>使用 trait object 的好处是，<strong>当在某个上下文中需要满足某个 trait 的类型，且这样的类型可能有很多，当前上下文无法确定会得到哪一个类型时，我们可以用 trait object 来统一处理行为</strong>。和泛型参数一样，trait object 也是一种延迟绑定，它让决策可以延迟到运行时，从而得到最大的灵活性。</p>

<p>当然，有得必有失。trait object 把决策延迟到运行时，带来的后果是执行效率的打折。在 Rust 里，函数或者方法的执行就是一次跳转指令，而 trait object 方法的执行还多一步，它涉及额外的内存访问，才能得到要跳转的位置再进行跳转，执行的效率要低一些。</p>

<p>此外，如果要把 trait object 作为返回值返回，或者要在线程间传递 trait object，都免不了使用 Box<dyn T> 或者 Arc<dyn T>，会带来额外的堆分配的开销。</p>

<p>好，对 trait object 的回顾就到这里，如果你对它还一知半解，请复习 [13 讲]，并且阅读 Rust book 里的：<a href="https://doc.rust-lang.org/book/ch17-02-trait-objects.html" target="_blank">Using Trait Objects that allow for values of different types</a>。接下来我们讲讲实战中 trait object 的主要使用场景。</p>

<h2 id="在函数中使用-trait-object">在函数中使用 trait object</h2>

<p>我们可以在函数的参数或者返回值中使用 trait object。</p>

<p>先来看在参数中使用 trait object。下面的代码构建了一个 Executor trait，并对比做静态分发的 impl Executor、做动态分发的 &amp;dyn Executor 和 Box<dyn Executor> 这几种不同的参数的使用：</p>

<pre><code>use std::{error::Error, process::Command};

pub type BoxedError = Box&lt;dyn Error + Send + Sync&gt;;

pub trait Executor {
    fn run(&amp;self) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt;;
}

pub struct Shell&lt;'a, 'b&gt; {
    cmd: &amp;'a str,
    args: &amp;'b [&amp;'a str],
}

impl&lt;'a, 'b&gt; Shell&lt;'a, 'b&gt; {
    pub fn new(cmd: &amp;'a str, args: &amp;'b [&amp;'a str]) -&gt; Self {
        Self { cmd, args }
    }
}

impl&lt;'a, 'b&gt; Executor for Shell&lt;'a, 'b&gt; {
    fn run(&amp;self) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
        let output = Command::new(self.cmd).args(self.args).output()?;
        Ok(output.status.code())
    }
}

/// 使用泛型参数
pub fn execute_generics(cmd: &amp;impl Executor) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
    cmd.run()
}

/// 使用 trait object: &amp;dyn T
pub fn execute_trait_object(cmd: &amp;dyn Executor) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
    cmd.run()
}

/// 使用 trait object: Box&lt;dyn T&gt;
pub fn execute_boxed_trait_object(cmd: Box&lt;dyn Executor&gt;) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
    cmd.run()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn shell_shall_work() {
        let cmd = Shell::new(&quot;ls&quot;, &amp;[]);
        let result = cmd.run().unwrap();
        assert_eq!(result, Some(0));
    }

    #[test]
    fn execute_shall_work() {
        let cmd = Shell::new(&quot;ls&quot;, &amp;[]);

        let result = execute_generics(&amp;cmd).unwrap();
        assert_eq!(result, Some(0));
        let result = execute_trait_object(&amp;cmd).unwrap();
        assert_eq!(result, Some(0));
        let boxed = Box::new(cmd);
        let result = execute_boxed_trait_object(boxed).unwrap();
        assert_eq!(result, Some(0));
    }
}
</code></pre>

<p>其中，impl Executor 使用的是泛型参数的简化版本，而 &amp;dyn Executor 和 Box<dyn Executor> 是 trait object，前者在栈上，后者分配在堆上。值得注意的是，分配在堆上的 trait object 也可以作为返回值返回，比如示例中的 Result<Option<i32>, BoxedError&gt; 里使用了 trait object。</p>

<p>这里为了简化代码，我使用了 type 关键字创建了一个<strong>BoxedError 类型，是 Box<dyn Error + Send + Sync + 'static> 的别名，它是 Error trait 的 trait object</strong>，除了要求类型实现了 Error trait 外，它还有额外的约束：类型必须满足 Send/Sync 这两个 trait。</p>

<p>在参数中使用 trait object 比较简单，再来看一个实战中的<a href="https://docs.rs/reqwest/0.11.5/reqwest/cookie/trait.CookieStore.html" target="_blank">例子</a>巩固一下：</p>

<pre><code>pub trait CookieStore: Send + Sync {
    fn set_cookies(
        &amp;self, 
        cookie_headers: &amp;mut dyn Iterator&lt;Item = &amp;HeaderValue&gt;, 
        url: &amp;Url
    );
    fn cookies(&amp;self, url: &amp;Url) -&gt; Option&lt;HeaderValue&gt;;
}
</code></pre>

<p>这是我们之前使用过的 reqwest 库中的一个处理 CookieStore 的 trait。在 set_cookies 方法中使用了 &amp;mut dyn Iterator 这样一个 trait object。</p>

<h3 id="在函数返回值中使用">在函数返回值中使用</h3>

<p>好，相信你对在参数中如何使用 trait object 已经没有什么问题了，我们再看返回值中使用 trait object，这是 trait object 使用频率比较高的场景。</p>

<p>之前已经出现过很多次了。比如上一讲已经详细介绍的，为何 KV server 里的 Storage trait 不能使用泛型参数来处理返回的 iterator，只能用 Box<dyn Iterator>：</p>

<pre><code>pub trait Storage: Send + Sync + 'static {
    ...
    /// 遍历 HashTable，返回 kv pair 的 Iterator
    fn get_iter(&amp;self, table: &amp;str) -&gt; Result&lt;Box&lt;dyn Iterator&lt;Item = Kvpair&gt;&gt;, KvError&gt;;
}
</code></pre>

<p>再来看一些实战中会遇到的例子。</p>

<p>首先是 <a href="https://docs.rs/async-trait" target="_blank">async_trait</a>。它是一种特殊的 trait，方法中包含 async fn。目前 <a href="https://smallcultfollowing.com/babysteps/blog/2019/10/26/async-fn-in-traits-are-hard/" target="_blank">Rust 并不支持 trait 中使用 async fn</a>，一个变通的方法是使用 async_trait 宏。</p>

<p>在 get hands dirty 系列中，我们就使用过 async trait。下面是[第 6 讲]SQL查询工具数据源的获取中定义的 Fetch trait：</p>

<pre><code>// Rust 的 async trait 还没有稳定，可以用 async_trait 宏
#[async_trait]
pub trait Fetch {
    type Error;
    async fn fetch(&amp;self) -&gt; Result&lt;String, Self::Error&gt;;
}
</code></pre>

<p>这里宏展开后，类似于：</p>

<pre><code>pub trait Fetch {
    type Error;
    fn fetch&lt;'a&gt;(&amp;'a self) -&gt; 
        Result&lt;Pin&lt;Box&lt;dyn Future&lt;Output = String&gt; + Send + 'a&gt;&gt;, Self::Error&gt;;
}
</code></pre>

<p>它使用了 trait object 作为返回值。这样，不管 fetch() 的实现，返回什么样的 Future 类型，都可以被 trait object 统一起来，调用者只需要按照正常 Future 的接口使用即可。</p>

<p>我们再看一个 <a href="https://github.com/mcginty/snow" target="_blank">snow</a> 下的 <a href="https://docs.rs/snow/0.8.0/snow/resolvers/trait.CryptoResolver.html" target="_blank">CryptoResolver</a> 的例子：</p>

<pre><code>/// An object that resolves the providers of Noise crypto choices
pub trait CryptoResolver {
    /// Provide an implementation of the Random trait or None if none available.
    fn resolve_rng(&amp;self) -&gt; Option&lt;Box&lt;dyn Random&gt;&gt;;

    /// Provide an implementation of the Dh trait for the given DHChoice or None if unavailable.
    fn resolve_dh(&amp;self, choice: &amp;DHChoice) -&gt; Option&lt;Box&lt;dyn Dh&gt;&gt;;

    /// Provide an implementation of the Hash trait for the given HashChoice or None if unavailable.
    fn resolve_hash(&amp;self, choice: &amp;HashChoice) -&gt; Option&lt;Box&lt;dyn Hash&gt;&gt;;

    /// Provide an implementation of the Cipher trait for the given CipherChoice or None if unavailable.
    fn resolve_cipher(&amp;self, choice: &amp;CipherChoice) -&gt; Option&lt;Box&lt;dyn Cipher&gt;&gt;;

    /// Provide an implementation of the Kem trait for the given KemChoice or None if unavailable
    #[cfg(feature = &quot;hfs&quot;)]
    fn resolve_kem(&amp;self, _choice: &amp;KemChoice) -&gt; Option&lt;Box&lt;dyn Kem&gt;&gt; {
        None
    }
}
</code></pre>

<p>这是一个处理 <a href="https://zhuanlan.zhihu.com/p/96944134" target="_blank">Noise Protocol</a> 使用何种加密算法的一个 trait。这个 trait 的每个方法，都返回一个 trait object，每个 trait object 都提供加密算法中所需要的不同的能力，比如随机数生成算法（Random）、DH 算法（Dh）、哈希算法（Hash）、对称加密算法（Cipher）和密钥封装算法（Kem）。</p>

<p>所有这些，都有一系列的具体的算法实现，通过 CryptoResolver trait，可以得到当前使用的某个具体算法的 trait object，<strong>这样，在处理业务逻辑时，我们不用关心当前究竟使用了什么算法，就能根据这些 trait object 构筑相应的实现</strong>，比如下面的 generate_keypair：</p>

<pre><code>pub fn generate_keypair(&amp;self) -&gt; Result&lt;Keypair, Error&gt; {
    // 拿到当前的随机数生成算法
    let mut rng = self.resolver.resolve_rng().ok_or(InitStage::GetRngImpl)?;
		// 拿到当前的 DH 算法
    let mut dh = self.resolver.resolve_dh(&amp;self.params.dh).ok_or(InitStage::GetDhImpl)?;
    let mut private = vec![0u8; dh.priv_len()];
    let mut public = vec![0u8; dh.pub_len()];
    // 使用随机数生成器 和 DH 生成密钥对
    dh.generate(&amp;mut *rng);

    private.copy_from_slice(dh.privkey());
    public.copy_from_slice(dh.pubkey());

    Ok(Keypair { private, public })
}
</code></pre>

<p>说句题外话，如果你想更好地学习 trait 和 trait object 的使用，snow 是一个很好的学习资料。你可以顺着 CryptoResolver 梳理它用到的这几个主要的加密算法相关的 trait，看看别人是怎么定义 trait、如何把各个 trait 关联起来，以及最终如何把 trait 和核心数据结构联系起来的（小提示：<a href="https://docs.rs/snow/0.8.0/snow/struct.Builder.html" target="_blank">Builder</a> 以及 <a href="https://docs.rs/snow/0.8.0/snow/struct.HandshakeState.html" target="_blank">HandshakeState</a>）。</p>

<h2 id="在数据结构中使用-trait-object">在数据结构中使用 trait object</h2>

<p>了解了在函数中是如何使用 trait object 的，接下来我们再看看在数据结构中，如何使用 trait object。</p>

<p>继续以 snow 的代码为例，看 HandshakeState这个用于处理 Noise Protocol 握手协议的数据结构，用到了哪些 trait object（<a href="https://docs.rs/snow/0.8.0/src/snow/handshakestate.rs.html#29-48" target="_blank">代码</a>）：</p>

<pre><code>pub struct HandshakeState {
    pub(crate) rng:              Box&lt;dyn Random&gt;,
    pub(crate) symmetricstate:   SymmetricState,
    pub(crate) cipherstates:     CipherStates,
    pub(crate) s:                Toggle&lt;Box&lt;dyn Dh&gt;&gt;,
    pub(crate) e:                Toggle&lt;Box&lt;dyn Dh&gt;&gt;,
    pub(crate) fixed_ephemeral:  bool,
    pub(crate) rs:               Toggle&lt;[u8; MAXDHLEN]&gt;,
    pub(crate) re:               Toggle&lt;[u8; MAXDHLEN]&gt;,
    pub(crate) initiator:        bool,
    pub(crate) params:           NoiseParams,
    pub(crate) psks:             [Option&lt;[u8; PSKLEN]&gt;; 10],
    #[cfg(feature = &quot;hfs&quot;)]
    pub(crate) kem:              Option&lt;Box&lt;dyn Kem&gt;&gt;,
    #[cfg(feature = &quot;hfs&quot;)]
    pub(crate) kem_re:           Option&lt;[u8; MAXKEMPUBLEN]&gt;,
    pub(crate) my_turn:          bool,
    pub(crate) message_patterns: MessagePatterns,
    pub(crate) pattern_position: usize,
}
</code></pre>

<p>你不需要了解 Noise protocol，也能够大概可以明白这里 Random、Dh 以及 Kem 三个 trait object 的作用：它们为握手期间使用的加密协议提供最大的灵活性。</p>

<p>想想看，如果这里不用 trait object，这个数据结构该怎么处理？</p>

<p>可以用泛型参数，也就是说：</p>

<pre><code>pub struct HandshakeState&lt;R, D, K&gt;
where
    R: Random,
    D: Dh,
    K: Kem
{
  ...
}
</code></pre>

<p>这是我们大部分时候处理这样的数据结构的选择。但是，过多的泛型参数会带来两个问题：首先，代码实现过程中，所有涉及的接口都变得非常臃肿，你在使用 HandshakeState<R, D, K> 的任何地方，都必须带着这几个泛型参数以及它们的约束。其次，这些参数所有被使用到的情况，组合起来，会生成大量的代码。</p>

<p>而使用 trait object，我们在牺牲一点性能的前提下，消除了这些泛型参数，实现的代码更干净清爽，且代码只会有一份实现。</p>

<p>在数据结构中使用 trait object 还有一种很典型的场景是，<strong>闭包</strong>。</p>

<p>因为在 Rust 中，闭包都是以匿名类型的方式出现，我们无法直接在数据结构中使用其类型，只能用泛型参数。而对闭包使用泛型参数后，如果捕获的数据太大，可能造成数据结构本身太大；但有时，我们并不在意一点点性能损失，更愿意让代码处理起来更方便。</p>

<p>比如用于做 RBAC 的库 <a href="https://github.com/osohq/oso" target="_blank">oso</a> 里的 AttributeGetter，它包含了一个 Fn：</p>

<pre><code>#[derive(Clone)]
pub struct AttributeGetter(
    Arc&lt;dyn Fn(&amp;Instance, &amp;mut Host) -&gt; crate::Result&lt;PolarValue&gt; + Send + Sync&gt;,
);
</code></pre>

<p>如果你对在 Rust 中如何实现 Python 的 getattr 感兴趣，可以看看 <a href="https://github.com/osohq/oso" target="_blank">oso</a> 的代码。</p>

<p>再比如做交互式 CLI 的 <a href="https://github.com/mitsuhiko/dialoguer" target="_blank">dialoguer</a> 的 <a href="https://docs.rs/dialoguer/0.8.0/dialoguer/struct.Input.html" target="_blank">Input</a>，它的 validator 就是一个 FnMut：</p>

<pre><code>pub struct Input&lt;'a, T&gt; {
    prompt: String,
    default: Option&lt;T&gt;,
    show_default: bool,
    initial_text: Option&lt;String&gt;,
    theme: &amp;'a dyn Theme,
    permit_empty: bool,
    validator: Option&lt;Box&lt;dyn FnMut(&amp;T) -&gt; Option&lt;String&gt; + 'a&gt;&gt;,
    #[cfg(feature = &quot;history&quot;)]
    history: Option&lt;&amp;'a mut dyn History&lt;T&gt;&gt;,
}
</code></pre>

<h2 id="用-trait-object-处理-kv-server-的-service-结构">用 trait object 处理 KV server 的 Service 结构</h2>

<p>好，到这里用 trait object 做动态分发的几个场景我们就介绍完啦，来写段代码练习一下。</p>

<p>就用之前写的 KV server 的 Service 结构来趁热打铁，我们尝试对它做个处理，使其内部使用 trait object。</p>

<p>其实对于 KV server 而言，使用泛型是更好的选择，因为此处泛型并不会造成太多的复杂性，我们也不希望丢掉哪怕一点点性能。然而，出于学习的目的，我们可以看看如果 store 使用 trait object，代码会变成什么样子。你自己可以先尝试一下，再来看下面的示例（<a href="https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=bf13294ace57216e529ac1f93efd9291" target="_blank">代码</a>）：</p>

<pre><code>use std::{error::Error, sync::Arc};

// 定义类型，让 KV server 里的 trait 可以被编译通过
pub type KvError = Box&lt;dyn Error + Send + Sync&gt;;
pub struct Value(i32);
pub struct Kvpair(i32, i32);

/// 对存储的抽象，我们不关心数据存在哪儿，但需要定义外界如何和存储打交道
pub trait Storage: Send + Sync + 'static {
    fn get(&amp;self, table: &amp;str, key: &amp;str) -&gt; Result&lt;Option&lt;Value&gt;, KvError&gt;;
    fn set(&amp;self, table: &amp;str, key: String, value: Value) -&gt; Result&lt;Option&lt;Value&gt;, KvError&gt;;
    fn contains(&amp;self, table: &amp;str, key: &amp;str) -&gt; Result&lt;bool, KvError&gt;;
    fn del(&amp;self, table: &amp;str, key: &amp;str) -&gt; Result&lt;Option&lt;Value&gt;, KvError&gt;;
    fn get_all(&amp;self, table: &amp;str) -&gt; Result&lt;Vec&lt;Kvpair&gt;, KvError&gt;;
    fn get_iter(&amp;self, table: &amp;str) -&gt; Result&lt;Box&lt;dyn Iterator&lt;Item = Kvpair&gt;&gt;, KvError&gt;;
}

// 使用 trait object，不需要泛型参数，也不需要 ServiceInner 了
pub struct Service {
    pub store: Arc&lt;dyn Storage&gt;,
}

// impl 的代码略微简单一些
impl Service {
    pub fn new&lt;S: Storage&gt;(store: S) -&gt; Self {
        Self {
            store: Arc::new(store),
        }
    }
}

// 实现 trait 时也不需要带着泛型参数
impl Clone for Service {
    fn clone(&amp;self) -&gt; Self {
        Self {
            store: Arc::clone(&amp;self.store),
        }
    }
}
</code></pre>

<p>从这段代码中可以看到，通过牺牲一点性能，我们让代码整体撰写和使用起来方便了不少。</p>

<h2 id="小结">小结</h2>

<p>无论是上一讲的泛型参数，还是今天的 trait object，都是 Rust 处理多态的手段。当系统需要使用多态来解决复杂多变的需求，让同一个接口可以展现不同的行为时，我们要决定究竟是编译时的静态分发更好，还是运行时的动态分发更好。</p>

<p>一般情况下，作为 Rust 开发者，我们不介意泛型参数带来的稍微复杂的代码结构，愿意用开发时的额外付出，换取运行时的高效；但<strong>有时候，当泛型参数过多，导致代码出现了可读性问题，或者运行效率并不是主要矛盾的时候，我们可以通过使用 trait object 做动态分发，来降低代码的复杂度</strong>。</p>

<p>具体看，在有些情况，我们不太容易使用泛型参数，比如希望函数返回某个 trait 的实现，或者数据结构中某些参数在运行时的组合过于复杂，比如上文提到的 HandshakeState，此时，使用 trait object 是更好的选择。</p>

<h3 id="思考题">思考题</h3>

<p>期中测试中我给出的 <a href="https://github.com/tyrchen/geektime-rust" target="_blank">rgrep 的代码</a>，如果把 StrategyFn 的接口改成使用 trait object：</p>

<pre><code>/// 定义类型，这样，在使用时可以简化复杂类型的书写
pub type StrategyFn = fn(&amp;Path, &amp;mut dyn BufRead, &amp;Regex, &amp;mut dyn Write) -&gt; Result&lt;(), GrepError&gt;;
</code></pre>

<p>你能把实现部分修改，使测试通过么？对比修改前后的代码，你觉得对 rgrep，哪种实现更好？为什么？</p>

<p>今天你完成了Rust学习的第24次打卡。如果你觉得有收获，也欢迎分享给你身边的朋友，邀他一起讨论。我们下节课见。</p>

<h2 id="延伸阅读">延伸阅读</h2>

<p>我们总说 trait object 性能会差一些，因为需要从 vtable 中额外加载对应的方法的地址，才能跳转执行。那么这个性能差异究竟有多大呢？网上有人说调用 trait object 的方法，性能会比直接调用类型的方法差一个数量级，真的么？</p>

<p>我用 criterion 做了一个简单的测试，测试的 trait 使用的就是我们这一讲使用的 Executor trait。测试代码如下（你可以访问 <a href="https://github.com/tyrchen/geektime-rust" target="_blank">GitHub repo</a> 中这一讲的代码）：</p>

<pre><code>use advanced_trait_objects::{
    execute_boxed_trait_object, execute_generics, execute_trait_object, Shell,
};
use criterion::{black_box, criterion_group, criterion_main, Criterion};

pub fn generics_benchmark(c: &amp;mut Criterion) {
    c.bench_function(&quot;generics&quot;, |b| {
        b.iter(|| {
            let cmd = Shell::new(&quot;ls&quot;, &amp;[]);
            execute_generics(black_box(&amp;cmd)).unwrap();
        })
    });
}

pub fn trait_object_benchmark(c: &amp;mut Criterion) {
    c.bench_function(&quot;trait object&quot;, |b| {
        b.iter(|| {
            let cmd = Shell::new(&quot;ls&quot;, &amp;[]);
            execute_trait_object(black_box(&amp;cmd)).unwrap();
        })
    });
}

pub fn boxed_object_benchmark(c: &amp;mut Criterion) {
    c.bench_function(&quot;boxed object&quot;, |b| {
        b.iter(|| {
            let cmd = Box::new(Shell::new(&quot;ls&quot;, &amp;[]));
            execute_boxed_trait_object(black_box(cmd)).unwrap();
        })
    });
}

criterion_group!(
    benches,
    generics_benchmark,
    trait_object_benchmark,
    boxed_object_benchmark
);
criterion_main!(benches);
</code></pre>

<p>为了不让实现本身干扰接口调用的速度，我们在 trait 的方法中什么也不做，直接返回：</p>

<pre><code>impl&lt;'a, 'b&gt; Executor for Shell&lt;'a, 'b&gt; {
    fn run(&amp;self) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
        // let output = Command::new(self.cmd).args(self.args).output()?;
        // Ok(output.status.code())
        Ok(Some(0))
    }
}
</code></pre>

<p>测试结果如下：</p>

<pre><code>generics                time:   [3.0995 ns 3.1549 ns 3.2172 ns]                      
                        change: [-96.890% -96.810% -96.732%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
Found 5 outliers among 100 measurements (5.00%)
  4 (4.00%) high mild
  1 (1.00%) high severe

trait object            time:   [4.0348 ns 4.0934 ns 4.1552 ns]                          
                        change: [-96.024% -95.893% -95.753%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
Found 8 outliers among 100 measurements (8.00%)
  3 (3.00%) high mild
  5 (5.00%) high severe

boxed object            time:   [65.240 ns 66.473 ns 67.777 ns]                         
                        change: [-67.403% -66.462% -65.530%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
Found 2 outliers among 100 measurements (2.00%)
</code></pre>

<p>可以看到，使用泛型做静态分发最快，平均 3.15ns；使用 &amp;dyn Executor 平均速度 4.09ns，要慢 30%；而使用 Box<dyn Executor> 平均速度 66.47ns，慢了足足 20 倍。可见，额外的内存访问并不是 trait object 的效率杀手，有些场景下为了使用 trait object 不得不做的额外的堆内存分配，才是主要的效率杀手。</p>

<p>那么，这个性能差异重要么？</p>

<p>在回答这个问题之前，我们把 run() 方法改回来：</p>

<pre><code>impl&lt;'a, 'b&gt; Executor for Shell&lt;'a, 'b&gt; {
    fn run(&amp;self) -&gt; Result&lt;Option&lt;i32&gt;, BoxedError&gt; {
        let output = Command::new(self.cmd).args(self.args).output()?;
        Ok(output.status.code())
    }
}
</code></pre>

<p>我们知道 Command 的执行速度比较慢，但是想再看看，对于执行效率低的方法，这个性能差异是否重要。</p>

<p>新的测试结果不出所料：</p>

<pre><code>generics                time:   [4.6901 ms 4.7267 ms 4.7678 ms]                      
                        change: [+145694872% +148496855% +151187366%] (p = 0.00 &lt; 0.05)
                        Performance has regressed.
Found 7 outliers among 100 measurements (7.00%)
  3 (3.00%) high mild
  4 (4.00%) high severe

trait object            time:   [4.7452 ms 4.7912 ms 4.8438 ms]                          
                        change: [+109643581% +113478268% +116908330%] (p = 0.00 &lt; 0.05)
                        Performance has regressed.
Found 7 outliers among 100 measurements (7.00%)
  4 (4.00%) high mild
  3 (3.00%) high severe

boxed object            time:   [4.7867 ms 4.8336 ms 4.8874 ms]                          
                        change: [+6935303% +7085465% +7238819%] (p = 0.00 &lt; 0.05)
                        Performance has regressed.
Found 8 outliers among 100 measurements (8.00%)
  4 (4.00%) high mild
  4 (4.00%) high severe
</code></pre>

<p>因为执行一个 Shell 命令的效率实在太低，到毫秒的量级，虽然 generics 依然最快，但使用 &amp;dyn Executor 和 Box<dyn executor> 也不过只比它慢了 1% 和 2%。</p>

<p>所以，如果是那种执行效率在数百纳秒以内的函数，是否使用 trait object，尤其是 boxed trait object，性能差别会比较明显；但当函数本身的执行需要数微秒到数百微秒时，性能差别就很小了；到了毫秒的量级，性能的差别几乎无关紧要。</p>

<p>总的来说，大部分情况，我们在撰写代码的时候，不必太在意 trait object 的性能问题。如果你实在在意关键路径上 trait object 的性能，那么先尝试看能不能不要做额外的堆内存分配。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c3afafaffaf7f2f2f3f483a4aea2aaafeda0acae" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a4fa49bb57755',t:'MTczNDEzODU5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=42&#32;阶段实操（7）：构建一个简单的KV&#32;server-如何做大的重构？>
        <link rel="icon" href="/static/favicon.png">
        <title>42 阶段实操（7）：构建一个简单的KV server-如何做大的重构？ </title>
        
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
                            <h1 id="title" data-id="42 阶段实操（7）：构建一个简单的KV server-如何做大的重构？" class="title">42 阶段实操（7）：构建一个简单的KV server-如何做大的重构？</h1>
                            <div><p>你好，我是陈天。</p>

<p>在软件开发的过程中，一开始设计得再精良，也扛不住无缘无故的需求变更。所以我们要妥善做架构设计，让它能满足潜在的需求；但也不能过度设计，让它去适应一些虚无缥缈的需求。好的开发者，要能够把握这个度。</p>

<p>到目前为止，我们的 KV server 已经羽翼丰满，作为一个基本的 KV 存储够用了。</p>

<p>这时候，产品经理突然抽风，想让你在这个 Server 上加上类似 Redis 的 Pub/Sub 支持。你说：别闹，这根本就是两个产品。产品经理回应： Redis 也支持 Pub/Sub。你怼回去：那干脆用 Redis 的 Pub/Sub 得了。产品经理听了哈哈一笑：行，用 Redis 挺好，我们还能把你的工钱省下来呢。天都聊到这份上了，你只好妥协：那啥，姐，我做，我做还不行么？</p>

<p>这虽是个虚构的故事，但类似的大需求变更在我们开发者的日常工作中相当常见。我们就以这个具备不小难度的挑战，来看看，如何对一个已经成形的系统进行大的重构。</p>

<h2 id="现有架构分析">现有架构分析</h2>

<p>先简单回顾一下 Redis 对 Pub/Sub 的支持：客户端可以随时发起 SUBSCRIBE、PUBLISH 和 UNSUBSCRIBE。如果客户端 A 和 B SUBSCRIBE 了一个叫 lobby 的主题，客户端 C 往 lobby 里发了 “hello”，A 和 B 都将立即收到这个信息。</p>

<p>使用起来是这个样子的：</p>

<pre><code>A: SUBSCRIBE &quot;lobby&quot;
A: SUBSCRIBE &quot;王者荣耀&quot;
B: SUBSCRIBE &quot;lobby&quot;
C: PUBLISH &quot;lobby&quot; &quot;hello&quot;
// A/B 都收到 &quot;hello&quot;
B: UNSUBSCRIBE &quot;lobby&quot;
B: SUBSCRIBE &quot;王者荣耀&quot;
D: PUBLISH &quot;lobby&quot; &quot;goodbye&quot;
// 只有 A 收到 &quot;goodbye&quot;
C: PUBLISH &quot;王者荣耀&quot; &quot;good game&quot;
// A/B 都收到 &quot;good game&quot;
</code></pre>

<p>带着这个需求，我们重新审视目前的架构：</p>

<p><img src="assets/82da823b4eb16935fdeyy727e3b3262c.jpg" alt="图片" /></p>

<p>要支持 Pub/Sub，现有架构有两个很大的问题。</p>

<p><strong>首先，CommandService 是一个同步的处理</strong>，来一个命令，立刻就能计算出一个值返回。但现在来一个 SUBSCRIBE 命令，它期待的不是一个值，而是未来可能产生的若干个值。我们讲过 Stream 代表未来可能产生的一系列值，所以这里需要返回一个异步的 Stream。</p>

<p>因此，我们要么需要牺牲 CommandService 这个 trait 来适应新的需求，要么构建一个新的、和 CommandService trait 并列的 trait，来处理和 Pub/Sub 有关的命令。</p>

<p>其次，<strong>如果直接在 TCP/TLS 之上构建 Pub/Sub 的支持，我们需要在 Request 和 Response 之间建立“流”的概念</strong>，为什么呢？</p>

<p>之前我们的协议运行模式是同步的，一来一回：</p>

<p><img src="assets/7byy9cdb2c3651e4cd77bdda89a52968.jpg" alt="图片" /></p>

<p>但是，如果继续采用这样的方式，就会有应用层的 <a href="https://en.wikipedia.org/wiki/Head-of-line_blocking" target="_blank">head of line blocking</a>（队头阻塞）问题，一个 SUBSCRIBE 命令，因为其返回结果不知道什么时候才结束，会阻塞后续的所有命令。所以，我们需要在一个连接里，划分出很多彼此独立的“流”，让它们的收发不受影响：</p>

<p><img src="assets/67659457626d12eba6e26b37ayy08edb.jpg" alt="图片" /></p>

<p>这种流式处理的典型协议是使用了多路复用（multiplex）的 HTTP/2。所以，一种方案是我们可以把 KV server 构建在使用 HTTP/2 的 gRPC 上。不过，HTTP 是个太过庞杂的协议，对于 KV server 这种性能非常重要的服务来说，不必要的额外开销太多，所以它不太适合。</p>

<p>另一种方式是使用 <a href="https://github.com/hashicorp/yamux/blob/master/spec.md" target="_blank">Yamux</a> 协议，之前介绍过，它是一个简单的、和 HTTP/2 内部多路复用机制非常类似的协议。如果使用它，整个协议的交互看上去是这个样子的：</p>

<p><img src="assets/31f3efcd510ff6a3yy0caf32dbfd8667.jpg" alt="图片" /></p>

<p>Yamux 适合不希望引入 HTTP 的繁文缛节（大量的头信息），在 TCP 层做多路复用的场景，今天就用它来支持我们所要实现的 Pub/Sub。</p>

<h2 id="使用-yamux-做多路复用">使用 yamux 做多路复用</h2>

<p>Rust 下有 <a href="https://github.com/libp2p/rust-yamux" target="_blank">rust-yamux</a> 这个库，来支持 yamux。除此之外，我们还需要 tokio-util，它提供了 tokio 下的 trait 和 futures 下的 trait 的兼容能力。在 Cargo.toml 中引入它们：</p>

<pre><code>[dependencies]
...
tokio-util = { version = &quot;0.6&quot;, features = [&quot;compat&quot;]} # tokio 和 futures 的兼容性库
...
yamux = &quot;0.9&quot; # yamux 多路复用支持
...
</code></pre>

<p>然后创建 src/network/multiplex.rs（记得在 <a href="http://mod.rs" target="_blank">mod.rs</a> 里引用），添入如下代码：</p>

<pre><code>use futures::{future, Future, TryStreamExt};
use std::marker::PhantomData;
use tokio::io::{AsyncRead, AsyncWrite};
use tokio_util::compat::{Compat, FuturesAsyncReadCompatExt, TokioAsyncReadCompatExt};
use yamux::{Config, Connection, ConnectionError, Control, Mode, WindowUpdateMode};

/// Yamux 控制结构
pub struct YamuxCtrl&lt;S&gt; {
    /// yamux control，用于创建新的 stream
    ctrl: Control,
    _conn: PhantomData&lt;S&gt;,
}

impl&lt;S&gt; YamuxCtrl&lt;S&gt;
where
    S: AsyncRead + AsyncWrite + Unpin + Send + 'static,
{
    /// 创建 yamux 客户端
    pub fn new_client(stream: S, config: Option&lt;Config&gt;) -&gt; Self {
        Self::new(stream, config, true, |_stream| future::ready(Ok(())))
    }

    /// 创建 yamux 服务端，服务端我们需要具体处理 stream
    pub fn new_server&lt;F, Fut&gt;(stream: S, config: Option&lt;Config&gt;, f: F) -&gt; Self
    where
        F: FnMut(yamux::Stream) -&gt; Fut,
        F: Send + 'static,
        Fut: Future&lt;Output = Result&lt;(), ConnectionError&gt;&gt; + Send + 'static,
    {
        Self::new(stream, config, false, f)
    }

    // 创建 YamuxCtrl
    fn new&lt;F, Fut&gt;(stream: S, config: Option&lt;Config&gt;, is_client: bool, f: F) -&gt; Self
    where
        F: FnMut(yamux::Stream) -&gt; Fut,
        F: Send + 'static,
        Fut: Future&lt;Output = Result&lt;(), ConnectionError&gt;&gt; + Send + 'static,
    {
        let mode = if is_client {
            Mode::Client
        } else {
            Mode::Server
        };

        // 创建 config
        let mut config = config.unwrap_or_default();
        config.set_window_update_mode(WindowUpdateMode::OnRead);

        // 创建 config，yamux::Stream 使用的是 futures 的 trait 所以需要 compat() 到 tokio 的 trait
        let conn = Connection::new(stream.compat(), config, mode);

        // 创建 yamux ctrl
        let ctrl = conn.control();

        // pull 所有 stream 下的数据
        tokio::spawn(yamux::into_stream(conn).try_for_each_concurrent(None, f));

        Self {
            ctrl,
            _conn: PhantomData::default(),
        }
    }

    /// 打开一个新的 stream
    pub async fn open_stream(&amp;mut self) -&gt; Result&lt;Compat&lt;yamux::Stream&gt;, ConnectionError&gt; {
        let stream = self.ctrl.open_stream().await?;
        Ok(stream.compat())
    }
}
</code></pre>

<p>这段代码提供了 Yamux 的基本处理。如果有些地方你看不明白，比如 WindowUpdateMode，yamux::into_stream() 等，很正常，需要看看 <a href="https://github.com/libp2p/rust-yamux" target="_blank">yamux crate</a> 的文档和例子。</p>

<p>这里有一个复杂的接口，我们稍微解释一下：</p>

<pre><code>pub fn new_server&lt;F, Fut&gt;(stream: S, config: Option&lt;Config&gt;, f: F) -&gt; Self
where
    F: FnMut(yamux::Stream) -&gt; Fut,
    F: Send + 'static,
    Fut: Future&lt;Output = Result&lt;(), ConnectionError&gt;&gt; + Send + 'static,
{
    Self::new(stream, config, false, f)
}
</code></pre>

<p>它的意思是，参数 f 是一个 FnMut 闭包，接受一个 yamux::Stream 参数，返回 Future。这样的结构我们之前见过，之所以接口这么复杂，是因为 Rust 还没有把 async 闭包稳定下来。所以，如果要想写一个 <code>async || {}</code>，这是最佳的方式。</p>

<p>还是写一段测试测一下（篇幅关系，完整的代码就不放了，你可以到 GitHub repo 下对照 diff_yamux 看修改）：</p>

<pre><code>#[tokio::test]
async fn yamux_ctrl_client_server_should_work() -&gt; Result&lt;()&gt; {
    // 创建使用了 TLS 的 yamux server
    let acceptor = tls_acceptor(false)?;
    let addr = start_yamux_server(&quot;127.0.0.1:0&quot;, acceptor, MemTable::new()).await?;

    let connector = tls_connector(false)?;
    let stream = TcpStream::connect(addr).await?;
    let stream = connector.connect(stream).await?;
    // 创建使用了 TLS 的 yamux client
    let mut ctrl = YamuxCtrl::new_client(stream, None);

    // 从 client ctrl 中打开一个新的 yamux stream
    let stream = ctrl.open_stream().await?;
    // 封装成 ProstClientStream
    let mut client = ProstClientStream::new(stream);

    let cmd = CommandRequest::new_hset(&quot;t1&quot;, &quot;k1&quot;, &quot;v1&quot;.into());
    client.execute(cmd).await.unwrap();

    let cmd = CommandRequest::new_hget(&quot;t1&quot;, &quot;k1&quot;);
    let res = client.execute(cmd).await.unwrap();
    assert_res_ok(res, &amp;[&quot;v1&quot;.into()], &amp;[]);

    Ok(())
}
</code></pre>

<p>可以看到，经过简单的封装，yamux 就很自然地融入到我们现有的架构中。因为 open_stream() 得到的是符合 tokio AsyncRead/AsyncWrite 的 stream，所以它可以直接配合 ProstClientStream 使用。也就是说，我们网络层又改动了一下，但后面逻辑依然不用变。</p>

<p>运行 <code>cargo test</code> ，所有测试都能通过。</p>

<h2 id="支持-pub-sub">支持 pub/sub</h2>

<p>好，现在网络层已经支持了 yamux，为多路复用打下了基础。我们来看 pub/sub 具体怎么实现。</p>

<p>首先修改 abi.proto，加入新的几个命令：</p>

<pre><code>// 来自客户端的命令请求
message CommandRequest {
  oneof request_data {
    ...
    Subscribe subscribe = 10;
    Unsubscribe unsubscribe = 11;
    Publish publish = 12;
  }
}

// subscribe 到某个主题，任何发布到这个主题的数据都会被收到
// 成功后，第一个返回的 CommandResponse，我们返回一个唯一的 subscription id
message Subscribe { string topic = 1; }

// 取消对某个主题的订阅
message Unsubscribe {
  string topic = 1;
  uint32 id = 2;
}

// 发布数据到某个主题
message Publish {
  string topic = 1;
  repeated Value data = 2;
}
</code></pre>

<p>命令的响应我们不用改变。当客户端 Subscribe 时，返回的 stream 里的第一个值包含订阅 ID，这是一个全局唯一的 ID，这样，客户端后续可以用 Unsubscribe 取消。</p>

<h3 id="pub-sub-如何设计">Pub/Sub 如何设计？</h3>

<p>那么，Pub/Sub 该如何实现呢？</p>

<p>我们可以用<strong>两张表</strong>：一张 Topic Table，存放主题和对应的订阅列表；一张 Subscription Table，存放订阅 ID 和 channel 的发送端。</p>

<p>当 SUBSCRIBE 时，我们获取一个订阅 ID，插入到 Topic Table，然后再创建一个 MPSC channel，把 channel 的发送端和订阅 ID 存入 subscription table。</p>

<p>这样，当有人 PUBLISH 时，可以从 Topic table 中找到对应的订阅 ID 的列表，然后循环从 subscription table 中找到对应的 Sender，往里面写入数据。此时，channel 的 Receiver 端会得到数据，这个数据会被 yamux stream poll 到，然后发给客户端。</p>

<p>整个流程如下图所示：</p>

<p><img src="assets/7ce3046af823dbbdaa7b47d12d04ce30.jpg" alt="图片" /></p>

<p>有了这个基本设计，我们可以着手接口和数据结构的构建了：</p>

<pre><code>/// 下一个 subscription id
static NEXT_ID: AtomicU32 = AtomicU32::new(1);

/// 获取下一个 subscription id
fn get_next_subscription_id() -&gt; u32 {
    NEXT_ID.fetch_add(1, Ordering::Relaxed)
}

pub trait Topic: Send + Sync + 'static {
    /// 订阅某个主题
    fn subscribe(self, name: String) -&gt; mpsc::Receiver&lt;Arc&lt;CommandResponse&gt;&gt;;
    /// 取消对主题的订阅
    fn unsubscribe(self, name: String, id: u32);
    /// 往主题里发布一个数据
    fn publish(self, name: String, value: Arc&lt;CommandResponse&gt;);
}

/// 用于主题发布和订阅的数据结构
#[derive(Default)]
pub struct Broadcaster {
    /// 所有的主题列表
    topics: DashMap&lt;String, DashSet&lt;u32&gt;&gt;,
    /// 所有的订阅列表
    subscriptions: DashMap&lt;u32, mpsc::Sender&lt;Arc&lt;CommandResponse&gt;&gt;&gt;,
}
</code></pre>

<p>这里，subscription_id 我们用一个 AtomicU32 来表述。</p>

<p>对于这样一个全局唯一的 ID，很多同学喜欢用 UUID4 来表述。注意使用 UUID 的话，存储时一定不要存它的字符串表现形式，太浪费内存且每次都有额外的堆分配，应该用它 u128 的表现形式。</p>

<p>不过即便 u128，也比 u32 浪费很多空间。假设某个主题 M 下有一万个订阅，要往这个 M 里发送一条消息，就意味着整个 DashSet<u32> 的一次拷贝，乘上一万，u32 的话做 40k 内存的拷贝，而 u128 要做 160k 内存的拷贝。这个性能上的差距就很明显了。</p>

<p>另外，我们把 CommandResponse 封装进了一个 Arc。如果一条消息要发送给一万个客户端，那么我们不希望这条消息被复制后，再被发送，而是直接发送同一份数据。</p>

<p>这里对 Pub/Sub 的接口，构建了一个 Topic trait。虽然目前我们只有 Broadcaster 会实现 Topic trait，但未来也许会换不同的实现方式，所以，抽象出 Topic trait 很有意义。</p>

<h3 id="pub-sub-的实现">Pub/Sub 的实现</h3>

<p>好，我们来写代码。创建 src/service/topic.rs（记得在 <a href="http://mod.rs" target="_blank">mod.rs</a> 里引用），并添入：</p>

<pre><code>use dashmap::{DashMap, DashSet};
use std::sync::{
    atomic::{AtomicU32, Ordering},
    Arc,
};
use tokio::sync::mpsc;
use tracing::{debug, info, warn};

use crate::{CommandResponse, Value};

/// topic 里最大存放的数据
const BROADCAST_CAPACITY: usize = 128;

/// 下一个 subscription id
static NEXT_ID: AtomicU32 = AtomicU32::new(1);

/// 获取下一个 subscription id
fn get_next_subscription_id() -&gt; u32 {
    NEXT_ID.fetch_add(1, Ordering::Relaxed)
}

pub trait Topic: Send + Sync + 'static {
    /// 订阅某个主题
    fn subscribe(self, name: String) -&gt; mpsc::Receiver&lt;Arc&lt;CommandResponse&gt;&gt;;
    /// 取消对主题的订阅
    fn unsubscribe(self, name: String, id: u32);
    /// 往主题里发布一个数据
    fn publish(self, name: String, value: Arc&lt;CommandResponse&gt;);
}

/// 用于主题发布和订阅的数据结构
#[derive(Default)]
pub struct Broadcaster {
    /// 所有的主题列表
    topics: DashMap&lt;String, DashSet&lt;u32&gt;&gt;,
    /// 所有的订阅列表
    subscriptions: DashMap&lt;u32, mpsc::Sender&lt;Arc&lt;CommandResponse&gt;&gt;&gt;,
}

impl Topic for Arc&lt;Broadcaster&gt; {
    fn subscribe(self, name: String) -&gt; mpsc::Receiver&lt;Arc&lt;CommandResponse&gt;&gt; {
        let id = {
            let entry = self.topics.entry(name).or_default();
            let id = get_next_subscription_id();
            entry.value().insert(id);
            id
        };

        // 生成一个 mpsc channel
        let (tx, rx) = mpsc::channel(BROADCAST_CAPACITY);

        let v: Value = (id as i64).into();

        // 立刻发送 subscription id 到 rx
        let tx1 = tx.clone();
        tokio::spawn(async move {
            if let Err(e) = tx1.send(Arc::new(v.into())).await {
                // TODO: 这个很小概率发生，但目前我们没有善后
                warn!(&quot;Failed to send subscription id: {}. Error: {:?}&quot;, id, e);
            }
        });

        // 把 tx 存入 subscription table
        self.subscriptions.insert(id, tx);
        debug!(&quot;Subscription {} is added&quot;, id);

        // 返回 rx 给网络处理的上下文
        rx
    }

    fn unsubscribe(self, name: String, id: u32) {
        if let Some(v) = self.topics.get_mut(&amp;name) {
            // 在 topics 表里找到 topic 的 subscription id，删除
            v.remove(&amp;id);

            // 如果这个 topic 为空，则也删除 topic
            if v.is_empty() {
                info!(&quot;Topic: {:?} is deleted&quot;, &amp;name);
                drop(v);
                self.topics.remove(&amp;name);
            }
        }

        debug!(&quot;Subscription {} is removed!&quot;, id);
        // 在 subscription 表中同样删除
        self.subscriptions.remove(&amp;id);
    }

    fn publish(self, name: String, value: Arc&lt;CommandResponse&gt;) {
        tokio::spawn(async move {
            match self.topics.get(&amp;name) {
                Some(chan) =&gt; {
                    // 复制整个 topic 下所有的 subscription id
                    // 这里我们每个 id 是 u32，如果一个 topic 下有 10k 订阅，复制的成本
                    // 也就是 40k 堆内存（外加一些控制结构），所以效率不算差
                    // 这也是为什么我们用 NEXT_ID 来控制 subscription id 的生成
                    let chan = chan.value().clone();

                    // 循环发送
                    for id in chan.into_iter() {
                        if let Some(tx) = self.subscriptions.get(&amp;id) {
                            if let Err(e) = tx.send(value.clone()).await {
                                warn!(&quot;Publish to {} failed! error: {:?}&quot;, id, e);
                            }
                        }
                    }
                }
                None =&gt; {}
            }
        });
    }
}
</code></pre>

<p>这段代码就是 Pub/Sub 的核心功能了。你可以对照着上面的设计图和代码中的详细注释去理解。我们来写一个测试确保它正常工作：</p>

<pre><code>#[cfg(test)]
mod tests {
    use std::convert::TryInto;

    use crate::assert_res_ok;

    use super::*;

    #[tokio::test]
    async fn pub_sub_should_work() {
        let b = Arc::new(Broadcaster::default());
        let lobby = &quot;lobby&quot;.to_string();

        // subscribe
        let mut stream1 = b.clone().subscribe(lobby.clone());
        let mut stream2 = b.clone().subscribe(lobby.clone());

        // publish
        let v: Value = &quot;hello&quot;.into();
        b.clone().publish(lobby.clone(), Arc::new(v.clone().into()));

        // subscribers 应该能收到 publish 的数据
        let id1: i64 = stream1.recv().await.unwrap().as_ref().try_into().unwrap();
        let id2: i64 = stream2.recv().await.unwrap().as_ref().try_into().unwrap();

        assert!(id1 != id2);

        let res1 = stream1.recv().await.unwrap();
        let res2 = stream2.recv().await.unwrap();

        assert_eq!(res1, res2);
        assert_res_ok(&amp;res1, &amp;[v.clone()], &amp;[]);

        // 如果 subscriber 取消订阅，则收不到新数据
        b.clone().unsubscribe(lobby.clone(), id1 as _);

        // publish
        let v: Value = &quot;world&quot;.into();
        b.clone().publish(lobby.clone(), Arc::new(v.clone().into()));

        assert!(stream1.recv().await.is_none());
        let res2 = stream2.recv().await.unwrap();
        assert_res_ok(&amp;res2, &amp;[v.clone()], &amp;[]);
    }
}
</code></pre>

<p>这个测试需要一系列新的改动，比如 assert_res_ok() 的接口变化了，我们需要在 src/pb/mod.rs 里添加新的 TryFrom 支持等等，详细代码你可以看 repo 里的 diff_topic。</p>

<h3 id="在处理流程中引入-pub-sub">在处理流程中引入 Pub/Sub</h3>

<p>好，再来看它和用户传入的 CommandRequest 如何发生关系。我们之前设计了 CommandService trait，它虽然可以处理其它命令，但对 Pub/Sub 相关的几个新命令无法处理，因为接口没有任何和 Topic 有关的参数：</p>

<pre><code>/// 对 Command 的处理的抽象
pub trait CommandService {
    /// 处理 Command，返回 Response
    fn execute(self, store: &amp;impl Storage) -&gt; CommandResponse;
}
</code></pre>

<p>但是如果直接修改这个接口，对已有的代码就非常不友好。所以我们还是对比着创建一个新的 trait：</p>

<pre><code>pub type StreamingResponse = Pin&lt;Box&lt;dyn Stream&lt;Item = Arc&lt;CommandResponse&gt;&gt; + Send&gt;&gt;;
pub trait TopicService {
    /// 处理 Command，返回 Response
    fn execute&lt;T&gt;(self, chan: impl Topic) -&gt; StreamingResponse;
}
</code></pre>

<p>因为 Stream 是一个 trait，在 trait 的方法里我们无法返回一个 impl Stream，所以用 trait object：<code>Pin&lt;Box\&lt;dyn Stream&gt;&gt;</code>。</p>

<p>实现它很简单，我们创建 src/service/topic_service.rs（记得在 <a href="http://mod.rs" target="_blank">mod.rs</a> 引用），然后添加：</p>

<pre><code>use futures::{stream, Stream};
use std::{pin::Pin, sync::Arc};
use tokio_stream::wrappers::ReceiverStream;

use crate::{CommandResponse, Publish, Subscribe, Topic, Unsubscribe};

pub type StreamingResponse = Pin&lt;Box&lt;dyn Stream&lt;Item = Arc&lt;CommandResponse&gt;&gt; + Send&gt;&gt;;

pub trait TopicService {
    /// 处理 Command，返回 Response
    fn execute&lt;T, S&gt;(self, topic: impl Topic) -&gt; StreamingResponse;
}

impl TopicService for Subscribe {
    fn execute&lt;T, S&gt;(self, topic: impl Topic) -&gt; StreamingResponse {
        let rx = topic.subscribe(self.topic);
        Box::pin(ReceiverStream::new(rx))
    }
}

impl TopicService for Unsubscribe {
    fn execute&lt;T, S&gt;(self, topic: impl Topic) -&gt; StreamingResponse {
        topic.unsubscribe(self.topic, self.id);
        Box::pin(stream::once(async { Arc::new(CommandResponse::ok()) }))
    }
}

impl TopicService for Publish {
    fn execute&lt;T, S&gt;(self, topic: impl Topic) -&gt; StreamingResponse {
        topic.publish(self.topic, Arc::new(self.data.into()));
        Box::pin(stream::once(async { Arc::new(CommandResponse::ok()) }))
    }
}
</code></pre>

<p>我们使用了 <a href="https://docs.rs/tokio-stream/0.1.7/tokio_stream/" target="_blank">tokio-stream</a> 的 wrapper 把一个 mpsc::Receiver 转换成 ReceiverStream。这样 Subscribe 的处理就能返回一个 Stream。对于 Unsubscribe 和 Publish，它们都返回单个值，我们使用 <code>stream::once</code> 将其统一起来。</p>

<p>同样地，要在 src/pb/mod.rs 里添加一些新的方法，比如 CommandResponse::ok()，它返回一个状态码是 OK 的 response：</p>

<pre><code>impl CommandResponse {
    pub fn ok() -&gt; Self {
        let mut result = CommandResponse::default();
        result.status = StatusCode::OK.as_u16() as _;
        result
    }
}
</code></pre>

<p>好，接下来看 src/service/mod.rs，我们可以对应着原来的 dispatch 做一个 dispatch_stream。同样地，已有的接口应该少动，我们平行添加一个新的：</p>

<pre><code>/// 从 Request 中得到 Response，目前处理所有 HGET/HSET/HDEL/HEXIST
pub fn dispatch(cmd: CommandRequest, store: &amp;impl Storage) -&gt; CommandResponse {
    match cmd.request_data {
        Some(RequestData::Hget(param)) =&gt; param.execute(store),
        Some(RequestData::Hgetall(param)) =&gt; param.execute(store),
        Some(RequestData::Hmget(param)) =&gt; param.execute(store),
        Some(RequestData::Hset(param)) =&gt; param.execute(store),
        Some(RequestData::Hmset(param)) =&gt; param.execute(store),
        Some(RequestData::Hdel(param)) =&gt; param.execute(store),
        Some(RequestData::Hmdel(param)) =&gt; param.execute(store),
        Some(RequestData::Hexist(param)) =&gt; param.execute(store),
        Some(RequestData::Hmexist(param)) =&gt; param.execute(store),
        None =&gt; KvError::InvalidCommand(&quot;Request has no data&quot;.into()).into(),
        // 处理不了的返回一个啥都不包括的 Response，这样后续可以用 dispatch_stream 处理
        _ =&gt; CommandResponse::default(),
    }
}

/// 从 Request 中得到 Response，目前处理所有 PUBLISH/SUBSCRIBE/UNSUBSCRIBE
pub fn dispatch_stream(cmd: CommandRequest, topic: impl Topic) -&gt; StreamingResponse {
    match cmd.request_data {
        Some(RequestData::Publish(param)) =&gt; param.execute(topic),
        Some(RequestData::Subscribe(param)) =&gt; param.execute(topic),
        Some(RequestData::Unsubscribe(param)) =&gt; param.execute(topic),
        // 如果走到这里，就是代码逻辑的问题，直接 crash 出来
        _ =&gt; unreachable!(),
    }
}
</code></pre>

<p>为了使用这个新的接口，Service 结构也需要相应改动：</p>

<pre><code>/// Service 数据结构
pub struct Service&lt;Store = MemTable&gt; {
    inner: Arc&lt;ServiceInner&lt;Store&gt;&gt;,
    broadcaster: Arc&lt;Broadcaster&gt;,
}

impl&lt;Store&gt; Clone for Service&lt;Store&gt; {
    fn clone(&amp;self) -&gt; Self {
        Self {
            inner: Arc::clone(&amp;self.inner),
            broadcaster: Arc::clone(&amp;self.broadcaster),
        }
    }
}

impl&lt;Store: Storage&gt; From&lt;ServiceInner&lt;Store&gt;&gt; for Service&lt;Store&gt; {
    fn from(inner: ServiceInner&lt;Store&gt;) -&gt; Self {
        Self {
            inner: Arc::new(inner),
            broadcaster: Default::default(),
        }
    }
}

impl&lt;Store: Storage&gt; Service&lt;Store&gt; {
    pub fn execute(&amp;self, cmd: CommandRequest) -&gt; StreamingResponse {
        debug!(&quot;Got request: {:?}&quot;, cmd);
        self.inner.on_received.notify(&amp;cmd);
        let mut res = dispatch(cmd, &amp;self.inner.store);

        if res == CommandResponse::default() {
            dispatch_stream(cmd, Arc::clone(&amp;self.broadcaster))
        } else {
            debug!(&quot;Executed response: {:?}&quot;, res);
            self.inner.on_executed.notify(&amp;res);
            self.inner.on_before_send.notify(&amp;mut res);
            if !self.inner.on_before_send.is_empty() {
                debug!(&quot;Modified response: {:?}&quot;, res);
            }

            Box::pin(stream::once(async { Arc::new(res) }))
        }
    }
}
</code></pre>

<p>这里，为了处理 Pub/Sub，我们引入了一个破坏性的更新。<strong>execute() 方法的返回值变成了 StreamingResponse，这就意味着所有围绕着这个方法的调用，包括测试，都需要相应更新</strong>。这是迫不得已的，不过通过构建和 CommandService/dispatch 平行的 TopicService/dispatch_stream，我们已经让这个破坏性更新尽可能地在比较高层，否则，改动会更大。</p>

<p>目前，代码无法编译通过，这是因为如下的代码，res 现在是个 stream，我们需要处理一下：</p>

<pre><code>let res = service.execute(CommandRequest::new_hget(&quot;t1&quot;, &quot;k1&quot;));
assert_res_ok(&amp;res, &amp;[&quot;v1&quot;.into()], &amp;[]);

// 需要变更为读取 stream 里的一个值
let res = service.execute(CommandRequest::new_hget(&quot;t1&quot;, &quot;k1&quot;));
let data = res.next().await.unwrap();
assert_res_ok(&amp;data, &amp;[&quot;v1&quot;.into()], &amp;[]);
</code></pre>

<p>当然，这样的改动也意味着，原本的函数需要变成 async。</p>

<p>如果是个 test，需要使用 <code>#[tokio::test]</code>。你可以自己试着把所有相关的代码都改一下。当你改到 src/network/mod.rs 里 ProstServerStream 的 process 方法时，会发现 <code>stream.send(data)</code> 时，我们目前的 data 是 Arc<CommandResponse>：</p>

<pre><code>impl&lt;S&gt; ProstServerStream&lt;S&gt;
where
    S: AsyncRead + AsyncWrite + Unpin + Send + 'static,
{
		...

    pub async fn process(mut self) -&gt; Result&lt;(), KvError&gt; {
        let stream = &amp;mut self.inner;
        while let Some(Ok(cmd)) = stream.next().await {
            info!(&quot;Got a new command: {:?}&quot;, cmd);
            let mut res = self.service.execute(cmd);
            while let Some(data) = res.next().await {
								// 目前 data 是 Arc&lt;CommandResponse&gt;，
								// 所以我们 send 最好用 &amp;CommandResponse
                stream.send(&amp;data).await.unwrap();
            }
        }
        // info!(&quot;Client {:?} disconnected&quot;, self.addr);
        Ok(())
    }
}
</code></pre>

<p>所以我们还需要稍微改动一下 src/network/stream.rs：</p>

<pre><code>// impl&lt;S, In, Out&gt; Sink&lt;Out&gt; for ProstStream&lt;S, In, Out&gt;
impl&lt;S, In, Out&gt; Sink&lt;&amp;Out&gt; for ProstStream&lt;S, In, Out&gt;
</code></pre>

<p>这会引发一系列的变动，你可以试着自己改一下。</p>

<p>如果你把所有编译错误都改正，<code>cargo test</code> 会全部通过。你也可以看 repo 里的 diff_service，看看所有改动的代码。</p>

<h3 id="继续重构-弥补设计上的小问题">继续重构：弥补设计上的小问题</h3>

<p>现在看上去大功告成，但你有没有注意，我们在撰写 src/service/topic_service.rs 时，没有写测试。你也许会说：这段代码如此简单，还有必要测试么？</p>

<p>还是那句话，测试是体验和感受接口完备性的一种手段。<strong>测试并不是为了测试实现本身，而是看接口是否好用，是否遗漏了某些产品需求</strong>。</p>

<p>当开始写测试的时候，我们就会思考：unsubscribe 接口如果遇到不存在的 subscription，要不要返回一个 404？publish 的时候遇到错误，是不是意味着客户端非正常退出了？我们要不要把它从 subscription 中移除掉？</p>

<pre><code>#[cfg(test)]
mod tests {
    use super::*;
    use crate::{assert_res_error, assert_res_ok, dispatch_stream, Broadcaster, CommandRequest};
    use futures::StreamExt;
    use std::{convert::TryInto, time::Duration};
    use tokio::time;

    #[tokio::test]
    async fn dispatch_publish_should_work() {
        let topic = Arc::new(Broadcaster::default());
        let cmd = CommandRequest::new_publish(&quot;lobby&quot;, vec![&quot;hello&quot;.into()]);
        let mut res = dispatch_stream(cmd, topic);
        let data = res.next().await.unwrap();
        assert_res_ok(&amp;data, &amp;[], &amp;[]);
    }

    #[tokio::test]
    async fn dispatch_subscribe_should_work() {
        let topic = Arc::new(Broadcaster::default());
        let cmd = CommandRequest::new_subscribe(&quot;lobby&quot;);
        let mut res = dispatch_stream(cmd, topic);
        let id: i64 = res.next().await.unwrap().as_ref().try_into().unwrap();
        assert!(id &gt; 0);
    }

    #[tokio::test]
    async fn dispatch_subscribe_abnormal_quit_should_be_removed_on_next_publish() {
        let topic = Arc::new(Broadcaster::default());
        let id = {
            let cmd = CommandRequest::new_subscribe(&quot;lobby&quot;);
            let mut res = dispatch_stream(cmd, topic.clone());
            let id: i64 = res.next().await.unwrap().as_ref().try_into().unwrap();
            drop(res);
            id as u32
        };

        // publish 时，这个 subscription 已经失效，所以会被删除
        let cmd = CommandRequest::new_publish(&quot;lobby&quot;, vec![&quot;hello&quot;.into()]);
        dispatch_stream(cmd, topic.clone());
        time::sleep(Duration::from_millis(10)).await;

        // 如果再尝试删除，应该返回 KvError
        let result = topic.unsubscribe(&quot;lobby&quot;.into(), id);
        assert!(result.is_err());
    }

    #[tokio::test]
    async fn dispatch_unsubscribe_should_work() {
        let topic = Arc::new(Broadcaster::default());
        let cmd = CommandRequest::new_subscribe(&quot;lobby&quot;);
        let mut res = dispatch_stream(cmd, topic.clone());
        let id: i64 = res.next().await.unwrap().as_ref().try_into().unwrap();

        let cmd = CommandRequest::new_unsubscribe(&quot;lobby&quot;, id as _);
        let mut res = dispatch_stream(cmd, topic);
        let data = res.next().await.unwrap();

        assert_res_ok(&amp;data, &amp;[], &amp;[]);
    }

    #[tokio::test]
    async fn dispatch_unsubscribe_random_id_should_error() {
        let topic = Arc::new(Broadcaster::default());

        let cmd = CommandRequest::new_unsubscribe(&quot;lobby&quot;, 9527);
        let mut res = dispatch_stream(cmd, topic);
        let data = res.next().await.unwrap();

        assert_res_error(&amp;data, 404, &quot;Not found: subscription 9527&quot;);
    }
}
</code></pre>

<p>在撰写这些测试，并试图使测试通过的过程中，我们又进一步重构了代码。具体的代码变更，你可以参考 repo 里的 diff_refactor。</p>

<h3 id="让客户端能更好地使用新的接口">让客户端能更好地使用新的接口</h3>

<p>目前，我们 ProstClientStream 还是一个统一的 execute() 方法：</p>

<pre><code>impl&lt;S&gt; ProstClientStream&lt;S&gt;
where
    S: AsyncRead + AsyncWrite + Unpin + Send,
{
	  ...

    pub async fn execute(&amp;mut self, cmd: CommandRequest) -&gt; Result&lt;CommandResponse, KvError&gt; {
        let stream = &amp;mut self.inner;
        stream.send(&amp;cmd).await?;

        match stream.next().await {
            Some(v) =&gt; v,
            None =&gt; Err(KvError::Internal(&quot;Didn't get any response&quot;.into())),
        }
    }
}
</code></pre>

<p>它并没有妥善处理 SUBSCRIBE。为了支持 SUBSCRIBE，我们需要两个接口：execute_unary 和 execute_streaming。在 src/network/mod.rs 修改这个代码：</p>

<pre><code>impl&lt;S&gt; ProstClientStream&lt;S&gt;
where
    S: AsyncRead + AsyncWrite + Unpin + Send + 'static,
{
    ...

    pub async fn execute_unary(
        &amp;mut self,
        cmd: &amp;CommandRequest,
    ) -&gt; Result&lt;CommandResponse, KvError&gt; {
        let stream = &amp;mut self.inner;
        stream.send(cmd).await?;

        match stream.next().await {
            Some(v) =&gt; v,
            None =&gt; Err(KvError::Internal(&quot;Didn't get any response&quot;.into())),
        }
    }

    pub async fn execute_streaming(self, cmd: &amp;CommandRequest) -&gt; Result&lt;StreamResult, KvError&gt; {
        let mut stream = self.inner;

        stream.send(cmd).await?;
        stream.close().await?;

        StreamResult::new(stream).await
    }
}
</code></pre>

<p>注意，因为 execute_streaming 里返回 Box:pin(stream)，我们需要对 ProstClientStream 的 S 限制是 &lsquo;static，否则编译器会抱怨。这个改动会导致使用 execute() 方法的测试都无法编译，你可以试着修改掉它们。</p>

<p>此外我们还创建了一个新的文件 src/network/stream_result.rs，用来帮助客户端更好地使用 execute_streaming() 接口。所有改动的具体代码可以看 repo 中的 diff_client。</p>

<p>现在，代码一切就绪。打开一个命令行窗口，运行：<code>RUST_LOG=info cargo run --bin kvs --quiet，</code>然后在另一个命令行窗口，运行：<code>RUST_LOG=info cargo run --bin kvc --quiet</code>。</p>

<p>此时，服务器和客户端都收到了彼此的请求和响应，即便混合 HSET/HGET 和 PUBLISH/SUBSCRIBE 命令，一切都依旧处理正常！今天我们做了一个比较大的重构，但比预想中对原有代码的改动要小，这简直太棒了！</p>

<h2 id="小结">小结</h2>

<p>当一个项目越来越复杂，且新加的功能并不能很好地融入已有的系统时，大的重构是不可避免的。在重构的时候，我们一定要首先要弄清楚现有的流程和架构，然后再思考如何重构，这样对系统的侵入才是最小的。</p>

<p>重构一般会带来对现有测试的破坏，在修复被破坏的测试时，我们要注意不要变动原有测试的逻辑。在做因为新功能添加导致的重构时，如果伴随着大量测试的删除和大量新测试的添加，那么，说明要么原来的测试写得很有问题，要么重构对原有系统的侵入性太强。我们要尽量避免这种事情发生。</p>

<p><strong>在架构和设计都相对不错的情况下，撰写代码的终极目标是对使用者友好的抽象</strong>。所谓对使用者友好的抽象，是指让别人调用我们写的接口时，不用想太多，接口本身就是自解释的。</p>

<p>如果你仔细阅读 diff_client，可以看到类似 StreamResult 这样的抽象。它避免了调用者需要了解如何手工从 Stream 中取第一个值作为 subscription_id 这样的实现细节，直接替调用者完成了这个工作，并以一个优雅的 ID 暴露给调用者。</p>

<p>你可以仔细阅读这一讲中的代码，好好品味这些接口的设计。它们并非完美，世上没有完美的代码，只有不断完善的代码。如果把一行行代码比作一段段文字，起码它们都需要努力地推敲和不断地迭代。</p>

<h3 id="思考题">思考题</h3>

<ol>
<li>现在我们的系统对 Pub/Sub 已经有比较完整的支持，但你有没有注意到，有一个潜在的内存泄漏的 bug。如果客户端 A subscribe 了 Topic M，但客户端意外终止，且随后也没有任何人往 Topic M publish 消息。这样，A 的 subscription 就一直放在表中。你能做一个 GC 来处理这种情况么？</li>
<li>Redis 还支持 PSUBSCRIBE，也就是说除了可以 subscribe “chat” 这样固定的 topic，还可以是 “chat.*”，一并订阅所有 “chat”、“chat.rust”、“chat.elixir” 。想想看，如果要支持 PSUBSCRIBE，你该怎么设计 Broadcaster 里的两张表？</li>
</ol>

<p>欢迎在留言区分享你的思考和学习感悟。感谢你的收听，如果觉得有收获，也欢迎分享给你身边的朋友，邀他一起讨论。恭喜你完成了Rust学习的第42次打卡，我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c5a9a9a9fcf1f4f4f5f285a2a8a4aca9eba6aaa8" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a53b7388b7755',t:'MTczNDEzODc2MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
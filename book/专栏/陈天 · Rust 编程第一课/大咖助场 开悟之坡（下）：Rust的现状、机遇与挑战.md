<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=大咖助场&#32;开悟之坡（下）：Rust的现状、机遇与挑战>
        <link rel="icon" href="/static/favicon.png">
        <title>大咖助场 开悟之坡（下）：Rust的现状、机遇与挑战 </title>
        
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
                            <h1 id="title" data-id="大咖助场 开悟之坡（下）：Rust的现状、机遇与挑战" class="title">大咖助场 开悟之坡（下）：Rust的现状、机遇与挑战</h1>
                            <div><p>你好，我是张汉东。</p>

<p>上篇我们聊了Rust语言的现状和机遇，从语言自身的成熟度、语言的生态和应用场景，以及语言的可持续发展能力这三个方面，比较系统地说明Rust发展相对成熟的现状。</p>

<p>Rust 语言作为一门新生语言，虽然目前倍受欢迎，但是面临的挑战还很多。我们今天就聊一聊这个话题。</p>

<p>挑战主要来自两个方面：</p>

<ol>
<li><strong>领域的选择</strong>。一门语言唱的再好，如果不被应用，也是没有什么用处。Rust 语言当前面临的挑战就是在领域中的应用。而目前最受关注的是，Rust 进入 Linux 内核开发，如果成功，其意义是划时代的。</li>
<li><strong>语言自身特性的进化</strong>。Rust 语言还有很多特性需要支持和进化，后面也会罗列一些待完善的相关特性。</li>
</ol>

<h3 id="rust-for-linux-的进展和预判">Rust For Linux 的进展和预判</h3>

<p>从 2020 年 6 月，Rust 进入<code>Linux</code> 就开始成为一个话题。<code>Linux</code> 创建者 Linus 在当时的开源峰会和嵌入式<code>Linux</code> 会议上，谈到了为开源内核寻找未来维护者的问题。</p>

<p>简单跟你讲一讲背景情况。</p>

<p>Linus 提到：“内核很无聊，至少大多数人认为它很无聊。许多新技术对很多人来说应该更加有趣。事实证明，开源内核很难找到维护者。虽然有很多人编写代码，但是很难找到站在上游对别人代码进行 Review 的人选。这不仅仅是来自其他维护者的信任，也来自所有编写代码的人的信任……这只是需要时间的”。</p>

<p><strong>而 Rust 作为一门天生安全的语言，作为<code>C</code>的备选语言，在帮助内核开发者之间建立彼此的信任，是非常有帮助的</strong>。三分之二的 Linux 内核安全漏洞( <a href="https://static.sched.com/hosted_files/lssna19/d6/kernel-modules-in-rust-lssna2019.pdf" target="_blank">PDF</a> )来自内存安全问题，在 Linux 中引入 Rust 会让其更加安全，这目前基本已经达成一种共识。</p>

<p>而且在今年（2021）的开源峰会上， Linus 说：“我认为C语言是一种伟大的语言，对我来说，C 语言确实是一种在相当低的水平上控制硬件的方法。因此，当我看到C语言代码时，我可以非常接近地猜测编译器的工作，它是如此接近硬件，以至于你可以用它来做任何事情。”</p>

<p>“但是，C语言微妙的类型交互，并不总是合乎逻辑的，对几乎所有人来说都是陷阱，它们很容易被忽视，而在内核中，这并不总是一件好事。”</p>

<p>“<strong>Rust 语言是我看到的、第一种看起来像是真的可以解决问题的语言</strong>。人们现在已经谈论Rust在内核中的应用很久了，但它还没有完成，可能在明年，我们会开始看到一些首次用Rust编写的无畏模块，也许会被整合到主线内核中。”</p>

<p>Linus 认为 <code>Linux</code> 之所以如此长青，其中一个重要的基石就是乐趣（Fun），并且乐趣也是他一直追求的东西。当人们讨论使用Rust编写一些<code>Linux</code>内核模块的可能性时，乐趣就出现了。</p>

<h3 id="大会进展">大会进展</h3>

<p>在刚过去的 2021 年 9 月 的 Linux Plumbers 大会上， 再一次讨论了 Rust 进入 Linux 内核的进展。</p>

<p>首先是Rust的参与角色问题。</p>

<p>Rust for Linux 的主力开发者 Miguel Ojedal 说，Rust 如果进入内核，就应该是一等公民的角色。Linus 则回答，内核社区几乎肯定会用该语言进行试验。</p>

<p>对Rust代码的review问题也简单讨论过。</p>

<p>Rust 进入内核肯定会有一些维护者需要学习该语言，用来 review Rust 代码。Linus 说， Rust 并不难懂，内核社区任何有能力 review patch 的人，都应该掌握 Rust 语言到足以 Review 该语言代码的程度。</p>

<p>另外还有一些Rust自身特性的稳定问题：</p>

<ol>
<li>目前内核工作还在使用一些 Unstable 的 Rust 特性，导致兼容性不够好，不能确保以后更新的 Rust 编译器能正常编译相关代码。</li>
</ol>

<p>Ojedal 说，但是如果 Rust 进入 Linux 内核，就会改变这种情况，对于一些 Unstable Rust 特性，Rust 官方团队也会考虑让其稳定。这是一种推动力，迟早会建立一个只使用 Rust 稳定版的内核，到时候兼容问题就会消失。</p>

<ol>
<li>另一位内核开发者 Thomas Gleixner 担心 Rust 并没有正式支持内存顺序，这可能会有问题。</li>
</ol>

<p>但是另一位从事三十年cpp 并发编程的 Linux 内核维护者 Paul McKenney 则写了<a href="https://paulmck.livejournal.com/62436.html" target="_blank">一系列文章</a>来探讨 Rust 社区该如何就Rust 进入 Linux 内核这件事正确处理内存顺序模型。对此我也写了另一篇文章<a href="https://mp.weixin.qq.com/s/9OSjVWj14lwf-ICYhU144g" target="_blank">【我读】Rust 语言应该使用什么内存模型？</a> 。</p>

<ol>
<li>关于 Rust 对 GCC 的支持，其中 <code>rustc_codegen_gcc</code>进展最快，目前已通过了部分的 <code>rustc</code> 测试，<code>rustc_codegen_llvm</code>是目前的主要开发项目，<code>Rust GCC</code>预计在 1~2 年内完成。</li>
</ol>

<p>这次大会的结论有2点：</p>

<ol>
<li>Rust <strong>肯定</strong>会在 Linux 内核中进行一次具有时代意义的实验。</li>
<li>Rust 进入 Linux 内核，对推动 Rust 进化具有很重要的战略意义。</li>
</ol>

<h3 id="最新消息">最新消息</h3>

<p>2021 年 11 月 11 日，在 Linux 基金会网站上，又放出另一场录制的网络会议： <a href="https://linuxfoundation.org/webinars/rust-for-linux-writing-abstractions-and-drivers/" target="_blank">Rust for Linux：编写安全抽象和驱动程序</a>，该视频中 Miguel Ojedal 介绍了 Rust 如何在内核中工作，包括整体基础设施、编译模型、文档、测试和编码指南等。</p>

<p>我对这部分视频内容做了一个简要总结，你可以对照要点找自己需要的看一看。</p>

<ol>
<li>介绍 Unsafe Rust 和 Safe Rust。</li>
<li>在 Linux 内核中使用 Rust ，采用一个理念：封装 Unsafe 操作，提供一个安全抽象给内核开发者使用。这个安全抽象位于 <a href="https://github.com/Rust-for-Linux/linux/tree/rust/rust" target="_blank">https://github.com/Rust-for-Linux/linux/tree/rust/rust</a> 的 <code>kernel</code> 模块中。</li>
<li>给出一个简单的示例来说明如何编写内核驱动。</li>
<li>对比 C 语言示例，给出在 Rust 中什么是 Safety 的行为。</li>
<li>介绍了文档、测试和遵循的编码准则。</li>
</ol>

<p>综合上面我们了解到的这些信息，可以推测，Rust for Linux 在不远的将来会进入到 Linux 进行一次试验，这次试验的意义是划时代的。如果试验成功，那么就意味着 Rust 正式从 C 语言手里拿到了时代的交接棒。</p>

<h2 id="rust-语言特性的完善">Rust 语言特性的完善</h2>

<p>下面来聊一聊最近Rust语言又完善了哪些特性。特别说明一下，这些本来就是高级知识，是Rust 语言的挑战，所以这些知识点你现在也许不太理解，但不用害怕，这些只是 Rust 语言进化路上必须要完善的东西，改进只是为了让 Rust 更好。目前并不影响你学习和使用 Rust 。</p>

<p>我们会讲4个已完善的特性，最后也顺带介绍一下还有哪些待完善的特性，供你参考。</p>

<p><img src="assets/354a4559160553c7b196f7cae4af695c.jpg" alt="图片" /></p>

<h3 id="安全-i-o-问题">安全 I/O 问题</h3>

<p>最近Rust官方合并了一个<a href="*https://github.com/rust-lang/rfcs/blob/master/text/3128-io-safety.md*" target="_blank">RFC</a> ，通过引入I/O安全的概念和一套新的类型和特质，为<code>AsRawFd</code>和相关特质的用户提供关于其原始资源句柄的保证，从而弥补Rust中封装边界的漏洞。</p>

<p>之前Rust 标准库提供了 I/O 安全性，保证程序持有私有的原始句柄（raw handle），其他部分无法访问它。</p>

<p>但是 <code>FromRawFd::from_raw_fd</code> 是 Unsafe 的，所以在 Safe Rust中无法做到 <code>File::from_raw(7)</code> 这种事，在这个文件描述符上面进行<code>I/O</code> 操作，而这个文件描述符可能被程序的其他部分私自持有。</p>

<p>而且，很多 API 通过接受原始句柄来进行 I/O 操作：</p>

<pre><code>pub fn do_some_io&lt;FD: AsRawFd&gt;(input: &amp;FD) -&gt; io::Result&lt;()&gt; {
    some_syscall(input.as_raw_fd())
}
</code></pre>

<p><code>AsRawFd</code>并没有限制<code>as_raw_fd</code>的返回值，所以<code>do_some_io</code>最终可以在任意的<code>RawFd</code>值上进行 <code>I/O</code>操作，甚至可以写<code>do_some_io(&amp;7)</code>，因为<code>RawFd</code>本身实现了<code>AsRawFd</code>。这可能会导致程序访问错误的资源。甚至通过创建在其他部分私有的句柄别名来打破封装边界，导致一些诡异的远隔作用（Action at a distance）。</p>

<blockquote>
<p><strong>远隔作用（Action at a distance）</strong>是一种程式设计中的<strong>反模式</strong>，是指程式某一部分的行为会广泛的受到程式其他部分指令的影响，而且要找到影响其他程式的指令很困难，甚至根本无法进行。</p>
</blockquote>

<p>在一些特殊的情况下，违反 I/O 安全甚至会导致内存安全。</p>

<p>所以Rust新增了<code>OwnedFd</code> 和 <code>BorrowedFd&lt;'fd&gt;</code>这两种类型，用于替代 <code>RawFd</code> ，对句柄值赋予所有权语义，代表句柄值的拥有和借用。<code>OwnedFd</code> 拥有一个 <code>fd</code> ，会在析构的时候关闭它。<code>BorrowedFd&lt;'fd&gt;</code> 中的生命周期参数，表示对这个 <code>fd</code> 的访问被借用多长时间。</p>

<p>对于Windows来说，也有类似的类型，但都是<code>Handle</code>和<code>Socket</code>形式。</p>

<p><img src="assets/a519d4d672c487c41c0388d54d375560.jpg" alt="" /></p>

<p>和其他类型相比，<code>I/O</code> 类型并不区分可变和不可变。操作系统资源可以在<code>Rust</code>的控制之外以各种方式共享，所以<code>I/O</code>可以被认为是使用内部可变性。</p>

<p>然后新增了三个概念，<code>AsFd</code>、<code>Into&lt;OwnedFd&gt;</code>和<code>From&lt;OwnedFd&gt;</code>。</p>

<p>这三个概念是<code>AsRawFd::as_raw_fd</code>、<code>IntoRawFd::into_raw_fd</code>和<code>FromRawFd::from_raw_fd</code>的概念性替代，分别适用于大多数使用情况。它们以<code>OwnedFd</code>和<code>BorrowedFd</code>的方式工作，所以它们自动执行其<code>I/O</code>安全不变性。</p>

<pre><code>pub fn do_some_io&lt;FD: AsFd&gt;(input: &amp;FD) -&gt; io::Result&lt;()&gt; {
    some_syscall(input.as_fd())
}
</code></pre>

<p>使用这个类型，就会避免之前那个问题。由于<code>AsFd</code>只针对那些适当拥有或借用其文件描述符的类型实现，这个版本的<code>do_some_io</code>不必担心被传递假的或悬空的文件描述符。</p>

<h3 id="错误处理改进-try">错误处理改进 Try</h3>

<p>目前 Rust 允许通过 <code>?</code> 操作符，自动返回 <code>Result&lt;T,E&gt;</code> 的 <code>Err(e)</code> ，但是对于 <code>Ok(o)</code> 还需要手动包装。</p>

<p>比如：</p>

<pre><code>fn foo() -&gt; Result&lt;PathBuf, io::Error&gt; {
    let base = env::current_dir()?;
    Ok(base.join(&quot;foo&quot;))
}
</code></pre>

<p>那么这就引出了一个术语： Ok-Wrapping 。很明显，这个写法不够优雅，还有很大的改进空间。</p>

<p>因此 Rust 官方成员 withoutboats 开发了一个库 <a href="https://github.com/withoutboats/fehler" target="_blank">fehler</a>，引入了一个 throw 语法。用法如下：</p>

<pre><code>#[throws(i32)]
fn foo(x: bool) -&gt; i32 {
    if x {
        0
    } else {
        throw!(1);
    }
}

// 上面foo函数错误处理等价于下面bar函数

fn bar(x: bool) -&gt; Result&lt;i32, i32&gt; {
    if x {
        Ok(0)
    } else {
        Err(1)
    }
}
</code></pre>

<p>通过 throw 宏语法，来帮助开发者省略 Ok-wrapping 和 Err-wrapping 的手动操作。这个库一时在社区引起了一些讨论，它也在促进着 Rust 错误处理的体验提升。</p>

<p>于是错误处理就围绕着 Ok-wrapping 和 Err-wrapping 这两条路径发展着，该如何设计语法才更加优雅，成为了讨论的焦点。</p>

<p>经过很久很久的讨论，try-trait-v2 RFC 被合并了，意味着一个确定的方案出现了。在这个方案中，引入了一个新类型<code>ControlFlow</code>和一个新的trait <code>FromResidual</code>。</p>

<p><code>ControlFlow</code> 的源码：</p>

<pre><code>enum ControlFlow&lt;B, C = ()&gt; {
    /// Exit the operation without running subsequent phases.
    Break(B),
    /// Move on to the next phase of the operation as normal.
    Continue(C),
}

impl&lt;B, C&gt; ControlFlow&lt;B, C&gt; {
    fn is_break(&amp;self) -&gt; bool;
    fn is_continue(&amp;self) -&gt; bool;
    fn break_value(self) -&gt; Option&lt;B&gt;;
    fn continue_value(self) -&gt; Option&lt;C&gt;;
}
</code></pre>

<p><code>ControlFlow</code> 中包含了两个值：</p>

<ul>
<li><code>ControlFlow::Break</code>，表示提前退出。但不一定是<code>Error</code> 的情况，也可能是 <code>Ok</code>。</li>
<li><code>ControlFlow::Continue</code>，表示继续。</li>
</ul>

<p>新的trait <code>FromResidual</code>：</p>

<pre><code>trait FromResidual&lt;Residual = &lt;Self as Try&gt;::Residual&gt; {
    fn from_residual(r: Residual) -&gt; Self;
}
</code></pre>

<p>Residual 这个单词有“剩余”的意思，因为要把 Result/Option/ ControlFlow 之类的类型，拆分成两部分（两条路径），用这个词也就好理解了。</p>

<p>而 <code>Try</code> trait 继承自 <code>FromResidual</code> trait ：</p>

<pre><code>pub trait Try: FromResidual {
    /// The type of the value consumed or produced when not short-circuiting.
    type Output;

    /// A type that &quot;colours&quot; the short-circuit value so it can stay associated
    /// with the type constructor from which it came.
    type Residual;

    /// Used in `try{}` blocks to wrap the result of the block.
    fn from_output(x: Self::Output) -&gt; Self;

    /// Determine whether to short-circuit (by returning `ControlFlow::Break`)
    /// or continue executing (by returning `ControlFlow::Continue`).
    fn branch(self) -&gt; ControlFlow&lt;Self::Residual, Self::Output&gt;;
}

pub trait FromResidual&lt;Residual = &lt;Self as Try&gt;::Residual&gt; {
    /// Recreate the type implementing `Try` from a related residual
    fn from_residual(x: Residual) -&gt; Self;
}
</code></pre>

<p>所以，在 <code>Try</code> trait 中有两个关联类型：</p>

<ul>
<li><code>Output</code>，如果是 Result 的话，就对应 Ok-wrapping 。</li>
<li><code>Residual</code>，如果是 Result 的话，就对应 Err-wrapping 。</li>
</ul>

<p>所以，现在 <code>?</code> 操作符的行为就变成了：</p>

<pre><code>match Try::branch(x) {
    ControlFlow::Continue(v) =&gt; v,
    ControlFlow::Break(r) =&gt; return FromResidual::from_residual(r),
}
</code></pre>

<p>然后内部给 Rusult 实现 <code>Try</code> ：</p>

<pre><code>impl&lt;T, E&gt; ops::Try for Result&lt;T, E&gt; {
    type Output = T;
    type Residual = Result&lt;!, E&gt;;

    #[inline]
    fn from_output(c: T) -&gt; Self {
        Ok(c)
    }

    #[inline]
    fn branch(self) -&gt; ControlFlow&lt;Self::Residual, T&gt; {
        match self {
            Ok(c) =&gt; ControlFlow::Continue(c),
            Err(e) =&gt; ControlFlow::Break(Err(e)),
        }
    }
}

impl&lt;T, E, F: From&lt;E&gt;&gt; ops::FromResidual&lt;Result&lt;!, E&gt;&gt; for Result&lt;T, F&gt; {
    fn from_residual(x: Result&lt;!, E&gt;) -&gt; Self {
        match x {
            Err(e) =&gt; Err(From::from(e)),
        }
    }
}
</code></pre>

<p>再给 <code>Option/Poll</code> 实现 <code>Try</code> ，就能达成错误处理大一统。</p>

<h3 id="泛型关联类型-gat">泛型关联类型 <code>GAT</code></h3>

<p>泛型关联类型在 <a href="https://github.com/rust-lang/rfcs/blob/master/text/1598-generic_associated_types.md" target="_blank">RFC 1598</a> 中被定义。该功能特性经常被对比于 Haskell 中的 HKT(Higher Kinded Type)，也就是高阶类型。</p>

<p>虽然这两个类型相似，但是 Rust 并没有把 Haskell 的<code>HKT</code> 原样照搬，而是针对 Rust 自身特性给出GAT(Generic associated type) 的概念。目前<code>GAT</code> 支持的进展可以在<a href="https://github.com/rust-lang/rust/issues/44265" target="_blank">issues #44265</a> 中被跟踪，也许在年内可以稳定。</p>

<p>什么是泛型关联类型？ 见下面代码：</p>

<pre><code>trait Iterable {
    type Item&lt;'a&gt;; // 'a 也是泛型参数
}

trait Foo {
    type Bar&lt;T&gt;;
}
</code></pre>

<p>就是这样一个简单的语法，让我们在关联类型里也能参与类型构造，就是实现起来却非常复杂。</p>

<p>但无论多复杂，这个特性是 Rust 语言必须要支持的功能，它非常有用。<strong>最典型的就是用来实现流迭代器：</strong></p>

<pre><code>trait StreamingIterator {
    type Item&lt;'a&gt;;
    fn next&lt;'a&gt;(&amp;'a mut self) -&gt; Option&lt;Self::Item&lt;'a&gt;&gt;;
}
</code></pre>

<p>现在 Rust 还不支持这种写法。这种写法可以解决当前迭代器性能慢的问题。-
比如标准库中的<code>std::io::lines</code> 方法，可以为 <code>io::BufRead</code> 类型生成一个迭代器，但是它当前只能返回 <code>io::Result&lt;Vec&lt;u8&gt;&gt;</code>，这就意味着它会为每一行进行内存分配，而产生一个新的<code>Vec&lt;u8&gt;</code> ，导致迭代器性能很慢。<a href="https://stackoverflow.com/questions/45455275/why-is-my-rust-version-of-wc-slower-than-the-one-from-gnu-coreutils" target="_blank">StackOverflow上有这个问题的讨论和优化方案</a>。</p>

<p>但是如果支持 <code>GAT</code> 的话，解决这个问题将变得非常简单：</p>

<pre><code>trait Iterator {
    type Item&lt;'s&gt;;
    fn next(&amp;mut self) -&gt; Option&lt;Self::Item&lt;'_&gt;&gt;;
}

impl&lt;B: BufRead&gt; Iterator for Lines&lt;B&gt; {
    type Item&lt;'s&gt; = io::Result&lt;&amp;'s str&gt;;
    fn next(&amp;mut self) -&gt; Option&lt;Self::Item&lt;'_&gt;&gt; { … }
}
</code></pre>

<p><code>GAT</code> 的实现还能推进“异步 trait”的支持。目前 Rust 异步还有很多限制，比如 trait 无法支持 <code>async</code> 方法，也是因为<code>GAT</code> 功能未完善而导致的。</p>

<h3 id="泛型特化specialization">泛型特化Specialization</h3>

<p>泛型特化这个概念，对应 Cpp 的模版特化。但是 Cpp 对特化的支持是相当完善，而 Rust 中特化还未稳定。</p>

<p>在 <a href="https://rust-lang.github.io/rfcs/1210-impl-specialization.html" target="_blank">RFC #1210</a> 中定义了 Rust 的泛型特化的实现标准，在 <a href="https://github.com/rust-lang/rust/issues/31844" target="_blank">issue #31844</a> 对其实现状态进行了跟踪。目前还有很多未解决的问题。</p>

<p>什么是泛型特化呢？</p>

<pre><code>trait Example {
    type Output;
    fn generate(self) -&gt; Self::Output;
}

impl&lt;T&gt; Example for T {
    type Output = Box&lt;T&gt;;
    fn generate(self) -&gt; Box&lt;T&gt; { Box::new(self) }
}

impl Example for bool {
    type Output = bool;
    fn generate(self) -&gt; bool { self }
}
</code></pre>

<p><strong>简单来说，就是可以为泛型以及更加具体的类型来实现同一个 trait 。在调用该trait 方法时，倾向于优先使用更具体的类型实现</strong>。这就是对“泛型特化”最直观的一个理解。</p>

<p>泛型特化带来两个重要意义：</p>

<ol>
<li>性能优化。特化扩展了零成本抽象的范围，可以为某个统一抽象下的具体实现，定制高性能实现。</li>
<li>代码重用。泛型特化可以提供一些默认（但不完整的）实现，某些情况下可以减少重复代码。</li>
</ol>

<p>其实曾经特化还要为“高效继承（efficient-inheritance）”做为实现基础，但是现在高效继承这个提议并未被正式采纳。但我想，作为代码高效重用的一种手段，在未来肯定会被重新提及。</p>

<p>泛型特化功能，离最终稳定还有很长的路，目前官方正准备稳定特化的一个子集（subset）叫 <code>min_specialization</code>，旨在让泛型特化有一个最小化可用（mvp）的实现，在此基础上再慢慢稳定整体功能。现在 <code>min_specialization</code> 还没有具体稳定的日期，如果要使用此功能，只能在 Nightly Rust 下添加 <code>#![feature(min_specialization)]</code> 来使用。</p>

<pre><code>#![feature(min_specialization)]
use std::fmt::Debug;

trait Destroy {
    fn destroy(self);
}

impl&lt;T: Debug&gt; Destroy for T {
    default fn destroy(self) {
        println!(&quot;Destroyed something!&quot;);
    }
}

struct Special;

impl Destroy for Special {
    fn destroy(self) {
        println!(&quot;Destroyed Special something!&quot;);
    }
}

fn main() {
    &quot;hello&quot;.destroy(); // Destroyed something!
    let sp = Special;
    sp.destroy(); // Destroyed Special something!
}
</code></pre>

<h3 id="其他待完善特性">其他待完善特性</h3>

<p><img src="assets/20b2fd045a0f23635694e44f752a994b.jpg" alt="图片" /></p>

<h3 id="异步-async-trait-async-drop">异步 async trait、async drop</h3>

<p>Rust 目前异步虽然早已稳定，但还有很多需要完善的地方。为此，官方创建了异步工作组，并且创建了<a href="https://rust-lang.github.io/async-fundamentals-initiative/index.html" target="_blank">异步基础计划</a>来推动这一过程。</p>

<p>对于异步 trait 功能，首先会稳定的一个 mvp 功能是：trait 中的静态的 <code>async fn</code> 方法。</p>

<pre><code>trait Service {
    async fn request(&amp;self, key: i32) -&gt; Response;
}

struct MyService {
    db: Database
}

impl Service for MyService {
    async fn request(&amp;self, key: i32) -&gt; Response {
        Response {
            contents: self.db.query(key).await.to_string()
        }
    }
}
</code></pre>

<p>在 trait 中支持 <code>async fn</code> 非常有用。但是目前只能通过 <code>async-trait</code> 来支持这个功能。因为当前 trait 中直接写 <code>async fn</code> 不是动态安全的（dyn safety，之前叫对象安全）。</p>

<p>现在这个 mvp 功能提出将 <code>async</code> <code>fn</code> 脱糖为静态分发的 trait，比如这样：</p>

<pre><code>trait Service {
    type RequestFut&lt;'a&gt;: Future&lt;Output = Response&gt;
    where
        Self: 'a;
    fn request(&amp;self, key: i32) -&gt; RequestFut;
}

impl Service for MyService {
    type RequestFut&lt;'a&gt; = impl Future + 'a
    where
        Self: 'a;
    fn request&lt;'a&gt;(&amp;'a self, key: i32) -&gt; RequestFut&lt;'a&gt; {
        async { ... }
    }
}
</code></pre>

<p>对于 异步 drop 功能，目前也给出了一个方案，但没有类似 mvp 的落地计划。更多解释可以去查看异步基础计划的内容。</p>

<h3 id="协程的稳定化">协程的稳定化</h3>

<p>目前 Rust 的异步是基于一种半协程机制生成器（ Generator） 来实现的，但生成器特性并未稳定。围绕“生成器特性”稳定的话题，在 Rust 论坛不定期会提出，因为生成器这个特性在其他语言中，也是比较常见且有用的特性。</p>

<p>但目前 Rust 团队对此并没有一个确切的设计，当前 Rust 内部的生成器机制只是为了稳定实现 异步编程而采取的临时设计。 所以这个特性也是 Rust 语言未来的挑战之一。</p>

<h3 id="simd">SIMD</h3>

<p>众所周知，计算机程序需要编译成指令才能让 CPU 识别并执行运算。所以，CPU 指令处理数据的能力，是衡量 CPU 性能的重要指标。</p>

<p>为了提高 CPU 指令处理数据的能力，半导体厂商在 CPU 中推出了一些可以同时并行处理多个数据的指令 —— <strong>SIMD指令</strong>。SIMD 的全称是 Single Instruction Multiple Data，中文名“单指令多数据”。顾名思义，一条指令处理多个数据。</p>

<p>经过多年的发展，支持 SIMD 的指令集有很多。各种 CPU 架构都提供各自的 SIMD 指令集，比如 <code>X86/MMX/SSE/AVX</code>等指令集。Rust 目前有很多架构平台下的指令集，但目前还未稳定，你可以在 <code>core::arch</code> 模块下找到，但这些都是可以具体架构平台相关的，并不能方便编写跨平台的 SIMD 代码。如果想编写跨平台 SIMD 代码，需要用到第三方库 <a href="https://github.com/rust-lang/packed_simd" target="_blank">packed_simd</a> 。</p>

<p>最近几天，Rust 官方团队发布了 <code>portable-simd</code> ，你可以在 Nightly 下使用这个库来代替 <a href="https://github.com/rust-lang/packed_simd" target="_blank">packed_simd</a> 了。这个库使得用 Rust 开发跨平台 SIMD 更加容易和安全。在不久的将来，也会引入到标准库中稳定下来。</p>

<h3 id="新的-asm-支持">新的 asm! 支持</h3>

<p><code>asm!</code> 宏允许在 Rust 中内联汇编。</p>

<p>在 <a href="https://github.com/rust-lang/rfcs/blob/master/text/2873-inline-asm.md" target="_blank">RFC #2873</a> 中规定了新的 asm!宏语法，将用于兼容 ARM、x86 和 RISC-V 等架构，方便在未来添加更多架构支持。之前的 asm! 宏被重命名为 llvm_asm!。目前新的 asm! 已经接近稳定状态，可在 <a href="https://github.com/rust-lang/rust/issues/72016" target="_blank">issue #72016</a> 中跟踪。</p>

<p>总的来说，就是让 <code>asm!</code> 宏更加通用，相比于 <code>llvm_asm!</code>，它有更好的语法。</p>

<pre><code>// 旧的 asm! 宏写法
let i: u64 = 3;
let o: u64;
unsafe {
    asm!(
        &quot;mov {0}, {1}&quot;,
        &quot;add {0}, {number}&quot;,
        out(reg) o,
        in(reg) i,
        number = const 5,
    );
}
assert_eq!(o, 8);

// 新的 asm! 宏写法：
let x: u64 = 3;
let y: u64;
unsafe {
    asm!(&quot;add {0}, {number}&quot;, inout(reg) x =&gt; y, number = const 5);
}
assert_eq!(y, 8);
</code></pre>

<p>上面示例中，<code>inout(reg) x</code>语句表示编译器应该找到一个合适的通用寄存器，用<code>x</code>的当前值准备该寄存器，将加法指令的输出存储在同一个通用寄存器中，然后将该通用寄存器的值存储在<code>x</code>中。</p>

<p>新的 <code>asm!</code> 宏的写法更像 <code>println!</code> 宏，这样更加易读。而旧的写法，需要和具体的汇编语法相绑定，并不通用。</p>

<h3 id="rustdoc-提升">Rustdoc 提升</h3>

<p>Rust 是一门优雅的语言，并且这份优雅是非常完整的。除了语言的诸多特性设计优雅之外，还有一个亮点就是 Rustdoc，Rust 官方 doc 工作组立志让 Rustdoc 成为一个伟大的工具。</p>

<p>Rustdoc 使用简单，可以创建非常漂亮的页面，使编写文档成为一种乐趣。关于 Rustdoc 详细介绍你可以去看 <a href="https://doc.rust-lang.org/rustdoc/what-is-rustdoc.html" target="_blank">Rustdoc book</a> 。</p>

<p>Rustdoc 工作组最近在不断更新其功能，宗旨就是让编写文档更加轻松，消除重复的工作。比如，可以把项目的<code>README</code>文档，通过 <code>#[doc]</code> 属性来指派给某个模块，从而可以减少没必要的重复。</p>

<p>当然，未来的改进还有很多工作要做，这也算是 Rust 未来一大挑战。</p>

<h3 id="deref-pattern">deref pattern</h3>

<p><code>deref pattern</code> 是一个代表，它可以看作是Rust 官方对 Rust 语言诸多持续改进中的一个影子。</p>

<p>该特性简单来说，就是想让 Rust 语言在 <code>match</code> 模式匹配中也支持 <code>deref</code>：</p>

<pre><code>let x: Option&lt;Rc&lt;bool&gt;&gt; = ...;
match x {
    Some(deref true) =&gt; ...,
    Some(x) =&gt; ...,
    None =&gt; ...,
}
</code></pre>

<p>比如上面代码，匹配 <code>Option&lt;Rc&lt;bool&gt;&gt;</code> 的时候，可以无视其中的 <code>Rc</code>，直接透明操作 <code>bool</code>。上面例子里是一种解决方案，就是增加一个 <code>deref</code> 关键字。当然最终使用什么方案并未确定。</p>

<p>这里提到这个特性，是想说，Rust 语言目前在人体工程学方面，还有很多提升的空间；并且，Rust 团队也在不断的努力，让 Rust 语言使用起来更加方便和优雅。</p>

<h2 id="小结">小结</h2>

<p>Rust 语言自身相对已经成熟，生态也足够丰富，并且在一些应用领域崭露头角。</p>

<p>Rust在系统语言的地位上，更像是当年的 C 语言。同样是通用语言，Rust现在在操作系统、云原生、物联网等关键系统领域成为刚需。因为“安全”现在已经是必选项了，这是 Rust 语言的时代机遇。同时，Rust 语言也在不同领域造就了新的职业岗位。</p>

<p>我们也看到，Rust 语言还有很多需要完善的地方，但这些都在官方团队的计划之中。我相信，在 Rust 基金会的引领下，Rust 肯定会迈向广泛应用的美好未来。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#28444444111c1919181f684f45494144064b4745" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a57076fe37755',t:'MTczNDEzODg5Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
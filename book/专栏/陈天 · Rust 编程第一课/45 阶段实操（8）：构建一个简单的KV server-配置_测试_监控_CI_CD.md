<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=45&#32;阶段实操（8）：构建一个简单的KV&#32;server-配置_测试_监控_CI_CD>
        <link rel="icon" href="/static/favicon.png">
        <title>45 阶段实操（8）：构建一个简单的KV server-配置_测试_监控_CI_CD </title>
        
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
                            <h1 id="title" data-id="45 阶段实操（8）：构建一个简单的KV server-配置_测试_监控_CI_CD" class="title">45 阶段实操（8）：构建一个简单的KV server-配置_测试_监控_CI_CD</h1>
                            <div><p>你好，我是陈天。</p>

<p>终于来到了我们这个 KV server 系列的终章。其实原本 KV server 我只计划了 4 讲，但现在 8 讲似乎都还有些意犹未尽。虽然这是一个“简单”的 KV server，它没有复杂的性能优化 —— 我们只用了一句 unsafe；也没有复杂的生命周期处理 —— 只有零星 &lsquo;static 标注；更没有支持集群的处理。</p>

<p>然而，如果你能够理解到目前为止的代码，甚至能独立写出这样的代码，那么，你已经具备足够的、能在一线大厂开发的实力了，国内我不是特别清楚，但在北美这边，保守一些地说，300k+ USD 的 package 应该可以轻松拿到。</p>

<p>今天我们就给KV server项目收个尾，结合之前梳理的实战中 Rust 项目应该考虑的问题，来聊聊和生产环境有关的一些处理，按开发流程，主要讲五个方面：配置、集成测试、性能测试、测量和监控、CI/CD。</p>

<h2 id="配置">配置</h2>

<p>首先在 Cargo.toml 里添加 <a href="https://github.com/serde-rs/serde" target="_blank">serde</a> 和 <a href="https://github.com/alexcrichton/toml-rs" target="_blank">toml</a>。我们计划使用 toml 做配置文件，serde 用来处理配置的序列化和反序列化：</p>

<pre><code>[dependencies]
...
serde = { version = &quot;1&quot;, features = [&quot;derive&quot;] } # 序列化/反序列化
...
toml = &quot;0.5&quot; # toml 支持
...
</code></pre>

<p>然后来创建一个 src/config.rs，构建 KV server 的配置：</p>

<pre><code>use crate::KvError;
use serde::{Deserialize, Serialize};
use std::fs;

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct ServerConfig {
    pub general: GeneralConfig,
    pub storage: StorageConfig,
    pub tls: ServerTlsConfig,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct ClientConfig {
    pub general: GeneralConfig,
    pub tls: ClientTlsConfig,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct GeneralConfig {
    pub addr: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
#[serde(tag = &quot;type&quot;, content = &quot;args&quot;)]
pub enum StorageConfig {
    MemTable,
    SledDb(String),
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct ServerTlsConfig {
    pub cert: String,
    pub key: String,
    pub ca: Option&lt;String&gt;,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct ClientTlsConfig {
    pub domain: String,
    pub identity: Option&lt;(String, String)&gt;,
    pub ca: Option&lt;String&gt;,
}

impl ServerConfig {
    pub fn load(path: &amp;str) -&gt; Result&lt;Self, KvError&gt; {
        let config = fs::read_to_string(path)?;
        let config: Self = toml::from_str(&amp;config)?;
        Ok(config)
    }
}

impl ClientConfig {
    pub fn load(path: &amp;str) -&gt; Result&lt;Self, KvError&gt; {
        let config = fs::read_to_string(path)?;
        let config: Self = toml::from_str(&amp;config)?;
        Ok(config)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn server_config_should_be_loaded() {
        let result: Result&lt;ServerConfig, toml::de::Error&gt; =
            toml::from_str(include_str!(&quot;../fixtures/server.conf&quot;));
        assert!(result.is_ok());
    }

    #[test]
    fn client_config_should_be_loaded() {
        let result: Result&lt;ClientConfig, toml::de::Error&gt; =
            toml::from_str(include_str!(&quot;../fixtures/client.conf&quot;));
        assert!(result.is_ok());
    }
}
</code></pre>

<p>你可以看到，在 Rust 下，有了 serde 的帮助，处理任何已知格式的配置文件，是多么容易的一件事情。我们<strong>只需要定义数据结构，并为数据结构使用 Serialize/Deserialize 派生宏，就可以处理任何支持 serde 的数据结构</strong>。</p>

<p>我还写了个 examples/gen_config.rs（你可以自行去查阅它的代码），用来生成配置文件，下面是生成的服务端的配置：</p>

<pre><code>[general]
addr = '127.0.0.1:9527'

[storage]
type = 'SledDb'
args = '/tmp/kv_server'

[tls]
cert = &quot;&quot;&quot;
-----BEGIN CERTIFICATE-----\r
MIIBdzCCASmgAwIBAgIICpy02U2yuPowBQYDK2VwMDMxCzAJBgNVBAYMAkNOMRIw\r
EAYDVQQKDAlBY21lIEluYy4xEDAOBgNVBAMMB0FjbWUgQ0EwHhcNMjEwOTI2MDEy\r
NTU5WhcNMjYwOTI1MDEyNTU5WjA6MQswCQYDVQQGDAJDTjESMBAGA1UECgwJQWNt\r
ZSBJbmMuMRcwFQYDVQQDDA5BY21lIEtWIHNlcnZlcjAqMAUGAytlcAMhAK2Z2AjF\r
A0uiltNuCvl6EVFl6tpaS/wJYB5IdWT2IISdo1QwUjAcBgNVHREEFTATghFrdnNl\r
cnZlci5hY21lLmluYzATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMEBTADAQEA\r
MA8GA1UdDwEB/wQFAwMH4AAwBQYDK2VwA0EASGOmOWFPjbGhXNOmYNCa3lInbgRy\r
iTNtB/5kElnbKkhKhRU7yQ8HTHWWkyU5WGWbOOIXEtYp+5ERUJC+mzP9Bw==\r
-----END CERTIFICATE-----\r
&quot;&quot;&quot;
key = &quot;&quot;&quot;
-----BEGIN PRIVATE KEY-----\r
MFMCAQEwBQYDK2VwBCIEIPMyINaewhXwuTPUufFO2mMt/MvQMHrGDGxgdgfy/kUu\r
oSMDIQCtmdgIxQNLopbTbgr5ehFRZeraWkv8CWAeSHVk9iCEnQ==\r
-----END PRIVATE KEY-----\r
&quot;&quot;&quot;
</code></pre>

<p>有了配置文件的支持，就可以在 <a href="http://lib.rs" target="_blank">lib.rs</a> 下写一些辅助函数，让我们创建服务端和客户端更加简单：</p>

<pre><code>mod config;
mod error;
mod network;
mod pb;
mod service;
mod storage;

pub use config::*;
pub use error::KvError;
pub use network::*;
pub use pb::abi::*;
pub use service::*;
pub use storage::*;

use anyhow::Result;
use tokio::net::{TcpListener, TcpStream};
use tokio_rustls::client;
use tokio_util::compat::FuturesAsyncReadCompatExt;
use tracing::info;

/// 通过配置创建 KV 服务器
pub async fn start_server_with_config(config: &amp;ServerConfig) -&gt; Result&lt;()&gt; {
    let acceptor =
        TlsServerAcceptor::new(&amp;config.tls.cert, &amp;config.tls.key, config.tls.ca.as_deref())?;

    let addr = &amp;config.general.addr;
    match &amp;config.storage {
        StorageConfig::MemTable =&gt; start_tls_server(addr, MemTable::new(), acceptor).await?,
        StorageConfig::SledDb(path) =&gt; start_tls_server(addr, SledDb::new(path), acceptor).await?,
    };

    Ok(())
}

/// 通过配置创建 KV 客户端
pub async fn start_client_with_config(
    config: &amp;ClientConfig,
) -&gt; Result&lt;YamuxCtrl&lt;client::TlsStream&lt;TcpStream&gt;&gt;&gt; {
    let addr = &amp;config.general.addr;
    let tls = &amp;config.tls;

    let identity = tls.identity.as_ref().map(|(c, k)| (c.as_str(), k.as_str()));
    let connector = TlsClientConnector::new(&amp;tls.domain, identity, tls.ca.as_deref())?;
    let stream = TcpStream::connect(addr).await?;
    let stream = connector.connect(stream).await?;

    // 打开一个 stream
    Ok(YamuxCtrl::new_client(stream, None))
}

async fn start_tls_server&lt;Store: Storage&gt;(
    addr: &amp;str,
    store: Store,
    acceptor: TlsServerAcceptor,
) -&gt; Result&lt;()&gt; {
    let service: Service&lt;Store&gt; = ServiceInner::new(store).into();
    let listener = TcpListener::bind(addr).await?;
    info!(&quot;Start listening on {}&quot;, addr);
    loop {
        let tls = acceptor.clone();
        let (stream, addr) = listener.accept().await?;
        info!(&quot;Client {:?} connected&quot;, addr);

        let svc = service.clone();
        tokio::spawn(async move {
            let stream = tls.accept(stream).await.unwrap();
            YamuxCtrl::new_server(stream, None, move |stream| {
                let svc1 = svc.clone();
                async move {
                    let stream = ProstServerStream::new(stream.compat(), svc1.clone());
                    stream.process().await.unwrap();
                    Ok(())
                }
            });
        });
    }
}
</code></pre>

<p>有了 start_server_with_config 和 start_client_with_config 这两个辅助函数，我们就可以简化 src/server.rs 和 src/client.rs 了。下面是 src/server.rs 的新代码：</p>

<pre><code>use anyhow::Result;
use kv6::{start_server_with_config, ServerConfig};

#[tokio::main]
async fn main() -&gt; Result&lt;()&gt; {
    tracing_subscriber::fmt::init();
    let config: ServerConfig = toml::from_str(include_str!(&quot;../fixtures/server.conf&quot;))?;

    start_server_with_config(&amp;config).await?;

    Ok(())
}
</code></pre>

<p>可以看到，整个代码简洁了很多。在这个重构的过程中，还有一些其它改动，你可以看 GitHub repo 下 45 讲的 diff_config。</p>

<h2 id="集成测试">集成测试</h2>

<p>之前我们写了很多单元测试，但还没有写过一行集成测试。今天就来写一个简单的集成测试，确保客户端和服务器完整的交互工作正常。</p>

<p>之前提到在 Rust 里，集成测试放在 tests 目录下，每个测试编成单独的二进制。所以首先，我们创建和 src 平行的 tests 目录。然后再创建 tests/server.rs，填入以下代码：</p>

<pre><code>use anyhow::Result;
use kv6::{
    start_client_with_config, start_server_with_config, ClientConfig, CommandRequest,
    ProstClientStream, ServerConfig, StorageConfig,
};
use std::time::Duration;
use tokio::time;

#[tokio::test]
async fn yamux_server_client_full_tests() -&gt; Result&lt;()&gt; {
    let addr = &quot;127.0.0.1:10086&quot;;

    let mut config: ServerConfig = toml::from_str(include_str!(&quot;../fixtures/server.conf&quot;))?;
    config.general.addr = addr.into();
    config.storage = StorageConfig::MemTable;

    // 启动服务器
    tokio::spawn(async move {
        start_server_with_config(&amp;config).await.unwrap();
    });

    time::sleep(Duration::from_millis(10)).await;
    let mut config: ClientConfig = toml::from_str(include_str!(&quot;../fixtures/client.conf&quot;))?;
    config.general.addr = addr.into();

    let mut ctrl = start_client_with_config(&amp;config).await.unwrap();
    let stream = ctrl.open_stream().await?;
    let mut client = ProstClientStream::new(stream);

    // 生成一个 HSET 命令
    let cmd = CommandRequest::new_hset(&quot;table1&quot;, &quot;hello&quot;, &quot;world&quot;.to_string().into());
    client.execute_unary(&amp;cmd).await?;

    // 生成一个 HGET 命令
    let cmd = CommandRequest::new_hget(&quot;table1&quot;, &quot;hello&quot;);
    let data = client.execute_unary(&amp;cmd).await?;

    assert_eq!(data.status, 200);
    assert_eq!(data.values, &amp;[&quot;world&quot;.into()]);

    Ok(())
}
</code></pre>

<p>可以看到，<strong>集成测试的写法和单元测试其实很类似，只不过我们不需要再使用 #[cfg(test)] 来做条件编译</strong>。</p>

<p>如果你的集成测试比较复杂，需要比较多的辅助代码，那么你还可以在 tests 下 cargo new 出一个项目，然后在那个项目里撰写辅助代码和测试代码。如果你对此感兴趣，可以看 <a href="https://github.com/hyperium/tonic/tree/master/tests" target="_blank">tonic 的集成测试</a>。不过注意了，集成测试和你的 crate 用同样的条件编译，所以在集成测试里，无法使用单元测试中构建的辅助代码。</p>

<h2 id="性能测试">性能测试</h2>

<p>在之前不断完善 KV server 的过程中，你一定会好奇：我们的 KV server 性能究竟如何呢？那来写一个关于 Pub/Sub 的性能测试吧。</p>

<p>基本的想法是我们连上 100 个 subscriber 作为背景，然后看 publisher publish 的速度。</p>

<p>因为 BROADCAST_CAPACITY 有限，是 128，当 publisher 速度太快，而导致 server 不能及时往 subscriber 发送时，server 接收 client 数据的速度就会降下来，无法接收新的 client，整体的 publish 的速度也会降下来，所以这个测试能够了解 server 处理 publish 的速度。</p>

<p>为了确认这一点，我们在 start_tls_server() 函数中，在 process() 之前，再加个 100ms 的延时，人为减缓系统的处理速度：</p>

<pre><code>async move {
    let stream = ProstServerStream::new(stream.compat(), svc1.clone());
    // 延迟 100ms 处理
    time::sleep(Duration::from_millis(100)).await;
    stream.process().await.unwrap();
    Ok(())
}
</code></pre>

<p>好，现在可以写性能测试了。</p>

<p>在 Rust 下，我们可以用 <a href="https://github.com/bheisler/criterion.rs" target="_blank">criterion</a> 库。它可以处理基本的性能测试，并生成漂亮的报告。所以在 Cargo.toml 中加入：</p>

<pre><code>[dev-dependencies]
...
criterion = { version = &quot;0.3&quot;, features = [&quot;async_futures&quot;, &quot;async_tokio&quot;, &quot;html_reports&quot;] } # benchmark
...
rand = &quot;0.8&quot; # 随机数处理
...

[[bench]]
name = &quot;pubsub&quot;
harness = false
</code></pre>

<p>最后这个 bench section，描述了性能测试的名字，它对应 benches 目录下的同名文件。</p>

<p>我们创建和 src 平级的 benches，然后再创建 benches/pubsub.rs，添入如下代码：</p>

<pre><code>use anyhow::Result;
use criterion::{criterion_group, criterion_main, Criterion};
use futures::StreamExt;
use kv6::{
    start_client_with_config, start_server_with_config, ClientConfig, CommandRequest, ServerConfig,
    StorageConfig, YamuxCtrl,
};
use rand::prelude::SliceRandom;
use std::time::Duration;
use tokio::net::TcpStream;
use tokio::runtime::Builder;
use tokio::time;
use tokio_rustls::client::TlsStream;
use tracing::info;

async fn start_server() -&gt; Result&lt;()&gt; {
    let addr = &quot;127.0.0.1:9999&quot;;
    let mut config: ServerConfig = toml::from_str(include_str!(&quot;../fixtures/server.conf&quot;))?;
    config.general.addr = addr.into();
    config.storage = StorageConfig::MemTable;

    tokio::spawn(async move {
        start_server_with_config(&amp;config).await.unwrap();
    });

    Ok(())
}

async fn connect() -&gt; Result&lt;YamuxCtrl&lt;TlsStream&lt;TcpStream&gt;&gt;&gt; {
    let addr = &quot;127.0.0.1:9999&quot;;
    let mut config: ClientConfig = toml::from_str(include_str!(&quot;../fixtures/client.conf&quot;))?;
    config.general.addr = addr.into();

    Ok(start_client_with_config(&amp;config).await?)
}

async fn start_subscribers(topic: &amp;'static str) -&gt; Result&lt;()&gt; {
    let mut ctrl = connect().await?;
    let stream = ctrl.open_stream().await?;
    info!(&quot;C(subscriber): stream opened&quot;);
    let cmd = CommandRequest::new_subscribe(topic.to_string());
    tokio::spawn(async move {
        let mut stream = stream.execute_streaming(&amp;cmd).await.unwrap();
        while let Some(Ok(data)) = stream.next().await {
            drop(data);
        }
    });

    Ok(())
}

async fn start_publishers(topic: &amp;'static str, values: &amp;'static [&amp;'static str]) -&gt; Result&lt;()&gt; {
    let mut rng = rand::thread_rng();
    let v = values.choose(&amp;mut rng).unwrap();

    let mut ctrl = connect().await.unwrap();
    let mut stream = ctrl.open_stream().await.unwrap();
    info!(&quot;C(publisher): stream opened&quot;);

    let cmd = CommandRequest::new_publish(topic.to_string(), vec![(*v).into()]);
    stream.execute_unary(&amp;cmd).await.unwrap();

    Ok(())
}

fn pubsub(c: &amp;mut Criterion) {
    // tracing_subscriber::fmt::init();
    // 创建 Tokio runtime
    let runtime = Builder::new_multi_thread()
        .worker_threads(4)
        .thread_name(&quot;pubsub&quot;)
        .enable_all()
        .build()
        .unwrap();
    let values = &amp;[&quot;Hello&quot;, &quot;Tyr&quot;, &quot;Goodbye&quot;, &quot;World&quot;];
    let topic = &quot;lobby&quot;;

    // 运行服务器和 100 个 subscriber，为测试准备
    runtime.block_on(async {
        eprint!(&quot;preparing server and subscribers&quot;);
        start_server().await.unwrap();
        time::sleep(Duration::from_millis(50)).await;
        for _ in 0..100 {
            start_subscribers(topic).await.unwrap();
            eprint!(&quot;.&quot;);
        }
        eprintln!(&quot;Done!&quot;);
    });

    // 进行 benchmark
    c.bench_function(&quot;publishing&quot;, move |b| {
        b.to_async(&amp;runtime)
            .iter(|| async { start_publishers(topic, values).await })
    });
}

criterion_group! {
    name = benches;
    config = Criterion::default().sample_size(10);
    targets = pubsub
}
criterion_main!(benches);
</code></pre>

<p>大部分的代码都很好理解，就是创建服务器和客户端，为测试做准备。说一下这里面核心的 benchmark 代码：</p>

<pre><code>c.bench_function(&quot;publishing&quot;, move |b| {
    b.to_async(&amp;runtime)
        .iter(|| async { start_publishers(topic, values).await })
});
</code></pre>

<p>对于要测试的代码，我们可以封装成一个函数进行测试。<strong>这里因为要做 async 函数的测试，需要使用 runtime。普通的函数不需要调用 to_async</strong>。对于更多有关 criterion 的用法，可以参考它的文档。</p>

<p>运行 <code>cargo bench</code> 后，会见到如下打印（如果你的代码无法通过，可以参考 repo 里的 diff_benchmark，我顺便做了一点小重构）：</p>

<pre><code>preparing server and subscribers....................................................................................................Done!
publishing              time:   [419.73 ms 426.84 ms 434.20 ms]                     
                        change: [-1.6712% +1.0499% +3.6586%] (p = 0.48 &gt; 0.05)
                        No change in performance detected.
</code></pre>

<p>可以看到，单个 publish 的处理速度要 426ms，好慢！我们把之前在 start_tls_server() 里加的延迟去掉，再次测试：</p>

<pre><code>preparing server and subscribers....................................................................................................Done!
publishing              time:   [318.61 ms 324.48 ms 329.81 ms]                     
                        change: [-25.854% -23.980% -22.144%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
</code></pre>

<p>嗯，这下 324ms，正好是减去刚才加的 100ms。可是这个速度依旧不合理，凭直觉我们感觉一下这个速度，是 Python 这样的语言还正常，如果是 Rust 也太慢了吧？</p>

<h2 id="测量和监控">测量和监控</h2>

<p>工业界有句名言：如果你无法测量，那你就无法改进（If you can’t measure it, you can’t improve it）。现在知道了 KV server 性能有问题，但并不知道问题出在哪里。我们需要使用合适的测量方式。</p>

<p>目前，<strong>比较好的端对端的性能监控和测量工具是 jaeger</strong>，我们可以在 KV server/client 侧收集监控信息，发送给 jaeger 来查看在服务器和客户端的整个处理流程中，时间都花费到哪里去了。</p>

<p>之前我们在 KV server 里使用的日志工具是 tracing，不过日志只是它的诸多功能之一，它还能做 <a href="https://docs.rs/tracing/0.1.28/tracing/attr.instrument.html" target="_blank">instrument</a>，然后配合 <a href="https://github.com/open-telemetry/opentelemetry-rust" target="_blank">opentelemetry</a> 库，我们就可以把 instrument 的结果发送给 jaeger 了。</p>

<p>好，在 Cargo.toml 里添加新的依赖：</p>

<pre><code>[dependencies]
...
opentelemetry-jaeger = &quot;0.15&quot; # opentelemetry jaeger 支持
...
tracing-appender = &quot;0.1&quot; # 文件日志
tracing-opentelemetry = &quot;0.15&quot; # opentelemetry 支持
tracing-subscriber = { version = &quot;0.2&quot;, features = [&quot;json&quot;, &quot;chrono&quot;] } # 日志处理
</code></pre>

<p>有了这些依赖后，在 benches/pubsub.rs 里，我们可以在初始化 tracing_subscriber 时，使用 jaeger 和 opentelemetry tracer：</p>

<pre><code>fn pubsub(c: &amp;mut Criterion) {
    let tracer = opentelemetry_jaeger::new_pipeline()
        .with_service_name(&quot;kv-bench&quot;)
        .install_simple()
        .unwrap();
    let opentelemetry = tracing_opentelemetry::layer().with_tracer(tracer);

    tracing_subscriber::registry()
        .with(EnvFilter::from_default_env())
        .with(opentelemetry)
        .init();

    let root = span!(tracing::Level::INFO, &quot;app_start&quot;, work_units = 2);
    let _enter = root.enter();
    // 创建 Tokio runtime
		...
}
</code></pre>

<p>设置好 tracing 后，就在系统的主流程上添加相应的 instrument：-
<img src="assets/f1680244d5c7901ec26181c01bfea8a1.jpg" alt="" /></p>

<p>新添加的代码你可以看 repo 中的 diff_telemetry。注意 instrument 可以用不同的名称，比如，对于 TlsConnector::new() 函数，可以用 <code>#[instrument(name = &quot;tls_connector_new&quot;)]</code>，这样它的名字辨识度高一些。</p>

<p>为主流程中的函数添加完 instrument 后，你需要先打开一个窗口，运行 jaeger（需要 docker）：</p>

<pre><code>docker run -d -p6831:6831/udp -p6832:6832/udp -p16686:16686 -p14268:14268 jaegertracing/all-in-one:latest
</code></pre>

<p>然后带着 RUST_LOG=info 运行 benchmark：</p>

<pre><code>RUST_LOG=info cargo bench
</code></pre>

<p>由于我的 OS X 上没装 docker（docker 不支持 Mac，需要 Linux VM 中转），我就在一个 Ubuntu 虚拟机里运行这两条命令：</p>

<pre><code>preparing server and subscribers....................................................................................................Done!
publishing              time:   [1.7464 ms 1.9556 ms 2.2343 ms]                       
Found 2 outliers among 10 measurements (20.00%)
  1 (10.00%) high mild
  1 (10.00%) high severe
</code></pre>

<p>并没有做任何事情，似乎只是换了个系统，性能就提升了很多，这给我们一个 tip：也许问题出在 OS X 和 Linux 系统相关的部分。</p>

<p>不管怎样，已经发送了不少数据给 jaeger，我们到 jaeger 上看看问题出在哪里。</p>

<p>打开 <a href="http://localhost:16686/" target="_blank">http://localhost:16686/</a>，service 选 kv-bench，Operation 选 app_start，点击 “Find Traces”，我们可以看到捕获的 trace。因为运行了两次 benchmark，所以有两个 app_start 的查询结果：-
<img src="assets/ecd9b1d06debe7fb3fe507befd803877.png" alt="" /></p>

<p>可以看到，每次 start_client_with_config 都要花 1.6-2.5ms，其中有差不多一小半时间花在了 TlsClientConnector::new() 上：-
<img src="assets/fe574ccac09ce5434027fce2afebaeb6.png" alt="" /></p>

<p>如果说 TlsClientConnector::connect() 花不少时间还情有可原，因为这是整个 TLS 协议的握手过程，涉及到网络调用、包的加解密等。<strong>但 TlsClientConnector::new() 就是加载一些证书、创建 TlsConnector 这个数据结构而已，为何这么慢</strong>？</p>

<p>仔细阅读 TlsClientConnector::new() 的代码，你可以对照注释看：</p>

<pre><code>#[instrument(name = &quot;tls_connector_new&quot;, skip_all)]
pub fn new(
    domain: impl Into&lt;String&gt; + std::fmt::Debug,
    identity: Option&lt;(&amp;str, &amp;str)&gt;,
    server_ca: Option&lt;&amp;str&gt;,
) -&gt; Result&lt;Self, KvError&gt; {
    let mut config = ClientConfig::new();

    // 如果有客户端证书，加载之
    if let Some((cert, key)) = identity {
        let certs = load_certs(cert)?;
        let key = load_key(key)?;
        config.set_single_client_cert(certs, key)?;
    }

    // 加载本地信任的根证书链
    config.root_store = match rustls_native_certs::load_native_certs() {
        Ok(store) | Err((Some(store), _)) =&gt; store,
        Err((None, error)) =&gt; return Err(error.into()),
    };

    // 如果有签署服务器的 CA 证书，则加载它，这样服务器证书不在根证书链
    // 但是这个 CA 证书能验证它，也可以
    if let Some(cert) = server_ca {
        let mut buf = Cursor::new(cert);
        config.root_store.add_pem_file(&amp;mut buf).unwrap();
    }

    Ok(Self {
        config: Arc::new(config),
        domain: Arc::new(domain.into()),
    })
}
</code></pre>

<p>可以发现，它的代码唯一可能影响性能的就是加载本地信任的根证书链的部分。这个代码会和操作系统交互，获取信任的根证书链。也许，这就是影响性能的原因之一？</p>

<p><strong>那我们将其简单重构一下</strong>。因为根证书链，只有在客户端没有提供用于验证服务器证书的 CA 证书时，才需要，所以可以在没有 CA 证书时，才加载本地的根证书链：</p>

<pre><code>#[instrument(name = &quot;tls_connector_new&quot;, skip_all)]
pub fn new(
    domain: impl Into&lt;String&gt; + std::fmt::Debug,
    identity: Option&lt;(&amp;str, &amp;str)&gt;,
    server_ca: Option&lt;&amp;str&gt;,
) -&gt; Result&lt;Self, KvError&gt; {
    let mut config = ClientConfig::new();

    // 如果有客户端证书，加载之
    if let Some((cert, key)) = identity {
        let certs = load_certs(cert)?;
        let key = load_key(key)?;
        config.set_single_client_cert(certs, key)?;
    }

    // 如果有签署服务器的 CA 证书，则加载它，这样服务器证书不在根证书链
    // 但是这个 CA 证书能验证它，也可以
    if let Some(cert) = server_ca {
        let mut buf = Cursor::new(cert);
        config.root_store.add_pem_file(&amp;mut buf).unwrap();
    } else {
        // 加载本地信任的根证书链
        config.root_store = match rustls_native_certs::load_native_certs() {
            Ok(store) | Err((Some(store), _)) =&gt; store,
            Err((None, error)) =&gt; return Err(error.into()),
        };
    }

    Ok(Self {
        config: Arc::new(config),
        domain: Arc::new(domain.into()),
    })
}
</code></pre>

<p>完成这个修改后，我们再运行 <code>RUST_LOG=info cargo bench</code>，现在的性能达到了 1.64ms，相比之前的 1.95ms，提升了 16%。</p>

<p>打开 jaeger，看最新的 app_start 结果，发现 TlsClientConnector::new() 所花时间降到了 ~12us 左右。嗯，虽然没有抓到服务器本身的 bug，但客户端的 bug 倒是解决了一个。-
<img src="assets/3cfde740dbe0d4a897e2d4c3684b530b.png" alt="" /></p>

<p>至于服务器，如果我们看 Service::execute 的主流程，执行速度在 40-60us，问题不大：-
<img src="assets/7be6139668c82fb8b79fb66f3ed06d31.png" alt="" /></p>

<p>再看服务器的主流程 server_process：-
<img src="assets/076402ac25b507295d022b980378e363.png" alt="" /></p>

<p>这是我们在 start_tls_server() 里额外添加的 tracing span：</p>

<pre><code>loop {
		let root = span!(tracing::Level::INFO, &quot;server_process&quot;);
		let _enter = root.enter();
		...
}
</code></pre>

<p>把右上角的 trace timeline 改成 trace graph，然后点右侧的 time：-
<img src="assets/1499657924a241e43c9d1be467793041.png" alt="" /></p>

<p>可以看到，主要的服务器时间都花在了 TLS accept 上，所以，<strong>目前服务器没有太多值得优化的地方</strong>。</p>

<p>由于 tracing 本身也占用不少 CPU，所以我们直接 <code>cargo bench</code> 看看目前的结果：</p>

<pre><code>preparing server and subscribers....................................................................................................Done!
publishing              time:   [1.3986 ms 1.4140 ms 1.4474 ms]                       
                        change: [-26.647% -19.977% -10.798%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
Found 2 outliers among 10 measurements (20.00%)
  2 (20.00%) high severe
</code></pre>

<p>不加 RUST_LOG=info 后，整体性能到了 1.4ms。这是我在 Ubuntu 虚拟机下的结果。</p>

<p>我们再回到 OS X 下测试，看看 TlsClientConnector::new() 的修改，对OS X 是否有效：</p>

<pre><code>preparing server and subscribers....................................................................................................Done!
publishing              time:   [1.4086 ms 1.4229 ms 1.4315 ms]                       
                        change: [-99.570% -99.563% -99.554%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
</code></pre>

<p>嗯，在我的 OS X下，现在整体性能也到了 1.4ms 的水平。这也意味着，在有 100 个 subscribers 的情况下，我们的 KV server 每秒钟可以处理 714k publish 请求；而在 1000 个 subscribers 的情况下，性能在 11.1ms 的水平，也就是每秒可以处理 90k publish 请求：</p>

<pre><code>publishing              time:   [11.007 ms 11.095 ms 11.253 ms]                      
                        change: [-96.618% -96.556% -96.486%] (p = 0.00 &lt; 0.05)
                        Performance has improved.
</code></pre>

<p>你也许会觉得目前 publish 的 value 太小，那换一些更加贴近实际的字符串大小：</p>

<pre><code>// let values = &amp;[&quot;Hello&quot;, &quot;Tyr&quot;, &quot;Goodbye&quot;, &quot;World&quot;];
let base_str = include_str!(&quot;../fixtures/server.conf&quot;); // 891 bytes

let values: &amp;'static [&amp;'static str] = Box::leak(
    vec![
        &amp;base_str[..64],
        &amp;base_str[..128],
        &amp;base_str[..256],
        &amp;base_str[..512],
    ]
    .into_boxed_slice(),
);
</code></pre>

<p>测试结果差不太多：</p>

<pre><code>publishing              time:   [10.917 ms 11.098 ms 11.428 ms]                      
                        change: [-0.4822% +2.3311% +4.9631%] (p = 0.12 &gt; 0.05)
                        No change in performance detected.
</code></pre>

<p>criterion 还会生成漂亮的 report，你可以用浏览器打开 ./target/criterion/publishing/report/index.html 查看（名字是publishing ，因为 benchmark ID 是 publishing）：-
<img src="assets/d3cebd8e3c164171febbe34e43916885.png" alt="" /></p>

<p>好，处理完性能相关的问题，我们来<strong>为 server 添加日志和性能监测的支持</strong>：</p>

<pre><code>use std::env;

use anyhow::Result;
use kv6::{start_server_with_config, RotationConfig, ServerConfig};
use tokio::fs;
use tracing::span;
use tracing_subscriber::{
    fmt::{self, format},
    layer::SubscriberExt,
    prelude::*,
    EnvFilter,
};

#[tokio::main]
async fn main() -&gt; Result&lt;()&gt; {
    // 如果有环境变量，使用环境变量中的 config
    let config = match env::var(&quot;KV_SERVER_CONFIG&quot;) {
        Ok(path) =&gt; fs::read_to_string(&amp;path).await?,
        Err(_) =&gt; include_str!(&quot;../fixtures/server.conf&quot;).to_string(),
    };
    let config: ServerConfig = toml::from_str(&amp;config)?;

    let tracer = opentelemetry_jaeger::new_pipeline()
        .with_service_name(&quot;kv-server&quot;)
        .install_simple()?;
    let opentelemetry = tracing_opentelemetry::layer().with_tracer(tracer);

    // 添加
    let log = &amp;config.log;
    let file_appender = match log.rotation {
        RotationConfig::Hourly =&gt; tracing_appender::rolling::hourly(&amp;log.path, &quot;server.log&quot;),
        RotationConfig::Daily =&gt; tracing_appender::rolling::daily(&amp;log.path, &quot;server.log&quot;),
        RotationConfig::Never =&gt; tracing_appender::rolling::never(&amp;log.path, &quot;server.log&quot;),
    };

    let (non_blocking, _guard1) = tracing_appender::non_blocking(file_appender);
    let fmt_layer = fmt::layer()
        .event_format(format().compact())
        .with_writer(non_blocking);

    tracing_subscriber::registry()
        .with(EnvFilter::from_default_env())
        .with(fmt_layer)
        .with(opentelemetry)
        .init();

    let root = span!(tracing::Level::INFO, &quot;app_start&quot;, work_units = 2);
    let _enter = root.enter();

    start_server_with_config(&amp;config).await?;

    Ok(())
}
</code></pre>

<p>为了让日志能在配置文件中配置，需要更新一下 src/config.rs：</p>

<pre><code>#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct ServerConfig {
    pub general: GeneralConfig,
    pub storage: StorageConfig,
    pub tls: ServerTlsConfig,
    pub log: LogConfig,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub struct LogConfig {
    pub path: String,
    pub rotation: RotationConfig,
}

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq)]
pub enum RotationConfig {
    Hourly,
    Daily,
    Never,
}
</code></pre>

<p>你还需要更新 examples/gen_config.rs。相关的改变可以看 repo 下的 diff_logging。-
tracing 和 opentelemetry 还支持 <a href="https://github.com/prometheus/prometheus" target="_blank">prometheus</a>，你可以使用 <a href="https://docs.rs/opentelemetry-prometheus" target="_blank">opentelemetry-prometheus</a> 来和 prometheus 交互，如果有兴趣，你可以自己深入研究一下。</p>

<h2 id="ci-cd">CI/CD</h2>

<p>为了讲述方便，我把 CI/CD 放在最后，但 CI/CD 应该是在一开始的时候就妥善设置的。</p>

<p>先说CI吧。这个课程的 repo <a href="https://github.com/tyrchen/geektime-rust" target="_blank">tyrchen/geektime-rust</a> 在一开始就设置了 github action，每次 commit 都会运行：</p>

<ul>
<li>代码格式检查：cargo fmt</li>
<li>依赖 license 检查：cargo deny</li>
<li>linting：cargo check 和 cargo clippy</li>
<li>单元测试和集成测试：cargo test</li>
<li>生成文档：cargo doc</li>
</ul>

<p>github action 配置如下，供你参考：</p>

<pre><code>name: build

on:
  push:
    branches:
      - master
  pull_request:
    branches:
      - master

jobs:
  build-rust:
    strategy:
      matrix:
        platform: [ubuntu-latest, windows-latest]
    runs-on: ${{ matrix.platform }}
    steps:
      - uses: actions/checkout@v2
      - name: Cache cargo registry
        uses: actions/cache@v1
        with:
          path: ~/.cargo/registry
          key: ${{ runner.os }}-cargo-registry
      - name: Cache cargo index
        uses: actions/cache@v1
        with:
          path: ~/.cargo/git
          key: ${{ runner.os }}-cargo-index
      - name: Cache cargo build
        uses: actions/cache@v1
        with:
          path: target
          key: ${{ runner.os }}-cargo-build-target
      - name: Install stable
        uses: actions-rs/toolchain@v1
        with:
          profile: minimal
          toolchain: stable
          override: true
      - name: Check code format
        run: cargo fmt -- --check
      - name: Check the package for errors
        run: cargo check --all
      - name: Lint rust sources
        run: cargo clippy --all-targets --all-features --tests --benches -- -D warnings
      - name: Run tests
        run: cargo test --all-features -- --test-threads=1 --nocapture
      - name: Generate docs
        run: cargo doc --all-features --no-deps
      - name: Deploy docs to gh-page
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./target/doc
</code></pre>

<p>除此之外，我们还可以在每次 push tag 时做 release：</p>

<pre><code>name: release

on:
  push:
    tags:
      - &quot;v*&quot; # Push events to matching v*, i.e. v1.0, v20.15.10

jobs:
  build:
    name: Upload Release Asset
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest]
    steps:
      - name: Cache cargo registry
        uses: actions/cache@v1
        with:
          path: ~/.cargo/registry
          key: ${{ runner.os }}-cargo-registry
      - name: Cache cargo index
        uses: actions/cache@v1
        with:
          path: ~/.cargo/git
          key: ${{ runner.os }}-cargo-index
      - name: Cache cargo build
        uses: actions/cache@v1
        with:
          path: target
          key: ${{ runner.os }}-cargo-build-target
      - name: Checkout code
        uses: actions/checkout@v2
        with:
          token: ${{ secrets.GH_TOKEN }}
          submodules: recursive
      - name: Build project
        run: |
          make build-release
      - name: Create Release
        id: create_release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ github.ref }}
          release_name: Release ${{ github.ref }}
          draft: false
          prerelease: false
      - name: Upload asset
        id: upload-kv-asset
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          upload_url: ${{ steps.create_release.outputs.upload_url }}
          asset_path: ./target/release/kvs
          asset_name: kvs
          asset_content_type: application/octet-stream
      - name: Set env
        run: echo &quot;RELEASE_VERSION=${GITHUB_REF#refs/*/}&quot; &gt;&gt; $GITHUB_ENV
      - name: Deploy docs to gh-page
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./target/doc/simple_kv
          destination_dir: ${{ env.RELEASE_VERSION }}
</code></pre>

<p>这样，每次 push tag 时，都可以打包出来 Linux 的 kvs 版本：-
<img src="assets/1c61b7f58dd176bd25a565577d75af19.png" alt="" /></p>

<p>如果你不希望直接使用编译出来的二进制，也可以打包成 docker，在 Kubernetes 下使用。</p>

<p><strong>在做 CI 的过程中，我们也可以触发 CD</strong>，比如：</p>

<ul>
<li>PR merge 到 master，在 build 完成后，触发 dev 服务器的部署，团队内部可以尝试；</li>
<li>如果 release tag 包含 alpha，在 build 完成后，触发 staging 服务器的部署，公司内部可以使用；</li>
<li>如果 release tag 包含 beta，在 build 完成后，触发 beta 服务器的部署，beta 用户可以使用；</li>
<li>正式的 release tag 会触发生产环境的滚动升级，升级覆盖到的用户可以使用。</li>
</ul>

<p>一般来说，每家企业都有自己的 CI/CD 的工具链，这里为了展示方便，我们演示了如何使用 github action 对 Rust 代码做 CI，你可以按照自己的需要来处理。</p>

<p>在刚才的 action 代码中，还编译并上传了文档，所以我们可以通过 github pages 很方便地访问文档：-
<img src="assets/885d092273f8cacda1a65867a2489ea7.png" alt="" /></p>

<h2 id="小结">小结</h2>

<p>我们的 KV server 之旅就到此为止了。在整整 7 堂课里，我们一点点从零构造了一个完整的 KV server，包括注释在内，撰写了近三千行代码：</p>

<pre><code>❯ tokei .
-------------------------------------------------------------------------------
 Language            Files        Lines         Code     Comments       Blanks
-------------------------------------------------------------------------------
 Makefile                1           24           16            1            7
 Markdown                1            7            7            0            0
 Protocol Buffers        1          119           79           23           17
 Rust                   25         3366         2730          145          491
 TOML                    2          268          107          142           19
-------------------------------------------------------------------------------
 Total                  30         3784         2939          311          534
-------------------------------------------------------------------------------
</code></pre>

<p>这是一个非常了不起的成就！我们应该为自己感到自豪！</p>

<p>在这个系列里，我们大量使用 trait 和泛型，构建了很多复杂的数据结构；还为自己的类型实现了 AsyncRead/AsyncWrite/Stream/Sink 这些比较高阶的 trait。通过良好的设计，我们把网络层和业务层划分地非常清晰，网络层的变化不会影响到业务层，反之亦然：-
<img src="assets/53f5e5cf68b4300c3231885b10c784f3.jpeg" alt="" /></p>

<p>我们还模拟了比较真实的开发场景，通过大的需求变更，引发了一次不小的代码重构。</p>

<p>最终，通过性能测试，发现了一个客户端实现的小 bug。在处理这个 bug 的时候，我们欣喜地看到，Rust 有着非常强大的测试工具链，除了我们使用的单元测试、集成测试、性能测试，Rust 还支持模糊测试（fuzzy testing）和基于特性的测试（property testing）。</p>

<p>对于测试过程中发现的问题，Rust 有着非常完善的 tracing 工具链，可以和整个 opentelemetry 生态系统（包括 jaeger、prometheus 等工具）打通。我们就是通过使用 jaeger 找到并解决了问题。除此之外，Rust tracing 工具链还支持生成 <a href="https://github.com/tokio-rs/tracing/tree/master/tracing-flame" target="_blank">flamegraph</a>，篇幅关系，没有演示，你感兴趣的话可以试试。</p>

<p>最后，我们完善了 KV server 的配置、日志以及 CI。完整的代码我放在了 <a href="https://github.com/tyrchen/simple-kv" target="_blank">github.com/tyrchen/simple-kv</a> 上，欢迎查看最终的版本。</p>

<p>希望通过这个系列，你对如何使用 Rust 的特性来构造应用程序有了深度的认识。我相信，如果你能够跟得上这个系列的节奏，另外如果遇到新的库，用[第 20 讲]阅读代码的方式快速掌握，那么，大部分 Rust 开发中的挑战，对你而言都不是难事。</p>

<h3 id="思考题">思考题</h3>

<p>我们目前并未对日志做任何配置。一般来说，怎么做日志，会有相应的开关以及日志级别，如果希望能通过如下的配置记录日志，该怎么做？试试看：</p>

<pre><code>[log]
enable_log_file = true
enable_jaeger = false
log_level = 'info'
path = '/tmp/kv-log'
rotation = 'Daily'
</code></pre>

<p>欢迎在留言区分享自己做 KV server 系列的想法和感悟。你已经完成了第45次打卡，我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#3d51515104090c0c0d0a7d5a505c5451135e5250" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a54873c657755',t:'MTczNDEzODc5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=07&#32;构建模式：Go&#32;Module的6类常规操作>
        <link rel="icon" href="/static/favicon.png">
        <title>07 构建模式：Go Module的6类常规操作 </title>
        
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
                        <a class="menu-item" id="00 开篇词 这样入门Go，才能少走弯路.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e8%bf%99%e6%a0%b7%e5%85%a5%e9%97%a8Go%ef%bc%8c%e6%89%8d%e8%83%bd%e5%b0%91%e8%b5%b0%e5%bc%af%e8%b7%af.md">00 开篇词 这样入门Go，才能少走弯路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 前世今生：你不得不了解的Go的历史和现状.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/01%20%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f%ef%bc%9a%e4%bd%a0%e4%b8%8d%e5%be%97%e4%b8%8d%e4%ba%86%e8%a7%a3%e7%9a%84Go%e7%9a%84%e5%8e%86%e5%8f%b2%e5%92%8c%e7%8e%b0%e7%8a%b6.md">01 前世今生：你不得不了解的Go的历史和现状.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 拒绝“Hello and Bye”：Go语言的设计哲学是怎么一回事？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/02%20%e6%8b%92%e7%bb%9d%e2%80%9cHello%20and%20Bye%e2%80%9d%ef%bc%9aGo%e8%af%ad%e8%a8%80%e7%9a%84%e8%ae%be%e8%ae%a1%e5%93%b2%e5%ad%a6%e6%98%af%e6%80%8e%e4%b9%88%e4%b8%80%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">02 拒绝“Hello and Bye”：Go语言的设计哲学是怎么一回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 配好环境：选择一种最适合你的Go安装方法.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/03%20%e9%85%8d%e5%a5%bd%e7%8e%af%e5%a2%83%ef%bc%9a%e9%80%89%e6%8b%a9%e4%b8%80%e7%a7%8d%e6%9c%80%e9%80%82%e5%90%88%e4%bd%a0%e7%9a%84Go%e5%ae%89%e8%a3%85%e6%96%b9%e6%b3%95.md">03 配好环境：选择一种最适合你的Go安装方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 初窥门径：一个Go程序的结构是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/04%20%e5%88%9d%e7%aa%a5%e9%97%a8%e5%be%84%ef%bc%9a%e4%b8%80%e4%b8%aaGo%e7%a8%8b%e5%ba%8f%e7%9a%84%e7%bb%93%e6%9e%84%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">04 初窥门径：一个Go程序的结构是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 标准先行：Go项目的布局标准是什么？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/05%20%e6%a0%87%e5%87%86%e5%85%88%e8%a1%8c%ef%bc%9aGo%e9%a1%b9%e7%9b%ae%e7%9a%84%e5%b8%83%e5%b1%80%e6%a0%87%e5%87%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">05 标准先行：Go项目的布局标准是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 构建模式：Go是怎么解决包依赖管理问题的？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/06%20%e6%9e%84%e5%bb%ba%e6%a8%a1%e5%bc%8f%ef%bc%9aGo%e6%98%af%e6%80%8e%e4%b9%88%e8%a7%a3%e5%86%b3%e5%8c%85%e4%be%9d%e8%b5%96%e7%ae%a1%e7%90%86%e9%97%ae%e9%a2%98%e7%9a%84%ef%bc%9f.md">06 构建模式：Go是怎么解决包依赖管理问题的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 构建模式：Go Module的6类常规操作.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/07%20%e6%9e%84%e5%bb%ba%e6%a8%a1%e5%bc%8f%ef%bc%9aGo%20Module%e7%9a%846%e7%b1%bb%e5%b8%b8%e8%a7%84%e6%93%8d%e4%bd%9c.md">07 构建模式：Go Module的6类常规操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 入口函数与包初始化：搞清Go程序的执行次序.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/08%20%e5%85%a5%e5%8f%a3%e5%87%bd%e6%95%b0%e4%b8%8e%e5%8c%85%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%9a%e6%90%9e%e6%b8%85Go%e7%a8%8b%e5%ba%8f%e7%9a%84%e6%89%a7%e8%a1%8c%e6%ac%a1%e5%ba%8f.md">08 入口函数与包初始化：搞清Go程序的执行次序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 即学即练：构建一个Web服务就是这么简单.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/09%20%e5%8d%b3%e5%ad%a6%e5%8d%b3%e7%bb%83%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aaWeb%e6%9c%8d%e5%8a%a1%e5%b0%b1%e6%98%af%e8%bf%99%e4%b9%88%e7%ae%80%e5%8d%95.md">09 即学即练：构建一个Web服务就是这么简单.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 变量声明：静态语言有别于动态语言的重要特征.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/10%20%e5%8f%98%e9%87%8f%e5%a3%b0%e6%98%8e%ef%bc%9a%e9%9d%99%e6%80%81%e8%af%ad%e8%a8%80%e6%9c%89%e5%88%ab%e4%ba%8e%e5%8a%a8%e6%80%81%e8%af%ad%e8%a8%80%e7%9a%84%e9%87%8d%e8%a6%81%e7%89%b9%e5%be%81.md">10 变量声明：静态语言有别于动态语言的重要特征.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 代码块与作用域：如何保证变量不会被遮蔽？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/11%20%e4%bb%a3%e7%a0%81%e5%9d%97%e4%b8%8e%e4%bd%9c%e7%94%a8%e5%9f%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e5%8f%98%e9%87%8f%e4%b8%8d%e4%bc%9a%e8%a2%ab%e9%81%ae%e8%94%bd%ef%bc%9f.md">11 代码块与作用域：如何保证变量不会被遮蔽？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 基本数据类型：Go原生支持的数值类型有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/12%20%e5%9f%ba%e6%9c%ac%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9aGo%e5%8e%9f%e7%94%9f%e6%94%af%e6%8c%81%e7%9a%84%e6%95%b0%e5%80%bc%e7%b1%bb%e5%9e%8b%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">12 基本数据类型：Go原生支持的数值类型有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 基本数据类型：为什么Go要原生支持字符串类型？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/13%20%e5%9f%ba%e6%9c%ac%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Go%e8%a6%81%e5%8e%9f%e7%94%9f%e6%94%af%e6%8c%81%e5%ad%97%e7%ac%a6%e4%b8%b2%e7%b1%bb%e5%9e%8b%ef%bc%9f.md">13 基本数据类型：为什么Go要原生支持字符串类型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 常量：Go在“常量”设计上的创新有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/14%20%e5%b8%b8%e9%87%8f%ef%bc%9aGo%e5%9c%a8%e2%80%9c%e5%b8%b8%e9%87%8f%e2%80%9d%e8%ae%be%e8%ae%a1%e4%b8%8a%e7%9a%84%e5%88%9b%e6%96%b0%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">14 常量：Go在“常量”设计上的创新有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 同构复合类型：从定长数组到变长切片.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/15%20%e5%90%8c%e6%9e%84%e5%a4%8d%e5%90%88%e7%b1%bb%e5%9e%8b%ef%bc%9a%e4%bb%8e%e5%ae%9a%e9%95%bf%e6%95%b0%e7%bb%84%e5%88%b0%e5%8f%98%e9%95%bf%e5%88%87%e7%89%87.md">15 同构复合类型：从定长数组到变长切片.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 复合数据类型：原生map类型的实现机制是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/16%20%e5%a4%8d%e5%90%88%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9a%e5%8e%9f%e7%94%9fmap%e7%b1%bb%e5%9e%8b%e7%9a%84%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">16 复合数据类型：原生map类型的实现机制是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 复合数据类型：用结构体建立对真实世界的抽象.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/17%20%e5%a4%8d%e5%90%88%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9a%e7%94%a8%e7%bb%93%e6%9e%84%e4%bd%93%e5%bb%ba%e7%ab%8b%e5%af%b9%e7%9c%9f%e5%ae%9e%e4%b8%96%e7%95%8c%e7%9a%84%e6%8a%bd%e8%b1%a1.md">17 复合数据类型：用结构体建立对真实世界的抽象.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 控制结构：if的“快乐路径”原则.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/18%20%e6%8e%a7%e5%88%b6%e7%bb%93%e6%9e%84%ef%bc%9aif%e7%9a%84%e2%80%9c%e5%bf%ab%e4%b9%90%e8%b7%af%e5%be%84%e2%80%9d%e5%8e%9f%e5%88%99.md">18 控制结构：if的“快乐路径”原则.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 控制结构：Go的for循环，仅此一种.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/19%20%e6%8e%a7%e5%88%b6%e7%bb%93%e6%9e%84%ef%bc%9aGo%e7%9a%84for%e5%be%aa%e7%8e%af%ef%bc%8c%e4%bb%85%e6%ad%a4%e4%b8%80%e7%a7%8d.md">19 控制结构：Go的for循环，仅此一种.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 控制结构：Go中的switch语句有哪些变化？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/20%20%e6%8e%a7%e5%88%b6%e7%bb%93%e6%9e%84%ef%bc%9aGo%e4%b8%ad%e7%9a%84switch%e8%af%ad%e5%8f%a5%e6%9c%89%e5%93%aa%e4%ba%9b%e5%8f%98%e5%8c%96%ef%bc%9f.md">20 控制结构：Go中的switch语句有哪些变化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 函数：请叫我“一等公民”.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/21%20%e5%87%bd%e6%95%b0%ef%bc%9a%e8%af%b7%e5%8f%ab%e6%88%91%e2%80%9c%e4%b8%80%e7%ad%89%e5%85%ac%e6%b0%91%e2%80%9d.md">21 函数：请叫我“一等公民”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 函数：怎么结合多返回值进行错误处理？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/22%20%e5%87%bd%e6%95%b0%ef%bc%9a%e6%80%8e%e4%b9%88%e7%bb%93%e5%90%88%e5%a4%9a%e8%bf%94%e5%9b%9e%e5%80%bc%e8%bf%9b%e8%a1%8c%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%9f.md">22 函数：怎么结合多返回值进行错误处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 函数：怎么让函数更简洁健壮？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/23%20%e5%87%bd%e6%95%b0%ef%bc%9a%e6%80%8e%e4%b9%88%e8%ae%a9%e5%87%bd%e6%95%b0%e6%9b%b4%e7%ae%80%e6%b4%81%e5%81%a5%e5%a3%ae%ef%bc%9f.md">23 函数：怎么让函数更简洁健壮？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 方法：理解“方法”的本质.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/24%20%e6%96%b9%e6%b3%95%ef%bc%9a%e7%90%86%e8%a7%a3%e2%80%9c%e6%96%b9%e6%b3%95%e2%80%9d%e7%9a%84%e6%9c%ac%e8%b4%a8.md">24 方法：理解“方法”的本质.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 方法：方法集合与如何选择receiver类型？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/25%20%e6%96%b9%e6%b3%95%ef%bc%9a%e6%96%b9%e6%b3%95%e9%9b%86%e5%90%88%e4%b8%8e%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9receiver%e7%b1%bb%e5%9e%8b%ef%bc%9f.md">25 方法：方法集合与如何选择receiver类型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 方法：如何用类型嵌入模拟实现“继承”？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/26%20%e6%96%b9%e6%b3%95%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e7%b1%bb%e5%9e%8b%e5%b5%8c%e5%85%a5%e6%a8%a1%e6%8b%9f%e5%ae%9e%e7%8e%b0%e2%80%9c%e7%bb%a7%e6%89%bf%e2%80%9d%ef%bc%9f.md">26 方法：如何用类型嵌入模拟实现“继承”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 即学即练：跟踪函数调用链，理解代码更直观.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/27%20%e5%8d%b3%e5%ad%a6%e5%8d%b3%e7%bb%83%ef%bc%9a%e8%b7%9f%e8%b8%aa%e5%87%bd%e6%95%b0%e8%b0%83%e7%94%a8%e9%93%be%ef%bc%8c%e7%90%86%e8%a7%a3%e4%bb%a3%e7%a0%81%e6%9b%b4%e7%9b%b4%e8%a7%82.md">27 即学即练：跟踪函数调用链，理解代码更直观.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 接口：接口即契约.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/28%20%e6%8e%a5%e5%8f%a3%ef%bc%9a%e6%8e%a5%e5%8f%a3%e5%8d%b3%e5%a5%91%e7%ba%a6.md">28 接口：接口即契约.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 接口：为什么nil接口不等于nil？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/29%20%e6%8e%a5%e5%8f%a3%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88nil%e6%8e%a5%e5%8f%a3%e4%b8%8d%e7%ad%89%e4%ba%8enil%ef%bc%9f.md">29 接口：为什么nil接口不等于nil？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 接口：Go中最强大的魔法.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/30%20%e6%8e%a5%e5%8f%a3%ef%bc%9aGo%e4%b8%ad%e6%9c%80%e5%bc%ba%e5%a4%a7%e7%9a%84%e9%ad%94%e6%b3%95.md">30 接口：Go中最强大的魔法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 并发：Go的并发方案实现方案是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/31%20%e5%b9%b6%e5%8f%91%ef%bc%9aGo%e7%9a%84%e5%b9%b6%e5%8f%91%e6%96%b9%e6%a1%88%e5%ae%9e%e7%8e%b0%e6%96%b9%e6%a1%88%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">31 并发：Go的并发方案实现方案是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 并发：聊聊Goroutine调度器的原理.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/32%20%e5%b9%b6%e5%8f%91%ef%bc%9a%e8%81%8a%e8%81%8aGoroutine%e8%b0%83%e5%ba%a6%e5%99%a8%e7%9a%84%e5%8e%9f%e7%90%86.md">32 并发：聊聊Goroutine调度器的原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 并发：小channel中蕴含大智慧.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/33%20%e5%b9%b6%e5%8f%91%ef%bc%9a%e5%b0%8fchannel%e4%b8%ad%e8%95%b4%e5%90%ab%e5%a4%a7%e6%99%ba%e6%85%a7.md">33 并发：小channel中蕴含大智慧.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 并发：如何使用共享变量？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/34%20%e5%b9%b6%e5%8f%91%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%ef%bc%9f.md">34 并发：如何使用共享变量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 即学即练：如何实现一个轻量级线程池？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/35%20%e5%8d%b3%e5%ad%a6%e5%8d%b3%e7%bb%83%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e8%bd%bb%e9%87%8f%e7%ba%a7%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f.md">35 即学即练：如何实现一个轻量级线程池？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 打稳根基：怎么实现一个TCP服务器？（上）.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/36%20%e6%89%93%e7%a8%b3%e6%a0%b9%e5%9f%ba%ef%bc%9a%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aaTCP%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">36 打稳根基：怎么实现一个TCP服务器？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 代码操练：怎么实现一个TCP服务器？（中）.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/37%20%e4%bb%a3%e7%a0%81%e6%93%8d%e7%bb%83%ef%bc%9a%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aaTCP%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9f%ef%bc%88%e4%b8%ad%ef%bc%89.md">37 代码操练：怎么实现一个TCP服务器？（中）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 成果优化：怎么实现一个TCP服务器？（下）.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/38%20%e6%88%90%e6%9e%9c%e4%bc%98%e5%8c%96%ef%bc%9a%e6%80%8e%e4%b9%88%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aaTCP%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">38 成果优化：怎么实现一个TCP服务器？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 驯服泛型：了解类型参数.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/39%20%e9%a9%af%e6%9c%8d%e6%b3%9b%e5%9e%8b%ef%bc%9a%e4%ba%86%e8%a7%a3%e7%b1%bb%e5%9e%8b%e5%8f%82%e6%95%b0.md">39 驯服泛型：了解类型参数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 驯服泛型：定义泛型约束.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/40%20%e9%a9%af%e6%9c%8d%e6%b3%9b%e5%9e%8b%ef%bc%9a%e5%ae%9a%e4%b9%89%e6%b3%9b%e5%9e%8b%e7%ba%a6%e6%9d%9f.md">40 驯服泛型：定义泛型约束.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 驯服泛型：明确使用时机.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/41%20%e9%a9%af%e6%9c%8d%e6%b3%9b%e5%9e%8b%ef%bc%9a%e6%98%8e%e7%a1%ae%e4%bd%bf%e7%94%a8%e6%97%b6%e6%9c%ba.md">41 驯服泛型：明确使用时机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="元旦快乐 这是一份暂时停更的声明.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%85%83%e6%97%a6%e5%bf%ab%e4%b9%90%20%e8%bf%99%e6%98%af%e4%b8%80%e4%bb%bd%e6%9a%82%e6%97%b6%e5%81%9c%e6%9b%b4%e7%9a%84%e5%a3%b0%e6%98%8e.md">元旦快乐 这是一份暂时停更的声明.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 作为Go Module的作者，你应该知道的几件事.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e4%bd%9c%e4%b8%baGo%20Module%e7%9a%84%e4%bd%9c%e8%80%85%ef%bc%8c%e4%bd%a0%e5%ba%94%e8%af%a5%e7%9f%a5%e9%81%93%e7%9a%84%e5%87%a0%e4%bb%b6%e4%ba%8b.md">加餐 作为Go Module的作者，你应该知道的几件事.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 如何拉取私有的Go Module？.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e5%a6%82%e4%bd%95%e6%8b%89%e5%8f%96%e7%a7%81%e6%9c%89%e7%9a%84Go%20Module%ef%bc%9f.md">加餐 如何拉取私有的Go Module？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 我“私藏”的那些优质且权威的Go语言学习资料.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e6%88%91%e2%80%9c%e7%a7%81%e8%97%8f%e2%80%9d%e7%9a%84%e9%82%a3%e4%ba%9b%e4%bc%98%e8%b4%a8%e4%b8%94%e6%9d%83%e5%a8%81%e7%9a%84Go%e8%af%ad%e8%a8%80%e5%ad%a6%e4%b9%a0%e8%b5%84%e6%96%99.md">加餐 我“私藏”的那些优质且权威的Go语言学习资料.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 聊聊Go 1.17版本的那些新特性.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e8%81%8a%e8%81%8aGo%201.17%e7%89%88%e6%9c%ac%e7%9a%84%e9%82%a3%e4%ba%9b%e6%96%b0%e7%89%b9%e6%80%a7.md">加餐 聊聊Go 1.17版本的那些新特性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 聊聊Go语言的指针.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e8%81%8a%e8%81%8aGo%e8%af%ad%e8%a8%80%e7%9a%84%e6%8c%87%e9%92%88.md">加餐 聊聊Go语言的指针.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 聊聊最近大热的Go泛型.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%8a%a0%e9%a4%90%20%e8%81%8a%e8%81%8a%e6%9c%80%e8%bf%91%e5%a4%a7%e7%83%ad%e7%9a%84Go%e6%b3%9b%e5%9e%8b.md">加餐 聊聊最近大热的Go泛型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 叶剑峰：Go语言中常用的那些代码优化点.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e5%8f%b6%e5%89%91%e5%b3%b0%ef%bc%9aGo%e8%af%ad%e8%a8%80%e4%b8%ad%e5%b8%b8%e7%94%a8%e7%9a%84%e9%82%a3%e4%ba%9b%e4%bb%a3%e7%a0%81%e4%bc%98%e5%8c%96%e7%82%b9.md">大咖助阵 叶剑峰：Go语言中常用的那些代码优化点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 大明：Go泛型，泛了，但没有完全泛.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e5%a4%a7%e6%98%8e%ef%bc%9aGo%e6%b3%9b%e5%9e%8b%ef%bc%8c%e6%b3%9b%e4%ba%86%ef%bc%8c%e4%bd%86%e6%b2%a1%e6%9c%89%e5%ae%8c%e5%85%a8%e6%b3%9b.md">大咖助阵 大明：Go泛型，泛了，但没有完全泛.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 孔令飞：从小白到“老鸟”，我的Go语言进阶之路.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e5%ad%94%e4%bb%a4%e9%a3%9e%ef%bc%9a%e4%bb%8e%e5%b0%8f%e7%99%bd%e5%88%b0%e2%80%9c%e8%80%81%e9%b8%9f%e2%80%9d%ef%bc%8c%e6%88%91%e7%9a%84Go%e8%af%ad%e8%a8%80%e8%bf%9b%e9%98%b6%e4%b9%8b%e8%b7%af.md">大咖助阵 孔令飞：从小白到“老鸟”，我的Go语言进阶之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 徐祥曦：从销售到分布式存储工程师，我与 Go  的故事.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e5%be%90%e7%a5%a5%e6%9b%a6%ef%bc%9a%e4%bb%8e%e9%94%80%e5%94%ae%e5%88%b0%e5%88%86%e5%b8%83%e5%bc%8f%e5%ad%98%e5%82%a8%e5%b7%a5%e7%a8%8b%e5%b8%88%ef%bc%8c%e6%88%91%e4%b8%8e%20Go%20%20%e7%9a%84%e6%95%85%e4%ba%8b.md">大咖助阵 徐祥曦：从销售到分布式存储工程师，我与 Go  的故事.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 曹春晖：聊聊 Go 语言的 GC 实现.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e6%9b%b9%e6%98%a5%e6%99%96%ef%bc%9a%e8%81%8a%e8%81%8a%20Go%20%e8%af%ad%e8%a8%80%e7%9a%84%20GC%20%e5%ae%9e%e7%8e%b0.md">大咖助阵 曹春晖：聊聊 Go 语言的 GC 实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助阵 海纳：聊聊语言中的类型系统与泛型.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e5%a4%a7%e5%92%96%e5%8a%a9%e9%98%b5%20%e6%b5%b7%e7%ba%b3%ef%bc%9a%e8%81%8a%e8%81%8a%e8%af%ad%e8%a8%80%e4%b8%ad%e7%9a%84%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%e4%b8%8e%e6%b3%9b%e5%9e%8b.md">大咖助阵 海纳：聊聊语言中的类型系统与泛型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="期中测试 一起检验下你的学习成果吧.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e6%9c%9f%e4%b8%ad%e6%b5%8b%e8%af%95%20%e4%b8%80%e8%b5%b7%e6%a3%80%e9%aa%8c%e4%b8%8b%e4%bd%a0%e7%9a%84%e5%ad%a6%e4%b9%a0%e6%88%90%e6%9e%9c%e5%90%a7.md">期中测试 一起检验下你的学习成果吧.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 罗杰：我的Go语言学习之路.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e7%bd%97%e6%9d%b0%ef%bc%9a%e6%88%91%e7%9a%84Go%e8%af%ad%e8%a8%80%e5%ad%a6%e4%b9%a0%e4%b9%8b%e8%b7%af.md">用户故事 罗杰：我的Go语言学习之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 和你一起迎接Go的黄金十年.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%92%8c%e4%bd%a0%e4%b8%80%e8%b5%b7%e8%bf%8e%e6%8e%a5Go%e7%9a%84%e9%bb%84%e9%87%91%e5%8d%81%e5%b9%b4.md">结束语 和你一起迎接Go的黄金十年.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结课测试 快来检验下你的学习成果吧！.md" href="/%e4%b8%93%e6%a0%8f/Tony%20Bai%20%c2%b7%20Go%e8%af%ad%e8%a8%80%e7%ac%ac%e4%b8%80%e8%af%be/%e7%bb%93%e8%af%be%e6%b5%8b%e8%af%95%20%e5%bf%ab%e6%9d%a5%e6%a3%80%e9%aa%8c%e4%b8%8b%e4%bd%a0%e7%9a%84%e5%ad%a6%e4%b9%a0%e6%88%90%e6%9e%9c%e5%90%a7%ef%bc%81.md">结课测试 快来检验下你的学习成果吧！.md</a>
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
                            <h1 id="title" data-id="07 构建模式：Go Module的6类常规操作" class="title">07 构建模式：Go Module的6类常规操作</h1>
                            <div><p>你好，我是Tony Bai。</p>

<p>通过上一节课的讲解，我们掌握了Go Module构建模式的基本概念和工作原理，也初步学会了如何通过go mod命令，将一个Go项目转变为一个Go Module，并通过Go Module构建模式进行构建。</p>

<p>但是，围绕一个Go Module，Go开发人员每天要执行很多Go命令对其进行维护。这些维护又是怎么进行的呢？</p>

<p>具体来说，维护Go Module 无非就是对Go Module 依赖包的管理。但在具体工作中还有很多情况，我们接下来会拆分成六个场景，层层深入给你分析。可以说，学好这些是每个Go开发人员成长的必经之路。</p>

<p>我们首先来看一下日常进行Go应用开发时遇到的最为频繁的一个场景：<strong>为当前项目添加一个依赖包</strong>。</p>

<h2 id="为当前module添加一个依赖">为当前module添加一个依赖</h2>

<p>在一个项目的初始阶段，我们会经常为项目引入第三方包，并借助这些包完成特定功能。即便是项目进入了稳定阶段，随着项目的演进，我们偶尔还需要在代码中引入新的第三方包。</p>

<p>那么我们如何为一个Go Module添加一个新的依赖包呢？</p>

<p>我们还是以上一节课中讲过的module-mode项目为例。如果我们要为这个项目增加一个新依赖：github.com/google/uuid，那需要怎么做呢？</p>

<p>我们首先会更新源码，就像下面代码中这样：</p>

<pre><code class="language-plain">package main

import (
	&quot;github.com/google/uuid&quot;
	&quot;github.com/sirupsen/logrus&quot;
)

func main() {
	logrus.Println(&quot;hello, go module mode&quot;)
	logrus.Println(uuid.NewString())
}
</code></pre>

<p>新源码中，我们通过import语句导入了github.com/google/uuid，并在main函数中调用了uuid包的函数NewString。此时，如果我们直接构建这个module，我们会得到一个错误提示：</p>

<pre><code class="language-plain">$go build
main.go:4:2: no required module provides package github.com/google/uuid; to add it:
	go get github.com/google/uuid
</code></pre>

<p>Go编译器提示我们，go.mod里的require段中，没有哪个module提供了github.com/google/uuid包，如果我们要增加这个依赖，可以手动执行go get命令。那我们就来按照提示手工执行一下这个命令：</p>

<pre><code class="language-plain">$go get github.com/google/uuid
go: downloading github.com/google/uuid v1.3.0
go get: added github.com/google/uuid v1.3.0
</code></pre>

<p>你会发现，go get命令将我们新增的依赖包下载到了本地module缓存里，并在go.mod文件的require段中新增了一行内容：</p>

<pre><code class="language-plain">require (
	github.com/google/uuid v1.3.0 //新增的依赖
	github.com/sirupsen/logrus v1.8.1
)
</code></pre>

<p>这新增的一行表明，我们当前项目依赖的是uuid的v1.3.0版本。我们也可以使用go mod tidy命令，在执行构建前自动分析源码中的依赖变化，识别新增依赖项并下载它们：</p>

<pre><code class="language-plain">$go mod tidy
go: finding module for package github.com/google/uuid
go: found github.com/google/uuid in github.com/google/uuid v1.3.0
</code></pre>

<p>对于我们这个例子而言，手工执行go get新增依赖项，和执行go mod tidy自动分析和下载依赖项的最终效果，是等价的。但对于复杂的项目变更而言，逐一手工添加依赖项显然很没有效率，go mod tidy是更佳的选择。</p>

<p>到这里，我们已经了解了怎么为当前的module添加一个新的依赖。但是在日常开发场景中，我们需要对依赖的版本进行更改。那这又要怎么做呢？下面我们就来看看下面升、降级修改依赖版本的场景。</p>

<h2 id="升级-降级依赖的版本">升级/降级依赖的版本</h2>

<p>我们先以对依赖的版本进行降级为例，分析一下。</p>

<p>在实际开发工作中，如果我们认为Go命令自动帮我们确定的某个依赖的版本存在一些问题，比如，引入了不必要复杂性导致可靠性下降、性能回退等等，我们可以手工将它降级为之前发布的某个兼容版本。</p>

<p>那这个操作依赖于什么原理呢？</p>

<p>答案就是我们上一节课讲过“语义导入版本”机制。我们再来简单复习一下，Go Module的版本号采用了语义版本规范，也就是版本号使用vX.Y.Z的格式。其中X是主版本号，Y为次版本号(minor)，Z为补丁版本号(patch)。主版本号相同的两个版本，较新的版本是兼容旧版本的。如果主版本号不同，那么两个版本是不兼容的。</p>

<p>有了语义版本号作为基础和前提，我们就可以从容地手工对依赖的版本进行升降级了，Go命令也可以根据版本兼容性，自动选择出合适的依赖版本了。</p>

<p>我们还是以上面提到过的logrus为例，logrus现在就存在着多个发布版本，我们可以通过下面命令来进行查询：</p>

<pre><code class="language-plain">$go list -m -versions github.com/sirupsen/logrus
github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1
</code></pre>

<p>在这个例子中，基于初始状态执行的go mod tidy命令，帮我们选择了logrus的最新发布版本v1.8.1。如果你觉得这个版本存在某些问题，想将logrus版本降至某个之前发布的兼容版本，比如v1.7.0，<strong>那么我们可以在项目的module根目录下，执行带有版本号的go get命令：</strong></p>

<pre><code class="language-plain">$go get github.com/sirupsen/<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="365a595144434576400718011806">[email&#160;protected]</a>
go: downloading github.com/sirupsen/logrus v1.7.0
go get: downgraded github.com/sirupsen/logrus v1.8.1 =&gt; v1.7.0
</code></pre>

<p>从这个执行输出的结果，我们可以看到，go get命令下载了logrus v1.7.0版本，并将go.mod中对logrus的依赖版本从v1.8.1降至v1.7.0。</p>

<p>当然我们也可以使用万能命令go mod tidy来帮助我们降级，但前提是首先要用go mod edit命令，明确告知我们要依赖v1.7.0版本，而不是v1.8.1，这个执行步骤是这样的：</p>

<pre><code class="language-plain">$go mod edit -require=github.com/sirupsen/<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="bdd1d2dacfc8cefdcb8c938a938d">[email&#160;protected]</a>
$go mod tidy       
go: downloading github.com/sirupsen/logrus v1.7.0
</code></pre>

<p>降级后，我们再假设logrus v1.7.1版本是一个安全补丁升级，修复了一个严重的安全漏洞，而且我们必须使用这个安全补丁版本，这就意味着我们需要将logrus依赖从v1.7.0升级到v1.7.1。</p>

<p>我们可以使用与降级同样的步骤来完成升级，这里我只列出了使用go get实现依赖版本升级的命令和输出结果，你自己动手试一下。</p>

<pre><code class="language-plain">$go get github.com/sirupsen/<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="600c0f071215132016514e574e51">[email&#160;protected]</a>
go: downloading github.com/sirupsen/logrus v1.7.1
go get: upgraded github.com/sirupsen/logrus v1.7.0 =&gt; v1.7.1
</code></pre>

<p>好了，到这里你就学会了如何对项目依赖包的版本进行升降级了。</p>

<p>但是你可能会发现一个问题，在前面的例子中，Go Module的依赖的主版本号都是1。根据我们上节课中学习的语义导入版本的规范，在Go Module构建模式下，当依赖的主版本号为0或1的时候，我们在Go源码中导入依赖包，不需要在包的导入路径上增加版本号，也就是：</p>

<pre><code class="language-plain">import github.com/user/repo/v0 等价于 import github.com/user/repo
import github.com/user/repo/v1 等价于 import github.com/user/repo
</code></pre>

<p>但是，如果我们要依赖的module的主版本号大于1，这又要怎么办呢？接着我们就来看看这个场景下该如何去做。</p>

<h2 id="添加一个主版本号大于1的依赖">添加一个主版本号大于1的依赖</h2>

<p>这里，我们还是先来回顾一下，上节课我们讲的语义版本规则中对主版本号大于1情况有没有相应的说明。</p>

<p>有的。语义导入版本机制有一个原则：<strong>如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的</strong>。也就是说，如果新旧两个包不兼容，那么我们就应该采用不同的导入路径。</p>

<p>按照语义版本规范，如果我们要为项目引入主版本号大于1的依赖，比如v2.0.0，那么由于这个版本与v1、v0开头的包版本都不兼容，我们在导入v2.0.0包时，不能再直接使用github.com/user/repo，而要使用像下面代码中那样不同的包导入路径：</p>

<pre><code class="language-plain">import github.com/user/repo/v2/xxx
</code></pre>

<p>也就是说，如果我们要为Go项目添加主版本号大于1的依赖，我们就需要使用“语义导入版本”机制，<strong>在声明它的导入路径的基础上，加上版本号信息</strong>。我们以“向module-mode项目添加github.com/go-redis/redis依赖包的v7版本”为例，看看添加步骤。</p>

<p>首先，我们在源码中，以空导入的方式导入v7版本的github.com/go-redis/redis包：</p>

<pre><code class="language-plain">package main

import (
	_ &quot;github.com/go-redis/redis/v7&quot; // “_”为空导入
	&quot;github.com/google/uuid&quot;
	&quot;github.com/sirupsen/logrus&quot;
)

func main() {
	logrus.Println(&quot;hello, go module mode&quot;)
	logrus.Println(uuid.NewString())
}
</code></pre>

<p>接下来的步骤就与添加兼容依赖一样，我们通过go get获取redis的v7版本：</p>

<pre><code class="language-plain">$go get github.com/go-redis/redis/v7
go: downloading github.com/go-redis/redis/v7 v7.4.1
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go get: added github.com/go-redis/redis/v7 v7.4.1
</code></pre>

<p>我们可以看到，go get为我们选择了go-redis v7版本下当前的最新版本v7.4.1。</p>

<p>不过呢，这里说的是为项目添加一个主版本号大于1的依赖的步骤。有些时候，出于要使用依赖包最新功能特性等原因，我们可能需要将某个依赖的版本升级为其不兼容版本，也就是主版本号不同的版本，这又该怎么做呢？</p>

<p>我们还以go-redis/redis这个依赖为例，将这个依赖从v7版本升级到最新的v8版本看看。</p>

<h2 id="升级依赖版本到一个不兼容版本">升级依赖版本到一个不兼容版本</h2>

<p>我们前面说了，按照语义导入版本的原则，不同主版本的包的导入路径是不同的。所以，同样地，我们这里也需要先将代码中redis包导入路径中的版本号改为v8：</p>

<pre><code class="language-plain">import (
	_ &quot;github.com/go-redis/redis/v8&quot;
	&quot;github.com/google/uuid&quot;
	&quot;github.com/sirupsen/logrus&quot;
)
</code></pre>

<p>接下来，我们再通过go get来获取v8版本的依赖包：</p>

<pre><code class="language-plain">$go get github.com/go-redis/redis/v8
go: downloading github.com/go-redis/redis/v8 v8.11.1
go: downloading github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f
go: downloading github.com/cespare/xxhash/v2 v2.1.1
go get: added github.com/go-redis/redis/v8 v8.11.1
</code></pre>

<p>这样，我们就完成了向一个不兼容依赖版本的升级。是不是很简单啊！</p>

<p>但是项目继续演化到一个阶段的时候，我们可能还需要移除对之前某个包的依赖。</p>

<h2 id="移除一个依赖">移除一个依赖</h2>

<p>我们还是看前面go-redis/redis示例，如果我们这个时候不需要再依赖go-redis/redis了，你会怎么做呢？</p>

<p>你可能会删除掉代码中对redis的空导入这一行，之后再利用go build命令成功地构建这个项目。</p>

<p>但你会发现，与添加一个依赖时Go命令给出友好提示不同，这次go build没有给出任何关于项目已经将go-redis/redis删除的提示，并且go.mod里require段中的go-redis/redis/v8的依赖依旧存在着。</p>

<p>我们再通过go list命令列出当前module的所有依赖，你也会发现go-redis/redis/v8仍出现在结果中：</p>

<pre><code class="language-plain">$go list -m all
github.com/bigwhite/module-mode
github.com/cespare/xxhash/v2 v2.1.1
github.com/davecgh/go-spew v1.1.1
... ...
github.com/go-redis/redis/v8 v8.11.1
... ...
gopkg.in/yaml.v2 v2.3.0
</code></pre>

<p>这是怎么回事呢？</p>

<p>其实，要想彻底从项目中移除go.mod中的依赖项，仅从源码中删除对依赖项的导入语句还不够。这是因为如果源码满足成功构建的条件，go build命令是不会“多管闲事”地清理go.mod中多余的依赖项的。</p>

<p>那正确的做法是怎样的呢？我们还得用go mod tidy命令，将这个依赖项彻底从Go Module构建上下文中清除掉。go mod tidy会自动分析源码依赖，而且将不再使用的依赖从go.mod和go.sum中移除。</p>

<p>到这里，其实我们已经分析了Go Module依赖包管理的5个常见情况了，但其实还有一种特殊情况，需要我们借用vendor机制。</p>

<h2 id="特殊情况-使用vendor">特殊情况：使用vendor</h2>

<p>你可能会感到有点奇怪，为什么Go Module的维护，还有要用vendor的情况？</p>

<p>其实，vendor机制虽然诞生于GOPATH构建模式主导的年代，但在Go Module构建模式下，它依旧被保留了下来，并且成为了Go Module构建机制的一个很好的补充。特别是在一些不方便访问外部网络，并且对Go应用构建性能敏感的环境，比如在一些内部的持续集成或持续交付环境（CI/CD）中，使用vendor机制可以实现与Go Module等价的构建。</p>

<p>和GOPATH构建模式不同，Go Module构建模式下，我们再也无需手动维护vendor目录下的依赖包了，Go提供了可以快速建立和更新vendor的命令，我们还是以前面的module-mode项目为例，通过下面命令为该项目建立vendor：</p>

<pre><code class="language-plain">$go mod vendor
$tree -LF 2 vendor
vendor
├── github.com/
│   ├── google/
│   ├── magefile/
│   └── sirupsen/
├── golang.org/
│   └── x/
└── modules.txt
</code></pre>

<p>我们看到，go mod vendor命令在vendor目录下，创建了一份这个项目的依赖包的副本，并且通过vendor/modules.txt记录了vendor下的module以及版本。</p>

<p>如果我们要基于vendor构建，而不是基于本地缓存的Go Module构建，我们需要在go build后面加上-mod=vendor参数。</p>

<p>在Go 1.14及以后版本中，如果Go项目的顶层目录下存在vendor目录，那么go build默认也会优先基于vendor构建，除非你给go build传入-mod=mod的参数。</p>

<h2 id="小结">小结</h2>

<p>好了，到这里，我们就完成了维护Go Module的全部常见场景的学习了，现在我们一起来回顾一下吧。</p>

<p>在通过go mod init为当前Go项目创建一个新的module后，随着项目的演进，我们在日常开发过程中，会遇到多种常见的维护Go Module的场景。</p>

<p>其中最常见的就是为项目添加一个依赖包，我们可以通过go get命令手工获取该依赖包的特定版本，更好的方法是通过go mod tidy命令让Go命令自动去分析新依赖并决定使用新依赖的哪个版本。</p>

<p>另外，还有几个场景需要你记住：</p>

<ul>
<li>通过go get我们可以升级或降级某依赖的版本，如果升级或降级前后的版本不兼容，这里千万注意别忘了变化包导入路径中的版本号，这是Go语义导入版本机制的要求；</li>
<li>通过go mod tidy，我们可以自动分析Go源码的依赖变更，包括依赖的新增、版本变更以及删除，并更新go.mod中的依赖信息。</li>
<li>通过go mod vendor，我们依旧可以支持vendor机制，并且可以对vendor目录下缓存的依赖包进行自动管理。</li>
</ul>

<p>在了解了如何应对Go Modules维护的日常工作场景后，你是不是有一种再也不担心Go源码构建问题的感觉了呢？</p>

<h2 id="思考题">思考题</h2>

<p>如果你是一个公共Go包的作者，在发布你的Go包时，有哪些需要注意的地方？</p>

<p>感谢你和我一起学习，也欢迎你把这节课分享给更多对Go构建模式感兴趣的朋友。我是Tony Bai，我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#731f1f1f4a474242434433141e121a1f5d101c1e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1297238d46ede4',t:'MTczNDA1NzYzNi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;复合数据类型：原生map类型的实现机制是怎样的？>
        <link rel="icon" href="/static/favicon.png">
        <title>16 复合数据类型：原生map类型的实现机制是怎样的？ </title>
        
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
                            <h1 id="title" data-id="16 复合数据类型：原生map类型的实现机制是怎样的？" class="title">16 复合数据类型：原生map类型的实现机制是怎样的？</h1>
                            <div><p>你好，我是Tony Bai。</p>

<p>上一节课，我们学习了Go语言中最常用的两个复合类型：数组与切片。它们代表<strong>一组连续存储的同构类型元素集合。</strong>不同的是，数组的长度是确定的，而切片，我们可以理解为一种“动态数组”，它的长度在运行时是可变的。</p>

<p>这一节课，我们会继续前面的脉络，学习另外一种日常Go编码中比较常用的复合类型，这种类型可以让你将一个值（Value）唯一关联到一个特定的键（Key）上，可以用于实现特定键值的快速查找与更新，这个复合数据类型就是<strong>map</strong>。很多中文Go编程语言类技术书籍都会将它翻译为映射、哈希表或字典，但在我的课程中，<strong>为了保持原汁原味，我就直接使用它的英文名，map</strong>。</p>

<p>map是我们既切片之后，学到的第二个由Go编译器与运行时联合实现的复合数据类型，它有着复杂的内部实现，但却提供了十分简单友好的开发者使用接口。这一节课，我将从map类型的定义，到它的使用，再到map内部实现机制，由浅到深地让你吃透map类型。</p>

<h2 id="什么是map类型">什么是map类型？</h2>

<p>map是Go语言提供的一种抽象数据类型，它表示一组无序的键值对。在后面的讲解中，我们会直接使用key和value分别代表map的键和值。而且，map集合中每个key都是唯一的：</p>

<p><img src="assets/f818dcaf10cd42f1bb5d4049037f6f06.jpg" alt="图片" /></p>

<p>和切片类似，作为复合类型的map，它在Go中的类型表示也是由key类型与value类型组成的，就像下面代码：</p>

<pre><code class="language-plain">map[key_type]value_type
</code></pre>

<p>key与value的类型可以相同，也可以不同：</p>

<pre><code class="language-plain">map[string]string // key与value元素的类型相同
map[int]string    // key与value元素的类型不同
</code></pre>

<p>如果两个map类型的key元素类型相同，value元素类型也相同，那么我们可以说它们是同一个map类型，否则就是不同的map类型。</p>

<p>这里，我们要注意，map类型对value的类型没有限制，但是对key的类型却有严格要求，因为map类型要保证key的唯一性。Go语言中要求，<strong>key的类型必须支持“==”和“!=”两种比较操作符</strong>。</p>

<p>但是，在Go语言中，函数类型、map类型自身，以及切片只支持与nil的比较，而不支持同类型两个变量的比较。如果像下面代码这样，进行这些类型的比较，Go编译器将会报错：</p>

<pre><code class="language-plain">s1 := make([]int, 1)
s2 := make([]int, 2)
f1 := func() {}
f2 := func() {}
m1 := make(map[int]string)
m2 := make(map[int]string)
println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)
</code></pre>

<p>因此在这里，你一定要注意：<strong>函数类型、map类型自身，以及切片类型是不能作为map的key类型的</strong>。</p>

<p>知道如何表示一个map类型后，接下来，我们来看看如何声明和初始化一个map类型的变量。</p>

<h2 id="map变量的声明和初始化">map变量的声明和初始化</h2>

<p>我们可以这样声明一个map变量：</p>

<pre><code class="language-plain">var m map[string]int // 一个map[string]int类型的变量
</code></pre>

<p>和切片类型变量一样，如果我们没有显式地赋予map变量初值，map类型变量的默认值为nil。</p>

<p>不过切片变量和map变量在这里也有些不同。初值为零值nil的切片类型变量，可以借助内置的append的函数进行操作，这种在Go语言中被称为“<strong>零值可用</strong>”。定义“零值可用”的类型，可以提升我们开发者的使用体验，我们不用再担心变量的初始状态是否有效。</p>

<p><strong>但map类型，因为它内部实现的复杂性，无法“零值可用”</strong>。所以，如果我们对处于零值状态的map变量直接进行操作，就会导致运行时异常（panic），从而导致程序进程异常退出：</p>

<pre><code class="language-plain">var m map[string]int // m = nil
m[&quot;key&quot;] = 1         // 发生运行时异常：panic: assignment to entry in nil map
</code></pre>

<p>所以，我们必须对map类型变量进行显式初始化后才能使用。那我们怎样对map类型变量进行初始化呢？</p>

<p>和切片一样，为map类型变量显式赋值有两种方式：一种是使用复合字面值；另外一种是使用make这个预声明的内置函数。</p>

<p><strong>方法一：使用复合字面值初始化map类型变量。</strong></p>

<p>我们先来看这句代码：</p>

<pre><code class="language-plain">m := map[int]string{}
</code></pre>

<p>这里，我们显式初始化了map类型变量m。不过，你要注意，虽然此时map类型变量m中没有任何键值对，但变量m也不等同于初值为nil的map变量。这个时候，我们对m进行键值对的插入操作，不会引发运行时异常。</p>

<p>这里我们再看看怎么通过稍微复杂一些的复合字面值，对map类型变量进行初始化：</p>

<pre><code class="language-plain">m1 := map[int][]string{
    1: []string{&quot;val1_1&quot;, &quot;val1_2&quot;},
    3: []string{&quot;val3_1&quot;, &quot;val3_2&quot;, &quot;val3_3&quot;},
    7: []string{&quot;val7_1&quot;},
}

type Position struct { 
    x float64 
    y float64
}

m2 := map[Position]string{
    Position{29.935523, 52.568915}: &quot;school&quot;,
    Position{25.352594, 113.304361}: &quot;shopping-mall&quot;,
    Position{73.224455, 111.804306}: &quot;hospital&quot;,
}
</code></pre>

<p>我们看到，上面代码虽然完成了对两个map类型变量m1和m2的显式初始化，但不知道你有没有发现一个问题，作为初值的字面值似乎有些“臃肿”。你看，作为初值的字面值采用了复合类型的元素类型，而且在编写字面值时还带上了各自的元素类型，比如作为map[int] []string值类型的[]string，以及作为map[Position]string的key类型的Position。</p>

<p>别急！针对这种情况，Go提供了“语法糖”。这种情况下，<strong>Go允许省略字面值中的元素类型</strong>。因为map类型表示中包含了key和value的元素类型，Go编译器已经有足够的信息，来推导出字面值中各个值的类型了。我们以m2为例，这里的显式初始化代码和上面变量m2的初始化代码是等价的：</p>

<pre><code class="language-plain">m2 := map[Position]string{
    {29.935523, 52.568915}: &quot;school&quot;,
    {25.352594, 113.304361}: &quot;shopping-mall&quot;,
    {73.224455, 111.804306}: &quot;hospital&quot;,
}
</code></pre>

<p>以后在无特殊说明的情况下，我们都将使用这种简化后的字面值初始化方式。</p>

<p><strong>方法二：使用make为map类型变量进行显式初始化。</strong></p>

<p>和切片通过make进行初始化一样，通过make的初始化方式，我们可以为map类型变量指定键值对的初始容量，但无法进行具体的键值对赋值，就像下面代码这样：</p>

<pre><code class="language-plain">m1 := make(map[int]string) // 未指定初始容量
m2 := make(map[int]string, 8) // 指定初始容量为8
</code></pre>

<p>不过，map类型的容量不会受限于它的初始容量值，当其中的键值对数量超过初始容量后，Go运行时会自动增加map类型的容量，保证后续键值对的正常插入。</p>

<p>了解完map类型变量的声明与初始化后，我们就来看看，在日常开发中，map类型都有哪些基本操作和注意事项。</p>

<h2 id="map的基本操作">map的基本操作</h2>

<p>针对一个map类型变量，我们可以进行诸如插入新键值对、获取当前键值对数量、查找特定键和读取对应值、删除键值对，以及遍历键值等操作。我们一个个来学习。</p>

<p><strong>操作一：插入新键值对。</strong></p>

<p>面对一个非nil的map类型变量，我们可以在其中插入符合map类型定义的任意新键值对。插入新键值对的方式很简单，我们只需要把value赋值给map中对应的key就可以了：</p>

<pre><code class="language-plain">m := make(map[int]string)
m[1] = &quot;value1&quot;
m[2] = &quot;value2&quot;
m[3] = &quot;value3&quot;
</code></pre>

<p>而且，我们不需要自己判断数据有没有插入成功，因为Go会保证插入总是成功的。这里，Go运行时会负责map变量内部的内存管理，因此除非是系统内存耗尽，我们可以不用担心向map中插入新数据的数量和执行结果。</p>

<p>不过，如果我们插入新键值对的时候，某个key已经存在于map中了，那我们的插入操作就会用新值覆盖旧值：</p>

<pre><code class="language-plain">m := map[string]int {
	&quot;key1&quot; : 1,
	&quot;key2&quot; : 2,
}

m[&quot;key1&quot;] = 11 // 11会覆盖掉&quot;key1&quot;对应的旧值1
m[&quot;key3&quot;] = 3  // 此时m为map[key1:11 key2:2 key3:3]
</code></pre>

<p>从这段代码中你可以看到，map类型变量m在声明的同时就做了初始化，它的内部建立了两个键值对，其中就包含键key1。所以后面我们再给键key1进行赋值时，Go不会重新创建key1键，而是会用新值(11)把key1键对应的旧值(1)替换掉。</p>

<p><strong>操作二：获取键值对数量。</strong></p>

<p>如果我们在编码中，想知道当前map类型变量中已经建立了多少个键值对，那我们可以怎么做呢？和切片一样，map类型也可以通过内置函数<strong>len</strong>，获取当前变量已经存储的键值对数量：</p>

<pre><code class="language-plain">m := map[string]int {
	&quot;key1&quot; : 1,
	&quot;key2&quot; : 2,
}

fmt.Println(len(m)) // 2
m[&quot;key3&quot;] = 3  
fmt.Println(len(m)) // 3
</code></pre>

<p>不过，这里要注意的是<strong>我们不能对map类型变量调用cap，来获取当前容量</strong>，这是map类型与切片类型的一个不同点。</p>

<p><strong>操作三：查找和数据读取</strong></p>

<p>和写入相比，map类型更多用在查找和数据读取场合。所谓查找，就是判断某个key是否存在于某个map中。有了前面向map插入键值对的基础，我们可能自然而然地想到，可以用下面代码去查找一个键并获得该键对应的值：</p>

<pre><code class="language-plain">m := make(map[string]int)
v := m[&quot;key1&quot;]
</code></pre>

<p>乍一看，第二行代码在语法上好像并没有什么不当之处，但其实通过这行语句，我们还是无法确定键key1是否真实存在于map中。这是因为，当我们尝试去获取一个键对应的值的时候，如果这个键在map中并不存在，我们也会得到一个值，这个值是value元素类型的<strong>零值</strong>。</p>

<p>我们以上面这个代码为例，如果键key1在map中并不存在，那么v的值就会被赋予value元素类型int的零值，也就是0。所以我们无法通过v值判断出，究竟是因为key1不存在返回的零值，还是因为key1本身对应的value就是0。</p>

<p>那么在map中查找key的正确姿势是什么呢？Go语言的map类型支持通过用一种名为“<strong>comma ok</strong>”的惯用法，进行对某个key的查询。接下来我们就用“comma ok”惯用法改造一下上面的代码：</p>

<pre><code class="language-plain">m := make(map[string]int)
v, ok := m[&quot;key1&quot;]
if !ok {
    // &quot;key1&quot;不在map中
}

// &quot;key1&quot;在map中，v将被赋予&quot;key1&quot;键对应的value
</code></pre>

<p>我们看到，这里我们通过了一个布尔类型变量ok，来判断键“key1”是否存在于map中。如果存在，变量v就会被正确地赋值为键“key1”对应的value。</p>

<p>不过，如果我们并不关心某个键对应的value，而只关心某个键是否在于map中，我们可以使用空标识符替代变量v，忽略可能返回的value：</p>

<pre><code class="language-plain">m := make(map[string]int)
_, ok := m[&quot;key1&quot;]
... ...
</code></pre>

<p>因此，你一定要记住：<strong>在Go语言中，请使用“comma ok”惯用法对map进行键查找和键值读取操作。</strong></p>

<p><strong>操作四：删除数据。</strong></p>

<p>接下来，我们再看看看如何从map中删除某个键值对。在Go中，我们需要借助<strong>内置函数delete</strong>来从map中删除数据。使用delete函数的情况下，传入的第一个参数是我们的map类型变量，第二个参数就是我们想要删除的键。我们可以看看这个代码示例：</p>

<pre><code class="language-plain">m := map[string]int {
	&quot;key1&quot; : 1,
	&quot;key2&quot; : 2,
}

fmt.Println(m) // map[key1:1 key2:2]
delete(m, &quot;key2&quot;) // 删除&quot;key2&quot;
fmt.Println(m) // map[key1:1]
</code></pre>

<p>这里要注意的是，<strong>delete函数是从map中删除键的唯一方法</strong>。即便传给delete的键在map中并不存在，delete函数的执行也不会失败，更不会抛出运行时的异常。</p>

<p><strong>操作五：遍历map中的键值数据</strong></p>

<p>最后，我们来说一下如何遍历map中的键值数据。这一点虽然不像查询和读取操作那么常见，但日常开发中我们还是有这个需求的。在Go中，遍历map的键值对只有一种方法，那就是<strong>像对待切片那样通过for range语句对map数据进行遍历</strong>。我们看一个例子：</p>

<pre><code class="language-plain">package main
  
import &quot;fmt&quot;

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    fmt.Printf(&quot;{ &quot;)
    for k, v := range m {
        fmt.Printf(&quot;[%d, %d] &quot;, k, v)
    }
    fmt.Printf(&quot;}\n&quot;)
}
</code></pre>

<p>你看，通过for range遍历map变量m，每次迭代都会返回一个键值对，其中键存在于变量k中，它对应的值存储在变量v中。我们可以运行一下这段代码，可以得到符合我们预期的结果：</p>

<pre><code class="language-plain">{ [1, 11] [2, 12] [3, 13] }
</code></pre>

<p>如果我们只关心每次迭代的键，我们可以使用下面的方式对map进行遍历：</p>

<pre><code class="language-plain">for k, _ := range m { 
	// 使用k
}
</code></pre>

<p>当然更地道的方式是这样的：</p>

<pre><code class="language-plain">for k := range m {
	// 使用k
}
</code></pre>

<p>如果我们只关心每次迭代返回的键所对应的value，我们同样可以通过空标识符替代变量k，就像下面这样：</p>

<pre><code class="language-plain">for _, v := range m {
	// 使用v
}
</code></pre>

<p>不过，前面map遍历的输出结果都非常理想，给我们的表象好像是迭代器按照map中元素的插入次序逐一遍历。那事实是不是这样呢？我们再来试试，多遍历几次这个map看看。</p>

<p>我们先来改造一下代码：</p>

<pre><code class="language-plain">package main
  
import &quot;fmt&quot;

func doIteration(m map[int]int) {
    fmt.Printf(&quot;{ &quot;)
    for k, v := range m {
        fmt.Printf(&quot;[%d, %d] &quot;, k, v)
    }
    fmt.Printf(&quot;}\n&quot;)
}

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    for i := 0; i &lt; 3; i++ {
        doIteration(m)
    }
}
</code></pre>

<p>运行一下上述代码，我们可以得到这样结果：</p>

<pre><code class="language-plain">{ [3, 13] [1, 11] [2, 12] }
{ [1, 11] [2, 12] [3, 13] }
{ [3, 13] [1, 11] [2, 12] }
</code></pre>

<p>我们可以看到，<strong>对同一map做多次遍历的时候，每次遍历元素的次序都不相同</strong>。这是Go语言map类型的一个重要特点，也是很容易让Go初学者掉入坑中的一个地方。所以这里你一定要记住：<strong>程序逻辑千万不要依赖遍历map所得到的的元素次序</strong>。</p>

<p>从我们前面的讲解，你应该也感受到了，map类型非常好用，那么，我们在各个函数方法间传递map变量会不会有很大开销呢？</p>

<h2 id="map变量的传递开销">map变量的传递开销</h2>

<p>其实你不用担心开销的问题。</p>

<p>和切片类型一样，map也是引用类型。这就意味着map类型变量作为参数被传递给函数或方法的时候，实质上传递的只是一个<strong>“描述符”</strong>（后面我们再讲这个描述符究竟是什么)，而不是整个map的数据拷贝，所以这个传递的开销是固定的，而且也很小。</p>

<p>并且，当map变量被传递到函数或方法内部后，我们在函数内部对map类型参数的修改在函数外部也是可见的。比如你从这个示例中就可以看到，函数foo中对map类型变量m进行了修改，而这些修改在foo函数外也可见。</p>

<pre><code class="language-plain">package main
  
import &quot;fmt&quot;

func foo(m map[string]int) {
    m[&quot;key1&quot;] = 11
    m[&quot;key2&quot;] = 12
}

func main() {
    m := map[string]int{
        &quot;key1&quot;: 1,
        &quot;key2&quot;: 2,
    }

    fmt.Println(m) // map[key1:1 key2:2]  
    foo(m)
    fmt.Println(m) // map[key1:11 key2:12] 
}
</code></pre>

<h2 id="map的内部实现">map的内部实现</h2>

<p>和切片相比，map类型的内部实现要更加复杂。Go运行时使用一张哈希表来实现抽象的map类型。运行时实现了map类型操作的所有功能，包括查找、插入、删除等。在编译阶段，Go编译器会将Go语法层面的map操作，重写成运行时对应的函数调用。大致的对应关系是这样的：</p>

<pre><code class="language-plain">// 创建map类型变量实例
m := make(map[keyType]valType, capacityhint) → m := runtime.makemap(maptype, capacityhint, m)

// 插入新键值对或给键重新赋值
m[&quot;key&quot;] = &quot;value&quot; → v := runtime.mapassign(maptype, m, &quot;key&quot;) v是用于后续存储value的空间的地址

// 获取某键的值 
v := m[&quot;key&quot;]      → v := runtime.mapaccess1(maptype, m, &quot;key&quot;)
v, ok := m[&quot;key&quot;]  → v, ok := runtime.mapaccess2(maptype, m, &quot;key&quot;)

// 删除某键
delete(m, &quot;key&quot;)   → runtime.mapdelete(maptype, m, “key”)
</code></pre>

<p>这是map类型在Go运行时层实现的示意图：</p>

<p><img src="assets/32d0f6b128b646fdb7d80af77d0f81d2.jpg" alt="" /></p>

<p>我们可以看到，和切片的运行时表示图相比，map的实现示意图显然要复杂得多。接下来，我们结合这张图来简要描述一下map在运行时层的实现原理。我们重点讲解一下一个map变量在初始状态、进行键值对操作后，以及在并发场景下的Go运行时层的实现原理。</p>

<h3 id="初始状态">初始状态</h3>

<p>从图中我们可以看到，与语法层面 map 类型变量（m）一一对应的是*runtime.hmap 的实例，即runtime.hmap类型的指针，也就是我们前面在讲解 map 类型变量传递开销时提到的 <strong>map 类型的描述符</strong>。hmap 类型是 map 类型的头部结构（header），它存储了后续 map 类型操作所需的所有信息，包括：</p>

<p><img src="assets/becf4c3571a04ceaafa7c26c4318af59.jpg" alt="图片" /></p>

<p>真正用来存储键值对数据的是桶，也就是bucket，每个bucket中存储的是Hash值低bit位数值相同的元素，默认的元素个数为 BUCKETSIZE（值为 8，Go 1.17版本中在$GOROOT/src/cmd/compile/internal/reflectdata/reflect.go中定义，与 runtime/map.go 中常量 bucketCnt 保持一致）。</p>

<p>当某个bucket（比如buckets[0])的8个空槽slot）都填满了，且map尚未达到扩容的条件的情况下，运行时会建立overflow bucket，并将这个overflow bucket挂在上面bucket（如buckets[0]）末尾的overflow指针上，这样两个buckets形成了一个链表结构，直到下一次map扩容之前，这个结构都会一直存在。</p>

<p>从图中我们可以看到，每个bucket由三部分组成，从上到下分别是tophash区域、key存储区域和value存储区域。</p>

<ul>
<li><strong>tophash区域</strong></li>
</ul>

<p>当我们向map插入一条数据，或者是从map按key查询数据的时候，运行时都会使用哈希函数对key做哈希运算，并获得一个哈希值（hashcode）。这个hashcode非常关键，运行时会把hashcode“一分为二”来看待，其中低位区的值用于选定bucket，高位区的值用于在某个bucket中确定key的位置。我把这一过程整理成了下面这张示意图，你理解起来可以更直观：</p>

<p><img src="assets/0de27d5e3bc34c10a04960b9249c0eba.jpg" alt="图片" /></p>

<p>因此，每个bucket的tophash区域其实是用来快速定位key位置的，这样就避免了逐个key进行比较这种代价较大的操作。尤其是当key是size较大的字符串类型时，好处就更突出了。这是一种以空间换时间的思路。</p>

<ul>
<li><strong>key存储区域</strong></li>
</ul>

<p>接着，我们看tophash区域下面是一块连续的内存区域，存储的是这个bucket承载的所有key数据。运行时在分配bucket的时候需要知道key的Size。那么运行时是如何知道key的size的呢？</p>

<p>当我们声明一个map类型变量，比如var m map[string]int时，Go运行时就会为这个变量对应的特定map类型，生成一个runtime.maptype实例。如果这个实例已经存在，就会直接复用。maptype实例的结构是这样的：</p>

<pre><code class="language-plain">type maptype struct {
    typ        _type
    key        *_type
    elem       *_type
    bucket     *_type // internal type representing a hash bucket
    keysize    uint8  // size of key slot
    elemsize   uint8  // size of elem slot
    bucketsize uint16 // size of bucket
    flags      uint32
} 
</code></pre>

<p>我们可以看到，这个实例包含了我们需要的map类型中的所有&rdquo;元信息&rdquo;。我们前面提到过，编译器会把语法层面的map操作重写成运行时对应的函数调用，这些运行时函数都有一个共同的特点，那就是第一个参数都是maptype指针类型的参数。</p>

<p><strong>Go运行时就是利用maptype参数中的信息确定key的类型和大小的。</strong>map所用的hash函数也存放在maptype.key.alg.hash(key, hmap.hash0)中。同时maptype的存在也让Go中所有map类型都共享一套运行时map操作函数，而不是像C++那样为每种map类型创建一套map操作函数，这样就节省了对最终二进制文件空间的占用。</p>

<ul>
<li><strong>value存储区域</strong></li>
</ul>

<p>我们再接着看key存储区域下方的另外一块连续的内存区域，这个区域存储的是key对应的value。和key一样，这个区域的创建也是得到了maptype中信息的帮助。Go运行时采用了把key和value分开存储的方式，而不是采用一个kv接着一个kv的kv紧邻方式存储，这带来的其实是算法上的复杂性，但却减少了因内存对齐带来的内存浪费。</p>

<p>我们以map[int8]int64为例，看看下面的存储空间利用率对比图：</p>

<p><img src="assets/bb51290ebd26427aa5d6bc389cabc01e.jpg" alt="图片" /></p>

<p>你会看到，当前Go运行时使用的方案内存利用效率很高，而kv紧邻存储的方案在map[int8]int64这样的例子中内存浪费十分严重，它的内存利用率是72/128=56.25%，有近一半的空间都浪费掉了。</p>

<p>另外，还有一点我要跟你强调一下，如果key或value的数据长度大于一定数值，那么运行时不会在bucket中直接存储数据，而是会存储key或value数据的指针。目前Go运行时定义的最大key和value的长度是这样的：</p>

<pre><code class="language-plain">// $GOROOT/src/runtime/map.go
const (
    maxKeySize  = 128
    maxElemSize = 128
)
</code></pre>

<h3 id="map扩容">map扩容</h3>

<p>我们前面提到过，map会对底层使用的内存进行自动管理。因此，在使用过程中，当插入元素个数超出一定数值后，map一定会存在自动扩容的问题，也就是怎么扩充bucket的数量，并重新在bucket间均衡分配数据的问题。</p>

<p>那么map在什么情况下会进行扩容呢？Go运行时的map实现中引入了一个LoadFactor（负载因子），当<strong>count &gt; LoadFactor * 2^B</strong>或overflow bucket过多时，运行时会自动对map进行扩容。目前Go最新1.17版本LoadFactor设置为6.5（loadFactorNum/loadFactorDen）。这里是Go中与map扩容相关的部分源码：</p>

<pre><code class="language-plain">// $GOROOT/src/runtime/map.go
const (
	... ...

	loadFactorNum = 13
	loadFactorDen = 2
	... ...
)

func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	... ...
	if !h.growing() &amp;&amp; (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}
	... ...
}
</code></pre>

<p>这两方面原因导致的扩容，在运行时的操作其实是不一样的。如果是因为overflow bucket过多导致的“扩容”，实际上运行时会新建一个和现有规模一样的bucket数组，然后在assign和delete时做排空和迁移。</p>

<p>如果是因为当前数据数量超出LoadFactor指定水位而进行的扩容，那么运行时会建立一个<strong>两倍于现有规模的bucket数组</strong>，但真正的排空和迁移工作也是在assign和delete时逐步进行的。原bucket数组会挂在hmap的oldbuckets指针下面，直到原buckets数组中所有数据都迁移到新数组后，原buckets数组才会被释放。你可以结合下面的map扩容示意图来理解这个过程，这会让你理解得更深刻一些：</p>

<p><img src="assets/11cc9ceaac72450eae10a77b17ed1df2.jpg" alt="图片" /></p>

<h3 id="map与并发">map与并发</h3>

<p>接着我们来看一下map和并发。从上面的实现原理来看，充当map描述符角色的hmap实例自身是有状态的（hmap.flags），而且对状态的读写是没有并发保护的。所以说map实例不是并发写安全的，也不支持并发读写。如果我们对map实例进行并发读写，程序运行时就会抛出异常。你可以看看下面这个并发读写map的例子：</p>

<pre><code class="language-plain">package main

import (
    &quot;fmt&quot;
    &quot;time&quot;
)

func doIteration(m map[int]int) {
    for k, v := range m {
        _ = fmt.Sprintf(&quot;[%d, %d] &quot;, k, v)
    }
}

func doWrite(m map[int]int) {
    for k, v := range m {
        m[k] = v + 1
    }
}

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    go func() {
        for i := 0; i &lt; 1000; i++ {
            doIteration(m)
        }
    }()

    go func() {
        for i := 0; i &lt; 1000; i++ {
            doWrite(m)
        }
    }()

    time.Sleep(5 * time.Second)
}
</code></pre>

<p>运行这个示例程序，我们会得到下面的执行错误结果：</p>

<pre><code class="language-plain">fatal error: concurrent map iteration and map write
</code></pre>

<p>不过，如果我们仅仅是进行并发读，map是没有问题的。而且，Go 1.9版本中引入了支持并发写安全的sync.Map类型，可以在并发读写的场景下替换掉map。如果你有这方面的需求，可以查看一下<a href="https://pkg.go.dev/sync#Map" target="_blank">sync.Map的手册</a>。</p>

<p>另外，你要注意，考虑到map可以自动扩容，map中数据元素的value位置可能在这一过程中发生变化，所以<strong>Go不允许获取map中value的地址，这个约束是在编译期间就生效的</strong>。下面这段代码就展示了Go编译器识别出获取map中value地址的语句后，给出的编译错误：</p>

<pre><code class="language-plain">p := &amp;m[key]  // cannot take the address of m[key]
fmt.Println(p)
</code></pre>

<h2 id="小结">小结</h2>

<p>好了，今天的课讲到这里就结束了。这一节课，我们讲解了Go语言的另一类十分常用的复合数据类型：map。</p>

<p>在Go语言中，map类型是一个无序的键值对的集合。它有两种类型元素，一类是键（key），另一类是值（value）。在一个map中，键是唯一的，在集合中不能有两个相同的键。Go也是通过这两种元素类型来表示一个map类型，你要记得这个通用的map类型表示：“map[key_type]value_type”。</p>

<p>map类型对key元素的类型是有约束的，它要求key元素的类型必须支持&rdquo;==“和”!=&ldquo;两个比较操作符。value元素的类型可以是任意的。</p>

<p>不过，map类型变量声明后必须对它进行初始化后才能操作。map类型支持插入新键值对、查找和数据读取、删除键值对、遍历map中的键值数据等操作，Go为开发者提供了十分简单的操作接口。这里要你重点记住的是，我们在查找和数据读取时一定要使用“comma ok”惯用法。此外，map变量在函数与方法间传递的开销很小，并且在函数内部通过map描述符对map的修改会对函数外部可见。</p>

<p>另外，map的内部实现要比切片复杂得多，它是由Go编译器与运行时联合实现的。Go编译器在编译阶段会将语法层面的map操作，重写为运行时对应的函数调用。Go运行时则采用了高效的算法实现了map类型的各类操作，这里我建议你要结合Go项目源码来理解map的具体实现。</p>

<p>和切片一样，map是Go语言提供的重要数据类型，也是Gopher日常Go编码是最常使用的类型之一。我们在日常使用map的场合要把握住下面几个要点，不要走弯路：</p>

<ul>
<li>不要依赖map的元素遍历顺序；</li>
<li>map不是线程安全的，不支持并发读写；</li>
<li>不要尝试获取map中元素（value）的地址。</li>
</ul>

<h2 id="思考题">思考题</h2>

<p>通过上面的学习，我们知道对map类型进行遍历所得到的键的次序是随机的，那么我想请你思考并实现一个方法，让我们能对map的进行稳定次序遍历？期待在留言区看到你的想法。</p>

<p>欢迎你把这节课分享给更多对Go语言map类型感兴趣的朋友。我是Tony Bai，我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#dfb3b3b3e6ebeeeeefe89fb8b2beb6b3f1bcb0b2" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12995f2828ede4',t:'MTczNDA1NzcyOC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
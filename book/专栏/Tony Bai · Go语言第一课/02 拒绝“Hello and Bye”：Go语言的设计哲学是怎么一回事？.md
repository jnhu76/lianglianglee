<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=02&#32;拒绝“Hello&#32;and&#32;Bye”：Go语言的设计哲学是怎么一回事？>
        <link rel="icon" href="/static/favicon.png">
        <title>02 拒绝“Hello and Bye”：Go语言的设计哲学是怎么一回事？ </title>
        
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
                            <h1 id="title" data-id="02 拒绝“Hello and Bye”：Go语言的设计哲学是怎么一回事？" class="title">02 拒绝“Hello and Bye”：Go语言的设计哲学是怎么一回事？</h1>
                            <div><p>你好，我是Tony Bai。</p>

<p>上一讲，我们探讨了<strong>“Go从哪里来，并可能要往哪里去”</strong>的问题。根据“绝大多数主流编程语言将在其15至20年间大步前进”这个依据，我们给出了一个结论：<strong>Go语言即将进入自己的黄金5~10年</strong>。</p>

<p>那么此时此刻，想必你已经跃跃欲试，想要尽快开启Go编程之旅。但在正式学习Go语法之前，我还是要再来给你<strong>泼泼冷水</strong>，因为这将决定你后续的学习结果，是“从入门到继续”还是“从入门到放弃”。</p>

<p>很多编程语言的初学者在学习初期，可能都会遇到这样的问题：最初兴致勃勃地开始学习一门编程语言，学着学着就发现了很多“别扭”的地方，比如想要的语言特性缺失、语法风格冷僻与主流语言差异较大、语言的不同版本间无法兼容、语言的语法特性过多导致学习曲线陡峭、语言的工具链支持较差，等等。</p>

<p>其实以上的这些问题，本质上都与语言设计者的设计哲学有关。所谓编程语言的设计哲学，就是指决定这门语言演化进程的高级原则和依据。</p>

<p><strong>设计哲学之于编程语言，就好比一个人的价值观之于这个人的行为。</strong></p>

<p>因为如果你不认同一个人的价值观，那你其实很难与之持续交往下去，即所谓道不同不相为谋。类似的，如果你不认同一门编程语言的设计哲学，那么大概率你在后续的语言学习中，就会遇到上面提到的这些问题，而且可能会让你失去继续学习的精神动力。</p>

<p>因此，在真正开始学习Go语法和编码之前，我们还需要先来了解一下Go语言的设计哲学，等学完这一讲之后，你就能更深刻地认识到自己学习Go语言的原因了。</p>

<p>我将Go语言的设计哲学总结为五点：简单、显式、组合、并发和面向工程。下面，我们就先从Go语言的第一设计哲学“<strong>简单</strong>”开始了解吧。</p>

<h3 id="简单">简单</h3>

<p>知名Go开发者戴维·切尼（Dave Cheney）曾说过：“大多数编程语言创建伊始都致力于成为一门简单的语言，但最终都只是满足于做一个强大的编程语言”。</p>

<p><strong>而Go语言是一个例外。Go语言的设计者们在语言设计之初，就拒绝了走语言特性融合的道路，选择了“做减法”并致力于打造一门简单的编程语言。</strong></p>

<p>选择了“简单”，就意味着Go不会像C++、Java那样将其他编程语言的新特性兼蓄并收，所以你在Go语言中看不到传统的面向对象的类、构造函数与继承，看不到结构化的异常处理，也看不到本属于函数编程范式的语法元素。</p>

<p>其实，Go语言也没它看起来那么简单，自身实现起来并不容易，但这些复杂性被Go语言的设计者们“隐藏”了，所以Go语法层面上呈现了这样的状态：</p>

<ul>
<li>仅有25个关键字，主流编程语言最少；</li>
<li>内置垃圾收集，降低开发人员内存管理的心智负担；</li>
<li>首字母大小写决定可见性，无需通过额外关键字修饰；</li>
<li>变量初始为类型零值，避免以随机值作为初值的问题；</li>
<li>内置数组边界检查，极大减少越界访问带来的安全隐患；</li>
<li>内置并发支持，简化并发程序设计；</li>
<li>内置接口类型，为组合的设计哲学奠定基础；</li>
<li>原生提供完善的工具链，开箱即用；</li>
<li>… …</li>
</ul>

<p>看，我说的没错吧，确实挺简单的。当然了，任何的设计都存在着权衡与折中。我们看到Go设计者选择的“简单”，其实是站在巨人肩膀上，去除或优化了以往语言中，已经被开发者证明为体验不好或难以驾驭的语法元素和语言机制，并提出了自己的一些创新性的设计。比如，首字母大小写决定可见性、变量初始为类型零值、内置以go关键字实现的并发支持等。</p>

<p>Go这种有些“逆潮流”的“简单哲学”并不是一开始就能得到程序员的理解的，但在真正使用Go之后，我们才能真正体会到这种简单所带来的收益：简单意味着可以使用更少的代码实现相同的功能；简单意味着代码具有更好的可读性，而可读性好的代码通常意味着更好的可维护性以及可靠性。</p>

<p>总之，在软件工程化的今天，这些都意味着对生产效率提升的极大促进，我们可以认为<strong>简单的设计哲学是Go生产力的源泉</strong>。</p>

<h3 id="显式">显式</h3>

<p>好，接下来我们继续来了解学习下Go语言的第二大设计哲学：<strong>显式</strong>。</p>

<p>首先，我想先带你来看一段C程序，我们一起来看看“隐式”代码的行为特征。</p>

<p>在C语言中，下面这段代码可以正常编译并输出正确结果：</p>

<pre><code class="language-plain">#include &lt;stdio.h&gt;

int main() {
		short int a = 5;

		int b = 8;
		long c = 0;
		
		c = a + b;
		printf(&quot;%ld\n&quot;, c);
}
</code></pre>

<p>我们看到在上面这段代码中，变量a、b和c的类型均不相同，C语言编译器在编译<code>c = a + b</code>这一行时，会自动将短整型变量a和整型变量b，先转换为long类型然后相加，并将所得结果存储在long类型变量c中。那如果换成Go来实现这个计算会怎么样呢？我们先把上面的C程序转化成等价的Go代码：</p>

<pre><code class="language-plain">package main

import &quot;fmt&quot;

func main() {
    var a int16 = 5
    var b int = 8
    var c int64

    c = a + b
    fmt.Printf(&quot;%d\n&quot;, c)
}
</code></pre>

<p>如果我们编译这段程序，将得到类似这样的编译器错误：“invalid operation: a + b (mismatched types int16 and int)”。我们能看到Go与C语言的隐式自动类型转换不同，Go不允许不同类型的整型变量进行混合计算，它同样也不会对其进行隐式的自动转换。</p>

<p>因此，如果要使这段代码通过编译，我们就需要对变量a和b进行<strong>显式转型</strong>，就像下面代码段中这样：</p>

<pre><code class="language-plain">c = int64(a) + int64(b)
fmt.Printf(&quot;%d\n&quot;, c)
</code></pre>

<p>而这其实就是Go语言<strong>显式设计哲学</strong>的一个体现。</p>

<p>在Go语言中，不同类型变量是不能在一起进行混合计算的，这是因为<strong>Go希望开发人员明确知道自己在做什么</strong>，这与C语言的“信任程序员”原则完全不同，因此你需要以显式的方式通过转型统一参与计算各个变量的类型。</p>

<p>除此之外，Go设计者所崇尚的显式哲学还直接决定了Go语言错误处理的形态：Go语言采用了<strong>显式的基于值比较的错误处理方案</strong>，函数/方法中的错误都会通过return语句显式地返回，并且通常调用者不能忽略对返回的错误的处理。</p>

<p>这种有悖于“主流语言潮流”的错误处理机制还一度让开发者诟病，社区也提出了多个新错误处理方案，但或多或少都包含隐式的成分，都被Go开发团队一一否决了，这也与显式的设计哲学不无关系。</p>

<h3 id="组合">组合</h3>

<p>接着，我们来看第三个设计哲学：<strong>组合</strong>。</p>

<p>这个设计哲学和我们各个程序之间的耦合有关，Go语言不像C++、Java等主流面向对象语言，我们在Go中是找不到经典的面向对象语法元素、类型体系和继承机制的，Go推崇的是组合的设计哲学。</p>

<p>在诠释组合之前，我们需要先来了解一下Go在语法元素设计时，是如何为“组合”哲学的应用奠定基础的。</p>

<p>在Go语言设计层面，Go设计者为开发者们提供了正交的语法元素，以供后续组合使用，包括：</p>

<ul>
<li>Go语言无类型层次体系，各类型之间是相互独立的，没有子类型的概念；</li>
<li>每个类型都可以有自己的方法集合，类型定义与方法实现是正交独立的；</li>
<li>实现某个接口时，无需像Java那样采用特定关键字修饰；</li>
<li>包之间是相对独立的，没有子包的概念。</li>
</ul>

<p>我们可以看到，无论是包、接口还是一个个具体的类型定义，Go语言其实是为我们呈现了这样的一幅图景：一座座没有关联的“孤岛”，但每个岛内又都很精彩。那么现在摆在面前的工作，就是在这些孤岛之间以最适当的方式建立关联，并形成一个整体。而<strong>Go选择采用的组合方式，也是最主要的方式</strong>。</p>

<p>Go语言为支撑组合的设计提供了<strong>类型嵌入</strong>（Type Embedding）。通过类型嵌入，我们可以将已经实现的功能嵌入到新类型中，以快速满足新类型的功能需求，这种方式有些类似经典面向对象语言中的“继承”机制，但在原理上却与面向对象中的继承完全不同，这是一种Go设计者们精心设计的“语法糖”。</p>

<p>被嵌入的类型和新类型两者之间没有任何关系，甚至相互完全不知道对方的存在，更没有经典面向对象语言中的那种父类、子类的关系，以及向上、向下转型（Type Casting）。通过新类型实例调用方法时，方法的匹配主要取决于方法名字，而不是类型。这种组合方式，我称之为<strong>垂直组合</strong>，即通过类型嵌入，快速让一个新类型“复用”其他类型已经实现的能力，实现功能的垂直扩展。</p>

<p>你可以看看下面这个Go标准库中的一段使用类型嵌入的组合方式的代码段：</p>

<pre><code class="language-plain">// $GOROOT/src/sync/pool.go
type poolLocal struct {
    private interface{}   
    shared  []interface{}
    Mutex               
    pad     [128]byte  
}
</code></pre>

<p>在代码段中，我们在poolLocal这个结构体类型中嵌入了类型Mutex，这就使得poolLocal这个类型具有了互斥同步的能力，我们可以通过poolLocal类型的变量，直接调用Mutex类型的方法Lock或Unlock。-
另外，我们在标准库中还会经常看到类似如下定义接口类型的代码段：</p>

<pre><code class="language-plain">// $GOROOT/src/io/io.go
type ReadWriter interface {
    Reader
    Writer
}
</code></pre>

<p>这里，标准库通过嵌入接口类型的方式来实现接口行为的聚合，组成大接口，这种方式在标准库中尤为常用，并且已经成为了Go语言的一种惯用法。-
垂直组合本质上是一种“能力继承”，采用嵌入方式定义的新类型继承了嵌入类型的能力。Go还有一种常见的组合方式，叫<strong>水平组合</strong>。和垂直组合的能力继承不同，水平组合是一种能力委托（Delegate），我们通常使用接口类型来实现水平组合。</p>

<p>Go语言中的接口是一个创新设计，它只是方法集合，并且它与实现者之间的关系无需通过显式关键字修饰，它让程序内部各部分之间的耦合降至最低，同时它也是连接程序各个部分之间“纽带”。</p>

<p>水平组合的模式有很多，比如一种常见方法就是，通过接受接口类型参数的普通函数进行组合，如以下代码段所示：</p>

<pre><code class="language-plain">// $GOROOT/src/io/ioutil/ioutil.go
func ReadAll(r io.Reader)([]byte, error)

// $GOROOT/src/io/io.go
func Copy(dst Writer, src Reader)(written int64, err error)
</code></pre>

<p>也就是说，函数ReadAll通过io.Reader这个接口，将io.Reader的实现与ReadAll所在的包低耦合地水平组合在一起了，从而达到从任意实现io.Reader的数据源读取所有数据的目的。类似的水平组合“模式”还有点缀器、中间件等，这里我就不展开了，在后面讲到接口类型时再详细叙述。</p>

<p>此外，我们还可以将Go语言内置的并发能力进行灵活组合以实现，比如，通过goroutine+channel的组合，可以实现类似Unix Pipe的能力。</p>

<p>总之，组合原则的应用实质上是塑造了Go程序的骨架结构。类型嵌入为类型提供了垂直扩展能力，而接口是水平组合的关键，它好比程序肌体上的“关节”，给予连接“关节”的两个部分各自“自由活动”的能力，而整体上又实现了某种功能。并且，组合也让遵循“简单”原则的Go语言，在表现力上丝毫不逊色于其他复杂的主流编程语言。</p>

<h3 id="并发">并发</h3>

<p>好，前面我们已经看过3个设计哲学了，紧接着我带你看的是第4个：<strong>并发</strong>。</p>

<p>“并发”这个设计哲学的出现有它的背景，你也知道CPU都是靠提高主频来改进性能的，但是现在这个做法已经遇到了瓶颈。主频提高导致CPU的功耗和发热量剧增，反过来制约了CPU性能的进一步提高。2007年开始，处理器厂商的竞争焦点从主频转向了多核。</p>

<p>在这种大背景下，Go的设计者在决定去创建一门新语言的时候，果断将面向多核、<strong>原生支持并发</strong>作为了新语言的设计原则之一。并且，Go放弃了传统的基于操作系统线程的并发模型，而采用了<strong>用户层轻量级线程</strong>，Go将之称为<strong>goroutine</strong>。</p>

<p>goroutine占用的资源非常小，Go运行时默认为每个goroutine分配的栈空间仅2KB。goroutine调度的切换也不用陷入（trap）操作系统内核层完成，代价很低。因此，一个Go程序中可以创建成千上万个并发的goroutine。而且，所有的Go代码都在goroutine中执行，哪怕是go运行时的代码也不例外。</p>

<p>在提供了开销较低的goroutine的同时，Go还在语言层面内置了辅助并发设计的原语：channel和select。开发者可以通过语言内置的channel传递消息或实现同步，并通过select实现多路channel的并发控制。相较于传统复杂的线程并发模型，Go对并发的原生支持将大大降低开发人员在开发并发程序时的心智负担。</p>

<p>此外，并发的设计哲学不仅仅让Go在语法层面提供了并发原语支持，其对Go应用程序设计的影响更为重要。并发是一种程序结构设计的方法，它使得并行成为可能。</p>

<p>采用并发方案设计的程序在单核处理器上也是可以正常运行的，也许在单核上的处理性能可能不如非并发方案。但随着处理器核数的增多，并发方案可以自然地提高处理性能。</p>

<p>而且，并发与组合的哲学是一脉相承的，并发是一个更大的组合的概念，它在程序设计的全局层面对程序进行拆解组合，再映射到程序执行层面上：goroutines各自执行特定的工作，通过channel+select将goroutines组合连接起来。并发的存在鼓励程序员在程序设计时进行独立计算的分解，而对并发的原生支持让Go语言也更适应现代计算环境。</p>

<h3 id="面向工程">面向工程</h3>

<p>最后，我们来看一下Go的最后一条设计哲学：面向工程。</p>

<p>Go语言设计的初衷，就是<strong>面向解决真实世界中Google内部大规模软件开发存在的各种问题，为这些问题提供答案</strong>，这些问题包括：程序构建慢、依赖管理失控、代码难于理解、跨语言构建难等。</p>

<p>很多编程语言设计者和他们的粉丝们认为这些问题并不是一门编程语言应该去解决的，但Go语言的设计者并不这么看，他们在Go语言最初设计阶段就<strong>将解决工程问题作为Go的设计原则之一</strong>去考虑Go语法、工具链与标准库的设计，这也是Go与其他偏学院派、偏研究型的编程语言在设计思路上的一个重大差异。</p>

<p>语法是编程语言的用户接口，它直接影响开发人员对于这门语言的使用体验。在面向工程设计哲学的驱使下，Go在语法设计细节上做了精心的打磨。比如：</p>

<ul>
<li>重新设计编译单元和目标文件格式，实现Go源码快速构建，让大工程的构建时间缩短到类似动态语言的交互式解释的编译速度；</li>
<li>如果源文件导入它不使用的包，则程序将无法编译。这可以充分保证任何Go程序的依赖树是精确的。这也可以保证在构建程序时不会编译额外的代码，从而最大限度地缩短编译时间；</li>
<li>去除包的循环依赖，循环依赖会在大规模的代码中引发问题，因为它们要求编译器同时处理更大的源文件集，这会减慢增量构建；</li>
<li>包路径是唯一的，而包名不必唯一的。导入路径必须唯一标识要导入的包，而名称只是包的使用者如何引用其内容的约定。“包名称不必是唯一的”这个约定，大大降低了开发人员给包起唯一名字的心智负担；</li>
<li>故意不支持默认函数参数。因为在规模工程中，很多开发者利用默认函数参数机制，向函数添加过多的参数以弥补函数API的设计缺陷，这会导致函数拥有太多的参数，降低清晰度和可读性；</li>
<li>增加类型别名（type alias），支持大规模代码库的重构。</li>
</ul>

<p>在标准库方面，Go被称为“自带电池”的编程语言。如果说一门编程语言是“自带电池”，则说明这门语言标准库功能丰富，多数功能不需要依赖外部的第三方包或库，Go语言恰恰就是这类编程语言。</p>

<p>由于诞生年代较晚，而且目标比较明确，Go在标准库中提供了各类高质量且性能优良的功能包，其中的<code>net/http</code>、<code>crypto</code>、<code>encoding</code>等包充分迎合了云原生时代的关于API/RPC Web服务的构建需求，Go开发者可以直接基于标准库提供的这些包实现一个满足生产要求的API服务，从而减少对外部第三方包或库的依赖，降低工程代码依赖管理的复杂性，也降低了开发人员学习第三方库的心理负担。</p>

<p>而且，开发人员在工程过程中肯定是需要使用工具的，Go语言就提供了足以让所有其它主流语言开发人员羡慕的工具链，工具链涵盖了编译构建、代码格式化、包依赖管理、静态代码检查、测试、文档生成与查看、性能剖析、语言服务器、运行时程序跟踪等方方面面。</p>

<p>这里值得重点介绍的是<strong>gofmt</strong>，它统一了Go语言的代码风格，在其他语言开发者还在为代码风格争论不休的时候，Go开发者可以更加专注于领域业务中。同时，相同的代码风格让以往困扰开发者的代码阅读、理解和评审工作变得容易了很多，至少Go开发者再也不会有那种因代码风格的不同而产生的陌生感。Go的这种统一代码风格思路也在开始影响着后续新编程语言的设计，并且一些现有的主流编程语言也在借鉴Go的一些设计。</p>

<p>在提供丰富的工具链的同时，Go在标准库中提供了官方的词法分析器、语法解析器和类型检查器相关包，开发者可以基于这些包快速构建并扩展Go工具链。</p>

<h3 id="小结">小结</h3>

<p>好了，今天的课讲到这里就结束了，现在我们一起来回顾一下吧。</p>

<p>在这一讲中，我和你一起了解了Go语言的设计哲学：<strong>简单</strong>、<strong>显式</strong>、<strong>组合</strong>、<strong>并发和面向工程</strong>。</p>

<ul>
<li>简单是指Go语言特性始终保持在少且足够的水平，不走语言特性融合的道路，但又不乏生产力。简单是Go生产力的源泉，也是Go对开发者的最大吸引力；</li>
<li>显式是指任何代码行为都需开发者明确知晓，不存在因“暗箱操作”而导致可维护性降低和不安全的结果；</li>
<li>组合是构建Go程序骨架的主要方式，它可以大幅降低程序元素间的耦合，提高程序的可扩展性和灵活性；</li>
<li>并发是Go敏锐地把握了CPU向多核方向发展这一趋势的结果，可以让开发人员在多核时代更容易写出充分利用系统资源、支持性能随CPU核数增加而自然提升的应用程序；</li>
<li>面向工程是Go语言在语言设计上的一个重大创新，它将语言要解决的问题域扩展到那些原本并不是由编程语言去解决的领域，从而覆盖了更多开发者在开发过程遇到的“痛点”，为开发者提供了更好的使用体验。</li>
</ul>

<p>这些设计哲学直接影响了Go语言自身的设计。理解这些设计哲学，也能帮助我们理解Go语言语法、标准库以及工具链的演化决策过程。</p>

<p>好了，学完这节课之后，你认同Go的设计哲学吗？认同的话就继续跟着我学下去吧。</p>

<h3 id="思考题">思考题</h3>

<p>今天，我还想问下你，你还能举出哪些符合Go语言设计哲学的例子吗？欢迎在留言区多多和我分享讨论。</p>

<p>感谢你和我一起学习，也欢迎你把这节课分享给更多对Go语言的设计哲学感兴趣的朋友。我是Tony Bai，我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#aac6c6c6939e9b9b9a9deacdc7cbc3c684c9c5c7" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12961c2c0bede4',t:'MTczNDA1NzU5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
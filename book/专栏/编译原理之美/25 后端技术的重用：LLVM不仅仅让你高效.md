<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=25&#32;后端技术的重用：LLVM不仅仅让你高效>
        <link rel="icon" href="/static/favicon.png">
        <title>25 后端技术的重用：LLVM不仅仅让你高效 </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么你要学习编译原理？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e8%a6%81%e5%ad%a6%e4%b9%a0%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%ef%bc%9f.md">00 开篇词 为什么你要学习编译原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 理解代码：编译器的前端技术.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/01%20%e7%90%86%e8%a7%a3%e4%bb%a3%e7%a0%81%ef%bc%9a%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e5%89%8d%e7%ab%af%e6%8a%80%e6%9c%af.md">01 理解代码：编译器的前端技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 正则文法和有限自动机：纯手工打造词法分析器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/02%20%e6%ad%a3%e5%88%99%e6%96%87%e6%b3%95%e5%92%8c%e6%9c%89%e9%99%90%e8%87%aa%e5%8a%a8%e6%9c%ba%ef%bc%9a%e7%ba%af%e6%89%8b%e5%b7%a5%e6%89%93%e9%80%a0%e8%af%8d%e6%b3%95%e5%88%86%e6%9e%90%e5%99%a8.md">02 正则文法和有限自动机：纯手工打造词法分析器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 语法分析（一）：纯手工打造公式计算器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/03%20%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e7%ba%af%e6%89%8b%e5%b7%a5%e6%89%93%e9%80%a0%e5%85%ac%e5%bc%8f%e8%ae%a1%e7%ae%97%e5%99%a8.md">03 语法分析（一）：纯手工打造公式计算器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 语法分析（二）：解决二元表达式中的难点.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/04%20%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e8%a7%a3%e5%86%b3%e4%ba%8c%e5%85%83%e8%a1%a8%e8%be%be%e5%bc%8f%e4%b8%ad%e7%9a%84%e9%9a%be%e7%82%b9.md">04 语法分析（二）：解决二元表达式中的难点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 语法分析（三）：实现一门简单的脚本语言.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/05%20%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%ae%9e%e7%8e%b0%e4%b8%80%e9%97%a8%e7%ae%80%e5%8d%95%e7%9a%84%e8%84%9a%e6%9c%ac%e8%af%ad%e8%a8%80.md">05 语法分析（三）：实现一门简单的脚本语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 编译器前端工具（一）：用Antlr生成词法、语法分析器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/06%20%e7%bc%96%e8%af%91%e5%99%a8%e5%89%8d%e7%ab%af%e5%b7%a5%e5%85%b7%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e7%94%a8Antlr%e7%94%9f%e6%88%90%e8%af%8d%e6%b3%95%e3%80%81%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%e5%99%a8.md">06 编译器前端工具（一）：用Antlr生成词法、语法分析器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 编译器前端工具（二）：用Antlr重构脚本语言.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/07%20%e7%bc%96%e8%af%91%e5%99%a8%e5%89%8d%e7%ab%af%e5%b7%a5%e5%85%b7%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%94%a8Antlr%e9%87%8d%e6%9e%84%e8%84%9a%e6%9c%ac%e8%af%ad%e8%a8%80.md">07 编译器前端工具（二）：用Antlr重构脚本语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 作用域和生存期：实现块作用域和函数.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/08%20%e4%bd%9c%e7%94%a8%e5%9f%9f%e5%92%8c%e7%94%9f%e5%ad%98%e6%9c%9f%ef%bc%9a%e5%ae%9e%e7%8e%b0%e5%9d%97%e4%bd%9c%e7%94%a8%e5%9f%9f%e5%92%8c%e5%87%bd%e6%95%b0.md">08 作用域和生存期：实现块作用域和函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 面向对象：实现数据和方法的封装.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/09%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%9a%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e5%92%8c%e6%96%b9%e6%b3%95%e7%9a%84%e5%b0%81%e8%a3%85.md">09 面向对象：实现数据和方法的封装.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 闭包： 理解了原理，它就不反直觉了.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/10%20%e9%97%ad%e5%8c%85%ef%bc%9a%20%e7%90%86%e8%a7%a3%e4%ba%86%e5%8e%9f%e7%90%86%ef%bc%8c%e5%ae%83%e5%b0%b1%e4%b8%8d%e5%8f%8d%e7%9b%b4%e8%a7%89%e4%ba%86.md">10 闭包： 理解了原理，它就不反直觉了.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 语义分析（上）：如何建立一个完善的类型系统？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/11%20%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bb%ba%e7%ab%8b%e4%b8%80%e4%b8%aa%e5%ae%8c%e5%96%84%e7%9a%84%e7%b1%bb%e5%9e%8b%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">11 语义分析（上）：如何建立一个完善的类型系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 语义分析（下）：如何做上下文相关情况的处理？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/12%20%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%81%9a%e4%b8%8a%e4%b8%8b%e6%96%87%e7%9b%b8%e5%85%b3%e6%83%85%e5%86%b5%e7%9a%84%e5%a4%84%e7%90%86%ef%bc%9f.md">12 语义分析（下）：如何做上下文相关情况的处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 继承和多态：面向对象运行期的动态特性.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/13%20%e7%bb%a7%e6%89%bf%e5%92%8c%e5%a4%9a%e6%80%81%ef%bc%9a%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e8%bf%90%e8%a1%8c%e6%9c%9f%e7%9a%84%e5%8a%a8%e6%80%81%e7%89%b9%e6%80%a7.md">13 继承和多态：面向对象运行期的动态特性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 前端技术应用（一）：如何透明地支持数据库分库分表？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/14%20%e5%89%8d%e7%ab%af%e6%8a%80%e6%9c%af%e5%ba%94%e7%94%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%8f%e6%98%8e%e5%9c%b0%e6%94%af%e6%8c%81%e6%95%b0%e6%8d%ae%e5%ba%93%e5%88%86%e5%ba%93%e5%88%86%e8%a1%a8%ef%bc%9f.md">14 前端技术应用（一）：如何透明地支持数据库分库分表？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 前端技术应用（二）：如何设计一个报表工具？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/15%20%e5%89%8d%e7%ab%af%e6%8a%80%e6%9c%af%e5%ba%94%e7%94%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e4%b8%80%e4%b8%aa%e6%8a%a5%e8%a1%a8%e5%b7%a5%e5%85%b7%ef%bc%9f.md">15 前端技术应用（二）：如何设计一个报表工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 NFA和DFA：如何自己实现一个正则表达式工具？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/16%20NFA%e5%92%8cDFA%ef%bc%9a%e5%a6%82%e4%bd%95%e8%87%aa%e5%b7%b1%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e6%ad%a3%e5%88%99%e8%a1%a8%e8%be%be%e5%bc%8f%e5%b7%a5%e5%85%b7%ef%bc%9f.md">16 NFA和DFA：如何自己实现一个正则表达式工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 First和Follow集合：用LL算法推演一个实例.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/17%20First%e5%92%8cFollow%e9%9b%86%e5%90%88%ef%bc%9a%e7%94%a8LL%e7%ae%97%e6%b3%95%e6%8e%a8%e6%bc%94%e4%b8%80%e4%b8%aa%e5%ae%9e%e4%be%8b.md">17 First和Follow集合：用LL算法推演一个实例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 移进和规约：用LR算法推演一个实例.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/18%20%e7%a7%bb%e8%bf%9b%e5%92%8c%e8%a7%84%e7%ba%a6%ef%bc%9a%e7%94%a8LR%e7%ae%97%e6%b3%95%e6%8e%a8%e6%bc%94%e4%b8%80%e4%b8%aa%e5%ae%9e%e4%be%8b.md">18 移进和规约：用LR算法推演一个实例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 案例总结与热点问题答疑：对于左递归的语法，为什么我的推导不是左递归的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/19%20%e6%a1%88%e4%be%8b%e6%80%bb%e7%bb%93%e4%b8%8e%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%9a%e5%af%b9%e4%ba%8e%e5%b7%a6%e9%80%92%e5%bd%92%e7%9a%84%e8%af%ad%e6%b3%95%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e7%9a%84%e6%8e%a8%e5%af%bc%e4%b8%8d%e6%98%af%e5%b7%a6%e9%80%92%e5%bd%92%e7%9a%84%ef%bc%9f.md">19 案例总结与热点问题答疑：对于左递归的语法，为什么我的推导不是左递归的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 高效运行：编译器的后端技术.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/20%20%e9%ab%98%e6%95%88%e8%bf%90%e8%a1%8c%ef%bc%9a%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af.md">20 高效运行：编译器的后端技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 运行时机制：突破现象看本质，透过语法看运行时.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/21%20%e8%bf%90%e8%a1%8c%e6%97%b6%e6%9c%ba%e5%88%b6%ef%bc%9a%e7%aa%81%e7%a0%b4%e7%8e%b0%e8%b1%a1%e7%9c%8b%e6%9c%ac%e8%b4%a8%ef%bc%8c%e9%80%8f%e8%bf%87%e8%af%ad%e6%b3%95%e7%9c%8b%e8%bf%90%e8%a1%8c%e6%97%b6.md">21 运行时机制：突破现象看本质，透过语法看运行时.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 生成汇编代码（一）：汇编语言其实不难学.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/22%20%e7%94%9f%e6%88%90%e6%b1%87%e7%bc%96%e4%bb%a3%e7%a0%81%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%b1%87%e7%bc%96%e8%af%ad%e8%a8%80%e5%85%b6%e5%ae%9e%e4%b8%8d%e9%9a%be%e5%ad%a6.md">22 生成汇编代码（一）：汇编语言其实不难学.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 生成汇编代码（二）：把脚本编译成可执行文件.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/23%20%e7%94%9f%e6%88%90%e6%b1%87%e7%bc%96%e4%bb%a3%e7%a0%81%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%8a%8a%e8%84%9a%e6%9c%ac%e7%bc%96%e8%af%91%e6%88%90%e5%8f%af%e6%89%a7%e8%a1%8c%e6%96%87%e4%bb%b6.md">23 生成汇编代码（二）：把脚本编译成可执行文件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 中间代码：兼容不同的语言和硬件.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/24%20%e4%b8%ad%e9%97%b4%e4%bb%a3%e7%a0%81%ef%bc%9a%e5%85%bc%e5%ae%b9%e4%b8%8d%e5%90%8c%e7%9a%84%e8%af%ad%e8%a8%80%e5%92%8c%e7%a1%ac%e4%bb%b6.md">24 中间代码：兼容不同的语言和硬件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 后端技术的重用：LLVM不仅仅让你高效.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/25%20%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e7%9a%84%e9%87%8d%e7%94%a8%ef%bc%9aLLVM%e4%b8%8d%e4%bb%85%e4%bb%85%e8%ae%a9%e4%bd%a0%e9%ab%98%e6%95%88.md">25 后端技术的重用：LLVM不仅仅让你高效.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 生成IR：实现静态编译的语言.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/26%20%e7%94%9f%e6%88%90IR%ef%bc%9a%e5%ae%9e%e7%8e%b0%e9%9d%99%e6%80%81%e7%bc%96%e8%af%91%e7%9a%84%e8%af%ad%e8%a8%80.md">26 生成IR：实现静态编译的语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 代码优化：为什么你的代码比他的更高效？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/27%20%e4%bb%a3%e7%a0%81%e4%bc%98%e5%8c%96%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e7%9a%84%e4%bb%a3%e7%a0%81%e6%af%94%e4%bb%96%e7%9a%84%e6%9b%b4%e9%ab%98%e6%95%88%ef%bc%9f.md">27 代码优化：为什么你的代码比他的更高效？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 数据流分析：你写的程序，它更懂.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/28%20%e6%95%b0%e6%8d%ae%e6%b5%81%e5%88%86%e6%9e%90%ef%bc%9a%e4%bd%a0%e5%86%99%e7%9a%84%e7%a8%8b%e5%ba%8f%ef%bc%8c%e5%ae%83%e6%9b%b4%e6%87%82.md">28 数据流分析：你写的程序，它更懂.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 目标代码的生成和优化（一）：如何适应各种硬件架构？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/29%20%e7%9b%ae%e6%a0%87%e4%bb%a3%e7%a0%81%e7%9a%84%e7%94%9f%e6%88%90%e5%92%8c%e4%bc%98%e5%8c%96%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%82%e5%ba%94%e5%90%84%e7%a7%8d%e7%a1%ac%e4%bb%b6%e6%9e%b6%e6%9e%84%ef%bc%9f.md">29 目标代码的生成和优化（一）：如何适应各种硬件架构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 目标代码的生成和优化（二）：如何适应各种硬件架构？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/30%20%e7%9b%ae%e6%a0%87%e4%bb%a3%e7%a0%81%e7%9a%84%e7%94%9f%e6%88%90%e5%92%8c%e4%bc%98%e5%8c%96%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%82%e5%ba%94%e5%90%84%e7%a7%8d%e7%a1%ac%e4%bb%b6%e6%9e%b6%e6%9e%84%ef%bc%9f.md">30 目标代码的生成和优化（二）：如何适应各种硬件架构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 内存计算：对海量数据做计算，到底可以有多快？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/31%20%e5%86%85%e5%ad%98%e8%ae%a1%e7%ae%97%ef%bc%9a%e5%af%b9%e6%b5%b7%e9%87%8f%e6%95%b0%e6%8d%ae%e5%81%9a%e8%ae%a1%e7%ae%97%ef%bc%8c%e5%88%b0%e5%ba%95%e5%8f%af%e4%bb%a5%e6%9c%89%e5%a4%9a%e5%bf%ab%ef%bc%9f.md">31 内存计算：对海量数据做计算，到底可以有多快？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 字节码生成：为什么Spring技术很强大？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/32%20%e5%ad%97%e8%8a%82%e7%a0%81%e7%94%9f%e6%88%90%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Spring%e6%8a%80%e6%9c%af%e5%be%88%e5%bc%ba%e5%a4%a7%ef%bc%9f.md">32 字节码生成：为什么Spring技术很强大？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 垃圾收集：能否不停下整个世界？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/33%20%e5%9e%83%e5%9c%be%e6%94%b6%e9%9b%86%ef%bc%9a%e8%83%bd%e5%90%a6%e4%b8%8d%e5%81%9c%e4%b8%8b%e6%95%b4%e4%b8%aa%e4%b8%96%e7%95%8c%ef%bc%9f.md">33 垃圾收集：能否不停下整个世界？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 运行时优化：即时编译的原理和作用.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/34%20%e8%bf%90%e8%a1%8c%e6%97%b6%e4%bc%98%e5%8c%96%ef%bc%9a%e5%8d%b3%e6%97%b6%e7%bc%96%e8%af%91%e7%9a%84%e5%8e%9f%e7%90%86%e5%92%8c%e4%bd%9c%e7%94%a8.md">34 运行时优化：即时编译的原理和作用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 案例总结与热点问题答疑：后端部分真的比前端部分难吗？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/35%20%e6%a1%88%e4%be%8b%e6%80%bb%e7%bb%93%e4%b8%8e%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%9a%e5%90%8e%e7%ab%af%e9%83%a8%e5%88%86%e7%9c%9f%e7%9a%84%e6%af%94%e5%89%8d%e7%ab%af%e9%83%a8%e5%88%86%e9%9a%be%e5%90%97%ef%bc%9f.md">35 案例总结与热点问题答疑：后端部分真的比前端部分难吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 当前技术的发展趋势以及其对编译技术的影响.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/36%20%e5%bd%93%e5%89%8d%e6%8a%80%e6%9c%af%e7%9a%84%e5%8f%91%e5%b1%95%e8%b6%8b%e5%8a%bf%e4%bb%a5%e5%8f%8a%e5%85%b6%e5%af%b9%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%e7%9a%84%e5%bd%b1%e5%93%8d.md">36 当前技术的发展趋势以及其对编译技术的影响.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  云编程：云计算会如何改变编程模式？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/37%20%20%e4%ba%91%e7%bc%96%e7%a8%8b%ef%bc%9a%e4%ba%91%e8%ae%a1%e7%ae%97%e4%bc%9a%e5%a6%82%e4%bd%95%e6%94%b9%e5%8f%98%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%bc%8f%ef%bc%9f.md">37  云编程：云计算会如何改变编程模式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 元编程：一边写程序，一边写语言.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/38%20%e5%85%83%e7%bc%96%e7%a8%8b%ef%bc%9a%e4%b8%80%e8%be%b9%e5%86%99%e7%a8%8b%e5%ba%8f%ef%bc%8c%e4%b8%80%e8%be%b9%e5%86%99%e8%af%ad%e8%a8%80.md">38 元编程：一边写程序，一边写语言.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 汇编代码编程与栈帧管理.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/%e5%8a%a0%e9%a4%90%20%e6%b1%87%e7%bc%96%e4%bb%a3%e7%a0%81%e7%bc%96%e7%a8%8b%e4%b8%8e%e6%a0%88%e5%b8%a7%e7%ae%a1%e7%90%86.md">加餐 汇编代码编程与栈帧管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 因为热爱，所以坚持.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e5%9b%a0%e4%b8%ba%e7%83%ad%e7%88%b1%ef%bc%8c%e6%89%80%e4%bb%a5%e5%9d%9a%e6%8c%81.md">用户故事 因为热爱，所以坚持.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第二季回归 这次，我们一起实战解析真实世界的编译器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/%e7%ac%ac%e4%ba%8c%e5%ad%a3%e5%9b%9e%e5%bd%92%20%e8%bf%99%e6%ac%a1%ef%bc%8c%e6%88%91%e4%bb%ac%e4%b8%80%e8%b5%b7%e5%ae%9e%e6%88%98%e8%a7%a3%e6%9e%90%e7%9c%9f%e5%ae%9e%e4%b8%96%e7%95%8c%e7%9a%84%e7%bc%96%e8%af%91%e5%99%a8.md">第二季回归 这次，我们一起实战解析真实世界的编译器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 用程序语言，推动这个世界的演化.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e4%b9%8b%e7%be%8e/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e7%94%a8%e7%a8%8b%e5%ba%8f%e8%af%ad%e8%a8%80%ef%bc%8c%e6%8e%a8%e5%8a%a8%e8%bf%99%e4%b8%aa%e4%b8%96%e7%95%8c%e7%9a%84%e6%bc%94%e5%8c%96.md">结束语 用程序语言，推动这个世界的演化.md</a>
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
                            <h1 id="title" data-id="25 后端技术的重用：LLVM不仅仅让你高效" class="title">25 后端技术的重用：LLVM不仅仅让你高效</h1>
                            <div><p>在编译器后端，做代码优化和为每个目标平台生成汇编代码，工作量是很大的。那么，有什么办法能降低这方面的工作量，提高我们的工作效率呢？<strong>答案就是利用现成的工具。</strong></p>

<p>在前端部分，我就带你使用Antlr生成了词法分析器和语法分析器。那么在后端部分，我们也可以获得类似的帮助，比如利用LLVM和GCC这两个后端框架。</p>

<p>相比前端的编译器工具，如Lex（Flex）、Yacc（Bison）和Antlr等，对于后端工具，了解的人比较少，资料也更稀缺，如果你是初学者，那么上手的确有一些难度。不过我们已经用20～24讲，铺垫了必要的基础知识，也尝试了手写汇编代码，这些知识足够你学习和掌握后端工具了。</p>

<p>本节课，我想先让你了解一些背景信息，所以会先概要地介绍一下LLVM和GCC这两个有代表性的框架的情况，这样，当我再更加详细地讲解LLVM，带你实际使用一下它的时候，你接受起来就会更加容易了。</p>

<h2 id="两个编译器后端框架-llvm和gcc">两个编译器后端框架：LLVM和GCC</h2>

<p>LLVM是一个开源的编译器基础设施项目，主要聚焦于编译器的后端功能（代码生成、代码优化、JIT……）。它最早是美国伊利诺伊大学的一个研究性项目，核心主持人员是Chris Lattner（克里斯·拉特纳）。</p>

<p>LLVM的出名是由于苹果公司全面采用了这个框架。苹果系统上的C语言、C++、Objective-C的编译器Clang就是基于LLVM的，最新的Swift编程语言也是基于LLVM，支撑了无数的移动应用和桌面应用。无独有偶，在Android平台上最新的开发语言Kotlin，也支持基于LLVM编译成本地代码。</p>

<p>另外，由Mozilla公司（Firefox就是这个公司的产品）开发的系统级编程语言RUST，也是基于LLVM开发的。还有一门相对小众的科学计算领域的语言，叫做Julia，它既能像脚本语言一样灵活易用，又可以具有C语言一样的速度，在数据计算方面又有特别的优化，它的背后也有LLVM的支撑。</p>

<p>OpenGL和一些图像处理领域也在用LLVM，我还看到一个资料，<strong>说阿里云的工程师实现了一个Cava脚本语言，用于配合其搜索引擎系统HA3。</strong></p>

<p><a href="https://en.wikipedia.org/wiki/File:LLVM_Logo.svg" target="_blank">LLVM的logo，一只漂亮的龙：</a></p>

<p><img src="assets/bae220d6d55346b0a9c6cc6504c17b0b.jpg" alt="" /></p>

<p>还有，在人工智能领域炙手可热的TensorFlow框架，在后端也是用LLVM来编译。它把机器学习的IR翻译成LLVM的IR，然后再翻译成支持CPU、GPU和TPU的程序。</p>

<p>所以这样看起来，你所使用的很多语言和工具，背后都有LLVM的影子，只不过你可能没有留意罢了。所以在我看来，要了解编译器的后端技术，就不能不了解LLVM。</p>

<p>与LLVM起到类似作用的后端编译框架是GCC（GNU Compiler Collection，GNU编译器套件）。它支持了GNU Linux上的很多语言，例如C、C++、Objective-C、Fortran、Go语言和Java语言等。其实，它最初只是一个C语言的编译器，后来把公共的后端功能也提炼了出来，形成了框架，支持多种前端语言和后端平台。最近华为发布的方舟编译器，据说也是建立在GCC基础上的。</p>

<p>LLVM和GCC很难比较优劣，因为这两个项目都取得了很大的成功。</p>

<p>在本课程中，我们主要采用LLVM，但其中学到的一些知识，比如IR的设计、代码优化算法、适配不同硬件的策略，在学习GCC或其他编译器后端的时候，也是有用的，从而大大提升学习效率。</p>

<p>接下来，我们先来看看LLVM的构成和特点，让你对它有个宏观的认识。</p>

<h2 id="了解llvm的特点">了解LLVM的特点</h2>

<p>LLVM能够支持多种语言的前端、多种后端CPU架构。在LLVM内部，使用类型化的和SSA特点的IR进行各种分析、优化和转换：</p>

<p><img src="assets/6b03f1e0504c4d5686e87053821881bc.jpg" alt="" /></p>

<p>LLVM项目包含了很多组成部分：</p>

<ul>
<li><p>LLVM核心（core）。就是上图中的优化和分析工具，还包括了为各种CPU生成目标代码的功能；这些库采用的是LLVM IR，一个良好定义的中间语言，在上一讲，我们已经初步了解它了。</p></li>

<li><p>Clang前端（是基于LLVM的C、C++、Objective-C编译器）。</p></li>

<li><p>LLDB（一个调试工具）。</p></li>

<li><p>LLVM版本的C++标准类库。</p></li>

<li><p>其他一些子项目。</p></li>
</ul>

<p><strong>我个人很喜欢LLVM，想了想，主要有几点原因：</strong></p>

<p>首先，LLVM有良好的模块化设计和接口。以前的编译器后端技术很难复用，而LLVM具备定义了良好接口的库，方便使用者选择在什么时候，复用哪些后端功能。比如，针对代码优化，LLVM提供了很多算法，语言的设计者可以自己选择合适的算法，或者实现自己特殊的算法，具有很好的灵活性。</p>

<p>第二，LLVM同时支持JIT（即时编译）和AOT（提前编译）两种模式。过去的语言要么是解释型的，要么编译后运行。习惯了使用解释型语言的程序员，很难习惯必须等待一段编译时间才能看到运行效果。很多科学工作者，习惯在一个REPL界面中一边写脚本，一边实时看到反馈。LLVM既可以通过JIT技术支持解释执行，又可以完全编译后才执行，这对于语言的设计者很有吸引力。</p>

<p>第三，有很多可以学习借鉴的项目。Swift、Rust、Julia这些新生代的语言，实现了很多吸引人的特性，还有很多其他的开源项目，而我们可以研究、借鉴它们是如何充分利用LLVM的。</p>

<p>第四，全过程优化的设计思想。LLVM在设计上支持全过程的优化。Lattner和Adve最早关于LLVM设计思想的文章<a href="https://llvm.org/pubs/2003-09-30-LifelongOptimizationTR.pdf" target="_blank">《LLVM: 一个全生命周期分析和转换的编译框架》，</a>就提出计算机语言可以在各个阶段进行优化，包括编译时、链接时、安装时，甚至是运行时。</p>

<p>以运行时优化为例，基于LLVM我们能够在运行时，收集一些性能相关的数据对代码编译优化，可以是实时优化的、动态修改内存中的机器码；也可以收集这些性能数据，然后做离线的优化，重新生成可执行文件，然后再加载执行，<strong>这一点非常吸引我，</strong>因为在现代计算环境下，每种功能的计算特点都不相同，确实需要针对不同的场景做不同的优化。下图展现了这个过程（图片来源《 LLVM: A Compilation Framework for Lifelong Program Analysis &amp; Transformation》）：</p>

<p><img src="assets/60cbd3bf31ed4a9f956161116451785f.jpg" alt="" /></p>

<p><strong>我建议你读一读Lattner和Adve的这篇论文</strong>（另外强调一下，当你深入学习编译技术的时候，阅读领域内的论文就是必不可少的一项功课了）。</p>

<p>第五，LLVM的授权更友好。GNU的很多软件都是采用GPL协议的，所以如果用GCC的后端工具来编写你的语言，你可能必须要按照GPL协议开源。而LLVM则更友好一些，你基于LLVM所做的工作，完全可以是闭源的软件产品。</p>

<p>而我之所以说：“LLVM不仅仅让你更高效”，就是因为上面它的这些特点。</p>

<p>现在，你已经对LLVM的构成和特点有一定的了解了，接下来，我带你亲自动手操作和体验一下LLVM的功能，这样你就可以迅速消除对它的陌生感，快速上手了。</p>

<h2 id="体验一下llvm的功能">体验一下LLVM的功能</h2>

<p>首先你需要安装一下LLVM（参照<a href="http://releases.llvm.org/" target="_blank">官方网站</a>上的相关介绍下载安装）。因为我使用的是macOS，所以用brew就可以安装。</p>

<pre><code>brew install llvm
</code></pre>

<p>因为LLVM里面带了一个版本的Clang和C++的标准库，与本机原来的工具链可能会有冲突，所以brew安装的时候并没有在/usr/local下建立符号链接。你在用LLVM工具的时候，要配置好相关的环境变量。</p>

<pre><code># 可执行文件的路径
export PATH=&quot;/usr/local/opt/llvm/bin:$PATH&quot;
# 让编译器能够找到LLVM
export LDFLAGS=&quot;-L/usr/local/opt/llvm/lib&quot;
export CPPFLAGS=&quot;-I/usr/local/opt/llvm/include”
</code></pre>

<p>安装完毕之后，我们使用一下LLVM自带的命令行工具，分几步体验一下LLVM的功能：</p>

<p>1.从C语言代码生成IR；-
2.优化IR；-
3.从文本格式的IR生成二进制的字节码；-
4.把IR编译成汇编代码和可执行文件。</p>

<p>从C语言代码生成IR代码比较简单，上一讲中我们已经用到过一个C语言的示例代码：</p>

<pre><code>//fun1.c 
int fun1(int a, int b){
    int c = 10;
    return a+b+c;
}
</code></pre>

<p>用前端工具Clang就可以把它编译成IR代码：</p>

<pre><code>clang -emit-llvm -S fun1.c -o fun1.ll
</code></pre>

<p>其中，-emit-llvm参数告诉Clang生成LLVM的汇编码，也就是IR代码（如果不带这个参数，就会生成针对目标机器的汇编码）所生成的IR我们上一讲也见过，你现在应该能够读懂它了。你可以多写几个不同的程序，看看生成的IR是什么样的，比如if语句、循环语句等等（这时你完成了第一步）：</p>

<pre><code>; ModuleID = 'function-call1.c'
source_filename = &quot;function-call1.c&quot;
target datalayout = &quot;e-m:o-i64:64-f80:128-n8:16:32:64-S128&quot;
target triple = &quot;x86_64-apple-macosx10.14.0&quot;

; Function Attrs: noinline nounwind optnone ssp uwtable
define i32 @fun1(i32, i32) #0 {
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  store i32 %0, i32* %3, align 4
  store i32 %1, i32* %4, align 4
  store i32 10, i32* %5, align 4
  %6 = load i32, i32* %3, align 4
  %7 = load i32, i32* %4, align 4
  %8 = add nsw i32 %6, %7
  %9 = load i32, i32* %5, align 4
  %10 = add nsw i32 %8, %9
  ret i32 %10
}

attributes #0 = { noinline nounwind optnone ssp uwtable &quot;correctly-rounded-divide-sqrt-fp-math&quot;=&quot;false&quot; &quot;disable-tail-calls&quot;=&quot;false&quot; &quot;less-precise-fpmad&quot;=&quot;false&quot; &quot;min-legal-vector-width&quot;=&quot;0&quot; &quot;no-frame-pointer-elim&quot;=&quot;true&quot; &quot;no-frame-pointer-elim-non-leaf&quot; &quot;no-infs-fp-math&quot;=&quot;false&quot; &quot;no-jump-tables&quot;=&quot;false&quot; &quot;no-nans-fp-math&quot;=&quot;false&quot; &quot;no-signed-zeros-fp-math&quot;=&quot;false&quot; &quot;no-trapping-math&quot;=&quot;false&quot; &quot;stack-protector-buffer-size&quot;=&quot;8&quot; &quot;target-cpu&quot;=&quot;penryn&quot; &quot;target-features&quot;=&quot;+cx16,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87&quot; &quot;unsafe-fp-math&quot;=&quot;false&quot; &quot;use-soft-float&quot;=&quot;false&quot; }

!llvm.module.flags = !{!0, !1}
!llvm.ident = !{!2}

!0 = !{i32 1, !&quot;wchar_size&quot;, i32 4}
!1 = !{i32 7, !&quot;PIC Level&quot;, i32 2}
!2 = !{!&quot;clang version 8.0.0 (tags/RELEASE_800/final)&quot;}
</code></pre>

<p>上一讲我们提到过，可以对生成的IR做优化，让代码更短，你只要在上面的命令中加上-O2参数就可以了（这时你完成了第二步）：</p>

<pre><code>clang -emit-llvm -S -O2 fun1.c -o fun1.ll
</code></pre>

<p>这个时候，函数体的核心代码就变短了很多。这里面最重要的优化动作，是从原来使用内存（alloca指令是在栈中分配空间，store指令是往内存里写入值），优化到只使用寄存器（%0、%1是参数，%3、%4也是寄存器）。</p>

<pre><code>define i32 @fun1(i32, i32) #0 {
  %3 = add nsw i32 %0, %1
  %4 = add nsw i32 %3, 10
  ret i32 %4
}
</code></pre>

<p>你还可以用opt命令来完成上面的优化，具体我们在27、28讲中讲优化算法的时候再细化。</p>

<p><strong>另外，LLVM的IR有两种格式。</strong>在示例代码中显示的，是它的文本格式，文件名一般以.ll结尾。第二种格式是字节码（bitcode）格式，文件名以.bc结尾。<strong>为什么要用两种格式呢？</strong>因为文本格式的文件便于程序员阅读，而字节码格式的是二进制文件，便于机器处理，比如即时编译和执行。生成字节码格式之后，所占空间会小很多，所以可以快速加载进内存，并转换为内存中的对象格式。而如果加载文本文件，则还需要一个解析的过程，才能变成内存中的格式，效率比较慢。</p>

<p>调用llvm-as命令，我们可以把文本格式转换成字节码格式：</p>

<pre><code>llvm-as fun1.ll -o fun1.bc
</code></pre>

<p>我们也可以用clang直接生成字节码，这时不需要带-S参数，而是要用-c参数。</p>

<pre><code>clang -emit-llvm -c fun1.c -o fun1.bc
</code></pre>

<p>因为.bc文件是二进制文件，不能直接用文本编辑器查看，而要用hexdump命令查看（这时你完成了第三步）：</p>

<pre><code>hexdump -C fun1.bc
</code></pre>

<p><img src="assets/22d90fe73704407c9be7c56c266711eb.jpg" alt="" /></p>

<p>LLVM的一个优点，就是可以即时编译运行字节码，不一定非要编译生成汇编码和可执行文件才能运行（这一点有点儿像Java语言），这也让LLVM具有极高的灵活性，比如，可以在运行时根据收集的性能信息，改变优化策略，生成更高效的机器码。</p>

<p>再进一步，我们可以把字节码编译成目标平台的汇编代码。我们使用的是llc命令，命令如下：</p>

<pre><code>llc fun1.bc -o fun1.s
</code></pre>

<p>用clang命令也能从字节码生成汇编代码，要注意带上-S参数就行了：</p>

<pre><code>clang -S fun1.bc -o fun1.s
</code></pre>

<p><strong>到了这一步，我们已经得到了汇编代码，</strong>接着就可以进一步生成目标文件和可执行文件了。</p>

<p>实际上，使用LLVM从源代码到生成可执行文件有两条可能的路径：</p>

<p><img src="assets/233e7f1409e643c9acd4be00b7ad2149.jpg" alt="" /></p>

<ul>
<li><p>第一条路径，是把每个源文件分别编译成字节码文件，然后再编译成目标文件，最后链接成可执行文件。</p></li>

<li><p>第二条路径，是先把编译好的字节码文件链接在一起，形成一个更大的字节码文件，然后对这个字节码文件进行进一步的优化，之后再生成目标文件和可执行文件。</p></li>
</ul>

<p>第二条路径比第一条路径多了一个优化的步骤，第一条路径只对每个模块做了优化，没有做整体的优化。所以，如有可能，尽量采用第二条路径，这样能够生成更加优化的代码。</p>

<p>现在你完成了第四步，对LLVM的命令行工具有了一定的了解。总结一下，我们用到的命令行工具包括：clang前端、llvm-as、llc，其他命令还有opt（代码优化）、llvm-dis（将字节码再反编译回ll文件）、llvm-link（链接）等，你可以看它们的help信息，并练习使用。</p>

<p>在熟悉了命令行工具之后，我们就可以进一步在编程环境中使用LLVM了，不过在此之前，需要搭建一个开发环境。</p>

<h2 id="建立c-开发环境来使用llvm">建立C++开发环境来使用LLVM</h2>

<p>LLVM本身是用C++开发的，所以最好采用C++调用它的功能。当然，采用其他语言也有办法调用LLVM：</p>

<ul>
<li>C语言可以调用专门的C接口；</li>
<li>像Go、Rust、Python、Ocaml、甚至Node.js都有对LLVM API的绑定；</li>
<li>如果使用Java，也可以通过JavaCPP（类似JNI）技术调用LLVM。</li>
</ul>

<p>在课程中，我用C++来做实现，因为这样能够最近距离地跟LLVM打交道。与此同时，我们前端工具采用的Antlr，也能够支持C++开发环境。<strong>所以，我为playscript建立了一个C++的开发环境。</strong></p>

<p><strong>开发工具方面：</strong>原则上只要一个编辑器加上工具链就行，但为了提高效率，有IDE的支持会更好（我用的是JetBrains的Clion）。</p>

<p><strong>构建工具方面：</strong>目前LLVM本身用的是CMake，而Clion刚好也采用CMake，所以很方便。</p>

<p><strong>这里我想针对CMake多解释几句，</strong>因为越来越多的C++项目都是用CMake来管理的，LLVM以及Antlr的C++版本也采用了CMake，<strong>你最好对它有一定了解。</strong></p>

<p>CMake是一款优秀的工程构建工具，它类似于Java程序员们习惯使用的Maven工具。对于只包含少量文件或模块的C或C++程序，你可以仅仅通过命令行带上一些参数就能编译。</p>

<p>不过，实际的项目都会比较复杂，往往会包含比较多的模块，存在比较复杂的依赖关系，编译过程也不是一步能完成的，要分成多步。这时候我们一般用make管理项目的构建过程，这就要学会写make文件。但手工写make文件工作量会比较大，而CMake就是在make的基础上再封装了一层，它能通过更简单的配置文件，帮我们生成make文件，帮助程序员提升效率。</p>

<p>整个开发环境的搭建我在课程里就不多写了，你可以参见示例代码所附带的文档。文档里有比较清晰的说明，可以帮助你把环境搭建起来，并运行示例程序。</p>

<p>另外，我知道你可能对C++并不那么熟悉。但你应该学过C语言，所以示例代码还是能看懂的。</p>

<h2 id="课程小结">课程小结</h2>

<p>本节课，为了帮助你理解后端工具，我先概要介绍了后端工具的情况，接着着重介绍了LLVM的构成和特点，然后又带你熟悉了它的命令行工具，让你能够生成文本和字节码两种格式的IR，并生成可执行文件，最后带你了解了LLVM的开发环境。</p>

<p>本节课的内容比较好理解，因为侧重让你建立跟LLVM的熟悉感，没有什么复杂的算法和原理，而我想强调的是以下几点：</p>

<p>1.后端工具对于语言设计者很重要，我们必须学会善加利用；-
2.LLVM有很好的模块化设计，支持即时编译（JIT）和提前编译（AOT），支持全过程的优化，并且具备友好的授权，值得我们好好掌握；-
3.你要熟悉LLVM的命令行工具，这样可以上手做很多实验，加深对LLVM的了解。</p>

<p>最后，我想给你的建议是：一定要动手安装和使用LLVM，写点代码测试它的功能。比如，写点儿C、C++等语言的程序，并翻译成IR，进一步熟悉LLVM的IR。下一讲，我们就要进入它的内部，调用它的API来生成IR和运行了！</p>

<h2 id="一课一思">一课一思</h2>

<p>很多语言都获得了后端工具的帮助，比如可以把Android应用直接编译成机器码，提升运行效率。你所经常使用的计算机语言采用了什么后端工具？有什么特点？欢迎在留言区分享。</p>

<p>最后，感谢你的阅读，如果这篇文章让你有所收获，也欢迎你分享给更多的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#95f9f9f9aca1a4a4a5a2d5f2f8f4fcf9bbf6faf8" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f17fff0cef548cd',t:'MTczNDExNDM1OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
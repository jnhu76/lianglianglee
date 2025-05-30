<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;Java&#32;JIT编译器（四）：Graal的后端是如何工作的？>
        <link rel="icon" href="/static/favicon.png">
        <title>16 Java JIT编译器（四）：Graal的后端是如何工作的？ </title>
        
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
                        <a class="menu-item" id="00 学习指南 如何学习这门编译原理实战课？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%ad%a6%e4%b9%a0%e6%8c%87%e5%8d%97%20%e5%a6%82%e4%bd%95%e5%ad%a6%e4%b9%a0%e8%bf%99%e9%97%a8%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be%ef%bc%9f.md">00 学习指南 如何学习这门编译原理实战课？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 在真实世界的编译器中游历.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%9c%a8%e7%9c%9f%e5%ae%9e%e4%b8%96%e7%95%8c%e7%9a%84%e7%bc%96%e8%af%91%e5%99%a8%e4%b8%ad%e6%b8%b8%e5%8e%86.md">00 开篇词 在真实世界的编译器中游历.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 编译的全过程都悄悄做了哪些事情？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/01%20%e7%bc%96%e8%af%91%e7%9a%84%e5%85%a8%e8%bf%87%e7%a8%8b%e9%83%bd%e6%82%84%e6%82%84%e5%81%9a%e4%ba%86%e5%93%aa%e4%ba%9b%e4%ba%8b%e6%83%85%ef%bc%9f.md">01 编译的全过程都悄悄做了哪些事情？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 词法分析：用两种方式构造有限自动机.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/02%20%e8%af%8d%e6%b3%95%e5%88%86%e6%9e%90%ef%bc%9a%e7%94%a8%e4%b8%a4%e7%a7%8d%e6%96%b9%e5%bc%8f%e6%9e%84%e9%80%a0%e6%9c%89%e9%99%90%e8%87%aa%e5%8a%a8%e6%9c%ba.md">02 词法分析：用两种方式构造有限自动机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 语法分析：两个基本功和两种算法思路.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/03%20%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%ef%bc%9a%e4%b8%a4%e4%b8%aa%e5%9f%ba%e6%9c%ac%e5%8a%9f%e5%92%8c%e4%b8%a4%e7%a7%8d%e7%ae%97%e6%b3%95%e6%80%9d%e8%b7%af.md">03 语法分析：两个基本功和两种算法思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 语义分析：让程序符合语义规则.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/04%20%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90%ef%bc%9a%e8%ae%a9%e7%a8%8b%e5%ba%8f%e7%ac%a6%e5%90%88%e8%af%ad%e4%b9%89%e8%a7%84%e5%88%99.md">04 语义分析：让程序符合语义规则.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 运行时机制：程序如何运行，你有发言权.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/05%20%e8%bf%90%e8%a1%8c%e6%97%b6%e6%9c%ba%e5%88%b6%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%a6%82%e4%bd%95%e8%bf%90%e8%a1%8c%ef%bc%8c%e4%bd%a0%e6%9c%89%e5%8f%91%e8%a8%80%e6%9d%83.md">05 运行时机制：程序如何运行，你有发言权.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 中间代码：不是只有一副面孔.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/06%20%e4%b8%ad%e9%97%b4%e4%bb%a3%e7%a0%81%ef%bc%9a%e4%b8%8d%e6%98%af%e5%8f%aa%e6%9c%89%e4%b8%80%e5%89%af%e9%9d%a2%e5%ad%94.md">06 中间代码：不是只有一副面孔.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 代码优化：跟编译器做朋友，让你的代码飞起来.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/07%20%e4%bb%a3%e7%a0%81%e4%bc%98%e5%8c%96%ef%bc%9a%e8%b7%9f%e7%bc%96%e8%af%91%e5%99%a8%e5%81%9a%e6%9c%8b%e5%8f%8b%ef%bc%8c%e8%ae%a9%e4%bd%a0%e7%9a%84%e4%bb%a3%e7%a0%81%e9%a3%9e%e8%b5%b7%e6%9d%a5.md">07 代码优化：跟编译器做朋友，让你的代码飞起来.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 代码生成：如何实现机器相关的优化？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/08%20%e4%bb%a3%e7%a0%81%e7%94%9f%e6%88%90%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%99%a8%e7%9b%b8%e5%85%b3%e7%9a%84%e4%bc%98%e5%8c%96%ef%bc%9f.md">08 代码生成：如何实现机器相关的优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 Java编译器（一）：手写的编译器有什么优势？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/09%20Java%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%89%8b%e5%86%99%e7%9a%84%e7%bc%96%e8%af%91%e5%99%a8%e6%9c%89%e4%bb%80%e4%b9%88%e4%bc%98%e5%8a%bf%ef%bc%9f.md">09 Java编译器（一）：手写的编译器有什么优势？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Java编译器（二）：语法分析之后，还要做些什么？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/10%20Java%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e8%af%ad%e6%b3%95%e5%88%86%e6%9e%90%e4%b9%8b%e5%90%8e%ef%bc%8c%e8%bf%98%e8%a6%81%e5%81%9a%e4%ba%9b%e4%bb%80%e4%b9%88%ef%bc%9f.md">10 Java编译器（二）：语法分析之后，还要做些什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Java编译器（三）：属性分析和数据流分析.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/11%20Java%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%b1%9e%e6%80%a7%e5%88%86%e6%9e%90%e5%92%8c%e6%95%b0%e6%8d%ae%e6%b5%81%e5%88%86%e6%9e%90.md">11 Java编译器（三）：属性分析和数据流分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 Java编译器（四）：去除语法糖和生成字节码.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/12%20Java%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e5%8e%bb%e9%99%a4%e8%af%ad%e6%b3%95%e7%b3%96%e5%92%8c%e7%94%9f%e6%88%90%e5%ad%97%e8%8a%82%e7%a0%81.md">12 Java编译器（四）：去除语法糖和生成字节码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Java JIT编译器（一）：动手修改Graal编译器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/13%20Java%20JIT%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%8a%a8%e6%89%8b%e4%bf%ae%e6%94%b9Graal%e7%bc%96%e8%af%91%e5%99%a8.md">13 Java JIT编译器（一）：动手修改Graal编译器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Java JIT编译器（二）：Sea of Nodes为何如此强大？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/14%20Java%20JIT%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aSea%20of%20Nodes%e4%b8%ba%e4%bd%95%e5%a6%82%e6%ad%a4%e5%bc%ba%e5%a4%a7%ef%bc%9f.md">14 Java JIT编译器（二）：Sea of Nodes为何如此强大？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Java JIT编译器（三）：探究内联和逃逸分析的算法原理.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/15%20Java%20JIT%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%8e%a2%e7%a9%b6%e5%86%85%e8%81%94%e5%92%8c%e9%80%83%e9%80%b8%e5%88%86%e6%9e%90%e7%9a%84%e7%ae%97%e6%b3%95%e5%8e%9f%e7%90%86.md">15 Java JIT编译器（三）：探究内联和逃逸分析的算法原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Java JIT编译器（四）：Graal的后端是如何工作的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/16%20Java%20JIT%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9aGraal%e7%9a%84%e5%90%8e%e7%ab%af%e6%98%af%e5%a6%82%e4%bd%95%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%9f.md">16 Java JIT编译器（四）：Graal的后端是如何工作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 Python编译器（一）：如何用工具生成编译器？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/17%20Python%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e5%b7%a5%e5%85%b7%e7%94%9f%e6%88%90%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%9f.md">17 Python编译器（一）：如何用工具生成编译器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Python编译器（二）：从AST到字节码.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/18%20Python%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e4%bb%8eAST%e5%88%b0%e5%ad%97%e8%8a%82%e7%a0%81.md">18 Python编译器（二）：从AST到字节码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Python编译器（三）：运行时机制.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/19%20Python%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e8%bf%90%e8%a1%8c%e6%97%b6%e6%9c%ba%e5%88%b6.md">19 Python编译器（三）：运行时机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 JavaScript编译器（一）：V8的解析和编译过程.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/20%20JavaScript%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aV8%e7%9a%84%e8%a7%a3%e6%9e%90%e5%92%8c%e7%bc%96%e8%af%91%e8%bf%87%e7%a8%8b.md">20 JavaScript编译器（一）：V8的解析和编译过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 JavaScript编译器（二）：V8的解释器和优化编译器.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/21%20JavaScript%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aV8%e7%9a%84%e8%a7%a3%e9%87%8a%e5%99%a8%e5%92%8c%e4%bc%98%e5%8c%96%e7%bc%96%e8%af%91%e5%99%a8.md">21 JavaScript编译器（二）：V8的解释器和优化编译器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Julia编译器（一）：如何让动态语言性能很高？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/22%20Julia%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e5%8a%a8%e6%80%81%e8%af%ad%e8%a8%80%e6%80%a7%e8%83%bd%e5%be%88%e9%ab%98%ef%bc%9f.md">22 Julia编译器（一）：如何让动态语言性能很高？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Julia编译器（二）：如何利用LLVM的优化和后端功能？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/23%20Julia%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8LLVM%e7%9a%84%e4%bc%98%e5%8c%96%e5%92%8c%e5%90%8e%e7%ab%af%e5%8a%9f%e8%83%bd%ef%bc%9f.md">23 Julia编译器（二）：如何利用LLVM的优化和后端功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 Go语言编译器：把它当作教科书吧.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/24%20Go%e8%af%ad%e8%a8%80%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%9a%e6%8a%8a%e5%ae%83%e5%bd%93%e4%bd%9c%e6%95%99%e7%a7%91%e4%b9%a6%e5%90%a7.md">24 Go语言编译器：把它当作教科书吧.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 MySQL编译器（一）：解析一条SQL语句的执行过程.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/25%20MySQL%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e8%a7%a3%e6%9e%90%e4%b8%80%e6%9d%a1SQL%e8%af%ad%e5%8f%a5%e7%9a%84%e6%89%a7%e8%a1%8c%e8%bf%87%e7%a8%8b.md">25 MySQL编译器（一）：解析一条SQL语句的执行过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 MySQL编译器（二）：编译技术如何帮你提升数据库性能？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/26%20MySQL%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%e5%a6%82%e4%bd%95%e5%b8%ae%e4%bd%a0%e6%8f%90%e5%8d%87%e6%95%b0%e6%8d%ae%e5%ba%93%e6%80%a7%e8%83%bd%ef%bc%9f.md">26 MySQL编译器（二）：编译技术如何帮你提升数据库性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 课前导读：学习现代语言设计的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/27%20%e8%af%be%e5%89%8d%e5%af%bc%e8%af%bb%ef%bc%9a%e5%ad%a6%e4%b9%a0%e7%8e%b0%e4%bb%a3%e8%af%ad%e8%a8%80%e8%ae%be%e8%ae%a1%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">27 课前导读：学习现代语言设计的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 前端总结：语言设计也有人机工程学.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/28%20%e5%89%8d%e7%ab%af%e6%80%bb%e7%bb%93%ef%bc%9a%e8%af%ad%e8%a8%80%e8%ae%be%e8%ae%a1%e4%b9%9f%e6%9c%89%e4%ba%ba%e6%9c%ba%e5%b7%a5%e7%a8%8b%e5%ad%a6.md">28 前端总结：语言设计也有人机工程学.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 中端总结：不遗余力地进行代码优化.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/29%20%e4%b8%ad%e7%ab%af%e6%80%bb%e7%bb%93%ef%bc%9a%e4%b8%8d%e9%81%97%e4%bd%99%e5%8a%9b%e5%9c%b0%e8%bf%9b%e8%a1%8c%e4%bb%a3%e7%a0%81%e4%bc%98%e5%8c%96.md">29 中端总结：不遗余力地进行代码优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 后端总结：充分发挥硬件的能力.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/30%20%e5%90%8e%e7%ab%af%e6%80%bb%e7%bb%93%ef%bc%9a%e5%85%85%e5%88%86%e5%8f%91%e6%8c%a5%e7%a1%ac%e4%bb%b6%e7%9a%84%e8%83%bd%e5%8a%9b.md">30 后端总结：充分发挥硬件的能力.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 运行时（一）：从0到语言级的虚拟化.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/31%20%e8%bf%90%e8%a1%8c%e6%97%b6%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e4%bb%8e0%e5%88%b0%e8%af%ad%e8%a8%80%e7%ba%a7%e7%9a%84%e8%99%9a%e6%8b%9f%e5%8c%96.md">31 运行时（一）：从0到语言级的虚拟化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 运行时（二）：垃圾收集与语言的特性有关吗？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/32%20%e8%bf%90%e8%a1%8c%e6%97%b6%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%9e%83%e5%9c%be%e6%94%b6%e9%9b%86%e4%b8%8e%e8%af%ad%e8%a8%80%e7%9a%84%e7%89%b9%e6%80%a7%e6%9c%89%e5%85%b3%e5%90%97%ef%bc%9f.md">32 运行时（二）：垃圾收集与语言的特性有关吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 并发中的编译技术（一）：如何从语言层面支持线程？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/33%20%e5%b9%b6%e5%8f%91%e4%b8%ad%e7%9a%84%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e8%af%ad%e8%a8%80%e5%b1%82%e9%9d%a2%e6%94%af%e6%8c%81%e7%ba%bf%e7%a8%8b%ef%bc%9f.md">33 并发中的编译技术（一）：如何从语言层面支持线程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 并发中的编译技术（二）：如何从语言层面支持协程？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/34%20%e5%b9%b6%e5%8f%91%e4%b8%ad%e7%9a%84%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e8%af%ad%e8%a8%80%e5%b1%82%e9%9d%a2%e6%94%af%e6%8c%81%e5%8d%8f%e7%a8%8b%ef%bc%9f.md">34 并发中的编译技术（二）：如何从语言层面支持协程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 并发中的编译技术（三）：Erlang语言厉害在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/35%20%e5%b9%b6%e5%8f%91%e4%b8%ad%e7%9a%84%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9aErlang%e8%af%ad%e8%a8%80%e5%8e%89%e5%ae%b3%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">35 并发中的编译技术（三）：Erlang语言厉害在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 高级特性（一）：揭秘元编程的实现机制.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/36%20%e9%ab%98%e7%ba%a7%e7%89%b9%e6%80%a7%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%8f%ad%e7%a7%98%e5%85%83%e7%bc%96%e7%a8%8b%e7%9a%84%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6.md">36 高级特性（一）：揭秘元编程的实现机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 高级特性（二）：揭秘泛型编程的实现机制.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/37%20%e9%ab%98%e7%ba%a7%e7%89%b9%e6%80%a7%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%8f%ad%e7%a7%98%e6%b3%9b%e5%9e%8b%e7%bc%96%e7%a8%8b%e7%9a%84%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6.md">37 高级特性（二）：揭秘泛型编程的实现机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 综合实现（一）：如何实现面向对象编程？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/38%20%e7%bb%bc%e5%90%88%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e7%bc%96%e7%a8%8b%ef%bc%9f.md">38 综合实现（一）：如何实现面向对象编程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 综合实现（二）：如何实现函数式编程？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/39%20%e7%bb%bc%e5%90%88%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%87%bd%e6%95%b0%e5%bc%8f%e7%bc%96%e7%a8%8b%ef%bc%9f.md">39 综合实现（二）：如何实现函数式编程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 成果检验：方舟编译器的优势在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/40%20%e6%88%90%e6%9e%9c%e6%a3%80%e9%aa%8c%ef%bc%9a%e6%96%b9%e8%88%9f%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e4%bc%98%e5%8a%bf%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">40 成果检验：方舟编译器的优势在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不定期加餐1 远程办公，需要你我具备什么样的素质？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e4%b8%8d%e5%ae%9a%e6%9c%9f%e5%8a%a0%e9%a4%901%20%e8%bf%9c%e7%a8%8b%e5%8a%9e%e5%85%ac%ef%bc%8c%e9%9c%80%e8%a6%81%e4%bd%a0%e6%88%91%e5%85%b7%e5%a4%87%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e7%b4%a0%e8%b4%a8%ef%bc%9f.md">不定期加餐1 远程办公，需要你我具备什么样的素质？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不定期加餐2 学习技术的过程，其实是训练心理素质的过程.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e4%b8%8d%e5%ae%9a%e6%9c%9f%e5%8a%a0%e9%a4%902%20%e5%ad%a6%e4%b9%a0%e6%8a%80%e6%9c%af%e7%9a%84%e8%bf%87%e7%a8%8b%ef%bc%8c%e5%85%b6%e5%ae%9e%e6%98%af%e8%ae%ad%e7%bb%83%e5%bf%83%e7%90%86%e7%b4%a0%e8%b4%a8%e7%9a%84%e8%bf%87%e7%a8%8b.md">不定期加餐2 学习技术的过程，其实是训练心理素质的过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不定期加餐3 这几年，打动我的两本好书.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e4%b8%8d%e5%ae%9a%e6%9c%9f%e5%8a%a0%e9%a4%903%20%e8%bf%99%e5%87%a0%e5%b9%b4%ef%bc%8c%e6%89%93%e5%8a%a8%e6%88%91%e7%9a%84%e4%b8%a4%e6%9c%ac%e5%a5%bd%e4%b9%a6.md">不定期加餐3 这几年，打动我的两本好书.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不定期加餐4 从身边的牛人身上，我学到的一些优秀品质.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e4%b8%8d%e5%ae%9a%e6%9c%9f%e5%8a%a0%e9%a4%904%20%e4%bb%8e%e8%ba%ab%e8%be%b9%e7%9a%84%e7%89%9b%e4%ba%ba%e8%ba%ab%e4%b8%8a%ef%bc%8c%e6%88%91%e5%ad%a6%e5%88%b0%e7%9a%84%e4%b8%80%e4%ba%9b%e4%bc%98%e7%a7%80%e5%93%81%e8%b4%a8.md">不定期加餐4 从身边的牛人身上，我学到的一些优秀品质.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="不定期加餐5 借助实例，探究C&#43;&#43;编译器的内部机制.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e4%b8%8d%e5%ae%9a%e6%9c%9f%e5%8a%a0%e9%a4%905%20%e5%80%9f%e5%8a%a9%e5%ae%9e%e4%be%8b%ef%bc%8c%e6%8e%a2%e7%a9%b6C&#43;&#43;%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e5%86%85%e9%83%a8%e6%9c%ba%e5%88%b6.md">不定期加餐5 借助实例，探究C&#43;&#43;编译器的内部机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 7种编译器的核心概念与算法.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e5%88%92%e9%87%8d%e7%82%b9%207%e7%a7%8d%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e6%a0%b8%e5%bf%83%e6%a6%82%e5%bf%b5%e4%b8%8e%e7%ae%97%e6%b3%95.md">划重点 7种编译器的核心概念与算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="期末答疑与总结 再次审视学习编译原理的作用.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e6%9c%9f%e6%9c%ab%e7%ad%94%e7%96%91%e4%b8%8e%e6%80%bb%e7%bb%93%20%e5%86%8d%e6%ac%a1%e5%ae%a1%e8%a7%86%e5%ad%a6%e4%b9%a0%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e7%9a%84%e4%bd%9c%e7%94%a8.md">期末答疑与总结 再次审视学习编译原理的作用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="热点问题答疑 如何吃透7种真实的编译器？.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%20%e5%a6%82%e4%bd%95%e5%90%83%e9%80%8f7%e7%a7%8d%e7%9c%9f%e5%ae%9e%e7%9a%84%e7%bc%96%e8%af%91%e5%99%a8%ef%bc%9f.md">热点问题答疑 如何吃透7种真实的编译器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 易昊：程序员不止有Bug和加班，还有诗和远方.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%98%93%e6%98%8a%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%91%98%e4%b8%8d%e6%ad%a2%e6%9c%89Bug%e5%92%8c%e5%8a%a0%e7%8f%ad%ef%bc%8c%e8%bf%98%e6%9c%89%e8%af%97%e5%92%8c%e8%bf%9c%e6%96%b9.md">用户故事 易昊：程序员不止有Bug和加班，还有诗和远方.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="知识地图 一起来复习编译技术核心概念与算法.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e7%9f%a5%e8%af%86%e5%9c%b0%e5%9b%be%20%e4%b8%80%e8%b5%b7%e6%9d%a5%e5%a4%8d%e4%b9%a0%e7%bc%96%e8%af%91%e6%8a%80%e6%9c%af%e6%a0%b8%e5%bf%83%e6%a6%82%e5%bf%b5%e4%b8%8e%e7%ae%97%e6%b3%95.md">知识地图 一起来复习编译技术核心概念与算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 实战是唯一标准！.md" href="/%e4%b8%93%e6%a0%8f/%e7%bc%96%e8%af%91%e5%8e%9f%e7%90%86%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%ae%9e%e6%88%98%e6%98%af%e5%94%af%e4%b8%80%e6%a0%87%e5%87%86%ef%bc%81.md">结束语 实战是唯一标准！.md</a>
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
                            <h1 id="title" data-id="16 Java JIT编译器（四）：Graal的后端是如何工作的？" class="title">16 Java JIT编译器（四）：Graal的后端是如何工作的？</h1>
                            <div><p>你好，我是宫文学。</p>

<p>前面两讲中，我介绍了Sea of Nodes类型的HIR，以及基于HIR的各种分析处理，这可以看做是编译器的中端。</p>

<p>可编译器最终还是要生成机器码的。那么，这个过程是怎么实现的呢？与硬件架构相关的LIR是什么样子的呢？指令选择是怎么做的呢？</p>

<p>这一讲，我就带你了解Graal编译器的后端功能，回答以上这些问题，破除你对后端处理过程的神秘感。</p>

<p>首先，我们来直观地了解一下后端处理的流程。</p>

<h2 id="后端的处理流程">后端的处理流程</h2>

<p>在<a href="https://time.geekbang.org/column/article/256914" target="_blank">第14讲</a>中，我们在运行Java示例程序的时候（比如<code>atLeastTen()</code>方法），使用了“<code>-Dgraal.Dump=:5</code>”的选项，这个选项会dump出整个编译过程最详细的信息。</p>

<p>对于HIR的处理过程，程序会通过网络端口，dump到IdealGraphVisualizer里面。而后端的处理过程，缺省则会dump到工作目录下的一个“<code>graal_dumps</code>”子目录下。你可以用文本编辑器打开查看里面的信息。</p>

<pre><code>//至少返回10
public int atLeastTen(int a){
    if (a &lt; 10)
        return 10;
    else
        return a;    
}
</code></pre>

<p>不过，你还可以再偷懒一下，使用一个图形工具<a href="http://lafo.ssw.uni-linz.ac.at/c1visualizer/" target="_blank">c1visualizer</a>来查看。</p>

<p>补充：c1visualizer原本是用于查看Hopspot的C1编译器（也就是客户端编译器）的LIR的工具，这也就是说，Graal的LIR和C1的是一样的。另外，该工具不能用太高版本的JDK运行，我用的是JDK1.8。</p>

<p><img src="assets/468339edbd404568bb71cfedf50c8d40.jpg" alt="" /></p>

<p>图1：atLeatTen()方法对应的LIR</p>

<p>在窗口的左侧，你能看到后端的处理流程。</p>

<ul>
<li>首先是把HIR做最后一次排序（HIR Final Schedule），这个处理会把HIR节点分配到基本块，并且排序；</li>
<li>第二是生成LIR，在这个过程中要做指令选择；</li>
<li>第三，寄存器分配工作，Graal采用的算法是线性扫描（Linear Scan）；</li>
<li>第四，是基于LIR的一些优化工作，比如ControlFlowOptimizer等；</li>
<li>最后一个步骤，是生成目标代码。</li>
</ul>

<p>接下来，我们来认识一下这个LIR：它是怎样生成的，用什么数据结构保存的，以及都有什么特点。</p>

<h2 id="认识lir">认识LIR</h2>

<p>在对HIR的处理过程中，前期（High Tier、Mid Tier）基本上都是与硬件无关。到了后期（Low Tier），你会看到IR中的一些节点逐步开始带有硬件的特征，比如上一讲中，计算AMD64地址的节点。而LIR就更加反映目标硬件的特征了，基本上可以跟机器码一对一地翻译。所以，从HIR生成LIR的过程，就要做指令选择。</p>

<p>我把与LIR相关的包和类整理成了类图，里面划分成了三个包，分别包含了与HIR、LIR和CFG有关的类。你可以重点看看它们之间的相互关系。</p>

<p><img src="assets/d527bd9332b44171bdb609c23533151d.jpg" alt="" /></p>

<p>图2：HIR、LIR和CFG的关联关系</p>

<p>在HIR的最后的处理阶段，程序会通过一个Schedule过程，把HIR节点排序，并放到控制流图中，为生成LIR和目标代码做准备。我之前说过，HIR的一大好处，就是那些浮动节点，可以最大程度地免受控制流的约束。但在最后生成的目标代码中，我们还是要把每行指令归属到某个具体的基本块的。而且，基本块中的HIR节点是按照顺序排列的，在ScheduleResult中保存着这个顺序（blockToNodesMap中顺序保存了每个Block中的节点）。</p>

<p><strong>你要注意</strong>，这里所说的Schedule，跟编译器后端的指令排序不是一回事儿。这里是把图变成线性的程序；而编译器后端的指令排序（也叫做Schedule），则是为了实现指令级并行的优化。</p>

<p>当然，把HIR节点划分到不同的基本块，优化程度是不同的。比如，与循环无关的代码，放在循环内部和外部都是可以的，但显然放在循环外部更好一些。把HIR节点排序的Schedule算法，复杂度比较高，所以使用了很多<strong>启发式</strong>的规则。刚才提到的把循环无关代码放在循环外面，就是一种启发式的规则。</p>

<p>图2中的ControlFlowGraph类和Block类构成了控制流图，控制流图和最后阶段的HIR是互相引用的。这样，你就可以知道HIR中的每个节点属于哪个基本块，也可以知道每个基本块中包含的HIR节点。</p>

<p>做完Schedule以后，接着就会生成LIR。<strong>与声明式的HIR不同，LIR是命令式的，由一行行指令构成。</strong></p>

<p>图1显示的是Foo.atLeatTen方法对应的LIR。你会看到一个控制流图（CFG），里面有三个基本块。B0是B1和B2的前序基本块，B0中的最后一个语句是分支语句（基本块中，只有最后一个语句才可以是导致指令跳转的语句）。</p>

<p>LIR中的指令是放到基本块中的，LIR对象的LIRInstructions属性中，保存了每个基本块中的指令列表。</p>

<p>OK，那接下来，我们来看看LIR的指令都有哪些，它们都有什么特点。</p>

<p>LIRInstruction的子类，主要放在三个包中，你可以看看下面的类图。</p>

<p><img src="assets/35feae7a1e0949d69cdb2dfcfddc5ce0.jpg" alt="" /></p>

<p>图3：LIR中的指令类型</p>

<p>首先，在<strong>org.graalvm.compiler.lir包</strong>中，声明了一些与架构无关的指令，比如跳转指令、标签指令等。因为无论什么架构的CPU，一定都会有跳转指令，也一定有作为跳转目标的标签。</p>

<p>然后，在<strong>org.graalvm.compiler.lir.amd64包</strong>中，声明了几十个AMD64架构的指令，为了降低你的阅读负担，这里我只列出了有代表性的几个。这些指令是LIR代码中的主体。</p>

<p>最后，在<strong>org.graalvm.compiler.hotspot.amd64包</strong>中，也声明了几个指令。这几个指令是利用HotSpot虚拟机的功能实现的。比如，要获取某个类的定义的地址，只能由虚拟机提供。</p>

<p>好了，通过这样的一个分析，你应该对LIR有更加具体的认识了：<strong>LIR中的指令，大多数是与架构相关的。</strong>这样才适合运行后端的一些算法，比如指令选择、寄存器分配等。你也可以据此推测，其他编译器的LIR，差不多也是这个特点。</p>

<p>接下来，我们就来了解一下Graal编译器是如何生成LIR，并且在这个过程中，它是如何实现指令选择的。</p>

<h2 id="生成lir及指令选择">生成LIR及指令选择</h2>

<p>我们已经知道了，Graal在生成LIR的过程中，要进行指令选择。</p>

<p>我们先看一下Graal对一个简单的示例程序Foo.add1，是如何生成LIR的。</p>

<pre><code>public static int add1(int x, int y){
    return x + y + 10;
}
</code></pre>

<p>这个示例程序，在转LIR之前，它的HIR是下面这样。其中有两个加法节点，操作数包括了参数（ParameterNode）和常数（ConstantNode）两种类型。最后是一个Return节点。这个例子足够简单。实际上，它简单到只是一棵树，而不是图。</p>

<p><img src="assets/ca6d9962c3784711ade6447c6a7bef1f.jpg" alt="" /></p>

<p>图4：add1方法对应的HIR</p>

<p>你可以想一下，对于这么简单的一棵树，编译器要如何生成指令呢？</p>

<p>最简单的方法，是做一个语法制导的简单翻译。我们可以深度遍历这棵树，针对不同的节点，分别使用不同的规则来生成指令。比如：</p>

<ul>
<li>在遇到参数节点的时候，我们要搞清楚它的存放位置。因为参数要么是在寄存器中，要么是在栈中，可以直接用于各种计算。</li>
<li>遇到常数节点的时候，我们记下这个常数，用于在下一条指令中作为立即数使用。</li>
<li>在遇到加法节点的时候，生成一个add指令，左右两棵子树的计算结果分别是其操作数。在处理到6号节点的时候，可以不用add指令，而是生成一个lea指令，这样可以直接把结果写入rax寄存器，作为返回值。这算是一个优化，因为可以减少一次从寄存器到寄存器的拷贝工作。</li>
<li>遇到Return节点的时候，看看其子树的计算结果是否放在rax寄存器中。如果不是，那么就要生成一个mov指令，把返回值放入rax寄存器，然后再生成一条返回指令（ret）。通常，在返回之前，编译器还要做一些栈帧的处理工作，把栈指针还原。</li>
</ul>

<p>对于这个简单的例子来说，按照这个翻译规则来生成代码，是完全没有问题的。你可以看下，Graal生成LIR，然后再基于LIR生成的目标代码的示例程序，它只有三行，足够精简和优化：</p>

<pre><code>add    esi,edx       #将参数1加到参数0上，结果保存在esi寄存器
lea    eax,[rsi+0xa] #将rsi加10,结果放入eax寄存器
ret                  #返回
</code></pre>

<p>补充：-
1.我去掉了一些额外的汇编代码，比如用于跟JVM握手，让JVM有机会做垃圾收集的代码。-
2. lea指令原本是用于计算地址的。上面的指令的意思是把rsi寄存器的值作为地址，然后再偏移10个字节，把新的地址放到eax寄存器。-
x86计算机支持间接寻址方式：<strong>偏移量（基址，索引值，字节数）</strong>-
其地址是：<strong>基址 + 索引值*字节数 + 偏移量</strong>-
所以，你可以利用这个特点，计算出<code>a+b*c+d</code>的值。但c（也就是字节数）只能取1、2、4、8。就算让c取1，那也能完成<code>a+b+c</code>的计算。并且，它还可以在另一个操作数里指定把结果写到哪个寄存器，而不像add指令，只能从一个操作数加到另一个操作数上。这些优点，使得x86汇编中经常使用lea指令做加法计算。</p>

<p>Graal编译器实际上大致也是这么做的。</p>

<p>首先，它通过Schedule的步骤，把HIR的节点排序，并放入基本块。对于这个简单的程序，只有一个基本块。</p>

<p><img src="assets/b27d51ef8b81439db3faad9832abf60a.jpg" alt="" /></p>

<p>接着，编译器会对基本块中的节点做遍历（参考：NodeLIRBuilder.java中的<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core/src/org/graalvm/compiler/core/gen/NodeLIRBuilder.java#L363" target="_blank">代码</a>）。针对每个节点转换（Lower）成LIR。</p>

<ul>
<li>把参数节点转换成了MoveFromRegOp指令，在示例程序里，其实这是冗余的，因为可以直接把存放参数的两个寄存器，用于加法计算；</li>
<li>把第一个加法节点转换成了CommutativeTwoOp指令；</li>
<li>把第二个加法节点转换成了LeaOp指令，并且把常数节点变成了LeaOp指令的操作数；</li>
<li>Return节点生成了两条指令，一条是把加法运算的值放到rax寄存器，作为返回值，这条我们知道是冗余的，所以就要看看后面的优化算法如何消除这个冗余；第二条是返回指令。</li>
</ul>

<p><img src="assets/cc5cee6f3d9f4df39af146ad97fd2c2d.jpg" alt="" /></p>

<p>一开始生成的LIR，使用的寄存器都是虚拟的寄存器名称，用v1、v2、v3这些来表示。等把这些虚拟的寄存器对应到具体的物理寄存器以后，就可以消除掉刚才提到的冗余代码了。</p>

<p>我们在c1visualizer中检查一下优化过程，可以发现这是在LinearScanAssignLocationsPhase做的优化。加法指令中对参数1和参数2的引用，变成了对物理寄存器的引用，从而优化掉了两条指令。lea指令中的返回值，也直接赋给了rax寄存器。这样呢，也就省掉了把计算结果mov到rax的指令。这样优化后的LIR，基本上已经等价于目标代码了。</p>

<p><img src="assets/a058ce7ad3ca420dad3e83868ab18305.jpg" alt="" /></p>

<p>好了，通过这样一个分析，你应该理解了从HIR生成LIR的过程。但是还有个问题，这中间似乎也没有做什么指令选择呀？唯一的一处，就是把加法操作优化成了lea指令。而这个也比较简单，基于单独的Add节点就能做出这个优化选择。<strong>那么，更复杂的模式匹配是怎么做的呢？</strong></p>

<p>不要着急，我们接下来就看看Graal是如何实现复杂一点的指令选择的。这一次，我们用了另一个示例程序：Foo.addMemory方法。它把一个类成员变量m和参数a相加。</p>

<pre><code>public class Foo{
    static int m = 3;
    public static int addMemory(int a){
        return m + a;
    } 
    ...
}
</code></pre>

<p>这跟add1方法有所不同，因为它要使用一个成员变量，所以一定要访问内存。而add1方法的所有操作，都是在寄存器里完成的，是“空中作业”，根本不在内存里落地。</p>

<p>我们来看一下这个示例程序对应的HIR。其中一个黄色节点“Read#Foo.m”，是读取内存的节点，也就是读取成员变量m的值。而这又需要通过AMD64Address节点来计算m的地址。由于m是个静态成员，所以它的地址要通过类的地址加上一定的偏移量来计算。</p>

<p><img src="assets/247035473a7b4003869576fdfbcdd347.jpg" alt="" /></p>

<p>图5：addMemory()方法对应的HIR</p>

<p>这里有一个小的知识点，我在第14讲中也提过：对内存操作的节点（如图中的ReadNode），是要加入控制流中的。因为内存里的值，会由于别的操作而改变。如果你把它变成浮动节点，就有可能破坏对内存读写的顺序，从而出现错误。</p>

<p>回到主题，我们来看看怎么为addMemory生成LIR。</p>

<p>如果还是像处理add1方法一样，那么你就会这么做：</p>

<ul>
<li><strong>计算m变量的地址，并放入一个寄存器；</strong></li>
<li><strong>基于这个地址，取出m的值，放入另一个寄存器；</strong></li>
<li><strong>把m的值和参数a做加法。</strong></li>
</ul>

<p>不过这样做，至少要生成3条指令。</p>

<p>在<a href="https://time.geekbang.org/column/article/249261" target="_blank">第8讲</a>中我曾经讲过，像AMD64这样使用复杂指令集（CICS）的架构，具有强大的地址表达能力，并且可以在做算术运算的时候，直接使用内存。所以上面的三条指令，其实能够缩减成一条指令。</p>

<p>这就需要编译器把刚才这种基于内存访问做加法的模式识别出来，以便生成优化的LIR，进而生成优化的目标代码。这也是指令选择算法要完成的任务。可是，<strong>如何识别这种模式呢？</strong></p>

<p>跟踪Graal的执行，你会发现HIR在生成LIR之前，有一个对基本块中的节点做<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core/src/org/graalvm/compiler/core/LIRGenerationPhase.java#L72" target="_blank">模式匹配</a>的操作，进而又调用匹配复杂表达式（<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core/src/org/graalvm/compiler/core/gen/NodeLIRBuilder.java#L430" target="_blank">matchComplexExpressions</a>）。在这里，编译器会把节点跟一个个匹配规则（<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core/src/org/graalvm/compiler/core/match/MatchStatement.java" target="_blank">MatchStatement</a>）做匹配。<strong>注意</strong>，匹配的时候是逆序做的，相当于从树根开始遍历。</p>

<p>在匹配加法节点的时候，Graal匹配上了一个MatchStatement，这个规则的名字叫“addMemory”，是专门针对涉及内存操作的加法运算提供的一个匹配规则。这个MatchStatement包含了一个匹配模式（MatchPattern），该模式的要求是：</p>

<ul>
<li>节点类型是AddNode；</li>
<li>第一个输入（也就是子节点）是一个值节点（value）；</li>
<li>第二个输入是一个ReadNode，而且必须只有一个使用者（singleUser=true）。</li>
</ul>

<p><img src="assets/08a71f3a41904e768ffbf8f4d55ffab5.jpg" alt="" /></p>

<p>图6：匹配规则和匹配模式</p>

<p>这个MatchStatement是在<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core.amd64/src/org/graalvm/compiler/core/amd64/AMD64NodeMatchRules.java#L487" target="_blank">AMD64NodeMatchRules</a>中用注解生成的。利用这样的一个匹配规则，就能够匹配示例程序中的Add节点。</p>

<p>匹配上以后，Graal会把AddNode和ReadNode做上特殊标记，这样在生成LIR的时候，就会按照新的生成规则。生成的LIR如下：</p>

<p><img src="assets/b9c9a19057c04823a39c5b175f5df093.jpg" alt="" /></p>

<p>你可以发现，优化后，编译器把取参数a的指令省略掉了，直接使用了传入参数a的寄存器rsi：</p>

<p><img src="assets/a4a905b3647f4b1590f1744383455ef6.jpg" alt="" /></p>

<p>最后生成的目标代码如下：</p>

<pre><code>movabs rax,0x797b00690          #把Foo类的地址放入rax寄存器
add    esi,DWORD PTR [rax+0x68] #偏移0x68后，是m的地址。做加法
mov    eax,esi                  #设置返回值
ret                             #返回
</code></pre>

<p>到目前为止，你已经了解了Graal是如何匹配一个模式，并选择优化的指令的了。</p>

<p>你可以看看<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core.amd64/src/org/graalvm/compiler/core/amd64/AMD64NodeMatchRules.java#L487" target="_blank">AMD64NodeMatchRules</a>类，它的里面定义了不少这种匹配规则。通过阅读和理解这些规则，你就会对为什么要做指令选择有更加具体的理解了。</p>

<p>Graal的指令选择算法算是比较简单的。在HotSpot的C2编译器中，指令选择采用的是BURS（Bottom-Up Rewrite System，自底向上的重写系统）。这个算法会更加复杂一点，消耗的时间更长，但优化效果更好一些。</p>

<p>这里我补充一个分享，我曾经请教过ARM公司的研发人员，他们目前认为Graal对针对AArch64的指令选择是比较初级的，你可以参考这个<a href="https://static.linaro.org/connect/san19/presentations/san19-514.pdf" target="_blank">幻灯片</a>。所以，他们也正在帮助Graal做改进。</p>

<h2 id="后端的其他功能">后端的其他功能</h2>

<p>出于突出特色功能的目的，这一讲我着重讲了LIR的特点和指令选择算法。不过在考察编译器的后端的时候，我们通常还要注意一些其他功能，比如寄存器分配算法、指令排序，等等。我这里就把Graal在这些功能上的实现特点，给你简单地介绍一下，你如果有兴趣的话，可以根据我的提示去做深入了解：</p>

<ul>
<li><strong>寄存器分配</strong>：Graal采用了线性扫描（Linear Scan）算法。这个算法的特点是速度比较快，但优化效果不如图染色算法。在HotSpot的C2中采用的是后者。</li>
<li><strong>指令排序</strong>：Graal没有为了实现指令级并行而去做指令排序。这里一个主要原因，是现在的很多CPU都已经支持乱序（out-of-order）执行，再做重排序的收益不大。</li>
<li><strong>窥孔优化</strong>：Graal在生成LIR的时候，会做一点窥孔优化（AMD64NodeLIRBuilder类的<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core.amd64/src/org/graalvm/compiler/core/amd64/AMD64NodeLIRBuilder.java#L62" target="_blank">peephole</a>方法）。不过它的优化功能有限，只实现了针对除法和求余数计算的一点优化。</li>
<li><strong>从LIR生成目标代码</strong>：由于LIR已经跟目标代码很接近了，所以这个翻译过程已经比较简单，没有太难的算法了，需要的只是了解和熟悉汇编代码和调用约定。</li>
</ul>

<h2 id="课程小结">课程小结</h2>

<p>这一讲，我带你对Graal的后端做了一个直观的认识，让你的后端知识有了第一个真实世界中编译器的参考系。</p>

<p>第一，把LIR从比较抽象的概念中落地。你现在可以确切地知道哪些指令是必须跟架构相关的，而哪些指令可以跟架构无关。</p>

<p>第二，把指令选择算法落地。虽然Graal的指令选择算法并不复杂，但这毕竟提供了一个可以借鉴的思路，是你认知的一个阶梯。如果你仔细阅读代码，你还可以具象地了解到，符合哪些模式的表达式，是可以从指令选择中受益的。这又是一个理论印证实践的点。</p>

<p>我把这讲的思维导图也放在了下面，供你参考。</p>

<p><img src="assets/df5c2e628a424f1b8250bd9961d07686.jpg" alt="" /></p>

<p>同时，这一讲之后，我们对Java编译器的探讨也就告一段落了。但是，我希望你对它的研究不要停止。</p>

<p>我们讨论的两个编译器（javac和Graal）中的很多知识点，你只要稍微深入挖掘一下，就可以得出不错的成果了。比如，我看到有国外的硕士学生研究了一下HotSpot，就可以发表不错的<a href="http://ssw.jku.at/Research/Papers/Schwaighofer09Master/schwaighofer09master.pdf" target="_blank">论文</a>。如果你是在校大学生，我相信你也可以通过顺着这门课程提供的信息做一些研究，从而得到不错的成果。如果是已经工作的同学，我们可以在极客时间的社群（比如留言区和部落）里保持对Java编译技术的讨论，也一定会对于你的工作有所助益。</p>

<h2 id="一课一思">一课一思</h2>

<p>请你阅读<a href="https://github.com/oracle/graal/blob/vm-20.0.1/compiler/src/org.graalvm.compiler.core.amd64/src/org/graalvm/compiler/core/amd64/AMD64NodeMatchRules.java#L487" target="_blank">AMD64NodeMatchRules</a>中的匹配规则，自己设计另一个例子，能够测试出指令选择的效果。如果降低一下工作量的话，你可以把它里面的某些规则解读一下，在留言区发表你的见解。</p>

<p>好，就到这里。感谢你的阅读，欢迎你把今天的内容分享给更多的朋友，我们下一讲再见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b8d4d4d4818c8989888ff8dfd5d9d1d496dbd7d5" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1820f79b3e63dd',t:'MTczNDExNTcxMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
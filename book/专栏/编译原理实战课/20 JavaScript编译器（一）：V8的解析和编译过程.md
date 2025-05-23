<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=20&#32;JavaScript编译器（一）：V8的解析和编译过程>
        <link rel="icon" href="/static/favicon.png">
        <title>20 JavaScript编译器（一）：V8的解析和编译过程 </title>
        
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
                            <h1 id="title" data-id="20 JavaScript编译器（一）：V8的解析和编译过程" class="title">20 JavaScript编译器（一）：V8的解析和编译过程</h1>
                            <div><p>你好，我是宫文学。从这一讲开始，我们就进入另一个非常重要的编译器：V8编译器。</p>

<p>V8是谷歌公司在2008年推出的一款JavaScript编译器，它也可能是世界上使用最广泛的编译器。即使你不是编程人员，你每天也会运行很多次V8，因为JavaScript是Web的语言，我们在电脑和手机上浏览的每个页面，几乎都会运行一点JavaScript脚本。</p>

<p>扩展：V8这个词，原意是8缸的发动机，换算成排量，大约是4.0排量，属于相当强劲的发动机了。它的编译器，叫做Ignition，是点火装置的意思。而它最新的JIT编译器，叫做TurboFan，是涡轮风扇发动机的意思。</p>

<p>在浏览器诞生的早期，就开始支持JavaScript了。但在V8推出以后，它重新定义了Web应用可以胜任的工作。到今天，在浏览器里，我们可以运行很多高度复杂的应用，比如办公套件等，这些都得益于以V8为代表的JavaScript引擎的进步。2008年V8发布时，就已经比当时的竞争对手快10倍了；到目前，它的速度又已经提升了10倍以上。从中你可以看到，编译技术有多大的潜力可挖掘！</p>

<p>对JavaScript编译器来说，它最大的挑战就在于，当我们打开一个页面的时候，源代码的下载、解析（Parse）、编译（Compile）和执行，都要在很短的时间内完成，否则就会影响到用户的体验。</p>

<p><strong>那么，V8是如何做到既编译得快，又要运行得快的呢？</strong>所以接下来，我将会花两讲的时间，来带你一起剖析一下V8里面的编译技术。在这个过程中，你能了解到V8是如何完成前端解析、后端优化等功能的，它都有哪些突出的特点；另外，了解了V8的编译原理，对你以后编写更容易优化的程序，也会非常有好处。</p>

<p>今天这一讲，我们先来透彻了解一下V8的编译过程，以及每个编译阶段的工作原理，看看它跟我们已经了解的其他编译器相比，有什么不同。</p>

<h2 id="初步了解v8">初步了解V8</h2>

<p>首先，按照惯例，我们肯定要下载V8的源代码。按照<a href="https://v8.dev/docs/build" target="_blank">官方文档</a>中的步骤，你可以下载源代码，并在本地编译。注意，你最好把它编译成Debug模式，这样便于用调试工具去跟踪它的执行，所以你要使用下面的命令来进行编译。</p>

<pre><code>tools/dev/gm.py x64.debug
</code></pre>

<p>编译完毕以后，进入v8/out/x64.debug目录，你可以运行./d8，这就是编译好的V8的命令行工具。如果你用过Node.js，那么d8的使用方法，其实跟它几乎是完全一样的，因为Node.js就封装了一个V8引擎。你还可以用GDB或LLDB工具来调试d8，这样你就可以知道，它是怎么编译和运行JavaScript程序了。</p>

<p>而v8/src目录下的，就是V8的源代码了。V8是用C++编写的。你可以重点关注这几个目录中的代码，它们是与编译有关的功能，而别的代码主要是运行时功能：</p>

<p><img src="assets/5563c6c20f0043c3bbd7a47288cfd7d6.jpg" alt="" /></p>

<p>V8的编译器的构成跟Java的编译器很像，它们都有从源代码编译到字节码的编译器，也都有解释器（叫Ignition），也都有JIT编译器（叫TurboFan）。你可以看下V8的编译过程的图例。在这个图中，你能注意到两个陌生的节点：<strong>流处理节点（Stream）和预解析器（PreParser）</strong>，这是V8编译过程中比较有特色的两个处理阶段。</p>

<p><img src="assets/7cb209a7d24140738cb7fd8c899e4892.jpg" alt="" /></p>

<p>图1：V8的编译过程</p>

<p>注意：这是比较新的V8版本的架构。在更早的版本里，有时会用到两个JIT编译器，类似于HotSpot的C1和C2，分别强调编译速度和优化效果。在更早的版本里，还没有字节码解释器。现在的架构，引入了字节码解释器，其速度够快，所以就取消了其中一级的JIT编译器。</p>

<p>下面我们就进入到V8编译过程中的各个阶段，去了解一些编译器的细节。</p>

<h2 id="超级快的解析过程-词法分析和语法分析">超级快的解析过程（词法分析和语法分析）</h2>

<p>首先，我们来了解一下V8解析源代码的过程。我在开头就已经说过，V8解析源代码的速度必须要非常快才行。源代码边下载边解析完毕，在这个过程中，用户几乎感觉不到停顿。<strong>那它是如何实现的呢？</strong></p>

<p>有两篇文章就非常好地解释了V8解析速度快的原因。</p>

<p>一个是“<a href="https://v8.dev/blog/scanner" target="_blank">optimizing the scanner</a>”这篇文章，它解释了V8在词法分析上做的优化。V8的作者们真是锱铢必较地在每一个可能优化的步骤上去做优化，他们所采用的技术很具备参考价值。</p>

<p>那我就按照我对这篇文章的理解，来给你解释一下V8解析速度快的原因吧：</p>

<p>第一个原因，是<strong>V8的整个解析过程是流（Stream）化的</strong>，也就是一边从网络下载源代码，一边解析。在下载后，各种不同的编码还被统一转化为UTF-16编码单位，这样词法解析器就不需要处理多种编码了。</p>

<p>第二个原因，是<strong>识别标识符时所做的优化</strong>，这也让V8的解析速度更快了一点。你应该知道，标识符的第一个字符（ID_START）只允许用字母、下划线和$来表示，而之后的字符（ID_CONTINUE）还可以包括数字。所以，当词法解析器遇到一个字符的时候，我们首先要判断它是否是合法的ID_START。</p>

<p><strong>那么，这样一个逻辑，通常你会怎么写？</strong>我一般想也不想，肯定是这样的写法：</p>

<pre><code>if(ch &gt;= 'A' &amp;&amp; ch &lt;= 'Z' || ch &gt;='a' &amp;&amp; ch&lt;='z' || ch == '$' || ch == '_'){
    return true;
}
</code></pre>

<p>但你要注意这里的一个问题，<strong>if语句中的判断条件需要做多少个运算？</strong></p>

<p>最坏的情况下，要做6次比较运算和3次逻辑“或”运算。不过，V8的作者们认为这太奢侈了。所以他们通过查表的方法，来识别每个ASCII字符是否是合法的标识符开头字符。</p>

<p>这相当于准备了一张大表，每个字符在里面对应一个位置，标明了该字符是否是合法的标识符开头字符。这是典型的牺牲空间来换效率的方法。虽然你在阅读代码的时候，会发现它调用了几层函数来实现这个功能，但这些函数其实是内联的，并且在编译优化以后，产生的指令要少很多，所以这个方法的性能更高。</p>

<p>第三个原因，是<strong>如何从标识符中挑出关键字</strong>。</p>

<p>与Java的编译器一样，JavaScript的Scanner，也是把标识符和关键字一起识别出来，然后再从中挑出关键字。所以，你可以认为这是一个最佳实践。那你应该也会想到，识别一个字符串是否是关键字的过程，使用的方法仍然是查表。查表用的技术是“<strong>完美哈希（perfect hashing）</strong>”，也就是每个关键字对应的哈希值都是不同的，不会发生碰撞。并且，计算哈希值只用了三个元素：前两个字符（ID_START、ID_CONTINUE），以及字符串的长度，不需要把每个字符都考虑进来，进一步降低了计算量。</p>

<p>文章里还有其他细节，比如通过缩窄对Unicode字符的处理范围来进行优化，等等。从中你能体会到V8的作者们在提升性能方面，无所不用其极的设计思路。</p>

<p>除了词法分析，在语法分析方面，V8也做了很多的优化来保证高性能。其中，最重要的是“<strong>懒解析</strong>”技术（<a href="https://v8.dev/blog/preparser" target="_blank">lazy parsing</a>）。</p>

<p>一个页面中包含的代码，并不会马上被程序用到。如果在一开头就把它们全部解析成AST并编译成字节码，就会产生很多开销：占用了太多CPU时间；过早地占用内存；编译后的代码缓存到硬盘上，导致磁盘IO的时间很长，等等。</p>

<p>所以，所有浏览器中的JavaScript编译器，都采用了懒解析技术。在V8里，首先由预解析器，也就是Preparser粗略地解析一遍程序，在正式运行某个函数的时候，编译器才会按需解析这个函数。你要注意，Preparser只检查语法的正确性，而基于上下文的检查则不是这个阶段的任务。你如果感兴趣的话，可以深入阅读一下这篇<a href="https://v8.dev/blog/preparser" target="_blank">介绍Preparser的文章</a>，我在这里就不重复了。</p>

<p>你可以在终端测试一下懒解析和完整解析的区别。针对foo.js示例程序，你输入“./d8 – ast-print foo.js”命令。</p>

<pre><code>function add(a,b){
    return a + b;
}

//add(1,2)    //一开始，先不调用add函数
</code></pre>

<p>得到的输出结果是：</p>

<p><img src="assets/063aa660433a4a30a1c4de7c29bbdbed.jpg" alt="" /></p>

<p>里面有一个没有名字的函数（也就是程序的顶层函数），并且它记录了一个add函数的声明，仅此而已。你可以看到，Preparser的解析结果确实够粗略。</p>

<p>而如果你把foo.js中最后一行的注释去掉，调用一下add函数，再次让d8运行一下foo.js，就会输出完整解析后的AST，你可以看看二者相差有多大：</p>

<p><img src="assets/9cf357ee11bf4ae68d60a7c423b67249.jpg" alt="" /></p>

<p>最后，你可以去看看正式的Parser（在parser.h、parser-base.h、parser.cc代码中）。学完了这么多编译器的实现机制以后，以你现在的经验，打开一看，你就能知道，这又是用手写的递归下降算法实现的。</p>

<p>在看算法的过程中，我一般第一个就会去看它是如何处理二元表达式的。因为二元表达式看上去很简单，但它需要解决一系列难题，包括左递归、优先级和结合性。</p>

<p>V8的Parser中，对于二元表达式的处理，采取的也是一种很常见的算法：<strong>操作符优先级解析器</strong>（Operator-precedence parser）。这跟Java的Parser也很像，它本质上是自底向上的一个LR(1)算法。所以我们可以得出结论，在手写语法解析器的时候，遇到二元表达式，采用操作符优先级的方法，算是最佳实践了！</p>

<p>好了，现在我们了解了V8的解析过程，那V8是如何把AST编译成字节码和机器码并运行的呢？我们接着来看看它的编译过程。</p>

<h2 id="编译成字节码">编译成字节码</h2>

<p>我们在执行刚才的foo.js文件时，加上“–print-bytecode”参数，就能打印出生成的字节码了。其中，add函数的字节码如下：</p>

<p><img src="assets/e9e8056d1f2944ad9ad90d8944ea6730.jpg" alt="" /></p>

<p>怎么理解这几行字节码呢？我来给你解释一下：</p>

<ul>
<li>Ldar a1：把参数1从寄存器加载到累加器（Ld=load，a=accumulator, r=register）。</li>
<li>Add a0, [0]：把参数0加到累加器上。</li>
<li>Return：返回（返回值在累加器上）。</li>
</ul>

<p>不过，要想充分理解这几行简单的字节码，你还需要真正理解Ignition的设计。因为这些字节码是由Ignition来解释执行的。</p>

<p>Ignition是一个基于寄存器的解释器。它把函数的参数、变量等保存在寄存器里。不过，这里的寄存器并不是物理寄存器，而是指栈帧中的一个位置。下面是一个示例的栈帧：</p>

<p><img src="assets/e5802efb4597487b83f4cbff098cc1d9.jpg" alt="" /></p>

<p>图2：Ignition的栈帧</p>

<p>这个栈帧里包含了执行函数所需要的所有信息：</p>

<ul>
<li>参数和本地变量。</li>
<li>临时变量：它是在计算表达式的时候会用到的。比如，计算2+3+4的时候，就需要引入一个临时变量。</li>
<li>上下文：用来在函数闭包之间维护状态。</li>
<li>pc：调用者的代码地址。</li>
</ul>

<p>栈帧里的a0、a1、r0、r1这些都是寄存器的名称，可以在指令里引用。而在字节码里，会用一个操作数的值代替。</p>

<p>整个栈帧的长度是在编译成字节码的时候就计算好了的。这就让Ignition的栈帧能适应不同架构对栈帧对齐的要求。比如AMD64架构的CPU，它就要求栈帧是16位对齐的。</p>

<p>Ignition也用到了一些物理寄存器，来提高运算的性能：</p>

<ul>
<li><strong>累加器：</strong>在做算术运算的时候，一定会用到累加器作为指令的其中一个操作数，所以它就不用在指令里体现了；指令里只要指定另一个操作数（寄存器）就行了。</li>
<li><strong>字节码数组寄存器：</strong>指向当前正在解释执行的字节码数组开头的指针。</li>
<li><strong>字节码偏移量寄存器：</strong>当前正在执行的指令，在字节码数组中的偏移量（与pc寄存器的作用一样）。</li>
<li>…</li>
</ul>

<p>Ignition是我们见到的第一个寄存器机，它跟我们之前见到的Java和Python的栈机有明显的不同。所以，你可以先思考一下，Ignition会有什么特点呢？</p>

<p>我来给你总结一下吧。</p>

<ol>
<li>它在指令里会引用寄存器作为操作数，寄存器在进入函数时就被分配了存储位置，在函数运行时，栈帧的结构是不变的。而对比起来，栈机的指令从操作数栈里获取操作数，操作数栈随着函数的执行会动态伸缩。</li>
<li>Ignition还引入了累加器这个物理寄存器作为缺省的操作数。这样既降低了指令的长度，又能够加快执行速度。</li>
</ol>

<p>当然，Ignition没有像生成机器码那样，用一个寄存器分配算法，让本地变量、参数等也都尽量采用物理寄存器。这样做的原因，一方面是因为，寄存器分配算法会增加编译的时间；另一方面，这样不利于代码在解释器和TurboFan生成的机器代码之间来回切换（因为它要在调用约定之间做转换）。采用固定格式的栈帧，Ignition就能够在从机器代码切换回来的时候，很容易地设置正确的解释器栈帧状态。</p>

<p>我把更多的字节码指令列在了下面，你可以仔细看一看Ignition都有哪些指令，从而加深对Ignition解释运行机制的理解。同时，你也可以跟我们已经学过的Java和Python的字节码做个对比。这样呀，你对字节码、解释器的了解就更丰富了。</p>

<p><img src="assets/6334f9b8686747059dd6b39e57649128.jpg" alt="" /></p>

<p>来源：<a href="https://docs.google.com/document/d/11T2CRex9hXxoJwbYqVQ32yIPMh0uouUZLdyrtmMoL44/mobilebasic" target="_blank">Ignition Design Doc</a></p>

<h2 id="编译成机器码">编译成机器码</h2>

<p>好，前面我提到了，V8也有自己的JIT编译器，叫做TurboFan。在学过Java的JIT编译器以后，你可以预期到，TurboFan也会有一些跟Java JIT编译器类似的特性，比如它们都是把字节码编译生成机器码，都是针对热点代码才会启动即时编译的。那接下来，我们就来验证一下自己的想法，并一起来看看TurboFan的运行效果究竟如何。</p>

<p>我们来看一个示例程序add.js：</p>

<pre><code>function add(a,b){
    return a+b;
}

for (i = 0; i&lt;100000; i++){
    add(i, i+1);
    if (i%1000==0)
        console.log(i);
}
</code></pre>

<p>你可以用下面的命令，要求V8打印出优化过程、优化后的汇编代码、注释等信息。其中，“–turbo-filter=add”参数会告诉V8，只优化add函数，否则的话，V8会把add函数内联到外层函数中去。</p>

<pre><code>./d8 --trace-opt-verbose \
    --trace-turbo \
    --turbo-filter=add \
    --print-code \
    --print-opt-code \
    --code-comments \
    add.js
</code></pre>

<p>注释：你用./d8 &ndash;help，就能列出V8可以使用的各种选项及其说明，我把上面几个选项的含义解释一下。-
–trace-opt-verbose：跟踪优化过程，并输出详细信息-
–trace-turbo：跟踪TurboFan的运行过程-
–print-code：打印生成的代码-
–print-opt-code：打印优化的代码-
–code-comment：在汇编代码里输出注释</p>

<p>程序一开头是解释执行的。在循环了24000次以后，V8认为这是热点代码，于是启动了Turbofan做即时编译。</p>

<p>最后生成的汇编代码有好几十条指令。不过你可以看到，大部分指令是用于初始化栈帧，以及处理逆优化的情况。真正用于计算的指令，是下面几行指令：</p>

<p><img src="assets/678c511b92c1485e93359a1ebeb22ea7.jpg" alt="" /></p>

<p>对这些汇编代码的解读，以及这些指令的产生和优化过程，我会在下一讲继续给你讲解。</p>

<h2 id="课程小结">课程小结</h2>

<p>今天这讲，我们从总体上考察了V8的编译过程，我希望你记住几个要点：</p>

<ul>
<li>首先，是<strong>编译速度</strong>。由于JavaScript是在浏览器下载完页面后马上编译并执行，它对编译速度有更高的要求。因此，V8使用了一边下载一边编译的技术：懒解析技术。并且，在解析阶段，V8也比其他编译器更加关注处理速度，你可以从中学到通过查表减少计算量的技术。</li>
<li>其次，我们认识了一种新的<strong>解释器Ignition</strong>，它是基于寄存器的解释器，或者叫寄存器机。Ignition比起栈机来，更有性能优势。</li>
<li>最后，我们初步使用了一下V8的<strong>即时编译器TurboFan</strong>。在下一讲中，我们会更细致地探讨TurboFan的特性。</li>
</ul>

<p>按照惯例，这一讲的思维导图我也给你整理出来了，供你参考：</p>

<p><img src="assets/50e554731fd74591b2fb2627f7aee50f.jpg" alt="" /></p>

<h2 id="一课一思">一课一思</h2>

<p>你能否把Ignition的字节码和Java、Python的字节码对比一下。看看它们有哪些共同之处，有哪些不同之处？</p>

<p>欢迎在留言区分享你的答案，也欢迎你把今天的内容分享给更多的朋友。</p>

<h2 id="参考资料">参考资料</h2>

<ol>
<li>这两篇文章分析了V8的解析器为什么速度非常快：<a href="https://v8.dev/blog/scanner" target="_blank">Blazingly fast parsing, part 1: optimizing the scanner</a>，<a href="https://v8.dev/blog/preparser" target="_blank">Blazingly fast parsing, part 2: lazy parsing</a></li>
<li>这篇文章描述了Ignition的设计：<a href="https://docs.google.com/document/d/11T2CRex9hXxoJwbYqVQ32yIPMh0uouUZLdyrtmMoL44/mobilebasic" target="_blank">Ignition Design Doc</a>，我在GitHub上也放了一个<a href="https://github.com/RichardGong/CompilersInPractice/blob/master/v8/Ignition%20Design%20Doc.pdf" target="_blank">拷贝</a></li>
<li>这篇文章有助于你了解Ignition的字节码：<a href="https://medium.com/dailyjs/understanding-v8s-bytecode-317d46c94775" target="_blank">Understanding V8’s bytecode</a></li>
<li>V8项目的<a href="https://v8.dev/" target="_blank">官网</a>，这里有一些重要的博客文章和文档</li>
</ol>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#117d7d7d28252020212651767c70787d3f727e7c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1822e56c3663dd',t:'MTczNDExNTc4OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=43&#32;元编程：通过Proxies和Reflect赋能元编程>
        <link rel="icon" href="/static/favicon.png">
        <title>43 元编程：通过Proxies和Reflect赋能元编程 </title>
        
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
                        <a class="menu-item" id="00 开篇词 JavaScript的进阶之路.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20JavaScript%e7%9a%84%e8%bf%9b%e9%98%b6%e4%b9%8b%e8%b7%af.md">00 开篇词 JavaScript的进阶之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 函数式vs.面向对象：响应未知和不确定.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/01%20%e5%87%bd%e6%95%b0%e5%bc%8fvs.%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%9a%e5%93%8d%e5%ba%94%e6%9c%aa%e7%9f%a5%e5%92%8c%e4%b8%8d%e7%a1%ae%e5%ae%9a.md">01 函数式vs.面向对象：响应未知和不确定.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 如何通过闭包对象管理程序中状态的变化？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/02%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e9%97%ad%e5%8c%85%e5%af%b9%e8%b1%a1%e7%ae%a1%e7%90%86%e7%a8%8b%e5%ba%8f%e4%b8%ad%e7%8a%b6%e6%80%81%e7%9a%84%e5%8f%98%e5%8c%96%ef%bc%9f.md">02 如何通过闭包对象管理程序中状态的变化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 如何通过部分应用和柯里化让函数具象化？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/03%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e9%83%a8%e5%88%86%e5%ba%94%e7%94%a8%e5%92%8c%e6%9f%af%e9%87%8c%e5%8c%96%e8%ae%a9%e5%87%bd%e6%95%b0%e5%85%b7%e8%b1%a1%e5%8c%96%ef%bc%9f.md">03 如何通过部分应用和柯里化让函数具象化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 如何通过组合、管道和reducer让函数抽象化？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/04%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e7%bb%84%e5%90%88%e3%80%81%e7%ae%a1%e9%81%93%e5%92%8creducer%e8%ae%a9%e5%87%bd%e6%95%b0%e6%8a%bd%e8%b1%a1%e5%8c%96%ef%bc%9f.md">04 如何通过组合、管道和reducer让函数抽象化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 map、reduce和monad如何围绕值进行操作？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/05%20map%e3%80%81reduce%e5%92%8cmonad%e5%a6%82%e4%bd%95%e5%9b%b4%e7%bb%95%e5%80%bc%e8%bf%9b%e8%a1%8c%e6%93%8d%e4%bd%9c%ef%bc%9f.md">05 map、reduce和monad如何围绕值进行操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 如何通过模块化、异步和观察做到动态加载？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/06%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%a8%a1%e5%9d%97%e5%8c%96%e3%80%81%e5%bc%82%e6%ad%a5%e5%92%8c%e8%a7%82%e5%af%9f%e5%81%9a%e5%88%b0%e5%8a%a8%e6%80%81%e5%8a%a0%e8%bd%bd%ef%bc%9f.md">06 如何通过模块化、异步和观察做到动态加载？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 深入理解对象的私有和静态属性.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/07%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e5%af%b9%e8%b1%a1%e7%9a%84%e7%a7%81%e6%9c%89%e5%92%8c%e9%9d%99%e6%80%81%e5%b1%9e%e6%80%a7.md">07 深入理解对象的私有和静态属性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 深入理解继承、Delegation和组合.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/08%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e7%bb%a7%e6%89%bf%e3%80%81Delegation%e5%92%8c%e7%bb%84%e5%90%88.md">08 深入理解继承、Delegation和组合.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 面向对象：通过词法作用域和调用点理解this绑定.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/09%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%9a%e9%80%9a%e8%bf%87%e8%af%8d%e6%b3%95%e4%bd%9c%e7%94%a8%e5%9f%9f%e5%92%8c%e8%b0%83%e7%94%a8%e7%82%b9%e7%90%86%e8%a7%a3this%e7%bb%91%e5%ae%9a.md">09 面向对象：通过词法作用域和调用点理解this绑定.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 JS有哪8种数据类型，你需要注意什么？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/10%20JS%e6%9c%89%e5%93%aa8%e7%a7%8d%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%8c%e4%bd%a0%e9%9c%80%e8%a6%81%e6%b3%a8%e6%84%8f%e4%bb%80%e4%b9%88%ef%bc%9f.md">10 JS有哪8种数据类型，你需要注意什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 通过JS引擎的堆栈了解闭包原理.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/11%20%e9%80%9a%e8%bf%87JS%e5%bc%95%e6%93%8e%e7%9a%84%e5%a0%86%e6%a0%88%e4%ba%86%e8%a7%a3%e9%97%ad%e5%8c%85%e5%8e%9f%e7%90%86.md">11 通过JS引擎的堆栈了解闭包原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 JS语义分析该用迭代还是递归？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/12%20JS%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90%e8%af%a5%e7%94%a8%e8%bf%ad%e4%bb%a3%e8%bf%98%e6%98%af%e9%80%92%e5%bd%92%ef%bc%9f.md">12 JS语义分析该用迭代还是递归？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 JS引擎如何实现数组的稳定排序？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/13%20JS%e5%bc%95%e6%93%8e%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%95%b0%e7%bb%84%e7%9a%84%e7%a8%b3%e5%ae%9a%e6%8e%92%e5%ba%8f%ef%bc%9f.md">13 JS引擎如何实现数组的稳定排序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 通过SparkPlug深入了解调用栈.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/14%20%e9%80%9a%e8%bf%87SparkPlug%e6%b7%b1%e5%85%a5%e4%ba%86%e8%a7%a3%e8%b0%83%e7%94%a8%e6%a0%88.md">14 通过SparkPlug深入了解调用栈.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 如何通过哈希查找JS对象内存地址？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/15%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e5%93%88%e5%b8%8c%e6%9f%a5%e6%89%beJS%e5%af%b9%e8%b1%a1%e5%86%85%e5%ad%98%e5%9c%b0%e5%9d%80%ef%bc%9f.md">15 如何通过哈希查找JS对象内存地址？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 为什么环形队列适合做Node数据流缓存？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/16%20%e4%b8%ba%e4%bb%80%e4%b9%88%e7%8e%af%e5%bd%a2%e9%98%9f%e5%88%97%e9%80%82%e5%90%88%e5%81%9aNode%e6%95%b0%e6%8d%ae%e6%b5%81%e7%bc%93%e5%ad%98%ef%bc%9f.md">16 为什么环形队列适合做Node数据流缓存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 如何通过链表做LRU_LFU缓存？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/17%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e9%93%be%e8%a1%a8%e5%81%9aLRU_LFU%e7%bc%93%e5%ad%98%ef%bc%9f.md">17 如何通过链表做LRU_LFU缓存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 TurboFan如何用图做JS编译优化？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/18%20TurboFan%e5%a6%82%e4%bd%95%e7%94%a8%e5%9b%be%e5%81%9aJS%e7%bc%96%e8%af%91%e4%bc%98%e5%8c%96%ef%bc%9f.md">18 TurboFan如何用图做JS编译优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 通过树和图看如何在无序中找到路径和秩序.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/19%20%e9%80%9a%e8%bf%87%e6%a0%91%e5%92%8c%e5%9b%be%e7%9c%8b%e5%a6%82%e4%bd%95%e5%9c%a8%e6%97%a0%e5%ba%8f%e4%b8%ad%e6%89%be%e5%88%b0%e8%b7%af%e5%be%84%e5%92%8c%e7%a7%a9%e5%ba%8f.md">19 通过树和图看如何在无序中找到路径和秩序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 算法思想：JS中分治、贪心、回溯和动态规划.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/20%20%e7%ae%97%e6%b3%95%e6%80%9d%e6%83%b3%ef%bc%9aJS%e4%b8%ad%e5%88%86%e6%b2%bb%e3%80%81%e8%b4%aa%e5%bf%83%e3%80%81%e5%9b%9e%e6%ba%af%e5%92%8c%e5%8a%a8%e6%80%81%e8%a7%84%e5%88%92.md">20 算法思想：JS中分治、贪心、回溯和动态规划.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 创建型：为什么说Redux可以替代单例状态管理.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/21%20%e5%88%9b%e5%bb%ba%e5%9e%8b%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4Redux%e5%8f%af%e4%bb%a5%e6%9b%bf%e4%bb%a3%e5%8d%95%e4%be%8b%e7%8a%b6%e6%80%81%e7%ae%a1%e7%90%86.md">21 创建型：为什么说Redux可以替代单例状态管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 结构型：Vue.js如何通过代理实现响应式编程.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/22%20%e7%bb%93%e6%9e%84%e5%9e%8b%ef%bc%9aVue.js%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e4%bb%a3%e7%90%86%e5%ae%9e%e7%8e%b0%e5%93%8d%e5%ba%94%e5%bc%8f%e7%bc%96%e7%a8%8b.md">22 结构型：Vue.js如何通过代理实现响应式编程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结构型：通过jQuery看结构型模式.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/23%20%e7%bb%93%e6%9e%84%e5%9e%8b%ef%bc%9a%e9%80%9a%e8%bf%87jQuery%e7%9c%8b%e7%bb%93%e6%9e%84%e5%9e%8b%e6%a8%a1%e5%bc%8f.md">23 结构型：通过jQuery看结构型模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 行为型：通过观察者、迭代器模式看JS异步回调.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/24%20%e8%a1%8c%e4%b8%ba%e5%9e%8b%ef%bc%9a%e9%80%9a%e8%bf%87%e8%a7%82%e5%af%9f%e8%80%85%e3%80%81%e8%bf%ad%e4%bb%a3%e5%99%a8%e6%a8%a1%e5%bc%8f%e7%9c%8bJS%e5%bc%82%e6%ad%a5%e5%9b%9e%e8%b0%83.md">24 行为型：通过观察者、迭代器模式看JS异步回调.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 行为型：模版、策略和状态模式有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/25%20%e8%a1%8c%e4%b8%ba%e5%9e%8b%ef%bc%9a%e6%a8%a1%e7%89%88%e3%80%81%e7%ad%96%e7%95%a5%e5%92%8c%e7%8a%b6%e6%80%81%e6%a8%a1%e5%bc%8f%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">25 行为型：模版、策略和状态模式有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 特殊型：前端有哪些处理加载和渲染的特殊“模式”？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/26%20%e7%89%b9%e6%ae%8a%e5%9e%8b%ef%bc%9a%e5%89%8d%e7%ab%af%e6%9c%89%e5%93%aa%e4%ba%9b%e5%a4%84%e7%90%86%e5%8a%a0%e8%bd%bd%e5%92%8c%e6%b8%b2%e6%9f%93%e7%9a%84%e7%89%b9%e6%ae%8a%e2%80%9c%e6%a8%a1%e5%bc%8f%e2%80%9d%ef%bc%9f.md">26 特殊型：前端有哪些处理加载和渲染的特殊“模式”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 性能：如何理解JavaScript中的并行、并发？（上）.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/27%20%e6%80%a7%e8%83%bd%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3JavaScript%e4%b8%ad%e7%9a%84%e5%b9%b6%e8%a1%8c%e3%80%81%e5%b9%b6%e5%8f%91%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">27 性能：如何理解JavaScript中的并行、并发？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 性能：如何理解JavaScript中的并行、并发？（下）.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/28%20%e6%80%a7%e8%83%bd%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3JavaScript%e4%b8%ad%e7%9a%84%e5%b9%b6%e8%a1%8c%e3%80%81%e5%b9%b6%e5%8f%91%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">28 性能：如何理解JavaScript中的并行、并发？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 性能：通过Orinoco、Jank Busters看垃圾回收.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/29%20%e6%80%a7%e8%83%bd%ef%bc%9a%e9%80%9a%e8%bf%87Orinoco%e3%80%81Jank%20Busters%e7%9c%8b%e5%9e%83%e5%9c%be%e5%9b%9e%e6%94%b6.md">29 性能：通过Orinoco、Jank Busters看垃圾回收.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 网络：从HTTP_1到HTTP_3，你都需要了解什么？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/30%20%e7%bd%91%e7%bb%9c%ef%bc%9a%e4%bb%8eHTTP_1%e5%88%b0HTTP_3%ef%bc%8c%e4%bd%a0%e9%83%bd%e9%9c%80%e8%a6%81%e4%ba%86%e8%a7%a3%e4%bb%80%e4%b9%88%ef%bc%9f.md">30 网络：从HTTP_1到HTTP_3，你都需要了解什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 安全：JS代码和程序都需要注意哪些安全问题？.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/31%20%e5%ae%89%e5%85%a8%ef%bc%9aJS%e4%bb%a3%e7%a0%81%e5%92%8c%e7%a8%8b%e5%ba%8f%e9%83%bd%e9%9c%80%e8%a6%81%e6%b3%a8%e6%84%8f%e5%93%aa%e4%ba%9b%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98%ef%bc%9f.md">31 安全：JS代码和程序都需要注意哪些安全问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 测试（一）：开发到重构中的测试.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/32%20%e6%b5%8b%e8%af%95%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%bc%80%e5%8f%91%e5%88%b0%e9%87%8d%e6%9e%84%e4%b8%ad%e7%9a%84%e6%b5%8b%e8%af%95.md">32 测试（一）：开发到重构中的测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 测试（二）：功能性测试.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/33%20%e6%b5%8b%e8%af%95%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%8a%9f%e8%83%bd%e6%80%a7%e6%b5%8b%e8%af%95.md">33 测试（二）：功能性测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 测试（三）：非功能性测试.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/34%20%e6%b5%8b%e8%af%95%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e9%9d%9e%e5%8a%9f%e8%83%bd%e6%80%a7%e6%b5%8b%e8%af%95.md">34 测试（三）：非功能性测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 静态类型检查：ESLint语法规则和代码风格的检查.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/35%20%e9%9d%99%e6%80%81%e7%b1%bb%e5%9e%8b%e6%a3%80%e6%9f%a5%ef%bc%9aESLint%e8%af%ad%e6%b3%95%e8%a7%84%e5%88%99%e5%92%8c%e4%bb%a3%e7%a0%81%e9%a3%8e%e6%a0%bc%e7%9a%84%e6%a3%80%e6%9f%a5.md">35 静态类型检查：ESLint语法规则和代码风格的检查.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Flow：通过Flow类看JS的类型检查.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/36%20Flow%ef%bc%9a%e9%80%9a%e8%bf%87Flow%e7%b1%bb%e7%9c%8bJS%e7%9a%84%e7%b1%bb%e5%9e%8b%e6%a3%80%e6%9f%a5.md">36 Flow：通过Flow类看JS的类型检查.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 包管理和分发：通过NPM做包的管理和分发.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/37%20%e5%8c%85%e7%ae%a1%e7%90%86%e5%92%8c%e5%88%86%e5%8f%91%ef%bc%9a%e9%80%9a%e8%bf%87NPM%e5%81%9a%e5%8c%85%e7%9a%84%e7%ae%a1%e7%90%86%e5%92%8c%e5%88%86%e5%8f%91.md">37 包管理和分发：通过NPM做包的管理和分发.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 编译和打包：通过Webpack、Babel做编译和打包.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/38%20%e7%bc%96%e8%af%91%e5%92%8c%e6%89%93%e5%8c%85%ef%bc%9a%e9%80%9a%e8%bf%87Webpack%e3%80%81Babel%e5%81%9a%e7%bc%96%e8%af%91%e5%92%8c%e6%89%93%e5%8c%85.md">38 编译和打包：通过Webpack、Babel做编译和打包.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 语法扩展：通过JSX来做语法扩展.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/39%20%e8%af%ad%e6%b3%95%e6%89%a9%e5%b1%95%ef%bc%9a%e9%80%9a%e8%bf%87JSX%e6%9d%a5%e5%81%9a%e8%af%ad%e6%b3%95%e6%89%a9%e5%b1%95.md">39 语法扩展：通过JSX来做语法扩展.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 Polyfill：通过Polyfill让浏览器提供原生支持.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/40%20Polyfill%ef%bc%9a%e9%80%9a%e8%bf%87Polyfill%e8%ae%a9%e6%b5%8f%e8%a7%88%e5%99%a8%e6%8f%90%e4%be%9b%e5%8e%9f%e7%94%9f%e6%94%af%e6%8c%81.md">40 Polyfill：通过Polyfill让浏览器提供原生支持.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 微前端：从MVC贫血模式到DDD充血模式.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/41%20%e5%be%ae%e5%89%8d%e7%ab%af%ef%bc%9a%e4%bb%8eMVC%e8%b4%ab%e8%a1%80%e6%a8%a1%e5%bc%8f%e5%88%b0DDD%e5%85%85%e8%a1%80%e6%a8%a1%e5%bc%8f.md">41 微前端：从MVC贫血模式到DDD充血模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 大前端：通过一云多端搭建跨PC_移动的平台应用.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/42%20%e5%a4%a7%e5%89%8d%e7%ab%af%ef%bc%9a%e9%80%9a%e8%bf%87%e4%b8%80%e4%ba%91%e5%a4%9a%e7%ab%af%e6%90%ad%e5%bb%ba%e8%b7%a8PC_%e7%a7%bb%e5%8a%a8%e7%9a%84%e5%b9%b3%e5%8f%b0%e5%ba%94%e7%94%a8.md">42 大前端：通过一云多端搭建跨PC_移动的平台应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 元编程：通过Proxies和Reflect赋能元编程.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/43%20%e5%85%83%e7%bc%96%e7%a8%8b%ef%bc%9a%e9%80%9a%e8%bf%87Proxies%e5%92%8cReflect%e8%b5%8b%e8%83%bd%e5%85%83%e7%bc%96%e7%a8%8b.md">43 元编程：通过Proxies和Reflect赋能元编程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 JavaScript的未来之路：源于一个以终为始的初心.md" href="/%e4%b8%93%e6%a0%8f/JavaScript%20%e8%bf%9b%e9%98%b6%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20JavaScript%e7%9a%84%e6%9c%aa%e6%9d%a5%e4%b9%8b%e8%b7%af%ef%bc%9a%e6%ba%90%e4%ba%8e%e4%b8%80%e4%b8%aa%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%e7%9a%84%e5%88%9d%e5%bf%83.md">结束语 JavaScript的未来之路：源于一个以终为始的初心.md</a>
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
                            <h1 id="title" data-id="43 元编程：通过Proxies和Reflect赋能元编程" class="title">43 元编程：通过Proxies和Reflect赋能元编程</h1>
                            <div><p>你好，我是石川。</p>

<p>今天，我们来到了这个单元最后的一课。前面的两节课中，我们分别学习了微前端和大前端“一大一小”两个趋势，今天我们再来看看“元”编程。那么，元编程中的“元”代表什么呢？“元”有“之上”或“超出一般限制”的意思。听上去，比较玄乎，实际上，要理解和使用元编程，并不难，甚至你可能每天都在用，只是没有觉察到而已。所以今天，就让我们来一步步了解下元编程的概念及使用。</p>

<p>在 JavaScript 中，我们可以把元编程的功能分为几类：第一类是查找和添加对象属性相关的功能；第二类是创建 DSL 这样的特定领域语言；第三类就是可以作为代理用于对象的装饰器。今天，就让我们一一来看一下。</p>

<h2 id="对象属性的属性">对象属性的属性</h2>

<h3 id="1-对象属性的设置">1. 对象属性的设置</h3>

<p>首先，我们先来看下对象相关属性的查找和添加。我们都知道 Javascript 对象的属性包含了名称和值。但是另外，我们需要了解的是，每个属性本身也有三个相关的属性，它们分别为<strong>可写属性（writable）、可枚举属性（enumerable）以及可配置属性（configurable）</strong>。</p>

<p>这三个属性指定了该属性的行为方式，以及我们可以使用它们做什么。这里的可写属性指定了属性的值是否可以更改；可枚举属性指定了属性是否可以由 <code>for/in</code> 循环和 <code>Object.keys()</code> 方法枚举；可配置属性指定了是否可以删除属性或更改属性的属性。</p>

<p>这里，需要注意的是，我们自定义的对象字面量和通过赋值定义的对象属性都是可写、可枚举和可配置的，但 JavaScript 标准库中定义的很多对象属性都不是。从这里，你应该可以看出，只要你使用过 <code>for/in</code>，那么，恭喜你，基本你已经使用过元编程开发了。</p>

<p><strong>那么对属性的查询和设置有什么用呢？</strong>这对于很多第三方库的开发者来说是很重要的，因为它允许开发者向原型对象中添加方法，而且也可以像标准库中的很多内置方法一样，将它们设置成不可枚举；同时，它也可以允许开发者“锁定”对象，让属性定义无法被更改或删除。下面，我们就来看看可以为 JavaScript 第三方库的开发赋能的查询和设置属性的 API。</p>

<p>我们可以将属性分为两类，一类是“<strong>数据属性</strong>”，一类是“<strong>访问器属性</strong>”。如果我们把值、getter、setter 都看做是值的话，那么，数据属性就包含了值、可写、可枚举和可配置这4个属性；而访问器属性则包含了 get、set、可枚举和可配置这4个属性。其中可写、可枚举和可配置属性是布尔值，get 和 set 属性是函数值。</p>

<p>我们可以通过 <code>object.getOwnPropertyDescriptor()</code> 来获取对象属性的属性描述，顾名思义，这个方法只适用于获取对象自身的属性，如果要查询继承属性的属性，就需要通过 <code>Object.getPrototypeOf()</code> 的方法来遍历原型链。如果要设置属性的属性或者使用指定的属性创建新属性，就需要用到 <code>Object.defineProperty()</code> 的方法。</p>

<h3 id="2-对象的可延展性">2. 对象的可延展性</h3>

<p>除了对对象的属性的获取和设置外，我们对对象本身也可以设置它的可延展性。我们可以通过 <code>Object.isExtensible()</code> 让一个对象可延展，同样的，我们也可以通过 <code>Object.preventExtensions()</code> 将一个对象设置为不可延展。</p>

<p>不过这里需要注意的是，一旦我们把一个对象设置为不可延展，我们不仅不可以在对象上再设置属性，而且我们也不可以再把对象改回为可延展。还有就是这个不可延展只影响对象本身的属性，对于对象的原型对象上的属性来说，是不受影响的。</p>

<p>对象的可延展性通常的作用是用于对像状态的锁定。通常我们把它和属性设置中的可写属性和可配置属性结合来使用。在 JavaScript 中，我们可以通过 <code>Object.seal()</code> 把不可延展属性和不可配置属性结合；通过 <code>Object.freeze()</code> 我们可以把不可延展、不可配置和不可写属性结合起来。</p>

<p>对于 JavaScript 第三方库的编写来说，如果库将对象传递给库的回调函数，那么我们就可以使用 <code>Object.freeze()</code> 来防止用户的代码对它们进行修改了。不过需要注意的是，这样的方法可能对 JavaScript 测试策略形成干扰。</p>

<h3 id="3-对象的原型对象">3. 对象的原型对象</h3>

<p>前面，我们介绍的 <code>Object.freeze()</code> 、 <code>Object.seal()</code> 和属性设置的方法一样，都是仅作用于对象本身的，都不会对对象的原型造成影响。我们知道，通过 new 创建的对象会使用创建函数的原型值作为自己的原型，通过 <code>Object.create()</code> 创建的对象会使用第一个参数作为对象的原型。</p>

<p>我们可以通过 <code>Object.getPrototypeOf()</code> 来获取对象的原型；通过 <code>isPrototypeOf()</code> 我们可以判断一个对象是不是另外一个对象的原型；同时，如果我们想要修改一个对象的原型，可以通过 <code>Object.setPrototypeOf()</code>。不过有一点需要注意的是，通常在原型已经设置后，就很少被改变了，使用 <code>Object.setPrototypeOf()</code> 有可能对性能产生影响。</p>

<h2 id="用于dsl的模版标签">用于DSL的模版标签</h2>

<p>我们知道，在 JavaScript 中，在反引号内的字符串被称为模板字面量。当一个值为函数的表达式，并且后面跟着一个模板字面量时，它会变成一个函数被调用，我们将它称之为“<strong>带标签的模板字面量</strong>”。</p>

<p>为什么我们说定义一个新的标签函数，用于标签模板字面量可以被当做是一种元编程呢？因为标签模板通常用于定义 DSL，也就是域特定语言，这样定义新的标签函数就如同向 JavaScript 中添加了新的语法。标签模板字面量已被许多前端 JavaScript 库采用。GraphQL 查询语言通过使用 gql`` 标签函数，可以使查询被嵌入到 JavaScript 代码中。Emotion 库使用 css`` 标签函数，使 CSS 样式同样可以被嵌入到 JavaScript 中。</p>

<p>当函数表达式后面有模板字面量时，该函数将被调用。第一个参数是字符串数组，后面是零或多个其它参数，这些参数可以具有任何类型的值。参数的数量取决于插入到模板字面量的值的数量，模板字面量的值始终是字符串，但是标签模板字面量的值是标签函数返回的值。它可能是一个字符串，但当使用标签函数实现 DSL 时，返回值通常是一个非字符串数据结构，它是字符串的解析表示。</p>

<p>当我们想将一个值安全地插入到 HTML 字符串中时，模版会非常得有用。我们拿 html`` 为例，在使用标签构建最终字符串之前，标签会对每个值执行 HTML 转义。</p>

<pre><code class="language-javascript">function html(str, ...val) {
    var escaped = val.map(v =&gt; String(v)
                                  .replace(&quot;&amp;&quot;, &quot;&amp;amp;&quot;)
                                  .replace(&quot;'&quot;, &quot;&amp;#39;&quot;));
    var result = str[0];
    for(var i = 0; i &lt; escaped.length; i++) {
        result += escaped[i] + str[i+1];
    }
    return result;
}

var operator = &quot;&amp;&quot;;
html`&lt;b&gt;x ${operator} y&lt;/b&gt;`             // =&gt; &quot;&lt;b&gt;x &amp;amp; y&lt;/b&gt;&quot;
</code></pre>

<p>下面，我们再来看看 Reflect 对象。Reflect 并不是一个类，和 Math 对象类似，它的属性只是定义了一组相关的函数。ES6 中添加的这些函数都在一个命名空间中，它们模仿核心语言的行为，并且复制了各种预先存在于对象函数的特性。</p>

<p>尽管 Reflect 函数没有提供任何新功能，但它们确实将这些功能组合在一个 API 中方便使用。比如，我们在上面提到的对象属性的设置、可延展性以及对象的原型对象在 Reflect 中都有对应的方法，如 <code>Reflect.set()</code>、<code>Reflect.isExtensible()</code> 和 <code>Reflect.getPrototypeOf()</code>，等等。下面，我们会看到 Reflect 函数集与 Proxy 的处理程序方法集也可以一一对应。</p>

<h2 id="proxy-和-reflect">Proxy 和 Reflect</h2>

<p>在 ES6 和更高版本中提供的 Proxy 类可以算是 JavaScript 中最强大的元编程功能了。它允许我们编写改变 JavaScript 对象基本行为的代码。我们在前面提到的 Reflect API 是一组函数，它使我们可以直接访问 JavaScript 对象上的一组基本操作。当我们创建 Proxy 对象时，我们指定了另外两个对象，目标对象和处理程序对象。</p>

<pre><code class="language-javascript">var target = {
  message1: &quot;hello&quot;,
  message2: &quot;world&quot;,
};
var handler = {};

var proxy = new Proxy(target, handler);
</code></pre>

<p><strong>生成的代理对象没有自己的状态或行为。</strong>无论何时对其执行操作（读取属性、写入属性、定义新属性、查找原型、将其作为函数调用），它都会将这些操作分派给处理程序对象或目标对象。代理对象支持的操作与 Reflect API 定义的操作相同。Proxy 的工作机制是，如果 handler 是空的，那么代理对象只是一层透明的装饰器。所以在上面的例子中，如果我们执行代理，那么它返回的结果就是目标对象上本来自有的属性。</p>

<pre><code class="language-javascript">console.log(proxy.message1); // hello
console.log(proxy.message2); // world
</code></pre>

<p>通常，我们会把 Proxy 和 Reflect 结合起来使用，这样的好处是，对于我们不想自定义的部分，我们可以使用 Reflect 来调用对象内置的方法。</p>

<pre><code class="language-javascript">const target = {
  message1: &quot;hello&quot;,
  message2: &quot;world&quot;,
};

const handler = {
  get(target, prop, receiver) {
    if (prop === &quot;message2&quot;) {
      return &quot;Jackson&quot;;
    }
    return Reflect.get(...arguments);
  },
};
var proxy = new Proxy(target, handler);
console.log(proxy.message1); // hello
console.log(proxy.message2); // Jackson
</code></pre>

<h2 id="总结">总结</h2>

<p>通过今天这节课，我们看到了 JavaScript 对象的属性本身也带有可以查找和添加的属性。同时，我们学习了，JavaScript 定义的函数允许我们遍历对象的原型链，甚至更改对象的原型。</p>

<p>标签模板字面量是一种函数调用语法，我们可以用它来定义新的标签函数，这样做有点像向语言中添加新的字面量语法。通过定义一个解析其模板字符串参数的标签函数，我们就可以在JavaScript 代码中嵌入 DSL。标签函数还提供对字符串字面量的原始、非转义形式的访问，其中反斜杠不带有任何特殊含义。</p>

<p>最后，我们又看了 Proxy 类和相关的 Reflect API。Proxy 和 Reflect 允许我们对 JavaScript 中的对象的基本行为进行低级控制。Proxy 对象可以用作可选的、可撤销的包装器，以改进代码封装，还可以用于实现非标准对象行为（如早期 Web 浏览器定义的一些特殊情况下的 API）。</p>

<h2 id="思考题">思考题</h2>

<p>我们知道 Symbol 对象的属性的值可以用于定义对象和类的属性或方法的名称。这样做可以控制对象如何与 JavaScript 语言特性和核心库交互。那么，你觉得 Symbol 在这种使用场景下算不算是一种元编程呢？</p>

<p>欢迎在留言区分享你的观点、交流学习心得或者提出问题，如果觉得有收获，也欢迎你把今天的内容分享给更多的朋友。我们下节课再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#18747474212c2929282f587f75797174367b7775" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f4b73c822cd81',t:'MTczNDAyMzA3OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
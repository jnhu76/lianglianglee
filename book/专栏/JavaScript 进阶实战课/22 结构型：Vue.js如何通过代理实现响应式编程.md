<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=22&#32;结构型：Vue.js如何通过代理实现响应式编程>
        <link rel="icon" href="/static/favicon.png">
        <title>22 结构型：Vue.js如何通过代理实现响应式编程 </title>
        
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
                            <h1 id="title" data-id="22 结构型：Vue.js如何通过代理实现响应式编程" class="title">22 结构型：Vue.js如何通过代理实现响应式编程</h1>
                            <div><p>你好，我是石川。</p>

<p>上一讲我们介绍了几种不同的创建型模式，今天我们来说说设计模式中的结构型模式。在结构型模式中，最经典、最常用的就非代理模式莫属了。在JavaScript的开发中，代理模式也是出现频率比较高的一种设计模式。</p>

<p>前端的朋友们应该都对Vue.js不会感到陌生。Vue.js的一个很大的特点就是用到了如今流行的响应式编程（Reactive Programming）。那它是怎么做到这一点的呢？这里面离不开代理模式，这一讲我们主要解答的就是这个问题。但是在解开谜底之前，我们先来看看代理模式比较传统直观的一些应用场景和实现方式吧。</p>

<h2 id="代理的应用场景">代理的应用场景</h2>

<p>在代理设计模式中，一个代理对象充当另一个主体对象的接口。很多人会把代理模式和门面模式做对比。后面几讲，我们会介绍到门面模式。这里我们需要了解的是代理模式与门面模式不同，门面模式最主要的功能是简化了接口的设计，把复杂的逻辑实现隐藏在背后，把不同的方法调用结合成更便捷的方法提供出来。而代理对象在调用者和主体对象之间，主要起到的作用是保护和控制调用者对主体对象的访问。代理会拦截所有或部分要在主体对象上执行的操作，有时会增强或补充它的行为。</p>

<p><img src="assets/e577e494cd174c469e49a08424a74137.jpg" alt="图片" /></p>

<p>如上图所示，一般代理和主体具有相同的接口，这对调用者来说是透明的。代理将每个操作转发给主体，通过额外的预处理或后处理增强其行为。这种模式可能看起来像“二道贩子”，但存在即合理，代理特别是在性能优化方面还是起到了很大作用的。下面我们就一一来看看。</p>

<h3 id="延迟初始化和缓存">延迟初始化和缓存</h3>

<p>因为代理充当了主体对象的保护作用，减少了客户端对代理背后真实主体无效的消耗。这就好像是公司里的销售，他们不是收到所有的客户需求都直接透传给到项目组来开发，而是接到客户的需求后，先给对方一个产品服务手册，当对方真的选择某项服务以及确定购买后，这时候售前才会转到项目团队来实施，再把结果交付给客户。</p>

<p><img src="assets/6d84a410d87a416ba539c9b080dd7495.jpg" alt="图片" /></p>

<p>这个例子是我们可以称之为<strong>延迟初始化</strong>的应用。在这种情况下，代理可以作为接口提供帮助。代理接受初始化请求，在明确调用者确实会使用主体之前不会传递给它。客户端发出初始化请求并且代理先做响应，但实际上并没有将消息传递给主体对象，直到客户端明显需要主体完成一些工作，只有这样的情况下，代理才会将两条消息一起传递。</p>

<p>在这个基础之上，我们就可以看到第二个例子。在这个例子中，代理除了起到延迟初始化的作用外，还可以增加<strong>一层缓存</strong>。当客户端第一次访问的时候，代理会合并请求给主体，并且把结果先缓存，再分开返回给客户端。在客户端第二次发起请求2的时候，代理可以直接从缓存中读取信息，不需要访问主体，就直接返回给客户端。</p>

<p><img src="assets/41336131c55c4727ba00b47ec30cfa1a.jpg" alt="图片" /></p>

<h3 id="代理的其它应用场景">代理的其它应用场景</h3>

<p>作为一门Web语言，JavaScript经常要和网络请求打交道，基于上述的方式，就能对性能优化起到很大的作用。除了延迟初始化和缓存外，代理还在很多其它方面有着很重要的用途。比如数据验证，代理可以在将输入转发给主体之前对输入的内容进行验证，确保无误后，再传给后端。除此之外，代理模式也可以用于安全验证，代理可以用来验证客户端是否被授权执行操作，只有在检查结果为肯定的情况下，才将请求发送给后端。</p>

<p>代理还有一个应用场景是日志记录，代理可以通过拦截方法调用和相关参数，重新编码。另外，它还可以获取远程对象并放到本地。说完了代理的应用场景，接下来我们可以看一下它在JavaScript中的实现方式。</p>

<h2 id="代理的实现方式">代理的实现方式</h2>

<p>代理模式在JavaScript中有很多种实现方式。其中包含了：1. 对象组合或对象字面量加工厂模式；2. 对象增强；3. 使用从ES6开始自带的内置的Proxy。这几种方式分别有它们的优劣势。</p>

<h3 id="组合模式">组合模式</h3>

<p>我们先看看组合模式，在上一节我们讲过了，基于函数式编程的思想，我们在编程中，应该尽量保证主体的不变性。基于这个原则，组合可以被认为是创建代理的一种简单而安全的方法，因为它使主体保持不变，从而不会改变其原始行为。它唯一的缺点是我们必须手动 delegate 所有方法，即使我们只想代理其中的一个方法。此外，我们可能必须 delegate 对主题属性的访问。还有一点就是如果想要做到延迟初始化的话，基本只可以用到组合。下面我用伪代码做个展示：</p>

<pre><code class="language-javascript">class Calculator {
  constructor () {
    /*...*/
  }
  plus () { /*...*/ }
  minus () { /*...*/ }
}

class ProxyCalculator {
  constructor (calculator) {
    this.calculator = calculator
  }
  // 代理的方法
  plus () { return this.calculator.divide() }
  minus () { return this.calculator.multiply() }
}

var calculator = new Calculator();
var proxyCalculator = new ProxyCalculator(calculator);
</code></pre>

<p>除了上述的方式外，我们也可以基于组合的思路用工厂函数来做代理创建。</p>

<pre><code class="language-javascript">function factoryProxyCalculator (calculator) {
  return {
    // 代理的方法
    plus () { return calculator.divide() },
    minus () { return calculator.multiply() }
  }
}

var calculator = new Calculator();
var proxyCalculator = new factoryProxyCalculator(calculator);
</code></pre>

<h3 id="对象增强">对象增强</h3>

<p>再来说第二种模式对象增强（Object Augmentation），对象增强还有一个名字叫猴子补丁（Monkey Patching）。对于对象增强来说，它的优点就是不需要 delegate 所有方法。但是它最大的问题是改变了主体对象。用这种方式确实是简化了代理创建的工作，但弊端是会造成函数式编程思想中的“副作用”，因为在这里，主体不再具有不可变性。</p>

<pre><code class="language-javascript">function patchingCalculator (calculator) {
  var plusOrig = calculator.plus
  calculator.plus = () =&gt; {
    // 额外的逻辑
    // 委托给主体
    return plusOrig.apply(calculator)
  }
  return calculator
}
var calculator = new Calculator();
var safeCalculator = patchingCalculator(calculator);
</code></pre>

<h3 id="内置proxy">内置Proxy</h3>

<p>最后，我们再来看看使用ES6内置的Proxy。从ES6之后，JavaScript便支持了Proxy。它结合了对象组合和对象增强各自的优点，我们既不需要手动的去 delegate 所有的方法，也不会改变主体对象，保持了主体对象的不变性。但是它也有一个缺点，就是它几乎没有polyfill。也就是说，如果使用内置的代理，就要考虑在兼容性上做出一定的牺牲。真的是鱼和熊掌不能兼得。</p>

<pre><code class="language-javascript">var ProxyCalculatorHandler = {
  get: (target, property) =&gt; {
    if (property === 'plus') {
      // 代理的方法
      return function () {
        // 额外的逻辑
        // 委托给主体
        return target.divide();
      }
    }
    // 委托的方法和属性
    return target[property]
  }
}
var calculator = new Calculator();
var proxyCalculator = new Proxy(calculator, ProxyCalculatorHandler);
</code></pre>

<h2 id="vue如何用代理实现响应式编程">VUE如何用代理实现响应式编程</h2>

<p>我们在上一讲讲到单例模式时，从解决状态管理时用到的Redux和reducer，可以看出一些三方库对传统面向对象的模式下，加入函数式编程来解决问题的思路。今天我们再来剖析下另外一个库，也就是Vue.js状态管理的思想。回到开篇的问题，Vue.js 是如何用代理实现响应式编程的呢？这里Vue.js通过代理创建了一种Change Obsverver的设计模式。</p>

<p>Vue.js 最显着的特点之一是<strong>无侵入的反应系统（unobtrusive reactivity system）</strong>。组件状态是响应式 JavaScript 对象，当被修改时，UI会更新。就像我们使用Excel时，如果我们在A2这个格子里设置了一个 A0 和 A1 相加的公式 “= A0 + A1”的话，当我们改动 A0 或 A1 的值的时候，A2也会随之变化。这也是我们在前面说过很多次的在函数式编程思想中的<strong>副作用（side effect）</strong>。</p>

<p><img src="assets/7d47de2a022340259cee839596e8eb9a.jpg" alt="图片" /></p>

<p>在JavaScript中，如果我们用命令式编程的方式， 可以看到这种副作用是不存在的。</p>

<pre><code class="language-javascript">var A0 = 1;
var A1 = 2;
var A2 = A0 + A1;
console.log(A2) // 返回是 3
A0 = 2;
console.log(A2) // 返回仍然是 3
</code></pre>

<p>但响应式编程（Reactive Programming）是一种基于声明式编程的范式。如果要做到响应式编程，我们就会需要下面示例中这样一个 update 的更新功能。这个功能会使得每次当 A0 或 A1 发生变化时，更新 A2 的值。这样做，其实就产生了副作用，update 就是这个副作用。A0 和 A1 被称为这个副作用的依赖。这个副作用是依赖状态变化的订阅者。whenDepsChange 在这里是个伪代码的订阅功能。</p>

<pre><code class="language-javascript">var A2;
function update() {
  A2 = A0 + A1;
}
whenDepsChange(update);
</code></pre>

<p>在 JavaScript 中没有 whenDepsChange 这样的机制可以跟踪局部变量的读取和写入 。Vue.js 能做的，是拦截对象属性的读写。JavaScript 中有两种拦截属性访问的方法：getter/setter 和 Proxies。由于浏览器支持限制，Vue 2 仅使用 getter/setter。在 Vue 3 中，Proxies 用于响应式对象，getter/setter 用于通过属性获取元素的 refs。下面是一些说明响应式对象如何工作的伪代码：</p>

<pre><code class="language-javascript">function reactive(obj) {
  return new Proxy(obj, {
    get(target, key) {
      track(target, key)
     return target[key]
    },
    set(target, key, value) {
      target[key] = value
      trigger(target, key)
    }
  })
}
</code></pre>

<p>你可能会想，上面对set和get的拦截和自定义，怎么就能做到对变化的观察呢？这就要说到 handler 里包含一系列具有预定义名称的可选方法了，称为<strong>陷阱方法</strong>（trap methods），例如：apply、get、set 和 has，在代理实例上执行相应操作时会自动调用这些方法。所以我们在拦截和自定义后，它们会在对象发生相关变化时被自动调用。所以假设我们可以订阅A0和A1的值的话，那么在这两个值有改动的情况下，就可以自动计算A2的更新。</p>

<pre><code class="language-javascript">import { reactive, computed } from 'vue'
var A0 = reactive(0);
var A1 = reactive(1);
var A2 = computed(() =&gt; A0.value + A1.value);
A0.value = 2;
</code></pre>

<h2 id="延伸-proxy还可以用于哪些场景">延伸：Proxy还可以用于哪些场景</h2>

<p>JavaScript内置的Proxy除了作为代理以外，还有很多作用。基于它的拦截和定制化的特点，Proxy也广泛用于对象虚拟化（object virtualization）、运算符重载（operator overloading）和最近很火的元编程（meta programming）。这里我们不用伪代码，换上一些简单的真代码，看看<strong>陷阱方法</strong>（trap methods）的强大之处。</p>

<h3 id="对象虚拟化">对象虚拟化</h3>

<p>我们先来看一下对象虚拟化，下面的例子中的 oddNumArr 单数数组就是一个虚拟的对象。我们可以查看一个单双数是不是在单数数组里，我们也可以获取一个单数，但是实际上这个数组里并没有储存任何数据。</p>

<pre><code class="language-javascript">const oddNumArr = new Proxy([], {
  get: (target, index) =&gt; index % 2 === 1 ? index : Number(index)+1,
  has: (target, number) =&gt; number % 2 === 1
})

console.log(4 in oddNumArr) // false
console.log(7 in oddNumArr) // true
console.log(oddNumArr[15])   // 15
console.log(oddNumArr[16])   // 17
</code></pre>

<h3 id="运算符重载">运算符重载</h3>

<p>运算符重载就是对已有的运算符重新进行定义，赋予其另一种功能，以适应不同的数据类型。比如在下面的例子中，我们就是通过重载“.”这个符号，所以在执行obj.count时，我们看到它同时返回了拦截get和set自定义的方法，以及返回了计数的结果。</p>

<pre><code class="language-javascript">var obj = new Proxy({}, {
  get: function (target, key, receiver) {
    console.log(`获取 ${key}!`);
    return Reflect.get(target, key, receiver);
  },
  set: function (target, key, value, receiver) {
    console.log(`设置 ${key}!`);
    return Reflect.set(target, key, value, receiver);
  }
});

obj.count = 1; // 返回：设置 count!
obj.count; 
// 返回：获取 count!
// 返回：设置 count!
// 返回：1
</code></pre>

<p>除了对象虚拟化和运算符重载 ，Proxy的强大之处还在于元编程的实现。但是元编程这个话题过大，我们后面会专门花一节课来讲。</p>

<h2 id="总结">总结</h2>

<p>通过今天的内容，我们再一次看到，面向对象、函数式和响应式的交集。经典的关于设计模式的书对代理的解释更多是单纯的面向对象，但是通过开篇的问题，我们可以看到代理其实也是解决响应式编程的状态追踪问题的一把利器。所以，从Vue.js用到的基于代理的变化观察者（change observer）模式，我们也可以看出任何设计模式都不是绝对的，而是可以互相结合形成新的模式来解决问题。</p>

<h2 id="思考题">思考题</h2>

<p>我们说响应式设计用到了很多函数式编程的思想，但是又不完全是函数式编程，你能说出它们的一些差别吗？</p>

<p>欢迎在留言区分享你的答案、交流学习心得或者提出问题，如果觉得有收获，也欢迎你把今天的内容分享给更多的朋友。我们下期再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#dab6b6b6e3eeebebeaed9abdb7bbb3b6f4b9b5b7" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f459dc82bcd81',t:'MTczNDAyMjg0MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
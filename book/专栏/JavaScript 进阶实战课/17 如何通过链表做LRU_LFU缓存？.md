<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=17&#32;如何通过链表做LRU_LFU缓存？>
        <link rel="icon" href="/static/favicon.png">
        <title>17 如何通过链表做LRU_LFU缓存？ </title>
        
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
                            <h1 id="title" data-id="17 如何通过链表做LRU_LFU缓存？" class="title">17 如何通过链表做LRU_LFU缓存？</h1>
                            <div><p>你好，我是石川。</p>

<p>前面我们在第10-13讲讲过了数组这种数据类型，以及通过它衍生出的栈和队列的数据结构。之后，我们在讲到散列表的时候，曾经提到过一种链表的数据结构。今天，我们就来深入了解下链表和它所支持的相关的缓存算法。链表有很多使用场景，最火的例子当属目前热门的区块链了。它就是用了链表的思想，每一个区块儿之间都是通过链表的方式相连的。在链表结构里，每一个元素都是节点，每个节点有自己的内容和相连的下一个元素的地址参考。</p>

<p>既然我们讲的是JavaScript，还是回到身边的例子，就是缓存。无论是我们的操作系统，还是浏览器，都离不开缓存。我们把链表和散列表结合起来，就可以应用于我们常说的缓存。那这种缓存是如何实现的呢？接下来就先从链表说起吧。</p>

<h2 id="如何实现单双循环链表">如何实现单双循环链表</h2>

<h3 id="单链表">单链表</h3>

<p>下面我们先来看看一个单向链表是怎么实现的。链表就是把零散的节点（node）串联起来的一种数据结构。在每个节点里，会有两个核心元素，一个是数据，一个是下一个节点的地址，我们称之为后继指针（next）。在下面的例子里，我们先创建了一个节点（node），里面有节点的数据（data）和下一个节点的地址（next）。</p>

<p><img src="assets/3302e44b2e1a42c19bc4d06c6b4b7098.jpg" alt="图片" /></p>

<p>在实现上，我们可以先创建一个节点的函数类，里面包含存储数据和下一节点两个属性。</p>

<pre><code class="language-javascript">class Node {
  constructor(data){
    this.data = data;
    this.next = null;
  }
}
</code></pre>

<p>之后，我们再创建一个链表的函数类。在链表里，最常见的方法就是头尾的插入和删除，这几项操作的复杂度都是<span class="math inline">\(O(1)\)</span>。应用在缓存当中的时候呢，通常是头部的插入，尾部删除。但倘若你想从前往后遍历的话，或是在任意位置做删除或插入的操作，那么相应的复杂度就会是<span class="math inline">\(O(n)\)</span>。这里的重点不是实现，而是要了解链表里面的主要功能。因此我没有把所有代码都贴在这里，但是我们可以看到它核心的结构。</p>

<pre><code class="language-javascript">class LinkedList {
  constructor(){
    this.head = null;
    this.size = 0;
  }
  isEmpty(){ /*判断是否为空*/ }
  size() { /*获取链表大小*/ }
  getHead() { /*获取链表头节点*/ } 
  indexOf(element) { /*获取某个元素的位置*/ } 
  insertAtTail(element) { /*在表尾插入元素*/ }
  insertAt(element, index) { /*在指定位置插入*/ }
  remove(element) { /*删除某个元素*/ }
  removeAt(index) { /*在指定位置删除*/ }
  getElementAt(index) { /*根据某个位置获取元素*/ }
  removeAtHead(){ /*在表头位置删除元素*/ }
}
</code></pre>

<p>注意观察的你，可能会发现，最后一个节点是null或undefined的，那它是做什么的呢？这个是一个哨兵节点，它的目的啊，是对插入和删除性能的优化。因为在这种情况下，插入和删除首尾节点的操作就可以和处理中间节点用同样的代码来实现。在小规模的数据中，这种区别可能显得微不足道。可是在处理缓存等需要大量频繁的增删改查的操作中，就会有很大的影响。</p>

<h3 id="双链表">双链表</h3>

<p>说完单链表，我们再来看看双链表。在前面单链表的例子里，我们可以看到它是单向的，也就是每一个节点都有下一个节点的地址，但是没有上一个的节点的指针。双链表的意思就是双向的，也就是说我们也可以顺着最后一个节点中关于上一个节点的地址，从后往前找到第一个节点。</p>

<p><img src="assets/888d7a62bc9849f3b9307d624b431192.jpg" alt="图片" /></p>

<p>所以在双链表节点的实现上，我们可以在单链表基础上增加一个上一节点指针的属性。同样的，双链表也可以基于单链表扩展，在里面加一个表尾节点。对于从后往前的遍历，和从前往后的遍历一样，复杂度都是<span class="math inline">\(O(n)\)</span>。</p>

<pre><code class="language-javascript">class DoublyNode extends Node { 
  constructor(data, next, prev) {
    super(data, next); 
    this.prev = prev; 
  }
}

class DoublyLinkedList extends LinkedList {
  constructor() {
    this.tail = undefined;
  }
}
</code></pre>

<h3 id="循环链表">循环链表</h3>

<p>我们接下来再来看看循环链表（circular list）。循环链表，顾名思义，就是一个头尾相连的链表。</p>

<p><img src="assets/98b1a6f2f69843f5847af6c237c6c740.jpg" alt="图片" /></p>

<p>如果是双向循环列表的话，就是除了顺时针的头尾相接以外，从逆时针也可以循环的双循环链表。</p>

<p><img src="assets/7ca0f407f1704f9e9b7cb9cbf0f14cdd.jpg" alt="图片" /></p>

<h2 id="如何通过链表来做缓存">如何通过链表来做缓存</h2>

<p>了解了单双循环链表后，现在我们回到题目，那我们如何能通过链表来做缓存呢？我们先了解下缓存的目的。缓存主要做的就是把数据放到临时的内存空间，便于再次访问。比如数据库会有缓存，目的是减少对硬盘的访问，我们的浏览器也有缓存，目的是再次访问页面的时候不用重新下载相关的文本、图片或其它素材。通过链表，我们可以做到缓存。下面，我们会看两种缓存的算法，他们是最近最少使用（LRU，least recently used）和最不经常使用（LFU，least frequently used），简称LRU缓存和LFU缓存。</p>

<p>一个最优的缓存方案应该把最不可能在未来访问的内容换成最有可能被访问的新元素。但是在实际情况中，这是不可能的，因为没人能预测未来。所以作为一个次优解，在缓存当中，通常会考虑两个因素，一个是时间局部性（Temporal Locality），我们认为一个最近被访问的内存位有可能被再次访问；另外一个因素是空间局部性(Spatial Locality)，空间局部性考虑的是一个最近被访问的内存位相近的位置也有可能被再次访问。</p>

<p>在实际的使用中，LRU缓存要比LFU的使用要多，你可能要问为什么？这里我先卖个关子，咱们先看看这两种方法的核心机制和实现逻辑。之后我们再回到这个问题。</p>

<h3 id="lfu-缓存">LFU 缓存</h3>

<p>我们先来看看LFU缓存。一个经典的LFU缓存中，有两个散列表，一个是键值散列表，里面存放的是节点。还有一个是频率散列表，里面根据每个访问频率，会有一个双链表，如下图中所示，如果键值为2和3的两个节点内容都是被访问一次，那么他们在频率1下面会按照访问顺序被加到链表中。</p>

<p><img src="assets/5800b831c76b43df860eef6b1922be21.jpg" alt="图片" /></p>

<p>在这里，我们看一个LFU关键的代码实现部分的说明。LFU双链表节点，可以通过继承双链表的节点类，在里面增加需要用到的键值，除此之外还有一个计数器来计算一个元素被获取和设置的频率。</p>

<pre><code class="language-javascript">class LFUNode extends DoublyNode { 
  constructor(key) {
    this.key = key; 
    this.freqCount = 1;
  }
}

class LFUDoublyLinkedList extends LinkedList {
  constructor() {/* LFU 双链表 */ }
}
</code></pre>

<p>前面说到一个LFU缓存里面有两个散列表，一个是键值散列表，一个是频率散列表。这两个散列表都可以通过对象来创建。键值散列表保存着每个节点的实例。频率散列表里有从1到n的频率的键名，被访问和设置次数最多的内容，就是n。频率散列表里每一个键值都是一个双向链表。</p>

<p>那围绕它呢，这里有3个比较核心的操作场景，分别是插入、更新和读取。</p>

<p>对于插入的场景，也就是<strong>插入新的节点</strong>，这里我们要看缓存是否已经满了，如果是，那么在插入时它的频率是1。如果没有满的话，那么新的元素在插入的同时，尾部的元素就会被删除。</p>

<p>对于更新场景，也就是<strong>更新旧的节点</strong>，这时，这个元素会被移动到表头。同时，为了计算下一个被删除的元素，最小的频率minFreq会减1。</p>

<p>对于读取场景，也就是<strong>获取节点的动作</strong>，缓存可以返回节点，并且增加它的调用频率的计数，同时将这个元素移动到双链表的头部。同样与插入场景类似，最后一步，最低频率minFreq的值会被调整，用来计算下次操作中会被取代的元素。</p>

<pre><code class="language-javascript">class LFUCache {
  constructor() {
    this.keys = {}; // 用来存储LFU节点的键值散列表
    this.freq = {}; // 用来存储LFU链表的频率散列表
    this.capacity = capacity; // 用来定义缓存的承载量
    this.minFreq = 0; // 把最低访问频率初始设置为0
    this.size = 0; // 把缓存大小初始设置为0
  }
  set() {/* 插入一个节点 */}
  get() {/* 读取一个节点 */}
}
</code></pre>

<h3 id="lru-缓存">LRU 缓存</h3>

<p>以上是LFU的机制和核心实现逻辑，下面我们再来看看最近最少使用（LRU least recently used）。LRU缓存的目的是在有限的内存空间中，当有新的内容需要缓存的时候，清除最老的内容缓存。当一个缓存中的元素被访问了，这个最新的元素就会被放到列表的最后。当一个没有在缓存中的内容被访问的时候，最前面的也就是最老的内容就会被删除。</p>

<p><img src="assets/0989142bd7514733b648afcde8ae84cb.jpg" alt="图片" /></p>

<p>在LRU缓存的实现中，最需要注意的是要追踪哪个节点在什么时间被访问了。为了达到这个目的，我们同样要用到链表和散列表。之所以要用到双向链表，是因为需要追踪在头部需要删除的最老的内容，和在尾部插入最新被访问的内容。</p>

<p>从实现的角度看，LRU当中的节点和LFU没有特别大的差别，也需要一个键值散列表，但是不需要频率散列表。LRU缓存可以通过传入承载量capacity参数来定义可以允许缓存的量。同样和LFU类似的，LRU也需要在链表头部删除节点和链表尾部增加节点的功能。在此基础之上，有着获取和设置的两个对外使用的方法。所以总体看来，LRU的实现和LFU相比是简化了的。</p>

<pre><code class="language-javascript">class LRUNode extends DoublyNode { 
  constructor(key) {
    this.key = key; 
  }
}

class LRUCache {
  constructor() {
    this.keys = {}; // 用来存储LFU节点的键值散列表
    this.capacity = capacity; // 用来定义缓存的承载量
    this.size = 0; // 把缓存大小初始设置为0
  }
  set() {/* 插入一个节点 */}
  get() {/* 读取一个节点 */}
}
</code></pre>

<h2 id="总结">总结</h2>

<p>如果从直觉上来看，你可能会觉得对比LRU，似乎LFU是更合理的一种缓存方法，它根据访问的频率来决定对内容进行缓存或清除。回到上面的问题，为什么在实际应用中，更多被使用的是LRU呢？这是因为一是短时间高频访问本身不能证明缓存有长期价值，二是它可能把一些除了短时间高频访问外的内容挤掉，第三就是刚被访问的内容也很可能不会被频繁访问，但是仍然可能被重复访问，这里也增加了它们被意外删除的可能性。</p>

<p>所以总的来说，LFU没有考虑时间局部性，也就是“最近被访问的可能再次被访问”这个问题。所以在实际的应用中，使用LRU的概率要大于LFU。希望通过这节课，你能对链表和散列表的使用有更多的了解，更主要的是，当你的应用需要写自己的缓存的时候，也可以把这些因素考虑进去。</p>

<h2 id="思考题">思考题</h2>

<p>下面到了思考题时间，虽然说LFU的应用场景比较少，但是它在某些特定场合还是很有用处的，你能想到一些相关的场景和实现吗？</p>

<p>欢迎在留言区分享你的答案、交流学习心得或者提出问题，如果觉得有收获，也欢迎你把今天的内容分享给更多的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#86eaeaeabfb2b7b7b6b1c6e1ebe7efeaa8e5e9eb" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f42ef5b44cd81',t:'MTczNDAyMjczMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=36&#32;Flow：通过Flow类看JS的类型检查>
        <link rel="icon" href="/static/favicon.png">
        <title>36 Flow：通过Flow类看JS的类型检查 </title>
        
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
                            <h1 id="title" data-id="36 Flow：通过Flow类看JS的类型检查" class="title">36 Flow：通过Flow类看JS的类型检查</h1>
                            <div><p>你好，我是石川。</p>

<p>前面我们讲了除了功能性和非功能性测试外，代码的质量检查和风格检查也能帮助我们发现和避免程序中潜在的问题，今天我们再来看看另外一种发现和避免潜在问题的方法——代码类型的检查。说到类型检查，TypeScript 可能是更合适的一门语言，但既然我们这个专栏主讲的是JavaScript，所以今天我们就通过 Flow ——这个 <strong>JavaScript 的语言扩展</strong>，来学习 JavaScript 中的类型检查。</p>

<p>如果你有 C 或 Java 语言的开发经验，那么对类型注释应该不陌生。我们拿 C 语言最经典的 Hello World 来举个例子，这里的 int 就代表整数，也就是函数的类型。那么，我们知道在JavaScript中，是没有类型注释要求的。而通过 Flow，我们既可以做类型注释，也可以对注释和未注释的代码做检查。</p>

<pre><code class="language-javascript">#include &lt;stdio.h&gt;

int main() {
  printf(&quot;Hello World! \n&quot;);
  return 0;
}
</code></pre>

<p>它的<strong>工作原理</strong>很简单，总结起来就3步。</p>

<ol>
<li>给代码加类型注释；</li>
<li>运行Flow工具来分析代码类型和相关的错误报告；</li>
<li>当问题修复后，我们可以通过Babel或其它自动化的代码打包流程来去掉代码中的类型注释。</li>
</ol>

<p>这里，你可能会想，为什么我们在第3步会删除注释呢？这是因为 Flow 的语言扩展本身并没有改变 JavaScript 本身的编译或语法，所以它只是我们<strong>代码编写阶段的静态类型检查</strong>。</p>

<h2 id="为什么需要类型">为什么需要类型</h2>

<p>在讲 Flow 前，我们先来熟悉下类型检查的使用场景和目的是什么？类型注释和检查最大的应用场景是较为大型的复杂项目开发。在这种场景下，严格的类型检查可以避免代码执行时<strong>由于类型的出入而引起的潜在问题</strong>。</p>

<p>另外，我们也来看看 Flow 和 TypeScript 有什么共同点和区别。这里，我们可以简单了解下。先说说<strong>相同点</strong>：</p>

<ul>
<li>首先，TypeScript 本身也是 JavaScript 的一个扩展，顾名思义就是“类型脚本语言”；</li>
<li>其次，TypeScript 加 TSC 编译器的代码注释和检查流程与 Flow 加 Babel 的模式大同小异；</li>
<li>对于简单类型的注释，两者也是相似的；</li>
<li>最后，它们的应用场景和目的也都是类似的。</li>
</ul>

<p>再看看 Flow 和 TypeScript 这两者的<strong>区别</strong>：</p>

<ul>
<li>首先，在对于相对高阶的类型上，两者在语法上有所不同，但要实现转换并不难；</li>
<li>其次，TypeScript 是早于 ES6，在2012年发布的，所以它虽然名叫 TypeScript，但在当时除了强调类型外，也为了弥补当时 ES5 中没有 class、for/of 循环、模块化或 Promises 等不足而增加了很多功能；</li>
<li>除此之外，TypeScript 也增加了很多自己独有的枚举类型（enum）和命名空间（namespace）等关键词，所以它可以说是一个独立的语言了。而 Flow 则更加轻量，它建立在JavaScript的基础上，加上了类型检查的功能，而不是自成一派，但改变了 JavaScript 语言本身。</li>
</ul>

<h2 id="安装和运行">安装和运行</h2>

<p>有了一些 Flow 相关的基础知识以后，下面我们正式进入主题，来看看Flow的安装和使用。</p>

<p>和我们之前介绍的 JavaScript 之器中的其它工具类似，我们也可以通过 NPM 来安装 Flow。这里，我们同样也使用了 <code>-g</code>，这样的选项可以让我们通过命令行来运行相关的程序。</p>

<pre><code class="language-javascript">npm install -g flow-bin
</code></pre>

<p>在使用 Flow 做代码类型检查前，我们需要通过下面的命令在项目所在地文件目录下做初始化。通过初始化，会创建一个以 .flowconfig 为后缀的配置文件。虽然我们一般不需要修改这个文件，但是对于 Flow 而言，可以知道我们项目的位置。</p>

<pre><code class="language-javascript">npm run flow --init
</code></pre>

<p>在第一次初始化之后，后续的检查可以直接通过 <code>npm run flow</code> 来执行。</p>

<pre><code class="language-javascript">npm run flow
</code></pre>

<p>Flow 会找到项目所在位置的所有 JavaScript 代码，但只会对头部标注了 <code>// @flow</code> 的代码文件做类型检查。这样的好处是<strong>对于已有项目，我们可以对文件逐个来处理，按阶段有计划地加上类型注释。</strong></p>

<p>前面，我们说过即使对于没有注释的代码，Flow 也可以进行检查，比如下面的例子，因为我们没有把 for 循环中的 <code>i</code> 设置为本地变量，就可能造成对全局的污染，所以在没有注释的情况下，也会收到报错。一个简单的解决方案就是在 for 循环中加一段 <code>for (let i = 0)</code> 这样的代码。</p>

<pre><code class="language-javascript">// @flow
let i = { x: 0, y: 1 };
for(i = 0; i &lt; 10; i++) { 
  console.log(i);
}
i.x = 1;  
</code></pre>

<p>同样，下面的例子在没有注释的情况下也会报错。虽然 Flow 开始不知道 <code>msg</code> 参数的类型，但是看到了长度 <code>length</code> 这个属性后，就知道它不会是一个数字。但是后面在函数调用过程中传入的实参却是数字，明显会产生问题，所以就会报错。</p>

<pre><code class="language-javascript">// @flow
function msgSize(msg) {
  return msg.length;
}
let message = msgSize(10000);
</code></pre>

<h2 id="类型注释的使用">类型注释的使用</h2>

<p>上面，我们讲完了在没有注释的情况下的类型检查。下面，我们再来看看类型注释的使用。当你声明一个 JavaScript 变量的时候，可以通过一个冒号和类型名称来增加一个 Flow 的类型注释。比如下面的例子中，我们声明了数字、字符串和布尔的类型。</p>

<pre><code class="language-javascript">// @flow
let num: number = 32;
let msg: string = &quot;Hello world&quot;;
let flag: boolean = false;
</code></pre>

<p>同上面未注释的例子一样，即使在没有注释的情况下，Flow 也可以通过变量声明赋值来判断值的类型。<strong>唯一的区别是在有注释的情况下，Flow 会对比注释和赋值的类型，如果发现两者间有出入，便会报错。</strong></p>

<p>函数参数和返回值的注释与变量的注释类似，也是通过冒号和类型名称。在下面的例子中，我们把参数的类型注释为字符串，返回值的类型注释为了数字。当我们运行检查的时候，虽然函数本身可以返回结果，但是会报错。这是因为我们期待的返回值是字符串，而数组长度返回的结果却是数字。</p>

<pre><code class="language-javascript">// 普通函数
function msgSize(msg: string): string {
  return msg.length;
}
console.log(msgSize([1,2,3]));
</code></pre>

<p>不过有一点需要注意的是 JavaScript 和 Flow 中 <code>null</code> 的类型是一致的，但是 JavaScript 中的 <code>undefined</code> 在 Flow 中是 <code>void</code>。而且针对函数没有返回值的情况，我们也可以用 <code>void</code> 来注释。</p>

<p>如果你想允许 <code>null</code> 或 <code>undefiend</code> 作为合法的变量或参数值，只需要在类型前<strong>加一个问号的前缀</strong>。比如在下面的例子中，我们使用了 <code>?string</code>，这个时候，虽然它不会对 <code>null</code> 参数本身的类型报错，但会报错说 <code>msg.length</code> 是不安全的，这是因为 <code>msg</code> 可能是 <code>null</code> 或 <code>undefined</code> 这些没有长度的值。为了解决这个报错，我们可以使用一个判断条件，只有在判断结果是<strong>真值</strong>的情况下，会再返回 <code>msg.length</code>。</p>

<pre><code class="language-javascript">// @flow

function msgSize(msg: ?string): number {
  return msg.length;
}
console.log(msgSize(null));

function msgSize(msg: ?string): number {
  return msg ? msg.length : -1;
}
console.log(msgSize(null));
</code></pre>

<h2 id="复杂数据类型的支持">复杂数据类型的支持</h2>

<p>到目前为止，我们学习了几个原始数据类型：字符串、数字、布尔、null 和 undefined 的检查，并且也学习了 Flow 在变量声明赋值、函数的参数和返回值中的使用。下面我们来看一些 Flow 对其它更加复杂的数据类型检查的支持。</p>

<h3 id="类">类</h3>

<p>首先我们先来看一下类。类的关键词 <code>class</code> 不需要额外的注释，但是我们可以对里面的属性和方法做类型注释。比如在下面的例子中，<code>prop</code> 属性的类型是数字。方法 <code>method</code> 的参数是字符串，返回值我们可以定义为数字。</p>

<pre><code class="language-javascript">// @flow
class MyClass {
  prop: number = 42;
  method(value: string): number { /* ... */ }
}
</code></pre>

<h3 id="对象和类型别名">对象和类型别名</h3>

<p>Flow 中的对象类型看上去很像是对象字面量，区别是 Flow 的属性值是类型。</p>

<pre><code class="language-javascript">// @flow
var obj1: { foo: boolean } = { foo: true };
</code></pre>

<p>在对象中，如果一个属性是可选的，我们可以通过下面的问号的方式来代替 <code>void</code> 和 <code>undefined</code>。如果没有注释为可选，那就默认是必须存在的。如果我们想改成可选，同样需要加一个问号。</p>

<pre><code class="language-plain">var obj: { foo?: boolean } = {};
</code></pre>

<p>Flow 对函数中没有标注的额外的属性是不会报错的。如果我们想要 Flow 来严格执行只允许明确声明的属性类型，可以通过增加以下<strong>竖线</strong>的方式声明相关的对象类型。</p>

<pre><code class="language-javascript">// @flow
function method(obj: { foo: string }) {
  // ...
}
method({
  foo: &quot;test&quot;, // 通过
  bar: 42      // 通过
});

{| foo: string, bar: number |} 
</code></pre>

<p>对于过长的类型对象参数，我们可以通过<strong>自定义类型名称</strong>的方式将参数类抽象提炼出来。</p>

<pre><code class="language-javascript">// @flow
export type MyObject = {
  x: number,
  y: string,
};

export default function method(val: MyObject) {
  // ... 
}
</code></pre>

<p>我们可以像导出一个模块一样地导出类型。其它的模块可以用导入的方式来引用类型定义。但是，这里需要注意的是，导入类型是 Flow 语言的一个延伸，不是一个实际的 JavaScript 导入指令。类型的导入导出只是被 Flow 作为类型检查来使用的，在最终执行的代码中，会被删除。最后需要注意的是，和创建一个 type 比起来，更简洁的方式是<strong>直接定义一个 MyObject 类，用来作为类型。</strong></p>

<p>我们知道在 JavaScript 中，对象有时被用来当做字典或字符串到值的映射。属性的名称是后知的，没法声明成一个 Flow 类型。但是，我们还是可以通过 Flow 来描述数据结构。假设你有一个对象，它的属性是城市的名称，值是城市的位置。我们可以将数据类型通过下面的方式声明。</p>

<pre><code class="language-javascript">// @flow
var cityLocations : {[string]: {long:number, lat:number}} = {
    &quot;上海&quot;: { long: 31.22222, lat: 121.45806 }
};
export default cityLocations;
</code></pre>

<h3 id="数组">数组</h3>

<p>数组中的同类元素类型可以在角括号中说明。一个有着固定长度和不同类型元素的数组，叫做<strong>元组（tuple）</strong>。在元组中，元素的类型在用逗号相隔的方括号中声明。</p>

<p>我们可以通过解构（destructuring）赋值的方式，加上 Flow 的类型别名功能来使用元组。</p>

<p>如果我们希望函数能够接受一个任意长度的数组作为参数时就不能使用元组了，这时我们需要用的是 <code>Array&lt;mixed&gt;</code>。<code>mixed</code> 表示数组的元素可以是任意类型。</p>

<pre><code class="language-javascript">// @flow

function average(data: Array&lt;number&gt;) {
 // ...
}

let tuple: [number, boolean, string] = [1, true, &quot;three&quot;];
let num  : number  = tuple[0]; // 通过
let bool : boolean = tuple[1]; // 通过
let str  : string  = tuple[2]; // 通过

function size(s: Array&lt;mixed&gt;): number {
    return s.length;
}
console.log(size([1,true,&quot;three&quot;]));
</code></pre>

<p>如果我们的函数对数组中的元素进行检索和使用，Flow 检查会使用类型检查或其他测试来确定元素的类型。如果你愿意放弃类型检查，你也可以使用 any 而不是 mixed，它允许你对数组的值做任何处理，而<strong>不需要确保这些值是期望的类型</strong>。</p>

<h3 id="函数">函数</h3>

<p>我们已经了解了如何通过添加类型注释来指定函数的参数类型和返回值的类型。但是，在高阶函数中，当函数的参数之一本身是函数时，我们也需要能够指定该函数参数的类型。要用 Flow 表示函数的类型，需要写出用逗号分隔、再用括号括起来的每个参数的类型，然后用箭头表示，最后键入函数的返回类型。</p>

<p>下面是一个期望传递回调函数的示例函数。这里注意下，我们是如何为回调函数的类型定义类型别名的。</p>

<pre><code class="language-javascript">// @flow
type FetchTextCallback = (?Error, ?number, ?string) =&gt; void;
function fetchText(url: string, callback: FetchTextCallback) {
  // ...  
}
</code></pre>

<h2 id="总结">总结</h2>

<p>虽然使用类型检查需要很多额外的工作量，当你开始使用 Flow 的时候，也可能会扫出来大量的问题，但这是很正常的。一旦你掌握了类型规范，就会发现它可以避免很多的潜在问题，比如函数中的输入和输出值类型与期待的参数或结果不一致。</p>

<p>另外，除了我们介绍的数字、字符串、函数、对象和数组等这几种核心的数据类型，类型检查还有很多用法，你也可以通过<a href="https://flow.org/en/" target="_blank">Flow 的官网</a>了解更多。</p>

<h2 id="思考题">思考题</h2>

<p>在类型检查中，有两种思想，一种是可靠性（soundness），一种是完整性（completeness）。可靠性是检查任何可能会在运行时发生的问题，完整性是检查一定会在运行时发生的问题。第一种思想是“宁可误杀，也不放过”；而后者的问题是有时可能会让问题“逃脱”。那么你知道Flow或TypeScript遵循的是哪种思想吗？你觉得哪种思想更适合实现？</p>

<p>欢迎在留言区分享你的看法、交流学习心得或者提出问题，如果觉得有收获，也欢迎你把今天的内容分享给更多的朋友。我们下节课再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#610d0d0d58555050515621060c00080d4f020e0c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f4a482f66cd81',t:'MTczNDAyMzAzMS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
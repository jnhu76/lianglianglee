<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=39&#32;案例分析（二）：高性能网络应用框架Netty>
        <link rel="icon" href="/static/favicon.png">
        <title>39 案例分析（二）：高性能网络应用框架Netty </title>
        
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
                        <a class="menu-item" id="00 学习攻略 如何才能学好并发编程？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/00%20%e5%ad%a6%e4%b9%a0%e6%94%bb%e7%95%a5%20%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e5%ad%a6%e5%a5%bd%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%ef%bc%9f.md">00 学习攻略 如何才能学好并发编程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 你为什么需要学习并发编程？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bd%a0%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e5%ad%a6%e4%b9%a0%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%ef%bc%9f.md">00 开篇词 你为什么需要学习并发编程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 可见性、原子性和有序性问题：并发编程Bug的源头.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/01%20%e5%8f%af%e8%a7%81%e6%80%a7%e3%80%81%e5%8e%9f%e5%ad%90%e6%80%a7%e5%92%8c%e6%9c%89%e5%ba%8f%e6%80%a7%e9%97%ae%e9%a2%98%ef%bc%9a%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8bBug%e7%9a%84%e6%ba%90%e5%a4%b4.md">01 可见性、原子性和有序性问题：并发编程Bug的源头.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Java内存模型：看Java如何解决可见性和有序性问题.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/02%20Java%e5%86%85%e5%ad%98%e6%a8%a1%e5%9e%8b%ef%bc%9a%e7%9c%8bJava%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e5%8f%af%e8%a7%81%e6%80%a7%e5%92%8c%e6%9c%89%e5%ba%8f%e6%80%a7%e9%97%ae%e9%a2%98.md">02 Java内存模型：看Java如何解决可见性和有序性问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 互斥锁（上）：解决原子性问题.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/03%20%e4%ba%92%e6%96%a5%e9%94%81%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e8%a7%a3%e5%86%b3%e5%8e%9f%e5%ad%90%e6%80%a7%e9%97%ae%e9%a2%98.md">03 互斥锁（上）：解决原子性问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 互斥锁（下）：如何用一把锁保护多个资源？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/04%20%e4%ba%92%e6%96%a5%e9%94%81%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e4%b8%80%e6%8a%8a%e9%94%81%e4%bf%9d%e6%8a%a4%e5%a4%9a%e4%b8%aa%e8%b5%84%e6%ba%90%ef%bc%9f.md">04 互斥锁（下）：如何用一把锁保护多个资源？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 一不小心就死锁了，怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/05%20%e4%b8%80%e4%b8%8d%e5%b0%8f%e5%bf%83%e5%b0%b1%e6%ad%bb%e9%94%81%e4%ba%86%ef%bc%8c%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">05 一不小心就死锁了，怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 用“等待-通知”机制优化循环等待.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/06%20%e7%94%a8%e2%80%9c%e7%ad%89%e5%be%85-%e9%80%9a%e7%9f%a5%e2%80%9d%e6%9c%ba%e5%88%b6%e4%bc%98%e5%8c%96%e5%be%aa%e7%8e%af%e7%ad%89%e5%be%85.md">06 用“等待-通知”机制优化循环等待.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 安全性、活跃性以及性能问题.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/07%20%e5%ae%89%e5%85%a8%e6%80%a7%e3%80%81%e6%b4%bb%e8%b7%83%e6%80%a7%e4%bb%a5%e5%8f%8a%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98.md">07 安全性、活跃性以及性能问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 管程：并发编程的万能钥匙.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/08%20%e7%ae%a1%e7%a8%8b%ef%bc%9a%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e7%9a%84%e4%b8%87%e8%83%bd%e9%92%a5%e5%8c%99.md">08 管程：并发编程的万能钥匙.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 Java线程（上）：Java线程的生命周期.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/09%20Java%e7%ba%bf%e7%a8%8b%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aJava%e7%ba%bf%e7%a8%8b%e7%9a%84%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f.md">09 Java线程（上）：Java线程的生命周期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Java线程（中）：创建多少线程才是合适的？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/10%20Java%e7%ba%bf%e7%a8%8b%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%88%9b%e5%bb%ba%e5%a4%9a%e5%b0%91%e7%ba%bf%e7%a8%8b%e6%89%8d%e6%98%af%e5%90%88%e9%80%82%e7%9a%84%ef%bc%9f.md">10 Java线程（中）：创建多少线程才是合适的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Java线程（下）：为什么局部变量是线程安全的？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/11%20Java%e7%ba%bf%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%b1%80%e9%83%a8%e5%8f%98%e9%87%8f%e6%98%af%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e7%9a%84%ef%bc%9f.md">11 Java线程（下）：为什么局部变量是线程安全的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 如何用面向对象思想写好并发程序？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/12%20%e5%a6%82%e4%bd%95%e7%94%a8%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e6%80%9d%e6%83%b3%e5%86%99%e5%a5%bd%e5%b9%b6%e5%8f%91%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">12 如何用面向对象思想写好并发程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 理论基础模块热点问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/13%20%e7%90%86%e8%ae%ba%e5%9f%ba%e7%a1%80%e6%a8%a1%e5%9d%97%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">13 理论基础模块热点问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Lock和Condition（上）：隐藏在并发包中的管程.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/14%20Lock%e5%92%8cCondition%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e9%9a%90%e8%97%8f%e5%9c%a8%e5%b9%b6%e5%8f%91%e5%8c%85%e4%b8%ad%e7%9a%84%e7%ae%a1%e7%a8%8b.md">14 Lock和Condition（上）：隐藏在并发包中的管程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Lock和Condition（下）：Dubbo如何用管程实现异步转同步？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/15%20Lock%e5%92%8cCondition%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aDubbo%e5%a6%82%e4%bd%95%e7%94%a8%e7%ae%a1%e7%a8%8b%e5%ae%9e%e7%8e%b0%e5%bc%82%e6%ad%a5%e8%bd%ac%e5%90%8c%e6%ad%a5%ef%bc%9f.md">15 Lock和Condition（下）：Dubbo如何用管程实现异步转同步？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Semaphore：如何快速实现一个限流器？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/16%20Semaphore%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e9%99%90%e6%b5%81%e5%99%a8%ef%bc%9f.md">16 Semaphore：如何快速实现一个限流器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 ReadWriteLock：如何快速实现一个完备的缓存？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/17%20ReadWriteLock%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e5%ae%8c%e5%a4%87%e7%9a%84%e7%bc%93%e5%ad%98%ef%bc%9f.md">17 ReadWriteLock：如何快速实现一个完备的缓存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 StampedLock：有没有比读写锁更快的锁？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/18%20StampedLock%ef%bc%9a%e6%9c%89%e6%b2%a1%e6%9c%89%e6%af%94%e8%af%bb%e5%86%99%e9%94%81%e6%9b%b4%e5%bf%ab%e7%9a%84%e9%94%81%ef%bc%9f.md">18 StampedLock：有没有比读写锁更快的锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 CountDownLatch和CyclicBarrier：如何让多线程步调一致？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/19%20CountDownLatch%e5%92%8cCyclicBarrier%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e5%a4%9a%e7%ba%bf%e7%a8%8b%e6%ad%a5%e8%b0%83%e4%b8%80%e8%87%b4%ef%bc%9f.md">19 CountDownLatch和CyclicBarrier：如何让多线程步调一致？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 并发容器：都有哪些“坑”需要我们填？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/20%20%e5%b9%b6%e5%8f%91%e5%ae%b9%e5%99%a8%ef%bc%9a%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e2%80%9c%e5%9d%91%e2%80%9d%e9%9c%80%e8%a6%81%e6%88%91%e4%bb%ac%e5%a1%ab%ef%bc%9f.md">20 并发容器：都有哪些“坑”需要我们填？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 原子类：无锁工具类的典范.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/21%20%e5%8e%9f%e5%ad%90%e7%b1%bb%ef%bc%9a%e6%97%a0%e9%94%81%e5%b7%a5%e5%85%b7%e7%b1%bb%e7%9a%84%e5%85%b8%e8%8c%83.md">21 原子类：无锁工具类的典范.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Executor与线程池：如何创建正确的线程池？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/22%20Executor%e4%b8%8e%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%9b%e5%bb%ba%e6%ad%a3%e7%a1%ae%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f.md">22 Executor与线程池：如何创建正确的线程池？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Future：如何用多线程实现最优的“烧水泡茶”程序？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/23%20Future%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e5%a4%9a%e7%ba%bf%e7%a8%8b%e5%ae%9e%e7%8e%b0%e6%9c%80%e4%bc%98%e7%9a%84%e2%80%9c%e7%83%a7%e6%b0%b4%e6%b3%a1%e8%8c%b6%e2%80%9d%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">23 Future：如何用多线程实现最优的“烧水泡茶”程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 CompletableFuture：异步编程没那么难.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/24%20CompletableFuture%ef%bc%9a%e5%bc%82%e6%ad%a5%e7%bc%96%e7%a8%8b%e6%b2%a1%e9%82%a3%e4%b9%88%e9%9a%be.md">24 CompletableFuture：异步编程没那么难.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 CompletionService：如何批量执行异步任务？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/25%20CompletionService%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%b9%e9%87%8f%e6%89%a7%e8%a1%8c%e5%bc%82%e6%ad%a5%e4%bb%bb%e5%8a%a1%ef%bc%9f.md">25 CompletionService：如何批量执行异步任务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 Fork_Join：单机版的MapReduce.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/26%20Fork_Join%ef%bc%9a%e5%8d%95%e6%9c%ba%e7%89%88%e7%9a%84MapReduce.md">26 Fork_Join：单机版的MapReduce.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 并发工具类模块热点问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/27%20%e5%b9%b6%e5%8f%91%e5%b7%a5%e5%85%b7%e7%b1%bb%e6%a8%a1%e5%9d%97%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">27 并发工具类模块热点问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 Immutability模式：如何利用不变性解决并发问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/28%20Immutability%e6%a8%a1%e5%bc%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e4%b8%8d%e5%8f%98%e6%80%a7%e8%a7%a3%e5%86%b3%e5%b9%b6%e5%8f%91%e9%97%ae%e9%a2%98%ef%bc%9f.md">28 Immutability模式：如何利用不变性解决并发问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 Copy-on-Write模式：不是延时策略的COW.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/29%20Copy-on-Write%e6%a8%a1%e5%bc%8f%ef%bc%9a%e4%b8%8d%e6%98%af%e5%bb%b6%e6%97%b6%e7%ad%96%e7%95%a5%e7%9a%84COW.md">29 Copy-on-Write模式：不是延时策略的COW.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="3 个用户来信 打开一个新的并发世界.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/3%20%e4%b8%aa%e7%94%a8%e6%88%b7%e6%9d%a5%e4%bf%a1%20%e6%89%93%e5%bc%80%e4%b8%80%e4%b8%aa%e6%96%b0%e7%9a%84%e5%b9%b6%e5%8f%91%e4%b8%96%e7%95%8c.md">3 个用户来信 打开一个新的并发世界.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 线程本地存储模式：没有共享，就没有伤害.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/30%20%e7%ba%bf%e7%a8%8b%e6%9c%ac%e5%9c%b0%e5%ad%98%e5%82%a8%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%b2%a1%e6%9c%89%e5%85%b1%e4%ba%ab%ef%bc%8c%e5%b0%b1%e6%b2%a1%e6%9c%89%e4%bc%a4%e5%ae%b3.md">30 线程本地存储模式：没有共享，就没有伤害.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 Guarded Suspension模式：等待唤醒机制的规范实现.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/31%20Guarded%20Suspension%e6%a8%a1%e5%bc%8f%ef%bc%9a%e7%ad%89%e5%be%85%e5%94%a4%e9%86%92%e6%9c%ba%e5%88%b6%e7%9a%84%e8%a7%84%e8%8c%83%e5%ae%9e%e7%8e%b0.md">31 Guarded Suspension模式：等待唤醒机制的规范实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 Balking模式：再谈线程安全的单例模式.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/32%20Balking%e6%a8%a1%e5%bc%8f%ef%bc%9a%e5%86%8d%e8%b0%88%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e7%9a%84%e5%8d%95%e4%be%8b%e6%a8%a1%e5%bc%8f.md">32 Balking模式：再谈线程安全的单例模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 Thread-Per-Message模式：最简单实用的分工方法.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/33%20Thread-Per-Message%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%9c%80%e7%ae%80%e5%8d%95%e5%ae%9e%e7%94%a8%e7%9a%84%e5%88%86%e5%b7%a5%e6%96%b9%e6%b3%95.md">33 Thread-Per-Message模式：最简单实用的分工方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 Worker Thread模式：如何避免重复创建线程？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/34%20Worker%20Thread%e6%a8%a1%e5%bc%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e9%87%8d%e5%a4%8d%e5%88%9b%e5%bb%ba%e7%ba%bf%e7%a8%8b%ef%bc%9f.md">34 Worker Thread模式：如何避免重复创建线程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 两阶段终止模式：如何优雅地终止线程？.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/35%20%e4%b8%a4%e9%98%b6%e6%ae%b5%e7%bb%88%e6%ad%a2%e6%a8%a1%e5%bc%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e9%9b%85%e5%9c%b0%e7%bb%88%e6%ad%a2%e7%ba%bf%e7%a8%8b%ef%bc%9f.md">35 两阶段终止模式：如何优雅地终止线程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 生产者-消费者模式：用流水线思想提高效率.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/36%20%e7%94%9f%e4%ba%a7%e8%80%85-%e6%b6%88%e8%b4%b9%e8%80%85%e6%a8%a1%e5%bc%8f%ef%bc%9a%e7%94%a8%e6%b5%81%e6%b0%b4%e7%ba%bf%e6%80%9d%e6%83%b3%e6%8f%90%e9%ab%98%e6%95%88%e7%8e%87.md">36 生产者-消费者模式：用流水线思想提高效率.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 设计模式模块热点问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/37%20%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%e6%a8%a1%e5%9d%97%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">37 设计模式模块热点问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 案例分析（一）：高性能限流器Guava RateLimiter.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/38%20%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e9%99%90%e6%b5%81%e5%99%a8Guava%20RateLimiter.md">38 案例分析（一）：高性能限流器Guava RateLimiter.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 案例分析（二）：高性能网络应用框架Netty.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/39%20%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e7%bd%91%e7%bb%9c%e5%ba%94%e7%94%a8%e6%a1%86%e6%9e%b6Netty.md">39 案例分析（二）：高性能网络应用框架Netty.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 案例分析（三）：高性能队列Disruptor.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/40%20%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e9%98%9f%e5%88%97Disruptor.md">40 案例分析（三）：高性能队列Disruptor.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 案例分析（四）：高性能数据库连接池HiKariCP.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/41%20%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e6%95%b0%e6%8d%ae%e5%ba%93%e8%bf%9e%e6%8e%a5%e6%b1%a0HiKariCP.md">41 案例分析（四）：高性能数据库连接池HiKariCP.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 Actor模型：面向对象原生的并发模型.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/42%20Actor%e6%a8%a1%e5%9e%8b%ef%bc%9a%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e5%8e%9f%e7%94%9f%e7%9a%84%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%9e%8b.md">42 Actor模型：面向对象原生的并发模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 软件事务内存：借鉴数据库的并发经验.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/43%20%e8%bd%af%e4%bb%b6%e4%ba%8b%e5%8a%a1%e5%86%85%e5%ad%98%ef%bc%9a%e5%80%9f%e9%89%b4%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e5%b9%b6%e5%8f%91%e7%bb%8f%e9%aa%8c.md">43 软件事务内存：借鉴数据库的并发经验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 协程：更轻量级的线程.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/44%20%e5%8d%8f%e7%a8%8b%ef%bc%9a%e6%9b%b4%e8%bd%bb%e9%87%8f%e7%ba%a7%e7%9a%84%e7%ba%bf%e7%a8%8b.md">44 协程：更轻量级的线程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 CSP模型：Golang的主力队员.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/45%20CSP%e6%a8%a1%e5%9e%8b%ef%bc%9aGolang%e7%9a%84%e4%b8%bb%e5%8a%9b%e9%98%9f%e5%91%98.md">45 CSP模型：Golang的主力队员.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户来信 真好，面试考到这些并发编程，我都答对了！.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/%e7%94%a8%e6%88%b7%e6%9d%a5%e4%bf%a1%20%e7%9c%9f%e5%a5%bd%ef%bc%8c%e9%9d%a2%e8%af%95%e8%80%83%e5%88%b0%e8%bf%99%e4%ba%9b%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%ef%bc%8c%e6%88%91%e9%83%bd%e7%ad%94%e5%af%b9%e4%ba%86%ef%bc%81.md">用户来信 真好，面试考到这些并发编程，我都答对了！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 十年之后，初心依旧.md" href="/%e4%b8%93%e6%a0%8f/Java%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%8d%81%e5%b9%b4%e4%b9%8b%e5%90%8e%ef%bc%8c%e5%88%9d%e5%bf%83%e4%be%9d%e6%97%a7.md">结束语 十年之后，初心依旧.md</a>
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
                            <h1 id="title" data-id="39 案例分析（二）：高性能网络应用框架Netty" class="title">39 案例分析（二）：高性能网络应用框架Netty</h1>
                            <div><p>Netty是一个高性能网络应用框架，应用非常普遍，目前在Java领域里，Netty基本上成为网络程序的标配了。Netty框架功能丰富，也非常复杂，今天我们主要分析Netty框架中的线程模型，而<strong>线程模型直接影响着网络程序的性能</strong>。</p>

<p>在介绍Netty的线程模型之前，我们首先需要把问题搞清楚，了解网络编程性能的瓶颈在哪里，然后再看Netty的线程模型是如何解决这个问题的。</p>

<h2 id="网络编程性能的瓶颈">网络编程性能的瓶颈</h2>

<p>在<a href="https://time.geekbang.org/column/article/95098" target="_blank">《33 | Thread-Per-Message模式：最简单实用的分工方法》</a>中，我们写过一个简单的网络程序echo，采用的是阻塞式I/O（BIO）。BIO模型里，所有read()操作和write()操作都会阻塞当前线程的，如果客户端已经和服务端建立了一个连接，而迟迟不发送数据，那么服务端的read()操作会一直阻塞，所以<strong>使用BIO模型，一般都会为每个socket分配一个独立的线程</strong>，这样就不会因为线程阻塞在一个socket上而影响对其他socket的读写。BIO的线程模型如下图所示，每一个socket都对应一个独立的线程；为了避免频繁创建、消耗线程，可以采用线程池，但是socket和线程之间的对应关系并不会变化。</p>

<p><img src="assets/e712c37ea0483e9dde0d6efe76e687e2.png" alt="" /></p>

<p>BIO的线程模型</p>

<p>BIO这种线程模型适用于socket连接不是很多的场景；但是现在的互联网场景，往往需要服务器能够支撑十万甚至百万连接，而创建十万甚至上百万个线程显然并不现实，所以BIO线程模型无法解决百万连接的问题。如果仔细观察，你会发现互联网场景中，虽然连接多，但是每个连接上的请求并不频繁，所以线程大部分时间都在等待I/O就绪。也就是说线程大部分时间都阻塞在那里，这完全是浪费，如果我们能够解决这个问题，那就不需要这么多线程了。</p>

<p>顺着这个思路，我们可以将线程模型优化为下图这个样子，可以用一个线程来处理多个连接，这样线程的利用率就上来了，同时所需的线程数量也跟着降下来了。这个思路很好，可是使用BIO相关的API是无法实现的，这是为什么呢？因为BIO相关的socket读写操作都是阻塞式的，而一旦调用了阻塞式API，在I/O就绪前，调用线程会一直阻塞，也就无法处理其他的socket连接了。</p>

<p><img src="assets/eafed0787b82b0b428e1ec0927029f1f.png" alt="" /></p>

<p>理想的线程模型图</p>

<p>好在Java里还提供了非阻塞式（NIO）API，<strong>利用非阻塞式API就能够实现一个线程处理多个连接了</strong>。那具体如何实现呢？现在普遍都是<strong>采用Reactor模式</strong>，包括Netty的实现。所以，要想理解Netty的实现，接下来我们就需要先了解一下Reactor模式。</p>

<h2 id="reactor模式">Reactor模式</h2>

<p>下面是Reactor模式的类结构图，其中Handle指的是I/O句柄，在Java网络编程里，它本质上就是一个网络连接。Event Handler很容易理解，就是一个事件处理器，其中handle_event()方法处理I/O事件，也就是每个Event Handler处理一个I/O Handle；get_handle()方法可以返回这个I/O的Handle。Synchronous Event Demultiplexer可以理解为操作系统提供的I/O多路复用API，例如POSIX标准里的select()以及Linux里面的epoll()。</p>

<p><img src="assets/a7ba3c8d6c49e50d9288baf0c03fa240.png" alt="" /></p>

<p>Reactor模式类结构图</p>

<p>Reactor模式的核心自然是<strong>Reactor这个类</strong>，其中register_handler()和remove_handler()这两个方法可以注册和删除一个事件处理器；<strong>handle_events()方式是核心</strong>，也是Reactor模式的发动机，这个方法的核心逻辑如下：首先通过同步事件多路选择器提供的select()方法监听网络事件，当有网络事件就绪后，就遍历事件处理器来处理该网络事件。由于网络事件是源源不断的，所以在主程序中启动Reactor模式，需要以 <code>while(true){}</code> 的方式调用handle_events()方法。</p>

<pre><code>void Reactor::handle_events(){
  //通过同步事件多路选择器提供的
  //select()方法监听网络事件
  select(handlers);
  //处理网络事件
  for(h in handlers){
    h.handle_event();
  }
}
// 在主程序中启动事件循环
while (true) {
  handle_events();
</code></pre>

<h2 id="netty中的线程模型">Netty中的线程模型</h2>

<p>Netty的实现虽然参考了Reactor模式，但是并没有完全照搬，<strong>Netty中最核心的概念是事件循环（EventLoop）</strong>，其实也就是Reactor模式中的Reactor，<strong>负责监听网络事件并调用事件处理器进行处理</strong>。在4.x版本的Netty中，网络连接和EventLoop是稳定的多对1关系，而EventLoop和Java线程是1对1关系，这里的稳定指的是关系一旦确定就不再发生变化。也就是说一个网络连接只会对应唯一的一个EventLoop，而一个EventLoop也只会对应到一个Java线程，所以<strong>一个网络连接只会对应到一个Java线程</strong>。</p>

<p>一个网络连接对应到一个Java线程上，有什么好处呢？最大的好处就是对于一个网络连接的事件处理是单线程的，这样就<strong>避免了各种并发问题</strong>。</p>

<p>Netty中的线程模型可以参考下图，这个图和前面我们提到的理想的线程模型图非常相似，核心目标都是用一个线程处理多个网络连接。</p>

<p><img src="assets/034756f1d76bb3af09e125de9f3c2f04.png" alt="" /></p>

<p>Netty中的线程模型</p>

<p>Netty中还有一个核心概念是<strong>EventLoopGroup</strong>，顾名思义，一个EventLoopGroup由一组EventLoop组成。实际使用中，一般都会创建两个EventLoopGroup，一个称为bossGroup，一个称为workerGroup。为什么会有两个EventLoopGroup呢？</p>

<p>这个和socket处理网络请求的机制有关，socket处理TCP网络连接请求，是在一个独立的socket中，每当有一个TCP连接成功建立，都会创建一个新的socket，之后对TCP连接的读写都是由新创建处理的socket完成的。也就是说<strong>处理TCP连接请求和读写请求是通过两个不同的socket完成的</strong>。上面我们在讨论网络请求的时候，为了简化模型，只是讨论了读写请求，而没有讨论连接请求。</p>

<p><strong>在Netty中，bossGroup就用来处理连接请求的，而workerGroup是用来处理读写请求的</strong>。bossGroup处理完连接请求后，会将这个连接提交给workerGroup来处理， workerGroup里面有多个EventLoop，那新的连接会交给哪个EventLoop来处理呢？这就需要一个负载均衡算法，Netty中目前使用的是<strong>轮询算法</strong>。</p>

<p>下面我们用Netty重新实现以下echo程序的服务端，近距离感受一下Netty。</p>

<h2 id="用netty实现echo程序服务端">用Netty实现Echo程序服务端</h2>

<p>下面的示例代码基于Netty实现了echo程序服务端：首先创建了一个事件处理器（等同于Reactor模式中的事件处理器），然后创建了bossGroup和workerGroup，再之后创建并初始化了ServerBootstrap，代码还是很简单的，不过有两个地方需要注意一下。</p>

<p>第一个，如果NettybossGroup只监听一个端口，那bossGroup只需要1个EventLoop就可以了，多了纯属浪费。</p>

<p>第二个，默认情况下，Netty会创建“2*CPU核数”个EventLoop，由于网络连接与EventLoop有稳定的关系，所以事件处理器在处理网络事件的时候是不能有阻塞操作的，否则很容易导致请求大面积超时。如果实在无法避免使用阻塞操作，那可以通过线程池来异步处理。</p>

<pre><code>//事件处理器
final EchoServerHandler serverHandler 
  = new EchoServerHandler();
//boss线程组  
EventLoopGroup bossGroup 
  = new NioEventLoopGroup(1); 
//worker线程组  
EventLoopGroup workerGroup 
  = new NioEventLoopGroup();
try {
  ServerBootstrap b = new ServerBootstrap();
  b.group(bossGroup, workerGroup)
   .channel(NioServerSocketChannel.class)
   .childHandler(new ChannelInitializer&lt;SocketChannel&gt;() {
     @Override
     public void initChannel(SocketChannel ch){
       ch.pipeline().addLast(serverHandler);
     }
    });
  //bind服务端端口  
  ChannelFuture f = b.bind(9090).sync();
  f.channel().closeFuture().sync();
} finally {
  //终止工作线程组
  workerGroup.shutdownGracefully();
  //终止boss线程组
  bossGroup.shutdownGracefully();
}

//socket连接处理器
class EchoServerHandler extends 
    ChannelInboundHandlerAdapter {
  //处理读事件  
  @Override
  public void channelRead(
    ChannelHandlerContext ctx, Object msg){
      ctx.write(msg);
  }
  //处理读完成事件
  @Override
  public void channelReadComplete(
    ChannelHandlerContext ctx){
      ctx.flush();
  }
  //处理异常事件
  @Override
  public void exceptionCaught(
    ChannelHandlerContext ctx,  Throwable cause) {
      cause.printStackTrace();
      ctx.close();
  }
}
</code></pre>

<h2 id="总结">总结</h2>

<p>Netty是一个款优秀的网络编程框架，性能非常好，为了实现高性能的目标，Netty做了很多优化，例如优化了ByteBuffer、支持零拷贝等等，和并发编程相关的就是它的线程模型了。Netty的线程模型设计得很精巧，每个网络连接都关联到了一个线程上，这样做的好处是：对于一个网络连接，读写操作都是单线程执行的，从而避免了并发程序的各种问题。</p>

<p>你要想深入理解Netty的线程模型，还需要对网络相关知识有一定的理解，关于Java IO的演进过程，你可以参考<a href="http://gee.cs.oswego.edu/dl/cpjslides/nio.pdf" target="_blank">Scalable IO in Java</a>，至于TCP/IP网络编程的知识你可以参考韩国尹圣雨写的经典教程——《TCP/IP网络编程》。</p>

<p>欢迎在留言区与我分享你的想法，也欢迎你在留言区记录你的思考过程。感谢阅读，如果你觉得这篇文章对你有帮助的话，也欢迎把它分享给更多的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#96fafafaafa2a7a7a6a1d6f1fbf7fffab8f5f9fb" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f551b3deccd81',t:'MTczNDAyMzQ3NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
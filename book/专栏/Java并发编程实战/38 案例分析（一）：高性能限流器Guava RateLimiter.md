<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=38&#32;案例分析（一）：高性能限流器Guava&#32;RateLimiter>
        <link rel="icon" href="/static/favicon.png">
        <title>38 案例分析（一）：高性能限流器Guava RateLimiter </title>
        
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
                            <h1 id="title" data-id="38 案例分析（一）：高性能限流器Guava RateLimiter" class="title">38 案例分析（一）：高性能限流器Guava RateLimiter</h1>
                            <div><p>从今天开始，我们就进入案例分析模块了。 这个模块我们将分析四个经典的开源框架，看看它们是如何处理并发问题的，通过这四个案例的学习，相信你会对如何解决并发问题有个更深入的认识。</p>

<p>首先我们来看看<strong>Guava RateLimiter是如何解决高并发场景下的限流问题的</strong>。Guava是Google开源的Java类库，提供了一个工具类RateLimiter。我们先来看看RateLimiter的使用，让你对限流有个感官的印象。假设我们有一个线程池，它每秒只能处理两个任务，如果提交的任务过快，可能导致系统不稳定，这个时候就需要用到限流。</p>

<p>在下面的示例代码中，我们创建了一个流速为2个请求/秒的限流器，这里的流速该怎么理解呢？直观地看，2个请求/秒指的是每秒最多允许2个请求通过限流器，其实在Guava中，流速还有更深一层的意思：是一种匀速的概念，2个请求/秒等价于1个请求/500毫秒。</p>

<p>在向线程池提交任务之前，调用 <code>acquire()</code> 方法就能起到限流的作用。通过示例代码的执行结果，任务提交到线程池的时间间隔基本上稳定在500毫秒。</p>

<pre><code>//限流器流速：2个请求/秒
RateLimiter limiter = 
  RateLimiter.create(2.0);
//执行任务的线程池
ExecutorService es = Executors
  .newFixedThreadPool(1);
//记录上一次执行时间
prev = System.nanoTime();
//测试执行20次
for (int i=0; i&lt;20; i++){
  //限流器限流
  limiter.acquire();
  //提交任务异步执行
  es.execute(()-&gt;{
    long cur=System.nanoTime();
    //打印时间间隔：毫秒
    System.out.println(
      (cur-prev)/1000_000);
    prev = cur;
  });
}

输出结果：
...
500
499
499
500
499
</code></pre>

<h2 id="经典限流算法-令牌桶算法">经典限流算法：令牌桶算法</h2>

<p>Guava的限流器使用上还是很简单的，那它是如何实现的呢？Guava采用的是<strong>令牌桶算法</strong>，其<strong>核心是要想通过限流器，必须拿到令牌</strong>。也就是说，只要我们能够限制发放令牌的速率，那么就能控制流速了。令牌桶算法的详细描述如下：</p>

<ol>
<li>令牌以固定的速率添加到令牌桶中，假设限流的速率是 r/秒，则令牌每 1/r 秒会添加一个；</li>
<li>假设令牌桶的容量是 b ，如果令牌桶已满，则新的令牌会被丢弃；</li>
<li>请求能够通过限流器的前提是令牌桶中有令牌。</li>
</ol>

<p>这个算法中，限流的速率 r 还是比较容易理解的，但令牌桶的容量 b 该怎么理解呢？b 其实是burst的简写，意义是<strong>限流器允许的最大突发流量</strong>。比如b=10，而且令牌桶中的令牌已满，此时限流器允许10个请求同时通过限流器，当然只是突发流量而已，这10个请求会带走10个令牌，所以后续的流量只能按照速率 r 通过限流器。</p>

<p>令牌桶这个算法，如何用Java实现呢？很可能你的直觉会告诉你生产者-消费者模式：一个生产者线程定时向阻塞队列中添加令牌，而试图通过限流器的线程则作为消费者线程，只有从阻塞队列中获取到令牌，才允许通过限流器。</p>

<p>这个算法看上去非常完美，而且实现起来非常简单，如果并发量不大，这个实现并没有什么问题。可实际情况却是使用限流的场景大部分都是高并发场景，而且系统压力已经临近极限了，此时这个实现就有问题了。问题就出在定时器上，在高并发场景下，当系统压力已经临近极限的时候，定时器的精度误差会非常大，同时定时器本身会创建调度线程，也会对系统的性能产生影响。</p>

<p>那还有什么好的实现方式呢？当然有，Guava的实现就没有使用定时器，下面我们就来看看它是如何实现的。</p>

<h2 id="guava如何实现令牌桶算法">Guava如何实现令牌桶算法</h2>

<p>Guava实现令牌桶算法，用了一个很简单的办法，其关键是<strong>记录并动态计算下一令牌发放的时间</strong>。下面我们以一个最简单的场景来介绍该算法的执行过程。假设令牌桶的容量为 b=1，限流速率 r = 1个请求/秒，如下图所示，如果当前令牌桶中没有令牌，下一个令牌的发放时间是在第3秒，而在第2秒的时候有一个线程T1请求令牌，此时该如何处理呢？</p>

<p><img src="assets/391179821a55fc798c9c17a6991c1dce.png" alt="" /></p>

<p>线程T1请求令牌示意图</p>

<p>对于这个请求令牌的线程而言，很显然需要等待1秒，因为1秒以后（第3秒）它就能拿到令牌了。此时需要注意的是，下一个令牌发放的时间也要增加1秒，为什么呢？因为第3秒发放的令牌已经被线程T1预占了。处理之后如下图所示。</p>

<p><img src="assets/1a4069c830e18de087ba7f490aa78087.png" alt="" /></p>

<p>线程T1请求结束示意图</p>

<p>假设T1在预占了第3秒的令牌之后，马上又有一个线程T2请求令牌，如下图所示。</p>

<p><img src="assets/2cf695d0888a93e1e2d020d9514f5a2e.png" alt="" /></p>

<p>线程T2请求令牌示意图</p>

<p>很显然，由于下一个令牌产生的时间是第4秒，所以线程T2要等待两秒的时间，才能获取到令牌，同时由于T2预占了第4秒的令牌，所以下一令牌产生时间还要增加1秒，完全处理之后，如下图所示。</p>

<p><img src="assets/68c09a96049aacda7936c52b801c22f7.png" alt="" /></p>

<p>线程T2请求结束示意图</p>

<p>上面线程T1、T2都是在<strong>下一令牌产生时间之前</strong>请求令牌，如果线程在<strong>下一令牌产生时间之后</strong>请求令牌会如何呢？假设在线程T1请求令牌之后的5秒，也就是第7秒，线程T3请求令牌，如下图所示。</p>

<p><img src="assets/e3125d72eb3d84eabf6de6ab987e695c.png" alt="" /></p>

<p>线程T3请求令牌示意图</p>

<p>由于在第5秒已经产生了一个令牌，所以此时线程T3可以直接拿到令牌，而无需等待。在第7秒，实际上限流器能够产生3个令牌，第5、6、7秒各产生一个令牌。由于我们假设令牌桶的容量是1，所以第6、7秒产生的令牌就丢弃了，其实等价地你也可以认为是保留的第7秒的令牌，丢弃的第5、6秒的令牌，也就是说第7秒的令牌被线程T3占有了，于是下一令牌的的产生时间应该是第8秒，如下图所示。</p>

<p><img src="assets/baf159d05b2abf650839e29a2399a4fc.png" alt="" /></p>

<p>线程T3请求结束示意图</p>

<p>通过上面简要地分析，你会发现，我们<strong>只需要记录一个下一令牌产生的时间，并动态更新它，就能够轻松完成限流功能</strong>。我们可以将上面的这个算法代码化，示例代码如下所示，依然假设令牌桶的容量是1。关键是<strong>reserve()方法</strong>，这个方法会为请求令牌的线程预分配令牌，同时返回该线程能够获取令牌的时间。其实现逻辑就是上面提到的：如果线程请求令牌的时间在下一令牌产生时间之后，那么该线程立刻就能够获取令牌；反之，如果请求时间在下一令牌产生时间之前，那么该线程是在下一令牌产生的时间获取令牌。由于此时下一令牌已经被该线程预占，所以下一令牌产生的时间需要加上1秒。</p>

<pre><code>class SimpleLimiter {
  //下一令牌产生时间
  long next = System.nanoTime();
  //发放令牌间隔：纳秒
  long interval = 1000_000_000;
  //预占令牌，返回能够获取令牌的时间
  synchronized long reserve(long now){
    //请求时间在下一令牌产生时间之后
    //重新计算下一令牌产生时间
    if (now &gt; next){
      //将下一令牌产生时间重置为当前时间
      next = now;
    }
    //能够获取令牌的时间
    long at=next;
    //设置下一令牌产生时间
    next += interval;
    //返回线程需要等待的时间
    return Math.max(at, 0L);
  }
  //申请令牌
  void acquire() {
    //申请令牌时的时间
    long now = System.nanoTime();
    //预占令牌
    long at=reserve(now);
    long waitTime=max(at-now, 0);
    //按照条件等待
    if(waitTime &gt; 0) {
      try {
        TimeUnit.NANOSECONDS
          .sleep(waitTime);
      }catch(InterruptedException e){
        e.printStackTrace();
      }
    }
  }
}
</code></pre>

<p>如果令牌桶的容量大于1，又该如何处理呢？按照令牌桶算法，令牌要首先从令牌桶中出，所以我们需要按需计算令牌桶中的数量，当有线程请求令牌时，先从令牌桶中出。具体的代码实现如下所示。我们增加了一个<strong>resync()方法</strong>，在这个方法中，如果线程请求令牌的时间在下一令牌产生时间之后，会重新计算令牌桶中的令牌数，<strong>新产生的令牌的计算公式是：(now-next)/interval</strong>，你可对照上面的示意图来理解。reserve()方法中，则增加了先从令牌桶中出令牌的逻辑，不过需要注意的是，如果令牌是从令牌桶中出的，那么next就无需增加一个 interval 了。</p>

<pre><code>class SimpleLimiter {
  //当前令牌桶中的令牌数量
  long storedPermits = 0;
  //令牌桶的容量
  long maxPermits = 3;
  //下一令牌产生时间
  long next = System.nanoTime();
  //发放令牌间隔：纳秒
  long interval = 1000_000_000;

  //请求时间在下一令牌产生时间之后,则
  // 1.重新计算令牌桶中的令牌数
  // 2.将下一个令牌发放时间重置为当前时间
  void resync(long now) {
    if (now &gt; next) {
      //新产生的令牌数
      long newPermits=(now-next)/interval;
      //新令牌增加到令牌桶
      storedPermits=min(maxPermits, 
        storedPermits + newPermits);
      //将下一个令牌发放时间重置为当前时间
      next = now;
    }
  }
  //预占令牌，返回能够获取令牌的时间
  synchronized long reserve(long now){
    resync(now);
    //能够获取令牌的时间
    long at = next;
    //令牌桶中能提供的令牌
    long fb=min(1, storedPermits);
    //令牌净需求：首先减掉令牌桶中的令牌
    long nr = 1 - fb;
    //重新计算下一令牌产生时间
    next = next + nr*interval;
    //重新计算令牌桶中的令牌
    this.storedPermits -= fb;
    return at;
  }
  //申请令牌
  void acquire() {
    //申请令牌时的时间
    long now = System.nanoTime();
    //预占令牌
    long at=reserve(now);
    long waitTime=max(at-now, 0);
    //按照条件等待
    if(waitTime &gt; 0) {
      try {
        TimeUnit.NANOSECONDS
          .sleep(waitTime);
      }catch(InterruptedException e){
        e.printStackTrace();
      }
    }
  }
}
</code></pre>

<h2 id="总结">总结</h2>

<p>经典的限流算法有两个，一个是<strong>令牌桶算法（Token Bucket）</strong>，另一个是<strong>漏桶算法（Leaky Bucket）</strong>。令牌桶算法是定时向令牌桶发送令牌，请求能够从令牌桶中拿到令牌，然后才能通过限流器；而漏桶算法里，请求就像水一样注入漏桶，漏桶会按照一定的速率自动将水漏掉，只有漏桶里还能注入水的时候，请求才能通过限流器。令牌桶算法和漏桶算法很像一个硬币的正反面，所以你可以参考令牌桶算法的实现来实现漏桶算法。</p>

<p>上面我们介绍了Guava是如何实现令牌桶算法的，我们的示例代码是对Guava RateLimiter的简化，Guava RateLimiter扩展了标准的令牌桶算法，比如还能支持预热功能。对于按需加载的缓存来说，预热后缓存能支持5万TPS的并发，但是在预热前5万TPS的并发直接就把缓存击垮了，所以如果需要给该缓存限流，限流器也需要支持预热功能，在初始阶段，限制的流速 r 很小，但是动态增长的。预热功能的实现非常复杂，Guava构建了一个积分函数来解决这个问题，如果你感兴趣，可以继续深入研究。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#224e4e4e1b161313121562454f434b4e0c414d4f" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f54d7185dcd81',t:'MTczNDAyMzQ2My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
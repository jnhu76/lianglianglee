<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=40&#32;案例分析（三）：高性能队列Disruptor>
        <link rel="icon" href="/static/favicon.png">
        <title>40 案例分析（三）：高性能队列Disruptor </title>
        
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
                            <h1 id="title" data-id="40 案例分析（三）：高性能队列Disruptor" class="title">40 案例分析（三）：高性能队列Disruptor</h1>
                            <div><p>我们在<a href="https://time.geekbang.org/column/article/90201" target="_blank">《20 | 并发容器：都有哪些“坑”需要我们填？》</a>介绍过Java SDK提供了2个有界队列：ArrayBlockingQueue 和 LinkedBlockingQueue，它们都是基于ReentrantLock实现的，在高并发场景下，锁的效率并不高，那有没有更好的替代品呢？有，今天我们就介绍一种性能更高的有界队列：Disruptor。</p>

<p><strong>Disruptor是一款高性能的有界内存队列</strong>，目前应用非常广泛，Log4j2、Spring Messaging、HBase、Storm都用到了Disruptor，那Disruptor的性能为什么这么高呢？Disruptor项目团队曾经写过一篇论文，详细解释了其原因，可以总结为如下：</p>

<ol>
<li>内存分配更加合理，使用RingBuffer数据结构，数组元素在初始化时一次性全部创建，提升缓存命中率；对象循环利用，避免频繁GC。</li>
<li>能够避免伪共享，提升缓存利用率。</li>
<li>采用无锁算法，避免频繁加锁、解锁的性能消耗。</li>
<li>支持批量消费，消费者可以无锁方式消费多个消息。</li>
</ol>

<p>其中，前三点涉及到的知识比较多，所以今天咱们重点讲解前三点，不过在详细介绍这些知识之前，我们先来聊聊Disruptor如何使用，好让你先对Disruptor有个感官的认识。</p>

<p>下面的代码出自官方示例，我略做了一些修改，相较而言，Disruptor的使用比Java SDK提供BlockingQueue要复杂一些，但是总体思路还是一致的，其大致情况如下：</p>

<ul>
<li><p>在Disruptor中，生产者生产的对象（也就是消费者消费的对象）称为Event，使用Disruptor必须自定义Event，例如示例代码的自定义Event是LongEvent；</p></li>

<li><p>构建Disruptor对象除了要指定队列大小外，还需要传入一个EventFactory，示例代码中传入的是<code>LongEvent::new</code>；</p></li>

<li><p>消费Disruptor中的Event需要通过handleEventsWith()方法注册一个事件处理器，发布Event则需要通过publishEvent()方法。</p>

<p>//自定义Event
class LongEvent {
  private long value;
  public void set(long value) {</p>

<pre><code>this.value = value;
</code></pre>
<p>}
}
//指定RingBuffer大小,
//必须是2的N次方
int bufferSize = 1024;</p>

<p>//构建Disruptor
Disruptor<LongEvent> disruptor
  = new Disruptor&lt;&gt;(</p>

<pre><code>LongEvent::new,
bufferSize,
DaemonThreadFactory.INSTANCE);
</code></pre>
<p>//注册事件处理器
disruptor.handleEventsWith(
  (event, sequence, endOfBatch) -&gt;</p>

<pre><code>System.out.println(&quot;E: &quot;+event));
</code></pre>
<p>//启动Disruptor
disruptor.start();</p>

<p>//获取RingBuffer
RingBuffer<LongEvent> ringBuffer
  = disruptor.getRingBuffer();
//生产Event
ByteBuffer bb = ByteBuffer.allocate(8);
for (long l = 0; true; l++){
  bb.putLong(0, l);
  //生产者生产消息
  ringBuffer.publishEvent(</p>

<pre><code>(event, sequence, buffer) -&gt; 
  event.set(buffer.getLong(0)), bb);
</code></pre>
<p>Thread.sleep(1000);
}</p></li>
</ul>

<h2 id="ringbuffer如何提升性能">RingBuffer如何提升性能</h2>

<p>Java SDK中ArrayBlockingQueue使用<strong>数组</strong>作为底层的数据存储，而Disruptor是使用<strong>RingBuffer</strong>作为数据存储。RingBuffer本质上也是数组，所以仅仅将数据存储从数组换成RingBuffer并不能提升性能，但是Disruptor在RingBuffer的基础上还做了很多优化，其中一项优化就是和内存分配有关的。</p>

<p>在介绍这项优化之前，你需要先了解一下程序的局部性原理。简单来讲，<strong>程序的局部性原理指的是在一段时间内程序的执行会限定在一个局部范围内</strong>。这里的“局部性”可以从两个方面来理解，一个是时间局部性，另一个是空间局部性。<strong>时间局部性</strong>指的是程序中的某条指令一旦被执行，不久之后这条指令很可能再次被执行；如果某条数据被访问，不久之后这条数据很可能再次被访问。而<strong>空间局部性</strong>是指某块内存一旦被访问，不久之后这块内存附近的内存也很可能被访问。</p>

<p>CPU的缓存就利用了程序的局部性原理：CPU从内存中加载数据X时，会将数据X缓存在高速缓存Cache中，实际上CPU缓存X的同时，还缓存了X周围的数据，因为根据程序具备局部性原理，X周围的数据也很有可能被访问。从另外一个角度来看，如果程序能够很好地体现出局部性原理，也就能更好地利用CPU的缓存，从而提升程序的性能。Disruptor在设计RingBuffer的时候就充分考虑了这个问题，下面我们就对比着ArrayBlockingQueue来分析一下。</p>

<p>首先是ArrayBlockingQueue。生产者线程向ArrayBlockingQueue增加一个元素，每次增加元素E之前，都需要创建一个对象E，如下图所示，ArrayBlockingQueue内部有6个元素，这6个元素都是由生产者线程创建的，由于创建这些元素的时间基本上是离散的，所以这些元素的内存地址大概率也不是连续的。</p>

<p><img src="assets/848fd30644355ea86f3f91b06bfafa90.png" alt="" /></p>

<p>ArrayBlockingQueue内部结构图</p>

<p>下面我们再看看Disruptor是如何处理的。Disruptor内部的RingBuffer也是用数组实现的，但是这个数组中的所有元素在初始化时是一次性全部创建的，所以这些元素的内存地址大概率是连续的，相关的代码如下所示。</p>

<pre><code>for (int i=0; i&lt;bufferSize; i++){
  //entries[]就是RingBuffer内部的数组
  //eventFactory就是前面示例代码中传入的LongEvent::new
  entries[BUFFER_PAD + i] 
    = eventFactory.newInstance();
}
</code></pre>

<p>Disruptor内部RingBuffer的结构可以简化成下图，那问题来了，数组中所有元素内存地址连续能提升性能吗？能！为什么呢？因为消费者线程在消费的时候，是遵循空间局部性原理的，消费完第1个元素，很快就会消费第2个元素；当消费第1个元素E1的时候，CPU会把内存中E1后面的数据也加载进Cache，如果E1和E2在内存中的地址是连续的，那么E2也就会被加载进Cache中，然后当消费第2个元素的时候，由于E2已经在Cache中了，所以就不需要从内存中加载了，这样就能大大提升性能。</p>

<p><img src="assets/33bc0d35615f5d5f7869871e0cfed037.png" alt="" /></p>

<p>Disruptor内部RingBuffer结构图</p>

<p>除此之外，在Disruptor中，生产者线程通过publishEvent()发布Event的时候，并不是创建一个新的Event，而是通过event.set()方法修改Event， 也就是说RingBuffer创建的Event是可以循环利用的，这样还能避免频繁创建、删除Event导致的频繁GC问题。</p>

<h2 id="如何避免-伪共享">如何避免“伪共享”</h2>

<p>高效利用Cache，能够大大提升性能，所以要努力构建能够高效利用Cache的内存结构。而从另外一个角度看，努力避免不能高效利用Cache的内存结构也同样重要。</p>

<p>有一种叫做“伪共享（False sharing）”的内存布局就会使Cache失效，那什么是“伪共享”呢？</p>

<p>伪共享和CPU内部的Cache有关，Cache内部是按照缓存行（Cache Line）管理的，缓存行的大小通常是64个字节；CPU从内存中加载数据X，会同时加载X后面（64-size(X)）个字节的数据。下面的示例代码出自Java SDK的ArrayBlockingQueue，其内部维护了4个成员变量，分别是队列数组items、出队索引takeIndex、入队索引putIndex以及队列中的元素总数count。</p>

<pre><code>/** 队列数组 */
final Object[] items;
/** 出队索引 */
int takeIndex;
/** 入队索引 */
int putIndex;
/** 队列中元素总数 */
int count;
</code></pre>

<p>当CPU从内存中加载takeIndex的时候，会同时将putIndex以及count都加载进Cache。下图是某个时刻CPU中Cache的状况，为了简化，缓存行中我们仅列出了takeIndex和putIndex。</p>

<p><img src="assets/fdccf96bda79453e55ed75e418864b5c.png" alt="" /></p>

<p>CPU缓存示意图</p>

<p>假设线程A运行在CPU-1上，执行入队操作，入队操作会修改putIndex，而修改putIndex会导致其所在的所有核上的缓存行均失效；此时假设运行在CPU-2上的线程执行出队操作，出队操作需要读取takeIndex，由于takeIndex所在的缓存行已经失效，所以CPU-2必须从内存中重新读取。入队操作本不会修改takeIndex，但是由于takeIndex和putIndex共享的是一个缓存行，就导致出队操作不能很好地利用Cache，这其实就是<strong>伪共享</strong>。简单来讲，<strong>伪共享指的是由于共享缓存行导致缓存无效的场景</strong>。</p>

<p>ArrayBlockingQueue的入队和出队操作是用锁来保证互斥的，所以入队和出队不会同时发生。如果允许入队和出队同时发生，那就会导致线程A和线程B争用同一个缓存行，这样也会导致性能问题。所以为了更好地利用缓存，我们必须避免伪共享，那如何避免呢？</p>

<p><img src="assets/d5d5afc11fe6b1aaf8c9be7dba643827.png" alt="" /></p>

<p>CPU缓存失效示意图</p>

<p>方案很简单，<strong>每个变量独占一个缓存行、不共享缓存行</strong>就可以了，具体技术是<strong>缓存行填充</strong>。比如想让takeIndex独占一个缓存行，可以在takeIndex的前后各填充56个字节，这样就一定能保证takeIndex独占一个缓存行。下面的示例代码出自Disruptor，Sequence 对象中的value属性就能避免伪共享，因为这个属性前后都填充了56个字节。Disruptor中很多对象，例如RingBuffer、RingBuffer内部的数组都用到了这种填充技术来避免伪共享。</p>

<pre><code>//前：填充56字节
class LhsPadding{
    long p1, p2, p3, p4, p5, p6, p7;
}
class Value extends LhsPadding{
    volatile long value;
}
//后：填充56字节
class RhsPadding extends Value{
    long p9, p10, p11, p12, p13, p14, p15;
}
class Sequence extends RhsPadding{
  //省略实现
}
</code></pre>

<h2 id="disruptor中的无锁算法">Disruptor中的无锁算法</h2>

<p>ArrayBlockingQueue是利用管程实现的，中规中矩，生产、消费操作都需要加锁，实现起来简单，但是性能并不十分理想。Disruptor采用的是无锁算法，很复杂，但是核心无非是生产和消费两个操作。Disruptor中最复杂的是入队操作，所以我们重点来看看入队操作是如何实现的。</p>

<p>对于入队操作，最关键的要求是不能覆盖没有消费的元素；对于出队操作，最关键的要求是不能读取没有写入的元素，所以Disruptor中也一定会维护类似出队索引和入队索引这样两个关键变量。Disruptor中的RingBuffer维护了入队索引，但是并没有维护出队索引，这是因为在Disruptor中多个消费者可以同时消费，每个消费者都会有一个出队索引，所以RingBuffer的出队索引是所有消费者里面最小的那一个。</p>

<p>下面是Disruptor生产者入队操作的核心代码，看上去很复杂，其实逻辑很简单：如果没有足够的空余位置，就出让CPU使用权，然后重新计算；反之则用CAS设置入队索引。</p>

<pre><code>//生产者获取n个写入位置
do {
  //cursor类似于入队索引，指的是上次生产到这里
  current = cursor.get();
  //目标是在生产n个
  next = current + n;
  //减掉一个循环
  long wrapPoint = next - bufferSize;
  //获取上一次的最小消费位置
  long cachedGatingSequence = gatingSequenceCache.get();
  //没有足够的空余位置
  if (wrapPoint&gt;cachedGatingSequence || cachedGatingSequence&gt;current){
    //重新计算所有消费者里面的最小值位置
    long gatingSequence = Util.getMinimumSequence(
        gatingSequences, current);
    //仍然没有足够的空余位置，出让CPU使用权，重新执行下一循环
    if (wrapPoint &gt; gatingSequence){
      LockSupport.parkNanos(1);
      continue;
    }
    //从新设置上一次的最小消费位置
    gatingSequenceCache.set(gatingSequence);
  } else if (cursor.compareAndSet(current, next)){
    //获取写入位置成功，跳出循环
    break;
  }
} while (true);
</code></pre>

<h2 id="总结">总结</h2>

<p>Disruptor在优化并发性能方面可谓是做到了极致，优化的思路大体是两个方面，一个是利用无锁算法避免锁的争用，另外一个则是将硬件（CPU）的性能发挥到极致。尤其是后者，在Java领域基本上属于经典之作了。</p>

<p>发挥硬件的能力一般是C这种面向硬件的语言常干的事儿，C语言领域经常通过调整内存布局优化内存占用，而Java领域则用的很少，原因在于Java可以智能地优化内存布局，内存布局对Java程序员的透明的。这种智能的优化大部分场景是很友好的，但是如果你想通过填充方式避免伪共享就必须绕过这种优化，关于这方面Disruptor提供了经典的实现，你可以参考。</p>

<p>由于伪共享问题如此重要，所以Java也开始重视它了，比如Java 8中，提供了避免伪共享的注解：@sun.misc.Contended，通过这个注解就能轻松避免伪共享（需要设置JVM参数-XX:-RestrictContended）。不过避免伪共享是以牺牲内存为代价的，所以具体使用的时候还是需要仔细斟酌。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#caa6a6a6f3fefbfbfafd8aada7aba3a6e4a9a5a7" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f96948d577791',t:'MTczNDAyNjE1Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=42&#32;Actor模型：面向对象原生的并发模型>
        <link rel="icon" href="/static/favicon.png">
        <title>42 Actor模型：面向对象原生的并发模型 </title>
        
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
                            <h1 id="title" data-id="42 Actor模型：面向对象原生的并发模型" class="title">42 Actor模型：面向对象原生的并发模型</h1>
                            <div><p>上学的时候，有门计算机专业课叫做面向对象编程，学这门课的时候有个问题困扰了我很久，按照面向对象编程的理论，对象之间通信需要依靠<strong>消息</strong>，而实际上，像C++、Java这些面向对象的语言，对象之间通信，依靠的是<strong>对象方法</strong>。对象方法和过程语言里的函数本质上没有区别，有入参、有出参，思维方式很相似，使用起来都很简单。那面向对象理论里的消息是否就等价于面向对象语言里的对象方法呢？很长一段时间里，我都以为对象方法是面向对象理论中消息的一种实现，直到接触到Actor模型，才明白消息压根不是这个实现法。</p>

<h2 id="hello-actor模型">Hello Actor模型</h2>

<p>Actor模型本质上是一种计算模型，基本的计算单元称为Actor，换言之，<strong>在Actor模型中，所有的计算都是在Actor中执行的</strong>。在面向对象编程里面，一切都是对象；在Actor模型里，一切都是Actor，并且Actor之间是完全隔离的，不会共享任何变量。</p>

<p>当看到“不共享任何变量”的时候，相信你一定会眼前一亮，并发问题的根源就在于共享变量，而Actor模型中Actor之间不共享变量，那用Actor模型解决并发问题，一定是相当顺手。的确是这样，所以很多人就把Actor模型定义为一种<strong>并发计算模型</strong>。其实Actor模型早在1973年就被提出来了，只是直到最近几年才被广泛关注，一个主要原因就在于它是解决并发问题的利器，而最近几年随着多核处理器的发展，并发问题被推到了风口浪尖上。</p>

<p>但是Java语言本身并不支持Actor模型，所以如果你想在Java语言里使用Actor模型，就需要借助第三方类库，目前能完备地支持Actor模型而且比较成熟的类库就是<strong>Akka</strong>了。在详细介绍Actor模型之前，我们就先基于Akka写一个Hello World程序，让你对Actor模型先有个感官的印象。</p>

<p>在下面的示例代码中，我们首先创建了一个ActorSystem（Actor不能脱离ActorSystem存在）；之后创建了一个HelloActor，Akka中创建Actor并不是new一个对象出来，而是通过调用system.actorOf()方法创建的，该方法返回的是ActorRef，而不是HelloActor；最后通过调用ActorRef的tell()方法给HelloActor发送了一条消息 “Actor” 。</p>

<pre><code>//该Actor当收到消息message后，
//会打印Hello message
static class HelloActor 
    extends UntypedActor {
  @Override
  public void onReceive(Object message) {
    System.out.println(&quot;Hello &quot; + message);
  }
}

public static void main(String[] args) {
  //创建Actor系统
  ActorSystem system = ActorSystem.create(&quot;HelloSystem&quot;);
  //创建HelloActor
  ActorRef helloActor = 
    system.actorOf(Props.create(HelloActor.class));
  //发送消息给HelloActor
  helloActor.tell(&quot;Actor&quot;, ActorRef.noSender());
}
</code></pre>

<p>通过这个例子，你会发现Actor模型和面向对象编程契合度非常高，完全可以用Actor类比面向对象编程里面的对象，而且Actor之间的通信方式完美地遵守了消息机制，而不是通过对象方法来实现对象之间的通信。那Actor中的消息机制和面向对象语言里的对象方法有什么区别呢？</p>

<h2 id="消息和对象方法的区别">消息和对象方法的区别</h2>

<p>在没有计算机的时代，异地的朋友往往是通过写信来交流感情的，但信件发出去之后，也许会在寄送过程中弄丢了，也有可能寄到后，对方一直没有时间写回信……这个时候都可以让邮局“背个锅”，不过无论如何，也不过是重写一封，生活继续。</p>

<p>Actor中的消息机制，就可以类比这现实世界里的写信。Actor内部有一个邮箱（Mailbox），接收到的消息都是先放到邮箱里，如果邮箱里有积压的消息，那么新收到的消息就不会马上得到处理，也正是因为Actor使用单线程处理消息，所以不会出现并发问题。你可以把Actor内部的工作模式想象成只有一个消费者线程的生产者-消费者模式。</p>

<p>所以，在Actor模型里，发送消息仅仅是把消息发出去而已，接收消息的Actor在接收到消息后，也不一定会立即处理，也就是说<strong>Actor中的消息机制完全是异步的</strong>。而<strong>调用对象方法</strong>，实际上是<strong>同步</strong>的，对象方法return之前，调用方会一直等待。</p>

<p>除此之外，<strong>调用对象方法</strong>，需要持有对象的引用，<strong>所有的对象必须在同一个进程中</strong>。而在Actor中发送消息，类似于现实中的写信，只需要知道对方的地址就可以，<strong>发送消息和接收消息的Actor可以不在一个进程中，也可以不在同一台机器上</strong>。因此，Actor模型不但适用于并发计算，还适用于分布式计算。</p>

<h2 id="actor的规范化定义">Actor的规范化定义</h2>

<p>通过上面的介绍，相信你应该已经对Actor有一个感官印象了，下面我们再来看看Actor规范化的定义是什么样的。Actor是一种基础的计算单元，具体来讲包括三部分能力，分别是：</p>

<ol>
<li>处理能力，处理接收到的消息。</li>
<li>存储能力，Actor可以存储自己的内部状态，并且内部状态在不同Actor之间是绝对隔离的。</li>
<li>通信能力，Actor可以和其他Actor之间通信。</li>
</ol>

<p>当一个Actor接收的一条消息之后，这个Actor可以做以下三件事：</p>

<ol>
<li>创建更多的Actor；</li>
<li>发消息给其他Actor；</li>
<li>确定如何处理下一条消息。</li>
</ol>

<p>其中前两条还是很好理解的，就是最后一条，该如何去理解呢？前面我们说过Actor具备存储能力，它有自己的内部状态，所以你也可以把Actor看作一个状态机，把Actor处理消息看作是触发状态机的状态变化；而状态机的变化往往要基于上一个状态，触发状态机发生变化的时刻，上一个状态必须是确定的，所以确定如何处理下一条消息，本质上不过是改变内部状态。</p>

<p>在多线程里面，由于可能存在竞态条件，所以根据当前状态确定如何处理下一条消息还是有难度的，需要使用各种同步工具，但在Actor模型里，由于是单线程处理，所以就不存在竞态条件问题了。</p>

<h2 id="用actor实现累加器">用Actor实现累加器</h2>

<p>支持并发的累加器可能是最简单并且有代表性的并发问题了，可以基于互斥锁方案实现，也可以基于原子类实现，但今天我们要尝试用Actor来实现。</p>

<p>在下面的示例代码中，CounterActor内部持有累计值counter，当CounterActor接收到一个数值型的消息message时，就将累计值counter += message；但如果是其他类型的消息，则打印当前累计值counter。在main()方法中，我们启动了4个线程来执行累加操作。整个程序没有锁，也没有CAS，但是程序是线程安全的。</p>

<pre><code>//累加器
static class CounterActor extends UntypedActor {
  private int counter = 0;
  @Override
  public void onReceive(Object message){
    //如果接收到的消息是数字类型，执行累加操作，
    //否则打印counter的值
    if (message instanceof Number) {
      counter += ((Number) message).intValue();
    } else {
      System.out.println(counter);
    }
  }
}
public static void main(String[] args) throws InterruptedException {
  //创建Actor系统
  ActorSystem system = ActorSystem.create(&quot;HelloSystem&quot;);
  //4个线程生产消息
  ExecutorService es = Executors.newFixedThreadPool(4);
  //创建CounterActor 
  ActorRef counterActor = 
    system.actorOf(Props.create(CounterActor.class));
  //生产4*100000个消息 
  for (int i=0; i&lt;4; i++) {
    es.execute(()-&gt;{
      for (int j=0; j&lt;100000; j++) {
        counterActor.tell(1, ActorRef.noSender());
      }
    });
  }
  //关闭线程池
  es.shutdown();
  //等待CounterActor处理完所有消息
  Thread.sleep(1000);
  //打印结果
  counterActor.tell(&quot;&quot;, ActorRef.noSender());
  //关闭Actor系统
  system.shutdown();
}
</code></pre>

<h2 id="总结">总结</h2>

<p>Actor模型是一种非常简单的计算模型，其中Actor是最基本的计算单元，Actor之间是通过消息进行通信。Actor与面向对象编程（OOP）中的对象匹配度非常高，在面向对象编程里，系统由类似于生物细胞那样的对象构成，对象之间也是通过消息进行通信，所以在面向对象语言里使用Actor模型基本上不会有违和感。</p>

<p>在Java领域，除了可以使用Akka来支持Actor模型外，还可以使用Vert.x，不过相对来说Vert.x更像是Actor模型的隐式实现，对应关系不像Akka那样明显，不过本质上也是一种Actor模型。</p>

<p>Actor可以创建新的Actor，这些Actor最终会呈现出一个树状结构，非常像现实世界里的组织结构，所以利用Actor模型来对程序进行建模，和现实世界的匹配度非常高。Actor模型和现实世界一样都是异步模型，理论上不保证消息百分百送达，也不保证消息送达的顺序和发送的顺序是一致的，甚至无法保证消息会被百分百处理。虽然实现Actor模型的厂商都在试图解决这些问题，但遗憾的是解决得并不完美，所以使用Actor模型也是有成本的。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#610d0d0d58555050515621060c00080d4f020e0c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f97049bb37791',t:'MTczNDAyNjE3NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
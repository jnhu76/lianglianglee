<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=39&#32;原子类是如何利用&#32;CAS&#32;保证线程安全的？>
        <link rel="icon" href="/static/favicon.png">
        <title>39 原子类是如何利用 CAS 保证线程安全的？ </title>
        
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
                        <a class="menu-item" id="00 由点及面，搭建你的 Java 并发知识网.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/00%20%e7%94%b1%e7%82%b9%e5%8f%8a%e9%9d%a2%ef%bc%8c%e6%90%ad%e5%bb%ba%e4%bd%a0%e7%9a%84%20Java%20%e5%b9%b6%e5%8f%91%e7%9f%a5%e8%af%86%e7%bd%91.md">00 由点及面，搭建你的 Java 并发知识网.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 为何说只有 1 种实现线程的方法？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/01%20%e4%b8%ba%e4%bd%95%e8%af%b4%e5%8f%aa%e6%9c%89%201%20%e7%a7%8d%e5%ae%9e%e7%8e%b0%e7%ba%bf%e7%a8%8b%e7%9a%84%e6%96%b9%e6%b3%95%ef%bc%9f.md">01 为何说只有 1 种实现线程的方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 如何正确停止线程？为什么 volatile 标记位的停止方法是错误的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/02%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e5%81%9c%e6%ad%a2%e7%ba%bf%e7%a8%8b%ef%bc%9f%e4%b8%ba%e4%bb%80%e4%b9%88%20volatile%20%e6%a0%87%e8%ae%b0%e4%bd%8d%e7%9a%84%e5%81%9c%e6%ad%a2%e6%96%b9%e6%b3%95%e6%98%af%e9%94%99%e8%af%af%e7%9a%84%ef%bc%9f.md">02 如何正确停止线程？为什么 volatile 标记位的停止方法是错误的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 线程是如何在 6 种状态之间转换的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/03%20%e7%ba%bf%e7%a8%8b%e6%98%af%e5%a6%82%e4%bd%95%e5%9c%a8%206%20%e7%a7%8d%e7%8a%b6%e6%80%81%e4%b9%8b%e9%97%b4%e8%bd%ac%e6%8d%a2%e7%9a%84%ef%bc%9f.md">03 线程是如何在 6 种状态之间转换的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 waitnotifynotifyAll 方法的使用注意事项？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/04%20waitnotifynotifyAll%20%e6%96%b9%e6%b3%95%e7%9a%84%e4%bd%bf%e7%94%a8%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9%ef%bc%9f.md">04 waitnotifynotifyAll 方法的使用注意事项？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 有哪几种实现生产者消费者模式的方法？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/05%20%e6%9c%89%e5%93%aa%e5%87%a0%e7%a7%8d%e5%ae%9e%e7%8e%b0%e7%94%9f%e4%ba%a7%e8%80%85%e6%b6%88%e8%b4%b9%e8%80%85%e6%a8%a1%e5%bc%8f%e7%9a%84%e6%96%b9%e6%b3%95%ef%bc%9f.md">05 有哪几种实现生产者消费者模式的方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 一共有哪 3 类线程安全问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/06%20%e4%b8%80%e5%85%b1%e6%9c%89%e5%93%aa%203%20%e7%b1%bb%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98%ef%bc%9f.md">06 一共有哪 3 类线程安全问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 哪些场景需要额外注意线程安全问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/07%20%e5%93%aa%e4%ba%9b%e5%9c%ba%e6%99%af%e9%9c%80%e8%a6%81%e9%a2%9d%e5%a4%96%e6%b3%a8%e6%84%8f%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98%ef%bc%9f.md">07 哪些场景需要额外注意线程安全问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 为什么多线程会带来性能问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/08%20%e4%b8%ba%e4%bb%80%e4%b9%88%e5%a4%9a%e7%ba%bf%e7%a8%8b%e4%bc%9a%e5%b8%a6%e6%9d%a5%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%ef%bc%9f.md">08 为什么多线程会带来性能问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 使用线程池比手动创建线程好在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/09%20%e4%bd%bf%e7%94%a8%e7%ba%bf%e7%a8%8b%e6%b1%a0%e6%af%94%e6%89%8b%e5%8a%a8%e5%88%9b%e5%bb%ba%e7%ba%bf%e7%a8%8b%e5%a5%bd%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">09 使用线程池比手动创建线程好在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 线程池的各个参数的含义？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/10%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%e7%9a%84%e5%90%84%e4%b8%aa%e5%8f%82%e6%95%b0%e7%9a%84%e5%90%ab%e4%b9%89%ef%bc%9f.md">10 线程池的各个参数的含义？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 线程池有哪 4 种拒绝策略？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/11%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%e6%9c%89%e5%93%aa%204%20%e7%a7%8d%e6%8b%92%e7%bb%9d%e7%ad%96%e7%95%a5%ef%bc%9f.md">11 线程池有哪 4 种拒绝策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 有哪 6 种常见的线程池？什么是 Java8 的 ForkJoinPool？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/12%20%e6%9c%89%e5%93%aa%206%20%e7%a7%8d%e5%b8%b8%e8%a7%81%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f%e4%bb%80%e4%b9%88%e6%98%af%20Java8%20%e7%9a%84%20ForkJoinPool%ef%bc%9f.md">12 有哪 6 种常见的线程池？什么是 Java8 的 ForkJoinPool？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 线程池常用的阻塞队列有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/13%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%e5%b8%b8%e7%94%a8%e7%9a%84%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">13 线程池常用的阻塞队列有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 为什么不应该自动创建线程池？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/14%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%8d%e5%ba%94%e8%af%a5%e8%87%aa%e5%8a%a8%e5%88%9b%e5%bb%ba%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f.md">14 为什么不应该自动创建线程池？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 合适的线程数量是多少？CPU 核心数和线程数的关系？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/15%20%e5%90%88%e9%80%82%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%95%b0%e9%87%8f%e6%98%af%e5%a4%9a%e5%b0%91%ef%bc%9fCPU%20%e6%a0%b8%e5%bf%83%e6%95%b0%e5%92%8c%e7%ba%bf%e7%a8%8b%e6%95%b0%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">15 合适的线程数量是多少？CPU 核心数和线程数的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 如何根据实际需要，定制自己的线程池？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/16%20%e5%a6%82%e4%bd%95%e6%a0%b9%e6%8d%ae%e5%ae%9e%e9%99%85%e9%9c%80%e8%a6%81%ef%bc%8c%e5%ae%9a%e5%88%b6%e8%87%aa%e5%b7%b1%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9f.md">16 如何根据实际需要，定制自己的线程池？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 如何正确关闭线程池？shutdown 和 shutdownNow 的区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/17%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e5%85%b3%e9%97%ad%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9fshutdown%20%e5%92%8c%20shutdownNow%20%e7%9a%84%e5%8c%ba%e5%88%ab%ef%bc%9f.md">17 如何正确关闭线程池？shutdown 和 shutdownNow 的区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 线程池实现“线程复用”的原理？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/18%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%e5%ae%9e%e7%8e%b0%e2%80%9c%e7%ba%bf%e7%a8%8b%e5%a4%8d%e7%94%a8%e2%80%9d%e7%9a%84%e5%8e%9f%e7%90%86%ef%bc%9f.md">18 线程池实现“线程复用”的原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 你知道哪几种锁？分别有什么特点？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/19%20%e4%bd%a0%e7%9f%a5%e9%81%93%e5%93%aa%e5%87%a0%e7%a7%8d%e9%94%81%ef%bc%9f%e5%88%86%e5%88%ab%e6%9c%89%e4%bb%80%e4%b9%88%e7%89%b9%e7%82%b9%ef%bc%9f.md">19 你知道哪几种锁？分别有什么特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 悲观锁和乐观锁的本质是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/20%20%e6%82%b2%e8%a7%82%e9%94%81%e5%92%8c%e4%b9%90%e8%a7%82%e9%94%81%e7%9a%84%e6%9c%ac%e8%b4%a8%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">20 悲观锁和乐观锁的本质是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 如何看到 synchronized 背后的“monitor 锁”？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/21%20%e5%a6%82%e4%bd%95%e7%9c%8b%e5%88%b0%20synchronized%20%e8%83%8c%e5%90%8e%e7%9a%84%e2%80%9cmonitor%20%e9%94%81%e2%80%9d%ef%bc%9f.md">21 如何看到 synchronized 背后的“monitor 锁”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 synchronized 和 Lock 孰优孰劣，如何选择？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/22%20synchronized%20%e5%92%8c%20Lock%20%e5%ad%b0%e4%bc%98%e5%ad%b0%e5%8a%a3%ef%bc%8c%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%ef%bc%9f.md">22 synchronized 和 Lock 孰优孰劣，如何选择？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Lock 有哪几个常用方法？分别有什么用？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/23%20Lock%20%e6%9c%89%e5%93%aa%e5%87%a0%e4%b8%aa%e5%b8%b8%e7%94%a8%e6%96%b9%e6%b3%95%ef%bc%9f%e5%88%86%e5%88%ab%e6%9c%89%e4%bb%80%e4%b9%88%e7%94%a8%ef%bc%9f.md">23 Lock 有哪几个常用方法？分别有什么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 讲一讲公平锁和非公平锁，为什么要“非公平”？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/24%20%e8%ae%b2%e4%b8%80%e8%ae%b2%e5%85%ac%e5%b9%b3%e9%94%81%e5%92%8c%e9%9d%9e%e5%85%ac%e5%b9%b3%e9%94%81%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e2%80%9c%e9%9d%9e%e5%85%ac%e5%b9%b3%e2%80%9d%ef%bc%9f.md">24 讲一讲公平锁和非公平锁，为什么要“非公平”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 读写锁 ReadWriteLock 获取锁有哪些规则？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/25%20%e8%af%bb%e5%86%99%e9%94%81%20ReadWriteLock%20%e8%8e%b7%e5%8f%96%e9%94%81%e6%9c%89%e5%93%aa%e4%ba%9b%e8%a7%84%e5%88%99%ef%bc%9f.md">25 读写锁 ReadWriteLock 获取锁有哪些规则？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 读锁应该插队吗？什么是读写锁的升降级？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/26%20%e8%af%bb%e9%94%81%e5%ba%94%e8%af%a5%e6%8f%92%e9%98%9f%e5%90%97%ef%bc%9f%e4%bb%80%e4%b9%88%e6%98%af%e8%af%bb%e5%86%99%e9%94%81%e7%9a%84%e5%8d%87%e9%99%8d%e7%ba%a7%ef%bc%9f.md">26 读锁应该插队吗？什么是读写锁的升降级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 什么是自旋锁？自旋的好处和后果是什么呢？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/27%20%e4%bb%80%e4%b9%88%e6%98%af%e8%87%aa%e6%97%8b%e9%94%81%ef%bc%9f%e8%87%aa%e6%97%8b%e7%9a%84%e5%a5%bd%e5%a4%84%e5%92%8c%e5%90%8e%e6%9e%9c%e6%98%af%e4%bb%80%e4%b9%88%e5%91%a2%ef%bc%9f.md">27 什么是自旋锁？自旋的好处和后果是什么呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 JVM 对锁进行了哪些优化？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/28%20JVM%20%e5%af%b9%e9%94%81%e8%bf%9b%e8%a1%8c%e4%ba%86%e5%93%aa%e4%ba%9b%e4%bc%98%e5%8c%96%ef%bc%9f.md">28 JVM 对锁进行了哪些优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 HashMap 为什么是线程不安全的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/29%20HashMap%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%98%af%e7%ba%bf%e7%a8%8b%e4%b8%8d%e5%ae%89%e5%85%a8%e7%9a%84%ef%bc%9f.md">29 HashMap 为什么是线程不安全的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 ConcurrentHashMap 在 Java7 和 8 有何不同？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/30%20ConcurrentHashMap%20%e5%9c%a8%20Java7%20%e5%92%8c%208%20%e6%9c%89%e4%bd%95%e4%b8%8d%e5%90%8c%ef%bc%9f.md">30 ConcurrentHashMap 在 Java7 和 8 有何不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 为什么 Map 桶中超过 8 个才转为红黑树？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/31%20%e4%b8%ba%e4%bb%80%e4%b9%88%20Map%20%e6%a1%b6%e4%b8%ad%e8%b6%85%e8%bf%87%208%20%e4%b8%aa%e6%89%8d%e8%bd%ac%e4%b8%ba%e7%ba%a2%e9%bb%91%e6%a0%91%ef%bc%9f.md">31 为什么 Map 桶中超过 8 个才转为红黑树？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 同样是线程安全，ConcurrentHashMap 和 Hashtable 的区别.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/32%20%e5%90%8c%e6%a0%b7%e6%98%af%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%ef%bc%8cConcurrentHashMap%20%e5%92%8c%20Hashtable%20%e7%9a%84%e5%8c%ba%e5%88%ab.md">32 同样是线程安全，ConcurrentHashMap 和 Hashtable 的区别.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 CopyOnWriteArrayList 有什么特点？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/33%20CopyOnWriteArrayList%20%e6%9c%89%e4%bb%80%e4%b9%88%e7%89%b9%e7%82%b9%ef%bc%9f.md">33 CopyOnWriteArrayList 有什么特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 什么是阻塞队列？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/34%20%e4%bb%80%e4%b9%88%e6%98%af%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%ef%bc%9f.md">34 什么是阻塞队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 阻塞队列包含哪些常用的方法？add、offer、put 等方法的区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/35%20%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%e5%8c%85%e5%90%ab%e5%93%aa%e4%ba%9b%e5%b8%b8%e7%94%a8%e7%9a%84%e6%96%b9%e6%b3%95%ef%bc%9fadd%e3%80%81offer%e3%80%81put%20%e7%ad%89%e6%96%b9%e6%b3%95%e7%9a%84%e5%8c%ba%e5%88%ab%ef%bc%9f.md">35 阻塞队列包含哪些常用的方法？add、offer、put 等方法的区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 有哪几种常见的阻塞队列？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/36%20%e6%9c%89%e5%93%aa%e5%87%a0%e7%a7%8d%e5%b8%b8%e8%a7%81%e7%9a%84%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%ef%bc%9f.md">36 有哪几种常见的阻塞队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 阻塞和非阻塞队列的并发安全原理是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/37%20%e9%98%bb%e5%a1%9e%e5%92%8c%e9%9d%9e%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%e7%9a%84%e5%b9%b6%e5%8f%91%e5%ae%89%e5%85%a8%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">37 阻塞和非阻塞队列的并发安全原理是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 如何选择适合自己的阻塞队列？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/38%20%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e9%80%82%e5%90%88%e8%87%aa%e5%b7%b1%e7%9a%84%e9%98%bb%e5%a1%9e%e9%98%9f%e5%88%97%ef%bc%9f.md">38 如何选择适合自己的阻塞队列？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 原子类是如何利用 CAS 保证线程安全的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/39%20%e5%8e%9f%e5%ad%90%e7%b1%bb%e6%98%af%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20CAS%20%e4%bf%9d%e8%af%81%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e7%9a%84%ef%bc%9f.md">39 原子类是如何利用 CAS 保证线程安全的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 AtomicInteger 在高并发下性能不好，如何解决？为什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/40%20AtomicInteger%20%e5%9c%a8%e9%ab%98%e5%b9%b6%e5%8f%91%e4%b8%8b%e6%80%a7%e8%83%bd%e4%b8%8d%e5%a5%bd%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%ef%bc%9f%e4%b8%ba%e4%bb%80%e4%b9%88%ef%bc%9f.md">40 AtomicInteger 在高并发下性能不好，如何解决？为什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 原子类和 volatile 有什么异同？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/41%20%e5%8e%9f%e5%ad%90%e7%b1%bb%e5%92%8c%20volatile%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%bc%82%e5%90%8c%ef%bc%9f.md">41 原子类和 volatile 有什么异同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 AtomicInteger 和 synchronized 的异同点？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/42%20AtomicInteger%20%e5%92%8c%20synchronized%20%e7%9a%84%e5%bc%82%e5%90%8c%e7%82%b9%ef%bc%9f.md">42 AtomicInteger 和 synchronized 的异同点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 Java 8 中 Adder 和 Accumulator 有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/43%20Java%208%20%e4%b8%ad%20Adder%20%e5%92%8c%20Accumulator%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">43 Java 8 中 Adder 和 Accumulator 有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 ThreadLocal 适合用在哪些实际生产的场景中？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/44%20ThreadLocal%20%e9%80%82%e5%90%88%e7%94%a8%e5%9c%a8%e5%93%aa%e4%ba%9b%e5%ae%9e%e9%99%85%e7%94%9f%e4%ba%a7%e7%9a%84%e5%9c%ba%e6%99%af%e4%b8%ad%ef%bc%9f.md">44 ThreadLocal 适合用在哪些实际生产的场景中？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 ThreadLocal 是用来解决共享资源的多线程访问的问题吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/45%20ThreadLocal%20%e6%98%af%e7%94%a8%e6%9d%a5%e8%a7%a3%e5%86%b3%e5%85%b1%e4%ba%ab%e8%b5%84%e6%ba%90%e7%9a%84%e5%a4%9a%e7%ba%bf%e7%a8%8b%e8%ae%bf%e9%97%ae%e7%9a%84%e9%97%ae%e9%a2%98%e5%90%97%ef%bc%9f.md">45 ThreadLocal 是用来解决共享资源的多线程访问的问题吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 多个 ThreadLocal 在 Thread 中的 threadlocals 里是怎么存储的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/46%20%e5%a4%9a%e4%b8%aa%20ThreadLocal%20%e5%9c%a8%20Thread%20%e4%b8%ad%e7%9a%84%20threadlocals%20%e9%87%8c%e6%98%af%e6%80%8e%e4%b9%88%e5%ad%98%e5%82%a8%e7%9a%84%ef%bc%9f.md">46 多个 ThreadLocal 在 Thread 中的 threadlocals 里是怎么存储的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 内存泄漏——为何每次用完 ThreadLocal 都要调用 remove()？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/47%20%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%e2%80%94%e2%80%94%e4%b8%ba%e4%bd%95%e6%af%8f%e6%ac%a1%e7%94%a8%e5%ae%8c%20ThreadLocal%20%e9%83%bd%e8%a6%81%e8%b0%83%e7%94%a8%20remove%28%29%ef%bc%9f.md">47 内存泄漏——为何每次用完 ThreadLocal 都要调用 remove()？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 Callable 和 Runnable 的不同？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/48%20Callable%20%e5%92%8c%20Runnable%20%e7%9a%84%e4%b8%8d%e5%90%8c%ef%bc%9f.md">48 Callable 和 Runnable 的不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 Future 的主要功能是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/49%20Future%20%e7%9a%84%e4%b8%bb%e8%a6%81%e5%8a%9f%e8%83%bd%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">49 Future 的主要功能是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 使用 Future 有哪些注意点？Future 产生新的线程了吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/50%20%e4%bd%bf%e7%94%a8%20Future%20%e6%9c%89%e5%93%aa%e4%ba%9b%e6%b3%a8%e6%84%8f%e7%82%b9%ef%bc%9fFuture%20%e4%ba%a7%e7%94%9f%e6%96%b0%e7%9a%84%e7%ba%bf%e7%a8%8b%e4%ba%86%e5%90%97%ef%bc%9f.md">50 使用 Future 有哪些注意点？Future 产生新的线程了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 如何利用 CompletableFuture 实现“旅游平台”问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/51%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20CompletableFuture%20%e5%ae%9e%e7%8e%b0%e2%80%9c%e6%97%85%e6%b8%b8%e5%b9%b3%e5%8f%b0%e2%80%9d%e9%97%ae%e9%a2%98%ef%bc%9f.md">51 如何利用 CompletableFuture 实现“旅游平台”问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52 信号量能被 FixedThreadPool 替代吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/52%20%e4%bf%a1%e5%8f%b7%e9%87%8f%e8%83%bd%e8%a2%ab%20FixedThreadPool%20%e6%9b%bf%e4%bb%a3%e5%90%97%ef%bc%9f.md">52 信号量能被 FixedThreadPool 替代吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="53 CountDownLatch 是如何安排线程执行顺序的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/53%20CountDownLatch%20%e6%98%af%e5%a6%82%e4%bd%95%e5%ae%89%e6%8e%92%e7%ba%bf%e7%a8%8b%e6%89%a7%e8%a1%8c%e9%a1%ba%e5%ba%8f%e7%9a%84%ef%bc%9f.md">53 CountDownLatch 是如何安排线程执行顺序的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="54 CyclicBarrier 和 CountdownLatch 有什么异同？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/54%20CyclicBarrier%20%e5%92%8c%20CountdownLatch%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%bc%82%e5%90%8c%ef%bc%9f.md">54 CyclicBarrier 和 CountdownLatch 有什么异同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="55 Condition、object.wait() 和 notify() 的关系？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/55%20Condition%e3%80%81object.wait%28%29%20%e5%92%8c%20notify%28%29%20%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">55 Condition、object.wait() 和 notify() 的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="56 讲一讲什么是 Java 内存模型？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/56%20%e8%ae%b2%e4%b8%80%e8%ae%b2%e4%bb%80%e4%b9%88%e6%98%af%20Java%20%e5%86%85%e5%ad%98%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">56 讲一讲什么是 Java 内存模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="57 什么是指令重排序？为什么要重排序？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/57%20%e4%bb%80%e4%b9%88%e6%98%af%e6%8c%87%e4%bb%a4%e9%87%8d%e6%8e%92%e5%ba%8f%ef%bc%9f%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e9%87%8d%e6%8e%92%e5%ba%8f%ef%bc%9f.md">57 什么是指令重排序？为什么要重排序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="58 Java 中的原子操作有哪些注意事项？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/58%20Java%20%e4%b8%ad%e7%9a%84%e5%8e%9f%e5%ad%90%e6%93%8d%e4%bd%9c%e6%9c%89%e5%93%aa%e4%ba%9b%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9%ef%bc%9f.md">58 Java 中的原子操作有哪些注意事项？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="59 什么是“内存可见性”问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/59%20%e4%bb%80%e4%b9%88%e6%98%af%e2%80%9c%e5%86%85%e5%ad%98%e5%8f%af%e8%a7%81%e6%80%a7%e2%80%9d%e9%97%ae%e9%a2%98%ef%bc%9f.md">59 什么是“内存可见性”问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="60 主内存和工作内存的关系？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/60%20%e4%b8%bb%e5%86%85%e5%ad%98%e5%92%8c%e5%b7%a5%e4%bd%9c%e5%86%85%e5%ad%98%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">60 主内存和工作内存的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="61 什么是 happens-before 规则？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/61%20%e4%bb%80%e4%b9%88%e6%98%af%20happens-before%20%e8%a7%84%e5%88%99%ef%bc%9f.md">61 什么是 happens-before 规则？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="62 volatile 的作用是什么？与 synchronized 有什么异同？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/62%20volatile%20%e7%9a%84%e4%bd%9c%e7%94%a8%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e4%b8%8e%20synchronized%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%bc%82%e5%90%8c%ef%bc%9f.md">62 volatile 的作用是什么？与 synchronized 有什么异同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="63 单例模式的双重检查锁模式为什么必须加 volatile？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/63%20%e5%8d%95%e4%be%8b%e6%a8%a1%e5%bc%8f%e7%9a%84%e5%8f%8c%e9%87%8d%e6%a3%80%e6%9f%a5%e9%94%81%e6%a8%a1%e5%bc%8f%e4%b8%ba%e4%bb%80%e4%b9%88%e5%bf%85%e9%a1%bb%e5%8a%a0%20volatile%ef%bc%9f.md">63 单例模式的双重检查锁模式为什么必须加 volatile？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="64 你知道什么是 CAS 吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/64%20%e4%bd%a0%e7%9f%a5%e9%81%93%e4%bb%80%e4%b9%88%e6%98%af%20CAS%20%e5%90%97%ef%bc%9f.md">64 你知道什么是 CAS 吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="65 CAS 和乐观锁的关系，什么时候会用到 CAS？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/65%20CAS%20%e5%92%8c%e4%b9%90%e8%a7%82%e9%94%81%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%8c%e4%bb%80%e4%b9%88%e6%97%b6%e5%80%99%e4%bc%9a%e7%94%a8%e5%88%b0%20CAS%ef%bc%9f.md">65 CAS 和乐观锁的关系，什么时候会用到 CAS？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="66 CAS 有什么缺点？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/66%20CAS%20%e6%9c%89%e4%bb%80%e4%b9%88%e7%bc%ba%e7%82%b9%ef%bc%9f.md">66 CAS 有什么缺点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="67 如何写一个必然死锁的例子？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/67%20%e5%a6%82%e4%bd%95%e5%86%99%e4%b8%80%e4%b8%aa%e5%bf%85%e7%84%b6%e6%ad%bb%e9%94%81%e7%9a%84%e4%be%8b%e5%ad%90%ef%bc%9f.md">67 如何写一个必然死锁的例子？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="68 发生死锁必须满足哪 4 个条件？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/68%20%e5%8f%91%e7%94%9f%e6%ad%bb%e9%94%81%e5%bf%85%e9%a1%bb%e6%bb%a1%e8%b6%b3%e5%93%aa%204%20%e4%b8%aa%e6%9d%a1%e4%bb%b6%ef%bc%9f.md">68 发生死锁必须满足哪 4 个条件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="69 如何用命令行和代码定位死锁？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/69%20%e5%a6%82%e4%bd%95%e7%94%a8%e5%91%bd%e4%bb%a4%e8%a1%8c%e5%92%8c%e4%bb%a3%e7%a0%81%e5%ae%9a%e4%bd%8d%e6%ad%bb%e9%94%81%ef%bc%9f.md">69 如何用命令行和代码定位死锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="70 有哪些解决死锁问题的策略？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/70%20%e6%9c%89%e5%93%aa%e4%ba%9b%e8%a7%a3%e5%86%b3%e6%ad%bb%e9%94%81%e9%97%ae%e9%a2%98%e7%9a%84%e7%ad%96%e7%95%a5%ef%bc%9f.md">70 有哪些解决死锁问题的策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="71 讲一讲经典的哲学家就餐问题.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/71%20%e8%ae%b2%e4%b8%80%e8%ae%b2%e7%bb%8f%e5%85%b8%e7%9a%84%e5%93%b2%e5%ad%a6%e5%ae%b6%e5%b0%b1%e9%a4%90%e9%97%ae%e9%a2%98.md">71 讲一讲经典的哲学家就餐问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="72 final 的三种用法是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/72%20final%20%e7%9a%84%e4%b8%89%e7%a7%8d%e7%94%a8%e6%b3%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">72 final 的三种用法是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="73 为什么加了 final 却依然无法拥有“不变性”？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/73%20%e4%b8%ba%e4%bb%80%e4%b9%88%e5%8a%a0%e4%ba%86%20final%20%e5%8d%b4%e4%be%9d%e7%84%b6%e6%97%a0%e6%b3%95%e6%8b%a5%e6%9c%89%e2%80%9c%e4%b8%8d%e5%8f%98%e6%80%a7%e2%80%9d%ef%bc%9f.md">73 为什么加了 final 却依然无法拥有“不变性”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="74 为什么 String 被设计为是不可变的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/74%20%e4%b8%ba%e4%bb%80%e4%b9%88%20String%20%e8%a2%ab%e8%ae%be%e8%ae%a1%e4%b8%ba%e6%98%af%e4%b8%8d%e5%8f%af%e5%8f%98%e7%9a%84%ef%bc%9f.md">74 为什么 String 被设计为是不可变的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="75 为什么需要 AQS？AQS 的作用和重要性是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/75%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%20AQS%ef%bc%9fAQS%20%e7%9a%84%e4%bd%9c%e7%94%a8%e5%92%8c%e9%87%8d%e8%a6%81%e6%80%a7%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">75 为什么需要 AQS？AQS 的作用和重要性是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="76 AQS 的内部原理是什么样的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/76%20AQS%20%e7%9a%84%e5%86%85%e9%83%a8%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%ef%bc%9f.md">76 AQS 的内部原理是什么样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="77 AQS 在 CountDownLatch 等类中的应用原理是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/77%20AQS%20%e5%9c%a8%20CountDownLatch%20%e7%ad%89%e7%b1%bb%e4%b8%ad%e7%9a%84%e5%ba%94%e7%94%a8%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">77 AQS 在 CountDownLatch 等类中的应用原理是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="78 一份独家的 Java 并发工具图谱.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%2078%20%e8%ae%b2-%e5%ae%8c/78%20%e4%b8%80%e4%bb%bd%e7%8b%ac%e5%ae%b6%e7%9a%84%20Java%20%e5%b9%b6%e5%8f%91%e5%b7%a5%e5%85%b7%e5%9b%be%e8%b0%b1.md">78 一份独家的 Java 并发工具图谱.md</a>
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
                            <h1 id="title" data-id="39 原子类是如何利用 CAS 保证线程安全的？" class="title">39 原子类是如何利用 CAS 保证线程安全的？</h1>
                            <div><p>本课时主要讲解原子类是如何利用 CAS 保证线程安全的。</p>

<h3 id="什么是原子类-原子类有什么作用">什么是原子类？原子类有什么作用？</h3>

<p>要想回答这个问题，首先我们需要知道什么是原子类，以及它有什么作用。</p>

<p>在编程领域里，原子性意味着“一组操作要么全都操作成功，要么全都失败，不能只操作成功其中的一部分”。而 java.util.concurrent.atomic 下的类，就是具有原子性的类，可以原子性地执行添加、递增、递减等操作。比如之前多线程下的线程不安全的 i++ 问题，到了原子类这里，就可以用功能相同且线程安全的 getAndIncrement 方法来优雅地解决。</p>

<p><strong>原子类的作用</strong>和锁有类似之处，是为了保证并发情况下线程安全。不过原子类相比于锁，有一定的优势：</p>

<ul>
<li>粒度更细：原子变量可以把竞争范围缩小到变量级别，通常情况下，锁的粒度都要大于原子变量的粒度。</li>
<li>效率更高：除了高度竞争的情况之外，使用原子类的效率通常会比使用同步互斥锁的效率更高，因为原子类底层利用了 CAS 操作，不会阻塞线程。</li>
</ul>

<h3 id="6-类原子类纵览">6 类原子类纵览</h3>

<p>下面我们来看下一共有哪些原子类，原子类一共可以分为以下这 6 类，我们来逐一介绍：</p>

<table>
<thead>
<tr>
<th>类型</th>
<th>具体类</th>
</tr>
</thead>

<tbody>
<tr>
<td>Atomic* 基本类型原子类</td>
<td>AtomicInteger、AtomicLong、AtomicBoolean</td>
</tr>

<tr>
<td>Atomic*Array 数组类型原子类</td>
<td>AtomicIntegerArray、AtomicLongArray、AtomicReferenceArray</td>
</tr>

<tr>
<td>Atomic*Reference 引用类型原子类</td>
<td>AtomicReference、AtomicStampedReference、AtomicMarkableReference</td>
</tr>

<tr>
<td>Atomic*FieldUpdater 升级类型原子类</td>
<td>AtomicIntegerfieldupdater、AtomicLongFieldUpdater、AtomicReferenceFieldUpdater</td>
</tr>

<tr>
<td>Adder 累加器</td>
<td>LongAdder、DoubleAdder</td>
</tr>

<tr>
<td>Accumulator 积累器</td>
<td>LongAccumulator、DoubleAccumulator</td>
</tr>
</tbody>
</table>

<h3 id="atomic-基本类型原子类">Atomic\ 基本类型原子类</h3>

<p>首先看到第一类 Atomic*，我们把它称为基本类型原子类，它包括三种，分别是 AtomicInteger、AtomicLong 和 AtomicBoolean。</p>

<p>我们来介绍一下最为典型的 AtomicInteger。对于这个类型而言，它是对于 int 类型的封装，并且提供了原子性的访问和更新。也就是说，我们如果需要一个整型的变量，并且这个变量会被运用在并发场景之下，我们可以不用基本类型 int，也不使用包装类型 Integer，而是直接使用 AtomicInteger，这样一来就自动具备了原子能力，使用起来非常方便。</p>

<h4 id="atomicinteger-类常用方法">AtomicInteger 类常用方法</h4>

<p>AtomicInteger 类有以下几个常用的方法：</p>

<ul>
<li>public final int get() //获取当前的值</li>
</ul>

<p>因为它本身是一个 Java 类，而不再是一个基本类型，所以要想获取值还是需要一些方法，比如通过 get 方法就可以获取到当前的值。</p>

<ul>
<li>public final int getAndSet(int newValue) //获取当前的值，并设置新的值</li>
</ul>

<p>接下来的几个方法和它平时的操作相关：</p>

<ul>
<li>public final int getAndIncrement() //获取当前的值，并自增</li>
<li>public final int getAndDecrement() //获取当前的值，并自减</li>
<li>public final int getAndAdd(int delta) //获取当前的值，并加上预期的值</li>
</ul>

<p>这个参数就是我想让当前这个原子类改变多少值，可以是正数也可以是负数，如果是正数就是增加，如果是负数就是减少。而刚才的 getAndIncrement 和 getAndDecrement 修改的数值默认为 +1 或 -1，如果不能满足需求，我们就可以使用  getAndAdd 方法来直接一次性地加减我们想要的数值。</p>

<ul>
<li>boolean compareAndSet(int expect, int update) //如果输入的数值等于预期值，则以原子方式将该值更新为输入值（update）</li>
</ul>

<p>这个方法也是 CAS 的一个重要体现。</p>

<h4 id="array-数组类型原子类">Array 数组类型原子类</h4>

<p>下面我们来看第二大类 Atomic*Array 数组类型原子类，数组里的元素，都可以保证其原子性，比如 AtomicIntegerArray 相当于把 AtomicInteger 聚合起来，组合成一个数组。这样一来，我们如果想用一个每一个元素都具备原子性的数组的话， 就可以使用 Atomic*Array。</p>

<p>它一共分为 3 种，分别是：</p>

<ul>
<li>AtomicIntegerArray：整形数组原子类；</li>
<li>AtomicLongArray：长整形数组原子类；</li>
<li>AtomicReferenceArray ：引用类型数组原子类。</li>
</ul>

<h4 id="atomic-reference-引用类型原子类">Atomic\Reference 引用类型原子类</h4>

<p>下面我们介绍第三种 AtomicReference 引用类型原子类。AtomicReference 类的作用和AtomicInteger 并没有本质区别， AtomicInteger 可以让一个整数保证原子性，而AtomicReference 可以让一个对象保证原子性。这样一来，AtomicReference 的能力明显比 AtomicInteger 强，因为一个对象里可以包含很多属性。</p>

<p>在这个类别之下，除了 AtomicReference 之外，还有：</p>

<ul>
<li>AtomicStampedReference：它是对 AtomicReference 的升级，在此基础上还加了时间戳，用于解决 CAS 的 ABA 问题。</li>
<li>AtomicMarkableReference：和 AtomicReference 类似，多了一个绑定的布尔值，可以用于表示该对象已删除等场景。</li>
</ul>

<h4 id="atomic-fieldupdater-原子更新器">Atomic\FieldUpdater 原子更新器</h4>

<p>第四类我们将要介绍的是 Atomic\FieldUpdater，我们把它称为原子更新器，一共有三种，分别是。</p>

<ul>
<li>AtomicIntegerFieldUpdater：原子更新整形的更新器；</li>
<li>AtomicLongFieldUpdater：原子更新长整形的更新器；</li>
<li>AtomicReferenceFieldUpdater：原子更新引用的更新器。</li>
</ul>

<p>如果我们之前已经有了一个变量，比如是整型的 int，实际它并不具备原子性。可是木已成舟，这个变量已经被定义好了，此时我们有没有办法可以让它拥有原子性呢？办法是有的，就是利用 Atomic*FieldUpdater，如果它是整型的，就使用 AtomicIntegerFieldUpdater 把已经声明的变量进行升级，这样一来这个变量就拥有了 CAS 操作的能力。</p>

<p>这里的非互斥同步手段，是把我们已经声明好的变量进行 CAS 操作以达到同步的目的。那么你可能会想，既然想让这个变量具备原子性，为什么不在一开始就声明为 AtomicInteger？这样也免去了升级的过程，难道是一开始设计的时候不合理吗？这里有以下几种情况：</p>

<p>第一种情况是出于历史原因考虑，那么如果出于历史原因的话，之前这个变量已经被声明过了而且被广泛运用，那么修改它成本很高，所以我们可以利用升级的原子类。</p>

<p>另外还有一个使用场景，如果我们在大部分情况下并不需要使用到它的原子性，只在少数情况，比如每天只有定时一两次需要原子操作的话，我们其实没有必要把原来的变量声明为原子类型的变量，因为 AtomicInteger 比普通的变量更加耗费资源。所以如果我们有成千上万个原子类的实例的话，它占用的内存也会远比我们成千上万个普通类型占用的内存高。所以在这种情况下，我们可以利用 AtomicIntegerFieldUpdater 进行合理升级，节约内存。</p>

<p>下面我们看一段代码：</p>

<pre><code class="language-java">public class AtomicIntegerFieldUpdaterDemo implements Runnable{

   static Score math;

   static Score computer;

   public static AtomicIntegerFieldUpdater&lt;Score&gt; scoreUpdater = AtomicIntegerFieldUpdater

           .newUpdater(Score.class, &quot;score&quot;);

   @Override

   public void run() {

       for (int i = 0; i &lt; 1000; i++) {

           computer.score++;

           scoreUpdater.getAndIncrement(math);

       }

   }

   public static class Score {

       volatile int score;

   }

   public static void main(String[] args) throws InterruptedException {

       math =new Score();

       computer =new Score();

       AtomicIntegerFieldUpdaterDemo2 r = new AtomicIntegerFieldUpdaterDemo2();

       Thread t1 = new Thread(r);

       Thread t2 = new Thread(r);

       t1.start();

       t2.start();

       t1.join();

       t2.join();

       System.out.println(&quot;普通变量的结果：&quot;+ computer.score);

       System.out.println(&quot;升级后的结果：&quot;+ math.score);

   }

}
</code></pre>

<p>这段代码就演示了这个类的用法，比如说我们有两个类，它们都是 Score 类型的，Score 类型内部会有一个分数，也叫作 core，那么这两个分数的实例分别叫作数学 math 和计算机  computer，然后我们还声明了一个 AtomicIntegerFieldUpdater，在它构造的时候传入了两个参数，第一个是 Score.class，这是我们的类名，第二个是属性名，叫作 score。</p>

<p>接下来我们看一下 run 方法，run 方法里面会对这两个实例分别进行自加操作。</p>

<p>第一个是 computer，这里的 computer 我们调用的是它内部的 score，也就是说我们直接调用了 int 变量的自加操作，这在多线程下是线程非安全的。</p>

<p>第二个自加是利用了刚才声明的 scoreUpdater 并且使用了它的 getAndIncrement 方法并且传入了 math，这是一种正确使用AtomicIntegerFieldUpdater 的用法，这样可以线程安全地进行自加操作。</p>

<p>接下来我们看下 main 函数。在 main 函数中，我们首先把 math 和 computer 定义了出来，然后分别启动了两个线程，每个线程都去执行我们刚才所介绍过的 run 方法。这样一来，两个 score，也就是 math 和 computer 都会分别被加 2000 次，最后我们在 join 等待之后把结果打印了出来，这个程序的运行结果如下：</p>

<p>普通变量的结果：1942
升级后的结果：2000</p>

<p>可以看出，正如我们所预料的那样，普通变量由于不具备线程安全性，所以在多线程操作的情况下，它虽然看似进行了 2000 次操作，但有一些操作被冲突抵消了，所以最终结果小于 2000。可是使用 AtomicIntegerFieldUpdater  这个工具之后，就可以做到把一个普通类型的 score 变量进行原子的自加操作，最后的结果也和加的次数是一样的，也就是 2000。可以看出，这个类的功能还是非常强大的。</p>

<p>下面我们继续看最后两种原子类。</p>

<h3 id="adder-加法器">Adder 加法器</h3>

<p>它里面有两种加法器，分别叫作 LongAdder 和 DoubleAdder。</p>

<h4 id="accumulator-积累器">Accumulator 积累器</h4>

<p>最后一种叫 Accumulator 积累器，分别是 LongAccumulator 和 DoubleAccumulator。</p>

<p>这两种原子类我们会在后面的课时中展开介绍。</p>

<h3 id="以-atomicinteger-为例-分析在-java-中如何利用-cas-实现原子操作">以 AtomicInteger 为例，分析在 Java 中如何利用 CAS 实现原子操作？</h3>

<p>让我们回到标题中的问题，在充分了解了原子类的作用和种类之后，我们来看下  AtomicInteger 是如何通过 CAS 操作实现并发下的累加操作的，以其中一个重要方法 getAndAdd 方法为突破口。</p>

<h4 id="getandadd方法">getAndAdd方法</h4>

<p>这个方法的代码在 Java 1.8 中的实现如下：</p>

<pre><code class="language-java">//JDK 1.8实现

public final int getAndAdd(int delta) {

   return unsafe.getAndAddInt(this, valueOffset, delta);

}
</code></pre>

<p>可以看出，里面使用了 Unsafe 这个类，并且调用了 unsafe.getAndAddInt 方法。所以这里需要简要介绍一下 Unsafe 类。</p>

<h4 id="unsafe-类">Unsafe 类</h4>

<p>Unsafe 其实是 CAS 的核心类。由于 Java 无法直接访问底层操作系统，而是需要通过 native 方法来实现。不过尽管如此，JVM 还是留了一个后门，在 JDK 中有一个 Unsafe 类，它提供了硬件级别的原子操作，我们可以利用它直接操作内存数据。</p>

<p>那么我们就来看一下 AtomicInteger 的一些重要代码，如下所示：</p>

<pre><code class="language-java">public class AtomicInteger extends Number implements java.io.Serializable {

   // setup to use Unsafe.compareAndSwapInt for updates

   private static final Unsafe unsafe = Unsafe.getUnsafe();

   private static final long valueOffset;

   static {

       try {

           valueOffset = unsafe.objectFieldOffset

               (AtomicInteger.class.getDeclaredField(&quot;value&quot;));

       } catch (Exception ex) { throw new Error(ex); }

   }

   private volatile int value;

   public final int get() {return value;}

   ...

}
</code></pre>

<p>可以看出，在数据定义的部分，首先还获取了 Unsafe 实例，并且定义了 valueOffset。我们往下看到 static 代码块，这个代码块会在类加载的时候执行，执行时我们会调用 Unsafe 的 objectFieldOffset 方法，从而得到当前这个原子类的 value 的偏移量，并且赋给 valueOffset 变量，这样一来我们就获取到了 value 的偏移量，它的含义是在内存中的偏移地址，因为 Unsafe 就是根据内存偏移地址获取数据的原值的，这样我们就能通过 Unsafe 来实现 CAS 了。</p>

<p>value 是用 volatile 修饰的，它就是我们原子类存储的值的变量，由于它被 volatile 修饰，我们就可以保证在多线程之间看到的 value 是同一份，保证了可见性。</p>

<p>接下来继续看 Unsafe 的 getAndAddInt 方法的实现，代码如下：</p>

<pre><code class="language-java">public final int getAndAddInt(Object var1, long var2, int var4) {

   int var5;

   do {

       var5 = this.getIntVolatile(var1, var2);

   } while(!this.compareAndSwapInt(var1, var2, var5, var5 + var4));

   return var5;

}
</code></pre>

<p>首先我们看一下结构，它是一个 do-while 循环，所以这是一个死循环，直到满足循环的退出条件时才可以退出。</p>

<p>那么我们来看一下 do 后面的这一行代码 var5 = this.getIntVolatile(var1, var2) 是什么意思。这是个 native 方法，作用就是获取在 var1 中的 var2 偏移处的值。</p>

<p>那传入的是什么呢？传入的两个参数，第一个就是当前原子类，第二个是我们最开始获取到的 offset，这样一来我们就可以获取到当前内存中偏移量的值，并且保存到 var5 里面。此时 var5 实际上代表当前时刻下的原子类的数值。</p>

<p>现在再来看 while 的退出条件，也就是 compareAndSwapInt 这个方法，它一共传入了 4 个参数，这 4 个参数是 var1、var2、var5、var5 + var4，为了方便理解，我们给它们取了新了变量名，分别 object、offset、expectedValue、newValue，具体含义如下：</p>

<ul>
<li>第一个参数 object 就是将要操作的对象，传入的是 this，也就是 atomicInteger 这个对象本身；</li>
<li>第二个参数是 offset，也就是偏移量，借助它就可以获取到 value 的数值；</li>
<li>第三个参数 expectedValue，代表“期望值”，传入的是刚才获取到的 var5；</li>
<li>而最后一个参数 newValue 是希望修改的数值 ，等于之前取到的数值 var5 再加上 var4，而 var4 就是我们之前所传入的 delta，delta 就是我们希望原子类所改变的数值，比如可以传入 +1，也可以传入 -1。</li>
</ul>

<p>所以 compareAndSwapInt 方法的作用就是，判断如果现在原子类里 value 的值和之前获取到的 var5 相等的话，那么就把计算出来的 var5 + var4 给更新上去，所以说这行代码就实现了 CAS 的过程。</p>

<p>一旦 CAS 操作成功，就会退出这个 while 循环，但是也有可能操作失败。如果操作失败就意味着在获取到 var5 之后，并且在 CAS 操作之前，value 的数值已经发生变化了，证明有其他线程修改过这个变量。</p>

<p>这样一来，就会再次执行循环体里面的代码，重新获取 var5 的值，也就是获取最新的原子变量的数值，并且再次利用 CAS 去尝试更新，直到更新成功为止，所以这是一个死循环。</p>

<p>我们总结一下，Unsafe 的 getAndAddInt 方法是通过循环 + CAS 的方式来实现的，在此过程中，它会通过 compareAndSwapInt 方法来尝试更新 value 的值，如果更新失败就重新获取，然后再次尝试更新，直到更新成功。</p>

<h3 id="总结">总结</h3>

<p>在本课时我们首先介绍了原子类的作用，然后对 6 类原子类进行了介绍，分别是 Atomic* 基本类型原子类、Atomic*Array 数组类型原子类、Atomic*Reference 引用类型原子类、Atomic*FieldUpdater 升级类型原子类、Adder 加法器和 Accumulator 积累器。</p>

<p>然后我们对它们逐一进行了展开介绍，了解了它们的基本作用和用法，接下来我们以  AtomicInteger 为例，分析了在 Java 中是如何利用 CAS 实现原子操作的。</p>

<p>我们从 getAndAdd 方法出发，逐步深入，最后到了 Unsafe 的 getAndAddInt 方法。所以通过源码分析之后，我们也清楚地看到了，它实现的原理是利用自旋去不停地尝试，直到成功为止。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b9d5d5d5808d8888898ef9ded4d8d0d597dad6d4" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e2875babe653b',t:'MTczNDAxMTE1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
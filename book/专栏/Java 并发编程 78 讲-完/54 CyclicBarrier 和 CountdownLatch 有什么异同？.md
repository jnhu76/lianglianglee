<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=54&#32;CyclicBarrier&#32;和&#32;CountdownLatch&#32;有什么异同？>
        <link rel="icon" href="/static/favicon.png">
        <title>54 CyclicBarrier 和 CountdownLatch 有什么异同？ </title>
        
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
                            <h1 id="title" data-id="54 CyclicBarrier 和 CountdownLatch 有什么异同？" class="title">54 CyclicBarrier 和 CountdownLatch 有什么异同？</h1>
                            <div><p>本课时我们主要介绍 CyclicBarrier 和 CountDownLatch 有什么不同。</p>

<h3 id="cyclicbarrier">CyclicBarrier</h3>

<h4 id="作用">作用</h4>

<p>CyclicBarrier 和 CountDownLatch 确实有一定的相似性，它们都能阻塞一个或者一组线程，直到某种预定的条件达到之后，这些之前在等待的线程才会统一出发，继续向下执行。正因为它们有这个相似点，你可能会认为它们的作用是完全一样的，其实并不是。</p>

<p>CyclicBarrier 可以构造出一个集结点，当某一个线程执行 await() 的时候，它就会到这个集结点开始等待，等待这个栅栏被撤销。直到预定数量的线程都到了这个集结点之后，这个栅栏就会被撤销，之前等待的线程就在此刻统一出发，继续去执行剩下的任务。</p>

<p>举一个生活中的例子。假设我们班级春游去公园里玩，并且会租借三人自行车，每个人都可以骑，但由于这辆自行车是三人的，所以要凑齐三个人才能骑一辆，而且从公园大门走到自行车驿站需要一段时间。那么我们模拟这个场景，写出如下代码：</p>

<pre><code>public class CyclicBarrierDemo {

    public static void main(String[] args) {

        CyclicBarrier cyclicBarrier = new CyclicBarrier(3);

        for (int i = 0; i &lt; 6; i++) {

            new Thread(new Task(i + 1, cyclicBarrier)).start();

        }

    }

    static class Task implements Runnable {

        private int id;

        private CyclicBarrier cyclicBarrier;

        public Task(int id, CyclicBarrier cyclicBarrier) {

            this.id = id;

            this.cyclicBarrier = cyclicBarrier;

        }

        @Override

        public void run() {

            System.out.println(&quot;同学&quot; + id + &quot;现在从大门出发，前往自行车驿站&quot;);

            try {

                Thread.sleep((long) (Math.random() * 10000));

                System.out.println(&quot;同学&quot; + id + &quot;到了自行车驿站，开始等待其他人到达&quot;);

                cyclicBarrier.await();

                System.out.println(&quot;同学&quot; + id + &quot;开始骑车&quot;);

            } catch (InterruptedException e) {

                e.printStackTrace();

            } catch (BrokenBarrierException e) {

                e.printStackTrace();

            }

        }

    }

}
</code></pre>

<p>在这段代码中可以看到，首先建了一个参数为 3 的 CyclicBarrier，参数为 3 的意思是需要等待 3 个线程到达这个集结点才统一放行；然后我们又在 for 循环中去开启了 6 个线程，每个线程中执行的 Runnable 对象就在下方的 Task 类中，直接看到它的 run 方法，它首先会打印出&rdquo;同学某某现在从大门出发，前往自行车驿站&rdquo;，然后是一个随机时间的睡眠，这就代表着从大门开始步行走到自行车驿站的时间，由于每个同学的步行速度不一样，所以时间用随机值来模拟。</p>

<p>当同学们都到了驿站之后，比如某一个同学到了驿站，首先会打印出“同学某某到了自行车驿站，开始等待其他人到达”的消息，然后去调用 CyclicBarrier 的 await() 方法。一旦它调用了这个方法，它就会陷入等待，直到三个人凑齐，才会继续往下执行，一旦开始继续往下执行，就意味着 3 个同学开始一起骑车了，所以打印出“某某开始骑车”这个语句。</p>

<p>接下来我们运行一下这个程序，结果如下所示：</p>

<pre><code>同学1现在从大门出发，前往自行车驿站

同学3现在从大门出发，前往自行车驿站

同学2现在从大门出发，前往自行车驿站

同学4现在从大门出发，前往自行车驿站

同学5现在从大门出发，前往自行车驿站

同学6现在从大门出发，前往自行车驿站

同学5到了自行车驿站，开始等待其他人到达

同学2到了自行车驿站，开始等待其他人到达

同学3到了自行车驿站，开始等待其他人到达

同学3开始骑车

同学5开始骑车

同学2开始骑车

同学6到了自行车驿站，开始等待其他人到达

同学4到了自行车驿站，开始等待其他人到达

同学1到了自行车驿站，开始等待其他人到达

同学1开始骑车

同学6开始骑车

同学4开始骑车
</code></pre>

<p>可以看到 6 个同学纷纷从大门出发走到自行车驿站，因为每个人的速度不一样，所以会有 3 个同学先到自行车驿站，不过在这 3 个先到的同学里面，前面 2 个到的都必须等待第 3 个人到齐之后，才可以开始骑车。后面的同学也一样，由于第一辆车已经被骑走了，第二辆车依然也要等待 3 个人凑齐才能统一发车。</p>

<p>要想实现这件事情，如果你不利用 CyclicBarrier 去做的话，逻辑可能会非常复杂，因为你也不清楚哪个同学先到、哪个后到。而用了 CyclicBarrier 之后，可以非常简洁优雅的实现这个逻辑，这就是它的一个非常典型的应用场景。</p>

<h4 id="执行动作-barrieraction">执行动作 barrierAction</h4>

<p>public CyclicBarrier(int parties, Runnable barrierAction)：当 parties 线程到达集结点时，继续往下执行前，会执行这一次这个动作。</p>

<p>接下来我们再介绍一下它的一个额外功能，就是执行动作 barrierAction 功能。CyclicBarrier 还有一个构造函数是传入两个参数的，第一个参数依然是 parties，代表需要几个线程到齐；第二个参数是一个 Runnable 对象，它就是我们下面所要介绍的 barrierAction。</p>

<p>当预设数量的线程到达了集结点之后，在出发的时候，便会执行这里所传入的 Runnable 对象，那么假设我们把刚才那个代码的构造函数改成如下这个样子：</p>

<pre><code>CyclicBarrier cyclicBarrier = new CyclicBarrier(3, new Runnable() {

    @Override

    public void run() {

        System.out.println(&quot;凑齐3人了，出发！&quot;);

    }

});
</code></pre>

<p>可以看出，我们传入了第二个参数，它是一个 Runnable 对象，在这里传入了这个 Runnable 之后，这个任务就会在到齐的时候去打印&rdquo;凑齐3人了，出发！&rdquo;。上面的代码如果改成这个样子，则执行结果如下所示：</p>

<pre><code>同学1现在从大门出发，前往自行车驿站

同学3现在从大门出发，前往自行车驿站

同学2现在从大门出发，前往自行车驿站

同学4现在从大门出发，前往自行车驿站

同学5现在从大门出发，前往自行车驿站

同学6现在从大门出发，前往自行车驿站

同学2到了自行车驿站，开始等待其他人到达

同学4到了自行车驿站，开始等待其他人到达

同学6到了自行车驿站，开始等待其他人到达

凑齐3人了，出发！

同学6开始骑车

同学2开始骑车

同学4开始骑车

同学1到了自行车驿站，开始等待其他人到达

同学3到了自行车驿站，开始等待其他人到达

同学5到了自行车驿站，开始等待其他人到达

凑齐3人了，出发！

同学5开始骑车

同学1开始骑车

同学3开始骑车
</code></pre>

<p>可以看出，三个人凑齐了一组之后，就会打印出“凑齐 3 人了，出发！”这样的语句，该语句恰恰是我们在这边传入 Runnable 所执行的结果。</p>

<p>值得注意的是，这个语句每个周期只打印一次，不是说你有几个线程在等待就打印几次，而是说这个任务只在“开闸”的时候执行一次。</p>

<h3 id="cyclicbarrier-和-countdownlatch-的异同">CyclicBarrier 和 CountDownLatch 的异同</h3>

<p>下面我们来总结一下 CyclicBarrier 和 CountDownLatch 有什么异同。</p>

<p>相同点：都能阻塞一个或一组线程，直到某个预设的条件达成发生，再统一出发。</p>

<p>但是它们也有很多不同点，具体如下。</p>

<ul>
<li><strong>作用对象不同</strong>：CyclicBarrier 要等固定数量的线程都到达了栅栏位置才能继续执行，而 CountDownLatch 只需等待数字倒数到 0，也就是说 CountDownLatch 作用于事件，但 CyclicBarrier 作用于线程；CountDownLatch 是在调用了 countDown 方法之后把数字倒数减 1，而 CyclicBarrier 是在某线程开始等待后把计数减 1。</li>
<li><strong>可重用性不同</strong>：CountDownLatch 在倒数到 0  并且触发门闩打开后，就不能再次使用了，除非新建一个新的实例；而 CyclicBarrier 可以重复使用，在刚才的代码中也可以看出，每 3 个同学到了之后都能出发，并不需要重新新建实例。CyclicBarrier 还可以随时调用 reset 方法进行重置，如果重置时有线程已经调用了 await 方法并开始等待，那么这些线程则会抛出 BrokenBarrierException 异常。</li>
<li><strong>执行动作不同</strong>：CyclicBarrier 有执行动作 barrierAction，而 CountDownLatch 没这个功能。</li>
</ul>

<h3 id="总结">总结</h3>

<p>以上就是本课时的内容，在本课时中，首先介绍了 CyclicBarrier 的作用、代码示例和执行动作，然后对 CyclicBarrier 和 CountDownLatch 的异同进行了总结。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#573b3b3b6e636666676017303a363e3b7934383a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e2be89a51653b',t:'MTczNDAxMTMwMS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
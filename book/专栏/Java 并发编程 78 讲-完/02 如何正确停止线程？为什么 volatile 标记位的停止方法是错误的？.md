<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=02&#32;如何正确停止线程？为什么&#32;volatile&#32;标记位的停止方法是错误的？>
        <link rel="icon" href="/static/favicon.png">
        <title>02 如何正确停止线程？为什么 volatile 标记位的停止方法是错误的？ </title>
        
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
                            <h1 id="title" data-id="02 如何正确停止线程？为什么 volatile 标记位的停止方法是错误的？" class="title">02 如何正确停止线程？为什么 volatile 标记位的停止方法是错误的？</h1>
                            <div><p>在本课时我们主要学习如何正确停止一个线程？以及为什么用 volatile 标记位的停止方法是错误的？</p>

<p>首先，我们来复习如何启动一个线程，想要启动线程需要调用 Thread 类的 start() 方法，并在 run() 方法中定义需要执行的任务。启动一个线程非常简单，但如果想要正确停止它就没那么容易了。</p>

<h3 id="原理介绍">原理介绍</h3>

<p>通常情况下，我们不会手动停止一个线程，而是允许线程运行到结束，然后让它自然停止。但是依然会有许多特殊的情况需要我们提前停止线程，比如：用户突然关闭程序，或程序运行出错重启等。</p>

<p>在这种情况下，即将停止的线程在很多业务场景下仍然很有价值。尤其是我们想写一个健壮性很好，能够安全应对各种场景的程序时，正确停止线程就显得格外重要。但是Java 并没有提供简单易用，能够直接安全停止线程的能力。</p>

<h3 id="为什么不强制停止-而是通知-协作">为什么不强制停止？而是通知、协作</h3>

<p>对于 Java 而言，最正确的停止线程的方式是使用 interrupt。但 interrupt 仅仅起到通知被停止线程的作用。而对于被停止的线程而言，它拥有完全的自主权，它既可以选择立即停止，也可以选择一段时间后停止，也可以选择压根不停止。那么为什么 Java 不提供强制停止线程的能力呢？</p>

<p>事实上，Java 希望程序间能够相互通知、相互协作地管理线程，因为如果不了解对方正在做的工作，贸然强制停止线程就可能会造成一些安全的问题，为了避免造成问题就需要给对方一定的时间来整理收尾工作。比如：线程正在写入一个文件，这时收到终止信号，它就需要根据自身业务判断，是选择立即停止，还是将整个文件写入成功后停止，而如果选择立即停止就可能造成数据不完整，不管是中断命令发起者，还是接收者都不希望数据出现问题。</p>

<h3 id="如何用-interrupt-停止线程">如何用 interrupt 停止线程</h3>

<pre><code>while (!Thread.currentThread().isInterrupted() &amp;&amp; more work to do) {

    do more work

}
</code></pre>

<p>明白 Java 停止线程的设计原则之后，我们看看如何用代码实现停止线程的逻辑。我们一旦调用某个线程的 interrupt() 之后，这个线程的中断标记位就会被设置成 true。每个线程都有这样的标记位，当线程执行时，应该定期检查这个标记位，如果标记位被设置成 true，就说明有程序想终止该线程。回到源码，可以看到在 while 循环体判断语句中，首先通过 Thread.currentThread().isInterrupt() 判断线程是否被中断，随后检查是否还有工作要做。&amp;&amp; 逻辑表示只有当两个判断条件同时满足的情况下，才会去执行下面的工作。</p>

<p>我们再看看具体例子。</p>

<pre><code>public class StopThread implements Runnable {

    @Override

    public void run() {

        int count = 0;

        while (!Thread.currentThread().isInterrupted() &amp;&amp; count &lt; 1000) {

            System.out.println(&quot;count = &quot; + count++);

        }

    }

    public static void main(String[] args) throws InterruptedException {

        Thread thread = new Thread(new StopThread());

        thread.start();

        Thread.sleep(5);

        thread.interrupt();

    }

}
</code></pre>

<p>在 StopThread 类的 run() 方法中，首先判断线程是否被中断，然后判断 count 值是否小于 1000。这个线程的工作内容很简单，就是打印 0~999 的数字，每打印一个数字 count 值加 1，可以看到，线程会在每次循环开始之前，检查是否被中断了。接下来在 main 函数中会启动该线程，然后休眠 5 毫秒后立刻中断线程，该线程会检测到中断信号，于是在还没打印完1000个数的时候就会停下来，这种就属于通过 interrupt 正确停止线程的情况。</p>

<h3 id="sleep-期间能否感受到中断">sleep 期间能否感受到中断</h3>

<pre><code>Runnable runnable = () -&gt; {

    int num = 0;

    try {

        while (!Thread.currentThread().isInterrupted() &amp;&amp; 

        num &lt;= 1000) {

            System.out.println(num);

            num++;

            Thread.sleep(1000000);

        }

    } catch (InterruptedException e) {

        e.printStackTrace();

    }

};
</code></pre>

<p>那么我们考虑一种特殊情况，改写上面的代码，如果线程在执行任务期间有休眠需求，也就是每打印一个数字，就进入一次 sleep ，而此时将 Thread.sleep() 的休眠时间设置为 1000 秒钟。</p>

<pre><code>public class StopDuringSleep {

    public static void main(String[] args) throws InterruptedException {

        Runnable runnable = () -&gt; {

            int num = 0;

            try {

                while (!Thread.currentThread().isInterrupted() &amp;&amp; num &lt;= 1000) {

                    System.out.println(num);

                    num++;

                    Thread.sleep(1000000);

                }

            } catch (InterruptedException e) {

                e.printStackTrace();

            }

        };

        Thread thread = new Thread(runnable);

        thread.start();

        Thread.sleep(5);

        thread.interrupt();

    }

}
</code></pre>

<p>主线程休眠 5 毫秒后，通知子线程中断，此时子线程仍在执行 sleep 语句，处于休眠中。那么就需要考虑一点，在休眠中的线程是否能够感受到中断通知呢？是否需要等到休眠结束后才能中断线程呢？如果是这样，就会带来严重的问题，因为响应中断太不及时了。正因为如此，Java 设计者在设计之初就考虑到了这一点。</p>

<p>如果 sleep、wait 等可以让线程进入阻塞的方法使线程休眠了，而处于休眠中的线程被中断，那么线程是可以感受到中断信号的，并且会抛出一个 InterruptedException 异常，同时清除中断信号，将中断标记位设置成 false。这样一来就不用担心长时间休眠中线程感受不到中断了，因为即便线程还在休眠，仍然能够响应中断通知，并抛出异常。</p>

<h3 id="两种最佳处理方式">两种最佳处理方式</h3>

<p>在实际开发中肯定是团队协作的，不同的人负责编写不同的方法，然后相互调用来实现整个业务的逻辑。那么如果我们负责编写的方法需要被别人调用，同时我们的方法内调用了 sleep 或者 wait 等能响应中断的方法时，仅仅 catch 住异常是不够的。</p>

<pre><code>void subTas() {

    try {

        Thread.sleep(1000);

    } catch (InterruptedException e) {

        // 在这里不处理该异常是非常不好的

    }

}
</code></pre>

<p>我们可以在方法中使用 try/catch 或在方法签名中声明 throws  InterruptedException。</p>

<h3 id="方法签名抛异常-run-强制-try-catch">方法签名抛异常，run() 强制 try/catch</h3>

<p>我们先来看下 try/catch 的处理逻辑。如上面的代码所示，catch 语句块里代码是空的，它并没有进行任何处理。假设线程执行到这个方法，并且正在 sleep，此时有线程发送 interrupt 通知试图中断线程，就会立即抛出异常，并清除中断信号。抛出的异常被 catch 语句块捕捉。</p>

<p>但是，捕捉到异常的 catch 没有进行任何处理逻辑，相当于把中断信号给隐藏了，这样做是非常不合理的，那么究竟应该怎么处理呢？首先，可以选择在方法签名中抛出异常。</p>

<pre><code>void subTask2() throws InterruptedException {

    Thread.sleep(1000);

}
</code></pre>

<p>正如代码所示，要求每一个方法的调用方有义务去处理异常。调用方要不使用 try/catch 并在 catch 中正确处理异常，要不将异常声明到方法签名中。如果每层逻辑都遵守规范，便可以将中断信号层层传递到顶层，最终让 run() 方法可以捕获到异常。而对于 run() 方法而言，它本身没有抛出 checkedException 的能力，只能通过 try/catch 来处理异常。层层传递异常的逻辑保障了异常不会被遗漏，而对 run() 方法而言，就可以根据不同的业务逻辑来进行相应的处理。</p>

<h3 id="再次中断">再次中断</h3>

<pre><code>private void reInterrupt() {

    try {

        Thread.sleep(2000);

    } catch (InterruptedException e) {

        Thread.currentThread().interrupt();

        e.printStackTrace();

    }

}
</code></pre>

<p>除了刚才推荐的将异常声明到方法签名中的方式外，还可以在 catch 语句中再次中断线程。如代码所示，需要在 catch 语句块中调用 Thread.currentThread().interrupt() 函数。因为如果线程在休眠期间被中断，那么会自动清除中断信号。如果这时手动添加中断信号，中断信号依然可以被捕捉到。这样后续执行的方法依然可以检测到这里发生过中断，可以做出相应的处理，整个线程可以正常退出。</p>

<p>我们需要注意，我们在实际开发中不能盲目吞掉中断，如果不在方法签名中声明，也不在 catch 语句块中再次恢复中断，而是在 catch 中不作处理，我们称这种行为是“屏蔽了中断请求”。如果我们盲目地屏蔽了中断请求，会导致中断信号被完全忽略，最终导致线程无法正确停止。</p>

<h3 id="为什么用-volatile-标记位的停止方法是错误的">为什么用 volatile 标记位的停止方法是错误的</h3>

<p>下面我们来看一看本课时的第二个问题，为什么用 volatile 标记位的停止方法是错误的？</p>

<h4 id="错误的停止方法">错误的停止方法</h4>

<p>首先，我们来看几种停止线程的错误方法。比如 stop()，suspend() 和 resume()，这些方法已经被 Java 直接标记为 @Deprecated。如果再调用这些方法，IDE 会友好地提示，我们不应该再使用它们了。但为什么它们不能使用了呢？是因为 stop() 会直接把线程停止，这样就没有给线程足够的时间来处理想要在停止前保存数据的逻辑，任务戛然而止，会导致出现数据完整性等问题。</p>

<p>而对于 suspend() 和 resume() 而言，它们的问题在于如果线程调用 suspend()，它并不会释放锁，就开始进入休眠，但此时有可能仍持有锁，这样就容易导致死锁问题，因为这把锁在线程被 resume() 之前，是不会被释放的。</p>

<p>假设线程 A 调用了 suspend() 方法让线程 B 挂起，线程 B 进入休眠，而线程 B 又刚好持有一把锁，此时假设线程 A 想访问线程 B 持有的锁，但由于线程 B 并没有释放锁就进入休眠了，所以对于线程 A 而言，此时拿不到锁，也会陷入阻塞，那么线程 A 和线程 B 就都无法继续向下执行。</p>

<p>正是因为有这样的风险，所以 suspend() 和 resume() 组合使用的方法也被废弃了。那么接下来我们来看看，为什么用 volatile 标记位的停止方法也是错误的？</p>

<h4 id="volatile-修饰标记位适用的场景">volatile 修饰标记位适用的场景</h4>

<pre><code>public class VolatileCanStop implements Runnable {

    private volatile boolean canceled = false;

    @Override

    public void run() {

        int num = 0;

        try {

            while (!canceled &amp;&amp; num &lt;= 1000000) {

                if (num % 10 == 0) {

                    System.out.println(num + &quot;是10的倍数。&quot;);

                }

                num++;

                Thread.sleep(1);

            }

        } catch (InterruptedException e) {

            e.printStackTrace();

        }

    }

    public static void main(String[] args) throws InterruptedException {

        VolatileCanStop r = new VolatileCanStop();

        Thread thread = new Thread(r);

        thread.start();

        Thread.sleep(3000);

        r.canceled = true;

    }

}
</code></pre>

<p>什么场景下 volatile 修饰标记位可以让线程正常停止呢？如代码所示，声明了一个叫作 VolatileStopThread 的类， 它实现了 Runnable 接口，然后在 run() 中进行 while 循环，在循环体中又进行了两层判断，首先判断 canceled 变量的值，canceled 变量是一个被 volatile 修饰的初始值为 false 的布尔值，当该值变为 true 时，while 跳出循环，while 的第二个判断条件是 num 值小于1000000（一百万），在while 循环体里，只要是 10 的倍数就打印出来，然后 num++。</p>

<p>接下来，首先启动线程，然后经过 3 秒钟的时间，把用 volatile 修饰的布尔值的标记位设置成 true，这样，正在运行的线程就会在下一次 while 循环中判断出 canceled 的值已经变成 true 了，这样就不再满足 while 的判断条件，跳出整个 while 循环，线程就停止了，这种情况是演示 volatile 修饰的标记位可以正常工作的情况，但是如果我们说某个方法是正确的，那么它应该不仅仅是在一种情况下适用，而在其他情况下也应该是适用的。</p>

<h4 id="volatile-修饰标记位不适用的场景">volatile 修饰标记位不适用的场景</h4>

<p>接下来我们就用一个生产者/消费者模式的案例来演示为什么说  volatile 标记位的停止方法是不完美的。</p>

<pre><code>class Producer implements Runnable {

    public volatile boolean canceled = false;

    BlockingQueue storage;

    public Producer(BlockingQueue storage) {

        this.storage = storage;

    }

    @Override

    public void run() {

        int num = 0;

        try {

            while (num &lt;= 100000 &amp;&amp; !canceled) {

                if (num % 50 == 0) {

                    storage.put(num);

                    System.out.println(num + &quot;是50的倍数,被放到仓库中了。&quot;);

                }

                num++;

            }

        } catch (InterruptedException e) {

            e.printStackTrace();

        } finally {

            System.out.println(&quot;生产者结束运行&quot;);

        }

    }

}
</code></pre>

<p>首先，声明了一个生产者 Producer，通过 volatile 标记的初始值为 false 的布尔值 canceled 来停止线程。而在 run() 方法中，while 的判断语句是 num 是否小于 100000 及 canceled 是否被标记。while 循环体中判断 num 如果是 50 的倍数就放到 storage 仓库中，storage 是生产者与消费者之间进行通信的存储器，当 num 大于 100000 或被通知停止时，会跳出 while 循环并执行 finally 语句块，告诉大家“生产者结束运行”。</p>

<pre><code>class Consumer {

    BlockingQueue storage;

    public Consumer(BlockingQueue storage) {

        this.storage = storage;

    }

    public boolean needMoreNums() {

        if (Math.random() &gt; 0.97) {

            return false;

        }

        return true;

    }

}
</code></pre>

<p>而对于消费者 Consumer，它与生产者共用同一个仓库 storage，并且在方法内通过 needMoreNums() 方法判断是否需要继续使用更多的数字，刚才生产者生产了一些 50 的倍数供消费者使用，消费者是否继续使用数字的判断条件是产生一个随机数并与 0.97 进行比较，大于 0.97 就不再继续使用数字。</p>

<pre><code>public static void main(String[] args) throws InterruptedException {

        ArrayBlockingQueue storage = new ArrayBlockingQueue(8);

        Producer producer = new Producer(storage);

        Thread producerThread = new Thread(producer);

        producerThread.start();

        Thread.sleep(500);

        Consumer consumer = new Consumer(storage);

        while (consumer.needMoreNums()) {

            System.out.println(consumer.storage.take() + &quot;被消费了&quot;);

            Thread.sleep(100);

        }

        System.out.println(&quot;消费者不需要更多数据了。&quot;);

        //一旦消费不需要更多数据了，我们应该让生产者也停下来，但是实际情况却停不下来

        producer.canceled = true;

        System.out.println(producer.canceled);

    }

}
</code></pre>

<p>下面来看下 main 函数，首先创建了生产者/消费者共用的仓库 BlockingQueue storage，仓库容量是 8，并且建立生产者并将生产者放入线程后启动线程，启动后进行 500 毫秒的休眠，休眠时间保障生产者有足够的时间把仓库塞满，而仓库达到容量后就不会再继续往里塞，这时生产者会阻塞，500 毫秒后消费者也被创建出来，并判断是否需要使用更多的数字，然后每次消费后休眠 100 毫秒，这样的业务逻辑是有可能出现在实际生产中的。</p>

<p>当消费者不再需要数据，就会将 canceled 的标记位设置为 true，理论上此时生产者会跳出 while 循环，并打印输出“生产者运行结束”。</p>

<p>然而结果却不是我们想象的那样，尽管已经把 canceled 设置成 true，但生产者仍然没有停止，这是因为在这种情况下，生产者在执行 storage.put(num) 时发生阻塞，在它被叫醒之前是没有办法进入下一次循环判断 canceled 的值的，所以在这种情况下用 volatile 是没有办法让生产者停下来的，相反如果用 interrupt 语句来中断，即使生产者处于阻塞状态，仍然能够感受到中断信号，并做响应处理。</p>

<h3 id="总结">总结</h3>

<p>好了，本课时的内容就全部讲完了，我们来总结下学到了什么，首先学习了如何正确停止线程，其次是掌握了为什么说 volatile 修饰标记位停止方法是错误的。</p>

<p>如果我们在面试中被问到“你知不知道如何正确停止线程”这样的问题，我想你一定可以完美地回答了，首先，从原理上讲应该用 interrupt 来请求中断，而不是强制停止，因为这样可以避免数据错乱，也可以让线程有时间结束收尾工作。</p>

<p>如果我们是子方法的编写者，遇到了 interruptedException，应该如何处理呢？</p>

<p>我们可以把异常声明在方法中，以便顶层方法可以感知捕获到异常，或者也可以在 catch 中再次声明中断，这样下次循环也可以感知中断，所以要想正确停止线程就要求我们停止方，被停止方，子方法的编写者相互配合，大家都按照一定的规范来编写代码，就可以正确地停止线程了。</p>

<p>最后我们再来看下有哪些方法是不够好的，比如说已经被舍弃的 stop()、suspend() 和 resume()，它们由于有很大的安全风险比如死锁风险而被舍弃，而 volatile 这种方法在某些特殊的情况下，比如线程被长时间阻塞的情况，就无法及时感受中断，所以 volatile 是不够全面的停止线程的方法。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b1dddddd888580808186f1d6dcd0d8dd9fd2dedc" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e20419e69653b',t:'MTczNDAxMDgyMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
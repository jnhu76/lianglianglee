<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;HashMap&#32;为什么是线程不安全的？>
        <link rel="icon" href="/static/favicon.png">
        <title>29 HashMap 为什么是线程不安全的？ </title>
        
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
                            <h1 id="title" data-id="29 HashMap 为什么是线程不安全的？" class="title">29 HashMap 为什么是线程不安全的？</h1>
                            <div><p>本课时我们主要讲解为什么 HashMap 是线程不安全的？而对于 HashMap，相信你一定并不陌生，HashMap 是我们平时工作和学习中用得非常非常多的一个容器，也是 Map 最主要的实现类之一，但是它自身并不具备线程安全的特点，可以从多种情况中体现出来，下面我们就对此进行具体的分析。</p>

<h3 id="源码分析">源码分析</h3>

<p>第一步，我们来看一下 HashMap 中 put 方法的源码：</p>

<pre><code class="language-java">public V put(K key, V value) {

    if (key == null)

        return putForNullKey(value);

    int hash = hash(key.hashCode());

    int i = indexFor(hash, table.length);

    for (Entry&lt;K,V&gt; e = table[i]; e != null; e = e.next) {

        Object k;

        if (e.hash == hash &amp;&amp; ((k = e.key) == key || key.equals(k))) {

            V oldValue = e.value;

            e.value = value;

            e.recordAccess(this);

            return oldValue;

        }

    } 

    //modCount++ 是一个复合操作

    modCount++;

    addEntry(hash, key, value, i);

    return null;

}
</code></pre>

<p>在 HashMap 的 put() 方法中，可以看出里面进行了很多操作，那么在这里，我们把目光聚焦到标记出来的 modCount++ 这一行代码中，相信有经验的小伙伴一定发现了，这相当于是典型的“i++”操作，正是我们在 06 课时讲过的线程不安全的“运行结果错误”的情况。从表面上看 i++ 只是一行代码，但实际上它并不是一个原子操作，它的执行步骤主要分为三步，而且在每步操作之间都有可能被打断。</p>

<ul>
<li>第一个步骤是读取；</li>
<li>第二个步骤是增加；</li>
<li>第三个步骤是保存。</li>
</ul>

<p>那么我们接下来具体看一下如何发生的线程不安全问题。</p>

<p><img src="assets/Cgq2xl4YRJeAC6fuAAA8JO4TxM0077.png" alt="img" /></p>

<p>我们根据箭头指向依次看，假设线程 1 首先拿到 i=1 的结果，然后进行 i+1 操作，但此时 i+1 的结果并没有保存下来，线程 1 就被切换走了，于是 CPU 开始执行线程 2，它所做的事情和线程 1 是一样的 i++ 操作，但此时我们想一下，它拿到的 i 是多少？实际上和线程 1 拿到的 i 的结果一样都是 1，为什么呢？因为线程 1 虽然对 i 进行了 +1 操作，但结果没有保存，所以线程 2 看不到修改后的结果。</p>

<p>然后假设等线程 2 对 i 进行 +1 操作后，又切换到线程 1，让线程 1 完成未完成的操作，即将 i + 1 的结果 2 保存下来，然后又切换到线程 2 完成 i = 2 的保存操作，虽然两个线程都执行了对 i 进行 +1 的操作，但结果却最终保存了 i = 2 的结果，而不是我们期望的 i = 3，这样就发生了线程安全问题，导致了数据结果错误，这也是最典型的线程安全问题。</p>

<p>所以，从源码的角度，或者说从理论上来讲，这完全足以证明 HashMap 是线程非安全的了。因为如果有多个线程同时调用 put() 方法的话，它很有可能会把 modCount 的值计算错（上述的源码分析针对的是 Java 7 版本的源码，而在 Java 8 版本的 HashMap 的 put 方法中会调用 putVal 方法，里面同样有 ++modCount 语句，所以原理是一样的）。</p>

<h3 id="实验-扩容期间取出的值不准确">实验：扩容期间取出的值不准确</h3>

<p>刚才我们分析了源码，你可能觉得不过瘾，下面我们就打开代码编辑器，用一个实验来证明 HashMap 是线程不安全的。</p>

<p>为什么说 HashMap 不是线程安全的呢？我们先来讲解下原理。HashMap 本身默认的容量不是很大，如果不停地往 map 中添加新的数据，它便会在合适的时机进行扩容。而在扩容期间，它会新建一个新的空数组，并且用旧的项填充到这个新的数组中去。那么，在这个填充的过程中，如果有线程获取值，很可能会取到 null 值，而不是我们所希望的、原来添加的值。所以我们程序就想演示这种情景，我们来看一下这段代码：</p>

<pre><code class="language-java">public class HashMapNotSafe {

    public static void main(String[] args) {

        final Map&lt;Integer, String&gt; map = new HashMap&lt;&gt;();

        final Integer targetKey = 0b1111_1111_1111_1111; // 65 535

        final String targetValue = &quot;v&quot;;

        map.put(targetKey, targetValue);

        new Thread(() -&gt; {

            IntStream.range(0, targetKey).forEach(key -&gt; map.put(key, &quot;someValue&quot;));

        }).start();

        while (true) {

            if (null == map.get(targetKey)) {

                throw new RuntimeException(&quot;HashMap is not thread safe.&quot;);

            }

        }

    }

}
</code></pre>

<p>代码中首先建立了一个 HashMap，并且定义了 key 和 value， key 的值是一个二进制的 1111_1111_1111_1111，对应的十进制是 65535。之所以选取这样的值，就是为了让它在扩容往回填充数据的时候，尽量不要填充得太快，比便于我们能捕捉到错误的发生。而对应的 value 是无所谓的，我们随意选取了一个非 null 的 &ldquo;v&rdquo; 来表示它，并且把这个值放到了 map 中。</p>

<p>接下来，我们就用一个新的线程不停地往我们的 map 中去填入新的数据，我们先来看是怎么填入的。首先它用了一个 IntStream，这个 range 是从 0 到之前所讲过的 65535，这个 range 是一个左闭右开的区间，所以会从 0、1、2、3……一直往上加，并且每一次加的时候，这个 0、1、2、3、4 都会作为 key 被放到 map 中去。而它的 value 是统一的，都是 &ldquo;someValue&rdquo;，因为 value 不是我们所关心的。</p>

<p>然后，我们就会把这个线程启动起来，随后就进入一个 while 循环，这个 while 循环是关键，在 while 循环中我们会不停地检测之前放入的 key 所对应的 value 还是不是我们所期望的字符串 &ldquo;v&rdquo;。我们在 while 循环中会不停地从 map 中取 key 对应的值。如果 HashMap 是线程安全的，那么无论怎样它所取到的值都应该是我们最开始放入的字符串 &ldquo;v&rdquo;，可是如果取出来是一个 null，就会满足这个 if 条件并且随即抛出一个异常，因为如果取出 null 就证明它所取出来的值和我们一开始放入的值是不一致的，也就证明了它是线程不安全的，所以在此我们要抛出一个 RuntimeException 提示我们。</p>

<p>下面就让我们运行这个程序来看一看是否会抛出这个异常。一旦抛出就代表它是线程不安全的，这段代码的运行结果：</p>

<pre><code class="language-java">Exception in thread &quot;main&quot; java.lang.RuntimeException: HashMap is not thread safe.

at lesson29.HashMapNotSafe.main(HashMapNotSafe.java:25)
</code></pre>

<p>很明显，很快这个程序就抛出了我们所希望看到的 RuntimeException，并且我们把它描述为：HashMap is not thread safe，一旦它能进入到这个 if 语句，就已经证明它所取出来的值是 null，而不是我们期望的字符串 &ldquo;v&rdquo;。</p>

<p>通过以上这个例子，我们也证明了HashMap 是线程非安全的。</p>

<p>除了刚才的例子之外，还有很多种线程不安全的情况，例如：</p>

<h4 id="同时-put-碰撞导致数据丢失">同时 put 碰撞导致数据丢失</h4>

<p>比如，有多个线程同时使用 put 来添加元素，而且恰好两个 put 的 key 是一样的，它们发生了碰撞，也就是根据 hash 值计算出来的 bucket 位置一样，并且两个线程又同时判断该位置是空的，可以写入，所以这两个线程的两个不同的 value 便会添加到数组的同一个位置，这样最终就只会保留一个数据，丢失一个数据。</p>

<h4 id="可见性问题无法保证">可见性问题无法保证</h4>

<p>我们再从可见性的角度去考虑一下。可见性也是线程安全的一部分，如果某一个数据结构声称自己是线程安全的，那么它同样需要保证可见性，也就是说，当一个线程操作这个容器的时候，该操作需要对另外的线程都可见，也就是其他线程都能感知到本次操作。可是 HashMap 对此是做不到的，如果线程 1 给某个 key 放入了一个新值，那么线程 2 在获取对应的 key 的值的时候，它的可见性是无法保证的，也就是说线程 2 可能可以看到这一次的更改，但也有可能看不到。所以从可见性的角度出发，HashMap 同样是线程非安全的。</p>

<h4 id="死循环造成-cpu-100">死循环造成 CPU 100%</h4>

<p>下面我们再举一个死循环造成 CPU 100% 的例子。HashMap 有可能会发生死循环并且造成  CPU 100% ，这种情况发生最主要的原因就是在扩容的时候，也就是内部新建新的 HashMap 的时候，扩容的逻辑会反转散列桶中的节点顺序，当有多个线程同时进行扩容的时候，由于 HashMap 并非线程安全的，所以如果两个线程同时反转的话，便可能形成一个循环，并且这种循环是链表的循环，相当于 A 节点指向 B 节点，B 节点又指回到 A 节点，这样一来，在下一次想要获取该 key 所对应的 value 的时候，便会在遍历链表的时候发生永远无法遍历结束的情况，也就发生 CPU 100% 的情况。</p>

<p>所以综上所述，HashMap 是线程不安全的，在多线程使用场景中如果需要使用 Map，应该尽量避免使用线程不安全的 HashMap。同时，虽然 Collections.synchronizedMap(new HashMap()) 是线程安全的，但是效率低下，因为内部用了很多的 synchronized，多个线程不能同时操作。推荐使用线程安全同时性能比较好的 ConcurrentHashMap。关于 ConcurrentHashMap 我们会在下一个课时中介绍。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#5e323232676a6f6f6e691e39333f3732703d3133" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e26597ad9653b',t:'MTczNDAxMTA3My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
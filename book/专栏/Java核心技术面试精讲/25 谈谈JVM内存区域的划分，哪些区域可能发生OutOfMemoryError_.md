<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=25&#32;谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_>
        <link rel="icon" href="/static/favicon.png">
        <title>25 谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_ </title>
        
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
                        <a class="menu-item" id="00 开篇词 以面试题为切入点，有效提升你的Java内功.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%a5%e9%9d%a2%e8%af%95%e9%a2%98%e4%b8%ba%e5%88%87%e5%85%a5%e7%82%b9%ef%bc%8c%e6%9c%89%e6%95%88%e6%8f%90%e5%8d%87%e4%bd%a0%e7%9a%84Java%e5%86%85%e5%8a%9f.md">00 开篇词 以面试题为切入点，有效提升你的Java内功.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 谈谈你对Java平台的理解？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/01%20%e8%b0%88%e8%b0%88%e4%bd%a0%e5%af%b9Java%e5%b9%b3%e5%8f%b0%e7%9a%84%e7%90%86%e8%a7%a3%ef%bc%9f.md">01 谈谈你对Java平台的理解？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Exception和Error有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/02%20Exception%e5%92%8cError%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">02 Exception和Error有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 谈谈final、finally、 finalize有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/03%20%e8%b0%88%e8%b0%88final%e3%80%81finally%e3%80%81%20finalize%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">03 谈谈final、finally、 finalize有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 强引用、软引用、弱引用、幻象引用有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/04%20%e5%bc%ba%e5%bc%95%e7%94%a8%e3%80%81%e8%bd%af%e5%bc%95%e7%94%a8%e3%80%81%e5%bc%b1%e5%bc%95%e7%94%a8%e3%80%81%e5%b9%bb%e8%b1%a1%e5%bc%95%e7%94%a8%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">04 强引用、软引用、弱引用、幻象引用有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 String、StringBuffer、StringBuilder有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/05%20String%e3%80%81StringBuffer%e3%80%81StringBuilder%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">05 String、StringBuffer、StringBuilder有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 动态代理是基于什么原理？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/06%20%e5%8a%a8%e6%80%81%e4%bb%a3%e7%90%86%e6%98%af%e5%9f%ba%e4%ba%8e%e4%bb%80%e4%b9%88%e5%8e%9f%e7%90%86%ef%bc%9f.md">06 动态代理是基于什么原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 int和Integer有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/07%20int%e5%92%8cInteger%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">07 int和Integer有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 对比Vector、ArrayList、LinkedList有何区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/08%20%e5%af%b9%e6%af%94Vector%e3%80%81ArrayList%e3%80%81LinkedList%e6%9c%89%e4%bd%95%e5%8c%ba%e5%88%ab%ef%bc%9f.md">08 对比Vector、ArrayList、LinkedList有何区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 对比Hashtable、HashMap、TreeMap有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/09%20%e5%af%b9%e6%af%94Hashtable%e3%80%81HashMap%e3%80%81TreeMap%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">09 对比Hashtable、HashMap、TreeMap有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 如何保证集合是线程安全的_ ConcurrentHashMap如何实现高效地线程安全？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/10%20%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e9%9b%86%e5%90%88%e6%98%af%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e7%9a%84_%20ConcurrentHashMap%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e5%9c%b0%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%ef%bc%9f.md">10 如何保证集合是线程安全的_ ConcurrentHashMap如何实现高效地线程安全？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Java提供了哪些IO方式？ NIO如何实现多路复用？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/11%20Java%e6%8f%90%e4%be%9b%e4%ba%86%e5%93%aa%e4%ba%9bIO%e6%96%b9%e5%bc%8f%ef%bc%9f%20NIO%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%a4%9a%e8%b7%af%e5%a4%8d%e7%94%a8%ef%bc%9f.md">11 Java提供了哪些IO方式？ NIO如何实现多路复用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 Java有几种文件拷贝方式？哪一种最高效？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/12%20Java%e6%9c%89%e5%87%a0%e7%a7%8d%e6%96%87%e4%bb%b6%e6%8b%b7%e8%b4%9d%e6%96%b9%e5%bc%8f%ef%bc%9f%e5%93%aa%e4%b8%80%e7%a7%8d%e6%9c%80%e9%ab%98%e6%95%88%ef%bc%9f.md">12 Java有几种文件拷贝方式？哪一种最高效？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 谈谈接口和抽象类有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/13%20%e8%b0%88%e8%b0%88%e6%8e%a5%e5%8f%a3%e5%92%8c%e6%8a%bd%e8%b1%a1%e7%b1%bb%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">13 谈谈接口和抽象类有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 谈谈你知道的设计模式？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/14%20%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9f%a5%e9%81%93%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%9f.md">14 谈谈你知道的设计模式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 synchronized和ReentrantLock有什么区别呢？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/15%20synchronized%e5%92%8cReentrantLock%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%e5%91%a2%ef%bc%9f.md">15 synchronized和ReentrantLock有什么区别呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 synchronized底层如何实现？什么是锁的升级、降级？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/16%20synchronized%e5%ba%95%e5%b1%82%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f%e4%bb%80%e4%b9%88%e6%98%af%e9%94%81%e7%9a%84%e5%8d%87%e7%ba%a7%e3%80%81%e9%99%8d%e7%ba%a7%ef%bc%9f.md">16 synchronized底层如何实现？什么是锁的升级、降级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 一个线程两次调用start()方法会出现什么情况？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/17%20%e4%b8%80%e4%b8%aa%e7%ba%bf%e7%a8%8b%e4%b8%a4%e6%ac%a1%e8%b0%83%e7%94%a8start%28%29%e6%96%b9%e6%b3%95%e4%bc%9a%e5%87%ba%e7%8e%b0%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%ef%bc%9f.md">17 一个线程两次调用start()方法会出现什么情况？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 什么情况下Java程序会产生死锁？如何定位、修复？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/18%20%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8bJava%e7%a8%8b%e5%ba%8f%e4%bc%9a%e4%ba%a7%e7%94%9f%e6%ad%bb%e9%94%81%ef%bc%9f%e5%a6%82%e4%bd%95%e5%ae%9a%e4%bd%8d%e3%80%81%e4%bf%ae%e5%a4%8d%ef%bc%9f.md">18 什么情况下Java程序会产生死锁？如何定位、修复？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Java并发包提供了哪些并发工具类？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/19%20Java%e5%b9%b6%e5%8f%91%e5%8c%85%e6%8f%90%e4%be%9b%e4%ba%86%e5%93%aa%e4%ba%9b%e5%b9%b6%e5%8f%91%e5%b7%a5%e5%85%b7%e7%b1%bb%ef%bc%9f.md">19 Java并发包提供了哪些并发工具类？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 并发包中的ConcurrentLinkedQueue和LinkedBlockingQueue有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/20%20%e5%b9%b6%e5%8f%91%e5%8c%85%e4%b8%ad%e7%9a%84ConcurrentLinkedQueue%e5%92%8cLinkedBlockingQueue%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">20 并发包中的ConcurrentLinkedQueue和LinkedBlockingQueue有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Java并发类库提供的线程池有哪几种？ 分别有什么特点？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/21%20Java%e5%b9%b6%e5%8f%91%e7%b1%bb%e5%ba%93%e6%8f%90%e4%be%9b%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%b1%a0%e6%9c%89%e5%93%aa%e5%87%a0%e7%a7%8d%ef%bc%9f%20%e5%88%86%e5%88%ab%e6%9c%89%e4%bb%80%e4%b9%88%e7%89%b9%e7%82%b9%ef%bc%9f.md">21 Java并发类库提供的线程池有哪几种？ 分别有什么特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 AtomicInteger底层实现原理是什么？如何在自己的产品代码中应用CAS操作？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/22%20AtomicInteger%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e5%a6%82%e4%bd%95%e5%9c%a8%e8%87%aa%e5%b7%b1%e7%9a%84%e4%ba%a7%e5%93%81%e4%bb%a3%e7%a0%81%e4%b8%ad%e5%ba%94%e7%94%a8CAS%e6%93%8d%e4%bd%9c%ef%bc%9f.md">22 AtomicInteger底层实现原理是什么？如何在自己的产品代码中应用CAS操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 请介绍类加载过程，什么是双亲委派模型？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/23%20%e8%af%b7%e4%bb%8b%e7%bb%8d%e7%b1%bb%e5%8a%a0%e8%bd%bd%e8%bf%87%e7%a8%8b%ef%bc%8c%e4%bb%80%e4%b9%88%e6%98%af%e5%8f%8c%e4%ba%b2%e5%a7%94%e6%b4%be%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">23 请介绍类加载过程，什么是双亲委派模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 有哪些方法可以在运行时动态生成一个Java类？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/24%20%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b9%e6%b3%95%e5%8f%af%e4%bb%a5%e5%9c%a8%e8%bf%90%e8%a1%8c%e6%97%b6%e5%8a%a8%e6%80%81%e7%94%9f%e6%88%90%e4%b8%80%e4%b8%aaJava%e7%b1%bb%ef%bc%9f.md">24 有哪些方法可以在运行时动态生成一个Java类？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/25%20%e8%b0%88%e8%b0%88JVM%e5%86%85%e5%ad%98%e5%8c%ba%e5%9f%9f%e7%9a%84%e5%88%92%e5%88%86%ef%bc%8c%e5%93%aa%e4%ba%9b%e5%8c%ba%e5%9f%9f%e5%8f%af%e8%83%bd%e5%8f%91%e7%94%9fOutOfMemoryError_.md">25 谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 如何监控和诊断JVM堆内和堆外内存使用？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/26%20%e5%a6%82%e4%bd%95%e7%9b%91%e6%8e%a7%e5%92%8c%e8%af%8a%e6%96%adJVM%e5%a0%86%e5%86%85%e5%92%8c%e5%a0%86%e5%a4%96%e5%86%85%e5%ad%98%e4%bd%bf%e7%94%a8%ef%bc%9f.md">26 如何监控和诊断JVM堆内和堆外内存使用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 Java常见的垃圾收集器有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/27%20Java%e5%b8%b8%e8%a7%81%e7%9a%84%e5%9e%83%e5%9c%be%e6%94%b6%e9%9b%86%e5%99%a8%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">27 Java常见的垃圾收集器有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 谈谈你的GC调优思路_.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/28%20%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9a%84GC%e8%b0%83%e4%bc%98%e6%80%9d%e8%b7%af_.md">28 谈谈你的GC调优思路_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 Java内存模型中的happen-before是什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/29%20Java%e5%86%85%e5%ad%98%e6%a8%a1%e5%9e%8b%e4%b8%ad%e7%9a%84happen-before%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">29 Java内存模型中的happen-before是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 Java程序运行在Docker等容器环境有哪些新问题？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/30%20Java%e7%a8%8b%e5%ba%8f%e8%bf%90%e8%a1%8c%e5%9c%a8Docker%e7%ad%89%e5%ae%b9%e5%99%a8%e7%8e%af%e5%a2%83%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b0%e9%97%ae%e9%a2%98%ef%bc%9f.md">30 Java程序运行在Docker等容器环境有哪些新问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 你了解Java应用开发中的注入攻击吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/31%20%e4%bd%a0%e4%ba%86%e8%a7%a3Java%e5%ba%94%e7%94%a8%e5%bc%80%e5%8f%91%e4%b8%ad%e7%9a%84%e6%b3%a8%e5%85%a5%e6%94%bb%e5%87%bb%e5%90%97%ef%bc%9f.md">31 你了解Java应用开发中的注入攻击吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 如何写出安全的Java代码？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/32%20%e5%a6%82%e4%bd%95%e5%86%99%e5%87%ba%e5%ae%89%e5%85%a8%e7%9a%84Java%e4%bb%a3%e7%a0%81%ef%bc%9f.md">32 如何写出安全的Java代码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 后台服务出现明显“变慢”，谈谈你的诊断思路？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/33%20%e5%90%8e%e5%8f%b0%e6%9c%8d%e5%8a%a1%e5%87%ba%e7%8e%b0%e6%98%8e%e6%98%be%e2%80%9c%e5%8f%98%e6%85%a2%e2%80%9d%ef%bc%8c%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9a%84%e8%af%8a%e6%96%ad%e6%80%9d%e8%b7%af%ef%bc%9f.md">33 后台服务出现明显“变慢”，谈谈你的诊断思路？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 有人说“Lambda能让Java程序慢30倍”，你怎么看？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/34%20%e6%9c%89%e4%ba%ba%e8%af%b4%e2%80%9cLambda%e8%83%bd%e8%ae%a9Java%e7%a8%8b%e5%ba%8f%e6%85%a230%e5%80%8d%e2%80%9d%ef%bc%8c%e4%bd%a0%e6%80%8e%e4%b9%88%e7%9c%8b%ef%bc%9f.md">34 有人说“Lambda能让Java程序慢30倍”，你怎么看？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 JVM优化Java代码时都做了什么？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/35%20JVM%e4%bc%98%e5%8c%96Java%e4%bb%a3%e7%a0%81%e6%97%b6%e9%83%bd%e5%81%9a%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">35 JVM优化Java代码时都做了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 谈谈MySQL支持的事务隔离级别，以及悲观锁和乐观锁的原理和应用场景？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/36%20%e8%b0%88%e8%b0%88MySQL%e6%94%af%e6%8c%81%e7%9a%84%e4%ba%8b%e5%8a%a1%e9%9a%94%e7%a6%bb%e7%ba%a7%e5%88%ab%ef%bc%8c%e4%bb%a5%e5%8f%8a%e6%82%b2%e8%a7%82%e9%94%81%e5%92%8c%e4%b9%90%e8%a7%82%e9%94%81%e7%9a%84%e5%8e%9f%e7%90%86%e5%92%8c%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%ef%bc%9f.md">36 谈谈MySQL支持的事务隔离级别，以及悲观锁和乐观锁的原理和应用场景？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 谈谈Spring Bean的生命周期和作用域？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/37%20%e8%b0%88%e8%b0%88Spring%20Bean%e7%9a%84%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%e5%92%8c%e4%bd%9c%e7%94%a8%e5%9f%9f%ef%bc%9f.md">37 谈谈Spring Bean的生命周期和作用域？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 对比Java标准NIO类库，你知道Netty是如何实现更高性能的吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/38%20%e5%af%b9%e6%af%94Java%e6%a0%87%e5%87%86NIO%e7%b1%bb%e5%ba%93%ef%bc%8c%e4%bd%a0%e7%9f%a5%e9%81%93Netty%e6%98%af%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%9b%b4%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84%e5%90%97%ef%bc%9f.md">38 对比Java标准NIO类库，你知道Netty是如何实现更高性能的吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 谈谈常用的分布式ID的设计方案？Snowflake是否受冬令时切换影响？.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/39%20%e8%b0%88%e8%b0%88%e5%b8%b8%e7%94%a8%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8fID%e7%9a%84%e8%ae%be%e8%ae%a1%e6%96%b9%e6%a1%88%ef%bc%9fSnowflake%e6%98%af%e5%90%a6%e5%8f%97%e5%86%ac%e4%bb%a4%e6%97%b6%e5%88%87%e6%8d%a2%e5%bd%b1%e5%93%8d%ef%bc%9f.md">39 谈谈常用的分布式ID的设计方案？Snowflake是否受冬令时切换影响？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="周末福利 谈谈我对Java学习和面试的看法.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/%e5%91%a8%e6%9c%ab%e7%a6%8f%e5%88%a9%20%e8%b0%88%e8%b0%88%e6%88%91%e5%af%b9Java%e5%ad%a6%e4%b9%a0%e5%92%8c%e9%9d%a2%e8%af%95%e7%9a%84%e7%9c%8b%e6%b3%95.md">周末福利 谈谈我对Java学习和面试的看法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 技术没有终点.md" href="/%e4%b8%93%e6%a0%8f/Java%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%8a%80%e6%9c%af%e6%b2%a1%e6%9c%89%e7%bb%88%e7%82%b9.md">结束语 技术没有终点.md</a>
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
                            <h1 id="title" data-id="25 谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_" class="title">25 谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError_</h1>
                            <div><p>今天，我将从内存管理的角度，进一步探索Java虚拟机（JVM）。垃圾收集机制为我们打理了很多繁琐的工作，大大提高了开发的效率，但是，垃圾收集也不是万能的，懂得JVM内部的内存结构、工作机制，是设计高扩展性应用和诊断运行时问题的基础，也是Java工程师进阶的必备能力。</p>

<p>今天我要问你的问题是，谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError？</p>

<h2 id="典型回答">典型回答</h2>

<p>通常可以把JVM内存区域分为下面几个方面，其中，有的区域是以线程为单位，而有的区域则是整个JVM进程唯一的。</p>

<p>首先，<strong>程序计数器</strong>（PC，Program Counter Register）。在JVM规范中，每个线程都有它自己的程序计数器，并且任何时间一个线程都只有一个方法在执行，也就是所谓的当前方法。程序计数器会存储当前线程正在执行的Java方法的JVM指令地址；或者，如果是在执行本地方法，则是未指定值（undefined）。</p>

<p>第二，<strong>Java虚拟机栈</strong>（Java Virtual Machine Stack），早期也叫Java栈。每个线程在创建时都会创建一个虚拟机栈，其内部保存一个个的栈帧（Stack Frame），对应着一次次的Java方法调用。</p>

<p>前面谈程序计数器时，提到了当前方法；同理，在一个时间点，对应的只会有一个活动的栈帧，通常叫作当前帧，方法所在的类叫作当前类。如果在该方法中调用了其他方法，对应的新的栈帧会被创建出来，成为新的当前帧，一直到它返回结果或者执行结束。JVM直接对Java栈的操作只有两个，就是对栈帧的压栈和出栈。</p>

<p>栈帧中存储着局部变量表、操作数（operand）栈、动态链接、方法正常退出或者异常退出的定义等。</p>

<p>第三，<strong>堆</strong>（Heap），它是Java内存管理的核心区域，用来放置Java对象实例，几乎所有创建的Java对象实例都是被直接分配在堆上。堆被所有的线程共享，在虚拟机启动时，我们指定的“Xmx”之类参数就是用来指定最大堆空间等指标。</p>

<p>理所当然，堆也是垃圾收集器重点照顾的区域，所以堆内空间还会被不同的垃圾收集器进行进一步的细分，最有名的就是新生代、老年代的划分。</p>

<p>第四，<strong>方法区</strong>（Method Area）。这也是所有线程共享的一块内存区域，用于存储所谓的元（Meta）数据，例如类结构信息，以及对应的运行时常量池、字段、方法代码等。</p>

<p>由于早期的Hotspot JVM实现，很多人习惯于将方法区称为永久代（Permanent Generation）。Oracle JDK 8中将永久代移除，同时增加了元数据区（Metaspace）。</p>

<p>第五，<strong>运行时常量池</strong>（Run-Time Constant Pool），这是方法区的一部分。如果仔细分析过反编译的类文件结构，你能看到版本号、字段、方法、超类、接口等各种信息，还有一项信息就是常量池。Java的常量池可以存放各种常量信息，不管是编译期生成的各种字面量，还是需要在运行时决定的符号引用，所以它比一般语言的符号表存储的信息更加宽泛。</p>

<p>第六，<strong>本地方法栈</strong>（Native Method Stack）。它和Java虚拟机栈是非常相似的，支持对本地方法的调用，也是每个线程都会创建一个。在Oracle Hotspot JVM中，本地方法栈和Java虚拟机栈是在同一块儿区域，这完全取决于技术实现的决定，并未在规范中强制。</p>

<h2 id="考点分析">考点分析</h2>

<p>这是个JVM领域的基础题目，我给出的答案依据的是<a href="https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-2.html#jvms-2.5" target="_blank">JVM规范</a>中运行时数据区定义，这也和大多数书籍和资料解读的角度类似。</p>

<p>JVM内部的概念庞杂，对于初学者比较晦涩，我的建议是在工作之余，还是要去阅读经典书籍，比如我推荐过多次的《深入理解Java虚拟机》。</p>

<p>今天这一讲作为Java虚拟机内存管理的开篇，我会侧重于：</p>

<ul>
<li><p>分析广义上的JVM内存结构或者说Java进程内存结构。</p></li>

<li><p>谈到Java内存模型，不可避免的要涉及OutOfMemory（OOM）问题，那么在Java里面存在哪些种OOM的可能性，分别对应哪个内存区域的异常状况呢？</p></li>
</ul>

<p>注意，具体JVM的内存结构，其实取决于其实现，不同厂商的JVM，或者同一厂商发布的不同版本，都有可能存在一定差异。我在下面的分析中，还会介绍Oracle Hotspot JVM的部分设计变化。</p>

<h2 id="知识扩展">知识扩展</h2>

<p>首先，为了让你有个更加直观、清晰的印象，我画了一个简单的内存结构图，里面展示了我前面提到的堆、线程栈等区域，并从数量上说明了什么是线程私有，例如，程序计数器、Java栈等，以及什么是Java进程唯一。另外，还额外划分出了直接内存等区域。-
<img src="assets/360b8f453e016cb641208a6a8fb589bc.png" alt="" /></p>

<p>这张图反映了实际中Java进程内存占用，与规范中定义的JVM运行时数据区之间的差别，它可以看作是运行时数据区的一个超集。毕竟理论上的视角和现实中的视角是有区别的，规范侧重的是通用的、无差别的部分，而对于应用开发者来说，只要是Java进程在运行时会占用，都会影响到我们的工程实践。</p>

<p>我这里简要介绍两点区别：</p>

<ul>
<li><p>直接内存（Direct Memory）区域，它就是我在[专栏第12讲]中谈到的Direct Buffer所直接分配的内存，也是个容易出现问题的地方。尽管，在JVM工程师的眼中，并不认为它是JVM内部内存的一部分，也并未体现JVM内存模型中。</p></li>

<li><p>JVM本身是个本地程序，还需要其他的内存去完成各种基本任务，比如，JIT Compiler在运行时对热点方法进行编译，就会将编译后的方法储存在Code Cache里面；GC等功能需要运行在本地线程之中，类似部分都需要占用内存空间。这些是实现JVM JIT等功能的需要，但规范中并不涉及。</p></li>
</ul>

<p>如果深入到JVM的实现细节，你会发现一些结论似乎有些模棱两可，比如：</p>

<ul>
<li>Java对象是不是都创建在堆上的呢？</li>
</ul>

<p>我注意到有一些观点，认为通过<a href="https://en.wikipedia.org/wiki/Escape_analysis" target="_blank">逃逸分析</a>，JVM会在栈上分配那些不会逃逸的对象，这在理论上是可行的，但是取决于JVM设计者的选择。据我所知，Oracle Hotspot JVM中并未这么做，这一点在逃逸分析相关的<a href="https://docs.oracle.com/javase/8/docs/technotes/guides/vm/performance-enhancements-7.html#escapeAnalysis" target="_blank">文档</a>里已经说明，所以可以明确所有的对象实例都是创建在堆上。</p>

<ul>
<li>目前很多书籍还是基于JDK 7以前的版本，JDK已经发生了很大变化，Intern字符串的缓存和静态变量曾经都被分配在永久代上，而永久代已经被元数据区取代。但是，Intern字符串缓存和静态变量并不是被转移到元数据区，而是直接在堆上分配，所以这一点同样符合前面一点的结论：对象实例都是分配在堆上。</li>
</ul>

<p>接下来，我们来看看什么是OOM问题，它可能在哪些内存区域发生？</p>

<p>首先，OOM如果通俗点儿说，就是JVM内存不够用了，javadoc中对<a href="https://docs.oracle.com/javase/9/docs/api/java/lang/OutOfMemoryError.html" target="_blank">OutOfMemoryError</a>的解释是，没有空闲内存，并且垃圾收集器也无法提供更多内存。</p>

<p>这里面隐含着一层意思是，在抛出OutOfMemoryError之前，通常垃圾收集器会被触发，尽其所能去清理出空间，例如：</p>

<ul>
<li><p>我在[专栏第4讲]的引用机制分析中，已经提到了JVM会去尝试回收软引用指向的对象等。</p></li>

<li><p>在<a href="http://hg.openjdk.java.net/jdk/jdk/file/9f62267e79df/src/java.base/share/classes/java/nio/Bits.java" target="_blank">java.nio.BIts.reserveMemory()</a> 方法中，我们能清楚的看到，System.gc()会被调用，以清理空间，这也是为什么在大量使用NIO的Direct Buffer之类时，通常建议不要加下面的参数，毕竟是个最后的尝试，有可能避免一定的内存不足问题。</p></li>
</ul>

<pre><code>    -XX:+DisableExplicitGC
</code></pre>

<p>当然，也不是在任何情况下垃圾收集器都会被触发的，比如，我们去分配一个超大对象，类似一个超大数组超过堆的最大值，JVM可以判断出垃圾收集并不能解决这个问题，所以直接抛出OutOfMemoryError。</p>

<p>从我前面分析的数据区的角度，除了程序计数器，其他区域都有可能会因为可能的空间不足发生OutOfMemoryError，简单总结如下：</p>

<ul>
<li><p>堆内存不足是最常见的OOM原因之一，抛出的错误信息是“java.lang.OutOfMemoryError:Java heap space”，原因可能千奇百怪，例如，可能存在内存泄漏问题；也很有可能就是堆的大小不合理，比如我们要处理比较可观的数据量，但是没有显式指定JVM堆大小或者指定数值偏小；或者出现JVM处理引用不及时，导致堆积起来，内存无法释放等。</p></li>

<li><p>而对于Java虚拟机栈和本地方法栈，这里要稍微复杂一点。如果我们写一段程序不断的进行递归调用，而且没有退出条件，就会导致不断地进行压栈。类似这种情况，JVM实际会抛出StackOverFlowError；当然，如果JVM试图去扩展栈空间的的时候失败，则会抛出OutOfMemoryError。</p></li>

<li><p>对于老版本的Oracle JDK，因为永久代的大小是有限的，并且JVM对永久代垃圾回收（如，常量池回收、卸载不再需要的类型）非常不积极，所以当我们不断添加新类型的时候，永久代出现OutOfMemoryError也非常多见，尤其是在运行时存在大量动态类型生成的场合；类似Intern字符串缓存占用太多空间，也会导致OOM问题。对应的异常信息，会标记出来和永久代相关：“java.lang.OutOfMemoryError: PermGen space”。</p></li>

<li><p>随着元数据区的引入，方法区内存已经不再那么窘迫，所以相应的OOM有所改观，出现OOM，异常信息则变成了：“java.lang.OutOfMemoryError: Metaspace”。</p></li>

<li><p>直接内存不足，也会导致OOM，这个已经[专栏第11讲]介绍过。</p></li>
</ul>

<p>今天是JVM内存部分的第一讲，算是我们先进行了热身准备，我介绍了主要的内存区域，以及在不同版本Hotspot JVM内部的变化，并且分析了各区域是否可能产生OutOfMemoryError，以及OOME发生的典型情况。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？今天的思考题是，我在试图分配一个100M bytes大数组的时候发生了OOME，但是GC日志显示，明明堆上还有远不止100M的空间，你觉得可能问题的原因是什么？想要弄清楚这个问题，还需要什么信息呢？</p>

<p>请你在留言区写写你对这个问题的思考，我会选出经过认真思考的留言，送给你一份学习奖励礼券，欢迎你与我一起讨论。</p>

<p>你的朋友是不是也在准备面试呢？你可以“请朋友读”，把今天的题目分享给好友，或许你能帮到他。</p>
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
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f9c71acba7791',t:'MTczNDAyNjM5Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=32&#32;如何写出安全的Java代码？>
        <link rel="icon" href="/static/favicon.png">
        <title>32 如何写出安全的Java代码？ </title>
        
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
                            <h1 id="title" data-id="32 如何写出安全的Java代码？" class="title">32 如何写出安全的Java代码？</h1>
                            <div><p>在上一讲中，我们已经初步接触了Java安全，今天我们将一起探讨更多Java开发中可能影响到安全的场合。很多安全问题，在特定的上下文，存在着不同的定义，尽管本质是相似或一致的，这是由于Java平台自身的特性所带来特有的问题。今天这一讲我将侧重于Java开发者的角度谈代码安全，而不是讲广义的安全风险。</p>

<p>今天我要问你的问题是，如何写出安全的Java代码？</p>

<h2 id="典型回答">典型回答</h2>

<p>这个问题可能有点宽泛，我们可以用特定类型的安全风险为例，如拒绝服务（DoS）攻击，分析Java开发者需要重点考虑的点。</p>

<p>DoS是一种常见的网络攻击，有人也称其为“洪水攻击”。最常见的表现是，利用大量机器发送请求，将目标网站的带宽或者其他资源耗尽，导致其无法响应正常用户的请求。</p>

<p>我认为，从Java语言的角度，更加需要重视的是程序级别的攻击，也就是利用Java、JVM或应用程序的瑕疵，进行低成本的DoS攻击，这也是想要写出安全的Java代码所必须考虑的。例如：</p>

<ul>
<li><p>如果使用的是早期的JDK和Applet等技术，攻击者构建合法但恶劣的程序就相对容易，例如，将其线程优先级设置为最高，做一些看起来无害但空耗资源的事情。幸运的是类似技术已经逐步退出历史舞台，在JDK 9以后，相关模块就已经被移除。</p></li>

<li><p>上一讲中提到的哈希碰撞攻击，就是个典型的例子，对方可以轻易消耗系统有限的CPU和线程资源。从这个角度思考，类似加密、解密、图形处理等计算密集型任务，都要防范被恶意滥用，以免攻击者通过直接调用或者间接触发方式，消耗系统资源。</p></li>

<li><p>利用Java构建类似上传文件或者其他接受输入的服务，需要对消耗系统内存或存储的上限有所控制，因为我们不能将系统安全依赖于用户的合理使用。其中特别注意的是涉及解压缩功能时，就需要防范<a href="https://en.wikipedia.org/wiki/Zip_bomb" target="_blank">Zip bomb</a>等特定攻击。</p></li>

<li><p>另外，Java程序中需要明确释放的资源有很多种，比如文件描述符、数据库连接，甚至是再入锁，任何情况下都应该保证资源释放成功，否则即使平时能够正常运行，也可能被攻击者利用而耗尽某类资源，这也算是可能的DoS攻击来源。</p></li>
</ul>

<p>所以可以看出，实现安全的Java代码，需要从功能设计到实现细节，都充分考虑可能的安全影响。</p>

<h2 id="考点分析">考点分析</h2>

<p>关于今天的问题，以典型的DoS攻击作为切入点，将问题聚焦在Java开发中，我介绍了Java应用设计、实现的注意事项，后面还会介绍更加全面的实践。</p>

<p>其实安全问题实际就是软件的缺陷，软件安全并不存在一劳永逸的秘籍，既离不开设计、架构中的风险分析，也离不开编码、测试等阶段的安全实践手段。对于面试官来说，考察安全问题，除了对特定安全领域知识的考察，更多是要看面试者的Java编程基本功和知识的积累。</p>

<p>所以，我会在后面会循序渐进探讨Java安全编程，这里面没有什么黑科技，只有规范的开发标准，很多安全问题其实是态度问题，取决于你是否真的认真对待它。</p>

<ul>
<li><p>我将以一些典型的代码片段为出发点，分析一些非常容易被忽略的安全风险，并介绍安全问题频发的热点场景，如Java序列化和反序列化。</p></li>

<li><p>从软件生命周期的角度，探讨设计、开发、测试、部署等不同阶段，有哪些常见的安全策略或工具。</p></li>
</ul>

<h2 id="知识扩展">知识扩展</h2>

<p>首先，我们一起来看一段不起眼的条件判断代码，这里可能有什么问题吗？</p>

<pre><code class="language-java">    // a, b, c都是int类型的数值
    if (a + b &lt; c) {
    // …
    }
</code></pre>

<p>你可能会纳闷，这是再常见不过的一个条件判断了，能有什么安全隐患？</p>

<p>这里的隐患是数值类型需要防范溢出，否则这不仅仅可能会带来逻辑错误，在特定情况下可能导致严重的安全漏洞。</p>

<p>从语言特性来说，Java和JVM提供了很多基础性的改进，相比于传统的C、C++等语言，对于数组越界等处理要完善的多，原生的避免了<a href="https://en.wikipedia.org/wiki/Buffer_overflow" target="_blank">缓冲区溢出</a>等攻击方式，提高了软件的安全性。但这并不代表完全杜绝了问题，Java程序可能调用本地代码，也就是JNI技术，错误的数值可能导致C/C++层面的数据越界等问题，这是很危险的。</p>

<p>所以，上面的条件判断，需要判断其数值范围，例如，写成类似下面结构。</p>

<pre><code class="language-java">    if (a &lt; c – b)
</code></pre>

<p>再来看一个例子，请看下面的一段异常处理代码：</p>

<pre><code class="language-java">    try {
    // 业务代码
    } catch (Exception e) {
    throw new RuntimeException(hostname + port + “ doesn’t response”);
    }
</code></pre>

<p>这段代码将敏感信息包含在异常消息中，试想，如果是一个Web应用，异常也没有良好的包装起来，很有可能就把内部信息暴露给终端客户。古人曾经告诫我们“言多必失”是很有道理的，虽然其本意不是指软件安全，但尽量少暴露信息，也是保证安全的基本原则之一。即使我们并不认为某个信息有安全风险，我的建议也是如果没有必要，不要暴露出来。</p>

<p>这种暴露还可能通过其他方式发生，比如某著名的编程技术网站，就被曝光过所有用户名和密码。这些信息都是明文存储，传输过程也未必进行加密，类似这种情况，暴露只是个时间早晚的问题。</p>

<p>对于安全标准特别高的系统，甚至可能要求敏感信息被使用后，要立即明确在内存中销毁，以免被探测；或者避免在发生core dump时，意外暴露。</p>

<p>第三，Java提供了序列化等创新的特性，广泛使用在远程调用等方面，但也带来了复杂的安全问题。直到今天，序列化仍然是个安全问题频发的场景。</p>

<p>针对序列化，通常建议：</p>

<ul>
<li><p>敏感信息不要被序列化！在编码中，建议使用transient关键字将其保护起来。</p></li>

<li><p>反序列化中，建议在readObject中实现与对象构件过程相同的安全检查和数据检查。</p></li>
</ul>

<p>另外，在JDK 9中，Java引入了过滤器机制，以保证反序列化过程中数据都要经过基本验证才可以使用。其原理是通过黑名单和白名单，限定安全或者不安全的类型，并且你可以进行定制，然后通过环境变量灵活进行配置， 更加具体的使用你可以参考 <a href="https://docs.oracle.com/javase/9/docs/api/java/io/ObjectInputFilter.html" target="_blank">ObjectInputFilter</a>。</p>

<p>通过前面的介绍，你可能注意到，很多安全问题都是源于非常基本的编程细节，类似Immutable、封装等设计，都存在着安全性的考虑。从实践的角度，让每个人都了解和掌握这些原则，有必要但并不太现实，有没有什么工程实践手段，可以帮助我们排查安全隐患呢？</p>

<p><strong>开发和测试阶段</strong></p>

<p>在实际开发中，各种功能点五花八门，未必能考虑的全面。我建议没有必要所有都需要自己去从头实现，尽量使用广泛验证过的工具、类库，不管是来自于JDK自身，还是Apache等第三方组织，都在社区的反馈下持续地完善代码安全。</p>

<p>开发过程中应用代码规约标准，是避免安全问题的有效手段。我特别推荐来自孤尽的《阿里巴巴Java开发手册》，以及其配套工具，充分总结了业界在Java等领域的实践经验，将规约实践系统性地引入国内的软件开发，可以有效提高代码质量。</p>

<p>当然，凡事都是有代价的，规约会增加一定的开发成本，可能对迭代的节奏产生一定影响，所以对于不同阶段、不同需求的团队，可以根据自己的情况对规约进行适应性的调整。</p>

<p>落实到实际开发流程中，以OpenJDK团队为例，我们应用了几个不同角度的实践：</p>

<ul>
<li><p>在早期设计阶段，就由安全专家组对新特性进行风险评估。</p></li>

<li><p>开发过程中，尤其是code review阶段，应用OpenJDK自身定制的代码规范。</p></li>

<li><p>利用多种静态分析工具如<a href="http://findbugs.sourceforge.net/" target="_blank">FindBugs</a>、<a href="https://labs.oracle.com/pls/apex/f?p=labs:49:::::P49_PROJECT_ID:13" target="_blank">Parfait</a>等，帮助早期发现潜在安全风险，并对相应问题采取零容忍态度，强制要求解决。</p></li>

<li><p>甚至OpenJDK会默认将任何（编译等）警告，都当作错误对待，并体现在CI流程中。</p></li>

<li><p>在代码check-in等关键环节，利用hook机制去调用规则检查工具，以保证不合规代码不能进入OpenJDK代码库。</p></li>
</ul>

<p>关于静态分析工具的选择，我们选取的原则是“足够好”。没有什么工具能够发现所有问题，所以在保证功能的前提下，影响更大的是分析效率，换句话说是代码分析的噪音高低。不管分析有多么的完备，如果太多误报，就会导致有用信息被噪音覆盖，也不利于后续其他程序化的处理，反倒不利于排查问题。</p>

<p>以上这些是为了保证JDK作为基础平台的苛刻质量要求，在实际产品中，你需要斟酌具体什么程度的要求是合理的。</p>

<p><strong>部署阶段</strong></p>

<p>JDK自身的也是个软件，难免会存在实现瑕疵，我们平时看到JDK更新的安全漏洞补丁，其实就是在修补这些漏洞。我最近还注意到，某大厂后台被曝出了使用的JDK版本存在序列化相关的漏洞。类似这种情况，大多数都是因为使用的JDK是较低版本，算是可以通过部署解决的问题。</p>

<p>如果是安全敏感型产品，建议关注JDK在加解密方面的<a href="https://java.com/en/jre-jdk-cryptoroadmap.html" target="_blank">路线图</a>，同样的标准也应用于其他语言和平台，很多早期认为非常安全的算法，已经被攻破，及时地升级基础软件是安全的必要条件。</p>

<p>攻击和防守是不对称的，只要有一个严重漏洞，对于攻击者就足够了，所以，不能对黑盒形式的部署心存侥幸，这并不能保证系统的安全，攻击者可以利用对软件设计的猜测，结合一系列手段，探测出漏洞。</p>

<p>今天我以DoS等典型攻击方式为例，分析了其在Java平台上的特定表现，并从更多安全编码的细节帮你体会安全问题的普遍性，最后我介绍了软件开发周期中的安全实践，希望能对你的工作有所帮助。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？你在开发中遇到过Java特定的安全问题吗？是怎么解决的呢？</p>

<p>请你在留言区写写你对这个问题的思考，我会选出经过认真思考的留言，送给你一份学习奖励礼券，欢迎你与我一起讨论。</p>

<p>别忘了今晚8点半我会做客“极客Live”，和你一起聊聊Java面试那些事儿。在“极客时间”App内点击“极客Live”即可加入直播，今晚我们不见不散。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#79151515404d4848494e391e14181015571a1614" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f9e089ef47791',t:'MTczNDAyNjQ2MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
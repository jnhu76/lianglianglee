<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=26&#32;如何监控和诊断JVM堆内和堆外内存使用？>
        <link rel="icon" href="/static/favicon.png">
        <title>26 如何监控和诊断JVM堆内和堆外内存使用？ </title>
        
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
                            <h1 id="title" data-id="26 如何监控和诊断JVM堆内和堆外内存使用？" class="title">26 如何监控和诊断JVM堆内和堆外内存使用？</h1>
                            <div><p>上一讲我介绍了JVM内存区域的划分，总结了相关的一些概念，今天我将结合JVM参数、工具等方面，进一步分析JVM内存结构，包括外部资料相对较少的堆外部分。</p>

<p>今天我要问你的问题是，如何监控和诊断JVM堆内和堆外内存使用？</p>

<h2 id="典型回答">典型回答</h2>

<p>了解JVM内存的方法有很多，具体能力范围也有区别，简单总结如下：</p>

<ul>
<li>可以使用综合性的图形化工具，如JConsole、VisualVM（注意，从Oracle JDK 9开始，VisualVM已经不再包含在JDK安装包中）等。这些工具具体使用起来相对比较直观，直接连接到Java进程，然后就可以在图形化界面里掌握内存使用情况。</li>
</ul>

<p>以JConsole为例，其内存页面可以显示常见的<strong>堆内存</strong>和<strong>各种堆外部分</strong>使用状态。</p>

<ul>
<li><p>也可以使用命令行工具进行运行时查询，如jstat和jmap等工具都提供了一些选项，可以查看堆、方法区等使用数据。</p></li>

<li><p>或者，也可以使用jmap等提供的命令，生成堆转储（Heap Dump）文件，然后利用jhat或Eclipse MAT等堆转储分析工具进行详细分析。</p></li>

<li><p>如果你使用的是Tomcat、Weblogic等Java EE服务器，这些服务器同样提供了内存管理相关的功能。</p></li>

<li><p>另外，从某种程度上来说，GC日志等输出，同样包含着丰富的信息。</p></li>
</ul>

<p>这里有一个相对特殊的部分，就是是堆外内存中的直接内存，前面的工具基本不适用，可以使用JDK自带的Native Memory Tracking（NMT）特性，它会从JVM本地内存分配的角度进行解读。</p>

<h2 id="考点分析">考点分析</h2>

<p>今天选取的问题是Java内存管理相关的基础实践，对于普通的内存问题，掌握上面我给出的典型工具和方法就足够了。这个问题也可以理解为考察两个基本方面能力，第一，你是否真的理解了JVM的内部结构；第二，具体到特定内存区域，应该使用什么工具或者特性去定位，可以用什么参数调整。</p>

<p>对于JConsole等工具的使用细节，我在专栏里不再赘述，如果你还没有接触过，你可以参考<a href="https://docs.oracle.com/javase/7/docs/technotes/guides/management/jconsole.html" target="_blank">JConsole官方教程</a>。我这里特别推荐<a href="http://www.oracle.com/technetwork/java/javaseproducts/mission-control/java-mission-control-1998576.html" target="_blank">Java Mission Control</a>（JMC），这是一个非常强大的工具，不仅仅能够使用<a href="https://en.wikipedia.org/wiki/Java_Management_Extensions" target="_blank">JMX</a>进行普通的管理、监控任务，还可以配合<a href="https://docs.oracle.com/javacomponents/jmc-5-4/jfr-runtime-guide/about.htm#JFRUH171" target="_blank">Java Flight Recorder</a>（JFR）技术，以非常低的开销，收集和分析JVM底层的Profiling和事件等信息。目前， Oracle已经将其开源，如果你有兴趣请可以查看OpenJDK的<a href="http://openjdk.java.net/projects/jmc/" target="_blank">Mission Control</a>项目。</p>

<p>关于内存监控与诊断，我会在知识扩展部分结合JVM参数和特性，尽量从庞杂的概念和JVM参数选项中，梳理出相对清晰的框架：</p>

<ul>
<li><p>细化对各部分内存区域的理解，堆内结构是怎样的？如何通过参数调整？</p></li>

<li><p>堆外内存到底包括哪些部分？具体大小受哪些因素影响？</p></li>
</ul>

<h2 id="知识扩展">知识扩展</h2>

<p>今天的分析，我会结合相关JVM参数和工具，进行对比以加深你对内存区域更细粒度的理解。</p>

<p>首先，堆内部是什么结构？</p>

<p>对于堆内存，我在上一讲介绍了最常见的新生代和老年代的划分，其内部结构随着JVM的发展和新GC方式的引入，可以有不同角度的理解，下图就是年代视角的堆结构示意图。-
<img src="assets/721e97abc93449fbdb4c071f7b3b5289.png" alt="" /></p>

<p>你可以看到，按照通常的GC年代方式划分，Java堆内分为：</p>

<ol>
<li>新生代</li>
</ol>

<p>新生代是大部分对象创建和销毁的区域，在通常的Java应用中，绝大部分对象生命周期都是很短暂的。其内部又分为Eden区域，作为对象初始分配的区域；两个Survivor，有时候也叫from、to区域，被用来放置从Minor GC中保留下来的对象。</p>

<ul>
<li><p>JVM会随意选取一个Survivor区域作为“to”，然后会在GC过程中进行区域间拷贝，也就是将Eden中存活下来的对象和from区域的对象，拷贝到这个“to”区域。这种设计主要是为了防止内存的碎片化，并进一步清理无用对象。</p></li>

<li><p>从内存模型而不是垃圾收集的角度，对Eden区域继续进行划分，Hotspot JVM还有一个概念叫做Thread Local Allocation Buffer（TLAB），据我所知所有OpenJDK衍生出来的JVM都提供了TLAB的设计。这是JVM为每个线程分配的一个私有缓存区域，否则，多线程同时分配内存时，为避免操作同一地址，可能需要使用加锁等机制，进而影响分配速度，你可以参考下面的示意图。从图中可以看出，TLAB仍然在堆上，它是分配在Eden区域内的。其内部结构比较直观易懂，start、end就是起始地址，top（指针）则表示已经分配到哪里了。所以我们分配新对象，JVM就会移动top，当top和end相遇时，即表示该缓存已满，JVM会试图再从Eden里分配一块儿。-
<img src="assets/f546839e98ea5d43b595235849b0f2bd.png" alt="" /></p></li>
</ul>

<ol>
<li>老年代</li>
</ol>

<p>放置长生命周期的对象，通常都是从Survivor区域拷贝过来的对象。当然，也有特殊情况，我们知道普通的对象会被分配在TLAB上；如果对象较大，JVM会试图直接分配在Eden其他位置上；如果对象太大，完全无法在新生代找到足够长的连续空闲空间，JVM就会直接分配到老年代。</p>

<ol>
<li>永久代</li>
</ol>

<p>这部分就是早期Hotspot JVM的方法区实现方式了，储存Java类元数据、常量池、Intern字符串缓存，在JDK 8之后就不存在永久代这块儿了。</p>

<p>那么，我们如何利用JVM参数，直接影响堆和内部区域的大小呢？我来简单总结一下：</p>

<ul>
<li><p>最大堆体积</p>

<p>-Xmx value</p></li>

<li><p>初始的最小堆体积</p>

<p>-Xms value</p></li>

<li><p>老年代和新生代的比例</p>

<p>-XX:NewRatio=value</p></li>
</ul>

<p>默认情况下，这个数值是2，意味着老年代是新生代的2倍大；换句话说，新生代是堆大小的1/3。</p>

<ul>
<li><p>当然，也可以不用比例的方式调整新生代的大小，直接指定下面的参数，设定具体的内存大小数值。</p>

<p>-XX:NewSize=value</p></li>

<li><p>Eden和Survivor的大小是按照比例设置的，如果SurvivorRatio是8，那么Survivor区域就是Eden的1/8大小，也就是新生代的1/10，因为YoungGen=Eden + 2*Survivor，JVM参数格式是</p>

<p>-XX:SurvivorRatio=value</p></li>

<li><p>TLAB当然也可以调整，JVM实现了复杂的适应策略，如果你有兴趣可以参考这篇<a href="https://blogs.oracle.com/jonthecollector/the-real-thing" target="_blank">说明</a>。</p></li>
</ul>

<p>不知道你有没有注意到，我在年代视角的堆结构示意图也就是第一张图中，还标记出了Virtual区域，这是块儿什么区域呢？</p>

<p>在JVM内部，如果Xms小于Xmx，堆的大小并不会直接扩展到其上限，也就是说保留的空间（reserved）大于实际能够使用的空间（committed）。当内存需求不断增长的时候，JVM会逐渐扩展新生代等区域的大小，所以Virtual区域代表的就是暂时不可用（uncommitted）的空间。</p>

<p>第二，分析完堆内空间，我们一起来看看JVM堆外内存到底包括什么？</p>

<p>在JMC或JConsole的内存管理界面，会统计部分非堆内存，但提供的信息相对有限，下图就是JMC活动内存池的截图。-
<img src="assets/fa491795ffe21c1f49982de8b7810c2e.png" alt="" /></p>

<p>接下来我会依赖NMT特性对JVM进行分析，它所提供的详细分类信息，非常有助于理解JVM内部实现。</p>

<p>首先来做些准备工作，开启NMT并选择summary模式，</p>

<pre><code>    -XX:NativeMemoryTracking=summary
</code></pre>

<p>为了方便获取和对比NMT输出，选择在应用退出时打印NMT统计信息</p>

<pre><code>    -XX:+UnlockDiagnosticVMOptions -XX:+PrintNMTStatistics
</code></pre>

<p>然后，执行一个简单的在标准输出打印HelloWorld的程序，就可以得到下面的输出-
<img src="assets/55f1c7f0550adbbcc885c97a4dd426bb.png" alt="" /></p>

<p>我来仔细分析一下，NMT所表征的JVM本地内存使用：</p>

<ul>
<li><p>第一部分非常明显是Java堆，我已经分析过使用什么参数调整，不再赘述。</p></li>

<li><p>第二部分是Class内存占用，它所统计的就是Java类元数据所占用的空间，JVM可以通过类似下面的参数调整其大小：</p></li>
</ul>

<pre><code>    -XX:MaxMetaspaceSize=value
</code></pre>

<p>对于本例，因为HelloWorld没有什么用户类库，所以其内存占用主要是启动类加载器（Bootstrap）加载的核心类库。你可以使用下面的小技巧，调整启动类加载器元数据区，这主要是为了对比以加深理解，也许只有在hack JDK时才有实际意义。</p>

<pre><code>    -XX:InitialBootClassLoaderMetaspaceSize=30720
</code></pre>

<ul>
<li>下面是Thread，这里既包括Java线程，如程序主线程、Cleaner线程等，也包括GC等本地线程。你有没有注意到，即使是一个HelloWorld程序，这个线程数量竟然还有25。似乎有很多浪费，设想我们要用Java作为Serverless运行时，每个function是非常短暂的，如何降低线程数量呢？-
如果你充分理解了专栏讲解的内容，对JVM内部有了充分理解，思路就很清晰了：-
JDK 9的默认GC是G1，虽然它在较大堆场景表现良好，但本身就会比传统的Parallel GC或者Serial GC之类复杂太多，所以要么降低其并行线程数目，要么直接切换GC类型；-
JIT编译默认是开启了TieredCompilation的，将其关闭，那么JIT也会变得简单，相应本地线程也会减少。-
我们来对比一下，这是默认参数情况的输出：-</li>
</ul>

<p><img src="assets/97d060b306e44af3a8443f932a0a4d42.png" alt="" /></p>

<p>下面是替换了默认GC，并关闭TieredCompilation的命令行-
<img src="assets/b07d6da56f588cbfadbb7b381346213b.png" alt="" /></p>

<p>得到的统计信息如下，线程数目从25降到了17，消耗的内存也下降了大概1/3。-
<img src="assets/593735623f6917695602095fd249d527.png" alt="" /></p>

<ul>
<li>接下来是Code统计信息，显然这是CodeCache相关内存，也就是JIT compiler存储编译热点方法等信息的地方，JVM提供了一系列参数可以限制其初始值和最大值等，例如：</li>
</ul>

<pre><code>    -XX:InitialCodeCacheSize=value

    -XX:ReservedCodeCacheSize=value
</code></pre>

<p>你可以设置下列JVM参数，也可以只设置其中一个，进一步判断不同参数对CodeCache大小的影响。-
<img src="assets/945740c37433f783d2d877c67dcc1170.png" alt="" />-
<img src="assets/82d1fbc9ca09698c01ccff18fb97c8cd.png" alt="" /></p>

<p>很明显，CodeCache空间下降非常大，这是因为我们关闭了复杂的TieredCompilation，而且还限制了其初始大小。</p>

<ul>
<li>下面就是GC部分了，就像我前面介绍的，G1等垃圾收集器其本身的设施和数据结构就非常复杂和庞大，例如Remembered Set通常都会占用20%~30%的堆空间。如果我把GC明确修改为相对简单的Serial GC，会有什么效果呢？</li>
</ul>

<p>使用命令：</p>

<pre><code>    -XX:+UseSerialGC
</code></pre>

<p><img src="assets/6eeee6624c7dc6be54bfce5e93064233.png" alt="" /></p>

<p>可见，不仅总线程数大大降低（25 → 13），而且GC设施本身的内存开销就少了非常多。据我所知，AWS Lambda中Java运行时就是使用的Serial GC，可以大大降低单个function的启动和运行开销。</p>

<ul>
<li><p>Compiler部分，就是JIT的开销，显然关闭TieredCompilation会降低内存使用。</p></li>

<li><p>其他一些部分占比都非常低，通常也不会出现内存使用问题，请参考<a href="https://docs.oracle.com/javase/8/docs/technotes/guides/troubleshoot/tooldescr022.html#BABCBGFA" target="_blank">官方文档</a>。唯一的例外就是Internal（JDK 11以后在Other部分）部分，其统计信息<strong>包含着Direct Buffer的直接内存</strong>，这其实是堆外内存中比较敏感的部分，很多堆外内存OOM就发生在这里，请参考专栏第12讲的处理步骤。原则上Direct Buffer是不推荐频繁创建或销毁的，如果你怀疑直接内存区域有问题，通常可以通过类似instrument构造函数等手段，排查可能的问题。</p></li>
</ul>

<p>JVM内部结构就介绍到这里，主要目的是为了加深理解，很多方面只有在定制或调优JVM运行时才能真正涉及，随着微服务和Serverless等技术的兴起，JDK确实存在着为新特征的工作负载进行定制的需求。</p>

<p>今天我结合JVM参数和特性，系统地分析了JVM堆内和堆外内存结构，相信你一定对JVM内存结构有了比较深入的了解，在定制Java运行时或者处理OOM等问题的时候，思路也会更加清晰。JVM问题千奇百怪，如果你能快速将问题缩小，大致就能清楚问题可能出在哪里，例如如果定位到问题可能是堆内存泄漏，往往就已经有非常清晰的<a href="https://docs.oracle.com/javase/8/docs/technotes/guides/troubleshoot/memleaks004.html#CIHIEEFH" target="_blank">思路和工具</a>可以去解决了。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？今天的思考题是，如果用程序的方式而不是工具，对Java内存使用进行监控，有哪些技术可以做到?</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#6e020202575a5f5f5e592e09030f0702400d0103" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f9c9dca967791',t:'MTczNDAyNjQwMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
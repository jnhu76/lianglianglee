<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=12&#32;Java有几种文件拷贝方式？哪一种最高效？>
        <link rel="icon" href="/static/favicon.png">
        <title>12 Java有几种文件拷贝方式？哪一种最高效？ </title>
        
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
                            <h1 id="title" data-id="12 Java有几种文件拷贝方式？哪一种最高效？" class="title">12 Java有几种文件拷贝方式？哪一种最高效？</h1>
                            <div><p>我在专栏上一讲提到，NIO不止是多路复用，NIO 2也不只是异步IO，今天我们来看看Java IO体系中，其他不可忽略的部分。</p>

<p>今天我要问你的问题是，Java有几种文件拷贝方式？哪一种最高效？</p>

<h2 id="典型回答">典型回答</h2>

<p>Java有多种比较典型的文件拷贝实现方式，比如：</p>

<p>利用java.io类库，直接为源文件构建一个FileInputStream读取，然后再为目标文件构建一个FileOutputStream，完成写入工作。</p>

<pre><code class="language-java">    public static void copyFileByStream(File source, File dest) throws
            IOException {
        try (InputStream is = new FileInputStream(source);
             OutputStream os = new FileOutputStream(dest);){
            byte[] buffer = new byte[1024];
            int length;
            while ((length = is.read(buffer)) &gt; 0) {
                os.write(buffer, 0, length);
            }
        }
     }
</code></pre>

<p>或者，利用java.nio类库提供的transferTo或transferFrom方法实现。</p>

<pre><code class="language-java">    public static void copyFileByChannel(File source, File dest) throws
            IOException {
        try (FileChannel sourceChannel = new FileInputStream(source)
                .getChannel();
             FileChannel targetChannel = new FileOutputStream(dest).getChannel
                     ();){
            for (long count = sourceChannel.size() ;count&gt;0 ;) {
                long transferred = sourceChannel.transferTo(
                        sourceChannel.position(), count, targetChannel);            sourceChannel.position(sourceChannel.position() + transferred);
                count -= transferred;
            }
        }
     }
</code></pre>

<p>当然，Java标准类库本身已经提供了几种Files.copy的实现。</p>

<p>对于Copy的效率，这个其实与操作系统和配置等情况相关，总体上来说，NIO transferTo/From的方式<strong>可能更快</strong>，因为它更能利用现代操作系统底层机制，避免不必要拷贝和上下文切换。</p>

<h2 id="考点分析">考点分析</h2>

<p>今天这个问题，从面试的角度来看，确实是一个面试考察的点，针对我上面的典型回答，面试官还可能会从实践角度，或者IO底层实现机制等方面进一步提问。这一讲的内容从面试题出发，主要还是为了让你进一步加深对Java IO类库设计和实现的了解。</p>

<p>从实践角度，我前面并没有明确说NIO transfer的方案一定最快，真实情况也确实未必如此。我们可以根据理论分析给出可行的推断，保持合理的怀疑，给出验证结论的思路，有时候面试官考察的就是如何将猜测变成可验证的结论，思考方式远比记住结论重要。</p>

<p>从技术角度展开，下面这些方面值得注意：</p>

<ul>
<li><p>不同的copy方式，底层机制有什么区别？</p></li>

<li><p>为什么零拷贝（zero-copy）可能有性能优势？</p></li>

<li><p>Buffer分类与使用。</p></li>

<li><p>Direct Buffer对垃圾收集等方面的影响与实践选择。</p></li>
</ul>

<p>接下来，我们一起来分析一下吧。</p>

<h2 id="知识扩展">知识扩展</h2>

<ol>
<li>拷贝实现机制分析</li>
</ol>

<p>先来理解一下，前面实现的不同拷贝方法，本质上有什么明显的区别。</p>

<p>首先，你需要理解用户态空间（User Space）和内核态空间（Kernel Space），这是操作系统层面的基本概念，操作系统内核、硬件驱动等运行在内核态空间，具有相对高的特权；而用户态空间，则是给普通应用和服务使用。你可以参考：<a href="https://en.wikipedia.org/wiki/User_space" target="_blank">https://en.wikipedia.org/wiki/User_space</a>。</p>

<p>当我们使用输入输出流进行读写时，实际上是进行了多次上下文切换，比如应用读取数据时，先在内核态将数据从磁盘读取到内核缓存，再切换到用户态将数据从内核缓存读取到用户缓存。</p>

<p>写入操作也是类似，仅仅是步骤相反，你可以参考下面这张图。</p>

<p><img src="assets/6d2368424431f1b0d2b935386324b585.png" alt="" /></p>

<p>所以，这种方式会带来一定的额外开销，可能会降低IO效率。</p>

<p>而基于NIO transferTo的实现方式，在Linux和Unix上，则会使用到零拷贝技术，数据传输并不需要用户态参与，省去了上下文切换的开销和不必要的内存拷贝，进而可能提高应用拷贝性能。注意，transferTo不仅仅是可以用在文件拷贝中，与其类似的，例如读取磁盘文件，然后进行Socket发送，同样可以享受这种机制带来的性能和扩展性提高。</p>

<p>transferTo的传输过程是：</p>

<p><img src="assets/b0c8226992bb97adda5ad84fe25372ea.png" alt="" /></p>

<ol>
<li>Java IO/NIO源码结构</li>
</ol>

<p>前面我在典型回答中提了第三种方式，即Java标准库也提供了文件拷贝方法（java.nio.file.Files.copy）。如果你这样回答，就一定要小心了，因为很少有问题的答案是仅仅调用某个方法。从面试的角度，面试官往往会追问：既然你提到了标准库，那么它是怎么实现的呢？有的公司面试官以喜欢追问而出名，直到追问到你说不知道。</p>

<p>其实，这个问题的答案还真不是那么直观，因为实际上有几个不同的copy方法。</p>

<pre><code class="language-java">    public static Path copy(Path source, Path target, CopyOption... options)
        throws IOException
    

    public static long copy(InputStream in, Path target, CopyOption... options)
        throws IOException
    

    public static long copy(Path source, OutputStream out) 
    throws IOException
</code></pre>

<p>可以看到，copy不仅仅是支持文件之间操作，没有人限定输入输出流一定是针对文件的，这是两个很实用的工具方法。</p>

<p>后面两种copy实现，能够在方法实现里直接看到使用的是InputStream.transferTo()，你可以直接看源码，其内部实现其实是stream在用户态的读写；而对于第一种方法的分析过程要相对麻烦一些，可以参考下面片段。简单起见，我只分析同类型文件系统拷贝过程。</p>

<pre><code class="language-java">    public static Path copy(Path source, Path target, CopyOption... options)
        throws IOException
     {
        FileSystemProvider provider = provider(source);
        if (provider(target) == provider) {
            // same provider
            provider.copy(source, target, options);//这是本文分析的路径
        } else {
            // different providers
            CopyMoveHelper.copyToForeignTarget(source, target, options);
        }
        return target;
    }
</code></pre>

<p>我把源码分析过程简单记录如下，JDK的源代码中，内部实现和公共API定义也不是可以能够简单关联上的，NIO部分代码甚至是定义为模板而不是Java源文件，在build过程自动生成源码，下面顺便介绍一下部分JDK代码机制和如何绕过隐藏障碍。</p>

<ul>
<li><p>首先，直接跟踪，发现FileSystemProvider只是个抽象类，阅读它的<a href="http://hg.openjdk.java.net/jdk/jdk/file/f84ae8aa5d88/src/java.base/share/classes/java/nio/file/spi/FileSystemProvider.java" target="_blank">源码</a>能够理解到，原来文件系统实际逻辑存在于JDK内部实现里，公共API其实是通过ServiceLoader机制加载一系列文件系统实现，然后提供服务。</p></li>

<li><p>我们可以在JDK源码里搜索FileSystemProvider和nio，可以定位到<a href="http://hg.openjdk.java.net/jdk/jdk/file/f84ae8aa5d88/src/java.base/share/classes/sun/nio/fs" target="_blank">sun/nio/fs</a>，我们知道NIO底层是和操作系统紧密相关的，所以每个平台都有自己的部分特有文件系统逻辑。</p></li>
</ul>

<p><img src="assets/5e0bf3130dffa8e56f398f0856eb76f7.png" alt="" /></p>

<ul>
<li><p>省略掉一些细节，最后我们一步步定位到UnixFileSystemProvider → UnixCopyFile.Transfer，发现这是个本地方法。</p></li>

<li><p>最后，明确定位到<a href="http://hg.openjdk.java.net/jdk/jdk/file/f84ae8aa5d88/src/java.base/unix/native/libnio/fs/UnixCopyFile.c" target="_blank">UnixCopyFile.c</a>，其内部实现清楚说明竟然只是简单的用户态空间拷贝！</p></li>
</ul>

<p>所以，我们明确这个最常见的copy方法其实不是利用transferTo，而是本地技术实现的用户态拷贝。</p>

<p>前面谈了不少机制和源码，我简单从实践角度总结一下，如何提高类似拷贝等IO操作的性能，有一些宽泛的原则：</p>

<ul>
<li><p>在程序中，使用缓存等机制，合理减少IO次数（在网络通信中，如TCP传输，window大小也可以看作是类似思路）。</p></li>

<li><p>使用transferTo等机制，减少上下文切换和额外IO操作。</p></li>

<li><p>尽量减少不必要的转换过程，比如编解码；对象序列化和反序列化，比如操作文本文件或者网络通信，如果不是过程中需要使用文本信息，可以考虑不要将二进制信息转换成字符串，直接传输二进制信息。</p></li>
</ul>

<ol>
<li>掌握NIO Buffer</li>
</ol>

<p>我在上一讲提到Buffer是NIO操作数据的基本工具，Java为每种原始数据类型都提供了相应的Buffer实现（布尔除外），所以掌握和使用Buffer是十分必要的，尤其是涉及Direct Buffer等使用，因为其在垃圾收集等方面的特殊性，更要重点掌握。</p>

<p><img src="assets/5220029e92bc21e99920937a8210276e.png" alt="" /></p>

<p>Buffer有几个基本属性：</p>

<ul>
<li><p>capacity，它反映这个Buffer到底有多大，也就是数组的长度。</p></li>

<li><p>position，要操作的数据起始位置。</p></li>

<li><p>limit，相当于操作的限额。在读取或者写入时，limit的意义很明显是不一样的。比如，读取操作时，很可能将limit设置到所容纳数据的上限；而在写入时，则会设置容量或容量以下的可写限度。</p></li>

<li><p>mark，记录上一次postion的位置，默认是0，算是一个便利性的考虑，往往不是必须的。</p></li>
</ul>

<p>前面三个是我们日常使用最频繁的，我简单梳理下Buffer的基本操作：</p>

<ul>
<li><p>我们创建了一个ByteBuffer，准备放入数据，capacity当然就是缓冲区大小，而position就是0，limit默认就是capacity的大小。</p></li>

<li><p>当我们写入几个字节的数据时，position就会跟着水涨船高，但是它不可能超过limit的大小。</p></li>

<li><p>如果我们想把前面写入的数据读出来，需要调用flip方法，将position设置为0，limit设置为以前的position那里。</p></li>

<li><p>如果还想从头再读一遍，可以调用rewind，让limit不变，position再次设置为0。</p></li>
</ul>

<p>更进一步的详细使用，我建议参考相关<a href="http://tutorials.jenkov.com/java-nio/buffers.html" target="_blank">教程</a>。</p>

<ol>
<li>Direct Buffer和垃圾收集</li>
</ol>

<p>我这里重点介绍两种特别的Buffer。</p>

<ul>
<li><p>Direct Buffer：如果我们看Buffer的方法定义，你会发现它定义了isDirect()方法，返回当前Buffer是否是Direct类型。这是因为Java提供了堆内和堆外（Direct）Buffer，我们可以以它的allocate或者allocateDirect方法直接创建。</p></li>

<li><p>MappedByteBuffer：它将文件按照指定大小直接映射为内存区域，当程序访问这个内存区域时将直接操作这块儿文件数据，省去了将数据从内核空间向用户空间传输的损耗。我们可以使用<a href="https://docs.oracle.com/javase/9/docs/api/java/nio/channels/FileChannel.html#map-java.nio.channels.FileChannel.MapMode-long-long-" target="_blank">FileChannel.map</a>创建MappedByteBuffer，它本质上也是种Direct Buffer。</p></li>
</ul>

<p>在实际使用中，Java会尽量对Direct Buffer仅做本地IO操作，对于很多大数据量的IO密集操作，可能会带来非常大的性能优势，因为：</p>

<ul>
<li><p>Direct Buffer生命周期内内存地址都不会再发生更改，进而内核可以安全地对其进行访问，很多IO操作会很高效。</p></li>

<li><p>减少了堆内对象存储的可能额外维护工作，所以访问效率可能有所提高。</p></li>
</ul>

<p>但是请注意，Direct Buffer创建和销毁过程中，都会比一般的堆内Buffer增加部分开销，所以通常都建议用于长期使用、数据较大的场景。</p>

<p>使用Direct Buffer，我们需要清楚它对内存和JVM参数的影响。首先，因为它不在堆上，所以Xmx之类参数，其实并不能影响Direct Buffer等堆外成员所使用的内存额度，我们可以使用下面参数设置大小：</p>

<pre><code>-XX:MaxDirectMemorySize=512M
</code></pre>

<p>从参数设置和内存问题排查角度来看，这意味着我们在计算Java可以使用的内存大小的时候，不能只考虑堆的需要，还有Direct Buffer等一系列堆外因素。如果出现内存不足，堆外内存占用也是一种可能性。</p>

<p>另外，大多数垃圾收集过程中，都不会主动收集Direct Buffer，它的垃圾收集过程，就是基于我在专栏前面所介绍的Cleaner（一个内部实现）和幻象引用（PhantomReference）机制，其本身不是public类型，内部实现了一个Deallocator负责销毁的逻辑。对它的销毁往往要拖到full GC的时候，所以使用不当很容易导致OutOfMemoryError。</p>

<p>对于Direct Buffer的回收，我有几个建议：</p>

<ul>
<li><p>在应用程序中，显式地调用System.gc()来强制触发。</p></li>

<li><p>另外一种思路是，在大量使用Direct Buffer的部分框架中，框架会自己在程序中调用释放方法，Netty就是这么做的，有兴趣可以参考其实现（PlatformDependent0）。</p></li>

<li><p>重复使用Direct Buffer。</p></li>
</ul>

<ol>
<li>跟踪和诊断Direct Buffer内存占用？</li>
</ol>

<p>因为通常的垃圾收集日志等记录，并不包含Direct Buffer等信息，所以Direct Buffer内存诊断也是个比较头疼的事情。幸好，在JDK 8之后的版本，我们可以方便地使用Native Memory Tracking（NMT）特性来进行诊断，你可以在程序启动时加上下面参数：</p>

<pre><code>    -XX:NativeMemoryTracking={summary|detail}
</code></pre>

<p>注意，激活NMT通常都会导致JVM出现5%~10%的性能下降，请谨慎考虑。</p>

<p>运行时，可以采用下面命令进行交互式对比：</p>

<pre><code>    // 打印NMT信息
    jcmd &lt;pid&gt; VM.native_memory detail 
    
    // 进行baseline，以对比分配内存变化
    jcmd &lt;pid&gt; VM.native_memory baseline
    
    // 进行baseline，以对比分配内存变化
    jcmd &lt;pid&gt; VM.native_memory detail.diff
</code></pre>

<p>我们可以在Internal部分发现Direct Buffer内存使用的信息，这是因为其底层实际是利用unsafe_allocatememory。严格说，这不是JVM内部使用的内存，所以在JDK 11以后，其实它是归类在other部分里。</p>

<p>JDK 9的输出片段如下，“+”表示的就是diff命令发现的分配变化：</p>

<pre><code>    -Internal (reserved=679KB +4KB, committed=679KB +4KB)
                  (malloc=615KB +4KB #1571 +4)
                  (mmap: reserved=64KB, committed=64KB)
</code></pre>

<p><strong>注意</strong>：JVM的堆外内存远不止Direct Buffer，NMT输出的信息当然也远不止这些，我在专栏后面有综合分析更加具体的内存结构的主题。</p>

<p>今天我分析了Java IO/NIO底层文件操作数据的机制，以及如何实现零拷贝的高性能操作，梳理了Buffer的使用和类型，并针对Direct Buffer的生命周期管理和诊断进行了较详细的分析。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？你可以思考下，如果我们需要在channel读取的过程中，将不同片段写入到相应的Buffer里面（类似二进制消息分拆成消息头、消息体等），可以采用NIO的什么机制做到呢？</p>

<p>请你在留言区写写你对这个问题的思考，我会选出经过认真思考的留言，送给你一份学习鼓励金，欢迎你与我一起讨论。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#2d41414114191c1c1d1a6d4a404c4441034e4240" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f9a16db507791',t:'MTczNDAyNjI5OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=05&#32;String、StringBuffer、StringBuilder有什么区别？>
        <link rel="icon" href="/static/favicon.png">
        <title>05 String、StringBuffer、StringBuilder有什么区别？ </title>
        
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
                            <h1 id="title" data-id="05 String、StringBuffer、StringBuilder有什么区别？" class="title">05 String、StringBuffer、StringBuilder有什么区别？</h1>
                            <div><p>今天我会聊聊日常使用的字符串，别看它似乎很简单，但其实字符串几乎在所有编程语言里都是个特殊的存在，因为不管是数量还是体积，字符串都是大多数应用中的重要组成。</p>

<p>今天我要问你的问题是，理解Java的字符串，String、StringBuffer、StringBuilder有什么区别？</p>

<h2 id="典型回答">典型回答</h2>

<p>String是Java语言非常基础和重要的类，提供了构造和管理字符串的各种基本逻辑。它是典型的Immutable类，被声明成为final class，所有属性也都是final的。也由于它的不可变性，类似拼接、裁剪字符串等动作，都会产生新的String对象。由于字符串操作的普遍性，所以相关操作的效率往往对应用性能有明显影响。</p>

<p>StringBuffer是为解决上面提到拼接产生太多中间对象的问题而提供的一个类，我们可以用append或者add方法，把字符串添加到已有序列的末尾或者指定位置。StringBuffer本质是一个线程安全的可修改字符序列，它保证了线程安全，也随之带来了额外的性能开销，所以除非有线程安全的需要，不然还是推荐使用它的后继者，也就是StringBuilder。</p>

<p>StringBuilder是Java 1.5中新增的，在能力上和StringBuffer没有本质区别，但是它去掉了线程安全的部分，有效减小了开销，是绝大部分情况下进行字符串拼接的首选。</p>

<h2 id="考点分析">考点分析</h2>

<p>几乎所有的应用开发都离不开操作字符串，理解字符串的设计和实现以及相关工具如拼接类的使用，对写出高质量代码是非常有帮助的。关于这个问题，我前面的回答是一个通常的概要性回答，至少你要知道String是Immutable的，字符串操作不当可能会产生大量临时字符串，以及线程安全方面的区别。</p>

<p>如果继续深入，面试官可以从各种不同的角度考察，比如可以：</p>

<ul>
<li><p>通过String和相关类，考察基本的线程安全设计与实现，各种基础编程实践。</p></li>

<li><p>考察JVM对象缓存机制的理解以及如何良好地使用。</p></li>

<li><p>考察JVM优化Java代码的一些技巧。</p></li>

<li><p>String相关类的演进，比如Java 9中实现的巨大变化。</p></li>

<li><p>…</p></li>
</ul>

<p>针对上面这几方面，我会在知识扩展部分与你详细聊聊。</p>

<h2 id="知识扩展">知识扩展</h2>

<ol>
<li>字符串设计和实现考量</li>
</ol>

<p>我在前面介绍过，String是Immutable类的典型实现，原生的保证了基础线程安全，因为你无法对它内部数据进行任何修改，这种便利甚至体现在拷贝构造函数中，由于不可变，Immutable对象在拷贝时不需要额外复制数据。</p>

<p>我们再来看看StringBuffer实现的一些细节，它的线程安全是通过把各种修改数据的方法都加上synchronized关键字实现的，非常直白。其实，这种简单粗暴的实现方式，非常适合我们常见的线程安全类实现，不必纠结于synchronized性能之类的，有人说“过早优化是万恶之源”，考虑可靠性、正确性和代码可读性才是大多数应用开发最重要的因素。</p>

<p>为了实现修改字符序列的目的，StringBuffer和StringBuilder底层都是利用可修改的（char，JDK 9以后是byte）数组，二者都继承了AbstractStringBuilder，里面包含了基本操作，区别仅在于最终的方法是否加了synchronized。</p>

<p>另外，这个内部数组应该创建成多大的呢？如果太小，拼接的时候可能要重新创建足够大的数组；如果太大，又会浪费空间。目前的实现是，构建时初始字符串长度加16（这意味着，如果没有构建对象时输入最初的字符串，那么初始值就是16）。我们如果确定拼接会发生非常多次，而且大概是可预计的，那么就可以指定合适的大小，避免很多次扩容的开销。扩容会产生多重开销，因为要抛弃原有数组，创建新的（可以简单认为是倍数）数组，还要进行arraycopy。</p>

<p>前面我讲的这些内容，在具体的代码书写中，应该如何选择呢？</p>

<p>在没有线程安全问题的情况下，全部拼接操作是应该都用StringBuilder实现吗？毕竟这样书写的代码，还是要多敲很多字的，可读性也不理想，下面的对比非常明显。</p>

<pre><code class="language-java">    String strByBuilder  = new
    StringBuilder().append(&quot;aa&quot;).append(&quot;bb&quot;).append(&quot;cc&quot;).append
                (&quot;dd&quot;).toString();
                 
    String strByConcat = &quot;aa&quot; + &quot;bb&quot; + &quot;cc&quot; + &quot;dd&quot;;
</code></pre>

<p>其实，在通常情况下，没有必要过于担心，要相信Java还是非常智能的。</p>

<p>我们来做个实验，把下面一段代码，利用不同版本的JDK编译，然后再反编译，例如：</p>

<pre><code class="language-java">    public class StringConcat {
         public static String concat(String str) {
           return str + “aa” + “bb”;
         }
    }
</code></pre>

<p>先编译再反编译，比如使用不同版本的JDK：</p>

<pre><code>    ${JAVA_HOME}/bin/javac StringConcat.java
    ${JAVA_HOME}/bin/javap -v StringConcat.class
</code></pre>

<p>JDK 8的输出片段是：</p>

<pre><code>             0: new           #2                  // class java/lang/StringBuilder
             3: dup
             4: invokespecial #3                  // Method java/lang/StringBuilder.&quot;&lt;init&gt;&quot;:()V
             7: aload_0
             8: invokevirtual #4                  // Method java/lang/StringBuilder.append:(Ljava/lang/String;)Ljava/lang/StringBuilder;
            11: ldc           #5                  // String aa
            13: invokevirtual #4                  // Method java/lang/StringBuilder.append:(Ljava/lang/String;)Ljava/lang/StringBuilder;
            16: ldc           #6                  // String bb
            18: invokevirtual #4                  // Method java/lang/StringBuilder.append:(Ljava/lang/String;)Ljava/lang/StringBuilder;
            21: invokevirtual #7                  // Method java/lang/StringBuilder.toString:()Ljava/lang/String;
</code></pre>

<p>而在JDK 9中，反编译的结果就会有点特别了，片段是：</p>

<pre><code>             // concat method
             1: invokedynamic #2,  0              // InvokeDynamic #0:makeConcatWithConstants:(Ljava/lang/String;)Ljava/lang/String;
             
             // ...
             // 实际是利用了MethodHandle,统一了入口
             0: #15 REF_invokeStatic java/lang/invoke/StringConcatFactory.makeConcatWithConstants:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/String;[Ljava/lang/Object;)Ljava/lang/invoke/CallSite;
</code></pre>

<p>你可以看到，非静态的拼接逻辑在JDK 8中会自动被javac转换为StringBuilder操作；而在JDK 9里面，则是体现了思路的变化。Java 9利用InvokeDynamic，将字符串拼接的优化与javac生成的字节码解耦，假设未来JVM增强相关运行时实现，将不需要依赖javac的任何修改。</p>

<p>在日常编程中，保证程序的可读性、可维护性，往往比所谓的最优性能更重要，你可以根据实际需求酌情选择具体的编码方式。</p>

<ol>
<li>字符串缓存</li>
</ol>

<p>我们粗略统计过，把常见应用进行堆转储（Dump Heap），然后分析对象组成，会发现平均25%的对象是字符串，并且其中约半数是重复的。如果能避免创建重复字符串，可以有效降低内存消耗和对象创建开销。</p>

<p>String在Java 6以后提供了intern()方法，目的是提示JVM把相应字符串缓存起来，以备重复使用。在我们创建字符串对象并调用intern()方法的时候，如果已经有缓存的字符串，就会返回缓存里的实例，否则将其缓存起来。一般来说，JVM会将所有的类似“abc”这样的文本字符串，或者字符串常量之类缓存起来。</p>

<p>看起来很不错是吧？但实际情况估计会让你大跌眼镜。一般使用Java 6这种历史版本，并不推荐大量使用intern，为什么呢？魔鬼存在于细节中，被缓存的字符串是存在所谓PermGen里的，也就是臭名昭著的“永久代”，这个空间是很有限的，也基本不会被FullGC之外的垃圾收集照顾到。所以，如果使用不当，OOM就会光顾。</p>

<p>在后续版本中，这个缓存被放置在堆中，这样就极大避免了永久代占满的问题，甚至永久代在JDK 8中被MetaSpace（元数据区）替代了。而且，默认缓存大小也在不断地扩大中，从最初的1009，到7u40以后被修改为60013。你可以使用下面的参数直接打印具体数字，可以拿自己的JDK立刻试验一下。</p>

<pre><code>    -XX:+PrintStringTableStatistics
</code></pre>

<p>你也可以使用下面的JVM参数手动调整大小，但是绝大部分情况下并不需要调整，除非你确定它的大小已经影响了操作效率。</p>

<pre><code>    -XX:StringTableSize=N
</code></pre>

<p>Intern是一种<strong>显式地排重机制</strong>，但是它也有一定的副作用，因为需要开发者写代码时明确调用，一是不方便，每一个都显式调用是非常麻烦的；另外就是我们很难保证效率，应用开发阶段很难清楚地预计字符串的重复情况，有人认为这是一种污染代码的实践。</p>

<p>幸好在Oracle JDK 8u20之后，推出了一个新的特性，也就是G1 GC下的字符串排重。它是通过将相同数据的字符串指向同一份数据来做到的，是JVM底层的改变，并不需要Java类库做什么修改。</p>

<p>注意这个功能目前是默认关闭的，你需要使用下面参数开启，并且记得指定使用G1 GC：</p>

<pre><code>    -XX:+UseStringDeduplication
</code></pre>

<p>前面说到的几个方面，只是Java底层对字符串各种优化的一角，在运行时，字符串的一些基础操作会直接利用JVM内部的Intrinsic机制，往往运行的就是特殊优化的本地代码，而根本就不是Java代码生成的字节码。Intrinsic可以简单理解为，是一种利用native方式hard-coded的逻辑，算是一种特别的内联，很多优化还是需要直接使用特定的CPU指令，具体可以看相关<a href="http://hg.openjdk.java.net/jdk/jdk/file/44b64fc0baa3/src/hotspot/share/classfile/vmSymbols.hpp" target="_blank">源码</a>，搜索“string”以查找相关Intrinsic定义。当然，你也可以在启动实验应用时，使用下面参数，了解intrinsic发生的状态。</p>

<pre><code>    -XX:+PrintCompilation -XX:+UnlockDiagnosticVMOptions -XX:+PrintInlining
        //样例输出片段    
            180    3       3       java.lang.String::charAt (25 bytes)  
                                      @ 1   java.lang.String::isLatin1 (19 bytes)   
                                      ...  
                                      @ 7 java.lang.StringUTF16::getChar (60 bytes) intrinsic 
</code></pre>

<p>可以看出，仅仅是字符串一个实现，就需要Java平台工程师和科学家付出如此大且默默无闻的努力，我们得到的很多便利都是来源于此。</p>

<p>我会在专栏后面的JVM和性能等主题，详细介绍JVM内部优化的一些方法，如果你有兴趣可以再深入学习。即使你不做JVM开发或者暂时还没有使用到特别的性能优化，这些知识也能帮助你增加技术深度。</p>

<ol>
<li>String自身的演化</li>
</ol>

<p>如果你仔细观察过Java的字符串，在历史版本中，它是使用char数组来存数据的，这样非常直接。但是Java中的char是两个bytes大小，拉丁语系语言的字符，根本就不需要太宽的char，这样无区别的实现就造成了一定的浪费。密度是编程语言平台永恒的话题，因为归根结底绝大部分任务是要来操作数据的。</p>

<p>其实在Java 6的时候，Oracle JDK就提供了压缩字符串的特性，但是这个特性的实现并不是开源的，而且在实践中也暴露出了一些问题，所以在最新的JDK版本中已经将它移除了。</p>

<p>在Java 9中，我们引入了Compact Strings的设计，对字符串进行了大刀阔斧的改进。将数据存储方式从char数组，改变为一个byte数组加上一个标识编码的所谓coder，并且将相关字符串操作类都进行了修改。另外，所有相关的Intrinsic之类也都进行了重写，以保证没有任何性能损失。</p>

<p>虽然底层实现发生了这么大的改变，但是Java字符串的行为并没有任何大的变化，所以这个特性对于绝大部分应用来说是透明的，绝大部分情况不需要修改已有代码。</p>

<p>当然，在极端情况下，字符串也出现了一些能力退化，比如最大字符串的大小。你可以思考下，原来char数组的实现，字符串的最大长度就是数组本身的长度限制，但是替换成byte数组，同样数组长度下，存储能力是退化了一倍的！还好这是存在于理论中的极限，还没有发现现实应用受此影响。</p>

<p>在通用的性能测试和产品实验中，我们能非常明显地看到紧凑字符串带来的优势，<strong>即更小的内存占用、更快的操作速度</strong>。</p>

<p>今天我从String、StringBuffer和StringBuilder的主要设计和实现特点开始，分析了字符串缓存的intern机制、非代码侵入性的虚拟机层面排重、Java 9中紧凑字符的改进，并且初步接触了JVM的底层优化机制intrinsic。从实践的角度，不管是Compact Strings还是底层intrinsic优化，都说明了使用Java基础类库的优势，它们往往能够得到最大程度、最高质量的优化，而且只要升级JDK版本，就能零成本地享受这些益处。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？限于篇幅有限，还有很多字符相关的问题没有来得及讨论，比如编码相关的问题。可以思考一下，很多字符串操作，比如getBytes()/<a href="https://docs.oracle.com/javase/9/docs/api/java/lang/String.html#String-byte:A-" target="_blank">String</a>(byte[] bytes)等都是隐含着使用平台默认编码，这是一种好的实践吗？是否有利于避免乱码？</p>

<p>请你在留言区写写你对这个问题的思考，或者分享一下你在操作字符串时掉过的坑，我会选出经过认真思考的留言，送给你一份学习鼓励金，欢迎你与我一起讨论。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#08646464313c3939383f486f65696164266b6765" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0f98e38a997791',t:'MTczNDAyNjI1MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
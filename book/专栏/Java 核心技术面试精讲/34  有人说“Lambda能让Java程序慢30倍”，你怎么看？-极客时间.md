<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=34&#32;&#32;有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间>
        <link rel="icon" href="/static/favicon.png">
        <title>34  有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间 </title>
        
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
                        <a class="menu-item" id="00 开篇词  以面试题为切入点，有效提升你的Java内功-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%bb%a5%e9%9d%a2%e8%af%95%e9%a2%98%e4%b8%ba%e5%88%87%e5%85%a5%e7%82%b9%ef%bc%8c%e6%9c%89%e6%95%88%e6%8f%90%e5%8d%87%e4%bd%a0%e7%9a%84Java%e5%86%85%e5%8a%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">00 开篇词  以面试题为切入点，有效提升你的Java内功-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  谈谈你对Java平台的理解？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/01%20%20%e8%b0%88%e8%b0%88%e4%bd%a0%e5%af%b9Java%e5%b9%b3%e5%8f%b0%e7%9a%84%e7%90%86%e8%a7%a3%ef%bc%9f.md">01  谈谈你对Java平台的理解？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  Exception和Error有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/02%20%20Exception%e5%92%8cError%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">02  Exception和Error有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  谈谈final、finally、 finalize有什么不同？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/03%20%20%e8%b0%88%e8%b0%88final%e3%80%81finally%e3%80%81%20finalize%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">03  谈谈final、finally、 finalize有什么不同？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  强引用、软引用、弱引用、幻象引用有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/04%20%20%e5%bc%ba%e5%bc%95%e7%94%a8%e3%80%81%e8%bd%af%e5%bc%95%e7%94%a8%e3%80%81%e5%bc%b1%e5%bc%95%e7%94%a8%e3%80%81%e5%b9%bb%e8%b1%a1%e5%bc%95%e7%94%a8%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">04  强引用、软引用、弱引用、幻象引用有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  String、StringBuffer、StringBuilder有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/05%20%20String%e3%80%81StringBuffer%e3%80%81StringBuilder%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">05  String、StringBuffer、StringBuilder有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  动态代理是基于什么原理？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/06%20%20%e5%8a%a8%e6%80%81%e4%bb%a3%e7%90%86%e6%98%af%e5%9f%ba%e4%ba%8e%e4%bb%80%e4%b9%88%e5%8e%9f%e7%90%86%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">06  动态代理是基于什么原理？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  int和Integer有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/07%20%20int%e5%92%8cInteger%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">07  int和Integer有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  对比Vector、ArrayList、LinkedList有何区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/08%20%20%e5%af%b9%e6%af%94Vector%e3%80%81ArrayList%e3%80%81LinkedList%e6%9c%89%e4%bd%95%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">08  对比Vector、ArrayList、LinkedList有何区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  对比Hashtable、HashMap、TreeMap有什么不同？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/09%20%20%e5%af%b9%e6%af%94Hashtable%e3%80%81HashMap%e3%80%81TreeMap%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">09  对比Hashtable、HashMap、TreeMap有什么不同？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  如何保证集合是线程安全的 ConcurrentHashMap如何实现高效地线程安全？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/10%20%20%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e9%9b%86%e5%90%88%e6%98%af%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e7%9a%84%20ConcurrentHashMap%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e5%9c%b0%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">10  如何保证集合是线程安全的 ConcurrentHashMap如何实现高效地线程安全？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  Java提供了哪些IO方式？ NIO如何实现多路复用？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/11%20%20Java%e6%8f%90%e4%be%9b%e4%ba%86%e5%93%aa%e4%ba%9bIO%e6%96%b9%e5%bc%8f%ef%bc%9f%20NIO%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%a4%9a%e8%b7%af%e5%a4%8d%e7%94%a8%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">11  Java提供了哪些IO方式？ NIO如何实现多路复用？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  Java有几种文件拷贝方式？哪一种最高效？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/12%20%20Java%e6%9c%89%e5%87%a0%e7%a7%8d%e6%96%87%e4%bb%b6%e6%8b%b7%e8%b4%9d%e6%96%b9%e5%bc%8f%ef%bc%9f%e5%93%aa%e4%b8%80%e7%a7%8d%e6%9c%80%e9%ab%98%e6%95%88%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">12  Java有几种文件拷贝方式？哪一种最高效？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  谈谈接口和抽象类有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/13%20%20%e8%b0%88%e8%b0%88%e6%8e%a5%e5%8f%a3%e5%92%8c%e6%8a%bd%e8%b1%a1%e7%b1%bb%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">13  谈谈接口和抽象类有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  谈谈你知道的设计模式？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/14%20%20%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9f%a5%e9%81%93%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">14  谈谈你知道的设计模式？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  synchronized和ReentrantLock有什么区别呢？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/15%20%20synchronized%e5%92%8cReentrantLock%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%e5%91%a2%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">15  synchronized和ReentrantLock有什么区别呢？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  synchronized底层如何实现？什么是锁的升级、降级？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/16%20%20synchronized%e5%ba%95%e5%b1%82%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f%e4%bb%80%e4%b9%88%e6%98%af%e9%94%81%e7%9a%84%e5%8d%87%e7%ba%a7%e3%80%81%e9%99%8d%e7%ba%a7%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">16  synchronized底层如何实现？什么是锁的升级、降级？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  一个线程两次调用start()方法会出现什么情况？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/17%20%20%e4%b8%80%e4%b8%aa%e7%ba%bf%e7%a8%8b%e4%b8%a4%e6%ac%a1%e8%b0%83%e7%94%a8start%28%29%e6%96%b9%e6%b3%95%e4%bc%9a%e5%87%ba%e7%8e%b0%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">17  一个线程两次调用start()方法会出现什么情况？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  什么情况下Java程序会产生死锁？如何定位、修复？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/18%20%20%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8bJava%e7%a8%8b%e5%ba%8f%e4%bc%9a%e4%ba%a7%e7%94%9f%e6%ad%bb%e9%94%81%ef%bc%9f%e5%a6%82%e4%bd%95%e5%ae%9a%e4%bd%8d%e3%80%81%e4%bf%ae%e5%a4%8d%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">18  什么情况下Java程序会产生死锁？如何定位、修复？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  Java并发包提供了哪些并发工具类？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/19%20%20Java%e5%b9%b6%e5%8f%91%e5%8c%85%e6%8f%90%e4%be%9b%e4%ba%86%e5%93%aa%e4%ba%9b%e5%b9%b6%e5%8f%91%e5%b7%a5%e5%85%b7%e7%b1%bb%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">19  Java并发包提供了哪些并发工具类？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  并发包中的ConcurrentLinkedQueue和LinkedBlockingQueue有什么区别？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/20%20%20%e5%b9%b6%e5%8f%91%e5%8c%85%e4%b8%ad%e7%9a%84ConcurrentLinkedQueue%e5%92%8cLinkedBlockingQueue%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">20  并发包中的ConcurrentLinkedQueue和LinkedBlockingQueue有什么区别？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  Java并发类库提供的线程池有哪几种？ 分别有什么特点？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/21%20%20Java%e5%b9%b6%e5%8f%91%e7%b1%bb%e5%ba%93%e6%8f%90%e4%be%9b%e7%9a%84%e7%ba%bf%e7%a8%8b%e6%b1%a0%e6%9c%89%e5%93%aa%e5%87%a0%e7%a7%8d%ef%bc%9f%20%e5%88%86%e5%88%ab%e6%9c%89%e4%bb%80%e4%b9%88%e7%89%b9%e7%82%b9%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">21  Java并发类库提供的线程池有哪几种？ 分别有什么特点？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  AtomicInteger底层实现原理是什么？如何在自己的产品代码中应用CAS操作？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/22%20%20AtomicInteger%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e5%a6%82%e4%bd%95%e5%9c%a8%e8%87%aa%e5%b7%b1%e7%9a%84%e4%ba%a7%e5%93%81%e4%bb%a3%e7%a0%81%e4%b8%ad%e5%ba%94%e7%94%a8CAS%e6%93%8d%e4%bd%9c%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">22  AtomicInteger底层实现原理是什么？如何在自己的产品代码中应用CAS操作？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  请介绍类加载过程，什么是双亲委派模型？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/23%20%20%e8%af%b7%e4%bb%8b%e7%bb%8d%e7%b1%bb%e5%8a%a0%e8%bd%bd%e8%bf%87%e7%a8%8b%ef%bc%8c%e4%bb%80%e4%b9%88%e6%98%af%e5%8f%8c%e4%ba%b2%e5%a7%94%e6%b4%be%e6%a8%a1%e5%9e%8b%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">23  请介绍类加载过程，什么是双亲委派模型？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  有哪些方法可以在运行时动态生成一个Java类？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/24%20%20%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b9%e6%b3%95%e5%8f%af%e4%bb%a5%e5%9c%a8%e8%bf%90%e8%a1%8c%e6%97%b6%e5%8a%a8%e6%80%81%e7%94%9f%e6%88%90%e4%b8%80%e4%b8%aaJava%e7%b1%bb%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">24  有哪些方法可以在运行时动态生成一个Java类？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/25%20%20%e8%b0%88%e8%b0%88JVM%e5%86%85%e5%ad%98%e5%8c%ba%e5%9f%9f%e7%9a%84%e5%88%92%e5%88%86%ef%bc%8c%e5%93%aa%e4%ba%9b%e5%8c%ba%e5%9f%9f%e5%8f%af%e8%83%bd%e5%8f%91%e7%94%9fOutOfMemoryError-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">25  谈谈JVM内存区域的划分，哪些区域可能发生OutOfMemoryError-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  如何监控和诊断JVM堆内和堆外内存使用？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/26%20%20%e5%a6%82%e4%bd%95%e7%9b%91%e6%8e%a7%e5%92%8c%e8%af%8a%e6%96%adJVM%e5%a0%86%e5%86%85%e5%92%8c%e5%a0%86%e5%a4%96%e5%86%85%e5%ad%98%e4%bd%bf%e7%94%a8%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">26  如何监控和诊断JVM堆内和堆外内存使用？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  Java常见的垃圾收集器有哪些？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/27%20%20Java%e5%b8%b8%e8%a7%81%e7%9a%84%e5%9e%83%e5%9c%be%e6%94%b6%e9%9b%86%e5%99%a8%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">27  Java常见的垃圾收集器有哪些？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  谈谈你的GC调优思路-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/28%20%20%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9a%84GC%e8%b0%83%e4%bc%98%e6%80%9d%e8%b7%af-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">28  谈谈你的GC调优思路-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  Java内存模型中的happen-before是什么？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/29%20%20Java%e5%86%85%e5%ad%98%e6%a8%a1%e5%9e%8b%e4%b8%ad%e7%9a%84happen-before%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">29  Java内存模型中的happen-before是什么？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  Java程序运行在Docker等容器环境有哪些新问题？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/30%20%20Java%e7%a8%8b%e5%ba%8f%e8%bf%90%e8%a1%8c%e5%9c%a8Docker%e7%ad%89%e5%ae%b9%e5%99%a8%e7%8e%af%e5%a2%83%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b0%e9%97%ae%e9%a2%98%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">30  Java程序运行在Docker等容器环境有哪些新问题？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  你了解Java应用开发中的注入攻击吗？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/31%20%20%e4%bd%a0%e4%ba%86%e8%a7%a3Java%e5%ba%94%e7%94%a8%e5%bc%80%e5%8f%91%e4%b8%ad%e7%9a%84%e6%b3%a8%e5%85%a5%e6%94%bb%e5%87%bb%e5%90%97%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">31  你了解Java应用开发中的注入攻击吗？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  如何写出安全的Java代码？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/32%20%20%e5%a6%82%e4%bd%95%e5%86%99%e5%87%ba%e5%ae%89%e5%85%a8%e7%9a%84Java%e4%bb%a3%e7%a0%81%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">32  如何写出安全的Java代码？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  后台服务出现明显“变慢”，谈谈你的诊断思路？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/33%20%20%e5%90%8e%e5%8f%b0%e6%9c%8d%e5%8a%a1%e5%87%ba%e7%8e%b0%e6%98%8e%e6%98%be%e2%80%9c%e5%8f%98%e6%85%a2%e2%80%9d%ef%bc%8c%e8%b0%88%e8%b0%88%e4%bd%a0%e7%9a%84%e8%af%8a%e6%96%ad%e6%80%9d%e8%b7%af%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">33  后台服务出现明显“变慢”，谈谈你的诊断思路？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/34%20%20%e6%9c%89%e4%ba%ba%e8%af%b4%e2%80%9cLambda%e8%83%bd%e8%ae%a9Java%e7%a8%8b%e5%ba%8f%e6%85%a230%e5%80%8d%e2%80%9d%ef%bc%8c%e4%bd%a0%e6%80%8e%e4%b9%88%e7%9c%8b%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">34  有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  JVM优化Java代码时都做了什么？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/35%20%20JVM%e4%bc%98%e5%8c%96Java%e4%bb%a3%e7%a0%81%e6%97%b6%e9%83%bd%e5%81%9a%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">35  JVM优化Java代码时都做了什么？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  谈谈MySQL支持的事务隔离级别，以及悲观锁和乐观锁的原理和应用场景？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/36%20%20%e8%b0%88%e8%b0%88MySQL%e6%94%af%e6%8c%81%e7%9a%84%e4%ba%8b%e5%8a%a1%e9%9a%94%e7%a6%bb%e7%ba%a7%e5%88%ab%ef%bc%8c%e4%bb%a5%e5%8f%8a%e6%82%b2%e8%a7%82%e9%94%81%e5%92%8c%e4%b9%90%e8%a7%82%e9%94%81%e7%9a%84%e5%8e%9f%e7%90%86%e5%92%8c%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">36  谈谈MySQL支持的事务隔离级别，以及悲观锁和乐观锁的原理和应用场景？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  谈谈Spring Bean的生命周期和作用域？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/37%20%20%e8%b0%88%e8%b0%88Spring%20Bean%e7%9a%84%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%e5%92%8c%e4%bd%9c%e7%94%a8%e5%9f%9f%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">37  谈谈Spring Bean的生命周期和作用域？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  对比Java标准NIO类库，你知道Netty是如何实现更高性能的吗？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/38%20%20%e5%af%b9%e6%af%94Java%e6%a0%87%e5%87%86NIO%e7%b1%bb%e5%ba%93%ef%bc%8c%e4%bd%a0%e7%9f%a5%e9%81%93Netty%e6%98%af%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%9b%b4%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84%e5%90%97%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">38  对比Java标准NIO类库，你知道Netty是如何实现更高性能的吗？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  谈谈常用的分布式ID的设计方案？Snowflake是否受冬令时切换影响？-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/39%20%20%e8%b0%88%e8%b0%88%e5%b8%b8%e7%94%a8%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8fID%e7%9a%84%e8%ae%be%e8%ae%a1%e6%96%b9%e6%a1%88%ef%bc%9fSnowflake%e6%98%af%e5%90%a6%e5%8f%97%e5%86%ac%e4%bb%a4%e6%97%b6%e5%88%87%e6%8d%a2%e5%bd%b1%e5%93%8d%ef%bc%9f-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">39  谈谈常用的分布式ID的设计方案？Snowflake是否受冬令时切换影响？-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="周末福利  一份Java工程师必读书单-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/%e5%91%a8%e6%9c%ab%e7%a6%8f%e5%88%a9%20%20%e4%b8%80%e4%bb%bdJava%e5%b7%a5%e7%a8%8b%e5%b8%88%e5%bf%85%e8%af%bb%e4%b9%a6%e5%8d%95-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">周末福利  一份Java工程师必读书单-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="周末福利  谈谈我对Java学习和面试的看法-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/%e5%91%a8%e6%9c%ab%e7%a6%8f%e5%88%a9%20%20%e8%b0%88%e8%b0%88%e6%88%91%e5%af%b9Java%e5%ad%a6%e4%b9%a0%e5%92%8c%e9%9d%a2%e8%af%95%e7%9a%84%e7%9c%8b%e6%b3%95-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">周末福利  谈谈我对Java学习和面试的看法-极客时间.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  技术没有终点-极客时间.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e7%b2%be%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e6%8a%80%e6%9c%af%e6%b2%a1%e6%9c%89%e7%bb%88%e7%82%b9-%e6%9e%81%e5%ae%a2%e6%97%b6%e9%97%b4.md">结束语  技术没有终点-极客时间.md</a>
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
                            <h1 id="title" data-id="34  有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间" class="title">34  有人说“Lambda能让Java程序慢30倍”，你怎么看？-极客时间</h1>
                            <div><p>在上一讲中，我介绍了 Java 性能问题分析的一些基本思路。但在实际工作中，我们不能仅仅等待性能出现问题再去试图解决，而是需要定量的、可对比的方法，去评估 Java 应用性能，来判断其是否能够符合业务支撑目标。今天这一讲，我会介绍从 Java 开发者角度，如何从代码级别判断应用的性能表现，重点理解最广泛使用的基准测试（Benchmark）。</p>

<p>今天我要问你的问题是，<strong>有人说“Lambda 能让 Java 程序慢 30 倍”，你怎么看？</strong></p>

<p>为了让你清楚地了解这个背景，请参考下面的代码片段。在实际运行中，基于 Lambda/Stream 的版本（lambdaMaxInteger），比传统的 for-each 版本（forEachLoopMaxInteger）慢很多。</p>

<pre><code class="language-java">// 一个大的ArrayList，内部是随机的整形数据
volatile List&lt;Integer&gt; integers = …
 
// 基准测试1
public int forEachLoopMaxInteger() {
   int max = Integer.MIN_VALUE;
   for (Integer n : integers) {
    max = Integer.max(max, n);
   }
   return max;
}
 
// 基准测试2
public int lambdaMaxInteger() {
   return integers.stream().reduce(Integer.MIN_VALUE, (a, b) -&gt; Integer.max(a, b));
}
</code></pre>

<h2 id="典型回答">典型回答</h2>

<p>我认为，“Lambda 能让 Java 程序慢 30 倍”这个争论实际反映了几个方面：</p>

<p>第一，基准测试是一个非常有效的通用手段，让我们以直观、量化的方式，判断程序在特定条件下的性能表现。</p>

<p>第二，基准测试必须明确定义自身的范围和目标，否则很有可能产生误导的结果。前面代码片段本身的逻辑就有瑕疵，更多的开销是源于自动装箱、拆箱（auto-boxing/unboxing），而不是源自 Lambda 和 Stream，所以得出的初始结论是没有说服力的。</p>

<p>第三，虽然 Lambda/Stream 为 Java 提供了强大的函数式编程能力，但是也需要正视其局限性：</p>

<ul>
<li>一般来说，我们可以认为 Lambda/Stream 提供了与传统方式接近对等的性能，但是如果对于性能非常敏感，就不能完全忽视它在特定场景的性能差异了，例如：<strong>初始化的开销</strong>。 Lambda 并不算是语法糖，而是一种新的工作机制，在首次调用时，JVM 需要为其构建<a href="https://docs.oracle.com/javase/8/docs/api/java/lang/invoke/CallSite.html" target="_blank">CallSite</a>实例。这意味着，如果 Java 应用启动过程引入了很多 Lambda 语句，会导致启动过程变慢。其实现特点决定了 JVM 对它的优化可能与传统方式存在差异。</li>
<li>增加了程序诊断等方面的复杂性，程序栈要复杂很多，Fluent 风格本身也不算是对于调试非常友好的结构，并且在可检查异常的处理方面也存在着局限性等。</li>
</ul>

<h2 id="考点分析">考点分析</h2>

<p>今天的题目是源自于一篇有争议的<a href="https://blog.takipi.com/benchmark-how-java-8-lambdas-and-streams-can-make-your-code-5-times-slower/" target="_blank">文章</a>，原文后来更正为“如果 Stream 使用不当，会让你的代码慢 5 倍”。针对这个问题我给出的回答，并没有纠结于所谓的“快”与“慢”，而是从工程实践的角度指出了基准测试本身存在的问题，以及 Lambda 自身的局限性。</p>

<p>从知识点的角度，这个问题考察了我在【专栏第 7 讲】中介绍过的自动装箱 / 拆箱机制对性能的影响，并且考察了 Java 8 中引入的 Lambda 特性的相关知识。除了这些知识点，面试官还可能更加深入探讨如何用基准测试之类的方法，将含糊的观点变成可验证的结论。</p>

<p>对于 Java 语言的很多特性，经常有很多似是而非的 “秘籍”，我们有必要去伪存真，以定量、定性的方式探究真相，探讨更加易于推广的实践。找到结论的能力，比结论本身更重要，因此在今天这一讲中，我们来探讨一下：</p>

<ul>
<li>基准测试的基础要素，以及如何利用主流框架构建简单的基准测试。</li>
<li>进一步分析，针对保证基准测试的有效性，如何避免偏离测试目的，如何保证基准测试的正确性。</li>
</ul>

<h2 id="知识扩展">知识扩展</h2>

<p>首先，我们先来整体了解一下基准测试的主要目的和特征，专栏里我就不重复那些<a href="https://baike.baidu.com/item/%E5%9F%BA%E5%87%86%E6%B5%8B%E8%AF%95" target="_blank">书面的定义</a>了。</p>

<p>性能往往是特定情景下的评价，泛泛地说性能“好”或者“快”，往往是具有误导性的。通过引入基准测试，我们可以定义性能对比的明确条件、具体的指标，进而保证得到<strong>定量的、可重复的</strong>对比数据，这是工程中的实际需要。</p>

<p>不同的基准测试其具体内容和范围也存在很大的不同。如果是专业的性能工程师，更加熟悉的可能是类似<a href="https://www.spec.org/" target="_blank">SPEC</a>提供的工业标准的系统级测试；而对于大多数 Java 开发者，更熟悉的则是范围相对较小、关注点更加细节的微基准测试（Micro-Benchmark）。我在文章开头提的问题，就是典型的微基准测试，也是我今天的侧重点。</p>

<p><strong>什么时候需要开发微基准测试呢？</strong></p>

<p>我认为，当需要对一个大型软件的某小部分的性能进行评估时，就可以考虑微基准测试。换句话说，微基准测试大多是 API 级别的验证，或者与其他简单用例场景的对比，例如：</p>

<ul>
<li>你在开发共享类库，为其他模块提供某种服务的 API 等。</li>
<li>你的 API 对于性能，如延迟、吞吐量有着严格的要求，例如，实现了定制的 HTTP 客户端 API，需要明确它对 HTTP 服务器进行大量 GET 请求时的吞吐能力，或者需要对比其他 API，保证至少对等甚至更高的性能标准。</li>
</ul>

<p>所以微基准测试更是偏基础、底层平台开发者的需求，当然，也是那些追求极致性能的前沿工程师的最爱。</p>

<p><strong>如何构建自己的微基准测试，选择什么样的框架比较好？</strong></p>

<p>目前应用最为广泛的框架之一就是<a href="https://openjdk.org/projects/code-tools/jmh/" target="_blank">JMH</a>，OpenJDK 自身也大量地使用 JMH 进行性能对比，如果你是做 Java API 级别的性能对比，JMH 往往是你的首选。</p>

<p>JMH 是由 Hotspot JVM 团队专家开发的，除了支持完整的基准测试过程，包括预热、运行、统计和报告等，还支持 Java 和其他 JVM 语言。更重要的是，它针对 Hotspot JVM 提供了各种特性，以保证基准测试的正确性，整体准确性大大优于其他框架，并且，JMH 还提供了用近乎白盒的方式进行 Profiling 等工作的能力。</p>

<p>使用 JMH 也非常简单，你可以直接将其依赖加入 Maven 工程，如下图：</p>

<p><img src="assets/0dd290f8842959cb02d6c3a434a58e68-20221127212308-eu9ogc7.png" alt="" /></p>

<p>也可以，利用类似下面的命令，直接生成一个 Maven 项目。</p>

<pre><code class="language-java">$ mvn archetype:generate \
        -DinteractiveMode=false \
        -DarchetypeGroupId=org.openjdk.jmh \
          -DarchetypeArtifactId=jmh-java-benchmark-archetype \
        -DgroupId=org.sample \
        -DartifactId=test \
        -Dversion=1.0
</code></pre>

<p>JMH 利用注解（Annotation），定义具体的测试方法，以及基准测试的详细配置。例如，至少要加上“@Benchmark”以标识它是个基准测试方法，而 BenchmarkMode 则指定了基准测试模式，例如下面例子指定了吞吐量（Throughput）模式，还可以根据需要指定平均时间（AverageTime）等其他模式。</p>

<pre><code class="language-java">@Benchmark
@BenchmarkMode(Mode.Throughput)
public void testMethod() {
   // Put your benchmark code here.
}
</code></pre>

<p>当我们实现了具体的测试后，就可以利用下面的 Maven 命令构建。</p>

<pre><code class="language-java">mvn clean install
</code></pre>

<p>运行基准测试则与运行不同的 Java 应用没有明显区别。</p>

<pre><code class="language-java">java -jar target/benchmarks.jar
</code></pre>

<p>更加具体的上手步骤，请参考相关<a href="https://www.baeldung.com/java-microbenchmark-harness" target="_blank">指南</a>。JMH 处处透着浓浓的工程师味道，并没有纠结于完善的文档，而是提供了非常棒的<a href="http://hg.openjdk.java.net/code-tools/jmh/file/3769055ad883/jmh-samples/src/main/java/org/openjdk/jmh/samples" target="_blank">样例代码</a>，所以你需要习惯于直接从代码中学习。</p>

<p><strong>如何保证微基准测试的正确性，有哪些坑需要规避？</strong></p>

<p>首先，构建微基准测试，需要从白盒层面理解代码，尤其是具体的性能开销，不管是 CPU 还是内存分配。这有两个方面的考虑，第一，需要保证我们写出的基准测试符合测试目的，确实验证的是我们要覆盖的功能点，这一讲的问题就是个典型例子；第二，通常对于微基准测试，我们通常希望代码片段确实是有限的，例如，执行时间如果需要很多毫秒（ms），甚至是秒级，那么这个有效性就要存疑了，也不便于诊断问题所在。</p>

<p>更加重要的是，由于微基准测试基本上都是体量较小的 API 层面测试，最大的威胁来自于过度“聪明”的 JVM！Brain Goetz 曾经很早就指出了微基准测试中的<a href="https://www.ibm.com/developerworks/java/library/j-jtp02225/" target="_blank">典型问题</a>。</p>

<p>由于我们执行的是非常有限的代码片段，必须要保证 JVM 优化过程不影响原始测试目的，下面几个方面需要重点关注：</p>

<ul>
<li>保证代码经过了足够并且合适的预热。我在【专栏第 1 讲】中提到过，默认情况，在 server 模式下，JIT 会在一段代码执行 10000 次后，将其编译为本地代码，client 模式则是 1500 次以后。我们需要排除代码执行初期的噪音，保证真正采样到的统计数据符合其稳定运行状态。</li>
</ul>

<p>通常建议使用下面的参数来判断预热工作到底是经过了多久。</p>

<pre><code class="language-java">-XX:+PrintCompilation
</code></pre>

<p>我这里建议考虑另外加上一个参数，否则 JVM 将默认开启后台编译，也就是在其他线程进行，可能导致输出的信息有些混淆。</p>

<pre><code class="language-java">-Xbatch
</code></pre>

<p>与此同时，也要保证预热阶段的代码路径和采集阶段的代码路径是一致的，并且可以观察 PrintCompilation 输出是否在后期运行中仍然有零星的编译语句出现。</p>

<ul>
<li>防止 JVM 进行无效代码消除（Dead Code Elimination），例如下面的代码片段中，由于我们并没有使用计算结果 mul，那么 JVM 就可能直接判断无效代码，根本就不执行它。</li>
</ul>

<pre><code class="language-java">public void testMethod() {
   int left = 10;
   int right = 100;
   int mul = left * right;
}
</code></pre>

<p>如果你发现代码统计数据发生了数量级程度上的提高，需要警惕是否出现了无效代码消除的问题。</p>

<p>解决办法也很直接，尽量保证方法有返回值，而不是 void 方法，或者使用 JMH 提供的<a href="">BlackHole</a>设施，在方法中添加下面语句。</p>

<pre><code class="language-java">public void testMethod(Blackhole blackhole) {
   // …
   blackhole.consume(mul);
}
</code></pre>

<ul>
<li>防止发生常量折叠（Constant Folding）。JVM 如果发现计算过程是依赖于常量或者事实上的常量，就可能会直接计算其结果，所以基准测试并不能真实反映代码执行的性能。JMH 提供了 State 机制来解决这个问题，将本地变量修改为 State 对象信息，请参考下面示例。</li>
</ul>

<pre><code class="language-java">@State(Scope.Thread)
public static class MyState {
   public int left = 10;
   public int right = 100;
}

public void testMethod(MyState state, Blackhole blackhole) {
   int left = state.left;
   int right = state.right;
   int mul = left * right;
   blackhole.consume(mul);
}
</code></pre>

<ul>
<li>另外 JMH 还会对 State 对象进行额外的处理，以尽量消除伪共享（<a href="https://blogs.oracle.com/dave/java-contended-annotation-to-help-reduce-false-sharing" target="_blank">False Sharing</a>）的影响，标记 @State，JMH 会自动进行补齐。</li>
<li>如果你希望确定方法内联（Inlining）对性能的影响，可以考虑打开下面的选项。</li>
</ul>

<pre><code class="language-java">-XX:+PrintInlining
</code></pre>

<p>从上面的总结，可以看出来微基准测试是一个需要高度了解 Java、JVM 底层机制的技术，是个非常好的深入理解程序背后效果的工具，但是也反映了我们需要审慎对待微基准测试，不被可能的假象蒙蔽。</p>

<p>我今天介绍的内容是相对常见并易于把握的，对于微基准测试，GC 等基层机制同样会影响其统计数据。我在前面提到，微基准测试通常希望执行时间和内存分配速率都控制在有限范围内，而在这个过程中发生 GC，很可能导致数据出现偏差，所以 Serial GC 是个值得考虑的选项。另外，JDK 11 引入了<a href="http://openjdk.java.net/jeps/318" target="_blank">Epsilon GC</a>，可以考虑使用这种什么也不做的 GC 方式，从最大可能性去排除相关影响。</p>

<p>今天我从一个争议性的程序开始，探讨了如何从开发者角度而不是性能工程师角度，利用（微）基准测试验证你在性能上的判断，并且介绍了其基础构建方式和需要重点规避的风险点。</p>

<h2 id="一课一练">一课一练</h2>

<p>关于今天我们讨论的题目你做到心中有数了吗？我们在项目中需要评估系统的容量，以计划和保证其业务支撑能力，谈谈你的思路是怎么样的？常用手段有哪些？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#08646464313c3939383f486f65696164266b6765" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0ef3c1be43ecfb',t:'MTczNDAxOTQ4Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
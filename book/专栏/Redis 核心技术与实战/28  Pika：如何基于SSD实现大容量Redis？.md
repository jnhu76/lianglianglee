<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=28&#32;&#32;Pika：如何基于SSD实现大容量Redis？>
        <link rel="icon" href="/static/favicon.png">
        <title>28  Pika：如何基于SSD实现大容量Redis？ </title>
        
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
                        <a class="menu-item" id="00  开篇词  这样学Redis，才能技高一筹.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e8%bf%99%e6%a0%b7%e5%ad%a6Redis%ef%bc%8c%e6%89%8d%e8%83%bd%e6%8a%80%e9%ab%98%e4%b8%80%e7%ad%b9.md">00  开篇词  这样学Redis，才能技高一筹.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  基本架构：一个键值数据库包含什么？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%20%e5%9f%ba%e6%9c%ac%e6%9e%b6%e6%9e%84%ef%bc%9a%e4%b8%80%e4%b8%aa%e9%94%ae%e5%80%bc%e6%95%b0%e6%8d%ae%e5%ba%93%e5%8c%85%e5%90%ab%e4%bb%80%e4%b9%88%ef%bc%9f.md">01  基本架构：一个键值数据库包含什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  数据结构：快速的Redis有哪些慢操作？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/02%20%20%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%ef%bc%9a%e5%bf%ab%e9%80%9f%e7%9a%84Redis%e6%9c%89%e5%93%aa%e4%ba%9b%e6%85%a2%e6%93%8d%e4%bd%9c%ef%bc%9f.md">02  数据结构：快速的Redis有哪些慢操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  高性能IO模型：为什么单线程Redis能那么快？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/03%20%20%e9%ab%98%e6%80%a7%e8%83%bdIO%e6%a8%a1%e5%9e%8b%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%8d%95%e7%ba%bf%e7%a8%8bRedis%e8%83%bd%e9%82%a3%e4%b9%88%e5%bf%ab%ef%bc%9f.md">03  高性能IO模型：为什么单线程Redis能那么快？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  AOF日志：宕机了，Redis如何避免数据丢失？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/04%20%20AOF%e6%97%a5%e5%bf%97%ef%bc%9a%e5%ae%95%e6%9c%ba%e4%ba%86%ef%bc%8cRedis%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e6%95%b0%e6%8d%ae%e4%b8%a2%e5%a4%b1%ef%bc%9f.md">04  AOF日志：宕机了，Redis如何避免数据丢失？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  内存快照：宕机后，Redis如何实现快速恢复？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%20%e5%86%85%e5%ad%98%e5%bf%ab%e7%85%a7%ef%bc%9a%e5%ae%95%e6%9c%ba%e5%90%8e%ef%bc%8cRedis%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%bf%ab%e9%80%9f%e6%81%a2%e5%a4%8d%ef%bc%9f.md">05  内存快照：宕机后，Redis如何实现快速恢复？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  数据同步：主从库如何实现数据一致？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/06%20%20%e6%95%b0%e6%8d%ae%e5%90%8c%e6%ad%a5%ef%bc%9a%e4%b8%bb%e4%bb%8e%e5%ba%93%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e4%b8%80%e8%87%b4%ef%bc%9f.md">06  数据同步：主从库如何实现数据一致？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  哨兵机制：主库挂了，如何不间断服务？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%20%e5%93%a8%e5%85%b5%e6%9c%ba%e5%88%b6%ef%bc%9a%e4%b8%bb%e5%ba%93%e6%8c%82%e4%ba%86%ef%bc%8c%e5%a6%82%e4%bd%95%e4%b8%8d%e9%97%b4%e6%96%ad%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">07  哨兵机制：主库挂了，如何不间断服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  哨兵集群：哨兵挂了，主从库还能切换吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%20%e5%93%a8%e5%85%b5%e9%9b%86%e7%be%a4%ef%bc%9a%e5%93%a8%e5%85%b5%e6%8c%82%e4%ba%86%ef%bc%8c%e4%b8%bb%e4%bb%8e%e5%ba%93%e8%bf%98%e8%83%bd%e5%88%87%e6%8d%a2%e5%90%97%ef%bc%9f.md">08  哨兵集群：哨兵挂了，主从库还能切换吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  切片集群：数据增多了，是该加内存还是加实例？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%20%e5%88%87%e7%89%87%e9%9b%86%e7%be%a4%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%a2%9e%e5%a4%9a%e4%ba%86%ef%bc%8c%e6%98%af%e8%af%a5%e5%8a%a0%e5%86%85%e5%ad%98%e8%bf%98%e6%98%af%e5%8a%a0%e5%ae%9e%e4%be%8b%ef%bc%9f.md">09  切片集群：数据增多了，是该加内存还是加实例？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  第1～9讲课后思考题答案及常见问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%20%e7%ac%ac1%ef%bd%9e9%e8%ae%b2%e8%af%be%e5%90%8e%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%8f%8a%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">10  第1～9讲课后思考题答案及常见问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  “万金油”的String，为什么不好用了？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%20%e2%80%9c%e4%b8%87%e9%87%91%e6%b2%b9%e2%80%9d%e7%9a%84String%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%8d%e5%a5%bd%e7%94%a8%e4%ba%86%ef%bc%9f.md">11  “万金油”的String，为什么不好用了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  有一亿个keys要统计，应该用哪种集合？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%20%e6%9c%89%e4%b8%80%e4%ba%bf%e4%b8%aakeys%e8%a6%81%e7%bb%9f%e8%ae%a1%ef%bc%8c%e5%ba%94%e8%af%a5%e7%94%a8%e5%93%aa%e7%a7%8d%e9%9b%86%e5%90%88%ef%bc%9f.md">12  有一亿个keys要统计，应该用哪种集合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  GEO是什么？还可以定义新的数据类型吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%20GEO%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e8%bf%98%e5%8f%af%e4%bb%a5%e5%ae%9a%e4%b9%89%e6%96%b0%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%e5%90%97%ef%bc%9f.md">13  GEO是什么？还可以定义新的数据类型吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  如何在Redis中保存时间序列数据？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%20%e5%a6%82%e4%bd%95%e5%9c%a8Redis%e4%b8%ad%e4%bf%9d%e5%ad%98%e6%97%b6%e9%97%b4%e5%ba%8f%e5%88%97%e6%95%b0%e6%8d%ae%ef%bc%9f.md">14  如何在Redis中保存时间序列数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  消息队列的考验：Redis有哪些解决方案？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/15%20%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e7%9a%84%e8%80%83%e9%aa%8c%ef%bc%9aRedis%e6%9c%89%e5%93%aa%e4%ba%9b%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88%ef%bc%9f.md">15  消息队列的考验：Redis有哪些解决方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  异步机制：如何避免单线程模型的阻塞？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%20%e5%bc%82%e6%ad%a5%e6%9c%ba%e5%88%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e5%8d%95%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%9e%8b%e7%9a%84%e9%98%bb%e5%a1%9e%ef%bc%9f.md">16  异步机制：如何避免单线程模型的阻塞？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  为什么CPU结构也会影响Redis的性能？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%20%e4%b8%ba%e4%bb%80%e4%b9%88CPU%e7%bb%93%e6%9e%84%e4%b9%9f%e4%bc%9a%e5%bd%b1%e5%93%8dRedis%e7%9a%84%e6%80%a7%e8%83%bd%ef%bc%9f.md">17  为什么CPU结构也会影响Redis的性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  波动的响应延迟：如何应对变慢的Redis？（上）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/18%20%20%e6%b3%a2%e5%8a%a8%e7%9a%84%e5%93%8d%e5%ba%94%e5%bb%b6%e8%bf%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e5%8f%98%e6%85%a2%e7%9a%84Redis%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">18  波动的响应延迟：如何应对变慢的Redis？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  波动的响应延迟：如何应对变慢的Redis？（下）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/19%20%20%e6%b3%a2%e5%8a%a8%e7%9a%84%e5%93%8d%e5%ba%94%e5%bb%b6%e8%bf%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e5%8f%98%e6%85%a2%e7%9a%84Redis%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">19  波动的响应延迟：如何应对变慢的Redis？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  删除数据后，为什么内存占用率还是很高？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%20%e5%88%a0%e9%99%a4%e6%95%b0%e6%8d%ae%e5%90%8e%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e5%86%85%e5%ad%98%e5%8d%a0%e7%94%a8%e7%8e%87%e8%bf%98%e6%98%af%e5%be%88%e9%ab%98%ef%bc%9f.md">20  删除数据后，为什么内存占用率还是很高？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  缓冲区：一个可能引发“惨案”的地方.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/21%20%20%e7%bc%93%e5%86%b2%e5%8c%ba%ef%bc%9a%e4%b8%80%e4%b8%aa%e5%8f%af%e8%83%bd%e5%bc%95%e5%8f%91%e2%80%9c%e6%83%a8%e6%a1%88%e2%80%9d%e7%9a%84%e5%9c%b0%e6%96%b9.md">21  缓冲区：一个可能引发“惨案”的地方.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  第11～21讲课后思考题答案及常见问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%20%e7%ac%ac11%ef%bd%9e21%e8%ae%b2%e8%af%be%e5%90%8e%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%8f%8a%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">22  第11～21讲课后思考题答案及常见问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  旁路缓存：Redis是如何工作的？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%20%e6%97%81%e8%b7%af%e7%bc%93%e5%ad%98%ef%bc%9aRedis%e6%98%af%e5%a6%82%e4%bd%95%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%9f.md">23  旁路缓存：Redis是如何工作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  替换策略：缓存满了怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%20%e6%9b%bf%e6%8d%a2%e7%ad%96%e7%95%a5%ef%bc%9a%e7%bc%93%e5%ad%98%e6%bb%a1%e4%ba%86%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">24  替换策略：缓存满了怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  缓存异常（上）：如何解决缓存和数据库的数据不一致问题？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%20%e7%bc%93%e5%ad%98%e5%bc%82%e5%b8%b8%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e7%bc%93%e5%ad%98%e5%92%8c%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e6%95%b0%e6%8d%ae%e4%b8%8d%e4%b8%80%e8%87%b4%e9%97%ae%e9%a2%98%ef%bc%9f.md">25  缓存异常（上）：如何解决缓存和数据库的数据不一致问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  缓存异常（下）：如何解决缓存雪崩、击穿、穿透难题？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%20%e7%bc%93%e5%ad%98%e5%bc%82%e5%b8%b8%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e7%bc%93%e5%ad%98%e9%9b%aa%e5%b4%a9%e3%80%81%e5%87%bb%e7%a9%bf%e3%80%81%e7%a9%bf%e9%80%8f%e9%9a%be%e9%a2%98%ef%bc%9f.md">26  缓存异常（下）：如何解决缓存雪崩、击穿、穿透难题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  缓存被污染了，该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%20%e7%bc%93%e5%ad%98%e8%a2%ab%e6%b1%a1%e6%9f%93%e4%ba%86%ef%bc%8c%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">27  缓存被污染了，该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  Pika：如何基于SSD实现大容量Redis？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%20Pika%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8eSSD%e5%ae%9e%e7%8e%b0%e5%a4%a7%e5%ae%b9%e9%87%8fRedis%ef%bc%9f.md">28  Pika：如何基于SSD实现大容量Redis？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  无锁的原子操作：Redis如何应对并发访问？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%20%e6%97%a0%e9%94%81%e7%9a%84%e5%8e%9f%e5%ad%90%e6%93%8d%e4%bd%9c%ef%bc%9aRedis%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e5%b9%b6%e5%8f%91%e8%ae%bf%e9%97%ae%ef%bc%9f.md">29  无锁的原子操作：Redis如何应对并发访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  如何使用Redis实现分布式锁？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%20%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8Redis%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%ef%bc%9f.md">30  如何使用Redis实现分布式锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  事务机制：Redis能实现ACID属性吗？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/31%20%20%e4%ba%8b%e5%8a%a1%e6%9c%ba%e5%88%b6%ef%bc%9aRedis%e8%83%bd%e5%ae%9e%e7%8e%b0ACID%e5%b1%9e%e6%80%a7%e5%90%97%ef%bc%9f.md">31  事务机制：Redis能实现ACID属性吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  Redis主从同步与故障切换，有哪些坑？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%20Redis%e4%b8%bb%e4%bb%8e%e5%90%8c%e6%ad%a5%e4%b8%8e%e6%95%85%e9%9a%9c%e5%88%87%e6%8d%a2%ef%bc%8c%e6%9c%89%e5%93%aa%e4%ba%9b%e5%9d%91%ef%bc%9f.md">32  Redis主从同步与故障切换，有哪些坑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  脑裂：一次奇怪的数据丢失.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/33%20%20%e8%84%91%e8%a3%82%ef%bc%9a%e4%b8%80%e6%ac%a1%e5%a5%87%e6%80%aa%e7%9a%84%e6%95%b0%e6%8d%ae%e4%b8%a2%e5%a4%b1.md">33  脑裂：一次奇怪的数据丢失.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  第23~33讲课后思考题答案及常见问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/34%20%20%e7%ac%ac23~33%e8%ae%b2%e8%af%be%e5%90%8e%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%8f%8a%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">34  第23~33讲课后思考题答案及常见问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  Codis VS Redis Cluster：我该选择哪一个集群方案？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/35%20%20Codis%20VS%20Redis%20Cluster%ef%bc%9a%e6%88%91%e8%af%a5%e9%80%89%e6%8b%a9%e5%93%aa%e4%b8%80%e4%b8%aa%e9%9b%86%e7%be%a4%e6%96%b9%e6%a1%88%ef%bc%9f.md">35  Codis VS Redis Cluster：我该选择哪一个集群方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  Redis支撑秒杀场景的关键技术和实践都有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/36%20%20Redis%e6%94%af%e6%92%91%e7%a7%92%e6%9d%80%e5%9c%ba%e6%99%af%e7%9a%84%e5%85%b3%e9%94%ae%e6%8a%80%e6%9c%af%e5%92%8c%e5%ae%9e%e8%b7%b5%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">36  Redis支撑秒杀场景的关键技术和实践都有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  数据分布优化：如何应对数据倾斜？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/37%20%20%e6%95%b0%e6%8d%ae%e5%88%86%e5%b8%83%e4%bc%98%e5%8c%96%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e6%95%b0%e6%8d%ae%e5%80%be%e6%96%9c%ef%bc%9f.md">37  数据分布优化：如何应对数据倾斜？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  通信开销：限制Redis Cluster规模的关键因素.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/38%20%20%e9%80%9a%e4%bf%a1%e5%bc%80%e9%94%80%ef%bc%9a%e9%99%90%e5%88%b6Redis%20Cluster%e8%a7%84%e6%a8%a1%e7%9a%84%e5%85%b3%e9%94%ae%e5%9b%a0%e7%b4%a0.md">38  通信开销：限制Redis Cluster规模的关键因素.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  Redis 6.0的新特性：多线程、客户端缓存与安全.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/39%20%20Redis%206.0%e7%9a%84%e6%96%b0%e7%89%b9%e6%80%a7%ef%bc%9a%e5%a4%9a%e7%ba%bf%e7%a8%8b%e3%80%81%e5%ae%a2%e6%88%b7%e7%ab%af%e7%bc%93%e5%ad%98%e4%b8%8e%e5%ae%89%e5%85%a8.md">39  Redis 6.0的新特性：多线程、客户端缓存与安全.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  Redis的下一步：基于NVM内存的实践.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/40%20%20Redis%e7%9a%84%e4%b8%8b%e4%b8%80%e6%ad%a5%ef%bc%9a%e5%9f%ba%e4%ba%8eNVM%e5%86%85%e5%ad%98%e7%9a%84%e5%ae%9e%e8%b7%b5.md">40  Redis的下一步：基于NVM内存的实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41   第35～40讲课后思考题答案及常见问题答疑.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/41%20%20%20%e7%ac%ac35%ef%bd%9e40%e8%ae%b2%e8%af%be%e5%90%8e%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%8f%8a%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">41   第35～40讲课后思考题答案及常见问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 01  经典的Redis学习资料有哪些？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2001%20%20%e7%bb%8f%e5%85%b8%e7%9a%84Redis%e5%ad%a6%e4%b9%a0%e8%b5%84%e6%96%99%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">加餐 01  经典的Redis学习资料有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 02  用户Kaito：我是如何学习Redis的？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2002%20%20%e7%94%a8%e6%88%b7Kaito%ef%bc%9a%e6%88%91%e6%98%af%e5%a6%82%e4%bd%95%e5%ad%a6%e4%b9%a0Redis%e7%9a%84%ef%bc%9f.md">加餐 02  用户Kaito：我是如何学习Redis的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 03  用户Kaito：我希望成为在压力中成长的人.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2003%20%20%e7%94%a8%e6%88%b7Kaito%ef%bc%9a%e6%88%91%e5%b8%8c%e6%9c%9b%e6%88%90%e4%b8%ba%e5%9c%a8%e5%8e%8b%e5%8a%9b%e4%b8%ad%e6%88%90%e9%95%bf%e7%9a%84%e4%ba%ba.md">加餐 03  用户Kaito：我希望成为在压力中成长的人.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 04   Redis客户端如何与服务器端交换命令和数据？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2004%20%20%20Redis%e5%ae%a2%e6%88%b7%e7%ab%af%e5%a6%82%e4%bd%95%e4%b8%8e%e6%9c%8d%e5%8a%a1%e5%99%a8%e7%ab%af%e4%ba%a4%e6%8d%a2%e5%91%bd%e4%bb%a4%e5%92%8c%e6%95%b0%e6%8d%ae%ef%bc%9f.md">加餐 04   Redis客户端如何与服务器端交换命令和数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 05   Redis有哪些好用的运维工具？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2005%20%20%20Redis%e6%9c%89%e5%93%aa%e4%ba%9b%e5%a5%bd%e7%94%a8%e7%9a%84%e8%bf%90%e7%bb%b4%e5%b7%a5%e5%85%b7%ef%bc%9f.md">加餐 05   Redis有哪些好用的运维工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 06   Redis的使用规范小建议.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2006%20%20%20Redis%e7%9a%84%e4%bd%bf%e7%94%a8%e8%a7%84%e8%8c%83%e5%b0%8f%e5%bb%ba%e8%ae%ae.md">加餐 06   Redis的使用规范小建议.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 07  从微博的Redis实践中，我们可以学到哪些经验？.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%2007%20%20%e4%bb%8e%e5%be%ae%e5%8d%9a%e7%9a%84Redis%e5%ae%9e%e8%b7%b5%e4%b8%ad%ef%bc%8c%e6%88%91%e4%bb%ac%e5%8f%af%e4%bb%a5%e5%ad%a6%e5%88%b0%e5%93%aa%e4%ba%9b%e7%bb%8f%e9%aa%8c%ef%bc%9f.md">加餐 07  从微博的Redis实践中，我们可以学到哪些经验？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  从学习Redis到向Redis学习.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bb%8e%e5%ad%a6%e4%b9%a0Redis%e5%88%b0%e5%90%91Redis%e5%ad%a6%e4%b9%a0.md">结束语  从学习Redis到向Redis学习.md</a>
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
                            <h1 id="title" data-id="28  Pika：如何基于SSD实现大容量Redis？" class="title">28  Pika：如何基于SSD实现大容量Redis？</h1>
                            <div><p>我们在应用 Redis 时，随着业务数据的增加（比如说电商业务中，随着用户规模和商品数量的增加），就需要 Redis 能保存更多的数据。你可能会想到使用 Redis 切片集群，把数据分散保存到多个实例上。但是这样做的话，会有一个问题，如果要保存的数据总量很大，但是每个实例保存的数据量较小的话，就会导致集群的实例规模增加，这会让集群的运维管理变得复杂，增加开销。</p>

<p>你可能又会说，我们可以通过增加 Redis 单实例的内存容量，形成大内存实例，每个实例可以保存更多的数据，这样一来，在保存相同的数据总量时，所需要的大内存实例的个数就会减少，就可以节省开销。</p>

<p>这是一个好主意，但这也并不是完美的方案：基于大内存的大容量实例在实例恢复、主从同步过程中会引起一系列潜在问题，例如恢复时间增长、主从切换开销大、缓冲区易溢出。</p>

<p>那怎么办呢？我推荐你使用固态硬盘（Solid State Drive，SSD）。它的成本很低（每 GB 的成本约是内存的十分之一），而且容量大，读写速度快，我们可以基于 SSD 来实现大容量的 Redis 实例。360 公司 DBA 和基础架构组联合开发的 Pika<a href="https://github.com/Qihoo360/pika" target="_blank">键值数据库</a>，正好实现了这一需求。</p>

<p>Pika 在刚开始设计的时候，就有两个目标：一是，单实例可以保存大容量数据，同时避免了实例恢复和主从同步时的潜在问题；二是，和 Redis 数据类型保持兼容，可以支持使用 Redis 的应用平滑地迁移到 Pika 上。所以，如果你一直在使用 Redis，并且想使用 SSD 来扩展单实例容量，Pika 就是一个很好的选择。</p>

<p>这节课，我就和你聊聊 Pika。在介绍 Pika 前，我先给你具体解释下基于大内存实现大容量 Redis 实例的潜在问题。只有知道了这些问题，我们才能选择更合适的方案。另外呢，我还会带你一步步分析下 Pika 是如何实现刚刚我们所说的两个设计目标，解决这些问题的。</p>

<h2 id="大内存-redis-实例的潜在问题">大内存 Redis 实例的潜在问题</h2>

<p>Redis 使用内存保存数据，内存容量增加后，就会带来两方面的潜在问题，分别是，内存快照 RDB 生成和恢复效率低，以及主从节点全量同步时长增加、缓冲区易溢出。我来一一解释下，</p>

<p>我们先看内存快照 RDB 受到的影响。内存大小和内存快照 RDB 的关系是非常直接的：实例内存容量大，RDB 文件也会相应增大，那么，RDB 文件生成时的 fork 时长就会增加，这就会导致 Redis 实例阻塞。而且，RDB 文件增大后，使用 RDB 进行恢复的时长也会增加，会导致 Redis 较长时间无法对外提供服务。</p>

<p>接下来我们再来看下主从同步受到的影响，</p>

<p>主从节点间的同步的第一步就是要做全量同步。全量同步是主节点生成 RDB 文件，并传给从节点，从节点再进行加载。试想一下，如果 RDB 文件很大，肯定会导致全量同步的时长增加，效率不高，而且还可能会导致复制缓冲区溢出。一旦缓冲区溢出了，主从节点间就会又开始全量同步，影响业务应用的正常使用。如果我们增加复制缓冲区的容量，这又会消耗宝贵的内存资源。</p>

<p>此外，如果主库发生了故障，进行主从切换后，其他从库都需要和新主库进行一次全量同步。如果 RDB 文件很大，也会导致主从切换的过程耗时增加，同样会影响业务的可用性。</p>

<p>那么，Pika 是如何解决这两方面的问题呢？这就要提到 Pika 中的关键模块 RocksDB、binlog 机制和 Nemo 了，这些模块都是 Pika 架构中的重要组成部分。所以，接下来，我们就来先看下 Pika 的整体架构。</p>

<h2 id="pika-的整体架构">Pika 的整体架构</h2>

<p>Pika 键值数据库的整体架构中包括了五部分，分别是网络框架、Pika 线程模块、Nemo 存储模块、RocksDB 和 binlog 机制，如下图所示：</p>

<p><img src="assets/a1421b8dbca6bb1ee9b6c1be7a929ae7-20221015223826-i08atjs.jpg" alt="" /></p>

<p>这五个部分分别实现了不同的功能，下面我一个个来介绍下。</p>

<p>首先，网络框架主要负责底层网络请求的接收和发送。Pika 的网络框架是对操作系统底层的网络函数进行了封装。Pika 在进行网络通信时，可以直接调用网络框架封装好的函数。</p>

<p>其次，Pika 线程模块采用了多线程模型来具体处理客户端请求，包括一个请求分发线程（DispatchThread）、一组工作线程（WorkerThread）以及一个线程池（ThreadPool）。</p>

<p>请求分发线程专门监听网络端口，一旦接收到客户端的连接请求后，就和客户端建立连接，并把连接交由工作线程处理。工作线程负责接收客户端连接上发送的具体命令请求，并把命令请求封装成 Task，再交给线程池中的线程，由这些线程进行实际的数据存取处理，如下图所示：</p>

<p><img src="assets/4627f13848167cdaa3b30370d9b80a06-20221015223826-ctxrr3y.jpg" alt="" /></p>

<p>在实际应用 Pika 的时候，我们可以通过增加工作线程数和线程池中的线程数，来提升 Pika 的请求处理吞吐率，进而满足业务层对数据处理性能的需求。</p>

<p>Nemo 模块很容易理解，它实现了 Pika 和 Redis 的数据类型兼容。这样一来，当我们把 Redis 服务迁移到 Pika 时，不用修改业务应用中操作 Redis 的代码，而且还可以继续应用运维 Redis 的经验，这使得 Pika 的学习成本就较低。Nemo 模块对数据类型的具体转换机制是我们要重点关心的，下面我会具体介绍。</p>

<p>最后，我们再来看看 RocksDB 提供的基于 SSD 保存数据的功能。它使得 Pika 可以不用大容量的内存，就能保存更多数据，还避免了使用内存快照。而且，Pika 使用 binlog 机制记录写命令，用于主从节点的命令同步，避免了刚刚所说的大内存实例在主从同步过程中的潜在问题。</p>

<p>接下来，我们就来具体了解下，Pika 是如何使用 RocksDB 和 binlog 机制的。</p>

<h2 id="pika-如何基于-ssd-保存更多数据">Pika 如何基于 SSD 保存更多数据？</h2>

<p>为了把数据保存到 SSD，Pika 使用了业界广泛应用的持久化键值数据库<a href="https://rocksdb.org/" target="_blank">RocksDB</a>。RocksDB 本身的实现机制较为复杂，你不需要全部弄明白，你只要记住 RocksDB 的基本数据读写机制，对于学习了解 Pika 来说，就已经足够了。下面我来解释下这个基本读写机制。</p>

<p>下面我结合一张图片，来给你具体介绍下 RocksDB 写入数据的基本流程。</p>

<p><img src="assets/95d97d3cf0f1555b65b47fb256b7b81d-20221015223826-bxs2sbv.jpg" alt="" /></p>

<p>当 Pika 需要保存数据时，RocksDB 会使用两小块内存空间（Memtable1 和 Memtable2）来交替缓存写入的数据。Memtable 的大小可以设置，一个 Memtable 的大小一般为几 MB 或几十 MB。当有数据要写入 RocksDB 时，RocksDB 会先把数据写入到 Memtable1。等到 Memtable1 写满后，RocksDB 再把数据以文件的形式，快速写入底层的 SSD。同时，RocksDB 会使用 Memtable2 来代替 Memtable1，缓存新写入的数据。等到 Memtable1 的数据都写入 SSD 了，RocksDB 会在 Memtable2 写满后，再用 Memtable1 缓存新写入的数据。</p>

<p>这么一分析你就知道了，RocksDB 会先用 Memtable 缓存数据，再将数据快速写入 SSD，即使数据量再大，所有数据也都能保存到 SSD 中。而且，Memtable 本身容量不大，即使 RocksDB 使用了两个 Memtable，也不会占用过多的内存，这样一来，Pika 在保存大容量数据时，也不用占据太大的内存空间了。</p>

<p>当 Pika 需要读取数据的时候，RocksDB 会先在 Memtable 中查询是否有要读取的数据。这是因为，最新的数据都是先写入到 Memtable 中的。如果 Memtable 中没有要读取的数据，RocksDB 会再查询保存在 SSD 上的数据文件，如下图所示：</p>

<p><img src="assets/aa70655efbb767af499a83bd6521ee3b-20221015223826-xw696z9.jpg" alt="" /></p>

<p>到这里，你就了解了，当使用了 RocksDB 保存数据后，Pika 就可以把大量数据保存到大容量的 SSD 上了，实现了大容量实例。不过，我刚才向你介绍过，当使用大内存实例保存大量数据时，Redis 会面临 RDB 生成和恢复的效率问题，以及主从同步时的效率和缓冲区溢出问题。那么，当 Pika 保存大量数据时，还会面临相同的问题吗？</p>

<p>其实不会了，我们来分析一下。</p>

<p>一方面，Pika 基于 RocksDB 保存了数据文件，直接读取数据文件就能恢复，不需要再通过内存快照进行恢复了。而且，Pika 从库在进行全量同步时，可以直接从主库拷贝数据文件，不需要使用内存快照，这样一来，Pika 就避免了大内存快照生成效率低的问题。</p>

<p>另一方面，Pika 使用了 binlog 机制实现增量命令同步，既节省了内存，还避免了缓冲区溢出的问题。binlog 是保存在 SSD 上的文件，Pika 接收到写命令后，在把数据写入 Memtable 时，也会把命令操作写到 binlog 文件中。和 Redis 类似，当全量同步结束后，从库会从 binlog 中把尚未同步的命令读取过来，这样就可以和主库的数据保持一致。当进行增量同步时，从库也是把自己已经复制的偏移量发给主库，主库把尚未同步的命令发给从库，来保持主从库的数据一致。</p>

<p>不过，和 Redis 使用缓冲区相比，使用 binlog 好处是非常明显的：binlog 是保存在 SSD 上的文件，文件大小不像缓冲区，会受到内存容量的较多限制。而且，当 binlog 文件增大后，还可以通过轮替操作，生成新的 binlog 文件，再把旧的 binlog 文件独立保存。这样一来，即使 Pika 实例保存了大量的数据，在同步过程中也不会出现缓冲区溢出的问题了。</p>

<p>现在，我们先简单小结下。Pika 使用 RocksDB 把大量数据保存到了 SSD，同时避免了内存快照的生成和恢复问题。而且，Pika 使用 binlog 机制进行主从同步，避免大内存时的影响，Pika 的第一个设计目标就实现了。</p>

<p>接下来，我们再来看 Pika 是如何实现第二个设计目标的，也就是如何和 Redis 兼容。毕竟，如果不兼容的话，原来使用 Redis 的业务就无法平滑迁移到 Pika 上使用了，也就没办法利用 Pika 保存大容量数据的优势了。</p>

<h2 id="pika-如何实现-redis-数据类型兼容">Pika 如何实现 Redis 数据类型兼容？</h2>

<p>Pika 的底层存储使用了 RocksDB 来保存数据，但是，RocksDB 只提供了单值的键值对类型，RocksDB 键值对中的值就是单个值，而 Redis 键值对中的值还可以是集合类型。</p>

<p>对于 Redis 的 String 类型来说，它本身就是单值的键值对，我们直接用 RocksDB 保存就行。但是，对于集合类型来说，我们就无法直接把集合保存为单值的键值对，而是需要进行转换操作。</p>

<p>为了保持和 Redis 的兼容性，Pika 的 Nemo 模块就负责把 Redis 的集合类型转换成单值的键值对。简单来说，我们可以把 Redis 的集合类型分成两类：</p>

<ul>
<li>一类是 List 和 Set 类型，它们的集合中也只有单值；</li>
<li>另一类是 Hash 和 Sorted Set 类型，它们的集合中的元素是成对的，其中，Hash 集合元素是 field-value 类型，而 Sorted Set 集合元素是 member-score 类型。</li>
</ul>

<p>Nemo 模块通过转换操作，把这 4 种集合类型的元素表示为单值的键值对。具体怎么转换呢？下面我们来分别看下每种类型的转换。</p>

<p>首先我们来看 List 类型。在 Pika 中，List 集合的 key 被嵌入到了单值键值对的键当中，用 key 字段表示；而 List 集合的元素值，则被嵌入到单值键值对的值当中，用 value 字段表示。因为 List 集合中的元素是有序的，所以，Nemo 模块还在单值键值对的 key 后面增加了 sequence 字段，表示当前元素在 List 中的顺序，同时，还在 value 的前面增加了 previous sequence 和 next sequence 这两个字段，分别表示当前元素的前一个元素和后一个元素。</p>

<p>此外，在单值键值对的 key 前面，Nemo 模块还增加了一个值“l”，表示当前数据是 List 类型，以及增加了一个 1 字节的 size 字段，表示 List 集合 key 的大小。在单值键值对的 value 后面，Nemo 模块还增加了 version 和 ttl 字段，分别表示当前数据的版本号和剩余存活时间（用来支持过期 key 功能），如下图所示：</p>

<p><img src="assets/066465f1a28b6f14a42c1fc3a3f73105-20221015223826-cb9jtod.jpg" alt="" /></p>

<p>我们再来看看 Set 集合。</p>

<p>Set 集合的 key 和元素 member 值，都被嵌入到了 Pika 单值键值对的键当中，分别用 key 和 member 字段表示。同时，和 List 集合类似，单值键值对的 key 前面有值“s”，用来表示数据是 Set 类型，同时还有 size 字段，用来表示 key 的大小。Pika 单值键值对的值只保存了数据的版本信息和剩余存活时间，如下图所示：</p>

<p><img src="assets/aa20c1456526dbf3f7d30f9d865f0f71-20221015223826-yqm2lrb.jpg" alt="" /></p>

<p>对于 Hash 类型来说，Hash 集合的 key 被嵌入到单值键值对的键当中，用 key 字段表示，而 Hash 集合元素的 field 也被嵌入到单值键值对的键当中，紧接着 key 字段，用 field 字段表示。Hash 集合元素的 value 则是嵌入到单值键值对的值当中，并且也带有版本信息和剩余存活时间，如下图所示：</p>

<p><img src="assets/6378f7045393ae342632189a4ab601b9-20221015223826-12yej5u.jpg" alt="" /></p>

<p>最后，对于 Sorted Set 类型来说，该类型是需要能够按照集合元素的 score 值排序的，而 RocksDB 只支持按照单值键值对的键来排序。所以，Nemo 模块在转换数据时，就把 Sorted Set 集合 key、元素的 score 和 member 值都嵌入到了单值键值对的键当中，此时，单值键值对中的值只保存了数据的版本信息和剩余存活时间，如下图所示：</p>

<p><img src="assets/a0bc4d00a5d95e7fd2699945ff7a56a8-20221015223826-v7wucuf.jpg" alt="" /></p>

<p>采用了上面的转换方式之后，Pika 不仅能兼容支持 Redis 的数据类型，而且还保留了这些数据类型的特征，例如 List 的元素保序、Sorted Set 的元素按 score 排序。了解了 Pika 的转换机制后，你就会明白，如果你有业务应用计划从使用 Redis 切换到使用 Pika，就不用担心面临因为操作接口不兼容而要修改业务应用的问题了。</p>

<p>经过刚刚的分析，我们可以知道，Pika 能够基于 SSD 保存大容量数据，而且和 Redis 兼容，这是它的两个优势。接下来，我们再来看看，跟 Redis 相比，Pika 的其他优势，以及潜在的不足。当在实际应用 Pika 时，Pika 的不足之处是你需要特别注意的地方，这些可能都需要你进行系统配置或参数上的调优。</p>

<h2 id="pika-的其他优势与不足">Pika 的其他优势与不足</h2>

<p>跟 Redis 相比，Pika 最大的特点就是使用了 SSD 来保存数据，这个特点能带来的最直接好处就是，Pika 单实例能保存更多的数据了，实现了实例数据扩容。</p>

<p>除此之外，Pika 使用 SSD 来保存数据，还有额外的两个优势。</p>

<p>首先，<strong>实例重启快</strong>。Pika 的数据在写入数据库时，是会保存到 SSD 上的。当 Pika 实例重启时，可以直接从 SSD 上的数据文件中读取数据，不需要像 Redis 一样，从 RDB 文件全部重新加载数据或是从 AOF 文件中全部回放操作，这极大地提高了 Pika 实例的重启速度，可以快速处理业务应用请求。</p>

<p>另外，主从库重新执行全量同步的风险低。Pika 通过 binlog 机制实现写命令的增量同步，不再受内存缓冲区大小的限制，所以，即使在数据量很大导致主从库同步耗时很长的情况下，Pika 也不用担心缓冲区溢出而触发的主从库重新全量同步。</p>

<p>但是，就像我在前面的课程中和你说的，“硬币都是有正反两面的”，Pika 也有自身的一些不足。</p>

<p>虽然它保持了 Redis 操作接口，也能实现数据库扩容，但是，当把数据保存到 SSD 上后，会降低数据的访问性能。这是因为，数据操作毕竟不能在内存中直接执行了，而是要在底层的 SSD 中进行存取，这肯定会影响，Pika 的性能。而且，我们还需要把 binlog 机制记录的写命令同步到 SSD 上，这会降低 Pika 的写性能。</p>

<p>不过，Pika 的多线程模型，可以同时使用多个线程进行数据读写，这在一定程度上弥补了从 SSD 存取数据造成的性能损失。当然，你也可以使用高配的 SSD 来提升访问性能，进而减少读写 SSD 对 Pika 性能的影响。</p>

<p>为了帮助你更直观地了解 Pika 的性能情况，我再给你提供一张表，这是 Pika<a href="">官网</a>上提供的测试数据。</p>

<p><img src="assets/6fed4a269a79325efd6fa4fb17fc44c5-20221015223826-bpxdg26.jpg" alt="" /></p>

<p>这些数据是在 Pika 3.2 版本中，String 和 Hash 类型在多线程情况下的基本操作性能结果。从表中可以看到，在不写 binlog 时，Pika 的 SET/GET、HSET/HGET 的性能都能达到 200K OPS 以上，而一旦增加了写 binlog 操作，SET 和 HSET 操作性能大约下降了 41%，只有约 120K OPS。</p>

<p>所以，我们在使用 Pika 时，需要在单实例扩容的必要性和可能的性能损失间做个权衡。如果保存大容量数据是我们的首要需求，那么，Pika 是一个不错的解决方案。</p>

<h2 id="小结">小结</h2>

<p>这节课，我们学习了基于 SSD 给 Redis 单实例进行扩容的技术方案 Pika。跟 Redis 相比，Pika 的好处非常明显：既支持 Redis 操作接口，又能支持保存大容量的数据。如果你原来就在应用 Redis，现在想进行扩容，那么，Pika 无疑是一个很好的选择，无论是代码迁移还是运维管理，Pika 基本不需要额外的工作量。</p>

<p>不过，Pika 毕竟是把数据保存到了 SSD 上，数据访问要读写 SSD，所以，读写性能要弱于 Redis。针对这一点，我给你提供两个降低读写 SSD 对 Pika 的性能影响的小建议：</p>

<p>利用 Pika 的多线程模型，增加线程数量，提升 Pika 的并发请求处理能力；</p>

<p>为 Pika 配置高配的 SSD，提升 SSD 自身的访问性能。</p>

<p>最后，我想再给你一个小提示。Pika 本身提供了很多工具，可以帮助我们把 Redis 数据迁移到 Pika，或者是把 Redis 请求转发给 Pika。比如说，我们使用 aof_to_pika 命令，并且指定 Redis 的 AOF 文件以及 Pika 的连接信息，就可以把 Redis 数据迁移到 Pika 中了，如下所示：</p>

<pre><code class="language-java">aof_to_pika -i [Redis AOF文件] -h [Pika IP] -p [Pika port] -a [认证信息]
</code></pre>

<p>关于这些工具的信息，你都可以直接在 Pika 的<a href="https://github.com/Qihoo360/pika/wiki" target="_blank">GitHub</a>上找到。而且，Pika 本身也还在迭代开发中，我也建议你多去看看 GitHub，进一步地了解它。这样，你就可以获得 Pika 的最新进展，也能更好地把它应用到你的业务实践中。</p>

<h2 id="每课一问">每课一问</h2>

<p>按照惯例，我给你提个小问题。这节课，我向你介绍的是使用 SSD 作为内存容量的扩展，增加 Redis 实例的数据保存量，我想请你来聊一聊，我们可以使用机械硬盘来作为实例容量扩展吗，有什么好处或不足吗？</p>

<p>欢迎在留言区写下你的思考和答案，我们一起交流讨论。如果你觉得今天的内容对你有所帮助，也欢迎你分享给你的朋友或同事。我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#afc3c3c3969b9e9e9f98efc8c2cec6c381ccc0c2" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1223222dbccdc2',t:'MTczNDA1Mjg4NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
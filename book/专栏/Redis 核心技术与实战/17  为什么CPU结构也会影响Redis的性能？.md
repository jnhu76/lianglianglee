<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=17&#32;&#32;为什么CPU结构也会影响Redis的性能？>
        <link rel="icon" href="/static/favicon.png">
        <title>17  为什么CPU结构也会影响Redis的性能？ </title>
        
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
                            <h1 id="title" data-id="17  为什么CPU结构也会影响Redis的性能？" class="title">17  为什么CPU结构也会影响Redis的性能？</h1>
                            <div><p>很多人都认为 Redis 和 CPU 的关系很简单，就是 Redis 的线程在 CPU 上运行，CPU 快，Redis 处理请求的速度也很快。</p>

<p>这种认知其实是片面的。CPU 的多核架构以及多 CPU 架构，也会影响到 Redis 的性能。如果不了解 CPU 对 Redis 的影响，在对 Redis 的性能进行调优时，就可能会遗漏一些调优方法，不能把 Redis 的性能发挥到极限。</p>

<p>今天，我们就来学习下目前主流服务器的 CPU 架构，以及基于 CPU 多核架构和多 CPU 架构优化 Redis 性能的方法。</p>

<h2 id="主流的-cpu-架构">主流的 CPU 架构</h2>

<p>要了解 CPU 对 Redis 具体有什么影响，我们得先了解一下 CPU 架构。</p>

<p>一个 CPU 处理器中一般有多个运行核心，我们把一个运行核心称为一个物理核，每个物理核都可以运行应用程序。每个物理核都拥有私有的一级缓存（Level 1 cache，简称 L1 cache），包括一级指令缓存和一级数据缓存，以及私有的二级缓存（Level 2 cache，简称 L2 cache）。</p>

<p>这里提到了一个概念，就是物理核的私有缓存。它其实是指缓存空间只能被当前的这个物理核使用，其他的物理核无法对这个核的缓存空间进行数据存取。我们来看一下 CPU 物理核的架构。</p>

<p><img src="assets/c2d620c012a82e825570df631a7fbc3a-20221015223416-wtqsmp0.jpg" alt="" /></p>

<p>因为 L1 和 L2 缓存是每个物理核私有的，所以，当数据或指令保存在 L1、L2 缓存时，物理核访问它们的延迟不超过 10 纳秒，速度非常快。那么，如果 Redis 把要运行的指令或存取的数据保存在 L1 和 L2 缓存的话，就能高速地访问这些指令和数据。</p>

<p>但是，这些 L1 和 L2 缓存的大小受限于处理器的制造技术，一般只有 KB 级别，存不下太多的数据。如果 L1、L2 缓存中没有所需的数据，应用程序就需要访问内存来获取数据。而应用程序的访存延迟一般在百纳秒级别，是访问 L1、L2 缓存的延迟的近 10 倍，不可避免地会对性能造成影响。</p>

<p>所以，不同的物理核还会共享一个共同的三级缓存（Level 3 cache，简称为 L3 cache）。L3 缓存能够使用的存储资源比较多，所以一般比较大，能达到几 MB 到几十 MB，这就能让应用程序缓存更多的数据。当 L1、L2 缓存中没有数据缓存时，可以访问 L3，尽可能避免访问内存。</p>

<p>另外，现在主流的 CPU 处理器中，每个物理核通常都会运行两个超线程，也叫作逻辑核。同一个物理核的逻辑核会共享使用 L1、L2 缓存。</p>

<p>为了方便你理解，我用一张图展示一下物理核和逻辑核，以及一级、二级缓存的关系。</p>

<p><img src="assets/d9689a38cbe67c3008d8ba99663c2f09-20221015223416-qxjwqsp.jpg" alt="" /></p>

<p>在主流的服务器上，一个 CPU 处理器会有 10 到 20 多个物理核。同时，为了提升服务器的处理能力，服务器上通常还会有多个 CPU 处理器（也称为多 CPU Socket），每个处理器有自己的物理核（包括 L1、L2 缓存），L3 缓存，以及连接的内存，同时，不同处理器间通过总线连接。</p>

<p>下图显示的就是多 CPU Socket 的架构，图中有两个 Socket，每个 Socket 有两个物理核。</p>

<p><img src="assets/5ceb2ab6f61c064284c8f8811431bc3d-20221015223416-kni484v.jpg" alt="" /></p>

<p><strong>在多 CPU 架构上，应用程序可以在不同的处理器上运行</strong>。在刚才的图中，Redis 可以先在 Socket  1 上运行一段时间，然后再被调度到 Socket  2 上运行。</p>

<p>但是，有个地方需要你注意一下：如果应用程序先在一个 Socket 上运行，并且把数据保存到了内存，然后被调度到另一个 Socket 上运行，此时，应用程序再进行内存访问时，就需要访问之前 Socket 上连接的内存，这种访问属于<strong>远端内存访问</strong>。<strong>和访问 Socket 直接连接的内存相比，远端内存访问会增加应用程序的延迟。</strong></p>

<p>在多 CPU 架构下，一个应用程序访问所在 Socket 的本地内存和访问远端内存的延迟并不一致，所以，我们也把这个架构称为非统一内存访问架构（Non-Uniform Memory Access，NUMA 架构）。</p>

<p>到这里，我们就知道了主流的 CPU 多核架构和多 CPU 架构，我们来简单总结下 CPU 架构对应用程序运行的影响。</p>

<ul>
<li>L1、L2 缓存中的指令和数据的访问速度很快，所以，充分利用 L1、L2 缓存，可以有效缩短应用程序的执行时间；</li>
<li>在 NUMA 架构下，如果应用程序从一个 Socket 上调度到另一个 Socket 上，就可能会出现远端内存访问的情况，这会直接增加应用程序的执行时间。</li>
</ul>

<p>接下来，我们就先来了解下 CPU 多核是如何影响 Redis 性能的。</p>

<h2 id="cpu-多核对-redis-性能的影响">CPU 多核对 Redis 性能的影响</h2>

<p>在一个 CPU 核上运行时，应用程序需要记录自身使用的软硬件资源信息（例如栈指针、CPU 核的寄存器值等），我们把这些信息称为运行时信息。同时，应用程序访问最频繁的指令和数据还会被缓存到 L1、L2 缓存上，以便提升执行速度。</p>

<p>但是，在多核 CPU 的场景下，一旦应用程序需要在一个新的 CPU 核上运行，那么，运行时信息就需要重新加载到新的 CPU 核上。而且，新的 CPU 核的 L1、L2 缓存也需要重新加载数据和指令，这会导致程序的运行时间增加。</p>

<p>说到这儿，我想跟你分享一个我曾经在多核 CPU 环境下对 Redis 性能进行调优的案例。希望借助这个案例，帮你全方位地了解到多核 CPU 对 Redis 的性能的影响。</p>

<p>当时，我们的项目需求是要对 Redis 的 99% 尾延迟进行优化，要求 GET 尾延迟小于 300 微秒，PUT 尾延迟小于 500 微秒。</p>

<p>可能有同学不太清楚 99% 尾延迟是啥，我先解释一下。我们把所有请求的处理延迟从小到大排个序，<strong>99% 的请求延迟小于的值就是 99% 尾延迟。</strong>比如说，我们有 1000 个请求，假设按请求延迟从小到大排序后，第 991 个请求的延迟实测值是 1ms，而前 990 个请求的延迟都小于 1ms，所以，这里的 99% 尾延迟就是 1ms。</p>

<p>刚开始的时候，我们使用 GET/PUT 复杂度为 O(1) 的 String 类型进行数据存取，同时关闭了 RDB 和 AOF，而且，Redis 实例中没有保存集合类型的其他数据，也就没有 bigkey 操作，避免了可能导致延迟增加的许多情况。</p>

<p>但是，即使这样，我们在一台有 24 个 CPU 核的服务器上运行 Redis 实例，GET 和 PUT 的 99% 尾延迟分别是 504 微秒和 1175 微秒，明显大于我们设定的目标。</p>

<p>后来，我们仔细检测了 Redis 实例运行时的服务器 CPU 的状态指标值，这才发现，CPU 的 context switch 次数比较多。</p>

<p>context switch 是指线程的上下文切换，这里的上下文就是线程的运行时信息。在 CPU 多核的环境中，一个线程先在一个 CPU 核上运行，之后又切换到另一个 CPU 核上运行，这时就会发生 context switch。</p>

<p>当 context switch 发生后，Redis 主线程的运行时信息需要被重新加载到另一个 CPU 核上，而且，此时，另一个 CPU 核上的 L1、L2 缓存中，并没有 Redis 实例之前运行时频繁访问的指令和数据，所以，这些指令和数据都需要重新从 L3 缓存，甚至是内存中加载。这个重新加载的过程是需要花费一定时间的。而且，Redis 实例需要等待这个重新加载的过程完成后，才能开始处理请求，所以，这也会导致一些请求的处理时间增加。</p>

<p>如果在 CPU 多核场景下，Redis 实例被频繁调度到不同 CPU 核上运行的话，那么，对 Redis 实例的请求处理时间影响就更大了。<strong>每调度一次，一些请求就会受到运行时信息、指令和数据重新加载过程的影响，这就会导致某些请求的延迟明显高于其他请求。</strong>分析到这里，我们就知道了刚刚的例子中 99% 尾延迟的值始终降不下来的原因。</p>

<p>所以，我们要避免 Redis 总是在不同 CPU 核上来回调度执行。于是，我们尝试着把 Redis 实例和 CPU 核绑定了，让一个 Redis 实例固定运行在一个 CPU 核上。我们可以使用** taskset 命令**把一个程序绑定在一个核上运行。</p>

<p>比如说，我们执行下面的命令，就把 Redis 实例绑在了 0 号核上，其中，“-c”选项用于设置要绑定的核编号。</p>

<pre><code>taskset -c 0 ./redis-server
</code></pre>

<p>绑定以后，我们进行了测试。我们发现，Redis 实例的 GET 和 PUT 的 99% 尾延迟一下子就分别降到了 260 微秒和 482 微秒，达到了我们期望的目标。</p>

<p>我们来看一下绑核前后的 Redis 的 99% 尾延迟。</p>

<p><img src="assets/eb72b9f58052d6a6023d3e1dac522157-20221015223416-itmf8u6.jpg" alt="" /></p>

<p>可以看到，在 CPU 多核的环境下，通过绑定 Redis 实例和 CPU 核，可以有效降低 Redis 的尾延迟。当然，绑核不仅对降低尾延迟有好处，同样也能降低平均延迟、提升吞吐率，进而提升 Redis 性能。</p>

<p>接下来，我们再来看看多 CPU 架构，也就是 NUMA 架构，对 Redis 性能的影响。</p>

<h2 id="cpu-的-numa-架构对-redis-性能的影响">CPU 的 NUMA 架构对 Redis 性能的影响</h2>

<p>在实际应用 Redis 时，我经常看到一种做法，为了提升 Redis 的网络性能，把操作系统的网络中断处理程序和 CPU 核绑定。这个做法可以避免网络中断处理程序在不同核上来回调度执行，的确能有效提升 Redis 的网络处理性能。</p>

<p>但是，网络中断程序是要和 Redis 实例进行网络数据交互的，一旦把网络中断程序绑核后，我们就需要注意 Redis 实例是绑在哪个核上了，这会关系到 Redis 访问网络数据的效率高低。</p>

<p>我们先来看下 Redis 实例和网络中断程序的数据交互：网络中断处理程序从网卡硬件中读取数据，并把数据写入到操作系统内核维护的一块内存缓冲区。内核会通过 epoll 机制触发事件，通知 Redis 实例，Redis 实例再把数据从内核的内存缓冲区拷贝到自己的内存空间，如下图所示：</p>

<p><img src="assets/8753ce6985fd08bb9cf9a3813c8b2cd2-20221015223416-e5ld65n.jpg" alt="" /></p>

<p>那么，在 CPU 的 NUMA 架构下，当网络中断处理程序、Redis 实例分别和 CPU 核绑定后，就会有一个潜在的风险：<strong>如果网络中断处理程序和 Redis 实例各自所绑的 CPU 核不在同一个 CPU Socket 上，那么，Redis 实例读取网络数据时，就需要跨 CPU Socket 访问内存，这个过程会花费较多时间。</strong></p>

<p>这么说可能有点抽象，我再借助一张图来解释下。</p>

<p><img src="assets/30cd42yy86debc0eb6e7c5b069533ab0-20221015223416-fpu2z51.jpg" alt="" /></p>

<p>可以看到，图中的网络中断处理程序被绑在了 CPU Socket 1 的某个核上，而 Redis 实例则被绑在了 CPU Socket  2 上。此时，网络中断处理程序读取到的网络数据，被保存在 CPU Socket  1 的本地内存中，当 Redis 实例要访问网络数据时，就需要 Socket 2 通过总线把内存访问命令发送到 Socket 1 上，进行远程访问，时间开销比较大。</p>

<p>我们曾经做过测试，和访问 CPU Socket 本地内存相比，跨 CPU Socket 的内存访问延迟增加了 18%，这自然会导致 Redis 处理请求的延迟增加。</p>

<p>所以，为了避免 Redis 跨 CPU Socket 访问网络数据，我们最好把网络中断程序和 Redis 实例绑在同一个 CPU Socket 上，这样一来，Redis 实例就可以直接从本地内存读取网络数据了，如下图所示：</p>

<p><img src="assets/41f02b2afb08ec54249680e8cac30179-20221015223416-y165yv4.jpg" alt="" /></p>

<p>不过，需要注意的是，<strong>在 CPU 的 NUMA 架构下，对 CPU 核的编号规则，并不是先把一个 CPU Socket 中的所有逻辑核编完，再对下一个 CPU Socket 中的逻辑核编码，而是先给每个 CPU Socket 中每个物理核的第一个逻辑核依次编号，再给每个 CPU Socket 中的物理核的第二个逻辑核依次编号。</strong></p>

<p>我给你举个例子。假设有 2 个 CPU Socket，每个 Socket 上有 6 个物理核，每个物理核又有 2 个逻辑核，总共 24 个逻辑核。我们可以执行 <strong>lscpu 命令</strong>，查看到这些核的编号：</p>

<pre><code>lscpu

Architecture: x86_64
...
NUMA node0 CPU(s): 0-5,12-17
NUMA node1 CPU(s): 6-11,18-23
...
</code></pre>

<p>可以看到，NUMA node0 的 CPU 核编号是 0 到 5、12 到 17。其中，0 到 5 是 node0 上的 6 个物理核中的第一个逻辑核的编号，12 到 17 是相应物理核中的第二个逻辑核编号。NUMA node1 的 CPU 核编号规则和 node0 一样。</p>

<p>所以，在绑核时，我们一定要注意，不能想当然地认为第一个 Socket 上的 12 个逻辑核的编号就是 0 到 11。否则，网络中断程序和 Redis 实例就可能绑在了不同的 CPU Socket 上。</p>

<p>比如说，如果我们把网络中断程序和 Redis 实例分别绑到编号为 1 和 7 的 CPU 核上，此时，它们仍然是在 2 个 CPU Socket 上，Redis 实例仍然需要跨 Socket 读取网络数据。</p>

<p><strong>所以，你一定要注意 NUMA 架构下 CPU 核的编号方法，这样才不会绑错核。</strong></p>

<p>我们先简单地总结下刚刚学习的内容。在 CPU 多核的场景下，用 taskset 命令把 Redis 实例和一个核绑定，可以减少 Redis 实例在不同核上被来回调度执行的开销，避免较高的尾延迟；在多 CPU 的 NUMA 架构下，如果你对网络中断程序做了绑核操作，建议你同时把 Redis 实例和网络中断程序绑在同一个 CPU Socket 的不同核上，这样可以避免 Redis 跨 Socket 访问内存中的网络数据的时间开销。</p>

<p>不过，“硬币都是有两面的”，绑核也存在一定的风险。接下来，我们就来了解下它的潜在风险点和解决方案。</p>

<h2 id="绑核的风险和解决方案">绑核的风险和解决方案</h2>

<p>Redis 除了主线程以外，还有用于 RDB 生成和 AOF 重写的子进程（可以回顾看下【第 4 讲】和【第 5 讲】）。此外，我们还在【第 16 讲】学习了 Redis 的后台线程。</p>

<p>当我们把 Redis 实例绑到一个 CPU 逻辑核上时，就会导致子进程、后台线程和 Redis 主线程竞争 CPU 资源，一旦子进程或后台线程占用 CPU 时，主线程就会被阻塞，导致 Redis 请求延迟增加。</p>

<p>针对这种情况，我来给你介绍两种解决方案，分别是<strong>一个 Redis 实例对应绑一个物理核和优化 Redis 源码。</strong></p>

<p><strong>方案一：一个 Redis 实例对应绑一个物理核</strong></p>

<p>在给 Redis 实例绑核时，我们不要把一个实例和一个逻辑核绑定，而要和一个物理核绑定，也就是说，把一个物理核的 2 个逻辑核都用上。</p>

<p>我们还是以刚才的 NUMA 架构为例，NUMA node0 的 CPU 核编号是 0 到 5、12 到 17。其中，编号 0 和 12、1 和 13、2 和 14 等都是表示一个物理核的 2 个逻辑核。所以，在绑核时，我们使用属于同一个物理核的 2 个逻辑核进行绑核操作。例如，我们执行下面的命令，就把 Redis 实例绑定到了逻辑核 0 和 12 上，而这两个核正好都属于物理核 1。</p>

<pre><code>taskset -c 0,12 ./redis-server
</code></pre>

<p>和只绑一个逻辑核相比，把 Redis 实例和物理核绑定，可以让主线程、子进程、后台线程共享使用 2 个逻辑核，可以在一定程度上缓解 CPU 资源竞争。但是，因为只用了 2 个逻辑核，它们相互之间的 CPU 竞争仍然还会存在。如果你还想进一步减少 CPU 竞争，我再给你介绍一种方案。</p>

<p><strong>方案二：优化 Redis 源码</strong></p>

<p>这个方案就是通过修改 Redis 源码，把子进程和后台线程绑到不同的 CPU 核上。</p>

<p>如果你对 Redis 的源码不太熟悉，也没关系，因为这是通过编程实现绑核的一个通用做法。学会了这个方案，你可以在熟悉了源码之后把它用上，也可以应用在其他需要绑核的场景中。</p>

<p>接下来，我先介绍一下通用的做法，然后，再具体说说可以把这个做法对应到 Redis 的哪部分源码中。</p>

<p>通过编程实现绑核时，要用到操作系统提供的 1 个数据结构 cpu_set_t 和 3 个函数 CPU_ZERO、CPU_SET 和 sched_setaffinity，我先来解释下它们。</p>

<ul>
<li>cpu_set_t 数据结构：是一个位图，每一位用来表示服务器上的一个 CPU 逻辑核。</li>
<li>CPU_ZERO 函数：以 cpu_set_t 结构的位图为输入参数，把位图中所有的位设置为 0。</li>
<li>CPU_SET 函数：以 CPU 逻辑核编号和 cpu_set_t 位图为参数，把位图中和输入的逻辑核编号对应的位设置为 1。</li>
<li>sched_setaffinity 函数：以进程 / 线程 ID 号和 cpu_set_t 为参数，检查 cpu_set_t 中哪一位为 1，就把输入的 ID 号所代表的进程 / 线程绑在对应的逻辑核上。</li>
</ul>

<p>那么，怎么在编程时把这三个函数结合起来实现绑核呢？很简单，我们分四步走就行。</p>

<ul>
<li>第一步：创建一个 cpu_set_t 结构的位图变量；</li>
<li>第二步：使用 CPU_ZERO 函数，把 cpu_set_t 结构的位图所有的位都设置为 0；</li>
<li>第三步：根据要绑定的逻辑核编号，使用 CPU_SET 函数，把 cpu_set_t 结构的位图相应位设置为 1；</li>
<li>第四步：使用 sched_setaffinity 函数，把程序绑定在 cpu_set_t 结构位图中为 1 的逻辑核上。</li>
</ul>

<p>下面，我就具体介绍下，分别把后台线程、子进程绑到不同的核上的做法。</p>

<p>先说后台线程。为了让你更好地理解编程实现绑核，你可以看下这段示例代码，它实现了为线程绑核的操作：</p>

<pre><code class="language-c">//线程函数
void worker(int bind_cpu){
    cpu_set_t cpuset;  //创建位图变量
    CPU_ZERO(&amp;cpu_set); //位图变量所有位设置0
    CPU_SET(bind_cpu, &amp;cpuset); //根据输入的bind_cpu编号，把位图对应为设置为1
    sched_setaffinity(0, sizeof(cpuset), &amp;cpuset); //把程序绑定在cpu_set_t结构位图中为1的逻辑核

    //实际线程函数工作
}

int main(){
    pthread_t pthread1
    //把创建的pthread1绑在编号为3的逻辑核上
    pthread_create(&amp;pthread1, NULL, (void *)worker, 3);
}
</code></pre>

<p>对于 Redis 来说，它是在 bio.c 文件中的 bioProcessBackgroundJobs 函数中创建了后台线程。bioProcessBackgroundJobs 函数类似于刚刚的例子中的 worker 函数，在这个函数中实现绑核四步操作，就可以把后台线程绑到和主线程不同的核上了。</p>

<p>和给线程绑核类似，当我们使用 fork 创建子进程时，也可以把刚刚说的四步操作实现在 fork 后的子进程代码中，示例代码如下：</p>

<pre><code class="language-c">int main(){
   //用fork创建一个子进程
   pid_t p = fork();
   if(p &lt; 0){
      printf(&quot; fork error\n&quot;);
   }
   //子进程代码部分
   else if(!p){
      cpu_set_t cpuset;  //创建位图变量
      CPU_ZERO(&amp;cpu_set); //位图变量所有位设置0
      CPU_SET(3, &amp;cpuset); //把位图的第3位设置为1
      sched_setaffinity(0, sizeof(cpuset), &amp;cpuset);  //把程序绑定在3号逻辑核
      //实际子进程工作
      exit(0);
   }
   ...
}
</code></pre>

<p>对于 Redis 来说，生成 RDB 和 AOF 日志重写的子进程分别是下面两个文件的函数中实现的。</p>

<ul>
<li>rdb.c 文件：rdbSaveBackground 函数；</li>
<li>aof.c 文件：rewriteAppendOnlyFileBackground 函数。</li>
</ul>

<p>这两个函数中都调用了 fork 创建子进程，所以，我们可以在子进程代码部分加上绑核的四步操作。</p>

<p>使用源码优化方案，我们既可以实现 Redis 实例绑核，避免切换核带来的性能影响，还可以让子进程、后台线程和主线程不在同一个核上运行，避免了它们之间的 CPU 资源竞争。相比使用 taskset 绑核来说，这个方案可以进一步降低绑核的风险。</p>

<h2 id="小结">小结</h2>

<p>这节课，我们学习了 CPU 架构对 Redis 性能的影响。首先，我们了解了目前主流的多核 CPU 架构，以及 NUMA 架构。</p>

<p>在多核 CPU 架构下，Redis 如果在不同的核上运行，就需要频繁地进行上下文切换，这个过程会增加 Redis 的执行时间，客户端也会观察到较高的尾延迟了。所以，建议你在 Redis 运行时，把实例和某个核绑定，这样，就能重复利用核上的 L1、L2 缓存，可以降低响应延迟。</p>

<p>为了提升 Redis 的网络性能，我们有时还会把网络中断处理程序和 CPU 核绑定。在这种情况下，如果服务器使用的是 NUMA 架构，Redis 实例一旦被调度到和中断处理程序不在同一个 CPU Socket，就要跨 CPU Socket 访问网络数据，这就会降低 Redis 的性能。所以，我建议你把 Redis 实例和网络中断处理程序绑在同一个 CPU Socket 下的不同核上，这样可以提升 Redis 的运行性能。</p>

<p>虽然绑核可以帮助 Redis 降低请求执行时间，但是，除了主线程，Redis 还有用于 RDB 和 AOF 重写的子进程，以及 4.0 版本之后提供的用于惰性删除的后台线程。当 Redis 实例和一个逻辑核绑定后，这些子进程和后台线程会和主线程竞争 CPU 资源，也会对 Redis 性能造成影响。所以，我给了你两个建议：</p>

<ul>
<li>如果你不想修改 Redis 代码，可以把按一个 Redis 实例一个物理核方式进行绑定，这样，Redis 的主线程、子进程和后台线程可以共享使用一个物理核上的两个逻辑核。</li>
<li>如果你很熟悉 Redis 的源码，就可以在源码中增加绑核操作，把子进程和后台线程绑到不同的核上，这样可以避免对主线程的 CPU 资源竞争。不过，如果你不熟悉 Redis 源码，也不用太担心，Redis 6.0 出来后，可以支持 CPU 核绑定的配置操作了，我将在第 38 讲中向你介绍 Redis 6.0 的最新特性。</li>
</ul>

<p>Redis 的低延迟是我们永恒的追求目标，而多核 CPU 和 NUMA 架构已经成为了目前服务器的主流配置，所以，希望你能掌握绑核优化方案，并把它应用到实践中。</p>

<h2 id="每课一问">每课一问</h2>

<p>按照惯例，我给你提个小问题。</p>

<p>在一台有 2 个 CPU Socket（每个 Socket 8 个物理核）的服务器上，我们部署了有 8 个实例的 Redis 切片集群（8 个实例都为主节点，没有主备关系），现在有两个方案：</p>

<ol>
<li>在同一个 CPU Socket 上运行 8 个实例，并和 8 个 CPU 核绑定；</li>
<li>在 2 个 CPU Socket 上各运行 4 个实例，并和相应 Socket 上的核绑定。</li>
</ol>

<p>如果不考虑网络数据读取的影响，你会选择哪个方案呢？</p>

<p>欢迎在留言区写下你的思考和答案，如果你觉得有所收获，也欢迎你帮我把今天的内容分享给你的朋友。我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#016d6d6d38353030313641666c60686d2f626e6c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f122064ef7acdc2',t:'MTczNDA1Mjc3Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
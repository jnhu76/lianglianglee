<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=15&#32;&#32;消息队列的考验：Redis有哪些解决方案？>
        <link rel="icon" href="/static/favicon.png">
        <title>15  消息队列的考验：Redis有哪些解决方案？ </title>
        
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
                            <h1 id="title" data-id="15  消息队列的考验：Redis有哪些解决方案？" class="title">15  消息队列的考验：Redis有哪些解决方案？</h1>
                            <div><p>现在的互联网应用基本上都是采用分布式系统架构进行设计的，而很多分布式系统必备的一个基础软件就是消息队列。</p>

<p>消息队列要能支持组件通信消息的快速读写，而 Redis 本身支持数据的高速访问，正好可以满足消息队列的读写性能需求。不过，除了性能，消息队列还有其他的要求，所以，很多人都很关心一个问题：“Redis 适合做消息队列吗？”</p>

<p>其实，这个问题的背后，隐含着两方面的核心问题：</p>

<ul>
<li>消息队列的消息存取需求是什么？</li>
<li>Redis 如何实现消息队列的需求？</li>
</ul>

<p>这节课，我们就来聊一聊消息队列的特征和 Redis 提供的消息队列方案。只有把这两方面的知识和实践经验串连起来，才能彻底理解基于 Redis 实现消息队列的技术实践。以后当你需要为分布式系统组件做消息队列选型时，就可以根据组件通信量和消息通信速度的要求，选择出适合的 Redis 消息队列方案了。</p>

<p>我们先来看下第一个问题：消息队列的消息读取有什么样的需求？</p>

<h2 id="消息队列的消息存取需求">消息队列的消息存取需求</h2>

<p>我先介绍一下消息队列存取消息的过程。在分布式系统中，当两个组件要基于消息队列进行通信时，一个组件会把要处理的数据以消息的形式传递给消息队列，然后，这个组件就可以继续执行其他操作了；远端的另一个组件从消息队列中把消息读取出来，再在本地进行处理。</p>

<p>为了方便你理解，我还是借助一个例子来解释一下。</p>

<p>假设组件 1 需要对采集到的数据进行求和计算，并写入数据库，但是，消息到达的速度很快，组件 1 没有办法及时地既做采集，又做计算，并且写入数据库。所以，我们可以使用基于消息队列的通信，让组件 1 把数据 x 和 y 保存为 JSON 格式的消息，再发到消息队列，这样它就可以继续接收新的数据了。组件 2 则异步地从消息队列中把数据读取出来，在服务器 2 上进行求和计算后，再写入数据库。这个过程如下图所示：</p>

<p><img src="assets/d79d46ec4aa22bf46fde3ae1a99fc2bc-20221015223336-w1nto6h.jpg" alt="" /></p>

<p>我们一般把消息队列中发送消息的组件称为生产者（例子中的组件 1），把接收消息的组件称为消费者（例子中的组件 2），下图展示了一个通用的消息队列的架构模型：</p>

<p><img src="assets/f470bb957c1faff674c08b1fa65a3a62-20221015223336-61kx1hk.jpg" alt="" /></p>

<p>在使用消息队列时，消费者可以异步读取生产者消息，然后再进行处理。这样一来，即使生产者发送消息的速度远远超过了消费者处理消息的速度，生产者已经发送的消息也可以缓存在消息队列中，避免阻塞生产者，这是消息队列作为分布式组件通信的一大优势。</p>

<p><strong>不过，消息队列在存取消息时，必须要满足三个需求，分别是消息保序、处理重复的消息和保证消息可靠性。</strong>分别是消息保序、处理重复的消息和保证消息可靠性。</p>

<h3 id="需求一-消息保序">需求一：消息保序</h3>

<p>虽然消费者是异步处理消息，但是，消费者仍然需要按照生产者发送消息的顺序来处理消息，避免后发送的消息被先处理了。对于要求消息保序的场景来说，一旦出现这种消息被乱序处理的情况，就可能会导致业务逻辑被错误执行，从而给业务方造成损失。</p>

<p>我们来看一个更新商品库存的场景。</p>

<p>假设生产者负责接收库存更新请求，消费者负责实际更新库存，现有库存量是 10。生产者先后发送了消息 1 和消息 2，消息 1 要把商品 X 的库存记录更新为 5，消息 2 是把商品 X 库存更新为 3。如果消息 1 和 2 在消息队列中无法保序，出现消息 2 早于消息 1 被处理的情况，那么，很显然，库存更新就出错了。这是业务应用无法接受的。</p>

<p>面对这种情况，你可能会想到一种解决方案：不要把更新后的库存量作为生产者发送的消息，而是<strong>把库存扣除值作为消息的内容</strong>。这样一来，消息 1 是扣减库存量 5，消息 2 是扣减库存量 2。如果消息 1 和消息 2 之间没有库存查询请求的话，即使消费者先处理消息 2，再处理消息 1，这个方案也能够保证最终的库存量是正确的，也就是库存量为 3。</p>

<p>但是，我们还需要考虑这样一种情况：假如消费者收到了这样三条消息：消息 1 是扣减库存量 5，消息 2 是读取库存量，消息 3 是扣减库存量 2，此时，如果消费者先处理了消息 3（把库存量扣减 2），那么库存量就变成了 8。然后，消费者处理了消息 2，读取当前的库存量是 8，这就会出现库存量查询不正确的情况。从业务应用层面看，消息 1、2、3 应该是顺序执行的，所以，消息 2 查询到的应该是扣减了 5 以后的库存量，而不是扣减了 2 以后的库存量。所以，用库存扣除值作为消息的方案，在消息中同时包含读写操作的场景下，会带来数据读取错误的问题。而且，这个方案还会面临一个问题，那就是重复消息处理。</p>

<h3 id="需求二-重复消息处理">需求二：重复消息处理</h3>

<p>消费者从消息队列读取消息时，有时会因为网络堵塞而出现消息重传的情况。此时，消费者可能会收到多条重复的消息。对于重复的消息，消费者如果多次处理的话，就可能造成一个业务逻辑被多次执行，如果业务逻辑正好是要修改数据，那就会出现数据被多次修改的问题了。</p>

<p>还是以库存更新为例，假设消费者收到了一次消息 1，要扣减库存量 5，然后又收到了一次消息 1，那么，如果消费者无法识别这两条消息实际是一条相同消息的话，就会执行两次扣减库存量 5 的操作，此时，库存量就不对了。这当然也是无法接受的。</p>

<h3 id="需求三-消息可靠性保证">需求三：消息可靠性保证</h3>

<p>另外，消费者在处理消息的时候，还可能出现因为故障或宕机导致消息没有处理完成的情况。此时，消息队列需要能提供消息可靠性的保证，也就是说，当消费者重启后，可以重新读取消息再次进行处理，否则，就会出现消息漏处理的问题了。</p>

<p>Redis 的 List 和 Streams 两种数据类型，就可以满足消息队列的这三个需求。我们先来了解下基于 List 的消息队列实现方法。</p>

<h2 id="基于-list-的消息队列解决方案">基于 List 的消息队列解决方案</h2>

<p>List 本身就是按先进先出的顺序对数据进行存取的，所以，如果使用 List 作为消息队列保存消息的话，就已经能满足消息保序的需求了。</p>

<p>具体来说，生产者可以使用 LPUSH 命令把要发送的消息依次写入 List，而消费者则可以使用 RPOP 命令，从 List 的另一端按照消息的写入顺序，依次读取消息并进行处理。</p>

<p>如下图所示，生产者先用 LPUSH 写入了两条库存消息，分别是 5 和 3，表示要把库存更新为 5 和 3；消费者则用 RPOP 把两条消息依次读出，然后进行相应的处理。</p>

<p><img src="assets/b0959216cbce7ac383ce206b8884777c-20221015223336-1f7blty.jpg" alt="" /></p>

<p>不过，在消费者读取数据时，有一个潜在的性能风险点。</p>

<p>在生产者往 List 中写入数据时，List 并不会主动地通知消费者有新消息写入，如果消费者想要及时处理消息，就需要在程序中不停地调用 RPOP 命令（比如使用一个 while(1) 循环）。如果有新消息写入，RPOP 命令就会返回结果，否则，RPOP 命令返回空值，再继续循环。</p>

<p>所以，即使没有新消息写入 List，消费者也要不停地调用 RPOP 命令，这就会导致消费者程序的 CPU 一直消耗在执行 RPOP 命令上，带来不必要的性能损失。</p>

<p>为了解决这个问题，Redis 提供了 BRPOP 命令。<strong>BRPOP 命令也称为阻塞式读取，客户端在没有读到队列数据时，自动阻塞，直到有新的数据写入队列，再开始读取新数据。</strong>和消费者程序自己不停地调用 RPOP 命令相比，这种方式能节省 CPU 开销。</p>

<p>消息保序的问题解决了，接下来，我们还需要考虑解决重复消息处理的问题，这里其实有一个要求：<strong>消费者程序本身能对重复消息进行判断。</strong></p>

<p>一方面，消息队列要能给每一个消息提供全局唯一的 ID 号；另一方面，消费者程序要把已经处理过的消息的 ID 号记录下来。</p>

<p>当收到一条消息后，消费者程序就可以对比收到的消息 ID 和记录的已处理过的消息 ID，来判断当前收到的消息有没有经过处理。如果已经处理过，那么，消费者程序就不再进行处理了。这种处理特性也称为幂等性，幂等性就是指，对于同一条消息，消费者收到一次的处理结果和收到多次的处理结果是一致的。</p>

<p>不过，List 本身是不会为每个消息生成 ID 号的，所以，消息的全局唯一 ID 号就需要生产者程序在发送消息前自行生成。生成之后，我们在用 LPUSH 命令把消息插入 List 时，需要在消息中包含这个全局唯一 ID。</p>

<p>例如，我们执行以下命令，就把一条全局 ID 为 101030001、库存量为 5 的消息插入了消息队列：</p>

<pre><code>LPUSH mq &quot;101030001:stock:5&quot;
(integer) 1
</code></pre>

<p>最后，我们再来看下，List 类型是如何保证消息可靠性的。</p>

<p>当消费者程序从 List 中读取一条消息后，List 就不会再留存这条消息了。所以，如果消费者程序在处理消息的过程出现了故障或宕机，就会导致消息没有处理完成，那么，消费者程序再次启动后，就没法再次从 List 中读取消息了。</p>

<p>为了留存消息，List 类型提供了 BRPOPLPUSH 命令，这个命令的作用是让消费者程序从一个 List 中读取消息，同时，Redis 会把这个消息再插入到另一个 List（可以叫作备份 List）留存。这样一来，如果消费者程序读了消息但没能正常处理，等它重启后，就可以从备份 List 中重新读取消息并进行处理了。</p>

<p>我画了一张示意图，展示了使用 BRPOPLPUSH 命令留存消息，以及消费者再次读取消息的过程，你可以看下。</p>

<p><img src="assets/5045395da08317b546aab7eb698d013d-20221015223336-apywldy.jpg" alt="" /></p>

<p>生产者先用 LPUSH 把消息“5”“3”插入到消息队列 mq 中。消费者程序使用 BRPOPLPUSH 命令读取消息“5”，同时，消息“5”还会被 Redis 插入到 mqback 队列中。如果消费者程序处理消息“5”时宕机了，等它重启后，可以从 mqback 中再次读取消息“5”，继续处理。</p>

<p>好了，到这里，你可以看到，基于 List 类型，我们可以满足分布式组件对消息队列的三大需求。但是，在用 List 做消息队列时，我们还可能遇到过一个问题：<strong>生产者消息发送很快，而消费者处理消息的速度比较慢，这就导致 List 中的消息越积越多，给 Redis 的内存带来很大压力。</strong></p>

<p>这个时候，我们希望启动多个消费者程序组成一个消费组，一起分担处理 List 中的消息。但是，List 类型并不支持消费组的实现。那么，还有没有更合适的解决方案呢？这就要说到 Redis 从 5.0 版本开始提供的 Streams 数据类型了。</p>

<p>和 List 相比，Streams 同样能够满足消息队列的三大需求。而且，它还支持消费组形式的消息读取。接下来，我们就来了解下 Streams 的使用方法。</p>

<h2 id="基于-streams-的消息队列解决方案">基于 Streams 的消息队列解决方案</h2>

<p>Streams 是 Redis 专门为消息队列设计的数据类型，它提供了丰富的消息队列操作命令。</p>

<ul>
<li>XADD：插入消息，保证有序，可以自动生成全局唯一 ID；</li>
<li>XREAD：用于读取消息，可以按 ID 读取数据；</li>
<li>XREADGROUP：按消费组形式读取消息；</li>
<li>XPENDING 和 XACK：XPENDING 命令可以用来查询每个消费组内所有消费者已读取但尚未确认的消息，而 XACK 命令用于向消息队列确认消息处理已完成。</li>
</ul>

<p>首先，我们来学习下 Streams 类型存取消息的操作 XADD。</p>

<p>XADD 命令可以往消息队列中插入新消息，消息的格式是键 - 值对形式。对于插入的每一条消息，Streams 可以自动为其生成一个全局唯一的 ID。</p>

<p>比如说，我们执行下面的命令，就可以往名称为 mqstream 的消息队列中插入一条消息，消息的键是 repo，值是 5。其中，消息队列名称后面的*，表示让 Redis 为插入的数据自动生成一个全局唯一的 ID，例如“1599203861727-0”。当然，我们也可以不用*，直接在消息队列名称后自行设定一个 ID 号，只要保证这个 ID 号是全局唯一的就行。不过，相比自行设定 ID 号，使用*会更加方便高效。</p>

<pre><code>XADD mqstream * repo 5
&quot;1599203861727-0&quot;
</code></pre>

<p>可以看到，消息的全局唯一 ID 由两部分组成，第一部分“1599203861727”是数据插入时，以毫秒为单位计算的当前服务器时间，第二部分表示插入消息在当前毫秒内的消息序号，这是从 0 开始编号的。例如，“1599203861727-0”就表示在“1599203861727”毫秒内的第 1 条消息。</p>

<p>当消费者需要读取消息时，可以直接使用 XREAD 命令从消息队列中读取。</p>

<p>XREAD 在读取消息时，可以指定一个消息 ID，并从这个消息 ID 的下一条消息开始进行读取。</p>

<p>例如，我们可以执行下面的命令，从 ID 号为 1599203861727-0 的消息开始，读取后续的所有消息（示例中一共 3 条）。</p>

<pre><code>XREAD BLOCK 100 STREAMS  mqstream 1599203861727-0
1) 1) &quot;mqstream&quot;
   2) 1) 1) &quot;1599274912765-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;3&quot;
      2) 1) &quot;1599274925823-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;2&quot;
      3) 1) &quot;1599274927910-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;1&quot;
</code></pre>

<p>另外，消费者也可以在调用 XRAED 时设定 block 配置项，实现类似于 BRPOP 的阻塞读取操作。当消息队列中没有消息时，一旦设置了 block 配置项，XREAD 就会阻塞，阻塞的时长可以在 block 配置项进行设置。</p>

<p>举个例子，我们来看一下下面的命令，其中，命令最后的“$”符号表示读取最新的消息，同时，我们设置了 block 10000 的配置项，10000 的单位是毫秒，表明 XREAD 在读取最新消息时，如果没有消息到来，XREAD 将阻塞 10000 毫秒（即 10 秒），然后再返回。下面命令中的 XREAD 执行后，消息队列 mqstream 中一直没有消息，所以，XREAD 在 10 秒后返回空值（nil）。</p>

<pre><code>XREAD block 10000 streams mqstream $
(nil)
(10.00s)
</code></pre>

<p>刚刚讲到的这些操作是 List 也支持的，接下来，我们再来学习下 Streams 特有的功能。</p>

<p>Streams 本身可以使用 XGROUP 创建消费组，创建消费组之后，Streams 可以使用 XREADGROUP 命令让消费组内的消费者读取消息，</p>

<p>例如，我们执行下面的命令，创建一个名为 group1 的消费组，这个消费组消费的消息队列是 mqstream。</p>

<pre><code>XGROUP create mqstream group1 0
OK
</code></pre>

<p>然后，我们再执行一段命令，让 group1 消费组里的消费者 consumer1 从 mqstream 中读取所有消息，其中，命令最后的参数“&gt;”，表示从第一条尚未被消费的消息开始读取。因为在 consumer1 读取消息前，group1 中没有其他消费者读取过消息，所以，consumer1 就得到 mqstream 消息队列中的所有消息了（一共 4 条）。</p>

<pre><code>XREADGROUP group group1 consumer1 streams mqstream &gt;
1) 1) &quot;mqstream&quot;
   2) 1) 1) &quot;1599203861727-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;5&quot;
      2) 1) &quot;1599274912765-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;3&quot;
      3) 1) &quot;1599274925823-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;2&quot;
      4) 1) &quot;1599274927910-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;1&quot;
</code></pre>

<p>需要注意的是，消息队列中的消息一旦被消费组里的一个消费者读取了，就不能再被该消费组内的其他消费者读取了。比如说，我们执行完刚才的 XREADGROUP 命令后，再执行下面的命令，让 group1 内的 consumer2 读取消息时，consumer2 读到的就是空值，因为消息已经被 consumer1 读取完了，如下所示：</p>

<pre><code>XREADGROUP group group1 consumer2  streams mqstream 0
1) 1) &quot;mqstream&quot;
   2) (empty list or set)
</code></pre>

<p>使用消费组的目的是让组内的多个消费者共同分担读取消息，所以，我们通常会让每个消费者读取部分消息，从而实现消息读取负载在多个消费者间是均衡分布的。例如，我们执行下列命令，让 group2 中的 consumer1、2、3 各自读取一条消息。</p>

<pre><code>XREADGROUP group group2 consumer1 count 1 streams mqstream &gt;
1) 1) &quot;mqstream&quot;
   2) 1) 1) &quot;1599203861727-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;5&quot;

XREADGROUP group group2 consumer2 count 1 streams mqstream &gt;
1) 1) &quot;mqstream&quot;
   2) 1) 1) &quot;1599274912765-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;3&quot;

XREADGROUP group group2 consumer3 count 1 streams mqstream &gt;
1) 1) &quot;mqstream&quot;
   2) 1) 1) &quot;1599274925823-0&quot;
         2) 1) &quot;repo&quot;
            2) &quot;2&quot;
</code></pre>

<p>为了保证消费者在发生故障或宕机再次重启后，仍然可以读取未处理完的消息，Streams 会自动使用内部队列（也称为 PENDING List）留存消费组里每个消费者读取的消息，直到消费者使用 XACK 命令通知 Streams“消息已经处理完成”。如果消费者没有成功处理消息，它就不会给 Streams 发送 XACK 命令，消息仍然会留存。此时，消费者可以在重启后，用 XPENDING 命令查看已读取、但尚未确认处理完成的消息。</p>

<p>例如，我们来查看一下 group2 中各个消费者已读取、但尚未确认的消息个数。其中，XPENDING 返回结果的第二、三行分别表示 group2 中所有消费者读取的消息最小 ID 和最大 ID。</p>

<pre><code>XPENDING mqstream group2
1) (integer) 3
2) &quot;1599203861727-0&quot;
3) &quot;1599274925823-0&quot;
4) 1) 1) &quot;consumer1&quot;
      2) &quot;1&quot;
   2) 1) &quot;consumer2&quot;
      2) &quot;1&quot;
   3) 1) &quot;consumer3&quot;
      2) &quot;1&quot;
</code></pre>

<p>如果我们还需要进一步查看某个消费者具体读取了哪些数据，可以执行下面的命令：</p>

<pre><code>XPENDING mqstream group2 - + 10 consumer2
1) 1) &quot;1599274912765-0&quot;
   2) &quot;consumer2&quot;
   3) (integer) 513336
   4) (integer) 1
</code></pre>

<p>可以看到，consumer2 已读取的消息的 ID 是 1599274912765-0。</p>

<p>一旦消息 1599274912765-0 被 consumer2 处理了，consumer2 就可以使用 XACK 命令通知 Streams，然后这条消息就会被删除。当我们再使用 XPENDING 命令查看时，就可以看到，consumer2 已经没有已读取、但尚未确认处理的消息了。</p>

<pre><code>XACK mqstream group2 1599274912765-0
(integer) 1
XPENDING mqstream group2 - + 10 consumer2
(empty list or set)
</code></pre>

<p>现在，我们就知道了用 Streams 实现消息队列的方法，我还想再强调下，Streams 是 Redis 5.0 专门针对消息队列场景设计的数据类型，如果你的 Redis 是 5.0 及 5.0 以后的版本，就可以考虑把 Streams 用作消息队列了。</p>

<h2 id="小结">小结</h2>

<p>这节课，我们学习了分布式系统组件使用消息队列时的三大需求：消息保序、重复消息处理和消息可靠性保证，这三大需求可以进一步转换为对消息队列的三大要求：消息数据有序存取，消息数据具有全局唯一编号，以及消息数据在消费完成后被删除。</p>

<p>我画了一张表格，汇总了用 List 和 Streams 实现消息队列的特点和区别。当然，在实践的过程中，你也可以根据新的积累，进一步补充和完善这张表。</p>

<p><img src="assets/b2d6581e43f573da6218e790bb8c6814-20221015223336-eq8kjy5.jpg" alt="" /></p>

<p>其实，关于 Redis 是否适合做消息队列，业界一直是有争论的。很多人认为，要使用消息队列，就应该采用 Kafka、RabbitMQ 这些专门面向消息队列场景的软件，而 Redis 更加适合做缓存。</p>

<p>根据这些年做 Redis 研发工作的经验，我的看法是：Redis 是一个非常轻量级的键值数据库，部署一个 Redis 实例就是启动一个进程，部署 Redis 集群，也就是部署多个 Redis 实例。而 Kafka、RabbitMQ 部署时，涉及额外的组件，例如 Kafka 的运行就需要再部署 ZooKeeper。相比 Redis 来说，Kafka 和 RabbitMQ 一般被认为是重量级的消息队列。</p>

<p>所以，关于是否用 Redis 做消息队列的问题，不能一概而论，我们需要考虑业务层面的数据体量，以及对性能、可靠性、可扩展性的需求。如果分布式系统中的组件消息通信量不大，那么，Redis 只需要使用有限的内存空间就能满足消息存储的需求，而且，Redis 的高性能特性能支持快速的消息读写，不失为消息队列的一个好的解决方案。</p>

<h2 id="每课一问">每课一问</h2>

<p>按照惯例，我给你提个小问题。如果一个生产者发送给消息队列的消息，需要被多个消费者进行读取和处理（例如，一个消息是一条从业务系统采集的数据，既要被消费者 1 读取进行实时计算，也要被消费者 2 读取并留存到分布式文件系统 HDFS 中，以便后续进行历史查询），你会使用 Redis 的什么数据类型来解决这个问题呢？</p>

<p>欢迎在留言区写下你的思考和答案，如果觉得今天的内容对你有所帮助，也欢迎你帮我分享给更多人。我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#731f1f1f4a474242434433141e121a1f5d101c1e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f121fdfc8d5cdc2',t:'MTczNDA1Mjc1MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
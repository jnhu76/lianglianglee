<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=11&#32;&#32;“万金油”的String，为什么不好用了？>
        <link rel="icon" href="/static/favicon.png">
        <title>11  “万金油”的String，为什么不好用了？ </title>
        
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
                            <h1 id="title" data-id="11  “万金油”的String，为什么不好用了？" class="title">11  “万金油”的String，为什么不好用了？</h1>
                            <div><p>从今天开始，我们就要进入“实践篇”了。接下来，我们会用 5 节课的时间学习“数据结构”。我会介绍节省内存开销以及保存和统计海量数据的数据类型及其底层数据结构，还会围绕典型的应用场景（例如地址位置查询、时间序列数据库读写和消息队列存取），跟你分享使用 Redis 的数据类型和 module 扩展功能来满足需求的具体方案。</p>

<p>今天，我们先了解下 String 类型的内存空间消耗问题，以及选择节省内存开销的数据类型的解决方案。</p>

<p>先跟你分享一个我曾经遇到的需求。</p>

<p>当时，我们要开发一个图片存储系统，要求这个系统能快速地记录图片 ID 和图片在存储系统中保存时的 ID（可以直接叫作图片存储对象 ID）。同时，还要能够根据图片 ID 快速查找到图片存储对象 ID。</p>

<p>因为图片数量巨大，所以我们就用 10 位数来表示图片 ID 和图片存储对象 ID，例如，图片 ID 为 1101000051，它在存储系统中对应的 ID 号是 3301000051。</p>

<pre><code class="language-c">photo_id: 1101000051
photo_obj_id: 3301000051
</code></pre>

<p>可以看到，图片 ID 和图片存储对象 ID 正好一一对应，是典型的“键 - 单值”模式。所谓的“单值”，就是指键值对中的值就是一个值，而不是一个集合，这和 String 类型提供的“一个键对应一个值的数据”的保存形式刚好契合。</p>

<p>而且，String 类型可以保存二进制字节流，就像“万金油”一样，只要把数据转成二进制字节数组，就可以保存了。</p>

<p>所以，我们的第一个方案就是用 String 保存数据。我们把图片 ID 和图片存储对象 ID 分别作为键值对的 key 和 value 来保存，其中，图片存储对象 ID 用了 String 类型。</p>

<p>刚开始，我们保存了 1 亿张图片，大约用了 6.4GB 的内存。但是，随着图片数据量的不断增加，我们的 Redis 内存使用量也在增加，结果就遇到了大内存 Redis 实例因为生成 RDB 而响应变慢的问题。很显然，String 类型并不是一种好的选择，我们还需要进一步寻找能节省内存开销的数据类型方案。</p>

<p>在这个过程中，我深入地研究了 String 类型的底层结构，找到了它内存开销大的原因，对“万金油”的 String 类型有了全新的认知：String 类型并不是适用于所有场合的，它有一个明显的短板，就是它保存数据时所消耗的内存空间较多。</p>

<p>同时，我还仔细研究了集合类型的数据结构。我发现，集合类型有非常节省内存空间的底层实现结构，但是，集合类型保存的数据模式，是一个键对应一系列值，并不适合直接保存单值的键值对。所以，我们就使用二级编码的方法，实现了用集合类型保存单值键值对，Redis 实例的内存空间消耗明显下降了。</p>

<p>这节课，我就把在解决这个问题时学到的经验和方法分享给你，包括 String 类型的内存空间消耗在哪儿了、用什么数据结构可以节省内存，以及如何用集合类型保存单值键值对。如果你在使用 String 类型时也遇到了内存空间消耗较多的问题，就可以尝试下今天的解决方案了。</p>

<p>接下来，我们先来看看 String 类型的内存都消耗在哪里了。</p>

<h2 id="为什么-string-类型内存开销大">为什么 String 类型内存开销大？</h2>

<p>在刚才的案例中，我们保存了 1 亿张图片的信息，用了约 6.4GB 的内存，一个图片 ID 和图片存储对象 ID 的记录平均用了 64 字节。</p>

<p>但问题是，一组图片 ID 及其存储对象 ID 的记录，实际只需要 16 字节就可以了。</p>

<p>我们来分析一下。图片 ID 和图片存储对象 ID 都是 10 位数，我们可以用两个 8 字节的 Long 类型表示这两个 ID。因为 8 字节的 Long 类型最大可以表示 2 的 64 次方的数值，所以肯定可以表示 10 位数。但是，为什么 String 类型却用了 64 字节呢？</p>

<p>其实，除了记录实际数据，String 类型还需要额外的内存空间记录数据长度、空间使用等信息，这些信息也叫作元数据。当实际保存的数据较小时，元数据的空间开销就显得比较大了，有点“喧宾夺主”的意思。</p>

<p>那么，String 类型具体是怎么保存数据的呢？我来解释一下。</p>

<p>当你保存 64 位有符号整数时，String 类型会把它保存为一个 8 字节的 Long 类型整数，这种保存方式通常也叫作 int 编码方式。</p>

<p>但是，当你保存的数据中包含字符时，String 类型就会用简单动态字符串（Simple Dynamic String，SDS）结构体来保存，如下图所示：</p>

<p><img src="assets/37c6a8d5abd65906368e7c4a6b938657-20221015223128-ksc0um3.jpg" alt="" /></p>

<ul>
<li><strong>buf</strong>：字节数组，保存实际数据。为了表示字节数组的结束，Redis 会自动在数组最后加一个“\0”，这就会额外占用 1 个字节的开销。</li>
<li><strong>len</strong>：占 4 个字节，表示 buf 的已用长度。</li>
<li><strong>alloc</strong>：也占个 4 字节，表示 buf 的实际分配长度，一般大于 len。</li>
</ul>

<p>可以看到，在 SDS 中，buf 保存实际数据，而 len 和 alloc 本身其实是 SDS 结构体的额外开销。</p>

<p>另外，对于 String 类型来说，除了 SDS 的额外开销，还有一个来自于 RedisObject 结构体的开销。</p>

<p>因为 Redis 的数据类型有很多，而且，不同数据类型都有些相同的元数据要记录（比如最后一次访问的时间、被引用的次数等），所以，Redis 会用一个 RedisObject 结构体来统一记录这些元数据，同时指向实际数据。</p>

<p>一个 RedisObject 包含了 8 字节的元数据和一个 8 字节指针，这个指针再进一步指向具体数据类型的实际数据所在，例如指向 String 类型的 SDS 结构所在的内存地址，可以看一下下面的示意图。关于 RedisObject 的具体结构细节，我会在后面的课程中详细介绍，现在你只要了解它的基本结构和元数据开销就行了。</p>

<p><img src="assets/3409948e9d3e8aa5cd7cafb9b66c2857-20221015223128-9rrrq48.jpg" alt="" /></p>

<p>为了节省内存空间，Redis 还对 Long 类型整数和 SDS 的内存布局做了专门的设计。</p>

<p>一方面，当保存的是 Long 类型整数时，RedisObject 中的指针就直接赋值为整数数据了，这样就不用额外的指针再指向整数了，节省了指针的空间开销。</p>

<p>另一方面，当保存的是字符串数据，并且字符串小于等于 44 字节时，RedisObject 中的元数据、指针和 SDS 是一块连续的内存区域，这样就可以避免内存碎片。这种布局方式也被称为 embstr 编码方式。</p>

<p>当然，当字符串大于 44 字节时，SDS 的数据量就开始变多了，Redis 就不再把 SDS 和 RedisObject 布局在一起了，而是会给 SDS 分配独立的空间，并用指针指向 SDS 结构。这种布局方式被称为 raw 编码模式。</p>

<p>为了帮助你理解 int、embstr 和 raw 这三种编码模式，我画了一张示意图，如下所示：</p>

<p><img src="assets/ce83d1346c9642fdbbf5ffbe701bfbe3-20221015223128-f2kz2pu.jpg" alt="" /></p>

<p>好了，知道了 RedisObject 所包含的额外元数据开销，现在，我们就可以计算 String 类型的内存使用量了。</p>

<p>因为 10 位数的图片 ID 和图片存储对象 ID 是 Long 类型整数，所以可以直接用 int 编码的 RedisObject 保存。每个 int 编码的 RedisObject 元数据部分占 8 字节，指针部分被直接赋值为 8 字节的整数了。此时，每个 ID 会使用 16 字节，加起来一共是 32 字节。但是，另外的 32 字节去哪儿了呢？</p>

<p>我在【第 2 讲】中说过，Redis 会使用一个全局哈希表保存所有键值对，哈希表的每一项是一个 dictEntry 的结构体，用来指向一个键值对。dictEntry 结构中有三个 8 字节的指针，分别指向 key、value 以及下一个 dictEntry，三个指针共 24 字节，如下图所示：</p>

<p><img src="assets/b6cbc5161388fdf4c9b49f3802ef53e7-20221015223128-adrxmxz.jpg" alt="" /></p>

<p>但是，这三个指针只有 24 字节，为什么会占用了 32 字节呢？这就要提到 Redis 使用的内存分配库 jemalloc 了。</p>

<p>jemalloc 在分配内存时，会根据我们申请的字节数 N，找一个比 N 大，但是最接近 N 的 2 的幂次数作为分配的空间，这样可以减少频繁分配的次数。</p>

<p>举个例子。如果你申请 6 字节空间，jemalloc 实际会分配 8 字节空间；如果你申请 24 字节空间，jemalloc 则会分配 32 字节。所以，在我们刚刚说的场景里，dictEntry 结构就占用了 32 字节。</p>

<p>好了，到这儿，你应该就能理解，为什么用 String 类型保存图片 ID 和图片存储对象 ID 时需要用 64 个字节了。</p>

<p>你看，明明有效信息只有 16 字节，使用 String 类型保存时，却需要 64 字节的内存空间，有 48 字节都没有用于保存实际的数据。我们来换算下，如果要保存的图片有 1 亿张，那么 1 亿条的图片 ID 记录就需要 6.4GB 内存空间，其中有 4.8GB 的内存空间都用来保存元数据了，额外的内存空间开销很大。那么，有没有更加节省内存的方法呢？</p>

<h2 id="用什么数据结构可以节省内存">用什么数据结构可以节省内存？</h2>

<p>Redis 有一种底层数据结构，叫压缩列表（ziplist），这是一种非常节省内存的结构。</p>

<p>我们先回顾下压缩列表的构成。表头有三个字段 zlbytes、zltail 和 zllen，分别表示列表长度、列表尾的偏移量，以及列表中的 entry 个数。压缩列表尾还有一个 zlend，表示列表结束。</p>

<p><img src="assets/f6d4df5f7d6e80de29e2c6446b02429f-20221015223128-5dijzgy.jpg" alt="" /></p>

<p>压缩列表之所以能节省内存，就在于它是用一系列连续的 entry 保存数据。每个 entry 的元数据包括下面几部分。</p>

<ul>
<li><strong>prev_len</strong>，表示前一个 entry 的长度。prev_len 有两种取值情况：1 字节或 5 字节。取值 1 字节时，表示上一个 entry 的长度小于 254 字节。虽然 1 字节的值能表示的数值范围是 0 到 255，但是压缩列表中 zlend 的取值默认是 255，因此，就默认用 255 表示整个压缩列表的结束，其他表示长度的地方就不能再用 255 这个值了。所以，当上一个 entry 长度小于 254 字节时，prev_len 取值为 1 字节，否则，就取值为 5 字节。</li>
<li><strong>len</strong>：表示自身长度，4 字节；</li>
<li><strong>encoding</strong>：表示编码方式，1 字节；</li>
<li><strong>content</strong>：保存实际数据。</li>
</ul>

<p>这些 entry 会挨个儿放置在内存中，不需要再用额外的指针进行连接，这样就可以节省指针所占用的空间。</p>

<p>我们以保存图片存储对象 ID 为例，来分析一下压缩列表是如何节省内存空间的。</p>

<p>每个 entry 保存一个图片存储对象 ID（8 字节），此时，每个 entry 的 prev_len 只需要 1 个字节就行，因为每个 entry 的前一个 entry 长度都只有 8 字节，小于 254 字节。这样一来，一个图片的存储对象 ID 所占用的内存大小是 14 字节（1+4+1+8=14），实际分配 16 字节。</p>

<p>Redis 基于压缩列表实现了 List、Hash 和 Sorted Set 这样的集合类型，这样做的最大好处就是节省了 dictEntry 的开销。当你用 String 类型时，一个键值对就有一个 dictEntry，要用 32 字节空间。但采用集合类型时，一个 key 就对应一个集合的数据，能保存的数据多了很多，但也只用了一个 dictEntry，这样就节省了内存。</p>

<p>这个方案听起来很好，但还存在一个问题：在用集合类型保存键值对时，一个键对应了一个集合的数据，但是在我们的场景中，一个图片 ID 只对应一个图片的存储对象 ID，我们该怎么用集合类型呢？换句话说，在一个键对应一个值（也就是单值键值对）的情况下，我们该怎么用集合类型来保存这种单值键值对呢？</p>

<h2 id="如何用集合类型保存单值的键值对">如何用集合类型保存单值的键值对？</h2>

<p>在保存单值的键值对时，可以采用基于 Hash 类型的二级编码方法。这里说的二级编码，就是把一个单值的数据拆分成两部分，前一部分作为 Hash 集合的 key，后一部分作为 Hash 集合的 value，这样一来，我们就可以把单值数据保存到 Hash 集合中了。</p>

<p>以图片 ID 1101000060 和图片存储对象 ID 3302000080 为例，我们可以把图片 ID 的前 7 位（1101000）作为 Hash 类型的键，把图片 ID 的最后 3 位（060）和图片存储对象 ID 分别作为 Hash 类型值中的 key 和 value。</p>

<p>按照这种设计方法，我在 Redis 中插入了一组图片 ID 及其存储对象 ID 的记录，并且用 info 命令查看了内存开销，我发现，增加一条记录后，内存占用只增加了 16 字节，如下所示：</p>

<pre><code class="language-c">127.0.0.1:6379&gt; info memory
# Memory
used_memory:1039120
127.0.0.1:6379&gt; hset 1101000 060 3302000080
(integer) 1
127.0.0.1:6379&gt; info memory
# Memory
used_memory:1039136
</code></pre>

<p>在使用 String 类型时，每个记录需要消耗 64 字节，这种方式却只用了 16 字节，所使用的内存空间是原来的 1/4，满足了我们节省内存空间的需求。</p>

<p>不过，你可能也会有疑惑：“二级编码一定要把图片 ID 的前 7 位作为 Hash 类型的键，把最后 3 位作为 Hash 类型值中的 key 吗？”<strong>其实，</strong><strong>二级编码方法中采用的 ID 长度是有讲究的</strong>。</p>

<p>在【第 2 讲】中，我介绍过 Redis Hash 类型的两种底层实现结构，分别是压缩列表和哈希表。</p>

<p>那么，Hash 类型底层结构什么时候使用压缩列表，什么时候使用哈希表呢？其实，Hash 类型设置了用压缩列表保存数据时的两个阈值，一旦超过了阈值，Hash 类型就会用哈希表来保存数据了。</p>

<p>这两个阈值分别对应以下两个配置项：</p>

<ul>
<li>hash-max-ziplist-entries：表示用压缩列表保存时哈希集合中的最大元素个数。</li>
<li>hash-max-ziplist-value：表示用压缩列表保存时哈希集合中单个元素的最大长度。</li>
</ul>

<p>如果我们往 Hash 集合中写入的元素个数超过了 hash-max-ziplist-entries，或者写入的单个元素大小超过了 hash-max-ziplist-value，Redis 就会自动把 Hash 类型的实现结构由压缩列表转为哈希表。</p>

<p>一旦从压缩列表转为了哈希表，Hash 类型就会一直用哈希表进行保存，而不会再转回压缩列表了。在节省内存空间方面，哈希表就没有压缩列表那么高效了。</p>

<p><strong>为了能充分使用压缩列表的精简内存布局，我们一般要控制保存在 Hash 集合中的元素个数</strong>。所以，在刚才的二级编码中，我们只用图片 ID 最后 3 位作为 Hash 集合的 key，也就保证了 Hash 集合的元素个数不超过 1000，同时，我们把 hash-max-ziplist-entries 设置为 1000，这样一来，Hash 集合就可以一直使用压缩列表来节省内存空间了。</p>

<h2 id="小结">小结</h2>

<p>这节课，我们打破了对 String 的认知误区，以前，我们认为 String 是“万金油”，什么场合都适用，但是，在保存的键值对本身占用的内存空间不大时（例如这节课里提到的的图片 ID 和图片存储对象 ID），String 类型的元数据开销就占据主导了，这里面包括了 RedisObject 结构、SDS 结构、dictEntry 结构的内存开销。</p>

<p>针对这种情况，我们可以使用压缩列表保存数据。当然，使用 Hash 这种集合类型保存单值键值对的数据时，我们需要将单值数据拆分成两部分，分别作为 Hash 集合的键和值，就像刚才案例中用二级编码来表示图片 ID，希望你能把这个方法用到自己的场景中。</p>

<p>最后，我还想再给你提供一个小方法：如果你想知道键值对采用不同类型保存时的内存开销，可以在<a href="http://www.redis.cn/redis_memory/" target="_blank">这个网址</a>里输入你的键值对长度和使用的数据类型，这样就能知道实际消耗的内存大小了。建议你把这个小工具用起来，它可以帮助你充分地节省内存。</p>

<h2 id="每课一问">每课一问</h2>

<p>按照惯例，给你提个小问题：除了 String 类型和 Hash 类型，你觉得，还有其他合适的类型可以应用在这节课所说的保存图片的例子吗？</p>

<p>欢迎在留言区写下你的思考和答案，我们一起交流讨论，也欢迎你把今天的内容分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#caa6a6a6f3fefbfbfafd8aada7aba3a6e4a9a5a7" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f121eb3187ccdc2',t:'MTczNDA1MjcwMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
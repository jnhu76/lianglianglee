<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=04&#32;Redis&#32;持久化——AOF>
        <link rel="icon" href="/static/favicon.png">
        <title>04 Redis 持久化——AOF </title>
        
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
                        <a class="menu-item" id="01 Redis 是如何执行的.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/01%20Redis%20%e6%98%af%e5%a6%82%e4%bd%95%e6%89%a7%e8%a1%8c%e7%9a%84.md">01 Redis 是如何执行的.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Redis 快速搭建与使用.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/02%20Redis%20%e5%bf%ab%e9%80%9f%e6%90%ad%e5%bb%ba%e4%b8%8e%e4%bd%bf%e7%94%a8.md">02 Redis 快速搭建与使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 Redis 持久化——RDB.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/03%20Redis%20%e6%8c%81%e4%b9%85%e5%8c%96%e2%80%94%e2%80%94RDB.md">03 Redis 持久化——RDB.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 Redis 持久化——AOF.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/04%20Redis%20%e6%8c%81%e4%b9%85%e5%8c%96%e2%80%94%e2%80%94AOF.md">04 Redis 持久化——AOF.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 Redis 持久化——混合持久化.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/05%20Redis%20%e6%8c%81%e4%b9%85%e5%8c%96%e2%80%94%e2%80%94%e6%b7%b7%e5%90%88%e6%8c%81%e4%b9%85%e5%8c%96.md">05 Redis 持久化——混合持久化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 字符串使用与内部实现原理.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/06%20%e5%ad%97%e7%ac%a6%e4%b8%b2%e4%bd%bf%e7%94%a8%e4%b8%8e%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">06 字符串使用与内部实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 附录：更多字符串操作命令.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%e9%99%84%e5%bd%95%ef%bc%9a%e6%9b%b4%e5%a4%9a%e5%ad%97%e7%ac%a6%e4%b8%b2%e6%93%8d%e4%bd%9c%e5%91%bd%e4%bb%a4.md">07 附录：更多字符串操作命令.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 字典使用与内部实现原理.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%e5%ad%97%e5%85%b8%e4%bd%bf%e7%94%a8%e4%b8%8e%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">08 字典使用与内部实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 附录：更多字典操作命令.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%e9%99%84%e5%bd%95%ef%bc%9a%e6%9b%b4%e5%a4%9a%e5%ad%97%e5%85%b8%e6%93%8d%e4%bd%9c%e5%91%bd%e4%bb%a4.md">09 附录：更多字典操作命令.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 列表使用与内部实现原理.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%e5%88%97%e8%a1%a8%e4%bd%bf%e7%94%a8%e4%b8%8e%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">10 列表使用与内部实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 附录：更多列表操作命令.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%e9%99%84%e5%bd%95%ef%bc%9a%e6%9b%b4%e5%a4%9a%e5%88%97%e8%a1%a8%e6%93%8d%e4%bd%9c%e5%91%bd%e4%bb%a4.md">11 附录：更多列表操作命令.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 集合使用与内部实现原理.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%e9%9b%86%e5%90%88%e4%bd%bf%e7%94%a8%e4%b8%8e%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">12 集合使用与内部实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 附录：更多集合操作命令.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%e9%99%84%e5%bd%95%ef%bc%9a%e6%9b%b4%e5%a4%9a%e9%9b%86%e5%90%88%e6%93%8d%e4%bd%9c%e5%91%bd%e4%bb%a4.md">13 附录：更多集合操作命令.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 有序集合使用与内部实现原理.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%e6%9c%89%e5%ba%8f%e9%9b%86%e5%90%88%e4%bd%bf%e7%94%a8%e4%b8%8e%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86.md">14 有序集合使用与内部实现原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 附录：更多有序集合操作命令.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/15%20%e9%99%84%e5%bd%95%ef%bc%9a%e6%9b%b4%e5%a4%9a%e6%9c%89%e5%ba%8f%e9%9b%86%e5%90%88%e6%93%8d%e4%bd%9c%e5%91%bd%e4%bb%a4.md">15 附录：更多有序集合操作命令.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Redis 事务深入解析.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/16%20Redis%20%e4%ba%8b%e5%8a%a1%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90.md">16 Redis 事务深入解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 Redis 键值过期操作.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/17%20Redis%20%e9%94%ae%e5%80%bc%e8%bf%87%e6%9c%9f%e6%93%8d%e4%bd%9c.md">17 Redis 键值过期操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Redis 过期策略与源码分析.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/18%20Redis%20%e8%bf%87%e6%9c%9f%e7%ad%96%e7%95%a5%e4%b8%8e%e6%ba%90%e7%a0%81%e5%88%86%e6%9e%90.md">18 Redis 过期策略与源码分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Redis 管道技术——Pipeline.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/19%20Redis%20%e7%ae%a1%e9%81%93%e6%8a%80%e6%9c%af%e2%80%94%e2%80%94Pipeline.md">19 Redis 管道技术——Pipeline.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 查询附近的人——GEO.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%e6%9f%a5%e8%af%a2%e9%99%84%e8%bf%91%e7%9a%84%e4%ba%ba%e2%80%94%e2%80%94GEO.md">20 查询附近的人——GEO.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 游标迭代器（过滤器）——Scan.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/21%20%e6%b8%b8%e6%a0%87%e8%bf%ad%e4%bb%a3%e5%99%a8%ef%bc%88%e8%bf%87%e6%bb%a4%e5%99%a8%ef%bc%89%e2%80%94%e2%80%94Scan.md">21 游标迭代器（过滤器）——Scan.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 优秀的基数统计算法——HyperLogLog.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%e4%bc%98%e7%a7%80%e7%9a%84%e5%9f%ba%e6%95%b0%e7%bb%9f%e8%ae%a1%e7%ae%97%e6%b3%95%e2%80%94%e2%80%94HyperLogLog.md">22 优秀的基数统计算法——HyperLogLog.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 内存淘汰机制与算法.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%e5%86%85%e5%ad%98%e6%b7%98%e6%b1%b0%e6%9c%ba%e5%88%b6%e4%b8%8e%e7%ae%97%e6%b3%95.md">23 内存淘汰机制与算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 消息队列——发布订阅模式.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e2%80%94%e2%80%94%e5%8f%91%e5%b8%83%e8%ae%a2%e9%98%85%e6%a8%a1%e5%bc%8f.md">24 消息队列——发布订阅模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 消息队列的其他实现方式.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e7%9a%84%e5%85%b6%e4%bb%96%e5%ae%9e%e7%8e%b0%e6%96%b9%e5%bc%8f.md">25 消息队列的其他实现方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 消息队列终极解决方案——Stream（上）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e7%bb%88%e6%9e%81%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88%e2%80%94%e2%80%94Stream%ef%bc%88%e4%b8%8a%ef%bc%89.md">26 消息队列终极解决方案——Stream（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 消息队列终极解决方案——Stream（下）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%e6%b6%88%e6%81%af%e9%98%9f%e5%88%97%e7%bb%88%e6%9e%81%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88%e2%80%94%e2%80%94Stream%ef%bc%88%e4%b8%8b%ef%bc%89.md">27 消息队列终极解决方案——Stream（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 实战：分布式锁详解与代码.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e8%af%a6%e8%a7%a3%e4%b8%8e%e4%bb%a3%e7%a0%81.md">28 实战：分布式锁详解与代码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 实战：布隆过滤器安装与使用及原理分析.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%b8%83%e9%9a%86%e8%bf%87%e6%bb%a4%e5%99%a8%e5%ae%89%e8%a3%85%e4%b8%8e%e4%bd%bf%e7%94%a8%e5%8f%8a%e5%8e%9f%e7%90%86%e5%88%86%e6%9e%90.md">29 实战：布隆过滤器安装与使用及原理分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 完整案例：实现延迟队列的两种方法.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%e5%ae%8c%e6%95%b4%e6%a1%88%e4%be%8b%ef%bc%9a%e5%ae%9e%e7%8e%b0%e5%bb%b6%e8%bf%9f%e9%98%9f%e5%88%97%e7%9a%84%e4%b8%a4%e7%a7%8d%e6%96%b9%e6%b3%95.md">30 完整案例：实现延迟队列的两种方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 实战：定时任务案例.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/31%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1%e6%a1%88%e4%be%8b.md">31 实战：定时任务案例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 实战：RediSearch 高性能的全文搜索引擎.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%e5%ae%9e%e6%88%98%ef%bc%9aRediSearch%20%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84%e5%85%a8%e6%96%87%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e.md">32 实战：RediSearch 高性能的全文搜索引擎.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 实战：Redis 性能测试.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/33%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95.md">33 实战：Redis 性能测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 实战：Redis 慢查询.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/34%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e6%85%a2%e6%9f%a5%e8%af%a2.md">34 实战：Redis 慢查询.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 实战：Redis 性能优化方案.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/35%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e6%96%b9%e6%a1%88.md">35 实战：Redis 性能优化方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 实战：Redis 主从同步.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/36%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e4%b8%bb%e4%bb%8e%e5%90%8c%e6%ad%a5.md">36 实战：Redis 主从同步.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 实战：Redis哨兵模式（上）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/37%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%e5%93%a8%e5%85%b5%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%b8%8a%ef%bc%89.md">37 实战：Redis哨兵模式（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 实战：Redis 哨兵模式（下）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/38%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e5%93%a8%e5%85%b5%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%b8%8b%ef%bc%89.md">38 实战：Redis 哨兵模式（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 实战：Redis 集群模式（上）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/39%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e9%9b%86%e7%be%a4%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%b8%8a%ef%bc%89.md">39 实战：Redis 集群模式（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 实战：Redis 集群模式（下）.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/40%20%e5%ae%9e%e6%88%98%ef%bc%9aRedis%20%e9%9b%86%e7%be%a4%e6%a8%a1%e5%bc%8f%ef%bc%88%e4%b8%8b%ef%bc%89.md">40 实战：Redis 集群模式（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 案例：Redis 问题汇总和相关解决方案.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/41%20%e6%a1%88%e4%be%8b%ef%bc%9aRedis%20%e9%97%ae%e9%a2%98%e6%b1%87%e6%80%bb%e5%92%8c%e7%9b%b8%e5%85%b3%e8%a7%a3%e5%86%b3%e6%96%b9%e6%a1%88.md">41 案例：Redis 问题汇总和相关解决方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 技能学习指南.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/42%20%e6%8a%80%e8%83%bd%e5%ad%a6%e4%b9%a0%e6%8c%87%e5%8d%97.md">42 技能学习指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 加餐：Redis 的可视化管理工具.md" href="/%e4%b8%93%e6%a0%8f/Redis%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%e4%b8%8e%e5%ae%9e%e6%88%98/43%20%e5%8a%a0%e9%a4%90%ef%bc%9aRedis%20%e7%9a%84%e5%8f%af%e8%a7%86%e5%8c%96%e7%ae%a1%e7%90%86%e5%b7%a5%e5%85%b7.md">43 加餐：Redis 的可视化管理工具.md</a>
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
                            <h1 id="title" data-id="04 Redis 持久化——AOF" class="title">04 Redis 持久化——AOF</h1>
                            <div><p>使用 RDB 持久化有一个风险，它可能会造成最新数据丢失的风险。因为 RDB 的持久化有一定的时间间隔，在这个时间段内如果 Redis 服务意外终止的话，就会造成最新的数据全部丢失。</p>

<p>可能会操作 Redis 服务意外终止的条件：</p>

<ul>
<li>安装 Redis 的机器停止运行，蓝屏或者系统崩溃；</li>
<li>安装 Redis 的机器出现电源故障，例如突然断电；</li>
<li>使用 <code>kill -9 Redis_PID</code> 等。</li>
</ul>

<p>那么如何解决以上的这些问题呢？Redis 为我们提供了另一种持久化的方案——AOF。</p>

<h3 id="1-简介">1 简介</h3>

<p>AOF（Append Only File）中文是附加到文件，顾名思义 AOF 可以把 Redis 每个键值对操作都记录到文件（appendonly.aof）中。</p>

<h3 id="2-持久化查询和设置">2 持久化查询和设置</h3>

<h4 id="1-查询-aof-启动状态">1）查询 AOF 启动状态</h4>

<p>使用 <code>config get appendonly</code> 命令，如下图所示： <img src="assets/2020-02-24-122526.png" alt="image.png" /> 其中，第一行为 AOF 文件的名称，而最后一行表示 AOF 启动的状态，yes 表示已启动，no 表示未启动。</p>

<h4 id="2-开启-aof-持久化">2）开启 AOF 持久化</h4>

<p>Redis 默认是关闭 AOF 持久化的，想要开启 AOF 持久化，有以下两种方式：</p>

<ul>
<li>通过命令行的方式；</li>
<li>通过修改配置文件的方式（redis.conf）。</li>
</ul>

<p>下面分别来看以上两种方式的实现。</p>

<h5 id="①-命令行启动-aof">① 命令行启动 AOF</h5>

<p>命令行启动 AOF，使用 <code>config set appendonly yes</code> 命令，如下图所示： <img src="assets/2020-02-24-122527.png" alt="image.png" /> <strong>命令行启动 AOF 的优缺点</strong>：命令行启动优点是无需重启 Redis 服务，缺点是如果 Redis 服务重启，则之前使用命令行设置的配置就会失效。</p>

<h5 id="②-配置文件启动-aof">② 配置文件启动 AOF</h5>

<p>Redis 的配置文件在它的根路径下的 redis.conf 文件中，获取 Redis 的根目录可以使用命令 <code>config get dir</code> 获取，如下图所示： <img src="assets/2020-02-24-122528.png" alt="image.png" /> 只需要在配置文件中设置 <code>appendonly yes</code> 即可，默认 <code>appendonly no</code> 表示关闭 AOF 持久化。 <strong>配置文件启动 AOF 的优缺点</strong>：修改配置文件的缺点是每次修改配置文件都要重启 Redis 服务才能生效，优点是无论重启多少次 Redis 服务，配置文件中设置的配置信息都不会失效。</p>

<h3 id="3-触发持久化">3 触发持久化</h3>

<p>AOF 持久化开启之后，只要满足一定条件，就会触发 AOF 持久化。AOF 的触发条件分为两种：自动触发和手动触发。</p>

<h4 id="1-自动触发">1）自动触发</h4>

<p>有两种情况可以自动触发 AOF 持久化，分为是：<strong>满足 AOF 设置的策略触发</strong>和<strong>满足 AOF 重写触发。</strong>其中，AOF 重写触发会在本文的后半部分详细介绍，这里重点来说 AOF 持久化策略都有哪些。 AOF 持久化策略，分为以下三种：</p>

<ul>
<li>always：每条 Redis 操作命令都会写入磁盘，最多丢失一条数据；</li>
<li>everysec：每秒钟写入一次磁盘，最多丢失一秒的数据；</li>
<li>no：不设置写入磁盘的规则，根据当前操作系统来决定何时写入磁盘，Linux 默认 30s 写入一次数据至磁盘。</li>
</ul>

<p>这三种配置可以在 Redis 的配置文件（redis.conf）中设置，如下代码所示：</p>

<pre><code># 开启每秒写入一次的持久化策略
appendfsync everysec

</code></pre>

<blockquote>
<p>小贴士：因为每次写入磁盘都会对 Redis 的性能造成一定的影响，所以要根据用户的实际情况设置相应的策略，一般设置每秒写入一次磁盘的频率就可以满足大部分的使用场景了。</p>
</blockquote>

<p>触发自动持久化的两种情况，如下图所示： <img src="assets/2020-02-24-122529.png" alt="image.png" /></p>

<h4 id="2-手动触发">2）手动触发</h4>

<p>在客户端执行 <code>bgrewriteaof</code> 命令就可以手动触发 AOF 持久化，如下图所示： <img src="assets/2020-02-24-122530.png" alt="image.png" /> 可以看出执行完 <code>bgrewriteaof</code> 命令之后，AOF 持久化就会被触发。</p>

<h3 id="4-aof-文件重写">4 AOF 文件重写</h3>

<p>AOF 是通过记录 Redis 的执行命令来持久化（保存）数据的，所以随着时间的流逝 AOF 文件会越来越多，这样不仅增加了服务器的存储压力，也会造成 Redis 重启速度变慢，为了解决这个问题 Redis 提供了 AOF 重写的功能。</p>

<h4 id="1-什么是-aof-重写">1）什么是 AOF 重写？</h4>

<p>AOF 重写指的是它会直接读取 Redis 服务器当前的状态，并压缩保存为 AOF 文件。例如，我们增加了一个计数器，并对它做了 99 次修改，如果不做 AOF 重写的话，那么持久化文件中就会有 100 条记录执行命令的信息，而 AOF 重写之后，之后记录一条此计数器最终的结果信息，这样就去除了所有的无效信息。</p>

<h4 id="2-aof-重写实现">2）AOF 重写实现</h4>

<p>触发 AOF 文件重写，要满足两个条件，这两个条件也是配置在 Redis 配置文件中的，它们分别：</p>

<ul>
<li>auto-aof-rewrite-min-size：允许 AOF 重写的最小文件容量，默认是 64mb 。</li>
<li>auto-aof-rewrite-percentage：AOF 文件重写的大小比例，默认值是 100，表示 100%，也就是只有当前 AOF 文件，比最后一次（上次）的 AOF 文件大一倍时，才会启动 AOF 文件重写。</li>
</ul>

<p>查询 auto-aof-rewrite-min-size 和 auto-aof-rewrite-percentage 的值，可使用 <code>config get xxx</code> 命令，如下图所示： <img src="assets/2020-02-24-122531.png" alt="image.png" /></p>

<blockquote>
<p>小贴士：只有同时满足 auto-aof-rewrite-min-size 和 auto-aof-rewrite-percentage 设置的条件，才会触发 AOF 文件重写。</p>
</blockquote>

<p><strong>注意</strong>：使用 <code>bgrewriteaof</code> 命令，可以自动触发 AOF 文件重写。</p>

<h4 id="3-aof-重写流程">3）AOF 重写流程</h4>

<p>AOF 文件重写是生成一个全新的文件，并把当前数据的最少操作命令保存到新文件上，当把所有的数据都保存至新文件之后，Redis 会交换两个文件，并把最新的持久化操作命令追加到新文件上。</p>

<h3 id="5-配置说明">5 配置说明</h3>

<p>合理的设置 AOF 的配置，可以保障 Redis 高效且稳定的运行，以下是 AOF 的全部配置信息和说明。</p>

<p>AOF 的配置参数在  Redis 的配置文件中，也就是 Redis 根路径下的 <code>redis.conf</code> 文件中，配置参数和说明如下：</p>

<pre><code># 是否开启 AOF，yes 为开启，默认是关闭
appendonly no

# AOF 默认文件名
appendfilename &quot;appendonly.aof&quot;

# AOF 持久化策略配置
# appendfsync always
appendfsync everysec
# appendfsync no

# AOF 文件重写的大小比例，默认值是 100，表示 100%，也就是只有当前 AOF 文件，比最后一次的 AOF 文件大一倍时，才会启动 AOF 文件重写。
auto-aof-rewrite-percentage 100

# 允许 AOF 重写的最小文件容量
auto-aof-rewrite-min-size 64mb

# 是否开启启动时加载 AOF 文件效验，默认值是 yes，表示尽可能的加载 AOF 文件，忽略错误部分信息，并启动 Redis 服务。
# 如果值为 no，则表示，停止启动 Redis，用户必须手动修复 AOF 文件才能正常启动 Redis 服务。
aof-load-truncated yes

</code></pre>

<p>其中比较重要的是 appendfsync 参数，用它来设置 AOF 的持久化策略，可以选择按时间间隔或者操作次数来存储 AOF 文件，这个参数的三个值在文章开头有说明，这里就不再复述了。</p>

<h3 id="6-数据恢复">6 数据恢复</h3>

<h4 id="1-正常数据恢复">1）正常数据恢复</h4>

<p>正常情况下，只要开启了 AOF 持久化，并且提供了正常的 appendonly.aof 文件，在 Redis 启动时就会自定加载 AOF 文件并启动，执行如下图所示： <img src="assets/2020-02-24-122532.png" alt="image.png" /> 其中 <code>DB loaded from append only file......</code> 表示 Redis 服务器在启动时，先去加载了 AOF 持久化文件。</p>

<blockquote>
<p>小贴士：默认情况下 appendonly.aof 文件保存在 Redis 的根目录下。</p>
</blockquote>

<p><strong>持久化文件加载规则</strong></p>

<ul>
<li>如果只开启了 AOF 持久化，Redis 启动时只会加载 AOF 文件（appendonly.aof），进行数据恢复；</li>
<li>如果只开启了 RDB 持久化，Redis 启动时只会加载 RDB 文件（dump.rdb），进行数据恢复；</li>
<li>如果同时开启了 RDB 和 AOF 持久化，Redis 启动时只会加载 AOF 文件（appendonly.aof），进行数据恢复。</li>
</ul>

<p>在 AOF 开启的情况下，即使 AOF 文件不存在，只有 RDB 文件，也不会加载 RDB 文件。 AOF 和 RDB 的加载流程如下图所示： <img src="assets/2020-02-24-122534.png" alt="image.png" /></p>

<h4 id="2-简单异常数据恢复">2）简单异常数据恢复</h4>

<p>在 AOF 写入文件时如果服务器崩溃，或者是 AOF 存储已满的情况下，AOF 的最后一条命令可能被截断，这就是异常的 AOF 文件。</p>

<p>在 AOF 文件异常的情况下，如果为修改 Redis 的配置文件，也就是使用 <code>aof-load-truncated</code> 等于 <code>yes</code> 的配置，Redis 在启动时会忽略最后一条命令，并顺利启动 Redis，执行结果如下：</p>

<pre><code>* Reading RDB preamble from AOF file...
* Reading the remaining AOF tail...
# !!! Warning: short read while loading the AOF file !!!
# !!! Truncating the AOF at offset 439 !!!
# AOF loaded anyway because aof-load-truncated is enabled

</code></pre>

<h4 id="3-复杂异常数据恢复">3）复杂异常数据恢复</h4>

<p>AOF 文件可能出现更糟糕的情况，当 AOF 文件不仅被截断，而且中间的命令也被破坏，这个时候再启动 Redis 会提示错误信息并中止运行，错误信息如下：</p>

<pre><code>* Reading the remaining AOF tail...
# Bad file format reading the append only file: make a backup of your AOF file, then use ./redis-check-aof --fix &lt;filename&gt;

</code></pre>

<p>出现此类问题的解决方案如下：</p>

<ol>
<li>首先使用 AOF 修复工具，检测出现的问题，在命令行中输入 <code>redis-check-aof</code> 命令，它会跳转到出现问题的命令行，这个时候可以尝试手动修复此文件；</li>
<li>如果无法手动修复，我们可以使用 <code>redis-check-aof --fix</code> 自动修复 AOF 异常文件，不过执行此命令，可能会导致异常部分至文件末尾的数据全部被丢弃。</li>
</ol>

<h3 id="7-优缺点">7 优缺点</h3>

<h4 id="aof-优点">AOF 优点</h4>

<ul>
<li>AOF 持久化保存的数据更加完整，AOF 提供了三种保存策略：每次操作保存、每秒钟保存一次、跟随系统的持久化策略保存，其中每秒保存一次，从数据的安全性和性能两方面考虑是一个不错的选择，也是 AOF 默认的策略，即使发生了意外情况，最多只会丢失 1s 钟的数据；</li>
<li>AOF 采用的是命令追加的写入方式，所以不会出现文件损坏的问题，即使由于某些意外原因，导致了最后操作的持久化数据写入了一半，也可以通过 redis-check-aof 工具轻松的修复；</li>
<li>AOF 持久化文件，非常容易理解和解析，它是把所有 Redis 键值操作命令，以文件的方式存入了磁盘。即使不小心使用 <code>flushall</code> 命令删除了所有键值信息，只要使用 AOF 文件，删除最后的 <code>flushall</code> 命令，重启 Redis 即可恢复之前误删的数据。</li>
</ul>

<h4 id="aof-缺点">AOF 缺点</h4>

<ul>
<li>对于相同的数据集来说，AOF 文件要大于 RDB 文件；</li>
<li>在 Redis 负载比较高的情况下，RDB 比 AOF 性能更好；</li>
<li>RDB 使用快照的形式来持久化整个 Redis 数据，而 AOF 只是将每次执行的命令追加到 AOF 文件中，因此从理论上说，RDB 比 AOF 更健壮。</li>
</ul>

<h3 id="8-小结">8 小结</h3>

<p>AOF 保存数据更加完整，它可以记录每次 Redis 的键值变化，或者是选择每秒保存一次数据。AOF 的持久化文件更加易读，但相比与二进制的 RDB 来说，所占的存储空间也越大，为了解决这个问题，AOF 提供自动化重写机制，最大程度的减少了 AOF 占用空间大的问题。同时 AOF 也提供了很方便的异常文件恢复命令： <code>redis-check-aof --fix</code> ，为使用 AOF 提供了很好的保障。</p>

<p><strong>参考&amp;鸣谢</strong> <a href="https://redis.io/topics/persistence" target="_blank">https://redis.io/topics/persistence</a> <a href="https://blog.csdn.net/qq_36318234/article/details/79994133" target="_blank">https://blog.csdn.net/qq_36318234/article/details/79994133</a> <a href="https://www.cnblogs.com/wdliu/p/9377278.html" target="_blank">https://www.cnblogs.com/wdliu/p/9377278.html</a></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b8d4d4d4818c8989888ff8dfd5d9d1d496dbd7d5" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f121438dcd0cdc2',t:'MTczNDA1MjI3NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=38&#32;加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑>
        <link rel="icon" href="/static/favicon.png">
        <title>38 加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑 </title>
        
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
                        <a class="menu-item" id="00 开篇词 业务代码真的会有这么多坑？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%9a%e5%8a%a1%e4%bb%a3%e7%a0%81%e7%9c%9f%e7%9a%84%e4%bc%9a%e6%9c%89%e8%bf%99%e4%b9%88%e5%a4%9a%e5%9d%91%ef%bc%9f.md">00 开篇词 业务代码真的会有这么多坑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 使用了并发工具类库，线程安全就高枕无忧了吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/01%20%e4%bd%bf%e7%94%a8%e4%ba%86%e5%b9%b6%e5%8f%91%e5%b7%a5%e5%85%b7%e7%b1%bb%e5%ba%93%ef%bc%8c%e7%ba%bf%e7%a8%8b%e5%ae%89%e5%85%a8%e5%b0%b1%e9%ab%98%e6%9e%95%e6%97%a0%e5%bf%a7%e4%ba%86%e5%90%97%ef%bc%9f.md">01 使用了并发工具类库，线程安全就高枕无忧了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 代码加锁：不要让“锁”事成为烦心事.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/02%20%e4%bb%a3%e7%a0%81%e5%8a%a0%e9%94%81%ef%bc%9a%e4%b8%8d%e8%a6%81%e8%ae%a9%e2%80%9c%e9%94%81%e2%80%9d%e4%ba%8b%e6%88%90%e4%b8%ba%e7%83%a6%e5%bf%83%e4%ba%8b.md">02 代码加锁：不要让“锁”事成为烦心事.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 线程池：业务代码最常用也最容易犯错的组件.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/03%20%e7%ba%bf%e7%a8%8b%e6%b1%a0%ef%bc%9a%e4%b8%9a%e5%8a%a1%e4%bb%a3%e7%a0%81%e6%9c%80%e5%b8%b8%e7%94%a8%e4%b9%9f%e6%9c%80%e5%ae%b9%e6%98%93%e7%8a%af%e9%94%99%e7%9a%84%e7%bb%84%e4%bb%b6.md">03 线程池：业务代码最常用也最容易犯错的组件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 连接池：别让连接池帮了倒忙.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/04%20%e8%bf%9e%e6%8e%a5%e6%b1%a0%ef%bc%9a%e5%88%ab%e8%ae%a9%e8%bf%9e%e6%8e%a5%e6%b1%a0%e5%b8%ae%e4%ba%86%e5%80%92%e5%bf%99.md">04 连接池：别让连接池帮了倒忙.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 HTTP调用：你考虑到超时、重试、并发了吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/05%20HTTP%e8%b0%83%e7%94%a8%ef%bc%9a%e4%bd%a0%e8%80%83%e8%99%91%e5%88%b0%e8%b6%85%e6%97%b6%e3%80%81%e9%87%8d%e8%af%95%e3%80%81%e5%b9%b6%e5%8f%91%e4%ba%86%e5%90%97%ef%bc%9f.md">05 HTTP调用：你考虑到超时、重试、并发了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 2成的业务代码的Spring声明式事务，可能都没处理正确.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/06%202%e6%88%90%e7%9a%84%e4%b8%9a%e5%8a%a1%e4%bb%a3%e7%a0%81%e7%9a%84Spring%e5%a3%b0%e6%98%8e%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%8c%e5%8f%af%e8%83%bd%e9%83%bd%e6%b2%a1%e5%a4%84%e7%90%86%e6%ad%a3%e7%a1%ae.md">06 2成的业务代码的Spring声明式事务，可能都没处理正确.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 数据库索引：索引并不是万能药.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/07%20%e6%95%b0%e6%8d%ae%e5%ba%93%e7%b4%a2%e5%bc%95%ef%bc%9a%e7%b4%a2%e5%bc%95%e5%b9%b6%e4%b8%8d%e6%98%af%e4%b8%87%e8%83%bd%e8%8d%af.md">07 数据库索引：索引并不是万能药.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 判等问题：程序里如何确定你就是你？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/08%20%e5%88%a4%e7%ad%89%e9%97%ae%e9%a2%98%ef%bc%9a%e7%a8%8b%e5%ba%8f%e9%87%8c%e5%a6%82%e4%bd%95%e7%a1%ae%e5%ae%9a%e4%bd%a0%e5%b0%b1%e6%98%af%e4%bd%a0%ef%bc%9f.md">08 判等问题：程序里如何确定你就是你？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 数值计算：注意精度、舍入和溢出问题.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/09%20%e6%95%b0%e5%80%bc%e8%ae%a1%e7%ae%97%ef%bc%9a%e6%b3%a8%e6%84%8f%e7%b2%be%e5%ba%a6%e3%80%81%e8%88%8d%e5%85%a5%e5%92%8c%e6%ba%a2%e5%87%ba%e9%97%ae%e9%a2%98.md">09 数值计算：注意精度、舍入和溢出问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 集合类：坑满地的List列表操作.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/10%20%e9%9b%86%e5%90%88%e7%b1%bb%ef%bc%9a%e5%9d%91%e6%bb%a1%e5%9c%b0%e7%9a%84List%e5%88%97%e8%a1%a8%e6%93%8d%e4%bd%9c.md">10 集合类：坑满地的List列表操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 空值处理：分不清楚的null和恼人的空指针.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/11%20%e7%a9%ba%e5%80%bc%e5%a4%84%e7%90%86%ef%bc%9a%e5%88%86%e4%b8%8d%e6%b8%85%e6%a5%9a%e7%9a%84null%e5%92%8c%e6%81%bc%e4%ba%ba%e7%9a%84%e7%a9%ba%e6%8c%87%e9%92%88.md">11 空值处理：分不清楚的null和恼人的空指针.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 异常处理：别让自己在出问题的时候变为瞎子.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/12%20%e5%bc%82%e5%b8%b8%e5%a4%84%e7%90%86%ef%bc%9a%e5%88%ab%e8%ae%a9%e8%87%aa%e5%b7%b1%e5%9c%a8%e5%87%ba%e9%97%ae%e9%a2%98%e7%9a%84%e6%97%b6%e5%80%99%e5%8f%98%e4%b8%ba%e7%9e%8e%e5%ad%90.md">12 异常处理：别让自己在出问题的时候变为瞎子.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 日志：日志记录真没你想象的那么简单.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/13%20%e6%97%a5%e5%bf%97%ef%bc%9a%e6%97%a5%e5%bf%97%e8%ae%b0%e5%bd%95%e7%9c%9f%e6%b2%a1%e4%bd%a0%e6%83%b3%e8%b1%a1%e7%9a%84%e9%82%a3%e4%b9%88%e7%ae%80%e5%8d%95.md">13 日志：日志记录真没你想象的那么简单.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 文件IO：实现高效正确的文件读写并非易事.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/14%20%e6%96%87%e4%bb%b6IO%ef%bc%9a%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e6%ad%a3%e7%a1%ae%e7%9a%84%e6%96%87%e4%bb%b6%e8%af%bb%e5%86%99%e5%b9%b6%e9%9d%9e%e6%98%93%e4%ba%8b.md">14 文件IO：实现高效正确的文件读写并非易事.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 序列化：一来一回你还是原来的你吗？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/15%20%e5%ba%8f%e5%88%97%e5%8c%96%ef%bc%9a%e4%b8%80%e6%9d%a5%e4%b8%80%e5%9b%9e%e4%bd%a0%e8%bf%98%e6%98%af%e5%8e%9f%e6%9d%a5%e7%9a%84%e4%bd%a0%e5%90%97%ef%bc%9f.md">15 序列化：一来一回你还是原来的你吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 用好Java 8的日期时间类，少踩一些“老三样”的坑.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/16%20%e7%94%a8%e5%a5%bdJava%208%e7%9a%84%e6%97%a5%e6%9c%9f%e6%97%b6%e9%97%b4%e7%b1%bb%ef%bc%8c%e5%b0%91%e8%b8%a9%e4%b8%80%e4%ba%9b%e2%80%9c%e8%80%81%e4%b8%89%e6%a0%b7%e2%80%9d%e7%9a%84%e5%9d%91.md">16 用好Java 8的日期时间类，少踩一些“老三样”的坑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 别以为“自动挡”就不可能出现OOM.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/17%20%e5%88%ab%e4%bb%a5%e4%b8%ba%e2%80%9c%e8%87%aa%e5%8a%a8%e6%8c%a1%e2%80%9d%e5%b0%b1%e4%b8%8d%e5%8f%af%e8%83%bd%e5%87%ba%e7%8e%b0OOM.md">17 别以为“自动挡”就不可能出现OOM.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 当反射、注解和泛型遇到OOP时，会有哪些坑？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/18%20%e5%bd%93%e5%8f%8d%e5%b0%84%e3%80%81%e6%b3%a8%e8%a7%a3%e5%92%8c%e6%b3%9b%e5%9e%8b%e9%81%87%e5%88%b0OOP%e6%97%b6%ef%bc%8c%e4%bc%9a%e6%9c%89%e5%93%aa%e4%ba%9b%e5%9d%91%ef%bc%9f.md">18 当反射、注解和泛型遇到OOP时，会有哪些坑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Spring框架：IoC和AOP是扩展的核心.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/19%20Spring%e6%a1%86%e6%9e%b6%ef%bc%9aIoC%e5%92%8cAOP%e6%98%af%e6%89%a9%e5%b1%95%e7%9a%84%e6%a0%b8%e5%bf%83.md">19 Spring框架：IoC和AOP是扩展的核心.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Spring框架：框架帮我们做了很多工作也带来了复杂度.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/20%20Spring%e6%a1%86%e6%9e%b6%ef%bc%9a%e6%a1%86%e6%9e%b6%e5%b8%ae%e6%88%91%e4%bb%ac%e5%81%9a%e4%ba%86%e5%be%88%e5%a4%9a%e5%b7%a5%e4%bd%9c%e4%b9%9f%e5%b8%a6%e6%9d%a5%e4%ba%86%e5%a4%8d%e6%9d%82%e5%ba%a6.md">20 Spring框架：框架帮我们做了很多工作也带来了复杂度.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 代码重复：搞定代码重复的三个绝招.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/21%20%e4%bb%a3%e7%a0%81%e9%87%8d%e5%a4%8d%ef%bc%9a%e6%90%9e%e5%ae%9a%e4%bb%a3%e7%a0%81%e9%87%8d%e5%a4%8d%e7%9a%84%e4%b8%89%e4%b8%aa%e7%bb%9d%e6%8b%9b.md">21 代码重复：搞定代码重复的三个绝招.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 接口设计：系统间对话的语言，一定要统一.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/22%20%e6%8e%a5%e5%8f%a3%e8%ae%be%e8%ae%a1%ef%bc%9a%e7%b3%bb%e7%bb%9f%e9%97%b4%e5%af%b9%e8%af%9d%e7%9a%84%e8%af%ad%e8%a8%80%ef%bc%8c%e4%b8%80%e5%ae%9a%e8%a6%81%e7%bb%9f%e4%b8%80.md">22 接口设计：系统间对话的语言，一定要统一.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 缓存设计：缓存可以锦上添花也可以落井下石.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/23%20%e7%bc%93%e5%ad%98%e8%ae%be%e8%ae%a1%ef%bc%9a%e7%bc%93%e5%ad%98%e5%8f%af%e4%bb%a5%e9%94%a6%e4%b8%8a%e6%b7%bb%e8%8a%b1%e4%b9%9f%e5%8f%af%e4%bb%a5%e8%90%bd%e4%ba%95%e4%b8%8b%e7%9f%b3.md">23 缓存设计：缓存可以锦上添花也可以落井下石.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 业务代码写完，就意味着生产就绪了？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/24%20%e4%b8%9a%e5%8a%a1%e4%bb%a3%e7%a0%81%e5%86%99%e5%ae%8c%ef%bc%8c%e5%b0%b1%e6%84%8f%e5%91%b3%e7%9d%80%e7%94%9f%e4%ba%a7%e5%b0%b1%e7%bb%aa%e4%ba%86%ef%bc%9f.md">24 业务代码写完，就意味着生产就绪了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 异步处理好用，但非常容易用错.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/25%20%e5%bc%82%e6%ad%a5%e5%a4%84%e7%90%86%e5%a5%bd%e7%94%a8%ef%bc%8c%e4%bd%86%e9%9d%9e%e5%b8%b8%e5%ae%b9%e6%98%93%e7%94%a8%e9%94%99.md">25 异步处理好用，但非常容易用错.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 数据存储：NoSQL与RDBMS如何取长补短、相辅相成？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/26%20%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%ef%bc%9aNoSQL%e4%b8%8eRDBMS%e5%a6%82%e4%bd%95%e5%8f%96%e9%95%bf%e8%a1%a5%e7%9f%ad%e3%80%81%e7%9b%b8%e8%be%85%e7%9b%b8%e6%88%90%ef%bc%9f.md">26 数据存储：NoSQL与RDBMS如何取长补短、相辅相成？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 数据源头：任何客户端的东西都不可信任.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/27%20%e6%95%b0%e6%8d%ae%e6%ba%90%e5%a4%b4%ef%bc%9a%e4%bb%bb%e4%bd%95%e5%ae%a2%e6%88%b7%e7%ab%af%e7%9a%84%e4%b8%9c%e8%a5%bf%e9%83%bd%e4%b8%8d%e5%8f%af%e4%bf%a1%e4%bb%bb.md">27 数据源头：任何客户端的东西都不可信任.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 安全兜底：涉及钱时，必须考虑防刷、限量和防重.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/28%20%e5%ae%89%e5%85%a8%e5%85%9c%e5%ba%95%ef%bc%9a%e6%b6%89%e5%8f%8a%e9%92%b1%e6%97%b6%ef%bc%8c%e5%bf%85%e9%a1%bb%e8%80%83%e8%99%91%e9%98%b2%e5%88%b7%e3%80%81%e9%99%90%e9%87%8f%e5%92%8c%e9%98%b2%e9%87%8d.md">28 安全兜底：涉及钱时，必须考虑防刷、限量和防重.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 数据和代码：数据就是数据，代码就是代码.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/29%20%e6%95%b0%e6%8d%ae%e5%92%8c%e4%bb%a3%e7%a0%81%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%b0%b1%e6%98%af%e6%95%b0%e6%8d%ae%ef%bc%8c%e4%bb%a3%e7%a0%81%e5%b0%b1%e6%98%af%e4%bb%a3%e7%a0%81.md">29 数据和代码：数据就是数据，代码就是代码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 如何正确保存和传输敏感数据？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/30%20%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e4%bf%9d%e5%ad%98%e5%92%8c%e4%bc%a0%e8%be%93%e6%95%8f%e6%84%9f%e6%95%b0%e6%8d%ae%ef%bc%9f.md">30 如何正确保存和传输敏感数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 加餐1：带你吃透课程中Java 8的那些重要知识点（一）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/31%20%e5%8a%a0%e9%a4%901%ef%bc%9a%e5%b8%a6%e4%bd%a0%e5%90%83%e9%80%8f%e8%af%be%e7%a8%8b%e4%b8%adJava%208%e7%9a%84%e9%82%a3%e4%ba%9b%e9%87%8d%e8%a6%81%e7%9f%a5%e8%af%86%e7%82%b9%ef%bc%88%e4%b8%80%ef%bc%89.md">31 加餐1：带你吃透课程中Java 8的那些重要知识点（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 加餐2：带你吃透课程中Java 8的那些重要知识点（二）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/32%20%e5%8a%a0%e9%a4%902%ef%bc%9a%e5%b8%a6%e4%bd%a0%e5%90%83%e9%80%8f%e8%af%be%e7%a8%8b%e4%b8%adJava%208%e7%9a%84%e9%82%a3%e4%ba%9b%e9%87%8d%e8%a6%81%e7%9f%a5%e8%af%86%e7%82%b9%ef%bc%88%e4%ba%8c%ef%bc%89.md">32 加餐2：带你吃透课程中Java 8的那些重要知识点（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 加餐3：定位应用问题，排错套路很重要.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/33%20%e5%8a%a0%e9%a4%903%ef%bc%9a%e5%ae%9a%e4%bd%8d%e5%ba%94%e7%94%a8%e9%97%ae%e9%a2%98%ef%bc%8c%e6%8e%92%e9%94%99%e5%a5%97%e8%b7%af%e5%be%88%e9%87%8d%e8%a6%81.md">33 加餐3：定位应用问题，排错套路很重要.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 加餐4：分析定位Java问题，一定要用好这些工具（一）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/34%20%e5%8a%a0%e9%a4%904%ef%bc%9a%e5%88%86%e6%9e%90%e5%ae%9a%e4%bd%8dJava%e9%97%ae%e9%a2%98%ef%bc%8c%e4%b8%80%e5%ae%9a%e8%a6%81%e7%94%a8%e5%a5%bd%e8%bf%99%e4%ba%9b%e5%b7%a5%e5%85%b7%ef%bc%88%e4%b8%80%ef%bc%89.md">34 加餐4：分析定位Java问题，一定要用好这些工具（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 加餐5：分析定位Java问题，一定要用好这些工具（二）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/35%20%e5%8a%a0%e9%a4%905%ef%bc%9a%e5%88%86%e6%9e%90%e5%ae%9a%e4%bd%8dJava%e9%97%ae%e9%a2%98%ef%bc%8c%e4%b8%80%e5%ae%9a%e8%a6%81%e7%94%a8%e5%a5%bd%e8%bf%99%e4%ba%9b%e5%b7%a5%e5%85%b7%ef%bc%88%e4%ba%8c%ef%bc%89.md">35 加餐5：分析定位Java问题，一定要用好这些工具（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 加餐6：这15年来，我是如何在工作中学习技术和英语的？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/36%20%e5%8a%a0%e9%a4%906%ef%bc%9a%e8%bf%9915%e5%b9%b4%e6%9d%a5%ef%bc%8c%e6%88%91%e6%98%af%e5%a6%82%e4%bd%95%e5%9c%a8%e5%b7%a5%e4%bd%9c%e4%b8%ad%e5%ad%a6%e4%b9%a0%e6%8a%80%e6%9c%af%e5%92%8c%e8%8b%b1%e8%af%ad%e7%9a%84%ef%bc%9f.md">36 加餐6：这15年来，我是如何在工作中学习技术和英语的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 加餐7：程序员成长28计.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/37%20%e5%8a%a0%e9%a4%907%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%91%98%e6%88%90%e9%95%bf28%e8%ae%a1.md">37 加餐7：程序员成长28计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/38%20%e5%8a%a0%e9%a4%908%ef%bc%9aJava%e7%a8%8b%e5%ba%8f%e4%bb%8e%e8%99%9a%e6%8b%9f%e6%9c%ba%e8%bf%81%e7%a7%bb%e5%88%b0Kubernetes%e7%9a%84%e4%b8%80%e4%ba%9b%e5%9d%91.md">38 加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：代码篇思考题集锦（一）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e4%bb%a3%e7%a0%81%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e9%9b%86%e9%94%a6%ef%bc%88%e4%b8%80%ef%bc%89.md">答疑篇：代码篇思考题集锦（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：代码篇思考题集锦（三）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e4%bb%a3%e7%a0%81%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e9%9b%86%e9%94%a6%ef%bc%88%e4%b8%89%ef%bc%89.md">答疑篇：代码篇思考题集锦（三）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：代码篇思考题集锦（二）.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e4%bb%a3%e7%a0%81%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e9%9b%86%e9%94%a6%ef%bc%88%e4%ba%8c%ef%bc%89.md">答疑篇：代码篇思考题集锦（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：加餐篇思考题答案合集.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e5%8a%a0%e9%a4%90%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%90%88%e9%9b%86.md">答疑篇：加餐篇思考题答案合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：安全篇思考题答案合集.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e5%ae%89%e5%85%a8%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%90%88%e9%9b%86.md">答疑篇：安全篇思考题答案合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑篇：设计篇思考题答案合集.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%ad%94%e7%96%91%e7%af%87%ef%bc%9a%e8%ae%be%e8%ae%a1%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%90%88%e9%9b%86.md">答疑篇：设计篇思考题答案合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 写代码时，如何才能尽量避免踩坑？.md" href="/%e4%b8%93%e6%a0%8f/Java%20%e4%b8%9a%e5%8a%a1%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%94%99%e8%af%af%20100%20%e4%be%8b/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%86%99%e4%bb%a3%e7%a0%81%e6%97%b6%ef%bc%8c%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e5%b0%bd%e9%87%8f%e9%81%bf%e5%85%8d%e8%b8%a9%e5%9d%91%ef%bc%9f.md">结束语 写代码时，如何才能尽量避免踩坑？.md</a>
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
                            <h1 id="title" data-id="38 加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑" class="title">38 加餐8：Java程序从虚拟机迁移到Kubernetes的一些坑</h1>
                            <div><p>我们又见面了。结课并不意味着结束，我非常高兴能持续把好的内容分享给你，也希望你能继续在留言区与我保持交流，分享你的学习心得和实践经验。</p>

<p>使用 Kubernetes 大规模部署应用程序，可以提升整体资源利用率，提高集群稳定性，还能提供快速的集群扩容能力，甚至还可以实现集群根据压力自动扩容。因此，现在越来越多的公司开始把程序从虚拟机（VM）迁移到 Kubernetes 了。</p>

<p>在大多数的公司中，Kubernetes 集群由运维来搭建，而程序的发布一般也是由 CI/CD 平台完成。从虚拟机到 Kubernetes 的整个迁移过程，基本不需要修改任何代码，可能只是重新发布一次而已。所以，我们 Java 开发人员可能对迁移这个事情本身感知不强烈，认为 Kubernetes 只是运维需要知道的事情。但是程序一旦部署到了 Kubernetes 集群中，在容器环境中运行，总是会出现各种各样之前没有的奇怪的问题。</p>

<p>今天的加餐，就让我们一起看下这其中大概会遇到哪些“坑”，还有相应的“避坑方法”。</p>

<h2 id="pod-ip-不固定带来的坑">Pod IP 不固定带来的坑</h2>

<p>Pod 是 Kubernetes 中能够创建和部署应用的最小单元，我们可以通过 Pod IP 来访问到某一个应用实例，但需要注意的是，如果没有经过特殊配置，Pod IP 并不是固定不变的，会在 Pod 重启后会发生变化。</p>

<p>不过好在，通常我们的 Java 微服务都是没有状态的，我们并不需要通过 Pod IP 来访问到某一个特定的 Java 服务实例。通常来说，要访问到部署在 Kubernetes 中的微服务集群，有两种服务发现和访问的方式：</p>

<p>通过 Kubernetes 来实现。也就是通过 Service 进行内部服务的互访，通过 Ingress 从外部访问到服务集群。</p>

<p>通过微服务注册中心（比如 Eureka）来实现。也就是服务之间的互访通过客户端负载均衡后 + 直接访问 Pod IP 进行，外部访问到服务集群通过微服务网关转发请求。</p>

<p>使用这两种方式进行微服务的访问，我们都没有和 Pod IP 直接打交道，也不会把 Pod IP 记录持久化，所以一般不需要太关注 Pod IP 变动的问题。不过，在一些场景下，Pod IP 的变动会造成一些问题。</p>

<p>之前我就遇到过这样的情况：某任务调度中间件会记录被调度节点的 IP 到数据库，随后通过访问节点 IP 查看任务节点执行日志的时候，如果节点部署在 Kubernetes 中，那么节点重启后 Pod IP 就会变动。这样，之前记录在数据库中的老节点的 Pod IP 必然访问不到，那么就会发生无法查看任务日志的情况。</p>

<p>遇到这种情况，我们应该怎么做呢？这时候，可能就需要修改这个中间件，把任务执行日志也进行持久化，从而避免这种访问任务节点来查看日志的行为。</p>

<p>总之，我们需要意识到 Pod IP 不固定的问题，并且进行“避坑操作”：在迁移到 Kubernetes 集群之前，摸排一下是否会存在需要通过 IP 访问到老节点的情况，如果有的话需要进行改造。</p>

<h2 id="程序因为-oom-被杀进程的坑">程序因为 OOM 被杀进程的坑</h2>

<p>在 Kubernetes 集群中部署程序的时候，我们通常会为容器设置一定的内存限制（limit），容器不可以使用超出其资源 limit 属性所设置的资源量。如果容器内的 Java 程序使用了大量内存，可能会出现各种 OOM 的情况。</p>

<p>第一种情况，是 OS OOM Kill 问题。如果过量内存导致操作系统 Kernel 不稳定，操作系统可能就会杀死 Java 进程。这时候，你能在操作系统 /var/log/messages 日志中找到类似 oom_kill_process 的关键字。</p>

<p>第二种情况，是我们最常遇到的 Java 程序的 OOM 问题。程序超出堆内存的限制申请内存，导致 Heap OOM，后续可能会因为健康检测没有通过被 Kubernetes 重启 Pod。</p>

<p><img src="assets/2cf6d48915a0bce6834cf46edb462c04.png" alt="img" /></p>

<p>在 Kubernetes 中部署 Java 程序时，这两种情况都很常见，表现出的症状也都是 OOM 关键字 + 重启。所以，当运维同学说程序因为 OOM 被杀死或重启的时候，我们一定要和运维同学一起去区分清楚，到底是哪一种情况，然后再对症处理。</p>

<p>对于情况 1，问题的原因往往不是 Java 堆内存不够，更可能是程序使用了太多的堆外内存，超过了内存限制。这个时候，调大 JVM 最大堆内存只会让问题更严重，因为堆内存是可以通过 GC 回收的。我们需要分析 Java 进程哪部分区域内存占用过大，是不是合理，以及是否可能存在内存泄露问题。Java 进程的内存占用除了堆之外，还包括</p>

<p>直接内存</p>

<p>元数据区</p>

<p>线程栈大小 Xss * 线程数</p>

<p>JIT 代码缓存</p>

<p>GC、编译器使用额外空间</p>

<p>……</p>

<p>我们可以使用 NMT 打印各部分区域大小，从而判断到底是哪部分内存区域占用了过多内存，或是可能有内存泄露问题：</p>

<pre><code>java -XX:NativeMemoryTracking=smmary/detail -XX:+UnlockDiagnosticVMOptions -XX:+PrintNMTStatistics

</code></pre>

<p>如果你确认 OOM 是情况 2，那么我同样不建议直接调大堆内存的限制，防止之后再出现情况 1。我会更建议你把堆内存限制为容器内存限制的 50%~70%，预留出足够多的内存给堆外和 OS 核心。如果需要扩容堆内存的话，那么也需要同步扩容容器的内存 limit。此外，也需要通过 Heap Dump（你可以回顾下第 35 讲的相关内容）等手段来排查为什么堆内存占用会这么大，排除潜在的内存泄露的可能性。</p>

<h2 id="内存和-cpu-资源配置不适配容器的坑">内存和 CPU 资源配置不适配容器的坑</h2>

<p>刚刚我们提到了，堆内存扩容需要结合容器内存 limit 同步进行。其实，我们更希望的是，Java 程序的堆内存配置能随着容器的资源配置，实现自动扩容或缩容，而不是写死 Xmx 和 Xms。这样一来，运维同学可以更方便地针对整个集群进行扩容或缩容。</p>

<p>对于 JDK&gt;8u191 的版本，我们可以设置下面这些 JVM 参数，来让 JVM 自动根据容器内存限制来设置堆内存用量。比如，下面配置相当于把 Xmx 和 Xms 都设置为了容器内存 limit 的 50%：</p>

<pre><code>XX:MaxRAMPercentage=50.0 -XX:InitialRAMPercentage=50.0 -XX:MinRAMPercentage=50.0

</code></pre>

<p>接下来，我们看看 CPU 资源配置不适配容器的坑，以及对应的解决方案。</p>

<p>对于 CPU 资源的使用，我们主要需要注意的是，代码中的各种组件甚至是 JVM 本身，会根据 CPU 数来配置并发数等重要参数指标，那么：</p>

<p>如果这个值因为 JVM 对容器的兼容性问题取到了 Kubernetes 工作节点的 CPU 数量，那么这个数量可能就不是 4 或 8，而是 128 以上，进而导致并发数过高。</p>

<p>对于 JDK&gt;8u191 的版本可能会对容器兼容性较好，但是其获取到的 Runtime.getRuntime().availableProcessors() 其实是 request 的值而不是 limit 的值（比如我们设置 request 为 2、limit 为 8、CICompilerCount 和 ParallelGCThreads 可能只是 2），那么可能并发数就会过低，进而影响 JVM 的 GC 或编译性能。</p>

<p>所以，我的建议是：</p>

<p>第一，通过 -XX:+PrintFlagsFinal 开关，来确认 ActiveProcessorCount 是不是符合我们的期望，并且确认 CICompilerCount、ParallelGCThreads 等重要参数配置是否合理。</p>

<p>第二，直接设置 CPU 的 request 和 limit 一致，或是对于 JDK&gt;8u191 的版本可以通过 -XX:ActiveProcessorCount=xxx 直接把 ActiveProcessorCount 设置为容器的 CPU limit。</p>

<h2 id="pod-重启以及重启后没有现场的坑">Pod 重启以及重启后没有现场的坑</h2>

<p>除非宿主机有问题，否则虚拟机不太会自己重启或被重启，而 Kubernetes 中 Pod 的重启绝非小概率事件。在存活检测不通过、Pod 重新进行节点调度等情况下，Pod 都会进行重启。对于 Pod 的重启，我们需要关注两个问题。</p>

<p>第一个问题是，分析 Pod 为什么会重启。</p>

<p>其中，除了“程序因为 OOM 被杀进程的坑”这部分提到的 OOM 的问题之外，我们还需要关注存活检查不通过的情况。</p>

<p>Kubernetes 有 readinessProbe 和 livenessProbe 两个探针，前者用于检查应用是否已经启动完成，后者用于持续探活。一般而言，运维同学会配置这 2 个探针为一个健康检测的断点，如果健康检测访问一次需要消耗比较长的时间（比如涉及到存储或外部服务可用性检测），那么很可能可以通过 readinessProbe 的检查但不通过 livenessProbe 检查（毕竟我们通常会为 readinessProbe 设置比较长的超时时间，而对于 livenessProbe 则没有那么宽容）。此外，健康检测也可能会受到 Full GC 的干扰导致超时。所以，我们需要和运维同学一起确认 livenessProbe 的配置地址和超时时间设置是否合理，防止偶发的 livenessProbe 探活失败导致的 Pod 重启。</p>

<p>第二个问题是，要理解 Pod 和虚拟机不同。</p>

<p>虚拟机一般都是有状态的，即便部署在虚拟机内的 Java 程序重启了，我们始终能有现场。而对于 Pod 重启来说，则是新建一个 Pod，这就意味着老的 Pod 无法进入。因此，如果因为堆 OOM 问题导致重启，我们希望事后查看当时 OS 的一些日志或是在现场执行一些命令来分析问题，就不太可能了。</p>

<p>所以，我们需要想办法在 Pod 关闭之前尽可能保留现场，比如：</p>

<p>对于程序的应用日志、标准输出、GC 日志等可以直接挂载到持久卷，不要保存在容器内部。</p>

<p>对于程序的堆栈现场保留，可以配置 -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath 在堆 OOM 的时候生成 Dump；还可以让 JVM 调用任一个 shell 脚本，通过脚本来保留线程栈等信息：</p>

<pre><code>-XX:OnOutOfMemoryError=saveinfo.sh

</code></pre>

<p>对于容器的现场保留，可以让运维配置 preStop 钩子，在 Pod 关闭之前把必要的信息上传到持久卷或云上。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>今天，我们探讨了 Java 应用部署到 Kubernetes 集群会遇到的 4 类问题。</p>

<p>第一类问题是，我们需要理解应用的 IP 会动态变化，因此要在设计上解除对 Pod IP 的强依赖，使用依赖服务发现来定位到应用。</p>

<p>第二类问题是，在出现 OOM 问题的时候，首先要区分 OOM 的原因来自 Java 进程层面还是容器层面。如果是容器层面的话，我们还需要进一步分析到底是哪个内存区域占用了过多内存，定位到问题后再根据容器资源设置合理的 JVM 参数或进行资源扩容。</p>

<p>第三类问题是，需要确保程序使用的内存和 CPU 资源匹配容器的资源限制，既要确保程序所“看”到的主机资源信息是容器本身的而不是物理机的，又要确保程序能尽可能随着容器扩容而扩容其资源限制。</p>

<p>第四类问题是，我们需要重点关注程序非发布期重启的问题，并且针对 Pod 的重启问题做好现场保留的准备工作，排除资源配置不合理、存活检查不通过等可能性，以避免因为程序频繁重启导致的偶发性能问题或可用性问题。</p>

<p>只有解决了这些隐患，才能让 Kubernetes 集群更好地发挥作用。</p>

<h2 id="思考与讨论">思考与讨论</h2>

<p>在你的工作中，还遇到过 Java+Kubernetes 中的其他坑吗？</p>

<p>我是朱晔，欢迎在评论区与我留言分享，也欢迎你把今天的内容分享给你的朋友或同事，一起交流。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a5c9c9c99c9194949592e5c2c8c4ccc98bc6cac8" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0e1dd93ef5653b',t:'MTczNDAxMDcyNS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
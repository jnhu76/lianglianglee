<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=31&#32;&#32;数据库文件系统实例：MySQL&#32;中&#32;B&#32;树和&#32;B&#43;&#32;树有什么区别？>
        <link rel="icon" href="/static/favicon.png">
        <title>31  数据库文件系统实例：MySQL 中 B 树和 B&#43; 树有什么区别？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  为什么大厂面试必考操作系统？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%b8%ba%e4%bb%80%e4%b9%88%e5%a4%a7%e5%8e%82%e9%9d%a2%e8%af%95%e5%bf%85%e8%80%83%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">00 开篇词  为什么大厂面试必考操作系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 课前必读  构建知识体系，可以这样做！.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/00%20%e8%af%be%e5%89%8d%e5%bf%85%e8%af%bb%20%20%e6%9e%84%e5%bb%ba%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%ef%bc%8c%e5%8f%af%e4%bb%a5%e8%bf%99%e6%a0%b7%e5%81%9a%ef%bc%81.md">00 课前必读  构建知识体系，可以这样做！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  计算机是什么：“如何把程序写好”这个问题是可计算的吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/01%20%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9a%e2%80%9c%e5%a6%82%e4%bd%95%e6%8a%8a%e7%a8%8b%e5%ba%8f%e5%86%99%e5%a5%bd%e2%80%9d%e8%bf%99%e4%b8%aa%e9%97%ae%e9%a2%98%e6%98%af%e5%8f%af%e8%ae%a1%e7%ae%97%e7%9a%84%e5%90%97%ef%bc%9f.md">01  计算机是什么：“如何把程序写好”这个问题是可计算的吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  程序的执行：相比 32 位，64 位的优势是什么？（上）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/02%20%20%e7%a8%8b%e5%ba%8f%e7%9a%84%e6%89%a7%e8%a1%8c%ef%bc%9a%e7%9b%b8%e6%af%94%2032%20%e4%bd%8d%ef%bc%8c64%20%e4%bd%8d%e7%9a%84%e4%bc%98%e5%8a%bf%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">02  程序的执行：相比 32 位，64 位的优势是什么？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  程序的执行：相比 32 位，64 位的优势是什么？（下）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/03%20%20%e7%a8%8b%e5%ba%8f%e7%9a%84%e6%89%a7%e8%a1%8c%ef%bc%9a%e7%9b%b8%e6%af%94%2032%20%e4%bd%8d%ef%bc%8c64%20%e4%bd%8d%e7%9a%84%e4%bc%98%e5%8a%bf%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">03  程序的执行：相比 32 位，64 位的优势是什么？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  构造复杂的程序：将一个递归函数转成非递归函数的通用方法.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/04%20%20%e6%9e%84%e9%80%a0%e5%a4%8d%e6%9d%82%e7%9a%84%e7%a8%8b%e5%ba%8f%ef%bc%9a%e5%b0%86%e4%b8%80%e4%b8%aa%e9%80%92%e5%bd%92%e5%87%bd%e6%95%b0%e8%bd%ac%e6%88%90%e9%9d%9e%e9%80%92%e5%bd%92%e5%87%bd%e6%95%b0%e7%9a%84%e9%80%9a%e7%94%a8%e6%96%b9%e6%b3%95.md">04  构造复杂的程序：将一个递归函数转成非递归函数的通用方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  存储器分级：L1 Cache 比内存和 SSD 快多少倍？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/05%20%20%e5%ad%98%e5%82%a8%e5%99%a8%e5%88%86%e7%ba%a7%ef%bc%9aL1%20Cache%20%e6%af%94%e5%86%85%e5%ad%98%e5%92%8c%20SSD%20%e5%bf%ab%e5%a4%9a%e5%b0%91%e5%80%8d%ef%bc%9f.md">05  存储器分级：L1 Cache 比内存和 SSD 快多少倍？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 (1) 加餐  练习题详解（一）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/05%20%281%29%20%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e4%b8%80%ef%bc%89.md">05 (1) 加餐  练习题详解（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  目录结构和文件管理指令：rm  -rf 指令的作用是？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/06%20%20%e7%9b%ae%e5%bd%95%e7%bb%93%e6%9e%84%e5%92%8c%e6%96%87%e4%bb%b6%e7%ae%a1%e7%90%86%e6%8c%87%e4%bb%a4%ef%bc%9arm%20%20-rf%20%e6%8c%87%e4%bb%a4%e7%9a%84%e4%bd%9c%e7%94%a8%e6%98%af%ef%bc%9f.md">06  目录结构和文件管理指令：rm  -rf 指令的作用是？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  进程、重定向和管道指令：xargs 指令的作用是？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/07%20%20%e8%bf%9b%e7%a8%8b%e3%80%81%e9%87%8d%e5%ae%9a%e5%90%91%e5%92%8c%e7%ae%a1%e9%81%93%e6%8c%87%e4%bb%a4%ef%bc%9axargs%20%e6%8c%87%e4%bb%a4%e7%9a%84%e4%bd%9c%e7%94%a8%e6%98%af%ef%bc%9f.md">07  进程、重定向和管道指令：xargs 指令的作用是？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  用户和权限管理指令： 请简述 Linux 权限划分的原则？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/08%20%20%e7%94%a8%e6%88%b7%e5%92%8c%e6%9d%83%e9%99%90%e7%ae%a1%e7%90%86%e6%8c%87%e4%bb%a4%ef%bc%9a%20%e8%af%b7%e7%ae%80%e8%bf%b0%20Linux%20%e6%9d%83%e9%99%90%e5%88%92%e5%88%86%e7%9a%84%e5%8e%9f%e5%88%99%ef%bc%9f.md">08  用户和权限管理指令： 请简述 Linux 权限划分的原则？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  Linux 中的网络指令：如何查看一个域名有哪些 NS 记录？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/09%20%20Linux%20%e4%b8%ad%e7%9a%84%e7%bd%91%e7%bb%9c%e6%8c%87%e4%bb%a4%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9f%a5%e7%9c%8b%e4%b8%80%e4%b8%aa%e5%9f%9f%e5%90%8d%e6%9c%89%e5%93%aa%e4%ba%9b%20NS%20%e8%ae%b0%e5%bd%95%ef%bc%9f.md">09  Linux 中的网络指令：如何查看一个域名有哪些 NS 记录？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  软件的安装： 编译安装和包管理器安装有什么优势和劣势？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/10%20%20%e8%bd%af%e4%bb%b6%e7%9a%84%e5%ae%89%e8%a3%85%ef%bc%9a%20%e7%bc%96%e8%af%91%e5%ae%89%e8%a3%85%e5%92%8c%e5%8c%85%e7%ae%a1%e7%90%86%e5%99%a8%e5%ae%89%e8%a3%85%e6%9c%89%e4%bb%80%e4%b9%88%e4%bc%98%e5%8a%bf%e5%92%8c%e5%8a%a3%e5%8a%bf%ef%bc%9f.md">10  软件的安装： 编译安装和包管理器安装有什么优势和劣势？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  高级技巧之日志分析：利用 Linux 指令分析 Web 日志.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/11%20%20%e9%ab%98%e7%ba%a7%e6%8a%80%e5%b7%a7%e4%b9%8b%e6%97%a5%e5%bf%97%e5%88%86%e6%9e%90%ef%bc%9a%e5%88%a9%e7%94%a8%20Linux%20%e6%8c%87%e4%bb%a4%e5%88%86%e6%9e%90%20Web%20%e6%97%a5%e5%bf%97.md">11  高级技巧之日志分析：利用 Linux 指令分析 Web 日志.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  高级技巧之集群部署：利用 Linux 指令同时在多台机器部署程序.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/12%20%20%e9%ab%98%e7%ba%a7%e6%8a%80%e5%b7%a7%e4%b9%8b%e9%9b%86%e7%be%a4%e9%83%a8%e7%bd%b2%ef%bc%9a%e5%88%a9%e7%94%a8%20Linux%20%e6%8c%87%e4%bb%a4%e5%90%8c%e6%97%b6%e5%9c%a8%e5%a4%9a%e5%8f%b0%e6%9c%ba%e5%99%a8%e9%83%a8%e7%bd%b2%e7%a8%8b%e5%ba%8f.md">12  高级技巧之集群部署：利用 Linux 指令同时在多台机器部署程序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 (1)加餐  练习题详解（二）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/12%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e4%ba%8c%ef%bc%89.md">12 (1)加餐  练习题详解（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  操作系统内核：Linux 内核和 Windows 内核有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/13%20%20%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%86%85%e6%a0%b8%ef%bc%9aLinux%20%e5%86%85%e6%a0%b8%e5%92%8c%20Windows%20%e5%86%85%e6%a0%b8%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">13  操作系统内核：Linux 内核和 Windows 内核有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  用户态和内核态：用户态线程和内核态线程有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/14%20%20%e7%94%a8%e6%88%b7%e6%80%81%e5%92%8c%e5%86%85%e6%a0%b8%e6%80%81%ef%bc%9a%e7%94%a8%e6%88%b7%e6%80%81%e7%ba%bf%e7%a8%8b%e5%92%8c%e5%86%85%e6%a0%b8%e6%80%81%e7%ba%bf%e7%a8%8b%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">14  用户态和内核态：用户态线程和内核态线程有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  中断和中断向量：Javajs 等语言为什么可以捕获到键盘输入？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/15%20%20%e4%b8%ad%e6%96%ad%e5%92%8c%e4%b8%ad%e6%96%ad%e5%90%91%e9%87%8f%ef%bc%9aJavajs%20%e7%ad%89%e8%af%ad%e8%a8%80%e4%b8%ba%e4%bb%80%e4%b9%88%e5%8f%af%e4%bb%a5%e6%8d%95%e8%8e%b7%e5%88%b0%e9%94%ae%e7%9b%98%e8%be%93%e5%85%a5%ef%bc%9f.md">15  中断和中断向量：Javajs 等语言为什么可以捕获到键盘输入？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  WinMacUnixLinux 的区别和联系：为什么 Debian 漏洞排名第一还这么多人用？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/16%20%20WinMacUnixLinux%20%e7%9a%84%e5%8c%ba%e5%88%ab%e5%92%8c%e8%81%94%e7%b3%bb%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%20Debian%20%e6%bc%8f%e6%b4%9e%e6%8e%92%e5%90%8d%e7%ac%ac%e4%b8%80%e8%bf%98%e8%bf%99%e4%b9%88%e5%a4%9a%e4%ba%ba%e7%94%a8%ef%bc%9f.md">16  WinMacUnixLinux 的区别和联系：为什么 Debian 漏洞排名第一还这么多人用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 (1)加餐  练习题详解（三）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/16%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e4%b8%89%ef%bc%89.md">16 (1)加餐  练习题详解（三）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  进程和线程：进程的开销比线程大在了哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/17%20%20%e8%bf%9b%e7%a8%8b%e5%92%8c%e7%ba%bf%e7%a8%8b%ef%bc%9a%e8%bf%9b%e7%a8%8b%e7%9a%84%e5%bc%80%e9%94%80%e6%af%94%e7%ba%bf%e7%a8%8b%e5%a4%a7%e5%9c%a8%e4%ba%86%e5%93%aa%e9%87%8c%ef%bc%9f.md">17  进程和线程：进程的开销比线程大在了哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  锁、信号量和分布式锁：如何控制同一时间只有 2 个线程运行？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/18%20%20%e9%94%81%e3%80%81%e4%bf%a1%e5%8f%b7%e9%87%8f%e5%92%8c%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8e%a7%e5%88%b6%e5%90%8c%e4%b8%80%e6%97%b6%e9%97%b4%e5%8f%aa%e6%9c%89%202%20%e4%b8%aa%e7%ba%bf%e7%a8%8b%e8%bf%90%e8%a1%8c%ef%bc%9f.md">18  锁、信号量和分布式锁：如何控制同一时间只有 2 个线程运行？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  乐观锁、区块链：除了上锁还有哪些并发控制方法？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/19%20%20%e4%b9%90%e8%a7%82%e9%94%81%e3%80%81%e5%8c%ba%e5%9d%97%e9%93%be%ef%bc%9a%e9%99%a4%e4%ba%86%e4%b8%8a%e9%94%81%e8%bf%98%e6%9c%89%e5%93%aa%e4%ba%9b%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e6%96%b9%e6%b3%95%ef%bc%9f.md">19  乐观锁、区块链：除了上锁还有哪些并发控制方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  线程的调度：线程调度都有哪些方法？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/20%20%20%e7%ba%bf%e7%a8%8b%e7%9a%84%e8%b0%83%e5%ba%a6%ef%bc%9a%e7%ba%bf%e7%a8%8b%e8%b0%83%e5%ba%a6%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b9%e6%b3%95%ef%bc%9f.md">20  线程的调度：线程调度都有哪些方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  哲学家就餐问题：什么情况下会触发饥饿和死锁？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/21%20%20%e5%93%b2%e5%ad%a6%e5%ae%b6%e5%b0%b1%e9%a4%90%e9%97%ae%e9%a2%98%ef%bc%9a%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e4%bc%9a%e8%a7%a6%e5%8f%91%e9%a5%a5%e9%a5%bf%e5%92%8c%e6%ad%bb%e9%94%81%ef%bc%9f.md">21  哲学家就餐问题：什么情况下会触发饥饿和死锁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  进程间通信： 进程间通信都有哪些方法？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/22%20%20%e8%bf%9b%e7%a8%8b%e9%97%b4%e9%80%9a%e4%bf%a1%ef%bc%9a%20%e8%bf%9b%e7%a8%8b%e9%97%b4%e9%80%9a%e4%bf%a1%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b9%e6%b3%95%ef%bc%9f.md">22  进程间通信： 进程间通信都有哪些方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  分析服务的特性：我的服务应该开多少个进程、多少个线程？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/23%20%20%e5%88%86%e6%9e%90%e6%9c%8d%e5%8a%a1%e7%9a%84%e7%89%b9%e6%80%a7%ef%bc%9a%e6%88%91%e7%9a%84%e6%9c%8d%e5%8a%a1%e5%ba%94%e8%af%a5%e5%bc%80%e5%a4%9a%e5%b0%91%e4%b8%aa%e8%bf%9b%e7%a8%8b%e3%80%81%e5%a4%9a%e5%b0%91%e4%b8%aa%e7%ba%bf%e7%a8%8b%ef%bc%9f.md">23  分析服务的特性：我的服务应该开多少个进程、多少个线程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 (1)加餐  练习题详解（四）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/23%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e5%9b%9b%ef%bc%89.md">23 (1)加餐  练习题详解（四）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  虚拟内存 ：一个程序最多能使用多少内存？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/24%20%20%e8%99%9a%e6%8b%9f%e5%86%85%e5%ad%98%20%ef%bc%9a%e4%b8%80%e4%b8%aa%e7%a8%8b%e5%ba%8f%e6%9c%80%e5%a4%9a%e8%83%bd%e4%bd%bf%e7%94%a8%e5%a4%9a%e5%b0%91%e5%86%85%e5%ad%98%ef%bc%9f.md">24  虚拟内存 ：一个程序最多能使用多少内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  内存管理单元： 什么情况下使用大内存分页？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/25%20%20%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%e5%8d%95%e5%85%83%ef%bc%9a%20%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e4%bd%bf%e7%94%a8%e5%a4%a7%e5%86%85%e5%ad%98%e5%88%86%e9%a1%b5%ef%bc%9f.md">25  内存管理单元： 什么情况下使用大内存分页？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  缓存置换算法： LRU 用什么数据结构实现更合理？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/26%20%20%e7%bc%93%e5%ad%98%e7%bd%ae%e6%8d%a2%e7%ae%97%e6%b3%95%ef%bc%9a%20LRU%20%e7%94%a8%e4%bb%80%e4%b9%88%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e5%ae%9e%e7%8e%b0%e6%9b%b4%e5%90%88%e7%90%86%ef%bc%9f.md">26  缓存置换算法： LRU 用什么数据结构实现更合理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  内存回收上篇：如何解决内存的循环引用问题？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/27%20%20%e5%86%85%e5%ad%98%e5%9b%9e%e6%94%b6%e4%b8%8a%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e5%86%85%e5%ad%98%e7%9a%84%e5%be%aa%e7%8e%af%e5%bc%95%e7%94%a8%e9%97%ae%e9%a2%98%ef%bc%9f.md">27  内存回收上篇：如何解决内存的循环引用问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  内存回收下篇：三色标记-清除算法是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/28%20%20%e5%86%85%e5%ad%98%e5%9b%9e%e6%94%b6%e4%b8%8b%e7%af%87%ef%bc%9a%e4%b8%89%e8%89%b2%e6%a0%87%e8%ae%b0-%e6%b8%85%e9%99%a4%e7%ae%97%e6%b3%95%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">28  内存回收下篇：三色标记-清除算法是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 (1)加餐  练习题详解（五）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/28%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e4%ba%94%ef%bc%89.md">28 (1)加餐  练习题详解（五）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  Linux 下的各个目录有什么作用？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/29%20%20Linux%20%e4%b8%8b%e7%9a%84%e5%90%84%e4%b8%aa%e7%9b%ae%e5%bd%95%e6%9c%89%e4%bb%80%e4%b9%88%e4%bd%9c%e7%94%a8%ef%bc%9f.md">29  Linux 下的各个目录有什么作用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  文件系统的底层实现：FAT、NTFS 和 Ext3 有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/30%20%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e7%9a%84%e5%ba%95%e5%b1%82%e5%ae%9e%e7%8e%b0%ef%bc%9aFAT%e3%80%81NTFS%20%e5%92%8c%20Ext3%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">30  文件系统的底层实现：FAT、NTFS 和 Ext3 有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  数据库文件系统实例：MySQL 中 B 树和 B&#43; 树有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/31%20%20%e6%95%b0%e6%8d%ae%e5%ba%93%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%ae%9e%e4%be%8b%ef%bc%9aMySQL%20%e4%b8%ad%20B%20%e6%a0%91%e5%92%8c%20B&#43;%20%e6%a0%91%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">31  数据库文件系统实例：MySQL 中 B 树和 B&#43; 树有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  HDFS 介绍：分布式文件系统是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/32%20%20HDFS%20%e4%bb%8b%e7%bb%8d%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">32  HDFS 介绍：分布式文件系统是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 (1)加餐  练习题详解（六）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/32%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e5%85%ad%ef%bc%89.md">32 (1)加餐  练习题详解（六）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  互联网协议群（TCPIP）：多路复用是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/33%20%20%e4%ba%92%e8%81%94%e7%bd%91%e5%8d%8f%e8%ae%ae%e7%be%a4%ef%bc%88TCPIP%ef%bc%89%ef%bc%9a%e5%a4%9a%e8%b7%af%e5%a4%8d%e7%94%a8%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">33  互联网协议群（TCPIP）：多路复用是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  UDP 协议：UDP 和 TCP 相比快在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/34%20%20UDP%20%e5%8d%8f%e8%ae%ae%ef%bc%9aUDP%20%e5%92%8c%20TCP%20%e7%9b%b8%e6%af%94%e5%bf%ab%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">34  UDP 协议：UDP 和 TCP 相比快在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  Linux 的 IO 模式：selectpollepoll 有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/35%20%20Linux%20%e7%9a%84%20IO%20%e6%a8%a1%e5%bc%8f%ef%bc%9aselectpollepoll%20%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">35  Linux 的 IO 模式：selectpollepoll 有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  公私钥体系和网络安全：什么是中间人攻击？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/36%20%20%e5%85%ac%e7%a7%81%e9%92%a5%e4%bd%93%e7%b3%bb%e5%92%8c%e7%bd%91%e7%bb%9c%e5%ae%89%e5%85%a8%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%af%e4%b8%ad%e9%97%b4%e4%ba%ba%e6%94%bb%e5%87%bb%ef%bc%9f.md">36  公私钥体系和网络安全：什么是中间人攻击？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 (1)加餐  练习题详解（七）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/36%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e4%b8%83%ef%bc%89.md">36 (1)加餐  练习题详解（七）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  虚拟化技术介绍：VMware 和 Docker 的区别？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/37%20%20%e8%99%9a%e6%8b%9f%e5%8c%96%e6%8a%80%e6%9c%af%e4%bb%8b%e7%bb%8d%ef%bc%9aVMware%20%e5%92%8c%20Docker%20%e7%9a%84%e5%8c%ba%e5%88%ab%ef%bc%9f.md">37  虚拟化技术介绍：VMware 和 Docker 的区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  容器编排技术：如何利用 K8s 和 Docker Swarm 管理微服务？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/38%20%20%e5%ae%b9%e5%99%a8%e7%bc%96%e6%8e%92%e6%8a%80%e6%9c%af%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20K8s%20%e5%92%8c%20Docker%20Swarm%20%e7%ae%a1%e7%90%86%e5%be%ae%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">38  容器编排技术：如何利用 K8s 和 Docker Swarm 管理微服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  Linux 架构优秀在哪里.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/39%20%20Linux%20%e6%9e%b6%e6%9e%84%e4%bc%98%e7%a7%80%e5%9c%a8%e5%93%aa%e9%87%8c.md">39  Linux 架构优秀在哪里.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  商业操作系统：电商操作系统是不是一个噱头？.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/40%20%20%e5%95%86%e4%b8%9a%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%ef%bc%9a%e7%94%b5%e5%95%86%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e6%98%af%e4%b8%8d%e6%98%af%e4%b8%80%e4%b8%aa%e5%99%b1%e5%a4%b4%ef%bc%9f.md">40  商业操作系统：电商操作系统是不是一个噱头？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 (1)加餐  练习题详解（八）.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/40%20%281%29%e5%8a%a0%e9%a4%90%20%20%e7%bb%83%e4%b9%a0%e9%a2%98%e8%af%a6%e8%a7%a3%ef%bc%88%e5%85%ab%ef%bc%89.md">40 (1)加餐  练习题详解（八）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 结束语  论程序员的发展——信仰、选择和博弈.md" href="/%e4%b8%93%e6%a0%8f/%e9%87%8d%e5%ad%a6%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f-%e5%ae%8c/41%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e8%ae%ba%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e5%8f%91%e5%b1%95%e2%80%94%e2%80%94%e4%bf%a1%e4%bb%b0%e3%80%81%e9%80%89%e6%8b%a9%e5%92%8c%e5%8d%9a%e5%bc%88.md">41 结束语  论程序员的发展——信仰、选择和博弈.md</a>
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
                            <h1 id="title" data-id="31  数据库文件系统实例：MySQL 中 B 树和 B&#43; 树有什么区别？" class="title">31  数据库文件系统实例：MySQL 中 B 树和 B&#43; 树有什么区别？</h1>
                            <div><p><strong>这一讲给你带来的关联面试题是：MySQL 中 B 树和 B+ 树的区别？</strong></p>

<p>B 树和 B+ 树是两种数据结构（关于它们的名字为什么以 B 开头，因为众说纷纭，本讲我就不介绍了），构建了磁盘中的高速索引结构，因此不仅 MySQL 在用，MongoDB、Oracle 等也在用，基本属于数据库的标配常规操作。</p>

<p>数据库要经常和磁盘与内存打交道，为了提升性能，通常需要自己去构建类似文件系统的结构。这一讲的内容有限，我只是先带你入一个门，如果你感兴趣后续可以自己深入学习。下面我们一起来探讨数据库如何利用磁盘空间设计索引。</p>

<h3 id="行存储和列存储">行存储和列存储</h3>

<p>在学习构建磁盘数据的索引结构前，我们先通过行存储、列存储的学习来了解一些基本的存储概念，帮助你建立一个基本的认知。</p>

<p>目前数据库存储一张表格主要是行存储（Row Storage）和列存储（Column Storage）两种存储方式。行存储将表格看作一个个记录，每个记录是一行。以包含订单号、金额、下单时间 3 项的表为例，行存储如下图所示：</p>

<p><img src="assets/CgpVE1_sZ-GAVZtxAABQMz4nr2o681.png" alt="Drawing 1.png" /></p>

<p>如上图所示，在计算机中没有真正的行的概念。行存储本质就是数据一个接着一个排列，一行数据后面马上跟着另一行数据。如果订单表很大，一个磁盘块（Block）存不下，那么实际上就是每个块存储一定的行数。 类似下图这样的结构：</p>

<p><img src="assets/Ciqc1F_sZ-qAWYp7AABVE_FTHT8211.png" alt="Drawing 3.png" /></p>

<p>行存储更新一行的操作，往往可以在一个块（Block）中进行。而查询数据，聚合数据（比如求 4 月份的订单数），往往需要跨块（Block）。因此，<strong>行存储优点很明显，更新快、单条记录的数据集中，适合事务。但缺点也很明显，查询慢</strong>。</p>

<p>还有一种表格的存储方式是列存储（Column Storage），列存储中数据是一列一列存的。还以订单表为例，如下图所示：</p>

<p><img src="assets/CgqCHl_sZ_GAV5BXAABYyBf77uc003.png" alt="Drawing 5.png" /></p>

<p>你可以看到订单号在一起、姓名在一起、时间在一起、金额也在一起——每个列的数据都聚集在一起。乍一看这样的结构很低效，比如说你想取出第一条订单，需要取第 1 列的第 1 个数据<code>1001</code>，然后取第 2 列的第 1 个数据<code>小明</code>，以此类推，需要 4 次磁盘读取。特别是更新某一条记录的时候，需要更新多处，速度很慢。那么列存储优势在哪里呢？<strong>优势其实是在查询和聚合运算</strong>。</p>

<p>在列存储中同一列数据总是存放在一起，比如要查找某个时间段，很有可能在一个块中就可以找到，因为时间是集中存储的。假设磁盘块的大小是 4KB，一条记录是 100 字节， 那么 4KB 可以存 40 条记录；但是存储时间戳只需要一个 32 位整数，4KB 可以存储 1000 个时间。更关键的是，我们可以把一片连续的硬盘空间通过 DMA 技术直接映射到内存，这样就大大减少了搜索需要的时间。所以有时候在行存储需要几分钟的搜索操作，在列存储中只需几秒钟就可以完成。</p>

<p><strong>总结一下，行存储、列存储，最终都需要把数据存到磁盘块</strong>。行存储记录一个接着一个，列存储一列接着一列。前面我们提到行存储适合更新及事务处理，更新好理解，因为一个订单可以在相同的 Block 中更新，那么为什么适合事务呢？</p>

<p>其实适合不适合是相对的，说行存储适合是因为列存储非常不适合事务。试想一下，你更新一个表的若干个数据，如果要在不同块中更新，就可能产生多次更新操作。更新次数越多，保证一致性越麻烦。在单机环境我们可以上锁，可以用阻塞队列，可以用屏障……但是分布式场景中保证一致性（特别是强一致性）开销很大。因此我们说行存储适合事务，而列存储不适合。</p>

<h3 id="索引">索引</h3>

<p><strong>接下来，我们在行存储、列存储的基础上，讨论如何创建一些更高效的查询结构，这种结构通常称为索引</strong>。我们经常会遇到根据一个订单编号查订单的情况，比如说<code>select * from order where id=1000000</code>，这个时候就需要用到索引。而下面我将试图通过二分查找的场景，和你一起讨论<strong>索引是什么。</strong></p>

<p>在亿级的订单 ID 中查找某个编号，很容易想到二分查找。要理解二分查找，最需要关心的是算法的进步机制。这个算法每进行一次查找，都会让问题的规模减半。当然，也有场景限制，二分查找只能应用在排序好的数据上。</p>

<p>比如我们要在下面排序好的数组中查找 3：</p>

<p>1,3,5,8,11,12,15,19,21,25</p>

<p>数组中一共有 10 个元素，因此我们第一次查找从数组正中间的元素找起。如果数组正中间有两个元素，就取左边的那个——对于这个例子是 11。我们比较 11 和 3 的值，因为 11 大于 3，因此可以排除包括 11 在内的所有 11 右边的元素。相当于我们通过一次运算将数据的规模减半。假设我们有 240 （1T 数据）个元素需要查询（规模已经相当大了，万亿级别），用二分查找只需要 40 次运算。</p>

<p>所以按照这个思路，我们需要做的是将数据按照订单 ID 排好序，查询的时候就可以应用二分查找了。而且按照二分查找的思路，也可以进行范围查找。比如要查找 [a,b] 之间的数据，可以先通过二分查找找到 a 的序号，再二分找到 b 的序号，两个序号之间的数据就是目标结果。</p>

<p>但是直接在原始数据上排序，我们可能会把数据弄乱，常规做法是设计冗余数据描述这种排序关系——这就是索引。下面我通过一个简单的例子告诉你为什么不能在原始数据上直接排序。</p>

<p>假设我们有一个订单表，里面有订单 ID 和金额。使用列存储做演示如下：</p>

<p>订单 ID 列：</p>

<p>10005 10001 ……</p>

<p>订单金额列：</p>

<p>99.00 100.00 ……</p>

<p>可以看到，订单（10001）是第 2 个订单。但是进行排序后，订单（10001）会到第 1 个位置。这样会弄乱订单 ID（10001）和 金额（100.00）对应的关系。</p>

<p>因此我们必须用空间换时间，额外将订单列拷贝一份排序：</p>

<p>10001，2，10005， 1</p>

<p>以上这种<strong>专门用来进行数据查询的额外数据，就是索引</strong>。<strong>索引中的一个数据，也称作索引条目</strong>。上面的索引条目一个接着一个，每个索引条目是 &lt;订单 ID, 序号&gt; 的二元组。</p>

<p>如果你考虑是行存储（比如 MySQL），那么依然可以生成上面的索引，订单 ID 和序号（行号）关联。如果有多个索引，就需要创造多个上面的数据结构。如果有复合索引，比如 &lt;订单状态、日期、序号&gt; 作为一个索引条目，其实就是先按照订单状态，再按照日期排序的索引。</p>

<p>所以复合索引，无非就是多消耗一些空间，排序维度多一些。而且你可以看出复合索引和单列索引完全是独立关系，所以我们可以认为每创造一组索引，就创造了一份冗余的数据。也创造了一种特别的查询方式。关于索引还有很多有趣的知识，我们先介绍这些，如果感兴趣可以自己查资料深挖。</p>

<p>接下来，请分析一个非常<strong>核心的问题：上面的索引是一个连续的、从小到大的索引，那么应不应该使用这种从小到大排序的索引呢</strong>？例如，我们需要查询订单，就事先创建另一个根据订单 ID 从小到大排序的索引，当用户查找某个订单的时候，无论是行存储、还是列存储，我们就用二分查找查询对应的索引条目。这种方式，我们姑且称为线性排序索引——看似很不错的一个方式，但是并不是非常好的一种做法，请看我接下来的讨论。</p>

<h3 id="二叉搜索树">二叉搜索树</h3>

<p>线性排序的数据虽然好用，但是插入新元素的时候性能太低。如果是内存操作，插入一个元素，需要将这个元素之后的所有元素后移一位。但如果这个操作发生在磁盘中呢？这必然是灾难性的。因为磁盘的速度比内存慢至少 10-1000 倍，如果是机械硬盘可能慢几十万到百万倍。</p>

<p>所以我们不能用一种线性结构将磁盘排序。那么树呢？ 比如二叉搜索树（Binary Serach Tree）行不行呢？利用磁盘的空间形成一个二叉搜索树，例如将订单 ID 作为二叉搜索树的 Key。</p>

<p>如下图所示，二叉搜索树的特点是一个节点的左子树的所有节点都小于这个节点，右子树的所有节点都大于这个节点。而且，因为索引条目较少，确实可以考虑在查询的时候，先将足够大的树导入内存，然后再进行搜索。搜索的算法是递归的，与二分查找非常类似，每次计算可以将问题规模减半。当然，具体有多少数据可以导入内存，受实际可以使用的内存数量的限制。</p>

<p><img src="assets/Ciqc1F_saAWAQLD2AAEHA2b-QK4732.png" alt="Drawing 7.png" /></p>

<p>在上面的二叉搜索树中，每个节点的数据分成 Key 和 Value。Key 就是索引值，比如订单 ID 创建索引，那么 Key 就是订单 ID。值中至少需要序号（对行存储也就是行号）。这样，如果们想找 18 对应的行，就可以先通过二叉搜索树找到对应的行号，然后再去对应的行读取数据。</p>

<p><img src="assets/CgqCHl_saAyAIC_4AAFHdR111Pg040.png" alt="Drawing 9.png" /></p>

<p>二叉搜索树是一个天生的二分查找结构，每次查找都可以减少一半的问题规模。而且二叉搜索树解决了插入新节点的问题，因为二叉搜索树是一个跳跃结构，不必在内存中连续排列。这样在插入的时候，新节点可以放在任何位置，不会像线性结构那样插入一个元素，所有元素都需要向后排列。</p>

<p><strong>那么回到本质问题，在使用磁盘的时候，二叉搜索树是不是一种合理的查询结构</strong>？</p>

<p>当然还不算，因此还需要继续优化我们的算法。二叉搜索树，在内存中是一个高效的数据结构。这是因为内存速度快，不仅可以随机存取，还可以高频操作。注意 CPU 缓存的穿透率只有 5% 左右，也就是 95% 的操作是在更快的 CPU 缓存中执行的。而且即便穿透，内存操作也是在纳秒级别可以完成。</p>

<p>但是，这个逻辑在磁盘中是不存在的，磁盘的速度慢太多了。我们可以尝试把尽可能多的二叉搜索树读入磁盘，但是如果数据量大，只能读入一部分呢？因此我们还需要继续改进算法。</p>

<h3 id="b-树和-b-树">B 树和 B+ 树</h3>

<p>二叉搜索树解决了连续结构插入新元素开销很大的问题，同时又保持着天然的二分结构。但是，当需要索引的数据量很大，无法在一个磁盘 Block 中存下整棵二叉搜索树的时候。每一次递归向下的搜索，实际都是读取不同的磁盘块。这个时候二叉搜索树的开销很大。</p>

<p>试想一个一万亿条订单的表，进行 40 次查找找到答案，在内存中不是问题，要考虑到 CPU 缓存有 90% 以上的命中率（当然前提是内存足够大）。通常情况下我们没有这么大的内存空间，如果 40 次查找发生在磁盘上，也是非常耗时的。那么有没有更好的方案呢？</p>

<p><strong>一个更好的方案，就是继续沿用树状结构，利用好磁盘的分块让每个节点多一些数据，并且允许子节点也多一些，树就会更矮。因为树的高度决定了搜索的次数</strong>。</p>

<p><img src="assets/Ciqc1F_saBaAXK5-AAFO9nLONPo957.png" alt="Drawing 11.png" /></p>

<p><strong>上图中我们构造的树被称为 B 树</strong>（<strong>B-Tree</strong>），开头说过，B 这个字母具体是哪个单词或者人名的缩写，至今有争议，具体你可以查查资料。</p>

<p>B-Tree 是一种递归的搜索结构，与二叉搜索树非常类似。不同的是，B 树中的父节点中的数据会对子树进行区段分割。比如上图中节点 1 有 3 个子节点，并用数字 9,30 对子树的区间进行了划分。</p>

<p>上图中的 B 树是一个 3-4 B 树，3 指的是每个非叶子节点允许最大 3 个索引，4 指的是每个节点最多允许 4 个子节点，4 也指每个叶子节点可以存 4 个索引。上面只是一个例子，在实际的操作中，子节点有几十个、甚至上百个索引也很常见，因为我们希望树变矮，好减少磁盘操作。</p>

<p>B 树的每个节点是一个索引条目（例如：一个 &lt;订单 ID，序号&gt; 的组合），如果是行数据库可以索引到一条存储在磁盘上的记录。</p>

<h4 id="继承-b-树-b-树">继承 B 树：B+ 树</h4>

<p>为了达到最高的效率，实战中我们往往使用的是一种继承于 B 树设计的结构，称为 B+ 树。B+ 树只有叶子节点才映射数据，下图中是对 B 树设计的一种改进，节点 1 为冗余节点，它不存储数据，只划定子树数据的范围。你可以看到节点 1 的索引 Key：12 和 30，在节点 3 和 4 中也有一份。</p>

<p><img src="assets/Ciqc1F_saCKADMzAAAEHeJ2-HvI282.png" alt="Drawing 13.png" /></p>

<h4 id="树的形成-插入">树的形成：插入</h4>

<p>下面我以一棵 2-3 B+ 树来演示 B+ 树的插入过程。2 指的是 B+ 树每个非叶子节点允许 2 个数据，叶子节点最多允许 3 个索引，每个节点允许最多 3 个子节点。我们要在 2-3 B+ 树中依次插入 3,6,9,12,19,15,26,8,30。下图是演示：
<img src="assets/Ciqc1F_yuGaAe34hAIFHQoSI9GI738.gif" alt="10.gif" /></p>

<p>插入 3,6,9 过程很简单，都写入一个节点即可，因为叶子节点最多允许每个 3 个索引。接下来我们插入 12，会发生一次过载，然后节点就需要拆分，这个时候按照 B+ 树的设计会产生冗余节点。</p>

<p><img src="assets/Ciqc1F_yt1uAUp9PAABNVsx4HM8009.png" alt="6.png" /></p>

<p>然后插入 15 非常简单，直接加入即可：</p>

<p><img src="assets/CgqCHl_yt36ANCL0AABO--TyyME355.png" alt="7.png" /></p>

<p>接下来插入 19， 这个时候下图中红色部分发生过载：</p>

<p><img src="assets/Ciqc1F_yt6GAarKkAABRlrOx8wQ337.png" alt="8.png" /></p>

<p>因此需要拆分节点数据，我们从中间把红色的节点拆开，15 作为冗余的索引写入父节点，就形成下图的情况：</p>

<p><img src="assets/Ciqc1F_yt9aAWPFwAABWeUsN-3w700.png" alt="9.png" /></p>

<p>接着插入 26， 写入到对应位置即可。</p>

<p><img src="assets/Cip5yF_ytRmAXVtMAABY0urP73M758.png" alt="1.png" /></p>

<p>接下来，插入 8 到对应位置即可。</p>

<p><img src="assets/Ciqc1F_ytbOABwB3AABY1XWgGI4978.png" alt="2.png" /></p>

<p>然后我们插入 30，此时右边节点发生过载：</p>

<p><img src="assets/CgpVE1_yteaATRGyAABdLxm53J0846.png" alt="3.png" /></p>

<p>解决完一次过载问题之后，因为 26 会浮上去，根节点又发生了过载：</p>

<p><img src="assets/CgpVE1_ytjSAduc9AABnqqExax0913.png" alt="4.png" /></p>

<p>再次解决过载，拆分红色部分，得到最后结果：</p>

<p><img src="assets/Cip5yF_ytoaACk0NAAB41Q4jOn8459.png" alt="5.png" /></p>

<p>在上述过程中，B+ 树始终可以保持平衡状态，而且所有叶子节点都在同一层级。更复杂的数学证明，我就不在这里讲解了。不过建议对算法感兴趣对同学，可以学习《算法导论》中关于树的部分。</p>

<h4 id="插入和删除效率">插入和删除效率</h4>

<p>B+ 树有大量的冗余节点，比如删除一个节点的时候，可以直接从叶子节点中删除，甚至可以不动非叶子节点。这样删除非常快。B 树则不同，B 树没有冗余节点，删除节点的时候非常复杂。比如删除根节点中的数据，可能涉及复杂的树的变形。</p>

<p>B+ 树的插入也是一样，有冗余节点，插入可能存在节点的拆分（如果节点饱和），但是最多只涉及树的一条路径。而且 B+ 树会自动平衡，不需要更多复杂的算法，类似红黑树的旋转操作等。</p>

<p>因此，<strong>B+ 树的插入和删除效率更高</strong>。</p>

<h4 id="搜索-链表的作用">搜索：链表的作用</h4>

<p>B 树和 B+ 树搜索原理基本一致。先从根节点查找，然后对比目标数据的范围，最后递归的进入子节点查找。</p>

<p>你可能会注意到，B+ 树所有叶子节点间还有一个链表进行连接。这种设计对范围查找非常有帮助，比如说我们想知道 1 月 20 日和 1 月 22 日之间的订单，这个时候可以先查找到 1 月 20 日所在的叶子节点，然后利用链表向右遍历，直到找到 1 月22 日的节点。这样我们就进一步节省搜索需要的时间。</p>

<h3 id="总结">总结</h3>

<p>这一讲我们学习了在数据库中如何利用文件系统造索引。无论是行存储还是列存储，构造索引的过程都是类似的。索引有很多做法，除了 B+ 树，还有 HashTable、倒排表等。<strong>如果是存储海量数据的数据库，我们的思考点需要放在 I/O 的效率上</strong>。<strong>如果把今天的知识放到分布式数据库上，那除了需要节省磁盘读写还需要节省网络 I/O</strong>。</p>

<p><strong>那么通过这一讲的学习，你现在可以尝试来回答本讲关联的面试题目：MySQL 中的 B 树和 B+ 树有什么区别？</strong></p>

<p>【<strong>解析</strong>】B+ 树继承于 B 树，都限定了节点中数据数目和子节点的数目。B 树所有节点都可以映射数据，B+ 树只有叶子节点可以映射数据。</p>

<p>单独看这部分设计，看不出 B+ 树的优势。为了只有叶子节点可以映射数据，B+ 树创造了很多冗余的索引（所有非叶子节点都是冗余索引），这些冗余索引让 B+ 树在插入、删除的效率都更高，而且可以自动平衡，因此 B+ 树的所有叶子节点总是在一个层级上。所以 B+ 树可以用一条链表串联所有的叶子节点，也就是索引数据，这让 B+ 树的范围查找和聚合运算更快。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#fc909090c5c8cdcdcccbbc9b919d9590d29f9391" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a0dec6e4c888b',t:'MTczNDEzNTkwMi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
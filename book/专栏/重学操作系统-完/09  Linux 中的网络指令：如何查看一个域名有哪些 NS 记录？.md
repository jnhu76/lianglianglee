<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;&#32;Linux&#32;中的网络指令：如何查看一个域名有哪些&#32;NS&#32;记录？>
        <link rel="icon" href="/static/favicon.png">
        <title>09  Linux 中的网络指令：如何查看一个域名有哪些 NS 记录？ </title>
        
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
                            <h1 id="title" data-id="09  Linux 中的网络指令：如何查看一个域名有哪些 NS 记录？" class="title">09  Linux 中的网络指令：如何查看一个域名有哪些 NS 记录？</h1>
                            <div><p><strong>我看到过一道关于 Linux 指令的面试题：如何查看一个域名有哪些 NS 记录</strong>？</p>

<p>这类题目是根据一个场景，考察一件具体的事情如何处理。虽然你可以通过查资料找到解决方案，但是，这类问题在面试中还是有必要穿插一下，用于确定求职者技能是否熟练、经验是否丰富。特别是计算机网络相关的指令，平时在远程操作、开发、联调、Debug 线上问题的时候，会经常用到。</p>

<p>Linux 中提供了不少网络相关的指令，因为网络指令比较分散，本课时会从下面几个维度给你介绍，帮助你梳理常用的网络指令：</p>

<ul>
<li>远程操作；</li>
<li>查看本地网络状态；</li>
<li>网络测试；</li>
<li>DNS 查询；</li>
<li>HTTP。</li>
</ul>

<p>这块知识从体系上属于 Linux 指令，同时也关联了很多计算机网络的知识，比如说 TCP/IP 协议、UDP 协议，我会在“<strong>模块七</strong>”为你简要介绍。</p>

<p>如果你对这部分指令背后的网络原理有什么困惑，可以在评论区提问。另外，你也可以关注我将在拉勾教育推出的《<strong>计算机网络</strong>》课程。下面我们开始学习今天的内容。</p>

<h3 id="远程操作指令">远程操作指令</h3>

<p>远程操作指令用得最多的是<code>ssh</code>，<code>ssh</code>指令允许远程登录到目标计算机并进行远程操作和管理。还有一个比较常用的远程指令是<code>scp</code>，<code>scp</code>帮助我们远程传送文件。</p>

<h4 id="ssh-secure-shell">ssh（Secure Shell）</h4>

<p>有一种场景需要远程登录一个 Linux 系统，这时我们会用到<code>ssh</code>指令。比如你想远程登录一台机器，可以使用<code>ssh user@ip</code>的方式，如下图所示：</p>

<p><img src="assets/CgqCHl92j8GAMNHAAAPCrIyhHHk744.png" alt="Drawing 0.png" /></p>

<p>上图中，我在使用<code>ssh</code>指令从机器<code>u1</code>登录我的另一台虚拟机<code>u2</code>。这里<code>u1</code>和<code>u2</code>对应着 IP 地址，是我在<code>/etc/hosts</code>中设置的，如下图所示：</p>

<p><img src="assets/CgqCHl92j8mAIMPdAACTOATTrQM694.png" alt="Drawing 1.png" /></p>

<p><code>/etc/hosts</code>这个文件可以设置 IP 地址对应的域名。我这里是一个小集群，总共有两台机器，因此我设置了方便记忆和操作的名字。</p>

<h4 id="scp">scp</h4>

<p>另一种场景是我需要拷贝一个文件到远程，这时可以使用<code>scp</code>指令，如下图，我使用<code>scp</code>指令将本地计算机的一个文件拷贝到了 ubuntu 虚拟机用户的家目录中。</p>

<p>比如从<code>u1</code>拷贝家目录下的文件<code>a.txt</code>到<code>u2</code>。家目录有一个简写，就是用<code>~</code>。具体指令见下图：</p>

<p><img src="assets/Ciqc1F92j9OADjTcAAPER8w5DNg904.png" alt="Drawing 2.png" /></p>

<p>输入 scp 指令之后会弹出一个提示，要求输入密码，系统验证通过后文件会被成功拷贝。</p>

<h3 id="查看本地网络状态">查看本地网络状态</h3>

<p>如果你想要了解本地的网络状态，比较常用的网络指令是<code>ifconfig</code>和<code>netstat</code>。</p>

<h4 id="ifconfig">ifconfig</h4>

<p>当你想知道本地<code>ip</code>以及本地有哪些网络接口时，就可以使用<code>ifconfig</code>指令。你可以把一个网络接口理解成一个网卡，有时候虚拟机会装虚拟网卡，虚拟网卡是用软件模拟的网卡。</p>

<p>比如：VMware 为每个虚拟机创造一个虚拟网卡，通过虚拟网卡接入虚拟网络。当然物理机也可以接入虚拟网络，它可以通过虚拟网络向虚拟机的虚拟网卡上发送信息。</p>

<p>下图是我的 ubuntu 虚拟机用 ifconfig 查看网络接口信息。</p>

<p><img src="assets/Ciqc1F92j9yAaioXAAbz00ZJYlw555.png" alt="Drawing 3.png" /></p>

<p>可以看到我的这台 ubuntu 虚拟机一共有 2 个网卡，ens33 和 lo。<code>lo</code>是本地回路（local lookback），发送给<code>lo</code>就相当于发送给本机。<code>ens33</code>是一块连接着真实网络的虚拟网卡。</p>

<h4 id="netstat">netstat</h4>

<p>另一个查看网络状态的场景是想看目前本机的网络使用情况，这个时候可以用<code>netstat</code>。</p>

<p><strong>默认行为</strong></p>

<p>不传任何参数的<code>netstat</code>帮助查询所有的本地 socket，下图是<code>netstat | less</code>的结果。</p>

<p><img src="assets/Ciqc1F92j-aAAZlfAAizLye7uc4727.png" alt="Drawing 4.png" /></p>

<p>如上图，我们看到的是 socket 文件。socket 是网络插槽被抽象成了文件，负责在客户端、服务器之间收发数据。当客户端和服务端发生连接时，客户端和服务端会同时各自生成一个 socket 文件，用于管理这个连接。这里，可以用<code>wc -l</code>数一下有多少个<code>socket</code>。</p>

<p><img src="assets/Ciqc1F92j-2AVEYjAAA8xcVMQzc068.png" alt="Drawing 5.png" /></p>

<p>你可以看到一共有 615 个 socket 文件，因为有很多 socket 在解决进程间的通信。就是将两个进程一个想象成客户端，一个想象成服务端。并不是真的有 600 多个连接着互联网的请求。</p>

<p><strong>查看 TCP 连接</strong></p>

<p>如果想看有哪些 TCP 连接，可以使用<code>netstat -t</code>。比如下面我通过<code>netstat -t</code>看<code>tcp</code>协议的网络情况：</p>

<p><img src="assets/CgqCHl92j_aAbxdlAAEAdzG3a2s636.png" alt="Drawing 6.png" /></p>

<p>这里没有找到连接中的<code>tcp</code>，因为我们这台虚拟机当时没有发生任何的网络连接。因此我们尝试从机器<code>u2</code>（另一台机器）ssh 登录进<code>u1</code>，再看一次：</p>

<p><img src="assets/CgqCHl92kAaAMuMDAAFWQdSNGfk978.png" alt="Drawing 7.png" /></p>

<p>如上图所示，可以看到有一个 TCP 连接了。</p>

<p><strong>查看端口占用</strong></p>

<p>还有一种非常常见的情形，我们想知道某个端口是哪个应用在占用。如下图所示：</p>

<p><img src="assets/Ciqc1F92kBKAHr2RAAEnmEOZ8RM010.png" alt="Drawing 8.png" /></p>

<p>这里我们看到 22 端口被 sshd，也就是远程登录模块被占用了。<code>-n</code>是将一些特殊的端口号用数字显示，<code>-t</code>是指看 TCP 协议，<code>-l</code>是只显示连接中的连接，<code>-p</code>是显示程序名称。</p>

<h3 id="网络测试">网络测试</h3>

<p>当我们需要测试网络延迟、测试服务是否可用时，可能会用到<code>ping</code>和<code>telnet</code>指令。</p>

<h4 id="ping">ping</h4>

<p>想知道本机到某个网站的网络延迟，就可以使用<code>ping</code>指令。如下图所示：</p>

<p><img src="assets/CgqCHl92kB-ARKR5AAP30Xk0nBg068.png" alt="Drawing 9.png" /></p>

<p><code>ping</code>一个网站需要使用 ICMP 协议。因此你可以在上图中看到 icmp 序号。 这里的时间<code>time</code>是往返一次的时间。<code>ttl</code>叫作 time to live，是封包的生存时间。就是说，一个封包从发出就开始倒计时，如果途中超过 128ms，这个包就会被丢弃。如果包被丢弃，就会被算进丢包率。</p>

<p>另外<code>ping</code>还可以帮助我们看到一个网址的 IP 地址。 通过网址获得 IP 地址的过程叫作 DNS Lookup（DNS 查询）。<code>ping</code>利用了 DNS 查询，但是没有显示全部的 DNS 查询结果。</p>

<h4 id="telnet">telnet</h4>

<p>有时候我们想知道本机到某个 IP + 端口的网络是否通畅，也就是想知道对方服务器是否在这个端口上提供了服务。这个时候可以用<code>telnet</code>指令。 如下图所示：</p>

<p><img src="assets/CgqCHl92kCmAcPQzAADcRdxOtdw609.png" alt="Drawing 10.png" /></p>

<p>telnet 执行后会进入一个交互式的界面，比如这个时候，我们输入下图中的文字就可以发送 HTTP 请求了。如果你对 HTTP 协议还不太了解，建议自学一下 HTTP 协议。如果希望和林老师一起学习，可以等待下我之后的《<strong>计算机网络</strong>》专栏。</p>

<p><img src="assets/Ciqc1F92kDKAcYUbAASLFyOyBg4948.png" alt="Drawing 11.png" /></p>

<p>如上图所示，第 5 行的<code>GET</code> 和第 6 行的<code>HOST</code>是我输入的。 拉勾网返回了一个 301 永久跳转。这是因为拉勾网尝试把<code>http</code>协议链接重定向到<code>https</code>。</p>

<h3 id="dns-查询">DNS 查询</h3>

<p>我们排查网络故障时想要进行一次 DNS Lookup，想知道一个网址 DNS 的解析过程。这个时候有多个指令可以用。</p>

<h4 id="host">host</h4>

<p>host 就是一个 DNS 查询工具。比如我们查询拉勾网的 DNS，如下图所示：</p>

<p><img src="assets/Ciqc1F92kD6AOJPQAAGW1va0D9c041.png" alt="Drawing 12.png" /></p>

<p>我们看到拉勾网 <a href="http://www.lagou.comw/" target="_blank">www.lagou.com</a> 是一个别名，它的原名是 lgmain 开头的一个域名，这说明拉勾网有可能在用 CDN 分发主页（关于 CDN，我们《计算机网络》专栏见）。</p>

<p>上图中，可以找到 3 个域名对应的 IP 地址。</p>

<p>如果想追查某种类型的记录，可以使用<code>host -t</code>。比如下图我们追查拉勾的 AAAA 记录，因为拉勾网还没有部署 IPv6，所以没有找到。</p>

<p><img src="assets/CgqCHl92kFWAHIqAAACvpo6qaOs100.png" alt="Drawing 13.png" /></p>

<h4 id="dig">dig</h4>

<p><code>dig</code>指令也是一个做 DNS 查询的。不过<code>dig</code>指令显示的内容更详细。下图是<code>dig</code>拉勾网的结果。</p>

<p><img src="assets/CgqCHl92kGaADOhxAAR-BfryZ5g689.png" alt="Drawing 14.png" /></p>

<p>从结果可以看到<a href="http://www.lagou.c/" target="_blank">www.lagou.com</a> 有一个别名，用 CNAME 记录定义 lgmain 开头的一个域名，然后有 3 条 A 记录，通常这种情况是为了均衡负载或者分发内容。</p>

<h3 id="http-相关">HTTP 相关</h3>

<p>最后我们来说说<code>http</code>协议相关的指令。</p>

<h4 id="curl">curl</h4>

<p>如果要在命令行请求一个网页，或者请求一个接口，可以用<code>curl</code>指令。<code>curl</code>支持很多种协议，比如 LDAP、SMTP、FTP、HTTP 等。</p>

<p>我们可以直接使用 curl 请求一个网址，获取资源，比如我用 curl 直接获取了拉勾网的主页，如下图所示：</p>

<p><img src="assets/Ciqc1F92kG-AJPyrAANJZYQ4u5w784.png" alt="Drawing 15.png" /></p>

<p>如果只想看 HTTP 返回头，可以使用<code>curl -I</code>。</p>

<p>另外<code>curl</code>还可以执行 POST 请求，比如下面这个语句：</p>

<pre><code>curl -d '{&quot;x&quot; : 1}' -H &quot;Content-Type: application/json&quot; -X POST http://localhost:3000/api
</code></pre>

<p>curl在向<code>localhost:3000</code>发送 POST 请求。<code>-d</code>后面跟着要发送的数据， -<code>X</code>后面是用到的 HTTP 方法，<code>-H</code>是指定自定义的请求头。</p>

<h3 id="总结">总结</h3>

<p>这节课我们学习了不少网络相关的 Linux 指令，这些指令是将来开发和调试的强大工具。这里再给你复习一下这些指令：</p>

<ul>
<li>远程登录的 ssh 指令；</li>
<li>远程拷贝文件的 scp 指令；</li>
<li>查看网络接口的 ifconfig 指令；</li>
<li>查看网络状态的 netstat 指令；</li>
<li>测试网络延迟的 ping 指令；</li>
<li>可以交互式调试和服务端的 telnet 指令；</li>
<li>两个 DNS 查询指令 host 和 dig；</li>
<li>可以发送各种请求包括 HTTPS 的 curl 指令。</li>
</ul>

<p><strong>那么通过这节课的学习，你现在可以来回答本节关联的面试题目：如何查看一个域名有哪些 NS 记录了吗？</strong></p>

<p>老规矩，请你先在脑海里构思下给面试官的表述，并把你的思考写在留言区，然后再来看我接下来的分析。</p>

<p><strong>【解析】</strong> host 指令提供了一个<code>-t</code>参数指定需要查找的记录类型。我们可以使用<code>host -t ns {网址}</code>。另外 dig 也提供了同样的能力。如果你感兴趣，还可以使用<code>man</code>对系统进行操作。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b5d9d9d98c8184848582f5d2d8d4dcd99bd6dad8" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a051e2881888b',t:'MTczNDEzNTU0MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
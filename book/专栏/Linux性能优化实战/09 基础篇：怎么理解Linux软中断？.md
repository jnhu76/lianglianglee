<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;基础篇：怎么理解Linux软中断？>
        <link rel="icon" href="/static/favicon.png">
        <title>09 基础篇：怎么理解Linux软中断？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 别再让Linux性能问题成为你的绊脚石.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%88%ab%e5%86%8d%e8%ae%a9Linux%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%e6%88%90%e4%b8%ba%e4%bd%a0%e7%9a%84%e7%bb%8a%e8%84%9a%e7%9f%b3.md">00 开篇词 别再让Linux性能问题成为你的绊脚石.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 如何学习Linux性能优化？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/01%20%e5%a6%82%e4%bd%95%e5%ad%a6%e4%b9%a0Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%9f.md">01 如何学习Linux性能优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 基础篇：到底应该怎么理解“平均负载”？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/02%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e5%88%b0%e5%ba%95%e5%ba%94%e8%af%a5%e6%80%8e%e4%b9%88%e7%90%86%e8%a7%a3%e2%80%9c%e5%b9%b3%e5%9d%87%e8%b4%9f%e8%bd%bd%e2%80%9d%ef%bc%9f.md">02 基础篇：到底应该怎么理解“平均负载”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 基础篇：经常说的 CPU 上下文切换是什么意思？（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/03%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e7%bb%8f%e5%b8%b8%e8%af%b4%e7%9a%84%20CPU%20%e4%b8%8a%e4%b8%8b%e6%96%87%e5%88%87%e6%8d%a2%e6%98%af%e4%bb%80%e4%b9%88%e6%84%8f%e6%80%9d%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">03 基础篇：经常说的 CPU 上下文切换是什么意思？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 基础篇：经常说的 CPU 上下文切换是什么意思？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/04%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e7%bb%8f%e5%b8%b8%e8%af%b4%e7%9a%84%20CPU%20%e4%b8%8a%e4%b8%8b%e6%96%87%e5%88%87%e6%8d%a2%e6%98%af%e4%bb%80%e4%b9%88%e6%84%8f%e6%80%9d%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">04 基础篇：经常说的 CPU 上下文切换是什么意思？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 基础篇：某个应用的CPU使用率居然达到100%，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/05%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e6%9f%90%e4%b8%aa%e5%ba%94%e7%94%a8%e7%9a%84CPU%e4%bd%bf%e7%94%a8%e7%8e%87%e5%b1%85%e7%84%b6%e8%be%be%e5%88%b0100%25%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">05 基础篇：某个应用的CPU使用率居然达到100%，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 案例篇：系统的 CPU 使用率很高，但为啥却找不到高 CPU 的应用？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/06%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e7%b3%bb%e7%bb%9f%e7%9a%84%20CPU%20%e4%bd%bf%e7%94%a8%e7%8e%87%e5%be%88%e9%ab%98%ef%bc%8c%e4%bd%86%e4%b8%ba%e5%95%a5%e5%8d%b4%e6%89%be%e4%b8%8d%e5%88%b0%e9%ab%98%20CPU%20%e7%9a%84%e5%ba%94%e7%94%a8%ef%bc%9f.md">06 案例篇：系统的 CPU 使用率很高，但为啥却找不到高 CPU 的应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 案例篇：系统中出现大量不可中断进程和僵尸进程怎么办？（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/07%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e7%b3%bb%e7%bb%9f%e4%b8%ad%e5%87%ba%e7%8e%b0%e5%a4%a7%e9%87%8f%e4%b8%8d%e5%8f%af%e4%b8%ad%e6%96%ad%e8%bf%9b%e7%a8%8b%e5%92%8c%e5%83%b5%e5%b0%b8%e8%bf%9b%e7%a8%8b%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">07 案例篇：系统中出现大量不可中断进程和僵尸进程怎么办？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 案例篇：系统中出现大量不可中断进程和僵尸进程怎么办？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/08%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e7%b3%bb%e7%bb%9f%e4%b8%ad%e5%87%ba%e7%8e%b0%e5%a4%a7%e9%87%8f%e4%b8%8d%e5%8f%af%e4%b8%ad%e6%96%ad%e8%bf%9b%e7%a8%8b%e5%92%8c%e5%83%b5%e5%b0%b8%e8%bf%9b%e7%a8%8b%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">08 案例篇：系统中出现大量不可中断进程和僵尸进程怎么办？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 基础篇：怎么理解Linux软中断？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/09%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e6%80%8e%e4%b9%88%e7%90%86%e8%a7%a3Linux%e8%bd%af%e4%b8%ad%e6%96%ad%ef%bc%9f.md">09 基础篇：怎么理解Linux软中断？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 案例篇：系统的软中断CPU使用率升高，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/10%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e7%b3%bb%e7%bb%9f%e7%9a%84%e8%bd%af%e4%b8%ad%e6%96%adCPU%e4%bd%bf%e7%94%a8%e7%8e%87%e5%8d%87%e9%ab%98%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">10 案例篇：系统的软中断CPU使用率升高，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 套路篇：如何迅速分析出系统CPU的瓶颈在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/11%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%85%e9%80%9f%e5%88%86%e6%9e%90%e5%87%ba%e7%b3%bb%e7%bb%9fCPU%e7%9a%84%e7%93%b6%e9%a2%88%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">11 套路篇：如何迅速分析出系统CPU的瓶颈在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 套路篇：CPU 性能优化的几个思路.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/12%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9aCPU%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84%e5%87%a0%e4%b8%aa%e6%80%9d%e8%b7%af.md">12 套路篇：CPU 性能优化的几个思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 答疑（一）：无法模拟出 RES 中断的问题，怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/13%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%97%a0%e6%b3%95%e6%a8%a1%e6%8b%9f%e5%87%ba%20RES%20%e4%b8%ad%e6%96%ad%e7%9a%84%e9%97%ae%e9%a2%98%ef%bc%8c%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">13 答疑（一）：无法模拟出 RES 中断的问题，怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 答疑（二）：如何用perf工具分析Java程序？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/14%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8perf%e5%b7%a5%e5%85%b7%e5%88%86%e6%9e%90Java%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">14 答疑（二）：如何用perf工具分析Java程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 基础篇：Linux内存是怎么工作的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/15%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9aLinux%e5%86%85%e5%ad%98%e6%98%af%e6%80%8e%e4%b9%88%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%9f.md">15 基础篇：Linux内存是怎么工作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 基础篇：怎么理解内存中的Buffer和Cache？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/16%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9a%e6%80%8e%e4%b9%88%e7%90%86%e8%a7%a3%e5%86%85%e5%ad%98%e4%b8%ad%e7%9a%84Buffer%e5%92%8cCache%ef%bc%9f.md">16 基础篇：怎么理解内存中的Buffer和Cache？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 案例篇：如何利用系统缓存优化程序的运行效率？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/17%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e7%b3%bb%e7%bb%9f%e7%bc%93%e5%ad%98%e4%bc%98%e5%8c%96%e7%a8%8b%e5%ba%8f%e7%9a%84%e8%bf%90%e8%a1%8c%e6%95%88%e7%8e%87%ef%bc%9f.md">17 案例篇：如何利用系统缓存优化程序的运行效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 案例篇：内存泄漏了，我该如何定位和处理？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/18%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%e4%ba%86%ef%bc%8c%e6%88%91%e8%af%a5%e5%a6%82%e4%bd%95%e5%ae%9a%e4%bd%8d%e5%92%8c%e5%a4%84%e7%90%86%ef%bc%9f.md">18 案例篇：内存泄漏了，我该如何定位和处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 案例篇：为什么系统的Swap变高了（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/19%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e7%b3%bb%e7%bb%9f%e7%9a%84Swap%e5%8f%98%e9%ab%98%e4%ba%86%ef%bc%88%e4%b8%8a%ef%bc%89.md">19 案例篇：为什么系统的Swap变高了（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 案例篇：为什么系统的Swap变高了？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/20%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e7%b3%bb%e7%bb%9f%e7%9a%84Swap%e5%8f%98%e9%ab%98%e4%ba%86%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">20 案例篇：为什么系统的Swap变高了？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 套路篇：如何“快准狠”找到系统内存的问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/21%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e2%80%9c%e5%bf%ab%e5%87%86%e7%8b%a0%e2%80%9d%e6%89%be%e5%88%b0%e7%b3%bb%e7%bb%9f%e5%86%85%e5%ad%98%e7%9a%84%e9%97%ae%e9%a2%98%ef%bc%9f.md">21 套路篇：如何“快准狠”找到系统内存的问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 答疑（三）：文件系统与磁盘的区别是什么？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/22%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e4%b8%8e%e7%a3%81%e7%9b%98%e7%9a%84%e5%8c%ba%e5%88%ab%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">22 答疑（三）：文件系统与磁盘的区别是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 基础篇：Linux 文件系统是怎么工作的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/23%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9aLinux%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e6%98%af%e6%80%8e%e4%b9%88%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%9f.md">23 基础篇：Linux 文件系统是怎么工作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 基础篇：Linux 磁盘I_O是怎么工作的（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/24%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9aLinux%20%e7%a3%81%e7%9b%98I_O%e6%98%af%e6%80%8e%e4%b9%88%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%88%e4%b8%8a%ef%bc%89.md">24 基础篇：Linux 磁盘I_O是怎么工作的（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 基础篇：Linux 磁盘I_O是怎么工作的（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/25%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9aLinux%20%e7%a3%81%e7%9b%98I_O%e6%98%af%e6%80%8e%e4%b9%88%e5%b7%a5%e4%bd%9c%e7%9a%84%ef%bc%88%e4%b8%8b%ef%bc%89.md">25 基础篇：Linux 磁盘I_O是怎么工作的（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 案例篇：如何找出狂打日志的“内鬼”？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/26%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%be%e5%87%ba%e7%8b%82%e6%89%93%e6%97%a5%e5%bf%97%e7%9a%84%e2%80%9c%e5%86%85%e9%ac%bc%e2%80%9d%ef%bc%9f.md">26 案例篇：如何找出狂打日志的“内鬼”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 案例篇：为什么我的磁盘I_O延迟很高？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/27%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e7%9a%84%e7%a3%81%e7%9b%98I_O%e5%bb%b6%e8%bf%9f%e5%be%88%e9%ab%98%ef%bc%9f.md">27 案例篇：为什么我的磁盘I_O延迟很高？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 案例篇：一个SQL查询要15秒，这是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/28%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e4%b8%80%e4%b8%aaSQL%e6%9f%a5%e8%af%a2%e8%a6%8115%e7%a7%92%ef%bc%8c%e8%bf%99%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">28 案例篇：一个SQL查询要15秒，这是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 案例篇：Redis响应严重延迟，如何解决？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/29%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9aRedis%e5%93%8d%e5%ba%94%e4%b8%a5%e9%87%8d%e5%bb%b6%e8%bf%9f%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%ef%bc%9f.md">29 案例篇：Redis响应严重延迟，如何解决？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 套路篇：如何迅速分析出系统I_O的瓶颈在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/30%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%85%e9%80%9f%e5%88%86%e6%9e%90%e5%87%ba%e7%b3%bb%e7%bb%9fI_O%e7%9a%84%e7%93%b6%e9%a2%88%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">30 套路篇：如何迅速分析出系统I_O的瓶颈在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 套路篇：磁盘 I_O 性能优化的几个思路.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/31%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e7%a3%81%e7%9b%98%20I_O%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84%e5%87%a0%e4%b8%aa%e6%80%9d%e8%b7%af.md">31 套路篇：磁盘 I_O 性能优化的几个思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 答疑（四）：阻塞、非阻塞 I_O 与同步、异步 I_O 的区别和联系.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/32%20%e7%ad%94%e7%96%91%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e9%98%bb%e5%a1%9e%e3%80%81%e9%9d%9e%e9%98%bb%e5%a1%9e%20I_O%20%e4%b8%8e%e5%90%8c%e6%ad%a5%e3%80%81%e5%bc%82%e6%ad%a5%20I_O%20%e7%9a%84%e5%8c%ba%e5%88%ab%e5%92%8c%e8%81%94%e7%b3%bb.md">32 答疑（四）：阻塞、非阻塞 I_O 与同步、异步 I_O 的区别和联系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 关于 Linux 网络，你必须知道这些（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/33%20%e5%85%b3%e4%ba%8e%20Linux%20%e7%bd%91%e7%bb%9c%ef%bc%8c%e4%bd%a0%e5%bf%85%e9%a1%bb%e7%9f%a5%e9%81%93%e8%bf%99%e4%ba%9b%ef%bc%88%e4%b8%8a%ef%bc%89.md">33 关于 Linux 网络，你必须知道这些（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 关于 Linux 网络，你必须知道这些（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/34%20%e5%85%b3%e4%ba%8e%20Linux%20%e7%bd%91%e7%bb%9c%ef%bc%8c%e4%bd%a0%e5%bf%85%e9%a1%bb%e7%9f%a5%e9%81%93%e8%bf%99%e4%ba%9b%ef%bc%88%e4%b8%8b%ef%bc%89.md">34 关于 Linux 网络，你必须知道这些（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 基础篇：C10K 和 C1000K 回顾.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/35%20%e5%9f%ba%e7%a1%80%e7%af%87%ef%bc%9aC10K%20%e5%92%8c%20C1000K%20%e5%9b%9e%e9%a1%be.md">35 基础篇：C10K 和 C1000K 回顾.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 套路篇：怎么评估系统的网络性能？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/36%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e6%80%8e%e4%b9%88%e8%af%84%e4%bc%b0%e7%b3%bb%e7%bb%9f%e7%9a%84%e7%bd%91%e7%bb%9c%e6%80%a7%e8%83%bd%ef%bc%9f.md">36 套路篇：怎么评估系统的网络性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 案例篇：DNS 解析时快时慢，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/37%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9aDNS%20%e8%a7%a3%e6%9e%90%e6%97%b6%e5%bf%ab%e6%97%b6%e6%85%a2%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">37 案例篇：DNS 解析时快时慢，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 案例篇：怎么使用 tcpdump 和 Wireshark 分析网络流量？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/38%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e6%80%8e%e4%b9%88%e4%bd%bf%e7%94%a8%20tcpdump%20%e5%92%8c%20Wireshark%20%e5%88%86%e6%9e%90%e7%bd%91%e7%bb%9c%e6%b5%81%e9%87%8f%ef%bc%9f.md">38 案例篇：怎么使用 tcpdump 和 Wireshark 分析网络流量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 案例篇：怎么缓解 DDoS 攻击带来的性能下降问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/39%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e6%80%8e%e4%b9%88%e7%bc%93%e8%a7%a3%20DDoS%20%e6%94%bb%e5%87%bb%e5%b8%a6%e6%9d%a5%e7%9a%84%e6%80%a7%e8%83%bd%e4%b8%8b%e9%99%8d%e9%97%ae%e9%a2%98%ef%bc%9f.md">39 案例篇：怎么缓解 DDoS 攻击带来的性能下降问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 案例篇：网络请求延迟变大了，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/40%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e7%bd%91%e7%bb%9c%e8%af%b7%e6%b1%82%e5%bb%b6%e8%bf%9f%e5%8f%98%e5%a4%a7%e4%ba%86%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">40 案例篇：网络请求延迟变大了，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 案例篇：如何优化 NAT 性能？（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/41%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%20NAT%20%e6%80%a7%e8%83%bd%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">41 案例篇：如何优化 NAT 性能？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 案例篇：如何优化 NAT 性能？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/42%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%20NAT%20%e6%80%a7%e8%83%bd%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">42 案例篇：如何优化 NAT 性能？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 套路篇：网络性能优化的几个思路（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/43%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e7%bd%91%e7%bb%9c%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84%e5%87%a0%e4%b8%aa%e6%80%9d%e8%b7%af%ef%bc%88%e4%b8%8a%ef%bc%89.md">43 套路篇：网络性能优化的几个思路（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 套路篇：网络性能优化的几个思路（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/44%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e7%bd%91%e7%bb%9c%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84%e5%87%a0%e4%b8%aa%e6%80%9d%e8%b7%af%ef%bc%88%e4%b8%8b%ef%bc%89.md">44 套路篇：网络性能优化的几个思路（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 答疑（五）：网络收发过程中，缓冲区位置在哪里？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/45%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%94%ef%bc%89%ef%bc%9a%e7%bd%91%e7%bb%9c%e6%94%b6%e5%8f%91%e8%bf%87%e7%a8%8b%e4%b8%ad%ef%bc%8c%e7%bc%93%e5%86%b2%e5%8c%ba%e4%bd%8d%e7%bd%ae%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">45 答疑（五）：网络收发过程中，缓冲区位置在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 案例篇：为什么应用容器化后，启动慢了很多？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/46%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%ba%94%e7%94%a8%e5%ae%b9%e5%99%a8%e5%8c%96%e5%90%8e%ef%bc%8c%e5%90%af%e5%8a%a8%e6%85%a2%e4%ba%86%e5%be%88%e5%a4%9a%ef%bc%9f.md">46 案例篇：为什么应用容器化后，启动慢了很多？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 案例篇：服务器总是时不时丢包，我该怎么办？（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/47%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%80%bb%e6%98%af%e6%97%b6%e4%b8%8d%e6%97%b6%e4%b8%a2%e5%8c%85%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">47 案例篇：服务器总是时不时丢包，我该怎么办？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 案例篇：服务器总是时不时丢包，我该怎么办？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/48%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%80%bb%e6%98%af%e6%97%b6%e4%b8%8d%e6%97%b6%e4%b8%a2%e5%8c%85%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">48 案例篇：服务器总是时不时丢包，我该怎么办？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 案例篇：内核线程 CPU 利用率太高，我该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/49%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%86%85%e6%a0%b8%e7%ba%bf%e7%a8%8b%20CPU%20%e5%88%a9%e7%94%a8%e7%8e%87%e5%a4%aa%e9%ab%98%ef%bc%8c%e6%88%91%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">49 案例篇：内核线程 CPU 利用率太高，我该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 案例篇：动态追踪怎么用？（上）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/50%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%8a%a8%e6%80%81%e8%bf%bd%e8%b8%aa%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">50 案例篇：动态追踪怎么用？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 案例篇：动态追踪怎么用？（下）.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/51%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e5%8a%a8%e6%80%81%e8%bf%bd%e8%b8%aa%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">51 案例篇：动态追踪怎么用？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52 案例篇：服务吞吐量下降很厉害，怎么分析？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/52%20%e6%a1%88%e4%be%8b%e7%af%87%ef%bc%9a%e6%9c%8d%e5%8a%a1%e5%90%9e%e5%90%90%e9%87%8f%e4%b8%8b%e9%99%8d%e5%be%88%e5%8e%89%e5%ae%b3%ef%bc%8c%e6%80%8e%e4%b9%88%e5%88%86%e6%9e%90%ef%bc%9f.md">52 案例篇：服务吞吐量下降很厉害，怎么分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="53 套路篇：系统监控的综合思路.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/53%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e7%b3%bb%e7%bb%9f%e7%9b%91%e6%8e%a7%e7%9a%84%e7%bb%bc%e5%90%88%e6%80%9d%e8%b7%af.md">53 套路篇：系统监控的综合思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="54 套路篇：应用监控的一般思路.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/54%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e5%ba%94%e7%94%a8%e7%9b%91%e6%8e%a7%e7%9a%84%e4%b8%80%e8%88%ac%e6%80%9d%e8%b7%af.md">54 套路篇：应用监控的一般思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="55 套路篇：分析性能问题的一般步骤.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/55%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e5%88%86%e6%9e%90%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%e7%9a%84%e4%b8%80%e8%88%ac%e6%ad%a5%e9%aa%a4.md">55 套路篇：分析性能问题的一般步骤.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="56 套路篇：优化性能问题的一般方法.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/56%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9a%e4%bc%98%e5%8c%96%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%e7%9a%84%e4%b8%80%e8%88%ac%e6%96%b9%e6%b3%95.md">56 套路篇：优化性能问题的一般方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="57 套路篇：Linux 性能工具速查.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/57%20%e5%a5%97%e8%b7%af%e7%af%87%ef%bc%9aLinux%20%e6%80%a7%e8%83%bd%e5%b7%a5%e5%85%b7%e9%80%9f%e6%9f%a5.md">57 套路篇：Linux 性能工具速查.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="58 答疑（六）：容器冷启动如何性能分析？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/58%20%e7%ad%94%e7%96%91%ef%bc%88%e5%85%ad%ef%bc%89%ef%bc%9a%e5%ae%b9%e5%99%a8%e5%86%b7%e5%90%af%e5%8a%a8%e5%a6%82%e4%bd%95%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%ef%bc%9f.md">58 答疑（六）：容器冷启动如何性能分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐（一） 书单推荐：性能优化和Linux 系统原理.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%ef%bc%88%e4%b8%80%ef%bc%89%20%e4%b9%a6%e5%8d%95%e6%8e%a8%e8%8d%90%ef%bc%9a%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%92%8cLinux%20%e7%b3%bb%e7%bb%9f%e5%8e%9f%e7%90%86.md">加餐（一） 书单推荐：性能优化和Linux 系统原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐（二） 书单推荐：网络原理和 Linux 内核实现.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%ef%bc%88%e4%ba%8c%ef%bc%89%20%e4%b9%a6%e5%8d%95%e6%8e%a8%e8%8d%90%ef%bc%9a%e7%bd%91%e7%bb%9c%e5%8e%9f%e7%90%86%e5%92%8c%20Linux%20%e5%86%85%e6%a0%b8%e5%ae%9e%e7%8e%b0.md">加餐（二） 书单推荐：网络原理和 Linux 内核实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 “半路出家 ”，也要顺利拿下性能优化！.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e2%80%9c%e5%8d%8a%e8%b7%af%e5%87%ba%e5%ae%b6%20%e2%80%9d%ef%bc%8c%e4%b9%9f%e8%a6%81%e9%a1%ba%e5%88%a9%e6%8b%bf%e4%b8%8b%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%81.md">用户故事 “半路出家 ”，也要顺利拿下性能优化！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 运维和开发工程师们怎么说？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e8%bf%90%e7%bb%b4%e5%92%8c%e5%bc%80%e5%8f%91%e5%b7%a5%e7%a8%8b%e5%b8%88%e4%bb%ac%e6%80%8e%e4%b9%88%e8%af%b4%ef%bc%9f.md">用户故事 运维和开发工程师们怎么说？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 愿你攻克性能难关.md" href="/%e4%b8%93%e6%a0%8f/Linux%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%84%bf%e4%bd%a0%e6%94%bb%e5%85%8b%e6%80%a7%e8%83%bd%e9%9a%be%e5%85%b3.md">结束语 愿你攻克性能难关.md</a>
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
                            <h1 id="title" data-id="09 基础篇：怎么理解Linux软中断？" class="title">09 基础篇：怎么理解Linux软中断？</h1>
                            <div><p>你好，我是倪朋飞。</p>

<p>上一期，我用一个不可中断进程的案例，带你学习了iowait（也就是等待I/O的CPU使用率）升高时的分析方法。这里你要记住，进程的不可中断状态是系统的一种保护机制，可以保证硬件的交互过程不被意外打断。所以，短时间的不可中断状态是很正常的。</p>

<p>但是，当进程长时间都处于不可中断状态时，你就得当心了。这时，你可以使用 dstat、pidstat 等工具，确认是不是磁盘 I/O 的问题，进而排查相关的进程和磁盘设备。关于磁盘 I/O 的性能问题，你暂且不用专门去背，我会在后续的 I/O 部分详细介绍，到时候理解了也就记住了。</p>

<p>其实除了 iowait，软中断（softirq）CPU使用率升高也是最常见的一种性能问题。接下来的两节课，我们就来学习软中断的内容，我还会以最常见的反向代理服务器 Nginx 的案例，带你分析这种情况。</p>

<h2 id="从-取外卖-看中断">从“取外卖”看中断</h2>

<p>说到中断，我在前面<a href="https://time.geekbang.org/column/article/69859" target="_blank">关于“上下文切换”的文章</a>，简单说过中断的含义，先来回顾一下。中断是系统用来响应硬件设备请求的一种机制，它会打断进程的正常调度和执行，然后调用内核中的中断处理程序来响应设备的请求。</p>

<p>你可能要问了，为什么要有中断呢？我可以举个生活中的例子，让你感受一下中断的魅力。</p>

<p>比如说你订了一份外卖，但是不确定外卖什么时候送到，也没有别的方法了解外卖的进度，但是，配送员送外卖是不等人的，到了你这儿没人取的话，就直接走人了。所以你只能苦苦等着，时不时去门口看看外卖送到没，而不能干其他事情。</p>

<p>不过呢，如果在订外卖的时候，你就跟配送员约定好，让他送到后给你打个电话，那你就不用苦苦等待了，就可以去忙别的事情，直到电话一响，接电话、取外卖就可以了。</p>

<p>这里的“打电话”，其实就是一个中断。没接到电话的时候，你可以做其他的事情；只有接到了电话（也就是发生中断），你才要进行另一个动作：取外卖。</p>

<p>这个例子你就可以发现，<strong>中断其实是一种异步的事件处理机制，可以提高系统的并发处理能力</strong>。</p>

<p>由于中断处理程序会打断其他进程的运行，所以，<strong>为了减少对正常进程运行调度的影响，中断处理程序就需要尽可能快地运行</strong>。如果中断本身要做的事情不多，那么处理起来也不会有太大问题；但如果中断要处理的事情很多，中断服务程序就有可能要运行很长时间。</p>

<p>特别是，中断处理程序在响应中断时，还会临时关闭中断。这就会导致上一次中断处理完成之前，其他中断都不能响应，也就是说中断有可能会丢失。</p>

<p>那么还是以取外卖为例。假如你订了 2 份外卖，一份主食和一份饮料，并且是由2个不同的配送员来配送。这次你不用时时等待着，两份外卖都约定了电话取外卖的方式。但是，问题又来了。</p>

<p>当第一份外卖送到时，配送员给你打了个长长的电话，商量发票的处理方式。与此同时，第二个配送员也到了，也想给你打电话。</p>

<p>但是很明显，因为电话占线（也就是关闭了中断响应），第二个配送员的电话是打不通的。所以，第二个配送员很可能试几次后就走掉了（也就是丢失了一次中断）。</p>

<h2 id="软中断">软中断</h2>

<p>如果你弄清楚了“取外卖”的模式，那对系统的中断机制就很容易理解了。事实上，为了解决中断处理程序执行过长和中断丢失的问题，Linux 将中断处理过程分成了两个阶段，也就是<strong>上半部和下半部</strong>：</p>

<ul>
<li><p><strong>上半部用来快速处理中断</strong>，它在中断禁止模式下运行，主要处理跟硬件紧密相关的或时间敏感的工作。</p></li>

<li><p><strong>下半部用来延迟处理上半部未完成的工作，通常以内核线程的方式运行</strong>。</p></li>
</ul>

<p>比如说前面取外卖的例子，上半部就是你接听电话，告诉配送员你已经知道了，其他事儿见面再说，然后电话就可以挂断了；下半部才是取外卖的动作，以及见面后商量发票处理的动作。</p>

<p>这样，第一个配送员不会占用你太多时间，当第二个配送员过来时，照样能正常打通你的电话。</p>

<p>除了取外卖，我再举个最常见的网卡接收数据包的例子，让你更好地理解。</p>

<p>网卡接收到数据包后，会通过<strong>硬件中断</strong>的方式，通知内核有新的数据到了。这时，内核就应该调用中断处理程序来响应它。你可以自己先想一下，这种情况下的上半部和下半部分别负责什么工作呢？</p>

<p>对上半部来说，既然是快速处理，其实就是要把网卡的数据读到内存中，然后更新一下硬件寄存器的状态（表示数据已经读好了），最后再发送一个<strong>软中断</strong>信号，通知下半部做进一步的处理。</p>

<p>而下半部被软中断信号唤醒后，需要从内存中找到网络数据，再按照网络协议栈，对数据进行逐层解析和处理，直到把它送给应用程序。</p>

<p>所以，这两个阶段你也可以这样理解：</p>

<ul>
<li><p>上半部直接处理硬件请求，也就是我们常说的硬中断，特点是快速执行；</p></li>

<li><p>而下半部则是由内核触发，也就是我们常说的软中断，特点是延迟执行。</p></li>
</ul>

<p>实际上，上半部会打断CPU正在执行的任务，然后立即执行中断处理程序。而下半部以内核线程的方式执行，并且每个 CPU 都对应一个软中断内核线程，名字为 “ksoftirqd/CPU编号”，比如说， 0 号CPU对应的软中断内核线程的名字就是 ksoftirqd/0。</p>

<p>不过要注意的是，软中断不只包括了刚刚所讲的硬件设备中断处理程序的下半部，一些内核自定义的事件也属于软中断，比如内核调度和RCU锁（Read-Copy Update 的缩写，RCU 是 Linux 内核中最常用的锁之一）等。</p>

<p>那要怎么知道你的系统里有哪些软中断呢？</p>

<h2 id="查看软中断和内核线程">查看软中断和内核线程</h2>

<p>不知道你还记不记得，前面提到过的 proc 文件系统。它是一种内核空间和用户空间进行通信的机制，可以用来查看内核的数据结构，或者用来动态修改内核的配置。其中：</p>

<ul>
<li><p>/proc/softirqs 提供了软中断的运行情况；</p></li>

<li><p>/proc/interrupts 提供了硬中断的运行情况。</p></li>
</ul>

<p>运行下面的命令，查看 /proc/softirqs 文件的内容，你就可以看到各种类型软中断在不同 CPU 上的累积运行次数：</p>

<pre><code>$ cat /proc/softirqs
                    CPU0       CPU1
          HI:          0          0
       TIMER:     811613    1972736
      NET_TX:         49          7
      NET_RX:    1136736    1506885
       BLOCK:          0          0
    IRQ_POLL:          0          0
     TASKLET:     304787       3691
       SCHED:     689718    1897539
     HRTIMER:          0          0
         RCU:    1330771    1354737
</code></pre>

<p>在查看 /proc/softirqs 文件内容时，你要特别注意以下这两点。</p>

<p>第一，要注意软中断的类型，也就是这个界面中第一列的内容。从第一列你可以看到，软中断包括了10个类别，分别对应不同的工作类型。比如 NET_RX 表示网络接收中断，而 NET_TX 表示网络发送中断。</p>

<p>第二，要注意同一种软中断在不同 CPU 上的分布情况，也就是同一行的内容。正常情况下，同一种中断在不同 CPU 上的累积次数应该差不多。比如这个界面中，NET_RX 在 CPU0 和 CPU1 上的中断次数基本是同一个数量级，相差不大。</p>

<p>不过你可能发现，TASKLET 在不同CPU上的分布并不均匀。TASKLET 是最常用的软中断实现机制，每个 TASKLET 只运行一次就会结束 ，并且只在调用它的函数所在的 CPU 上运行。</p>

<p>因此，使用 TASKLET 特别简便，当然也会存在一些问题，比如说由于只在一个CPU上运行导致的调度不均衡，再比如因为不能在多个 CPU 上并行运行带来了性能限制。</p>

<p>另外，刚刚提到过，软中断实际上是以内核线程的方式运行的，每个 CPU 都对应一个软中断内核线程，这个软中断内核线程就叫做 ksoftirqd/CPU编号。那要怎么查看这些线程的运行状况呢？</p>

<p>其实用 ps 命令就可以做到，比如执行下面的指令：</p>

<pre><code>$ ps aux | grep softirq
root         7  0.0  0.0      0     0 ?        S    Oct10   0:01 [ksoftirqd/0]
root        16  0.0  0.0      0     0 ?        S    Oct10   0:01 [ksoftirqd/1]
</code></pre>

<p>注意，这些线程的名字外面都有中括号，这说明 ps 无法获取它们的命令行参数（cmline）。一般来说，ps 的输出中，名字括在中括号里的，一般都是内核线程。</p>

<h2 id="小结">小结</h2>

<p>Linux 中的中断处理程序分为上半部和下半部：</p>

<ul>
<li><p>上半部对应硬件中断，用来快速处理中断。</p></li>

<li><p>下半部对应软中断，用来异步处理上半部未完成的工作。</p></li>
</ul>

<p>Linux 中的软中断包括网络收发、定时、调度、RCU锁等各种类型，可以通过查看 /proc/softirqs 来观察软中断的运行情况。</p>

<h2 id="思考">思考</h2>

<p>最后，我想请你一起聊聊，你是怎么理解软中断的？你有没有碰到过因为软中断出现的性能问题？你又是怎么分析它们的瓶颈的呢？你可以结合今天的内容，总结自己的思路，写下自己的问题。</p>

<p>欢迎在留言区和我讨论，也欢迎把这篇文章分享给你的同事、朋友。我们一起在实战中演练，在交流中进步。</p>

<p><img src="assets/7f5bf5b283db49d7b83568c70e8a44ed.jpg" alt="" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#eb878787d2dfdadadbdcab8c868a8287c5888486" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f10e7e1addc6365',t:'MTczNDAzOTk3Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
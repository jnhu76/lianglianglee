<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=23&#32;基础篇：Linux&#32;文件系统是怎么工作的？>
        <link rel="icon" href="/static/favicon.png">
        <title>23 基础篇：Linux 文件系统是怎么工作的？ </title>
        
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
                            <h1 id="title" data-id="23 基础篇：Linux 文件系统是怎么工作的？" class="title">23 基础篇：Linux 文件系统是怎么工作的？</h1>
                            <div><p>你好，我是倪朋飞。</p>

<p>通过前面CPU和内存模块的学习，我相信，你已经掌握了CPU和内存的性能分析以及优化思路。从这一节开始，我们将进入下一个重要模块——文件系统和磁盘的I/O性能。</p>

<p>同CPU、内存一样，磁盘和文件系统的管理，也是操作系统最核心的功能。</p>

<ul>
<li><p>磁盘为系统提供了最基本的持久化存储。</p></li>

<li><p>文件系统则在磁盘的基础上，提供了一个用来管理文件的树状结构。</p></li>
</ul>

<p>那么，磁盘和文件系统是怎么工作的呢？又有哪些指标可以衡量它们的性能呢？</p>

<p>今天，我就带你先来看看，Linux文件系统的工作原理。磁盘的工作原理，我们下一节再来学习。</p>

<h2 id="索引节点和目录项">索引节点和目录项</h2>

<p>文件系统，本身是对存储设备上的文件，进行组织管理的机制。组织方式不同，就会形成不同的文件系统。</p>

<p>你要记住最重要的一点，在Linux中一切皆文件。不仅普通的文件和目录，就连块设备、套接字、管道等，也都要通过统一的文件系统来管理。</p>

<p>为了方便管理，Linux文件系统为每个文件都分配两个数据结构，索引节点（index node）和目录项（directory entry）。它们主要用来记录文件的元信息和目录结构。</p>

<ul>
<li><p>索引节点，简称为inode，用来记录文件的元数据，比如inode编号、文件大小、访问权限、修改日期、数据的位置等。索引节点和文件一一对应，它跟文件内容一样，都会被持久化存储到磁盘中。所以记住，索引节点同样占用磁盘空间。</p></li>

<li><p>目录项，简称为dentry，用来记录文件的名字、索引节点指针以及与其他目录项的关联关系。多个关联的目录项，就构成了文件系统的目录结构。不过，不同于索引节点，目录项是由内核维护的一个内存数据结构，所以通常也被叫做目录项缓存。</p></li>
</ul>

<p>换句话说，索引节点是每个文件的唯一标志，而目录项维护的正是文件系统的树状结构。目录项和索引节点的关系是多对一，你可以简单理解为，一个文件可以有多个别名。</p>

<p>举个例子，通过硬链接为文件创建的别名，就会对应不同的目录项，不过这些目录项本质上还是链接同一个文件，所以，它们的索引节点相同。</p>

<p>索引节点和目录项纪录了文件的元数据，以及文件间的目录关系，那么具体来说，文件数据到底是怎么存储的呢？是不是直接写到磁盘中就好了呢？</p>

<p>实际上，磁盘读写的最小单位是扇区，然而扇区只有512B 大小，如果每次都读写这么小的单位，效率一定很低。所以，文件系统又把连续的扇区组成了逻辑块，然后每次都以逻辑块为最小单元，来管理数据。常见的逻辑块大小为4KB，也就是由连续的8个扇区组成。</p>

<p>为了帮助你理解目录项、索引节点以及文件数据的关系，我画了一张示意图。你可以对照着这张图，来回忆刚刚讲过的内容，把知识和细节串联起来。</p>

<p><img src="assets/2e883fa9578e4fd495e0279ee810857c.jpg" alt="" /></p>

<p>不过，这里有两点需要你注意。</p>

<p>第一，目录项本身就是一个内存缓存，而索引节点则是存储在磁盘中的数据。在前面的Buffer和Cache原理中，我曾经提到过，为了协调慢速磁盘与快速CPU的性能差异，文件内容会缓存到页缓存Cache中。</p>

<p>那么，你应该想到，这些索引节点自然也会缓存到内存中，加速文件的访问。</p>

<p>第二，磁盘在执行文件系统格式化时，会被分成三个存储区域，超级块、索引节点区和数据块区。其中，</p>

<ul>
<li><p>超级块，存储整个文件系统的状态。</p></li>

<li><p>索引节点区，用来存储索引节点。</p></li>

<li><p>数据块区，则用来存储文件数据。</p></li>
</ul>

<h2 id="虚拟文件系统">虚拟文件系统</h2>

<p>目录项、索引节点、逻辑块以及超级块，构成了Linux文件系统的四大基本要素。不过，为了支持各种不同的文件系统，Linux内核在用户进程和文件系统的中间，又引入了一个抽象层，也就是虚拟文件系统VFS（Virtual File System）。</p>

<p>VFS 定义了一组所有文件系统都支持的数据结构和标准接口。这样，用户进程和内核中的其他子系统，只需要跟VFS 提供的统一接口进行交互就可以了，而不需要再关心底层各种文件系统的实现细节。</p>

<p>这里，我画了一张Linux文件系统的架构图，帮你更好地理解系统调用、VFS、缓存、文件系统以及块存储之间的关系。</p>

<p><img src="assets/84541910def0493d9cca932684b5cae2.jpg" alt="" /></p>

<p>通过这张图，你可以看到，在VFS的下方，Linux支持各种各样的文件系统，如Ext4、XFS、NFS等等。按照存储位置的不同，这些文件系统可以分为三类。</p>

<ul>
<li><p>第一类是基于磁盘的文件系统，也就是把数据直接存储在计算机本地挂载的磁盘中。常见的Ext4、XFS、OverlayFS等，都是这类文件系统。</p></li>

<li><p>第二类是基于内存的文件系统，也就是我们常说的虚拟文件系统。这类文件系统，不需要任何磁盘分配存储空间，但会占用内存。我们经常用到的 /proc 文件系统，其实就是一种最常见的虚拟文件系统。此外，/sys 文件系统也属于这一类，主要向用户空间导出层次化的内核对象。</p></li>

<li><p>第三类是网络文件系统，也就是用来访问其他计算机数据的文件系统，比如NFS、SMB、iSCSI等。</p></li>
</ul>

<p>这些文件系统，要先挂载到 VFS 目录树中的某个子目录（称为挂载点），然后才能访问其中的文件。拿第一类，也就是基于磁盘的文件系统为例，在安装系统时，要先挂载一个根目录（/），在根目录下再把其他文件系统（比如其他的磁盘分区、/proc文件系统、/sys文件系统、NFS等）挂载进来。</p>

<h2 id="文件系统i-o">文件系统I/O</h2>

<p>把文件系统挂载到挂载点后，你就能通过挂载点，再去访问它管理的文件了。VFS 提供了一组标准的文件访问接口。这些接口以系统调用的方式，提供给应用程序使用。</p>

<p>就拿cat 命令来说，它首先调用 open() ，打开一个文件；然后调用 read() ，读取文件的内容；最后再调用 write() ，把文件内容输出到控制台的标准输出中：</p>

<pre><code>int open(const char *pathname, int flags, mode_t mode); 
ssize_t read(int fd, void *buf, size_t count); 
ssize_t write(int fd, const void *buf, size_t count); 
</code></pre>

<p>文件读写方式的各种差异，导致 I/O的分类多种多样。最常见的有，缓冲与非缓冲I/O、直接与非直接I/O、阻塞与非阻塞I/O、同步与异步I/O等。 接下来，我们就详细看这四种分类。</p>

<p>第一种，根据是否利用标准库缓存，可以把文件I/O分为缓冲I/O与非缓冲I/O。</p>

<ul>
<li><p>缓冲I/O，是指利用标准库缓存来加速文件的访问，而标准库内部再通过系统调度访问文件。</p></li>

<li><p>非缓冲I/O，是指直接通过系统调用来访问文件，不再经过标准库缓存。</p></li>
</ul>

<p>注意，这里所说的“缓冲”，是指标准库内部实现的缓存。比方说，你可能见到过，很多程序遇到换行时才真正输出，而换行前的内容，其实就是被标准库暂时缓存了起来。</p>

<p>无论缓冲I/O还是非缓冲I/O，它们最终还是要经过系统调用来访问文件。而根据上一节内容，我们知道，系统调用后，还会通过页缓存，来减少磁盘的I/O操作。</p>

<p>第二，根据是否利用操作系统的页缓存，可以把文件I/O分为直接I/O与非直接I/O。</p>

<ul>
<li><p>直接I/O，是指跳过操作系统的页缓存，直接跟文件系统交互来访问文件。</p></li>

<li><p>非直接I/O正好相反，文件读写时，先要经过系统的页缓存，然后再由内核或额外的系统调用，真正写入磁盘。</p></li>
</ul>

<p>想要实现直接I/O，需要你在系统调用中，指定 O_DIRECT 标志。如果没有设置过，默认的是非直接I/O。</p>

<p>不过要注意，直接I/O、非直接I/O，本质上还是和文件系统交互。如果是在数据库等场景中，你还会看到，跳过文件系统读写磁盘的情况，也就是我们通常所说的裸I/O。</p>

<p>第三，根据应用程序是否阻塞自身运行，可以把文件I/O分为阻塞I/O和非阻塞I/O：</p>

<ul>
<li><p>所谓阻塞I/O，是指应用程序执行I/O操作后，如果没有获得响应，就会阻塞当前线程，自然就不能执行其他任务。</p></li>

<li><p>所谓非阻塞I/O，是指应用程序执行I/O操作后，不会阻塞当前的线程，可以继续执行其他的任务，随后再通过轮询或者事件通知的形式，获取调用的结果。</p></li>
</ul>

<p>比方说，访问管道或者网络套接字时，设置 O_NONBLOCK 标志，就表示用非阻塞方式访问；而如果不做任何设置，默认的就是阻塞访问。</p>

<p>第四，根据是否等待响应结果，可以把文件I/O分为同步和异步I/O：</p>

<ul>
<li><p>所谓同步I/O，是指应用程序执行I/O操作后，要一直等到整个I/O完成后，才能获得I/O响应。</p></li>

<li><p>所谓异步I/O，是指应用程序执行I/O操作后，不用等待完成和完成后的响应，而是继续执行就可以。等到这次 I/O完成后，响应会用事件通知的方式，告诉应用程序。</p></li>
</ul>

<p>举个例子，在操作文件时，如果你设置了 O_SYNC 或者 O_DSYNC 标志，就代表同步I/O。如果设置了O_DSYNC，就要等文件数据写入磁盘后，才能返回；而O_SYNC，则是在O_DSYNC基础上，要求文件元数据也要写入磁盘后，才能返回。</p>

<p>再比如，在访问管道或者网络套接字时，设置了O_ASYNC选项后，相应的I/O就是异步I/O。这样，内核会再通过SIGIO或者SIGPOLL，来通知进程文件是否可读写。</p>

<p>你可能发现了，这里的好多概念也经常出现在网络编程中。比如非阻塞I/O，通常会跟select/poll配合，用在网络套接字的I/O中。</p>

<p>你也应该可以理解，“Linux 一切皆文件”的深刻含义。无论是普通文件和块设备、还是网络套接字和管道等，它们都通过统一的VFS 接口来访问。</p>

<h2 id="性能观测">性能观测</h2>

<p>学了这么多文件系统的原理，你估计也是迫不及待想上手，观察一下文件系统的性能情况了。</p>

<p>接下来，打开一个终端，SSH登录到服务器上，然后跟我一起来探索，如何观测文件系统的性能。</p>

<h3 id="容量">容量</h3>

<p>对文件系统来说，最常见的一个问题就是空间不足。当然，你可能本身就知道，用 df 命令，就能查看文件系统的磁盘空间使用情况。比如：</p>

<pre><code>$ df /dev/sda1 
Filesystem     1K-blocks    Used Available Use% Mounted on 
/dev/sda1       30308240 3167020  27124836  11% / 
</code></pre>

<p>你可以看到，我的根文件系统只使用了11%的空间。这里还要注意，总空间用1K-blocks的数量来表示，你可以给df加上-h选项，以获得更好的可读性：</p>

<pre><code>$ df -h /dev/sda1 
Filesystem      Size  Used Avail Use% Mounted on 
/dev/sda1        29G  3.1G   26G  11% / 
</code></pre>

<p>不过有时候，明明你碰到了空间不足的问题，可是用df查看磁盘空间后，却发现剩余空间还有很多。这是怎么回事呢？</p>

<p>不知道你还记不记得，刚才我强调的一个细节。除了文件数据，索引节点也占用磁盘空间。你可以给df命令加上 -i 参数，查看索引节点的使用情况，如下所示：</p>

<pre><code>$ df -i /dev/sda1 
Filesystem      Inodes  IUsed   IFree IUse% Mounted on 
/dev/sda1      3870720 157460 3713260    5% / 
</code></pre>

<p>索引节点的容量，（也就是Inode个数）是在格式化磁盘时设定好的，一般由格式化工具自动生成。当你发现索引节点空间不足，但磁盘空间充足时，很可能就是过多小文件导致的。</p>

<p>所以，一般来说，删除这些小文件，或者把它们移动到索引节点充足的其他磁盘中，就可以解决这个问题。</p>

<h3 id="缓存">缓存</h3>

<p>在前面Cache案例中，我已经介绍过，可以用 free 或 vmstat，来观察页缓存的大小。复习一下，free输出的Cache，是页缓存和可回收Slab缓存的和，你可以从 /proc/meminfo ，直接得到它们的大小：</p>

<pre><code>$ cat /proc/meminfo | grep -E &quot;SReclaimable|Cached&quot; 
Cached:           748316 kB 
SwapCached:            0 kB 
SReclaimable:     179508 kB 
</code></pre>

<p>话说回来，文件系统中的目录项和索引节点缓存，又该如何观察呢？</p>

<p>实际上，内核使用Slab机制，管理目录项和索引节点的缓存。/proc/meminfo只给出了Slab的整体大小，具体到每一种Slab缓存，还要查看/proc/slabinfo这个文件。</p>

<p>比如，运行下面的命令，你就可以得到，所有目录项和各种文件系统索引节点的缓存情况：</p>

<pre><code>$ cat /proc/slabinfo | grep -E '^#|dentry|inode' 
# name            &lt;active_objs&gt; &lt;num_objs&gt; &lt;objsize&gt; &lt;objperslab&gt; &lt;pagesperslab&gt; : tunables &lt;limit&gt; &lt;batchcount&gt; &lt;sharedfactor&gt; : slabdata &lt;active_slabs&gt; &lt;num_slabs&gt; &lt;sharedavail&gt; 
xfs_inode              0      0    960   17    4 : tunables    0    0    0 : slabdata      0      0      0 
... 
ext4_inode_cache   32104  34590   1088   15    4 : tunables    0    0    0 : slabdata   2306   2306      0hugetlbfs_inode_cache     13     13    624   13    2 : tunables    0    0    0 : slabdata      1      1      0 
sock_inode_cache    1190   1242    704   23    4 : tunables    0    0    0 : slabdata     54     54      0 
shmem_inode_cache   1622   2139    712   23    4 : tunables    0    0    0 : slabdata     93     93      0 
proc_inode_cache    3560   4080    680   12    2 : tunables    0    0    0 : slabdata    340    340      0 
inode_cache        25172  25818    608   13    2 : tunables    0    0    0 : slabdata   1986   1986      0 
dentry             76050 121296    192   21    1 : tunables    0    0    0 : slabdata   5776   5776      0 
</code></pre>

<p>这个界面中，dentry行表示目录项缓存，inode_cache行，表示VFS索引节点缓存，其余的则是各种文件系统的索引节点缓存。</p>

<p>/proc/slabinfo 的列比较多，具体含义你可以查询 man slabinfo。在实际性能分析中，我们更常使用 slabtop ，来找到占用内存最多的缓存类型。</p>

<p>比如，下面就是我运行slabtop得到的结果：</p>

<pre><code># 按下c按照缓存大小排序，按下a按照活跃对象数排序 
$ slabtop 
Active / Total Objects (% used)    : 277970 / 358914 (77.4%) 
Active / Total Slabs (% used)      : 12414 / 12414 (100.0%) 
Active / Total Caches (% used)     : 83 / 135 (61.5%) 
Active / Total Size (% used)       : 57816.88K / 73307.70K (78.9%) 
Minimum / Average / Maximum Object : 0.01K / 0.20K / 22.88K 

  OBJS ACTIVE  USE OBJ SIZE  SLABS OBJ/SLAB CACHE SIZE NAME 
69804  23094   0%    0.19K   3324       21     13296K dentry 
16380  15854   0%    0.59K   1260       13     10080K inode_cache 
58260  55397   0%    0.13K   1942       30      7768K kernfs_node_cache 
   485    413   0%    5.69K     97        5      3104K task_struct 
  1472   1397   0%    2.00K     92       16      2944K kmalloc-2048 
</code></pre>

<p>从这个结果你可以看到，在我的系统中，目录项和索引节点占用了最多的Slab缓存。不过它们占用的内存其实并不大，加起来也只有23MB左右。</p>

<h2 id="小结">小结</h2>

<p>今天，我带你梳理了Linux文件系统的工作原理。</p>

<p>文件系统，是对存储设备上的文件，进行组织管理的一种机制。为了支持各类不同的文件系统，Linux在各种文件系统实现上，抽象了一层虚拟文件系统（VFS）。</p>

<p>VFS 定义了一组所有文件系统都支持的数据结构和标准接口。这样，用户进程和内核中的其他子系统，就只需要跟 VFS 提供的统一接口进行交互。</p>

<p>为了降低慢速磁盘对性能的影响，文件系统又通过页缓存、目录项缓存以及索引节点缓存，缓和磁盘延迟对应用程序的影响。</p>

<p>在性能观测方面，今天主要讲了容量和缓存的指标。下一节，我们将会学习Linux磁盘 I/O的工作原理，并掌握磁盘I/O的性能观测方法。</p>

<h2 id="思考">思考</h2>

<p>最后，给你留一个思考题。在实际工作中，我们经常会根据文件名字，查找它所在路径，比如：</p>

<pre><code>$ find / -name file-name
</code></pre>

<p>今天的问题就是，这个命令，会不会导致系统的缓存升高呢？如果有影响，又会导致哪种类型的缓存升高呢？你可以结合今天内容，自己先去操作和分析，看看观察到的结果跟你分析的是否一样。</p>

<p>欢迎在留言区和我讨论，也欢迎把这篇文章分享给你的同事、朋友。我们一起在实战中演练，在交流中进步。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#18747474212c2929282f587f75797174367b7775" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f10ec68acac6365',t:'MTczNDA0MDE1Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=10&#32;案例篇：系统的软中断CPU使用率升高，我该怎么办？>
        <link rel="icon" href="/static/favicon.png">
        <title>10 案例篇：系统的软中断CPU使用率升高，我该怎么办？ </title>
        
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
                            <h1 id="title" data-id="10 案例篇：系统的软中断CPU使用率升高，我该怎么办？" class="title">10 案例篇：系统的软中断CPU使用率升高，我该怎么办？</h1>
                            <div><p>你好，我是倪朋飞。</p>

<p>上一期我给你讲了软中断的基本原理，我们先来简单复习下。</p>

<p>中断是一种异步的事件处理机制，用来提高系统的并发处理能力。中断事件发生，会触发执行中断处理程序，而中断处理程序被分为上半部和下半部这两个部分。</p>

<ul>
<li><p>上半部对应硬中断，用来快速处理中断；</p></li>

<li><p>下半部对应软中断，用来异步处理上半部未完成的工作。</p></li>
</ul>

<p>Linux 中的软中断包括网络收发、定时、调度、RCU锁等各种类型，我们可以查看 proc 文件系统中的 /proc/softirqs ，观察软中断的运行情况。</p>

<p>在 Linux 中，每个 CPU 都对应一个软中断内核线程，名字是 ksoftirqd/CPU编号。当软中断事件的频率过高时，内核线程也会因为CPU 使用率过高而导致软中断处理不及时，进而引发网络收发延迟、调度缓慢等性能问题。</p>

<p>软中断 CPU 使用率过高也是一种最常见的性能问题。今天，我就用最常见的反向代理服务器 Nginx 的案例，教你学会分析这种情况。</p>

<h2 id="案例">案例</h2>

<h3 id="你的准备">你的准备</h3>

<p>接下来的案例基于 Ubuntu 18.04，也同样适用于其他的 Linux 系统。我使用的案例环境是这样的：</p>

<ul>
<li><p>机器配置：2 CPU、8 GB 内存。</p></li>

<li><p>预先安装 docker、sysstat、sar 、hping3、tcpdump 等工具，比如 apt-get install <a href="http://docker.io" target="_blank">docker.io</a> sysstat hping3 tcpdump。</p></li>
</ul>

<p>这里我又用到了三个新工具，sar、 hping3 和 tcpdump，先简单介绍一下：</p>

<ul>
<li><p>sar 是一个系统活动报告工具，既可以实时查看系统的当前活动，又可以配置保存和报告历史统计数据。</p></li>

<li><p>hping3 是一个可以构造 TCP/IP 协议数据包的工具，可以对系统进行安全审计、防火墙测试等。</p></li>

<li><p>tcpdump 是一个常用的网络抓包工具，常用来分析各种网络问题。</p></li>
</ul>

<p>本次案例用到两台虚拟机，我画了一张图来表示它们的关系。</p>

<p><img src="assets/e832f675711442ef8bddd3139eddd18f.jpg" alt="" /></p>

<p>你可以看到，其中一台虚拟机运行 Nginx ，用来模拟待分析的 Web 服务器；而另一台当作Web 服务器的客户端，用来给 Nginx 增加压力请求。使用两台虚拟机的目的，是为了相互隔离，避免“交叉感染”。</p>

<p>接下来，我们打开两个终端，分别 SSH 登录到两台机器上，并安装上面提到的这些工具。</p>

<p>同以前的案例一样，下面的所有命令都默认以 root 用户运行，如果你是用普通用户身份登陆系统，请运行 sudo su root 命令切换到 root 用户。</p>

<p>如果安装过程中有什么问题，同样鼓励你先自己搜索解决，解决不了的，可以在留言区向我提问。如果你以前已经安装过了，就可以忽略这一点了。</p>

<h3 id="操作和分析">操作和分析</h3>

<p>安装完成后，我们先在第一个终端，执行下面的命令运行案例，也就是一个最基本的 Nginx 应用：</p>

<pre><code># 运行Nginx服务并对外开放80端口
$ docker run -itd --name=nginx -p 80:80 nginx
</code></pre>

<p>然后，在第二个终端，使用 curl 访问 Nginx 监听的端口，确认 Nginx 正常启动。假设 192.168.0.30 是 Nginx 所在虚拟机的 IP 地址，运行 curl 命令后你应该会看到下面这个输出界面：</p>

<pre><code>$ curl http://192.168.0.30/
&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;head&gt;
&lt;title&gt;Welcome to nginx!&lt;/title&gt;
...
</code></pre>

<p>接着，还是在第二个终端，我们运行 hping3 命令，来模拟 Nginx 的客户端请求：</p>

<pre><code># -S参数表示设置TCP协议的SYN（同步序列号），-p表示目的端口为80
# -i u100表示每隔100微秒发送一个网络帧
# 注：如果你在实践过程中现象不明显，可以尝试把100调小，比如调成10甚至1
$ hping3 -S -p 80 -i u100 192.168.0.30
</code></pre>

<p>现在我们再回到第一个终端，你应该发现了异常。是不是感觉系统响应明显变慢了，即便只是在终端中敲几个回车，都得很久才能得到响应？这个时候应该怎么办呢？</p>

<p>虽然在运行 hping3 命令时，我就已经告诉你，这是一个 SYN FLOOD 攻击，你肯定也会想到从网络方面入手，来分析这个问题。不过，在实际的生产环境中，没人直接告诉你原因。</p>

<p>所以，我希望你把 hping3 模拟 SYN FLOOD 这个操作暂时忘掉，然后重新从观察到的问题开始，分析系统的资源使用情况，逐步找出问题的根源。</p>

<p>那么，该从什么地方入手呢？刚才我们发现，简单的 SHELL 命令都明显变慢了，先看看系统的整体资源使用情况应该是个不错的注意，比如执行下 top 看看是不是出现了 CPU 的瓶颈。我们在第一个终端运行 top 命令，看一下系统整体的资源使用情况。</p>

<pre><code># top运行后按数字1切换到显示所有CPU
$ top
top - 10:50:58 up 1 days, 22:10,  1 user,  load average: 0.00, 0.00, 0.00
Tasks: 122 total,   1 running,  71 sleeping,   0 stopped,   0 zombie
%Cpu0  :  0.0 us,  0.0 sy,  0.0 ni, 96.7 id,  0.0 wa,  0.0 hi,  3.3 si,  0.0 st
%Cpu1  :  0.0 us,  0.0 sy,  0.0 ni, 95.6 id,  0.0 wa,  0.0 hi,  4.4 si,  0.0 st
...

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
    7 root      20   0       0      0      0 S   0.3  0.0   0:01.64 ksoftirqd/0
   16 root      20   0       0      0      0 S   0.3  0.0   0:01.97 ksoftirqd/1
 2663 root      20   0  923480  28292  13996 S   0.3  0.3   4:58.66 docker-containe
 3699 root      20   0       0      0      0 I   0.3  0.0   0:00.13 kworker/u4:0
 3708 root      20   0   44572   4176   3512 R   0.3  0.1   0:00.07 top
    1 root      20   0  225384   9136   6724 S   0.0  0.1   0:23.25 systemd
    2 root      20   0       0      0      0 S   0.0  0.0   0:00.03 kthreadd
...
</code></pre>

<p>这里你有没有发现异常的现象？我们从第一行开始，逐个看一下：</p>

<ul>
<li><p>平均负载全是0，就绪队列里面只有一个进程（1 running）。</p></li>

<li><p>每个CPU的使用率都挺低，最高的CPU1的使用率也只有4.4%，并不算高。</p></li>

<li><p>再看进程列表，CPU使用率最高的进程也只有 0.3%，还是不高呀。</p></li>
</ul>

<p>那为什么系统的响应变慢了呢？既然每个指标的数值都不大，那我们就再来看看，这些指标对应的更具体的含义。毕竟，哪怕是同一个指标，用在系统的不同部位和场景上，都有可能对应着不同的性能问题。</p>

<p>仔细看 top 的输出，两个 CPU的使用率虽然分别只有 3.3%和4.4%，但都用在了软中断上；而从进程列表上也可以看到，CPU使用率最高的也是软中断进程 ksoftirqd。看起来，软中断有点可疑了。</p>

<p>根据上一期的内容，既然软中断可能有问题，那你先要知道，究竟是哪类软中断的问题。停下来想想，上一节我们用了什么方法，来判断软中断类型呢？没错，还是proc文件系统。观察 /proc/softirqs 文件的内容，你就能知道各种软中断类型的次数。</p>

<p>不过，这里的各类软中断次数，又是什么时间段里的次数呢？它是系统运行以来的<strong>累积中断次数</strong>。所以我们直接查看文件内容，得到的只是累积中断次数，对这里的问题并没有直接参考意义。因为，这些<strong>中断次数的变化速率</strong>才是我们需要关注的。</p>

<p>那什么工具可以观察命令输出的变化情况呢？我想你应该想起来了，在前面案例中用过的 watch 命令，就可以定期运行一个命令来查看输出；如果再加上 -d 参数，还可以高亮出变化的部分，从高亮部分我们就可以直观看出，哪些内容变化得更快。</p>

<p>比如，还是在第一个终端，我们运行下面的命令：</p>

<pre><code>$ watch -d cat /proc/softirqs
                    CPU0       CPU1
          HI:          0          0
       TIMER:    1083906    2368646
      NET_TX:         53          9
      NET_RX:    1550643    1916776
       BLOCK:          0          0
    IRQ_POLL:          0          0
     TASKLET:     333637       3930
       SCHED:     963675    2293171
     HRTIMER:          0          0
         RCU:    1542111    1590625
</code></pre>

<p>通过 /proc/softirqs 文件内容的变化情况，你可以发现， TIMER（定时中断）、NET_RX（网络接收）、SCHED（内核调度）、RCU（RCU锁）等这几个软中断都在不停变化。</p>

<p>其中，NET_RX，也就是网络数据包接收软中断的变化速率最快。而其他几种类型的软中断，是保证 Linux 调度、时钟和临界区保护这些正常工作所必需的，所以它们有一定的变化倒是正常的。</p>

<p>那么接下来，我们就从网络接收的软中断着手，继续分析。既然是网络接收的软中断，第一步应该就是观察系统的网络接收情况。这里你可能想起了很多网络工具，不过，我推荐今天的主人公工具 sar 。</p>

<p>sar 可以用来查看系统的网络收发情况，还有一个好处是，不仅可以观察网络收发的吞吐量（BPS，每秒收发的字节数），还可以观察网络收发的 PPS，即每秒收发的网络帧数。</p>

<p>我们在第一个终端中运行 sar 命令，并添加 -n DEV 参数显示网络收发的报告：</p>

<pre><code># -n DEV 表示显示网络收发的报告，间隔1秒输出一组数据
$ sar -n DEV 1
15:03:46        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
15:03:47         eth0  12607.00   6304.00    664.86    358.11      0.00      0.00      0.00      0.01
15:03:47      docker0   6302.00  12604.00    270.79    664.66      0.00      0.00      0.00      0.00
15:03:47           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
15:03:47    veth9f6bbcd   6302.00  12604.00    356.95    664.66      0.00      0.00      0.00      0.05
</code></pre>

<p>对于 sar 的输出界面，我先来简单介绍一下，从左往右依次是：</p>

<ul>
<li><p>第一列：表示报告的时间。</p></li>

<li><p>第二列：IFACE 表示网卡。</p></li>

<li><p>第三、四列：rxpck/s 和 txpck/s 分别表示每秒接收、发送的网络帧数，也就是 PPS。</p></li>

<li><p>第五、六列：rxkB/s 和 txkB/s 分别表示每秒接收、发送的千字节数，也就是 BPS。</p></li>

<li><p>后面的其他参数基本接近0，显然跟今天的问题没有直接关系，你可以先忽略掉。</p></li>
</ul>

<p>我们具体来看输出的内容，你可以发现：</p>

<ul>
<li><p>对网卡 eth0来说，每秒接收的网络帧数比较大，达到了 12607，而发送的网络帧数则比较小，只有 6304；每秒接收的千字节数只有 664 KB，而发送的千字节数更小，只有 358 KB。</p></li>

<li><p>docker0 和 veth9f6bbcd 的数据跟 eth0 基本一致，只是发送和接收相反，发送的数据较大而接收的数据较小。这是 Linux 内部网桥转发导致的，你暂且不用深究，只要知道这是系统把 eth0 收到的包转发给 Nginx 服务即可。具体工作原理，我会在后面的网络部分详细介绍。</p></li>
</ul>

<p>从这些数据，你有没有发现什么异常的地方？</p>

<p>既然怀疑是网络接收中断的问题，我们还是重点来看 eth0 ：接收的 PPS 比较大，达到 12607，而接收的 BPS 却很小，只有 664 KB。直观来看网络帧应该都是比较小的，我们稍微计算一下，664*<sup>1024</sup>&frasl;<sub>12607</sub> = 54 字节，说明平均每个网络帧只有 54 字节，这显然是很小的网络帧，也就是我们通常所说的小包问题。</p>

<p>那么，有没有办法知道这是一个什么样的网络帧，以及从哪里发过来的呢？</p>

<p>使用 tcpdump 抓取 eth0 上的包就可以了。我们事先已经知道， Nginx 监听在 80 端口，它所提供的 HTTP 服务是基于 TCP 协议的，所以我们可以指定 TCP 协议和 80 端口精确抓包。</p>

<p>接下来，我们在第一个终端中运行 tcpdump 命令，通过 -i eth0 选项指定网卡 eth0，并通过 tcp port 80 选项指定 TCP 协议的 80 端口：</p>

<pre><code># -i eth0 只抓取eth0网卡，-n不解析协议名和主机名
# tcp port 80表示只抓取tcp协议并且端口号为80的网络帧
$ tcpdump -i eth0 -n tcp port 80
15:11:32.678966 IP 192.168.0.2.18238 &gt; 192.168.0.30.80: Flags [S], seq 458303614, win 512, length 0
...
</code></pre>

<p>从 tcpdump 的输出中，你可以发现</p>

<ul>
<li><p>192.168.0.2.18238 &gt; 192.168.0.30.80 ，表示网络帧从 192.168.0.2 的 18238 端口发送到 192.168.0.30 的 80 端口，也就是从运行 hping3 机器的 18238 端口发送网络帧，目的为 Nginx 所在机器的 80 端口。</p></li>

<li><p>Flags [S] 则表示这是一个 SYN 包。</p></li>
</ul>

<p>再加上前面用 sar 发现的， PPS 超过 12000的现象，现在我们可以确认，这就是从 192.168.0.2 这个地址发送过来的 SYN FLOOD 攻击。</p>

<p>到这里，我们已经做了全套的性能诊断和分析。从系统的软中断使用率高这个现象出发，通过观察 /proc/softirqs 文件的变化情况，判断出软中断类型是网络接收中断；再通过 sar 和 tcpdump ，确认这是一个 SYN FLOOD 问题。</p>

<p>SYN FLOOD 问题最简单的解决方法，就是从交换机或者硬件防火墙中封掉来源 IP，这样 SYN FLOOD 网络帧就不会发送到服务器中。</p>

<p>至于 SYN FLOOD 的原理和更多解决思路，你暂时不需要过多关注，后面的网络章节里我们都会学到。</p>

<p>案例结束后，也不要忘了收尾，记得停止最开始启动的 Nginx 服务以及 hping3 命令。</p>

<p>在第一个终端中，运行下面的命令就可以停止 Nginx 了：</p>

<pre><code># 停止 Nginx 服务
$ docker rm -f nginx
</code></pre>

<p>然后到第二个终端中按下 Ctrl+C 就可以停止 hping3。</p>

<h2 id="小结">小结</h2>

<p>软中断CPU使用率（softirq）升高是一种很常见的性能问题。虽然软中断的类型很多，但实际生产中，我们遇到的性能瓶颈大多是网络收发类型的软中断，特别是网络接收的软中断。</p>

<p>在碰到这类问题时，你可以借用 sar、tcpdump 等工具，做进一步分析。不要害怕网络性能，后面我会教你更多的分析方法。</p>

<h2 id="思考">思考</h2>

<p>最后，我想请你一起来聊聊，你所碰到的软中断问题。你所碰到的软中问题是哪种类型，是不是这个案例中的小包问题？你又是怎么分析它们的来源并解决的呢？可以结合今天的案例，总结你自己的思路和感受。如果遇到过其他问题，也可以留言给我一起解决。</p>

<p>欢迎在留言区和我讨论，也欢迎你把这篇文章分享给你的同事、朋友。我们一起在实战中演练，在交流中进步。</p>

<p><img src="assets/b7f27dc9584c4462b69adda51fd4b946.jpg" alt="" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#deb2b2b2e7eaefefeee99eb9b3bfb7b2f0bdb1b3" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f10e815d86c6365',t:'MTczNDAzOTk4MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
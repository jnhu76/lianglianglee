<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=25&#32;基础篇：Linux&#32;磁盘I_O是怎么工作的（下）>
        <link rel="icon" href="/static/favicon.png">
        <title>25 基础篇：Linux 磁盘I_O是怎么工作的（下） </title>
        
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
                            <h1 id="title" data-id="25 基础篇：Linux 磁盘I_O是怎么工作的（下）" class="title">25 基础篇：Linux 磁盘I_O是怎么工作的（下）</h1>
                            <div><p>你好，我是倪朋飞。</p>

<p>上一节我们学习了 Linux 磁盘 I/O 的工作原理，并了解了由文件系统层、通用块层和设备层构成的 Linux 存储系统 I/O 栈。</p>

<p>其中，通用块层是 Linux 磁盘 I/O 的核心。向上，它为文件系统和应用程序，提供访问了块设备的标准接口；向下，把各种异构的磁盘设备，抽象为统一的块设备，并会对文件系统和应用程序发来的 I/O 请求，进行重新排序、请求合并等，提高了磁盘访问的效率。</p>

<p>掌握了磁盘 I/O 的工作原理，你估计迫不及待想知道，怎么才能衡量磁盘的 I/O 性能。</p>

<p>接下来，我们就来看看，磁盘的性能指标，以及观测这些指标的方法。</p>

<h2 id="磁盘性能指标">磁盘性能指标</h2>

<p>说到磁盘性能的衡量标准，必须要提到五个常见指标，也就是我们经常用到的，使用率、饱和度、IOPS、吞吐量以及响应时间等。这五个指标，是衡量磁盘性能的基本指标。</p>

<ul>
<li><p>使用率，是指磁盘处理I/O的时间百分比。过高的使用率（比如超过80%），通常意味着磁盘 I/O 存在性能瓶颈。</p></li>

<li><p>饱和度，是指磁盘处理 I/O 的繁忙程度。过高的饱和度，意味着磁盘存在严重的性能瓶颈。当饱和度为 100% 时，磁盘无法接受新的 I/O 请求。</p></li>

<li><p>IOPS（Input/Output Per Second），是指每秒的 I/O 请求数。</p></li>

<li><p>吞吐量，是指每秒的 I/O 请求大小。</p></li>

<li><p>响应时间，是指 I/O 请求从发出到收到响应的间隔时间。</p></li>
</ul>

<p>这里要注意的是，使用率只考虑有没有 I/O，而不考虑 I/O 的大小。换句话说，当使用率是 100% 的时候，磁盘依然有可能接受新的 I/O 请求。</p>

<p>这些指标，很可能是你经常挂在嘴边的，一讨论磁盘性能必定提起的对象。不过我还是要强调一点，不要孤立地去比较某一指标，而要结合读写比例、I/O类型（随机还是连续）以及 I/O 的大小，综合来分析。</p>

<p>举个例子，在数据库、大量小文件等这类随机读写比较多的场景中，IOPS 更能反映系统的整体性能；而在多媒体等顺序读写较多的场景中，吞吐量才更能反映系统的整体性能。</p>

<p>一般来说，我们在为应用程序的服务器选型时，要先对磁盘的 I/O 性能进行基准测试，以便可以准确评估，磁盘性能是否可以满足应用程序的需求。</p>

<p>这一方面，我推荐用性能测试工具 fio ，来测试磁盘的IOPS、吞吐量以及响应时间等核心指标。但还是那句话，因地制宜，灵活选取。在基准测试时，一定要注意根据应用程序 I/O 的特点，来具体评估指标。</p>

<p>当然，这就需要你测试出，不同 I/O 大小（一般是 512B 至 1MB 中间的若干值）分别在随机读、顺序读、随机写、顺序写等各种场景下的性能情况。</p>

<p>用性能工具得到的这些指标，可以作为后续分析应用程序性能的依据。一旦发生性能问题，你就可以把它们作为磁盘性能的极限值，进而评估磁盘 I/O 的使用情况。</p>

<p>了解磁盘的性能指标，只是我们I/O性能测试的第一步。接下来，又该用什么方法来观测它们呢？这里，我给你介绍几个常用的I/O性能观测方法。</p>

<h2 id="磁盘i-o观测"><strong>磁盘I/O观测</strong></h2>

<p>第一个要观测的，是每块磁盘的使用情况。</p>

<p>iostat 是最常用的磁盘I/O性能观测工具，它提供了每个磁盘的使用率、IOPS、吞吐量等各种常见的性能指标，当然，这些指标实际上来自 /proc/diskstats。</p>

<p>iostat 的输出界面如下。</p>

<pre><code># -d -x表示显示所有磁盘I/O的指标
$ iostat -d -x 1 
Device            r/s     w/s     rkB/s     wkB/s   rrqm/s   wrqm/s  %rrqm  %wrqm r_await w_await aqu-sz rareq-sz wareq-sz  svctm  %util 
loop0            0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00 
loop1            0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00 
sda              0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00 
sdb              0.00    0.00      0.00      0.00     0.00     0.00   0.00   0.00    0.00    0.00   0.00     0.00     0.00   0.00   0.00 
</code></pre>

<p>从这里你可以看到，iostat 提供了非常丰富的性能指标。第一列的 Device 表示磁盘设备的名字，其他各列指标，虽然数量较多，但是每个指标的含义都很重要。为了方便你理解，我把它们总结成了一个表格。</p>

<p><img src="assets/0122db0b86ba4a6684ad0da76336a176.jpg" alt="" /></p>

<p>这些指标中，你要注意：</p>

<ul>
<li><p>%util ，就是我们前面提到的磁盘I/O使用率；</p></li>

<li><p>r/s+ w/s ，就是 IOPS；</p></li>

<li><p>rkB/s+wkB/s ，就是吞吐量；</p></li>

<li><p>r_await+w_await ，就是响应时间。</p></li>
</ul>

<p>在观测指标时，也别忘了结合请求的大小（ rareq-sz 和wareq-sz）一起分析。</p>

<p>你可能注意到，从 iostat 并不能直接得到磁盘饱和度。事实上，饱和度通常也没有其他简单的观测方法，不过，你可以把观测到的，平均请求队列长度或者读写请求完成的等待时间，跟基准测试的结果（比如通过 fio）进行对比，综合评估磁盘的饱和情况。</p>

<h2 id="进程i-o观测"><strong>进程I/O观测</strong></h2>

<p>除了每块磁盘的 I/O 情况，每个进程的 I/O 情况也是我们需要关注的重点。</p>

<p>上面提到的 iostat 只提供磁盘整体的 I/O 性能数据，缺点在于，并不能知道具体是哪些进程在进行磁盘读写。要观察进程的I/O情况，你还可以使用 pidstat 和 iotop 这两个工具。</p>

<p>pidstat 是我们的老朋友了，这里我就不再啰嗦它的功能了。给它加上 -d 参数，你就可以看到进程的I/O情况，如下所示：</p>

<pre><code>$ pidstat -d 1 
13:39:51      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command 
13:39:52      102       916      0.00      4.00      0.00       0  rsyslogd
</code></pre>

<p>从pidstat的输出你能看到，它可以实时查看每个进程的I/O情况，包括下面这些内容。</p>

<ul>
<li><p>用户ID（UID）和进程ID（PID） 。</p></li>

<li><p>每秒读取的数据大小（kB_rd/s） ，单位是 KB。</p></li>

<li><p>每秒发出的写请求数据大小（kB_wr/s） ，单位是 KB。</p></li>

<li><p>每秒取消的写请求数据大小（kB_ccwr/s） ，单位是 KB。</p></li>

<li><p>块I/O延迟（iodelay），包括等待同步块I/O和换入块I/O结束的时间，单位是时钟周期。</p></li>
</ul>

<p>除了可以用 pidstat 实时查看，根据 I/O 大小对进程排序，也是性能分析中一个常用的方法。这一点，我推荐另一个工具， iotop。它是一个类似于 top 的工具，你可以按照 I/O 大小对进程排序，然后找到I/O较大的那些进程。</p>

<p>iotop 的输出如下所示：</p>

<pre><code>$ iotop
Total DISK READ :       0.00 B/s | Total DISK WRITE :       7.85 K/s 
Actual DISK READ:       0.00 B/s | Actual DISK WRITE:       0.00 B/s 
  TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN     IO&gt;    COMMAND 
15055 be/3 root        0.00 B/s    7.85 K/s  0.00 %  0.00 % systemd-journald 
</code></pre>

<p>从这个输出，你可以看到，前两行分别表示，进程的磁盘读写大小总数和磁盘真实的读写大小总数。因为缓存、缓冲区、I/O合并等因素的影响，它们可能并不相等。</p>

<p>剩下的部分，则是从各个角度来分别表示进程的I/O情况，包括线程ID、I/O优先级、每秒读磁盘的大小、每秒写磁盘的大小、换入和等待I/O的时钟百分比等。</p>

<p>这两个工具，是我们分析磁盘 I/O 性能时最常用到的。你先了解它们的功能和指标含义，具体的使用方法，接下来的案例实战中我们一起学习。</p>

<h2 id="小结">小结</h2>

<p>今天，我们梳理了 Linux 磁盘 I/O 的性能指标和性能工具。我们通常用IOPS、吞吐量、使用率、饱和度以及响应时间等几个指标，来评估磁盘的 I/O 性能。</p>

<p>你可以用 iostat 获得磁盘的 I/O 情况，也可以用 pidstat、iotop 等观察进程的 I/O 情况。不过在分析这些性能指标时，你要注意结合读写比例、I/O 类型以及 I/O 大小等，进行综合分析。</p>

<h2 id="思考">思考</h2>

<p>最后，我想请你一起来聊聊，你碰到过的磁盘 I/O 问题。在碰到磁盘 I/O 性能问题时，你是怎么分析和定位的呢？你可以结合今天学到的磁盘 I/O 指标和工具，以及上一节学过的磁盘 I/O 原理，来总结你的思路。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#4f232323767b7e7e7f780f28222e2623612c2022" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f10ecd56f176365',t:'MTczNDA0MDE3NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
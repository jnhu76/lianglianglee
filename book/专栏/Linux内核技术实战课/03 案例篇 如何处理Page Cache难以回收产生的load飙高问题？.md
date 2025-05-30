<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=03&#32;案例篇&#32;如何处理Page&#32;Cache难以回收产生的load飙高问题？>
        <link rel="icon" href="/static/favicon.png">
        <title>03 案例篇 如何处理Page Cache难以回收产生的load飙高问题？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 如何让Linux内核更好地服务应用程序？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%a6%82%e4%bd%95%e8%ae%a9Linux%e5%86%85%e6%a0%b8%e6%9b%b4%e5%a5%bd%e5%9c%b0%e6%9c%8d%e5%8a%a1%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">00 开篇词 如何让Linux内核更好地服务应用程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 基础篇 如何用数据观测Page Cache？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/01%20%e5%9f%ba%e7%a1%80%e7%af%87%20%e5%a6%82%e4%bd%95%e7%94%a8%e6%95%b0%e6%8d%ae%e8%a7%82%e6%b5%8bPage%20Cache%ef%bc%9f.md">01 基础篇 如何用数据观测Page Cache？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 基础篇 Page Cache是怎样产生和释放的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/02%20%e5%9f%ba%e7%a1%80%e7%af%87%20Page%20Cache%e6%98%af%e6%80%8e%e6%a0%b7%e4%ba%a7%e7%94%9f%e5%92%8c%e9%87%8a%e6%94%be%e7%9a%84%ef%bc%9f.md">02 基础篇 Page Cache是怎样产生和释放的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 案例篇 如何处理Page Cache难以回收产生的load飙高问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/03%20%e6%a1%88%e4%be%8b%e7%af%87%20%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86Page%20Cache%e9%9a%be%e4%bb%a5%e5%9b%9e%e6%94%b6%e4%ba%a7%e7%94%9f%e7%9a%84load%e9%a3%99%e9%ab%98%e9%97%ae%e9%a2%98%ef%bc%9f.md">03 案例篇 如何处理Page Cache难以回收产生的load飙高问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 案例篇 如何处理Page Cache容易回收引起的业务性能问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/04%20%e6%a1%88%e4%be%8b%e7%af%87%20%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86Page%20Cache%e5%ae%b9%e6%98%93%e5%9b%9e%e6%94%b6%e5%bc%95%e8%b5%b7%e7%9a%84%e4%b8%9a%e5%8a%a1%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%ef%bc%9f.md">04 案例篇 如何处理Page Cache容易回收引起的业务性能问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 分析篇 如何判断问题是否由Page Cache产生的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/05%20%e5%88%86%e6%9e%90%e7%af%87%20%e5%a6%82%e4%bd%95%e5%88%a4%e6%96%ad%e9%97%ae%e9%a2%98%e6%98%af%e5%90%a6%e7%94%b1Page%20Cache%e4%ba%a7%e7%94%9f%e7%9a%84%ef%bc%9f.md">05 分析篇 如何判断问题是否由Page Cache产生的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  基础篇 进程的哪些内存类型容易引起内存泄漏？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/06%20%20%e5%9f%ba%e7%a1%80%e7%af%87%20%e8%bf%9b%e7%a8%8b%e7%9a%84%e5%93%aa%e4%ba%9b%e5%86%85%e5%ad%98%e7%b1%bb%e5%9e%8b%e5%ae%b9%e6%98%93%e5%bc%95%e8%b5%b7%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%ef%bc%9f.md">06  基础篇 进程的哪些内存类型容易引起内存泄漏？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 案例篇 如何预防内存泄漏导致的系统假死？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/07%20%e6%a1%88%e4%be%8b%e7%af%87%20%e5%a6%82%e4%bd%95%e9%a2%84%e9%98%b2%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%e5%af%bc%e8%87%b4%e7%9a%84%e7%b3%bb%e7%bb%9f%e5%81%87%e6%ad%bb%ef%bc%9f.md">07 案例篇 如何预防内存泄漏导致的系统假死？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 案例篇 Shmem：进程没有消耗内存，内存哪去了？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/08%20%e6%a1%88%e4%be%8b%e7%af%87%20Shmem%ef%bc%9a%e8%bf%9b%e7%a8%8b%e6%b2%a1%e6%9c%89%e6%b6%88%e8%80%97%e5%86%85%e5%ad%98%ef%bc%8c%e5%86%85%e5%ad%98%e5%93%aa%e5%8e%bb%e4%ba%86%ef%bc%9f.md">08 案例篇 Shmem：进程没有消耗内存，内存哪去了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 分析篇 如何对内核内存泄漏做些基础的分析？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/09%20%e5%88%86%e6%9e%90%e7%af%87%20%e5%a6%82%e4%bd%95%e5%af%b9%e5%86%85%e6%a0%b8%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%e5%81%9a%e4%ba%9b%e5%9f%ba%e7%a1%80%e7%9a%84%e5%88%86%e6%9e%90%ef%bc%9f.md">09 分析篇 如何对内核内存泄漏做些基础的分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 分析篇 内存泄漏时，我们该如何一步步找到根因？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/10%20%e5%88%86%e6%9e%90%e7%af%87%20%e5%86%85%e5%ad%98%e6%b3%84%e6%bc%8f%e6%97%b6%ef%bc%8c%e6%88%91%e4%bb%ac%e8%af%a5%e5%a6%82%e4%bd%95%e4%b8%80%e6%ad%a5%e6%ad%a5%e6%89%be%e5%88%b0%e6%a0%b9%e5%9b%a0%ef%bc%9f.md">10 分析篇 内存泄漏时，我们该如何一步步找到根因？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 基础篇 TCP连接的建立和断开受哪些系统配置影响？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/11%20%e5%9f%ba%e7%a1%80%e7%af%87%20TCP%e8%bf%9e%e6%8e%a5%e7%9a%84%e5%bb%ba%e7%ab%8b%e5%92%8c%e6%96%ad%e5%bc%80%e5%8f%97%e5%93%aa%e4%ba%9b%e7%b3%bb%e7%bb%9f%e9%85%8d%e7%bd%ae%e5%bd%b1%e5%93%8d%ef%bc%9f.md">11 基础篇 TCP连接的建立和断开受哪些系统配置影响？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 基础篇 TCP收发包过程会受哪些配置项影响？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/12%20%e5%9f%ba%e7%a1%80%e7%af%87%20TCP%e6%94%b6%e5%8f%91%e5%8c%85%e8%bf%87%e7%a8%8b%e4%bc%9a%e5%8f%97%e5%93%aa%e4%ba%9b%e9%85%8d%e7%bd%ae%e9%a1%b9%e5%bd%b1%e5%93%8d%ef%bc%9f.md">12 基础篇 TCP收发包过程会受哪些配置项影响？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 案例篇 TCP拥塞控制是如何导致业务性能抖动的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/13%20%e6%a1%88%e4%be%8b%e7%af%87%20TCP%e6%8b%a5%e5%a1%9e%e6%8e%a7%e5%88%b6%e6%98%af%e5%a6%82%e4%bd%95%e5%af%bc%e8%87%b4%e4%b8%9a%e5%8a%a1%e6%80%a7%e8%83%bd%e6%8a%96%e5%8a%a8%e7%9a%84%ef%bc%9f.md">13 案例篇 TCP拥塞控制是如何导致业务性能抖动的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 案例篇  TCP端到端时延变大，怎样判断是哪里出现了问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/14%20%e6%a1%88%e4%be%8b%e7%af%87%20%20TCP%e7%ab%af%e5%88%b0%e7%ab%af%e6%97%b6%e5%bb%b6%e5%8f%98%e5%a4%a7%ef%bc%8c%e6%80%8e%e6%a0%b7%e5%88%a4%e6%96%ad%e6%98%af%e5%93%aa%e9%87%8c%e5%87%ba%e7%8e%b0%e4%ba%86%e9%97%ae%e9%a2%98%ef%bc%9f.md">14 案例篇  TCP端到端时延变大，怎样判断是哪里出现了问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 分析篇 如何高效地分析TCP重传问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/15%20%e5%88%86%e6%9e%90%e7%af%87%20%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%9c%b0%e5%88%86%e6%9e%90TCP%e9%87%8d%e4%bc%a0%e9%97%ae%e9%a2%98%ef%bc%9f.md">15 分析篇 如何高效地分析TCP重传问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 套路篇 如何分析常见的TCP问题？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/16%20%e5%a5%97%e8%b7%af%e7%af%87%20%e5%a6%82%e4%bd%95%e5%88%86%e6%9e%90%e5%b8%b8%e8%a7%81%e7%9a%84TCP%e9%97%ae%e9%a2%98%ef%bc%9f.md">16 套路篇 如何分析常见的TCP问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 基础篇 CPU是如何执行任务的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/17%20%e5%9f%ba%e7%a1%80%e7%af%87%20CPU%e6%98%af%e5%a6%82%e4%bd%95%e6%89%a7%e8%a1%8c%e4%bb%bb%e5%8a%a1%e7%9a%84%ef%bc%9f.md">17 基础篇 CPU是如何执行任务的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 案例篇 业务是否需要使用透明大页：水可载舟，亦可覆舟？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/18%20%e6%a1%88%e4%be%8b%e7%af%87%20%e4%b8%9a%e5%8a%a1%e6%98%af%e5%90%a6%e9%9c%80%e8%a6%81%e4%bd%bf%e7%94%a8%e9%80%8f%e6%98%8e%e5%a4%a7%e9%a1%b5%ef%bc%9a%e6%b0%b4%e5%8f%af%e8%bd%bd%e8%88%9f%ef%bc%8c%e4%ba%a6%e5%8f%af%e8%a6%86%e8%88%9f%ef%bc%9f.md">18 案例篇 业务是否需要使用透明大页：水可载舟，亦可覆舟？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 案例篇 网络吞吐高的业务是否需要开启网卡特性呢？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/19%20%e6%a1%88%e4%be%8b%e7%af%87%20%e7%bd%91%e7%bb%9c%e5%90%9e%e5%90%90%e9%ab%98%e7%9a%84%e4%b8%9a%e5%8a%a1%e6%98%af%e5%90%a6%e9%9c%80%e8%a6%81%e5%bc%80%e5%90%af%e7%bd%91%e5%8d%a1%e7%89%b9%e6%80%a7%e5%91%a2%ef%bc%9f.md">19 案例篇 网络吞吐高的业务是否需要开启网卡特性呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 分析篇 如何分析CPU利用率飙高问题 ？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/20%20%e5%88%86%e6%9e%90%e7%af%87%20%e5%a6%82%e4%bd%95%e5%88%86%e6%9e%90CPU%e5%88%a9%e7%94%a8%e7%8e%87%e9%a3%99%e9%ab%98%e9%97%ae%e9%a2%98%20%ef%bc%9f.md">20 分析篇 如何分析CPU利用率飙高问题 ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 我是如何使用tracepoint来分析内核Bug的？.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%90%20%e6%88%91%e6%98%af%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8tracepoint%e6%9d%a5%e5%88%86%e6%9e%90%e5%86%85%e6%a0%b8Bug%e7%9a%84%ef%bc%9f.md">加餐 我是如何使用tracepoint来分析内核Bug的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 第一次看内核代码，我也很懵逼.md" href="/%e4%b8%93%e6%a0%8f/Linux%e5%86%85%e6%a0%b8%e6%8a%80%e6%9c%af%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e7%ac%ac%e4%b8%80%e6%ac%a1%e7%9c%8b%e5%86%85%e6%a0%b8%e4%bb%a3%e7%a0%81%ef%bc%8c%e6%88%91%e4%b9%9f%e5%be%88%e6%87%b5%e9%80%bc.md">结束语 第一次看内核代码，我也很懵逼.md</a>
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
                            <h1 id="title" data-id="03 案例篇 如何处理Page Cache难以回收产生的load飙高问题？" class="title">03 案例篇 如何处理Page Cache难以回收产生的load飙高问题？</h1>
                            <div><p>你好，我是邵亚方。今天这节课，我想跟你聊一聊怎么处理在生产环境中，因为Page Cache管理不当引起的系统load飙高的问题。</p>

<p>相信你在平时的工作中，应该会或多或少遇到过这些情形：系统很卡顿，敲命令响应非常慢；应用程序的RT变得很高，或者抖动得很厉害。在发生这些问题时，很有可能也伴随着系统load飙得很高。</p>

<p>那这是什么原因导致的呢？据我观察，大多是有三种情况：</p>

<ul>
<li>直接内存回收引起的load飙高；</li>
<li>系统中脏页积压过多引起的load飙高；</li>
<li>系统NUMA策略配置不当引起的load飙高。</li>
</ul>

<p>这是应用开发者和运维人员向我咨询最多的几种情况。问题看似很简单，但如果对问题产生的原因理解得不深，解决起来就会很棘手，甚至配置得不好，还会带来负面的影响。</p>

<p>所以这节课，我们一起来分析下这三种情况，可以说，搞清楚了这几种情况，你差不多就能解决掉绝大部分Page Cache引起的load飙高问题了。如果你对问题的原因排查感兴趣，也不要着急，在第5讲，我会带你学习load飙高问题的分析方法。</p>

<p>接下来，我们就来逐一分析下这几类情况。</p>

<h2 id="直接内存回收引起load飙高或者业务时延抖动">直接内存回收引起load飙高或者业务时延抖动</h2>

<p>直接内存回收是指在进程上下文同步进行内存回收，那么它具体是怎么引起load飙高的呢？</p>

<p>因为直接内存回收是在进程申请内存的过程中同步进行的回收，而这个回收过程可能会消耗很多时间，进而导致进程的后续行为都被迫等待，这样就会造成很长时间的延迟，以及系统的CPU利用率会升高，最终引起load飙高。</p>

<p>我们详细地描述一下这个过程，为了尽量不涉及太多技术细节，我会用一张图来表示，这样你理解起来会更容易。</p>

<p><img src="assets/fe84eb2bd4956bbbdd5b0259df8c9400.jpg" alt="" title="内存回收过程" /></p>

<p>从图里你可以看到，在开始内存回收后，首先进行后台异步回收（上图中蓝色标记的地方），这不会引起进程的延迟；如果后台异步回收跟不上进程内存申请的速度，就会开始同步阻塞回收，导致延迟（上图中红色和粉色标记的地方，这就是引起load高的地址）。</p>

<p>那么，针对直接内存回收引起load飙高或者业务RT抖动的问题，一个解决方案就是<strong>及早地触发后台回收来避免应用程序进行直接内存回收，那具体要怎么做呢？</strong></p>

<p>我们先来了解一下后台回收的原理，如图：</p>

<p><img src="assets/44d471fdae7376eb13e6e6bfc70b3172.jpg" alt="" /></p>

<p>它的意思是：当内存水位低于watermark low时，就会唤醒kswapd进行后台回收，然后kswapd会一直回收到watermark high。</p>

<p>那么，我们可以增大min_free_kbytes这个配置选项来及早地触发后台回收，该选项最终控制的是内存回收水位，不过，内存回收水位是内核里面非常细节性的知识点，我们可以先不去讨论。</p>

<blockquote>
<p>vm.min_free_kbytes = 4194304</p>
</blockquote>

<p>对于大于等于128G的系统而言，将min_free_kbytes设置为4G比较合理，这是我们在处理很多这种问题时总结出来的一个经验值，既不造成较多的内存浪费，又能避免掉绝大多数的直接内存回收。</p>

<p>该值的设置和总的物理内存并没有一个严格对应的关系，我们在前面也说过，如果配置不当会引起一些副作用，所以在调整该值之前，我的建议是：你可以渐进式地增大该值，比如先调整为1G，观察sar -B中pgscand是否还有不为0的情况；如果存在不为0的情况，继续增加到2G，再次观察是否还有不为0的情况来决定是否增大，以此类推。</p>

<p>在这里你需要注意的是，即使将该值增加得很大，还是可能存在pgscand不为0的情况（这个略复杂，涉及到内存碎片和连续内存申请，我们在此先不展开，你知道有这么回事儿就可以了）。那么这个时候你要考虑的是，业务是否可以容忍，如果可以容忍那就没有必要继续增加了，也就是说，增大该值并不是完全避免直接内存回收，而是尽量将直接内存回收行为控制在业务可以容忍的范围内。</p>

<p>这个方法可以用在3.10.0以后的内核上（对应的操作系统为CentOS-7以及之后更新的操作系统）。</p>

<p>当然了，这样做也有一些缺陷：提高了内存水位后，应用程序可以直接使用的内存量就会减少，这在一定程度上浪费了内存。所以在调整这一项之前，你需要先思考一下，<strong>应用程序更加关注什么，如果关注延迟那就适当地增大该值，如果关注内存的使用量那就适当地调小该值。</strong></p>

<p>除此之外，对CentOS-6(对应于2.6.32内核版本)而言，还有另外一种解决方案：</p>

<blockquote>
<p>vm.extra_free_kbytes = 4194304</p>
</blockquote>

<p>那就是将extra_free_kbytes 配置为4G。extra_free_kbytes在3.10以及以后的内核上都被废弃掉了，不过由于在生产环境中还存在大量的机器运行着较老版本内核，你使用到的也可能会是较老版本的内核，所以在这里还是有必要提一下。它的大致原理如下所示：</p>

<p><img src="assets/7d9e537e23489cd4f5f34fedcd6f89d2.jpg" alt="" /></p>

<p>extra_free_kbytes的目的是为了解决min_free_kbyte造成的内存浪费，但是这种做法并没有被内核主线接收，因为这种行为很难维护会带来一些麻烦，感兴趣的可以看一下这个讨论：<a href="https://lkml.org/lkml/2013/2/17/210" target="_blank">add extra free kbytes tunable</a></p>

<p>总的来说，通过调整内存水位，在一定程度上保障了应用的内存申请，但是同时也带来了一定的内存浪费，因为系统始终要保障有这么多的free内存，这就压缩了Page Cache的空间。调整的效果你可以通过/proc/zoneinfo来观察：</p>

<pre><code>$ egrep &quot;min|low|high&quot; /proc/zoneinfo 
...
        min      7019
        low      8773
        high     10527
...
</code></pre>

<p>其中min、low、high分别对应上图中的三个内存水位。你可以观察一下调整前后min、low、high的变化。需要提醒你的是，内存水位是针对每个内存zone进行设置的，所以/proc/zoneinfo里面会有很多zone以及它们的内存水位，你可以不用去关注这些细节。</p>

<h2 id="系统中脏页过多引起load飙高">系统中脏页过多引起load飙高</h2>

<p>接下来，我们分析下由于系统脏页过多引起load飙高的情况。在前一个案例中我们也提到，直接回收过程中，如果存在较多脏页就可能涉及在回收过程中进行回写，这可能会造成非常大的延迟，而且因为这个过程本身是阻塞式的，所以又可能进一步导致系统中处于D状态的进程数增多，最终的表现就是系统的load值很高。</p>

<p>我们来看一下这张图，这是一个典型的脏页引起系统load值飙高的问题场景：</p>

<p><img src="assets/90c693c95d67cfaf89b86edbd1228d75.jpg" alt="" /></p>

<p>如图所示，如果系统中既有快速I/O设备，又有慢速I/O设备（比如图中的ceph RBD设备，或者其他慢速存储设备比如HDD），直接内存回收过程中遇到了正在往慢速I/O设备回写的page，就可能导致非常大的延迟。</p>

<p>这里我多说一点。这类问题其实是不太好去追踪的，为了更好追踪这种慢速I/O设备引起的抖动问题，我也给Linux Kernel提交了一个patch来进行更好的追踪：<a href="https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=19343b5bdd16ad4ae6b845ef829f68b683c4dfb5" target="_blank">mm/page-writeback: introduce tracepoint for wait_on_page_writeback()</a>，这种做法是在原来的基础上增加了回写的设备，这样子用户就能更好地将回写和具体设备关联起来，从而判断问题是否是由慢速I/O设备导致的（具体的分析方法我会在后面第5讲分析篇里重点来讲）。</p>

<p>那如何解决这类问题呢？一个比较省事的解决方案是控制好系统中积压的脏页数据。很多人知道需要控制脏页，但是往往并不清楚如何来控制好这个度，脏页控制的少了可能会影响系统整体的效率，脏页控制的多了还是会触发问题，所以我们接下来看下如何来衡量好这个“度”。</p>

<p>首先你可以通过sar -r来观察系统中的脏页个数：</p>

<pre><code>$ sar -r 1
07:30:01 PM kbmemfree kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
09:20:01 PM   5681588   2137312     27.34         0   1807432    193016      2.47    534416   1310876         4
09:30:01 PM   5677564   2141336     27.39         0   1807500    204084      2.61    539192   1310884        20
09:40:01 PM   5679516   2139384     27.36         0   1807508    196696      2.52    536528   1310888        20
09:50:01 PM   5679548   2139352     27.36         0   1807516    196624      2.51    536152   1310892        24
</code></pre>

<p>kbdirty就是系统中的脏页大小，它同样也是对/proc/vmstat中nr_dirty的解析。你可以通过调小如下设置来将系统脏页个数控制在一个合理范围:</p>

<blockquote>
<p>vm.dirty_background_bytes = 0-
vm.dirty_background_ratio = 10-
vm.dirty_bytes = 0-
vm.dirty_expire_centisecs = 3000-
vm.dirty_ratio = 20</p>
</blockquote>

<p>调整这些配置项有利有弊，调大这些值会导致脏页的积压，但是同时也可能减少了I/O的次数，从而提升单次刷盘的效率；调小这些值可以减少脏页的积压，但是同时也增加了I/O的次数，降低了I/O的效率。</p>

<p><strong>至于这些值调整大多少比较合适，也是因系统和业务的不同而异，我的建议也是一边调整一边观察，将这些值调整到业务可以容忍的程度就可以了，即在调整后需要观察业务的服务质量(SLA)，要确保SLA在可接受范围内</strong>。调整的效果你可以通过/proc/vmstat来查看：</p>

<pre><code>$ grep &quot;nr_dirty_&quot; /proc/vmstat
nr_dirty_threshold 366998
nr_dirty_background_threshold 183275
</code></pre>

<p>你可以观察一下调整前后这两项的变化。<strong>这里我要给你一个避免踩坑的提示</strong>，解决该方案中的设置项如果设置不妥会触发一个内核Bug，这是我在2017年进行性能调优时发现的一个内核Bug，我给社区提交了一个patch将它fix掉了，具体的commit见<a href="https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=94af584692091347baea4d810b9fc6e0f5483d42" target="_blank">writeback: schedule periodic writeback with sysctl</a> <a href="https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=94af584692091347baea4d810b9fc6e0f5483d42" target="_blank"></a>, commit log清晰地描述了该问题，我建议你有时间看一看。</p>

<h2 id="系统numa策略配置不当引起的load飙高">系统NUMA策略配置不当引起的load飙高</h2>

<p>除了我前面提到的这两种引起系统load飙高或者业务延迟抖动的场景之外，还有另外一种场景也会引起load飙高，那就是系统NUMA策略配置不当引起的load飙高。</p>

<p>比如说，我们在生产环境上就曾经遇到这样的问题：系统中还有一半左右的free内存，但还是频频触发direct reclaim，导致业务抖动得比较厉害。后来经过排查发现是由于设置了zone_reclaim_mode，这是NUMA策略的一种。</p>

<p>设置zone_reclaim_mode的目的是为了增加业务的NUMA亲和性，但是在实际生产环境中很少会有对NUMA特别敏感的业务，这也是为什么内核将该配置从默认配置1修改为了默认配置0: <a href="https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=4f9b16a64753d0bb607454347036dc997fd03b82" target="_blank">mm: disable zone_reclaim_mode by default</a> ，配置为0之后，就避免了在其他node有空闲内存时，不去使用这些空闲内存而是去回收当前node的Page Cache，也就是说，通过减少内存回收发生的可能性从而避免它引发的业务延迟。</p>

<p>那么如何来有效地衡量业务延迟问题是否由zone reclaim引起的呢？它引起的延迟究竟有多大呢？这个衡量和观察方法也是我贡献给Linux Kernel的：<a href="https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=132bb8cfc9e081238e7e2fd0c37c8c75ad0d2963" target="_blank">mm/vmscan: add tracepoints for node reclaim</a> ，大致的思路就是利用linux的tracepoint来做这种量化分析，这是性能开销相对较小的一个方案。</p>

<p>我们可以通过numactl来查看服务器的NUMA信息，如下是两个node的服务器：</p>

<pre><code>$ numactl --hardware
available: 2 nodes (0-1)
node 0 cpus: 0 1 2 3 4 5 6 7 8 9 10 11 24 25 26 27 28 29 30 31 32 33 34 35
node 0 size: 130950 MB
node 0 free: 108256 MB
node 1 cpus: 12 13 14 15 16 17 18 19 20 21 22 23 36 37 38 39 40 41 42 43 44 45 46 47
node 1 size: 131072 MB
node 1 free: 122995 MB
node distances:
node   0   1 
  0:  10  21 
  1:  21  10 
</code></pre>

<p>其中CPU0～11，24～35的local node为node 0；而CPU12～23，36～47的local node为node 1。如下图所示：</p>

<p><img src="assets/80e7c19a8f310d5bf30d368cef86bbe2.jpg" alt="" /></p>

<p>推荐将zone_reclaim_mode配置为0。</p>

<blockquote>
<p>vm.zone_reclaim_mode = 0</p>
</blockquote>

<p>因为相比内存回收的危害而言，NUMA带来的性能提升几乎可以忽略，所以配置为0，利远大于弊。</p>

<p>好了，对于Page Cache管理不当引起的系统load飙高和业务时延抖动问题，我们就分析到这里，希望通过这篇的学习，在下次你遇到直接内存回收引起的load飙高问题时不再束手无策。</p>

<p>总的来说，这些问题都是Page Cache难以释放而产生的问题，那你是否想过，是不是Page Cache很容易释放就不会产生问题了？这个答案可能会让你有些意料不到：Page Cache容易释放也有容易释放的问题。这到底是怎么回事呢，我们下节课来分析下这方面的案例。</p>

<h2 id="课堂总结">课堂总结</h2>

<p>这节课我们讲的这几个案例都是内存回收过程中引起的load飙高问题。关于内存回收这事，我们可以做一个形象的类比。我们知道，内存是操作系统中很重要的一个资源，它就像我们在生活过程中很重要的一个资源——钱一样，如果你的钱（内存）足够多，那想买什么就可以买什么，而不用担心钱花完（内存用完）后要吃土（引起load飙高）。</p>

<p>但是现实情况是我们每个人用来双十一购物的钱（内存）总是有限的，在买东西（运行程序）的时候总需要精打细算，一旦预算快超了（内存快不够了），就得把一些不重要的东西（把一些不活跃的内容）从购物车里删除掉（回收掉），好腾出资金（空闲的内存）来买更想买的东西（运行需要运行的程序）。</p>

<p>我们讲的这几个案例都可以通过调整系统参数/配置来解决，调整系统参数/配置也是应用开发者和运维人员在发生了内核问题时所能做的改动。比如说，直接内存回收引起load飙高时，就去调整内存水位设置；脏页积压引起load飙高时，就需要去调整脏页的水位；NUMA策略配置不当引起load飙高时，就去检查是否需要关闭该策略。同时我们在做这些调整的时候，一定要边调整边观察业务的服务质量，确保SLA是可以接受的。</p>

<p>如果你想要你的系统更加稳定，你的业务性能更好，你不妨去研究一下系统中的可配置项，看看哪些配置可以帮助你的业务。</p>

<h2 id="课后作业">课后作业</h2>

<p>这节课我给你布置的作业是针对直接内存回收的，现在你已经知道直接内存回收容易产生问题，是我们需要尽量避免的，那么我的问题是：请你执行一些模拟程序来构造出直接内存回收的场景（小提示： 你可以通过sar -B中的pgscand来判断是否有了直接内存回收）。欢迎在留言区分享你的看法。</p>

<p>感谢你的阅读，如果你认为这节课的内容有收获，也欢迎把它分享给你的朋友，我们下一讲见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#cfa3a3a3f6fbfefefff88fa8a2aea6a3e1aca0a2" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f10c47de8c160fe',t:'MTczNDAzODUyMi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=04&#32;震撼的Linux全景图：业界成熟的内核架构长什么样？>
        <link rel="icon" href="/static/favicon.png">
        <title>04 震撼的Linux全景图：业界成熟的内核架构长什么样？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么要学写一个操作系统？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6%e5%86%99%e4%b8%80%e4%b8%aa%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">00 开篇词 为什么要学写一个操作系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 编辑手记 升级认知，迭代自己的操作系统.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/00%20%e7%bc%96%e8%be%91%e6%89%8b%e8%ae%b0%20%e5%8d%87%e7%ba%a7%e8%ae%a4%e7%9f%a5%ef%bc%8c%e8%bf%ad%e4%bb%a3%e8%87%aa%e5%b7%b1%e7%9a%84%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f.md">00 编辑手记 升级认知，迭代自己的操作系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 程序的运行过程：从代码到机器运行.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/01%20%e7%a8%8b%e5%ba%8f%e7%9a%84%e8%bf%90%e8%a1%8c%e8%bf%87%e7%a8%8b%ef%bc%9a%e4%bb%8e%e4%bb%a3%e7%a0%81%e5%88%b0%e6%9c%ba%e5%99%a8%e8%bf%90%e8%a1%8c.md">01 程序的运行过程：从代码到机器运行.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 几行汇编几行C：实现一个最简单的内核.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/02%20%e5%87%a0%e8%a1%8c%e6%b1%87%e7%bc%96%e5%87%a0%e8%a1%8cC%ef%bc%9a%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e6%9c%80%e7%ae%80%e5%8d%95%e7%9a%84%e5%86%85%e6%a0%b8.md">02 几行汇编几行C：实现一个最简单的内核.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 黑盒之中有什么：内核结构与设计.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/03%20%e9%bb%91%e7%9b%92%e4%b9%8b%e4%b8%ad%e6%9c%89%e4%bb%80%e4%b9%88%ef%bc%9a%e5%86%85%e6%a0%b8%e7%bb%93%e6%9e%84%e4%b8%8e%e8%ae%be%e8%ae%a1.md">03 黑盒之中有什么：内核结构与设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 震撼的Linux全景图：业界成熟的内核架构长什么样？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/04%20%e9%9c%87%e6%92%bc%e7%9a%84Linux%e5%85%a8%e6%99%af%e5%9b%be%ef%bc%9a%e4%b8%9a%e7%95%8c%e6%88%90%e7%86%9f%e7%9a%84%e5%86%85%e6%a0%b8%e6%9e%b6%e6%9e%84%e9%95%bf%e4%bb%80%e4%b9%88%e6%a0%b7%ef%bc%9f.md">04 震撼的Linux全景图：业界成熟的内核架构长什么样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 CPU工作模式：执行程序的三种模式.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/05%20CPU%e5%b7%a5%e4%bd%9c%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%89%a7%e8%a1%8c%e7%a8%8b%e5%ba%8f%e7%9a%84%e4%b8%89%e7%a7%8d%e6%a8%a1%e5%bc%8f.md">05 CPU工作模式：执行程序的三种模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 虚幻与真实：程序中的地址如何转换？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/06%20%e8%99%9a%e5%b9%bb%e4%b8%8e%e7%9c%9f%e5%ae%9e%ef%bc%9a%e7%a8%8b%e5%ba%8f%e4%b8%ad%e7%9a%84%e5%9c%b0%e5%9d%80%e5%a6%82%e4%bd%95%e8%bd%ac%e6%8d%a2%ef%bc%9f.md">06 虚幻与真实：程序中的地址如何转换？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Cache与内存：程序放在哪儿？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/07%20Cache%e4%b8%8e%e5%86%85%e5%ad%98%ef%bc%9a%e7%a8%8b%e5%ba%8f%e6%94%be%e5%9c%a8%e5%93%aa%e5%84%bf%ef%bc%9f.md">07 Cache与内存：程序放在哪儿？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 锁：并发操作中，解决数据同步的四种方法.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/08%20%e9%94%81%ef%bc%9a%e5%b9%b6%e5%8f%91%e6%93%8d%e4%bd%9c%e4%b8%ad%ef%bc%8c%e8%a7%a3%e5%86%b3%e6%95%b0%e6%8d%ae%e5%90%8c%e6%ad%a5%e7%9a%84%e5%9b%9b%e7%a7%8d%e6%96%b9%e6%b3%95.md">08 锁：并发操作中，解决数据同步的四种方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 瞧一瞧Linux：Linux的自旋锁和信号量如何实现？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/09%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9aLinux%e7%9a%84%e8%87%aa%e6%97%8b%e9%94%81%e5%92%8c%e4%bf%a1%e5%8f%b7%e9%87%8f%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">09 瞧一瞧Linux：Linux的自旋锁和信号量如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 设置工作模式与环境（上）：建立计算机.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/10%20%e8%ae%be%e7%bd%ae%e5%b7%a5%e4%bd%9c%e6%a8%a1%e5%bc%8f%e4%b8%8e%e7%8e%af%e5%a2%83%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%bb%ba%e7%ab%8b%e8%ae%a1%e7%ae%97%e6%9c%ba.md">10 设置工作模式与环境（上）：建立计算机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 设置工作模式与环境（中）：建造二级引导器.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/11%20%e8%ae%be%e7%bd%ae%e5%b7%a5%e4%bd%9c%e6%a8%a1%e5%bc%8f%e4%b8%8e%e7%8e%af%e5%a2%83%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%bb%ba%e9%80%a0%e4%ba%8c%e7%ba%a7%e5%bc%95%e5%af%bc%e5%99%a8.md">11 设置工作模式与环境（中）：建造二级引导器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 设置工作模式与环境（下）：探查和收集信息.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/12%20%e8%ae%be%e7%bd%ae%e5%b7%a5%e4%bd%9c%e6%a8%a1%e5%bc%8f%e4%b8%8e%e7%8e%af%e5%a2%83%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%8e%a2%e6%9f%a5%e5%92%8c%e6%94%b6%e9%9b%86%e4%bf%a1%e6%81%af.md">12 设置工作模式与环境（下）：探查和收集信息.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 第一个C函数：如何实现板级初始化？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/13%20%e7%ac%ac%e4%b8%80%e4%b8%aaC%e5%87%bd%e6%95%b0%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%9d%bf%e7%ba%a7%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%9f.md">13 第一个C函数：如何实现板级初始化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Linux初始化（上）：GRUB与vmlinuz的结构.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/14%20Linux%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aGRUB%e4%b8%8evmlinuz%e7%9a%84%e7%bb%93%e6%9e%84.md">14 Linux初始化（上）：GRUB与vmlinuz的结构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Linux初始化（下）：从_start到第一个进程.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/15%20Linux%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%bb%8e_start%e5%88%b0%e7%ac%ac%e4%b8%80%e4%b8%aa%e8%bf%9b%e7%a8%8b.md">15 Linux初始化（下）：从_start到第一个进程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 划分土地（上）：如何划分与组织内存？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/16%20%e5%88%92%e5%88%86%e5%9c%9f%e5%9c%b0%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%92%e5%88%86%e4%b8%8e%e7%bb%84%e7%bb%87%e5%86%85%e5%ad%98%ef%bc%9f.md">16 划分土地（上）：如何划分与组织内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 划分土地（中）：如何实现内存页面初始化？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/17%20%e5%88%92%e5%88%86%e5%9c%9f%e5%9c%b0%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%86%85%e5%ad%98%e9%a1%b5%e9%9d%a2%e5%88%9d%e5%a7%8b%e5%8c%96%ef%bc%9f.md">17 划分土地（中）：如何实现内存页面初始化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 划分土地（下）：如何实现内存页的分配与释放？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/18%20%e5%88%92%e5%88%86%e5%9c%9f%e5%9c%b0%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%86%85%e5%ad%98%e9%a1%b5%e7%9a%84%e5%88%86%e9%85%8d%e4%b8%8e%e9%87%8a%e6%94%be%ef%bc%9f.md">18 划分土地（下）：如何实现内存页的分配与释放？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 土地不能浪费：如何管理内存对象？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/19%20%e5%9c%9f%e5%9c%b0%e4%b8%8d%e8%83%bd%e6%b5%aa%e8%b4%b9%ef%bc%9a%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e5%86%85%e5%ad%98%e5%af%b9%e8%b1%a1%ef%bc%9f.md">19 土地不能浪费：如何管理内存对象？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 土地需求扩大与保障：如何表示虚拟内存？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/20%20%e5%9c%9f%e5%9c%b0%e9%9c%80%e6%b1%82%e6%89%a9%e5%a4%a7%e4%b8%8e%e4%bf%9d%e9%9a%9c%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a1%a8%e7%a4%ba%e8%99%9a%e6%8b%9f%e5%86%85%e5%ad%98%ef%bc%9f.md">20 土地需求扩大与保障：如何表示虚拟内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 土地需求扩大与保障：如何分配和释放虚拟内存？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/21%20%e5%9c%9f%e5%9c%b0%e9%9c%80%e6%b1%82%e6%89%a9%e5%a4%a7%e4%b8%8e%e4%bf%9d%e9%9a%9c%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%86%e9%85%8d%e5%92%8c%e9%87%8a%e6%94%be%e8%99%9a%e6%8b%9f%e5%86%85%e5%ad%98%ef%bc%9f.md">21 土地需求扩大与保障：如何分配和释放虚拟内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 瞧一瞧Linux：伙伴系统如何分配内存？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/22%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e4%bc%99%e4%bc%b4%e7%b3%bb%e7%bb%9f%e5%a6%82%e4%bd%95%e5%88%86%e9%85%8d%e5%86%85%e5%ad%98%ef%bc%9f.md">22 瞧一瞧Linux：伙伴系统如何分配内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 瞧一瞧Linux：SLAB如何分配内存？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/23%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9aSLAB%e5%a6%82%e4%bd%95%e5%88%86%e9%85%8d%e5%86%85%e5%ad%98%ef%bc%9f.md">23 瞧一瞧Linux：SLAB如何分配内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 活动的描述：到底什么是进程？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/24%20%e6%b4%bb%e5%8a%a8%e7%9a%84%e6%8f%8f%e8%bf%b0%ef%bc%9a%e5%88%b0%e5%ba%95%e4%bb%80%e4%b9%88%e6%98%af%e8%bf%9b%e7%a8%8b%ef%bc%9f.md">24 活动的描述：到底什么是进程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 多个活动要安排（上）：多进程如何调度？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/25%20%e5%a4%9a%e4%b8%aa%e6%b4%bb%e5%8a%a8%e8%a6%81%e5%ae%89%e6%8e%92%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a4%9a%e8%bf%9b%e7%a8%8b%e5%a6%82%e4%bd%95%e8%b0%83%e5%ba%a6%ef%bc%9f.md">25 多个活动要安排（上）：多进程如何调度？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 多个活动要安排（下）：如何实现进程的等待与唤醒机制？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/26%20%e5%a4%9a%e4%b8%aa%e6%b4%bb%e5%8a%a8%e8%a6%81%e5%ae%89%e6%8e%92%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e8%bf%9b%e7%a8%8b%e7%9a%84%e7%ad%89%e5%be%85%e4%b8%8e%e5%94%a4%e9%86%92%e6%9c%ba%e5%88%b6%ef%bc%9f.md">26 多个活动要安排（下）：如何实现进程的等待与唤醒机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 瞧一瞧Linux：Linux如何实现进程与进程调度_.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/27%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9aLinux%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e8%bf%9b%e7%a8%8b%e4%b8%8e%e8%bf%9b%e7%a8%8b%e8%b0%83%e5%ba%a6_.md">27 瞧一瞧Linux：Linux如何实现进程与进程调度_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 部门分类：如何表示设备类型与设备驱动？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/28%20%e9%83%a8%e9%97%a8%e5%88%86%e7%b1%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a1%a8%e7%a4%ba%e8%ae%be%e5%a4%87%e7%b1%bb%e5%9e%8b%e4%b8%8e%e8%ae%be%e5%a4%87%e9%a9%b1%e5%8a%a8%ef%bc%9f.md">28 部门分类：如何表示设备类型与设备驱动？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 部门建立：如何在内核中注册设备？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/29%20%e9%83%a8%e9%97%a8%e5%bb%ba%e7%ab%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%86%85%e6%a0%b8%e4%b8%ad%e6%b3%a8%e5%86%8c%e8%ae%be%e5%a4%87%ef%bc%9f.md">29 部门建立：如何在内核中注册设备？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 部门响应：设备如何处理内核I_O包？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/30%20%e9%83%a8%e9%97%a8%e5%93%8d%e5%ba%94%ef%bc%9a%e8%ae%be%e5%a4%87%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e5%86%85%e6%a0%b8I_O%e5%8c%85%ef%bc%9f.md">30 部门响应：设备如何处理内核I_O包？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 瞧一瞧Linux：如何获取所有设备信息？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/31%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e5%a6%82%e4%bd%95%e8%8e%b7%e5%8f%96%e6%89%80%e6%9c%89%e8%ae%be%e5%a4%87%e4%bf%a1%e6%81%af%ef%bc%9f.md">31 瞧一瞧Linux：如何获取所有设备信息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 仓库结构：如何组织文件_.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/32%20%e4%bb%93%e5%ba%93%e7%bb%93%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e7%bb%84%e7%bb%87%e6%96%87%e4%bb%b6_.md">32 仓库结构：如何组织文件_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 仓库划分：文件系统的格式化操作.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/33%20%e4%bb%93%e5%ba%93%e5%88%92%e5%88%86%ef%bc%9a%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%a0%bc%e5%bc%8f%e5%8c%96%e6%93%8d%e4%bd%9c.md">33 仓库划分：文件系统的格式化操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 仓库管理：如何实现文件的六大基本操作？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/34%20%e4%bb%93%e5%ba%93%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%96%87%e4%bb%b6%e7%9a%84%e5%85%ad%e5%a4%a7%e5%9f%ba%e6%9c%ac%e6%93%8d%e4%bd%9c%ef%bc%9f.md">34 仓库管理：如何实现文件的六大基本操作？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 瞧一瞧Linux：虚拟文件系统如何管理文件？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/35%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e8%99%9a%e6%8b%9f%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e6%96%87%e4%bb%b6%ef%bc%9f.md">35 瞧一瞧Linux：虚拟文件系统如何管理文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 从URL到网卡：如何全局观察网络数据流动？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/36%20%e4%bb%8eURL%e5%88%b0%e7%bd%91%e5%8d%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e5%85%a8%e5%b1%80%e8%a7%82%e5%af%9f%e7%bd%91%e7%bb%9c%e6%95%b0%e6%8d%ae%e6%b5%81%e5%8a%a8%ef%bc%9f.md">36 从URL到网卡：如何全局观察网络数据流动？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 从内核到应用：网络数据在内核中如何流转.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/37%20%e4%bb%8e%e5%86%85%e6%a0%b8%e5%88%b0%e5%ba%94%e7%94%a8%ef%bc%9a%e7%bd%91%e7%bb%9c%e6%95%b0%e6%8d%ae%e5%9c%a8%e5%86%85%e6%a0%b8%e4%b8%ad%e5%a6%82%e4%bd%95%e6%b5%81%e8%bd%ac.md">37 从内核到应用：网络数据在内核中如何流转.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 从单排到团战：详解操作系统的宏观网络架构.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/38%20%e4%bb%8e%e5%8d%95%e6%8e%92%e5%88%b0%e5%9b%a2%e6%88%98%ef%bc%9a%e8%af%a6%e8%a7%a3%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e7%9a%84%e5%ae%8f%e8%a7%82%e7%bd%91%e7%bb%9c%e6%9e%b6%e6%9e%84.md">38 从单排到团战：详解操作系统的宏观网络架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 瞧一瞧Linux：详解socket实现与网络编程接口.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/39%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e8%af%a6%e8%a7%a3socket%e5%ae%9e%e7%8e%b0%e4%b8%8e%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e6%8e%a5%e5%8f%a3.md">39 瞧一瞧Linux：详解socket实现与网络编程接口.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 瞧一瞧Linux：详解socket的接口实现.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/40%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e8%af%a6%e8%a7%a3socket%e7%9a%84%e6%8e%a5%e5%8f%a3%e5%ae%9e%e7%8e%b0.md">40 瞧一瞧Linux：详解socket的接口实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 服务接口：如何搭建沟通桥梁？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/41%20%e6%9c%8d%e5%8a%a1%e6%8e%a5%e5%8f%a3%ef%bc%9a%e5%a6%82%e4%bd%95%e6%90%ad%e5%bb%ba%e6%b2%9f%e9%80%9a%e6%a1%a5%e6%a2%81%ef%bc%9f.md">41 服务接口：如何搭建沟通桥梁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 瞧一瞧Linux：如何实现系统API？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/42%20%e7%9e%a7%e4%b8%80%e7%9e%a7Linux%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e7%b3%bb%e7%bb%9fAPI%ef%bc%9f.md">42 瞧一瞧Linux：如何实现系统API？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 虚拟机内核：KVM是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/43%20%e8%99%9a%e6%8b%9f%e6%9c%ba%e5%86%85%e6%a0%b8%ef%bc%9aKVM%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">43 虚拟机内核：KVM是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 容器：如何理解容器的实现机制？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/44%20%e5%ae%b9%e5%99%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e5%ae%b9%e5%99%a8%e7%9a%84%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6%ef%bc%9f.md">44 容器：如何理解容器的实现机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 ARM新宠：苹果的M1芯片因何而快？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/45%20ARM%e6%96%b0%e5%ae%a0%ef%bc%9a%e8%8b%b9%e6%9e%9c%e7%9a%84M1%e8%8a%af%e7%89%87%e5%9b%a0%e4%bd%95%e8%80%8c%e5%bf%ab%ef%bc%9f.md">45 ARM新宠：苹果的M1芯片因何而快？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 AArch64体系：ARM最新编程架构模型剖析.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/46%20AArch64%e4%bd%93%e7%b3%bb%ef%bc%9aARM%e6%9c%80%e6%96%b0%e7%bc%96%e7%a8%8b%e6%9e%b6%e6%9e%84%e6%a8%a1%e5%9e%8b%e5%89%96%e6%9e%90.md">46 AArch64体系：ARM最新编程架构模型剖析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="LMOS来信：第二季课程带你“手撕”计算机基础.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/LMOS%e6%9d%a5%e4%bf%a1%ef%bc%9a%e7%ac%ac%e4%ba%8c%e5%ad%a3%e8%af%be%e7%a8%8b%e5%b8%a6%e4%bd%a0%e2%80%9c%e6%89%8b%e6%92%95%e2%80%9d%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80.md">LMOS来信：第二季课程带你“手撕”计算机基础.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助场 以无法为有法，以无限为有限.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e5%a4%a7%e5%92%96%e5%8a%a9%e5%9c%ba%20%e4%bb%a5%e6%97%a0%e6%b3%95%e4%b8%ba%e6%9c%89%e6%b3%95%ef%bc%8c%e4%bb%a5%e6%97%a0%e9%99%90%e4%b8%ba%e6%9c%89%e9%99%90.md">大咖助场 以无法为有法，以无限为有限.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 yiyang：我的上机实验“爬坑指南”.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20yiyang%ef%bc%9a%e6%88%91%e7%9a%84%e4%b8%8a%e6%9c%ba%e5%ae%9e%e9%aa%8c%e2%80%9c%e7%88%ac%e5%9d%91%e6%8c%87%e5%8d%97%e2%80%9d.md">用户故事 yiyang：我的上机实验“爬坑指南”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 成为面向“知识库”的工程师.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%88%90%e4%b8%ba%e9%9d%a2%e5%90%91%e2%80%9c%e7%9f%a5%e8%af%86%e5%ba%93%e2%80%9d%e7%9a%84%e5%b7%a5%e7%a8%8b%e5%b8%88.md">用户故事 成为面向“知识库”的工程师.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 技术人如何做选择，路才越走越宽？.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%8a%80%e6%9c%af%e4%ba%ba%e5%a6%82%e4%bd%95%e5%81%9a%e9%80%89%e6%8b%a9%ef%bc%8c%e8%b7%af%e6%89%8d%e8%b6%8a%e8%b5%b0%e8%b6%8a%e5%ae%bd%ef%bc%9f.md">用户故事 技术人如何做选择，路才越走越宽？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 操作系统发烧友：看不懂？因为你没动手.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%8f%91%e7%83%a7%e5%8f%8b%ef%bc%9a%e7%9c%8b%e4%b8%8d%e6%87%82%ef%bc%9f%e5%9b%a0%e4%b8%ba%e4%bd%a0%e6%b2%a1%e5%8a%a8%e6%89%8b.md">用户故事 操作系统发烧友：看不懂？因为你没动手.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 用好动态调试，助力课程学习.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e7%94%a8%e5%a5%bd%e5%8a%a8%e6%80%81%e8%b0%83%e8%af%95%ef%bc%8c%e5%8a%a9%e5%8a%9b%e8%af%be%e7%a8%8b%e5%ad%a6%e4%b9%a0.md">用户故事 用好动态调试，助力课程学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 艾同学：路虽远，行则将至.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e8%89%be%e5%90%8c%e5%ad%a6%ef%bc%9a%e8%b7%af%e8%99%bd%e8%bf%9c%ef%bc%8c%e8%a1%8c%e5%88%99%e5%b0%86%e8%87%b3.md">用户故事 艾同学：路虽远，行则将至.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 生活可以一地鸡毛，但操作系统却是心中的光.md" href="/%e4%b8%93%e6%a0%8f/%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%ae%9e%e6%88%9845%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e7%94%9f%e6%b4%bb%e5%8f%af%e4%bb%a5%e4%b8%80%e5%9c%b0%e9%b8%a1%e6%af%9b%ef%bc%8c%e4%bd%86%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e5%8d%b4%e6%98%af%e5%bf%83%e4%b8%ad%e7%9a%84%e5%85%89.md">结束语 生活可以一地鸡毛，但操作系统却是心中的光.md</a>
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
                            <h1 id="title" data-id="04 震撼的Linux全景图：业界成熟的内核架构长什么样？" class="title">04 震撼的Linux全景图：业界成熟的内核架构长什么样？</h1>
                            <div><p>你好，我是LMOS。</p>

<p>什么？你想成为计算机黑客？</p>

<p>梦想坐在计算机前敲敲键盘，银行账号里的数字就会自己往上涨。拜托，估计明天你就该被警察逮捕了。真正的黑客是对计算机技术有近乎极致的追求，而不是干坏事。</p>

<p>下面我就带你认识这样一个计算机黑客，看看他是怎样创造出影响世界的Linux，然后进一步了解一下Linux的内部结构。</p>

<p>同时，我也会带你看看Windows NT和Darwin的内部结构，三者形成对比，你能更好地了解它们之间的差异和共同点，这对我们后面写操作系统会很有帮助。</p>

<h2 id="关于linus">关于Linus</h2>

<p>Linus Benedict Torvalds，这个名字很长，下面简称Linus，他1969年12月28日出生在芬兰的赫尔辛基市，并不是美国人。Linus在赫尔辛基大学学的就是计算机，妻子还是空手道高手，一个“码林高手”和一个“武林高手”真的是绝配啊。</p>

<p>Linus在小时候就对各种事情充满好奇，这点非常具有黑客精神，后来有了自己的计算机更是痴迷其中，开始自己控制计算机做一些事情，并深挖其背后的原理。就是这种黑客精神促使他后来写出了颠覆世界的软件——Linux，也因此登上了美国《时代》周刊。</p>

<p>你是否对很多垃圾软件感到愤慨，但自己又无法改变。Linus就不一样，他为了方便访问大学服务器中的资源 ，而在自己的机器上写了一个文件系统和硬盘驱动，这样就可以把自己需要的资源下载到自己的机器中。</p>

<p>再后来，这成为了Linux的第一个版本。看看，牛人之所以为牛人就是敢于对现有的规则说不，并勇于改变。</p>

<p>如果仅仅如此，那么也不会有后来的Linux内核。Linus随后做了一个重要决定，他把这款操作系统雏形开源，并加入到自由软件运动，以GPL协议授权，允许用户自由复制或者改动程序代码，但用户必须公开自己的修改并传播。</p>

<p>无疑，正是Linus的这一重要决定使得Linux和他自己名声大振。短短几年时间，就已经聚集了成千上万的狂热分子，大家不计得失的为Linux添砖加瓦，很多程序员更是对Linus像神明一样顶礼膜拜。</p>

<h2 id="linux内核">Linux内核</h2>

<p>好了回到正题，回到Linux。Linus也不是什么神明，现有的Linux，99.9%的代码都不是Linus所写，而且他的代码，也不一定比你我的代码写得更好。</p>

<p>Linux，全称GNU/Linux，是一套免费使用和自由传播的操作系统，支持类UNIX、POSIX标准接口，也支持多用户、多进程、多线程，可以在多CPU的机器上运行。由于互联网的发展，Linux吸引了来自全世界各地软件爱好者、科技公司的支持，它已经从大型机到服务器蔓延至个人电脑、嵌入式系统等领域。</p>

<p>Linux系统性能稳定且开源。在很多公司企业网络中被当作服务器来使用，这是Linux的一大亮点，也是它得以壮大的关键。</p>

<p>Linux的基本思想是一切都是文件：每个文件都有确定的用途，包括用户数据、命令、配置参数、硬件设备等对于操作系统内核而言，都被视为各种类型的文件。Linux支持多用户，各个用户对于自己的文件有自己特殊的权利，保证了各用户之间互不影响。多任务则是现代操作系统最重要的一个特点，Linux可以使多个程序同时并独立地运行。</p>

<p>Linux发展到今天，不是哪一个人能做到的，更不是一群计算机黑客能做到的，而是由很多世界级的顶尖科技公司联合开发，如IBM、甲骨文、红帽、英特尔、微软，它们开发Linux并向Linux社区提供补丁，使Linux工作在它们的服务器上，向客户出售业务服务。</p>

<p>Linux发展到今天其代码量近2000万行，可以用浩如烟海来形容，没人能在短时间内弄清楚。但是你也不用害怕，我们可以先看看Linux内部的全景图，从全局了解一下Linux的内部结构，如下图。</p>

<p><img src="assets/ab62a335c9804e3d911563673da02a3b.jpg" alt="" /></p>

<p>啊哈！是不是感觉壮观之后一阵头晕目眩，头晕目眩就对了，因为Linux太大了，别怕，下面我们来分解一下。但这里我要先解释一下，上图仍然不足于描述Linux的全部，只是展示了重要且显而易见的部分。</p>

<p>上图中大致分为<strong>五大重要组件</strong>，每个组件又分成许多模块从上到下贯穿各个层次，每个模块中有重要的函数和数据结构。具体每个模块的主要功能，我都给你列在了文稿里，你可以详细看看后面这张图。</p>

<p><img src="assets/3d4ffb2f6be647f3a6ed410530cf1050.jpg" alt="" /></p>

<p>不要着急，不要心慌，因为现在我们不需要搞清楚这些Linux模块的全部实现细节，只要在心里默念Linux的模块真多啊，大概有五大组件，有好几十个模块，每个模块主要完成什么功能就行了。</p>

<p>是不是松了口气，先定定神，然后我们就能发现Linux这么多模块挤在一起，之间的通信主要是函数调用，而且函数间的调用没有一定的层次关系，更加没有左右边界的限定。函数的调用路径是纵横交错的，从图中的线条可以得到印证。</p>

<p>继续深入思考你就会发现，这些纵横交错的路径上有一个函数出现了问题，就麻烦大了，它会波及到全部组件，导致整个系统崩溃。当然调试解决这个问题，也是相当困难的。同样，模块之间没有隔离，安全隐患也是巨大的。</p>

<p>当然，这种结构不是一无是处，它的性能极高，而性能是衡量操作系统的一个重要指标。这种结构就是传统的内核结构，也称为<strong>宏内核架构</strong>。</p>

<p>想要评判一个产品好不好，最直接的方法就是用相似的产品对比。你说Linux很好，但是什么为好呢？我说Linux很差，它又差在什么地方呢？</p>

<p>下面我们就拿出Windows和macOS进行对比，注意我们只是对比它们的内核架构。</p>

<h2 id="darwin-xnu内核">Darwin-XNU内核</h2>

<p>我们先来看看Darwin，Darwin是由苹果公司在2000年开发的一个开放源代码的操作系统。</p>

<p>一个经久不衰的公司，必然有自己的核心竞争力，也许是商业策略，也许是技术产品，又或是这两者的结合。而作为苹果公司各种产品和强大的应用生态系统的支撑者——Darwin，更是公司核心竞争力中的核心。</p>

<p>苹果公司有台式计算机、笔记本、平板、手机，台式计算机、笔记本使用了macOS操作系统，平板和手机则使用了iOS操作系统。Darwin作为macOS与iOS操作系统的核心，从技术实现角度说，它必然要支持PowerPC、x86、ARM架构的处理器。</p>

<p>Darwin 使用了一种微内核（Mach）和相应的固件来支持不同的处理器平台，并提供操作系统原始的基础服务，上层的功能性系统服务和工具则是整合了BSD系统所提供的。苹果公司还为其开发了大量的库、框架和服务，不过它们都工作在用户态且闭源。</p>

<p>下面我们先从整体看一下Darwin的架构。</p>

<p><img src="assets/a1470e3434da4898ac41d8fa174cf7e7.jpg" alt="" /></p>

<p>什么？两套内核？惊不惊喜？由于我们是研究Darwin内核，所以上图中我们只需要关注内核-用户转换层以下的部分即可。显然它有两个内核层——<strong>Mach层与BSD层</strong>。</p>

<p>Mach内核是卡耐基梅隆大学开发的经典微内核，意在提供最基本的操作系统服务，从而达到高性能、安全、可扩展的目的，而BSD则是伯克利大学开发的类UNIX操作系统，提供一整套操作系统服务。</p>

<p>那为什么两套内核会同时存在呢？</p>

<p>MAC OS X（2011年之前的称呼）的发展经过了不同时期，随着时代的进步，产品功能需求增加，单纯的Mach之上实现出现了性能瓶颈，但是为了兼容之前为Mach开发的应用和设备驱动，就保留了Mach内核，同时加入了BSD内核。</p>

<p>Mach内核仍然提供十分简单的进程、线程、IPC通信、虚拟内存设备驱动相关的功能服务，BSD则提供强大的安全特性，完善的网络服务，各种文件系统的支持，同时对Mach的进程、线程、IPC、虚拟内核组件进行细化、扩展延伸。</p>

<p>那么应用如何使用Darwin系统的服务呢？应用会通过用户层的框架和库来请求Darwin系统的服务，即调用Darwin系统API。</p>

<p>在调用Darwin系统API时，会传入一个API号码，用这个号码去索引Mach陷入中断服务表中的函数。此时，API号码如果小于0，则表明请求的是Mach内核的服务，API号码如果大于0，则表明请求的是BSD内核的服务，它提供一整套标准的POSIX接口。</p>

<p>就这样，Mach和BSD就同时存在了。</p>

<p>Mach中还有一个重要的组件Libkern，它是一个库，提供了很多底层的操作函数，同时支持C++运行环境。</p>

<p>依赖这个库的还有IOKit，IOKit管理所有的设备驱动和内核功能扩展模块。驱动程序开发人员则可以使用C++面向对象的方式开发驱动，这个方式很优雅，你完全可以找一个成熟的驱动程序作为父类继承它，要特别实现某个功能就重载其中的函数，也可以同时继承其它驱动程序，这大大节省了内存，也大大降低了出现BUG的可能。</p>

<p>如果你要详细了解Darwin内核的话，可以自行阅读<a href="https://github.com/apple/darwin-xnu" target="_blank">相应的代码</a>。而在这里，你只要从全局认识一下它的结构就行了。</p>

<h2 id="windows-nt内核">Windows NT内核</h2>

<p>接下来我们再看下 NT 内核。现代Windows的内核就是NT，我们不妨先看看NT的历史。</p>

<p>如果你是90后，大概没有接触过MS-DOS，它的交互方式是你在键盘上输入相应的功能命令，它完成相应的功能后给用户返回相应的操作信息，没有图形界面。</p>

<p>在MS-DOS内核的实现上，也没有应用现代硬件的保护机制，这导致后来微软基于它开发的图形界面的操作系统，如Windows 3.1、Windows95/98/ME，极其不稳定，且容易死机。</p>

<p>加上类UNIX操作系统在互联网领域大行其道，所以微软急需一款全新的操作系统来与之竞争。所以，Windows NT诞生了。</p>

<p>Windows NT是微软于1993年推出的面向工作站、网络服务器和大型计算机的网络操作系统，也可做PC操作系统。它是一款全新从零开始开发的新操作系统，并应用了现代硬件的所有特性，“NT”所指的便是“新技术”（New Technology）。</p>

<p>而普通用户第一次接触基于NT内核的Windows是Windows 2000，一开始用户其实是不愿意接受的，因为Windows 2000对用户的硬件和应用存在兼容性问题。</p>

<p>随着硬件厂商和应用厂商对程序的升级，这个兼容性问题被缓解了，加之Windows 2000的高性能、高稳定性、高安全性，用户很快便接受了这个操作系统。这可以从Windows 2000的迭代者Windows XP的巨大成功，得到验证。</p>

<p>现在，NT内核在设计上层次非常清晰明了，各组件之间界限耦合程度很低。下面我们就来看看NT内核架构图，了解一下NT内核是如何“庄严宏伟”。如下图：</p>

<p><img src="assets/f8ef132c6e8c407193ad22550db82f7e.jpg" alt="" /></p>

<p>这样看NT内核架构，是不是就清晰了很多？但这并不是我画图画得清晰，事实上的NT确实如此。</p>

<p>这里我要提示一下，上图中我们只关注内核模式下的东西，也就是传统意义上的内核。</p>

<p>当然微软自己在HAL层上是定义了一个小内核，小内核之下是硬件抽象层HAL，这个HAL存在的好处是：不同的硬件平台只要提供对应的HAL就可以移植系统了。小内核之上是各种内核组件，微软称之为内核执行体，它们完成进程、内存、配置、I/O文件缓存、电源与即插即用、安全等相关的服务。</p>

<p>每个执行体互相独立，只对外提供相应的接口，其它执行体要通过内核模式可调用接口和其它执行体通信或者请求其完成相应的功能服务。所有的设备驱动和文件系统都由I/O管理器统一管理，驱动程序可以堆叠形成I/O驱动栈，功能请求被封装成I/O包，在栈中一层层流动处理。Windows引以为傲的图形子系统也在内核中。</p>

<p>显而易见，NT内核中各层次分明，各个执行体互相独立，这种“高内聚、低偶合”的特性，正是检验一个软件工程是否优秀的重要标准。而这些你都可以通过微软公开的WRK代码得到佐证，如果你觉得WRK代码量太少，也可以看一看<a href="https://reactos.org/" target="_blank">REACT OS</a>这个号称“开源版”的NT。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>到这里，我们了解了Linux、Darwin-XNU和Windows的发展历史，也清楚了它们内部的组件和结构，并对它们的架构进行了对比，对比后我们发现：<strong>Linux性能良好，结构异常复杂，不利于问题的排查和功能的扩展，而Darwin-XNU和Windows结构良好，层面分明，利于功能扩展，不容易产生问题且性能稳定。</strong></p>

<p>下面我们来回顾下这节课的重点。</p>

<p>首先，我们从一名计算机黑客切入，简单介绍了一下Linus，他由于沉迷于技术，对不好的规则敢于挑战而写出了Linux雏形，并且利用了GNU开源软件的精神推动了Linux后来的发展，这样的精神很值得我们学习。</p>

<p>然后我们探讨了Linux内核架构，大致搞清楚了Linux内核中的各种组件，它们是系统、进程、内存、储存、网络。其中，每个组件都是从接口到硬件经过了几个层次，组件与组件之间的层次互联调用。这些组件组合在一起，其调用关系形成了一个巨大的网状结构。因此，Linux也成了宏内核的代表。</p>

<p>为了有所对比，我们研究了苹果的Darwin-XNU内核结构，发现其分层更细，固件层、Mach层屏蔽了硬件平台的细节，向上层提供了最基础的服务。在Mach层之上的BSD层提供了更完善的服务，它们是进程与线程、IPC通信、虚拟内存、安全、网络协议栈以及文件系统。通过Mach中断嵌入表，可以让应用自己决定使用Mach层服务还是使用BSD层的服务，因此Darwin-XNU拥有了两套内核，Darwin-XNU内核层也成为了多内核架构的代表。</p>

<p>最后，我们研究了迄今为止，最成功的商业操作系统——Windows，它的内核是NT，其结构清晰明了，各组件完全遵循了软件工程<strong>高内聚、低偶合</strong>的设计标准。最下层是HAL（硬件抽象），HAL层是为了适配各种不同的硬件平台；在HAL层之上就是微软定义的小内核，你可以理解成是NT内核的内核；在这个小内核之上就是各种执行体了，这些执行体提供了操作系统的进程、虚拟内存、文件数据缓存、安全、对象管理、配置等服务，还有Windows的技术核心图形系统。</p>

<h2 id="思考题">思考题</h2>

<p>Windows NT内核属于哪种架构类型？</p>

<p>很期待在留言区看到你的分享，也欢迎你把这节课分享给身边的同事、朋友。</p>

<p>我是LMOS，让我们下节课见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#9df1f1f1a4a9acacadaaddfaf0fcf4f1b3fef2f0" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f15058c4e254596',t:'MTczNDA4MzEzMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
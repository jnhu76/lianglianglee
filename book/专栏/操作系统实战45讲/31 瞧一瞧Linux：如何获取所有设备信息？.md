<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=31&#32;瞧一瞧Linux：如何获取所有设备信息？>
        <link rel="icon" href="/static/favicon.png">
        <title>31 瞧一瞧Linux：如何获取所有设备信息？ </title>
        
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
                            <h1 id="title" data-id="31 瞧一瞧Linux：如何获取所有设备信息？" class="title">31 瞧一瞧Linux：如何获取所有设备信息？</h1>
                            <div><p>你好，我是LMOS。</p>

<p>前面我们已经完成了Cosmos的驱动设备的建立，还写好了一个真实的设备驱动。</p>

<p>今天，我们就来看看Linux是如何管理设备的。我们将从Linux如何组织设备开始，然后研究设备驱动相关的数据结构，最后我们还是要一起写一个Linux设备驱动实例，这样才能真正理解它。</p>

<h2 id="感受一下linux下的设备信息">感受一下Linux下的设备信息</h2>

<p>Linux的设计哲学就是一切皆文件，各种设备在Linux系统下自然也是一个个文件。不过这个文件并不对应磁盘上的数据文件，而是对应着存在内存当中的设备文件。实际上，我们对设备文件进行操作，就等同于操作具体的设备。</p>

<p>既然我们了解万事万物，都是从最直观的感受开始的，想要理解Linux对设备的管理，自然也是同样的道理。那么Linux设备文件在哪个目录下呢？其实现在我们在/sys/bus目录下，就可以查看所有的设备了。</p>

<p>Linux用BUS（总线）组织设备和驱动，我们在/sys/bus目录下输入tree命令，就可以看到所有总线下的所有设备了，如下图所示。</p>

<p><img src="assets/511510c8f4744ccf8f5e5c5695c1dce4.jpg" alt="" /></p>

<p>上图中，显示了部分Linux设备文件，有些设备文件是链接到其它目录下文件，这不是重点，重点是你要在心中有这个目录层次结构，即<strong>总线目录下有设备目录，设备目录下是设备文件</strong>。</p>

<h2 id="数据结构">数据结构</h2>

<p>我们接着刚才的图往下说，我们能感觉到Linux的驱动模型至少有三个核心数据结构，分别是总线、设备和驱动，但是要像上图那样有层次化地组织它们，只有总线、设备、驱动这三个数据结构是不够的，还得有两个数据结构来组织它们，那就是kobject和kset，下面我们就去研究它们。</p>

<h3 id="kobject与kset">kobject与kset</h3>

<p>kobject和kset是构成/sys目录下的目录节点和文件节点的核心，也是层次化组织总线、设备、驱动的核心数据结构，kobject、kset数据结构都能表示一个目录或者文件节点。下面我们先来研究一下kobject数据结构，代码如下所示。</p>

<pre><code>struct kobject {
    const char      *name;           //名称，反映在sysfs中
    struct list_head    entry;       //挂入kset结构的链表
    struct kobject      *parent;     //指向父结构 
    struct kset     *kset;           //指向所属的kset
    struct kobj_type    *ktype;
    struct kernfs_node  *sd;         //指向sysfs文件系统目录项 
    struct kref     kref;            //引用计数器结构
    unsigned int state_initialized:1;//初始化状态
    unsigned int state_in_sysfs:1;   //是否在sysfs中
    unsigned int state_add_uevent_sent:1;
    unsigned int state_remove_uevent_sent:1;
    unsigned int uevent_suppress:1;
};
</code></pre>

<p>每一个 kobject，都对应着 /sys目录下（其实是sysfs文件系统挂载在/sys目录下） 的一个目录或者文件，目录或者文件的名字就是kobject结构中的name。</p>

<p>我们从kobject结构中可以看出，它挂载在kset下，并且指向了kset，那kset是什么呢？我们来分析分析，它是kobject结构的容器吗？</p>

<p>其实是也不是，因为kset结构中本身又包含一个kobject结构，所以它既是kobject的容器，同时本身还是一个kobject。kset结构代码如下所示。</p>

<pre><code>struct kset {
    struct list_head list; //挂载kobject结构的链表
    spinlock_t list_lock; //自旋锁
    struct kobject kobj;//自身包含一个kobject结构
    const struct kset_uevent_ops *uevent_ops;//暂时不关注
} __randomize_layout;
</code></pre>

<p>看到这里你应该知道了，kset不仅仅自己是个kobject，还能挂载多个kobject，这说明kset是kobject的集合容器。在Linux内核中，至少有两个顶层kset，代码如下所示。</p>

<pre><code>struct kset *devices_kset;//管理所有设备
static struct kset *bus_kset;//管理所有总线
static struct kset *system_kset;
int __init devices_init(void)
{
    devices_kset = kset_create_and_add(&quot;devices&quot;, &amp;device_uevent_ops, NULL);//建立设备kset
    return 0;
}
int __init buses_init(void)
{
    bus_kset = kset_create_and_add(&quot;bus&quot;, &amp;bus_uevent_ops, NULL);//建立总线kset
    if (!bus_kset)
        return -ENOMEM;
    system_kset = kset_create_and_add(&quot;system&quot;, NULL, &amp;devices_kset-&gt;kobj);//在设备kset之下建立system的kset
    if (!system_kset)
        return -ENOMEM;
    return 0;
}
</code></pre>

<p>我知道，你可能很难想象许多个kset和kobject在逻辑上形成的层次结构，所以我为你画了一幅图，你可以结合这张示意图理解这个结构。</p>

<p><img src="assets/078dd6ab7cee413dab1fc0b1267f09a6.jpg" alt="" /></p>

<p>上图中展示了一个类似文件目录的结构，这正是kset与kobject设计的目标之一。kset与kobject结构只是基础数据结构，但是仅仅只有它的话，也就只能实现这个层次结构，其它的什么也不能干，根据我们以往的经验可以猜出，kset与kobject结构肯定是嵌入到更高级的数据结构之中使用，下面我们继续探索。</p>

<h3 id="总线">总线</h3>

<p>kset、kobject结构只是开胃菜，这个基础了解了，我们还要回到研究Linux设备与驱动的正题上。我们之前说过了，Linux用总线组织设备和驱动，由此可见总线是Linux设备的基础，它可以表示CPU与设备的连接，那么总线的数据结构是什么样呢？我们一起来看看。</p>

<p>Linux把总线抽象成bus_type结构，代码如下所示。</p>

<pre><code>struct bus_type {
    const char      *name;//总线名称
    const char      *dev_name;//用于列举设备，如（&quot;foo%u&quot;, dev-&gt;id）
    struct device       *dev_root;//父设备
    const struct attribute_group **bus_groups;//总线的默认属性
    const struct attribute_group **dev_groups;//总线上设备的默认属性
    const struct attribute_group **drv_groups;//总线上驱动的默认属性
    //每当有新的设备或驱动程序被添加到这个总线上时调用
    int (*match)(struct device *dev, struct device_driver *drv);
    //当一个设备被添加、移除或其他一些事情时被调用产生uevent来添加环境变量。
    int (*uevent)(struct device *dev, struct kobj_uevent_env *env);
    //当一个新的设备或驱动程序添加到这个总线时被调用，并回调特定驱动程序探查函数，以初始化匹配的设备
    int (*probe)(struct device *dev);
    //将设备状态同步到软件状态时调用
    void (*sync_state)(struct device *dev);
    //当一个设备从这个总线上删除时被调用
    int (*remove)(struct device *dev);
    //当系统关闭时被调用
    void (*shutdown)(struct device *dev);
    //调用以使设备重新上线（在下线后）
    int (*online)(struct device *dev);
    //调用以使设备离线，以便热移除。可能会失败。
    int (*offline)(struct device *dev);
    //当这个总线上的设备想进入睡眠模式时调用
    int (*suspend)(struct device *dev, pm_message_t state);
    //调用以使该总线上的一个设备脱离睡眠模式
    int (*resume)(struct device *dev);
    //调用以找出该总线上的一个设备支持多少个虚拟设备功能
    int (*num_vf)(struct device *dev);
    //调用以在该总线上的设备配置DMA
    int (*dma_configure)(struct device *dev);
    //该总线的电源管理操作，回调特定的设备驱动的pm-ops
    const struct dev_pm_ops *pm;
    //此总线的IOMMU具体操作，用于将IOMMU驱动程序实现到总线上
    const struct iommu_ops *iommu_ops;
    //驱动核心的私有数据，只有驱动核心能够接触这个
    struct subsys_private *p;
    struct lock_class_key lock_key;
    //当探测或移除该总线上的一个设备时，设备驱动核心应该锁定该设备
    bool need_parent_lock;
};
</code></pre>

<p>可以看出，上面代码的bus_type结构中，包括总线名字、总线属性，还有操作该总线下所有设备通用操作函数的指针，其各个函数的功能我在代码注释中已经写清楚了。</p>

<p>从这一点可以发现，<strong>总线不仅仅是组织设备和驱动的容器，还是同类设备的共有功能的抽象层。</strong>下面我们来看看subsys_private，它是总线的驱动核心的私有数据，其中有我们想知道的秘密，代码如下所示。</p>

<pre><code>//通过kobject找到对应的subsys_private
#define to_subsys_private(obj) container_of(obj, struct subsys_private, subsys.kobj)
struct subsys_private {
    struct kset subsys;//定义这个子系统结构的kset
    struct kset *devices_kset;//该总线的&quot;设备&quot;目录，包含所有的设备
    struct list_head interfaces;//总线相关接口的列表
    struct mutex mutex;//保护设备，和接口列表
    struct kset *drivers_kset;//该总线的&quot;驱动&quot;目录，包含所有的驱动
    struct klist klist_devices;//挂载总线上所有设备的可迭代链表
    struct klist klist_drivers;//挂载总线上所有驱动的可迭代链表
    struct blocking_notifier_head bus_notifier;
    unsigned int drivers_autoprobe:1;
    struct bus_type *bus;   //指向所属总线
    struct kset glue_dirs;
    struct class *class;//指向这个结构所关联类结构的指针
};
</code></pre>

<p>看到这里，你应该明白kset的作用了，我们通过bus_kset可以找到所有的kset，通过kset又能找到subsys_private，再通过subsys_private就可以找到总线了，也可以找到该总线上所有的设备与驱动。</p>

<h3 id="设备">设备</h3>

<p>虽然Linux抽象出了总线结构，但是Linux还需要表示一个设备，下面我们来探索Linux是如何表示一个设备的。</p>

<p>其实，在Linux系统中设备也是一个数据结构，里面包含了一个设备的所有信息。代码如下所示。</p>

<pre><code>struct device {
    struct kobject kobj;
    struct device       *parent;//指向父设备
    struct device_private   *p;//设备的私有数据
    const char      *init_name; //设备初始化名字
    const struct device_type *type;//设备类型
    struct bus_type *bus;  //指向设备所属总线
    struct device_driver *driver;//指向设备的驱动
    void        *platform_data;//设备平台数据
    void        *driver_data;//设备驱动的私有数据
    struct dev_links_info   links;//设备供应商链接
    struct dev_pm_info  power;//用于设备的电源管理
    struct dev_pm_domain    *pm_domain;//提供在系统暂停时执行调用
#ifdef CONFIG_GENERIC_MSI_IRQ
    struct list_head    msi_list;//主机的MSI描述符链表
#endif
    struct dev_archdata archdata;
    struct device_node  *of_node; //用访问设备树节点
    struct fwnode_handle    *fwnode; //设备固件节点
    dev_t           devt;   //用于创建sysfs &quot;dev&quot;
    u32         id; //设备实例id
    spinlock_t      devres_lock;//设备资源链表锁
    struct list_head    devres_head;//设备资源链表
    struct class        *class;//设备的类
    const struct attribute_group **groups;  //可选的属性组
    void    (*release)(struct device *dev);//在所有引用结束后释放设备
    struct iommu_group  *iommu_group;//该设备属于的IOMMU组
    struct dev_iommu    *iommu;//每个设备的通用IOMMU运行时数据
};
</code></pre>

<p>device结构很大，这里删除了我们不需要关心的内容。另外，我们看到device结构中同样包含了kobject结构，这使得设备可以加入kset和kobject组建的层次结构中。device结构中有总线和驱动指针，这能帮助设备找到自己的驱动程序和总线。</p>

<h3 id="驱动">驱动</h3>

<p>有了设备结构，还需要有设备对应的驱动，Linux是如何表示一个驱动的呢？同样也是一个数据结构，其中包含了驱动程序的相关信息。其实在device结构中我们就看到了，就是device_driver结构，代码如下。</p>

<pre><code>struct device_driver {
    const char      *name;//驱动名称
    struct bus_type     *bus;//指向总线
    struct module       *owner;//模块持有者
    const char      *mod_name;//用于内置模块
    bool suppress_bind_attrs;//禁用通过sysfs的绑定/解绑
    enum probe_type probe_type;//要使用的探查类型（同步或异步）
    const struct of_device_id   *of_match_table;//开放固件表
    const struct acpi_device_id *acpi_match_table;//ACPI匹配表
    //被调用来查询一个特定设备的存在
    int (*probe) (struct device *dev);
    //将设备状态同步到软件状态时调用
    void (*sync_state)(struct device *dev);
    //当设备被从系统中移除时被调用，以便解除设备与该驱动的绑定
    int (*remove) (struct device *dev);
    //关机时调用，使设备停止
    void (*shutdown) (struct device *dev);
    //调用以使设备进入睡眠模式，通常是进入一个低功率状态
    int (*suspend) (struct device *dev, pm_message_t state);
    //调用以使设备从睡眠模式中恢复
    int (*resume) (struct device *dev);
    //默认属性
    const struct attribute_group **groups;
    //绑定设备的属性
    const struct attribute_group **dev_groups;
    //设备电源操作
    const struct dev_pm_ops *pm;
    //当sysfs目录被写入时被调用
    void (*coredump) (struct device *dev);
    //驱动程序私有数据
    struct driver_private *p;
};
struct driver_private {
    struct kobject kobj;
    struct klist klist_devices;//驱动管理的所有设备的链表
    struct klist_node knode_bus;//加入bus链表的节点
    struct module_kobject *mkobj;//指向用kobject管理模块节点
    struct device_driver *driver;//指向驱动本身
};
</code></pre>

<p>在device_driver结构中，包含了驱动程序的名字、驱动程序所在模块、设备探查和电源相关的回调函数的指针。在driver_private结构中同样包含了kobject结构，用于组织所有的驱动，还指向了驱动本身，你发现没有，bus_type中的subsys_private结构的机制如出一辙。</p>

<h3 id="文件操作函数">文件操作函数</h3>

<p>前面我们学习的都是Linux驱动程序的核心数据结构，我们很少用到，只是为了让你了解最基础的原理。</p>

<p>其实，在Linux系统中提供了更为高级的封装，Linux将设备分成几类分别是：字符设备、块设备、网络设备以及杂项设备。具体情况你可以参考我后面梳理的图表。</p>

<p><img src="assets/a381cb72daf34a048cfe59ba3545e6aa.jpg" alt="" /></p>

<p>这些类型的设备的数据结构，都会直接或者间接包含基础的device结构，我们以杂项设备为例子研究一下，Linux用miscdevice结构表示一个杂项设备，代码如下。</p>

<pre><code>struct miscdevice  {
    int minor;//设备号
    const char *name;//设备名称
    const struct file_operations *fops;//文件操作函数结构
    struct list_head list;//链表
    struct device *parent;//指向父设备的device结构
    struct device *this_device;//指向本设备的device结构
    const struct attribute_group **groups;
    const char *nodename;//节点名字
    umode_t mode;//访问权限
};
</code></pre>

<p>miscdevice结构就是一个杂项设备，它一般在驱动程序代码文件中静态定义。我们清楚地看见有个this_device指针，它指向下层的、属于这个杂项设备的device结构。</p>

<p>但是这里重点是<strong>file_operations结构</strong>，设备一经注册，就会在sys相关的目录下建立设备对应的文件结点，对这个文件结点打开、读写等操作，最终会调用到驱动程序对应的函数，而对应的函数指针就保存在file_operations结构中，我们现在来看看这个结构。</p>

<pre><code>struct file_operations {
    struct module *owner;//所在的模块
    loff_t (*llseek) (struct file *, loff_t, int);//调整读写偏移
    ssize_t (*read) (struct file *, char __user *, size_t, loff_t *);//读
    ssize_t (*write) (struct file *, const char __user *, size_t, loff_t *);//写
    int (*mmap) (struct file *, struct vm_area_struct *);//映射
    int (*open) (struct inode *, struct file *);//打开
    int (*flush) (struct file *, fl_owner_t id);//刷新
    int (*release) (struct inode *, struct file *);//关闭
} __randomize_layout;
</code></pre>

<p>file_operations结构中的函数指针有31个，我删除了我们不熟悉的函数指针，我们了解原理，不需要搞清楚所有函数指针的功能。</p>

<p>那么，Linux如何调用到这个file_operations结构中的函数呢？我以打开操作为例给你讲讲，Linux的打开系统调用接口会调用filp_open函数，filp_open函数的调用路径如下所示。</p>

<pre><code>//filp_open
//file_open_name
//do_filp_open
//path_openat
static int do_o_path(struct nameidata *nd, unsigned flags, struct file *file)
{
    struct path path;
    int error = path_lookupat(nd, flags, &amp;path);//解析文件路径得到文件inode节点
    if (!error) {
        audit_inode(nd-&gt;name, path.dentry, 0);
        error = vfs_open(&amp;path, file);//vfs层打开文件接口
        path_put(&amp;path);
    }
    return error;
}
int vfs_open(const struct path *path, struct file *file)
{
    file-&gt;f_path = *path;
    return do_dentry_open(file, d_backing_inode(path-&gt;dentry), NULL);
}
static int do_dentry_open(struct file *f, struct inode *inode,int (*open)(struct inode *, struct file *))
{
    //略过我们不想看的代码
    f-&gt;f_op = fops_get(inode-&gt;i_fop);//获取文件节点的file_operations
    if (!open)//如果open为空则调用file_operations结构中的open函数
        open = f-&gt;f_op-&gt;open;
    if (open) {
        error = open(inode, f);
    }
    //略过我们不想看的代码
    return 0;
}
</code></pre>

<p>看到这里，我们就知道了，file_operations结构的地址存在一个文件的inode结构中。在Linux系统中，都是用inode结构表示一个文件，不管它是数据文件还是设备文件。</p>

<p>到这里，我们已经清楚了文件操作函数以及它的调用流程。</p>

<h2 id="驱动程序实例">驱动程序实例</h2>

<p>我们想要真正理解Linux设备驱动，最好的方案就是写一个真实的驱动程序实例。下面我们一起应用前面的基础，结合Linux提供的驱动程序开发接口，一起实现一个真实驱动程序。</p>

<p>这个驱动程序的主要工作，就是获取所有总线和其下所有设备的名字。为此我们需要先了解驱动程序的整体框架，接着建立我们总线和设备，然后实现驱动程序的打开、关闭，读写操作函数，最后我们写个应用程序，来测试我们的驱动程序。</p>

<h3 id="驱动程序框架">驱动程序框架</h3>

<p>Linux内核的驱动程序是在一个可加载的内核模块中实现，可加载的内核模块只需要两个函数和模块信息就行，但是我们要在模块中实现总线和设备驱动，所以需要更多的函数和数据结构，它们的代码如下。</p>

<pre><code>#define DEV_NAME  &quot;devicesinfo&quot;
#define BUS_DEV_NAME  &quot;devicesinfobus&quot;

static int misc_find_match(struct device *dev, void *data)
{
    printk(KERN_EMERG &quot;device name is:%s\n&quot;, dev-&gt;kobj.name);
    return 0;
}
//对应于设备文件的读操作函数
static ssize_t misc_read (struct file *pfile, char __user *buff, size_t size, loff_t *off)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);
    return 0;
}
//对应于设备文件的写操作函数
static ssize_t misc_write(struct file *pfile, const char __user *buff, size_t size, loff_t *off)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);    
    return 0;
}
//对应于设备文件的打开操作函数
static int  misc_open(struct inode *pinode, struct file *pfile)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);
    return 0;
} 
//对应于设备文件的关闭操作函数
static int misc_release(struct inode *pinode, struct file *pfile)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);
    return 0;
}

static int devicesinfo_bus_match(struct device *dev, struct device_driver *driver)
{
        return !strncmp(dev-&gt;kobj.name, driver-&gt;name, strlen(driver-&gt;name));
}
//对应于设备文件的操作函数结构
static const  struct file_operations misc_fops = {
    .read     = misc_read,
    .write    = misc_write,
    .release  = misc_release,
    .open     = misc_open,
};
//misc设备的结构
static struct miscdevice  misc_dev =  {
    .fops  =  &amp;misc_fops,         //设备文件操作方法
    .minor =  255,                //次设备号
    .name  =  DEV_NAME,           //设备名/dev/下的设备节点名
};
//总线结构
struct bus_type devicesinfo_bus = {
        .name = BUS_DEV_NAME, //总线名字
        .match = devicesinfo_bus_match, //总线match函数指针
};
//内核模块入口函数
static int __init miscdrv_init(void)
{
    printk(KERN_EMERG &quot;INIT misc\n&quot;)；
    return 0;
}
//内核模块退出函数
static void  __exit miscdrv_exit(void)
{
    printk(KERN_EMERG &quot;EXIT,misc\n&quot;);
}
module_init(miscdrv_init);//申明内核模块入口函数
module_exit(miscdrv_exit);//申明内核模块退出函数
MODULE_LICENSE(&quot;GPL&quot;);//模块许可
MODULE_AUTHOR(&quot;LMOS&quot;);//模块开发者
</code></pre>

<p>一个最简单的驱动程序框架的内核模块就写好了，该有的函数和数据结构都有了，那些数据结构都是静态定义的，它们的内部字段我们在前面也已经了解了。这个模块一旦加载就会执行miscdrv_init函数，卸载时就会执行miscdrv_exit函数。</p>

<h3 id="建立设备">建立设备</h3>

<p>Linux系统也提供了很多专用接口函数，用来建立总线和设备。下面我们先来建立一个总线，然后在总线下建立一个设备。</p>

<p>首先来说说建立一个总线，Linux系统提供了一个bus_register函数向内核注册一个总线，相当于建立了一个总线，我们需要在miscdrv_init函数中调用它，代码如下所示。</p>

<pre><code>static int __init miscdrv_init(void)
{
    printk(KERN_EMERG &quot;INIT misc\n&quot;);
    busok = bus_register(&amp;devicesinfo_bus);//注册总线
    return 0;
}
</code></pre>

<p>bus_register函数会在系统中注册一个总线，所需参数就是总线结构的地址(&amp;devicesinfo_bus)，返回非0表示注册失败。现在我们来看看，在bus_register函数中都做了些什么事情，代码如下所示。</p>

<pre><code>int bus_register(struct bus_type *bus)
{
    int retval;
    struct subsys_private *priv;
    //分配一个subsys_private结构
    priv = kzalloc(sizeof(struct subsys_private), GFP_KERNEL);
    //bus_type和subsys_private结构互相指向
    priv-&gt;bus = bus;
    bus-&gt;p = priv;
    //把总线的名称加入subsys_private的kobject中
    retval = kobject_set_name(&amp;priv-&gt;subsys.kobj, &quot;%s&quot;, bus-&gt;name);
    priv-&gt;subsys.kobj.kset = bus_kset;//指向bus_kset
    //把subsys_private中的kset注册到系统中
    retval = kset_register(&amp;priv-&gt;subsys);
    //建立总线的文件结构在sysfs中
    retval = bus_create_file(bus, &amp;bus_attr_uevent);
    //建立subsys_private中的devices和drivers的kset
    priv-&gt;devices_kset = kset_create_and_add(&quot;devices&quot;, NULL,
                         &amp;priv-&gt;subsys.kobj);
    priv-&gt;drivers_kset = kset_create_and_add(&quot;drivers&quot;, NULL,
                         &amp;priv-&gt;subsys.kobj);
    //建立subsys_private中的devices和drivers链表，用于属于总线的设备和驱动
    klist_init(&amp;priv-&gt;klist_devices, klist_devices_get, klist_devices_put);
    klist_init(&amp;priv-&gt;klist_drivers, NULL, NULL);
    return 0;
}
</code></pre>

<p>我删除了很多你不用关注的代码，看到这里，你应该知道总线是怎么通过subsys_private把设备和驱动关联起来的（通过bus_type和subsys_private结构互相指向），下面我们看看怎么建立设备。我们这里建立一个misc杂项设备。misc杂项设备需要定一个数据结构，然后调用misc杂项设备注册接口函数，代码如下。</p>

<pre><code>#define DEV_NAME  &quot;devicesinfo&quot;
static const  struct file_operations misc_fops = {
    .read     = misc_read,
    .write    = misc_write,
    .release  = misc_release,
    .open     = misc_open,
};
static struct miscdevice  misc_dev =  {
    .fops  =  &amp;misc_fops,         //设备文件操作方法
    .minor =  255,                //次设备号
    .name  =  DEV_NAME,           //设备名/dev/下的设备节点名
};
static int __init miscdrv_init(void)
{
    misc_register(&amp;misc_dev);//注册misc杂项设备
    printk(KERN_EMERG &quot;INIT misc busok\n&quot;);
    busok = bus_register(&amp;devicesinfo_bus);//注册总线
    return 0;
}
</code></pre>

<p>上面的代码中，静态定义了miscdevice结构的变量misc_dev，miscdevice结构我们在前面已经了解过了，最后调用misc_register函数注册了misc杂项设备。</p>

<p>misc_register函数到底做了什么，我们一起来看看，代码如下所示。</p>

<pre><code>int misc_register(struct miscdevice *misc)
{
    dev_t dev;
    int err = 0;
    bool is_dynamic = (misc-&gt;minor == MISC_DYNAMIC_MINOR);
    INIT_LIST_HEAD(&amp;misc-&gt;list);
    mutex_lock(&amp;misc_mtx);
    if (is_dynamic) {//minor次设备号如果等于255就自动分配次设备
        int i = find_first_zero_bit(misc_minors, DYNAMIC_MINORS);
        if (i &gt;= DYNAMIC_MINORS) {
            err = -EBUSY;
            goto out;
        }
        misc-&gt;minor = DYNAMIC_MINORS - i - 1;
        set_bit(i, misc_minors);
    } else {//否则检查次设备号是否已经被占有
        struct miscdevice *c;
        list_for_each_entry(c, &amp;misc_list, list) {
            if (c-&gt;minor == misc-&gt;minor) {
                err = -EBUSY;
                goto out;
            }
        }
    }
    dev = MKDEV(MISC_MAJOR, misc-&gt;minor);//合并主、次设备号
    //建立设备
    misc-&gt;this_device =
        device_create_with_groups(misc_class, misc-&gt;parent, dev,
                      misc, misc-&gt;groups, &quot;%s&quot;, misc-&gt;name);
    //把这个misc加入到全局misc_list链表
    list_add(&amp;misc-&gt;list, &amp;misc_list);
 out:
    mutex_unlock(&amp;misc_mtx);
    return err;
}
</code></pre>

<p>可以看出，misc_register函数只是负责分配设备号，以及把miscdev加入链表，真正的核心工作由device_create_with_groups函数来完成，代码如下所示。</p>

<pre><code>struct device *device_create_with_groups(struct class *class,
                     struct device *parent, dev_t devt,void *drvdata,const struct attribute_group **groups,const char *fmt, ...)
{
    va_list vargs;
    struct device *dev;
    va_start(vargs, fmt);
    dev = device_create_groups_vargs(class, parent, devt, drvdata, groups,fmt, vargs);
    va_end(vargs);
    return dev;
}
struct device *device_create_groups_vargs(struct class *class, struct device *parent, dev_t devt, void *drvdata,const struct attribute_group **groups,const char *fmt, va_list args)
{
    struct device *dev = NULL;
    int retval = -ENODEV;
    dev = kzalloc(sizeof(*dev), GFP_KERNEL);//分配设备结构的内存空间
    device_initialize(dev);//初始化设备结构
    dev-&gt;devt = devt;//设置设备号
    dev-&gt;class = class;//设置设备类
    dev-&gt;parent = parent;//设置设备的父设备
    dev-&gt;groups = groups;////设置设备属性
    dev-&gt;release = device_create_release;
    dev_set_drvdata(dev, drvdata);//设置miscdev的地址到设备结构中
    retval = kobject_set_name_vargs(&amp;dev-&gt;kobj, fmt, args);//把名称设置到设备的kobjext中去
    retval = device_add(dev);//把设备加入到系统中
    if (retval)
        goto error;
    return dev;//返回设备
error:
    put_device(dev);
    return ERR_PTR(retval);
}
</code></pre>

<p>到这里，misc设备的注册就搞清楚了，下面我们来测试一下看看结果，看看Linux系统是不是多了一个总线和设备。</p>

<p>你可以在本课程的代码目录中，执行make指令，就会产生一个miscdvrv.ko内核模块文件，我们把这个模块文件加载到Linux系统中就行了。</p>

<p>为了看到效果，我们还必须要做另一件事情。 在终端中用sudo cat /proc/kmsg 指令读取/proc/kmsg文件，该文件是内核prink函数输出信息的文件。指令如下所示。</p>

<pre><code>#第一步在终端中执行如下指令
sudo cat /proc/kmsg
#第二步在另一个终端中执行如下指令
make
sudo insmod miscdrv.ko
#不用这个模块了可以用以下指令卸载
sudo rmmod miscdrv.ko
</code></pre>

<p>insmod指令是加载一个内核模块，一旦加载成功就会执行miscdrv_init函数。如果不出意外，你在终端中会看到如下图所示的情况。</p>

<p><img src="assets/11d9f4f2d9374a51a863616a819b26ea.jpg" alt="" /></p>

<p>这说明我们设备已经建立了，你应该可以在/dev目录看到一个devicesinfo文件，同时你在/sys/bus/目录下也可以看到一个devicesinfobus文件。这就是我们建立的设备和总线的文件节点的名称。</p>

<h3 id="打开-关闭-读写函数">打开、关闭、读写函数</h3>

<p>建立了设备和总线，有了设备文件节点，应用程序就可以打开、关闭以及读写这个设备文件了。</p>

<p>虽然现在确实可以操作设备文件了，只不过还不能完成任何实际功能，因为我们只是写好了框架函数，所以我们下面就去写好并填充这些框架函数，代码如下所示。</p>

<pre><code>//打开
static int  misc_open(struct inode *pinode, struct file *pfile)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);//打印这个函数所在文件的行号和名称
    return 0;
}
//关闭 
static int misc_release(struct inode *pinode, struct file *pfile)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);//打印这个函数所在文件的行号和名称
    return 0;
}
//写
static ssize_t misc_write(struct file *pfile, const char __user *buff, size_t size, loff_t *off)
{
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);//打印这个函数所在文件的行号和名称    
    return 0;
}
</code></pre>

<p>以上三个函数，仍然没干什么实际工作，就是打印该函数所在文件的行号和名称，然后返回0就完事了。回到前面，我们的目的是要获取Linux中所有总线上的所有设备，所以在读函数中来实现是合理的。</p>

<p>具体实现的代码如下所示。</p>

<pre><code>#define to_subsys_private(obj) container_of(obj, struct subsys_private, subsys.kobj)//从kobject上获取subsys_private的地址
struct kset *ret_buskset(void)
{
    struct subsys_private *p;
    if(busok)
        return NULL;
    if(!devicesinfo_bus.p)
        return NULL;
    p = devicesinfo_bus.p;
    if(!p-&gt;subsys.kobj.kset)
        return NULL;
    //返回devicesinfo_bus总线上的kset，正是bus_kset
    return p-&gt;subsys.kobj.kset;
}
static int misc_find_match(struct device *dev, void *data)
{
    struct bus_type* b = (struct bus_type*)data;
    printk(KERN_EMERG &quot;%s----&gt;device name is:%s\n&quot;, b-&gt;name, dev-&gt;kobj.name);//打印总线名称和设备名称
    return 0;
}
static ssize_t misc_read (struct file *pfile, char __user *buff, size_t size, loff_t *off)
{
    struct kobject* kobj;
    struct kset* kset;
    struct subsys_private* p;
    kset = ret_buskset();//获取bus_kset的地址
    if(!kset)
        return 0;
    printk(KERN_EMERG &quot;line:%d,%s is call\n&quot;,__LINE__,__FUNCTION__);//打印这个函数所在文件的行号和名称
    //扫描所有总线的kobject
    list_for_each_entry(kobj, &amp;kset-&gt;list, entry)
    {
        p = to_subsys_private(kobj);
        printk(KERN_EMERG &quot;Bus name is:%s\n&quot;,p-&gt;bus-&gt;name);
        //遍历具体总线上的所有设备
        bus_for_each_dev(p-&gt;bus, NULL, p-&gt;bus, misc_find_match);
    }
    return 0;
}
</code></pre>

<p>正常情况下，我们是不能获取bus_kset地址的，它是所有总线的根，包含了所有总线的kobject，Linux为了保护bus_kset，并没有在bus_type结构中直接包含kobject，而是让总线指向一个subsys_private结构，在其中包含了kobject结构。</p>

<p>所以，我们要注册一个总线，这样就能拔出萝卜带出泥，得到bus_kset，根据它又能找到所有subsys_private结构中的kobject，接着找到subsys_private结构，反向查询到bus_type结构的地址。</p>

<p>然后调用Linux提供的bus_for_each_dev函数，就可以遍历一个总线上的所有设备，它每遍历到一个设备，就调用一个函数，这个函数是用参数的方式传给它的，在我们代码中就是misc_find_match函数。</p>

<p>在调用misc_find_match函数时，会把一个设备结构的地址和另一个指针作为参数传递进来。最后就能打印每个设备的名称了。</p>

<h3 id="测试驱动">测试驱动</h3>

<p>驱动程序已经写好，加载之后会自动建立设备文件，但是驱动程序不会主动工作，我们还需要写一个应用程序，对设备文件进行读写，才能测试驱动。我们这里这个驱动对打开、关闭、写操作没有什么实际的响应，但是只要一读就会打印所有设备的信息了。</p>

<p>下面我们来写好这个应用，代码如下所示。</p>

<pre><code>#include &lt;stdio.h&gt;
#include &lt;stdlib.h&gt;
#include &lt;string.h&gt;
#include &lt;unistd.h&gt;
#include &lt;sys/types.h&gt;
#include &lt;sys/stat.h&gt;
#include &lt;fcntl.h&gt;
#define DEV_NAME &quot;/dev/devicesinfo&quot;
int main(void)
{
    char buf[] = {0, 0, 0, 0};
    int fd;
    //打开设备文件
    fd = open(DEV_NAME, O_RDWR);
    if (fd &lt; 0) {
        printf(&quot;打开 :%s 失败!\n&quot;, DEV_NAME);
    }
    //写数据到内核空间
    write(fd, buf, 4);
    //从内核空间中读取数据
    read(fd, buf, 4);
    //关闭设备,也可以不调用，程序关闭时系统自动调用
    close(fd);
    return 0;
}
</code></pre>

<p>你可以这样操作：切换到本课程的代码目录make一下，然后加载miscdrv.ko模块，最后在终端中执行sudo ./app，就能在另一个已经执行了sudo cat /proc/kmsg的终端中，看到后面图片这样形式的数据。</p>

<p><img src="assets/033e9744384e4ad489f1931cef21474c.jpg" alt="" /></p>

<p>上图是我系统中总线名和设备名，你的计算机上可能略有差异，因为我们的计算机硬件可能不同，所以有差异是正常的，不必奇怪。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>尽管Linux驱动模型异常复杂，我们还是以最小的成本，领会了Linux驱动模型设计的要点，还动手写了个小小的驱动程序。现在我来为你梳理一下这节课的重点。</p>

<p>首先，我们通过查看sys目录下的文件层次结构，直观感受了一下Linux系统的总线、设备、驱动是什么情况。</p>

<p>然后，我们了解一些重要的数据结构，它们分别是总线、驱动、设备、文件操作函数结构，还有非常关键的<strong>kset和kobject</strong>，这两个结构一起组织了总线、设备、驱动，最终形成了类目录文件这样的层次结构。</p>

<p>最后，我们建立一个驱动程序实例，从驱动程序框架开始，我们了解如何建立一个总线和设备，编写了对应的文件操作函数，在读操作函数中实现扫描了所有总线上的所有设备，并打印总线名称和设备名称，还写了个应用程序进行了测试，检查有没有达到预期的功能。</p>

<p>如果你对Linux是怎么在总线上注册设备和驱动，又对驱动和设备怎么进行匹配感兴趣的话，也可以自己阅读Linux内核代码，其中有很多驱动实例，你可以研究和实验，动手和动脑相结合，我相信你一定可以搞清楚的。</p>

<h2 id="思考题">思考题</h2>

<p>为什么无论是我们加载miscdrv.ko内核模块，还是运行app测试，都要在前面加上sudo呢？</p>

<p>欢迎你在留言区记录你的学习收获，也欢迎你把这节课分享给你身边的小伙伴，一起拿下Linux设备驱动的内容。</p>

<p>我是LMOS，我们下节课见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#0d61616134393c3c3d3a4d6a606c6461236e6260" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1511632c7a4596',t:'MTczNDA4MzYxNS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=48&#32;&#32;DMA：为什么Kafka这么快？>
        <link rel="icon" href="/static/favicon.png">
        <title>48  DMA：为什么Kafka这么快？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  为什么你需要学习计算机组成原理？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e9%9c%80%e8%a6%81%e5%ad%a6%e4%b9%a0%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86%ef%bc%9f.md">00 开篇词  为什么你需要学习计算机组成原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  冯·诺依曼体系结构：计算机组成的金字塔.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/01%20%20%e5%86%af%c2%b7%e8%af%ba%e4%be%9d%e6%9b%bc%e4%bd%93%e7%b3%bb%e7%bb%93%e6%9e%84%ef%bc%9a%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e7%9a%84%e9%87%91%e5%ad%97%e5%a1%94.md">01  冯·诺依曼体系结构：计算机组成的金字塔.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  给你一张知识地图，计算机组成原理应该这么学.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/02%20%20%e7%bb%99%e4%bd%a0%e4%b8%80%e5%bc%a0%e7%9f%a5%e8%af%86%e5%9c%b0%e5%9b%be%ef%bc%8c%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86%e5%ba%94%e8%af%a5%e8%bf%99%e4%b9%88%e5%ad%a6.md">02  给你一张知识地图，计算机组成原理应该这么学.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  通过你的CPU主频，我们来谈谈“性能”究竟是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/03%20%20%e9%80%9a%e8%bf%87%e4%bd%a0%e7%9a%84CPU%e4%b8%bb%e9%a2%91%ef%bc%8c%e6%88%91%e4%bb%ac%e6%9d%a5%e8%b0%88%e8%b0%88%e2%80%9c%e6%80%a7%e8%83%bd%e2%80%9d%e7%a9%b6%e7%ab%9f%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">03  通过你的CPU主频，我们来谈谈“性能”究竟是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  穿越功耗墙，我们该从哪些方面提升“性能”？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/04%20%20%e7%a9%bf%e8%b6%8a%e5%8a%9f%e8%80%97%e5%a2%99%ef%bc%8c%e6%88%91%e4%bb%ac%e8%af%a5%e4%bb%8e%e5%93%aa%e4%ba%9b%e6%96%b9%e9%9d%a2%e6%8f%90%e5%8d%87%e2%80%9c%e6%80%a7%e8%83%bd%e2%80%9d%ef%bc%9f.md">04  穿越功耗墙，我们该从哪些方面提升“性能”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  计算机指令：让我们试试用纸带编程.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/05%20%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e6%8c%87%e4%bb%a4%ef%bc%9a%e8%ae%a9%e6%88%91%e4%bb%ac%e8%af%95%e8%af%95%e7%94%a8%e7%ba%b8%e5%b8%a6%e7%bc%96%e7%a8%8b.md">05  计算机指令：让我们试试用纸带编程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  指令跳转：原来if...else就是goto.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/06%20%20%e6%8c%87%e4%bb%a4%e8%b7%b3%e8%bd%ac%ef%bc%9a%e5%8e%9f%e6%9d%a5if...else%e5%b0%b1%e6%98%afgoto.md">06  指令跳转：原来if...else就是goto.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  函数调用：为什么会发生stack overflow？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/07%20%20%e5%87%bd%e6%95%b0%e8%b0%83%e7%94%a8%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bc%9a%e5%8f%91%e7%94%9fstack%20overflow%ef%bc%9f.md">07  函数调用：为什么会发生stack overflow？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  ELF和静态链接：为什么程序无法同时在Linux和Windows下运行？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/08%20%20ELF%e5%92%8c%e9%9d%99%e6%80%81%e9%93%be%e6%8e%a5%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e7%a8%8b%e5%ba%8f%e6%97%a0%e6%b3%95%e5%90%8c%e6%97%b6%e5%9c%a8Linux%e5%92%8cWindows%e4%b8%8b%e8%bf%90%e8%a1%8c%ef%bc%9f.md">08  ELF和静态链接：为什么程序无法同时在Linux和Windows下运行？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  程序装载：“640K内存”真的不够用么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/09%20%20%e7%a8%8b%e5%ba%8f%e8%a3%85%e8%bd%bd%ef%bc%9a%e2%80%9c640K%e5%86%85%e5%ad%98%e2%80%9d%e7%9c%9f%e7%9a%84%e4%b8%8d%e5%a4%9f%e7%94%a8%e4%b9%88%ef%bc%9f.md">09  程序装载：“640K内存”真的不够用么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  动态链接：程序内部的“共享单车”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/10%20%20%e5%8a%a8%e6%80%81%e9%93%be%e6%8e%a5%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%86%85%e9%83%a8%e7%9a%84%e2%80%9c%e5%85%b1%e4%ba%ab%e5%8d%95%e8%bd%a6%e2%80%9d.md">10  动态链接：程序内部的“共享单车”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  二进制编码：“手持两把锟斤拷，口中疾呼烫烫烫”？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/11%20%20%e4%ba%8c%e8%bf%9b%e5%88%b6%e7%bc%96%e7%a0%81%ef%bc%9a%e2%80%9c%e6%89%8b%e6%8c%81%e4%b8%a4%e6%8a%8a%e9%94%9f%e6%96%a4%e6%8b%b7%ef%bc%8c%e5%8f%a3%e4%b8%ad%e7%96%be%e5%91%bc%e7%83%ab%e7%83%ab%e7%83%ab%e2%80%9d%ef%bc%9f.md">11  二进制编码：“手持两把锟斤拷，口中疾呼烫烫烫”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  理解电路：从电报机到门电路，我们如何做到“千里传信”？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/12%20%20%e7%90%86%e8%a7%a3%e7%94%b5%e8%b7%af%ef%bc%9a%e4%bb%8e%e7%94%b5%e6%8a%a5%e6%9c%ba%e5%88%b0%e9%97%a8%e7%94%b5%e8%b7%af%ef%bc%8c%e6%88%91%e4%bb%ac%e5%a6%82%e4%bd%95%e5%81%9a%e5%88%b0%e2%80%9c%e5%8d%83%e9%87%8c%e4%bc%a0%e4%bf%a1%e2%80%9d%ef%bc%9f.md">12  理解电路：从电报机到门电路，我们如何做到“千里传信”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  加法器：如何像搭乐高一样搭电路（上）？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/13%20%20%e5%8a%a0%e6%b3%95%e5%99%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%83%8f%e6%90%ad%e4%b9%90%e9%ab%98%e4%b8%80%e6%a0%b7%e6%90%ad%e7%94%b5%e8%b7%af%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9f.md">13  加法器：如何像搭乐高一样搭电路（上）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  乘法器：如何像搭乐高一样搭电路（下）？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/14%20%20%e4%b9%98%e6%b3%95%e5%99%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%83%8f%e6%90%ad%e4%b9%90%e9%ab%98%e4%b8%80%e6%a0%b7%e6%90%ad%e7%94%b5%e8%b7%af%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9f.md">14  乘法器：如何像搭乐高一样搭电路（下）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  浮点数和定点数（上）：怎么用有限的Bit表示尽可能多的信息？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/15%20%20%e6%b5%ae%e7%82%b9%e6%95%b0%e5%92%8c%e5%ae%9a%e7%82%b9%e6%95%b0%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%80%8e%e4%b9%88%e7%94%a8%e6%9c%89%e9%99%90%e7%9a%84Bit%e8%a1%a8%e7%a4%ba%e5%b0%bd%e5%8f%af%e8%83%bd%e5%a4%9a%e7%9a%84%e4%bf%a1%e6%81%af%ef%bc%9f.md">15  浮点数和定点数（上）：怎么用有限的Bit表示尽可能多的信息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  浮点数和定点数（下）：深入理解浮点数到底有什么用？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/16%20%20%e6%b5%ae%e7%82%b9%e6%95%b0%e5%92%8c%e5%ae%9a%e7%82%b9%e6%95%b0%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e6%b5%ae%e7%82%b9%e6%95%b0%e5%88%b0%e5%ba%95%e6%9c%89%e4%bb%80%e4%b9%88%e7%94%a8%ef%bc%9f.md">16  浮点数和定点数（下）：深入理解浮点数到底有什么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  建立数据通路（上）：指令加运算=CPU.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/17%20%20%e5%bb%ba%e7%ab%8b%e6%95%b0%e6%8d%ae%e9%80%9a%e8%b7%af%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%8c%87%e4%bb%a4%e5%8a%a0%e8%bf%90%e7%ae%97=CPU.md">17  建立数据通路（上）：指令加运算=CPU.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  建立数据通路（中）：指令加运算=CPU.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/18%20%20%e5%bb%ba%e7%ab%8b%e6%95%b0%e6%8d%ae%e9%80%9a%e8%b7%af%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e6%8c%87%e4%bb%a4%e5%8a%a0%e8%bf%90%e7%ae%97=CPU.md">18  建立数据通路（中）：指令加运算=CPU.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  建立数据通路（下）：指令加运算=CPU.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/19%20%20%e5%bb%ba%e7%ab%8b%e6%95%b0%e6%8d%ae%e9%80%9a%e8%b7%af%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%8c%87%e4%bb%a4%e5%8a%a0%e8%bf%90%e7%ae%97=CPU.md">19  建立数据通路（下）：指令加运算=CPU.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  面向流水线的指令设计（上）：一心多用的现代CPU.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/20%20%20%e9%9d%a2%e5%90%91%e6%b5%81%e6%b0%b4%e7%ba%bf%e7%9a%84%e6%8c%87%e4%bb%a4%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%b8%80%e5%bf%83%e5%a4%9a%e7%94%a8%e7%9a%84%e7%8e%b0%e4%bb%a3CPU.md">20  面向流水线的指令设计（上）：一心多用的现代CPU.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  面向流水线的指令设计（下）：奔腾4是怎么失败的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/21%20%20%e9%9d%a2%e5%90%91%e6%b5%81%e6%b0%b4%e7%ba%bf%e7%9a%84%e6%8c%87%e4%bb%a4%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a5%94%e8%85%be4%e6%98%af%e6%80%8e%e4%b9%88%e5%a4%b1%e8%b4%a5%e7%9a%84%ef%bc%9f.md">21  面向流水线的指令设计（下）：奔腾4是怎么失败的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  冒险和预测（一）：hazard是“危”也是“机”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/22%20%20%e5%86%92%e9%99%a9%e5%92%8c%e9%a2%84%e6%b5%8b%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9ahazard%e6%98%af%e2%80%9c%e5%8d%b1%e2%80%9d%e4%b9%9f%e6%98%af%e2%80%9c%e6%9c%ba%e2%80%9d.md">22  冒险和预测（一）：hazard是“危”也是“机”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  冒险和预测（二）：流水线里的接力赛.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/23%20%20%e5%86%92%e9%99%a9%e5%92%8c%e9%a2%84%e6%b5%8b%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%b5%81%e6%b0%b4%e7%ba%bf%e9%87%8c%e7%9a%84%e6%8e%a5%e5%8a%9b%e8%b5%9b.md">23  冒险和预测（二）：流水线里的接力赛.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  冒险和预测（三）：CPU里的“线程池”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/24%20%20%e5%86%92%e9%99%a9%e5%92%8c%e9%a2%84%e6%b5%8b%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9aCPU%e9%87%8c%e7%9a%84%e2%80%9c%e7%ba%bf%e7%a8%8b%e6%b1%a0%e2%80%9d.md">24  冒险和预测（三）：CPU里的“线程池”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  冒险和预测（四）：今天下雨了，明天还会下雨么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/25%20%20%e5%86%92%e9%99%a9%e5%92%8c%e9%a2%84%e6%b5%8b%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e4%bb%8a%e5%a4%a9%e4%b8%8b%e9%9b%a8%e4%ba%86%ef%bc%8c%e6%98%8e%e5%a4%a9%e8%bf%98%e4%bc%9a%e4%b8%8b%e9%9b%a8%e4%b9%88%ef%bc%9f.md">25  冒险和预测（四）：今天下雨了，明天还会下雨么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  Superscalar和VLIW：如何让CPU的吞吐率超过1？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/26%20%20Superscalar%e5%92%8cVLIW%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9CPU%e7%9a%84%e5%90%9e%e5%90%90%e7%8e%87%e8%b6%85%e8%bf%871%ef%bc%9f.md">26  Superscalar和VLIW：如何让CPU的吞吐率超过1？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  SIMD：如何加速矩阵乘法？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/27%20%20SIMD%ef%bc%9a%e5%a6%82%e4%bd%95%e5%8a%a0%e9%80%9f%e7%9f%a9%e9%98%b5%e4%b9%98%e6%b3%95%ef%bc%9f.md">27  SIMD：如何加速矩阵乘法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  异常和中断：程序出错了怎么办？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/28%20%20%e5%bc%82%e5%b8%b8%e5%92%8c%e4%b8%ad%e6%96%ad%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%87%ba%e9%94%99%e4%ba%86%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">28  异常和中断：程序出错了怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  CISC和RISC：为什么手机芯片都是ARM？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/29%20%20CISC%e5%92%8cRISC%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e6%89%8b%e6%9c%ba%e8%8a%af%e7%89%87%e9%83%bd%e6%98%afARM%ef%bc%9f.md">29  CISC和RISC：为什么手机芯片都是ARM？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  GPU（上）：为什么玩游戏需要使用GPU？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/30%20%20GPU%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e7%8e%a9%e6%b8%b8%e6%88%8f%e9%9c%80%e8%a6%81%e4%bd%bf%e7%94%a8GPU%ef%bc%9f.md">30  GPU（上）：为什么玩游戏需要使用GPU？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  GPU（下）：为什么深度学习需要使用GPU？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/31%20%20GPU%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e9%9c%80%e8%a6%81%e4%bd%bf%e7%94%a8GPU%ef%bc%9f.md">31  GPU（下）：为什么深度学习需要使用GPU？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  FPGA、ASIC和TPU（上）：计算机体系结构的黄金时代.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/32%20%20FPGA%e3%80%81ASIC%e5%92%8cTPU%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e8%ae%a1%e7%ae%97%e6%9c%ba%e4%bd%93%e7%b3%bb%e7%bb%93%e6%9e%84%e7%9a%84%e9%bb%84%e9%87%91%e6%97%b6%e4%bb%a3.md">32  FPGA、ASIC和TPU（上）：计算机体系结构的黄金时代.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  解读TPU：设计和拆解一块ASIC芯片.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/33%20%20%e8%a7%a3%e8%af%bbTPU%ef%bc%9a%e8%ae%be%e8%ae%a1%e5%92%8c%e6%8b%86%e8%a7%a3%e4%b8%80%e5%9d%97ASIC%e8%8a%af%e7%89%87.md">33  解读TPU：设计和拆解一块ASIC芯片.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  理解虚拟机：你在云上拿到的计算机是什么样的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/34%20%20%e7%90%86%e8%a7%a3%e8%99%9a%e6%8b%9f%e6%9c%ba%ef%bc%9a%e4%bd%a0%e5%9c%a8%e4%ba%91%e4%b8%8a%e6%8b%bf%e5%88%b0%e7%9a%84%e8%ae%a1%e7%ae%97%e6%9c%ba%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%ef%bc%9f.md">34  理解虚拟机：你在云上拿到的计算机是什么样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  存储器层次结构全景：数据存储的大金字塔长什么样？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/35%20%20%e5%ad%98%e5%82%a8%e5%99%a8%e5%b1%82%e6%ac%a1%e7%bb%93%e6%9e%84%e5%85%a8%e6%99%af%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e7%9a%84%e5%a4%a7%e9%87%91%e5%ad%97%e5%a1%94%e9%95%bf%e4%bb%80%e4%b9%88%e6%a0%b7%ef%bc%9f.md">35  存储器层次结构全景：数据存储的大金字塔长什么样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  局部性原理：数据库性能跟不上，加个缓存就好了？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/36%20%20%e5%b1%80%e9%83%a8%e6%80%a7%e5%8e%9f%e7%90%86%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%ba%93%e6%80%a7%e8%83%bd%e8%b7%9f%e4%b8%8d%e4%b8%8a%ef%bc%8c%e5%8a%a0%e4%b8%aa%e7%bc%93%e5%ad%98%e5%b0%b1%e5%a5%bd%e4%ba%86%ef%bc%9f.md">36  局部性原理：数据库性能跟不上，加个缓存就好了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  理解CPU Cache（上）：“4毫秒”究竟值多少钱？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/37%20%20%e7%90%86%e8%a7%a3CPU%20Cache%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e2%80%9c4%e6%af%ab%e7%a7%92%e2%80%9d%e7%a9%b6%e7%ab%9f%e5%80%bc%e5%a4%9a%e5%b0%91%e9%92%b1%ef%bc%9f.md">37  理解CPU Cache（上）：“4毫秒”究竟值多少钱？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  高速缓存（下）：你确定你的数据更新了么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/38%20%20%e9%ab%98%e9%80%9f%e7%bc%93%e5%ad%98%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%bd%a0%e7%a1%ae%e5%ae%9a%e4%bd%a0%e7%9a%84%e6%95%b0%e6%8d%ae%e6%9b%b4%e6%96%b0%e4%ba%86%e4%b9%88%ef%bc%9f.md">38  高速缓存（下）：你确定你的数据更新了么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  MESI协议：如何让多核CPU的高速缓存保持一致？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/39%20%20MESI%e5%8d%8f%e8%ae%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e5%a4%9a%e6%a0%b8CPU%e7%9a%84%e9%ab%98%e9%80%9f%e7%bc%93%e5%ad%98%e4%bf%9d%e6%8c%81%e4%b8%80%e8%87%b4%ef%bc%9f.md">39  MESI协议：如何让多核CPU的高速缓存保持一致？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  理解内存（上）：虚拟内存和内存保护是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/40%20%20%e7%90%86%e8%a7%a3%e5%86%85%e5%ad%98%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e8%99%9a%e6%8b%9f%e5%86%85%e5%ad%98%e5%92%8c%e5%86%85%e5%ad%98%e4%bf%9d%e6%8a%a4%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">40  理解内存（上）：虚拟内存和内存保护是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41  理解内存（下）：解析TLB和内存保护.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/41%20%20%e7%90%86%e8%a7%a3%e5%86%85%e5%ad%98%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%a7%a3%e6%9e%90TLB%e5%92%8c%e5%86%85%e5%ad%98%e4%bf%9d%e6%8a%a4.md">41  理解内存（下）：解析TLB和内存保护.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42  总线：计算机内部的高速公路.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/42%20%20%e6%80%bb%e7%ba%bf%ef%bc%9a%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%86%85%e9%83%a8%e7%9a%84%e9%ab%98%e9%80%9f%e5%85%ac%e8%b7%af.md">42  总线：计算机内部的高速公路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43  输入输出设备：我们并不是只能用灯泡显示“0”和“1”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/43%20%20%e8%be%93%e5%85%a5%e8%be%93%e5%87%ba%e8%ae%be%e5%a4%87%ef%bc%9a%e6%88%91%e4%bb%ac%e5%b9%b6%e4%b8%8d%e6%98%af%e5%8f%aa%e8%83%bd%e7%94%a8%e7%81%af%e6%b3%a1%e6%98%be%e7%a4%ba%e2%80%9c0%e2%80%9d%e5%92%8c%e2%80%9c1%e2%80%9d.md">43  输入输出设备：我们并不是只能用灯泡显示“0”和“1”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44  理解IO_WAIT：IO性能到底是怎么回事儿？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/44%20%20%e7%90%86%e8%a7%a3IO_WAIT%ef%bc%9aIO%e6%80%a7%e8%83%bd%e5%88%b0%e5%ba%95%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%e5%84%bf%ef%bc%9f.md">44  理解IO_WAIT：IO性能到底是怎么回事儿？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45  机械硬盘：Google早期用过的“黑科技”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/45%20%20%e6%9c%ba%e6%a2%b0%e7%a1%ac%e7%9b%98%ef%bc%9aGoogle%e6%97%a9%e6%9c%9f%e7%94%a8%e8%bf%87%e7%9a%84%e2%80%9c%e9%bb%91%e7%a7%91%e6%8a%80%e2%80%9d.md">45  机械硬盘：Google早期用过的“黑科技”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46  SSD硬盘（上）：如何完成性能优化的KPI？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/46%20%20SSD%e7%a1%ac%e7%9b%98%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%8c%e6%88%90%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84KPI%ef%bc%9f.md">46  SSD硬盘（上）：如何完成性能优化的KPI？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47  SSD硬盘（下）：如何完成性能优化的KPI？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/47%20%20SSD%e7%a1%ac%e7%9b%98%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%8c%e6%88%90%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e7%9a%84KPI%ef%bc%9f.md">47  SSD硬盘（下）：如何完成性能优化的KPI？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48  DMA：为什么Kafka这么快？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/48%20%20DMA%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Kafka%e8%bf%99%e4%b9%88%e5%bf%ab%ef%bc%9f.md">48  DMA：为什么Kafka这么快？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49  数据完整性（上）：硬件坏了怎么办？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/49%20%20%e6%95%b0%e6%8d%ae%e5%ae%8c%e6%95%b4%e6%80%a7%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e7%a1%ac%e4%bb%b6%e5%9d%8f%e4%ba%86%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">49  数据完整性（上）：硬件坏了怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50  数据完整性（下）：如何还原犯罪现场？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/50%20%20%e6%95%b0%e6%8d%ae%e5%ae%8c%e6%95%b4%e6%80%a7%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%98%e5%8e%9f%e7%8a%af%e7%bd%aa%e7%8e%b0%e5%9c%ba%ef%bc%9f.md">50  数据完整性（下）：如何还原犯罪现场？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51  分布式计算：如果所有人的大脑都联网会怎样？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/51%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e8%ae%a1%e7%ae%97%ef%bc%9a%e5%a6%82%e6%9e%9c%e6%89%80%e6%9c%89%e4%ba%ba%e7%9a%84%e5%a4%a7%e8%84%91%e9%83%bd%e8%81%94%e7%bd%91%e4%bc%9a%e6%80%8e%e6%a0%b7%ef%bc%9f.md">51  分布式计算：如果所有人的大脑都联网会怎样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52  设计大型DMP系统（上）：MongoDB并不是什么灵丹妙药.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/52%20%20%e8%ae%be%e8%ae%a1%e5%a4%a7%e5%9e%8bDMP%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aMongoDB%e5%b9%b6%e4%b8%8d%e6%98%af%e4%bb%80%e4%b9%88%e7%81%b5%e4%b8%b9%e5%a6%99%e8%8d%af.md">52  设计大型DMP系统（上）：MongoDB并不是什么灵丹妙药.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="53  设计大型DMP系统（下）：SSD拯救了所有的DBA.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/53%20%20%e8%ae%be%e8%ae%a1%e5%a4%a7%e5%9e%8bDMP%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aSSD%e6%8b%af%e6%95%91%e4%ba%86%e6%89%80%e6%9c%89%e7%9a%84DBA.md">53  设计大型DMP系统（下）：SSD拯救了所有的DBA.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="54  理解Disruptor（上）：带你体会CPU高速缓存的风驰电掣.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/54%20%20%e7%90%86%e8%a7%a3Disruptor%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%b8%a6%e4%bd%a0%e4%bd%93%e4%bc%9aCPU%e9%ab%98%e9%80%9f%e7%bc%93%e5%ad%98%e7%9a%84%e9%a3%8e%e9%a9%b0%e7%94%b5%e6%8e%a3.md">54  理解Disruptor（上）：带你体会CPU高速缓存的风驰电掣.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="55  理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/55%20%20%e7%90%86%e8%a7%a3Disruptor%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%8d%e9%9c%80%e8%a6%81%e6%8d%a2%e6%8c%a1%e5%92%8c%e8%b8%a9%e5%88%b9%e8%bd%a6%e7%9a%84CPU%ef%bc%8c%e6%9c%89%e5%a4%9a%e5%bf%ab%ef%bc%9f.md">55  理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  知也无涯，愿你也享受发现的乐趣.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%bb%84%e6%88%90%e5%8e%9f%e7%90%86/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e7%9f%a5%e4%b9%9f%e6%97%a0%e6%b6%af%ef%bc%8c%e6%84%bf%e4%bd%a0%e4%b9%9f%e4%ba%ab%e5%8f%97%e5%8f%91%e7%8e%b0%e7%9a%84%e4%b9%90%e8%b6%a3.md">结束语  知也无涯，愿你也享受发现的乐趣.md</a>
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
                            <h1 id="title" data-id="48  DMA：为什么Kafka这么快？" class="title">48  DMA：为什么Kafka这么快？</h1>
                            <div><p>过去几年里，整个计算机产业届，都在尝试不停地提升 I/O 设备的速度。把 HDD 硬盘换成 SSD 硬盘，我们仍然觉得不够快；用 PCI Express 接口的 SSD 硬盘替代 SATA 接口的 SSD 硬盘，我们还是觉得不够快，所以，现在就有了傲腾（Optane）这样的技术。</p>

<p>但是，无论 I/O 速度如何提升，比起 CPU，总还是太慢。SSD 硬盘的 IOPS 可以到 2 万、4 万，但是我们 CPU 的主频有 2GHz 以上，也就意味着每秒会有 20 亿次的操作。</p>

<p>如果我们对于 I/O 的操作，都是由 CPU 发出对应的指令，然后等待 I/O 设备完成操作之后返回，那 CPU 有大量的时间其实都是在等待 I/O 设备完成操作。</p>

<p>但是，这个 CPU 的等待，在很多时候，其实并没有太多的实际意义。我们对于 I/O 设备的大量操作，其实都只是把内存里面的数据，传输到 I/O 设备而已。在这种情况下，其实 CPU 只是在傻等而已。特别是当传输的数据量比较大的时候，比如进行大文件复制，如果所有数据都要经过 CPU，实在是有点儿太浪费时间了。</p>

<p>因此，计算机工程师们，就发明了 DMA 技术，也就是<strong>直接内存访问</strong>（Direct Memory Access）技术，来减少 CPU 等待的时间。</p>

<h2 id="理解-dma-一个协处理器">理解 DMA，一个协处理器</h2>

<p>其实 DMA 技术很容易理解，本质上，DMA 技术就是我们在主板上放一块独立的芯片。在进行内存和 I/O 设备的数据传输的时候，我们不再通过 CPU 来控制数据传输，而直接通过<strong>DMA 控制器</strong>（DMA Controller，简称 DMAC）。这块芯片，我们可以认为它其实就是一个<strong>协处理器</strong>（Co-Processor）。</p>

<p><strong>DMAC 最有价值的地方体现在，当我们要传输的数据特别大、速度特别快，或者传输的数据特别小、速度特别慢的时候。</strong></p>

<p>比如说，我们用千兆网卡或者硬盘传输大量数据的时候，如果都用 CPU 来搬运的话，肯定忙不过来，所以可以选择 DMAC。而当数据传输很慢的时候，DMAC 可以等数据到齐了，再发送信号，给到 CPU 去处理，而不是让 CPU 在那里忙等待。</p>

<p>好了，现在你应该明白 DMAC 的价值，知道了它适合用在什么情况下。那我们现在回过头来看。我们上面说，DMAC 是一块“协处理器芯片”，这是为什么呢？</p>

<p>注意，这里面的“协”字。DMAC 是在“协助”CPU，完成对应的数据传输工作。在 DMAC 控制数据传输的过程中，我们还是需要 CPU 的。</p>

<p>除此之外，DMAC 其实也是一个特殊的 I/O 设备，它和 CPU 以及其他 I/O 设备一样，通过连接到总线来进行实际的数据传输。总线上的设备呢，其实有两种类型。一种我们称之为<strong>主设备</strong>（Master），另外一种，我们称之为<strong>从设备</strong>（Slave）。</p>

<p>想要主动发起数据传输，必须要是一个主设备才可以，CPU 就是主设备。而我们从设备（比如硬盘）只能接受数据传输。所以，如果通过 CPU 来传输数据，要么是 CPU 从 I/O 设备读数据，要么是 CPU 向 I/O 设备写数据。</p>

<p>这个时候你可能要问了，那我们的 I/O 设备不能向主设备发起请求么？可以是可以，不过这个发送的不是数据内容，而是控制信号。I/O 设备可以告诉 CPU，我这里有数据要传输给你，但是实际数据是 CPU 从拉走的，而不是 I/O 设备推给 CPU 的。</p>

<p><img src="assets/6388a8322103eb4d4db288d68aaa3db4.jpeg" alt="img" /></p>

<p>不过，DMAC 就很有意思了，它既是一个主设备，又是一个从设备。对于 CPU 来说，它是一个从设备；对于硬盘这样的 IO 设备来说呢，它又变成了一个主设备。那使用 DMAC 进行数据传输的过程究竟是什么样的呢？下面我们来具体看看。</p>

<p>\1. 首先，CPU 还是作为一个主设备，向 DMAC 设备发起请求。这个请求，其实就是在 DMAC 里面修改配置寄存器。</p>

<p>2.CPU 修改 DMAC 的配置的时候，会告诉 DMAC 这样几个信息：</p>

<ul>
<li>首先是<strong>源地址的初始值以及传输时候的地址增减方式</strong>。
所谓源地址，就是数据要从哪里传输过来。如果我们要从内存里面写入数据到硬盘上，那么就是要读取的数据在内存里面的地址。如果是从硬盘读取数据到内存里，那就是硬盘的 I/O 接口的地址。
我们讲过总线的时候说过，I/O 的地址可以是一个内存地址，也可以是一个端口地址。而地址的增减方式就是说，数据是从大的地址向小的地址传输，还是从小的地址往大的地址传输。</li>
<li>其次是<strong>目标地址初始值和传输时候的地址增减方式</strong>。目标地址自然就是和源地址对应的设备，也就是我们数据传输的目的地。</li>
<li>第三个自然是<strong>要传输的数据长度</strong>，也就是我们一共要传输多少数据。</li>
</ul>

<p>\3. 设置完这些信息之后，DMAC 就会变成一个空闲的状态（Idle）。</p>

<p>\4. 如果我们要从硬盘上往内存里面加载数据，这个时候，硬盘就会向 DMAC 发起一个数据传输请求。这个请求并不是通过总线，而是通过一个额外的连线。</p>

<p>\5. 然后，我们的 DMAC 需要再通过一个额外的连线响应这个申请。</p>

<p>\6. 于是，DMAC 这个芯片，就向硬盘的接口发起要总线读的传输请求。数据就从硬盘里面，读到了 DMAC 的控制器里面。</p>

<p>\7. 然后，DMAC 再向我们的内存发起总线写的数据传输请求，把数据写入到内存里面。</p>

<p>8.DMAC 会反复进行上面第 6、7 步的操作，直到 DMAC 的寄存器里面设置的数据长度传输完成。</p>

<p>\9. 数据传输完成之后，DMAC 重新回到第 3 步的空闲状态。</p>

<p>所以，整个数据传输的过程中，我们不是通过 CPU 来搬运数据，而是由 DMAC 这个芯片来搬运数据。但是 CPU 在这个过程中也是必不可少的。因为传输什么数据，从哪里传输到哪里，其实还是由 CPU 来设置的。这也是为什么，DMAC 被叫作“协处理器”。</p>

<p><img src="assets/c9ed34b47b0cd33867c581772d8eff8e.jpeg" alt="img" /></p>

<p>现在的外设里面，很多都内置了 DMAC</p>

<p>最早，计算机里是没有 DMAC 的，所有数据都是由 CPU 来搬运的。随着对于数据传输的需求越来越多，先是出现了主板上独立的 DMAC 控制器。到了今天，各种 I/O 设备越来越多，数据传输的需求越来越复杂，使用的场景各不相同。加之显示器、网卡、硬盘对于数据传输的需求都不一样，所以各个设备里面都有自己的 DMAC 芯片了。</p>

<h2 id="为什么那么快-一起来看-kafka-的实现原理">为什么那么快？一起来看 Kafka 的实现原理</h2>

<p>了解了 DMAC 是怎么回事儿，那你可能要问了，这和我们实际进行程序开发有什么关系呢？有什么 API，我们直接调用一下，就能加速数据传输，减少 CPU 占用吗？</p>

<p>你还别说，过去几年的大数据浪潮里面，还真有一个开源项目很好地利用了 DMA 的数据传输方式，通过 DMA 的方式实现了非常大的性能提升。这个项目就是<strong>Kafka</strong>。下面我们就一起来看看它究竟是怎么利用 DMA 的。</p>

<p>Kafka 是一个用来处理实时数据的管道，我们常常用它来做一个消息队列，或者用来收集和落地海量的日志。作为一个处理实时数据和日志的管道，瓶颈自然也在 I/O 层面。</p>

<p>Kafka 里面会有两种常见的海量数据传输的情况。一种是从网络中接收上游的数据，然后需要落地到本地的磁盘上，确保数据不丢失。另一种情况呢，则是从本地磁盘上读取出来，通过网络发送出去。</p>

<p>我们来看一看后一种情况，从磁盘读数据发送到网络上去。如果我们自己写一个简单的程序，最直观的办法，自然是用一个文件读操作，从磁盘上把数据读到内存里面来，然后再用一个 Socket，把这些数据发送到网络上去。</p>

<pre><code>File.read(fileDesc, buf, len);
Socket.send(socket, buf, len);
</code></pre>

<p><a href="https://developer.ibm.com/articles/j-zerocopy/" target="_blank">代码来源</a></p>

<p>这段伪代码，来自 IBM Developer Works 上关于 Zero Copy 的文章</p>

<p>在这个过程中，数据一共发生了四次传输的过程。其中两次是 DMA 的传输，另外两次，则是通过 CPU 控制的传输。下面我们来具体看看这个过程。</p>

<p>第一次传输，是从硬盘上，读到操作系统内核的缓冲区里。这个传输是通过 DMA 搬运的。</p>

<p>第二次传输，需要从内核缓冲区里面的数据，复制到我们应用分配的内存里面。这个传输是通过 CPU 搬运的。</p>

<p>第三次传输，要从我们应用的内存里面，再写到操作系统的 Socket 的缓冲区里面去。这个传输，还是由 CPU 搬运的。</p>

<p>最后一次传输，需要再从 Socket 的缓冲区里面，写到网卡的缓冲区里面去。这个传输又是通过 DMA 搬运的。</p>

<p><img src="assets/e0e85505e793e804e3b396fc50871cd5.jpg" alt="img" /></p>

<p>这个时候，你可以回过头看看这个过程。我们只是要“搬运”一份数据，结果却整整搬运了四次。而且这里面，从内核的读缓冲区传输到应用的内存里，再从应用的内存里传输到 Socket 的缓冲区里，其实都是把同一份数据在内存里面搬运来搬运去，特别没有效率。</p>

<p>像 Kafka 这样的应用场景，其实大部分最终利用到的硬件资源，其实又都是在干这个搬运数据的事儿。所以，我们就需要尽可能地减少数据搬运的需求。</p>

<p>事实上，Kafka 做的事情就是，把这个数据搬运的次数，从上面的四次，变成了两次，并且只有 DMA 来进行数据搬运，而不需要 CPU。</p>

<pre><code>@Override
public long transferFrom(FileChannel fileChannel, long position, long count) throws IOException {
    return fileChannel.transferTo(position, count, socketChannel);
}
</code></pre>

<p>如果你层层追踪 Kafka 的代码，你会发现，最终它调用了 Java NIO 库里的 transferTo 方法</p>

<p>Kafka 的代码调用了 Java NIO 库，具体是 FileChannel 里面的 transferTo 方法。我们的数据并没有读到中间的应用内存里面，而是直接通过 Channel，写入到对应的网络设备里。并且，对于 Socket 的操作，也不是写入到 Socket 的 Buffer 里面，而是直接根据描述符（Descriptor）写入到网卡的缓冲区里面。于是，在这个过程之中，我们只进行了两次数据传输。</p>

<p><img src="assets/596042d111ad9b871045d970a10464ab.jpg" alt="img" /></p>

<p>第一次，是通过 DMA，从硬盘直接读到操作系统内核的读缓冲区里面。第二次，则是根据 Socket 的描述符信息，直接从读缓冲区里面，写入到网卡的缓冲区里面。</p>

<p>这样，我们同一份数据传输的次数从四次变成了两次，并且没有通过 CPU 来进行数据搬运，所有的数据都是通过 DMA 来进行传输的。</p>

<p>在这个方法里面，我们没有在内存层面去“复制（Copy）”数据，所以这个方法，也被称之为<strong>零拷贝</strong>（Zero-Copy）。</p>

<p>IBM Developer Works 里面有一篇文章，专门写过程序来测试过在同样的硬件下，使用零拷贝能够带来的性能提升。我在这里放上这篇文章<a href="https://developer.ibm.com/articles/j-zerocopy/" target="_blank">链接</a>。在这篇文章最后，你可以看到，无论传输数据量的大小，传输同样的数据，使用了零拷贝能够缩短 65% 的时间，大幅度提升了机器传输数据的吞吐量。想要深入了解零拷贝，建议你可以仔细读一读这篇文章。</p>

<h2 id="总结延伸">总结延伸</h2>

<p>讲到这里，相信你对 DMA 的原理、作用和效果都有所理解了。那么，我们一起来回顾总结一下。</p>

<p>如果我们始终让 CPU 来进行各种数据传输工作，会特别浪费。一方面，我们的数据传输工作用不到多少 CPU 核心的“计算”功能。另一方面，CPU 的运转速度也比 I/O 操作要快很多。所以，我们希望能够给 CPU“减负”。</p>

<p>于是，工程师们就在主板上放上了 DMAC 这样一个协处理器芯片。通过这个芯片，CPU 只需要告诉 DMAC，我们要传输什么数据，从哪里来，到哪里去，就可以放心离开了。后续的实际数据传输工作，都会有 DMAC 来完成。随着现代计算机各种外设硬件越来越多，光一个通用的 DMAC 芯片不够了，我们在各个外设上都加上了 DMAC 芯片，使得 CPU 很少再需要关心数据传输的工作了。</p>

<p>在我们实际的系统开发过程中，利用好 DMA 的数据传输机制，也可以大幅提升 I/O 的吞吐率。最典型的例子就是 Kafka。</p>

<p>传统地从硬盘读取数据，然后再通过网卡上向外发送，我们需要进行四次数据传输，其中有两次是发生在内存里的缓冲区和对应的硬件设备之间，我们没法节省掉。但是还有两次，完全是通过 CPU 在内存里面进行数据复制。</p>

<p>在 Kafka 里，通过 Java 的 NIO 里面 FileChannel 的 transferTo 方法调用，我们可以不用把数据复制到我们应用程序的内存里面。通过 DMA 的方式，我们可以把数据从内存缓冲区直接写到网卡的缓冲区里面。在使用了这样的零拷贝的方法之后呢，我们传输同样数据的时间，可以缩减为原来的 1/3，相当于提升了 3 倍的吞吐率。</p>

<p>这也是为什么，Kafka 是目前实时数据传输管道的标准解决方案。</p>

<h2 id="推荐阅读">推荐阅读</h2>

<p>学完了这一讲之后，我推荐你阅读一下 Kafka 的论文，<a href="http://notes.stephenholiday.com/Kafka.pdf" target="_blank">Kakfa:a Distrubted Messaging System for Log Processing.</a>。Kafka 的论文其实非常简单易懂，是一个很好的让你了解系统、日志、分布式系统的入门材料。</p>

<p>如果你想要进一步去了解 Kafka，也可以订阅极客时间的专栏“<a href="https://time.geekbang.org/column/intro/191" target="_blank">Kafka 核心技术与实战</a>”。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#375b5b5b0e030606070077505a565e5b1954585a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1683fafb35f667',t:'MTczNDA5ODc5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=55&#32;&#32;理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？>
        <link rel="icon" href="/static/favicon.png">
        <title>55  理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？ </title>
        
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
                            <h1 id="title" data-id="55  理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？" class="title">55  理解Disruptor（下）：不需要换挡和踩刹车的CPU，有多快？</h1>
                            <div><p>上一讲，我们学习了一个精妙的想法，Disruptor 通过缓存行填充，来利用好 CPU 的高速缓存。不知道你做完课后思考题之后，有没有体会到高速缓存在实践中带来的速度提升呢？</p>

<p>不过，利用 CPU 高速缓存，只是 Disruptor“快”的一个因素，那今天我们就来看一看 Disruptor 快的另一个因素，也就是“无锁”，而尽可能发挥 CPU 本身的高速处理性能。</p>

<h2 id="缓慢的锁">缓慢的锁</h2>

<p>Disruptor 作为一个高性能的生产者 - 消费者队列系统，一个核心的设计就是通过 RingBuffer 实现一个无锁队列。</p>

<p>上一讲里我们讲过，Java 里面的基础库里，就有像 LinkedBlockingQueue 这样的队列库。但是，这个队列库比起 Disruptor 里用的 RingBuffer 要慢上很多。慢的第一个原因我们说过，因为链表的数据在内存里面的布局对于高速缓存并不友好，而 RingBuffer 所使用的数组则不然。</p>

<p><img src="assets/9ce732cb22c49a8a26e870dddde66b69.jpeg" alt="img" /></p>

<p>LinkedBlockingQueue 慢，有另外一个重要的因素，那就是它对于锁的依赖。在生产者 - 消费者模式里，我们可能有多个消费者，同样也可能有多个生产者。多个生产者都要往队列的尾指针里面添加新的任务，就会产生多个线程的竞争。于是，在做这个事情的时候，生产者就需要拿到对于队列尾部的锁。同样地，在多个消费者去消费队列头的时候，也就产生竞争。同样消费者也要拿到锁。</p>

<p>那只有一个生产者，或者一个消费者，我们是不是就没有这个锁竞争的问题了呢？很遗憾，答案还是否定的。一般来说，在生产者 - 消费者模式下，消费者要比生产者快。不然的话，队列会产生积压，队列里面的任务会越堆越多。</p>

<p>一方面，你会发现越来越多的任务没有能够及时完成；另一方面，我们的内存也会放不下。虽然生产者 - 消费者模型下，我们都有一个队列来作为缓冲区，但是大部分情况下，这个缓冲区里面是空的。也就是说，即使只有一个生产者和一个消费者者，这个生产者指向的队列尾和消费者指向的队列头是同一个节点。于是，这两个生产者和消费者之间一样会产生锁竞争。</p>

<p>在 LinkedBlockingQueue 上，这个锁机制是通过 synchronized 这个 Java 关键字来实现的。一般情况下，这个锁最终会对应到操作系统层面的加锁机制（OS-based Lock），这个锁机制需要由操作系统的内核来进行裁决。这个裁决，也需要通过一次上下文切换（Context Switch），把没有拿到锁的线程挂起等待。</p>

<p>不知道你还记不记得，我们在<a href="https://time.geekbang.org/column/article/103717" target="_blank">第 28 讲</a>讲过的异常和中断，这里的上下文切换要做的和异常和中断里的是一样的。上下文切换的过程，需要把当前执行线程的寄存器等等的信息，保存到线程栈里面。而这个过程也必然意味着，已经加载到高速缓存里面的指令或者数据，又回到了主内存里面，会进一步拖慢我们的性能。</p>

<p>我们可以按照 Disruptor 介绍资料里提到的 Benchmark，写一段代码来看看，是不是真是这样的。这里我放了一段 Java 代码，代码的逻辑很简单，就是把一个 long 类型的 counter，从 0 自增到 5 亿。一种方式是没有任何锁，另外一个方式是每次自增的时候都要去取一个锁。</p>

<p>你可以在自己的电脑上试试跑一下这个程序。在我这里，两个方式执行所需要的时间分别是 207 毫秒和 9603 毫秒，性能差出了将近 50 倍。</p>

<pre><code>package com.xuwenhao.perf.jmm;
 
 
import java.util.concurrent.atomic.AtomicLong;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
 
 
public class LockBenchmark{
 
 
    public static void runIncrement()
    {
        long counter = 0;
        long max = 500000000L;
        long start = System.currentTimeMillis();
        while (counter &lt; max) {
            counter++;
        }
        long end = System.currentTimeMillis();
        System.out.println(&quot;Time spent is &quot; + (end-start) + &quot;ms without lock&quot;);
    }
 
 
    public static void runIncrementWithLock()
    {
        Lock lock = new ReentrantLock();
        long counter = 0;
        long max = 500000000L;
        long start = System.currentTimeMillis();
        while (counter &lt; max) {
            if (lock.tryLock()){
                counter++;
                lock.unlock();
            }
        }
        long end = System.currentTimeMillis();
        System.out.println(&quot;Time spent is &quot; + (end-start) + &quot;ms with lock&quot;);
    }
 
 
    public static void main(String[] args) {
        runIncrement();
        runIncrementWithLock();
</code></pre>

<p>加锁和不加锁自增 counter</p>

<pre><code>Time spent is 207ms without lock
Time spent is 9603ms with lock
</code></pre>

<p>性能差出将近 10 倍</p>

<h2 id="无锁的-ringbuffer">无锁的 RingBuffer</h2>

<p>加锁很慢，所以 Disruptor 的解决方案就是“无锁”。这个“无锁”指的是没有操作系统层面的锁。实际上，Disruptor 还是利用了一个 CPU 硬件支持的指令，称之为 CAS（Compare And Swap，比较和交换）。在 Intel CPU 里面，这个对应的指令就是 cmpxchg。那么下面，我们就一起从 Disruptor 的源码，到具体的硬件指令来看看这是怎么一回事儿。</p>

<p>Disruptor 的 RingBuffer 是这么设计的，它和直接在链表的头和尾加锁不同。Disruptor 的 RingBuffer 创建了一个 Sequence 对象，用来指向当前的 RingBuffer 的头和尾。这个头和尾的标识呢，不是通过一个指针来实现的，而是通过一个<strong>序号</strong>。这也是为什么对应源码里面的类名叫 Sequence。</p>

<p><img src="assets/b64487a7b6b45393fdfa7e2d63e176ec.jpeg" alt="img" /></p>

<p>在这个 RingBuffer 当中，进行生产者和消费者之间的资源协调，采用的是对比序号的方式。当生产者想要往队列里加入新数据的时候，它会把当前的生产者的 Sequence 的序号，加上需要加入的新数据的数量，然后和实际的消费者所在的位置进行对比，看看队列里是不是有足够的空间加入这些数据，而不会覆盖掉消费者还没有处理完的数据。</p>

<p>在 Sequence 的代码里面，就是通过 compareAndSet 这个方法，并且最终调用到了 UNSAFE.compareAndSwapLong，也就是直接使用了 CAS 指令。</p>

<pre><code> public boolean compareAndSet(final long expectedValue, final long newValue)
	    {
	        return UNSAFE.compareAndSwapLong(this, VALUE_OFFSET, expectedValue, newValue);
	    }
 
 
public long addAndGet(final long increment)
    {
        long currentValue;
        long newValue;
 
 
        do
        {
            currentValue = get();
            newValue = currentValue + increment;
        }
        while (!compareAndSet(currentValue, newValue));
 
 
        return newValue;
</code></pre>

<p>Sequence 源码中的 addAndGet，如果 CAS 的操作没有成功，它会不断忙等待地重试</p>

<p>这个 CAS 指令，也就是比较和交换的操作，并不是基础库里的一个函数。它也不是操作系统里面实现的一个系统调用，而是<strong>一个 CPU 硬件支持的机器指令</strong>。在我们服务器所使用的 Intel CPU 上，就是 cmpxchg 这个指令。</p>

<pre><code>compxchg [ax] (隐式参数，EAX 累加器), [bx] (源操作数地址), [cx] (目标操作数地址)
复制代码
</code></pre>

<p>cmpxchg 指令，一共有三个操作数，第一个操作数不在指令里面出现，是一个隐式的操作数，也就是 EAX 累加寄存器里面的值。第二个操作数就是源操作数，并且指令会对比这个操作数和上面的累加寄存器里面的值。</p>

<p>如果值是相同的，那一方面，CPU 会把 ZF（也就是条件码寄存器里面零标志位的值）设置为 1，然后再把第三个操作数（也就是目标操作数），设置到源操作数的地址上。如果不相等的话，就会把源操作数里面的值，设置到累加器寄存器里面。</p>

<p>我在这里放了这个逻辑对应的伪代码，你可以看一下。如果你对汇编指令、条件码寄存器这些知识点有点儿模糊了，可以回头去看看<a href="https://time.geekbang.org/column/article/93359" target="_blank">第 5</a><a href="https://time.geekbang.org/column/article/93359" target="_blank">讲</a>、<a href="https://time.geekbang.org/column/article/94075" target="_blank">第 6 讲</a>关于汇编指令的部分。</p>

<pre><code>IF [ax]&lt; == [bx] THEN [ZF] = 1, [bx] = [cx]
                 ELSE [ZF] = 0, [ax] = [bx] 
</code></pre>

<p>单个指令是原子的，这也就意味着在使用 CAS 操作的时候，我们不再需要单独进行加锁，直接调用就可以了。</p>

<p>没有了锁，CPU 这部高速跑车就像在赛道上行驶，不会遇到需要上下文切换这样的红灯而停下来。虽然会遇到像 CAS 这样复杂的机器指令，就好像赛道上会有 U 型弯一样，不过不用完全停下来等待，我们 CPU 运行起来仍然会快很多。</p>

<p>那么，CAS 操作到底会有多快呢？我们还是用一段 Java 代码来看一下。</p>

<pre><code>package com.xuwenhao.perf.jmm;
 
 
import java.util.concurrent.atomic.AtomicLong;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
 
 
public class LockBenchmark {
 
 
    public static void runIncrementAtomic()
    {
        AtomicLong counter = new AtomicLong(0);
        long max = 500000000L;
        long start = System.currentTimeMillis();
        while (counter.incrementAndGet() &lt; max) {
        }
        long end = System.currentTimeMillis();
        System.out.println(&quot;Time spent is &quot; + (end-start) + &quot;ms with cas&quot;);
    }
 
 
    public static void main(String[] args) {
        runIncrementAtomic();
    }
Time spent is 3867ms with cas
复制代码
</code></pre>

<p>和上面的 counter 自增一样，只不过这一次，自增我们采用了 AtomicLong 这个 Java 类。里面的 incrementAndGet 最终到了 CPU 指令层面，在实现的时候用的就是 CAS 操作。可以看到，它所花费的时间，虽然要比没有任何锁的操作慢上一个数量级，但是比起使用 ReentrantLock 这样的操作系统锁的机制，还是减少了一半以上的时间。</p>

<h2 id="总结延伸">总结延伸</h2>

<p>好了，咱们专栏的正文内容到今天就要结束了。今天最后一讲，我带着你一起看了 Disruptor 代码的一个核心设计，也就是它的 RingBuffer 是怎么做到无锁的。</p>

<p>Java 基础库里面的 BlockingQueue，都需要通过显示地加锁来保障生产者之间、消费者之间，乃至生产者和消费者之间，不会发生锁冲突的问题。</p>

<p>但是，加锁会大大拖慢我们的性能。在获取锁过程中，CPU 没有去执行计算的相关指令，而要等待操作系统进行锁竞争的裁决。而那些没有拿到锁而被挂起等待的线程，则需要进行上下文切换。这个上下文切换，会把挂起线程的寄存器里的数据放到线程的程序栈里面去。这也意味着，加载到高速缓存里面的数据也失效了，程序就变得更慢了。</p>

<p>Disruptor 里的 RingBuffer 采用了一个无锁的解决方案，通过 CAS 这样的操作，去进行序号的自增和对比，使得 CPU 不需要获取操作系统的锁。而是能够继续顺序地执行 CPU 指令。没有上下文切换、没有操作系统锁，自然程序就跑得快了。不过因为采用了 CAS 这样的忙等待（Busy-Wait）的方式，会使得我们的 CPU 始终满负荷运转，消耗更多的电，算是一个小小的缺点。</p>

<p>程序里面的 CAS 调用，映射到我们的 CPU 硬件层面，就是一个机器指令，这个指令就是 cmpxchg。可以看到，当想要追求最极致的性能的时候，我们会从应用层、贯穿到操作系统，乃至最后的 CPU 硬件，搞清楚从高级语言到系统调用，乃至最后的汇编指令，这整个过程是怎么执行代码的。而这个，也是学习组成原理这门专栏的意义所在。</p>

<h2 id="推荐阅读">推荐阅读</h2>

<p>不知道上一讲说的 Disruptor 相关材料，你有没有读完呢？如果没有读完的话，我建议你还是先去研读一下。</p>

<p>如果你已经读完了，这里再给你推荐一些额外的阅读材料，那就是著名的<a href="http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.53.8674&amp;rep=rep1&amp;type=pdf" target="_blank">Implement Lock-Free Queues</a>这篇论文。你可以更深入地学习一下，怎么实现一个无锁队列。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#96fafafaafa2a7a7a6a1d6f1fbf7fffab8f5f9fb" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f16862cbfddf667',t:'MTczNDA5ODg4NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
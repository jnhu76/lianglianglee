<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=26&#32;&#32;Superscalar和VLIW：如何让CPU的吞吐率超过1？>
        <link rel="icon" href="/static/favicon.png">
        <title>26  Superscalar和VLIW：如何让CPU的吞吐率超过1？ </title>
        
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
                            <h1 id="title" data-id="26  Superscalar和VLIW：如何让CPU的吞吐率超过1？" class="title">26  Superscalar和VLIW：如何让CPU的吞吐率超过1？</h1>
                            <div><p>到今天为止，专栏已经过半了。过去的 20 多讲里，我给你讲的内容，很多都是围绕着怎么提升 CPU 的性能这个问题展开的。</p>

<p>我们先回顾一下<a href="https://time.geekbang.org/column/article/93246" target="_blank">第 4 讲</a>，不知道你是否还记得这个公式：</p>

<p>程序的 CPU 执行时间 = 指令数 × CPI × Clock Cycle Time</p>

<p>这个公式里，有一个叫 CPI 的指标。我们知道，CPI 的倒数，又叫作 IPC（Instruction Per Clock），也就是一个时钟周期里面能够执行的指令数，代表了 CPU 的吞吐率。那么，这个指标，放在我们前面几节反复优化流水线架构的 CPU 里，能达到多少呢？</p>

<p>答案是，最佳情况下，IPC 也只能到 1。因为无论做了哪些流水线层面的优化，即使做到了指令执行层面的乱序执行，CPU 仍然只能在一个时钟周期里面，取一条指令。</p>

<p><img src="assets/dd88d0dbf3a88b09d5e8fb6d9e3aea13.jpeg" alt="img" /></p>

<p>这说明，无论指令后续能优化得多好，一个时钟周期也只能执行完这样一条指令，CPI 只能是 1。但是，我们现在用的 Intel CPU 或者 ARM 的 CPU，一般的 CPI 都能做到 2 以上，这是怎么做到的呢？</p>

<p>今天，我们就一起来看看，现代 CPU 都使用了什么“黑科技”。</p>

<h2 id="多发射与超标量-同一实践执行的两条指令">多发射与超标量：同一实践执行的两条指令</h2>

<p>之前讲 CPU 的硬件组成的时候，我们把所有算术和逻辑运算都抽象出来，变成了一个 ALU 这样的“黑盒子”。你应该还记得第 13 讲到第 16 讲，关于加法器、乘法器、乃至浮点数计算的部分，其实整数的计算和浮点数的计算过程差异还是不小的。实际上，整数和浮点数计算的电路，在 CPU 层面也是分开的。</p>

<p>一直到 80386，我们的 CPU 都是没有专门的浮点数计算的电路的。当时的浮点数计算，都是用软件进行模拟的。所以，在 80386 时代，Intel 给 386 配了单独的 387 芯片，专门用来做浮点数运算。那个时候，你买 386 芯片的话，会有 386sx 和 386dx 这两种芯片可以选择。386dx 就是带了 387 浮点数计算芯片的，而 sx 就是不带浮点数计算芯片的。</p>

<p>其实，我们现在用的 Intel CPU 芯片也是一样的。虽然浮点数计算已经变成 CPU 里的一部分，但并不是所有计算功能都在一个 ALU 里面，真实的情况是，我们会有多个 ALU。这也是为什么，在<a href="https://time.geekbang.org/column/article/101436" target="_blank">第 24 讲</a>讲乱序执行的时候，你会看到，其实指令的执行阶段，是由很多个功能单元（FU）并行（Parallel）进行的。</p>

<p>不过，在指令乱序执行的过程中，我们的取指令（IF）和指令译码（ID）部分并不是并行进行的。</p>

<p>既然指令的执行层面可以并行进行，为什么取指令和指令译码不行呢？如果想要实现并行，该怎么办呢？</p>

<p>其实只要我们把取指令和指令译码，也一样通过增加硬件的方式，并行进行就好了。我们可以一次性从内存里面取出多条指令，然后分发给多个并行的指令译码器，进行译码，然后对应交给不同的功能单元去处理。这样，我们在一个时钟周期里，能够完成的指令就不只一条了。IPC 也就能做到大于 1 了。</p>

<p><img src="assets/85f15ec667d09fd2d368822904029b32.jpeg" alt="img" /></p>

<p>这种 CPU 设计，我们叫作<strong>多发射</strong>（Mulitple Issue）和<strong>超标量</strong>（Superscalar）。</p>

<p>什么叫多发射呢？这个词听起来很抽象，其实它意思就是说，我们同一个时间，可能会同时把多条指令发射（Issue）到不同的译码器或者后续处理的流水线中去。</p>

<p>在超标量的 CPU 里面，有很多条并行的流水线，而不是只有一条流水线。“超标量“这个词是说，本来我们在一个时钟周期里面，只能执行一个标量（Scalar）的运算。在多发射的情况下，我们就能够超越这个限制，同时进行多次计算。</p>

<p><img src="assets/2e96fe0985a4ae3bd7a58c345def29d3.jpeg" alt="img" /></p>

<p>你可以看我画的这个超标量设计的流水线示意图。仔细看，你应该能看到一个有意思的现象，每一个功能单元的流水线的长度是不同的。事实上，不同的功能单元的流水线长度本来就不一样。我们平时所说的 14 级流水线，指的通常是进行整数计算指令的流水线长度。如果是浮点数运算，实际的流水线长度则会更长一些。</p>

<h2 id="intel-的失败之作-安腾的超长指令字设计">Intel 的失败之作：安腾的超长指令字设计</h2>

<p>无论是之前几讲里讲的乱序执行，还是现在更进一步的超标量技术，在实际的硬件层面，其实实施起来都挺麻烦的。这是因为，在乱序执行和超标量的体系里面，我们的 CPU 要解决依赖冲突的问题。这也就是前面几讲我们讲的冒险问题。</p>

<p>CPU 需要在指令执行之前，去判断指令之间是否有依赖关系。如果有对应的依赖关系，指令就不能分发到执行阶段。因为这样，上面我们所说的超标量 CPU 的多发射功能，又被称为<strong>动态多发射处理器</strong>。这些对于依赖关系的检测，都会使得我们的 CPU 电路变得更加复杂。</p>

<p>于是，计算机科学家和工程师们就又有了一个大胆的想法。我们能不能不把分析和解决依赖关系的事情，放在硬件里面，而是放到软件里面来干呢？</p>

<p>如果你还记得的话，我在第 4 讲也讲过，要想优化 CPU 的执行时间，关键就是拆解这个公式：</p>

<p>程序的 CPU 执行时间 = 指令数 × CPI × Clock Cycle Time</p>

<p>当时我们说过，这个公式里面，我们可以通过改进编译器来优化指令数这个指标。那接下来，我们就来看看一个非常大胆的 CPU 设计想法，叫作<strong>超长指令字设计</strong>（Very Long Instruction Word，VLIW）。这个设计呢，不仅想让编译器来优化指令数，还想直接通过编译器，来优化 CPI。</p>

<p>围绕着这个设计的，是 Intel 一个著名的“史诗级”失败，也就是著名的 IA-64 架构的安腾（Itanium）处理器。只不过，这一次，责任不全在 Intel，还要拉上可以称之为硅谷起源的另一家公司，也就是惠普。</p>

<p>之所以称为“史诗”级失败，这个说法来源于惠普最早给这个架构取的名字，<strong>显式并发指令运算</strong>（Explicitly Parallel Instruction Computer），这个名字的缩写<strong>EPIC</strong>，正好是“史诗”的意思。</p>

<p>好巧不巧，安腾处理器和和我之前给你介绍过的 Pentium 4 一样，在市场上是一个失败的产品。在经历了 12 年之久的设计研发之后，安腾一代只卖出了几千套。而安腾二代，在从 2002 年开始反复挣扎了 16 年之后，最终在 2018 年被 Intel 宣告放弃，退出了市场。自此，世上再也没有这个“史诗”服务器了。</p>

<p>那么，我们就来看看，这个超长指令字的安腾处理器是怎么回事儿。</p>

<p>在乱序执行和超标量的 CPU 架构里，指令的前后依赖关系，是由 CPU 内部的硬件电路来检测的。而到了<strong>超长指令字</strong>的架构里面，这个工作交给了编译器这个软件。</p>

<p><img src="assets/22b3f723ceee5950ac20a7b874dabbde.jpeg" alt="img" /></p>

<p>我从专栏第 5 讲开始，就给你看了不少 C 代码到汇编代码和机器代码的对照。编译器在这个过程中，其实也能够知道前后数据的依赖。于是，我们可以让编译器把没有依赖关系的代码位置进行交换。然后，再把多条连续的指令打包成一个指令包。安腾的 CPU 就是把 3 条指令变成一个指令包。</p>

<p><img src="assets/f16a1ae443418caca0dc2fc3cec200f6.jpeg" alt="img" /></p>

<p>CPU 在运行的时候，不再是取一条指令，而是取出一个指令包。然后，译码解析整个指令包，解析出 3 条指令直接并行运行。可以看到，使用<strong>超长指令字</strong>架构的 CPU，同样是采用流水线架构的。也就是说，一组（Group）指令，仍然要经历多个时钟周期。同样的，下一组指令并不是等上一组指令执行完成之后再执行，而是在上一组指令的指令译码阶段，就开始取指令了。</p>

<p>值得注意的一点是，流水线停顿这件事情在<strong>超长指令字</strong>里面，很多时候也是由编译器来做的。除了停下整个处理器流水线，<strong>超长指令字</strong>的 CPU 不能在某个时钟周期停顿一下，等待前面依赖的操作执行完成。编译器需要在适当的位置插入 NOP 操作，直接在编译出来的机器码里面，就把流水线停顿这个事情在软件层面就安排妥当。</p>

<p>虽然安腾的设想很美好，Intel 也曾经希望能够让安腾架构成为替代 x86 的新一代架构，但是最终安腾还是在前前后后折腾将近 30 年后失败了。2018 年，Intel 宣告安腾 9500 会在 2021 年停止供货。</p>

<p>安腾失败的原因有很多，其中有一个重要的原因就是“向前兼容”。</p>

<p>一方面，安腾处理器的指令集和 x86 是不同的。这就意味着，原来 x86 上的所有程序是没有办法在安腾上运行的，而需要通过编译器重新编译才行。</p>

<p>另一方面，安腾处理器的 VLIW 架构决定了，如果安腾需要提升并行度，就需要增加一个指令包里包含的指令数量，比方说从 3 个变成 6 个。一旦这么做了，虽然同样是 VLIW 架构，同样指令集的安腾 CPU，程序也需要重新编译。因为原来编译器判断的依赖关系是在 3 个指令以及由 3 个指令组成的指令包之间，现在要变成 6 个指令和 6 个指令组成的指令包。编译器需要重新编译，交换指令顺序以及 NOP 操作，才能满足条件。甚至，我们需要重新来写编译器，才能让程序在新的 CPU 上跑起来。</p>

<p>于是，安腾就变成了一个既不容易向前兼容，又不容易向后兼容的 CPU。那么，它的失败也就不足为奇了。</p>

<p>可以看到，技术思路上的先进想法，在实际的业界应用上会遇到更多具体的实践考验。无论是指令集向前兼容性，还是对应 CPU 未来的扩展，在设计的时候，都需要更多地去考虑实践因素。</p>

<h2 id="总结延伸">总结延伸</h2>

<p>这一讲里，我和你一起向 CPU 的性能发起了一个新的挑战：让 CPU 的吞吐率，也就是 IPC 能够超过 1。</p>

<p>我先是为你介绍了超标量，也就是 Superscalar 这个方法。超标量可以让 CPU 不仅在指令执行阶段是并行的，在取指令和指令译码的时候，也是并行的。通过超标量技术，可以使得你所使用的 CPU 的 IPC 超过 1。</p>

<p>在 Intel 的 x86 的 CPU 里，从 Pentium 时代，第一次开始引入超标量技术，整个 CPU 的性能上了一个台阶。对应的技术，一直沿用到了现在。超标量技术和你之前看到的其他流水线技术一样，依赖于在硬件层面，能够检测到对应的指令的先后依赖关系，解决“冒险”问题。所以，它也使得 CPU 的电路变得更复杂了。</p>

<p>因为这些复杂性，惠普和 Intel 又共同推出了著名的安腾处理器。通过在编译器层面，直接分析出指令的前后依赖关系。于是，硬件在代码编译之后，就可以直接拿到调换好先后顺序的指令。并且这些指令中，可以并行执行的部分，会打包在一起组成一个指令包。安腾处理器在取指令和指令译码的时候，拿到的不再是单个指令，而是这样一个指令包。并且在指令执行阶段，可以并行执行指令包里所有的指令。</p>

<p>虽然看起来，VLIW 在技术层面更具有颠覆性，不仅仅只是一个硬件层面的改造，而且利用了软件层面的编译器，来组合解决提升 CPU 指令吞吐率的问题。然而，最终 VLIW 却没有得到市场和业界的认可。</p>

<p>惠普和 Intel 强强联合开发的安腾处理器命运多舛。从 1989 开始研发，直到 2001 年才发布了第一代安腾处理器。然而 12 年的开发过程后，第一代安腾处理器最终只卖出了几千套。而 2002 年发布的安腾 2 处理器，也没能拯救自己的命运。最终在 2018 年，Intel 宣布安腾退出市场。自此之后，市面上再没有能够大规模商用的 VLIW 架构的处理器了。</p>

<h2 id="推荐阅读">推荐阅读</h2>

<p>关于超标量和多发射的相关知识，你可以多看一看《计算机组成与设计：硬件 / 软件接口》的 4.10 部分。其中，4.10.1 和 4.10.2 的推测和静态多发射，其实就是今天我们讲的超长指令字（VLIW）的知识点。4.10.2 的动态多发射，其实就是今天我们讲的超标量（Superscalar）的知识点。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#472b2b2b7e737676777007202a262e2b6924282a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f167d557878f667',t:'MTczNDA5ODUyMi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=21&#32;RISC-V指令精讲（六）：加载指令实现与调试>
        <link rel="icon" href="/static/favicon.png">
        <title>21 RISC-V指令精讲（六）：加载指令实现与调试 </title>
        
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
                        <a class="menu-item" id="00 开篇词 练好基本功，优秀工程师成长第一步.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e7%bb%83%e5%a5%bd%e5%9f%ba%e6%9c%ac%e5%8a%9f%ef%bc%8c%e4%bc%98%e7%a7%80%e5%b7%a5%e7%a8%8b%e5%b8%88%e6%88%90%e9%95%bf%e7%ac%ac%e4%b8%80%e6%ad%a5.md">00 开篇词 练好基本功，优秀工程师成长第一步.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 CISC &amp; RISC：从何而来，何至于此.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/01%20CISC%20&amp;%20RISC%ef%bc%9a%e4%bb%8e%e4%bd%95%e8%80%8c%e6%9d%a5%ef%bc%8c%e4%bd%95%e8%87%b3%e4%ba%8e%e6%ad%a4.md">01 CISC &amp; RISC：从何而来，何至于此.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 RISC特性与发展：RISC-V凭什么成为“半导体行业的Linux”？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/02%20RISC%e7%89%b9%e6%80%a7%e4%b8%8e%e5%8f%91%e5%b1%95%ef%bc%9aRISC-V%e5%87%ad%e4%bb%80%e4%b9%88%e6%88%90%e4%b8%ba%e2%80%9c%e5%8d%8a%e5%af%bc%e4%bd%93%e8%a1%8c%e4%b8%9a%e7%9a%84Linux%e2%80%9d%ef%bc%9f.md">02 RISC特性与发展：RISC-V凭什么成为“半导体行业的Linux”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 硬件语言筑基（一）：从硬件语言开启手写CPU之旅.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/03%20%e7%a1%ac%e4%bb%b6%e8%af%ad%e8%a8%80%e7%ad%91%e5%9f%ba%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e4%bb%8e%e7%a1%ac%e4%bb%b6%e8%af%ad%e8%a8%80%e5%bc%80%e5%90%af%e6%89%8b%e5%86%99CPU%e4%b9%8b%e6%97%85.md">03 硬件语言筑基（一）：从硬件语言开启手写CPU之旅.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 硬件语言筑基（二）_ 代码是怎么生成具体电路的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/04%20%e7%a1%ac%e4%bb%b6%e8%af%ad%e8%a8%80%e7%ad%91%e5%9f%ba%ef%bc%88%e4%ba%8c%ef%bc%89_%20%e4%bb%a3%e7%a0%81%e6%98%af%e6%80%8e%e4%b9%88%e7%94%9f%e6%88%90%e5%85%b7%e4%bd%93%e7%94%b5%e8%b7%af%e7%9a%84%ef%bc%9f.md">04 硬件语言筑基（二）_ 代码是怎么生成具体电路的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 指令架构：RISC-V在CPU设计上到底有哪些优势？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/05%20%e6%8c%87%e4%bb%a4%e6%9e%b6%e6%9e%84%ef%bc%9aRISC-V%e5%9c%a8CPU%e8%ae%be%e8%ae%a1%e4%b8%8a%e5%88%b0%e5%ba%95%e6%9c%89%e5%93%aa%e4%ba%9b%e4%bc%98%e5%8a%bf%ef%bc%9f.md">05 指令架构：RISC-V在CPU设计上到底有哪些优势？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 手写CPU（一）：迷你CPU架构设计与取指令实现.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/06%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e8%bf%b7%e4%bd%a0CPU%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%e4%b8%8e%e5%8f%96%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0.md">06 手写CPU（一）：迷你CPU架构设计与取指令实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 手写CPU（二）：如何实现指令译码模块？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/07%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%8c%87%e4%bb%a4%e8%af%91%e7%a0%81%e6%a8%a1%e5%9d%97%ef%bc%9f.md">07 手写CPU（二）：如何实现指令译码模块？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 手写CPU（三）：如何实现指令执行模块？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/08%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%8c%87%e4%bb%a4%e6%89%a7%e8%a1%8c%e6%a8%a1%e5%9d%97%ef%bc%9f.md">08 手写CPU（三）：如何实现指令执行模块？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 手写CPU（四）：如何实现CPU流水线的访存阶段？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/09%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0CPU%e6%b5%81%e6%b0%b4%e7%ba%bf%e7%9a%84%e8%ae%bf%e5%ad%98%e9%98%b6%e6%ae%b5%ef%bc%9f.md">09 手写CPU（四）：如何实现CPU流水线的访存阶段？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 手写CPU（五）：CPU流水线的写回模块如何实现？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/10%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e4%ba%94%ef%bc%89%ef%bc%9aCPU%e6%b5%81%e6%b0%b4%e7%ba%bf%e7%9a%84%e5%86%99%e5%9b%9e%e6%a8%a1%e5%9d%97%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">10 手写CPU（五）：CPU流水线的写回模块如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 手写CPU（六）：如何让我们的CPU跑起来？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/11%20%e6%89%8b%e5%86%99CPU%ef%bc%88%e5%85%ad%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e6%88%91%e4%bb%ac%e7%9a%84CPU%e8%b7%91%e8%b5%b7%e6%9d%a5%ef%bc%9f.md">11 手写CPU（六）：如何让我们的CPU跑起来？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 QEMU：支持RISC-V的QEMU如何构建？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/12%20QEMU%ef%bc%9a%e6%94%af%e6%8c%81RISC-V%e7%9a%84QEMU%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%ef%bc%9f.md">12 QEMU：支持RISC-V的QEMU如何构建？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 小试牛刀：跑通RISC-V平台的Hello World程序.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/13%20%e5%b0%8f%e8%af%95%e7%89%9b%e5%88%80%ef%bc%9a%e8%b7%91%e9%80%9aRISC-V%e5%b9%b3%e5%8f%b0%e7%9a%84Hello%20World%e7%a8%8b%e5%ba%8f.md">13 小试牛刀：跑通RISC-V平台的Hello World程序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 走进C语言：高级语言怎样抽象执行逻辑？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/14%20%e8%b5%b0%e8%bf%9bC%e8%af%ad%e8%a8%80%ef%bc%9a%e9%ab%98%e7%ba%a7%e8%af%ad%e8%a8%80%e6%80%8e%e6%a0%b7%e6%8a%bd%e8%b1%a1%e6%89%a7%e8%a1%8c%e9%80%bb%e8%be%91%ef%bc%9f.md">14 走进C语言：高级语言怎样抽象执行逻辑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 C与汇编：揭秘C语言编译器的“搬砖”日常.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/15%20C%e4%b8%8e%e6%b1%87%e7%bc%96%ef%bc%9a%e6%8f%ad%e7%a7%98C%e8%af%ad%e8%a8%80%e7%bc%96%e8%af%91%e5%99%a8%e7%9a%84%e2%80%9c%e6%90%ac%e7%a0%96%e2%80%9d%e6%97%a5%e5%b8%b8.md">15 C与汇编：揭秘C语言编译器的“搬砖”日常.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 RISC-V指令精讲（一）：算术指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/16%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e7%ae%97%e6%9c%af%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">16 RISC-V指令精讲（一）：算术指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 RISC-V指令精讲（二）：算术指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/17%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%ae%97%e6%9c%af%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">17 RISC-V指令精讲（二）：算术指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 RISC-V指令精讲（三）：跳转指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/18%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e8%b7%b3%e8%bd%ac%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">18 RISC-V指令精讲（三）：跳转指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 RISC-V指令精讲（四）：跳转指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/19%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e8%b7%b3%e8%bd%ac%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">19 RISC-V指令精讲（四）：跳转指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 RISC-V指令精讲（五）：原子指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/20%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e4%ba%94%ef%bc%89%ef%bc%9a%e5%8e%9f%e5%ad%90%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">20 RISC-V指令精讲（五）：原子指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 RISC-V指令精讲（六）：加载指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/21%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e5%85%ad%ef%bc%89%ef%bc%9a%e5%8a%a0%e8%bd%bd%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">21 RISC-V指令精讲（六）：加载指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 RISC-V指令精讲（七）：访存指令实现与调试.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/22%20RISC-V%e6%8c%87%e4%bb%a4%e7%b2%be%e8%ae%b2%ef%bc%88%e4%b8%83%ef%bc%89%ef%bc%9a%e8%ae%bf%e5%ad%98%e6%8c%87%e4%bb%a4%e5%ae%9e%e7%8e%b0%e4%b8%8e%e8%b0%83%e8%af%95.md">22 RISC-V指令精讲（七）：访存指令实现与调试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 内存地址空间：程序中地址的三种产生方式.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/23%20%e5%86%85%e5%ad%98%e5%9c%b0%e5%9d%80%e7%a9%ba%e9%97%b4%ef%bc%9a%e7%a8%8b%e5%ba%8f%e4%b8%ad%e5%9c%b0%e5%9d%80%e7%9a%84%e4%b8%89%e7%a7%8d%e4%ba%a7%e7%94%9f%e6%96%b9%e5%bc%8f.md">23 内存地址空间：程序中地址的三种产生方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 虚实结合：虚拟内存和物理内存.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/24%20%e8%99%9a%e5%ae%9e%e7%bb%93%e5%90%88%ef%bc%9a%e8%99%9a%e6%8b%9f%e5%86%85%e5%ad%98%e5%92%8c%e7%89%a9%e7%90%86%e5%86%85%e5%ad%98.md">24 虚实结合：虚拟内存和物理内存.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 堆&amp;栈：堆与栈的区别和应用.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/25%20%e5%a0%86&amp;%e6%a0%88%ef%bc%9a%e5%a0%86%e4%b8%8e%e6%a0%88%e7%9a%84%e5%8c%ba%e5%88%ab%e5%92%8c%e5%ba%94%e7%94%a8.md">25 堆&amp;栈：堆与栈的区别和应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 延迟分配：提高内存利用率的三种机制.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/26%20%e5%bb%b6%e8%bf%9f%e5%88%86%e9%85%8d%ef%bc%9a%e6%8f%90%e9%ab%98%e5%86%85%e5%ad%98%e5%88%a9%e7%94%a8%e7%8e%87%e7%9a%84%e4%b8%89%e7%a7%8d%e6%9c%ba%e5%88%b6.md">26 延迟分配：提高内存利用率的三种机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 应用内存管理：Linux的应用与内存管理.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/27%20%e5%ba%94%e7%94%a8%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%ef%bc%9aLinux%e7%9a%84%e5%ba%94%e7%94%a8%e4%b8%8e%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86.md">27 应用内存管理：Linux的应用与内存管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 进程调度：应用为什么能并行执行？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/28%20%e8%bf%9b%e7%a8%8b%e8%b0%83%e5%ba%a6%ef%bc%9a%e5%ba%94%e7%94%a8%e4%b8%ba%e4%bb%80%e4%b9%88%e8%83%bd%e5%b9%b6%e8%a1%8c%e6%89%a7%e8%a1%8c%ef%bc%9f.md">28 进程调度：应用为什么能并行执行？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 应用间通信（一）：详解Linux进程IPC.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/29%20%e5%ba%94%e7%94%a8%e9%97%b4%e9%80%9a%e4%bf%a1%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e8%af%a6%e8%a7%a3Linux%e8%bf%9b%e7%a8%8bIPC.md">29 应用间通信（一）：详解Linux进程IPC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  应用间通信（二）：详解Linux进程IPC.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/30%20%20%e5%ba%94%e7%94%a8%e9%97%b4%e9%80%9a%e4%bf%a1%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e8%af%a6%e8%a7%a3Linux%e8%bf%9b%e7%a8%8bIPC.md">30  应用间通信（二）：详解Linux进程IPC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 外设通信：IO Cache与IO调度.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/31%20%e5%a4%96%e8%ae%be%e9%80%9a%e4%bf%a1%ef%bc%9aIO%20Cache%e4%b8%8eIO%e8%b0%83%e5%ba%a6.md">31 外设通信：IO Cache与IO调度.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 IO管理：Linux如何管理多个外设？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/32%20IO%e7%ae%a1%e7%90%86%ef%bc%9aLinux%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e5%a4%9a%e4%b8%aa%e5%a4%96%e8%ae%be%ef%bc%9f.md">32 IO管理：Linux如何管理多个外设？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 lotop与lostat命令：聊聊命令背后的故事与工作原理.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/33%20lotop%e4%b8%8elostat%e5%91%bd%e4%bb%a4%ef%bc%9a%e8%81%8a%e8%81%8a%e5%91%bd%e4%bb%a4%e8%83%8c%e5%90%8e%e7%9a%84%e6%95%85%e4%ba%8b%e4%b8%8e%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86.md">33 lotop与lostat命令：聊聊命令背后的故事与工作原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 文件仓库：初识文件与文件系统.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/34%20%e6%96%87%e4%bb%b6%e4%bb%93%e5%ba%93%ef%bc%9a%e5%88%9d%e8%af%86%e6%96%87%e4%bb%b6%e4%b8%8e%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f.md">34 文件仓库：初识文件与文件系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 Linux文件系统（一）：Linux如何存放文件？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/35%20Linux%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aLinux%e5%a6%82%e4%bd%95%e5%ad%98%e6%94%be%e6%96%87%e4%bb%b6%ef%bc%9f.md">35 Linux文件系统（一）：Linux如何存放文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Linux文件系统（二）：Linux如何存放文件？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/36%20Linux%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aLinux%e5%a6%82%e4%bd%95%e5%ad%98%e6%94%be%e6%96%87%e4%bb%b6%ef%bc%9f.md">36 Linux文件系统（二）：Linux如何存放文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 浏览器原理（一）：浏览器为什么要用多进程模型？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/37%20%e6%b5%8f%e8%a7%88%e5%99%a8%e5%8e%9f%e7%90%86%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%b5%8f%e8%a7%88%e5%99%a8%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e7%94%a8%e5%a4%9a%e8%bf%9b%e7%a8%8b%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">37 浏览器原理（一）：浏览器为什么要用多进程模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 浏览器原理（二）：浏览器进程通信与网络渲染详解.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/38%20%e6%b5%8f%e8%a7%88%e5%99%a8%e5%8e%9f%e7%90%86%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%b5%8f%e8%a7%88%e5%99%a8%e8%bf%9b%e7%a8%8b%e9%80%9a%e4%bf%a1%e4%b8%8e%e7%bd%91%e7%bb%9c%e6%b8%b2%e6%9f%93%e8%af%a6%e8%a7%a3.md">38 浏览器原理（二）：浏览器进程通信与网络渲染详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 源码解读：V8 执行 JS 代码的全过程.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/39%20%e6%ba%90%e7%a0%81%e8%a7%a3%e8%af%bb%ef%bc%9aV8%20%e6%89%a7%e8%a1%8c%20JS%20%e4%bb%a3%e7%a0%81%e7%9a%84%e5%85%a8%e8%bf%87%e7%a8%8b.md">39 源码解读：V8 执行 JS 代码的全过程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 内功心法（一）：内核和后端通用的设计思想有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/40%20%e5%86%85%e5%8a%9f%e5%bf%83%e6%b3%95%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%86%85%e6%a0%b8%e5%92%8c%e5%90%8e%e7%ab%af%e9%80%9a%e7%94%a8%e7%9a%84%e8%ae%be%e8%ae%a1%e6%80%9d%e6%83%b3%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">40 内功心法（一）：内核和后端通用的设计思想有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 内功心法（二）：内核和后端通用的设计思想有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/41%20%e5%86%85%e5%8a%9f%e5%bf%83%e6%b3%95%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%86%85%e6%a0%b8%e5%92%8c%e5%90%8e%e7%ab%af%e9%80%9a%e7%94%a8%e7%9a%84%e8%ae%be%e8%ae%a1%e6%80%9d%e6%83%b3%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">41 内功心法（二）：内核和后端通用的设计思想有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 性能调优：性能调优工具eBPF和调优方法.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/42%20%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%ef%bc%9a%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%b7%a5%e5%85%b7eBPF%e5%92%8c%e8%b0%83%e4%bc%98%e6%96%b9%e6%b3%95.md">42 性能调优：性能调优工具eBPF和调优方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="先睹为快：迷你CPU项目效果演示.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%85%88%e7%9d%b9%e4%b8%ba%e5%bf%ab%ef%bc%9a%e8%bf%b7%e4%bd%a0CPU%e9%a1%b9%e7%9b%ae%e6%95%88%e6%9e%9c%e6%bc%94%e7%a4%ba.md">先睹为快：迷你CPU项目效果演示.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐01 云计算基础：自己动手搭建一款IAAS虚拟化平台.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%9001%20%e4%ba%91%e8%ae%a1%e7%ae%97%e5%9f%ba%e7%a1%80%ef%bc%9a%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%e6%90%ad%e5%bb%ba%e4%b8%80%e6%ac%beIAAS%e8%99%9a%e6%8b%9f%e5%8c%96%e5%b9%b3%e5%8f%b0.md">加餐01 云计算基础：自己动手搭建一款IAAS虚拟化平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐02 学习攻略（一）：大数据&amp;云计算，究竟怎么学？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%9002%20%e5%ad%a6%e4%b9%a0%e6%94%bb%e7%95%a5%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a4%a7%e6%95%b0%e6%8d%ae&amp;%e4%ba%91%e8%ae%a1%e7%ae%97%ef%bc%8c%e7%a9%b6%e7%ab%9f%e6%80%8e%e4%b9%88%e5%ad%a6%ef%bc%9f.md">加餐02 学习攻略（一）：大数据&amp;云计算，究竟怎么学？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐03 学习攻略（二）：大数据&amp;云计算，究竟怎么学？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%9003%20%e5%ad%a6%e4%b9%a0%e6%94%bb%e7%95%a5%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a4%a7%e6%95%b0%e6%8d%ae&amp;%e4%ba%91%e8%ae%a1%e7%ae%97%ef%bc%8c%e7%a9%b6%e7%ab%9f%e6%80%8e%e4%b9%88%e5%ad%a6%ef%bc%9f.md">加餐03 学习攻略（二）：大数据&amp;云计算，究竟怎么学？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐04 谈谈容器云与和CaaS平台.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%9004%20%e8%b0%88%e8%b0%88%e5%ae%b9%e5%99%a8%e4%ba%91%e4%b8%8e%e5%92%8cCaaS%e5%b9%b3%e5%8f%b0.md">加餐04 谈谈容器云与和CaaS平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐05 分布式微服务与智能SaaS.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%8a%a0%e9%a4%9005%20%e5%88%86%e5%b8%83%e5%bc%8f%e5%be%ae%e6%9c%8d%e5%8a%a1%e4%b8%8e%e6%99%ba%e8%83%bdSaaS.md">加餐05 分布式微服务与智能SaaS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="国庆策划01 知识挑战赛：检验一下学习成果吧！.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%9b%bd%e5%ba%86%e7%ad%96%e5%88%9201%20%e7%9f%a5%e8%af%86%e6%8c%91%e6%88%98%e8%b5%9b%ef%bc%9a%e6%a3%80%e9%aa%8c%e4%b8%80%e4%b8%8b%e5%ad%a6%e4%b9%a0%e6%88%90%e6%9e%9c%e5%90%a7%ef%bc%81.md">国庆策划01 知识挑战赛：检验一下学习成果吧！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="国庆策划02 来自课代表的学习锦囊.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%9b%bd%e5%ba%86%e7%ad%96%e5%88%9202%20%e6%9d%a5%e8%87%aa%e8%af%be%e4%bb%a3%e8%a1%a8%e7%9a%84%e5%ad%a6%e4%b9%a0%e9%94%a6%e5%9b%8a.md">国庆策划02 来自课代表的学习锦囊.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="国庆策划03 揭秘代码优化操作和栈保护机制.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e5%9b%bd%e5%ba%86%e7%ad%96%e5%88%9203%20%e6%8f%ad%e7%a7%98%e4%bb%a3%e7%a0%81%e4%bc%98%e5%8c%96%e6%93%8d%e4%bd%9c%e5%92%8c%e6%a0%88%e4%bf%9d%e6%8a%a4%e6%9c%ba%e5%88%b6.md">国庆策划03 揭秘代码优化操作和栈保护机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="温故知新 思考题参考答案（一）.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e6%b8%a9%e6%95%85%e7%9f%a5%e6%96%b0%20%e6%80%9d%e8%80%83%e9%a2%98%e5%8f%82%e8%80%83%e7%ad%94%e6%a1%88%ef%bc%88%e4%b8%80%ef%bc%89.md">温故知新 思考题参考答案（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 我是怎样学习Verilog的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%88%91%e6%98%af%e6%80%8e%e6%a0%b7%e5%ad%a6%e4%b9%a0Verilog%e7%9a%84%ef%bc%9f.md">用户故事 我是怎样学习Verilog的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 心若有所向往，何惧道阻且长.md" href="/%e4%b8%93%e6%a0%8f/%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%9f%ba%e7%a1%80%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%bf%83%e8%8b%a5%e6%9c%89%e6%89%80%e5%90%91%e5%be%80%ef%bc%8c%e4%bd%95%e6%83%a7%e9%81%93%e9%98%bb%e4%b8%94%e9%95%bf.md">结束语 心若有所向往，何惧道阻且长.md</a>
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
                            <h1 id="title" data-id="21 RISC-V指令精讲（六）：加载指令实现与调试" class="title">21 RISC-V指令精讲（六）：加载指令实现与调试</h1>
                            <div><p>你好，我是LMOS。</p>

<p>之前我们已经学过了RISC-V中的算术指令、逻辑指令、原子指令。这些指令主要的操作对象是寄存器，即对寄存器中的数据进行加工，这是RISC体系的重要特性。</p>

<p>但你是否想过寄存器中的数据从哪里来呢？答案是从内存中来，经过存储指令加载到寄存器当中。</p>

<p>RISC-V是一个典型的加载储存体系结构，这种体系类型的CPU，只有加载与储存指令可以访问内存，运算指令不能访问内存。这节课我们就来学习一下RISC-V的加载指令。</p>

<p>顾名思义，加载指令就是从一个地址指向的内存单元中，加载数据到一个寄存器中。根据加载数据大小和类型的不同，加载指令还可以细分成五条加载指令，分别是加载字节指令、无符号加载字节指令、加载半字指令、无符号加载半字指令、加载字指令。</p>

<p>这节课的代码，你可以从<a href="https://gitee.com/lmos/Geek-time-computer-foundation/tree/master/lesson21~22" target="_blank">这里</a>下载。</p>

<h2 id="加载字节指令-lb指令">加载字节指令：lb指令</h2>

<p>我们先从加载字节指令开始说起。在研究加载字节指令之前，我们先来看看RISC-V的加载指令的格式，其对应的汇编语句格式如下：</p>

<pre><code>指令助记符 目标寄存器，源操作数2(源操作数1）
</code></pre>

<p>对于加载指令，指令助记符可以是lb、lbu、lh、lhu、lw，目标寄存器可以是任何通用寄存器，源操作数1也可以是任何通用寄存器，源操作数2则是立即数。</p>

<p>我们用汇编代码来描述一下加载字节指令，形式如下：</p>

<pre><code>lb rd,imm(rs1)
#lb 加载字节指令
#rd 目标寄存器
#rs1 源寄存器
#imm 立即数（-2048~2047）
</code></pre>

<p>上述代码中rd和rs1可以是任何通用寄存器。立即数imm为12位二进制数据，其范围是-2048~2047，前面课程已经说明了，RISC-V指令集中所有的立即数都是有符号数据，这里的imm在其他的文档里也称为偏移量，为了一致性，我们继续沿用立即数的叫法。</p>

<p>lb指令完成的操作，用伪代码描述如下所示：</p>

<pre><code>rd = 符号扩展（[rs1+imm][7:0]）
</code></pre>

<p>我来为你解释一下，上面的伪代码执行的操作是怎样的。</p>

<p>首先lb指令会从内存单元里rs1+imm这个地址里取得8位数据，也就是第0位到第7位的数据。然后，把这个数据进行符号扩展，扩展成32位数据。如果符号位为1，则该32位的高24位为1，否则为0。最后lb指令再把这个32位的数据赋给rd。</p>

<p>下面我们一起写代码验证一下。为了方便之后的调试，我们需要先设计好代码的组织结构，这个过程前面几节课我们反复做过，现在估计你已经相当熟练了。首先创建main.c文件并在上面写好main函数。然后写一个load.S文件，用汇编写上lb_ins函数。</p>

<p>lb_ins函数的代码如下所示：</p>

<pre><code>.text
.globl lb_ins
#a0内存地址
#a0返回值
lb_ins:
    lb a0, 0(a0)       #加载a0+0地址处的字节到a0中
    jr ra              #返回
</code></pre>

<p>对照代码，我们可以看到这个函数只有两条指令，第一条指令把a0+0地址处的字节加载到a0中，第二条指令就是返回指令，a0作为函数的返回值返回。</p>

<p>你可以用VSCode打开工程目录，按下“F5”键调试一下。首先，我们把断点停在lb a0，0(a0) 指令处，如下所示：</p>

<p><img src="assets/2d988c29026a922b85000c442e5d7c73.jpg" alt="图片" /></p>

<p>上图中是刚刚执行完lb a0，0(a0)指令之后，执行jr ra指令之前的状态。</p>

<p>我们可以看到，a0寄存器中的值已经变成了0xfffffffb，我们继续单步调试返回到main函数中执行printf函数，打印一下lb_ins函数返回的结果，如下图所示：</p>

<p><img src="assets/891a34aa8f3f81d499f8ca1465a47f09.jpg" alt="图片" /></p>

<p>如上图所示，byte变量的值为-5，其补码为0xfb，我们把byte的地址强制为无符号，整体传给lb_ins函数。</p>

<p>调用规范告诉我们，C语言函数用a0寄存器传递第一个参数。lb指令虽然只加载了内存地址处的8位数据（0xfb），但是它会用数据的符号位把数据扩展成32位（0xfffffffb），再传给目标寄存器，即a0寄存器，这样a0就会作为返回值返回，所以结果为0xfffffffb。这证明了lb指令工作是正常的。</p>

<h2 id="无符号加载字节指令-lbu指令">无符号加载字节指令：lbu指令</h2>

<p>接着我们来看一看lb指令的另一个版本，就是无符号加载字节指令，它的汇编代码是这样写的：</p>

<pre><code>lbu rd,imm(rs1)
#lbu 无符号加载字节指令
#rd 目标寄存器
#rs1 源寄存器
#imm 立即数（-2048~2047）
</code></pre>

<p>上述代码里，rd，rs1，imm与lb指令的用法和规则是一样的。lbu指令完成的操作，我们用伪代码描述如下：</p>

<pre><code>rd = 符号扩展（[rs1+imm][7:0]）
</code></pre>

<p>因为lbu指令获取8位数据的位置，还有把数据扩展成32位赋给rd的过程，都和lb指令一样，我就不重复了。<strong>注意是无符号扩展，即符号位为0。</strong></p>

<p>接下来咱们写个代码验证一下，同样在load.S文件中用汇编写上lbu_ins函数 ，代码如下所示：</p>

<pre><code>.globl lbu_ins
#a0内存地址
#a0返回值
lbu_ins:
    lbu a0, 0(a0)      #加载a0+0地址处的字节到a0中
    jr ra              #返回
</code></pre>

<p>在lbu_ins函数中，第一条指令把a0+0地址处的字节加载到a0中，之后a0会作为函数的返回值返回。</p>

<p>同样地，用VSCode打开工程目录，这里我们需要在lbu a0，0(a0) 指令处打下断点，随后按下“F5”进行调试，如下所示：</p>

<p><img src="assets/00a95a306fbb62873838810c57c0e6a7.jpg" alt="图片" /></p>

<p>上图中是执行完lbu a0，0(a0)指令之后，执行jr ra指令之前的状态，现在a0寄存器中的值已经变成了0xfb。</p>

<p>我们继续单步调试，返回到main函数中，让printf函数打印lbu_ins函数，返回的结果如下图所示：</p>

<p><img src="assets/61ebb72785926ed2ecbf9cb98c4c5808.jpg" alt="图片" /></p>

<p>同样的，byte变量的值为-5，其补码为0xfb，我们把byte变量的地址强制为无符号整体传给lbu_ins函数并调用它。</p>

<p>在lbu_ins函数中，lbu指令只加载内存地址处的8位数据（(0xfb），但是它与lb指令不同，它会用0把数据扩展成32位（0x000000fb），再传给目标寄存器，即a0寄存器。这样a0就会作为返回值返回，故而result为0xfb（251）。这证明了lbu指令是正常工作的。lbu指令的这种<strong>无符号扩展特性</strong>，非常易于处理无符号类型的变量。</p>

<h2 id="加载半字指令-lh指令">加载半字指令：lh指令</h2>

<p>有了能够加载一个字节的指令，我们还需要加载双字节的指令，也叫加载半字指令。在RISC-V规范中，一个字是四字节，所以两个字节也称为半字。</p>

<p>下面我们一起来学习加载半字指令。我们还是先从汇编代码的书写形式来熟悉它，如下所示：</p>

<pre><code>lh rd,imm(rs1)
#lh 加载半字指令
#rd 目标寄存器
#rs1 源寄存器
#imm 立即数（-2048~2047）
</code></pre>

<p>上述代码中，rd和rs1可以是任何通用寄存器。立即数imm为12位二进制数据，其范围是-2048~2047。</p>

<p>lh指令完成的操作，用伪代码描述如下：</p>

<pre><code>rd = 符号扩展（[rs1+imm][15:0]）
</code></pre>

<p>经过前面的学习，相信你已经找到了规律，现在自己也能解读这样的伪代码了。还是熟悉的过程：先读取数据，找到内存单元rs1+imm这个地址，从里面获取第0位到第15位的数据，再对这个16位数据进行符号扩展（扩展为32位数据）；接着根据符号位分情况处理，如果符号位为1则该32位的高16位为1，否则为0；最后把这个32位数据赋值给rd。</p>

<p>下面是写代码验证时间。我们在load.S文件中用汇编写上lh_ins函数，代码如下：</p>

<pre><code>.globl lh_ins
#a0内存地址
#a0返回值
lh_ins:
    lh a0, 0(a0)       #加载a0+0地址处的半字到a0中
    jr ra              #返回
</code></pre>

<p>上面的lh_ins函数中，第一条指令把a0+0地址处的半字加载到a0中，而a0将会作为函数的返回值返回。</p>

<p>我们用VSCode打开工程目录，在lh a0，0(a0) 指令处打下断点，随后按下“F5”键调试一下，如下所示：</p>

<p><img src="assets/f46710136ab52368c7b172976b9yy42e.jpg" alt="图片" /></p>

<p>上图中是执行完lh a0，0(a0)指令之后，执行jr ra指令之前的状态。从图中我们可以看到a0寄存器中的值已经变成了0xffffffff。</p>

<p>我们继续单步调试，返回到main函数中，让printf函数打印一下lh_ins函数返回的结果，如下图所示：</p>

<p><img src="assets/a293e6ca840c2b2269fac8f79633aef6.jpg" alt="图片" /></p>

<p>对照图片不难发现，short类型的half变量占用两个字节，其值为-1，它的补码为0xffff，我们把half的地址强制为无符号整体传给lh_ins函数。</p>

<p>在lh_ins函数中，lh指令虽然只加载内存地址处的16位数据（0xffff），但是它会用数据的符号位把数据扩展成32位（0xffffffff），再把扩展后的数据传递给a0寄存器，这样a0就会作为返回值返回，故而result为0xffffffff。这证明了lh指令工作正常。</p>

<h2 id="无符号加载半字指令-lhu指令">无符号加载半字指令：lhu指令</h2>

<p>加载半字指令也分为两种版本，即有符号版本和无符号版本。我们再看看无符号加载半字指令lhu，它的汇编代码书写形式如下所示。</p>

<pre><code>lhu rd,imm(rs1)
#lhu 无符号加载半字指令
#rd 目标寄存器
#rs1 源寄存器
#imm 立即数（-2048~2047）
</code></pre>

<p>上述代码中rd，rs1，imm与lh指令的用法和规则是一样的。</p>

<p>我用伪代码为你描述一下lhu指令完成的功能。</p>

<pre><code>rd = 符号扩展（[rs1+imm][15:0]）
</code></pre>

<p>lhu指令的操作过程与lh指令一样，我就不重复了，但符号位为0，lhu会进行无符号扩展，即数据的高16位为0。</p>

<p>接下来就是代码验证环节，我们同样在load.S文件中用汇编写上lhu_ins函数，代码如下所示：</p>

<pre><code>.globl lhu_ins
#a0内存地址
#a0返回值
lhu_ins:
    lhu a0, 0(a0)       #加载a0+0地址处的半字到a0中
    jr ra               #返回
</code></pre>

<p>可以看到上面的lhu_ins函数中，第一条指令会把a0+0地址处的字节加载到a0中，而a0将会作为函数的返回值返回。</p>

<p>我们用VSCode打开工程目录，在lhu a0，0(a0) 指令处打下断点，随后按“F5”键调试，如下所示：</p>

<p><img src="assets/10f28eced7dca150e1c64cb33967fb13.jpg" alt="图片" /></p>

<p>上图是执行完lhu a0，0(a0)指令之后，执行jr ra指令之前的状态，可以看到a0寄存器中的值已经变成了0xffff。</p>

<p>我们继续单步调试，返回到main函数中，让printf函数打印一下lhu_ins函数返回的结果，如下图所示：</p>

<p><img src="assets/8666b703ceffce261aefffa971c07574.jpg" alt="图片" /></p>

<p>如上图所示，我们把half的地址强制为无符号整体传给lhu_ins函数。在lhu_ins函数中，lh指令虽然只加载内存地址处的16位数据（0xffff），但是它会用数据的符号位把数据扩展成32位（0x0000ffff），给a0寄存器作为返回值返回，故而result为0xffff，也就是65535。这证明了lhu指令工作正常。与lbu指令一样，这里同样是为了让编译器方便处理无符号类型的变量。</p>

<h2 id="加载字指令-lw指令">加载字指令：lw指令</h2>

<p>对于一款处理器来说，最常用的是加载其自身位宽的数据为32位的RISC-V处理器，加载字指令是非常常用且必要的指令，一个字的储存大小通常和处理器位宽相等。</p>

<p>现在。我们一起来学习最后一条加载指令，即加载字指令。我们先来看看加载字指令lw，它的汇编代码书写形式如下：</p>

<pre><code>lw rd,imm(rs1)
#lw 加载字指令
#rd 目标寄存器
#rs1 源寄存器
#imm 立即数（-2048~2047）
</code></pre>

<p>lw指令完成的操作，用伪代码描述是这样的：</p>

<pre><code>rd = （[rs1+imm][31:0]）
</code></pre>

<p>我们看看上面的伪代码执行的操作。首先找到内存单元rs1+imm这个地址，从里面获取第0位到第31位的数据，注意数据无需进行符号扩展，最后把这个32位数据赋值给rd。</p>

<p>写代码验证的思路，现在你应该也很熟悉了。同样还是在load.S文件中，用汇编写上lw_ins函数，代码如下所示：</p>

<pre><code>.globl lw_ins
#a0内存地址
#a0返回值
lw_ins:
    lw a0, 0(a0)        #加载a0+0地址处的字到a0中
    jr ra               #返回
</code></pre>

<p>我们可以看到，lw_ins函数完成的操作就是，先把a0+0地址处的一个字加载到a0中，再把a0作为函数的返回值返回。</p>

<p>用VSCode打开工程目录，在lw a0，0(a0) 指令处打下断点，随后按下“F5”键调试，调试截图如下所示：</p>

<p><img src="assets/c8a5868f94e54ed58def8417e9a65fa0.jpg" alt="图片" /></p>

<p>上图中是执行完lw a0，0(a0)指令之后，执行jr ra指令之前的状态，现在a0寄存器中的值已经变成了0xffffffff。继续单步调试执行，就可以返回到main函数中。</p>

<p>我们通过printf函数打印一下lw_ins函数返回的结果，如下图所示：</p>

<p><img src="assets/f8e0e3830db13c00a6e7609917ee56c3.jpg" alt="图片" /></p>

<p>这里我们把word的地址强制为无符号整体传给lw_ins函数。在lw_ins函数中，lw指令会直接加载内存地址处的32位数据(0xffffffff)，给a0寄存器作为返回值返回，result值为0xffffffff，但因为它是有符号类型，故而0xffffffff表示为-1。而word为无符号整形，0xffffffff则表示为4294967295，这证明了lw指令功能是正确无误的。</p>

<p>到这里，我们已经完成了对lb、lbu、lh、lhu、lw这五条指令的调试，也熟悉了它们的功能细节。现在我们继续研究一下lb_ins、lbu_ins、lh_ins、lhu_ins、lw_ins函数的二进制数据。</p>

<p>你只需要打开终端，切换到该工程目录下，输入命令：riscv64-unknown-elf-objdump -d ./main.elf &gt; ./main.ins，就会得到main.elf的反汇编数据文件main.ins。接着，我们打开这个文件，就会看到上述函数的二进制数据，如下所示：</p>

<p><img src="assets/aedebe83c009ab9a255c7e7b6e9e3f01.jpg" alt="图片" /></p>

<p>上图反汇编代码中包括伪指令和两个字节的压缩指令。比如ret的机器码是0x8082，lw a0,0(a0)机器码是0x4108，它们只占用16位编码，即二字节。截图里五条加载指令的机器码与指令的对应关系，你可以参考后面这张表格。</p>

<p><img src="assets/4fba716e94cc23548af53e1373eec202.jpg" alt="图片" /></p>

<p>下面我们继续一起拆分一下lb、lbu、lh、lhu、lw指令的各位段的数据，看看它们都是如何编码的。如下图所示：</p>

<p><img src="assets/23d68b684c69fe5yy96e418231571e1c.jpg" alt="图片" /></p>

<p>对照上图可以看到，lb、lbu、lh、lhu、lw指令的功能码都不一样，我们可以借此区分这些指令。而这些加载指令的操作码都一样，立即数也相同（都是0），这和我们编写的代码有关。</p>

<p>需要注意的是<strong>lw a0,0(a0)指令</strong>，上图的情况和反汇编出来的数据可能不一致，这是因为<strong>编译器使用了压缩指令</strong>。</p>

<p>我还原了lw a0,0(a0)正常的编码，你可以手动在lw_ins函数中插入这个数据0x00052503，进行验证。怎么插入这个数据使之变成一条指令呢？代码如下所示：</p>

<pre><code>.globl lw_ins
#a0内存地址
#a0返回值
lw_ins:
    .word 0x00052503    #lw a0, 0(a0)        #加载a0+0地址处的字到a0中
    jr ra               #返回
</code></pre>

<h2 id="重点回顾">重点回顾</h2>

<p>今天我们一共学习了五条加载指令，分别是加载字节指令、无符号加载字节指令、加载半字指令、无符号加载半字指令、加载字指令，它们可以加载不同大小的数据，同时又能处理数据的符号。</p>

<p>而且这五条指令组合起来，既可以加载不同位宽的数据，又能处理加载有、无符号的数据。这些指令为高级语言实现有无符号的类型变量提供了基础，让我们的开发工作更便利。比方说，在C语言中，实现的各种数据类型：unsigned、int、char、unsigned、char等都离不开加载指令。</p>

<p>最后我给你总结了一张导图，供你参考复习。下节课，我们继续学习储存指令，敬请期待。</p>

<p><img src="assets/6ef7fbfcef0ecc577b558d63370fa3b5.jpg" alt="" /></p>

<h2 id="思考题">思考题</h2>

<p>为什么加载字节与加载半字指令，需要处理数据符号问题，而加载字指令却不需要呢？</p>

<p>欢迎你在留言区跟我交流，也推荐你把这节课分享给更多同事、朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#59353535606d6868696e193e34383035773a3634" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f18623aeea4ecfb',t:'MTczNDExODM4NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
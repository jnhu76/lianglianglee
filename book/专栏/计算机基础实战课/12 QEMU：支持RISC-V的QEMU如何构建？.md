<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=12&#32;QEMU：支持RISC-V的QEMU如何构建？>
        <link rel="icon" href="/static/favicon.png">
        <title>12 QEMU：支持RISC-V的QEMU如何构建？ </title>
        
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
                            <h1 id="title" data-id="12 QEMU：支持RISC-V的QEMU如何构建？" class="title">12 QEMU：支持RISC-V的QEMU如何构建？</h1>
                            <div><p>你好，我是LMOS。</p>

<p>工欲善其事，必先利其器。作为开发者，学习过程中我们尤其要重视动手实践，不断巩固和验证自己学到的知识点。而动手实践的前提，就是要建立一个开发环境，这个环境具体包括编译环境、执行环境，以及各种常用的工具软件。</p>

<p>我会用两节课带你动手搭好环境，今天这节课咱们先热个身，搞清楚什么是主环境，还有怎么基于它生成交叉编译工具。</p>

<p>代码你可以从<a href="https://gitee.com/lmos/Geek-time-computer-foundation/tree/master/lesson12~13" target="_blank">这里</a>下载。</p>

<h2 id="主环境">主环境</h2>

<p>主环境，有时也叫作HOST环境，也就是我们使用的计算机环境，即使用什么样的操作系统、什么架构的计算机作为开发环境。</p>

<p>比方说我们经常用PC机作为开发机使用，它实际就是一个基于x86架构（或其他架构）的硬件平台，再加上Windows或者Linux等操作系统共同组成的开发环境。</p>

<p>普通用户的电脑上经常安装的操作系统是Windows，因为界面友好方便、操作简单且娱乐影音、游戏办公等应用软件也是不胜枚举。</p>

<p>Windows对普通用户来说的确非常友好。但是作为软件开发者，对于志存高远、想要精研技术的我们而言，更喜欢用的是Linux系统。</p>

<p>它虽然没有漂亮的GUI，却暴露了更多的计算机底层接口，也生产了更多的开发工具和各种各样的工具软件。比如大名鼎鼎的编译器GCC、声名远扬的编辑器EMACS、VIM，还有自动化的脚本工具shell、make等。这些工具对开发者非常友好，配合使用可以让我们的工作事半功倍，后面你会逐渐体会到这点。</p>

<p>当然Linux只是一个内核，我们不能直接使用，还需要各种工具、库和桌面GUI，把这些和Linux打包在一起发行，这就构成了我们常说的Linux发行版。</p>

<p>我最喜欢的Linux发行版是Deepin和Ubuntu。为了统一，我建议你使用Deepin最新版，你也可以使用Ubuntu，它们是差不多的。只是操作界面稍有不同。我先给你展示下我的Deepin，如下图，刚装上它的时候，我就觉得它颇为惊艳。</p>

<p><img src="assets/782c1d620430ba46b584b0cdbd86dc28.jpg" alt="图片" /></p>

<p>这里最基础的安装我就不讲了，因为安装Deepin十分简单，无论是虚拟机还是在物理机上安装，我相信你通过互联网都可以自行解决，搞不定也可以看看<a href="https://www.deepin.org/zh/installation/" target="_blank">这里</a>。</p>

<p>这两种方式我也替你对比过，虚拟机中的Linux较物理机上的Linux性能稍差一点，但并不影响我们实验操作和结果。</p>

<h2 id="为什么需要交叉编译">为什么需要交叉编译</h2>

<p>虽然主环境搞定了，但现在我们还不能直接跑代码。为什么呢？</p>

<p>先回想一下，平时我们正常开发软件需要什么？我猜，哪怕你不能抢答，也会知道个大概：需要电脑（PC）、特定的操作系统（比如Windows或Linux等），在这个操作系统上还能运行相应的编辑器和编译器。编辑器用来编写源代码，而编译器用来把源代码编译成可执行程序。</p>

<p>似乎不需要更多东西了，毕竟我们日常开发的软件，宿主平台和目标平台是相同的。如果我们把限制条件变一变，情况就不同了。如果我们想尝试在RISC-V平台上跑程序，要怎么办呢？</p>

<p>你或许会说，这简单，买一台RISC-V的机器不就行了。可是先不说购买硬件的经济成本，实际上，很多RISC-V平台硬件资源（如内存、SD卡容量）有限，不足以运行复杂的编译器软件，有的甚至没有操作系统，更别说在上面运行编译器或者编辑器软件了。</p>

<p>面对这样的困境，就要用到<strong>交叉编译</strong>了。什么是交叉编译呢？简单来说，就是在一个硬件平台上，生成另一个硬件平台的可执行程序。</p>

<p>举个例子，我们在x86平台上编译生成ARM平台的可执行程序；再比如说，之后的课里我们将在x86平台上，生成RISC-V的可执行程序。这些都属于交叉编译，在这个过程中编译生成可执行程序的平台，称为宿主机或者主机；执行特定程序的平台（如ARM或者RISC-V平台），称为目标机。</p>

<p>我特意准备了图解，为你展示在x86平台上，交叉编译生成RISC-V平台可执行程序的过程，你可以仔细看看：</p>

<p><img src="assets/a5fb7d81fa49da02bfeee45c35aef255.jpg" alt="图片" /></p>

<h2 id="如何构建risc-v交叉编译器">如何构建RISC-V交叉编译器</h2>

<p>前面说了交叉编译的本质就是生成其他平台体系上的可执行程序，这个体系又不同于我们宿主平台。我们的目的很简单，就是要<strong>在x86平台上编写源代码，然后编译出RISC-V平台的可执行程序，最后放在RISC-V平台上去运行</strong>。</p>

<p>因此，我们需要用宿主机编译器A，编译出一个编译器B，这个编译器B是本地平台上的可执行程序。</p>

<p>说得再具体点，你可以把编译器B看作是 <strong>x86 Linux上的一个应用</strong>。但它的特殊之处就是，能根据源代码生成RISC-V平台上的可执行程序。补充一句，这里的编译器A和B都是C语言编译器。</p>

<p>下面我们开始构造编译器B。编译器B不仅仅是C语言编译器，还有很多额外的程序。比如RISC-V平台上使用的二进制文件分析（objcopy）、反汇编（objdump）、elf结构分析工具（readelf）、静态库归档（ar）、汇编器（as）、链接器（ld）、GDB、C语言库（Newlib、Glib、Multlib）等。</p>

<p>为了简单、便于区分，我们把这些对应于RISC-V平台的编译器相关的软件，统称为 <strong>RISC-V工具链</strong>。</p>

<p>构建RISC-V工具链的主要步骤如下：</p>

<ol>
<li>安装依赖工具：在宿主平台上安装编译器A，以及相应的工具和库。-</li>
<li>下载RISC-V工具链的源代码；-</li>
<li>配置RISC-V工具链；-</li>
<li>编译RISC-V工具链，并安装在宿主平台上。</li>
</ol>

<h3 id="第一步-安装依赖工具">第一步：安装依赖工具</h3>

<p>我们先从第一步开始，编译器A主要是宿主平台上的GCC，工具主要是Make、Git、Autoconf、Automake、CURL、Python3、Bison、Flex等。这里GCC主要在build-essential包中，我们只要在Linux终端中输入如下指令就可以了：</p>

<pre><code>sudo apt-get install git autoconf automake autotools-dev curl python3 libmpc-dev libmpfr-dev libgmp-dev gawk build-essential bison flex texinfo gperf patchutils bc libexpat-dev libglib2.0-dev ninja-build zlib1g-dev pkg-config libboost-all-dev libtool libssl-dev libpixman-1-dev libpython-dev virtualenv libmount-dev libsdl2-dev
</code></pre>

<p><img src="assets/778c6aea3fa8c112c732fc3d7b653f79.jpg" alt="图片" /></p>

<p>如果不出意外，这些工具和库会通过网络由Linux的apt包管理器，全自动地给你安装完毕。</p>

<h3 id="第二步-下载工具链源代码">第二步：下载工具链源代码</h3>

<p>接着进入第二步下载RISC-V工具链源代码。通常来说，我们只要用Git克隆一个riscv-gnu-toolchain仓库即可，其它的由riscv-gnu-toolchain仓库中的仓库子模块自动处理。</p>

<h4 id="手动配置环节">手动配置环节</h4>

<p>由于众所周知的网络原因，你可能连riscv-gnu-toolchain仓库都下载不下来，更别说自动下载仓库子模块了。为了照顾卡壳的人，我把手动处理的情况也顺便讲一下，能够直接自动安装的同学可以跳过这部分，直接翻到7条指令之后的最终截图对一下结果就行。</p>

<p>子模块如下：</p>

<pre><code>riscv-qemu（虚拟机）
riscv-newlib (用嵌入式的轻量级C库)
riscv-binutils(包含一些二进制工具集合，如objcopy等)
riscv-gdb(用于调试代码的调试器)
riscv-dejagnu(用于测试其它程序的框架)
riscv-glibc(GNU的C库)
riscv-gcc (C语言编译器)
</code></pre>

<p>这些子模块我们需要手动从Gitee网站上下载。下载前，我们先在终端上输入后面的指令，建立一个目录，并切换到该目录中：</p>

<pre><code>mkdir RISCV_TOOLS
cd RISCV_TOOLS
</code></pre>

<p>把RISC-V工具链的源代码手动下载好，步骤稍微多了一些，我在后面分步骤列出，方便你跟上节奏。</p>

<p>其实也就是7条指令的事儿，并不复杂。先统一说明下，<strong>后面这些命令都是切换到riscv-gnu-toolchain目录的终端下</strong>，输入我给你列出的指令即可。</p>

<ol>
<li>开始下载riscv-gnu-toolchain，命令如下：</li>
</ol>

<pre><code class="language-bash">    git clone https://gitee.com/mirrors/riscv-gnu-toolchain
    cd riscv-gnu-toolchain
</code></pre>

<ol>
<li>下载RISC-V平台的C语言编译器源代码仓库，输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone -b riscv-gcc-10.2.0 https://gitee.com/mirrors/riscv-gcc
</code></pre>

<ol>
<li>下载测试框架源代码仓库，即riscv-dejagnu。输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone https://gitee.com/mirrors/riscv-dejagnu
</code></pre>

<ol>
<li>下载GNU的C库源代码仓库，也就是riscv-glibc，输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone -b riscv-glibc-2.29 https://gitee.com/mirrors/riscv-glibc
</code></pre>

<ol>
<li>下载用于嵌入式的轻量级C库源代码仓库，即riscv-newlib。输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone https://gitee.com/mirrors/riscv-newlib
</code></pre>

<ol>
<li>下载二进制工具集合源代码仓库riscv-binutils，输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone -b riscv-binutils-2.35 https://gitee.com/mirrors/riscv-binutils-gdb riscv-binutils
</code></pre>

<ol>
<li>最后，下载GDB软件调试器源代码仓库riscv-gdb，输入如下指令：</li>
</ol>

<pre><code class="language-bash">    git clone -b fsf-gdb-10.1-with-sim https://gitee.com/mirrors/riscv-binutils-gdb riscv-gdb
</code></pre>

<p>好，现在所有的RISC-V工具链的源代码，我们已经下载完了。我们一起来同步一下，确保你我的riscv-gnu-toolchain目录下的目录和文件，完全一致。</p>

<p>在riscv-gnu-toolchain目录的终端下输入ls指令，你应该得到和后面这张图一样的结果。</p>

<p><img src="assets/e8307c6cabf00a3fccc4a0e0a93b3034.jpg" alt="图片" /></p>

<h3 id="第三步-配置工具链">第三步：配置工具链</h3>

<p>在我们用宿主编译器编译所有的RISC-V工具链的源代码之前，还有最重要的一步，那就是配置RISC-V工具链的功能。</p>

<p>RISC-V工具链有很多配置选项，不同的配置操作会生成具有特定功能的RISC-V工具链。此外，配置操作还有一个功能，就是检查编译RISC-V工具链所依赖的工具和库。检查通过，就会生成相应的配置选项文件，还有用于编译操作的Makefile文件。</p>

<p>下面我们开始配置操作。为了不污染源代码目录，我们可以在riscv-gnu-toolchain目录下建立一个build目录，用于存放编译RISC-V工具链所产生的文件。还是在切换到riscv-gnu-toolchain目录的终端下，输入如下指令：</p>

<pre><code class="language-bash">    mkdir build  #建立build目录
    #配置操作，终端一定要切换到build目录下再执行如下指令
    ../configure --prefix=/opt/riscv/gcc --enable-multilib --target=riscv64-multlib-elf
</code></pre>

<p>我给你解释一下指令里的关键内容。</p>

<p>–prefix表示RISC-V的工具链的安装目录，我们一起约定为“/opt/riscv/gcc”这个目录。</p>

<p>–enable-multilib表示使用multlib库，使用该库编译出的RISC-V工具链，既可以生成RISCV32的可执行程序，也可以生成RISCV64的可执行程序，而默认的Newlib库则不行，它只能生成RISCV（32/64）其中之一的可执行程序。</p>

<p>–target表示生成的RISC-V工具链中，软件名称的前缀是riscv64-multlib-elf-xxxx。若配置操作执行成功了，build目录中会出现如下所示的文件：</p>

<p><img src="assets/86a3b65f445394f1e8be61e3470e9f5e.jpg" alt="图片" /></p>

<h3 id="第四步-编译工具链">第四步：编译工具链</h3>

<p>最后我们来完成第四步，编译RISC-V工具链。只要配置操作成功了，就已经成功了90%。其实编译操作是简单且高度自动化的，我们只要在切换到build目录的终端下，输入如下指令即可：</p>

<pre><code>sudo make -j8
</code></pre>

<p>这个指令在编译完成后会自动安装到“/opt/riscv/gcc”目录，由于要操作“/opt/riscv/gcc”目录需要超级管理员权限，所以我们要记得加上sudo。</p>

<p>另外，如果你的宿主机的CPU有n个核心，就在make 后面加-j（n*2），这样才能使用多线程加速编译。</p>

<p>好了，一通操作猛如虎，现在最重要的事情是等待计算机“搬砖”了。你不妨播放音乐，泡上一杯新鲜的热茶，一边听歌，一边喝茶……估计要喝很多杯茶，才会编译完成。最最重要的是这期间不能断电，否则几个小时就白费了。</p>

<p>如果终端中不出现任何错误，就说明编译成功了。我们在终端中切换到“/opt/riscv/gcc/bin”目录下，执行如下指令：</p>

<pre><code>riscv64-unknown-elf-gcc -v
</code></pre>

<p>上述指令执行以后，会输出riscv64-unknown-elf-gcc的版本信息，这证明RISC-V工具链构建成功了。如下所示：</p>

<p><img src="assets/d62cf1b12d532aea1643118ea117febe.jpg" alt="图片" /></p>

<p>到这里，我们环境已经成功了一半，有了交叉编译器，并且这种交叉编译器能生成32位的RISC-V平台的可执行程序，也能生成64位的RISC-V平台的可执行程序。</p>

<p>你可能会好奇，成功了一半，那另一半呢？这需要我们接着干另一件事。什么事呢？容我先在这里卖个关子，下节课再揭秘。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>通过这节课的学习，我们成功构建了RISC-V工具链，这样就能在X86平台上生成RISC-V平台的可执行程序了。下面让我们一起回顾一下，这节课中都做了些什么。</p>

<p>我们首先约定了宿主环境，需要用到Ubuntu或者Deepin的Linux发行版，无论你是将它们安装在物理PC上，还是安装在虚拟机上。</p>

<p>然后我们了解了什么是交叉编译。为了方便后面课程学习动手实践，我们要在x86平台的宿主机上编译生成RISC-V平台的可执行程序。</p>

<p>明确了目标，我们一起动手开始构建了一个RISC-V交叉编译器。你会发现其中不只是C/C++编译器，还有很多处理二进制可执行程序的工具，我们把这些统称为RISC-V工具链。</p>

<h2 id="思考题">思考题</h2>

<p>请你说一说交叉编译的过程？</p>

<p>期待你再留言区分享自己的实验笔记，或者与我交流讨论。也推荐你把这节课分享给更多朋友，我们一起玩转交叉编译。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#711d1d1d48454040414631161c10181d5f121e1c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f185dfe3ed9ecfb',t:'MTczNDExODIxMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
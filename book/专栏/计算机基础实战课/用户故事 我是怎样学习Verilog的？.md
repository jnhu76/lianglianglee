<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=用户故事&#32;我是怎样学习Verilog的？>
        <link rel="icon" href="/static/favicon.png">
        <title>用户故事 我是怎样学习Verilog的？ </title>
        
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
                            <h1 id="title" data-id="用户故事 我是怎样学习Verilog的？" class="title">用户故事 我是怎样学习Verilog的？</h1>
                            <div><p>你好，我是咻咻咻。</p>

<p>先做个自我介绍，我是学 FPGA的 （写Verilog的）。今年参加了工信部组织的集创赛（<a href="http://univ.ciciec.com" target="_blank">全国大学生集成电路创新创业大赛</a>），取得了省级一等奖，国家级三等奖的成绩。</p>

<p>成绩还不错，但主要是团队能力比较强，我深知自己还有很大的提升空间。刚比赛归来，就听说 LMOS 老师开了一门新课，还是手撸 CPU 的，于是我便兴致冲冲地准备完成男人的两大梦想（手撸 CPU造高达）之一。</p>

<p>因为之前参赛时也用到了RISC-V，主要编程语言也是Verilog，回来继续学习的时候，比赛的熟悉感又扑面而来了。这次用户故事，我主要想和你聊聊我对 Verilog 和专栏学习的一些粗浅理解，希望对你有启发。</p>

<h2 id="我是怎样学verilog的">我是怎样学Verilog的？</h2>

<p>想必没有人是看完一整本谭的 C 语言再去写“hello，world!”的吧？</p>

<p>学Verilog也是一样，如果你<strong>从头开始</strong>学Verilog中的数据类型、符号常量、运算符，条件、分支、循环语句，过程结构……学完这些，再去看课里的配套代码，结果很可能是看每个词都<strong>模模糊糊有点印象，但连起来就一脸懵逼</strong>。而且这样做往往进度很慢，大多数人多半坚持不下去。</p>

<p>因此，在学习前，我们先要明确，在这门课里Verilog只是一种实现迷你CPU的<strong>工具</strong>，我们真正需要掌握的是<strong>这个迷你CPU的实现原理。</strong>Verilog 够用就行，不需要过多深究它的原理。</p>

<p>那我们学 Verilog 语言怎样上手才有效呢？如果一直在纸面理论徘徊，大概率会把自己绕晕。</p>

<p>所以，我的方法就是从简单的开始，多写写，代码写的多了就会了。比如说，阻塞与非阻塞想不通，不妨直接写程序看仿真，毕竟仿真波形更加清晰明了，这比跟理论说明的几行字较劲更管用。</p>

<p>如果一定这样你还是觉得Verilog还是太难，建议先刷题、补补基础语法（我用的参考资料列在了最后）。</p>

<h2 id="how-what-why三步代码学习法">How-What-Why三步代码学习法</h2>

<p>另外，有了比赛经历练习代码，再回看专栏突然觉得课程里代码量并没那么多，但非常规范和完善。</p>

<p>我试着按“<strong>How-What-Why</strong>”的三步走策略执行，感觉效果还不错，因此分享出来，给你提供一个思路。</p>

<p><img src="assets/94a16a30b8533c205576a1bd56fa58b4.jpg" alt="图片" /></p>

<p>简单说说我的做法。首先解决How的问题，具体就是把LMOS老师提供的例程 CTRL+C，CTRL+V，整个跑一遍。</p>

<p>跑通后，你就能看到仿真与通过Yosys生成的RTL图，这能帮你大概感知到整个程序是如何运转的，让你提升信心，消除陌生感的同时，也能<strong>建立一个大体的框架</strong>，更加明确各部分作用与相互关系，学起来也连贯。</p>

<p>接下来是What，也是核心部分。跑通了老师的代码后，我们就可以开始慢慢看程序详细代码，补足这部分配套代码相关的Verilog语法知识。我建议将程序分割成很多的小部分，一点一点去了解各部分是什么，程序中起什么作用。这样以<strong>问题为导向</strong>来理解知识，学起来也相对轻松。</p>

<p>最后就是理解Why的环节，此时我们需要回到文章中，结合课程尤其是老师给出的代码注释做理解。在理解过程中，别忘了结合coding去实践自己的想法。<strong>实践的时候总会有些与理想状况的差异，而这些差异往往可以促进我们更快成长。</strong></p>

<p>与Verilog语法相比，确实是完全理解CPU内部的思想和工作原理更难，我对完全理解的定义是，脑子里一想到这个知识点就可以熟练拓展开，明白它大致的作用，涉及到的原理，为什么这样写。</p>

<p>遇到想不通的地方，建议暂时先记录下来，相信随着你课程看得多了，有些问题会迎刃而解。如果实在不懂的，也可以留言提问。</p>

<p>当然，这样做往往会花费你大量的时间，对于一个零基础的人来说，可能理解完二十行代码（前提：每行代码都有截然不同的作用），四五个小时就过去了。但我们刚起步，这也很正常，不用认为自己理解力差，不适合这门课。</p>

<p>如此来回往复不断学习，Verilog语言你也会不知不觉学个七七八八。<strong>更重要的是，你将掌握构建迷你CPU的核心内容，不仅能明白 CPU的基本原理，也会获得如何手写一个CPU的实操经验。</strong></p>

<h2 id="什么是以问题为导向">什么是以问题为导向？</h2>

<p>在学习课程的时候，“以问题为导向”这个原则对我很有帮助。这里最重要的是<strong>明确自己的目的</strong>，通过不断给自己提出问题，使自己专注，提高效率。</p>

<p>专栏每节课的内容信息密度都很高，硬着头皮从头看到尾，很容易让自己信息“过载”了。</p>

<p>我个人会对一篇内容进行分块学习，先借助思维导图/重点回顾了解整节课的结构/重点，知道主线是什么，了解大体框架，再选自己最有兴趣的部分上手，或者自己给自己提几个问题，带着问题细读课程。</p>

<p><strong>只有当你自己感兴趣、有问题，想不明白又特别好奇</strong>的时候，才会充分调动自己的大脑去思考，拼尽全力去尝试解决问题。这个过程中，你并不是被动输入，而是主动地加工理解。最终把一个问题弄明白的时候，也会非常有成就感。</p>

<p>对此，我有两个建议。第一点，先想清楚自己想要找什么，再去搜寻答案。如果有问题，但不明确搜寻方向，可以从这几个方面下手去提问：是什么、有什么用、为什么要这样写。<strong>明确问题，可以帮助你提升阅读效率。</strong></p>

<p>以 [06｜手写CPU（一）：迷你CPU架构设计与取指令实现] 这节课为例，“CPU流水线”这部分内容可以分解为以下几个问题：</p>

<p>1.流水线是什么？什么是五级流水线？（What）-
2.CPU为什么要使用流水线？这么做的好处是什么？（Why）-
3.流水线思想在代码中是如何体现的？（How）</p>

<p>第二点建议，阅读的时候分清“主次”，明确哪些对自己解决这个问题有帮助，哪些对自己解决问题没有帮助。时刻注意这一点，这样做可以很好的<strong>帮你把注意力拉回来</strong>，不会读着读着就忘了目的。</p>

<p>还是拿第六课为例，我当时的第一个疑惑是 “CPU流水线是什么？”带着这个疑惑，我通过页内搜索，找到第一处出现“流水线”这个词汇的地方。从这个地方开始往下读。</p>

<p>在这个过程中，因为想着我疑惑的问题，我会下意识着重寻找作者对流水线的定义，相关词汇以及作者的理解。</p>

<p>在看到相关词汇的时候，我会停下快速<strong>瞟一眼</strong>上下文做一个简单的判断，判断这是对“流水线”的解释还是定义。如果是定义，就停下仔细读。如果不是的话，就忽略掉这个词，接着往下找。</p>

<p>比如这句话：</p>

<blockquote>
<p>说到流水线，你是否会马上想到我们打工人的工厂流水线？没错，高大上的 CPU 流水线其实和我们打工人的流水线是一样的。</p>
</blockquote>

<p>在这句话中，看到“工厂流水线”，我暂时停下，快速看了看这个词的前后，发现这段话其实是作者为解释“CPU流水线”做的一个类比，用来辅助大家理解，并不是给“流水线”下的结论，是“次要的”，可以忽略。于是我接着往下读。</p>

<p>然后我注意到这样一句话：</p>

<blockquote>
<p>这样，后续生产中就能够保证五个工人一直处于工作状态，不会造成人员的闲置而产线的冰墩墩就好像流水一样源源不断地产出，<strong>因此我们称这种生产方式为流水线</strong>。</p>
</blockquote>

<p>可以看到，在最后有一句“称……为流水线”这句话，不难推测这就是作者对流水线的定义，是“主要的”。于是，我开始仔仔细细精读这段话。</p>

<blockquote>
<p>这句话，大致可以省略为：……保证……一直处于工作状态，不会造成人员的闲置……源源不断地产出，<strong>……称这种生产方式为流水线</strong>。</p>
</blockquote>

<p>然后我会用自己的话重新概括，比如这样概括流水线：“不会造成人员闲置，可以不断产出商品的一种生产方式”。这样，我们就把“流水线”的定义这个问题解决了。如果看不懂，也没关系，重新返回前文，借助冰墩墩的例子辅助理解即可。</p>

<p>接下来的学习也是如法炮制，比如在理解完流水线，回顾的时候，发现在工厂中每个工人是同时工作的，那么就可以思考“在<strong>CPU运行时，每个地方也是同时工作的么？</strong>”这样的问题，带着这个疑问，你再回看内容，就会发现：在CPU运行时也是同时工作的。</p>

<p>接下来，你可能又会去思考，在Verilog代码中，又是怎么实现同时工作的呢？如果你是按课程顺序学习的，可能就会想起第四课的这段内容：</p>

<blockquote>
<p>Verilog 代码和 C 语言、Java 等这些计算机编程语言有本质的不同，在可综合（这里的“可综合”和代码“编译”的意思差不多）的 <strong>Verilog 代码里，基本所有写出来的东西都对应着实际的电路。</strong></p>
</blockquote>

<p>在留言中也会看到这样一句话：</p>

<blockquote>
<p><strong>硬件设计是特定电路</strong>实现更符合项目，并且是<strong>真正的并行结构</strong>，<strong>软件</strong>是在特定的处理器下进行项目实现，<strong>顺序结构</strong>。效率远低于直接硬件设计实现。</p>
</blockquote>

<p>再综合一些相关知识，你可能就理解了：Verilog是一个硬件描述语言，编译以后生成的是电路。每一个模块、每一个always/assign 语句是同时进行的，这与软件不一样，软件编程是顺序结构，是从上而下依次执行的……</p>

<p>概括一下整个过程就是：不断给自己提问，再结合课程解谜，然后整合搜集到的“线索”，用自己的话概括复盘。</p>

<p>做完这些，你可能对“CPU流水线”在脑中已经有了一个大概的理解，可以去做一个总结概述（最好写下来），验证自己对这个知识点是否理解透彻。如果<strong>有什么阐述不流畅的地方</strong>，写不出来，这就说明你还<strong>未对这个知识点掌握透彻</strong>，还需要重新去看文章或搜索 Google 加强理解。</p>

<p>当然，如果没有疑问，也可以去作者文章里“找找茬”。一切都出于自己的兴趣去搜寻，不必强逼自己。</p>

<p>好，例子就说到这，我实践下来的感觉就是。自问自答很容易有一种成就感，让自己开心一整天，内容理解也更加扎实。</p>

<h2 id="学习资料">学习资料</h2>

<p>最后。我列了一些觉得不错的学习资料：</p>

<ul>
<li><a href="https://hdlbits.01xz.net/wiki/Step_one" target="_blank">HDLBits</a> （全英文刷题网站，无需科学上网，解析详尽。如果英语好/想提升英语水平推荐使用。）</li>
<li><a href="https://www.runoob.com/w3cnote/verilog-basic-syntax.html" target="_blank">verilog基础语法 | 菜鸟教程</a> （Berilog(中文) 学习网站 ）</li>
<li><a href="https://b23.tv/hgfpLYH" target="_blank">入行十年，我总结了这份FPGA学习路线</a> （FPGA入门思维导图）</li>
<li><a href="https://blog.csdn.net/fzr_en/article/details/89552323?spm=1001.2101.3001.6661.1&amp;utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-89552323-blog-104636214.pc_relevant_multi_platform_whitelistv3&amp;depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-89552323-blog-104636214.pc_relevant_multi_platform_whitelistv3&amp;utm_relevant_index=1" target="_blank">阻塞赋值与非阻塞赋值的区别</a> （阻塞与非阻塞解答，建议自己手写直接看仿真理解）</li>
<li><a href="https://b23.tv/XW3S1gZ" target="_blank">【硬件科普】带你认识CPU</a> （CPU 相关科普课程）</li>
<li><a href="http://riscvbook.com/chinese/RISC-V-Reader-Chinese-v2p1.pdf" target="_blank">RISC-V中文手册</a></li>
<li><a href="https://blog.csdn.net/weixin_43240387/article/details/88592279" target="_blank">verilog宏定义用法</a>（ &lsquo;define ）</li>
<li><a href="https://space.bilibili.com/3280670/channel/collectiondetail?sid=305846" target="_blank">RISCV部分原理 不完全解答 讲解_哔哩哔哩_bilibili</a>（感谢Geek_6a1eb9的推荐，这里一并列出）</li>
</ul>

<p>好，我的分享就到这里。我继续去啃代码了。各位加油，一起实现RISC-V 五级流水线 CPU呀！fighting！！！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ee828282d7dadfdfded9ae89838f8782c08d8183" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f186be30fe2ecfb',t:'MTczNDExODc3OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=15&#32;过不了的坎：聊聊GUI自动化过程中的测试数据>
        <link rel="icon" href="/static/favicon.png">
        <title>15 过不了的坎：聊聊GUI自动化过程中的测试数据 </title>
        
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
                        <a class="menu-item" id="00 开篇词 从“小工”到“专家”，我的软件测试修炼之道.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%e2%80%9c%e5%b0%8f%e5%b7%a5%e2%80%9d%e5%88%b0%e2%80%9c%e4%b8%93%e5%ae%b6%e2%80%9d%ef%bc%8c%e6%88%91%e7%9a%84%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%95%e4%bf%ae%e7%82%bc%e4%b9%8b%e9%81%93.md">00 开篇词 从“小工”到“专家”，我的软件测试修炼之道.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 你真的懂测试吗？从“用户登录”测试谈起.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/01%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e6%87%82%e6%b5%8b%e8%af%95%e5%90%97%ef%bc%9f%e4%bb%8e%e2%80%9c%e7%94%a8%e6%88%b7%e7%99%bb%e5%bd%95%e2%80%9d%e6%b5%8b%e8%af%95%e8%b0%88%e8%b5%b7.md">01 你真的懂测试吗？从“用户登录”测试谈起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 如何设计一个“好的”测试用例？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/02%20%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e4%b8%80%e4%b8%aa%e2%80%9c%e5%a5%bd%e7%9a%84%e2%80%9d%e6%b5%8b%e8%af%95%e7%94%a8%e4%be%8b%ef%bc%9f.md">02 如何设计一个“好的”测试用例？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 什么是单元测试？如何做好单元测试？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/03%20%e4%bb%80%e4%b9%88%e6%98%af%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%ef%bc%9f%e5%a6%82%e4%bd%95%e5%81%9a%e5%a5%bd%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%ef%bc%9f.md">03 什么是单元测试？如何做好单元测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 为什么要做自动化测试？什么样的项目适合做自动化测试？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/04%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%81%9a%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%ef%bc%9f%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e9%a1%b9%e7%9b%ae%e9%80%82%e5%90%88%e5%81%9a%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%ef%bc%9f.md">04 为什么要做自动化测试？什么样的项目适合做自动化测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 你知道软件开发各阶段都有哪些自动化测试技术吗？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/05%20%e4%bd%a0%e7%9f%a5%e9%81%93%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%90%84%e9%98%b6%e6%ae%b5%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%e6%8a%80%e6%9c%af%e5%90%97%ef%bc%9f.md">05 你知道软件开发各阶段都有哪些自动化测试技术吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 你真的懂测试覆盖率吗？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/06%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e6%87%82%e6%b5%8b%e8%af%95%e8%a6%86%e7%9b%96%e7%8e%87%e5%90%97%ef%bc%9f.md">06 你真的懂测试覆盖率吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 如何高效填写软件缺陷报告？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/07%20%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%a1%ab%e5%86%99%e8%bd%af%e4%bb%b6%e7%bc%ba%e9%99%b7%e6%8a%a5%e5%91%8a%ef%bc%9f.md">07 如何高效填写软件缺陷报告？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 以终为始，如何才能做好测试计划？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/08%20%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%ef%bc%8c%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e5%81%9a%e5%a5%bd%e6%b5%8b%e8%af%95%e8%ae%a1%e5%88%92%ef%bc%9f.md">08 以终为始，如何才能做好测试计划？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 软件测试工程师的核心竞争力是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/09%20%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%95%e5%b7%a5%e7%a8%8b%e5%b8%88%e7%9a%84%e6%a0%b8%e5%bf%83%e7%ab%9e%e4%ba%89%e5%8a%9b%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">09 软件测试工程师的核心竞争力是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 软件测试工程师需要掌握的非测试知识有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/10%20%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%95%e5%b7%a5%e7%a8%8b%e5%b8%88%e9%9c%80%e8%a6%81%e6%8e%8c%e6%8f%a1%e7%9a%84%e9%9d%9e%e6%b5%8b%e8%af%95%e7%9f%a5%e8%af%86%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">10 软件测试工程师需要掌握的非测试知识有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 互联网产品的测试策略应该如何设计？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/11%20%e4%ba%92%e8%81%94%e7%bd%91%e4%ba%a7%e5%93%81%e7%9a%84%e6%b5%8b%e8%af%95%e7%ad%96%e7%95%a5%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%ef%bc%9f.md">11 互联网产品的测试策略应该如何设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 从0到1：你的第一个GUI自动化测试.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/12%20%e4%bb%8e0%e5%88%b01%ef%bc%9a%e4%bd%a0%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aaGUI%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95.md">12 从0到1：你的第一个GUI自动化测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 效率为王：脚本与数据的解耦 &#43; Page Object模型.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/13%20%e6%95%88%e7%8e%87%e4%b8%ba%e7%8e%8b%ef%bc%9a%e8%84%9a%e6%9c%ac%e4%b8%8e%e6%95%b0%e6%8d%ae%e7%9a%84%e8%a7%a3%e8%80%a6%20&#43;%20Page%20Object%e6%a8%a1%e5%9e%8b.md">13 效率为王：脚本与数据的解耦 &#43; Page Object模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 更接近业务的抽象：让自动化测试脚本更好地描述业务.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/14%20%e6%9b%b4%e6%8e%a5%e8%bf%91%e4%b8%9a%e5%8a%a1%e7%9a%84%e6%8a%bd%e8%b1%a1%ef%bc%9a%e8%ae%a9%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%e8%84%9a%e6%9c%ac%e6%9b%b4%e5%a5%bd%e5%9c%b0%e6%8f%8f%e8%bf%b0%e4%b8%9a%e5%8a%a1.md">14 更接近业务的抽象：让自动化测试脚本更好地描述业务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 过不了的坎：聊聊GUI自动化过程中的测试数据.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/15%20%e8%bf%87%e4%b8%8d%e4%ba%86%e7%9a%84%e5%9d%8e%ef%bc%9a%e8%81%8a%e8%81%8aGUI%e8%87%aa%e5%8a%a8%e5%8c%96%e8%bf%87%e7%a8%8b%e4%b8%ad%e7%9a%84%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae.md">15 过不了的坎：聊聊GUI自动化过程中的测试数据.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 脑洞大开：GUI测试还能这么玩（Page Code Gen &#43; Data Gen &#43; Headless）？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/16%20%e8%84%91%e6%b4%9e%e5%a4%a7%e5%bc%80%ef%bc%9aGUI%e6%b5%8b%e8%af%95%e8%bf%98%e8%83%bd%e8%bf%99%e4%b9%88%e7%8e%a9%ef%bc%88Page%20Code%20Gen%20&#43;%20Data%20Gen%20&#43;%20Headless%ef%bc%89%ef%bc%9f.md">16 脑洞大开：GUI测试还能这么玩（Page Code Gen &#43; Data Gen &#43; Headless）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 精益求精：聊聊提高GUI测试稳定性的关键技术.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/17%20%e7%b2%be%e7%9b%8a%e6%b1%82%e7%b2%be%ef%bc%9a%e8%81%8a%e8%81%8a%e6%8f%90%e9%ab%98GUI%e6%b5%8b%e8%af%95%e7%a8%b3%e5%ae%9a%e6%80%a7%e7%9a%84%e5%85%b3%e9%94%ae%e6%8a%80%e6%9c%af.md">17 精益求精：聊聊提高GUI测试稳定性的关键技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 眼前一亮：带你玩转GUI自动化的测试报告.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/18%20%e7%9c%bc%e5%89%8d%e4%b8%80%e4%ba%ae%ef%bc%9a%e5%b8%a6%e4%bd%a0%e7%8e%a9%e8%bd%acGUI%e8%87%aa%e5%8a%a8%e5%8c%96%e7%9a%84%e6%b5%8b%e8%af%95%e6%8a%a5%e5%91%8a.md">18 眼前一亮：带你玩转GUI自动化的测试报告.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 真实的战场：如何在大型项目中设计GUI自动化测试策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/19%20%e7%9c%9f%e5%ae%9e%e7%9a%84%e6%88%98%e5%9c%ba%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e5%a4%a7%e5%9e%8b%e9%a1%b9%e7%9b%ae%e4%b8%ad%e8%ae%be%e8%ae%a1GUI%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%e7%ad%96%e7%95%a5.md">19 真实的战场：如何在大型项目中设计GUI自动化测试策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 与时俱进：浅谈移动应用测试方法与思路.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/20%20%e4%b8%8e%e6%97%b6%e4%bf%b1%e8%bf%9b%ef%bc%9a%e6%b5%85%e8%b0%88%e7%a7%bb%e5%8a%a8%e5%ba%94%e7%94%a8%e6%b5%8b%e8%af%95%e6%96%b9%e6%b3%95%e4%b8%8e%e6%80%9d%e8%b7%af.md">20 与时俱进：浅谈移动应用测试方法与思路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 移动测试神器：带你玩转Appium.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/21%20%e7%a7%bb%e5%8a%a8%e6%b5%8b%e8%af%95%e7%a5%9e%e5%99%a8%ef%bc%9a%e5%b8%a6%e4%bd%a0%e7%8e%a9%e8%bd%acAppium.md">21 移动测试神器：带你玩转Appium.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 从0到1：API测试怎么做？常用API测试工具简介.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/22%20%e4%bb%8e0%e5%88%b01%ef%bc%9aAPI%e6%b5%8b%e8%af%95%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f%e5%b8%b8%e7%94%a8API%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7%e7%ae%80%e4%bb%8b.md">22 从0到1：API测试怎么做？常用API测试工具简介.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 知其然知其所以然：聊聊API自动化测试框架的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/23%20%e7%9f%a5%e5%85%b6%e7%84%b6%e7%9f%a5%e5%85%b6%e6%89%80%e4%bb%a5%e7%84%b6%ef%bc%9a%e8%81%8a%e8%81%8aAPI%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%e6%a1%86%e6%9e%b6%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">23 知其然知其所以然：聊聊API自动化测试框架的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 紧跟时代步伐：微服务模式下API测试要怎么做？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/24%20%e7%b4%a7%e8%b7%9f%e6%97%b6%e4%bb%a3%e6%ad%a5%e4%bc%90%ef%bc%9a%e5%be%ae%e6%9c%8d%e5%8a%a1%e6%a8%a1%e5%bc%8f%e4%b8%8bAPI%e6%b5%8b%e8%af%95%e8%a6%81%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f.md">24 紧跟时代步伐：微服务模式下API测试要怎么做？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 不破不立：掌握代码级测试的基本理念与方法.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/25%20%e4%b8%8d%e7%a0%b4%e4%b8%8d%e7%ab%8b%ef%bc%9a%e6%8e%8c%e6%8f%a1%e4%bb%a3%e7%a0%81%e7%ba%a7%e6%b5%8b%e8%af%95%e7%9a%84%e5%9f%ba%e6%9c%ac%e7%90%86%e5%bf%b5%e4%b8%8e%e6%96%b9%e6%b3%95.md">25 不破不立：掌握代码级测试的基本理念与方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 深入浅出之静态测试方法.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/26%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e4%b9%8b%e9%9d%99%e6%80%81%e6%b5%8b%e8%af%95%e6%96%b9%e6%b3%95.md">26 深入浅出之静态测试方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 深入浅出之动态测试方法.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/27%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e4%b9%8b%e5%8a%a8%e6%80%81%e6%b5%8b%e8%af%95%e6%96%b9%e6%b3%95.md">27 深入浅出之动态测试方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 带你一起解读不同视角的软件性能与性能指标.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/28%20%e5%b8%a6%e4%bd%a0%e4%b8%80%e8%b5%b7%e8%a7%a3%e8%af%bb%e4%b8%8d%e5%90%8c%e8%a7%86%e8%a7%92%e7%9a%84%e8%bd%af%e4%bb%b6%e6%80%a7%e8%83%bd%e4%b8%8e%e6%80%a7%e8%83%bd%e6%8c%87%e6%a0%87.md">28 带你一起解读不同视角的软件性能与性能指标.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 聊聊性能测试的基本方法与应用领域.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/29%20%e8%81%8a%e8%81%8a%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e7%9a%84%e5%9f%ba%e6%9c%ac%e6%96%b9%e6%b3%95%e4%b8%8e%e5%ba%94%e7%94%a8%e9%a2%86%e5%9f%9f.md">29 聊聊性能测试的基本方法与应用领域.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 工欲善其事必先利其器：后端性能测试工具原理与行业常用工具简介.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/30%20%e5%b7%a5%e6%ac%b2%e5%96%84%e5%85%b6%e4%ba%8b%e5%bf%85%e5%85%88%e5%88%a9%e5%85%b6%e5%99%a8%ef%bc%9a%e5%90%8e%e7%ab%af%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7%e5%8e%9f%e7%90%86%e4%b8%8e%e8%a1%8c%e4%b8%9a%e5%b8%b8%e7%94%a8%e5%b7%a5%e5%85%b7%e7%ae%80%e4%bb%8b.md">30 工欲善其事必先利其器：后端性能测试工具原理与行业常用工具简介.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 工欲善其事必先利其器：前端性能测试工具原理与行业常用工具简介.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/31%20%e5%b7%a5%e6%ac%b2%e5%96%84%e5%85%b6%e4%ba%8b%e5%bf%85%e5%85%88%e5%88%a9%e5%85%b6%e5%99%a8%ef%bc%9a%e5%89%8d%e7%ab%af%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7%e5%8e%9f%e7%90%86%e4%b8%8e%e8%a1%8c%e4%b8%9a%e5%b8%b8%e7%94%a8%e5%b7%a5%e5%85%b7%e7%ae%80%e4%bb%8b.md">31 工欲善其事必先利其器：前端性能测试工具原理与行业常用工具简介.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 无实例无真相：基于LoadRunner实现企业级服务器端性能测试的实践（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/32%20%e6%97%a0%e5%ae%9e%e4%be%8b%e6%97%a0%e7%9c%9f%e7%9b%b8%ef%bc%9a%e5%9f%ba%e4%ba%8eLoadRunner%e5%ae%9e%e7%8e%b0%e4%bc%81%e4%b8%9a%e7%ba%a7%e6%9c%8d%e5%8a%a1%e5%99%a8%e7%ab%af%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e7%9a%84%e5%ae%9e%e8%b7%b5%ef%bc%88%e4%b8%8a%ef%bc%89.md">32 无实例无真相：基于LoadRunner实现企业级服务器端性能测试的实践（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 无实例无真相：基于LoadRunner实现企业级服务器端性能测试的实践（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/33%20%e6%97%a0%e5%ae%9e%e4%be%8b%e6%97%a0%e7%9c%9f%e7%9b%b8%ef%bc%9a%e5%9f%ba%e4%ba%8eLoadRunner%e5%ae%9e%e7%8e%b0%e4%bc%81%e4%b8%9a%e7%ba%a7%e6%9c%8d%e5%8a%a1%e5%99%a8%e7%ab%af%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e7%9a%84%e5%ae%9e%e8%b7%b5%ef%bc%88%e4%b8%8b%ef%bc%89.md">33 无实例无真相：基于LoadRunner实现企业级服务器端性能测试的实践（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 站在巨人的肩膀：企业级实际性能测试案例与经验分享.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/34%20%e7%ab%99%e5%9c%a8%e5%b7%a8%e4%ba%ba%e7%9a%84%e8%82%a9%e8%86%80%ef%bc%9a%e4%bc%81%e4%b8%9a%e7%ba%a7%e5%ae%9e%e9%99%85%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e6%a1%88%e4%be%8b%e4%b8%8e%e7%bb%8f%e9%aa%8c%e5%88%86%e4%ba%ab.md">34 站在巨人的肩膀：企业级实际性能测试案例与经验分享.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 如何准备测试数据？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/35%20%e5%a6%82%e4%bd%95%e5%87%86%e5%a4%87%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%ef%bc%9f.md">35 如何准备测试数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 浅谈测试数据的痛点.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/36%20%e6%b5%85%e8%b0%88%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e7%9a%84%e7%97%9b%e7%82%b9.md">36 浅谈测试数据的痛点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 测试数据的“银弹”- 统一测试数据平台（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/37%20%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e7%9a%84%e2%80%9c%e9%93%b6%e5%bc%b9%e2%80%9d-%20%e7%bb%9f%e4%b8%80%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0%ef%bc%88%e4%b8%8a%ef%bc%89.md">37 测试数据的“银弹”- 统一测试数据平台（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 测试数据的“银弹”- 统一测试数据平台（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/38%20%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e7%9a%84%e2%80%9c%e9%93%b6%e5%bc%b9%e2%80%9d-%20%e7%bb%9f%e4%b8%80%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0%ef%bc%88%e4%b8%8b%ef%bc%89.md">38 测试数据的“银弹”- 统一测试数据平台（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 从小作坊到工厂：什么是Selenium Grid？如何搭建Selenium Grid？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/39%20%e4%bb%8e%e5%b0%8f%e4%bd%9c%e5%9d%8a%e5%88%b0%e5%b7%a5%e5%8e%82%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%afSelenium%20Grid%ef%bc%9f%e5%a6%82%e4%bd%95%e6%90%ad%e5%bb%baSelenium%20Grid%ef%bc%9f.md">39 从小作坊到工厂：什么是Selenium Grid？如何搭建Selenium Grid？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 从小工到专家：聊聊测试执行环境的架构设计（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/40%20%e4%bb%8e%e5%b0%8f%e5%b7%a5%e5%88%b0%e4%b8%93%e5%ae%b6%ef%bc%9a%e8%81%8a%e8%81%8a%e6%b5%8b%e8%af%95%e6%89%a7%e8%a1%8c%e7%8e%af%e5%a2%83%e7%9a%84%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89.md">40 从小工到专家：聊聊测试执行环境的架构设计（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 从小工到专家：聊聊测试执行环境的架构设计（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/41%20%e4%bb%8e%e5%b0%8f%e5%b7%a5%e5%88%b0%e4%b8%93%e5%ae%b6%ef%bc%9a%e8%81%8a%e8%81%8a%e6%b5%8b%e8%af%95%e6%89%a7%e8%a1%8c%e7%8e%af%e5%a2%83%e7%9a%84%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89.md">41 从小工到专家：聊聊测试执行环境的架构设计（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 实战：大型全球化电商的测试基础架构设计.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/42%20%e5%ae%9e%e6%88%98%ef%bc%9a%e5%a4%a7%e5%9e%8b%e5%85%a8%e7%90%83%e5%8c%96%e7%94%b5%e5%95%86%e7%9a%84%e6%b5%8b%e8%af%95%e5%9f%ba%e7%a1%80%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">42 实战：大型全球化电商的测试基础架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 发挥人的潜能：探索式测试.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/43%20%e5%8f%91%e6%8c%a5%e4%ba%ba%e7%9a%84%e6%bd%9c%e8%83%bd%ef%bc%9a%e6%8e%a2%e7%b4%a2%e5%bc%8f%e6%b5%8b%e8%af%95.md">43 发挥人的潜能：探索式测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 测试先行：测试驱动开发(TDD).md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/44%20%e6%b5%8b%e8%af%95%e5%85%88%e8%a1%8c%ef%bc%9a%e6%b5%8b%e8%af%95%e9%a9%b1%e5%8a%a8%e5%bc%80%e5%8f%91%28TDD%29.md">44 测试先行：测试驱动开发(TDD).md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 打蛇打七寸：精准测试.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/45%20%e6%89%93%e8%9b%87%e6%89%93%e4%b8%83%e5%af%b8%ef%bc%9a%e7%b2%be%e5%87%86%e6%b5%8b%e8%af%95.md">45 打蛇打七寸：精准测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 安全第一：渗透测试.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/46%20%e5%ae%89%e5%85%a8%e7%ac%ac%e4%b8%80%ef%bc%9a%e6%b8%97%e9%80%8f%e6%b5%8b%e8%af%95.md">46 安全第一：渗透测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 用机器设计测试用例：基于模型的测试.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/47%20%e7%94%a8%e6%9c%ba%e5%99%a8%e8%ae%be%e8%ae%a1%e6%b5%8b%e8%af%95%e7%94%a8%e4%be%8b%ef%bc%9a%e5%9f%ba%e4%ba%8e%e6%a8%a1%e5%9e%8b%e7%9a%84%e6%b5%8b%e8%af%95.md">47 用机器设计测试用例：基于模型的测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 优秀的测试工程师为什么要懂大型网站的架构设计？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/48%20%e4%bc%98%e7%a7%80%e7%9a%84%e6%b5%8b%e8%af%95%e5%b7%a5%e7%a8%8b%e5%b8%88%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e6%87%82%e5%a4%a7%e5%9e%8b%e7%bd%91%e7%ab%99%e7%9a%84%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%9f.md">48 优秀的测试工程师为什么要懂大型网站的架构设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 深入浅出网站高性能架构设计.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/49%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e7%bd%91%e7%ab%99%e9%ab%98%e6%80%a7%e8%83%bd%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">49 深入浅出网站高性能架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 深入浅出网站高可用架构设计.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/50%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e7%bd%91%e7%ab%99%e9%ab%98%e5%8f%af%e7%94%a8%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">50 深入浅出网站高可用架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 深入浅出网站伸缩性架构设计.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/51%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e7%bd%91%e7%ab%99%e4%bc%b8%e7%bc%a9%e6%80%a7%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">51 深入浅出网站伸缩性架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52 深入浅出网站可扩展性架构设计.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/52%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e7%bd%91%e7%ab%99%e5%8f%af%e6%89%a9%e5%b1%95%e6%80%a7%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">52 深入浅出网站可扩展性架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 浅谈全链路压测.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e6%b5%85%e8%b0%88%e5%85%a8%e9%93%be%e8%b7%af%e5%8e%8b%e6%b5%8b.md">测试专栏特别放送 浅谈全链路压测.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第一期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e4%b8%80%e6%9c%9f.md">测试专栏特别放送 答疑解惑第一期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第七期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e4%b8%83%e6%9c%9f.md">测试专栏特别放送 答疑解惑第七期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第三期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e4%b8%89%e6%9c%9f.md">测试专栏特别放送 答疑解惑第三期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第二期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e4%ba%8c%e6%9c%9f.md">测试专栏特别放送 答疑解惑第二期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第五期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e4%ba%94%e6%9c%9f.md">测试专栏特别放送 答疑解惑第五期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第六期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e5%85%ad%e6%9c%9f.md">测试专栏特别放送 答疑解惑第六期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="测试专栏特别放送 答疑解惑第四期.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e6%b5%8b%e8%af%95%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%e7%ac%ac%e5%9b%9b%e6%9c%9f.md">测试专栏特别放送 答疑解惑第四期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 不是结束，而是开始.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%9552%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e4%b8%8d%e6%98%af%e7%bb%93%e6%9d%9f%ef%bc%8c%e8%80%8c%e6%98%af%e5%bc%80%e5%a7%8b.md">结束语 不是结束，而是开始.md</a>
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
                            <h1 id="title" data-id="15 过不了的坎：聊聊GUI自动化过程中的测试数据" class="title">15 过不了的坎：聊聊GUI自动化过程中的测试数据</h1>
                            <div><p>在前面几篇文章中，我从页面操作的角度介绍了GUI自动化测试，讲解了页面对象模型和业务流程封装，今天我将从测试数据的角度再来谈谈GUI自动化测试。</p>

<p>为了顺利进行GUI测试，往往需要准备测试数据来配合测试的进行，如果不采用事先数据准备的方式，测试效率将会大打折扣，而且还会引入大量不必要的依赖关系。</p>

<p>以“用户登录”功能的测试为例，如果你的目的仅仅是测试用户是否可以正常登录，比较理想的方式是这个用户已经存在于被测系统中了，或者你可以通过很方便的方式在测试用例中生成这个用户。否则，难道你要为了测试用户登录功能，而以GUI的方式当场注册一个新用户吗？显然，这是不可取的。</p>

<p>其实从这里，你就可以看出测试数据准备是实现测试用例解耦的重要手段，你完全不必为了测试GUI用户登录功能而去执行用户注册，只要你能够有方法快速创建出这个登录用户就可以了。</p>

<p>在正式讨论测试数据的创建方法前，我先来分析一下GUI测试中两种常见的数据类型：</p>

<ul>
<li><p>第一大类是，测试输入数据，也就是GUI测试过程中，通过界面输入的数据。比如“用户登录”测试中输入的用户名和密码就就属于这一类数据；再比如，数据驱动测试中的测试数据，也是指这一类。</p></li>

<li><p>第二大类是，为了完成GUI测试而需要准备的测试数据。比如，“用户登录”测试中，我们需要事先准备好用户账户，以便进行用户的登录测试。今天我分享的测试数据创建的方法，也都是围着这一部分的数据展开的。</p></li>
</ul>

<p>那么接下来，我就带你一起去看看创建测试数据的方法都有哪些，以及它们各自的优缺点，和适用场景。</p>

<p>从创建的技术手段上来讲，创建测试数据的方法主要分为三种：</p>

<ol>
<li><p>API调用；</p></li>

<li><p>数据库操作；</p></li>

<li><p>综合运用API调用和数据库操作。</p></li>
</ol>

<p>从创建的时机来讲，创建测试数据的方法主要分为两种：</p>

<ol>
<li><p>测试用例执行过程中，实时创建测试数据，我们通常称这种方式为On-the-fly。</p></li>

<li><p>测试用例执行前，事先创建好“开箱即用”的测试数据，我们通常称这种方式为Out-of-box。</p></li>
</ol>

<p><strong>在实际项目中，对于创建数据的技术手段而言，最佳的选择是利用API来创建数据，只有当API不能满足数据创建的需求时，才会使用数据库操作的手段。</strong></p>

<p>实际上，往往很多测试数据的创建是基于API和数据库操作两者的结合来完成，即先通过API创建基本的数据，然后调用数据库操作来修改数据，以达到对测试数据的特定要求。</p>

<p><strong>而对于创建数据的时机，在实际项目中，往往是On-the-fly和Out-of-box结合在一起使用。</strong></p>

<p>对于相对稳定的测试数据，比如商品类型、图书类型等，往往采用Out-of-box的方式以提高效率；而对于那些只能一次性使用的测试数据，比如商品、订单、优惠券等，往往采用On-the-fly的方式以保证不存在脏数据问题。</p>

<p>接下来，我就先从测试数据创建的技术手段开始今天的分享吧。</p>

<h2 id="基于api调用创建测试数据">基于API调用创建测试数据</h2>

<p>先看一个电商网站“新用户注册”的例子，当用户通过GUI界面完成新用户注册信息填写后，向系统后台递交表单，系统后台就会调用createUser的API完成用户的创建。</p>

<p>而互联网产品，尤其是现在大量采用微服务架构的网站，这个API往往以Web Service的形式暴露接口。那么，在这种架构下，你完全可以直接调用这个API来创建新用户，而无须再向后台递交表单。</p>

<p>由于API通常都有安全相关的token机制来保护，所以实际项目中，通常会把对这些API的调用以代码的形式封装为测试数据工具（Test Data Utility）。</p>

<p>这种方式最大的好处就是，测试数据的准确性直接由产品API保证，缺点是并不是所有的测试数据都有相关的API来支持。</p>

<p>另外，对需要大量创建数据的测试来说，基于API调用方式的执行效率，即使采用了并发机制也不会十分理想。为了解决执行效率的问题，就有了基于数据库操作的测试数据创建手段。</p>

<h2 id="基于数据库操作创建测试数据">基于数据库操作创建测试数据</h2>

<p>实际项目中，并不是所有的数据都可以通过API的方式实现创建和修改，很多数据的创建和修改直接在产品代码内完成，而且并没有对外暴露供测试使用的接口。</p>

<p>那么，这种情况下，你就需要通过直接操作数据库的方式来产生测试数据。</p>

<p>同样地，我们可以把创建和修改数据的相关SQL语句封装成测试数据工具，以方便测试用例的使用。但是，如果你正尝试在实际项目中运用这个方法，不可避免地会遇到如何才能找到正确的SQL语句来创建和修改数据的问题。</p>

<p>因为，创建或修改一条测试数据往往会涉及很多业务表，任何的遗漏都会造成测试数据的不准确，从而导致有些测试因为数据问题而无法进行。</p>

<p>那么，现在我就提供两个思路来帮你解决这个问题：</p>

<ol>
<li><p>手工方式。查阅设计文档和产品代码，找到相关的SQL语句集合。或者，直接找开发人员索要相关的SQL语句集合。</p></li>

<li><p>自动方式。在测试环境中，先在只有一个活跃用户的情况下，通过GUI界面操作完成数据的创建、修改，然后利用数据库监控工具获取这段时间内所有的业务表修改记录，以此为依据开发SQL语句集。</p></li>
</ol>

<p>需要注意的是，这两种思路的前提都是，假定产品功能正确，否则就会出现“一错到底”的尴尬局面。</p>

<p>基于数据库操作创建测试数据的最大好处是，可以创建和修改API不支持的测试数据，并且由于是直接数据库操作，执行效率会远远高于API调用方法。</p>

<p>但是，数据库操作这种方式的缺点也显而易见，数据库表操作的任何变更，都必须同步更新测试数据工具中的SQL语句。</p>

<p>但很不幸的是，在实际项目中，经常出现因为SQL语句更新不及时而导致测试数据错误的问题，而且这里的数据不准确往往只是局部错误，因此这类问题往往比较隐蔽，只有在特定的测试场景下才会暴露。</p>

<p>所以，在实际工程项目中，需要引入测试数据工具的版本管理，并通过开发流程来保证SQL的变更能够及时通知到测试数据工具团队。</p>

<h2 id="综合运用api调用和数据库操作创建测试数据">综合运用API调用和数据库操作创建测试数据</h2>

<p>你如果已经理解了基于API调用和基于数据库操作创建测试数据这两类方法，那么综合运用这两类方法，就是使得测试数据工具能够提供更多种类的业务测试数据。</p>

<p>具体来讲，当你要创建一种特定的测试数据时，你发现没有直接API支持，但是可以通过API先创建一个基本的数据，然后再通过修改数据库的方式来更新这个数据，以此来达到创建特定测试数据的要求。</p>

<p>比如，你需要创建一个已经绑定了信用卡的新用户，如果创建新用户有直接的API，而绑定信用卡需要操作数据库，那这种情况下就需要综合运用这两种方式完成测试数据工具的开发。</p>

<h2 id="实时创建数据-on-the-fly">实时创建数据：On-the-fly</h2>

<p><strong>GUI测试脚本中，在开始执行界面操作前，我们往往会通过调用测试数据工具实时创建测试数据，也就是On-the-fly方式。</strong></p>

<p>这种方式不依赖被测试系统中的任何原有数据，也不会对原有数据产生影响，可以很好地从数据层面隔离测试用例，让测试用例实现“自包含”。</p>

<p>从理论上讲，On-the-fly是很好的方法，但在实际测试项目中却并不是那么回事儿，往往会存在三个问题：</p>

<ol>
<li><p><strong>在用例执行过程中实时创建数据，导致测试的执行时间比较长。</strong> 我曾经粗略统计过一个大型Web GUI自动化测试项目的执行时间，将近30%的时间都花在了测试数据的准备上。</p></li>

<li><p><strong>业务数据的连带关系，导致测试数据的创建效率非常低。</strong> 比如，你需要创建一个订单数据，而这个订单必然会绑定买家和卖家，以及订单商品信息。-
如果完全基于On-the-fly模式，你就需要先实时创建买家和卖家这两个用户，然后再创建订单中的商品，最后才是创建这个订单本身。-
显然，这样的测试数据创建方式虽然是“自包含”的，但创建效率非常低，会使得测试用例执行时间变得更长，而这恰恰与互联网产品的测试策略产生冲突。</p></li>

<li><p><strong>更糟糕的情况是，实时创建测试数据的方式对测试环境的依赖性很强。</strong> 比如，你要测试用户登录功能，基于On-the-fly方式，你就应该先调用测试数据工具实时创建一个用户，然后再用这个用户完成登录测试。-
这时，创建用户的API由于各种原因处于不可用的状态（这种情况在采用微服务架构的系统中很常见），那么这时就会因为无法创建用户，而无法完成用户登录测试。</p></li>
</ol>

<p>基于这三种常见问题，实际项目中还会引入Out-of-box方式（即在执行测试用例前，预先创建好测试数据）准备测试数据。</p>

<h2 id="事先创建测试数据-out-of-box">事先创建测试数据：Out-of-box</h2>

<p><strong>Out-of-box的含义是开箱即用，也就是说，已经在被测系统中预先创建好了充足的、典型的测试数据。这些数据通常是在搭建测试环境时通过数据库脚本“预埋”在系统中的，后续的测试用例可以直接使用。</strong></p>

<p>Out-of-box的方式有效解决了On-the-fly的很多问题，但是这种方法的缺点也很明显，主要体现在以下三个方面：</p>

<ol>
<li><p><strong>测试用例中需要硬编码（hardcode）测试数据，额外引入了测试数据和用例之间的依赖。</strong></p></li>

<li><p><strong>只能被一次性使用的测试数据不适合Out-of-box的方式。</strong> 测试用例往往会需要修改测试数据，而且有些测试数据只能被一次性使用。比如，一个商品被买下一次后就不能再用了；再比如，优惠券在一个订单中被使用后，就失效了，等等。所以如果没有很好的全局测试数据管理，很容易因为测试数据失效而造成测试失败。</p></li>

<li><p><strong>“预埋”的测试数据的可靠性远不如实时创建的数据。</strong> 在测试用例执行过程中，经常会出现测试数据被修改的情况。比如，手动测试，或者是自动化测试用例的调试等情况。</p></li>
</ol>

<h2 id="on-the-fly和out-of-box的互补">On-the-fly和Out-of-box的互补</h2>

<p>基于On-the-fly和Out-of-box的优缺点和互补性，在实际的大型测试项目中，我们往往会采用两者相结合的方式，从测试数据本身的特点入手，选取不同的测试数据创建方式。</p>

<p>针对应该选择什么时机创建测试数据，结合多年的实践经验，我为你总结了以下三点：</p>

<ol>
<li><p>对于相对稳定、很少有修改的数据，建议采用Out-of-box的方式，比如商品类目、厂商品牌、部分标准的卖家和买家账号等。</p></li>

<li><p>对于一次性使用、经常需要修改、状态经常变化的数据，建议使用On-the-fly的方式。</p></li>

<li><p>用On-the-fly方式创建测试数据时，上游数据的创建可以采用Out-of-box方式，以提高测试数据创建的效率。以订单数据为例，订单的创建可以采用On-the-fly方式，而与订单相关联的卖家、买家和商品信息可以使用Out-of-box方式创建。</p></li>
</ol>

<p>其实，为了更好地解决测试数据本身组合的复杂性和多样性，充分发挥测试数据工具的威力，还有很多大型企业的最佳实践值得讨论，在本专栏后面的测试数据章节，我会再为你详细介绍。</p>

<h2 id="总结">总结</h2>

<p>今天我从创建测试数据的技术手段和时机两个方面，介绍了GUI测试数据的准备。</p>

<p>在实际测试项目中，往往需要综合运用API调用和数据库操作来创建测试数据，并且会根据测试数据自身的特点，分而治之地采用On-the-fly和Out-of-box的方式，以寻求数据稳定性和数据准备效率之间的最佳平衡。</p>

<h2 id="思考题">思考题</h2>

<p>你所在的公司是如何准备GUI测试的测试数据的？遇到了哪些问题，对应的有哪些解决方案呢？</p>

<p>欢迎你给我留言。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#533f3f3f6a676262636413343e323a3f7d303c3e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f19b387eb42cdba',t:'MTczNDEzMjE5OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
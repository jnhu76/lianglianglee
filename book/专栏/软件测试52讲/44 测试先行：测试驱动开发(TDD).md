<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=44&#32;测试先行：测试驱动开发(TDD)>
        <link rel="icon" href="/static/favicon.png">
        <title>44 测试先行：测试驱动开发(TDD) </title>
        
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
                            <h1 id="title" data-id="44 测试先行：测试驱动开发(TDD)" class="title">44 测试先行：测试驱动开发(TDD)</h1>
                            <div><p>你好，我是茹炳晟。今天我和你分享的主题是“测试先行：测试驱动开发（TDD）”。</p>

<p>通过上一篇文章，我们已经深入理解了什么是探索式测试，以及如何用探索式测试开展具体的测试。今天我这次分享的目的，就是和你聊聊软件测试领域中的另一个很热门的话题：测试驱动开发，也就是Test-Driven Development，通常简称为TDD。</p>

<p>听上去有些迷惑是不是？测试怎么可能驱动开发呢？在传统软件的开发流程中，软件开发人员先开发好功能代码，再针对这些功能设计测试用例、实现测试脚本，以此保证开发的这些功能的正确性和稳定性。那么，TDD从字面上理解就是要让测试先行，这又是怎么一回事呢？</p>

<p>确切地说，TDD并不是一门技术，而是一种开发理念。它的核心思想，是在开发人员实现功能代码前，先设计好测试用例的代码，然后再根据测试用例的代码编写产品的功能代码，最终目的是让开发前设计的测试用例代码都能够顺利执行通过。</p>

<p>这样对于开发人员来说，他就需要参与到这个功能的完整设计过程中，而不是凭自己想象去开发一个功能。他有一个非常明确的目标，就是要让提前设计的测试用例都可以顺利通过，为此，他先实现测试用例要求的功能，再通过不断修改和完善，让产品代码可以满足测试用例，可以说是“小而美”的开发过程。</p>

<p>所以，从本质上来讲，TDD并不属于测试技术的范畴。那么，我为什么还要单独用一篇文章和你分享这个主题呢？因为，TDD中通常会用到很多常见的自动化测试技术，使得测试在整个软件生命周期中的重要性和地位得到了大幅提升。</p>

<p>可以说，TDD的思想和理念给软件研发流程带来了颠覆性的变化，使得测试工作从原本软件研发生命周期的最后端走向了最前端。也就是说，原本测试工作是软件研发生命周期最后的一个环节，而现在TDD相当于把测试提到了需求定义的高度，跑到了软件研发生命周期最前面。</p>

<p>那么，接下来我们就一起看看TDD的优势有哪些，以及TDD的具体实施过程。</p>

<h2 id="tdd的优势">TDD的优势</h2>

<p>TDD的优势，可以概括为以下五个方面：</p>

<ol>
<li><strong>保证开发的功能一定是符合实际需求的。</strong></li>
</ol>

<p>用户需求才应该是软件开发的源头，但在实际的软件开发过程中，往往会在不知情的情况下，或者自己的主观判断下，开发出一个完全没有实际应用场景的功能。而这些没有实际应用场景的功能，却因为产品验证和测试工作介入的时机都在项目后期，所以往往在集成测试中或者产品上线后才会被发现。</p>

<p>比如，开发人员在实现用户注册的功能时，认为需要提供使用手机号注册的功能。但是，这个功能开发完成后，测试人员却告知开发人员这个功能用不上，或者产品上线后才发现这个功能在实际场景中完全不是必须的，因为用户可以使用邮箱注册，然后再通过绑定手机号实现手机号登陆。所以，直接用手机号注册这个功能是不需要的，真正需要的是绑定邮箱和手机号的功能。</p>

<p>试想一下，如果是测试驱动开发，即先根据用户的实际需求编写测试用例，再根据测试用例来完成功能代码，就不会出现这种既浪费时间、精力，又没有必要的功能了。</p>

<ol>
<li><strong>更加灵活的迭代方式。</strong></li>
</ol>

<p>传统的需求文档，往往会从比较高的层次去描述功能。开发人员面对这种抽象的需求文档，往往会感觉无从下手。但是，在TDD的流程里，需求是以测试用例描述的，非常具体。那么，开发人员拿到这样的需求时，就可以先开发一个很明确的、针对用户某一个小需求的功能代码。</p>

<p>在开发过程中，开发人员可以不断的调试这个功能，通过测试-&gt;失败-修改/重构-&gt;测试-&gt;成功的过程，使开发的代码符合预期，而不是等所有功能开发完成后，再将一个笨重的产品交给测试人员进行一个长周期的测试，发现缺陷后再整个打回来修改，然后由此又可能会引入新的缺陷。</p>

<p>另外，如果用户需求有变化，我们能够很快地定位到要修改的功能，从而实现快速修改。</p>

<ol>
<li><strong>保证系统的可扩展性。</strong></li>
</ol>

<p>为了满足测试先行的灵活迭代方式，我们会要求开发人员设计更松耦合的系统，以保证它的可扩展性和易修改性。这就要求，开发人员在设计系统时，要考虑它的整体架构，搭建系统的骨架，提供规范的接口定义而非具体的功能类。</p>

<p>这样，当用户需求有变化时，或者有新增测试用例时，能够通过设计的接口快速实现新功能，满足新的测试场景。</p>

<ol>
<li><strong>更好的质量保证。</strong></li>
</ol>

<p>TDD要求测试先于开发，也就是说在每次新增功能时，都需要先用测试用例去验证功能是否运行正常，并运行所有的测试来保证整个系统的质量。在这个测试先行的过程中，开发人员会不断调试功能模块、优化设计、重构代码，使其能够满足所有测试场景。所以，很多的代码实现缺陷和系统设计漏洞，都会在这个不断调优的过程中暴露出来。</p>

<p>也就是说，TDD可以保证更好的产品质量。</p>

<ol>
<li><strong>测试用例即文档。</strong></li>
</ol>

<p>因为在TDD过程中编写的测试用例，首先一定是贴合用户实际需求的，然后又在开发调试的过程中经过了千锤百炼，即一定是符合系统的业务逻辑的，所以我们直接将测试用例生成需求文档。</p>

<p>这里，直接将测试用例生成需求文档的方法有很多、很简单的方法，比如JavaDoc。</p>

<p>这样，我们就无须再花费额外的精力，去撰写需求文档了。</p>

<p>你看，TDD真的是优势多多吧。那么，接下来我们就一起来看看实施TDD的具体过程。</p>

<h2 id="测试驱动开发的实施过程">测试驱动开发的实施过程</h2>

<p>站在全局的角度来看，TDD的整个过程遵循以下流程：</p>

<ol>
<li><p>为需要实现的新功能添加一批测试；</p></li>

<li><p>运行所有测试，看看新添加的测试是否失败；</p></li>

<li><p>编写实现软件新功能的实现代码；</p></li>

<li><p>再次运行所有的测试，看是否有测试失败；</p></li>

<li><p>重构代码；</p></li>

<li><p>重复以上步骤直到所有测试通过。</p></li>
</ol>

<p>接下来，我们就通过一个具体的例子，来看看TDD的整个流程吧。</p>

<p>我们现在要实现这么一个功能：用户输入自己的生日，就可以输出还要多少天到下次生日。</p>

<p>根据TDD测试先行的原则，我们<strong>首先要做的是设计测试用例。</strong></p>

<p>测试用例一，用户输入空字符串或者null：</p>

<pre><code>@Test
//测试输入空字符串null时，是否抛出&quot;Birthday should not be null or empty&quot;异常
public void birthdayIsNull() {
    RuntimeException exception = null;
    try {
        BirthdayCaculator.caculate(null);
    }catch(RuntimeException e) {
       exception = e;
    }
    Assert.assertNotNull(exception);
    Assert.assertEquals(exception.getMessage(), &quot;Birthday should not be null or empty&quot;);
}

@Test
//测试输入空字符串&quot;&quot;时，是否抛出&quot;Birthday should not be null or empty&quot;异常
public void birthdayIsEmpty() {
    RuntimeException exception = null;
    try {
        BirthdayCaculator.caculate(&quot;&quot;);
    }catch(RuntimeException e) {
        exception = e;
    }
    Assert.assertNotNull(exception);
    Assert.assertEquals(exception.getMessage(), &quot;Birthday should not be null or empty&quot;);
}
</code></pre>

<p>根据这个测试用例，我们可以很容易地写出这部分的Java代码：</p>

<pre><code>public static int caculate(String birthday) {
    if(birthday == null || birthday.isEmpty()) {
        throw new RuntimeException(&quot;Birthday should not be null or empty&quot;);
    }
}
</code></pre>

<p>测试用例二，用户输入的生日格式不符合YYYY-MM-dd的格式：</p>

<pre><code>@Test
//测试输入错误的时间格式，是否抛出&quot;Birthday format is invalid!&quot;异常
public void birthdayFormatIsInvalid() {
    RuntimeException exception = null;
    try {
        BirthdayCaculator.caculate(&quot;Sep 3, 1996&quot;);
    }catch(RuntimeException e) {
        exception = e;
    }
    Assert.assertNotNull(exception);
    Assert.assertEquals(exception.getMessage(), &quot;Birthday format is invalid！&quot;);
}
</code></pre>

<p>那么，这部分的Java代码实现便要catch住ParseException, 重新自定义错误信息并抛出异常。</p>

<pre><code>SimpleDateFormat sdf = new SimpleDateFormat(&quot;yyyy-MM-dd&quot;);
Calendar birthDate = Calendar.getInstance();
try {
    //使用SimpleDateFormat来格式化输入日期值
    birthDate.setTime(sdf.parse(birthday));
} catch (ParseException e) {
    throw new RuntimeException(&quot;Birthday format is invalid!&quot;);
}
</code></pre>

<p>测试用例三，用户输入的生日格式正确，但是今年的生日已经过了，就应该返回离明年的生日还有多少天：</p>

<pre><code>@Test
//测试用户输入的日期晚于今年生日的情况，判断是否返回离明年的生日有多少天
public void thisYearBirthdayPassed() {
    Calendar birthday = Calendar.getInstance();
    birthday.add(Calendar.DATE, -1);
    SimpleDateFormat sdf = new SimpleDateFormat(&quot;YYYY-MM-dd&quot;);
    String date = sdf.format(birthday.getTime());
    int days = BirthdayCaculator.caculate(date);
    //天数不应该出现负数
    Assert.assertTrue(days &gt; 0);
}
</code></pre>

<p>测试用例四，用户输入的生日格式正确且今年生日还没过，返回的结果应该不大于365天：</p>

<pre><code>@Test
//测试用户输入的日期早于今年生日的情况，判断返回的天数是否小于365
public void thisYearBirthdayNotPass() {
    Calendar birthday = Calendar.getInstance();
    birthday.add(Calendar.DATE, 5);
    SimpleDateFormat sdf = new SimpleDateFormat(&quot;YYYY-MM-dd&quot;);
    String date = sdf.format(birthday.getTime());
    int days = BirthdayCaculator.caculate(date);
    //天数不应该大于一年的天数，365天
    Assert.assertTrue(days &lt; 365);
}
</code></pre>

<p>测试用例五，用户输入的生日格式正确并且是今天，返回的结果应该为0：</p>

<pre><code>@Test
//测试用户输入的日期恰好等于今年生日的情况，判断返回的天数是否是0
public void todayIsBirthday() {
    Calendar birthday = Calendar.getInstance();
    SimpleDateFormat sdf = new SimpleDateFormat(&quot;YYYY-MM-dd&quot;);
    String date = sdf.format(birthday.getTime());
    int days = BirthdayCaculator.caculate(date);
    Assert.assertEquals(days, 0);
}
</code></pre>

<p>综合上述五种测试场景，根据测试用例，我们可以编写完整的功能代码覆盖所有类型的用户输入，完整代码如下：</p>

<pre><code>public static int caculate(String birthday) {
    //首先对输入的日期是否是null或者是&quot;&quot;进行判断
    if(birthday == null || birthday.isEmpty()) {
        throw new RuntimeException(&quot;Birthday should not be null or empty&quot;);
    }

    SimpleDateFormat sdf = new SimpleDateFormat(&quot;yyyy-MM-dd&quot;);
    Calendar today = Calendar.getInstance();

    //处理输入的日期恰好等于今年生日的情况
    if(birthday.equals(sdf.format(today.getTime()))) {
        return 0;
    }

    //输入日期格式的有效性检查
    Calendar birthDate = Calendar.getInstance();
    try {
        birthDate.setTime(sdf.parse(birthday));
    } catch (ParseException e) {
        throw new RuntimeException(&quot;Birthday format is invalid!&quot;);
    }
    birthDate.set(Calendar.YEAR, today.get(Calendar.YEAR));

    //实际计算的逻辑
    int days;
    if (birthDate.get(Calendar.DAY_OF_YEAR) &lt; today.get(Calendar.DAY_OF_YEAR)) {
        days = today.getActualMaximum(Calendar.DAY_OF_YEAR) - today.get(Calendar.DAY_OF_YEAR);
        days += birthDate.get(Calendar.DAY_OF_YEAR);
    } else {
        days = birthDate.get(Calendar.DAY_OF_YEAR) - today.get(Calendar.DAY_OF_YEAR);
    }
    return days;
}
</code></pre>

<p>以上场景，每添加一个新的功能点，都会添加一个测试方法；完成新功能点的软件代码后，接着运行当前所有的测试用例，以保证新加的功能代码能够满足现有的测试需求。这就是一个典型的TDD过程了。但是，在实际开发场景，肯定会更复杂， 你想要用TDD思想写出健壮稳定的代码，就需要深入理解TDD中的每一步。</p>

<p><strong>首先，需要控制TDD测试用例的粒度</strong>。如果测试用例并不是最小粒度的单元测试，开发人员就不能不假思索地直接根据测试用例开发功能代码，而应该先把测试用例分解成更小粒度的任务列表，保证每一个任务列表都是一个最小的功能模块。</p>

<p>在开发过程中，要把测试用例当成用户，不断分析他可能会怎样调用这个功能，大到功能的设计是用类还是接口，小到方法的参数类型，都要充分考虑到用户的使用场景。</p>

<p><strong>其次，要注意代码的简洁和高效</strong>。随着功能代码的增加，开发人员为了让测试能顺利通过，很可能会简单粗暴地使用复制粘贴来完成某个功能，而这就违背了TDD的初衷，本来是为了写出更优雅的代码，结果反而造成了代码冗余混乱。因此，在开发-测试循环过程中，我们要不断地检查代码，时刻注意是否有重复代码、以及不需要的功能，将功能代码变得更加高效优雅。</p>

<p><strong>最后，通过重构保证最终交付代码的优雅和简洁</strong>。所有功能代码都完成，所有测试都通过之后，我们就要考虑重构了。这里可以考虑类名、方法名甚至变量名命名，是否规范且有意义，太长的类可以考虑拆分；从系统角度检查是否有重复代码，是否有可以合并的代码，你也可以参考市面上比较权威的关于重构的书完成整个系统的重构和优化。这里我建议你阅读Martin Fowler的《重构：改善既有代码的设计》这本书。</p>

<p><strong>总的来说，TDD有其优于传统开发的特点，但在实际开发过程中，我们应该具体场景具体分析。</strong></p>

<p>比如，最典型的一个场景就是，一个旧系统需要翻新重做，并且针对这个老系统已经有很多不错的测试用例了，这就很适合选用TDD。</p>

<p>总之，我们可以通过分析当前的时间、人、方式、效果各要素来最终决定是否选用TDD。另外，需要特别注意的是，选用TDD并不是测试人员或者测试部门的事情，而是需要公司层面的流程和体系的配合，也正是这种原因，虽然大家都能看到TDD的优势，但是在实际项目中的运用还是比较有限。</p>

<h2 id="总结">总结</h2>

<p>今天我和你分享了测试驱动开发的核心理念，以及TDD的优势。</p>

<p>TDD的核心思想便是在开发人员实现功能代码前，先设计好测试用例，编写测试代码，然后再针对新增的测试代码来编写产品的功能代码，最终目的是让新增的测试代码能够通过。</p>

<p>相对于传统软件开发流程，TDD的优势主要包括对需求精准的把控、更灵活的迭代、促使更好的系统设计、更好的交付质量以及轻量级的文档等。</p>

<p>最后，我用用“用户输入自己的生日，就可以输出还要多少天到下次生日”作为例子，展示了测试驱动开发的完整流程，希望帮助你对TDD有更直观的认识。</p>

<h2 id="思考题">思考题</h2>

<p>在实际的工程项目中，你实际使用过TDD吗？如果有的话，是否可以分享一下你的实践心得？如果没有的话，你是否可以设象一下你会怎么规划和设计一个TDD的项目？</p>

<p>感谢你的收听，欢迎你给我留言一起讨论。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#650909095c5154545552250208040c094b060a08" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f19b9dbef7bcdba',t:'MTczNDEzMjQ1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=35&#32;总是在说MVC分层架构，但你真的理解分层吗？>
        <link rel="icon" href="/static/favicon.png">
        <title>35 总是在说MVC分层架构，但你真的理解分层吗？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 程序员解决的问题，大多不是程序问题.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e7%a8%8b%e5%ba%8f%e5%91%98%e8%a7%a3%e5%86%b3%e7%9a%84%e9%97%ae%e9%a2%98%ef%bc%8c%e5%a4%a7%e5%a4%9a%e4%b8%8d%e6%98%af%e7%a8%8b%e5%ba%8f%e9%97%ae%e9%a2%98.md">00 开篇词 程序员解决的问题，大多不是程序问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 10x程序员是如何思考的？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/01%2010x%e7%a8%8b%e5%ba%8f%e5%91%98%e6%98%af%e5%a6%82%e4%bd%95%e6%80%9d%e8%80%83%e7%9a%84%ef%bc%9f.md">01 10x程序员是如何思考的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 以终为始：如何让你的努力不白费？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/02%20%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e4%bd%a0%e7%9a%84%e5%8a%aa%e5%8a%9b%e4%b8%8d%e7%99%bd%e8%b4%b9%ef%bc%9f.md">02 以终为始：如何让你的努力不白费？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 DoD的价值：你完成了工作，为什么他们还不满意？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/03%20DoD%e7%9a%84%e4%bb%b7%e5%80%bc%ef%bc%9a%e4%bd%a0%e5%ae%8c%e6%88%90%e4%ba%86%e5%b7%a5%e4%bd%9c%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bb%96%e4%bb%ac%e8%bf%98%e4%b8%8d%e6%bb%a1%e6%84%8f%ef%bc%9f.md">03 DoD的价值：你完成了工作，为什么他们还不满意？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 接到需求任务，你要先做哪件事？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/04%20%e6%8e%a5%e5%88%b0%e9%9c%80%e6%b1%82%e4%bb%bb%e5%8a%a1%ef%bc%8c%e4%bd%a0%e8%a6%81%e5%85%88%e5%81%9a%e5%93%aa%e4%bb%b6%e4%ba%8b%ef%bc%9f.md">04 接到需求任务，你要先做哪件事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 持续集成：集成本身就是写代码的一个环节.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/05%20%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%ef%bc%9a%e9%9b%86%e6%88%90%e6%9c%ac%e8%ba%ab%e5%b0%b1%e6%98%af%e5%86%99%e4%bb%a3%e7%a0%81%e7%9a%84%e4%b8%80%e4%b8%aa%e7%8e%af%e8%8a%82.md">05 持续集成：集成本身就是写代码的一个环节.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 精益创业：产品经理不靠谱，你该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/06%20%e7%b2%be%e7%9b%8a%e5%88%9b%e4%b8%9a%ef%bc%9a%e4%ba%a7%e5%93%81%e7%bb%8f%e7%90%86%e4%b8%8d%e9%9d%a0%e8%b0%b1%ef%bc%8c%e4%bd%a0%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">06 精益创业：产品经理不靠谱，你该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 解决了很多技术问题，为什么你依然在“坑”里？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/07%20%e8%a7%a3%e5%86%b3%e4%ba%86%e5%be%88%e5%a4%9a%e6%8a%80%e6%9c%af%e9%97%ae%e9%a2%98%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e4%be%9d%e7%84%b6%e5%9c%a8%e2%80%9c%e5%9d%91%e2%80%9d%e9%87%8c%ef%bc%9f.md">07 解决了很多技术问题，为什么你依然在“坑”里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 为什么说做事之前要先进行推演？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/08%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4%e5%81%9a%e4%ba%8b%e4%b9%8b%e5%89%8d%e8%a6%81%e5%85%88%e8%bf%9b%e8%a1%8c%e6%8e%a8%e6%bc%94%ef%bc%9f.md">08 为什么说做事之前要先进行推演？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 你的工作可以用数字衡量吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/09%20%e4%bd%a0%e7%9a%84%e5%b7%a5%e4%bd%9c%e5%8f%af%e4%bb%a5%e7%94%a8%e6%95%b0%e5%ad%97%e8%a1%a1%e9%87%8f%e5%90%97%ef%bc%9f.md">09 你的工作可以用数字衡量吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 迭代0_ 启动开发之前，你应该准备什么？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/10%20%e8%bf%ad%e4%bb%a30_%20%e5%90%af%e5%8a%a8%e5%bc%80%e5%8f%91%e4%b9%8b%e5%89%8d%ef%bc%8c%e4%bd%a0%e5%ba%94%e8%af%a5%e5%87%86%e5%a4%87%e4%bb%80%e4%b9%88%ef%bc%9f.md">10 迭代0_ 启动开发之前，你应该准备什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 向埃隆·马斯克学习任务分解.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/11%20%e5%90%91%e5%9f%83%e9%9a%86%c2%b7%e9%a9%ac%e6%96%af%e5%85%8b%e5%ad%a6%e4%b9%a0%e4%bb%bb%e5%8a%a1%e5%88%86%e8%a7%a3.md">11 向埃隆·马斯克学习任务分解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 测试也是程序员的事吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/12%20%e6%b5%8b%e8%af%95%e4%b9%9f%e6%98%af%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e4%ba%8b%e5%90%97%ef%bc%9f.md">12 测试也是程序员的事吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 先写测试，就是测试驱动开发吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/13%20%e5%85%88%e5%86%99%e6%b5%8b%e8%af%95%ef%bc%8c%e5%b0%b1%e6%98%af%e6%b5%8b%e8%af%95%e9%a9%b1%e5%8a%a8%e5%bc%80%e5%8f%91%e5%90%97%ef%bc%9f.md">13 先写测试，就是测试驱动开发吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 大师级程序员的工作秘笈.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/14%20%e5%a4%a7%e5%b8%88%e7%ba%a7%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e5%b7%a5%e4%bd%9c%e7%a7%98%e7%ac%88.md">14 大师级程序员的工作秘笈.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 一起练习：手把手带你分解任务.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/15%20%e4%b8%80%e8%b5%b7%e7%bb%83%e4%b9%a0%ef%bc%9a%e6%89%8b%e6%8a%8a%e6%89%8b%e5%b8%a6%e4%bd%a0%e5%88%86%e8%a7%a3%e4%bb%bb%e5%8a%a1.md">15 一起练习：手把手带你分解任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 为什么你的测试不够好？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/16%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e7%9a%84%e6%b5%8b%e8%af%95%e4%b8%8d%e5%a4%9f%e5%a5%bd%ef%bc%9f.md">16 为什么你的测试不够好？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 程序员也可以“砍”需求吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/17%20%e7%a8%8b%e5%ba%8f%e5%91%98%e4%b9%9f%e5%8f%af%e4%bb%a5%e2%80%9c%e7%a0%8d%e2%80%9d%e9%9c%80%e6%b1%82%e5%90%97%ef%bc%9f.md">17 程序员也可以“砍”需求吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 需求管理：太多人给你安排任务，怎么办？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/18%20%e9%9c%80%e6%b1%82%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a4%aa%e5%a4%9a%e4%ba%ba%e7%bb%99%e4%bd%a0%e5%ae%89%e6%8e%92%e4%bb%bb%e5%8a%a1%ef%bc%8c%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">18 需求管理：太多人给你安排任务，怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 如何用最小的代价做产品？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/19%20%e5%a6%82%e4%bd%95%e7%94%a8%e6%9c%80%e5%b0%8f%e7%9a%84%e4%bb%a3%e4%bb%b7%e5%81%9a%e4%ba%a7%e5%93%81%ef%bc%9f.md">19 如何用最小的代价做产品？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 为什么世界和你的理解不一样？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/20%20%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%96%e7%95%8c%e5%92%8c%e4%bd%a0%e7%9a%84%e7%90%86%e8%a7%a3%e4%b8%8d%e4%b8%80%e6%a0%b7%ef%bc%9f.md">20 为什么世界和你的理解不一样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 你的代码为谁而写？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/21%20%e4%bd%a0%e7%9a%84%e4%bb%a3%e7%a0%81%e4%b8%ba%e8%b0%81%e8%80%8c%e5%86%99%ef%bc%9f.md">21 你的代码为谁而写？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 轻量级沟通：你总是在开会吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/22%20%e8%bd%bb%e9%87%8f%e7%ba%a7%e6%b2%9f%e9%80%9a%ef%bc%9a%e4%bd%a0%e6%80%bb%e6%98%af%e5%9c%a8%e5%bc%80%e4%bc%9a%e5%90%97%ef%bc%9f.md">22 轻量级沟通：你总是在开会吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 可视化：一种更为直观的沟通方式.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/23%20%e5%8f%af%e8%a7%86%e5%8c%96%ef%bc%9a%e4%b8%80%e7%a7%8d%e6%9b%b4%e4%b8%ba%e7%9b%b4%e8%a7%82%e7%9a%84%e6%b2%9f%e9%80%9a%e6%96%b9%e5%bc%8f.md">23 可视化：一种更为直观的沟通方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 快速反馈：为什么你们公司总是做不好持续集成？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/24%20%e5%bf%ab%e9%80%9f%e5%8f%8d%e9%a6%88%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e4%bb%ac%e5%85%ac%e5%8f%b8%e6%80%bb%e6%98%af%e5%81%9a%e4%b8%8d%e5%a5%bd%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%ef%bc%9f.md">24 快速反馈：为什么你们公司总是做不好持续集成？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 开发中的问题一再出现，应该怎么办？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/25%20%e5%bc%80%e5%8f%91%e4%b8%ad%e7%9a%84%e9%97%ae%e9%a2%98%e4%b8%80%e5%86%8d%e5%87%ba%e7%8e%b0%ef%bc%8c%e5%ba%94%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">25 开发中的问题一再出现，应该怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 作为程序员，你也应该聆听用户声音.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/26%20%e4%bd%9c%e4%b8%ba%e7%a8%8b%e5%ba%8f%e5%91%98%ef%bc%8c%e4%bd%a0%e4%b9%9f%e5%ba%94%e8%af%a5%e8%81%86%e5%90%ac%e7%94%a8%e6%88%b7%e5%a3%b0%e9%9f%b3.md">26 作为程序员，你也应该聆听用户声音.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 尽早暴露问题： 为什么被指责的总是你？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/27%20%e5%b0%bd%e6%97%a9%e6%9a%b4%e9%9c%b2%e9%97%ae%e9%a2%98%ef%bc%9a%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a2%ab%e6%8c%87%e8%b4%a3%e7%9a%84%e6%80%bb%e6%98%af%e4%bd%a0%ef%bc%9f.md">27 尽早暴露问题： 为什么被指责的总是你？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 结构化：写文档也是一种学习方式.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/28%20%e7%bb%93%e6%9e%84%e5%8c%96%ef%bc%9a%e5%86%99%e6%96%87%e6%a1%a3%e4%b9%9f%e6%98%af%e4%b8%80%e7%a7%8d%e5%ad%a6%e4%b9%a0%e6%96%b9%e5%bc%8f.md">28 结构化：写文档也是一种学习方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 “懒惰”应该是所有程序员的骄傲.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/29%20%e2%80%9c%e6%87%92%e6%83%b0%e2%80%9d%e5%ba%94%e8%af%a5%e6%98%af%e6%89%80%e6%9c%89%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e9%aa%84%e5%82%b2.md">29 “懒惰”应该是所有程序员的骄傲.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 一个好的项目自动化应该是什么样子的？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/30%20%e4%b8%80%e4%b8%aa%e5%a5%bd%e7%9a%84%e9%a1%b9%e7%9b%ae%e8%87%aa%e5%8a%a8%e5%8c%96%e5%ba%94%e8%af%a5%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e5%ad%90%e7%9a%84%ef%bc%9f.md">30 一个好的项目自动化应该是什么样子的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 程序员怎么学习运维知识？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/31%20%e7%a8%8b%e5%ba%8f%e5%91%98%e6%80%8e%e4%b9%88%e5%ad%a6%e4%b9%a0%e8%bf%90%e7%bb%b4%e7%9f%a5%e8%af%86%ef%bc%9f.md">31 程序员怎么学习运维知识？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 持续交付：有持续集成就够了吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/32%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%ef%bc%9a%e6%9c%89%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%e5%b0%b1%e5%a4%9f%e4%ba%86%e5%90%97%ef%bc%9f.md">32 持续交付：有持续集成就够了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 如何做好验收测试？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/33%20%e5%a6%82%e4%bd%95%e5%81%9a%e5%a5%bd%e9%aa%8c%e6%94%b6%e6%b5%8b%e8%af%95%ef%bc%9f.md">33 如何做好验收测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 你的代码是怎么变混乱的？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/34%20%e4%bd%a0%e7%9a%84%e4%bb%a3%e7%a0%81%e6%98%af%e6%80%8e%e4%b9%88%e5%8f%98%e6%b7%b7%e4%b9%b1%e7%9a%84%ef%bc%9f.md">34 你的代码是怎么变混乱的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 总是在说MVC分层架构，但你真的理解分层吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/35%20%e6%80%bb%e6%98%af%e5%9c%a8%e8%af%b4MVC%e5%88%86%e5%b1%82%e6%9e%b6%e6%9e%84%ef%bc%8c%e4%bd%86%e4%bd%a0%e7%9c%9f%e7%9a%84%e7%90%86%e8%a7%a3%e5%88%86%e5%b1%82%e5%90%97%ef%bc%9f.md">35 总是在说MVC分层架构，但你真的理解分层吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 为什么总有人觉得5万块钱可以做一个淘宝？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/36%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%80%bb%e6%9c%89%e4%ba%ba%e8%a7%89%e5%be%975%e4%b8%87%e5%9d%97%e9%92%b1%e5%8f%af%e4%bb%a5%e5%81%9a%e4%b8%80%e4%b8%aa%e6%b7%98%e5%ae%9d%ef%bc%9f.md">36 为什么总有人觉得5万块钱可以做一个淘宝？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 先做好DDD再谈微服务吧，那只是一种部署形式.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/37%20%e5%85%88%e5%81%9a%e5%a5%bdDDD%e5%86%8d%e8%b0%88%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%90%a7%ef%bc%8c%e9%82%a3%e5%8f%aa%e6%98%af%e4%b8%80%e7%a7%8d%e9%83%a8%e7%bd%b2%e5%bd%a2%e5%bc%8f.md">37 先做好DDD再谈微服务吧，那只是一种部署形式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 新入职一家公司，怎么快速进入工作状态？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/38%20%e6%96%b0%e5%85%a5%e8%81%8c%e4%b8%80%e5%ae%b6%e5%85%ac%e5%8f%b8%ef%bc%8c%e6%80%8e%e4%b9%88%e5%bf%ab%e9%80%9f%e8%bf%9b%e5%85%a5%e5%b7%a5%e4%bd%9c%e7%8a%b6%e6%80%81%ef%bc%9f.md">38 新入职一家公司，怎么快速进入工作状态？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 面对遗留系统，你应该这样做.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/39%20%e9%9d%a2%e5%af%b9%e9%81%97%e7%95%99%e7%b3%bb%e7%bb%9f%ef%bc%8c%e4%bd%a0%e5%ba%94%e8%af%a5%e8%bf%99%e6%a0%b7%e5%81%9a.md">39 面对遗留系统，你应该这样做.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 我们应该如何保持竞争力？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/40%20%e6%88%91%e4%bb%ac%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e4%bf%9d%e6%8c%81%e7%ab%9e%e4%ba%89%e5%8a%9b%ef%bc%9f.md">40 我们应该如何保持竞争力？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 “综合运用”主题内容的全盘回顾.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%88%92%e9%87%8d%e7%82%b9%20%e2%80%9c%e7%bb%bc%e5%90%88%e8%bf%90%e7%94%a8%e2%80%9d%e4%b8%bb%e9%a2%98%e5%86%85%e5%ae%b9%e7%9a%84%e5%85%a8%e7%9b%98%e5%9b%9e%e9%a1%be.md">划重点 “综合运用”主题内容的全盘回顾.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 “自动化”主题的重点内容回顾汇总.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%88%92%e9%87%8d%e7%82%b9%20%e2%80%9c%e8%87%aa%e5%8a%a8%e5%8c%96%e2%80%9d%e4%b8%bb%e9%a2%98%e7%9a%84%e9%87%8d%e7%82%b9%e5%86%85%e5%ae%b9%e5%9b%9e%e9%a1%be%e6%b1%87%e6%80%bb.md">划重点 “自动化”主题的重点内容回顾汇总.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 一次关于“沟通反馈”主题内容的复盘.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%88%92%e9%87%8d%e7%82%b9%20%e4%b8%80%e6%ac%a1%e5%85%b3%e4%ba%8e%e2%80%9c%e6%b2%9f%e9%80%9a%e5%8f%8d%e9%a6%88%e2%80%9d%e4%b8%bb%e9%a2%98%e5%86%85%e5%ae%b9%e7%9a%84%e5%a4%8d%e7%9b%98.md">划重点 一次关于“沟通反馈”主题内容的复盘.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 关于“以终为始”，你要记住的9句话.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%88%92%e9%87%8d%e7%82%b9%20%e5%85%b3%e4%ba%8e%e2%80%9c%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%e2%80%9d%ef%bc%8c%e4%bd%a0%e8%a6%81%e8%ae%b0%e4%bd%8f%e7%9a%849%e5%8f%a5%e8%af%9d.md">划重点 关于“以终为始”，你要记住的9句话.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="划重点 关于“任务分解”，你要重点掌握哪些事？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%88%92%e9%87%8d%e7%82%b9%20%e5%85%b3%e4%ba%8e%e2%80%9c%e4%bb%bb%e5%8a%a1%e5%88%86%e8%a7%a3%e2%80%9d%ef%bc%8c%e4%bd%a0%e8%a6%81%e9%87%8d%e7%82%b9%e6%8e%8c%e6%8f%a1%e5%93%aa%e4%ba%9b%e4%ba%8b%ef%bc%9f.md">划重点 关于“任务分解”，你要重点掌握哪些事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 你真的了解重构吗？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e5%8a%a0%e9%a4%90%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e4%ba%86%e8%a7%a3%e9%87%8d%e6%9e%84%e5%90%97%ef%bc%9f.md">加餐 你真的了解重构吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="总复习 重新审视“最佳实践”.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e6%80%bb%e5%a4%8d%e4%b9%a0%20%e9%87%8d%e6%96%b0%e5%ae%a1%e8%a7%86%e2%80%9c%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5%e2%80%9d.md">总复习 重新审视“最佳实践”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="总复习 重新来“看书”.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e6%80%bb%e5%a4%8d%e4%b9%a0%20%e9%87%8d%e6%96%b0%e6%9d%a5%e2%80%9c%e7%9c%8b%e4%b9%a6%e2%80%9d.md">总复习 重新来“看书”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑解惑 如何分解一个你不了解的技术任务？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%20%e5%a6%82%e4%bd%95%e5%88%86%e8%a7%a3%e4%b8%80%e4%b8%aa%e4%bd%a0%e4%b8%8d%e4%ba%86%e8%a7%a3%e7%9a%84%e6%8a%80%e6%9c%af%e4%bb%bb%e5%8a%a1%ef%bc%9f.md">答疑解惑 如何分解一个你不了解的技术任务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑解惑 如何在实际工作中推行新观念？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%20%e5%a6%82%e4%bd%95%e5%9c%a8%e5%ae%9e%e9%99%85%e5%b7%a5%e4%bd%9c%e4%b8%ad%e6%8e%a8%e8%a1%8c%e6%96%b0%e8%a7%82%e5%bf%b5%ef%bc%9f.md">答疑解惑 如何在实际工作中推行新观念？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑解惑 如何管理你的上级？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%20%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e4%bd%a0%e7%9a%84%e4%b8%8a%e7%ba%a7%ef%bc%9f.md">答疑解惑 如何管理你的上级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑解惑 持续集成、持续交付，然后呢？.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%20%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%e3%80%81%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%ef%bc%8c%e7%84%b6%e5%90%8e%e5%91%a2%ef%bc%9f.md">答疑解惑 持续集成、持续交付，然后呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑解惑 持续集成，一条贯穿诸多实践的主线.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91%20%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%ef%bc%8c%e4%b8%80%e6%9d%a1%e8%b4%af%e7%a9%bf%e8%af%b8%e5%a4%9a%e5%ae%9e%e8%b7%b5%e7%9a%84%e4%b8%bb%e7%ba%bf.md">答疑解惑 持续集成，一条贯穿诸多实践的主线.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 少做事，才能更有效地工作.md" href="/%e4%b8%93%e6%a0%8f/10x%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%a5%e4%bd%9c%e6%b3%95/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%b0%91%e5%81%9a%e4%ba%8b%ef%bc%8c%e6%89%8d%e8%83%bd%e6%9b%b4%e6%9c%89%e6%95%88%e5%9c%b0%e5%b7%a5%e4%bd%9c.md">结束语 少做事，才能更有效地工作.md</a>
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
                            <h1 id="title" data-id="35 总是在说MVC分层架构，但你真的理解分层吗？" class="title">35 总是在说MVC分层架构，但你真的理解分层吗？</h1>
                            <div><p>你好，我是郑晔。</p>

<p>作为程序员，你一定听说过分层，比如，最常见的 Java 服务端应用的三层结构，在《<a href="http://time.geekbang.org/column/article/78542" target="_blank">15 | 一起练习：手把手带你分解任务</a>》中，我曾提到过：</p>

<ul>
<li>数据访问层，按照传统的说法，叫 DAO（Data Access Object，数据访问对象），按照领域驱动开发的术语，称之为 Repository；</li>
<li>服务层，提供应用服务；</li>
<li>资源层，提供对外访问的资源，采用传统做法就是 Controller。</li>
</ul>

<p>这几乎成为了写 Java 服务的标准模式。但不知道你有没有想过，为什么要分层呢？</p>

<h2 id="设计上的分解">设计上的分解</h2>

<p>其实，分层并不是一个特别符合直觉的做法，符合直觉的做法应该是直接写在一起。</p>

<p>在编程框架还不是特别流行的时候，人们就是直接把页面和逻辑混在一起写的。如果你有机会看看写得不算理想的 PHP 程序，这种现象还是大概率会出现的。</p>

<p>即便像 Java 这个如此重视架构的社区，分层也是很久之后才出现的，早期的 JSP 和 PHP 并没有什么本质区别。</p>

<p>那为什么要分层呢？原因很简单，当代码复杂到一定程度，人们维护代码的难度就急剧上升。一旦出现任何问题，在所有一切都混在一起的代码中定位问题，本质上就是一个“大海捞针”的活。</p>

<p>前面讲任务分解的时候，我不断在强调的观点就是，人们擅长解决的是小问题，大问题怎么办？拆小了就好。</p>

<p><strong>分层架构，实际上，就是一种在设计上的分解。</strong></p>

<p>回到前面所说的三层架构，这是行业中最早普及的一种架构模式，最开始是 MVC，也就是 Model、View 和 Controller。</p>

<p>MVC 的概念起源于 GUI （Graphical User Interface，图形用户界面）编程，人们希望将图形界面上展示的部分（View）与 UI 的数据模型（Model）分开，它们之间的联动由 Controller 负责。这个概念在 GUI 编程中是没有问题的，但也仅限于在与 UI 有交互的部分。</p>

<p>很多人误以为这也适合服务端程序，他们就把模型部分误解成了数据库里的模型，甚至把它理解成数据库访问。于是，你会看到有人在 Controller 里访问数据库。</p>

<p>不知道你是不是了解 <a href="http://rubyonrails.org" target="_blank">Ruby on Rails</a>，这是当年改变了行业认知的一个 Web 开发框架，带来很多颠覆性的做法。它采用的就是这样一种编程模型。当年写 Rails 程序的时候我发现，当业务复杂到了一定规模，代码就开始难以维护了。我想了好久，终于发现，在 Rails 的常规做法中少了服务层（Service）的设计。</p>

<p>这个问题在 Java 领域，爆发得要比 Rails 里早，因为 Ruby 语言的优越性，Rails 实现的数据访问非常优雅。正是因为 Rails 的数据访问实在太容易了，很多服务实际上写到 Model 层里。在代码规模不大时，代码看上去是不复杂的，甚至还有些优雅。</p>

<p>而那时的 Java 可是要一行一行地写数据访问，所以，代码不太可能放在 Model 层，而放在Controller 里也会让代码变复杂，于是，为业务逻辑而生的 Service 层就呼之欲出了。</p>

<p>至此，常见的 Java 服务端开发的基础就全部成型了，只不过，由于后来 REST 服务的兴起，资源层替代了 Controller 层。</p>

<p>到这里，我给你讲了常见的 Java 服务三层架构的来龙去脉。但实际上，在软件开发中，分层几乎是无处不在的，因为好的分层往往需要有好的抽象。</p>

<h2 id="无处不在的分层">无处不在的分层</h2>

<p>作为程序员，我们几乎每天都在与分层打交道。比如说，程序员都对网络编程模型很熟悉，无论是 ISO 的七层还是 TCP/IP 的五层。</p>

<p>但不知道你有没有发现，虽然学习的时候，你要学习网络有那么多层，但在使用的时候，大多数情况下，你只要了解最上面的那层，比如，HTTP。</p>

<p>很多人对底层的协议的理解几乎就停留在“学过”的水平上，因为在大多数情况下，除非你要写协议栈，不然你很难用得到。即便偶尔用到，90%的问题靠搜索引擎就解决了，你也很少有动力去系统学习。</p>

<p>之所以你可以这么放心大胆地“忽略”底层协议，一个关键点就在于，网络模型的分层架构实现得太好了，好到你作为上层的使用者几乎可以忽略底层。而这正是分层真正的价值：<strong>构建一个良好的抽象。</strong></p>

<p>这种构建良好的抽象在软件开发中随处可见，比如，你作为一个程序员，每天写着在 CPU 上运行的代码，但你读过指令集吗？你之所以可以不去了解，是因为已经有编译器做好了分层，让你可以只用它们构建出的“抽象”——编程语言去思考问题。</p>

<p>比如，每天写着 Java 程序的程序员，你知道 Java 程序是如何管理内存的吗？这可是令很多 C/C++程序员寝食难安的问题，而你之所以不用关心这些，正是托了 Java 这种“抽象”的福。对了，你甚至可能没有注意到编程语言也是一种抽象。</p>

<h2 id="有抽象有发展">有抽象有发展</h2>

<p>只有构建起抽象，人们才能在此基础上做出更复杂的东西。如果今天的游戏依然是面向显示屏的像素编程，那么，精彩的游戏视觉效果就只能由极少数真正的高手来开发。我们今天的大部分游戏应该依然停留在《超级玛丽》的水准。</p>

<p>同样，近些年前端领域风起云涌，但你是否想过，为什么 Web 的概念早就出现了，但前端作为一个专门的职位，真正的蓬勃发展却是最近十年的事？</p>

<p>2009年，Ryan Dahl 发布了Node.js，人们才真正认识到，原来 JavaScript 不仅仅可以用于浏览器，还能做服务器开发。</p>

<p>于是，JavaScript 社区大发展，各种在其他社区已经很流行的工具终于在 JavaScript 世界中发展了起来。正是有了这些工具的支持，人们才能用 JavaScript 构建更复杂的工程，前端领域才能得到了极大的发展。</p>

<p>如今，JavaScript 已经发展成唯一一门全平台语言，当然，发展最好的依然是在它的大本营：前端领域。前端程序员才有了今天幸福的烦恼：各种前端框架层出不穷。</p>

<p>在这里，Node.js 的出现让 JavaScript 成为了一个更好的抽象。</p>

<h2 id="构建你的抽象">构建你的抽象</h2>

<p>理解了分层实际上是在构建抽象，你或许会关心，我该怎么把它运用在自己的工作中。</p>

<p>构建抽象，最核心的一步是构建出你的核心模型。什么是核心模型呢？就是表达你业务的那部分代码，换句话说，别的东西都可以变，但这部分不能变。</p>

<p>这么说可能还是有点抽象，我们回到前面的三层架构。</p>

<p>在前面介绍三层架构的演变时，提到了一个变迁：REST服务的兴起，让 Controller 逐渐退出了历史舞台，资源层取而代之。</p>

<p>换句话说，访问服务的方式可能会变。放到计算机编程的发展中，这种趋势就更明显了，从命令行到网络，从 CS（Client-Server） 到 BS（Browser-Server），从浏览器到移动端。所以，怎么访问不应该是你关注的核心。</p>

<p>同样， 关系型数据库也不是你关注的核心，它只是今天的主流而已。从前用文件，今天还有各种 NoSQL。</p>

<p>如此说来，三层架构中的两层重要性都不是那么高，那重要的是什么？答案便呼之欲出了，没错，就是剩下的部分，我们习惯上称之为服务层，但这个名字其实不能很好地反映它的作用，更恰当的说法应该可以叫领域模型（Domain Model）。</p>

<p>它便是我们的核心模型，也是我们在做软件设计时，真正应该着力的地方。</p>

<p>为什么叫“服务层”不是一个好的说法呢？这里会遗漏领域模型中一个重要的组成部分：领域对象。</p>

<p>很多人理解领域对象有一个严重的误区，认为领域对象属于数据层。数据存储只是领域对象的一种用途，它更重要的用途还是用在各种领域服务中。</p>

<p>由此还能引出另一个常见的设计错误，领域对象中只包含数据访问，也就是常说的 getter 和 setter，而没有任何逻辑。</p>

<p>如果只用于数据存储，只有数据访问就够了，但如果是领域对象，就应该有业务逻辑。比如，给一个用户修改密码，用户这个对象上应该有一个 changePassword 方法，而不是每次去 setPassword。</p>

<p>严格地说，领域对象和存储对象应该是两个类，只不过它俩实在太像了，很多人经常使用一个类，这还是个小问题。但很多人却把这种内部方案用到了外部，比如，第三方集成。</p>

<p>为数不少的团队都在自己的业务代码中直接使用了第三方代码中的对象，第三方的任何修改都会让你的代码跟着改，你的团队就只能疲于奔命。</p>

<p>解决这个问题最好的办法就是把它们分开，<strong>你的领域层只依赖于你的领域对象，第三方发过来的内容先做一次转换，转换成你的领域对象</strong>。这种做法称为防腐层。</p>

<p>当我们把领域模型看成了整个设计的核心，看待其他层的视角也会随之转变，它们只不过是适配到不同地方的一种方式而已，而这种理念的推广，就是一些人在说的六边形架构。</p>

<p><img src="assets/5b87222af8b840a08f5efa6820c0b215.jpg" alt="" /></p>

<p>怎么设计好领域模型是一个庞大的主题，推荐你去了解一下领域驱动设计（Domain Driven Design，DDD），这个话题我们后面还会再次提到。</p>

<p>讨论其实还可以继续延伸下去，已经构建好的领域模型怎么更好地提供给其他部分使用呢？一个好的做法是封装成领域特定语言（Domain Specific Language，DSL）。当然，这也是一个庞大的话题，就不继续展开了。</p>

<h2 id="总结时刻">总结时刻</h2>

<p>我从最常见的服务端三层架构入手，给你讲了它们的来龙去脉。分层架构实际是一种设计上的分解，将不同的内容放在不同的地方，降低软件开发和维护的成本。</p>

<p>分层，更关键的是，提供抽象。这种分层抽象在计算机领域无处不在，无论是编程语言，还是网络协议，都体现着分层抽象的价值。有了分层抽象，人们才能更好地在抽象的基础上构建更复杂的东西。</p>

<p>在日常工作中，我们应该把精力重点放在构建自己的领域模型上，因为它才是工作最核心、不易变的东西。</p>

<p>如果今天的内容你只能记住一件事，那请记住：<strong>构建好你的领域模型。</strong></p>

<p>最后我想请你思考一下，你还知道哪些技术是体现分层抽象的思想吗？欢迎在留言区写下你的想法。</p>

<p>感谢阅读，如果你觉得这篇文章对你有帮助的话，也欢迎把它分享给你的朋友。</p>
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
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0cdfb139bf94de',t:'MTczMzk5NzY5My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
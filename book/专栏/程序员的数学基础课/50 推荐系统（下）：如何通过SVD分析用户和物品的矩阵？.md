<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=50&#32;推荐系统（下）：如何通过SVD分析用户和物品的矩阵？>
        <link rel="icon" href="/static/favicon.png">
        <title>50 推荐系统（下）：如何通过SVD分析用户和物品的矩阵？ </title>
        
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
                        <a class="menu-item" id="01 二进制：不了解计算机的源头，你学什么编程.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/01%20%e4%ba%8c%e8%bf%9b%e5%88%b6%ef%bc%9a%e4%b8%8d%e4%ba%86%e8%a7%a3%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%9a%84%e6%ba%90%e5%a4%b4%ef%bc%8c%e4%bd%a0%e5%ad%a6%e4%bb%80%e4%b9%88%e7%bc%96%e7%a8%8b.md">01 二进制：不了解计算机的源头，你学什么编程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 余数：原来取余操作本身就是个哈希函数.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/02%20%e4%bd%99%e6%95%b0%ef%bc%9a%e5%8e%9f%e6%9d%a5%e5%8f%96%e4%bd%99%e6%93%8d%e4%bd%9c%e6%9c%ac%e8%ba%ab%e5%b0%b1%e6%98%af%e4%b8%aa%e5%93%88%e5%b8%8c%e5%87%bd%e6%95%b0.md">02 余数：原来取余操作本身就是个哈希函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 迭代法：不用编程语言的自带函数，你会如何计算平方根？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/03%20%e8%bf%ad%e4%bb%a3%e6%b3%95%ef%bc%9a%e4%b8%8d%e7%94%a8%e7%bc%96%e7%a8%8b%e8%af%ad%e8%a8%80%e7%9a%84%e8%87%aa%e5%b8%a6%e5%87%bd%e6%95%b0%ef%bc%8c%e4%bd%a0%e4%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a1%e7%ae%97%e5%b9%b3%e6%96%b9%e6%a0%b9%ef%bc%9f.md">03 迭代法：不用编程语言的自带函数，你会如何计算平方根？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 数学归纳法：如何用数学归纳提升代码的运行效率？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/04%20%e6%95%b0%e5%ad%a6%e5%bd%92%e7%ba%b3%e6%b3%95%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e6%95%b0%e5%ad%a6%e5%bd%92%e7%ba%b3%e6%8f%90%e5%8d%87%e4%bb%a3%e7%a0%81%e7%9a%84%e8%bf%90%e8%a1%8c%e6%95%88%e7%8e%87%ef%bc%9f.md">04 数学归纳法：如何用数学归纳提升代码的运行效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 递归（上）：泛化数学归纳，如何将复杂问题简单化？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/05%20%e9%80%92%e5%bd%92%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%b3%9b%e5%8c%96%e6%95%b0%e5%ad%a6%e5%bd%92%e7%ba%b3%ef%bc%8c%e5%a6%82%e4%bd%95%e5%b0%86%e5%a4%8d%e6%9d%82%e9%97%ae%e9%a2%98%e7%ae%80%e5%8d%95%e5%8c%96%ef%bc%9f.md">05 递归（上）：泛化数学归纳，如何将复杂问题简单化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 递归（下）：分而治之，从归并排序到MapReduce.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/06%20%e9%80%92%e5%bd%92%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%88%86%e8%80%8c%e6%b2%bb%e4%b9%8b%ef%bc%8c%e4%bb%8e%e5%bd%92%e5%b9%b6%e6%8e%92%e5%ba%8f%e5%88%b0MapReduce.md">06 递归（下）：分而治之，从归并排序到MapReduce.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 排列：如何让计算机学会“田忌赛马”？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/07%20%e6%8e%92%e5%88%97%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%ad%a6%e4%bc%9a%e2%80%9c%e7%94%b0%e5%bf%8c%e8%b5%9b%e9%a9%ac%e2%80%9d%ef%bc%9f.md">07 排列：如何让计算机学会“田忌赛马”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 组合：如何让计算机安排世界杯的赛程？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/08%20%e7%bb%84%e5%90%88%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%ae%89%e6%8e%92%e4%b8%96%e7%95%8c%e6%9d%af%e7%9a%84%e8%b5%9b%e7%a8%8b%ef%bc%9f.md">08 组合：如何让计算机安排世界杯的赛程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 动态规划（上）：如何实现基于编辑距离的查询推荐？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/09%20%e5%8a%a8%e6%80%81%e8%a7%84%e5%88%92%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%9f%ba%e4%ba%8e%e7%bc%96%e8%be%91%e8%b7%9d%e7%a6%bb%e7%9a%84%e6%9f%a5%e8%af%a2%e6%8e%a8%e8%8d%90%ef%bc%9f.md">09 动态规划（上）：如何实现基于编辑距离的查询推荐？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 动态规划（下）：如何求得状态转移方程并进行编程实现？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/10%20%e5%8a%a8%e6%80%81%e8%a7%84%e5%88%92%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e6%b1%82%e5%be%97%e7%8a%b6%e6%80%81%e8%bd%ac%e7%a7%bb%e6%96%b9%e7%a8%8b%e5%b9%b6%e8%bf%9b%e8%a1%8c%e7%bc%96%e7%a8%8b%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">10 动态规划（下）：如何求得状态转移方程并进行编程实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 树的深度优先搜索（上）：如何才能高效率地查字典？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/11%20%e6%a0%91%e7%9a%84%e6%b7%b1%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e9%ab%98%e6%95%88%e7%8e%87%e5%9c%b0%e6%9f%a5%e5%ad%97%e5%85%b8%ef%bc%9f.md">11 树的深度优先搜索（上）：如何才能高效率地查字典？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 树的深度优先搜索（下）：如何才能高效率地查字典？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/12%20%e6%a0%91%e7%9a%84%e6%b7%b1%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e9%ab%98%e6%95%88%e7%8e%87%e5%9c%b0%e6%9f%a5%e5%ad%97%e5%85%b8%ef%bc%9f.md">12 树的深度优先搜索（下）：如何才能高效率地查字典？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 树的广度优先搜索（上）：人际关系的六度理论是真的吗？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/13%20%e6%a0%91%e7%9a%84%e5%b9%bf%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%ba%ba%e9%99%85%e5%85%b3%e7%b3%bb%e7%9a%84%e5%85%ad%e5%ba%a6%e7%90%86%e8%ae%ba%e6%98%af%e7%9c%9f%e7%9a%84%e5%90%97%ef%bc%9f.md">13 树的广度优先搜索（上）：人际关系的六度理论是真的吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 树的广度优先搜索（下）：为什么双向广度优先搜索的效率更高？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/14%20%e6%a0%91%e7%9a%84%e5%b9%bf%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%8f%8c%e5%90%91%e5%b9%bf%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2%e7%9a%84%e6%95%88%e7%8e%87%e6%9b%b4%e9%ab%98%ef%bc%9f.md">14 树的广度优先搜索（下）：为什么双向广度优先搜索的效率更高？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 从树到图：如何让计算机学会看地图？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/15%20%e4%bb%8e%e6%a0%91%e5%88%b0%e5%9b%be%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%ad%a6%e4%bc%9a%e7%9c%8b%e5%9c%b0%e5%9b%be%ef%bc%9f.md">15 从树到图：如何让计算机学会看地图？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 时间和空间复杂度（上）：优化性能是否只是“纸上谈兵”？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/16%20%e6%97%b6%e9%97%b4%e5%92%8c%e7%a9%ba%e9%97%b4%e5%a4%8d%e6%9d%82%e5%ba%a6%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%bc%98%e5%8c%96%e6%80%a7%e8%83%bd%e6%98%af%e5%90%a6%e5%8f%aa%e6%98%af%e2%80%9c%e7%ba%b8%e4%b8%8a%e8%b0%88%e5%85%b5%e2%80%9d%ef%bc%9f.md">16 时间和空间复杂度（上）：优化性能是否只是“纸上谈兵”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 时间和空间复杂度（下）：如何使用六个法则进行复杂度分析？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/17%20%e6%97%b6%e9%97%b4%e5%92%8c%e7%a9%ba%e9%97%b4%e5%a4%8d%e6%9d%82%e5%ba%a6%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%85%ad%e4%b8%aa%e6%b3%95%e5%88%99%e8%bf%9b%e8%a1%8c%e5%a4%8d%e6%9d%82%e5%ba%a6%e5%88%86%e6%9e%90%ef%bc%9f.md">17 时间和空间复杂度（下）：如何使用六个法则进行复杂度分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 总结课：数据结构、编程语句和基础算法体现了哪些数学思想？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/18%20%e6%80%bb%e7%bb%93%e8%af%be%ef%bc%9a%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e3%80%81%e7%bc%96%e7%a8%8b%e8%af%ad%e5%8f%a5%e5%92%8c%e5%9f%ba%e7%a1%80%e7%ae%97%e6%b3%95%e4%bd%93%e7%8e%b0%e4%ba%86%e5%93%aa%e4%ba%9b%e6%95%b0%e5%ad%a6%e6%80%9d%e6%83%b3%ef%bc%9f.md">18 总结课：数据结构、编程语句和基础算法体现了哪些数学思想？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 概率和统计：编程为什么需要概率和统计？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/19%20%e6%a6%82%e7%8e%87%e5%92%8c%e7%bb%9f%e8%ae%a1%ef%bc%9a%e7%bc%96%e7%a8%8b%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e6%a6%82%e7%8e%87%e5%92%8c%e7%bb%9f%e8%ae%a1%ef%bc%9f.md">19 概率和统计：编程为什么需要概率和统计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 概率基础（上）：一篇文章帮你理解随机变量、概率分布和期望值.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/20%20%e6%a6%82%e7%8e%87%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%b8%80%e7%af%87%e6%96%87%e7%ab%a0%e5%b8%ae%e4%bd%a0%e7%90%86%e8%a7%a3%e9%9a%8f%e6%9c%ba%e5%8f%98%e9%87%8f%e3%80%81%e6%a6%82%e7%8e%87%e5%88%86%e5%b8%83%e5%92%8c%e6%9c%9f%e6%9c%9b%e5%80%bc.md">20 概率基础（上）：一篇文章帮你理解随机变量、概率分布和期望值.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 概率基础（下）：联合概率、条件概率和贝叶斯法则，这些概率公式究竟能做什么？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/21%20%e6%a6%82%e7%8e%87%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%81%94%e5%90%88%e6%a6%82%e7%8e%87%e3%80%81%e6%9d%a1%e4%bb%b6%e6%a6%82%e7%8e%87%e5%92%8c%e8%b4%9d%e5%8f%b6%e6%96%af%e6%b3%95%e5%88%99%ef%bc%8c%e8%bf%99%e4%ba%9b%e6%a6%82%e7%8e%87%e5%85%ac%e5%bc%8f%e7%a9%b6%e7%ab%9f%e8%83%bd%e5%81%9a%e4%bb%80%e4%b9%88%ef%bc%9f.md">21 概率基础（下）：联合概率、条件概率和贝叶斯法则，这些概率公式究竟能做什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 朴素贝叶斯：如何让计算机学会自动分类？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/22%20%e6%9c%b4%e7%b4%a0%e8%b4%9d%e5%8f%b6%e6%96%af%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%ad%a6%e4%bc%9a%e8%87%aa%e5%8a%a8%e5%88%86%e7%b1%bb%ef%bc%9f.md">22 朴素贝叶斯：如何让计算机学会自动分类？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 文本分类：如何区分特定类型的新闻？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/23%20%e6%96%87%e6%9c%ac%e5%88%86%e7%b1%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e5%8c%ba%e5%88%86%e7%89%b9%e5%ae%9a%e7%b1%bb%e5%9e%8b%e7%9a%84%e6%96%b0%e9%97%bb%ef%bc%9f.md">23 文本分类：如何区分特定类型的新闻？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 语言模型：如何使用链式法则和马尔科夫假设简化概率模型？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/24%20%e8%af%ad%e8%a8%80%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e9%93%be%e5%bc%8f%e6%b3%95%e5%88%99%e5%92%8c%e9%a9%ac%e5%b0%94%e7%a7%91%e5%a4%ab%e5%81%87%e8%ae%be%e7%ae%80%e5%8c%96%e6%a6%82%e7%8e%87%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">24 语言模型：如何使用链式法则和马尔科夫假设简化概率模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 马尔科夫模型：从PageRank到语音识别，背后是什么模型在支撑？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/25%20%e9%a9%ac%e5%b0%94%e7%a7%91%e5%a4%ab%e6%a8%a1%e5%9e%8b%ef%bc%9a%e4%bb%8ePageRank%e5%88%b0%e8%af%ad%e9%9f%b3%e8%af%86%e5%88%ab%ef%bc%8c%e8%83%8c%e5%90%8e%e6%98%af%e4%bb%80%e4%b9%88%e6%a8%a1%e5%9e%8b%e5%9c%a8%e6%94%af%e6%92%91%ef%bc%9f.md">25 马尔科夫模型：从PageRank到语音识别，背后是什么模型在支撑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 信息熵：如何通过几个问题，测出你对应的武侠人物？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/26%20%e4%bf%a1%e6%81%af%e7%86%b5%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e5%87%a0%e4%b8%aa%e9%97%ae%e9%a2%98%ef%bc%8c%e6%b5%8b%e5%87%ba%e4%bd%a0%e5%af%b9%e5%ba%94%e7%9a%84%e6%ad%a6%e4%be%a0%e4%ba%ba%e7%89%a9%ef%bc%9f.md">26 信息熵：如何通过几个问题，测出你对应的武侠人物？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 决策树：信息增益、增益比率和基尼指数的运用.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/27%20%e5%86%b3%e7%ad%96%e6%a0%91%ef%bc%9a%e4%bf%a1%e6%81%af%e5%a2%9e%e7%9b%8a%e3%80%81%e5%a2%9e%e7%9b%8a%e6%af%94%e7%8e%87%e5%92%8c%e5%9f%ba%e5%b0%bc%e6%8c%87%e6%95%b0%e7%9a%84%e8%bf%90%e7%94%a8.md">27 决策树：信息增益、增益比率和基尼指数的运用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 熵、信息增益和卡方：如何寻找关键特征？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/28%20%e7%86%b5%e3%80%81%e4%bf%a1%e6%81%af%e5%a2%9e%e7%9b%8a%e5%92%8c%e5%8d%a1%e6%96%b9%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%bb%e6%89%be%e5%85%b3%e9%94%ae%e7%89%b9%e5%be%81%ef%bc%9f.md">28 熵、信息增益和卡方：如何寻找关键特征？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 归一化和标准化：各种特征如何综合才是最合理的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/29%20%e5%bd%92%e4%b8%80%e5%8c%96%e5%92%8c%e6%a0%87%e5%87%86%e5%8c%96%ef%bc%9a%e5%90%84%e7%a7%8d%e7%89%b9%e5%be%81%e5%a6%82%e4%bd%95%e7%bb%bc%e5%90%88%e6%89%8d%e6%98%af%e6%9c%80%e5%90%88%e7%90%86%e7%9a%84%ef%bc%9f.md">29 归一化和标准化：各种特征如何综合才是最合理的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 统计意义（上）：如何通过显著性检验，判断你的A_B测试结果是不是巧合？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/30%20%e7%bb%9f%e8%ae%a1%e6%84%8f%e4%b9%89%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%98%be%e8%91%97%e6%80%a7%e6%a3%80%e9%aa%8c%ef%bc%8c%e5%88%a4%e6%96%ad%e4%bd%a0%e7%9a%84A_B%e6%b5%8b%e8%af%95%e7%bb%93%e6%9e%9c%e6%98%af%e4%b8%8d%e6%98%af%e5%b7%a7%e5%90%88%ef%bc%9f.md">30 统计意义（上）：如何通过显著性检验，判断你的A_B测试结果是不是巧合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 统计意义（下）：如何通过显著性检验，判断你的A_B测试结果是不是巧合？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/31%20%e7%bb%9f%e8%ae%a1%e6%84%8f%e4%b9%89%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%98%be%e8%91%97%e6%80%a7%e6%a3%80%e9%aa%8c%ef%bc%8c%e5%88%a4%e6%96%ad%e4%bd%a0%e7%9a%84A_B%e6%b5%8b%e8%af%95%e7%bb%93%e6%9e%9c%e6%98%af%e4%b8%8d%e6%98%af%e5%b7%a7%e5%90%88%ef%bc%9f.md">31 统计意义（下）：如何通过显著性检验，判断你的A_B测试结果是不是巧合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 概率统计篇答疑和总结：为什么会有欠拟合和过拟合？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/32%20%e6%a6%82%e7%8e%87%e7%bb%9f%e8%ae%a1%e7%af%87%e7%ad%94%e7%96%91%e5%92%8c%e6%80%bb%e7%bb%93%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bc%9a%e6%9c%89%e6%ac%a0%e6%8b%9f%e5%90%88%e5%92%8c%e8%bf%87%e6%8b%9f%e5%90%88%ef%bc%9f.md">32 概率统计篇答疑和总结：为什么会有欠拟合和过拟合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 线性代数：线性代数到底都讲了些什么？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/33%20%e7%ba%bf%e6%80%a7%e4%bb%a3%e6%95%b0%ef%bc%9a%e7%ba%bf%e6%80%a7%e4%bb%a3%e6%95%b0%e5%88%b0%e5%ba%95%e9%83%bd%e8%ae%b2%e4%ba%86%e4%ba%9b%e4%bb%80%e4%b9%88%ef%bc%9f.md">33 线性代数：线性代数到底都讲了些什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 向量空间模型：如何让计算机理解现实事物之间的关系？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/34%20%e5%90%91%e9%87%8f%e7%a9%ba%e9%97%b4%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e7%90%86%e8%a7%a3%e7%8e%b0%e5%ae%9e%e4%ba%8b%e7%89%a9%e4%b9%8b%e9%97%b4%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">34 向量空间模型：如何让计算机理解现实事物之间的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 文本检索：如何让计算机处理自然语言？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/35%20%e6%96%87%e6%9c%ac%e6%a3%80%e7%b4%a2%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%a4%84%e7%90%86%e8%87%aa%e7%84%b6%e8%af%ad%e8%a8%80%ef%bc%9f.md">35 文本检索：如何让计算机处理自然语言？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 文本聚类：如何过滤冗余的新闻？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/36%20%e6%96%87%e6%9c%ac%e8%81%9a%e7%b1%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%87%e6%bb%a4%e5%86%97%e4%bd%99%e7%9a%84%e6%96%b0%e9%97%bb%ef%bc%9f.md">36 文本聚类：如何过滤冗余的新闻？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 矩阵（上）：如何使用矩阵操作进行PageRank计算？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/37%20%e7%9f%a9%e9%98%b5%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e7%9f%a9%e9%98%b5%e6%93%8d%e4%bd%9c%e8%bf%9b%e8%a1%8cPageRank%e8%ae%a1%e7%ae%97%ef%bc%9f.md">37 矩阵（上）：如何使用矩阵操作进行PageRank计算？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 矩阵（下）：如何使用矩阵操作进行协同过滤推荐？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/38%20%e7%9f%a9%e9%98%b5%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e7%9f%a9%e9%98%b5%e6%93%8d%e4%bd%9c%e8%bf%9b%e8%a1%8c%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%e6%8e%a8%e8%8d%90%ef%bc%9f.md">38 矩阵（下）：如何使用矩阵操作进行协同过滤推荐？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 线性回归（上）：如何使用高斯消元求解线性方程组？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/39%20%e7%ba%bf%e6%80%a7%e5%9b%9e%e5%bd%92%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e9%ab%98%e6%96%af%e6%b6%88%e5%85%83%e6%b1%82%e8%a7%a3%e7%ba%bf%e6%80%a7%e6%96%b9%e7%a8%8b%e7%bb%84%ef%bc%9f.md">39 线性回归（上）：如何使用高斯消元求解线性方程组？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 线性回归（中）：如何使用最小二乘法进行直线拟合？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/40%20%e7%ba%bf%e6%80%a7%e5%9b%9e%e5%bd%92%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e6%9c%80%e5%b0%8f%e4%ba%8c%e4%b9%98%e6%b3%95%e8%bf%9b%e8%a1%8c%e7%9b%b4%e7%ba%bf%e6%8b%9f%e5%90%88%ef%bc%9f.md">40 线性回归（中）：如何使用最小二乘法进行直线拟合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 线性回归（下）：如何使用最小二乘法进行效果验证？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/41%20%e7%ba%bf%e6%80%a7%e5%9b%9e%e5%bd%92%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e6%9c%80%e5%b0%8f%e4%ba%8c%e4%b9%98%e6%b3%95%e8%bf%9b%e8%a1%8c%e6%95%88%e6%9e%9c%e9%aa%8c%e8%af%81%ef%bc%9f.md">41 线性回归（下）：如何使用最小二乘法进行效果验证？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 PCA主成分分析（上）：如何利用协方差矩阵来降维？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/42%20PCA%e4%b8%bb%e6%88%90%e5%88%86%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e5%8d%8f%e6%96%b9%e5%b7%ae%e7%9f%a9%e9%98%b5%e6%9d%a5%e9%99%8d%e7%bb%b4%ef%bc%9f.md">42 PCA主成分分析（上）：如何利用协方差矩阵来降维？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 PCA主成分分析（下）：为什么要计算协方差矩阵的特征值和特征向量？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/43%20PCA%e4%b8%bb%e6%88%90%e5%88%86%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e8%ae%a1%e7%ae%97%e5%8d%8f%e6%96%b9%e5%b7%ae%e7%9f%a9%e9%98%b5%e7%9a%84%e7%89%b9%e5%be%81%e5%80%bc%e5%92%8c%e7%89%b9%e5%be%81%e5%90%91%e9%87%8f%ef%bc%9f.md">43 PCA主成分分析（下）：为什么要计算协方差矩阵的特征值和特征向量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 奇异值分解：如何挖掘潜在的语义关系？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/44%20%e5%a5%87%e5%bc%82%e5%80%bc%e5%88%86%e8%a7%a3%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8c%96%e6%8e%98%e6%bd%9c%e5%9c%a8%e7%9a%84%e8%af%ad%e4%b9%89%e5%85%b3%e7%b3%bb%ef%bc%9f.md">44 奇异值分解：如何挖掘潜在的语义关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 线性代数篇答疑和总结：矩阵乘法的几何意义是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/45%20%e7%ba%bf%e6%80%a7%e4%bb%a3%e6%95%b0%e7%af%87%e7%ad%94%e7%96%91%e5%92%8c%e6%80%bb%e7%bb%93%ef%bc%9a%e7%9f%a9%e9%98%b5%e4%b9%98%e6%b3%95%e7%9a%84%e5%87%a0%e4%bd%95%e6%84%8f%e4%b9%89%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">45 线性代数篇答疑和总结：矩阵乘法的几何意义是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 缓存系统：如何通过哈希表和队列实现高效访问？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/46%20%e7%bc%93%e5%ad%98%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e5%93%88%e5%b8%8c%e8%a1%a8%e5%92%8c%e9%98%9f%e5%88%97%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e8%ae%bf%e9%97%ae%ef%bc%9f.md">46 缓存系统：如何通过哈希表和队列实现高效访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 搜索引擎（上）：如何通过倒排索引和向量空间模型，打造一个简单的搜索引擎？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/47%20%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e5%80%92%e6%8e%92%e7%b4%a2%e5%bc%95%e5%92%8c%e5%90%91%e9%87%8f%e7%a9%ba%e9%97%b4%e6%a8%a1%e5%9e%8b%ef%bc%8c%e6%89%93%e9%80%a0%e4%b8%80%e4%b8%aa%e7%ae%80%e5%8d%95%e7%9a%84%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%ef%bc%9f.md">47 搜索引擎（上）：如何通过倒排索引和向量空间模型，打造一个简单的搜索引擎？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 搜索引擎（下）：如何通过查询的分类，让电商平台的搜索结果更相关？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/48%20%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%9f%a5%e8%af%a2%e7%9a%84%e5%88%86%e7%b1%bb%ef%bc%8c%e8%ae%a9%e7%94%b5%e5%95%86%e5%b9%b3%e5%8f%b0%e7%9a%84%e6%90%9c%e7%b4%a2%e7%bb%93%e6%9e%9c%e6%9b%b4%e7%9b%b8%e5%85%b3%ef%bc%9f.md">48 搜索引擎（下）：如何通过查询的分类，让电商平台的搜索结果更相关？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 推荐系统（上）：如何实现基于相似度的协同过滤？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/49%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%9f%ba%e4%ba%8e%e7%9b%b8%e4%bc%bc%e5%ba%a6%e7%9a%84%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%ef%bc%9f.md">49 推荐系统（上）：如何实现基于相似度的协同过滤？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 推荐系统（下）：如何通过SVD分析用户和物品的矩阵？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/50%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87SVD%e5%88%86%e6%9e%90%e7%94%a8%e6%88%b7%e5%92%8c%e7%89%a9%e5%93%81%e7%9a%84%e7%9f%a9%e9%98%b5%ef%bc%9f.md">50 推荐系统（下）：如何通过SVD分析用户和物品的矩阵？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 综合应用篇答疑和总结：如何进行个性化用户画像的设计？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/51%20%e7%bb%bc%e5%90%88%e5%ba%94%e7%94%a8%e7%af%87%e7%ad%94%e7%96%91%e5%92%8c%e6%80%bb%e7%bb%93%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e4%b8%aa%e6%80%a7%e5%8c%96%e7%94%a8%e6%88%b7%e7%94%bb%e5%83%8f%e7%9a%84%e8%ae%be%e8%ae%a1%ef%bc%9f.md">51 综合应用篇答疑和总结：如何进行个性化用户画像的设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="导读：程序员应该怎么学数学？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e5%af%bc%e8%af%bb%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%91%98%e5%ba%94%e8%af%a5%e6%80%8e%e4%b9%88%e5%ad%a6%e6%95%b0%e5%ad%a6%ef%bc%9f.md">导读：程序员应该怎么学数学？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="开篇词 作为程序员，为什么你应该学好数学？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bd%9c%e4%b8%ba%e7%a8%8b%e5%ba%8f%e5%91%98%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e5%ba%94%e8%af%a5%e5%ad%a6%e5%a5%bd%e6%95%b0%e5%ad%a6%ef%bc%9f.md">开篇词 作为程序员，为什么你应该学好数学？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="数学专栏课外加餐（一） 我们为什么需要反码和补码？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e6%95%b0%e5%ad%a6%e4%b8%93%e6%a0%8f%e8%af%be%e5%a4%96%e5%8a%a0%e9%a4%90%ef%bc%88%e4%b8%80%ef%bc%89%20%e6%88%91%e4%bb%ac%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e5%8f%8d%e7%a0%81%e5%92%8c%e8%a1%a5%e7%a0%81%ef%bc%9f.md">数学专栏课外加餐（一） 我们为什么需要反码和补码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="数学专栏课外加餐（三）：程序员需要读哪些数学书？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e6%95%b0%e5%ad%a6%e4%b8%93%e6%a0%8f%e8%af%be%e5%a4%96%e5%8a%a0%e9%a4%90%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e7%a8%8b%e5%ba%8f%e5%91%98%e9%9c%80%e8%a6%81%e8%af%bb%e5%93%aa%e4%ba%9b%e6%95%b0%e5%ad%a6%e4%b9%a6%ef%bc%9f.md">数学专栏课外加餐（三）：程序员需要读哪些数学书？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="数学专栏课外加餐（二） 位操作的三个应用实例.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e6%95%b0%e5%ad%a6%e4%b8%93%e6%a0%8f%e8%af%be%e5%a4%96%e5%8a%a0%e9%a4%90%ef%bc%88%e4%ba%8c%ef%bc%89%20%e4%bd%8d%e6%93%8d%e4%bd%9c%e7%9a%84%e4%b8%89%e4%b8%aa%e5%ba%94%e7%94%a8%e5%ae%9e%e4%be%8b.md">数学专栏课外加餐（二） 位操作的三个应用实例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 从数学到编程，本身就是一个很长的链条.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e5%9f%ba%e7%a1%80%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e4%bb%8e%e6%95%b0%e5%ad%a6%e5%88%b0%e7%bc%96%e7%a8%8b%ef%bc%8c%e6%9c%ac%e8%ba%ab%e5%b0%b1%e6%98%af%e4%b8%80%e4%b8%aa%e5%be%88%e9%95%bf%e7%9a%84%e9%93%be%e6%9d%a1.md">结束语 从数学到编程，本身就是一个很长的链条.md</a>
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
                            <h1 id="title" data-id="50 推荐系统（下）：如何通过SVD分析用户和物品的矩阵？" class="title">50 推荐系统（下）：如何通过SVD分析用户和物品的矩阵？</h1>
                            <div><p>你好，我是黄申。</p>

<p>上一节，我们讲了如何使用矩阵操作，实现基于用户或者物品的协同过滤。实际上，推荐系统是个很大的课题，你可以尝试不同的想法。比如，对于用户给电影评分的案例，是不是可以使用SVD奇异值的分解，来分解用户评分的矩阵，并找到“潜在”的电影主题呢？如果在一定程度上实现这个目标，那么我们可以通过用户和主题，以及电影和主题之间的关系来进行推荐。今天，我们继续使用MovieLens中的一个数据集，尝试Python代码中的SVD分解，并分析一些结果所代表的含义。</p>

<h2 id="svd回顾以及在推荐中的应用">SVD回顾以及在推荐中的应用</h2>

<p>在实现SVD分解之前，我们先来回顾一下SVD的主要概念和步骤。如果矩阵<span class="math inline">\(X\)</span>是对称的方阵，那么我们可以求得这个矩阵的特征值和特征向量，并把矩阵<span class="math inline">\(X\)</span>分解为特征值和特征向量的乘积。</p>

<p>假设我们求出了矩阵<span class="math inline">\(X\)</span>的<span class="math inline">\(n\)</span>个特征值<span class="math inline">\(λ\_1，λ\_2，…，λ\_n\)</span>，以及这<span class="math inline">\(n\)</span>个特征值所对应的特征向量<span class="math inline">\(v\_1，v\_2，…，v\_n\)</span>，那么矩阵<span class="math inline">\(X\)</span>可以表示为：</p>

<p><span class="math inline">\(X=VΣV^{-1}\)</span></p>

<p>其中，<span class="math inline">\(V\)</span>是这<span class="math inline">\(n\)</span>个特征向量所组成的<span class="math inline">\(n×n\)</span>维矩阵，而<span class="math inline">\(Σ\)</span>是这<span class="math inline">\(n\)</span>个特征值为主对角线的<span class="math inline">\(n×n\)</span>维矩阵。这个过程就是特征分解（Eigendecomposition）。</p>

<p>如果我们会把<span class="math inline">\(V\)</span>的这<span class="math inline">\(n\)</span>个特征向量进行标准化处理，那么对于每个特征向量<span class="math inline">\(V\_i\)</span>，就有<span class="math inline">\(||V\_{i}||\_{2}=1\)</span>，而这表示<span class="math inline">\(V’\_iV\_i=1\)</span>，此时<span class="math inline">\(V\)</span>的<span class="math inline">\(n\)</span>个特征向量为标准正交基，满足<span class="math inline">\(V’V=I\)</span>， 也就是说，<span class="math inline">\(V\)</span>为酉矩阵，有<span class="math inline">\(V’=V^{-1}\)</span> 。这样一来，我们就可以把特征分解表达式写作：</p>

<p><span class="math inline">\(X=VΣV'\)</span></p>

<p>可是，如果矩阵<span class="math inline">\(X\)</span>不是对称的方阵，那么我们不一定能得到有实数解的特征分解。但是，SVD分解可以避免这个问题。</p>

<p>我们可以把<span class="math inline">\(X\)</span>的转置<span class="math inline">\(X’\)</span>和<span class="math inline">\(X\)</span>做矩阵乘法，得到一个<span class="math inline">\(n×n\)</span>维的对称方阵<span class="math inline">\(X’X\)</span>，并对这个对称方阵进行特征分解。分解的时候，我们得到了矩阵<span class="math inline">\(X’X\)</span>的<span class="math inline">\(n\)</span>个特征值和对应的<span class="math inline">\(n\)</span>个特征向量<span class="math inline">\(v\)</span>，其中所有的特征向量叫作<span class="math inline">\(X\)</span>的右奇异向量。通过所有右奇异向量我们可以构造一个<span class="math inline">\(n×n\)</span>维的矩阵<span class="math inline">\(V\)</span>。</p>

<p>类似地，如果我们把<span class="math inline">\(X\)</span>和<span class="math inline">\(X’\)</span>做矩阵乘法，那么会得到一个<span class="math inline">\(m×m\)</span>维的对称方阵<span class="math inline">\(XX’\)</span>。由于<span class="math inline">\(XX’\)</span>也是方阵，因此我们同样可以对它进行特征分解，并得到矩阵<span class="math inline">\(XX’\)</span>的<span class="math inline">\(m\)</span>个特征值和对应的<span class="math inline">\(m\)</span>个特征向量<span class="math inline">\(u\)</span>，其中所有的特征向量向叫作<span class="math inline">\(X\)</span>的左奇异向量。通过所有左奇异向量我们可以构造一个<span class="math inline">\(m×m\)</span>的矩阵<span class="math inline">\(U\)</span>。</p>

<p>现在，包含左右奇异向量的<span class="math inline">\(U\)</span>和<span class="math inline">\(V\)</span>都求解出来了，只剩下奇异值矩阵<span class="math inline">\(Σ\)</span>了。<span class="math inline">\(Σ\)</span>除了对角线上是奇异值之外，其他位置的元素都是0，所以我们只需要求出每个奇异值<span class="math inline">\(σ\)</span>就可以了。之前我们已经推导过，<span class="math inline">\(σ\)</span>可以通过两种方式获得。第一种方式是计算下面这个式子：</p>

<p><span class="math inline">\(σ\_i=\\frac{X\_{v\_{i}}}{u\_{i}}\)</span></p>

<p>其中<span class="math inline">\(v\_i\)</span>和<span class="math inline">\(u\_i\)</span>都是列向量。一旦我们求出了每个奇异值<span class="math inline">\(σ\)</span>，那么就能得到奇异值矩阵<span class="math inline">\(Σ\)</span>。</p>

<p>第二种方式是通过<span class="math inline">\(X’X\)</span>矩阵或者<span class="math inline">\(XX’\)</span>矩阵的特征值之平方根，来求奇异值。计算出每个奇异值<span class="math inline">\(σ\)</span>，那么就能得到奇异值矩阵<span class="math inline">\(Σ\)</span>了。</p>

<p>通过上述几个步骤，我们就能把一个<span class="math inline">\(mxn\)</span>维的实数矩阵，分解成<span class="math inline">\(X=UΣV’\)</span>的形式。那么这种分解对于推荐系统来说，又有怎样的意义呢？</p>

<p>之前我讲过，在潜在语义分析LSA的应用场景下，分解之后所得到的奇异值<span class="math inline">\(σ\)</span>，对应一个语义上的“概念”，而<span class="math inline">\(σ\)</span>值的大小表示这个概念在整个文档集合中的重要程度。<span class="math inline">\(U\)</span>中的左奇异向量表示了每个文档和这些语义“概念”的关系强弱，<span class="math inline">\(V\)</span>中的右奇异向量表示每个词条和这些语义“概念”的关系强弱。</p>

<p>最终，SVD分解把原来的“词条-文档”关系，转换成了“词条-语义概念-文档”的关系。而在推荐系统的应用场景下，对用户评分矩阵的SVD分解，能够帮助我们找到电影中潜在的“主题”，比如科幻类、动作类、浪漫类、传记类等等。</p>

<p>分解之后所得到的奇异值<span class="math inline">\(σ\)</span>对应了一个“主题”，<span class="math inline">\(σ\)</span>值的大小表示这个主题在整个电影集合中的重要程度。<span class="math inline">\(U\)</span>中的左奇异向量表示了每位用户对这些“主题”的喜好程度，<span class="math inline">\(V\)</span>中的右奇异向量表示每部电影和这些“主题”的关系强弱。</p>

<p>最终，SVD分解把原来的“用户-电影”关系，转换成了“用户-主题-电影”的关系。有了这种新的关系，即使我们没有人工标注的电影类型，同样可以使用更多基于电影主题的推荐方法，比如通过用户对电影主题的评分矩阵，进行基于用户或者电影的协同过滤。</p>

<p>接下来，我会使用同样一个MovieLens的数据集，一步步展示如何通过Python语言，对用户评分的矩阵进行SVD分解，并分析一些结果的示例。</p>

<h2 id="python中的svd实现和结果分析">Python中的SVD实现和结果分析</h2>

<p>和上节的代码类似，首先我们需要加载用户对电影的评分。不过，由于非并行SVD分解的时间复杂度是3次方数量级，而空间复杂度是2次方数量级，所以对硬件资源要求很高。这里为了节省测试的时间，我增加了一些语句，只取大约十分之一的数据。</p>

<pre><code>import pandas as pd
from numpy import *


# 加载用户对电影的评分数据
df_ratings = pd.read_csv(&quot;/Users/shenhuang/Data/ml-latest-small/ratings.csv&quot;)


# 获取用户的数量和电影的数量，这里我们只取前1/10来减小数据规模
user_num = int(df_ratings[&quot;userId&quot;].max() / 10)
movie_num = int(df_ratings[&quot;movieId&quot;].max() / 10)


# 构造用户对电影的二元关系矩阵
user_rating = [[0.0] * movie_num for i in range(user_num)]


i = 0
for index, row in df_ratings.iterrows():   # 获取每行的index、row


  # 由于用户和电影的ID都是从1开始，为了和Python的索引一致，减去1
  userId = int(row[&quot;userId&quot;]) - 1
  movieId = int(row[&quot;movieId&quot;]) - 1


  # 我们只取前1/10来减小数据规模
  if (userId &gt;= user_num) or (movieId &gt;= movie_num):
    continue


  # 设置用户对电影的评分
  user_rating[userId][movieId] = row[&quot;rati
</code></pre>

<p>之后，二维数组转为矩阵，以及标准化矩阵的代码和之前是一致的。</p>

<pre><code># 把二维数组转化为矩阵
x = mat(user_rating)


# 标准化每位用户的评分数据
from sklearn.preprocessing import scale


# 对每一行的数据，进行标准化
x_s = scale(x, with_mean=True, with_std=True, axis=1)
print(&quot;标准化后的矩阵：&quot;, x_s
</code></pre>

<p>Python的numpy库，已经实现了一种SVD分解，我们只调用一个函数就行了。</p>

<pre><code># 进行SVD分解
from numpy import linalg as LA


u,sigma,vt = LA.svd(x_s, full_matrices=False, compute_uv=True)
print(&quot;U矩阵：&quot;, u)
print(&quot;Sigma奇异值：&quot;, sigma)
print(&quot;V矩阵：&quot;, vt)
</code></pre>

<p>最后输出的Sigma奇异值大概是这样的：</p>

<pre><code>Sigma奇异值： [416.56942602 285.42546812 202.25724866 ... 79.26188177  76.35167406 74.96719708]
</code></pre>

<p>最后几个奇异值不是0，说明我们没有办法完全忽略它们，不过它们相比最大的几个奇异值还是很小的，我们可以去掉这些值来求得近似的解。</p>

<p>为了验证一下SVD的效果，我们还可以加载电影的元信息，包括电影的标题和类型等等。我在这里使用了一个基于哈希的Python字典结构来存储电影ID到标题和类型的映射。</p>

<pre><code># 加载电影元信息
df_movies = pd.read_csv(&quot;/Users/shenhuang/Data/ml-latest-small/movies.csv&quot;)
dict_movies = {}


for index, row in df_movies.iterrows():   # 获取每行的index、row
  dict_movies[row[&quot;movieId&quot;]] = &quot;{0},{1}&quot;.format(row[&quot;title&quot;], row[&quot;genres&quot;])
print(dict_movies)
</code></pre>

<p>我刚刚提到，分解之后所得到的奇异值<span class="math inline">\(σ\)</span>对应了一个“主题”，<span class="math inline">\(σ\)</span>值的大小表示这个主题在整个电影集合中的重要程度，而V中的右奇异向量表示每部电影和这些“主题”的关系强弱。所以，我们可以对分解后的每个奇异值，通过<span class="math inline">\(V\)</span>中的向量，找找看哪些电影和这个奇异值所对应的主题更相关，然后看看SVD分解所求得的电影主题是不是合理。比如，我们可以使用下面的代码，来查看和向量<span class="math inline">\(Vt1\)</span>,相关的电影主要有哪些。</p>

<pre><code># 输出和某个奇异值高度相关的电影，这些电影代表了一个主题
print(max(vt[1,:]))
for i in range(movie_num):
    if (vt[1][i] &gt; 0.1):
        print(i + 1, vt[1][i], dict_movies[i + 1])
</code></pre>

<p>需要注意的是，向量中的电影ID和原始的电影ID差1，所以在读取dict_movies时需要使用(i + 1)。这个向量中最大的分值大约是0.173，所以我把阈值设置为0.1，并输出了所有分值大于0.1的电影，电影列表如下：</p>

<pre><code>0.17316444479201024
260 0.14287410901699643 Star Wars: Episode IV - A New Hope (1977),Action|Adventure|Sci-Fi
1196 0.1147295905497075 Star Wars: Episode V - The Empire Strikes Back (1980),Action|Adventure|Sci-Fi
1198 0.15453176747222075 Raiders of the Lost Ark (Indiana Jones and the Raiders of the Lost Ark) (1981),Action|Adventure
1210 0.10411193224648774 Star Wars: Episode VI - Return of the Jedi (1983),Action|Adventure|Sci-Fi
2571 0.17316444479201024 Matrix, The (1999),Action|Sci-Fi|Thriller
3578 0.1268370902126096 Gladiator (2000),Action|Adventure|Drama
4993 0.12445203514448012 Lord of the Rings: The Fellowship of the Ring, The (2001),Adventure|Fantasy
5952 0.12535012292041953 Lord of the Rings: The Two Towers, The (2002),Adventure|Fantasy
7153 0.10972312192709989 Lord of the Rings: The Return of the King, The (2003),Action|Adventure|Drama|Fantasy
</code></pre>

<p>从这个列表可以看出，这个主题是关于科幻或者奇幻类的动作冒险题材。</p>

<p>使用类似的代码和同样的阈值0.1，我们来看看和向量<span class="math inline">\(Vt5\)</span>,相关的电影主要有哪些。</p>

<pre><code># 输出和某个奇异值高度相关的电影，这些电影代表了一个主题
print(max(vt[5,:]))
for i in range(movie_num):
    if (vt[5][i] &gt; 0.1):
        print(i + 1, vt[5][i], dict_movies[i + 1])
</code></pre>

<p>电影列表如下：</p>

<pre><code>0.13594520920117012
21 0.13557812349701226 Get Shorty (1995),Comedy|Crime|Thriller
50 0.11870851441884082 Usual Suspects, The (1995),Crime|Mystery|Thriller
62 0.11407971751480048 Mr. Holland's Opus (1995),Drama
168 0.10295400456394468 First Knight (1995),Action|Drama|Romance
222 0.12587492482374366 Circle of Friends (1995),Drama|Romance
261 0.13594520920117012 Little Women (1994),Drama
339 0.10815473505804706 While You Were Sleeping (1995),Comedy|Romance
357 0.11108191756350501 Four Weddings and a Funeral (1994),Comedy|Romance
527 0.1305895737838763 Schindler's List (1993),Drama|War
595 0.11155774544755555 Beauty and the Beast (1991),Animation|Children|Fantasy|Musical|Romance|IMAX
</code></pre>

<p>从这个列表可以看出，这个主题更多的是关于剧情类题材。就目前所看的两个向量来说，SVD在一定程度上区分了不同的电影主题，你也可以使用类似的方式查看更多的向量，以及对应的电影名称和类型。</p>

<h2 id="总结">总结</h2>

<p>在今天的内容中，我们回顾了SVD奇异值分解的核心思想，解释了如何通过<span class="math inline">\(XX’\)</span>和<span class="math inline">\(X’X\)</span>这两个对称矩阵的特征分解，求得分解后的<span class="math inline">\(U\)</span>矩阵、<span class="math inline">\(V\)</span>矩阵和<span class="math inline">\(Σ\)</span>矩阵。另外，我们也解释了在用户对电影评分的应用场景下，SVD分解后的<span class="math inline">\(U\)</span>矩阵、<span class="math inline">\(V\)</span>矩阵和<span class="math inline">\(Σ\)</span>矩阵各自代表的意义，其中<span class="math inline">\(Σ\)</span>矩阵中的奇异值表示了SVD挖掘出来的电影主题，<span class="math inline">\(U\)</span>矩阵中的奇异向量表示用户对这些电影主题的评分，而<span class="math inline">\(V\)</span>矩阵中的奇异向量表示了电影和这些主题的相关程度。</p>

<p>我们还通过Python代码，实践了这种思想在推荐算法中的运用。从结果的奇异值和奇异向量可以看出，SVD分解找到了一些MovieLens数据集上的电影主题。这样我们就可以把用户针对电影的评分转化为用户针对主题的评分。由于主题通常远远小于电影，所以SVD的分解也帮助我们实现了降低特征维度的目的。</p>

<p>SVD分解能够找到一些“潜在的”因素，例如语义上的概念、电影的主题等等。虽然这样操作可以降低特征维度，去掉一些噪音信息，但是由于SVD分解本身的计算量也很大，所以从单次的执行效率来看，SVD往往无法起到优化的作用。在这种情况下，我们可以考虑把它和一些监督式的学习相结合，使用一次分解的结果构建分类器，提升日后的执行效率。</p>

<h2 id="思考题">思考题</h2>

<p>刚才SVD分解实验中得到的<span class="math inline">\(U\)</span>矩阵，是用户对不同电影主题的评分矩阵。请你使用这个<span class="math inline">\(U\)</span>矩阵，进行基于用户或者基于主题（物品）的协同过滤。</p>

<p>欢迎留言和我分享，也欢迎你在留言区写下今天的学习笔记。你可以点击“请朋友读”，把今天的内容分享给你的好友，和他一起精进。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c8a4a4a4f1fcf9f9f8ff88afa5a9a1a4e6aba7a5" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f17b1f30c6dbef8',t:'MTczNDExMTE2My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;&#32;似然估计：如何利用&#32;MLE&#32;对参数进行估计？>
        <link rel="icon" href="/static/favicon.png">
        <title>09  似然估计：如何利用 MLE 对参数进行估计？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  数学，编程能力的营养根基.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e6%95%b0%e5%ad%a6%ef%bc%8c%e7%bc%96%e7%a8%8b%e8%83%bd%e5%8a%9b%e7%9a%84%e8%90%a5%e5%85%bb%e6%a0%b9%e5%9f%ba.md">00 开篇词  数学，编程能力的营养根基.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  从计数开始，程序员必知必会的数制转换法.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/01%20%20%e4%bb%8e%e8%ae%a1%e6%95%b0%e5%bc%80%e5%a7%8b%ef%bc%8c%e7%a8%8b%e5%ba%8f%e5%91%98%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a%e7%9a%84%e6%95%b0%e5%88%b6%e8%bd%ac%e6%8d%a2%e6%b3%95.md">01  从计数开始，程序员必知必会的数制转换法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  逻辑与沟通，怎样才能讲出有逻辑的话？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/02%20%20%e9%80%bb%e8%be%91%e4%b8%8e%e6%b2%9f%e9%80%9a%ef%bc%8c%e6%80%8e%e6%a0%b7%e6%89%8d%e8%83%bd%e8%ae%b2%e5%87%ba%e6%9c%89%e9%80%bb%e8%be%91%e7%9a%84%e8%af%9d%ef%bc%9f.md">02  逻辑与沟通，怎样才能讲出有逻辑的话？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  用数学决策，如何规划好投入、转化和产出？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/03%20%20%e7%94%a8%e6%95%b0%e5%ad%a6%e5%86%b3%e7%ad%96%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%84%e5%88%92%e5%a5%bd%e6%8a%95%e5%85%a5%e3%80%81%e8%bd%ac%e5%8c%96%e5%92%8c%e4%ba%a7%e5%87%ba%ef%bc%9f.md">03  用数学决策，如何规划好投入、转化和产出？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  万物可数学，经典公式是如何在生活中应用的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/04%20%20%e4%b8%87%e7%89%a9%e5%8f%af%e6%95%b0%e5%ad%a6%ef%bc%8c%e7%bb%8f%e5%85%b8%e5%85%ac%e5%bc%8f%e6%98%af%e5%a6%82%e4%bd%95%e5%9c%a8%e7%94%9f%e6%b4%bb%e4%b8%ad%e5%ba%94%e7%94%a8%e7%9a%84%ef%bc%9f.md">04  万物可数学，经典公式是如何在生活中应用的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  求极值：如何找到复杂业务的最优解？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/05%20%20%e6%b1%82%e6%9e%81%e5%80%bc%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%be%e5%88%b0%e5%a4%8d%e6%9d%82%e4%b8%9a%e5%8a%a1%e7%9a%84%e6%9c%80%e4%bc%98%e8%a7%a3%ef%bc%9f.md">05  求极值：如何找到复杂业务的最优解？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  向量及其导数：计算机如何完成对海量高维度数据计算？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/06%20%20%e5%90%91%e9%87%8f%e5%8f%8a%e5%85%b6%e5%af%bc%e6%95%b0%ef%bc%9a%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%a6%82%e4%bd%95%e5%ae%8c%e6%88%90%e5%af%b9%e6%b5%b7%e9%87%8f%e9%ab%98%e7%bb%b4%e5%ba%a6%e6%95%b0%e6%8d%ae%e8%ae%a1%e7%ae%97%ef%bc%9f.md">06  向量及其导数：计算机如何完成对海量高维度数据计算？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  线性回归：如何在离散点中寻找数据规律？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/07%20%20%e7%ba%bf%e6%80%a7%e5%9b%9e%e5%bd%92%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e7%a6%bb%e6%95%a3%e7%82%b9%e4%b8%ad%e5%af%bb%e6%89%be%e6%95%b0%e6%8d%ae%e8%a7%84%e5%be%8b%ef%bc%9f.md">07  线性回归：如何在离散点中寻找数据规律？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  加乘法则：如何计算复杂事件发生的概率？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/08%20%20%e5%8a%a0%e4%b9%98%e6%b3%95%e5%88%99%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a1%e7%ae%97%e5%a4%8d%e6%9d%82%e4%ba%8b%e4%bb%b6%e5%8f%91%e7%94%9f%e7%9a%84%e6%a6%82%e7%8e%87%ef%bc%9f.md">08  加乘法则：如何计算复杂事件发生的概率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  似然估计：如何利用 MLE 对参数进行估计？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/09%20%20%e4%bc%bc%e7%84%b6%e4%bc%b0%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20MLE%20%e5%af%b9%e5%8f%82%e6%95%b0%e8%bf%9b%e8%a1%8c%e4%bc%b0%e8%ae%a1%ef%bc%9f.md">09  似然估计：如何利用 MLE 对参数进行估计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  信息熵：事件的不确定性如何计算？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/10%20%20%e4%bf%a1%e6%81%af%e7%86%b5%ef%bc%9a%e4%ba%8b%e4%bb%b6%e7%9a%84%e4%b8%8d%e7%a1%ae%e5%ae%9a%e6%80%a7%e5%a6%82%e4%bd%95%e8%ae%a1%e7%ae%97%ef%bc%9f.md">10  信息熵：事件的不确定性如何计算？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  灰度实验：如何设计灰度实验并计算实验的收益？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/11%20%20%e7%81%b0%e5%ba%a6%e5%ae%9e%e9%aa%8c%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e7%81%b0%e5%ba%a6%e5%ae%9e%e9%aa%8c%e5%b9%b6%e8%ae%a1%e7%ae%97%e5%ae%9e%e9%aa%8c%e7%9a%84%e6%94%b6%e7%9b%8a%ef%bc%9f.md">11  灰度实验：如何设计灰度实验并计算实验的收益？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  统计学方法：如何证明灰度实验效果不是偶然得到的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/12%20%20%e7%bb%9f%e8%ae%a1%e5%ad%a6%e6%96%b9%e6%b3%95%ef%bc%9a%e5%a6%82%e4%bd%95%e8%af%81%e6%98%8e%e7%81%b0%e5%ba%a6%e5%ae%9e%e9%aa%8c%e6%95%88%e6%9e%9c%e4%b8%8d%e6%98%af%e5%81%b6%e7%84%b6%e5%be%97%e5%88%b0%e7%9a%84%ef%bc%9f.md">12  统计学方法：如何证明灰度实验效果不是偶然得到的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  复杂度：如何利用数学推导对程序进行优化？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/13%20%20%e5%a4%8d%e6%9d%82%e5%ba%a6%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e6%95%b0%e5%ad%a6%e6%8e%a8%e5%af%bc%e5%af%b9%e7%a8%8b%e5%ba%8f%e8%bf%9b%e8%a1%8c%e4%bc%98%e5%8c%96%ef%bc%9f.md">13  复杂度：如何利用数学推导对程序进行优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  程序的循环：如何利用数学归纳法进行程序开发？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/14%20%20%e7%a8%8b%e5%ba%8f%e7%9a%84%e5%be%aa%e7%8e%af%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e6%95%b0%e5%ad%a6%e5%bd%92%e7%ba%b3%e6%b3%95%e8%bf%9b%e8%a1%8c%e7%a8%8b%e5%ba%8f%e5%bc%80%e5%8f%91%ef%bc%9f.md">14  程序的循环：如何利用数学归纳法进行程序开发？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  递归：如何计算汉诺塔问题的移动步数？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/15%20%20%e9%80%92%e5%bd%92%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a1%e7%ae%97%e6%b1%89%e8%af%ba%e5%a1%94%e9%97%ae%e9%a2%98%e7%9a%84%e7%a7%bb%e5%8a%a8%e6%ad%a5%e6%95%b0%ef%bc%9f.md">15  递归：如何计算汉诺塔问题的移动步数？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  二分法：如何利用指数爆炸优化程序？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/16%20%20%e4%ba%8c%e5%88%86%e6%b3%95%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e6%8c%87%e6%95%b0%e7%88%86%e7%82%b8%e4%bc%98%e5%8c%96%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">16  二分法：如何利用指数爆炸优化程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  动态规划：如何利用最优子结构解决问题？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/17%20%20%e5%8a%a8%e6%80%81%e8%a7%84%e5%88%92%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e6%9c%80%e4%bc%98%e5%ad%90%e7%bb%93%e6%9e%84%e8%a7%a3%e5%86%b3%e9%97%ae%e9%a2%98%ef%bc%9f.md">17  动态规划：如何利用最优子结构解决问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  AI 入门：利用 3 个公式搭建最简 AI 框架.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/18%20%20AI%20%e5%85%a5%e9%97%a8%ef%bc%9a%e5%88%a9%e7%94%a8%203%20%e4%b8%aa%e5%85%ac%e5%bc%8f%e6%90%ad%e5%bb%ba%e6%9c%80%e7%ae%80%20AI%20%e6%a1%86%e6%9e%b6.md">18  AI 入门：利用 3 个公式搭建最简 AI 框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  逻辑回归：如何让计算机做出二值化决策？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/19%20%20%e9%80%bb%e8%be%91%e5%9b%9e%e5%bd%92%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%81%9a%e5%87%ba%e4%ba%8c%e5%80%bc%e5%8c%96%e5%86%b3%e7%ad%96%ef%bc%9f.md">19  逻辑回归：如何让计算机做出二值化决策？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  决策树：如何对 NP 难复杂问题进行启发式求解？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/20%20%20%e5%86%b3%e7%ad%96%e6%a0%91%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%b9%20NP%20%e9%9a%be%e5%a4%8d%e6%9d%82%e9%97%ae%e9%a2%98%e8%bf%9b%e8%a1%8c%e5%90%af%e5%8f%91%e5%bc%8f%e6%b1%82%e8%a7%a3%ef%bc%9f.md">20  决策树：如何对 NP 难复杂问题进行启发式求解？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  神经网络与深度学习：计算机是如何理解图像、文本和语音的？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/21%20%20%e7%a5%9e%e7%bb%8f%e7%bd%91%e7%bb%9c%e4%b8%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%ef%bc%9a%e8%ae%a1%e7%ae%97%e6%9c%ba%e6%98%af%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e5%9b%be%e5%83%8f%e3%80%81%e6%96%87%e6%9c%ac%e5%92%8c%e8%af%ad%e9%9f%b3%e7%9a%84%ef%bc%9f.md">21  神经网络与深度学习：计算机是如何理解图像、文本和语音的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  面试中那些坑了无数人的算法题.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/22%20%20%e9%9d%a2%e8%af%95%e4%b8%ad%e9%82%a3%e4%ba%9b%e5%9d%91%e4%ba%86%e6%97%a0%e6%95%b0%e4%ba%ba%e7%9a%84%e7%ae%97%e6%b3%95%e9%a2%98.md">22  面试中那些坑了无数人的算法题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  站在生活的十字路口，如何用数学抉择？.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/23%20%20%e7%ab%99%e5%9c%a8%e7%94%9f%e6%b4%bb%e7%9a%84%e5%8d%81%e5%ad%97%e8%b7%af%e5%8f%a3%ef%bc%8c%e5%a6%82%e4%bd%95%e7%94%a8%e6%95%b0%e5%ad%a6%e6%8a%89%e6%8b%a9%ef%bc%9f.md">23  站在生活的十字路口，如何用数学抉择？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 结束语  数学底子好，学啥都快.md" href="/%e4%b8%93%e6%a0%8f/%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e6%95%b0%e5%ad%a6%e8%af%be/24%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e6%95%b0%e5%ad%a6%e5%ba%95%e5%ad%90%e5%a5%bd%ef%bc%8c%e5%ad%a6%e5%95%a5%e9%83%bd%e5%bf%ab.md">24 结束语  数学底子好，学啥都快.md</a>
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
                            <h1 id="title" data-id="09  似然估计：如何利用 MLE 对参数进行估计？" class="title">09  似然估计：如何利用 MLE 对参数进行估计？</h1>
                            <div><p>你好，欢迎来到第 09 课时——似然估计：如何利用 MLE 对参数进行估计？</p>

<p>前面我们学会了如何计算概率，这一讲我们学习如何利用概率对某个参数进行估计。在读书的时候，你一定接触过极大似然估计，它是数学课程的难点之一，它名字背后的含义，以及它的推导过程都非常复杂，需要你对它有深刻的理解。</p>

<p>不过，有了前面“形式化定义”“概率计算的加乘法则”和求函数最值的“求导法”“梯度下降法”的知识储备，相信极大似然估计也能迎刃而解。</p>

<h3 id="白话理解-极大似然估计">白话理解“极大似然估计”</h3>

<p>如果你是刚刚学习概率，极大似然估计这六个字一定会让你产生不解。</p>

<p><strong>似然</strong>（Likelihood），可以理解为可能性，也就是概率。举个例子，某个同学毕业于华中科技大学这样的工科院校，那么这位同学是男生的可能性（或者说概率、似然）就更大；相反，某个同学毕业于北京外国语学院这样的文科院校，那么这位同学是女生的可能性（或者说概率、似然）就更大。</p>

<p>那么反过来思考，如果大漂亮是个美丽又可爱的女生，现在有两个候选项：A.大漂亮毕业于华中科技大学；B.大漂亮毕业于北京外国语学院。在对其他信息都毫不知情的情况下，你更愿意相信哪个呢？很显然，相信 B 是更好的选项，因为 B 的概率（或者说似然）更大。</p>

<p>其实，在刚刚的思考逻辑中，我们已经不知不觉地用了极大似然估计的思想了——<strong>估计</strong>（Estimate），用大白话说就是“猜”。</p>

<p>例如，你对于大漂亮毕业院校的“估计”是她来自北京外国语学院；这就是说，你“猜测”大漂亮毕业于北京外国语学院。那么，为何你猜测她毕业于北京外国语学院，而不是华中科技大学呢？原因就是前者的可能性更大，而后者可能性更小。换句话说，从可能性的视角看，前者是个<strong>极大</strong>值（Maximum）。</p>

<p>我们将上面思考过程的 3 个关键词“<strong>极大</strong>（Maximum）”“<strong>似然</strong>（Likelihood）”“<strong>估计</strong>（Estimate）”给提炼出来，就得到了极大似然估计这个方法，通常也可以用这 3 个单词的首个字母来表示——MLE。</p>

<h3 id="极大似然估计的方法路径">极大似然估计的方法路径</h3>

<p>从刚才的例子不难看出，极大似然估计做的事情，就是<strong>通过已知条件对某个未知参数进行估计，它根据观测的样本构建似然函数，再通过让这个函数取得极大值，来完成估计</strong>。接着，我们用数学语言来描述整个过程。</p>

<p>极大似然估计的流程可以分为 3 步，分别是似然、极大和估计。</p>

<ul>
<li>第一步<strong>似然</strong>，即根据观测的样本建立似然函数，也是概率函数或可能性函数。
这个步骤的数学表达如下：假设观测的样本或集合为 D，待估计的参数为 θ。则观察到样本集合的概率，就是在参数 θ 条件下，D 发生的条件概率 P(D|θ)。这就是似然函数，也是极大似然估计中最难的一步。</li>
<li>第二步<strong>极大</strong>，也就是求解似然函数的极大值。
你可以通过求导法、梯度下降法等方式求解。这个步骤的数学表达就简单许多，即 max P(D|θ)。</li>
<li>第三步<strong>估计</strong>，利用求解出的极大值，对未知参数进行估计。
<img src="assets/CgqCHl-3jnCAQHn-AADNrmqedmI922.png" alt="图片1.png" /></li>
</ul>

<p>利用这 3 步就完成了极大似然估计的整个流程。</p>

<p><strong>接下来，我们将这个方法路径用在对“大漂亮毕业院校的极大似然估计表达”上。</strong></p>

<ul>
<li>第一步 <strong>似然</strong></li>
</ul>

<p>我们观测的样本结果 D 是“大漂亮是个女生”，待估计的变量 θ 是“大漂亮毕业于哪个学校”。这样，似然函数就是 P(D|θ) = P(大漂亮是个女生|大漂亮毕业于 θ 学校)，其中 θ∈(北京外国语学院,华中科技大学)。</p>

<p>接着，我们还需要了解工科院校、文科院校的男女比例情况，把似然函数写出具体的数字表达。假设华中科技大学的男女比例为 7:1，北京外国语学院的男女比例为 1:8，则有下表的概率值：</p>

<p><img src="assets/Ciqc1F-3jn2AWcQRAADZzUOWHhg550.png" alt="图片2.png" /></p>

<ul>
<li>第二步 <strong>极大</strong></li>
</ul>

<p>有了前面的信息，我们就能求解似然函数的极大值了。似然函数中参数 θ 是离散值，只有两个可能的取值。因此，我们既不需要求导法，也不需要梯度下降法，只需要把两种可能性都算一下，再进行比较就可以了。</p>

<p>不难发现，因为 P(女|北外)=<sup>8</sup>&frasl;<sub>9</sub> &gt; P(女|华科) = 1/8，所以似然函数的极大值是 8/9。</p>

<ul>
<li>第三步 <strong>估计</strong></li>
</ul>

<p>求解出似然函数的极大值之后，我们利用取得极大值的参数值作为结果，则有
<img src="assets/CgqCHl-3jomAa0uOAAA8fD9AnMk614.png" alt="图片3.png" /></p>

<h3 id="极大似然估计的拓展">极大似然估计的拓展</h3>

<p>前面的例子很简单，而实际中你可能还会遇到很复杂的拓展问题。</p>

<h4 id="1-第一个复杂的拓展问题-为单样本拓展为多样本">1.第一个复杂的拓展问题，为单样本拓展为多样本</h4>

<p>刚刚的观察样本集合中，只有一个样本（即大漂亮是个女生）。而如果有多个样本又该怎么办呢？</p>

<p>此时我们需要用到概率计算的乘法法则。通常，我们都会认为同一个事件的不同观测结果是独立的，因此可以用乘法法则计算它们共同发生的概率。</p>

<p>这个过程用数学语言表达，就是假设观测的样本集合为 D = (d1,d2,d3……dn)，待估计的参数为θ，则似然函数 P(D|θ) = P(d1,d2,d3……dn|θ)。</p>

<p>因为观测样本独立，满足 P(AB) = P(A)·P(B)，则有</p>

<p><img src="assets/CgqCHl-3jqGAF23JAABMly9eS8A616.png" alt="图片4.png" /></p>

<h4 id="2-第二个拓展问题-是似然函数到对数似然函数">2.第二个拓展问题，是似然函数到对数似然函数</h4>

<p>刚刚的推导结果非常吓人。大型连乘算式中，直接求解最值是非常困难的。不过，庆幸的是数学中有个化乘法为加法的函数——对数函数。因为对数函数是单调的，所以在化乘法为加法的过程中，不会改变最大值发生的位置，即 ln(xy) = ln x + ln y。</p>

<p><img src="assets/CgqCHl-3jsCAfDcaAABZD-jbRFI179.png" alt="图片5.png" /></p>

<h3 id="mle-梳理">MLE 梳理</h3>

<p>到这里，关于 MLE 所有的知识点就讲完了，我们做个简单的梳理。</p>

<p>极大似然估计的目标，是通过观察样本估计某个参数的值，它估计的方法路径如下。</p>

<ul>
<li>第一步，通过观察到的样本，建立代表这些样本发生可能性的似然函数。</li>
<li>第二步，利用求导法、梯度下降法等算法，求解似然函数的极大值。</li>
<li>第三步，用似然函数取得极大值的参数值，作为结果的估计值并输出。
在实际应用，样本很多的时候，通常认为样本之间是独立的，满足概率相乘的乘法法则；而面对连乘的复杂运算，通常采用对数似然函数的处理方式，化连乘为求和运算。</li>
</ul>

<p>以上就是 MLE 基础原理的知识。</p>

<h3 id="极大似然估计在工作场景中的应用">极大似然估计在工作场景中的应用</h3>

<p>我们看一个利用极大似然估计解决实际工作问题的案例。</p>

<p>假设大迷糊是某个电商公司负责质量检测的工程师，这个公司的商品质量可以分为三档，分别是优质品、合格品和残次品。BI 的同事根据调研，发现商品的质量满足如下概率分布：</p>

<p><img src="assets/CgqCHl-3jtSAJhzDAACaJLpVCsc726.png" alt="图片6.png" /></p>

<p>其中 θ 是个未知参数，大迷糊想用 MLE 的方法估计出 θ 的值。于是，大迷糊对商品进行了采样，得到的采样值分别为优质品、优质品和合格品。现在，让我们用 MLE 帮助大迷糊来估计未知数 θ 的值吧。</p>

<ul>
<li>第一步 似然</li>
</ul>

<p>我们发现，样本集合有 3 个样本，则 D = (d1,d2,d3) = (优质品,优质品,合格品)。待估计的未知数为θ，则似然函数为 P(D|θ) = P(d1,d2,d3|θ) = P(d1|θ)·P(d2|θ)·P(d3|θ)。</p>

<p>代入 d1～d3 的值，以及对应的概率，则有 P(D|θ) = P(优质品|θ)·P(优质品|θ)·P(合格品|θ) = θ4 * 2θ(1-θ)。</p>

<p>那么，对数似然就是 ln P(D|θ) = ln (θ4 * 2θ(1-θ)) = ln 2 + 5 ln θ + ln (1-θ)。</p>

<ul>
<li>第二步 极大</li>
</ul>

<p>有了似然函数，我们就来尝试求解它的极大值吧。首先求对数似然函数关于 θ 的导数，则有
<img src="assets/CgqCHl-3juyAanYJAAAyG8k3Pfs920.png" alt="图片7.png" />
推导到这里，你会发现直接用求导法建立导函数为零的方程就能得到结果。这是因为，商品质量函数都是比较简单的多项式。如果里面包含了复杂的函数，例如指数函数、正弦函数等，就必须要借助梯度下降法来求解了。</p>

<p>为了再次说明梯度下降法的使用，我们这里尝试采用梯度下降法来求解，我们直接给出代码：</p>

<pre><code>import math

def grad(x):

	return (5 - 6 * x) / (x*(1-x))

def main():

	a = 0.01

	maxloop = 1000

	theta = 0.1

	for _ in range(maxloop):

		g = grad(theta)

		theta = theta + a*g

	print theta

if __name__ == '__main__':

	main()
</code></pre>

<p>我们对代码进行走读。</p>

<ul>
<li>主函数中，设置学习率为 0.01，最大迭代轮数为 1000 次，θ 的初始值设置为 0.1。</li>
<li>接下来，第 10～12 行，是 1000 次的循环体。每次循环执行两个动作，分别是计算梯度，并把结果保存在 g 变量中；再用学习率和梯度的乘积，去更新 θ。</li>
<li>在计算梯度的函数 grad() 内部，直接返回一阶导数值。这是因为对于单变量而言，一阶导数的值就是其梯度的值。</li>
</ul>

<p>我们执行这段代码，打印的结果如下图所示：
<img src="assets/CgqCHl-3jyWAAphzAABBhDmkpZk268.png" alt="image" /></p>

<p>如果我们用求导法，则有(5-6θ)/(θ*(1-θ)) = 0，解得 θ = <sup>5</sup>&frasl;<sub>6</sub> = 0.8333，这与我们用梯度下降法求得的结果一致。</p>

<ul>
<li>第三步 估计</li>
</ul>

<p>我们求解出的 θ* 值为 0.8333。它的含义是当 θ = θ* 时，大迷糊随机抽取 3 个样本恰好是优质品、优质品、合格品的概率最大。因此，我们有理由相信，θ* 是最有可能让这个观测结果出现的参数值。因此，0.8333 就是这里 θ 的估计结果。</p>

<h3 id="小结">小结</h3>

<p>MLE 覆盖的知识点比较多。要想利用 MLE 去解决问题，你首先需要会计算概率，构建似然函数；接着，你还需要一些算法知识的储备，才能让你面对任何一个复杂函数，都能快速求解其最大值；最后，你还需要一个小技巧，那就是似然函数转化为对数似然函数后，最优估计值是不变的。</p>

<p>正是 MLE 的背后需要很多知识和能力，才让它成为数学学习过程中的一个难点。不过，庆幸的是，它的编程实现还是非常简单的。如果你掌握了梯度下降法的开发，那么 MLE 的开发也一定难不倒你。</p>

<p>最后，我们给一个练习题。假设在本例中，商品质量的分布如下：</p>

<p><img src="assets/Ciqc1F-3jviAJErpAABn_kNDrps965.png" alt="图片8.png" /></p>

<p>试着再来帮大迷糊来估计下 θ 的值吧。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#09656565303d3838393e496e64686065276a6664" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f17b645586abef8',t:'MTczNDExMTM0MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
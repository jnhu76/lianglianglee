<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=24&#32;Spark&#32;3.0（一）：AQE的3个特性怎么才能用好？>
        <link rel="icon" href="/static/favicon.png">
        <title>24 Spark 3.0（一）：AQE的3个特性怎么才能用好？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 Spark性能调优，你该掌握这些“套路”.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%ef%bc%8c%e4%bd%a0%e8%af%a5%e6%8e%8c%e6%8f%a1%e8%bf%99%e4%ba%9b%e2%80%9c%e5%a5%97%e8%b7%af%e2%80%9d.md">00 开篇词 Spark性能调优，你该掌握这些“套路”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 性能调优的必要性：Spark本身就很快，为啥还需要我调优？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/01%20%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e7%9a%84%e5%bf%85%e8%a6%81%e6%80%a7%ef%bc%9aSpark%e6%9c%ac%e8%ba%ab%e5%b0%b1%e5%be%88%e5%bf%ab%ef%bc%8c%e4%b8%ba%e5%95%a5%e8%bf%98%e9%9c%80%e8%a6%81%e6%88%91%e8%b0%83%e4%bc%98%ef%bc%9f.md">01 性能调优的必要性：Spark本身就很快，为啥还需要我调优？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 性能调优的本质：调优的手段五花八门，该从哪里入手？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/02%20%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e7%9a%84%e6%9c%ac%e8%b4%a8%ef%bc%9a%e8%b0%83%e4%bc%98%e7%9a%84%e6%89%8b%e6%ae%b5%e4%ba%94%e8%8a%b1%e5%85%ab%e9%97%a8%ef%bc%8c%e8%af%a5%e4%bb%8e%e5%93%aa%e9%87%8c%e5%85%a5%e6%89%8b%ef%bc%9f.md">02 性能调优的本质：调优的手段五花八门，该从哪里入手？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 RDD：为什么你必须要理解弹性分布式数据集？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/03%20RDD%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bd%a0%e5%bf%85%e9%a1%bb%e8%a6%81%e7%90%86%e8%a7%a3%e5%bc%b9%e6%80%a7%e5%88%86%e5%b8%83%e5%bc%8f%e6%95%b0%e6%8d%ae%e9%9b%86%ef%bc%9f.md">03 RDD：为什么你必须要理解弹性分布式数据集？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 DAG与流水线：到底啥叫“内存计算”？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/04%20DAG%e4%b8%8e%e6%b5%81%e6%b0%b4%e7%ba%bf%ef%bc%9a%e5%88%b0%e5%ba%95%e5%95%a5%e5%8f%ab%e2%80%9c%e5%86%85%e5%ad%98%e8%ae%a1%e7%ae%97%e2%80%9d%ef%bc%9f.md">04 DAG与流水线：到底啥叫“内存计算”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 调度系统：“数据不动代码动”到底是什么意思？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/05%20%e8%b0%83%e5%ba%a6%e7%b3%bb%e7%bb%9f%ef%bc%9a%e2%80%9c%e6%95%b0%e6%8d%ae%e4%b8%8d%e5%8a%a8%e4%bb%a3%e7%a0%81%e5%8a%a8%e2%80%9d%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%e6%84%8f%e6%80%9d%ef%bc%9f.md">05 调度系统：“数据不动代码动”到底是什么意思？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 存储系统：空间换时间，还是时间换空间？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/06%20%e5%ad%98%e5%82%a8%e7%b3%bb%e7%bb%9f%ef%bc%9a%e7%a9%ba%e9%97%b4%e6%8d%a2%e6%97%b6%e9%97%b4%ef%bc%8c%e8%bf%98%e6%98%af%e6%97%b6%e9%97%b4%e6%8d%a2%e7%a9%ba%e9%97%b4%ef%bc%9f.md">06 存储系统：空间换时间，还是时间换空间？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 内存管理基础：Spark如何高效利用有限的内存空间？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/07%20%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%e5%9f%ba%e7%a1%80%ef%bc%9aSpark%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%88%a9%e7%94%a8%e6%9c%89%e9%99%90%e7%9a%84%e5%86%85%e5%ad%98%e7%a9%ba%e9%97%b4%ef%bc%9f.md">07 内存管理基础：Spark如何高效利用有限的内存空间？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 应用开发三原则：如何拓展自己的开发边界？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/08%20%e5%ba%94%e7%94%a8%e5%bc%80%e5%8f%91%e4%b8%89%e5%8e%9f%e5%88%99%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8b%93%e5%b1%95%e8%87%aa%e5%b7%b1%e7%9a%84%e5%bc%80%e5%8f%91%e8%be%b9%e7%95%8c%ef%bc%9f.md">08 应用开发三原则：如何拓展自己的开发边界？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 调优一筹莫展，配置项速查手册让你事半功倍！（上）.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/09%20%e8%b0%83%e4%bc%98%e4%b8%80%e7%ad%b9%e8%8e%ab%e5%b1%95%ef%bc%8c%e9%85%8d%e7%bd%ae%e9%a1%b9%e9%80%9f%e6%9f%a5%e6%89%8b%e5%86%8c%e8%ae%a9%e4%bd%a0%e4%ba%8b%e5%8d%8a%e5%8a%9f%e5%80%8d%ef%bc%81%ef%bc%88%e4%b8%8a%ef%bc%89.md">09 调优一筹莫展，配置项速查手册让你事半功倍！（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 调优一筹莫展，配置项速查手册让你事半功倍！（下）.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/10%20%e8%b0%83%e4%bc%98%e4%b8%80%e7%ad%b9%e8%8e%ab%e5%b1%95%ef%bc%8c%e9%85%8d%e7%bd%ae%e9%a1%b9%e9%80%9f%e6%9f%a5%e6%89%8b%e5%86%8c%e8%ae%a9%e4%bd%a0%e4%ba%8b%e5%8d%8a%e5%8a%9f%e5%80%8d%ef%bc%81%ef%bc%88%e4%b8%8b%ef%bc%89.md">10 调优一筹莫展，配置项速查手册让你事半功倍！（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 为什么说Shuffle是一时无两的性能杀手？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/11%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4Shuffle%e6%98%af%e4%b8%80%e6%97%b6%e6%97%a0%e4%b8%a4%e7%9a%84%e6%80%a7%e8%83%bd%e6%9d%80%e6%89%8b%ef%bc%9f.md">11 为什么说Shuffle是一时无两的性能杀手？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 广播变量（一）：克制Shuffle，如何一招制胜！.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/12%20%e5%b9%bf%e6%92%ad%e5%8f%98%e9%87%8f%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%85%8b%e5%88%b6Shuffle%ef%bc%8c%e5%a6%82%e4%bd%95%e4%b8%80%e6%8b%9b%e5%88%b6%e8%83%9c%ef%bc%81.md">12 广播变量（一）：克制Shuffle，如何一招制胜！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 广播变量（二）：如何让Spark SQL选择Broadcast Joins？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/13%20%e5%b9%bf%e6%92%ad%e5%8f%98%e9%87%8f%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9Spark%20SQL%e9%80%89%e6%8b%a9Broadcast%20Joins%ef%bc%9f.md">13 广播变量（二）：如何让Spark SQL选择Broadcast Joins？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 CPU视角：如何高效地利用CPU？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/14%20CPU%e8%a7%86%e8%a7%92%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%9c%b0%e5%88%a9%e7%94%a8CPU%ef%bc%9f.md">14 CPU视角：如何高效地利用CPU？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 内存视角（一）：如何最大化内存的使用效率？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/15%20%e5%86%85%e5%ad%98%e8%a7%86%e8%a7%92%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9c%80%e5%a4%a7%e5%8c%96%e5%86%85%e5%ad%98%e7%9a%84%e4%bd%bf%e7%94%a8%e6%95%88%e7%8e%87%ef%bc%9f.md">15 内存视角（一）：如何最大化内存的使用效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 内存视角（二）：如何有效避免Cache滥用？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/16%20%e5%86%85%e5%ad%98%e8%a7%86%e8%a7%92%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9c%89%e6%95%88%e9%81%bf%e5%85%8dCache%e6%bb%a5%e7%94%a8%ef%bc%9f.md">16 内存视角（二）：如何有效避免Cache滥用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 内存视角（三）：OOM都是谁的锅？怎么破？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/17%20%e5%86%85%e5%ad%98%e8%a7%86%e8%a7%92%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9aOOM%e9%83%bd%e6%98%af%e8%b0%81%e7%9a%84%e9%94%85%ef%bc%9f%e6%80%8e%e4%b9%88%e7%a0%b4%ef%bc%9f.md">17 内存视角（三）：OOM都是谁的锅？怎么破？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 磁盘视角：如果内存无限大，磁盘还有用武之地吗？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/18%20%e7%a3%81%e7%9b%98%e8%a7%86%e8%a7%92%ef%bc%9a%e5%a6%82%e6%9e%9c%e5%86%85%e5%ad%98%e6%97%a0%e9%99%90%e5%a4%a7%ef%bc%8c%e7%a3%81%e7%9b%98%e8%bf%98%e6%9c%89%e7%94%a8%e6%ad%a6%e4%b9%8b%e5%9c%b0%e5%90%97%ef%bc%9f.md">18 磁盘视角：如果内存无限大，磁盘还有用武之地吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 网络视角：如何有效降低网络开销？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/19%20%e7%bd%91%e7%bb%9c%e8%a7%86%e8%a7%92%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9c%89%e6%95%88%e9%99%8d%e4%bd%8e%e7%bd%91%e7%bb%9c%e5%bc%80%e9%94%80%ef%bc%9f.md">19 网络视角：如何有效降低网络开销？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 RDD和DataFrame：既生瑜，何生亮？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/20%20RDD%e5%92%8cDataFrame%ef%bc%9a%e6%97%a2%e7%94%9f%e7%91%9c%ef%bc%8c%e4%bd%95%e7%94%9f%e4%ba%ae%ef%bc%9f.md">20 RDD和DataFrame：既生瑜，何生亮？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Catalyst逻辑计划：你的SQL语句是怎么被优化的？（上）.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/21%20Catalyst%e9%80%bb%e8%be%91%e8%ae%a1%e5%88%92%ef%bc%9a%e4%bd%a0%e7%9a%84SQL%e8%af%ad%e5%8f%a5%e6%98%af%e6%80%8e%e4%b9%88%e8%a2%ab%e4%bc%98%e5%8c%96%e7%9a%84%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">21 Catalyst逻辑计划：你的SQL语句是怎么被优化的？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Catalyst物理计划：你的SQL语句是怎么被优化的（下）？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/22%20Catalyst%e7%89%a9%e7%90%86%e8%ae%a1%e5%88%92%ef%bc%9a%e4%bd%a0%e7%9a%84SQL%e8%af%ad%e5%8f%a5%e6%98%af%e6%80%8e%e4%b9%88%e8%a2%ab%e4%bc%98%e5%8c%96%e7%9a%84%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9f.md">22 Catalyst物理计划：你的SQL语句是怎么被优化的（下）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 钨丝计划：Tungsten给开发者带来了哪些福报？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/23%20%e9%92%a8%e4%b8%9d%e8%ae%a1%e5%88%92%ef%bc%9aTungsten%e7%bb%99%e5%bc%80%e5%8f%91%e8%80%85%e5%b8%a6%e6%9d%a5%e4%ba%86%e5%93%aa%e4%ba%9b%e7%a6%8f%e6%8a%a5%ef%bc%9f.md">23 钨丝计划：Tungsten给开发者带来了哪些福报？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 Spark 3.0（一）：AQE的3个特性怎么才能用好？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/24%20Spark%203.0%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aAQE%e7%9a%843%e4%b8%aa%e7%89%b9%e6%80%a7%e6%80%8e%e4%b9%88%e6%89%8d%e8%83%bd%e7%94%a8%e5%a5%bd%ef%bc%9f.md">24 Spark 3.0（一）：AQE的3个特性怎么才能用好？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 Spark 3.0（二）：DPP特性该怎么用？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/25%20Spark%203.0%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aDPP%e7%89%b9%e6%80%a7%e8%af%a5%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f.md">25 Spark 3.0（二）：DPP特性该怎么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 Join Hints指南：不同场景下，如何选择Join策略？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/26%20Join%20Hints%e6%8c%87%e5%8d%97%ef%bc%9a%e4%b8%8d%e5%90%8c%e5%9c%ba%e6%99%af%e4%b8%8b%ef%bc%8c%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9Join%e7%ad%96%e7%95%a5%ef%bc%9f.md">26 Join Hints指南：不同场景下，如何选择Join策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 大表Join小表：广播变量容不下小表怎么办？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/27%20%e5%a4%a7%e8%a1%a8Join%e5%b0%8f%e8%a1%a8%ef%bc%9a%e5%b9%bf%e6%92%ad%e5%8f%98%e9%87%8f%e5%ae%b9%e4%b8%8d%e4%b8%8b%e5%b0%8f%e8%a1%a8%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9f.md">27 大表Join小表：广播变量容不下小表怎么办？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 大表Join大表（一）：什么是“分而治之”的调优思路？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/28%20%e5%a4%a7%e8%a1%a8Join%e5%a4%a7%e8%a1%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%af%e2%80%9c%e5%88%86%e8%80%8c%e6%b2%bb%e4%b9%8b%e2%80%9d%e7%9a%84%e8%b0%83%e4%bc%98%e6%80%9d%e8%b7%af%ef%bc%9f.md">28 大表Join大表（一）：什么是“分而治之”的调优思路？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 大表Join大表（二）：什么是负隅顽抗的调优思路？.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/29%20%e5%a4%a7%e8%a1%a8Join%e5%a4%a7%e8%a1%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%af%e8%b4%9f%e9%9a%85%e9%a1%bd%e6%8a%97%e7%9a%84%e8%b0%83%e4%bc%98%e6%80%9d%e8%b7%af%ef%bc%9f.md">29 大表Join大表（二）：什么是负隅顽抗的调优思路？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 应用开发：北京市小客车（汽油车）摇号趋势分析.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/30%20%e5%ba%94%e7%94%a8%e5%bc%80%e5%8f%91%ef%bc%9a%e5%8c%97%e4%ba%ac%e5%b8%82%e5%b0%8f%e5%ae%a2%e8%bd%a6%ef%bc%88%e6%b1%bd%e6%b2%b9%e8%bd%a6%ef%bc%89%e6%91%87%e5%8f%b7%e8%b6%8b%e5%8a%bf%e5%88%86%e6%9e%90.md">30 应用开发：北京市小客车（汽油车）摇号趋势分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 性能调优：手把手带你提升应用的执行性能.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/31%20%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%ef%bc%9a%e6%89%8b%e6%8a%8a%e6%89%8b%e5%b8%a6%e4%bd%a0%e6%8f%90%e5%8d%87%e5%ba%94%e7%94%a8%e7%9a%84%e6%89%a7%e8%a1%8c%e6%80%a7%e8%83%bd.md">31 性能调优：手把手带你提升应用的执行性能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spark UI（上）深入解读Spark作业的“体检报告”.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/Spark%20UI%ef%bc%88%e4%b8%8a%ef%bc%89%e6%b7%b1%e5%85%a5%e8%a7%a3%e8%af%bbSpark%e4%bd%9c%e4%b8%9a%e7%9a%84%e2%80%9c%e4%bd%93%e6%a3%80%e6%8a%a5%e5%91%8a%e2%80%9d.md">Spark UI（上）深入解读Spark作业的“体检报告”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="Spark UI（下）：深入解读Spark作业的“体检报告”.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/Spark%20UI%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%85%a5%e8%a7%a3%e8%af%bbSpark%e4%bd%9c%e4%b8%9a%e7%9a%84%e2%80%9c%e4%bd%93%e6%a3%80%e6%8a%a5%e5%91%8a%e2%80%9d.md">Spark UI（下）：深入解读Spark作业的“体检报告”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="期末考试 “Spark性能调优”100分试卷等你来挑战！.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/%e6%9c%9f%e6%9c%ab%e8%80%83%e8%af%95%20%e2%80%9cSpark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e2%80%9d100%e5%88%86%e8%af%95%e5%8d%b7%e7%ad%89%e4%bd%a0%e6%9d%a5%e6%8c%91%e6%88%98%ef%bc%81.md">期末考试 “Spark性能调优”100分试卷等你来挑战！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 在时间面前，做一个笃定学习的人.md" href="/%e4%b8%93%e6%a0%8f/Spark%e6%80%a7%e8%83%bd%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%9c%a8%e6%97%b6%e9%97%b4%e9%9d%a2%e5%89%8d%ef%bc%8c%e5%81%9a%e4%b8%80%e4%b8%aa%e7%ac%83%e5%ae%9a%e5%ad%a6%e4%b9%a0%e7%9a%84%e4%ba%ba.md">结束语 在时间面前，做一个笃定学习的人.md</a>
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
                            <h1 id="title" data-id="24 Spark 3.0（一）：AQE的3个特性怎么才能用好？" class="title">24 Spark 3.0（一）：AQE的3个特性怎么才能用好？</h1>
                            <div><p>你好，我是吴磊。</p>

<p>目前，距离Spark 3.0版本的发布已经将近一年的时间了，这次版本升级添加了自适应查询执行（AQE）、动态分区剪裁（DPP）和扩展的 Join Hints 等新特性。利用好这些新特性，可以让我们的性能调优如虎添翼。因此，我会用三讲的时间和你聊聊它们。今天，我们先来说说AQE。</p>

<p>我发现，同学们在使用AQE的时候总会抱怨说：“AQE的开关打开了，相关的配置项也设了，可应用性能还是没有提升。”这往往是因为我们对于AQE的理解不够透彻，调优总是照葫芦画瓢，所以这一讲，我们就先从AQE的设计初衷说起，然后说说它的工作原理，最后再去探讨怎样才能用好AQE。</p>

<h2 id="spark为什么需要aqe">Spark为什么需要AQE？</h2>

<p>在2.0版本之前，Spark SQL仅仅支持启发式、静态的优化过程，就像我们在第21、22、23三讲介绍的一样。</p>

<p>启发式的优化又叫RBO（Rule Based Optimization，基于规则的优化），它往往基于一些规则和策略实现，如谓词下推、列剪枝，这些规则和策略来源于数据库领域已有的应用经验。也就是说，<strong>启发式的优化实际上算是一种经验主义</strong>。</p>

<p>经验主义的弊端就是不分青红皂白、胡子眉毛一把抓，对待相似的问题和场景都使用同一类套路。Spark社区正是因为意识到了RBO的局限性，因此在2.2版本中推出了CBO（Cost Based Optimization，基于成本的优化）。</p>

<p><strong>CBO的特点是“实事求是”</strong>，基于数据表的统计信息（如表大小、数据列分布）来选择优化策略。CBO支持的统计信息很丰富，比如数据表的行数、每列的基数（Cardinality）、空值数、最大值、最小值和直方图等等。因为有统计数据做支持，所以CBO选择的优化策略往往优于RBO选择的优化规则。</p>

<p>但是，<strong>CBO也面临三个方面的窘境：“窄、慢、静”</strong>。窄指的是适用面太窄，CBO仅支持注册到Hive Metastore的数据表，但在大量的应用场景中，数据源往往是存储在分布式文件系统的各类文件，如Parquet、ORC、CSV等等。</p>

<p>慢指的是统计信息的搜集效率比较低。对于注册到Hive Metastore的数据表，开发者需要调用ANALYZE TABLE COMPUTE STATISTICS语句收集统计信息，而各类信息的收集会消耗大量时间。</p>

<p>静指的是静态优化，这一点与RBO一样。CBO结合各类统计信息制定执行计划，一旦执行计划交付运行，CBO的使命就算完成了。换句话说，如果在运行时数据分布发生动态变化，CBO先前制定的执行计划并不会跟着调整、适配。</p>

<h2 id="aqe到底是什么">AQE到底是什么？</h2>

<p>考虑到RBO和CBO的种种限制，Spark在3.0版本推出了AQE（Adaptive Query Execution，自适应查询执行）。如果用一句话来概括，<strong>AQE是Spark SQL的一种动态优化机制，在运行时，每当Shuffle Map阶段执行完毕，AQE都会结合这个阶段的统计信息，基于既定的规则动态地调整、修正尚未执行的逻辑计划和物理计划，来完成对原始查询语句的运行时优化</strong>。</p>

<p>从定义中，我们不难发现，<strong>AQE优化机制触发的时机是Shuffle Map阶段执行完毕。也就是说，AQE优化的频次与执行计划中Shuffle的次数一致</strong>。反过来说，如果你的查询语句不会引入Shuffle操作，那么Spark SQL是不会触发AQE的。对于这样的查询，无论你怎么调整AQE相关的配置项，AQE也都爱莫能助。</p>

<p>对于AQE的定义，我相信你还有很多问题，比如，AQE依赖的统计信息具体是什么？既定的规则和策略具体指什么？接下来，我们一一来解答。</p>

<p><strong>首先，AQE赖以优化的统计信息与CBO不同，这些统计信息并不是关于某张表或是哪个列，而是Shuffle Map阶段输出的中间文件。</strong>学习过Shuffle的工作原理之后，我们知道，每个Map Task都会输出以data为后缀的数据文件，还有以index为结尾的索引文件，这些文件统称为中间文件。每个data文件的大小、空文件数量与占比、每个Reduce Task对应的分区大小，所有这些基于中间文件的统计值构成了AQE进行优化的信息来源。</p>

<p><strong>其次，结合Spark SQL端到端优化流程图我们可以看到，AQE从运行时获取统计信息，在条件允许的情况下，优化决策会分别作用到逻辑计划和物理计划。</strong></p>

<p><img src="assets/5d4d4646e12245e4858ed381bb1763e7.jpg" alt="" /></p>

<p><strong>AQE既定的规则和策略主要有4个，分为1个逻辑优化规则和3个物理优化策略。</strong>我把这些规则与策略，和相应的AQE特性，以及每个特性仰仗的统计信息，都汇总到了如下的表格中，你可以看一看。</p>

<p><img src="assets/fdd3f765482f49a981df53e5081c635f.jpg" alt="" /></p>

<h2 id="如何用好aqe">如何用好AQE？</h2>

<p>那么，AQE是如何根据Map阶段的统计信息以及这4个规则与策略，来动态地调整和修正尚未执行的逻辑计划和物理计划的呢？这就要提到AQE的三大特性，也就是Join策略调整、自动分区合并，以及自动倾斜处理，我们需要借助它们去分析AQE动态优化的过程。它们的基本概念我们在<a href="https://time.geekbang.org/column/article/357342" target="_blank">第9讲</a>说过，这里我再带你简单回顾一下。</p>

<ul>
<li>Join策略调整：如果某张表在过滤之后，尺寸小于广播变量阈值，这张表参与的数据关联就会从Shuffle Sort Merge Join降级（Demote）为执行效率更高的Broadcast Hash Join。</li>
<li>自动分区合并：在Shuffle过后，Reduce Task数据分布参差不齐，AQE将自动合并过小的数据分区。</li>
<li>自动倾斜处理：结合配置项，AQE自动拆分Reduce阶段过大的数据分区，降低单个Reduce Task的工作负载。</li>
</ul>

<p>接下来，我们就一起来分析这3个特性的动态优化过程。</p>

<h3 id="join策略调整">Join策略调整</h3>

<p>我们先来说说Join策略调整，这个特性涉及了一个逻辑规则和一个物理策略，它们分别是DemoteBroadcastHashJoin和OptimizeLocalShuffleReader。</p>

<p><strong>DemoteBroadcastHashJoin规则的作用，是把Shuffle Joins降级为Broadcast Joins。需要注意的是，这个规则仅适用于Shuffle Sort Merge Join这种关联机制，其他机制如Shuffle Hash Join、Shuffle Nested Loop Join都不支持。</strong>对于参与Join的两张表来说，在它们分别完成Shuffle Map阶段的计算之后，DemoteBroadcastHashJoin会判断中间文件是否满足如下条件：</p>

<ul>
<li>中间文件尺寸总和小于广播阈值spark.sql.autoBroadcastJoinThreshold</li>
<li>空文件占比小于配置项spark.sql.adaptive.nonEmptyPartitionRatioForBroadcastJoin</li>
</ul>

<p>只要有任意一张表的统计信息满足这两个条件，Shuffle Sort Merge Join就会降级为Broadcast Hash Join。说到这儿，你可能会问：“既然DemoteBroadcastHashJoin逻辑规则可以把Sort Merge Join转换为Broadcast Join，那同样用来调整Join策略的OptimizeLocalShuffleReader规则又是干什么用的呢？看上去有些多余啊！”</p>

<p>不知道你注意到没有，我一直强调，<strong>AQE依赖的统计信息来自于Shuffle Map阶段生成的中间文件</strong>。这意味什么呢？这就意味着AQE在开始优化之前，Shuffle操作已经执行过半了！</p>

<p>我们来举个例子，现在有两张表：事实表Order和维度表User，它们的查询语句和初始的执行计划如下。</p>

<pre><code>//订单表与用户表关联
select sum(order.price * order.volume), user.id
from order inner join user
on order.userId = user.id
where user.type = ‘Head Users’
group by user.id
</code></pre>

<p>由于两张表大都到超过了广播阈值，因此Spark SQL在最初的执行计划中选择了Sort Merge Join。AQE需要同时结合两个分支中的Shuffle（Exchange）输出，才能判断是否可以降级为Broadcast Join，以及用哪张表降级。这就意味着，不论大表还是小表都要完成Shuffle Map阶段的计算，并且把中间文件落盘，AQE才能做出决策。</p>

<p><img src="assets/12e0c35eff85471ab05dbeec1189c6da.jpg" alt="" /></p>

<p>你可能会说：“根本不需要大表做Shuffle呀，AQE只需要去判断小表Shuffle的中间文件就好啦”。可问题是，AQE可分不清哪张是大表、哪张是小表。在Shuffle Map阶段结束之前，数据表的尺寸大小对于AQE来说是“透明的”。因此，AQE必须等待两张表都完成Shuffle Map的计算，然后统计中间文件，才能判断降级条件是否成立，以及用哪张表做广播变量。</p>

<p>在常规的Shuffle计算流程中，Reduce阶段的计算需要跨节点访问中间文件拉取数据分片。如果遵循常规步骤，即便AQE在运行时把Shuffle Sort Merge Join降级为Broadcast Join，大表的中间文件还是需要通过网络进行分发。这个时候，AQE的动态Join策略调整也就失去了实用价值。原因很简单，负载最重的大表Shuffle计算已经完成，再去决定切换到Broadcast Join已经没有任何意义。</p>

<p>在这样的背景下，OptimizeLocalShuffleReader物理策略就非常重要了。既然大表已经完成Shuffle Map阶段的计算，这些计算可不能白白浪费掉。<strong>采取OptimizeLocalShuffleReader策略可以省去Shuffle常规步骤中的网络分发，Reduce Task可以就地读取本地节点（Local）的中间文件，完成与广播小表的关联操作。</strong></p>

<p>不过，需要我们特别注意的是，<strong>OptimizeLocalShuffleReader物理策略的生效与否由一个配置项决定</strong>。这个配置项是spark.sql.adaptive.localShuffleReader.enabled，尽管它的默认值是True，但是你千万不要把它的值改为False。否则，就像我们刚才说的，AQE的Join策略调整就变成了形同虚设。</p>

<p>说到这里，你可能会说：“这么看，AQE的Join策略调整有些鸡肋啊！毕竟Shuffle计算都已经过半，Shuffle Map阶段的内存消耗和磁盘I/O是半点没省！”确实，Shuffle Map阶段的计算开销是半点没省。但是，OptimizeLocalShuffleReader策略避免了Reduce阶段数据在网络中的全量分发，仅凭这一点，大多数的应用都能获益匪浅。因此，对于AQE的Join策略调整，<strong>我们可以用一个成语来形容：“亡羊补牢、犹未为晚”</strong>。</p>

<h3 id="自动分区合并">自动分区合并</h3>

<p>接下来，我们再来说说自动分区合并。分区合并的原理比较简单，<strong>在Reduce阶段，当Reduce Task从全网把数据分片拉回，AQE按照分区编号的顺序，依次把小于目标尺寸的分区合并在一起</strong>。目标分区尺寸由以下两个参数共同决定。这部分我们在第10讲详细讲过，如果不记得，你可以翻回去看一看。</p>

<ul>
<li>spark.sql.adaptive.advisoryPartitionSizeInBytes，由开发者指定分区合并后的推荐尺寸。</li>
<li>spark.sql.adaptive.coalescePartitions.minPartitionNum，分区合并后，分区数不能低于该值。</li>
</ul>

<p><img src="assets/b72466254f614ba2a868daf917a71cf9.jpg" alt="" /></p>

<p>除此之外，我们还要注意，在Shuffle Map阶段完成之后，AQE优化机制被触发，CoalesceShufflePartitions策略“无条件”地被添加到新的物理计划中。读取配置项、计算目标分区大小、依序合并相邻分区这些计算逻辑，在Tungsten WSCG的作用下融合进“手写代码”于Reduce阶段执行。</p>

<h3 id="自动倾斜处理">自动倾斜处理</h3>

<p>与自动分区合并相反，自动倾斜处理的操作是“拆”。在Reduce阶段，当Reduce Task所需处理的分区尺寸大于一定阈值时，利用OptimizeSkewedJoin策略，AQE会把大分区拆成多个小分区。倾斜分区和拆分粒度由以下这些配置项决定。关于它们的含义与作用，我们在<a href="https://time.geekbang.org/column/article/357342" target="_blank">第10讲</a>说过，你可以再翻回去看一看。</p>

<ul>
<li>spark.sql.adaptive.skewJoin.skewedPartitionFactor，判定倾斜的膨胀系数</li>
<li>spark.sql.adaptive.skewJoin.skewedPartitionThresholdInBytes，判定倾斜的最低阈值</li>
<li>spark.sql.adaptive.advisoryPartitionSizeInBytes，以字节为单位，定义拆分粒度</li>
</ul>

<p>自动倾斜处理的拆分操作也是在Reduce阶段执行的。在同一个Executor内部，本该由一个Task去处理的大分区，被AQE拆成多个小分区并交由多个Task去计算。这样一来，Task之间的计算负载就可以得到平衡。但是，这并不能解决不同Executors之间的负载均衡问题。</p>

<p>我们来举个例子，假设有个Shuffle操作，它的Map阶段有3个分区，Reduce阶段有4个分区。4个分区中的两个都是倾斜的大分区，而且这两个倾斜的大分区刚好都分发到了Executor 0。通过下图，我们能够直观地看到，尽管两个大分区被拆分，但横向来看，整个作业的主要负载还是落在了Executor 0的身上。Executor 0的计算能力依然是整个作业的瓶颈，这一点并没有因为分区拆分而得到实质性的缓解。</p>

<p><img src="assets/a443064257e148ada0f5b12d4972ce60.jpg" alt="" /></p>

<p>另外，在数据关联的场景中，对于参与Join的两张表，我们暂且把它们记做数据表1和数据表2，如果表1存在数据倾斜，表2不倾斜，那在关联的过程中，AQE除了对表1做拆分之外，还需要对表2对应的数据分区做复制，来保证关联关系不被破坏。</p>

<p><img src="assets/788cded1192649d792a72e2fae8f8a1c.jpg" alt="" /></p>

<p>在这样的运行机制下，如果两张表都存在数据倾斜怎么办？这个时候，事情就开始变得逐渐复杂起来了。对于上图中的表1和表2，我们假设表1还是拆出来两个分区，表2因为倾斜也拆出来两个分区。这个时候，为了不破坏逻辑上的关联关系，表1、表2拆分出来的分区还要各自复制出一份，如下图所示。</p>

<p><img src="assets/59ae8cb1d4eb4be6a9a8a8a76155c831.jpg" alt="" /></p>

<p>如果现在问题变得更复杂了，左表拆出M个分区，右表拆出N各分区，那么每张表最终都需要保持M x N份分区数据，才能保证关联逻辑的一致性。当M和N逐渐变大时，AQE处理数据倾斜所需的计算开销将会面临失控的风险。</p>

<p><strong>总的来说，当应用场景中的数据倾斜比较简单，比如虽然有倾斜但数据分布相对均匀，或是关联计算中只有一边倾斜，我们完全可以依赖AQE的自动倾斜处理机制。但是，当我们的场景中数据倾斜变得复杂，比如数据中不同Key的分布悬殊，或是参与关联的两表都存在大量的倾斜，我们就需要衡量AQE的自动化机制与手工处理倾斜之间的利害得失。</strong>关于手工处理倾斜，我们留到第28讲再去展开。</p>

<h2 id="小结">小结</h2>

<p>AQE是Spark SQL的一种动态优化机制，它的诞生解决了RBO、CBO，这些启发式、静态优化机制的局限性。想要用好AQE，我们就要掌握它的特点，以及它支持的三种优化特性的工作原理和使用方法。</p>

<p>如果用一句话来概括AQE的定义，就是每当Shuffle Map阶段执行完毕，它都会结合这个阶段的统计信息，根据既定的规则和策略动态地调整、修正尚未执行的逻辑计划和物理计划，从而完成对原始查询语句的运行时优化。也因此，只有当你的查询语句会引入Shuffle操作的时候，Spark SQL才会触发AQE。</p>

<p>AQE支持的三种优化特性分别是Join策略调整、自动分区合并和自动倾斜处理。</p>

<p>关于Join策略调整，我们首先要知道DemoteBroadcastHashJoin规则仅仅适用于Shuffle Sort Merge Join这种关联机制，对于其他Shuffle Joins类型，AQE暂不支持把它们转化为Broadcast Joins。其次，为了确保AQE的Join策略调整正常运行，我们要确保spark.sql.adaptive.localShuffleReader.enabled配置项始终为开启状态。</p>

<p>关于自动分区合并，我们要知道，在Shuffle Map阶段完成之后，结合分区推荐尺寸与分区数量限制，AQE会自动帮我们完成分区合并的计算过程。</p>

<p>关于AQE的自动倾斜处理我们要知道，它只能以Task为粒度缓解数据倾斜，并不能解决不同Executors之间的负载均衡问题。针对场景较为简单的倾斜问题，比如关联计算中只涉及单边倾斜，我们完全可以依赖AQE的自动倾斜处理机制。但是，当数据倾斜问题变得复杂的时候，我们需要衡量AQE的自动化机制与手工处理倾斜之间的利害得失。</p>

<h2 id="每日一练">每日一练</h2>

<ol>
<li>我们知道，AQE依赖的统计信息来源于Shuffle Map阶段输出的中间文件。你觉得，在运行时，AQE还有其他渠道可以获得同样的统计信息吗？</li>
<li>AQE的自动倾斜处理机制只能以Task为粒度来平衡工作负载，如果让你重新实现这个机制，你有什么更好的办法能让AQE以Executors为粒度做到负载均衡吗？</li>
</ol>

<p>期待在留言区看到你的思考和答案，我们下一讲见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#48242424717c7979787f082f25292124662b2725" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f126c22dc57ede4',t:'MTczNDA1NTg3NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
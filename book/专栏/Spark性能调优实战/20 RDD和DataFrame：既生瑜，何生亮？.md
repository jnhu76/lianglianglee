<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=20&#32;RDD和DataFrame：既生瑜，何生亮？>
        <link rel="icon" href="/static/favicon.png">
        <title>20 RDD和DataFrame：既生瑜，何生亮？ </title>
        
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
                            <h1 id="title" data-id="20 RDD和DataFrame：既生瑜，何生亮？" class="title">20 RDD和DataFrame：既生瑜，何生亮？</h1>
                            <div><p>你好，我是吴磊。</p>

<p>从今天开始，我们进入Spark SQL性能调优篇的学习。在这一篇中，我会先带你学习Spark SQL已有的优化机制，如Catalyst、Tungsten这些核心组件，以及AQE、DPP等新特性。深入理解这些内置的优化机制，会让你在开发应用之初就有一个比较高的起点。然后，针对数据分析中的典型场景，如数据关联，我们再去深入探讨性能调优的方法和技巧。</p>

<p>今天这一讲，我们先来说说RDD和DataFrame的渊源。这也是面试的时候，面试官经常会问的。比如说：“Spark 3.0大版本发布，Spark SQL的优化占比将近50%；而像PySpark、Mllib和Streaming的优化占比都不超过10%，Graph的占比几乎可以忽略不计。这是否意味着Spark社区逐渐放弃了其他计算领域，只专注于数据分析？”</p>

<p><a href="https://databricks.com/blog/2020/06/18/introducing-apache-spark-3-0-now-available-in-databricks-runtime-7-0.html" target="_blank"><img src="assets/d6510a30b69a452888816813d9ca173c.jpg" alt="" /></a></p>

<p>这个问题的标准答案是：“Spark SQL取代Spark Core，成为新一代的引擎内核，所有其他子框架如Mllib、Streaming和Graph，都可以共享Spark SQL的性能优化，都能从Spark社区对于Spark SQL的投入中受益。”不过，面试官可没有那么好对付，一旦你这么说，他/她可能会追问：“为什么需要Spark SQL这个新一代引擎内核？Spark Core有什么问题吗？Spark SQL解决了Spark Core的哪些问题？怎么解决的？”</p>

<p>面对这一连串“箭如雨发”的追问，你还能回答出来吗？接下来，我就从RDD的痛点说起，一步一步带你探讨DataFrame出现的必然性，Spark Core的局限性，以及它和Spark SQL的关系。</p>

<h2 id="rdd之痛-优化空间受限">RDD之痛：优化空间受限</h2>

<p>自从Spark社区在1.3版本发布了DataFrame，它就开始代替RDD，逐渐成为开发者的首选。我们知道，新抽象的诞生一定是为了解决老抽象不能搞定的问题。那么，这些问题都是什么呢？下面，我们就一起来分析一下。</p>

<p>在RDD的开发框架下，我们调用RDD算子进行适当的排列组合，就可以很轻松地实现业务逻辑。我把这些使用频繁的RDD算子总结到了下面的表格里，你可以看一看。</p>

<p><a href="https://spark.apache.org/docs/latest/rdd-programming-guide.html" target="_blank"><img src="assets/dfea632260164eb786fc3cbb8fe70f86.jpg" alt="" /></a></p>

<p>表格中高亮显示的就是RDD转换和聚合算子，它们都是高阶函数。高阶函数指的是形参包含函数的函数，或是返回结果包含函数的函数。为了叙述方便，我们把那些本身是高阶函数的RDD算子，简称“高阶算子”。</p>

<p><strong>对于这些高阶算子，开发者需要以Lambda函数的形式自行提供具体的计算逻辑。</strong>以map为例，我们需要明确对哪些字段做映射，以什么规则映射。再以filter为例，我们需要指明以什么条件在哪些字段上过滤。</p>

<p>但这样一来，Spark只知道开发者要做map、filter，但并不知道开发者打算怎么做map和filter。也就是说，<strong>在RDD的开发模式下，Spark Core只知道“做什么”，而不知道“怎么做”。</strong>这会让Spark Core两眼一抹黑，除了把Lambda函数用闭包的形式打发到Executors以外，实在是没有什么额外的优化空间。</p>

<p><strong>对于Spark Core来说，优化空间受限最主要的影响，莫过于让应用的执行性能变得低下。</strong>一个典型的例子，就是相比Java或者Scala，PySpark实现的应用在执行性能上相差悬殊。原因在于，在RDD的开发模式下，即便是同一个应用，不同语言实现的版本在运行时也会有着天壤之别。</p>

<p><img src="assets/f8e0aeeba425463784184c390475b01b.jpg" alt="" /></p>

<p>当我们使用Java或者Scala语言做开发时，所有的计算都在JVM进程内完成，如图中左侧的Spark计算节点所示。</p>

<p>而当我们在PySpark上做开发的时候，只能把由RDD算子构成的计算代码，一股脑地发送给Python进程。Python进程负责执行具体的脚本代码，完成计算之后，再把结果返回给Executor进程。由于每一个Task都需要一个Python进程，如果RDD的并行度为#N，那么整个集群就需要#N个这样的Python进程与Executors交互。不难发现，其中的任务调度、数据计算和数据通信等开销，正是PySpark性能低下的罪魁祸首。</p>

<h2 id="dataframe应运而生">DataFrame应运而生</h2>

<p>针对优化空间受限这个核心问题，Spark社区痛定思痛，在2013年在1.3版本中发布了DataFrame。那么，DataFrame的特点是什么，它和RDD又有什么不同呢？</p>

<p>首先，用一句话来概括，DataFrame就是携带数据模式（Data Schema）的结构化分布式数据集，而RDD是不带Schema的分布式数据集。<strong>因此，从数据表示（Data Representation）的角度来看，是否携带Schema是它们唯一的区别。</strong>带Schema的数据表示形式决定了DataFrame只能封装结构化数据，而RDD则没有这个限制，所以除了结构化数据，它还可以封装半结构化和非结构化数据。</p>

<p>其次，从开发API上看，<strong>RDD算子多是高阶函数，这些算子允许开发者灵活地实现业务逻辑，表达能力极强</strong>。</p>

<p>DataFrame的表达能力却很弱。一来，它定义了一套DSL（Domain Specific Language）算子，如select、filter、agg、groupBy等等。由于DSL语言是为解决某一类任务而专门设计的计算机语言，非图灵完备，因此，语言表达能力非常有限。二来，DataFrame中的绝大多数算子都是标量函数（Scalar Functions），它们的形参往往是结构化的数据列（Columns），表达能力也很弱。</p>

<p>你可能会问：“相比RDD，DataFrame的表示和表达能力都变弱了，那它是怎么解决RDD优化空间受限的核心痛点呢？”</p>

<p>当然，仅凭DataFrame在API上的改动就想解决RDD的核心痛点，比登天还难。<strong>DataFrame API最大的意义在于，它为Spark引擎的内核优化打开了全新的空间。</strong></p>

<p>首先，DataFrame中Schema所携带的类型信息，让Spark可以根据明确的字段类型设计定制化的数据结构，从而大幅提升数据的存储和访问效率。其次，DataFrame中标量算子确定的计算逻辑，让Spark可以基于启发式的规则和策略，甚至是动态的运行时信息去优化DataFrame的计算过程。</p>

<h2 id="spark-sql智能大脑">Spark SQL智能大脑</h2>

<p>那么问题来了，有了DataFrame API，负责引擎内核优化的那个幕后英雄是谁？为了支持DataFrame开发模式，Spark从1.3版本开始推出Spark SQL。<strong>Spark SQL的核心组件有二，其一是Catalyst优化器，其二是Tungsten。</strong>关于Catalyst和Tungsten的特性和优化过程，我们在后面的两讲再去展开，今天这一讲，咱们专注在它们和DataFrame的关系。</p>

<h3 id="catalyst-执行过程优化">Catalyst：执行过程优化</h3>

<p>我们先来说说Catalyst的优化过程。当开发者通过Actions算子触发DataFrame的计算请求时，Spark内部会发生一系列有趣的事情。</p>

<p>首先，基于DataFrame确切的计算逻辑，Spark会使用第三方的SQL解析器ANTLR来生成抽象语法树（AST，Abstract Syntax Tree）。既然是树，就会有节点和边这两个基本的构成元素。节点记录的是标量算子（如select、filter）的处理逻辑，边携带的是数据信息：关系表和数据列，如下图所示。这样的语法树描述了从源数据到DataFrame结果数据的转换过程。</p>

<p><img src="assets/822994014c804f60a5ff405f60db7702.jpg" alt="" /></p>

<p>在Spark中，语法树还有个别名叫做“Unresolved Logical Plan”。它正是Catalyst优化过程的起点。之所以取名“Unresolved”，是因为边上记录的关系表和数据列仅仅是一些字符串，还没有和实际数据对应起来。举个例子，Filter之后的那条边，输出的数据列是joinKey和payLoad。这些字符串的来源是DataFrame的DSL查询，Catalyst并不确定这些字段名是不是有效的，更不知道每个字段都是什么类型。</p>

<p>因此，<strong>Catalyst做的第一步优化，就是结合DataFrame的Schema信息，确认计划中的表名、字段名、字段类型与实际数据是否一致</strong>。这个过程也叫做把“Unresolved Logical Plan”转换成“Analyzed Logical Plan”。</p>

<p><img src="assets/bb86b42250614523b8415f54368cf26c.jpg" alt="" /></p>

<p>基于解析过后的“Analyzed Logical Plan”，Catalyst才能继续做优化。利用启发式的规则和执行策略，Catalyst最终把逻辑计划转换为可执行的物理计划。总之，Catalyst的优化空间来源DataFrame的开发模式。</p>

<h3 id="tungsten-数据结构优化">Tungsten：数据结构优化</h3>

<p>说完Catalyst，我接着再来说说Tungsten。在开发原则那一讲，我们提到过Tungsten使用定制化的数据结构Unsafe Row来存储数据，Unsafe Row的优点是存储效率高、GC效率高。Tungsten之所以能够设计这样的数据结构，仰仗的也是DataFrame携带的Schema。Unsafe Row我们之前讲过，这里我再带你简单回顾一下。</p>

<p><img src="assets/7ae4d24f9a4145eeb6c1d5251f663994.jpg" alt="" /></p>

<p>Tungsten是用二进制字节序列来存储每一条用户数据的，因此在存储效率上完胜Java Object。比如说，如果我们要存储上表中的数据，用Java Object来存储会消耗100个字节数，而使用Tungsten仅需要不到20个字节，如下图所示。</p>

<p><img src="assets/511e8ca128244d8cb3ef6bf2d20777c7.jpg" alt="" /></p>

<p>但是，要想实现上图中的二进制序列，Tungsten必须要知道数据条目的Schema才行。也就是说，它需要知道每一个字段的数据类型，才能决定在什么位置安放定长字段、安插Offset，以及存放变长字段的数据值。DataFrame刚好能满足这个前提条件。</p>

<p>我们不妨想象一下，如果数据是用RDD封装的，Tungsten还有可能做到这一点吗？当然不可能。这是因为，虽然RDD也带类型，如RDD[Int]、RDD[(Int, String)]，但如果RDD中携带的是开发者自定义的数据类型，如RDD[User]或是RDD[Product]，Tungsten就会两眼一抹黑，完全不知道你的User和Product抽象到底是什么。成也萧何、败也萧何，RDD的通用性是一柄双刃剑，在提供开发灵活性的同时，也让引擎内核的优化变得无比困难。</p>

<p><strong>总的来说，基于DataFrame简单的标量算子和明确的Schema定义，借助Catalyst优化器和Tungsten，Spark SQL有能力在运行时构建起一套端到端的优化机制。这套机制运用启发式的规则与策略，以及运行时的执行信息，将原本次优、甚至是低效的查询计划转换为高效的执行计划，从而提升端到端的执行性能。</strong>因此，在DataFrame的开发框架下，不论你使用哪种开发语言，开发者都能共享Spark SQL带来的性能福利。</p>

<p><img src="assets/df8f6d7e00354c49b9499ef267cf0b3b.jpg" alt="" /></p>

<p>最后，我们再来回顾最开始提到的面试题：“从2.0版本至今，Spark对于其他子框架的完善与优化，相比Spark SQL占比很低。这是否意味着，Spark未来的发展重心是数据分析，其他场景如机器学习、流计算会逐渐边缘化吗？”</p>

<p>最初，Spark SQL确实仅仅是运行SQL和DataFrame应用的子框架，但随着优化机制的日趋完善，Spark SQL逐渐取代Spark Core，演进为新一代的引擎内核。到目前为止，所有子框架的源码实现都已从RDD切换到DataFrame。因此，和PySpark一样，像Streaming、Graph、Mllib这些子框架实际上都是通过DataFrame API运行在Spark SQL之上，它们自然可以共享Spark SQL引入的种种优化机制。</p>

<p>形象地说，Spark SQL就像是Spark的智能大脑，凡是通过DataFrame这双“眼睛”看到的问题，都会经由智能大脑这个指挥中心，统筹地进行分析与优化，优化得到的行动指令，最终再交由Executors这些“四肢”去执行。</p>

<h2 id="小结">小结</h2>

<p>今天，我们围绕RDD的核心痛点，探讨了DataFrame出现的必然性，Spark Core的局限性，以及它和Spark SQL的关系，对Spark SQL有了更深刻的理解。</p>

<p>RDD的核心痛点是优化空间有限，它指的是RDD高阶算子中封装的函数对于Spark来说完全透明，因此Spark对于计算逻辑的优化无从下手。</p>

<p>相比RDD，DataFrame是携带Schema的分布式数据集，只能封装结构化数据。DataFrame的算子大多数都是普通的标量函数，以消费数据列为主。但是，DataFrame更弱的表示能力和表达能力，反而为Spark引擎的内核优化打开了全新的空间。</p>

<p>根据DataFrame简单的标量算子和明确的Schema定义，借助Catalyst优化器和Tungsten，Spark SQL有能力在运行时，构建起一套端到端的优化机制。这套机制运用启发式的规则与策略和运行时的执行信息，将原本次优、甚至是低效的查询计划转换为高效的执行计划，从而提升端到端的执行性能。</p>

<p>在DataFrame的开发模式下，所有子框架、以及PySpark，都运行在Spark SQL之上，都可以共享Spark SQL提供的种种优化机制，这也是为什么Spark历次发布新版本、Spark SQL占比最大的根本原因。</p>

<h2 id="每日一练">每日一练</h2>

<ol>
<li>Java Object在对象存储上为什么会有比较大的开销？JVM需要多少个字节才能存下字符串“abcd”？</li>
<li>在DataFrame的开发框架下， PySpark中还有哪些操作是“顽固分子”，会导致计算与数据在JVM进程与Python进程之间频繁交互？(提示：参考RDD的局限性，那些对Spark透明的计算逻辑，Spark是没有优化空间的)</li>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#375b5b5b0e030606070077505a565e5b1954585a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f126abb7a48ede4',t:'MTczNDA1NTgxNy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
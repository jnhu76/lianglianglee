<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=13&#32;广播变量（二）：如何让Spark&#32;SQL选择Broadcast&#32;Joins？>
        <link rel="icon" href="/static/favicon.png">
        <title>13 广播变量（二）：如何让Spark SQL选择Broadcast Joins？ </title>
        
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
                            <h1 id="title" data-id="13 广播变量（二）：如何让Spark SQL选择Broadcast Joins？" class="title">13 广播变量（二）：如何让Spark SQL选择Broadcast Joins？</h1>
                            <div><p>你好，我是吴磊。</p>

<p>上一讲我们说到，在数据关联场景中，广播变量是克制Shuffle的杀手锏，用Broadcast Joins取代Shuffle Joins可以大幅提升执行性能。但是，很多同学只会使用默认的广播变量，不会去调优。那么，我们该怎么保证Spark在运行时优先选择Broadcast Joins策略呢？</p>

<p>今天这一讲，我就围绕着数据关联场景，从配置项和开发API两个方面，帮你梳理出两类调优手段，让你能够游刃有余地运用广播变量。</p>

<h2 id="利用配置项强制广播">利用配置项强制广播</h2>

<p>我们先来从配置项的角度说一说，有哪些办法可以让Spark优先选择Broadcast Joins。在Spark SQL配置项那一讲，我们提到过spark.sql.autoBroadcastJoinThreshold这个配置项。它的设置值是存储大小，默认是10MB。它的含义是，<strong>对于参与Join的两张表来说，任意一张表的尺寸小于10MB，Spark就在运行时采用Broadcast Joins的实现方式去做数据关联。</strong>另外，AQE在运行时尝试动态调整Join策略时，也是基于这个参数来判定过滤后的数据表是否足够小，从而把原本的Shuffle Joins调整为Broadcast Joins。</p>

<p><img src="assets/45a518f9c6be4b399f5f293fbb16e515.jpg" alt="" /></p>

<p>为了方便你理解，我来举个例子。在数据仓库中，我们经常会看到两张表：一张是订单事实表，为了方便，我们把它记成Fact；另一张是用户维度表，记成Dim。事实表体量很大在100GB量级，维度表很小在1GB左右。两张表的Schema如下所示：</p>

<pre><code>//订单事实表Schema
orderID: Int
userID: Int
trxDate: Timestamp
productId: Int
price: Float
volume: Int
 
//用户维度表Schema
userID: Int
name: String
age: Int
gender: String
</code></pre>

<p>当Fact表和Dim表基于userID做关联的时候，由于两张表的尺寸大小都远超spark.sql.autoBroadcastJoinThreshold参数的默认值10MB，因此Spark不得不选择Shuffle Joins的实现方式。但如果我们把这个参数的值调整为2GB，因为Dim表的尺寸比2GB小，所以，Spark在运行时会选择把Dim表封装到广播变量里，并采用Broadcast Join的方式来完成两张表的数据关联。</p>

<p>显然，对于绝大多数的Join场景来说，autoBroadcastJoinThreshold参数的默认值10MB太低了，因为现在企业的数据体量都在TB，甚至PB级别。因此，想要有效地利用Broadcast Joins，我们需要把参数值调大，一般来说，2GB左右是个不错的选择。</p>

<p>现在我们已经知道了，<strong>使用广播阈值配置项让Spark优先选择Broadcast Joins的关键，就是要确保至少有一张表的存储尺寸小于广播阈值</strong>。</p>

<p>但是，在设置广播阈值的时候，不少同学都跟我抱怨：“我的数据量明明小于autoBroadcastJoinThreshold参数设定的广播阈值，为什么Spark SQL在运行时并没有选择Broadcast Joins呢？”</p>

<p>详细了解后我才知道，这些同学所说的数据量，<strong>其实指的是数据表在磁盘上的存储大小</strong>，比如用<code>ls</code>或是<code>du -sh</code>等系统命令查看文件得到的结果。要知道，同一份数据在内存中的存储大小往往会比磁盘中的存储大小膨胀数倍，甚至十数倍。这主要有两方面原因。</p>

<p>一方面，为了提升存储和访问效率，开发者一般采用Parquet或是ORC等压缩格式把数据落盘。这些高压缩比的磁盘数据展开到内存之后，数据量往往会翻上数倍。</p>

<p>另一方面，受限于对象管理机制，在堆内内存中，JVM往往需要比数据原始尺寸更大的内存空间来存储对象。</p>

<p>我们来举个例子，字符串“abcd”按理说只需要消耗4个字节，但是，JVM在堆内存储这4个字符串总共需要消耗48个字节！那在运行时，一份看上去不大的磁盘数据展开到内存，翻上个4、5倍并不稀奇。因此，如果你按照磁盘上的存储大小去配置autoBroadcastJoinThreshold广播阈值，大概率也会遇到同样的窘境。</p>

<p><strong>那么问题来了，有什么办法能准确地预估一张表在内存中的存储大小呢？</strong></p>

<p>首先，我们要避开一个坑。我发现，有很多资料推荐用Spark内置的SizeEstimator去预估分布式数据集的存储大小。结合多次实战和踩坑经验，咱们必须要指出，<strong>SizeEstimator的估算结果不准确</strong>。因此，你可以直接跳过这种方法，这也能节省你调优的时间和精力。</p>

<p>我认为比较靠谱的办法是：<strong>第一步，把要预估大小的数据表缓存到内存，比如直接在DataFrame或是Dataset上调用cache方法；第二步，读取Spark SQL执行计划的统计数据</strong>。这是因为，Spark SQL在运行时，就是靠这些统计数据来制定和调整执行策略的。</p>

<pre><code>val df: DataFrame = _
df.cache.count
 
val plan = df.queryExecution.logical
val estimated: BigInt = spark
.sessionState
.executePlan(plan)
.optimizedPlan
.stats
.sizeInBytes
</code></pre>

<p>你可能会说：“这种办法虽然精确，但是这么做，实际上已经是在运行时进行调优了。把数据先缓存到内存，再去计算它的存储尺寸，当然更准确了。”没错，采用这种计算方式，调优所需花费的时间和精力确实更多，但在很多情况下，尤其是Shuffle Joins的执行效率让你痛不欲生的时候，这样的付出是值得的。</p>

<h2 id="利用api强制广播">利用API强制广播</h2>

<p>既然数据量的预估这么麻烦，有没有什么办法，不需要配置广播阈值，就可以让Spark SQL选择Broadcast Joins？还真有，而且办法还不止一种。</p>

<p>开发者可以通过Join Hints或是SQL functions中的broadcast函数，来强制Spark SQL在运行时采用Broadcast Joins的方式做数据关联。下面我就来分别讲一讲它们的含义和作用，以及该如何使用。必须要说明的是，这两种方式是等价的，并无优劣之分，只不过有了多样化的选择之后，你就可以根据自己的偏好和习惯来灵活地进行开发。</p>

<h3 id="用join-hints强制广播">用Join Hints强制广播</h3>

<p>Join Hints中的Hints表示“提示”，它指的是在开发过程中使用特殊的语法，明确告知Spark SQL在运行时采用哪种Join策略。一旦你启用了Join Hints，不管你的数据表是不是满足广播阈值，Spark SQL都会尽可能地尊重你的意愿和选择，使用Broadcast Joins去完成数据关联。</p>

<p>我们来举个例子，假设有两张表，一张表的内存大小在100GB量级，另一张小一些，2GB左右。在广播阈值被设置为2GB的情况下，并没有触发Broadcast Joins，但我们又不想花费时间和精力去精确计算小表的内存占用到底是多大。在这种情况下，我们就可以用Join Hints来帮我们做优化，仅仅几句提示就可以帮我们达到目的。</p>

<pre><code>val table1: DataFrame = spark.read.parquet(path1)
val table2: DataFrame = spark.read.parquet(path2)
table1.createOrReplaceTempView(&quot;t1&quot;)
table2.createOrReplaceTempView(&quot;t2&quot;)
 
val query: String = “select /*+ broadcast(t2) */ * from t1 inner join t2 on t1.key = t2.key”
val queryResutls: DataFrame = spark.sql(query)
</code></pre>

<p>你看，在上面的代码示例中，只要在SQL结构化查询语句里面加上一句<code>/*+ broadcast(t2) */</code>提示，我们就可以强制Spark SQL对小表t2进行广播，在运行时选择Broadcast Joins的实现方式。提示语句中的关键字，除了使用broadcast外，我们还可以用broadcastjoin或者mapjoin，它们实现的效果都一样。</p>

<p>如果你不喜欢用SQL结构化查询语句，尤其不想频繁地在Spark SQL上下文中注册数据表，你也可以在DataFrame的DSL语法中使用Join Hints。</p>

<pre><code>table1.join(table2.hint(“broadcast”), Seq(“key”), “inner”)

</code></pre>

<p>在上面的DSL语句中，我们只要在table2上调用hint方法，然后指定broadcast关键字，就可以同样达到强制广播表2的效果。</p>

<p>总之，Join Hints让开发者可以灵活地选择运行时的Join策略，对于熟悉业务、了解数据的同学来说，Join Hints允许开发者把专家经验凌驾于Spark SQL的优化引擎之上，更好地服务业务。</p>

<p>不过，Join Hints也有个小缺陷。如果关键字拼写错误，Spark SQL在运行时并不会显示地抛出异常，而是默默地忽略掉拼写错误的hints，假装它压根不存在。因此，在使用Join Hints的时候，需要我们在编译时自行确认Debug和纠错。</p>

<h3 id="用broadcast函数强制广播">用broadcast函数强制广播</h3>

<p>如果你不想等到运行时才发现问题，想让编译器帮你检查类似的拼写错误，那么你可以使用强制广播的第二种方式：broadcast函数。这个函数是类库org.apache.spark.sql.functions中的broadcast函数。调用方式非常简单，比Join Hints还要方便，只需要用broadcast函数封装需要广播的数据表即可，如下所示。</p>

<pre><code>import org.apache.spark.sql.functions.broadcast
table1.join(broadcast(table2), Seq(“key”), “inner”)
</code></pre>

<p>你可能会问：“既然开发者可以通过Join Hints和broadcast函数强制Spark SQL选择Broadcast Joins，那我是不是就可以不用理会广播阈值的配置项了？”其实还真不是。我认为，<strong>以广播阈值配置为主，以强制广播为辅</strong>，往往是不错的选择。</p>

<p><strong>广播阈值的设置，更多的是把选择权交给Spark SQL，尤其是在AQE的机制下，动态Join策略调整需要这样的设置在运行时做出选择。强制广播更多的是开发者以专家经验去指导Spark SQL该如何选择运行时策略。</strong>二者相辅相成，并不冲突，开发者灵活地运用就能平衡Spark SQL优化策略与专家经验在应用中的比例。</p>

<h2 id="广播变量不是银弹">广播变量不是银弹</h2>

<p>不过，虽然我们一直在强调，数据关联场景中广播变量是克制Shuffle的杀手锏，但广播变量并不是银弹。</p>

<p>就像有的同学会说：“开发者有这么多选项，甚至可以强制Spark选择Broadcast Joins，那我们是不是可以把所有Join操作都用Broadcast Joins来实现？”答案当然是否定的，广播变量不能解决所有的数据关联问题。</p>

<p><strong>首先，从性能上来讲，Driver在创建广播变量的过程中，需要拉取分布式数据集所有的数据分片。</strong>在这个过程中，网络开销和Driver内存会成为性能隐患。广播变量尺寸越大，额外引入的性能开销就会越多。更何况，如果广播变量大小超过8GB，Spark会直接抛异常中断任务执行。</p>

<p><strong>其次，从功能上来讲，并不是所有的Joins类型都可以转换为Broadcast Joins。</strong>一来，Broadcast Joins不支持全连接（Full Outer Joins）；二来，在所有的数据关联中，我们不能广播基表。或者说，即便开发者强制广播基表，也无济于事。比如说，在左连接（Left Outer Join）中，我们只能广播右表；在右连接（Right Outer Join）中，我们只能广播左表。在下面的代码中，即便我们强制用broadcast函数进行广播，Spark SQL在运行时还是会选择Shuffle Joins。</p>

<pre><code>import org.apache.spark.sql.functions.broadcast
broadcast (table1).join(table2, Seq(“key”), “left”)
table1.join(broadcast(table2), Seq(“key”), “right”)
</code></pre>

<h2 id="小结">小结</h2>

<p>这一讲，我们总结了2种方法，让Spark SQL在运行时能够选择Broadcast Joins策略，分别是设置配置项和用API强制广播。</p>

<p><strong>首先，设置配置项主要是设置autoBroadcastJoinThreshold配置项。</strong>开发者通过这个配置项指示Spark SQL优化器。只要参与Join的两张表中，有一张表的尺寸小于这个参数值，就在运行时采用Broadcast Joins的实现方式。</p>

<p>为了让Spark SQL采用Broadcast Joins，开发者要做的，就是让数据表在内存中的尺寸小于autoBroadcastJoinThreshold参数的设定值。</p>

<p>除此之外，在设置广播阈值的时候，因为磁盘数据展开到内存的时候，存储大小会成倍增加，往往导致Spark SQL无法采用Broadcast Joins的策略。因此，我们在做数据关联的时候，还要先预估一张表在内存中的存储大小。一种精确的预估方法是先把DataFrame缓存，然后读取执行计划的统计数据。</p>

<p><strong>其次，用API强制广播有两种方法，分别是设置Join Hints和用broadcast函数。</strong>设置Join Hints的方法就是在SQL结构化查询语句里面加上一句“/*+ broadcast(某表) */”的提示就可以了，这里的broadcast关键字也可以换成broadcastjoin或者mapjoin。另外，你也可以在DataFrame的DSL语法中使用调用hint方法，指定broadcast关键字，来达到同样的效果。设置broadcast函数的方法非常简单，只要用broadcast函数封装需要广播的数据表就可以了。</p>

<p>总的来说，不管是设置配置项还是用API强制广播都有各自的优缺点，所以，<strong>以广播阈值配置为主、强制广播为辅</strong>，往往是一个不错的选择。</p>

<p>最后，不过，我们也要注意，广播变量不是银弹，它并不能解决所有的数据关联问题，所以在日常的开发工作中，你要注意避免滥用广播。</p>

<h2 id="每日一练">每日一练</h2>

<ol>
<li>除了broadcast关键字外，在Spark 3.0版本中，Join Hints还支持哪些关联类型和关键字？</li>
<li>DataFrame可以用sparkContext.broadcast函数来广播吗？它和org.apache.spark.sql.functions.broadcast函数之间的区别是什么？</li>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#533f3f3f6a676262636413343e323a3f7d303c3e" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1268f7dc76ede4',t:'MTczNDA1NTc0NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
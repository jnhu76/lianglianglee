<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=22&#32;Spark&#32;UI（下）：如何高效地定位性能问题？>
        <link rel="icon" href="/static/favicon.png">
        <title>22 Spark UI（下）：如何高效地定位性能问题？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 入门Spark，你需要学会“三步走”.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%85%a5%e9%97%a8Spark%ef%bc%8c%e4%bd%a0%e9%9c%80%e8%a6%81%e5%ad%a6%e4%bc%9a%e2%80%9c%e4%b8%89%e6%ad%a5%e8%b5%b0%e2%80%9d.md">00 开篇词 入门Spark，你需要学会“三步走”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 Spark：从“大数据的Hello World”开始.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/01%20Spark%ef%bc%9a%e4%bb%8e%e2%80%9c%e5%a4%a7%e6%95%b0%e6%8d%ae%e7%9a%84Hello%20World%e2%80%9d%e5%bc%80%e5%a7%8b.md">01 Spark：从“大数据的Hello World”开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 RDD与编程模型：延迟计算是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/02%20RDD%e4%b8%8e%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%bb%b6%e8%bf%9f%e8%ae%a1%e7%ae%97%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">02 RDD与编程模型：延迟计算是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 RDD常用算子（一）：RDD内部的数据转换.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/03%20RDD%e5%b8%b8%e7%94%a8%e7%ae%97%e5%ad%90%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aRDD%e5%86%85%e9%83%a8%e7%9a%84%e6%95%b0%e6%8d%ae%e8%bd%ac%e6%8d%a2.md">03 RDD常用算子（一）：RDD内部的数据转换.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 进程模型与分布式部署：分布式计算是怎么回事？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/04%20%e8%bf%9b%e7%a8%8b%e6%a8%a1%e5%9e%8b%e4%b8%8e%e5%88%86%e5%b8%83%e5%bc%8f%e9%83%a8%e7%bd%b2%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%e8%ae%a1%e7%ae%97%e6%98%af%e6%80%8e%e4%b9%88%e5%9b%9e%e4%ba%8b%ef%bc%9f.md">04 进程模型与分布式部署：分布式计算是怎么回事？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 调度系统：如何把握分布式计算的精髓？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/05%20%e8%b0%83%e5%ba%a6%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8a%8a%e6%8f%a1%e5%88%86%e5%b8%83%e5%bc%8f%e8%ae%a1%e7%ae%97%e7%9a%84%e7%b2%be%e9%ab%93%ef%bc%9f.md">05 调度系统：如何把握分布式计算的精髓？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Shuffle管理：为什么Shuffle是性能瓶颈？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/06%20Shuffle%e7%ae%a1%e7%90%86%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Shuffle%e6%98%af%e6%80%a7%e8%83%bd%e7%93%b6%e9%a2%88%ef%bc%9f.md">06 Shuffle管理：为什么Shuffle是性能瓶颈？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 RDD常用算子（二）：Spark如何实现数据聚合？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/07%20RDD%e5%b8%b8%e7%94%a8%e7%ae%97%e5%ad%90%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aSpark%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%95%b0%e6%8d%ae%e8%81%9a%e5%90%88%ef%bc%9f.md">07 RDD常用算子（二）：Spark如何实现数据聚合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 内存管理：Spark如何使用内存？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/08%20%e5%86%85%e5%ad%98%e7%ae%a1%e7%90%86%ef%bc%9aSpark%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e5%86%85%e5%ad%98%ef%bc%9f.md">08 内存管理：Spark如何使用内存？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 RDD常用算子（三）：数据的准备、重分布与持久化.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/09%20RDD%e5%b8%b8%e7%94%a8%e7%ae%97%e5%ad%90%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%95%b0%e6%8d%ae%e7%9a%84%e5%87%86%e5%a4%87%e3%80%81%e9%87%8d%e5%88%86%e5%b8%83%e4%b8%8e%e6%8c%81%e4%b9%85%e5%8c%96.md">09 RDD常用算子（三）：数据的准备、重分布与持久化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 广播变量 &amp; 累加器：共享变量是用来做什么的？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/10%20%e5%b9%bf%e6%92%ad%e5%8f%98%e9%87%8f%20&amp;%20%e7%b4%af%e5%8a%a0%e5%99%a8%ef%bc%9a%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e6%98%af%e7%94%a8%e6%9d%a5%e5%81%9a%e4%bb%80%e4%b9%88%e7%9a%84%ef%bc%9f.md">10 广播变量 &amp; 累加器：共享变量是用来做什么的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 存储系统：数据到底都存哪儿了？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/11%20%e5%ad%98%e5%82%a8%e7%b3%bb%e7%bb%9f%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%88%b0%e5%ba%95%e9%83%bd%e5%ad%98%e5%93%aa%e5%84%bf%e4%ba%86%ef%bc%9f.md">11 存储系统：数据到底都存哪儿了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 基础配置详解：哪些参数会影响应用程序稳定性？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/12%20%e5%9f%ba%e7%a1%80%e9%85%8d%e7%bd%ae%e8%af%a6%e8%a7%a3%ef%bc%9a%e5%93%aa%e4%ba%9b%e5%8f%82%e6%95%b0%e4%bc%9a%e5%bd%b1%e5%93%8d%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%9f.md">12 基础配置详解：哪些参数会影响应用程序稳定性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Spark SQL：让我们从“小汽车摇号分析”开始.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/13%20Spark%20SQL%ef%bc%9a%e8%ae%a9%e6%88%91%e4%bb%ac%e4%bb%8e%e2%80%9c%e5%b0%8f%e6%b1%bd%e8%bd%a6%e6%91%87%e5%8f%b7%e5%88%86%e6%9e%90%e2%80%9d%e5%bc%80%e5%a7%8b.md">13 Spark SQL：让我们从“小汽车摇号分析”开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 台前幕后：DataFrame与Spark SQL的由来.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/14%20%e5%8f%b0%e5%89%8d%e5%b9%95%e5%90%8e%ef%bc%9aDataFrame%e4%b8%8eSpark%20SQL%e7%9a%84%e7%94%b1%e6%9d%a5.md">14 台前幕后：DataFrame与Spark SQL的由来.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 数据源与数据格式：DataFrame从何而来？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/15%20%e6%95%b0%e6%8d%ae%e6%ba%90%e4%b8%8e%e6%95%b0%e6%8d%ae%e6%a0%bc%e5%bc%8f%ef%bc%9aDataFrame%e4%bb%8e%e4%bd%95%e8%80%8c%e6%9d%a5%ef%bc%9f.md">15 数据源与数据格式：DataFrame从何而来？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 数据转换：如何在DataFrame之上做数据处理？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/16%20%e6%95%b0%e6%8d%ae%e8%bd%ac%e6%8d%a2%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8DataFrame%e4%b9%8b%e4%b8%8a%e5%81%9a%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%ef%bc%9f.md">16 数据转换：如何在DataFrame之上做数据处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 数据关联：不同的关联形式与实现机制该怎么选？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/17%20%e6%95%b0%e6%8d%ae%e5%85%b3%e8%81%94%ef%bc%9a%e4%b8%8d%e5%90%8c%e7%9a%84%e5%85%b3%e8%81%94%e5%bd%a2%e5%bc%8f%e4%b8%8e%e5%ae%9e%e7%8e%b0%e6%9c%ba%e5%88%b6%e8%af%a5%e6%80%8e%e4%b9%88%e9%80%89%ef%bc%9f.md">17 数据关联：不同的关联形式与实现机制该怎么选？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 数据关联优化：都有哪些Join策略，开发者该如何取舍？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/18%20%e6%95%b0%e6%8d%ae%e5%85%b3%e8%81%94%e4%bc%98%e5%8c%96%ef%bc%9a%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9bJoin%e7%ad%96%e7%95%a5%ef%bc%8c%e5%bc%80%e5%8f%91%e8%80%85%e8%af%a5%e5%a6%82%e4%bd%95%e5%8f%96%e8%88%8d%ef%bc%9f.md">18 数据关联优化：都有哪些Join策略，开发者该如何取舍？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 配置项详解：哪些参数会影响应用程序执行性能？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/19%20%e9%85%8d%e7%bd%ae%e9%a1%b9%e8%af%a6%e8%a7%a3%ef%bc%9a%e5%93%aa%e4%ba%9b%e5%8f%82%e6%95%b0%e4%bc%9a%e5%bd%b1%e5%93%8d%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e6%89%a7%e8%a1%8c%e6%80%a7%e8%83%bd%ef%bc%9f.md">19 配置项详解：哪些参数会影响应用程序执行性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Hive &#43; Spark强强联合：分布式数仓的不二之选.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/20%20Hive%20&#43;%20Spark%e5%bc%ba%e5%bc%ba%e8%81%94%e5%90%88%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%e6%95%b0%e4%bb%93%e7%9a%84%e4%b8%8d%e4%ba%8c%e4%b9%8b%e9%80%89.md">20 Hive &#43; Spark强强联合：分布式数仓的不二之选.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Spark UI（上）：如何高效地定位性能问题？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/21%20Spark%20UI%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%9c%b0%e5%ae%9a%e4%bd%8d%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%ef%bc%9f.md">21 Spark UI（上）：如何高效地定位性能问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Spark UI（下）：如何高效地定位性能问题？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/22%20Spark%20UI%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%9c%b0%e5%ae%9a%e4%bd%8d%e6%80%a7%e8%83%bd%e9%97%ae%e9%a2%98%ef%bc%9f.md">22 Spark UI（下）：如何高效地定位性能问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 Spark MLlib：从“房价预测”开始.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/23%20Spark%20MLlib%ef%bc%9a%e4%bb%8e%e2%80%9c%e6%88%bf%e4%bb%b7%e9%a2%84%e6%b5%8b%e2%80%9d%e5%bc%80%e5%a7%8b.md">23 Spark MLlib：从“房价预测”开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 特征工程（上）：有哪些常用的特征处理函数？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/24%20%e7%89%b9%e5%be%81%e5%b7%a5%e7%a8%8b%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%9c%89%e5%93%aa%e4%ba%9b%e5%b8%b8%e7%94%a8%e7%9a%84%e7%89%b9%e5%be%81%e5%a4%84%e7%90%86%e5%87%bd%e6%95%b0%ef%bc%9f.md">24 特征工程（上）：有哪些常用的特征处理函数？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 特征工程（下）：有哪些常用的特征处理函数？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/25%20%e7%89%b9%e5%be%81%e5%b7%a5%e7%a8%8b%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%9c%89%e5%93%aa%e4%ba%9b%e5%b8%b8%e7%94%a8%e7%9a%84%e7%89%b9%e5%be%81%e5%a4%84%e7%90%86%e5%87%bd%e6%95%b0%ef%bc%9f.md">25 特征工程（下）：有哪些常用的特征处理函数？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 模型训练（上）：决策树系列算法详解.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/26%20%e6%a8%a1%e5%9e%8b%e8%ae%ad%e7%bb%83%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%86%b3%e7%ad%96%e6%a0%91%e7%b3%bb%e5%88%97%e7%ae%97%e6%b3%95%e8%af%a6%e8%a7%a3.md">26 模型训练（上）：决策树系列算法详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 模型训练（中）：回归、分类和聚类算法详解.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/27%20%e6%a8%a1%e5%9e%8b%e8%ae%ad%e7%bb%83%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e5%9b%9e%e5%bd%92%e3%80%81%e5%88%86%e7%b1%bb%e5%92%8c%e8%81%9a%e7%b1%bb%e7%ae%97%e6%b3%95%e8%af%a6%e8%a7%a3.md">27 模型训练（中）：回归、分类和聚类算法详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 模型训练（下）：协同过滤与频繁项集算法详解.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/28%20%e6%a8%a1%e5%9e%8b%e8%ae%ad%e7%bb%83%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%e4%b8%8e%e9%a2%91%e7%b9%81%e9%a1%b9%e9%9b%86%e7%ae%97%e6%b3%95%e8%af%a6%e8%a7%a3.md">28 模型训练（下）：协同过滤与频繁项集算法详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 Spark MLlib Pipeline：高效开发机器学习应用.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/29%20Spark%20MLlib%20Pipeline%ef%bc%9a%e9%ab%98%e6%95%88%e5%bc%80%e5%8f%91%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e5%ba%94%e7%94%a8.md">29 Spark MLlib Pipeline：高效开发机器学习应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 Structured Streaming：从“流动的Word Count”开始.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/30%20Structured%20Streaming%ef%bc%9a%e4%bb%8e%e2%80%9c%e6%b5%81%e5%8a%a8%e7%9a%84Word%20Count%e2%80%9d%e5%bc%80%e5%a7%8b.md">30 Structured Streaming：从“流动的Word Count”开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 新一代流处理框架：Batch mode和Continuous mode哪家强？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/31%20%e6%96%b0%e4%b8%80%e4%bb%a3%e6%b5%81%e5%a4%84%e7%90%86%e6%a1%86%e6%9e%b6%ef%bc%9aBatch%20mode%e5%92%8cContinuous%20mode%e5%93%aa%e5%ae%b6%e5%bc%ba%ef%bc%9f.md">31 新一代流处理框架：Batch mode和Continuous mode哪家强？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 Window操作&amp;Watermark：流处理引擎提供了哪些优秀机制？.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/32%20Window%e6%93%8d%e4%bd%9c&amp;Watermark%ef%bc%9a%e6%b5%81%e5%a4%84%e7%90%86%e5%bc%95%e6%93%8e%e6%8f%90%e4%be%9b%e4%ba%86%e5%93%aa%e4%ba%9b%e4%bc%98%e7%a7%80%e6%9c%ba%e5%88%b6%ef%bc%9f.md">32 Window操作&amp;Watermark：流处理引擎提供了哪些优秀机制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 流计算中的数据关联：流与流、流与批.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/33%20%e6%b5%81%e8%ae%a1%e7%ae%97%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e5%85%b3%e8%81%94%ef%bc%9a%e6%b5%81%e4%b8%8e%e6%b5%81%e3%80%81%e6%b5%81%e4%b8%8e%e6%89%b9.md">33 流计算中的数据关联：流与流、流与批.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 Spark &#43; Kafka：流计算中的“万金油”.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/34%20Spark%20&#43;%20Kafka%ef%bc%9a%e6%b5%81%e8%ae%a1%e7%ae%97%e4%b8%ad%e7%9a%84%e2%80%9c%e4%b8%87%e9%87%91%e6%b2%b9%e2%80%9d.md">34 Spark &#43; Kafka：流计算中的“万金油”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 小王：保持空杯心态，不做井底之蛙.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e5%b0%8f%e7%8e%8b%ef%bc%9a%e4%bf%9d%e6%8c%81%e7%a9%ba%e6%9d%af%e5%bf%83%e6%80%81%ef%bc%8c%e4%b8%8d%e5%81%9a%e4%ba%95%e5%ba%95%e4%b9%8b%e8%9b%99.md">用户故事 小王：保持空杯心态，不做井底之蛙.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 进入时间裂缝，持续学习.md" href="/%e4%b8%93%e6%a0%8f/%e9%9b%b6%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8Spark/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e8%bf%9b%e5%85%a5%e6%97%b6%e9%97%b4%e8%a3%82%e7%bc%9d%ef%bc%8c%e6%8c%81%e7%bb%ad%e5%ad%a6%e4%b9%a0.md">结束语 进入时间裂缝，持续学习.md</a>
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
                            <h1 id="title" data-id="22 Spark UI（下）：如何高效地定位性能问题？" class="title">22 Spark UI（下）：如何高效地定位性能问题？</h1>
                            <div><p>你好，我是吴磊。</p>

<p>在上一讲，我们一起梳理了Spark UI的一级入口。其中Executors、Environment、Storage是详情页，开发者可以通过这3个页面，迅速地了解集群整体的计算负载、运行环境，以及数据集缓存的详细情况。不过SQL、Jobs、Stages，更多地是一种罗列式的展示，想要了解其中的细节，还需要进入到二级入口。</p>

<p>沿用之前的比喻，身为“大夫”的开发者想要结合经验，迅速定位“病灶”，离不开各式各样的指标项。而今天要讲的二级入口，相比一级入口，内容更加丰富、详尽。要想成为一名“临床经验丰富”的老医生，咱们先要做到熟练解读这些度量指标。</p>

<p><img src="assets/56563537c4e0ef597629d42618df21d2.png" alt="图片" title="Spark UI导航条：一级入口" /></p>

<p>所谓二级入口，它指的是，<strong>通过一次超链接跳转才能访问到的页面</strong>。对于SQL、Jobs和Stages这3类入口来说，二级入口往往已经提供了足够的信息，基本覆盖了“体检报告”的全部内容。因此，尽管Spark UI也提供了少量的三级入口（需要两跳才能到达的页面），但是这些隐藏在“犄角旮旯”的三级入口，往往并不需要开发者去特别关注。</p>

<p>接下来，我们就沿着SQL -&gt; Jobs -&gt; Stages的顺序，依次地去访问它们的二级入口，从而针对全局DAG、作业以及执行阶段，获得更加深入的探索与洞察。</p>

<h4 id="sql详情页">SQL详情页</h4>

<p>在SQL Tab一级入口，我们看到有3个条目，分别是count（统计申请编号）、count（统计中签编号）和save。前两者的计算过程，都是读取数据源、缓存数据并触发缓存的物化，相对比较简单，因此，我们把目光放在<strong>save</strong>这个条目上。</p>

<p><img src="assets/dd3231ca21492ff00c63a111d96516cb.png" alt="图片" title="SQL概览页" /></p>

<p>点击图中的“save at:27”，即可进入到该作业的执行计划页面，如下图所示。</p>

<p><img src="assets/5e9fa6231dc66db829bb043446c73db9.png" alt="图片" title="SQL页面二级入口（部分截图）" /></p>

<p>为了聚焦重点，这里我们仅截取了部分的执行计划，想要获取完整的执行计划，你可以通过访问<a href="https://github.com/wulei-bj-cn/learn-spark/blob/main/chapter22/demo%20-%20Details%20for%20Query%202.webarchive" target="_blank">这里</a>来获得。为了方便你阅读，这里我手绘出了执行计划的示意图，供你参考，如下图所示。</p>

<p><img src="assets/0d5db8ed21155563ec2a9bb8204368dd.jpg" alt="图片" title="“倍率与中签率分析”SQL执行计划示意图" /></p>

<p>可以看到，“倍率与中签率分析”应用的计算过程，非常具有代表性，它涵盖了数据分析场景中大部分的操作，也即<strong>过滤、投影、关联、分组聚合和排序</strong>。图中红色的部分为Exchange，代表的是Shuffle操作，蓝色的部分为Sort，也就是排序，而绿色的部分是Aggregate，表示的是（局部与全局的）数据聚合。</p>

<p>无疑，这三部分是硬件资源的主要消费者，同时，对于这3类操作，Spark UI更是提供了详细的Metrics来刻画相应的硬件资源消耗。接下来，咱们就重点研究一下这3类操作的度量指标。</p>

<p><strong>Exchange</strong></p>

<p>下图中并列的两个Exchange，对应的是示意图中SortMergeJoin之前的两个Exchange。它们的作用是对申请编码数据与中签编码数据做Shuffle，为数据关联做准备。</p>

<p><img src="assets/e506b76d435de956e4e20d62f82e10dc.png" alt="图片" title="Exchange（左）：申请编码数据Shuffle，Exchange（右）：中签编码数据Shuffle" /></p>

<p>可以看到，对于每一个Exchange，Spark UI都提供了丰富的Metrics来刻画Shuffle的计算过程。从Shuffle Write到Shuffle Read，从数据量到处理时间，应有尽有。为了方便说明，对于Metrics的解释与释义，我以表格的方式进行了整理，供你随时查阅。</p>

<p><img src="assets/73e87a9d741a1859b287397e46abe16f.jpg" alt="图片" title="Shuffle Metrics" /></p>

<p>结合这份Shuffle的“体检报告”，我们就能以量化的方式，去掌握Shuffle过程的计算细节，从而为调优提供更多的洞察与思路。</p>

<p>为了让你获得直观感受，我还是举个例子说明。比方说，我们观察到过滤之后的中签编号数据大小不足10MB（7.4MB），这时我们首先会想到，对于这样的大表Join小表，Spark SQL选择了SortMergeJoin策略是不合理的。</p>

<p>基于这样的判断，我们完全可以让Spark SQL选择BroadcastHashJoin策略来提供更好的执行性能。至于调优的具体方法，想必不用我多说，你也早已心领神会：<strong>要么用强制广播，要么利用Spark 3.x版本提供的AQE特性</strong>。</p>

<p>你不妨结合本讲开头的代码，去完成SortMergeJoin到BroadcastHashJoin策略转换的调优，期待你在留言区分享你的调优结果。</p>

<p><strong>Sort</strong></p>

<p>接下来，我们再来说说Sort。相比Exchange，Sort的度量指标没那么多，不过，他们足以让我们一窥Sort在运行时，对于内存的消耗，如下图所示。</p>

<p><img src="assets/50e7dba6b9c1700c8b466077e8c34990.png" alt="图片" title="Sort（左）：申请编码数据排序，Sort（右）：中签编码数据排序" /></p>

<p>按照惯例，我们还是先把这些Metrics整理到表格中，方便后期查看。</p>

<p><img src="assets/3db747647eyy03b2bed8f972ff967c39.jpg" alt="图片" title="Sort Metrics" /></p>

<p>可以看到，“Peak memory total”和“Spill size total”这两个数值，足以指导我们更有针对性地去设置spark.executor.memory、spark.memory.fraction、spark.memory.storageFraction，从而使得Execution Memory区域得到充分的保障。</p>

<p>以上图为例，结合18.8GB的峰值消耗，以及12.5GB的磁盘溢出这两条信息，我们可以判断出，当前3GB的Executor Memory是远远不够的。那么我们自然要去调整上面的3个参数，来加速Sort的执行性能。</p>

<p><strong>Aggregate</strong></p>

<p>与Sort类似，衡量Aggregate的度量指标，主要记录的也是操作的内存消耗，如图所示。</p>

<p><img src="assets/cc4617577712fc0b1619bc2d67cb7fc8.png" alt="图片" title="Aggregate Metrics" /></p>

<p>可以看到，对于Aggregate操作，Spark UI也记录着磁盘溢出与峰值消耗，即Spill size和Peak memory total。这两个数值也为内存的调整提供了依据，以上图为例，零溢出与3.2GB的峰值消耗，证明当前3GB的Executor Memory设置，对于Aggregate计算来说是绰绰有余的。</p>

<p>到此为止，我们分别介绍了Exchange、Sort和Aggregate的度量指标，并结合“倍率与中签率分析”的例子，进行了简单的调优分析。</p>

<p>纵观“倍率与中签率分析”完整的DAG，我们会发现它包含了若干个Exchange、Sort、Aggregate以及Filter和Project。<strong>结合上述的各类Metrics，对于执行计划的观察与洞见，我们需要以统筹的方式，由点到线、由局部到全局地去进行</strong>。</p>

<h4 id="jobs详情页">Jobs详情页</h4>

<p>接下来，我们再来说说Jobs详情页。Jobs详情页非常的简单、直观，它罗列了隶属于当前Job的所有Stages。要想访问每一个Stage的执行细节，我们还需要通过“Description”的超链接做跳转。</p>

<p><img src="assets/9ec76b98622cff2b766dfc097d682af2.png" alt="图片" title="Jobs详情页" /></p>

<h4 id="stages详情页">Stages详情页</h4>

<p>实际上，要访问Stage详情，我们还有另外一种选择，那就是直接从Stages一级入口进入，然后完成跳转。因此，Stage详情页也归类到二级入口。接下来，我们以Id为10的Stage为例，去看一看详情页都记录着哪些关键信息。</p>

<p>在所有二级入口中，Stage详情页的信息量可以说是最大的。点进Stage详情页，可以看到它主要包含3大类信息，分别是Stage DAG、Event Timeline与Task Metrics。</p>

<p>其中，Task Metrics又分为“Summary”与“Entry details”两部分，提供不同粒度的信息汇总。而Task Metrics中记录的指标类别，还可以通过“Show Additional Metrics”选项进行扩展。</p>

<p><img src="assets/612b82f355072e03400fd162557967d9.png" alt="图片" title="Stage详情页概览" /></p>

<p><strong>Stage DAG</strong></p>

<p>接下来，我们沿着“Stage DAG -&gt; Event Timeline -&gt; Task Metrics”的顺序，依次讲讲这些页面所包含的内容。</p>

<p>首先，我们先来看最简单的Stage DAG。点开蓝色的“DAG Visualization”按钮，我们就能获取到当前Stage的DAG，如下图所示。</p>

<p><img src="assets/b4fb2fc255674897cb749b2469e32c1b.png" alt="图片" title="Stage DAG" /></p>

<p>之所以说Stage DAG简单，是因为咱们在SQL二级入口，已经对DAG做过详细的说明。而Stage DAG仅仅是SQL页面完整DAG的一个子集，毕竟，SQL页面的DAG，针对的是作业（Job）。因此，只要掌握了作业的DAG，自然也就掌握了每一个Stage的DAG。</p>

<p><strong>Event Timeline</strong></p>

<p>与“DAG Visualization”并列，在“Summary Metrics”之上，有一个“Event Timeline”按钮，点开它，我们可以得到如下图所示的可视化信息。</p>

<p><img src="assets/51d2218b6f2f25a2a15bc0385f51ee0c.png" alt="图片" title="Event Timeline" /></p>

<p>Event Timeline，记录着分布式任务调度与执行的过程中，不同计算环节主要的时间花销。图中的每一个条带，都代表着一个分布式任务，条带由不同的颜色构成。其中不同颜色的矩形，代表不同环节的计算时间。</p>

<p>为了方便叙述，我还是用表格形式帮你梳理了这些环节的含义与作用，你可以保存以后随时查看。</p>

<p><img src="assets/de11412fbf47989aeyycd7a1c86e0c40.jpg" alt="图片" title="不同环节的计算时间" /></p>

<p>理想情况下，条带的大部分应该都是绿色的（如图中所示），也就是任务的时间消耗，大部分都是执行时间。不过，实际情况并不总是如此，比如，有些时候，蓝色的部分占比较多，或是橙色的部分占比较大。</p>

<p>在这些情况下，我们就可以结合Event Timeline，来判断作业是否存在调度开销过大、或是Shuffle负载过重的问题，从而有针对性地对不同环节做调优。</p>

<p>比方说，如果条带中深蓝的部分（Scheduler Delay）很多，那就说明任务的调度开销很重。这个时候，我们就需要参考公式：D / P ~ M / C，来相应地调整CPU、内存与并行度，从而减低任务的调度开销。其中，D是数据集尺寸，P为并行度，M是Executor内存，而C是Executor的CPU核数。波浪线~表示的是，等式两边的数值，要在同一量级。</p>

<p>再比如，如果条带中黄色（Shuffle Write Time）与橙色（Shuffle Read Time）的面积较大，就说明任务的Shuffle负载很重，这个时候，我们就需要考虑，有没有可能通过利用Broadcast Join来消除Shuffle，从而缓解任务的Shuffle负担。</p>

<p><strong>Task Metrics</strong></p>

<p>说完Stage DAG与Event Timeline，最后，我们再来说一说Stage详情页的重头戏：Task Metrics。</p>

<p>之所以说它是重头戏，在于Task Metrics以不同的粒度，提供了详尽的量化指标。其中，“Tasks”以Task为粒度，记录着每一个分布式任务的执行细节，而“Summary Metrics”则是对于所有Tasks执行细节的统计汇总。我们先来看看粗粒度的“Summary Metrics”，然后再去展开细粒度的“Tasks”。</p>

<p><strong>Summary Metrics</strong></p>

<p>首先，我们点开“Show Additional Metrics”按钮，勾选“Select All”，让所有的度量指标都生效，如下图所示。这么做的目的，在于获取最详尽的Task执行信息。</p>

<p><img src="assets/bf916cabf5de22fbf16bcbda1bfb640a.png" alt="图片" title="Summary Metrics" /></p>

<p>可以看到，“Select All”生效之后，Spark UI打印出了所有的执行细节。老规矩，为了方便叙述，我还是把这些Metrics整理到表格中，方便你随时查阅。其中，Task Deserialization Time、Result Serialization Time、Getting Result Time、Scheduler Delay与刚刚表格中的含义相同，不再赘述，这里我们仅整理新出现的Task Metrics。</p>

<p><img src="assets/c6d49f5ae074078dcfa9bc28619eebd8.jpg" alt="图片" title="不同环节的计算时间" /></p>

<p>对于这些详尽的Task Metrics，难能可贵地，Spark UI以最大最小（max、min）以及分位点（25%分位、50%分位、75%分位）的方式，提供了不同Metrics的统计分布。这一点非常重要，原因在于，<strong>这些Metrics的统计分布，可以让我们非常清晰地量化任务的负载分布</strong>。</p>

<p>换句话说，根据不同Metrics的统计分布信息，我们就可以轻而易举地判定，当前作业的不同任务之间，是相对均衡，还是存在严重的倾斜。如果判定计算负载存在倾斜，那么我们就要利用AQE的自动倾斜处理，去消除任务之间的不均衡，从而改善作业性能。</p>

<p>在上面的表格中，有一半的Metrics是与Shuffle直接相关的，比如Shuffle Read Size / Records，Shuffle Remote Reads，等等。</p>

<p>这些Metrics我们在介绍SQL详情的时候，已经详细说过了。另外，Duration、GC Time、以及Peak Execution Memory，这些Metrics的含义，要么已经讲过，要么过于简单、无需解释。因此，对于这3个指标，咱们也不再多着笔墨。</p>

<p>这里特别值得你关注的，是<strong>Spill（Memory）和Spill（Disk）这两个指标</strong>。Spill，也即溢出数据，它指的是因内存数据结构（PartitionedPairBuffer、AppendOnlyMap，等等）空间受限，而腾挪出去的数据。Spill（Memory）表示的是，这部分数据在内存中的存储大小，而Spill（Disk）表示的是，这些数据在磁盘中的大小。</p>

<p>因此，用Spill（Memory）除以Spill（Disk），就可以得到“数据膨胀系数”的近似值，我们把它记为<strong>Explosion ratio</strong>。有了Explosion ratio，对于一份存储在磁盘中的数据，我们就可以估算它在内存中的存储大小，从而准确地把握数据的内存消耗。</p>

<p><strong>Tasks</strong></p>

<p>介绍完粗粒度的Summary Metrics，接下来，我们再来说说细粒度的“Tasks”。实际上，Tasks的不少指标，与Summary是高度重合的，如下图所示。同理，这些重合的Metrics，咱们不再赘述，你可以参考Summary的部分，来理解这些Metrics。唯一的区别，就是这些指标是针对每一个Task进行度量的。</p>

<p><img src="assets/c23bc53203358611328e656d64c2a43b.png" alt="图片" title="Tasks执行细节" /></p>

<p>按照惯例，咱们还是把Tasks中那些新出现的指标，整理到表格中，以备后续查看。</p>

<p><img src="assets/57182b44ca360239a1b4777458b73982.jpg" alt="图片" title="Tasks度量指标" /></p>

<p>可以看到，新指标并不多，这里最值得关注的，是<strong>Locality level</strong>，也就是本地性级别。在调度系统中，我们讲过，每个Task都有自己的本地性倾向。结合本地性倾向，调度系统会把Tasks调度到合适的Executors或是计算节点，尽可能保证“<strong>数据不动、代码动</strong>”。</p>

<p>Logs与Errors属于Spark UI的三级入口，它们是Tasks的执行日志，详细记录了Tasks在执行过程中的运行时状态。一般来说，我们不需要深入到三级入口去进行Debug。Errors列提供的报错信息，往往足以让我们迅速地定位问题所在。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>好啦，今天的课程，到这里就讲完啦。今天这一讲，我们分别学习了二级入口的SQL、Jobs与Stages。每个二级入口的内容都很丰富，提前知道它们所涵盖的信息，对我们寻找、启发与探索性能调优的思路非常有帮助。</p>

<p>到此为止，关于Spark UI的全部内容就讲完啦。Spark UI涉及的Metrics纷繁而又复杂，一次性记住确实有难度，所以通过这一讲，你只要清楚各级入口怎么找到，知道各个指标能给我们提供什么信息就好了。当然，仅仅跟着我去用“肉眼”学习一遍只是第一步，之后还需要你结合日常的开发，去多多摸索与体会，加油！</p>

<p>最后的最后，还是想提醒你，由于我们的应用是通过spark-shell提交的，因此节点8080端口的Spark UI会一直展示应用的“体检报告”。在我们退出spark-shell之后，节点8080端口的内存也随即消失（404 Page not found）。</p>

<p>要想再次查看应用的“体检报告”，需要移步至节点的18080端口，这里是Spark History Server的领地，它收集了所有（已执行完毕）应用的“体检报告”，并同样使用Spark UI的形式进行展示，切记切记。</p>

<h2 id="每课一练">每课一练</h2>

<p>今天的思考题，需要你发散思维。学习过Spark UI之后，请你说一说，都可以通过哪些途径，来定位数据倾斜问题？</p>

<p>欢迎你把Spark UI使用的心得体会，分享到课后的评论区，我们一起讨论，共同进步！也推荐你把这一讲分享更多同事、朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#87ebebebbeb3b6b6b7b0c7e0eae6eeeba9e4e8ea" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a5f8fed777755',t:'MTczNDEzOTI0Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
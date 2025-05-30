<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;Spark&#32;MLlib&#32;Pipeline：高效开发机器学习应用>
        <link rel="icon" href="/static/favicon.png">
        <title>29 Spark MLlib Pipeline：高效开发机器学习应用 </title>
        
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
                            <h1 id="title" data-id="29 Spark MLlib Pipeline：高效开发机器学习应用" class="title">29 Spark MLlib Pipeline：高效开发机器学习应用</h1>
                            <div><p>你好，我是吴磊。</p>

<p>前面我们一起学习了如何在Spark MLlib框架下做特征工程与模型训练。不论是特征工程，还是模型训练，针对同一个机器学习问题，我们往往需要尝试不同的特征处理方法或是模型算法。</p>

<p>结合之前的大量实例，细心的你想必早已发现，针对同一问题，不同的算法选型在开发的过程中，存在着大量的重复性代码。</p>

<p>以GBDT和随机森林为例，它们处理数据的过程是相似的，原始数据都是经过StringIndexer、VectorAssembler和VectorIndexer这三个环节转化为训练样本，只不过GBDT最后用GBTRegressor来做回归，而随机森林用RandomForestClassifier来做分类。</p>

<p><img src="assets/51a23f4c00c6048f262eb6006f66600b.jpg" alt="图片" title="重复的处理逻辑" /></p>

<p>不仅如此，在之前验证模型效果的时候我们也没有闭环，仅仅检查了训练集上的拟合效果，并没有在测试集上进行推理并验证。如果我们尝试去加载新的测试数据集，那么所有的特征处理过程，都需要在测试集上重演一遍。无疑，这同样会引入大量冗余的重复代码。</p>

<p>那么，有没有什么办法，能够避免上述的重复开发，让Spark MLlib框架下的机器学习开发更加高效呢？答案是肯定的，今天这一讲，我们就来说说Spark MLlib Pipeline，看看它如何帮助开发者大幅提升机器学习应用的开发效率。</p>

<h2 id="spark-mllib-pipeline">Spark MLlib Pipeline</h2>

<p>什么是Spark MLlib Pipeline呢？简单地说，Pipeline是一套基于DataFrame的高阶开发API，它让开发者以一种高效的方式，来打造端到端的机器学习流水线。这么说可能比较抽象，我们不妨先来看看，Pipeline都有哪些核心组件，它们又提供了哪些功能。</p>

<p>Pipeline的核心组件有两类，一类是Transformer，我们不妨把它称作“转换器”，另一类是Estimator，我把它叫作“模型生成器”。我们之前接触的各类特征处理函数，实际上都属于转换器，比如StringIndexer、MinMaxScaler、Bucketizer、VectorAssembler，等等。而前面3讲提到的模型算法，全部都是Estimator。</p>

<p><img src="assets/c2aa44370d88f8a73975a315c37aeb4f.jpg" alt="图片" title="Pipeline核心组件" /></p>

<h3 id="transformer">Transformer</h3>

<p>我们先来说说Transformer，数据转换器。在形式上，Transformer的输入是DataFrame，输出也是DataFrame。结合特定的数据处理逻辑，Transformer基于原有的DataFrame数据列，去创建新的数据列，而新的数据列中，往往包含着不同形式的特征。</p>

<p>以StringIndexer为例，它的转换逻辑很简单，就是把字符串转换为数值。在创建StringIndexer实例的时候，我们需要使用setInputCol(s)和setOutputCol(s)方法，来指定原始数据列和期待输出的数据列，而输出数据列中的内容就是我们需要的特征，如下图所示。</p>

<p><img src="assets/7723f13198234b75549b968d87816e82.jpg" alt="图片" title="以StringIndexer为例，演示Transformer的作用" /></p>

<p>结合图示可以看到，Transformer消费原有DataFrame的数据列，然后把生成的数据列再追加到该DataFrame，就会生成新的DataFrame。换句话说，Transformer并不是“就地”（Inline）修改原有的DataFrame，而是基于它去创建新的DataFrame。</p>

<p>实际上，每个Transformer都实现了setInputCol(s)和setOutputCol(s)这两个（接口）方法。除此之外，Transformer还提供了transform接口，用于封装具体的转换逻辑。正是基于这些核心接口，Pipeline才能把各式各样的Transformer拼接在一起，打造出了特征工程流水线。</p>

<p>一般来说，在一个机器学习应用中，我们往往需要多个Transformer来对数据做各式各样的转换，才能生成所需的训练样本。在逻辑上，多个基于同一份原始数据生成的、不同“版本”数据的DataFrame，它们会同时存在于系统中。</p>

<p>不过，受益于Spark的惰性求值（Lazy Evaluation）设计，应用在运行时并不会出现多份冗余数据重复占用内存的情况。</p>

<p>不过，为了开发上的遍历，我们还是会使用var而不是用val来命名原始的DataFrame。原因很简单，如果用val的话，我们需要反复使用新的变量名，来命名新生成的DataFrame。关于这部分开发小细节，你可以通过回顾[上一讲]的代码来体会。</p>

<h3 id="estimator">Estimator</h3>

<p>接下来，我们来说说Estimator。相比Transformer，Estimator要简单得多，它实际上就是各类模型算法，如GBDT、随机森林、线性回归，等等。Estimator的核心接口，只有一个，那就是fit，中文可以翻译成“拟合”。</p>

<p>Estimator的作用，就是定义模型算法，然后通过拟合DataFrame所囊括的训练样本，来生产模型（Models）。这也是为什么我把Estimator称作是“模型生成器”。</p>

<p>不过，有意思的是，虽然模型算法是Estimator，但是Estimator生产的模型，却是不折不扣的Transformer。</p>

<p>要搞清楚为什么模型是Transformer，我们得先弄明白模型到底是什么。所谓机器学习模型，它本质上就是一个参数（Parameters，又称权重，Weights）矩阵，外加一个模型结构。模型结构与模型算法有关，比如决策树结构、GBDT结构、神经网络结构，等等。</p>

<p>模型的核心用途就是做推断（Inference）或者说预测。给定数据样本，模型可以推断房价、推断房屋类型，等等。在Spark MLlib框架下，数据样本往往是由DataFrame封装的，而模型推断的结果，还是保存在（新的）DataFrame中，结果的默认列名是“predictions”。</p>

<p>其实基于训练好的推理逻辑，通过增加“predictions”列，把一个DataFrame转化成一个新的DataFrame，这不就是Transformer在做的事情吗？而这，也是为什么在模型算法上，我们调用的是fit方法，而在做模型推断时，我们在模型上调用的是transform方法。</p>

<h2 id="构建pipeline">构建Pipeline</h2>

<p>好啦，了解了Transformer和Estimator之后，我们就可以基于它们去构建Pipeline，来打造端到端的机器学习流水线。实际上，一旦Transformer、Estimator准备就绪，定义Pipeline只需一行代码就可以轻松拿下，如下所示。</p>

<pre><code>import org.apache.spark.ml.Pipeline

// 像之前一样，定义各种特征处理对象与模型算法
val stringIndexer = _
val vectorAssembler = _
val vectorIndexer = _
val gbtRegressor = _

// 将所有的Transformer、Estimator依序放入数组
val stages = Array(stringIndexer, vectorAssembler, vectorIndexer, gbtRegressor)

// 定义Spark MLlib Pipeline
val newPipeline = new Pipeline()
.setStages(stages)
</code></pre>

<p>可以看到，要定义Pipeline，只需创建Pipeline实例，然后把之前定义好的Transformer、Estimator纷纷作为参数，传入setStages方法即可。需要注意的是，<strong>一个Pipeline可以包含多个Transformer和Estimator，不过，Pipeline的最后一个环节，必须是Estimator，切记</strong>。</p>

<p>到此为止，Pipeline的作用、定义以及核心组件，我们就讲完了。不过，你可能会说：“概念是讲完了，不过我还是不知道Pipeline具体怎么用，以及它到底有什么优势？”别着急，光说不练假把式，接下来，我们就结合GBDT与随机森林的例子，来说说Pipeline的具体用法，以及怎么用它帮你大幅度提升开发效率。</p>

<p>首先，我们来看看，在一个机器学习应用中，Pipeline如何帮助我们提高效率。在上一讲，我们用GBDT来拟合房价，并给出了代码示例。</p>

<p>现在，咱们把代码稍微调整一下，用Spark MLlib Pipeline来实现模型训练。第一步，我们还是先从文件创建DataFrame，然后把数值型字段与非数值型字段区分开，如下所示。</p>

<pre><code>import org.apache.spark.sql.DataFrame

// rootPath为房价预测数据集根目录
val rootPath: String = _
val filePath: String = s&quot;${rootPath}/train.csv&quot;

// 读取文件，创建DataFrame
var engineeringDF: DataFrame = spark.read.format(&quot;csv&quot;).option(&quot;header&quot;, true).load(filePath)

// 所有数值型字段
val numericFields: Array[String] = Array(&quot;LotFrontage&quot;, &quot;LotArea&quot;, &quot;MasVnrArea&quot;, &quot;BsmtFinSF1&quot;, &quot;BsmtFinSF2&quot;, &quot;BsmtUnfSF&quot;, &quot;TotalBsmtSF&quot;, &quot;1stFlrSF&quot;, &quot;2ndFlrSF&quot;, &quot;LowQualFinSF&quot;, &quot;GrLivArea&quot;, &quot;BsmtFullBath&quot;, &quot;BsmtHalfBath&quot;, &quot;FullBath&quot;, &quot;HalfBath&quot;, &quot;BedroomAbvGr&quot;, &quot;KitchenAbvGr&quot;, &quot;TotRmsAbvGrd&quot;, &quot;Fireplaces&quot;, &quot;GarageCars&quot;, &quot;GarageArea&quot;, &quot;WoodDeckSF&quot;, &quot;OpenPorchSF&quot;, &quot;EnclosedPorch&quot;, &quot;3SsnPorch&quot;, &quot;ScreenPorch&quot;, &quot;PoolArea&quot;)

// Label字段
val labelFields: Array[String] = Array(&quot;SalePrice&quot;)

import org.apache.spark.sql.types.IntegerType

for (field &lt;- (numericFields ++ labelFields)) {
engineeringDF = engineeringDF
.withColumn(s&quot;${field}Int&quot;,col(field).cast(IntegerType))
.drop(field)
}
</code></pre>

<p>数据准备好之后，接下来，我们就可以开始着手，为Pipeline的构建打造零件：依次定义转换器Transformer和模型生成器Estimator。在上一讲，我们用StringIndexer把非数值字段转换为数值字段，这一讲，咱们也依法炮制。</p>

<pre><code>import org.apache.spark.ml.feature.StringIndexer

// 所有非数值型字段
val categoricalFields: Array[String] = Array(&quot;MSSubClass&quot;, &quot;MSZoning&quot;, &quot;Street&quot;, &quot;Alley&quot;, &quot;LotShape&quot;, &quot;LandContour&quot;, &quot;Utilities&quot;, &quot;LotConfig&quot;, &quot;LandSlope&quot;, &quot;Neighborhood&quot;, &quot;Condition1&quot;, &quot;Condition2&quot;, &quot;BldgType&quot;, &quot;HouseStyle&quot;, &quot;OverallQual&quot;, &quot;OverallCond&quot;, &quot;YearBuilt&quot;, &quot;YearRemodAdd&quot;, &quot;RoofStyle&quot;, &quot;RoofMatl&quot;, &quot;Exterior1st&quot;, &quot;Exterior2nd&quot;, &quot;MasVnrType&quot;, &quot;ExterQual&quot;, &quot;ExterCond&quot;, &quot;Foundation&quot;, &quot;BsmtQual&quot;, &quot;BsmtCond&quot;, &quot;BsmtExposure&quot;, &quot;BsmtFinType1&quot;, &quot;BsmtFinType2&quot;, &quot;Heating&quot;, &quot;HeatingQC&quot;, &quot;CentralAir&quot;, &quot;Electrical&quot;, &quot;KitchenQual&quot;, &quot;Functional&quot;, &quot;FireplaceQu&quot;, &quot;GarageType&quot;, &quot;GarageYrBlt&quot;, &quot;GarageFinish&quot;, &quot;GarageQual&quot;, &quot;GarageCond&quot;, &quot;PavedDrive&quot;, &quot;PoolQC&quot;, &quot;Fence&quot;, &quot;MiscFeature&quot;, &quot;MiscVal&quot;, &quot;MoSold&quot;, &quot;YrSold&quot;, &quot;SaleType&quot;, &quot;SaleCondition&quot;)

// StringIndexer期望的输出列名
val indexFields: Array[String] = categoricalFields.map(_ + &quot;Index&quot;).toArray

// 定义StringIndexer实例
val stringIndexer = new StringIndexer()
// 批量指定输入列名
.setInputCols(categoricalFields)
// 批量指定输出列名，输出列名与输入列名，必须要一一对应
.setOutputCols(indexFields)
.setHandleInvalid(&quot;keep&quot;) 
</code></pre>

<p>在上一讲，定义完StringIndexer实例之后，我们立即拿它去对engineeringDF做转换。不过在构建Pipeline的时候，我们不需要这么做，只需要把这个“零件”定义好即可。接下来，我们来打造下一个零件：VectorAssembler。</p>

<pre><code>import org.apache.spark.ml.feature.VectorAssembler

// 转换为整型的数值型字段
val numericFeatures: Array[String] = numericFields.map(_ + &quot;Int&quot;).toArray

val vectorAssembler = new VectorAssembler()
/** 输入列为：数值型字段 + 非数值型字段
注意，非数值型字段的列名，要用indexFields，
而不能用原始的categoricalFields，不妨想一想为什么？
*/
.setInputCols(numericFeatures ++ indexFields)
.setOutputCol(&quot;features&quot;)
.setHandleInvalid(&quot;keep&quot;)
</code></pre>

<p>与上一讲相比，VectorAssembler的定义并没有什么两样。</p>

<p>下面，我们继续来打造第三个零件：VectorIndexer，它用于帮助模型算法区分连续特征与离散特征。</p>

<pre><code>import org.apache.spark.ml.feature.VectorIndexer

val vectorIndexer = new VectorIndexer()
// 指定输入列
.setInputCol(&quot;features&quot;)
// 指定输出列
.setOutputCol(&quot;indexedFeatures&quot;)
// 指定连续、离散判定阈值
.setMaxCategories(30)
.setHandleInvalid(&quot;keep&quot;)
</code></pre>

<p>到此为止，Transformer就全部定义完了，原始数据经过StringIndexer、VectorAssembler和VectorIndexer的转换之后，会生成新的DataFrame。在这个最新的DataFrame中，会有多个由不同Transformer生成的数据列，其中“indexedFeatures”列包含的数据内容即是特征向量。</p>

<p>结合DataFrame一路携带过来的“SalePriceInt”列，特征向量与预测标的终于结合在一起了，就是我们常说的训练样本。有了训练样本，接下来，我们就可以着手定义Estimator。</p>

<pre><code>import org.apache.spark.ml.regression.GBTRegressor

val gbtRegressor = new GBTRegressor()
// 指定预测标的
.setLabelCol(&quot;SalePriceInt&quot;)
// 指定特征向量
.setFeaturesCol(&quot;indexedFeatures&quot;)
// 指定决策树的数量
.setMaxIter(30)
// 指定决策树的最大深度
.setMaxDepth(5)
</code></pre>

<p>好啦，到这里，Pipeline所需的零件全部打造完毕，零件就位，只欠组装。我们需要通过Spark MLlib提供的“流水线工艺”，把所有零件组装成Pipeline。</p>

<pre><code>import org.apache.spark.ml.Pipeline

val components = Array(stringIndexer, vectorAssembler, vectorIndexer, gbtRegressor)

val pipeline = new Pipeline()
.setStages(components)
</code></pre>

<p>怎么样，是不是很简单？接下来的问题是，有了Pipeline，我们都能用它做些什么呢？</p>

<pre><code>// Pipeline保存地址的根目录
val savePath: String = _

// 将Pipeline物化到磁盘，以备后用（复用）
pipeline.write
.overwrite()
.save(s&quot;${savePath}/unfit-gbdt-pipeline&quot;)

// 划分出训练集和验证集
val Array(trainingData, validationData) = engineeringDF.randomSplit(Array(0.7, 0.3))

// 调用fit方法，触发Pipeline计算，并最终拟合出模型
val pipelineModel = pipeline.fit(trainingData)
</code></pre>

<p>首先，我们可以把Pipeline保存下来，以备后用，至于怎么复用，我们待会再说。再者，把之前准备好的训练样本，传递给Pipeline的fit方法，即可触发整条Pipeline从头至尾的计算逻辑，从各式各样的数据转换，到最终的模型训练，一步到位。-
Pipeline fit方法的输出结果，即是训练好的机器学习模型。我们最开始说过，模型也是Transformer，它可以用来推断预测结果。</p>

<p>看到这里，你可能会说：“和之前的代码实现相比，Pipeline也没有什么特别之处，无非是用Pipeline API把之前的环节拼接起来而已”。其实不然，基于构建好的Pipeline，我们可以在不同范围对其进行复用。对于机器学习应用来说，我们<strong>既可以在作业内部实现复用，也可以在作业之间实现复用，从而大幅度提升开发效率</strong>。</p>

<h3 id="作业内的代码复用">作业内的代码复用</h3>

<p>在之前的模型训练过程中，我们仅仅在训练集与验证集上评估了模型效果。实际上，在工业级应用中，我们最关心的，是模型在测试集上的泛化能力。就拿Kaggle竞赛来说，对于每一个机器学习项目，Kaggle都会同时提供train.csv和test.csv两个文件。</p>

<p>其中train.csv是带标签的，用于训练模型，而test.csv是不带标签的。我们需要对test.csv中的数据做推断，然后把预测结果提交到Kaggle线上平台，平台会结合房屋的实际价格来评判我们的模型，到那时我们才能知道，模型对于房价的预测到底有多准（或是有多不准）。</p>

<p>要完成对test.csv的推断，我们需要把原始数据转换为特征向量，也就是把“粗粮”转化为“细粮”，然后才能把它“喂给”模型。</p>

<p>在之前的代码实现中，要做到这一点，我们必须把之前加持到train.csv的所有转换逻辑都重写一遍，比如StringIndexer、VectorAssembler和VectorIndexer。毫无疑问，这样的开发方式是极其低效的，更糟的是，手工重写很容易会造成测试样本与训练样本不一致，而这样的不一致是机器学习应用中的大忌。</p>

<p>不过，有了Pipeline，我们就可以省去这些麻烦。首先，我们把test.csv加载进来并创建DataFrame，然后把数值字段从String转为Int。</p>

<pre><code>import org.apache.spark.sql.DataFrame

val rootPath: String = _
val filePath: String = s&quot;${rootPath}/test.csv&quot;

// 加载test.csv，并创建DataFrame
var testData: DataFrame = spark.read.format(&quot;csv&quot;).option(&quot;header&quot;, true).load(filePath)

// 所有数值型字段
val numericFields: Array[String] = Array(&quot;LotFrontage&quot;, &quot;LotArea&quot;, &quot;MasVnrArea&quot;, &quot;BsmtFinSF1&quot;, &quot;BsmtFinSF2&quot;, &quot;BsmtUnfSF&quot;, &quot;TotalBsmtSF&quot;, &quot;1stFlrSF&quot;, &quot;2ndFlrSF&quot;, &quot;LowQualFinSF&quot;, &quot;GrLivArea&quot;, &quot;BsmtFullBath&quot;, &quot;BsmtHalfBath&quot;, &quot;FullBath&quot;, &quot;HalfBath&quot;, &quot;BedroomAbvGr&quot;, &quot;KitchenAbvGr&quot;, &quot;TotRmsAbvGrd&quot;, &quot;Fireplaces&quot;, &quot;GarageCars&quot;, &quot;GarageArea&quot;, &quot;WoodDeckSF&quot;, &quot;OpenPorchSF&quot;, &quot;EnclosedPorch&quot;, &quot;3SsnPorch&quot;, &quot;ScreenPorch&quot;, &quot;PoolArea&quot;)

// 所有非数值型字段
val categoricalFields: Array[String] = Array(&quot;MSSubClass&quot;, &quot;MSZoning&quot;, &quot;Street&quot;, &quot;Alley&quot;, &quot;LotShape&quot;, &quot;LandContour&quot;, &quot;Utilities&quot;, &quot;LotConfig&quot;, &quot;LandSlope&quot;, &quot;Neighborhood&quot;, &quot;Condition1&quot;, &quot;Condition2&quot;, &quot;BldgType&quot;, &quot;HouseStyle&quot;, &quot;OverallQual&quot;, &quot;OverallCond&quot;, &quot;YearBuilt&quot;, &quot;YearRemodAdd&quot;, &quot;RoofStyle&quot;, &quot;RoofMatl&quot;, &quot;Exterior1st&quot;, &quot;Exterior2nd&quot;, &quot;MasVnrType&quot;, &quot;ExterQual&quot;, &quot;ExterCond&quot;, &quot;Foundation&quot;, &quot;BsmtQual&quot;, &quot;BsmtCond&quot;, &quot;BsmtExposure&quot;, &quot;BsmtFinType1&quot;, &quot;BsmtFinType2&quot;, &quot;Heating&quot;, &quot;HeatingQC&quot;, &quot;CentralAir&quot;, &quot;Electrical&quot;, &quot;KitchenQual&quot;, &quot;Functional&quot;, &quot;FireplaceQu&quot;, &quot;GarageType&quot;, &quot;GarageYrBlt&quot;, &quot;GarageFinish&quot;, &quot;GarageQual&quot;, &quot;GarageCond&quot;, &quot;PavedDrive&quot;, &quot;PoolQC&quot;, &quot;Fence&quot;, &quot;MiscFeature&quot;, &quot;MiscVal&quot;, &quot;MoSold&quot;, &quot;YrSold&quot;, &quot;SaleType&quot;, &quot;SaleCondition&quot;)

import org.apache.spark.sql.types.IntegerType

// 注意，test.csv没有SalePrice字段，也即没有Label
for (field &lt;- (numericFields)) {
testData = testData
.withColumn(s&quot;${field}Int&quot;,col(field).cast(IntegerType))
.drop(field)
}
</code></pre>

<p>接下来，我们只需要调用Pipeline Model的transform方法，就可以对测试集做推理。还记得吗？模型是Transformer，而transform是Transformer用于数据转换的统一接口。</p>

<pre><code>val predictions = pipelineModel.transform(testData)
</code></pre>

<p>有了Pipeline，我们就可以省去StringIndexer、VectorAssembler这些特征处理函数的重复定义，在提升开发效率的同时，消除样本不一致的隐患。除了在同一个作业内部复用Pipeline之外，我们还可以在不同的作业之间对其进行复用，从而进一步提升开发效率。</p>

<h3 id="作业间的代码复用">作业间的代码复用</h3>

<p>对于同一个机器学习问题，我们往往会尝试不同的模型算法，以期获得更好的模型效果。例如，对于房价预测，我们既可以用GBDT，也可以用随机森林。不过，尽管模型算法不同，但是它们的训练样本往往是类似的，甚至是完全一样的。如果每尝试一种模型算法，就需要从头处理一遍数据，这未免过于低效，也容易出错。</p>

<p>有了Pipeline，我们就可以把算法选型这件事变得异常简单。还是拿房价预测来举例，之前我们尝试使用GBTRegressor来训练模型，这一次，咱们来试试RandomForestRegressor，也即使用随机森林来解决回归问题。按照惯例，我们还是结合代码来进行讲解。</p>

<pre><code>import org.apache.spark.ml.Pipeline

val savePath: String = _

// 加载之前保存到磁盘的Pipeline
val unfitPipeline = Pipeline.load(s&quot;${savePath}/unfit-gbdt-pipeline&quot;)

// 获取Pipeline中的每一个Stage（Transformer或Estimator）
val formerStages = unfitPipeline.getStages

// 去掉Pipeline中最后一个组件，也即Estimator：GBTRegressor
val formerStagesWithoutModel = formerStages.dropRight(1)

import org.apache.spark.ml.regression.RandomForestRegressor

// 定义新的Estimator：RandomForestRegressor
val rf = new RandomForestRegressor()
.setLabelCol(&quot;SalePriceInt&quot;)
.setFeaturesCol(&quot;indexedFeatures&quot;)
.setNumTrees(30)
.setMaxDepth(5)

// 将老的Stages与新的Estimator拼接在一起
val stages = formerStagesWithoutModel ++ Array(rf)

// 重新定义新的Pipeline
val newPipeline = new Pipeline()
.setStages(stages)
</code></pre>

<p>首先，我们把之前保存下来的Pipeline，重新加载进来。然后，用新的RandomForestRegressor替换原来的GBTRegressor。最后，再把原有的Stages和新的Estimator拼接在一起，去创建新的Pipeline即可。接下来，只要调用fit方法，就可以触发新Pipeline的运转，并最终拟合出新的随机森林模型。</p>

<pre><code>// 像之前一样，从train.csv创建DataFrame，准备数据
var engineeringDF = _

val Array(trainingData, testData) = engineeringDF.randomSplit(Array(0.7, 0.3))

// 调用fit方法，触发Pipeline运转，拟合新模型
val pipelineModel = newPipeline.fit(trainingData)
</code></pre>

<p>可以看到，短短的几行代码，就可以让我们轻松地完成模型选型。到此，Pipeline在开发效率与容错上的优势，可谓一览无余。</p>

<h2 id="重点回顾">重点回顾</h2>

<p>今天的内容就讲完啦，今天这一讲，我们一起学习了Spark MLlib Pipeline。你需要理解Pipeline的优势所在，并掌握它的核心组件与具体用法。Pipeline的核心组件是Transformer与Estimator。</p>

<p>其中，Transformer完成从DataFrame到DataFrame的转换，基于固有的转换逻辑，生成新的数据列。Estimator主要是模型算法，它基于DataFrame中封装的训练样本，去生成机器学习模型。将若干Transformer与Estimator拼接在一起，通过调用Pipeline的setStages方法，即可完成Pipeline的创建。</p>

<p><strong>Pipeline的核心优势在于提升机器学习应用的开发效率，并同时消除测试样本与训练样本之间不一致这一致命隐患。Pipeline可用于作业内的代码复用，或是作业间的代码复用。</strong></p>

<p>在同一作业内，Pipeline能够轻松地在测试集之上，完成数据推断。而在作业之间，开发者可以加载之前保存好的Pipeline，然后用“新零件”替换“旧零件”的方式，在复用大部分处理逻辑的同时，去打造新的Pipeline，从而实现高效的模型选型过程。</p>

<p>在今后的机器学习开发中，我们要充分利用Pipeline提供的优势，来降低开发成本，从而把主要精力放在特征工程与模型调优上。</p>

<p>到此为止，Spark MLlib模块的全部内容，我们就讲完了。</p>

<p>在这个模块中，我们主要围绕着特征工程、模型训练和机器学习流水线等几个方面，梳理了Spark MLlib子框架为开发者提供的种种能力。换句话说，我们知道了Spark MLlib能做哪些事情、擅长做哪些事情。如果我们能够做到对这些能力了如指掌，在日常的机器学习开发中，就可以灵活地对其进行取舍，从而去应对不断变化的业务需求，加油！</p>

<h2 id="每日一练">每日一练</h2>

<p>我们今天一直在讲Pipeline的优势，你能说一说，Pipeline有哪些可能的劣势吗？</p>

<p>欢迎你在留言区和我交流互动，也推荐你把这一讲分享给更多同事、朋友，说不定就能让他进一步理解Pipeline。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#fa969696c3cecbcbcacdba9d979b9396d4999597" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a6230ae827755',t:'MTczNDEzOTM1My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=18&#32;如何自己开发一个大数据SQL引擎？>
        <link rel="icon" href="/static/favicon.png">
        <title>18 如何自己开发一个大数据SQL引擎？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 为什么说每个软件工程师都应该懂大数据技术？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4%e6%af%8f%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e5%b8%88%e9%83%bd%e5%ba%94%e8%af%a5%e6%87%82%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%8a%80%e6%9c%af%ef%bc%9f.md">00 开篇词 为什么说每个软件工程师都应该懂大数据技术？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 大数据技术发展史：大数据的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/01%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%8a%80%e6%9c%af%e5%8f%91%e5%b1%95%e5%8f%b2%ef%bc%9a%e5%a4%a7%e6%95%b0%e6%8d%ae%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">01 大数据技术发展史：大数据的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 大数据应用发展史：从搜索引擎到人工智能.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/02%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%ba%94%e7%94%a8%e5%8f%91%e5%b1%95%e5%8f%b2%ef%bc%9a%e4%bb%8e%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%e5%88%b0%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd.md">02 大数据应用发展史：从搜索引擎到人工智能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 大数据应用领域：数据驱动一切.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/03%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%ba%94%e7%94%a8%e9%a2%86%e5%9f%9f%ef%bc%9a%e6%95%b0%e6%8d%ae%e9%a9%b1%e5%8a%a8%e4%b8%80%e5%88%87.md">03 大数据应用领域：数据驱动一切.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 移动计算比移动数据更划算.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/04%20%e7%a7%bb%e5%8a%a8%e8%ae%a1%e7%ae%97%e6%af%94%e7%a7%bb%e5%8a%a8%e6%95%b0%e6%8d%ae%e6%9b%b4%e5%88%92%e7%ae%97.md">04 移动计算比移动数据更划算.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 从RAID看垂直伸缩到水平伸缩的演化.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/05%20%e4%bb%8eRAID%e7%9c%8b%e5%9e%82%e7%9b%b4%e4%bc%b8%e7%bc%a9%e5%88%b0%e6%b0%b4%e5%b9%b3%e4%bc%b8%e7%bc%a9%e7%9a%84%e6%bc%94%e5%8c%96.md">05 从RAID看垂直伸缩到水平伸缩的演化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 新技术层出不穷，HDFS依然是存储的王者.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/06%20%e6%96%b0%e6%8a%80%e6%9c%af%e5%b1%82%e5%87%ba%e4%b8%8d%e7%a9%b7%ef%bc%8cHDFS%e4%be%9d%e7%84%b6%e6%98%af%e5%ad%98%e5%82%a8%e7%9a%84%e7%8e%8b%e8%80%85.md">06 新技术层出不穷，HDFS依然是存储的王者.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 为什么说MapReduce既是编程模型又是计算框架？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/07%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4MapReduce%e6%97%a2%e6%98%af%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%9e%8b%e5%8f%88%e6%98%af%e8%ae%a1%e7%ae%97%e6%a1%86%e6%9e%b6%ef%bc%9f.md">07 为什么说MapReduce既是编程模型又是计算框架？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 MapReduce如何让数据完成一次旅行？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/08%20MapReduce%e5%a6%82%e4%bd%95%e8%ae%a9%e6%95%b0%e6%8d%ae%e5%ae%8c%e6%88%90%e4%b8%80%e6%ac%a1%e6%97%85%e8%a1%8c%ef%bc%9f.md">08 MapReduce如何让数据完成一次旅行？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 为什么我们管Yarn叫作资源调度框架？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/09%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e4%bb%ac%e7%ae%a1Yarn%e5%8f%ab%e4%bd%9c%e8%b5%84%e6%ba%90%e8%b0%83%e5%ba%a6%e6%a1%86%e6%9e%b6%ef%bc%9f.md">09 为什么我们管Yarn叫作资源调度框架？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 模块答疑：我们能从Hadoop学到什么？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/10%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e6%88%91%e4%bb%ac%e8%83%bd%e4%bb%8eHadoop%e5%ad%a6%e5%88%b0%e4%bb%80%e4%b9%88%ef%bc%9f.md">10 模块答疑：我们能从Hadoop学到什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Hive是如何让MapReduce实现SQL操作的？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/11%20Hive%e6%98%af%e5%a6%82%e4%bd%95%e8%ae%a9MapReduce%e5%ae%9e%e7%8e%b0SQL%e6%93%8d%e4%bd%9c%e7%9a%84%ef%bc%9f.md">11 Hive是如何让MapReduce实现SQL操作的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 我们并没有觉得MapReduce速度慢，直到Spark出现.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/12%20%e6%88%91%e4%bb%ac%e5%b9%b6%e6%b2%a1%e6%9c%89%e8%a7%89%e5%be%97MapReduce%e9%80%9f%e5%ba%a6%e6%85%a2%ef%bc%8c%e7%9b%b4%e5%88%b0Spark%e5%87%ba%e7%8e%b0.md">12 我们并没有觉得MapReduce速度慢，直到Spark出现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 同样的本质，为何Spark可以更高效？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/13%20%e5%90%8c%e6%a0%b7%e7%9a%84%e6%9c%ac%e8%b4%a8%ef%bc%8c%e4%b8%ba%e4%bd%95Spark%e5%8f%af%e4%bb%a5%e6%9b%b4%e9%ab%98%e6%95%88%ef%bc%9f.md">13 同样的本质，为何Spark可以更高效？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 BigTable的开源实现：HBase.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/14%20BigTable%e7%9a%84%e5%bc%80%e6%ba%90%e5%ae%9e%e7%8e%b0%ef%bc%9aHBase.md">14 BigTable的开源实现：HBase.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 流式计算的代表：Storm、Flink、Spark Streaming.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/15%20%e6%b5%81%e5%bc%8f%e8%ae%a1%e7%ae%97%e7%9a%84%e4%bb%a3%e8%a1%a8%ef%bc%9aStorm%e3%80%81Flink%e3%80%81Spark%20Streaming.md">15 流式计算的代表：Storm、Flink、Spark Streaming.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 ZooKeeper是如何保证数据一致性的？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/16%20ZooKeeper%e6%98%af%e5%a6%82%e4%bd%95%e4%bf%9d%e8%af%81%e6%95%b0%e6%8d%ae%e4%b8%80%e8%87%b4%e6%80%a7%e7%9a%84%ef%bc%9f.md">16 ZooKeeper是如何保证数据一致性的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 模块答疑：这么多技术，到底都能用在什么场景里？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/17%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e8%bf%99%e4%b9%88%e5%a4%9a%e6%8a%80%e6%9c%af%ef%bc%8c%e5%88%b0%e5%ba%95%e9%83%bd%e8%83%bd%e7%94%a8%e5%9c%a8%e4%bb%80%e4%b9%88%e5%9c%ba%e6%99%af%e9%87%8c%ef%bc%9f.md">17 模块答疑：这么多技术，到底都能用在什么场景里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 如何自己开发一个大数据SQL引擎？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/18%20%e5%a6%82%e4%bd%95%e8%87%aa%e5%b7%b1%e5%bc%80%e5%8f%91%e4%b8%80%e4%b8%aa%e5%a4%a7%e6%95%b0%e6%8d%aeSQL%e5%bc%95%e6%93%8e%ef%bc%9f.md">18 如何自己开发一个大数据SQL引擎？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Spark的性能优化案例分析（上）.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/19%20Spark%e7%9a%84%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89.md">19 Spark的性能优化案例分析（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Spark的性能优化案例分析（下）.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/20%20Spark%e7%9a%84%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89.md">20 Spark的性能优化案例分析（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 从阿里内部产品看海量数据处理系统的设计（上）：Doris的立项.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/21%20%e4%bb%8e%e9%98%bf%e9%87%8c%e5%86%85%e9%83%a8%e4%ba%a7%e5%93%81%e7%9c%8b%e6%b5%b7%e9%87%8f%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e7%b3%bb%e7%bb%9f%e7%9a%84%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aDoris%e7%9a%84%e7%ab%8b%e9%a1%b9.md">21 从阿里内部产品看海量数据处理系统的设计（上）：Doris的立项.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 从阿里内部产品看海量数据处理系统的设计（下）：架构与创新.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/22%20%e4%bb%8e%e9%98%bf%e9%87%8c%e5%86%85%e9%83%a8%e4%ba%a7%e5%93%81%e7%9c%8b%e6%b5%b7%e9%87%8f%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e7%b3%bb%e7%bb%9f%e7%9a%84%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%9e%b6%e6%9e%84%e4%b8%8e%e5%88%9b%e6%96%b0.md">22 从阿里内部产品看海量数据处理系统的设计（下）：架构与创新.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 大数据基准测试可以带来什么好处？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/23%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%9f%ba%e5%87%86%e6%b5%8b%e8%af%95%e5%8f%af%e4%bb%a5%e5%b8%a6%e6%9d%a5%e4%bb%80%e4%b9%88%e5%a5%bd%e5%a4%84%ef%bc%9f.md">23 大数据基准测试可以带来什么好处？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 从大数据性能测试工具Dew看如何快速开发大数据系统.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/24%20%e4%bb%8e%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7Dew%e7%9c%8b%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%bc%80%e5%8f%91%e5%a4%a7%e6%95%b0%e6%8d%ae%e7%b3%bb%e7%bb%9f.md">24 从大数据性能测试工具Dew看如何快速开发大数据系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 模块答疑：我能从大厂的大数据开发实践中学到什么？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/25%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e6%88%91%e8%83%bd%e4%bb%8e%e5%a4%a7%e5%8e%82%e7%9a%84%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e8%b7%b5%e4%b8%ad%e5%ad%a6%e5%88%b0%e4%bb%80%e4%b9%88%ef%bc%9f.md">25 模块答疑：我能从大厂的大数据开发实践中学到什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 互联网产品 &#43; 大数据产品 = 大数据平台.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/26%20%e4%ba%92%e8%81%94%e7%bd%91%e4%ba%a7%e5%93%81%20&#43;%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e4%ba%a7%e5%93%81%20=%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0.md">26 互联网产品 &#43; 大数据产品 = 大数据平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 大数据从哪里来？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/27%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e4%bb%8e%e5%93%aa%e9%87%8c%e6%9d%a5%ef%bc%9f.md">27 大数据从哪里来？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 知名大厂如何搭建大数据平台？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/28%20%e7%9f%a5%e5%90%8d%e5%a4%a7%e5%8e%82%e5%a6%82%e4%bd%95%e6%90%ad%e5%bb%ba%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0%ef%bc%9f.md">28 知名大厂如何搭建大数据平台？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 盘点可供中小企业参考的商业大数据平台.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/29%20%e7%9b%98%e7%82%b9%e5%8f%af%e4%be%9b%e4%b8%ad%e5%b0%8f%e4%bc%81%e4%b8%9a%e5%8f%82%e8%80%83%e7%9a%84%e5%95%86%e4%b8%9a%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0.md">29 盘点可供中小企业参考的商业大数据平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 当大数据遇上物联网.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/30%20%e5%bd%93%e5%a4%a7%e6%95%b0%e6%8d%ae%e9%81%87%e4%b8%8a%e7%89%a9%e8%81%94%e7%bd%91.md">30 当大数据遇上物联网.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 模块答疑：为什么大数据平台至关重要？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/31%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e5%a4%a7%e6%95%b0%e6%8d%ae%e5%b9%b3%e5%8f%b0%e8%87%b3%e5%85%b3%e9%87%8d%e8%a6%81%ef%bc%9f.md">31 模块答疑：为什么大数据平台至关重要？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 互联网运营数据指标与可视化监控.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/32%20%e4%ba%92%e8%81%94%e7%bd%91%e8%bf%90%e8%90%a5%e6%95%b0%e6%8d%ae%e6%8c%87%e6%a0%87%e4%b8%8e%e5%8f%af%e8%a7%86%e5%8c%96%e7%9b%91%e6%8e%a7.md">32 互联网运营数据指标与可视化监控.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 一个电商网站订单下降的数据分析案例.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/33%20%e4%b8%80%e4%b8%aa%e7%94%b5%e5%95%86%e7%bd%91%e7%ab%99%e8%ae%a2%e5%8d%95%e4%b8%8b%e9%99%8d%e7%9a%84%e6%95%b0%e6%8d%ae%e5%88%86%e6%9e%90%e6%a1%88%e4%be%8b.md">33 一个电商网站订单下降的数据分析案例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 A_B测试与灰度发布必知必会.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/34%20A_B%e6%b5%8b%e8%af%95%e4%b8%8e%e7%81%b0%e5%ba%a6%e5%8f%91%e5%b8%83%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a.md">34 A_B测试与灰度发布必知必会.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 如何利用大数据成为“增长黑客”？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/35%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%88%90%e4%b8%ba%e2%80%9c%e5%a2%9e%e9%95%bf%e9%bb%91%e5%ae%a2%e2%80%9d%ef%bc%9f.md">35 如何利用大数据成为“增长黑客”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 模块答疑：为什么说数据驱动运营？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/36%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4%e6%95%b0%e6%8d%ae%e9%a9%b1%e5%8a%a8%e8%bf%90%e8%90%a5%ef%bc%9f.md">36 模块答疑：为什么说数据驱动运营？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 如何对数据进行分类和预测？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/37%20%e5%a6%82%e4%bd%95%e5%af%b9%e6%95%b0%e6%8d%ae%e8%bf%9b%e8%a1%8c%e5%88%86%e7%b1%bb%e5%92%8c%e9%a2%84%e6%b5%8b%ef%bc%9f.md">37 如何对数据进行分类和预测？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 如何发掘数据之间的关系？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/38%20%e5%a6%82%e4%bd%95%e5%8f%91%e6%8e%98%e6%95%b0%e6%8d%ae%e4%b9%8b%e9%97%b4%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">38 如何发掘数据之间的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 如何预测用户的喜好？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/39%20%e5%a6%82%e4%bd%95%e9%a2%84%e6%b5%8b%e7%94%a8%e6%88%b7%e7%9a%84%e5%96%9c%e5%a5%bd%ef%bc%9f.md">39 如何预测用户的喜好？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 机器学习的数学原理是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/40%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%95%b0%e5%ad%a6%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">40 机器学习的数学原理是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 从感知机到神经网络算法.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/41%20%e4%bb%8e%e6%84%9f%e7%9f%a5%e6%9c%ba%e5%88%b0%e7%a5%9e%e7%bb%8f%e7%bd%91%e7%bb%9c%e7%ae%97%e6%b3%95.md">41 从感知机到神经网络算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 模块答疑：软件工程师如何进入人工智能领域？.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/42%20%e6%a8%a1%e5%9d%97%e7%ad%94%e7%96%91%ef%bc%9a%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e5%b8%88%e5%a6%82%e4%bd%95%e8%bf%9b%e5%85%a5%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e9%a2%86%e5%9f%9f%ef%bc%9f.md">42 模块答疑：软件工程师如何进入人工智能领域？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="所有的不确定都是机会——智慧写给你的新年寄语.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/%e6%89%80%e6%9c%89%e7%9a%84%e4%b8%8d%e7%a1%ae%e5%ae%9a%e9%83%bd%e6%98%af%e6%9c%ba%e4%bc%9a%e2%80%94%e2%80%94%e6%99%ba%e6%85%a7%e5%86%99%e7%bb%99%e4%bd%a0%e7%9a%84%e6%96%b0%e5%b9%b4%e5%af%84%e8%af%ad.md">所有的不确定都是机会——智慧写给你的新年寄语.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第2季回归丨大数据之后，让我们回归后端.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/%e7%ac%ac2%e5%ad%a3%e5%9b%9e%e5%bd%92%e4%b8%a8%e5%a4%a7%e6%95%b0%e6%8d%ae%e4%b9%8b%e5%90%8e%ef%bc%8c%e8%ae%a9%e6%88%91%e4%bb%ac%e5%9b%9e%e5%bd%92%e5%90%8e%e7%ab%af.md">第2季回归丨大数据之后，让我们回归后端.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 未来的你，有无限可能.md" href="/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%a4%a7%e6%95%b0%e6%8d%ae/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%9c%aa%e6%9d%a5%e7%9a%84%e4%bd%a0%ef%bc%8c%e6%9c%89%e6%97%a0%e9%99%90%e5%8f%af%e8%83%bd.md">结束语 未来的你，有无限可能.md</a>
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
                            <h1 id="title" data-id="18 如何自己开发一个大数据SQL引擎？" class="title">18 如何自己开发一个大数据SQL引擎？</h1>
                            <div><p>从今天开始我们就进入了专栏的第三个模块，一起来看看大数据开发实践过程中的门道。学习一样技术，如果只是作为学习者，被动接受总是困难的。<strong>但如果从开发者的视角看，很多东西就豁然开朗了，明白了原理，有时甚至不需要学习，顺着原理就可以推导出各种实现细节</strong>。</p>

<p>各种知识从表象上看，总是杂乱无章的，如果只是学习这些繁杂的知识点，固然自己的知识面是有限的，并且遇到问题的应变能力也很难提高。所以有些高手看起来似乎无所不知，不论谈论起什么技术，都能头头是道，其实并不是他们学习、掌握了所有技术，而是他们是在谈到这个问题的时候，才开始进行推导，并迅速得出结论。</p>

<p>我在Intel的时候，面试过一个交大的实习生，她大概只学过一点MapReduce的基本知识，我问她如何用MapReduce实现数据库的join操作，可以明显看出她没学习过这部分知识。她说：我想一下，然后盯着桌子看了两三秒的时间，就开始回答，基本跟Hive的实现机制一样。从她的回答就能看出这个女生就是一个高手，高手不一定要很资深、经验丰富，把握住了技术的核心本质，掌握了快速分析推导的能力，能够迅速将自己的知识技能推进到陌生的领域，就是高手。</p>

<p>这也是我这个专栏的目的，讲述大数据技术的核心原理，分享一些高效的思考和思维方式，帮助你构建起自己的技术知识体系。</p>

<p>在这个模块里，我将以大数据开发者的视角，讲述大数据开发需要关注的各种问题，以及相应的解决方案。希望你可以进入我的角色，跳出纷繁复杂的知识表象，掌握核心原理和思维方式，进而融会贯通各种技术，再通过各种实践训练，最终成为真正的高手。</p>

<p>前面专栏我们提到过，程序员的三大梦想，也就是开发数据库、操作系统和编译器。今天我们通过一个支持标准SQL语法的大数据仓库引擎的设计开发案例，看看如何自己开发一个大数据SQL引擎。</p>

<p>在学习今天的内容前，我们先回顾一下前面讨论过的大数据仓库Hive。作为一个成功的大数据仓库，它将SQL语句转换成MapReduce执行过程，并把大数据应用的门槛下降到普通数据分析师和工程师就可以很快上手的地步。这部分内容如果你忘记了，可以返回<a href="http://time.geekbang.org/column/article/69459" target="_blank">专栏第11期</a>复习一下。</p>

<p>但是Hive也有自己的问题，由于它使用自己定义的Hive QL语法，这对已经熟悉Oracle等传统数据仓库的分析师来说，还是有一定的上手难度。特别是很多企业使用传统数据仓库进行数据分析已经由来已久，沉淀了大量的SQL语句，并且这些SQL经过多年的修改打磨，非常庞大也非常复杂。我曾经见过某银行的一条统计报表SQL，打印出来足足两张A4纸。这样的SQL光是完全理解可能就要花很长时间，再转化成Hive QL就更加费力，还不说转化修改过程中可能引入的bug。</p>

<p>2012年那会，我还在Intel亚太研发中心大数据团队，当时团队决定开发一款能够支持标准数据库SQL的大数据仓库引擎，希望让那些在Oracle上运行良好的SQL可以直接运行在Hadoop上，而不需要重写成Hive QL。这就是后来的<a href="https://github.com/intel-hadoop/project-panthera-ase" target="_blank">Panthera</a>项目。</p>

<p>在开发Panthera前，我们分析一下Hive的主要处理过程，大体上分成三步：</p>

<p>1.将输入的Hive QL经过语法解析器转换成Hive抽象语法树（Hive AST）。-
2.将Hive AST经过语义分析器转换成MapReduce执行计划。-
3.将生成的MapReduce执行计划和Hive执行函数代码提交到Hadoop上执行。</p>

<p>Panthera的设计思路是保留Hive语义分析器不动，替换Hive语法解析器，使其将标准SQL语句转换成Hive语义分析器能够处理的Hive抽象语法树。用图形来表示的话，是用红框内的部分代替黑框内原来Hive的部分。</p>

<p><img src="assets/ea8ecd0af0eb4459819225070673a6d3.jpg" alt="" /></p>

<p>红框内的组件我们重新开发过，浅蓝色的是我们使用的一个开源的SQL语法解析器，将标准SQL解析成标准SQL抽象语法树（SQL AST），后面深蓝色的就是团队自己开发的SQL抽象语法树分析与转换器，将SQL AST转换成Hive AST。</p>

<p>那么标准SQL和Hive QL的差别在哪里呢？</p>

<p>标准SQL和Hive QL的差别主要有两个方面，一个是语法表达方式，Hive QL语法和标准SQL语法略有不同；另一个是Hive QL支持的语法元素比标准SQL要少很多，比如，数据仓库领域主要的测试集<a href="http://www.tpc.org/tpch/" target="_blank">TPC-H</a>所有的SQL语句Hive都不支持。尤其是Hive不支持复杂的嵌套子查询，而对于数据仓库分析而言，嵌套子查询几乎是无处不在的。比如下面这样的SQL，在where查询条件existes里面包含了另一条SQL语句。</p>

<pre><code>select o_orderpriority, count(*) as order_count 
from orders 
where o_orderdate &gt;= date '[DATE]' 
and o_orderdate &lt; date '[DATE]' + interval '3' month 
and exists 
( select * from lineitem 
where l_orderkey = o_orderkey and l_commitdate &lt; l_receiptdate ) 
group by o_orderpriority order by o_orderpriority;
</code></pre>

<p>所以开发支持标准SQL语法的SQL引擎的难点，就变成<strong>如何将复杂的嵌套子查询消除掉</strong>，也就是where条件里不包含select。</p>

<p>SQL的理论基础是关系代数，而关系代数的主要操作只有5种，分别是并、差、积、选择、投影。所有的SQL语句最后都能用这5种操作组合完成。而一个嵌套子查询可以等价转换成一个连接（join）操作。</p>

<p>比如这条SQL</p>

<pre><code>select s_grade from staff where s_city not in (select p_city from proj where s_empname=p_pname)
</code></pre>

<p>这是一个在where条件里嵌套了not in子查询的SQL语句，它可以用left outer join和left semi join进行等价转换，示例如下，这是Panthera自动转换完成得到的等价SQL。这条SQL语句不再包含嵌套子查询，</p>

<pre><code>select panthera_10.panthera_1 as s_grade from (select panthera_1, panthera_4, panthera_6, s_empname, s_city from (select s_grade as panthera_1, s_city as panthera_4, s_empname as panthera_6, s_empname as s_empname, s_city as s_city from staff) panthera_14 left outer join (select panthera_16.panthera_7 as panthera_7, panthera_16.panthera_8 as panthera_8, panthera_16.panthera_9 as panthera_9, panthera_16.panthera_12 as panthera_12, panthera_16.panthera_13 as panthera_13 from (select panthera_0.panthera_1 as panthera_7, panthera_0.panthera_4 as panthera_8, panthera_0.panthera_6 as panthera_9, panthera_0.s_empname as panthera_12, panthera_0.s_city as panthera_13 from (select s_grade as panthera_1, s_city as panthera_4, s_empname as panthera_6, s_empname, s_city from staff) panthera_0 left semi join (select p_city as panthera_3, p_pname as panthera_5 from proj) panthera_2 on (panthera_0.panthera_4 = panthera_2.panthera_3) and (panthera_0.panthera_6 = panthera_2.panthera_5) where true) panthera_16 group by panthera_16.panthera_7, panthera_16.panthera_8, panthera_16.panthera_9, panthera_16.panthera_12, panthera_16.panthera_13) panthera_15 on ((((panthera_14.panthera_1 &lt;=&gt; panthera_15.panthera_7) and (panthera_14.panthera_4 &lt;=&gt; panthera_15.panthera_8)) and (panthera_14.panthera_6 &lt;=&gt; panthera_15.panthera_9)) and (panthera_14.s_empname &lt;=&gt; panthera_15.panthera_12)) and (panthera_14.s_city &lt;=&gt; panthera_15.panthera_13) where ((((panthera_15.panthera_7 is null) and (panthera_15.panthera_8 is null)) and (panthera_15.panthera_9 is null)) and (panthera_15.panthera_12 is null)) and (panthera_15.panthera_13 is null)) panthera_10 ;
</code></pre>

<p>通过可视化工具将上面两条SQL的语法树展示出来，是这样的。</p>

<p><img src="assets/8902ee3b43c24250aee0e2fa9a9f187b.jpg" alt="" /></p>

<p>这是原始的SQL抽象语法树。</p>

<p><img src="assets/466b680e69f74a60b3c32e083f75b977.jpg" alt="" /></p>

<p>这是等价转换后的抽象语法树，内容太多被压缩得无法看清，不过你可以感受一下（笑）。</p>

<p>那么，在程序设计上如何实现这样复杂的语法转换呢？当时Panthera项目组合使用了几种经典的设计模式，每个语法点被封装到一个类里去处理，每个类通常不过几十行代码，这样整个程序非常简单、清爽。如果在测试过程中遇到不支持的语法点，只需为这个语法点新增加一个类即可，团队协作与代码维护非常容易。</p>

<p>使用装饰模式的语法等价转换类的构造，Panthera每增加一种新的语法转换能力，只需要开发一个新的Transformer类，然后添加到下面的构造函数代码里即可。</p>

<pre><code>   private static SqlASTTransformer tf =
      new RedundantSelectGroupItemTransformer(
      new DistinctTransformer(
      new GroupElementNormalizeTransformer(
      new PrepareQueryInfoTransformer(
      new OrderByTransformer(
      new OrderByFunctionTransformer(
      new MinusIntersectTransformer(
      new PrepareQueryInfoTransformer(
      new UnionTransformer(
      new Leftsemi2LeftJoinTransformer(
      new CountAsteriskPositionTransformer(
      new FilterInwardTransformer(
      //use leftJoin method to handle not exists for correlated
      new CrossJoinTransformer(
      new PrepareQueryInfoTransformer(
      new SubQUnnestTransformer(
      new PrepareFilterBlockTransformer(
      new PrepareQueryInfoTransformer(
      new TopLevelUnionTransformer(
      new FilterBlockAdjustTransformer(
      new PrepareFilterBlockTransformer(
      new ExpandAsteriskTransformer(
      new PrepareQueryInfoTransformer(
      new CrossJoinTransformer(
      new PrepareQueryInfoTransformer(
      new ConditionStructTransformer(
      new MultipleTableSelectTransformer(
      new WhereConditionOptimizationTransformer(
      new PrepareQueryInfoTransformer(
      new InTransformer(
      new TopLevelUnionTransformer(
      new MinusIntersectTransformer(
      new NaturalJoinTransformer(
      new OrderByNotInSelectListTransformer(
      new RowNumTransformer(
      new BetweenTransformer(
      new UsingTransformer(
      new SchemaDotTableTransformer(
      new NothingTransformer())))))))))))))))))))))))))))))))))))));
</code></pre>

<p>而在具体的Transformer类中，则使用组合模式对抽象语法树AST进行遍历，以下为Between语法节点的遍历。我们看到使用组合模式进行树的遍历不需要用递归算法，因为递归的特性已经隐藏在树的结构里面了。</p>

<pre><code>  @Override
  protected void transform(CommonTree tree, TranslateContext context) throws SqlXlateException {
    tf.transformAST(tree, context);
    trans(tree, context);
  }

  void trans(CommonTree tree, TranslateContext context) {
    // deep firstly
    for (int i = 0; i &lt; tree.getChildCount(); i++) {
      trans((CommonTree) (tree.getChild(i)), context);
    }
    if (tree.getType() == PantheraExpParser.SQL92_RESERVED_BETWEEN) {
      transBetween(false, tree, context);
    }
    if (tree.getType() == PantheraExpParser.NOT_BETWEEN) {
      transBetween(true, tree, context);
    }
  }
</code></pre>

<p>将等价转换后的抽象语法树AST再进一步转换成Hive格式的抽象语法树，就可以交给Hive的语义分析器去处理了，从而也就实现了对标准SQL的支持。</p>

<p>当时Facebook为了证明Hive对数据仓库的支持，Facebook的工程师手工将TPC-H的测试SQL转换成Hive QL，我们将这些手工Hive QL和Panthera进行对比测试，两者性能各有所长，总体上不相上下，这说明Panthera自动进行语法分析和转换的效率还是不错的。</p>

<p>Panthera（ASE）和Facebook手工Hive QL对比测试结果如下。</p>

<p><img src="assets/1ee0e4203db04dfd93dc6cc832dee803.jpg" alt="" /></p>

<p>事实上，标准SQL语法集的语法点非常多，我们经过近两年的努力，绞尽脑汁进行各种关系代数等价变形，依然没有全部适配所有的标准SQL语法。</p>

<p>我在开发Panthera的时候，查阅了很多关于SQL和数据库的网站和文档，发现在我们耳熟能详的那些主流数据库之外还有很多名不见经传的数据库，此外，还有大量的关于SQL的语法解析器、测试数据和脚本集、各种周边工具。我们经常看到的MySQL、Oracle这些产品仅仅是整个数据库生态体系的冰山一角。还有很多优秀的数据库在竞争中落了下风、默默无闻，而更多的支撑起这些优秀数据库的论文、工具，非业内人士几乎闻所未闻。</p>

<p>这个认识给了我很大触动，我一直期待我们中国人能开发出自己的操作系统、数据库、编程语言，也看到有很多人前仆后继投入其中。但是这么多年过去了，大部分努力惨淡收场，小部分结果沦落为笑柄，成为人们饭后的谈资。我曾经思考过，为什么会这样？</p>

<p>开发Panthera之后，我想，我们国家虽然从事软件开发的人数很多，但是绝大多数都在做最顶层的应用开发，底层技术开发和研究的人太少，做出来的成果也太少。我们在一个缺乏周边生态体系、没有足够的竞争产品的情况下，就想直接开发出自己有影响力的操作系统、数据库、编程语言，无异于想在沙漠里种出几棵参天大树。</p>

<p>不过，庆幸的是，我也看到越来越多有全球影响力的底层技术产品中出现中国人的身影，小草已经在默默生长，假以时日，料想必有大树出现。</p>

<p>今天我讲的是一个SQL引擎是如何设计出来的，也许在你的工作几乎不可能去开发SQL引擎，但是了解这些基础的知识，了解一些设计的技巧，对你用好数据库，开发更加灵活、有弹性的系统也会很有帮助。</p>

<h2 id="思考题">思考题</h2>

<p>SQL注入是一种常见的Web攻击手段，如下图所示，攻击者在HTTP请求中注入恶意SQL命令（drop table users;），服务器用请求参数构造数据库SQL命令时，恶意SQL被一起构造，并在数据库中执行。</p>

<p><img src="assets/b256b6d952144caa912509e2fa151d6b.jpg" alt="" /></p>

<p>但是JDBC的PrepareStatement可以阻止SQL注入攻击，MyBatis之类的ORM框架也可以阻止SQL注入，请从数据库引擎的工作机制解释PrepareStatement和MyBatis的防注入攻击的原理。</p>

<p>欢迎你点击“请朋友读”，把今天的文章分享给好友。也欢迎你写下自己的思考或疑问，与我和其他同学一起讨论。</p>
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
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1339b179ce76e1',t:'MTczNDA2NDI5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;Spark&#32;Streaming：Spark的实时流计算API>
        <link rel="icon" href="/static/favicon.png">
        <title>16 Spark Streaming：Spark的实时流计算API </title>
        
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
                        <a class="menu-item" id="00 开篇词 从这里开始，带你走上硅谷一线系统架构师之路.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%e8%bf%99%e9%87%8c%e5%bc%80%e5%a7%8b%ef%bc%8c%e5%b8%a6%e4%bd%a0%e8%b5%b0%e4%b8%8a%e7%a1%85%e8%b0%b7%e4%b8%80%e7%ba%bf%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84%e5%b8%88%e4%b9%8b%e8%b7%af.md">00 开篇词 从这里开始，带你走上硅谷一线系统架构师之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 为什么MapReduce会被硅谷一线公司淘汰？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/01%20%e4%b8%ba%e4%bb%80%e4%b9%88MapReduce%e4%bc%9a%e8%a2%ab%e7%a1%85%e8%b0%b7%e4%b8%80%e7%ba%bf%e5%85%ac%e5%8f%b8%e6%b7%98%e6%b1%b0%ef%bc%9f.md">01 为什么MapReduce会被硅谷一线公司淘汰？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 MapReduce后谁主沉浮：怎样设计下一代数据处理技术？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/02%20MapReduce%e5%90%8e%e8%b0%81%e4%b8%bb%e6%b2%89%e6%b5%ae%ef%bc%9a%e6%80%8e%e6%a0%b7%e8%ae%be%e8%ae%a1%e4%b8%8b%e4%b8%80%e4%bb%a3%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e6%8a%80%e6%9c%af%ef%bc%9f.md">02 MapReduce后谁主沉浮：怎样设计下一代数据处理技术？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 大规模数据处理初体验：怎样实现大型电商热销榜？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/03%20%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%88%9d%e4%bd%93%e9%aa%8c%ef%bc%9a%e6%80%8e%e6%a0%b7%e5%ae%9e%e7%8e%b0%e5%a4%a7%e5%9e%8b%e7%94%b5%e5%95%86%e7%83%ad%e9%94%80%e6%a6%9c%ef%bc%9f.md">03 大规模数据处理初体验：怎样实现大型电商热销榜？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 分布式系统（上）：学会用服务等级协议SLA来评估你的系统.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/04%20%e5%88%86%e5%b8%83%e5%bc%8f%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%ad%a6%e4%bc%9a%e7%94%a8%e6%9c%8d%e5%8a%a1%e7%ad%89%e7%ba%a7%e5%8d%8f%e8%ae%aeSLA%e6%9d%a5%e8%af%84%e4%bc%b0%e4%bd%a0%e7%9a%84%e7%b3%bb%e7%bb%9f.md">04 分布式系统（上）：学会用服务等级协议SLA来评估你的系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 分布式系统（下）：架构师不得不知的三大指标.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/05%20%e5%88%86%e5%b8%83%e5%bc%8f%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%9e%b6%e6%9e%84%e5%b8%88%e4%b8%8d%e5%be%97%e4%b8%8d%e7%9f%a5%e7%9a%84%e4%b8%89%e5%a4%a7%e6%8c%87%e6%a0%87.md">05 分布式系统（下）：架构师不得不知的三大指标.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 如何区分批处理还是流处理？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/06%20%e5%a6%82%e4%bd%95%e5%8c%ba%e5%88%86%e6%89%b9%e5%a4%84%e7%90%86%e8%bf%98%e6%98%af%e6%b5%81%e5%a4%84%e7%90%86%ef%bc%9f.md">06 如何区分批处理还是流处理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Workflow设计模式：让你在大规模数据世界中君临天下.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/07%20Workflow%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%9a%e8%ae%a9%e4%bd%a0%e5%9c%a8%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e4%b8%96%e7%95%8c%e4%b8%ad%e5%90%9b%e4%b8%b4%e5%a4%a9%e4%b8%8b.md">07 Workflow设计模式：让你在大规模数据世界中君临天下.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 发布_订阅模式：流处理架构中的瑞士军刀.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/08%20%e5%8f%91%e5%b8%83_%e8%ae%a2%e9%98%85%e6%a8%a1%e5%bc%8f%ef%bc%9a%e6%b5%81%e5%a4%84%e7%90%86%e6%9e%b6%e6%9e%84%e4%b8%ad%e7%9a%84%e7%91%9e%e5%a3%ab%e5%86%9b%e5%88%80.md">08 发布_订阅模式：流处理架构中的瑞士军刀.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 CAP定理：三选二，架构师必须学会的取舍.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/09%20CAP%e5%ae%9a%e7%90%86%ef%bc%9a%e4%b8%89%e9%80%89%e4%ba%8c%ef%bc%8c%e6%9e%b6%e6%9e%84%e5%b8%88%e5%bf%85%e9%a1%bb%e5%ad%a6%e4%bc%9a%e7%9a%84%e5%8f%96%e8%88%8d.md">09 CAP定理：三选二，架构师必须学会的取舍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Lambda架构：Twitter亿级实时数据分析架构背后的倚天剑.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/10%20Lambda%e6%9e%b6%e6%9e%84%ef%bc%9aTwitter%e4%ba%bf%e7%ba%a7%e5%ae%9e%e6%97%b6%e6%95%b0%e6%8d%ae%e5%88%86%e6%9e%90%e6%9e%b6%e6%9e%84%e8%83%8c%e5%90%8e%e7%9a%84%e5%80%9a%e5%a4%a9%e5%89%91.md">10 Lambda架构：Twitter亿级实时数据分析架构背后的倚天剑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Kappa架构：利用Kafka锻造的屠龙刀.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/11%20Kappa%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%88%a9%e7%94%a8Kafka%e9%94%bb%e9%80%a0%e7%9a%84%e5%b1%a0%e9%be%99%e5%88%80.md">11 Kappa架构：利用Kafka锻造的屠龙刀.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 我们为什么需要Spark？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/12%20%e6%88%91%e4%bb%ac%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81Spark%ef%bc%9f.md">12 我们为什么需要Spark？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 弹性分布式数据集：Spark大厦的地基（上）.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/13%20%e5%bc%b9%e6%80%a7%e5%88%86%e5%b8%83%e5%bc%8f%e6%95%b0%e6%8d%ae%e9%9b%86%ef%bc%9aSpark%e5%a4%a7%e5%8e%a6%e7%9a%84%e5%9c%b0%e5%9f%ba%ef%bc%88%e4%b8%8a%ef%bc%89.md">13 弹性分布式数据集：Spark大厦的地基（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 弹性分布式数据集：Spark大厦的地基（下）.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/14%20%e5%bc%b9%e6%80%a7%e5%88%86%e5%b8%83%e5%bc%8f%e6%95%b0%e6%8d%ae%e9%9b%86%ef%bc%9aSpark%e5%a4%a7%e5%8e%a6%e7%9a%84%e5%9c%b0%e5%9f%ba%ef%bc%88%e4%b8%8b%ef%bc%89.md">14 弹性分布式数据集：Spark大厦的地基（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Spark SQL：Spark数据查询的利器.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/15%20Spark%20SQL%ef%bc%9aSpark%e6%95%b0%e6%8d%ae%e6%9f%a5%e8%af%a2%e7%9a%84%e5%88%a9%e5%99%a8.md">15 Spark SQL：Spark数据查询的利器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Spark Streaming：Spark的实时流计算API.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/16%20Spark%20Streaming%ef%bc%9aSpark%e7%9a%84%e5%ae%9e%e6%97%b6%e6%b5%81%e8%ae%a1%e7%ae%97API.md">16 Spark Streaming：Spark的实时流计算API.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 Structured Streaming：如何用DataFrame API进行实时数据分析_.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/17%20Structured%20Streaming%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8DataFrame%20API%e8%bf%9b%e8%a1%8c%e5%ae%9e%e6%97%b6%e6%95%b0%e6%8d%ae%e5%88%86%e6%9e%90_.md">17 Structured Streaming：如何用DataFrame API进行实时数据分析_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Word Count：从零开始运行你的第一个Spark应用.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/18%20Word%20Count%ef%bc%9a%e4%bb%8e%e9%9b%b6%e5%bc%80%e5%a7%8b%e8%bf%90%e8%a1%8c%e4%bd%a0%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aaSpark%e5%ba%94%e7%94%a8.md">18 Word Count：从零开始运行你的第一个Spark应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 综合案例实战：处理加州房屋信息，构建线性回归模型.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/19%20%e7%bb%bc%e5%90%88%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e5%a4%84%e7%90%86%e5%8a%a0%e5%b7%9e%e6%88%bf%e5%b1%8b%e4%bf%a1%e6%81%af%ef%bc%8c%e6%9e%84%e5%bb%ba%e7%ba%bf%e6%80%a7%e5%9b%9e%e5%bd%92%e6%a8%a1%e5%9e%8b.md">19 综合案例实战：处理加州房屋信息，构建线性回归模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 流处理案例实战：分析纽约市出租车载客信息.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/20%20%e6%b5%81%e5%a4%84%e7%90%86%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%ef%bc%9a%e5%88%86%e6%9e%90%e7%ba%bd%e7%ba%a6%e5%b8%82%e5%87%ba%e7%a7%9f%e8%bd%a6%e8%bd%bd%e5%ae%a2%e4%bf%a1%e6%81%af.md">20 流处理案例实战：分析纽约市出租车载客信息.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 深入对比Spark与Flink：帮你系统设计两开花.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/21%20%e6%b7%b1%e5%85%a5%e5%af%b9%e6%af%94Spark%e4%b8%8eFlink%ef%bc%9a%e5%b8%ae%e4%bd%a0%e7%b3%bb%e7%bb%9f%e8%ae%be%e8%ae%a1%e4%b8%a4%e5%bc%80%e8%8a%b1.md">21 深入对比Spark与Flink：帮你系统设计两开花.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 Apache Beam的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/22%20Apache%20Beam%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">22 Apache Beam的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 站在Google的肩膀上学习Beam编程模型.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/23%20%e7%ab%99%e5%9c%a8Google%e7%9a%84%e8%82%a9%e8%86%80%e4%b8%8a%e5%ad%a6%e4%b9%a0Beam%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%9e%8b.md">23 站在Google的肩膀上学习Beam编程模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 PCollection：为什么Beam要如此抽象封装数据？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/24%20PCollection%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88Beam%e8%a6%81%e5%a6%82%e6%ad%a4%e6%8a%bd%e8%b1%a1%e5%b0%81%e8%a3%85%e6%95%b0%e6%8d%ae%ef%bc%9f.md">24 PCollection：为什么Beam要如此抽象封装数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 Transform：Beam数据转换操作的抽象方法.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/25%20Transform%ef%bc%9aBeam%e6%95%b0%e6%8d%ae%e8%bd%ac%e6%8d%a2%e6%93%8d%e4%bd%9c%e7%9a%84%e6%8a%bd%e8%b1%a1%e6%96%b9%e6%b3%95.md">25 Transform：Beam数据转换操作的抽象方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 Pipeline：Beam如何抽象多步骤的数据流水线？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/26%20Pipeline%ef%bc%9aBeam%e5%a6%82%e4%bd%95%e6%8a%bd%e8%b1%a1%e5%a4%9a%e6%ad%a5%e9%aa%a4%e7%9a%84%e6%95%b0%e6%8d%ae%e6%b5%81%e6%b0%b4%e7%ba%bf%ef%bc%9f.md">26 Pipeline：Beam如何抽象多步骤的数据流水线？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 Pipeline I_O_ Beam数据中转的设计模式.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/27%20Pipeline%20I_O_%20Beam%e6%95%b0%e6%8d%ae%e4%b8%ad%e8%bd%ac%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f.md">27 Pipeline I_O_ Beam数据中转的设计模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 如何设计创建好一个Beam Pipeline？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/28%20%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e5%88%9b%e5%bb%ba%e5%a5%bd%e4%b8%80%e4%b8%aaBeam%20Pipeline%ef%bc%9f.md">28 如何设计创建好一个Beam Pipeline？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 如何测试Beam Pipeline？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/29%20%e5%a6%82%e4%bd%95%e6%b5%8b%e8%af%95Beam%20Pipeline%ef%bc%9f.md">29 如何测试Beam Pipeline？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 Apache Beam实战冲刺：Beam如何run everywhere_.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/30%20Apache%20Beam%e5%ae%9e%e6%88%98%e5%86%b2%e5%88%ba%ef%bc%9aBeam%e5%a6%82%e4%bd%95run%20everywhere_.md">30 Apache Beam实战冲刺：Beam如何run everywhere_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 WordCount Beam Pipeline实战.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/31%20WordCount%20Beam%20Pipeline%e5%ae%9e%e6%88%98.md">31 WordCount Beam Pipeline实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 Beam Window：打通流处理的任督二脉.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/32%20Beam%20Window%ef%bc%9a%e6%89%93%e9%80%9a%e6%b5%81%e5%a4%84%e7%90%86%e7%9a%84%e4%bb%bb%e7%9d%a3%e4%ba%8c%e8%84%89.md">32 Beam Window：打通流处理的任督二脉.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 横看成岭侧成峰：再战Streaming WordCount.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/33%20%e6%a8%aa%e7%9c%8b%e6%88%90%e5%b2%ad%e4%be%a7%e6%88%90%e5%b3%b0%ef%bc%9a%e5%86%8d%e6%88%98Streaming%20WordCount.md">33 横看成岭侧成峰：再战Streaming WordCount.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 Amazon热销榜Beam Pipeline实战.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/34%20Amazon%e7%83%ad%e9%94%80%e6%a6%9cBeam%20Pipeline%e5%ae%9e%e6%88%98.md">34 Amazon热销榜Beam Pipeline实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 Facebook游戏实时流处理Beam Pipeline实战（上）.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/35%20Facebook%e6%b8%b8%e6%88%8f%e5%ae%9e%e6%97%b6%e6%b5%81%e5%a4%84%e7%90%86Beam%20Pipeline%e5%ae%9e%e6%88%98%ef%bc%88%e4%b8%8a%ef%bc%89.md">35 Facebook游戏实时流处理Beam Pipeline实战（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Facebook游戏实时流处理Beam Pipeline实战（下）.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/36%20Facebook%e6%b8%b8%e6%88%8f%e5%ae%9e%e6%97%b6%e6%b5%81%e5%a4%84%e7%90%86Beam%20Pipeline%e5%ae%9e%e6%88%98%ef%bc%88%e4%b8%8b%ef%bc%89.md">36 Facebook游戏实时流处理Beam Pipeline实战（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 5G时代，如何处理超大规模物联网数据.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/37%205G%e6%97%b6%e4%bb%a3%ef%bc%8c%e5%a6%82%e4%bd%95%e5%a4%84%e7%90%86%e8%b6%85%e5%a4%a7%e8%a7%84%e6%a8%a1%e7%89%a9%e8%81%94%e7%bd%91%e6%95%b0%e6%8d%ae.md">37 5G时代，如何处理超大规模物联网数据.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 大规模数据处理在深度学习中如何应用？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/38%20%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%9c%a8%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e4%b8%ad%e5%a6%82%e4%bd%95%e5%ba%94%e7%94%a8%ef%bc%9f.md">38 大规模数据处理在深度学习中如何应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 从SQL到Streaming SQL：突破静态数据查询的次元.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/39%20%e4%bb%8eSQL%e5%88%b0Streaming%20SQL%ef%bc%9a%e7%aa%81%e7%a0%b4%e9%9d%99%e6%80%81%e6%95%b0%e6%8d%ae%e6%9f%a5%e8%af%a2%e7%9a%84%e6%ac%a1%e5%85%83.md">39 从SQL到Streaming SQL：突破静态数据查询的次元.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 大规模数据处理未来之路.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/40%20%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e6%9c%aa%e6%9d%a5%e4%b9%8b%e8%b7%af.md">40 大规模数据处理未来之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="FAQ第一期 学习大规模数据处理需要什么基础？.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/FAQ%e7%ac%ac%e4%b8%80%e6%9c%9f%20%e5%ad%a6%e4%b9%a0%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e9%9c%80%e8%a6%81%e4%bb%80%e4%b9%88%e5%9f%ba%e7%a1%80%ef%bc%9f.md">FAQ第一期 学习大规模数据处理需要什么基础？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="FAQ第三期 Apache Beam基础答疑.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/FAQ%e7%ac%ac%e4%b8%89%e6%9c%9f%20Apache%20Beam%e5%9f%ba%e7%a1%80%e7%ad%94%e7%96%91.md">FAQ第三期 Apache Beam基础答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="FAQ第二期 Spark案例实战答疑.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/FAQ%e7%ac%ac%e4%ba%8c%e6%9c%9f%20Spark%e6%a1%88%e4%be%8b%e5%ae%9e%e6%88%98%e7%ad%94%e7%96%91.md">FAQ第二期 Spark案例实战答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加油站 Practice makes perfect！.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/%e5%8a%a0%e6%b2%b9%e7%ab%99%20Practice%20makes%20perfect%ef%bc%81.md">加油站 Practice makes perfect！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 世间所有的相遇，都是久别重逢.md" href="/%e4%b8%93%e6%a0%8f/%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e4%b8%96%e9%97%b4%e6%89%80%e6%9c%89%e7%9a%84%e7%9b%b8%e9%81%87%ef%bc%8c%e9%83%bd%e6%98%af%e4%b9%85%e5%88%ab%e9%87%8d%e9%80%a2.md">结束语 世间所有的相遇，都是久别重逢.md</a>
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
                            <h1 id="title" data-id="16 Spark Streaming：Spark的实时流计算API" class="title">16 Spark Streaming：Spark的实时流计算API</h1>
                            <div><p>你好，我是蔡元楠。</p>

<p>今天我要与你分享的内容是“Spark Streaming”。</p>

<p>通过上一讲的内容，我们深入了解了Spark SQL API。通过它，我们可以像查询关系型数据库一样查询Spark的数据，并且对原生数据做相应的转换和动作。</p>

<p>但是，无论是DataFrame API还是DataSet API，都是基于批处理模式对静态数据进行处理的。比如，在每天某个特定的时间对一天的日志进行处理分析。</p>

<p>在第二章中你已经知道了，批处理和流处理是大数据处理最常见的两个场景。那么作为当下最流行的大数据处理平台之一，Spark是否支持流处理呢？</p>

<p>答案是肯定的。</p>

<p>早在2013年，Spark的流处理组件Spark Streaming就发布了。之后经过好几年的迭代与改进，现在的Spark Streaming已经非常成熟，在业界应用十分广泛。</p>

<p>今天就让我们一起揭开Spark Streaming的神秘面纱，让它成为我们手中的利器。</p>

<h2 id="spark-streaming的原理">Spark Streaming的原理</h2>

<p>Spark Streaming的原理与微积分的思想很类似。</p>

<p>在大学的微积分课上，你的老师一定说过，微分就是无限细分，积分就是对无限细分的每一段进行求和。它本质上把一个连续的问题转换成了无限个离散的问题。</p>

<p>比如，用微积分思想求下图中阴影部分S的面积。</p>

<p><img src="assets/a881c974a2b041e986574626c78bb1ae.jpg" alt="" /></p>

<p>我们可以把S无限细分成无数个小矩形，因为矩形的宽足够短，所以它顶端的边近似是一个直线。这样，把容易计算的矩形面积相加，就得到不容易直接计算的不规则图形面积。</p>

<p>你知道，流处理的数据是一系列连续不断变化，且无边界的。我们永远无法预测下一秒的数据是什么样。Spark Streaming用时间片拆分了无限的数据流，然后对每一个数据片用类似于批处理的方法进行处理，输出的数据也是一块一块的。如下图所示。</p>

<p><img src="assets/08c5409ab51d40c19cbaba785e799243.jpg" alt="" /></p>

<p>Spark Streaming提供一个对于流数据的抽象DStream。DStream可以由来自Apache Kafka、Flume或者HDFS的流数据生成，也可以由别的DStream经过各种转换操作得来。讲到这里，你是不是觉得内容似曾相识？</p>

<p>没错，底层DStream也是由很多个序列化的RDD构成，按时间片（比如一秒）切分成的每个数据单位都是一个RDD。然后，Spark核心引擎将对DStream的Transformation操作变为针对Spark中对 RDD的Transformation操作，将RDD经过操作变成中间结果保存在内存中。</p>

<p>之前的DataFrame和DataSet也是同样基于RDD，所以说RDD是Spark最基本的数据抽象。就像Java里的基本数据类型（Primitive Type）一样，所有的数据都可以用基本数据类型描述。</p>

<p>也正是因为这样，无论是DataFrame，还是DStream，都具有RDD的不可变性、分区性和容错性等特质。</p>

<p>所以，Spark是一个高度统一的平台，所有的高级API都有相同的性质，它们之间可以很容易地相互转化。Spark的野心就是用这一套工具统一所有数据处理的场景。</p>

<p>由于Spark Streaming将底层的细节封装起来了，所以对于开发者来说，只需要操作DStream就行。接下来，让我们一起学习DStream的结构以及它支持的转换操作。</p>

<h2 id="dstream">DStream</h2>

<p>下图就是DStream的内部形式，即一个连续的RDD序列，每一个RDD代表一个时间窗口的输入数据流。</p>

<p><img src="assets/9b168c38101b4c9fa6946c2d9a135823.jpg" alt="" /></p>

<p>对DStream的转换操作，意味着对它包含的每一个RDD进行同样的转换操作。比如下边的例子。</p>

<pre><code>sc = SparkContext(master, appName)
ssc = StreamingContext(sc, 1)
lines = sc.socketTextStream(&quot;localhost&quot;, 9999)
words = lines.flatMap(lambda line: line.split(&quot; &quot;))
</code></pre>

<p><img src="assets/4adcdb8670214349b8370e6b4ca894be.jpg" alt="" /></p>

<p>首先，我们创建了一个lines的DStream，去监听来自本机9999端口的数据流，每一个数据代表一行文本。然后，对lines进行flatMap的转换操作，把每一个文本行拆分成词语。</p>

<p>本质上，对一个DStream进行flatMap操作，就是对它里边的每一个RDD进行flatMap操作，生成了一系列新的RDD，构成了一个新的代表词语的DStream。</p>

<p>正因为DStream和RDD的关系，RDD支持的所有转换操作，DStream都支持，比如map、flatMap、filter、union等。这些操作我们在前边学习RDD时都详细介绍过，在此不做赘述。</p>

<p>此外，DStream还有一些特有操作，如滑动窗口操作，我们可以一起探讨。</p>

<h3 id="滑动窗口操作">滑动窗口操作</h3>

<p>任何Spark Streaming的程序都要首先创建一个<strong>StreamingContext</strong>的对象，它是所有Streaming操作的入口。</p>

<p>比如，我们可以通过StreamingContext来创建DStream。前边提到的例子中，lines这个DStream就是由名为sc的StreamingContext创建的。</p>

<p>StreamingContext中最重要的参数是批处理的<strong>时间间隔</strong>，即把流数据细分成数据块的粒度。</p>

<p>这个时间间隔决定了流处理的延迟性，所以，需要我们根据需求和资源来权衡间隔的长度。上边的例子中，我们把输入的数据流以秒为单位划分，每一秒的数据会生成一个RDD进行运算。</p>

<p>有些场景中，我们需要每隔一段时间，统计过去某个时间段内的数据。比如，对热点搜索词语进行统计，每隔10秒钟输出过去60秒内排名前十位的热点词。这是流处理的一个基本应用场景，很多流处理框架如Apache Flink都有原生的支持。所以，Spark也同样支持滑动窗口操作。</p>

<p>从统计热点词这个例子，你可以看出滑动窗口操作有两个基本参数：</p>

<ul>
<li>窗口长度（window length）：每次统计的数据的时间跨度，在例子中是60秒；</li>
<li>滑动间隔（sliding interval）：每次统计的时间间隔，在例子中是10秒。</li>
</ul>

<p>显然，由于Spark Streaming流处理的最小时间单位就是StreamingContext的时间间隔，所以这两个参数一定是它的整数倍。</p>

<p><img src="assets/3a2345dddeaf45dc897ba72dcdf02441.jpg" alt="" /></p>

<p>最基本的滑动窗口操作是window，它可以返回一个新的DStream，这个DStream中每个RDD代表一段时间窗口内的数据，如下例所示。</p>

<pre><code>windowed_words = words.window(60, 10)
</code></pre>

<p>windowed_words代表的就是热词统计例子中我们所需的DStream，即它里边每一个数据块都包含过去60秒内的词语，而且这样的块每10秒钟就会生成一个。</p>

<p>此外，Spark Streaming还支持一些“进阶”窗口操作。如countByWindow、reduceByWindow、reduceByKeyAndWindow和countByValueAndWindow，在此不做深入讨论。</p>

<h2 id="spark-streaming的优缺点">Spark Streaming的优缺点</h2>

<p>讲了这么多Spark Streaming，不管内部实现也好，支持的API也好，我们还并不明白它的优势是什么，相比起其他流处理框架的缺点是什么。只有明白了这些，才能帮助我们在实际工作中决定是否使用Spark Streaming。</p>

<p>首先，Spark Streaming的优点很明显，由于它的底层是基于RDD实现的，所以RDD的优良特性在它这里都有体现。</p>

<p>比如，数据容错性，如果RDD 的某些分区丢失了，可以通过依赖信息重新计算恢复。</p>

<p>再比如运行速度，DStream同样也能通过persist()方法将数据流存放在内存中。这样做的好处是遇到需要多次迭代计算的程序时，速度优势十分明显。</p>

<p>而且，Spark Streaming是Spark生态的一部分。所以，它可以和Spark的核心引擎、Spark SQL、MLlib等无缝衔接。换句话说，对实时处理出来的中间数据，我们可以立即在程序中无缝进行批处理、交互式查询等操作。这个特点大大增强了Spark Streaming的优势和功能，使得基于Spark Streaming的应用程序很容易扩展。</p>

<p>而Spark Streaming的主要缺点是实时计算延迟较高，一般在秒的级别。这是由于Spark Streaming不支持太小的批处理的时间间隔。</p>

<p>在第二章中，我们讲过准实时和实时系统，无疑Spark Streaming是一个准实时系统。别的流处理框架，如Storm的延迟性就好很多，可以做到毫秒级。</p>

<h2 id="小结">小结</h2>

<p>Spark Streaming，作为Spark中的流处理组件，把连续的流数据按时间间隔划分为一个个数据块，然后对每个数据块分别进行批处理。</p>

<p>在内部，每个数据块就是一个RDD，所以Spark Streaming有RDD的所有优点，处理速度快，数据容错性好，支持高度并行计算。</p>

<p>但是，它的实时延迟相比起别的流处理框架比较高。在实际工作中，我们还是要具体情况具体分析，选择正确的处理框架。</p>

<h2 id="思考题">思考题</h2>

<p>如果想要优化一个Spark Streaming程序，你会从哪些角度入手？</p>

<p>欢迎你把答案写在留言区，与我和其他同学一起讨论。</p>

<p>如果你觉得有所收获，也欢迎把文章分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b6dadada8f8287878681f6d1dbd7dfda98d5d9db" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1439454d639508',t:'MTczNDA3NDc2My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
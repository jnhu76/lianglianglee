<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=17&#32;Structured&#32;Streaming：如何用DataFrame&#32;API进行实时数据分析_>
        <link rel="icon" href="/static/favicon.png">
        <title>17 Structured Streaming：如何用DataFrame API进行实时数据分析_ </title>
        
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
                            <h1 id="title" data-id="17 Structured Streaming：如何用DataFrame API进行实时数据分析_" class="title">17 Structured Streaming：如何用DataFrame API进行实时数据分析_</h1>
                            <div><p>你好，我是蔡元楠。</p>

<p>上一讲中，我们介绍了Spark中的流处理库Spark Streaming。它将无边界的流数据抽象成DStream，按特定的时间间隔，把数据流分割成一个个RDD进行批处理。所以，DStream API与RDD API高度相似，也拥有RDD的各种性质。</p>

<p>在第15讲中，我们比较过RDD和DataSet/DataFrame。你还记得DataSet/DataFrame的优点吗？你有没有想过，既然已经有了RDD API，我们为什么还要引入DataSet/DataFrame呢？</p>

<p>让我们来回顾一下DataSet/DataFrame的优点（为了方便描述，下文中我们统一用DataFrame来代指DataSet和DataFrame）：</p>

<ul>
<li>DataFrame 是<strong>高级API</strong>，提供类似于<strong>SQL</strong>的query接口，方便熟悉关系型数据库的开发人员使用；</li>
<li><strong>Spark SQL执行引擎会自动优化DataFrame程序</strong>，而用RDD API开发的程序本质上需要工程师自己构造RDD的DAG执行图，所以依赖于工程师自己去优化。</li>
</ul>

<p>那么我们自然会想到，如果可以拥有一个基于DataFrame API的流处理模块，作为工程师的我们就不需要去用相对low level的DStream API去处理无边界数据，这样会大大提升我们的开发效率。</p>

<p>基于这个思想，2016年，Spark在其2.0版本中推出了结构化流数据处理的模块Structured Streaming。</p>

<p>Structured Streaming是基于Spark SQL引擎实现的，依靠Structured Streaming，在开发者眼里，流数据和静态数据没有区别。我们完全可以像批处理静态数据那样去处理流数据。随着流数据的持续输入，Spark SQL引擎会帮助我们持续地处理新数据，并且更新计算结果。</p>

<p>今天，就让我们来一起学习Structured Streaming的原理以及应用。</p>

<h2 id="structured-streaming模型">Structured Streaming模型</h2>

<p>流数据处理最基本的问题就是如何对不断更新的无边界数据建模。</p>

<p>之前讲的Spark Streaming就是把流数据按一定的时间间隔分割成许多个小的数据块进行批处理。在Structured Streaming的模型中，我们要把数据看成一个无边界的关系型的数据表。每一个数据都是表中的一行，不断会有新的数据行被添加到表里来。我们可以对这个表做任何类似批处理的查询，Spark会帮我们不断对新加入的数据进行处理，并更新计算结果。</p>

<p><img src="assets/e4dcd36f64604316b3dd50683f85a127.jpg" alt="" /></p>

<p>与Spark Streaming类似，Structured Streaming也是将输入的数据流按照时间间隔（以一秒为例）划分成数据段。每一秒都会把新输入的数据添加到表中，Spark也会每秒更新输出结果。输出结果也是表的形式，输出表可以写入硬盘或者HDFS。</p>

<p>这里我要介绍一下Structured Streaming的三种输出模式。</p>

<ol>
<li>完全模式（Complete Mode）：整个更新过的输出表都被写入外部存储；</li>
<li>附加模式（Append Mode）：上一次触发之后新增加的行才会被写入外部存储。如果老数据有改动则不适合这个模式；</li>
<li>更新模式（Update Mode）：上一次触发之后被更新的行才会被写入外部存储。</li>
</ol>

<p>需要注意的是，Structured Streaming并不会完全存储输入数据。每个时间间隔它都会读取最新的输入，进行处理，更新输出表，然后把这次的输入删除。Structured Streaming只会存储更新输出表所需要的信息。</p>

<p>Structured Streaming的模型在根据事件时间（Event Time）处理数据时十分方便。</p>

<p>我们在第六讲中曾经讲过事件时间和处理时间（Processing Time）的区别。这里我再简单说一下。事件时间指的是事件发生的时间，是数据本身的属性；而处理时间指的是Spark接收到数据的时间。</p>

<p>很多情况下，我们需要基于事件时间来处理数据。比如说，统计每个小时接到的订单数量，一个订单很有可能在12:59被创建，但是到了13:01才被处理。</p>

<p>在Structured Streaming的模型中，由于每个数据都是输入数据表中的一行，那么事件时间就是行中的一列。依靠DataFrame API提供的类似于SQL的接口，我们可以很方便地执行基于时间窗口的查询。</p>

<h2 id="streaming-dataframe-api">Streaming DataFrame API</h2>

<p>在Structured Streaming发布以后，DataFrame既可以代表静态的有边界数据，也可以代表无边界数据。之前对静态DataFrame的各种操作同样也适用于流式DataFrame。接下来，让我们看几个例子。</p>

<h3 id="创建dataframe">创建DataFrame</h3>

<p>SparkSession.readStream()返回的DataStreamReader可以用于创建流DataFrame。它支持多种类型的数据流作为输入，比如文件、Kafka、socket等。</p>

<pre><code>socketDataFrame = spark
   .readStream
   .format(&quot;socket&quot;）
   .option(&quot;host&quot;, &quot;localhost&quot;)
   .option(&quot;port&quot;, 9999)
   .load()
</code></pre>

<p>上边的代码例子创建了一个DataFrame，用来监听来自localhost:9999的数据流。</p>

<h3 id="基本的查询操作">基本的查询操作</h3>

<p>流DataFrame同静态DataFrame一样，不仅支持类似SQL的查询操作（如select和where等），还支持RDD的转换操作（如map和filter）。 让我们一起来看下面的例子。</p>

<p>假设我们已经有一个DataFrame代表一个学生的数据流，即每个数据是一个学生，每个学生有名字（name）、年龄（age）、身高（height）和年级（grade）四个属性，我们可以用DataFrame API去做类似于SQL的Query。</p>

<pre><code>df = … // 这个DataFrame代表学校学生的数据流，schema是{name: string, age: number, height: number, grade: string}
df.select(&quot;name&quot;).where(&quot;age &gt; 10&quot;) // 返回年龄大于10岁的学生名字列表
df.groupBy(&quot;grade&quot;).count() // 返回每个年级学生的人数
df.sort_values([‘age’], ascending=False).head(100) //返回100个年龄最大的学生 
</code></pre>

<p>在这个例子中，通过第二行我们可以得到所有年龄在10岁以上的学生名字，第三行可以得到每个年级学生的人数，第四行得到100个年龄最大的学生信息。此外，DataFrame还支持很多基本的查询操作，在此不做赘述。</p>

<p>我们还可以通过isStreaming函数来判断一个DataFrame是否代表流数据。</p>

<pre><code>df.isStreaming()
</code></pre>

<p>基于事件时间的时间窗口操作</p>

<p>在学习Spark Streaming的时间窗口操作时，我们举过一个例子，是每隔10秒钟输出过去60秒的前十热点词。这个例子是基于处理时间而非事件时间的。</p>

<p>现在让我们设想一下，如果数据流中的每个词语都有一个时间戳代表词语产生的时间，那么要怎样实现，每隔10秒钟输出过去60秒内产生的前十热点词呢？你可以看看下边的代码。</p>

<pre><code>words = ...  #这个DataFrame代表词语的数据流，schema是 { timestamp: Timestamp, word: String}


windowedCounts = words.groupBy(
   window(words.timestamp, &quot;1 minute&quot;, &quot;10 seconds&quot;),
   words.word
).count()
.sort(desc(&quot;count&quot;))
.limit(10)
</code></pre>

<p>基于词语的生成时间，我们创建了一个窗口长度为1分钟，滑动间隔为10秒的window。然后，把输入的词语表根据window和词语本身聚合起来，并统计每个window内每个词语的数量。之后，再根据词语的数量进行排序，只返回前10的词语。</p>

<p>在Structured Streaming基于时间窗口的聚合操作中，groupBy是非常常用的。</p>

<h3 id="输出结果流">输出结果流</h3>

<p>当经过各种SQL查询操作之后，我们创建好了代表最终结果的DataFrame。下一步就是开始对输入数据流的处理，并且持续输出结果。</p>

<p>我们可以用Dataset.writeStream()返回的DataStreamWriter对象去输出结果。它支持多种写入位置，如硬盘文件、Kafka、console和内存等。</p>

<pre><code>query = wordCounts
   .writeStream
   .outputMode(&quot;complete&quot;)
   .format(&quot;csv&quot;)
   .option(&quot;path&quot;, &quot;path/to/destination/dir&quot;)
   .start()


query.awaitTermination()
</code></pre>

<p>在上面这个代码例子中，我们选择了完全模式，把输出结果流写入了CSV文件。</p>

<h2 id="structured-streaming与spark-streaming对比">Structured Streaming与Spark Streaming对比</h2>

<p>接下来，让我们对比一下Structured Streaming和上一讲学过的Spark Streaming。看看同为流处理的组件的它们各有什么优缺点。</p>

<h3 id="简易度和性能">简易度和性能</h3>

<p>Spark Streaming提供的DStream API与RDD API很类似，相对比较低level。</p>

<p>当我们编写 Spark Streaming 程序的时候，本质上就是要去构造RDD的DAG执行图，然后通过 Spark Engine 运行。这样开发者身上的担子就很重，很多时候要自己想办法去提高程序的处理效率。这不是Spark作为一个数据处理框架想看到的。对于好的框架来说，开发者只需要专注在业务逻辑上，而不用操心别的配置、优化等繁杂事项。</p>

<p>Structured Streaming提供的DataFrame API就是这么一个相对高level的API，大部分开发者都很熟悉关系型数据库和SQL。<strong>这样的数据抽象可以让他们用一套统一的方案去处理批处理和流处理</strong>，不用去关心具体的执行细节。</p>

<p>而且，DataFrame API是在Spark SQL的引擎上执行的，Spark SQL有非常多的优化功能，比如执行计划优化和内存管理等，所以Structured Streaming的应用程序性能很好。</p>

<h3 id="实时性">实时性</h3>

<p>在上一讲中我们了解到，Spark Streaming是准实时的，它能做到的最小延迟在一秒左右。</p>

<p>虽然Structured Streaming用的也是类似的微批处理思想，每过一个时间间隔就去拿来最新的数据加入到输入数据表中并更新结果，但是相比起Spark Streaming来说，它更像是实时处理，能做到用更小的时间间隔，最小延迟在100毫秒左右。</p>

<p>而且在Spark 2.3版本中，Structured Streaming引入了连续处理的模式，可以做到真正的毫秒级延迟，这无疑大大拓展了Structured Streaming的应用广度。不过现在连续处理模式还有很多限制，让我们期待它的未来吧。</p>

<h3 id="对事件时间的支持">对事件时间的支持</h3>

<p>就像我们在前边讲过的，Structured Streaming对基于事件时间的处理有很好的支持。</p>

<p>由于Spark Streaming是把数据按接收到的时间切分成一个个RDD来进行批处理，所以它很难基于数据本身的产生时间来进行处理。如果某个数据的处理时间和事件时间不一致的话，就容易出问题。比如，统计每秒的词语数量，有的数据先产生，但是在下一个时间间隔才被处理，这样几乎不可能输出正确的结果。</p>

<p>Structured Streaming还有很多其他优点。比如，它有更好的容错性，保证了端到端exactly once的语义等等。所以综合来说，Structured Streaming是比Spark Streaming更好的流处理工具。</p>

<h2 id="思考题">思考题</h2>

<p>在基于事件时间的窗口操作中，Structured Streaming是怎样处理晚到达的数据，并且返回正确结果的呢？</p>

<p>比如，在每十分钟统计词频的例子中，一个词语在1:09被生成，在1:11被处理，程序在1:10和1:20都输出了对应的结果，在1:20输出时为什么可以把这个词语统计在内？这样的机制有没有限制？</p>

<p>欢迎你把自己的答案写在留言区，与我和其他同学一起讨论。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d7bbbbbbeee3e6e6e7e097b0bab6bebbf9b4b8ba" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1439923e3f9508',t:'MTczNDA3NDc3NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
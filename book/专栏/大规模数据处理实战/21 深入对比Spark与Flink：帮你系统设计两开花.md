<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=21&#32;深入对比Spark与Flink：帮你系统设计两开花>
        <link rel="icon" href="/static/favicon.png">
        <title>21 深入对比Spark与Flink：帮你系统设计两开花 </title>
        
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
                            <h1 id="title" data-id="21 深入对比Spark与Flink：帮你系统设计两开花" class="title">21 深入对比Spark与Flink：帮你系统设计两开花</h1>
                            <div><p>你好，我是蔡元楠。</p>

<p>今天我要与你分享的主题是“深入对比Spark与Flink”。</p>

<p>相信通过这一模块前9讲的学习，你对Spark已经有了基本的认识。现在，我们先来回顾整个模块，理清一下思路。</p>

<p>首先，从MapReduce框架存在的问题入手，我们知道了Spark的主要优点，比如用内存运算来提高性能；提供很多High-level API；开发者无需用map和reduce两个操作实现复杂逻辑；支持流处理等等。</p>

<p>接下来，我们学习了Spark的数据抽象——RDD。RDD是整个Spark的核心概念，所有的新API在底层都是基于RDD实现的。但是RDD是否就是完美无缺的呢？显然不是，它还是很底层，不方便开发者使用，而且用RDD API写的应用程序需要大量的人工调优来提高性能。</p>

<p>Spark SQL提供的DataFrame/DataSet API就解决了这个问题，它提供类似SQL的查询接口，把数据看成关系型数据库的表，提升了熟悉关系型数据库的开发者的工作效率。这部分内容都是专注于数据的批处理，那么我们很自然地就过渡到下一个问题：Spark是怎样支持流处理的呢？</p>

<p>那就讲到了Spark Streaming和新的Structured Streaming，这是Spark的流处理组件，其中Structured Streaming也可以使用DataSet/DataFrame API，这就实现了Spark批流处理的统一。</p>

<p>通过这个简单的回顾我们发现，Spark的发布，和之后各个版本新功能的发布，并不是开发人员拍脑袋的决定，每个新版本发布的功能都是在解决旧功能的问题。在如此多的开源工作者的努力下，Spark生态系统才有今天的规模，成为了当前最流行的大数据处理框架之一。</p>

<p>在开篇词中我就提到过，我希望你能通过这个专栏建立自己的批判性思维，遇到一个新的技术，多问为什么，而不是盲目的接受和学习。只有这样我们才能不随波逐流，成为这个百花齐放的技术时代的弄潮儿。</p>

<p>所以，这里我想问你一个问题，Spark有什么缺点？</p>

<p>这个缺点我们之前已经提到过一个——无论是Spark Streaming还是Structured Streaming，Spark流处理的实时性还不够，所以无法用在一些对实时性要求很高的流处理场景中。</p>

<p>这是因为Spark的流处理是基于所谓微批处理（Micro-batch processing）的思想，即它把流处理看作是批处理的一种特殊形式，每次接收到一个时间间隔的数据才会去处理，所以天生很难在实时性上有所提升。</p>

<p>虽然在Spark 2.3中提出了连续处理模型（Continuous Processing Model），但是现在只支持很有限的功能，并不能在大的项目中使用。Spark还需要做出很大的努力才能改进现有的流处理模型。</p>

<p>想要在流处理的实时性上提升，就不能继续用微批处理的模式，而要想办法实现真正的流处理，即每当有一条数据输入就立刻处理，不做等待。那么当今时代有没有这样的流处理框架呢？</p>

<p>Apache Flink就是其中的翘楚。它采用了基于操作符（Operator）的连续流模型，可以做到微秒级别的延迟。今天我就带你一起了解一下这个流行的数据处理平台，并将Flink与Spark做深入对比，方便你在今后的实际项目中做出选择。</p>

<h2 id="flink核心模型简介">Flink核心模型简介</h2>

<p>Flink中最核心的数据结构是Stream，它代表一个运行在多个分区上的并行流。</p>

<p>在Stream上同样可以进行各种转换操作（Transformation）。与Spark的RDD不同的是，Stream代表一个数据流而不是静态数据的集合。所以，它包含的数据是随着时间增长而变化的。而且Stream上的转换操作都是逐条进行的，即每当有新的数据进来，整个流程都会被执行并更新结果。这样的基本处理模式决定了Flink会比Spark Streaming有更低的流处理延迟性。</p>

<p>当一个Flink程序被执行的时候，它会被映射为Streaming Dataflow，下图就是一个Streaming Dataflow的示意图。</p>

<p><img src="assets/573472c4cad24c289151d850ae1e92d5.jpg" alt="" /></p>

<p>在图中，你可以看出Streaming Dataflow包括Stream和Operator（操作符）。转换操作符把一个或多个Stream转换成多个Stream。每个Dataflow都有一个输入数据源（Source）和输出数据源（Sink）。与Spark的RDD转换图类似，Streaming Dataflow也会被组合成一个有向无环图去执行。</p>

<p>在Flink中，程序天生是并行和分布式的。一个Stream可以包含多个分区（Stream Partitions），一个操作符可以被分成多个操作符子任务，每一个子任务是在不同的线程或者不同的机器节点中独立执行的。如下图所示：</p>

<p><img src="assets/58bc855535324aeea425d89dd25a6dea.jpg" alt="" /></p>

<p>从上图你可以看出，Stream在操作符之间传输数据的形式有两种：一对一和重新分布。</p>

<ul>
<li>一对一（One-to-one）：Stream维护着分区以及元素的顺序，比如上图从输入数据源到map间。这意味着map操作符的子任务处理的数据和输入数据源的子任务生产的元素的数据相同。你有没有发现，它与RDD的窄依赖类似。</li>
<li>重新分布（Redistributing）：Stream中数据的分区会发生改变，比如上图中map与keyBy之间。操作符的每一个子任务把数据发送到不同的目标子任务。</li>
</ul>

<h2 id="flink的架构">Flink的架构</h2>

<p>当前版本Flink的架构如下图所示。</p>

<p><img src="assets/1f0ad779199f41e3a2282d8d4c6bf790.jpg" alt="" /></p>

<p>我们可以看到，这个架构和<a href="https://time.geekbang.org/column/article/94410" target="_blank">第12讲</a>中介绍的Spark架构比较类似，都分为四层：存储层、部署层、核心处理引擎、high-level的API和库。</p>

<p>从存储层来看，Flink同样兼容多种主流文件系统如HDFS、Amazon S3，多种数据库如HBase和多种数据流如Kafka和Flume。</p>

<p>从部署层来看，Flink不仅支持本地运行，还能在独立集群或者在被 YARN 或 Mesos 管理的集群上运行，也能部署在云端。</p>

<p>核心处理引擎就是我们刚才提到的分布式Streaming Dataflow，所有的高级API及应用库都会被翻译成包含Stream和Operator的Dataflow来执行。</p>

<p>Flink提供的两个核心API就是DataSet API和DataStream API。你没看错，名字和Spark的DataSet、DataFrame非常相似。顾名思义，DataSet代表有界的数据集，而DataStream代表流数据。所以，DataSet API是用来做批处理的，而DataStream API是做流处理的。</p>

<p>也许你会问，Flink这样基于流的模型是怎样支持批处理的？在内部，DataSet其实也用Stream表示，静态的有界数据也可以被看作是特殊的流数据，而且DataSet与DataStream可以无缝切换。所以，Flink的核心是DataStream。</p>

<p>DataSet和DataStream都支持各种基本的转换操作如map、filter、count、groupBy等，让我们来看一个用DataStream实现的统计词频例子。</p>

<pre><code>public class WindowWordCount {
 public static void main(String[] args) throws Exception {
   StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();


   DataStream&lt;Tuple2&lt;String, Integer&gt;&gt; dataStream = env
　　.socketTextStream(&quot;localhost&quot;, 9999)
　　.flatMap(new Splitter())
　　.keyBy(0)
　　.timeWindow(Time.seconds(5))
　　.sum(1);


   dataStream.print();
   env.execute(&quot;Window WordCount&quot;);
}


 public static class Splitter implements FlatMapFunction&lt;String, Tuple2&lt;String, Integer&gt;&gt; {
   @Override
   public void flatMap(String sentence, Collector&lt;Tuple2&lt;String, Integer&gt;&gt; out) {
     for (String word: sentence.split(&quot; &quot;)) {
       out.collect(new Tuple2&lt;String, Integer&gt;(word, 1));
     }
   }
 }
</code></pre>

<p>这里我是用Java来示范的，因为Flink就是用Java开发的，所以它对Java有原生的支持。此外，也可以用Scala来开发Flink程序，在1.0版本后更是支持了Python。</p>

<p>在这个例子中，我们首先创建了一个Splitter类，来把输入的句子拆分成（词语，1）的对。在主程序中用StreamExecutionEnvironment创建DataStream，来接收本地Web Socket的文本流，并进行了4步操作。</p>

<ol>
<li>用flatMap把输入文本拆分成（词语，1）的对；</li>
<li>用keyBy把相同的词语分配到相同的分区；</li>
<li>设好5秒的时间窗口；</li>
<li>对词语的出现频率用sum求和。</li>
</ol>

<p>可以看出，DataStream的使用方法和RDD比较相似，都是把程序拆分成一系列的转换操作并分布式地执行。</p>

<p>在DataSet和DataStream之上，有更高层次的Table API。Table API和Spark SQL的思想类似，是关系型的API，用户可以像操作SQL数据库表一样的操作数据，而不需要通过写Java代码、操作DataStream/DataSet的方式进行数据处理，更不需要手动优化代码的执行逻辑。</p>

<p>此外，Table API同样统一了Flink的批处理和流处理。</p>

<h2 id="flink和spark对比">Flink和Spark对比</h2>

<p>通过前面的学习，我们了解到，Spark和Flink都支持批处理和流处理，接下来让我们对这两种流行的数据处理框架在各方面进行对比。</p>

<p>首先，这两个数据处理框架有很多相同点。</p>

<ul>
<li>都基于内存计算；</li>
<li>都有统一的批处理和流处理API，都支持类似SQL的编程接口；</li>
<li>都支持很多相同的转换操作，编程都是用类似于Scala Collection API的函数式编程模式；</li>
<li>都有完善的错误恢复机制；</li>
<li>都支持Exactly once的语义一致性。</li>
</ul>

<p>当然，它们的不同点也是相当明显，我们可以从4个不同的角度来看。</p>

<p><strong>从流处理的角度来讲</strong>，Spark基于微批量处理，把流数据看成是一个个小的批处理数据块分别处理，所以延迟性只能做到秒级。而Flink基于每个事件处理，每当有新的数据输入都会立刻处理，是真正的流式计算，支持毫秒级计算。由于相同的原因，Spark只支持基于时间的窗口操作（处理时间或者事件时间），而Flink支持的窗口操作则非常灵活，不仅支持时间窗口，还支持基于数据本身的窗口，开发者可以自由定义想要的窗口操作。</p>

<p><strong>从SQL功能的角度来讲</strong>，Spark和Flink分别提供SparkSQL和Table API提供SQL交互支持。两者相比较，Spark对SQL支持更好，相应的优化、扩展和性能更好，而 Flink 在 SQL 支持方面还有很大提升空间。</p>

<p><strong>从迭代计算的角度来讲</strong>，Spark对机器学习的支持很好，因为可以在内存中缓存中间计算结果来加速机器学习算法的运行。但是大部分机器学习算法其实是一个有环的数据流，在Spark中，却是用无环图来表示。而Flink支持在运行时间中的有环数据流，从而可以更有效的对机器学习算法进行运算。</p>

<p><strong>从相应的生态系统角度来讲</strong>，Spark的社区无疑更加活跃。Spark可以说有着Apache旗下最多的开源贡献者，而且有很多不同的库来用在不同场景。而Flink由于较新，现阶段的开源社区不如Spark活跃，各种库的功能也不如Spark全面。但是Flink还在不断发展，各种功能也在逐渐完善。</p>

<h2 id="小结">小结</h2>

<p>今天我们从Spark存在的一个缺点——无法高效应对低延迟的流处理场景入手，一起学习了另一个主流流数据处理框架Flink，还对比了这两个框架的异同，相信现在你对两个框架肯定有了更多的认识。</p>

<p>我经常被问到的一个问题是：Spark和Flink到底该选哪一个？对于这个问题，我们还是要分一下场景。</p>

<p>对于以下场景，你可以选择Spark。</p>

<ol>
<li>数据量非常大而且逻辑复杂的批数据处理，并且对计算效率有较高要求（比如用大数据分析来构建推荐系统进行个性化推荐、广告定点投放等）；</li>
<li>基于历史数据的交互式查询，要求响应较快；</li>
<li>基于实时数据流的数据处理，延迟性要求在在数百毫秒到数秒之间。</li>
</ol>

<p>Spark完美满足这些场景的需求， 而且它可以一站式解决这些问题，无需用别的数据处理平台。</p>

<p>由于Flink是为了提升流处理而创建的平台，所以它适用于各种需要非常低延迟（微秒到毫秒级）的实时数据处理场景，比如实时日志报表分析。</p>

<p>而且Flink用流处理去模拟批处理的思想，比Spark用批处理去模拟流处理的思想扩展性更好，所以我相信将来Flink会发展的越来越好，生态和社区各方面追上Spark。比如，阿里巴巴就基于Flink构建了公司范围内全平台使用的数据处理平台Blink，美团、饿了么等公司也都接受Flink作为数据处理解决方案。</p>

<p>可以说，Spark和Flink都在某种程度上统一了批处理和流处理，但也都有一些不足。下一模块中，让我们来一起学习一个全新的、完全统一批流处理的数据处理平台——Apache Beam，到时候我们会对Spark的优缺点有更加深入的认识。</p>

<h2 id="思考题">思考题</h2>

<p>除了高延迟的流处理这一缺点外，你认为Spark还有什么不足？可以怎样改进？</p>

<p>欢迎你把答案写在留言区，与我和其他同学一起讨论。如果你觉得有所收获，也欢迎把文章分享给你的朋友。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b0dcdcdc898481818087f0d7ddd1d9dc9ed3dfdd" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f143a6b99899508',t:'MTczNDA3NDgxMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
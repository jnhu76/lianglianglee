<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;如何测试Beam&#32;Pipeline？>
        <link rel="icon" href="/static/favicon.png">
        <title>29 如何测试Beam Pipeline？ </title>
        
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
                            <h1 id="title" data-id="29 如何测试Beam Pipeline？" class="title">29 如何测试Beam Pipeline？</h1>
                            <div><p>你好，我是蔡元楠。</p>

<p>今天我要与你分享的主题是“如何测试Beam Pipeline”。</p>

<p>在上一讲中，我们结合了第7讲的内容，一起学习了在Beam的世界中我们该怎么设计好对应的设计模式。而在今天这一讲中，我想要讲讲在日常开发中经常会被忽略的，但是又非常重要的一个开发环节——测试。</p>

<p>你知道，我们设计好的Beam数据流水线通常都会被放在分布式环境下执行，具体每一步的Transform都会被分配到任意的机器上面执行。如果我们在运行数据流水线时发现结果出错了，那么想要定位到具体的机器，再到上面去做调试是不现实的。</p>

<p>当然还有另一种方法，读取一些样本数据集，再运行整个数据流水线去验证哪一步逻辑出错了。但这是一项非常耗时耗力的工作。即便我们可以把样本数据集定义得非常小，从而缩短运行数据流水线运行所需的时间。但是万一我们所写的是多步骤数据流水线的话，就不知道到底在哪一步出错了，我们必须把每一步的中间结果输出出来进行调试。</p>

<p>基于以上种种的原因，在我们正式将数据流水线放在分布式环境上面运行之前，先完整地测试好整个数据流水线逻辑，就变得尤为重要了。</p>

<p>为了解决这些问题，Beam提供了一套完整的测试SDK。让我们可以在开发数据流水线的同时，能够实现对一个Transform逻辑的单元测试，也可以对整个数据流水线端到端（End-to-End）地测试。</p>

<p>在Beam所支持的各种Runners当中，有一个Runner叫作DirectRunner。DirectRunner其实就是我们的本地机器。也就是说，如果我们指定Beam的Runner为DirectRunner的话，整个Beam数据流水线都会放在本地机器上面运行。我们在运行测试程序的时候可以利用这个DirectRunner来跑测试逻辑。</p>

<p>在正式讲解之前，有一点是我需要提醒你的。如果你喜欢自行阅读Beam的相关技术文章或者是示例代码的话，可能你会看见一些测试代码使用了在Beam SDK中的一个测试类，叫作DoFnTester来进行单元测试。这个DoFnTester类可以让我们传入一个用户自定义的函数（User Defined Function/UDF）来进行测试。</p>

<p>通过<a href="https://time.geekbang.org/column/article/101735" target="_blank">第25讲</a>的内容我们已经知道，一个最简单的Transform可以用一个ParDo来表示，在使用它的时候，我们需要继承DoFn这个抽象类。这个DoFnTester接收的对象就是我们继承实现的DoFn。在这里，我们把一个DoFn看作是一个单元来进行测试了。但这并不是Beam所提倡的。</p>

<p>因为在Beam中，数据转换的逻辑都是被抽象成Transform，而不是Transform里面的ParDo这些具体的实现。<strong>每个Runner具体怎么运行这些ParDo</strong>，<strong>对于用户来说应该都是透明的。</strong>所以，在Beam的2.4.0版本之后，Beam SDK将这个类标记成了Deprecated，转而推荐使用Beam SDK中的TestPipeline。</p>

<p>所以，我在这里也建议你，在写测试代码的时候，不要使用任何和DoFnTester有关的SDK。</p>

<h2 id="beam的transform单元测试">Beam的Transform单元测试</h2>

<p>说完了注意事项，那事不宜迟，我们就先从一个Transform的单元测试开始，看看在Beam是如何做测试的（以下所有的测试示例代码都是以Java为编程语言来讲解）。</p>

<p>一般来说，Transform的单元测试可以通过以下五步来完成：</p>

<ol>
<li>创建一个Beam测试SDK中所提供的TestPipeline实例。</li>
<li>创建一个静态（Static）的、用于测试的输入数据集。</li>
<li>使用Create Transform来创建一个PCollection作为输入数据集。</li>
<li>在测试数据集上调用我们需要测试的Transform上并将结果保存在一个PCollection上。</li>
<li>使用PAssert类的相关函数来验证输出的PCollection是否是我所期望的结果。</li>
</ol>

<p>假设我们要处理的数据集是一个整数集合，处理逻辑是过滤掉数据集中的奇数，将输入数据集中的偶数输出。为此，我们通过继承DoFn类来实现一个产生偶数的Transform，它的输入和输出数据类型都是Integer。</p>

<p>Java</p>

<pre><code>static class EvenNumberFn extends DoFn&lt;Integer, Integer&gt; {
 @ProcessElement
 public void processElement(@Element Integer in, OutputReceiver&lt;Integer&gt; out) {
   if (in % 2 == 0) {
     out.output(in);
   }
 }
}
</code></pre>

<p>那我们接下来就根据上面所讲的测试流程，测试这个EvenNumerFn Transform，来一步步创建我们的单元测试。</p>

<h3 id="创建testpipeline">创建TestPipeline</h3>

<p>第一步，创建TestPipeline。创建一个TestPipeline实例的代码非常简单，示例如下：</p>

<p>Java</p>

<pre><code>...
Pipeline p = TestPipeline.create();
...
</code></pre>

<p>如果你还记得在<a href="https://time.geekbang.org/column/article/102182" target="_blank">第26讲</a>中如何创建数据流水线的话，可以发现，TestPipeline实例的创建其实不用给这个TestPipeline定义选项（Options）。因为TestPipeline中create函数已经在内部帮我们创建好一个测试用的Options了。</p>

<h3 id="创建静态输入数据集">创建静态输入数据集</h3>

<p>Java</p>

<pre><code>...
static final List&lt;Integer&gt; INPUTS = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
...
</code></pre>

<p>第二步，创建静态的输入数据集。创建静态的输入数据集的操作就和我们平时所写的普通Java代码一样，在示例中，我调用了Arrays类的asList接口来创建一个拥有10个整数的数据集。</p>

<h3 id="使用create-transform创建pcollection">使用Create Transform创建PCollection</h3>

<p>在创建完静态数据集后，我们进入第三步，创建一个PCollection作为输入数据集。在Beam原生支持的Transform里面，有一种叫作Create Transform，我们可以利用这个Create Transform将Java Collection的数据转换成为Beam的数据抽象PCollection，具体的做法如下：</p>

<p>Java</p>

<pre><code>...
PCollection&lt;Integer&gt; input = p.apply(Create.of(INPUTS)).setCoder(VarIntCoder.of());
...
</code></pre>

<h3 id="调用transform处理逻辑">调用Transform处理逻辑</h3>

<p>第四步，调用Transform处理逻辑。有了数据抽象PCollection，我们就需要在测试数据集上调用我们需要测试的Transform处理逻辑，并将结果保存在一个PCollection上。</p>

<p>Java</p>

<pre><code>...
PCollection&lt;String&gt; output = input.apply(ParDo.of(new EvenNumberFn()));
...
</code></pre>

<p>根据<a href="https://time.geekbang.org/column/article/101735" target="_blank">第25讲</a>的内容，我们只需要在这个输入数据集上调用apply抽象函数，生成一个需要测试的Transform，并且传入apply函数中就可以了。</p>

<h3 id="验证输出结果">验证输出结果</h3>

<p>第五步，验证输出结果。在验证结果的阶段，我们需要调用PAssert类中的函数来验证输出结果是否和我们期望的一致，示例如下。</p>

<p>Java</p>

<pre><code>...
PAssert.that(output).containsInAnyOrder(2, 4, 6, 8, 10);
...
</code></pre>

<p>完成了所有的步骤，我们就差运行这个测试的数据流水线了。很简单，就是调用TestPipeline的run函数，整个Transform的单元测试示例如下：</p>

<p>Java</p>

<pre><code>final class TestClass {
    static final List&lt;Integer&gt; INPUTS = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

    public void testFn() {
      Pipeline p = TestPipeline.create();
      PCollection&lt;Integer&gt; input = p.apply(Create.of(INPUTS)).setCoder(VarIntCoder.of());
      PCollection&lt;String&gt; output = input.apply(ParDo.of(new EvenNumberFn()));
      PAssert.that(output).containsInAnyOrder(2, 4, 6, 8, 10);
      p.run();
    }
}
</code></pre>

<p>有一点需要注意的是，TestPipeline的run函数是在单元测试的结尾处调用的，PAssert的调用必须在TestPipeliner调用run函数之前调用。</p>

<h2 id="beam的端到端测试">Beam的端到端测试</h2>

<p>在一般的现实应用中，我们设计的都是多步骤数据流水线，就拿我在<a href="https://time.geekbang.org/column/article/90081" target="_blank">第一讲</a>中举到的处理美团外卖电动车的图片为例子，其中就涉及到了多个输入数据集，而结果也有可能会根据实际情况有多个输出。</p>

<p>所以，我们在做测试的时候，往往希望能有一个端到端的测试。在Beam中，端到端的测试和Transform的单元测试非常相似。唯一的不同点在于，我们要为所有的输入数据集创建测试数据集，而不是只针对某一个Transform来创建。对于在数据流水线的每一个应用到Write Transfrom的地方，我们都需要用到PAssert类来验证输出数据集。</p>

<p>所以，端到端测试的步骤也分五步，具体内容如下：</p>

<ol>
<li>创建一个Beam测试SDK中所提供的TestPipeline实例。</li>
<li>对于多步骤数据流水线中的每个输入数据源，创建相对应的静态（Static）测试数据集。</li>
<li>使用Create Transform，将所有的这些静态测试数据集转换成PCollection作为输入数据集。</li>
<li>按照真实数据流水线逻辑，调用所有的Transforms操作。</li>
<li>在数据流水线中所有应用到Write Transform的地方，都使用PAssert来替换这个Write Transform，并且验证输出的结果是否我们期望的结果相匹配。</li>
</ol>

<p>为了方便说明，我们就在之前的例子中多加一步Transform和一个输出操作来解释如何写端到端测试。假设，我们要处理数据集是一个整数集合，处理逻辑是过滤掉奇数，将输入数据集中的偶数转换成字符串输出。同时，我们也希望对这些偶数求和并将结果输出，示例如下：</p>

<p>Java</p>

<pre><code>final class Foo {
  static class EvenNumberFn extends DoFn&lt;Integer, Integer&gt; {
    @ProcessElement
    public void processElement(@Element Integer in, OutputReceiver&lt;Integer&gt; out) {
      if (in % 2 == 0) {
        out.output(in);
      }
    }
  }

  static class ParseIntFn extends DoFn&lt;String, Integer&gt; {
    @ProcessElement
    public void processElement(@Element String in, OutputReceiver&lt;Integer&gt; out) {
      out.output(Integer.parseInt(in));
    }
  }

  public static void main(String[] args) {
    PipelineOptions options = PipelineOptionsFactory.create();
    Pipeline p = Pipeline.create(options);
    PCollection&lt;Integer&gt; input = p.apply(TextIO.read().from(&quot;filepath/input&quot;)).apply(ParDo.of(new ParseIntFn()));
    PCollection&lt;Integer&gt; output1 = input.apply(ParDo.of(new EvenNumberFn()));
    output1.apply(ToString.elements()).apply(TextIO.write().to(&quot;filepath/evenNumbers&quot;));
    PCollection&lt;Integer&gt; sum = output1.apply(Combine.globally(new SumInts()));
    sum.apply(ToString.elements()).apply(TextIO.write().to(&quot;filepath/sum&quot;));
    p.run();
  }
}
</code></pre>

<p>从上面的示例代码中你可以看到，我们从一个外部源读取了一系列输入数据进来，将它转换成了整数集合。同时，将我们自己编写的EvenNumberFn Transform应用在了这个输入数据集上。得到了所有偶数集合之后，我们先将这个中间结果输出，然后再针对这个偶数集合求和，最后将这个结果输出。</p>

<p>整个数据流水线总共有一次对外部数据源的读取和两次的输出，我们按照端到端测试的步骤，为所有的输入数据集创建静态数据，然后将所有有输出的地方都使用PAssert类来进行验证。整个测试程序如下所示：</p>

<p>Java</p>

<pre><code>final class TestClass {

  static final List&lt;String&gt; INPUTS =
      Arrays.asList(&quot;1&quot;, &quot;2&quot;, &quot;3&quot;, &quot;4&quot;, &quot;5&quot;, &quot;6&quot;, &quot;7&quot;, &quot;8&quot;, &quot;9&quot;, &quot;10&quot;);

  static class EvenNumberFn extends DoFn&lt;Integer, Integer&gt; {
    @ProcessElement
    public void processElement(@Element Integer in, OutputReceiver&lt;Integer&gt; out) {
      if (in % 2 == 0) {
        out.output(in);
      }
    }
  }

  static class ParseIntFn extends DoFn&lt;String, Integer&gt; {
    @ProcessElement
    public void processElement(@Element String in, OutputReceiver&lt;Integer&gt; out) {
      out.output(Integer.parseInt(in));
    }
  }

  public void testFn() {
    Pipeline p = TestPipeline.create();
    PCollection&lt;String&gt; input = p.apply(Create.of(INPUTS)).setCoder(StringUtf8Coder.of());
    PCollection&lt;Integer&gt; output1 = input.apply(ParDo.of(new ParseIntFn())).apply(ParDo.of(new EvenNumberFn()));
    PAssert.that(output1).containsInAnyOrder(2, 4, 6, 8, 10);
    PCollection&lt;Integer&gt; sum = output1.apply(Combine.globally(new SumInts()));
    PAssert.that(sum).is(30);
    p.run();
  }
}
</code></pre>

<p>在上面的示例代码中，我们用TestPipeline替换了原来的Pipeline，创建了一个静态输入数据集并用Create Transform转换成了PCollection，最后将所有用到Write Transform的地方都用PAssert替换掉，来验证输出结果是否是我们期望的结果。</p>

<h2 id="小结">小结</h2>

<p>今天我们一起学习了在Beam中写编写测试逻辑的两种方式，分别是针对一个Transform的单元测试和针对整个数据流水线的端到端测试。Beam提供的SDK能够让我们不需要在分布式环境下运行程序而是本地机器上运行。测试在整个开发环节中是非常的一环，我强烈建议你在正式上线自己的业务逻辑之前先对此有一个完整的测试。</p>

<h2 id="思考题">思考题</h2>

<p>如果让你来利用Beam SDK来测试你日常处理的数据逻辑，你会如何编写测试呢？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ed818181d4d9dcdcdddaad8a808c8481c38e8280" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f143cb5ef219508',t:'MTczNDA3NDkwMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
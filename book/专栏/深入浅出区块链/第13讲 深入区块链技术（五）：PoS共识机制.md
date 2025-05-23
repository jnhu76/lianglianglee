<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=第13讲&#32;深入区块链技术（五）：PoS共识机制>
        <link rel="icon" href="/static/favicon.png">
        <title>第13讲 深入区块链技术（五）：PoS共识机制 </title>
        
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
                        <a class="menu-item" id="00 开篇词 帮你从0到1深入学习区块链技术.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%b8%ae%e4%bd%a0%e4%bb%8e0%e5%88%b01%e6%b7%b1%e5%85%a5%e5%ad%a6%e4%b9%a0%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af.md">00 开篇词 帮你从0到1深入学习区块链技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="尾声篇 授人以鱼，不如授人以渔.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e5%b0%be%e5%a3%b0%e7%af%87%20%e6%8e%88%e4%ba%ba%e4%bb%a5%e9%b1%bc%ef%bc%8c%e4%b8%8d%e5%a6%82%e6%8e%88%e4%ba%ba%e4%bb%a5%e6%b8%94.md">尾声篇 授人以鱼，不如授人以渔.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="新书首发《区块链第一课：深入浅出技术与应用》.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e6%96%b0%e4%b9%a6%e9%a6%96%e5%8f%91%e3%80%8a%e5%8c%ba%e5%9d%97%e9%93%be%e7%ac%ac%e4%b8%80%e8%af%be%ef%bc%9a%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ba%94%e7%94%a8%e3%80%8b.md">新书首发《区块链第一课：深入浅出技术与应用》.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第01讲 到底什么才是区块链？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac01%e8%ae%b2%20%e5%88%b0%e5%ba%95%e4%bb%80%e4%b9%88%e6%89%8d%e6%98%af%e5%8c%ba%e5%9d%97%e9%93%be%ef%bc%9f.md">第01讲 到底什么才是区块链？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第02讲 区块链到底是怎么运行的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac02%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e5%88%b0%e5%ba%95%e6%98%af%e6%80%8e%e4%b9%88%e8%bf%90%e8%a1%8c%e7%9a%84%ef%bc%9f.md">第02讲 区块链到底是怎么运行的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第03讲 浅说区块链共识机制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac03%e8%ae%b2%20%e6%b5%85%e8%af%b4%e5%8c%ba%e5%9d%97%e9%93%be%e5%85%b1%e8%af%86%e6%9c%ba%e5%88%b6.md">第03讲 浅说区块链共识机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第04讲 区块链的应用类型.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac04%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e7%9a%84%e5%ba%94%e7%94%a8%e7%b1%bb%e5%9e%8b.md">第04讲 区块链的应用类型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第05讲 如何理解数字货币？它与区块链又是什么样的关系？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac05%e8%ae%b2%20%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81%ef%bc%9f%e5%ae%83%e4%b8%8e%e5%8c%ba%e5%9d%97%e9%93%be%e5%8f%88%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">第05讲 如何理解数字货币？它与区块链又是什么样的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第06讲 理解区块链之前，先上手体验一把数字货币.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac06%e8%ae%b2%20%e7%90%86%e8%a7%a3%e5%8c%ba%e5%9d%97%e9%93%be%e4%b9%8b%e5%89%8d%ef%bc%8c%e5%85%88%e4%b8%8a%e6%89%8b%e4%bd%93%e9%aa%8c%e4%b8%80%e6%8a%8a%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81.md">第06讲 理解区块链之前，先上手体验一把数字货币.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第07讲 区块链的常见误区.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac07%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e7%9a%84%e5%b8%b8%e8%a7%81%e8%af%af%e5%8c%ba.md">第07讲 区块链的常见误区.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第08讲 最主流区块链项目有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac08%e8%ae%b2%20%e6%9c%80%e4%b8%bb%e6%b5%81%e5%8c%ba%e5%9d%97%e9%93%be%e9%a1%b9%e7%9b%ae%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">第08讲 最主流区块链项目有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第09讲 深入区块链技术（一）：技术基础.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac09%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%8a%80%e6%9c%af%e5%9f%ba%e7%a1%80.md">第09讲 深入区块链技术（一）：技术基础.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第10讲 深入区块链技术（二）：P2P网络.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac10%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aP2P%e7%bd%91%e7%bb%9c.md">第10讲 深入区块链技术（二）：P2P网络.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第11讲 深入区块链技术（三）：共识算法与分布式一致性算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac11%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%85%b1%e8%af%86%e7%ae%97%e6%b3%95%e4%b8%8e%e5%88%86%e5%b8%83%e5%bc%8f%e4%b8%80%e8%87%b4%e6%80%a7%e7%ae%97%e6%b3%95.md">第11讲 深入区块链技术（三）：共识算法与分布式一致性算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第12讲 深入区块链技术（四）：PoW共识.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac12%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9aPoW%e5%85%b1%e8%af%86.md">第12讲 深入区块链技术（四）：PoW共识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第13讲 深入区块链技术（五）：PoS共识机制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac13%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e4%ba%94%ef%bc%89%ef%bc%9aPoS%e5%85%b1%e8%af%86%e6%9c%ba%e5%88%b6.md">第13讲 深入区块链技术（五）：PoS共识机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第14讲 深入区块链技术（六）：DPoS共识机制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac14%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e5%85%ad%ef%bc%89%ef%bc%9aDPoS%e5%85%b1%e8%af%86%e6%9c%ba%e5%88%b6.md">第14讲 深入区块链技术（六）：DPoS共识机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第15讲 深入区块链技术（七）：哈希与加密算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac15%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%83%ef%bc%89%ef%bc%9a%e5%93%88%e5%b8%8c%e4%b8%8e%e5%8a%a0%e5%af%86%e7%ae%97%e6%b3%95.md">第15讲 深入区块链技术（七）：哈希与加密算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第16讲 深入区块链技术（八）： UTXO与普通账户模型.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac16%e8%ae%b2%20%e6%b7%b1%e5%85%a5%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%ef%bc%88%e5%85%ab%ef%bc%89%ef%bc%9a%20UTXO%e4%b8%8e%e6%99%ae%e9%80%9a%e8%b4%a6%e6%88%b7%e6%a8%a1%e5%9e%8b.md">第16讲 深入区块链技术（八）： UTXO与普通账户模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第17讲  去中心化与区块链交易性能.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac17%e8%ae%b2%20%20%e5%8e%bb%e4%b8%ad%e5%bf%83%e5%8c%96%e4%b8%8e%e5%8c%ba%e5%9d%97%e9%93%be%e4%ba%a4%e6%98%93%e6%80%a7%e8%83%bd.md">第17讲  去中心化与区块链交易性能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第18讲 智能合约与以太坊.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac18%e8%ae%b2%20%e6%99%ba%e8%83%bd%e5%90%88%e7%ba%a6%e4%b8%8e%e4%bb%a5%e5%a4%aa%e5%9d%8a.md">第18讲 智能合约与以太坊.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第19讲 上手搭建一条自己的智能合约.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac19%e8%ae%b2%20%e4%b8%8a%e6%89%8b%e6%90%ad%e5%bb%ba%e4%b8%80%e6%9d%a1%e8%87%aa%e5%b7%b1%e7%9a%84%e6%99%ba%e8%83%bd%e5%90%88%e7%ba%a6.md">第19讲 上手搭建一条自己的智能合约.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第20讲 区块链项目详解：比特股BTS.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac20%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e9%a1%b9%e7%9b%ae%e8%af%a6%e8%a7%a3%ef%bc%9a%e6%af%94%e7%89%b9%e8%82%a1BTS.md">第20讲 区块链项目详解：比特股BTS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第21讲 引人瞩目的区块链项目：EOS、IOTA、Cardano.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac21%e8%ae%b2%20%e5%bc%95%e4%ba%ba%e7%9e%a9%e7%9b%ae%e7%9a%84%e5%8c%ba%e5%9d%97%e9%93%be%e9%a1%b9%e7%9b%ae%ef%bc%9aEOS%e3%80%81IOTA%e3%80%81Cardano.md">第21讲 引人瞩目的区块链项目：EOS、IOTA、Cardano.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第22讲 国内区块链项目技术一览.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac22%e8%ae%b2%20%e5%9b%bd%e5%86%85%e5%8c%ba%e5%9d%97%e9%93%be%e9%a1%b9%e7%9b%ae%e6%8a%80%e6%9c%af%e4%b8%80%e8%a7%88.md">第22讲 国内区块链项目技术一览.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第23讲 联盟链和它的困境.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac23%e8%ae%b2%20%e8%81%94%e7%9b%9f%e9%93%be%e5%92%8c%e5%ae%83%e7%9a%84%e5%9b%b0%e5%a2%83.md">第23讲 联盟链和它的困境.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第24讲 比特币专题（一）历史与货币.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac24%e8%ae%b2%20%e6%af%94%e7%89%b9%e5%b8%81%e4%b8%93%e9%a2%98%ef%bc%88%e4%b8%80%ef%bc%89%e5%8e%86%e5%8f%b2%e4%b8%8e%e8%b4%a7%e5%b8%81.md">第24讲 比特币专题（一）历史与货币.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第25讲 比特币专题（二）：扩容之争、IFO与链上治理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac25%e8%ae%b2%20%e6%af%94%e7%89%b9%e5%b8%81%e4%b8%93%e9%a2%98%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e6%89%a9%e5%ae%b9%e4%b9%8b%e4%ba%89%e3%80%81IFO%e4%b8%8e%e9%93%be%e4%b8%8a%e6%b2%bb%e7%90%86.md">第25讲 比特币专题（二）：扩容之争、IFO与链上治理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第26讲 数字货币和数字资产.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac26%e8%ae%b2%20%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81%e5%92%8c%e6%95%b0%e5%ad%97%e8%b5%84%e4%ba%a7.md">第26讲 数字货币和数字资产.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第27讲 弄懂数字货币交易平台（一）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac27%e8%ae%b2%20%e5%bc%84%e6%87%82%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81%e4%ba%a4%e6%98%93%e5%b9%b3%e5%8f%b0%ef%bc%88%e4%b8%80%ef%bc%89.md">第27讲 弄懂数字货币交易平台（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第28讲 弄懂数字货币交易平台（二）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac28%e8%ae%b2%20%e5%bc%84%e6%87%82%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81%e4%ba%a4%e6%98%93%e5%b9%b3%e5%8f%b0%ef%bc%88%e4%ba%8c%ef%bc%89.md">第28讲 弄懂数字货币交易平台（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第29讲 互联网身份与区块链数字身份.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac29%e8%ae%b2%20%e4%ba%92%e8%81%94%e7%bd%91%e8%ba%ab%e4%bb%bd%e4%b8%8e%e5%8c%ba%e5%9d%97%e9%93%be%e6%95%b0%e5%ad%97%e8%ba%ab%e4%bb%bd.md">第29讲 互联网身份与区块链数字身份.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第30讲 区块链即服务BaaS.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac30%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e5%8d%b3%e6%9c%8d%e5%8a%a1BaaS.md">第30讲 区块链即服务BaaS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第31讲 数字货币钱包服务.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac31%e8%ae%b2%20%e6%95%b0%e5%ad%97%e8%b4%a7%e5%b8%81%e9%92%b1%e5%8c%85%e6%9c%8d%e5%8a%a1.md">第31讲 数字货币钱包服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第32讲 区块链与供应链（一）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac32%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e4%b8%8e%e4%be%9b%e5%ba%94%e9%93%be%ef%bc%88%e4%b8%80%ef%bc%89.md">第32讲 区块链与供应链（一）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第33讲 区块链与供应链（二）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac33%e8%ae%b2%20%e5%8c%ba%e5%9d%97%e9%93%be%e4%b8%8e%e4%be%9b%e5%ba%94%e9%93%be%ef%bc%88%e4%ba%8c%ef%bc%89.md">第33讲 区块链与供应链（二）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第34讲 从业区块链需要了解什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac34%e8%ae%b2%20%e4%bb%8e%e4%b8%9a%e5%8c%ba%e5%9d%97%e9%93%be%e9%9c%80%e8%a6%81%e4%ba%86%e8%a7%a3%e4%bb%80%e4%b9%88%ef%bc%9f.md">第34讲 从业区块链需要了解什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第35讲 搭建你的迷你区块链（设计篇 ）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac35%e8%ae%b2%20%e6%90%ad%e5%bb%ba%e4%bd%a0%e7%9a%84%e8%bf%b7%e4%bd%a0%e5%8c%ba%e5%9d%97%e9%93%be%ef%bc%88%e8%ae%be%e8%ae%a1%e7%af%87%20%ef%bc%89.md">第35讲 搭建你的迷你区块链（设计篇 ）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="第36讲 搭建你的迷你区块链（实践篇）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8c%ba%e5%9d%97%e9%93%be/%e7%ac%ac36%e8%ae%b2%20%e6%90%ad%e5%bb%ba%e4%bd%a0%e7%9a%84%e8%bf%b7%e4%bd%a0%e5%8c%ba%e5%9d%97%e9%93%be%ef%bc%88%e5%ae%9e%e8%b7%b5%e7%af%87%ef%bc%89.md">第36讲 搭建你的迷你区块链（实践篇）.md</a>
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
                            <h1 id="title" data-id="第13讲 深入区块链技术（五）：PoS共识机制" class="title">第13讲 深入区块链技术（五）：PoS共识机制</h1>
                            <div><p>上一篇我们讲到了PoW共识机制，这一篇我们就来分享另外一种共识机制，PoS共识机制。</p>

<p>PoS全称是Proof of Stake，中文翻译为权益证明。这一篇我们会将PoS与PoW对比讲解，帮助你加深理解。</p>

<h2 id="pos的由来">PoS的由来</h2>

<p>PoS最早出现在点点币的创始人Sunny King的白皮书中，它的目的就是为了解决使用PoW挖矿出现大量资源浪费的问题。PoS共识机制一经提出就引起了广泛关注，Sunny King 也基于PoW的基础框架实现了第一代PoS区块链：点点币。</p>

<p>PoW的具体实现有很多版本，但它们大多只是在挖矿算法上有所改进，主体逻辑并没有发生质的变化。PoS包含了多个变种实现，每个变种往往会涉及区块链代币经济模型的改动，可以说是牵一发而动全身。</p>

<p>这些实现有点点币、黑币、未来币、瑞迪币，它们都推动了PoS机制的发展，PoS研究前沿还有以太坊的 Casper ，以及 Cardano 的Ouroboros。</p>

<p>那到底是什么样的机制导致PoS具有这样的特性呢？让我们来看一看。</p>

<h2 id="什么是pos">什么是PoS？</h2>

<p>在讲PoS之前，我先来讲一个叫做币龄的概念，币龄这个概念其实很好理解，它的英文是 CoinAge，字面意思就是币数量乘以天数。</p>

<p>比如你有100个币，在某个地址上9天没有动，那么产生的币龄就是900，如果你把这个地址上这100币转移到任意地址，包括你自己的地址，那么900个币龄就在转移过程中被花费了，你的币数量虽然还是100个，但是币龄变更为0。币龄在数据链上就可以取到，任何人都可以验证。</p>

<p>我们回过头来看看PoS究竟是什么，区块链共识机制的第一步就是随机筛选一个记账者，PoW是通过计算能力来获得记账权，计算能力越强，获得记账权的概率越大。</p>

<p>PoS则将此处的计算能力更换为财产证明，就是节点所拥有的币龄越多，获得的记账的概率就越大。这有点像公司的股权结构，股权占比大的合伙人话语权越重。</p>

<p>以上算是简述了PoS的概念，实际上，PoS的发展经历了三个版本，第一个版本是以点点币为代表的PoS1.0版本，这个版本中使用的是币龄；第二个版本的代表是黑币（blackcoin），对应使用的是币数量，相当于是财产证明，后面黑币又升级到PoS3.0，这个版本又回到了币龄。</p>

<p>PoW早在比特币出现之前就已经应用了，而PoS是才是真正意义上为了区块链而创造出来的概念。</p>

<h2 id="pos的实现原理">PoS的实现原理</h2>

<p>好了，现在我们开始讲解PoS的具体实现原理吧。这一部分公式较多，如果你在收听音频，可以点击文稿查看。</p>

<p>通过上一篇我们知道PoW挖矿的基本逻辑和步骤，我们先寻求一个nonce小于目标值，这一步用公式可表示为：</p>

<p><strong>Hash (block_header) &lt; Target</strong></p>

<p>从公式中我们可以看到，PoW下所有矿工的目标值是一样的，只要计算结果哈希小于目标值即可，简化来看就是前导0的个数。</p>

<p>而在PoS系统中，这个公式变更为：</p>

<p><strong>Hash (block_header) &lt; Target * CoinAge</strong></p>

<p>我们可以看出多引入了一个变量叫做CoinAge，也就是币龄，这里就有意思了。</p>

<p>这个变量为会造成每个矿工看到的目标值不一样，如果你的币龄越大，也就意味着你的获得答案越容易。这里的Target与PoW一致，与全网难度成反比，用来控制出块速度的。</p>

<p>例如当前全网的目标是4369，A矿工的输入的币龄是15，那么A矿工的目标值为65535，换算成十六进制就是0xFFFF，完整的哈希长度假设是8个字节，也就是0x0000FFFF。</p>

<p>而B矿工比较有钱，他输入的币龄是240，那么B矿工的目标值就是0x000FFFFF。你如果仔细观察肯定会发现，相比A矿工的目标值，B直接少了一个零。即如下：</p>

<ul>
<li>A 矿工 Hash( block_header ) &lt; 0x0000FFFF</li>
<li>B 矿工 Hash( block_header ) &lt; 0x000FFFFF</li>
</ul>

<p>所以B矿工获得记账权的概率肯定要比A高。</p>

<p>具体代码分析这里就不讲解了，这里需要币龄作为输入，如果我们写示例代码也只是一个简单的参数，PoS需要放到区块链的语境中才能运作。</p>

<h2 id="pos的相关问题">PoS的相关问题</h2>

<p>通过上述的介绍我们知道：PoS似乎完美地解决了PoW挖矿资源浪费的问题，甚至还顺带解决了51%攻击的问题，这里可以顺便讲一下51%攻击是什么，它是指PoW矿工如果积累了超过51%的算力，则可以一定程度篡改账本。</p>

<p>这里顺便科普一下，什么是51%攻击呢，我们发现，矿工挖矿的成本不再是物理设备和电费，而是虚拟代币，它的边际成本几乎为零，本质上PoS让挖矿者和使用者合二为一了。</p>

<p>这也意味着如果挖矿者发起51%攻击，就需要拥有全网51%的币或币龄，这几乎不可能办到，即使你成功地实施了51%攻击，那么也意味着作为全网最大的持币大户的你，损失也会最大。</p>

<p>PoS看起来相当完美，其实并不然，PoS有很多缺陷。</p>

<p>PoS遇到的第一个问题就是币发行的问题。一开始的时候，只有创始区块上有币，意味着只有这一个节点可以挖矿，所以让币分散出去才能让整个网络壮大，那么如何分散出去又是另外一个难题了。</p>

<p>所以早期PoS币种基本都采用了分阶段挖矿，有的叫混合挖矿，其实，我并不同意混合挖矿这个说法，混合就意味着同时。很多币种其实是分了阶段的，即第一阶段是PoW挖矿，到第二阶段才是PoS挖矿。</p>

<p>随着ERC20类型的标准合约代币的出现，这个问题被解决了，不再需要第一阶段改成PoW，也可以将代币分散出去。</p>

<p>第二个问题是由于币龄是与时间挂钩的，这也意味着用户可以无限囤积一定的币，等过了很久再一次性挖矿发起攻击；所以解决方案是：PoS机制需要引入一个时间上限来控制时间因素的自然增长。</p>

<p>第三个问题是虽然引入了时间上下限，用户还是倾向于囤积代币，这会造成币流通的不充分；基于此，所以瑞迪币引入了币龄按时间衰减，构造了权益速度证明，鼓励用户流动代币，而不是倾向于囤积代币。</p>

<p>第四个问题是离线攻击，即使引入了时间上下限，时间仍然是自然流动的，也就是可以不需要求挖矿节点长时间在线。挖矿是可以离线的，这简直是灾难，所以任意一个PoS机制的实践形式都必须避免这个问题，因为网络节点数量的多少直接关系到区块链网络的健壮性。</p>

<p>当然这些问题都不是致命问题，还记得我们一开始提到了PoS经历了三个版本，而第二个版本PoS 2.0使用的不是币龄，而直接是币的数量。</p>

<p>这会造成完全不同的结果，上述第二、三、四问题都不存在了，似乎看起来直接使用币的数量会更好一些，但却出现了整个PoS机制的致命问题。</p>

<p>这个问题叫做Nothing at Stake，翻译过来叫做无成本利益问题。大体的意思在PoS系统中做任何事几乎没有成本，比如在PoS系统上挖矿几乎没有成本，这也就意味着分叉非常方便。</p>

<p>方便到什么程度呢，每个诚实矿工在产生孤块的时候都可以继续挖下去，反正也没什么成本，反正分叉链和主链都可以同时挖，也就是任何持币较少的用户都可以尝试分叉，并且把分叉链广播出去。</p>

<p>这个时候如果其他诚实矿工看到了，第一反应也是没有成本，那么咱们也来挖吧，说不定什么时候就值钱了，意思就是说任何逐利的矿工并不会使这个系统变得更强壮稳定，而是更加的混乱。</p>

<p>无成本利益问题无论以币龄还是币数量作为PoS的参数，都无法避免。</p>

<p>而PoW则没有这样的问题，我们回到PoW系统中来看，因为任何的分叉都会造成挖矿成本直接变成负收益，所以这会抵抗分叉的产生，矿工倾向于跟随“最长”的链。</p>

<p>由于以太坊部分采用了PoS共识，它的名字叫做Casper，它必须解决上述无成本利益问题攻击。所以Casper协议要求PoS矿工需通过抵押保证金的方法对共识结果进行下注，具体实践结果我们还需要拭目以待。</p>

<h2 id="总结">总结</h2>

<p>最后我们来总结一下PoS共识机制，PoS的区块链系统无需外部物理输入，所以它相比PoW更为环保不费电，并且矿工就是使用者，这会在一定程度上抵御了51%攻击，所以基于PoS机制的数字货币属于理想状态的数字货币。</p>

<p>PoS的缺点是缺乏工业级的区块链应用，从逻辑上来看有点循环自证明的味道，就是用自己的币来维护系统的安全，而币的安全性是由系统保证的，所以现阶段PoS共识机制往往不是独立运行的，而是混合了PoW一起运行，这就可以弥补PoS的缺陷。</p>

<p>PoS共识机制目前也出现了矿池，也可能会出现中心化挖矿的风险。</p>

<p>虽然PoS共识机制未来变数依然很多，但它的可塑性比PoW好，技术上的探索空间大，目前PoS币种相比较PoW币种风险也较高。</p>

<p>那么有哪些区块链项目使用了PoS共识机制呢？你可以给我留言，我们一起讨论，感谢你的收听，我们下期再见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d3bfbfbfeae7e2e2e3e493b4beb2babffdb0bcbe" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f163e36dbe03696',t:'MTczNDA5NTkzNy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=20&#32;用RNN构建个性化音乐播单>
        <link rel="icon" href="/static/favicon.png">
        <title>20 用RNN构建个性化音乐播单 </title>
        
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
                        <a class="menu-item" id="00 开篇词 用知识去对抗技术不平等.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e7%94%a8%e7%9f%a5%e8%af%86%e5%8e%bb%e5%af%b9%e6%8a%97%e6%8a%80%e6%9c%af%e4%b8%8d%e5%b9%b3%e7%ad%89.md">00 开篇词 用知识去对抗技术不平等.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 你真的需要个性化推荐系统吗_.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/01%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e9%9c%80%e8%a6%81%e4%b8%aa%e6%80%a7%e5%8c%96%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e5%90%97_.md">01 你真的需要个性化推荐系统吗_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 个性化推荐系统有哪些绕不开的经典问题？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/02%20%e4%b8%aa%e6%80%a7%e5%8c%96%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bb%95%e4%b8%8d%e5%bc%80%e7%9a%84%e7%bb%8f%e5%85%b8%e9%97%ae%e9%a2%98%ef%bc%9f.md">02 个性化推荐系统有哪些绕不开的经典问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 这些你必须应该具备的思维模式.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/03%20%e8%bf%99%e4%ba%9b%e4%bd%a0%e5%bf%85%e9%a1%bb%e5%ba%94%e8%af%a5%e5%85%b7%e5%a4%87%e7%9a%84%e6%80%9d%e7%bb%b4%e6%a8%a1%e5%bc%8f.md">03 这些你必须应该具备的思维模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 画鬼容易画人难：用户画像的“能”和“不能”.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/04%20%e7%94%bb%e9%ac%bc%e5%ae%b9%e6%98%93%e7%94%bb%e4%ba%ba%e9%9a%be%ef%bc%9a%e7%94%a8%e6%88%b7%e7%94%bb%e5%83%8f%e7%9a%84%e2%80%9c%e8%83%bd%e2%80%9d%e5%92%8c%e2%80%9c%e4%b8%8d%e8%83%bd%e2%80%9d.md">04 画鬼容易画人难：用户画像的“能”和“不能”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 从文本到用户画像有多远.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/05%20%e4%bb%8e%e6%96%87%e6%9c%ac%e5%88%b0%e7%94%a8%e6%88%b7%e7%94%bb%e5%83%8f%e6%9c%89%e5%a4%9a%e8%bf%9c.md">05 从文本到用户画像有多远.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 超越标签的内容推荐系统.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/06%20%e8%b6%85%e8%b6%8a%e6%a0%87%e7%ad%be%e7%9a%84%e5%86%85%e5%ae%b9%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f.md">06 超越标签的内容推荐系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 人以群分，你是什么人就看到什么世界.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/07%20%e4%ba%ba%e4%bb%a5%e7%be%a4%e5%88%86%ef%bc%8c%e4%bd%a0%e6%98%af%e4%bb%80%e4%b9%88%e4%ba%ba%e5%b0%b1%e7%9c%8b%e5%88%b0%e4%bb%80%e4%b9%88%e4%b8%96%e7%95%8c.md">07 人以群分，你是什么人就看到什么世界.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 解密“看了又看”和“买了又买”.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/08%20%e8%a7%a3%e5%af%86%e2%80%9c%e7%9c%8b%e4%ba%86%e5%8f%88%e7%9c%8b%e2%80%9d%e5%92%8c%e2%80%9c%e4%b9%b0%e4%ba%86%e5%8f%88%e4%b9%b0%e2%80%9d.md">08 解密“看了又看”和“买了又买”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 协同过滤中的相似度计算方法有哪些.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/09%20%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%e4%b8%ad%e7%9a%84%e7%9b%b8%e4%bc%bc%e5%ba%a6%e8%ae%a1%e7%ae%97%e6%96%b9%e6%b3%95%e6%9c%89%e5%93%aa%e4%ba%9b.md">09 协同过滤中的相似度计算方法有哪些.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 那些在Netflix Prize中大放异彩的推荐算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/10%20%e9%82%a3%e4%ba%9b%e5%9c%a8Netflix%20Prize%e4%b8%ad%e5%a4%a7%e6%94%be%e5%bc%82%e5%bd%a9%e7%9a%84%e6%8e%a8%e8%8d%90%e7%ae%97%e6%b3%95.md">10 那些在Netflix Prize中大放异彩的推荐算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 Facebook是怎么为十亿人互相推荐好友的.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/11%20Facebook%e6%98%af%e6%80%8e%e4%b9%88%e4%b8%ba%e5%8d%81%e4%ba%bf%e4%ba%ba%e4%ba%92%e7%9b%b8%e6%8e%a8%e8%8d%90%e5%a5%bd%e5%8f%8b%e7%9a%84.md">11 Facebook是怎么为十亿人互相推荐好友的.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 如果关注排序效果，那么这个模型可以帮到你.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/12%20%e5%a6%82%e6%9e%9c%e5%85%b3%e6%b3%a8%e6%8e%92%e5%ba%8f%e6%95%88%e6%9e%9c%ef%bc%8c%e9%82%a3%e4%b9%88%e8%bf%99%e4%b8%aa%e6%a8%a1%e5%9e%8b%e5%8f%af%e4%bb%a5%e5%b8%ae%e5%88%b0%e4%bd%a0.md">12 如果关注排序效果，那么这个模型可以帮到你.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 经典模型融合办法：线性模型和树模型的组合拳.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/13%20%e7%bb%8f%e5%85%b8%e6%a8%a1%e5%9e%8b%e8%9e%8d%e5%90%88%e5%8a%9e%e6%b3%95%ef%bc%9a%e7%ba%bf%e6%80%a7%e6%a8%a1%e5%9e%8b%e5%92%8c%e6%a0%91%e6%a8%a1%e5%9e%8b%e7%9a%84%e7%bb%84%e5%90%88%e6%8b%b3.md">13 经典模型融合办法：线性模型和树模型的组合拳.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 一网打尽协同过滤、矩阵分解和线性模型.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/14%20%e4%b8%80%e7%bd%91%e6%89%93%e5%b0%bd%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%e3%80%81%e7%9f%a9%e9%98%b5%e5%88%86%e8%a7%a3%e5%92%8c%e7%ba%bf%e6%80%a7%e6%a8%a1%e5%9e%8b.md">14 一网打尽协同过滤、矩阵分解和线性模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 深度和宽度兼具的融合模型 Wide and Deep.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/15%20%e6%b7%b1%e5%ba%a6%e5%92%8c%e5%ae%bd%e5%ba%a6%e5%85%bc%e5%85%b7%e7%9a%84%e8%9e%8d%e5%90%88%e6%a8%a1%e5%9e%8b%20Wide%20and%20Deep.md">15 深度和宽度兼具的融合模型 Wide and Deep.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 简单却有效的Bandit算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/16%20%e7%ae%80%e5%8d%95%e5%8d%b4%e6%9c%89%e6%95%88%e7%9a%84Bandit%e7%ae%97%e6%b3%95.md">16 简单却有效的Bandit算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 结合上下文信息的Bandit算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/17%20%e7%bb%93%e5%90%88%e4%b8%8a%e4%b8%8b%e6%96%87%e4%bf%a1%e6%81%af%e7%9a%84Bandit%e7%ae%97%e6%b3%95.md">17 结合上下文信息的Bandit算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 如何将Bandit算法与协同过滤结合使用.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/18%20%e5%a6%82%e4%bd%95%e5%b0%86Bandit%e7%ae%97%e6%b3%95%e4%b8%8e%e5%8d%8f%e5%90%8c%e8%bf%87%e6%bb%a4%e7%bb%93%e5%90%88%e4%bd%bf%e7%94%a8.md">18 如何将Bandit算法与协同过滤结合使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 深度学习在推荐系统中的应用有哪些_.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/19%20%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e5%9c%a8%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%ad%e7%9a%84%e5%ba%94%e7%94%a8%e6%9c%89%e5%93%aa%e4%ba%9b_.md">19 深度学习在推荐系统中的应用有哪些_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 用RNN构建个性化音乐播单.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/20%20%e7%94%a8RNN%e6%9e%84%e5%bb%ba%e4%b8%aa%e6%80%a7%e5%8c%96%e9%9f%b3%e4%b9%90%e6%92%ad%e5%8d%95.md">20 用RNN构建个性化音乐播单.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 构建一个科学的排行榜体系.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/21%20%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e7%a7%91%e5%ad%a6%e7%9a%84%e6%8e%92%e8%a1%8c%e6%a6%9c%e4%bd%93%e7%b3%bb.md">21 构建一个科学的排行榜体系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 实用的加权采样算法.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/22%20%e5%ae%9e%e7%94%a8%e7%9a%84%e5%8a%a0%e6%9d%83%e9%87%87%e6%a0%b7%e7%ae%97%e6%b3%95.md">22 实用的加权采样算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 推荐候选池的去重策略.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/23%20%e6%8e%a8%e8%8d%90%e5%80%99%e9%80%89%e6%b1%a0%e7%9a%84%e5%8e%bb%e9%87%8d%e7%ad%96%e7%95%a5.md">23 推荐候选池的去重策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 典型的信息流架构是什么样的.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/24%20%e5%85%b8%e5%9e%8b%e7%9a%84%e4%bf%a1%e6%81%af%e6%b5%81%e6%9e%b6%e6%9e%84%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84.md">24 典型的信息流架构是什么样的.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 Netflix个性化推荐架构.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/25%20Netflix%e4%b8%aa%e6%80%a7%e5%8c%96%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84.md">25 Netflix个性化推荐架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 总览推荐架构和搜索、广告的关系.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/26%20%e6%80%bb%e8%a7%88%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84%e5%92%8c%e6%90%9c%e7%b4%a2%e3%80%81%e5%b9%bf%e5%91%8a%e7%9a%84%e5%85%b3%e7%b3%bb.md">26 总览推荐架构和搜索、广告的关系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 巧妇难为无米之炊：数据采集关键要素.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/27%20%e5%b7%a7%e5%a6%87%e9%9a%be%e4%b8%ba%e6%97%a0%e7%b1%b3%e4%b9%8b%e7%82%8a%ef%bc%9a%e6%95%b0%e6%8d%ae%e9%87%87%e9%9b%86%e5%85%b3%e9%94%ae%e8%a6%81%e7%b4%a0.md">27 巧妇难为无米之炊：数据采集关键要素.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 让你的推荐系统反应更快：实时推荐.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/28%20%e8%ae%a9%e4%bd%a0%e7%9a%84%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e5%8f%8d%e5%ba%94%e6%9b%b4%e5%bf%ab%ef%bc%9a%e5%ae%9e%e6%97%b6%e6%8e%a8%e8%8d%90.md">28 让你的推荐系统反应更快：实时推荐.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 让数据驱动落地，你需要一个实验平台.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/29%20%e8%ae%a9%e6%95%b0%e6%8d%ae%e9%a9%b1%e5%8a%a8%e8%90%bd%e5%9c%b0%ef%bc%8c%e4%bd%a0%e9%9c%80%e8%a6%81%e4%b8%80%e4%b8%aa%e5%ae%9e%e9%aa%8c%e5%b9%b3%e5%8f%b0.md">29 让数据驱动落地，你需要一个实验平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 推荐系统服务化、存储选型及API设计.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/30%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e6%9c%8d%e5%8a%a1%e5%8c%96%e3%80%81%e5%ad%98%e5%82%a8%e9%80%89%e5%9e%8b%e5%8f%8aAPI%e8%ae%be%e8%ae%a1.md">30 推荐系统服务化、存储选型及API设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 推荐系统的测试方法及常用指标介绍.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/31%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%b5%8b%e8%af%95%e6%96%b9%e6%b3%95%e5%8f%8a%e5%b8%b8%e7%94%a8%e6%8c%87%e6%a0%87%e4%bb%8b%e7%bb%8d.md">31 推荐系统的测试方法及常用指标介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 道高一尺魔高一丈：推荐系统的攻防.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/32%20%e9%81%93%e9%ab%98%e4%b8%80%e5%b0%ba%e9%ad%94%e9%ab%98%e4%b8%80%e4%b8%88%ef%bc%9a%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%94%bb%e9%98%b2.md">32 道高一尺魔高一丈：推荐系统的攻防.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 和推荐系统有关的开源工具及框架介绍.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/33%20%e5%92%8c%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e6%9c%89%e5%85%b3%e7%9a%84%e5%bc%80%e6%ba%90%e5%b7%a5%e5%85%b7%e5%8f%8a%e6%a1%86%e6%9e%b6%e4%bb%8b%e7%bb%8d.md">33 和推荐系统有关的开源工具及框架介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 推荐系统在互联网产品商业链条中的地位.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/34%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e5%9c%a8%e4%ba%92%e8%81%94%e7%bd%91%e4%ba%a7%e5%93%81%e5%95%86%e4%b8%9a%e9%93%be%e6%9d%a1%e4%b8%ad%e7%9a%84%e5%9c%b0%e4%bd%8d.md">34 推荐系统在互联网产品商业链条中的地位.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 说说信息流的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/35%20%e8%af%b4%e8%af%b4%e4%bf%a1%e6%81%af%e6%b5%81%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">35 说说信息流的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 组建推荐团队及工程师的学习路径.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/36%20%e7%bb%84%e5%bb%ba%e6%8e%a8%e8%8d%90%e5%9b%a2%e9%98%9f%e5%8f%8a%e5%b7%a5%e7%a8%8b%e5%b8%88%e7%9a%84%e5%ad%a6%e4%b9%a0%e8%b7%af%e5%be%84.md">36 组建推荐团队及工程师的学习路径.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 推荐系统的参考阅读.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/%e5%8a%a0%e9%a4%90%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e7%9a%84%e5%8f%82%e8%80%83%e9%98%85%e8%af%bb.md">加餐 推荐系统的参考阅读.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 遇“荐”之后，江湖再见.md" href="/%e4%b8%93%e6%a0%8f/%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e4%b8%89%e5%8d%81%e5%85%ad%e5%bc%8f/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e9%81%87%e2%80%9c%e8%8d%90%e2%80%9d%e4%b9%8b%e5%90%8e%ef%bc%8c%e6%b1%9f%e6%b9%96%e5%86%8d%e8%a7%81.md">结束语 遇“荐”之后，江湖再见.md</a>
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
                            <h1 id="title" data-id="20 用RNN构建个性化音乐播单" class="title">20 用RNN构建个性化音乐播单</h1>
                            <div><p>时间是一个客观存在的物理属性，很多数据都有时间属性，只不过大多时候都把它忽略掉了。前面讲到的绝大多数推荐算法，也都没有考虑“用户在产品上作出任何行为”都是有时间先后的。</p>

<p>正是认识到这一点，有一些矩阵分解算法考虑了时间属性，比如Time-SVD；但是，这种做法只是把时间作为一个独立特征加入到模型中，仍然没有给时间一个正确的名分。</p>

<h2 id="时间的重要性">时间的重要性</h2>

<p>时间属性反应在序列的先后上，比如用户在视频网站上观看电视剧会先看第一集再看第二集，股市数据先有昨天的再有今天的，说“我订阅了《推荐系统36式》专栏”这句话时，词语也有先后，这种先后的关系就是时间序列。</p>

<p>具体到推荐系统领域，时间序列就是用户操作行为的先后。绝大数推荐算法都忽略操作的先后顺序，为什么要采取这样简化的做法呢？因为一方面的确也能取得不错的效果，另一方面是深度学习和推荐系统还迟迟没有相见。</p>

<p>在深度学习大火之后，对时间序列建模被提上议事日程，业界有很多尝试，今天以Spotify的音乐推荐为例，介绍循环神经网络在推荐系统中的应用。</p>

<h2 id="循环神经网络">循环神经网络</h2>

<p>循环神经网络，也常被简称为RNN，是一种特殊的神经网络。再回顾一下神经网络的结构，示意图如下：</p>

<p><img src="assets/61177f813f5f493966f8f9beaa395dc6.png" alt="" /></p>

<p>普通神经网络有三个部分，输入层x，隐藏层h，输出层o，深度神经网络的区别就是隐藏层数量有很多，具体多少算深，这个可没有定论，有几层的，也有上百层的。</p>

<p>把输入层和隐藏层之间的关系表示成公式后就是：</p>
<p><span class="math display">\[h = F(Wx) \]</span></p><p>就是输入层x经过连接参数线性加权后，再有激活函数F变换成非线性输出给输出层。</p>

<p>在输出层就是：</p>
<p><span class="math display">\[O = \\phi(Vh) \]</span></p><p>隐藏层输出经过输出层的网络连接参数线性加权后，再由输出函数变换成最终输出，比如分类任务就是Softmax函数。</p>

<p>那循环神经网络和普通神经网络的区别在哪？</p>

<p>区别就在于：普通神经网络的隐藏层参数只有输入x决定，因为当神经网络在面对一条样本时，这条样本是孤立的，不考虑前一个样本是什么，循环神经网络的隐藏层不只是受输入x影响，还受上一个时刻的隐藏层参数影响。</p>

<p>我们把这个表示成示意图如下：-
<img src="assets/285fdbb62e517ddb2099d3b6c87f8371.png" alt="" /></p>

<p>解释一下这个示意图。在时刻t，输入是xt，而隐藏层的输出不再是只有输入层xt，还有时刻t-1的隐藏层输出h(t-1)，表示成公式就是：</p>
<p><span class="math display">\[h_{t} = F(Wx_{t} + Uh_{t-1})\]</span></p><p>对比这个公式和前面普通神经网络的隐藏层输出，就是在激活函数的输入处多了一个 <span class="math inline">\(Uh_{t-1}\)</span> 。别小看多这一个小东西，它背后的意义非凡。</p>

<p>我一直在传递一个观点，隐藏层的东西，包括矩阵分解和各种Embedding得到的隐因子，是对很多表面纷繁复杂的现象所做的信息抽取和信息压缩。</p>

<p>那么上一个时刻得到的隐藏层，就是对时间序列上一个时刻的信息压缩，让它参与到这一个时刻的隐藏层建设上来，物理意义就是认为现在这个时刻的信息不只和现在的输入有关，还和上一个时刻的状态有关。这是时间序列本来的意义，也就是循环神经网络的意义。</p>

<h2 id="播单生成">播单生成</h2>

<p>了解了循环神经网络原理之后，我再和你一起来看下它如何应用在推荐系统中的。</p>

<p>在网络音乐推荐中，尤其是各类FM类App，提倡的是一直听下去，比如是你在做家务时，你在开车时，一首歌接着一首歌地播下去，就很适合这些场景。</p>

<p>通常要做到这样的效果，有这么几种做法。</p>

<ol>
<li>电台音乐DJ手工编排播单，然后一直播放下去，传统广播电台都是这样的。</li>
<li>用非时序数据离线计算出推荐集合，然后按照分数顺序逐一输出。</li>
<li>利用循环神经网络，把音乐播单的生成看成是歌曲时间序列的生成，每一首歌的得到不但受用户当前的特征影响，还受上一首歌影响。</li>
</ol>

<p>Spotify采用了第三种办法，下面我就详细讲解这个推荐算法。</p>

<h3 id="1-数据">1.数据</h3>

<p>个性化的播单生成，不再是推荐一个一个独立的音乐，而是推荐一个序列给用户。所用的数据就是已有播单，或者用户的会话信息。其中用户会话信息的意思就是，当一个用户在App上所做的一系列操作。</p>

<p>把这些数据，看成一个一个的文档，每一个音乐文件就是一个一个的词。听完什么再听什么，就像是语言中的词和词的关系。</p>

<h3 id="2-建模">2.建模</h3>

<p>你可以把播单生成看成由若干步骤组成，每一步吐出一个音乐来。这个吐出音乐的动作实际上是一个多分类问题，类别数目就是总共可以选择的音乐数目，如果有100万首歌可以选择，那么就是一个100万分类任务。</p>

<p>这个分类任务计算输入是当前神经网络的隐藏状态，然后每一首歌都得到一个线性加权值，再由Softmax函数为每一首歌计算得到一个概率。表示如下：</p>
<p><span class="math display">\[p(o_{ti} | h_{t}) = \\frac{e^{v_{i}h}}{\\sum_{j \\in M}{e^{v_{j}h}}} \]</span></p><p>假如隐藏层有k个神经元，也就是说h是一个k维向量，输出层有m首歌可选，所以是一个One-hot编码的向量，也就是说一个m维向量，只有真正输出那首歌i是1，其他都是0，那么输出层就有k乘以m个未知参数。</p>

<p>再往前，计算隐藏层神经元输出时，不但用到输入层的信息，在这里，输入层也是一首歌，也有m首歌可以选择，所以输入向量仍然是一个One-hot编码的向量。</p>

<p>除此之外，每一个隐藏层神经元还依赖上一个时刻自己的输出值，隐藏层神经元是k个，一个k维向量。</p>

<p>按照隐藏层计算公式就是下面的样子。</p>
<p><span class="math display">\[h_{t} = F(Wx_{t} + Uh_{t-1})\]</span></p><p>W就是一个m乘以k的参数矩阵，U就是一个k乘以k的参数矩阵。</p>

<p>如此一来，循环神经网络在预测时的计算过程就是：</p>

<p>当用户听完一首歌，要预测下一首歌该推荐什么时，输入就是一个One-hot编码的m维度向量，用m乘以k形状的输入层参数矩阵，乘以这个m向量，然后用隐藏层之间的k乘k参数矩阵，去乘以上一个隐藏状态向量，两者都得到一个k维向量，相加后经过非线性激活函数，比如ReLU，这样就得到当前时刻的隐藏层输出值。</p>

<p>再用当前时刻的隐藏层输出值，经过k乘以m形状的输出层参数矩阵，得到一个m维向量，再用Softmax把这个m维向量归一化成概率值，就是对下一首歌的预测，可以挑选最大概率的若干首歌作为输出，或者直接输出概率最高的那首歌直接播放。</p>

<p>这个计算过程示意图如下：</p>

<p><img src="assets/79e67fc2dd3172dd1e331d6ebd5e9a28.png" alt="" /></p>

<p>一个播单生成模型的参数就是这么三大块。</p>

<ol>
<li>连接输入和隐藏之间的矩阵 <span class="math inline">\(W_{m\\times{k}}\)</span>；</li>
<li>连接上一个隐藏状态和当前隐藏状态的矩阵： <span class="math inline">\(U_{k \\times {k}}\)</span>；</li>
<li>连接隐藏层和输出层的矩阵 <span class="math inline">\(V_{k\\times{m}}\)</span>。</li>
</ol>

<p>得到了这些参数，就得到了播单推荐模型，怎么得到呢？这里就再简要讲一下神经网络的参数如何训练得到。</p>

<p>你知道一个简单的逻辑回归模型参数如何训练得到吗？大致是这样几步：</p>

<ol>
<li>初始化参数；</li>
<li>用当前的参数预测样本的类别概率；</li>
<li>用预测的概率计算交叉熵；</li>
<li>用交叉熵计算参数的梯度；</li>
<li>用学习步长和梯度更新参数；</li>
<li>迭代上述过程直到满足设置的条件。</li>
</ol>

<p>神经网络的参数学习大致也是这个过程，但略为复杂的地方就是第4步和第5步，因为逻辑回归没有隐藏层，神经网络有隐藏层。那怎么办呢？我不打算讲解具体的做法，我打算给你建立一个直观印象。</p>

<p>还记得下面这个函数对x求导是怎么计算的吗？</p>
<p><span class="math display">\[f(x) = g(x)^2;\]</span></p><p><span class="math display">\[g(x) = e^x\]</span></p><p>函数f(x)是另一个函数gx的平方，函数g(x)又是一个指数函数。那么要对f(x0求导，就是一个链式规则，先把g(x)看成个一个整体求导，再乘以g(x)的求导结果：</p>
<p><span class="math display">\[f^{’}(x) = 2g(x)e^{x} = 2e^{x}e^{x} = 2e^{2x}\]</span></p><p>你就需要记住一点：链式规则，一路求导下去。</p>

<p>现在回到神经网路的训练，这个方法有个高大上的名字，叫做误差方向传播。</p>

<p>实际上就是链式求导法则，因为要更新参数，就需要计算参数在当前取值时的梯度，要计算梯度就要求导，要求导就要从交叉熵函数开始，先对输出层参数求导计算梯度，更新输出层参数，接着链式下去，对输入层参数求导计算梯度，更新输入层参数。</p>

<p>交叉熵是模型的目标函数，训练模型的目的就是要最小化它，也就是“误差反向传播”的“误差”。</p>

<p>相信聪明如你已经在直观上理解了一个普通神经网络是怎么训练的了，那么一个循环神经网络的参数训练有何不同呢？唯一不同就是多了一个参数矩阵，连接当前隐藏层和上一次隐藏层的参数矩阵U，也是链式求导法则的传播路径，也就是多了一些求导计算，更新参数方式并没有什么不同。</p>

<h2 id="总结">总结</h2>

<p>好了，今天介绍了如何使用循环神经网络推荐音乐播单，播单是一个时间序列，听完上一首歌会影响下一首歌。</p>

<p>循环神经网络和普通神经网络相比，就是在两个时刻的隐藏状态之间多了网络连接。看上去这个网络连接只与上一个时刻有关，事实上，上一个状态又与上上个状态有关，所以实际上任意一个时刻的状态是与此前所有的状态有关的。</p>

<p>今天的应用虽然是以播单推荐为例，但其实循环神经网络还可以应用在很多其他地方，你对循环神经网络的应用有任何问题都可以留言给我，我们一起讨论。</p>

<p>感谢你的收听，我们下期再见。</p>

<p><img src="assets/1025014d381c4913aae6eaef553d7750.jpg" alt="" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a6cacaca9f9297979691e6c1cbc7cfca88c5c9cb" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f14fea5fd474596',t:'MTczNDA4Mjg0Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
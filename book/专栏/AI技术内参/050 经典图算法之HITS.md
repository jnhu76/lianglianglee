<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=050&#32;经典图算法之HITS>
        <link rel="icon" href="/static/favicon.png">
        <title>050 经典图算法之HITS </title>
        
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
                        <a class="menu-item" id="000 开篇词 你的360度人工智能信息助理.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/000%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bd%a0%e7%9a%84360%e5%ba%a6%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e4%bf%a1%e6%81%af%e5%8a%a9%e7%90%86.md">000 开篇词 你的360度人工智能信息助理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="001 聊聊2017年KDD大会的时间检验奖.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/001%20%e8%81%8a%e8%81%8a2017%e5%b9%b4KDD%e5%a4%a7%e4%bc%9a%e7%9a%84%e6%97%b6%e9%97%b4%e6%a3%80%e9%aa%8c%e5%a5%96.md">001 聊聊2017年KDD大会的时间检验奖.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="002 精读2017年KDD最佳研究论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/002%20%e7%b2%be%e8%af%bb2017%e5%b9%b4KDD%e6%9c%80%e4%bd%b3%e7%a0%94%e7%a9%b6%e8%ae%ba%e6%96%87.md">002 精读2017年KDD最佳研究论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="003 精读2017年KDD最佳应用数据科学论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/003%20%e7%b2%be%e8%af%bb2017%e5%b9%b4KDD%e6%9c%80%e4%bd%b3%e5%ba%94%e7%94%a8%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e8%ae%ba%e6%96%87.md">003 精读2017年KDD最佳应用数据科学论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="004 精读2017年EMNLP最佳长论文之一.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/004%20%e7%b2%be%e8%af%bb2017%e5%b9%b4EMNLP%e6%9c%80%e4%bd%b3%e9%95%bf%e8%ae%ba%e6%96%87%e4%b9%8b%e4%b8%80.md">004 精读2017年EMNLP最佳长论文之一.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="005 精读2017年EMNLP最佳长论文之二.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/005%20%e7%b2%be%e8%af%bb2017%e5%b9%b4EMNLP%e6%9c%80%e4%bd%b3%e9%95%bf%e8%ae%ba%e6%96%87%e4%b9%8b%e4%ba%8c.md">005 精读2017年EMNLP最佳长论文之二.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="006 精读2017年EMNLP最佳短论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/006%20%e7%b2%be%e8%af%bb2017%e5%b9%b4EMNLP%e6%9c%80%e4%bd%b3%e7%9f%ad%e8%ae%ba%e6%96%87.md">006 精读2017年EMNLP最佳短论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="007 精读2017年ICCV最佳研究论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/007%20%e7%b2%be%e8%af%bb2017%e5%b9%b4ICCV%e6%9c%80%e4%bd%b3%e7%a0%94%e7%a9%b6%e8%ae%ba%e6%96%87.md">007 精读2017年ICCV最佳研究论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="008 精读2017年ICCV最佳学生论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/008%20%e7%b2%be%e8%af%bb2017%e5%b9%b4ICCV%e6%9c%80%e4%bd%b3%e5%ad%a6%e7%94%9f%e8%ae%ba%e6%96%87.md">008 精读2017年ICCV最佳学生论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="009 如何将深度强化学习应用到视觉问答系统？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/009%20%e5%a6%82%e4%bd%95%e5%b0%86%e6%b7%b1%e5%ba%a6%e5%bc%ba%e5%8c%96%e5%ad%a6%e4%b9%a0%e5%ba%94%e7%94%a8%e5%88%b0%e8%a7%86%e8%a7%89%e9%97%ae%e7%ad%94%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">009 如何将深度强化学习应用到视觉问答系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="010 精读2017年NIPS最佳研究论文之一：如何解决非凸优化问题？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/010%20%e7%b2%be%e8%af%bb2017%e5%b9%b4NIPS%e6%9c%80%e4%bd%b3%e7%a0%94%e7%a9%b6%e8%ae%ba%e6%96%87%e4%b9%8b%e4%b8%80%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e9%9d%9e%e5%87%b8%e4%bc%98%e5%8c%96%e9%97%ae%e9%a2%98%ef%bc%9f.md">010 精读2017年NIPS最佳研究论文之一：如何解决非凸优化问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="011 精读2017年NIPS最佳研究论文之二：KSD测试如何检验两个分布的异同？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/011%20%e7%b2%be%e8%af%bb2017%e5%b9%b4NIPS%e6%9c%80%e4%bd%b3%e7%a0%94%e7%a9%b6%e8%ae%ba%e6%96%87%e4%b9%8b%e4%ba%8c%ef%bc%9aKSD%e6%b5%8b%e8%af%95%e5%a6%82%e4%bd%95%e6%a3%80%e9%aa%8c%e4%b8%a4%e4%b8%aa%e5%88%86%e5%b8%83%e7%9a%84%e5%bc%82%e5%90%8c%ef%bc%9f.md">011 精读2017年NIPS最佳研究论文之二：KSD测试如何检验两个分布的异同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="012 精读2017年NIPS最佳研究论文之三：如何解决非完美信息博弈问题？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/012%20%e7%b2%be%e8%af%bb2017%e5%b9%b4NIPS%e6%9c%80%e4%bd%b3%e7%a0%94%e7%a9%b6%e8%ae%ba%e6%96%87%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e9%9d%9e%e5%ae%8c%e7%be%8e%e4%bf%a1%e6%81%af%e5%8d%9a%e5%bc%88%e9%97%ae%e9%a2%98%ef%bc%9f.md">012 精读2017年NIPS最佳研究论文之三：如何解决非完美信息博弈问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="013 WSDM 2018论文精读：看谷歌团队如何做位置偏差估计.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/013%20WSDM%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e7%9c%8b%e8%b0%b7%e6%ad%8c%e5%9b%a2%e9%98%9f%e5%a6%82%e4%bd%95%e5%81%9a%e4%bd%8d%e7%bd%ae%e5%81%8f%e5%b7%ae%e4%bc%b0%e8%ae%a1.md">013 WSDM 2018论文精读：看谷歌团队如何做位置偏差估计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="014 WSDM 2018论文精读：看京东团队如何挖掘商品的替代信息和互补信息.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/014%20WSDM%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e7%9c%8b%e4%ba%ac%e4%b8%9c%e5%9b%a2%e9%98%9f%e5%a6%82%e4%bd%95%e6%8c%96%e6%8e%98%e5%95%86%e5%93%81%e7%9a%84%e6%9b%bf%e4%bb%a3%e4%bf%a1%e6%81%af%e5%92%8c%e4%ba%92%e8%a1%a5%e4%bf%a1%e6%81%af.md">014 WSDM 2018论文精读：看京东团队如何挖掘商品的替代信息和互补信息.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="015 WSDM 2018论文精读：深度学习模型中如何使用上下文信息？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/015%20WSDM%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b%e4%b8%ad%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%e4%b8%8a%e4%b8%8b%e6%96%87%e4%bf%a1%e6%81%af%ef%bc%9f.md">015 WSDM 2018论文精读：深度学习模型中如何使用上下文信息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="016 The Web 2018论文精读：如何对商品的图片美感进行建模？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/016%20The%20Web%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%b9%e5%95%86%e5%93%81%e7%9a%84%e5%9b%be%e7%89%87%e7%be%8e%e6%84%9f%e8%bf%9b%e8%a1%8c%e5%bb%ba%e6%a8%a1%ef%bc%9f.md">016 The Web 2018论文精读：如何对商品的图片美感进行建模？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="017 The Web 2018论文精读：如何改进经典的推荐算法BPR？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/017%20The%20Web%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e6%94%b9%e8%bf%9b%e7%bb%8f%e5%85%b8%e7%9a%84%e6%8e%a8%e8%8d%90%e7%ae%97%e6%b3%95BPR%ef%bc%9f.md">017 The Web 2018论文精读：如何改进经典的推荐算法BPR？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="018 The Web 2018论文精读：如何从文本中提取高元关系？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/018%20The%20Web%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e6%96%87%e6%9c%ac%e4%b8%ad%e6%8f%90%e5%8f%96%e9%ab%98%e5%85%83%e5%85%b3%e7%b3%bb%ef%bc%9f.md">018 The Web 2018论文精读：如何从文本中提取高元关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="019 SIGIR 2018论文精读：偏差和流行度之间的关系.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/019%20SIGIR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%81%8f%e5%b7%ae%e5%92%8c%e6%b5%81%e8%a1%8c%e5%ba%a6%e4%b9%8b%e9%97%b4%e7%9a%84%e5%85%b3%e7%b3%bb.md">019 SIGIR 2018论文精读：偏差和流行度之间的关系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="020 SIGIR 2018论文精读：如何利用对抗学习来增强排序模型的普适性？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/020%20SIGIR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e5%af%b9%e6%8a%97%e5%ad%a6%e4%b9%a0%e6%9d%a5%e5%a2%9e%e5%bc%ba%e6%8e%92%e5%ba%8f%e6%a8%a1%e5%9e%8b%e7%9a%84%e6%99%ae%e9%80%82%e6%80%a7%ef%bc%9f.md">020 SIGIR 2018论文精读：如何利用对抗学习来增强排序模型的普适性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="021 SIGIR 2018论文精读：如何对搜索页面上的点击行为进行序列建模？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/021%20SIGIR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%b9%e6%90%9c%e7%b4%a2%e9%a1%b5%e9%9d%a2%e4%b8%8a%e7%9a%84%e7%82%b9%e5%87%bb%e8%a1%8c%e4%b8%ba%e8%bf%9b%e8%a1%8c%e5%ba%8f%e5%88%97%e5%bb%ba%e6%a8%a1%ef%bc%9f.md">021 SIGIR 2018论文精读：如何对搜索页面上的点击行为进行序列建模？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="022 CVPR 2018论文精读：如何研究计算机视觉任务之间的关系？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/022%20CVPR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e7%a0%94%e7%a9%b6%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e4%bb%bb%e5%8a%a1%e4%b9%8b%e9%97%b4%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">022 CVPR 2018论文精读：如何研究计算机视觉任务之间的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="023 CVPR 2018论文精读：如何从整体上对人体进行三维建模？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/023%20CVPR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e6%95%b4%e4%bd%93%e4%b8%8a%e5%af%b9%e4%ba%ba%e4%bd%93%e8%bf%9b%e8%a1%8c%e4%b8%89%e7%bb%b4%e5%bb%ba%e6%a8%a1%ef%bc%9f.md">023 CVPR 2018论文精读：如何从整体上对人体进行三维建模？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="024 CVPR 2018论文精读：如何解决排序学习计算复杂度高这个问题？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/024%20CVPR%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e6%8e%92%e5%ba%8f%e5%ad%a6%e4%b9%a0%e8%ae%a1%e7%ae%97%e5%a4%8d%e6%9d%82%e5%ba%a6%e9%ab%98%e8%bf%99%e4%b8%aa%e9%97%ae%e9%a2%98%ef%bc%9f.md">024 CVPR 2018论文精读：如何解决排序学习计算复杂度高这个问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="025 ICML 2018论文精读：模型经得起对抗样本的攻击？这或许只是个错觉.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/025%20ICML%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e6%a8%a1%e5%9e%8b%e7%bb%8f%e5%be%97%e8%b5%b7%e5%af%b9%e6%8a%97%e6%a0%b7%e6%9c%ac%e7%9a%84%e6%94%bb%e5%87%bb%ef%bc%9f%e8%bf%99%e6%88%96%e8%ae%b8%e5%8f%aa%e6%98%af%e4%b8%aa%e9%94%99%e8%a7%89.md">025 ICML 2018论文精读：模型经得起对抗样本的攻击？这或许只是个错觉.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="026 ICML 2018论文精读：聊一聊机器学习算法的公平性问题.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/026%20ICML%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e8%81%8a%e4%b8%80%e8%81%8a%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e7%ae%97%e6%b3%95%e7%9a%84%e5%85%ac%e5%b9%b3%e6%80%a7%e9%97%ae%e9%a2%98.md">026 ICML 2018论文精读：聊一聊机器学习算法的公平性问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="027 ICML 2018论文精读：优化目标函数的时候，有可能放大了不公平？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/027%20ICML%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e4%bc%98%e5%8c%96%e7%9b%ae%e6%a0%87%e5%87%bd%e6%95%b0%e7%9a%84%e6%97%b6%e5%80%99%ef%bc%8c%e6%9c%89%e5%8f%af%e8%83%bd%e6%94%be%e5%a4%a7%e4%ba%86%e4%b8%8d%e5%85%ac%e5%b9%b3%ef%bc%9f.md">027 ICML 2018论文精读：优化目标函数的时候，有可能放大了不公平？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="028 ACL 2018论文精读：问答系统场景下，如何提出好问题？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/028%20ACL%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e9%97%ae%e7%ad%94%e7%b3%bb%e7%bb%9f%e5%9c%ba%e6%99%af%e4%b8%8b%ef%bc%8c%e5%a6%82%e4%bd%95%e6%8f%90%e5%87%ba%e5%a5%bd%e9%97%ae%e9%a2%98%ef%bc%9f.md">028 ACL 2018论文精读：问答系统场景下，如何提出好问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="029 ACL 2018论文精读：什么是对话中的前提触发？如何检测？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/029%20ACL%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%af%e5%af%b9%e8%af%9d%e4%b8%ad%e7%9a%84%e5%89%8d%e6%8f%90%e8%a7%a6%e5%8f%91%ef%bc%9f%e5%a6%82%e4%bd%95%e6%a3%80%e6%b5%8b%ef%bc%9f.md">029 ACL 2018论文精读：什么是对话中的前提触发？如何检测？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="030 ACL 2018论文精读：什么是端到端的语义哈希？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/030%20ACL%202018%e8%ae%ba%e6%96%87%e7%b2%be%e8%af%bb%ef%bc%9a%e4%bb%80%e4%b9%88%e6%98%af%e7%ab%af%e5%88%b0%e7%ab%af%e7%9a%84%e8%af%ad%e4%b9%89%e5%93%88%e5%b8%8c%ef%bc%9f.md">030 ACL 2018论文精读：什么是端到端的语义哈希？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="030 复盘 7 一起来读人工智能国际顶级会议论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/030%20%e5%a4%8d%e7%9b%98%207%20%e4%b8%80%e8%b5%b7%e6%9d%a5%e8%af%bb%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e5%9b%bd%e9%99%85%e9%a1%b6%e7%ba%a7%e4%bc%9a%e8%ae%ae%e8%ae%ba%e6%96%87.md">030 复盘 7 一起来读人工智能国际顶级会议论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="031 经典搜索核心算法：TF-IDF及其变种.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/031%20%e7%bb%8f%e5%85%b8%e6%90%9c%e7%b4%a2%e6%a0%b8%e5%bf%83%e7%ae%97%e6%b3%95%ef%bc%9aTF-IDF%e5%8f%8a%e5%85%b6%e5%8f%98%e7%a7%8d.md">031 经典搜索核心算法：TF-IDF及其变种.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="032 经典搜索核心算法：BM25及其变种（内附全年目录）.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/032%20%e7%bb%8f%e5%85%b8%e6%90%9c%e7%b4%a2%e6%a0%b8%e5%bf%83%e7%ae%97%e6%b3%95%ef%bc%9aBM25%e5%8f%8a%e5%85%b6%e5%8f%98%e7%a7%8d%ef%bc%88%e5%86%85%e9%99%84%e5%85%a8%e5%b9%b4%e7%9b%ae%e5%bd%95%ef%bc%89.md">032 经典搜索核心算法：BM25及其变种（内附全年目录）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="033 经典搜索核心算法：语言模型及其变种.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/033%20%e7%bb%8f%e5%85%b8%e6%90%9c%e7%b4%a2%e6%a0%b8%e5%bf%83%e7%ae%97%e6%b3%95%ef%bc%9a%e8%af%ad%e8%a8%80%e6%a8%a1%e5%9e%8b%e5%8f%8a%e5%85%b6%e5%8f%98%e7%a7%8d.md">033 经典搜索核心算法：语言模型及其变种.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="034 机器学习排序算法：单点法排序学习.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/034%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%ef%bc%9a%e5%8d%95%e7%82%b9%e6%b3%95%e6%8e%92%e5%ba%8f%e5%ad%a6%e4%b9%a0.md">034 机器学习排序算法：单点法排序学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="035 机器学习排序算法：配对法排序学习.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/035%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%ef%bc%9a%e9%85%8d%e5%af%b9%e6%b3%95%e6%8e%92%e5%ba%8f%e5%ad%a6%e4%b9%a0.md">035 机器学习排序算法：配对法排序学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="036 机器学习排序算法：列表法排序学习.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/036%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%ef%bc%9a%e5%88%97%e8%a1%a8%e6%b3%95%e6%8e%92%e5%ba%8f%e5%ad%a6%e4%b9%a0.md">036 机器学习排序算法：列表法排序学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="037 查询关键字理解三部曲之分类.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/037%20%e6%9f%a5%e8%af%a2%e5%85%b3%e9%94%ae%e5%ad%97%e7%90%86%e8%a7%a3%e4%b8%89%e9%83%a8%e6%9b%b2%e4%b9%8b%e5%88%86%e7%b1%bb.md">037 查询关键字理解三部曲之分类.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="038 查询关键字理解三部曲之解析.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/038%20%e6%9f%a5%e8%af%a2%e5%85%b3%e9%94%ae%e5%ad%97%e7%90%86%e8%a7%a3%e4%b8%89%e9%83%a8%e6%9b%b2%e4%b9%8b%e8%a7%a3%e6%9e%90.md">038 查询关键字理解三部曲之解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="039 查询关键字理解三部曲之扩展.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/039%20%e6%9f%a5%e8%af%a2%e5%85%b3%e9%94%ae%e5%ad%97%e7%90%86%e8%a7%a3%e4%b8%89%e9%83%a8%e6%9b%b2%e4%b9%8b%e6%89%a9%e5%b1%95.md">039 查询关键字理解三部曲之扩展.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="040 搜索系统评测，有哪些基础指标？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/040%20%e6%90%9c%e7%b4%a2%e7%b3%bb%e7%bb%9f%e8%af%84%e6%b5%8b%ef%bc%8c%e6%9c%89%e5%93%aa%e4%ba%9b%e5%9f%ba%e7%a1%80%e6%8c%87%e6%a0%87%ef%bc%9f.md">040 搜索系统评测，有哪些基础指标？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="041 搜索系统评测，有哪些高级指标？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/041%20%e6%90%9c%e7%b4%a2%e7%b3%bb%e7%bb%9f%e8%af%84%e6%b5%8b%ef%bc%8c%e6%9c%89%e5%93%aa%e4%ba%9b%e9%ab%98%e7%ba%a7%e6%8c%87%e6%a0%87%ef%bc%9f.md">041 搜索系统评测，有哪些高级指标？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="042 如何评测搜索系统的在线表现？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/042%20%e5%a6%82%e4%bd%95%e8%af%84%e6%b5%8b%e6%90%9c%e7%b4%a2%e7%b3%bb%e7%bb%9f%e7%9a%84%e5%9c%a8%e7%ba%bf%e8%a1%a8%e7%8e%b0%ef%bc%9f.md">042 如何评测搜索系统的在线表现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="043 文档理解第一步：文档分类.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/043%20%e6%96%87%e6%a1%a3%e7%90%86%e8%a7%a3%e7%ac%ac%e4%b8%80%e6%ad%a5%ef%bc%9a%e6%96%87%e6%a1%a3%e5%88%86%e7%b1%bb.md">043 文档理解第一步：文档分类.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="044 文档理解的关键步骤：文档聚类.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/044%20%e6%96%87%e6%a1%a3%e7%90%86%e8%a7%a3%e7%9a%84%e5%85%b3%e9%94%ae%e6%ad%a5%e9%aa%a4%ef%bc%9a%e6%96%87%e6%a1%a3%e8%81%9a%e7%b1%bb.md">044 文档理解的关键步骤：文档聚类.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="045 文档理解的重要特例：多模文档建模.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/045%20%e6%96%87%e6%a1%a3%e7%90%86%e8%a7%a3%e7%9a%84%e9%87%8d%e8%a6%81%e7%89%b9%e4%be%8b%ef%bc%9a%e5%a4%9a%e6%a8%a1%e6%96%87%e6%a1%a3%e5%bb%ba%e6%a8%a1.md">045 文档理解的重要特例：多模文档建模.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="046 大型搜索框架宏观视角：发展、特点及趋势.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/046%20%e5%a4%a7%e5%9e%8b%e6%90%9c%e7%b4%a2%e6%a1%86%e6%9e%b6%e5%ae%8f%e8%a7%82%e8%a7%86%e8%a7%92%ef%bc%9a%e5%8f%91%e5%b1%95%e3%80%81%e7%89%b9%e7%82%b9%e5%8f%8a%e8%b6%8b%e5%8a%bf.md">046 大型搜索框架宏观视角：发展、特点及趋势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="047 多轮打分系统概述.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/047%20%e5%a4%9a%e8%bd%ae%e6%89%93%e5%88%86%e7%b3%bb%e7%bb%9f%e6%a6%82%e8%bf%b0.md">047 多轮打分系统概述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="048 搜索索引及其相关技术概述.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/048%20%e6%90%9c%e7%b4%a2%e7%b4%a2%e5%bc%95%e5%8f%8a%e5%85%b6%e7%9b%b8%e5%85%b3%e6%8a%80%e6%9c%af%e6%a6%82%e8%bf%b0.md">048 搜索索引及其相关技术概述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="049 PageRank算法的核心思想是什么？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/049%20PageRank%e7%ae%97%e6%b3%95%e7%9a%84%e6%a0%b8%e5%bf%83%e6%80%9d%e6%83%b3%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">049 PageRank算法的核心思想是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="050 经典图算法之HITS.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/050%20%e7%bb%8f%e5%85%b8%e5%9b%be%e7%ae%97%e6%b3%95%e4%b9%8bHITS.md">050 经典图算法之HITS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="051 社区检测算法之模块最大化" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/051%20%e7%a4%be%e5%8c%ba%e6%a3%80%e6%b5%8b%e7%ae%97%e6%b3%95%e4%b9%8b%e6%a8%a1%e5%9d%97%e6%9c%80%e5%a4%a7%e5%8c%96">051 社区检测算法之模块最大化</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="052 机器学习排序算法经典模型：RankSVM.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/052%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%e7%bb%8f%e5%85%b8%e6%a8%a1%e5%9e%8b%ef%bc%9aRankSVM.md">052 机器学习排序算法经典模型：RankSVM.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="053 机器学习排序算法经典模型：GBDT.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/053%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%e7%bb%8f%e5%85%b8%e6%a8%a1%e5%9e%8b%ef%bc%9aGBDT.md">053 机器学习排序算法经典模型：GBDT.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="054 机器学习排序算法经典模型：LambdaMART.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/054%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8e%92%e5%ba%8f%e7%ae%97%e6%b3%95%e7%bb%8f%e5%85%b8%e6%a8%a1%e5%9e%8b%ef%bc%9aLambdaMART.md">054 机器学习排序算法经典模型：LambdaMART.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="055 基于深度学习的搜索算法：深度结构化语义模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/055%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%90%9c%e7%b4%a2%e7%ae%97%e6%b3%95%ef%bc%9a%e6%b7%b1%e5%ba%a6%e7%bb%93%e6%9e%84%e5%8c%96%e8%af%ad%e4%b9%89%e6%a8%a1%e5%9e%8b.md">055 基于深度学习的搜索算法：深度结构化语义模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="056 基于深度学习的搜索算法：卷积结构下的隐含语义模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/056%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%90%9c%e7%b4%a2%e7%ae%97%e6%b3%95%ef%bc%9a%e5%8d%b7%e7%a7%af%e7%bb%93%e6%9e%84%e4%b8%8b%e7%9a%84%e9%9a%90%e5%90%ab%e8%af%ad%e4%b9%89%e6%a8%a1%e5%9e%8b.md">056 基于深度学习的搜索算法：卷积结构下的隐含语义模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="057 基于深度学习的搜索算法：局部和分布表征下的搜索模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/057%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%90%9c%e7%b4%a2%e7%ae%97%e6%b3%95%ef%bc%9a%e5%b1%80%e9%83%a8%e5%92%8c%e5%88%86%e5%b8%83%e8%a1%a8%e5%be%81%e4%b8%8b%e7%9a%84%e6%90%9c%e7%b4%a2%e6%a8%a1%e5%9e%8b.md">057 基于深度学习的搜索算法：局部和分布表征下的搜索模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="057 复盘 1 搜索核心技术模块.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/057%20%e5%a4%8d%e7%9b%98%201%20%e6%90%9c%e7%b4%a2%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e6%a8%a1%e5%9d%97.md">057 复盘 1 搜索核心技术模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="058 简单推荐模型之一：基于流行度的推荐模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/058%20%e7%ae%80%e5%8d%95%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e5%9f%ba%e4%ba%8e%e6%b5%81%e8%a1%8c%e5%ba%a6%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b.md">058 简单推荐模型之一：基于流行度的推荐模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="059 简单推荐模型之二：基于相似信息的推荐模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/059%20%e7%ae%80%e5%8d%95%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e5%9f%ba%e4%ba%8e%e7%9b%b8%e4%bc%bc%e4%bf%a1%e6%81%af%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b.md">059 简单推荐模型之二：基于相似信息的推荐模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="060 简单推荐模型之三：基于内容信息的推荐模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/060%20%e7%ae%80%e5%8d%95%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%9f%ba%e4%ba%8e%e5%86%85%e5%ae%b9%e4%bf%a1%e6%81%af%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b.md">060 简单推荐模型之三：基于内容信息的推荐模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="061 基于隐变量的模型之一：矩阵分解.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/061%20%e5%9f%ba%e4%ba%8e%e9%9a%90%e5%8f%98%e9%87%8f%e7%9a%84%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e7%9f%a9%e9%98%b5%e5%88%86%e8%a7%a3.md">061 基于隐变量的模型之一：矩阵分解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="062 基于隐变量的模型之二：基于回归的矩阵分解.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/062%20%e5%9f%ba%e4%ba%8e%e9%9a%90%e5%8f%98%e9%87%8f%e7%9a%84%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e5%9f%ba%e4%ba%8e%e5%9b%9e%e5%bd%92%e7%9a%84%e7%9f%a9%e9%98%b5%e5%88%86%e8%a7%a3.md">062 基于隐变量的模型之二：基于回归的矩阵分解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="063 基于隐变量的模型之三：分解机.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/063%20%e5%9f%ba%e4%ba%8e%e9%9a%90%e5%8f%98%e9%87%8f%e7%9a%84%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%88%86%e8%a7%a3%e6%9c%ba.md">063 基于隐变量的模型之三：分解机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="064 高级推荐模型之一：张量分解模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/064%20%e9%ab%98%e7%ba%a7%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e5%bc%a0%e9%87%8f%e5%88%86%e8%a7%a3%e6%a8%a1%e5%9e%8b.md">064 高级推荐模型之一：张量分解模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="065 高级推荐模型之二：协同矩阵分解.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/065%20%e9%ab%98%e7%ba%a7%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e5%8d%8f%e5%90%8c%e7%9f%a9%e9%98%b5%e5%88%86%e8%a7%a3.md">065 高级推荐模型之二：协同矩阵分解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="066 高级推荐模型之三：优化复杂目标函数.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/066%20%e9%ab%98%e7%ba%a7%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%89%ef%bc%9a%e4%bc%98%e5%8c%96%e5%a4%8d%e6%9d%82%e7%9b%ae%e6%a0%87%e5%87%bd%e6%95%b0.md">066 高级推荐模型之三：优化复杂目标函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="067 推荐的Exploit和Explore算法之一：EE算法综述.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/067%20%e6%8e%a8%e8%8d%90%e7%9a%84Exploit%e5%92%8cExplore%e7%ae%97%e6%b3%95%e4%b9%8b%e4%b8%80%ef%bc%9aEE%e7%ae%97%e6%b3%95%e7%bb%bc%e8%bf%b0.md">067 推荐的Exploit和Explore算法之一：EE算法综述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="068 推荐的Exploit和Explore算法之二：UCB算法.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/068%20%e6%8e%a8%e8%8d%90%e7%9a%84Exploit%e5%92%8cExplore%e7%ae%97%e6%b3%95%e4%b9%8b%e4%ba%8c%ef%bc%9aUCB%e7%ae%97%e6%b3%95.md">068 推荐的Exploit和Explore算法之二：UCB算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="069 推荐的Exploit和Explore算法之三：汤普森采样算法.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/069%20%e6%8e%a8%e8%8d%90%e7%9a%84Exploit%e5%92%8cExplore%e7%ae%97%e6%b3%95%e4%b9%8b%e4%b8%89%ef%bc%9a%e6%b1%a4%e6%99%ae%e6%a3%ae%e9%87%87%e6%a0%b7%e7%ae%97%e6%b3%95.md">069 推荐的Exploit和Explore算法之三：汤普森采样算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="070 推荐系统评测之一：传统线下评测.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/070%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e8%af%84%e6%b5%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e4%bc%a0%e7%bb%9f%e7%ba%bf%e4%b8%8b%e8%af%84%e6%b5%8b.md">070 推荐系统评测之一：传统线下评测.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="071 推荐系统评测之二：线上评测.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/071%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e8%af%84%e6%b5%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e7%ba%bf%e4%b8%8a%e8%af%84%e6%b5%8b.md">071 推荐系统评测之二：线上评测.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="072 推荐系统评测之三：无偏差估计.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/072%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e8%af%84%e6%b5%8b%e4%b9%8b%e4%b8%89%ef%bc%9a%e6%97%a0%e5%81%8f%e5%b7%ae%e4%bc%b0%e8%ae%a1.md">072 推荐系统评测之三：无偏差估计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="073 现代推荐架构剖析之一：基于线下离线计算的推荐架构.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/073%20%e7%8e%b0%e4%bb%a3%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84%e5%89%96%e6%9e%90%e4%b9%8b%e4%b8%80%ef%bc%9a%e5%9f%ba%e4%ba%8e%e7%ba%bf%e4%b8%8b%e7%a6%bb%e7%ba%bf%e8%ae%a1%e7%ae%97%e7%9a%84%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84.md">073 现代推荐架构剖析之一：基于线下离线计算的推荐架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="074 现代推荐架构剖析之二：基于多层搜索架构的推荐系统.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/074%20%e7%8e%b0%e4%bb%a3%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84%e5%89%96%e6%9e%90%e4%b9%8b%e4%ba%8c%ef%bc%9a%e5%9f%ba%e4%ba%8e%e5%a4%9a%e5%b1%82%e6%90%9c%e7%b4%a2%e6%9e%b6%e6%9e%84%e7%9a%84%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f.md">074 现代推荐架构剖析之二：基于多层搜索架构的推荐系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="075 现代推荐架构剖析之三：复杂现代推荐架构漫谈.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/075%20%e7%8e%b0%e4%bb%a3%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84%e5%89%96%e6%9e%90%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%a4%8d%e6%9d%82%e7%8e%b0%e4%bb%a3%e6%8e%a8%e8%8d%90%e6%9e%b6%e6%9e%84%e6%bc%ab%e8%b0%88.md">075 现代推荐架构剖析之三：复杂现代推荐架构漫谈.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="076 基于深度学习的推荐模型之一：受限波兹曼机.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/076%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e5%8f%97%e9%99%90%e6%b3%a2%e5%85%b9%e6%9b%bc%e6%9c%ba.md">076 基于深度学习的推荐模型之一：受限波兹曼机.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="077 基于深度学习的推荐模型之二：基于RNN的推荐系统.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/077%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e5%9f%ba%e4%ba%8eRNN%e7%9a%84%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f.md">077 基于深度学习的推荐模型之二：基于RNN的推荐系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="078 基于深度学习的推荐模型之三：利用深度学习来扩展推荐系统.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/078%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e6%8e%a8%e8%8d%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%88%a9%e7%94%a8%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%9d%a5%e6%89%a9%e5%b1%95%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f.md">078 基于深度学习的推荐模型之三：利用深度学习来扩展推荐系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="078 复盘 2 推荐系统核心技术模块.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/078%20%e5%a4%8d%e7%9b%98%202%20%e6%8e%a8%e8%8d%90%e7%b3%bb%e7%bb%9f%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e6%a8%a1%e5%9d%97.md">078 复盘 2 推荐系统核心技术模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="079 广告系统概述.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/079%20%e5%b9%bf%e5%91%8a%e7%b3%bb%e7%bb%9f%e6%a6%82%e8%bf%b0.md">079 广告系统概述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="080 广告系统架构.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/080%20%e5%b9%bf%e5%91%8a%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84.md">080 广告系统架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="081 广告回馈预估综述.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/081%20%e5%b9%bf%e5%91%8a%e5%9b%9e%e9%a6%88%e9%a2%84%e4%bc%b0%e7%bb%bc%e8%bf%b0.md">081 广告回馈预估综述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="082 Google的点击率系统模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/082%20Google%e7%9a%84%e7%82%b9%e5%87%bb%e7%8e%87%e7%b3%bb%e7%bb%9f%e6%a8%a1%e5%9e%8b.md">082 Google的点击率系统模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="083 Facebook的广告点击率预估模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/083%20Facebook%e7%9a%84%e5%b9%bf%e5%91%8a%e7%82%b9%e5%87%bb%e7%8e%87%e9%a2%84%e4%bc%b0%e6%a8%a1%e5%9e%8b.md">083 Facebook的广告点击率预估模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="084 雅虎的广告点击率预估模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/084%20%e9%9b%85%e8%99%8e%e7%9a%84%e5%b9%bf%e5%91%8a%e7%82%b9%e5%87%bb%e7%8e%87%e9%a2%84%e4%bc%b0%e6%a8%a1%e5%9e%8b.md">084 雅虎的广告点击率预估模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="085 LinkedIn的广告点击率预估模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/085%20LinkedIn%e7%9a%84%e5%b9%bf%e5%91%8a%e7%82%b9%e5%87%bb%e7%8e%87%e9%a2%84%e4%bc%b0%e6%a8%a1%e5%9e%8b.md">085 LinkedIn的广告点击率预估模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="086 Twitter的广告点击率预估模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/086%20Twitter%e7%9a%84%e5%b9%bf%e5%91%8a%e7%82%b9%e5%87%bb%e7%8e%87%e9%a2%84%e4%bc%b0%e6%a8%a1%e5%9e%8b.md">086 Twitter的广告点击率预估模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="087 阿里巴巴的广告点击率预估模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/087%20%e9%98%bf%e9%87%8c%e5%b7%b4%e5%b7%b4%e7%9a%84%e5%b9%bf%e5%91%8a%e7%82%b9%e5%87%bb%e7%8e%87%e9%a2%84%e4%bc%b0%e6%a8%a1%e5%9e%8b.md">087 阿里巴巴的广告点击率预估模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="088 什么是基于第二价位的广告竞拍？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/088%20%e4%bb%80%e4%b9%88%e6%98%af%e5%9f%ba%e4%ba%8e%e7%ac%ac%e4%ba%8c%e4%bb%b7%e4%bd%8d%e7%9a%84%e5%b9%bf%e5%91%8a%e7%ab%9e%e6%8b%8d%ef%bc%9f.md">088 什么是基于第二价位的广告竞拍？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="089 广告的竞价策略是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/089%20%e5%b9%bf%e5%91%8a%e7%9a%84%e7%ab%9e%e4%bb%b7%e7%ad%96%e7%95%a5%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">089 广告的竞价策略是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="090 如何优化广告的竞价策略？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/090%20%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96%e5%b9%bf%e5%91%8a%e7%9a%84%e7%ab%9e%e4%bb%b7%e7%ad%96%e7%95%a5%ef%bc%9f.md">090 如何优化广告的竞价策略？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="091 如何控制广告预算？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/091%20%e5%a6%82%e4%bd%95%e6%8e%a7%e5%88%b6%e5%b9%bf%e5%91%8a%e9%a2%84%e7%ae%97%ef%bc%9f.md">091 如何控制广告预算？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="092 如何设置广告竞价的底价？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/092%20%e5%a6%82%e4%bd%95%e8%ae%be%e7%bd%ae%e5%b9%bf%e5%91%8a%e7%ab%9e%e4%bb%b7%e7%9a%84%e5%ba%95%e4%bb%b7%ef%bc%9f.md">092 如何设置广告竞价的底价？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="093 聊一聊程序化直接购买和广告期货.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/093%20%e8%81%8a%e4%b8%80%e8%81%8a%e7%a8%8b%e5%ba%8f%e5%8c%96%e7%9b%b4%e6%8e%a5%e8%b4%ad%e4%b9%b0%e5%92%8c%e5%b9%bf%e5%91%8a%e6%9c%9f%e8%b4%a7.md">093 聊一聊程序化直接购买和广告期货.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="094 归因模型：如何来衡量广告的有效性.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/094%20%e5%bd%92%e5%9b%a0%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9d%a5%e8%a1%a1%e9%87%8f%e5%b9%bf%e5%91%8a%e7%9a%84%e6%9c%89%e6%95%88%e6%80%a7.md">094 归因模型：如何来衡量广告的有效性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="095 广告投放如何选择受众？如何扩展受众群？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/095%20%e5%b9%bf%e5%91%8a%e6%8a%95%e6%94%be%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e5%8f%97%e4%bc%97%ef%bc%9f%e5%a6%82%e4%bd%95%e6%89%a9%e5%b1%95%e5%8f%97%e4%bc%97%e7%be%a4%ef%bc%9f.md">095 广告投放如何选择受众？如何扩展受众群？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="096 复盘 4 广告系统核心技术模块.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/096%20%e5%a4%8d%e7%9b%98%204%20%e5%b9%bf%e5%91%8a%e7%b3%bb%e7%bb%9f%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e6%a8%a1%e5%9d%97.md">096 复盘 4 广告系统核心技术模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="096 如何利用机器学习技术来检测广告欺诈？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/096%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%e6%8a%80%e6%9c%af%e6%9d%a5%e6%a3%80%e6%b5%8b%e5%b9%bf%e5%91%8a%e6%ac%ba%e8%af%88%ef%bc%9f.md">096 如何利用机器学习技术来检测广告欺诈？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="097 LDA模型的前世今生.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/097%20LDA%e6%a8%a1%e5%9e%8b%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">097 LDA模型的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="098 LDA变种模型知多少.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/098%20LDA%e5%8f%98%e7%a7%8d%e6%a8%a1%e5%9e%8b%e7%9f%a5%e5%a4%9a%e5%b0%91.md">098 LDA变种模型知多少.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="099 针对大规模数据，如何优化LDA算法？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/099%20%e9%92%88%e5%af%b9%e5%a4%a7%e8%a7%84%e6%a8%a1%e6%95%b0%e6%8d%ae%ef%bc%8c%e5%a6%82%e4%bd%95%e4%bc%98%e5%8c%96LDA%e7%ae%97%e6%b3%95%ef%bc%9f.md">099 针对大规模数据，如何优化LDA算法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="100 基础文本分析模型之一：隐语义分析.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/100%20%e5%9f%ba%e7%a1%80%e6%96%87%e6%9c%ac%e5%88%86%e6%9e%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%80%ef%bc%9a%e9%9a%90%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90.md">100 基础文本分析模型之一：隐语义分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="101 基础文本分析模型之二：概率隐语义分析.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/101%20%e5%9f%ba%e7%a1%80%e6%96%87%e6%9c%ac%e5%88%86%e6%9e%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%ba%8c%ef%bc%9a%e6%a6%82%e7%8e%87%e9%9a%90%e8%af%ad%e4%b9%89%e5%88%86%e6%9e%90.md">101 基础文本分析模型之二：概率隐语义分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="102 基础文本分析模型之三：EM算法.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/102%20%e5%9f%ba%e7%a1%80%e6%96%87%e6%9c%ac%e5%88%86%e6%9e%90%e6%a8%a1%e5%9e%8b%e4%b9%8b%e4%b8%89%ef%bc%9aEM%e7%ae%97%e6%b3%95.md">102 基础文本分析模型之三：EM算法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="103 为什么需要Word2Vec算法？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/103%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81Word2Vec%e7%ae%97%e6%b3%95%ef%bc%9f.md">103 为什么需要Word2Vec算法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="104 Word2Vec算法有哪些扩展模型？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/104%20Word2Vec%e7%ae%97%e6%b3%95%e6%9c%89%e5%93%aa%e4%ba%9b%e6%89%a9%e5%b1%95%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">104 Word2Vec算法有哪些扩展模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="105 Word2Vec算法有哪些应用？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/105%20Word2Vec%e7%ae%97%e6%b3%95%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%ef%bc%9f.md">105 Word2Vec算法有哪些应用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="106 序列建模的深度学习利器：RNN基础架构.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/106%20%e5%ba%8f%e5%88%97%e5%bb%ba%e6%a8%a1%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e5%88%a9%e5%99%a8%ef%bc%9aRNN%e5%9f%ba%e7%a1%80%e6%9e%b6%e6%9e%84.md">106 序列建模的深度学习利器：RNN基础架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="107 基于门机制的RNN架构：LSTM与GRU.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/107%20%e5%9f%ba%e4%ba%8e%e9%97%a8%e6%9c%ba%e5%88%b6%e7%9a%84RNN%e6%9e%b6%e6%9e%84%ef%bc%9aLSTM%e4%b8%8eGRU.md">107 基于门机制的RNN架构：LSTM与GRU.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="108 RNN在自然语言处理中有哪些应用场景？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/108%20RNN%e5%9c%a8%e8%87%aa%e7%84%b6%e8%af%ad%e8%a8%80%e5%a4%84%e7%90%86%e4%b8%ad%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%ef%bc%9f.md">108 RNN在自然语言处理中有哪些应用场景？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="109 对话系统之经典的对话模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/109%20%e5%af%b9%e8%af%9d%e7%b3%bb%e7%bb%9f%e4%b9%8b%e7%bb%8f%e5%85%b8%e7%9a%84%e5%af%b9%e8%af%9d%e6%a8%a1%e5%9e%8b.md">109 对话系统之经典的对话模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="110 任务型对话系统有哪些技术要点？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/110%20%e4%bb%bb%e5%8a%a1%e5%9e%8b%e5%af%b9%e8%af%9d%e7%b3%bb%e7%bb%9f%e6%9c%89%e5%93%aa%e4%ba%9b%e6%8a%80%e6%9c%af%e8%a6%81%e7%82%b9%ef%bc%9f.md">110 任务型对话系统有哪些技术要点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="111 聊天机器人有哪些核心技术要点？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/111%20%e8%81%8a%e5%a4%a9%e6%9c%ba%e5%99%a8%e4%ba%ba%e6%9c%89%e5%93%aa%e4%ba%9b%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e8%a6%81%e7%82%b9%ef%bc%9f.md">111 聊天机器人有哪些核心技术要点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="112 什么是文档情感分类？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/112%20%e4%bb%80%e4%b9%88%e6%98%af%e6%96%87%e6%a1%a3%e6%83%85%e6%84%9f%e5%88%86%e7%b1%bb%ef%bc%9f.md">112 什么是文档情感分类？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="113 如何来提取情感实体和方面呢？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/113%20%e5%a6%82%e4%bd%95%e6%9d%a5%e6%8f%90%e5%8f%96%e6%83%85%e6%84%9f%e5%ae%9e%e4%bd%93%e5%92%8c%e6%96%b9%e9%9d%a2%e5%91%a2%ef%bc%9f.md">113 如何来提取情感实体和方面呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="114 复盘 3 自然语言处理及文本处理核心技术模块.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/114%20%e5%a4%8d%e7%9b%98%203%20%e8%87%aa%e7%84%b6%e8%af%ad%e8%a8%80%e5%a4%84%e7%90%86%e5%8f%8a%e6%96%87%e6%9c%ac%e5%a4%84%e7%90%86%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e6%a8%a1%e5%9d%97.md">114 复盘 3 自然语言处理及文本处理核心技术模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="114 文本情感分析中如何做意见总结和搜索？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/114%20%e6%96%87%e6%9c%ac%e6%83%85%e6%84%9f%e5%88%86%e6%9e%90%e4%b8%ad%e5%a6%82%e4%bd%95%e5%81%9a%e6%84%8f%e8%a7%81%e6%80%bb%e7%bb%93%e5%92%8c%e6%90%9c%e7%b4%a2%ef%bc%9f.md">114 文本情感分析中如何做意见总结和搜索？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="115 什么是计算机视觉？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/115%20%e4%bb%80%e4%b9%88%e6%98%af%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%ef%bc%9f.md">115 什么是计算机视觉？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="116 掌握计算机视觉任务的基础模型和操作.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/116%20%e6%8e%8c%e6%8f%a1%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e4%bb%bb%e5%8a%a1%e7%9a%84%e5%9f%ba%e7%a1%80%e6%a8%a1%e5%9e%8b%e5%92%8c%e6%93%8d%e4%bd%9c.md">116 掌握计算机视觉任务的基础模型和操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="117 计算机视觉中的特征提取难在哪里？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/117%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e4%b8%ad%e7%9a%84%e7%89%b9%e5%be%81%e6%8f%90%e5%8f%96%e9%9a%be%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">117 计算机视觉中的特征提取难在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="118 基于深度学习的计算机视觉技术（一）：深度神经网络入门.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/118%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%ba%a6%e7%a5%9e%e7%bb%8f%e7%bd%91%e7%bb%9c%e5%85%a5%e9%97%a8.md">118 基于深度学习的计算机视觉技术（一）：深度神经网络入门.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="119 基于深度学习的计算机视觉技术（二）：基本的深度学习模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/119%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e6%8a%80%e6%9c%af%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%9f%ba%e6%9c%ac%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b.md">119 基于深度学习的计算机视觉技术（二）：基本的深度学习模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="120 基于深度学习的计算机视觉技术（三）：深度学习模型的优化.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/120%20%e5%9f%ba%e4%ba%8e%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e7%9a%84%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e6%8a%80%e6%9c%af%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b%e7%9a%84%e4%bc%98%e5%8c%96.md">120 基于深度学习的计算机视觉技术（三）：深度学习模型的优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="121 计算机视觉领域的深度学习模型（一）：AlexNet.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/121%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%a2%86%e5%9f%9f%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aAlexNet.md">121 计算机视觉领域的深度学习模型（一）：AlexNet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="122 计算机视觉领域的深度学习模型（二）：VGG &amp; GoogleNet.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/122%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%a2%86%e5%9f%9f%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aVGG%20&amp;%20GoogleNet.md">122 计算机视觉领域的深度学习模型（二）：VGG &amp; GoogleNet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="123 计算机视觉领域的深度学习模型（三）：ResNet.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/123%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%a2%86%e5%9f%9f%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ad%a6%e4%b9%a0%e6%a8%a1%e5%9e%8b%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9aResNet.md">123 计算机视觉领域的深度学习模型（三）：ResNet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="124 计算机视觉高级话题（一）：图像物体识别和分割.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/124%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%ab%98%e7%ba%a7%e8%af%9d%e9%a2%98%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%9b%be%e5%83%8f%e7%89%a9%e4%bd%93%e8%af%86%e5%88%ab%e5%92%8c%e5%88%86%e5%89%b2.md">124 计算机视觉高级话题（一）：图像物体识别和分割.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="125 计算机视觉高级话题（二）：视觉问答.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/125%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%ab%98%e7%ba%a7%e8%af%9d%e9%a2%98%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e8%a7%86%e8%a7%89%e9%97%ae%e7%ad%94.md">125 计算机视觉高级话题（二）：视觉问答.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="126 计算机视觉高级话题（三）：产生式模型.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/126%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e9%ab%98%e7%ba%a7%e8%af%9d%e9%a2%98%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e4%ba%a7%e7%94%9f%e5%bc%8f%e6%a8%a1%e5%9e%8b.md">126 计算机视觉高级话题（三）：产生式模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="126复盘 5 计算机视觉核心技术模块.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/126%e5%a4%8d%e7%9b%98%205%20%e8%ae%a1%e7%ae%97%e6%9c%ba%e8%a7%86%e8%a7%89%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e6%a8%a1%e5%9d%97.md">126复盘 5 计算机视觉核心技术模块.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="127 数据科学家基础能力之概率统计.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/127%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%9f%ba%e7%a1%80%e8%83%bd%e5%8a%9b%e4%b9%8b%e6%a6%82%e7%8e%87%e7%bb%9f%e8%ae%a1.md">127 数据科学家基础能力之概率统计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="128 数据科学家基础能力之机器学习.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/128%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%9f%ba%e7%a1%80%e8%83%bd%e5%8a%9b%e4%b9%8b%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0.md">128 数据科学家基础能力之机器学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="129 数据科学家基础能力之系统.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/129%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%9f%ba%e7%a1%80%e8%83%bd%e5%8a%9b%e4%b9%8b%e7%b3%bb%e7%bb%9f.md">129 数据科学家基础能力之系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="130 数据科学家高阶能力之分析产品.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/130%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e9%ab%98%e9%98%b6%e8%83%bd%e5%8a%9b%e4%b9%8b%e5%88%86%e6%9e%90%e4%ba%a7%e5%93%81.md">130 数据科学家高阶能力之分析产品.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="131 数据科学家高阶能力之评估产品.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/131%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e9%ab%98%e9%98%b6%e8%83%bd%e5%8a%9b%e4%b9%8b%e8%af%84%e4%bc%b0%e4%ba%a7%e5%93%81.md">131 数据科学家高阶能力之评估产品.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="132 数据科学家高阶能力之如何系统提升产品性能.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/132%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e9%ab%98%e9%98%b6%e8%83%bd%e5%8a%9b%e4%b9%8b%e5%a6%82%e4%bd%95%e7%b3%bb%e7%bb%9f%e6%8f%90%e5%8d%87%e4%ba%a7%e5%93%81%e6%80%a7%e8%83%bd.md">132 数据科学家高阶能力之如何系统提升产品性能.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="133 职场话题：当数据科学家遇见产品团队.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/133%20%e8%81%8c%e5%9c%ba%e8%af%9d%e9%a2%98%ef%bc%9a%e5%bd%93%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e9%81%87%e8%a7%81%e4%ba%a7%e5%93%81%e5%9b%a2%e9%98%9f.md">133 职场话题：当数据科学家遇见产品团队.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="134 职场话题：数据科学家应聘要具备哪些能力？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/134%20%e8%81%8c%e5%9c%ba%e8%af%9d%e9%a2%98%ef%bc%9a%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%ba%94%e8%81%98%e8%a6%81%e5%85%b7%e5%a4%87%e5%93%aa%e4%ba%9b%e8%83%bd%e5%8a%9b%ef%bc%9f.md">134 职场话题：数据科学家应聘要具备哪些能力？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="135 职场话题：聊聊数据科学家的职场规划.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/135%20%e8%81%8c%e5%9c%ba%e8%af%9d%e9%a2%98%ef%bc%9a%e8%81%8a%e8%81%8a%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e7%9a%84%e8%81%8c%e5%9c%ba%e8%a7%84%e5%88%92.md">135 职场话题：聊聊数据科学家的职场规划.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="136 如何组建一个数据科学团队？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/136%20%e5%a6%82%e4%bd%95%e7%bb%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%ef%bc%9f.md">136 如何组建一个数据科学团队？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="137 数据科学团队养成：电话面试指南.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/137%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%e5%85%bb%e6%88%90%ef%bc%9a%e7%94%b5%e8%af%9d%e9%9d%a2%e8%af%95%e6%8c%87%e5%8d%97.md">137 数据科学团队养成：电话面试指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="138 数据科学团队养成：Onsite面试面面观.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/138%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%e5%85%bb%e6%88%90%ef%bc%9aOnsite%e9%9d%a2%e8%af%95%e9%9d%a2%e9%9d%a2%e8%a7%82.md">138 数据科学团队养成：Onsite面试面面观.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="139 成为香饽饽的数据科学家，如何衡量他们的工作呢？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/139%20%e6%88%90%e4%b8%ba%e9%a6%99%e9%a5%bd%e9%a5%bd%e7%9a%84%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a1%a1%e9%87%8f%e4%bb%96%e4%bb%ac%e7%9a%84%e5%b7%a5%e4%bd%9c%e5%91%a2%ef%bc%9f.md">139 成为香饽饽的数据科学家，如何衡量他们的工作呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="140 人工智能领域知识体系更新周期只有5～6年，数据科学家如何培养？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/140%20%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e9%a2%86%e5%9f%9f%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e6%9b%b4%e6%96%b0%e5%91%a8%e6%9c%9f%e5%8f%aa%e6%9c%895%ef%bd%9e6%e5%b9%b4%ef%bc%8c%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%a6%82%e4%bd%95%e5%9f%b9%e5%85%bb%ef%bc%9f.md">140 人工智能领域知识体系更新周期只有5～6年，数据科学家如何培养？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="141 数据科学家团队组织架构：水平还是垂直，这是个问题.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/141%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%9b%a2%e9%98%9f%e7%bb%84%e7%bb%87%e6%9e%b6%e6%9e%84%ef%bc%9a%e6%b0%b4%e5%b9%b3%e8%bf%98%e6%98%af%e5%9e%82%e7%9b%b4%ef%bc%8c%e8%bf%99%e6%98%af%e4%b8%aa%e9%97%ae%e9%a2%98.md">141 数据科学家团队组织架构：水平还是垂直，这是个问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="142 数据科学家必备套路之一：搜索套路.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/142%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%bf%85%e5%a4%87%e5%a5%97%e8%b7%af%e4%b9%8b%e4%b8%80%ef%bc%9a%e6%90%9c%e7%b4%a2%e5%a5%97%e8%b7%af.md">142 数据科学家必备套路之一：搜索套路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="143 数据科学家必备套路之二：推荐套路.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/143%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%bf%85%e5%a4%87%e5%a5%97%e8%b7%af%e4%b9%8b%e4%ba%8c%ef%bc%9a%e6%8e%a8%e8%8d%90%e5%a5%97%e8%b7%af.md">143 数据科学家必备套路之二：推荐套路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="144 数据科学家必备套路之三：广告套路.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/144%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e5%bf%85%e5%a4%87%e5%a5%97%e8%b7%af%e4%b9%8b%e4%b8%89%ef%bc%9a%e5%b9%bf%e5%91%8a%e5%a5%97%e8%b7%af.md">144 数据科学家必备套路之三：广告套路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="145 如何做好人工智能项目的管理？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/145%20%e5%a6%82%e4%bd%95%e5%81%9a%e5%a5%bd%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e9%a1%b9%e7%9b%ae%e7%9a%84%e7%ae%a1%e7%90%86%ef%bc%9f.md">145 如何做好人工智能项目的管理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="146 数据科学团队必备的工程流程三部曲.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/146%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%e5%bf%85%e5%a4%87%e7%9a%84%e5%b7%a5%e7%a8%8b%e6%b5%81%e7%a8%8b%e4%b8%89%e9%83%a8%e6%9b%b2.md">146 数据科学团队必备的工程流程三部曲.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="147 数据科学团队怎么选择产品和项目？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/147%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%e6%80%8e%e4%b9%88%e9%80%89%e6%8b%a9%e4%ba%a7%e5%93%81%e5%92%8c%e9%a1%b9%e7%9b%ae%ef%bc%9f.md">147 数据科学团队怎么选择产品和项目？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="148 曾经辉煌的雅虎研究院.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/148%20%e6%9b%be%e7%bb%8f%e8%be%89%e7%85%8c%e7%9a%84%e9%9b%85%e8%99%8e%e7%a0%94%e7%a9%b6%e9%99%a2.md">148 曾经辉煌的雅虎研究院.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="149 微软研究院：工业界研究机构的楷模.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/149%20%e5%be%ae%e8%bd%af%e7%a0%94%e7%a9%b6%e9%99%a2%ef%bc%9a%e5%b7%a5%e4%b8%9a%e7%95%8c%e7%a0%94%e7%a9%b6%e6%9c%ba%e6%9e%84%e7%9a%84%e6%a5%b7%e6%a8%a1.md">149 微软研究院：工业界研究机构的楷模.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="150 复盘 6 数据科学家与数据科学团队是怎么养成的？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/150%20%e5%a4%8d%e7%9b%98%206%20%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%ae%b6%e4%b8%8e%e6%95%b0%e6%8d%ae%e7%a7%91%e5%ad%a6%e5%9b%a2%e9%98%9f%e6%98%af%e6%80%8e%e4%b9%88%e5%85%bb%e6%88%90%e7%9a%84%ef%bc%9f.md">150 复盘 6 数据科学家与数据科学团队是怎么养成的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="150 聊一聊谷歌特立独行的混合型研究.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/150%20%e8%81%8a%e4%b8%80%e8%81%8a%e8%b0%b7%e6%ad%8c%e7%89%b9%e7%ab%8b%e7%8b%ac%e8%a1%8c%e7%9a%84%e6%b7%b7%e5%90%88%e5%9e%8b%e7%a0%94%e7%a9%b6.md">150 聊一聊谷歌特立独行的混合型研究.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="151 精读AlphaGo Zero论文.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/151%20%e7%b2%be%e8%af%bbAlphaGo%20Zero%e8%ae%ba%e6%96%87.md">151 精读AlphaGo Zero论文.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="152 2017人工智能技术发展盘点.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/152%202017%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e6%8a%80%e6%9c%af%e5%8f%91%e5%b1%95%e7%9b%98%e7%82%b9.md">152 2017人工智能技术发展盘点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="153 如何快速学习国际顶级学术会议的内容？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/153%20%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%ad%a6%e4%b9%a0%e5%9b%bd%e9%99%85%e9%a1%b6%e7%ba%a7%e5%ad%a6%e6%9c%af%e4%bc%9a%e8%ae%ae%e7%9a%84%e5%86%85%e5%ae%b9%ef%bc%9f.md">153 如何快速学习国际顶级学术会议的内容？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="154 在人工智能领域，如何快速找到学习的切入点？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/154%20%e5%9c%a8%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e9%a2%86%e5%9f%9f%ef%bc%8c%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e6%89%be%e5%88%b0%e5%ad%a6%e4%b9%a0%e7%9a%84%e5%88%87%e5%85%a5%e7%82%b9%ef%bc%9f.md">154 在人工智能领域，如何快速找到学习的切入点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="155 人工智能技术选择，该从哪里获得灵感？.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/155%20%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e6%8a%80%e6%9c%af%e9%80%89%e6%8b%a9%ef%bc%8c%e8%af%a5%e4%bb%8e%e5%93%aa%e9%87%8c%e8%8e%b7%e5%be%97%e7%81%b5%e6%84%9f%ef%bc%9f.md">155 人工智能技术选择，该从哪里获得灵感？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="156 内参特刊 和你聊聊每个人都关心的人工智能热点话题.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/156%20%e5%86%85%e5%8f%82%e7%89%b9%e5%88%8a%20%e5%92%8c%e4%bd%a0%e8%81%8a%e8%81%8a%e6%af%8f%e4%b8%aa%e4%ba%ba%e9%83%bd%e5%85%b3%e5%bf%83%e7%9a%84%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e7%83%ad%e7%82%b9%e8%af%9d%e9%a2%98.md">156 内参特刊 和你聊聊每个人都关心的人工智能热点话题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="156 近在咫尺，走进人工智能研究.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/156%20%e8%bf%91%e5%9c%a8%e5%92%ab%e5%b0%ba%ef%bc%8c%e8%b5%b0%e8%bf%9b%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e7%a0%94%e7%a9%b6.md">156 近在咫尺，走进人工智能研究.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 雄关漫道真如铁，而今迈步从头越.md" href="/%e4%b8%93%e6%a0%8f/AI%e6%8a%80%e6%9c%af%e5%86%85%e5%8f%82/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e9%9b%84%e5%85%b3%e6%bc%ab%e9%81%93%e7%9c%9f%e5%a6%82%e9%93%81%ef%bc%8c%e8%80%8c%e4%bb%8a%e8%bf%88%e6%ad%a5%e4%bb%8e%e5%a4%b4%e8%b6%8a.md">结束语 雄关漫道真如铁，而今迈步从头越.md</a>
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
                            <h1 id="title" data-id="050 经典图算法之HITS" class="title">050 经典图算法之HITS</h1>
                            <div><p>这周我们分享的内容是如何理解网页和网页之间的关系。周一我们介绍了用图（Graph）来表达网页与网页之间的关系并计算网页的重要性，就是经典算法PageRank。今天我来介绍一下PageRank的姊妹算法：<strong>HITS算法</strong>。</p>

<h2 id="hits的简要历史">HITS的简要历史</h2>

<p>HITS是Hypertext-Induced Topic Search算法的简称。这个算法是由康奈尔大学计算机科学教授乔·克莱恩堡（Jon Kleinberg）于1998年发明的，正好和我们周一讲的布林和佩奇发表PageRank算法是同一年。</p>

<p>这里有必要简单介绍一下乔这个人。乔于1971年出生在马萨诸塞州波士顿。1993年他毕业于康奈尔大学获得计算机科学学士学位，并于1996年从麻省理工大学获得计算机博士学位。1998的时候，乔正在位于美国西海岸硅谷地区的IBM阿尔玛登（Almaden）研究院做博士后研究。HITS的工作最早发表于1998年在旧金山举办的第九届ACM-SIAM离散算法年会上（详细论述可参阅参考文献）。</p>

<p>乔目前是美国国家工程院（National Academy of Engineering）和美国自然与人文科学院（American Academy of Arts and Sciences）院士。顺便提一下，乔的弟弟罗伯特·克莱恩堡也在康奈尔大学计算机系任教职。</p>

<h2 id="hits的基本原理">HITS的基本原理</h2>

<p>在介绍HITS算法的基本原理之前，我们首先来复习一下网页的网络结构。每一个网页都有一个“输出链接”（Outlink）的集合。输出链接指的是从当前网页出发所指向的其他页面，比如从页面A有一个链接到页面B，那么B就是A的输出链接。根据这个定义，我们来看“输入链接”（Inlink），指的就是指向当前页面的其他页面，比如页面C指向页面A，那么C就是A的输入链接。</p>

<p>要理解HITS算法，我们还需要引入一组概念：<strong>“权威”（Authority）结点</strong>和<strong>“枢纽”（Hub）结点</strong>。这两类结点到底是什么意思呢？</p>

<p>HITS给出了一种“循环”的定义：<strong>好的“权威”结点是很多“枢纽”结点的输出链接，好的“枢纽”结点则指向很多好的“权威”结点</strong>。这种循环定义我们在PageRank的定义中已经见识过了。</p>

<p>很明显，要用数学的方法来表述权威结点和枢纽结点之间的关系就必须要为每一个页面准备两个值。因为从直觉上来说，不可能有一个页面完全是权威，也不可能有一个页面完全是枢纽。绝大多数页面都在这两种角色中转换，或者说同时扮演这两类角色。</p>

<p>数学上，对于每一个页面I，<strong>我们用X来表达这个页面的“权威值”，用Y来表达这个页面的“枢纽值”</strong>。那么，一个最直观的定义，对于I的权威值X来说，它是所有I页面的输入链接的枢纽值的总和。同理，I的枢纽值是所有I页面输出链接的权威值的总和。这就是HITS算法的原始定义。</p>

<p>我们可以看到，如果I页面的输入链接的枢纽值大，说明I页面经常被一些好的“枢纽”结点链接到，那么I自身的权威性自然也就增加了。反之，如果I能够经常指向好的“权威”结点，那I自身的“枢纽”性质也就显得重要了。</p>

<p>当然，和PageRank值一样，X和Y在HITS算法里也都是事先不可知的。因此，<strong>HITS算法的重点就是要求解X和Y</strong>。如果把所有页面的X和Y都表达成向量的形式，那么HITS算法可以写成X是矩阵L的转置和Y的乘积，而Y是矩阵L和X的乘积，这里的矩阵L就是一个邻接矩阵，每一行列表达某两个页面是否相连。进行一下代数变形，我们就可以得到X其实是一个矩阵A乘以X，这里的A是L的转置乘以L。Y其实是一个矩阵B乘以Y，这里的B是L乘以L的转置。</p>

<p>于是，惊人的一点出现了，那就是HITS算法其实是需要求解矩阵A或者矩阵B的主特征向量，也就是特征值最大所对应的特征向量，用于求解X或者Y。这一点和PageRank用矩阵表达的形式不谋而和。也就是说，尽管PageRank和HITS在思路和概念上完全不同，并且在最初的定义式上南辕北辙，但是经过一番变形之后，我们能够把两者都划归为<strong>某种形式的矩阵求解特征向量的问题</strong>。</p>

<p>实际上，<strong>把图表达为矩阵，并且通过特征向量对图的一些特性进行分析是图算法中的一个重要分支</strong>（当然，我们这里说的主要是最大的值对应的特征向量，还有其他的特征向量也有含义）。既然我们已经知道了需要计算最大的特征向量，那么之前计算PageRank所使用的“乘幂法”（Power Method）在这里也是可以使用的，我们在这里就不展开了。</p>

<p>如何把HITS算法用于搜索中呢？最开始提出HITS的时候是这么使用的。</p>

<p><strong>首先，我们根据某个查询关键字构建一个“相邻图”</strong>（Neighborhood Graph）。这个图包括所有和这个查询关键字相关的页面。这里，我们可以简化为所有包含查询关键字的页面。这一步在现代搜索引擎中通过“倒排索引”（Inverted Index）就可以很容易地得到。</p>

<p><strong>有了这个相邻图以后，我们根据这个图建立邻接矩阵，然后就可以通过邻接矩阵计算这些结点的权威值和枢纽值</strong>。当计算出这两组值之后，我们就可以根据这两组值给用户展现两种网页排序的结果，分别是根据不同的假设。</p>

<p>值得注意的是，PageRank是“查询关键字无关”（Query-Independent）的算法，也就是说每个页面的PageRank值并不随着查询关键字的不同而产生不同。而HITS算法是“查询关键字相关”（Query-Dependent）的算法。从这一点来说，HITS就和PageRank有本质的不同。</p>

<h2 id="hits算法的一些特点">HITS算法的一些特点</h2>

<p>HITS算法依靠这种迭代的方法来计算权威值和枢纽值，你一定很好奇，这样的计算究竟收敛吗？是不是也需要像PageRank一样来进行特别的处理呢？</p>

<p><strong>答案是HITS一定是收敛的</strong>。这点比原始的PageRank情况要好。然而，HITS在原始的情况下，不一定收敛到唯一一组权威值和枢纽值，也就是说，解是不唯一的。因此，我们其实需要对HITS进行一部分类似于PageRank的处理，那就是让HITS的邻接矩阵里面所有的结点都能够达到其他任何结点，只是以比较小的概率。经过这样修改，HITS就能够收敛到唯一的权威值和枢纽值了。</p>

<p><strong>HITS算法的好处是为用户提供了一种全新的视角，对于同一个查询关键字，HITS提供的权威排序和枢纽排序能够帮助用户理解自己的需求</strong>。</p>

<p>当然，<strong>HITS的弱点也来自于这个依赖于查询关键字的问题</strong>。如果把所有的计算都留在用户输入查询关键字以后，并且需要在响应时间内计算出所有的权威值和枢纽值然后进行排序，这里面的计算量是很大的。所以，后来有研究者开始使用全局的网页图，提前来计算所有页面的权威值和枢纽值，然而这样做就失去了对某一个关键字的相关信息。</p>

<h2 id="小结">小结</h2>

<p>今天我为你讲了HITS算法的核心思想 。 一起来回顾下要点：第一，我们讲了HITS的一些简明历史。第二，我们讲了HITS最原始的定义和算法，并且联系PageRank，讲了两者的异同之处。第三，我们分析了HITS的一些特点。</p>

<p>最后，给你留一个思考题，有没有办法把权威值和枢纽值所对应的两个排序合并成为一个排序呢？</p>

<p><strong>参考文献</strong></p>

<ol>
<li>Jon M. Kleinberg. <em>Authoritative sources in a hyperlinked environment</em>. J. ACM 46, 5 (September 1999), 604-632，1999.</li>
</ol>

<p><strong>论文链接</strong></p>

<ul>
<li><a href="http://www.woodmann.com/searchlores/library/authoratitativesources.pdf" target="_blank">Authoritative sources in a hyperlinked environment</a></li>
</ul>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#9ef2f2f2a7aaafafaea9def9f3fff7f2b0fdf1f3" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0d3b211e9ee8fe',t:'MTczNDAwMTQzOC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
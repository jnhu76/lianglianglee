<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=37&#32;快速构建持续交付系统（四）：Ansible&#32;解决自动部署问题>
        <link rel="icon" href="/static/favicon.png">
        <title>37 快速构建持续交付系统（四）：Ansible 解决自动部署问题 </title>
        
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
                        <a class="menu-item" id="00 开篇词 量身定制你的持续交付体系.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e9%87%8f%e8%ba%ab%e5%ae%9a%e5%88%b6%e4%bd%a0%e7%9a%84%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e4%bd%93%e7%b3%bb.md">00 开篇词 量身定制你的持续交付体系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 持续交付到底有什么价值？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/01%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e5%88%b0%e5%ba%95%e6%9c%89%e4%bb%80%e4%b9%88%e4%bb%b7%e5%80%bc%ef%bc%9f.md">01 持续交付到底有什么价值？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 影响持续交付的因素有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/02%20%e5%bd%b1%e5%93%8d%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%9a%84%e5%9b%a0%e7%b4%a0%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">02 影响持续交付的因素有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 持续交付和DevOps是一对好基友.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/03%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e5%92%8cDevOps%e6%98%af%e4%b8%80%e5%af%b9%e5%a5%bd%e5%9f%ba%e5%8f%8b.md">03 持续交付和DevOps是一对好基友.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 一切的源头，代码分支策略的选择.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/04%20%e4%b8%80%e5%88%87%e7%9a%84%e6%ba%90%e5%a4%b4%ef%bc%8c%e4%bb%a3%e7%a0%81%e5%88%86%e6%94%af%e7%ad%96%e7%95%a5%e7%9a%84%e9%80%89%e6%8b%a9.md">04 一切的源头，代码分支策略的选择.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 手把手教你依赖管理.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/05%20%e6%89%8b%e6%8a%8a%e6%89%8b%e6%95%99%e4%bd%a0%e4%be%9d%e8%b5%96%e7%ae%a1%e7%90%86.md">05 手把手教你依赖管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 代码回滚，你真的理解吗？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/06%20%e4%bb%a3%e7%a0%81%e5%9b%9e%e6%bb%9a%ef%bc%8c%e4%bd%a0%e7%9c%9f%e7%9a%84%e7%90%86%e8%a7%a3%e5%90%97%ef%bc%9f.md">06 代码回滚，你真的理解吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  “两个披萨”团队的代码管理实际案例.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/07%20%c2%a0%e2%80%9c%e4%b8%a4%e4%b8%aa%e6%8a%ab%e8%90%a8%e2%80%9d%e5%9b%a2%e9%98%9f%e7%9a%84%e4%bb%a3%e7%a0%81%e7%ae%a1%e7%90%86%e5%ae%9e%e9%99%85%e6%a1%88%e4%be%8b.md">07  “两个披萨”团队的代码管理实际案例.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 测试环境要多少？从现实需求说起.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/08%20%e6%b5%8b%e8%af%95%e7%8e%af%e5%a2%83%e8%a6%81%e5%a4%9a%e5%b0%91%ef%bc%9f%e4%bb%8e%e7%8e%b0%e5%ae%9e%e9%9c%80%e6%b1%82%e8%af%b4%e8%b5%b7.md">08 测试环境要多少？从现实需求说起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 测试环境要多少？从成本与效率说起.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/09%20%e6%b5%8b%e8%af%95%e7%8e%af%e5%a2%83%e8%a6%81%e5%a4%9a%e5%b0%91%ef%bc%9f%e4%bb%8e%e6%88%90%e6%9c%ac%e4%b8%8e%e6%95%88%e7%8e%87%e8%af%b4%e8%b5%b7.md">09 测试环境要多少？从成本与效率说起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 让环境自己说话，论环境自描述的重要性.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/10%20%e8%ae%a9%e7%8e%af%e5%a2%83%e8%87%aa%e5%b7%b1%e8%af%b4%e8%af%9d%ef%bc%8c%e8%ae%ba%e7%8e%af%e5%a2%83%e8%87%aa%e6%8f%8f%e8%bf%b0%e7%9a%84%e9%87%8d%e8%a6%81%e6%80%a7.md">10 让环境自己说话，论环境自描述的重要性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 “配置”是把双刃剑，带你了解各种配置方法.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/11%20%e2%80%9c%e9%85%8d%e7%bd%ae%e2%80%9d%e6%98%af%e6%8a%8a%e5%8f%8c%e5%88%83%e5%89%91%ef%bc%8c%e5%b8%a6%e4%bd%a0%e4%ba%86%e8%a7%a3%e5%90%84%e7%a7%8d%e9%85%8d%e7%bd%ae%e6%96%b9%e6%b3%95.md">11 “配置”是把双刃剑，带你了解各种配置方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 极限挑战，如何做到分钟级搭建环境？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/12%20%e6%9e%81%e9%99%90%e6%8c%91%e6%88%98%ef%bc%8c%e5%a6%82%e4%bd%95%e5%81%9a%e5%88%b0%e5%88%86%e9%92%9f%e7%ba%a7%e6%90%ad%e5%bb%ba%e7%8e%af%e5%a2%83%ef%bc%9f.md">12 极限挑战，如何做到分钟级搭建环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 容器技术真的是环境管理的救星吗？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/13%20%e5%ae%b9%e5%99%a8%e6%8a%80%e6%9c%af%e7%9c%9f%e7%9a%84%e6%98%af%e7%8e%af%e5%a2%83%e7%ae%a1%e7%90%86%e7%9a%84%e6%95%91%e6%98%9f%e5%90%97%ef%bc%9f.md">13 容器技术真的是环境管理的救星吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 如何做到构建的提速，再提速！.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/14%20%e5%a6%82%e4%bd%95%e5%81%9a%e5%88%b0%e6%9e%84%e5%bb%ba%e7%9a%84%e6%8f%90%e9%80%9f%ef%bc%8c%e5%86%8d%e6%8f%90%e9%80%9f%ef%bc%81.md">14 如何做到构建的提速，再提速！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 构建检测，无规矩不成方圆.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/15%20%e6%9e%84%e5%bb%ba%e6%a3%80%e6%b5%8b%ef%bc%8c%e6%97%a0%e8%a7%84%e7%9f%a9%e4%b8%8d%e6%88%90%e6%96%b9%e5%9c%86.md">15 构建检测，无规矩不成方圆.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 构建资源的弹性伸缩.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/16%20%e6%9e%84%e5%bb%ba%e8%b5%84%e6%ba%90%e7%9a%84%e5%bc%b9%e6%80%a7%e4%bc%b8%e7%bc%a9.md">16 构建资源的弹性伸缩.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 容器镜像构建的那些事儿.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/17%20%e5%ae%b9%e5%99%a8%e9%95%9c%e5%83%8f%e6%9e%84%e5%bb%ba%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">17 容器镜像构建的那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 如何做好容器镜像的个性化及合规检查？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/18%20%e5%a6%82%e4%bd%95%e5%81%9a%e5%a5%bd%e5%ae%b9%e5%99%a8%e9%95%9c%e5%83%8f%e7%9a%84%e4%b8%aa%e6%80%a7%e5%8c%96%e5%8f%8a%e5%90%88%e8%a7%84%e6%a3%80%e6%9f%a5%ef%bc%9f.md">18 如何做好容器镜像的个性化及合规检查？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  发布是持续交付的最后一公里.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/19%20%20%e5%8f%91%e5%b8%83%e6%98%af%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%9a%84%e6%9c%80%e5%90%8e%e4%b8%80%e5%85%ac%e9%87%8c.md">19  发布是持续交付的最后一公里.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 Immutable！任何变更都需要发布.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/20%20Immutable%ef%bc%81%e4%bb%bb%e4%bd%95%e5%8f%98%e6%9b%b4%e9%83%bd%e9%9c%80%e8%a6%81%e5%8f%91%e5%b8%83.md">20 Immutable！任何变更都需要发布.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 发布系统一定要注意用户体验.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/21%20%e5%8f%91%e5%b8%83%e7%b3%bb%e7%bb%9f%e4%b8%80%e5%ae%9a%e8%a6%81%e6%b3%a8%e6%84%8f%e7%94%a8%e6%88%b7%e4%bd%93%e9%aa%8c.md">21 发布系统一定要注意用户体验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 发布系统的核心架构和功能设计.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/22%20%e5%8f%91%e5%b8%83%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%a0%b8%e5%bf%83%e6%9e%b6%e6%9e%84%e5%92%8c%e5%8a%9f%e8%83%bd%e8%ae%be%e8%ae%a1.md">22 发布系统的核心架构和功能设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 业务及系统架构对发布的影响.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/23%20%e4%b8%9a%e5%8a%a1%e5%8f%8a%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84%e5%af%b9%e5%8f%91%e5%b8%83%e7%9a%84%e5%bd%b1%e5%93%8d.md">23 业务及系统架构对发布的影响.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 如何利用监控保障发布质量？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/24%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e7%9b%91%e6%8e%a7%e4%bf%9d%e9%9a%9c%e5%8f%91%e5%b8%83%e8%b4%a8%e9%87%8f%ef%bc%9f.md">24 如何利用监控保障发布质量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 代码静态检查实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/25%20%e4%bb%a3%e7%a0%81%e9%9d%99%e6%80%81%e6%a3%80%e6%9f%a5%e5%ae%9e%e8%b7%b5.md">25 代码静态检查实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 越来越重要的破坏性测试.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/26%20%e8%b6%8a%e6%9d%a5%e8%b6%8a%e9%87%8d%e8%a6%81%e7%9a%84%e7%a0%b4%e5%9d%8f%e6%80%a7%e6%b5%8b%e8%af%95.md">26 越来越重要的破坏性测试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 利用Mock与回放技术助力自动化回归.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/27%20%e5%88%a9%e7%94%a8Mock%e4%b8%8e%e5%9b%9e%e6%94%be%e6%8a%80%e6%9c%af%e5%8a%a9%e5%8a%9b%e8%87%aa%e5%8a%a8%e5%8c%96%e5%9b%9e%e5%bd%92.md">27 利用Mock与回放技术助力自动化回归.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 持续交付为什么要平台化设计？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/28%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%b9%b3%e5%8f%b0%e5%8c%96%e8%ae%be%e8%ae%a1%ef%bc%9f.md">28 持续交付为什么要平台化设计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 计算资源也是交付的内容.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/29%20%e8%ae%a1%e7%ae%97%e8%b5%84%e6%ba%90%e4%b9%9f%e6%98%af%e4%ba%a4%e4%bb%98%e7%9a%84%e5%86%85%e5%ae%b9.md">29 计算资源也是交付的内容.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 持续交付中有哪些宝贵数据？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/30%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e4%b8%ad%e6%9c%89%e5%93%aa%e4%ba%9b%e5%ae%9d%e8%b4%b5%e6%95%b0%e6%8d%ae%ef%bc%9f.md">30 持续交付中有哪些宝贵数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 了解移动App的持续交付生命周期.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/31%20%e4%ba%86%e8%a7%a3%e7%a7%bb%e5%8a%a8App%e7%9a%84%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f.md">31 了解移动App的持续交付生命周期.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 细谈移动APP的交付流水线（pipeline）.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/32%20%e7%bb%86%e8%b0%88%e7%a7%bb%e5%8a%a8APP%e7%9a%84%e4%ba%a4%e4%bb%98%e6%b5%81%e6%b0%b4%e7%ba%bf%ef%bc%88pipeline%ef%bc%89.md">32 细谈移动APP的交付流水线（pipeline）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 进阶，如何进一步提升移动APP的交付效率？.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/33%20%e8%bf%9b%e9%98%b6%ef%bc%8c%e5%a6%82%e4%bd%95%e8%bf%9b%e4%b8%80%e6%ad%a5%e6%8f%90%e5%8d%87%e7%a7%bb%e5%8a%a8APP%e7%9a%84%e4%ba%a4%e4%bb%98%e6%95%88%e7%8e%87%ef%bc%9f.md">33 进阶，如何进一步提升移动APP的交付效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 快速构建持续交付系统（一）：需求分析.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/34%20%e5%bf%ab%e9%80%9f%e6%9e%84%e5%bb%ba%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e9%9c%80%e6%b1%82%e5%88%86%e6%9e%90.md">34 快速构建持续交付系统（一）：需求分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 快速构建持续交付系统（二）：GitLab 解决代码管理问题.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/35%20%e5%bf%ab%e9%80%9f%e6%9e%84%e5%bb%ba%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aGitLab%20%e8%a7%a3%e5%86%b3%e4%bb%a3%e7%a0%81%e7%ae%a1%e7%90%86%e9%97%ae%e9%a2%98.md">35 快速构建持续交付系统（二）：GitLab 解决代码管理问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 快速构建持续交付系统（三）：Jenkins 解决集成打包问题.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/36%20%e5%bf%ab%e9%80%9f%e6%9e%84%e5%bb%ba%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%b3%bb%e7%bb%9f%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9aJenkins%20%e8%a7%a3%e5%86%b3%e9%9b%86%e6%88%90%e6%89%93%e5%8c%85%e9%97%ae%e9%a2%98.md">36 快速构建持续交付系统（三）：Jenkins 解决集成打包问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 快速构建持续交付系统（四）：Ansible 解决自动部署问题.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/37%20%e5%bf%ab%e9%80%9f%e6%9e%84%e5%bb%ba%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e7%b3%bb%e7%bb%9f%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9aAnsible%20%e8%a7%a3%e5%86%b3%e8%87%aa%e5%8a%a8%e9%83%a8%e7%bd%b2%e9%97%ae%e9%a2%98.md">37 快速构建持续交付系统（四）：Ansible 解决自动部署问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="持续交付专栏特别放送 答疑解惑.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%ad%94%e7%96%91%e8%a7%a3%e6%83%91.md">持续交付专栏特别放送 答疑解惑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="持续交付专栏特别放送 高效学习指南.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%e4%b8%93%e6%a0%8f%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e9%ab%98%e6%95%88%e5%ad%a6%e4%b9%a0%e6%8c%87%e5%8d%97.md">持续交付专栏特别放送 高效学习指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 越痛苦的事，越要经常做.md" href="/%e4%b8%93%e6%a0%8f/%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%9836%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e8%b6%8a%e7%97%9b%e8%8b%a6%e7%9a%84%e4%ba%8b%ef%bc%8c%e8%b6%8a%e8%a6%81%e7%bb%8f%e5%b8%b8%e5%81%9a.md">结束语 越痛苦的事，越要经常做.md</a>
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
                            <h1 id="title" data-id="37 快速构建持续交付系统（四）：Ansible 解决自动部署问题" class="title">37 快速构建持续交付系统（四）：Ansible 解决自动部署问题</h1>
                            <div><p>今天这篇文章，已经是实践案例系列的最后一篇了。在《快速构建持续交付系统（二）：GitLab 解决配置管理问题》和《快速构建持续交付系统（三）：Jenkins 解决集成打包问题》这两篇文章中，我们已经分别基于GitLab搭建了代码管理平台、基于Jenkins搭建了集成与编译系统，并解决了这两个平台之间的联动、配合问题，从而满足了在代码平台 push 代码时，驱动集成编译系统工作的需求。</p>

<p>算下来，我们已经通过前面这两篇文章，跑完了整个持续交付体系三分之二的路程，剩下的就是解决如何利用开源工具搭建发布平台完成代码发布，跑完持续交付最后一公里的问题了。</p>

<h2 id="利用ansible完成部署">利用Ansible完成部署</h2>

<p>Ansible 是一个自动化运维管理工具，支持Linux/Windows跨平台的配置管理，任务分发等操作，可以帮我们大大减少在变更环境时所花费的时间。</p>

<p>与其他三大主流的配置管理工具Chef、Puppet、Salt相比，Ansible最大的特点在于“agentless”，即无需在目标机器装安装agent进程，即可通过SSH或者PowerShell对一个环境中的集群进行中心化的管理。</p>

<p>所以，这个“agentless”特性，可以大大减少我们配置管理平台的学习成本，尤其适合于第一次尝试使用此类配置管理工具。</p>

<p>另外，利用Ansible，我们可以完成虚拟机的初始化，以及Tomcat Java程序的发布更新。</p>

<p>现在，我们就先看看如何在我们的机器上安装Ansible，以及如何用它来搭建我们的代码发布平台。这里，我们再一起回顾下，我在第34篇文章《快速构建持续交付系统（一）：需求分析》中提到的对发布系统的需求：</p>

<blockquote>
<p>同时支持 Jar、War、Docker的生产发布，以及统一的部署标准。</p>
</blockquote>

<p>对于移动App，我们只要发布到内部测试集市即可，所以只需要在编译打包之后上传至指定地址，这个操作在Jenkins Pipeline里执行就可以了，所以本篇就不累述了。</p>

<h3 id="ansible安装"><strong>Ansible安装</strong></h3>

<p>对于Ansible环境的准备，我推荐使用pip的方式安装。</p>

<pre><code>sudo pip install Ansible
</code></pre>

<p>安装完之后, 我们可以简单测试一下：</p>

<ol>
<li><p>提交一个Ansible的Inventory文件 hosts，该文件代表要管理的目标对象：</p>

<p>$ cat hosts
[Jenkinsservers]
10.1.77.79</p></li>

<li><p>打通本机和测试机的SSH访问：</p>

<p>$ ssh-copy-id deployer@localhost</p></li>

<li><p>尝试远程访问主机 10.1.77.79：</p>

<p>$ Ansible -i hosts  all -u deployer -a &ldquo;cat /etc/hosts”</p>

<p>10.1.77.79 | SUCCESS | rc=0 &gt;&gt;
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6</p></li>
</ol>

<p>如果返回SUCCESS，则表示我们已经可以通过Ansible管理该主机了。</p>

<p>接下来，我们再看一下如何使用Ansible达到我们的发布目标吧。</p>

<h3 id="ansible使用">Ansible使用</h3>

<p>现在，我先简单介绍下，在初次接触Ansible时，你应该掌握的两个最关键的概念：Inventory和PlayBook。</p>

<ol>
<li><strong>Inventory</strong></li>
</ol>

<p>对于被Ansible管理的机器清单，我们可以通过Inventory文件，分组管理其中一些集群的机器列表分组，并为其设置不同变量。</p>

<p>比如，我们可以通过Ansible_user ，指定不同机器的Ansible用户。</p>

<pre><code>[Jenkinsservers]
10.1.77.79 Ansible_user=root
10.1.77.80 Ansible_user=deployer
[Gitlabservers]
10.1.77.77
</code></pre>

<ol>
<li><strong>PlayBook</strong></li>
</ol>

<p>PlayBook是Ansible的脚本文件，使用YAML语言编写，包含需要远程执行的核心命令、定义任务具体内容，等等。</p>

<p>我们一起看一个Ansible官方提供的一个例子吧。</p>

<pre><code>---
- hosts: webservers
 remote_user: root

 tasks:
 - name: ensure apache is at the latest version
 yum:
 name: httpd
 state: latest
 - name: write the apache config file
 template:
 src: /srv/httpd.j2
 dest: /etc/httpd.conf
- hosts: databases
 remote_user: root

 tasks:
 - name: ensure postgresql is at the latest version
 yum:
 name: postgresql
 state: latest
 - name: ensure that postgresql is started
 service:
 name: postgresql
 state: started
</code></pre>

<p>这段代码的最主要功能是，使用yum完成了Apache服务器和PostgreSQL的安装。其中，包含了编写Ansible PlayBook的三个常用模块。</p>

<ol>
<li><p>yum 调用目标机器上的包管理工具完成软件安装 。Ansible对于不同的Linux操作系统包管理进行了封装，在CentOS上相当于yum， 在Ubuntu上相当于APT。</p></li>

<li><p>Template 远程文件渲染，可以把本地机器的文件模板渲染后放到远程主机上。</p></li>

<li><p>Service 服务管理，同样封装了不同Linux操作系统实际执行的Service命令。</p></li>
</ol>

<p><strong>通常情况下，我们用脚本的方式使用Ansible，只要使用好Inventory和PlayBook这两个组件就可以了，即：使用PlayBook编写Ansible脚本，然后用Inventory维护好需要管理的机器列表</strong>。这样，就能解决90%以上使用Ansible的需求。</p>

<p>但如果你有一些更复杂的需求，比如通过代码调用Ansible，可能还要用到API组件。感兴趣的话，你可以参考Ansible的官方文档。</p>

<h3 id="使用ansible进行java应用部署">使用Ansible进行Java应用部署</h3>

<p>我先来整理下，针对Java后端服务部署的需求：</p>

<blockquote>
<p>完成Ansible的PlayBook后，在Jenkins Pipeline中调用相关的脚本，从而完成Java Tomcat应用的发布。</p>
</blockquote>

<p><strong>首先，在目标机器上安装Tomcat，并初始化。</strong></p>

<p>我们可以通过编写Ansible PlayBook完成这个操作。一个最简单的Tomcat初始化脚本只要十几行代码，但是如果我们要对Tomcat进行更复杂的配置，比如修改Tomcat的CATALINA_OPTS参数，工作量就相当大了，而且还容易出错。</p>

<p>在这种情况下，一个更简单的做法是，使用开源第三方的PlayBook的复用文件roles。你可以访问<a href="https://galaxy.ansible.com" target="_blank">https://galaxy.Ansible.com</a> ，这里有数千个第三方的roles可供使用。</p>

<p>在GitHub上搜索一下Ansible-Tomcat，并下载，就可以很方便地使用了。</p>

<p>这里，我和你一起看一个具体roles的例子：</p>

<pre><code>---
- hosts: Tomcat_server
 roles:
 - { role: Ansible-Tomcat }
</code></pre>

<p>你只需要这简单的三行代码，就可以完成Tomcat的安装，以及服务注册。与此同时，你只要添加Tomcat_default_catalina_opts参数，就可以修改CATALINA_OPTS了。</p>

<p>这样一来，Java应用所需要的Web容器就部署好了。</p>

<p><strong>然后，部署具体的业务代码</strong>。</p>

<p>这个过程就是指，把编译完后的War包推送到目标机器上的指定目录下，供Tomcat加载。</p>

<p>完成这个需求，我们只需要通过Ansible的SCP模块把War包从Jenkins推送到目标机器上即可。</p>

<p>具体的命令如下：</p>

<pre><code>- name: Copy a war file to the remote machine 
  copy:
    src: /tmp/waimai-service.war
    dest: /opt/Tomcat/webapps/waimai-service.war
</code></pre>

<p>但是，这样编写Ansible的方式会有一个问题，就是把Ansible的发布过程和Jenkins的编译耦合了起来。</p>

<p>而在上一篇文章《快速构建持续交付系统（三）：Jenkins 解决集成打包问题》中，我提到，要在编译之后，把构建产物统一上传到Nexus或者Artifactory之类的构建产物仓库中。</p>

<p>所以，<strong>此时更好的做法是直接在部署本地从仓库下载War包</strong>。这样，之后我们有独立部署或者回滚的需求时，也可以通过在Ansible的脚本中选择版本实现。当然，此处你仍旧可以使用Ansible的SCP模块复制War包，只不过是换成了在部署机上执行而已。</p>

<p><strong>最后，重启 Tomcat 服务，整个应用的部署过程就完成了</strong>。</p>

<h2 id="ansible-tower-简介">Ansible Tower 简介</h2>

<p>通过上面的几个步骤，我们已经使用Ansible脚本简单实现了Tomcat War包分发的过程。</p>

<p>这样的持续交付工作流，虽然可以工作，但依然存在两个问题。</p>

<ol>
<li><p>用户体验问题。-
我们一起回顾下第21篇文章《发布系统一定要注意用户体验》中的相关内容，用户体验对发布系统来说是相当重要的。-
在上面使用Ansible进行部署Java应用的方案中，我们采用的Jenkins Pipeline和Ansible命令行直接集成的方式，就所有的信息都集中到了Jenkins的console log下面，从而缺少了对发布状态、异常日志的直观展示，整个发布体验很糟糕。</p></li>

<li><p>统一管理问题。-
Ansible缺乏集中式管理，需要在每个Jenkins节点上进行Ansible的初始化，增加了管理成本。</p></li>
</ol>

<p>而这两个问题，我们都可以通过Ansible Tower解决。</p>

<p><img src="assets/88fc43236173194403445fe383a42bb5.png" alt="" /></p>

<p>图1 Ansible Dashboard（ 来源 Ansible 官网）</p>

<p>Ansible Tower是Ansible的中心化管理节点，既提供了Web页面以达到可视化能力，也提供了Rest API以达到调用Ansible的PlayBook的目的。</p>

<p>如图1所示为Ansible Tower的Dashboard页面。我们可以看到，这个页面提供了整个Ansible集群发布的趋势图，以及每次发布在每台被部署机器上的详细结果。</p>

<h2 id="灰度发布的处理">灰度发布的处理</h2>

<p>通过上面的内容，我们已经可以通过合理使用Ansible，顺利地部署一个Java应用了，而且还可以通过Ansible Tower监控整个发布过程。而对于灰度发布过程的处理，你只需要在Jenkins Pipeline中编写相应的脚本，控制具体的发布过程就可以了。</p>

<p>比如，通过Inventory定义灰度分批策略，再利用Pipeline驱动PlayBook，就是一个典型的灰度发布的处理过程。其实，这只是将原子化的单机操作批量化了而已。</p>

<p>当然，这个过程中我们还需要考虑其他一些问题。而对于这些问题如何解决，你就可以参考发布及监控系列的六篇文章（即，第19篇至第24篇）了。</p>

<p>至此，标准的Java应用的发布就已经大功告成了。接下来，我再和你说说其他产物（Jar包、Docker镜像）的发布方式。</p>

<h2 id="jar包的发布">Jar包的发布</h2>

<p><strong>Jar包的发布本身就比较简单，执行一条Maven命令（即，mvn deploy）就可以完成。但，Jar包发布的关键在于，如何通过工具提升Jar包发布的质量。</strong></p>

<p>在不引入任何工具和流程辅助时，我们在携程尝试过让开发人员自行通过“mvn deploy”进行发布。但结果可想而知，造成了很多问题。诸如，大量低质量的代码进入仓库；release版本的Jar包被多次覆盖；版本管理混乱，Bug难以排查等等。</p>

<p>后来，我们初次收紧了发布权限，也就是只把“mvn deploy”的权限下放给每个团队的技术经理（tech leader）。这种方案，虽然在一定程度上解决了Jar包质量的问题，但同时也降低了发布效率。这里发布效率的降低，主要体现在两个方面：</p>

<ul>
<li>一方面，每次发布都需要经过技术经理，增加了他的工作负担；</li>
<li>另一方面，“mvn deploy”权限需要由SCM人员手工完成，增加了发布的沟通成本，降低了整个团队的开发效率。</li>
</ul>

<p>再后来，为了解决这些问题，我们在GitLab上进行了二次开发，即：允许开发人员自主选择某个pom module的Jar包进行发布，并记录下每次的Jar包发布的记录。</p>

<p>在Jar包发布的第一步，我们使用Maven Enforcer插件进行构建检测，以保证所有进入仓库的Jar包是合规的。这部分内容，你可以参考第15篇文章《构建检测，无规矩不成方圆》。</p>

<p>如果你不想通过在GitLab上进行二次开发控制Jar包发布的话，简单的做法是，通过Jenkins任务，参数化创建一个Jar包发布的job。让用户在每次发布前填入所需的代码仓库和module名，并在job的逻辑中保证Jar包编译时已经通过了Enforcer检查。</p>

<p>这样，我们就可以顺利解决掉Jar包发布的问题了。</p>

<h2 id="使用spinnaker处理docker">使用Spinnaker处理Docker</h2>

<p>现在，我们再来看一下如何选择开源的Docker交付平台。</p>

<p>在携程，我们第一版的Docker发布流程，是基于自研发布工具Tars和mesos framework集成实现的。这个方案成型于2016年底，那时容器编排平台的局面还是Mesos、Swarm,，以及Kubernetes的三方大战，三方各有优势和支持者。</p>

<p>时至今日，Kubernetes基本已经一统容器编排平台。为了更多地获取开源红利，携程也在向Kubernetes的全面迁移中。</p>

<p>目前，携程对接Kubernetes的方案是，使用StatefulSet管理Pod，并且保持实例的IP不会因为发布而产生变化，而负载均衡器依然使用之前的SLB中间件，并未使用Kubernetes天然支持的Ingress。这和我在第23篇文章《业务及系统机构对发布的影响》中提到的markdown、markup机制有关系，你可以再回顾一下这篇文章的内容。</p>

<p>但，如果今天让我再重新实现一次的话，我更推荐使用Kubernetes原生方案作为Docker编排平台的第一方案，这样更简单有效。如果你还没有在持续交付平台中支持Kubernetes的话，我的建议是：直接考虑搭建持续交付平台Spinnaker。</p>

<p>Spinnaker 是 Netflix 的开源项目，致力于解除持续交付平台和云平台之间的耦合。这个持续交付平台的优点，主要包括：</p>

<ol>
<li><p>发布支持多个云平台，比如AWS EC2、Microsoft Azure、Kubernetes等。如果你未来有在多数据中心使用混合云的打算，Spinnaker可以给你提供很多帮助。</p></li>

<li><p>支持集成多个持续集成平台，包括Jenkins、Travis CI等。</p></li>

<li><p>Netflix 是金丝雀发布的早期实践者，Spinnaker中已经天然集成了蓝绿发布和金丝雀发布这两种发布策略，减少了开发发布系统的工作量。 在此，你可以回顾一下我在第19篇文章《发布是持续交付的最后一公里》中，和你分享的蓝绿发布和金丝雀发布。</p></li>
</ol>

<p><img src="assets/1b7002b9c0c85bb87075a3a7531ea2d4.png" alt="" /></p>

<p>图2 Spinnaker 金丝雀发布配置图（来源 Spinnaker 官网）</p>

<p>虽然，我并未在携程的生产环境中使用过Spinnaker，但由处于持续交付领域领头羊地位的Netflix出品，并且在国内也已经有了小红书的成功案例，Spinnaker还是值得信任的。你可以放心大胆的用到自己的持续交付体系中。</p>

<p>好了，现在我们已经一起完成了发布平台的搭建。至此，整个持续交付体系，从代码管理到集成编译再到程序发布上线的完整过程，就算是顺利完成了。</p>

<h2 id="总结与实践">总结与实践</h2>

<p>在今天这篇文章中，我主要基于Ansible系统的能力，和你分享了搭建一套部署系统的过程。在搭建过程中，你最需要关注的两部分内容是：</p>

<ol>
<li><p>利用Inventory做好部署目标的管理；</p></li>

<li><p>利用PlayBook编写部署过程的具体逻辑。</p></li>
</ol>

<p>同时，我还介绍了Ansible Tower这样一个可视化工具，可以帮助你更好地管理整个部署过程。</p>

<p>另外，对于Jar包的发布，以及Docker的处理，我也结合着携程的经验，和你分享了一些方法和建议，希望可以帮到你。</p>

<p>至此，我们要搭建的整个持续交付系统，也算是顺利完成了。</p>

<p>同样地，最后我还是建议你动手去搭建一套发布系统，看看是否能够顺利地完成这个过程。如果你在这个过程中碰到了任何问题，欢迎你给我留言一起讨论。-</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#177b7b7b2e232626272057707a767e7b3974787a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f14f92e1b284596',t:'MTczNDA4MjYyMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;自动化测试：如何把Bug杀死在摇篮里？>
        <link rel="icon" href="/static/favicon.png">
        <title>29 自动化测试：如何把Bug杀死在摇篮里？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 你为什么应该学好软件工程？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bd%a0%e4%b8%ba%e4%bb%80%e4%b9%88%e5%ba%94%e8%af%a5%e5%ad%a6%e5%a5%bd%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%ef%bc%9f.md">00 开篇词 你为什么应该学好软件工程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 到底应该怎样理解软件工程？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/01%20%e5%88%b0%e5%ba%95%e5%ba%94%e8%af%a5%e6%80%8e%e6%a0%b7%e7%90%86%e8%a7%a3%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%ef%bc%9f.md">01 到底应该怎样理解软件工程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 工程思维：把每件事都当作一个项目来推进.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/02%20%e5%b7%a5%e7%a8%8b%e6%80%9d%e7%bb%b4%ef%bc%9a%e6%8a%8a%e6%af%8f%e4%bb%b6%e4%ba%8b%e9%83%bd%e5%bd%93%e4%bd%9c%e4%b8%80%e4%b8%aa%e9%a1%b9%e7%9b%ae%e6%9d%a5%e6%8e%a8%e8%bf%9b.md">02 工程思维：把每件事都当作一个项目来推进.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 瀑布模型：像工厂流水线一样把软件开发分层化.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/03%20%e7%80%91%e5%b8%83%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%83%8f%e5%b7%a5%e5%8e%82%e6%b5%81%e6%b0%b4%e7%ba%bf%e4%b8%80%e6%a0%b7%e6%8a%8a%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%88%86%e5%b1%82%e5%8c%96.md">03 瀑布模型：像工厂流水线一样把软件开发分层化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 瀑布模型之外，还有哪些开发模型？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/04%20%e7%80%91%e5%b8%83%e6%a8%a1%e5%9e%8b%e4%b9%8b%e5%a4%96%ef%bc%8c%e8%bf%98%e6%9c%89%e5%93%aa%e4%ba%9b%e5%bc%80%e5%8f%91%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">04 瀑布模型之外，还有哪些开发模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 敏捷开发到底是想解决什么问题？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/05%20%e6%95%8f%e6%8d%b7%e5%bc%80%e5%8f%91%e5%88%b0%e5%ba%95%e6%98%af%e6%83%b3%e8%a7%a3%e5%86%b3%e4%bb%80%e4%b9%88%e9%97%ae%e9%a2%98%ef%bc%9f.md">05 敏捷开发到底是想解决什么问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 大厂都在用哪些敏捷方法？（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/06%20%e5%a4%a7%e5%8e%82%e9%83%bd%e5%9c%a8%e7%94%a8%e5%93%aa%e4%ba%9b%e6%95%8f%e6%8d%b7%e6%96%b9%e6%b3%95%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">06 大厂都在用哪些敏捷方法？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 大厂都在用哪些敏捷方法？（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/07%20%e5%a4%a7%e5%8e%82%e9%83%bd%e5%9c%a8%e7%94%a8%e5%93%aa%e4%ba%9b%e6%95%8f%e6%8d%b7%e6%96%b9%e6%b3%95%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">07 大厂都在用哪些敏捷方法？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 怎样平衡软件质量与时间成本范围的关系？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/08%20%e6%80%8e%e6%a0%b7%e5%b9%b3%e8%a1%a1%e8%bd%af%e4%bb%b6%e8%b4%a8%e9%87%8f%e4%b8%8e%e6%97%b6%e9%97%b4%e6%88%90%e6%9c%ac%e8%8c%83%e5%9b%b4%e7%9a%84%e5%85%b3%e7%b3%bb%ef%bc%9f.md">08 怎样平衡软件质量与时间成本范围的关系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 为什么软件工程项目普遍不重视可行性分析？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/09%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e9%a1%b9%e7%9b%ae%e6%99%ae%e9%81%8d%e4%b8%8d%e9%87%8d%e8%a7%86%e5%8f%af%e8%a1%8c%e6%80%a7%e5%88%86%e6%9e%90%ef%bc%9f.md">09 为什么软件工程项目普遍不重视可行性分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 如果你想技术转管理，先来试试管好一个项目.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/10%20%e5%a6%82%e6%9e%9c%e4%bd%a0%e6%83%b3%e6%8a%80%e6%9c%af%e8%bd%ac%e7%ae%a1%e7%90%86%ef%bc%8c%e5%85%88%e6%9d%a5%e8%af%95%e8%af%95%e7%ae%a1%e5%a5%bd%e4%b8%80%e4%b8%aa%e9%a1%b9%e7%9b%ae.md">10 如果你想技术转管理，先来试试管好一个项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 项目计划：代码未动，计划先行.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/11%20%e9%a1%b9%e7%9b%ae%e8%ae%a1%e5%88%92%ef%bc%9a%e4%bb%a3%e7%a0%81%e6%9c%aa%e5%8a%a8%ef%bc%8c%e8%ae%a1%e5%88%92%e5%85%88%e8%a1%8c.md">11 项目计划：代码未动，计划先行.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 流程和规范：红绿灯不是约束，而是用来提高效率.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/12%20%e6%b5%81%e7%a8%8b%e5%92%8c%e8%a7%84%e8%8c%83%ef%bc%9a%e7%ba%a2%e7%bb%bf%e7%81%af%e4%b8%8d%e6%98%af%e7%ba%a6%e6%9d%9f%ef%bc%8c%e8%80%8c%e6%98%af%e7%94%a8%e6%9d%a5%e6%8f%90%e9%ab%98%e6%95%88%e7%8e%87.md">12 流程和规范：红绿灯不是约束，而是用来提高效率.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 白天开会，加班写代码的节奏怎么破？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/13%20%e7%99%bd%e5%a4%a9%e5%bc%80%e4%bc%9a%ef%bc%8c%e5%8a%a0%e7%8f%ad%e5%86%99%e4%bb%a3%e7%a0%81%e7%9a%84%e8%8a%82%e5%a5%8f%e6%80%8e%e4%b9%88%e7%a0%b4%ef%bc%9f.md">13 白天开会，加班写代码的节奏怎么破？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 项目管理工具：一切管理问题，都应思考能否通过工具解决.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/14%20%e9%a1%b9%e7%9b%ae%e7%ae%a1%e7%90%86%e5%b7%a5%e5%85%b7%ef%bc%9a%e4%b8%80%e5%88%87%e7%ae%a1%e7%90%86%e9%97%ae%e9%a2%98%ef%bc%8c%e9%83%bd%e5%ba%94%e6%80%9d%e8%80%83%e8%83%bd%e5%90%a6%e9%80%9a%e8%bf%87%e5%b7%a5%e5%85%b7%e8%a7%a3%e5%86%b3.md">14 项目管理工具：一切管理问题，都应思考能否通过工具解决.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 风险管理：不能盲目乐观，凡事都应该有B计划.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/15%20%e9%a3%8e%e9%99%a9%e7%ae%a1%e7%90%86%ef%bc%9a%e4%b8%8d%e8%83%bd%e7%9b%b2%e7%9b%ae%e4%b9%90%e8%a7%82%ef%bc%8c%e5%87%a1%e4%ba%8b%e9%83%bd%e5%ba%94%e8%af%a5%e6%9c%89B%e8%ae%a1%e5%88%92.md">15 风险管理：不能盲目乐观，凡事都应该有B计划.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 怎样才能写好项目文档？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/16%20%e6%80%8e%e6%a0%b7%e6%89%8d%e8%83%bd%e5%86%99%e5%a5%bd%e9%a1%b9%e7%9b%ae%e6%96%87%e6%a1%a3%ef%bc%9f.md">16 怎样才能写好项目文档？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 需求分析到底要分析什么？怎么分析？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/17%20%e9%9c%80%e6%b1%82%e5%88%86%e6%9e%90%e5%88%b0%e5%ba%95%e8%a6%81%e5%88%86%e6%9e%90%e4%bb%80%e4%b9%88%ef%bc%9f%e6%80%8e%e4%b9%88%e5%88%86%e6%9e%90%ef%bc%9f.md">17 需求分析到底要分析什么？怎么分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 原型设计：如何用最小的代价完成产品特性？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/18%20%e5%8e%9f%e5%9e%8b%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e6%9c%80%e5%b0%8f%e7%9a%84%e4%bb%a3%e4%bb%b7%e5%ae%8c%e6%88%90%e4%ba%a7%e5%93%81%e7%89%b9%e6%80%a7%ef%bc%9f.md">18 原型设计：如何用最小的代价完成产品特性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 作为程序员，你应该有产品意识.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/19%20%e4%bd%9c%e4%b8%ba%e7%a8%8b%e5%ba%8f%e5%91%98%ef%bc%8c%e4%bd%a0%e5%ba%94%e8%af%a5%e6%9c%89%e4%ba%a7%e5%93%81%e6%84%8f%e8%af%86.md">19 作为程序员，你应该有产品意识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 如何应对让人头疼的需求变更问题？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/20%20%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e8%ae%a9%e4%ba%ba%e5%a4%b4%e7%96%bc%e7%9a%84%e9%9c%80%e6%b1%82%e5%8f%98%e6%9b%b4%e9%97%ae%e9%a2%98%ef%bc%9f.md">20 如何应对让人头疼的需求变更问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 架构设计：普通程序员也能实现复杂系统？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/21%20%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%9a%e6%99%ae%e9%80%9a%e7%a8%8b%e5%ba%8f%e5%91%98%e4%b9%9f%e8%83%bd%e5%ae%9e%e7%8e%b0%e5%a4%8d%e6%9d%82%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">21 架构设计：普通程序员也能实现复杂系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 如何为项目做好技术选型？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/22%20%e5%a6%82%e4%bd%95%e4%b8%ba%e9%a1%b9%e7%9b%ae%e5%81%9a%e5%a5%bd%e6%8a%80%e6%9c%af%e9%80%89%e5%9e%8b%ef%bc%9f.md">22 如何为项目做好技术选型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 架构师：不想当架构师的程序员不是好程序员.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/23%20%e6%9e%b6%e6%9e%84%e5%b8%88%ef%bc%9a%e4%b8%8d%e6%83%b3%e5%bd%93%e6%9e%b6%e6%9e%84%e5%b8%88%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%e4%b8%8d%e6%98%af%e5%a5%bd%e7%a8%8b%e5%ba%8f%e5%91%98.md">23 架构师：不想当架构师的程序员不是好程序员.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 技术债务：是继续修修补补凑合着用，还是推翻重来？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/24%20%e6%8a%80%e6%9c%af%e5%80%ba%e5%8a%a1%ef%bc%9a%e6%98%af%e7%bb%a7%e7%bb%ad%e4%bf%ae%e4%bf%ae%e8%a1%a5%e8%a1%a5%e5%87%91%e5%90%88%e7%9d%80%e7%94%a8%ef%bc%8c%e8%bf%98%e6%98%af%e6%8e%a8%e7%bf%bb%e9%87%8d%e6%9d%a5%ef%bc%9f.md">24 技术债务：是继续修修补补凑合着用，还是推翻重来？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 有哪些方法可以提高开发效率？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/25%20%e6%9c%89%e5%93%aa%e4%ba%9b%e6%96%b9%e6%b3%95%e5%8f%af%e4%bb%a5%e6%8f%90%e9%ab%98%e5%bc%80%e5%8f%91%e6%95%88%e7%8e%87%ef%bc%9f.md">25 有哪些方法可以提高开发效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 持续交付：如何做到随时发布新版本到生产环境？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/26%20%e6%8c%81%e7%bb%ad%e4%ba%a4%e4%bb%98%ef%bc%9a%e5%a6%82%e4%bd%95%e5%81%9a%e5%88%b0%e9%9a%8f%e6%97%b6%e5%8f%91%e5%b8%83%e6%96%b0%e7%89%88%e6%9c%ac%e5%88%b0%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%ef%bc%9f.md">26 持续交付：如何做到随时发布新版本到生产环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 软件工程师的核心竞争力是什么？（上）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/27%20%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e5%b8%88%e7%9a%84%e6%a0%b8%e5%bf%83%e7%ab%9e%e4%ba%89%e5%8a%9b%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">27 软件工程师的核心竞争力是什么？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 软件工程师的核心竞争力是什么？（下）.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/28%20%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e5%b8%88%e7%9a%84%e6%a0%b8%e5%bf%83%e7%ab%9e%e4%ba%89%e5%8a%9b%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">28 软件工程师的核心竞争力是什么？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 自动化测试：如何把Bug杀死在摇篮里？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/29%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b5%8b%e8%af%95%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8a%8aBug%e6%9d%80%e6%ad%bb%e5%9c%a8%e6%91%87%e7%af%ae%e9%87%8c%ef%bc%9f.md">29 自动化测试：如何把Bug杀死在摇篮里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 用好源代码管理工具，让你的协作更高效.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/30%20%e7%94%a8%e5%a5%bd%e6%ba%90%e4%bb%a3%e7%a0%81%e7%ae%a1%e7%90%86%e5%b7%a5%e5%85%b7%ef%bc%8c%e8%ae%a9%e4%bd%a0%e7%9a%84%e5%8d%8f%e4%bd%9c%e6%9b%b4%e9%ab%98%e6%95%88.md">30 用好源代码管理工具，让你的协作更高效.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 软件测试要为产品质量负责吗？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/31%20%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%95%e8%a6%81%e4%b8%ba%e4%ba%a7%e5%93%81%e8%b4%a8%e9%87%8f%e8%b4%9f%e8%b4%a3%e5%90%97%ef%bc%9f.md">31 软件测试要为产品质量负责吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 软件测试：什么样的公司需要专职测试？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/32%20%e8%bd%af%e4%bb%b6%e6%b5%8b%e8%af%95%ef%bc%9a%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e5%85%ac%e5%8f%b8%e9%9c%80%e8%a6%81%e4%b8%93%e8%81%8c%e6%b5%8b%e8%af%95%ef%bc%9f.md">32 软件测试：什么样的公司需要专职测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 测试工具：为什么不应该通过QQ微信邮件报Bug？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/33%20%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%8d%e5%ba%94%e8%af%a5%e9%80%9a%e8%bf%87QQ%e5%be%ae%e4%bf%a1%e9%82%ae%e4%bb%b6%e6%8a%a5Bug%ef%bc%9f.md">33 测试工具：为什么不应该通过QQ微信邮件报Bug？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 账号密码泄露成灾，应该怎样预防？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/34%20%e8%b4%a6%e5%8f%b7%e5%af%86%e7%a0%81%e6%b3%84%e9%9c%b2%e6%88%90%e7%81%be%ef%bc%8c%e5%ba%94%e8%af%a5%e6%80%8e%e6%a0%b7%e9%a2%84%e9%98%b2%ef%bc%9f.md">34 账号密码泄露成灾，应该怎样预防？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 版本发布：软件上线只是新的开始.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/35%20%e7%89%88%e6%9c%ac%e5%8f%91%e5%b8%83%ef%bc%9a%e8%bd%af%e4%bb%b6%e4%b8%8a%e7%ba%bf%e5%8f%aa%e6%98%af%e6%96%b0%e7%9a%84%e5%bc%80%e5%a7%8b.md">35 版本发布：软件上线只是新的开始.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 DevOps工程师到底要做什么事情？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/36%20DevOps%e5%b7%a5%e7%a8%8b%e5%b8%88%e5%88%b0%e5%ba%95%e8%a6%81%e5%81%9a%e4%bb%80%e4%b9%88%e4%ba%8b%e6%83%85%ef%bc%9f.md">36 DevOps工程师到底要做什么事情？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 遇到线上故障，你和高手的差距在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/37%20%e9%81%87%e5%88%b0%e7%ba%bf%e4%b8%8a%e6%95%85%e9%9a%9c%ef%bc%8c%e4%bd%a0%e5%92%8c%e9%ab%98%e6%89%8b%e7%9a%84%e5%b7%ae%e8%b7%9d%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">37 遇到线上故障，你和高手的差距在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 日志管理：如何借助工具快速发现和定位产品问题 ？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/38%20%e6%97%a5%e5%bf%97%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%80%9f%e5%8a%a9%e5%b7%a5%e5%85%b7%e5%bf%ab%e9%80%9f%e5%8f%91%e7%8e%b0%e5%92%8c%e5%ae%9a%e4%bd%8d%e4%ba%a7%e5%93%81%e9%97%ae%e9%a2%98%20%ef%bc%9f.md">38 日志管理：如何借助工具快速发现和定位产品问题 ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 项目总结：做好项目复盘，把经验变成能力.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/39%20%e9%a1%b9%e7%9b%ae%e6%80%bb%e7%bb%93%ef%bc%9a%e5%81%9a%e5%a5%bd%e9%a1%b9%e7%9b%ae%e5%a4%8d%e7%9b%98%ef%bc%8c%e6%8a%8a%e7%bb%8f%e9%aa%8c%e5%8f%98%e6%88%90%e8%83%bd%e5%8a%9b.md">39 项目总结：做好项目复盘，把经验变成能力.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 最佳实践：小团队如何应用软件工程？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/40%20%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5%ef%bc%9a%e5%b0%8f%e5%9b%a2%e9%98%9f%e5%a6%82%e4%bd%95%e5%ba%94%e7%94%a8%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%ef%bc%9f.md">40 最佳实践：小团队如何应用软件工程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 为什么程序员的业余项目大多都死了？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/41%20%e4%b8%ba%e4%bb%80%e4%b9%88%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e4%b8%9a%e4%bd%99%e9%a1%b9%e7%9b%ae%e5%a4%a7%e5%a4%9a%e9%83%bd%e6%ad%bb%e4%ba%86%ef%bc%9f.md">41 为什么程序员的业余项目大多都死了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 反面案例：盘点那些失败的软件项目.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/42%20%e5%8f%8d%e9%9d%a2%e6%a1%88%e4%be%8b%ef%bc%9a%e7%9b%98%e7%82%b9%e9%82%a3%e4%ba%9b%e5%a4%b1%e8%b4%a5%e7%9a%84%e8%bd%af%e4%bb%b6%e9%a1%b9%e7%9b%ae.md">42 反面案例：盘点那些失败的软件项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 以VS Code为例，看大型开源项目是如何应用软件工程的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/43%20%e4%bb%a5VS%20Code%e4%b8%ba%e4%be%8b%ef%bc%8c%e7%9c%8b%e5%a4%a7%e5%9e%8b%e5%bc%80%e6%ba%90%e9%a1%b9%e7%9b%ae%e6%98%af%e5%a6%82%e4%bd%95%e5%ba%94%e7%94%a8%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e7%9a%84%ef%bc%9f.md">43 以VS Code为例，看大型开源项目是如何应用软件工程的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 微软、谷歌、阿里巴巴等大厂是怎样应用软件工程的？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/44%20%e5%be%ae%e8%bd%af%e3%80%81%e8%b0%b7%e6%ad%8c%e3%80%81%e9%98%bf%e9%87%8c%e5%b7%b4%e5%b7%b4%e7%ad%89%e5%a4%a7%e5%8e%82%e6%98%af%e6%80%8e%e6%a0%b7%e5%ba%94%e7%94%a8%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e7%9a%84%ef%bc%9f.md">44 微软、谷歌、阿里巴巴等大厂是怎样应用软件工程的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 从软件工程的角度看微服务、云计算、人工智能这些新技术.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/45%20%e4%bb%8e%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e7%9a%84%e8%a7%92%e5%ba%a6%e7%9c%8b%e5%be%ae%e6%9c%8d%e5%8a%a1%e3%80%81%e4%ba%91%e8%ae%a1%e7%ae%97%e3%80%81%e4%ba%ba%e5%b7%a5%e6%99%ba%e8%83%bd%e8%bf%99%e4%ba%9b%e6%96%b0%e6%8a%80%e6%9c%af.md">45 从软件工程的角度看微服务、云计算、人工智能这些新技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="一问一答第1期 30个软件开发常见问题解决策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e4%b8%80%e9%97%ae%e4%b8%80%e7%ad%94%e7%ac%ac1%e6%9c%9f%2030%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e8%a7%a3%e5%86%b3%e7%ad%96%e7%95%a5.md">一问一答第1期 30个软件开发常见问题解决策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="一问一答第2期 30个软件开发常见问题解决策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e4%b8%80%e9%97%ae%e4%b8%80%e7%ad%94%e7%ac%ac2%e6%9c%9f%2030%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e8%a7%a3%e5%86%b3%e7%ad%96%e7%95%a5.md">一问一答第2期 30个软件开发常见问题解决策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="一问一答第3期 18个软件开发常见问题解决策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e4%b8%80%e9%97%ae%e4%b8%80%e7%ad%94%e7%ac%ac3%e6%9c%9f%2018%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e8%a7%a3%e5%86%b3%e7%ad%96%e7%95%a5.md">一问一答第3期 18个软件开发常见问题解决策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="一问一答第4期 14个软件开发常见问题解决策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e4%b8%80%e9%97%ae%e4%b8%80%e7%ad%94%e7%ac%ac4%e6%9c%9f%2014%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e8%a7%a3%e5%86%b3%e7%ad%96%e7%95%a5.md">一问一答第4期 14个软件开发常见问题解决策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="一问一答第5期 22个软件开发常见问题解决策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e4%b8%80%e9%97%ae%e4%b8%80%e7%ad%94%e7%ac%ac5%e6%9c%9f%2022%e4%b8%aa%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98%e8%a7%a3%e5%86%b3%e7%ad%96%e7%95%a5.md">一问一答第5期 22个软件开发常见问题解决策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="学习攻略 怎样学好软件工程？.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e5%ad%a6%e4%b9%a0%e6%94%bb%e7%95%a5%20%e6%80%8e%e6%a0%b7%e5%ad%a6%e5%a5%bd%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%ef%bc%9f.md">学习攻略 怎样学好软件工程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 从软件工程的角度解读任正非的新年公开信.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e4%bb%8e%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e7%9a%84%e8%a7%92%e5%ba%a6%e8%a7%a3%e8%af%bb%e4%bb%bb%e6%ad%a3%e9%9d%9e%e7%9a%84%e6%96%b0%e5%b9%b4%e5%85%ac%e5%bc%80%e4%bf%a1.md">特别放送 从软件工程的角度解读任正非的新年公开信.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 万事皆项目，软件工程无处不在.md" href="/%e4%b8%93%e6%a0%8f/%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e4%b9%8b%e7%be%8e/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e4%b8%87%e4%ba%8b%e7%9a%86%e9%a1%b9%e7%9b%ae%ef%bc%8c%e8%bd%af%e4%bb%b6%e5%b7%a5%e7%a8%8b%e6%97%a0%e5%a4%84%e4%b8%8d%e5%9c%a8.md">结束语 万事皆项目，软件工程无处不在.md</a>
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
                            <h1 id="title" data-id="29 自动化测试：如何把Bug杀死在摇篮里？" class="title">29 自动化测试：如何把Bug杀死在摇篮里？</h1>
                            <div><p>你好，我是宝玉。前不久我所在项目组完成了一个大项目，把一个网站前端的 jQuery 代码全部换成 React 代码，涉及改动的项目源代码文件有一百多个，变动的代码有几千行，最终上线后出乎意料的稳定，只有几个不算太严重的 Bug。</p>

<p>能做到重构还这么稳定，是因为我们技术水平特别高吗？当然不是。还是同样一组人，一年前做一个比这还简单的项目，上线后却跟噩梦一样，频繁出各种问题，导致上线后不停打补丁，一段时间才逐步稳定下来。</p>

<p>这其中的差别，只是因为在那次失败的上线后，我们总结经验，逐步增加了自动化测试代码的覆盖率。等我们再做大的重构时，这些自动化测试代码就能帮助我们发现很多问题。</p>

<p>当我们确保每一个以前写好的测试用例能正常通过后，就相当于把 Bug 杀死在摇篮里，再配合少量的人工手动测试，就可以保证上线后的系统是稳定的。</p>

<p>其实对于自动化测试，我们专栏已经多次提及，它是敏捷开发能快速迭代很重要的质量保障，是持续交付的基础前提。</p>

<p>所以今天我将带你一起了解什么是自动化测试，以及如何在项目中应用自动化测试。</p>

<h2 id="为什么自动化测试能保障质量">为什么自动化测试能保障质量？</h2>

<p>自动化测试并不难理解，你可以想想人是怎么做测试的：首先根据需求写成测试用例，设计好输入值和期望的输出，然后按照测试用例一个个操作，输入一些内容，做一些操作，观察是不是和期望的结果一致，一致就通过，不一致就不通过。</p>

<p>自动化测试，就是把这些操作，用程序脚本来完成的，本质上还是要输入和操作，要检查输出是不是和期望值一致。只要能按照测试用例操作和检查，其实是人来做还是程序来做，结果都是一样的。</p>

<p>不过，自动化测试有一个手工测试没有的优势，那就是可以直接绕过界面，对程序内部的类、函数进行直接测试，如果有一定量的自动化测试代码覆盖，相对来说软件质量是更有保障的。</p>

<p>而且，一旦实现了自动化，每测试一次的成本其实大幅降低了的，几百个测试用例可能几分钟就跑完了。尤其是每次修改完代码，合并到主干之前，把这几百个测试用例跑一遍，可以有效地预防“修复一个 Bug 而产生新 Bug”的情况发生。</p>

<p>但现阶段，自动化测试还是不能完全代替手工测试的，有些测试，自动化测试成本比手工测试成本要高，比如说测试界面布局、颜色等，还是需要一定量的手工测试配合。</p>

<h2 id="有哪些类型的自动化测试">有哪些类型的自动化测试？</h2>

<p>现在说到自动化测试，已经有很多的概念，除了大家熟悉的单元测试，还有像集成测试、UI 测试、端到端测试、契约测试、组件测试等。而很多时候同一个名字还有不同的解读，很容易混淆。</p>

<p>在对自动化测试类型的定义方面，Google 的分类方法我觉得比较科学：根据数据做出决策，而不仅仅是依靠直觉或无法衡量和评估的内容。Google 将自动化测试分成了三大类：小型测试、中型测试和大型测试。</p>

<p>假设我们有一个网站，是基于三层架构（如下图所示），业务逻辑层的类叫 UserService 类，数据访问层的类叫 UserDA，我们将以用户注册的功能来说明几种测试的区别 。</p>

<p><img src="assets/9b9fbf93cf03fa33b381ee144a26a92b.png" alt="img" /></p>

<h4 id="小型测试">小型测试</h4>

<p>小型测试是为了验证一个代码单元的功能，例如针对一个函数或者一个类的测试。我们平时说的单元测试就是一个典型的小型测试。</p>

<p>比如说 UserService 这个类，有一个注册用户的函数，现在要对它写一个单元测试代码，那么看起来就像下面这样：</p>

<p><img src="assets/02aa850792c8fbb3c6bc626b9c944161.png" alt="img" /></p>

<p>通过这样的测试代码，就可以清楚的知道 UserService 类的 create 这个函数是不是能正常工作。</p>

<p>小型测试的运行，不需要依赖外部。如果有外部服务（比如文件操作、网络服务、数据库等），必须使用一个模拟的外部服务。比如上面例子中我们就使用了 FakeUserDA 这个模拟的数据库访问类，实际上它不会访问真实的数据库。这样可以保证小型测试在很短时间内就可以完成。</p>

<p><img src="assets/43ce39715dddae2e51728d13714c31ee.png" alt="img" /></p>

<p>小型测试，图片来源：《Google软件测试之道》</p>

<h4 id="中型测试">中型测试</h4>

<p>中型测试是验证两个或多个模块应用之间的交互，通常也叫集成测试。</p>

<p>如果说要对用户注册的功能写集成测试，那么就会同时测试业务逻辑层的 UserService 类和数据访问层的 UserDA 类。如下所示：</p>

<p><img src="assets/7dcb05ac985cdc8b554c9ba3b5691ee7.png" alt="img" /></p>

<p>对于中型测试，可以使用外部服务（比如文件操作、网络服务、数据库等），可以模拟也可以使用真实的服务。比如上面这个例子，就是真实的数据库访问类，但是用的内存数据库，这样可以提高性能，也可以减少依赖。</p>

<p>至于中型测试要不要使用模拟的服务，有个简单的标准，就是看能不能在单机情况下完成集成测试，如果可以就不需要模拟，如果不能，则要模拟避免外部依赖。</p>

<p><img src="assets/cfb41a9f0a490a3e1aa54555e4d35eab.png" alt="img" /></p>

<p>中型测试，图片来源：《Google软件测试之道》</p>

<h4 id="大型测试">大型测试</h4>

<p>大型测试则是从较高的层次运行，把系统作为一个整体验证。会验证系统的一个或者所有子系统，从前端一直到后端数据存储。大型测试也叫系统测试或者端对端测试。</p>

<p>如果说要对用户注册写一个端对端测试的例子，那么看起来会像这样：</p>

<p><img src="assets/736f1fc4609ba5d5c408b243b30834b0.png" alt="img" /></p>

<p>对于大型测试，通常会直接使用外部服务（比如文件操作、网络服务、数据库等），而不会去模拟。比如上面这个例子，就是直接访问测试环境的地址，通过测试库提供的 API 操作浏览器界面，输入测试的用户名密码，点击注册按钮，最后检查输出的结果是不是符合预期。</p>

<p><img src="assets/30ab7a154bc2324f1d4b858e36ad03af.png" alt="img" /></p>

<p>大型测试，图片来源：《Google软件测试之道》</p>

<h4 id="区分测试类型的依据是什么">区分测试类型的依据是什么？</h4>

<p>以上就是主要的自动测试类型了。捎带着补充一个测试类型，那就是契约测试，这个测试最近出现的频率比较高，主要是针对微服务的。其实就是让微服务在测试时，不需要依赖于引用的外部的微服务，在本地就可以模拟运行，同时又可以保证外部微服务的接口更新时，本地模拟的接口（契约）也能同步更新。对契约服务更多的说明可以参考这篇文章：<a href="http://insights.thoughtworks.cn/about-contract-test/" target="_blank">聊一聊契约测试</a></p>

<p>那么契约测试，属于大型测试还是中型测试呢？</p>

<p>Google 针对这几种测试类型列了一张表，根据数据给出了明确区分：</p>

<p><img src="assets/a72fcd3b3f358e4512fa5694ad526dbd.png" alt="img" /></p>

<p>图片来源：Google Testing Blog</p>

<p>结合上面的表格其实就很好区分了：</p>

<ul>
<li><p>小型测试，没有外部服务的依赖，都是要模拟的；</p></li>

<li><p>中型测试，所有的测试几乎都不需要依赖其他服务器的资源，如果有涉及其他机器的服务，则本地模拟，这样本机就可以完成测试；</p></li>

<li><p>大型测试，几乎不模拟，直接访问相关的外部服务。</p></li>
</ul>

<p>所以现在你应该就知道契约测试，也是中型测试的一种了，因为它不需要依赖外部服务，本机就可以完成测试。</p>

<p>为什么中型测试这么看重“能单机运行”这一点呢？因为这样才方便在持续集成上跑中型测试，不用担心外部服务不稳定而导致测试失败的问题。</p>

<p>上面的表中还反映出一个事实：<strong>越是小型测试，执行速度越快，越是大型测试，执行速度越慢。</strong>通常一个项目的小型测试，不超过一分钟就能全部跑完，一个中型测试，包括一些环境准备的时间，可能要几分钟甚至更久，而大型测试就更久了。</p>

<p><strong>另外越是大型测试，写起来的成本也相应的会更高，所以一般项目中，小型测试最多，中型测试次之，大型测试最少。</strong>就像下面这张金字塔图一样。所以我们也常用测试金字塔来区分不同类型的测试粒度。</p>

<p><img src="assets/616bb4cdb13884dde562b10568ba77cf.png" alt="img" /></p>

<p>测试金字塔，图片来源： TestPyramid</p>

<p>如果你对测试类型很感兴趣，可以参考<a href="https://insights.thoughtworks.cn/practical-test-pyramid/" target="_blank">测试金字塔实战</a>这篇文章作为补充。</p>

<h4 id="怎么写好自动化测试代码">怎么写好自动化测试代码？</h4>

<p>很多人认为写自动化测试很复杂，其实测试代码其实写起来不难，包含四部分内容即可，也就是：准备、执行、断言和清理，我再把第一段代码示例贴一下：</p>

<p><img src="assets/02aa850792c8fbb3c6bc626b9c944161.png" alt="img" /></p>

<p>第一步就是准备，例如创建实例，创建模拟对象；第二步就是执行要测试的方法，传入要测试的参数；第三步断言就是检查结果对不对，如果不对测试会失败；第四步还要对数据进行清理，这样不影响下一次测试。</p>

<p>上面还有几个测试代码示例，都是这样的四部分内容。</p>

<p>这是针对写一个自动化测试的代码结构。对于同一个功能，通常需要写几个自动化测试才完整。</p>

<p>一个完整的自动化测试要包括三个部分的测试：</p>

<ul>
<li><p><strong>验证功能是不是正确</strong>：例如说输入正确的用户名和密码，要能正常注册账号；</p></li>

<li><p><strong>覆盖边界条件</strong>： 比如说如果用户名或密码为空，应该不允许注册成功；</p></li>

<li><p><strong>异常和错误处理</strong>：比如说使用一个已经用过的用户名，应该提示用户名被使用。</p></li>
</ul>

<p><img src="assets/055004435b9bcd81cfa13050a8f42aa8.png" alt="img" /></p>

<p>所以你看，写一个测试代码并没有你想的那么复杂，那还有什么理由不去写测试呢？</p>

<h2 id="如何为你的项目实施自动化测试">如何为你的项目实施自动化测试？</h2>

<p>现在你了解了有哪些类型的测试，如何写自动化测试代码，也许迫不及待想在项目中实施自动化测试。</p>

<h4 id="选择好自动化测试框架">选择好自动化测试框架</h4>

<p>要写好自动化测试代码，首先要找对自动测试化框架。不同的语言，不同的平台，测试的框架都不一样。好在现在搜索引擎很方便，根据“你的语言 + 自动测试框架”的关键字，就能找到很多的结果。这里我也帮你找了一些，供参考。</p>

<ul>
<li>Web 前端</li>
</ul>

<p>Jest： Facebook 的前端测试框架；</p>

<p>Mocha：历史悠久的一个 JS 测试框架；</p>

<p>Nighwatch: 一个 API 很简单，但是功能很强大，可以直接操作浏览器的自动测试框架。</p>

<ul>
<li>iOS 开发</li>
</ul>

<p>可以参考这篇文章《iOS 自动化测试框架对比》。</p>

<ul>
<li>安卓开发</li>
</ul>

<p>可以参考这篇文章《Android 谈谈自动化测试》。</p>

<h4 id="在持续集成环境上跑你的自动化测试">在持续集成环境上跑你的自动化测试</h4>

<p>选好自动化测试框架后，你的自动化测试代码，其中的小型测试和中型测试，最好要能在持续集成环境上运行起来。</p>

<p><strong>让自动化测试在持续集成上运行非常重要，只有这样才能最大化地发挥自动化测试的作用。</strong></p>

<p>因为持续集成，会强制测试通过才能合并代码，在合并代码之前就能知道测试是不是都通过了，可以帮助程序员获得最直观的反馈，知道哪里可能存在问题，这样才能真正做到防患于未然，把 Bug 杀死在摇篮里。</p>

<p>下图描述的就是自动测试配合持续集成的一个标准流程：</p>

<ul>
<li><p>在提交代码前，先本地跑一遍单元测试，这个过程很快的，失败了需要继续修改；</p></li>

<li><p>单元测试成功后就可以提交到源代码管理中心，提交后持续集成服务会自动运行完整的自动化测试，不仅包括小型测试，还有中型测试；</p></li>

<li><p>通过所有的测试后，就可以合并到主分支，如果失败，需要本地修改后再次提交，直到通过所有的测试为止。</p></li>
</ul>

<p><img src="assets/7bbc58d82864974ff2f9ec31347fa538.png" alt="img" /></p>

<p>图片来源：Microservice Testing: Unit Tests</p>

<h4 id="新项目和老项目的不同策略">新项目和老项目的不同策略</h4>

<p>如果是新项目，那么可以在一开始就保持一定的自动化测试代码的覆盖率，你甚至还可以试试测试驱动（TDD）的开发模式，也就是先写测试代码，再写实现代码，保证测试通过，最后对代码进行重构。</p>

<p><img src="assets/67f2886f7dee6f24e5a833e6b4c94f66.png" alt="img" /></p>

<p>图片来源：郑晔 《10x程序员工作法》专栏</p>

<p>如果是老项目，短期内要让自动化测试代码有覆盖是有难度的，可以先把主要的功能场景的中型测试写起来，这样可以保证这些主要功能不会轻易出问题。</p>

<p>后面在维护的过程中：</p>

<ul>
<li><p>增加新功能的时候，同步对新功能增加自动化测试代码；</p></li>

<li><p>修复 Bug 的时候，针对 Bug 补写自动化测试代码。</p></li>
</ul>

<p>这样一点一点，把自动化测试代码覆盖率加上去。</p>

<h4 id="如果时间紧任务重-来不及写自动化测试怎么办">如果时间紧任务重，来不及写自动化测试怎么办？</h4>

<p>确实遇到时间紧的情况，我建议你要优先保证中型测试代码的覆盖，因为这样至少可以保证主要的用户使用场景是正常的。然后把来不及完成的部分，创建一个 Ticket，放到任务跟踪系统里面，后面补上。</p>

<h2 id="总结">总结</h2>

<p>今天我带你一起学习了关于自动化测试有关的知识。自动化测试，分为三类：</p>

<ul>
<li><p>小型测试，主要针对函数或者类进行验证，不调用外部服务，执行速度快；</p></li>

<li><p>中型测试，主要验证两个或多个模块应用之间的交互，可能会调用外部服务，尽可能让所有测试能在本机即可完成，执行速度比较快；</p></li>

<li><p>大型测试，对服务整体进行验证，执行速度慢。</p></li>
</ul>

<p>写好单元测试代码，基本结构就是：准备、执行、断言和清理；基本原则就是：</p>

<ul>
<li><p>要验证正确性；</p></li>

<li><p>覆盖边界条件；</p></li>

<li><p>验证是否有异常和错误的处理。</p></li>
</ul>

<p>自动化测试，一定要配合好持续集成，才能最大化发挥其效用。</p>

<p>对于自动化测试的实施，开头是最难的，因为需要花时间选择自动化测试框架，需要针对自动化测试框架搭建环境，甚至要去搭建持续集成环境。但搭建持续集成和搭建自动化测试环境，并且保证持续更新维护自动测试代码，这个技术投资，一定是你在项目中最有价值的投资之一。</p>

<p>搭建持续集成环境和集成自动化测试框架的事情，要作为一个正式的项目任务去做，当作一个很重要的任务去推进。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#5b373737626f6a6a6b6c1b3c363a323775383436" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f19aba3cb99cdba',t:'MTczNDEzMTg3Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
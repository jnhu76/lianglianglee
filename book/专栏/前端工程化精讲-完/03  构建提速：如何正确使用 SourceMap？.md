<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=03&#32;&#32;构建提速：如何正确使用&#32;SourceMap？>
        <link rel="icon" href="/static/favicon.png">
        <title>03  构建提速：如何正确使用 SourceMap？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  建立上帝视角，全面系统掌握前端效率工程化.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e5%bb%ba%e7%ab%8b%e4%b8%8a%e5%b8%9d%e8%a7%86%e8%a7%92%ef%bc%8c%e5%85%a8%e9%9d%a2%e7%b3%bb%e7%bb%9f%e6%8e%8c%e6%8f%a1%e5%89%8d%e7%ab%af%e6%95%88%e7%8e%87%e5%b7%a5%e7%a8%8b%e5%8c%96.md">00 开篇词  建立上帝视角，全面系统掌握前端效率工程化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  项目基石：前端脚手架工具探秘.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/01%20%20%e9%a1%b9%e7%9b%ae%e5%9f%ba%e7%9f%b3%ef%bc%9a%e5%89%8d%e7%ab%af%e8%84%9a%e6%89%8b%e6%9e%b6%e5%b7%a5%e5%85%b7%e6%8e%a2%e7%a7%98.md">01  项目基石：前端脚手架工具探秘.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  界面调试：热更新技术如何开着飞机修引擎？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/02%20%20%e7%95%8c%e9%9d%a2%e8%b0%83%e8%af%95%ef%bc%9a%e7%83%ad%e6%9b%b4%e6%96%b0%e6%8a%80%e6%9c%af%e5%a6%82%e4%bd%95%e5%bc%80%e7%9d%80%e9%a3%9e%e6%9c%ba%e4%bf%ae%e5%bc%95%e6%93%8e%ef%bc%9f.md">02  界面调试：热更新技术如何开着飞机修引擎？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  构建提速：如何正确使用 SourceMap？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/03%20%20%e6%9e%84%e5%bb%ba%e6%8f%90%e9%80%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e4%bd%bf%e7%94%a8%20SourceMap%ef%bc%9f.md">03  构建提速：如何正确使用 SourceMap？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  接口调试：Mock 工具如何快速进行接口调试？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/04%20%20%e6%8e%a5%e5%8f%a3%e8%b0%83%e8%af%95%ef%bc%9aMock%20%e5%b7%a5%e5%85%b7%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e8%bf%9b%e8%a1%8c%e6%8e%a5%e5%8f%a3%e8%b0%83%e8%af%95%ef%bc%9f.md">04  接口调试：Mock 工具如何快速进行接口调试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  编码效率：如何提高编写代码的效率？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/05%20%20%e7%bc%96%e7%a0%81%e6%95%88%e7%8e%87%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e7%bc%96%e5%86%99%e4%bb%a3%e7%a0%81%e7%9a%84%e6%95%88%e7%8e%87%ef%bc%9f.md">05  编码效率：如何提高编写代码的效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  团队工具：如何利用云开发提升团队开发效率？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/06%20%20%e5%9b%a2%e9%98%9f%e5%b7%a5%e5%85%b7%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e4%ba%91%e5%bc%80%e5%8f%91%e6%8f%90%e5%8d%87%e5%9b%a2%e9%98%9f%e5%bc%80%e5%8f%91%e6%95%88%e7%8e%87%ef%bc%9f.md">06  团队工具：如何利用云开发提升团队开发效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  低代码工具：如何用更少的代码实现更灵活的需求.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/07%20%20%e4%bd%8e%e4%bb%a3%e7%a0%81%e5%b7%a5%e5%85%b7%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a8%e6%9b%b4%e5%b0%91%e7%9a%84%e4%bb%a3%e7%a0%81%e5%ae%9e%e7%8e%b0%e6%9b%b4%e7%81%b5%e6%b4%bb%e7%9a%84%e9%9c%80%e6%b1%82.md">07  低代码工具：如何用更少的代码实现更灵活的需求.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  无代码工具：如何做到不写代码就能高效交付？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/08%20%20%e6%97%a0%e4%bb%a3%e7%a0%81%e5%b7%a5%e5%85%b7%ef%bc%9a%e5%a6%82%e4%bd%95%e5%81%9a%e5%88%b0%e4%b8%8d%e5%86%99%e4%bb%a3%e7%a0%81%e5%b0%b1%e8%83%bd%e9%ab%98%e6%95%88%e4%ba%a4%e4%bb%98%ef%bc%9f.md">08  无代码工具：如何做到不写代码就能高效交付？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  构建总览：前端构建工具的演进.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/09%20%20%e6%9e%84%e5%bb%ba%e6%80%bb%e8%a7%88%ef%bc%9a%e5%89%8d%e7%ab%af%e6%9e%84%e5%bb%ba%e5%b7%a5%e5%85%b7%e7%9a%84%e6%bc%94%e8%bf%9b.md">09  构建总览：前端构建工具的演进.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  流程分解：Webpack 的完整构建流程.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/10%20%20%e6%b5%81%e7%a8%8b%e5%88%86%e8%a7%a3%ef%bc%9aWebpack%20%e7%9a%84%e5%ae%8c%e6%95%b4%e6%9e%84%e5%bb%ba%e6%b5%81%e7%a8%8b.md">10  流程分解：Webpack 的完整构建流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  编译提效：如何为 Webpack 编译阶段提速？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/11%20%20%e7%bc%96%e8%af%91%e6%8f%90%e6%95%88%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%ba%20Webpack%20%e7%bc%96%e8%af%91%e9%98%b6%e6%ae%b5%e6%8f%90%e9%80%9f%ef%bc%9f.md">11  编译提效：如何为 Webpack 编译阶段提速？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  打包提效：如何为 Webpack 打包阶段提速？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/12%20%20%e6%89%93%e5%8c%85%e6%8f%90%e6%95%88%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%ba%20Webpack%20%e6%89%93%e5%8c%85%e9%98%b6%e6%ae%b5%e6%8f%90%e9%80%9f%ef%bc%9f.md">12  打包提效：如何为 Webpack 打包阶段提速？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  缓存优化：那些基于缓存的优化方案.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/13%20%20%e7%bc%93%e5%ad%98%e4%bc%98%e5%8c%96%ef%bc%9a%e9%82%a3%e4%ba%9b%e5%9f%ba%e4%ba%8e%e7%bc%93%e5%ad%98%e7%9a%84%e4%bc%98%e5%8c%96%e6%96%b9%e6%a1%88.md">13  缓存优化：那些基于缓存的优化方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  增量构建：Webpack 中的增量构建.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/14%20%20%e5%a2%9e%e9%87%8f%e6%9e%84%e5%bb%ba%ef%bc%9aWebpack%20%e4%b8%ad%e7%9a%84%e5%a2%9e%e9%87%8f%e6%9e%84%e5%bb%ba.md">14  增量构建：Webpack 中的增量构建.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  版本特性：Webpack 5 中的优化细节.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/15%20%20%e7%89%88%e6%9c%ac%e7%89%b9%e6%80%a7%ef%bc%9aWebpack%205%20%e4%b8%ad%e7%9a%84%e4%bc%98%e5%8c%96%e7%bb%86%e8%8a%82.md">15  版本特性：Webpack 5 中的优化细节.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  无包构建：盘点那些 No-bundle 的构建方案.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/16%20%20%e6%97%a0%e5%8c%85%e6%9e%84%e5%bb%ba%ef%bc%9a%e7%9b%98%e7%82%b9%e9%82%a3%e4%ba%9b%20No-bundle%20%e7%9a%84%e6%9e%84%e5%bb%ba%e6%96%b9%e6%a1%88.md">16  无包构建：盘点那些 No-bundle 的构建方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  部署初探：为什么一般不在开发环境下部署代码？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/17%20%20%e9%83%a8%e7%bd%b2%e5%88%9d%e6%8e%a2%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e4%b8%80%e8%88%ac%e4%b8%8d%e5%9c%a8%e5%bc%80%e5%8f%91%e7%8e%af%e5%a2%83%e4%b8%8b%e9%83%a8%e7%bd%b2%e4%bb%a3%e7%a0%81%ef%bc%9f.md">17  部署初探：为什么一般不在开发环境下部署代码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  工具盘点：掌握那些流行的代码部署工具.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/18%20%20%e5%b7%a5%e5%85%b7%e7%9b%98%e7%82%b9%ef%bc%9a%e6%8e%8c%e6%8f%a1%e9%82%a3%e4%ba%9b%e6%b5%81%e8%a1%8c%e7%9a%84%e4%bb%a3%e7%a0%81%e9%83%a8%e7%bd%b2%e5%b7%a5%e5%85%b7.md">18  工具盘点：掌握那些流行的代码部署工具.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  安装提效：部署流程中的依赖安装效率优化.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/19%20%20%e5%ae%89%e8%a3%85%e6%8f%90%e6%95%88%ef%bc%9a%e9%83%a8%e7%bd%b2%e6%b5%81%e7%a8%8b%e4%b8%ad%e7%9a%84%e4%be%9d%e8%b5%96%e5%ae%89%e8%a3%85%e6%95%88%e7%8e%87%e4%bc%98%e5%8c%96.md">19  安装提效：部署流程中的依赖安装效率优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  流程优化：部署流程中的构建流程策略优化.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/20%20%20%e6%b5%81%e7%a8%8b%e4%bc%98%e5%8c%96%ef%bc%9a%e9%83%a8%e7%bd%b2%e6%b5%81%e7%a8%8b%e4%b8%ad%e7%9a%84%e6%9e%84%e5%bb%ba%e6%b5%81%e7%a8%8b%e7%ad%96%e7%95%a5%e4%bc%98%e5%8c%96.md">20  流程优化：部署流程中的构建流程策略优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  容器方案：从构建到部署，容器化方案的优势有哪些？.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/21%20%20%e5%ae%b9%e5%99%a8%e6%96%b9%e6%a1%88%ef%bc%9a%e4%bb%8e%e6%9e%84%e5%bb%ba%e5%88%b0%e9%83%a8%e7%bd%b2%ef%bc%8c%e5%ae%b9%e5%99%a8%e5%8c%96%e6%96%b9%e6%a1%88%e7%9a%84%e4%bc%98%e5%8a%bf%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">21  容器方案：从构建到部署，容器化方案的优势有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  案例分析：搭建基本的前端高效部署系统.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/22%20%20%e6%a1%88%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%9a%e6%90%ad%e5%bb%ba%e5%9f%ba%e6%9c%ac%e7%9a%84%e5%89%8d%e7%ab%af%e9%ab%98%e6%95%88%e9%83%a8%e7%bd%b2%e7%b3%bb%e7%bb%9f.md">22  案例分析：搭建基本的前端高效部署系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结束语  前端效率工程化的未来展望.md" href="/%e4%b8%93%e6%a0%8f/%e5%89%8d%e7%ab%af%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%b2%be%e8%ae%b2-%e5%ae%8c/23%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e5%89%8d%e7%ab%af%e6%95%88%e7%8e%87%e5%b7%a5%e7%a8%8b%e5%8c%96%e7%9a%84%e6%9c%aa%e6%9d%a5%e5%b1%95%e6%9c%9b.md">23 结束语  前端效率工程化的未来展望.md</a>
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
                            <h1 id="title" data-id="03  构建提速：如何正确使用 SourceMap？" class="title">03  构建提速：如何正确使用 SourceMap？</h1>
                            <div><p>在上一课时中我们聊了开发时的热更新机制和其中的技术细节，热更新能帮助我们在开发时快速预览代码效果，免去了手动执行编译和刷新浏览器的操作。课后留的思考题是找一个实现了 HMR 的 Loader，查看其中用到哪些热替换的 API，希望通过这个题目能让你加深对相关知识点的印象。</p>

<p>那么除了热更新以外，项目的开发环境还有哪些在影响着我们的开发效率呢？在过去的工作中，公司同事就曾问过我一个问题：为什么我的项目在开发环境下每次构建还是很卡？每次保存完代码都要过 1~2 秒才能看到效果，这是怎么回事呢？其实这里面的原因主要是这位同事在开发时选择的 Source Map 设定不对。今天我们就来具体讨论下这个问题。首先，什么是 <strong>Source Map</strong> 呢？</p>

<h3 id="什么是-source-map">什么是 Source Map</h3>

<p>在前端开发过程中，通常我们编写的源代码会经过多重处理（编译、封装、压缩等），最后形成产物代码。于是在浏览器中调试产物代码时，我们往往会发现代码变得面目全非，例如：</p>

<p><img src="assets/Ciqc1F85_FmAA4UeAABWNiHqsWQ595.png" alt="Drawing 0.png" /></p>

<p>因此，我们需要一种在调试时将产物代码显示回源代码的功能，<strong>source map</strong> 就是实现这一目标的工具。</p>

<p>source-map 的基本原理是，在编译处理的过程中，在生成产物代码的同时生成产物代码中被转换的部分与源代码中相应部分的映射关系表。有了这样一张完整的映射表，我们就可以通过 Chrome 控制台中的&rdquo;Enable Javascript source map&rdquo;来实现调试时的显示与定位源代码功能。</p>

<p>对于同一个源文件，根据不同的目标，可以生成不同效果的 source map。它们在<strong>构建速度</strong>、<strong>质量</strong>（反解代码与源代码的接近程度以及调试时行号列号等辅助信息的对应情况）、<strong>访问方式</strong>（在产物文件中或是单独生成 source map 文件）和<strong>文件大小</strong>等方面各不相同。在开发环境和生产环境下，我们对于 source map 功能的期望也有所不同：</p>

<ul>
<li><strong>在开发环境中</strong>，通常我们关注的是构建速度快，质量高，以便于提升开发效率，而不关注生成文件的大小和访问方式。</li>
<li><strong>在生产环境中</strong>，通常我们更关注是否需要提供线上 source map , 生成的文件大小和访问方式是否会对页面性能造成影响等，其次才是质量和构建速度。</li>
</ul>

<h3 id="webpack-中的-source-map-预设">Webpack 中的 source map 预设</h3>

<p>在 Webpack 中，通过设置 devtool 来选择 source map 的预设类型，文档中共有 <a href="https://webpack.js.org/configuration/devtool/#devtool" target="_blank">20 余种</a> source map 的预设（注意：其中部分预设实际效果与其他预设相同，即页面表格中空白行条目）可供选择，这些预设通常包含了 &ldquo;eval&rdquo; &ldquo;cheap&rdquo; &ldquo;module&rdquo; &ldquo;inline&rdquo; &ldquo;hidden&rdquo; &ldquo;nosource&rdquo; &ldquo;source-map&rdquo; 等关键字的组合，这些关键字的具体逻辑如下：</p>

<pre><code>webpack/lib/WebpackOptionsApply.js:232 

if (options.devtool.includes(&quot;source-map&quot;)) { 

  const hidden = options.devtool.includes(&quot;hidden&quot;); 

  const inline = options.devtool.includes(&quot;inline&quot;); 

  const evalWrapped = options.devtool.includes(&quot;eval&quot;); 

  const cheap = options.devtool.includes(&quot;cheap&quot;); 

  const moduleMaps = options.devtool.includes(&quot;module&quot;); 

  const noSources = options.devtool.includes(&quot;nosources&quot;); 

  const Plugin = evalWrapped 

    ? require(&quot;./EvalSourceMapDevToolPlugin&quot;) 

    : require(&quot;./SourceMapDevToolPlugin&quot;); 

  new Plugin({ 

    filename: inline ? null : options.output.sourceMapFilename, 

    moduleFilenameTemplate: options.output.devtoolModuleFilenameTemplate, 

    fallbackModuleFilenameTemplate: 

      options.output.devtoolFallbackModuleFilenameTemplate, 

    append: hidden ? false : undefined, 

    module: moduleMaps ? true : cheap ? false : true, 

    columns: cheap ? false : true, 

    noSources: noSources, 

    namespace: options.output.devtoolNamespace 

  }).apply(compiler); 

} else if (options.devtool.includes(&quot;eval&quot;)) { 

  const EvalDevToolModulePlugin = require(&quot;./EvalDevToolModulePlugin&quot;); 

  new EvalDevToolModulePlugin({ 

    moduleFilenameTemplate: options.output.devtoolModuleFilenameTemplate, 

    namespace: options.output.devtoolNamespace 

  }).apply(compiler); 

}
</code></pre>

<p>如上面的代码所示， devtool 的值匹配并非精确匹配，某个关键字只要包含在赋值中即可获得匹配，例如：&rsquo;foo-eval-bar&rsquo; 等同于 &lsquo;eval&rsquo;，&rsquo;cheapfoo-source-map&rsquo; 等同于 &lsquo;cheap-source-map&rsquo;。</p>

<h4 id="source-map-名称关键字">Source Map 名称关键字</h4>

<p>各字段的作用各不相同，为了便于记忆，我们在这里简单整理下这些关键字的作用：</p>

<ul>
<li><strong>false</strong>：即不开启 source map 功能，其他不符合上述规则的赋值也等价于 false。</li>
<li><strong>eval</strong>：是指在编译器中使用 EvalDevToolModulePlugin 作为 source map 的处理插件。</li>
<li><strong>[xxx-&hellip;]source-map</strong>：根据 devtool 对应值中是否有 <strong>eval</strong> 关键字来决定使用 EvalSourceMapDevToolPlugin 或 SourceMapDevToolPlugin 作为 source map 的处理插件，其余关键字则决定传入到插件的相关字段赋值。</li>
<li><strong>inline</strong>：决定是否传入插件的 filename 参数，作用是决定单独生成 source map 文件还是在行内显示，<strong>该参数在 eval- 参数存在时无效</strong>。</li>
<li><strong>hidden</strong>：决定传入插件 append 的赋值，作用是判断是否添加 SourceMappingURL 的注释，<strong>该参数在 eval- 参数存在时无效</strong>。</li>
<li><strong>module</strong>：为 true 时传入插件的 module 为 true ，作用是为加载器（Loaders）生成 source map。</li>
<li><strong>cheap</strong>：这个关键字有两处作用。首先，当 module 为 false 时，它决定插件 module 参数的最终取值，最终取值与 cheap 相反。其次，它决定插件 columns 参数的取值，作用是决定生成的 source map 中是否包含列信息，在不包含列信息的情况下，调试时只能定位到指定代码所在的行而定位不到所在的列。</li>
<li><strong>nosource</strong>：nosource 决定了插件中 noSource 变量的取值，作用是决定生成的 source map 中是否包含源代码信息，不包含源码情况下只能显示调用堆栈信息。</li>
</ul>

<h4 id="source-map-处理插件">Source Map 处理插件</h4>

<p>从上面的规则中我们还可以看到，根据不同规则，实际上 Webpack 是从三种插件中选择其一作为 source map 的处理插件。</p>

<ul>
<li><a href="https://github.com/webpack/webpack/blob/master/lib/EvalDevToolModulePlugin.js" target="_blank">EvalDevToolModulePlugin</a>：模块代码后添加 sourceURL=webpack:///+ 模块引用路径，不生成 source map 内容，模块产物代码通过 eval() 封装。</li>
<li><a href="https://github.com/webpack/webpack/blob/master/lib/EvalSourceMapDevToolPlugin.js" target="_blank">EvalSourceMapDevToolPlugin</a>：生成 base64 格式的 source map 并附加在模块代码之后， source map 后添加 sourceURL=webpack:///+ 模块引用路径，不单独生成文件，模块产物代码通过 eval() 封装。</li>
<li><a href="https://github.com/webpack/webpack/blob/master/lib/SourceMapDevToolPlugin.js" target="_blank">SourceMapDevToolPlugin</a>：生成单独的 .map 文件，模块产物代码不通过 eval 封装。</li>
</ul>

<p>通过上面的代码分析，我们了解了不同参数在 Webpack 运行时起到的作用。那么这些不同参数组合下的各种预设对我们的 source map 生成又各自会产生什么样的效果呢？下面我们通过示例来看一下。</p>

<h3 id="不同预设的示例结果对比">不同预设的示例结果对比</h3>

<p>下面，以课程示例代码 <a href="https://github.com/fe-efficiency/lessons_fe_efficiency/tree/master/03_develop_environment" target="_blank">03_develop_environment</a> 为例，我们来对比下几种常用预设的差异（为了使时间差异更明显，示例中引入了几个大的类库文件）：</p>

<p><img src="assets/Ciqc1F87kyiAZvHdAAIGvohk2F4144.png" alt="12.png" /></p>

<blockquote>
<p>*注1：“/”前后分别表示产物 js 大小和对应 .map 大小。
*注2：“/”前后分别表示初次构建时间和开启 watch 模式下 rebuild 时间。对应统计的都是 development 模式下的笔者机器环境下几次构建时间的平均值，只作为相对快慢与量级的比较。</p>
</blockquote>

<h4 id="不同预设的效果总结">不同预设的效果总结</h4>

<p>从上图的数据中我们不难发现，选择不同的 devtool 类型在以下几个方面会产生不同的效果。</p>

<ul>
<li>质量：生成的 source map 的质量分为 5 个级别，对应的调试便捷性依次降低：源代码 &gt; 缺少列信息的源代码 &gt; loader 转换后的代码 &gt; 生成后的产物代码 &gt; 无法显示代码（具体参见下面的<strong>不同质量的源码示例</strong>小节）。对应对质量产生影响的预设关键字优先级为 souce-map = eval-source-map &gt; cheap-module- &gt; cheap- &gt; eval = none &gt; nosource-。</li>
<li>构建的速度：再次构建速度都要显著快于初次构建速度。不同环境下关注的速度也不同：

<ul>
<li>在开发环境下：一直开着 devServer，再次构建的速度对我们的效率影响远大于初次构建的速度。从结果中可以看到，eval- 对应的 EvalSourceMapDevToolPlugin 整体要快于不带 eval- 的 SourceMapDevToolPlugin。尤其在质量最佳的配置下，eval-source-map 的再次构建速度要远快于其他几种。而同样插件配置下，不同质量配置与构建速度成反比，但差异程度有限，更多是看具体项目的大小而定。</li>
<li>在生产环境下：通常不会开启再次构建，因此相比再次构建，初次构建的速度更值得关注，甚至对构建速度以外因素的考虑要优先于对构建速度的考虑，这一部分我们在之后的构建优化的课程里会再次讨论到。</li>
</ul></li>
<li>包的大小和生成方式：在开发环境下我们并不需要关注这些因素，正如在开发环境下也通常不考虑使用分包等优化方式。我们需要关注速度和质量来保证我们的高效开发体验，而其他的部分则是在生产环境下需要考虑的问题。</li>
</ul>

<h4 id="不同质量的源码示例">不同质量的源码示例</h4>

<ul>
<li>源码且包含列信息</li>
</ul>

<p><img src="assets/CgqCHl85_KuANSVfAADSE8VO7Qg572.png" alt="Drawing 1.png" /></p>

<ul>
<li>源码不包含列信息</li>
</ul>

<p><img src="assets/Ciqc1F85_LCAMTlgAADhqpZ4v9o628.png" alt="Drawing 2.png" /></p>

<ul>
<li>Loader转换后代码</li>
</ul>

<p><img src="assets/CgqCHl85_LqAPrYzAADfmUwS_JE006.png" alt="Drawing 3.png" /></p>

<ul>
<li>生成后的产物代码</li>
</ul>

<p><img src="assets/CgqCHl85_MGAHhmMAAKGwvDeXIM418.png" alt="Drawing 4.png" /></p>

<h4 id="开发环境下-source-map-推荐预设">开发环境下 Source Map 推荐预设</h4>

<p>在这里我们对开发环境下使用的推荐预设做一个总结（生产环境的预设我们将在之后的构建效率篇中再具体分析）：</p>

<ul>
<li>通常来说，开发环境首选哪一种预设取决于 source map 对于我们的帮助程度。</li>
<li>如果对项目代码了如指掌，查看产物代码也可以无障碍地了解对应源代码的部分，那就可以关闭 devtool 或使用 eval 来获得最快构建速度。</li>
<li>反之如果在调试时，需要通过 source map 来快速定位到源代码，则优先考虑使用 <strong>eval-cheap-modulesource-map</strong>，它的质量与初次/再次构建速度都属于次优级，以牺牲定位到列的功能为代价换取更快的构建速度通常也是值得的。</li>
<li>在其他情况下，根据对质量要求更高或是对速度要求更高的不同情况，可以分别考虑使用 <strong>eval-source-map</strong> 或 <strong>eval-cheap-source-map</strong>。</li>
</ul>

<p>了解了开发环境下如何选择 source map 预设后，我们再来补充几种工具和脚手架中的默认预设：</p>

<ul>
<li>Webpack 配置中，如果不设定 devtool，则使用默认值 eval，即速度与 devtool:false 几乎相同、但模块代码后多了 sourceURL 以帮助定位模块的文件名称。</li>
<li><a href="https://github.com/facebook/create-react-app/blob/fa648daca1dedd97aec4fa3bae8752c4dcf37e6f/packages/react-scripts/config/webpack.config.js" target="_blank">create-react-app 中</a>，在生产环境下，根据 shouldUseSourceMap 参数决定使用‘source-map’或 false；在开发环境下使用‘cheap-module-source-map’（不包含列信息的源代码，但更快）。</li>
<li><a href="https://github.com/vuejs/vue-cli/blob/36f961e43dc76705878659247b563e2af83138ce/packages/%40vue/cli-service/lib/commands/serve.js" target="_blank">vue-cli-service 中</a>，与 creat-react-app 中相同。</li>
</ul>

<p>除了上面讨论的这些简单的预设外，Webpack 还允许开发者直接使用对应插件来进行更精细化的 source map 控制，在开发环境下我们首选的还是 EvalSourceMapDevToolPlugin。下面我们再来看看如何直接使用这个插件进行优化。</p>

<h4 id="evalsourcemapdevtoolplugin-的使用">EvalSourceMapDevToolPlugin 的使用</h4>

<p>在 EvalSourceMapDevToolPlugin 的 <a href="https://webpack.js.org/plugins/eval-source-map-dev-tool-plugin/" target="_blank">传入参数</a>中，除了上面和预设相关的 filename、append、module、columns 外，还有影响注释内容的 moduleFilenameTemplate 和 protocol，以及影响处理范围的 test、include、exclude。这里重点看处理范围的参数，因为通常我们需要调试的是开发的业务代码部分，而非依赖的第三方模块部分。因此在生成 source map 的时候如果可以排除第三方模块的部分而只生成业务代码的 source map，无疑能进一步提升构建的速度，例如示例：</p>

<pre><code>webpack.config.js 

  ... 

  //devtool: 'eval-source-map', 

  devtool: false, 

  plugins: [ 

    new webpack.EvalSourceMapDevToolPlugin({ 

      exclude: /node_modules/, 

      module: true, 

      columns: false 

    }) 

  ], 

  ...
</code></pre>

<p>在上面的示例中，我们将 devtool 设为 false，而直接使用 EvalSourceMapDevToolPlugin，通过传入 module: true 和 column:false，达到和预设 eval-cheap-module-source-map 一样的质量，同时传入 exclude 参数，排除第三方依赖包的 source map 生成。保存设定后通过运行可以看到，在文件体积减小（尽管开发环境并不关注文件大小）的同时，再次构建的速度相比上面表格中的速度提升了将近一倍，达到了最快一级。</p>

<p><img src="assets/CgqCHl85_N2AUkcpAAEqvMKhgVQ549.png" alt="Drawing 5.png" /></p>

<p>类似这样的优化可以帮助我们在一些大型项目中，通过自定义设置来获取比预设更好的开发体验。</p>

<h3 id="总结">总结</h3>

<p>在今天这一课时中，我们主要了解了提升开发效率的另一个重要工具——source map 的用途和使用方法。我们分析了 Webpack 中 devtool 的各种参数预设的组合规则、使用效果及其背后的原理。对于开发环境，我们根据一组示例对比分析来了解通常情况下的最佳选择，也知道了如何直接使用插件来达到更细致的优化。</p>

<p>限于篇幅原因，关于 source map 这一课时还有两个与提效无关的小细节没有提到，一个是生成的 source map 的内容，即浏览器工具是如何将 source map 内容映射回源文件的，如果你感兴趣可以通过这个<a href="http://www.ruanyifeng.com/blog/2013/01/javascript_source_map.html" target="_blank">链接</a>进一步了解；另一个是我们在控制台的网络面板中通常看不到 source map 文件的请求，其原因是出于安全考虑 Chrome 隐藏了 source map 的请求，需要通过 <a href="chrome://net-export/" target="_blank">net-log</a> 来查询。</p>

<p><strong>最后还是留一个小作业</strong>：不知道你有没有留意过自己项目里的 source map 使用的是哪一种生成方式吗？可以根据这一课时的内容对它进行调整和观察效果，也欢迎你在课后留言区讨论项目里对 source map 的优化方案。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#59353535606d6868696e193e34383035773a3634" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f13e9b77f6d9508',t:'MTczNDA3MTUwNC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
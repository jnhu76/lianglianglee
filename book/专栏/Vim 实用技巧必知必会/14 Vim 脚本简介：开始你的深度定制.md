<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=14&#32;Vim&#32;脚本简介：开始你的深度定制>
        <link rel="icon" href="/static/favicon.png">
        <title>14 Vim 脚本简介：开始你的深度定制 </title>
        
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
                        <a class="menu-item" id="00 导读 池建强：Vim 就是四个字“唯快不破”.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/00%20%e5%af%bc%e8%af%bb%20%e6%b1%a0%e5%bb%ba%e5%bc%ba%ef%bc%9aVim%20%e5%b0%b1%e6%98%af%e5%9b%9b%e4%b8%aa%e5%ad%97%e2%80%9c%e5%94%af%e5%bf%ab%e4%b8%8d%e7%a0%b4%e2%80%9d.md">00 导读 池建强：Vim 就是四个字“唯快不破”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 我们为什么要学 Vim？.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e6%88%91%e4%bb%ac%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6%20Vim%ef%bc%9f.md">00 开篇词 我们为什么要学 Vim？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 各平台下的 Vim 安装方法：上路前准备好你的宝马.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/01%20%e5%90%84%e5%b9%b3%e5%8f%b0%e4%b8%8b%e7%9a%84%20Vim%20%e5%ae%89%e8%a3%85%e6%96%b9%e6%b3%95%ef%bc%9a%e4%b8%8a%e8%b7%af%e5%89%8d%e5%87%86%e5%a4%87%e5%a5%bd%e4%bd%a0%e7%9a%84%e5%ae%9d%e9%a9%ac.md">01 各平台下的 Vim 安装方法：上路前准备好你的宝马.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 基本概念和基础命令：应对简单的编辑任务.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/02%20%e5%9f%ba%e6%9c%ac%e6%a6%82%e5%bf%b5%e5%92%8c%e5%9f%ba%e7%a1%80%e5%91%bd%e4%bb%a4%ef%bc%9a%e5%ba%94%e5%af%b9%e7%ae%80%e5%8d%95%e7%9a%84%e7%bc%96%e8%be%91%e4%bb%bb%e5%8a%a1.md">02 基本概念和基础命令：应对简单的编辑任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 更多常用命令：应对稍复杂的编辑任务.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/03%20%e6%9b%b4%e5%a4%9a%e5%b8%b8%e7%94%a8%e5%91%bd%e4%bb%a4%ef%bc%9a%e5%ba%94%e5%af%b9%e7%a8%8d%e5%a4%8d%e6%9d%82%e7%9a%84%e7%bc%96%e8%be%91%e4%bb%bb%e5%8a%a1.md">03 更多常用命令：应对稍复杂的编辑任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 初步定制：让你的 Vim 更顺手.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/04%20%e5%88%9d%e6%ad%a5%e5%ae%9a%e5%88%b6%ef%bc%9a%e8%ae%a9%e4%bd%a0%e7%9a%84%20Vim%20%e6%9b%b4%e9%a1%ba%e6%89%8b.md">04 初步定制：让你的 Vim 更顺手.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 多文件打开与缓冲区：复制粘贴的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/05%20%e5%a4%9a%e6%96%87%e4%bb%b6%e6%89%93%e5%bc%80%e4%b8%8e%e7%bc%93%e5%86%b2%e5%8c%ba%ef%bc%9a%e5%a4%8d%e5%88%b6%e7%b2%98%e8%b4%b4%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">05 多文件打开与缓冲区：复制粘贴的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 窗口和标签页：修改、对比多个文件的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/06%20%e7%aa%97%e5%8f%a3%e5%92%8c%e6%a0%87%e7%ad%be%e9%a1%b5%ef%bc%9a%e4%bf%ae%e6%94%b9%e3%80%81%e5%af%b9%e6%af%94%e5%a4%9a%e4%b8%aa%e6%96%87%e4%bb%b6%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">06 窗口和标签页：修改、对比多个文件的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 正则表达式：实现文件内容的搜索和替换.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/07%20%e6%ad%a3%e5%88%99%e8%a1%a8%e8%be%be%e5%bc%8f%ef%bc%9a%e5%ae%9e%e7%8e%b0%e6%96%87%e4%bb%b6%e5%86%85%e5%ae%b9%e7%9a%84%e6%90%9c%e7%b4%a2%e5%92%8c%e6%9b%bf%e6%8d%a2.md">07 正则表达式：实现文件内容的搜索和替换.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 基本编程支持：规避、解决编程时的常见问题.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/08%20%e5%9f%ba%e6%9c%ac%e7%bc%96%e7%a8%8b%e6%94%af%e6%8c%81%ef%bc%9a%e8%a7%84%e9%81%bf%e3%80%81%e8%a7%a3%e5%86%b3%e7%bc%96%e7%a8%8b%e6%97%b6%e7%9a%84%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98.md">08 基本编程支持：规避、解决编程时的常见问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 七大常用技巧：让编辑效率再上一个台阶.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/09%20%e4%b8%83%e5%a4%a7%e5%b8%b8%e7%94%a8%e6%8a%80%e5%b7%a7%ef%bc%9a%e8%ae%a9%e7%bc%96%e8%be%91%e6%95%88%e7%8e%87%e5%86%8d%e4%b8%8a%e4%b8%80%e4%b8%aa%e5%8f%b0%e9%98%b6.md">09 七大常用技巧：让编辑效率再上一个台阶.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 代码重构实验：在实战中提高编辑熟练度.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/10%20%e4%bb%a3%e7%a0%81%e9%87%8d%e6%9e%84%e5%ae%9e%e9%aa%8c%ef%bc%9a%e5%9c%a8%e5%ae%9e%e6%88%98%e4%b8%ad%e6%8f%90%e9%ab%98%e7%bc%96%e8%be%91%e7%86%9f%e7%bb%83%e5%ba%a6.md">10 代码重构实验：在实战中提高编辑熟练度.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 文本的细节：关于字符、编码、行你所需要知道的一切.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/11%20%e6%96%87%e6%9c%ac%e7%9a%84%e7%bb%86%e8%8a%82%ef%bc%9a%e5%85%b3%e4%ba%8e%e5%ad%97%e7%ac%a6%e3%80%81%e7%bc%96%e7%a0%81%e3%80%81%e8%a1%8c%e4%bd%a0%e6%89%80%e9%9c%80%e8%a6%81%e7%9f%a5%e9%81%93%e7%9a%84%e4%b8%80%e5%88%87.md">11 文本的细节：关于字符、编码、行你所需要知道的一切.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 语法加亮和配色方案：颜即正义.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/12%20%e8%af%ad%e6%b3%95%e5%8a%a0%e4%ba%ae%e5%92%8c%e9%85%8d%e8%89%b2%e6%96%b9%e6%a1%88%ef%bc%9a%e9%a2%9c%e5%8d%b3%e6%ad%a3%e4%b9%89.md">12 语法加亮和配色方案：颜即正义.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 YouCompleteMe：Vim 里的自动完成.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/13%20YouCompleteMe%ef%bc%9aVim%20%e9%87%8c%e7%9a%84%e8%87%aa%e5%8a%a8%e5%ae%8c%e6%88%90.md">13 YouCompleteMe：Vim 里的自动完成.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Vim 脚本简介：开始你的深度定制.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/14%20Vim%20%e8%84%9a%e6%9c%ac%e7%ae%80%e4%bb%8b%ef%bc%9a%e5%bc%80%e5%a7%8b%e4%bd%a0%e7%9a%84%e6%b7%b1%e5%ba%a6%e5%ae%9a%e5%88%b6.md">14 Vim 脚本简介：开始你的深度定制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 插件荟萃：不可或缺的插件.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/15%20%e6%8f%92%e4%bb%b6%e8%8d%9f%e8%90%83%ef%bc%9a%e4%b8%8d%e5%8f%af%e6%88%96%e7%bc%ba%e7%9a%84%e6%8f%92%e4%bb%b6.md">15 插件荟萃：不可或缺的插件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 终端和 GDB 支持：不离开 Vim 完成开发任务.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/16%20%e7%bb%88%e7%ab%af%e5%92%8c%20GDB%20%e6%94%af%e6%8c%81%ef%bc%9a%e4%b8%8d%e7%a6%bb%e5%bc%80%20Vim%20%e5%ae%8c%e6%88%90%e5%bc%80%e5%8f%91%e4%bb%bb%e5%8a%a1.md">16 终端和 GDB 支持：不离开 Vim 完成开发任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="拓展1 纯文本编辑：使用 Vim 书写中英文文档.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e6%8b%93%e5%b1%951%20%e7%ba%af%e6%96%87%e6%9c%ac%e7%bc%96%e8%be%91%ef%bc%9a%e4%bd%bf%e7%94%a8%20Vim%20%e4%b9%a6%e5%86%99%e4%b8%ad%e8%8b%b1%e6%96%87%e6%96%87%e6%a1%a3.md">拓展1 纯文本编辑：使用 Vim 书写中英文文档.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="拓展2 C 程序员的 Vim 工作环境：C 代码的搜索、提示和自动完成.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e6%8b%93%e5%b1%952%20C%20%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%20Vim%20%e5%b7%a5%e4%bd%9c%e7%8e%af%e5%a2%83%ef%bc%9aC%20%e4%bb%a3%e7%a0%81%e7%9a%84%e6%90%9c%e7%b4%a2%e3%80%81%e6%8f%90%e7%a4%ba%e5%92%8c%e8%87%aa%e5%8a%a8%e5%ae%8c%e6%88%90.md">拓展2 C 程序员的 Vim 工作环境：C 代码的搜索、提示和自动完成.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="拓展3 Python 程序员的 Vim 工作环境：完整的 Python 开发环境.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e6%8b%93%e5%b1%953%20Python%20%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%20Vim%20%e5%b7%a5%e4%bd%9c%e7%8e%af%e5%a2%83%ef%bc%9a%e5%ae%8c%e6%95%b4%e7%9a%84%20Python%20%e5%bc%80%e5%8f%91%e7%8e%af%e5%a2%83.md">拓展3 Python 程序员的 Vim 工作环境：完整的 Python 开发环境.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="拓展4 插件样例分析：自己动手改进插件.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e6%8b%93%e5%b1%954%20%e6%8f%92%e4%bb%b6%e6%a0%b7%e4%be%8b%e5%88%86%e6%9e%90%ef%bc%9a%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%e6%94%b9%e8%bf%9b%e6%8f%92%e4%bb%b6.md">拓展4 插件样例分析：自己动手改进插件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="拓展5 其他插件和技巧：吴咏炜的箱底私藏.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e6%8b%93%e5%b1%955%20%e5%85%b6%e4%bb%96%e6%8f%92%e4%bb%b6%e5%92%8c%e6%8a%80%e5%b7%a7%ef%bc%9a%e5%90%b4%e5%92%8f%e7%82%9c%e7%9a%84%e7%ae%b1%e5%ba%95%e7%a7%81%e8%97%8f.md">拓展5 其他插件和技巧：吴咏炜的箱底私藏.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 Vim 森林探秘，一切才刚刚开始.md" href="/%e4%b8%93%e6%a0%8f/Vim%20%e5%ae%9e%e7%94%a8%e6%8a%80%e5%b7%a7%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a/%e7%bb%93%e6%9d%9f%e8%af%ad%20Vim%20%e6%a3%ae%e6%9e%97%e6%8e%a2%e7%a7%98%ef%bc%8c%e4%b8%80%e5%88%87%e6%89%8d%e5%88%9a%e5%88%9a%e5%bc%80%e5%a7%8b.md">结束语 Vim 森林探秘，一切才刚刚开始.md</a>
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
                            <h1 id="title" data-id="14 Vim 脚本简介：开始你的深度定制" class="title">14 Vim 脚本简介：开始你的深度定制</h1>
                            <div><p>你好，我是吴咏炜。</p>

<p>学到今天，我们已经看到了很多的 Vim 脚本，只是还没有正式地把它作为一门语言来介绍。今天，我就正式向你介绍把 Vim 的功能粘合到一起的语言——Vim 脚本（Vim script）。掌握 Vim 脚本的基本语法之后，你就可以得心应手地定制你的 Vim 环境啦。</p>

<h2 id="语法概要">语法概要</h2>

<p>首先，我们需要知道，通过命令行模式执行的命令就是 Vim 脚本。它是一种图灵完全的脚本语言：图灵完全，说明它的功能够强大，理论上可以完成任何计算任务；脚本语言，说明它不需要编译，可以直接通过解释方式来执行。</p>

<p>当然，这并没有说出 Vim 脚本的真正特点。下面，我们就通过各个不同的角度，进行了解，把 Vim 脚本这头“大象”的基本形状完整地摸出来。</p>

<p>在这一讲里，我们改变一下惯例，除非明确说“正常模式命令”，否则用代码方式显示的都是脚本文件里的代码或者命令行模式命令，也就是说，它们前面都不会加 <code>:</code>。毕竟我们这一讲介绍的全是 Vim 脚本，而不是正常模式的快捷操作。</p>

<h3 id="打印输出和字符串">打印输出和字符串</h3>

<p>学习任何一门语言，我们常常以“Hello world!”开始。对于 Vim 脚本，我们不妨也这样——毕竟，打印是一种重要的调试方式，尤其对于没有专门调试器的脚本语言来说。</p>

<p>Vim 脚本的“Hello world!”是下面这样的：</p>

<pre><code>echo 'Hello world!'
</code></pre>

<p><code>echo</code> 是 Vim 用来显示信息的内置命令，而 <code>'Hello world!'</code> 是一个字符串字面量。Vim 里也可以使用 <code>&quot;</code> 来引起一个字符串。<code>'</code> 和 <code>&quot;</code> 的区别和在 shell 里比较相似，前者里面不允许有任何转义字符，而后者则可以使用常见的转义字符序列，如 <code>\n</code> 和 <code>\u....</code> 等。和 shell 不同的是，我们可以在 <code>'</code> 括起的字符里把 <code>'</code> 重复一次来得到这个字符本身，即 <code>'It''s'</code> 相当于 <code>&quot;It's&quot;</code>。不过，在这个例子里，显然还是后者更清晰了。</p>

<p>因为 <code>&quot;</code> 还有开始注释的作用，一般情况下我推荐在 Vim 脚本里使用 <code>'</code>，除非你需要转义字符序列或者需要把 <code>'</code> 本身放到字符串里。</p>

<p>字符串可以用 <code>.</code> 运算符来拼接。由于字典访问也可以用 <code>.</code> ，为了避免歧义，Bram 推荐开发者在新的 Vim 脚本中使用 <code>..</code> 来拼接。但要注意，这个写法在 Vim 7 及之前的版本里不支持。我目前仍暂时使用 <code>.</code> 进行字符串拼接，并和其他大部分运算符一样，前后空一格。这样跟不空格的字典用法比起来，差异就相当明显了。</p>

<p>除了 <code>echo</code>，Vim 还可以用 <code>echomsg</code>（缩写 <code>echom</code>）命令，来显示一条消息。跟 <code>echo</code> 不同的是，这条消息不仅会显示在屏幕上，还会保留在消息历史里，可以在之后用 <code>message</code> 命令查看。</p>

<h3 id="变量">变量</h3>

<p>跟大部分语言一样，Vim 脚本里有变量。变量可以用 <code>let</code> 命令来赋值，如下所示：</p>

<pre><code>let answer = 42
</code></pre>

<p>然后你当然就可以使用 <code>answer</code> 这个变量了，如：</p>

<pre><code>echo 'The meaning of life, the universe and everything is ' . answer
</code></pre>

<p>Vim 的变量可以手工取消，需要的命令是 <code>unlet</code>。在你写了 <code>unlet answer</code> 之后，你就不能再读取 <code>answer</code> 这个变量了。</p>

<h3 id="数字">数字</h3>

<p>上面的赋值语句用到了整数。Vim 脚本里的数字支持整数和浮点数，在大部分平台上，两者都是 64 位的有符号数字类型，基本对应于大部分 C 语言环境里的 <code>int64_t</code> 和 <code>double</code>。表示方式也和 C 里面差不多：整数可以使用 <code>0</code>（八进制）、<code>0b</code>（二进制）和 <code>0x</code>（十六进制）前缀；浮点数含小数点（不可省略），可选使用科学计数法。</p>

<h3 id="复杂数据结构">复杂数据结构</h3>

<p>Vim 脚本内置支持的复杂数据结构是列表（list）和字典（dictionary）。这两者都和 Python 里的对应数据结构一样。对于 C++ 的程序员来说，列表基本上就是数组/array/vector，但大小可变，而且可以直接使用方括号表达式来初始化，如：</p>

<pre><code>let primes = [2, 3, 5, 7, 11, 13, 17, 19]
</code></pre>

<p>然后你可以用下标访问，比如用 <code>primes[0]</code> 就可以得到 <code>2</code>。</p>

<p>字典基本上就是 map，可以使用花括号来初始化，如：</p>

<pre><code>let int_squares = {
      \0: 0,
      \1: 1,
      \2: 4,
      \3: 9,
      \4: 16,
      \}
</code></pre>

<p>键会自动转换成字符串，而值会保留其类型。上面也用到了 Vim 脚本的续行——下一行的第一个非空白字符如果是 <code>\</code>，则表示这一行跟上一行在逻辑上是同一行，这一点和大部分其他语言是不同的。</p>

<p>访问字典里的某一个元素可以用方括号（跟大部分语言一样），如 <code>int_squares['2']</code>；或使用 <code>.</code>，如 <code>int_squares.2</code>。</p>

<h3 id="表达式">表达式</h3>

<p>跟大部分编程语言类似，Vim 脚本的表达式里可以使用括号，可以调用函数（形如 <code>func(…)</code>），支持加（<code>+</code>）、减（<code>-</code>）、乘（<code>*</code>）、除（<code>/</code>）和取模（<code>%</code>），支持逻辑操作（<code>&amp;&amp;</code>、<code>||</code> 和 <code>!</code>），还支持三元条件表达式（<code>a ? b : c</code>）。前面我们已经学过，可以使用 <code>[]</code> 访问列表成员，可以使用 <code>[]</code> 或 <code>.</code> 访问字典的成员，也可以使用 <code>.</code> 或 <code>..</code> 进行字符串拼接。<code>==</code> 和 <code>!=</code> 运算符对所有类型都有效，而 <code>&lt;</code>、<code>&gt;=</code> 等运算符对整数、浮点数和字符串都有效。</p>

<p>对于文本处理，常见的情况是我们使用 <code>=~</code> 和 <code>!~</code> 进行正则表达式匹配，前者表示匹配的判断，后者表示不匹配的判断。比较操作符可以后面添加 <code>#</code> 或 <code>?</code> 来强制进行大小写敏感或不敏感的匹配（缺省受 Vim 选项 <code>ignorecase</code> 影响）。表达式的左侧是待匹配的字符串，右侧则是用来匹配的正则表达式。</p>

<p>注意表达式不是一个合法的 Vim 命令或脚本语句。在表达式的左侧，需要有 <code>echo</code> 这样的命令。如果你只想调用一个函数，而不需要使用其返回的结果，则应使用 <code>call func(…)</code> 这样的写法。</p>

<p>此外，我们在插入模式和命令行模式下都可以使用按键 <code>&lt;C-R&gt;=</code>（两个键）后面跟一个表达式来使用表达式的结果。在替换命令中，我们在 <code>\=</code> 后面也同样可以跟一个表达式，来表示使用该表达式的结果。比如，下面的命令可以在当前编辑文件的每一行前面插入行号和空格：</p>

<pre><code>:%s/^/\=line('.') . ' '/
</code></pre>

<p><code>line</code> 是 Vim 的一个内置函数，<code>line('.')</code> 表示“当前”行的行号，剩下部分你应该直接就明白了吧？</p>

<h3 id="控制结构">控制结构</h3>

<p>作为一门完整的编程语言，标准的控制结构当然也少不了。Vim 支持标准的 <code>if</code>、<code>while</code> 和 <code>for</code> 语句。语法上，Vim 的写法有点老派，跟当前的主流语言不太一样，每种结构都要用一个对应的 <code>endif</code>、<code>endwhile</code> 和 <code>endfor</code> 来结束，如下面所示：</p>

<pre><code>&quot; 简单条件语句
if 表达式
  语句
endif

&quot; 有 else 分支的条件语句
if 表达式
  语句
else
  语句
endif

&quot; 更复杂的条件语句
if 表达式
  语句
elseif 表达式
  语句
else
  语句
endif

&quot; 循环语句
while 表达式
  语句
endwhile
</code></pre>

<p>在 <code>while</code> 和 <code>for</code> 循环语句里，你可以使用 <code>break</code> 来退出循环，也可以使用 <code>continue</code> 来跳过循环体内的其他语句。作为一个程序员，理解它们肯定没有任何困难。</p>

<p>Vim 脚本的 <code>for</code> 语句跟 Python 非常相似，形式是：</p>

<pre><code>for var in object
  这儿可以使用 var
endfor
</code></pre>

<p>表示遍历 <code>object</code>（通常是个列表）对象里面的所有元素。</p>

<p>哦，跟 Python 一样，Vim 脚本也没有 <code>switch/case</code> 语句。</p>

<h3 id="函数和匿名函数">函数和匿名函数</h3>

<p>为了方便开发，函数肯定也是少不了的。Vim 脚本里定义函数使用下面的语法：</p>

<pre><code>function 函数名(参数1, 参数2, ...)
  函数内容
endfunction
</code></pre>

<p>Vim 里用户自定义函数必须首字母大写（和内置函数相区别），或者使用 <code>s:</code> 表示该函数只在当前脚本文件有效。<code>...</code> 可以出现在参数列表的结尾，表示可以传递额外的无名参数。使用有名字的参数时，你需要加上 <code>a:</code> 前缀。要访问额外参数，则需要使用 <code>a:1</code>、<code>a:2</code> 这样的形式。特殊名字 <code>a:0</code> 表示额外参数的数量，<code>a:000</code> 表示把额外参数当成列表来使用，因而 <code>a:000[0]</code> 就相当于 <code>a:1</code>。</p>

<p>在函数里面，跟大部分语言一样，你可以使用 <code>return</code> 命令返回一个结果，或提前结束函数的执行。</p>

<p>Vim 脚本里允许匿名函数，形式是 <code>{逗号分隔开的参数 -&gt; 表达式}</code>。如果你对函数式编程完全没有概念，你可以跳过匿名函数。如果你喜欢函数式编程，那你应该会很欣喜地看到，在 Vim 脚本里可以使用类似下面的语句：</p>

<pre><code>echo map(range(1, 5), {idx, val -&gt; val * val})
</code></pre>

<p>结果是 <code>[1, 4, 9, 16, 25]</code>。跟常见的 <code>map</code> 函数不同，Vim 会传过去两个参数，分别是列表索引和值；同时，它会修改列表的内容。不想修改的话，要把列表复制一份，如 <code>copy(mylist)</code>。</p>

<h2 id="vim-特性">Vim 特性</h2>

<p>上面描述的只是一般性的编程语言语法，但 Vim 脚本如果只当作通用编程语言来用的话，就没啥意义了。我们使用 Vim 脚本，肯定是为了和 Vim 进行交互。下面我们就来仔细检查一下 Vim 脚本里的 Vim 特性。</p>

<h3 id="变量的前缀">变量的前缀</h3>

<p>我们上面已经提到了变量的 <code>a:</code> 前缀。变量的前缀实际上有更多，通用编程概念上很容易理解的是下面四个：</p>

<ul>
<li><code>a:</code> 表示这个变量是函数参数，只能在函数内使用。</li>
<li><code>g:</code> 表示这个变量是全局变量，可以在任何地方访问。</li>
<li><code>l:</code> 表示这个变量是本地变量，但一般这个前缀不需要使用，除非你跟系统的某个名字发生了冲突。</li>
<li><code>s:</code> 表示这个变量（或函数，它也能用在函数上）只能用于当前脚本，有点像 C 里面的 static 变量和函数，只在当前脚本文件有效，因而不会影响其他脚本文件里定义的有冲突的名字。</li>
</ul>

<p>一般编程语言里没有的，是下面这些前缀：</p>

<ul>
<li><code>b:</code> 表示这个变量是当前缓冲区的，不同的缓冲区可以有同名的 <code>b:</code> 变量。比如，在 Vim 里，<code>b:current_syntax</code> 这个变量表示当前缓冲区使用的语法名字。</li>
<li><code>w:</code> 表示这个变量是当前窗口的，不同的窗口可以有同名的 <code>w:</code> 变量。</li>
<li><code>t:</code> 表示这个变量是当前标签页的，不同的标签页可以有同名的 <code>t:</code> 变量。</li>
<li><code>v:</code> 表示这个变量是特殊的 Vim 内置变量，如 <code>v:version</code> 是 Vim 的版本号，等等（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/eval.html#v:var" target="_blank"><code>:help v:var</code></a>）。</li>
</ul>

<p>还有下面这些前缀，可以让我们像使用变量一样使用环境变量和 Vim 选项：</p>

<ul>
<li><code>$</code> 表示紧接着的名字是一个环境变量。注意，一些环境变量是由 Vim 自己设置的，如 <code>$VIMRUNTIME</code>。</li>
<li><code>&amp;</code> 表示紧接着的名字是一个选项，比如， <code>echo &amp;filetype</code> 和 <code>set filetype?</code> 效果相似，都能用来显示当前缓冲区的文件类型。</li>
<li><code>&amp;g:</code> 表示访问一个选项的全局（global）值。对于有本地值的选项，如 <code>tabstop</code>，我们用 <code>&amp;tabstop</code> 直接读到的是本地值了，要访问全局值就必须使用 <code>&amp;g:tabstop</code>。</li>
<li><code>&amp;l:</code> 表示访问一个选项的本地（local）值。对于有本地值的选项，如 <code>tabstop</code>，我们用 <code>&amp;tabstop</code> 直接读到的已经是本地值了，但修改则和 <code>set</code> 一样，同时修改本地值和全局值。使用 <code>&amp;l:</code> 前缀可以允许我们仅修改本地值，像 <code>setlocal</code> 命令一样。</li>
</ul>

<p>你可能要问，什么时候我们会需要用变量形式来访问选项，而不是使用 <code>set</code>、<code>setlocal</code> 这样的命令呢？答案是，当我们需要计算出选项值的时候。<code>set filetype=cpp</code> 基本上和 <code>let &amp;filetype = 'cpp'</code> 等效，我们需要注意到后者里面 <code>cpp</code> 是个字符串，可以是通过某种方式算出来的。光使用 <code>set</code>，就不方便做到这样的灵活性了。</p>

<h3 id="重要命令">重要命令</h3>

<p>Vim 里有很多命令，很多我们已经介绍过，或者直接在 vimrc 配置文件里使用了。这节里我们会介绍跟 Vim 脚本相关性比较大的一些命令。</p>

<p>首先是 <code>execute</code>（缩写 <code>exe</code>），它能用来把后面跟的字符串当成命令来解释。跟上一节使用选项还是 <code>&amp;</code> 变量一样，这样做可以增加脚本的灵活性。除此之外，它还有两种常见特殊用法：</p>

<ul>
<li>在使用键盘映射等场合、需要在一行里放多个命令时，一般可以使用 <code>|</code> 来分隔，但某些命令会把 <code>|</code> 当成命令的一部分（如 <code>!</code>、<code>command</code>、<code>nmap</code> 和用户自定义命令），这种时候就可以使用 <code>execute</code> 把这样的命令包起来，如：<code>exe '!ls' | echo 'See file list above'</code>。</li>
<li><code>normal</code> 命令把后面跟的字符直接当成正常模式命令解释，但如果其中包含有特殊字符时就不方便了。这时可以用 <code>execute</code> 命令，然后在 <code>&quot;</code> 里可以使用转义字符。我们上面讲字符串时没说的是，按键也可以这样转义，比如，<code>&quot;\&lt;C-W&gt;&quot;</code> 就代表 Ctrl-W 这个按键。所以，如果你想在脚本中控制切换到下一个窗口，可以写成：<code>exe &quot;normal \&lt;C-W&gt;w&quot;</code>。</li>
</ul>

<p>然后，我要介绍一下 <code>source</code>（缩写 <code>so</code>）命令。它用来载入一个 Vim 脚本文件，并执行其中的内容。我们已经多次在 vimrc 配置文件中使用它来载入系统提供的 Vim 脚本了，如：</p>

<pre><code>source $VIMRUNTIME/vimrc_example.vim
…
command! PackUpdate packadd minpac | source $MYVIMRC | call minpac#update('', {'do': 'call minpac#status()'})
…
</code></pre>

<p>这里要注意的地方是，要允许一个文件被 <code>source</code> 多次，是需要一些特殊处理的。我目前给出的 vimrc 配置文件由于需要被载入多次，进行了下面的特殊处理：</p>

<ul>
<li>清除缺省自动命令组里当前的所有命令，以免定义的自动命令被执行超过一次</li>
<li>使用 <code>command!</code> 来定义命令，避免重复命令定义的错误</li>
<li>使用 <code>function!</code> 来定义函数，避免重复函数定义的错误</li>
<li>没有手工设置 <code>set nocompatible</code>，因为该设置可能会有较多的副作用（在 defaults.vim 里会确保只设置该选项一次）</li>
</ul>

<p>上面我已经展示了一个 <code>command</code> 命令的例子。这个命令允许我们自定义 Vim 的命令，并允许用户来定制自动完成之类的效果（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/map.html#user-commands" target="_blank"><code>:help user-commands</code></a>）。注意这个命令的定义要写在一行里，所以如果命令很长，或者中间出现会吞掉 <code>|</code> 的命令的话，我们就会需要用上 <code>execute</code> 命令了。</p>

<p>最后，我再说明一下我们用过的 <code>map</code> 系列键映射命令（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/map.html#key-mapping" target="_blank"><code>:help key-mapping</code></a>）。这些命令的主干是 <code>map</code>，然后前面可以插入 <code>nore</code> 表示键映射的结果不再重新映射，最前面用 <code>n</code>、<code>v</code>、<code>i</code> 等字母表示适用的 Vim 模式。在绝大部分情况下，我们都会使用带 <code>nore</code> 这种方式，表示结果不再进行映射（排除偶尔偷懒的情况）。但是，如果我们的 <code>map</code> 命令的右侧用到了已有的（如某个插件带来的）键映射，我们就必须使用没有 <code>nore</code> 的版本了。</p>

<h3 id="事件">事件</h3>

<p>和用户主动发起的命令相对应，Vim 里的自动处理依赖于 Vim 里的事件。迄今为止，我们已经遇到了下面这些事件：</p>

<ul>
<li>BufNewFile 事件在创建一个新文件时触发</li>
<li>BufRead（跟 BufReadPost 相同）事件在读入一个文件后触发</li>
<li>BufWritePost 事件在把整个缓冲区写回到文件之后触发</li>
<li>FileType 事件在设置文件类型（<code>filetype</code> 选项）时被触发</li>
</ul>

<p>Vim 里的事件还有很多（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/autocmd.html#autocmd-events-abc" target="_blank"><code>:help autocmd-events-abc</code></a>），我们就不一一介绍了。上面这些是我们最常用的，你应该了解它们的意义。</p>

<h3 id="内置函数">内置函数</h3>

<p>Vim 里内置了很多函数（列表见 <a href="https://yianwillis.github.io/vimcdoc/doc/usr_41.html#function-list" target="_blank"><code>:help function-list</code></a>），可以实现编程语言所需要的基本功能。我们目前用得比较多的是下面这两个：</p>

<ul>
<li><code>exists</code> 用来检测某一符号（变量、函数等）是否已经存在。在 Vim 脚本里最常见的用途是检测某一变量是否已经被定义。</li>
<li><code>has</code> 用来检测某一 Vim 特性（列表见 <a href="https://yianwillis.github.io/vimcdoc/doc/eval.html#feature-list" target="_blank"><code>:help feature-list</code></a>）是否存在。帮助文档里已经描述得很清楚，我就不详细介绍了。你可以对照看一下我们的 vimrc 配置文件里的用法，应该就明白了。</li>
</ul>

<p>Vim 的内置函数真的很多，我也没法一一介绍。你可以稍作浏览，了解其大概，然后在使用中根据需要查询。别忘了，在看 Vim 脚本时，在关键字上按下 <code>K</code> 就可以查看这个关键字的帮助，如下图所示：</p>

<p><img src="assets/3477651fbc0c79f285e9612180cfc030.gif" alt="Fig14.1" title="在 Vim 脚本里使用 K 键查看帮助" /></p>

<h2 id="风格指南">风格指南</h2>

<p>结束 Vim 脚本的介绍之前，我向你推荐一下 Google 出品的 Vim 脚本风格指南，<a href="https://google.github.io/styleguide/vimscriptguide.xml" target="_blank">Google Vimscript Style Guide</a>。写一种语言，有一个风格指南肯定是会有帮助的，尤其对于初学者而言。</p>

<h2 id="python-集成-选学">Python 集成（选学）</h2>

<p>Vim 脚本功能再强大，也还是一种小众的编程语言。所以，Vim 里内置了跟多种脚本语言的集成，包括：</p>

<ul>
<li>Python</li>
<li>Perl</li>
<li>Tcl</li>
<li>Ruby</li>
<li>Lua</li>
<li>MzScheme</li>
</ul>

<p>由于 Python 的高流行度，目前 Vim 插件里常常见到对 Python 的要求——至少我还没有用过哪个插件要求有其他语言的支持。所以，在这儿我就以 Python 为例，简单介绍一下 Vim 对其他脚本语言的支持。各个语言当然有不同的特性，但支持的方式非常相似，可以说是大同小异。</p>

<p>这部分作为选学提供，相当于本讲内部的一个小加餐。Python 程序员一定要把这部分读完，其他同学则可以选择跳到内容小结。</p>

<p>Vim 很早就支持了 Python 2，Vim 的命令 <code>python</code>（缩写 <code>py</code>）就是用来执行 Python 2 的代码的。后来，Vim 也支持了 Python 3，使用 <code>python3</code>（缩写 <code>py3</code>）来执行 Python 3 的代码。鉴于 Python 的代码还是有不少是 2、3 兼容的，Vim 还有命令 <code>pythonx</code>（缩写 <code>pyx</code>）可以自动选择一个可用的 Python 版本来执行。</p>

<p>我在[拓展 3]里给出了一段代码，用 Python 来检测当前目录是不是在一个 Git 库里。我们先用 <code>pythonx</code> 命令定义了一个 Python 函数，然后用 <code>pyxeval</code> 函数来调用该函数。这就是一种典型的使用方式：在 Python 里定义某个功能，然后在 Vim 脚本里调用该功能。这种情况下，Python 部分的代码一般不需要对 Vim 有任何特殊处理，只是简单实现某个特定功能。</p>

<p>下面是另一个小例子，通过 Python 来获得当前时区和协调世界时的时间差值（对于中国，应当返回 <code>␣+0800</code>）：</p>

<pre><code>function! Timezone()
  if has('pythonx')
pythonx &lt;&lt; EOF
import time

def my_timezone():
    is_dst = time.daylight and time.localtime().tm_isdst
    offset = time.altzone if is_dst else time.timezone
    (hours, seconds) = divmod(abs(offset), 3600)
    if offset &gt; 0: hours = -hours
    minutes = seconds // 60
    return '{:+03d}{:02d}'.format(hours, minutes)
EOF
    return ' ' . pyxeval('my_timezone()')
  else
    return ''
  endif
endfunction
</code></pre>

<p>从 <code>pythonx &lt;&lt; EOF</code> 到 <code>EOF</code>，中间是 Python 代码，定义了一个叫 <code>my_timezone</code> 的函数，我们然后调用该函数来获得结果。对于不支持 Python 的情况，我们就直接返回一个空字符串了。</p>

<p>另一种更复杂的情况是，我们的主干处理逻辑就放在 Python 里。这种情况下，我们就需要在 Python 里调用 Vim 的功能了。在 Vim 调用 Python 代码时，Python 可以访问 <code>vim</code> 模块，其中提供多个 Vim 的专门方法和对象，如：</p>

<ul>
<li><code>vim.command</code> 可以执行 Vim 的命令</li>
<li><code>vim.eval</code> 可以对表达式进行估值</li>
<li><code>vim.buffers</code> 代表 Vim 里的缓冲区</li>
<li><code>vim.windows</code> 代表当前标签页里的 Vim 窗口</li>
<li><code>vim.tabpages</code> 代表 Vim 里的标签页</li>
<li><code>vim.current</code> 代表各种 Vim 的“当前”对象（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/if_pyth.html#python-current" target="_blank"><code>:help python-current</code></a>），包括行、缓冲区、窗口等</li>
</ul>

<p>此外，在拓展 2 里我们给出的使用 <code>pyxf</code> 来执行一个 Python 脚本文件，也是一种在 Vim 里调用 Python 的方式（详见 <a href="https://yianwillis.github.io/vimcdoc/doc/if_pyth.html#:pyxfile" target="_blank"><code>:help pyxfile</code></a>）。那段 clang-format 的代码，总体上也就是访问 <code>vim.current.buffer</code> 对象，调用外部命令格式化指定行，然后把修改的内容写回到 Vim 缓冲区里。</p>

<h2 id="内容小结">内容小结</h2>

<p>好了，我们的 Vim 脚本介绍就到这里了。这一讲和大部分其他讲不同，只是给了你一个 Vim 脚本的概览，目的是让你全面了解一下 Vim 脚本，能够读懂一般的 Vim 脚本，而不是真正教会你如何去写脚本。这讲的主要知识点是：</p>

<ul>
<li>Vim 脚本的基本语法，包括变量、数字、字符串、复杂数据结构、表达式、控制结构和函数</li>
<li>Vim 的专门特性，包括变量的前缀、脚本相关命令、Vim 里的事件和内置函数</li>
<li>Vim 脚本风格指南</li>
<li>Vim 对 Python 等其他脚本语言的支持</li>
</ul>

<p>作为一门编程语言，只有在实践中不断操练，才能真正学会它的使用。如果你对 Vim 脚本有兴趣的话，我们下一讲会剖析几个 Vim 脚本来分析一下，让你有更深入的体会。</p>

<h2 id="课后练习">课后练习</h2>

<p>请查看几个现有的 Vim 脚本来仔细分析一下，理解各行的意义。建议可以从我们在 vimrc 配置文件中包含的 vimrc_example.vim 开始，然后查看其中使用的 defaults.vim。别忘了，我们可以使用普通模式快捷键 <code>gf</code> 或 <code>&lt;C-W&gt;f</code> 直接跳转到光标下的文件里。</p>

<p>如果遇到什么问题，欢迎留言和我讨论。我们下一讲再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b2dedede8b8683838285f2d5dfd3dbde9cd1dddf" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12adbdc895ede4',t:'MTczNDA1ODU2Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
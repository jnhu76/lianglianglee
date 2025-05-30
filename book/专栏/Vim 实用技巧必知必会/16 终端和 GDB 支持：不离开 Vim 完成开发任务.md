<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;终端和&#32;GDB&#32;支持：不离开&#32;Vim&#32;完成开发任务>
        <link rel="icon" href="/static/favicon.png">
        <title>16 终端和 GDB 支持：不离开 Vim 完成开发任务 </title>
        
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
                            <h1 id="title" data-id="16 终端和 GDB 支持：不离开 Vim 完成开发任务" class="title">16 终端和 GDB 支持：不离开 Vim 完成开发任务</h1>
                            <div><p>你好，我是吴咏炜。</p>

<p>早在 Vim 和 Emacs 的“圣战”时期，Emacs 有个功能可是 Vim 用户一直暗暗垂涎的，那就是可以集成 GDB 来调试程序。Emacs 之所以能够实现这个功能，是因为它可以模拟一个终端环境，像终端一样跟一个程序进行输入输出的交互。这样一来，我们不离开编辑器，也能调试程序，既可以方便地看到目前执行在源代码的第几行，也可以直接在编辑器里跟执行中的程序进行交互。</p>

<p>很多主流的开发环境都支持类似的功能。但 Vim 一直不支持这样的功能，直到 Vim 8。虽然到得有点晚，但 Vim 也算是厚积薄发，利用 libvterm 给出了完整的终端支持。今天，我们就拿终端窗口支持和 GDB 支持，作为我们最后的技术话题来介绍了。</p>

<h2 id="终端窗口支持">终端窗口支持</h2>

<h3 id="基本用法">基本用法</h3>

<p>使用 <code>:terminal</code>（缩写 <code>:term</code>）命令，我们可以在 Vim 的窗口中运行终端模拟器。基本的用法就是下面两种：</p>

<ul>
<li>使用 <code>:terminal</code>，后面不跟其他命令，分割一个新窗口，并使用默认的 shell 程序进行终端模拟；shell 退出后窗口自动关闭（可用使用命令参数 <code>++noclose</code> 改变这一行为）。</li>
<li>使用 <code>:terminal 命令</code> 的方式，分割一个新窗口，在其中运行指定的命令并进行终端模拟；命令执行完成退出后窗口不自动关闭，保留执行中显示的信息（可用使用命令参数 <code>++close</code> 改变这一行为）。</li>
</ul>

<p>跟其他的多窗口命令一样，<code>:terminal</code> 默认会进行横向分割，但你也可以在 <code>terminal</code> 前面加上 <code>vert</code> 来进行纵向分割，或加上 <code>tab</code> 来把终端窗口打开到一个新的标签页里。</p>

<p>跟 quickfix 窗口里只能看到程序的输出不同，在终端模拟器里我们既可以看到程序的输出，也可以向程序提供输入。同时，这个终端模拟器像一个真正的终端一样，能够支持色彩和其他的文本控制。你甚至可以在里面运行 Vim，就像 <em>Matrix</em> 电影里层层嵌套的世界一样。</p>

<p><img src="assets/c384d46392e63ebbb162c7f8f888b103.png" alt="Fig16.1" title="开了两个终端窗口的 Vim，其中上面那个又再次运行 Vim" /></p>

<p>当然，从实用的角度，我并不建议你这么做——那样可能会让人头昏，并且容易在使用 <code>&lt;C-W&gt;</code> 和 <code>:q</code> 这样的命令时，出现结果跟自己预想不一致的情况。</p>

<p>终端模拟器的行为应当跟普通的终端一致；因此在 Vim 的终端模拟器里，你可以直接使用的命令跟一般的 Vim 窗口很不一样。毕竟，你在终端模拟器里输入 <code>:</code> 时，肯定不是想进入 Vim 的命令行模式吧？这时候，你需要知道下面这些在“终端作业模式”下的特殊命令（完整列表见 <a href="https://yianwillis.github.io/vimcdoc/doc/terminal.html#t_CTRL-W_N" target="_blank"><code>:help t_CTRL-W</code></a>）：</p>

<ul>
<li><code>&lt;C-W&gt;N</code>（注意大写）或 <code>&lt;C-\&gt;&lt;C-N&gt;</code> 退出终端作业模式，进入终端普通模式。这时终端窗口变成一个普通的文本窗口（终端缓冲区），不再显示色彩，但可以像普通的只读窗口一样自由使用，只是不能修改其中的内容而已。按下 <code>a</code> 或 <code>i</code> 可重新激活终端模拟器，进入终端作业模式。</li>
<li><code>&lt;C-W&gt;&quot;</code> 后面跟寄存器号，表示粘贴该寄存器中的内容到终端里。</li>
<li><code>&lt;C-W&gt;:</code> 相当于普通窗口中的 <code>:</code>，执行命令行模式的命令。</li>
<li><code>&lt;C-W&gt;.</code> 可以给终端窗口发送一个普通的 Ctrl-W。</li>
<li><code>&lt;C-W&gt;&lt;C-\&gt;</code> 可以给终端窗口发送一个普通的 Ctrl-\。</li>
<li>大部分的 <code>&lt;C-W&gt;</code> 开始的命令仍然可以使用，如窗口跳转命令（后面跟 <code>j</code>、<code>k</code> 等）、窗口大小调整命令（后面跟 <code>+</code>、<code>_</code> 等），等等。</li>
</ul>

<p>需要注意，终端模拟器里的光标只能用正常终端里的光标移动键来移动，比如在 Bash 默认配置下，可以用 <code>&lt;C-A&gt;</code> 或 <code>&lt;Home&gt;</code> 移到行首，用 <code>&lt;C-E&gt;</code> 或 <code>&lt;End&gt;</code> 移到行尾等。在退出终端作业模式后，光标就只是普通文本窗口的光标，不会影响终端模式里的光标位置——在你按下 <code>a</code> 或 <code>i</code> 时，光标还是在原来的位置，而不是退出终端作业模式后你移动到的新位置。你也不能修改终端缓冲区中的内容。只要稍微仔细想一想，你就知道这些是完全符合逻辑的。</p>

<p>当你从终端窗口切到另外一个窗口时，终端窗口里面的程序仍然在继续运行；如果你不退出终端作业模式的话，终端窗口里面的内容也会持续更新，跟正常的终端行为一致。要结束终端运行的话（而不只是临时退出终端模式），也跟普通的终端情况一下，可使用 <code>exit</code> 命令或 <code>&lt;C-D&gt;</code>。如果由于某种原因无法正常退出终端的话，则可以使用 <code>&lt;C-W&gt;&lt;C-C&gt;</code> 来强行退出。</p>

<h3 id="使用提示">使用提示</h3>

<p>如果你觉得自己不会在终端里另外启动 Vim，似乎也就很少有机会用到 <code>&lt;Esc&gt;</code> 了，那我们干吗不把这个键用作退出终端作业模式呢？说干就干：</p>

<pre><code>tnoremap &lt;Esc&gt;      &lt;C-\&gt;&lt;C-N&gt;
tnoremap &lt;C-V&gt;&lt;Esc&gt; &lt;Esc&gt;
</code></pre>

<p>前缀 <code>t</code> 表示在终端作业模式下的键映射。我们把 <code>&lt;Esc&gt;</code> 映射到我们上面说的退出终端作业模式的快捷键；同时，我们又把 <code>&lt;C-V&gt;&lt;Esc&gt;</code> 这一在终端里等价于 <code>&lt;Esc&gt;</code> 的按键组合映射为 <code>&lt;Esc&gt;</code>，这样万一我们需要 <code>&lt;Esc&gt;</code>，仍然可以用一种较为自然的方式获得这个按键。</p>

<p>遗憾的是，在 Unix 终端的情况下，很多功能键本身包含 <code>&lt;Esc&gt;</code>，因而会误触发这个键映射。对于这种情况，我们使用下面的键映射，用连按两下 <code>&lt;Esc&gt;</code> 退出终端作业模式效果更好：</p>

<pre><code>tnoremap &lt;Esc&gt;&lt;Esc&gt; &lt;C-\&gt;&lt;C-N&gt;
</code></pre>

<p>此外，对于大部分人而言（像 Bram 这样，用 Vim 调试 Vim，不属于大众需求吧），在 Vim 的终端模式里启动 Vim，恐怕是失误的可能性最大。为了防止这样的失误发生，我们可以在 Vim 启动时检查一下，检测这种嵌套的 Vim 使用。你只需要把下面的代码加到 vimrc 配置文件的开头即可：</p>

<pre><code>if exists('$VIM_TERMINAL')
  echoerr 'Do not run Vim inside a Vim terminal'
  quit
endif
</code></pre>

<p>你可以试验一下在 Vim 的终端窗口里再运行 Vim，看一下上面的代码产生的出错效果。</p>

<h3 id="终端的用途">终端的用途</h3>

<p>说了这么多，你可能有点疑惑，单独起一个终端有什么问题吗？我为什么要在 Vim 里运行终端呢？</p>

<p>我是这么理解的：</p>

<ol>
<li><strong>方便。</strong>特别在远程连接的时候，有可能新开一个连接在某些环境里需要特别的认证，比较麻烦。即使连接没有任何障碍，你总还需要重新 cd 到工作目录里吧？而如果在一个现有的 Vim 会话里开一个新的终端，可以一个命令搞定，然后用你已经很熟悉的 Vim 命令在不同的窗口或标签页里切换。</li>
<li><strong>文本。</strong>我们可以从终端作业模式切换到终端普通模式，然后用我们熟悉的 Vim 命令来对缓冲区中的文本进行搜索、复制等处理工作。</li>
<li><strong>控制。</strong>你可以发送命令给终端，也可以读取终端屏幕上的信息。这样，事实上就打开了一片新天地，可以在 Vim 里做很多之前做不到的事情，比如，用 Vim 来比较两个屏幕输出的区别（<a href="https://yianwillis.github.io/vimcdoc/doc/terminal.html#terminal-diff" target="_blank"><code>:help terminal-diff</code></a>）。</li>
</ol>

<p>终端窗口相关的函数名称都以 <code>term_</code> 打头（可以查看帮助文件 <a href="https://yianwillis.github.io/vimcdoc/doc/terminal.html#terminal-function-details" target="_blank"><code>:help terminal-function-details</code></a>）。比如，如果我们想要用程序向缓冲区编号为 2（可以用 <code>:ls</code> 和 <code>:echo term_list()</code> 等命令来检查）的终端发送 <code>ls</code> 命令来显示当前目录下的文件列表的话，我们可以使用（注意转义字符序列要求使用双引号）：</p>

<pre><code>call term_sendkeys(2, &quot;ls\n&quot;)
</code></pre>

<p>下面这个比较无聊的例子，可以用来获取 ~/.vim 目录下的文件清单：</p>

<pre><code>let term_nbr = term_start('bash')
call term_wait(term_nbr, 100)
let line_pos1 = term_getcursor(term_nbr)[0]
call term_sendkeys(term_nbr, &quot;ls ~/.vim|cat\n&quot;)
call term_wait(term_nbr, 500)
let line_pos2 = term_getcursor(term_nbr)[0]
let result = []
let line_pos1 += 1
while line_pos1 &lt; line_pos2
  call add(result, term_getline(term_nbr, line_pos1))
  let line_pos1 += 1
endwhile
call term_sendkeys(term_nbr, &quot;\&lt;C-D&gt;&quot;)
while term_getstatus(term_nbr) != 'finished'
  call term_wait(term_nbr, 100)
endwhile
exe term_nbr . 'bd'
echo join(result, &quot;\n&quot;)
</code></pre>

<p>这当然不是完成这件任务的最好方法，但上面的代码展示了终端相关函数的一些基本用法：</p>

<ol>
<li>我们用 <code>term_start</code> 命令创建一个新的终端，得到终端缓冲区的编号</li>
<li>我们用 <code>term_wait</code> 等待 100 毫秒，待其就绪</li>
<li>我们用 <code>term_getcursor</code> 获取光标的当前行号</li>
<li>我们用 <code>term_sendkeys</code> 发送一个命令到终端上；ls 之后用 cat 是为了防止 ls 看到输出是终端而产生多列的输出</li>
<li>然后我们等待命令执行完成并更新终端</li>
<li>我们获取光标的当前位置，然后用 <code>term_getline</code> 获得上一次的行号和这一次的行号之间的行的内容，放到变量 <code>result</code> 里</li>
<li>我们然后发送一个 <code>&lt;C-D&gt;</code> 到终端，结束作业</li>
<li>然后我们等待到 <code>term_getstaus</code> 返回的状态成为 <code>'finished'</code>，即终端作业已经执行结束</li>
<li>最后我们用缓冲区编号加 <code>bd</code> 命令删除缓冲区（所以屏幕上我们看不到这个终端窗口），并用换行符作为分隔符打印 ls 返回的内容</li>
</ol>

<p>你可以实际测试一下这个脚本，体会一下这些基本功能。比如，可以把脚本存盘为 test.vim，然后用 <code>:so %</code> 来运行。</p>

<h2 id="gdb-支持">GDB 支持</h2>

<p>为什么 Vim 直到最近才支持 GDB 呢？因为这真不是件容易的事情啊。为了能在 Vim 里顺畅地使用 GDB，Bram 需要在 Vim 里实现下面这些不同的功能：</p>

<ul>
<li>终端支持</li>
<li>作业（job）和通道（channel）</li>
<li>窗口工具条、弹出窗口和弹出式菜单</li>
</ul>

<p>有了这些功能之后，Vim 通过一个内置的插件，就可以提供 GDB 的调试支持了。我们可以通过 <code>:packadd termdebug</code> 命令来加载这个插件，然后通过 <code>:Termdebug 可执行程序名称</code> 来调试一个可执行程序。</p>

<p>下面这个动图可以说明最主要的流程：</p>

<p><img src="assets/493d863ebaa07a528b484e2d01093fa4.gif" alt="Fig16.2" title="在 Vim 里进行调试的过程示例" /></p>

<p>我简要说明一下需要注意的几点：</p>

<ul>
<li><code>:Termdebug</code> 命令会把屏幕分成三个区域，从上到下分别是 gdb 命令行，程序输出，以及含调试控制按钮的源代码窗口。</li>
<li>在最上面的 gdb 窗口中，我们可以输入 gdb 的命令，但程序的输出和纯终端使用 gdb 的情况不同，是在中间的窗口输出的。</li>
<li>最下面的的源代码窗口里，我们有五个按钮可以用，允许习惯图形界面的用户使用鼠标进行操作。我们也可以使用鼠标右键直接在源代码行上设置断点。（当然，我们仍然可以在最上面的 gdb 窗口用命令来完成这些任务。）</li>
<li>鼠标在变量上悬停时，可以显示变量的值。只要 gdb 能打印的信息，它就能用浮动提示显示出来。这比手工使用 gdb 的 <code>p</code> 命令还是要方便多了。</li>
</ul>

<p>还有一个需要稍微注意的地方是，如果你在不同的作用域有两个同名变量，那浮动提示只能显示当前作用域的变量的信息，即使你把光标放到不在当前作用域的变量上也是如此。这点上，Vim 还是比较笨的——毕竟它不理解代码。</p>

<h2 id="内容小结">内容小结</h2>

<p>这讲我们介绍了 Vim 8 带来的新功能：终端支持。这个功能给 Vim 打开了一片新的天空。使用终端支持，我们可以不离开 Vim 打开一个或多个新的终端窗口，里面可以模拟真正的终端功能，包括色彩控制。我们可以使用 Vim 命令来处理新的终端缓冲区中的文本。我们还可以利用代码来控制这个终端和读取其中的内容。有了这些支持，Vim 也就顺理成章地支持使用 GDB 像集成开发环境一样地调试程序了。</p>

<p>根据我个人的经验，在使用了这个功能之后，我开启新远程连接比之前少了，而经常在一个服务器上只开一个连接，里面开一个 Vim 来完成所需的任务。编译和执行，可以全部在这个 Vim 会话里完成。</p>

<p>本讲我们的配置文件中加入了针对终端窗口的键映射和防 Vim 重入，对应的标签是 <code>l16-unix</code> 和 <code>l16-windows</code>。</p>

<h2 id="课后练习">课后练习</h2>

<p>请尝试使用 <code>:terminal</code> 命令，打开一个新窗口，并在其中进行操作，然后退出终端作业模式，把终端缓冲区中的内容复制到新的缓冲区中。</p>

<p>如果你使用一种可以用 GDB 调试的编译语言的话，也请你尝试一下使用 <code>:Termdebug</code> 命令进行调试。如果你之前用的是纯命令行的 gdb 的话，这个功能还是有很大的易用性提升的。</p>

<p>最后，同样地，如果有任何问题或疑问，欢迎留言和我讨论！</p>

<p>我是吴咏炜，让我们在告别这个课程之前，再道一次再见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#39555555000d0808090e795e54585055175a5654" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12ae952ba8ede4',t:'MTczNDA1ODU5Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
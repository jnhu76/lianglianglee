<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=15&#32;PowerShell脚本：如何实现文件批量处理的自动化？>
        <link rel="icon" href="/static/favicon.png">
        <title>15 PowerShell脚本：如何实现文件批量处理的自动化？ </title>
        
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
                        <a class="menu-item" id="00 导读 入门Python的必备知识.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%af%bc%e8%af%bb%20%e5%85%a5%e9%97%a8Python%e7%9a%84%e5%bf%85%e5%a4%87%e7%9f%a5%e8%af%86.md">00 导读 入门Python的必备知识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 重复工作这么多，怎样才能提高工作效率？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e9%87%8d%e5%a4%8d%e5%b7%a5%e4%bd%9c%e8%bf%99%e4%b9%88%e5%a4%9a%ef%bc%8c%e6%80%8e%e6%a0%b7%e6%89%8d%e8%83%bd%e6%8f%90%e9%ab%98%e5%b7%a5%e4%bd%9c%e6%95%88%e7%8e%87%ef%bc%9f.md">00 开篇词 重复工作这么多，怎样才能提高工作效率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 拆分与合并：如何快速地批量处理内容相似的Excel？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/01%20%e6%8b%86%e5%88%86%e4%b8%8e%e5%90%88%e5%b9%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%9c%b0%e6%89%b9%e9%87%8f%e5%a4%84%e7%90%86%e5%86%85%e5%ae%b9%e7%9b%b8%e4%bc%bc%e7%9a%84Excel%ef%bc%9f.md">01 拆分与合并：如何快速地批量处理内容相似的Excel？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 善用Python扩展库：如何批量合并多个文档？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/02%20%e5%96%84%e7%94%a8Python%e6%89%a9%e5%b1%95%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%b9%e9%87%8f%e5%90%88%e5%b9%b6%e5%a4%9a%e4%b8%aa%e6%96%87%e6%a1%a3%ef%bc%9f.md">02 善用Python扩展库：如何批量合并多个文档？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 图片转文字：如何提高识别准确率？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/03%20%e5%9b%be%e7%89%87%e8%bd%ac%e6%96%87%e5%ad%97%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e8%af%86%e5%88%ab%e5%87%86%e7%a1%ae%e7%8e%87%ef%bc%9f.md">03 图片转文字：如何提高识别准确率？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  函数与字典：如何实现多次替换.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/04%20%20%e5%87%bd%e6%95%b0%e4%b8%8e%e5%ad%97%e5%85%b8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e5%a4%9a%e6%ac%a1%e6%9b%bf%e6%8d%a2.md">04  函数与字典：如何实现多次替换.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 图像处理库：如何实现长图拼接？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/05%20%e5%9b%be%e5%83%8f%e5%a4%84%e7%90%86%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%95%bf%e5%9b%be%e6%8b%bc%e6%8e%a5%ef%bc%9f.md">05 图像处理库：如何实现长图拼接？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 jieba分词：如何基于感情色彩进行单词数量统计？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/06%20jieba%e5%88%86%e8%af%8d%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%e6%84%9f%e6%83%85%e8%89%b2%e5%bd%a9%e8%bf%9b%e8%a1%8c%e5%8d%95%e8%af%8d%e6%95%b0%e9%87%8f%e7%bb%9f%e8%ae%a1%ef%bc%9f.md">06 jieba分词：如何基于感情色彩进行单词数量统计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 快速读写文件：如何实现跨文件的字数统计？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/07%20%e5%bf%ab%e9%80%9f%e8%af%bb%e5%86%99%e6%96%87%e4%bb%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e8%b7%a8%e6%96%87%e4%bb%b6%e7%9a%84%e5%ad%97%e6%95%b0%e7%bb%9f%e8%ae%a1%ef%bc%9f.md">07 快速读写文件：如何实现跨文件的字数统计？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 正则表达式：如何提高搜索内容的精确度？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/08%20%e6%ad%a3%e5%88%99%e8%a1%a8%e8%be%be%e5%bc%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e6%90%9c%e7%b4%a2%e5%86%85%e5%ae%b9%e7%9a%84%e7%b2%be%e7%a1%ae%e5%ba%a6%ef%bc%9f.md">08 正则表达式：如何提高搜索内容的精确度？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 扩展搜索：如何快速找到想要的文件？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/09%20%e6%89%a9%e5%b1%95%e6%90%9c%e7%b4%a2%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e6%89%be%e5%88%b0%e6%83%b3%e8%a6%81%e7%9a%84%e6%96%87%e4%bb%b6%ef%bc%9f.md">09 扩展搜索：如何快速找到想要的文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 按指定顺序给词语排序，提高查找效率.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/10%20%e6%8c%89%e6%8c%87%e5%ae%9a%e9%a1%ba%e5%ba%8f%e7%bb%99%e8%af%8d%e8%af%ad%e6%8e%92%e5%ba%8f%ef%bc%8c%e6%8f%90%e9%ab%98%e6%9f%a5%e6%89%be%e6%95%88%e7%8e%87.md">10 按指定顺序给词语排序，提高查找效率.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  通过程序并行计算，避免CPU资源浪费.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/11%20%20%e9%80%9a%e8%bf%87%e7%a8%8b%e5%ba%8f%e5%b9%b6%e8%a1%8c%e8%ae%a1%e7%ae%97%ef%bc%8c%e9%81%bf%e5%85%8dCPU%e8%b5%84%e6%ba%90%e6%b5%aa%e8%b4%b9.md">11  通过程序并行计算，避免CPU资源浪费.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 文本处理函数：三招解决数据对齐问题.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/12%20%e6%96%87%e6%9c%ac%e5%a4%84%e7%90%86%e5%87%bd%e6%95%b0%ef%bc%9a%e4%b8%89%e6%8b%9b%e8%a7%a3%e5%86%b3%e6%95%b0%e6%8d%ae%e5%af%b9%e9%bd%90%e9%97%ae%e9%a2%98.md">12 文本处理函数：三招解决数据对齐问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Excel插件：如何扩展Excel的基本功能？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/13%20Excel%e6%8f%92%e4%bb%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%a9%e5%b1%95Excel%e7%9a%84%e5%9f%ba%e6%9c%ac%e5%8a%9f%e8%83%bd%ef%bc%9f.md">13 Excel插件：如何扩展Excel的基本功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 VBA脚本编程：如何扩展Excel，实现文件的批量打印？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/14%20VBA%e8%84%9a%e6%9c%ac%e7%bc%96%e7%a8%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%a9%e5%b1%95Excel%ef%bc%8c%e5%ae%9e%e7%8e%b0%e6%96%87%e4%bb%b6%e7%9a%84%e6%89%b9%e9%87%8f%e6%89%93%e5%8d%b0%ef%bc%9f.md">14 VBA脚本编程：如何扩展Excel，实现文件的批量打印？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 PowerShell脚本：如何实现文件批量处理的自动化？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/15%20PowerShell%e8%84%9a%e6%9c%ac%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%96%87%e4%bb%b6%e6%89%b9%e9%87%8f%e5%a4%84%e7%90%86%e7%9a%84%e8%87%aa%e5%8a%a8%e5%8c%96%ef%bc%9f.md">15 PowerShell脚本：如何实现文件批量处理的自动化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 循环与文件目录管理：如何实现文件的批量重命名？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/16%20%e5%be%aa%e7%8e%af%e4%b8%8e%e6%96%87%e4%bb%b6%e7%9b%ae%e5%bd%95%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%96%87%e4%bb%b6%e7%9a%84%e6%89%b9%e9%87%8f%e9%87%8d%e5%91%bd%e5%90%8d%ef%bc%9f.md">16 循环与文件目录管理：如何实现文件的批量重命名？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 不同操作系统下，如何通过网络同步文件？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/17%20%e4%b8%8d%e5%90%8c%e6%93%8d%e4%bd%9c%e7%b3%bb%e7%bb%9f%e4%b8%8b%ef%bc%8c%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e7%bd%91%e7%bb%9c%e5%90%8c%e6%ad%a5%e6%96%87%e4%bb%b6%ef%bc%9f.md">17 不同操作系统下，如何通过网络同步文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 http库：如何批量下载在线内容，解放鼠标（上）？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/18%20http%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%b9%e9%87%8f%e4%b8%8b%e8%bd%bd%e5%9c%a8%e7%ba%bf%e5%86%85%e5%ae%b9%ef%bc%8c%e8%a7%a3%e6%94%be%e9%bc%a0%e6%a0%87%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9f.md">18 http库：如何批量下载在线内容，解放鼠标（上）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 http库：如何批量下载在线内容，解放鼠标（下）？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/19%20http%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%b9%e9%87%8f%e4%b8%8b%e8%bd%bd%e5%9c%a8%e7%ba%bf%e5%86%85%e5%ae%b9%ef%bc%8c%e8%a7%a3%e6%94%be%e9%bc%a0%e6%a0%87%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9f.md">19 http库：如何批量下载在线内容，解放鼠标（下）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 不同文件混在一起，怎么快速分类？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/20%20%e4%b8%8d%e5%90%8c%e6%96%87%e4%bb%b6%e6%b7%b7%e5%9c%a8%e4%b8%80%e8%b5%b7%ef%bc%8c%e6%80%8e%e4%b9%88%e5%bf%ab%e9%80%9f%e5%88%86%e7%b1%bb%ef%bc%9f.md">20 不同文件混在一起，怎么快速分类？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 SQLite文本数据库：如何进行数据管理（上）？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/21%20SQLite%e6%96%87%e6%9c%ac%e6%95%b0%e6%8d%ae%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e6%95%b0%e6%8d%ae%e7%ae%a1%e7%90%86%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9f.md">21 SQLite文本数据库：如何进行数据管理（上）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 SQLite文本数据库：如何进行数据管理（下）？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/22%20SQLite%e6%96%87%e6%9c%ac%e6%95%b0%e6%8d%ae%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e6%95%b0%e6%8d%ae%e7%ae%a1%e7%90%86%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9f.md">22 SQLite文本数据库：如何进行数据管理（下）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 怎么用数据透视表更直观地展示汇报成果？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/23%20%e6%80%8e%e4%b9%88%e7%94%a8%e6%95%b0%e6%8d%ae%e9%80%8f%e8%a7%86%e8%a1%a8%e6%9b%b4%e7%9b%b4%e8%a7%82%e5%9c%b0%e5%b1%95%e7%a4%ba%e6%b1%87%e6%8a%a5%e6%88%90%e6%9e%9c%ef%bc%9f.md">23 怎么用数据透视表更直观地展示汇报成果？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 条形、饼状、柱状图最适合用在什么场景下？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/24%20%e6%9d%a1%e5%bd%a2%e3%80%81%e9%a5%bc%e7%8a%b6%e3%80%81%e6%9f%b1%e7%8a%b6%e5%9b%be%e6%9c%80%e9%80%82%e5%90%88%e7%94%a8%e5%9c%a8%e4%bb%80%e4%b9%88%e5%9c%ba%e6%99%af%e4%b8%8b%ef%bc%9f.md">24 条形、饼状、柱状图最适合用在什么场景下？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 图表库：想要生成动态图表，用Echarts就够了.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/25%20%e5%9b%be%e8%a1%a8%e5%ba%93%ef%bc%9a%e6%83%b3%e8%a6%81%e7%94%9f%e6%88%90%e5%8a%a8%e6%80%81%e5%9b%be%e8%a1%a8%ef%bc%8c%e7%94%a8Echarts%e5%b0%b1%e5%a4%9f%e4%ba%86.md">25 图表库：想要生成动态图表，用Echarts就够了.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 快速提取图片中的色块，模仿一张大师的照片.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/26%20%e5%bf%ab%e9%80%9f%e6%8f%90%e5%8f%96%e5%9b%be%e7%89%87%e4%b8%ad%e7%9a%84%e8%89%b2%e5%9d%97%ef%bc%8c%e6%a8%a1%e4%bb%bf%e4%b8%80%e5%bc%a0%e5%a4%a7%e5%b8%88%e7%9a%84%e7%85%a7%e7%89%87.md">26 快速提取图片中的色块，模仿一张大师的照片.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 zipfile压缩库：如何给数据压缩&amp;加密备份？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/27%20zipfile%e5%8e%8b%e7%bc%a9%e5%ba%93%ef%bc%9a%e5%a6%82%e4%bd%95%e7%bb%99%e6%95%b0%e6%8d%ae%e5%8e%8b%e7%bc%a9&amp;%e5%8a%a0%e5%af%86%e5%a4%87%e4%bb%bd%ef%bc%9f.md">27 zipfile压缩库：如何给数据压缩&amp;加密备份？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 Celery库：让计算机定时执行任务，解放人力.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/28%20Celery%e5%ba%93%ef%bc%9a%e8%ae%a9%e8%ae%a1%e7%ae%97%e6%9c%ba%e5%ae%9a%e6%97%b6%e6%89%a7%e8%a1%8c%e4%bb%bb%e5%8a%a1%ef%bc%8c%e8%a7%a3%e6%94%be%e4%ba%ba%e5%8a%9b.md">28 Celery库：让计算机定时执行任务，解放人力.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 网络和邮件库：定时收发邮件，减少手动操作.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/29%20%e7%bd%91%e7%bb%9c%e5%92%8c%e9%82%ae%e4%bb%b6%e5%ba%93%ef%bc%9a%e5%ae%9a%e6%97%b6%e6%94%b6%e5%8f%91%e9%82%ae%e4%bb%b6%ef%bc%8c%e5%87%8f%e5%b0%91%e6%89%8b%e5%8a%a8%e6%93%8d%e4%bd%9c.md">29 网络和邮件库：定时收发邮件，减少手动操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 怎么快速把任意文件格式转成PDF，并批量加水印？.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/30%20%e6%80%8e%e4%b9%88%e5%bf%ab%e9%80%9f%e6%8a%8a%e4%bb%bb%e6%84%8f%e6%96%87%e4%bb%b6%e6%a0%bc%e5%bc%8f%e8%bd%ac%e6%88%90PDF%ef%bc%8c%e5%b9%b6%e6%89%b9%e9%87%8f%e5%8a%a0%e6%b0%b4%e5%8d%b0%ef%bc%9f.md">30 怎么快速把任意文件格式转成PDF，并批量加水印？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节特别放送1 实体水果店转线上销售的数据统计问题.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/%e6%98%a5%e8%8a%82%e7%89%b9%e5%88%ab%e6%94%be%e9%80%811%20%e5%ae%9e%e4%bd%93%e6%b0%b4%e6%9e%9c%e5%ba%97%e8%bd%ac%e7%ba%bf%e4%b8%8a%e9%94%80%e5%94%ae%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%9f%e8%ae%a1%e9%97%ae%e9%a2%98.md">春节特别放送1 实体水果店转线上销售的数据统计问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节特别放送2 用自顶至底的思路解决数据统计问题.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/%e6%98%a5%e8%8a%82%e7%89%b9%e5%88%ab%e6%94%be%e9%80%812%20%e7%94%a8%e8%87%aa%e9%a1%b6%e8%87%b3%e5%ba%95%e7%9a%84%e6%80%9d%e8%b7%af%e8%a7%a3%e5%86%b3%e6%95%b0%e6%8d%ae%e7%bb%9f%e8%ae%a1%e9%97%ae%e9%a2%98.md">春节特别放送2 用自顶至底的思路解决数据统计问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节特别放送3 揭晓项目作业的答案.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/%e6%98%a5%e8%8a%82%e7%89%b9%e5%88%ab%e6%94%be%e9%80%813%20%e6%8f%ad%e6%99%93%e9%a1%b9%e7%9b%ae%e4%bd%9c%e4%b8%9a%e7%9a%84%e7%ad%94%e6%a1%88.md">春节特别放送3 揭晓项目作业的答案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 和我一起成为10X效率职场人.md" href="/%e4%b8%93%e6%a0%8f/Python%e8%87%aa%e5%8a%a8%e5%8c%96%e5%8a%9e%e5%85%ac%e5%ae%9e%e6%88%98%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%92%8c%e6%88%91%e4%b8%80%e8%b5%b7%e6%88%90%e4%b8%ba10X%e6%95%88%e7%8e%87%e8%81%8c%e5%9c%ba%e4%ba%ba.md">结束语 和我一起成为10X效率职场人.md</a>
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
                            <h1 id="title" data-id="15 PowerShell脚本：如何实现文件批量处理的自动化？" class="title">15 PowerShell脚本：如何实现文件批量处理的自动化？</h1>
                            <div><p>你好，我是尹会生。</p>

<p>在工作中，会遇到和Windows操作系统紧密结合又需要批量操作的工作需求，比如文件的批量重命名，还有按照扩展名搜索文件。那么今天这节课，我将给你介绍一个主要应用于Windows操作系统自动化的脚本&ndash;PowerShell。</p>

<p>你肯定会有疑问，为什么我们不用Python，而是要另外学习PowerShell脚本呢？原因就在于，遇到这类工作需求，PowerShell脚本会比Python功能更强大，使用更方便，学习起来也更容易。</p>

<p>首先，Windows的所有操作，都有对应的PowerShell操作，可以达到办公自动化的目标。而使用Python，会因为Windows没有提供接口，有些功能就不能完全实现，或者有些操作还需要手动执行。</p>

<p>其次，PowerShell的语法简洁，比Python更加友好，降低了你阅读代码的难度。这一点你在这节课我讲解的例子中会有更深刻的体会。</p>

<p>最后，PowerShell在Windows上能做到开箱即用，安装完成后就可以正常运行了。而Python还需要安装解释器和配置环境。比如在金融和证券领域中，基于公司的规定，你可能无法安装操作系统之外的软件，这时候PowerShell的优势就体现出来了。</p>

<h2 id="什么是powershell">什么是PowerShell？</h2>

<p>PowerShell是<strong>开源的终端命令解释器</strong>，之所以被称作解释器，是因为它会把用户输入的命令翻译给操作系统去执行，也能把操作系统执行的结果返回给用户。</p>

<p>作为PowerShell的初学者，你刚开始理解起来会有点难度，不过别担心，我举个例子你就明白了。</p>

<p>比如，在商场中的娃娃机，你需要通过对摇杆的控制来移动爪子的位置，抓取想要的礼物，因为你自己是没法直接移动爪子的。那么把你移动摇杆的操作翻译成爪子移动操作的功能，这就是解释器。</p>

<p>为了强化你对PowerShell的理解，我再把使用PowerShell前后的情况给你做个对比。</p>

<p>Windows操作系统在日常工作中，我们会采用图形界面（GUI）进行各类办公操作。在没有使用PowerShell的时候，如果我要想添加、删除、修改用户，就得点击各种按钮、选择下拉菜单、鼠标右键点击用户图标。</p>

<p>想象一下，如果我给100个新员工创建账户，创建一个账户需要5分钟，那仅仅在用户添加这一个操作上，我就需要花费500分钟，效率是不是很低？</p>

<p>但是利用PowerShell，就可以自动化完成<strong>用户添加</strong>这一重复性的操作了。</p>

<p>从Windows Server 2012版本（一听就是个很古老的操作系统版本）开始，所有的GUI操作就完全可以基于PowerShell实现了。这意味着，你能用GUI界面实现Windows操作系统中的所有功能。</p>

<p>你看，微软已经给了我们这么现成的自动化工具，所以我们可以放心大胆地把Windows操作系统上的重复操作，都交给PowerShell，提高工作效率。像这些和Windows操作系统紧密结合又需要批量操作的工作需求，都是PowerShell最适用的工作场景。</p>

<p>那么接下来，我就以Windows中最常用的文件管理为例，为你讲解一下PowerShell如何对文件进行批量处理。</p>

<h2 id="怎样使用powershell实现文件批量处理">怎样使用PowerShell实现文件批量处理？</h2>

<p>在对文件的批量处理中，我们经常会遇到两类场景。</p>

<ol>
<li>文件夹下的文件按照某一规则批量重命名。比如，为了项目文件名称工整，需要批量给文件夹/文件按照某些规律进行重命名。</li>
<li>将符合多个条件的扩展名从大量文件中找出来。比如，为了找到outlook里附件中的所有的压缩包，你需要使用搜索功能，一次性查找包含.zip .rar .7z多个扩展名的文件。</li>
</ol>

<p>接下来就以文件批量的重命名，来为你讲解一下用Powershell怎么实现。</p>

<h3 id="一行代码-实现文件批量重命名">一行代码，实现文件批量重命名</h3>

<p>使用PowerShell脚本实现文件的批量重命名，我们只需要一行代码。不过在这一行代码中，需要你掌握的PowerShell的概念比较多。所以我会通过一个案例，帮助你在学习的同时也能迅速掌握、内化。</p>

<p>首先我通过一行代码创建了10个文本文件，创建的PowerShell脚本是：</p>

<pre><code>foreach($num in (1..10)) { New-Item $num&quot;.txt&quot; -type file }
</code></pre>

<p>通过这行脚本，我在当前目录下创建了名字为“1.txt-10.txt”这样10个文件。由于批量改名操作也会用到类似的语法，所以我先给你详细讲解一下<strong>批量创建文件的脚本</strong>。这个脚本包含了<strong>循环、变量和命令三个部分</strong>。</p>

<p>第一个部分，是我在脚本中使用的<strong>foreach的循环功能。</strong></p>

<p>循环的运行次数由in关键字后面的次数确定，这里我为了创建10个文本文件，就使用了“1..10”的形式来产生1到10这十个数字，这样的话，foreach就会执行10次，而每一次循环num变量都会得到一个新的数字，并在foreach循环体“{ }”中被New-Item项执行一次。</p>

<p>第二个需要你掌握的部分叫做<strong>自定义变量。</strong></p>

<p>在PowerShell中，变量名都是以”<span class="math inline">\(”开头的，剩余字符可以是字母、数字和下划线等任意字符。并且PowerShell的变量名是不区分大小写的，比如“\)</span>num”和“<span class="math inline">\(NUM”都表示相同的变量。我在代码中使用了“\)</span>num”自定义变量来记录每次循环的数字。</p>

<p>第三个需要掌握的部分叫做<strong>命令</strong>。</p>

<p>PowerShell的命令有三种类型，例如New-Item叫做Cmdlet类型， foreach()叫做工作流类型，还有PowerShell内置的函数和自定义函数类型。这三种命令类型有一个共同的特点，就是都<strong>能被PowerShell执行</strong>。那么为了讲解方便，我在这节课中把这三种命令类型统称为“<strong>命令</strong>”。</p>

<p>在代码中，“New-Item”命令用来创建新的文本文件。我在创建时使用了两个参数，第一个参数<strong>“$num&rdquo;.txt&rdquo;”用来指定文件名</strong>，文件名为num变量和字符串“.txt”连接后的名字。第二个参数<strong>“-type”用来指定创建的类型是文件</strong>（而非文件夹）。</p>

<p><strong>PowerShell的脚本编写完成之后，接下来我们就需要运行它。</strong></p>

<p>它的执行方式非常简单。在Windows中使用快捷键“win+R”可以弹出运行界面，在运行界面输入<strong>“powershell”</strong>之后再按回车，便会弹出PowerShell的交互界面，此时我们只需要把代码输入到交互界面就行了。</p>

<p>如果你希望脚本内容能多次运行，你就可以把脚本保存为以<strong>“ps1”作为扩展名的文件中。</strong>例如，我可以把创建10个文件的脚本保存到“newtxt.ps1”文件中，反复执行时，可以进入PowerShell交互界面，并输入“.\路径\newtxt.ps1”来运行它。</p>

<p>我把执行的命令和结果截图，贴在下方供你参考。</p>

<p><img src="assets/a9b104f36a914522823244423dfedddb.jpg" alt="" /></p>

<p>创建了10个文本文件后，我想把它们统一进行改名，新的名字是“new_数字_new.txt”，即在这10个文件的文件名前后，都增加“new”字符串，改名之后的格式如下图：</p>

<p><img src="assets/230b94906410493eb5ea82d235a852a2.jpg" alt="" /></p>

<p>我把批量改名的代码从截图中单独拿出来，来重点讲解一下其中的<strong>dir命令、管道符和Rename-Item命令及其参数，因为PowerShell就是通过管道符将多个命令组合在一起，实现批量改名的。</strong></p>

<pre><code>dir *.txt | foreach { Rename-Item $_ -NewName (&quot;new_&quot;+$_.BaseName+&quot;_new.txt&quot;)  }

</code></pre>

<p>我来带你按照代码的执行顺序，从左到右，依次讲解这行代码的主要实现逻辑。</p>

<p>首先，这行代码最左侧的<strong>dir命令</strong>，它的功能是<strong>在终端输出当前目录下所有文件和目录的名称。</strong></p>

<p>如果你要操作的文件没有在当前目录，可以使用“cd 要访问的目录”方式进入该目录下，再使用“dir”命令查看。为了更改文本文件的名称，我需要先通过“dir *.txt”来获得当前目录下所有以“.txt”结尾的文件和目录。</p>

<p>接下来是这行代码中出现的<strong>“｜”，被称作管道符</strong>。就像它的名字一样，用一条管道把两个程序连接在一起。它的作用主要有两个：</p>

<ol>
<li>连接两条命令，PowerShell会从左到右依次执行。如果左侧程序没有执行完成，右侧程序则会进入阻塞状态，等待左侧程序执行完成后再执行。</li>
<li>把左侧命令的执行（输出）结果，通过管道，作为右侧程序的输入。在这行代码中，“dir”命令的输出，会通过管道符成为“Rename-Item”命令的输入，再通过foreach循环，实现对每个文本文件进行重命名的操作。</li>
</ol>

<p>最后，我们来学习一下“Rename-Item”命令和它的参数。和“New-Item”一样，“Rename-Item”也是<strong>PowerShell独有的Cmdlet类型命令</strong>，它的功能是<strong>实现文件的改名</strong>。</p>

<p>PowerShell中一共有9个与“Item”相关的Cmdlet，你可以使用下面的命令获得所有的命令和它的帮助。</p>

<pre><code>PS&gt; Get-Command -Noun Item

CommandType     Name                            Definition
-----------     ----                            ----------
Cmdlet          Clear-Item                      Clear-Item [-Path] &lt;String[]...
Cmdlet          Copy-Item                       Copy-Item [-Path] &lt;String[]&gt;...
Cmdlet          Get-Item                        Get-Item [-Path] &lt;String[]&gt; ...
Cmdlet          Invoke-Item                     Invoke-Item [-Path] &lt;String[...
Cmdlet          Move-Item                       Move-Item [-Path] &lt;String[]&gt;...
Cmdlet          New-Item                        New-Item [-Path] &lt;String[]&gt; ...
Cmdlet          Remove-Item                     Remove-Item [-Path] &lt;String[...
Cmdlet          Rename-Item                     Rename-Item [-Path] &lt;String&gt;...
Cmdlet          Set-Item                        Set-Item [-Path] &lt;String[]&gt; ...
</code></pre>

<p><strong>“Item”</strong>被官方文档翻译为“项”，你可以和我一样，把它们理解为<strong>要操作的对象</strong>，通过上面这些Cmdlet命令，就可以实现这些“项”的增删改查。当你需要改名，就可以用“Rename-Item” Cmdlet来实现。</p>

<p>我们了解了“Rename-Item”命令的功能后，再继续学习它的<strong>参数格式和内容</strong>。它的参数格式是： <strong>“Rename-Item</strong> 旧的文件名 <strong>-NewName</strong> 新的文件名<strong>”</strong>。我在代码中使用了这样一行命令对文件名称进行了修改：</p>

<pre><code>Rename-Item $_ -NewName (&quot;new_&quot;+$_.BaseName+&quot;_new.txt&quot;) 
</code></pre>

<p>在“Rename-Item”命令的参数中，你需要关注<strong>“$_”这个内置变量</strong>。它的功能是<strong>表示当前对象。</strong>在我们的例子中，它表示foreach每次循环时，通过“｜”传入的文件名称。</p>

<p>不过当你需要改名时，还有另外一种方法，那就是使用“<span class="math inline">\(_.BaseName”来获得文件(不包含扩展名的)基本名称，以及通过“\)</span>_.extension”来获得扩展名。例如：“10.txt”文件的基本名称是“10”，扩展名是“.txt”。当我把“10.txt”赋值给“<span class="math inline">\(_”之后，就可以使用“\)</span>_.BaseName”取得基本名称“10”，使用“$_.extension”取得扩展名“.txt”。</p>

<p>在“Rename-Item”命令的参数“(&ldquo;new_&rdquo;+$_.BaseName+&rdquo;_new.txt&rdquo;)”，这部分代码中，除了提取了“10.txt”的基本名称“10”之外，我还利用“+”符号把&rdquo;new&rdquo;字符串和基本名称“10”连接，形成新的文件名“new_10_new.txt”字符串。也就是说，把将、新的字符串作为“Rename-Item”的参数，就实现了文件重命名。</p>

<p>在这里有一点需要你注意，“<span class="math inline">\(\_”从PowerShell3.0版本开始，就可以使用“\)</span>PSItem”</p>

<p>替代“<span class="math inline">\(_”，虽然“\)</span>PSItem”在阅读上更加友好，但“<span class="math inline">\(_”能够向后兼容，而且输入的内容更少。所以我会更推荐你使用“\)</span>_”。</p>

<p>通过对批量改名代码的例子与分析，你会发现虽然只有一行代码，但是其中包含了内置变量、循环、管道和重命名命令等功能的组合，PowerShell就是这样一种简洁而功能强大的脚本语言。</p>

<p>为了让你能对PowerShell更加得心应手，我再为你讲解一个利用同样逻辑，就可以实现的场景：从一个文件夹找到.zip和.rar扩展名的文件，带你来一起看一下怎么使用PowerShell一行命令，来按扩展名搜索文件。</p>

<h3 id="按扩展名搜索文件怎么实现">按扩展名搜索文件怎么实现</h3>

<p>当工作时间久了，你会发现自己的Outlook文件夹里有海量的附件，特别是压缩包占用空间非常大。如果我们逐一搜索，就会像大海捞针一样，找到自己想要的文件非常困难。那么此时，我们就可以按照扩展名通过一行脚本将它们全部找出来进行备份或整理到新的文件夹中。</p>

<p>要想实现按扩展名搜索文件，我们需要继续利用“dir”命令来取得当前目录下所有文件，以及利用“$_.extension”来取得文件的扩展名。</p>

<p>所以我首先使用如下命令，来获取当前目录下的扩展名都有哪些，便于接下来的搜索和查找。</p>

<pre><code>dir | foreach{$_.extension} | Get-Unique

执行结果
.txt
.zip
.rar
</code></pre>

<p>这行代码会把当前目录下所有文件的扩展名显示出来，为了避免重复的扩展名会影响我后续的脚本编写，那么在这里，我使用了“Get-Unique”命令，把相同的扩展名去重，只保留一个。</p>

<p>代码的执行过程依然是从左向右依次执行。</p>

<ul>
<li>首先，通过“dir”命令获取当前目录下所有的文件名称，把所有的文件名称作为“foreach”循环的输入，赋值给“$_”。</li>
<li>然后再通过“$_.extension”取得每个文件的扩展名，之后进行输出，再次通过管道符，作为“Get-Unique”命令的输入。</li>
<li>最后，“Get-Unique”会把所有扩展名作为输入处理，输出不重复的扩展名。</li>
</ul>

<p>通过观察所有扩展名的脚本执行结果，确定要搜索的目录中，包含了我们需要的“.zip”和“.rar”之后，我们再取出符合扩展名的文件名称。为了实现这一功能，我们还需要使用一个新的命令：<strong>“Where-Object”。</strong></p>

<p><strong>“Where-Object”</strong>也是Cmdlet类型的命令，它的用法是<strong>对管道中的每个对象进行筛选，把不符合条件的对象删除</strong>。那怎么判断管道中的对象是否符合条件呢？依据就是“Where-Object”命令后面“{}”中的参数。</p>

<p>“{}”中可以使用命令、也可以使用比较运算符。如果“{}”中使用命令，且命令执行结果为True，则“Where-Object”会保留对象进行输出或传入后续的管道，如果命令结果为False，则删除对象。</p>

<p>此外，如果“{}”中使用了比较运算符，就根据比较运算符的结果对对象进行相应操作。同命令结果一样，如果为True，则保留对象；如果为False，则删除对象。</p>

<p>比较运算符一共有10个，它是以“-”开头，跟着运算符名称，运算符的名称是大于、等于、小于的首字母缩写，我把比较运算符整理成表格，供你参考：</p>

<p><img src="assets/a62e1fe816a14dd7a737d29083db4da9.jpg" alt="" /></p>

<p>为了便于你理解比较运算符，我通过一个例子，来给你讲解怎么搜索“.zip”扩展名的文件。</p>

<pre><code>dir | Where-Object{ $_.extension -eq &quot;.zip&quot; }
</code></pre>

<p>在这段代码中，我使用了“-eq”比较运算符，实现了文件扩展名是否和“.zip”字符串相等的判断。</p>

<p>脚本通过dir命令可以获得当前目录下所有的文件，再通过“Where-Object”命令，依次对管道中的对象进行判断。如果扩展名为“.zip”则比较的结果为True，“Where-Object”命令执行完成后，PowerShell会在终端显示该文件名称，否则就会删除该对象，不在终端进行显示。</p>

<p>如果你还需要对文件大小、文件名称相似性等其他方式比较，可以参考比较运算符表格进行实现。</p>

<p>我们实现了单个扩展名的搜索之后，再来让这一行脚本的功能继续增强，让它能搜索多个扩展名。</p>

<p>为了同时搜索到“.zip”和“.rar”文件，我们需要<strong>引入逻辑运算符</strong>。在“Where-Object”命令的参数中，使用逻辑运算符，然后根据它两侧的执行结果是否为True，实现对扩展名的组合判断。例如下面的脚本：</p>

<pre><code>dir | Where-Object{ ($_.extension -eq &quot;.zip&quot; ) -or ($_.extension -eq &quot;.rar&quot; ) }
</code></pre>

<p>这行代码，过滤了管道中文件的扩展名是否包含“.zip”或者包含“.rar”。我在代码中使用的“-or”逻辑运算符的含义是，它的两侧扩展名和“.zip”、“.rar”任意一个相等，返回结果就是True，“Where-Object”命令就会把文件名称打印到屏幕上，如果返回的结果为False，对象仍然会被删除掉。</p>

<p>标准逻辑运算符我也为你整理了一个表格，你可以参考。</p>

<p><img src="assets/09630bdc4b3845d189456dd62074aa83.jpg" alt="" /></p>

<p>总结来说，通过增加逻辑运算符，我们就能实现对多个扩展名的搜索功能。</p>

<h2 id="小结">小结</h2>

<p>最后我来为你做个总结，我通过批量改名和按多个扩展名搜索文件这两个例子，为你展示了PowerShell脚本的命令和管道。</p>

<p>命令中的Cmdlet类型能够支持Windows中所有的对象，其中对文件操作最常用的四个Cmdlet类型命令是：New-Item、Rename-Item、Where-Object和Get-Unique。它们也是我在这节课中为你着重介绍的命令。在PowerShell中，你掌握的Cmdlet越多，能够实现的功能就越强大，因此，我把<a href="https://docs.microsoft.com/zh-cn/powershell/scripting/samples/manipulating-items-directly?view=powershell-7.1" target="_blank">官方文档</a>地址提供给你，你可以通过扩展学习，掌握更多的Cmdlet命令。</p>

<p>此外，命令中的工作流类型可以实现循环，我在这节课给你介绍了foreach循环，通过foreach就可以实现批量操作，优化办公效率。</p>

<p>不过在你掌握了足够多的命令后，还需要掌握管道符。因为通过管道符能够连接命令，让命令可以按顺序执行，而通过对命令的组合，就可以实现多个命名的自动化运行，今儿实现Windows操作系统相关操作的自动化。</p>

<h2 id="思考题">思考题</h2>

<p>按照惯例，我来为你留一道思考题，你能否通过官方文档查找到删除项的命令，并实现“.txt”扩展名文件的批量删除？请你大胆尝试一下。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ee828282d7dadfdfded9ae89838f8782c08d8183" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f11ea5128774194',t:'MTczNDA1MDU1Ny4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=02&#32;善用Python扩展库：如何批量合并多个文档？>
        <link rel="icon" href="/static/favicon.png">
        <title>02 善用Python扩展库：如何批量合并多个文档？ </title>
        
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
                            <h1 id="title" data-id="02 善用Python扩展库：如何批量合并多个文档？" class="title">02 善用Python扩展库：如何批量合并多个文档？</h1>
                            <div><p>你好，我是尹会生。</p>

<p>在日常工作中，我们打交道最多的文件就要数Word和Excel了。我们经常面临这么一种场景：需要将Excel的内容合并到Word中。你可以想一想，完成这个需求，需要手动进行几个步骤的操作呢？很显然，有4步。</p>

<ul>
<li>首先，要手动打开Excel、Word文件；</li>
<li>接着，复制一个单元格的文字到Word指定位置；</li>
<li>然后，如果有多个单元格，就需要重复复制多次；</li>
<li>最后，保存Word文件，并关闭Excel和Word文件。</li>
</ul>

<p>如果只有两个文件，这几步手动操作一定不成问题，不会耗费太多的时间。但是如果文件特别多，哪怕只有十几个，手动操作就相当耗费时间了，而且一不小心还容易出错。幸运的是，现在我们可以通过Python来实现批量文件合并功能，你只需要执行一个Python程序就能搞定所有文件的合并操作。</p>

<p>所以今天这节课，我们先从比较简单的内容讲起：用Python自动合并两个Word文件。然后再进阶，学习如何合并Word和其他类型的文件。一步一步来，相信你会掌握得既牢固又扎实。</p>

<h2 id="手工操作和用python操作的区别">手工操作和用Python操作的区别</h2>

<p>首先我们要知道，为什么在合并文件的时候用Python更高效。我用一个例子来给你讲解手工操作和用Python操作的区别。比如下面这一段文字：</p>

<blockquote>
<p>简易流<strong>程</strong>——集团原则在每年<strong>1</strong><strong>～</strong><strong>5</strong>月、<strong>7</strong><strong>～</strong><strong>11</strong>月生产工作任务较重或<strong>考核时间较紧</strong>的情况下运用。</p>
</blockquote>

<p>在Word文件中，不但有文字内容，还有加粗、红色等格式，而且这些特殊的格式和文字内容是混合在一起的。</p>

<p>但如果用Python来读取Word文件，这段文字会被分为纯文字、段落、字体、字号以及表格等更具体的部分，而且每一个部分都对应着Python的变量和函数。</p>

<p>这样一来，你就可以根据自己的需求，只提取某一具体部分的内容。比如说，你看到哪一段文字的字体很好看，就可以直接读取之后套用到新的文字段落就行，非常便捷。</p>

<p>不过，用Python读取文件时，你需要记住很多个Python变量和函数。当然了，针对这一点，你也不用担心，这些变量和函数在Python的Word扩展库<a href="https://python-docx.readthedocs.io/en/latest/" target="_blank">官方文档</a>可以查看。所以如果你需要某个功能，但是不知道应该用什么变量和函数名称时，可以在官方文档中找到它的名字和描述信息。</p>

<p>总之，用Python读取文件的方式是非常有助于提高工作效率的。所以接下来我们用Python处理Word文件时，就需要通过刚才介绍的变量和函数来替代手动操作。</p>

<p>接下来，我先带你学习怎样用Python合并多个Word文件，然后再讲怎么把Word文件与纯文本、图片和Excel进行合并。</p>

<h2 id="如何合并多个word文件-只保留文字内容">如何合并多个Word文件，只保留文字内容？</h2>

<p>先从合并两个Word文件说起。假设你现在有两个Word文件，需要进行文件中的文字合并操作。两个文件的内容分别是:</p>

<p>文件一（内容包含字体、字号、颜色等额外信息）：</p>

<p><strong>1.简易流程</strong>——集团原则在每年<strong>1</strong><strong>～</strong><strong>5</strong>月、<strong>7</strong><strong>～</strong><strong>11</strong>月生产工作任务较重或<strong>考核时间较紧</strong>的情况下运用。</p>

<p>文件二（内容文字出现在表格中）：</p>

<p><img src="assets/ffb6a4b572b24728a917aa1d5fa52ee0.jpg" alt="" /></p>

<p>我把两个文件的信息总结如下：</p>

<ul>
<li>第一个文件中，字体使用了黑体和宋体字，此外还有红色字体和加粗等格式。现在我需要只提取其中的文字内容，不带任何格式。</li>
<li>第二个文件中，文字被放在了一张表格里。现在我需要用Python把表格中的文字提取出来，合并成一个新的文件。</li>
</ul>

<p>我先把核心代码给你贴出来，然后再给你详细讲解具体的操作方法。</p>

<p>首先，我们可以使用一段Python代码提取Word文件里的内容，然后合并成一个文件。</p>

<pre><code>import docx
def merge_without_format(docx_files: list):
    '''
    只获取内容进行合并
    '''
    # 遍历每个文件
    for docx_file in sorted(docx_files):
        another_doc = Document(docx_file)
        # 获取每个文件的所有“段落”
        paras = another_doc.paragraphs
        # 获取所有段落的文字内容
        # paras_content = [para.text for para in paras]
        for para in paras:
            # 为新的word文件创建一个新段落
            newpar = doc.add_paragraph('')
            # 将提取的内容写入新的文本段落中
            newpar.add_run(para.text)

    # 所有文件合并完成后在指定路径进行保存
    doc.save(Path(word_files_path, 'new.docx'))
        

# 调用函数
merge_without_format(files)
</code></pre>

<p>在这段代码中你可以看到，我使用了一个Python的扩展库，它叫做python-docx，这也是我想重点给你讲解的一个扩展库。</p>

<p>python-docx是Python专门用来编辑Word文档的库，我在实现Word文档自动化操作的工作中经常会用到它。使用它的好处就是不必自己研究docx文件类型的底层实现细节，你可以像操作.txt文本一样直接打开、修改和保存关闭文件。可以说，python-docx扩展库降低了用户使用Python的复杂度。</p>

<p>我再举个例子展示一下具体的操作过程。例如python-docx库支持函数Document，它实现了Word文件的打开功能，底层也做了很多对Word格式的处理工作，让你可以直接使用paragraphs变量就能读取一整段Word文件。</p>

<p>Document函数格式如下：</p>

<pre><code>Document(docx_file)
</code></pre>

<p>还有，函数save也是python-docx扩展库提供的word文件保存函数。同样的，save函数在底层也做了很多对docx格式兼容的操作。像下面的代码一样，你就可以直接给这个函数传递一个文件路径，然后进行保存。是不是降低了编写代码的难度呢?</p>

<pre><code>doc.save(Path(word_files_path, 'new.docx'))
</code></pre>

<p>通过我举的例子，你就能更直观地感受到Python扩展库的方便之处了。接下来我们再回到刚才那两个文件的合并，合并之后的结果如下：</p>

<pre><code>1.简易流程——集团原则在每年1-5月、7-11月生产工作任务较重或考核时间较紧的情况下运用。
人力资源、生产、品管、财务等部门整理、提供绩效考核数据。
人力资源部门收集各部门提供的考核数据，依据员工绩效考核评分标准对集团所有员工进行绩效考核得分计算。
人力资源部门将核计的员工绩效考核结果提交部门经理确认后报集团主管领导核定。
</code></pre>

<p>现在你已经掌握了两个Word文件的合并方法了。但如果我想让这段程序适用于三个、四个，甚至更多个Word文件的合并，那该怎么操作呢?</p>

<p>一个好消息就是，上面的代码我们不需要做任何修改，就可以合并多个Word文件。因为我使用了一个叫做<strong>函数</strong>的功能。<strong>函数有时候也被称作过程、方法</strong>，*<strong><em>它的作用是将那些需要反复使用的代码组合在一起</em></strong>*。</p>

<p>之前我们使用过函数，这些函数是Python自带的，或是扩展库提供的。这些函数我们可以直接拿来使用，使用函数在计算机术语中被称作函数调用。通过函数我们可以实现程序的模块化，多次使用可以多次调用，从而减少代码的重复性。但如果你需要自己编写函数怎么办呢?</p>

<p>你可以将重复的代码功能写在自己定义的函数中，在需要使用的地方调用就可以了。这种自己编写的函数就被称作自定义函数。自定义函数和Python自带的函数一样，也可以实现减少代码重复性的作用。</p>

<p>关于自定义函数，你需要熟悉它们的相关语法，主要是函数名、函数定义和调用方法。通常编写一个函数要为函数并起一个名字，这个名字叫做函数名。当你需要使用函数的功能时，可以使用函数名加“()”的方式来使用它，而且使用一个函数一般被称作调用函数。</p>

<p>我把函数定义和函数调用的写法单独拿出来给你看下，定义函数的格式是def后面跟着函数名称，调用函数是函数名称后面跟着一个&rdquo;()&rdquo; ，这是它的语法格式：</p>

<pre><code># 定义一个函数
def 函数名(参数列表):
    函数体
# 调用一个函数
函数名(参数)
</code></pre>

<p>知道了自定义函数的语法，接下来我们就可以在程序中使用自定义函数了。</p>

<p>举个简单的例子。像我在合并Word文件的程序中，第2行的merge_without_format就是我定义的一个函数，第24行merge_without_format(files) 就是对函数进行调用，files叫做函数的参数。通过函数参数，可以在调用函数的时候为函数指定要操作的对象。</p>

<p>相信你不难发现，使用函数以后，不但可以提高代码的重复利用率，还能提高代码的可读性。</p>

<p>那这段程序是怎样处理多个文件合并的呢? 我在调用函数merge_without_format时，使用了files变量作为参数，而files变量包含了大量的文件。因为是多个文件合并，所以在函数中我使用了一个小技巧，就是你熟悉的for循环语句，for循环语句能够遍历files变量的值，这样就可以将files指向的全部Word文件逐一进行文件内容的提取，进行两两合并，从而实现任意多个Word文件的合并操作。</p>

<p>通过对多个word合并，我希望你能学会怎么提取Word中的文字内容，如果你需要编写大量重复的代码，可以将它们写成自己定义的函数。</p>

<h2 id="怎样合并不同类型的文件">怎样合并不同类型的文件？</h2>

<p>通过上面的例子，我们实现了Word文件之间的合并。在工作中，我们经常需要处理Word和Txt文件、图片、Excel这些类型合并的情况，又该如何操作呢？接下来，我就一个一个来讲一讲。</p>

<h3 id="将纯文本和word文件合并">将纯文本和Word文件合并</h3>

<p>如果是为了支持信息丰富，我们Word和Txt合并之后保存到新的Word中，会出现Txt里的字体字号和原有文件不统一的问题，我们可以使用python-docx扩展库为Txt文件中的文字增加格式。</p>

<p>如果合并前Word文件是仿宋字体，而且有下划线和红色字体，我们将Txt合并之后如何进行字体、样式和颜色的统一呢？我们可以使用下面这段代码。</p>

<pre><code>def add_content_mode1(content):
    '''
    增加内容
    '''
    para = doc.add_paragraph().add_run(content)
    # 设置字体格式
    para.font.name = '仿宋'
    # 设置下划线
    para.font.underline = True
    # 设置颜色
    para.font.color.rgb = RGBColor(255,128,128)  
</code></pre>

<p>首先，我定义了一个叫做add_content_mode1的函数。考虑到Word合并Txt是否会有多个Txt进行合并操作，所以我使用自定义函数功能。</p>

<p>当你需要对多个Txt进行合并，就调用函数依次对它们进行处理，这样你就不用编写重复的代码了，这也是我在编写代码时进行提效的一个小技巧。</p>

<p>接下来，我们将每个新合并的txt内容作为一个新的段落合并到原有的文字中，这个功能使用python-docx的add_paragraph函数就可以增加了一个新的段落。</p>

<p>最后，把这一段所有文字设置成和原有的Word统一的字体、下划线和颜色，保证新的段落在格式上的统一。</p>

<p>在具体操作的时候，我还要提醒你，Word文件支持的格式丰富程度远远高于Txt文件，所以当这两种格式丰富程度不一致的文件进行合并时，要么向下兼容，去掉Txt不支持的格式；要么向上兼容，对Txt进行格式再调整。否则容易出现合并之后仍需要手动调整格式的问题，影响工作效率。</p>

<h3 id="将图片和word文件合并">将图片和Word文件合并</h3>

<p>我们再来看一下第2种情况，怎么把图片和Word文件进行合并呢?</p>

<p>想一下，我们经常见到的图片格式就有.jpg、.png、.gif等，由于这些格式应用范围广，格式没有被商业软件加密，所以python-docx库的add_picture函数就能实现把图片插入Word的功能。代码如下：</p>

<pre><code>from docx import Document
from docx import shared

doc = Document()
# 按英寸设置宽度,添加图片
doc.add_picture('test.jpg', width=shared.Inches(1)) 
</code></pre>

<p>那有没有被商业保护、不能直接支持的格式呢？比如Pohotshop自带的.ps格式，我们如果将.ps格式插入Word文档，.ps格式不能被add_picture所支持，就只能以附件的形式添加到Word文件中，作为附件添加的文件无法直接展示图片的内容，和add_picture相比不够直观。</p>

<p>所以如果不需要进行内容的加密等商业目的的时候，建议使用通用和公开格式，这些格式对编程语言的兼容性更好。</p>

<p>总的来说，python-docx的功能非常强大，除了将文本和图片合并到Word文件中，还可以和第一节课我们学过的xlrd扩展库相配合，将Excel和Word进行合并。</p>

<h3 id="将excel和word文件合并">将Excel和Word文件合并</h3>

<p>为了让你更好地理解如何进行Word和Excel文件的合并，我用一个利用Excel和Word批量制作邀请函的例子来给你讲解。</p>

<p>我在Word中保存了邀请函的标准公文格式，但是其中的被邀请人、性别（先生、女士）以及发出邀请的时间，分别用“&lt;姓名&gt;”“&lt;性别&gt;”“&lt;时间&gt;”替代。邀请函格式如下：</p>

<p><strong>尊敬的 &lt;姓名&gt; &lt;性别&gt;：</strong></p>

<p><strong>… 邀请函内容 …</strong></p>

<p><strong>&lt;今天日期&gt;</strong></p>

<p>我在Excel的每一行中写了被邀请人的姓名、性别信息。格式如下：</p>

<p><img src="assets/101df2cb498f463d9bb25eb22023692d.jpg" alt="" /></p>

<p>现在，我们需要将Excel和Word进行合并操作，为每个被邀请人自动生成一个Word格式的邀请函。</p>

<p>虽然Word中自带的邮件功能可以批量制作邀请函，但是在灵活性还是较差的。比如我要在邀请函制作完成的时候自动添加制作时间等功能，就无法通过Word自带的邮件功能实现。接下来我就用Python来生成邀请函，代码如下：</p>

<pre><code>def generat_invitation():
    '''
    生成邀请函文件
    '''
    doc = Document(invitation)
    # 取出每一段
    for para in doc.paragraphs:
        for key, value in replace_content.items():
            if key in para.text:
                # 逐个关键字进行替换
                para.text = para.text.replace(key, value)

    file_name = PurePath(invitation_path).with_name(replace_content['&lt;姓名&gt;']).with_suffix('.docx')
    doc.save(file_name)
</code></pre>

<p>对于这个问题，我是这样思考的。如果手动操作，我需要：</p>

<ul>
<li>先将Excel中的每一行中的姓名、性别填入Word文件中；</li>
<li>再将当前日期填入到Word文件中；</li>
<li>最后再按照姓名另存为一个文件。</li>
</ul>

<p>但如果使用Python来实现呢？就会非常简单。首先在整个过程中，Word文档是被反复使用到的，所以对Word文档进行修改的这个动作，我会将它写入到循环语句当中。</p>

<p>接着，我需要一个循环语句来处理Excel里的每一行循环，因为我们需要把Excel的每一行读取出来，然后替换“&lt;姓名&gt;”“&lt;性别&gt;”。</p>

<p>最后我们要解决的就是替换问题了。python-docx功能非常强大，它自带了替换函数–replace函数，能够将&rdquo;&lt;姓名&gt;“”&lt;性别&gt;&ldquo;替换成Excel真实的用户和性别。</p>

<p>我再用代码解释一下。对应上面的代码：</p>

<ul>
<li>第7行的for循环实现了遍历每个段落功能，para变量就是表示每个段落的变量。</li>
<li>第8行我们将excel提前处理为python的基础类型–字典（链接）， for循环实现了姓名、性别的遍历。</li>
<li>第11行实现了内容的替换功能。</li>
<li>第12行我将姓名作为文件名称，将.docx作为扩展名指定为新的文件名称，通过第13行的save函数进行了邀请函的保存。</li>
</ul>

<p>你看，多次读取Word文件的循环、多次按行读取Excel文件的循环、替换的函数都有了，那我们就可以实现自动化生成邀请函的功能了。</p>

<p>最终每张邀请函实现的效果如下图：</p>

<p><img src="assets/4b35740e2a1e463a99f2b523748a360b.jpg" alt="" /></p>

<p>如你所见，我们在对不同类型文件进行合并时，要考虑不同的问题：</p>

<ol>
<li>对于支持格式丰富不同的文件时要考虑格式的兼容性；</li>
<li>对于图片、音乐、视频和Word合并时要考虑是否是受到word支持的通用格式；</li>
<li>对于像Excel格式于Word合并时能实现更复杂的功能，代码的复杂程度也会随之提高，一般需要先分析功能，再进行代码编写。</li>
</ol>

<h2 id="小结">小结</h2>

<p>通过上面对Word文件的批量处理，我为你总结了Word和各种类型合并增效的几个通用法则。</p>

<ul>
<li>首先，尽量选择Word兼容的格式,这些格式往往也是python-docx库能直接支持的类型。</li>
<li>第二，善于将手工操作转换为Python程序实现。如果无法直接转换为Python程序，可以尝试将手工操作继续细化拆分。</li>
<li>第三，反复在程序中出现的代码可以编写为函数功能，函数可以让你的程序更健壮，较短的代码数量也减少了出现Bug的机率。</li>
</ul>

<h2 id="思考题">思考题</h2>

<p>在最后我也想留一个问题给你思考，如果邀请函的格式从Word文件改为图片，你将会如何去解决呢?</p>

<p>如果你觉得这节课有用，能解决你的办公效率问题，欢迎你点击“请朋友读”，分享给你的朋友或同事。</p>

<p>编辑小提示：专栏的完整代码位置是<a href="https://github.com/wilsonyin123/python_productivity，可点击链接下载查看。或者通过网盘链接提取后下载，链接是:" target="_blank">https://github.com/wilsonyin123/python_productivity，可点击链接下载查看。或者通过网盘链接提取后下载，链接是:</a> <a href="https://pan.baidu.com/s/1UvEKDCGnU6yb0a7gHLSE4Q?pwd=5wf1%EF%BC%8C%E6%8F%90%E5%8F%96%E7%A0%81:" target="_blank">https://pan.baidu.com/s/1UvEKDCGnU6yb0a7gHLSE4Q?pwd=5wf1，提取码:</a> 5wf1。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#f09c9c9cc9c4c1c1c0c7b0979d91999cde939f9d" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f11d920e8dccd35',t:'MTczNDA0OTg1My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
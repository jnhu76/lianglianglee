<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=39&#32;Django：搭建监控平台>
        <link rel="icon" href="/static/favicon.png">
        <title>39 Django：搭建监控平台 </title>
        
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
                        <a class="menu-item" id="00 开篇词 从工程的角度深入理解Python.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%e5%b7%a5%e7%a8%8b%e7%9a%84%e8%a7%92%e5%ba%a6%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3Python.md">00 开篇词 从工程的角度深入理解Python.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 如何逐步突破，成为Python高手？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/01%20%e5%a6%82%e4%bd%95%e9%80%90%e6%ad%a5%e7%aa%81%e7%a0%b4%ef%bc%8c%e6%88%90%e4%b8%baPython%e9%ab%98%e6%89%8b%ef%bc%9f.md">01 如何逐步突破，成为Python高手？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 Jupyter Notebook为什么是现代Python的必学技术？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/02%20Jupyter%20Notebook%e4%b8%ba%e4%bb%80%e4%b9%88%e6%98%af%e7%8e%b0%e4%bb%a3Python%e7%9a%84%e5%bf%85%e5%ad%a6%e6%8a%80%e6%9c%af%ef%bc%9f.md">02 Jupyter Notebook为什么是现代Python的必学技术？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 列表和元组，到底用哪一个？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/03%20%e5%88%97%e8%a1%a8%e5%92%8c%e5%85%83%e7%bb%84%ef%bc%8c%e5%88%b0%e5%ba%95%e7%94%a8%e5%93%aa%e4%b8%80%e4%b8%aa%ef%bc%9f.md">03 列表和元组，到底用哪一个？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 字典、集合，你真的了解吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/04%20%e5%ad%97%e5%85%b8%e3%80%81%e9%9b%86%e5%90%88%ef%bc%8c%e4%bd%a0%e7%9c%9f%e7%9a%84%e4%ba%86%e8%a7%a3%e5%90%97%ef%bc%9f.md">04 字典、集合，你真的了解吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 深入浅出字符串.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/05%20%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%ad%97%e7%ac%a6%e4%b8%b2.md">05 深入浅出字符串.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Python “黑箱”：输入与输出.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/06%20Python%20%e2%80%9c%e9%bb%91%e7%ae%b1%e2%80%9d%ef%bc%9a%e8%be%93%e5%85%a5%e4%b8%8e%e8%be%93%e5%87%ba.md">06 Python “黑箱”：输入与输出.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 修炼基本功：条件与循环.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/07%20%e4%bf%ae%e7%82%bc%e5%9f%ba%e6%9c%ac%e5%8a%9f%ef%bc%9a%e6%9d%a1%e4%bb%b6%e4%b8%8e%e5%be%aa%e7%8e%af.md">07 修炼基本功：条件与循环.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 异常处理：如何提高程序的稳定性？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/08%20%e5%bc%82%e5%b8%b8%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e7%a8%8b%e5%ba%8f%e7%9a%84%e7%a8%b3%e5%ae%9a%e6%80%a7%ef%bc%9f.md">08 异常处理：如何提高程序的稳定性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 不可或缺的自定义函数.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/09%20%e4%b8%8d%e5%8f%af%e6%88%96%e7%bc%ba%e7%9a%84%e8%87%aa%e5%ae%9a%e4%b9%89%e5%87%bd%e6%95%b0.md">09 不可或缺的自定义函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 简约不简单的匿名函数.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/10%20%e7%ae%80%e7%ba%a6%e4%b8%8d%e7%ae%80%e5%8d%95%e7%9a%84%e5%8c%bf%e5%90%8d%e5%87%bd%e6%95%b0.md">10 简约不简单的匿名函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 面向对象（上）：从生活中的类比说起.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/11%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%bb%8e%e7%94%9f%e6%b4%bb%e4%b8%ad%e7%9a%84%e7%b1%bb%e6%af%94%e8%af%b4%e8%b5%b7.md">11 面向对象（上）：从生活中的类比说起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 面向对象（下）：如何实现一个搜索引擎？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/12%20%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%ef%bc%9f.md">12 面向对象（下）：如何实现一个搜索引擎？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 搭建积木：Python 模块化.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/13%20%e6%90%ad%e5%bb%ba%e7%a7%af%e6%9c%a8%ef%bc%9aPython%20%e6%a8%a1%e5%9d%97%e5%8c%96.md">13 搭建积木：Python 模块化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 答疑（一）：列表和元组的内部实现是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/14%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%88%97%e8%a1%a8%e5%92%8c%e5%85%83%e7%bb%84%e7%9a%84%e5%86%85%e9%83%a8%e5%ae%9e%e7%8e%b0%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">14 答疑（一）：列表和元组的内部实现是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Python对象的比较、拷贝.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/15%20Python%e5%af%b9%e8%b1%a1%e7%9a%84%e6%af%94%e8%be%83%e3%80%81%e6%8b%b7%e8%b4%9d.md">15 Python对象的比较、拷贝.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 值传递，引用传递or其他，Python里参数是如何传递的？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/16%20%e5%80%bc%e4%bc%a0%e9%80%92%ef%bc%8c%e5%bc%95%e7%94%a8%e4%bc%a0%e9%80%92or%e5%85%b6%e4%bb%96%ef%bc%8cPython%e9%87%8c%e5%8f%82%e6%95%b0%e6%98%af%e5%a6%82%e4%bd%95%e4%bc%a0%e9%80%92%e7%9a%84%ef%bc%9f.md">16 值传递，引用传递or其他，Python里参数是如何传递的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 强大的装饰器.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/17%20%e5%bc%ba%e5%a4%a7%e7%9a%84%e8%a3%85%e9%a5%b0%e5%99%a8.md">17 强大的装饰器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 metaclass，是潘多拉魔盒还是阿拉丁神灯？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/18%20metaclass%ef%bc%8c%e6%98%af%e6%bd%98%e5%a4%9a%e6%8b%89%e9%ad%94%e7%9b%92%e8%bf%98%e6%98%af%e9%98%bf%e6%8b%89%e4%b8%81%e7%a5%9e%e7%81%af%ef%bc%9f.md">18 metaclass，是潘多拉魔盒还是阿拉丁神灯？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 深入理解迭代器和生成器.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/19%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e8%bf%ad%e4%bb%a3%e5%99%a8%e5%92%8c%e7%94%9f%e6%88%90%e5%99%a8.md">19 深入理解迭代器和生成器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 揭秘 Python 协程.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/20%20%e6%8f%ad%e7%a7%98%20Python%20%e5%8d%8f%e7%a8%8b.md">20 揭秘 Python 协程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 Python并发编程之Futures.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/21%20Python%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e4%b9%8bFutures.md">21 Python并发编程之Futures.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 并发编程之Asyncio.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/22%20%e5%b9%b6%e5%8f%91%e7%bc%96%e7%a8%8b%e4%b9%8bAsyncio.md">22 并发编程之Asyncio.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 你真的懂Python GIL（全局解释器锁）吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/23%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e6%87%82Python%20GIL%ef%bc%88%e5%85%a8%e5%b1%80%e8%a7%a3%e9%87%8a%e5%99%a8%e9%94%81%ef%bc%89%e5%90%97%ef%bc%9f.md">23 你真的懂Python GIL（全局解释器锁）吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 带你解析 Python 垃圾回收机制.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/24%20%e5%b8%a6%e4%bd%a0%e8%a7%a3%e6%9e%90%20Python%20%e5%9e%83%e5%9c%be%e5%9b%9e%e6%94%b6%e6%9c%ba%e5%88%b6.md">24 带你解析 Python 垃圾回收机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 答疑（二）：GIL与多线程是什么关系呢？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/25%20%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9aGIL%e4%b8%8e%e5%a4%9a%e7%ba%bf%e7%a8%8b%e6%98%af%e4%bb%80%e4%b9%88%e5%85%b3%e7%b3%bb%e5%91%a2%ef%bc%9f.md">25 答疑（二）：GIL与多线程是什么关系呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 活都来不及干了，还有空注意代码风格？！.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/26%20%e6%b4%bb%e9%83%bd%e6%9d%a5%e4%b8%8d%e5%8f%8a%e5%b9%b2%e4%ba%86%ef%bc%8c%e8%bf%98%e6%9c%89%e7%a9%ba%e6%b3%a8%e6%84%8f%e4%bb%a3%e7%a0%81%e9%a3%8e%e6%a0%bc%ef%bc%9f%ef%bc%81.md">26 活都来不及干了，还有空注意代码风格？！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 学会合理分解代码，提高代码可读性.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/27%20%e5%ad%a6%e4%bc%9a%e5%90%88%e7%90%86%e5%88%86%e8%a7%a3%e4%bb%a3%e7%a0%81%ef%bc%8c%e6%8f%90%e9%ab%98%e4%bb%a3%e7%a0%81%e5%8f%af%e8%af%bb%e6%80%a7.md">27 学会合理分解代码，提高代码可读性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 如何合理利用assert？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/28%20%e5%a6%82%e4%bd%95%e5%90%88%e7%90%86%e5%88%a9%e7%94%a8assert%ef%bc%9f.md">28 如何合理利用assert？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 巧用上下文管理器和With语句精简代码.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/29%20%e5%b7%a7%e7%94%a8%e4%b8%8a%e4%b8%8b%e6%96%87%e7%ae%a1%e7%90%86%e5%99%a8%e5%92%8cWith%e8%af%ad%e5%8f%a5%e7%b2%be%e7%ae%80%e4%bb%a3%e7%a0%81.md">29 巧用上下文管理器和With语句精简代码.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 真的有必要写单元测试吗？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/30%20%e7%9c%9f%e7%9a%84%e6%9c%89%e5%bf%85%e8%a6%81%e5%86%99%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%e5%90%97%ef%bc%9f.md">30 真的有必要写单元测试吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 pdb &amp; cProfile：调试和性能分析的法宝.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/31%20pdb%20&amp;%20cProfile%ef%bc%9a%e8%b0%83%e8%af%95%e5%92%8c%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%e7%9a%84%e6%b3%95%e5%ae%9d.md">31 pdb &amp; cProfile：调试和性能分析的法宝.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 答疑（三）：如何选择合适的异常处理方式？.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/32%20%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e5%90%88%e9%80%82%e7%9a%84%e5%bc%82%e5%b8%b8%e5%a4%84%e7%90%86%e6%96%b9%e5%bc%8f%ef%bc%9f.md">32 答疑（三）：如何选择合适的异常处理方式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 带你初探量化世界.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/33%20%e5%b8%a6%e4%bd%a0%e5%88%9d%e6%8e%a2%e9%87%8f%e5%8c%96%e4%b8%96%e7%95%8c.md">33 带你初探量化世界.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 RESTful &amp; Socket：搭建交易执行层核心.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/34%20RESTful%20&amp;%20Socket%ef%bc%9a%e6%90%ad%e5%bb%ba%e4%ba%a4%e6%98%93%e6%89%a7%e8%a1%8c%e5%b1%82%e6%a0%b8%e5%bf%83.md">34 RESTful &amp; Socket：搭建交易执行层核心.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 RESTful &amp; Socket：行情数据对接和抓取.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/35%20RESTful%20&amp;%20Socket%ef%bc%9a%e8%a1%8c%e6%83%85%e6%95%b0%e6%8d%ae%e5%af%b9%e6%8e%a5%e5%92%8c%e6%8a%93%e5%8f%96.md">35 RESTful &amp; Socket：行情数据对接和抓取.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 Pandas &amp; Numpy：策略与回测系统.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/36%20Pandas%20&amp;%20Numpy%ef%bc%9a%e7%ad%96%e7%95%a5%e4%b8%8e%e5%9b%9e%e6%b5%8b%e7%b3%bb%e7%bb%9f.md">36 Pandas &amp; Numpy：策略与回测系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 Kafka &amp; ZMQ：自动化交易流水线.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/37%20Kafka%20&amp;%20ZMQ%ef%bc%9a%e8%87%aa%e5%8a%a8%e5%8c%96%e4%ba%a4%e6%98%93%e6%b5%81%e6%b0%b4%e7%ba%bf.md">37 Kafka &amp; ZMQ：自动化交易流水线.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 MySQL：日志和数据存储系统.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/38%20MySQL%ef%bc%9a%e6%97%a5%e5%bf%97%e5%92%8c%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e7%b3%bb%e7%bb%9f.md">38 MySQL：日志和数据存储系统.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 Django：搭建监控平台.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/39%20Django%ef%bc%9a%e6%90%ad%e5%bb%ba%e7%9b%91%e6%8e%a7%e5%b9%b3%e5%8f%b0.md">39 Django：搭建监控平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 总结：Python中的数据结构与算法全景.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/40%20%e6%80%bb%e7%bb%93%ef%bc%9aPython%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e4%b8%8e%e7%ae%97%e6%b3%95%e5%85%a8%e6%99%af.md">40 总结：Python中的数据结构与算法全景.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 硅谷一线互联网公司的工作体验.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/41%20%e7%a1%85%e8%b0%b7%e4%b8%80%e7%ba%bf%e4%ba%92%e8%81%94%e7%bd%91%e5%85%ac%e5%8f%b8%e7%9a%84%e5%b7%a5%e4%bd%9c%e4%bd%93%e9%aa%8c.md">41 硅谷一线互联网公司的工作体验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 细数技术研发的注意事项.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/42%20%e7%bb%86%e6%95%b0%e6%8a%80%e6%9c%af%e7%a0%94%e5%8f%91%e7%9a%84%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9.md">42 细数技术研发的注意事项.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 Q&amp;A：聊一聊职业发展和选择.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/43%20Q&amp;A%ef%bc%9a%e8%81%8a%e4%b8%80%e8%81%8a%e8%81%8c%e4%b8%9a%e5%8f%91%e5%b1%95%e5%92%8c%e9%80%89%e6%8b%a9.md">43 Q&amp;A：聊一聊职业发展和选择.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 带你上手SWIG：一份清晰好用的SWIG编程实践指南.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e5%8a%a0%e9%a4%90%20%e5%b8%a6%e4%bd%a0%e4%b8%8a%e6%89%8bSWIG%ef%bc%9a%e4%b8%80%e4%bb%bd%e6%b8%85%e6%99%b0%e5%a5%bd%e7%94%a8%e7%9a%84SWIG%e7%bc%96%e7%a8%8b%e5%ae%9e%e8%b7%b5%e6%8c%87%e5%8d%97.md">加餐 带你上手SWIG：一份清晰好用的SWIG编程实践指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 技术之外的几点成长建议.md" href="/%e4%b8%93%e6%a0%8f/Python%e6%a0%b8%e5%bf%83%e6%8a%80%e6%9c%af%e4%b8%8e%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%8a%80%e6%9c%af%e4%b9%8b%e5%a4%96%e7%9a%84%e5%87%a0%e7%82%b9%e6%88%90%e9%95%bf%e5%bb%ba%e8%ae%ae.md">结束语 技术之外的几点成长建议.md</a>
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
                            <h1 id="title" data-id="39 Django：搭建监控平台" class="title">39 Django：搭建监控平台</h1>
                            <div><p>你好，我是景霄。</p>

<p>通过前几节课的学习，相信你对量化交易系统已经有了一个最基本的认知，也能通过自己的代码，搭建一个简单的量化交易系统来进行盈利。</p>

<p>前面几节课，我们的重点在后台代码、中间件、分布式系统和设计模式上。这节课，我们重点来看前端交互。</p>

<p>监控和运维，是互联网工业链上非常重要的一环。监控的目的就是防患于未然。通过监控，我们能够及时了解到企业网络的运行状态。一旦出现安全隐患，你就可以及时预警，或者是以其他方式通知运维人员，让运维监控人员有时间处理和解决隐患，避免影响业务系统的正常使用，将一切问题的根源扼杀在摇篮当中。</p>

<p>在硅谷互联网大公司中，监控和运维被称为 SRE，是公司正常运行中非常重要的一环。作为 billion 级别的 Facebook，内部自然也有着大大小小、各种各样的监控系统和运维工具，有的对标业务数据，有的对标服务器的健康状态，有的则是面向数据库和微服务的控制信息。</p>

<p>不过，万变不离其宗，运维工作最重要的就是维护系统的稳定性。除了熟悉运用各种提高运维效率的工具来辅助工作外，云资源费用管理、安全管理、监控等，都需要耗费不少精力和时间。运维监控不是一朝一夕得来的，而是随着业务发展的过程中同步和发展的。</p>

<p>作为量化实践内容的最后一节，今天我们就使用 Django 这个 Web 框架，来搭建一个简单的量化监控平台。</p>

<h2 id="django-简介和安装">Django 简介和安装</h2>

<p>Django 是用 Python 开发的一个免费开源的 Web 框架，可以用来快速搭建优雅的高性能网站。它采用的是“MVC”的框架模式，即模型 M、视图 V 和控制器 C。</p>

<p>Django 最大的特色，在于将网页和数据库中复杂的关系，转化为 Python 中对应的简单关系。它的设计目的，是使常见的Web开发任务变得快速而简单。Django是开源的，不是商业项目或者科研项目，并且集中力量解决Web开发中遇到的一系列问题。所以，Django 每天都会在现有的基础上进步，以适应不断更迭的开发需求。这样既节省了开发时间，也提高了后期维护的效率。</p>

<p>说了这么多，接下来，我们通过上手使用进一步来了解。先来看一下，如何安装和使用 Django。你可以先按照下面代码块的内容来操作，安装Django ：</p>

<pre><code>pip3 install Django
django-admin --version

########## 输出 ##########

2.2.3
</code></pre>

<p>接着，我们来创建一个新的 Django 项目：</p>

<pre><code>django-admin startproject TradingMonitor
cd TradingMonitor/
python3 manage.py migrate

########## 输出 ##########
</code></pre>

<pre><code>  Applying contenttypes.0001_initial... OK
  Applying auth.0001_initial... OK
  Applying admin.0001_initial... OK
  Applying admin.0002_logentry_remove_auto_add... OK
  Applying admin.0003_logentry_add_action_flag_choices... OK
  Applying contenttypes.0002_remove_content_type_name... OK
  Applying auth.0002_alter_permission_name_max_length... OK
  Applying auth.0003_alter_user_email_max_length... OK
  Applying auth.0004_alter_user_username_opts... OK
  Applying auth.0005_alter_user_last_login_null... OK
  Applying auth.0006_require_contenttypes_0002... OK
  Applying auth.0007_alter_validators_add_error_messages... OK
  Applying auth.0008_alter_user_username_max_length... OK
  Applying auth.0009_alter_user_last_name_max_length... OK
  Applying auth.0010_alter_group_name_max_length... OK
  Applying auth.0011_update_proxy_permissions... OK
  Applying sessions.0001_initial... OK
</code></pre>

<p>这时，你能看到文件系统大概是下面这样的：</p>

<pre><code>TradingMonitor/
├── TradingMonitor
│   ├── __init__.py
│   ├── settings.py
│   ├── urls.py
│   └── wsgi.py
├── db.sqlite3
└── manage.py
</code></pre>

<p>我简单解释一下它的意思：</p>

<ul>
<li>TradingMonitor/TradingMonitor，表示项目最初的 Python 包；</li>
<li>TradingMonitor/init.py，表示一个空文件，声明所在目录的包为一个 Python 包；</li>
<li>TradingMonitor/settings.py，管理项目的配置信息；</li>
<li>TradingMonitor/urls.py，声明请求 URL 的映射关系；</li>
<li>TradingMonitor/wsgi.py，表示Python 程序和 Web 服务器的通信协议；</li>
<li>manage.py，表示一个命令行工具，用来和 Django 项目进行交互；</li>
<li>Db.sqlite3，表示默认的数据库，可以在设置中替换成其他数据库。</li>
</ul>

<p>另外，你可能注意到了上述命令中的<code>python3 manage.py migrate</code>，这个命令表示创建或更新数据库模式。每当 model 源代码被改变后，如果我们要将其应用到数据库上，就需要执行一次这个命令。</p>

<p>接下来，我们为这个系统添加管理员账户：</p>

<pre><code>python3 manage.py createsuperuser

########## 输出 ##########

Username (leave blank to use 'ubuntu'): admin
Email address:  
Password: 
Password (again): 
Superuser created successfully.
</code></pre>

<p>然后，我们来启动 Django 的 debugging 模式：</p>

<pre><code>python3 manage.py runserver
</code></pre>

<p>最后，打开浏览器输入：<code>http://127.0.0.1:8000</code>。如果你能看到下面这个画面，就说明 Django 已经部署成功了。</p>

<p><img src="assets/d9cf613d333046ab81652f2ab24318d5.jpg" alt="" /></p>

<p>Django 的安装是不是非常简单呢？这其实也是 Python 一贯的理念，简洁，并简化入门的门槛。</p>

<p>OK，现在我们再定位到 <code>http://127.0.0.1:8000/admin</code>，你会看到 Django 的后台管理网页，这里我就不过多介绍了。</p>

<p><img src="assets/96fafb04739a4533b55d71d3e0c9a90a.jpg" alt="" />-
<img src="assets/4effeb0ad6164effae7801c03e094888.jpg" alt="" /></p>

<p>到此，Django 就已经成功安装，并且正常启动啦。</p>

<h2 id="mvc-架构">MVC 架构</h2>

<p>刚刚我说过，MVC 架构是 Django 设计模式的精髓。接下来，我们就来具体看一下这个架构，并通过 Django 动手搭建一个服务端。</p>

<h3 id="设计模型-model">设计模型 Model</h3>

<p>在之前的日志和存储系统这节课中，我介绍过 peewee 这个库，它能避开通过繁琐的 SQL 语句来操作 MySQL，直接使用 Python 的 class 来进行转换。事实上，这也是 Django 采取的方式。</p>

<p>Django 无需数据库就可以使用，它通过对象关系映射器（object-relational mapping），仅使用Python代码就可以描述数据结构。</p>

<p>我们先来看下面这段 Model 代码：</p>

<pre><code>#  TradingMonitor/models.py

from django.db import models


class Position(models.Model):
    asset = models.CharField(max_length=10)
    timestamp = models.DateTimeField()
    amount = models.DecimalField(max_digits=10, decimal_places=3)
</code></pre>

<p>models.py 文件主要用一个 Python 类来描述数据表，称为模型 。运用这个类，你可以通过简单的 Python 代码来创建、检索、更新、删除数据库中的记录，而不用写一条又一条的SQL语句，这也是我们之前所说的避免通过 SQL 操作数据库。</p>

<p>在这里，我们创建了一个 Position 模型，用来表示我们的交易仓位信息。其中，</p>

<ul>
<li>asset 表示当前持有资产的代码，例如 btc；</li>
<li>timestamp 表示时间戳；</li>
<li>amount 则表示时间戳时刻的持仓信息。</li>
</ul>

<h3 id="设计视图-views">设计视图 Views</h3>

<p>在模型被定义之后，我们便可以在视图中引用模型了。通常，视图会根据参数检索数据，加载一个模板，并使用检索到的数据呈现模板。</p>

<p>设计视图，则是我们用来实现业务逻辑的地方。我们来看 render_positions 这个代码，它接受 request 和 asset 两个参数，我们先不用管 request。这里的 asset 表示指定一个资产名称，例如 btc，然后这个函数返回一个渲染页面。</p>

<pre><code>#  TradingMonitor/views.py

from django.shortcuts import render
from .models import Position

def render_positions(request, asset):
    positions = Position.objects.filter(asset = asset)
    context = {'asset': asset, 'positions': positions}
    return render(request, 'positions.html', context)
</code></pre>

<p>不过，这个函数具体是怎么工作的呢？我们一行行来看。</p>

<p><code>positions = Position.objects.filter(asset = asset)</code>，这行代码向数据库中执行一个查询操作，其中， filter 表示筛选，意思是从数据库中选出所有我们需要的 asset 的信息。不过，这里我只是为你举例做示范；真正做监控的时候，我们一般会更有针对性地从数据库中筛选读取信息，而不是一口气读取出所有的信息。</p>

<p><code>context = {'asset': asset, 'positions': positions}</code>，这行代码没什么好说的，封装一个字典。至于这个字典的用处，下面的内容中可以体现。</p>

<p><code>return render(request, 'positions.html', context)</code>，最后这行代码返回一个页面。这里我们采用的模板设计，这也是 Django 非常推荐的开发方式，也就是让模板和数据分离，这样，数据只需要向其中填充即可。</p>

<p>最后的模板文件是 <code>position.html</code>，你应该注意到了， context 作为变量传给了模板，下面我们就来看一下设计模板的内容。</p>

<h3 id="设计模板templates">设计模板Templates</h3>

<p>模板文件，其实就是 HTML 文件和部分代码的综合。你可以想象成，这个HTML 在最终送给用户之前，需要被我们预先处理一下，而预先处理的方式就是找到对应的地方进行替换。</p>

<p>我们来看下面这段示例代码：</p>

<pre><code>#  TradingMonitor/templates/positions.html

&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en-US&quot;&gt;
&lt;head&gt;
&lt;title&gt;Positions for {{asset}}&lt;/title&gt;
&lt;/head&gt;

&lt;body&gt;
&lt;h1&gt;Positions for {{asset}}&lt;/h1&gt;

&lt;table&gt;
&lt;tr&gt;
    &lt;th&gt;Time&lt;/th&gt;
    &lt;th&gt;Amount&lt;/th&gt;
&lt;/tr&gt;
{% for position in positions %}
&lt;tr&gt;
    &lt;th&gt;{{position.timestamp}}&lt;/th&gt;
    &lt;th&gt;{{position.amount}}&lt;/th&gt;
&lt;/tr&gt;
{% endfor %}
&lt;/table&gt;
&lt;/body&gt;
</code></pre>

<p>我重点说一下几个地方。首先是<code>&lt;title&gt;Positions for {{asset}}&lt;/title&gt;</code>，这里双大括号括住 asset 这个变量，这个变量对应的正是前面 context 字典中的 asset key。Django 的渲染引擎会将 asset ，替换成 context 中 asset 对应的内容，此处是替换成了 btc。</p>

<p>再来看<code>{% for position in positions %}</code>，这是个很关键的地方。我们需要处理一个列表的情况，用 for 对 positions 进行迭代就行了。这里的 positions ，同样对应的是 context 中的 positions。</p>

<p>末尾的<code>{% endfor %}</code>，自然就表示结束了。这样，我们就将数据封装到了一个列表之中。</p>

<h3 id="设计链接-urls">设计链接 Urls</h3>

<p>最后，我们需要为我们的操作提供 URL 接口，具体操作我放在了下面的代码中，内容比较简单，我就不详细展开讲解了。</p>

<pre><code>#  TradingMonitor/urls.py

from django.contrib import admin
from django.urls import path
from . import views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('positions/&lt;str:asset&gt;', views.render_positions),
]
</code></pre>

<p>到这里，我们就可以通过 <code>http://127.0.0.1:8000/positions/btc</code> 来访问啦！</p>

<h3 id="测试">测试</h3>

<p>当然，除了主要流程外，我还需要强调几个很简单但非常关键的细节，不然，我们这些改变就不能被真正地应用。</p>

<p>第一步，在 <code>TradingMonitor/TradingMonitor</code> 下，新建一个文件夹 migrations；并在这个文件夹中，新建一个空文件 <code>__init__.py</code>。</p>

<pre><code>mkdir TradingMonitor/migrations
touch TradingMonitor/migrations/__init__.py
</code></pre>

<p>此时，你的目录结构应该长成下面这样：</p>

<pre><code>TradingMonitor/
├── TradingMonitor
│   ├── migrations
│       └── __init__.py
│   ├── templates
│       └── positions.html
│   ├── __init__.py
│   ├── settings.py
│   ├── urls.py
│   ├── models.py
│   ├── views.py
│   └── wsgi.py
├── db.sqlite3
└── manage.py
</code></pre>

<p>第二步，修改 <code>TradingMonitor/settings.py</code>：</p>

<pre><code>INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'TradingMonitor',  # 这里把我们的 app 加上
]
</code></pre>

<pre><code>TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, 'TradingMonitor/templates')],  # 这里把 templates 的目录加上
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]
</code></pre>

<p>第三步，运行 <code>python manage.py makemigrations</code>：</p>

<pre><code>python manage.py makemigrations

########## 输出 ##########

Migrations for 'TradingMonitor':
  TradingMonitor/migrations/0001_initial.py
    - Create model Position
</code></pre>

<p>第四步，运行 <code>python manage.py migrate</code>：</p>

<pre><code>python manage.py migrate


########## 输出 ##########


Operations to perform:
  Apply all migrations: TradingMonitor, admin, auth, contenttypes, sessions
Running migrations:
  Applying TradingMonitor.0001_initial... OK
</code></pre>

<p>这几步的具体操作，我都用代码和注释表示了出来，你完全可以同步进行操作。操作完成后，现在，我们的数据结构就已经被成功同步到数据库中了。</p>

<p>最后，输入 <code>python manage.py runserver</code>，然后打开浏览器输入<code>http://127.0.0.1:8000/positions/btc</code>，你就能看到效果啦。-
<img src="assets/1d2fbdf452994e5c8c67f72e8220325e.jpg" alt="" /></p>

<p>现在，我们再回过头来看一下 MVC 模式，通过我画的这张图，你可以看到，M、V、C这三者，以一种插件似的、松耦合的方式连接在一起：</p>

<p><img src="assets/9abc2b1b5f8d427a8b0a1ba8b06591de.jpg" alt="" /></p>

<p>当然，我带你写的只是一个简单的 Django 应用程序，对于真正的量化平台监控系统而言，这还只是一个简单的开始。</p>

<p>除此之外，对于监控系统来说，其实还有着非常多的开源插件可以使用。有一些界面非常酷炫，有一些可以做到很高的稳定性和易用性，它们很多都可以结合 Django 做出很好的效果来。比较典型的有：</p>

<ul>
<li>Graphite 是一款存储时间序列数据，并通过 Django Web 应用程序在图形中显示的插件；</li>
<li>Vimeo 则是一个基于 Graphite 的仪表板，具有附加功能和平滑的设计；</li>
<li>Scout 监控 Django和Flask应用程序的性能，提供自动检测视图、SQL查询、模板等。</li>
</ul>

<h2 id="总结">总结</h2>

<p>这一节课的内容更靠近上游应用层，我们以 Django 这个 Python 后端为例，讲解了搭建一个服务端的过程。你应该发现了，使用 RESTful Framework 搭建服务器，是一个如此简单的过程，你可以去开一个自己的交易所了（笑）。相比起具体的技术，今天我所讲的 MVC 框架和 Django 的思想，更值得你去深入学习和领会。</p>

<h2 id="思考题">思考题</h2>

<p>今天我想给你留一个难度比较高的作业。RESTful API 在 Django 中是如何实现安全认证的？你能通过搜索和自学掌握这个知识点吗？希望可以在留言区看到你的认真学习记录和总结，我会一一给出建议。也欢迎你把这篇文章分享给你的朋友、同事，一起交流、一起进步。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#2f434343161b1e1e1f186f48424e4643014c4042" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f11d74b5e92cd35',t:'MTczNDA0OTc3OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;软件设计实践：如何使用UML完成一个设计文档？>
        <link rel="icon" href="/static/favicon.png">
        <title>09 软件设计实践：如何使用UML完成一个设计文档？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 掌握软件开发技术的第一性原理.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e6%8e%8c%e6%8f%a1%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%e6%8a%80%e6%9c%af%e7%9a%84%e7%ac%ac%e4%b8%80%e6%80%a7%e5%8e%9f%e7%90%86.md">00 开篇词 掌握软件开发技术的第一性原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 程序运行原理：程序是如何运行又是如何崩溃的？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/01%20%e7%a8%8b%e5%ba%8f%e8%bf%90%e8%a1%8c%e5%8e%9f%e7%90%86%ef%bc%9a%e7%a8%8b%e5%ba%8f%e6%98%af%e5%a6%82%e4%bd%95%e8%bf%90%e8%a1%8c%e5%8f%88%e6%98%af%e5%a6%82%e4%bd%95%e5%b4%a9%e6%ba%83%e7%9a%84%ef%bc%9f.md">01 程序运行原理：程序是如何运行又是如何崩溃的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 数据结构原理：Hash表的时间复杂度为什么是O(1)？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/02%20%e6%95%b0%e6%8d%ae%e7%bb%93%e6%9e%84%e5%8e%9f%e7%90%86%ef%bc%9aHash%e8%a1%a8%e7%9a%84%e6%97%b6%e9%97%b4%e5%a4%8d%e6%9d%82%e5%ba%a6%e4%b8%ba%e4%bb%80%e4%b9%88%e6%98%afO%281%29%ef%bc%9f.md">02 数据结构原理：Hash表的时间复杂度为什么是O(1)？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 Java虚拟机原理：JVM为什么被称为机器（machine）？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/03%20Java%e8%99%9a%e6%8b%9f%e6%9c%ba%e5%8e%9f%e7%90%86%ef%bc%9aJVM%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a2%ab%e7%a7%b0%e4%b8%ba%e6%9c%ba%e5%99%a8%ef%bc%88machine%ef%bc%89%ef%bc%9f.md">03 Java虚拟机原理：JVM为什么被称为机器（machine）？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 网络编程原理：一个字符的互联网之旅.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/04%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%e5%8e%9f%e7%90%86%ef%bc%9a%e4%b8%80%e4%b8%aa%e5%ad%97%e7%ac%a6%e7%9a%84%e4%ba%92%e8%81%94%e7%bd%91%e4%b9%8b%e6%97%85.md">04 网络编程原理：一个字符的互联网之旅.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 文件系统原理：如何用1分钟遍历一个100TB的文件？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/05%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%8e%9f%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a81%e5%88%86%e9%92%9f%e9%81%8d%e5%8e%86%e4%b8%80%e4%b8%aa100TB%e7%9a%84%e6%96%87%e4%bb%b6%ef%bc%9f.md">05 文件系统原理：如何用1分钟遍历一个100TB的文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 数据库原理：为什么PrepareStatement性能更好更安全？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/06%20%e6%95%b0%e6%8d%ae%e5%ba%93%e5%8e%9f%e7%90%86%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88PrepareStatement%e6%80%a7%e8%83%bd%e6%9b%b4%e5%a5%bd%e6%9b%b4%e5%ae%89%e5%85%a8%ef%bc%9f.md">06 数据库原理：为什么PrepareStatement性能更好更安全？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 答疑 Java Web程序的运行时环境到底是怎样的？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/07%20%e7%ad%94%e7%96%91%20Java%20Web%e7%a8%8b%e5%ba%8f%e7%9a%84%e8%bf%90%e8%a1%8c%e6%97%b6%e7%8e%af%e5%a2%83%e5%88%b0%e5%ba%95%e6%98%af%e6%80%8e%e6%a0%b7%e7%9a%84%ef%bc%9f.md">07 答疑 Java Web程序的运行时环境到底是怎样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 编程语言原理：面向对象编程是编程的终极形态吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/07%20%e7%bc%96%e7%a8%8b%e8%af%ad%e8%a8%80%e5%8e%9f%e7%90%86%ef%bc%9a%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e7%bc%96%e7%a8%8b%e6%98%af%e7%bc%96%e7%a8%8b%e7%9a%84%e7%bb%88%e6%9e%81%e5%bd%a2%e6%80%81%e5%90%97%ef%bc%9f.md">07 编程语言原理：面向对象编程是编程的终极形态吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 软件设计的方法论：软件为什么要建模？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/08%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e6%96%b9%e6%b3%95%e8%ae%ba%ef%bc%9a%e8%bd%af%e4%bb%b6%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%bb%ba%e6%a8%a1%ef%bc%9f.md">08 软件设计的方法论：软件为什么要建模？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 软件设计实践：如何使用UML完成一个设计文档？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/09%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e5%ae%9e%e8%b7%b5%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8UML%e5%ae%8c%e6%88%90%e4%b8%80%e4%b8%aa%e8%ae%be%e8%ae%a1%e6%96%87%e6%a1%a3%ef%bc%9f.md">09 软件设计实践：如何使用UML完成一个设计文档？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 软件设计的目的：糟糕的程序员比优秀的程序员差在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/10%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e7%9b%ae%e7%9a%84%ef%bc%9a%e7%b3%9f%e7%b3%95%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%e6%af%94%e4%bc%98%e7%a7%80%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%ae%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">10 软件设计的目的：糟糕的程序员比优秀的程序员差在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 软件设计的开闭原则：如何不修改代码却能实现需求变更？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/11%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e5%bc%80%e9%97%ad%e5%8e%9f%e5%88%99%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%8d%e4%bf%ae%e6%94%b9%e4%bb%a3%e7%a0%81%e5%8d%b4%e8%83%bd%e5%ae%9e%e7%8e%b0%e9%9c%80%e6%b1%82%e5%8f%98%e6%9b%b4%ef%bc%9f.md">11 软件设计的开闭原则：如何不修改代码却能实现需求变更？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 软件设计的依赖倒置原则：如何不依赖代码却可以复用它的功能？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/12%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e4%be%9d%e8%b5%96%e5%80%92%e7%bd%ae%e5%8e%9f%e5%88%99%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%8d%e4%be%9d%e8%b5%96%e4%bb%a3%e7%a0%81%e5%8d%b4%e5%8f%af%e4%bb%a5%e5%a4%8d%e7%94%a8%e5%ae%83%e7%9a%84%e5%8a%9f%e8%83%bd%ef%bc%9f.md">12 软件设计的依赖倒置原则：如何不依赖代码却可以复用它的功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 软件设计的里氏替换原则：正方形可以继承长方形吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/13%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e9%87%8c%e6%b0%8f%e6%9b%bf%e6%8d%a2%e5%8e%9f%e5%88%99%ef%bc%9a%e6%ad%a3%e6%96%b9%e5%bd%a2%e5%8f%af%e4%bb%a5%e7%bb%a7%e6%89%bf%e9%95%bf%e6%96%b9%e5%bd%a2%e5%90%97%ef%bc%9f.md">13 软件设计的里氏替换原则：正方形可以继承长方形吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 软件设计的单一职责原则：为什么说一个类文件打开最好不要超过一屏？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/14%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e5%8d%95%e4%b8%80%e8%81%8c%e8%b4%a3%e5%8e%9f%e5%88%99%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4%e4%b8%80%e4%b8%aa%e7%b1%bb%e6%96%87%e4%bb%b6%e6%89%93%e5%bc%80%e6%9c%80%e5%a5%bd%e4%b8%8d%e8%a6%81%e8%b6%85%e8%bf%87%e4%b8%80%e5%b1%8f%ef%bc%9f.md">14 软件设计的单一职责原则：为什么说一个类文件打开最好不要超过一屏？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 软件设计的接口隔离原则：如何对类的调用者隐藏类的公有方法？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/15%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e7%9a%84%e6%8e%a5%e5%8f%a3%e9%9a%94%e7%a6%bb%e5%8e%9f%e5%88%99%ef%bc%9a%e5%a6%82%e4%bd%95%e5%af%b9%e7%b1%bb%e7%9a%84%e8%b0%83%e7%94%a8%e8%80%85%e9%9a%90%e8%97%8f%e7%b1%bb%e7%9a%84%e5%85%ac%e6%9c%89%e6%96%b9%e6%b3%95%ef%bc%9f.md">15 软件设计的接口隔离原则：如何对类的调用者隐藏类的公有方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 设计模式基础：不会灵活应用设计模式，你就没有掌握面向对象编程.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/16%20%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%e5%9f%ba%e7%a1%80%ef%bc%9a%e4%b8%8d%e4%bc%9a%e7%81%b5%e6%b4%bb%e5%ba%94%e7%94%a8%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%8c%e4%bd%a0%e5%b0%b1%e6%b2%a1%e6%9c%89%e6%8e%8c%e6%8f%a1%e9%9d%a2%e5%90%91%e5%af%b9%e8%b1%a1%e7%bc%96%e7%a8%8b.md">16 设计模式基础：不会灵活应用设计模式，你就没有掌握面向对象编程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 设计模式应用：编程框架中的设计模式.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/17%20%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%e5%ba%94%e7%94%a8%ef%bc%9a%e7%bc%96%e7%a8%8b%e6%a1%86%e6%9e%b6%e4%b8%ad%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f.md">17 设计模式应用：编程框架中的设计模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 反应式编程框架设计：如何使程序调用不阻塞等待，立即响应？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/18%20%e5%8f%8d%e5%ba%94%e5%bc%8f%e7%bc%96%e7%a8%8b%e6%a1%86%e6%9e%b6%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%a8%8b%e5%ba%8f%e8%b0%83%e7%94%a8%e4%b8%8d%e9%98%bb%e5%a1%9e%e7%ad%89%e5%be%85%ef%bc%8c%e7%ab%8b%e5%8d%b3%e5%93%8d%e5%ba%94%ef%bc%9f.md">18 反应式编程框架设计：如何使程序调用不阻塞等待，立即响应？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 组件设计原则：组件的边界在哪里？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/19%20%e7%bb%84%e4%bb%b6%e8%ae%be%e8%ae%a1%e5%8e%9f%e5%88%99%ef%bc%9a%e7%bb%84%e4%bb%b6%e7%9a%84%e8%be%b9%e7%95%8c%e5%9c%a8%e5%93%aa%e9%87%8c%ef%bc%9f.md">19 组件设计原则：组件的边界在哪里？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 答疑 对于设计模式而言，场景到底有多重要？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/20%20%e7%ad%94%e7%96%91%20%e5%af%b9%e4%ba%8e%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%e8%80%8c%e8%a8%80%ef%bc%8c%e5%9c%ba%e6%99%af%e5%88%b0%e5%ba%95%e6%9c%89%e5%a4%9a%e9%87%8d%e8%a6%81%ef%bc%9f.md">20 答疑 对于设计模式而言，场景到底有多重要？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 领域驱动设计：35岁的程序员应该写什么样的代码？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/20%20%e9%a2%86%e5%9f%9f%e9%a9%b1%e5%8a%a8%e8%ae%be%e8%ae%a1%ef%bc%9a35%e5%b2%81%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%e5%ba%94%e8%af%a5%e5%86%99%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e4%bb%a3%e7%a0%81%ef%bc%9f.md">20 领域驱动设计：35岁的程序员应该写什么样的代码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 分布式架构：如何应对高并发的用户请求.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/21%20%e5%88%86%e5%b8%83%e5%bc%8f%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e9%ab%98%e5%b9%b6%e5%8f%91%e7%9a%84%e7%94%a8%e6%88%b7%e8%af%b7%e6%b1%82.md">21 分布式架构：如何应对高并发的用户请求.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 缓存架构：如何减少不必要的计算？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/22%20%e7%bc%93%e5%ad%98%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e5%87%8f%e5%b0%91%e4%b8%8d%e5%bf%85%e8%a6%81%e7%9a%84%e8%ae%a1%e7%ae%97%ef%bc%9f.md">22 缓存架构：如何减少不必要的计算？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 异步架构：如何避免互相依赖的系统间耦合？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/23%20%e5%bc%82%e6%ad%a5%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e4%ba%92%e7%9b%b8%e4%be%9d%e8%b5%96%e7%9a%84%e7%b3%bb%e7%bb%9f%e9%97%b4%e8%80%a6%e5%90%88%ef%bc%9f.md">23 异步架构：如何避免互相依赖的系统间耦合？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 负载均衡架构：如何用10行代码实现一个负载均衡服务？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/24%20%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%a810%e8%a1%8c%e4%bb%a3%e7%a0%81%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">24 负载均衡架构：如何用10行代码实现一个负载均衡服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 数据存储架构：如何改善系统的数据存储能力？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/25%20%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e6%94%b9%e5%96%84%e7%b3%bb%e7%bb%9f%e7%9a%84%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e8%83%bd%e5%8a%9b%ef%bc%9f.md">25 数据存储架构：如何改善系统的数据存储能力？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 搜索引擎架构：如何瞬间完成海量数据检索？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/26%20%e6%90%9c%e7%b4%a2%e5%bc%95%e6%93%8e%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e7%9e%ac%e9%97%b4%e5%ae%8c%e6%88%90%e6%b5%b7%e9%87%8f%e6%95%b0%e6%8d%ae%e6%a3%80%e7%b4%a2%ef%bc%9f.md">26 搜索引擎架构：如何瞬间完成海量数据检索？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 微服务架构：微服务究竟是灵丹还是毒药？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/27%20%e5%be%ae%e6%9c%8d%e5%8a%a1%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%be%ae%e6%9c%8d%e5%8a%a1%e7%a9%b6%e7%ab%9f%e6%98%af%e7%81%b5%e4%b8%b9%e8%bf%98%e6%98%af%e6%af%92%e8%8d%af%ef%bc%9f.md">27 微服务架构：微服务究竟是灵丹还是毒药？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 高性能架构：除了代码，你还可以在哪些地方优化性能？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/28%20%e9%ab%98%e6%80%a7%e8%83%bd%e6%9e%b6%e6%9e%84%ef%bc%9a%e9%99%a4%e4%ba%86%e4%bb%a3%e7%a0%81%ef%bc%8c%e4%bd%a0%e8%bf%98%e5%8f%af%e4%bb%a5%e5%9c%a8%e5%93%aa%e4%ba%9b%e5%9c%b0%e6%96%b9%e4%bc%98%e5%8c%96%e6%80%a7%e8%83%bd%ef%bc%9f.md">28 高性能架构：除了代码，你还可以在哪些地方优化性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 高可用架构：我们为什么感觉不到淘宝应用升级时的停机？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/29%20%e9%ab%98%e5%8f%af%e7%94%a8%e6%9e%b6%e6%9e%84%ef%bc%9a%e6%88%91%e4%bb%ac%e4%b8%ba%e4%bb%80%e4%b9%88%e6%84%9f%e8%a7%89%e4%b8%8d%e5%88%b0%e6%b7%98%e5%ae%9d%e5%ba%94%e7%94%a8%e5%8d%87%e7%ba%a7%e6%97%b6%e7%9a%84%e5%81%9c%e6%9c%ba%ef%bc%9f.md">29 高可用架构：我们为什么感觉不到淘宝应用升级时的停机？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 安全性架构：为什么说用户密码泄漏是程序员的锅？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/30%20%e5%ae%89%e5%85%a8%e6%80%a7%e6%9e%b6%e6%9e%84%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4%e7%94%a8%e6%88%b7%e5%af%86%e7%a0%81%e6%b3%84%e6%bc%8f%e6%98%af%e7%a8%8b%e5%ba%8f%e5%91%98%e7%9a%84%e9%94%85%ef%bc%9f.md">30 安全性架构：为什么说用户密码泄漏是程序员的锅？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 大数据架构：大数据技术架构的思想和原理是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/31%20%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%8a%80%e6%9c%af%e6%9e%b6%e6%9e%84%e7%9a%84%e6%80%9d%e6%83%b3%e5%92%8c%e5%8e%9f%e7%90%86%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">31 大数据架构：大数据技术架构的思想和原理是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 AI与物联网架构：从智能引擎到物联网平台.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/32%20AI%e4%b8%8e%e7%89%a9%e8%81%94%e7%bd%91%e6%9e%b6%e6%9e%84%ef%bc%9a%e4%bb%8e%e6%99%ba%e8%83%bd%e5%bc%95%e6%93%8e%e5%88%b0%e7%89%a9%e8%81%94%e7%bd%91%e5%b9%b3%e5%8f%b0.md">32 AI与物联网架构：从智能引擎到物联网平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 区块链技术架构：区块链到底能做什么？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/33%20%e5%8c%ba%e5%9d%97%e9%93%be%e6%8a%80%e6%9c%af%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%8c%ba%e5%9d%97%e9%93%be%e5%88%b0%e5%ba%95%e8%83%bd%e5%81%9a%e4%bb%80%e4%b9%88%ef%bc%9f.md">33 区块链技术架构：区块链到底能做什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 答疑 互联网需要解决的技术问题是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/33%20%e7%ad%94%e7%96%91%20%e4%ba%92%e8%81%94%e7%bd%91%e9%9c%80%e8%a6%81%e8%a7%a3%e5%86%b3%e7%9a%84%e6%8a%80%e6%9c%af%e9%97%ae%e9%a2%98%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">33 答疑 互联网需要解决的技术问题是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 技术修炼之道：同样工作十几年，为什么有的人成为大厂架构师，有的人失业？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/34%20%e6%8a%80%e6%9c%af%e4%bf%ae%e7%82%bc%e4%b9%8b%e9%81%93%ef%bc%9a%e5%90%8c%e6%a0%b7%e5%b7%a5%e4%bd%9c%e5%8d%81%e5%87%a0%e5%b9%b4%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e6%9c%89%e7%9a%84%e4%ba%ba%e6%88%90%e4%b8%ba%e5%a4%a7%e5%8e%82%e6%9e%b6%e6%9e%84%e5%b8%88%ef%bc%8c%e6%9c%89%e7%9a%84%e4%ba%ba%e5%a4%b1%e4%b8%9a%ef%bc%9f.md">34 技术修炼之道：同样工作十几年，为什么有的人成为大厂架构师，有的人失业？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 技术进阶之道：你和这个星球最顶级的程序员差几个等级？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/35%20%e6%8a%80%e6%9c%af%e8%bf%9b%e9%98%b6%e4%b9%8b%e9%81%93%ef%bc%9a%e4%bd%a0%e5%92%8c%e8%bf%99%e4%b8%aa%e6%98%9f%e7%90%83%e6%9c%80%e9%a1%b6%e7%ba%a7%e7%9a%84%e7%a8%8b%e5%ba%8f%e5%91%98%e5%b7%ae%e5%87%a0%e4%b8%aa%e7%ad%89%e7%ba%a7%ef%bc%9f.md">35 技术进阶之道：你和这个星球最顶级的程序员差几个等级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 技术落地之道：你真的知道自己要解决的问题是什么吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/36%20%e6%8a%80%e6%9c%af%e8%90%bd%e5%9c%b0%e4%b9%8b%e9%81%93%ef%bc%9a%e4%bd%a0%e7%9c%9f%e7%9a%84%e7%9f%a5%e9%81%93%e8%87%aa%e5%b7%b1%e8%a6%81%e8%a7%a3%e5%86%b3%e7%9a%84%e9%97%ae%e9%a2%98%e6%98%af%e4%bb%80%e4%b9%88%e5%90%97%ef%bc%9f.md">36 技术落地之道：你真的知道自己要解决的问题是什么吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 技术沟通之道：如何解决问题？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/37%20%e6%8a%80%e6%9c%af%e6%b2%9f%e9%80%9a%e4%b9%8b%e9%81%93%ef%bc%9a%e5%a6%82%e4%bd%95%e8%a7%a3%e5%86%b3%e9%97%ae%e9%a2%98%ef%bc%9f.md">37 技术沟通之道：如何解决问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 技术管理之道：你真的要转管理吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/38%20%e6%8a%80%e6%9c%af%e7%ae%a1%e7%90%86%e4%b9%8b%e9%81%93%ef%bc%9a%e4%bd%a0%e7%9c%9f%e7%9a%84%e8%a6%81%e8%bd%ac%e7%ae%a1%e7%90%86%e5%90%97%ef%bc%9f.md">38 技术管理之道：你真的要转管理吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 答疑 工作中的交往和沟通，都有哪些小技巧呢？.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/38%20%e7%ad%94%e7%96%91%20%e5%b7%a5%e4%bd%9c%e4%b8%ad%e7%9a%84%e4%ba%a4%e5%be%80%e5%92%8c%e6%b2%9f%e9%80%9a%ef%bc%8c%e9%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%e5%b0%8f%e6%8a%80%e5%b7%a7%e5%91%a2%ef%bc%9f.md">38 答疑 工作中的交往和沟通，都有哪些小技巧呢？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 软件设计文档示例模板.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/%e5%8a%a0%e9%a4%90%20%e8%bd%af%e4%bb%b6%e8%ae%be%e8%ae%a1%e6%96%87%e6%a1%a3%e7%a4%ba%e4%be%8b%e6%a8%a1%e6%9d%bf.md">加餐 软件设计文档示例模板.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 期待未来的你，成为优秀的软件架构师.md" href="/%e4%b8%93%e6%a0%8f/%e5%90%8e%e7%ab%af%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%9538%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%9c%9f%e5%be%85%e6%9c%aa%e6%9d%a5%e7%9a%84%e4%bd%a0%ef%bc%8c%e6%88%90%e4%b8%ba%e4%bc%98%e7%a7%80%e7%9a%84%e8%bd%af%e4%bb%b6%e6%9e%b6%e6%9e%84%e5%b8%88.md">结束语 期待未来的你，成为优秀的软件架构师.md</a>
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
                            <h1 id="title" data-id="09 软件设计实践：如何使用UML完成一个设计文档？" class="title">09 软件设计实践：如何使用UML完成一个设计文档？</h1>
                            <div><p>在上一篇文章中，我们讨论了为什么要建模，以及建模的4+1视图模型，4+1视图模型很好地向我们展示了如何对一个软件的不同方面用不同的模型图进行建模与设计，以完整描述一个软件的业务场景与技术实现。但是软件开发是有阶段性的，在不同的开发阶段用不同的模型图描述业务场景与设计思路，在不同阶段输出不同的设计文档，对于现实的开发更有实践意义。</p>

<p>软件建模与设计过程可以拆分成需求分析、概要设计和详细设计三个阶段。UML规范包含了十多种模型图，常用的有7种：类图、序列图、组件图、部署图、用例图、状态图和活动图。下面我们讨论如何画这7种模型图，以及如何在需求分析、概要设计、详细设计三个阶段使用这7种模型输出合适的设计文档。</p>

<h2 id="类图">类图</h2>

<p>类图是最常见的UML图形，用来描述类的特性和类之间的静态关系。</p>

<p>一个类包含三个部分：类的名字、类的属性列表和类的方法列表。类之间有6种静态关系：关联、依赖、组合、聚合、继承、泛化。把相关的一组类及其关系用一张图画出来，就是类图。</p>

<p>类图主要是在<strong>详细设计阶段</strong>画，如果类图已经设计出来了，那么开发工程师只需要按照类图实现代码就可以了，只要类方法的逻辑不是太复杂，不同的工程师实现出来的代码几乎是一样的，这样可以保证软件的规范、统一。在实践中，我们通常不需要把一个软件所有的类都画出来，把核心的、有代表性的、有一定技术难度的类图画出来，一般就可以了。</p>

<p><img src="assets/84755193120d23e06e098642185bf152.png" alt="img" />
除了在详细设计阶段画类图，在<strong>需求分析阶段</strong>，也可以将关键的领域模型对象用类图画出来，在这个阶段中，我们需要关注的是领域对象的识别及其关系，所以用简化的类图来描述，只画出类的名字及关系就可以了。</p>

<h2 id="序列图">序列图</h2>

<p>类图之外，另一种常用的图是序列图，类图描述类之间的静态关系，序列图则用来描述参与者之间的动态调用关系。</p>

<p>每个参与者有一条垂直向下的生命线，这条线用虚线表示，而参与者之间的消息也从上到下表示其调用的前后顺序关系，这正是序列图这个词的由来。每个生命线都有一个激活条，只有在参与者活动的时候才是激活的。</p>

<p>序列图通常用于表示对象之间的交互，这个对象可以是类对象，也可以是更大粒度的参与者，比如组件、服务器、子系统等，总之，只要是描述不同参与者之间交互的，都可以使用序列图，也就是说，在软<strong>件设计的不同阶段</strong>，都可以画序列图。</p>

<h2 id="组件图">组件图</h2>

<p>组件是比类粒度更大的设计元素，一个组件中通常包含很多个类。组件图有的时候和包图的用途比较接近，组件图通常用来描述物理上的组件，比如一个JAR，一个DLL等等。在实践中，我们进行模块设计的时候更多的是用组件图。</p>

<p><img src="assets/5d4e41aa7769a011ef2bb9f22ee9808d.png" alt="img" />
组件图描述组件之间的静态关系，主要是依赖关系，如果想要描述组件之间的动态调用关系，可以使用组件序列图，以组件作为参与者，描述组件之间的消息调用关系。</p>

<p>因为组件的粒度比较粗，通常用以描述和设计软件的模块及其之间的关系，需要在设计早期阶段就画出来，因此组件图一般用在<strong>概要设计阶段</strong>。</p>

<h2 id="部署图">部署图</h2>

<p>部署图描述软件系统的最终部署情况，比如需要部署多少服务器，关键组件都部署在哪些服务器上。</p>

<p><img src="assets/32931c58184b79744efa559bf4e0b00c.png" alt="img" />
部署图是软件系统最终物理呈现的蓝图，根据部署图，所有相关者，诸如客户、老板、工程师都能清晰地了解到最终运行的系统在物理上是什么样子，和现有的系统服务器的关系，和第三方服务器的关系。根据部署图，还可以估算服务器和第三方软件的采购成本。</p>

<p>因此部署图是整个软件设计模型中，比较宏观的一种图，是在设计早期就需要画的一种模型图。根据部署图，各方可以讨论对这个方案是否认可。只有对部署图达成共识，才能继续后面的细节设计。部署图主要用在<strong>概要设计阶段</strong>。</p>

<h2 id="用例图">用例图</h2>

<p>用例图主要用在<strong>需求分析阶段</strong>，通过反映用户和软件系统的交互，描述系统的功能需求。</p>

<p><img src="assets/b79540c40111f3be97fc0b81d5a3060e.png" alt="img" />
图中小人形象的元素，被称为角色，角色可以是人，也可以是其他的系统。系统的功能可能会很复杂，所以一张用例图可能只包含其中一小部分功能，这些功能被一个矩形框框起来，这个矩形框被称为用例的边界。框里的椭圆表示一个一个的功能，功能之间可以调用依赖，也可以进行功能扩展。</p>

<p>因为用例图中功能描述比较简单，通常还需要对用例图配以文字说明，形成需求文档。</p>

<h2 id="状态图">状态图</h2>

<p>状态图用来展示单个对象生命周期的状态变迁。</p>

<p>业务系统中，很多重要的领域对象都有比较复杂的状态变迁，比如账号，有创建状态、激活状态、冻结状态、欠费状态等等各种状态。此外，用户、订单、商品、红包这些常见的领域模型都有多种状态。</p>

<p>这些状态的变迁描述可以在用例图中用文字描述，随着角色的各种操作而改变，但是用这种方式描述，状态散乱在各处，不要说开发的时候容易搞错，就是产品经理自己在设计的时候，也容易搞错对象的状态变迁。</p>

<p>UML的状态图可以很好地解决这一问题，一张状态图描述一个对象生命周期的各种状态，及其变迁的关系。如图所示，门的状态有开opened、关closed和锁locked三种，状态与变迁关系用一张状态图就可以搞定。</p>

<p><img src="assets/34ed399aa400f124ead77bad31468104.png" alt="img" />
状态图要在<strong>需求分析阶段</strong>画，描述状态变迁的逻辑关系，在<strong>详细设计</strong>阶段也要画，这个时候，状态要用枚举值表示，以指导具体的开发。</p>

<h2 id="活动图">活动图</h2>

<p>活动图主要用来描述过程逻辑和业务流程。UML中没有流程图，很多时候，人们用活动图代替流程图。</p>

<p><img src="assets/976fdc5c09c37d7b547d6c5d64bea107.png" alt="img" />
活动图和早期流程图的图形元素也很接近，实心圆代表流程开始，空心圆代表流程结束，圆角矩形表示活动，菱形表示分支判断。</p>

<p>此外，活动图引入了一个重要的概念——泳道。活动图可以根据活动的范围，将活动根据领域、系统和角色等划分到不同的泳道中，使流程边界更加清晰。</p>

<p>活动图也比较有普适性，可以在<strong>需求分析阶段</strong>描述业务流程，也可以在<strong>概要设计阶段</strong>描述子系统和组件的交互，还可以在<strong>详细设计阶段</strong>描述一个类方法内部的计算流程。</p>

<h2 id="使用合适的uml模型构建一个设计文档">使用合适的UML模型构建一个设计文档</h2>

<p>UML模型图本身并不复杂，几分钟的时间就可以学习一个模型图的画法。但难的是如何在合适的场合下用正确的UML模型表达自己的设计意图，形成一套完整的软件模型，进而组织成一个言之有物，层次分明，既可以指导开发，又可以在团队内外达成共识的设计文档。</p>

<p>下面我们就从软件设计的不同阶段这一维度，重新梳理下如何使用正确的模型进行软件建模。</p>

<p>在<strong>需求分析阶段</strong>，主要是通过用例图来描述系统的功能与使用场景；对于关键的业务流程，可以通过活动图描述；如果在需求阶段就提出要和现有的某些子系统整合，那么可以通过时序图描述新系统和原来的子系统的调用关系；可以通过简化的类图进行领域模型抽象，并描述核心领域对象之间的关系；如果某些对象内部会有复杂的状态变化，比如用户、订单这些，可以用状态图进行描述。</p>

<p>在<strong>概要设计阶段</strong>，通过部署图描述系统最终的物理蓝图；通过组件图以及组件时序图设计软件主要模块及其关系；还可以通过组件活动图描述组件间的流程逻辑。</p>

<p>在<strong>详细设计阶段</strong>，主要输出的就是类图和类的时序图，指导最终的代码开发，如果某个类方法内部有比较复杂的逻辑，那么可以用画方法的活动图进行描述。</p>

<p>下一篇文章我会通过一个示例模板为你展示设计文档的写法和UML模型在文档中的应用。</p>

<h2 id="小结">小结</h2>

<p>UML建模可以很复杂，也可以很简单，简单掌握类图、时序图、组件图、部署图、用例图、状态图、活动图这7种模型图，根据场景的不同，灵活在需求分析、概要设计和详细设计阶段绘制对应的模型图，可以实实在在地做好软件建模，搞好系统设计，做一个掌控局面、引领技术团队的架构师。</p>

<p>画UML的工具，可以是很复杂的，用像EA这样的大型软件设计工具，不过是收费的，也可以是draw.io这样在线、免费的工具，一般来说，都建议先从简单的用起。</p>

<h2 id="思考题">思考题</h2>

<p>你现在开发的软件是否会用到UML建模呢？如果没有，你觉得应该画哪些UML模型？又该如何画呢？</p>

<p>欢迎你在评论区写下你的思考，我会和你一起交流，也欢迎把这篇文章分享给你的朋友或者同事，一起交流进步吧！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#bcd0d0d085888d8d8c8bfcdbd1ddd5d092dfd3d1" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f140a5b5c289508',t:'MTczNDA3Mjg0MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=07&#32;&#32;错误处理：如何通过&#32;error、deferred、panic&#32;等处理错误？>
        <link rel="icon" href="/static/favicon.png">
        <title>07  错误处理：如何通过 error、deferred、panic 等处理错误？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  Go 为开发者的需求设计，带你实现高效工作.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20Go%20%e4%b8%ba%e5%bc%80%e5%8f%91%e8%80%85%e7%9a%84%e9%9c%80%e6%b1%82%e8%ae%be%e8%ae%a1%ef%bc%8c%e5%b8%a6%e4%bd%a0%e5%ae%9e%e7%8e%b0%e9%ab%98%e6%95%88%e5%b7%a5%e4%bd%9c.md">00 开篇词  Go 为开发者的需求设计，带你实现高效工作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  基础入门：编写你的第一个 Go 语言程序.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/01%20%20%e5%9f%ba%e7%a1%80%e5%85%a5%e9%97%a8%ef%bc%9a%e7%bc%96%e5%86%99%e4%bd%a0%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aa%20Go%20%e8%af%ad%e8%a8%80%e7%a8%8b%e5%ba%8f.md">01  基础入门：编写你的第一个 Go 语言程序.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  数据类型：你必须掌握的数据类型有哪些？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/02%20%20%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%ef%bc%9a%e4%bd%a0%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%9f.md">02  数据类型：你必须掌握的数据类型有哪些？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  控制结构：if、for、switch 逻辑语句的那些事儿.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/03%20%20%e6%8e%a7%e5%88%b6%e7%bb%93%e6%9e%84%ef%bc%9aif%e3%80%81for%e3%80%81switch%20%e9%80%bb%e8%be%91%e8%af%ad%e5%8f%a5%e7%9a%84%e9%82%a3%e4%ba%9b%e4%ba%8b%e5%84%bf.md">03  控制结构：if、for、switch 逻辑语句的那些事儿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  集合类型：如何正确使用 array、slice 和 map？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/04%20%20%e9%9b%86%e5%90%88%e7%b1%bb%e5%9e%8b%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e4%bd%bf%e7%94%a8%20array%e3%80%81slice%20%e5%92%8c%20map%ef%bc%9f.md">04  集合类型：如何正确使用 array、slice 和 map？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  函数和方法：Go 语言中的函数和方法到底有什么不同？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/05%20%20%e5%87%bd%e6%95%b0%e5%92%8c%e6%96%b9%e6%b3%95%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e4%b8%ad%e7%9a%84%e5%87%bd%e6%95%b0%e5%92%8c%e6%96%b9%e6%b3%95%e5%88%b0%e5%ba%95%e6%9c%89%e4%bb%80%e4%b9%88%e4%b8%8d%e5%90%8c%ef%bc%9f.md">05  函数和方法：Go 语言中的函数和方法到底有什么不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  struct 和 interface：结构体与接口都实现了哪些功能？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/06%20%20struct%20%e5%92%8c%20interface%ef%bc%9a%e7%bb%93%e6%9e%84%e4%bd%93%e4%b8%8e%e6%8e%a5%e5%8f%a3%e9%83%bd%e5%ae%9e%e7%8e%b0%e4%ba%86%e5%93%aa%e4%ba%9b%e5%8a%9f%e8%83%bd%ef%bc%9f.md">06  struct 和 interface：结构体与接口都实现了哪些功能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  错误处理：如何通过 error、deferred、panic 等处理错误？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/07%20%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20error%e3%80%81deferred%e3%80%81panic%20%e7%ad%89%e5%a4%84%e7%90%86%e9%94%99%e8%af%af%ef%bc%9f.md">07  错误处理：如何通过 error、deferred、panic 等处理错误？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  并发基础：Goroutines 和 Channels 的声明与使用.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/08%20%20%e5%b9%b6%e5%8f%91%e5%9f%ba%e7%a1%80%ef%bc%9aGoroutines%20%e5%92%8c%20Channels%20%e7%9a%84%e5%a3%b0%e6%98%8e%e4%b8%8e%e4%bd%bf%e7%94%a8.md">08  并发基础：Goroutines 和 Channels 的声明与使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  同步原语：sync 包让你对并发控制得心应手.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/09%20%20%e5%90%8c%e6%ad%a5%e5%8e%9f%e8%af%ad%ef%bc%9async%20%e5%8c%85%e8%ae%a9%e4%bd%a0%e5%af%b9%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e5%be%97%e5%bf%83%e5%ba%94%e6%89%8b.md">09  同步原语：sync 包让你对并发控制得心应手.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  Context：你必须掌握的多线程并发控制神器.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/10%20%20Context%ef%bc%9a%e4%bd%a0%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84%e5%a4%9a%e7%ba%bf%e7%a8%8b%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e7%a5%9e%e5%99%a8.md">10  Context：你必须掌握的多线程并发控制神器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  并发模式：Go 语言中即学即用的高效并发模式.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/11%20%20%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%bc%8f%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e4%b8%ad%e5%8d%b3%e5%ad%a6%e5%8d%b3%e7%94%a8%e7%9a%84%e9%ab%98%e6%95%88%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%bc%8f.md">11  并发模式：Go 语言中即学即用的高效并发模式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  指针详解：在什么情况下应该使用指针？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/12%20%20%e6%8c%87%e9%92%88%e8%af%a6%e8%a7%a3%ef%bc%9a%e5%9c%a8%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e5%ba%94%e8%af%a5%e4%bd%bf%e7%94%a8%e6%8c%87%e9%92%88%ef%bc%9f.md">12  指针详解：在什么情况下应该使用指针？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  参数传递：值、引用及指针之间的区别？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/13%20%20%e5%8f%82%e6%95%b0%e4%bc%a0%e9%80%92%ef%bc%9a%e5%80%bc%e3%80%81%e5%bc%95%e7%94%a8%e5%8f%8a%e6%8c%87%e9%92%88%e4%b9%8b%e9%97%b4%e7%9a%84%e5%8c%ba%e5%88%ab%ef%bc%9f.md">13  参数传递：值、引用及指针之间的区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  内存分配：new 还是 make？什么情况下该用谁？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/14%20%20%e5%86%85%e5%ad%98%e5%88%86%e9%85%8d%ef%bc%9anew%20%e8%bf%98%e6%98%af%20make%ef%bc%9f%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e8%af%a5%e7%94%a8%e8%b0%81%ef%bc%9f.md">14  内存分配：new 还是 make？什么情况下该用谁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  运行时反射：字符串和结构体之间如何转换？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/15%20%20%e8%bf%90%e8%a1%8c%e6%97%b6%e5%8f%8d%e5%b0%84%ef%bc%9a%e5%ad%97%e7%ac%a6%e4%b8%b2%e5%92%8c%e7%bb%93%e6%9e%84%e4%bd%93%e4%b9%8b%e9%97%b4%e5%a6%82%e4%bd%95%e8%bd%ac%e6%8d%a2%ef%bc%9f.md">15  运行时反射：字符串和结构体之间如何转换？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  非类型安全：让你既爱又恨的 unsafe.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/16%20%20%e9%9d%9e%e7%b1%bb%e5%9e%8b%e5%ae%89%e5%85%a8%ef%bc%9a%e8%ae%a9%e4%bd%a0%e6%97%a2%e7%88%b1%e5%8f%88%e6%81%a8%e7%9a%84%20unsafe.md">16  非类型安全：让你既爱又恨的 unsafe.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  SliceHeader：slice 如何高效处理数据？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/17%20%20SliceHeader%ef%bc%9aslice%20%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%a4%84%e7%90%86%e6%95%b0%e6%8d%ae%ef%bc%9f.md">17  SliceHeader：slice 如何高效处理数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  质量保证：Go 语言如何通过测试保证质量？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/18%20%20%e8%b4%a8%e9%87%8f%e4%bf%9d%e8%af%81%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%b5%8b%e8%af%95%e4%bf%9d%e8%af%81%e8%b4%a8%e9%87%8f%ef%bc%9f.md">18  质量保证：Go 语言如何通过测试保证质量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  性能优化：Go 语言如何进行代码检查和优化？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/19%20%20%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e4%bb%a3%e7%a0%81%e6%a3%80%e6%9f%a5%e5%92%8c%e4%bc%98%e5%8c%96%ef%bc%9f.md">19  性能优化：Go 语言如何进行代码检查和优化？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  协作开发：模块化管理为什么能够提升研发效能？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/20%20%20%e5%8d%8f%e4%bd%9c%e5%bc%80%e5%8f%91%ef%bc%9a%e6%a8%a1%e5%9d%97%e5%8c%96%e7%ae%a1%e7%90%86%e4%b8%ba%e4%bb%80%e4%b9%88%e8%83%bd%e5%a4%9f%e6%8f%90%e5%8d%87%e7%a0%94%e5%8f%91%e6%95%88%e8%83%bd%ef%bc%9f.md">20  协作开发：模块化管理为什么能够提升研发效能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  网络编程：Go 语言如何玩转 RESTful API 服务？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/21%20%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e7%8e%a9%e8%bd%ac%20RESTful%20API%20%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">21  网络编程：Go 语言如何玩转 RESTful API 服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  网络编程：Go 语言如何通过 RPC 实现跨平台服务？.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/22%20%20%e7%bd%91%e7%bb%9c%e7%bc%96%e7%a8%8b%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20RPC%20%e5%ae%9e%e7%8e%b0%e8%b7%a8%e5%b9%b3%e5%8f%b0%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">22  网络编程：Go 语言如何通过 RPC 实现跨平台服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结束语  你的 Go 语言成长之路.md" href="/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/23%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bd%a0%e7%9a%84%20Go%20%e8%af%ad%e8%a8%80%e6%88%90%e9%95%bf%e4%b9%8b%e8%b7%af.md">23 结束语  你的 Go 语言成长之路.md</a>
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
                            <h1 id="title" data-id="07  错误处理：如何通过 error、deferred、panic 等处理错误？" class="title">07  错误处理：如何通过 error、deferred、panic 等处理错误？</h1>
                            <div><p>上节课我为你讲解了结构体和接口，并留了一个小作业，让你自己练习实现有两个方法的接口。现在我就以“人既会走也会跑”为例进行讲解。</p>

<p>首先定义一个接口 WalkRun，它有两个方法 Walk 和 Run，如下面的代码所示：</p>

<pre><code class="language-go">type WalkRun interface {
   Walk()
   Run()
}
</code></pre>

<p>现在就可以让结构体 person 实现这个接口了，如下所示：</p>

<pre><code class="language-go">func (p *person) Walk(){
   fmt.Printf(&quot;%s能走\n&quot;,p.name)
}

func (p *person) Run(){
   fmt.Printf(&quot;%s能跑\n&quot;,p.name)
}
</code></pre>

<p>关键点在于，让接口的每个方法都实现，也就实现了这个接口。</p>

<blockquote>
<p>提示：%s 是占位符，和 p.name 对应，也就是 p.name 的值，具体可以参考 fmt.Printf 函数的文档。</p>
</blockquote>

<p>下面进行本节课的讲解。这节课我会带你学习 Go 语言的错误和异常，在我们编写程序的时候，可能会遇到一些问题，该怎么处理它们呢？</p>

<h3 id="错误">错误</h3>

<p>在 Go 语言中，错误是可以预期的，并且不是非常严重，不会影响程序的运行。对于这类问题，可以用返回错误给调用者的方法，让调用者自己决定如何处理。</p>

<h4 id="error-接口">error 接口</h4>

<p>在 Go 语言中，错误是通过内置的 error 接口表示的。它非常简单，只有一个 Error 方法用来返回具体的错误信息，如下面的代码所示：</p>

<pre><code class="language-go">type error interface {
   Error() string
}
</code></pre>

<p>在下面的代码中，我演示了一个字符串转整数的例子：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">func main() {
   i,err:=strconv.Atoi(&quot;a&quot;)
   if err!=nil {
      fmt.Println(err)
   }else {
      fmt.Println(i)
   }
}
</code></pre>

<p>这里我故意使用了字符串 &ldquo;a&rdquo;，尝试把它转为整数。我们知道 &ldquo;a&rdquo; 是无法转为数字的，所以运行这段程序，会打印出如下错误信息：</p>

<pre><code class="language-log">strconv.Atoi: parsing &quot;a&quot;: invalid syntax
</code></pre>

<p>这个错误信息就是通过接口 error 返回的。我们来看关于函数 strconv.Atoi 的定义，如下所示：</p>

<pre><code class="language-go">func Atoi(s string) (int, error)
</code></pre>

<p>一般而言，error 接口用于当方法或者函数执行遇到错误时进行返回，而且是第二个返回值。通过这种方式，可以让调用者自己根据错误信息决定如何进行下一步处理。</p>

<blockquote>
<p>小提示：因为方法和函数基本上差不多，区别只在于有无接收者，所以以后当我称方法或函数，表达的是一个意思，不会把这两个名字都写出来。</p>
</blockquote>

<h4 id="error-工厂函数">error 工厂函数</h4>

<p>除了可以使用其他函数，自己定义的函数也可以返回错误信息给调用者，如下面的代码所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">func add(a,b int) (int,error){
   if a&lt;0 || b&lt;0 {
      return 0,errors.New(&quot;a或者b不能为负数&quot;)
   }else {
      return a+b,nil
   }
}
</code></pre>

<p>add 函数会在 a 或者 b 任何一个为负数的情况下，返回一个错误信息，如果 a、b 都不为负数，错误信息部分会返回 nil，这也是常见的做法。所以调用者可以通过错误信息是否为 nil 进行判断。</p>

<p>下面的 add 函数示例，是使用 errors.New 这个工厂函数生成的错误信息，它接收一个字符串参数，返回一个 error 接口，这些在上节课的结构体和接口部分有过详细介绍，不再赘述。</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">sum,err:=add(-1,2)
if err!=nil {
   fmt.Println(err)
}else {
   fmt.Println(sum)
}
</code></pre>

<h4 id="自定义-error">自定义 error</h4>

<p>你可能会想，上面采用工厂返回错误信息的方式只能传递一个字符串，也就是携带的信息只有字符串，如果想要携带更多信息（比如错误码信息）该怎么办呢？这个时候就需要自定义 error 。</p>

<p>自定义 error 其实就是先自定义一个新类型，比如结构体，然后让这个类型实现 error 接口，如下面的代码所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">type commonError struct {
   errorCode int //错误码
   errorMsg string //错误信息
}

func (ce *commonError) Error() string{
   return ce.errorMsg
}
</code></pre>

<p>有了自定义的 error，就可以使用它携带更多的信息，现在我改造上面的例子，返回刚刚自定义的 commonError，如下所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">return 0, &amp;commonError{
   errorCode: 1,
   errorMsg:  &quot;a或者b不能为负数&quot;}
</code></pre>

<p>我通过字面量的方式创建一个 *commonError 返回，其中 errorCode 值为 1，errorMsg 值为 “a 或者 b 不能为负数”。</p>

<h4 id="error-断言">error 断言</h4>

<p>有了自定义的 error，并且携带了更多的错误信息后，就可以使用这些信息了。你需要先把返回的 error 接口转换为自定义的错误类型，用到的知识是上节课的类型断言。</p>

<p>下面代码中的 err.(*commonError) 就是类型断言在 error 接口上的应用，也可以称为 error 断言。</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">sum, err := add(-1, 2)
if cm,ok:=err.(*commonError);ok{
   fmt.Println(&quot;错误代码为:&quot;,cm.errorCode,&quot;，错误信息为：&quot;,cm.errorMsg)
} else {
   fmt.Println(sum)
}
</code></pre>

<p>如果返回的 ok 为 true，说明 error 断言成功，正确返回了 *commonError 类型的变量 cm，所以就可以像示例中一样使用变量 cm 的 errorCode 和 errorMsg 字段信息了。</p>

<h3 id="错误嵌套">错误嵌套</h3>

<h4 id="error-wrapping">Error Wrapping</h4>

<p>error 接口虽然比较简洁，但是功能也比较弱。想象一下，假如我们有这样的需求：基于一个存在的 error 再生成一个 error，需要怎么做呢？这就是错误嵌套。</p>

<p>这种需求是存在的，比如调用一个函数，返回了一个错误信息 error，在不想丢失这个 error 的情况下，又想添加一些额外信息返回新的 error。这时候，我们首先想到的应该是自定义一个 struct，如下面的代码所示：</p>

<pre><code class="language-go">type MyError struct {
    err error
    msg string
}
</code></pre>

<p>这个结构体有两个字段，其中 error 类型的 err 字段用于存放已存在的 error，string 类型的 msg 字段用于存放新的错误信息，<strong>这种方式就是 error 的嵌套</strong>。</p>

<p>现在让 MyError 这个 struct 实现 error 接口，然后在初始化 MyError 的时候传递存在的 error 和新的错误信息，如下面的代码所示：</p>

<pre><code class="language-go">func (e *MyError) Error() string {
    return e.err.Error() + e.msg
}

func main() {
    //err是一个存在的错误，可以从另外一个函数返回
    newErr := MyError{err, &quot;数据上传问题&quot;}
}
</code></pre>

<p>这种方式可以满足我们的需求，但是非常烦琐，因为既要定义新的类型还要实现 error 接口。所以从 Go 语言 1.13 版本开始，Go 标准库新增了 Error Wrapping 功能，让我们可以基于一个存在的 error 生成新的 error，并且可以保留原 error 信息，如下面的代码所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">e := errors.New(&quot;原始错误e&quot;)
w := fmt.Errorf(&quot;Wrap了一个错误:%w&quot;, e)
fmt.Println(w)
</code></pre>

<p>Go 语言没有提供 Wrap 函数，而是扩展了 fmt.Errorf 函数，然后加了一个 %w，通过这种方式，便可以生成 wrapping error。</p>

<h4 id="errors-unwrap-函数">errors.Unwrap 函数</h4>

<p>既然 error 可以包裹嵌套生成一个新的 error，那么也可以被解开，即通过 errors.Unwrap 函数得到被嵌套的 error。</p>

<p>Go 语言提供了 errors.Unwrap 用于获取被嵌套的 error，比如以上例子中的错误变量 w ，就可以对它进行 unwrap，获取被嵌套的原始错误 e。</p>

<p>下面我们运行以下代码：</p>

<pre><code class="language-go">fmt.Println(errors.Unwrap(w))
</code></pre>

<p>可以看到这样的信息，即“原始错误 e”。</p>

<pre><code class="language-log">原始错误e
</code></pre>

<h4 id="errors-is-函数">errors.Is 函数</h4>

<p>有了 Error Wrapping 后，你会发现原来用的判断两个 error 是不是同一个 error 的方法失效了，比如 Go 语言标准库经常用到的如下代码中的方式：</p>

<pre><code class="language-go">if err == os.ErrExist
</code></pre>

<p>为什么会出现这种情况呢？由于 Go 语言的 Error Wrapping 功能，令人不知道返回的 err 是否被嵌套，又嵌套了几层？</p>

<p>于是 Go 语言为我们提供了 errors.Is 函数，用来判断两个 error 是否是同一个，如下所示：</p>

<pre><code class="language-go">func Is(err, target error) bool
</code></pre>

<p>以上就是errors.Is 函数的定义，可以解释为：</p>

<ul>
<li>如果 err 和 target 是同一个，那么返回 true。</li>
<li>如果 err 是一个 wrapping error，target 也包含在这个嵌套 error 链中的话，也返回 true。</li>
</ul>

<p>可以简单地概括为，两个 error 相等或 err 包含 target 的情况下返回 true，其余返回 false。我们可以用上面的示例判断错误 w 中是否包含错误 e，试试运行下面的代码，来看打印的结果是不是 true。</p>

<pre><code class="language-go">fmt.Println(errors.Is(w,e))
</code></pre>

<h4 id="errors-as-函数">errors.As 函数</h4>

<p>同样的原因，有了 error 嵌套后，error 断言也不能用了，因为你不知道一个 error 是否被嵌套，又嵌套了几层。所以 Go 语言为解决这个问题提供了 errors.As 函数，比如前面 error 断言的例子，可以使用 errors.As 函数重写，效果是一样的，如下面的代码所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">var cm *commonError
if errors.As(err,&amp;cm){
   fmt.Println(&quot;错误代码为:&quot;,cm.errorCode,&quot;，错误信息为：&quot;,cm.errorMsg)
} else {
   fmt.Println(sum)
}
</code></pre>

<p>所以在 Go 语言提供的 Error Wrapping 能力下，我们写的代码要尽可能地使用 Is、As 这些函数做判断和转换。</p>

<h3 id="deferred-函数">Deferred 函数</h3>

<p>在一个自定义函数中，你打开了一个文件，然后需要关闭它以释放资源。不管你的代码执行了多少分支，是否出现了错误，文件是一定要关闭的，这样才能保证资源的释放。</p>

<p>如果这个事情由开发人员来做，随着业务逻辑的复杂会变得非常麻烦，而且还有可能会忘记关闭。基于这种情况，Go 语言为我们提供了 defer 函数，可以保证文件关闭后一定会被执行，不管你自定义的函数出现异常还是错误。</p>

<p>下面的代码是 Go 语言标准包 ioutil 中的 ReadFile 函数，它需要打开一个文件，然后通过 defer 关键字确保在 ReadFile 函数执行结束后，f.Close() 方法被执行，这样文件的资源才一定会释放。</p>

<pre><code class="language-go">func ReadFile(filename string) ([]byte, error) {
   f, err := os.Open(filename)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   //省略无关代码
   return readAll(f, n)
}
</code></pre>

<p>defer 关键字用于修饰一个函数或者方法，使得该函数或者方法在返回前才会执行，也就说被延迟，但又可以保证一定会执行。</p>

<p>以上面的 ReadFile 函数为例，被 defer 修饰的 f.Close 方法延迟执行，也就是说会先执行 readAll(f, n)，然后在整个 ReadFile 函数 return 之前执行 f.Close 方法。</p>

<p>defer 语句常被用于成对的操作，如文件的打开和关闭，加锁和释放锁，连接的建立和断开等。不管多么复杂的操作，都可以保证资源被正确地释放。</p>

<h3 id="panic-异常">Panic 异常</h3>

<p>Go 语言是一门静态的强类型语言，很多问题都尽可能地在编译时捕获，但是有一些只能在运行时检查，比如数组越界访问、不相同的类型强制转换等，这类运行时的问题会引起 panic 异常。</p>

<p>除了运行时可以产生 panic 外，我们自己也可以抛出 panic 异常。假设我需要连接 MySQL 数据库，可以写一个连接 MySQL 的函数connectMySQL，如下面的代码所示：</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">func connectMySQL(ip,username,password string){
   if ip ==&quot;&quot; {
      panic(&quot;ip不能为空&quot;)
   }
   //省略其他代码
}
</code></pre>

<p>在 connectMySQL 函数中，如果 ip 为空会直接抛出 panic 异常。这种逻辑是正确的，因为数据库无法连接成功的话，整个程序运行起来也没有意义，所以就抛出 panic 终止程序的运行。</p>

<p>panic 是 Go 语言内置的函数，可以接受 interface{} 类型的参数，也就是任何类型的值都可以传递给 panic 函数，如下所示：</p>

<pre><code class="language-go">func panic(v interface{})
</code></pre>

<blockquote>
<p>小提示：interface{} 是空接口的意思，在 Go 语言中代表任意类型。</p>
</blockquote>

<p>panic 异常是一种非常严重的情况，会让程序中断运行，使程序崩溃，所以<strong>如果是不影响程序运行的错误，不要使用 panic，使用普通错误 error 即可。</strong></p>

<p><img src="assets/CgqCHl-15ZSAAsw5AAUnpsfN34w061.png" alt="pDE7ppQNyfRSIn1Q__thumbnail.png" /></p>

<h3 id="recover-捕获-panic-异常">Recover 捕获 Panic 异常</h3>

<p>通常情况下，我们不对 panic 异常做任何处理，因为既然它是影响程序运行的异常，就让它直接崩溃即可。但是也的确有一些特例，比如在程序崩溃前做一些资源释放的处理，这时候就需要从 panic 异常中恢复，才能完成处理。</p>

<p>在 Go 语言中，可以通过内置的 recover 函数恢复 panic 异常。因为在程序 panic 异常崩溃的时候，只有被 defer 修饰的函数才能被执行，所以 recover 函数要结合 defer 关键字使用才能生效。</p>

<p>下面的示例是通过 defer 关键字 + 匿名函数 + recover 函数从 panic 异常中恢复的方式。</p>

<p><strong><em>ch07/main.go</em></strong></p>

<pre><code class="language-go">func main() {
   defer func() {
      if p:=recover();p!=nil{
         fmt.Println(p)
      }
   }()
   connectMySQL(&quot;&quot;,&quot;root&quot;,&quot;123456&quot;)
}
</code></pre>

<p>运行这个代码，可以看到如下的打印输出，这证明 recover 函数成功捕获了 panic 异常。</p>

<pre><code class="language-log">ip 不能为空
</code></pre>

<p>通过这个输出的结果也可以发现，recover 函数返回的值就是通过 panic 函数传递的参数值。</p>

<h3 id="总结">总结</h3>

<p>这节课主要讲了 Go 语言的错误处理机制，包括 error、defer、panic 等。在 error、panic 这两种错误机制中，Go 语言更提倡 error 这种轻量错误，而不是 panic。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#244848481d1015151413644349454d480a474b49" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0ce70b7f1a94de',t:'MTczMzk5Nzk5NC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
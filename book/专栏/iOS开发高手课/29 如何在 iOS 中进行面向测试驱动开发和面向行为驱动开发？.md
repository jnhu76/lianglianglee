<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;如何在&#32;iOS&#32;中进行面向测试驱动开发和面向行为驱动开发？>
        <link rel="icon" href="/static/favicon.png">
        <title>29 如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 锚定一个点，然后在这个点上深耕.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e9%94%9a%e5%ae%9a%e4%b8%80%e4%b8%aa%e7%82%b9%ef%bc%8c%e7%84%b6%e5%90%8e%e5%9c%a8%e8%bf%99%e4%b8%aa%e7%82%b9%e4%b8%8a%e6%b7%b1%e8%80%95.md">00 开篇词 锚定一个点，然后在这个点上深耕.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 建立你自己的iOS开发知识体系.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/01%20%e5%bb%ba%e7%ab%8b%e4%bd%a0%e8%87%aa%e5%b7%b1%e7%9a%84iOS%e5%bc%80%e5%8f%91%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb.md">01 建立你自己的iOS开发知识体系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 App 启动速度怎么做优化与监控？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/02%20App%20%e5%90%af%e5%8a%a8%e9%80%9f%e5%ba%a6%e6%80%8e%e4%b9%88%e5%81%9a%e4%bc%98%e5%8c%96%e4%b8%8e%e7%9b%91%e6%8e%a7%ef%bc%9f.md">02 App 启动速度怎么做优化与监控？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 Auto Layout 是怎么进行自动布局的，性能如何？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/03%20Auto%20Layout%20%e6%98%af%e6%80%8e%e4%b9%88%e8%bf%9b%e8%a1%8c%e8%87%aa%e5%8a%a8%e5%b8%83%e5%b1%80%e7%9a%84%ef%bc%8c%e6%80%a7%e8%83%bd%e5%a6%82%e4%bd%95%ef%bc%9f.md">03 Auto Layout 是怎么进行自动布局的，性能如何？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 项目大了人员多了，架构怎么设计更合理？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/04%20%e9%a1%b9%e7%9b%ae%e5%a4%a7%e4%ba%86%e4%ba%ba%e5%91%98%e5%a4%9a%e4%ba%86%ef%bc%8c%e6%9e%b6%e6%9e%84%e6%80%8e%e4%b9%88%e8%ae%be%e8%ae%a1%e6%9b%b4%e5%90%88%e7%90%86%ef%bc%9f.md">04 项目大了人员多了，架构怎么设计更合理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 链接器：符号是怎么绑定到地址上的？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/05%20%e9%93%be%e6%8e%a5%e5%99%a8%ef%bc%9a%e7%ac%a6%e5%8f%b7%e6%98%af%e6%80%8e%e4%b9%88%e7%bb%91%e5%ae%9a%e5%88%b0%e5%9c%b0%e5%9d%80%e4%b8%8a%e7%9a%84%ef%bc%9f.md">05 链接器：符号是怎么绑定到地址上的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 App 如何通过注入动态库的方式实现极速编译调试？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/06%20App%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%b3%a8%e5%85%a5%e5%8a%a8%e6%80%81%e5%ba%93%e7%9a%84%e6%96%b9%e5%bc%8f%e5%ae%9e%e7%8e%b0%e6%9e%81%e9%80%9f%e7%bc%96%e8%af%91%e8%b0%83%e8%af%95%ef%bc%9f.md">06 App 如何通过注入动态库的方式实现极速编译调试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Clang、Infer 和 OCLint ，我们应该使用谁来做静态分析？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/07%20Clang%e3%80%81Infer%20%e5%92%8c%20OCLint%20%ef%bc%8c%e6%88%91%e4%bb%ac%e5%ba%94%e8%af%a5%e4%bd%bf%e7%94%a8%e8%b0%81%e6%9d%a5%e5%81%9a%e9%9d%99%e6%80%81%e5%88%86%e6%9e%90%ef%bc%9f.md">07 Clang、Infer 和 OCLint ，我们应该使用谁来做静态分析？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 如何利用 Clang 为 App 提质？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/08%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20Clang%20%e4%b8%ba%20App%20%e6%8f%90%e8%b4%a8%ef%bc%9f.md">08 如何利用 Clang 为 App 提质？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 无侵入的埋点方案如何实现？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/09%20%e6%97%a0%e4%be%b5%e5%85%a5%e7%9a%84%e5%9f%8b%e7%82%b9%e6%96%b9%e6%a1%88%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">09 无侵入的埋点方案如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 包大小：如何从资源和代码层面实现全方位瘦身？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/10%20%e5%8c%85%e5%a4%a7%e5%b0%8f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e8%b5%84%e6%ba%90%e5%92%8c%e4%bb%a3%e7%a0%81%e5%b1%82%e9%9d%a2%e5%ae%9e%e7%8e%b0%e5%85%a8%e6%96%b9%e4%bd%8d%e7%98%a6%e8%ba%ab%ef%bc%9f.md">10 包大小：如何从资源和代码层面实现全方位瘦身？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 热点问题答疑（一）：基础模块问题答疑.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/11%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%9f%ba%e7%a1%80%e6%a8%a1%e5%9d%97%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">11 热点问题答疑（一）：基础模块问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 iOS 崩溃千奇百怪，如何全面监控？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/12%20iOS%20%e5%b4%a9%e6%ba%83%e5%8d%83%e5%a5%87%e7%99%be%e6%80%aa%ef%bc%8c%e5%a6%82%e4%bd%95%e5%85%a8%e9%9d%a2%e7%9b%91%e6%8e%a7%ef%bc%9f.md">12 iOS 崩溃千奇百怪，如何全面监控？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 如何利用 RunLoop 原理去监控卡顿？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/13%20%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%20RunLoop%20%e5%8e%9f%e7%90%86%e5%8e%bb%e7%9b%91%e6%8e%a7%e5%8d%a1%e9%a1%bf%ef%bc%9f.md">13 如何利用 RunLoop 原理去监控卡顿？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 临近 OOM，如何获取详细内存分配信息，分析内存问题？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/14%20%e4%b8%b4%e8%bf%91%20OOM%ef%bc%8c%e5%a6%82%e4%bd%95%e8%8e%b7%e5%8f%96%e8%af%a6%e7%bb%86%e5%86%85%e5%ad%98%e5%88%86%e9%85%8d%e4%bf%a1%e6%81%af%ef%bc%8c%e5%88%86%e6%9e%90%e5%86%85%e5%ad%98%e9%97%ae%e9%a2%98%ef%bc%9f.md">14 临近 OOM，如何获取详细内存分配信息，分析内存问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 日志监控：怎样获取 App 中的全量日志？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/15%20%e6%97%a5%e5%bf%97%e7%9b%91%e6%8e%a7%ef%bc%9a%e6%80%8e%e6%a0%b7%e8%8e%b7%e5%8f%96%20App%20%e4%b8%ad%e7%9a%84%e5%85%a8%e9%87%8f%e6%97%a5%e5%bf%97%ef%bc%9f.md">15 日志监控：怎样获取 App 中的全量日志？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 性能监控：衡量 App 质量的那把尺.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/16%20%e6%80%a7%e8%83%bd%e7%9b%91%e6%8e%a7%ef%bc%9a%e8%a1%a1%e9%87%8f%20App%20%e8%b4%a8%e9%87%8f%e7%9a%84%e9%82%a3%e6%8a%8a%e5%b0%ba.md">16 性能监控：衡量 App 质量的那把尺.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 远超你想象的多线程的那些坑.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/17%20%e8%bf%9c%e8%b6%85%e4%bd%a0%e6%83%b3%e8%b1%a1%e7%9a%84%e5%a4%9a%e7%ba%bf%e7%a8%8b%e7%9a%84%e9%82%a3%e4%ba%9b%e5%9d%91.md">17 远超你想象的多线程的那些坑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 怎么减少 App 电量消耗？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/18%20%e6%80%8e%e4%b9%88%e5%87%8f%e5%b0%91%20App%20%e7%94%b5%e9%87%8f%e6%b6%88%e8%80%97%ef%bc%9f.md">18 怎么减少 App 电量消耗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 热点问题答疑（二）：基础模块问题答疑.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/19%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%9f%ba%e7%a1%80%e6%a8%a1%e5%9d%97%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91.md">19 热点问题答疑（二）：基础模块问题答疑.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 iOS开发的最佳学习路径是什么？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/20%20iOS%e5%bc%80%e5%8f%91%e7%9a%84%e6%9c%80%e4%bd%b3%e5%ad%a6%e4%b9%a0%e8%b7%af%e5%be%84%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">20 iOS开发的最佳学习路径是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 除了 Cocoa，iOS还可以用哪些 GUI 框架开发？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/21%20%e9%99%a4%e4%ba%86%20Cocoa%ef%bc%8ciOS%e8%bf%98%e5%8f%af%e4%bb%a5%e7%94%a8%e5%93%aa%e4%ba%9b%20GUI%20%e6%a1%86%e6%9e%b6%e5%bc%80%e5%8f%91%ef%bc%9f.md">21 除了 Cocoa，iOS还可以用哪些 GUI 框架开发？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 细说 iOS 响应式框架变迁，哪些思想可以为我所用？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/22%20%e7%bb%86%e8%af%b4%20iOS%20%e5%93%8d%e5%ba%94%e5%bc%8f%e6%a1%86%e6%9e%b6%e5%8f%98%e8%bf%81%ef%bc%8c%e5%93%aa%e4%ba%9b%e6%80%9d%e6%83%b3%e5%8f%af%e4%bb%a5%e4%b8%ba%e6%88%91%e6%89%80%e7%94%a8%ef%bc%9f.md">22 细说 iOS 响应式框架变迁，哪些思想可以为我所用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 如何构造酷炫的物理效果和过场动画效果？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/23%20%e5%a6%82%e4%bd%95%e6%9e%84%e9%80%a0%e9%85%b7%e7%82%ab%e7%9a%84%e7%89%a9%e7%90%86%e6%95%88%e6%9e%9c%e5%92%8c%e8%bf%87%e5%9c%ba%e5%8a%a8%e7%94%bb%e6%95%88%e6%9e%9c%ef%bc%9f.md">23 如何构造酷炫的物理效果和过场动画效果？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 A_B 测试：验证决策效果的利器.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/24%20A_B%20%e6%b5%8b%e8%af%95%ef%bc%9a%e9%aa%8c%e8%af%81%e5%86%b3%e7%ad%96%e6%95%88%e6%9e%9c%e7%9a%84%e5%88%a9%e5%99%a8.md">24 A_B 测试：验证决策效果的利器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 怎样构建底层的发布和订阅事件总线？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/25%20%e6%80%8e%e6%a0%b7%e6%9e%84%e5%bb%ba%e5%ba%95%e5%b1%82%e7%9a%84%e5%8f%91%e5%b8%83%e5%92%8c%e8%ae%a2%e9%98%85%e4%ba%8b%e4%bb%b6%e6%80%bb%e7%ba%bf%ef%bc%9f.md">25 怎样构建底层的发布和订阅事件总线？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 如何提高 JSON 解析的性能？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/26%20%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%20JSON%20%e8%a7%a3%e6%9e%90%e7%9a%84%e6%80%a7%e8%83%bd%ef%bc%9f.md">26 如何提高 JSON 解析的性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 如何用 Flexbox 思路开发？跟自动布局比，Flexbox 好在哪？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/27%20%e5%a6%82%e4%bd%95%e7%94%a8%20Flexbox%20%e6%80%9d%e8%b7%af%e5%bc%80%e5%8f%91%ef%bc%9f%e8%b7%9f%e8%87%aa%e5%8a%a8%e5%b8%83%e5%b1%80%e6%af%94%ef%bc%8cFlexbox%20%e5%a5%bd%e5%9c%a8%e5%93%aa%ef%bc%9f.md">27 如何用 Flexbox 思路开发？跟自动布局比，Flexbox 好在哪？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 怎么应对各种富文本表现需求？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/28%20%e6%80%8e%e4%b9%88%e5%ba%94%e5%af%b9%e5%90%84%e7%a7%8d%e5%af%8c%e6%96%87%e6%9c%ac%e8%a1%a8%e7%8e%b0%e9%9c%80%e6%b1%82%ef%bc%9f.md">28 怎么应对各种富文本表现需求？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/29%20%e5%a6%82%e4%bd%95%e5%9c%a8%20iOS%20%e4%b8%ad%e8%bf%9b%e8%a1%8c%e9%9d%a2%e5%90%91%e6%b5%8b%e8%af%95%e9%a9%b1%e5%8a%a8%e5%bc%80%e5%8f%91%e5%92%8c%e9%9d%a2%e5%90%91%e8%a1%8c%e4%b8%ba%e9%a9%b1%e5%8a%a8%e5%bc%80%e5%8f%91%ef%bc%9f.md">29 如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 如何制定一套适合自己团队的 iOS 编码规范？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/30%20%e5%a6%82%e4%bd%95%e5%88%b6%e5%ae%9a%e4%b8%80%e5%a5%97%e9%80%82%e5%90%88%e8%87%aa%e5%b7%b1%e5%9b%a2%e9%98%9f%e7%9a%84%20iOS%20%e7%bc%96%e7%a0%81%e8%a7%84%e8%8c%83%ef%bc%9f.md">30 如何制定一套适合自己团队的 iOS 编码规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 iOS 开发学习资料和书单推荐.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/31%20iOS%20%e5%bc%80%e5%8f%91%e5%ad%a6%e4%b9%a0%e8%b5%84%e6%96%99%e5%92%8c%e4%b9%a6%e5%8d%95%e6%8e%a8%e8%8d%90.md">31 iOS 开发学习资料和书单推荐.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 热点问题答疑（三）.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/32%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%88%e4%b8%89%ef%bc%89.md">32 热点问题答疑（三）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 iOS 系统内核 XNU：App 如何加载？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/33%20iOS%20%e7%b3%bb%e7%bb%9f%e5%86%85%e6%a0%b8%20XNU%ef%bc%9aApp%20%e5%a6%82%e4%bd%95%e5%8a%a0%e8%bd%bd%ef%bc%9f.md">33 iOS 系统内核 XNU：App 如何加载？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 iOS 黑魔法 Runtime Method Swizzling 背后的原理.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/34%20iOS%20%e9%bb%91%e9%ad%94%e6%b3%95%20Runtime%20Method%20Swizzling%20%e8%83%8c%e5%90%8e%e7%9a%84%e5%8e%9f%e7%90%86.md">34 iOS 黑魔法 Runtime Method Swizzling 背后的原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 libffi：动态调用和定义 C 函数.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/35%20libffi%ef%bc%9a%e5%8a%a8%e6%80%81%e8%b0%83%e7%94%a8%e5%92%8c%e5%ae%9a%e4%b9%89%20C%20%e5%87%bd%e6%95%b0.md">35 libffi：动态调用和定义 C 函数.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 iOS 是怎么管理内存的？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/36%20iOS%20%e6%98%af%e6%80%8e%e4%b9%88%e7%ae%a1%e7%90%86%e5%86%85%e5%ad%98%e7%9a%84%ef%bc%9f.md">36 iOS 是怎么管理内存的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 如何编写 Clang 插件？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/37%20%e5%a6%82%e4%bd%95%e7%bc%96%e5%86%99%20Clang%20%e6%8f%92%e4%bb%b6%ef%bc%9f.md">37 如何编写 Clang 插件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 热点问题答疑（四）.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/38%20%e7%83%ad%e7%82%b9%e9%97%ae%e9%a2%98%e7%ad%94%e7%96%91%ef%bc%88%e5%9b%9b%ef%bc%89.md">38 热点问题答疑（四）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 打通前端与原生的桥梁：JavaScriptCore 能干哪些事情？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/39%20%e6%89%93%e9%80%9a%e5%89%8d%e7%ab%af%e4%b8%8e%e5%8e%9f%e7%94%9f%e7%9a%84%e6%a1%a5%e6%a2%81%ef%bc%9aJavaScriptCore%20%e8%83%bd%e5%b9%b2%e5%93%aa%e4%ba%9b%e4%ba%8b%e6%83%85%ef%bc%9f.md">39 打通前端与原生的桥梁：JavaScriptCore 能干哪些事情？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 React Native、Flutter 等，这些跨端方案怎么选？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/40%20React%20Native%e3%80%81Flutter%20%e7%ad%89%ef%bc%8c%e8%bf%99%e4%ba%9b%e8%b7%a8%e7%ab%af%e6%96%b9%e6%a1%88%e6%80%8e%e4%b9%88%e9%80%89%ef%bc%9f.md">40 React Native、Flutter 等，这些跨端方案怎么选？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 原生布局转到前端布局，开发思路有哪些转变？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/41%20%e5%8e%9f%e7%94%9f%e5%b8%83%e5%b1%80%e8%bd%ac%e5%88%b0%e5%89%8d%e7%ab%af%e5%b8%83%e5%b1%80%ef%bc%8c%e5%bc%80%e5%8f%91%e6%80%9d%e8%b7%af%e6%9c%89%e5%93%aa%e4%ba%9b%e8%bd%ac%e5%8f%98%ef%bc%9f.md">41 原生布局转到前端布局，开发思路有哪些转变？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 iOS原生、大前端和Flutter分别是怎么渲染的？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/42%20iOS%e5%8e%9f%e7%94%9f%e3%80%81%e5%a4%a7%e5%89%8d%e7%ab%af%e5%92%8cFlutter%e5%88%86%e5%88%ab%e6%98%af%e6%80%8e%e4%b9%88%e6%b8%b2%e6%9f%93%e7%9a%84%ef%bc%9f.md">42 iOS原生、大前端和Flutter分别是怎么渲染的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 剖析使 App 具有动态化和热更新能力的方案.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/43%20%e5%89%96%e6%9e%90%e4%bd%bf%20App%20%e5%85%b7%e6%9c%89%e5%8a%a8%e6%80%81%e5%8c%96%e5%92%8c%e7%83%ad%e6%9b%b4%e6%96%b0%e8%83%bd%e5%8a%9b%e7%9a%84%e6%96%b9%e6%a1%88.md">43 剖析使 App 具有动态化和热更新能力的方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="用户故事 我是如何学习这个专栏的？.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/%e7%94%a8%e6%88%b7%e6%95%85%e4%ba%8b%20%e6%88%91%e6%98%af%e5%a6%82%e4%bd%95%e5%ad%a6%e4%b9%a0%e8%bf%99%e4%b8%aa%e4%b8%93%e6%a0%8f%e7%9a%84%ef%bc%9f.md">用户故事 我是如何学习这个专栏的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 慢几步，深几度.md" href="/%e4%b8%93%e6%a0%8f/iOS%e5%bc%80%e5%8f%91%e9%ab%98%e6%89%8b%e8%af%be/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%85%a2%e5%87%a0%e6%ad%a5%ef%bc%8c%e6%b7%b1%e5%87%a0%e5%ba%a6.md">结束语 慢几步，深几度.md</a>
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
                            <h1 id="title" data-id="29 如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发？" class="title">29 如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发？</h1>
                            <div><p>你好，我是戴铭。今天，我要和你分享的话题是，如何在 iOS 中进行面向测试驱动开发和面向行为驱动开发。</p>

<p>每当你编写完代码后，都会编译看看运行结果是否符合预期。如果这段代码的影响范围小，你很容易就能看出结果是否符合预期，而如果验证的结果是不符合预期，那么你就会检查刚才编写的代码是否有问题。</p>

<p>但是，如果这段代码的影响范围比较大，这时需要检查的地方就会非常多，相应地，人工检查的时间成本也会非常大。特别是团队成员多、工程代码量大时，判断这段代码的影响面都需要耗费很多时间。那么，每次编写完代码，先判断它的影响面，然后再手动编译进行检查的开发方式，效率就非常低了，会浪费大量时间。</p>

<p>虽说一般公司都会有专门的测试团队对产品进行大量测试，但是如果不能在开发阶段及时发现问题，当各团队代码集成到一起，把所有问题都堆积到测试阶段去发现、解决，就会浪费大量的沟通时间，不光是开发同学和测试同学之间的沟通时间，还有开发团队之间的沟通时间也会呈指数级增加。</p>

<p>那么，有没有什么好的开发方式，能够提高在编写代码后及时检验结果的效率呢？</p>

<p>所谓好的开发方式，就是开发、测试同步进行，尽早发现问题。从测试范围和开发模式的角度，我们还可以把这种开发模式细分出更多类型。</p>

<p><strong>从测试范围上来划分的话</strong>，软件测试可以分为单元测试、集成测试、系统测试。测试团队负责的是集成测试以及系统测试，而单元测试则是有开发者负责的。对于开发者来说，通过单元测试就可以有效提高编写代码后快速发现问题的效率。</p>

<p>概括来说，单元测试，也叫作模块测试，就是对单一的功能代码进行测试。这个功能代码，可能是一个类的方法，也可能是一个模块的某个函数。</p>

<p>单元测试会使用 Mock 方式模拟外部使用，通过编写的各种测试用例去检验代码的功能是否正常。一个系统都是由各个功能组合而成，功能模块划分得越小，功能职责就越清晰。清晰的功能职责可以确保单个功能的测试不会出现问题，是单元测试的基础。</p>

<p><strong>从开发模式划分的话，</strong>开发方式可以分为 TDD（Test-driven development，面向测试驱动开发）和 BDD（Behavior-driven development ，面向行为驱动开发）。</p>

<ul>
<li>TDD 的开发思路是，先编写测试用例，然后在不考虑代码优化的情况下快速编写功能实现代码，等功能开发完成后，在测试用例的保障下，再进行代码重构，以提高代码质量。</li>
<li>BDD 是 TDD 的进化，基于行为进行功能测试，使用 DSL（Domain Specific Language，领域特定语言）来描述测试用例，让测试用例看起来和文档一样，更易读、更好维护。</li>
</ul>

<p>TDD 编写的测试用例主要针对的是开发中最小单元进行测试，适合单元测试。而 BDD 的测试用例是对行为的描述，测试范围可以更大一些，在集成测试和系统测试时都可以使用。同时，不仅开发者可以使用BDD的测试用例高效地发现问题，测试团队也能够很容易参与编写。这，都得益于 BDD 可以使用易于编写行为功能测试的 DSL 语言。</p>

<p>接下来，我就和你详细聊聊 TDD 和 BDD。</p>

<h2 id="tdd">TDD</h2>

<p>我刚刚也已经提到了，TDD在确定功能需求后，首先就会开始编写测试用例，用来检验每次的代码更新，能够让我们更快地发现问题，并能保正不会漏掉问题。其实，这就是通过测试用例来推动开发。</p>

<p>在思想上，和拿到功能需求后直接开发功能的区别是，TDD会先考虑如何对功能进行测试，然后再去考虑如何编写代码，这就给优化代码提供了更多的时间和空间，即使几个版本过后再来优化，只要能够通过先前写好的测试用例，就能够保证代码质量。</p>

<p>所以说，TDD 非常适合快速迭代的节奏，先尽快实现功能，然后再进行重构和优化。如果我们不使用 TDD 来进行快速迭代开发，虽然在最开始的时候开发效率会比 TDD 高，但是过几个版本再进行功能更新时，就需要在功能验证上花费大量的时间，反而得不偿失。</p>

<p>其实，TDD 这种开发模式和画漫画的工作方式非常类似：草稿就类似 TDD 中的测试用例，漫画家先画草稿，细节由漫画家和助手一起完成，无论助手怎么换，有了草稿的保障，内容都不会有偏差。分镜的草稿没有细节，人物眼睛、鼻子都可能没有，场景也只需要几条透视线就可以。虽然没有细节，但是草稿基本就确定了漫画完成后要表达的所有内容。</p>

<h2 id="bdd">BDD</h2>

<p>相比 TDD，BDD更关注的是行为方式的设计，通过对行为的描述来验证功能的可用性。行为描述使用的 DSL，规范、标准而且可读性高，可以当作文档来使用。</p>

<p>BDD 的 Objective-C 框架有 <a href="https://github.com/kiwi-bdd/Kiwi" target="_blank">Kiwi</a>、<a href="https://github.com/specta/specta" target="_blank">Specta</a>、<a href="https://github.com/specta/expecta" target="_blank">Expecta</a>等，Swift 框架有 <a href="https://github.com/Quick/Quick" target="_blank">Quick</a>。</p>

<p>Kiwi框架不光有 Specta 的 DSL 模式，Expecta框架的期望语法，还有 Mocks 和 Stubs 这样的模拟存根能力。所以接下来，我就跟你说说这个iOS中非常有名并且好用的BDD框架，以及怎么用它来进行 BDD 开发。</p>

<h2 id="kiwi">Kiwi</h2>

<p>将Kiwi集成到你的App里，只需要在 Podfile 里添加 pod ‘Kiwi’ 即可。下面这段代码，是 Kiwi 的使用示例：</p>

<pre><code>// describe 表示要测试的对象
describe(@&quot;RSSListViewController&quot;, ^{
    // context 表示的是不同场景下的行为
    context(@&quot;when get RSS data&quot;, ^{
        // 同一个 context 下每个 it 调用之前会调用一次 beforeEach
        beforeEach(^{
            id dataStore = [DataStore new];
        });


        // it 表示测试内容，一个 context 可以有多个 it
        it(@&quot;load data&quot;, ^{
            // Kiwi 使用链式调用，should 表示一个期待，用来验证对象行为是否满足期望
            [[theValue(dataStore.count) shouldNot] beNil];
        });
    });
});
</code></pre>

<p>上面这代码描述的是在 RSS 列表页面，当获取 RSS 数据时去读取数据这个行为的测试用例。这段测试用例代码，包含了 Kiwi 的基本元素，也就是describe、context、it。这些元素间的关系可以表述为：</p>

<ul>
<li>describe 表示要测试的对象，context 表示的是不同场景下的行为，一个 describe 里可以包含多个 context。</li>
<li>it表示的是需要测试的内容，同一个场景下的行为会有多个需要测试的内容，也就是说一个 context 下可以有多个 it。</li>
</ul>

<p>测试内容使用的是 Kiwi 的 DSL 语法，采用的是链式调用。上面示例代码中 shouldNot 是期望语法，期望是用来验证对象行为是否满足期望。</p>

<p>期望语法可以是期望数值和数字，也可以是期望字符串的匹配，比如：</p>

<pre><code>[[string should] containString:@&quot;rss&quot;];
</code></pre>

<p>should containString 语法表示的是，期望 string 包含了 rss 字符串。Kiwi 里的期望语法非常丰富，还有正则表达式匹配、数量变化、对象测试、集合、交互和消息、通知、异步调用、异常等。完整的期望语法描述，你可以查看Wiki的 <a href="https://github.com/allending/Kiwi/wiki/Expectations" target="_blank">Expectations 部分</a>。</p>

<p>除了期望语法外，Kiwi 还支持模拟对象和存根语法。</p>

<p>模拟对象能够降低对象之间的依赖，可以模拟难以出现的情况。模拟对象包含了模拟 Null 对象、模拟类的实例、模拟协议的实例等。存根可以返回指定选择器或消息模式的请求，可以存根对象和模拟对象。</p>

<p>模拟对象和存根的详细语法定义，你可以查看Wiki 的 <a href="https://github.com/allending/Kiwi/wiki/Mocks-and-Stubs" target="_blank">Mocks and Stubs 部分</a>。</p>

<h2 id="小结">小结</h2>

<p>按照 TDD 和 BDD 方式开发，有助于更好地进行模块化设计，划清模块边界，让代码更容易维护。TDD 在测试用例的保障下更容易进行代码重构优化，减少 debug 时间。而使用 BDD 编写的测试用例，则更是好的文档，可读性非常强。通过这些测试用例，在修改代码时，我们能够更方便地了解开发 App 的工作状态。同时，修改完代码后还能够快速全面地测试验证问题。</p>

<p>无论是 TDD 还是 BDD，开发中对于每个实现的方法都要编写测试用例，而且要注意先编写测试用例代码，再编写方法实现代码。测试用例需要考虑到各种异常条件，以及输入输出的边界。编写完测试用例还需要检查如果输入为错时，测试用例是否会显示为错。</p>

<p>最后需要强调一点，好的模块化架构和 TDD 、BDD 是相辅相成的。TDD 和 BDD 开发方式会让你的代码更加模块化，而模块化的架构更容易使用 TDD 和 BDD 的方式进行开发。</p>

<p>在团队中推行 TDD 和 BDD 的最大困难，就是业务迭代太快时，没有时间去写测试用例。我的建议是，优先对基础能力的功能开发使用 TDD 和 BDD，保证了基础能力的稳定，业务怎么变，底子还都是稳固的；当有了业务迭代、有了间隙时，再考虑在核心业务上采用 BDD，最大程度的保证 App 核心功能的稳定。</p>

<h2 id="课后作业">课后作业</h2>

<p>今天我跟你聊了很多 TDD 和 BDD 的优点，但是很多团队并没有使用这样的开发方式，你觉得这其中的原因是什么呢？</p>

<p>感谢你的收听，欢迎你在评论区给我留言分享你的观点，也欢迎把它分享给更多的朋友一起阅读。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#8be7e7e7b2bfbababbbccbece6eae2e7a5e8e4e6" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12fe679c5876e1',t:'MTczNDA2MTg2Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
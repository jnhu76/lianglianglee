<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;性能监控：衡量&#32;App&#32;质量的那把尺>
        <link rel="icon" href="/static/favicon.png">
        <title>16 性能监控：衡量 App 质量的那把尺 </title>
        
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
                            <h1 id="title" data-id="16 性能监控：衡量 App 质量的那把尺" class="title">16 性能监控：衡量 App 质量的那把尺</h1>
                            <div><p>你好，我是戴铭。</p>

<p>通常情况下，App 的性能问题虽然不会导致 App不可用，但依然会影响到用户体验。如果这个性能问题不断累积，达到临界点以后，问题就会爆发出来。这时，影响到的就不仅仅是用户了，还有负责App开发的你。</p>

<p>为了能够主动、高效地发现性能问题，避免App质量进入无人监管的失控状态，我们就需要对App的性能进行监控。目前，对App的性能监控，主要是从线下和线上两个维度展开。</p>

<p>今天这篇文章，我就从这两个方面来和你聊聊如何做性能监控这个话题。接下来，我们就先看看苹果官方的线下性能监控王牌 Instruments。</p>

<h2 id="instruments">Instruments</h2>

<p>关于线下性能监控，苹果公司官方就有一个性能监控工具Instruments。它是一款被集成在 Xcode 里，专门用来在线下进行性能分析的工具。</p>

<p>Instruments的功能非常强大，比如说Energy Log就是用来监控耗电量的，Leaks就是专门用来监控内存泄露问题的，Network就是用来专门检查网络情况的，Time Profiler就是通过时间采样来分析页面卡顿问题的。</p>

<p>如下图所示，就是Instruments的各种性能检测工具。</p>

<p><img src="assets/f11072c01a594c82b31f49526eee04c9.jpg" alt="" /></p>

<p>图1 Instruments 提供的各种性能检测工具</p>

<p>除了对各种性能问题进行监控外，<strong>最新版本的Instruments 10还有以下两大优势</strong>：</p>

<ol>
<li><p>Instruments基于os_signpost 架构，可以支持所有平台。</p></li>

<li><p>Instruments由于标准界面（Standard UI）和分析核心（Analysis Core）技术，使得我们可以非常方便地进行自定义性能监测工具的开发。当你想要给Instruments内置的工具换个交互界面，或者新创建一个工具的时候，都可以通过自定义工具这个功能来实现。</p></li>
</ol>

<p>其实，Instruments的这些优势也不是与生俱来的，都是伴随着移动开发技术的发展而演进来的。就比如说自定义工具的功能吧，这是因为App的规模越来越大，往往还涉及到多个团队合作开发、集成多个公司SDK的情况，所以我们就需要以黑盒的方式来进行性能监控。这样的需求，也就迫使苹果公司要不断地增强Instruments的功能。</p>

<p>从整体架构来看，Instruments 包括Standard UI 和 Analysis Core 两个组件，它的所有工具都是基于这两个组件开发的。而且，你如果要开发自定义的性能分析工具的话，完全基于这两个组件就可以实现。</p>

<p><strong>开发一款自定义Instruments工具</strong>，主要包括以下这几个步骤：</p>

<ol>
<li><p>在Xcode中，点击File &gt; New &gt; Project；</p></li>

<li><p>在弹出的Project模板选择界面，将其设置为macOS；</p></li>

<li><p>选择 Instruments Package，点击后即可开始自定义工具的开发了。如下图所示。</p></li>
</ol>

<p><img src="assets/2bafcc62cf104d3eaa99fadc3f86b95b.jpg" alt="" /></p>

<p>图2 开发自定义Instrument工具</p>

<p>经过上面的三步之后，会在新创建的工程里面生成一个.instrpkg 文件，接下来的开发过程主要就是对这个文件的配置工作了。这些配置工作中最主要的是要完成Standard UI 和 Analysis Core 的配置。</p>

<p>上面这些内容，就是你在开发一个自定义Instruments工具时，需要完成的编码工作了。可以看到，Instruments 10版本的自定义工具开发还是比较简单的。与此同时，苹果公司还提供了大量的代码片段，帮助你进行个性化的配置。你可以<a href="https://help.apple.com/instruments/developer/mac/current/" target="_blank">点击这个链接</a>，查看官方指南中的详细教程。</p>

<p>如果你想要更好地进行个性化定制，就还需要再了解Instruments收集和处理数据的机制，也就是<strong>分析核心（Analysis Core ）的工作原理</strong>。</p>

<p>Analysis Core收集和处理数据的过程，可以大致分为以下这三步：</p>

<ol>
<li><p>处理我们配置好的各种数据表，并申请存储空间 store；</p></li>

<li><p>store去找数据提供者，如果不能直接找到，就会通过 Modeler 接收其他store 的输入信号进行合成；</p></li>

<li><p>store 获得数据源后，会进行 Binding Solution 工作来优化数据处理过程。</p></li>
</ol>

<p>这里需要强调的是，在我们通过store找到的这些数据提供者中，对开发者来说最重要的就是 os_signpost。os_signpost 的主要作用，是让你可以在程序中通过编写代码来获取数据。你可以在工程中的任何地方通过 os_signpost API ，将需要的数据提供给 Analysis Core。</p>

<p>苹果公司在 WWDC 2018 Session 410 <a href="https://developer.apple.com/videos/play/wwdc2018/410" target="_blank">Creating Custom Instruments</a> 里提供了一个范例：通过 os_signpost API 将图片下载的数据提供给 Analysis Core 进行监控观察。这个示例在 App 的代码如下所示：</p>

<pre><code>os_signpost(.begin, log: parsinglog, name:&quot;Parsing&quot;, &quot;Parsing started SIZE:%ld&quot;, data.count)
// Decode the JSON we just downloaded
let result = try jsonDecoder.decode(Trail.self, from: data)
os_signpost(.end, log: parsingLog, name:&quot;Parsing&quot;, &quot;Parsing finished&quot;)
</code></pre>

<p>需要注意的是，上面代码中，os_signpost 的 begin 和 end 需要成对出现。</p>

<p>上面这段代码就是使用 os_signpost 的 API 获取了程序里的数据。接下来，我们再看看 Instruments 是如何通过配置数据表来使用这些数据的。配置的数据表的 XML 设计如下所示：</p>

<pre><code>&lt;os-signpost-interval-schema&gt;
&lt;id&gt;json-parse&lt;/id&gt;
&lt;title&gt;Image Download&lt;/title&gt;
&lt;subsystem&gt;&quot;com.apple.trailblazer&lt;/subsystem&gt;
&lt;category&gt;&quot;Networking&lt;/category&gt;
&lt;name&gt;&quot;Parsing&quot;&lt;/name&gt;
&lt;start-pattern&gt;
&lt;message&gt;&quot;Parsing started SIZE:&quot; ?data-size&lt;/message&gt; 
&lt;/start-pattern&gt;
&lt;column&gt;
&lt;mnemonic&gt;data-size&lt;/mnemonic&gt;
&lt;title&gt;JSON Data Size&lt;/title&gt;
&lt;type&gt;size-in-bytes&lt;/type&gt;
&lt;expression&gt;?data-size&lt;/expression&gt;
&lt;/column&gt;
&lt;/os-signpost-interval-schema&gt;
</code></pre>

<p>这里，我们配置数据表是要对数据输出进行可视化配置，从而可以将代码中的数据展示出来。如下图所示，就是对下载图片大小监控的效果。</p>

<p><img src="assets/0d037395a6d84ec5977ba0a7fa05d272.jpg" alt="" /></p>

<p>图3 对下载图片大小的监控</p>

<p>通过上面的分析我们可以看到，Instruments 10通过提供 os_signpost API 的方式使得开发者监控自定义的性能指标时更方便，从而解决了在此之前只能通过重新建设工具来完成的问题。并且，Instruments通过 XML 标准数据接口解耦展示和数据分析的思路，也非常值得我们借鉴和学习。</p>

<p>在线下性能监控中，Instruments可以说是王者，但却对线上监控无能为力。那么，对于线上的性能监控，我们应该怎么实现呢？</p>

<h2 id="线上性能监控">线上性能监控</h2>

<p>对于线上性能监控，我们需要先明白两个原则：</p>

<ol>
<li><p>监控代码不要侵入到业务代码中；</p></li>

<li><p>采用性能消耗最小的监控方案。</p></li>
</ol>

<p>线上性能监控，主要集中在CPU使用率、FPS的帧率和内存这三个方面。接下来，我们就分别从这三个方面展开讨论吧。</p>

<h3 id="cpu使用率的线上监控方法">CPU使用率的线上监控方法</h3>

<p>App作为进程运行起来后会有多个线程，每个线程对CPU 的使用率不同。各个线程对CPU使用率的总和，就是当前App对CPU 的使用率。明白了这一点以后，我们也就摸清楚了对CPU使用率进行线上监控的思路。</p>

<p>在iOS系统中，你可以在 usr/include/mach/thread_info.h 里看到线程基本信息的结构体，其中的cpu_usage 就是 CPU使用率。结构体的完整代码如下所示：</p>

<pre><code>struct thread_basic_info {
  time_value_t    user_time;     // 用户运行时长
  time_value_t    system_time;   // 系统运行时长
  integer_t       cpu_usage;     // CPU 使用率
  policy_t        policy;        // 调度策略
  integer_t       run_state;     // 运行状态
  integer_t       flags;         // 各种标记
  integer_t       suspend_count; // 暂停线程的计数
  integer_t       sleep_time;    // 休眠的时间
};
</code></pre>

<p>因为每个线程都会有这个 thread_basic_info 结构体，所以接下来的事情就好办了，你只需要定时（比如，将定时间隔设置为2s）去遍历每个线程，累加每个线程的 cpu_usage 字段的值，就能够得到当前App所在进程的 CPU 使用率了。实现代码如下：</p>

<pre><code>+ (integer_t)cpuUsage {
    thread_act_array_t threads; //int 组成的数组比如 thread[1] = 5635
    mach_msg_type_number_t threadCount = 0; //mach_msg_type_number_t 是 int 类型
    const task_t thisTask = mach_task_self();
    //根据当前 task 获取所有线程
    kern_return_t kr = task_threads(thisTask, &amp;threads, &amp;threadCount);
    
    if (kr != KERN_SUCCESS) {
        return 0;
    }
    
    integer_t cpuUsage = 0;
    // 遍历所有线程
    for (int i = 0; i &lt; threadCount; i++) {
        
        thread_info_data_t threadInfo;
        thread_basic_info_t threadBaseInfo;
        mach_msg_type_number_t threadInfoCount = THREAD_INFO_MAX;
        
        if (thread_info((thread_act_t)threads[i], THREAD_BASIC_INFO, (thread_info_t)threadInfo, &amp;threadInfoCount) == KERN_SUCCESS) {
            // 获取 CPU 使用率
            threadBaseInfo = (thread_basic_info_t)threadInfo;
            if (!(threadBaseInfo-&gt;flags &amp; TH_FLAGS_IDLE)) {
                cpuUsage += threadBaseInfo-&gt;cpu_usage;
            }
        }
    }
    assert(vm_deallocate(mach_task_self(), (vm_address_t)threads, threadCount * sizeof(thread_t)) == KERN_SUCCESS);
    return cpuUsage;
}
</code></pre>

<p>在上面这段代码中，task_threads 方法能够取到当前进程中的线程总数 threadCount 和所有线程的数组 threads。</p>

<p>接下来，我们就可以通过遍历这个数组来获取单个线程的基本信息。其中，线程基本信息的结构体是 thread_basic_info_t，这个结构体里就包含了我们需要的 CPU 使用率的字段 cpu_usage。然后，我们累加这个字段就能够获取到当前的整体 CPU 使用率。</p>

<p>到此，我们就实现了对CPU使用率的线上监控。接下来，我们再看看对FPS的线上监控方法吧。</p>

<h3 id="fps-线上监控方法">FPS 线上监控方法</h3>

<p>FPS 是指图像连续在显示设备上出现的频率。FPS低，表示App不够流畅，还需要进行优化。</p>

<p>但是，和前面对CPU使用率和内存使用量的监控不同，iOS系统中没有一个专门的结构体，用来记录与FPS相关的数据。但是，对FPS的监控也可以比较简单的实现：通过注册 CADisplayLink 得到屏幕的同步刷新率，记录每次刷新时间，然后就可以得到 FPS。具体的实现代码如下：</p>

<pre><code>- (void)start {
    self.dLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(fpsCount:)];
    [self.dLink addToRunLoop:[NSRunLoop mainRunLoop] forMode:NSRunLoopCommonModes];
}

// 方法执行帧率和屏幕刷新率保持一致
- (void)fpsCount:(CADisplayLink *)displayLink {
    if (lastTimeStamp == 0) {
        lastTimeStamp = self.dLink.timestamp;
    } else {
        total++;
        // 开始渲染时间与上次渲染时间差值
        NSTimeInterval useTime = self.dLink.timestamp - lastTimeStamp;
        if (useTime &lt; 1) return;
        lastTimeStamp = self.dLink.timestamp;
        // fps 计算
        fps = total / useTime; 
        total = 0;
    }
}
</code></pre>

<h3 id="内存使用量的线上监控方法">内存使用量的线上监控方法</h3>

<p>通常情况下，我们在获取 iOS 应用内存使用量时，都是使用task_basic_info 里的 resident_size 字段信息。但是，我们发现这样获得的内存使用量和 Instruments 里看到的相差很大。后来，在 2018 WWDC Session 416 <a href="https://developer.apple.com/videos/play/wwdc2018/416/" target="_blank">iOS Memory Deep Dive</a>中，苹果公司介绍说 phys_footprint 才是实际使用的物理内存。</p>

<p>内存信息存在 task_info.h （完整路径 usr/include/mach/task.info.h）文件的 task_vm_info 结构体中，其中phys_footprint 就是物理内存的使用，而不是驻留内存 resident_size。结构体里和内存相关的代码如下：</p>

<pre><code>struct task_vm_info {
  mach_vm_size_t  virtual_size;       // 虚拟内存大小
  integer_t region_count;             // 内存区域的数量
  integer_t page_size;
  mach_vm_size_t  resident_size;      // 驻留内存大小
  mach_vm_size_t  resident_size_peak; // 驻留内存峰值

  ...

  /* added for rev1 */
  mach_vm_size_t  phys_footprint;     // 物理内存

  ...
</code></pre>

<p>OK，类似于对CPU使用率的监控，我们只要从这个结构体里取出phys_footprint 字段的值，就能够监控到实际物理内存的使用情况了。具体实现代码如下：</p>

<pre><code>uint64_t memoryUsage() {
    task_vm_info_data_t vmInfo;
    mach_msg_type_number_t count = TASK_VM_INFO_COUNT;
    kern_return_t result = task_info(mach_task_self(), TASK_VM_INFO, (task_info_t) &amp;vmInfo, &amp;count);
    if (result != KERN_SUCCESS)
        return 0;
    return vmInfo.phys_footprint;
}
</code></pre>

<p>从以上三个线上性能监控方案可以看出，它们的代码和业务逻辑是完全解耦的，监控时基本都是直接获取系统本身提供的数据，没有额外的计算量，因此对 App 本身的性能影响也非常小，满足了我们要考虑的两个原则。</p>

<h2 id="小结">小结</h2>

<p>在今天这篇文章中，我和你分享了如何通过线下和线上监控，去掌控App的性能。</p>

<p>关于线下的性能监控，我们可以使用苹果官方的Instruments 去解决性能监控的问题。同时，我还和你分享了如何使用 Instruments 的 os_signpost API 来完成自定义的性能数据监控工具开发。</p>

<p>关于线上的性能监控，我们需要在不影响性能的前提下，去监控线上的性能问题。在这一部分内容中，我主要和你介绍了对CPU使用率、内存使用量和FPS的线上监控方案。</p>

<p>最后，我还要再和你提一个建议。作为一名 iOS 开发者，与其一起开始到处去寻找各种解决方案，不如先摸透苹果公司自己的库和工具，这里面的设计思想和演进包含有大量可以吸取和学习的知识。掌握好了这些知识，你也就能够开发出适合自己团队的工具了。这，也正是我没有在这篇文章中和你介绍第三方线上性能监控工具的原因。</p>

<h2 id="课后小作业">课后小作业</h2>

<p>Instruments 可以自定义性能数据的监控，那么接下来就请你看下，你现在工程中有哪些数据是需要监控的，然后新建一个自定义 Instruments 工具将其监控起来吧。</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#f5999999ccc1c4c4c5c2b59298949c99db969a98" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12fbc99fbf76e1',t:'MTczNDA2MTc1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
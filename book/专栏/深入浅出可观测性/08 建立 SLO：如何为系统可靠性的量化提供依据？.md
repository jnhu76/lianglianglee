<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=08&#32;建立&#32;SLO：如何为系统可靠性的量化提供依据？>
        <link rel="icon" href="/static/favicon.png">
        <title>08 建立 SLO：如何为系统可靠性的量化提供依据？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 可观测性，让开发和维护系统的你每晚都能睡个好觉！.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%ef%bc%8c%e8%ae%a9%e5%bc%80%e5%8f%91%e5%92%8c%e7%bb%b4%e6%8a%a4%e7%b3%bb%e7%bb%9f%e7%9a%84%e4%bd%a0%e6%af%8f%e6%99%9a%e9%83%bd%e8%83%bd%e7%9d%a1%e4%b8%aa%e5%a5%bd%e8%a7%89%ef%bc%81.md">00 开篇词 可观测性，让开发和维护系统的你每晚都能睡个好觉！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 以史鉴今：监控是如何一步步发展而来的？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/01%20%e4%bb%a5%e5%8f%b2%e9%89%b4%e4%bb%8a%ef%bc%9a%e7%9b%91%e6%8e%a7%e6%98%af%e5%a6%82%e4%bd%95%e4%b8%80%e6%ad%a5%e6%ad%a5%e5%8f%91%e5%b1%95%e8%80%8c%e6%9d%a5%e7%9a%84%ef%bc%9f.md">01 以史鉴今：监控是如何一步步发展而来的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 基本概念：指标&#43;日志&#43;链路追踪=可观测性？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/02%20%e5%9f%ba%e6%9c%ac%e6%a6%82%e5%bf%b5%ef%bc%9a%e6%8c%87%e6%a0%87&#43;%e6%97%a5%e5%bf%97&#43;%e9%93%be%e8%b7%af%e8%bf%bd%e8%b8%aa=%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%ef%bc%9f.md">02 基本概念：指标&#43;日志&#43;链路追踪=可观测性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 相互对比：可观测性和传统监控有什么区别？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/03%20%e7%9b%b8%e4%ba%92%e5%af%b9%e6%af%94%ef%bc%9a%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%e5%92%8c%e4%bc%a0%e7%bb%9f%e7%9b%91%e6%8e%a7%e6%9c%89%e4%bb%80%e4%b9%88%e5%8c%ba%e5%88%ab%ef%bc%9f.md">03 相互对比：可观测性和传统监控有什么区别？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 OpenTelemetry：如何利用OpenTelemetry采集可观测数据？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/04%20OpenTelemetry%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8OpenTelemetry%e9%87%87%e9%9b%86%e5%8f%af%e8%a7%82%e6%b5%8b%e6%95%b0%e6%8d%ae%ef%bc%9f.md">04 OpenTelemetry：如何利用OpenTelemetry采集可观测数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 构建可观测平台：应该购买商业产品还是自己搭建？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/05%20%e6%9e%84%e5%bb%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e5%b9%b3%e5%8f%b0%ef%bc%9a%e5%ba%94%e8%af%a5%e8%b4%ad%e4%b9%b0%e5%95%86%e4%b8%9a%e4%ba%a7%e5%93%81%e8%bf%98%e6%98%af%e8%87%aa%e5%b7%b1%e6%90%ad%e5%bb%ba%ef%bc%9f.md">05 构建可观测平台：应该购买商业产品还是自己搭建？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 团队合作：不同团队如何高效共建可观测性？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/06%20%e5%9b%a2%e9%98%9f%e5%90%88%e4%bd%9c%ef%bc%9a%e4%b8%8d%e5%90%8c%e5%9b%a2%e9%98%9f%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%85%b1%e5%bb%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%ef%bc%9f.md">06 团队合作：不同团队如何高效共建可观测性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 软件开发：如何利用可观测性进行驱动？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/07%20%e8%bd%af%e4%bb%b6%e5%bc%80%e5%8f%91%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%a9%e7%94%a8%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%e8%bf%9b%e8%a1%8c%e9%a9%b1%e5%8a%a8%ef%bc%9f.md">07 软件开发：如何利用可观测性进行驱动？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 建立 SLO：如何为系统可靠性的量化提供依据？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/08%20%e5%bb%ba%e7%ab%8b%20SLO%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%ba%e7%b3%bb%e7%bb%9f%e5%8f%af%e9%9d%a0%e6%80%a7%e7%9a%84%e9%87%8f%e5%8c%96%e6%8f%90%e4%be%9b%e4%be%9d%e6%8d%ae%ef%bc%9f.md">08 建立 SLO：如何为系统可靠性的量化提供依据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 跟踪 SLO：如何基于错误预算采取行动？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/09%20%e8%b7%9f%e8%b8%aa%20SLO%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%e9%94%99%e8%af%af%e9%a2%84%e7%ae%97%e9%87%87%e5%8f%96%e8%a1%8c%e5%8a%a8%ef%bc%9f.md">09 跟踪 SLO：如何基于错误预算采取行动？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 实战 1：从 0 到 1 建立系统应用的可观测性.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/10%20%e5%ae%9e%e6%88%98%201%ef%bc%9a%e4%bb%8e%200%20%e5%88%b0%201%20%e5%bb%ba%e7%ab%8b%e7%b3%bb%e7%bb%9f%e5%ba%94%e7%94%a8%e7%9a%84%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7.md">10 实战 1：从 0 到 1 建立系统应用的可观测性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 实战 2：基于可观测性数据进行问题分析和根因定位.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/11%20%e5%ae%9e%e6%88%98%202%ef%bc%9a%e5%9f%ba%e4%ba%8e%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%e6%95%b0%e6%8d%ae%e8%bf%9b%e8%a1%8c%e9%97%ae%e9%a2%98%e5%88%86%e6%9e%90%e5%92%8c%e6%a0%b9%e5%9b%a0%e5%ae%9a%e4%bd%8d.md">11 实战 2：基于可观测性数据进行问题分析和根因定位.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 实战 3：构建基于 Kubernetes 的系统可观测性.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/12%20%e5%ae%9e%e6%88%98%203%ef%bc%9a%e6%9e%84%e5%bb%ba%e5%9f%ba%e4%ba%8e%20Kubernetes%20%e7%9a%84%e7%b3%bb%e7%bb%9f%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7.md">12 实战 3：构建基于 Kubernetes 的系统可观测性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 实战 4：建立软件持续集成和发布的可观测性.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/13%20%e5%ae%9e%e6%88%98%204%ef%bc%9a%e5%bb%ba%e7%ab%8b%e8%bd%af%e4%bb%b6%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%e5%92%8c%e5%8f%91%e5%b8%83%e7%9a%84%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7.md">13 实战 4：建立软件持续集成和发布的可观测性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 文化建设：如何构建可观测性的文化和框架_.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/14%20%e6%96%87%e5%8c%96%e5%bb%ba%e8%ae%be%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%e7%9a%84%e6%96%87%e5%8c%96%e5%92%8c%e6%a1%86%e6%9e%b6_.md">14 文化建设：如何构建可观测性的文化和框架_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 展望未来：可观测性在未来的发展趋势是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/15%20%e5%b1%95%e6%9c%9b%e6%9c%aa%e6%9d%a5%ef%bc%9a%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7%e5%9c%a8%e6%9c%aa%e6%9d%a5%e7%9a%84%e5%8f%91%e5%b1%95%e8%b6%8b%e5%8a%bf%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">15 展望未来：可观测性在未来的发展趋势是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 在实践中，用科学的方法分析和解决问题.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e6%b5%85%e5%87%ba%e5%8f%af%e8%a7%82%e6%b5%8b%e6%80%a7/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%9c%a8%e5%ae%9e%e8%b7%b5%e4%b8%ad%ef%bc%8c%e7%94%a8%e7%a7%91%e5%ad%a6%e7%9a%84%e6%96%b9%e6%b3%95%e5%88%86%e6%9e%90%e5%92%8c%e8%a7%a3%e5%86%b3%e9%97%ae%e9%a2%98.md">结束语 在实践中，用科学的方法分析和解决问题.md</a>
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
                            <h1 id="title" data-id="08 建立 SLO：如何为系统可靠性的量化提供依据？" class="title">08 建立 SLO：如何为系统可靠性的量化提供依据？</h1>
                            <div><p>你好，我是翁一磊。</p>

<p>在前面两节课，相信你已经意识到，建立可观测性需要开发、测试、运维等团队共同的努力，而大家一起努力的目标就是维护好系统可靠性。</p>

<p>这时候你可能会有一个疑问，系统的可靠性究竟应该如何衡量呢？在这一讲，我就来为你详细介绍一下服务水平目标，也就是 Service Level Objective（SLO）。通过设定具体且可衡量的可靠性目标，能够帮助企业在产品开发迭代和稳定性保障工作之间取得适当的平衡，带来积极的最终用户体验。</p>

<h2 id="什么是-slo">什么是 SLO？</h2>

<p>说到 SLO，相信你一定会联想到另一个英文缩写 SLA。<strong>SLA 即 Service Level Agreement</strong>，也就是服务等级协议，它指的是系统服务提供者（Provider）对客户（Customer）的一个服务承诺。</p>

<p>在移动互联网时代，我们对基于互联网应用的需求日益旺盛（电商、社交网络、游戏、云服务商、SaaS…），任何一个互联网业务应用（也就是这个时代的系统服务提供者）出现故障，都会对用户乃至于整个社会产生巨大的影响，因此服务提供者需要明确能够提供的服务保障。</p>

<p>而 SLO 就是 SLA 的具体目标管理办法，它由一系列相关的指标 <strong>SLI （Service Level Indicator）</strong>来进行衡量。虽然我们中文里也常提到指标，但 SLI 和我们之前讨论的 Metric（指标）有所不同：<strong>不是所有的 Metric 都是 SLI，SLI 应该更靠近使用产品和服务的最终用户，用于衡量提供给最终用户的服务水平，具体包括可用性、响应时间等等。</strong></p>

<p>那具体什么是 SLI 呢？</p>

<p>举个例子来说，如果你负责的是一个电子书刊或杂志的应用程序，最终用户进行“订阅”的这个操作对业务收入来说就至关重要。那么，针对这个订阅服务的 API 可用性，就应该作为一个 SLI，因为它关系到用户能否正常完成订阅的操作。如果一次请求返回的 HTTP 状态代码为 200，则可以视为一次成功的请求；但是如果不是 200，比如出现 5xx 的错误，我们就认为这次请求是失败的。在这节课的后半部分我们还会进一步介绍如何选择合适的 SLI。</p>

<p>有了 SLI，接下来，我们就可以检测在每个检测周期内各个 SLI 是否满足要求，从而计算整体的 SLO 情况了。<strong>SLO 具体来说，指的是在一个时间窗口内，各项 SLI 预期的累计成功百分比。这个时间窗口可以根据业务的需要来定义，一般来说为 30 天。</strong></p>

<p>举个例子，过去 30 天（总计 43200 分钟），如果发生异常的时间为 2 分钟，则 SLO 的状态为 （43200 - 2）/ 43200 * 100% = 99.995%。这里有一个对应的概念叫做<strong>错误预算（Error Budget），它指的是初始状态时 100% 可靠性和 SLO 目标之间的差额。</strong></p>

<p>每当和 SLO 相关联的 SLI 没有满足要求，我们就要扣除一部分错误预算，那剩余的错误预算就是当前 SLO 状态和 SLO 目标之间的差额；等到错误预算不足时，这个周期内的 SLO 就达不到目标了。</p>

<p>在这个例子中，如果 SLO 的目标是 99.995%（意味着 30 天内总共只有 2 分钟 10 秒的错误预算），那么扣除 2 分钟之后，剩余的错误预算只有 10 秒了。</p>

<p>我再给你举一个在观测云上具体的例子：如下图所示，假设 SLI 检测周期为 5 分钟，而这一次监测发现了一些异常事件（也就是有 SLI 不达标），根据叠加的结果，异常事件覆盖时间为 3 分钟，需要扣除 SLO 的错误预算额度 3 分钟。那这个时候，如果仍然 SLO 以 99.995% 为目标，这个月的错误预算就已经用完了（总共有 2 分钟 10 秒），也就是说这个月的 SLO 不达标了。</p>

<p><img src="assets/1ae733858aa981b21f0bbb5097df27b5.png" alt="图片" /></p>

<h2 id="slo-对哪些人很重要">SLO 对哪些人很重要？</h2>

<p>就像前面说的，SLO 最重要的是衡量最终用户使用产品和服务的质量。除此之外，SLO 的状态还会影响到企业内部人员（包括开发和运维）对产品和服务采取的措施。因此， SLO 对于最终用户、开发和运维都很重要。</p>

<h3 id="最终用户">最终用户</h3>

<p>无论是什么样的产品，最终用户都对他们获得的服务质量抱有期望。他们希望可以在任何给定时间访问应用程序，希望应用能够快速加载并返回正确的数据。虽然我们可以通过工单或者客服电话来听取最终用户的反馈，对产品进行改进，但这些渠道非常有局限。而且即使我们能够解决所有工单和问题，也不意味着最终用户一定满意，不代表达到了最终用户期望的服务水平。</p>

<p>实际上，始终保持 100% 的可靠性是不可能的。SLO 可以帮助你在产品创新（这将帮助你为最终用户提供更大价值，但有一定破坏稳定性的风险）和可靠性（这将使最终用户在使用产品和服务的时候感到满意）之间找到正确的平衡点。你的错误预算决定了，在你的服务质量下降到真正影响最终用户正常使用之前，开发工作能承受的不可靠性的程度。</p>

<h3 id="企业内部人员">企业内部人员</h3>

<p>为了让整个企业的主要利益相关者采用 SLO，需要他们就实际可实现的可靠性目标达成一致，尤其是考虑到业务的优先级和他们希望开展的项目。</p>

<p>过去，开发人员和运维工程师之间的分歧源于他们对立的目标和职责：开发人员旨在为产品和服务添加更多功能，而运维工程师则负责维护这些服务的稳定性。SLO 不仅可以推进业务成果，还可以促进文化转变，让开发和运维团队对应用程序的可靠性形成一种共同的责任感。</p>

<p>有了 SLO 和错误预算之后，团队就能够客观地决定优先考虑哪些项目或计划了。只要有剩余的错误预算，开发人员就可以发布新功能以此提高产品的整体质量，而运维工程师则可以更专注于长期可靠性项目，例如数据库维护和流程自动化等。</p>

<p>但是，当错误预算即将耗尽时，开发人员需要放慢或冻结功能工作，与运维团队密切合作，在违反任何 SLO 之前重新稳定系统。简而言之，错误预算是一种可量化的方法，它可以调整开发人员和运维工程师的工作和目标。</p>

<h2 id="如何选取合适的-sli">如何选取合适的 SLI？</h2>

<h3 id="哪些指标适合作为-sli">哪些指标适合作为 SLI</h3>

<p>现在，我们已经定义了一些和 SLO 相关的概念，那么具体应该关联哪些 SLI 呢？这需要我们先深入了解你的最终用户是怎样使用你的产品的，这是第一步，也是最重要的一步。</p>

<p>你需要了解你的最终用户如何与你的应用程序交互，他们会希望通过应用来达到什么目的，一般的使用习惯是怎样的，他们使用的这些功能后面对应的服务和基础设施又有哪些。</p>

<p>以电商平台为例，该如何选取 SLI 从而设置 SLO 呢？你需要首先弄清楚，你的最终用户如何使用网站或者 App。通常来说，最终用户需要能够登录、搜索商品、查看单个商品的详细信息、将商品添加到购物车，最后进行支付。这就是你的应用的关键用户旅程，这对于选择 SLI 很重要，因为这是会对最终用户的体验造成影响的。</p>

<p>随着你的基础架构越来越复杂，为每个数据库、消息队列和负载均衡器设置外部 SLO 变得越来越麻烦。相反，我建议你将你的系统组件组织成几个主要类别（例如，响应/请求、存储、数据管道），并在每个类别中指定 SLI。在选取 SLI 的时候，请记住：“所有 SLI 都是指标，但并非所有指标都是好的 SLI。” 这意味着，虽然你可能要跟踪成百上千个指标，<strong>但你应该关注最重要的指标：最能捕捉用户体验的指标。</strong></p>

<p>你可以使用下表（<a href="https://sre.google/workbook/implementing-slos/#slis-for-different-types-of-services" target="_blank">来自 Google 的 SRE 书籍</a>）作为参考。</p>

<ul>
<li><p>响应或者请求类型的服务。</p>

<ul>
<li>可用性：服务成功响应的请求比例。</li>
<li>延迟：响应请求需要多长时间，超过某个阈值的请求比例。</li>
<li>吞吐量：可以处理多少个请求。</li>
</ul></li>

<li><p>数据存储类型的服务。</p>

<ul>
<li>可用性：数据是否可以按需访问，可以成功读取和写入的比例。</li>
<li>延迟：读取和写入需要多长时间，超过某个阈值的比例。</li>
<li>耐用性：用户所需要的特定数据是否存在。</li>
</ul></li>

<li><p>数据管道（Pipeline，将输入的数据进行转换并进行输出，例如从多种来源收集日志并生成报告）。</p>

<ul>
<li>正确性：进入管道的产生正确的值的记录所占的比例。</li>
<li>新鲜度：新数据或处理结果需要多长时间出现。</li>
</ul></li>
</ul>

<p>让我再给你举一些例子。想象一下你的最终用户，也就是你的电商平台的购物者，卡在了结账页面上，他们要等待缓慢的支付端点返回响应。他们等待的时间越长，就越有可能放弃并且转向其他平台，这会给业务带来损失。所以从这个角度来说，页面上每一秒延迟的增加都与收入的显著减少相关。因此从这个例子中，我们可以看到响应延迟是在线电商平台跟踪的一个特别重要的 SLI，它能够确保他们的客户可以快速完成关键业务交易。</p>

<p>我们可以再看一下另一个层面的指标，基础设施，我们拿 CPU 利用率举例。我们说这不是一个很好的 SLI 候选，因为即使服务器的 CPU 利用率比较高了，但对于最终用户来说，可能仍然可以在页面上完成商品浏览、对比、加入购物车以及结账等操作，用户体验没有受到影响。</p>

<p>当然，并不是说基础设施就完全不用去关注了，我们还是需要设置监控，但并不是作为 SLI 来与 SLO 直接关联。这里的要点是，<strong>无论一个指标对你的内部团队有多重要，如果它的价值不直接影响用户满意度，那么它作为 SLI 就没有用处，反而可能带来告警的风暴，淹没了更加重要的信息。</strong></p>

<h3 id="将-sli-与-slo-进行关联">将 SLI 与 SLO 进行关联</h3>

<p>一旦确定了 SLI，你就需要为 SLI 设置目标值（或值范围），把它和 SLO 关联起来了。一般我们对 SLI 完整的定义是在过去多久的时间内，该指标需要满足的正常阈值范围。</p>

<p>例如，跟踪请求延迟可能是“在 30 天内，95% 的身份验证服务请求的延迟将小于 250 毫秒”。需要指出的是，这里的 95% 是 P95 的含义，即将响应耗时从小到大排列，顺序处于 95% 位置的值即为 P95 值。这里不选择平均数，是因为偶尔发生的极端值可能会极大地影响平均数，让平均数的统计失去了意义。</p>

<h2 id="小结">小结</h2>

<p>好了，这节课，我们探讨了如何选择正确的 SLI 并将其转换为定义明确的 SLO。通过使用 SLI 来衡量你们为用户提供的服务水平，并根据实际 SLO 和错误预算进行跟踪，你将能够更好地做出决策，提高功能速度和系统可靠性。</p>

<p>从另一个角度来说，设定多个 SLI 与相应的影响 SLO 的规则，相当于为这个系统可靠性工程定义了 OKR（Objective and Key Results，现在很多企业的目标管理方式），确保 SLO 的错误预算不会出现不足，这是整个团队的核心目标。</p>

<p>在下一讲，我会为你介绍如何使用可观测平台来进行 SLO 的跟踪和维护，并通过 SLO 的状态来决定下一步的行动计划。</p>

<h2 id="思考题">思考题</h2>

<p>在这节课的最后，留给你一道思考题。</p>

<p>如果你也参与了系统可靠性的维护，在这个过程中，你的应用软件是什么类型，属于什么行业？你是通过重点关注哪些指标来量化和保障系统可靠性的？</p>

<p>欢迎你在留言区和我交流讨论，我们下节课见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#91fdfdfda8a5a0a0a1a6d1f6fcf0f8fdbff2fefc" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f16455d4bcb3696',t:'MTczNDA5NjIzMC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
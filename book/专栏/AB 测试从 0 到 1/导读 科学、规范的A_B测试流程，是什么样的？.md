<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=导读&#32;科学、规范的A_B测试流程，是什么样的？>
        <link rel="icon" href="/static/favicon.png">
        <title>导读 科学、规范的A_B测试流程，是什么样的？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 用好A_B测试，你得这么学.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e7%94%a8%e5%a5%bdA_B%e6%b5%8b%e8%af%95%ef%bc%8c%e4%bd%a0%e5%be%97%e8%bf%99%e4%b9%88%e5%ad%a6.md">00 开篇词 用好A_B测试，你得这么学.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 统计基础（上）：系统掌握指标的统计属性.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/01%20%e7%bb%9f%e8%ae%a1%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e7%b3%bb%e7%bb%9f%e6%8e%8c%e6%8f%a1%e6%8c%87%e6%a0%87%e7%9a%84%e7%bb%9f%e8%ae%a1%e5%b1%9e%e6%80%a7.md">01 统计基础（上）：系统掌握指标的统计属性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 统计基础（下）：深入理解A_B测试中的假设检验.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/02%20%e7%bb%9f%e8%ae%a1%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3A_B%e6%b5%8b%e8%af%95%e4%b8%ad%e7%9a%84%e5%81%87%e8%ae%be%e6%a3%80%e9%aa%8c.md">02 统计基础（下）：深入理解A_B测试中的假设检验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 确定指标：指标这么多，到底如何来选择？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/04%20%e7%a1%ae%e5%ae%9a%e6%8c%87%e6%a0%87%ef%bc%9a%e6%8c%87%e6%a0%87%e8%bf%99%e4%b9%88%e5%a4%9a%ef%bc%8c%e5%88%b0%e5%ba%95%e5%a6%82%e4%bd%95%e6%9d%a5%e9%80%89%e6%8b%a9%ef%bc%9f.md">04 确定指标：指标这么多，到底如何来选择？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 选取实验单位：什么样的实验单位是合适的？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/05%20%e9%80%89%e5%8f%96%e5%ae%9e%e9%aa%8c%e5%8d%95%e4%bd%8d%ef%bc%9a%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%e5%ae%9e%e9%aa%8c%e5%8d%95%e4%bd%8d%e6%98%af%e5%90%88%e9%80%82%e7%9a%84%ef%bc%9f.md">05 选取实验单位：什么样的实验单位是合适的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 选择实验样本量：样本量越多越好吗？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/06%20%e9%80%89%e6%8b%a9%e5%ae%9e%e9%aa%8c%e6%a0%b7%e6%9c%ac%e9%87%8f%ef%bc%9a%e6%a0%b7%e6%9c%ac%e9%87%8f%e8%b6%8a%e5%a4%9a%e8%b6%8a%e5%a5%bd%e5%90%97%ef%bc%9f.md">06 选择实验样本量：样本量越多越好吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  分析测试结果：你得到的测试结果真的靠谱吗？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/07%20%20%e5%88%86%e6%9e%90%e6%b5%8b%e8%af%95%e7%bb%93%e6%9e%9c%ef%bc%9a%e4%bd%a0%e5%be%97%e5%88%b0%e7%9a%84%e6%b5%8b%e8%af%95%e7%bb%93%e6%9e%9c%e7%9c%9f%e7%9a%84%e9%9d%a0%e8%b0%b1%e5%90%97%ef%bc%9f.md">07  分析测试结果：你得到的测试结果真的靠谱吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 案例串讲：从0开始，搭建一个规范的A_B测试框架.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/08%20%e6%a1%88%e4%be%8b%e4%b8%b2%e8%ae%b2%ef%bc%9a%e4%bb%8e0%e5%bc%80%e5%a7%8b%ef%bc%8c%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%e8%a7%84%e8%8c%83%e7%9a%84A_B%e6%b5%8b%e8%af%95%e6%a1%86%e6%9e%b6.md">08 案例串讲：从0开始，搭建一个规范的A_B测试框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  测试结果不显著，要怎么改善？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/09%20%20%e6%b5%8b%e8%af%95%e7%bb%93%e6%9e%9c%e4%b8%8d%e6%98%be%e8%91%97%ef%bc%8c%e8%a6%81%e6%80%8e%e4%b9%88%e6%94%b9%e5%96%84%ef%bc%9f.md">09  测试结果不显著，要怎么改善？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 常见误区及解决方法（上）：多重检验问题和学习效应.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/10%20%e5%b8%b8%e8%a7%81%e8%af%af%e5%8c%ba%e5%8f%8a%e8%a7%a3%e5%86%b3%e6%96%b9%e6%b3%95%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a4%9a%e9%87%8d%e6%a3%80%e9%aa%8c%e9%97%ae%e9%a2%98%e5%92%8c%e5%ad%a6%e4%b9%a0%e6%95%88%e5%ba%94.md">10 常见误区及解决方法（上）：多重检验问题和学习效应.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 常见误区及解决方法（下）：辛普森悖论和实验组_对照组的独立性.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/11%20%e5%b8%b8%e8%a7%81%e8%af%af%e5%8c%ba%e5%8f%8a%e8%a7%a3%e5%86%b3%e6%96%b9%e6%b3%95%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%be%9b%e6%99%ae%e6%a3%ae%e6%82%96%e8%ae%ba%e5%92%8c%e5%ae%9e%e9%aa%8c%e7%bb%84_%e5%af%b9%e7%85%a7%e7%bb%84%e7%9a%84%e7%8b%ac%e7%ab%8b%e6%80%a7.md">11 常见误区及解决方法（下）：辛普森悖论和实验组_对照组的独立性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 什么情况下不适合做A_B测试？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/12%20%e4%bb%80%e4%b9%88%e6%83%85%e5%86%b5%e4%b8%8b%e4%b8%8d%e9%80%82%e5%90%88%e5%81%9aA_B%e6%b5%8b%e8%af%95%ef%bc%9f.md">12 什么情况下不适合做A_B测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 融会贯通：A_B测试面试必知必会（上）.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/13%20%e8%9e%8d%e4%bc%9a%e8%b4%af%e9%80%9a%ef%bc%9aA_B%e6%b5%8b%e8%af%95%e9%9d%a2%e8%af%95%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a%ef%bc%88%e4%b8%8a%ef%bc%89.md">13 融会贯通：A_B测试面试必知必会（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 举一反三：A_B测试面试必知必会（下）.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/14%20%e4%b8%be%e4%b8%80%e5%8f%8d%e4%b8%89%ef%bc%9aA_B%e6%b5%8b%e8%af%95%e9%9d%a2%e8%af%95%e5%bf%85%e7%9f%a5%e5%bf%85%e4%bc%9a%ef%bc%88%e4%b8%8b%ef%bc%89.md">14 举一反三：A_B测试面试必知必会（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 用R_Shiny，教你制作一个样本量计算器.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/15%20%e7%94%a8R_Shiny%ef%bc%8c%e6%95%99%e4%bd%a0%e5%88%b6%e4%bd%9c%e4%b8%80%e4%b8%aa%e6%a0%b7%e6%9c%ac%e9%87%8f%e8%ae%a1%e7%ae%97%e5%99%a8.md">15 用R_Shiny，教你制作一个样本量计算器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 试验意识改变决策模式，推动业务增长.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/%e5%8a%a0%e9%a4%90%20%e8%af%95%e9%aa%8c%e6%84%8f%e8%af%86%e6%94%b9%e5%8f%98%e5%86%b3%e7%ad%96%e6%a8%a1%e5%bc%8f%ef%bc%8c%e6%8e%a8%e5%8a%a8%e4%b8%9a%e5%8a%a1%e5%a2%9e%e9%95%bf.md">加餐 试验意识改变决策模式，推动业务增长.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="导读 科学、规范的A_B测试流程，是什么样的？.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/%e5%af%bc%e8%af%bb%20%e7%a7%91%e5%ad%a6%e3%80%81%e8%a7%84%e8%8c%83%e7%9a%84A_B%e6%b5%8b%e8%af%95%e6%b5%81%e7%a8%8b%ef%bc%8c%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%ef%bc%9f.md">导读 科学、规范的A_B测试流程，是什么样的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 实践是检验真理的唯一标准.md" href="/%e4%b8%93%e6%a0%8f/AB%20%e6%b5%8b%e8%af%95%e4%bb%8e%200%20%e5%88%b0%201/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%ae%9e%e8%b7%b5%e6%98%af%e6%a3%80%e9%aa%8c%e7%9c%9f%e7%90%86%e7%9a%84%e5%94%af%e4%b8%80%e6%a0%87%e5%87%86.md">结束语 实践是检验真理的唯一标准.md</a>
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
                            <h1 id="title" data-id="导读 科学、规范的A_B测试流程，是什么样的？" class="title">导读 科学、规范的A_B测试流程，是什么样的？</h1>
                            <div><p>你好，我是博伟。</p>

<p>前面两节课啊，我们花了很大力气去学习做A/B测试的理论前提，这也是为了让你夯实理论基础。不过啊，除非你是统计科班出身，否则我都会推荐你，在学习实战的时候呢，也要不断温习统计篇的内容，把理论与实践结合起来。如果觉得有必要，也可以把我在统计篇讲的统计概念和理论延伸开来，通过查看相关统计专业书籍来加深理解。</p>

<p>学完了统计理论，接下来就要开始设计实现做A/B测试了。不过在我总结A/B测试的流程之前呢，我要简单介绍下在实践中做A/B测试的准备工作，主要有两部分：<strong>数据</strong>和<strong>测试平台</strong>。</p>

<p>一方面，我们要有数据，包括用户在我们产品和业务中的各种行为，营销广告的表现效果等等，以便用来构建指标。因为A/B测试是建立在数据上的分析方法，正如“巧妇难为无米之炊”，没有数据的话，我们就不能通过A/B测试来比较谁好谁坏。</p>

<p>一般来说，只要是公司的数据基础架构做得好，埋点埋得到位的话，基本的常用指标都是可以满足的。</p>

<p>如果说我们要进行的A/B测试的指标比较新、比较特别，或者数据库没有很全面，没有现成的数据可以用来计算相应的指标，那么可以和数据团队进行协商，看能不能在现有的数据中找出可以替代的指标计算方法。</p>

<p>如果找不到相近的替代指标，那么就要和数据工程团队协商，看能不能构建这个数据，可能需要新的埋点，或者从第三方获得。</p>

<p>另一方面呢，我们要有合适的测试平台，来帮助我们具体实施A/B测试。可以是公司内部工程团队搭建的平台，也可以是第三方提供的平台。对于这些平台，我们在做A/B测试之前都是需要事先熟悉的，以便可以在平台上面设置和实施新的A/B测试。</p>

<p>当然，在做A/B测试时，数据库和测试平台是要通过API等方式有机结合起来的，这样我们在测试平台上设置和实施的A/B测试，才能通过数据来计算相应的指标。</p>

<p>以上的准备工作并不是每次A/B测试都要做的，更多的是第一次做A/B测试时才需要去做的准备，所以更像是A/B测试的基础设施。以下的流程才是我们每次去做A/B测试都要经历的，我把它们总结在一张图中，你可以看一下。</p>

<p><img src="assets/5fa96121d0fe4b2cb4060d118e057435.jpg" alt="" /></p>

<p>以上就是一个规范的做A/B测试的流程了。你看啊，A/B测试的实践性很强，但大体就是这么几步。在这门课里，我会着重讲最核心的5个部分，也就是确定目标和假设、确定指标、确定实验单位、估算样本量以及分析测试结果。</p>

<p>在整个流程中，除了随机分组的具体实施细节和具体实施测试外，其余环节我都会逐个讲解。你可能会问，为什么不能把全部环节讲解一遍呢？</p>

<p>其实啊，我会侧重讲解A/B测试的基本原理，实践中的具体流程，还有实践中遇到的常见问题及解决办法，这些都是偏经验和方法论的内容。不管你在哪家公司，处在哪个行业，用什么平台去实施A/B测试，这些经验和方法论都是通用的，学完之后你就可以应用到实践中。</p>

<p>至于随机分组的不同随机算法，以及实施测试所用的平台，这些更偏工程实施的细节，公司不同，平台不同，那么实施A/B测试时也会有很大的差别。比如A/B测试的平台，大公司一般会自己开发内部的测试平台，中小型公司则会利用第三方的测试平台。</p>

<p>所以啊，基础篇这几节课呢，我也希望你能在学习的同时，能够跟自己的工作联系起来。如果你在工作中做过A/B测试，但是觉得流程没有很系统化，你就可以把平时做的A/B测试和基础篇的流程进行参照对比，看看还有哪些不足的地方。同时，通过学习基础篇，也会让你知道为什么会有这些流程，它们背后的原理是什么，让你加深对流程的理解，应用起来更加得心应手。</p>

<p>如果你还没做过A/B测试，也没关系。我会结合实际案例，来给你深入讲解。如果有条件，学习完之后你就可以尝试做自己的第一个A/B测试啦！</p>

<p>最后，我还要说明一点。A/B测试的前提是数据，这里牵涉到一个公司的数据架构和埋点策略，更多的是工程和数据库建设的问题，不是我们A/B测试的重点。所以在接下来讲课的时候，我就假设我们已经能够追踪A/B测试所需要的数据了，至于如何追踪这些数据，如何埋点这种工程实施的细节我们这里就不展开讨论了。</p>

<p>好啦，了解了这些，就让我们正式开始A/B测试的旅程吧！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a8c4c4c4919c9999989fe8cfc5c9c1c486cbc7c5" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0d31705a0ae8fe',t:'MTczNDAwMTA0MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
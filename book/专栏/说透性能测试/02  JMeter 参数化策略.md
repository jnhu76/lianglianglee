<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=02&#32;&#32;JMeter&#32;参数化策略>
        <link rel="icon" href="/static/favicon.png">
        <title>02  JMeter 参数化策略 </title>
        
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
                        <a class="menu-item" id="00 开篇词  为什么每个测试人都要学好性能测试？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%af%8f%e4%b8%aa%e6%b5%8b%e8%af%95%e4%ba%ba%e9%83%bd%e8%a6%81%e5%ad%a6%e5%a5%bd%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%ef%bc%9f.md">00 开篇词  为什么每个测试人都要学好性能测试？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  JMeter 的核心概念.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/01%20%20JMeter%20%e7%9a%84%e6%a0%b8%e5%bf%83%e6%a6%82%e5%bf%b5.md">01  JMeter 的核心概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  JMeter 参数化策略.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/02%20%20JMeter%20%e5%8f%82%e6%95%b0%e5%8c%96%e7%ad%96%e7%95%a5.md">02  JMeter 参数化策略.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  构建并执行 JMeter 脚本的正确姿势.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/03%20%20%e6%9e%84%e5%bb%ba%e5%b9%b6%e6%89%a7%e8%a1%8c%20JMeter%20%e8%84%9a%e6%9c%ac%e7%9a%84%e6%ad%a3%e7%a1%ae%e5%a7%bf%e5%8a%bf.md">03  构建并执行 JMeter 脚本的正确姿势.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  JMeter 二次开发其实并不难.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/04%20%20JMeter%20%e4%ba%8c%e6%ac%a1%e5%bc%80%e5%8f%91%e5%85%b6%e5%ae%9e%e5%b9%b6%e4%b8%8d%e9%9a%be.md">04  JMeter 二次开发其实并不难.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  如何基于 JMeter API 开发性能测试平台？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/05%20%20%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20JMeter%20API%20%e5%bc%80%e5%8f%91%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%b9%b3%e5%8f%b0%ef%bc%9f.md">05  如何基于 JMeter API 开发性能测试平台？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  Nginx 在系统架构中的作用.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/06%20%20Nginx%20%e5%9c%a8%e7%b3%bb%e7%bb%9f%e6%9e%b6%e6%9e%84%e4%b8%ad%e7%9a%84%e4%bd%9c%e7%94%a8.md">06  Nginx 在系统架构中的作用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  你真的知道如何制定性能测试的目标吗？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/07%20%20%e4%bd%a0%e7%9c%9f%e7%9a%84%e7%9f%a5%e9%81%93%e5%a6%82%e4%bd%95%e5%88%b6%e5%ae%9a%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e7%9a%84%e7%9b%ae%e6%a0%87%e5%90%97%ef%bc%9f.md">07  你真的知道如何制定性能测试的目标吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  性能测试场景的分类和意义.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/08%20%20%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%9c%ba%e6%99%af%e7%9a%84%e5%88%86%e7%b1%bb%e5%92%8c%e6%84%8f%e4%b9%89.md">08  性能测试场景的分类和意义.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  如何制定一份有效的性能测试方案？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/09%20%20%e5%a6%82%e4%bd%95%e5%88%b6%e5%ae%9a%e4%b8%80%e4%bb%bd%e6%9c%89%e6%95%88%e7%9a%84%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e6%96%b9%e6%a1%88%ef%bc%9f.md">09  如何制定一份有效的性能测试方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  命令行监控 Linux 服务器的要点.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/10%20%20%e5%91%bd%e4%bb%a4%e8%a1%8c%e7%9b%91%e6%8e%a7%20Linux%20%e6%9c%8d%e5%8a%a1%e5%99%a8%e7%9a%84%e8%a6%81%e7%82%b9.md">10  命令行监控 Linux 服务器的要点.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  分布式服务链路监控以及报警方案.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/11%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e6%9c%8d%e5%8a%a1%e9%93%be%e8%b7%af%e7%9b%91%e6%8e%a7%e4%bb%a5%e5%8f%8a%e6%8a%a5%e8%ad%a6%e6%96%b9%e6%a1%88.md">11  分布式服务链路监控以及报警方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  如何把可视化监控也做得酷炫？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/12%20%20%e5%a6%82%e4%bd%95%e6%8a%8a%e5%8f%af%e8%a7%86%e5%8c%96%e7%9b%91%e6%8e%a7%e4%b9%9f%e5%81%9a%e5%be%97%e9%85%b7%e7%82%ab%ef%bc%9f.md">12  如何把可视化监控也做得酷炫？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  Docker 的制作、运行以及监控.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/13%20%20Docker%20%e7%9a%84%e5%88%b6%e4%bd%9c%e3%80%81%e8%bf%90%e8%a1%8c%e4%bb%a5%e5%8f%8a%e7%9b%91%e6%8e%a7.md">13  Docker 的制作、运行以及监控.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  如何从 CPU 飙升定位到热点方法？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/14%20%20%e5%a6%82%e4%bd%95%e4%bb%8e%20CPU%20%e9%a3%99%e5%8d%87%e5%ae%9a%e4%bd%8d%e5%88%b0%e7%83%ad%e7%82%b9%e6%96%b9%e6%b3%95%ef%bc%9f.md">14  如何从 CPU 飙升定位到热点方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  如何基于 JVM 分析内存使用对象？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/15%20%20%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20JVM%20%e5%88%86%e6%9e%90%e5%86%85%e5%ad%98%e4%bd%bf%e7%94%a8%e5%af%b9%e8%b1%a1%ef%bc%9f.md">15  如何基于 JVM 分析内存使用对象？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  如何通过 Arthas 定位代码链路问题？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/16%20%20%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20Arthas%20%e5%ae%9a%e4%bd%8d%e4%bb%a3%e7%a0%81%e9%93%be%e8%b7%af%e9%97%ae%e9%a2%98%ef%bc%9f.md">16  如何通过 Arthas 定位代码链路问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  如何应对 Redis 缓存穿透、击穿和雪崩？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/17%20%20%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%20Redis%20%e7%bc%93%e5%ad%98%e7%a9%bf%e9%80%8f%e3%80%81%e5%87%bb%e7%a9%bf%e5%92%8c%e9%9b%aa%e5%b4%a9%ef%bc%9f.md">17  如何应对 Redis 缓存穿透、击穿和雪崩？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  如何才能优化 MySQL 性能？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/18%20%20%e5%a6%82%e4%bd%95%e6%89%8d%e8%83%bd%e4%bc%98%e5%8c%96%20MySQL%20%e6%80%a7%e8%83%bd%ef%bc%9f.md">18  如何才能优化 MySQL 性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  如何根治慢 SQL？.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/19%20%20%e5%a6%82%e4%bd%95%e6%a0%b9%e6%b2%bb%e6%85%a2%20SQL%ef%bc%9f.md">19  如何根治慢 SQL？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 结束语  线上全链路性能测试实践总结.md" href="/%e4%b8%93%e6%a0%8f/%e8%af%b4%e9%80%8f%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95/20%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e7%ba%bf%e4%b8%8a%e5%85%a8%e9%93%be%e8%b7%af%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%ae%9e%e8%b7%b5%e6%80%bb%e7%bb%93.md">20 结束语  线上全链路性能测试实践总结.md</a>
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
                            <h1 id="title" data-id="02  JMeter 参数化策略" class="title">02  JMeter 参数化策略</h1>
                            <div><p>上一讲我梳理了 JMeter 的核心概念，希望你能够通过课程去理解并灵活的应用到实际工作中。这一讲我会带你学习一个重要的知识点：参数化。无论是从使用频率还是从参数化对性能测试结果的影响，它都是你做性能测试必须要掌握的。</p>

<h3 id="参数化是什么">参数化是什么</h3>

<p>简单来说，参数化就是选取不同的参数作为请求内容输入。使用 JMeter 测试时，测试数据的准备是一项重要的 工作。若要求每次传入的数据不一样，就需要进行参数化了。</p>

<h3 id="为什么要进行参数化">为什么要进行参数化</h3>

<p>刚才说到，若要求每次传入的数据不一样，就需要进行参数化。那为什么会有这种要求呢？我们来看两个场景。</p>

<ul>
<li><strong>数据被缓存导致测试结果不准确</strong></li>
</ul>

<p>缓存原本是为了让数据访问的速度接近 CPU 的处理速度而设置的临时存储区域，比如 cache。如今缓存的概念变得更广了，很多空间都可以设置客户端缓存、CDN 缓存等等。</p>

<p>当你频繁地请求某一条固定的数据时，这条数据就很容易被缓存，而不是每次都从数据库中去获取，这就可能导致和真实的场景有差别。</p>

<p>比如大促有 10w 用户会从获取不同的商品信息，而你的压测中并没有进行充分的参数化，此时用大量线程反复请求同一件商品，极有可能访问的是缓存数据。从业务逻辑看接口返回并没有问题，但这样的场景几乎不会发生，这就会导致无效压测，测试的结果并没有多少意义。</p>

<ul>
<li><strong>流程不能正常执行</strong></li>
</ul>

<p>数据被缓存可能会导致测试结果不准确，但至少业务能够走通。还有的情况是，在没有参数化的情况下，会产生大量的业务报错。</p>

<p>打个比方，你在测试限购商品抢购，如果用多线程模拟同一个用户操作可能会直接报错，因为在生成订单接口（支付等）都会判断是否是同一个用户。</p>

<p>在要求每次传入的数据不一样时，如果不进行参数化会造成很多问题，以上我列举的两个场景基本概括了没有参数化时会发生的问题，希望你能在性能测试时多加注意。下面我们就来看如何实现 JMeter 参数化。</p>

<h3 id="jmeter-参数化的实现方式">JMeter 参数化的实现方式</h3>

<p>我列举了 3 种比较常见的 JMeter 参数化的实现方式，你可以根据自己的需要进行选择。</p>

<ul>
<li><strong>CSV Data Set Config</strong>：将参数化的数据放入文件中，参数化读取依赖于文件操作。这样的参数化方式很常用，<strong>尤其适用于参数化数据量较多的场景</strong>，而且维护比较简单灵活。</li>
<li><strong>User Defined Variables</strong>：一般来说可以配置脚本中的公共参数，如域名，端口号，不需要随着压测进行动态改变，<strong>比较方便环境切换</strong>。</li>
<li><strong>Function Helper 中的函数</strong>：使用函数的方式生成参数，如果你需要随机数、uuid 等都可以使用函数生成。JMeter 还提供了相应的接口给你二次开发，自定义需要的功能。</li>
</ul>

<h4 id="csv-data-set-config">CSV Data Set Config</h4>

<p>CSV Data Set Config 的可配置选项较多，也是目前性能测试参数化时使用最多的插件，这里我就重点介绍一下 CSV Data Set Config。</p>

<p>在配置组件中添加元件 CSV Data Set Config，如下图所示：</p>

<p><img src="assets/Ciqc1F_7uM2AMFOoAAFcsKt5GDc012.png" alt="Drawing 0.png" /></p>

<p>图 1：CSV Data Set Config</p>

<p>我们来看一下 CSV Data Set Config 各项的含义。</p>

<p><strong>文件名</strong>：顾名思义，这里填写文件的名字即可。</p>

<p>保存参数化数据的文件目录，我这边是将 user.csv 和脚本放置在同一路径下。在这里我要推荐一个小技巧，就是“<strong>相对路径</strong>”。使用绝对路径，在做脚本迁移时大部分情况下都需要修改。如果你是先在 Windows 或 Mac 机器上修改脚本，再将脚本上传到 Linux 服务器上执行测试的，那你就可以用相对路径，这样就不用频繁修改该选项了。</p>

<p><strong>文件编码</strong>：指定文件的编码格式，设置的格式一般需要和文件编码格式一致，大部分情况下保存编码格式为 UTF-8 即可。</p>

<p><strong>变量名称</strong>：对应参数文件每列的变量名，类似 Excel 文件的文件头，主要是作为后续引用的标识符，一般使用英文。如下图所示：</p>

<p><img src="assets/CgqCHl_7uN-ALfhLAADCt_4kBrI773.png" alt="Drawing 1.png" /></p>

<p>图 2：user.csv</p>

<p>图中我标示了 name 和 password，相对应 user.csv 中的第一列和第二列。</p>

<p>那如何引用需要的文件数据？通过“${变量名称}”就可以了，如下图所示：</p>

<p><img src="assets/Ciqc1F_7uOaACU7pAACyoK2kugg837.png" alt="Drawing 2.png" /></p>

<p>图 3：引用演示图</p>

<p><strong>忽略首行</strong>： 第一行不读取。比如图 2 的第一行我只是标示这一列是什么类型的数据，实际上并不是需要读取的业务数据，此时就可以采用忽略首行。</p>

<p><strong>分隔符</strong>：用来标示参数文件中的分隔符号，与参数文件中的分隔符保持一致即可。</p>

<p><strong>遇到文件结束符再次循环</strong>：是否循环读取参数文件内容。因为 CSV Data Set Config 一次读入一行，如果设置的循环次数超过文本的行数，标示为 True 就是继续再从头开始读入。</p>

<p><strong>遇到文件结束符停止线程</strong>：读取到参数文件末尾时，是否停止读取线程，默认为 False。如果“遇到文件结束符再次循环”已经设置为 True 了，这个选项就没有意义了。</p>

<p><strong>线程共享模式</strong>：这个适用范围是一个脚本里多线程组的情况。所有线程是默认选项，代表当前测试计划中的所有线程中的所有的线程都有效；当前线程组代表当前线程组中的线程有效；当前线程则表示当前线程有效。一般情况下，我们选择默认选项“所有线程”就好，“当前线程组”和“当前线程”很少会用到。</p>

<p>上面我介绍了参数化的意义以及常见用法，参数化对于性能测试是基础且必需的，因为在性能场景中，很多时候不进行参数化，脚本也是可以跑通的，所以有一些测试同学在这方面就会“偷工减料”，但这会造成性能数据不准确。下面，我们就来看一种特殊的参数化：关联。</p>

<h3 id="特殊的参数化-关联">特殊的参数化：关联</h3>

<p>关联是将上个请求的数据提取需要的部分放到下个请求中，通过关联我们可以获取到满足业务特性的不同数据，因此可以认为是一种特殊的参数化。</p>

<h4 id="关联的使用场景">关联的使用场景</h4>

<p>我们来看一个例子，从例子中了解关联解决了什么问题。</p>

<p>我编写了一个查看订单接口，运行完成没有问题，并且返回正确的结果，如下所示：</p>

<pre><code>{&quot;data&quot;:{&quot;code&quot;:0,&quot;count&quot;:16,&quot;items&quot;:[{&quot;actualPrice&quot;:8900,&quot;gmtCreate&quot;:1601448530000,&quot;id&quot;:357,&quot;orderNo&quot;:&quot;1012020091448501002&quot;,&quot;skuList&quot;:[{&quot;barCode&quot;:&quot;150004&quot;,&quot;gmtCreate&quot;:1601448530000,&quot;gmtUpdate&quot;:1601448530000,&quot;id&quot;:389,&quot;img&quot;:&quot;https://demo.oss-cn-shenzhen.aliyuncs.com/bg/86338c9e576342baa0d079bc1caef9cc.jpg&quot;,&quot;num&quot;:1,&quot;orderId&quot;:357,&quot;orderNo&quot;:&quot;1012020091448501002&quot;,&quot;originalPrice&quot;:10690,&quot;price&quot;:8900,&quot;skuId&quot;:2777,&quot;spuId&quot;:1236771,&quot;spuTitle&quot;:&quot;昵趣 NaTruse 山羊奶配方狗狗洁齿骨 盒装 20g*40 支&quot;,&quot;title&quot;:&quot;山羊奶&quot;,&quot;unit&quot;:&quot;盒&quot;}],&quot;status&quot;:90},&quot;msg&quot;:&quot;第 1 页,共 1 条&quot;,&quot;pageNo&quot;:1,&quot;pageSize&quot;:1，&quot;total&quot;:1，&quot;totalPageNo&quot;:2},&quot;errmsg&quot;:&quot;成功&quot;,&quot;errno&quot;:200,&quot;timestamp&quot;:1609219480400}
</code></pre>

<p>一个小时之后，我再来运行这个接口时，却发现返回用户未登录：</p>

<pre><code> {&quot;errmsg&quot;:&quot;用户尚未登录&quot;,&quot;errno&quot;:10001,&quot;timestamp&quot;:1609220170295}
</code></pre>

<p>在所有入参都没有修改的情况下为什么会出现这样的情况呢？因为你看到返回的信息是用户未登录，也就是说用户信息是无效的。</p>

<p>这个接口使用 Token 验证用户，Token 有效期为 15 分钟，刚刚问题产生的原因就是 Token 过期了。</p>

<p>那如何保证查看订单接口信息中需要的 Token 都是有效的呢？其中一个方法就是查看订单接口之前调用登录接口获取 Token，把登录接口的 Token 传入查看订单接口中。这个过程其实就是“关联”。</p>

<h4 id="jmeter-如何实现关联">JMeter 如何实现关联</h4>

<p>JMeter实现关联有 3 种方式：边界提取器，通过左右边界的方式关联需要的数据；Json Extractor提取器，针对返回的 json 数据类型；正则表达式提取器，通过正则表达式去提取数据，实现关联。</p>

<p>正则表达式提取器是最为常用，也是这里我要向你介绍的关联方式。我们来看下面的例子：</p>

<p><img src="assets/Cip5yF_7uPGATVcpAABwbh3j2xc538.png" alt="Drawing 3.png" /></p>

<p>图 4：正则表达式提取器</p>

<p>我们来看一下正则表达式提取器中每一项的含义。</p>

<ul>
<li><strong>引用名称</strong>：自己定义的变量名称以及后续请求将要引用到的变量名。在图中我填写的是“token”，则引用方式是“${token}”。</li>
<li><strong>正则表达式</strong>：提取内容的正则表达式。“( )”括起来的部分就是需要提取的，“.”点号表示匹配任何字符串，“+”表示一次或多次，“？”表示找到第一个匹配项后停止。</li>
<li><strong>模板</strong>：用“<span class="math inline">\( \)</span>”引用，表示解析到的第几个值给 token，图 4 中的正则表达式如下：</li>
</ul>

<pre><code>&quot;accesstoken&quot;:(.+?),&quot;gender&quot;:(.+?)
</code></pre>

<p><span class="math inline">\(1\)</span> 表示匹配的第一个值，即 accesstoken 后匹配后的值，模板 <span class="math inline">\(2\)</span> 则匹配 gender 后的值。图 4 演示的实例中只有 1 个 token 值，所以使用的 <span class="math inline">\(1\)</span>。</p>

<ul>
<li><strong>匹配数字</strong>：0 代表随机取值，1 代表第一个值。假设我返回数据的结构是：</li>
</ul>

<pre><code>[{&quot;accesstoken&quot;:&quot;ABDS88WDWHJEHJSHWJW&quot;,&quot;gender&quot;:null},{{&quot;accesstoken&quot;:&quot;NDJNW3U98SJWKISXIWN&quot;,&quot;gender&quot;:null}]
</code></pre>

<p>此时提取结果是一个数组，accesstoken 对应了多个值相当于数组，1 代表匹配第一个 accsstoken 的值“ABDS88WDWHJEHJSHWJW”。</p>

<ul>
<li><strong>缺省值</strong>：正则匹配失败时的取值。比如这里我设置的是 null（token 值取不到时就会用 null 代替）。上面我们已经匹配了 token 值，在被测接口传参处直接用“${token}”就可以了。</li>
</ul>

<p><img src="assets/CgpVE1_7uPyAFTS7AAAim7aU0H8431.png" alt="Drawing 4.png" /></p>

<p>图 5：关联 Token</p>

<p>关联后就可以看到，每次都能进行正常的业务返回了。</p>

<p><img src="assets/CgpVE1_7uQWAYQQjAAIhTkpdTuc704.png" alt="Drawing 5.png" /></p>

<p>图 6：关联后正常业务返回</p>

<h3 id="总结">总结</h3>

<p>这一讲我介绍了参数化策略以及使用场景。作为性能测试中最常用到的操作，你不仅要学会基本操作，还需要思考参数化策略适合的场景以及参数化数据对性能测试的影响。说到这里，我就要问一个问题了：不同的参数对性能结果会不会有影响？</p>

<p>举个例子，在电商系统中，你准备了不同的用户数据，用户又分为不同的等级，不同的等级可能有不一样的优惠规则和对应的优惠券，每个会员的优惠券数量可能也不一样，那这些不同的会员信息分布会对性能测试的结果有什么样的的影响？欢迎在评论区给出你的思考。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d3bfbfbfeae7e2e2e3e493b4beb2babffdb0bcbe" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f18b281d92a635f',t:'MTczNDEyMTY3Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
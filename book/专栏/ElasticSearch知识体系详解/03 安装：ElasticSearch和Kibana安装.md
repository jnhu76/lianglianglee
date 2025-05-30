<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=03&#32;安装：ElasticSearch和Kibana安装>
        <link rel="icon" href="/static/favicon.png">
        <title>03 安装：ElasticSearch和Kibana安装 </title>
        
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
                        <a class="menu-item" id="01 认知：ElasticSearch基础概念.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/01%20%e8%ae%a4%e7%9f%a5%ef%bc%9aElasticSearch%e5%9f%ba%e7%a1%80%e6%a6%82%e5%bf%b5.md">01 认知：ElasticSearch基础概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 认知：Elastic Stack生态和场景方案.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/02%20%e8%ae%a4%e7%9f%a5%ef%bc%9aElastic%20Stack%e7%94%9f%e6%80%81%e5%92%8c%e5%9c%ba%e6%99%af%e6%96%b9%e6%a1%88.md">02 认知：Elastic Stack生态和场景方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 安装：ElasticSearch和Kibana安装.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/03%20%e5%ae%89%e8%a3%85%ef%bc%9aElasticSearch%e5%92%8cKibana%e5%ae%89%e8%a3%85.md">03 安装：ElasticSearch和Kibana安装.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 入门：查询和聚合的基础使用.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/04%20%e5%85%a5%e9%97%a8%ef%bc%9a%e6%9f%a5%e8%af%a2%e5%92%8c%e8%81%9a%e5%90%88%e7%9a%84%e5%9f%ba%e7%a1%80%e4%bd%bf%e7%94%a8.md">04 入门：查询和聚合的基础使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 索引：索引管理详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/05%20%e7%b4%a2%e5%bc%95%ef%bc%9a%e7%b4%a2%e5%bc%95%e7%ae%a1%e7%90%86%e8%af%a6%e8%a7%a3.md">05 索引：索引管理详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 索引：索引模板(Index Template)详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/06%20%e7%b4%a2%e5%bc%95%ef%bc%9a%e7%b4%a2%e5%bc%95%e6%a8%a1%e6%9d%bf%28Index%20Template%29%e8%af%a6%e8%a7%a3.md">06 索引：索引模板(Index Template)详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 查询：DSL查询之复合查询详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/07%20%e6%9f%a5%e8%af%a2%ef%bc%9aDSL%e6%9f%a5%e8%af%a2%e4%b9%8b%e5%a4%8d%e5%90%88%e6%9f%a5%e8%af%a2%e8%af%a6%e8%a7%a3.md">07 查询：DSL查询之复合查询详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 查询：DSL查询之全文搜索详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/08%20%e6%9f%a5%e8%af%a2%ef%bc%9aDSL%e6%9f%a5%e8%af%a2%e4%b9%8b%e5%85%a8%e6%96%87%e6%90%9c%e7%b4%a2%e8%af%a6%e8%a7%a3.md">08 查询：DSL查询之全文搜索详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 查询：DSL查询之Term详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/09%20%e6%9f%a5%e8%af%a2%ef%bc%9aDSL%e6%9f%a5%e8%af%a2%e4%b9%8bTerm%e8%af%a6%e8%a7%a3.md">09 查询：DSL查询之Term详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 聚合：聚合查询之Bucket聚合详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/10%20%e8%81%9a%e5%90%88%ef%bc%9a%e8%81%9a%e5%90%88%e6%9f%a5%e8%af%a2%e4%b9%8bBucket%e8%81%9a%e5%90%88%e8%af%a6%e8%a7%a3.md">10 聚合：聚合查询之Bucket聚合详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 聚合：聚合查询之Metric聚合详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/11%20%e8%81%9a%e5%90%88%ef%bc%9a%e8%81%9a%e5%90%88%e6%9f%a5%e8%af%a2%e4%b9%8bMetric%e8%81%9a%e5%90%88%e8%af%a6%e8%a7%a3.md">11 聚合：聚合查询之Metric聚合详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 聚合：聚合查询之Pipline聚合详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/12%20%e8%81%9a%e5%90%88%ef%bc%9a%e8%81%9a%e5%90%88%e6%9f%a5%e8%af%a2%e4%b9%8bPipline%e8%81%9a%e5%90%88%e8%af%a6%e8%a7%a3.md">12 聚合：聚合查询之Pipline聚合详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 原理：从图解构筑对ES原理的初步认知.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/13%20%e5%8e%9f%e7%90%86%ef%bc%9a%e4%bb%8e%e5%9b%be%e8%a7%a3%e6%9e%84%e7%ad%91%e5%af%b9ES%e5%8e%9f%e7%90%86%e7%9a%84%e5%88%9d%e6%ad%a5%e8%ae%a4%e7%9f%a5.md">13 原理：从图解构筑对ES原理的初步认知.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 原理：ES原理知识点补充和整体结构.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/14%20%e5%8e%9f%e7%90%86%ef%bc%9aES%e5%8e%9f%e7%90%86%e7%9f%a5%e8%af%86%e7%82%b9%e8%a1%a5%e5%85%85%e5%92%8c%e6%95%b4%e4%bd%93%e7%bb%93%e6%9e%84.md">14 原理：ES原理知识点补充和整体结构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 原理：ES原理之索引文档流程详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/15%20%e5%8e%9f%e7%90%86%ef%bc%9aES%e5%8e%9f%e7%90%86%e4%b9%8b%e7%b4%a2%e5%bc%95%e6%96%87%e6%a1%a3%e6%b5%81%e7%a8%8b%e8%af%a6%e8%a7%a3.md">15 原理：ES原理之索引文档流程详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 原理：ES原理之读取文档流程详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/16%20%e5%8e%9f%e7%90%86%ef%bc%9aES%e5%8e%9f%e7%90%86%e4%b9%8b%e8%af%bb%e5%8f%96%e6%96%87%e6%a1%a3%e6%b5%81%e7%a8%8b%e8%af%a6%e8%a7%a3.md">16 原理：ES原理之读取文档流程详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 优化：ElasticSearch性能优化详解.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/17%20%e4%bc%98%e5%8c%96%ef%bc%9aElasticSearch%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e8%af%a6%e8%a7%a3.md">17 优化：ElasticSearch性能优化详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 大厂实践：腾讯万亿级 Elasticsearch 技术实践.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/18%20%e5%a4%a7%e5%8e%82%e5%ae%9e%e8%b7%b5%ef%bc%9a%e8%85%be%e8%ae%af%e4%b8%87%e4%ba%bf%e7%ba%a7%20Elasticsearch%20%e6%8a%80%e6%9c%af%e5%ae%9e%e8%b7%b5.md">18 大厂实践：腾讯万亿级 Elasticsearch 技术实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 资料：Awesome Elasticsearch.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/19%20%e8%b5%84%e6%96%99%ef%bc%9aAwesome%20Elasticsearch.md">19 资料：Awesome Elasticsearch.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 WrapperQuery.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/20%20WrapperQuery.md">20 WrapperQuery.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 备份和迁移.md" href="/%e4%b8%93%e6%a0%8f/ElasticSearch%e7%9f%a5%e8%af%86%e4%bd%93%e7%b3%bb%e8%af%a6%e8%a7%a3/21%20%e5%a4%87%e4%bb%bd%e5%92%8c%e8%bf%81%e7%a7%bb.md">21 备份和迁移.md</a>
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
                            <h1 id="title" data-id="03 安装：ElasticSearch和Kibana安装" class="title">03 安装：ElasticSearch和Kibana安装</h1>
                            <div><h2 id="安装elasticsearch">安装ElasticSearch</h2>

<blockquote>
<p>ElasticSearch 是基于Java平台的，所以先要安装Java</p>
</blockquote>

<ul>
<li><strong>平台确认</strong></li>
</ul>

<p>这里我准备了一台Centos7虚拟机, 为方便选择后续安装的版本，所以需要看下系统版本信息。</p>

<pre><code class="language-bash">[root@VM-0-14-centos ~]# uname -a
Linux VM-0-14-centos 3.10.0-862.el7.x86_64 #1 SMP Fri Apr 20 16:44:24 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
    
</code></pre>

<ul>
<li><strong>安装Java</strong></li>
</ul>

<p>安装 Elasticsearch 之前，你需要先安装一个较新的版本的 Java，最好的选择是，你可以从 <a href="http://www.java.com" target="_blank">www.java.com  </a> 获得官方提供的最新版本的 Java。安装以后，确认是否安装成功：</p>

<pre><code class="language-bash">[root@VM-0-14-centos ~]# java --version
openjdk 14.0.2 2020-07-14
OpenJDK Runtime Environment 20.3 (slowdebug build 14.0.2+12)
OpenJDK 64-Bit Server VM 20.3 (slowdebug build 14.0.2+12, mixed mode, sharing)
    
</code></pre>

<ul>
<li><strong>下载ElasticSearch</strong></li>
</ul>

<p>从<a href="https://www.elastic.co/cn/downloads/elasticsearch" target="_blank">这里  </a>下载ElasticSearch</p>

<p>比如可以通过curl下载</p>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# curl -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.12.0-linux-x86_64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
    
</code></pre>

<ul>
<li><strong>解压</strong></li>
</ul>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# tar zxvf /opt/elasticsearch-7.12.0-linux-x86_64.tar.gz 
...
[root@VM-0-14-centos opt]# ll | grep elasticsearch
drwxr-xr-x  9 root root      4096 Mar 18 14:21 elasticsearch-7.12.0
-rw-r--r--  1 root root 327497331 Apr  5 21:05 elasticsearch-7.12.0-linux-x86_64.tar.gz
    
</code></pre>

<ul>
<li><strong>增加elasticSearch用户</strong></li>
</ul>

<p>必须创建一个非root用户来运行ElasticSearch(ElasticSearch5及以上版本，基于安全考虑，强制规定不能以root身份运行。)</p>

<p>如果你使用root用户来启动ElasticSearch，则会有如下错误信息：</p>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# cd elasticsearch-7.12.0/
[root@VM-0-14-centos elasticsearch-7.12.0]# ./bin/elasticsearch
[2021-04-05T21:36:46,510][ERROR][o.e.b.ElasticsearchUncaughtExceptionHandler] [VM-0-14-centos] uncaught exception in thread [main]
org.elasticsearch.bootstrap.StartupException: java.lang.RuntimeException: can not run elasticsearch as root
        at org.elasticsearch.bootstrap.Elasticsearch.init(Elasticsearch.java:163) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Elasticsearch.execute(Elasticsearch.java:150) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.cli.EnvironmentAwareCommand.execute(EnvironmentAwareCommand.java:75) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.cli.Command.mainWithoutErrorHandling(Command.java:116) ~[elasticsearch-cli-7.12.0.jar:7.12.0]
        at org.elasticsearch.cli.Command.main(Command.java:79) ~[elasticsearch-cli-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Elasticsearch.main(Elasticsearch.java:115) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Elasticsearch.main(Elasticsearch.java:81) ~[elasticsearch-7.12.0.jar:7.12.0]
Caused by: java.lang.RuntimeException: can not run elasticsearch as root
        at org.elasticsearch.bootstrap.Bootstrap.initializeNatives(Bootstrap.java:101) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Bootstrap.setup(Bootstrap.java:168) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Bootstrap.init(Bootstrap.java:397) ~[elasticsearch-7.12.0.jar:7.12.0]
        at org.elasticsearch.bootstrap.Elasticsearch.init(Elasticsearch.java:159) ~[elasticsearch-7.12.0.jar:7.12.0]
        ... 6 more
uncaught exception in thread [main]
java.lang.RuntimeException: can not run elasticsearch as root
        at org.elasticsearch.bootstrap.Bootstrap.initializeNatives(Bootstrap.java:101)
        at org.elasticsearch.bootstrap.Bootstrap.setup(Bootstrap.java:168)
        at org.elasticsearch.bootstrap.Bootstrap.init(Bootstrap.java:397)
        at org.elasticsearch.bootstrap.Elasticsearch.init(Elasticsearch.java:159)
        at org.elasticsearch.bootstrap.Elasticsearch.execute(Elasticsearch.java:150)
        at org.elasticsearch.cli.EnvironmentAwareCommand.execute(EnvironmentAwareCommand.java:75)
        at org.elasticsearch.cli.Command.mainWithoutErrorHandling(Command.java:116)
        at org.elasticsearch.cli.Command.main(Command.java:79)
        at org.elasticsearch.bootstrap.Elasticsearch.main(Elasticsearch.java:115)
        at org.elasticsearch.bootstrap.Elasticsearch.main(Elasticsearch.java:81)
For complete error details, refer to the log at /opt/elasticsearch-7.12.0/logs/elasticsearch.log
2021-04-05 13:36:46,979269 UTC [8846] INFO  Main.cc@106 Parent process died - ML controller exiting
    
</code></pre>

<p>所以我们增加一个独立的elasticsearch用户来运行</p>

<pre><code class="language-bash"># 增加elasticsearch用户
[root@VM-0-14-centos elasticsearch-7.12.0]# useradd elasticsearch
[root@VM-0-14-centos elasticsearch-7.12.0]# passwd elasticsearch
Changing password for user elasticsearch.
New password: 
BAD PASSWORD: The password contains the user name in some form
Retype new password: 
passwd: all authentication tokens updated successfully.

# 修改目录权限至新增的elasticsearch用户
[root@VM-0-14-centos elasticsearch-7.12.0]# chown -R elasticsearch /opt/elasticsearch-7.12.0
# 增加data和log存放区，并赋予elasticsearch用户权限
[root@VM-0-14-centos elasticsearch-7.12.0]# mkdir -p /data/es
[root@VM-0-14-centos elasticsearch-7.12.0]# chown -R elasticsearch /data/es
[root@VM-0-14-centos elasticsearch-7.12.0]# mkdir -p /var/log/es
[root@VM-0-14-centos elasticsearch-7.12.0]# chown -R elasticsearch /var/log/es

    
</code></pre>

<p>然后修改上述的data和log路径，<code>vi /opt/elasticsearch-7.12.0/config/elasticsearch.yml</code></p>

<pre><code class="language-bash"># ----------------------------------- Paths ------------------------------------
#
# Path to directory where to store the data (separate multiple locations by comma):
#
path.data: /data/es
#
# Path to log files:
#
path.logs: /var/log/es
    
</code></pre>

<ul>
<li><strong>修改Linux系统的限制配置</strong></li>
</ul>

<ol>
<li>修改系统中允许应用最多创建多少文件等的限制权限。Linux默认来说，一般限制应用最多创建的文件是65535个。但是ES至少需要65536的文件创建权限。</li>
<li>修改系统中允许用户启动的进程开启多少个线程。默认的Linux限制root用户开启的进程可以开启任意数量的线程，其他用户开启的进程可以开启1024个线程。必须修改限制数为4096+。因为ES至少需要4096的线程池预备。ES在5.x版本之后，强制要求在linux中不能使用root用户启动ES进程。所以必须使用其他用户启动ES进程才可以。</li>
<li>Linux低版本内核为线程分配的内存是128K。4.x版本的内核分配的内存更大。如果虚拟机的内存是1G，最多只能开启3000+个线程数。至少为虚拟机分配1.5G以上的内存。</li>
</ol>

<p>修改如下配置</p>

<pre><code class="language-bash">[root@VM-0-14-centos elasticsearch-7.12.0]# vi /etc/security/limits.conf

elasticsearch soft nofile 65536
elasticsearch hard nofile 65536
elasticsearch soft nproc 4096
elasticsearch hard nproc 4096
    
</code></pre>

<ul>
<li><strong>启动ElasticSearch</strong></li>
</ul>

<pre><code class="language-bash">[root@VM-0-14-centos elasticsearch-7.12.0]# su elasticsearch
[elasticsearch@VM-0-14-centos elasticsearch-7.12.0]$ ./bin/elasticsearch -d
[2021-04-05T22:03:38,332][INFO ][o.e.n.Node               ] [VM-0-14-centos] version[7.12.0], pid[13197], build[default/tar/78722783c38caa25a70982b5b042074cde5d3b3a/2021-03-18T06:17:15.410153305Z], OS[Linux/3.10.0-862.el7.x86_64/amd64], JVM[AdoptOpenJDK/OpenJDK 64-Bit Server VM/15.0.1/15.0.1+9]
[2021-04-05T22:03:38,348][INFO ][o.e.n.Node               ] [VM-0-14-centos] JVM home [/opt/elasticsearch-7.12.0/jdk], using bundled JDK [true]
[2021-04-05T22:03:38,348][INFO ][o.e.n.Node               ] [VM-0-14-centos] JVM arguments [-Xshare:auto, -Des.networkaddress.cache.ttl=60, -Des.networkaddress.cache.negative.ttl=10, -XX:+AlwaysPreTouch, -Xss1m, -Djava.awt.headless=true, -Dfile.encoding=UTF-8, -Djna.nosys=true, -XX:-OmitStackTraceInFastThrow, -XX:+ShowCodeDetailsInExceptionMessages, -Dio.netty.noUnsafe=true, -Dio.netty.noKeySetOptimization=true, -Dio.netty.recycler.maxCapacityPerThread=0, -Dio.netty.allocator.numDirectArenas=0, -Dlog4j.shutdownHookEnabled=false, -Dlog4j2.disable.jmx=true, -Djava.locale.providers=SPI,COMPAT, --add-opens=java.base/java.io=ALL-UNNAMED, -XX:+UseG1GC, -Djava.io.tmpdir=/tmp/elasticsearch-17264135248464897093, -XX:+HeapDumpOnOutOfMemoryError, -XX:HeapDumpPath=data, -XX:ErrorFile=logs/hs_err_pid%p.log, -Xlog:gc*,gc+age=trace,safepoint:file=logs/gc.log:utctime,pid,tags:filecount=32,filesize=64m, -Xms1894m, -Xmx1894m, -XX:MaxDirectMemorySize=993001472, -XX:G1HeapRegionSize=4m, -XX:InitiatingHeapOccupancyPercent=30, -XX:G1ReservePercent=15, -Des.path.home=/opt/elasticsearch-7.12.0, -Des.path.conf=/opt/elasticsearch-7.12.0/config, -Des.distribution.flavor=default, -Des.distribution.type=tar, -Des.bundled_jdk=true]
    
</code></pre>

<ul>
<li><strong>查看安装是否成功</strong></li>
</ul>

<pre><code class="language-bash">[root@VM-0-14-centos ~]# netstat -ntlp | grep 9200
tcp6       0      0 127.0.0.1:9200          :::*                    LISTEN      13549/java          
tcp6       0      0 ::1:9200                :::*                    LISTEN      13549/java          
[root@VM-0-14-centos ~]# curl 127.0.0.1:9200
{
  &quot;name&quot; : &quot;VM-0-14-centos&quot;,
  &quot;cluster_name&quot; : &quot;elasticsearch&quot;,
  &quot;cluster_uuid&quot; : &quot;ihttW8b2TfWSkwf_YgPH2Q&quot;,
  &quot;version&quot; : {
    &quot;number&quot; : &quot;7.12.0&quot;,
    &quot;build_flavor&quot; : &quot;default&quot;,
    &quot;build_type&quot; : &quot;tar&quot;,
    &quot;build_hash&quot; : &quot;78722783c38caa25a70982b5b042074cde5d3b3a&quot;,
    &quot;build_date&quot; : &quot;2021-03-18T06:17:15.410153305Z&quot;,
    &quot;build_snapshot&quot; : false,
    &quot;lucene_version&quot; : &quot;8.8.0&quot;,
    &quot;minimum_wire_compatibility_version&quot; : &quot;6.8.0&quot;,
    &quot;minimum_index_compatibility_version&quot; : &quot;6.0.0-beta1&quot;
  },
  &quot;tagline&quot; : &quot;You Know, for Search&quot;
}
    
</code></pre>

<h2 id="安装kibana">安装Kibana</h2>

<blockquote>
<p>Kibana是界面化的查询数据的工具，下载时尽量下载与ElasicSearch一致的版本。</p>
</blockquote>

<ul>
<li><strong>下载Kibana</strong></li>
</ul>

<p>从<a href="https://www.elastic.co/cn/downloads/kibana" target="_blank">这里  </a>下载Kibana</p>

<ul>
<li><strong>解压</strong></li>
</ul>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# tar -vxzf kibana-7.12.0-linux-x86_64.tar.gz
    
</code></pre>

<ul>
<li><strong>使用elasticsearch用户权限</strong></li>
</ul>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# chown -R elasticsearch /opt/kibana-7.12.0-linux-x86_64
#配置Kibana的远程访问
[root@VM-0-14-centos opt]# vi /opt/kibana-7.12.0-linux-x86_64/config/kibana.yml
server.host: 0.0.0.0
    
</code></pre>

<ul>
<li><strong>启动</strong></li>
</ul>

<p>需要切换至elasticsearch用户</p>

<pre><code class="language-bash">[root@VM-0-14-centos opt]# su elasticsearch
[elasticsearch@VM-0-14-centos opt]$ cd /opt/kibana-7.12.0-linux-x86_64/
[elasticsearch@VM-0-14-centos kibana-7.12.0-linux-x86_64]$ ./bin/kibana
  log   [22:30:22.185] [info][plugins-service] Plugin &quot;osquery&quot; is disabled.
  log   [22:30:22.283] [warning][config][deprecation] Config key [monitoring.cluster_alerts.email_notifications.email_address] will be required for email notifications to work in 8.0.&quot;
  log   [22:30:22.482] [info][plugins-system] Setting up [100] plugins: [taskManager,licensing,globalSearch,globalSearchProviders,banners,code,usageCollection,xpackLegacy,telemetryCollectionManager,telemetry,telemetryCollectionXpack,kibanaUsageCollection,securityOss,share,newsfeed,mapsLegacy,kibanaLegacy,translations,legacyExport,embeddable,uiActionsEnhanced,expressions,charts,esUiShared,bfetch,data,home,observability,console,consoleExtensions,apmOss,searchprofiler,painlessLab,grokdebugger,management,indexPatternManagement,advancedSettings,fileUpload,savedObjects,visualizations,visTypeVislib,visTypeVega,visTypeTimelion,features,licenseManagement,watcher,canvas,visTypeTagcloud,visTypeTable,visTypeMetric,visTypeMarkdown,tileMap,regionMap,visTypeXy,graph,timelion,dashboard,dashboardEnhanced,visualize,visTypeTimeseries,inputControlVis,discover,discoverEnhanced,savedObjectsManagement,spaces,security,savedObjectsTagging,maps,lens,reporting,lists,encryptedSavedObjects,dashboardMode,dataEnhanced,cloud,upgradeAssistant,snapshotRestore,fleet,indexManagement,rollup,remoteClusters,crossClusterReplication,indexLifecycleManagement,enterpriseSearch,beatsManagement,transform,ingestPipelines,eventLog,actions,alerts,triggersActionsUi,stackAlerts,ml,securitySolution,case,infra,monitoring,logstash,apm,uptime]
  log   [22:30:22.483] [info][plugins][taskManager] TaskManager is identified by the Kibana UUID: xxxxxx
  ...
    
</code></pre>

<p>如果是后台启动：</p>

<pre><code class="language-bash">[elasticsearch@VM-0-14-centos kibana-7.12.0-linux-x86_64]$ nohup ./bin/kibana &amp;
    
</code></pre>

<ul>
<li><strong>界面访问</strong></li>
</ul>

<p><img src="assets/es-install-1.png" alt="img" /></p>

<p>可以导入simple data</p>

<p><img src="assets/es-install-2.png" alt="img" /></p>

<p>查看数据</p>

<p><img src="assets/es-install-3.png" alt="img" /></p>

<h2 id="配置密码访问">配置密码访问</h2>

<blockquote>
<p>使用基本许可证时，默认情况下禁用Elasticsearch安全功能。由于我测试环境是放在公网上的，所以需要设置下密码访问。相关文档可以参考<a href="https://www.elastic.co/guide/en/elasticsearch/reference/7.12/security-minimal-setup.html" target="_blank">这里  </a></p>
</blockquote>

<ol>
<li>停止kibana和elasticsearch服务</li>
<li>将<code>xpack.security.enabled</code>设置添加到ES_PATH_CONF/elasticsearch.yml文件并将值设置为true</li>
<li>启动elasticsearch (<code>./bin/elasticsearch -d</code>)</li>
<li>执行如下密码设置器，<code>./bin/elasticsearch-setup-passwords interactive</code>来设置各个组件的密码</li>
<li>将elasticsearch.username设置添加到KIB_PATH_CONF/kibana.yml 文件并将值设置给elastic用户： <code>elasticsearch.username: &quot;elastic&quot;</code></li>
<li>创建kibana keystore, <code>./bin/kibana-keystore create</code></li>
<li>在kibana keystore 中添加密码 <code>./bin/kibana-keystore add elasticsearch.password</code></li>
<li>重启kibana 服务即可 <code>nohup ./bin/kibana &amp;</code></li>
</ol>

<p>然后就可以使用密码登录了：</p>

<p><img src="assets/es-install-4.png" alt="img" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#6e020202575a5f5f5e592e09030f0702400d0103" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0d9abcfc2ae8fe',t:'MTczNDAwNTM1NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
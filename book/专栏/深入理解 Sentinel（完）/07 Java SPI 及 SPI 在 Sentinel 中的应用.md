<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=07&#32;Java&#32;SPI&#32;及&#32;SPI&#32;在&#32;Sentinel&#32;中的应用>
        <link rel="icon" href="/static/favicon.png">
        <title>07 Java SPI 及 SPI 在 Sentinel 中的应用 </title>
        
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
                        <a class="menu-item" id="01 开篇词：一次服务雪崩问题排查经历.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/01%20%e5%bc%80%e7%af%87%e8%af%8d%ef%bc%9a%e4%b8%80%e6%ac%a1%e6%9c%8d%e5%8a%a1%e9%9b%aa%e5%b4%a9%e9%97%ae%e9%a2%98%e6%8e%92%e6%9f%a5%e7%bb%8f%e5%8e%86.md">01 开篇词：一次服务雪崩问题排查经历.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 为什么需要服务降级以及常见的几种降级方式.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/02%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e6%9c%8d%e5%8a%a1%e9%99%8d%e7%ba%a7%e4%bb%a5%e5%8f%8a%e5%b8%b8%e8%a7%81%e7%9a%84%e5%87%a0%e7%a7%8d%e9%99%8d%e7%ba%a7%e6%96%b9%e5%bc%8f.md">02 为什么需要服务降级以及常见的几种降级方式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 为什么选择 Sentinel，Sentinel 与 Hystrix 的对比.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/03%20%e4%b8%ba%e4%bb%80%e4%b9%88%e9%80%89%e6%8b%a9%20Sentinel%ef%bc%8cSentinel%20%e4%b8%8e%20Hystrix%20%e7%9a%84%e5%af%b9%e6%af%94.md">03 为什么选择 Sentinel，Sentinel 与 Hystrix 的对比.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 Sentinel 基于滑动窗口的实时指标数据统计.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/04%20Sentinel%20%e5%9f%ba%e4%ba%8e%e6%bb%91%e5%8a%a8%e7%aa%97%e5%8f%a3%e7%9a%84%e5%ae%9e%e6%97%b6%e6%8c%87%e6%a0%87%e6%95%b0%e6%8d%ae%e7%bb%9f%e8%ae%a1.md">04 Sentinel 基于滑动窗口的实时指标数据统计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 Sentinel 的一些概念与核心类介绍.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/05%20Sentinel%20%e7%9a%84%e4%b8%80%e4%ba%9b%e6%a6%82%e5%bf%b5%e4%b8%8e%e6%a0%b8%e5%bf%83%e7%b1%bb%e4%bb%8b%e7%bb%8d.md">05 Sentinel 的一些概念与核心类介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 Sentinel 中的责任链模式与 Sentinel 的整体工作流程.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/06%20Sentinel%20%e4%b8%ad%e7%9a%84%e8%b4%a3%e4%bb%bb%e9%93%be%e6%a8%a1%e5%bc%8f%e4%b8%8e%20Sentinel%20%e7%9a%84%e6%95%b4%e4%bd%93%e5%b7%a5%e4%bd%9c%e6%b5%81%e7%a8%8b.md">06 Sentinel 中的责任链模式与 Sentinel 的整体工作流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 Java SPI 及 SPI 在 Sentinel 中的应用.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/07%20Java%20SPI%20%e5%8f%8a%20SPI%20%e5%9c%a8%20Sentinel%20%e4%b8%ad%e7%9a%84%e5%ba%94%e7%94%a8.md">07 Java SPI 及 SPI 在 Sentinel 中的应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 资源指标数据统计的实现全解析（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/08%20%e8%b5%84%e6%ba%90%e6%8c%87%e6%a0%87%e6%95%b0%e6%8d%ae%e7%bb%9f%e8%ae%a1%e7%9a%84%e5%ae%9e%e7%8e%b0%e5%85%a8%e8%a7%a3%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89.md">08 资源指标数据统计的实现全解析（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 资源指标数据统计的实现全解析（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/09%20%e8%b5%84%e6%ba%90%e6%8c%87%e6%a0%87%e6%95%b0%e6%8d%ae%e7%bb%9f%e8%ae%a1%e7%9a%84%e5%ae%9e%e7%8e%b0%e5%85%a8%e8%a7%a3%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89.md">09 资源指标数据统计的实现全解析（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 限流降级与流量效果控制器（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/10%20%e9%99%90%e6%b5%81%e9%99%8d%e7%ba%a7%e4%b8%8e%e6%b5%81%e9%87%8f%e6%95%88%e6%9e%9c%e6%8e%a7%e5%88%b6%e5%99%a8%ef%bc%88%e4%b8%8a%ef%bc%89.md">10 限流降级与流量效果控制器（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 限流降级与流量效果控制器（中）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/11%20%e9%99%90%e6%b5%81%e9%99%8d%e7%ba%a7%e4%b8%8e%e6%b5%81%e9%87%8f%e6%95%88%e6%9e%9c%e6%8e%a7%e5%88%b6%e5%99%a8%ef%bc%88%e4%b8%ad%ef%bc%89.md">11 限流降级与流量效果控制器（中）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 限流降级与流量效果控制器（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/12%20%e9%99%90%e6%b5%81%e9%99%8d%e7%ba%a7%e4%b8%8e%e6%b5%81%e9%87%8f%e6%95%88%e6%9e%9c%e6%8e%a7%e5%88%b6%e5%99%a8%ef%bc%88%e4%b8%8b%ef%bc%89.md">12 限流降级与流量效果控制器（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 熔断降级与系统自适应限流.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/13%20%e7%86%94%e6%96%ad%e9%99%8d%e7%ba%a7%e4%b8%8e%e7%b3%bb%e7%bb%9f%e8%87%aa%e9%80%82%e5%ba%94%e9%99%90%e6%b5%81.md">13 熔断降级与系统自适应限流.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 黑白名单限流与热点参数限流.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/14%20%e9%bb%91%e7%99%bd%e5%90%8d%e5%8d%95%e9%99%90%e6%b5%81%e4%b8%8e%e7%83%ad%e7%82%b9%e5%8f%82%e6%95%b0%e9%99%90%e6%b5%81.md">14 黑白名单限流与热点参数限流.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 自定义 ProcessorSlot 实现开关降级.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/15%20%e8%87%aa%e5%ae%9a%e4%b9%89%20ProcessorSlot%20%e5%ae%9e%e7%8e%b0%e5%bc%80%e5%85%b3%e9%99%8d%e7%ba%a7.md">15 自定义 ProcessorSlot 实现开关降级.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Sentinel 动态数据源：规则动态配置.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/16%20Sentinel%20%e5%8a%a8%e6%80%81%e6%95%b0%e6%8d%ae%e6%ba%90%ef%bc%9a%e8%a7%84%e5%88%99%e5%8a%a8%e6%80%81%e9%85%8d%e7%bd%ae.md">16 Sentinel 动态数据源：规则动态配置.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 Sentinel 主流框架适配.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/17%20Sentinel%20%e4%b8%bb%e6%b5%81%e6%a1%86%e6%9e%b6%e9%80%82%e9%85%8d.md">17 Sentinel 主流框架适配.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 Sentinel 集群限流的实现（上）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/18%20Sentinel%20%e9%9b%86%e7%be%a4%e9%99%90%e6%b5%81%e7%9a%84%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8a%ef%bc%89.md">18 Sentinel 集群限流的实现（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Sentinel 集群限流的实现（下）.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/19%20Sentinel%20%e9%9b%86%e7%be%a4%e9%99%90%e6%b5%81%e7%9a%84%e5%ae%9e%e7%8e%b0%ef%bc%88%e4%b8%8b%ef%bc%89.md">19 Sentinel 集群限流的实现（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 结束语：Sentinel 对应用的性能影响如何？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/20%20%e7%bb%93%e6%9d%9f%e8%af%ad%ef%bc%9aSentinel%20%e5%af%b9%e5%ba%94%e7%94%a8%e7%9a%84%e6%80%a7%e8%83%bd%e5%bd%b1%e5%93%8d%e5%a6%82%e4%bd%95%ef%bc%9f.md">20 结束语：Sentinel 对应用的性能影响如何？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 番外篇：Sentinel 1.8.0 熔断降级新特性解读.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Sentinel%ef%bc%88%e5%ae%8c%ef%bc%89/21%20%e7%95%aa%e5%a4%96%e7%af%87%ef%bc%9aSentinel%201.8.0%20%e7%86%94%e6%96%ad%e9%99%8d%e7%ba%a7%e6%96%b0%e7%89%b9%e6%80%a7%e8%a7%a3%e8%af%bb.md">21 番外篇：Sentinel 1.8.0 熔断降级新特性解读.md</a>
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
                            <h1 id="title" data-id="07 Java SPI 及 SPI 在 Sentinel 中的应用" class="title">07 Java SPI 及 SPI 在 Sentinel 中的应用</h1>
                            <div><p>SPI 全称是 Service Provider Interface，直译就是服务提供者接口，是一种服务发现机制，是 Java 的一个内置标准，允许不同的开发者去实现某个特定的服务。SPI 的本质是将接口实现类的全限定名配置在文件中，由服务加载器读取配置文件，加载实现类，实现在运行时动态替换接口的实现类。</p>

<p>使用 SPI 机制能够实现按配置加载接口的实现类，SPI 机制在阿里开源的项目中被广泛使用，例如 Dubbo、RocketMQ、以及本文介绍的 Sentinel。RocketMQ 与 Sentinel 使用的都是 Java 提供的 SPI 机制，而 Dubbo 则是使用自实现的一套 SPI，与 Java SPI 的配置方式不同，Dubbo SPI 使用 Key-Value 方式配置，目的是实现自适应扩展机制。</p>

<h3 id="java-spi-简介">Java SPI 简介</h3>

<p>我们先来个“Hello Word”级别的 Demo 学习一下 Java SPI 怎么使用，通过这个例子认识 Java SPI。</p>

<p><strong>第一步：定义接口</strong></p>

<p>假设我们有多种登录方式，则创建一个 LoginService 接口。</p>

<pre><code class="language-java">public interface LoginService{
  void login(String username,String password);
}

</code></pre>

<p><strong>第二步：编写接口实现类</strong></p>

<p>假设一开始我们使用 Shiro 框架实现用户鉴权，提供了一个 ShiroLoginService。</p>

<pre><code class="language-java">public class ShiroLoginService implements LoginService{
  public void login(String username,String password){
    // 实现登陆
  }
}

</code></pre>

<p>现在我们不想搞那么麻烦，比如我们可以直接使用 Spring MVC 的拦截器实现用户鉴权，那么可以提供一个 SpringLoginService。</p>

<pre><code class="language-java">public class SpringLoginService implements LoginService{
  public void login(String username,String password){
    // 实现登陆
  }
}

</code></pre>

<p><strong>第三步：通过配置使用 SpringLoginService 或 ShiroLoginService。</strong></p>

<p>当我们想通过修改配置文件的方式而不修改代码实现权限验证框架的切换，就可以使用 Java 的 SPI。通过运行时从配置文件中读取实现类，加载使用配置的实现类。</p>

<p>需要在 resources 目录下新建一个目录 META-INF，并在 META-INF 目录下创建 services 目录，用来存放接口配置文件。</p>

<p>配置文件名为接口 LoginService 全类名，文件中写入使用的实现类的全类名。只要是在 META-INF/services 目录下，只要文件名是接口的全类名，那么编写配置文件内容的时候，IDEA 就会自动提示有哪些实现类。</p>

<p>文件：/resources/META-INF/services/接口名称，填写的内容为接口的实现类，多个实现类使用换行分开。</p>

<pre><code>com.wujiuye.spi.ShiroLoginService

</code></pre>

<p><strong>第四步：编写 main 方法测试使用 Java SPI 加载 LoginService。</strong></p>

<pre><code class="language-java">public class JavaSPI{
  public static void main(String[] args){
    ServiceLoader&lt;LoginService&gt; serviceLoader = ServiceLoader.load(ServiceLoader.class);
    for(LoginService serviceImpl:serviceLoader){
      serviceImpl.login(&quot;wujiuye&quot;,&quot;123456&quot;);
    }
  }
}

</code></pre>

<p>ServiceLoader（服务加载器）是 Java 提供的 SPI 机制的实现，调用 load 方法传入接口名就能获取到一个 ServiceLoader 实例，此时配置文件中注册的实现类是还没有加载到 JVM 的，只有通过 Iterator 遍历获取的时候，才会去加载实现类与实例化实现类。</p>

<p>需要说明的是，例子中配置文件只配置了一个实现类，但其实我们是可以配置 N 多个的，并且 iterator 遍历的顺序就是配置文件中注册实现类的顺序。如果非要想一个注册多实现类的适用场景的话，责任链（拦截器、过滤器）模式这种可插拔的设计模式最适合不过。又或者一个画图程序，定义一个形状接口，实现类可以有矩形、三角形等，如果后期添加了圆形，只需要在形状接口的配置文件中注册圆形就能支持画圆形，完全不用修改任何代码。</p>

<p>ServiceLoader 源码很容易理解，就是根据传入的接口获取接口的全类名，将前缀“/META-INF/services”与接口的全类名拼接定位到配置文件，读取配置文件中的字符串、解析字符串，将解析出来的实现类全类名添加到一个数组，返回一个 ServiceLoader 实例。只有在遍历迭代器的时候 ServiceLoader 才通过调用 Class#forName 方法加载类并且通过反射创建实例，如果不指定类加载器，就使用当前线程的上下文类加载器加载类。</p>

<h3 id="java-spi-在-sentinel-中的应用">Java SPI 在 Sentinel 中的应用</h3>

<p>在 sentinel-core 模块的 resources 资源目录下，有一个 META-INF/services 目录，该目录下有两个以接口全名命名的文件，其中 com.alibaba.csp.sentinel.slotchain.SlotChainBuilder 文件用于配置 SlotChainBuilder 接口的实现类，com.alibaba.csp.sentinel.init.InitFunc 文件用于配置 InitFunc 接口的实现类，并且这两个配置文件中都配置了接口的默认实现类，如果我们不添加新的配置，Sentinel 将使用默认配置的接口实现类。</p>

<p>com.alibaba.csp.sentinel.init.InitFunc 文件的默认配置如下：</p>

<pre><code class="language-txt">  com.alibaba.csp.sentinel.metric.extension.MetricCallbackInit

</code></pre>

<p>com.alibaba.csp.sentinel.slotchain.SlotChainBuilder 文件的默认配置如下：</p>

<pre><code class="language-java">  # Default slot chain builder
  com.alibaba.csp.sentinel.slots.DefaultSlotChainBuilder

</code></pre>

<p>ServiceLoader 可加载接口配置文件中配置的所有实现类并且使用反射创建对象，但是否全部加载以及实例化还是由使用者自己决定。Sentinel 的 core 模块使用 Java SPI 机制加载 InitFunc 与 SlotChainBuilder 的实现上稍有不同，如果 InitFunc 接口的配置文件注册了多个实现类，那么这些注册的 InitFunc 实现类都会被 Sentinel 加载、实例化，且都会被使用，但 SlotChainBuilder 不同，如果注册多个实现类，Sentinel 只会加载和使用第一个。</p>

<p>调用 ServiceLoader#load 方法传入接口可获取到一个 ServiceLoader 实例，ServiceLoader 实现了 Iterable 接口，所以可以使用 forEach 语法遍历，ServiceLoader 使用 layz 方式实现迭代器（Iterator），只有被迭代器的 next 方法遍历到的类才会被加载和实例化。如果只想使用接口配置文件中注册的第一个实现类，那么可在使用迭代器遍历时，使用 break 跳出循环。</p>

<p>Sentinel 在加载 SlotChainBuilder 时，只会获取第一个非默认（非 DefaultSlotChainBuilder）实现类的实例，如果接口配置文件中除了默认实现类没有注册别的实现类，则 Sentinel 会使用这个默认的 SlotChainBuilder。其实现源码在 SpiLoader 的 loadFirstInstanceOrDefault 方法中，代码如下。</p>

<pre><code class="language-java">public final class SpiLoader {
    public static &lt;T&gt; T loadFirstInstanceOrDefault(Class&lt;T&gt; clazz, Class&lt;? extends T&gt; defaultClass) {
        try {
            // 缓存的实现省略...
            // 返回第一个类型不等于 defaultClass 的实例
            // ServiceLoader 实现了 Iterable 接口
            for (T instance : serviceLoader) {
                // 获取第一个非默认类的实例
                if (instance.getClass() != defaultClass) {
                    return instance;
                }
            }
            // 没有则使用默认类的实例，反射创建对象
            return defaultClass.newInstance();
        } catch (Throwable t) {
            return null;
        }
    }
}

</code></pre>

<p>Sentinel 加载 InitFunc 则不同，因为 Sentinel 允许存在多个初始化方法。InitFunc 可用于初始化配置限流、熔断规则，但在 Web 项目中我们基本不会使用它，更多的是通过监听 Spring 容器刷新完成事件再去初始化为 Sentinel 配置规则，如果使用动态数据源还可在监听到动态配置改变事件时重新加载规则，所以 InitFunc 我们基本使用不上。</p>

<p>Sentinel 使用 ServiceLoader 加载注册的 InitFunc 实现代码如下：</p>

<pre><code class="language-java">public final class InitExecutor {

    public static void doInit() {
        try {
            // 加载配置
            ServiceLoader&lt;InitFunc&gt; loader = ServiceLoaderUtil.getServiceLoader(InitFunc.class);
            List&lt;OrderWrapper&gt; initList = new ArrayList&lt;OrderWrapper&gt;();
            for (InitFunc initFunc : loader) {
               // 插入数组并排序，同时将 InitFunc 包装为 OrderWrapper
                insertSorted(initList, initFunc);
            }
            // 遍历调用 InitFunc 的初始化方法
            for (OrderWrapper w : initList) {
                w.func.init();
            }
        } catch (Exception ex) {
            ex.printStackTrace();
        } catch (Error error) {
            error.printStackTrace();
        }
    }
}

</code></pre>

<p>虽然 InitFunc 接口与 SlotChainBuilder 接口的配置文件在 sentinel-core 模块下，但我们不需要去修改 Sentinel 的源码，不需要修改 sentinel-core 模块下的接口配置文件，而只需要在当前项目的 /resource/META-INF/services 目录下创建一个与接口全名相同名称的配置文件，并在配置文件中添加接口的实现类即可。项目编译后不会覆盖 sentinel-core 模块下的相同名称的配置文件，而是将两个配置文件合并成一个配置文件。</p>

<h3 id="自定义-processorslotchain-构造器">自定义 ProcessorSlotChain 构造器</h3>

<p>Sentinel 使用 SlotChainBuilder 将多个 ProcessorSlot 构造成一个 ProcessorSlotChain，由 ProcessorSlotChain 按照 ProcessorSlot 的注册顺序去调用这些 ProcessorSlot。Sentinel 使用 Java SPI 加载 SlotChainBuilder 支持使用者自定义 SlotChainBuilder，相当于是提供了插件的功能。</p>

<p>Sentinel 默认使用的 SlotChainBuilder 是 DefaultSlotChainBuilder，其源码如下：</p>

<pre><code class="language-java">public class DefaultSlotChainBuilder implements SlotChainBuilder {

    @Override
    public ProcessorSlotChain build() {
        ProcessorSlotChain chain = new DefaultProcessorSlotChain();
        chain.addLast(new NodeSelectorSlot());
        chain.addLast(new ClusterBuilderSlot());
        chain.addLast(new LogSlot());
        chain.addLast(new StatisticSlot());
        chain.addLast(new AuthoritySlot());
        chain.addLast(new SystemSlot());
        chain.addLast(new FlowSlot());
        chain.addLast(new DegradeSlot());
        return chain;
    }

}

</code></pre>

<p>DefaultSlotChainBuilder 构造的 ProcessorSlotChain 注册了 NodeSelectorSlot、ClusterBuilderSlot、LogSlot、StatisticSlot、AuthoritySlot、SystemSlot、FlowSlot、DegradeSlot，但这些 ProcessorSlot 并非都是必须的，如果注册的这些 ProcessorSlot 有些我们用不到，那么我们可以自己实现一个 SlotChainBuilder，自己构造 ProcessorSlotChain。例如，我们可以将 LogSlot、AuthoritySlot、SystemSlot 去掉。</p>

<p>第一步，编写 MySlotChainBuilder，实现 SlotChainBuilder 接口，代码如下：</p>

<pre><code class="language-java">public class MySlotChainBuilder implements SlotChainBuilder {

    @Override
    public ProcessorSlotChain build() {
        ProcessorSlotChain chain = new DefaultProcessorSlotChain();
        chain.addLast(new NodeSelectorSlot());
        chain.addLast(new ClusterBuilderSlot());
        chain.addLast(new StatisticSlot());
        chain.addLast(new FlowSlot());
        chain.addLast(new DegradeSlot());
        return chain;
    }

}

</code></pre>

<p>第二步，在当前项目的 /resources/META-INF/services 目录下添加名为 com.alibaba.csp.sentinel.slotchain.SlotChainBuilder 的接口配置文件，并在配置文件中注册 MySlotChainBuilder。</p>

<pre><code class="language-txt">com.wujiuye.sck.provider.config.MySlotChainBuilder

</code></pre>

<p>在构造 ProcessorSlotChain 时，需注意 ProcessorSlot 的注册顺序，例如，NodeSelectorSlot 需作为 ClusterBuilderSlot 的前驱节点，ClusterBuilderSlot 需作为 StatisticSlot 的前驱节点，否则 Sentinel 运行会出现 bug。但你可以将 DegradeSlot 放在 FlowSlot 的前面，这就是我们上一篇说到的 ProcessorSlot 的排序。</p>

<h3 id="总结">总结</h3>

<p>Sentinel 使用 Java SPI 为我们提供了插件的功能，也类似于 Spring Boot 提供的自动配置类注册功能。我们可以直接替换 Sentinel 提供的默认 SlotChainBuilder，使用自定义的 SlotChainBuilder 自己构造 ProcessorSlotChain，以此实现修改 ProcessorSlot 排序顺序以及增加或移除 ProcessorSlot。在 Sentinel 1.7.2 版本中，Sentinel 支持使用 SPI 注册 ProcessorSlot，并且支持排序。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c6aaaaaafff2f7f7f6f186a1aba7afaae8a5a9ab" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1688021ca5f667',t:'MTczNDA5ODk1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=16&#32;Sentinel&#32;动态数据源：规则动态配置>
        <link rel="icon" href="/static/favicon.png">
        <title>16 Sentinel 动态数据源：规则动态配置 </title>
        
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
                            <h1 id="title" data-id="16 Sentinel 动态数据源：规则动态配置" class="title">16 Sentinel 动态数据源：规则动态配置</h1>
                            <div><p>经过前面的学习，我们知道，为资源配置各种规则可使用 Sentinel 提供的各种规则对应的 loadRules API，但这种以编码的方式配置规则很难实现动态修改。但基于 Sentinel 提供的各种规则对应的 loadRules API，我们可以自己实现规则的动态更新，而这一功能几乎在每个需要使用 Sentinel 的微服务项目中都需要实现一遍。Sentinel 也考虑到了这点，所以提供了动态数据源接口，并且提供了多种动态数据源的实现，尽管我们可能不会用到。</p>

<p>动态数据源作为扩展功能放在 sentinel-extension 模块下，前面我们学习的热点参数限流模块 sentinel-parameter-flow-control 也是在该模块下。在 1.7.1 版本，sentinel-extension 模块下的子模块除 sentinel-parameter-flow-control、sentinel-annotation-aspectj 之外，其余子模块都是实现动态数据源的模块。</p>

<ul>
<li>sentinel-datasource-extension：定义动态数据源接口、提供抽象类</li>
<li>sentinel-datasource-redis：基于 Redis 实现的动态数据源</li>
<li>sentinel-datasource-zookeeper： 基于 ZooKeeper 实现的动态数据源</li>
<li>其它省略</li>
</ul>

<p>显然，sentinel-datasource-extension 模块才是我们主要研究的模块，这是 Sentinel 实现动态数据源的核心。</p>

<h3 id="sentinelproperty">SentinelProperty</h3>

<p>SentinelProperty 是 Sentinel 提供的一个接口，可注册到 Sentinel 提供的各种规则的 Manager，例如 FlowRuleManager，并且可以给 SentinelProperty 添加监听器，在配置改变时，你可以调用 SentinelProperty#updateValue 方法，由它负责调用监听器去更新规则，而不需要调用 FlowRuleManager#loadRules 方法。同时，你也可以注册额外的监听器，在配置改变时做些别的事情。</p>

<p>SentinelProperty 并非 sentinel-datasource-extension 模块中定义的接口，而是 sentinel-core 定义的接口，其源码如下：</p>

<pre><code class="language-java">public interface SentinelProperty&lt;T&gt; {
    void addListener(PropertyListener&lt;T&gt; listener);
    void removeListener(PropertyListener&lt;T&gt; listener);
    boolean updateValue(T newValue);
}

</code></pre>

<ul>
<li>addListener：添加监听器</li>
<li>removeListener：移除监听器</li>
<li>updateValue：通知所有监听器配置更新，参数 newValue 为新的配置</li>
</ul>

<p>默认使用的实现类是 DynamicSentinelProperty，其实现源码如下（有删减）：</p>

<pre><code class="language-java">public class DynamicSentinelProperty&lt;T&gt; implements SentinelProperty&lt;T&gt; {
    // 存储注册的监听器
    protected Set&lt;PropertyListener&lt;T&gt;&gt; listeners = Collections.synchronizedSet(new HashSet&lt;PropertyListener&lt;T&gt;&gt;());
    @Override
    public void addListener(PropertyListener&lt;T&gt; listener) {
        listeners.add(listener);
        listener.configLoad(value);
    }
    @Override
    public void removeListener(PropertyListener&lt;T&gt; listener) {
        listeners.remove(listener);
    }
    @Override
    public boolean updateValue(T newValue) {
        for (PropertyListener&lt;T&gt; listener : listeners) {
            listener.configUpdate(newValue);
        }
        return true;
    }
}

</code></pre>

<p>可见，DynamicSentinelProperty 使用 Set 存储已注册的监听器，updateValue 负责通知所有监听器，调用监听器的 configUpdate 方法。</p>

<p>在前面分析 FlowRuleManager 时，我们只关注了其 loadRules 方法，除了使用 loadRules 方法加载规则配置之外，FlowRuleManager 还提供 registerProperty API，用于注册 SentinelProperty。</p>

<p>使用 SentinelProperty 实现加载 FlowRule 的步骤如下：</p>

<ol>
<li>给 FlowRuleManager 注册一个 SentinelProperty，替换 FlowRuleManager 默认创建的 SentinelProperty（因为默认的 SentinelProperty 外部拿不到）；</li>
<li>这一步由 FlowRuleManager 完成，FlowRuleManager 会给 SentinelProperty 注册 FlowPropertyListener 监听器，该监听器负责更新 FlowRuleManager.flowRules 缓存的限流规则；</li>
<li>在应用启动或者规则配置改变时，只需要调用 SentinelProperty#updateValue 方法，由 updateValue 通知 FlowPropertyListener 监听器去更新规则。</li>
</ol>

<p>FlowRuleManager 支持使用 SentinelProperty 加载或更新限流规则的实现源码如下：</p>

<pre><code class="language-java">public class FlowRuleManager {
    // 缓存限流规则
    private static final Map&lt;String, List&lt;FlowRule&gt;&gt; flowRules = new ConcurrentHashMap&lt;String, List&lt;FlowRule&gt;&gt;();
    // PropertyListener 监听器
    private static final FlowPropertyListener LISTENER = new FlowPropertyListener();
    // SentinelProperty
    private static SentinelProperty&lt;List&lt;FlowRule&gt;&gt; currentProperty 
       // 提供默认的 SentinelProperty
       = new DynamicSentinelProperty&lt;List&lt;FlowRule&gt;&gt;();

  static {
       // 给默认的 SentinelProperty 注册监听器（FlowPropertyListener）
        currentProperty.addListener(LISTENER);
  }

  // 注册 SentinelProperty
  public static void register2Property(SentinelProperty&lt;List&lt;FlowRule&gt;&gt; property) {
        synchronized (LISTENER) {
            currentProperty.removeListener(LISTENER);
            // 注册监听器
            property.addListener(LISTENER);
            currentProperty = property;
        }
  }
}

</code></pre>

<p>实现更新限流规则缓存的 FlowPropertyListener 是 FlowRuleManager 的一个内部类，其源码如下：</p>

<pre><code class="language-java">private static final class FlowPropertyListener implements PropertyListener&lt;List&lt;FlowRule&gt;&gt; {
        @Override
        public void configUpdate(List&lt;FlowRule&gt; value) {
            Map&lt;String, List&lt;FlowRule&gt;&gt; rules = FlowRuleUtil.buildFlowRuleMap(value);
            if (rules != null) {
                // 先清空缓存再写入
                flowRules.clear();
                flowRules.putAll(rules);
            }
        }
        @Override
        public void configLoad(List&lt;FlowRule&gt; conf) {
            Map&lt;String, List&lt;FlowRule&gt;&gt; rules = FlowRuleUtil.buildFlowRuleMap(conf);
            if (rules != null) {
                flowRules.clear();
                flowRules.putAll(rules);
            }
        }
}

</code></pre>

<p>PropertyListener 接口定义的两个方法：</p>

<ul>
<li>configUpdate：在规则更新时被调用，被调用的时机就是在 SentinelProperty#updateValue 方法被调用时。</li>
<li>configLoad：在规则首次加载时被调用，是否会被调用由 SentinelProperty 决定。DynamicSentinelProperty 就没有调用这个方法。</li>
</ul>

<p>所以，现在我们有两种方法更新限流规则：</p>

<ul>
<li>调用 FlowRuleManager#loadRules 方法</li>
<li>注册 SentinelProperty，调用 SentinelProperty#updateValue 方法</li>
</ul>

<h3 id="readabledatasource">ReadableDataSource</h3>

<p>Sentinel 将读和写数据源抽离成两个接口，一开始只有读接口，写接口是后面才加的功能，目前来看，写接口只在热点参数限流模块中使用到。事实上，使用读接口就已经满足我们的需求。ReadableDataSource 接口的定义如下：</p>

<pre><code class="language-java">public interface ReadableDataSource&lt;S, T&gt; {
    T loadConfig() throws Exception;
    S readSource() throws Exception;
    SentinelProperty&lt;T&gt; getProperty();
    void close() throws Exception;
}

</code></pre>

<p>ReadableDataSource 是一个泛型接口，参数化类型 S 代表用于装载从数据源读取的配置的类型，参数化类型 T 代表对应 Sentinel 中的规则类型。例如，我们可以定义一个 FlowRuleProps 类，用于装载从 yml 配置文件中读取的限流规则配置，然后再将 FlowRuleProps 转为 FlowRule，所以 S 可以替换为 FlowRuleProps，T 可以替换为 <code>List&lt;FlowRule&gt;</code>。</p>

<p>ReadableDataSource 接口定义的方法解释说明如下：</p>

<ul>
<li>loadConfig：加载配置。</li>
<li>readSource：从数据源读取配置，数据源可以是 yaml 配置文件，可以是 MySQL、Redis 等，由实现类决定从哪种数据源读取配置。</li>
<li>getProperty：获取 SentinelProperty。</li>
<li>close：用于关闭数据源，例如使用文件存储配置时，可在此方法实现关闭文件输入流等。</li>
</ul>

<p>如果动态数据源提供 SentinelProperty，则可以调用 getProperty 方法获取动态数据源的 SentinelProperty，将 SentinelProperty 注册给规则管理器（XxxManager），动态数据源在读取到配置时就可以调用自身 SentinelProperty 的 updateValue 方法通知规则管理器（XxxManager）更新规则。</p>

<p>AbstractDataSource 是一个抽象类，该类实现 ReadableDataSource 接口，用于简化具体动态数据源的实现，子类只需要继承 AbstractDataSource 并实现 readSource 方法即可。AbstractDataSource 源码如下：</p>

<pre><code class="language-java">public abstract class AbstractDataSource&lt;S, T&gt; implements ReadableDataSource&lt;S, T&gt; {
    protected final Converter&lt;S, T&gt; parser;
    protected final SentinelProperty&lt;T&gt; property;

    public AbstractDataSource(Converter&lt;S, T&gt; parser) {
        if (parser == null) {
            throw new IllegalArgumentException(&quot;parser can't be null&quot;);
        }
        this.parser = parser;
        this.property = new DynamicSentinelProperty&lt;T&gt;();
    }

    @Override
    public T loadConfig() throws Exception {
        return loadConfig(readSource());
    }

    public T loadConfig(S conf) throws Exception {
        T value = parser.convert(conf);
        return value;
    }

    @Override
    public SentinelProperty&lt;T&gt; getProperty() {
        return property;
    }
}

</code></pre>

<p>从源码可以看出：</p>

<ul>
<li>AbstractDataSource 要求所有子类都必须提供一个数据转换器（Converter），Converter 用于将 S 类型的对象转为 T 类型对象，例如将 FlowRuleProps 对象转为 FlowRule 集合。</li>
<li>AbstractDataSource 在构造方法中创建 DynamicSentinelProperty，因此子类无需创建 SentinelProperty。</li>
<li>AbstractDataSource 实现 loadConfig 方法，首先调用子类实现的 readSource 方法从数据源读取配置，返回的对象类型为 S，再调用 Converter#convert 方法，将对象类型由 S 转为 T。</li>
</ul>

<p>Converter 接口的定义如下：</p>

<pre><code class="language-java">public interface Converter&lt;S, T&gt; {
    T convert(S source);
}

</code></pre>

<ul>
<li>convert：将类型为 S 的参数 source 转为类型为 T 的对象。</li>
</ul>

<h3 id="基于-spring-cloud-动态配置实现规则动态配置">基于 Spring Cloud 动态配置实现规则动态配置</h3>

<p>我们项目中并未使用 Sentinel 提供的任何一种动态数据源的实现，而是选择自己实现数据源，因为我们项目是部署在 Kubernetes 集群上的，我们可以利用 ConfigMap 资源存储限流、熔断降级等规则。Spring Cloud Kubernetes 提供了 Spring Cloud 动态配置接口的实现，因此，我们不需要关心怎么去读取 ConfigMap 资源。就相当于基于 Spring Cloud 动态配置实现 Sentinel 规则动态数据源。Spring Cloud 动态配置如何使用这里不做介绍。</p>

<p>以实现 FlowRule 动态配置为例，其步骤如下。</p>

<p>第一步：定义一个用于装载动态配置的类，如 FlowRuleProps 所示。</p>

<pre><code class="language-java">@Component
@RefreshScope
@ConfigurationProperties(prefix = &quot;sentinel.flow-rules&quot;)
public class FlowRuleProps{
  //....省略
}

</code></pre>

<p>第二步：创建数据转换器，实现将 FlowRuleProps 转为 <code>List&lt;FlowRule&gt;</code>，如 FlowRuleConverter 所示。</p>

<pre><code class="language-java">public class FlowRuleConverter implements Converter&lt;FlowRuleProps, List&lt;FlowRule&gt;&gt;{

    @Override
    public List&lt;FlowRule&gt; convert(FlowRuleProps source){
       //....省略
    }
}

</code></pre>

<p>第三步：创建 FlowRuleDataSource，继承 AbstractDataSource，实现 readSource 方法。readSource 方法只需要获取 FlowRuleProps 返回即可，代码如下。</p>

<pre><code class="language-java">@Component
public class FlowRuleDataSource extends AbstractDataSource&lt;FlowRuleProps, List&lt;FlowRule&gt;&gt;{
    @Autowired
    private FlowRuleProps flowRuleProps;

    public FlowRuleDataSource() {
        super(new FlowRuleConverter());
    }
    @Override
    public FlowRuleProps readSource() throws Exception {
        return this.flowRuleProps;
    }
    @Override
    public void close() throws Exception {
    }
}

</code></pre>

<p>第四步：增强 FlowRuleDataSource，让 FlowRuleDataSource 能够监听到 FlowRuleProps 配置改变。</p>

<pre><code class="language-java">@Component
public class FlowRuleDataSource extends AbstractDataSource&lt;FlowRuleProps, List&lt;FlowRule&gt;&gt;
   implements ApplicationListener&lt;RefreshScopeRefreshedEvent&gt;,
        InitializingBean{
       // .....
    @Override
    public void onApplicationEvent(RefreshScopeRefreshedEvent event) {
        getProperty().updateValue(loadConfig());
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        onApplicationEvent(new RefreshScopeRefreshedEvent());
    }
}

</code></pre>

<ul>
<li>实现 InitializingBean 方法，在数据源对象创建时，初始化加载一次规则配置。</li>
<li>实现 ApplicationListener 接口，监听动态配置改变事件（RefreshScopeRefreshedEvent），在监听到事件时，首先调用 loadConfig 方法加载所有限流规则配置，然后调用 getProperty 方法获取 SentinelProperty，最后调用 SentinelProperty#updateValue 方法通知 FlowRuleManager 的监听器更新限流规则缓存。</li>
</ul>

<p>第五步：写一个 ApplicationRunner 类，在 Spring 容器刷新完成后， 将数据源（FlowRuleDataSource）的 SentinelProperty 注册给 FlowRuleManager，代码如下。</p>

<pre><code class="language-java">@Component
public class FlowRuleDataSourceConfiguration implements ApplicationRunner{
    @Autowired
    private FlowRuleDataSource flowRuleDataSource;

    @Override
    public void run(ApplicationArguments args) throws Exception {
        // 将数据源的 SentinelProperty 注册给 FlowRuleManager
        FlowRuleManager.register2Property(flowRuleDataSource.getProperty());
    }
}

</code></pre>

<p>在调用 FlowRuleManager#register2Property 方法将 FlowRuleDataSource 动态数据源的 SentinelProperty 注册给 FlowRuleManager 时，FlowRuleManager 会自动给该 SentinelProperty 注册一个监听器（FlowPropertyListener）。</p>

<p>到此，一个基于 Spring Cloud 动态配置的限流规则动态数据源就已经完成，整个调用链路如下：</p>

<ol>
<li>当动态配置改变时，Spring Cloud 会发出 RefreshScopeRefreshedEvent 事件，FlowRuleDataSource 的 onApplicationEvent 方法被调用；</li>
<li>FlowRuleDataSource 调用 loadConfig 方法获取最新的配置；</li>
<li>FlowRuleDataSource#loadConfig 调用 readSource 方法获取 FlowRuleProps 对象，此时的 FlowRuleProps 对象已经装载了最新的配置；</li>
<li>FlowRuleDataSource#loadConfig 调用转换器（FlowRuleConverter）的 convert 方法将 FlowRuleProps 对象转为 <code>List&lt;FlowRule&gt;</code> 对象；</li>
<li>FlowRuleDataSource 调用自身的 SentinelProperty 的 updateValue 方法通知所有监听器，并携带新的规则配置；</li>
<li>FlowPropertyListener 监听器的 configUpdate 方法被调用；</li>
<li>FlowPropertyListener 在 configUpdate 方法中更新 FlowRuleManager 缓存的限流规则。</li>
</ol>

<h3 id="总结">总结</h3>

<p>了解 Sentinel 实现动态数据源的原理后，我们可以灵活地自定义规则动态数据源，例如本篇介绍的，利用 Kubernetes 的 ConfigMap 资源存储规则配置，并通过 Spring Cloud 动态配置实现 Sentinel 的规则动态数据源。不仅如此，Sentinel 实现动态数据源的整体框架的设计也是值得我们学习的，如数据转换器、监听器。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#bfd3d3d3868b8e8e8f88ffd8d2ded6d391dcd0d2" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1689e01a31f667',t:'MTczNDA5OTAzNi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>
## 分词规则:

```Go
types.EngineOpts{
	Using:         4,
	GseMode: true,
}
```

- 当 Using 为 0 并且 Content 分词不为空时, 从内容通过 gse分词和 Tokens 中得到关键词

- 当 Using 为 1 并且 Content 分词不为空时，优先从内容通过 gse分词中得到关键词

- 当 Using 为 2 、Using 为 1 或 3 并且 Content 分词为空时, 从 Tokens 中得到关键词

- 当 Using 为 3 并且 Content 分词不为空时, 从内容通过 gse分词和通过`""`分词中得到关键词

- 当 Using 为 4 并且 Content 分词不为空时，从内容通过空格分词中得到关键词

- 当 GseMode 为 true, gse 分词采用搜索模式, 给搜索引擎提供尽可能多的关键字
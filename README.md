# Dev.to Prometheus Exporter

Export basic stats about all articles written by the owner of the api_key provided.

Requires `DEVTO_API_KEY` env var.

Returns 

```
# HELP devto_article_comments_total Displays total number of comments to users articles
# TYPE devto_article_comments_total gauge
devto_article_comments_total{article_id="<article id>"} <int>
# HELP devto_article_reactions_total Displays total number of reactions to users articles
# TYPE devto_article_reactions_total gauge
devto_article_reactions_total{article_id="<article id>"} <int>
# HELP devto_article_views_total Displays total number of views of users articles
# TYPE devto_article_views_total gauge
devto_article_views_total{article_id="<article id>"} <int>
# HELP devto_probe_duration_seconds Displays how long the probe took in seconds
# TYPE devto_probe_duration_seconds gauge
devto_probe_duration_seconds <float64>
# HELP devto_probe_success Displays whether or not the probe was a success
# TYPE devto_probe_success gauge
devto_probe_success < 1|0 >
# HELP devto_published_articles_total Displays the total number of users published articles
# TYPE devto_published_articles_total gauge
devto_published_articles_total <int>
```
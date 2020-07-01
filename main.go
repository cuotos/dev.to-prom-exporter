package main

import (
	"context"
	"dev.to-prom-exporter/client"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

type App struct {
	DevtoApiKey string `required:"true" split_words:"true"`
}

func init() {
	prometheus.MustRegister(NewCollector("devto_exporter"))
}

func main() {

	var a App
	err := envconfig.Process("", &a)
	if err != nil {
		log.Fatal(err)
	}

	client := &client.DevtoClient{ApiKey: a.DevtoApiKey}

	http.HandleFunc("/probe", func(writer http.ResponseWriter, request *http.Request) {
		probeHandler(writer, request, client)
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}

func probeHandler(w http.ResponseWriter, r *http.Request, client *client.DevtoClient) {

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*60)
	defer cancel()

	r.WithContext(ctx)

	probeSuccessGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "devto",
		Name:      "probe_success",
		Help:      "Displays whether or not the probe was a success",
	})

	registry := prometheus.NewRegistry()

	registry.MustRegister(probeSuccessGauge)

	success := Probe(registry, client)
	if success {
		probeSuccessGauge.Set(1)
	}

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	h.ServeHTTP(w, r)
}

func Probe(registry *prometheus.Registry, client *client.DevtoClient) (success bool) {

	var (
		totalArticlesGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "devto",
			Name:      "published_articles_total",
			Help:      "Displays the total number of users published articles",
		}, []string{"user"})

		totalViewsGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "devto",
			Name:      "article_views_total",
			Help:      "Displays total number of views of users articles",
		}, []string{"article_id", "user"})

		totalReactionGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "devto",
			Name:      "article_reactions_total",
			Help:      "Displays total number of reactions to users articles",
		}, []string{"article_id", "user"})

		totalCommentsGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "devto",
			Name:      "article_comments_total",
			Help:      "Displays total number of comments to users articles",
		}, []string{"article_id", "user"})

		probeDuration = prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "devto",
			Name:      "probe_duration_seconds",
			Help:      "Displays how long the probe took in seconds",
		})
	)

	registry.MustRegister(totalArticlesGauge)
	registry.MustRegister(totalViewsGauge)
	registry.MustRegister(totalReactionGauge)
	registry.MustRegister(totalCommentsGauge)
	registry.MustRegister(probeDuration)

	start := time.Now()
	articles, err := client.GetUserArticleData()
	if err != nil {
		log.Println("unable to get data:", err.Error())
		return false
	}
	user, err := client.GetUserDetails()
	if err != nil {
		log.Println("unable to get data:", err.Error())
		return false
	}
	username := user.Username

	duration := time.Since(start).Seconds()
	probeDuration.Set(duration)

	var totalPublishedArticles int

	for _, a := range articles {
		id := strconv.Itoa(a.ID)

		totalViewsGauge.WithLabelValues(id,username).Set(float64(a.PageViewsCount))
		totalReactionGauge.WithLabelValues(id,username).Set(float64(a.PublicReactionsCount))
		totalCommentsGauge.WithLabelValues(id, username).Set(float64(a.CommentsCount))

		if a.Published {
			totalPublishedArticles++
		}
	}

	totalArticlesGauge.WithLabelValues(username).Set(float64(totalPublishedArticles))

	return true
}

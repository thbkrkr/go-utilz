package elastic

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
)

// NewClient creates an Elasticsearch client
func NewClient(elasticURL string) *elastic.Client {
	opts := elastic.SetURL(elasticURL)
	client, err := elastic.NewClient(opts)
	if err != nil {
		logrus.WithError(err).WithField("url", elasticURL).Fatal("Fail to create elastic client")
	}

	logrus.WithField("url", elasticURL).Info("Elastic client ready")

	return client
}

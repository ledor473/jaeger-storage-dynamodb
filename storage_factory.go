package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jaegertracing/jaeger/plugin"
	"github.com/jaegertracing/jaeger/storage"
	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
)

type DynamoDBStorageFactory struct {
	svc            *dynamodb.DynamoDB
	metricsFactory metrics.Factory
	logger         *zap.Logger
	tableName      string
}

func (d *DynamoDBStorageFactory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
	d.metricsFactory = metricsFactory
	d.logger = logger

	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	d.svc = dynamodb.New(sess)

	return nil
}

// CreateSpanReader creates a spanstore.Reader.
func (d *DynamoDBStorageFactory) CreateSpanReader() (spanstore.Reader, error) {
	d.logger.Info("Creating SpanReader for DynamoDB with table=" + d.tableName)
	return nil, nil
}

// CreateSpanWriter creates a spanstore.Writer.
func (d *DynamoDBStorageFactory) CreateSpanWriter() (spanstore.Writer, error) {
	d.logger.Info("Creating SpanWrtier for DynamoDB with table=" + d.tableName)
	return nil, nil
}

// CreateDependencyReader creates a dependencystore.Reader.
func (d *DynamoDBStorageFactory) CreateDependencyReader() (dependencystore.Reader, error) {
	d.logger.Info("Creating DependencyReader for DynamoDB with table=" + d.tableName)
	return nil, nil
}

func (d *DynamoDBStorageFactory) AddFlags(flagSet *flag.FlagSet) {
	flagSet.Int("dynamodb.table-name", 30, "DynamoDB Table Name")
}

func (d *DynamoDBStorageFactory) InitFromViper(v *viper.Viper) {
	d.tableName =  v.GetString("dynamodb.table-name")
}

var factory = DynamoDBStorageFactory{}

// Export the needed symbols
var Configurable plugin.Configurable = &factory
var StorageFactory storage.Factory = &factory;
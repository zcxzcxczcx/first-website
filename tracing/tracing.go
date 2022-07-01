package tracing

import (
	"context"
	"first_website/config"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/apex/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var jaegerConfig = config.GetJaegerConfig()

// 空的span，为了避免nil pointer出错
var emptySpan opentracing.Span

var tracer, closer = newOpentracingTracer()
var enabled = false

func newOpentracingTracer() (opentracing.Tracer, io.Closer) {
	emptySpan = opentracing.NoopTracer{}.StartSpan("noop")
	if len(jaegerConfig.CollectorEndpoint) == 0 {
		return opentracing.NoopTracer{}, nil
	}
	cfg := jaegercfg.Configuration{
		ServiceName: "onlinehall",

		Reporter: &jaegercfg.ReporterConfig{
			QueueSize:         2,
			CollectorEndpoint: jaegerConfig.CollectorEndpoint,
		},
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1.0,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Error(err.Error())
		return nil, nil
	}
	enabled = true
	return tracer, closer
}

func StartSpan(c *gin.Context) opentracing.Span {
	carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
	var span opentracing.Span
	name := c.Request.RequestURI
	clientContext, err := tracer.Extract(opentracing.HTTPHeaders, carrier)
	if err == nil {
		span = tracer.StartSpan(name, ext.RPCServerOption(clientContext))
	} else {
		span = tracer.StartSpan(name)
	}
	_ = opentracing.ContextWithSpan(c.Request.Context(), span)

	return span
}

func GetSpan(ctx context.Context) opentracing.Span {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		return span
	}
	return emptySpan
}

func SetTags(ctx context.Context, tags map[string]interface{}) opentracing.Span {
	span := GetSpan(ctx)
	for k, v := range tags {
		span.SetTag(k, v)
	}
	return span
}

func StartChildSpan(ctx context.Context, serviceName string) (opentracing.Span, context.Context) {
	parent := opentracing.SpanFromContext(ctx)
	if parent == nil {
		return emptySpan, ctx
	}
	spanCtx := parent.Context()
	span := tracer.StartSpan(serviceName, opentracing.ChildOf(spanCtx))
	ctx = opentracing.ContextWithSpan(ctx, span)

	return span, ctx
}

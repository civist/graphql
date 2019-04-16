package graphql_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/location"
	"github.com/graphql-go/graphql/testutil"
)

func tinit(t *testing.T) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Type",
			Fields: graphql.Fields{
				"a": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "foo", nil
					},
				},
			},
		}),
	})
	if err != nil {
		t.Fatalf("Error in schema %v", err.Error())
	}
	return schema
}

func TestExtensionInitPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.initFn = func(ctx context.Context, p *graphql.Params) context.Context {
		if true {
			panic(errors.New("test error"))
		}
		return ctx
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.Init: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionParseDidStartPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.parseDidStartFn = func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(err error) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ParseDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionParseFinishFuncPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.parseDidStartFn = func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
		return ctx, func(err error) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ParseFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionValidationDidStartPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.validationDidStartFn = func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func([]gqlerrors.FormattedError) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ValidationDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionValidationFinishFuncPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.validationDidStartFn = func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
		return ctx, func([]gqlerrors.FormattedError) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ValidationFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionExecutionDidStartPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.executionDidStartFn = func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(r *graphql.Result) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ExecutionDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionExecutionFinishFuncPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.executionDidStartFn = func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
		return ctx, func(r *graphql.Result) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ExecutionFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionResolveFieldDidStartPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(v interface{}, err error) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ResolveFieldDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionResolveFieldFinishFuncPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		return ctx, func(v interface{}, err error) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ResolveFieldFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionGetResultPanic(t *testing.T) {
	ext := newtestExt("testExt")
	ext.getResultFn = func(context.Context) interface{} {
		if true {
			panic(errors.New("test error"))
		}
		return nil
	}
	ext.hasResultFn = func() bool {
		return true
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.GetResult: %v", ext.Name(), errors.New("test error"))),
		},
		Extensions: make(map[string]interface{}),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionContextPropagation(t *testing.T) {
	testType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Test",
		Fields: graphql.Fields{
			"foo": &graphql.Field{Type: graphql.String},
			"bar": &graphql.Field{Type: graphql.String},
			"baz": &graphql.Field{Type: graphql.String},
		},
	})
	type test struct {
		Foo string
		Bar string
		Baz string
	}
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"a": &graphql.Field{
					Type: testType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return test{"foo", "bar", "baz"}, nil
					},
				},
				"b": &graphql.Field{
					Type: graphql.NewList(testType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return []test{{"foo", "bar", "baz"}, {"foo", "bar", "baz"}}, nil
					},
				},
			},
		}),
	})
	if err != nil {
		t.Fatal("Error in schema", err)
	}

	type ctxPathKey struct{}
	ctxPathTrace := map[string]int{}

	ext := newtestExt("testExt")
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		ctxPath := ctx.Value(ctxPathKey{}).([]string)
		ctxPath = append(append([]string(nil), ctxPath...), i.FieldName)
		ctx = context.WithValue(ctx, ctxPathKey{}, ctxPath)
		ctxPathTrace[strings.Join(ctxPath, ".")] += 1
		return ctx, func(interface{}, error) {}
	}

	query := `query { a { foo bar baz } b { foo bar baz } }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.Background(), ctxPathKey{}, []string{}),
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": map[string]interface{}{
				"foo": "foo",
				"bar": "bar",
				"baz": "baz",
			},
			"b": []interface{}{
				map[string]interface{}{
					"foo": "foo",
					"bar": "bar",
					"baz": "baz",
				},
				map[string]interface{}{
					"foo": "foo",
					"bar": "bar",
					"baz": "baz",
				},
			},
		},
	}
	expectedCtxPathTrace := map[string]int{
		"a":     1,
		"a.bar": 1,
		"a.baz": 1,
		"a.foo": 1,
		"b":     1,
		"b.bar": 2,
		"b.baz": 2,
		"b.foo": 2,
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}

	if !reflect.DeepEqual(expectedCtxPathTrace, ctxPathTrace) {
		t.Fatalf("Unexpected ctx path trace, Diff: %v", testutil.Diff(expectedCtxPathTrace, ctxPathTrace))
	}
}

func TestExtensionErrorPropagation(t *testing.T) {
	testType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Test",
		Fields: graphql.Fields{
			"foo": &graphql.Field{Type: graphql.String},
			"bar": &graphql.Field{Type: graphql.String},
			"baz": &graphql.Field{Type: graphql.String},
		},
	})
	type test struct {
		Foo string
		Bar string
		Baz string
	}
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"a": &graphql.Field{
					Type: testType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return nil, errors.New("test error")
					},
				},
			},
		}),
	})
	if err != nil {
		t.Fatal("Error in schema", err)
	}

	type ctxKey struct{}

	ext := newtestExt("testExt")
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		ctx = context.WithValue(ctx, ctxKey{}, "some value")
		return ctx, func(interface{}, error) {}
	}

	query := `query { a { foo bar baz } }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.Background(),
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": nil,
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(&gqlerrors.Error{
				Message: "test error",
				Locations: []location.SourceLocation{
					{Line: 1, Column: 9},
				},
				Path:          []interface{}{"a"},
				OriginalError: errors.New("test error"),
			}),
		},
	}

	if !testutil.EqualResults(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func newtestExt(name string) *testExt {
	ext := &testExt{
		name: name,
	}
	if ext.initFn == nil {
		ext.initFn = func(ctx context.Context, p *graphql.Params) context.Context {
			return ctx
		}
	}
	if ext.parseDidStartFn == nil {
		ext.parseDidStartFn = func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
			return ctx, func(err error) {

			}
		}
	}
	if ext.validationDidStartFn == nil {
		ext.validationDidStartFn = func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
			return ctx, func([]gqlerrors.FormattedError) {

			}
		}
	}
	if ext.executionDidStartFn == nil {
		ext.executionDidStartFn = func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
			return ctx, func(r *graphql.Result) {

			}
		}
	}
	if ext.resolveFieldDidStartFn == nil {
		ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
			return ctx, func(v interface{}, err error) {

			}
		}
	}
	if ext.hasResultFn == nil {
		ext.hasResultFn = func() bool {
			return false
		}
	}
	if ext.getResultFn == nil {
		ext.getResultFn = func(context.Context) interface{} {
			return nil
		}
	}
	return ext
}

type testExt struct {
	name                   string
	initFn                 func(ctx context.Context, p *graphql.Params) context.Context
	hasResultFn            func() bool
	getResultFn            func(context.Context) interface{}
	parseDidStartFn        func(ctx context.Context) (context.Context, graphql.ParseFinishFunc)
	validationDidStartFn   func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc)
	executionDidStartFn    func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc)
	resolveFieldDidStartFn func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc)
}

func (t *testExt) Init(ctx context.Context, p *graphql.Params) context.Context {
	return t.initFn(ctx, p)
}

func (t *testExt) Name() string {
	return t.name
}

func (t *testExt) HasResult() bool {
	return t.hasResultFn()
}

func (t *testExt) GetResult(ctx context.Context) interface{} {
	return t.getResultFn(ctx)
}

func (t *testExt) ParseDidStart(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
	return t.parseDidStartFn(ctx)
}

func (t *testExt) ValidationDidStart(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
	return t.validationDidStartFn(ctx)
}

func (t *testExt) ExecutionDidStart(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
	return t.executionDidStartFn(ctx)
}

func (t *testExt) ResolveFieldDidStart(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
	return t.resolveFieldDidStartFn(ctx, i)
}

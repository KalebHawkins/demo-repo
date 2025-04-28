package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Tags from stackforge.yaml
func applyGeneratedTags(stack awscdk.Stack) {
	awscdk.Tags_Of(stack).Add(jsii.String("AppOwner"), jsii.String("platform-team"), nil)
	awscdk.Tags_Of(stack).Add(jsii.String("Environment"), jsii.String("staging"), nil)
	awscdk.Tags_Of(stack).Add(jsii.String("Project"), jsii.String("project-abc"), nil)
}

func resolveRuntime(runtime string) awslambda.Runtime {
	switch runtime {
	case "go":
		return awslambda.Runtime_GO_1_X()
	case "python3.11":
		return awslambda.Runtime_PYTHON_3_11()
	case "python3.9":
		return awslambda.Runtime_PYTHON_3_9()
	case "python3.8":
		return awslambda.Runtime_PYTHON_3_8()
	default:
		panic("Unsupported Lambda runtime: " + runtime)
	}
}

func NewLambdaStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)
	applyGeneratedTags(stack)
	awslambda.NewFunction(stack, jsii.String("processUsers"), &awslambda.FunctionProps{
		FunctionName: jsii.String("processUsers"),
		Runtime:      resolveRuntime("python3.11"),
		Handler:      jsii.String("handler.lambda_handler"),
		Code:         awslambda.Code_FromAsset(jsii.String("../services/processUsers"), nil),
		MemorySize:   jsii.Number(256),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"STAGE": jsii.String("staging"),
		},
	})
	awslambda.NewFunction(stack, jsii.String("notifyAdmins"), &awslambda.FunctionProps{
		FunctionName: jsii.String("notifyAdmins"),
		Runtime:      resolveRuntime("python3.9"),
		Handler:      jsii.String("handler.lambda_handler"),
		Code:         awslambda.Code_FromAsset(jsii.String("../services/notifyAdmins"), nil),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(5)),
	})

	return stack
}

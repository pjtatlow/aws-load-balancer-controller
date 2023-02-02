package framework

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/go-logr/logr"
	elbv2api "github.com/pjtatlow/aws-load-balancer-controller/apis/elbv2/v1beta1"
	"github.com/pjtatlow/aws-load-balancer-controller/pkg/aws"
	"github.com/pjtatlow/aws-load-balancer-controller/pkg/aws/throttle"
	"github.com/pjtatlow/aws-load-balancer-controller/test/framework/controller"
	"github.com/pjtatlow/aws-load-balancer-controller/test/framework/helm"
	"github.com/pjtatlow/aws-load-balancer-controller/test/framework/http"
	awsresources "github.com/pjtatlow/aws-load-balancer-controller/test/framework/resources/aws"
	k8sresources "github.com/pjtatlow/aws-load-balancer-controller/test/framework/resources/k8s"
	"github.com/pjtatlow/aws-load-balancer-controller/test/framework/utils"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Framework struct {
	Options   Options
	RestCfg   *rest.Config
	K8sClient client.Client
	Cloud     aws.Cloud

	CTRLInstallationManager controller.InstallationManager
	NSManager               k8sresources.NamespaceManager
	DPManager               k8sresources.DeploymentManager
	SVCManager              k8sresources.ServiceManager
	INGManager              k8sresources.IngressManager
	LBManager               awsresources.LoadBalancerManager
	TGManager               awsresources.TargetGroupManager

	HTTPVerifier http.Verifier

	Logger         logr.Logger
	LoggerReporter httpexpect.LoggerReporter
}

func InitFramework() (*Framework, error) {
	err := globalOptions.Validate()
	if err != nil {
		return nil, err
	}
	restCfg := ctrl.GetConfigOrDie()

	k8sSchema := runtime.NewScheme()
	clientgoscheme.AddToScheme(k8sSchema)
	elbv2api.AddToScheme(k8sSchema)

	k8sClient, err := client.New(restCfg, client.Options{Scheme: k8sSchema})
	if err != nil {
		return nil, err
	}

	cloud, err := aws.NewCloud(aws.CloudConfig{
		Region:         globalOptions.AWSRegion,
		VpcID:          globalOptions.AWSVPCID,
		MaxRetries:     3,
		ThrottleConfig: throttle.NewDefaultServiceOperationsThrottleConfig(),
	}, nil)
	if err != nil {
		return nil, err
	}

	logger, loggerReporter := utils.NewGinkgoLogger()

	f := &Framework{
		Options:   globalOptions,
		RestCfg:   restCfg,
		K8sClient: k8sClient,
		Cloud:     cloud,

		CTRLInstallationManager: buildControllerInstallationManager(globalOptions, logger),
		NSManager:               k8sresources.NewDefaultNamespaceManager(k8sClient, logger),
		DPManager:               k8sresources.NewDefaultDeploymentManager(k8sClient, logger),
		SVCManager:              k8sresources.NewDefaultServiceManager(k8sClient, logger),
		INGManager:              k8sresources.NewDefaultIngressManager(k8sClient, logger),
		LBManager:               awsresources.NewDefaultLoadBalancerManager(cloud.ELBV2(), logger),
		TGManager:               awsresources.NewDefaultTargetGroupManager(cloud.ELBV2(), logger),

		HTTPVerifier: http.NewDefaultVerifier(),

		Logger:         logger,
		LoggerReporter: loggerReporter,
	}

	return f, nil
}

func buildControllerInstallationManager(options Options, logger logr.Logger) controller.InstallationManager {
	helmReleaseManager := helm.NewDefaultReleaseManager(options.KubeConfig, logger)
	ctrlInstallationManager := controller.NewDefaultInstallationManager(helmReleaseManager, options.ClusterName, options.AWSRegion, options.AWSVPCID, options.HelmChart, logger)
	return ctrlInstallationManager
}

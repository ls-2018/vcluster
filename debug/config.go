package debug

import (
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetHostConfig() *restclient.Config {
	// creates the connection
	inClusterConfig, _ := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "121.config"))
	return inClusterConfig
}
func GetVirtualConfig() *restclient.Config {
	// creates the connection
	inClusterConfig, _ := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "vcluster.config"))
	return inClusterConfig
}

func GetVirtualConfigPath() string {
	return filepath.Join(homedir.HomeDir(), ".kube", "vcluster.config")
}

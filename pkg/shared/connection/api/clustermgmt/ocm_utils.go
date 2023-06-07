package clustermgmt

import (
	"github.com/apicurio/apicurio-cli/pkg/core/localize"
	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func clustermgmtConnection(f *factory.Factory, accessToken string, clustermgmturl string) (*v1.Client, func(), error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, nil, err
	}
	client, closeConnection, err := conn.API().OCMClustermgmt(clustermgmturl, accessToken)
	if err != nil {
		return nil, nil, err
	}
	return client, closeConnection, nil
}

func GetClusterList(f *factory.Factory, accessToken string, clustermgmturl string, pageNumber int, pageLimit int) (*v1.ClusterList, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()

	resource := client.Clusters().List()
	resource = resource.Page(pageNumber)
	resource = resource.Size(pageLimit)
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	clusters := response.Items()
	return clusters, nil
}

func GetClusterById(f *factory.Factory, accessToken string, clustermgmturl string, clusterId string) (*v1.Cluster, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()

	resource := client.Clusters().Cluster(clusterId)
	response, err := resource.Get().Send()
	if err != nil {
		return nil, err
	}
	cluster := response.Body()
	return cluster, nil
}

func GetMachinePoolList(f *factory.Factory, clustermgmturl string, accessToken string, clusterId string) (*v1.MachinePoolsListResponse, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	resource := client.Clusters().Cluster(clusterId).MachinePools().List()
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetClusterListWithSearchParams(f *factory.Factory, clustermgmturl string, accessToken string, params string, page int, size int) (*v1.ClusterList, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	resource := client.Clusters().List().Search(params).Size(size).Page(page)
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	return response.Items(), nil
}

func CreateMachinePool(f *factory.Factory, clustermgmturl string, accessToken string, mprequest *v1.MachinePool, clusterId string) (*v1.MachinePool, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	response, err := client.Clusters().Cluster(clusterId).MachinePools().Add().Body(mprequest).Send()
	if err != nil {
		return nil, err
	}
	return response.Body(), nil
}

func GetMachinePoolIdByTaintKey(f *factory.Factory, clustermgmturl string, accessToken string, clusterId string, machinePoolTaintKey string) (string, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return "", err
	}
	defer closeConnection()
	response, err := client.Clusters().Cluster(clusterId).MachinePools().List().Send()
	if err != nil {
		return "", err
	}
	machinePools := response.Items()
	for _, machinePool := range machinePools.Slice() {
		for _, taint := range machinePool.Taints() {
			if taint.Key() == machinePoolTaintKey {
				return machinePool.ID(), nil
			}
		}
	}
	return "", nil
}

func DeleteMachinePool(f *factory.Factory, clustermgmturl string, accessToken string, clusterId string, machinePoolId string) error {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return err
	}
	defer closeConnection()
	_, err = client.Clusters().Cluster(clusterId).MachinePools().MachinePool(machinePoolId).Delete().Send()
	if err != nil {
		return err
	}
	return nil
}

func GetMachinePoolNodeCount(machinePool *v1.MachinePool) int {
	var nodeCount int
	replicas, ok := machinePool.GetReplicas()
	if ok {
		nodeCount = replicas
	} else {
		autoscaledReplicas, ok := machinePool.GetAutoscaling()
		if ok {
			nodeCount = autoscaledReplicas.MaxReplicas()
		}
	}
	return nodeCount
}

func RemoveAddonsFromCluster(f *factory.Factory, clusterManagementApiUrl string, accessToken string, cluster *clustersmgmtv1.Cluster, addonList []string) error {
	// create a new addon via ocm
	conn, err := f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt(clusterManagementApiUrl, accessToken)
	if err != nil {
		return err
	}
	defer cc()

	addons, err := client.Clusters().Cluster(cluster.ID()).Addons().List().Send()
	if err != nil {
		return err
	}

	for _, addonToDelete := range addonList {
		for i := 0; i < addons.Size(); i++ {
			addon := addons.Items().Get(i)

			if addon.ID() == addonToDelete {
				f.Logger.Info(f.Localizer.MustLocalize("kafka.openshiftCluster.common.addons.deleting.message", localize.NewEntry("Id", addon.ID())))
				_, err = client.Clusters().Cluster(cluster.ID()).Addons().Addoninstallation(addon.ID()).Delete().Send()
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

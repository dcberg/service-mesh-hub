package get_cluster_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	mock_table_printing "github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing/mocks"
	cli_mocks "github.com/solo-io/service-mesh-hub/cli/pkg/mocks"
	cli_test "github.com/solo-io/service-mesh-hub/cli/pkg/test"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	mock_zephyr_discovery "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Get Cluster Cmd", func() {
	var (
		ctrl               *gomock.Controller
		ctx                context.Context
		meshctl            *cli_test.MockMeshctl
		mockKubeLoader     *cli_mocks.MockKubeLoader
		mockClusterPrinter *mock_table_printing.MockKubernetesClusterPrinter
		mockClusterClient  *mock_zephyr_discovery.MockKubernetesClusterClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockKubeLoader = cli_mocks.NewMockKubeLoader(ctrl)
		mockClusterPrinter = mock_table_printing.NewMockKubernetesClusterPrinter(ctrl)
		mockClusterClient = mock_zephyr_discovery.NewMockKubernetesClusterClient(ctrl)
		meshctl = &cli_test.MockMeshctl{
			MockController: ctrl,
			Ctx:            ctx,
			Clients:        common.Clients{},
			KubeClients: common.KubeClients{
				KubeClusterClient: mockClusterClient,
			},
			KubeLoader: mockKubeLoader,
			Printers: common.Printers{
				KubeClusterPrinter: mockClusterPrinter,
			},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will call the kube cluster printer with the proper data", func() {

		clusters := []*zephyr_discovery.KubernetesCluster{
			{
				ObjectMeta: v1.ObjectMeta{
					Name: "mesh-1",
				},
			},
			{
				ObjectMeta: v1.ObjectMeta{
					Name: "mesh-2",
				},
			},
		}
		mockKubeLoader.EXPECT().
			GetRestConfigForContext("", "").
			Return(nil, nil)
		mockClusterClient.EXPECT().
			ListKubernetesCluster(ctx).
			Return(&zephyr_discovery.KubernetesClusterList{
				Items: []zephyr_discovery.KubernetesCluster{*clusters[0], *clusters[1]},
			}, nil)
		mockClusterPrinter.EXPECT().
			Print(gomock.Any(), clusters).
			Return(nil)
		_, err := meshctl.Invoke("get clusters")
		Expect(err).NotTo(HaveOccurred())
	})
})

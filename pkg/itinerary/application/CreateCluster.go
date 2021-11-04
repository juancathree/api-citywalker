package application

import (
	matrixDomain "citywalker/pkg/matrixCost/domain"
	placeDomain "citywalker/pkg/place/domain"
	"math"
	"math/rand"
)

type Cluster struct {
	NClusters    int
	CentroID     [][]float64
	Clusters     [][]placeDomain.Place
	ClusterAlloc [][]float64
}

func (c *Cluster) CreateCluster(visit *[]placeDomain.Place, matrixCost *matrixDomain.MatrixCost) *[][]placeDomain.Place {
	// Calculate num of clusters
	c.NClusters = int(math.Round(math.Sqrt(float64(len(*visit))))) + 1

	c.InitializeClusters(len(*visit), matrixCost)

	change := false

	for !change {
		c.FillCluster(visit, &change)
		if !change {
			break
		}
		c.RedefineCentroID()
		change = false
	}

	filteredCluster := make([][]placeDomain.Place, 0, len(c.Clusters))

	for _, v := range c.Clusters {
		if len(v) > 0 {
			filteredCluster = append(filteredCluster, v)
		}
	}

	return &filteredCluster
}

func (c *Cluster) InitializeClusters(lenVisit int, matrixCost *matrixDomain.MatrixCost) {
	// center of clusters
	c.CentroID = make([][]float64, c.NClusters)

	// cluster to store places
	c.Clusters = make([][]placeDomain.Place, c.NClusters)

	// cost of each cluster
	c.ClusterAlloc = make([][]float64, lenVisit)

	for i := 0; i < c.NClusters; i++ {
		c.CentroID[i] = make([]float64, 2)
		c.Clusters[i] = make([]placeDomain.Place, 0, lenVisit)
		c.CentroID[i][1] = (*matrixCost).MinLat*1.001 + rand.Float64()*((*matrixCost).MaxLat*0.999-(*matrixCost).MinLat*1.001)
		c.CentroID[i][0] = (*matrixCost).MinLong*1.001 + rand.Float64()*((*matrixCost).MaxLong*0.999-(*matrixCost).MinLong*1.001)
	}

	for i := 0; i < lenVisit; i++ {
		c.ClusterAlloc[i] = make([]float64, c.NClusters+1)
		c.ClusterAlloc[i][c.NClusters] = 100.0
	}
}

func (c *Cluster) FillCluster(visit *[]placeDomain.Place, change *bool) {
	for i, p := range *visit {
		cl := 0
		min := 100.0
		var j int
		for j = 0; j < c.NClusters; j++ {
			c.ClusterAlloc[i][j] = p.CalcDistance(c.CentroID[j][1], c.CentroID[j][0])
			if c.ClusterAlloc[i][j] < min {
				min = c.ClusterAlloc[i][j]
				cl = j
			}
		}
		if c.ClusterAlloc[i][j] != float64(cl) {
			c.ClusterAlloc[i][j] = float64(cl)
			*change = true
		}
		c.Clusters[cl] = append(c.Clusters[cl], p)
	}
}

func (c *Cluster) RedefineCentroID() {

	for i := range c.Clusters {
		maxLong, maxLat := 0.0, 0.0
		minLong, minLat := 100.0, 100.0
		for _, p := range c.Clusters[i] {
			lat := p.Latitude()
			long := p.Longitude()
			if lat > maxLat {
				maxLat = lat
			}
			if lat < minLat {
				minLat = lat
			}
			if long > maxLong {
				maxLong = long
			}
			if long < minLong {
				minLong = long
			}
		}
		c.Clusters[i] = c.Clusters[i][:0]
		c.CentroID[i][1] = minLat + rand.Float64()*(maxLat-minLat)
		c.CentroID[i][0] = minLong + rand.Float64()*(maxLong-minLong)
	}
}

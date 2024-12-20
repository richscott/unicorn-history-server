package model

import (
	"github.com/G-Research/yunikorn-core/pkg/webservice/dao"
)

type AppHistory struct {
	Metadata                      `json:",inline"`
	ID                            string `json:"id"`
	dao.ApplicationHistoryDAOInfo `json:",inline"`
}

func (h *AppHistory) MergeFromAppHistory(other *dao.ApplicationHistoryDAOInfo) {
	h.ApplicationHistoryDAOInfo = *other
}

type ContainerHistory struct {
	Metadata                    `json:",inline"`
	ID                          string `json:"id"`
	dao.ContainerHistoryDAOInfo `json:",inline"`
}

func (h *ContainerHistory) MergeFromContainerHistory(other *dao.ContainerHistoryDAOInfo) {
	h.ContainerHistoryDAOInfo = *other
}

package test

import (
	m "recon_test/model"
	svc "recon_test/service"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconService_PerformSuccess(t *testing.T) {
	recon := []m.ReconResult{
		{
			ID          : "bbbb",
			Amount      : "101",
			Description : "B",
			Date        : "2021-06-30",
			Remark      : m.AMOUNT_DIFF,
		},
	}
}
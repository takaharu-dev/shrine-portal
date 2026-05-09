package service

import "shrine-portal/backend/internal/response"

type BenefitService struct{}

func NewBenefitService() *BenefitService {
	return &BenefitService{}
}

func (s *BenefitService) FindAll() []response.BenefitResponse {
	return []response.BenefitResponse{
		{
			ID:   1,
			Name: "金運",
		},
		{
			ID:   2,
			Name: "縁結び",
		},
		{
			ID:   3,
			Name: "学業成就",
		},
		{
			ID:   4,
			Name: "健康",
		},
		{
			ID:   5,
			Name: "厄除け",
		},
		{
			ID:   6,
			Name: "仕事運",
		},
	}
}

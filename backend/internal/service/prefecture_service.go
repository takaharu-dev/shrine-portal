package service

import "shrine-portal/backend/internal/response"

type PrefectureService struct{}

func NewPrefectureService() *PrefectureService {
	return &PrefectureService{}
}

func (s *PrefectureService) FindAll() []response.PrefectureResponse {
	return []response.PrefectureResponse{
		{
			ID:   13,
			Name: "東京都",
		},
		{
			ID:   14,
			Name: "神奈川県",
		},
	}
}

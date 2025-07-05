// Package server implements all of the gRPC related server request handling.
package server

import (
	"context"

	tokenValidator "github.com/Aeriqu/kanikaki/common/token"
	"github.com/Aeriqu/kanikaki/services/srs/database"
	srspb "github.com/Aeriqu/kanikaki/services/srs/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SrsServer struct {
	srspb.UnimplementedSrsServer
	*database.Database
}

func (server *SrsServer) GetAllKanjiReviewData(req *srspb.IdentifierRequest, stream srspb.Srs_GetAllKanjiReviewDataServer) error {
	// Check auth token validity
	claims, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token provided")
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return status.Error(codes.Internal, "error retrieving subject from claims")
	}
	if subject != req.Identifier {
		return status.Error(codes.PermissionDenied, "token subject does not match identifier")
	}
	
	reviewData, err := server.Database.GetUserReviewData(req.Identifier, "kanji")

	if err != nil {
		return err
	}

	for _, review := range *reviewData {
		response := &srspb.KanjiReviewDataResponse{
			Character: review.Topic,
			SrsLevel: review.SrsLevel,
			ReviewTime: review.ReviewTime,
		}
		if err := stream.Send(response); err !=  nil {
			return err
		}
	}

	return nil
}

func (server *SrsServer) GetDueKanjiReviews(req *srspb.IdentifierRequest, stream srspb.Srs_GetDueKanjiReviewsServer) error {
	// Check auth token validity
	claims, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token provided")
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return status.Error(codes.Internal, "error retrieving subject from claims")
	}
	if subject != req.Identifier {
		return status.Error(codes.PermissionDenied, "token subject does not match identifier")
	}

	reviewData, err := server.Database.GetDueReviewData(req.Identifier, "kanji")
	if err != nil {
		return err
	}

	for _, review := range *reviewData {
		response := &srspb.KanjiReviewDataResponse{
			Character: review.Topic,
			SrsLevel: review.SrsLevel,
			ReviewTime: review.ReviewTime,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (server *SrsServer) UpdateKanjiReviewData(ctx context.Context, req *srspb.KanjiReviewDataRequest) (*srspb.KanjiReviewDataResponse, error) {
	// Check auth token validity
	claims, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token provided")
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return nil, status.Error(codes.Internal, "error retrieving subject from claims")
	}
	if subject != req.Identifier {
		return nil, status.Error(codes.PermissionDenied, "token subject does not match identifier")
	}

	_, err = server.Database.UpdateSrsItem(req.Identifier, "kanji", req.Character, req.SrsLevel, req.ReviewTime)
	if err != nil {
		return nil, err
	}

	response := &srspb.KanjiReviewDataResponse{
		Character: req.Character,
		SrsLevel: req.SrsLevel,
		ReviewTime: req.ReviewTime,
	}

	return response, nil
}
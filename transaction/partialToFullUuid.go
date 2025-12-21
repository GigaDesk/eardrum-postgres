package transaction

import (
	customErrors "github.com/GigaDesk/eardrum-interfaces/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func PartialToFullUuid(db *gorm.DB, partialScan string, inputPin string, checkPIN func(hashedPIN, PIN string) error) (uuid.UUID, *customErrors.PublicError) {
	if len(partialScan) < 10 {
		return uuid.Nil, NewAccountNotFoundError("Invalid QR code. Please scan again")
	}

	// 1. Fetch candidates WITH their similarity scores
	type candidate struct {
		QrCode  uuid.UUID
		PinCode string
		Score   float32 `gorm:"column:score"`
	}
	var results []candidate

	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Exec("SET LOCAL pg_trgm.word_similarity_threshold = 0.6")
		// We explicitly SELECT the similarity score for logging purposes
		return tx.Raw(`
			SELECT qr_code, pin_code, word_similarity(?, qr_code::text) as score 
			FROM users 
			WHERE ? <% qr_code::text 
			ORDER BY score DESC`, partialScan, partialScan).Scan(&results).Error
	})

	if err != nil {
		return uuid.Nil, ErrDBPersistenceFailure(err)
	}

	// LOG: Initial database results
	dbLog := log.Info().Int("count", len(results)).Str("partial_scan", partialScan)
	candidatesLog := []map[string]interface{}{}
	for _, r := range results {
		candidatesLog = append(candidatesLog, map[string]interface{}{
			"qr": r.QrCode.String(), "pin": r.PinCode, "score": r.Score,
		})
	}
	dbLog.Interface("db_candidates", candidatesLog).Msg("Fuzzy scan database results")

	// 2. Mapping and Disqualification logic
	finalMap := make(map[string]candidate) // Changed to store struct to keep score
	disqualified := make(map[string]candidate)
	subsequencePassed := []map[string]interface{}{}

	for _, res := range results {
		if isStrictSubsequence(partialScan, res.QrCode.String()) {
			subsequencePassed = append(subsequencePassed, map[string]interface{}{
				"qr": res.QrCode.String(), "pin": res.PinCode, "score": res.Score,
			})

			if _, exists := disqualified[res.PinCode]; exists {
				continue
			}

			if _, exists := finalMap[res.PinCode]; exists {
				disqualified[res.PinCode] = res
				delete(finalMap, res.PinCode)
				continue
			}

			finalMap[res.PinCode] = res
		}
	}

	// LOG: Subsequence and Disqualification
	log.Info().
		Interface("subsequence_passed", subsequencePassed).
		Interface("disqualified_pins", disqualified).
		Msg("Filter results")

	// 3. Final Verification
	for hashedPin, cand := range finalMap {
		if err1 := checkPIN(hashedPin, inputPin); err1 == nil {
			// LOG: Success
			log.Info().
				Str("qr", cand.QrCode.String()).
				Float32("score", cand.Score).
				Msg("Match found successfully")
			return cand.QrCode, nil
		}
	}

	// LOG: Failure
	log.Info().Str("input", partialScan).Msg("Account not found after fuzzy check")
	return uuid.Nil, NewAccountNotFoundError("Invalid QR code. Please scan again")
}

// isStrictSubsequence checks if partial is a subset of full in the correct order
func isStrictSubsequence(partial, full string) bool {
	pIdx, fIdx := 0, 0
	for pIdx < len(partial) && fIdx < len(full) {
		if partial[pIdx] == full[fIdx] {
			pIdx++
		}
		fIdx++
	}
	return pIdx == len(partial)
}

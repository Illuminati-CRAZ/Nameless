package scores

type invalidScoreDetections int

const (
	detectPassNoReplayData invalidScoreDetections = iota
	detectFailWithReplayData
	detectReplayDecodeError
	detectInvalidReplayMD5
	detectInvalidExecutingAssemblyMD5
	detectInvalidEntryAssemblyMD5
	detectInvalidMapMD5
	detectInvalidMapReplayMD5
	detectInvalidTotalScore
	detectMaxTotalScoreWithFailure
	detectInvalidAudioPlaybackRate
	detectInvalidMaxComboForJudgements
	detectMaxComboAndEndMismatch
	detectFailWithNonZeroHealth
	detectPassWithZeroHealth
)

package restriction

const (
	ValidateRestrictedStatusIneligible = "INELIGIBLE"

	REBlacklisted    = "blacklisted"
	REKYCNotVerified = "kyc_not_verified"
	REKYCPending     = "kyc_pending"
	REKYCUnderage    = "kyc_underage"
)

// GetRestrictedCategoryMessage will generate message for validation restriction category
func (resp *ValidateRestrictionResponse) GetRestrictedCategoryMessage(lang string, minAge int) string {
	for _, data := range resp.DataResponse {
		if data.Status == ValidateRestrictedStatusIneligible {
			if data.Action != nil {
				for _, action := range data.Action {
					switch action {
					case REBlacklisted:
						return ""
					case REKYCNotVerified:
						return ""
					case REKYCPending:
						return ""
					case REKYCUnderage:
						return ""
					}
				}
			}
		}
	}
	return ""
}

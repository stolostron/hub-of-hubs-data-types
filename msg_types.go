package datatypes

const (
	PolicyType = "PoliciesBundle"
	PlacementRuleType = "PlacementRulesBundle"
	PlacementBindingType = "PlacementBindingsBundle"
)

func GetBundleTypes() []string {
	var bundleTypes []string
	bundleTypes = append(bundleTypes, PolicyType)
	bundleTypes = append(bundleTypes, PlacementRuleType)
	bundleTypes = append(bundleTypes, PlacementBindingType)
	return bundleTypes
}
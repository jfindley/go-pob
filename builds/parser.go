package builds

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Vilsol/go-pob/pob"
)

var nilCleanupRegex = regexp.MustCompile(`\w+?="nil"`)

var pobGemIDtoGameGemIDs = map[string]string{
	"Metadata/Items/Gems/Smite":                   "Metadata/Items/Gems/SkillGemSmite",
	"Metadata/Items/Gems/ConsecratedPath":         "Metadata/Items/Gems/SkillGemConsecratedPath",
	"Metadata/Items/Gems/VaalAncestralWarchief":   "Metadata/Items/Gems/SkillGemVaalAncestralWarchief",
	"Metadata/Items/Gems/HeraldOfAgony":           "Metadata/Items/Gems/SkillGemHeraldOfAgony",
	"Metadata/Items/Gems/HeraldOfPurity":          "Metadata/Items/Gems/SkillGemHeraldOfPurity",
	"Metadata/Items/Gems/ScourgeArrow":            "Metadata/Items/Gems/SkillGemScourgeArrow",
	"Metadata/Items/Gems/RainOfSpores":            "Metadata/Items/Gems/SkillGemToxicRain",
	"Metadata/Items/Gems/SummonRelic":             "Metadata/Items/Gems/SkillGemSummonRelic",
	"Metadata/Items/Gems/SkillGemNewArcticArmour": "Metadata/Items/Gems/SkillGemArcticArmour",
}

func ParseBuildStr(rawXML string) (*pob.PathOfBuilding, error) {
	return ParseBuild([]byte(rawXML))
}

func ParseBuild(rawXML []byte) (*pob.PathOfBuilding, error) {
	clean := nilCleanupRegex.ReplaceAllLiteral(rawXML, []byte{})
	var build pob.PathOfBuilding
	err := xml.Unmarshal(clean, &build)
	if err != nil {
		return nil, fmt.Errorf("failed to parse build as xml: %w", err)
	}

	for i, set := range build.Skills.SkillSets {
		for j, skill := range set.Skills {
			for k, gem := range skill.Gems {
				if gameID, ok := pobGemIDtoGameGemIDs[gem.GemID]; ok {
					build.Skills.SkillSets[i].Skills[j].Gems[k].GemID = gameID
				}
			}
		}
	}

	build.Build.PassiveNodes = make([]int64, 0, 100)
	// TODO: This only grabs the activeSpec but there are many here
	var spec = build.Tree.Specs[build.Tree.ActiveSpec-1]
	var nodeStrs = strings.Split(spec.NodesAttr, ",")
	for _, str := range nodeStrs {
		var num, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("spec has some non-integer nodes: %s", spec.NodesAttr)
		}
		build.Build.PassiveNodes = append(build.Build.PassiveNodes, num)
	}

	return &build, nil
}

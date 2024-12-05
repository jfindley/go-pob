package moddb

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/Vilsol/go-pob/mod"
	"github.com/Vilsol/go-pob/utils"
)

func TestList(t *testing.T) {
	tc := []struct {
		name        string
		mods        []mod.Mod
		cfg         *ListCfg
		mappedNames []string
		expected    any
	}{
		{
			name:        "single mod, empty modlist",
			mods:        []mod.Mod{mod.NewList("testMod", "testVal")},
			mappedNames: []string{"testMod"},
			expected:    []any{"testVal"},
		},
		{
			name: "multiple list mods, empty modlist",
			mods: []mod.Mod{
				mod.NewList("testMod0", "testVal0"),
				mod.NewList("testMod1", "testVal1"),
				mod.NewList("testMod2", "testVal2"),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: []any{
				"testVal0", "testVal1", "testVal2",
			},
		},
		{
			name: "mixed mods, empty modlist",
			mods: []mod.Mod{
				mod.NewList("testMod0", "testVal0"),
				mod.NewFlag("testMod1", true),
				mod.NewFloat("testMod2", mod.TypeIncrease, 42.42),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: []any{
				"testVal0",
			},
		},
		{
			name: "multiple list mods, keyword modlist",
			mods: []mod.Mod{
				mod.NewList("testMod0", "testVal0").KeywordFlag(mod.KeywordFlagCold),
				mod.NewList("testMod1", "testVal1").KeywordFlag(mod.KeywordFlagFire),
				mod.NewList("testMod2", "testVal2").KeywordFlag(mod.KeywordFlagFire),
			},
			cfg: &ListCfg{
				KeywordFlags: utils.Ptr(mod.KeywordFlagCold),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: []any{
				"testVal0",
			},
		},
		{
			name: "multiple mods, flag modlist",
			mods: []mod.Mod{
				mod.NewList("testMod0", "testVal0").Flag(mod.MFlagProjectile),
				mod.NewList("testMod1", "testVal1").Flag(mod.MFlagAilment),
				mod.NewList("testMod2", "testVal2").Flag(mod.MFlagDot),
			},
			cfg: &ListCfg{
				Flags: utils.Ptr(mod.MFlagDot),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: []any{
				"testVal2",
			},
		},
		{
			name: "multiple mods, source modlist",
			mods: []mod.Mod{
				mod.NewList("testMod0", "testVal0").Source(mod.SourceShock),
				mod.NewList("testMod1", "testVal1").Source(mod.SourceScorch),
				mod.NewList("testMod2", "testVal2").Source(mod.SourceBrittle),
			},
			cfg: &ListCfg{
				Source: utils.Ptr(mod.SourceScorch),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: []any{
				"testVal1",
			},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			m := NewModList()
			for _, tm := range test.mods {
				m.AddMod(tm)
			}
			got := m.List(test.cfg, test.mappedNames...)
			testza.AssertEqual(t, test.expected, got)
			t.Run(test.name+"-parent", func(t *testing.T) {
				p := NewModList()
				for _, tm := range test.mods {
					p.AddMod(tm)
				}
				m := NewModList()
				m.Parent = p
				got := m.List(test.cfg, test.mappedNames...)
				testza.AssertEqual(t, test.expected, got)
			})
		})
	}

}

func TestSum(t *testing.T) {
	tc := []struct {
		name        string
		mods        []mod.Mod
		modType     mod.Type
		cfg         *ListCfg
		mappedNames []string
		expected    float64
	}{
		{
			name:        "single mod, empty modlist",
			mods:        []mod.Mod{mod.NewFloat("testMod", mod.TypeIncrease, 42)},
			modType:     mod.TypeIncrease,
			mappedNames: []string{"testMod"},
			expected:    42,
		},
		{
			name: "multiple mod types, empty modlist",
			mods: []mod.Mod{
				mod.NewFloat("testMod0", mod.TypeIncrease, 10),
				mod.NewFloat("testMod1", mod.TypeIncrease, 20),
				mod.NewFloat("testMod2", mod.TypeMultiplier, 40),
			},
			modType: mod.TypeIncrease,
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: 30,
		},
		{
			name: "multiple mod types, keyword modlist",
			mods: []mod.Mod{
				mod.NewFloat("testMod0", mod.TypeIncrease, 10).KeywordFlag(mod.KeywordFlagCold),
				mod.NewFloat("testMod1", mod.TypeIncrease, 20).KeywordFlag(mod.KeywordFlagFire),
				mod.NewFloat("testMod2", mod.TypeMultiplier, 40).KeywordFlag(mod.KeywordFlagFire),
			},
			cfg: &ListCfg{
				KeywordFlags: utils.Ptr(mod.KeywordFlagFire),
			},
			modType: mod.TypeIncrease,
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: 20,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			m := NewModList()
			for _, tm := range test.mods {
				m.AddMod(tm)
			}
			got := m.Sum(test.modType, test.cfg, test.mappedNames...)
			testza.AssertEqual(t, test.expected, got)
			t.Run(test.name+"-parent", func(t *testing.T) {
				p := NewModList()
				for _, tm := range test.mods {
					p.AddMod(tm)
				}
				m := NewModList()
				m.Parent = p
				got := m.Sum(test.modType, test.cfg, test.mappedNames...)
				testza.AssertEqual(t, test.expected, got)
			})
		})
	}

}

func TestMore(t *testing.T) {
	tc := []struct {
		name        string
		mods        []mod.Mod
		cfg         *ListCfg
		mappedNames []string
		expected    float64
	}{
		{
			name:        "single mod, empty modlist",
			mods:        []mod.Mod{mod.NewFloat("testMod", mod.TypeMore, 120)},
			mappedNames: []string{"testMod"},
			expected:    2.2,
		},
		{
			name: "multiple mod types, keyword modlist",
			mods: []mod.Mod{
				mod.NewFloat("testMod0", mod.TypeMore, 100).KeywordFlag(mod.KeywordFlagCold),
				mod.NewFloat("testMod1", mod.TypeMore, 200),
				mod.NewFloat("testMod2", mod.TypeMore, 400).KeywordFlag(mod.KeywordFlagCold),
			},
			cfg: &ListCfg{
				KeywordFlags: utils.Ptr(mod.KeywordFlagCold),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: 30,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			m := NewModList()
			for _, tm := range test.mods {
				m.AddMod(tm)
			}
			got := m.More(test.cfg, test.mappedNames...)
			testza.AssertEqual(t, test.expected, got)
			t.Run(test.name+"-parent", func(t *testing.T) {
				p := NewModList()
				for _, tm := range test.mods {
					p.AddMod(tm)
				}
				m := NewModList()
				m.Parent = p
				got := m.More(test.cfg, test.mappedNames...)
				testza.AssertEqual(t, test.expected, got)
			})
		})
	}

}

func TestFlag(t *testing.T) {
	tc := []struct {
		name        string
		mods        []mod.Mod
		cfg         *ListCfg
		mappedNames []string
		expected    bool
	}{
		{
			name:        "single mod, empty modlist",
			mods:        []mod.Mod{mod.NewFlag("testMod", true)},
			mappedNames: []string{"testMod"},
			expected:    true,
		},
		{
			name: "multiple mod types, keyword modlist",
			mods: []mod.Mod{
				mod.NewFlag("testMod0", true).KeywordFlag(mod.KeywordFlagCold),
				mod.NewFlag("testMod1", true),
				mod.NewFlag("testMod2", true).KeywordFlag(mod.KeywordFlagCold),
			},
			cfg: &ListCfg{
				KeywordFlags: utils.Ptr(mod.KeywordFlagCold),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: true,
		},
		{
			name:        "no matches, empty modlist",
			mods:        []mod.Mod{mod.NewFloat("testMod", mod.TypeIncrease, 32)},
			mappedNames: []string{"testMod"},
			expected:    false,
		},
		{
			name: "no matches, keyword modlist",
			mods: []mod.Mod{
				mod.NewFlag("testMod0", true).KeywordFlag(mod.KeywordFlagCold),
				mod.NewFlag("testMod1", true),
				mod.NewFlag("testMod2", true).KeywordFlag(mod.KeywordFlagCold),
			},
			cfg: &ListCfg{
				KeywordFlags: utils.Ptr(mod.KeywordFlagFire),
			},
			mappedNames: []string{
				"testMod0", "testMod1", "testMod2",
			},
			expected: true,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			m := NewModList()
			for _, tm := range test.mods {
				m.AddMod(tm)
			}
			got := m.Flag(test.cfg, test.mappedNames...)
			testza.AssertEqual(t, test.expected, got)
			t.Run(test.name+"-parent", func(t *testing.T) {
				p := NewModList()
				for _, tm := range test.mods {
					p.AddMod(tm)
				}
				m := NewModList()
				m.Parent = p
				got := m.Flag(test.cfg, test.mappedNames...)
				testza.AssertEqual(t, test.expected, got)
			})
		})
	}

}

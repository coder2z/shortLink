package test

import (
	"encoding/json"
	"fmt"
	string_plus "shortLink/pkg/string"
	"testing"
)

func TestJson(t *testing.T) {
	jsonStr := `{
    "result": {
        "195151": {
            "icon_url": "fWFc82js0fmoRAP-qOIPu5THSWqfSmTELLqcUywGkijVjZULUrsm1j-9xgEEZDgJUBrxhyxRjc3ZAfOeD-VOz45ss5JX22VqwgIvbbGwZ2cyIVKRA6VYBPY8rAvvUCJhupAyUIfjpvUWJ1tW2OBwwQ",
            "icon_url_large": "fWFc82js0fmoRAP-qOIPu5THSWqfSmTELLqcUywGkijVjZULUrsm1j-9xgEEZDgJUBrxhyxRjc3ZAfOeD-VOz45ss5JX22VqwgIvbbGwZ2cyIVKRA6VYBPY8rAvvUCJhupAyUIfjpvUWJ1tW2OBwwQ",
            "icon_drag_url": "",
            "name": "A Rather Festive Tree",
            "market_hash_name": "A Rather Festive Tree",
            "market_name": "A Rather Festive Tree",
            "name_color": "7D6D00",
            "background_color": "3C352E",
            "type": "Level 1 Hat",
            "tradable": "1",
            "marketable": "0",
            "commodity": "0",
            "market_tradable_restriction": "7",
            "market_marketable_restriction": "0",
            "fraudwarnings": "",
            "descriptions": "",
            "actions": {
                "0": {
                    "name": "Item Wiki Page...",
                    "link": "http://wiki.teamfortress.com/scripts/itemredirect.php?id=341&lang=en_US"
                },
                "1": {
                    "name": "Inspect in Game...",
                    "link": "steam://rungame/440/76561202255233023/+tf_econ_item_preview%20S%owner_steamid%A%assetid%D4614055367427583853"
                }
            },
            "market_actions": {
                "0": {
                    "name": "Inspect in Game...",
                    "link": "steam://rungame/440/76561202255233023/+tf_econ_item_preview%20M%listingid%A%assetid%D4614055367427583853"
                }
            },
            "tags": {
                "0": {
                    "internal_name": "Unique",
                    "name": "Unique",
                    "category": "Quality",
                    "color": "7D6D00",
                    "category_name": "Quality"
                },
                "1": {
                    "internal_name": "misc",
                    "name": "Cosmetic",
                    "category": "Type",
                    "category_name": "Type"
                },
                "2": {
                    "internal_name": "Scout",
                    "name": "Scout",
                    "category": "Class",
                    "category_name": "Class"
                },
                "3": {
                    "internal_name": "Sniper",
                    "name": "Sniper",
                    "category": "Class",
                    "category_name": "Class"
                },
                "4": {
                    "internal_name": "Soldier",
                    "name": "Soldier",
                    "category": "Class",
                    "category_name": "Class"
                },
                "5": {
                    "internal_name": "Demoman",
                    "name": "Demoman",
                    "category": "Class",
                    "category_name": "Class"
                },
                "6": {
                    "internal_name": "Medic",
                    "name": "Medic",
                    "category": "Class",
                    "category_name": "Class"
                },
                "7": {
                    "internal_name": "Heavy",
                    "name": "Heavy",
                    "category": "Class",
                    "category_name": "Class"
                },
                "8": {
                    "internal_name": "Pyro",
                    "name": "Pyro",
                    "category": "Class",
                    "category_name": "Class"
                },
                "9": {
                    "internal_name": "Spy",
                    "name": "Spy",
                    "category": "Class",
                    "category_name": "Class"
                },
                "10": {
                    "internal_name": "Engineer",
                    "name": "Engineer",
                    "category": "Class",
                    "category_name": "Class"
                }
            },
            "app_data": {
                "def_index": "341",
                "quality": "6",
                "slot": "Cosmetic",
                "filter_data": {
                    "931505789": {
                        "element_ids": {
                            "0": "991457757",
                            "1": "8"
                        }
                    },
                    "1662615936": {
                        "element_ids": {
                            "0": "991457757",
                            "1": "1",
                            "2": "2",
                            "3": "3",
                            "4": "4",
                            "5": "5",
                            "6": "6",
                            "7": "7",
                            "8": "8",
                            "9": "9"
                        }
                    }
                },
                "player_class_ids": {
                    "0": "1",
                    "1": "2",
                    "2": "3",
                    "3": "4",
                    "4": "5",
                    "5": "6",
                    "6": "7",
                    "7": "8",
                    "8": "9"
                },
                "highlight_color": "7a6e65"
            },
            "classid": "195151"
        },
        "16891096": {
            "icon_url": "fWFc82js0fmoRAP-qOIPu5THSWqfSmTELLqcUywGkijVjZULUrsm1j-9xgEMZAgCSSTjqyhGi9zZAfOeD-VOmIkw5sUFgWE_wwJ4Y7vgZTY-c1OTBaEIDqRirFC6XiI37pVlA9S39PUWJ1tHc_-5Qg",
            "icon_url_large": "fWFc82js0fmoRAP-qOIPu5THSWqfSmTELLqcUywGkijVjZULUrsm1j-9xgEMZAgCSSTjqyhGi9zZAfOeD-VOmIkw5sUFgWE_wwJ4Y7vgZTY-c1OTBaEIDqRirFC6XiI37pVlA9S39PUWJ1tHc_-5Qg",
            "icon_drag_url": "",
            "name": "Apparition's Aspect",
            "market_hash_name": "Apparition's Aspect",
            "market_name": "Apparition's Aspect",
            "name_color": "7D6D00",
            "background_color": "3C352E",
            "type": "Level 13 Mask",
            "tradable": "1",
            "marketable": "0",
            "commodity": "0",
            "market_tradable_restriction": "7",
            "market_marketable_restriction": "0",
            "fraudwarnings": "",
            "descriptions": "",
            "actions": {
                "0": {
                    "name": "Item Wiki Page...",
                    "link": "http://wiki.teamfortress.com/scripts/itemredirect.php?id=571&lang=en_US"
                },
                "1": {
                    "name": "Inspect in Game...",
                    "link": "steam://rungame/440/76561202255233023/+tf_econ_item_preview%20S%owner_steamid%A%assetid%D7973706962122134107"
                }
            },
            "market_actions": {
                "0": {
                    "name": "Inspect in Game...",
                    "link": "steam://rungame/440/76561202255233023/+tf_econ_item_preview%20M%listingid%A%assetid%D7973706962122134107"
                }
            },
            "tags": {
                "0": {
                    "internal_name": "Unique",
                    "name": "Unique",
                    "category": "Quality",
                    "color": "7D6D00",
                    "category_name": "Quality"
                },
                "1": {
                    "internal_name": "misc",
                    "name": "Cosmetic",
                    "category": "Type",
                    "category_name": "Type"
                },
                "2": {
                    "internal_name": "Pyro",
                    "name": "Pyro",
                    "category": "Class",
                    "category_name": "Class"
                }
            },
            "app_data": {
                "def_index": "571",
                "quality": "6",
                "slot": "Cosmetic",
                "filter_data": {
                    "931505789": {
                        "element_ids": {
                            "0": "991457757",
                            "1": "8"
                        }
                    },
                    "1662615936": {
                        "element_ids": {
                            "0": "991457757",
                            "1": "7"
                        }
                    }
                },
                "player_class_ids": {
                    "0": "7"
                },
                "highlight_color": "7a6e65"
            },
            "classid": "16891096"
        },
        "success": true
    }
}`

	var s map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &s)
	m := s["result"].(map[string]interface{})
	for _, i := range m {
		var a AutoGenerated
		m2, ok := i.(map[string]interface{})
		if ok {
			data, _ := json.Marshal(m2)
			json.Unmarshal(data, &a)
			fmt.Println(a)
		}

	}
}

type AutoGenerated struct {
	IconURL                     string `json:"icon_url"`
	IconURLLarge                string `json:"icon_url_large"`
	IconDragURL                 string `json:"icon_drag_url"`
	Name                        string `json:"name"`
	MarketHashName              string `json:"market_hash_name"`
	MarketName                  string `json:"market_name"`
	NameColor                   string `json:"name_color"`
	BackgroundColor             string `json:"background_color"`
	Type                        string `json:"type"`
	Tradable                    string `json:"tradable"`
	Marketable                  string `json:"marketable"`
	Commodity                   string `json:"commodity"`
	MarketTradableRestriction   string `json:"market_tradable_restriction"`
	MarketMarketableRestriction string `json:"market_marketable_restriction"`
	Fraudwarnings               string `json:"fraudwarnings"`
	Descriptions                string `json:"descriptions"`
	Actions                     struct {
		Num0 struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"0"`
		Num1 struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"1"`
	} `json:"actions"`
	MarketActions struct {
		Num0 struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"0"`
	} `json:"market_actions"`
	Tags struct {
		Num0 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			Color        string `json:"color"`
			CategoryName string `json:"category_name"`
		} `json:"0"`
		Num1 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"1"`
		Num2 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"2"`
		Num3 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"3"`
		Num4 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"4"`
		Num5 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"5"`
		Num6 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"6"`
		Num7 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"7"`
		Num8 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"8"`
		Num9 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"9"`
		Num10 struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"10"`
	} `json:"tags"`
	AppData struct {
		DefIndex   string `json:"def_index"`
		Quality    string `json:"quality"`
		Slot       string `json:"slot"`
		FilterData struct {
			Num931505789 struct {
				ElementIds struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
				} `json:"element_ids"`
			} `json:"931505789"`
			Num1662615936 struct {
				ElementIds struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"element_ids"`
			} `json:"1662615936"`
		} `json:"filter_data"`
		PlayerClassIds struct {
			Num0 string `json:"0"`
			Num1 string `json:"1"`
			Num2 string `json:"2"`
			Num3 string `json:"3"`
			Num4 string `json:"4"`
			Num5 string `json:"5"`
			Num6 string `json:"6"`
			Num7 string `json:"7"`
			Num8 string `json:"8"`
		} `json:"player_class_ids"`
		HighlightColor string `json:"highlight_color"`
	} `json:"app_data"`
	Classid string `json:"classid"`
}

func TestRandStr(t *testing.T) {
	string_plus.New()
	go func() {
		for {
			fmt.Println(string_plus.Get())
		}
	}()

	<-make(chan struct{})
}

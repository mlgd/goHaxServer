package wiiu

import (
	"github.com/astaxie/beego"
)

type SystemVersions struct {
	Name       string
	Value      string
	Constants  *constants
	LoaderName string
}

var (
	arraySystemVersions = []*SystemVersions{
		&SystemVersions{"US_5_5_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.12 NintendoBrowser/4.3.1.11264.US", Constants550, "stagefright.bin"},
		&SystemVersions{"US_5_5_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.11 NintendoBrowser/4.3.0.11224.US", Constants550, "stagefright.bin"},
		&SystemVersions{"US_5_4_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.9 NintendoBrowser/4.2.0.11146.US", Constants532, "stagefright.bin"},
		&SystemVersions{"US_5_3_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.US", Constants532, "stagefright.bin"},
		&SystemVersions{"US_5_3_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.US", nil, ""},
		&SystemVersions{"US_5_3_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/4.1.0.9584.US", nil, ""},
		&SystemVersions{"US_5_2_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.US", nil, ""},
		&SystemVersions{"US_5_1_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.US", nil, ""},
		&SystemVersions{"US_5_1_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.US", nil, ""},
		&SystemVersions{"US_5_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.US", nil, ""},
		&SystemVersions{"US_5_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.US", nil, ""},
		&SystemVersions{"US_4_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.6 NintendoBrowser/2.0.0.9362.US", nil, ""},
		&SystemVersions{"US_2_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.23 NintendoBrowser/1.1.0.7579.US", nil, ""},
		&SystemVersions{"US_2_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US", nil, ""},

		&SystemVersions{"EU_5_5_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.12 NintendoBrowser/4.3.1.11264.EU", Constants550, "stagefright.bin"},
		&SystemVersions{"EU_5_5_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.11 NintendoBrowser/4.3.0.11224.EU", Constants550, "stagefright.bin"},
		&SystemVersions{"EU_5_4_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.9 NintendoBrowser/4.2.0.11146.EU", Constants532, "stagefright.bin"},
		&SystemVersions{"EU_5_3_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.EU", Constants532, "stagefright.bin"},
		&SystemVersions{"EU_5_3_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.EU", nil, ""},
		&SystemVersions{"EU_5_3_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/4.1.0.9584.EU", nil, ""},
		&SystemVersions{"EU_5_2_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.EU", nil, ""},
		&SystemVersions{"EU_5_1_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.EU", nil, ""},
		&SystemVersions{"EU_5_1_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.EU", nil, ""},
		&SystemVersions{"EU_5_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.EU", nil, ""},
		&SystemVersions{"EU_5_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.EU", nil, ""},
		&SystemVersions{"EU_4_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.6 NintendoBrowser/2.0.0.9362.EU", nil, ""},
		&SystemVersions{"EU_2_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.23 NintendoBrowser/1.1.0.7579.EU", nil, ""},
		&SystemVersions{"EU_2_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.EU", nil, ""},

		&SystemVersions{"JP_5_4_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.9 NintendoBrowser/4.2.0.11146.JP", Constants532, "stagefright.bin"},
		&SystemVersions{"JP_5_3_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.JP", Constants532, "stagefright.bin"},
		&SystemVersions{"JP_5_3_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.15 NintendoBrowser/4.1.1.9601.JP", nil, ""},
		&SystemVersions{"JP_5_3_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/4.1.0.9584.JP", nil, ""},
		&SystemVersions{"JP_5_2_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.JP", nil, ""},
		&SystemVersions{"JP_5_1_2", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.JP", nil, ""},
		&SystemVersions{"JP_5_1_1", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.14 NintendoBrowser/3.1.1.9577.JP", nil, ""},
		&SystemVersions{"JP_5_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.JP", nil, ""},
		&SystemVersions{"JP_5_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.12 NintendoBrowser/3.0.0.9561.JP", nil, ""},
		&SystemVersions{"JP_4_0_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.6 NintendoBrowser/2.0.0.9362.JP", nil, ""},
		&SystemVersions{"JP_2_1_0", "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.23 NintendoBrowser/1.1.0.7579.JP", nil, ""},
	}
)

func NewSystemVersions() *SystemVersions {
	return &SystemVersions{}
}

func GetSystemVersion(propriety string) *SystemVersions {
	if propriety == "" {
		return nil
	}

	for _, sv := range arraySystemVersions {
		if sv.Value == propriety {
			return sv
		}
	}

	beego.Warning("Unknown system version:", propriety)

	return nil
}

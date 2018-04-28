package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
)

var (
	sDns   []string //短链接DNS
	lDns   []string //长链接DNS
	priKey []byte   //私钥
	pubKey []byte   //公钥
)

func main() {
	var bRet bool
	bRet, sDns, lDns = getDns()
	if !bRet {
		fmt.Println("提取长、短链接DNS出错！")
		return
	}
	fmt.Println(sDns)
	fmt.Println(lDns)

	bRet, priKey, pubKey = GenEcdhKey()
	if !bRet {
		fmt.Println("初始化出错！")
		return
	}

	fmt.Println(priKey)
	fmt.Println(pubKey)

	var l int32 = 16
	var nid int32 = 713
	var login_aes_key []byte = RandomStr(16)
	var pubKeyLen int32 = int32(len(pubKey))
	var username string = "yydhsb"
	var password string = GetMd5("te8780511")
	accountRequest := ManualAuthAccountRequest{
		Aes: &ManualAuthAccountRequest_AesKey{
			Len: &l,
			Key: login_aes_key,
		},
		Ecdh: &ManualAuthAccountRequest_Ecdh{
			Nid: &nid,
			EcdhKey: &ManualAuthAccountRequest_Ecdh_EcdhKey{
				Len: &pubKeyLen,
				Key: pubKey,
			},
		},
		UserName:  &username,
		Password1: &password,
		Password2: &password,
	}

	var uZero int32 = 0
	var uOne int32 = 1
	var uTwo int32 = 2
	var guid string = "Aff0aef642a31fc\000"
	var clientver int32 = 637927472
	var androidver string = "android-22"
	var imei string = "865166024671219"
	var softInfoXml string = "<softtype><lctmoc>0</lctmoc><level>1</level><k1>ARMv7 processor rev 1 (v7l) </k1><k2></k2><k3>5.1.1</k3><k4>865166024671219</k4><k5>460007337766541</k5><k6>89860012221746527381</k6><k7>d3151233cfbb4fd4</k7><k8>unknown</k8><k9>iPhone X</k9><k10>2</k10><k11>placeholder</k11><k12>0001</k12><k13>0000000000000001</k13><k14>01:61:19:58:78:d3</k14><k15></k15><k16>neon vfp swp half thumb fastmult edsp vfpv3 idiva idivt</k16><k18>e89b158e4bcf988ebd09eb83f5378e87</k18><k21>\"wireless\"</k21><k22></k22><k24>41:27:91:12:3e:14</k24><k26>0</k26><k30>\"wireless\"</k30><k33>com.tencent.mm</k33><k34>Android-x86/android_x86/x86:5.1.1/LMY48Z/denglibo08021647:userdebug/test-keys</k34><k35>vivo v3</k35><k36>unknown</k36><k37>iPhone</k37><k38>x86</k38><k39>android_x86</k39><k40>taurus</k40><k41>1</k41><k42>X</k42><k43>null</k43><k44>0</k44><k45></k45><k46></k46><k47>wifi</k47><k48>865166024671219</k48><k49>/data/data/com.tencent.mm/</k49><k52>0</k52><k53>0</k53><k57>1080</k57><k58></k58><k59>0</k59></softtype>"
	var clientSeqID string = "Aff0aef642a31fc2_1522827110765"
	var clientSeqIDSign string = "e89b158e4bcf988ebd09eb83f5378e87"
	var loginDeviceName string = "iPhone X"
	var deviceInfoXml string = "<deviceinfo><MANUFACTURER name=\"iPhone\"><MODEL name=\"X\"><VERSION_RELEASE name=\"5.1.1\"><VERSION_INCREMENTAL name=\"eng.denglibo.20171224.164708\"><DISPLAY name=\"android_x86-userdebug 5.1.1 LMY48Z eng.denglibo.20171224.164708 test-keys\"></DISPLAY></VERSION_INCREMENTAL></VERSION_RELEASE></MODEL></MANUFACTURER></deviceinfo>"
	var language string = "zh_CN"
	var timeZone string = "8.00"
	var deviceBrand string = "iPhone"
	var deviceModel string = "Xarmeabi-v7a"
	var osType string = "android-22"
	var realCountry string = "cn"
	deviceRequest := ManualAuthDeviceRequest{
		Login: &LoginInfo{
			AesKey:     login_aes_key,
			Uin:        &uZero,
			Guid:       &guid,
			ClientVer:  &clientver,
			AndroidVer: &androidver,
			Unknown:    &uOne,
		},
		Tag2:            &ManualAuthDeviceRequest__Tag2{},
		Imei:            &imei,
		SoftInfoXml:     &softInfoXml,
		Unknown5:        &uZero,
		ClientSeqID:     &clientSeqID,
		ClientSeqIDSign: &clientSeqIDSign,
		LoginDeviceName: &loginDeviceName,
		DeviceInfoXml:   &deviceInfoXml,
		Language:        &language,
		TimeZone:        &timeZone,
		Unknown13:       &uZero,
		Unknown14:       &uZero,
		DeviceBrand:     &deviceBrand,
		DeviceModel:     &deviceModel,
		OsType:          &osType,
		RealCountry:     &realCountry,
		Unknown22:       &uTwo,
	}

	accData, err := proto.Marshal(&accountRequest)
	if err != nil {
		fmt.Println("序列化accountRequest出错！")
		return
	}
	fmt.Println("-------------accData-------------")
	fmt.Println(accData)
	fmt.Println("-------------accData-------------")

	devData, err := proto.Marshal(&deviceRequest)
	if err != nil {
		fmt.Println("序列化deviceRequest出错！")
		return
	}
	fmt.Println("-------------devData-------------")
	fmt.Println(devData)
	fmt.Println("-------------devData-------------")

	reqAccount, _ := CompressAndRsaEnc(accData)
	reqDevice, _ := CompressAndAesEnc(devData, login_aes_key)
	fmt.Println(reqAccount)
	fmt.Println(reqDevice)
}
